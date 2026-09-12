package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 float32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v270 float64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(0)
	if v27 == v28 {
		v44 = v28
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v44&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v31 == int32(0) {
		v44 = v28
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != int32(7) {
		v44 = v28
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 != int32(17) {
		v44 = v28
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	v44 = v40 ^ int32(1)
	goto L4
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = F_get_fn_opclass_options(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v55 = int32(95)
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v58 = int32(base.Ui32(v56) >> (uint(int32(2)) % 32))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v61 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v55 = v50<<(uint(int32(3))%32) - int32(1)
	goto L11
L13:
	;
	if base.Ui32(int32(10)) < base.Ui32(v25) {
		goto L50
	} else {
		goto L51
	}
L14:
	;
	v134 = int32(4)
	v138 = F_generate_trgm(m, v20+v134, v58-v134)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if int32(base.Ui32(v64)>>(uint(int32(2))%32)) != v58 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v58) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v129 != 0 {
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v129 = int32(0)
	goto L17
L19:
	;
	v103 = v98
	v104 = v99
	v105 = v100
	goto L29
L20:
	;
	if (v61|v20)&int32(3) != 0 {
		v98 = v61
		v99 = v20
		v100 = v58
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v91 = v61
	v92 = v20
	v93 = v58
	goto L22
L22:
	;
	if v93 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L23:
	;
	v75 = v61
	v76 = v20
	v77 = v58
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v80 != v81 {
		v98 = v75
		v99 = v76
		v100 = v77
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v91 = v86
	v92 = v84
	v93 = v88
	goto L22
L26:
	;
	v83 = int32(4)
	v84 = v76 + v83
	v86 = v75 + v83
	v88 = v77 - v83
	if base.Ui32(int32(3)) < base.Ui32(v88) {
		v75 = v86
		v76 = v84
		v77 = v88
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v98 = v91
	v99 = v92
	v100 = v93
	goto L19
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 == v109 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v129 = v108 - v109
	goto L17
L31:
	;
	v111 = int32(1)
	v116 = v105 - v111
	if v116 != 0 {
		v103 = v103 + v111
		v104 = v104 + v111
		v105 = v116
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L18
L35:
	;
	v166 = (v58 + int32(7)) & int32(2147483640)
	v167 = v61
	goto L13
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	v145 = (v58 + int32(7)) & int32(2147483640)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v150 = F_MemoryContextAlloc(m, v141, v145+int32(base.Ui32(v146)>>(uint(int32(2))%32)))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v58 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v157 = int32(base.Ui32(v155) >> (uint(int32(2)) % 32))
	if v157 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v152 = F__emscripten_memcpy_bulkmem(m, v150, v20, v58)
	mBase = m.M
	v153 = v152
	goto L41
L40:
	;
	v153 = v150
	goto L41
L41:
	;
	goto L38
L42:
	;
	if v61 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v158 = F__emscripten_memcpy_bulkmem(m, v153+v145, v138, v157)
	mBase = m.M
	goto L45
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	F_pfree(m, v61)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v153
	v166 = v145
	v167 = v150
	goto L13
L49:
	;
	goto L48
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L67
	}
L51:
	;
	if int32(1)<<(uint(v25)%32)&int32(1284) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v176 = v166 + v167
	v178 = base.B2i32(v25 != int32(2))
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+16)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v181)+12)))
	if v183&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v271 = F_Float8GetDatum(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L66
	}
L54:
	;
	v186 = F_cnt_sml(m, v176, v59, v178)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
	if v194&int32(4) != 0 {
		v270 = float64(0)
		goto L53
	} else {
		goto L58
	}
L57:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v16)+12)) = v186
	v270 = base.F64_sub(float64(1), base.F64_promote_f32(v186))
	goto L53
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v202 = int32(base.Ui32(v198)>>(uint(int32(2))%32)) - int32(5)
	if base.Ui32(v202) < base.Ui32(int32(3)) {
		v270 = float64(-1)
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v205 = int32(5)
	v209 = int32(1)
	v211 = base.I32_div_u_s(v202, int32(3))
	if base.Ui32(v211) <= base.Ui32(v209) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v214 = v209
	goto L62
L61:
	;
	v214 = v211
	goto L62
L62:
	;
	v215 = int32(0)
	v217 = v215
	v218 = v215
	goto L63
L63:
	;
	v230 = int32(3)
	v232 = v176 + v205 + v217*v230
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232))))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+2)))
	v238 = base.I32_rem_u_s(v233|v234<<(uint(int32(16))%32), v55)
	v242 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59+v205+int32(base.Ui32(v238)>>(uint(v230)%32))))))
	v246 = int32(1)
	v248 = int32(base.Ui32(v242)>>(uint(v238&int32(7))%32))&v246 + v218
	v250 = v217 + v246
	if v250 != v214 {
		v217 = v250
		v218 = v248
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v270 = base.F64_sub(float64(1), base.F64_div(base.F64_convert_i32_u(v248), base.F64_convert_i32_u(v211)))
	goto L53
L65:
	;
	goto L64
L66:
	;
	m.G0 = v16 + int32(16)
	return v271
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v25
	F_errmsg_internal(m, int32(479325), v16)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(489588), int32(525), int32(414843))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtrgm_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(287231)
			F_errmsg(m, int32(191561), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(489588), int32(61), int32(277928))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_gtrgm_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v349 int64
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v506 int64
	_ = v506
	var v508 int32
	_ = v508
	var v510 int64
	_ = v510
	var v512 int32
	_ = v512
	var v514 int64
	_ = v514
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v520 int32
	_ = v520
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v545 int64
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v565 int64
	_ = v565
	var v566 int32
	_ = v566
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v607 int64
	_ = v607
	var v609 int32
	_ = v609
	var v611 int64
	_ = v611
	var v613 int32
	_ = v613
	var v615 int64
	_ = v615
	var v617 int32
	_ = v617
	var v619 int64
	_ = v619
	var v621 int32
	_ = v621
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v646 int64
	_ = v646
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v666 int64
	_ = v666
	var v667 int32
	_ = v667
	var v670 int64
	_ = v670
	var v671 int64
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v765 int64
	_ = v765
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int64
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v808 int64
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int64
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int64
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int64
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int64
	_ = v850
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int64
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int64
	_ = v870
	var v871 int64
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v883 int64
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v892 int64
	_ = v892
	var v893 int32
	_ = v893
	var v896 int64
	_ = v896
	var v897 int32
	_ = v897
	var v900 int64
	_ = v900
	var v901 int32
	_ = v901
	var v904 int64
	_ = v904
	var v905 int32
	_ = v905
	var v908 int64
	_ = v908
	var v912 int64
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v923 int64
	_ = v923
	var v928 int64
	_ = v928
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int64
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v971 int64
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int64
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int64
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int64
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int64
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int64
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int64
	_ = v1033
	var v1034 int64
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int64
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1055 int64
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int64
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int64
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int64
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 int64
	_ = v1071
	var v1075 int64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1086 int64
	_ = v1086
	var v1091 int64
	_ = v1091
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 int64
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1134 int64
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int64
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int64
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int64
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1176 int64
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int64
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int64
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1209 int64
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1218 int64
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int64
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int64
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int64
	_ = v1234
	var v1238 int64
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1249 int64
	_ = v1249
	var v1266 int64
	_ = v1266
	var v1287 int64
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1328 int64
	_ = v1328
	v2 = int32(0)
	v17 = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(12)
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v54 = v50 + int32(5)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if v55&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_div(base.F32_convert_i32_s(v280+(base.I32_wrap_i64(v1328)^int32(-1))), base.F32_convert_i32_s(v280))
	return v18
