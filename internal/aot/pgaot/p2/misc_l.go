package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LagTrackerRead(m *base.Module, l0 int32, l1 int64, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v104 int64
	_ = v104
	var v116 int32
	_ = v116
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v151 int64
	_ = v151
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v171 int64
	_ = v171
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_LagTrackerRead[0]))
	v19 = v16 + l0<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_LagTrackerRead[1])))
	if v22 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v49 == v50 {
		v104 = v51
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_LagTrackerRead[2])))
	v49 = v22
	v50 = v25
	v51 = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = v16 + l0<<(uint(int32(4))%32)
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[3])))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[4])))
	if base.Ui64(l1) < base.Ui64(v30) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 < v29 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[5]))) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LagTrackerRead[6]))) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_LagTrackerRead[2])))
	v47 = base.I32_rem_s(v43+int32(1), int32(_a_F_LagTrackerRead_0))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_LagTrackerRead[1]))) = v47
	v49 = v47
	v50 = v43
	v51 = v29
	goto L1
L8:
	;
	v35 = int64(-1)
	goto L10
L9:
	;
	v35 = l2 - v29
	goto L10
L10:
	;
	return v35
L11:
	;
	v127 = int64(-1)
	if l2 < v123 {
		v171 = v127
		goto L19
	} else {
		goto L20
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16+l0<<(uint(int32(4))%32))+uint32(_c_F_LagTrackerRead[5]))) = int64(0)
	v116 = v50
	v123 = v104
	goto L11
L13:
	;
	v54 = v16 + int32(8)
	v57 = v54 + v49<<(uint(int32(4))%32)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	if base.Ui64(l1) < base.Ui64(v58) {
		v116 = v49
		v123 = v51
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v62 = v16 + l0<<(uint(int32(4))%32)
	v68 = v49
	v71 = v57
	goto L15
L15:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_LagTrackerRead[5]))) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_LagTrackerRead[6]))) = v81
	v86 = base.I32_rem_s(v68+int32(1), int32(_a_F_LagTrackerRead_0))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_LagTrackerRead[1]))) = v86
	if v86 == v50 {
		v104 = v79
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v116 = v86
	v123 = v79
	goto L11
L17:
	;
	v91 = v54 + v86<<(uint(int32(4))%32)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
	if base.Ui64(v92) <= base.Ui64(l1) {
		v68 = v86
		v71 = v91
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	return v171
L20:
	;
	if v123 != int64(0) {
		v164 = v123
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v171 = l2 - v164
	goto L19
L22:
	;
	if v116 == v50 {
		v171 = v127
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v134 = v16 + l0<<(uint(int32(4))%32)
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v134)+uint32(_c_F_LagTrackerRead[5])))
	if v135 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v134)+uint32(_c_F_LagTrackerRead[6])))
	if base.Ui64(l1) < base.Ui64(v140) {
		v171 = v127
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v16+v116<<(uint(int32(4))%32))+16))
	v164 = v162
	goto L21
L27:
	;
	v144 = v16 + v116<<(uint(int32(4))%32)
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)+16))
	if v145 < v135 {
		v171 = v127
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v144)+8))
	v164 = base.I64_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v145-v135), base.F64_div(base.F64_convert_i64_u(l1-v140), base.F64_convert_i64_u(v151-v140))), base.F64_convert_i64_s(v135)))
	goto L21
}
func F_LockManagerShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LockManagerShmemInit[0]))
	v4 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3))), uint32(v4))
	return
}
func F_LockRelease(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v236 int64
	_ = v236
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v261 int32
	_ = v261
	var v266 int64
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v22-int32(3))&int32(255)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L11
	} else {
		goto L95
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L11
	} else {
		goto L92
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L11
	} else {
		goto L89
	}
L4:
	;
	if l1 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L86
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_LockRelease[0])))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 < l1 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+96)) = v36
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = l1
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[1]))
	v48 = F_hash_search(m, v43, v20+int32(88), v41, v41)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	m.G0 = v20 + int32(112)
	return v454
L10:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[2]))
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	return int32(0)
L12:
	;
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
	if int64(0) < v52 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v57 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v57 == int32(0) {
		v454 = v41
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v65
	F_errmsg_internal(m, int32(_a_F_LockRelease_0), v20+int32(32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2183), int32(_a_F_LockRelease_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v454 = v41
	goto L9
L21:
	;
	v80 = int32(0)
	goto L23
L22:
	;
	v80 = v79
	goto L23
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v86 = v81
	goto L28
L24:
	;
	F_RemoveLocalLock(m, v48)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L11
	} else {
		goto L85
	}
L25:
	;
	v360 = int32(1) << (uint(l1) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	if v360&v361 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L26:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[3]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v335 = int32(0)
	v337 = F_hash_search_with_hash_value(m, v333, l0, v334, v335, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L65
	}
L27:
	;
	v309 = int32(0)
	v312 = F_errstart(m, int32(19), v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L11
	} else {
		goto L61
	}
L28:
	;
	v101 = v86 - int32(1)
	if v101 < int32(0) {
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v106)+8))
	v111 = v109 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v111
	if v111 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v106 = v82 + v101<<(uint(int32(4))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v107 != v80 {
		v86 = v101
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
	v133 = v131 - int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v133
	if int64(0) < v133 {
		v454 = int32(1)
		goto L9
	} else {
		goto L39
	}
L33:
	;
	if v80 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ResourceOwnerForgetLock(m, v79, v48)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	v119 = v117 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+40)) = v119
	if v119 <= v101 {
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v124 = v82 + v119<<(uint(int32(4))%32)
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v124)))
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v127
	goto L32
L39:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+53)) = uint8(v138)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v140 != int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[4]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v301 = v293 + v294&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockRelease_3)
	v303 = F_LWLockAcquire(m, v301, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L11
	} else {
		goto L59
	}
L41:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v143|base.B2i32(int32(3) < l1) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[5]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v148 != v149)|base.B2i32(v148 == int32(0)) != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[6]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32((v155-int32(1))&(v158*int32(_a_F_LockRelease_4))<<(uint(int32(2))%32))+uint32(_c_F_LockRelease[7])))
	if v164 <= int32(0) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[8]))
	v172 = F_LWLockAcquire(m, v168+int32(548), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v175 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[6]))
	v178 = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v183 = (v177 - v178) & (v180 * int32(_a_F_LockRelease_4))
	v185 = v183 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+uint32(_c_F_LockRelease[7]))) = v175
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[8]))
	v199 = v183 & int32(268435455) << (uint(int32(3)) % 32)
	v202 = v195
	v203 = v175
	v213 = int64(0)
	goto L46
L46:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v202)+568))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217+v183<<(uint(int32(6))%32)+base.I32_wrap_i64(v213)<<(uint(int32(2))%32))))
	if v180 != v223 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[8]))
	F_LWLockRelease(m, v270+int32(548))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L11
	} else {
		goto L57
	}
L48:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247)+564))
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v253+v199)))
	if int64(base.Ui64(v255)>>(uint(v250)%64))&int64(7) != int64(0) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v247 = v202
	v248 = v203
	v250 = v213 * int64(3)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v202)+564))
	v228 = v227 + v199
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	v232 = v213 * int64(3)
	v236 = int64(1) << (uint(base.I64_extend_i32_u(l1-v178+base.I32_wrap_i64(v232))) % 64)
	if v229&v236 == int64(0) {
		v247 = v202
		v248 = v203
		v250 = v232
		goto L48
	} else {
		goto L52
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v229 & (v236 ^ int64(-1))
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[8]))
	v247 = v245
	v248 = int32(1)
	v250 = v232
	goto L48
