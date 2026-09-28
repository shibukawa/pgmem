package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstorePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	v14 = l1 << (uint(int32(3)) % 32)
	v16 = v14 + int32(8)
	v17 = l2 + v16
	v18 = F_palloc(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v23 = l1 | int32(-2147483648)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
		if l1 != 0 {
			v29 = v18 + int32(8)
			v32 = v29 + v14&int32(2147483640)
			if int32(0) < l1 {
				v37 = v29
				v38 = v32
				v45 = int32(0)
				for {
					v49 = l0 + v45*int32(20)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
					if v50 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						base.MemoryCopy(m, v38, v51, v50)
					} else {
					}
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
					v54 = v38 + v53
					v57 = (v54 - v32) & int32(1073741823)
					*(*int32)(unsafe.Add(mBase, uint32(v37))) = v57
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
					if v59 == int32(1) {
						v72 = v54
						v74 = v57 | int32(1073741824)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
						if v64 != 0 {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							base.MemoryCopy(m, v54, v65, v64)
						} else {
						}
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
						v68 = v54 + v67
						v72 = v68
						v74 = (v68 - v32) & int32(1073741823)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v74
					v79 = v45 + int32(1)
					if v79 != l1 {
						v37 = v37 + int32(8)
						v38 = v72
						v45 = v79
						continue
					} else {
						break
					}
					break
				}
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v84 = v81
				v85 = v72
			} else {
				v84 = v23
				v85 = v32
			}
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v94 | int32(-2147483648)
			if v84&int32(268435455) != l1 {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v23
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v16 - v32 + v85) << (uint(int32(2)) % 32)
		} else {
		}
		return v18
	}
}
func F_hstore_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_ge_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v6^int64(-1))>>(uint(int64(31))%64)) & int64(1)
	}
}
func F_hstore_hash_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_hstoreUpgrade(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(4)
		v10 = v5 + v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v15 = int32(base.Ui32(v11)>>(uint(int32(2))%32)) - v9
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = v15 - int32(1636608432)
		if v16 == int64(0) {
			v59 = v22
			v61 = v22
			v63 = v22
		} else {
			v26 = v22 + base.I32_wrap_i64(v16)
			v27 = v26 + v22
			v31 = int32(4)
			v33 = base.I32_wrap_i64(int64(base.Ui64(v16)>>(uint(int64(32))%64))) ^ base.I32_rotl(v22, v31)
			v37 = v26 - v33 ^ base.I32_rotl(v33, int32(6))
			v41 = v27 - v37 ^ base.I32_rotl(v37, int32(8))
			v42 = v27 + v33
			v43 = v37 + v42
			v44 = v41 + v43
			v48 = v42 - v41 ^ base.I32_rotl(v41, int32(16))
			v52 = v43 - v48 ^ base.I32_rotl(v48, int32(19))
			v57 = v44 + v48
			v59 = v57
			v61 = v44 - v52 ^ base.I32_rotl(v52, v31)
			v63 = v52 + v57
		}
		if v10&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v15) {
				v68 = v10
				v69 = v15
				v71 = v59
				v72 = v63
				v73 = v61
				for {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
					v76 = v75 + v72
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
					v80 = v79 + v73
					v82 = int32(4)
					v84 = v77 + v71 - v80 ^ base.I32_rotl(v80, v82)
					v88 = v76 - v84 ^ base.I32_rotl(v84, int32(6))
					v89 = v80 + v76
					v90 = v84 + v89
					v91 = v88 + v90
					v95 = v89 - v88 ^ base.I32_rotl(v88, int32(8))
					v99 = v90 - v95 ^ base.I32_rotl(v95, int32(16))
					v103 = v91 - v99 ^ base.I32_rotl(v99, int32(19))
					v104 = v95 + v91
					v105 = v99 + v104
					v106 = v103 + v105
					v110 = v104 - v103 ^ base.I32_rotl(v103, v82)
					v111 = int32(12)
					v112 = v68 + v111
					v114 = v69 - v111
					if base.Ui32(int32(11)) < base.Ui32(v114) {
						v68 = v112
						v69 = v114
						v71 = v105
						v72 = v106
						v73 = v110
						continue
					} else {
						break
					}
					break
				}
				v117 = v112
				v118 = v114
				v120 = v105
				v121 = v106
				v122 = v110
			} else {
				v117 = v10
				v118 = v15
				v120 = v59
				v121 = v63
				v122 = v61
			}
			switch v118 - int32(1) {
			case 0:
				v287 = v120
				v288 = v121
				v289 = v122
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 1:
				v280 = v120
				v281 = v121
				v282 = v122
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 2:
				v273 = v120
				v274 = v121
				v275 = v122
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 3:
				v267 = v121
				v268 = v122
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 4:
				v263 = v121
				v264 = v122
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 5:
				v257 = v121
				v258 = v122
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 6:
				v251 = v121
				v252 = v122
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
				v257 = v253<<(uint(int32(16))%32) + v251
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 7:
				v246 = v122
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+7)))
				v251 = v247<<(uint(int32(24))%32) + v121
				v252 = v246
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
				v257 = v253<<(uint(int32(16))%32) + v251
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 8:
				v241 = v122
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+8)))
				v246 = v242<<(uint(int32(8))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+7)))
				v251 = v247<<(uint(int32(24))%32) + v121
				v252 = v246
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
				v257 = v253<<(uint(int32(16))%32) + v251
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 9:
				v236 = v122
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+9)))
				v241 = v237<<(uint(int32(16))%32) + v236
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+8)))
				v246 = v242<<(uint(int32(8))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+7)))
				v251 = v247<<(uint(int32(24))%32) + v121
				v252 = v246
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
				v257 = v253<<(uint(int32(16))%32) + v251
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			case 10:
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+10)))
				v236 = v232<<(uint(int32(24))%32) + v122
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+9)))
				v241 = v237<<(uint(int32(16))%32) + v236
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+8)))
				v246 = v242<<(uint(int32(8))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+7)))
				v251 = v247<<(uint(int32(24))%32) + v121
				v252 = v246
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
				v257 = v253<<(uint(int32(16))%32) + v251
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)))
				v263 = v259<<(uint(int32(8))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
				v267 = v263 + v265
				v268 = v264
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)))
				v273 = v269<<(uint(int32(24))%32) + v120
				v274 = v267
				v275 = v268
				v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
				v280 = v276<<(uint(int32(16))%32) + v273
				v281 = v274
				v282 = v275
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
				v287 = v283<<(uint(int32(8))%32) + v280
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
				v295 = v287 + v290
				v296 = v288
				v297 = v289
			default:
				v295 = v120
				v296 = v121
				v297 = v122
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v15) {
				v128 = v10
				v129 = v15
				v131 = v59
				v132 = v63
				v133 = v61
				for {
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
					v136 = v135 + v132
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
					v140 = v139 + v133
					v142 = int32(4)
					v144 = v137 + v131 - v140 ^ base.I32_rotl(v140, v142)
					v148 = v136 - v144 ^ base.I32_rotl(v144, int32(6))
					v149 = v140 + v136
					v150 = v144 + v149
					v151 = v148 + v150
					v155 = v149 - v148 ^ base.I32_rotl(v148, int32(8))
					v159 = v150 - v155 ^ base.I32_rotl(v155, int32(16))
					v163 = v151 - v159 ^ base.I32_rotl(v159, int32(19))
					v164 = v155 + v151
					v165 = v159 + v164
					v166 = v163 + v165
					v170 = v164 - v163 ^ base.I32_rotl(v163, v142)
					v171 = int32(12)
					v172 = v128 + v171
					v174 = v129 - v171
					if base.Ui32(int32(11)) < base.Ui32(v174) {
						v128 = v172
						v129 = v174
						v131 = v165
						v132 = v166
						v133 = v170
						continue
					} else {
						break
					}
					break
				}
				v177 = v172
				v178 = v174
				v180 = v165
				v181 = v166
				v182 = v170
			} else {
				v177 = v10
				v178 = v15
				v180 = v59
				v181 = v63
				v182 = v61
			}
			switch v178 - int32(1) {
			case 0:
				v229 = v180
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
				v295 = v229 + v230
				v296 = v181
				v297 = v182
			case 1:
				v224 = v180
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
				v229 = v225<<(uint(int32(8))%32) + v224
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
				v295 = v229 + v230
				v296 = v181
				v297 = v182
			case 2:
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)))
				v224 = v220<<(uint(int32(16))%32) + v180
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
				v229 = v225<<(uint(int32(8))%32) + v224
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
				v295 = v229 + v230
				v296 = v181
				v297 = v182
			case 3:
				v217 = v181
				v218 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v295 = v218 + v180
				v296 = v217
				v297 = v182
			case 4:
				v214 = v181
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+4)))
				v217 = v214 + v215
				v218 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v295 = v218 + v180
				v296 = v217
				v297 = v182
			case 5:
				v209 = v181
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+5)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+4)))
				v217 = v214 + v215
				v218 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v295 = v218 + v180
				v296 = v217
				v297 = v182
			case 6:
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+6)))
				v209 = v205<<(uint(int32(16))%32) + v181
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+5)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+4)))
				v217 = v214 + v215
				v218 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v295 = v218 + v180
				v296 = v217
				v297 = v182
			case 7:
				v200 = v182
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
				v295 = v201 + v180
				v296 = v203 + v181
				v297 = v200
			case 8:
				v195 = v182
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
				v295 = v201 + v180
				v296 = v203 + v181
				v297 = v200
			case 9:
				v190 = v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
				v295 = v201 + v180
				v296 = v203 + v181
				v297 = v200
			case 10:
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+10)))
				v190 = v186<<(uint(int32(24))%32) + v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
				v295 = v201 + v180
				v296 = v203 + v181
				v297 = v200
			default:
				v295 = v180
				v296 = v181
				v297 = v182
			}
		}
		v300 = int32(14)
		v302 = v296 ^ v297 - base.I32_rotl(v296, v300)
		v306 = v302 ^ v295 - base.I32_rotl(v302, int32(11))
		v310 = v306 ^ v296 - base.I32_rotl(v306, int32(25))
		v314 = v310 ^ v302 - base.I32_rotl(v310, int32(16))
		v318 = v314 ^ v306 - base.I32_rotl(v314, int32(4))
		v322 = v318 ^ v310 - base.I32_rotl(v318, v300)
		v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v332 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v335 = m.ExcPending
			if v335 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v322)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v322^v314-base.I32_rotl(v322, int32(24)))
			}
		} else {
			return base.I64_extend_i32_u(v322)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v322^v314-base.I32_rotl(v322, int32(24)))
		}
	}
}
func F_hstore_populate_record(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v578 int32
	_ = v578
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int64
	_ = v621
	var v622 int32
	_ = v622
	var v644 int64
	_ = v644
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_get_fn_expr_argtype(m, v26, v2)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v24 + int32(32)
	return v644
L2:
	;
	v601 = F_heap_form_tuple(m, v68, v153, v155)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L5
	} else {
		goto L128
	}
