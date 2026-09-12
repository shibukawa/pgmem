package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildTupleFromCStrings(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_palloc(m, v12<<(uint(int32(2))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_palloc(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v12 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v88 = F_heap_form_tuple(m, v11, v15, v19)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(29)+v27<<(uint(int32(4))%32)))))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v76 = v27 + int32(1)
	if v76 != v12 {
		v27 = v76
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v42 = v27 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = l1 + v42
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50+v42)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53+v42)))
	v56 = F_InputFunctionCall(m, v44+v27*int32(28), v49, v52, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v27<<(uint(int32(2))%32)))) = int32(0)
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v19))) = uint8(v71)
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v42))) = v56
	v59 = v27 + v19
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v61)
	goto L9
L15:
	;
	goto L16
L16:
	;
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
	goto L9
L17:
	;
	goto L8
L18:
	;
	F_pfree(m, v15)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v19)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	return v88
}
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_palloc(m, v10*int32(116)+int32(20))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(-4294965047)
	v24 = int32(4)
	v27 = int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = v10 * int32(100)
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if int32(0) < v39 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v37 = F__emscripten_memcpy_bulkmem(m, v15+v10<<(uint(v24)%32)+v27, l0+v29<<(uint(v24)%32)+v27, v36)
	mBase = m.M
	goto L6
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v42 = int32(31)
	v48 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	if v9 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_populate_compact_attribute(m, v15, v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v58 = v48 << (uint(int32(4)) % 32)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+(l0+v42)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v42+v58))) = uint8(v61)
	v64 = v48 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v64 < v65 {
		v48 = v64
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v76 = F_palloc0(m, int32(20))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v278
	return v15
L17:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+16)) = uint8(v78)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+17)) = uint8(v80)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+18)) = uint8(v82)
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+12)) = uint16(v84)
	if v84 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v132 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v90 = F_palloc(m, v84<<(uint(int32(3))%32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v90
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+12)))
	v96 = v94 << (uint(int32(3)) % 32)
	if v96 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+12)))
	if v99 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L22:
	;
	v97 = F__emscripten_memcpy_bulkmem(m, v90, v93, v96)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v103 = v99
	goto L26
L26:
	;
	v111 = v103 - int32(1)
	v113 = v111 << (uint(int32(3)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113+v114)+4))
	v117 = F_pstrdup(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L18
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v119+v113)+4)) = v117
	if base.Ui32(int32(1)) < base.Ui32(v103) {
		v103 = v111
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)) = uint16(v193)
	if v193 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = F_palloc(m, v135<<(uint(int32(3))%32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = v142 << (uint(int32(3)) % 32)
	if v144 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = v147 - int32(1)
	if v149 < int32(0) {
		goto L30
	} else {
		goto L37
	}
L34:
	;
	v145 = F__emscripten_memcpy_bulkmem(m, v138, v141, v144)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v155 = v149
	goto L38
L38:
	;
	v163 = v155 << (uint(int32(3)) % 32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v165 = v163 + v164
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v166 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L30
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v172 = l0 + int32(20) + v155<<(uint(int32(4))%32)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+6)))
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v172)+4)))
	v175 = F_datumCopy(m, v169, v173, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if int32(0) < v155 {
		v155 = v155 - int32(1)
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v177+v163)+4)) = v175
	goto L42
L44:
	;
	goto L39
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v76
	goto L16
L46:
	;
	v199 = F_palloc(m, v193*int32(12))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)))
	v205 = v203 * int32(12)
	if v205 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)))
	if v208 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L49:
	;
	v206 = F__emscripten_memcpy_bulkmem(m, v199, v202, v205)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	v217 = v208
	goto L53
L53:
	;
	v220 = v217 - int32(1)
	v222 = v220 * int32(12)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222+v223)))
	v226 = F_pstrdup(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L45
L55:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228+v222))) = v226
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231+v222)+4))
	v234 = F_pstrdup(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v222)+4)) = v234
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v222)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v239+v222)+8)) = uint8(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247+v222)+9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v245+v222)+9)) = uint8(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v222)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v251+v222)+10)) = uint8(v255)
	if base.Ui32(int32(1)) < base.Ui32(v217) {
		v217 = v220
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
}
func F_ExecSetTupleBound(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	v5 = l1
	goto L2
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	switch v7 - int32(394) {
	case 0:
		goto L8
	default:
		goto L1
	case 3:
		goto L11
	case 4:
		goto L12
	case 17:
		goto L7
	case 32:
		goto L10
	case 33:
		goto L9
	case 38, 39:
		goto L6
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+120)) = l0
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+117)) = uint8(v206)
	goto L1
L4:
	;
	goto L3
