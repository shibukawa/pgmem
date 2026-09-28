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
func F_TypeCacheOpcCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = v7 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheOpcCallback[0]))
	F_hash_seq_init(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_hash_seq_search(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v7 + int32(32)
	return
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+312)) = v21 & int32(_a_F_TypeCacheOpcCallback_0)
	if base.B2i32(v21&int32(-1572866) == int32(0))|v21&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v47 = F_hash_seq_search(m, v7+int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
	if v32 != int32(99) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	if v35 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheOpcCallback[1]))
	v43 = F_hash_search(m, v37, v18+int32(16), int32(2), v7+int32(31))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	if v47 != 0 {
		v18 = v47
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v138
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(0) {
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
	v138 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_TypeIsVisibleExt_0), v11)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_TypeIsVisibleExt_1), int32(1068), int32(_a_F_TypeIsVisibleExt_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v40 = v36 + v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L40
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_TypeIsVisibleExt[0]))
	v46 = int32(0)
	if v45 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_TypeIsVisibleExt[0]))
	if v88 == int32(0) {
		v128 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v84 == int32(0) {
		v128 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v84 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v52 <= int32(0) {
		v78 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L18
L23:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v52
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v61 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = base.B2i32(v69 == v41)
	if v69 == v41 {
		v78 = v70
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L22
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v58 {
		v61 = v72
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
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v92 <= v91 {
		v128 = v3
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v98 = v91
	goto L34
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v98<<(uint(int32(2))%32))))
	v111 = base.B2i32(v110 == v41)
	if v110 == v41 {
		v128 = v111
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v128 = v111
	goto L14
L36:
	;
	v114 = int64(0)
	v116 = F_SearchSysCacheExists(m, int32(81), base.I64_extend_i32_u(v40+int32(4)), base.I64_extend_i32_u(v110), v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v116 != 0 {
		v128 = v111
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v119 = v98 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v119 < v120 {
		v98 = v119
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v138 = v128
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v453 int32
	_ = v453
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
	return v453
L2:
	;
	v453 = int32(1)
	goto L1
L3:
	;
	v36 = v5
	v39 = v5
	goto L4
L4:
	;
	v50 = v36 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+v50)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v50)))
	if v52 == v54 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v129 = int32(0)
	v136 = m.G0
	v138 = v136 - int32(416)
	m.G0 = v138
	if l0 <= v129 {
		v401 = int32(1)
		goto L47
	} else {
		goto L48
	}
L6:
	;
	goto L5
L7:
	;
	v122 = int32(1)
	v124 = v36 + v122
	if v124 != l0 {
		v36 = v124
		v39 = v122
		goto L4
	} else {
		goto L46
	}
L8:
	;
	v453 = int32(0)
	goto L1
L9:
	;
	v116 = v36 + int32(1)
	if v116 != l0 {
		v36 = v116
		goto L4
	} else {
		goto L44
	}
L10:
	;
	v56 = int32(2281)
	if base.B2i32(v52 == v56)|base.B2i32(v54 == v56) != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v54 <= int32(3830) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v52 == int32(705) {
		goto L9
	} else {
		goto L20
	}
L13:
	;
	switch v54 - int32(2276) {
	case 0:
		goto L9
	case 1, 7:
		goto L7
	case 2, 3, 4, 5, 6:
		goto L12
	default:
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.B2i32(base.Ui32(v54-int32(_a_F_can_coerce_type_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v54-int32(_a_F_can_coerce_type_1)) < base.Ui32(int32(2)))|base.B2i32(v54 == int32(3831)) != 0 {
		goto L7
	} else {
		goto L19
	}
L16:
	;
	if v54 == int32(2776) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v54 != int32(3500) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	goto L12
L20:
	;
	v85 = F_find_coercion_pathway(m, v54, v52, l3, v24+int32(12))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v85 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if v52 == int32(2249) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = F_typeOrDomainTypeRelid(m, v54)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v54 != int32(2287) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	if v91 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v108 = F_typeInheritsFrom(m, v52, v54)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L21
	} else {
		goto L40
	}
L30:
	;
	if v54 != int32(2249) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v101 = F_get_element_type(m, v52)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L36
	}
L33:
	;
	v97 = F_typeOrDomainTypeRelid(m, v52)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	if v97 == int32(0) {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L9
L36:
	;
	if v101 == int32(0) {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v105 = F_typeOrDomainTypeRelid(m, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	if v105 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	if v108 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v110 = F_typeIsOfTypedTable(m, v52, v54)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	if v110 == int32(0) {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	goto L9
L44:
	;
	if v39 == int32(0) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	goto L6
L46:
	;
	goto L6
L47:
	;
	m.G0 = v138 + int32(416)
	if v401 == int32(0) {
		v453 = v129
		goto L1
	} else {
		goto L176
	}
L48:
	;
	v145 = v129
	v146 = v129
	v147 = v129
	v148 = int32(0)
	v149 = v129
	v151 = v129
	v153 = v5
	v155 = v129
	v156 = v5
	v158 = v5
	v162 = v5
	v163 = v5
	v164 = v5
	goto L49
L49:
	;
	v166 = v148 << (uint(int32(2)) % 32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1+v166)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166+l2)))
	if v170 <= int32(3830) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	if base.B2i32(v259 == int32(0))|base.B2i32(v259 == int32(2277)) != 0 {
		v287 = v257
		goto L114
	} else {
		goto L115
	}
L51:
	;
	v271 = v148 + int32(1)
	if v271 != l0 {
		v145 = v257
		v146 = v258
		v147 = v259
		v148 = v271
		v149 = v260
		v151 = v262
		v153 = v263
		v155 = v264
		v156 = v265
		v158 = v266
		v162 = v267
		v163 = v268
		v164 = v269
		goto L49
	} else {
		goto L113
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138+v153<<(uint(int32(2))%32)))) = v245
	v257 = v145
	v258 = v146
	v259 = v147
	v260 = v149
	v262 = v247
	v263 = v153 + int32(1)
	v264 = v155
	v265 = v248
	v266 = v158
	v267 = v162
	v268 = v163
	v269 = v249
	goto L51
L53:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L105
	}
L54:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L97
	}