L3:
	;
	v332 = v58 + int32(8)
	v340 = int32(0)
	goto L66
L4:
	;
	F_heap_deform_tuple(m, v24+int32(12), v68, v153, v155)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L64
	}
L5:
	;
	return int64(0)
L6:
	;
	v32 = F_type_is_rowtype(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v34 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L60
	}
L11:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v58 = F_hstoreUpgrade(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v37 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v45 = F_pg_detoast_datum(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	v53 = int32(-1)
	v54 = v2
	v55 = v28
	goto L11
L16:
	;
	goto L17
L17:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
	v644 = int64(0)
	goto L1
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v48 == int32(1) {
		v644 = base.I64_extend_i32_u(v45)
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v53 = v51
	v54 = v45
	v55 = v52
	goto L11
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v62 = v60 & int32(268435455)
	v63 = int32(0)
	if v62|base.B2i32(v54 == v63) == v63 {
		v644 = base.I64_extend_i32_u(v54)
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v68 = F_lookup_rowtype_tupdesc_domain(m, v55, v53)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v54 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v54
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v73
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)) = uint16(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(base.Ui32(v71) >> (uint(int32(2)) % 32))
	goto L25
L24:
	;
	goto L25
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v84 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v153 = F_palloc(m, v70<<(uint(int32(3))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L46
	}
L27:
	;
	if v105 == v55 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v95 = F_MemoryContextAlloc(m, v90, v70*int32(40)+int32(16))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v87 != v70 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v104 = v84
	v105 = v89
	goto L27
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v95
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v95))) = int64(0)
	v104 = v95
	v105 = v99
	goto L27
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v107 == v53 {
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v112 = v70 * int32(40)
	v114 = v112 + int32(16)
	if v104&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v114)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v55
	goto L26
L37:
	;
	if v114 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v114 == int32(0) {
		goto L36
	} else {
		goto L45
	}
L40:
	;
	v126 = v104 + v112 + int32(16)
	v128 = v104 + int32(4)
	if base.Ui32(v128) < base.Ui32(v126) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v130 = v126
	goto L43
L42:
	;
	v130 = v128
	goto L43
L43:
	;
	v135 = (v104^int32(-1)+v130)&int32(-4) + int32(4)
	if v135 == int32(0) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	base.MemoryFill(m, v104, int32(0), v135)
	goto L36
L45:
	;
	base.MemoryFill(m, v104, int32(0), v114)
	goto L36
L46:
	;
	v155 = F_palloc(m, v70)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	if v54 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	if v70 <= int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v160 = v70 & int32(3)
	v161 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = v161
	v180 = v2
	goto L53
L51:
	;
	v234 = v161
	goto L52
L52:
	;
	v255 = v234
	v257 = v161
	goto L57
L53:
	;
	v188 = int32(3)
	v191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v153+v168<<(uint(v188)%32)))) = v191
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v168+v155))) = uint8(v194)
	v197 = v168 | v194
	*(*int64)(unsafe.Add(mBase, uint32(v153+v197<<(uint(v188)%32)))) = v191
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v197))) = uint8(v194)
	v207 = v168 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v153+v207<<(uint(v188)%32)))) = v191
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v207))) = uint8(v194)
	v217 = v168 | v188
	*(*int64)(unsafe.Add(mBase, uint32(v153+v217<<(uint(v188)%32)))) = v191
	*(*uint8)(unsafe.Add(mBase, uint32(v155+v217))) = uint8(v194)
	v226 = int32(4)
	v227 = v168 + v226
	v229 = v180 + v226
	if v229 != v70&int32(2147483644) {
		v168 = v227
		v180 = v229
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v160 == int32(0) {
		goto L3
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v234 = v227
	goto L52
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v153+v255<<(uint(int32(3))%32)))) = int64(0)
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v255+v155))) = uint8(v281)
	v286 = v257 + v281
	if v286 != v160 {
		v255 = v255 + v281
		v257 = v286
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L3
L59:
	;
	goto L58
L60:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_hstore_populate_record_0), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_hstore_populate_record_1), int32(1015), int32(_a_F_hstore_populate_record_2))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	if v70 <= int32(0) {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L3
L66:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v367 = v68 + v361<<(uint(int32(3))%32) + v340*int32(100)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+119)))
	if v368 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L2
