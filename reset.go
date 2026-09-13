package pgmem

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// Reset returns a fork to the state of the snapshot it was started from;
// see Restore. It fails on a server that Snapshot.Fork did not start.
func (s *Server) Reset(ctx context.Context) error {
	if s.origin == nil {
		return errors.New("pgmem: Reset needs a server started by Snapshot.Fork; use Restore")
	}
	return s.Restore(ctx, s.origin)
}

// Restore replaces the server's data with a copy of sn while the server
// keeps its port and its client connections, so a connection string or a
// connection pool captured earlier stays valid. It waits for open
// transactions to end (ctx bounds the wait) and returns at once when
// nothing has run since the server was last restored from sn.
//
// Live connections continue in a fresh session, as if their pool had
// reconnected: SET values and temp tables are gone. pgmem re-creates their
// named prepared statements and LISTEN registrations, so a driver that
// prepared a statement once keeps using it; a statement that no longer
// parses against sn fails when it is next used.
func (s *Server) Restore(ctx context.Context, sn *Snapshot) error {
	if sn.opts.Database != s.opts.Database || sn.opts.User != s.opts.User {
		return fmt.Errorf("pgmem: the snapshot serves database %q as %q, this server %q as %q",
			sn.opts.Database, sn.opts.User, s.opts.Database, s.opts.User)
	}
	if err := s.acquire(ctx, 0, false); err != nil {
		if errors.Is(err, errServerClosed) {
			return err
		}
		return fmt.Errorf("pgmem: restore waited for an open transaction to end (commit or close every connection first): %w", err)
	}
	defer s.release()
	if !s.dirty && s.restored == sn {
		return nil
	}
	fs := sn.fs.Clone()
	b, err := s.startBackend(fs)
	if err != nil {
		return err
	}
	if s.startupPkt != nil {
		resp, err := b.Startup(s.startupPkt)
		if err == nil {
			err = checkNoError(resp)
		}
		if err != nil {
			b.Close()
			return fmt.Errorf("pgmem: restore: startup: %w", err)
		}
	}
	s.bmu.Lock()
	if s.closed.Load() {
		s.bmu.Unlock()
		b.Close()
		return errServerClosed
	}
	old := s.b
	s.fs, s.b = fs, b
	s.bmu.Unlock()
	old.Close()
	s.restored, s.dirty = sn, false
	if s.startupPkt == nil {
		return nil // no client has connected yet; the first one starts the session
	}
	if err := s.applyUser(); err != nil {
		return fmt.Errorf("pgmem: restore: %w", err)
	}
	s.reestablish()
	return nil
}

// reestablish re-creates on a new backend what live connections had set
// up on the one it replaced: LISTEN registrations and named prepared
// statements. Runs with the backend held.
func (s *Server) reestablish() {
	clear(s.relisten)
	if len(s.listeners) > 0 {
		var q strings.Builder
		for ch := range s.listeners {
			q.WriteString("LISTEN " + quoteIdent(ch) + ";")
		}
		s.quietListen = true
		s.b.Exec(simpleQuery(q.String()))
		s.quietListen = false
	}
	for _, sess := range s.sessions {
		for _, parse := range sess.parses {
			s.b.Exec(append(append([]byte(nil), parse...), syncMessage...))
		}
	}
}
