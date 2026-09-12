package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bloom_add_element(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v435 int32
	_ = v435
	var __phi435 int32
	_ = __phi435
	var v436 int32
	_ = v436
	var __phi436 int32
	_ = __phi436
	var v437 int32
	_ = v437
	var __phi437 int32
	_ = __phi437
	var v438 int32
	_ = v438
	var __phi438 int32
	_ = __phi438
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = l2 - int32(1636608432)
	if v18 == int64(0) {
		v61 = v24
		v63 = v24
		v65 = v24
	} else {
		v28 = v24 + base.I32_wrap_i64(v18)
		v29 = v28 + v24
		v33 = int32(4)
		v35 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ base.I32_rotl(v24, v33)
		v39 = v28 - v35 ^ base.I32_rotl(v35, int32(6))
		v43 = v29 - v39 ^ base.I32_rotl(v39, int32(8))
		v44 = v35 + v29
		v45 = v39 + v44
		v46 = v43 + v45
		v50 = v44 - v43 ^ base.I32_rotl(v43, int32(16))
		v54 = v45 - v50 ^ base.I32_rotl(v50, int32(19))
		v59 = v50 + v46
		v61 = v59
		v63 = v46 - v54 ^ base.I32_rotl(v54, v33)
		v65 = v54 + v59
	}
	if l1&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(l2) {
			v70 = l1
			v71 = l2
			v73 = v61
			v74 = v65
			v75 = v63
			for {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
				v78 = v77 + v74
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
				v82 = v81 + v75
				v84 = int32(4)
				v86 = v79 + v73 - v82 ^ base.I32_rotl(v82, v84)
				v90 = v78 - v86 ^ base.I32_rotl(v86, int32(6))
				v91 = v82 + v78
				v92 = v86 + v91
				v93 = v90 + v92
				v97 = v91 - v90 ^ base.I32_rotl(v90, int32(8))
				v101 = v92 - v97 ^ base.I32_rotl(v97, int32(16))
				v105 = v93 - v101 ^ base.I32_rotl(v101, int32(19))
				v106 = v97 + v93
				v107 = v101 + v106
				v108 = v105 + v107
				v112 = v106 - v105 ^ base.I32_rotl(v105, v84)
				v113 = int32(12)
				v114 = v70 + v113
				v116 = v71 - v113
				if base.Ui32(int32(11)) < base.Ui32(v116) {
					v70 = v114
					v71 = v116
					v73 = v107
					v74 = v108
					v75 = v112
					continue
				} else {
					break
				}
				break
			}
			v119 = v114
			v120 = v116
			v122 = v107
			v123 = v108
			v124 = v112
		} else {
			v119 = l1
			v120 = l2
			v122 = v61
			v123 = v65
			v124 = v63
		}
		switch v120 - int32(1) {
		case 0:
			v289 = v122
			v290 = v123
			v291 = v124
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 1:
			v282 = v122
			v283 = v123
			v284 = v124
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 2:
			v275 = v122
			v276 = v123
			v277 = v124
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 3:
			v269 = v123
			v270 = v124
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 4:
			v265 = v123
			v266 = v124
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 5:
			v259 = v123
			v260 = v124
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 6:
			v253 = v123
			v254 = v124
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
			v259 = v255<<(uint(int32(16))%32) + v253
			v260 = v254
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 7:
			v248 = v124
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
			v253 = v249<<(uint(int32(24))%32) + v123
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
			v259 = v255<<(uint(int32(16))%32) + v253
			v260 = v254
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 8:
			v243 = v124
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
			v248 = v244<<(uint(int32(8))%32) + v243
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
			v253 = v249<<(uint(int32(24))%32) + v123
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
			v259 = v255<<(uint(int32(16))%32) + v253
			v260 = v254
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 9:
			v238 = v124
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+9)))
			v243 = v239<<(uint(int32(16))%32) + v238
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
			v248 = v244<<(uint(int32(8))%32) + v243
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
			v253 = v249<<(uint(int32(24))%32) + v123
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
			v259 = v255<<(uint(int32(16))%32) + v253
			v260 = v254
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		case 10:
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+10)))
			v238 = v234<<(uint(int32(24))%32) + v124
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+9)))
			v243 = v239<<(uint(int32(16))%32) + v238
			v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
			v248 = v244<<(uint(int32(8))%32) + v243
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+7)))
			v253 = v249<<(uint(int32(24))%32) + v123
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+6)))
			v259 = v255<<(uint(int32(16))%32) + v253
			v260 = v254
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+5)))
			v265 = v261<<(uint(int32(8))%32) + v259
			v266 = v260
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
			v269 = v265 + v267
			v270 = v266
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
			v275 = v271<<(uint(int32(24))%32) + v122
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
			v282 = v278<<(uint(int32(16))%32) + v275
			v283 = v276
			v284 = v277
			v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
			v289 = v285<<(uint(int32(8))%32) + v282
			v290 = v283
			v291 = v284
			v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
			v297 = v289 + v292
			v298 = v290
			v299 = v291
		default:
			v297 = v122
			v298 = v123
			v299 = v124
		}
	} else {
		if base.Ui32(int32(12)) <= base.Ui32(l2) {
			v130 = l1
			v131 = l2
			v133 = v61
			v134 = v65
			v135 = v63
			for {
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
				v138 = v137 + v134
				v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				v142 = v141 + v135
				v144 = int32(4)
				v146 = v139 + v133 - v142 ^ base.I32_rotl(v142, v144)
				v150 = v138 - v146 ^ base.I32_rotl(v146, int32(6))
				v151 = v142 + v138
				v152 = v146 + v151
				v153 = v150 + v152
				v157 = v151 - v150 ^ base.I32_rotl(v150, int32(8))
				v161 = v152 - v157 ^ base.I32_rotl(v157, int32(16))
				v165 = v153 - v161 ^ base.I32_rotl(v161, int32(19))
				v166 = v157 + v153
				v167 = v161 + v166
				v168 = v165 + v167
				v172 = v166 - v165 ^ base.I32_rotl(v165, v144)
				v173 = int32(12)
				v174 = v130 + v173
				v176 = v131 - v173
				if base.Ui32(int32(11)) < base.Ui32(v176) {
					v130 = v174
					v131 = v176
					v133 = v167
					v134 = v168
					v135 = v172
					continue
				} else {
					break
				}
				break
			}
			v179 = v174
			v180 = v176
			v182 = v167
			v183 = v168
			v184 = v172
		} else {
			v179 = l1
			v180 = l2
			v182 = v61
			v183 = v65
			v184 = v63
		}
		switch v180 - int32(1) {
		case 0:
			v231 = v182
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
			v297 = v231 + v232
			v298 = v183
			v299 = v184
		case 1:
			v226 = v182
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
			v231 = v227<<(uint(int32(8))%32) + v226
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
			v297 = v231 + v232
			v298 = v183
			v299 = v184
		case 2:
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)))
			v226 = v222<<(uint(int32(16))%32) + v182
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
			v231 = v227<<(uint(int32(8))%32) + v226
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
			v297 = v231 + v232
			v298 = v183
			v299 = v184
		case 3:
			v219 = v183
			v220 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v297 = v220 + v182
			v298 = v219
			v299 = v184
		case 4:
			v216 = v183
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
			v219 = v216 + v217
			v220 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v297 = v220 + v182
			v298 = v219
			v299 = v184
		case 5:
			v211 = v183
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)))
			v216 = v212<<(uint(int32(8))%32) + v211
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
			v219 = v216 + v217
			v220 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v297 = v220 + v182
			v298 = v219
			v299 = v184
		case 6:
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+6)))
			v211 = v207<<(uint(int32(16))%32) + v183
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)))
			v216 = v212<<(uint(int32(8))%32) + v211
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
			v219 = v216 + v217
			v220 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v297 = v220 + v182
			v298 = v219
			v299 = v184
		case 7:
			v202 = v184
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
			v297 = v203 + v182
			v298 = v205 + v183
			v299 = v202
		case 8:
			v197 = v184
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
			v297 = v203 + v182
			v298 = v205 + v183
			v299 = v202
		case 9:
			v192 = v184
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+9)))
			v197 = v193<<(uint(int32(16))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
			v297 = v203 + v182
			v298 = v205 + v183
			v299 = v202
		case 10:
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+10)))
			v192 = v188<<(uint(int32(24))%32) + v184
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+9)))
			v197 = v193<<(uint(int32(16))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+8)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
			v297 = v203 + v182
			v298 = v205 + v183
			v299 = v202
		default:
			v297 = v182
			v298 = v183
			v299 = v184
		}
	}
	v302 = int32(14)
	v304 = v298 ^ v299 - base.I32_rotl(v298, v302)
	v308 = v304 ^ v297 - base.I32_rotl(v304, int32(11))
	v312 = v308 ^ v298 - base.I32_rotl(v308, int32(25))
	v316 = v312 ^ v304 - base.I32_rotl(v312, int32(16))
	v320 = v316 ^ v308 - base.I32_rotl(v316, int32(4))
	v324 = v320 ^ v312 - base.I32_rotl(v320, v302)
	v334 = F_Int64GetDatum(m, base.I64_extend_i32_u(v324)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v316^v324-base.I32_rotl(v324, int32(24))))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		return
	} else {
		v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v338 = v336 - int32(1)
		v339 = *(*int64)(unsafe.Add(mBase, uint32(v334)))
		v341 = v338 & base.I32_wrap_i64(v339)
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v341
		v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v343 < int32(2) {
		} else {
			v346 = int32(1)
			v347 = v343 - v346
			v348 = int32(3)
			v349 = v347 & v348
			v352 = base.I32_wrap_i64(int64(base.Ui64(v339) >> (uint(int64(32)) % 64)))
			if base.Ui32(v348) <= base.Ui32(v343-int32(2)) {
				v363 = v346
				v364 = v352
				v365 = v341
				v366 = int32(0)
				for {
					v374 = int32(2)
					v377 = v338 & v364
					v379 = (v377 + v365) & v338
					*(*int32)(unsafe.Add(mBase, uint32(v16+v363<<(uint(v374)%32)))) = v379
					v382 = v363 + int32(1)
					v387 = (v363 + v377) & v338
					v389 = (v387 + v379) & v338
					*(*int32)(unsafe.Add(mBase, uint32(v16+v382<<(uint(v374)%32)))) = v389
					v392 = v363 + v374
					v397 = (v387 + v382) & v338
					v399 = (v397 + v389) & v338
					*(*int32)(unsafe.Add(mBase, uint32(v16+v392<<(uint(v374)%32)))) = v399
					v402 = v363 + int32(3)
					v407 = (v397 + v392) & v338
					v409 = (v407 + v399) & v338
					*(*int32)(unsafe.Add(mBase, uint32(v16+v402<<(uint(v374)%32)))) = v409
					v411 = v407 + v402
					v412 = int32(4)
					v413 = v363 + v412
					v415 = v366 + v412
					if v415 != v347&int32(-4) {
						v363 = v413
						v364 = v411
						v365 = v409
						v366 = v415
						continue
					} else {
						break
					}
					break
				}
				v419 = v413
				v420 = v411
				v421 = v409
			} else {
				v419 = v346
				v420 = v352
				v421 = v341
			}
			if v349 == int32(0) {
			} else {
				__phi435 = v419
				__phi436 = v420
				__phi437 = v421
				__phi438 = int32(0)
				v435 = __phi435
				v436 = __phi436
				v437 = __phi437
				v438 = __phi438
				for {
					v449 = v338 & v436
					v451 = (v449 + v437) & v338
					*(*int32)(unsafe.Add(mBase, uint32(v16+v435<<(uint(int32(2))%32)))) = v451
					v454 = int32(1)
					v457 = v438 + v454
					if v457 != v349 {
						__phi435 = v435 + v454
						__phi436 = v435 + v449
						__phi437 = v451
						__phi438 = v457
						v435 = __phi435
						v436 = __phi436
						v437 = __phi437
						v438 = __phi438
						continue
					} else {
						break
					}
					break
				}
			}
		}
		if v343 <= int32(0) {
		} else {
			v474 = int32(1)
			v477 = l0 + int32(24)
			v478 = int32(0)
			if v343 != v474 {
				v486 = v478
				v488 = int32(0)
				for {
					v497 = int32(2)
					v499 = v16 + v486<<(uint(v497)%32)
					v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
					v501 = int32(3)
					v503 = v477 + int32(base.Ui32(v500)>>(uint(v501)%32))
					v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
					v505 = int32(1)
					v506 = int32(7)
					v509 = v504 | v505<<(uint(v500&v506)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v509)
					v511 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
					v514 = v477 + int32(base.Ui32(v511)>>(uint(v501)%32))
					v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
					v520 = v515 | v505<<(uint(v511&v506)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v514))) = uint8(v520)
					v523 = v486 + v497
					v525 = v488 + v497
					if v525 != v343&int32(2147483646) {
						v486 = v523
						v488 = v525
						continue
					} else {
						break
					}
					break
				}
				v529 = v523
			} else {
				v529 = v478
			}
			if v343&v474 == int32(0) {
			} else {
				v545 = *(*int32)(unsafe.Add(mBase, uint32(v16+v529<<(uint(int32(2))%32))))
				v548 = v477 + int32(base.Ui32(v545)>>(uint(int32(3))%32))
				v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
				v554 = v549 | int32(1)<<(uint(v545&int32(7))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v548))) = uint8(v554)
			}
		}
		m.G0 = v16 + int32(48)
		return
	}
}
func F_initBloomState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)) = v8
	if int32(0) < v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v62 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v24 = l0 + v17*int32(28)
	v25 = int32(1)
	v26 = v17 + v25
	v29 = F_index_getprocinfo(m, l1, base.I32_extend16_s(v26), v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v34 = v24 + int32(16)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
	goto L8
L8:
	;
	v47 = v17 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49+v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(896)+v47))) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v26 < v54 {
		v17 = v26
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L37
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L34
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v67 = F_MemoryContextAlloc(m, v65, int32(136))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v119 = v62
	goto L14
L14:
	;
	goto L31
L15:
	;
	v70 = F_ReadBuffer(m, l1, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_LockBuffer(m, v70, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v70 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v108 != int32(-609481235) {
		goto L11
	} else {
		goto L24
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+(v70^int32(-1))<<(uint(int32(2))%32))))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85)+2)))
	if v87&int32(1) != 0 {
		v106 = v84
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v94 = v91 + v70<<(uint(int32(13))%32)
	v96 = v94 + int32(-8192)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94-int32(8176)))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v99)+2)))
	if v101&int32(1) == int32(0) {
		goto L10
	} else {
		goto L23
	}
L22:
	;
	goto L10
L23:
	;
	v106 = v96
	goto L18
L24:
	;
	goto L26
L25:
	;
	F_UnlockReleaseBuffer(m, v70)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L29
	}
L26:
	;
	v114 = F__emscripten_memcpy_bulkmem(m, v67, v106+int32(32), int32(136))
	mBase = m.M
	goto L28
L28:
	;
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+256)) = v114
	v119 = v67
	goto L14
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1028))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1164)) = v128<<(uint(int32(1))%32) + int32(6)
	return
L31:
	;
	v126 = F__emscripten_memcpy_bulkmem(m, l0+int32(1024), v119, int32(136))
	mBase = m.M
	goto L33
L33:
	;
	goto L30
L34:
	;
	F_errmsg_internal(m, int32(29570), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(518894), int32(201), int32(372323))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errmsg_internal(m, int32(29570), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(518894), int32(197), int32(372323))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
