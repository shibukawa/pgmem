package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_addTargetToGroupList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
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
	var v277 int32
	_ = v277
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_exprType(m, v19)
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
	if v20 == int32(705) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = int32(-1)
	v34 = F_coerce_type(m, l0, v27, int32(705), v26, v30, int32(0), int32(2), v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v37 = v20
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v39 = int32(0)
	if base.B2i32(v38 == v39)|base.B2i32(l2 == v39) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
	v37 = v26
	goto L5
L7:
	;
	m.G0 = v17 + int32(32)
	return v277
L8:
	;
	v86 = F_palloc0(m, int32(20))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v44 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v53 = int32(0)
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v47+v53<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v66 == v38 {
		v277 = l2
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v69 = v53 + int32(1)
	if v69 != v44 {
		v53 = v69
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(106)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v17
	v95 = int32(_a_F_addTargetToGroupList_0)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_addTargetToGroupList[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v96
	*(*int32)(unsafe.Add(mBase, _c_F_addTargetToGroupList[0])) = v17 + int32(8)
	goto L16
L16:
	;
	v102 = int32(0)
	F_get_sort_group_operators(m, v37, v102, int32(1), v102, v17+int32(28), v17+int32(24), v102, v17+int32(23))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_addTargetToGroupList[0])) = v115
	goto L18
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v117 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if l3 == int32(0) {
		v248 = int32(1)
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v255 = v117
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v255
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v268 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)) = uint16(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v267
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+23)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+18)) = uint8(v271)
	v273 = F_lappend(m, l2, v86)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L51
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v248
	v255 = v248
	goto L21
L23:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v124 <= int32(0) {
		v248 = int32(1)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v128 = v124 & int32(3)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v130 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v124) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v248 = v223 + int32(1)
	goto L22
L26:
	;
	v138 = v130
	v143 = v130
	v146 = int32(0)
	goto L29
L27:
	;
	v178 = v130
	v183 = v130
	goto L28
L28:
	;
	v192 = v178
	v197 = v183
	v201 = v130
	goto L45
L29:
	;
	v154 = v129 + v138<<(uint(int32(2))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	if base.Ui32(v143) < base.Ui32(v162) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v128 == int32(0) {
		v223 = v170
		goto L25
	} else {
		goto L44
	}
L31:
	;
	v164 = v162
	goto L33
L32:
	;
	v164 = v143
	goto L33
L33:
	;
	if base.Ui32(v164) < base.Ui32(v160) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v166 = v160
	goto L36
L35:
	;
	v166 = v164
	goto L36
L36:
	;
	if base.Ui32(v166) < base.Ui32(v158) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v168 = v158
	goto L39
L38:
	;
	v168 = v166
	goto L39
L39:
	;
	if base.Ui32(v168) < base.Ui32(v156) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v170 = v156
	goto L42
L41:
	;
	v170 = v168
	goto L42
L42:
	;
	v171 = int32(4)
	v172 = v138 + v171
	v174 = v146 + v171
	if v174 != v124&int32(2147483644) {
		v138 = v172
		v143 = v170
		v146 = v174
		goto L29
	} else {
		goto L43
	}
L43:
	;
	goto L30
L44:
	;
	v178 = v172
	v183 = v170
	goto L28
L45:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v129+v192<<(uint(int32(2))%32))))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	if base.Ui32(v197) < base.Ui32(v210) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v223 = v212
	goto L25
L47:
	;
	v212 = v210
	goto L49
L48:
	;
	v212 = v197
	goto L49
L49:
	;
	v213 = int32(1)
	v216 = v201 + v213
	if v216 != v128 {
		v192 = v192 + v213
		v197 = v212
		v201 = v216
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v277 = v273
	goto L7
}
func F_transformUpdateTargetList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v12 = m.G0
	v13 = int32(16)
	v14 = v12 - v13
	m.G0 = v14
	v17 = F_transformTargetList(m, l0, l1, v13)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+120)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v24 <= v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v23 + int32(1)
	goto L5
L4:
	;
	goto L5
L5:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = v29
	goto L8
L7:
	;
	v31 = int32(0)
	goto L8
L8:
	;
	if v17 == int32(0) {
		v231 = v31
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L67
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L64
	}
L11:
	;
	if v231 != 0 {
		goto L9
	} else {
		goto L63
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v34 <= int32(0) {
		v231 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v44 = v31
	v46 = int32(0)
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+26)))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v231 = v219
	goto L11
L16:
	;
	v223 = v46 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v223 < v224 {
		v44 = v219
		v46 = v223
		goto L14
	} else {
		goto L62
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v56 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+8)) = uint16(v56)
	v219 = v44
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v44 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v69 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+120)))
	if v69 < v72 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v127 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	v127 = v120
	goto L21
L23:
	;
	v120 = v78 + int32(1)
	goto L22
L24:
	;
	v78 = v69
	goto L27
L25:
	;
	goto L26
L26:
	;
	goto L35
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v87 = v80 + v81<<(uint(int32(4))%32) + v78*int32(100)
	v90 = F_namestrcmp(m, v87+int32(24), v67)
	mBase = m.M
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+111)))
	if v93 != int32(1) {
		goto L23
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v97 = v78 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+120)))
	if v97 < v99 {
		v78 = v97
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
L34:
	;
	v127 = int32(0)
	goto L21
L35:
	;
	v108 = F_SystemAttributeByName(m, v67)
	mBase = m.M
	if v108 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v108)+74)))
	if v111 != 0 {
		v120 = v111
		goto L22
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v197 = F_transformAssignedExpr(m, l0, v192, int32(17), v194, v127, v195, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L41:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v138 + int32(4)
	F_errmsg(m, int32(_a_F_transformUpdateTargetList_0), v14)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v147 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	F_parser_errposition(m, l0, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L55
	}
L45:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if base.B2i32(v156 == int32(0))|base.B2i32(v156 != v159) != 0 {
		v177 = v156
		v178 = v159
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v177-v178 != 0 {
		goto L44
	} else {
		goto L53
	}
L47:
	;
	goto L46
L48:
	;
	v162 = v150
	v163 = v153
	goto L49
L49:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v167 == int32(0) {
		v177 = v167
		v178 = v166
		goto L47
	} else {
		goto L51
	}
L50:
	;
	v177 = v167
	v178 = v166
	goto L47
L51:
	;
	v170 = int32(1)
	if v167 == v166 {
		v162 = v162 + v170
		v163 = v163 + v170
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	F_errhint(m, int32(_a_F_transformUpdateTargetList_1), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	F_errfinish(m, int32(_a_F_transformUpdateTargetList_2), int32(2581), int32(_a_F_transformUpdateTargetList_3))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v194
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+8)) = uint16(v127)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v205 = F_bms_add_member(m, v202, v127+int32(7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v205
	v209 = v44 + int32(4)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v209) < base.Ui32(v211+v212<<(uint(int32(2))%32)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = v209
	goto L61
L60:
	;
	v217 = int32(0)
	goto L61
L61:
	;
	v219 = v217
	goto L16
L62:
	;
	goto L15
L63:
	;
	m.G0 = v14 + int32(16)
	return v17
L64:
	;
	F_errmsg_internal(m, int32(_a_F_transformUpdateTargetList_4), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_transformUpdateTargetList_2), int32(2567), int32(_a_F_transformUpdateTargetList_3))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_transformUpdateTargetList_4), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_transformUpdateTargetList_2), int32(2595), int32(_a_F_transformUpdateTargetList_3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