L53:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v185)+uint32(_c_F_LockRelease[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+uint32(_c_F_LockRelease[7]))) = v261 + int32(1)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v266 = v213 + int64(1)
	if v266 != int64(16) {
		v202 = v247
		v203 = v248
		v213 = v266
		goto L46
	} else {
		goto L56
	}
L56:
	;
	goto L47
L57:
	;
	if v248 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	goto L40
L59:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v305 == int32(0) {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	v357 = v308
	v358 = v305
	goto L25
L61:
	;
	if v312 == int32(0) {
		v454 = v309
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v316+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v320
	F_errmsg_internal(m, int32(_a_F_LockRelease_0), v20-int32(-64))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2222), int32(_a_F_LockRelease_2))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v454 = v309
	goto L9
L65:
	;
	if v337 == int32(0) {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v337
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v344
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelease[9]))
	v350 = int32(0)
	v352 = F_hash_search(m, v347, v20+int32(80), v350, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v352
	if v352 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v357 = v352
	v358 = v337
	goto L25
L69:
	;
	F_LWLockRelease(m, v301)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L11
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v358)+84))
	v391 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+84)) = v390 - v391
	v395 = l1 << (uint(int32(2)) % 32)
	v396 = v358 + v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+44)) = v397 - v391
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v358)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+128)) = v401 - v391
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v396)+88))
	v407 = v405 - v391
	*(*int32)(unsafe.Add(mBase, uint32(v396)+88)) = v407
	v410 = v360 ^ int32(-1)
	if v407 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	v367 = int32(0)
	v370 = F_errstart(m, int32(19), v367)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	if v370 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v372+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v376
	F_errmsg_internal(m, int32(_a_F_LockRelease_0), v20+int32(48))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L11
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_RemoveLocalLock(m, v48)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2318), int32(_a_F_LockRelease_2))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v454 = v367
	goto L9
L80:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v358)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v413 & v410
	goto L82
L81:
	;
	goto L82
L82:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v416+v395)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v358)+20))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v357)+12)) = v420 & v410
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	F_CleanUpLock(m, v358, v357, v33, v423, base.B2i32(v418&v419 != int32(0)))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_LWLockRelease(m, v301)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	goto L24
L85:
	;
	v454 = int32(1)
	goto L9
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	F_errmsg_internal(m, int32(_a_F_LockRelease_5), v20)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2154), int32(_a_F_LockRelease_2))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_LockRelease_6), v20+int32(16))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2157), int32(_a_F_LockRelease_2))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errmsg_internal(m, int32(_a_F_LockRelease_7), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2293), int32(_a_F_LockRelease_2))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errmsg_internal(m, int32(_a_F_LockRelease_8), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_LockRelease_1), int32(2303), int32(_a_F_LockRelease_2))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LruDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l0 * int32(48)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[0]))
	v13 = v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	F_pgaio_closing_fd(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v18 = F_close(m, v17)
		mBase = m.M
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
			v48 = int32(_a_F_LruDelete_0)
			v50 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1])) = v50 - int32(1)
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[0]))
			v56 = v55 + v10
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
			v58 = int32(48)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v55+v57*v58)+16)) = v61
			*(*int32)(unsafe.Add(mBase, uint32(v55+v61*v58)+20)) = v57
			m.G0 = v7 + int32(16)
			return
		} else {
			v21 = int32(15)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LruDelete[2])))
			if v25 != 0 {
				v26 = v21
			} else {
				v26 = int32(24)
			}
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
			if v27&int32(4) != 0 {
				v30 = v21
			} else {
				v30 = v26
			}
			v32 = F_errstart(m, v30, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
					v48 = int32(_a_F_LruDelete_0)
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1])) = v50 - int32(1)
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[0]))
					v56 = v55 + v10
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
					v58 = int32(48)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v55+v57*v58)+16)) = v61
					*(*int32)(unsafe.Add(mBase, uint32(v55+v61*v58)+20)) = v57
					m.G0 = v7 + int32(16)
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36
					F_errmsg_internal(m, int32(_a_F_LruDelete_1), v7)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_LruDelete_2), int32(1297), int32(_a_F_LruDelete_3))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(-1)
							v48 = int32(_a_F_LruDelete_0)
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_LruDelete[1])) = v50 - int32(1)
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_LruDelete[0]))
							v56 = v55 + v10
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							v58 = int32(48)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v55+v57*v58)+16)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(v55+v61*v58)+20)) = v57
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = m.Wasi_snapshot_preview1.Fd_seek(m, l0, l1, l2&int32(255), v7+int32(8))
	mBase = m.M
	if v13 == int32(0) {
		v20 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___lseek[0])) = v13
		v20 = int32(-1)
	}
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	m.G0 = v7 + int32(16)
	if v20 != 0 {
		v26 = int64(-1)
	} else {
		v26 = v21
	}
	return v26
}
func F__ltq_rregex(m *base.Module, l0 int32) int64 {
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
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltq_rregex_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_l2_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 float32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 float32
	_ = v59
	var v61 float32
	_ = v61
	var v62 float32
	_ = v62
	var v65 float32
	_ = v65
	var v67 float32
	_ = v67
	var v68 float32
	_ = v68
	var v71 float32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v91 float32
	_ = v91
	var v94 int32
	_ = v94
	var v96 float32
	_ = v96
	var v98 float32
	_ = v98
	var v99 float32
	_ = v99
	var v112 float32
	_ = v112
	var v129 int64
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v10 = float32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			if v25 == v26 {
				v28 = base.I32_extend16_s(v25)
				if v28 <= int32(0) {
					v129 = int64(0)
				} else {
					v32 = int32(8)
					v33 = v23 + v32
					v35 = v18 + v32
					if v28 == int32(1) {
						v81 = int32(0)
						v91 = v10
						v94 = v81 << (uint(int32(2)) % 32)
						v96 = *(*float32)(unsafe.Add(mBase, uint32(v35+v94)))
						v98 = *(*float32)(unsafe.Add(mBase, uint32(v94+v33)))
						v99 = base.F32_sub(v96, v98)
						v112 = base.F32_add(base.F32_mul(v99, v99), v91)
					} else {
						v42 = int32(0)
						v50 = int32(0)
						v52 = v10
						for {
							v54 = int32(2)
							v55 = v42 << (uint(v54) % 32)
							v57 = v55 | int32(4)
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v35+v57)))
							v61 = *(*float32)(unsafe.Add(mBase, uint32(v33+v57)))
							v62 = base.F32_sub(v59, v61)
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v35+v55)))
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v33+v55)))
							v68 = base.F32_sub(v65, v67)
							v71 = base.F32_add(base.F32_mul(v62, v62), base.F32_add(base.F32_mul(v68, v68), v52))
							v73 = v42 + v54
							v75 = v50 + v54
							if v75 != v28&int32(_a_F_l2_distance_0) {
								v42 = v73
								v50 = v75
								v52 = v71
								continue
							} else {
								break
							}
							break
						}
						if v28&int32(1) == int32(0) {
							v112 = v71
						} else {
							v81 = v73
							v91 = v71
							v94 = v81 << (uint(int32(2)) % 32)
							v96 = *(*float32)(unsafe.Add(mBase, uint32(v35+v94)))
							v98 = *(*float32)(unsafe.Add(mBase, uint32(v94+v33)))
							v99 = base.F32_sub(v96, v98)
							v112 = base.F32_add(base.F32_mul(v99, v99), v91)
						}
					}
					v129 = base.I64_reinterpret_f64(base.F64_sqrt(base.F64_promote_f32(v112)))
				}
				m.G0 = v15 + int32(16)
				return v129
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int64(0)
					} else {
						v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v142
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v141
						F_errmsg(m, int32(_a_F_l2_distance_1), v15)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_l2_distance_2), int32(76), int32(_a_F_l2_distance_3))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_lca(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v10 = F_palloc_mul(m, int32(4), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = int32(0)
	goto L6
L4:
	;
	v42 = v14
	goto L5
L5:
	;
	v47 = F_lca_inner(m, v10, v42)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v20<<(uint(int32(4))%32))))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v42 = v38
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+v20<<(uint(int32(2))%32)))) = v33
	v37 = v20 + int32(1)
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v37 < v38 {
		v20 = v37
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v49 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = int32(0)
	v57 = v49
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_pfree(m, v10)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L21
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10+v56<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v56<<(uint(int32(4))%32))))
	if v65 != v69 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_pfree(m, v65)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v74 = v57
	goto L18