L5:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v5+v202)))
	v5 = v204
	goto L2
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+112)) = l0
	v202 = int32(36)
	goto L5
L7:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	if v197 == int32(0) {
		v202 = int32(116)
		goto L5
	} else {
		goto L80
	}
L8:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	if v195 != 0 {
		v5 = v195
		goto L2
	} else {
		goto L79
	}
L9:
	;
	if l0 < int64(0) {
		goto L76
	} else {
		goto L77
	}
L10:
	;
	if int64(0) <= l0 {
		goto L4
	} else {
		goto L75
	}
L11:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v97 <= int32(0) {
		goto L1
	} else {
		goto L44
	}
L12:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v16 = int32(0)
	goto L14
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+104))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v16<<(uint(int32(2))%32))))
	v24 = v21
	goto L18
L15:
	;
	goto L1
L16:
	;
	v94 = v16 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v94 < v95 {
		v16 = v94
		goto L14
	} else {
		goto L43
	}
L17:
	;
	goto L16
L18:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	switch v26 - int32(394) {
	case 0:
		goto L24
	default:
		goto L17
	case 3:
		goto L27
	case 4:
		goto L28
	case 17:
		goto L23
	case 32:
		goto L26
	case 33:
		goto L25
	case 38, 39:
		goto L22
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = l0
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+117)) = uint8(v85)
	goto L17
L20:
	;
	goto L19
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v24+v81)))
	v24 = v83
	goto L18
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+112)) = l0
	v81 = int32(36)
	goto L21
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	if v76 == int32(0) {
		v81 = int32(116)
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v74 != 0 {
		v24 = v74
		goto L18
	} else {
		goto L41
	}
L25:
	;
	if l0 < int64(0) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	if int64(0) <= l0 {
		goto L20
	} else {
		goto L37
	}
L27:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v46 <= int32(0) {
		goto L17
	} else {
		goto L33
	}
L28:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v29 <= int32(0) {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v35 = int32(0)
	goto L30
L30:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v40)
	mBase = m.M
	v43 = v35 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v43 < v44 {
		v35 = v43
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L17
L32:
	;
	goto L31
L33:
	;
	v52 = int32(0)
	goto L34
L34:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v52<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v57)
	mBase = m.M
	v60 = v52 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v60 < v61 {
		v52 = v60
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L17
L36:
	;
	goto L35
L37:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+117)) = uint8(v65)
	goto L16
L38:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+116)) = uint8(v69)
	goto L16
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = l0
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+116)) = uint8(v72)
	goto L16
L41:
	;
	goto L17
L42:
	;
	goto L17
L43:
	;
	goto L15
L44:
	;
	v103 = int32(0)
	goto L45
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v5)+104))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v103<<(uint(int32(2))%32))))
	v111 = v108
	goto L49
L46:
	;
	goto L1
L47:
	;
	v181 = v103 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v181 < v182 {
		v103 = v181
		goto L45
	} else {
		goto L74
	}
L48:
	;
	goto L47
L49:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	switch v113 - int32(394) {
	case 0:
		goto L55
	default:
		goto L48
	case 3:
		goto L58
	case 4:
		goto L59
	case 17:
		goto L54
	case 32:
		goto L57
	case 33:
		goto L56
	case 38, 39:
		goto L53
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+120)) = l0
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+117)) = uint8(v172)
	goto L48
L51:
	;
	goto L50
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v111+v168)))
	v111 = v170
	goto L49
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+112)) = l0
	v168 = int32(36)
	goto L52
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	if v163 == int32(0) {
		v168 = int32(116)
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	if v161 != 0 {
		v111 = v161
		goto L49
	} else {
		goto L72
	}
L56:
	;
	if l0 < int64(0) {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	if int64(0) <= l0 {
		goto L51
	} else {
		goto L68
	}
L58:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v133 <= int32(0) {
		goto L48
	} else {
		goto L64
	}
L59:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v116 <= int32(0) {
		goto L48
	} else {
		goto L60
	}
L60:
	;
	v122 = int32(0)
	goto L61
L61:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+104))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v122<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v127)
	mBase = m.M
	v130 = v122 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v130 < v131 {
		v122 = v130
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L48
L63:
	;
	goto L62
L64:
	;
	v139 = int32(0)
	goto L65
L65:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v111)+104))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v139<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v144)
	mBase = m.M
	v147 = v139 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v147 < v148 {
		v139 = v147
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L48
L67:
	;
	goto L66
L68:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+117)) = uint8(v152)
	goto L47
L69:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+116)) = uint8(v156)
	goto L47
L70:
	;
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+120)) = l0
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+116)) = uint8(v159)
	goto L47
L72:
	;
	goto L48
L73:
	;
	goto L48
L74:
	;
	goto L46
