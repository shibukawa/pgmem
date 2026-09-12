package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(25)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(26)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(27)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(28)
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
			*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(784)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = int32(785)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(786)
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(787)
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
			*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(527)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(528)
			*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(529)
			*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(530)
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
			*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(688)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(689)
			*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(690)
			*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(691)
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
			*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = int32(562)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = int32(563)
			*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = int32(564)
			*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(565)
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v269 float64
	_ = v269
	var v271 int32
	_ = v271
	var v273 float64
	_ = v273
	var v280 float64
	_ = v280
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int64
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v354 int32
	_ = v354
	var v361 float64
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int64
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	v12 = m.G0
	v14 = v12 - int32(1200)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+440)) = int32(1)
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	F_s_lock(m, v22+int32(440), int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_1), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+352))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v33)+344))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v33)+336))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)+328))
	goto L7
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+440)) = int32(0)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v48 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v43 = F__emscripten_memcpy_bulkmem(m, v14+int32(84), v33+int32(356), int32(76))
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	m.G0 = v14 + int32(1200)
	return v458
L11:
	;
	if v37 != int64(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
	v54 = base.B2i32(v52 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(v54)
	if v52 != int32(2) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v57 = int32(0)
	v60 = F_errstart(m, int32(13), v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v60 == int32(0) {
		v458 = v57
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_16), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_17), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v458 = v57
	goto L10
L20:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L36
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
	if base.Ui64(v77) < base.Ui64(v35) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = int32(0)
	v82 = F_errstart(m, int32(13), v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v35)
	v86 = int64(base.Ui64(v35) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v86)
	F_errmsg_internal(m, int32(_a_F_CreateRestartPoint_18), v14)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_UpdateMinRecoveryPoint(m, int64(0), int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_19), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if l0&int32(1) == int32(0) {
		v458 = v79
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v109 = F_LWLockAcquire(m, v105+int32(1152), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+16)) = int32(2)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v116, v112)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v120+int32(1152))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v458 = v79
	goto L10
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+152)) = v35
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3])) = v35
	F_WALInsertLockRelease(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+440)) = int32(1)
	if v136 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	F_s_lock(m, v140+int32(440), int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_3), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v150 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+440)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v149)+200)) = v35
	v157 = F__emscripten_memset_bulkmem(m, int32(_a_F_CreateRestartPoint_4), base.I32_extend8_s(v150), int32(80))
	mBase = m.M
	goto L42
L41:
	;
	goto L40
L42:
	;
	v162 = m.G0
	v163 = int32(16)
	v164 = v162 - v163
	m.G0 = v164
	F___gettimeofday(m, v164)
	mBase = m.M
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+8)))
	m.G0 = v164 + v163
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[4])) = v168 + v167*int64(1000000) - int64(946684800000000)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v179 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_LogCheckpointStart(m, l0, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l0&int32(3) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = int32(_a_F_CreateRestartPoint_5)
	if l0&int32(2) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_CheckPointGuts(m, v35, l0)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L58
	}
L51:
	;
	v193 = int32(_a_F_CreateRestartPoint_6)
	goto L53
L52:
	;
	v193 = int32(_a_F_CreateRestartPoint_7)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v193
	if l0&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v199 = int32(_a_F_CreateRestartPoint_8)
	goto L56
L55:
	;
	v199 = int32(_a_F_CreateRestartPoint_7)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v199
	v207 = F_pg_snprintf(m, v14+int32(160), int32(128), int32(_a_F_CreateRestartPoint_9), v14+int32(48))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v211 = F_strlen(m, v14+int32(160))
	mBase = m.M
	goto L50
L58:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v215)+40))
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	v222 = F_LWLockAcquire(m, v218+int32(1152), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[2]))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v225)+40))
	if base.Ui64(v226) < base.Ui64(v35) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+48)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v225)+40)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v225)+32)) = v37
	goto L64
L61:
	;
	goto L62
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[6]))
	F_LWLockRelease(m, v258+int32(1152))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L74
	}
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	if v238 != int32(5) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v236 = F__emscripten_memcpy_bulkmem(m, v225+int32(52), v14+int32(84), int32(76))
	mBase = m.M
	goto L66
L66:
	;
	goto L63
L67:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[8]))
	F_update_controlfile(m, v254, v225)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L73
	}
L68:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v225)+136))
	if base.Ui64(v241) < base.Ui64(v36) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+144)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v225)+136)) = v36
	*(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[7])) = v36
	goto L71
L70:
	;
	goto L71
L71:
	;
	if l0&int32(1) == int32(0) {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = int32(2)
	goto L67
L73:
	;
	goto L62
L74:
	;
	v264 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	if v216 != int64(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v269 = base.F64_convert_i64_u(v264 - v216)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[9])) = v269
	v271 = int32(_a_F_CreateRestartPoint_10)
	v273 = *(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10]))
	if base.F64_gt(v269, v273) != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v285 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v286 = base.I64_div_u_s(v264, v285)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v286
	v288 = int32(0)
	v290 = F_GetWalRcvFlushRecPtr(m, v288, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L81
	}
L78:
	;
	v280 = v269
	goto L80
L79:
	;
	v280 = base.F64_add(base.F64_mul(v273, float64(0.9)), base.F64_mul(v269, float64(0.1)))
	goto L80
L80:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[10])) = v280
	goto L77
L81:
	;
	v294 = F_GetXLogReplayRecPtr(m, v14+int32(80))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	if base.Ui64(v294) < base.Ui64(v290) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v297 = v290
	goto L85
L84:
	;
	v297 = v294
	goto L85
L85:
	;
	F_KeepLogSeg(m, v297, v14+int32(72))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	v304 = int32(0)
	v306 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v303, v304, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v306 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v309 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	v311 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11])))
	v312 = base.I64_div_u_s(v309, v311)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v312
	F_KeepLogSeg(m, v297, v14+int32(72))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v319 = v303
	goto L90
L90:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])))
	if v325 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	v319 = v318
	goto L90
L92:
	;
	v342 = *(*int64)(unsafe.Add(mBase, _c_F_CreateRestartPoint[3]))
	F_RemoveOldXlogFiles(m, v319-int64(1), v342, v297, v339)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L96
	}
L93:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v323)+308))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v337
	v339 = v337
	goto L92
L94:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323)+316))
	v330 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[1])) = uint8(base.B2i32(v329 != v330))
	if v329 == v330 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v339 = v335
	goto L92
L96:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[0]))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+320)))
	if v347 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[12])))
	if v400 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L98:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v352 = v297 - int64(1)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[11]))
	v361 = base.F64_mul(base.F64_convert_i32_s(v354), float64(0.75))
	if base.F64_lt(v361, float64(4.294967296e+09))&base.F64_ge(v361, float64(0)) != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if base.Ui64(v352&base.I64_extend_i32_s(v354-int32(1))) < base.Ui64(base.I64_extend_i32_u(v369)) {
		goto L97
	} else {
		goto L103
	}
L100:
	;
	v367 = base.I32_trunc_f64_u(v361)
	v369 = v367
	goto L99
L101:
	;
	goto L102
L102:
	;
	v369 = int32(0)
	goto L99
L103:
	;
	v373 = base.I64_div_u_s(v352, base.I64_extend_i32_s(v354))
	v380 = F_XLogFileInitInternal(m, v373+int64(1), v350, v14+int32(1199), v14+int32(160))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v380 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v384 = F_close(m, v380)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1199)))
	if v385 != int32(1) {
		goto L97
	} else {
		goto L108
	}
L108:
	;
	v388 = int32(_a_F_CreateRestartPoint_15)
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[14])) = v390 + int32(1)
	goto L97
L109:
	;
	v403 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_LogCheckpointEnd(m, int32(1))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	F_TruncateSUBTRANS(m, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v410 = F_GetLatestXTime(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateRestartPoint[5])))
	if v415 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v416 = int32(15)
	goto L118
L117:
	;
	v416 = int32(13)
	goto L118
L118:
	;
	v418 = F_errstart(m, v416, int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	if v418 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+36)) = uint32(v35)
	v422 = int64(base.Ui64(v35) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+32)) = uint32(v422)
	F_errmsg(m, int32(_a_F_CreateRestartPoint_11), v14+int32(32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v444 = int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_CreateRestartPoint[13]))
	if v446 == int32(0) {
		v458 = v444
		goto L10
	} else {
		goto L130
	}