L18:
	;
	v76 = v56 + int32(1)
	if v76 < base.I32_extend16_s(v74) {
		v56 = v76
		v57 = v74
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	v74 = v73
	goto L18
L20:
	;
	goto L15
L21:
	;
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return base.I64_extend_i32_u(v47)
L23:
	;
	goto L24
L24:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v90)
	return int64(0)
}
func F_lcons_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14325(m, l0, l1, int64(4294967776))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_left_oper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+140)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v6
	v20 = F_make_oper_cache_key(m, l0, v12+int32(4), l1, v6, l2, l4)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(192)
	return v160
L2:
	;
	v69 = F_OpernameGetOprid(m, l1, int32(0), l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L18
	}
L3:
	;
	return int32(0)
L4:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_left_oper[0]))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = v27
	goto L8
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = int64(601295421576)
	v36 = F_hash_create(m, int32(_a_F_left_oper_0), int64(256), v12+int32(144), int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v54 = int32(0)
	v56 = F_hash_search(m, v51, v12+int32(4), v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_left_oper[0])) = v36
	F_CacheRegisterSyscacheCallback(m, int32(39), int32(526), int64(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(12), int32(526), int64(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_left_oper[0]))
	v51 = v50
	goto L8
L12:
	;
	if v56 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+136))
	if v60 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v65 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v60))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if v65 != 0 {
		v160 = v65
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L2
L17:
	;
	if l3 != 0 {
		v160 = int32(0)
		goto L1
	} else {
		goto L38
	}
L18:
	;
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = F_OpernameGetCandidates(m, l1, int32(108), int32(0), v12)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L22
	}
L20:
	;
	v115 = v69
	v120 = v6
	goto L21
L21:
	;
	v124 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v115))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L34
	}
L22:
	;
	if v75 == int32(0) {
		v146 = v6
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v81 = v75
	goto L24
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+32)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v90 != 0 {
		v81 = v90
		goto L24
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v75
	v97 = F_func_match_argtypes(m, int32(1), v12+int32(140), v75, v12+int32(144))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L30
	}
L26:
	;
	goto L25
L27:
	;
	v109 = int32(2)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v110 == int32(0) {
		v146 = v109
		goto L17
	} else {
		goto L33
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	v107 = v106
	goto L27
L29:
	;
	v99 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
	v104 = F_func_select_candidate(m, v99, v12+int32(140), v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	switch v97 {
	case 0:
		v146 = v97
		goto L17
	case 1:
		goto L28
	default:
		goto L29
	}
L31:
	;
	if v104 != 0 {
		v107 = v104
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v146 = v99
	goto L17
L33:
	;
	v115 = v110
	v120 = v109
	goto L21
L34:
	;
	if v124 == int32(0) {
		v146 = v120
		goto L17
	} else {
		goto L35
	}
L35:
	;
	if v20 == int32(0) {
		v160 = v124
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_left_oper[0]))
	v136 = F_hash_search(m, v131, v12+int32(4), int32(1), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+136)) = v115
	v160 = v124
	goto L1
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	F_op_error(m, l0, l1, int32(0), v150, v146, v151, l4)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_levenshtein_less_equal_with_costs(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v20 == int32(1) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v49 = v36
			} else {
				v37 = int32(1)
				if v20&v37 != 0 {
					v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
			v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
			v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v54 = int32(1)
			if v20&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			v60 = int32(1)
			if v19&v60 != 0 {
				v64 = v60
			} else {
				v64 = int32(4)
			}
			if v19 == int32(1) {
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				if v71 == int32(18) {
					v74 = int32(16)
				} else {
					v74 = int32(0)
				}
				if base.Ui32((v71-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v81 = int32(4)
				} else {
					v81 = v74
				}
				v94 = v81
			} else {
				v82 = int32(1)
				if v19&v82 != 0 {
					v94 = int32(base.Ui32(v19)>>(uint(v82)%32)) - v82
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v94 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v100 = F_varstr_levenshtein_less_equal(m, v12+v58, v49, v17+v64, v94, base.I32_wrap_i64(v53), base.I32_wrap_i64(v52), base.I32_wrap_i64(v51), base.I32_wrap_i64(v50), int32(0))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_s(v100)
			}
		}
	}
}
func F_ln_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int64
	_ = v520
	var v522 int64
	_ = v522
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L152
	}
L4:
	;
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v16
	v32 = F_palloc(m, v12<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L7
	} else {
		goto L148
	}
L7:
	;
	return
L8:
	;
	v34 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v32))) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v36 <= v34 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v32 + int32(2)
	v57 = F_palloc(m, int32(4))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v40 = v36 << (uint(int32(1)) % 32)
	if v40 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v32+int32(2), v45, v40)
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(_a_F_ln_var_0)
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_ln_var[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v62
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_ln_var[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v57 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v57
	v72 = l2 + int32(8)
	v78 = int32(0)
	goto L13
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	if v82 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v585 = v10 + int32(72)
	F_sqrt_var(m, v585, v585, v72-v81<<(uint(int32(2))%32)>>(uint(int32(1))%32))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L7
	} else {
		goto L146
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v85 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v88 = int32(1)
	v89 = int32(-1)
	v90 = int32(0)
	if base.B2i32(v89 < v81)&base.B2i32(v90 < v82) == v90 {
		v121 = v81
		v125 = v90
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v263 <= int32(0) {
		goto L15
	} else {
		goto L61
	}
L19:
	;
	v263 = v253
	goto L18
L20:
	;
	if int32(0)|base.B2i32(v89 <= v121) != 0 {
		v157 = v89
		v159 = v90
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v102 = v81
	v106 = v90
	goto L22
L22:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86+v106<<(uint(int32(1))%32)))))
	if v112 != 0 {
		v253 = int32(1)
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v121 = v116
	v125 = v114
	goto L20
L24:
	;
	v113 = int32(1)
	v114 = v106 + v113
	v116 = v102 - v113
	if v116 <= v89 {
		v121 = v116
		v125 = v114
		goto L20
	} else {
		goto L25
	}
L25:
	;
	if v114 < v82 {
		v102 = v116
		v106 = v114
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v121 != v157 {
		v199 = v125
		v200 = v159
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v138 = v89
	v140 = v90
	goto L29
L29:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140<<(uint(int32(1))%32))+uint32(_c_F_ln_var[2]))))
	if v145 != 0 {
		v253 = int32(-1)
		goto L19
	} else {
		goto L31
	}
