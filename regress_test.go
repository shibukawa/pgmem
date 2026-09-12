package pgmem_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shibukawa/pgmem"
)

// contribRegress lists, per extension, the upstream regression files
// (contrib/<name>/sql and expected, PostgreSQL License) vendored under
// testdata/regress/<name>, in the order the contrib Makefile runs them.
var contribRegress = map[string][]string{
	// pgcrypto for a build with zlib: OpenPGP compression runs on the host
	// (wasm/pgmem_pgcrypto_compress.inc), so pgp-compression is in and
	// pgp-zlib-DISABLED is not.
	"pgcrypto": {
		"init", "md5", "sha1", "hmac-md5", "hmac-sha1", "blowfish", "rijndael",
		"sha2", "des", "3des", "cast5",
		"crypt-des", "crypt-md5", "crypt-blowfish", "crypt-xdes",
		"pgp-armor", "pgp-decrypt", "pgp-encrypt", "pgp-encrypt-md5", "pgp-compression",
		"pgp-pubkey-decrypt", "pgp-pubkey-encrypt", "pgp-pubkey-session",
		"pgp-info", "crypt-shacrypt",
	},
	"citext":  {"create_index_acl", "citext", "citext_utf8"},
	"pg_trgm": {"pg_trgm", "pg_utf8_trgm", "pg_word_trgm", "pg_strict_word_trgm"},
	"hstore":  {"hstore", "hstore_utf8"},
	"ltree":   {"ltree"},
	// without_overlaps uses \d, which the psql emulation does not have
	"btree_gist": {
		"init", "int2", "int4", "int8", "float4", "float8", "cash", "oid", "timestamp", "timestamptz",
		"time", "timetz", "date", "interval", "macaddr", "macaddr8", "inet", "cidr", "text", "varchar", "char",
		"bytea", "bit", "varbit", "numeric", "uuid", "not_equal", "enum", "bool", "partitions",
		"stratnum",
	},
	"btree_gin": {
		"install_btree_gin", "int2", "int4", "int8", "float4", "float8", "money", "oid",
		"timestamp", "timestamptz", "time", "timetz", "date", "interval",
		"macaddr", "macaddr8", "inet", "cidr", "text", "varchar", "char", "bytea", "bit", "varbit",
		"numeric", "enum", "uuid", "name", "bool", "bpchar",
	},
	"unaccent":      {"unaccent"},
	"tablefunc":     {"tablefunc"},
	"intarray":      {"_int"},
	"fuzzystrmatch": {"fuzzystrmatch", "fuzzystrmatch_utf8"},
	"cube":          {"cube", "cube_sci"},
	// earthdistance's file is cut before its extension create/drop part,
	// which is all \dT, \df, \do and \d
	"earthdistance": {"earthdistance"},
	// partition is cut before its two closing \d+ lines
	"seg":       {"security", "seg", "partition"},
	"bloom":     {"bloom"},
	"isn":       {"isn"},
	"dict_int":  {"dict_int"},
	"dict_xsyn": {"dict_xsyn"},
	"lo":        {"lo"},
}

// TestContribRegress replays PostgreSQL's own regression tests for every
// bundled extension against a fresh server each and compares with the
// upstream expected output, formatted the way psql -a -q prints it.
func TestContribRegress(t *testing.T) {
	names := make([]string, 0, len(contribRegress))
	for name := range contribRegress {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			// REGRESS_DSN points the same replay at another server (a real
			// PostgreSQL with the contrib modules installed) to tell a
			// pgmem difference from a harness one.
			dsn := os.Getenv("REGRESS_DSN")
			if dsn == "" {
				s, err := pgmem.Start(context.Background(), pgmem.Options{})
				if err != nil {
					t.Fatal(err)
				}
				defer s.Close()
				dsn = s.DSN()
			}
			runRegress(t, dsn, filepath.Join("testdata", "regress", name), contribRegress[name])
		})
	}
}