L123:
	;
	if v410 != int64(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v431 = F_timestamptz_to_str(m, v410)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_errfinish(m, int32(_a_F_CreateRestartPoint_0), int32(_a_F_CreateRestartPoint_13), int32(_a_F_CreateRestartPoint_2))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v431
	F_errdetail(m, int32(_a_F_CreateRestartPoint_12), v14+int32(16))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	goto L122
L130:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v449 == int32(0) {
		v458 = v444
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_ExecuteRecoveryCommand(m, v446, int32(_a_F_CreateRestartPoint_14), int32(0), int32(134217729))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v458 = v444
	goto L10
}
func F_CreateStatistics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
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
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
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
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1331 int32
	_ = v1331
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1556 int32
	_ = v1556
	var v1570 int32
	_ = v1570
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int64
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1720 int32
	_ = v1720
	var v1744 int32
	_ = v1744
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1767 int32
	_ = v1767
	var v1795 int32
	_ = v1795
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1840 int64
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(432)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v27 == v4 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L25
	} else {
		goto L441
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L25
	} else {
		goto L437
	}
L3:
	;
	v1840 = *(*int64)(unsafe.Add(mBase, uint32(v1824)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1840
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1842
	m.G0 = v23 + int32(432)
	return
L4:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1097 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L5:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1047 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L262
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L25
	} else {
		goto L258
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = v4
	goto L22
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L25
	} else {
		goto L253
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L25
	} else {
		goto L249
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L25
	} else {
		goto L243
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L25
	} else {
		goto L239
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L25
	} else {
		goto L235
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L25
	} else {
		goto L230
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L25
	} else {
		goto L226
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L25
	} else {
		goto L222
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L25
	} else {
		goto L218
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L25
	} else {
		goto L214
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L25
	} else {
		goto L210
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L25
	} else {
		goto L205
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L25
	} else {
		goto L201
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v41<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != int32(3) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v127 != 0 {
		goto L51
	} else {
		goto L52
	}
L24:
	;
	v66 = F_relation_openrv(m, v61, int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	v71 = v69 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v71) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if int32(1)<<(uint(v71)%32)&int32(_a_F_CreateStatistics_0) == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v82 = F_object_ownercheck(m, int32(1259), v81, v26)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	if v82 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v87)+119)))
	switch v88 - int32(73) {
	case 0, 32:
		goto L39
	default:
		v98 = int32(41)
		goto L34
	case 10:
		goto L38
	case 29:
		goto L35
	case 36:
		goto L36
	case 45:
		goto L37
	}
L31:
	;
	goto L32
L32:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[1])))
	if v107 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	F_aclcheck_error(m, int32(2), v100, v101+int32(4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L25
	} else {
		goto L40
	}
L34:
	;
	v100 = v98
	goto L33
L35:
	;
	v98 = int32(18)
	goto L34
L36:
	;
	v100 = int32(23)
	goto L33
L37:
	;
	v100 = int32(51)
	goto L33
L38:
	;
	v100 = int32(37)
	goto L33
L39:
	;
	v100 = int32(20)
	goto L33
L40:
	;
	goto L32
L41:
	;
	v111 = int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	if base.Ui32(v112) < base.Ui32(int32(_a_F_CreateStatistics_1)) {
		v121 = v111
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v123 = v41 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v123 < v124 {
		v41 = v123
		goto L22
	} else {
		goto L49
	}
L44:
	;
	if v121 != 0 {
		goto L19
	} else {
		goto L48
	}
L45:
	;
	goto L44
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
	if v116 == int32(99) {
		v121 = v111
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v119 = F_isTempToastNamespace(m, v116)
	mBase = m.M
	v121 = v119
	goto L45
L48:
	;
	goto L43
L49:
	;
	goto L23
L50:
	;
	v445 = F_strncpy(m, v23+int32(304), v426, int32(64))
	mBase = m.M
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v445)+63)) = uint8(v446)
	goto L116
L51:
	;
	v130 = F_QualifiedNameGetCreationNamespace(m, v127, v23+int32(284))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L25
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+68))
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+304)) = uint8(v136)
	v140 = v134 + int32(4)
	if v133 == v136 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	v426 = v132
	v434 = v130
	goto L50
L55:
	;
	v339 = F_pstrdup(m, v23+int32(304))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L25
	} else {
		goto L103
	}
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v143 <= int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v151 = v136
	v152 = int32(0)
	v155 = v143
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v151<<(uint(int32(2))%32))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v172 == int32(206) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L55
L60:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if int32(0) < v152 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v311 = v152
	v312 = v155
	goto L62
L62:
	;
	v315 = v151 + int32(1)
	if v315 < v312 {
		v151 = v315
		v152 = v311
		v155 = v312
		goto L58
	} else {
		goto L102
	}
L63:
	;
	v181 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(304)+v152))) = uint8(v181)
	v185 = v152 + int32(1)
	goto L65
L64:
	;
	v185 = v152
	goto L65
L65:
	;
	v188 = v23 + int32(304) + v185
	if v175 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v190 = v175
	goto L68
L67:
	;
	v190 = int32(_a_F_CreateStatistics_2)
	goto L68
L68:
	;
	goto L72
L69:
	;
	v306 = F_strlen(m, v188)
	mBase = m.M
	v307 = v306 + v185
	if int32(63) < v307 {
		goto L55
	} else {
		goto L101
	}
L70:
	;
	v303 = F_strlen(m, v292)
	mBase = m.M
	goto L69
L72:
	;
	goto L73
L73:
	;
	v197 = int32(63)
	if (v188^v190)&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v293))) = uint8(v296)
	goto L70
L75:
	;
	v277 = v272
	v278 = v273
	v279 = v274
	goto L97
L76:
	;
	if v267 == int32(0) {
		v292 = v265
		v293 = v266
		goto L74
	} else {
		goto L96
	}
L77:
	;
	v265 = v190
	v266 = v188
	v267 = v197
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v190&int32(3) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v234 == int32(0) {
		v292 = v231
		v293 = v232
		goto L74
	} else {
		goto L89
	}
L81:
	;
	v231 = v190
	v232 = v188
	v233 = v197
	v234 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v210 = v190
	v211 = v188
	v212 = v197
	goto L84
L84:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v214)
	if v214 == int32(0) {
		v272 = v210
		v273 = v211
		v274 = v212
		goto L75
	} else {
		goto L86
	}
L85:
	;
	v231 = v225
	v232 = v219
	v233 = v221
	v234 = v223
	goto L80
L86:
	;
	v218 = int32(1)
	v219 = v211 + v218
	v221 = v212 - v218
	v222 = int32(0)
	v223 = base.B2i32(v221 != v222)
	v225 = v210 + v218
	if v225&int32(3) == v222 {
		v231 = v225
		v232 = v219
		v233 = v221
		v234 = v223
		goto L80
	} else {
		goto L87
	}
L87:
	;
	if v221 != 0 {
		v210 = v225
		v211 = v219
		v212 = v221
		goto L84
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v237 == int32(0) {
		v265 = v231
		v266 = v232
		v267 = v233
		goto L76
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v233) < base.Ui32(int32(4)) {
		v265 = v231
		v266 = v232
		v267 = v233
		goto L76
	} else {
		goto L91
	}
L91:
	;
	v243 = v231
	v244 = v232
	v245 = v233
	goto L92
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v251 = int32(-2139062144)
	if (int32(16843008)-v248|v248)&v251 != v251 {
		v272 = v243
		v273 = v244
		v274 = v245
		goto L75
	} else {
		goto L94
	}
L93:
	;
	v265 = v259
	v266 = v257
	v267 = v261
	goto L76
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v248
	v256 = int32(4)
	v257 = v244 + v256
	v259 = v243 + v256
	v261 = v245 - v256
	if base.Ui32(int32(3)) < base.Ui32(v261) {
		v243 = v259
		v244 = v257
		v245 = v261
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v272 = v265
	v273 = v266
	v274 = v267
	goto L75
L97:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	*(*uint8)(unsafe.Add(mBase, uint32(v278))) = uint8(v281)
	if v281 == int32(0) {
		v292 = v277
		v293 = v278
		goto L74
	} else {
		goto L99
	}
L98:
	;
	v292 = v288
	v293 = v286
	goto L74
L99:
	;
	v285 = int32(1)
	v286 = v278 + v285
	v288 = v277 + v285
	v290 = v279 - v285
	if v290 != 0 {
		v277 = v288
		v278 = v286
		v279 = v290
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v311 = v307
	v312 = v310
	goto L62
L102:
	;
	goto L59
L103:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+308)) = uint8(v343)
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+304)) = v346
	v351 = F_makeObjectName(m, v140, v339, v23+int32(304))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L25
	} else {
		goto L104
	}