L75:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+117)) = uint8(v186)
	return
L76:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+116)) = uint8(v190)
	return
L77:
	;
	goto L78
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+120)) = l0
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+116)) = uint8(v193)
	return
L79:
	;
	goto L1
L80:
	;
	goto L1
}
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = int32(0)
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13 | v14<<(uint(v6)%32)
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v19)
	v25 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_ResetTupleHashTable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v9 = F__emscripten_memset_bulkmem(m, v3, base.I32_extend8_s(int32(0)), v5*int32(12))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(0)
	return
}
func F_TupleQueueReaderNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2 != 0 {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v10)
	} else {
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F_shm_mq_receive(m, v12, v8+int32(12), v8+int32(8), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		switch v17 - int32(1) {
		case 0:
			v28 = v4
		case 1:
			if l2 == int32(0) {
				v28 = v4
			} else {
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v25)
				v28 = v4
			}
		default:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v28 = v27
		}
		m.G0 = v8 + int32(16)
		return v28
	}
}
func F_exec_init_tuple_store(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(556508))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errmsg(m, int32(106894), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(500315), int32(3681), int32(364661))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		if v11 != int32(383) {
			F_errstart_cold(m, int32(21), int32(556508))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errmsg(m, int32(106894), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errfinish(m, int32(500315), int32(3681), int32(364661))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v14&int32(2) == int32(0) {
				F_errstart_cold(m, int32(21), int32(556508))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errmsg(m, int32(60436), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							F_errfinish(m, int32(500315), int32(3687), int32(364661))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(556508))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errmsg(m, int32(60436), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								F_errfinish(m, int32(500315), int32(3687), int32(364661))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v22 = int32(4515600)
					v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
					v26 = int32(4515652)
					v27 = *(*int32)(unsafe.Add(mBase, _consts[179]))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, _consts[179])) = v28
					v36 = *(*int32)(unsafe.Add(mBase, _consts[532]))
					v37 = F_tuplestore_begin_heap(m, int32(base.Ui32(v14&int32(4))>>(uint(int32(2))%32)), int32(0), v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v37
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
						*(*int32)(unsafe.Add(mBase, _consts[179])) = v27
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42
						return
					}
				}
			}
		}
	}
}
func F_expandTupleDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	v12 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = v19
	goto L3
L2:
	;
	v20 = v12
	goto L3
L3:
	;
	if l3 < v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v26 = v22 + l3<<(uint(int32(2))%32)
	goto L6
L5:
	;
	v26 = v12
	goto L6
L6:
	;
	if int32(0) < l2 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v42 = v26
	v44 = v12
	goto L10
L8:
	;
	goto L9
L9:
	;
	return
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = l0 + int32(20) + v48<<(uint(int32(4))%32) + v44*int32(100)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+91)))
	if v55 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	v147 = v44 + int32(1)
	if v147 != l2 {
		v42 = v145
		v44 = v147
		goto L10
	} else {
		goto L51
	}
L13:
	;
	if l8 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if l9 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v80 = int32(0)
	if v42 == v80 {
		v145 = v80
		goto L12
	} else {
		goto L28
	}
L17:
	;
	if l9 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v62 = F_pstrdup(m, int32(757603))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if l10 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	return
L22:
	;
	v64 = F_makeString(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v66 = F_lappend(m, v60, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v66
	goto L20
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v75 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v77 = F_lappend(m, v71, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v77
	goto L16
L28:
	;
	v84 = v42 + int32(4)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if base.Ui32(v84) < base.Ui32(v87+v88<<(uint(int32(2))%32)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = v84
	goto L31
L30:
	;
	v93 = int32(0)
	goto L31
L31:
	;
	v145 = v93
	goto L12
L32:
	;
	if l10 != 0 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v123 = v42
	goto L32
L34:
	;
	goto L35
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v42 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v116 = F_pstrdup(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L21
	} else {
		goto L43
	}
L37:
	;
	v98 = v42 + int32(4)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if base.Ui32(v98) < base.Ui32(v101+v102<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v113 = int32(0)
	v115 = v54 + int32(4)
	goto L36
L40:
	;
	v107 = v98
	goto L42
L41:
	;
	v107 = int32(0)
	goto L42
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v113 = v107
	v115 = v109
	goto L36
L43:
	;
	v118 = F_makeString(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v120 = F_lappend(m, v96, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v120
	v123 = v113
	goto L32
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v132 = F_makeVar(m, l4, base.I32_extend16_s(l3+v44+int32(1)), v129, v130, v131, l5)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L21
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v145 = v123
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v132)+32)) = l6
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v137 = F_lappend(m, v136, v132)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v137
	goto L48
L51:
	;
	goto L11
}
