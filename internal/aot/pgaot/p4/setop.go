package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_setop_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v8 = int32(0)
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v16)
	v27 = v8
	v31 = v16
	v32 = v8
	goto L1
L1:
	;
	v34 = int32(0)
	if l0 == v34 {
		v44 = v34
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v32
L3:
	;
	v45 = int32(0)
	if l1 == v45 {
		v56 = v45
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v27 {
		v44 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = v40 + v27<<(uint(int32(2))%32)
	goto L3
L6:
	;
	if l4 == int32(0) {
		v65 = v45
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 <= v27 {
		v56 = int32(0)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v56 = v52 + v27<<(uint(int32(2))%32)
	goto L6
L9:
	;
	v66 = int32(0)
	if l5 == v66 {
		v75 = v66
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v59 <= v27 {
		v65 = v45
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v65 = v61 + v27<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v76 = int32(0)
	if base.B2i32(v44 == v76)|base.B2i32(v56 == v76)|(base.B2i32(v65 == v76)|base.B2i32(v75 == v76)) == v76 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v69 <= v27 {
		v75 = v66
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v75 = v71 + v27<<(uint(int32(2))%32)
	goto L12
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if l3 == int32(0) {
		v102 = v93
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	goto L2
L18:
	;
	v118 = F_exprType(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L25
	} else {
		goto L30
	}
L19:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92)+8)))
	v104 = F_exprType(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	if v93 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v102 = int32(0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v99 == int32(7) {
		v117 = v93
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v102 = v93
	goto L19
L25:
	;
	return int32(0)
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v109 = F_exprTypmod(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v112 = F_exprCollation(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v115 = F_makeVar(m, l2, v103, v104, v109, v112, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v117 = v115
	goto L18
L30:
	;
	if v118 != v91 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v123 = F_coerce_to_common_type(m, int32(0), v117, v91, int32(_a_F_generate_setop_tlist_0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	v127 = v117
	goto L33
L33:
	;
	v128 = F_exprCollation(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L25
	} else {
		goto L35
	}
L34:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v125)
	v127 = v123
	goto L33
L35:
	;
	if v128 != v90 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = F_exprType(m, v127)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L25
	} else {
		goto L39
	}
L37:
	;
	v142 = v127
	goto L38
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v145 = F_pstrdup(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L42
	}
L39:
	;
	v133 = F_exprTypmod(m, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v138 = F_applyRelabelType(m, v127, v131, v133, v90, int32(2), int32(-1), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v140)
	v142 = v138
	goto L38
L42:
	;
	v148 = F_makeTargetEntry(m, v142, base.I32_extend16_s(v31), v145, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v150
	v152 = int32(1)
	v156 = F_lappend(m, v32, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v27 = v27 + v152
	v31 = v31 + v152
	v32 = v156
	goto L1
}
func F_setop_has_grouping(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v36
L2:
	;
	v36 = int32(0)
	goto L1
L3:
	;
	v5 = l0
	goto L4
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v7 != int32(142) {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	goto L2
L6:
	;
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	if v11 != 0 {
		v36 = v10
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v12 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v30 != 0 {
		v36 = v10
		goto L1
	} else {
		goto L18
	}
L9:
	;
	goto L8
L10:
	;
	v30 = int32(0)
	goto L9
L11:
	;
	v16 = v12
	goto L12
L12:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v18 != int32(142) {
		goto L10
	} else {
		goto L14
	}
L13:
	;
	goto L10
L14:
	;
	v21 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v22 != 0 {
		v30 = v21
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v24 = F_setop_has_grouping(m, v23)
	mBase = m.M
	if v24 != 0 {
		v30 = v21
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v25 != 0 {
		v16 = v25
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v31 != 0 {
		v5 = v31
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L5
}
func F_setop_load_group(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, v8, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v9&int32(2) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	m.T0[v16].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	goto L1
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_ExecStoreMinimalTuple(m, v24, v26, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v35 = int64(1)
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ExecReScan(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
	if v41 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v46&int32(2) != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = F_setop_compare_slots(m, v49, v41, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v50 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v52 + int64(1)
	goto L11
}
