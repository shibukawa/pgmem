package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CreateDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v25 int64
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	switch l0 - int32(1) {
	case 0:
		return int32(_a_F_CreateDestReceiver_0)
	case 1, 2:
		v7 = F_palloc0(m, int32(60))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(base.B2i32(l0 == int32(2)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(26)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(27)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(28)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(29)
			v25 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+28)) = v25
			*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v25
			return v7
		}
	case 3:
		v107 = int32(_a_F_CreateDestReceiver_1)
		return v107
	case 4:
		return int32(_a_F_CreateDestReceiver_2)
	case 5:
		v37 = F_palloc0(m, int32(56))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(830)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = int32(831)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(832)
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(833)
			return v37
		}
	case 6:
		v51 = F_CreateIntoRelDestReceiver(m, int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			return v51
		}
	case 7:
		v55 = F_palloc(m, int32(32))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = int64(8)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(566)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(567)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(568)
			*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(569)
			return v55
		}
	case 8:
		v71 = F_palloc0(m, int32(28))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = int32(9)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(734)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(735)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(736)
			*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(737)
			return v71
		}
	case 9:
		v85 = F_palloc0(m, int32(40))
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = int32(10)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = int32(604)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(605)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = int32(606)
			*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(607)
			return v85
		}
	case 10:
		v101 = F_CreateTupleQueueDestReceiver(m, int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			return v101
		}
	case 11:
		v105 = F_CreateExplainSerializeDestReceiver(m, int32(0))
		mBase = m.M
		v106 = m.ExcPending
		if v106 != 0 {
			return int32(0)
		} else {
			v107 = v105
			return v107
		}
	default:
		return int32(_a_F_CreateDestReceiver_3)
	}
}
func F_CreateRestartPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v227 int32
	_ = v227
	var v230 int64
	_ = v230
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int64
	_ = v253
	var v258 float64
	_ = v258
	var v260 int32
	_ = v260
	var v262 float64
	_ = v262
	var v269 float64
	_ = v269
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v283 int64
	_ = v283
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v352 int64
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int64
	_ = v411
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	v11 = m.G0
	v13 = v11 - int32(1200)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v19 = base.AtomicRmwXchg32(m, v16, int32(440), int32(1))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v16+int32(440), int32(_a_F_CreateRestartPoint_0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+344))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v28)+336))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)+328))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v28)+320))
	base.MemoryCopy(m, v13+int32(76), v28+int32(348), int32(84))
	v39 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v28)+440)), uint32(v39))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v43 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v13 + int32(1200)
	return v447
L7:
	;
	if v32 != int64(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+308))
	v49 = base.B2i32(v47 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(v49)
	if v47 != int32(2) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v52 = int32(0)
	v55 = F_errstart(m, int32(13), v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v55 == int32(0) {
		v447 = v52
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_15), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_10), int32(_a_F_CreateRestartPoint_16), int32(_a_F_CreateRestartPoint_12))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v447 = v52
	goto L6
L16:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L32
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+40))
	if base.Ui64(v72) < base.Ui64(v30) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v74 = int32(0)
	v77 = F_errstart(m, int32(13), v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v30)
	v81 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v81)
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_17), v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_10), int32(_a_F_CreateRestartPoint_18), int32(_a_F_CreateRestartPoint_12))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if l0&int32(1) == int32(0) {
		v447 = v74
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v104 = F_LWLockAcquire(m, v100+int32(1152), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = int32(2)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v111, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v115+int32(1152))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v447 = v74
	goto L6
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+152)) = v30
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3])) = v30
	F_WALInsertLockRelease(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v133 = base.AtomicRmwXchg32(m, v130, int32(440), int32(1))
	if v133 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_s_lock(m, v130+int32(440), int32(_a_F_CreateRestartPoint_0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v140)+200)) = v30
	v142 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v140)+440)), uint32(v142))
	v145 = int32(_a_F_CreateRestartPoint_1)
	base.MemoryFill(m, v145, v142, int32(80))
	v153 = m.G0
	v154 = int32(16)
	v155 = v153 - v154
	m.G0 = v155
	F_gettimeofday(m, v155)
	mBase = m.M
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
	v159 = int64(*(*int32)(unsafe.Add(mBase, uint32(v155)+8)))
	m.G0 = v155 + v154
	goto L38
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[4])) = v159 + v158*int64(1000000) - int64(946684800000000)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v170 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_LogCheckpointStart(m, l0, int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if l0&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = int32(_a_F_CreateRestartPoint_2)
	if l0&int32(2) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	F_CheckPointGuts(m, v30, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L53
	}
L46:
	;
	v184 = int32(_a_F_CreateRestartPoint_3)
	goto L48
L47:
	;
	v184 = int32(_a_F_CreateRestartPoint_4)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v184
	if l0&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v190 = int32(_a_F_CreateRestartPoint_5)
	goto L51
L50:
	;
	v190 = int32(_a_F_CreateRestartPoint_4)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v190
	v193 = v13 + int32(160)
	v198 = F_pg_snprintf(m, v193, int32(128), int32(_a_F_CreateRestartPoint_6), v13+int32(48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v200 = F_strlen(m, v193)
	mBase = m.M
	goto L45
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v205)+40))
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v212 = F_LWLockAcquire(m, v208+int32(1152), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v215)+40))
	if base.Ui64(v216) < base.Ui64(v30) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+48)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v215)+40)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v215)+32)) = v32
	base.MemoryCopy(m, v215+int32(52), v13+int32(76), int32(84))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v227 != int32(5) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v247+int32(1152))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L65
	}
L58:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v243, v215)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L64
	}
L59:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v215)+144))
	if base.Ui64(v230) < base.Ui64(v31) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+152)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v215)+144)) = v31
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[7])) = v31
	goto L62
L61:
	;
	goto L62
L62:
	;
	if l0&int32(1) == int32(0) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = int32(2)
	goto L58
L64:
	;
	goto L57
L65:
	;
	v253 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	if v206 != int64(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v258 = base.F64_convert_i64_u(v253 - v206)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[9])) = v258
	v260 = int32(_a_F_CreateRestartPoint_7)
	v262 = *(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10]))
	if base.F64_gt(v258, v262) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v274 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v275 = base.I64_div_u_s(v253, v274)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v275
	v277 = int32(0)
	v279 = F_GetWalRcvFlushRecPtr(m, v277, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L72
	}
L69:
	;
	v269 = v258
	goto L71
L70:
	;
	v269 = base.F64_add(base.F64_mul(v262, float64(0.9)), base.F64_mul(v258, float64(0.1)))
	goto L71
L71:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10])) = v269
	goto L68
L72:
	;
	v283 = F_GetXLogReplayRecPtr(m, v13+int32(72))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	if base.Ui64(v283) < base.Ui64(v279) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v286 = v279
	goto L76
L75:
	;
	v286 = v283
	goto L76
L76:
	;
	v288 = v13 - int32(-64)
	F_KeepLogSeg(m, v286, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	v293 = int32(0)
	v295 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v292, v293, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	if v295 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v298 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	v300 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v301 = base.I64_div_u_s(v298, v300)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v301
	F_KeepLogSeg(m, v286, v288)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	v306 = v292
	goto L81
L81:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v312 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	v306 = v305
	goto L81
L83:
	;
	v329 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	F_RemoveOldXlogFiles(m, v306-int64(1), v329, v286, v326)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L87
	}
L84:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v310)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v324
	v326 = v324
	goto L83
L85:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v310)+308))
	v317 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(base.B2i32(v316 != v317))
	if v316 == v317 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v326 = v322
	goto L83
L87:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+312)))
	if v334 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[12]))
	v380 = base.AtomicRmwXchg32(m, v377, int32(96), int32(1))
	if v380 != 0 {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v338 = v286 - int64(1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11]))
	if base.Ui64(v338&base.I64_extend_i32_s(v340-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v340), float64(0.75))))) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v352 = base.I64_div_u_s(v338, base.I64_extend_i32_s(v340))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v360 = F_XLogFileInitInternal(m, v352+int64(1), v355, v13+int32(1199), v13+int32(160))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if int32(0) <= v360 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v364 = F_close(m, v360)
	mBase = m.M
	goto L94
L93:
	;
	goto L94
L94:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1199)))
	if v365 != int32(1) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v368 = int32(_a_F_CreateRestartPoint_14)
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14])) = v370 + int32(1)
	goto L88
L96:
	;
	F_s_lock(m, v377+int32(96), int32(_a_F_CreateRestartPoint_0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[12]))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+2)))
	v389 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v387)+96)), uint32(v389))
	if v388 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v392 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	F_LogCheckpointEnd(m, int32(1), l0)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	F_TruncateSUBTRANS(m, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v399 = F_GetLatestXTime(m)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v404 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v405 = int32(15)
	goto L109
L108:
	;
	v405 = int32(13)
	goto L109
L109:
	;
	v407 = F_errstart(m, v405, int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	if v407 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v30)
	v411 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+32)) = uint32(v411)
	F_errmsg(m, int32(_a_F_CreateRestartPoint_8), v13+int32(32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v433 = int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[13]))
	if v435 == int32(0) {
		v447 = v433
		goto L6
	} else {
		goto L121
	}