L55:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L93
	}
L56:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v217
		goto L51
	} else {
		goto L92
	}
L57:
	;
	v217 = int32(1)
	goto L56
L58:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L86
	}
L59:
	;
	if v194 == v146 {
		v257 = v145
		v258 = v194
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L85
	}
L60:
	;
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L79
	}
L61:
	;
	if v168 == v145 {
		v257 = v168
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v181
		v268 = v182
		v269 = v164
		goto L51
	} else {
		goto L78
	}
L62:
	;
	switch v170 - int32(2277) {
	case 0:
		goto L60
	case 1, 2, 3, 4, 5:
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	case 6:
		v181 = v162
		v182 = v163
		goto L65
	default:
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	switch v170 - int32(_a_F_can_coerce_type_0) {
	case 0:
		v217 = v164
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
	if v168 == int32(705) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v181
		v268 = v182
		v269 = v164
		goto L51
	} else {
		goto L71
	}
L66:
	;
	if v170 == int32(2776) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v181 = int32(1)
	v182 = v163
	goto L65
L68:
	;
	goto L69
L69:
	;
	if v170 != int32(3500) {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L70
	}
L70:
	;
	v181 = v162
	v182 = int32(1)
	goto L65
L71:
	;
	if v145 != 0 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	v257 = v168
	v258 = v146
	v259 = v147
	v260 = v149
	v262 = v151
	v263 = v153
	v264 = v155
	v265 = v156
	v266 = v158
	v267 = v181
	v268 = v182
	v269 = v164
	goto L51
L73:
	;
	switch v170 - int32(_a_F_can_coerce_type_1) {
	case 0:
		goto L58
	case 1:
		goto L53
	default:
		goto L74
	}
L74:
	;
	if base.B2i32(v168 == int32(705))|base.B2i32(v170 != int32(3831)) != 0 {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L75
	}
L75:
	;
	v194 = F_getBaseType(m, v168)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	if v146 != 0 {
		goto L59
	} else {
		goto L77
	}
L77:
	;
	v257 = v145
	v258 = v194
	v259 = v147
	v260 = v149
	v262 = v151
	v263 = v153
	v264 = v155
	v265 = v156
	v266 = v158
	v267 = v162
	v268 = v163
	v269 = v164
	goto L51
L78:
	;
	v401 = int32(0)
	goto L47
L79:
	;
	v200 = F_getBaseType(m, v168)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	if v147 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v257 = v145
	v258 = v146
	v259 = v200
	v260 = v149
	v262 = v151
	v263 = v153
	v264 = v155
	v265 = v156
	v266 = v158
	v267 = v162
	v268 = v163
	v269 = v164
	goto L51
L82:
	;
	goto L83
L83:
	;
	if v200 == v147 {
		v257 = v145
		v258 = v146
		v259 = v200
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L84
	}
L84:
	;
	v401 = int32(0)
	goto L47
L85:
	;
	v401 = int32(0)
	goto L47
L86:
	;
	v210 = F_getBaseType(m, v168)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L21
	} else {
		goto L87
	}
