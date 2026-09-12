package pgmem_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shibukawa/pgmem"
)

// The pgcrypto regression suite from the PostgreSQL source tree
// (contrib/pgcrypto/sql and expected, PostgreSQL License), in the order the
// contrib Makefile runs it for a build with zlib: OpenPGP compression runs
// on the host (wasm/pgmem_pgcrypto_compress.inc), so pgp-compression is in
// and pgp-zlib-DISABLED is not.
var pgcryptoRegress = []string{
	"init", "md5", "sha1", "hmac-md5", "hmac-sha1", "blowfish", "rijndael",
	"sha2", "des", "3des", "cast5",
	"crypt-des", "crypt-md5", "crypt-blowfish", "crypt-xdes",
	"pgp-armor", "pgp-decrypt", "pgp-encrypt", "pgp-encrypt-md5", "pgp-compression",
	"pgp-pubkey-decrypt", "pgp-pubkey-encrypt", "pgp-pubkey-session",
	"pgp-info", "crypt-shacrypt",
}

// TestPgcryptoRegress replays PostgreSQL's own pgcrypto regression tests
// against the host-backed crypto (wasm/pgmem_pgcrypto_*.inc,
// internal/host/crypto.go) and compares with the upstream expected output,
// formatted the way psql -a -q prints it.
func TestPgcryptoRegress(t *testing.T) {
	s, err := pgmem.Start(context.Background(), pgmem.Options{})
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	runRegress(t, s.DSN(), filepath.Join("testdata", "regress", "pgcrypto"), pgcryptoRegress)
}

// runRegress executes dir/sql/<name>.sql for each name on one connection
// and compares the psql-style transcript with dir/expected/<name>.out.
func runRegress(t *testing.T, dsn, dir string, names []string) {
	t.Helper()
	ctx := context.Background()
	cfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		t.Fatal(err)
	}
	// psql speaks the simple query protocol; results come back as text.
	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	var notices strings.Builder
	cfg.OnNotice = func(_ *pgconn.PgConn, n *pgconn.Notice) {
		fmt.Fprintf(&notices, "%s:  %s\n", n.Severity, n.Message)
	}
	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	for _, name := range names {
		ok := t.Run(name, func(t *testing.T) {
			sqlText, err := os.ReadFile(filepath.Join(dir, "sql", name+".sql"))
			if err != nil {
				t.Fatal(err)
			}
			want, err := os.ReadFile(filepath.Join(dir, "expected", name+".out"))
			if err != nil {
				t.Fatal(err)
			}
			got := psqlTranscript(ctx, conn, &notices, string(sqlText))
			if diff := transcriptDiff(string(want), got); diff != "" {
				out := filepath.Join(t.TempDir(), name+".out")
				os.WriteFile(out, []byte(got), 0o644)
				t.Errorf("output differs from expected (actual saved to %s):\n%s", out, diff)
			}
		})
		if !ok && name == "init" {
			t.Fatal("init failed; the rest would be noise")
		}
	}
}

// psqlTranscript echoes every input line and, after each complete
// statement, what psql -a -q would print for it: notices, the result table
// in aligned format, or the error.
func psqlTranscript(ctx context.Context, conn *pgx.Conn, notices *strings.Builder, input string) string {
	var out strings.Builder
	var buf strings.Builder
	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	for _, line := range lines {
		// psql drops empty lines unless they sit inside a quoted string
		// (mainloop.c: "nothing left on line? then ignore").
		if line == "" && !endsInQuote(buf.String()) {
			continue
		}
		out.WriteString(line)
		out.WriteByte('\n')
		buf.WriteString(line)
		buf.WriteByte('\n')
		for {
			stmt, rest, found := splitStatement(buf.String())
			if !found {
				break
			}
			buf.Reset()
			buf.WriteString(rest)
			if strings.TrimSpace(stripComments(stmt)) == ";" {
				continue
			}
			notices.Reset()
			result := runStatement(ctx, conn, stmt)
			out.WriteString(notices.String())
			out.WriteString(result)
		}
	}
	return out.String()
}

