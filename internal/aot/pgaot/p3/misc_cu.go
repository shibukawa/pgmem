package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_currtid_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_palloc(m, int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_currtid_internal[0]))
	v21 = F_pg_class_aclcheck(m, v17, v19, int64(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+119)))
	switch v24 - int32(73) {
	case 0, 32:
		goto L13
	default:
		v34 = int32(41)
		goto L8
	case 10:
		goto L12
	case 29:
		goto L9
	case 36:
		goto L10
	case 45:
		goto L11
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	switch v43 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L18
	default:
		goto L19
	case 35:
		goto L20
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_aclcheck_error(m, v21, v36, v37+int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v36 = v34
	goto L7
L9:
	;
	v34 = int32(18)
	goto L8
L10:
	;
	v36 = int32(23)
	goto L7
L11:
	;
	v36 = int32(51)
	goto L7
L12:
	;
	v36 = int32(37)
	goto L7
L13:
	;
	v36 = int32(20)
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L102
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L98
	}
L17:
	;
	m.G0 = v10 + int32(16)
	return v315
L18:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v288
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v290)
	v292 = F_GetLatestSnapshot(m)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L92
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L87
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if int32(0) < v47 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	v58 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L44
	}
L25:
	;
	v65 = v46 + v47<<(uint(int32(4))%32) + int32(20) + v58*int32(100)
	v67 = v65 + int32(4)
	v68 = int32(_a_F_currtid_internal_0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_currtid_internal[1])))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(0) {
		v91 = v71
		v92 = v72
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L24
L27:
	;
	if v92-v91 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	goto L27
L29:
	;
	if v71 != v72 {
		v91 = v71
		v92 = v72
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v76 = v67
	v77 = v68
	goto L31
L31:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v80
		v92 = v81
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v91 = v80
	v92 = v81
	goto L28
L33:
	;
	v84 = int32(1)
	if v80 == v81 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	if v96 == int32(27) {
		goto L21
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v116 = v58 + int32(1)
	if v116 != v47 {
		v58 = v116
		goto L25
	} else {
		goto L43
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(356), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	goto L26
L44:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_4), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(364), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v163 = int32(0)
	goto L57
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 <= int32(0) {
		goto L15
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	goto L48
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_5), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(369), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v145+v163<<(uint(int32(2))%32))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v174 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	if v180 == int32(0) {
		goto L16
	} else {
		goto L63
	}
L59:
	;
	v178 = v163 + int32(1)
	if v142 != v178 {
		v163 = v178
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	goto L15
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v183 != int32(1) {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+76))
	if v188 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v229 == int32(0) {
		goto L15
	} else {
		goto L78
	}
L66:
	;
	goto L65
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v195 <= int32(0) {
		v229 = int32(0)
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v229 = int32(0)
	goto L66
L70:
	;
	v198 = int32(0)
	if v198 < v195 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v201 = v195
	goto L73
L72:
	;
	v201 = v198
	goto L73
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v206 = int32(0)
	goto L74
L74:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(int32(2))%32))))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+8)))
	if v215 == base.I32_extend16_s(v58+int32(1))&int32(_a_F_currtid_internal_6) {
		v229 = v214
		goto L66
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v218 = v206 + int32(1)
	if v218 != v201 {
		v206 = v218
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v233 == int32(0) {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v236 != int32(6) {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v239 < int32(0) {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+8)))
	if v242 != int32(_a_F_currtid_internal_6) {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v187)+52))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246+v239<<(uint(int32(2))%32)-int32(4))))
	if v252 == int32(0) {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	v257 = F_table_open(m, v255, int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v259 = F_currtid_internal(m, v257, l1)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_sequence_close(m, v257, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v315 = v259
	goto L17
L87:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+68))
	v273 = F_get_namespace_name(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v275 + int32(4)
	F_errmsg(m, int32(_a_F_currtid_internal_7), v10)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(319), int32(_a_F_currtid_internal_8))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v294 = F_RegisterSnapshot(m, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v296 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+8))
	v302 = m.T0[v301].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v294, v296, v296, v296, int32(8))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_table_tuple_get_latest_tid(m, v302, v13)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+188))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	m.T0[v308].(func(*base.Module, int32))(m, v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_UnregisterSnapshot(m, v294)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v315 = v13
	goto L17
L98:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_9), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(381), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errmsg_internal(m, int32(_a_F_currtid_internal_10), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(408), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
