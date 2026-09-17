package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetJsonPathVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v389
	m.G0 = v13 + int32(16)
	return v387
L2:
	;
	v387 = int32(0)
	v389 = int32(-1)
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		v387 = v6
		v389 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v18
	goto L7
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v6
	v34 = int32(1)
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if l2 == v41 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v97 = F_palloc(m, int32(20))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	goto L9
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if l2 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v91 = int32(1)
	v94 = v32 + v91
	if v24 != v94 {
		v32 = v94
		v34 = v34 + v91
		goto L8
	} else {
		goto L28
	}
L14:
	;
	if v88 == int32(0) {
		goto L10
	} else {
		goto L27
	}
L15:
	;
	v88 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = v43
	v51 = l1
	v52 = l2
	v53 = v49
	goto L22
L19:
	;
	v76 = l1
	v80 = int32(0)
	goto L20
L20:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v88 = v80 - v81
	goto L14
L21:
	;
	v76 = v71
	v80 = v73
	goto L20
L22:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.B2i32(v53 != v55)|base.B2i32(v55 == int32(0)) != 0 {
		v71 = v51
		v73 = v53
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v71 = v65
	v73 = int32(0)
	goto L21
L24:
	;
	v61 = v52 - int32(1)
	if v61 == int32(0) {
		v71 = v51
		v73 = v53
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v64 = int32(1)
	v65 = v51 + v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v66 != 0 {
		v50 = v50 + v64
		v51 = v65
		v52 = v61
		v53 = v66
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L13
L28:
	;
	goto L2
L29:
	;
	return int32(0)
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+20)))
	if v101 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v364
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v366
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v368
	v387 = v97
	v389 = v34
	goto L1
L32:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v104
	goto L31
L33:
	;
	goto L34
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v109 <= int32(1042) {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v308 = F_pg_detoast_datum(m, v108)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L29
	} else {
		goto L112
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)) = uint8(base.B2i32(v108 != int32(0)))
	goto L31
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L29
	} else {
		goto L107
	}
L38:
	;
	if v109 == int32(114) {
		goto L35
	} else {
		goto L106
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(18)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v250 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L40:
	;
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v205
	v208 = v108 + v205
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v211&v205 != 0 {
		goto L79
	} else {
		goto L80
	}
L41:
	;
	v198 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v108)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L29
	} else {
		goto L77
	}
L42:
	;
	v189 = F_DirectFunctionCall1Coll(m, int32(1400), int32(0), v108)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L29
	} else {
		goto L75
	}
L43:
	;
	v180 = F_DirectFunctionCall1Coll(m, int32(1399), int32(0), v108)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L29
	} else {
		goto L73
	}
L44:
	;
	v171 = F_DirectFunctionCall1Coll(m, int32(1398), int32(0), v108)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L29
	} else {
		goto L71
	}
L45:
	;
	v162 = F_DirectFunctionCall1Coll(m, int32(1397), int32(0), v108)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L29
	} else {
		goto L69
	}
L46:
	;
	switch v109 - int32(16) {
	case 0:
		goto L36
	case 1, 2, 3, 6, 8:
		goto L37
	case 4:
		goto L43
	case 5:
		goto L45
	case 7:
		goto L44
	case 9:
		goto L40
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v109 <= int32(1183) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	switch v109 - int32(700) {
	case 0:
		goto L42
	case 1:
		goto L41
	default:
		goto L38
	}
L50:
	;
	if v109 != int32(1700) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(32)
	goto L31
L52:
	;
	if base.Ui32(v109-int32(1082)) < base.Ui32(int32(2)) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if int32(1699) < v109 {
		goto L50
	} else {
		goto L58
	}
L55:
	;
	if v109 == int32(1043) {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	if v109 == int32(1114) {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L37
L58:
	;
	if v109 == int32(1184) {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	if v109 != int32(1266) {
		goto L37
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	if v109 != int32(3802) {
		goto L37
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v157 = F_pg_detoast_datum(m, v108)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L29
	} else {
		goto L68
	}
L64:
	;
	v144 = F_pg_detoast_datum(m, v108)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	v147 = v144 + int32(4)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+7)))
	if v148&int32(16) == int32(0) {
		goto L39
	} else {
		goto L66
	}
L66:
	;
	v153 = F_JsonbExtractScalar(m, v147, v97)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	goto L31
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v157
	goto L31
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v166 = F_pg_detoast_datum(m, v162)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v166
	goto L31
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v175 = F_pg_detoast_datum(m, v171)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L29
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v175
	goto L31
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v184 = F_pg_detoast_datum(m, v180)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L29
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v184
	goto L31
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v193 = F_pg_detoast_datum(m, v189)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v193
	goto L31
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(2)
	v202 = F_pg_detoast_datum(m, v198)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v202
	goto L31
L79:
	;
	v214 = v208
	goto L81
L80:
	;
	v214 = v108 + int32(4)
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v214
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v216 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v222 == int32(18) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	if v216&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	v225 = int32(16)
	goto L87
L86:
	;
	v225 = int32(0)
	goto L87
L87:
	;
	if base.Ui32((v222-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v232 = int32(4)
	goto L90
L89:
	;
	v232 = v225
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v232
	goto L31
L91:
	;
	v236 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v216)>>(uint(v236)%32)) - v236
	goto L31
L92:
	;
	goto L93
L93:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v241)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L94:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v256 == int32(18) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	if v250&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v259 = int32(16)
	goto L99
L98:
	;
	v259 = int32(0)
	goto L99
L99:
	;
	if base.Ui32((v256-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v266 = int32(4)
	goto L102
L101:
	;
	v266 = v259
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v266
	goto L31
L103:
	;
	v270 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v250)>>(uint(v270)%32)) - v270
	goto L31
L104:
	;
	goto L105
L105:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v275)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L106:
	;
	goto L37
