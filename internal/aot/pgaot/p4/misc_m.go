package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MarkAsPreparingGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int64
	_ = v40
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	v7 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_MarkAsPreparingGuts[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = v11 + v12*int32(640)
	base.MemoryFill(m, v15+int32(8), v7, int32(632))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_MarkAsPreparingGuts[1]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v33
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v15)+92)) = v40
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+74)) = uint16(v34)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+72)) = uint8(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v34
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+277)) = uint8(v34)
	v57 = v15 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v57
	v61 = v15 + int32(260)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v61
	v65 = v15 + int32(252)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v65
	v69 = v15 + int32(244)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v69
	v73 = v15 + int32(236)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v15)+236)) = v73
	v77 = v15 + int32(228)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+232)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+228)) = v77
	v81 = v15 + int32(220)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v81
	v85 = v15 + int32(212)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+216)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v85
	v89 = v15 + int32(204)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v89
	v93 = v15 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v15)+196)) = v93
	v97 = v15 + int32(188)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v15)+188)) = v97
	v101 = v15 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+184)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v101
	v105 = v15 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v105
	v109 = v15 + int32(164)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v109
	v113 = v15 + int32(156)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v113
	v117 = v15 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v117
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+276)) = uint8(v34)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_MarkAsPreparingGuts[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)) = uint8(v34)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v34)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v126
	v133 = l0 + int32(47)
	if (l2^v133)&int32(3) != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v27
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_MarkAsPreparingGuts[2]))
	v33 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l1
	v33 = int32(-1)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkAsPreparingGuts[3])) = l0
	return
L6:
	;
	goto L5
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v187)
	if v187&int32(255) == int32(0) {
		goto L6
	} else {
		goto L22
	}
L8:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v186 = l2
	v187 = v139
	v188 = v133
	goto L7
L9:
	;
	goto L10
L10:
	;
	if l2&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v143 = l2
	v145 = v133
	goto L14
L12:
	;
	v157 = l2
	v159 = v133
	goto L13
L13:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v164 = int32(-2139062144)
	if (int32(16843008)-v161|v161)&v164 != v164 {
		v186 = v157
		v187 = v161
		v188 = v159
		goto L7
	} else {
		goto L18
	}
L14:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v146)
	if v146 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v157 = v153
	v159 = v151
	goto L13
L16:
	;
	v150 = int32(1)
	v151 = v145 + v150
	v153 = v143 + v150
	if v153&int32(3) != 0 {
		v143 = v153
		v145 = v151
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v169 = v157
	v170 = v161
	v171 = v159
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v170
	v173 = int32(4)
	v174 = v171 + v173
	v176 = v169 + v173
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v181 = int32(-2139062144)
	if (int32(16843008)-v178|v178)&v181 == v181 {
		v169 = v176
		v170 = v178
		v171 = v174
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v186 = v176
	v187 = v178
	v188 = v174
	goto L7
L21:
	;
	goto L20
L22:
	;
	v195 = v186
	v197 = v188
	goto L23
L23:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)) = uint8(v198)
	v200 = int32(1)
	if v198 != 0 {
		v195 = v195 + v200
		v197 = v197 + v200
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L6
L25:
	;
	goto L24
}
func F___math_invalid(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	v2 = base.F64_sub(l0, l0)
	return base.F64_div(v2, v2)
}
func F___math_uflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(1.2882297539194267e-231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F__mdfd_openseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v15 = v12 + int32(93)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v15, v16, v17, v18, v19, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v110 = v12 + int32(93)
	base.MemoryCopy(m, v110, v12+int32(10), int32(83))
	v118 = *(*int32)(unsafe.Add(mBase, _c_F__mdfd_openseg[0]))
	if v118&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	v29 = F_pg_sprintf(m, v12+int32(10), int32(_a_F__mdfd_openseg_0), v12)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v32 = v12 + int32(10)
	v34 = v12 + int32(93)
	if (v34^v32)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L3
L8:
	;
	goto L3
L9:
	;
	goto L8
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v88)
	if v88&int32(255) == int32(0) {
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v87 = v34
	v88 = v40
	v89 = v32
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v34&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = v34
	v46 = v32
	goto L17
L15:
	;
	v58 = v34
	v60 = v32
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v65 = int32(-2139062144)
	if (int32(16843008)-v62|v62)&v65 != v65 {
		v87 = v58
		v88 = v62
		v89 = v60
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v47)
	if v47 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v58 = v54
	v60 = v52
	goto L16
L19:
	;
	v51 = int32(1)
	v52 = v46 + v51
	v54 = v44 + v51
	if v54&int32(3) != 0 {
		v44 = v54
		v46 = v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v70 = v58
	v71 = v62
	v72 = v60
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v71
	v74 = int32(4)
	v75 = v72 + v74
	v77 = v70 + v74
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v82 = int32(-2139062144)
	if (int32(16843008)-v79|v79)&v82 == v82 {
		v70 = v77
		v71 = v79
		v72 = v75
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v87 = v77
	v88 = v79
	v89 = v75
	goto L10
L24:
	;
	goto L23
L25:
	;
	v96 = v87
	v98 = v89
	goto L26
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)) = uint8(v99)
	v101 = int32(1)
	if v99 != 0 {
		v96 = v96 + v101
		v98 = v98 + v101
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	goto L27
L29:
	;
	v121 = int32(_a_F__mdfd_openseg_1)
	goto L31
L30:
	;
	v121 = int32(2)
	goto L31
L31:
	;
	v123 = F_PathNameOpenFile(m, v110, v121|l3)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if int32(0) <= v123 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v129 = l0 + l1<<(uint(int32(2))%32)
	v131 = v129 + int32(40)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v134 = l2 + int32(1)
	if v134 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v182 = int32(0)
	goto L35
L35:
	;
	m.G0 = v12 + int32(176)
	return v182
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v134
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56))
	v178 = v175 + l2<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v123
	v182 = v178
	goto L35
L37:
	;
	if v132 <= int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v132 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v140 = v129 + int32(56)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	F_pfree(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = int32(0)
	goto L36
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F__mdfd_openseg[1]))
	v155 = F_MemoryContextAlloc(m, v152, v134<<(uint(int32(3))%32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v134 <= v132 {
		goto L36
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56)) = v155
	goto L36
L46:
	;
	v163 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v167 = F_repalloc(m, v164, v134<<(uint(int32(3))%32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v167
	goto L36
}
func F_makeBoolean(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v1 = l0
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v1)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(467)
		return v4
	}
}
func F_makeCompoundFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v10 == v3 {
		v57 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(1056)
	return v57
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 == int32(0) {
		v57 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(0)
	goto L4
L4:
	;
	v31 = v8 + int32(16)
	F_getNextFlagFromString(m, l0, v8+int32(12), v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v57 = v49 & int32(15)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v37 = v8 + int32(1044)
	F_setCompoundAffixFlagValue(m, l0, v37, v31, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v45 = F_bsearch(m, v37, v41, v42, int32(12), int32(1151))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v49 = v47 | v24
	goto L12
L11:
	;
	v49 = v24
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v51 != 0 {
		v24 = v49
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
}
func F_makeWholeRowVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(2249)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v15 {
	case 0:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = F_get_rel_type_id(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v101 = v17
				v102 = v5
				v103 = v5
				v106 = F_palloc0(m, int32(48))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					v108 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
					*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
					v112 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
					m.G0 = v12 + int32(32)
					return v106
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v29 = F_get_rel_name(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v29
							F_errmsg(m, int32(_a_F_makeWholeRowVar_0), v12)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_makeWholeRowVar_1), int32(155), int32(_a_F_makeWholeRowVar_2))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	case 1:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v40 != 0 {
			v41 = F_get_rel_type_id(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 != 0 {
					v101 = v41
					v102 = v5
					v103 = v5
					v106 = F_palloc0(m, int32(48))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
						*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
						v112 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
						*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
						*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
						*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
						m.G0 = v12 + int32(32)
						return v106
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v51 = F_get_rel_name(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v51
								F_errmsg(m, int32(_a_F_makeWholeRowVar_0), v12+int32(16))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_makeWholeRowVar_1), int32(181), int32(_a_F_makeWholeRowVar_2))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v64 == int32(0) {
				v101 = v14
				v102 = v5
				v103 = v5
				v106 = F_palloc0(m, int32(48))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					v108 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
					*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
					v112 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
					m.G0 = v12 + int32(32)
					return v106
				}
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
				v70 = F_exprType(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v73 = F_type_is_rowtype(m, v70)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						if v73 != 0 {
							v75 = v70
						} else {
							v75 = int32(2249)
						}
						v101 = v75
						v102 = v5
						v103 = v5
						v106 = F_palloc0(m, int32(48))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v108 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
							*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
							v112 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
							*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
							*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
							*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
							*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
							m.G0 = v12 + int32(32)
							return v106
						}
					}
				}
			}
		}
	default:
		v101 = v14
		v102 = v5
		v103 = v5
		v106 = F_palloc0(m, int32(48))
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return int32(0)
		} else {
			v108 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
			*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
			v112 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
			*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
			*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
			*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
			*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
			*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
			m.G0 = v12 + int32(32)
			return v106
		}
	case 3:
		v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
		if v76 != 0 {
			v101 = v14
			v102 = v5
			v103 = v5
			v106 = F_palloc0(m, int32(48))
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int32(0)
			} else {
				v108 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
				*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
				v112 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
				*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
				*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
				*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
				*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
				*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
				m.G0 = v12 + int32(32)
				return v106
			}
		} else {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v77 == int32(0) {
				v101 = v14
				v102 = v5
				v103 = v5
				v106 = F_palloc0(m, int32(48))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					v108 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
					*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
					v112 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
					*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
					m.G0 = v12 + int32(32)
					return v106
				}
			} else {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				if v80 != int32(1) {
					v101 = v14
					v102 = v5
					v103 = v5
					v106 = F_palloc0(m, int32(48))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
						*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
						v112 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
						*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
						*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
						*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
						*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
						m.G0 = v12 + int32(32)
						return v106
					}
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
					v86 = F_exprType(m, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v89 = F_type_is_rowtype(m, v86)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							if v89 != 0 {
								v91 = v86
							} else {
								v91 = int32(2249)
							}
							if v89 != 0 {
								v101 = v91
								v102 = int32(0)
								v103 = v5
								v106 = F_palloc0(m, int32(48))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v108 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
									*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
									v112 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
									*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
									*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
									*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
									*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
									*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
									*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
									m.G0 = v12 + int32(32)
									return v106
								}
							} else {
								v93 = int32(0)
								if l3 == v93 {
									v101 = v91
									v102 = v93
									v103 = v5
									v106 = F_palloc0(m, int32(48))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
										*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
										v112 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
										*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
										*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
										*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
										*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
										*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
										*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
										m.G0 = v12 + int32(32)
										return v106
									}
								} else {
									v97 = F_exprCollation(m, v85)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										v101 = v86
										v102 = v97
										v103 = int32(1)
										v106 = F_palloc0(m, int32(48))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v108 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v108
											*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v102
											v112 = int32(-1)
											*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v112
											*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v101
											*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)) = uint16(v103)
											*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(6)
											*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v112
											*(*uint16)(unsafe.Add(mBase, uint32(v106)+40)) = uint16(v103)
											*(*int32)(unsafe.Add(mBase, uint32(v106)+36)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v108
											m.G0 = v12 + int32(32)
											return v106
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_make_absolute_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L39
	}
L2:
	;
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L35
	}
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != int32(47) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v102 = int32(0)
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v102
L6:
	;
	F_canonicalize_path_enc(m, v96)
	mBase = m.M
	v102 = v96
	goto L5
L7:
	;
	v13 = int32(1024)
	v15 = F_emscripten_builtin_malloc(m, v13)
	mBase = m.M
	if v15 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	goto L9
L9:
	;
	v84 = F_strlen(m, l0)
	mBase = m.M
	v86 = v84 + int32(1)
	v87 = F_emscripten_builtin_malloc(m, v86)
	mBase = m.M
	if v87 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	v68 = F_strlen(m, v18)
	mBase = m.M
	v69 = F_strlen(m, l0)
	mBase = m.M
	v73 = F_emscripten_builtin_malloc(m, v68+v69+int32(2))
	mBase = m.M
	if v73 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_absolute_path[0])) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L25
	}
L12:
	;
	v17 = v13
	v18 = v15
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v21 = F_getcwd(m, v18, v17)
	mBase = m.M
	if v21 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_make_absolute_path[0]))
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	if v23 != int32(68) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v28 = v17 << (uint(int32(1)) % 32)
	v29 = F_emscripten_builtin_malloc(m, v28)
	mBase = m.M
	if v29 != 0 {
		v17 = v28
		v18 = v29
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	return int32(0)
L21:
	;
	F_errcode(m, int32(_a_F_make_absolute_path_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_make_absolute_path_1), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_make_absolute_path_2), int32(829), int32(_a_F_make_absolute_path_3))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_make_absolute_path_4), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_make_absolute_path_2), int32(851), int32(_a_F_make_absolute_path_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	v79 = F_pg_sprintf(m, v73, int32(_a_F_make_absolute_path_5), v8)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	v96 = v73
	goto L6
L30:
	;
	if v92 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v92 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v91 = F___memcpy(m, v87, l0, v86)
	mBase = m.M
	v92 = v91
	goto L30
L34:
	;
	v96 = v92
	goto L6
L35:
	;
	F_errcode(m, int32(_a_F_make_absolute_path_0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_make_absolute_path_1), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_make_absolute_path_2), int32(866), int32(_a_F_make_absolute_path_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(_a_F_make_absolute_path_0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_make_absolute_path_1), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_make_absolute_path_2), int32(883), int32(_a_F_make_absolute_path_3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_attrmap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		v11 = F_palloc0(m, l0<<(uint(int32(1))%32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v11
			return v4
		}
	}
}
func F_make_canonical_pathkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v5 = l4
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = l1
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L22
	} else {
		goto L25
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v20 != 0 {
		v12 = v20
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	return v93
L8:
	;
	v66 = int32(_a_F_make_canonical_pathkey_0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_make_canonical_pathkey[0]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_make_canonical_pathkey[0])) = v69
	v72 = F_palloc0(m, int32(20))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v27 = int32(0)
	if v27 < v24 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v24
	goto L13
L12:
	;
	v31 = v27
	goto L13
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v34 = v27
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32+v34<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v12 != v46 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v55 = v34 + int32(1)
	if v55 != v31 {
		v34 = v55
		goto L14
	} else {
		goto L21
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if l2 != v48 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if l3 != v50 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)))
	if v52 == v5 {
		v93 = v45
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L15
L22:
	;
	return int32(0)
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(275)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = F_lappend(m, v82, v72)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_make_canonical_pathkey[0])) = v67
	v93 = v72
	goto L7