L114:
	;
	if v399 != int64(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v420 = F_timestamptz_to_str(m, v399)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_10), int32(_a_F_CreateRestartPoint_11), int32(_a_F_CreateRestartPoint_12))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v420
	v426 = F_errdetail(m, int32(_a_F_CreateRestartPoint_9), v13+int32(16))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L113
L121:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v438 == int32(0) {
		v447 = v433
		goto L6
	} else {
		goto L122
	}
L122:
	;
	F_ExecuteRecoveryCommand(m, v435, int32(_a_F_CreateRestartPoint_13), int32(0), int32(134217729))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v447 = v433
	goto L6
}
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
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
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int64
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v542 int64
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v790 int32
	_ = v790
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1032 int32
	_ = v1032
	var v1049 int32
	_ = v1049
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1352 int32
	_ = v1352
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1392 int32
	_ = v1392
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1433 int32
	_ = v1433
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1579 int32
	_ = v1579
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int64
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int64
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1720 int32
	_ = v1720
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1770 int32
	_ = v1770
	var v1794 int32
	_ = v1794
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1817 int32
	_ = v1817
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1848 int32
	_ = v1848
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1892 int32
	_ = v1892
	var v1894 int64
	_ = v1894
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(496)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v26 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v31 = v30
	goto L3
L2:
	;
	v31 = v26
	goto L3
L3:
	;
	if l1 == int32(0) {
		v184 = v5
		v185 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v194 != 0 {
		goto L46
	} else {
		goto L47
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 <= int32(0) {
		v184 = v5
		v185 = v5
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v42 = v5
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L41
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v42<<(uint(int32(2))%32))))
	v64 = F_relation_open(m, v62, int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L36
	}
L10:
	;
	return
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	v69 = v67 - int32(102)
	v76 = int32(0)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v69))|base.B2i32(int32(1)<<(uint(v69)%32)&int32(_a_F_CreateStatistics_0) == v76) == v76 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	goto L9
L15:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[1])))
	if v108 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
	v85 = F_object_ownercheck(m, int32(1259), v84, v31)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v85 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88)+119)))
	switch v89 - int32(73) {
	case 0, 32:
		v99 = int32(20)
		goto L20
	default:
		goto L21
	case 10:
		goto L25
	case 29:
		goto L22
	case 36:
		goto L23
	case 45:
		goto L24
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	F_aclcheck_error(m, int32(2), v101, v102+int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L26
	}
L20:
	;
	v101 = v99
	goto L19
L21:
	;
	v99 = int32(42)
	goto L20
L22:
	;
	v101 = int32(18)
	goto L19
L23:
	;
	v101 = int32(23)
	goto L19
L24:
	;
	v101 = int32(52)
	goto L19
L25:
	;
	v101 = int32(38)
	goto L19
L26:
	;
	goto L15
L27:
	;
	v112 = int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v64)+56))
	if base.Ui32(v113) < base.Ui32(int32(_a_F_CreateStatistics_1)) {
		v122 = v112
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v124 = v42 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v124 < v125 {
		v42 = v124
		goto L8
	} else {
		goto L35
	}
L30:
	;
	if v122 != 0 {
		goto L7
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+68))
	if v117 == int32(99) {
		v122 = v112
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v120 = F_isTempToastNamespace(m, v117)
	mBase = m.M
	v122 = v120
	goto L31
L34:
	;
	goto L29
L35:
	;
	v184 = v64
	v185 = v62
	goto L4
L36:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v134 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_2), v24+int32(176))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143)+119)))
	F_errdetail_relkind_not_supported(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(136), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v159 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_5), v24+int32(192))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(152), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v520 = F_strncpy(m, v24+int32(284), v501, int32(64))
	mBase = m.M
	v521 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+63)) = uint8(v521)
	goto L110
L46:
	;
	v197 = F_QualifiedNameGetCreationNamespace(m, v194, v24+int32(348))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v184)+48))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+68))
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+368)) = uint8(v203)
	v207 = v201 + int32(4)
	if v200 == v203 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v24)+348))
	v501 = v199
	v509 = v197
	goto L45
L50:
	;
	v411 = v24 + int32(368)
	v412 = F_pstrdup(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L97
	}
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v210 <= int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v215 = int32(0)
	v219 = v203
	v220 = v210
	goto L53
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v219<<(uint(int32(2))%32))))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v240 == int32(206) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L50
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if int32(0) < v215 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v383 = v215
	v384 = v220
	goto L57
L57:
	;
	v387 = v219 + int32(1)
	if v387 < v384 {
		v215 = v383
		v219 = v387
		v220 = v384
		goto L53
	} else {
		goto L96
	}
L58:
	;
	v249 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(368)+v215))) = uint8(v249)
	v253 = v215 + int32(1)
	goto L60
L59:
	;
	v253 = v215
	goto L60
L60:
	;
	v256 = v24 + int32(368) + v253
	if v243 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v258 = v243
	goto L63
L62:
	;
	v258 = int32(_a_F_CreateStatistics_6)
	goto L63
L63:
	;
	goto L67
L64:
	;
	v378 = F_strlen(m, v256)
	mBase = m.M
	v379 = v378 + v253
	if int32(63) < v379 {
		goto L50
	} else {
		goto L95
	}
L65:
	;
	v375 = F_strlen(m, v364)
	mBase = m.M
	goto L64
L67:
	;
	goto L68
L68:
	;
	v265 = int32(63)
	if (v256^v258)&int32(3) != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v368)
	goto L65
L70:
	;
	v349 = v344
	v350 = v345
	v351 = v346
	goto L91
L71:
	;
	if v339 == int32(0) {
		v364 = v337
		v365 = v338
		goto L69
	} else {
		goto L90
	}
L72:
	;
	v337 = v258
	v338 = v256
	v339 = v265
	goto L71
L73:
	;
	goto L74
L74:
	;
	v269 = int32(0)
	if base.B2i32(v258&int32(3) == v269)|int32(0) == v269 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v305 == int32(0) {
		v364 = v302
		v365 = v303
		goto L69
	} else {
		goto L84
	}
L76:
	;
	v281 = v258
	v282 = v256
	v283 = v265
	goto L79
L77:
	;
	goto L78
L78:
	;
	v302 = v258
	v303 = v256
	v304 = v265
	v305 = int32(1)
	goto L75
L79:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v285)
	if v285 == int32(0) {
		v344 = v281
		v345 = v282
		v346 = v283
		goto L70
	} else {
		goto L81
	}
L80:
	;
	v302 = v296
	v303 = v290
	v304 = v292
	v305 = v294
	goto L75
L81:
	;
	v289 = int32(1)
	v290 = v282 + v289
	v292 = v283 - v289
	v293 = int32(0)
	v294 = base.B2i32(v292 != v293)
	v296 = v281 + v289
	if v296&int32(3) == v293 {
		v302 = v296
		v303 = v290
		v304 = v292
		v305 = v294
		goto L75
	} else {
		goto L82
	}
L82:
	;
	if v292 != 0 {
		v281 = v296
		v282 = v290
		v283 = v292
		goto L79
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if base.B2i32(v308 == int32(0))|base.B2i32(base.Ui32(v304) < base.Ui32(int32(4))) != 0 {
		v337 = v302
		v338 = v303
		v339 = v304
		goto L71
	} else {
		goto L85
	}
L85:
	;
	v315 = v302
	v316 = v303
	v317 = v304
	goto L86
L86:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v323 = int32(-2139062144)
	if (int32(16843008)-v320|v320)&v323 != v323 {
		v344 = v315
		v345 = v316
		v346 = v317
		goto L70
	} else {
		goto L88
	}
L87:
	;
	v337 = v331
	v338 = v329
	v339 = v333
	goto L71
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v320
	v328 = int32(4)
	v329 = v316 + v328
	v331 = v315 + v328
	v333 = v317 - v328
	if base.Ui32(int32(3)) < base.Ui32(v333) {
		v315 = v331
		v316 = v329
		v317 = v333
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v344 = v337
	v345 = v338
	v346 = v339
	goto L70
L91:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v353)
	if v353 == int32(0) {
		v364 = v349
		v365 = v350
		goto L69
	} else {
		goto L93
	}
L92:
	;
	v364 = v360
	v365 = v358
	goto L69
L93:
	;
	v357 = int32(1)
	v358 = v350 + v357
	v360 = v349 + v357
	v362 = v351 - v357
	if v362 != 0 {
		v349 = v360
		v350 = v358
		v351 = v362
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v383 = v379
	v384 = v382
	goto L57
L96:
	;
	goto L54
L97:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+372)) = uint8(v416)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v419
	v422 = F_makeObjectName(m, v207, v412, v411)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v425 = base.I64_extend_i32_u(v202)
	v426 = int64(0)
	v428 = F_GetSysCacheOid(m, int32(63), base.I64_extend_i32_u(v422), v425, v426, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	if v428 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v431 = int32(0)
	v435 = v422
	goto L103
L101:
	;
	v479 = v422
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+348)) = v479
	v501 = v479
	v509 = v202
	goto L45