L13:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_convert_i32_s(v1294)
	return v18
L14:
	;
	v1294 = v48<<(uint(v475)%32) + (base.I32_wrap_i64(v1287) ^ int32(-1))
	goto L13
L15:
	;
	v1294 = v48<<(uint(v576)%32) + (base.I32_wrap_i64(v1266) ^ int32(-1))
	goto L13
L16:
	;
	v1091 = int64(0)
	if v48 < int32(4) {
		v1170 = v54
		v1171 = v48
		v1176 = v1091
		goto L211
	} else {
		goto L212
	}
L17:
	;
	v928 = int64(0)
	if v48 < int32(4) {
		v1007 = v474
		v1008 = v48
		v1013 = v928
		goto L183
	} else {
		goto L184
	}
L18:
	;
	v765 = int64(0)
	if v48 < int32(4) {
		v844 = v260
		v845 = v48
		v850 = v765
		goto L155
	} else {
		goto L156
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v60 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	v64 = (v48 + int32(7)) & int32(-8)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	if v66 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v467 = int32(4)
	v468 = v55 & v467
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)))
	if v469&v467 != 0 {
		goto L104
	} else {
		goto L105
	}
L22:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)))
	if v276&int32(4) != 0 {
		goto L73
	} else {
		goto L74
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v139 = F_MemoryContextAlloc(m, v137, v60+v64)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L45
	}
