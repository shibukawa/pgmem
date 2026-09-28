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
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	v12 = m.G0
	v14 = v12 - int32(128)
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
	case 18:
		goto L13
	case 19:
		goto L10
	case 23:
		v112 = int64(16511)
		v113 = int64(114)
		goto L9
	case 31:
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
	v112 = int64(262)
	v113 = int64(83)
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
	v112 = int64(256)
	v113 = int64(84)
	goto L9
L15:
	;
	v112 = int64(128)
	v113 = int64(102)
	goto L9
L16:
	;
	v112 = int64(768)
	v113 = int64(110)
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_SetDefaultACL_3)
	F_errmsg(m, int32(_a_F_SetDefaultACL_4), v14+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_SetDefaultACL_1), int32(1198), int32(_a_F_SetDefaultACL_2))
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
	v112 = int64(6)
	v113 = int64(76)
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_SetDefaultACL_5)
	F_errmsg(m, int32(_a_F_SetDefaultACL_4), v14+int32(32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_SetDefaultACL_1), int32(1209), int32(_a_F_SetDefaultACL_2))
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
	F_errmsg_internal(m, int32(_a_F_SetDefaultACL_0), v14)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_SetDefaultACL_1), int32(1217), int32(_a_F_SetDefaultACL_2))
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
	v116 = v112
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
	v120 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v121 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v122 = F_SearchSysCache3(m, int32(22), v120, v121, v113)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = F_merge_acl_with_grant(m, v149, v151, v152, v153, v154, v118, v155, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L49
	}
L40:
	;
	v145 = F_aclcopy(m, v38)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
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
	v136 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v130))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v136 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v142 = F_aclmembers(m, v136, v14+int32(124))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v149 = v136
	v150 = v142
	goto L39
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = int32(0)
	v149 = v145
	v150 = int32(0)
	goto L39
L49:
	;
	F_aclitemsort(m, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_aclitemsort(m, v38)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v162 = int32(0)
	if v156 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	F_relation_close(m, v19, int32(3))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L101
	}
L53:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L100
	}
L54:
	;
	if v206 != 0 {
		goto L72
	} else {
		goto L73
	}
L55:
	;
	if v38 == int32(0) {
		v202 = v162
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v164 != 0 {
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
	v206 = int32(1)
	goto L54
L61:
	;
	goto L62
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v206 = base.B2i32(v169 == int32(0))
	goto L54
L63:
	;
	v206 = v202
	goto L54
L64:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v164 != v174 {
		v202 = v162
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	if v176 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v184 = v176
	goto L68
L67:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v184 = (v177<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L68
L68:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v186 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v194 = v186
	goto L71
L70:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v194 = (v187<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L71
L71:
	;
	v198 = F_memcmp(m, v184+v156, v194+v38, v164<<(uint(int32(4))%32))
	mBase = m.M
	v202 = base.B2i32(v198 == int32(0))
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
	v224 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v224
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v234
	if v122 == v234 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(826)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+22)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v211+v212)))
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v214
	F_performDeletion(m, v14+int32(80), v215, v215)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	v244 = F_GetNewOidWithIndex(m, v19, int32(828), int32(1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+22)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306+v307)))
	v310 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+68)) = uint8(v310)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = base.I64_extend_i32_u(v156)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v321 = F_heap_modify_tuple(m, v122, v314, v14+int32(80), v14+int32(72), v14-int32(-64))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L93
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = base.I64_extend_i32_u(v244)
	v248 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v248
	v250 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = base.I64_extend_i32_u(v156)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v250
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v260 = F_heap_form_tuple(m, v255, v14+int32(80), v14+int32(72))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_CatalogTupleInsert(m, v19, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_recordDependencyOnOwner(m, int32(826), v244, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v268 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v269 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(826)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(2615)
	F_recordDependencyOn(m, v14+int32(52), v14+int32(40), int32(97))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v288 = F_aclmembers(m, v156, v14+int32(52))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_updateAclDependencies(m, int32(826), v244, int32(0), v292, v150, v293, v288, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_SetDefaultACL[0]))
	if v298 == int32(0) {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v302 = int32(0)
	F_RunObjectPostCreateHook(m, int32(826), v244, v302, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L77
L93:
	;
	F_CatalogTupleUpdate(m, v19, v321+int32(4), v321)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v329 = F_aclmembers(m, v156, v14+int32(52))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v14)+124))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_updateAclDependencies(m, int32(826), v309, int32(0), v333, v150, v334, v329, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_SetDefaultACL[0]))
	if v339 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L97:
	;
	v343 = int32(0)
	F_RunObjectPostAlterHook(m, int32(826), v309, v343, v343, v343)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
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
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	m.G0 = v14 + int32(128)
	return
}
func F_assign_default_text_search_config(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_assign_default_text_search_config[0])) = int32(0)
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
		*(*int32)(unsafe.Add(mBase, _c_F_check_default_with_oids[0])) = int32(1088)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_with_oids[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_default_with_oids[2])) = v11
		v17 = F_format_elog_string(m, int32(_a_F_check_default_with_oids_0), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_default_with_oids[3])) = v17
			return v4 ^ int32(1)
		}
	} else {
		return v4 ^ int32(1)
	}
}