L104:
	;
	v353 = int32(0)
	v355 = F_GetSysCacheOid(m, int32(63), v351, v135, v353, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L25
	} else {
		goto L105
	}
L105:
	;
	if v355 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v361 = v351
	v362 = int32(0)
	goto L109
L107:
	;
	v405 = v351
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+284)) = v405
	v426 = v405
	v434 = v135
	goto L50
L109:
	;
	F_pfree(m, v361)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L111
	}
L110:
	;
	v405 = v395
	goto L108
L111:
	;
	v380 = v362 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = int32(_a_F_CreateStatistics_3)
	v390 = F_pg_snprintf(m, v23+int32(304), int32(64), int32(_a_F_CreateStatistics_4), v23+int32(144))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	v395 = F_makeObjectName(m, v140, v339, v23+int32(304))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	v397 = int32(0)
	v399 = F_GetSysCacheOid(m, int32(63), v395, v135, v397, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L25
	} else {
		goto L114
	}
L114:
	;
	if v399 != 0 {
		v361 = v395
		v362 = v380
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	if l2 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v464 = int32(0)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	v469 = F_SearchSysCacheExists(m, int32(63), v466, v434, v464, v464)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L25
	} else {
		goto L123
	}
L118:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[0]))
	v454 = F_object_aclcheck(m, int32(2615), v434, v452, int64(512))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L25
	} else {
		goto L119
	}
L119:
	;
	if v454 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v459 = F_get_namespace_name(m, v434)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L25
	} else {
		goto L121
	}
L121:
	;
	F_aclcheck_error(m, v454, int32(36), v459)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L25
	} else {
		goto L122
	}
L122:
	;
	goto L117
L123:
	;
	if v469 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+25)))
	if v471 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v516 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L127:
	;
	v476 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L25
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L25
	} else {
		goto L138
	}
L130:
	;
	if v476 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_5))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L25
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L25
	} else {
		goto L137
	}
L134:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v481
	F_errmsg(m, int32(_a_F_CreateStatistics_6), v23+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L25
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(209), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L25
	} else {
		goto L136
	}
L136:
	;
	goto L133
L137:
	;
	v1824 = int32(_a_F_CreateStatistics_9)
	goto L3
L138:
	;
	F_errcode(m, int32(_a_F_CreateStatistics_5))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L25
	} else {
		goto L139
	}
L139:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v504
	F_errmsg(m, int32(_a_F_CreateStatistics_10), v23+int32(32))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L25
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(216), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L25
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	v1082 = int32(0)
	v1086 = v464
	v1096 = v4
	goto L4
L143:
	;
	goto L144
L144:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if int32(9) <= v520 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	v523 = int32(0)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v524 <= v523 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v1032 = int32(0)
	v1036 = v464
	goto L5
L147:
	;
	goto L148
L148:
	;
	v534 = int32(0)
	v536 = v523
	v538 = v464
	goto L149
L149:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v536<<(uint(int32(2))%32))))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	if v554 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v1032 = v747
	v1036 = v751
	goto L5
L151:
	;
	v763 = v536 + int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	if v763 < v764 {
		v534 = v747
		v536 = v763
		v538 = v751
		goto L149
	} else {
		goto L200
	}
L152:
	;
	v555 = F_SearchSysCacheAttName(m, v126, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L25
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	if v587 == int32(6) {
		goto L162
	} else {
		goto L163
	}
L155:
	;
	if v555 == int32(0) {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555)+16))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+22)))
	v561 = v559 + v560
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v561)+74)))
	if v562 <= int32(0) {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+90)))
	if v565 == int32(118) {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v561)+68))
	v570 = F_lookup_type_cache(m, v568, int32(2))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L25
	} else {
		goto L159
	}
L159:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v570)+56))
	if v572 == int32(0) {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	v580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v561)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v538<<(uint(int32(1))%32)))) = uint16(v580)
	F_ReleaseCatCache(m, v555)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L25
	} else {
		goto L161
	}
L161:
	;
	v747 = v534
	v751 = v538 + int32(1)
	goto L151
L162:
	;
	v590 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586)+8)))
	if v590 <= int32(0) {
		goto L13
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = int32(0)
	F_pull_varattnos(m, v586, int32(1), v23+int32(240))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L25
	} else {
		goto L170
	}
L165:
	;
	v593 = F_get_attgenerated(m, v126, v590)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L25
	} else {
		goto L166
	}
L166:
	;
	if v593 == int32(118) {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v599 = F_lookup_type_cache(m, v597, int32(2))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L25
	} else {
		goto L168
	}
L168:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v599)+56))
	if v601 == int32(0) {
		goto L11
	} else {
		goto L169
	}
L169:
	;
	v606 = int32(1)
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v586)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v538<<(uint(v606)%32)))) = uint16(v609)
	v747 = v534
	v751 = v538 + v606
	goto L151
L170:
	;
	v625 = int32(-1)
	goto L172
L171:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v725 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L172:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v23)+240))
	if v641 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L25
	} else {
		goto L189
	}
L174:
	;
	if v697 < int32(0) {
		goto L171
	} else {
		goto L185
	}
L175:
	;
	v697 = base.I32_ctz(v683) | v684<<(uint(int32(5))%32)
	goto L174
L176:
	;
	v697 = int32(-2)
	goto L174
L177:
	;
	v648 = v625 + int32(1)
	v650 = base.I32_div_s(v648, int32(32))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v651 <= v650 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v654 = v641 + int32(8)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v650<<(uint(int32(2))%32))))
	v661 = v658 & (int32(-1) << (uint(v648) % 32))
	if v661 != 0 {
		v683 = v661
		v684 = v650
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v663 = v650 + int32(1)
	if v663 == v651 {
		goto L176
	} else {
		goto L180
	}
L180:
	;
	v666 = v663
	goto L181
L181:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v654+v666<<(uint(int32(2))%32))))
	if v673 != 0 {
		v683 = v673
		v684 = v666
		goto L175
	} else {
		goto L183
	}
L182:
	;
	goto L176
L183:
	;
	v675 = v666 + int32(1)
	if v675 != v651 {
		v666 = v675
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v702 = base.I32_extend16_s(v697 - int32(7))
	if v702 <= int32(0) {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	v705 = F_get_attgenerated(m, v126, v702)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L25
	} else {
		goto L187
	}
L187:
	;
	if v705 != int32(118) {
		v625 = v697
		goto L172
	} else {
		goto L188
	}
L188:
	;
	goto L173
L189:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L25
	} else {
		goto L190
	}
L190:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L25
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(343), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L25
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v740 = F_lappend(m, v534, v586)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L25
	} else {
		goto L199
	}
L194:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	if v728 < int32(2) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v731 = F_exprType(m, v586)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L25
	} else {
		goto L196
	}
L196:
	;
	v734 = F_lookup_type_cache(m, v731, int32(2))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L25
	} else {
		goto L197
	}
L197:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v734)+56))
	if v736 == int32(0) {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	goto L193
L199:
	;
	v747 = v740
	v751 = v538
	goto L151
L200:
	;
	goto L150
L201:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L25
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_12), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L25
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(115), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L25
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L25
	} else {
		goto L206
	}
L206:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v789 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_13), v23)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L25
	} else {
		goto L207
	}
L207:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v797 = int32(*(*int8)(unsafe.Add(mBase, uint32(v796)+119)))
	F_errdetail_relkind_not_supported(m, v797)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L25
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(135), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L25
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L25
	} else {
		goto L211
	}
L211:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v812 + int32(4)
	F_errmsg(m, int32(_a_F_CreateStatistics_14), v23+int32(160))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L25
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(153), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L25
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L25
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(8)
	F_errmsg(m, int32(_a_F_CreateStatistics_15), v23-int32(-64))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L25
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(228), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L25
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L25
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v554
	F_errmsg(m, int32(_a_F_CreateStatistics_16), v23+int32(112))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L25
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(261), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L25
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L25
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L25
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(268), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L25
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L25
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L25
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(274), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L25
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L25
	} else {
		goto L231
	}
L231:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v561)+68))
	v903 = F_format_type_be(m, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L25
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v554
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v23+int32(128))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L25
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(282), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L25
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
	v923 = m.ExcPending
	if v923 != 0 {
		goto L25
	} else {
		goto L236
	}
L236:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L25
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(297), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L25
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
	v939 = m.ExcPending
	if v939 != 0 {
		goto L25
	} else {
		goto L240
	}
L240:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_11), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L25
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(303), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L25
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L25
	} else {
		goto L244
	}
L244:
	;
	v956 = int32(*(*int16)(unsafe.Add(mBase, uint32(v586)+8)))
	v958 = F_get_attname(m, v126, v956, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L25
	} else {
		goto L245
	}
