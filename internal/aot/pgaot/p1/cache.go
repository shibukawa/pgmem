package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_CacheInvalidateHeapTupleCommon(m, l0, l1, l2, int32(1582))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_PlanCacheObjectCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheObjectCallback[0]))
	if base.B2i32(v13 == v4)|base.B2i32(v13 == int32(_a_F_PlanCacheObjectCallback_0)) == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v13
	goto L4
L2:
	;
	goto L3
L3:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheObjectCallback[1]))
	v258 = int32(0)
	if base.B2i32(v257 == v258)|base.B2i32(v257 == int32(_a_F_PlanCacheObjectCallback_1)) == v258 {
		goto L65
	} else {
		goto L66
	}
L4:
	;
	v33 = v24 - int32(5)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v242 != int32(_a_F_PlanCacheObjectCallback_0) {
		v24 = v242
		goto L4
	} else {
		goto L64
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(96))))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(32))))
	if v77 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	switch v43 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v47 = int32(1)
		goto L13
	default:
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(92))))
	if v50 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v47 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v47 = int32(0)
	goto L13
L15:
	;
	goto L6
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v54 != int32(6) {
		v69 = int32(1)
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v69&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v61 = v59 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v61) {
		v69 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v69 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v61)) % 64)))
	goto L18
L21:
	;
	goto L8
L22:
	;
	v132 = v24 - int32(12)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v133 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(0)
	if v83 < v80 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v86 = v80
	goto L27
L26:
	;
	v86 = v83
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v89 = int32(0)
	goto L28
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v87+v89<<(uint(int32(2))%32))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v104 != l1 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L22
L30:
	;
	v118 = v89 + int32(1)
	if v118 != v86 {
		v89 = v118
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	if v106 != l2 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v108)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(12))))
	if v112 == v108 {
		goto L22
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+10)) = uint8(v115)
	goto L22
L37:
	;
	goto L29
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+10)))
	if v136 != int32(1) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v139 == int32(0) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v143 <= v142 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v150 = v133
	v153 = v142
	goto L42
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v153<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v162 != int32(6) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L6
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+80))
	if v165 == int32(0) {
		v206 = v150
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v220 = v150
	goto L46
L46:
	;
	v228 = v153 + int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v228 < v229 {
		v150 = v220
		v153 = v228
		goto L42
	} else {
		goto L63
	}
L47:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
	if v213 != int32(1) {
		goto L6
	} else {
		goto L62
	}
L48:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v168 <= int32(0) {
		v206 = v150
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v171 = int32(0)
	if v171 < v168 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v174 = v168
	goto L52
L51:
	;
	v174 = v171
	goto L52
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v177 = int32(0)
	goto L53
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v175+v177<<(uint(int32(2))%32))))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v192 != l1 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v206 = v150
	goto L47
L55:
	;
	v200 = v177 + int32(1)
	if v200 != v174 {
		v177 = v200
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
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v194 != l2 {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+10)) = uint8(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v206 = v198
	goto L47
L60:
	;
	goto L59
L61:
	;
	goto L54
L62:
	;
	v220 = v206
	goto L46
L63:
	;
	goto L43
L64:
	;
	goto L5
L65:
	;
	v268 = v257
	goto L68
L66:
	;
	goto L67
L67:
	;
	return
L68:
	;
	v277 = v268 - int32(16)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v278 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v330 != int32(_a_F_PlanCacheObjectCallback_1) {
		v268 = v330
		goto L68
	} else {
		goto L86
	}
L71:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v268-int32(8))))
	if v283 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	if v286 <= int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v289 = int32(0)
	if v289 < v286 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v292 = v286
	goto L76
L75:
	;
	v292 = v289
	goto L76