L68:
	;
	v578 = v340 + int32(1)
	if v578 != v70 {
		v340 = v578
		goto L66
	} else {
		goto L127
	}
L69:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+96))
	v376 = v367 + int32(32)
	v377 = F_strlen(m, v376)
	mBase = m.M
	goto L76
L70:
	;
	v563 = int32(1)
	goto L71
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v340+v155))) = uint8(v563)
	goto L68
L72:
	;
	v489 = int32(0)
	if v54 != 0 {
		goto L104
	} else {
		goto L105
	}
L73:
	;
	goto L72
L76:
	;
	v386 = int32(0)
	goto L77
L77:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v390 = v388 & int32(268435455)
	if v386 < v390 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v393 = v58 + int32(8)
	v402 = v386
	v403 = v390
	goto L81
L79:
	;
	goto L80
L80:
	;
	v478 = int32(-1)
	goto L73
L81:
	;
	v411 = int32(base.Ui32(v403-v402)>>(uint(int32(1))%32)) + v402
	v414 = v393 + v411<<(uint(int32(3))%32)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v417 = v415 & int32(1073741823)
	if int32(0) <= v415 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	goto L80
L83:
	;
	v448 = base.B2i32(v443 < int32(0))
	if v443 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L84:
	;
	v438 = F_memcmp(m, v436+(v393+v390<<(uint(int32(3))%32)), v376, v377)
	mBase = m.M
	if v438 != 0 {
		v443 = v438
		goto L83
	} else {
		goto L94
	}