L24:
	;
	v69 = v66 + v64
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if int32(base.Ui32(v70)>>(uint(int32(2))%32)) != v60 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v60) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v135 != 0 {
		goto L23
	} else {
		goto L44
	}
L27:
	;
	v135 = int32(0)
	goto L26
L28:
	;
	v109 = v104
	v110 = v105
	v111 = v106
	goto L38
L29:
	;
	if (v69|v49)&int32(3) != 0 {
		v104 = v69
		v105 = v49
		v106 = v60
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v97 = v69
	v98 = v49
	v99 = v60
	goto L31
L31:
	;
	if v99 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v81 = v69
	v82 = v49
	v83 = v60
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v86 != v87 {
		v104 = v81
		v105 = v82
		v106 = v83
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v97 = v92
	v98 = v90
	v99 = v94
	goto L31
L35:
	;
	v89 = int32(4)
	v90 = v82 + v89
	v92 = v81 + v89
	v94 = v83 - v89
	if base.Ui32(int32(3)) < base.Ui32(v94) {
		v81 = v92
		v82 = v90
		v83 = v94
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v104 = v97
	v105 = v98
	v106 = v99
	goto L28
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == v115 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v135 = v114 - v115
	goto L26
L40:
	;
	v117 = int32(1)
	v122 = v111 - v117
	if v122 != 0 {
		v109 = v109 + v117
		v110 = v110 + v117
		v111 = v122
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	v260 = v66
	goto L22
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v145 = int32(base.Ui32(v141)>>(uint(int32(2))%32)) - int32(5)
	v146 = int32(3)
	v147 = base.I32_div_u_s(v145, v146)
	if v139&v146 != 0 {
		v171 = v48
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v176 = int32(3)
	v179 = v48<<(uint(v176)%32) - int32(1)
	v181 = base.I32_div_s(v179, int32(8))
	v182 = v139 + v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v185 = v183 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v185)
	if base.Ui32(v176) <= base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v173 = F__emscripten_memset_bulkmem(m, v139, base.I32_extend8_s(int32(0)), v171)
	mBase = m.M
	goto L55
L48:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v48) {
		v171 = v48
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v48&int32(3) != 0 {
		v171 = v48
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v48 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v159 = v139 + v48
	v161 = v139 + int32(4)
	if base.Ui32(v161) < base.Ui32(v159) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v163 = v159
	goto L54
L53:
	;
	v163 = v161
	goto L54
L54:
	;
	v171 = (v139^int32(-1)+v163)&int32(-4) + int32(4)
	goto L47
L55:
	;
	goto L46
L56:
	;
	v191 = int32(1)
	if base.Ui32(v147) <= base.Ui32(v191) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v60 != 0 {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v194 = v191
	goto L61
L60:
	;
	v194 = v147
	goto L61
L61:
	;
	v198 = int32(0)
	goto L62
L62:
	;
	v213 = int32(3)
	v215 = v49 + int32(5) + v198*v213
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215))))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+2)))
	v221 = base.I32_rem_u_s(v216|v217<<(uint(int32(16))%32), v179)
	v224 = v139 + int32(base.Ui32(v221)>>(uint(v213)%32))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v226 = int32(1)
	v230 = v225 | v226<<(uint(v221&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v230)
	v233 = v198 + v226
	if v233 != v194 {
		v198 = v233
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L58
L64:
	;
	goto L63
L65:
	;
	if v66 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v253 = F__emscripten_memcpy_bulkmem(m, v139+v64, v49, v60)
	mBase = m.M
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	F_pfree(m, v66)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+16)) = v139
	v260 = v139
	goto L22
L72:
	;
	goto L71