L25:
	;
	F_errmsg_internal(m, int32(_a_F_make_canonical_pathkey_1), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_make_canonical_pathkey_2), int32(66), int32(_a_F_make_canonical_pathkey_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_empty_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)) = uint8(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v5)+3)) = v2
	v21 = F_make_range(m, l0, v5+int32(8), v5, v7, v2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v21
	}
}
func F_make_oper_cache_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v233 int32
	_ = v233
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_DeconstructQualifiedName(m, l2, v11+int32(28), v11+int32(24))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(0)
	base.MemoryFill(m, l1, v21, int32(136))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	goto L6
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l3
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v147 != 0 {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	v142 = F_strlen(m, v131)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v32 = int32(63)
	if (l1^v25)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
	goto L4
L9:
	;
	v116 = v111
	v117 = v112
	v118 = v113
	goto L30
L10:
	;
	if v106 == int32(0) {
		v131 = v104
		v132 = v105
		goto L8
	} else {
		goto L29
	}
L11:
	;
	v104 = v25
	v105 = l1
	v106 = v32
	goto L10
L12:
	;
	goto L13
L13:
	;
	v36 = int32(0)
	if base.B2i32(v25&int32(3) == v36)|int32(0) == v36 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v72 == int32(0) {
		v131 = v69
		v132 = v70
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v48 = v25
	v49 = l1
	v50 = v32
	goto L18
L16:
	;
	goto L17
L17:
	;
	v69 = v25
	v70 = l1
	v71 = v32
	v72 = int32(1)
	goto L14
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v52)
	if v52 == int32(0) {
		v111 = v48
		v112 = v49
		v113 = v50
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v69 = v63
	v70 = v57
	v71 = v59
	v72 = v61
	goto L14
L20:
	;
	v56 = int32(1)
	v57 = v49 + v56
	v59 = v50 - v56
	v60 = int32(0)
	v61 = base.B2i32(v59 != v60)
	v63 = v48 + v56
	if v63&int32(3) == v60 {
		v69 = v63
		v70 = v57
		v71 = v59
		v72 = v61
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v59 != 0 {
		v48 = v63
		v49 = v57
		v50 = v59
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if base.B2i32(v75 == int32(0))|base.B2i32(base.Ui32(v71) < base.Ui32(int32(4))) != 0 {
		v104 = v69
		v105 = v70
		v106 = v71
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v82 = v69
	v83 = v70
	v84 = v71
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 != v90 {
		v111 = v82
		v112 = v83
		v113 = v84
		goto L9
	} else {
		goto L27
	}
L26:
	;
	v104 = v98
	v105 = v96
	v106 = v100
	goto L10
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v87
	v95 = int32(4)
	v96 = v83 + v95
	v98 = v82 + v95
	v100 = v84 - v95
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v82 = v98
		v83 = v96
		v84 = v100
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v111 = v104
	v112 = v105
	v113 = v106
	goto L9
L30:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	if v120 == int32(0) {
		v131 = v116
		v132 = v117
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v131 = v127
	v132 = v125
	goto L8
L32:
	;
	v124 = int32(1)
	v125 = v117 + v124
	v127 = v116 + v124
	v129 = v118 - v124
	if v129 != 0 {
		v116 = v127
		v117 = v125
		v118 = v129
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	m.G0 = v11 + int32(32)
	return v233
L35:
	;
	v233 = int32(1)
	goto L34
L36:
	;
	v149 = v11 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = v149
	v155 = int32(_a_F_make_oper_cache_key_0)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_make_oper_cache_key[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = v156
	*(*int32)(unsafe.Add(mBase, _c_F_make_oper_cache_key[0])) = v11 + int32(12)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v172 = int32(0)
	F_recomputeNamespacePath(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L42
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v164 = F_LookupExplicitNamespace(m, v162, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v164
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_make_oper_cache_key[0])) = v168
	goto L41
L41:
	;
	goto L35
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_make_oper_cache_key[1]))
	if v177 == int32(0) {
		v212 = v172
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if int32(16) < v212 {
		v233 = v21
		goto L34
	} else {
		goto L55
	}
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v180 <= int32(0) {
		v212 = v172
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_make_oper_cache_key[2]))
	v185 = v172
	v188 = v172
	goto L46
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v188<<(uint(int32(2))%32))))
	if v184 != v197 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v212 = v207
	goto L43
L48:
	;
	if v185 < int32(16) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v207 = v185
	goto L50
L50:
	;
	v209 = v188 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v209 < v210 {
		v185 = v207
		v188 = v209
		goto L46
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(72)+v185<<(uint(int32(2))%32)))) = v197
	goto L53
L52:
	;
	goto L53
L53:
	;
	v207 = v185 + int32(1)
	goto L50
L54:
	;
	goto L47
L55:
	;
	goto L35
}
func F_make_placeholder_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(319)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
		v19 = v17 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v19
		return v6
	}
}
func F_make_trigrams(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var __phi161 int32
	_ = __phi161
	var v162 int32
	_ = v162
	var __phi162 int32
	_ = __phi162
	var v163 int32
	_ = v163
	var __phi163 int32
	_ = __phi163
	var v164 int32
	_ = v164
	var __phi164 int32
	_ = __phi164
	var v165 int32
	_ = v165
	var __phi165 int32
	_ = __phi165
	var v166 int32
	_ = v166
	var __phi166 int32
	_ = __phi166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	if int32(3) <= l2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L52
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = l2 + v18 - int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v21) <= base.Ui32(v22) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	v43 = v38 + v37*int32(3) + int32(5)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_make_trigrams[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	goto L13
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = v18
	v38 = v24
	goto L5
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(int32(357913942)) <= base.Ui32(v21) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = F_repalloc(m, v27, v21*int32(3)+int32(5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v36
	v38 = v32
	goto L5
L12:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v293 = base.I32_div_s(v277-v288-int32(5), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	goto L4
L13:
	;
	if base.Ui32(v46) <= base.Ui32(int32(41)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v53 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46*int32(28))+uint32(_c_F_make_trigrams[1])))
	v53 = v51
	goto L17
L16:
	;
	v53 = int32(1)
	goto L17
L17:
	;
	goto L14
L18:
	;
	v60 = l1
	v63 = v43
	goto L21
L19:
	;
	goto L20
L20:
	;
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v85 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v74)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)) = uint8(v76)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)) = uint8(v78)
	v81 = v63 + int32(3)
	v83 = v60 + int32(1)
	if base.Ui32(v83) < base.Ui32(l1+l2-int32(2)) {
		v60 = v83
		v63 = v81
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v277 = v81
	goto L12
L23:
	;
	goto L22
L24:
	;
	v151 = F_pg_mblen_unbounded(m, v141)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L38
	}
L25:
	;
	v127 = F_pg_mblen_unbounded(m, l1)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L34
	}
L26:
	;
	v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v88 < int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v98 = l1
	v99 = v43
	goto L28
L28:
	;
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98)+2)))
	if v110 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v277 = v123
	goto L12
L30:
	;
	v139 = v98
	v140 = v99
	v141 = v98 + int32(2)
	v142 = int32(1)
	v143 = int32(1)
	goto L24
L31:
	;
	goto L32
L32:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v116)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)) = uint8(v118)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+2)) = uint8(v120)
	v123 = v99 + int32(3)
	v125 = v98 + int32(1)
	if v125 != l1+l2-int32(2) {
		v98 = v125
		v99 = v123
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	if l2 <= v127 {
		v277 = v43
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v130 = l1 + v127
	v131 = F_pg_mblen_unbounded(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	if l2 <= v131+v127 {
		v277 = v43
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v139 = l1
	v140 = v43
	v141 = v130 + v131
	v142 = v131
	v143 = v127
	goto L24
L38:
	;
	v155 = v151 + (v139 + v143 + v142)
	v156 = l1 + l2
	if base.Ui32(v156) < base.Ui32(v155) {
		v277 = v140
		goto L12
	} else {
		goto L39
	}
L39:
	;
	__phi161 = v139
	__phi162 = v140
	__phi163 = v151
	__phi164 = v142
	__phi165 = v143
	__phi166 = v155
	v161 = __phi161
	v162 = __phi162
	v163 = __phi163
	v164 = __phi164
	v165 = __phi165
	v166 = __phi166
	goto L40
L40:
	;
	v173 = int32(255)
	v177 = v166 - v161
	switch v177 {
	case 0:
		v235 = v177
		goto L43
	default:
		goto L44
	case 3:
		goto L45
	}
L41:
	;
	v277 = v266
	goto L12
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)) = uint8(v263)
	v266 = v162 + int32(3)
	if v166 == v156 {
		v277 = v266
		goto L12
	} else {
		goto L49
	}
L43:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v162))) = uint16(v235)
	v263 = int32(base.Ui32(v235) >> (uint(int32(16)) % 32))
	goto L42
L44:
	;
	v184 = v161
	v188 = v177
	v194 = v173
	v195 = v173
	v196 = v173
	v197 = v173
	goto L46
L45:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v178)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)) = uint8(v180)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
	v263 = v182
	goto L42
L46:
	;
	v198 = int32(8)
	v202 = int32(16)
	v207 = int32(24)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v216 = *(*int32)(unsafe.Add(mBase, uint32((v194^v210)<<(uint(int32(2))%32))+uint32(_c_F_make_trigrams[2])))
	v217 = v195<<(uint(v198)%32)&int32(_a_F_make_trigrams_0) | v196<<(uint(v202)%32)&int32(16711680) | v197<<(uint(v207)%32) ^ v216
	v224 = int32(1)
	v227 = v188 - v224
	if v227 != 0 {
		v184 = v184 + v224
		v188 = v227
		v194 = int32(base.Ui32(v217) >> (uint(v207) % 32))
		v195 = v216
		v196 = int32(base.Ui32(v217) >> (uint(v198) % 32))
		v197 = int32(base.Ui32(v217) >> (uint(v202) % 32))
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v235 = v217 ^ int32(-1)
	goto L43
L48:
	;
	goto L47
L49:
	;
	v269 = F_pg_mblen_unbounded(m, v166)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v271 = v166 + v269
	if base.Ui32(v271) <= base.Ui32(v156) {
		__phi161 = v161 + v165
		__phi162 = v266
		__phi163 = v269
		__phi164 = v163
		__phi165 = v164
		__phi166 = v271
		v161 = __phi161
		v162 = __phi162
		v163 = __phi163
		v164 = __phi164
		v165 = __phi165
		v166 = __phi166
		goto L40
	} else {
		goto L51
	}
L51:
	;
	goto L41
L52:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_make_trigrams_1), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_make_trigrams_2), int32(134), int32(_a_F_make_trigrams_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makesearch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v10 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)))
	v16 = v10
	goto L4
L3:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v184 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	if v13 != v22 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_rainbow(m, l1, v27, int32(-1), v9, v9)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+58)))
	if v22 != v24 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v26 != 0 {
		v16 = v26
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	goto L3
L12:
	;
	return
L13:
	;
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+56)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_makesearch[0]))
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v36 <= v37 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L16
L18:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+58)))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_makesearch[0]))
	if v102 != 0 {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	F_createarc(m, l1, int32(112), v31, v9, v9)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L12
	} else {
		goto L39
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v39 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v60 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L23:
	;
	v46 = v39
	goto L24
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v52 != v9 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v59 != 0 {
		v46 = v59
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v54 != v31&int32(_a_F_makesearch_0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v56 == int32(112) {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L25
L31:
	;
	v67 = v60
	goto L32
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v73 != v9 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L19
L34:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	if v80 != 0 {
		v67 = v80
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v75 != v31&int32(_a_F_makesearch_0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v77 == int32(112) {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	goto L18
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v105 <= v106 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L42
L44:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if v169&int32(2) == int32(0) {
		goto L3
	} else {
		goto L66
	}
L45:
	;
	F_createarc(m, l1, int32(112), v100, v9, v9)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L65
	}
L46:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v108 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v129 == int32(0) {
		goto L45
	} else {
		goto L57
	}
L49:
	;
	v115 = v108
	goto L50
L50:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v121 != v9 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L45
L52:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	if v128 != 0 {
		v115 = v128
		goto L50
	} else {
		goto L56
	}
L53:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	if v123 != v100&int32(_a_F_makesearch_0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v125 == int32(112) {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	goto L51
L57:
	;
	v136 = v129
	goto L58
L58:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v142 != v9 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L45
L60:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	if v149 != 0 {
		v136 = v149
		goto L58
	} else {
		goto L64
	}
L61:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+4)))
	if v144 != v100&int32(_a_F_makesearch_0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v146 == int32(112) {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L59
L65:
	;
	goto L44
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = int32(256)
	goto L3
L67:
	;
	v191 = v184
	v193 = v3
	goto L68
L68:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	if v196 == int32(0) {
		v219 = v193
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v219 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v191)+16))
	if v221 != 0 {
		v191 = v221
		v193 = v219
		goto L68
	} else {
		goto L82
	}
L71:
	;
	v201 = v196
	goto L72
L72:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v9 == v207 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	if v210 != 0 {
		v219 = v193
		goto L70
	} else {
		goto L78
	}
L74:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	if v209 != 0 {
		v201 = v209
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v219 = v193
	goto L70
L78:
	;
	if v193 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v211 = v193
	goto L81
L80:
	;
	v211 = v195
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+24)) = v211
	v219 = v195
	goto L70
L82:
	;
	goto L69
L83:
	;
	v230 = v219
	goto L84
L84:
	;
	v232 = F_newstate(m, l1)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L86
	}
L85:
	;
	goto L1
L86:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v234 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	if v235 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v262 != 0 {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v236 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v241 = v236
	goto L91
L91:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+4)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	F_createarc(m, l1, v247, v248, v232, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L93
	}
L92:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v253 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	if v252 != 0 {
		v241 = v252
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	goto L88
L96:
	;
	v265 = v262
	goto L99
L97:
	;
	goto L98
L98:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v230)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(0)
	if v359 == v230 {
		goto L1
	} else {
		goto L132
	}
L99:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	if v9 != v272 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	F_cparc(m, l1, v265, v272, v232)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v271 != 0 {
		v265 = v271
		goto L99
	} else {
		goto L131
	}
L104:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v265)+4)))
	if v282 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L103
L106:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	if v317 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v287 = v285 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v287))|base.B2i32(int32(1)<<(uint(v287)%32)&int32(_a_F_makesearch_1) == int32(0)) != 0 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v297 != 0 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v265)+36))
	if v298 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v310 != 0 {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v265)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v302+v282*int32(24))+12)) = v306
	v310 = v306
	goto L110
L112:
	;
	goto L113
L113:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v265)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+32)) = v308
	v310 = v308
	goto L110
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+36)) = v298
	goto L116
L115:
	;
	goto L116
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v265)+32)) = int64(0)
	goto L106
L117:
	;
	if v316 != 0 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+20)) = v316
	goto L117
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+16)) = v316
	goto L117
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+20)) = v317
	goto L123
L122:
	;
	goto L123
L123:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = v323 - int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v265)+28))
	if v328 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v327 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v327
	goto L124
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+24)) = v327
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+28)) = v328
	goto L130
L129:
	;
	goto L130
L130:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v334 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = int32(0)
	v341 = v265 + int32(8)
	v342 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v341)+16)) = v342
	*(*int64)(unsafe.Add(mBase, uint32(v341)+8)) = v342
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v342
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+16)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v265
	goto L105
L131:
	;
	goto L100
L132:
	;
	if v359 != 0 {
		v230 = v359
		goto L84
	} else {
		goto L133
	}
L133:
	;
	goto L85
}
func F_markreachable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module) int32)(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(101)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v16 = v14
	goto L8
L7:
	;
	v16 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = v20
	goto L12
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	F_markreachable(m, l0, v26, l2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v29 != 0 {
		v24 = v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_maybe_reread_subscription(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[0])))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v396 == int32(2) {
		goto L120
	} else {
		goto L121
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L10
	} else {
		goto L117
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v19 = base.B2i32(v17 == int32(2))
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v9 + int32(80)
	return
L6:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v24 = int32(_a_F_maybe_reread_subscription_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[2]))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[2])) = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v34 = F_GetSubscription(m, v32, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+25)))
	if v65 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
	F_errmsg(m, int32(_a_F_maybe_reread_subscription_1), v9)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(2) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_errfinish(m, int32(_a_F_maybe_reread_subscription_2), int32(3991), int32(_a_F_maybe_reread_subscription_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+36))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if base.B2i32(v101 == int32(0))|base.B2i32(v101 != v104) != 0 {
		v122 = v101
		v123 = v104
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v68 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v68 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v72
	F_errmsg(m, int32(_a_F_maybe_reread_subscription_4), v9-int32(-64))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+16)))
	if base.B2i32(v87 == int32(1))&base.B2i32(v86 == int32(3)) != 0 {
		goto L27
	} else {
		goto L35
	}
L33:
	;
	F_errfinish(m, int32(_a_F_maybe_reread_subscription_2), int32(4005), int32(_a_F_maybe_reread_subscription_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v395 = v85
	v396 = v86
	goto L1
L36:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	if v278 != 0 {
		v333 = v277
		goto L88
	} else {
		goto L89
	}
L37:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+16)))
	if v227 != int32(1) {
		goto L78
	} else {
		goto L79
	}
L38:
	;
	if v122-v123 != 0 {
		goto L37
	} else {
		goto L45
	}
L39:
	;
	goto L38
L40:
	;
	v107 = v95
	v108 = v98
	goto L41
L41:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v112
		v123 = v111
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v122 = v112
	v123 = v111
	goto L39
