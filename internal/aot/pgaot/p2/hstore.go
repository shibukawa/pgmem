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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v14 = l1 << (uint(int32(3)) % 32)
	v16 = v14 + int32(8)
	v17 = v16 + l2
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
				v38 = v29
				v40 = v32
				v43 = int32(0)
				for {
					v49 = l0 + v43*int32(20)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
					if v51 != 0 {
						v52 = F__emscripten_memcpy_bulkmem(m, v40, v50, v51)
						mBase = m.M
						v53 = v52
					} else {
						v53 = v40
					}
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
					v55 = v53 + v54
					v58 = (v55 - v32) & int32(1073741823)
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v58
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+16)))
					if v60 == int32(1) {
						v74 = v55
						v75 = v58 | int32(1073741824)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
						if v66 != 0 {
							v67 = F__emscripten_memcpy_bulkmem(m, v55, v65, v66)
							mBase = m.M
							v68 = v67
						} else {
							v68 = v55
						}
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
						v70 = v68 + v69
						v74 = v70
						v75 = (v70 - v32) & int32(1073741823)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v75
					v80 = v43 + int32(1)
					if v80 != l1 {
						v38 = v38 + int32(8)
						v40 = v74
						v43 = v80
						continue
					} else {
						break
					}
					break
				}
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v85 = v82
				v88 = v74
			} else {
				v85 = v23
				v88 = v32
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v95 | int32(-2147483648)
			if v85&int32(268435455) != l1 {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v23
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v16 - v32 + v88) << (uint(int32(2)) % 32)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5467), int32(0), v4, v5)
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
			v43 = v34 + v28
			v44 = v38 + v43
			v45 = v42 + v44
			v49 = v43 - v42 ^ base.I32_rotl(v42, int32(16))
			v53 = v44 - v49 ^ base.I32_rotl(v49, int32(19))
			v58 = v49 + v45
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
		v333 = F_Int64GetDatum(m, base.I64_extend_i32_u(v323)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v315^v323-base.I32_rotl(v323, int32(24))))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v474 int32
	_ = v474
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
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
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
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
	var v606 int32
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
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
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
	return v628
L2:
	;
	v601 = F_heap_form_tuple(m, v64, v144, v146)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L5
	} else {
		goto L132
	}
L3:
	;
	if v66 <= int32(0) {
		goto L2
	} else {
		goto L68
	}
L4:
	;
	F_heap_deform_tuple(m, v24+int32(12), v64, v144, v146)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L67
	}