// scanSQL walks s and returns the index of the first ';' outside quotes,
// dollar quotes and comments (-1 if none) and whether s ends inside a
// quoted or dollar-quoted string.
func scanSQL(s string) (semi int, inQuote bool) {
	i := 0
	for i < len(s) {
		c := s[i]
		switch {
		case c == '-' && i+1 < len(s) && s[i+1] == '-':
			nl := strings.IndexByte(s[i:], '\n')
			if nl < 0 {
				return -1, false
			}
			i += nl + 1
		case c == '/' && i+1 < len(s) && s[i+1] == '*':
			end := strings.Index(s[i+2:], "*/")
			if end < 0 {
				return -1, false
			}
			i += end + 4
		case c == '\'' || c == '"':
			j := i + 1
			for {
				k := strings.IndexByte(s[j:], c)
				if k < 0 {
					return -1, true
				}
				j += k + 1
				if j < len(s) && s[j] == c { // doubled quote
					j++
					continue
				}
				break
			}
			i = j
		case c == '$':
			// $tag$ ... $tag$ or a bare $$
			j := i + 1
			for j < len(s) && (s[j] == '_' || s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z' || s[j] >= '0' && s[j] <= '9') {
				j++
			}
			if j < len(s) && s[j] == '$' {
				tag := s[i : j+1]
				end := strings.Index(s[j+1:], tag)
				if end < 0 {
					return -1, true
				}
				i = j + 1 + end + len(tag)
			} else {
				i++
			}
		case c == ';':
			return i, false
		default:
			i++
		}
	}
	return -1, false
}

func endsInQuote(s string) bool {
	_, q := scanSQL(s)
	return q
}

// splitStatement splits s at the first statement terminator.
func splitStatement(s string) (stmt, rest string, found bool) {
	i, _ := scanSQL(s)
	if i < 0 {
		return "", "", false
	}
	return s[:i+1], s[i+1:], true
}