L43:
	;
	v115 = int32(1)
	if v112 == v111 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v150-v151 != 0 {
		goto L37
	} else {
		goto L53
	}
L47:
	;
	goto L46
L48:
	;
	v135 = v125
	v136 = v126
	goto L49
L49:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L47
	} else {
		goto L51
	}
L50:
	;
	v150 = v140
	v151 = v139
	goto L47
L51:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v97)+40))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if base.B2i32(v157 == int32(0))|base.B2i32(v157 != v160) != 0 {
		v178 = v157
		v179 = v160
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v178-v179 != 0 {
		goto L37
	} else {
		goto L61
	}
L55:
	;
	goto L54
L56:
	;
	v163 = v153
	v164 = v154
	goto L57
L57:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v168 == int32(0) {
		v178 = v168
		v179 = v167
		goto L55
	} else {
		goto L59
	}
L58:
	;
	v178 = v168
	v179 = v167
	goto L55
L59:
	;
	v171 = int32(1)
	if v168 == v167 {
		v163 = v163 + v171
		v164 = v164 + v171
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+26)))
	if v181 != v182 {
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+27)))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+27)))
	if v184 != v185 {
		goto L37
	} else {
		goto L63
	}
L63:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+30)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+30)))
	if v187 != v188 {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if base.B2i32(v194 == int32(0))|base.B2i32(v194 != v197) != 0 {
		v215 = v194
		v216 = v197
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v215-v216 != 0 {
		goto L37
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	v200 = v190
	v201 = v191
	goto L68
L68:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	if v205 == int32(0) {
		v215 = v205
		v216 = v204
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v215 = v205
	v216 = v204
	goto L66
L70:
	;
	v208 = int32(1)
	if v205 == v204 {
		v200 = v200 + v208
		v201 = v201 + v208
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	if v218 != v219 {
		goto L37
	} else {
		goto L73
	}
L73:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v97)+48))
	v223 = F_equal(m, v221, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	if v223 != 0 {
		goto L36
	} else {
		goto L75
	}
L75:
	;
	goto L37
L76:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+16)))
	if base.B2i32(v268 == int32(1))&base.B2i32(v267 == int32(3)) != 0 {
		goto L36
	} else {
		goto L87
	}
L77:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v253
	F_errmsg(m, v249, v9+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L85
	}
L78:
	;
	v243 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L83
	}
L79:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v230 != int32(3) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v235 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	if v235 == int32(0) {
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v249 = int32(_a_F_maybe_reread_subscription_5)
	v250 = int32(4036)
	goto L77
L83:
	;
	if v243 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v249 = int32(_a_F_maybe_reread_subscription_6)
	v250 = int32(4040)
	goto L77
L85:
	;
	F_errfinish(m, int32(_a_F_maybe_reread_subscription_2), v250, int32(_a_F_maybe_reread_subscription_3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	goto L76
L87:
	;
	v395 = v266
	v396 = v267
	goto L1
L88:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v335 != v336 {
		goto L2
	} else {
		goto L103
	}
L89:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+24)))
	if v279 != int32(1) {
		v333 = v277
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+16)))
	if v284 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+16)))
	if base.B2i32(v325 != int32(1))|base.B2i32(v324 != int32(3)) != 0 {
		v395 = v323
		v396 = v324
		goto L1
	} else {
		goto L102
	}
L92:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v310
	F_errmsg(m, v306, v9+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L100
	}
L93:
	;
	v300 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L10
	} else {
		goto L98
	}
L94:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v287 != int32(3) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v292 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	if v292 == int32(0) {
		goto L91
	} else {
		goto L97
	}
L97:
	;
	v306 = int32(_a_F_maybe_reread_subscription_7)
	v307 = int32(4054)
	goto L92
L98:
	;
	if v300 == int32(0) {
		goto L91
	} else {
		goto L99
	}
L99:
	;
	v306 = int32(_a_F_maybe_reread_subscription_8)
	v307 = int32(4058)
	goto L92
L100:
	;
	F_errfinish(m, int32(_a_F_maybe_reread_subscription_2), v307, int32(_a_F_maybe_reread_subscription_3))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	goto L91
L102:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5]))
	v333 = v332
	goto L88
L103:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	F_pfree(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
	F_pfree(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v333)+40))
	if v344 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_pfree(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L10
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v333)+48))
	F_list_free_deep(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L10
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	F_pfree(m, v333)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[2])) = v25
	*(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[5])) = v34
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	F_SetConfigOption(m, int32(_a_F_maybe_reread_subscription_9), v357, int32(4), int32(10))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	if v19 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L10
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v367 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[0])) = uint8(v367)
	goto L5
L116:
	;
	goto L115
L117:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_maybe_reread_subscription[4]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v383
	F_errmsg_internal(m, int32(_a_F_maybe_reread_subscription_10), v9+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_maybe_reread_subscription_2), int32(4067), int32(_a_F_maybe_reread_subscription_3))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v395)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L10
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L10
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_md5_text(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13966(m, l0, int32(_a_F_md5_text_0), int32(49))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_md_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l1 + int32(104)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v13
	v17 = base.I32_wrap_i64(int64(base.Ui64(v13) >> (uint(int64(32)) % 64)))
	if v13 < int64(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(257) - v17<<(uint(int32(9))%32)
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v27
		F_pgaio_result_report(m, v9, v12, int32(16))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	} else {
		v33 = int32(base.Ui32(v17) >> (uint(int32(13)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
		if base.Ui64(v13) <= base.Ui64(int64(35184372088831)) {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(257)
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v39
			F_pgaio_result_report(m, v9+int32(8), v12, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v46 = base.I32_wrap_i64(v13)
			if v46&int32(448) == int32(256) {
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
				if base.Ui32(v51) <= base.Ui32(v33) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v46&int32(-512) | int32(129)
				}
			}
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_mda_get_offset_values(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	v5 = int32(0)
	v11 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1+l0<<(uint(v11)%32)-int32(4)))) = v5
	v19 = l0 - v11
	if v5 <= v19 {
		v26 = v19
		v30 = v5
		for {
			v33 = v26 << (uint(int32(2)) % 32)
			v34 = l1 + v33
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+v33)))
			v37 = int32(1)
			v38 = v36 - v37
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v38
			v41 = v26 + v37
			if l0 <= v41 {
			} else {
				if v30&int32(1) == int32(0) {
					v47 = int32(2)
					v48 = v41 << (uint(v47) % 32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l3+v48)))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v48)))
					v56 = v38 - (v50-int32(1))*v54
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v56
					v60 = v26 + v47
					v61 = v56
				} else {
					v60 = v41
					v61 = v38
				}
				if v30 == int32(0) {
				} else {
					v68 = v60
					v69 = v61
					for {
						v74 = int32(2)
						v75 = v68 << (uint(v74) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l3+v75)))
						v78 = int32(1)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l2+v75)))
						v83 = v69 - (v77-v78)*v81
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v83
						v86 = v75 + int32(4)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l3+v86)))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+v86)))
						v94 = v83 - (v88-v78)*v92
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v94
						v97 = v68 + v74
						if v97 != l0 {
							v68 = v97
							v69 = v94
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v109 = int32(1)
			if int32(0) < v26 {
				v26 = v26 - v109
				v30 = v30 + v109
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_mdcbuf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+116))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	m.T0[v5].(func(*base.Module, int32))(m, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v9
		base.MemoryFill(m, l0, v9, int32(_a_F_mdcbuf_free_0))
		F_pfree(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_mdcbuf_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(_a_F_mdcbuf_init_0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(_a_F_mdcbuf_init_1)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = int32(1)
		return int32(0)
	}
}
func F_mdcbuf_read(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 != 0 {
		v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
		if l2 < v206 {
			v210 = l2
		} else {
			v210 = v206
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
		v216 = v210
		m.G0 = v15 + int32(48)
		return v216
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l2 <= v18 {
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
			if l2 < v206 {
				v210 = l2
			} else {
				v210 = v206
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
			v216 = v210
			m.G0 = v15 + int32(48)
			return v216
		} else {
			if v18 <= int32(0) {
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v26 = l0 + int32(46)
				if base.B2i32(v18 == int32(0))|base.B2i32(v24 == v26) != 0 {
				} else {
					base.MemoryCopy(m, v26, v24, v18)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l0 + int32(46)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v43 = F_pullf_read(m, l1, v35-(v18+v36)+int32(22), v15+int32(12))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				if v43 < int32(0) {
					v216 = v43
					m.G0 = v15 + int32(48)
					return v216
				} else {
					if v43 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v54 != int32(211) {
							v101 = int32(_a_F_mdcbuf_read_0)
							F_px_debug(m, v101, int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v216 = int32(-100)
								m.G0 = v15 + int32(48)
								return v216
							}
						} else {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
							if v58 != int32(20) {
								v101 = int32(_a_F_mdcbuf_read_0)
								F_px_debug(m, v101, int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v216 = int32(-100)
									m.G0 = v15 + int32(48)
									return v216
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+116))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
								m.T0[v66].(func(*base.Module, int32, int32, int32))(m, v62, l0+int32(24), int32(2))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+116))
									v72 = v15 + int32(16)
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
									m.T0[v73].(func(*base.Module, int32, int32))(m, v70, v72)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+42))
										v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+34))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+26))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
										v81 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
										base.MemoryFill(m, v72, int32(0), int32(20))
										if base.I64_extend_i32_u(v78^v76)|(v80^v79|(v81^v77)) == int64(0) {
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
											if l2 < v206 {
												v210 = l2
											} else {
												v210 = v206
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
											v216 = v210
											m.G0 = v15 + int32(48)
											return v216
										} else {
											v101 = int32(_a_F_mdcbuf_read_1)
											F_px_debug(m, v101, int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v216 = int32(-100)
												m.G0 = v15 + int32(48)
												return v216
											}
										}
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(22)) <= base.Ui32(v43) {
							v109 = l0 + int32(24)
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v110 != 0 {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								base.MemoryCopy(m, v111+v112, v109, v110)
							} else {
							}
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+116))
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
							m.T0[v117].(func(*base.Module, int32, int32, int32))(m, v116, v109, v110)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
								v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v123 = v122 + v110
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v123
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								v127 = v43 - int32(22)
								if v127 != 0 {
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									base.MemoryCopy(m, v128+v123, v125, v127)
								} else {
								}
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+116))
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
								m.T0[v133].(func(*base.Module, int32, int32, int32))(m, v132, v125, v127)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v136 + v127
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
									v140 = v127 + v139
									v141 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
									v142 = *(*int64)(unsafe.Add(mBase, uint32(v140)+8))
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
									v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v145 = v109 + v144
									v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+20)))
									*(*uint16)(unsafe.Add(mBase, uint32(v145)+20)) = uint16(v146)
									*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = v143
									*(*int64)(unsafe.Add(mBase, uint32(v145)+8)) = v142
									*(*int64)(unsafe.Add(mBase, uint32(v145))) = v141
									v190 = int32(22)
									v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v195 + v190
									v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
									if l2 < v206 {
										v210 = l2
									} else {
										v210 = v206
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
									v216 = v210
									m.G0 = v15 + int32(48)
									return v216
								}
							}
						} else {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v153 = v152 + v43
							if v153 < int32(23) {
								v179 = v152
								if v43 == int32(0) {
									v190 = v43
								} else {
									v187 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
									base.MemoryCopy(m, l0+v179+int32(24), v187, v43)
									v190 = v43
								}
								v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v195 + v190
								v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
								if l2 < v206 {
									v210 = l2
								} else {
									v210 = v206
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
								v216 = v210
								m.G0 = v15 + int32(48)
								return v216
							} else {
								v157 = l0 + int32(24)
								v159 = v153 - int32(22)
								if v159 != 0 {
									v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									base.MemoryCopy(m, v160+v161, v157, v159)
								} else {
								}
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+116))
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
								m.T0[v166].(func(*base.Module, int32, int32, int32))(m, v165, v157, v159)
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v169 + v159
									v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v173 = v172 - v159
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v173
									if v173 == int32(0) {
										v179 = v173
									} else {
										base.MemoryCopy(m, v157, v159+v157, v173)
										v179 = v173
									}
									if v43 == int32(0) {
										v190 = v43
									} else {
										v187 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
										base.MemoryCopy(m, l0+v179+int32(24), v187, v43)
										v190 = v43
									}
									v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v195 + v190
									v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v207
									if l2 < v206 {
										v210 = l2
									} else {
										v210 = v206
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206 - v210
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207 + v210
									v216 = v210
									m.G0 = v15 + int32(48)
									return v216
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_mdsyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v410 int64
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v443 int64
	_ = v443
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int64
	_ = v494
	var v498 int32
	_ = v498
	var v510 int64
	_ = v510
	var v514 int32
	_ = v514
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v535 int64
	_ = v535
	var v540 int32
	_ = v540
	var v551 int32
	_ = v551
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	v20 = F_smgropen(m, v11+int32(8), int32(-1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v28 = v20 + v25<<(uint(int32(2))%32)
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+40)))
	if base.Ui64(v24) < base.Ui64(v29) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v11 + int32(192)
	return v551
L4:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_mdsyncfiletag[0])))
	v399 = m.G0
	v401 = v399 - int32(16)
	m.G0 = v401
	if v396 != 0 {
		goto L104
	} else {
		goto L105
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.I32_wrap_i64(v24)<<(uint(int32(3))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_mdsyncfiletag[1]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36*int32(48))+32))
	goto L8
L6:
	;
	goto L7
L7:
	;
	v163 = v11 + int32(109)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	F_GetRelationPath(m, v163, v164, v165, v166, v167, v25)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L8:
	;
	goto L12
L9:
	;
	v393 = v36
	goto L4
L10:
	;
	v159 = F_strlen(m, v148)
	mBase = m.M
	goto L9
L12:
	;
	goto L13
L13:
	;
	v49 = int32(1023)
	if (l1^v42)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v152)
	goto L10
L15:
	;
	v133 = v128
	v134 = v129
	v135 = v130
	goto L36
L16:
	;
	if v123 == int32(0) {
		v148 = v121
		v149 = v122
		goto L14
	} else {
		goto L35
	}
L17:
	;
	v121 = v42
	v122 = l1
	v123 = v49
	goto L16
L18:
	;
	goto L19
