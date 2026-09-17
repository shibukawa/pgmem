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
func F_PlanCacheRelCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRelCallback[0]))
	if base.B2i32(v8 == v3)|base.B2i32(v8 == int32(_a_F_PlanCacheRelCallback_0)) == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRelCallback[1]))
	v219 = int32(0)
	if base.B2i32(v218 == v219)|base.B2i32(v218 == int32(_a_F_PlanCacheRelCallback_1)) == v219 {
		goto L72
	} else {
		goto L73
	}
L4:
	;
	v23 = v16 - int32(5)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v24 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v208 != int32(_a_F_PlanCacheRelCallback_0) {
		v16 = v208
		goto L4
	} else {
		goto L71
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(96))))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(36))))
	if l1 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	switch v33 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v37 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(92))))
	if v40 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v37 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v37 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 != int32(6) {
		v59 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v59&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v51 = v49 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v51) {
		v59 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v59 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v51)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v122 = v16 - int32(12)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123 == int32(0) {
		goto L6
	} else {
		goto L43
	}
L23:
	;
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v111)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(12))))
	if v115 == v111 {
		goto L22
	} else {
		goto L42
	}
L24:
	;
	if v67 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v70 = int32(0)
	if v67 == v70 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L22
L28:
	;
	if v108 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L29:
	;
	v108 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v76 <= int32(0) {
		v102 = v70
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v108 = v102
	goto L28
L33:
	;
	v79 = int32(0)
	if v79 < v76 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v82 = v76
	goto L36
L35:
	;
	v82 = v79
	goto L36
L36:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v85 = int32(0)
	goto L37
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v83+v85<<(uint(int32(2))%32))))
	v94 = base.B2i32(v93 == l1)
	if v93 == l1 {
		v102 = v94
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v102 = v94
	goto L32
L39:
	;
	v96 = v85 + int32(1)
	if v96 != v82 {
		v85 = v96
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
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+10)) = uint8(v118)
	goto L22
L43:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+10)))
	if v126 != int32(1) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v129 == int32(0) {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v133 <= v132 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v139 = v132
	goto L47
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v139<<(uint(int32(2))%32))))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v147 == int32(6) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L6
L49:
	;
	v199 = v139 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v199 < v200 {
		v139 = v199
		goto L47
	} else {
		goto L70
	}
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+76))
	if l1 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+10)) = uint8(v195)
	goto L6
L52:
	;
	if v150 != 0 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v153 = int32(0)
	if v150 == v153 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L49
L56:
	;
	if v191 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L57:
	;
	v191 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v159 <= int32(0) {
		v185 = v153
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v191 = v185
	goto L56
L61:
	;
	v162 = int32(0)
	if v162 < v159 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v165 = v159
	goto L64
L63:
	;
	v165 = v162
	goto L64
L64:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v168 = int32(0)
	goto L65
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v166+v168<<(uint(int32(2))%32))))
	v177 = base.B2i32(v176 == l1)
	if v176 == l1 {
		v185 = v177
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v185 = v177
	goto L60
L67:
	;
	v179 = v168 + int32(1)
	if v179 != v165 {
		v168 = v179
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
	v226 = v218
	goto L75
L73:
	;
	goto L74
L74:
	;
	return
L75:
	;
	v233 = v226 - int32(16)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v234 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L74
L77:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v286 != int32(_a_F_PlanCacheRelCallback_1) {
		v226 = v286
		goto L75
	} else {
		goto L98
	}
L78:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v226-int32(12))))
	if l1 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v283 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v283)
	goto L77
L80:
	;
	if v239 != 0 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v242 = int32(0)
	if v239 == v242 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L77
L84:
	;
	if v280 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L85:
	;
	v280 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v248 <= int32(0) {
		v274 = v242
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v280 = v274
	goto L84
L89:
	;
	v251 = int32(0)
	if v251 < v248 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v254 = v248
	goto L92
L91:
	;
	v254 = v251
	goto L92
L92:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v257 = int32(0)
	goto L93
L93:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255+v257<<(uint(int32(2))%32))))
	v266 = base.B2i32(v265 == l1)
	if v265 == l1 {
		v274 = v266
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v274 = v266
	goto L88
L95:
	;
	v268 = v257 + int32(1)
	if v268 != v254 {
		v257 = v268
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
	v41 = v23 + v31<<(uint(int32(4))%32) + v34*int32(100)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+111)))
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
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
