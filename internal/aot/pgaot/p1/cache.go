package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_CacheInvalidateHeapTupleCommon(m, l0, l1, l2, int32(1795))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_PlanCacheObjectCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v251 int32
	_ = v251
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
	v4 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheObjectCallback[0]))
	if base.B2i32(v14 == v4)|base.B2i32(v14 == int32(_a_F_PlanCacheObjectCallback_0)) == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v14
	goto L4
L2:
	;
	goto L3
L3:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheObjectCallback[1]))
	v268 = int32(0)
	if base.B2i32(v267 == v268)|base.B2i32(v267 == int32(_a_F_PlanCacheObjectCallback_1)) == v268 {
		goto L65
	} else {
		goto L66
	}
L4:
	;
	v35 = v26 - int32(5)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v251 != int32(_a_F_PlanCacheObjectCallback_0) {
		v26 = v251
		goto L4
	} else {
		goto L64
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26-int32(96))))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v26-int32(32))))
	if v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	switch v45 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v49 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26-int32(92))))
	if v52 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v49 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v49 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v56 != int32(6) {
		v71 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v71&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v63 = v61 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v63) {
		v71 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v71 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v63)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v136 = v26 - int32(12)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v137 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v82 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v85 = int32(0)
	if v85 < v82 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v88 = v82
	goto L27
L26:
	;
	v88 = v85
	goto L27
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v94 = int32(0)
	goto L28
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89+v94<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v107 != l1 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L22
L30:
	;
	v121 = v94 + int32(1)
	if v121 != v88 {
		v94 = v121
		goto L28
	} else {
		goto L37
	}
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v109 != l2 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v111)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v26-int32(12))))
	if v115 == v111 {
		goto L22
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+10)) = uint8(v118)
	goto L22
L37:
	;
	goto L29
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+10)))
	if v140 != int32(1) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v143 == int32(0) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v147 <= v146 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v155 = v137
	v158 = v146
	goto L42
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v158<<(uint(int32(2))%32))))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v167 != int32(6) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L6
L44:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+92))
	if v170 == int32(0) {
		v213 = v155
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v228 = v155
	goto L46
L46:
	;
	v236 = v158 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v236 < v237 {
		v155 = v228
		v158 = v236
		goto L42
	} else {
		goto L63
	}
L47:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+10)))
	if v220 != int32(1) {
		goto L6
	} else {
		goto L62
	}
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v173 <= int32(0) {
		v213 = v155
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v176 = int32(0)
	if v176 < v173 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v179 = v173
	goto L52
L51:
	;
	v179 = v176
	goto L52
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v185 = int32(0)
	goto L53
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v180+v185<<(uint(int32(2))%32))))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v198 != l1 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v213 = v155
	goto L47
L55:
	;
	v206 = v185 + int32(1)
	if v206 != v179 {
		v185 = v206
		goto L53
	} else {
		goto L61
	}
L56:
	;
	if l2 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	if v200 != l2 {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+10)) = uint8(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v213 = v204
	goto L47
L60:
	;
	goto L59
L61:
	;
	goto L54
L62:
	;
	v228 = v213
	goto L46
L63:
	;
	goto L43
L64:
	;
	goto L5
L65:
	;
	v279 = v267
	goto L68
L66:
	;
	goto L67
L67:
	;
	return
L68:
	;
	v288 = v279 - int32(16)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v289 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v343 != int32(_a_F_PlanCacheObjectCallback_1) {
		v279 = v343
		goto L68
	} else {
		goto L86
	}
L71:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v279-int32(8))))
	if v294 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v297 <= int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v300 = int32(0)
	if v300 < v297 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v303 = v297
	goto L76
L75:
	;
	v303 = v300
	goto L76
L76:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v309 = int32(0)
	goto L77
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v304+v309<<(uint(int32(2))%32))))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v322 != l1 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L70
L79:
	;
	v329 = v309 + int32(1)
	if v329 != v303 {
		v309 = v329
		goto L77
	} else {
		goto L85
	}