L19:
	;
	v53 = int32(0)
	if base.B2i32(v42&int32(3) == v53)|int32(0) == v53 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v89 == int32(0) {
		v148 = v86
		v149 = v87
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v65 = v42
	v66 = l1
	v67 = v49
	goto L24
L22:
	;
	goto L23
L23:
	;
	v86 = v42
	v87 = l1
	v88 = v49
	v89 = int32(1)
	goto L20
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v69)
	if v69 == int32(0) {
		v128 = v65
		v129 = v66
		v130 = v67
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v86 = v80
	v87 = v74
	v88 = v76
	v89 = v78
	goto L20
L26:
	;
	v73 = int32(1)
	v74 = v66 + v73
	v76 = v67 - v73
	v77 = int32(0)
	v78 = base.B2i32(v76 != v77)
	v80 = v65 + v73
	if v80&int32(3) == v77 {
		v86 = v80
		v87 = v74
		v88 = v76
		v89 = v78
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v76 != 0 {
		v65 = v80
		v66 = v74
		v67 = v76
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if base.B2i32(v92 == int32(0))|base.B2i32(base.Ui32(v88) < base.Ui32(int32(4))) != 0 {
		v121 = v86
		v122 = v87
		v123 = v88
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v99 = v86
	v100 = v87
	v101 = v88
	goto L31
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v107 = int32(-2139062144)
	if (int32(16843008)-v104|v104)&v107 != v107 {
		v128 = v99
		v129 = v100
		v130 = v101
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v121 = v115
	v122 = v113
	v123 = v117
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
	v112 = int32(4)
	v113 = v100 + v112
	v115 = v99 + v112
	v117 = v101 - v112
	if base.Ui32(int32(3)) < base.Ui32(v117) {
		v99 = v115
		v100 = v113
		v101 = v117
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v128 = v121
	v129 = v122
	v130 = v123
	goto L15
L36:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	if v137 == int32(0) {
		v148 = v133
		v149 = v134
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v148 = v144
	v149 = v142
	goto L14
L38:
	;
	v141 = int32(1)
	v142 = v134 + v141
	v144 = v133 + v141
	v146 = v135 - v141
	if v146 != 0 {
		v133 = v144
		v134 = v142
		v135 = v146
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v170 = base.I32_wrap_i64(v24)
	if v170 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v257 = v11 + int32(109)
	base.MemoryCopy(m, v257, v11+int32(26), int32(83))
	goto L70
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v163
	v176 = F_pg_sprintf(m, v11+int32(26), int32(_a_F_mdsyncfiletag_0), v11)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v179 = v11 + int32(26)
	v181 = v11 + int32(109)
	if (v181^v179)&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	goto L41
L46:
	;
	goto L41
L47:
	;
	goto L46
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v235)
	if v235&int32(255) == int32(0) {
		goto L47
	} else {
		goto L63
	}
L49:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v234 = v181
	v235 = v187
	v236 = v179
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v181&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v191 = v181
	v193 = v179
	goto L55
L53:
	;
	v205 = v181
	v207 = v179
	goto L54
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v212 = int32(-2139062144)
	if (int32(16843008)-v209|v209)&v212 != v212 {
		v234 = v205
		v235 = v209
		v236 = v207
		goto L48
	} else {
		goto L59
	}
L55:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v194)
	if v194 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v205 = v201
	v207 = v199
	goto L54
L57:
	;
	v198 = int32(1)
	v199 = v193 + v198
	v201 = v191 + v198
	if v201&int32(3) != 0 {
		v191 = v201
		v193 = v199
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v217 = v205
	v218 = v209
	v219 = v207
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v218
	v221 = int32(4)
	v222 = v219 + v221
	v224 = v217 + v221
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v229 = int32(-2139062144)
	if (int32(16843008)-v226|v226)&v229 == v229 {
		v217 = v224
		v218 = v226
		v219 = v222
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v234 = v224
	v235 = v226
	v236 = v222
	goto L48
L62:
	;
	goto L61
L63:
	;
	v243 = v234
	v245 = v236
	goto L64
L64:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)) = uint8(v246)
	v248 = int32(1)
	if v246 != 0 {
		v243 = v243 + v248
		v245 = v245 + v248
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L47
L66:
	;
	goto L65
L67:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_mdsyncfiletag[2]))
	if v384&int32(1) != 0 {
		goto L98
	} else {
		goto L99
	}
L68:
	;
	v378 = F_strlen(m, v367)
	mBase = m.M
	goto L67
L70:
	;
	goto L71
L71:
	;
	v268 = int32(81)
	if (l1^v257)&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v371 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v368))) = uint8(v371)
	goto L68
L73:
	;
	v352 = v347
	v353 = v348
	v354 = v349
	goto L94
L74:
	;
	if v342 == int32(0) {
		v367 = v340
		v368 = v341
		goto L72
	} else {
		goto L93
	}
L75:
	;
	v340 = v257
	v341 = l1
	v342 = v268
	goto L74
L76:
	;
	goto L77
L77:
	;
	v272 = int32(0)
	if base.B2i32(v257&int32(3) == v272)|int32(0) == v272 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v308 == int32(0) {
		v367 = v305
		v368 = v306
		goto L72
	} else {
		goto L87
	}
L79:
	;
	v284 = v257
	v285 = l1
	v286 = v268
	goto L82
L80:
	;
	goto L81
L81:
	;
	v305 = v257
	v306 = l1
	v307 = v268
	v308 = int32(1)
	goto L78
L82:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v288)
	if v288 == int32(0) {
		v347 = v284
		v348 = v285
		v349 = v286
		goto L73
	} else {
		goto L84
	}
L83:
	;
	v305 = v299
	v306 = v293
	v307 = v295
	v308 = v297
	goto L78
L84:
	;
	v292 = int32(1)
	v293 = v285 + v292
	v295 = v286 - v292
	v296 = int32(0)
	v297 = base.B2i32(v295 != v296)
	v299 = v284 + v292
	if v299&int32(3) == v296 {
		v305 = v299
		v306 = v293
		v307 = v295
		v308 = v297
		goto L78
	} else {
		goto L85
	}
L85:
	;
	if v295 != 0 {
		v284 = v299
		v285 = v293
		v286 = v295
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if base.B2i32(v311 == int32(0))|base.B2i32(base.Ui32(v307) < base.Ui32(int32(4))) != 0 {
		v340 = v305
		v341 = v306
		v342 = v307
		goto L74
	} else {
		goto L88
	}
L88:
	;
	v318 = v305
	v319 = v306
	v320 = v307
	goto L89
L89:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v326 = int32(-2139062144)
	if (int32(16843008)-v323|v323)&v326 != v326 {
		v347 = v318
		v348 = v319
		v349 = v320
		goto L73
	} else {
		goto L91
	}
L90:
	;
	v340 = v334
	v341 = v332
	v342 = v336
	goto L74
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v323
	v331 = int32(4)
	v332 = v319 + v331
	v334 = v318 + v331
	v336 = v320 - v331
	if base.Ui32(int32(3)) < base.Ui32(v336) {
		v318 = v334
		v319 = v332
		v320 = v336
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v347 = v340
	v348 = v341
	v349 = v342
	goto L73
L94:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	*(*uint8)(unsafe.Add(mBase, uint32(v353))) = uint8(v356)
	if v356 == int32(0) {
		v367 = v352
		v368 = v353
		goto L72
	} else {
		goto L96
	}
L95:
	;
	v367 = v363
	v368 = v361
	goto L72
L96:
	;
	v360 = int32(1)
	v361 = v353 + v360
	v363 = v352 + v360
	v365 = v354 - v360
	if v365 != 0 {
		v352 = v363
		v353 = v361
		v354 = v365
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v387 = int32(_a_F_mdsyncfiletag_1)
	goto L100
L99:
	;
	v387 = int32(2)
	goto L100
L100:
	;
	v388 = F_PathNameOpenFile(m, l1, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if int32(0) <= v388 {
		v393 = v388
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v551 = int32(-1)
	goto L3
L103:
	;
	v415 = F_FileSync(m, v393, int32(167772182))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L107
	}
L104:
	;
	F___clock_gettime(m, int32(1), v401)
	mBase = m.M
	v405 = int64(*(*int32)(unsafe.Add(mBase, uint32(v401)+8)))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v401)))
	v410 = v405 + v406*int64(1000000000)
	goto L106
L105:
	;
	v410 = int64(0)
	goto L106
L106:
	;
	m.G0 = v401 + int32(16)
	goto L103
L107:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_mdsyncfiletag[3]))
	if base.Ui64(v29) <= base.Ui64(v24) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_FileClose(m, v393)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v424 = int32(1)
	v426 = int64(0)
	v430 = m.G0
	v432 = v430 - int32(16)
	m.G0 = v432
	if v410 != v426 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_mdsyncfiletag[3])) = v418
	v551 = v415
	goto L3
L113:
	;
	F___clock_gettime(m, int32(1), v432)
	mBase = m.M
	v438 = int64(*(*int32)(unsafe.Add(mBase, uint32(v432)+8)))
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	v443 = v438 + (v439*int64(1000000000) - v410)
	goto L117
L114:
	;
	goto L115
L115:
	;
	v530 = int32(200)
	v531 = *(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[4])) = v531 + base.I64_extend_i32_u(v424)
	v535 = *(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[5])) = v535 + v426
	F_pgstat_count_backend_io_op(m, int32(0), int32(3), v424, v424, v426)
	mBase = m.M
	v540 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_mdsyncfiletag[6])) = uint8(v540)
	*(*uint8)(unsafe.Add(mBase, _c_F_mdsyncfiletag[7])) = uint8(v540)
	m.G0 = v432 + int32(16)
	goto L112
L116:
	;
	v487 = int32(0)
	v493 = int32(200)
	v494 = *(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[8])) = v494 + v443
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_mdsyncfiletag[9]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v498))|base.B2i32(int32(1)<<(uint(v498)%32)&int32(_a_F_mdsyncfiletag_2) == v487) == v487 {
		goto L126
	} else {
		goto L127
	}
L117:
	;
	goto L119
L119:
	;
	goto L120
L120:
	;
	goto L116
L126:
	;
	v510 = *(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_mdsyncfiletag[10])) = v510 + v443
	v514 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_mdsyncfiletag[6])) = uint8(v514)
	*(*uint8)(unsafe.Add(mBase, _c_F_mdsyncfiletag[11])) = uint8(v514)
	goto L128
L127:
	;
	goto L128
L128:
	;
	goto L115
}
func F_merge_clump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v184 = F_lappend(m, v177, v178)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L14
	} else {
		goto L54
	}
L2:
	;
	v177 = int32(0)
	v178 = l2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = l1
	v14 = l2
	goto L5
L5:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 <= v20 {
		v140 = v21
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v144 == int32(1) {
		v177 = v13
		v178 = v14
		goto L1
	} else {
		goto L43
	}
L7:
	;
	goto L6
L8:
	;
	v29 = v20
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v140 = v134
	goto L7
L11:
	;
	v133 = v29 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v133 < v134 {
		v29 = v133
		goto L9
	} else {
		goto L42
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v51 = F_make_join_rel(m, l0, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L19
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v39 = F_have_relevant_joinclause(m, l0, v37, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v39 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v43 = F_have_join_order_restriction(m, l0, v37, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	if v51 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	F_generate_partitionwise_join_paths(m, l0, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v59 = int32(0)
	if base.B2i32(v57 == v59)|base.B2i32(v58 == v59) != 0 {
		v105 = base.B2i32(v57|v58 == v59)
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v105 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	goto L22
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v73 != v74 {
		v105 = int32(0)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v76 = int32(1)
	if v73 <= v76 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = v76
	goto L28
L27:
	;
	v79 = v73
	goto L28
L28:
	;
	v80 = int32(8)
	v85 = int32(0)
	goto L29
L29:
	;
	v93 = v85 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v57+v80+v93)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v58+v80+v93)))
	v98 = base.B2i32(v95 == v97)
	if v95 != v97 {
		v105 = v98
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v105 = v98
	goto L23
L31:
	;
	v101 = v85 + int32(1)
	if v101 != v79 {
		v85 = v101
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_generate_useful_gather_paths(m, l0, v51, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_set_cheapest(m, v51)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v51
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v118 + v119
	F_pfree(m, v14)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v124 = F_list_delete_nth_cell(m, v13, v29)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	if v124 != 0 {
		v13 = v124
		v14 = v36
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v127 = F_lappend(m, int32(0), v36)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	return v127
L42:
	;
	goto L10
L43:
	;
	if v140 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v150 = F_list_insert_nth(m, v13, int32(0), v14)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v160 = int32(0)
	goto L49
L47:
	;
	return v150
L48:
	;
	v173 = F_list_insert_nth(m, v13, v172, v14)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L53
	}
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v153+v160<<(uint(int32(2))%32))))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v167 < v144 {
		v172 = v160
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v172 = v140
	goto L48
L51:
	;
	v170 = v160 + int32(1)
	if v170 != v140 {
		v160 = v170
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	return v173
L54:
	;
	return v184
}
func F_miss(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v811 int32
	_ = v811
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v841 int32
	_ = v841
	var v868 int32
	_ = v868
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1132 int32
	_ = v1132
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1382 int32
	_ = v1382
	var v1398 int32
	_ = v1398
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1445 int32
	_ = v1445
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1519 int64
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1554 int64
	_ = v1554
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1588 int32
	_ = v1588
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1655 int32
	_ = v1655
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1707 int32
	_ = v1707
	var v1720 int32
	_ = v1720
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1745 int64
	_ = v1745
	var v1767 int32
	_ = v1767
	v7 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+l3<<(uint(int32(2))%32))))
	if v27 != 0 {
		v1767 = v27
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1767
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_miss[0]))
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v35 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L5
L8:
	;
	v44 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v92 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v60+v44<<(uint(int32(2))%32)))) = int32(0)
	v67 = v44 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v67 < v68 {
		v44 = v67
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+l3*int32(24))+20))
	v117 = v92
	v118 = v7
	v120 = v7
	v122 = int32(1)
	v124 = v7
	goto L17
L17:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(base.Ui32(v120)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v136)>>(uint(v120)%32))&int32(1) == int32(0) {
		v229 = v117
		v230 = v118
		v234 = v122
		v236 = v124
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v245 = int32(0)
	if v230 == v245 {
		v1767 = v245
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v243 = v120 + int32(1)
	if v243 < v229 {
		v117 = v229
		v118 = v230
		v120 = v243
		v122 = v234
		v124 = v236
		goto L17
	} else {
		goto L31
	}
L20:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v120<<(uint(int32(2))%32))))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146))))
	if v147 == int32(_a_F_miss_0) {
		v229 = v117
		v230 = v118
		v234 = v122
		v236 = v124
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v156 = v146
	v158 = v147
	v160 = v118
	v164 = v122
	v166 = v124
	goto L22
L22:
	;
	v172 = int32(_a_F_miss_0)
	v173 = v158 & v172
	if base.B2i32(v173 != l3&v172)&(base.B2i32(v173 != int32(_a_F_miss_1))|base.B2i32(v102&int32(2) != int32(0))) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v229 = v219
	v230 = v211
	v234 = v212
	v236 = v213
	goto L19
L24:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v189 = v183 + int32(base.Ui32(v184)>>(uint(int32(3))%32))&int32(536870908)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v191 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v190 | v191<<(uint(v184)%32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v196 == v197 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v211 = v160
	v212 = v164
	v213 = v166
	goto L26
L26:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)))
	if v214 != int32(_a_F_miss_0) {
		v156 = v156 + int32(8)
		v158 = v214
		v160 = v211
		v164 = v212
		v166 = v213
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v199 = v191
	goto L29
L28:
	;
	v199 = v166
	goto L29
L29:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v196))))
	v204 = int32(1)
	v211 = v204
	v212 = (int32(0) - v203&v204) & v164
	v213 = v199
	goto L26
L30:
	;
	goto L23
L31:
	;
	goto L18
L32:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v248&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v252 = l0
	v253 = l1
	v254 = l2
	v255 = l3
	v256 = l4
	v257 = l5
	v258 = int32(0)
	v261 = v229
	v263 = v28
	v266 = v234
	v268 = v236
	v270 = v245
	goto L36
L34:
	;
	v989 = l0
	v990 = l1
	v991 = l2
	v992 = l3
	v993 = l4
	v994 = l5
	v1003 = v234
	v1005 = v236
	v1007 = v245
	v1011 = int32(1)
	goto L35
L35:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v990)+28))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	if v1013 == int32(1) {
		goto L160
	} else {
		goto L161
	}
L36:
	;
	v274 = int32(0)
	if v261 <= v274 {
		v964 = v252
		v965 = v253
		v966 = v254
		v967 = v255
		v968 = v256
		v969 = v257
		v970 = v258
		v978 = v266
		v980 = v268
		v982 = v270
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v989 = v964
	v990 = v965
	v991 = v966
	v992 = v967
	v993 = v968
	v994 = v969
	v1003 = v978
	v1005 = v980
	v1007 = v982
	v1011 = base.B2i32(v970 == int32(0))
	goto L35
L38:
	;
	goto L37
L39:
	;
	v284 = v258
	v287 = v261
	v290 = v274
	v292 = v266
	v294 = v268
	v295 = v274
	goto L40
L40:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v253)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300+int32(base.Ui32(v290)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v306)>>(uint(v290)%32))&int32(1) == int32(0) {
		v945 = v284
		v948 = v287
		v953 = v292
		v955 = v294
		v956 = v295
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v956 != 0 {
		v258 = v945
		v261 = v948
		v266 = v953
		v268 = v955
		goto L36
	} else {
		goto L158
	}
L42:
	;
	v962 = v290 + int32(1)
	if v962 < v948 {
		v284 = v945
		v287 = v948
		v290 = v962
		v292 = v953
		v294 = v955
		v295 = v956
		goto L40
	} else {
		goto L157
	}
L43:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v263)+32))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312+v290<<(uint(int32(2))%32))))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316))))
	if v317 == int32(_a_F_miss_0) {
		v945 = v284
		v948 = v287
		v953 = v292
		v955 = v294
		v956 = v295
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v326 = v284
	v327 = v317
	v328 = v316
	v334 = v292
	v336 = v294
	v337 = v295
	goto L45
L45:
	;
	v342 = base.I32_extend16_s(v327)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v342 < v343 {
		v917 = v326
		v925 = v334
		v927 = v336
		v928 = v337
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v945 = v917
	v948 = v938
	v953 = v925
	v955 = v927
	v956 = v928
	goto L42
L47:
	;
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+8)))
	if v935 != int32(_a_F_miss_0) {
		v326 = v917
		v327 = v935
		v328 = v328 + int32(8)
		v334 = v925
		v336 = v927
		v337 = v928
		goto L45
	} else {
		goto L156
	}