L5:
	;
	return int32(0)
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
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
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
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L63
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v55 = F_hstoreUpgrade(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = F_pg_detoast_datum(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	v51 = v2
	v52 = int32(-1)
	v53 = v28
	goto L11
L16:
	;
	goto L17
L17:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
	v628 = v2
	goto L1
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v46 == int32(1) {
		v628 = v44
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v51 = v44
	v52 = v49
	v53 = v50
	goto L11
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v59 = v57 & int32(268435455)
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = v59
	goto L23
L22:
	;
	v61 = int32(1)
	goto L23
L23:
	;
	if v61 == int32(0) {
		v628 = v51
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v64 = F_lookup_rowtype_tupdesc_domain(m, v53, v52)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v51 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v51
	v69 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v69
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)) = uint16(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(base.Ui32(v67) >> (uint(int32(2)) % 32))
	goto L28
L27:
	;
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v80 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v144 = F_palloc(m, v66<<(uint(int32(2))%32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L49
	}
L30:
	;
	if v101 == v53 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v91 = F_MemoryContextAlloc(m, v86, v66*int32(40)+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v83 != v66 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v100 = v80
	v101 = v85
	goto L30
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v91
	v95 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = int64(0)
	v100 = v91
	v101 = v95
	goto L30
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v103 == v52 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v106 = v66 * int32(40)
	v108 = v106 + int32(16)
	if v100&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v53
	goto L29
L40:
	;
	v134 = F__emscripten_memset_bulkmem(m, v100, base.I32_extend8_s(int32(0)), v108)
	mBase = m.M
	goto L48
L41:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v108) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if v108 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v120 = v100 + v106 + int32(16)
	v122 = v100 + int32(4)
	if base.Ui32(v122) < base.Ui32(v120) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v124 = v120
	goto L46
L45:
	;
	v124 = v122
	goto L46
L46:
	;
	v131 = F__emscripten_memset_bulkmem(m, v100, base.I32_extend8_s(int32(0)), (v100^int32(-1)+v124)&int32(-4)+int32(4))
	mBase = m.M
	goto L47
L47:
	;
	goto L39
L48:
	;
	goto L39
L49:
	;
	v146 = F_palloc(m, v66)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
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
	if v66 <= int32(0) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v151 = v66 & int32(3)
	v152 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v66) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = v152
	v161 = int32(0)
	goto L56
L54:
	;
	v224 = v152
	goto L55
L55:
	;
	if v151 == int32(0) {
		goto L3
	} else {
		goto L59
	}
L56:
	;
	v180 = int32(2)
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v144+v160<<(uint(v180)%32)))) = v183
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v146))) = uint8(v186)
	v189 = v160 | v186
	*(*int32)(unsafe.Add(mBase, uint32(v144+v189<<(uint(v180)%32)))) = v183
	*(*uint8)(unsafe.Add(mBase, uint32(v189+v146))) = uint8(v186)
	v199 = v160 | v180
	*(*int32)(unsafe.Add(mBase, uint32(v144+v199<<(uint(v180)%32)))) = v183
	*(*uint8)(unsafe.Add(mBase, uint32(v199+v146))) = uint8(v186)
	v209 = v160 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v144+v209<<(uint(v180)%32)))) = v183
	*(*uint8)(unsafe.Add(mBase, uint32(v209+v146))) = uint8(v186)
	v218 = int32(4)
	v219 = v160 + v218
	v221 = v161 + v218
	if v221 != v66&int32(2147483644) {
		v160 = v219
		v161 = v221
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v224 = v219
	goto L55
L58:
	;
	goto L57
L59:
	;
	v247 = v224
	v254 = v152
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144+v247<<(uint(int32(2))%32)))) = int32(0)
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v146))) = uint8(v273)
	v278 = v254 + v273
	if v278 != v151 {
		v247 = v247 + v273
		v254 = v278
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
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(365954), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(496551), int32(1015), int32(420960))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	goto L3
L68:
	;
	v328 = v55 + int32(8)
	v338 = int32(0)
	goto L69
L69:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v365 = v64 + int32(20) + v359<<(uint(int32(4))%32) + v338*int32(100)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+91)))
	if v366 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L2
L71:
	;
	v578 = v338 + int32(1)
	if v578 != v66 {
		v338 = v578
		goto L69
	} else {
		goto L131
	}
L72:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365)+68))
	v372 = v365 + int32(4)
	v373 = F_strlen(m, v372)
	mBase = m.M
	goto L79
L73:
	;
	v565 = int32(1)
	goto L74
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v338+v146))) = uint8(v565)
	goto L71
L75:
	;
	v485 = int32(0)
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
	v382 = int32(0)
	goto L80
L80:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v386 = v384 & int32(268435455)
	if v382 < v386 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v389 = v55 + int32(8)
	v398 = v382
	v399 = v386
	goto L84
L82:
	;
	goto L83
L83:
	;
	v474 = int32(-1)
	goto L76
L84:
	;
	v406 = base.I32_div_s(v399-v398, int32(2))
	v407 = v406 + v398
	v410 = v389 + v407<<(uint(int32(3))%32)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v413 = v411 & int32(1073741823)
	if int32(0) <= v411 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	goto L83
L86:
	;
	v444 = base.B2i32(v439 < int32(0))
	if v439 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v434 = F_memcmp(m, v432+(v389+v386<<(uint(int32(3))%32)), v372, v373)
	mBase = m.M
	if v434 != 0 {
		v439 = v434
		goto L86
	} else {
		goto L97
	}
