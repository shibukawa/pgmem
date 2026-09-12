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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
	F_hash_seq_init(m, v6+int32(8), v11)
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
	v16 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v16
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+312)) = v21 & int32(1572865)
	if v21&int32(-1572866) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v46 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	if v21&int32(1) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
	if v31 != int32(99) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	if v34 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1092]))
	v42 = F_hash_search(m, v36, v18+int32(16), int32(2), v6+int32(31))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	if v46 != 0 {
		v18 = v46
		goto L7
	} else {
		goto L16
	}
L16:
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
	var v76 int32
	_ = v76
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
	F_errmsg_internal(m, int32(50330), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(499480), int32(1067), int32(64529))
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
	v44 = *(*int32)(unsafe.Add(mBase, _consts[437]))
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
	v87 = *(*int32)(unsafe.Add(mBase, _consts[437]))
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
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
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
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
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
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v452 int32
	_ = v452
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = int32(1)
	if l0 <= v5 {
		v452 = v26
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(16)
	return v452
L2:
	;
	v34 = v5
	v41 = v5
	goto L4
L3:
	;
	v452 = int32(0)
	goto L1
L4:
	;
	v51 = v41 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1+v51)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2+v51)))
	if v53 == v55 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if v116&int32(1) == int32(0) {
		v452 = v26
		goto L1
	} else {
		goto L47
	}
L6:
	;
	v118 = v41 + int32(1)
	if v118 != l0 {
		v34 = v115
		v41 = v118
		goto L4
	} else {
		goto L46
	}
L7:
	;
	v111 = F_typeIsOfTypedTable(m, v53, v55)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L23
	} else {
		goto L44
	}
L8:
	;
	v115 = v110
	v116 = v110
	goto L6
L9:
	;
	v110 = v34
	goto L8
L10:
	;
	if v55 == int32(2276) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v59 = int32(1)
	if v55 <= int32(3830) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v53 == int32(705) {
		goto L9
	} else {
		goto L22
	}
L13:
	;
	switch v55 - int32(2277) {
	case 0, 6:
		v110 = v59
		goto L8
	case 1, 2, 3, 4, 5:
		goto L12
	default:
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(v55-int32(5077)) < base.Ui32(int32(4)) {
		v110 = v59
		goto L8
	} else {
		goto L19
	}
L16:
	;
	if v55 == int32(2776) {
		v110 = v59
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v55 != int32(3500) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v110 = v59
	goto L8
L19:
	;
	if base.Ui32(v55-int32(4537)) < base.Ui32(int32(2)) {
		v110 = v59
		goto L8
	} else {
		goto L20
	}
L20:
	;
	if v55 == int32(3831) {
		v110 = v59
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	v82 = F_find_coercion_pathway(m, v55, v53, l3, v24+int32(12))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v82 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v53 == int32(2249) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = F_typeOrDomainTypeRelid(m, v55)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v55 != int32(2287) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v88 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v105 = F_typeInheritsFrom(m, v53, v55)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L23
	} else {
		goto L42
	}
L32:
	;
	if v55 != int32(2249) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v98 = F_get_element_type(m, v53)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L23
	} else {
		goto L38
	}
L35:
	;
	v94 = F_typeOrDomainTypeRelid(m, v53)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	if v94 == int32(0) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	goto L9
L38:
	;
	if v98 == int32(0) {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v102 = F_typeOrDomainTypeRelid(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	if v102 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	if v105 == int32(0) {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	goto L9
L44:
	;
	if v111 == int32(0) {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v115 = v34
	v116 = v34
	goto L6
L46:
	;
	goto L5
L47:
	;
	v124 = int32(0)
	v130 = m.G0
	v132 = v130 - int32(416)
	m.G0 = v132
	if l0 <= v124 {
		v395 = int32(1)
		goto L48
	} else {
		goto L49
	}
L48:
	;
	m.G0 = v132 + int32(416)
	if v395 != 0 {
		v452 = v26
		goto L1
	} else {
		goto L179
	}
L49:
	;
	v139 = v124
	v140 = v124
	v141 = v124
	v142 = int32(0)
	v144 = v124
	v145 = v124
	v146 = v124
	v147 = v5
	v149 = v5
	v151 = v5
	v154 = v5
	v155 = v5
	v156 = v5
	goto L50
L50:
	;
	v160 = v142 << (uint(int32(2)) % 32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1+v160)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+l2)))
	if v164 <= int32(3830) {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	if v252 == int32(0) {
		v279 = v250
		goto L116
	} else {
		goto L117
	}
L52:
	;
	v264 = v142 + int32(1)
	if v264 != l0 {
		v139 = v250
		v140 = v251
		v141 = v252
		v142 = v264
		v144 = v254
		v145 = v255
		v146 = v256
		v147 = v257
		v149 = v258
		v151 = v259
		v154 = v260
		v155 = v261
		v156 = v262
		goto L50
	} else {
		goto L115
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132+v147<<(uint(int32(2))%32)))) = v238
	v250 = v139
	v251 = v140
	v252 = v141
	v254 = v240
	v255 = v145
	v256 = v146
	v257 = v147 + int32(1)
	v258 = v241
	v259 = v151
	v260 = v242
	v261 = v155
	v262 = v156
	goto L52
L54:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L107
	}
