package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_word_similarity_dist_commutator_op(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14408(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