L88:
	;
	if base.Ui32(v373) < base.Ui32(v425) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v410-int32(4))))
	v420 = v418 & int32(1073741823)
	v421 = v413 - v420
	if v421 != v373 {
		v425 = v421
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v373 == v413 {
		v432 = int32(0)
		goto L87
	} else {
		goto L93
	}
L92:
	;
	v432 = v420
	goto L87
L93:
	;
	v425 = v413
	goto L88
L94:
	;
	v430 = int32(1)
	goto L96
L95:
	;
	v430 = int32(-1)
	goto L96
L96:
	;
	v439 = v430
	goto L86
L97:
	;
	v474 = v407
	goto L76
L99:
	;
	v445 = v407 + int32(1)
	goto L101
L100:
	;
	v445 = v398
	goto L101
L101:
	;
	if v439 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v446 = v399
	goto L104
L103:
	;
	v446 = v407
	goto L104
L104:
	;
	if v445 < v446 {
		v398 = v445
		v399 = v446
		goto L84
	} else {
		goto L105
	}
L105:
	;
	goto L85
L107:
	;
	v488 = base.B2i32(v474 < v485)
	goto L109
L108:
	;
	v488 = v485
	goto L109
L109:
	;
	if v488 != 0 {
		goto L71
	} else {
		goto L110
	}
L110:
	;
	v491 = v100 + int32(16) + v338*int32(40)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	if v369 != v492 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_getTypeInputInfo(m, v369, v491+int32(4), v491+int32(8))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v508 = int32(0)
	v509 = int32(1)
	if v474 < v508 {
		v551 = v509
		v552 = v508
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+20))
	F_fmgr_info_cxt(m, v500, v491+int32(12), v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v369
	goto L113
L116:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v365)+76))
	v562 = F_InputFunctionCall(m, v491+int32(12), v552, v560, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L130
	}
L117:
	;
	v516 = v328 + v474<<(uint(int32(3))%32) + int32(4)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	if v517&int32(1073741824) != 0 {
		v551 = v509
		v552 = v508
		goto L116
	} else {
		goto L118
	}
L118:
	;
	if v517 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v530 = v517 & int32(1073741823)
	goto L121
L120:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v516-int32(4))))
	v530 = v517 - v526&int32(1073741823)
	goto L121
L121:
	;
	v533 = F_palloc(m, v530+int32(1))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v535 = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	if v535 <= v536 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v516-int32(4))))
	v544 = v541 & int32(1073741823)
	goto L125
L124:
	;
	v544 = v535
	goto L125
L125:
	;
	if v530 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v549 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v547+v530))) = uint8(v549)
	v551 = v535
	v552 = v533
	goto L116
L127:
	;
	v546 = F__emscripten_memcpy_bulkmem(m, v533, v544+(v328+v59<<(uint(int32(3))%32)), v530)
	mBase = m.M
	v547 = v546
	goto L129
L128:
	;
	v547 = v533
	goto L129
L129:
	;
	goto L126
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144+v338<<(uint(int32(2))%32)))) = v562
	v565 = v551
	goto L74
L131:
	;
	goto L70
L132:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v603 != v28 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	v606 = F_HeapTupleHeaderGetDatum(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L5
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if int32(0) <= v615 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+20))
	F_domain_check(m, v606, int32(0), v28, v100+int32(8), v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	F_DecrTupleDescRefCount(m, v64)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	v621 = F_HeapTupleHeaderGetDatum(m, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L5
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v628 = v621
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	v27 = int32(24)
	v29 = int32(65280)
	v31 = int32(8)
	v39 = v18 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v25))) = v18<<(uint(v27)%32) | v18&v29<<(uint(v31)%32) | (int32(base.Ui32(v18)>>(uint(v31)%32))&v29 | int32(base.Ui32(v39)>>(uint(v27)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v24 + int32(4)
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v49 = v14 + int32(8)
	v52 = v49 + v39<<(uint(int32(3))%32)
	v58 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v198 << (uint(int32(2)) % 32)
	goto L32
L8:
	;
	v64 = v49 + v58<<(uint(int32(3))%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v67 = v65 & int32(1073741823)
	if int32(0) <= v65 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64-int32(4))))
	v76 = v67 - v72&int32(1073741823)
	goto L12
L11:
	;
	v76 = v67
	goto L12
L12:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v83 = int32(24)
	v85 = int32(65280)
	v87 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v80+v81))) = v76<<(uint(v83)%32) | v76&v85<<(uint(v87)%32) | (int32(base.Ui32(v76)>>(uint(v87)%32))&v85 | int32(base.Ui32(v76)>>(uint(v83)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v80 + int32(4)
	v102 = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v102 <= v103 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v64-int32(4))))
	v111 = v108 & int32(1073741823)
	goto L16