L55:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L99
	}
L56:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L95
	}
L57:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v210
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L94
	}
L58:
	;
	v210 = int32(1)
	goto L57
L59:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L88
	}
L60:
	;
	if v187 == v140 {
		v250 = v139
		v251 = v187
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L87
	}
L61:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L81
	}
L62:
	;
	if v162 == v139 {
		v250 = v162
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v175
		v262 = v176
		goto L52
	} else {
		goto L80
	}
L63:
	;
	switch v164 - int32(2277) {
	case 0:
		goto L61
	case 1, 2, 3, 4, 5:
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	case 6:
		v175 = v155
		v176 = v156
		goto L66
	default:
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	switch v164 - int32(5077) {
	case 0:
		v210 = v154
		goto L57
	case 1:
		goto L56
	case 2:
		goto L58
	case 3:
		goto L55
	default:
		goto L74
	}
L66:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v175
		v262 = v176
		goto L52
	} else {
		goto L72
	}
L67:
	;
	if v164 == int32(2776) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v175 = v155
	v176 = int32(1)
	goto L66
L69:
	;
	goto L70
L70:
	;
	if v164 != int32(3500) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L71
	}
L71:
	;
	v175 = int32(1)
	v176 = v156
	goto L66
L72:
	;
	if v139 != 0 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	v250 = v162
	v251 = v140
	v252 = v141
	v254 = v144
	v255 = v145
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v151
	v260 = v154
	v261 = v175
	v262 = v176
	goto L52
L74:
	;
	switch v164 - int32(4537) {
	case 0:
		goto L59
	case 1:
		goto L54
	default:
		goto L75
	}
L75:
	;
	if v164 != int32(3831) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L76
	}
L76:
	;
	if v162 == int32(705) {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L77
	}
L77:
	;
	v187 = F_getBaseType(m, v162)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L78
	}
L78:
	;
	if v140 != 0 {
		goto L60
	} else {
		goto L79
	}
L79:
	;
	v250 = v139
	v251 = v187
	v252 = v141
	v254 = v144
	v255 = v145
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v151
	v260 = v154
	v261 = v155
	v262 = v156
	goto L52
L80:
	;
	v395 = int32(0)
	goto L48
L81:
	;
	v193 = F_getBaseType(m, v162)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	if v141 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v250 = v139
	v251 = v140
	v252 = v193
	v254 = v144
	v255 = v145
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v151
	v260 = v154
	v261 = v155
	v262 = v156
	goto L52
L84:
	;
	goto L85
L85:
	;
	if v193 == v141 {
		v250 = v139
		v251 = v140
		v252 = v193
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L86
	}
L86:
	;
	v395 = int32(0)
	goto L48
L87:
	;
	v395 = int32(0)
	goto L48
L88:
	;
	v203 = F_getBaseType(m, v162)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L23
	} else {
		goto L89
	}
L89:
	;
	if v145 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v250 = v139
	v251 = v140
	v252 = v141
	v254 = v144
	v255 = v203
	v256 = v146
	v257 = v147
	v258 = v149
	v259 = v151
	v260 = v154
	v261 = v155
	v262 = v156
	goto L52
