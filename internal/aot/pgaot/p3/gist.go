package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistCompressValues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v5 = l4
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	if v18 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v5 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v84 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v30 = int32(0)
	goto L5
L5:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+l3))))
	if v43 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v84 = v80
	goto L1
L7:
	;
	v80 = v30 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+10)))
	if v80 < v82 {
		v30 = v80
		goto L5
	} else {
		goto L16
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5+v30<<(uint(int32(3))%32)))) = int64(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v52 = v30 << (uint(int32(3)) % 32)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l2+v52)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)) = uint8(v5)
	v56 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+24)) = uint16(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v54
	v65 = l0 + int32(1812) + v30*int32(28)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v66 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_gistCompressValues_0)+v30<<(uint(int32(2))%32))))
	v71 = F_FunctionCall1Coll(m, v65, v70, base.I64_extend_i32_u(v15+int32(8)))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v75 = v54
	goto L13
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5+v52))) = v75
	goto L7
L14:
	;
	return
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v71))))
	v75 = v74
	goto L13
L16:
	;
	goto L6
L17:
	;
	m.G0 = v15 + int32(32)
	return
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99 <= v84 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v101 = v84
	goto L20
L20:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+l3))))
	if v117 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v123 = int64(0)
	goto L24
L23:
	;
	v122 = *(*int64)(unsafe.Add(mBase, uint32(l2+v101<<(uint(int32(3))%32))))
	v123 = v122
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5+v101<<(uint(int32(3))%32)))) = v123
	v126 = v101 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v126 < v128 {
		v101 = v126
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
}
func F_gistFormTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	F_gistCompressValues(m, l0, l1, l2, l3, l4, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l4 != 0 {
			v17 = int32(8)
		} else {
			v17 = int32(12)
		}
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+v17)))
		v20 = F_index_form_tuple(m, v19, v9, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(_a_F_gistFormTuple_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)) = uint16(v22)
			m.G0 = v9 + int32(256)
			return v20
		}
	}
}
func F_gistMakeUnionItVec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v106 int32
	_ = v106
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v204 int64
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	v6 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v6
	v35 = F_palloc(m, l2*int32(24)+int32(56))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if int32(0) < v38 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v44 = v35 + int32(32)
	v50 = v35 + int32(8)
	v62 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v27 + int32(16)
	return
L6:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v79
	v81 = int32(1)
	if l2 <= v79 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3+v62<<(uint(int32(3))%32)))) = v204
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v62))) = uint8(v196)
	v214 = v62 + int32(1)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v214 < v216 {
		v62 = v214
		goto L6
	} else {
		goto L26
	}
L9:
	;
	v196 = v81
	v204 = int64(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v89 = l0 + int32(_a_F_gistMakeUnionItVec_0) + v62<<(uint(int32(2))%32)
	v91 = v62 * int32(28)
	v92 = l0 + int32(2708) + v91
	v106 = int32(0)
	goto L12
L12:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1+v106<<(uint(int32(2))%32))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v125 = F_index_getattr_2(m, v121, v62+int32(1), v122, v27+int32(11))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	switch v169 {
	case 0:
		v196 = v81
		v204 = int64(0)
		goto L8
	case 1:
		goto L24
	default:
		goto L23
	}
L14:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+11)))
	if v127 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v133 = v50 + v130*int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = v125
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+15)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v139 == v137 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v166 = v106 + int32(1)
	if v166 != l2 {
		v106 = v166
		goto L12
	} else {
		goto L22
	}
L18:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v159 + int32(1)
	goto L17
L19:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v144 = F_FunctionCall1Coll(m, v92, v142, base.I64_extend_i32_u(v133))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v146 = base.I32_wrap_i64(v144)
	if v133 == v146 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v152
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+16)) = uint16(v154)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+18)) = uint8(v156)
	goto L18
L22:
	;
	goto L13
L23:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v181 = F_FunctionCall2Coll(m, l0+int32(916)+v91, v180, base.I64_extend_i32_u(v35), base.I64_extend_i32_u(v27+int32(12)))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(2)
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v50)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v172
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = v174
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v44))) = v176
	goto L23
L25:
	;
	v196 = int32(0)
	v204 = v181
	goto L8
L26:
	;
	goto L7
}
func F_gistNewBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v184
L2:
	;
	return int32(0)