L73:
	;
	v279 = int32(3)
	v280 = v48 << (uint(v279) % 32)
	if v279 < v48 {
		goto L18
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v48 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	if v48 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v1328 = v17
	goto L12
L78:
	;
	goto L79
L79:
	;
	v285 = int32(3)
	v286 = v48 & v285
	if base.Ui32(v48-int32(1)) < base.Ui32(v285) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v286 == int32(0) {
		v1328 = v349
		goto L12
	} else {
		goto L87
	}
L81:
	;
	v334 = v260
	v349 = v17
	goto L80
L82:
	;
	goto L83
L83:
	;
	v295 = v260
	v299 = int32(0)
	v310 = v17
	goto L84
L84:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v312)+uint32(_consts[1117]))))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	v318 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v316)+uint32(_consts[1117]))))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)))
	v322 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v320)+uint32(_consts[1117]))))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+3)))
	v326 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v324)+uint32(_consts[1117]))))
	v327 = v310 + v314 + v318 + v322 + v326
	v328 = int32(4)
	v329 = v295 + v328
	v331 = v299 + v328
	if v331 != v48&int32(-4) {
		v295 = v329
		v299 = v331
		v310 = v327
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v334 = v329
	v349 = v327
	goto L80
L86:
	;
	goto L85
L87:
	;
	v354 = v334
	v355 = int32(0)
	v369 = v349
	goto L88
L88:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	v373 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v370)+uint32(_consts[1117]))))
	v374 = v369 + v373
	v375 = int32(1)
	v378 = v355 + v375
	if v378 != v286 {
		v354 = v354 + v375
		v355 = v378
		v369 = v374
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v1328 = v374
	goto L12
L90:
	;
	goto L89
L91:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = float32(0)
	return v18
L92:
	;
	goto L93
L93:
	;
	v385 = int32(1)
	if v48 == v385 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v48&v385 != 0 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v389 = int32(0)
	v439 = v389
	v441 = v389
	goto L94
L96:
	;
	goto L97
L97:
	;
	v393 = int32(0)
	v398 = v393
	v400 = v393
	v402 = v393
	goto L98
L98:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v398))))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v54))))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415^v417)+uint32(_consts[1117]))))
	v423 = v398 | int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v423))))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v423))))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425^v427)+uint32(_consts[1117]))))
	v431 = v400 + v420 + v430
	v432 = int32(2)
	v433 = v398 + v432
	v435 = v402 + v432
	if v435 != v48&int32(2147483646) {
		v398 = v433
		v400 = v431
		v402 = v435
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v439 = v433
	v441 = v431
	goto L94
L100:
	;
	goto L99
L101:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260+v439))))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+v54))))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455^v457)+uint32(_consts[1117]))))
	v463 = v441 + v461
	goto L103
L102:
	;
	v463 = v441
	goto L103
L103:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v18))) = base.F32_convert_i32_u(v463)
	return v18
L104:
	;
	if v468 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if v468 != 0 {
		goto L125
	} else {
		goto L126
	}
L107:
	;
	v1294 = int32(0)
	goto L13
L108:
	;
	goto L109
L109:
	;
	v474 = v49 + int32(5)
	v475 = int32(3)
	if v475 < v48 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	if v48 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v1287 = v17
	goto L14
L112:
	;
	goto L113
L113:
	;
	v481 = int32(3)
	v482 = v48 & v481
	if base.Ui32(v48-int32(1)) < base.Ui32(v481) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v482 == int32(0) {
		v1287 = v545
		goto L14
	} else {
		goto L121
	}
L115:
	;
	v530 = v474
	v545 = v17
	goto L114
L116:
	;
	goto L117
L117:
	;
	v491 = v474
	v495 = int32(0)
	v506 = v17
	goto L118
L118:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	v510 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v508)+uint32(_consts[1117]))))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	v514 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v512)+uint32(_consts[1117]))))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+2)))
	v518 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v516)+uint32(_consts[1117]))))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+3)))
	v522 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v520)+uint32(_consts[1117]))))
	v523 = v506 + v510 + v514 + v518 + v522
	v524 = int32(4)
	v525 = v491 + v524
	v527 = v495 + v524
	if v527 != v48&int32(-4) {
		v491 = v525
		v495 = v527
		v506 = v523
		goto L118
	} else {
		goto L120
	}
L119:
	;
	v530 = v525
	v545 = v523
	goto L114
L120:
	;
	goto L119
L121:
	;
	v550 = v530
	v551 = int32(0)
	v565 = v545
	goto L122