L30:
	;
	v157 = v149
	v159 = v147
	goto L27
L31:
	;
	v146 = int32(1)
	v147 = v140 + v146
	v149 = v138 - v146
	if v149 <= v121 {
		v157 = v149
		v159 = v147
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if v147 < v88 {
		v138 = v149
		v140 = v147
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if v82 < v199 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v168 = v125
	v169 = v159
	goto L36
L36:
	;
	if base.B2i32(v82 <= v168)|base.B2i32(v88 <= v169) != 0 {
		v199 = v168
		v200 = v169
		goto L34
	} else {
		goto L38
	}
L37:
	;
	if base.I32_extend16_s(v185) < base.I32_extend16_s(v183) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v174 = int32(1)
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86+v168<<(uint(v174)%32)))))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169<<(uint(v174)%32))+uint32(_c_F_ln_var[2]))))
	if v183 == v185 {
		v168 = v168 + v174
		v169 = v169 + v174
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v192 = int32(1)
	goto L42
L41:
	;
	v192 = int32(-1)
	goto L42
L42:
	;
	v263 = v192
	goto L18
L43:
	;
	v203 = v199
	goto L45
L44:
	;
	v203 = v82
	goto L45
L45:
	;
	v210 = v199
	goto L46
L46:
	;
	if v203 == v210 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v253 = v236
	goto L19
L48:
	;
	if v88 < v200 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v236 = int32(1)
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86+v210<<(uint(v236)%32)))))
	if v242 == int32(0) {
		v210 = v210 + v236
		goto L46
	} else {
		goto L60
	}
L51:
	;
	v215 = v200
	goto L53
L52:
	;
	v215 = v88
	goto L53
L53:
	;
	v223 = v200
	goto L54
L54:
	;
	if v215 == v223 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v253 = int32(-1)
	goto L19
L56:
	;
	v263 = int32(0)
	goto L18
L57:
	;
	goto L58
L58:
	;
	v227 = int32(1)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223<<(uint(v227)%32))+uint32(_c_F_ln_var[2]))))
	if v232 == int32(0) {
		v223 = v223 + v227
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	goto L47
L61:
	;
	v266 = v81
	v270 = v78
	v272 = v82
	goto L62
L62:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v275 = int32(2)
	v276 = int32(0)
	if base.B2i32(v276 < v266)&base.B2i32(v276 < v272) == v276 {
		v308 = v266
		v312 = v276
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v479 = v10 + int32(72)
	F_sub_var(m, v479, int32(_a_F_ln_var_1), l1)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L7
	} else {
		goto L113
	}
L64:
	;
	goto L63
L65:
	;
	if v450 < int32(0) {
		v476 = v270
		goto L64
	} else {
		goto L108
	}
L66:
	;
	v450 = v440
	goto L65
L67:
	;
	if int32(0)|base.B2i32(v276 <= v308) != 0 {
		v344 = v276
		v346 = v276
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v289 = v266
	v293 = v276
	goto L69
L69:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273+v293<<(uint(int32(1))%32)))))
	if v299 != 0 {
		v440 = int32(1)
		goto L66
	} else {
		goto L71
	}
L70:
	;
	v308 = v303
	v312 = v301
	goto L67
L71:
	;
	v300 = int32(1)
	v301 = v293 + v300
	v303 = v289 - v300
	if v303 <= v276 {
		v308 = v303
		v312 = v301
		goto L67
	} else {
		goto L72
	}
L72:
	;
	if v301 < v272 {
		v289 = v303
		v293 = v301
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	if v308 != v344 {
		v386 = v312
		v387 = v346
		goto L81
	} else {
		goto L82
	}
L75:
	;
	v325 = v276
	v327 = v276
	goto L76
L76:
	;
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v327<<(uint(int32(1))%32))+uint32(_c_F_ln_var[3]))))
	if v332 != 0 {
		v440 = int32(-1)
		goto L66
	} else {
		goto L78
	}
L77:
	;
	v344 = v336
	v346 = v334
	goto L74
L78:
	;
	v333 = int32(1)
	v334 = v327 + v333
	v336 = v325 - v333
	if v336 <= v308 {
		v344 = v336
		v346 = v334
		goto L74
	} else {
		goto L79
	}
L79:
	;
	if v334 < v275 {
		v325 = v336
		v327 = v334
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v272 < v386 {
		goto L90
	} else {
		goto L91
	}
L82:
	;
	v355 = v312
	v356 = v346
	goto L83
L83:
	;
	if base.B2i32(v272 <= v355)|base.B2i32(v275 <= v356) != 0 {
		v386 = v355
		v387 = v356
		goto L81
	} else {
		goto L85
	}
L84:
	;
	if base.I32_extend16_s(v372) < base.I32_extend16_s(v370) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v361 = int32(1)
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273+v355<<(uint(v361)%32)))))
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v356<<(uint(v361)%32))+uint32(_c_F_ln_var[3]))))
	if v370 == v372 {
		v355 = v355 + v361
		v356 = v356 + v361
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v379 = int32(1)
	goto L89
L88:
	;
	v379 = int32(-1)
	goto L89
L89:
	;
	v450 = v379
	goto L65
L90:
	;
	v390 = v386
	goto L92
L91:
	;
	v390 = v272
	goto L92
L92:
	;
	v397 = v386
	goto L93
L93:
	;
	if v390 == v397 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v440 = v423
	goto L66
L95:
	;
	if v275 < v387 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v423 = int32(1)
	v429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273+v397<<(uint(v423)%32)))))
	if v429 == int32(0) {
		v397 = v397 + v423
		goto L93
	} else {
		goto L107
	}
L98:
	;
	v402 = v387
	goto L100
L99:
	;
	v402 = v275
	goto L100
L100:
	;
	v410 = v387
	goto L101
L101:
	;
	if v402 == v410 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v440 = int32(-1)
	goto L66
L103:
	;
	v450 = int32(0)
	goto L65
L104:
	;
	goto L105
L105:
	;
	v414 = int32(1)
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410<<(uint(v414)%32))+uint32(_c_F_ln_var[3]))))
	if v419 == int32(0) {
		v410 = v410 + v414
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	goto L94
L108:
	;
	v454 = v10 + int32(72)
	F_sqrt_var(m, v454, v454, v72-v266<<(uint(int32(2))%32)>>(uint(int32(1))%32))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_mul_var(m, v10, int32(_a_F_ln_var_2), v10, int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	v467 = v270 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	if v468 == int32(0) {
		v476 = v467
		goto L64
	} else {
		goto L111
	}
L111:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v472 == int32(0) {
		v266 = v471
		v270 = v467
		v272 = v468
		goto L62
	} else {
		goto L112
	}
L112:
	;
	v476 = v467
	goto L64
L113:
	;
	v485 = v10 + int32(24)
	F_add_var(m, v479, int32(_a_F_ln_var_1), v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v488 = int32(1)
	v494 = v72 + base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_s(v476+v488), float64(0.301029995663981)))
	F_div_var(m, l1, v485, l1, v494, v488, int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v504 = F_palloc(m, v499<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v506 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v504))) = uint16(v506)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v508 <= v506 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v520 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v520
	v522 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v504 + int32(2)
	F_mul_var(m, l1, l1, v10+int32(72), v494)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	v512 = v508 << (uint(int32(1)) % 32)
	if v512 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	base.MemoryCopy(m, v504+int32(2), v517, v512)
	goto L117
L120:
	;
	v532 = int32(1)
	v535 = base.I32_div_s(v494<<(uint(v532)%32), int32(-4))
	v541 = v532
	goto L121