L91:
	;
	goto L92
L92:
	;
	if v203 == v145 {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v203
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L93
	}
L93:
	;
	v395 = int32(0)
	goto L48
L94:
	;
	v238 = v162
	v240 = v144
	v241 = v149
	v242 = v210
	goto L53
L95:
	;
	v215 = F_getBaseType(m, v162)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L23
	} else {
		goto L96
	}
L96:
	;
	v217 = F_get_element_type(m, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L97
	}
L97:
	;
	if v217 != 0 {
		v238 = v217
		v240 = v144
		v241 = v149
		v242 = v154
		goto L53
	} else {
		goto L98
	}
L98:
	;
	v395 = int32(0)
	goto L48
L99:
	;
	v222 = F_getBaseType(m, v162)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L23
	} else {
		goto L100
	}
L100:
	;
	if v144 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v222 == v144 {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v226 = F_get_range_subtype(m, v222)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L23
	} else {
		goto L105
	}
L104:
	;
	v395 = int32(0)
	goto L48
L105:
	;
	if v226 != 0 {
		v238 = v226
		v240 = v222
		v241 = v226
		v242 = v154
		goto L53
	} else {
		goto L106
	}
L106:
	;
	v395 = int32(0)
	goto L48
L107:
	;
	v231 = F_getBaseType(m, v162)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L23
	} else {
		goto L108
	}
L108:
	;
	if v146 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v231 == v146 {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v146
		v257 = v147
		v258 = v149
		v259 = v151
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v235 = F_get_multirange_range(m, v231)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L23
	} else {
		goto L113
	}
L112:
	;
	v395 = int32(0)
	goto L48
L113:
	;
	if v235 != 0 {
		v250 = v139
		v251 = v140
		v252 = v141
		v254 = v144
		v255 = v145
		v256 = v231
		v257 = v147
		v258 = v149
		v259 = v235
		v260 = v154
		v261 = v155
		v262 = v156
		goto L52
	} else {
		goto L114
	}
L114:
	;
	v395 = int32(0)
	goto L48
L115:
	;
	goto L51
L116:
	;
	if v255 != 0 {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	if v252 == int32(2277) {
		v279 = v250
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v270 = int32(0)
	v271 = F_get_element_type(m, v252)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L23
	} else {
		goto L119
	}
L119:
	;
	if v271 == int32(0) {
		v395 = v270
		goto L48
	} else {
		goto L120
	}
L120:
	;
	if v250 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v279 = v271
	goto L116
L122:
	;
	goto L123
L123:
	;
	if v271 != v250 {
		v395 = v270
		goto L48
	} else {
		goto L124
	}
L124:
	;
	v279 = v250
	goto L116
L125:
	;
	if v262 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L126:
	;
	v294 = int32(0)
	v295 = F_get_range_subtype(m, v291)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L23
	} else {
		goto L138
	}
L127:
	;
	if base.B2i32(v282 == v251) == int32(0) {
		v395 = v281
		goto L48
	} else {
		goto L137
	}
L128:
	;
	v281 = int32(0)
	v282 = F_get_multirange_range(m, v255)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L23
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v251 != 0 {
		v291 = v251
		goto L126
	} else {
		goto L136
	}
L131:
	;
	if v282 == int32(0) {
		v395 = v281
		goto L48
	} else {
		goto L132
	}
L132:
	;
	if v251 != 0 {
		goto L127
	} else {
		goto L133
	}
L133:
	;
	v286 = F_get_range_subtype(m, v282)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L23
	} else {
		goto L134
	}
L134:
	;
	if v286 != 0 {
		v291 = v282
		goto L126
	} else {
		goto L135
	}
L135:
	;
	v395 = v281
	goto L48
L136:
	;
	v303 = v279
	goto L125
L137:
	;
	v291 = v251
	goto L126
L138:
	;
	if v295 == int32(0) {
		v395 = v294
		goto L48
	} else {
		goto L139
	}
L139:
	;
	if v279 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v303 = v295
	goto L125
L141:
	;
	goto L142