L103:
	;
	F_pfree(m, v435)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L105
	}
L104:
	;
	v479 = v467
	goto L102
L105:
	;
	v454 = v431 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = int32(_a_F_CreateStatistics_7)
	v459 = v24 + int32(368)
	v464 = F_pg_snprintf(m, v459, int32(64), int32(_a_F_CreateStatistics_8), v24+int32(160))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	v467 = F_makeObjectName(m, v207, v412, v459)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	v470 = int64(0)
	v472 = F_GetSysCacheOid(m, int32(63), base.I64_extend_i32_u(v467), v425, v470, v470)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	if v472 != 0 {
		v431 = v454
		v435 = v467
		goto L103
	} else {
		goto L109
	}
L109:
	;
	goto L104
L110:
	;
	if l3 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v540 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+348)))
	v541 = base.I64_extend_i32_u(v509)
	v542 = int64(0)
	v544 = F_SearchSysCacheExists(m, int32(63), v540, v541, v542, v542)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L118
	}
L112:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v529 = F_object_aclcheck(m, int32(2615), v509, v527, int64(512))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	if v529 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v534 = F_get_namespace_name(m, v509)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	F_aclcheck_error(m, v529, int32(37), v534)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1892
	v1894 = *(*int64)(unsafe.Add(mBase, uint32(v1876)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1894
	m.G0 = v24 + int32(496)
	return
L118:
	;
	if v544 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+25)))
	if v546 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v589 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L122:
	;
	v551 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L10
	} else {
		goto L133
	}
L125:
	;
	if v551 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_9))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_relation_close(m, v184, int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L10
	} else {
		goto L132
	}
L129:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v24)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v556
	F_errmsg(m, int32(_a_F_CreateStatistics_10), v24)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(207), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	v1876 = int32(_a_F_CreateStatistics_11)
	goto L117
L133:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_9))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v24)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v577
	F_errmsg(m, int32(_a_F_CreateStatistics_12), v24+int32(16))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L10
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(214), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L10
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v1406 = int32(0)
	if base.B2i32(v1392 == v1406)|base.B2i32(v1097 < int32(2)) == v1406 {
		goto L339
	} else {
		goto L340
	}
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L10
	} else {
		goto L335
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L10
	} else {
		goto L331
	}
L140:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1098 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L141:
	;
	if v592 != int32(1) {
		v1097 = v592
		goto L140
	} else {
		goto L260
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L10
	} else {
		goto L255
	}
L143:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v592 < int32(9) {
		goto L151
	} else {
		goto L152
	}
L144:
	;
	if int32(2) <= v592 {
		v1097 = v592
		goto L140
	} else {
		goto L252
	}
L145:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L10
	} else {
		goto L246
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L10
	} else {
		goto L239
	}
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L10
	} else {
		goto L235
	}
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L10
	} else {
		goto L229
	}
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L10
	} else {
		goto L225
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L10
	} else {
		goto L221
	}
L151:
	;
	v595 = int32(0)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v596 <= v595 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L217
	}
L154:
	;
	v1025 = v595
	v1032 = int32(0)
	goto L144
L155:
	;
	goto L156
L156:
	;
	v600 = int32(0)
	v603 = v595
	v608 = v600
	v610 = v600
	goto L157
L157:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v623+v608<<(uint(int32(2))%32))))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	if v628 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v1025 = v849
	v1032 = v856
	goto L144
L159:
	;
	v870 = v608 + int32(1)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v870 < v871 {
		v603 = v849
		v608 = v870
		v610 = v856
		goto L157
	} else {
		goto L216
	}
L160:
	;
	v629 = F_SearchSysCacheAttName(m, v185, v628)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L10
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v627)+8))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	if v674 == int32(6) {
		goto L179
	} else {
		goto L180
	}
L163:
	;
	if v629 == int32(0) {
		goto L150
	} else {
		goto L164
	}
L164:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v629)+16))
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+22)))
	v635 = v633 + v634
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635)+74)))
	if v636 <= int32(0) {
		goto L149
	} else {
		goto L165
	}
L165:
	;
	if int32(2) <= v592 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v635)+68))
	v643 = F_lookup_type_cache(m, v641, int32(2))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L10
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v648 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635)+74)))
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635)+90)))
	if v649 == int32(118) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v643)+56))
	if v645 == int32(0) {
		goto L148
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v635)+68))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v635)+76))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v635)+96))
	v657 = F_makeVar(m, int32(1), v648, v653, v654, v655, int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L10
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v665 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(352)+v610<<(uint(v665)%32)))) = uint16(v648)
	F_ReleaseCatCache(m, v629)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L10
	} else {
		goto L177
	}
L174:
	;
	v659 = F_lappend(m, v603, v657)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	F_ReleaseCatCache(m, v629)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	v849 = v659
	v856 = v610
	goto L159
L177:
	;
	v849 = v603
	v856 = v610 + v665
	goto L159
L178:
	;
	v846 = F_lappend(m, v603, v673)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L10
	} else {
		goto L215
	}
L179:
	;
	v677 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673)+8)))
	if v677 <= int32(0) {
		goto L147
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = int32(0)
	F_pull_varattnos(m, v673, int32(1), v24+int32(368))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L10
	} else {
		goto L190
	}
L182:
	;
	if int32(2) <= v592 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v673)+12))
	v684 = F_lookup_type_cache(m, v682, int32(2))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L10
	} else {
		goto L186
	}
L184:
	;
	v690 = v677
	goto L185
L185:
	;
	v692 = F_get_attgenerated(m, v185, base.I32_extend16_s(v690))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L10
	} else {
		goto L188
	}
L186:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v684)+56))
	if v686 == int32(0) {
		goto L146
	} else {
		goto L187
	}
L187:
	;
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673)+8)))
	v690 = v689
	goto L185
L188:
	;
	if v692 == int32(118) {
		goto L178
	} else {
		goto L189
	}
L189:
	;
	v698 = int32(1)
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(352)+v610<<(uint(v698)%32)))) = uint16(v701)
	v849 = v603
	v856 = v610 + v698
	goto L159
L190:
	;
	v718 = int32(-1)
	goto L192
L191:
	;
	if v592 < int32(2) {
		goto L178
	} else {
		goto L211
	}
L192:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v24)+368))
	if v734 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L10
	} else {
		goto L207
	}
L194:
	;
	if v790 < int32(0) {
		goto L191
	} else {
		goto L205
	}
L195:
	;
	v790 = base.I32_ctz(v776) | v777<<(uint(int32(5))%32)
	goto L194
L196:
	;
	v790 = int32(-2)
	goto L194
L197:
	;
	v741 = v718 + int32(1)
	v743 = int32(base.Ui32(v741) >> (uint(int32(5)) % 32))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	if v744 <= v743 {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v747 = v734 + int32(8)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v747+v743<<(uint(int32(2))%32))))
	v754 = v751 & (int32(-1) << (uint(v741) % 32))
	if v754 != 0 {
		v776 = v754
		v777 = v743
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v756 = v743 + int32(1)
	if v756 == v744 {
		goto L196
	} else {
		goto L200
	}
L200:
	;
	v759 = v756
	goto L201
L201:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v747+v759<<(uint(int32(2))%32))))
	if v766 != 0 {
		v776 = v766
		v777 = v759
		goto L195
	} else {
		goto L203
	}
L202:
	;
	goto L196
L203:
	;
	v768 = v759 + int32(1)
	if v768 != v744 {
		v759 = v768
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	if int32(_a_F_CreateStatistics_13) < v790<<(uint(int32(16))%32)-int32(_a_F_CreateStatistics_14) {
		v718 = v790
		goto L192
	} else {
		goto L206
	}
L206:
	;
	goto L193
L207:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_15), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(364), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L10
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	v817 = F_exprType(m, v673)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	v820 = F_lookup_type_cache(m, v817, int32(2))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L10
	} else {
		goto L213
	}
L213:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v820)+56))
	if v822 == int32(0) {
		goto L145
	} else {
		goto L214
	}
L214:
	;
	goto L178
L215:
	;
	v849 = v846
	v856 = v610
	goto L159
L216:
	;
	goto L158
L217:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(8)
	F_errmsg(m, int32(_a_F_CreateStatistics_16), v24+int32(32))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(226), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L10
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L10
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v628
	F_errmsg(m, int32(_a_F_CreateStatistics_17), v24+int32(112))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(260), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_15), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(267), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v628
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v24+int32(144))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v635)+68))
	v940 = F_format_type_be(m, v939)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v940
	v946 = F_errdetail(m, int32(_a_F_CreateStatistics_19), v24+int32(128))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(282), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L10
	} else {
		goto L236
	}
L236:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_15), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L10
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(314), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L10
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	v976 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673)+8)))
	v978 = F_get_attname(m, v185, v976, int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v978
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v24-int32(-64))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L10
	} else {
		goto L242
	}