L80:
	;
	if l2 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	if v324 != l2 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v288))) = uint8(v326)
	goto L70
L84:
	;
	goto L83
L85:
	;
	goto L78
L86:
	;
	goto L69
}
func F_PlanCacheSysCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheSysCallback[0]))
	if base.B2i32(v6 == v4)|base.B2i32(v6 == int32(_a_F_PlanCacheSysCallback_0)) == v4 {
		v15 = v6
		for {
			v19 = v15 - int32(5)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			if v20 != int32(1) {
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(96))))
				if v25 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					switch v29 - int32(137) {
					case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
						v33 = int32(1)
					default:
						v33 = int32(0)
					}
					if v33 != 0 {
						v61 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v61)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(12))))
						if v65 == v61 {
						} else {
							v68 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)) = uint8(v68)
						}
					} else {
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(92))))
					if v36 == int32(0) {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						if v40 != int32(6) {
							v55 = int32(1)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							v47 = v45 - int32(201)
							if base.Ui32(int32(41)) < base.Ui32(v47) {
								v55 = int32(0)
							} else {
								v55 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v47)) % 64)))
							}
						}
						if v55&int32(1) == int32(0) {
						} else {
							v61 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v61)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(12))))
							if v65 == v61 {
							} else {
								v68 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)) = uint8(v68)
							}
						}
					}
				}
			}
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v72 != int32(_a_F_PlanCacheSysCallback_0) {
				v15 = v72
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheSysCallback[1]))
	v81 = int32(0)
	if base.B2i32(v80 == v81)|base.B2i32(v80 == int32(_a_F_PlanCacheSysCallback_1)) == v81 {
		v89 = v80
		for {
			v94 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v89-int32(16)))) = uint8(v94)
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
			if v96 != int32(_a_F_PlanCacheSysCallback_1) {
				v89 = v96
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_cache_locale_time(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	v1 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(3104)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cache_locale_time[0])))
	if v23 == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L118
	}
L2:
	;
	v28 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v20 + int32(3104)
	return
L5:
	;
	return
L6:
	;
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	F_errmsg_internal(m, int32(_a_F_cache_locale_time_0), v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[2])) = int32(44)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[1]))
	v47 = int32(0)
	v52 = m.G0
	v54 = v52 - int32(32)
	m.G0 = v54
	v59 = v47
	goto L15
L10:
	;
	F_errfinish(m, int32(_a_F_cache_locale_time_1), int32(718), int32(_a_F_cache_locale_time_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v182 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L13:
	;
	m.G0 = v54 + int32(32)
	goto L12
L14:
	;
	v182 = int32(0)
	goto L13
L15:
	;
	v64 = v59 << (uint(int32(2)) % 32)
	v70 = int32(1) << (uint(v59) % 32) & int32(2147483647)
	if v70|int32(1) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v89 = F___loc_is_allocated(m, v47)
	mBase = m.M
	if v89 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64+(v54+int32(8))))) = v81
	if v81 == int32(-1) {
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v47+v64)))
	v81 = v77
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v79 = v46
	goto L23
L22:
	;
	v79 = int32(_a_F_cache_locale_time_3)
	goto L23
L23:
	;
	v80 = F___get_locale(m, v59, v79)
	mBase = m.M
	v81 = v80
	goto L17
L24:
	;
	v86 = v59 + int32(1)
	if v86 != int32(6) {
		v59 = v86
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	v92 = int32(_a_F_cache_locale_time_4)
	v94 = v54 + int32(8)
	v97 = F_memcmp(m, v94, v92, int32(24))
	mBase = m.M
	if v97 == int32(0) {
		v182 = v92
		goto L13
	} else {
		goto L29
	}
L27:
	;
	v161 = v47
	goto L28
L28:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v161)+16)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v161)+8)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = v170
	v182 = v161
	goto L13