L76:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v283)+12))
	v295 = int32(0)
	goto L77
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v293+v295<<(uint(int32(2))%32))))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v310 != l1 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L70
L79:
	;
	v317 = v295 + int32(1)
	if v317 != v292 {
		v295 = v317
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
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v312 != l2 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v314)
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
func F_PlanCacheSysCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheSysCallback[0]))
	v6 = int32(0)
	if base.B2i32(v5 == v6)|base.B2i32(v5 == int32(_a_F_PlanCacheSysCallback_0)) == v6 {
		v13 = v5
		for {
			v17 = v13 - int32(5)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v18 != int32(1) {
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(96))))
				if v23 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					switch v27 - int32(137) {
					case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
						v31 = int32(1)
					default:
						v31 = int32(0)
					}
					if v31 != 0 {
						v59 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v59)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(12))))
						if v63 == v59 {
						} else {
							v66 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)) = uint8(v66)
						}
					} else {
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(92))))
					if v34 == int32(0) {
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						if v38 != int32(6) {
							v53 = int32(1)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v45 = v43 - int32(201)
							if base.Ui32(int32(41)) < base.Ui32(v45) {
								v53 = int32(0)
							} else {
								v53 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v45)) % 64)))
							}
						}
						if v53&int32(1) == int32(0) {
						} else {
							v59 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v59)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(12))))
							if v63 == v59 {
							} else {
								v66 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)) = uint8(v66)
							}
						}
					}
				}
			}
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v70 != int32(_a_F_PlanCacheSysCallback_0) {
				v13 = v70
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheSysCallback[1]))
	v78 = int32(0)
	if base.B2i32(v77 == v78)|base.B2i32(v77 == int32(_a_F_PlanCacheSysCallback_1)) == v78 {
		v85 = v77
		for {
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v85-int32(16)))) = uint8(v90)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
			if v92 != int32(_a_F_PlanCacheSysCallback_1) {
				v85 = v92
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
	F_errfinish(m, int32(_a_F_cache_locale_time_1), int32(744), int32(_a_F_cache_locale_time_2))
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
	F_errfinish(m, int32(_a_F_cache_locale_time_1), int32(810), int32(_a_F_cache_locale_time_2))
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v143 int64
	_ = v143
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
	var v155 int64
	_ = v155
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var __phi209 int32
	_ = __phi209
	var v210 int32
	_ = v210
	var __phi210 int32
	_ = __phi210
	var v211 int32
	_ = v211
	var __phi211 int32
	_ = __phi211
	var v214 int32
	_ = v214
	var __phi214 int32
	_ = __phi214
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v276 int32
	_ = v276
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
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
	v306 = m.ExcPending
	if v306 != 0 {
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
	v282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v282 + v281
	return v276 & int32(1)
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
	v276 = v34
	v281 = v39
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
	F_slot_getsomeattrs_int(m, v43, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v64 = v42 << (uint(int32(2)) % 32)
	if v64 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	base.MemoryCopy(m, v65, v66, v64)
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	base.MemoryCopy(m, v68, v69, v42)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v73 = v71 & int32(_a_F_cache_reduce_memory_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v73)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+6)) = uint16(v76)
	goto L26
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v79 = F_MemoizeHash_hash(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v83 = v79 & v82
	v86 = v81 + v83<<(uint(int32(4))%32)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+12)))
	if v87 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v92 = v83
	v93 = v86
	v95 = v82
	v98 = v81
	goto L29
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v105 == v79 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v121 != v52 {
		goto L7
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v108 = F_MemoizeHash_equal(m, v78, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	v112 = v95
	v113 = v98
	goto L34
L34:
	;
	v116 = v112 & (v92 + int32(1))
	v119 = v113 + v116<<(uint(int32(4))%32)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)))
	if v120 != 0 {
		v92 = v116
		v93 = v119
		v95 = v112
		v98 = v113
		goto L29
	} else {
		goto L37
	}
L35:
	;
	if v108 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v112 = v111
	v113 = v110
	goto L34
L37:
	;
	goto L7
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v126
	v128 = int64(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v129 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v132 = v129
	v143 = v128
	goto L42
L40:
	;
	v169 = v128
	goto L41
L41:
	;
	v172 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v172
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+13)) = uint8(v172)
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v177 = v176 - v169
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v177 - base.I64_extend_i32_u(v181+int32(28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v189 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = v188 - v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v195 = int32(4)
	v199 = v193 & ((v93-v192)>>(uint(v195)%32) + v189)
	v202 = v192 + v199<<(uint(v195)%32)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+12)))
	if v203 != v189 {
		v241 = v93
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	F_pfree(m, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L44
	}
L43:
	;
	v169 = v155
	goto L41
L44:
	;
	F_pfree(m, v132)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v155 = v143 + base.I64_extend_i32_u(v147+int32(8))
	if v145 != 0 {
		v132 = v145
		v143 = v155
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v254 = v34 & base.B2i32(l1 != v52)
	v255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+12)) = uint8(v255)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	F_pfree(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L55
	}
L48:
	;
	__phi209 = v93
	__phi210 = v202
	__phi211 = v193
	__phi214 = v199
	v209 = __phi209
	v210 = __phi210
	v211 = __phi211
	v214 = __phi214
	goto L49
L49:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	if v214 == v221&v211 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v241 = v210
	goto L47
L51:
	;
	v241 = v209
	goto L47
L52:
	;
	goto L53
L53:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v210)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v209)+8)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v230 = int32(1)
	v232 = v229 & (v214 + v230)
	v235 = v228 + v232<<(uint(int32(4))%32)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+12)))
	if v236 == v230 {
		__phi209 = v210
		__phi210 = v235
		__phi211 = v229
		__phi214 = v232
		v209 = __phi209
		v210 = __phi210
		v211 = __phi211
		v214 = __phi214
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	F_pfree(m, v121)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v263 = v39 + int64(1)
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v265 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v265) < base.Ui64(v264) {
		v34 = v254
		v35 = v41
		v39 = v263
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v276 = v254
	v281 = v263
	goto L12
L58:
	;
	F_errmsg_internal(m, int32(_a_F_cache_reduce_memory_1), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_cache_reduce_memory_2), int32(484), int32(_a_F_cache_reduce_memory_3))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
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