L85:
	;
	if base.Ui32(v377) < base.Ui32(v429) {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v414-int32(4))))
	v424 = v422 & int32(1073741823)
	v425 = v417 - v424
	if v425 != v377 {
		v429 = v425
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v377 == v417 {
		v436 = int32(0)
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v436 = v424
	goto L84
L90:
	;
	v429 = v417
	goto L85
L91:
	;
	v434 = int32(1)
	goto L93
L92:
	;
	v434 = int32(-1)
	goto L93
L93:
	;
	v443 = v434
	goto L83
L94:
	;
	v478 = v411
	goto L73
L96:
	;
	v449 = v411 + int32(1)
	goto L98
L97:
	;
	v449 = v402
	goto L98
L98:
	;
	if v443 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v450 = v403
	goto L101
L100:
	;
	v450 = v411
	goto L101
L101:
	;
	if v449 < v450 {
		v402 = v449
		v403 = v450
		goto L81
	} else {
		goto L102
	}
L102:
	;
	goto L82
L104:
	;
	v492 = base.B2i32(v478 < v489)
	goto L106
L105:
	;
	v492 = v489
	goto L106
L106:
	;
	if v492 != 0 {
		goto L68
	} else {
		goto L107
	}
L107:
	;
	v495 = v104 + int32(16) + v340*int32(40)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v371 != v496 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_getTypeInputInfo(m, v371, v495+int32(4), v495+int32(8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v512 = int32(0)
	v513 = int32(1)
	if v478 < v512 {
		v549 = v513
		v552 = v512
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	F_fmgr_info_cxt(m, v504, v495+int32(12), v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v371
	goto L110
L113:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v367+int32(28))+76))
	v560 = F_InputFunctionCall(m, v495+int32(12), v552, v558, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L126
	}
L114:
	;
	v518 = v332 + v478<<(uint(int32(3))%32)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v519&int32(1073741824) != 0 {
		v549 = v513
		v552 = v512
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if v519 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v530 = v519 & int32(1073741823)
	goto L118
L117:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v530 = v519 - v526&int32(1073741823)
	goto L118
L118:
	;
	v533 = F_palloc(m, v530+int32(1))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v535 = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v535 <= v536 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v542 = v539 & int32(1073741823)
	goto L122
L121:
	;
	v542 = v535
	goto L122
L122:
	;
	if v530 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	base.MemoryCopy(m, v533, v542+(v332+v62<<(uint(int32(3))%32)), v530)
	goto L125
L124:
	;
	goto L125
L125:
	;
	v545 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v530+v533))) = uint8(v545)
	v549 = v545
	v552 = v533
	goto L113
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v153+v340<<(uint(int32(3))%32)))) = v560
	v563 = v549
	goto L71
