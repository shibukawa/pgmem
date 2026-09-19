package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTypeFromTL(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_ExecTypeFromTLInternal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_TypeCacheOpcCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheOpcCallback[0]))
	F_hash_seq_init(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_hash_seq_search(m, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+312)) = v19 & int32(_a_F_TypeCacheOpcCallback_0)
	if base.B2i32(v19&int32(-1572866) == int32(0))|v19&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v45 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+13)))
	if v30 != int32(99) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheOpcCallback[1]))
	v41 = F_hash_search(m, v35, v16+int32(16), int32(2), v6+int32(31))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	if v45 != 0 {
		v16 = v45
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
}
func F_TypeIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v136
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v136 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_TypeIsVisibleExt_0), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_TypeIsVisibleExt_1), int32(1067), int32(_a_F_TypeIsVisibleExt_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L40
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_TypeIsVisibleExt[0]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_TypeIsVisibleExt[0]))
	if v87 == int32(0) {
		v126 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v126 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v77 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v77
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v77 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v77 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= v90 {
		v126 = v3
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v96 = v90
	goto L34
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v109 = base.B2i32(v108 == v40)
	if v108 == v40 {
		v126 = v109
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v126 = v109
	goto L14
L36:
	;
	v111 = int32(0)
	v113 = F_SearchSysCacheExists(m, int32(81), v39+int32(4), v108, v111, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v113 != 0 {
		v126 = v109
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v116 = v96 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v116 < v117 {
		v96 = v116
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v136 = v126
	goto L1
}
func F_can_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
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
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
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
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v451 int32
	_ = v451
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	if l0 <= v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v24 + int32(16)
	return v451
L2:
	;
	v451 = int32(1)
	goto L1
L3:
	;
	v34 = v5
	v38 = v5
	goto L4
L4:
	;
	v50 = v38 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+v50)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v50)))
	if base.B2i32(v52 == v54)|base.B2i32(v54 == int32(2276)) != 0 {
		v115 = v34
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v115&int32(1) == int32(0) {
		goto L2
	} else {
		goto L46
	}
L6:
	;
	v117 = v38 + int32(1)
	if v117 != l0 {
		v34 = v115
		v38 = v117
		goto L4
	} else {
		goto L45
	}
L7:
	;
	v59 = int32(1)
	if v54 <= int32(3830) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v52 == int32(705) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	switch v54 - int32(2277) {
	case 0, 6:
		v115 = v59
		goto L6
	case 1, 2, 3, 4, 5:
		goto L8
	default:
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if base.B2i32(base.Ui32(v54-int32(_a_F_can_coerce_type_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v54-int32(_a_F_can_coerce_type_1)) < base.Ui32(int32(2)))|base.B2i32(v54 == int32(3831)) != 0 {
		v115 = v59
		goto L6
	} else {
		goto L15
	}
L12:
	;
	if v54 == int32(2776) {
		v115 = v59
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v54 != int32(3500) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v115 = v59
	goto L6
L15:
	;
	goto L8
L16:
	;
	v115 = v34
	goto L6
L17:
	;
	goto L18
L18:
	;
	v84 = F_find_coercion_pathway(m, v54, v52, l3, v24+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v84 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v115 = v34
	goto L6
L22:
	;
	goto L23
L23:
	;
	if v52 != int32(2249) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v54 != int32(2287) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v90 = F_typeOrDomainTypeRelid(m, v54)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if v90 == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v115 = v34
	goto L6
L28:
	;
	v111 = F_typeInheritsFrom(m, v52, v54)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L39
	}
L29:
	;
	if v54 != int32(2249) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v102 = F_get_element_type(m, v52)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L35
	}
L32:
	;
	v98 = F_typeOrDomainTypeRelid(m, v52)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if v98 == int32(0) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v115 = v34
	goto L6
L35:
	;
	if v102 == int32(0) {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v106 = F_typeOrDomainTypeRelid(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	if v106 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v115 = v34
	goto L6
L39:
	;
	if v111 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = v34
	goto L6
L41:
	;
	goto L42
L42:
	;
	v113 = F_typeIsOfTypedTable(m, v52, v54)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	if v113 != 0 {
		v115 = v34
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v451 = v5
	goto L1
L45:
	;
	goto L5
L46:
	;
	v123 = int32(0)
	v129 = m.G0
	v131 = v129 - int32(416)
	m.G0 = v131
	if l0 <= v123 {
		v395 = int32(1)
		goto L47
	} else {
		goto L48
	}
L47:
	;
	m.G0 = v131 + int32(416)
	if v395 == int32(0) {
		v451 = v5
		goto L1
	} else {
		goto L176
	}
L48:
	;
	v138 = v123
	v139 = v123
	v140 = v123
	v142 = int32(0)
	v143 = v123
	v144 = v123
	v146 = v5
	v147 = v123
	v149 = v5
	v152 = v5
	v155 = v5
	v156 = v5
	v157 = v5
	goto L49
L49:
	;
	v159 = v142 << (uint(int32(2)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1+v159)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+l2)))
	if v163 <= int32(3830) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	if base.B2i32(v252 == int32(0))|base.B2i32(v252 == int32(2277)) != 0 {
		v280 = v250
		goto L114
	} else {
		goto L115
	}
L51:
	;
	v264 = v142 + int32(1)
	if v264 != l0 {
		v138 = v250
		v139 = v251
		v140 = v252
		v142 = v264
		v143 = v254
		v144 = v255
		v146 = v256
		v147 = v257
		v149 = v258
		v152 = v259
		v155 = v260
		v156 = v261
		v157 = v262
		goto L49
	} else {
		goto L113
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131+v146<<(uint(int32(2))%32)))) = v238
	v250 = v138
	v251 = v139
	v252 = v140
	v254 = v143
	v255 = v240
	v256 = v146 + int32(1)
	v257 = v147
	v258 = v241
	v259 = v152
	v260 = v155
	v261 = v156
	v262 = v242
	goto L51
L53:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L105
	}
L54:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L97
	}
L55:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L93
	}
