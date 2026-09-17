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
func F_hstore_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_ge_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_hstore_hash_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
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
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_hstoreUpgrade(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(4)
		v10 = v5 + v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v15 = int32(base.Ui32(v11)>>(uint(int32(2))%32)) - v9
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
		v23 = v15 - int32(1636608432)
		if v17 == int64(0) {
			v60 = v23
			v62 = v23
			v64 = v23
		} else {
			v27 = v23 + base.I32_wrap_i64(v17)
			v28 = v27 + v23
			v32 = int32(4)
			v34 = base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) ^ base.I32_rotl(v23, v32)
			v38 = v27 - v34 ^ base.I32_rotl(v34, int32(6))
			v42 = v28 - v38 ^ base.I32_rotl(v38, int32(8))
			v43 = v28 + v34
			v44 = v38 + v43
			v45 = v42 + v44
			v49 = v43 - v42 ^ base.I32_rotl(v42, int32(16))
			v53 = v44 - v49 ^ base.I32_rotl(v49, int32(19))
			v58 = v45 + v49
			v60 = v58
			v62 = v45 - v53 ^ base.I32_rotl(v53, v32)
			v64 = v53 + v58
		}
		if v10&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v15) {
				v69 = v10
				v70 = v15
				v72 = v60
				v73 = v64
				v74 = v62
				for {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
					v77 = v76 + v73
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
					v81 = v80 + v74
					v83 = int32(4)
					v85 = v78 + v72 - v81 ^ base.I32_rotl(v81, v83)
					v89 = v77 - v85 ^ base.I32_rotl(v85, int32(6))
					v90 = v81 + v77
					v91 = v85 + v90
					v92 = v89 + v91
					v96 = v90 - v89 ^ base.I32_rotl(v89, int32(8))
					v100 = v91 - v96 ^ base.I32_rotl(v96, int32(16))
					v104 = v92 - v100 ^ base.I32_rotl(v100, int32(19))
					v105 = v96 + v92
					v106 = v100 + v105
					v107 = v104 + v106
					v111 = v105 - v104 ^ base.I32_rotl(v104, v83)
					v112 = int32(12)
					v113 = v69 + v112
					v115 = v70 - v112
					if base.Ui32(int32(11)) < base.Ui32(v115) {
						v69 = v113
						v70 = v115
						v72 = v106
						v73 = v107
						v74 = v111
						continue
					} else {
						break
					}
					break
				}
				v118 = v113
				v119 = v115
				v121 = v106
				v122 = v107
				v123 = v111
			} else {
				v118 = v10
				v119 = v15
				v121 = v60
				v122 = v64
				v123 = v62
			}
			switch v119 - int32(1) {
			case 0:
				v288 = v121
				v289 = v122
				v290 = v123
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 1:
				v281 = v121
				v282 = v122
				v283 = v123
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 2:
				v274 = v121
				v275 = v122
				v276 = v123
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 3:
				v268 = v122
				v269 = v123
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 4:
				v264 = v122
				v265 = v123
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 5:
				v258 = v122
				v259 = v123
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 6:
				v252 = v122
				v253 = v123
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
				v258 = v254<<(uint(int32(16))%32) + v252
				v259 = v253
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 7:
				v247 = v123
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+7)))
				v252 = v248<<(uint(int32(24))%32) + v122
				v253 = v247
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
				v258 = v254<<(uint(int32(16))%32) + v252
				v259 = v253
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 8:
				v242 = v123
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+8)))
				v247 = v243<<(uint(int32(8))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+7)))
				v252 = v248<<(uint(int32(24))%32) + v122
				v253 = v247
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
				v258 = v254<<(uint(int32(16))%32) + v252
				v259 = v253
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 9:
				v237 = v123
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+9)))
				v242 = v238<<(uint(int32(16))%32) + v237
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+8)))
				v247 = v243<<(uint(int32(8))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+7)))
				v252 = v248<<(uint(int32(24))%32) + v122
				v253 = v247
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
				v258 = v254<<(uint(int32(16))%32) + v252
				v259 = v253
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			case 10:
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+10)))
				v237 = v233<<(uint(int32(24))%32) + v123
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+9)))
				v242 = v238<<(uint(int32(16))%32) + v237
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+8)))
				v247 = v243<<(uint(int32(8))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+7)))
				v252 = v248<<(uint(int32(24))%32) + v122
				v253 = v247
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+6)))
				v258 = v254<<(uint(int32(16))%32) + v252
				v259 = v253
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
				v264 = v260<<(uint(int32(8))%32) + v258
				v265 = v259
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
				v268 = v264 + v266
				v269 = v265
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
				v274 = v270<<(uint(int32(24))%32) + v121
				v275 = v268
				v276 = v269
				v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
				v281 = v277<<(uint(int32(16))%32) + v274
				v282 = v275
				v283 = v276
				v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
				v288 = v284<<(uint(int32(8))%32) + v281
				v289 = v282
				v290 = v283
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
				v296 = v288 + v291
				v297 = v289
				v298 = v290
			default:
				v296 = v121
				v297 = v122
				v298 = v123
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v15) {
				v129 = v10
				v130 = v15
				v132 = v60
				v133 = v64
				v134 = v62
				for {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
					v137 = v136 + v133
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
					v141 = v140 + v134
					v143 = int32(4)
					v145 = v138 + v132 - v141 ^ base.I32_rotl(v141, v143)
					v149 = v137 - v145 ^ base.I32_rotl(v145, int32(6))
					v150 = v141 + v137
					v151 = v145 + v150
					v152 = v149 + v151
					v156 = v150 - v149 ^ base.I32_rotl(v149, int32(8))
					v160 = v151 - v156 ^ base.I32_rotl(v156, int32(16))
					v164 = v152 - v160 ^ base.I32_rotl(v160, int32(19))
					v165 = v156 + v152
					v166 = v160 + v165
					v167 = v164 + v166
					v171 = v165 - v164 ^ base.I32_rotl(v164, v143)
					v172 = int32(12)
					v173 = v129 + v172
					v175 = v130 - v172
					if base.Ui32(int32(11)) < base.Ui32(v175) {
						v129 = v173
						v130 = v175
						v132 = v166
						v133 = v167
						v134 = v171
						continue
					} else {
						break
					}
					break
				}
				v178 = v173
				v179 = v175
				v181 = v166
				v182 = v167
				v183 = v171
			} else {
				v178 = v10
				v179 = v15
				v181 = v60
				v182 = v64
				v183 = v62
			}
			switch v179 - int32(1) {
			case 0:
				v230 = v181
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v296 = v230 + v231
				v297 = v182
				v298 = v183
			case 1:
				v225 = v181
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v230 = v226<<(uint(int32(8))%32) + v225
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v296 = v230 + v231
				v297 = v182
				v298 = v183
			case 2:
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v225 = v221<<(uint(int32(16))%32) + v181
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v230 = v226<<(uint(int32(8))%32) + v225
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v296 = v230 + v231
				v297 = v182
				v298 = v183
			case 3:
				v218 = v182
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v296 = v219 + v181
				v297 = v218
				v298 = v183
			case 4:
				v215 = v182
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v218 = v215 + v216
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v296 = v219 + v181
				v297 = v218
				v298 = v183
			case 5:
				v210 = v182
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v215 = v211<<(uint(int32(8))%32) + v210
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v218 = v215 + v216
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v296 = v219 + v181
				v297 = v218
				v298 = v183
			case 6:
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v210 = v206<<(uint(int32(16))%32) + v182
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v215 = v211<<(uint(int32(8))%32) + v210
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v218 = v215 + v216
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v296 = v219 + v181
				v297 = v218
				v298 = v183
			case 7:
				v201 = v183
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
				v296 = v202 + v181
				v297 = v204 + v182
				v298 = v201
			case 8:
				v196 = v183
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v201 = v197<<(uint(int32(8))%32) + v196
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
				v296 = v202 + v181
				v297 = v204 + v182
				v298 = v201
			case 9:
				v191 = v183
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v196 = v192<<(uint(int32(16))%32) + v191
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v201 = v197<<(uint(int32(8))%32) + v196
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
				v296 = v202 + v181
				v297 = v204 + v182
				v298 = v201
			case 10:
				v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
				v191 = v187<<(uint(int32(24))%32) + v183
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v196 = v192<<(uint(int32(16))%32) + v191
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v201 = v197<<(uint(int32(8))%32) + v196
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
				v296 = v202 + v181
				v297 = v204 + v182
				v298 = v201
			default:
				v296 = v181
				v297 = v182
				v298 = v183
			}
		}
		v301 = int32(14)
		v303 = v297 ^ v298 - base.I32_rotl(v297, v301)
		v307 = v303 ^ v296 - base.I32_rotl(v303, int32(11))
		v311 = v307 ^ v297 - base.I32_rotl(v307, int32(25))
		v315 = v311 ^ v303 - base.I32_rotl(v311, int32(16))
		v319 = v315 ^ v307 - base.I32_rotl(v315, int32(4))
		v323 = v319 ^ v311 - base.I32_rotl(v319, v301)
		v333 = F_Int64GetDatum(m, base.I64_extend_i32_u(v323)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v323^v315-base.I32_rotl(v323, int32(24))))
		mBase = m.M
		v334 = m.ExcPending
		if v334 != 0 {
			return int32(0)
		} else {
			v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v335 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v338 = m.ExcPending
				if v338 != 0 {
					return int32(0)
				} else {
					return v333
				}
			} else {
				return v333
			}
		}
	}
}
func F_hstore_populate_record(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
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
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v569 int32
	_ = v569
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v622 int32
	_ = v622
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_get_fn_expr_argtype(m, v25, v2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v23 + int32(32)
	return v622
L2:
	;
	v591 = F_heap_form_tuple(m, v63, v148, v150)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L5
	} else {
		goto L131
	}
L3:
	;
	v324 = v54 + int32(8)
	v332 = int32(0)
	goto L69
L4:
	;
	F_heap_deform_tuple(m, v23+int32(12), v63, v148, v150)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L67
	}
