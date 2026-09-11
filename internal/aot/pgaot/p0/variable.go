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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v40
L2:
	;
	v40 = int32(1)
	goto L1
L3:
	;
	v6 = int32(0)
	if v4 == v6 {
		v40 = v6
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
		v40 = int32(0)
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v12 == int32(0) {
		v31 = v11
		v32 = v12
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v32-v31 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v11 != v12 {
		v31 = v11
		v32 = v12
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v16 = v5
	v17 = v4
	goto L11
L11:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v20
		v32 = v21
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v31 = v20
	v32 = v21
	goto L8
L13:
	;
	v24 = int32(1)
	if v20 == v21 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v40 = v6
	goto L1
L16:
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