L245:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v961 = F_format_type_be(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L25
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v961
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v958
	F_errmsg(m, int32(_a_F_CreateStatistics_18), v23+int32(80))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L25
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(311), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L25
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L25
	} else {
		goto L250
	}
L250:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_17), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L25
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(337), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L25
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L25
	} else {
		goto L254
	}
L254:
	;
	v998 = F_format_type_be(m, v731)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L25
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v998
	F_errmsg(m, int32(_a_F_CreateStatistics_19), v23+int32(96))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L25
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(361), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L25
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L25
	} else {
		goto L259
	}
L259:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_12), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L25
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(106), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L25
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1050 != int32(1) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L263
	}
L263:
	;
	if v1032 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+4))
	if v1055 != int32(1) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L265
	}
L265:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v1058 == int32(0) {
		v1082 = v1032
		v1086 = v1036
		v1096 = v520
		goto L4
	} else {
		goto L266
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L25
	} else {
		goto L267
	}
L267:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L25
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_20), int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L25
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(381), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L25
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	v1384 = base.B2i32(v1082 == int32(0))
	if v1096 < int32(2) {
		goto L353
	} else {
		goto L354
	}
L272:
	;
	v1110 = int32(0)
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1114 = int32(_a_F_CreateStatistics_21)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1118 == v1110 {
		v1137 = v1117
		v1138 = v1118
		goto L281
	} else {
		goto L282
	}
L273:
	;
	v1365 = v1107
	v1376 = v1109
	v1379 = v4
	v1380 = v4
	goto L271
L274:
	;
	v1107 = int32(1)
	v1109 = int32(0)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if int32(0) < v1103 {
		goto L272
	} else {
		goto L277
	}
L277:
	;
	v1107 = int32(1)
	v1109 = int32(0)
	goto L273
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L25
	} else {
		goto L348
	}
L279:
	;
	v1201 = base.B2i32(v1139 == int32(0))
	v1202 = int32(1)
	if v1103 == v1202 {
		v1365 = v1110
		v1376 = v1198
		v1379 = v1199
		v1380 = v1201
		goto L271
	} else {
		goto L309
	}
L280:
	;
	if v1139 == int32(0) {
		v1198 = v4
		v1199 = v4
		goto L279
	} else {
		goto L288
	}
L281:
	;
	v1139 = v1138 - v1137
	goto L280
L282:
	;
	if v1117 != v1118 {
		v1137 = v1117
		v1138 = v1118
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1122 = v1113
	v1123 = v1114
	goto L284
L284:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123)+1)))
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+1)))
	if v1127 == int32(0) {
		v1137 = v1126
		v1138 = v1127
		goto L281
	} else {
		goto L286
	}
L285:
	;
	v1137 = v1126
	v1138 = v1127
	goto L281
L286:
	;
	v1130 = int32(1)
	if v1126 == v1127 {
		v1122 = v1122 + v1130
		v1123 = v1123 + v1130
		goto L284
	} else {
		goto L287
	}
L287:
	;
	goto L285
L288:
	;
	v1142 = int32(_a_F_CreateStatistics_22)
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1146 == int32(0) {
		v1165 = v1145
		v1166 = v1146
		goto L290
	} else {
		goto L291
	}
L289:
	;
	if v1166-v1165 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L290:
	;
	goto L289
L291:
	;
	if v1145 != v1146 {
		v1165 = v1145
		v1166 = v1146
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1150 = v1113
	v1151 = v1142
	goto L293
L293:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+1)))
	if v1155 == int32(0) {
		v1165 = v1154
		v1166 = v1155
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1165 = v1154
	v1166 = v1155
	goto L290
L295:
	;
	v1158 = int32(1)
	if v1154 == v1155 {
		v1150 = v1150 + v1158
		v1151 = v1151 + v1158
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v1198 = v4
	v1199 = int32(1)
	goto L279
L298:
	;
	goto L299
L299:
	;
	v1171 = int32(_a_F_CreateStatistics_23)
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	if v1175 == int32(0) {
		v1194 = v1174
		v1195 = v1175
		goto L301
	} else {
		goto L302
	}
L300:
	;
	if v1195-v1194 != 0 {
		v1331 = v1113
		goto L278
	} else {
		goto L308
	}
L301:
	;
	goto L300
L302:
	;
	if v1174 != v1175 {
		v1194 = v1174
		v1195 = v1175
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1179 = v1113
	v1180 = v1171
	goto L304
L304:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+1)))
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+1)))
	if v1184 == int32(0) {
		v1194 = v1183
		v1195 = v1184
		goto L301
	} else {
		goto L306
	}
L305:
	;
	v1194 = v1183
	v1195 = v1184
	goto L301
L306:
	;
	v1187 = int32(1)
	if v1183 == v1184 {
		v1179 = v1179 + v1187
		v1180 = v1180 + v1187
		goto L304
	} else {
		goto L307
	}
L307:
	;
	goto L305
L308:
	;
	v1198 = int32(1)
	v1199 = v4
	goto L279
L309:
	;
	v1205 = int32(0)
	if v1205 < v1103 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1208 = v1103
	goto L312
L311:
	;
	v1208 = v1205
	goto L312
L312:
	;
	v1213 = v1202
	v1222 = v1198
	v1225 = v1199
	v1226 = v1201
	goto L313
L313:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1111+v1213<<(uint(int32(2))%32))))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+4))
	v1234 = int32(_a_F_CreateStatistics_21)
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[4])))
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1238 == int32(0) {
		v1257 = v1237
		v1258 = v1238
		goto L317
	} else {
		goto L318
	}
L314:
	;
	v1365 = v1110
	v1376 = v1319
	v1379 = v1320
	v1380 = v1321
	goto L271
L315:
	;
	v1323 = v1213 + int32(1)
	if v1208 != v1323 {
		v1213 = v1323
		v1222 = v1319
		v1225 = v1320
		v1226 = v1321
		goto L313
	} else {
		goto L347
	}
L316:
	;
	if v1258-v1257 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L317:
	;
	goto L316
L318:
	;
	if v1237 != v1238 {
		v1257 = v1237
		v1258 = v1238
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1242 = v1233
	v1243 = v1234
	goto L320
L320:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+1)))
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
	if v1247 == int32(0) {
		v1257 = v1246
		v1258 = v1247
		goto L317
	} else {
		goto L322
	}
L321:
	;
	v1257 = v1246
	v1258 = v1247
	goto L317
L322:
	;
	v1250 = int32(1)
	if v1246 == v1247 {
		v1242 = v1242 + v1250
		v1243 = v1243 + v1250
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v1319 = v1222
	v1320 = v1225
	v1321 = int32(1)
	goto L315
L325:
	;
	goto L326
L326:
	;
	v1263 = int32(_a_F_CreateStatistics_22)
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[5])))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1267 == int32(0) {
		v1286 = v1266
		v1287 = v1267
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if v1287-v1286 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L328:
	;
	goto L327
L329:
	;
	if v1266 != v1267 {
		v1286 = v1266
		v1287 = v1267
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1271 = v1233
	v1272 = v1263
	goto L331
L331:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272)+1)))
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+1)))
	if v1276 == int32(0) {
		v1286 = v1275
		v1287 = v1276
		goto L328
	} else {
		goto L333
	}
L332:
	;
	v1286 = v1275
	v1287 = v1276
	goto L328
L333:
	;
	v1279 = int32(1)
	if v1275 == v1276 {
		v1271 = v1271 + v1279
		v1272 = v1272 + v1279
		goto L331
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1319 = v1222
	v1320 = int32(1)
	v1321 = v1226
	goto L315
L336:
	;
	goto L337
L337:
	;
	v1292 = int32(_a_F_CreateStatistics_23)
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateStatistics[6])))
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1296 == int32(0) {
		v1315 = v1295
		v1316 = v1296
		goto L339
	} else {
		goto L340
	}
L338:
	;
	if v1316-v1315 != 0 {
		v1331 = v1233
		goto L278
	} else {
		goto L346
	}
L339:
	;
	goto L338
L340:
	;
	if v1295 != v1296 {
		v1315 = v1295
		v1316 = v1296
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v1300 = v1233
	v1301 = v1292
	goto L342
L342:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301)+1)))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+1)))
	if v1305 == int32(0) {
		v1315 = v1304
		v1316 = v1305
		goto L339
	} else {
		goto L344
	}
L343:
	;
	v1315 = v1304
	v1316 = v1305
	goto L339
