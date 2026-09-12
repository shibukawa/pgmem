package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_ChangeVarNodesWalkExpression(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_expression_tree_walker_impl(m, l0, int32(1049), l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
