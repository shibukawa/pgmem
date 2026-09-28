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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
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
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
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
		v34 = int32(20)
		goto L8
	default:
		goto L9
	case 10:
		goto L13
	case 29:
		goto L10
	case 36:
		goto L11
	case 45:
		goto L12
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
		goto L19
	default:
		goto L20
	case 35:
		goto L21
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
	v34 = int32(42)
	goto L8
L10:
	;
	v36 = int32(18)
	goto L7
L11:
	;
	v36 = int32(23)
	goto L7
L12:
	;
	v36 = int32(52)
	goto L7
L13:
	;
	v36 = int32(38)
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L109
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L106
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L102
	}
L18:
	;
	m.G0 = v10 + int32(16)
	return v324
L19:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v291
	v293 = F_GetLatestSnapshot(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L92
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L87
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if int32(0) < v47 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v142 != 0 {
		goto L49
	} else {
		goto L50
	}
L23:
	;
	v58 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L44
	}
L26:
	;
	v65 = v46 + v47<<(uint(int32(3))%32) + int32(28) + v58*int32(100)
	v67 = v65 + int32(4)
	v68 = int32(_a_F_currtid_internal_0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_currtid_internal[1])))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L25
L28:
	;
	if v92-v93 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	v77 = v67
	v78 = v68
	goto L31
L31:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v92 = v82
	v93 = v81
	goto L29
L33:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	if v97 == int32(27) {
		goto L22
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v117 = v58 + int32(1)
	if v117 != v47 {
		v58 = v117
		goto L26
	} else {
		goto L43
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(384), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	goto L27
L44:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_4), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(392), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	v164 = int32(0)
	goto L57
L49:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v143 <= int32(0) {
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
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	goto L48
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_5), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(397), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
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
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v146+v164<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	if v181 == int32(0) {
		goto L17
	} else {
		goto L63
	}
L59:
	;
	v179 = v164 + int32(1)
	if v143 != v179 {
		v164 = v179
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
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v184 != int32(1) {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+76))
	if v189 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v230 == int32(0) {
		goto L15
	} else {
		goto L78
	}
L66:
	;
	goto L65
L67:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v196 <= int32(0) {
		v230 = int32(0)
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v230 = int32(0)
	goto L66
L70:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = v196
	goto L73
L72:
	;
	v202 = v199
	goto L73
L73:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v207 = int32(0)
	goto L74
L74:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203+v207<<(uint(int32(2))%32))))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+8)))
	if v216 == base.I32_extend16_s(v58+int32(1))&int32(_a_F_currtid_internal_6) {
		v230 = v215
		goto L66
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v219 = v207 + int32(1)
	if v219 != v202 {
		v207 = v219
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v234 == int32(0) {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v237 != int32(6) {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v240 < int32(0) {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+8)))
	if v243 != int32(_a_F_currtid_internal_6) {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v188)+52))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247+v240<<(uint(int32(2))%32)-int32(4))))
	if v253 == int32(0) {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v258 = F_table_open(m, v256, int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v260 = F_currtid_internal(m, v258, l1)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_relation_close(m, v258, int32(1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v324 = v260
	goto L18
L87:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+68))
	v274 = F_get_namespace_name(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v276 + int32(4)
	F_errmsg(m, int32(_a_F_currtid_internal_7), v10)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(347), int32(_a_F_currtid_internal_8))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
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
	v295 = F_RegisterSnapshot(m, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_currtid_internal[2]))
	if v298 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_currtid_internal[3])))
	if v300&int32(1) == int32(0) {
		goto L16
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v305 = int32(0)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	v311 = m.T0[v310].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v295, v305, v305, v305, int32(8))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	F_table_tuple_get_latest_tid(m, v311, v13)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+188))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	m.T0[v317].(func(*base.Module, int32))(m, v311)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_UnregisterSnapshot(m, v295)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v324 = v13
	goto L18
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_9), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(409), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errmsg_internal(m, int32(_a_F_currtid_internal_10), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_11), int32(931), int32(_a_F_currtid_internal_12))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errmsg_internal(m, int32(_a_F_currtid_internal_13), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(436), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