L344:
	;
	v1308 = int32(1)
	if v1304 == v1305 {
		v1300 = v1300 + v1308
		v1301 = v1301 + v1308
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	v1319 = int32(1)
	v1320 = v1225
	v1321 = v1226
	goto L315
L347:
	;
	goto L314
L348:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L25
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1331
	F_errmsg(m, int32(_a_F_CreateStatistics_24), v23+int32(48))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L25
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(411), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L25
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_pg_qsort(m, v23+int32(288), v1086, int32(2), int32(571))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L25
	} else {
		goto L359
	}
L353:
	;
	if int32(1) < v1096 {
		v1400 = v1376
		v1401 = v1384
		v1402 = v1379
		v1403 = v1380
		goto L352
	} else {
		goto L356
	}
L354:
	;
	if v1365 == int32(0) {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1389 = int32(1)
	v1400 = v1389
	v1401 = v1384
	v1402 = v1389
	v1403 = v1389
	goto L352
L356:
	;
	if v1082 == int32(0) {
		goto L2
	} else {
		goto L357
	}
L357:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1397 != int32(1) {
		goto L2
	} else {
		goto L358
	}
L358:
	;
	v1400 = v1376
	v1401 = int32(0)
	v1402 = v1379
	v1403 = v1380
	goto L352
L359:
	;
	if int32(2) <= v1086 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+288)))
	v1418 = int32(1)
	v1420 = v1412
	goto L363
L361:
	;
	goto L362
L362:
	;
	if v1401 != 0 {
		goto L367
	} else {
		goto L368
	}
L363:
	;
	v1441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(288)+v1418<<(uint(int32(1))%32)))))
	if v1420&int32(_a_F_CreateStatistics_25) == v1441 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	goto L362
L365:
	;
	v1444 = v1418 + int32(1)
	if v1444 != v1086 {
		v1418 = v1444
		v1420 = v1441
		goto L363
	} else {
		goto L366
	}
L366:
	;
	goto L364
L367:
	;
	v1594 = F_buildint2vector(m, v23+int32(288), v1086)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L25
	} else {
		goto L384
	}
L368:
	;
	v1466 = int32(0)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1467 <= v1466 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1472 = v1466
	v1477 = v1467
	goto L370
L370:
	;
	v1490 = int32(0)
	if v1477 <= v1490 {
		v1556 = v1477
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L367
L372:
	;
	v1570 = v1472 + int32(1)
	if v1570 < v1556 {
		v1472 = v1570
		v1477 = v1556
		goto L370
	} else {
		goto L383
	}
L373:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1493+v1472<<(uint(int32(2))%32))))
	v1503 = v1490
	v1505 = int32(0)
	goto L374
L374:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+12))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1519+v1503<<(uint(int32(2))%32))))
	v1524 = F_equal(m, v1497, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L25
	} else {
		goto L376
	}
L375:
	;
	if v1526 <= int32(1) {
		v1556 = v1529
		goto L372
	} else {
		goto L378
	}
L376:
	;
	v1526 = v1524 + v1505
	v1528 = v1503 + int32(1)
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	if v1528 < v1529 {
		v1503 = v1528
		v1505 = v1526
		goto L374
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L25
	} else {
		goto L379
	}
L379:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L25
	} else {
		goto L380
	}
L380:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_26), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L25
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(492), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L25
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	goto L371
L384:
	;
	if v1403 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	if v1402 != 0 {
		goto L389
	} else {
		goto L390
	}
L386:
	;
	v1604 = v23 + int32(176)
	v1605 = int32(0)
	goto L385
L387:
	;
	goto L388
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = int32(100)
	v1604 = v23 + int32(176) | int32(4)
	v1605 = int32(1)
	goto L385
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1604))) = int32(102)
	v1610 = v1605 + int32(1)
	goto L391
L390:
	;
	v1610 = v1605
	goto L391
L391:
	;
	if v1400 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(176)|v1610<<(uint(int32(2))%32)))) = int32(109)
	v1620 = v1610 + int32(1)
	goto L394
L393:
	;
	v1620 = v1610
	goto L394
L394:
	;
	if v1401 != 0 {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	v1652 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L25
	} else {
		goto L404
	}
L396:
	;
	v1625 = F_construct_array_builtin(m, v23+int32(176), v1620, int32(18))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L25
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1628 = v23 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v1628+v1620<<(uint(int32(2))%32)))) = int32(101)
	v1639 = F_construct_array_builtin(m, v1628, v1620+int32(1), int32(18))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L25
	} else {
		goto L400
	}
L399:
	;
	v1647 = v1625
	v1649 = int32(0)
	goto L395
L400:
	;
	v1641 = F_nodeToString(m, v1082)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L25
	} else {
		goto L401
	}
L401:
	;
	v1643 = F_cstring_to_text(m, v1641)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L25
	} else {
		goto L402
	}
L402:
	;
	F_pfree(m, v1641)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L25
	} else {
		goto L403
	}
L403:
	;
	v1647 = v1639
	v1649 = v1643
	goto L395
L404:
	;
	v1654 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+264)) = v1654
	v1656 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+232)) = uint8(v1656)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+256)) = v1654
	*(*int64)(unsafe.Add(mBase, uint32(v23)+224)) = v1654
	v1664 = F_GetNewOidWithIndex(m, v1652, int32(3380), int32(1))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L25
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+260)) = v1594
	*(*int32)(unsafe.Add(mBase, uint32(v23)+256)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+252)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v23)+244)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = v1664
	v1671 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+230)) = uint8(v1671)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+272)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(v23)+268)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(v23)+248)) = v23 + int32(304)
	if v1649 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1680 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+232)) = uint8(v1680)
	goto L408
L407:
	;
	goto L408
L408:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+52))
	v1687 = F_heap_form_tuple(m, v1682, v23+int32(240), v23+int32(224))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L25
	} else {
		goto L409
	}
L409:
	;
	F_CatalogTupleInsert(m, v1652, v1687)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L25
	} else {
		goto L410
	}
L410:
	;
	F_pfree(m, v1687)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L25
	} else {
		goto L411
	}
L411:
	;
	F_relation_close(m, v1652, int32(3))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L25
	} else {
		goto L412
	}
L412:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStatistics[7]))
	if v1697 != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1699 = int32(0)
	F_RunObjectPostCreateHook(m, int32(3381), v1664, v1699, v1699)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L25
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	F_CacheInvalidateRelcache(m, v66)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L25
	} else {
		goto L417
	}
L416:
	;
	goto L415
L417:
	;
	v1705 = int32(0)
	F_relation_close(m, v66, v1705)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L25
	} else {
		goto L418
	}
L418:
	;
	v1709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = int32(3381)
	if v1709 < v1086 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if v1401 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L420:
	;
	v1720 = v1705
	goto L423
L421:
	;
	goto L422
L422:
	;
	if v1086 != 0 {
		goto L419
	} else {
		goto L427
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(1259)
	v1744 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+int32(288)+v1720<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = v1744
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(97))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L25
	} else {
		goto L425
	}
L425:
	;
	v1754 = v1720 + int32(1)
	if v1754 != v1086 {
		v1720 = v1754
		goto L423
	} else {
		goto L426
	}
L426:
	;
	goto L419
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(1259)
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(97))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L25
	} else {
		goto L428
	}
L428:
	;
	goto L419
L429:
	;
	F_recordDependencyOnSingleRelExpr(m, v23+int32(200), v1082, v126, int32(97), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L25
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+220)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+216)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = int32(2615)
	F_recordDependencyOn(m, v23+int32(200), v23+int32(212), int32(110))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L25
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	F_recordDependencyOnOwner(m, int32(3381), v1664, v26)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L25
	} else {
		goto L434
	}
L434:
	;
	v1812 = v23 + int32(200)
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1813 == int32(0) {
		v1824 = v1812
		goto L3
	} else {
		goto L435
	}
L435:
	;
	F_CreateComments(m, v1664, int32(3381), int32(0), v1813)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L25
	} else {
		goto L436
	}
L436:
	;
	v1824 = v1812
	goto L3
L437:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L25
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_27), int32(0))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L25
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(439), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L25
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L25
	} else {
		goto L442
	}