L127:
	;
	goto L67
L128:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v603 != v28 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	v606 = F_HeapTupleHeaderGetDatum(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L5
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if int32(0) <= v615 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+20))
	F_domain_check(m, v606, int32(0), v28, v104+int32(8), v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	F_DecrTupleDescRefCount(m, v68)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	v621 = F_HeapTupleHeaderGetDatum(m, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L5
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v644 = v621
	goto L1
}
func F_hstore_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_pq_begintypsend(m, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v28 = v18 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v25))) = base.I32_rotr(v28, int32(24))&int32(16711695) | base.I32_rotr(v18&int32(16711935), int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v24 + int32(4)
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v43 = v14 + int32(8)
	v46 = v43 + v28<<(uint(int32(3))%32)
	v51 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v172 << (uint(int32(2)) % 32)
	goto L32
L8:
	;
	v58 = v43 + v51<<(uint(int32(3))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v61 = v59 & int32(1073741823)
	if int32(0) <= v59 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58-int32(4))))
	v70 = v61 - v66&int32(1073741823)
	goto L12
L11:
	;
	v70 = v61
	goto L12
L12:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v74+v75))) = base.I32_rotr(v70, int32(24))&v79 | base.I32_rotr(v70&v79, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v74 + int32(4)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if int32(0) <= v90 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v58-int32(4))))
	v99 = v95 & int32(1073741823)
	goto L16