L48:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v253)+28))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345+int32(base.Ui32(v346)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v352)>>(uint(v346)%32))&int32(1) != 0 {
		v917 = v326
		v925 = v334
		v927 = v336
		v928 = v337
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+28))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v359 = m.T0[v358].(func(*base.Module) int32)(m)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	if v359 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	if v361 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+424))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v370 = v342 - v369
	v372 = v370 << (uint(int32(2)) % 32)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372+v373)))
	if v375 != 0 {
		v393 = v375
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v363 = v361
	goto L56
L55:
	;
	v363 = int32(19)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+36)) = v363
	return int32(0)
L57:
	;
	if v868 != 0 {
		goto L150
	} else {
		goto L151
	}
L58:
	;
	if v841 == int32(0) {
		v917 = int32(1)
		v925 = v334
		v927 = v336
		v928 = v337
		goto L47
	} else {
		goto L149
	}
L59:
	;
	v396 = v368 + v370*int32(88)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+2)))
	if v397&int32(2) != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v384 = F_newdfa(m, v252, v368+v370*int32(88)+int32(36), v367+int32(72), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v386+v372))) = v384
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389+v372)))
	if v391 != 0 {
		v393 = v391
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	v841 = v392
	goto L58
L63:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
	v401 = int32(0)
	v403 = F_shortest(m, v252, v393, v256, v256, v400, v401, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v393)+40))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+8)))
	if v415&int32(2) != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	v406 = int32(0)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+2)))
	if base.B2i32(v403 == v406)^base.B2i32(v408&int32(1) == v406) != 0 {
		v841 = v405
		goto L58
	} else {
		goto L67
	}
L67:
	;
	v868 = v405
	goto L57
L68:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+2)))
	if v811 != (v827^int32(-1))&int32(1) {
		v868 = v826
		goto L57
	} else {
		goto L148
	}
L69:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+40))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v811 = base.B2i32(base.Ui32(v418) <= base.Ui32((v256-v419)>>(uint(int32(2))%32)))
	goto L68
L70:
	;
	goto L71
L71:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v252)+48))
	v425 = v424 + v372
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v393)+44))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v252)+52))
	v428 = v427 + v372
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	if base.Ui32(v429) <= base.Ui32(v256) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if base.Ui32(v256) <= base.Ui32(v681) {
		goto L121
	} else {
		goto L122
	}
L73:
	;
	v432 = v429
	goto L75
L74:
	;
	v432 = int32(0)
	goto L75
L75:
	;
	if v432 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v435 = int32(0)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v435 < v445 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	goto L78
L78:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v677 != 0 {
		v679 = v677
		v681 = v429
		goto L72
	} else {
		goto L119
	}
L79:
	;
	if v658 == int32(0) {
		v811 = v435
		goto L68
	} else {
		goto L116
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+20)) = v436
	*(*int64)(unsafe.Add(mBase, uint32(v393)+48)) = int64(0)
	v658 = v637
	goto L79
L81:
	;
	v612 = int32(0)
	goto L113
L82:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v393)+20))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)))
	if v449&int32(1) != 0 {
		v604 = v448
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v453 = F_getvacant(m, v252, v393, v436, v436)
	mBase = m.M
	if v453 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v658 = int32(0)
	goto L79
L87:
	;
	goto L88
L88:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if int32(0) < v457 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v461 = int32(0)
	goto L92
L90:
	;
	goto L91
L91:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v393)+40))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v500 = v493 + int32(base.Ui32(v495)>>(uint(int32(3))%32))&int32(536870908)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	v502 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501 | v502<<(uint(v495)%32)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if v507 == v502 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	*(*int32)(unsafe.Add(mBase, uint32(v472+v461<<(uint(int32(2))%32)))) = int32(0)
	v479 = v461 + int32(1)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	if v479 < v480 {
		v461 = v479
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	goto L93
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v586
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v597 <= int32(0) {
		v637 = v453
		goto L80
	} else {
		goto L112
	}
L96:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506)))
	v586 = v510
	goto L95
L97:
	;
	goto L98
L98:
	;
	if v507 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v586 = int32(0)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v515 = v507 & int32(3)
	v516 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v507) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v522 = v516
	v525 = v516
	v531 = v435
	goto L105
L103:
	;
	v551 = v516
	v554 = v516
	goto L104
L104:
	;
	v562 = v551
	v565 = v554
	v572 = v435
	goto L109
L105:
	;
	v535 = v506 + v522<<(uint(int32(2))%32)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v543 = v536 ^ (v537 ^ (v538 ^ (v539 ^ v525)))
	v544 = int32(4)
	v545 = v522 + v544
	v547 = v531 + v544
	if v547 != v507&int32(2147483644) {
		v522 = v545
		v525 = v543
		v531 = v547
		goto L105
	} else {
		goto L107
	}
L106:
	;
	if v515 == int32(0) {
		v586 = v543
		goto L95
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v551 = v545
	v554 = v543
	goto L104
L109:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v506+v562<<(uint(int32(2))%32))))
	v577 = v576 ^ v565
	v578 = int32(1)
	v581 = v572 + v578
	if v581 != v515 {
		v562 = v562 + v578
		v565 = v577
		v572 = v581
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v586 = v577
	goto L95
L111:
	;
	goto L110
L112:
	;
	v604 = v453
	goto L81
L113:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v393)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v623+v612<<(uint(int32(5))%32))+20)) = int32(0)
	v630 = v612 + int32(1)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v630 < v631 {
		v612 = v630
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v637 = v604
	goto L80
L115:
	;
	goto L114
L116:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v393)+40))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v665 = int32(1)
	v670 = int32(*(*int16)(unsafe.Add(mBase, uint32(v661+(v662^int32(-1))&v665<<(uint(v665)%32))+20)))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v672 = F_miss(m, v252, v393, v658, v670, v436, v671)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	if v672 == int32(0) {
		v811 = v435
		goto L68
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v672)+20)) = v436
	v679 = v672
	v681 = v436
	goto L72
L119:
	;
	v811 = int32(0)
	goto L68
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v743
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
	if base.Ui32(v743) < base.Ui32(v759) {
		goto L137
	} else {
		goto L138
	}
L121:
	;
	v742 = v679
	v743 = v681
	goto L120
L122:
	;
	goto L123
L123:
	;
	v689 = v679
	v691 = v681
	goto L124
L124:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	if base.Ui32(v705) <= base.Ui32(int32(2047)) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v742 = v730
	v743 = v732
	goto L120
L126:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v689)+24))
	v716 = base.I32_extend16_s(v714)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v715+v716<<(uint(int32(2))%32))))
	if v720 != 0 {
		v730 = v720
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708+v705<<(uint(int32(1))%32)))))
	v714 = v712
	goto L126
L128:
	;
	goto L129
L129:
	;
	v713 = F_pg_reg_getcolor(m, v426, v705)
	mBase = m.M
	v714 = v713
	goto L126
L130:
	;
	v732 = v691 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v730)+20)) = v732
	if base.Ui32(v732) < base.Ui32(v256) {
		v689 = v730
		v691 = v732
		goto L124
	} else {
		goto L134
	}
L131:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v724 = F_miss(m, v252, v393, v689, v716, v691+int32(4), v723)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	if v724 != 0 {
		v730 = v724
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v726 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v691
	v811 = v726
	goto L68
L134:
	;
	goto L125
L135:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v797)+8))
	v800 = int32(1)
	v811 = int32(base.Ui32(v799)>>(uint(v800)%32)) & v800
	goto L68
L136:
	;
	if v795 != 0 {
		v797 = v795
		goto L135
	} else {
		goto L147
	}
L137:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	if base.Ui32(v761) <= base.Ui32(int32(2047)) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	goto L139
L139:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v393)+40))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v789 = int32(*(*int16)(unsafe.Add(mBase, uint32(v782+(v783^int32(-1))&int32(2))+24)))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v791 = F_miss(m, v252, v393, v742, v789, v743, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L146
	}
L140:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v742)+24))
	v772 = base.I32_extend16_s(v770)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v771+v772<<(uint(int32(2))%32))))
	if v776 != 0 {
		v797 = v776
		goto L135
	} else {
		goto L144
	}
L141:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	v768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v764+v761<<(uint(int32(1))%32)))))
	v770 = v768
	goto L140
L142:
	;
	goto L143
L143:
	;
	v769 = F_pg_reg_getcolor(m, v426, v761)
	mBase = m.M
	v770 = v769
	goto L140
L144:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	v780 = F_miss(m, v252, v393, v742, v772, v743+int32(4), v779)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v795 = v780
	goto L136
L146:
	;
	v795 = v791
	goto L136
L147:
	;
	v811 = int32(0)
	goto L68
L148:
	;
	v841 = v826
	goto L58
L149:
	;
	return int32(0)
L150:
	;
	return int32(0)
L151:
	;
	goto L152
L152:
	;
	v884 = int32(1)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v253)+28))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v891 = v885 + int32(base.Ui32(v886)>>(uint(int32(3))%32))&int32(536870908)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)))
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v892 | v884<<(uint(v886)%32)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v898 == v899 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v901 = v884
	goto L155
L154:
	;
	v901 = v336
	goto L155
L155:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v263)+28))
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903+v898))))
	v906 = int32(1)
	v917 = v906
	v925 = (int32(0) - v905&v906) & v334
	v927 = v901
	v928 = v884
	goto L47
L156:
	;
	goto L46
L157:
	;
	goto L41
L158:
	;
	v964 = v252
	v965 = v253
	v966 = v254
	v967 = v255
	v968 = v256
	v969 = v257
	v970 = v945
	v978 = v953
	v980 = v955
	v982 = v270
	goto L38
L159:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	if int32(0) < v1147 {
		goto L178
	} else {
		goto L179
	}
L160:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1012)))
	v1132 = v1016
	goto L159
L161:
	;
	goto L162
L162:
	;
	if v1013 <= int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1132 = int32(0)
	goto L159
L164:
	;
	goto L165
L165:
	;
	v1021 = v1013 & int32(3)
	v1022 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1013) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1036 = v1022
	v1038 = v1022
	v1040 = int32(0)
	goto L169
L167:
	;
	v1076 = v1022
	v1078 = v1022
	goto L168
L168:
	;
	v1098 = v1076
	v1099 = v1022
	v1100 = v1078
	goto L173
L169:
	;
	v1054 = v1012 + v1036<<(uint(int32(2))%32)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+8))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1062 = v1055 ^ (v1056 ^ (v1057 ^ (v1058 ^ v1038)))
	v1063 = int32(4)
	v1064 = v1036 + v1063
	v1066 = v1040 + v1063
	if v1066 != v1013&int32(2147483644) {
		v1036 = v1064
		v1038 = v1062
		v1040 = v1066
		goto L169
	} else {
		goto L171
	}
L170:
	;
	if v1021 == int32(0) {
		v1132 = v1062
		goto L159
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v1076 = v1064
	v1078 = v1062
	goto L168
L173:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1012+v1098<<(uint(int32(2))%32))))
	v1118 = v1117 ^ v1100
	v1119 = int32(1)
	v1122 = v1099 + v1119
	if v1122 != v1021 {
		v1098 = v1098 + v1119
		v1099 = v1122
		v1100 = v1118
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v1132 = v1118
	goto L159
L175:
	;
	goto L174
L176:
	;
	if v1011 != 0 {
		goto L296
	} else {
		goto L297
	}
L177:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	if v1277 < v1278 {
		goto L210
	} else {
		goto L211
	}
L178:
	;
	v1151 = v1013 << (uint(int32(2)) % 32)
	v1160 = v1146
	v1161 = v1147
	goto L181
L179:
	;
	goto L180
L180:
	;
	if v1147 != 0 {
		v1720 = v1146
		goto L176
	} else {
		goto L207
	}
L181:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+4))
	if v1132 == v1176 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L177
L183:
	;
	if v1013 == int32(1) {
		v1720 = v1160
		goto L176
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v1245 = int32(1)
	if v1245 < v1161 {
		v1160 = v1160 + int32(32)
		v1161 = v1161 - v1245
		goto L181
	} else {
		goto L206
	}
L186:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1160)))
	if base.Ui32(int32(4)) <= base.Ui32(v1151) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	if v1240 == int32(0) {
		v1720 = v1160
		goto L176
	} else {
		goto L205
	}
L188:
	;
	v1240 = int32(0)
	goto L187
L189:
	;
	v1214 = v1209
	v1215 = v1210
	v1216 = v1211
	goto L199
L190:
	;
	if (v1012|v1178)&int32(3) != 0 {
		v1209 = v1012
		v1210 = v1178
		v1211 = v1151
		goto L189
	} else {
		goto L193
	}
L191:
	;
	v1202 = v1012
	v1203 = v1178
	v1204 = v1151
	goto L192
L192:
	;
	if v1204 == int32(0) {
		goto L188
	} else {
		goto L198
	}
L193:
	;
	v1186 = v1012
	v1187 = v1178
	v1188 = v1151
	goto L194
L194:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	if v1191 != v1192 {
		v1209 = v1186
		v1210 = v1187
		v1211 = v1188
		goto L189
	} else {
		goto L196
	}
L195:
	;
	v1202 = v1197
	v1203 = v1195
	v1204 = v1199
	goto L192
L196:
	;
	v1194 = int32(4)
	v1195 = v1187 + v1194
	v1197 = v1186 + v1194
	v1199 = v1188 - v1194
	if base.Ui32(int32(3)) < base.Ui32(v1199) {
		v1186 = v1197
		v1187 = v1195
		v1188 = v1199
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v1209 = v1202
	v1210 = v1203
	v1211 = v1204
	goto L189
L199:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214))))
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1215))))
	if v1219 == v1220 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v1240 = v1219 - v1220
	goto L187
L201:
	;
	v1222 = int32(1)
	v1227 = v1216 - v1222
	if v1227 != 0 {
		v1214 = v1214 + v1222
		v1215 = v1215 + v1222
		v1216 = v1227
		goto L199
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	goto L200
L204:
	;
	goto L188
L205:
	;
	goto L185
L206:
	;
	goto L182
L207:
	;
	goto L177
L208:
	;
	if v1641 == int32(0) {
		v1767 = v1007
		goto L1
	} else {
		goto L285
	}
L209:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+12))
	if v1453 != 0 {
		goto L249
	} else {
		goto L250
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v990)+4)) = v1277 + int32(1)
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v990)+24))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	v1285 = int32(0)
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	v1289 = v1286 + v1277<<(uint(int32(5))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1289)+16)) = uint16(v1285)
	*(*int64)(unsafe.Add(mBase, uint32(v1289)+8)) = int64(0)
	v1295 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1289))) = v1283 + v1277*v1284<<(uint(v1295)%32)
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v990)+32))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1289)+24)) = v1299 + v1300*v1277<<(uint(v1295)%32)
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v990)+36))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1289)+28)) = v1306 + v1307*v1277<<(uint(int32(3))%32)
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	if v1313 <= v1285 {
		v1445 = v1289
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1345 = base.I32_div_s(v1278<<(uint(int32(1))%32), int32(3))
	v1346 = int32(2)
	if v1345 < (v993-v994)>>(uint(v1346)%32) {
		goto L217
	} else {
		goto L218
	}
L213:
	;
	v1319 = v1285
	goto L214
L214:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+24))
	v1330 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1326+v1319<<(uint(int32(2))%32)))) = v1330
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1332+v1319<<(uint(int32(3))%32)))) = v1330
	v1339 = v1319 + int32(1)
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	if v1339 < v1340 {
		v1319 = v1339
		goto L214
	} else {
		goto L216
	}
