package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_hash_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v16 != 0 {
				v46 = int32(base.Ui32(v14)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = F_str_tolower(m, v8+v17, v46, int32(100))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int64(0)
		} else {
			v50 = F_strlen(m, v48)
			mBase = m.M
			v56 = v50 - int32(1636608432)
			if v19 == int64(0) {
				v93 = v56
				v95 = v56
				v97 = v56
			} else {
				v60 = v56 + base.I32_wrap_i64(v19)
				v61 = v60 + v56
				v65 = int32(4)
				v67 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(32))%64))) ^ base.I32_rotl(v56, v65)
				v71 = v60 - v67 ^ base.I32_rotl(v67, int32(6))
				v75 = v61 - v71 ^ base.I32_rotl(v71, int32(8))
				v76 = v61 + v67
				v77 = v71 + v76
				v78 = v75 + v77
				v82 = v76 - v75 ^ base.I32_rotl(v75, int32(16))
				v86 = v77 - v82 ^ base.I32_rotl(v82, int32(19))
				v91 = v78 + v82
				v93 = v91
				v95 = v78 - v86 ^ base.I32_rotl(v86, v65)
				v97 = v86 + v91
			}
			if v48&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v50) {
					v102 = v48
					v103 = v50
					v105 = v93
					v106 = v97
					v107 = v95
					for {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
						v110 = v109 + v106
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
						v114 = v113 + v107
						v116 = int32(4)
						v118 = v111 + v105 - v114 ^ base.I32_rotl(v114, v116)
						v122 = v110 - v118 ^ base.I32_rotl(v118, int32(6))
						v123 = v114 + v110
						v124 = v118 + v123
						v125 = v122 + v124
						v129 = v123 - v122 ^ base.I32_rotl(v122, int32(8))
						v133 = v124 - v129 ^ base.I32_rotl(v129, int32(16))
						v137 = v125 - v133 ^ base.I32_rotl(v133, int32(19))
						v138 = v129 + v125
						v139 = v133 + v138
						v140 = v137 + v139
						v144 = v138 - v137 ^ base.I32_rotl(v137, v116)
						v145 = int32(12)
						v146 = v102 + v145
						v148 = v103 - v145
						if base.Ui32(int32(11)) < base.Ui32(v148) {
							v102 = v146
							v103 = v148
							v105 = v139
							v106 = v140
							v107 = v144
							continue
						} else {
							break
						}
						break
					}
					v151 = v146
					v152 = v148
					v154 = v139
					v155 = v140
					v156 = v144
				} else {
					v151 = v48
					v152 = v50
					v154 = v93
					v155 = v97
					v156 = v95
				}
				switch v152 - int32(1) {
				case 0:
					v321 = v154
					v322 = v155
					v323 = v156
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 1:
					v314 = v154
					v315 = v155
					v316 = v156
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 2:
					v307 = v154
					v308 = v155
					v309 = v156
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 3:
					v301 = v155
					v302 = v156
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 4:
					v297 = v155
					v298 = v156
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 5:
					v291 = v155
					v292 = v156
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 6:
					v285 = v155
					v286 = v156
					v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
					v291 = v287<<(uint(int32(16))%32) + v285
					v292 = v286
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 7:
					v280 = v156
					v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
					v285 = v281<<(uint(int32(24))%32) + v155
					v286 = v280
					v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
					v291 = v287<<(uint(int32(16))%32) + v285
					v292 = v286
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 8:
					v275 = v156
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
					v280 = v276<<(uint(int32(8))%32) + v275
					v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
					v285 = v281<<(uint(int32(24))%32) + v155
					v286 = v280
					v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
					v291 = v287<<(uint(int32(16))%32) + v285
					v292 = v286
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 9:
					v270 = v156
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+9)))
					v275 = v271<<(uint(int32(16))%32) + v270
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
					v280 = v276<<(uint(int32(8))%32) + v275
					v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
					v285 = v281<<(uint(int32(24))%32) + v155
					v286 = v280
					v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
					v291 = v287<<(uint(int32(16))%32) + v285
					v292 = v286
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				case 10:
					v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+10)))
					v270 = v266<<(uint(int32(24))%32) + v156
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+9)))
					v275 = v271<<(uint(int32(16))%32) + v270
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
					v280 = v276<<(uint(int32(8))%32) + v275
					v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
					v285 = v281<<(uint(int32(24))%32) + v155
					v286 = v280
					v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
					v291 = v287<<(uint(int32(16))%32) + v285
					v292 = v286
					v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
					v297 = v293<<(uint(int32(8))%32) + v291
					v298 = v292
					v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
					v301 = v297 + v299
					v302 = v298
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
					v307 = v303<<(uint(int32(24))%32) + v154
					v308 = v301
					v309 = v302
					v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
					v314 = v310<<(uint(int32(16))%32) + v307
					v315 = v308
					v316 = v309
					v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
					v321 = v317<<(uint(int32(8))%32) + v314
					v322 = v315
					v323 = v316
					v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
					v329 = v321 + v324
					v330 = v322
					v331 = v323
				default:
					v329 = v154
					v330 = v155
					v331 = v156
				}
			} else {
				if base.Ui32(int32(12)) <= base.Ui32(v50) {
					v162 = v48
					v163 = v50
					v165 = v93
					v166 = v97
					v167 = v95
					for {
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
						v170 = v169 + v166
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
						v174 = v173 + v167
						v176 = int32(4)
						v178 = v171 + v165 - v174 ^ base.I32_rotl(v174, v176)
						v182 = v170 - v178 ^ base.I32_rotl(v178, int32(6))
						v183 = v174 + v170
						v184 = v178 + v183
						v185 = v182 + v184
						v189 = v183 - v182 ^ base.I32_rotl(v182, int32(8))
						v193 = v184 - v189 ^ base.I32_rotl(v189, int32(16))
						v197 = v185 - v193 ^ base.I32_rotl(v193, int32(19))
						v198 = v189 + v185
						v199 = v193 + v198
						v200 = v197 + v199
						v204 = v198 - v197 ^ base.I32_rotl(v197, v176)
						v205 = int32(12)
						v206 = v162 + v205
						v208 = v163 - v205
						if base.Ui32(int32(11)) < base.Ui32(v208) {
							v162 = v206
							v163 = v208
							v165 = v199
							v166 = v200
							v167 = v204
							continue
						} else {
							break
						}
						break
					}
					v211 = v206
					v212 = v208
					v214 = v199
					v215 = v200
					v216 = v204
				} else {
					v211 = v48
					v212 = v50
					v214 = v93
					v215 = v97
					v216 = v95
				}
				switch v212 - int32(1) {
				case 0:
					v263 = v214
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
					v329 = v263 + v264
					v330 = v215
					v331 = v216
				case 1:
					v258 = v214
					v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
					v263 = v259<<(uint(int32(8))%32) + v258
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
					v329 = v263 + v264
					v330 = v215
					v331 = v216
				case 2:
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+2)))
					v258 = v254<<(uint(int32(16))%32) + v214
					v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
					v263 = v259<<(uint(int32(8))%32) + v258
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
					v329 = v263 + v264
					v330 = v215
					v331 = v216
				case 3:
					v251 = v215
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v329 = v252 + v214
					v330 = v251
					v331 = v216
				case 4:
					v248 = v215
					v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
					v251 = v248 + v249
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v329 = v252 + v214
					v330 = v251
					v331 = v216
				case 5:
					v243 = v215
					v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
					v248 = v244<<(uint(int32(8))%32) + v243
					v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
					v251 = v248 + v249
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v329 = v252 + v214
					v330 = v251
					v331 = v216
				case 6:
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)))
					v243 = v239<<(uint(int32(16))%32) + v215
					v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
					v248 = v244<<(uint(int32(8))%32) + v243
					v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
					v251 = v248 + v249
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v329 = v252 + v214
					v330 = v251
					v331 = v216
				case 7:
					v234 = v216
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v237 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
					v329 = v235 + v214
					v330 = v237 + v215
					v331 = v234
				case 8:
					v229 = v216
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v237 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
					v329 = v235 + v214
					v330 = v237 + v215
					v331 = v234
				case 9:
					v224 = v216
					v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
					v229 = v225<<(uint(int32(16))%32) + v224
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v237 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
					v329 = v235 + v214
					v330 = v237 + v215
					v331 = v234
				case 10:
					v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+10)))
					v224 = v220<<(uint(int32(24))%32) + v216
					v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
					v229 = v225<<(uint(int32(16))%32) + v224
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
					v237 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
					v329 = v235 + v214
					v330 = v237 + v215
					v331 = v234
				default:
					v329 = v214
					v330 = v215
					v331 = v216
				}
			}
			v334 = int32(14)
			v336 = v330 ^ v331 - base.I32_rotl(v330, v334)
			v340 = v336 ^ v329 - base.I32_rotl(v336, int32(11))
			v344 = v340 ^ v330 - base.I32_rotl(v340, int32(25))
			v348 = v344 ^ v336 - base.I32_rotl(v344, int32(16))
			v352 = v348 ^ v340 - base.I32_rotl(v348, int32(4))
			v356 = v352 ^ v344 - base.I32_rotl(v352, v334)
			F_pfree(m, v48)
			mBase = m.M
			v367 = m.ExcPending
			if v367 != 0 {
				return int64(0)
			} else {
				v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v368 != v8 {
					F_pfree(m, v8)
					mBase = m.M
					v371 = m.ExcPending
					if v371 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v356)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v356^v348-base.I32_rotl(v356, int32(24)))
					}
				} else {
					return base.I64_extend_i32_u(v356)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v356^v348-base.I32_rotl(v356, int32(24)))
				}
			}
		}
	}
}
func F_citext_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(int32(base.Ui32(v14) >> (uint(int32(31)) % 32)))
							}
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v14) >> (uint(int32(31)) % 32)))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v14) >> (uint(int32(31)) % 32)))
						}
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v14) >> (uint(int32(31)) % 32)))
					}
				}
			}
		}
	}
}
func F_citext_pattern_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
					}
				}
			}
		}
	}
}
func F_citext_pattern_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v13 <= int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v13 <= int32(0)))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v13 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v13 <= int32(0)))
					}
				}
			}
		}
	}
}
