package pgmem

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shibukawa/pgmem/internal/vfs"
)

// A single-user backend serves one database, but a data directory can hold
// many (CREATE DATABASE works on the live session). The server serves them
// one at a time: a connection to another database waits for the backend
// like any other connection and then has the backend restarted on its own
// database. The switch shuts the old backend down with a checkpoint and
// takes about 10 ms; connections of the other databases stay open and get
// their prepared statements and LISTEN registrations back when their
// database is served again. This is what Prisma's shadow database and a
// database per test worker need.

// dbState is the session bookkeeping of one database. Guarded by sem.
type dbState struct {
	// startupPkt is the startup packet of the first connection to the
	// database, replayed whenever a backend is started on it again;
	// startupResp is what later connections are sent.
	startupPkt  []byte
	startupResp []byte
	live        int // sessions started and not yet ended
}

// listenKey is a notification channel of one database.
type listenKey struct{ database, channel string }

// noDatabaseError reports a connection to a database the server does not have.
type noDatabaseError struct{ name string }

func (e *noDatabaseError) Error() string { return fmt.Sprintf("database %q does not exist", e.name) }

// db returns the bookkeeping of database name, creating it. Runs with the
// backend held.
func (s *Server) db(name string) *dbState {
	st := s.dbs[name]
	if st == nil {
		st = &dbState{}
		s.dbs[name] = st
	}
	return st
}

// use makes the backend serve database name, restarting it on name when it
// serves another database. If name cannot be served the backend goes back
// to the database it served. Runs with the backend held and no transaction
// open.
func (s *Server) use(name string) error {
	if name == s.current {
		return nil
	}
	prev := s.current
	err := s.restart(s.fs, name)
	if err == nil || errors.Is(err, errServerClosed) {
		return err
	}
	if rerr := s.restart(s.fs, prev); rerr != nil {
		return fmt.Errorf("pgmem: starting database %q: %v; restarting %q: %w", name, err, prev, rerr)
	}
	if strings.Contains(err.Error(), "does not exist") {
		return &noDatabaseError{name: name}
	}
	return fmt.Errorf("pgmem: starting database %q: %w", name, err)
}

// restart replaces the backend with one serving database name on fs. fs is
// either the current filesystem (a database switch: the old backend is shut
// down first, as two backends must never share a data directory) or a new
// copy (Restore: the old backend runs until the new one is up). The new
// backend gets the database's first startup packet again, and the live
// connections of the database get their prepared statements and LISTEN
// registrations back. Runs with the backend held.
func (s *Server) restart(fs *vfs.FS, name string) error {
	s.bmu.Lock()
	closed, old, same := s.closed.Load(), s.b, fs == s.fs
	s.bmu.Unlock()
	if closed {
		return errServerClosed
	}
	if same {
		old.Close() // a shutdown checkpoint, so the next backend starts clean
	}
	b, err := s.startBackend(fs, name)
	if err != nil {
		return err
	}
	st := s.db(name)
	if st.startupPkt != nil {
		resp, err := b.Startup(st.startupPkt)
		if err == nil {
			err = checkNoError(resp)
		}
		if err != nil {
			b.Close()
			return fmt.Errorf("startup: %w", err)
		}
	}
	s.bmu.Lock()
	if s.closed.Load() {
		s.bmu.Unlock()
		b.Close()
		return errServerClosed
	}
	s.fs, s.b, s.current = fs, b, name
	s.bmu.Unlock()
	if !same {
		old.Close()
	}
	if st.startupPkt == nil {
		return nil // nobody has connected to the database yet; the first connection starts the session
	}
	if err := s.applyUser(); err != nil {
		return err
	}
	s.reestablish(name)
	return nil
}

// reestablish re-creates on a new backend what the live connections of
// database name had set up on the previous one: LISTEN registrations and
// named prepared statements. Runs with the backend held.
func (s *Server) reestablish(name string) {
	var q strings.Builder
	for k := range s.relisten {
		if k.database == name {
			delete(s.relisten, k)
		}
	}
	for k := range s.listeners {
		if k.database == name {
			q.WriteString("LISTEN " + quoteIdent(k.channel) + ";")
		}
	}
	if q.Len() > 0 {
		s.quietListen = true
		s.b.Exec(simpleQuery(q.String()))
		s.quietListen = false
	}
	for _, sess := range s.sessions {
		if sess.database != name {
			continue
		}
		for _, parse := range sess.parses {
			s.b.Exec(append(append([]byte(nil), parse...), syncMessage...))
		}
	}
}