L107:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L29
	} else {
		goto L108
	}
L108:
	;
	v290 = F_format_type_be(m, v109)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L29
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v290
	F_errmsg(m, int32(_a_F_GetJsonPathVar_0), v13)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L29
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_GetJsonPathVar_1), int32(3124), int32(_a_F_GetJsonPathVar_2))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L29
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v310 = F_text_to_cstring(m, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v312 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L29
	} else {
		goto L114
	}
L114:
	;
	v314 = F_pg_detoast_datum(m, v312)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L29
	} else {
		goto L115
	}
L115:
	;
	F_pfree(m, v310)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L29
	} else {
		goto L116
	}
L116:
	;
	v318 = F_pg_detoast_datum(m, v314)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L29
	} else {
		goto L117
	}
L117:
	;
	v321 = v318 + int32(4)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+7)))
	if v322&int32(16) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v325 = F_JsonbExtractScalar(m, v321, v97)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L29
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(18)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v330 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L31
L122:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v336 == int32(18) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v330&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	v339 = int32(16)
	goto L127
L126:
	;
	v339 = int32(0)
	goto L127
L127:
	;
	if base.Ui32((v336-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v346 = int32(4)
	goto L130
L129:
	;
	v346 = v339
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v346
	goto L31
L131:
	;
	v350 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v330)>>(uint(v350)%32)) - v350
	goto L31
L132:
	;
	goto L133