L121:
	;
	v545 = v10 + int32(48)
	F_mul_var(m, v545, v10+int32(72), v545, v494)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L7
	} else {
		goto L124
	}
L122:
	;
	F_mul_var(m, l1, v10, l1, l2)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L7
	} else {
		goto L129
	}
L123:
	;
	goto L122
L124:
	;
	v551 = v541 + int32(2)
	v554 = v10 + int32(24)
	F_div_var_int(m, v545, v551, int32(0), v554, v494, int32(1))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v558 == int32(0) {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	F_add_var(m, l1, v554, l1)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v564+v535 <= v563 {
		v541 = v551
		goto L121
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	if v569 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_pfree(m, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L7
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	if v572 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	F_pfree(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L7
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v575 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L136
L138:
	;
	F_pfree(m, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L7
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v578 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	F_pfree(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L7
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	m.G0 = v10 + int32(96)
	return
L145:
	;
	goto L144
L146:
	;
	F_mul_var(m, v10, int32(_a_F_ln_var_2), v10, int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	v78 = v78 + int32(1)
	goto L13
L148:
	;
	F_errcode(m, int32(352583810))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_ln_var_3), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_ln_var_4), int32(_a_F_ln_var_5), int32(_a_F_ln_var_6))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L7
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(352583810))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L7
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_ln_var_7), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_ln_var_4), int32(_a_F_ln_var_8), int32(_a_F_ln_var_6))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L7
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_load_relcache_init_file(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v599 int64
	_ = v599
	var v633 int32
	_ = v633
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v730 int32
	_ = v730
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v844 int32
	_ = v844
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(1104)
	m.G0 = v19
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = F_AllocateFile(m, v19+int32(80), int32(_a_F_load_relcache_init_file_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L10
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(_a_F_load_relcache_init_file_1)
	v29 = F_pg_snprintf(m, v19+int32(80), int32(1024), int32(_a_F_load_relcache_init_file_2), v19+int32(48))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(_a_F_load_relcache_init_file_1)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v36
	v44 = F_pg_snprintf(m, v19+int32(80), int32(1024), int32(_a_F_load_relcache_init_file_3), v19-int32(-64))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L1
L7:
	;
	goto L1
L8:
	;
	m.G0 = v19 + int32(1104)
	return v844
L9:
	;
	if int32(0) < v78 {
		goto L162
	} else {
		goto L163
	}
L10:
	;
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = F_palloc(m, int32(400))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v844 = int32(0)
	goto L8
L14:
	;
	v58 = F_fread(m, v19+int32(76), int32(1), int32(4), v49)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v698)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L5
	} else {
		goto L160
	}
L16:
	;
	if v58 != int32(4) {
		v698 = v52
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v62 != int32(_a_F_load_relcache_init_file_4) {
		v698 = v52
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v75 = v52
	v76 = v2
	v77 = v2
	v78 = v2
	v79 = int32(100)
	goto L19
L19:
	;
	v86 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v76 == int32(4))&base.B2i32(v77 == int32(7)) != 0 {
		goto L9
	} else {
		goto L155
	}
L21:
	;
	goto L20
L22:
	;
	if v86 != int32(4) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v86 != 0 {
		v698 = v75
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v115 != int32(276) {
		v698 = v75
		goto L15
	} else {
		goto L33
	}
L26:
	;
	if l0 == int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	if base.B2i32(v76 == int32(6))&base.B2i32(v77 == int32(8)) != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v99 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	if v99 == int32(0) {
		v698 = v75
		goto L15
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(34359738374)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v76
	F_errmsg_internal(m, int32(_a_F_load_relcache_init_file_5), v19)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_load_relcache_init_file_6), int32(_a_F_load_relcache_init_file_7), int32(_a_F_load_relcache_init_file_8))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v698 = v75
	goto L15
L33:
	;
	if v79 <= v78 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v122 = F_repalloc(m, v75, v79<<(uint(int32(3))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	v127 = int32(276)
	v128 = v75
	v129 = v79
	goto L36
L36:
	;
	v133 = F_palloc(m, v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v127 = v126
	v128 = v122
	v129 = v79 << (uint(int32(1)) % 32)
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128+v78<<(uint(int32(2))%32)))) = v133
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v138 = F_fread(m, v133, int32(1), v137, v49)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v138 != v137 {
		v698 = v128
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v145 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	if v145 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L42
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v150 = F_palloc(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v153 = F_fread(m, v150, int32(1), v149, v49)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v153 != v149 {
		v698 = v128
		goto L15
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+48)) = v150
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	v158 = F_CreateTemplateTupleDesc(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+52)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v164
	goto L49
L48:
	;
	v166 = int32(2249)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = int32(-1)
	v171 = int32(0)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v171 < v173 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v180 = v171
	v182 = v171
	goto L53
L51:
	;
	v236 = v171
	goto L52
L52:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v249 = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v249 < v258 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v198 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L55
	}
L54:
	;
	v236 = v227
	goto L52
L55:
	;
	if v198 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v202 != int32(100) {
		v698 = v128
		goto L15
	} else {
		goto L57
	}
L57:
	;
	v208 = int32(100)
	v212 = v192 + v193<<(uint(int32(3))%32) + v182*v208 + int32(28)
	v215 = F_fread(m, v212, int32(1), v208, v49)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v215 != int32(100) {
		v698 = v128
		goto L15
	} else {
		goto L59
	}
L59:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+86)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	F_populate_compact_attribute(m, v220, v182)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v223 = int32(1)
	v227 = base.B2i32(v219|v180&v223 != int32(0))
	v229 = v182 + v223
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v229 < v230 {
		v180 = v227
		v182 = v229
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	v340 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L81
	}
L63:
	;
	v262 = v248 + int32(28)
	v269 = v249
	v270 = v258
	v272 = v249
	goto L67
L64:
	;
	v326 = v249
	v333 = v258
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+20)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v326
	goto L62
L66:
	;
	v326 = v320
	v333 = v299
	goto L65
L67:
	;
	v278 = v262 + v258<<(uint(int32(3))%32) + v269*int32(100)
	v281 = v262 + v269<<(uint(int32(3))%32)
	if v258 != v270 {
		v299 = v270
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v320 = v258
	goto L66
L69:
	;
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v281)+2)))
	if v300 <= int32(0) {
		v320 = v269
		goto L66
	} else {
		goto L77
	}
L70:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+7)))
	if v283 != int32(118) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v299 = v269
	goto L69
L72:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+4)))
	if v286 != int32(1) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+6)))
	if v289&int32(6) != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v281)+2)))
	if v292 <= int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+90)))
	if v295 != int32(118) {
		v299 = v258
		goto L69
	} else {
		goto L76
	}
L76:
	;
	goto L71
L77:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+90)))
	if v303 == int32(118) {
		v320 = v269
		goto L66
	} else {
		goto L78
	}
L78:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+5)))
	v312 = (v272 + v306 - int32(1)) & (int32(0) - v306)
	if int32(_a_F_load_relcache_init_file_9) < v312 {
		v320 = v269
		goto L66
	} else {
		goto L79
	}
L79:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v281))) = uint16(v312)
	v318 = v269 + int32(1)
	if v318 != v258 {
		v269 = v318
		v270 = v299
		v272 = v312 + v300
		goto L67
	} else {
		goto L80
	}
L80:
	;
	goto L68
L81:
	;
	if v340 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v344 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v236 != 0 {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v345 = F_palloc(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+180)) = int32(0)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+180)) = v345
	v349 = F_fread(m, v345, int32(1), v344, v49)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	if v349 != v344 {
		v698 = v128
		goto L15
	} else {
		goto L89
	}