L29:
	;
	v100 = int32(_a_F_cache_locale_time_5)
	v103 = F_memcmp(m, v94, v100, int32(24))
	mBase = m.M
	if v103 == int32(0) {
		v182 = v100
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v106 = int32(0)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cache_locale_time[3])))
	if v108 == v106 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = v106
	goto L34
L32:
	;
	goto L33
L33:
	;
	v141 = int32(_a_F_cache_locale_time_6)
	v143 = v54 + int32(8)
	v146 = F_memcmp(m, v143, v141, int32(24))
	mBase = m.M
	if v146 == int32(0) {
		v182 = v141
		goto L13
	} else {
		goto L37
	}
L34:
	;
	v121 = F___get_locale(m, v114, int32(_a_F_cache_locale_time_3))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v114<<(uint(int32(2))%32))+uint32(_c_F_cache_locale_time[4]))) = v121
	v124 = v114 + int32(1)
	if v124 != int32(6) {
		v114 = v124
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_cache_locale_time[3])) = uint8(v128)
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[5])) = v132
	goto L33
L36:
	;
	goto L35
L37:
	;
	v149 = int32(_a_F_cache_locale_time_7)
	v152 = F_memcmp(m, v143, v149, int32(24))
	mBase = m.M
	if v152 == int32(0) {
		v182 = v149
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v156 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v156 == int32(0) {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v161 = v156
	goto L28
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[1]))
	F_report_newlocale_failure(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v196 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v196
	v202 = F___gmtime_r(m, v20+int32(56), v20+int32(12))
	mBase = m.M
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[2])) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = v204
	v212 = F___strftime_l(m, v20-int32(-64), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v218 = F___strftime_l(m, v20+int32(144), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(1)
	v226 = F___strftime_l(m, v20+int32(224), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v232 = F___strftime_l(m, v20+int32(304), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(2)
	v240 = F___strftime_l(m, v20+int32(384), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v246 = F___strftime_l(m, v20+int32(464), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(3)
	v254 = F___strftime_l(m, v20+int32(544), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v260 = F___strftime_l(m, v20+int32(624), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(4)
	v268 = F___strftime_l(m, v20+int32(704), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v274 = F___strftime_l(m, v20+int32(784), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(5)
	v282 = F___strftime_l(m, v20+int32(864), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v288 = F___strftime_l(m, v20+int32(944), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(6)
	v296 = F___strftime_l(m, v20+int32(1024), int32(80), int32(_a_F_cache_locale_time_8), v202, v182)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v298 = int32(0)
	v326 = F___strftime_l(m, v20+int32(1104), int32(80), int32(_a_F_cache_locale_time_9), v202, v182)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v328 = int32(0)
	v350 = base.B2i32(v212 == v298) | (base.B2i32(v218 == v298) | (base.B2i32(v226 == v298) | (base.B2i32(v232 == v298) | (base.B2i32(v240 == v298) | (base.B2i32(v246 == v298) | (base.B2i32(v254 == v298) | (base.B2i32(v260 == v298) | (base.B2i32(v268 == v298) | (base.B2i32(v274 == v298) | (base.B2i32(v282 == v298) | (base.B2i32(v288 == v298) | (base.B2i32(v326 == v328) | base.B2i32(v296 == v328)))))))))))))
	v352 = v20 + int32(1184)
	v353 = v1
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v353
	v369 = F___strftime_l(m, v352, int32(80), int32(_a_F_cache_locale_time_10), v202, v182)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L60
	}
L59:
	;
	v389 = F___loc_is_allocated(m, v182)
	mBase = m.M
	if v389 != 0 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v371 = int32(80)
	v375 = F___strftime_l(m, v352+v371, v371, int32(_a_F_cache_locale_time_11), v202, v182)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v377 = int32(0)
	v382 = base.B2i32(v375 == v377) | base.B2i32(v369 == v377) | v350
	v386 = v353 + int32(1)
	if v386 != int32(12) {
		v350 = v382
		v352 = v352 + int32(160)
		v353 = v386
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	if v382&int32(1) != 0 {
		goto L1
	} else {
		goto L67
	}
L64:
	;
	F_emscripten_builtin_free(m, v182)
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[1]))
	v396 = F_pg_get_encoding_from_locale(m, v394, int32(1))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v398 = int32(0)
	if v398 < v396 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v401 = v396
	goto L71
L70:
	;
	v401 = v398
	goto L71
L71:
	;
	v405 = v20 - int32(-64)
	v406 = int32(0)
	goto L72
L72:
	;
	v422 = F_strlen(m, v405)
	mBase = m.M
	v423 = F_pg_any_to_server(m, v405, v422, v401)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L74
	}
L73:
	;
	v461 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[6])) = v461
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[7])) = v461
	v467 = v455
	v468 = v461
	goto L95
L74:
	;
	v426 = v406 << (uint(int32(2)) % 32)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+uint32(_c_F_cache_locale_time[8])))
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[9]))
	v430 = F_MemoryContextStrdup(m, v429, v423)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426)+uint32(_c_F_cache_locale_time[8]))) = v430
	if v427 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v427)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v405 != v423 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	F_pfree(m, v423)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L5
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v439 = v405 + int32(80)
	v440 = F_strlen(m, v439)
	mBase = m.M
	v441 = F_pg_any_to_server(m, v439, v440, v401)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v426)+uint32(_c_F_cache_locale_time[10])))
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[9]))
	v446 = F_MemoryContextStrdup(m, v445, v441)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426)+uint32(_c_F_cache_locale_time[10]))) = v446
	if v443 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_pfree(m, v443)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v441 != v439 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_pfree(m, v441)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v455 = v405 + int32(160)
	v457 = v406 + int32(1)
	if v457 != int32(7) {
		v405 = v455
		v406 = v457
		goto L72
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	goto L73
L95:
	;
	v484 = F_strlen(m, v467)
	mBase = m.M
	v485 = F_pg_any_to_server(m, v467, v484, v401)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L97
	}