L215:
	;
	v1445 = v1289
	goto L209
L216:
	;
	goto L215
L217:
	;
	v1353 = v993 - v1345<<(uint(v1346)%32)
	goto L219
L218:
	;
	v1353 = v994
	goto L219
L219:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v990)+56))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v990)+20))
	v1358 = v1355 + v1278<<(uint(int32(5))%32)
	if base.Ui32(v1354) < base.Ui32(v1358) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v990)+56)) = v1432 + int32(32)
	v1445 = v1432
	goto L209
L221:
	;
	v1362 = v1354
	goto L224
L222:
	;
	goto L223
L223:
	;
	if base.Ui32(v1355) < base.Ui32(v1354) {
		goto L234
	} else {
		goto L235
	}
L224:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+20))
	if base.Ui32(v1353) <= base.Ui32(v1370) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L223
L226:
	;
	v1373 = v1370
	goto L228
L227:
	;
	v1373 = int32(0)
	goto L228
L228:
	;
	if v1373 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1362)+8)))
	if v1376&int32(4) == int32(0) {
		v1432 = v1362
		goto L220
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1382 = v1362 + int32(32)
	if base.Ui32(v1382) < base.Ui32(v1358) {
		v1362 = v1382
		goto L224
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	goto L225
L234:
	;
	v1398 = v1355
	goto L237
L235:
	;
	goto L236
L236:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v989)+36))
	if v1425 != 0 {
		goto L246
	} else {
		goto L247
	}
L237:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+20))
	if base.Ui32(v1353) <= base.Ui32(v1405) {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L236
L239:
	;
	v1413 = v1398 + int32(32)
	if base.Ui32(v1413) < base.Ui32(v1354) {
		v1398 = v1413
		goto L237
	} else {
		goto L245
	}
L240:
	;
	v1408 = v1405
	goto L242
L241:
	;
	v1408 = int32(0)
	goto L242
L242:
	;
	if v1408 != 0 {
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1398)+8)))
	if v1409&int32(4) != 0 {
		goto L239
	} else {
		goto L244
	}
L244:
	;
	v1432 = v1398
	goto L220
L245:
	;
	goto L238
L246:
	;
	v1427 = v1425
	goto L248
L247:
	;
	v1427 = int32(15)
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+36)) = v1427
	v1641 = int32(0)
	goto L208
L249:
	;
	v1454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1445)+16)))
	v1455 = v1454
	v1458 = v1453
	goto L252
L250:
	;
	goto L251
L251:
	;
	v1490 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1445)+12)) = v1490
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	if v1490 < v1492 {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+24))
	v1466 = base.I32_extend16_s(v1455)
	v1470 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1465+v1466<<(uint(int32(2))%32)))) = v1470
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+28))
	v1475 = v1472 + v1466<<(uint(int32(3))%32)
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1475)))
	*(*int32)(unsafe.Add(mBase, uint32(v1475))) = v1470
	v1479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1475)+4)))
	if v1476 != 0 {
		v1455 = v1479
		v1458 = v1476
		goto L252
	} else {
		goto L254
	}
L253:
	;
	goto L251
L254:
	;
	goto L253
L255:
	;
	v1496 = v1492
	v1501 = int32(0)
	goto L258
L256:
	;
	goto L257
L257:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+8))
	if v1600&int32(2) == int32(0) {
		v1615 = v1600
		goto L273
	} else {
		goto L274
	}
L258:
	;
	v1507 = v1501 << (uint(int32(2)) % 32)
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+24))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1507+v1508)))
	if v1510 != 0 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	goto L257
L260:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+12))
	if v1511 != v1445 {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v1577 = v1496
	goto L262
L262:
	;
	v1588 = v1501 + int32(1)
	if v1588 < v1577 {
		v1496 = v1577
		v1501 = v1588
		goto L258
	} else {
		goto L272
	}
L263:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+24))
	v1568 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1566+v1507))) = v1568
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1570+v1501<<(uint(int32(3))%32)))) = v1568
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v990)+12))
	v1577 = v1576
	goto L262
L264:
	;
	v1521 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1510)+16)))
	v1522 = v1511
	v1525 = v1521
	goto L268
L265:
	;
	v1513 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1510)+16)))
	if v1501 != v1513 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+28))
	v1519 = *(*int64)(unsafe.Add(mBase, uint32(v1515+v1501<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1510)+12)) = v1519
	goto L263
L267:
	;
	v1547 = int32(3)
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+28))
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(v1550+v1501<<(uint(v1547)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1546+v1545<<(uint(v1547)%32)))) = v1554
	goto L263
L268:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+28))
	v1535 = v1532 + v1525<<(uint(int32(3))%32)
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1535)))
	if v1536 == int32(0) {
		v1545 = v1525
		v1546 = v1532
		goto L267
	} else {
		goto L270
	}
L269:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+28))
	v1545 = base.I32_extend16_s(v1525)
	v1546 = v1543
	goto L267
L270:
	;
	v1540 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1535)+4)))
	if base.B2i32(v1536 != v1445)|base.B2i32(v1501 != v1540) != 0 {
		v1522 = v1536
		v1525 = v1540
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	goto L259
L273:
	;
	if v1615&int32(8) == int32(0) {
		goto L279
	} else {
		goto L280
	}
L274:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+20))
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v990)+48))
	if base.Ui32(v1605) <= base.Ui32(v1606) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1610 = v1606
	goto L277
L276:
	;
	v1610 = int32(0)
	goto L277
L277:
	;
	if base.B2i32(v1605 == v1606)|v1610 != 0 {
		v1615 = v1600
		goto L273
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v990)+48)) = v1605
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+8))
	v1615 = v1613
	goto L273
L279:
	;
	v1641 = v1445
	goto L208
L280:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+20))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v990)+52))
	if base.Ui32(v1621) <= base.Ui32(v1622) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1626 = v1622
	goto L283
L282:
	;
	v1626 = int32(0)
	goto L283
L283:
	;
	if base.B2i32(v1621 == v1622)|v1626 != 0 {
		goto L279
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v990)+52)) = v1621
	goto L279
L285:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	if int32(0) < v1644 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1655 = int32(0)
	goto L289
L287:
	;
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+4)) = v1132
	if v1005 != 0 {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	v1671 = v1655 << (uint(int32(2)) % 32)
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v990)+28))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1674+v1671)))
	*(*int32)(unsafe.Add(mBase, uint32(v1671+v1672))) = v1676
	v1679 = v1655 + int32(1)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v990)+16))
	if v1679 < v1680 {
		v1655 = v1679
		goto L289
	} else {
		goto L291
	}
L290:
	;
	goto L288
L291:
	;
	goto L290
L292:
	;
	v1707 = int32(2)
	goto L294
L293:
	;
	v1707 = int32(0)
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+8)) = v1707
	if v1003 == int32(0) {
		v1720 = v1641
		goto L176
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+8)) = v1707 | int32(8)
	v1720 = v1641
	goto L176
L296:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v991)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1736+v992<<(uint(int32(2))%32)))) = v1720
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v991)+28))
	v1745 = *(*int64)(unsafe.Add(mBase, uint32(v1720)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1741+v992<<(uint(int32(3))%32)))) = v1745
	*(*uint16)(unsafe.Add(mBase, uint32(v1720)+16)) = uint16(v992)
	*(*int32)(unsafe.Add(mBase, uint32(v1720)+12)) = v991
	goto L298
L297:
	;
	goto L298
L298:
	;
	v1767 = v1720
	goto L1
}
func F_missing_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 != v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v6 < v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v11 = int32(1)
	goto L6
L5:
	;
	v11 = int32(-1)
	goto L6
L6:
	;
	return v11
L7:
	;
	return v76
L8:
	;
	v76 = int32(0)
	goto L7
L9:
	;
	v50 = v45
	v51 = v46
	v52 = v47
	goto L19
L10:
	;
	if (v13|v14)&int32(3) != 0 {
		v45 = v13
		v46 = v14
		v47 = v5
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v38 = v13
	v39 = v14
	v40 = v5
	goto L12
L12:
	;
	if v40 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v22 = v13
	v23 = v14
	v24 = v5
	goto L14
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 != v28 {
		v45 = v22
		v46 = v23
		v47 = v24
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v38 = v33
	v39 = v31
	v40 = v35
	goto L12
L16:
	;
	v30 = int32(4)
	v31 = v23 + v30
	v33 = v22 + v30
	v35 = v24 - v30
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		v22 = v33
		v23 = v31
		v24 = v35
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v45 = v38
	v46 = v39
	v47 = v40
	goto L9
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v76 = v55 - v56
	goto L7
L21:
	;
	v58 = int32(1)
	v63 = v52 - v58
	if v63 != 0 {
		v50 = v50 + v58
		v51 = v51 + v58
		v52 = v63
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L8
}
func F_mix_decrypt_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 <= int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2 + v9
		return l2
	} else {
		v20 = l1
		v22 = l3
		v23 = v9
		for {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(84))))) = uint8(v28)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(52))))))
			v32 = v28 ^ v31
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v32)
			v34 = int32(1)
			v39 = v23 + v34
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = v40 + l2
			if v39 < v41 {
				v20 = v20 + v34
				v22 = v22 + v34
				v23 = v39
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		return l2
	}
}
func F_mix_encrypt_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 <= int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2 + v9
		return l2
	} else {
		v20 = l1
		v22 = l3
		v23 = v9
		for {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(52))))))
			v31 = v28 ^ v30
			*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(84))))) = uint8(v31)
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v31)
			v34 = int32(1)
			v39 = v23 + v34
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = v40 + l2
			if v39 < v41 {
				v20 = v20 + v34
				v22 = v22 + v34
				v23 = v39
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		return l2
	}
}
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var __phi73 int32
	_ = __phi73
	var v75 int32
	_ = v75
	var __phi75 int32
	_ = __phi75
	var v76 int32
	_ = v76
	var __phi76 int32
	_ = __phi76
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1138 int32
	_ = v1138
	var v1146 int32
	_ = v1146
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1223 int32
	_ = v1223
	var v1230 int32
	_ = v1230
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1329 int32
	_ = v1329
	var v1336 int32
	_ = v1336
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1394 int32
	_ = v1394
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1466 int32
	_ = v1466
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1514 int64
	_ = v1514
	var v1516 int64
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1525 int64
	_ = v1525
	var v1527 int64
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1534 int64
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1587 int64
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1598 int64
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1609 int64
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1620 int64
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1634 int32
	_ = v1634
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1698 int64
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1738 int32
	_ = v1738
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1781 int64
	_ = v1781
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1795 int64
	_ = v1795
	var v1796 int64
	_ = v1796
	var v1802 int64
	_ = v1802
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1833 int64
	_ = v1833
	var v1840 int32
	_ = v1840
	var v1841 int64
	_ = v1841
	var v1842 int64
	_ = v1842
	var v1847 int64
	_ = v1847
	var v1851 int64
	_ = v1851
	var v1852 int64
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1860 int64
	_ = v1860
	var v1861 int64
	_ = v1861
	var v1865 int64
	_ = v1865
	var v1869 int64
	_ = v1869
	var v1870 int64
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1881 int32
	_ = v1881
	var v1905 int64
	_ = v1905
	var v1912 int32
	_ = v1912
	var v1913 int64
	_ = v1913
	var v1914 int64
	_ = v1914
	var v1916 int64
	_ = v1916
	var v1919 int64
	_ = v1919
	var v1982 int64
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int64
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2044 int64
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int64
	_ = v2053
	var v2057 int64
	_ = v2057
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2072 int32
	_ = v2072
	var v2101 int32
	_ = v2101
	var v2102 int64
	_ = v2102
	var v2106 int64
	_ = v2106
	var v2137 int64
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2212 int32
	_ = v2212
	var v2231 int64
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2240 int64
	_ = v2240
	var v2241 int64
	_ = v2241
	var v2245 int32
	_ = v2245
	var v2249 int64
	_ = v2249
	var v2253 int64
	_ = v2253
	var v2254 int64
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2297 int32
	_ = v2297
	var v2309 int32
	_ = v2309
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2360 int32
	_ = v2360
	var v2364 int32
	_ = v2364
	var v2370 int32
	_ = v2370
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2462 int32
	_ = v2462
	var v2467 int32
	_ = v2467
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2514 int32
	_ = v2514
	var v2548 int32
	_ = v2548
	var v2552 int32
	_ = v2552
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2602 int64
	_ = v2602
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v32 < v33 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v2596 != 0 {
		goto L238
	} else {
		goto L239
	}
L2:
	;
	v35 = v32
	goto L4
L3:
	;
	v35 = v33
	goto L4
L4:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = base.B2i32(v33 < v32)
	if v33 < v32 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = v32
	goto L10
L9:
	;
	v39 = v33
	goto L10
L10:
	;
	if v33 < v32 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = l0
	goto L13
L12:
	;
	v40 = l1
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v33 < v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = l1
	goto L16
L15:
	;
	v42 = l0
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if int32(6) < v35 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v1256 = v33 + v32
	v1257 = int32(2)
	v1258 = base.I32_div_s(v1256, v1257)
	v1259 = int32(1)
	v1260 = v1258 + v1259
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v1275 = v1261 + (v1262 + ((v35^int32(-1))&v1259 - (v1256 + v39&v1259) + v1260<<(uint(v1259)%32)))
	v1276 = int32(3)
	v1279 = base.I32_div_s(l3+v1276, int32(4))
	v1284 = base.I32_div_s(v1275+v1279+v1276, v1257)
	v1286 = v1284 + v1259
	if v1260 < v1286 {
		goto L83
	} else {
		goto L84
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if l3 != v46+v47 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v54 = v33 + v32
	v56 = v54 << (uint(int32(1)) % 32)
	v59 = F_palloc(m, v56+int32(2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v61 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v59))) = uint16(v61)
	v64 = v59 + int32(2)
	switch v35 - int32(1) {
	case 0:
		goto L33
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	case 5:
		goto L28
	default:
		goto L22
	}
L22:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1085 != 0 {
		goto L60
	} else {
		goto L61
	}
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v64))) = uint16(v1023)
	goto L22
L24:
	;
	v1012 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	v1013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v1015 = v1012*v1013 + v984
	v1016 = int32(_a_F_mul_var_0)
	v1017 = base.I32_div_u_s(v1015, v1016)
	v1020 = v1015 - v1017*v1016
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)) = uint16(v1020)
	v1023 = v1017
	goto L23
L25:
	;
	v967 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+2)))
	v968 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v971 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	v972 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v974 = v967*v968 + v940 + v971*v972
	v975 = int32(_a_F_mul_var_0)
	v976 = base.I32_div_u_s(v974, v975)
	v979 = v974 - v976*v975
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+6)) = uint16(v979)
	v984 = v976
	goto L24
L26:
	;
	v918 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	v919 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+2)))
	v923 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v926 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v929 = v918*v919 + v891 + v922*v923 + v926*v927
	v930 = int32(_a_F_mul_var_0)
	v931 = base.I32_div_u_s(v929, v930)
	v934 = v929 - v931*v930
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+8)) = uint16(v934)
	v940 = v931
	goto L25
L27:
	;
	v865 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
	v866 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v869 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	v870 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v873 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+2)))
	v874 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v877 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	v878 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v880 = v865*v866 + v838 + v869*v870 + v873*v874 + v877*v878
	v881 = int32(_a_F_mul_var_0)
	v882 = base.I32_div_u_s(v880, v881)
	v885 = v880 - v882*v881
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+10)) = uint16(v885)
	v891 = v882
	goto L26