L242:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v673)+12))
	v987 = F_format_type_be(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L10
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v987
	v993 = F_errdetail(m, int32(_a_F_CreateStatistics_19), v24+int32(48))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(329), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L10
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_20), int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	v1011 = F_format_type_be(m, v817)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v1011
	v1017 = F_errdetail(m, int32(_a_F_CreateStatistics_19), v24+int32(96))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(380), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	if v1025 == int32(0) {
		goto L142
	} else {
		goto L253
	}
L253:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+4))
	if v1049 == int32(1) {
		goto L141
	} else {
		goto L254
	}
L254:
	;
	goto L142
L255:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L10
	} else {
		goto L256
	}
L256:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_21), int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	v1086 = F_errdetail(m, int32(_a_F_CreateStatistics_22), int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L10
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(396), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1096 != 0 {
		goto L139
	} else {
		goto L261
	}
L261:
	;
	v1097 = int32(1)
	goto L140
L262:
	;
	v1108 = int32(0)
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+4))
	v1112 = int32(_a_F_CreateStatistics_23)
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	if base.B2i32(v1115 == v1108)|base.B2i32(v1115 != v1118) != 0 {
		v1136 = v1115
		v1137 = v1118
		goto L270
	} else {
		goto L271
	}
L263:
	;
	v1392 = v1106
	v1401 = v5
	v1402 = v5
	v1403 = v5
	goto L137
L264:
	;
	v1106 = int32(1)
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	if int32(0) < v1103 {
		goto L262
	} else {
		goto L267
	}
L267:
	;
	v1106 = int32(1)
	goto L263
L268:
	;
	v1202 = base.B2i32(v1138 == int32(0))
	v1203 = int32(1)
	if v1103 == v1203 {
		v1392 = v1108
		v1401 = v1199
		v1402 = v1200
		v1403 = v1202
		goto L137
	} else {
		goto L295
	}
L269:
	;
	if v1138 == int32(0) {
		v1199 = v5
		v1200 = v5
		goto L268
	} else {
		goto L276
	}
L270:
	;
	v1138 = v1136 - v1137
	goto L269
L271:
	;
	v1121 = v1111
	v1122 = v1112
	goto L272
L272:
	;
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+1)))
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121)+1)))
	if v1126 == int32(0) {
		v1136 = v1126
		v1137 = v1125
		goto L270
	} else {
		goto L274
	}
L273:
	;
	v1136 = v1126
	v1137 = v1125
	goto L270
L274:
	;
	v1129 = int32(1)
	if v1126 == v1125 {
		v1121 = v1121 + v1129
		v1122 = v1122 + v1129
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v1141 = int32(_a_F_CreateStatistics_24)
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	if base.B2i32(v1144 == int32(0))|base.B2i32(v1144 != v1147) != 0 {
		v1165 = v1144
		v1166 = v1147
		goto L278
	} else {
		goto L279
	}
L277:
	;
	if v1165-v1166 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L278:
	;
	goto L277
L279:
	;
	v1150 = v1111
	v1151 = v1141
	goto L280
L280:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+1)))
	if v1155 == int32(0) {
		v1165 = v1155
		v1166 = v1154
		goto L278
	} else {
		goto L282
	}
L281:
	;
	v1165 = v1155
	v1166 = v1154
	goto L278
L282:
	;
	v1158 = int32(1)
	if v1155 == v1154 {
		v1150 = v1150 + v1158
		v1151 = v1151 + v1158
		goto L280
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v1199 = v5
	v1200 = int32(1)
	goto L268
L285:
	;
	goto L286
L286:
	;
	v1171 = int32(_a_F_CreateStatistics_25)
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	if base.B2i32(v1174 == int32(0))|base.B2i32(v1174 != v1177) != 0 {
		v1195 = v1174
		v1196 = v1177
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if v1195-v1196 != 0 {
		v1352 = v1111
		goto L138
	} else {
		goto L294
	}
L288:
	;
	goto L287
L289:
	;
	v1180 = v1111
	v1181 = v1171
	goto L290
L290:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181)+1)))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+1)))
	if v1185 == int32(0) {
		v1195 = v1185
		v1196 = v1184
		goto L288
	} else {
		goto L292
	}
L291:
	;
	v1195 = v1185
	v1196 = v1184
	goto L288
L292:
	;
	v1188 = int32(1)
	if v1185 == v1184 {
		v1180 = v1180 + v1188
		v1181 = v1181 + v1188
		goto L290
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v1199 = int32(1)
	v1200 = v5
	goto L268
L295:
	;
	v1206 = int32(0)
	if v1206 < v1103 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1209 = v1103
	goto L298
L297:
	;
	v1209 = v1206
	goto L298
L298:
	;
	v1215 = v1203
	v1226 = v1199
	v1227 = v1200
	v1228 = v1202
	goto L299
L299:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1109+v1215<<(uint(int32(2))%32))))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	v1236 = int32(_a_F_CreateStatistics_23)
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	if base.B2i32(v1239 == int32(0))|base.B2i32(v1239 != v1242) != 0 {
		v1260 = v1239
		v1261 = v1242
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1392 = v1108
	v1401 = v1324
	v1402 = v1325
	v1403 = v1326
	goto L137
L301:
	;
	v1328 = v1215 + int32(1)
	if v1209 != v1328 {
		v1215 = v1328
		v1226 = v1324
		v1227 = v1325
		v1228 = v1326
		goto L299
	} else {
		goto L330
	}
L302:
	;
	if v1260-v1261 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L303:
	;
	goto L302
L304:
	;
	v1245 = v1235
	v1246 = v1236
	goto L305
L305:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246)+1)))
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245)+1)))
	if v1250 == int32(0) {
		v1260 = v1250
		v1261 = v1249
		goto L303
	} else {
		goto L307
	}
L306:
	;
	v1260 = v1250
	v1261 = v1249
	goto L303
L307:
	;
	v1253 = int32(1)
	if v1250 == v1249 {
		v1245 = v1245 + v1253
		v1246 = v1246 + v1253
		goto L305
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	v1324 = v1226
	v1325 = v1227
	v1326 = int32(1)
	goto L301
L310:
	;
	goto L311
L311:
	;
	v1266 = int32(_a_F_CreateStatistics_24)
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	if base.B2i32(v1269 == int32(0))|base.B2i32(v1269 != v1272) != 0 {
		v1290 = v1269
		v1291 = v1272
		goto L313
	} else {
		goto L314
	}
L312:
	;
	if v1290-v1291 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L313:
	;
	goto L312
L314:
	;
	v1275 = v1235
	v1276 = v1266
	goto L315
L315:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276)+1)))
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275)+1)))
	if v1280 == int32(0) {
		v1290 = v1280
		v1291 = v1279
		goto L313
	} else {
		goto L317
	}
L316:
	;
	v1290 = v1280
	v1291 = v1279
	goto L313
L317:
	;
	v1283 = int32(1)
	if v1280 == v1279 {
		v1275 = v1275 + v1283
		v1276 = v1276 + v1283
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	v1324 = v1226
	v1325 = int32(1)
	v1326 = v1228
	goto L301
L320:
	;
	goto L321
L321:
	;
	v1296 = int32(_a_F_CreateStatistics_25)
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	if base.B2i32(v1299 == int32(0))|base.B2i32(v1299 != v1302) != 0 {
		v1320 = v1299
		v1321 = v1302
		goto L323
	} else {
		goto L324
	}
L322:
	;
	if v1320-v1321 != 0 {
		v1352 = v1235
		goto L138
	} else {
		goto L329
	}
L323:
	;
	goto L322
L324:
	;
	v1305 = v1235
	v1306 = v1296
	goto L325
L325:
	;
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306)+1)))
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305)+1)))
	if v1310 == int32(0) {
		v1320 = v1310
		v1321 = v1309
		goto L323
	} else {
		goto L327
	}
L326:
	;
	v1320 = v1310
	v1321 = v1309
	goto L323
L327:
	;
	v1313 = int32(1)
	if v1310 == v1309 {
		v1305 = v1305 + v1313
		v1306 = v1306 + v1313
		goto L325
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	v1324 = int32(1)
	v1325 = v1227
	v1326 = v1228
	goto L301
L330:
	;
	goto L300
L331:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L10
	} else {
		goto L332
	}
L332:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_26), int32(0))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L10
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(405), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L10
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L10
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v1352
	F_errmsg(m, int32(_a_F_CreateStatistics_27), v24+int32(80))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L10
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(433), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L10
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	v1413 = int32(1)
	v1416 = v1413
	v1417 = v1413
	v1418 = v1413
	goto L341
L340:
	;
	v1416 = v1401
	v1417 = v1402
	v1418 = v1403
	goto L341
L341:
	;
	F_pg_qsort(m, v24+int32(352), v1032, int32(2), int32(616))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L10
	} else {
		goto L342
	}
L342:
	;
	v1425 = int32(1)
	if v1425 < v1032 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1433 = v1425
	goto L346