L122:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v569 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v566)+uint32(_consts[1117]))))
	v570 = v565 + v569
	v571 = int32(1)
	v574 = v551 + v571
	if v574 != v482 {
		v550 = v550 + v571
		v551 = v574
		v565 = v570
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v1287 = v570
	goto L14
L124:
	;
	goto L123
L125:
	;
	v576 = int32(3)
	if v576 < v48 {
		goto L16
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v48 <= int32(0) {
		goto L143
	} else {
		goto L144
	}
L128:
	;
	if v48 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v1266 = v17
	goto L15
L130:
	;
	goto L131
L131:
	;
	v582 = int32(3)
	v583 = v48 & v582
	if base.Ui32(v48-int32(1)) < base.Ui32(v582) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v583 == int32(0) {
		v1266 = v646
		goto L15
	} else {
		goto L139
	}
L133:
	;
	v635 = v54
	v646 = v17
	goto L132
L134:
	;
	goto L135
L135:
	;
	v593 = int32(0)
	v596 = v54
	v607 = v17
	goto L136
L136:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	v611 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v609)+uint32(_consts[1117]))))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	v615 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v613)+uint32(_consts[1117]))))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+2)))
	v619 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v617)+uint32(_consts[1117]))))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+3)))
	v623 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v621)+uint32(_consts[1117]))))
	v624 = v607 + v611 + v615 + v619 + v623
	v625 = int32(4)
	v626 = v596 + v625
	v628 = v593 + v625
	if v628 != v48&int32(-4) {
		v593 = v628
		v596 = v626
		v607 = v624
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v635 = v626
	v646 = v624
	goto L132
L138:
	;
	goto L137
L139:
	;
	v651 = int32(0)
	v655 = v635
	v666 = v646
	goto L140
L140:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	v670 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v667)+uint32(_consts[1117]))))
	v671 = v666 + v670
	v672 = int32(1)
	v675 = v651 + v672
	if v675 != v583 {
		v651 = v675
		v655 = v655 + v672
		v666 = v671
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v1266 = v671
	goto L15
L142:
	;
	goto L141
L143:
	;
	v1294 = int32(0)
	goto L13
L144:
	;
	goto L145
L145:
	;
	v681 = v49 + int32(5)
	v682 = int32(1)
	if v48 == v682 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v48&v682 == int32(0) {
		v1294 = v735
		goto L13
	} else {
		goto L153
	}
L147:
	;
	v686 = int32(0)
	v734 = v686
	v735 = v686
	goto L146
L148:
	;
	goto L149
L149:
	;
	v690 = int32(0)
	v693 = v690
	v694 = v690
	v698 = v2
	goto L150
L150:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693+v54))))
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693+v681))))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711^v713)+uint32(_consts[1117]))))
	v719 = v693 | int32(1)
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681+v719))))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719+v54))))
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721^v723)+uint32(_consts[1117]))))
	v727 = v694 + v716 + v726
	v728 = int32(2)
	v729 = v693 + v728
	v731 = v698 + v728
	if v731 != v48&int32(2147483646) {
		v693 = v729
		v694 = v727
		v698 = v731
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v734 = v729
	v735 = v727
	goto L146
L152:
	;
	goto L151
L153:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734+v681))))
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734+v54))))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753^v755)+uint32(_consts[1117]))))
	v1294 = v735 + v759
	goto L13
L154:
	;
	v1328 = v923
	goto L12
L155:
	;
	if v845 == int32(0) {
		v923 = v850
		goto L169
	} else {
		goto L170
	}
