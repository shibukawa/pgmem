package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bloom_add_element(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
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
	var v439 int32
	_ = v439
	var __phi439 int32
	_ = __phi439
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = int32(-1636608428)
	if v19 == int64(0) {
		v62 = v25
		v64 = v25
		v66 = v25
	} else {
		v28 = base.I32_wrap_i64(v19)
		v30 = v28 + int32(1021750440)
		v34 = int32(4)
		v36 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(32))%64))) ^ base.I32_rotl(v25, v34)
		v40 = v25 + v28 - v36 ^ base.I32_rotl(v36, int32(6))
		v44 = v30 - v40 ^ base.I32_rotl(v40, int32(8))
		v45 = v36 + v30
		v46 = v40 + v45
		v47 = v44 + v46
		v51 = v45 - v44 ^ base.I32_rotl(v44, int32(16))
		v55 = v46 - v51 ^ base.I32_rotl(v51, int32(19))
		v60 = v51 + v47
		v62 = v60
		v64 = v47 - v55 ^ base.I32_rotl(v55, v34)
		v66 = v55 + v60
	}
	if l1&int32(3) != 0 {
		switch int32(3) {
		case 0:
			v290 = v62
			v291 = v66
			v292 = v64
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 1:
			v283 = v62
			v284 = v66
			v285 = v64
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 2:
			v276 = v62
			v277 = v66
			v278 = v64
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 3:
			v270 = v66
			v271 = v64
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 4:
			v266 = v66
			v267 = v64
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 5:
			v260 = v66
			v261 = v64
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 6:
			v254 = v66
			v255 = v64
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v260 = v256<<(uint(int32(16))%32) + v254
			v261 = v255
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 7:
			v249 = v64
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			v254 = v250<<(uint(int32(24))%32) + v66
			v255 = v249
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v260 = v256<<(uint(int32(16))%32) + v254
			v261 = v255
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 8:
			v244 = v64
			v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v249 = v245<<(uint(int32(8))%32) + v244
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			v254 = v250<<(uint(int32(24))%32) + v66
			v255 = v249
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v260 = v256<<(uint(int32(16))%32) + v254
			v261 = v255
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 9:
			v239 = v64
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
			v244 = v240<<(uint(int32(16))%32) + v239
			v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v249 = v245<<(uint(int32(8))%32) + v244
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			v254 = v250<<(uint(int32(24))%32) + v66
			v255 = v249
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v260 = v256<<(uint(int32(16))%32) + v254
			v261 = v255
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		case 10:
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
			v239 = v235<<(uint(int32(24))%32) + v64
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
			v244 = v240<<(uint(int32(16))%32) + v239
			v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v249 = v245<<(uint(int32(8))%32) + v244
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			v254 = v250<<(uint(int32(24))%32) + v66
			v255 = v249
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v260 = v256<<(uint(int32(16))%32) + v254
			v261 = v255
			v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v266 = v262<<(uint(int32(8))%32) + v260
			v267 = v261
			v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v270 = v266 + v268
			v271 = v267
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			v276 = v272<<(uint(int32(24))%32) + v62
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v283 = v279<<(uint(int32(16))%32) + v276
			v284 = v277
			v285 = v278
			v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v290 = v286<<(uint(int32(8))%32) + v283
			v291 = v284
			v292 = v285
			v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v290 + v293
			v299 = v291
			v300 = v292
		default:
			v298 = v62
			v299 = v66
			v300 = v64
		}
	} else {
		switch int32(3) {
		case 0:
			v232 = v62
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v232 + v233
			v299 = v66
			v300 = v64
		case 1:
			v227 = v62
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v232 = v228<<(uint(int32(8))%32) + v227
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v232 + v233
			v299 = v66
			v300 = v64
		case 2:
			v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v227 = v223<<(uint(int32(16))%32) + v62
			v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v232 = v228<<(uint(int32(8))%32) + v227
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v298 = v232 + v233
			v299 = v66
			v300 = v64
		case 3:
			v220 = v66
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v298 = v221 + v62
			v299 = v220
			v300 = v64
		case 4:
			v217 = v66
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v220 = v217 + v218
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v298 = v221 + v62
			v299 = v220
			v300 = v64
		case 5:
			v212 = v66
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v220 = v217 + v218
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v298 = v221 + v62
			v299 = v220
			v300 = v64
		case 6:
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
			v212 = v208<<(uint(int32(16))%32) + v66
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
			v220 = v217 + v218
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v298 = v221 + v62
			v299 = v220
			v300 = v64
		case 7:
			v203 = v64
			v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v298 = v204 + v62
			v299 = v206 + v66
			v300 = v203
		case 8:
			v198 = v64
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v298 = v204 + v62
			v299 = v206 + v66
			v300 = v203
		case 9:
			v193 = v64
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
			v198 = v194<<(uint(int32(16))%32) + v193
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v298 = v204 + v62
			v299 = v206 + v66
			v300 = v203
		case 10:
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
			v193 = v189<<(uint(int32(24))%32) + v64
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
			v198 = v194<<(uint(int32(16))%32) + v193
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v298 = v204 + v62
			v299 = v206 + v66
			v300 = v203
		default:
			v298 = v62
			v299 = v66
			v300 = v64
		}
	}
	v303 = int32(14)
	v305 = v299 ^ v300 - base.I32_rotl(v299, v303)
	v309 = v305 ^ v298 - base.I32_rotl(v305, int32(11))
	v313 = v309 ^ v299 - base.I32_rotl(v309, int32(25))
	v317 = v313 ^ v305 - base.I32_rotl(v313, int32(16))
	v321 = v317 ^ v309 - base.I32_rotl(v317, int32(4))
	v325 = v321 ^ v313 - base.I32_rotl(v321, v303)
	v335 = F_Int64GetDatum(m, base.I64_extend_i32_u(v325)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v317^v325-base.I32_rotl(v325, int32(24))))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		return
	} else {
		v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v339 = v337 - int32(1)
		v340 = *(*int64)(unsafe.Add(mBase, uint32(v335)))
		v342 = v339 & base.I32_wrap_i64(v340)
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v342
		v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v344 < int32(2) {
		} else {
			v347 = int32(1)
			v348 = v344 - v347
			v349 = int32(3)
			v350 = v348 & v349
			v353 = base.I32_wrap_i64(int64(base.Ui64(v340) >> (uint(int64(32)) % 64)))
			if base.Ui32(v349) <= base.Ui32(v344-int32(2)) {
				v364 = v347
				v365 = v353
				v366 = v342
				v367 = int32(0)
				for {
					v375 = int32(2)
					v378 = v339 & v365
					v380 = (v378 + v366) & v339
					*(*int32)(unsafe.Add(mBase, uint32(v16+v364<<(uint(v375)%32)))) = v380
					v383 = v364 + int32(1)
					v388 = (v364 + v378) & v339
					v390 = (v388 + v380) & v339
					*(*int32)(unsafe.Add(mBase, uint32(v16+v383<<(uint(v375)%32)))) = v390
					v393 = v364 + v375
					v398 = (v388 + v383) & v339
					v400 = (v398 + v390) & v339
					*(*int32)(unsafe.Add(mBase, uint32(v16+v393<<(uint(v375)%32)))) = v400
					v403 = v364 + int32(3)
					v408 = (v398 + v393) & v339
					v410 = (v408 + v400) & v339
					*(*int32)(unsafe.Add(mBase, uint32(v16+v403<<(uint(v375)%32)))) = v410
					v412 = v408 + v403
					v413 = int32(4)
					v414 = v364 + v413
					v416 = v367 + v413
					if v416 != v348&int32(-4) {
						v364 = v414
						v365 = v412
						v366 = v410
						v367 = v416
						continue
					} else {
						break
					}
					break
				}
				v420 = v414
				v421 = v412
				v422 = v410
			} else {
				v420 = v347
				v421 = v353
				v422 = v342
			}
			if v350 == int32(0) {
			} else {
				__phi436 = v420
				__phi437 = v421
				__phi438 = v422
				__phi439 = int32(0)
				v436 = __phi436
				v437 = __phi437
				v438 = __phi438
				v439 = __phi439
				for {
					v450 = v339 & v437
					v452 = (v450 + v438) & v339
					*(*int32)(unsafe.Add(mBase, uint32(v16+v436<<(uint(int32(2))%32)))) = v452
					v455 = int32(1)
					v458 = v439 + v455
					if v458 != v350 {
						__phi436 = v436 + v455
						__phi437 = v436 + v450
						__phi438 = v452
						__phi439 = v458
						v436 = __phi436
						v437 = __phi437
						v438 = __phi438
						v439 = __phi439
						continue
					} else {
						break
					}
					break
				}
			}
		}
		if v344 <= int32(0) {
		} else {
			v475 = int32(1)
			v478 = l0 + int32(24)
			v479 = int32(0)
			if v344 != v475 {
				v487 = v479
				v489 = int32(0)
				for {
					v498 = int32(2)
					v500 = v16 + v487<<(uint(v498)%32)
					v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
					v502 = int32(3)
					v504 = v478 + int32(base.Ui32(v501)>>(uint(v502)%32))
					v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
					v506 = int32(1)
					v507 = int32(7)
					v510 = v505 | v506<<(uint(v501&v507)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v504))) = uint8(v510)
					v512 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
					v515 = v478 + int32(base.Ui32(v512)>>(uint(v502)%32))
					v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
					v521 = v516 | v506<<(uint(v512&v507)%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v515))) = uint8(v521)
					v524 = v487 + v498
					v526 = v489 + v498
					if v526 != v344&int32(2147483646) {
						v487 = v524
						v489 = v526
						continue
					} else {
						break
					}
					break
				}
				v530 = v524
			} else {
				v530 = v479
			}
			if v344&v475 == int32(0) {
			} else {
				v546 = *(*int32)(unsafe.Add(mBase, uint32(v16+v530<<(uint(int32(2))%32))))
				v549 = v478 + int32(base.Ui32(v546)>>(uint(int32(3))%32))
				v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
				v555 = v550 | int32(1)<<(uint(v546&int32(7))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v549))) = uint8(v555)
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
	F_errmsg_internal(m, int32(28524), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(493634), int32(201), int32(353584))
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
	F_errmsg_internal(m, int32(28524), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(493634), int32(197), int32(353584))
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