L344:
	;
	goto L345
L345:
	;
	if v1025 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L346:
	;
	v1453 = v24 + int32(352) + v1433<<(uint(int32(1))%32)
	v1454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1453))))
	v1457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1453-int32(2)))))
	if v1454 == v1457 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	goto L345
L348:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L10
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	v1476 = v1433 + int32(1)
	if v1476 != v1032 {
		v1433 = v1476
		goto L346
	} else {
		goto L355
	}
L351:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L10
	} else {
		goto L352
	}
L352:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_28), int32(0))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L10
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(470), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L10
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	goto L347
L356:
	;
	v1633 = F_buildint2vector(m, v24+int32(352), v1032)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L10
	} else {
		goto L375
	}
L357:
	;
	v1501 = int32(0)
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+4))
	if v1502 <= v1501 {
		goto L356
	} else {
		goto L358
	}
L358:
	;
	v1512 = v1501
	v1514 = v1502
	goto L359
L359:
	;
	v1526 = int32(0)
	if v1526 < v1514 {
		goto L362
	} else {
		goto L363
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L10
	} else {
		goto L371
	}
L361:
	;
	goto L360
L362:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1529+v1512<<(uint(int32(2))%32))))
	v1540 = v1526
	v1541 = int32(0)
	goto L365
L363:
	;
	v1579 = v1514
	goto L364
L364:
	;
	v1592 = v1512 + int32(1)
	if v1592 < v1579 {
		v1512 = v1592
		v1514 = v1579
		goto L359
	} else {
		goto L370
	}
L365:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1556+v1540<<(uint(int32(2))%32))))
	v1561 = F_equal(m, v1533, v1560)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L10
	} else {
		goto L367
	}
L366:
	;
	if int32(2) <= v1563 {
		goto L361
	} else {
		goto L369
	}
L367:
	;
	v1563 = v1561 + v1541
	v1565 = v1540 + int32(1)
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+4))
	if v1565 < v1566 {
		v1540 = v1565
		v1541 = v1563
		goto L365
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v1579 = v1566
	goto L364
L370:
	;
	goto L356
L371:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L10
	} else {
		goto L372
	}
L372:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_29), int32(0))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L10
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_3), int32(505), int32(_a_F_CreateStatistics_4))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L10
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	if v1418 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	if v1417 != 0 {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v1647 = v24 + int32(208)
	v1648 = int32(0)
	goto L376
L378:
	;
	goto L379
L379:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+208)) = int64(100)
	v1647 = v24 + int32(208) | int32(8)
	v1648 = int32(1)
	goto L376
L380:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1647))) = int64(102)
	v1653 = v1648 + int32(1)
	goto L382
L381:
	;
	v1653 = v1648
	goto L382
L382:
	;
	if v1416 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(208)+v1653<<(uint(int32(3))%32)))) = int64(109)
	v1663 = v1653 + int32(1)
	goto L385
L384:
	;
	v1663 = v1653
	goto L385
L385:
	;
	if v1025 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	v1695 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L10
	} else {
		goto L395
	}
L387:
	;
	v1669 = F_construct_array_builtin(m, v24+int32(208), v1663, int32(18))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L10
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v1672 = v24 + int32(208)
	*(*int64)(unsafe.Add(mBase, uint32(v1672+v1663<<(uint(int32(3))%32)))) = int64(101)
	v1681 = F_construct_array_builtin(m, v1672, v1663+int32(1), int32(18))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L10
	} else {
		goto L391
	}
L390:
	;
	v1691 = v1669
	v1692 = int64(0)
	goto L386
L391:
	;
	v1683 = F_nodeToString(m, v1025)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L10
	} else {
		goto L392
	}
L392:
	;
	v1685 = F_cstring_to_text(m, v1683)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L10
	} else {
		goto L393
	}
L393:
	;
	F_pfree(m, v1683)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L10
	} else {
		goto L394
	}
L394:
	;
	v1691 = v1681
	v1692 = base.I64_extend_i32_u(v1685)
	goto L386
L395:
	;
	v1697 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+424)) = v1697
	*(*int64)(unsafe.Add(mBase, uint32(v24)+416)) = v1697
	*(*int64)(unsafe.Add(mBase, uint32(v24)+272)) = v1697
	v1703 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+280)) = uint8(v1703)
	v1707 = F_GetNewOidWithIndex(m, v1695, int32(3380), int32(1))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L10
	} else {
		goto L396
	}
L396:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+408)) = base.I64_extend_i32_u(v1633)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+400)) = base.I64_extend_i32_u(v31)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+392)) = v541
	*(*int64)(unsafe.Add(mBase, uint32(v24)+384)) = base.I64_extend_i32_u(v24 + int32(284))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+376)) = base.I64_extend_i32_u(v185)
	v1720 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+278)) = uint8(v1720)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+424)) = base.I64_extend_i32_u(v1691)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+368)) = base.I64_extend_i32_u(v1707)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+432)) = v1692
	if v1692 == int64(0) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1729 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+280)) = uint8(v1729)
	goto L399
L398:
	;
	goto L399
L399:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1695)+52))
	v1736 = F_heap_form_tuple(m, v1731, v24+int32(368), v24+int32(272))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L10
	} else {
		goto L400
	}
L400:
	;
	F_CatalogTupleInsert(m, v1695, v1736)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L10
	} else {
		goto L401
	}
L401:
	;
	F_pfree(m, v1736)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L10
	} else {
		goto L402
	}
L402:
	;
	F_relation_close(m, v1695, int32(3))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L10
	} else {
		goto L403
	}
L403:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[7]))
	if v1746 != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1748 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3381), v1707, v1748, v1748)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L10
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	F_CacheInvalidateRelcache(m, v184)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L10
	} else {
		goto L408
	}
L407:
	;
	goto L406
L408:
	;
	v1754 = int32(0)
	F_relation_close(m, v184, v1754)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L10
	} else {
		goto L409
	}
L409:
	;
	v1758 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = v1758
	*(*int32)(unsafe.Add(mBase, uint32(v24)+252)) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v24)+248)) = int32(3381)
	if v1758 < v1032 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	if v1025 != 0 {
		goto L420
	} else {
		goto L421
	}
L411:
	;
	v1770 = v1754
	goto L414
L412:
	;
	goto L413
L413:
	;
	if v1032 != 0 {
		goto L410
	} else {
		goto L418
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+264)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = int32(1259)
	v1794 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(352)+v1770<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+268)) = v1794
	F_recordDependencyOn(m, v24+int32(248), v24+int32(260), int32(97))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L10
	} else {
		goto L416
	}
L416:
	;
	v1804 = v1770 + int32(1)
	if v1804 != v1032 {
		v1770 = v1804
		goto L414
	} else {
		goto L417
	}
L417:
	;
	goto L410
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+264)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = int32(1259)
	F_recordDependencyOn(m, v24+int32(248), v24+int32(260), int32(97))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L10
	} else {
		goto L419
	}
L419:
	;
	goto L410
L420:
	;
	if l3 != 0 {
		goto L423
	} else {
		goto L424
	}
L421:
	;
	goto L422
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+264)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = int32(2615)
	v1855 = v24 + int32(248)
	F_recordDependencyOn(m, v1855, v24+int32(260), int32(110))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L10
	} else {
		goto L428
	}
L423:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	F_CheckUsageOnTypesInSingleRelExpr(m, v1025, v185, v1840)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L10
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	F_recordDependencyOnSingleRelExpr(m, v24+int32(248), v1025, v185, int32(97), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L10
	} else {
		goto L427
	}
L426:
	;
	goto L425
L427:
	;
	goto L422
L428:
	;
	F_recordDependencyOnOwner(m, int32(3381), v1707, v31)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L10
	} else {
		goto L429
	}
L429:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v1864 == int32(0) {
		v1876 = v1855
		goto L117
	} else {
		goto L430
	}
L430:
	;
	F_CreateComments(m, v1707, int32(3381), int32(0), v1864)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L10
	} else {
		goto L431
	}
L431:
	;
	v1876 = v1855
	goto L117
}
func F_create_empty_pathtarget(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc0(m, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(280)
		return v3
	}
}
func F_create_final_distinct_paths(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 float64
	_ = v114
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v43 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v33 = F_get_sortgrouplist_exprs(m, v31, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v42 = v30
	goto L1
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	if v25 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+36)))
	if v26 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+334)))
	if v27 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	return int32(0)