L156:
	;
	if v260 != (v260+int32(3))&int32(-4) {
		v844 = v260
		v845 = v48
		v850 = v765
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v774 = v48 - int32(4)
	v778 = int32(base.Ui32(v774)>>(uint(int32(2))%32)) + int32(1)
	v780 = v778 & int32(3)
	if base.Ui32(v774) < base.Ui32(int32(12)) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v780 == int32(0) {
		v844 = v816
		v845 = v817
		v850 = v822
		goto L155
	} else {
		goto L165
	}
L159:
	;
	v816 = v260
	v817 = v48
	v822 = v765
	goto L158
L160:
	;
	goto L161
L161:
	;
	v786 = v260
	v787 = v48
	v788 = int32(0)
	v792 = v765
	goto L162
L162:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v786)+8))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v786)+4))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v808 = base.I64_extend_i32_u(base.I32_popcnt(v793)) + (base.I64_extend_i32_u(base.I32_popcnt(v796)) + (base.I64_extend_i32_u(base.I32_popcnt(v799)) + (v792 + base.I64_extend_i32_u(base.I32_popcnt(v802)))))
	v809 = int32(16)
	v810 = v787 - v809
	v812 = v786 + v809
	v814 = v788 + int32(4)
	if v814 != v778&int32(2147483644) {
		v786 = v812
		v787 = v810
		v788 = v814
		v792 = v808
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v816 = v812
	v817 = v810
	v822 = v808
	goto L158
L164:
	;
	goto L163
L165:
	;
	v827 = v817
	v828 = v816
	v829 = int32(0)
	v832 = v822
	goto L166
L166:
	;
	v833 = int32(4)
	v834 = v827 - v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	v838 = v832 + base.I64_extend_i32_u(base.I32_popcnt(v835))
	v840 = v828 + v833
	v842 = v829 + int32(1)
	if v842 != v780 {
		v827 = v834
		v828 = v840
		v829 = v842
		v832 = v838
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v844 = v840
	v845 = v834
	v850 = v838
	goto L155
L168:
	;
	goto L167
L169:
	;
	goto L154
L170:
	;
	v854 = v845 & int32(3)
	if v854 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if base.Ui32(v845) < base.Ui32(int32(4)) {
		v923 = v883
		goto L169
	} else {
		goto L178
	}
L172:
	;
	v877 = v844
	v879 = v845
	v883 = v850
	goto L171
L173:
	;
	goto L174
L174:
	;
	v860 = v845
	v861 = v844
	v862 = int32(0)
	v864 = v850
	goto L175
L175:
	;
	v865 = int32(1)
	v866 = v860 - v865
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	v870 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v867)+uint32(_consts[1117]))))
	v871 = v864 + v870
	v873 = v861 + v865
	v875 = v862 + v865
	if v875 != v854 {
		v860 = v866
		v861 = v873
		v862 = v875
		v864 = v871
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v877 = v873
	v879 = v866
	v883 = v871
	goto L171
L177:
	;
	goto L176
L178:
	;
	v886 = v877
	v888 = v879
	v892 = v883
	goto L179
L179:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+3)))
	v896 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v893)+uint32(_consts[1117]))))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+2)))
	v900 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v897)+uint32(_consts[1117]))))
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+1)))
	v904 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v901)+uint32(_consts[1117]))))
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	v908 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v905)+uint32(_consts[1117]))))
	v912 = v896 + (v900 + (v904 + (v892 + v908)))
	v913 = int32(4)
	v916 = v888 - v913
	if v916 != 0 {
		v886 = v886 + v913
		v888 = v916
		v892 = v912
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v923 = v912
	goto L169
L181:
	;
	goto L180
L182:
	;
	v1287 = v1086
	goto L14
L183:
	;
	if v1008 == int32(0) {
		v1086 = v1013
		goto L197
	} else {
		goto L198
	}
L184:
	;
	if v474 != (v49+int32(8))&int32(-4) {
		v1007 = v474
		v1008 = v48
		v1013 = v928
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v937 = v48 - int32(4)
	v941 = int32(base.Ui32(v937)>>(uint(int32(2))%32)) + int32(1)
	v943 = v941 & int32(3)
	if base.Ui32(v937) < base.Ui32(int32(12)) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	if v943 == int32(0) {
		v1007 = v979
		v1008 = v980
		v1013 = v985
		goto L183
	} else {
		goto L193
	}
L187:
	;
	v979 = v474
	v980 = v48
	v985 = v928
	goto L186
L188:
	;
	goto L189
L189:
	;
	v949 = v474
	v950 = v48
	v951 = int32(0)
	v955 = v928
	goto L190
L190:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v949)+12))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v949)+8))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v949)+4))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	v971 = base.I64_extend_i32_u(base.I32_popcnt(v956)) + (base.I64_extend_i32_u(base.I32_popcnt(v959)) + (base.I64_extend_i32_u(base.I32_popcnt(v962)) + (v955 + base.I64_extend_i32_u(base.I32_popcnt(v965)))))
	v972 = int32(16)
	v973 = v950 - v972
	v975 = v949 + v972
	v977 = v951 + int32(4)
	if v977 != v941&int32(2147483644) {
		v949 = v975
		v950 = v973
		v951 = v977
		v955 = v971
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v979 = v975
	v980 = v973
	v985 = v971
	goto L186
L192:
	;
	goto L191
L193:
	;
	v990 = v980
	v991 = v979
	v992 = int32(0)
	v995 = v985
	goto L194
L194:
	;
	v996 = int32(4)
	v997 = v990 - v996
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v1001 = v995 + base.I64_extend_i32_u(base.I32_popcnt(v998))
	v1003 = v991 + v996
	v1005 = v992 + int32(1)
	if v1005 != v943 {
		v990 = v997
		v991 = v1003
		v992 = v1005
		v995 = v1001
		goto L194
	} else {
		goto L196
	}
L195:
	;
	v1007 = v1003
	v1008 = v997
	v1013 = v1001
	goto L183
L196:
	;
	goto L195
L197:
	;
	goto L182
L198:
	;
	v1017 = v1008 & int32(3)
	if v1017 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if base.Ui32(v1008) < base.Ui32(int32(4)) {
		v1086 = v1046
		goto L197
	} else {
		goto L206
	}
L200:
	;
	v1040 = v1007
	v1042 = v1008
	v1046 = v1013
	goto L199
L201:
	;
	goto L202
L202:
	;
	v1023 = v1008
	v1024 = v1007
	v1025 = int32(0)
	v1027 = v1013
	goto L203
L203:
	;
	v1028 = int32(1)
	v1029 = v1023 - v1028
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
	v1033 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1030)+uint32(_consts[1117]))))
	v1034 = v1027 + v1033
	v1036 = v1024 + v1028
	v1038 = v1025 + v1028
	if v1038 != v1017 {
		v1023 = v1029
		v1024 = v1036
		v1025 = v1038
		v1027 = v1034
		goto L203
	} else {
		goto L205
	}