L5:
	;
	return int32(0)
L6:
	;
	v31 = F_type_is_rowtype(m, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v33 == int32(1) {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L63
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = F_hstoreUpgrade(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v36 != int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = F_pg_detoast_datum(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	v50 = int32(-1)
	v51 = v2
	v52 = v27
	goto L11
L16:
	;
	goto L17
L17:
	;
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
	v622 = v2
	goto L1
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v45 == int32(1) {
		v622 = v43
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v50 = v48
	v51 = v43
	v52 = v49
	goto L11
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v58 = v56 & int32(268435455)
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v60 = v58
	goto L23
L22:
	;
	v60 = int32(1)
	goto L23
L23:
	;
	if v60 == int32(0) {
		v622 = v51
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v63 = F_lookup_rowtype_tupdesc_domain(m, v52, v50)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v51 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v51
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v68
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+20)) = uint16(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(base.Ui32(v66) >> (uint(int32(2)) % 32))
	goto L28
L27:
	;
	goto L28
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v79 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v148 = F_palloc(m, v65<<(uint(int32(2))%32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L49
	}
L30:
	;
	if v100 == v52 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v90 = F_MemoryContextAlloc(m, v85, v65*int32(40)+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v82 != v65 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v99 = v79
	v100 = v84
	goto L30
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v90
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = int64(0)
	v99 = v90
	v100 = v94
	goto L30
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 == v50 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v107 = v65 * int32(40)
	v109 = v107 + int32(16)
	if v99&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v109)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v52
	goto L29
L40:
	;
	if v109 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v109 == int32(0) {
		goto L39
	} else {
		goto L48
	}
L43:
	;
	v121 = v99 + v107 + int32(16)
	v123 = v99 + int32(4)
	if base.Ui32(v123) < base.Ui32(v121) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v125 = v121
	goto L46
L45:
	;
	v125 = v123
	goto L46
L46:
	;
	v130 = (v99^int32(-1)+v125)&int32(-4) + int32(4)
	if v130 == int32(0) {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	base.MemoryFill(m, v99, int32(0), v130)
	goto L39
L48:
	;
	base.MemoryFill(m, v99, int32(0), v109)
	goto L39
L49:
	;
	v150 = F_palloc(m, v65)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if v51 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	if v65 <= int32(0) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v155 = v65 & int32(3)
	v156 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v163 = v156
	v164 = int32(0)
	goto L56
L54:
	;
	v228 = v156
	goto L55
L55:
	;
	v249 = v228
	v250 = int32(0)
	goto L60
L56:
	;
	v182 = int32(2)
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v148+v163<<(uint(v182)%32)))) = v185
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v150))) = uint8(v188)
	v191 = v163 | v188
	*(*int32)(unsafe.Add(mBase, uint32(v148+v191<<(uint(v182)%32)))) = v185
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v150))) = uint8(v188)
	v201 = v163 | v182
	*(*int32)(unsafe.Add(mBase, uint32(v148+v201<<(uint(v182)%32)))) = v185
	*(*uint8)(unsafe.Add(mBase, uint32(v201+v150))) = uint8(v188)
	v211 = v163 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v148+v211<<(uint(v182)%32)))) = v185
	*(*uint8)(unsafe.Add(mBase, uint32(v211+v150))) = uint8(v188)
	v220 = int32(4)
	v221 = v163 + v220
	v223 = v164 + v220
	if v223 != v65&int32(2147483644) {
		v163 = v221
		v164 = v223
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v155 == int32(0) {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v228 = v221
	goto L55
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v249<<(uint(int32(2))%32)))) = int32(0)
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v249+v150))) = uint8(v274)
	v279 = v250 + v274
	if v279 != v155 {
		v249 = v249 + v274
		v250 = v279
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L3
L62:
	;
	goto L61
L63:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_hstore_populate_record_0), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_hstore_populate_record_1), int32(1015), int32(_a_F_hstore_populate_record_2))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	if v65 <= int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L3
L69:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v358 = v63 + v352<<(uint(int32(4))%32) + v332*int32(100)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+111)))
	if v359 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L2