L89:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v133)+180))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if v344 == int32(base.Ui32(v353)>>(uint(int32(2))%32)) {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	v698 = v128
	goto L15
L91:
	;
	v361 = F_palloc0(m, int32(20))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+25)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+119)))
	if v370 == int32(105) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v361)+16)) = uint8(v363)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+24)) = v361
	goto L93
L95:
	;
	v597 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+260)) = v597
	v599 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+128)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+68)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v133)+76)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+92)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+100)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+108)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+116)) = v599
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+124)) = uint8(v597)
	*(*int64)(unsafe.Add(mBase, uint32(v133)+228)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+236)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v133)+244)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v133)+176)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v133)+164)) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v133)+156)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v133)+136)) = v597
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+27)) = uint8(v597)
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+25)))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v633
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+88)) = uint8(v597)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+28)) = uint8(v597)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+84)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v133)+256)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(v133)+272)) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v133)+32)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+40)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(v133)+144)) = v599
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+152)) = uint8(v597)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v133)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+60)) = v653
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[1]))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658)+117)))
	if v659 != 0 {
		goto L151
	} else {
		goto L152
	}
L96:
	;
	v374 = v19 + int32(72)
	v377 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v574 = v76 + v368
	switch v370 - int32(83) {
	case 0, 26, 31, 33:
		goto L148
	default:
		v589 = v574
		v590 = v77
		goto L95
	}
L99:
	;
	if v377 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L100
	}
L100:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v382 = F_palloc(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+196)) = v382
	v386 = F_fread(m, v382, int32(1), v381, v49)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	if v386 != v381 {
		v698 = v128
		goto L15
	} else {
		goto L103
	}
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v133)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+16)) = v389 + int32(24)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v133)+196))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+192)) = v394 + v395
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[2]))
	v404 = F_AllocSetContextCreateInternal(m, v399, int32(_a_F_load_relcache_init_file_10), int32(0), int32(1024), int32(_a_F_load_relcache_init_file_11))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+200)) = v404
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v410 = F_MemoryContextStrdup(m, v404, v407+int32(4))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+36)) = v410
	v413 = int32(_a_F_load_relcache_init_file_12)
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[3]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v133)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[3])) = v416
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v133)+184))
	v419 = F_GetIndexAmRoutine(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+204)) = v419
	*(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[3])) = v414
	v426 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v426 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L108
	}
L108:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v431 = F_MemoryContextAlloc(m, v404, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v434 = F_fread(m, v431, int32(1), v430, v49)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	if v434 != v430 {
		v698 = v128
		goto L15
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+208)) = v431
	v440 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	if v440 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L113
	}
L113:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v445 = F_MemoryContextAlloc(m, v404, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v448 = F_fread(m, v445, int32(1), v444, v49)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	if v448 != v444 {
		v698 = v128
		goto L15
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+212)) = v445
	v454 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v454 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L118
	}
L118:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v459 = F_MemoryContextAlloc(m, v404, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v462 = F_fread(m, v459, int32(1), v458, v49)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	if v462 != v458 {
		v698 = v128
		goto L15
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+216)) = v459
	v468 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	if v468 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L123
	}
L123:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v473 = F_MemoryContextAlloc(m, v404, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v476 = F_fread(m, v473, int32(1), v472, v49)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if v476 != v472 {
		v698 = v128
		goto L15
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+248)) = v473
	v482 = F_fread(m, v374, int32(1), int32(4), v49)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	if v482 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L128
	}
L128:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v487 = F_MemoryContextAlloc(m, v404, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	v490 = F_fread(m, v487, int32(1), v486, v49)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	if v490 != v486 {
		v698 = v128
		goto L15
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+224)) = v487
	v494 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	v497 = F_MemoryContextAllocZero(m, v404, v494<<(uint(int32(2))%32))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+252)) = v497
	v500 = int32(0)
	v501 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v500 < v501 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v508 = v500
	goto L136
L134:
	;
	v555 = v501
	goto L135
L135:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v133)+204))
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v566)+6)))
	v571 = F_MemoryContextAllocZero(m, v404, v555*v567*int32(28))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L5
	} else {
		goto L147
	}
L136:
	;
	v524 = F_fread(m, v19+int32(72), int32(1), int32(4), v49)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L5
	} else {
		goto L138
	}
L137:
	;
	v555 = v547
	goto L135
L138:
	;
	if v524 != int32(4) {
		v698 = v128
		goto L15
	} else {
		goto L139
	}
L139:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v528 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v529 = F_MemoryContextAlloc(m, v404, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v546 = v508 + int32(1)
	v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+120)))
	if v546 < v547 {
		v508 = v546
		goto L136
	} else {
		goto L146
	}
L143:
	;
	v532 = v508 << (uint(int32(2)) % 32)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v133)+252))
	*(*int32)(unsafe.Add(mBase, uint32(v532+v533))) = v529
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v133)+252))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v536+v532)))
	v540 = F_fread(m, v538, int32(1), v528, v49)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	if v540 != v528 {
		v698 = v128
		goto L15
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	goto L137
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+220)) = v571
	v589 = v76
	v590 = v77 + v368
	goto L95
L148:
	;
	F_RelationInitTableAccessMethod(m, v133)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v589 = v574
	v590 = v77
	goto L95
L150:
	;
	F_RelationInitPhysicalAddr(m, v133)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L5
	} else {
		goto L154
	}
L151:
	;
	v660 = v597
	goto L153
L152:
	;
	v660 = v657
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+64)) = v660
	goto L150
L154:
	;
	v75 = v128
	v76 = v589
	v77 = v590
	v78 = v78 + int32(1)
	v79 = v129
	goto L19
L155:
	;
	v671 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L5
	} else {
		goto L156
	}
L156:
	;
	if v671 == int32(0) {
		v698 = v75
		goto L15
	} else {
		goto L157
	}
L157:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = int64(30064771076)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v76
	F_errmsg_internal(m, int32(_a_F_load_relcache_init_file_13), v19+int32(32))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L5
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_load_relcache_init_file_6), int32(_a_F_load_relcache_init_file_14), int32(_a_F_load_relcache_init_file_8))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	v698 = v75
	goto L15
L160:
	;
	v707 = F_FreeFile(m, v49)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	goto L13
L162:
	;
	v730 = int32(0)
	goto L165
L163:
	;
	goto L164
L164:
	;
	F_pfree(m, v75)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L5
	} else {
		goto L182
	}
L165:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[4]))
	v749 = v75 + v730<<(uint(int32(2))%32)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	v756 = F_hash_search(m, v746, v750+int32(56), int32(1), v19+int32(72))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L5
	} else {
		goto L167
	}
L166:
	;
	goto L164
L167:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
	if v758 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v798 = v730 + int32(1)
	if v798 != v78 {
		v730 = v798
		goto L165
	} else {
		goto L181
	}
L169:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+4)) = v762
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v761)+16))
	if v764 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+4)) = v794
	goto L168
L172:
	;
	F_RelationDestroyRelation(m, v761, int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L5
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_load_relcache_init_file[5]))
	if v771 == int32(0) {
		goto L168
	} else {
		goto L176
	}
L175:
	;
	goto L168
L176:
	;
	v776 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	if v776 == int32(0) {
		goto L168
	} else {
		goto L178
	}
