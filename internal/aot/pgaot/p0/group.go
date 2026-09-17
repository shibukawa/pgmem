package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformGroupClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+27)) = uint8(v8)
	v25 = F_flatten_grouping_sets(m, l1, int32(1), v16+int32(27))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v47 == int32(0) {
		v127 = v8
		goto L8
	} else {
		goto L9
	}
L2:
	;
	return int32(0)
L3:
	;
	if v25 != 0 {
		v47 = v25
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+27)))
	if v29&int32(1) == int32(0) {
		v47 = v25
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = int32(0)
	v36 = F_exprLocation(m, l1)
	mBase = m.M
	v37 = F_makeGroupingSet(m, v34, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v37
	v44 = F_list_make1_impl(m, int32(1), v16+int32(12))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v47 = v44
	goto L1
L8:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 <= int32(0) {
		v127 = v8
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v63 = v8
	v64 = v8
	v65 = v8
	goto L11
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v63<<(uint(int32(2))%32))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 == int32(107) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v127 = v110
	goto L8
L13:
	;
	v112 = v63 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v112 < v113 {
		v63 = v112
		v64 = v109
		v65 = v110
		goto L11
	} else {
		goto L27
	}
L14:
	;
	v105 = F_lappend(m, v65, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	switch v74 {
	case 0:
		v102 = v70
		v104 = v64
		goto L14
	default:
		v109 = v64
		v110 = v65
		goto L13
	case 2, 3, 4:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v82 = F_transformGroupClauseExpr(m, v16+int32(28), v64, l0, v70, l3, l4, l5, l6, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v77 = F_transformGroupingSet(m, v16+int32(28), l0, v70, l3, l4, l5, l6)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v102 = v77
	v104 = v64
	goto L14
L20:
	;
	if v82 == int32(0) {
		v109 = v64
		v110 = v65
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v86 = F_bms_add_member(m, v64, v82)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+27)))
	if v88 != int32(1) {
		v109 = v86
		v110 = v65
		goto L13
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v82
	v97 = F_list_make1_impl(m, int32(471), v16+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v99 = F_exprLocation(m, v70)
	mBase = m.M
	v100 = F_makeGroupingSet(m, int32(1), v97, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v102 = v100
	v104 = v86
	goto L14
L26:
	;
	v109 = v104
	v110 = v105
	goto L13
L27:
	;
	goto L12
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v127
	goto L30
L29:
	;
	goto L30
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	m.G0 = v16 + int32(32)
	return v129
}
