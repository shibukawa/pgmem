package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_word(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v8 = l0
	goto L3
L1:
	;
	return v42
L2:
	;
	v30 = v13
	v31 = v8
	goto L11
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32(v13-int32(9)) < base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v13 == int32(0) {
		v42 = v3
		goto L1
	} else {
		goto L10
	}
L5:
	;
	goto L4
L6:
	;
	v21 = F_pg_mblen_cstr(m, v8)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	switch v13 - int32(32) {
	case 0:
		goto L6
	case 1, 2:
		goto L2
	case 3:
		v42 = v3
		goto L1
	default:
		goto L5
	}
L8:
	;
	return int32(0)
L9:
	;
	v8 = v21 + v8
	goto L3
L10:
	;
	goto L2
L11:
	;
	switch v30 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L13
	default:
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
	v42 = v8
	goto L1
L13:
	;
	goto L12
L14:
	;
	v33 = F_pg_mblen_cstr(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v35 = v33 + v31
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v30 = v36
	v31 = v35
	goto L11
}
func F_word_similarity(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14021(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_word_similarity_commutator_op(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14022(m, l0, int32(_a_F_word_similarity_commutator_op_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_word_similarity_op(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14025(m, l0, int32(_a_F_word_similarity_op_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