// runRegress executes dir/sql/<name>.sql for each name, each on a fresh
// connection the way pg_regress starts a psql per file, and compares the
// psql-style transcript with dir/expected/<name>.out.
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
	var cur *psqlState // the file being replayed, for its VERBOSITY
	cfg.OnNotice = func(_ *pgconn.PgConn, n *pgconn.Notice) {
		verbosity := "default"
		if cur != nil {
			verbosity = cur.verbosity
		}
		// libpq's default show_context=errors: no CONTEXT line for notices
		notices.WriteString(formatMessage((*pgconn.PgError)(n), verbosity, "", false))
	}
	// what pg_regress puts in the environment for every test session
	cfg.RuntimeParams["datestyle"] = "Postgres, MDY"
	cfg.RuntimeParams["timezone"] = "America/Los_Angeles"
	cfg.RuntimeParams["intervalstyle"] = "postgres_verbose"

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
			conn, err := pgx.ConnectConfig(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close(ctx)
			// the startup parameters only reach the first session of a
			// pgmem server; SET makes them stick for every connection
			for _, q := range []string{
				"SET datestyle = 'Postgres, MDY'",
				"SET timezone = 'America/Los_Angeles'",
				"SET intervalstyle = 'postgres_verbose'",
			} {
				if _, err := conn.Exec(ctx, q); err != nil {
					t.Fatal(err)
				}
			}
			cur = &psqlState{vars: map[string]string{}, verbosity: "default", dir: dir, ctx: ctx, conn: conn}
			got := psqlTranscript(cur, &notices, string(sqlText))
			if diff := transcriptDiff(string(want), got); diff != "" {
				out := filepath.Join(t.TempDir(), name+".out")
				os.WriteFile(out, []byte(got), 0o644)
				t.Errorf("output differs from expected (actual saved to %s):\n%s", out, diff)
			}
		})
		if !ok && name == names[0] {
			t.Fatal("first file failed; the rest would be noise")
		}
	}
}

// psqlState is the little of psql's scripting state the contrib tests use:
// variables (\\set, \\gset, :name interpolation), \\if/\\else/\\endif
// with \\quit, and the VERBOSITY setting for error display.
type psqlState struct {
	vars      map[string]string
	verbosity string
	ifStack   []bool // whether each open \\if branch is being executed
	quit      bool
	dir       string // test directory, for \\copy file paths
	ctx       context.Context
	conn      *pgx.Conn
}

func (p *psqlState) active() bool {
	for _, on := range p.ifStack {
		if !on {
			return false
		}
	}
	return true
}

// psqlTranscript echoes every input line and, after each complete
// statement, what psql -a -q would print for it: notices, the result table
// in aligned format, or the error.
func psqlTranscript(ps *psqlState, notices *strings.Builder, input string) string {
	var out strings.Builder
	var buf strings.Builder
	ctx, conn := ps.ctx, ps.conn
	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	exec := func(stmt string, gset bool, prefix string) {
		if !ps.active() {
			return
		}
		notices.Reset()
		result := runStatement(ctx, conn, ps, interpolate(leadingCommentsStripped(stmt), ps.vars), gset, prefix)
		out.WriteString(notices.String())
		out.WriteString(result)
	}
	for _, line := range lines {
		if ps.quit {
			break
		}
		// psql drops empty lines unless they sit inside a quoted string
		// (mainloop.c: "nothing left on line? then ignore").
		if line == "" && !endsInQuote(buf.String()) {
			continue
		}
		out.WriteString(line)
		out.WriteByte('\n')
		// a backslash command on its own line, or \\gset ending a query
		if strings.HasPrefix(strings.TrimSpace(line), "\\") && !endsInQuote(buf.String()) {
			cmd := strings.TrimSpace(line)
			if strings.HasPrefix(cmd, "\\gset") {
				exec(buf.String(), true, strings.TrimSpace(strings.TrimPrefix(cmd, "\\gset")))
				buf.Reset()
				continue
			}
			if strings.TrimSpace(stripComments(buf.String())) != "" {
				panic("regress: meta-command with a pending statement: " + line)
			}
			out.WriteString(ps.meta(cmd))
			continue
		}
		if i := strings.Index(line, "\\gset"); i >= 0 && !endsInQuote(buf.String()+line[:i]) {
			buf.WriteString(line[:i])
			exec(buf.String(), true, strings.TrimSpace(line[i+len("\\gset"):]))
			buf.Reset()
			continue
		}
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
			exec(stmt, false, "")
		}
	}
	return out.String()
}

// meta runs one backslash command and returns what it prints.
func (ps *psqlState) meta(cmd string) string {
	name, args, _ := strings.Cut(cmd, " ")
	args = strings.TrimSpace(args)
	switch name {
	case "\\if":
		on := ps.active() && psqlBool(interpolate(args, ps.vars))
		ps.ifStack = append(ps.ifStack, on)
	case "\\elif":
		if n := len(ps.ifStack); n > 0 {
			ps.ifStack[n-1] = !ps.ifStack[n-1] && psqlBool(interpolate(args, ps.vars))
		}
	case "\\else":
		if n := len(ps.ifStack); n > 0 {
			ps.ifStack[n-1] = !ps.ifStack[n-1]
		}
	case "\\endif":
		if n := len(ps.ifStack); n > 0 {
			ps.ifStack = ps.ifStack[:n-1]
		}
	case "\\quit", "\\q":
		if ps.active() {
			ps.quit = true
		}
	case "\\set":
		if !ps.active() {
			return ""
		}
		k, v, _ := strings.Cut(args, " ")
		v = strings.TrimSpace(v)
		if k == "VERBOSITY" {
			ps.verbosity = v
		} else {
			ps.vars[k] = interpolate(v, ps.vars)
		}
	case "\\unset":
		delete(ps.vars, args)
	case "\\echo":
		if ps.active() {
			return interpolate(args, ps.vars) + "\n"
		}
	case "\\copy":
		if ps.active() {
			return ps.clientCopy(args)
		}
	default:
		panic("regress: unsupported meta-command " + cmd)
	}
	return ""
}