L56:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v210
		goto L51
	} else {
		goto L92
	}
L57:
	;
	v210 = int32(1)
	goto L56
L58:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L86
	}
L59:
	;
	if v187 == v139 {
		v250 = v138
		v251 = v187
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L85
	}
L60:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L79
	}
L61:
	;
	if v161 == v138 {
		v250 = v161
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v174
		v261 = v175
		v262 = v157
		goto L51
	} else {
		goto L78
	}
L62:
	;
	switch v163 - int32(2277) {
	case 0:
		goto L60
	case 1, 2, 3, 4, 5:
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	case 6:
		v174 = v155
		v175 = v156
		goto L65
	default:
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	switch v163 - int32(_a_F_can_coerce_type_0) {
	case 0:
		v210 = v157
		goto L56
	case 1:
		goto L55
	case 2:
		goto L57
	case 3:
		goto L54
	default:
		goto L73
	}
L65:
	;
	if v161 == int32(705) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v174
		v261 = v175
		v262 = v157
		goto L51
	} else {
		goto L71
	}
L66:
	;
	if v163 == int32(2776) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v174 = int32(1)
	v175 = v156
	goto L65
L68:
	;
	goto L69
L69:
	;
	if v163 != int32(3500) {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L70
	}
L70:
	;
	v174 = v155
	v175 = int32(1)
	goto L65
L71:
	;
	if v138 != 0 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	v250 = v161
	v251 = v139
	v252 = v140
	v254 = v143
	v255 = v144
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v152
	v260 = v174
	v261 = v175
	v262 = v157
	goto L51
L73:
	;
	switch v163 - int32(_a_F_can_coerce_type_1) {
	case 0:
		goto L58
	case 1:
		goto L53
	default:
		goto L74
	}
L74:
	;
	if base.B2i32(v161 == int32(705))|base.B2i32(v163 != int32(3831)) != 0 {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L75
	}
L75:
	;
	v187 = F_getBaseType(m, v161)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	if v139 != 0 {
		goto L59
	} else {
		goto L77
	}
L77:
	;
	v250 = v138
	v251 = v187
	v252 = v140
	v254 = v143
	v255 = v144
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v152
	v260 = v155
	v261 = v156
	v262 = v157
	goto L51
L78:
	;
	v395 = int32(0)
	goto L47
L79:
	;
	v193 = F_getBaseType(m, v161)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	if v140 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v250 = v138
	v251 = v139
	v252 = v193
	v254 = v143
	v255 = v144
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v152
	v260 = v155
	v261 = v156
	v262 = v157
	goto L51
L82:
	;
	goto L83
L83:
	;
	if v193 == v140 {
		v250 = v138
		v251 = v139
		v252 = v193
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L84
	}
L84:
	;
	v395 = int32(0)
	goto L47
L85:
	;
	v395 = int32(0)
	goto L47
L86:
	;
	v203 = F_getBaseType(m, v161)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L87
	}
L87:
	;
	if v143 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v250 = v138
	v251 = v139
	v252 = v140
	v254 = v203
	v255 = v144
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v152
	v260 = v155
	v261 = v156
	v262 = v157
	goto L51
L89:
	;
	goto L90
L90:
	;
	if v203 == v143 {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v203
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L91
	}
L91:
	;
	v395 = int32(0)
	goto L47