L15:
	;
	v99 = int32(0)
	goto L16
L16:
	;
	F_pq_sendtext(m, v11, v99+v46, v70)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v103&int32(1073741824) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v160 = v51 + int32(1)
	if v160 != v28 {
		v51 = v160
		goto L8
	} else {
		goto L31
	}
L19:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v103 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v109+v110))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v109 + int32(4)
	goto L18
L23:
	;
	v125 = v103 & int32(1073741823)
	goto L25
L24:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v125 = v103 - v121&int32(1073741823)
	goto L25
L25:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v134 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v129+v130))) = base.I32_rotr(v125, int32(24))&v134 | base.I32_rotr(v125&v134, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v129 + int32(4)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if int32(0) <= v145 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v152 = v148 & int32(1073741823)
	goto L29
L28:
	;
	v152 = int32(0)
	goto L29
L29:
	;
	F_pq_sendtext(m, v11, v152+v46, v125)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	goto L9
L32:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v171)
}
func F_hstore_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.B2i32(l1 == int32(0))|l3 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_hstore_subscript_transform_0), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = F_exprLocation(m, l1)
					mBase = m.M
					F_parser_errposition(m, l2, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(58), int32(_a_F_hstore_subscript_transform_2))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_hstore_subscript_transform_0), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v59 = F_exprLocation(m, l1)
						mBase = m.M
						F_parser_errposition(m, l2, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(58), int32(_a_F_hstore_subscript_transform_2))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
			v20 = F_transformExpr(m, l2, v18, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = F_exprType(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = int32(-1)
					v29 = F_coerce_to_target_type(m, l2, v20, v22, int32(25), v25, int32(1), int32(2), v25)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_hstore_subscript_transform_3), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
										v79 = F_exprLocation(m, v78)
										mBase = m.M
										F_parser_errposition(m, l2, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(76), int32(_a_F_hstore_subscript_transform_2))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v29
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v29
							v38 = F_list_make1_impl(m, int32(1), v8+int32(8))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v38
								*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(-4294967271)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_hstore_version_diag(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v30 = v28 & int32(268435455)
	if base.B2i32(v30 == v17)|base.B2i32(v28 < v17) != 0 {
		v161 = int32(2)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v181 < int32(0) {
		v282 = v2
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v180 = v161
	goto L3
L5:
	;
	v37 = v13 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if int32(0) <= v38 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v180 = int32(0)
	goto L3
L7:
	;
	goto L8
L8:
	;
	v42 = int32(0)
	v44 = v30 << (uint(int32(3)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v37-int32(4))))
	v53 = v44 + v48&int32(1073741823) + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v56 = int32(base.Ui32(v54) >> (uint(int32(2)) % 32))
	if base.Ui32(v56) < base.Ui32(v53) {
		v161 = v42
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v58 = int32(1)
	v62 = v58
	goto L10
L10:
	;
	v74 = v37 + v62<<(uint(int32(2))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v75 < int32(0) {
		v161 = v42
		goto L4
	} else {
		goto L12
	}
L11:
	;
	if v30 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v78 = int32(1073741823)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74-int32(4))))
	if base.Ui32(v75&v78) < base.Ui32(v82&v78) {
		v161 = v42
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v87 = v62 + int32(1)
	if v87 != v30<<(uint(v58)%32) {
		v62 = v87
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v91 = int32(2)
	if base.Ui32(v30) <= base.Ui32(v91) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v53 == v56 {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	v94 = v91
	goto L20
L19:
	;
	v94 = v30
	goto L20
L20:
	;
	v98 = int32(1)
	goto L21
L21:
	;
	v108 = v98 << (uint(int32(3)) % 32)
	v109 = v37 + v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = v110 & int32(1073741823)
	if int32(0) <= v110 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L17
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109-int32(4))))
	v121 = v112 - v117&int32(1073741823)
	goto L25
L24:
	;
	v121 = v112
	goto L25
L25:
	;
	v122 = v13 + v108
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = v123 & int32(1073741823)
	if int32(0) <= v123 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122-int32(4))))
	v134 = v125 - v130&int32(1073741823)
	goto L28