L133:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(base.Ui32(v355)>>(uint(int32(2))%32)) - int32(4)
	goto L31
}
func F_JsonTableFetchRow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableFetchRow_0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v8 = F_JsonTablePlanNextRow(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F__equalJsonIsPredicate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v24 = v3
			return v24
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v24 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v18 != v19 {
						v24 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
						v24 = base.B2i32(v21 == v22)
					}
				}
				return v24
			}
		}
	}
}
func F__equalJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v24 = v3
			return v24
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v24 = v3
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					if v18 != v19 {
						v24 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
						v24 = base.B2i32(v21 == v22)
					}
				}
				return v24
			}
		}
	}
}
func F_get_json_agg_constructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v17 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v32 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_appendStringInfoString(m, v10+int32(16), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	switch v16 - int32(1) {
	case 0, 2:
		v26 = int32(_a_F_get_json_agg_constructor_0)
		goto L4
	default:
		goto L3
	}
L6:
	;
	goto L7
L7:
	;
	switch v16 - int32(2) {
	case 0, 2:
		v26 = int32(_a_F_get_json_agg_constructor_6)
		goto L4
	default:
		goto L3
	}
L8:
	;
	goto L3
L9:
	;
	F_appendStringInfoString(m, v10+int32(16), int32(_a_F_get_json_agg_constructor_1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v40-int32(5)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_json_returning(m, v45, v10+int32(16), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v52 - int32(9) {
	case 0:
		goto L18
	default:
		goto L19
	case 2:
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v10 + int32(32)
	return
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_get_agg_expr_helper(m, v51, l1, v51, l2, v73, l3)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_get_windowfunc_expr_helper(m, v51, l1, l2, v55, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v63
	F_errmsg_internal(m, int32(_a_F_get_json_agg_constructor_2), v10)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_get_json_agg_constructor_3), int32(_a_F_get_json_agg_constructor_4), int32(_a_F_get_json_agg_constructor_5))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
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
	goto L17
}
func F_get_json_table_nested_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v11 - int32(50) {
	case 0:
		goto L4
	case 1:
		goto L5
	default:
		goto L1
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v55, int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v45, int32(44))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L4:
	;
	if l4 == int32(0) {
		v50 = l1
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_get_json_table_nested_columns(m, l0, v14, l2, l3, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	switch v18 - int32(50) {
	case 0:
		v40 = v17
		goto L3
	case 1:
		goto L8
	default:
		goto L1
	}
L8:
	;
	v22 = v17
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_get_json_table_nested_columns(m, l0, v27, l2, l3, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	if v32 == int32(50) {
		v40 = v31
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(51) {
		v22 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L1
L14:
	;
	v40 = l1
	goto L3
L15:
	;
	v50 = v40
	goto L2
L16:
	;
	v60 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_json_table_nested_columns_0), v60, v60, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	F_get_const_expr(m, v66, l2, int32(-1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = F_quote_identifier(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
	F_appendStringInfo(m, v70, int32(_a_F_get_json_table_nested_columns_1), v9)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_get_json_table_columns(m, l0, v50, l2, l3)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L1
}
func F_json_array_elements(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(_a_F_json_array_elements_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_array_elements_text(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(_a_F_json_array_elements_text_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_build_array_noargs(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_cstring_to_text_with_len(m, int32(_a_F_json_build_array_noargs_0), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_each(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_each_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v11 = v8 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_cstring_to_text(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_makeJsonLexContext(m, v11, v13, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = F_pg_parse_json_or_errsave(m, v11, int32(_a_F_json_in_0), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
					v28 = int32(0)
				} else {
					v28 = v13
				}
				m.G0 = v8 + int32(80)
				return v28
			}
		}
	}
}
func F_json_lex_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = l1 - v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14) <= base.Ui32(v13) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != 0 {
		goto L74
	} else {
		goto L75
	}
L2:
	;
	if base.Ui32(v14) <= base.Ui32(v72) {
		v134 = v70
		v136 = v72
		v139 = v75
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v70 = l1
	v72 = v13
	v75 = v11
	goto L2
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 == int32(48) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v70 = v59
	v72 = v61
	v75 = int32(0)
	goto L2
L7:
	;
	v19 = int32(1)
	v59 = l1 + v19
	v61 = v13 + v19
	goto L6
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(8)) < base.Ui32((v16-int32(49))&int32(255)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = l1
	v72 = v13
	v75 = v11
	goto L2
L11:
	;
	goto L12
L12:
	;
	v33 = l1
	v38 = v13
	goto L13
L13:
	;
	if v38 == v14-int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v59 = v48
	v61 = v45
	goto L6
L15:
	;
	v302 = v12 + v14
	v304 = v14
	v305 = int32(0)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v44 = int32(1)
	v45 = v38 + v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v48 = v33 + v44
	if base.Ui32((v46-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v33 = v48
		v38 = v45
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if base.Ui32(v14) <= base.Ui32(v136) {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v77 != int32(46) {
		v134 = v70
		v136 = v72
		v139 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v80 = int32(1)
	v82 = v70 + v80
	v84 = v72 + v80
	if v14 == v84 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v302 = v82
	v304 = v14
	v305 = v80
	goto L1
L23:
	;
	goto L24
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.Ui32((v86-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v134 = v124
	v136 = v125
	v139 = v129
	goto L19
L26:
	;
	v124 = v82
	v125 = v84
	v129 = int32(1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v95 = v72 + int32(2)
	if base.Ui32(v95) < base.Ui32(v14) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v14
	goto L31
L30:
	;
	v97 = v95
	goto L31
L31:
	;
	v102 = v82
	v103 = v84
	goto L32
L32:
	;
	v108 = int32(1)
	v109 = v102 + v108
	v111 = v103 + v108
	if base.Ui32(v14) <= base.Ui32(v111) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v124 = v109
	v125 = v111
	v129 = v75
	goto L25
L34:
	;
	v302 = v109
	v304 = v97
	v305 = v75
	goto L1
L35:
	;
	goto L36
L36:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.Ui32((v113-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v102 = v109
		v103 = v111
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if base.Ui32(v14) <= base.Ui32(v212) {
		goto L62
	} else {
		goto L63
	}
L39:
	;
	v211 = v134
	v212 = v136
	v215 = v139
	goto L38
L40:
	;
	goto L41
L41:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v141|int32(32) != int32(101) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v211 = v134
	v212 = v136
	v215 = v139
	goto L38
L43:
	;
	goto L44
L44:
	;
	v146 = int32(1)
	v148 = v134 + v146
	v150 = v136 + v146
	if base.Ui32(v14) <= base.Ui32(v150) {
		v159 = v150
		v160 = v148
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v159 == v14 {
		v302 = v160
		v304 = v14
		v305 = v146
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	switch v152 - int32(43) {
	case 0, 2:
		goto L47
	default:
		v159 = v150
		v160 = v148
		goto L45
	}
L47:
	;
	v155 = int32(2)
	v159 = v136 + v155
	v160 = v134 + v155
	goto L45
L48:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if base.Ui32((v162-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v211 = v201
	v212 = v197
	v215 = v205
	goto L38
L50:
	;
	v197 = v159
	v201 = v160
	v205 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v171 = v159 + int32(1)
	if base.Ui32(v171) < base.Ui32(v14) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v173 = v14
	goto L55
L54:
	;
	v173 = v171
	goto L55
L55:
	;
	v175 = v159
	v179 = v160
	goto L56
L56:
	;
	v184 = int32(1)
	v185 = v179 + v184
	v187 = v175 + v184
	if base.Ui32(v14) <= base.Ui32(v187) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v197 = v187
	v201 = v185
	v205 = v139
	goto L49
L58:
	;
	v302 = v185
	v304 = v173
	v305 = v139
	goto L1
L59:
	;
	goto L60
L60:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32((v189-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v175 = v187
		v179 = v185
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	v302 = v211
	v304 = v212
	v305 = v215
	goto L1
L63:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v211))))
	v220 = int32(255)
	v236 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32((v217-int32(48))&v220) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v217&int32(-33)-int32(65))&v220) < base.Ui32(int32(26)))|base.B2i32(v217 == int32(95)) == v236)&base.B2i32(v236 <= v217) != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v243 = int32(1)
	v245 = v212 + v243
	if v14 != v245 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v248 = v245
	v252 = v211
	goto L68
L66:
	;
	goto L67
L67:
	;
	v302 = v211 + (v14 - v212)
	v304 = v14
	v305 = v243
	goto L1
L68:
	;
	v258 = v252 + int32(1)
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v252)+1)))
	v262 = int32(255)
	v277 = int32(0)
	if base.B2i32(base.Ui32((v259-int32(48))&v262) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v259&int32(-33)-int32(65))&v262) < base.Ui32(int32(26)))|(base.B2i32(v259 == int32(95))|base.B2i32(v259 < v277)) == v277 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v302 = v258
	v304 = v248
	v305 = v243
	goto L1
L71:
	;
	goto L72
L72:
	;
	v284 = v248 + int32(1)
	if v284 != v14 {
		v248 = v284
		v252 = v258
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v304
	goto L76
L75:
	;
	goto L76
L76:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v308 != int32(1) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	return v338
L78:
	;
	if l2 != 0 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v312 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v304) < base.Ui32(v313) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v311+int32(4), v317, v302-v317)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	return int32(0)
L83:
	;
	if l2 == int32(0) {
		v338 = int32(1)
		goto L77
	} else {
		goto L84
	}
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v305)
	return int32(1)
L85:
	;
	v338 = int32(0)
	goto L77
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v305)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v302
	if v305 != 0 {
		v338 = int32(15)
		goto L77
	} else {
		goto L89
	}
L89:
	;
	goto L85
}
func F_json_manifest_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 - int32(6) {
	case 0, 4:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_json_manifest_array_end_0)
		m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v18, int32(_a_F_json_manifest_array_end_1), v6)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_json_to_record(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_json_to_record_0), int32(1), v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_json_to_recordset(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	F_populate_recordset_worker(m, l0, int32(_a_F_json_to_recordset_0), int32(1), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_typeof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = v6 + int32(12)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_makeJsonLexContext(m, v9, v11, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_json_lex(m, v9)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					F_json_errsave_error(m, v18, v9, int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
						v24 = int32(1)
						v25 = v23 - v24
						if base.B2i32(base.Ui32(v25) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1815))>>(uint(v25)%32))&v24) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
								F_errmsg_internal(m, int32(_a_F_json_typeof_0), v6)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_json_typeof_1), int32(1909), int32(_a_F_json_typeof_2))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_json_typeof[0])))
							v52 = F_cstring_to_text(m, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(80)
								return v52
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
					v24 = int32(1)
					v25 = v23 - v24
					if base.B2i32(base.Ui32(v25) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1815))>>(uint(v25)%32))&v24) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
							F_errmsg_internal(m, int32(_a_F_json_typeof_0), v6)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_json_typeof_1), int32(1909), int32(_a_F_json_typeof_2))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_json_typeof[0])))
						v52 = F_cstring_to_text(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(80)
							return v52
						}
					}
				}
			}
		}
	}
}
func F_json_unique_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v2 == int32(1) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		F_pfree(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
func F_json_unique_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v9 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v15 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = F_hash_search(m, v18, v7+int32(4), int32(1), v7+int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
	F_pfree(m, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_makeJsonTablePathSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v8 = F_palloc0(m, int32(20))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(123)
		v15 = F_palloc0(m, int32(20))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(2010044694600)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
			if l1 != 0 {
				v22 = F_pstrdup(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
					return v8
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
				return v8
			}
		}
	}
}