L9:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v38 = int32(0)
	v40 = F_estimate_num_groups(m, l0, v33, v37, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v42 = v40
	goto L1
L11:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v346 != 0 {
		goto L114
	} else {
		goto L115
	}
L12:
	;
	if v88 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L13:
	;
	v88 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v52 <= int32(0) {
		v80 = int32(1)
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v88 = v80
	goto L12
L17:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v58 = v52
	goto L20
L19:
	;
	v58 = v55
	goto L20
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v61 = int32(0)
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v71 = int32(0)
	v72 = base.B2i32(v70 != v71)
	if v70 == v71 {
		v80 = v72
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v80 = v72
	goto L16
L23:
	;
	v76 = v61 + int32(1)
	if v76 != v58 {
		v61 = v76
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)))
	if v92 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v106 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L27:
	;
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v105 = v91
	goto L26
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v96 = v95
	goto L32
L31:
	;
	v96 = v4
	goto L32
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v97 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v100 = v98
	goto L35
L34:
	;
	v100 = int32(0)
	goto L35
L35:
	;
	if v96 < v100 {
		v105 = v97
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 <= int32(0) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	if v91 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = float64(-1)
	goto L41
L40:
	;
	v114 = float64(1)
	goto L41
L41:
	;
	v128 = v4
	goto L42
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+64))
	v138 = F_get_useful_pathkeys_for_distinct(m, l0, v105, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L11
L44:
	;
	v326 = v128 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v326 < v327 {
		v128 = v326
		goto L42
	} else {
		goto L112
	}
L45:
	;
	if v138 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v143 <= v142 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v147 = v142
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v147<<(uint(int32(2))%32))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v136)+64))
	v170 = v20 + int32(12)
	if v167 == v168 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L44
L50:
	;
	v305 = v147 + int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v305 < v306 {
		v147 = v305
		goto L48
	} else {
		goto L111
	}
L51:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v276 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L52:
	;
	if v248 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v236
	v248 = int32(1)
	goto L52
L54:
	;
	if v167 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v167 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(0)
	v248 = int32(1)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(0)
	v248 = int32(1)
	goto L52
L59:
	;
	goto L60
L60:
	;
	if v168 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v188
	v248 = v188
	goto L52
L62:
	;
	goto L63
L63:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v192 = int32(0)
	if v192 < v191 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v195 = v191
	goto L66
L65:
	;
	v195 = v192
	goto L66
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v201 = int32(0)
	goto L67
L67:
	;
	if v201 < v196 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v212 = v208 + v201<<(uint(int32(2))%32)
	goto L71
L70:
	;
	v212 = int32(0)
	goto L71
L71:
	;
	if v201 == v195 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v195
	v248 = base.B2i32(v212 == int32(0))
	goto L52
L73:
	;
	goto L74
L74:
	;
	v218 = base.B2i32(v212 == int32(0))
	if v212 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = v218
	goto L52
L76:
	;
	goto L77
L77:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	if v222 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = v218
	goto L52
L79:
	;
	goto L80
L80:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v222+v201<<(uint(int32(2))%32))))
	if v226 != v230 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v201
	v248 = int32(0)
	goto L52
L82:
	;
	v201 = v201 + int32(1)
	goto L67
L84:
	;
	v273 = v136
	goto L51
L85:
	;
	goto L86
L86:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[0])))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v136 == v22 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v260&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v260 = v250
	goto L87
L89:
	;
	goto L90
L90:
	;
	if v251 == int32(0) {
		goto L50
	} else {
		goto L91
	}
L91:
	;
	v255 = int32(1)
	if v250&v255 == int32(0) {
		goto L50
	} else {
		goto L92
	}
L92:
	;
	v260 = v255
	goto L87
L93:
	;
	v264 = v251
	goto L95
L94:
	;
	v264 = int32(0)
	goto L95
L95:
	;
	if v264 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v267 = F_create_sort_path(m, l2, v136, v167, v114)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v269 = F_create_incremental_sort_path(m, l0, l2, v136, v167, v251, v114)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L101
	}
L99:
	;
	if v267 != 0 {
		v273 = v267
		goto L51
	} else {
		goto L100
	}
L100:
	;
	goto L50
L101:
	;
	if v269 == int32(0) {
		goto L50
	} else {
		goto L102
	}
L102:
	;
	v273 = v269
	goto L51
L103:
	;
	v279 = int32(0)
	v287 = F_makeConst(m, int32(20), int32(-1), v279, int32(8), int64(1), v279, int32(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v297 = F_create_unique_path(m, l2, v273, v296, v42)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L8
	} else {
		goto L109
	}
L106:
	;
	v292 = F_create_limit_path(m, l2, v273, v279, v287, int32(0), int64(0), int64(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	F_add_path(m, l2, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	goto L50
L109:
	;
	F_add_path(m, l2, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	goto L50
L111:
	;
	goto L49
L112:
	;
	goto L43
L113:
	;
	m.G0 = v20 + int32(16)
	return l2
L114:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)))
	if v347 != 0 {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v355 = int32(0)
	if v354 == v355 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[1])))
	if v349&int32(1) == int32(0) {
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	if v392 == int32(0) {
		goto L113
	} else {
		goto L132
	}
L120:
	;
	v392 = int32(1)
	goto L119
L121:
	;
	goto L122
L122:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v362 <= int32(0) {
		v386 = int32(1)
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v392 = v386
	goto L119
L124:
	;
	v365 = int32(0)
	if v365 < v362 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v368 = v362
	goto L127
L126:
	;
	v368 = v365
	goto L127
L127:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	v373 = v355
	goto L128
L128:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v369+v373<<(uint(int32(2))%32))))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+18)))
	if v378 != int32(1) {
		v386 = v378
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v386 = v378
	goto L123
L130:
	;
	v382 = v373 + int32(1)
	if v382 != v368 {
		v373 = v382
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v401 = F_create_agg_path(m, l0, l2, v22, v395, int32(2), v397, v398, v397, v397, v42)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	F_add_path(m, l2, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	goto L113
}
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == v5 {
		v70 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v72 = F_palloc0(m, int32(88))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L17
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v16 <= int32(0) {
		v70 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v26 = int32(1)
	v28 = v5
	v29 = v5
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v28<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v70 = v55
	goto L1
L6:
	;
	v36 = F_replace_nestloop_params_mutator(m, v34, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v40 = v34
	goto L8
L8:
	;
	v42 = int32(0)
	v44 = F_makeTargetEntry(m, v40, base.I32_extend16_s(v26), v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	v40 = v36
	goto L8
L11:
	;
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19+v26<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v51
	goto L14
L13:
	;
	goto L14
L14:
	;
	v55 = F_lappend(m, v29, v44)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v58 = v28 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v58 < v59 {
		v26 = v26 + int32(1)
		v28 = v58
		v29 = v55
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L5
L17:
	;
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+80)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v72)+76)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v72)+72)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v72)+56)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v72)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v72)+48)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v72)+44)) = v70
	v85 = int32(335)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v87 != v85 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v98
	v100 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v72)+8)) = v100
	v102 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v72)+16)) = v102
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v72)+24)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+36)) = uint8(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = v106
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+37)) = uint8(v110)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+37)) = uint8(v112)
	return v72
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v90 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	if v91 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+52)) = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+80)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+72)) = v96
	goto L18
}
func F_create_groupingsets_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
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
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 float64
	_ = v167
	var v170 int32
	_ = v170
	var v171 float64
	_ = v171
	var v177 float64
	_ = v177
	var v184 float64
	_ = v184
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 float64
	_ = v199
	var v200 float64
	_ = v200
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v21 = F_palloc0(m, int32(96))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(312)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(369)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v32
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v36 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v41 = v39
	goto L5
L4:
	;
	v41 = int32(0)
	goto L5
L5:
	;
	v42 = int32(1)
	v43 = v41 & v42
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+21)) = uint8(v43)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v45
	switch l4 - v42 {
	case 0:
		goto L9
	default:
		v71 = l4
		v72 = v8
		goto L6
	case 2:
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v72
	if l6 != 0 {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v71 = v53
	v72 = v70
	goto L6
L8:
	;
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if l5 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v71 = int32(1)
	v72 = v8
	goto L6
L11:
	;
	goto L12
L12:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v54 != v53 {
		v71 = v53
		v72 = v8
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v59 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v71 = int32(0)
	v72 = v8
	goto L6
L15:
	;
	v71 = int32(3)
	v72 = v8
	goto L6
L16:
	;
	goto L17
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v66 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(2)
	goto L20
L19:
	;
	v69 = int32(3)
	goto L20
L20:
	;
	v71 = v69
	v72 = v8
	goto L6
L21:
	;
	v77 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l6)+32)))
	v79 = v77
	goto L23
L22:
	;
	v79 = int64(0)
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v79
	if l5 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = base.F64_add(v229, v230)
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	v234 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v233, base.F64_add(base.F64_mul(v234, v235), v237))
	m.G0 = v18 + int32(144)
	return v21
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v84 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = int32(1)
	v99 = int32(1)
	v101 = v8
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v101<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v113 = v112
	goto L31
L30:
	;
	v113 = int32(0)
	goto L31
L31:
	;
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v211 = v101 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v211 < v212 {
		v89 = v207
		v99 = int32(0)
		v101 = v211
		goto L27
	} else {
		goto L47
	}
L33:
	;
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	F_cost_agg(m, v21, l0, v71, l6, v113, v114, l3, v115, v116, v117, v118, base.F64_convert_i32_s(v120))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	if (v126|v89)&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v207 = v124 & v89
	goto L32