L178:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v761)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v780 + int32(4)
	F_errmsg_internal(m, int32(_a_F_load_relcache_init_file_15), v19+int32(16))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_load_relcache_init_file_6), int32(_a_F_load_relcache_init_file_16), int32(_a_F_load_relcache_init_file_8))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	goto L168
L181:
	;
	goto L166
L182:
	;
	v818 = F_FreeFile(m, v49)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	if l0 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v821 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_load_relcache_init_file[6])) = uint8(v821)
	v844 = v821
	goto L8
L185:
	;
	goto L186
L186:
	;
	v825 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_load_relcache_init_file[7])) = uint8(v825)
	v844 = v825
	goto L8
}
func F_locate_var_of_level_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v3 = int32(0)
	if l0 == v3 {
		v41 = v3
		return v41
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(58) {
		case 0:
			v41 = v3
			return v41
		case 1, 2, 3, 4, 5, 6, 7, 8:
			v37 = F_expression_tree_walker_impl(m, l0, int32(953), l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v41 = v37
				return v41
			}
		case 9:
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v21 + int32(1)
			v27 = F_query_tree_walker_impl(m, l0, int32(953), l1, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v31 - int32(1)
				return v27
			}
		default:
			if v7 != int32(6) {
				v37 = F_expression_tree_walker_impl(m, l0, int32(953), l1)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v41 = v37
					return v41
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v12 != v13 {
					v41 = v3
					return v41
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v15 < int32(0) {
						v41 = v3
						return v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
						return int32(1)
					}
				}
			}
		}
	}
}
func F_locate_windowfunc_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(11) {
			v19 = F_expression_tree_walker_impl(m, l0, int32(1124), l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v11 < int32(0) {
				v19 = F_expression_tree_walker_impl(m, l0, int32(1124), l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
				return int32(1)
			}
		}
	}
}
func F_lookup_rowtype_tupdesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = F_lookup_rowtype_tupdesc_internal(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		if int32(0) <= v7 {
			F_IncrTupleDescRefCount(m, v3)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v3
			}
		} else {
			return v3
		}
	}
}
func F_ltsGetBlock(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v86 int64
	_ = v86
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v131 int64
	_ = v131
	var v152 int64
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v196 int64
	_ = v196
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v213 int64
	_ = v213
	var v218 int64
	_ = v218
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v252 int64
	_ = v252
	var v257 int64
	_ = v257
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v12 == int32(1) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		if int32(0) < v16 {
			v178 = v16
			v179 = v15
			v182 = v178 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v182
			v187 = *(*int64)(unsafe.Add(mBase, uint32(v179+v182<<(uint(int32(3))%32))))
			return v187
		} else {
			if v15 == int32(0) {
				v21 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v21
				v25 = F_palloc_mul(m, v21, v21)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v45 = v25
					*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v45
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v47
					if v47 <= int32(0) {
						v178 = v47
						v179 = v45
					} else {
						v51 = v47
						v61 = v51
						for {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							if base.Ui64(v65) <= base.Ui64(int64(1)) {
								if base.I32_wrap_i64(v65) != int32(1) {
									v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + int64(1)
									v152 = v71
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
									v77 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
									v152 = v77
								}
							} else {
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
								v80 = v65 - int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v80
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v80)<<(uint(int32(3))%32))))
								v92 = int64(0)
								for {
									v99 = int64(1)
									v100 = v92 << (uint(v99) % 64)
									v102 = v100 + int64(2)
									v104 = v100 | v99
									if base.B2i32(base.Ui64(v80) <= base.Ui64(v104))|base.B2i32(base.Ui64(v80) <= base.Ui64(v102)) == int32(0) {
										v111 = int32(3)
										v114 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v104)<<(uint(v111)%32))))
										v119 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v102)<<(uint(v111)%32))))
										if v114 < v119 {
											v121 = v104
										} else {
											v121 = v102
										}
										v125 = v121
										v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
										if v86 <= v131 {
											break
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
											v92 = v125
											continue
										}
										break
									} else {
										v122 = base.B2i32(base.Ui64(v104) < base.Ui64(v80))
										if base.Ui64(v104) < base.Ui64(v80) {
											v123 = v104
										} else {
											v123 = v102
										}
										if base.Ui64(v104) < base.Ui64(v80) {
											v125 = v123
											v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
											if v86 <= v131 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
												v92 = v125
												continue
											}
											break
										} else {
											if base.Ui64(v80) <= base.Ui64(v102) {
												break
											} else {
												v125 = v123
												v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
												if v86 <= v131 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
													v92 = v125
													continue
												}
												break
											}
											break
										}
										break
									}
									break
								}
								*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v86
								v152 = v78
							}
							v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
							*(*int64)(unsafe.Add(mBase, uint32(v157+v61<<(uint(int32(3))%32)-int32(8)))) = v152
							v164 = int32(1)
							if v164 < v61 {
								v61 = v61 - v164
								continue
							} else {
								break
							}
							break
						}
						v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
						v178 = v168
						v179 = v169
					}
					v182 = v178 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v182
					v187 = *(*int64)(unsafe.Add(mBase, uint32(v179+v182<<(uint(int32(3))%32))))
					return v187
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if int32(128) <= v29 {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v29
					v51 = v29
					v61 = v51
					for {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						if base.Ui64(v65) <= base.Ui64(int64(1)) {
							if base.I32_wrap_i64(v65) != int32(1) {
								v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + int64(1)
								v152 = v71
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
								v152 = v77
							}
						} else {
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
							v80 = v65 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v80
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v80)<<(uint(int32(3))%32))))
							v92 = int64(0)
							for {
								v99 = int64(1)
								v100 = v92 << (uint(v99) % 64)
								v102 = v100 + int64(2)
								v104 = v100 | v99
								if base.B2i32(base.Ui64(v80) <= base.Ui64(v104))|base.B2i32(base.Ui64(v80) <= base.Ui64(v102)) == int32(0) {
									v111 = int32(3)
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v104)<<(uint(v111)%32))))
									v119 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v102)<<(uint(v111)%32))))
									if v114 < v119 {
										v121 = v104
									} else {
										v121 = v102
									}
									v125 = v121
									v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
									if v86 <= v131 {
										break
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
										v92 = v125
										continue
									}
									break
								} else {
									v122 = base.B2i32(base.Ui64(v104) < base.Ui64(v80))
									if base.Ui64(v104) < base.Ui64(v80) {
										v123 = v104
									} else {
										v123 = v102
									}
									if base.Ui64(v104) < base.Ui64(v80) {
										v125 = v123
										v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
										if v86 <= v131 {
											break
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
											v92 = v125
											continue
										}
										break
									} else {
										if base.Ui64(v80) <= base.Ui64(v102) {
											break
										} else {
											v125 = v123
											v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
											if v86 <= v131 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
												v92 = v125
												continue
											}
											break
										}
										break
									}
									break
								}
								break
							}
							*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v86
							v152 = v78
						}
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
						*(*int64)(unsafe.Add(mBase, uint32(v157+v61<<(uint(int32(3))%32)-int32(8)))) = v152
						v164 = int32(1)
						if v164 < v61 {
							v61 = v61 - v164
							continue
						} else {
							break
						}
						break
					}
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
					v178 = v168
					v179 = v169
					v182 = v178 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v182
					v187 = *(*int64)(unsafe.Add(mBase, uint32(v179+v182<<(uint(int32(3))%32))))
					return v187
				} else {
					v33 = int32(128)
					v35 = v29 << (uint(int32(1)) % 32)
					if v33 <= v35 {
						v38 = v33
					} else {
						v38 = v35
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v38
					v42 = F_repalloc(m, v15, v38<<(uint(int32(3))%32))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						v45 = v42
						*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v45
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v47
						if v47 <= int32(0) {
							v178 = v47
							v179 = v45
						} else {
							v51 = v47
							v61 = v51
							for {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								if base.Ui64(v65) <= base.Ui64(int64(1)) {
									if base.I32_wrap_i64(v65) != int32(1) {
										v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + int64(1)
										v152 = v71
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
										v77 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
										v152 = v77
									}
								} else {
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
									v80 = v65 - int64(1)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v80
									v86 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v80)<<(uint(int32(3))%32))))
									v92 = int64(0)
									for {
										v99 = int64(1)
										v100 = v92 << (uint(v99) % 64)
										v102 = v100 + int64(2)
										v104 = v100 | v99
										if base.B2i32(base.Ui64(v80) <= base.Ui64(v104))|base.B2i32(base.Ui64(v80) <= base.Ui64(v102)) == int32(0) {
											v111 = int32(3)
											v114 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v104)<<(uint(v111)%32))))
											v119 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v102)<<(uint(v111)%32))))
											if v114 < v119 {
												v121 = v104
											} else {
												v121 = v102
											}
											v125 = v121
											v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
											if v86 <= v131 {
												break
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
												v92 = v125
												continue
											}
											break
										} else {
											v122 = base.B2i32(base.Ui64(v104) < base.Ui64(v80))
											if base.Ui64(v104) < base.Ui64(v80) {
												v123 = v104
											} else {
												v123 = v102
											}
											if base.Ui64(v104) < base.Ui64(v80) {
												v125 = v123
												v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
												if v86 <= v131 {
													break
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
													v92 = v125
													continue
												}
												break
											} else {
												if base.Ui64(v80) <= base.Ui64(v102) {
													break
												} else {
													v125 = v123
													v131 = *(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v125)<<(uint(int32(3))%32))))
													if v86 <= v131 {
														break
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v131
														v92 = v125
														continue
													}
													break
												}
												break
											}
											break
										}
										break
									}
									*(*int64)(unsafe.Add(mBase, uint32(v64+base.I32_wrap_i64(v92)<<(uint(int32(3))%32)))) = v86
									v152 = v78
								}
								v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
								*(*int64)(unsafe.Add(mBase, uint32(v157+v61<<(uint(int32(3))%32)-int32(8)))) = v152
								v164 = int32(1)
								if v164 < v61 {
									v61 = v61 - v164
									continue
								} else {
									break
								}
								break
							}
							v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
							v178 = v168
							v179 = v169
						}
						v182 = v178 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v182
						v187 = *(*int64)(unsafe.Add(mBase, uint32(v179+v182<<(uint(int32(3))%32))))
						return v187
					}
				}
			}
		}
	} else {
		v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v190 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
		if base.Ui64(v190) <= base.Ui64(int64(1)) {
			if base.I32_wrap_i64(v190) != int32(1) {
				v196 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v196 + int64(1)
				return v196
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = int64(0)
				v203 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
				return v203
			}
		} else {
			v205 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
			v207 = v190 - int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v207
			v213 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v207)<<(uint(int32(3))%32))))
			v218 = int64(0)
			for {
				v225 = int64(1)
				v226 = v218 << (uint(v225) % 64)
				v228 = v226 + int64(2)
				v230 = v226 | v225
				if base.B2i32(base.Ui64(v207) <= base.Ui64(v230))|base.B2i32(base.Ui64(v207) <= base.Ui64(v228)) == int32(0) {
					v237 = int32(3)
					v240 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v230)<<(uint(v237)%32))))
					v245 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v228)<<(uint(v237)%32))))
					if v240 < v245 {
						v247 = v230
					} else {
						v247 = v228
					}
					v252 = v247
					v257 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v252)<<(uint(int32(3))%32))))
					if v213 <= v257 {
						break
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v257
						v218 = v252
						continue
					}
					break
				} else {
					v248 = base.B2i32(base.Ui64(v230) < base.Ui64(v207))
					if base.Ui64(v230) < base.Ui64(v207) {
						v249 = v230
					} else {
						v249 = v228
					}
					if base.Ui64(v230) < base.Ui64(v207) {
						v252 = v249
						v257 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v252)<<(uint(int32(3))%32))))
						if v213 <= v257 {
							break
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v257
							v218 = v252
							continue
						}
						break
					} else {
						if base.Ui64(v207) <= base.Ui64(v228) {
							break
						} else {
							v252 = v249
							v257 = *(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v252)<<(uint(int32(3))%32))))
							if v213 <= v257 {
								break
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v257
								v218 = v252
								continue
							}
							break
						}
						break
					}
					break
				}
				break
			}
			*(*int64)(unsafe.Add(mBase, uint32(v189+base.I32_wrap_i64(v218)<<(uint(int32(3))%32)))) = v213
			return v205
		}
	}
}
func F_ltsReadFillBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v74 int64
	_ = v74
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v91 int64
	_ = v91
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v114 int32
	_ = v114
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = v17
	v23 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	if v19 == int64(-1) {
		v157 = v23
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return base.B2i32(int32(0) < v157)
L4:
	;
	goto L3
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v34 = v33 + v19
	v35 = F_BufFileSeekBlock(m, v32, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = v23 + v30
	F_BufFileReadExact(m, v39, v40, int32(_a_F_ltsReadFillBuffer_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_ltsReadFillBuffer[0])))
	if int64(0) <= v134 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+40)))
	if v46 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v47 < base.I64_extend_i32_u(v48) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+48)) = v64 + int64(1)
	if v64 == int64(0) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v64 = v47
	v65 = v51
	goto L13