var copyFromRe = regexp.MustCompile(`(?i)^(.+?)\s+from\s+'([^']+)'\s*(.*)$`)

// clientCopy runs "\\copy <table> from '<file>' [options]" as COPY FROM
// STDIN fed from the file, which is what psql does; it prints nothing.
func (ps *psqlState) clientCopy(args string) string {
	m := copyFromRe.FindStringSubmatch(strings.TrimSpace(args))
	if m == nil {
		panic("regress: unsupported \\copy form: " + args)
	}
	f, err := os.Open(filepath.Join(ps.dir, m[2]))
	if err != nil {
		return "CLIENT ERROR: " + err.Error() + "\n"
	}
	defer f.Close()
	sql := "COPY " + m[1] + " FROM STDIN " + m[3]
	if _, err := ps.conn.PgConn().CopyFrom(ps.ctx, f, sql); err != nil {
		return formatError(err, ps.verbosity, sql)
	}
	return ""
}

func psqlBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "t", "true", "on", "1", "yes", "y":
		return true
	}
	return false
}

// interpolate replaces :name, :'name' and :"name" with the variable's
// value the way psql does, leaving unknown names and ::casts alone.
func interpolate(s string, vars map[string]string) string {
	if len(vars) == 0 || !strings.Contains(s, ":") {
		return s
	}
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ':' || (i > 0 && (s[i-1] == ':' || isIdent(s[i-1]))) || i+1 >= len(s) || s[i+1] == ':' {
			b.WriteByte(c)
			continue
		}
		quote := byte(0)
		j := i + 1
		if s[j] == '\'' || s[j] == '"' {
			quote = s[j]
			j++
		}
		k := j
		for k < len(s) && isIdent(s[k]) {
			k++
		}
		name := s[j:k]
		if quote != 0 && (k >= len(s) || s[k] != quote) {
			name = ""
		}
		v, ok := vars[name]
		if name == "" || !ok {
			b.WriteByte(c)
			continue
		}
		switch quote {
		case '\'':
			b.WriteString("'" + strings.ReplaceAll(v, "'", "''") + "'")
			k++
		case '"':
			b.WriteString("\"" + strings.ReplaceAll(v, "\"", "\"\"") + "\"")
			k++
		default:
			b.WriteString(v)
		}
		i = k - 1
	}
	return b.String()
}

