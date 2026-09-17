package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalVariableShowStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v41
L2:
	;
	v41 = int32(1)
	goto L1
L3:
	;
	v6 = int32(0)
	if v4 == v6 {
		v41 = v6
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v4 != v5 {
		v41 = int32(0)
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if base.B2i32(v11 == int32(0))|base.B2i32(v11 != v14) != 0 {
		v32 = v11
		v33 = v14
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v32-v33 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v17 = v5
	v18 = v4
	goto L10
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v22
		v33 = v21
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v32 = v22
	v33 = v21
	goto L8
L12:
	;
	v25 = int32(1)
	if v22 == v21 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v41 = v6
	goto L1
L15:
	;
	goto L2
}
func F_check_variable_parameters(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if int32(0) < v5 {
		v10 = F_query_tree_walker_impl(m, l1, int32(495), l0, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