L87:
	;
	if v149 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v257 = v145
	v258 = v146
	v259 = v147
	v260 = v210
	v262 = v151
	v263 = v153
	v264 = v155
	v265 = v156
	v266 = v158
	v267 = v162
	v268 = v163
	v269 = v164
	goto L51
L89:
	;
	goto L90
L90:
	;
	if v210 == v149 {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v210
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L91
	}
L91:
	;
	v401 = int32(0)
	goto L47
L92:
	;
	v245 = v168
	v247 = v151
	v248 = v156
	v249 = v217
	goto L52
L93:
	;
	v222 = F_getBaseType(m, v168)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	v224 = F_get_element_type(m, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	if v224 != 0 {
		v245 = v224
		v247 = v151
		v248 = v156
		v249 = v164
		goto L52
	} else {
		goto L96
	}
L96:
	;
	v401 = int32(0)
	goto L47
L97:
	;
	v229 = F_getBaseType(m, v168)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L21
	} else {
		goto L98
	}
L98:
	;
	if v151 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v229 == v151 {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v233 = F_get_range_subtype(m, v229)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L21
	} else {
		goto L103
	}
L102:
	;
	v401 = int32(0)
	goto L47
L103:
	;
	if v233 != 0 {
		v245 = v233
		v247 = v229
		v248 = v233
		v249 = v164
		goto L52
	} else {
		goto L104
	}
L104:
	;
	v401 = int32(0)
	goto L47
L105:
	;
	v238 = F_getBaseType(m, v168)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L21
	} else {
		goto L106
	}
L106:
	;
	if v155 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v238 == v155 {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v155
		v265 = v156
		v266 = v158
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v242 = F_get_multirange_range(m, v238)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L21
	} else {
		goto L111
	}
L110:
	;
	v401 = int32(0)
	goto L47
L111:
	;
	if v242 != 0 {
		v257 = v145
		v258 = v146
		v259 = v147
		v260 = v149
		v262 = v151
		v263 = v153
		v264 = v238
		v265 = v156
		v266 = v242
		v267 = v162
		v268 = v163
		v269 = v164
		goto L51
	} else {
		goto L112
	}
L112:
	;
	v401 = int32(0)
	goto L47
L113:
	;
	goto L50
L114:
	;
	if v260 != 0 {
		goto L125
	} else {
		goto L126
	}
L115:
	;
	v278 = int32(0)
	v279 = F_get_element_type(m, v259)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L21
	} else {
		goto L116
	}
L116:
	;
	if v279 == int32(0) {
		v401 = v278
		goto L47
	} else {
		goto L117
	}
L117:
	;
	if v257 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v287 = v279
	goto L114
L119:
	;
	goto L120
L120:
	;
	if v279 != v257 {
		v401 = v278
		goto L47
	} else {
		goto L121
	}
L121:
	;
	v287 = v257
	goto L114
L122:
	;
	if v267 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L123:
	;
	v301 = int32(0)
	v302 = F_get_range_subtype(m, v299)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L21
	} else {
		goto L135
	}
L124:
	;
	if base.B2i32(v290 == v258) == int32(0) {
		v401 = v289
		goto L47
	} else {
		goto L134
	}
L125:
	;
	v289 = int32(0)
	v290 = F_get_multirange_range(m, v260)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L21
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v258 != 0 {
		v299 = v258
		goto L123
	} else {
		goto L133
	}
L128:
	;
	if v290 == int32(0) {
		v401 = v289
		goto L47
	} else {
		goto L129
	}
