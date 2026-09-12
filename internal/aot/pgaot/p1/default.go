package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetDefaultACL(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int64
	_ = v225
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_table_open(m, int32(826), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v41 - int32(19) {
	case 0:
		goto L15
	default:
		goto L11
	case 3:
		goto L12
	case 17:
		goto L13
	case 18:
		goto L10
	case 22:
		v112 = int32(114)
		v113 = int64(16511)
		goto L9
	case 30:
		goto L14
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_acldefault(m, v24, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = F_palloc0(m, int32(24))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v26
	goto L3
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(4436701216768)
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(4294967392)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = int64(4294967296)
	v38 = v29
	goto L3
L9:
	;
	if v16 == int64(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	v112 = int32(83)
	v113 = int64(262)
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v72 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v112 = int32(84)
	v113 = int64(256)
	goto L9
L15:
	;
	v112 = int32(102)
	v113 = int64(128)
	goto L9
L16:
	;
	v112 = int32(110)
	v113 = int64(768)
	goto L9
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(524315)
	F_errmsg(m, int32(186444), v14+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(497783), int32(1212), int32(533854))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v112 = int32(76)
	v113 = int64(6)
	goto L9
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(522706)
	F_errmsg(m, int32(186444), v14+int32(32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(497783), int32(1223), int32(533854))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v100
	F_errmsg_internal(m, int32(484942), v14)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(497783), int32(1231), int32(533854))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v116 = v113
	goto L35
L34:
	;
	v116 = v16
	goto L35
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v117 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v116
	goto L38
L37:
	;
	v118 = v16
	goto L38
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = F_SearchSysCache3(m, int32(22), v120, v121, v112)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = F_merge_acl_with_grant(m, v148, v150, v151, v152, v153, v118, v154, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L40:
	;
	v144 = F_aclcopy(m, v38)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L48
	}
L41:
	;
	if v122 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v130 = F_SysCacheGetAttr(m, int32(22), v122, int32(5), v14+int32(80))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v132 == int32(1) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v135 = F_pg_detoast_datum_copy(m, v130)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v135 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v141 = F_aclmembers(m, v135, v14+int32(108))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v148 = v135
	v149 = v141
	goto L39
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = int32(0)
	v148 = v144
	v149 = int32(0)
	goto L39
L49:
	;
	F_aclitemsort(m, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_aclitemsort(m, v38)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v161 = int32(0)
	if v155 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	F_sequence_close(m, v19, int32(3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L101
	}
L53:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L100
	}
L54:
	;
	if v205 != 0 {
		goto L72
	} else {
		goto L73
	}
L55:
	;
	if v38 == int32(0) {
		v201 = v161
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	if v163 != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v38 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v205 = int32(1)
	goto L54
L61:
	;
	goto L62
L62:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v205 = base.B2i32(v168 == int32(0))
	goto L54
L63:
	;
	v205 = v201
	goto L54
L64:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v163 != v173 {
		v201 = v161
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v175 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v183 = v175
	goto L68
L67:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v183 = (v176<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L68
L68:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v185 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v193 = v185
	goto L71
L70:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v193 = (v186<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L71
L71:
	;
	v197 = F_memcmp(m, v183+v155, v193+v38, v163<<(uint(int32(4))%32))
	mBase = m.M
	v201 = base.B2i32(v197 == int32(0))
	goto L63
L72:
	;
	if v122 == int32(0) {
		goto L52
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v223
	v225 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v225
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v225
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v223)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v223
	if v122 == v223 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(826)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+22)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210+v211)))
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v213
	F_performDeletion(m, v14+int32(80), v214, v214)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L53
L77:
	;
	if v122 == int32(0) {
		goto L52
	} else {
		goto L99
	}
L78:
	;
	v239 = F_GetNewOidWithIndex(m, v19, int32(828), int32(1))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+22)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v299+v300)))
	v303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+68)) = uint8(v303)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v155
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v313 = F_heap_modify_tuple(m, v122, v306, v14+int32(80), v14+int32(72), v14-int32(-64))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L93
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v239
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v244
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v253 = F_heap_form_tuple(m, v248, v14+int32(80), v14+int32(72))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_CatalogTupleInsert(m, v19, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_recordDependencyOnOwner(m, int32(826), v239, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v261 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v262 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(826)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(2615)
	F_recordDependencyOn(m, v14+int32(52), v14+int32(40), int32(97))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v281 = F_aclmembers(m, v155, v14+int32(52))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_updateAclDependencies(m, int32(826), v239, int32(0), v285, v149, v286, v281, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v291 == int32(0) {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v295 = int32(0)
	F_RunObjectPostCreateHook(m, int32(826), v239, v295, v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L77
L93:
	;
	F_CatalogTupleUpdate(m, v19, v313+int32(4), v313)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v321 = F_aclmembers(m, v155, v14+int32(52))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_updateAclDependencies(m, int32(826), v302, int32(0), v325, v149, v326, v321, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v331 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L97:
	;
	v335 = int32(0)
	F_RunObjectPostAlterHook(m, int32(826), v302, v335, v335, v335)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L77
L99:
	;
	goto L53
L100:
	;
	goto L52
L101:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	m.G0 = v14 + int32(112)
	return
}
func F_assign_default_text_search_config(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1172])) = int32(0)
	return
}
func F_check_default_with_oids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4 == int32(1) {
		*(*int32)(unsafe.Add(mBase, _consts[428])) = int32(1088)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		*(*int32)(unsafe.Add(mBase, _consts[426])) = v11
		v17 = F_format_elog_string(m, int32(443993), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[427])) = v17
			return v4 ^ int32(1)
		}
	} else {
		return v4 ^ int32(1)
	}
}