L71:
	;
	v569 = v332 + int32(1)
	if v569 != v65 {
		v332 = v569
		goto L69
	} else {
		goto L130
	}
L72:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+88))
	v367 = v358 + int32(24)
	v368 = F_strlen(m, v367)
	mBase = m.M
	goto L79
L73:
	;
	v554 = int32(1)
	goto L74
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v332+v150))) = uint8(v554)
	goto L71
L75:
	;
	v480 = int32(0)
	if v51 != 0 {
		goto L107
	} else {
		goto L108
	}
L76:
	;
	goto L75
L79:
	;
	v377 = int32(0)
	goto L80
L80:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v381 = v379 & int32(268435455)
	if v377 < v381 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v384 = v54 + int32(8)
	v393 = v377
	v394 = v381
	goto L84
L82:
	;
	goto L83
L83:
	;
	v469 = int32(-1)
	goto L76
L84:
	;
	v402 = int32(base.Ui32(v394-v393)>>(uint(int32(1))%32)) + v393
	v405 = v384 + v402<<(uint(int32(3))%32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v408 = v406 & int32(1073741823)
	if int32(0) <= v406 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	goto L83
L86:
	;
	v439 = base.B2i32(v434 < int32(0))
	if v434 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v429 = F_memcmp(m, v427+(v384+v381<<(uint(int32(3))%32)), v367, v368)
	mBase = m.M
	if v429 != 0 {
		v434 = v429
		goto L86
	} else {
		goto L97
	}
L88:
	;
	if base.Ui32(v368) < base.Ui32(v420) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v405-int32(4))))
	v415 = v413 & int32(1073741823)
	v416 = v408 - v415
	if v416 != v368 {
		v420 = v416
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v368 == v408 {
		v427 = int32(0)
		goto L87
	} else {
		goto L93
	}
L92:
	;
	v427 = v415
	goto L87
L93:
	;
	v420 = v408
	goto L88
L94:
	;
	v425 = int32(1)
	goto L96
L95:
	;
	v425 = int32(-1)
	goto L96
L96:
	;
	v434 = v425
	goto L86
L97:
	;
	v469 = v402
	goto L76
L99:
	;
	v440 = v402 + int32(1)
	goto L101
L100:
	;
	v440 = v393
	goto L101
L101:
	;
	if v434 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v441 = v394
	goto L104
L103:
	;
	v441 = v402
	goto L104
L104:
	;
	if v440 < v441 {
		v393 = v440
		v394 = v441
		goto L84
	} else {
		goto L105
	}
L105:
	;
	goto L85
L107:
	;
	v483 = base.B2i32(v469 < v480)
	goto L109
L108:
	;
	v483 = v480
	goto L109
L109:
	;
	if v483 != 0 {
		goto L71
	} else {
		goto L110
	}
L110:
	;
	v486 = v99 + int32(16) + v332*int32(40)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	if v362 != v487 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_getTypeInputInfo(m, v362, v486+int32(4), v486+int32(8))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L5
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v503 = int32(0)
	v504 = int32(1)
	if v469 < v503 {
		v540 = v504
		v543 = v503
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+20))
	F_fmgr_info_cxt(m, v495, v486+int32(12), v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v362
	goto L113
L116:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v486)+8))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v358+int32(20))+76))
	v551 = F_InputFunctionCall(m, v486+int32(12), v543, v549, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L5
	} else {
		goto L129
	}