L142:
	;
	if v295 != v279 {
		v395 = v294
		goto L48
	} else {
		goto L143
	}
L143:
	;
	v303 = v279
	goto L125
L144:
	;
	if v261 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v308 = F_get_base_element_type(m, v303)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L23
	} else {
		goto L146
	}
L146:
	;
	if v308 == int32(0) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v395 = int32(0)
	goto L48
L148:
	;
	if v256 == int32(0) {
		v333 = v257
		v334 = v258
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v315 = F_type_is_enum(m, v303)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L23
	} else {
		goto L150
	}
L150:
	;
	if v315 != 0 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v395 = int32(0)
	goto L48
L152:
	;
	if v333 <= int32(0) {
		v395 = int32(1)
		goto L48
	} else {
		goto L162
	}
L153:
	;
	if v254 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if v254 == v259 {
		v333 = v257
		v334 = v258
		goto L152
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v322 = F_get_range_subtype(m, v259)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L23
	} else {
		goto L158
	}
L157:
	;
	v395 = int32(0)
	goto L48
L158:
	;
	if v322 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v395 = int32(0)
	goto L48
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132+v257<<(uint(int32(2))%32)))) = v322
	v333 = v257 + int32(1)
	v334 = v322
	goto L152
L162:
	;
	v339 = F_select_common_type_from_oids(m, v333, v132, int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L23
	} else {
		goto L163
	}
L163:
	;
	if v339 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v395 = int32(0)
	goto L48
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+412)) = v339
	v350 = int32(0)
	goto L168
L167:
	;
	if v260 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	v374 = F_can_coerce_type(m, int32(1), v132+v350<<(uint(int32(2))%32), v132+int32(412), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L23
	} else {
		goto L170
	}
L169:
	;
	v395 = int32(0)
	goto L48
L170:
	;
	if v374 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v377 = v350 + int32(1)
	if v333 != v377 {
		v350 = v377
		goto L168
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	goto L169
L174:
	;
	goto L167
L175:
	;
	v395 = base.B2i32(v334 == int32(0)) | base.B2i32(v339 == v334)
	goto L48
L176:
	;
	v382 = F_get_base_element_type(m, v339)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L23
	} else {
		goto L177
	}
L177:
	;
	if v382 == int32(0) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v395 = int32(0)
	goto L48
L179:
	;
	goto L3
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
							F_errmsg(m, int32(304281), v9+int32(32))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(499145), int32(3094), int32(242145))
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
								F_errmsg(m, int32(189804), v9+int32(16))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(499145), int32(3099), int32(242145))
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
				F_errmsg_internal(m, int32(50330), v9)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(499145), int32(3087), int32(242145))
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
					v24 = int32(4515392)
					v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				v23 = v16
				v24 = int32(4515392)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
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
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
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
					F_errmsg(m, int32(186545), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errhint(m, int32(578364), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errfinish(m, int32(494614), int32(3677), int32(15845))
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v32 int32
	_ = v32
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v6 = F_get_typtype(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		switch v6&int32(255) - int32(98) {
		case 0, 3, 11, 16:
			v32 = int32(0)
			return v32
		case 1:
			v32 = int32(1)
			return v32
		case 2:
			v14 = F_getBaseType(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
				v17 = F_get_typtype(m, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v17 == int32(99)) << (uint(int32(1)) % 32)
				}
			}
		default:
			return int32(4)
		case 14:
			switch l0 - int32(2249) {
			case 0:
				v32 = int32(3)
				return v32
			default:
				return int32(4)
			case 26, 29:
				v32 = int32(0)
				return v32
			}
		}
	}
}
func F_has_type_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v22 = F_text_to_cstring(m, v11)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_DirectFunctionCall1Coll(m, int32(1254), int32(0), v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
								F_errmsg(m, int32(72238), v8)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497467), int32(4575), int32(379309))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
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
						v46 = F_convert_any_priv_string(m, v16, int32(1657376))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = F_object_aclcheck(m, int32(1247), v24, v19, v46)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return base.B2i32(v48 == int32(0))
							}
						}
					}
				}
			}
		}
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
	F_errmsg_internal(m, int32(50330), v6)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(498780), int32(699), int32(435216))
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