L204:
	;
	v1040 = v1036
	v1042 = v1029
	v1046 = v1034
	goto L199
L205:
	;
	goto L204
L206:
	;
	v1049 = v1040
	v1051 = v1042
	v1055 = v1046
	goto L207
L207:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+3)))
	v1059 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+uint32(_consts[1117]))))
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+2)))
	v1063 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1060)+uint32(_consts[1117]))))
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+1)))
	v1067 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1064)+uint32(_consts[1117]))))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049))))
	v1071 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+uint32(_consts[1117]))))
	v1075 = v1059 + (v1063 + (v1067 + (v1055 + v1071)))
	v1076 = int32(4)
	v1079 = v1051 - v1076
	if v1079 != 0 {
		v1049 = v1049 + v1076
		v1051 = v1079
		v1055 = v1075
		goto L207
	} else {
		goto L209
	}
L208:
	;
	v1086 = v1075
	goto L197
L209:
	;
	goto L208
L210:
	;
	v1266 = v1249
	goto L15
L211:
	;
	if v1171 == int32(0) {
		v1249 = v1176
		goto L225
	} else {
		goto L226
	}
L212:
	;
	if v54 != (v50+int32(8))&int32(-4) {
		v1170 = v54
		v1171 = v48
		v1176 = v1091
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1100 = v48 - int32(4)
	v1104 = int32(base.Ui32(v1100)>>(uint(int32(2))%32)) + int32(1)
	v1106 = v1104 & int32(3)
	if base.Ui32(v1100) < base.Ui32(int32(12)) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v1106 == int32(0) {
		v1170 = v1142
		v1171 = v1143
		v1176 = v1148
		goto L211
	} else {
		goto L221
	}
L215:
	;
	v1142 = v54
	v1143 = v48
	v1148 = v1091
	goto L214
L216:
	;
	goto L217
L217:
	;
	v1112 = v54
	v1113 = v48
	v1114 = int32(0)
	v1118 = v1091
	goto L218
L218:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+12))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+8))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1134 = base.I64_extend_i32_u(base.I32_popcnt(v1119)) + (base.I64_extend_i32_u(base.I32_popcnt(v1122)) + (base.I64_extend_i32_u(base.I32_popcnt(v1125)) + (v1118 + base.I64_extend_i32_u(base.I32_popcnt(v1128)))))
	v1135 = int32(16)
	v1136 = v1113 - v1135
	v1138 = v1112 + v1135
	v1140 = v1114 + int32(4)
	if v1140 != v1104&int32(2147483644) {
		v1112 = v1138
		v1113 = v1136
		v1114 = v1140
		v1118 = v1134
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v1142 = v1138
	v1143 = v1136
	v1148 = v1134
	goto L214
L220:
	;
	goto L219
L221:
	;
	v1153 = v1143
	v1154 = v1142
	v1155 = int32(0)
	v1158 = v1148
	goto L222
L222:
	;
	v1159 = int32(4)
	v1160 = v1153 - v1159
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1154)))
	v1164 = v1158 + base.I64_extend_i32_u(base.I32_popcnt(v1161))
	v1166 = v1154 + v1159
	v1168 = v1155 + int32(1)
	if v1168 != v1106 {
		v1153 = v1160
		v1154 = v1166
		v1155 = v1168
		v1158 = v1164
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v1170 = v1166
	v1171 = v1160
	v1176 = v1164
	goto L211
L224:
	;
	goto L223
L225:
	;
	goto L210
L226:
	;
	v1180 = v1171 & int32(3)
	if v1180 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	if base.Ui32(v1171) < base.Ui32(int32(4)) {
		v1249 = v1209
		goto L225
	} else {
		goto L234
	}
L228:
	;
	v1203 = v1170
	v1205 = v1171
	v1209 = v1176
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1186 = v1171
	v1187 = v1170
	v1188 = int32(0)
	v1190 = v1176
	goto L231
L231:
	;
	v1191 = int32(1)
	v1192 = v1186 - v1191
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187))))
	v1196 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1193)+uint32(_consts[1117]))))
	v1197 = v1190 + v1196
	v1199 = v1187 + v1191
	v1201 = v1188 + v1191
	if v1201 != v1180 {
		v1186 = v1192
		v1187 = v1199
		v1188 = v1201
		v1190 = v1197
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v1203 = v1199
	v1205 = v1192
	v1209 = v1197
	goto L227
L233:
	;
	goto L232
L234:
	;
	v1212 = v1203
	v1214 = v1205
	v1218 = v1209
	goto L235
L235:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212)+3)))
	v1222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+uint32(_consts[1117]))))
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212)+2)))
	v1226 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1223)+uint32(_consts[1117]))))
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212)+1)))
	v1230 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+uint32(_consts[1117]))))
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212))))
	v1234 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+uint32(_consts[1117]))))
	v1238 = v1222 + (v1226 + (v1230 + (v1218 + v1234)))
	v1239 = int32(4)
	v1242 = v1214 - v1239
	if v1242 != 0 {
		v1212 = v1212 + v1239
		v1214 = v1242
		v1218 = v1238
		goto L235
	} else {
		goto L237
	}
