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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[108]))
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
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	v8 = *(*int32)(unsafe.Add(mBase, _consts[983]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[984]))
	if v215 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L2:
	;
	if v8 == int32(4083376) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v8
	goto L4
L4:
	;
	v20 = v13 - int32(5)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v21 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v205 != int32(4083376) {
		v13 = v205
		goto L4
	} else {
		goto L71
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(96))))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(36))))
	if l1 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	switch v30 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v34 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(92))))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v34 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v34 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v41 != int32(6) {
		v56 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v56&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v48 = v46 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v48) {
		v56 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v56 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v48)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v119 = v13 - int32(12)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v120 == int32(0) {
		goto L6
	} else {
		goto L43
	}
L23:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v108)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(12))))
	if v112 == v108 {
		goto L22
	} else {
		goto L42
	}
L24:
	;
	if v64 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v67 = int32(0)
	if v64 == v67 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L22
L28:
	;
	if v105 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L29:
	;
	v105 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v73 <= int32(0) {
		v98 = v67
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v105 = v98
	goto L28
L33:
	;
	v76 = int32(0)
	if v76 < v73 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v79 = v73
	goto L36
L35:
	;
	v79 = v76
	goto L36
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v82 = int32(0)
	goto L37
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80+v82<<(uint(int32(2))%32))))
	v91 = base.B2i32(v90 == l1)
	if v90 == l1 {
		v98 = v91
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v98 = v91
	goto L32
L39:
	;
	v93 = v82 + int32(1)
	if v93 != v79 {
		v82 = v93
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
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+10)) = uint8(v115)
	goto L22
L43:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+10)))
	if v123 != int32(1) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v126 == int32(0) {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v129 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v130 <= v129 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v136 = v129
	goto L47
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v136<<(uint(int32(2))%32))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v144 == int32(6) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L6
L49:
	;
	v196 = v136 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v196 < v197 {
		v136 = v196
		goto L47
	} else {
		goto L70
	}
L50:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	if l1 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+10)) = uint8(v192)
	goto L6
L52:
	;
	if v147 != 0 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v150 = int32(0)
	if v147 == v150 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L49
L56:
	;
	if v188 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L57:
	;
	v188 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v156 <= int32(0) {
		v181 = v150
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v188 = v181
	goto L56
L61:
	;
	v159 = int32(0)
	if v159 < v156 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v162 = v156
	goto L64
L63:
	;
	v162 = v159
	goto L64
L64:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v165 = int32(0)
	goto L65
L65:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163+v165<<(uint(int32(2))%32))))
	v174 = base.B2i32(v173 == l1)
	if v173 == l1 {
		v181 = v174
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v181 = v174
	goto L60
L67:
	;
	v176 = v165 + int32(1)
	if v176 != v162 {
		v165 = v176
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
	return
L73:
	;
	if v215 == int32(4083384) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v220 = v215
	goto L75
L75:
	;
	v227 = v220 - int32(16)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v228 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L72
L77:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v280 != int32(4083384) {
		v220 = v280
		goto L75
	} else {
		goto L98
	}
L78:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v220-int32(12))))
	if l1 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v277)
	goto L77
L80:
	;
	if v233 != 0 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v236 = int32(0)
	if v233 == v236 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L77
L84:
	;
	if v274 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L85:
	;
	v274 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v242 <= int32(0) {
		v267 = v236
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v274 = v267
	goto L84
L89:
	;
	v245 = int32(0)
	if v245 < v242 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v248 = v242
	goto L92
L91:
	;
	v248 = v245
	goto L92
L92:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v251 = int32(0)
	goto L93
L93:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v249+v251<<(uint(int32(2))%32))))
	v260 = base.B2i32(v259 == l1)
	if v259 == l1 {
		v267 = v260
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v267 = v260
	goto L88
L95:
	;
	v262 = v251 + int32(1)
	if v262 != v248 {
		v251 = v262
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(2249) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v124 | int32(16384)
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v11 | int32(98304)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	switch v15 - int32(99) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L1
	}
L5:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v91 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_load_typcache_tupdesc(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v24 = v18
	goto L9
L9:
	;
	F_IncrTupleDescRefCount(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v24 = v23
	goto L9
L12:
	;
	v27 = int32(491520)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 <= int32(0) {
		v81 = v27
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v86 | v81
	F_DecrTupleDescRefCount(m, v24)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L37
	}
L14:
	;
	v35 = v28
	v36 = v27
	v37 = int32(0)
	goto L15
L15:
	;
	v46 = v24 + int32(20) + v35<<(uint(int32(4))%32) + v37*int32(100)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+91)))
	if v47 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v81 = v75
	goto L13
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v54 = F_lookup_type_cache(m, v52, int32(16409))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v74 = v35
	v75 = v36
	goto L19
L19:
	;
	v77 = v37 + int32(1)
	if v77 < v74 {
		v35 = v74
		v36 = v75
		v37 = v77
		goto L15
	} else {
		goto L36
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v56 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v57 = v36
	goto L23
L22:
	;
	v57 = v36 & int32(-32769)
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+64))
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v61 = v57
	goto L26
L25:
	;
	v61 = v57 & int32(-65537)
	goto L26
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
	if v64 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v65 = v61
	goto L29
L28:
	;
	v65 = v61 & int32(-131073)
	goto L29
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+72))
	if v68 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v69 = v65
	goto L32
L31:
	;
	v69 = v65 & int32(-262145)
	goto L32
L32:
	;
	if v69 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v81 = int32(0)
	goto L13
L34:
	;
	goto L35
L35:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v74 = v73
	v75 = v69
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
	v98 = F_getBaseTypeAndTypmod(m, v8, l0+int32(304))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	v101 = v91
	goto L40
L40:
	;
	v103 = F_lookup_type_cache(m, v101, int32(16409))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v98
	v101 = v98
	goto L40
L42:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+13)))
	if v105 != int32(99) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v110 = v108 | int32(1048576)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v112&int32(491520) | v110
	goto L1
}