L92:
	;
	v238 = v161
	v240 = v144
	v241 = v149
	v242 = v210
	goto L52
L93:
	;
	v215 = F_getBaseType(m, v161)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L94
	}
L94:
	;
	v217 = F_get_element_type(m, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L19
	} else {
		goto L95
	}
L95:
	;
	if v217 != 0 {
		v238 = v217
		v240 = v144
		v241 = v149
		v242 = v157
		goto L52
	} else {
		goto L96
	}
L96:
	;
	v395 = int32(0)
	goto L47
L97:
	;
	v222 = F_getBaseType(m, v161)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	if v144 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v222 == v144 {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v226 = F_get_range_subtype(m, v222)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L19
	} else {
		goto L103
	}
L102:
	;
	v395 = int32(0)
	goto L47
L103:
	;
	if v226 != 0 {
		v238 = v226
		v240 = v222
		v241 = v226
		v242 = v157
		goto L52
	} else {
		goto L104
	}
L104:
	;
	v395 = int32(0)
	goto L47
L105:
	;
	v231 = F_getBaseType(m, v161)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L106
	}
L106:
	;
	if v147 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v231 == v147 {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v152
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v235 = F_get_multirange_range(m, v231)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L111
	}
L110:
	;
	v395 = int32(0)
	goto L47
L111:
	;
	if v235 != 0 {
		v250 = v138
		v251 = v139
		v252 = v140
		v254 = v143
		v255 = v144
		v256 = v146
		v257 = v231
		v258 = v149
		v259 = v235
		v260 = v155
		v261 = v156
		v262 = v157
		goto L51
	} else {
		goto L112
	}
L112:
	;
	v395 = int32(0)
	goto L47
L113:
	;
	goto L50
L114:
	;
	if v254 != 0 {
		goto L125
	} else {
		goto L126
	}
L115:
	;
	v271 = int32(0)
	v272 = F_get_element_type(m, v252)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	if v272 == int32(0) {
		v395 = v271
		goto L47
	} else {
		goto L117
	}
L117:
	;
	if v250 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v280 = v272
	goto L114
L119:
	;
	goto L120
L120:
	;
	if v272 != v250 {
		v395 = v271
		goto L47
	} else {
		goto L121
	}
L121:
	;
	v280 = v250
	goto L114
L122:
	;
	if v260 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L123:
	;
	v294 = int32(0)
	v295 = F_get_range_subtype(m, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L19
	} else {
		goto L135
	}
L124:
	;
	if base.B2i32(v283 == v251) == int32(0) {
		v395 = v282
		goto L47
	} else {
		goto L134
	}
L125:
	;
	v282 = int32(0)
	v283 = F_get_multirange_range(m, v254)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L19
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v251 != 0 {
		v292 = v251
		goto L123
	} else {
		goto L133
	}
L128:
	;
	if v283 == int32(0) {
		v395 = v282
		goto L47
	} else {
		goto L129
	}
L129:
	;
	if v251 != 0 {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v287 = F_get_range_subtype(m, v283)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	if v287 != 0 {
		v292 = v283
		goto L123
	} else {
		goto L132
	}
L132:
	;
	v395 = v282
	goto L47
L133:
	;
	v303 = v280
	goto L122
L134:
	;
	v292 = v251
	goto L123
L135:
	;
	if v295 == int32(0) {
		v395 = v294
		goto L47
	} else {
		goto L136
	}
L136:
	;
	if v280 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v303 = v295
	goto L122
L138:
	;
	goto L139
L139:
	;
	if v295 != v280 {
		v395 = v294
		goto L47
	} else {
		goto L140
	}
L140:
	;
	v303 = v280
	goto L122
L141:
	;
	if v261 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v307 = F_get_base_element_type(m, v303)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	if v307 == int32(0) {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v395 = int32(0)
	goto L47
L145:
	;
	if v257 == int32(0) {
		v332 = v256
		v333 = v258
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v314 = F_type_is_enum(m, v303)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	if v314 != 0 {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v395 = int32(0)
	goto L47
L149:
	;
	if v332 <= int32(0) {
		v395 = int32(1)
		goto L47
	} else {
		goto L159
	}
L150:
	;
	if v255 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v255 == v259 {
		v332 = v256
		v333 = v258
		goto L149
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v321 = F_get_range_subtype(m, v259)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L19
	} else {
		goto L155
	}
L154:
	;
	v395 = int32(0)
	goto L47
L155:
	;
	if v321 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v395 = int32(0)
	goto L47
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131+v256<<(uint(int32(2))%32)))) = v321
	v332 = v256 + int32(1)
	v333 = v321
	goto L149
L159:
	;
	v338 = F_select_common_type_from_oids(m, v332, v131, int32(1))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	if v338 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v395 = int32(0)
	goto L47
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+412)) = v338
	v350 = int32(0)
	goto L165