L442:
	;
	F_errmsg(m, int32(_a_F_CreateStatistics_28), int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L25
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_CreateStatistics_7), int32(457), int32(_a_F_CreateStatistics_8))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L25
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(277)
		return v3
	}
}
func F_create_final_distinct_paths(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 float64
	_ = v111
	var v124 int32
	_ = v124
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
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v31 = F_get_sortgrouplist_exprs(m, v29, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v40 = v28
	goto L1
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
	if v23 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)))
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+318)))
	if v25 != int32(1) {
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
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v36 = int32(0)
	v38 = F_estimate_num_groups(m, l0, v31, v35, v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = v38
	goto L1
L11:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v331 != 0 {
		goto L113
	} else {
		goto L114
	}
L12:
	;
	if v86 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L13:
	;
	v86 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v50 <= int32(0) {
		v78 = int32(1)
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v86 = v78
	goto L12
L17:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = v50
	goto L20
L19:
	;
	v56 = v53
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = int32(0)
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = int32(0)
	v70 = base.B2i32(v68 != v69)
	if v68 == v69 {
		v78 = v70
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v78 = v70
	goto L16
L23:
	;
	v74 = v59 + int32(1)
	if v74 != v56 {
		v59 = v74
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v90 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v103 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L27:
	;
	if v89 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v102 = v89
	goto L26
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v94 = v93
	goto L32
L31:
	;
	v94 = v4
	goto L32
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = v96
	goto L35
L34:
	;
	v97 = v4
	goto L35
L35:
	;
	if v94 < v97 {
		v102 = v95
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v106 <= int32(0) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	if v89 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v111 = float64(-1)
	goto L41
L40:
	;
	v111 = float64(1)
	goto L41
L41:
	;
	v124 = v4
	goto L42
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v124<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+64))
	v133 = F_get_useful_pathkeys_for_distinct(m, l0, v102, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L11
L44:
	;
	v313 = v124 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v313 < v314 {
		v124 = v313
		goto L42
	} else {
		goto L111
	}
L45:
	;
	if v133 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v138 <= v137 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v142 = v137
	goto L48
L48:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v142<<(uint(int32(2))%32))))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v131)+64))
	v163 = v18 + int32(12)
	if v160 == v161 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L44
L50:
	;
	v294 = v142 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v294 < v295 {
		v142 = v294
		goto L48
	} else {
		goto L110
	}
L51:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v264 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L52:
	;
	if v241 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v229
	v241 = int32(1)
	goto L52
L54:
	;
	if v160 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v160 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	v241 = int32(1)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0)
	v241 = int32(1)
	goto L52
L59:
	;
	goto L60
L60:
	;
	if v161 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v181
	v241 = v181
	goto L52
L62:
	;
	goto L63
L63:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v185 = int32(0)
	if v185 < v184 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v188 = v184
	goto L66
L65:
	;
	v188 = v185
	goto L66
L66:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v193 = int32(0)
	goto L67
L67:
	;
	if v193 < v189 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v205 = v201 + v193<<(uint(int32(2))%32)
	goto L71
L70:
	;
	v205 = int32(0)
	goto L71
L71:
	;
	if v193 == v188 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v188
	v241 = base.B2i32(v205 == int32(0))
	goto L52
L73:
	;
	goto L74
L74:
	;
	v211 = base.B2i32(v205 == int32(0))
	if v205 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = v211
	goto L52
L76:
	;
	goto L77
L77:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v218 = v215 + v193<<(uint(int32(2))%32)
	if v218 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = v211
	goto L52
L79:
	;
	goto L80
L80:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if v222 != v223 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v193
	v241 = int32(0)
	goto L52
L82:
	;
	v193 = v193 + int32(1)
	goto L67
L84:
	;
	v262 = v131
	goto L51
L85:
	;
	goto L86
L86:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v131 != v20 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v258 = F_create_incremental_sort_path(m, l0, l2, v131, v160, v242, v111)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L99
	}
L88:
	;
	if v242 == int32(0) {
		goto L50
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v242 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[0])))
	if v247 == int32(0) {
		goto L50
	} else {
		goto L92
	}
L92:
	;
	goto L87
L93:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[0])))
	if v251&int32(1) != 0 {
		goto L87
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v254 = F_create_sort_path(m, l2, v131, v160, v111)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	if v254 == int32(0) {
		goto L50
	} else {
		goto L98
	}
L98:
	;
	v262 = v254
	goto L51
L99:
	;
	if v258 == int32(0) {
		goto L50
	} else {
		goto L100
	}
L100:
	;
	v262 = v258
	goto L51
L101:
	;
	v267 = int32(0)
	v273 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v287 = F_create_upper_unique_path(m, l2, v262, v286, v40)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L108
	}
L104:
	;
	v275 = int32(0)
	v277 = F_makeConst(m, int32(20), int32(-1), v267, int32(8), v273, v275, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v282 = F_create_limit_path(m, l2, v262, v267, v277, int32(0), int64(0), int64(1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_add_path(m, l2, v282)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	goto L50
L108:
	;
	F_add_path(m, l2, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L50
L110:
	;
	goto L49
L111:
	;
	goto L43
L112:
	;
	m.G0 = v18 + int32(16)
	return l2
L113:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v332 != 0 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v338 = int32(0)
	if v337 == v338 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_final_distinct_paths[1])))
	if v334 != int32(1) {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if v375 == int32(0) {
		goto L112
	} else {
		goto L131
	}
L119:
	;
	v375 = int32(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v345 <= int32(0) {
		v370 = int32(1)
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v375 = v370
	goto L118
L123:
	;
	v348 = int32(0)
	if v348 < v345 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v351 = v345
	goto L126
L125:
	;
	v351 = v348
	goto L126
L126:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v355 = v338
	goto L127
L127:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v352+v355<<(uint(int32(2))%32))))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+18)))
	if v361 != int32(1) {
		v370 = v361
		goto L122
	} else {
		goto L129
	}
L128:
	;
	v370 = v361
	goto L122
L129:
	;
	v365 = v355 + int32(1)
	if v365 != v351 {
		v355 = v365
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v380 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v384 = F_create_agg_path(m, l0, l2, v20, v378, int32(2), v380, v381, v380, v380, v40)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	F_add_path(m, l2, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	goto L112
}
func F_create_gating_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v5 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v13 != int32(331) {
		v20 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 == int32(0) {
		v79 = v5
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v16 != 0 {
		v20 = l2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = l2
	goto L6
L5:
	;
	v19 = int32(0)
	goto L6
L6:
	;
	v20 = v19
	goto L1
L7:
	;
	v86 = F_palloc0(m, int32(80))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L25
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v26 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v79 = v5
	goto L7
L10:
	;
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v36 = int32(1)
	v38 = v5
	v40 = v5
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v79 = v67
	goto L7
L14:
	;
	v50 = F_replace_nestloop_params_mutator(m, v48, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v54 = v48
	goto L16
L16:
	;
	v56 = int32(0)
	v58 = F_makeTargetEntry(m, v54, base.I32_extend16_s(v36), v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	v54 = v50
	goto L16
L19:
	;
	if v29 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v29-int32(4)+v36<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v63
	goto L22
L21:
	;
	goto L22
L22:
	;
	v67 = F_lappend(m, v38, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v70 = v40 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v70 < v71 {
		v36 = v36 + int32(1)
		v38 = v67
		v40 = v70
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+72)) = l3
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+56)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+44)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(331)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v97
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+8)) = v99
	v101 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+16)) = v101
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v86)+24)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+36)) = uint8(v89)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v105
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+37)) = uint8(v109)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+37)) = uint8(v111)
	return v86
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 float64
	_ = v169
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v179 float64
	_ = v179
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 float64
	_ = v201
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
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
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(310)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(365)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)) = uint8(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v32
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v37 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v41 = v40
	goto L5
L4:
	;
	v41 = v33
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
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
	v231 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = base.F64_add(v231, v232)
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	v236 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v235, base.F64_add(base.F64_mul(v236, v237), v239))
	m.G0 = v18 + int32(144)
	return v21
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v83 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(1)
	v98 = v33
	v99 = v86
	v100 = v86
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v98<<(uint(int32(2))%32))))
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
	if v100&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v213 = v98 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v213 < v214 {
		v98 = v213
		v99 = v209
		v100 = int32(0)
		goto L27
	} else {
		goto L47
	}
L33:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	F_cost_agg(m, v21, l0, v71, l6, v113, v116, l3, v117, v118, v119, v120, base.F64_convert_i32_s(v122))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	if (v128|v99)&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v209 = v126 & v99
	goto L32
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v197 + v198
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v21)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = base.F64_add(v201, v202)
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = base.F64_add(v205, v206)
	v209 = v196
	goto L32
