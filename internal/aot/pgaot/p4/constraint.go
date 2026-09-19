package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v7 = F_list_copy(m, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		v51 = v7
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v56 = F_predicate_implied_by(m, l1, v51, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L4:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)))
	if v15 == int32(0) {
		v51 = v7
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(0)
	v21 = v7
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v28 = v25 + v19*int32(12)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	if v29 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v51 = v45
	goto L3
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = F_stringToNode(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v45 = v21
	goto L10
L10:
	;
	v47 = v19 + int32(1)
	if v47 != v15 {
		v19 = v47
		v21 = v45
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v36 = F_eval_const_expressions(m, int32(0), v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v39 = F_canonicalize_qual(m, v36, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v41 = F_make_ands_implicit(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v43 = F_list_concat(m, v21, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v45 = v43
	goto L10
L16:
	;
	goto L7
L17:
	;
	return v56
}
func F_get_constraint_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13917(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
