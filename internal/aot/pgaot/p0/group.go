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
	var v31 int32
	_ = v31
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
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
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	v63 = v8
	v65 = v8
	v66 = v8
	goto L14
L3:
	;
	v125 = v8
	goto L1
L4:
	;
	return int32(0)
L5:
	;
	if v25 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+27)))
	if v31 != int32(1) {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v49 = v25
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if int32(0) < v50 {
		goto L2
	} else {
		goto L13
	}
L9:
	;
	v34 = int32(0)
	v36 = F_exprLocation(m, l1)
	mBase = m.M
	v37 = F_makeGroupingSet(m, v34, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v37
	v44 = F_list_make1_impl(m, int32(1), v16+int32(12))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v44 == int32(0) {
		v125 = v8
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v49 = v44
	goto L8
L13:
	;
	goto L3
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v65<<(uint(int32(2))%32))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 == int32(107) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v125 = v110
	goto L1
L16:
	;
	v114 = v65 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v114 < v115 {
		v63 = v110
		v65 = v114
		v66 = v111
		goto L14
	} else {
		goto L30
	}
L17:
	;
	v107 = F_lappend(m, v63, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L29
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	switch v76 {
	case 0:
		v104 = v72
		v105 = v66
		goto L17
	default:
		v110 = v63
		v111 = v66
		goto L16
	case 2, 3, 4:
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v84 = F_transformGroupClauseExpr(m, v16+int32(28), v66, l0, v72, l3, l4, l5, l6, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v79 = F_transformGroupingSet(m, v16+int32(28), l0, v72, l3, l4, l5, l6)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v104 = v79
	v105 = v66
	goto L17
L23:
	;
	if v84 == int32(0) {
		v110 = v63
		v111 = v66
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v88 = F_bms_add_member(m, v66, v84)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+27)))
	if v90 != int32(1) {
		v110 = v63
		v111 = v88
		goto L16
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v84
	v99 = F_list_make1_impl(m, int32(471), v16+int32(8))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v101 = F_exprLocation(m, v72)
	mBase = m.M
	v102 = F_makeGroupingSet(m, int32(1), v99, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v104 = v102
	v105 = v88
	goto L17
L29:
	;
	v110 = v107
	v111 = v105
	goto L16
L30:
	;
	goto L15
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v125
	goto L33
L32:
	;
	goto L33
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	m.G0 = v16 + int32(32)
	return v131
}