L15:
	;
	goto L16
L16:
	;
	v53 = v48 << (uint(int32(4)) % 32)
	if base.Ui32(int32(1073741823)) < base.Ui32(v53) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+56)) = v48 << (uint(int32(1)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v60 = F_repalloc(m, v59, v53)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+44)) = v60
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v64 = v63
	v65 = v60
	goto L13
L19:
	;
	v114 = int32(0)
	goto L21
L20:
	;
	v74 = v64
	goto L22
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v65+v114<<(uint(int32(3))%32)))) = v34
	goto L10
L22:
	;
	v83 = int64(1)
	v86 = int64(base.Ui64(v74-v83) >> (uint(v83) % 64))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v65+base.I32_wrap_i64(v86)<<(uint(int32(3))%32))))
	if v91 < v34 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v114 = base.I32_wrap_i64(v102)
	goto L21
L24:
	;
	goto L23
L25:
	;
	v102 = v74
	goto L24
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v65+base.I32_wrap_i64(v74)<<(uint(int32(3))%32)))) = v91
	v98 = int64(0)
	if v86 != v98 {
		v74 = v86
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v102 = v98
	goto L24
L29:
	;
	v139 = int32(_a_F_ltsReadFillBuffer_1)
	goto L31
L30:
	;
	v139 = int32(0) - base.I32_wrap_i64(v134)
	goto L31
L31:
	;
	v140 = v131 + v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v140
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_ltsReadFillBuffer[0])))
	if v142 < int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
	v157 = v140
	goto L4
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v142
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(_a_F_ltsReadFillBuffer_0) < v148-v140 {
		v19 = v142
		v23 = v140
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v157 = v140
	goto L4
L36:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v34
	F_errmsg(m, int32(_a_F_ltsReadFillBuffer_2), v13)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ltsReadFillBuffer_3), int32(288), int32(_a_F_ltsReadFillBuffer_4))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
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