L15:
	;
	v111 = v102
	goto L16
L16:
	;
	F_pq_sendtext(m, v11, v111+v52, v76)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v119 = v49 + v58<<(uint(int32(3))%32) + int32(4)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v120&int32(1073741824) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v186 = v58 + int32(1)
	if v186 != v39 {
		v58 = v186
		goto L8
	} else {
		goto L31
	}
L19:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v120 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v127))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v126 + int32(4)
	goto L18
L23:
	;
	v144 = v120 & int32(1073741823)
	goto L25
L24:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v119-int32(4))))
	v144 = v120 - v140&int32(1073741823)
	goto L25
L25:
	;
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v151 = int32(24)
	v153 = int32(65280)
	v155 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v148+v149))) = v144<<(uint(v151)%32) | v144&v153<<(uint(v155)%32) | (int32(base.Ui32(v144)>>(uint(v155)%32))&v153 | int32(base.Ui32(v144)>>(uint(v151)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
	v170 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v170 <= v171 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v119-int32(4))))
	v179 = v176 & int32(1073741823)
	goto L29
L28:
	;
	v179 = v170
	goto L29
L29:
	;
	F_pq_sendtext(m, v11, v179+v52, v144)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
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
	return v197
}
func F_hstore_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errmsg(m, int32(84033), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = F_exprLocation(m, l1)
					mBase = m.M
					F_parser_errposition(m, l2, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(495431), int32(57), int32(288197))
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
		if l3 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_errmsg(m, int32(84033), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = F_exprLocation(m, l1)
						mBase = m.M
						F_parser_errposition(m, l2, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							F_errfinish(m, int32(495431), int32(57), int32(288197))
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
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v12 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errmsg(m, int32(84033), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = F_exprLocation(m, l1)
							mBase = m.M
							F_parser_errposition(m, l2, v60)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_errfinish(m, int32(495431), int32(57), int32(288197))
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
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
				v19 = F_transformExpr(m, l2, v17, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = F_exprType(m, v19)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = int32(-1)
						v28 = F_coerce_to_target_type(m, l2, v19, v21, int32(25), v24, int32(1), int32(2), v24)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							if v28 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_errmsg(m, int32(63416), int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
											v84 = F_exprLocation(m, v83)
											mBase = m.M
											F_parser_errposition(m, l2, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												F_errfinish(m, int32(495431), int32(75), int32(288197))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v28
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
								v37 = F_list_make1_impl(m, int32(1), v8+int32(8))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v37
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
}
func F_hstore_version_diag(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v99 int32
	_ = v99
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
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
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
	v28 = int32(2)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v31 = v29 & int32(268435455)
	if v31 == int32(0) {
		v159 = v28
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v180 < int32(0) {
		v287 = v2
		goto L36
	} else {
		goto L37
	}
L4:
	;
	v179 = v159
	goto L3
L5:
	;
	if v29 < int32(0) {
		v159 = v28
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = v14 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if int32(0) <= v38 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v179 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v42 = int32(0)
	v44 = v31 << (uint(int32(3)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v37-int32(4))))
	v53 = v44 + v48&int32(1073741823) + int32(8)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v56 = int32(base.Ui32(v54) >> (uint(int32(2)) % 32))
	if base.Ui32(v56) < base.Ui32(v53) {
		v159 = v42
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v58 = int32(1)
	v62 = v58
	goto L11
L11:
	;
	v74 = v37 + v62<<(uint(int32(2))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v75 < int32(0) {
		v159 = v42
		goto L4
	} else {
		goto L13
	}
L12:
	;
	if v31 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v78 = int32(1073741823)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74-int32(4))))
	if base.Ui32(v75&v78) < base.Ui32(v82&v78) {
		v159 = v42
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v87 = v62 + int32(1)
	if v87 != v31<<(uint(v58)%32) {
		v62 = v87
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v91 = int32(2)
	if base.Ui32(v31) <= base.Ui32(v91) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v53 == v56 {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v94 = v91
	goto L21
L20:
	;
	v94 = v31
	goto L21
L21:
	;
	v99 = int32(1)
	goto L22
L22:
	;
	v108 = v99 << (uint(int32(3)) % 32)
	v109 = v37 + v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = v110 & int32(1073741823)
	if int32(0) <= v110 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L18
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109-int32(4))))
	v121 = v112 - v117&int32(1073741823)
	goto L26
L25:
	;
	v121 = v112
	goto L26
L26:
	;
	v122 = v14 + v108
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = v123 & int32(1073741823)
	if int32(0) <= v123 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122-int32(4))))
	v134 = v125 - v130&int32(1073741823)
	goto L29
L28:
	;
	v134 = v125
	goto L29
L29:
	;
	v135 = int32(0)
	if v110&int32(1073741824) != 0 {
		v159 = v135
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v121) < base.Ui32(v134) {
		v159 = v135
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v140 = v99 + int32(1)
	if v140 != v94 {
		v99 = v140
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	v156 = int32(2)
	goto L35
L34:
	;
	v156 = int32(1)
	goto L35
L35:
	;
	v159 = v156
	goto L4
L36:
	;
	return v179 + v287
L37:
	;
	if v180 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return v179 + int32(20)
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v180) {
		v287 = v2
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v193 = v180<<(uint(int32(3))%32) + int32(8)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v196 = int32(base.Ui32(v194) >> (uint(int32(2)) % 32))
	if base.Ui32(v196) < base.Ui32(v193) {
		v287 = v2
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v198 = int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v198) < base.Ui32(v199) {
		v287 = v2
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v203 = v14 + int32(8)
	if v180 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v208 = v198
	goto L47
L45:
	;
	goto L46
L46:
	;
	v240 = int32(1)
	if v180 <= v240 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v219 = v208 << (uint(int32(3)) % 32)
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203+v219))))
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v219))))
	if base.Ui32(v221) < base.Ui32(v223) {
		v287 = v2
		goto L36
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v226 = v208 + int32(1)
	if v226 != v180 {
		v208 = v226
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v243 = v240
	goto L53
L52:
	;
	v243 = v180
	goto L53
L53:
	;
	v244 = int32(0)
	v246 = v244
	v247 = v244
	goto L54
L54:
	;
	v260 = v203 + v247<<(uint(int32(3))%32)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if int32(base.Ui32(v261)>>(uint(int32(1))%32)) != v246 {
		v287 = v2
		goto L36
	} else {
		goto L56
	}
L55:
	;
	v276 = v272 + v193
	if base.Ui32(v196) < base.Ui32(v276) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260))))
	if v261&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v270 = int32(0)
	goto L59
L58:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+2)))
	v270 = v269
	goto L59
L59:
	;
	v272 = v270 + (v246 + v265)
	v274 = v247 + int32(1)
	if v274 != v243 {
		v246 = v272
		v247 = v274
		goto L54
	} else {
		goto L60
	}
L60:
	;
	goto L55
L61:
	;
	return v179
L62:
	;
	goto L63
L63:
	;
	if v276 == v196 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v282 = int32(20)
	goto L66
L65:
	;
	v282 = int32(10)
	goto L66
L66:
	;
	v287 = v282
	goto L36
}