L164:
	;
	if v262 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	v373 = F_can_coerce_type(m, int32(1), v131+v350<<(uint(int32(2))%32), v131+int32(412), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L19
	} else {
		goto L167
	}
L166:
	;
	v395 = int32(0)
	goto L47
L167:
	;
	if v373 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v376 = v350 + int32(1)
	if v332 != v376 {
		v350 = v376
		goto L165
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	goto L166
L171:
	;
	goto L164
L172:
	;
	v395 = base.B2i32(v333 == int32(0)) | base.B2i32(v338 == v333)
	goto L47
L173:
	;
	v381 = F_get_base_element_type(m, v338)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L19
	} else {
		goto L174
	}
L174:
	;
	if v381 == int32(0) {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v395 = int32(0)
	goto L47
L176:
	;
	goto L2
}
func F_format_type_be_qualified(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_format_type_extended(m, l0, int32(-1), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_getTypeBinaryInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+82)))
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = F_format_type_be(m, l0)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v56
							F_errmsg(m, int32(_a_F_getTypeBinaryInputInfo_0), v9+int32(32))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3094), int32(_a_F_getTypeBinaryInputInfo_2))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+108))
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v76 = F_format_type_be(m, l0)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
								F_errmsg(m, int32(_a_F_getTypeBinaryInputInfo_3), v9+int32(16))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3099), int32(_a_F_getTypeBinaryInputInfo_2))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
					v26 = v24 + v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
					if v27 != 0 {
						v29 = v27
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						v29 = v28
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeBinaryInputInfo_4), v9)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3087), int32(_a_F_getTypeBinaryInputInfo_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
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
func F_get_record_type_from_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_get_call_result_type(m, l0, int32(0), v7+int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(1) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
			if v19 != 0 {
				F_FreeTupleDesc(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v23 = v22
					v24 = int32(_a_F_get_record_type_from_query_0)
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0]))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
					*(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0])) = v27
					v29 = F_CreateTupleDescCopy(m, v23)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+52)) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+56)) = v33
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+60)) = v35
						*(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0])) = v25
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				v23 = v16
				v24 = int32(_a_F_get_record_type_from_query_0)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
				*(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0])) = v27
				v29 = F_CreateTupleDescCopy(m, v23)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2)+52)) = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+56)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+60)) = v35
					*(*int32)(unsafe.Add(mBase, _c_F_get_record_type_from_query[0])) = v25
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(_a_F_get_record_type_from_query_1), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_get_record_type_from_query_2), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_get_record_type_from_query_3), int32(3677), int32(_a_F_get_record_type_from_query_4))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
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
}
func F_get_type_func_class(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v6 = F_get_typtype(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		switch v6&int32(255) - int32(98) {
		case 0, 3, 11, 16:
			v33 = int32(0)
			return v33
		case 1:
			v33 = int32(1)
			return v33
		case 2:
			v14 = F_getBaseType(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
				v19 = F_get_typtype(m, v14)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 == int32(99) {
						v23 = int32(2)
					} else {
						v23 = int32(0)
					}
					return v23
				}
			}
		default:
			return int32(4)
		case 14:
			switch l0 - int32(2249) {
			case 0:
				v33 = int32(3)
				return v33
			default:
				return int32(4)
			case 26, 29:
				v33 = int32(0)
				return v33
			}
		}
	}
}
func F_has_type_privilege_name(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13919(m, l0, int32(_a_F_has_type_privilege_name_0), int32(1247), int32(_a_F_has_type_privilege_name_1), int32(_a_F_has_type_privilege_name_2), int32(_a_F_has_type_privilege_name_3), int32(67137668), int32(1238))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_makeTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_makeString(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v8
		v17 = F_list_make1_impl(m, int32(1), v6+int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = F_palloc0(m, int32(32))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(68)
				m.G0 = v6 + int32(16)
				return v20
			}
		}
	}
}
func F_typeOrDomainTypeRelid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v9
	goto L7
L5:
	;
	v28 = l0
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v18 = v16 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)))
	if v19 != int32(100) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v28 = v22
	goto L6
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v26 = F_SearchSysCache1(m, int32(82), v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v26 != 0 {
		v14 = v26
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28
	F_errmsg_internal(m, int32(_a_F_typeOrDomainTypeRelid_0), v6)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_typeOrDomainTypeRelid_1), int32(699), int32(_a_F_typeOrDomainTypeRelid_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	m.G0 = v6 + int32(16)
	return v44
}
