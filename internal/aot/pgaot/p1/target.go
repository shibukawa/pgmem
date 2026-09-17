package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_target_list(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_initStringInfo(m, v15+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	F_pfree(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = v3
	v32 = int32(_a_F_get_target_list_0)
	v37 = v3
	v39 = v3
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+26)))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoString(m, v17, v32)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v273 = v30
	v275 = v32
	v282 = v39
	goto L10
L10:
	;
	v284 = v37 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v284 < v285 {
		v30 = v273
		v32 = v275
		v37 = v284
		v39 = v282
		goto L6
	} else {
		goto L74
	}
L11:
	;
	v51 = v15 + int32(16)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v53
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v51
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v60 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v78 = v30 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	F_get_rule_expr(m, v60, l1, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v63 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v67 = F_get_variable(m, v60, int32(1), l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v76 = v67
	goto L13
L18:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+33)))
	if v74 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = int32(0)
	goto L21
L20:
	;
	v75 = int32(_a_F_get_target_list_1)
	goto L21
L21:
	;
	v76 = v75
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v137&int32(2) == int32(0) {
		v265 = v39
		goto L41
	} else {
		goto L42
	}
L23:
	;
	if v76 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v93 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v82 < v78 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = v79 + v82<<(uint(int32(4))%32) + v30*int32(100) + int32(24)
	goto L23
L27:
	;
	v96 = v93
	goto L23
L28:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if base.B2i32(v100 == int32(0))|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v126 = F_quote_identifier(m, v96)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L39
	}
L31:
	;
	if v121-v122 == int32(0) {
		goto L22
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	v106 = v76
	v107 = v96
	goto L34
L34:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v121 = v111
	v122 = v110
	goto L32
L36:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L30
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v126
	F_appendStringInfo(m, v15+int32(16), int32(_a_F_get_target_list_2), v15)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L22
L41:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_appendBinaryStringInfo(m, v17, v266, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v142 < int32(0) {
		v265 = v39
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v145 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v244 = int32(10)
	v245 = F___strchrnul(m, v240+v232+int32(1), v244)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v247 == v244 {
		goto L70
	} else {
		goto L71
	}
L45:
	;
	v186 = int32(-1)
	if v78 < int32(2) {
		v232 = v186
		goto L44
	} else {
		goto L57
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v149 != int32(10) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v152 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v232 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v158 = v152
	goto L51
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v158-int32(1)))))
	if v172 == int32(32) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v232 = int32(0)
	goto L44
L53:
	;
	v176 = v158 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v176
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v176+v168))) = uint8(v178)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v178 < v182 {
		v158 = v182
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v232 = v178
	goto L44
L57:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v193 = F_strlen(m, v189)
	mBase = m.M
	v200 = v193 + int32(1)
	goto L60
L58:
	;
	if v212 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	goto L58
L60:
	;
	v202 = int32(0)
	if v200 == v202 {
		v212 = v202
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v212 = v207
	goto L59
L62:
	;
	v206 = v200 - int32(1)
	v207 = v189 + v206
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v208 != int32(10) {
		v200 = v206
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v215 = v212 + int32(1)
	goto L66
L65:
	;
	v215 = v189
	goto L66
L66:
	;
	v216 = F_strlen(m, v215)
	mBase = m.M
	if v39|base.B2i32(base.Ui32(v142) < base.Ui32(v216+v145)) == int32(0) {
		v232 = v186
		goto L44
	} else {
		goto L67
	}
L67:
	;
	F_appendContextKeyword(m, l1, int32(_a_F_get_target_list_3), int32(-8), int32(8), int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v232 = v186
	goto L44
L69:
	;
	v265 = base.B2i32(v251 != int32(0))
	goto L41
L70:
	;
	v251 = v245
	goto L72
L71:
	;
	v251 = int32(0)
	goto L72
L72:
	;
	goto L69
L73:
	;
	v273 = v78
	v275 = int32(_a_F_get_target_list_4)
	v282 = v265
	goto L10
L74:
	;
	goto L7
L75:
	;
	m.G0 = v15 + int32(32)
	return
}
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = l1
	v4 = l3
	v7 = F_palloc0(m, int32(28))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)) = uint8(v4)
		v12 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+24)) = uint16(v12)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(62)
		return v7
	}
}
func F_markTargetListOrigins(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
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
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L47
	} else {
		goto L83
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L47
	} else {
		goto L80
	}
L3:
	;
	m.G0 = v13 + int32(32)
	return
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = v3
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v29<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v251 = v29 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v251 < v252 {
		v29 = v251
		goto L6
	} else {
		goto L79
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v38 != int32(6) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v43 = int32(0)
	if v42 <= v43 {
		v90 = l0
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+8)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	switch v105 {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L8
	case 6:
		goto L26
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+v41<<(uint(int32(2))%32)-int32(4))))
	goto L11
L13:
	;
	v49 = v42 & int32(7)
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if base.Ui32(v42) < base.Ui32(int32(8)) {
		v90 = v64
		goto L12
	} else {
		goto L21
	}
L15:
	;
	v64 = l0
	v67 = v42
	goto L14
L16:
	;
	goto L17
L17:
	;
	v52 = l0
	v55 = v42
	v57 = v43
	goto L18
L18:
	;
	v58 = int32(1)
	v59 = v55 - v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v62 = v57 + v58
	if v62 != v49 {
		v52 = v60
		v55 = v59
		v57 = v62
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v64 = v60
	v67 = v59
	goto L14
L20:
	;
	goto L19
L21:
	;
	v72 = v64
	v75 = v67
	goto L22
L22:
	;
	v78 = int32(8)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v78 < v75 {
		v72 = v87
		v75 = v75 - v78
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v90 = v87
	goto L12
L24:
	;
	goto L23
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)) = uint16(v240)
	goto L8