L38:
	;
	v133 = int32(1)
	if v128&v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v149 = int32(0)
	v151 = v18 + int32(72)
	v153 = float64(0)
	v154 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_create_groupingsets_path[0]))
	v162 = m.G0
	v163 = int32(16)
	v164 = v162 - v163
	m.G0 = v164
	F_cost_tuplesort(m, v164+int32(8), v164, v154, v156, v153, v159, float64(-1))
	mBase = m.M
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v164)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v151)+32)) = v154
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_groupingsets_path[1])))
	v173 = base.F64_add(v153, v169)
	*(*float64)(unsafe.Add(mBase, uint32(v151)+48)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v151)+40)) = v149 + (v172 ^ int32(1))
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v164)))
	*(*float64)(unsafe.Add(mBase, uint32(v151)+56)) = base.F64_add(v173, v179)
	m.G0 = v164 + v163
	goto L45
L41:
	;
	v136 = int32(2)
	goto L43
L42:
	;
	v136 = v133
	goto L43
L43:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v139 = float64(0)
	v141 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	F_cost_agg(m, v18, l0, v136, l6, v113, v137, l3, int32(0), v139, v139, v141, base.F64_convert_i32_s(v143))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+25)))
	v196 = v147 & v99
	goto L37
L45:
	;
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v18)+120))
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v18)+128))
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v18)+104))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+32))
	F_cost_agg(m, v18, l0, int32(1), l6, v113, v186, l3, v187, v188, v189, v190, base.F64_convert_i32_s(v192))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v196 = v149
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v154 int32
	_ = v154
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v240 float64
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v253 int32
	_ = v253
	var v256 float64
	_ = v256
	var v258 int32
	_ = v258
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v270 float64
	_ = v270
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v281 float64
	_ = v281
	var v285 float64
	_ = v285
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v306 float64
	_ = v306
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v459 float64
	_ = v459
	var v463 float64
	_ = v463
	var v493 int32
	_ = v493
	var v494 float64
	_ = v494
	var v497 float64
	_ = v497
	var v501 float64
	_ = v501
	var v503 float64
	_ = v503
	var v506 float64
	_ = v506
	var v533 float64
	_ = v533
	var v535 float64
	_ = v535
	var v539 int32
	_ = v539
	var v540 int64
	_ = v540
	var v545 float64
	_ = v545
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 float64
	_ = v592
	var v593 float64
	_ = v593
	var v611 float64
	_ = v611
	var v619 float64
	_ = v619
	var v621 float64
	_ = v621
	var v622 int32
	_ = v622
	var v623 float64
	_ = v623
	var v625 float64
	_ = v625
	var v626 float64
	_ = v626
	var v628 float64
	_ = v628
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
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(298)
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+228))
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
	if v41 == v46 {
		v87 = v46
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v87 != 0 {
		goto L23
	} else {
		goto L24
	}
L10:
	;
	goto L9
L11:
	;
	if v45 == int32(0) {
		v87 = v46
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v55 < v56 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v55
	goto L15
L14:
	;
	v58 = v56
	goto L15
L15:
	;
	if v58 <= int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v61 = int32(1)
	goto L18
L17:
	;
	v61 = v58
	goto L18
L18:
	;
	v62 = int32(8)
	v67 = int32(0)
	goto L19
L19:
	;
	v74 = v67 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v45+v62+v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v41+v62))))
	v79 = v76 & v78
	v81 = base.B2i32(v79 != int32(0))
	if v79 != 0 {
		v87 = v81
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v87 = v81
	goto L10
L21:
	;
	v83 = v67 + int32(1)
	if v83 != v61 {
		v67 = v83
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v91 = F_get_param_path_clause_serials(m, l6)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(356)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v196
	v198 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v202 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v199, l9, v28+int32(12))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L41
	}
L26:
	;
	if l7 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v154
	goto L25
L28:
	;
	v111 = v93
	v112 = int32(0)
	goto L33
L29:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v93 < v94 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v154 = int32(0)
	goto L27
L32:
	;
	goto L31
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v111<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+56))
	v131 = F_bms_is_member(m, v130, v91)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v154 = v137
	goto L27
L35:
	;
	if v131 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v135 = F_lappend(m, v112, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v137 = v112
	goto L38
L38:
	;
	v139 = v111 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v139 < v140 {
		v111 = v139
		v112 = v137
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v137 = v135
	goto L38
L40:
	;
	goto L34
L41:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+20)) = uint8(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v202
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v207 != int32(1) {
		v214 = v198
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v216 = v214 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)) = uint8(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v218
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v222)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v226
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	v232 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v238
	v240 = float64(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v244 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v210 != int32(1) {
		v214 = v198
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v214 = v213
	goto L42
L45:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v250)))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if int32(0) < v253 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v250 = v244 + int32(8)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v250 = v247 + int32(16)
	goto L45
L49:
	;
	v256 = base.F64_convert_i32_u(v253)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_nestloop_path[0])))
	if v258 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.F64_le(v237, v240) != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v264 = base.F64_add(base.F64_mul(v256, float64(-0.3)), float64(1))
	if base.F64_gt(v264, float64(0)) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v270 = v256
	goto L54
L54:
	;
	v272 = float64(1e+100)
	v273 = base.F64_div(v251, v270)
	if base.F64_gt(v273, v272) != 0 {
		v285 = v272
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v268 = v264
	goto L57
L56:
	;
	v268 = math.Float64frombits(uint64(0x8000000000000000))
	goto L57
L57:
	;
	v270 = base.F64_add(v268, v256)
	goto L54
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = v285
	goto L51
L59:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v273)&int64(9223372036854775807)) {
		v285 = v272
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v281 = float64(1)
	if base.F64_le(v273, v281) != 0 {
		v285 = v281
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v285 = base.F64_nearest(v273)
	goto L58
L62:
	;
	v290 = float64(1)
	goto L64
L63:
	;
	v290 = v237
	goto L64
L64:
	;
	if base.F64_le(v235, v240) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v292 = float64(1)
	goto L67
L66:
	;
	v292 = v235
	goto L67
L67:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v293&int32(-2) != int32(4) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v540 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v230)+24)) = v540
	*(*int64)(unsafe.Add(mBase, uint32(v230)+16)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = l0
	v545 = float64(0)
	if v539 == int32(0) {
		v611 = v545
		v619 = v545
		goto L126
	} else {
		goto L127
	}
L69:
	;
	v533 = v232
	v535 = base.F64_mul(v292, v290)
	goto L68
L70:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v298 != int32(1) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v301 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v303 = base.F64_nearest(base.F64_mul(v292, v301))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v309 = base.F64_div(float64(2), base.F64_add(v306, float64(1)))
	v310 = base.F64_mul(base.F64_mul(v290, v303), v309)
	v311 = base.F64_sub(v292, v303)
	v312 = *(*float64)(unsafe.Add(mBase, uint32(l3)+40))
	v313 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	if v314 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v493 = base.F64_ge(v311, float64(1))
	if v493 != 0 {
		goto L116
	} else {
		goto L117
	}
L75:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v315 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	switch v320 - int32(341) {
	case 0, 1:
		v327 = l6
		goto L77
	default:
		goto L74
	case 3:
		goto L78
	}
L77:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	if v328 == int32(0) {
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l6)+72))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v324 != int32(280) {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v327 = v323
	goto L77
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v331 <= int32(0) {
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v327)+76))
	v335 = int32(0)
	v340 = v335
	v344 = v335
	goto L83
L82:
	;
	v459 = base.F64_add(base.F64_mul(v313, v309), v232)
	if base.F64_gt(v303, float64(1)) != 0 {
		goto L113
	} else {
		goto L114
	}
L83:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v340<<(uint(int32(2))%32))))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v369 = int32(0)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366)+28))
	v371 = F_bms_is_subset(m, v370, v319)
	mBase = m.M
	if v371 == v369 {
		v382 = v369
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v344 == int32(0) {
		goto L74
	} else {
		goto L112
	}
L85:
	;
	if v382 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+28))
	v375 = F_bms_overlap(m, v368, v374)
	mBase = m.M
	if v375 == int32(0) {
		v382 = v369
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v366)+40))
	v379 = F_bms_overlap(m, v368, v378)
	mBase = m.M
	v382 = v379 ^ int32(1)
	goto L86
L89:
	;
	if v334 != 0 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	goto L91
L91:
	;
	v445 = v340 + int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v445 < v446 {
		v340 = v445
		goto L83
	} else {
		goto L111
	}
L92:
	;
	if v433 == int32(0) {
		goto L74
	} else {
		goto L109
	}
L93:
	;
	goto L92
L94:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v388 <= int32(0) {
		v433 = int32(0)
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v433 = int32(0)
	goto L93
L97:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v366)+60))
	v392 = int32(0)
	if v392 < v388 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v395 = v388
	goto L100
L99:
	;
	v395 = v392
	goto L100
L100:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v399 = int32(0)
	goto L101