L37:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v195 + v196
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
	v200 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v199, v200)
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = base.F64_add(v203, v204)
	v207 = v194
	goto L32
L38:
	;
	v131 = int32(1)
	if v126&v131 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v147 = int32(0)
	v149 = v18 + int32(72)
	v151 = float64(0)
	v152 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_create_groupingsets_path[0]))
	v160 = m.G0
	v161 = int32(16)
	v162 = v160 - v161
	m.G0 = v162
	F_cost_tuplesort(m, v162+int32(8), v162, v152, v154, v151, v157, float64(-1))
	mBase = m.M
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v162)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v149)+32)) = v152
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_groupingsets_path[1])))
	v171 = base.F64_add(v151, v167)
	*(*float64)(unsafe.Add(mBase, uint32(v149)+48)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v149)+40)) = v147 + (v170 ^ int32(1))
	v177 = *(*float64)(unsafe.Add(mBase, uint32(v162)))
	*(*float64)(unsafe.Add(mBase, uint32(v149)+56)) = base.F64_add(v171, v177)
	m.G0 = v162 + v161
	goto L45
L41:
	;
	v134 = int32(2)
	goto L43
L42:
	;
	v134 = v131
	goto L43
L43:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v137 = float64(0)
	v139 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+32))
	F_cost_agg(m, v18, l0, v134, l6, v113, v135, l3, int32(0), v137, v137, v139, base.F64_convert_i32_s(v141))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v194 = v145 & v89
	goto L37
L45:
	;
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v18)+120))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v18)+128))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v18)+104))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+32))
	F_cost_agg(m, v18, l0, int32(1), l6, v113, v184, l3, v185, v186, v187, v188, base.F64_convert_i32_s(v190))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v194 = v147
	goto L37
L47:
	;
	goto L28
}
func F_create_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v251 int32
	_ = v251
	var v254 float64
	_ = v254
	var v256 int32
	_ = v256
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v268 float64
	_ = v268
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v280 float64
	_ = v280
	var v284 float64
	_ = v284
	var v291 float64
	_ = v291
	var v295 float64
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 float64
	_ = v304
	var v306 float64
	_ = v306
	var v309 float64
	_ = v309
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
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
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v462 float64
	_ = v462
	var v466 float64
	_ = v466
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v500 float64
	_ = v500
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v509 float64
	_ = v509
	var v535 float64
	_ = v535
	var v538 float64
	_ = v538
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v548 float64
	_ = v548
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 float64
	_ = v595
	var v596 float64
	_ = v596
	var v613 float64
	_ = v613
	var v622 float64
	_ = v622
	var v624 float64
	_ = v624
	var v625 int32
	_ = v625
	var v626 float64
	_ = v626
	var v628 float64
	_ = v628
	var v629 float64
	_ = v629
	var v631 float64
	_ = v631
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = l7
	v32 = F_palloc0(m, int32(96))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(300)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v38 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v41 = v39
	goto L5
L4:
	;
	v41 = int32(0)
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+252))
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v45 = v43
	goto L8
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v45 = v44
	goto L8
L8:
	;
	v46 = int32(0)
	if base.B2i32(v41 == v46)|base.B2i32(v45 == v46) != 0 {
		v91 = v46
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v91 != 0 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	goto L9
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v56 < v57 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v56
	goto L14
L13:
	;
	v59 = v57
	goto L14
L14:
	;
	if v59 <= int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = int32(1)
	goto L17
L16:
	;
	v62 = v59
	goto L17
L17:
	;
	v63 = int32(8)
	v68 = int32(0)
	goto L18
L18:
	;
	v75 = v68 << (uint(int32(2)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45+v63+v75)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41+v63+v75)))
	v80 = v77 & v79
	v82 = base.B2i32(v80 != int32(0))
	if v80 != 0 {
		v91 = v82
		goto L10
	} else {
		goto L20
	}
L19:
	;
	v91 = v82
	goto L10
L20:
	;
	v84 = v68 + int32(1)
	if v84 != v62 {
		v68 = v84
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v92 = F_get_param_path_clause_serials(m, l6)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(360)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v198
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v204 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v201, l9, v28+int32(12))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L39
	}
L25:
	;
	if l7 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v156
	goto L24
L27:
	;
	v156 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v98 <= v97 {
		v156 = v97
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v113 = int32(0)
	v114 = v97
	goto L31
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v113<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+56))
	v133 = F_bms_is_member(m, v132, v92)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v156 = v139
	goto L26
L33:
	;
	if v133 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v137 = F_lappend(m, v114, v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v139 = v114
	goto L36
L36:
	;
	v141 = v113 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v141 < v142 {
		v113 = v141
		v114 = v139
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v139 = v137
	goto L36
L38:
	;
	goto L32
L39:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+20)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v204
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v209 != int32(1) {
		v216 = v200
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v218 = v216 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)) = uint8(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v220
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v224)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v228
	v230 = m.G0
	v232 = v230 - int32(32)
	m.G0 = v232
	v234 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v242 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v212 != int32(1) {
		v216 = v200
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v216 = v215
	goto L40
L43:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v248)))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if int32(0) < v251 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v248 = v242 + int32(8)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v248 = v245 + int32(16)
	goto L43
L47:
	;
	v254 = base.F64_convert_i32_u(v251)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_nestloop_path[0])))
	if v256 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if base.F64_le(v239, float64(0)) != 0 {
		goto L59
	} else {
		goto L60
	}
L50:
	;
	v262 = base.F64_add(base.F64_mul(v254, float64(-0.3)), float64(1))
	if base.F64_gt(v262, float64(0)) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v268 = v254
	goto L52
L52:
	;
	v270 = float64(1e+100)
	v271 = base.F64_div(v249, v268)
	if base.F64_gt(v271, v270)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v271)&int64(9223372036854775807))) != 0 {
		v284 = v270
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v266 = v262
	goto L55
L54:
	;
	v266 = math.Float64frombits(uint64(0x8000000000000000))
	goto L55
L55:
	;
	v268 = base.F64_add(v266, v254)
	goto L52
L56:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v284
	goto L49
L57:
	;
	v280 = float64(1)
	if base.F64_le(v271, v280) != 0 {
		v284 = v280
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v284 = base.F64_nearest(v271)
	goto L56
L59:
	;
	v291 = float64(1)
	goto L61
L60:
	;
	v291 = v239
	goto L61
L61:
	;
	if base.F64_le(v237, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v295 = float64(1)
	goto L64
L63:
	;
	v295 = v237
	goto L64
L64:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v296&int32(-2) != int32(4) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v543 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v232)+24)) = v543
	v548 = float64(0)
	if v542 == int32(0) {
		v613 = v548
		v622 = v548
		goto L123
	} else {
		goto L124
	}
L66:
	;
	v535 = v234
	v538 = base.F64_mul(v295, v291)
	goto L65
L67:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v301 != int32(1) {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v304 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v306 = base.F64_nearest(base.F64_mul(v295, v304))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v312 = base.F64_div(float64(2), base.F64_add(v309, float64(1)))
	v313 = base.F64_mul(base.F64_mul(v291, v306), v312)
	v314 = base.F64_sub(v295, v306)
	v315 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	v316 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	if v317 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v496 = base.F64_ge(v314, float64(1))
	if v496 != 0 {
		goto L113
	} else {
		goto L114
	}
L72:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v318 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	switch v323 - int32(345) {
	case 0, 1:
		v330 = l6
		goto L74
	default:
		goto L71
	case 3:
		goto L75
	}
L74:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	if v331 == int32(0) {
		goto L71
	} else {
		goto L77
	}
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l6)+72))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v327 != int32(283) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v330 = v326
	goto L74
L77:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v334 <= int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330)+76))
	v338 = int32(0)
	v343 = v338
	v347 = v338
	goto L80
L79:
	;
	v462 = base.F64_add(base.F64_mul(v316, v312), v234)
	if base.F64_gt(v306, float64(1)) != 0 {
		goto L110
	} else {
		goto L111
	}
L80:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v343<<(uint(int32(2))%32))))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+8))
	v372 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	v374 = F_bms_is_subset(m, v373, v322)
	mBase = m.M
	if v374 == v372 {
		v385 = v372
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v347 == int32(0) {
		goto L71
	} else {
		goto L109
	}
L82:
	;
	if v385 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L82
L84:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v369)+28))
	v378 = F_bms_overlap(m, v371, v377)
	mBase = m.M
	if v378 == int32(0) {
		v385 = v372
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v369)+40))
	v382 = F_bms_overlap(m, v371, v381)
	mBase = m.M
	v385 = v382 ^ int32(1)
	goto L83
L86:
	;
	if v337 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	goto L88
L88:
	;
	v448 = v343 + int32(1)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v448 < v449 {
		v343 = v448
		goto L80
	} else {
		goto L108
	}
L89:
	;
	if v438 == int32(0) {
		goto L71
	} else {
		goto L106
	}
L90:
	;
	goto L89
L91:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v391 <= int32(0) {
		v438 = int32(0)
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v438 = int32(0)
	goto L90
L94:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v369)+60))
	v395 = int32(0)
	if v395 < v391 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v398 = v391
	goto L97