L27:
	;
	v134 = v125
	goto L28
L28:
	;
	if v110&int32(1073741824)|base.B2i32(base.Ui32(v121) < base.Ui32(v134)) != 0 {
		v161 = int32(0)
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v141 = v98 + int32(1)
	if v141 != v94 {
		v98 = v141
		goto L21
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	v157 = int32(2)
	goto L33
L32:
	;
	v157 = int32(1)
	goto L33
L33:
	;
	v161 = v157
	goto L4
L34:
	;
	return base.I64_extend_i32_u(v282 + v180)
L35:
	;
	if v181 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v282 = int32(20)
	goto L34
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v181) {
		v282 = v2
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v192 = v181<<(uint(int32(3))%32) + int32(8)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v195 = int32(base.Ui32(v193) >> (uint(int32(2)) % 32))
	if base.Ui32(v195) < base.Ui32(v192) {
		v282 = v2
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if base.Ui32(v197) < base.Ui32(v198) {
		v282 = v2
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v202 = v13 + int32(8)
	if v181 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v207 = v197
	goto L45
L43:
	;
	goto L44
L44:
	;
	v237 = int32(1)
	if v181 <= v237 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v217 = v207 << (uint(int32(3)) % 32)
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202+v217))))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v217))))
	if base.Ui32(v219) < base.Ui32(v221) {
		v282 = v2
		goto L34
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v224 = v207 + int32(1)
	if v224 != v181 {
		v207 = v224
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v240 = v237
	goto L51
L50:
	;
	v240 = v181
	goto L51
L51:
	;
	v241 = int32(0)
	v243 = v241
	v244 = v241
	goto L52
L52:
	;
	v256 = v202 + v244<<(uint(int32(3))%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if int32(base.Ui32(v257)>>(uint(int32(1))%32)) != v243 {
		v282 = v2
		goto L34
	} else {
		goto L54
	}
L53:
	;
	v272 = v268 + v192
	if base.Ui32(v195) < base.Ui32(v272) {
		v282 = v2
		goto L34
	} else {
		goto L59
	}
L54:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256))))
	if v257&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v266 = int32(0)
	goto L57
L56:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+2)))
	v266 = v265
	goto L57
L57:
	;
	v268 = v266 + (v243 + v261)
	v270 = v244 + int32(1)
	if v270 != v240 {
		v243 = v268
		v244 = v270
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	if v272 == v195 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v277 = int32(20)
	goto L62
L61:
	;
	v277 = int32(10)
	goto L62
L62:
	;
	v282 = v277
	goto L34
}