L3:
	;
	if v14 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v171
	v173 = int32(8)
	v175 = int32(0)
	v178 = F_ExtendBufferedRel(m, v12+v173, v175, v175, v173)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L58
	}
L7:
	;
	v29 = F_ReadBuffer(m, l0, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v31 = F_ConditionalLockBuffer(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v29 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseBuffer(m, v29)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L55
	}
L14:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)))
	if v51 == int32(0) {
		v184 = v29
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v29^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L14
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[1]))
	v50 = v44 + v29<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	F_gistcheckpage(m, l0, v29)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)))
	if v56 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_UnlockBuffer(m, v29)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L54
	}
L21:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+16)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v57)+12)))
	if v59&int32(2) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[2]))
	if v76 <= int32(0) {
		v184 = v29
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v65) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	v70 = v68
	goto L27
L26:
	;
	v70 = int64(3)
	goto L27
L27:
	;
	v71 = F_GlobalVisCheckRemovableFullXid(m, int32(0), v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v71 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+118)))
	if v80 != int32(112) {
		v184 = v29
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v83) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	v88 = v86
	goto L34
L33:
	;
	v88 = int64(3)
	goto L34
L34:
	;
	v89 = m.G0
	v91 = v89 - int32(32)
	m.G0 = v91
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_gistNewBuffer[2]))
	if v94 <= int32(1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)) = uint8(v127)
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v132
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v26
	F_XLogBeginInsert(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L51
	}
L36:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gistNewBuffer[3])))
	if v98&int32(1) == int32(0) {
		v127 = v3
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+118)))
	if v104 != int32(112) {
		v127 = v3
		goto L35
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	if int32(0) < v94 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L45
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v109 != 0 {
		v127 = v3
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v110 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v127 = v3
	goto L35
L45:
	;
	if base.Ui32(v114) < base.Ui32(int32(_a_F_gistNewBuffer_0)) {
		v127 = int32(1)
		goto L35
	} else {
		goto L46
	}
L46:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v117 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v127 = int32(0)
	goto L35
L48:
	;
	goto L49
L49:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+119)))
	switch v123 - int32(109) {
	case 0, 5:
		goto L50
	default:
		v127 = int32(0)
		goto L35
	}
L50:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+112)))
	v127 = v126
	goto L35
L51:
	;
	F_XLogRegisterData(m, v91, int32(25))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v143 = F_XLogInsert(m, int32(14), int32(32))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v91 + int32(32)
	v184 = v29
	goto L1
L54:
	;
	goto L13
L55:
	;
	v153 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	if v153 != int32(-1) {
		v26 = v153
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L8
L58:
	;
	v184 = v178
	goto L1
}
func F_gistScanPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v180 int32
	_ = v180
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int64
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v692 int32
	_ = v692
	var v724 int32
	_ = v724
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = F_ReadBuffer(m, v34, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBufferInternal(m, v36, int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v36 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v34, v59, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v36^int32(-1))*int32(56))+16))
	v59 = v50
	goto L4
L6:
	;
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[1]))
	v53 = int32(56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+v36*v53-v53)+16))
	v59 = v58
	goto L4
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_gistcheckpage(m, v63, v36)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v36 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v85 == int64(0) {
		v130 = v84
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[2]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+(v36^int32(-1))<<(uint(int32(2))%32))))
	v83 = v75
	goto L10
L12:
	;
	goto L13
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[3]))
	v83 = v77 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v83)+12)))
	if v134&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v88 = v84 + v83
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+12)))
	if v89&int32(8) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	if base.Ui64(base.I64_rotl(v94, int64(32))) <= base.Ui64(v85) {
		v130 = v84
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v98 == int32(-1) {
		v130 = v84
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v101 = int32(_a_F_gistScanPage_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = F_palloc(m, v106<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = v117 << (uint(int32(4)) % 32)
	if v119 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	base.MemoryCopy(m, v111+int32(32), l2, v119)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_pairingheap_add(m, v123, v111)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v102
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v130 = v128
	goto L14
L26:
	;
	F_UnlockReleaseBuffer(m, v36)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L115
	}
L27:
	;
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v137
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[6])))
	if v141 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_MemoryContextReset(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v144 = F_BufferGetLSNAtomic(m, v36)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v144
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+12)))
	if base.Ui32(v147) < base.Ui32(int32(25)) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v155 = int32(base.Ui32(v147+int32(_a_F_gistScanPage_1))>>(uint(int32(2))%32)) & int32(_a_F_gistScanPage_2)
	if v155 == int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v159 = v32 + int32(48)
	v164 = base.I64_extend_i32_u(v30 + int32(30))
	v165 = base.I64_extend_i32_u(v30)
	v180 = int32(1)
	goto L35