L96:
	;
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_cache_locale_time[0])) = uint8(v523)
	v526 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[11])) = v526
	*(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[12])) = v526
	goto L4
L97:
	;
	v488 = v468 << (uint(int32(2)) % 32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_cache_locale_time[13])))
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[9]))
	v492 = F_MemoryContextStrdup(m, v491, v485)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_cache_locale_time[13]))) = v492
	if v489 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_pfree(m, v489)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v467 != v485 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L101
L103:
	;
	F_pfree(m, v485)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v501 = v467 + int32(80)
	v502 = F_strlen(m, v501)
	mBase = m.M
	v503 = F_pg_any_to_server(m, v501, v502, v401)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L5
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_cache_locale_time[14])))
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_cache_locale_time[9]))
	v508 = F_MemoryContextStrdup(m, v507, v503)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_cache_locale_time[14]))) = v508
	if v505 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_pfree(m, v505)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L5
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if v503 != v501 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	F_pfree(m, v503)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L5
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v519 = v468 + int32(1)
	if v519 != int32(12) {
		v467 = v467 + int32(160)
		v468 = v519
		goto L95
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	goto L96
L118:
	;
	F_errmsg_internal(m, int32(_a_F_cache_locale_time_12), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_cache_locale_time_1), int32(784), int32(_a_F_cache_locale_time_2))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cache_reduce_memory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int64
	_ = v145
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
	var v157 int64
	_ = v157
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var __phi211 int32
	_ = __phi211
	var v212 int32
	_ = v212
	var __phi212 int32
	_ = __phi212
	var v213 int32
	_ = v213
	var __phi213 int32
	_ = __phi213
	var v216 int32
	_ = v216
	var __phi216 int32
	_ = __phi216
	var v223 int32
	_ = v223
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v278 int32
	_ = v278
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if base.Ui64(v17) < base.Ui64(v16) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v16
	goto L3
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v22 = l0 + int32(180)
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v22
	goto L6
L6:
	;
	v34 = int32(1)
	v35 = v23
	v39 = int64(0)
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L58
	}