func isIdent(c byte) bool {
	return c == '_' || c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// scanSQL walks s and returns the index of the first ';' outside quotes,
// dollar quotes and comments (-1 if none) and whether s ends inside a
// quoted or dollar-quoted string.
func scanSQL(s string) (semi int, inQuote bool) {
	i := 0
	// psql's begin_depth: inside BEGIN ATOMIC ... END a ';' does not end
	// the statement; CASE ... END nests within it.
	depth := 0
	lastWord := ""
	for i < len(s) {
		c := s[i]
		switch {
		case isIdent(c) && !(c >= '0' && c <= '9'):
			j := i
			for j < len(s) && isIdent(s[j]) {
				j++
			}
			word := strings.ToLower(s[i:j])
			switch {
			case word == "atomic" && lastWord == "begin":
				depth++
			case word == "case" && depth > 0:
				depth++
			case word == "end" && depth > 0:
				depth--
			}
			lastWord = word
			i = j
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
			if depth == 0 {
				return i, false
			}
			i++
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

func runStatement(ctx context.Context, conn *pgx.Conn, ps *psqlState, stmt string, gset bool, prefix string) string {
	// The raw simple-query path: the text goes to the server untouched
	// (pgx's Query would try to bind $n placeholders itself) and several
	// statements in one string produce several results, like psql.
	var out strings.Builder
	errShown := false
	mrr := conn.PgConn().Exec(ctx, stmt)
	for mrr.NextResult() {
		rr := mrr.ResultReader()
		fields := append([]pgconn.FieldDescription(nil), rr.FieldDescriptions()...)
		var data [][]string
		for rr.NextRow() {
			raw := rr.Values()
			row := make([]string, len(raw))
			for i, v := range raw {
				row[i] = string(v)
			}
			data = append(data, row)
		}
		if _, err := rr.Close(); err != nil {
			out.WriteString(formatError(err, ps.verbosity, stmt))
			errShown = true
			break
		}
		if len(fields) == 0 {
			continue
		}
		if gset {
			// \\gset [prefix]: one row's columns become variables, no output
			if len(data) != 1 {
				fmt.Fprintf(&out, "CLIENT ERROR: \\gset with %d rows\n", len(data))
				continue
			}
			for i, f := range fields {
				ps.vars[prefix+f.Name] = data[0][i]
			}
			continue
		}
		out.WriteString(formatAligned(fields, data))
	}
	if err := mrr.Close(); err != nil && !errShown {
		out.WriteString(formatError(err, ps.verbosity, stmt))
	}
	return out.String()
}

func formatError(err error, verbosity string, query string) string {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return "CLIENT ERROR: " + err.Error() + "\n"
	}
	return formatMessage(pgErr, verbosity, query, true)
}

// formatMessage renders an error or notice the way psql prints it at the
// given VERBOSITY; withContext adds the CONTEXT line (errors only, as
// libpq's default show_context=errors).
func formatMessage(pgErr *pgconn.PgError, verbosity string, query string, withContext bool) string {
	var b strings.Builder
	if verbosity == "sqlstate" {
		fmt.Fprintf(&b, "%s:  %s\n", pgErr.Severity, pgErr.Code)
		return b.String()
	}
	fmt.Fprintf(&b, "%s:  %s\n", pgErr.Severity, pgErr.Message)
	if verbosity == "terse" {
		return b.String()
	}
	if pgErr.Position > 0 {
		b.WriteString(errorPosition(query, int(pgErr.Position)))
	}
	if pgErr.Detail != "" {
		fmt.Fprintf(&b, "DETAIL:  %s\n", pgErr.Detail)
	}
	if pgErr.Hint != "" {
		fmt.Fprintf(&b, "HINT:  %s\n", pgErr.Hint)
	}
	if pgErr.Where != "" && withContext {
		fmt.Fprintf(&b, "CONTEXT:  %s\n", pgErr.Where)
	}
	return b.String()
}

// errorPosition renders libpq's "LINE n: ..." plus caret for a 1-based
// character position in query, following reportErrorPosition in
// fe-protocol3.c: tabs become spaces, the line is cut to 60 screen
// columns (right first, then left while keeping the cursor 10 columns
// from the right edge) with "..." marking the cuts. Every character is
// taken as one column wide.
func errorPosition(query string, loc int) string {
	const displaySize, minRightCut = 60, 10
	loc--
	if loc < 0 {
		return ""
	}
	runes := []rune(strings.ReplaceAll(query, "\t", " "))
	locLine := 1
	ibeg := 0
	iend := -1
	for cno := 0; cno < len(runes); cno++ {
		ch := runes[cno]
		if ch == '\r' || ch == '\n' {
			if cno < loc {
				if ch == '\r' || cno == 0 || runes[cno-1] != '\r' {
					locLine++
				}
				ibeg = cno + 1
			} else {
				iend = cno
				break
			}
		}
	}
	if iend < 0 {
		iend = len(runes)
	}
	if loc > len(runes) {
		return ""
	}
	begTrunc, endTrunc := false, false
	if iend-ibeg > displaySize {
		if ibeg+displaySize >= loc+minRightCut {
			iend = ibeg + displaySize
			endTrunc = true
		} else {
			for loc+minRightCut < iend {
				iend--
				endTrunc = true
			}
			for iend-ibeg > displaySize {
				ibeg++
				begTrunc = true
			}
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "LINE %d: ", locLine)
	if begTrunc {
		b.WriteString("...")
	}
	prefixWidth := utf8.RuneCountInString(b.String())
	b.WriteString(string(runes[ibeg:iend]))
	if endTrunc {
		b.WriteString("...")
	}
	b.WriteString("\n")
	b.WriteString(strings.Repeat(" ", prefixWidth+loc-ibeg))
	b.WriteString("^\n")
	return b.String()
}

// leadingCommentsStripped drops the blank and comment-only lines before a
// statement, as psql's query buffer starts at the first token; error
// positions and LINE numbers count from there.
func leadingCommentsStripped(stmt string) string {
	lines := strings.Split(stmt, "\n")
	i := 0
	for i < len(lines) {
		t := strings.TrimSpace(lines[i])
		if t == "" || strings.HasPrefix(t, "--") {
			i++
			continue
		}
		break
	}
	return strings.Join(lines[i:], "\n")
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