L96:
	;
	v398 = v395
	goto L97
L97:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v402 = int32(0)
	goto L98
L98:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v399+v402<<(uint(int32(2))%32))))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+12)))
	if v412 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L93
L100:
	;
	v423 = v402 + int32(1)
	if v423 != v398 {
		v402 = v423
		goto L98
	} else {
		goto L105
	}
L101:
	;
	v413 = int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	if v369 == v414 {
		v438 = v413
		goto L90
	} else {
		goto L102
	}
L102:
	;
	if v394 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+60))
	if v418 == v394 {
		v438 = v413
		goto L90
	} else {
		goto L104
	}
L104:
	;
	goto L100
L105:
	;
	goto L99
L106:
	;
	v442 = int32(1)
	v444 = v343 + v442
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v444 < v445 {
		v343 = v444
		v347 = v442
		goto L80
	} else {
		goto L107
	}
L107:
	;
	goto L79
L108:
	;
	goto L81
L109:
	;
	goto L79
L110:
	;
	v466 = base.F64_add(base.F64_mul(base.F64_mul(v315, base.F64_add(v306, float64(-1))), v312), v462)
	goto L112
L111:
	;
	v466 = v462
	goto L112
L112:
	;
	v535 = base.F64_add(base.F64_div(base.F64_mul(v315, v314), v291), v466)
	v538 = v313
	goto L65
L113:
	;
	v497 = v306
	goto L115
L114:
	;
	v497 = base.F64_add(v306, float64(-1))
	goto L115
L115:
	;
	v500 = base.F64_add(v234, v316)
	if base.F64_gt(v497, float64(0)) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v504 = base.F64_add(base.F64_mul(base.F64_mul(v315, v497), v312), v500)
	goto L118
L117:
	;
	v504 = v500
	goto L118
L118:
	;
	v506 = base.F64_add(base.F64_mul(v314, v291), v313)
	if v496 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v509 = base.F64_add(v314, float64(-1))
	goto L121
L120:
	;
	v509 = v314
	goto L121
L121:
	;
	if base.F64_gt(v509, float64(0)) == int32(0) {
		v535 = v504
		v538 = v506
		goto L65
	} else {
		goto L122
	}
L122:
	;
	v535 = base.F64_add(base.F64_mul(v509, v315), v504)
	v538 = v506
	goto L65
L123:
	;
	v624 = *(*float64)(unsafe.Add(mBase, _c_F_create_nestloop_path[1]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v626 = *(*float64)(unsafe.Add(mBase, uint32(v625)+24))
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v625)+16))
	v629 = base.F64_add(base.F64_add(v235, v622), v628)
	*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = v629
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v32)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+56)) = base.F64_add(v629, base.F64_add(base.F64_mul(v626, v631), base.F64_add(base.F64_mul(base.F64_add(v613, v624), v538), v535)))
	m.G0 = v232 + int32(32)
	m.G0 = v28 + int32(16)
	return v32
L124:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v553 <= int32(0) {
		v613 = v548
		v622 = float64(0)
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v560 = int32(0)
	goto L126
L126:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582+v560<<(uint(int32(2))%32))))
	v589 = F_cost_qual_eval_walker(m, v586, v232+int32(8))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	v595 = *(*float64)(unsafe.Add(mBase, uint32(v232)+24))
	v596 = *(*float64)(unsafe.Add(mBase, uint32(v232)+16))
	v613 = v595
	v622 = v596
	goto L123
L128:
	;
	v592 = v560 + int32(1)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v592 < v593 {
		v560 = v592
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
}
func F_create_secmsg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v4
	if v18 <= v4 {
		v101 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v108 = v18 + int32(3)
	v109 = F_palloc(m, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v24 = v18 & int32(3)
	v26 = l0 + int32(136)
	if base.Ui32(int32(4)) <= base.Ui32(v18) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = v4
	v38 = v4
	v42 = v4
	goto L6
L4:
	;
	v63 = v4
	v67 = v4
	goto L5
L5:
	;
	v76 = v63
	v80 = v67
	v85 = v4
	goto L10
L6:
	;
	v44 = v34 + v26
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)))
	v52 = v38 + v45 + v47 + v49 + v51
	v53 = int32(4)
	v54 = v34 + v53
	v56 = v42 + v53
	if v56 != v18&int32(2147483644) {
		v34 = v54
		v38 = v52
		v42 = v56
		goto L6
	} else {
		goto L8
	}
L7:
	;
	if v24 == int32(0) {
		v101 = v52
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v63 = v54
	v67 = v52
	goto L5
L10:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v26))))
	v88 = v80 + v87
	v89 = int32(1)
	v92 = v85 + v89
	if v92 != v24 {
		v76 = v76 + v89
		v80 = v88
		v85 = v92
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v101 = v88
	goto L1
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v113)
	if v18 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, v109+int32(1), l0+int32(136), v18)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v121 = int32(8)
	v127 = v101<<(uint(v121)%32) | int32(base.Ui32(v101&int32(_a_F_create_secmsg_0))>>(uint(v121)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v18+v109)+1)) = uint16(v127)
	v130 = l2 - v108
	v132 = v130 - int32(2)
	if v132 < v121 {
		v304 = int32(-12)
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v108 != 0 {
		goto L73
	} else {
		goto L74
	}
L19:
	;
	v135 = F_palloc(m, l2)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v137 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v137)
	v140 = v135 + int32(1)
	v141 = int32(0)
	v145 = m.G0
	v147 = v145 - int32(16)
	m.G0 = v147
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v141
	v153 = F_open(m, int32(_a_F_create_secmsg_1), v141, v147)
	mBase = m.M
	if v153 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v186 != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v156 = int32(1)
	if v132 == int32(0) {
		v179 = v156
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v186 = v141
	goto L24
L24:
	;
	m.G0 = v147 + int32(16)
	goto L21
L25:
	;
	v181 = F_close(m, v153)
	mBase = m.M
	v186 = v179
	goto L24
L26:
	;
	v159 = v140
	v160 = v132
	goto L27
L27:
	;
	v165 = F_read(m, v153, v159, v160)
	mBase = m.M
	if v165 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v179 = v156
	goto L25
L29:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v169 == int32(27) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v174 = v160 - v165
	if v174 != 0 {
		v159 = v159 + v165
		v160 = v174
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v179 = int32(0)
	goto L25
L33:
	;
	goto L28
L34:
	;
	v195 = v140
	goto L38
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v135)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L71
	}
L37:
	;
	if l2 != 0 {
		goto L68
	} else {
		goto L69
	}
L38:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v205 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v132)+1)) = uint8(v267)
	if v108 != 0 {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v263 = int32(1)
	goto L42
L41:
	;
	v208 = int32(0)
	v212 = m.G0
	v214 = v212 - int32(16)
	m.G0 = v214
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v208
	v220 = F_open(m, int32(_a_F_create_secmsg_1), v208, v214)
	mBase = m.M
	if v220 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v264 = v263 + v195
	if base.Ui32(v264) < base.Ui32(v140+v132) {
		v195 = v264
		goto L38
	} else {
		goto L57
	}
L43:
	;
	if v253 == int32(0) {
		goto L37
	} else {
		goto L56
	}
L44:
	;
	goto L48
L45:
	;
	v253 = v208
	goto L46
L46:
	;
	m.G0 = v214 + int32(16)
	goto L43
L47:
	;
	v248 = F_close(m, v220)
	mBase = m.M
	v253 = v246
	goto L46
L48:
	;
	v226 = v195
	v227 = int32(1)
	goto L49
L49:
	;
	v232 = F_read(m, v220, v226, v227)
	mBase = m.M
	if v232 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v246 = int32(1)
	goto L47
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v236 == int32(27) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v241 = v227 - v232
	if v241 != 0 {
		v226 = v226 + v232
		v227 = v241
		goto L49
	} else {
		goto L55
	}
L54:
	;
	v246 = int32(0)
	goto L47
L55:
	;
	goto L50
L56:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v263 = base.B2i32(v260 != int32(0))
	goto L42
L57:
	;
	goto L39
L58:
	;
	base.MemoryCopy(m, v135+v130, v109, v108)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v277 = F_pgp_mpi_create(m, v135, l2<<(uint(int32(3))%32)-int32(6), v16+int32(12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	if l2 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_pfree(m, v135)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L13
	} else {
		goto L66
	}
L63:
	;
	base.MemoryFill(m, v135, int32(0), l2)
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	v304 = v277
	goto L18
L67:
	;
	goto L36
L68:
	;
	base.MemoryFill(m, v135, int32(0), l2)
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v304 = int32(-17)
	goto L18
L72:
	;
	F_pfree(m, v109)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L76
	}
L73:
	;
	base.MemoryFill(m, v109, int32(0), v108)
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	if int32(0) <= v304 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v320
	goto L79
L78:
	;
	goto L79
L79:
	;
	m.G0 = v16 + int32(16)
	return v304
}