L236:
	;
	v1249 = v1238
	goto L225
L237:
	;
	goto L236
}
func F_gtrgm_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(12)
	goto L9
L9:
	;
	v50 = v48 + int32(5)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	v53 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v53)
	v56 = v50 << (uint(v53) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
	v62 = F__emscripten_memset_bulkmem(m, v51+int32(5), base.I32_extend8_s(int32(0)), v48)
	mBase = m.M
	goto L13
L13:
	;
	if v20 <= int32(0) {
		v299 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v299) >> (uint(int32(2)) % 32))
	return v51
L15:
	;
	v67 = int32(3)
	v68 = v48 & v67
	v84 = v2
	goto L16
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v84<<(uint(int32(4))%32))))
	v97 = v95 + int32(5)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)))
	if v98&int32(2) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v299 = v56
	goto L14
L18:
	;
	v286 = v84 + int32(1)
	if v286 != v20 {
		v84 = v286
		goto L16
	} else {
		goto L39
	}
L19:
	;
	v263 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v263
	v266 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v266)
	v299 = v263
	goto L14
L20:
	;
	if v98&int32(4) != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if base.Ui32(int32(base.Ui32(v210)>>(uint(int32(2))%32))-int32(5)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L35
	}
L23:
	;
	if v48 <= int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v105 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v110 = v105
	v116 = v105
	goto L28
L26:
	;
	v162 = v105
	goto L27
L27:
	;
	if v68 == int32(0) {
		goto L18
	} else {
		goto L31
	}
L28:
	;
	v127 = v110 + v62
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v97))))
	v131 = v128 | v130
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v131)
	v134 = v110 | int32(1)
	v135 = v62 + v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v97))))
	v139 = v136 | v138
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v139)
	v142 = v110 | int32(2)
	v143 = v62 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v97))))
	v147 = v144 | v146
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v147)
	v150 = v110 | int32(3)
	v151 = v62 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v97))))
	v155 = v152 | v154
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v155)
	v157 = int32(4)
	v158 = v110 + v157
	v160 = v116 + v157
	if v160 != v48&int32(2147483644) {
		v110 = v158
		v116 = v160
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v162 = v158
	goto L27
L30:
	;
	goto L29
L31:
	;
	v181 = v162
	v186 = v105
	goto L32
L32:
	;
	v198 = v181 + v62
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v97))))
	v202 = v199 | v201
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v202)
	v204 = int32(1)
	v207 = v186 + v204
	if v207 != v68 {
		v181 = v181 + v204
		v186 = v207
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L18
L34:
	;
	goto L33
L35:
	;
	v217 = int32(0)
	goto L36
L36:
	;
	v234 = int32(3)
	v236 = v97 + v217*v234
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236))))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+2)))
	v242 = base.I32_rem_u_s(v237|v238<<(uint(int32(16))%32), v48<<(uint(v67)%32)-int32(1))
	v245 = v62 + int32(base.Ui32(v242)>>(uint(v234)%32))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v247 = int32(1)
	v251 = v246 | v247<<(uint(v242&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v251)
	v254 = v217 + v247
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v261 = base.I32_div_u_s(int32(base.Ui32(v255)>>(uint(int32(2))%32))-int32(5), v234)
	if base.Ui32(v254) < base.Ui32(v261) {
		v217 = v254
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L18
L38:
	;
	goto L37
L39:
	;
	goto L17
}
