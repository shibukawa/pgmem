package pgmem

import (
	"context"
	"errors"
	"fmt"
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
	if s.cl != nil {
		// the cluster is stopped and started on the copy; the sessions get
		// a new backend on their next message (see cluster.go)
		return s.restartCluster(ctx, sn.fs.Clone())
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
	if err := s.restart(sn.fs.Clone(), s.opts.Database); err != nil {
		if errors.Is(err, errServerClosed) {
			return err
		}
		return fmt.Errorf("pgmem: restore: %w", err)
	}
	s.restored, s.dirty = sn, false
	return nil
}