L129:
	;
	if v258 != 0 {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v294 = F_get_range_subtype(m, v290)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	if v294 != 0 {
		v299 = v290
		goto L123
	} else {
		goto L132
	}
L132:
	;
	v401 = v289
	goto L47
L133:
	;
	v310 = v287
	goto L122
L134:
	;
	v299 = v258
	goto L123
L135:
	;
	if v302 == int32(0) {
		v401 = v301
		goto L47
	} else {
		goto L136
	}
L136:
	;
	if v287 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v310 = v302
	goto L122
L138:
	;
	goto L139
L139:
	;
	if v302 != v287 {
		v401 = v301
		goto L47
	} else {
		goto L140
	}
L140:
	;
	v310 = v287
	goto L122
L141:
	;
	if v268 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v314 = F_get_base_element_type(m, v310)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	if v314 == int32(0) {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v401 = int32(0)
	goto L47
L145:
	;
	if v264 == int32(0) {
		v339 = v263
		v340 = v265
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v321 = F_type_is_enum(m, v310)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L21
	} else {
		goto L147
	}
L147:
	;
	if v321 != 0 {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v401 = int32(0)
	goto L47
L149:
	;
	if v339 <= int32(0) {
		v401 = int32(1)
		goto L47
	} else {
		goto L159
	}
L150:
	;
	if v262 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v262 == v266 {
		v339 = v263
		v340 = v265
		goto L149
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v328 = F_get_range_subtype(m, v266)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L21
	} else {
		goto L155
	}
L154:
	;
	v401 = int32(0)
	goto L47
L155:
	;
	if v328 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v401 = int32(0)
	goto L47
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138+v263<<(uint(int32(2))%32)))) = v328
	v339 = v263 + int32(1)
	v340 = v328
	goto L149
L159:
	;
	v345 = F_select_common_type_from_oids(m, v339, v138, int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L21
	} else {
		goto L160
	}
L160:
	;
	if v345 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v401 = int32(0)
	goto L47
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+412)) = v345
	v356 = int32(0)
	goto L165
L164:
	;
	if v269 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	v380 = F_can_coerce_type(m, int32(1), v138+v356<<(uint(int32(2))%32), v138+int32(412), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L21
	} else {
		goto L167
	}
L166:
	;
	v401 = int32(0)
	goto L47
L167:
	;
	if v380 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v383 = v356 + int32(1)
	if v339 != v383 {
		v356 = v383
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
	v401 = base.B2i32(v340 == int32(0)) | base.B2i32(v345 == v340)
	goto L47
L173:
	;
	v388 = F_get_base_element_type(m, v345)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L21
	} else {
		goto L174
	}
L174:
	;
	if v388 == int32(0) {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v401 = int32(0)
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+82)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = F_format_type_be(m, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
							F_errmsg(m, int32(_a_F_getTypeBinaryInputInfo_0), v9+int32(32))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3268), int32(_a_F_getTypeBinaryInputInfo_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = F_format_type_be(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77
								F_errmsg(m, int32(_a_F_getTypeBinaryInputInfo_3), v9+int32(16))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3273), int32(_a_F_getTypeBinaryInputInfo_2))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
					v27 = v25 + v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
					if v28 != 0 {
						v30 = v28
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v30 = v29
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
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
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeBinaryInputInfo_4), v9)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeBinaryInputInfo_1), int32(3261), int32(_a_F_getTypeBinaryInputInfo_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
func F_has_type_privilege_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14306(m, l0, int32(_a_F_has_type_privilege_name_0), int32(1247), int32(_a_F_has_type_privilege_name_1), int32(_a_F_has_type_privilege_name_2), int32(_a_F_has_type_privilege_name_3), int32(67137668), int32(1366))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = v10
	goto L7
L5:
	;
	v30 = l0
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
	v19 = v17 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+79)))
	if v20 != int32(100) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v30 = v23
	goto L6
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v28 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v23))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v28 != 0 {
		v15 = v28
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v30
	F_errmsg_internal(m, int32(_a_F_typeOrDomainTypeRelid_0), v6)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_typeOrDomainTypeRelid_1), int32(699), int32(_a_F_typeOrDomainTypeRelid_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	return v46
}