L8:
	;
	if v22 != v35 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v284 + v283
	return v278 & int32(1)
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	m.T0[v46].(func(*base.Module, int32))(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v278 = v34
	v283 = v39
	goto L12
L12:
	;
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	v52 = v35 - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = F_ExecStoreMinimalTuple(m, v53, v43, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	if v59 < v58 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	m.T0[v62].(func(*base.Module, int32, int32))(m, v43, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v66 = v42 << (uint(int32(3)) % 32)
	if v66 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	base.MemoryCopy(m, v67, v68, v66)
	goto L22
L21:
	;
	goto L22
L22:
	;
	if v42 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	base.MemoryCopy(m, v70, v71, v42)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v75 = v73 & int32(_a_F_cache_reduce_memory_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+6)) = uint16(v78)
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v81 = F_MemoizeHash_hash(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v85 = v81 & v84
	v88 = v83 + v85<<(uint(int32(4))%32)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+12)))
	if v89 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v94 = v85
	v95 = v88
	v97 = v84
	v100 = v83
	goto L29
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if v107 == v81 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v123 != v52 {
		goto L7
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v110 = F_MemoizeHash_equal(m, v80, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	v114 = v97
	v115 = v100
	goto L34
L34:
	;
	v118 = v114 & (v94 + int32(1))
	v121 = v115 + v118<<(uint(int32(4))%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)))
	if v122 != 0 {
		v94 = v118
		v95 = v121
		v97 = v114
		v100 = v115
		goto L29
	} else {
		goto L37
	}
L35:
	;
	if v110 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v114 = v113
	v115 = v112
	goto L34
L37:
	;
	goto L7
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	v130 = int64(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v131 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = v131
	v145 = v130
	goto L42
L40:
	;
	v171 = v130
	goto L41
L41:
	;
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v174
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+13)) = uint8(v174)
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v179 = v178 - v171
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v179 - base.I64_extend_i32_u(v183+int32(28))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v191 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v190 - v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+20))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v197 = int32(4)
	v201 = v195 & ((v95-v194)>>(uint(v197)%32) + v191)
	v204 = v194 + v201<<(uint(v197)%32)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+12)))
	if v205 != v191 {
		v243 = v95
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	F_pfree(m, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L44
	}
L43:
	;
	v171 = v157
	goto L41
L44:
	;
	F_pfree(m, v134)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v157 = v145 + base.I64_extend_i32_u(v149+int32(8))
	if v147 != 0 {
		v134 = v147
		v145 = v157
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v256 = v34 & base.B2i32(l1 != v52)
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+12)) = uint8(v257)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	F_pfree(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L55
	}
L48:
	;
	__phi211 = v95
	__phi212 = v204
	__phi213 = v195
	__phi216 = v201
	v211 = __phi211
	v212 = __phi212
	v213 = __phi213
	v216 = __phi216
	goto L49
L49:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	if v216 == v223&v213 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v243 = v212
	goto L47
L51:
	;
	v243 = v211
	goto L47
L52:
	;
	goto L53
L53:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v212)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+8)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v212)))
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = v228
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v189)+20))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v232 = int32(1)
	v234 = v231 & (v216 + v232)
	v237 = v230 + v234<<(uint(int32(4))%32)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+12)))
	if v238 == v232 {
		__phi211 = v212
		__phi212 = v237
		__phi213 = v231
		__phi216 = v234
		v211 = __phi211
		v212 = __phi212
		v213 = __phi213
		v216 = __phi216
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	F_pfree(m, v123)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v265 = v39 + int64(1)
	v266 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v267) < base.Ui64(v266) {
		v34 = v256
		v35 = v41
		v39 = v265
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v278 = v256
	v283 = v265
	goto L12
L58:
	;
	F_errmsg_internal(m, int32(_a_F_cache_reduce_memory_1), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_cache_reduce_memory_2), int32(485), int32(_a_F_cache_reduce_memory_3))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