func stripComments(s string) string {
	var b strings.Builder
	for _, line := range strings.Split(s, "\n") {
		if i := strings.Index(line, "--"); i >= 0 {
			line = line[:i]
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}

var rightAligned = map[uint32]bool{20: true, 21: true, 23: true, 26: true, 700: true, 701: true, 790: true, 1700: true}

func runStatement(ctx context.Context, conn *pgx.Conn, stmt string) string {
	rows, err := conn.Query(ctx, stmt)
	if err != nil {
		return formatError(err)
	}
	fields := rows.FieldDescriptions()
	var data [][]string
	for rows.Next() {
		raw := rows.RawValues()
		row := make([]string, len(raw))
		for i, v := range raw {
			row[i] = string(v)
		}
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		return formatError(err)
	}
	if len(fields) == 0 {
		return ""
	}
	return formatAligned(fields, data)
}

func formatError(err error) string {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return "CLIENT ERROR: " + err.Error() + "\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s:  %s\n", pgErr.Severity, pgErr.Message)
	if pgErr.Detail != "" {
		fmt.Fprintf(&b, "DETAIL:  %s\n", pgErr.Detail)
	}
	if pgErr.Hint != "" {
		fmt.Fprintf(&b, "HINT:  %s\n", pgErr.Hint)
	}
	if pgErr.Where != "" {
		fmt.Fprintf(&b, "CONTEXT:  %s\n", pgErr.Where)
	}
	return b.String()
}

// formatAligned renders rows like psql's default aligned format with
// border 1: centered headers, a dashed separator, right-aligned numbers,
// '+' continuation marks on multi-line values and a "(N rows)" footer.
func formatAligned(fields []pgconn.FieldDescription, data [][]string) string {
	ncol := len(fields)
	widths := make([]int, ncol)
	cells := make([][][]string, len(data))
	for i, f := range fields {
		widths[i] = utf8.RuneCountInString(f.Name)
	}
	for r, row := range data {
		cells[r] = make([][]string, ncol)
		for c, v := range row {
			ls := strings.Split(psqlEscapeControl(v), "\n")
			cells[r][c] = ls
			for _, l := range ls {
				if w := utf8.RuneCountInString(l); w > widths[c] {
					widths[c] = w
				}
			}
		}
	}
	var b strings.Builder
	// header
	for c, f := range fields {
		if c > 0 {
			b.WriteString("|")
		}
		w := utf8.RuneCountInString(f.Name)
		left := (widths[c] - w) / 2
		b.WriteString(" ")
		b.WriteString(strings.Repeat(" ", left))
		b.WriteString(f.Name)
		b.WriteString(strings.Repeat(" ", widths[c]-w-left))
		b.WriteString(" ")
	}
	b.WriteString("\n")
	for c := range fields {
		if c > 0 {
			b.WriteString("+")
		}
		b.WriteString(strings.Repeat("-", widths[c]+2))
	}
	b.WriteString("\n")
	for r := range data {
		nlines := 1
		for _, ls := range cells[r] {
			if len(ls) > nlines {
				nlines = len(ls)
			}
		}
		for j := 0; j < nlines; j++ {
			for c := range fields {
				if c > 0 {
					b.WriteString("|")
				}
				ls := cells[r][c]
				text := ""
				if j < len(ls) {
					text = ls[j]
				}
				pad := strings.Repeat(" ", widths[c]-utf8.RuneCountInString(text))
				b.WriteString(" ")
				if rightAligned[fields[c].DataTypeOID] {
					b.WriteString(pad)
					b.WriteString(text)
				} else {
					b.WriteString(text)
					b.WriteString(pad)
				}
				if j+1 < len(ls) {
					b.WriteString("+")
				} else {
					b.WriteString(" ")
				}
			}
			b.WriteString("\n")
		}
	}
	if len(data) == 1 {
		b.WriteString("(1 row)\n\n")
	} else {
		fmt.Fprintf(&b, "(%d rows)\n\n", len(data))
	}
	return b.String()
}

// psqlEscapeControl renders control characters and invalid bytes the way
// psql's pg_wcsformat does: \r as "\r", other controls and 0x7F as \xNN,
// newlines kept as line breaks.
func psqlEscapeControl(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		switch {
		case r == utf8.RuneError && size == 1:
			fmt.Fprintf(&b, "\\x%02X", s[i])
		case r == '\n':
			b.WriteByte('\n')
		case r == '\r':
			b.WriteString("\\r")
		case r < 0x20 || r == 0x7f:
			fmt.Fprintf(&b, "\\x%02X", r)
		default:
			b.WriteRune(r)
		}
		i += size
	}
	return b.String()
}

// transcriptDiff compares line by line ignoring trailing whitespace and
// returns the first mismatch with context, or "" when equal.
func transcriptDiff(want, got string) string {
	norm := func(s string) []string {
		ls := strings.Split(s, "\n")
		for i := range ls {
			ls[i] = strings.TrimRight(ls[i], " \t\r")
		}
		for len(ls) > 0 && ls[len(ls)-1] == "" {
			ls = ls[:len(ls)-1]
		}
		return ls
	}
	w, g := norm(want), norm(got)
	n := len(w)
	if len(g) > n {
		n = len(g)
	}
	for i := 0; i < n; i++ {
		var wl, gl string
		if i < len(w) {
			wl = w[i]
		}
		if i < len(g) {
			gl = g[i]
		}
		if wl != gl {
			var b strings.Builder
			fmt.Fprintf(&b, "line %d:\n", i+1)
			for j := i - 3; j <= i+3; j++ {
				if j < 0 {
					continue
				}
				if j < len(w) {
					fmt.Fprintf(&b, "  want %4d| %s\n", j+1, w[j])
				}
				if j < len(g) {
					fmt.Fprintf(&b, "  got  %4d| %s\n", j+1, g[j])
				}
			}
			return b.String()
		}
	}
	return ""
}
