package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_rule_orderby(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v23, int32(715480))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = F_get_rule_sortgroupclause(m, v27, l1, l2, l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v30 = F_exprType(m, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v33 = F_lookup_type_cache(m, v30, int32(6))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	if v35 != v36 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v64 = int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v65 <= v64 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	F_appendStringInfoString(m, v23, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	if v38 != v35 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	if v57 != int32(1) {
		goto L9
	} else {
		goto L24
	}
L14:
	;
	v40 = F_generate_operator_name(m, v35, v30, v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_appendStringInfoString(m, v23, int32(520327))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L22
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v40
	F_appendStringInfo(m, v23, int32(187937), v14+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(493466)
	goto L21
L20:
	;
	v51 = int32(493603)
	goto L21
L21:
	;
	v61 = v51
	goto L10
L22:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	if v55 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v61 = int32(493603)
	goto L10
L24:
	;
	v61 = int32(493466)
	goto L10
L25:
	;
	goto L9
L26:
	;
	v74 = v64
	goto L27
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v74<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v23, int32(704244))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L1
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v88 = F_get_rule_sortgroupclause(m, v87, l1, l2, l3)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v90 = F_exprType(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v93 = F_lookup_type_cache(m, v90, int32(6))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	if v95 == v96 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v123 = v74 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v123 < v124 {
		v74 = v123
		goto L27
	} else {
		goto L50
	}
L34:
	;
	F_appendStringInfoString(m, v23, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L49
	}
L35:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v98 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+60))
	if v102 == v95 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v119 = int32(493466)
	goto L34
L39:
	;
	F_appendStringInfoString(m, v23, int32(520327))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v109 = F_generate_operator_name(m, v95, v90, v90)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v107 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v119 = int32(493603)
	goto L34
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v109
	F_appendStringInfo(m, v23, int32(187937), v14)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+17)))
	if v117 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v118 = int32(493466)
	goto L48
L47:
	;
	v118 = int32(493603)
	goto L48
L48:
	;
	v119 = v118
	goto L34
L49:
	;
	goto L33
L50:
	;
	goto L28
}