L26:
	;
	if v104 == int32(0) {
		goto L8
	} else {
		goto L45
	}
L27:
	;
	if v104 == int32(0) {
		goto L8
	} else {
		goto L29
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v106
	v240 = v104
	goto L25
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+36))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+76))
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v149 == int32(0) {
		goto L2
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v115 <= int32(0) {
		v149 = int32(0)
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v149 = int32(0)
	goto L31
L35:
	;
	v118 = int32(0)
	if v118 < v115 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v121 = v115
	goto L38
L37:
	;
	v121 = v118
	goto L38
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v126 = int32(0)
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122+v126<<(uint(int32(2))%32))))
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+8)))
	if v135 == v104&int32(_a_F_markTargetListOrigins_0) {
		v149 = v134
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v138 = v126 + int32(1)
	if v138 != v121 {
		v126 = v138
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+26)))
	if v153 == int32(1) {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v156
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+24)))
	v240 = v158
	goto L25
L45:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+92)))
	if v161 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v162 = F_GetCTEForRTE(m, l0, v103, v42)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+16))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v167 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = int32(76)
	goto L51
L50:
	;
	v170 = int32(96)
	goto L51
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164+v170)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	v175 = base.B2i32(v173 != int32(0))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v162)+24))
	if v178 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v172 != 0 {
		goto L66
	} else {
		goto L67
	}
L53:
	;
	v179 = v175 | int32(2)
	goto L55
L54:
	;
	v179 = v175
	goto L55
L55:
	;
	if v179 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if v172 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v104 <= v189+v179 {
		goto L8
	} else {
		goto L63
	}
L58:
	;
	v184 = int32(0)
	if v184 < v104 {
		v189 = v184
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v104 <= v187 {
		goto L52
	} else {
		goto L62
	}
L61:
	;
	goto L52
L62:
	;
	v189 = v187
	goto L57
L63:
	;
	goto L52
L64:
	;
	if v230 == int32(0) {
		goto L1
	} else {
		goto L77
	}
L65:
	;
	goto L64
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v196 <= int32(0) {
		v230 = int32(0)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v230 = int32(0)
	goto L65
L69:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v202 = v196
	goto L72
L71:
	;
	v202 = v199
	goto L72
L72:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v207 = int32(0)
	goto L73
L73:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203+v207<<(uint(int32(2))%32))))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+8)))
	if v216 == v104&int32(_a_F_markTargetListOrigins_0) {
		v230 = v215
		goto L65
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v219 = v207 + int32(1)
	if v219 != v202 {
		v207 = v219
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+26)))
	if v234 == int32(1) {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v237
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+24)))
	v240 = v239
	goto L25
L79:
	;
	goto L7
L80:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v272
	F_errmsg_internal(m, int32(_a_F_markTargetListOrigins_1), v13)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_markTargetListOrigins_2), int32(372), int32(_a_F_markTargetListOrigins_3))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v288
	F_errmsg_internal(m, int32(_a_F_markTargetListOrigins_4), v13+int32(16))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L47
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_markTargetListOrigins_2), int32(418), int32(_a_F_markTargetListOrigins_3))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setTargetTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v17 = F_get_visible_ENR_metadata(m, v16, v15)
		mBase = m.M
		if v17 != int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
					F_errmsg(m, int32(_a_F_setTargetTable_0), v10)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_setTargetTable_1), int32(192), int32(_a_F_setTargetTable_2))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v20 != 0 {
				F_relation_close(m, v20, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = F_parserOpenTable(m, l0, l1, int32(3))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v27
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v33 = F_addRangeTableEntryForRelation(m, l0, v27, int32(3), v31, l2, int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v33
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = l4
							if l3 != 0 {
								v38 = int32(1)
								F_addNSItemToQuery(m, l0, v33, v38, v38, v38)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
									m.G0 = v10 + int32(16)
									return v43
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
								m.G0 = v10 + int32(16)
								return v43
							}
						}
					}
				}
			} else {
				v27 = F_parserOpenTable(m, l0, l1, int32(3))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v33 = F_addRangeTableEntryForRelation(m, l0, v27, int32(3), v31, l2, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v33
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = l4
						if l3 != 0 {
							v38 = int32(1)
							F_addNSItemToQuery(m, l0, v33, v38, v38, v38)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
								m.G0 = v10 + int32(16)
								return v43
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
							m.G0 = v10 + int32(16)
							return v43
						}
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v20 != 0 {
			F_relation_close(m, v20, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = F_parserOpenTable(m, l0, l1, int32(3))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v33 = F_addRangeTableEntryForRelation(m, l0, v27, int32(3), v31, l2, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v33
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = l4
						if l3 != 0 {
							v38 = int32(1)
							F_addNSItemToQuery(m, l0, v33, v38, v38, v38)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
								m.G0 = v10 + int32(16)
								return v43
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
							m.G0 = v10 + int32(16)
							return v43
						}
					}
				}
			}
		} else {
			v27 = F_parserOpenTable(m, l0, l1, int32(3))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v27
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v33 = F_addRangeTableEntryForRelation(m, l0, v27, int32(3), v31, l2, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v33
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = l4
					if l3 != 0 {
						v38 = int32(1)
						F_addNSItemToQuery(m, l0, v33, v38, v38, v38)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
							m.G0 = v10 + int32(16)
							return v43
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
						m.G0 = v10 + int32(16)
						return v43
					}
				}
			}
		}
	}
}