L117:
	;
	v509 = v324 + v469<<(uint(int32(3))%32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v510&int32(1073741824) != 0 {
		v540 = v504
		v543 = v503
		goto L116
	} else {
		goto L118
	}
L118:
	;
	if v510 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v521 = v510 & int32(1073741823)
	goto L121
L120:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v521 = v510 - v517&int32(1073741823)
	goto L121
L121:
	;
	v524 = F_palloc(m, v521+int32(1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v526 = int32(0)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v526 <= v527 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v533 = v530 & int32(1073741823)
	goto L125
L124:
	;
	v533 = v526
	goto L125
L125:
	;
	if v521 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	base.MemoryCopy(m, v524, v533+(v324+v58<<(uint(int32(3))%32)), v521)
	goto L128
L127:
	;
	goto L128
L128:
	;
	v536 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v521+v524))) = uint8(v536)
	v540 = v536
	v543 = v524
	goto L116
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v332<<(uint(int32(2))%32)))) = v551
	v554 = v540
	goto L74
L130:
	;
	goto L70
L131:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v593 != v27 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v591)+16))
	v596 = F_HeapTupleHeaderGetDatum(m, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if int32(0) <= v605 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+20))
	F_domain_check(m, v596, int32(0), v27, v99+int32(8), v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	F_DecrTupleDescRefCount(m, v63)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L5
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v591)+16))
	v611 = F_HeapTupleHeaderGetDatum(m, v610)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L5
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v622 = v611
	goto L1
}
func F_hstore_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v171
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
						F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(57), int32(_a_F_hstore_subscript_transform_2))
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
							F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(57), int32(_a_F_hstore_subscript_transform_2))
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
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
											F_errfinish(m, int32(_a_F_hstore_subscript_transform_1), int32(75), int32(_a_F_hstore_subscript_transform_2))
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
func F_hstore_version_diag(m *base.Module, l0 int32) int32 {
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
		v291 = v2
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
	return v180 + v291
L35:
	;
	if v181 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return v180 + int32(20)
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v181) {
		v291 = v2
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v194 = v181<<(uint(int32(3))%32) + int32(8)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v197 = int32(base.Ui32(v195) >> (uint(int32(2)) % 32))
	if base.Ui32(v197) < base.Ui32(v194) {
		v291 = v2
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v199 = int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if base.Ui32(v199) < base.Ui32(v200) {
		v291 = v2
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v204 = v13 + int32(8)
	if v181 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v209 = v199
	goto L45
L43:
	;
	goto L44
L44:
	;
	v239 = int32(1)
	if v181 <= v239 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v219 = v209 << (uint(int32(3)) % 32)
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204+v219))))
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v219))))
	if base.Ui32(v221) < base.Ui32(v223) {
		v291 = v2
		goto L34
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v226 = v209 + int32(1)
	if v226 != v181 {
		v209 = v226
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v242 = v239
	goto L51
L50:
	;
	v242 = v181
	goto L51
L51:
	;
	v243 = int32(0)
	v245 = v243
	v246 = v243
	goto L52
L52:
	;
	v258 = v204 + v246<<(uint(int32(3))%32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if int32(base.Ui32(v259)>>(uint(int32(1))%32)) != v245 {
		v291 = v2
		goto L34
	} else {
		goto L54
	}
L53:
	;
	v274 = v270 + v194
	if base.Ui32(v197) < base.Ui32(v274) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258))))
	if v259&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v268 = int32(0)
	goto L57
L56:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+2)))
	v268 = v267
	goto L57
L57:
	;
	v270 = v268 + (v245 + v263)
	v272 = v246 + int32(1)
	if v272 != v242 {
		v245 = v270
		v246 = v272
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	return v180
L60:
	;
	goto L61
L61:
	;
	if v274 == v197 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v280 = int32(20)
	goto L64
L63:
	;
	v280 = int32(10)
	goto L64
L64:
	;
	v291 = v280
	goto L34
}