L28:
	;
	v582 = int32(1)
	v583 = v39 - v582
	v586 = v41 + v583<<(uint(v582)%32)
	v587 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586))))
	v588 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v589 = v587 * v588
	v590 = int32(_a_F_mul_var_0)
	v591 = base.I32_div_u_s(v589, v590)
	v594 = v589 - v591*v590
	*(*uint16)(unsafe.Add(mBase, uint32(v59+v56))) = uint16(v594)
	v596 = v64 + v56
	v597 = int32(4)
	v599 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586))))
	v600 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v605 = v41 + v39<<(uint(v582)%32)
	v607 = v605 - v597
	v608 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607))))
	v609 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v611 = v599*v600 + v591 + v608*v609
	v613 = base.I32_div_u_s(v611, v590)
	v616 = v611 - v613*v590
	*(*uint16)(unsafe.Add(mBase, uint32(v596-v597))) = uint16(v616)
	v618 = int32(6)
	v620 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586))))
	v621 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607))))
	v625 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v629 = v605 - v618
	v630 = int32(*(*int16)(unsafe.Add(mBase, uint32(v629))))
	v631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v633 = v620*v621 + v613 + v624*v625 + v630*v631
	v635 = base.I32_div_u_s(v633, v590)
	v638 = v633 - v635*v590
	*(*uint16)(unsafe.Add(mBase, uint32(v596-v618))) = uint16(v638)
	v640 = int32(8)
	v642 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586))))
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v646 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607))))
	v647 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v650 = int32(*(*int16)(unsafe.Add(mBase, uint32(v629))))
	v651 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v655 = v605 - v640
	v656 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v657 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v659 = v642*v643 + v635 + v646*v647 + v650*v651 + v656*v657
	v661 = base.I32_div_u_s(v659, v590)
	v664 = v659 - v661*v590
	*(*uint16)(unsafe.Add(mBase, uint32(v596-v640))) = uint16(v664)
	v666 = int32(10)
	v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586))))
	v669 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607))))
	v673 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v676 = int32(*(*int16)(unsafe.Add(mBase, uint32(v629))))
	v677 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v680 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v681 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v686 = int32(*(*int16)(unsafe.Add(mBase, uint32(v605-v666))))
	v687 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v689 = v668*v669 + v661 + v672*v673 + v676*v677 + v680*v681 + v686*v687
	v691 = base.I32_div_u_s(v689, v590)
	v694 = v689 - v691*v590
	*(*uint16)(unsafe.Add(mBase, uint32(v596-v666))) = uint16(v694)
	if int32(5) <= v583 {
		goto L54
	} else {
		goto L55
	}
L29:
	;
	v422 = int32(1)
	v423 = v39 - v422
	v426 = v41 + v423<<(uint(v422)%32)
	v427 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426))))
	v428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v429 = v427 * v428
	v430 = int32(_a_F_mul_var_0)
	v431 = base.I32_div_u_s(v429, v430)
	v434 = v429 - v431*v430
	*(*uint16)(unsafe.Add(mBase, uint32(v59+v56))) = uint16(v434)
	v436 = v64 + v56
	v437 = int32(4)
	v439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426))))
	v440 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v445 = v41 + v39<<(uint(v422)%32)
	v447 = v445 - v437
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447))))
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v451 = v439*v440 + v431 + v448*v449
	v453 = base.I32_div_u_s(v451, v430)
	v456 = v451 - v453*v430
	*(*uint16)(unsafe.Add(mBase, uint32(v436-v437))) = uint16(v456)
	v458 = int32(6)
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426))))
	v461 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v464 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447))))
	v465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v469 = v445 - v458
	v470 = int32(*(*int16)(unsafe.Add(mBase, uint32(v469))))
	v471 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v473 = v460*v461 + v453 + v464*v465 + v470*v471
	v475 = base.I32_div_u_s(v473, v430)
	v478 = v473 - v475*v430
	*(*uint16)(unsafe.Add(mBase, uint32(v436-v458))) = uint16(v478)
	v480 = int32(8)
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426))))
	v483 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v486 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447))))
	v487 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v469))))
	v491 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v496 = int32(*(*int16)(unsafe.Add(mBase, uint32(v445-v480))))
	v497 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v499 = v482*v483 + v475 + v486*v487 + v490*v491 + v496*v497
	v501 = base.I32_div_u_s(v499, v430)
	v504 = v499 - v501*v430
	*(*uint16)(unsafe.Add(mBase, uint32(v436-v480))) = uint16(v504)
	if v423 < v437 {
		v838 = v501
		goto L27
	} else {
		goto L50
	}
L30:
	;
	v294 = int32(1)
	v295 = v39 - v294
	v298 = v41 + v295<<(uint(v294)%32)
	v299 = int32(*(*int16)(unsafe.Add(mBase, uint32(v298))))
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v301 = v299 * v300
	v302 = int32(_a_F_mul_var_0)
	v303 = base.I32_div_u_s(v301, v302)
	v306 = v301 - v303*v302
	*(*uint16)(unsafe.Add(mBase, uint32(v59+v56))) = uint16(v306)
	v308 = v64 + v56
	v309 = int32(4)
	v311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v298))))
	v312 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v317 = v41 + v39<<(uint(v294)%32)
	v319 = v317 - v309
	v320 = int32(*(*int16)(unsafe.Add(mBase, uint32(v319))))
	v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v323 = v311*v312 + v303 + v320*v321
	v325 = base.I32_div_u_s(v323, v302)
	v328 = v323 - v325*v302
	*(*uint16)(unsafe.Add(mBase, uint32(v308-v309))) = uint16(v328)
	v330 = int32(6)
	v332 = int32(*(*int16)(unsafe.Add(mBase, uint32(v298))))
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v319))))
	v337 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v342 = int32(*(*int16)(unsafe.Add(mBase, uint32(v317-v330))))
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v345 = v332*v333 + v325 + v336*v337 + v342*v343
	v347 = base.I32_div_u_s(v345, v302)
	v350 = v345 - v347*v302
	*(*uint16)(unsafe.Add(mBase, uint32(v308-v330))) = uint16(v350)
	if v295 < int32(3) {
		v891 = v347
		goto L26
	} else {
		goto L46
	}
L31:
	;
	v194 = int32(1)
	v195 = v39 - v194
	v198 = v41 + v195<<(uint(v194)%32)
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v198))))
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v201 = v199 * v200
	v202 = int32(_a_F_mul_var_0)
	v203 = base.I32_div_u_s(v201, v202)
	v206 = v201 - v203*v202
	*(*uint16)(unsafe.Add(mBase, uint32(v59+v56))) = uint16(v206)
	v209 = int32(4)
	v211 = int32(*(*int16)(unsafe.Add(mBase, uint32(v198))))
	v212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v39<<(uint(v194)%32)-v209))))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v223 = v211*v212 + v203 + v220*v221
	v225 = base.I32_div_u_s(v223, v202)
	v228 = v223 - v225*v202
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v56-v209))) = uint16(v228)
	if v195 < int32(2) {
		v940 = v225
		goto L25
	} else {
		goto L42
	}
L32:
	;
	v122 = int32(1)
	v123 = v39 - v122
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v123<<(uint(v122)%32)))))
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v129 = v127 * v128
	v130 = int32(_a_F_mul_var_0)
	v131 = base.I32_div_u_s(v129, v130)
	v134 = v129 - v131*v130
	*(*uint16)(unsafe.Add(mBase, uint32(v59+v56))) = uint16(v134)
	if v123 <= int32(0) {
		v984 = v131
		goto L24
	} else {
		goto L38
	}
L33:
	;
	v67 = int32(0)
	v69 = v39 - int32(1)
	if v69 < v67 {
		v1023 = v67
		goto L23
	} else {
		goto L34
	}
L34:
	;
	__phi73 = v67
	__phi75 = v69
	__phi76 = v39
	v73 = __phi73
	v75 = __phi75
	v76 = __phi76
	goto L35
L35:
	;
	v103 = int32(1)
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v75<<(uint(v103)%32)))))
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v112 = v109*v110 + v73
	v113 = int32(_a_F_mul_var_0)
	v114 = base.I32_div_u_s(v112, v113)
	v117 = v112 - v114*v113
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v76<<(uint(v103)%32)))) = uint16(v117)
	if v75 != 0 {
		__phi73 = v114
		__phi75 = v75 - v103
		__phi76 = v75
		v73 = __phi73
		v75 = __phi75
		v76 = __phi76
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v1023 = v114
	goto L23
L37:
	;
	goto L36
L38:
	;
	v139 = v123
	v141 = v131
	goto L39
L39:
	;
	v169 = int32(1)
	v170 = v139 << (uint(v169) % 32)
	v172 = v170 + v41
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v172))))
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v172-int32(2)))))
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v182 = v141 + v173*v174 + v179*v180
	v183 = int32(_a_F_mul_var_0)
	v184 = base.I32_div_u_s(v182, v183)
	v187 = v182 - v184*v183
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v170)+2)) = uint16(v187)
	if base.Ui32(v169) < base.Ui32(v139) {
		v139 = v139 - v169
		v141 = v184
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v984 = v184
	goto L24
L41:
	;
	goto L40
L42:
	;
	v233 = v195
	v236 = v225
	goto L43
L43:
	;
	v263 = int32(1)
	v264 = v233 << (uint(v263) % 32)
	v266 = v264 + v41
	v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(v266))))
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v271 = int32(2)
	v273 = int32(*(*int16)(unsafe.Add(mBase, uint32(v266-v271))))
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v279 = int32(*(*int16)(unsafe.Add(mBase, uint32(v266-int32(4)))))
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v282 = v267*v268 + v236 + v273*v274 + v279*v280
	v283 = int32(_a_F_mul_var_0)
	v284 = base.I32_div_u_s(v282, v283)
	v287 = v282 - v284*v283
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v264)+2)) = uint16(v287)
	if base.Ui32(v271) < base.Ui32(v233) {
		v233 = v233 - v263
		v236 = v284
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v940 = v284
	goto L25
L45:
	;
	goto L44
L46:
	;
	v355 = v295
	v358 = v347
	goto L47
L47:
	;
	v385 = int32(1)
	v386 = v355 << (uint(v385) % 32)
	v388 = v386 + v41
	v389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388))))
	v390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388-int32(2)))))
	v396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v401 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388-int32(4)))))
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v388-int32(6)))))
	v408 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v410 = v389*v390 + v358 + v395*v396 + v401*v402 + v407*v408
	v411 = int32(_a_F_mul_var_0)
	v412 = base.I32_div_u_s(v410, v411)
	v415 = v410 - v412*v411
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v386)+2)) = uint16(v415)
	if base.Ui32(int32(3)) < base.Ui32(v355) {
		v355 = v355 - v385
		v358 = v412
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v891 = v412
	goto L26
L49:
	;
	goto L48
L50:
	;
	v511 = v423
	v512 = v501
	goto L51
L51:
	;
	v539 = int32(1)
	v540 = v511 << (uint(v539) % 32)
	v542 = v540 + v41
	v543 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542))))
	v544 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v549 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542-int32(2)))))
	v550 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v553 = int32(4)
	v555 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542-v553))))
	v556 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v561 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542-int32(6)))))
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v567 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542-int32(8)))))
	v568 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v570 = v543*v544 + v512 + v549*v550 + v555*v556 + v561*v562 + v567*v568
	v571 = int32(_a_F_mul_var_0)
	v572 = base.I32_div_u_s(v570, v571)
	v575 = v570 - v572*v571
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v540)+2)) = uint16(v575)
	if base.Ui32(v553) < base.Ui32(v511) {
		v511 = v511 - v539
		v512 = v572
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v838 = v572
	goto L27
L53:
	;
	goto L52
L54:
	;
	v701 = v583
	v702 = v691
	goto L57
L55:
	;
	v781 = v691
	goto L56
L56:
	;
	v808 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+8)))
	v809 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v812 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
	v813 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v816 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	v817 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v820 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+2)))
	v821 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v824 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
	v825 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v827 = v808*v809 + v781 + v812*v813 + v816*v817 + v820*v821 + v824*v825
	v828 = int32(_a_F_mul_var_0)
	v829 = base.I32_div_u_s(v827, v828)
	v832 = v827 - v829*v828
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)) = uint16(v832)
	v838 = v829
	goto L27
L57:
	;
	v729 = int32(1)
	v730 = v701 << (uint(v729) % 32)
	v732 = v730 + v41
	v733 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732))))
	v734 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	v739 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732-int32(2)))))
	v740 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+2)))
	v745 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732-int32(4)))))
	v746 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v751 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732-int32(6)))))
	v752 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v757 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732-int32(8)))))
	v758 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v763 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732-int32(10)))))
	v764 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	v766 = v733*v734 + v702 + v739*v740 + v745*v746 + v751*v752 + v757*v758 + v763*v764
	v767 = int32(_a_F_mul_var_0)
	v768 = base.I32_div_u_s(v766, v767)
	v771 = v766 - v768*v767
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v730)+2)) = uint16(v771)
	if base.Ui32(int32(5)) < base.Ui32(v701) {
		v701 = v701 - v729
		v702 = v768
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v781 = v768
	goto L56
L59:
	;
	goto L58
L60:
	;
	F_pfree(m, v1085)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L20
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
	if v53 != v52 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v1094 = int32(_a_F_mul_var_1)
	goto L66
L65:
	;
	v1094 = int32(0)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1094
	v1098 = v50 + v51 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1098
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v1100 + v1101
	if int32(0) < v54 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1230
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v1223
	return
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v1223 = v1189
	v1230 = int32(0)
	goto L67
L69:
	;
	v1107 = v64
	v1111 = v1098
	v1114 = v54
	goto L72
L70:
	;
	goto L71
L71:
	;
	if v54 != 0 {
		v1223 = v64
		v1230 = v54
		goto L67
	} else {
		goto L82
	}
L72:
	;
	v1138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107))))
	if v1138 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v1189 = v64 + v56
	goto L68
L74:
	;
	v1146 = v1114
	goto L77
L75:
	;
	goto L76
L76:
	;
	v1180 = int32(1)
	v1181 = v1111 - v1180
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1181
	if v1180 < v1114 {
		v1107 = v1107 + int32(2)
		v1111 = v1181
		v1114 = v1114 - v1180
		goto L72
	} else {
		goto L81
	}
L77:
	;
	v1175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107+v1146<<(uint(int32(1))%32)-int32(2)))))
	if v1175 != 0 {
		v1223 = v1107
		v1230 = v1146
		goto L67
	} else {
		goto L79
	}
L79:
	;
	v1176 = int32(1)
	if v1176 < v1146 {
		v1146 = v1146 - v1176
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v1189 = v1107
	goto L68
L81:
	;
	goto L73
L82:
	;
	v1189 = v64
	goto L68
L83:
	;
	v1288 = v1260
	goto L85
L84:
	;
	v1288 = v1286
	goto L85
L85:
	;
	v1289 = int32(1)
	v1291 = int32(2)
	v1292 = base.I32_div_s(v39+v1289, v1291)
	v1296 = base.I32_div_s(v35+v1289, v1291)
	v1297 = v1292 + v1296
	v1298 = v1260 - v1297
	v1300 = v1298 + v1289
	if v1288 <= v1300 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	goto L1
L87:
	;
	goto L88
L88:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v1304 = v1288 - v1300
	v1307 = v1288 << (uint(int32(3)) % 32)
	if v1292 < v1304 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v1309 = v1292
	goto L91
L90:
	;
	v1309 = v1304
	goto L91
L91:
	;
	v1313 = F_palloc(m, v1307+v1309<<(uint(int32(2))%32))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L20
	} else {
		goto L92
	}
L92:
	;
	v1315 = v1313 + v1307
	v1316 = int32(0)
	if v1296 < v1304 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v1317 = v1296
	goto L95
L94:
	;
	v1317 = v1304
	goto L95
L95:
	;
	v1319 = v1309 - int32(1)
	if int32(0) < v1319 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v1309 != int32(2) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v1466 = v1316
	goto L98
L98:
	;
	v1491 = v1466 << (uint(int32(2)) % 32)
	v1493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v1491))))
	v1495 = v1493 * int32(_a_F_mul_var_0)
	v1497 = int32(1)
	v1500 = v1466<<(uint(v1497)%32) | v1497
	if v1500 < v39 {
		goto L107
	} else {
		goto L108
	}
L99:
	;
	v1466 = v1319
	goto L98
L100:
	;
	v1329 = int32(0)
	v1336 = v1316
	goto L103
L101:
	;
	v1394 = v1316
	goto L102
L102:
	;
	v1419 = v1394 << (uint(int32(2)) % 32)
	v1421 = v1419 + v41
	v1422 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1421))))
	v1425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1421)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1315+v1419))) = v1422*int32(_a_F_mul_var_0) + v1425
	goto L99
