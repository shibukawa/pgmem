package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateRelcacheAll(m *base.Module) {
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v1 = F_PrepareInvalidationState(m)
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		v3 = int32(0)
		F_RegisterRelcacheInvalidation(m, v1, v3, v3)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_CacheInvalidateRelcacheByTuple(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+22)))
	v6 = v4 + v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+117)))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelcacheByTuple[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v11 = F_PrepareInvalidationState(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v7 != 0 {
			v14 = int32(0)
		} else {
			v14 = v9
		}
		F_RegisterRelcacheInvalidation(m, v11, v14, v10)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_PlanCacheRelCallback(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRelCallback[0]))
	if base.B2i32(v9 == v3)|base.B2i32(v9 == int32(_a_F_PlanCacheRelCallback_0)) == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v9
	goto L4
L2:
	;
	goto L3
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRelCallback[1]))
	v224 = int32(0)
	if base.B2i32(v223 == v224)|base.B2i32(v223 == int32(_a_F_PlanCacheRelCallback_1)) == v224 {
		goto L72
	} else {
		goto L73
	}
L4:
	;
	v25 = v20 - int32(5)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v26 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v212 != int32(_a_F_PlanCacheRelCallback_0) {
		v20 = v212
		goto L4
	} else {
		goto L71
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(96))))
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(36))))
	if l1 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	switch v35 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v39 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(92))))
	if v42 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v39 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v39 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v46 != int32(6) {
		v61 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v61&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v53 = v51 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v53) {
		v61 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v61 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v53)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v124 = v20 - int32(12)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v125 == int32(0) {
		goto L6
	} else {
		goto L43
	}
L23:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v113)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(12))))
	if v117 == v113 {
		goto L22
	} else {
		goto L42
	}
L24:
	;
	if v69 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v72 = int32(0)
	if v69 == v72 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L22
L28:
	;
	if v110 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L29:
	;
	v110 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v78 <= int32(0) {
		v104 = v72
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v110 = v104
	goto L28
L33:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v84 = v78
	goto L36
L35:
	;
	v84 = v81
	goto L36
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v87 = int32(0)
	goto L37
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85+v87<<(uint(int32(2))%32))))
	v96 = base.B2i32(v95 == l1)
	if v95 == l1 {
		v104 = v96
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v104 = v96
	goto L32
L39:
	;
	v98 = v87 + int32(1)
	if v98 != v84 {
		v87 = v98
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L23
L42:
	;
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)) = uint8(v120)
	goto L22
L43:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+10)))
	if v128 != int32(1) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v131 == int32(0) {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v134 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v135 <= v134 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v142 = v134
	goto L47
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v142<<(uint(int32(2))%32))))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v150 == int32(6) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L6
L49:
	;
	v202 = v142 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v202 < v203 {
		v142 = v202
		goto L47
	} else {
		goto L70
	}
L50:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+88))
	if l1 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)) = uint8(v198)
	goto L6
L52:
	;
	if v153 != 0 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v156 = int32(0)
	if v153 == v156 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L49
L56:
	;
	if v194 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L57:
	;
	v194 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v162 <= int32(0) {
		v188 = v156
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v194 = v188
	goto L56
L61:
	;
	v165 = int32(0)
	if v165 < v162 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v168 = v162
	goto L64
L63:
	;
	v168 = v165
	goto L64
L64:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v171 = int32(0)
	goto L65
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v169+v171<<(uint(int32(2))%32))))
	v180 = base.B2i32(v179 == l1)
	if v179 == l1 {
		v188 = v180
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v188 = v180
	goto L60
L67:
	;
	v182 = v171 + int32(1)
	if v182 != v168 {
		v171 = v182
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	goto L51
L70:
	;
	goto L48
L71:
	;
	goto L5
L72:
	;
	v234 = v223
	goto L75
L73:
	;
	goto L74
L74:
	;
	return
L75:
	;
	v239 = v234 - int32(16)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v240 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L74
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v292 != int32(_a_F_PlanCacheRelCallback_1) {
		v234 = v292
		goto L75
	} else {
		goto L98
	}
L78:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234-int32(12))))
	if l1 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v289)
	goto L77
L80:
	;
	if v245 != 0 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v248 = int32(0)
	if v245 == v248 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L77
L84:
	;
	if v286 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L85:
	;
	v286 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v254 <= int32(0) {
		v280 = v248
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v286 = v280
	goto L84
L89:
	;
	v257 = int32(0)
	if v257 < v254 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v260 = v254
	goto L92
L91:
	;
	v260 = v257
	goto L92
L92:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v263 = int32(0)
	goto L93
L93:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261+v263<<(uint(int32(2))%32))))
	v272 = base.B2i32(v271 == l1)
	if v271 == l1 {
		v280 = v272
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v280 = v272
	goto L88
L95:
	;
	v274 = v263 + int32(1)
	if v274 != v260 {
		v263 = v274
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	goto L79
L98:
	;
	goto L76
}
func F_cache_record_field_properties(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int32(2249) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v117 | int32(_a_F_cache_record_field_properties_0)
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v10 | int32(_a_F_cache_record_field_properties_1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	switch v14 - int32(99) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L1
	}
L5:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v85 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_load_typcache_tupdesc(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v23 = v17
	goto L9
L9:
	;
	F_IncrTupleDescRefCount(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v23 = v22
	goto L9
L12:
	;
	v26 = int32(_a_F_cache_record_field_properties_2)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 <= int32(0) {
		v76 = v26
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v80 | v76
	F_DecrTupleDescRefCount(m, v23)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L37
	}
L14:
	;
	v31 = v27
	v32 = v26
	v34 = int32(0)
	goto L15
L15:
	;
	v41 = v23 + v31<<(uint(int32(3))%32) + v34*int32(100)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+119)))
	if v42 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v76 = v70
	goto L13
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	v49 = F_lookup_type_cache(m, v47, int32(_a_F_cache_record_field_properties_3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v69 = v31
	v70 = v32
	goto L19
L19:
	;
	v72 = v34 + int32(1)
	if v72 < v69 {
		v31 = v69
		v32 = v70
		v34 = v72
		goto L15
	} else {
		goto L36
	}
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+52))
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = v32
	goto L23
L22:
	;
	v52 = v32 & int32(-32769)
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	if v55 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v56 = v52
	goto L26
L25:
	;
	v56 = v52 & int32(-65537)
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	if v59 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v60 = v56
	goto L29
L28:
	;
	v60 = v56 & int32(-131073)
	goto L29
L29:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+72))
	if v63 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v64 = v60
	goto L32
L31:
	;
	v64 = v60 & int32(-262145)
	goto L32
L32:
	;
	if v64 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v76 = int32(0)
	goto L13
L34:
	;
	goto L35
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v69 = v68
	v70 = v64
	goto L19
L36:
	;
	goto L16
L37:
	;
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+304)) = int32(-1)
	v92 = F_getBaseTypeAndTypmod(m, v7, l0+int32(304))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	v95 = v85
	goto L40
L40:
	;
	v97 = F_lookup_type_cache(m, v95, int32(_a_F_cache_record_field_properties_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v92
	v95 = v92
	goto L40
L42:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+13)))
	if v99 != int32(99) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v104 = v102 | int32(_a_F_cache_record_field_properties_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v106&int32(_a_F_cache_record_field_properties_2) | v104
	goto L1
}