L35:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v198 = v180 & int32(_a_F_gistScanPage_2)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(20)+v198<<(uint(int32(2))%32))))
	v203 = int32(_a_F_gistScanPage_3)
	if base.B2i32(v194 == int32(1))&base.B2i32(v202&v203 == v203) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L26
L37:
	;
	v692 = v180 + int32(1)
	if base.Ui32(v692&int32(_a_F_gistScanPage_2)) <= base.Ui32(v155) {
		v180 = v692
		goto L35
	} else {
		goto L114
	}
L38:
	;
	v208 = int32(_a_F_gistScanPage_0)
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v217 = v83 + v202&int32(_a_F_gistScanPage_4)
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v218 != int32(_a_F_gistScanPage_5) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v209
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	F_MemoryContextReset(m, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L90
	}
L40:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v223 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v223 < v225 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v426)+12)))
	if v428&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L43:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v230 = v228
	v231 = v225
	v240 = v223
	goto L46
L44:
	;
	v327 = v223
	goto L45
L45:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v343 <= int32(0) {
		v531 = v327
		v535 = v223
		goto L39
	} else {
		goto L65
	}
L46:
	;
	v256 = int32(*(*int16)(unsafe.Add(mBase, uint32(v230)+4)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v260 = F_index_getattr_2(m, v217, v256, v257, v30+int32(31))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v327 = v309
	goto L45
L48:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v262&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v310 = int32(1)
	if v310 < v231 {
		v230 = v230 + int32(56)
		v231 = v231 - v310
		v240 = v309
		goto L46
	} else {
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v209
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	F_MemoryContextReset(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L63
	}
L51:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v300 == int32(0) {
		v309 = v240
		goto L49
	} else {
		goto L62
	}
L52:
	;
	if v262&int32(64) == int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v279 != 0 {
		goto L50
	} else {
		goto L58
	}
L55:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v269)+12)))
	if v271&int32(1) == int32(0) {
		v309 = v240
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v276&int32(1) != 0 {
		v309 = v240
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v230)+4)))
	F_gistdentryinit(m, v222, v280-int32(1), v30, v260, v221, v83, v198, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v286)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v230)+48))
	v292 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v230)+6)))
	v293 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v230)+8)))
	v294 = F_FunctionCall5Coll(m, v230+int32(16), v290, v165, v291, v292, v293, v164)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v294 == int64(0) {
		goto L50
	} else {
		goto L61
	}
L61:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)))
	v309 = v298 | v240
	goto L49
L62:
	;
	goto L50
L63:
	;
	goto L37
L64:
	;
	goto L47
L65:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	v349 = v346
	v350 = v343
	v353 = v347
	v363 = v223
	goto L66
L66:
	;
	v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349)+4)))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v379 = F_index_getattr_2(m, v217, v375, v376, v30+int32(31))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v531 = v327
	v535 = v416
	goto L39
L68:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v381&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v422 = int32(1)
	if v422 < v350 {
		v349 = v349 + int32(56)
		v350 = v350 - v422
		v353 = v353 + int32(16)
		v363 = v416
		goto L66
	} else {
		goto L77
	}
L70:
	;
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349)+4)))
	F_gistdentryinit(m, v222, v395-int32(1), v30, v379, v221, v83, v198, int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	if v386&int32(1) == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v391 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v353)+8)) = uint8(v391)
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = int64(0)
	v416 = v363
	goto L69
L74:
	;
	goto L73
L75:
	;
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v401)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v349)+48))
	v407 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v349)+6)))
	v408 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v349)+8)))
	v409 = F_FunctionCall5Coll(m, v349+int32(16), v405, v165, v406, v407, v408, v164)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)))
	v412 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v353)+8)) = uint8(v412)
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v409
	v416 = v411 | v363
	goto L69
L77:
	;
	goto L67
L78:
	;
	v531 = v433
	v535 = int32(0)
	goto L39
L79:
	;
	v452 = int32(0)
	goto L87