L103:
	;
	v1360 = int32(2)
	v1361 = v1336 << (uint(v1360) % 32)
	v1363 = v41 + v1361
	v1364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1363))))
	v1365 = int32(_a_F_mul_var_0)
	v1367 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1363)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1315+v1361))) = v1364*v1365 + v1367
	v1371 = v1361 | int32(4)
	v1373 = v41 + v1371
	v1374 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1373))))
	v1377 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1373)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1315+v1371))) = v1374*v1365 + v1377
	v1381 = v1336 + v1360
	v1383 = v1329 + v1360
	if v1383 != v1319&int32(2147483646) {
		v1329 = v1383
		v1336 = v1381
		goto L103
	} else {
		goto L105
	}
L104:
	;
	if v1319&int32(1) == int32(0) {
		goto L99
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v1394 = v1381
	goto L102
L107:
	;
	v1505 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v1500<<(uint(int32(1))%32)))))
	v1507 = v1495 + v1505
	goto L109
L108:
	;
	v1507 = v1495
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1491+v1315))) = v1507
	v1514 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1317<<(uint(int32(2))%32)+v43-int32(4)))))
	v1516 = v1514 * int64(10000)
	v1517 = int32(1)
	v1520 = v1317<<(uint(v1517)%32) - v1517
	if v1520 < v35 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v1525 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43+v1520<<(uint(int32(1))%32)))))
	v1527 = v1516 + v1525
	goto L112
L111:
	;
	v1527 = v1516
	goto L112
L112:
	;
	v1528 = v1317 + v1298
	v1530 = v1528 << (uint(int32(3)) % 32)
	if v1530 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	base.MemoryFill(m, v1313, int32(0), v1530)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v1534 = v1527 & int64(4294967295)
	v1535 = v1288 - v1528
	if v1309 < v1535 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v1738 = v1317 - int32(2)
	if int32(0) <= v1738 {
		goto L131
	} else {
		goto L132
	}
L117:
	;
	v1537 = v1309
	goto L119
L118:
	;
	v1537 = v1535
	goto L119
L119:
	;
	if v1537 <= int32(0) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v1540 = v1530 + v1313
	v1542 = v1537 & int32(3)
	v1543 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1537) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v1550 = int32(0)
	v1554 = v1543
	goto L124
L122:
	;
	v1634 = v1543
	goto L123
L123:
	;
	v1662 = v1543
	v1665 = v1634
	goto L128
L124:
	;
	v1581 = int32(3)
	v1584 = int32(2)
	v1587 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v1554<<(uint(v1584)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1540+v1554<<(uint(v1581)%32)))) = v1534 * v1587
	v1591 = v1554 | int32(1)
	v1598 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v1591<<(uint(v1584)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1540+v1591<<(uint(v1581)%32)))) = v1534 * v1598
	v1602 = v1554 | v1584
	v1609 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v1602<<(uint(v1584)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1540+v1602<<(uint(v1581)%32)))) = v1534 * v1609
	v1613 = v1554 | v1581
	v1620 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v1613<<(uint(v1584)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1540+v1613<<(uint(v1581)%32)))) = v1534 * v1620
	v1623 = int32(4)
	v1624 = v1554 + v1623
	v1626 = v1550 + v1623
	if v1626 != v1537&int32(2147483644) {
		v1550 = v1626
		v1554 = v1624
		goto L124
	} else {
		goto L126
	}
L125:
	;
	if v1542 == int32(0) {
		goto L116
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v1634 = v1624
	goto L123
L128:
	;
	v1698 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v1665<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1540+v1665<<(uint(int32(3))%32)))) = v1534 * v1698
	v1701 = int32(1)
	v1704 = v1662 + v1701
	if v1704 != v1542 {
		v1662 = v1704
		v1665 = v1665 + v1701
		goto L128
	} else {
		goto L130
	}
L129:
	;
	goto L116
L130:
	;
	goto L129
L131:
	;
	v1758 = v1738
	v1760 = v1288 + v1297 - (v1317 + v1258)
	v1781 = v1534
	goto L134
L132:
	;
	goto L133
L133:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v2178 != 0 {
		goto L178
	} else {
		goto L179
	}
L134:
	;
	v1787 = v43 + v1758<<(uint(int32(2))%32)
	v1788 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1787))))
	v1791 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1787)+2)))
	v1792 = v1788*int32(_a_F_mul_var_0) + v1791
	if v1792 == int32(0) {
		v2137 = v1781
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	v2141 = int32(1)
	if int32(0) < v1758 {
		v1758 = v1758 - v2141
		v1760 = v1760 + v2141
		v1781 = v2137
		goto L134
	} else {
		goto L177
	}
L137:
	;
	v1795 = base.I64_extend_i32_u(v1792)
	v1796 = v1781 + v1795
	if base.Ui64(int64(184467440738)) <= base.Ui64(v1796) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v1288 <= int32(0) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v1982 = v1796
	goto L140
L140:
	;
	v1986 = v1288 + (v1298 ^ int32(-1)) - v1758
	if v1309 < v1986 {
		goto L160
	} else {
		goto L161
	}
L141:
	;
	v1982 = base.I64_extend_i32_u(v1792 + int32(1))
	goto L140
L142:
	;
	v1802 = int64(0)
	if v1288 != int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1809 = v1288
	v1812 = int32(0)
	v1833 = v1802
	goto L146
L144:
	;
	v1881 = v1288
	v1905 = v1802
	goto L145
L145:
	;
	v1912 = v1313 + v1881<<(uint(int32(3))%32) - int32(8)
	v1913 = *(*int64)(unsafe.Add(mBase, uint32(v1912)))
	v1914 = v1913 + v1905
	v1916 = base.I64_rem_u_s(v1914, int64(100000000))
	if base.Ui64(int64(99999999)) < base.Ui64(v1914) {
		goto L157
	} else {
		goto L158
	}
L146:
	;
	v1840 = v1313 + v1809<<(uint(int32(3))%32) - int32(8)
	v1841 = *(*int64)(unsafe.Add(mBase, uint32(v1840)))
	v1842 = v1841 + v1833
	if base.Ui64(v1842) < base.Ui64(int64(100000000)) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	if v1288&int32(1) == int32(0) {
		goto L141
	} else {
		goto L156
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1840))) = v1851
	v1856 = v1809 - int32(2)
	v1859 = v1313 + v1856<<(uint(int32(3))%32)
	v1860 = *(*int64)(unsafe.Add(mBase, uint32(v1859)))
	v1861 = v1860 + v1852
	if base.Ui64(int64(100000000)) <= base.Ui64(v1861) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v1851 = v1842
	v1852 = int64(0)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v1847 = base.I64_div_u_s(v1842, int64(100000000))
	v1851 = v1847*int64(-100000000) + v1842
	v1852 = v1847
	goto L148
L152:
	;
	v1865 = base.I64_div_u_s(v1861, int64(100000000))
	v1869 = v1865*int64(-100000000) + v1861
	v1870 = v1865
	goto L154
L153:
	;
	v1869 = v1861
	v1870 = int64(0)
	goto L154
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1859))) = v1869
	v1873 = v1812 + int32(2)
	if v1873 != v1288&int32(2147483646) {
		v1809 = v1856
		v1812 = v1873
		v1833 = v1870
		goto L146
	} else {
		goto L155
	}
L155:
	;
	goto L147
L156:
	;
	v1881 = v1856
	v1905 = v1870
	goto L145
L157:
	;
	v1919 = v1916
	goto L159
L158:
	;
	v1919 = v1914
	goto L159
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1912))) = v1919
	goto L141
L160:
	;
	v1988 = v1309
	goto L162
L161:
	;
	v1988 = v1986
	goto L162
L162:
	;
	if v1988 <= int32(0) {
		v2137 = v1982
		goto L136
	} else {
		goto L163
	}
L163:
	;
	v1993 = v1313 + v1300<<(uint(int32(3))%32) + v1758<<(uint(int32(3))%32)
	v1994 = int32(0)
	if v1292 < v1760 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1996 = v1292
	goto L166
L165:
	;
	v1996 = v1760
	goto L166
L166:
	;
	if v1996 < v1304 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1998 = v1996
	goto L169
L168:
	;
	v1998 = v1304
	goto L169
L169:
	;
	if v1998 != int32(1) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v2007 = int32(0)
	v2010 = v1994
	goto L173
L171:
	;
	v2072 = v1994
	goto L172
L172:
	;
	v2101 = v1993 + v2072<<(uint(int32(3))%32)
	v2102 = *(*int64)(unsafe.Add(mBase, uint32(v2101)))
	v2106 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v2072<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2101))) = v2102 + v2106*v1795
	v2137 = v1982
	goto L136
L173:
	;
	v2037 = int32(3)
	v2039 = v1993 + v2010<<(uint(v2037)%32)
	v2040 = *(*int64)(unsafe.Add(mBase, uint32(v2039)))
	v2041 = int32(2)
	v2044 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v2010<<(uint(v2041)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2039))) = v2040 + v2044*v1795
	v2049 = v2010 | int32(1)
	v2052 = v1993 + v2049<<(uint(v2037)%32)
	v2053 = *(*int64)(unsafe.Add(mBase, uint32(v2052)))
	v2057 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1315+v2049<<(uint(v2041)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2052))) = v2053 + v2057*v1795
	v2062 = v2010 + v2041
	v2064 = v2007 + v2041
	if v2064 != v1998&int32(-2) {
		v2007 = v2064
		v2010 = v2062
		goto L173
	} else {
		goto L175
	}
L174:
	;
	if v1998&int32(1) == int32(0) {
		v2137 = v1982
		goto L136
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v2072 = v2062
	goto L172
L177:
	;
	goto L135
L178:
	;
	F_pfree(m, v2178)
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L20
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v1302 != v1303 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	goto L180
L182:
	;
	v2186 = int32(_a_F_mul_var_1)
	goto L184
L183:
	;
	v2186 = int32(0)
	goto L184
L184:
	;
	v2187 = int32(2)
	v2191 = F_palloc(m, v1288<<(uint(v2187)%32)|v2187)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v2191
	v2194 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2191))) = uint16(v2194)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1288 << (uint(int32(1)) % 32)
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2199 = v2197 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2199
	if v2194 < v1288 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v2212 = v1288
	v2231 = int64(0)
	goto L189
L187:
	;
	goto L188
L188:
	;
	F_pfree(m, v1313)
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L20
	} else {
		goto L195
	}
L189:
	;
	v2236 = v2212 - int32(1)
	v2240 = *(*int64)(unsafe.Add(mBase, uint32(v1313+v2236<<(uint(int32(3))%32))))
	v2241 = v2240 + v2231
	v2245 = v2199 + v2236<<(uint(int32(2))%32)
	if base.Ui64(int64(100000000)) <= base.Ui64(v2241) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L188
L191:
	;
	v2249 = base.I64_div_u_s(v2241, int64(100000000))
	v2253 = v2249
	v2254 = v2249*int64(-100000000) + v2241
	goto L193
L192:
	;
	v2253 = int64(0)
	v2254 = v2241
	goto L193
L193:
	;
	v2255 = base.I32_wrap_i64(v2254)
	v2256 = int32(_a_F_mul_var_0)
	v2257 = base.I32_div_u_s(v2255, v2256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2245))) = uint16(v2257)
	v2261 = v2255 - v2257*v2256
	*(*uint16)(unsafe.Add(mBase, uint32(v2245)+2)) = uint16(v2261)
	if base.Ui32(int32(1)) < base.Ui32(v2212) {
		v2212 = v2236
		v2231 = v2253
		goto L189
	} else {
		goto L194
	}
L194:
	;
	goto L190
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2186
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2309 = l3 + v1275<<(uint(int32(2))%32)
	if v2309+int32(4) < int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v2425 {
		goto L224
	} else {
		goto L225
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	goto L196
L198:
	;
	goto L199
L199:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2320 = l3 & int32(3)
	v2324 = base.I32_div_s(v2309+int32(7), int32(4))
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2325 <= v2324 {
		goto L204
	} else {
		goto L205
	}
L200:
	;
	goto L196
L201:
	;
	if int32(0) <= v2390 {
		goto L200
	} else {
		goto L221
	}
L202:
	;
	v2370 = v2364
	goto L215
L203:
	;
	v2339 = int32(1)
	v2340 = v2324 - v2339
	v2343 = v2318 + v2340<<(uint(v2339)%32)
	v2344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2343))))
	v2345 = int32(2)
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2320<<(uint(v2345)%32))+uint32(_c_F_mul_var[0])))
	v2348 = base.I32_rem_s(v2344, v2347)
	v2349 = v2344 - v2348
	*(*uint16)(unsafe.Add(mBase, uint32(v2343))) = uint16(v2349)
	v2352 = base.I32_div_s(v2347, v2345)
	if v2348 < v2352 {
		v2390 = v2340
		goto L201
	} else {
		goto L210
	}
L204:
	;
	if base.B2i32(v2320 == int32(0))|base.B2i32(v2324 != v2325) != 0 {
		goto L200
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2324
	if v2320 != 0 {
		goto L203
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2324
	goto L203
L208:
	;
	v2336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2318+v2324<<(uint(int32(1))%32)))))
	if v2336 <= int32(_a_F_mul_var_2) {
		v2390 = v2324
		goto L201
	} else {
		goto L209
	}
L209:
	;
	v2364 = v2324
	goto L202
L210:
	;
	v2355 = v2347 + base.I32_extend16_s(v2349)
	if int32(_a_F_mul_var_3) < v2355 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v2360 = v2355 + int32(_a_F_mul_var_4)
	goto L213
L212:
	;
	v2360 = v2355
	goto L213
L213:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2343))) = uint16(v2360)
	if v2355 < int32(_a_F_mul_var_0) {
		v2390 = v2340
		goto L201
	} else {
		goto L214
	}
L214:
	;
	v2364 = v2340
	goto L202
L215:
	;
	v2376 = int32(1)
	v2377 = v2370 - v2376
	v2380 = v2318 + v2377<<(uint(v2376)%32)
	v2383 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2380))))
	v2385 = base.B2i32(int32(_a_F_mul_var_5) < v2383)
	if int32(_a_F_mul_var_5) < v2383 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v2390 = v2377
	goto L201
L217:
	;
	v2386 = int32(-9999)
	goto L219
L218:
	;
	v2386 = v2376
	goto L219
L219:
	;
	v2387 = v2386 + v2383
	*(*uint16)(unsafe.Add(mBase, uint32(v2380))) = uint16(v2387)
	if int32(_a_F_mul_var_5) < v2383 {
		v2370 = v2377
		goto L215
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2398 - int32(2)
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2403 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2402 + v2403
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2406 + v2403
	goto L200
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2552
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2548
	return
L223:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v2548 = v2514
	v2552 = int32(0)
	goto L222
L224:
	;
	v2431 = v2424
	v2435 = v2425
	goto L227
L225:
	;
	goto L226
L226:
	;
	if v2425 != 0 {
		v2548 = v2424
		v2552 = v2425
		goto L222
	} else {
		goto L237
	}
L227:
	;
	v2462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2431))))
	if v2462 != 0 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v2514 = v2424 + v2425<<(uint(int32(1))%32)
	goto L223
L229:
	;
	v2467 = v2435
	goto L232
L230:
	;
	goto L231
L231:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2505 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2504 - v2505
	if v2505 < v2435 {
		v2431 = v2431 + int32(2)
		v2435 = v2435 - v2505
		goto L227
	} else {
		goto L236
	}
L232:
	;
	v2499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2431+v2467<<(uint(int32(1))%32)-int32(2)))))
	if v2499 != 0 {
		v2548 = v2431
		v2552 = v2467
		goto L222
	} else {
		goto L234
	}
L234:
	;
	v2500 = int32(1)
	if v2500 < v2467 {
		v2467 = v2467 - v2500
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v2514 = v2431
	goto L223
L236:
	;
	goto L228
L237:
	;
	v2514 = v2424
	goto L223
L238:
	;
	F_pfree(m, v2596)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L20
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v2602 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v2602
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v2602
	return
L241:
	;
	goto L240
}