L101:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v396+v399<<(uint(int32(2))%32))))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)))
	if v409 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	v420 = v399 + int32(1)
	if v420 != v395 {
		v399 = v420
		goto L101
	} else {
		goto L108
	}
L104:
	;
	v410 = int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v366 == v411 {
		v433 = v410
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if v391 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411)+60))
	if v415 == v391 {
		v433 = v410
		goto L93
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	goto L102
L109:
	;
	v439 = int32(1)
	v441 = v340 + v439
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v441 < v442 {
		v340 = v441
		v344 = v439
		goto L83
	} else {
		goto L110
	}
L110:
	;
	goto L82
L111:
	;
	goto L84
L112:
	;
	goto L82
L113:
	;
	v463 = base.F64_add(base.F64_mul(base.F64_mul(v312, base.F64_add(v303, float64(-1))), v309), v459)
	goto L115
L114:
	;
	v463 = v459
	goto L115
L115:
	;
	v533 = base.F64_add(base.F64_div(base.F64_mul(v312, v311), v290), v463)
	v535 = v310
	goto L68
L116:
	;
	v494 = v303
	goto L118
L117:
	;
	v494 = base.F64_add(v303, float64(-1))
	goto L118
L118:
	;
	v497 = base.F64_add(v232, v313)
	if base.F64_gt(v494, float64(0)) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v501 = base.F64_add(base.F64_mul(base.F64_mul(v312, v494), v309), v497)
	goto L121
L120:
	;
	v501 = v497
	goto L121
L121:
	;
	v503 = base.F64_add(base.F64_mul(v311, v290), v310)
	if v493 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v506 = base.F64_add(v311, float64(-1))
	goto L124
L123:
	;
	v506 = v311
	goto L124
L124:
	;
	if base.F64_gt(v506, float64(0)) == int32(0) {
		v533 = v501
		v535 = v503
		goto L68
	} else {
		goto L125
	}
L125:
	;
	v533 = base.F64_add(base.F64_mul(v506, v312), v501)
	v535 = v503
	goto L68
L126:
	;
	v621 = *(*float64)(unsafe.Add(mBase, _c_F_create_nestloop_path[1]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v623 = *(*float64)(unsafe.Add(mBase, uint32(v622)+24))
	v625 = *(*float64)(unsafe.Add(mBase, uint32(v622)+16))
	v626 = base.F64_add(base.F64_add(v233, v619), v625)
	*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = v626
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v32)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+56)) = base.F64_add(v626, base.F64_add(base.F64_mul(v623, v628), base.F64_add(base.F64_mul(base.F64_add(v611, v621), v535), v533)))
	m.G0 = v230 + int32(32)
	m.G0 = v28 + int32(16)
	return v32
L127:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v550 <= int32(0) {
		v611 = v545
		v619 = float64(0)
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v557 = int32(0)
	goto L129
L129:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v579+v557<<(uint(int32(2))%32))))
	v586 = F_cost_qual_eval_walker(m, v583, v230+int32(8))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	v592 = *(*float64)(unsafe.Add(mBase, uint32(v230)+24))
	v593 = *(*float64)(unsafe.Add(mBase, uint32(v230)+16))
	v611 = v592
	v619 = v593
	goto L126
L131:
	;
	v589 = v557 + int32(1)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v589 < v590 {
		v557 = v589
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v4
	if v18 <= v4 {
		v98 = v4
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
	v26 = l0 + int32(132)
	if base.Ui32(int32(4)) <= base.Ui32(v18) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = v4
	v35 = v4
	v36 = v4
	goto L6
L4:
	;
	v61 = v4
	v62 = v4
	goto L5
L5:
	;
	if v24 == int32(0) {
		v98 = v62
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v44 = v34 + v26
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)))
	v52 = v35 + v45 + v47 + v49 + v51
	v53 = int32(4)
	v54 = v34 + v53
	v56 = v36 + v53
	if v56 != v18&int32(2147483644) {
		v34 = v54
		v35 = v52
		v36 = v56
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v61 = v54
	v62 = v52
	goto L5
L8:
	;
	goto L7
L9:
	;
	v76 = v61
	v77 = v62
	v79 = v4
	goto L10
L10:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v26))))
	v88 = v77 + v87
	v89 = int32(1)
	v92 = v79 + v89
	if v92 != v24 {
		v76 = v76 + v89
		v77 = v88
		v79 = v92
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v98 = v88
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
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v122 = int32(8)
	v128 = v98<<(uint(v122)%32) | int32(base.Ui32(v98&int32(_a_F_create_secmsg_0))>>(uint(v122)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v109+v18)+1)) = uint16(v128)
	v131 = l2 - v108
	v133 = v131 - int32(2)
	if v133 < v122 {
		v307 = int32(-12)
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v109+int32(1), l0+int32(132), v18)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v318 = F___memset(m, v109, int32(0), v108)
	mBase = m.M
	goto L68
L20:
	;
	v136 = F_palloc(m, l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v138 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v138)
	v141 = v136 + int32(1)
	v142 = int32(0)
	v146 = m.G0
	v148 = v146 - int32(16)
	m.G0 = v148
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v142
	v154 = F_open(m, int32(_a_F_create_secmsg_1), v142, v148)
	mBase = m.M
	if v154 != int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v187 != 0 {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v157 = int32(1)
	if v133 == int32(0) {
		v180 = v157
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v187 = v142
	goto L25
L25:
	;
	m.G0 = v148 + int32(16)
	goto L22
L26:
	;
	v182 = F_close(m, v154)
	mBase = m.M
	v187 = v180
	goto L25
L27:
	;
	v160 = v141
	v161 = v133
	goto L28
L28:
	;
	v166 = F_read(m, v154, v160, v161)
	mBase = m.M
	if v166 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v180 = v157
	goto L26
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v170 == int32(27) {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v175 = v161 - v166
	if v175 != 0 {
		v160 = v160 + v166
		v161 = v175
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v180 = int32(0)
	goto L26
L34:
	;
	goto L29
L35:
	;
	v192 = v131 + v136
	v198 = v141
	goto L39
L36:
	;
	goto L37
L37:
	;
	F_pfree(m, v136)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L67
	}
L38:
	;
	v287 = F___memset(m, v136, int32(0), l2)
	mBase = m.M
	goto L66
L39:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v209 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136+v133)+1)) = uint8(v270)
	if v108 != 0 {
		goto L60
	} else {
		goto L61
	}
L41:
	;
	v266 = int32(1)
	goto L43
L42:
	;
	v211 = int32(0)
	v215 = m.G0
	v217 = v215 - int32(16)
	m.G0 = v217
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v211
	v223 = F_open(m, int32(_a_F_create_secmsg_1), v211, v217)
	mBase = m.M
	if v223 != int32(-1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v267 = v266 + v198
	if base.Ui32(v267) < base.Ui32(v192-int32(1)) {
		v198 = v267
		goto L39
	} else {
		goto L58
	}
L44:
	;
	if v256 == int32(0) {
		goto L38
	} else {
		goto L57
	}
L45:
	;
	goto L49
L46:
	;
	v256 = v211
	goto L47
L47:
	;
	m.G0 = v217 + int32(16)
	goto L44
L48:
	;
	v251 = F_close(m, v223)
	mBase = m.M
	v256 = v249
	goto L47
L49:
	;
	v229 = v198
	v230 = int32(1)
	goto L50
L50:
	;
	v235 = F_read(m, v223, v229, v230)
	mBase = m.M
	if v235 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v249 = int32(1)
	goto L48
L52:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_create_secmsg[0]))
	if v239 == int32(27) {
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v244 = v230 - v235
	if v244 != 0 {
		v229 = v229 + v235
		v230 = v244
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v249 = int32(0)
	goto L48
L56:
	;
	goto L51
L57:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v266 = base.B2i32(v263 != int32(0))
	goto L43
L58:
	;
	goto L40
L59:
	;
	v280 = F_pgp_mpi_create(m, v136, l2<<(uint(int32(3))%32)-int32(6), v16+int32(12))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L63
	}
L60:
	;
	v272 = F__emscripten_memcpy_bulkmem(m, v192, v109, v108)
	mBase = m.M
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v283 = F___memset(m, v136, int32(0), l2)
	mBase = m.M
	goto L64
L64:
	;
	F_pfree(m, v136)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v307 = v280
	goto L19
L66:
	;
	goto L37
L67:
	;
	v307 = int32(-17)
	goto L19
L68:
	;
	F_pfree(m, v109)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	if int32(0) <= v307 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
	goto L72
L71:
	;
	goto L72
L72:
	;
	m.G0 = v16 + int32(16)
	return v307
}