L80:
	;
	v433 = int32(0)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v433 < v434 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L78
L84:
	;
	F_errmsg_internal(m, int32(_a_F_gistScanPage_6), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_gistScanPage_7), int32(161), int32(_a_F_gistScanPage_8))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v479 = v452 << (uint(int32(4)) % 32)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v479+v480))) = int64(-4503599627370496)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v484+v479)+8)) = uint8(v486)
	v489 = v452 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v489 < v490 {
		v452 = v489
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L78
L89:
	;
	goto L88
L90:
	;
	if l3 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v571 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v555)+12)))
	if v557&int32(1) == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v562 = int32(1)
	F_tbm_add_tuples(m, l3, v217, v562, v531&v562)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v567 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v567 + int64(1)
	goto L37
L95:
	;
	v616 = int32(_a_F_gistScanPage_0)
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v619
	v622 = v571 << (uint(int32(4)) % 32)
	v625 = F_palloc(m, v622+int32(32))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L102
	}
L96:
	;
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v572)+12)))
	if v574&int32(1) == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	v580 = int32(4)
	v582 = v159 + v579<<(uint(v580)%32)
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v582)+4)) = uint16(v583)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = v585
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	v590 = v159 + v587<<(uint(v580)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v590)+6)) = uint8(v531)
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+12)) = uint16(v180)
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v593 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v596 = int32(_a_F_gistScanPage_0)
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4]))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[6])))
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v599
	v601 = F_gistFetchTuple(m, v33, v34, v217)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v610 = v587
	goto L100
L100:
	;
	v614 = v610 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))) = uint16(v614)
	goto L37
L101:
	;
	v603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_gistScanPage[5]))))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v603<<(uint(int32(4))%32))+8)) = v601
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v597
	v610 = v603
	goto L100
L102:
	;
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v627)+12)))
	if v629&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v622 != 0 {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v625)+12)) = int32(-1)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v625)+16)) = v634
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v625)+20)) = uint16(v636)
	*(*uint8)(unsafe.Add(mBase, uint32(v625)+23)) = uint8(v535)
	*(*uint8)(unsafe.Add(mBase, uint32(v625)+22)) = uint8(v531)
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v640 != int32(1) {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	*(*int32)(unsafe.Add(mBase, uint32(v625)+12)) = v646 | v647<<(uint(int32(16))%32)
	v652 = F_BufferGetLSNAtomic(m, v36)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v643 = F_gistFetchTuple(m, v33, v34, v217)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v625)+24)) = v643
	goto L103
L109:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v625)+16)) = v652
	goto L103
L110:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	base.MemoryCopy(m, v625+int32(32), v657, v622)
	goto L112
L111:
	;
	goto L112
L112:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_pairingheap_add(m, v659, v625)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistScanPage[4])) = v617
	goto L37
L114:
	;
	goto L36
L115:
	;
	m.G0 = v30 + int32(32)
	return
}
func F_gistSortedBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	v13 = int32(_a_F_gistSortedBuildCallback_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0])) = v17
	F_gistCompressValues(m, v16, l0, l2, l3, int32(1), v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
		F_tuplesort_putindextuplevalues(m, v22, v23, l1, v11, l3)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_gistSortedBuildCallback[0])) = v14
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			F_MemoryContextReset(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
				*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v32 + int64(1)
				m.G0 = v11 + int32(256)
				return
			}
		}
	}
}
func F_gist_poly_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		v17 = int32(0)
		if base.B2i32(v16 == v17)|base.B2i32(v8 == v17) != 0 {
			v35 = int64(0)
			return v35
		} else {
			v28 = F_rtree_internal_consistent(m, v16, v8+int32(8), base.I32_wrap_i64(v12)&int32(_a_F_gist_poly_consistent_0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v30 != v8 {
					F_pfree(m, v8)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v35 = base.I64_extend_i32_u(v28)
						return v35
					}
				} else {
					v35 = base.I64_extend_i32_u(v28)
					return v35
				}
			}
		}
	}
}
func F_gist_translate_cmptype_common(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = v2 - int32(1)
	if base.Ui32(v4) <= base.Ui32(int32(7)) {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v4<<(uint(int32(3))%32))+uint32(_c_F_gist_translate_cmptype_common[0])))
		v11 = v9
	} else {
		v11 = int64(0)
	}
	return v11
}
func F_gist_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_gist_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
