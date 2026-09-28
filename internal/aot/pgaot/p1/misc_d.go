package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_DecodeUnits(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	v3 = int32(_a_F_DecodeUnits_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeUnits[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v78
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeUnits[0])) = v59
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59)+11)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
	v78 = v65
	goto L1
L3:
	;
	v13 = F_strncmp(m, l0, v11, int32(10))
	mBase = m.M
	if v13 == int32(0) {
		v59 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v23 = int32(_a_F_DecodeUnits_1)
	v25 = int32(_a_F_DecodeUnits_2)
	goto L7
L6:
	;
	goto L5
L7:
	;
	v32 = v23 + (v25-v23)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
	v34 = v16 - v33
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v78 = int32(31)
	goto L1
L9:
	;
	v38 = F_strncmp(m, l0, v32, int32(10))
	mBase = m.M
	if v38 == int32(0) {
		v59 = v32
		goto L2
	} else {
		goto L12
	}
L10:
	;
	v41 = v34
	goto L11
L11:
	;
	v45 = base.B2i32(v41 < int32(0))
	if v41 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v41 = v38
	goto L11
L13:
	;
	v46 = v32 - int32(16)
	goto L15
L14:
	;
	v46 = v25
	goto L15
L15:
	;
	if v41 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = v23
	goto L18
L17:
	;
	v49 = v32 + int32(16)
	goto L18
L18:
	;
	if base.Ui32(v49) <= base.Ui32(v46) {
		v23 = v49
		v25 = v46
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
}
func F_DecodingContextFindStartpoint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	F_XLogBeginRead(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v19)
	v22 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v22)
	F_errmsg_internal(m, int32(_a_F_DecodingContextFindStartpoint_0), v8+int32(16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = F_XLogReadRecord(m, v37, v8+int32(28))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(_a_F_DecodingContextFindStartpoint_1), int32(721), int32(_a_F_DecodingContextFindStartpoint_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L11:
	;
	v94 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v48 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	if v48 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_LogicalDecodingProcessRecord(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(2) {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_DecodingContextFindStartpoint[0]))
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_XLogReadRecord(m, v65, v8+int32(28))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v70 == int32(0) {
		v48 = v68
		goto L15
	} else {
		goto L25
	}
L25:
	;
	goto L16
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v82
	F_errmsg_internal(m, int32(_a_F_DecodingContextFindStartpoint_3), v8)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_DecodingContextFindStartpoint_1), int32(732), int32(_a_F_DecodingContextFindStartpoint_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_s_lock(m, v11, int32(_a_F_DecodingContextFindStartpoint_4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+136)))
	if v101 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v105
	goto L35
L34:
	;
	goto L35
L35:
	;
	v107 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11))), uint32(v107))
	m.G0 = v8 + int32(32)
	return
L36:
	;
	F_errmsg_internal(m, int32(_a_F_DecodingContextFindStartpoint_5), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_DecodingContextFindStartpoint_1), int32(734), int32(_a_F_DecodingContextFindStartpoint_2))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DetermineTimeZoneAbbrevOffset(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v348 int64
	_ = v348
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int64
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	v8 = m.G0
	v10 = v8 - int32(288)
	m.G0 = v10
	v13 = v10 + int32(280)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v25 <= int32(-4713) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v178 = v10 + int32(16)
	goto L38
L2:
	;
	m.G0 = v23 + int32(32)
	goto L1
L3:
	;
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(0)
	v173 = v160
	goto L2
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = int32(60)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v56 = base.B2i32(int32(2) < v41)
	if int32(2) < v41 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	if v25 != int32(-4713) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v25 <= int32(_a_F_DetermineTimeZoneAbbrevOffset_0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(10) < v30 {
		v41 = v30
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v41 = v35
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v25 != int32(_a_F_DetermineTimeZoneAbbrevOffset_1) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(5) < v38 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v41 = v38
	goto L4
L15:
	;
	v57 = int32(_a_F_DetermineTimeZoneAbbrevOffset_2)
	goto L17
L16:
	;
	v57 = int32(_a_F_DetermineTimeZoneAbbrevOffset_3)
	goto L17
L17:
	;
	v58 = v57 + v25
	v66 = base.I32_div_u_s(v58, int32(100))
	v69 = base.I32_div_u_s(v58, int32(400))
	if int32(2) < v41 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = int32(1)
	goto L20
L19:
	;
	v73 = int32(13)
	goto L20
L20:
	;
	v78 = base.I32_div_s((v73+v41)*int32(_a_F_DetermineTimeZoneAbbrevOffset_4), int32(256))
	v81 = v52 + v58*int32(365) + int32(base.Ui32(v58)>>(uint(int32(2))%32)) - v66 + v69 + v78 - int32(_a_F_DetermineTimeZoneAbbrevOffset_5)
	v85 = base.I64_extend_i32_s(v42+(v43+v44*v45)*v45) + base.I64_extend_i32_s(v81)*int64(86400)
	if base.B2i32(int32(0) < v81)&base.B2i32(v85 < int64(0)) != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v85 - int64(86400)
	v104 = F_pg_next_dst_boundary(m, v23+int32(24), v23+int32(12), v23+int32(4), v23+int32(16), v23+int32(8), v23, l2)
	mBase = m.M
	if v104 < int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v104 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v85 - base.I64_extend_i32_s(v111)
	v173 = int32(0) - v111
	goto L2
L24:
	;
	goto L25
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v119 = v85 - base.I64_extend_i32_s(v117)
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v23)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v124 = v85 - base.I64_extend_i32_s(v122)
	if base.B2i32(v120 <= v119)|base.B2i32(v120 <= v124) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v119
	v173 = int32(0) - v117
	goto L2
L27:
	;
	goto L28
L28:
	;
	if base.B2i32(v124 < v120)|base.B2i32(v119 <= v120) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v124
	v173 = int32(0) - v122
	goto L2
L30:
	;
	goto L31
L31:
	;
	if v117 < v122 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v145
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v119
	v173 = int32(0) - v117
	goto L2
L33:
	;
	goto L34
L34:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v124
	v173 = int32(0) - v122
	goto L2
L35:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	if v298 != 0 {
		goto L66
	} else {
		goto L67
	}
L36:
	;
	v295 = F_strlen(m, v284)
	mBase = m.M
	goto L35
L38:
	;
	goto L39
L39:
	;
	v185 = int32(255)
	if (v178^l1)&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v288)
	goto L36
L41:
	;
	v269 = v264
	v270 = v265
	v271 = v266
	goto L62
L42:
	;
	if v259 == int32(0) {
		v284 = v257
		v285 = v258
		goto L40
	} else {
		goto L61
	}
L43:
	;
	v257 = l1
	v258 = v178
	v259 = v185
	goto L42
L44:
	;
	goto L45
L45:
	;
	v189 = int32(0)
	if base.B2i32(l1&int32(3) == v189)|int32(0) == v189 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v225 == int32(0) {
		v284 = v222
		v285 = v223
		goto L40
	} else {
		goto L55
	}
L47:
	;
	v201 = l1
	v202 = v178
	v203 = v185
	goto L50
L48:
	;
	goto L49
L49:
	;
	v222 = l1
	v223 = v178
	v224 = v185
	v225 = int32(1)
	goto L46
L50:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v205)
	if v205 == int32(0) {
		v264 = v201
		v265 = v202
		v266 = v203
		goto L41
	} else {
		goto L52
	}
L51:
	;
	v222 = v216
	v223 = v210
	v224 = v212
	v225 = v214
	goto L46
L52:
	;
	v209 = int32(1)
	v210 = v202 + v209
	v212 = v203 - v209
	v213 = int32(0)
	v214 = base.B2i32(v212 != v213)
	v216 = v201 + v209
	if v216&int32(3) == v213 {
		v222 = v216
		v223 = v210
		v224 = v212
		v225 = v214
		goto L46
	} else {
		goto L53
	}
L53:
	;
	if v212 != 0 {
		v201 = v216
		v202 = v210
		v203 = v212
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if base.B2i32(v228 == int32(0))|base.B2i32(base.Ui32(v224) < base.Ui32(int32(4))) != 0 {
		v257 = v222
		v258 = v223
		v259 = v224
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v235 = v222
	v236 = v223
	v237 = v224
	goto L57
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v243 = int32(-2139062144)
	if (int32(16843008)-v240|v240)&v243 != v243 {
		v264 = v235
		v265 = v236
		v266 = v237
		goto L41
	} else {
		goto L59
	}
L58:
	;
	v257 = v251
	v258 = v249
	v259 = v253
	goto L42
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v240
	v248 = int32(4)
	v249 = v236 + v248
	v251 = v235 + v248
	v253 = v237 - v248
	if base.Ui32(int32(3)) < base.Ui32(v253) {
		v235 = v251
		v236 = v249
		v237 = v253
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v264 = v257
	v265 = v258
	v266 = v259
	goto L41
L62:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v273)
	if v273 == int32(0) {
		v284 = v269
		v285 = v270
		goto L40
	} else {
		goto L64
	}
L63:
	;
	v284 = v280
	v285 = v278
	goto L40
L64:
	;
	v277 = int32(1)
	v278 = v270 + v277
	v280 = v269 + v277
	v282 = v271 - v277
	if v282 != 0 {
		v269 = v280
		v270 = v278
		v271 = v282
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v300 = v178
	v304 = v298
	goto L69
L67:
	;
	goto L68
L68:
	;
	v336 = int32(0)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	if v343 <= v336 {
		v496 = v336
		goto L77
	} else {
		goto L78
	}
L69:
	;
	if base.Ui32((v304-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L68
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v316)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v318 != 0 {
		v300 = v300 + int32(1)
		v304 = v318
		goto L69
	} else {
		goto L75
	}
L72:
	;
	v314 = v304 - int32(32)
	goto L74
L73:
	;
	v314 = v304
	goto L74
L74:
	;
	v316 = v314 & int32(255)
	goto L71
L75:
	;
	goto L70
L76:
	;
	if v496 != 0 {
		goto L111
	} else {
		goto L112
	}
L77:
	;
	goto L76
L78:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(280))))
	v351 = int32(0)
	goto L79
L79:
	;
	v362 = v351 + (l2 + int32(_a_F_DetermineTimeZoneAbbrevOffset_6))
	v363 = F_strcmp(m, v10+int32(16), v362)
	mBase = m.M
	if v363 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l2)+260))
	if v369 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v364 = F_strlen(m, v362)
	mBase = m.M
	v367 = v364 + v351 + int32(1)
	if v367 < v343 {
		v351 = v367
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v496 = v336
	goto L77
L85:
	;
	v414 = l2 + int32(_a_F_DetermineTimeZoneAbbrevOffset_7)
	v416 = l2 + int32(_a_F_DetermineTimeZoneAbbrevOffset_8)
	v417 = v406
	goto L99
L86:
	;
	v406 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v376 = v369
	v381 = int32(0)
	goto L89
L89:
	;
	v389 = int32(1)
	v390 = (v376 + v381) >> (uint(v389) % 32)
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(280)+v390<<(uint(int32(3))%32))))
	v397 = base.B2i32(v348 < v396)
	if v348 < v396 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v406 = v398
	goto L85
L91:
	;
	v398 = v381
	goto L93
L92:
	;
	v398 = v390 + v389
	goto L93
L93:
	;
	if v348 < v396 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v399 = v390
	goto L96
L95:
	;
	v399 = v376
	goto L96
L96:
	;
	if v398 < v399 {
		v376 = v399
		v381 = v398
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v481
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v483
	v496 = int32(1)
	goto L77
L99:
	;
	if int32(0) < v417 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_DetermineTimeZoneAbbrevOffset[0])))
	v443 = v416 + v440<<(uint(int32(4))%32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	if v444 == v351 {
		v475 = v443
		goto L98
	} else {
		goto L105
	}
L101:
	;
	v432 = v417 - int32(1)
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414+v432))))
	v437 = v416 + v434<<(uint(int32(4))%32)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	if v438 != v351 {
		v417 = v432
		goto L99
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	v475 = v437
	goto L98
L105:
	;
	if v369 <= v406 {
		v496 = v336
		goto L77
	} else {
		goto L106
	}
L106:
	;
	v452 = v406
	goto L107
L107:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+v414))))
	v463 = v416 + v460<<(uint(int32(4))%32)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	if v464 == v351 {
		v475 = v463
		goto L98
	} else {
		goto L109
	}
L108:
	;
	v496 = v336
	goto L77
L109:
	;
	v467 = v452 + int32(1)
	if v369 != v467 {
		v452 = v467
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v499
	v504 = int32(0) - v498
	goto L113
L112:
	;
	v504 = v173
	goto L113
L113:
	;
	m.G0 = v10 + int32(288)
	return v504
}
func F_Do_MultiXactIdWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(1)
	if l2&int32(_a_F_Do_MultiXactIdWait_0) == int32(_a_F_Do_MultiXactIdWait_1) {
		v264 = v18
		v266 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l7 != 0 {
		goto L68
	} else {
		goto L69
	}
L2:
	;
	v34 = F_GetMultiXactIdMembers(m, l0, v16+int32(12), int32(base.Ui32(l2&int32(128))>>(uint(int32(7))%32))|base.B2i32(l2&int32(_a_F_Do_MultiXactIdWait_2) == int32(64)))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v34 < int32(0) {
		v264 = v18
		v266 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v34 == int32(0) {
		v248 = v18
		v250 = v10
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_pfree(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L67
	}
L7:
	;
	v49 = int32(0)
	v58 = v10
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v61 = int32(3)
	v63 = v60 + v49<<(uint(v61)%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(v65) < base.Ui32(v61) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v248 = v235
	v250 = v234
	goto L6
L10:
	;
	v235 = int32(1)
	v237 = v49 + v235
	if v237 != v34 {
		v49 = v237
		v58 = v234
		goto L8
	} else {
		goto L66
	}
L11:
	;
	if v197 != 0 {
		goto L51
	} else {
		goto L52
	}
L12:
	;
	v197 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_Do_MultiXactIdWait[0]))
	if v77 == v65 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v197 = int32(1)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_Do_MultiXactIdWait[1]))
	if v81 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v197 = v187
	goto L11
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_Do_MultiXactIdWait[2]))
	if v85 == int32(0) {
		v187 = int32(0)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_Do_MultiXactIdWait[3]))
	v157 = int32(0)
	v160 = v81 - int32(1)
	goto L41
L22:
	;
	v90 = v85
	goto L23
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v96 == int32(4) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v187 = int32(0)
	goto L18
L25:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v90)+80))
	if v150 != 0 {
		v90 = v150
		goto L23
	} else {
		goto L40
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v99 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v102 = int32(1)
	if v65 == v99 {
		v187 = v102
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
	v106 = v104 - int32(1)
	if v106 < int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v90)+48))
	v112 = int32(0)
	v115 = v106
	goto L30
L30:
	;
	v120 = int32(2)
	v121 = base.I32_div_s(v115-v112, v120)
	v122 = v121 + v112
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v109+v122<<(uint(v120)%32))))
	if v126 == v65 {
		v187 = v102
		goto L18
	} else {
		goto L32
	}
L31:
	;
	goto L25
L32:
	;
	v135 = base.B2i32(v126-v65 < int32(0)) | base.B2i32(base.Ui32(v126) < base.Ui32(int32(3)))
	if v135 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v136 = v122 + int32(1)
	goto L35
L34:
	;
	v136 = v112
	goto L35
L35:
	;
	if v135 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v139 = v115
	goto L38
L37:
	;
	v139 = v122 - int32(1)
	goto L38
L38:
	;
	if v136 <= v139 {
		v112 = v136
		v115 = v139
		goto L30
	} else {
		goto L39
	}
L39:
	;
	goto L31
L40:
	;
	goto L24
L41:
	;
	v165 = int32(2)
	v166 = base.I32_div_s(v160-v157, v165)
	v167 = v166 + v157
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v155+v167<<(uint(v165)%32))))
	v172 = base.B2i32(v171 == v65)
	if v171 == v65 {
		v187 = v172
		goto L18
	} else {
		goto L43
	}
L42:
	;
	v187 = v172
	goto L18
L43:
	;
	v175 = base.B2i32(base.Ui32(v171) < base.Ui32(v65))
	if base.Ui32(v171) < base.Ui32(v65) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v176 = v167 + int32(1)
	goto L46
L45:
	;
	v176 = v157
	goto L46
L46:
	;
	if base.Ui32(v171) < base.Ui32(v65) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v179 = v160
	goto L49
L48:
	;
	v179 = v167 - int32(1)
	goto L49
L49:
	;
	if v176 <= v179 {
		v157 = v176
		v160 = v179
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v234 = v58 + int32(1)
	goto L10
L52:
	;
	goto L53
L53:
	;
	v200 = int32(2)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v200)%32))+uint32(_c_F_Do_MultiXactIdWait[4])))
	v203 = int32(12)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202*v203)+uint32(_c_F_Do_MultiXactIdWait[5])))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_Do_MultiXactIdWait[4])))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208*v203)+uint32(_c_F_Do_MultiXactIdWait[5])))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v207<<(uint(v200)%32))+uint32(_c_F_Do_MultiXactIdWait[6])))
	goto L54
L54:
	;
	if int32(base.Ui32(v218)>>(uint(v213)%32))&int32(1) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if l7 == int32(0) {
		v234 = v58
		goto L10
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if l3 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v226 = F_TransactionIdIsInProgress(m, v65)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v234 = v226 + v58
	goto L10
L60:
	;
	v229 = F_ConditionalXactLockTableWait(m, v65, l8)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_XactLockTableWait(m, v65, l4, l5, l6)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L3
	} else {
		goto L65
	}
L63:
	;
	if v229 != 0 {
		v234 = v58
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v248 = int32(0)
	v250 = v58
	goto L6
L65:
	;
	v234 = v58
	goto L10
L66:
	;
	goto L9
L67:
	;
	v264 = v248
	v266 = v250
	goto L1
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v266
	goto L70
L69:
	;
	goto L70
L70:
	;
	m.G0 = v16 + int32(16)
	return v264
}
func F_DynaHashAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_MemoryContextAllocExtended(m, l1, l0, int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F___divti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int64(63)
	v14 = l2 >> (uint(v13) % 64)
	v15 = l1 ^ v14
	v18 = v15 + int64(base.Ui64(l2)>>(uint(v13)%64))
	v24 = l4 >> (uint(v13) % 64)
	v25 = v24 ^ l3
	v28 = v25 + int64(base.Ui64(l4)>>(uint(v13)%64))
	F___udivmodti4(m, v11, v18, base.I64_extend_i32_u(base.B2i32(base.Ui64(v18) < base.Ui64(v15)))+(v14^l2), v28, base.I64_extend_i32_u(base.B2i32(base.Ui64(v28) < base.Ui64(v25)))+(v24^l4))
	mBase = m.M
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v35 = v14 ^ v24
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v37 = v35 ^ v36
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37 - v35
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v35 ^ v34 - v35 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v37) < base.Ui64(v35)))
	m.G0 = v11 + v10
	return
}
func F_daitch_mokotoff(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v119 int32
	_ = v119
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v283 int32
	_ = v283
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v443 int32
	_ = v443
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v533 int32
	_ = v533
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v700 int32
	_ = v700
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v857 int32
	_ = v857
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1020 int32
	_ = v1020
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1067 int64
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1116 int64
	_ = v1116
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(_a_F_daitch_mokotoff_0), int32(0), int32(_a_F_daitch_mokotoff_1), int32(_a_F_daitch_mokotoff_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = int32(_a_F_daitch_mokotoff_3)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0])) = v35
	v41 = F_text_to_cstring(m, v25)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v43 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v74 = F_pg_server_to_any(m, v41, v72, int32(6))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v49 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v60 = int32(1)
	if v43&v60 != 0 {
		v72 = int32(base.Ui32(v43)>>(uint(v60)%32)) - v60
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v52 = int32(16)
	goto L11
L10:
	;
	v52 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = int32(4)
	goto L14
L13:
	;
	v59 = v52
	goto L14
L14:
	;
	v72 = v59
	goto L5
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v72 = int32(base.Ui32(v66)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v78 = F_initArrayResult(m, int32(25), v35, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
	v83 = v22 + int32(20)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v94 = v88
	goto L22
L18:
	;
	m.G0 = v22 + int32(32)
	return v1116
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0])) = v38
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L256
	}
L20:
	;
	if v212 == int32(0) {
		goto L19
	} else {
		goto L52
	}
L21:
	;
	v212 = base.I32_extend8_s(v205)
	goto L20
L22:
	;
	v95 = v74 + v94
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
	v98 = v96 & int32(255)
	if int32(0) <= v96 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v205 = int32(92)
	goto L21
L24:
	;
	goto L23
L25:
	;
	if v187&int32(255) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L26:
	;
	v158 = F_pg_utf_mblen_private(m, v95)
	mBase = m.M
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v160 = v158 + v159
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v160
	v162 = int32(26)
	if base.Ui32(v157-int32(91)) < base.Ui32(int32(3)) {
		v187 = v162
		v190 = v160
		goto L25
	} else {
		goto L39
	}
L27:
	;
	v155 = v98
	goto L29
L28:
	;
	if v98&int32(224) == int32(192) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v155 != 0 {
		v157 = v155
		goto L26
	} else {
		goto L38
	}
L30:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v95))))
	v155 = v148 | v150&int32(63)
	goto L29
L31:
	;
	v147 = int32(1)
	v148 = v98 << (uint(int32(6)) % 32) & int32(1984)
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v98&int32(240) == int32(224) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	v147 = int32(2)
	v148 = v98<<(uint(int32(12))%32)&int32(_a_F_daitch_mokotoff_4) | v119&int32(63)<<(uint(int32(6))%32)
	goto L30
L35:
	;
	goto L36
L36:
	;
	if v98&int32(248) != int32(240) {
		v157 = int32(-1)
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	v136 = int32(63)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+2)))
	v147 = int32(3)
	v148 = v98<<(uint(int32(18))%32)&int32(_a_F_daitch_mokotoff_5) | v135&v136<<(uint(int32(12))%32) | v141&v136<<(uint(int32(6))%32)
	goto L30
L38:
	;
	v187 = int32(0)
	v190 = v94
	goto L25
L39:
	;
	if base.Ui32(v157) <= base.Ui32(int32(95)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v187 = v157
	v190 = v160
	goto L25
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(v157) <= base.Ui32(int32(255)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_daitch_mokotoff[1]))))
	v187 = v173
	v190 = v160
	goto L25
L44:
	;
	goto L45
L45:
	;
	switch v157 - int32(260) {
	case 0, 1:
		v205 = int32(91)
		goto L21
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		v187 = v162
		v190 = v160
		goto L25
	case 20, 21:
		goto L24
	default:
		goto L46
	}
L46:
	;
	v180 = int32(2)
	if base.B2i32(base.Ui32(v157-int32(354)) < base.Ui32(v180))|base.B2i32(base.Ui32(v157-int32(538)) < base.Ui32(v180)) != 0 {
		v205 = int32(93)
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v187 = v162
	v190 = v160
	goto L25
L48:
	;
	v212 = base.I32_extend8_s(v187)
	goto L20
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(int32(28)) < base.Ui32((v187-int32(65))&int32(255)) {
		v94 = v190
		goto L22
	} else {
		goto L51
	}
L51:
	;
	v205 = v187
	goto L21
L52:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v215
	v220 = v212*int32(12) + int32(_a_F_daitch_mokotoff_6)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220-int32(772))))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v220-int32(776))))
	if v226 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v229 = v215
	v231 = v223
	v233 = v226
	goto L56
L54:
	;
	v443 = v223
	goto L55
L55:
	;
	if v443 == int32(0) {
		goto L19
	} else {
		goto L106
	}
L56:
	;
	v247 = v22 + int32(24)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v258 = v252
	goto L61
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v421
	v443 = v423
	goto L55
L58:
	;
	goto L57
L59:
	;
	if v376 == int32(0) {
		v421 = v229
		v423 = v231
		goto L58
	} else {
		goto L91
	}
L60:
	;
	v376 = base.I32_extend8_s(v369)
	goto L59
L61:
	;
	v259 = v74 + v258
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v259))))
	v262 = v260 & int32(255)
	if int32(0) <= v260 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v369 = int32(92)
	goto L60
L63:
	;
	goto L62
L64:
	;
	if v351&int32(255) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L65:
	;
	v322 = F_pg_utf_mblen_private(m, v259)
	mBase = m.M
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v324 = v322 + v323
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v324
	v326 = int32(26)
	if base.Ui32(v321-int32(91)) < base.Ui32(int32(3)) {
		v351 = v326
		v354 = v324
		goto L64
	} else {
		goto L78
	}
L66:
	;
	v319 = v262
	goto L68
L67:
	;
	if v262&int32(224) == int32(192) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v319 != 0 {
		v321 = v319
		goto L65
	} else {
		goto L77
	}
L69:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311+v259))))
	v319 = v312 | v314&int32(63)
	goto L68
L70:
	;
	v311 = int32(1)
	v312 = v262 << (uint(int32(6)) % 32) & int32(1984)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v262&int32(240) == int32(224) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v311 = int32(2)
	v312 = v262<<(uint(int32(12))%32)&int32(_a_F_daitch_mokotoff_4) | v283&int32(63)<<(uint(int32(6))%32)
	goto L69
L74:
	;
	goto L75
L75:
	;
	if v262&int32(248) != int32(240) {
		v321 = int32(-1)
		goto L65
	} else {
		goto L76
	}
L76:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v300 = int32(63)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+2)))
	v311 = int32(3)
	v312 = v262<<(uint(int32(18))%32)&int32(_a_F_daitch_mokotoff_5) | v299&v300<<(uint(int32(12))%32) | v305&v300<<(uint(int32(6))%32)
	goto L69
L77:
	;
	v351 = int32(0)
	v354 = v258
	goto L64
L78:
	;
	if base.Ui32(v321) <= base.Ui32(int32(95)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v351 = v321
	v354 = v324
	goto L64
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(v321) <= base.Ui32(int32(255)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+uint32(_c_F_daitch_mokotoff[1]))))
	v351 = v337
	v354 = v324
	goto L64
L83:
	;
	goto L84
L84:
	;
	switch v321 - int32(260) {
	case 0, 1:
		v369 = int32(91)
		goto L60
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		v351 = v326
		v354 = v324
		goto L64
	case 20, 21:
		goto L63
	default:
		goto L85
	}
L85:
	;
	v344 = int32(2)
	if base.B2i32(base.Ui32(v321-int32(354)) < base.Ui32(v344))|base.B2i32(base.Ui32(v321-int32(538)) < base.Ui32(v344)) != 0 {
		v369 = int32(93)
		goto L60
	} else {
		goto L86
	}
L86:
	;
	v351 = v326
	v354 = v324
	goto L64
L87:
	;
	v376 = base.I32_extend8_s(v351)
	goto L59
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(int32(28)) < base.Ui32((v351-int32(65))&int32(255)) {
		v258 = v354
		goto L61
	} else {
		goto L90
	}
L90:
	;
	v369 = v351
	goto L60
L91:
	;
	v379 = int32(0)
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v380 == v379 {
		v421 = v229
		v423 = v231
		goto L58
	} else {
		goto L92
	}
L92:
	;
	v388 = v379
	v391 = v380
	goto L93
L93:
	;
	if v376 != v391&int32(255) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v414 = v233 + v388*int32(12)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+8))
	if v415 != 0 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v406 = v388 + int32(1)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+v406*int32(12)))))
	if v410 != 0 {
		v388 = v406
		v391 = v410
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	v421 = v229
	v423 = v231
	goto L58
L99:
	;
	v416 = v411
	goto L101
L100:
	;
	v416 = v229
	goto L101
L101:
	;
	if v415 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v417 = v415
	goto L104
L103:
	;
	v417 = v231
	goto L104
L104:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v418 != 0 {
		v229 = v416
		v231 = v417
		v233 = v418
		goto L56
	} else {
		goto L105
	}
L105:
	;
	v421 = v416
	v423 = v417
	goto L58
L106:
	;
	v461 = F_palloc(m, int32(84))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v461
	base.MemoryCopy(m, v461, int32(_a_F_daitch_mokotoff_7), int32(84))
	v467 = int32(0)
	v469 = v467
	v473 = v443
	v477 = v467
	goto L109
L108:
	;
	v1067 = F_makeArrayResult(m, v78, v38)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L254
	}
L109:
	;
	v489 = v469 << (uint(int32(2)) % 32)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v489+(v22+int32(12)))))
	if v493 == int32(0) {
		goto L108
	} else {
		goto L111
	}
L110:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	if v1010 == int32(0) {
		goto L108
	} else {
		goto L248
	}
L111:
	;
	v497 = v22 + int32(20)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v508 = v502
	goto L115
L112:
	;
	v874 = v469 ^ int32(1)
	v876 = v874 << (uint(int32(2)) % 32)
	v879 = v876 + (v22 + int32(12))
	v880 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v879))) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(24)+v876))) = v880
	v888 = v473 + int32(9)
	if v857 != 0 {
		goto L199
	} else {
		goto L200
	}
L113:
	;
	if v626 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L114:
	;
	v626 = base.I32_extend8_s(v619)
	goto L113
L115:
	;
	v509 = v74 + v508
	v510 = int32(*(*int8)(unsafe.Add(mBase, uint32(v509))))
	v512 = v510 & int32(255)
	if int32(0) <= v510 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	v619 = int32(92)
	goto L114
L117:
	;
	goto L116
L118:
	;
	if v601&int32(255) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L119:
	;
	v572 = F_pg_utf_mblen_private(m, v509)
	mBase = m.M
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v574 = v572 + v573
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v574
	v576 = int32(26)
	if base.Ui32(v571-int32(91)) < base.Ui32(int32(3)) {
		v601 = v576
		v604 = v574
		goto L118
	} else {
		goto L132
	}
L120:
	;
	v569 = v512
	goto L122
L121:
	;
	if v512&int32(224) == int32(192) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	if v569 != 0 {
		v571 = v569
		goto L119
	} else {
		goto L131
	}
L123:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561+v509))))
	v569 = v562 | v564&int32(63)
	goto L122
L124:
	;
	v561 = int32(1)
	v562 = v512 << (uint(int32(6)) % 32) & int32(1984)
	goto L123
L125:
	;
	goto L126
L126:
	;
	if v512&int32(240) == int32(224) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
	v561 = int32(2)
	v562 = v512<<(uint(int32(12))%32)&int32(_a_F_daitch_mokotoff_4) | v533&int32(63)<<(uint(int32(6))%32)
	goto L123
L128:
	;
	goto L129
L129:
	;
	if v512&int32(248) != int32(240) {
		v571 = int32(-1)
		goto L119
	} else {
		goto L130
	}
L130:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
	v550 = int32(63)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+2)))
	v561 = int32(3)
	v562 = v512<<(uint(int32(18))%32)&int32(_a_F_daitch_mokotoff_5) | v549&v550<<(uint(int32(12))%32) | v555&v550<<(uint(int32(6))%32)
	goto L123
L131:
	;
	v601 = int32(0)
	v604 = v508
	goto L118
L132:
	;
	if base.Ui32(v571) <= base.Ui32(int32(95)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v601 = v571
	v604 = v574
	goto L118
L134:
	;
	goto L135
L135:
	;
	if base.Ui32(v571) <= base.Ui32(int32(255)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+uint32(_c_F_daitch_mokotoff[1]))))
	v601 = v587
	v604 = v574
	goto L118
L137:
	;
	goto L138
L138:
	;
	switch v571 - int32(260) {
	case 0, 1:
		v619 = int32(91)
		goto L114
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		v601 = v576
		v604 = v574
		goto L118
	case 20, 21:
		goto L117
	default:
		goto L139
	}
L139:
	;
	v594 = int32(2)
	if base.B2i32(base.Ui32(v571-int32(354)) < base.Ui32(v594))|base.B2i32(base.Ui32(v571-int32(538)) < base.Ui32(v594)) != 0 {
		v619 = int32(93)
		goto L114
	} else {
		goto L140
	}
L140:
	;
	v601 = v576
	v604 = v574
	goto L118
L141:
	;
	v626 = base.I32_extend8_s(v601)
	goto L113
L142:
	;
	goto L143
L143:
	;
	if base.Ui32(int32(28)) < base.Ui32((v601-int32(65))&int32(255)) {
		v508 = v604
		goto L115
	} else {
		goto L144
	}
L144:
	;
	v619 = v601
	goto L114
L145:
	;
	v857 = int32(0)
	goto L112
L146:
	;
	goto L147
L147:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v630
	v635 = v626*int32(12) + int32(_a_F_daitch_mokotoff_6)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v635-int32(772))))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v635-int32(776))))
	if v641 == int32(0) {
		v857 = v638
		goto L112
	} else {
		goto L148
	}
L148:
	;
	v646 = v641
	v647 = v638
	v655 = v630
	goto L149
L149:
	;
	v664 = v22 + int32(24)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v675 = v669
	goto L154
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v845
	v857 = v837
	goto L112
L151:
	;
	goto L150
L152:
	;
	if v793 == int32(0) {
		v837 = v647
		v845 = v655
		goto L151
	} else {
		goto L184
	}
L153:
	;
	v793 = base.I32_extend8_s(v786)
	goto L152
L154:
	;
	v676 = v74 + v675
	v677 = int32(*(*int8)(unsafe.Add(mBase, uint32(v676))))
	v679 = v677 & int32(255)
	if int32(0) <= v677 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v786 = int32(92)
	goto L153
L156:
	;
	goto L155
L157:
	;
	if v768&int32(255) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L158:
	;
	v739 = F_pg_utf_mblen_private(m, v676)
	mBase = m.M
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v664)))
	v741 = v739 + v740
	*(*int32)(unsafe.Add(mBase, uint32(v664))) = v741
	v743 = int32(26)
	if base.Ui32(v738-int32(91)) < base.Ui32(int32(3)) {
		v768 = v743
		v771 = v741
		goto L157
	} else {
		goto L171
	}
L159:
	;
	v736 = v679
	goto L161
L160:
	;
	if v679&int32(224) == int32(192) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if v736 != 0 {
		v738 = v736
		goto L158
	} else {
		goto L170
	}
L162:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v676))))
	v736 = v729 | v731&int32(63)
	goto L161
L163:
	;
	v728 = int32(1)
	v729 = v679 << (uint(int32(6)) % 32) & int32(1984)
	goto L162
L164:
	;
	goto L165
L165:
	;
	if v679&int32(240) == int32(224) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	v728 = int32(2)
	v729 = v679<<(uint(int32(12))%32)&int32(_a_F_daitch_mokotoff_4) | v700&int32(63)<<(uint(int32(6))%32)
	goto L162
L167:
	;
	goto L168
L168:
	;
	if v679&int32(248) != int32(240) {
		v738 = int32(-1)
		goto L158
	} else {
		goto L169
	}
L169:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	v717 = int32(63)
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+2)))
	v728 = int32(3)
	v729 = v679<<(uint(int32(18))%32)&int32(_a_F_daitch_mokotoff_5) | v716&v717<<(uint(int32(12))%32) | v722&v717<<(uint(int32(6))%32)
	goto L162
L170:
	;
	v768 = int32(0)
	v771 = v675
	goto L157
L171:
	;
	if base.Ui32(v738) <= base.Ui32(int32(95)) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v768 = v738
	v771 = v741
	goto L157
L173:
	;
	goto L174
L174:
	;
	if base.Ui32(v738) <= base.Ui32(int32(255)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+uint32(_c_F_daitch_mokotoff[1]))))
	v768 = v754
	v771 = v741
	goto L157
L176:
	;
	goto L177
L177:
	;
	switch v738 - int32(260) {
	case 0, 1:
		v786 = int32(91)
		goto L153
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		v768 = v743
		v771 = v741
		goto L157
	case 20, 21:
		goto L156
	default:
		goto L178
	}
L178:
	;
	v761 = int32(2)
	if base.B2i32(base.Ui32(v738-int32(354)) < base.Ui32(v761))|base.B2i32(base.Ui32(v738-int32(538)) < base.Ui32(v761)) != 0 {
		v786 = int32(93)
		goto L153
	} else {
		goto L179
	}
L179:
	;
	v768 = v743
	v771 = v741
	goto L157
L180:
	;
	v793 = base.I32_extend8_s(v768)
	goto L152
L181:
	;
	goto L182
L182:
	;
	if base.Ui32(int32(28)) < base.Ui32((v768-int32(65))&int32(255)) {
		v675 = v771
		goto L154
	} else {
		goto L183
	}
L183:
	;
	v786 = v768
	goto L153
L184:
	;
	v796 = int32(0)
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	if v797 == v796 {
		v837 = v647
		v845 = v655
		goto L151
	} else {
		goto L185
	}
L185:
	;
	v806 = v796
	v807 = v797
	goto L186
L186:
	;
	if v807 != v793 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v829 = v646 + v806*int32(12)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+8))
	if v830 != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v821 = v806 + int32(1)
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v821*int32(12)))))
	if v825 != 0 {
		v806 = v821
		v807 = v825
		goto L186
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	goto L187
L191:
	;
	v837 = v647
	v845 = v655
	goto L151
L192:
	;
	v831 = v826
	goto L194
L193:
	;
	v831 = v655
	goto L194
L194:
	;
	if v830 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v832 = v830
	goto L197
L196:
	;
	v832 = v647
	goto L197
L197:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v829)+4))
	if v833 != 0 {
		v646 = v833
		v647 = v832
		v655 = v831
		goto L149
	} else {
		goto L198
	}
L198:
	;
	v837 = v832
	v845 = v831
	goto L151
L199:
	;
	v890 = v857
	goto L201
L200:
	;
	v890 = int32(_a_F_daitch_mokotoff_8)
	goto L201
L201:
	;
	v896 = v493
	goto L202
L202:
	;
	v910 = int32(*(*int8)(unsafe.Add(mBase, uint32(v473))))
	if v910 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v857 != 0 {
		v469 = v874
		v473 = v857
		v477 = v477 + int32(1)
		goto L109
	} else {
		goto L247
	}
L204:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v896+v489)+76))
	if v1007 != 0 {
		v896 = v1007
		goto L202
	} else {
		goto L246
	}
L205:
	;
	v913 = int32(*(*int8)(unsafe.Add(mBase, uint32(v890))))
	if v913 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v958 = int32(*(*int8)(unsafe.Add(mBase, uint32(v888))))
	if v958 == int32(0) {
		goto L204
	} else {
		goto L226
	}
L207:
	;
	v917 = v22 + int32(12)
	v919 = v22 + int32(24)
	if int32(49) < v910 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v924 = int32(2)
	goto L210
L209:
	;
	v924 = int32(1)
	goto L210
L210:
	;
	if v913 < int32(50) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v929 = int32(1)
	goto L213
L212:
	;
	v929 = int32(2)
	goto L213
L213:
	;
	if v477 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v931 = v929
	goto L216
L215:
	;
	v931 = int32(0)
	goto L216
L216:
	;
	F_update_node(m, v917, v919, v896, v874, v477, v924, v931, v473+v931*int32(3), int32(0), v78)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v938 = int32(*(*int8)(unsafe.Add(mBase, uint32(v890)+9)))
	if v938 == int32(0) {
		goto L206
	} else {
		goto L218
	}
L218:
	;
	if v938 < int32(50) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v945 = int32(1)
	goto L221
L220:
	;
	v945 = int32(2)
	goto L221
L221:
	;
	if v477 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v947 = v945
	goto L224
L223:
	;
	v947 = int32(0)
	goto L224
L224:
	;
	F_update_node(m, v917, v919, v896, v874, v477, v924, v947, v473+v947*int32(3), int32(0), v78)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	goto L206
L226:
	;
	v961 = int32(*(*int8)(unsafe.Add(mBase, uint32(v890))))
	if v961 == int32(0) {
		goto L204
	} else {
		goto L227
	}
L227:
	;
	v965 = v22 + int32(12)
	v967 = v22 + int32(24)
	if int32(49) < v958 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v972 = int32(2)
	goto L230
L229:
	;
	v972 = int32(1)
	goto L230
L230:
	;
	if v961 < int32(50) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v977 = int32(1)
	goto L233
L232:
	;
	v977 = int32(2)
	goto L233
L233:
	;
	if v477 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v979 = v977
	goto L236
L235:
	;
	v979 = int32(0)
	goto L236
L236:
	;
	F_update_node(m, v965, v967, v896, v874, v477, v972, v979, v888+v979*int32(3), int32(0), v78)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v986 = int32(*(*int8)(unsafe.Add(mBase, uint32(v890)+9)))
	if v986 == int32(0) {
		goto L204
	} else {
		goto L238
	}
L238:
	;
	if v986 < int32(50) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v993 = int32(1)
	goto L241
L240:
	;
	v993 = int32(2)
	goto L241
L241:
	;
	if v477 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v995 = v993
	goto L244
L243:
	;
	v995 = int32(0)
	goto L244
L244:
	;
	F_update_node(m, v965, v967, v896, v874, v477, v972, v995, v888+v995*int32(3), int32(0), v78)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	goto L204
L246:
	;
	goto L203
L247:
	;
	goto L110
L248:
	;
	v1020 = v1010
	goto L249
L249:
	;
	v1037 = F_cstring_to_text_with_len(m, v1020+int32(4), int32(6))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L251
	}
L250:
	;
	goto L108
L251:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0]))
	v1044 = F_accumArrayResult(m, v78, base.I64_extend_i32_u(v1037), int32(0), int32(25), v1043)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v874<<(uint(int32(2))%32)+v1020)+76))
	if v1047 != 0 {
		v1020 = v1047
		goto L249
	} else {
		goto L253
	}
L253:
	;
	goto L250
L254:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_daitch_mokotoff[0])) = v38
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1116 = v1067
	goto L18
L256:
	;
	v1096 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1096)
	v1116 = int64(0)
	goto L18
}
func F_dasin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 float64
	_ = v9
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v46 float64
	_ = v46
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v60 float64
	_ = v60
	var v65 float64
	_ = v65
	var v69 float64
	_ = v69
	var v78 float64
	_ = v78
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v100 float64
	_ = v100
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v4&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v9 = base.F64_reinterpret_i64(v4)
		if base.F64_gt(base.F64_abs(v9), float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_dasin_0), int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_dasin_1), int32(1844), int32(_a_F_dasin_2))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = base.I64_reinterpret_f64(v9)
			v23 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072693248)) <= base.Ui32(v23) {
				if base.I32_wrap_i64(v18)|(v23-int32(1072693248)) == int32(0) {
					v100 = base.F64_add(base.F64_mul(v9, float64(1.5707963267948966)), float64(7.52316384526264e-37))
				} else {
					v100 = base.F64_div(float64(0), base.F64_sub(v9, v9))
				}
			} else {
				if base.Ui32(v23) <= base.Ui32(int32(1071644671)) {
					if base.Ui32(v23+int32(-1048576)) < base.Ui32(int32(1044381696)) {
						v92 = v9
						v100 = v92
					} else {
						v46 = F_R(m, base.F64_mul(v9, v9))
						mBase = m.M
						v100 = base.F64_add(base.F64_mul(v9, v46), v9)
					}
				} else {
					v53 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v9)), float64(0.5))
					v54 = base.F64_sqrt(v53)
					v55 = F_R(m, v53)
					mBase = m.M
					if base.Ui32(int32(1072640819)) <= base.Ui32(v23) {
						v60 = base.F64_add(base.F64_mul(v54, v55), v54)
						v87 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v60, v60), float64(-6.123233995736766e-17)))
					} else {
						v65 = float64(0.7853981633974483)
						v69 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v54) & int64(-4294967296))
						v78 = base.F64_div(base.F64_sub(v53, base.F64_mul(v69, v69)), base.F64_add(v54, v69))
						v87 = base.F64_add(base.F64_sub(base.F64_sub(v65, base.F64_add(v69, v69)), base.F64_sub(base.F64_mul(base.F64_add(v54, v54), v55), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v78, v78)))), v65)
					}
					if v18 < int64(0) {
						v91 = base.F64_neg(v87)
					} else {
						v91 = v87
					}
					v92 = v91
					v100 = v92
				}
			}
			if base.F64_eq(base.F64_abs(v100), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v107 = base.I64_reinterpret_f64(v100)
				return v107
			}
		}
	} else {
		v107 = int64(9221120237041090560)
		return v107
	}
}
func F_dasinh(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v29 float64
	_ = v29
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v46 int64
	_ = v46
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v86 float64
	_ = v86
	var v91 float64
	_ = v91
	var v105 float64
	_ = v105
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v156 float64
	_ = v156
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v168 float64
	_ = v168
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.F64_abs(v5)
	v7 = base.I64_reinterpret_f64(v5)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1049)) <= base.Ui32(v12) {
		v15 = F_log(m, v6)
		mBase = m.M
		v164 = base.F64_add(v15, float64(0.6931471805599453))
	} else {
		if base.Ui32(int32(1024)) <= base.Ui32(v12) {
			v21 = float64(1)
			v29 = F_log(m, base.F64_add(base.F64_add(v6, v6), base.F64_div(v21, base.F64_add(v6, base.F64_sqrt(base.F64_add(base.F64_mul(v5, v5), v21))))))
			mBase = m.M
			v164 = v29
		} else {
			if base.Ui32(v12) < base.Ui32(int32(997)) {
				v164 = v6
			} else {
				v32 = base.F64_mul(v5, v5)
				v33 = float64(1)
				v39 = base.F64_add(v6, base.F64_div(v32, base.F64_add(base.F64_sqrt(base.F64_add(v32, v33)), v33)))
				v40 = float64(0)
				v46 = base.I64_reinterpret_f64(v39)
				if v46 <= int64(4601133429810003967) {
					if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v46) {
						if base.F64_eq(v39, float64(-1)) != 0 {
							v156 = math.Float64frombits(uint64(0xfff0000000000000))
							v163 = v156
						} else {
							v163 = base.F64_div(base.F64_sub(v39, v39), float64(0))
						}
					} else {
						if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v46)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
							v163 = v39
						} else {
							if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v46) {
								v66 = float64(1)
								v67 = base.F64_add(v39, v66)
								v68 = base.I64_reinterpret_f64(v67)
								v73 = base.I32_wrap_i64(int64(base.Ui64(v68)>>(uint(int64(32))%64))) + int32(_a_F_dasinh_0)
								if base.Ui32(int32(1074790399)) < base.Ui32(v73) {
									v86 = base.F64_add(base.F64_sub(v39, v67), v66)
								} else {
									v86 = base.F64_sub(v39, base.F64_add(v67, float64(-1)))
								}
								if base.Ui32(v73) <= base.Ui32(int32(1129316351)) {
									v91 = base.F64_div(v86, v67)
								} else {
									v91 = float64(0)
								}
								v105 = base.F64_convert_i32_s(int32(base.Ui32(v73)>>(uint(int32(20))%32)) - int32(1023))
								v109 = base.F64_add(base.F64_reinterpret_i64(v68&int64(4294967295)|base.I64_extend_i32_u(v73&int32(_a_F_dasinh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
								v110 = v105
								v111 = base.F64_add(base.F64_mul(v105, float64(1.9082149292705877e-10)), v91)
							} else {
								v109 = v39
								v110 = v40
								v111 = v40
							}
							v118 = base.F64_div(v109, base.F64_add(v109, float64(2)))
							v121 = base.F64_mul(v109, base.F64_mul(v109, float64(0.5)))
							v122 = base.F64_mul(v118, v118)
							v123 = base.F64_mul(v122, v122)
							v156 = base.F64_add(base.F64_mul(v110, float64(0.6931471803691238)), base.F64_add(v109, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v111), v121)))
							v163 = v156
						}
					}
				} else {
					if base.Ui64(int64(9218868437227405311)) < base.Ui64(v46) {
						v163 = v39
					} else {
						v66 = float64(1)
						v67 = base.F64_add(v39, v66)
						v68 = base.I64_reinterpret_f64(v67)
						v73 = base.I32_wrap_i64(int64(base.Ui64(v68)>>(uint(int64(32))%64))) + int32(_a_F_dasinh_0)
						if base.Ui32(int32(1074790399)) < base.Ui32(v73) {
							v86 = base.F64_add(base.F64_sub(v39, v67), v66)
						} else {
							v86 = base.F64_sub(v39, base.F64_add(v67, float64(-1)))
						}
						if base.Ui32(v73) <= base.Ui32(int32(1129316351)) {
							v91 = base.F64_div(v86, v67)
						} else {
							v91 = float64(0)
						}
						v105 = base.F64_convert_i32_s(int32(base.Ui32(v73)>>(uint(int32(20))%32)) - int32(1023))
						v109 = base.F64_add(base.F64_reinterpret_i64(v68&int64(4294967295)|base.I64_extend_i32_u(v73&int32(_a_F_dasinh_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v110 = v105
						v111 = base.F64_add(base.F64_mul(v105, float64(1.9082149292705877e-10)), v91)
						v118 = base.F64_div(v109, base.F64_add(v109, float64(2)))
						v121 = base.F64_mul(v109, base.F64_mul(v109, float64(0.5)))
						v122 = base.F64_mul(v118, v118)
						v123 = base.F64_mul(v122, v122)
						v156 = base.F64_add(base.F64_mul(v110, float64(0.6931471803691238)), base.F64_add(v109, base.F64_sub(base.F64_add(base.F64_mul(v118, base.F64_add(v121, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v122, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v111), v121)))
						v163 = v156
					}
				}
				v164 = v163
			}
		}
	}
	if v7 < int64(0) {
		v168 = base.F64_neg(v164)
	} else {
		v168 = v164
	}
	return base.I64_reinterpret_f64(v168)
}
func F_dbase_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	v16 = v14 & int32(240)
	switch v16 - int32(16) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L1
	case 16:
		goto L2
	default:
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return
L2:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_dbase_desc_0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L9
	}
L3:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_rotl(v30, int64(32))
	F_appendStringInfo(m, l0, int32(_a_F_dbase_desc_1), v10+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L8
	}
L4:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v21 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = base.I64_rotl(v20, v21)
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = base.I64_rotl(v19, v21)
	F_appendStringInfo(m, l0, int32(_a_F_dbase_desc_2), v10)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	goto L1
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v42 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)+v49<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v58
	F_appendStringInfo(m, l0, int32(_a_F_dbase_desc_3), v10+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L1
L13:
	;
	v68 = v49 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v68 < v69 {
		v49 = v68
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_dcs_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v9 == int32(0))|base.B2i32(v9 != v12) != 0 {
		v30 = v9
		v31 = v12
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v31
L2:
	;
	goto L1
L3:
	;
	v15 = v4
	v16 = v6
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v20
		v31 = v19
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v30 = v20
	v31 = v19
	goto L2
L6:
	;
	v23 = int32(1)
	if v20 == v19 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_deltraverse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = m.T0[v9].(func(*base.Module) int32)(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(101)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v18
	return
L9:
	;
	return
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v23 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v29 = v25
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	goto L9
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	F_deltraverse(m, l0, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if v35 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v42 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v111 != 0 {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if v77 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v47 = v45 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v47))|base.B2i32(int32(1)<<(uint(v47)%32)&int32(_a_F_deltraverse_0) == int32(0)) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v57 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v58 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v42*int32(24))+12)) = v66
	v70 = v66
	goto L24
L26:
	;
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v68
	v70 = v68
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v58
	goto L30
L29:
	;
	goto L30
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = int64(0)
	goto L20
L31:
	;
	if v76 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v76
	goto L31
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v76
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v77
	goto L37
L36:
	;
	goto L37
L37:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v83 - int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if v88 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v87 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v87
	goto L38
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = v87
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v88
	goto L44
L43:
	;
	goto L44
L44:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v94 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	v101 = v29 + int32(8)
	v102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+16)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v101)+8)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v101))) = v102
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29
	goto L19
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v131 != 0 {
		v29 = v131
		goto L15
	} else {
		goto L56
	}
L46:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v112 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(-1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v118 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v117 != 0 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+32)) = v117
	goto L48
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v117
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v31
	goto L45
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v121
	goto L52
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v121
	goto L52
L56:
	;
	goto L16
}
func F_dense_alloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v8 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_dense_alloc_0)) <= base.Ui32(v8) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v14 = F_MemoryContextAlloc(m, v11, v8+int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v8
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v14
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v14
			}
			return v14 + int32(16)
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		if v33 != 0 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
			if base.Ui32(v8) <= base.Ui32(v34-v35) {
				*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v8 + v35
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
				*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55 + int32(1)
				return v33 + v35 + int32(16)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v41 = F_MemoryContextAlloc(m, v39, int32(_a_F_dense_alloc_1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v8
					*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(140737488355329)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v41
					return v41 + int32(16)
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
			v41 = F_MemoryContextAlloc(m, v39, int32(_a_F_dense_alloc_1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v8
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(140737488355329)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v41
				return v41 + int32(16)
			}
		}
	}
}
func F_derf(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v36 float64
	_ = v36
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v88 int32
	_ = v88
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.I64_reinterpret_f64(v6)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v7) >> (uint(int64(32)) % 64)))
	v12 = v10 & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v12) {
		v81 = base.F64_add(base.F64_div(float64(1), v6), base.F64_convert_i32_s(int32(1)-int32(base.Ui32(v10)>>(uint(int32(30))%32))&int32(2)))
	} else {
		if base.Ui32(v12) <= base.Ui32(int32(1072365567)) {
			if base.Ui32(v12) <= base.Ui32(int32(1043333119)) {
				v81 = base.F64_mul(base.F64_add(base.F64_mul(v6, float64(8)), base.F64_mul(v6, float64(1.0270333367641007))), float64(0.125))
			} else {
				v36 = base.F64_mul(v6, v6)
				v81 = base.F64_add(base.F64_mul(v6, base.F64_div(base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(-2.3763016656650163e-05)), float64(-0.005770270296489442))), float64(-0.02848174957559851))), float64(-0.3250421072470015))), float64(0.12837916709551256)), base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(-3.960228278775368e-06)), float64(0.00013249473800432164))), float64(0.005081306281875766))), float64(0.0650222499887673))), float64(0.39791722395915535))), float64(1)))), v6)
			}
		} else {
			if base.Ui32(v12) <= base.Ui32(int32(1075314687)) {
				v73 = F_erfc2(m, v12, v6)
				mBase = m.M
				v75 = base.F64_sub(float64(1), v73)
			} else {
				v75 = float64(1)
			}
			if v7 < int64(0) {
				v79 = base.F64_neg(v75)
			} else {
				v79 = v75
			}
			v81 = v79
		}
	}
	if base.F64_eq(base.F64_abs(v81), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I64_reinterpret_f64(v81)
	}
}
func F_des_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v36 int32
	_ = v36
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int64
	_ = v744
	var v768 int32
	_ = v768
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v819 int32
	_ = v819
	var v823 int64
	_ = v823
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v967 int32
	_ = v967
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1594 int32
	_ = v1594
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1728 int32
	_ = v1728
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_des_init[0])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_des_init[1])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_des_init[2])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_des_init[3])) = v1
	v36 = v1
	for {
		v67 = v36&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v36)>>(uint(int32(1))%32))&int32(15)
		v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+16)))
		*(*uint8)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_des_init[4]))) = uint8(v68)
		v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
		*(*uint8)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_des_init[5]))) = uint8(v70)
		v73 = v36 + int32(2)
		if v73 != int32(64) {
			v36 = v73
			continue
		} else {
			break
		}
		break
	}
	v76 = v1
	for {
		v108 = v76&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v76)>>(uint(int32(1))%32))&int32(15)
		v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+80)))
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_des_init[6]))) = uint8(v109)
		v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+64)))
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_des_init[7]))) = uint8(v111)
		v114 = v76 + int32(2)
		if v114 != int32(64) {
			v76 = v114
			continue
		} else {
			break
		}
		break
	}
	v118 = int32(0)
	for {
		v150 = v118&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v118)>>(uint(int32(1))%32))&int32(15)
		v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+144)))
		*(*uint8)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_des_init[8]))) = uint8(v151)
		v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+128)))
		*(*uint8)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_des_init[9]))) = uint8(v153)
		v156 = v118 + int32(2)
		if v156 != int32(64) {
			v118 = v156
			continue
		} else {
			break
		}
		break
	}
	v160 = int32(0)
	for {
		v192 = v160&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v160)>>(uint(int32(1))%32))&int32(15)
		v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+208)))
		*(*uint8)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_des_init[10]))) = uint8(v193)
		v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+192)))
		*(*uint8)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_des_init[11]))) = uint8(v195)
		v198 = v160 + int32(2)
		if v198 != int32(64) {
			v160 = v198
			continue
		} else {
			break
		}
		break
	}
	v202 = int32(0)
	for {
		v234 = v202&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v202)>>(uint(int32(1))%32))&int32(15)
		v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+272)))
		*(*uint8)(unsafe.Add(mBase, uint32(v202)+uint32(_c_F_des_init[12]))) = uint8(v235)
		v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+256)))
		*(*uint8)(unsafe.Add(mBase, uint32(v202)+uint32(_c_F_des_init[13]))) = uint8(v237)
		v240 = v202 + int32(2)
		if v240 != int32(64) {
			v202 = v240
			continue
		} else {
			break
		}
		break
	}
	v244 = int32(0)
	for {
		v276 = v244&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v244)>>(uint(int32(1))%32))&int32(15)
		v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+336)))
		*(*uint8)(unsafe.Add(mBase, uint32(v244)+uint32(_c_F_des_init[14]))) = uint8(v277)
		v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+320)))
		*(*uint8)(unsafe.Add(mBase, uint32(v244)+uint32(_c_F_des_init[15]))) = uint8(v279)
		v282 = v244 + int32(2)
		if v282 != int32(64) {
			v244 = v282
			continue
		} else {
			break
		}
		break
	}
	v286 = int32(0)
	for {
		v318 = v286&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v286)>>(uint(int32(1))%32))&int32(15)
		v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+400)))
		*(*uint8)(unsafe.Add(mBase, uint32(v286)+uint32(_c_F_des_init[16]))) = uint8(v319)
		v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+384)))
		*(*uint8)(unsafe.Add(mBase, uint32(v286)+uint32(_c_F_des_init[17]))) = uint8(v321)
		v324 = v286 + int32(2)
		if v324 != int32(64) {
			v286 = v324
			continue
		} else {
			break
		}
		break
	}
	v328 = int32(0)
	for {
		v360 = v328&int32(32) + int32(_a_F_des_init_0) + int32(base.Ui32(v328)>>(uint(int32(1))%32))&int32(15)
		v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+464)))
		*(*uint8)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_des_init[18]))) = uint8(v361)
		v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+448)))
		*(*uint8)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_des_init[19]))) = uint8(v363)
		v366 = v328 + int32(2)
		if v366 != int32(64) {
			v328 = v366
			continue
		} else {
			break
		}
		break
	}
	v371 = v1
	for {
		v394 = v371<<(uint(int32(6))%32) + int32(_a_F_des_init_1)
		v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+uint32(_c_F_des_init[5]))))
		v399 = v397 << (uint(int32(4)) % 32)
		v401 = int32(0)
		for {
			v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+uint32(_c_F_des_init[7]))))
			v427 = v399 | v426
			*(*uint8)(unsafe.Add(mBase, uint32(v401+v394))) = uint8(v427)
			v430 = v401 | int32(1)
			v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+uint32(_c_F_des_init[7]))))
			v435 = v399 | v434
			*(*uint8)(unsafe.Add(mBase, uint32(v394+v430))) = uint8(v435)
			v438 = v401 | int32(2)
			v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+uint32(_c_F_des_init[7]))))
			v443 = v399 | v442
			*(*uint8)(unsafe.Add(mBase, uint32(v394+v438))) = uint8(v443)
			v446 = v401 | int32(3)
			v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+uint32(_c_F_des_init[7]))))
			v451 = v399 | v450
			*(*uint8)(unsafe.Add(mBase, uint32(v394+v446))) = uint8(v451)
			v454 = v401 + int32(4)
			if v454 != int32(64) {
				v401 = v454
				continue
			} else {
				break
			}
			break
		}
		v458 = v371 + int32(1)
		if v458 != int32(64) {
			v371 = v458
			continue
		} else {
			break
		}
		break
	}
	v464 = int32(0)
	for {
		v489 = v464<<(uint(int32(6))%32) + int32(_a_F_des_init_2)
		v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+uint32(_c_F_des_init[9]))))
		v494 = v492 << (uint(int32(4)) % 32)
		v495 = int32(0)
		for {
			v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+uint32(_c_F_des_init[11]))))
			v521 = v494 | v520
			*(*uint8)(unsafe.Add(mBase, uint32(v495+v489))) = uint8(v521)
			v524 = v495 | int32(1)
			v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+uint32(_c_F_des_init[11]))))
			v529 = v494 | v528
			*(*uint8)(unsafe.Add(mBase, uint32(v489+v524))) = uint8(v529)
			v532 = v495 | int32(2)
			v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+uint32(_c_F_des_init[11]))))
			v537 = v494 | v536
			*(*uint8)(unsafe.Add(mBase, uint32(v489+v532))) = uint8(v537)
			v540 = v495 | int32(3)
			v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_c_F_des_init[11]))))
			v545 = v494 | v544
			*(*uint8)(unsafe.Add(mBase, uint32(v489+v540))) = uint8(v545)
			v548 = v495 + int32(4)
			if v548 != int32(64) {
				v495 = v548
				continue
			} else {
				break
			}
			break
		}
		v552 = v464 + int32(1)
		if v552 != int32(64) {
			v464 = v552
			continue
		} else {
			break
		}
		break
	}
	v558 = int32(0)
	for {
		v583 = v558<<(uint(int32(6))%32) + int32(_a_F_des_init_3)
		v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558)+uint32(_c_F_des_init[13]))))
		v588 = v586 << (uint(int32(4)) % 32)
		v589 = int32(0)
		for {
			v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+uint32(_c_F_des_init[15]))))
			v615 = v588 | v614
			*(*uint8)(unsafe.Add(mBase, uint32(v589+v583))) = uint8(v615)
			v618 = v589 | int32(1)
			v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+uint32(_c_F_des_init[15]))))
			v623 = v588 | v622
			*(*uint8)(unsafe.Add(mBase, uint32(v583+v618))) = uint8(v623)
			v626 = v589 | int32(2)
			v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+uint32(_c_F_des_init[15]))))
			v631 = v588 | v630
			*(*uint8)(unsafe.Add(mBase, uint32(v583+v626))) = uint8(v631)
			v634 = v589 | int32(3)
			v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+uint32(_c_F_des_init[15]))))
			v639 = v588 | v638
			*(*uint8)(unsafe.Add(mBase, uint32(v583+v634))) = uint8(v639)
			v642 = v589 + int32(4)
			if v642 != int32(64) {
				v589 = v642
				continue
			} else {
				break
			}
			break
		}
		v646 = v558 + int32(1)
		if v646 != int32(64) {
			v558 = v646
			continue
		} else {
			break
		}
		break
	}
	v652 = int32(0)
	for {
		v677 = v652<<(uint(int32(6))%32) + int32(_a_F_des_init_4)
		v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+uint32(_c_F_des_init[17]))))
		v682 = v680 << (uint(int32(4)) % 32)
		v683 = int32(0)
		for {
			v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+uint32(_c_F_des_init[19]))))
			v709 = v682 | v708
			*(*uint8)(unsafe.Add(mBase, uint32(v683+v677))) = uint8(v709)
			v712 = v683 | int32(1)
			v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712)+uint32(_c_F_des_init[19]))))
			v717 = v682 | v716
			*(*uint8)(unsafe.Add(mBase, uint32(v677+v712))) = uint8(v717)
			v720 = v683 | int32(2)
			v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+uint32(_c_F_des_init[19]))))
			v725 = v682 | v724
			*(*uint8)(unsafe.Add(mBase, uint32(v677+v720))) = uint8(v725)
			v728 = v683 | int32(3)
			v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+uint32(_c_F_des_init[19]))))
			v733 = v682 | v732
			*(*uint8)(unsafe.Add(mBase, uint32(v677+v728))) = uint8(v733)
			v736 = v683 + int32(4)
			if v736 != int32(64) {
				v683 = v736
				continue
			} else {
				break
			}
			break
		}
		v740 = v652 + int32(1)
		if v740 != int32(64) {
			v652 = v740
			continue
		} else {
			break
		}
		break
	}
	v744 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[20])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[21])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[22])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[23])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[24])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[25])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[26])) = v744
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[27])) = v744
	v768 = int32(0)
	for {
		v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+uint32(_c_F_des_init[28]))))
		v795 = int32(1)
		v796 = v794 - v795
		*(*uint8)(unsafe.Add(mBase, uint32(v768)+uint32(_c_F_des_init[29]))) = uint8(v796)
		v798 = int32(255)
		*(*uint8)(unsafe.Add(mBase, uint32(v796&v798)+uint32(_c_F_des_init[30]))) = uint8(v768)
		v804 = v768 | v795
		v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+uint32(_c_F_des_init[28]))))
		v811 = v809 - v795
		*(*uint8)(unsafe.Add(mBase, uint32(v804)+uint32(_c_F_des_init[29]))) = uint8(v811)
		*(*uint8)(unsafe.Add(mBase, uint32(v811&v798)+uint32(_c_F_des_init[30]))) = uint8(v804)
		v819 = v768 + int32(2)
		if v819 != int32(64) {
			v768 = v819
			continue
		} else {
			break
		}
		break
	}
	v823 = int64(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[31])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[32])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[33])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[34])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[35])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[36])) = v823
	*(*int64)(unsafe.Add(mBase, _c_F_des_init[37])) = v823
	v843 = int32(0)
	v846 = v843
	for {
		v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+uint32(_c_F_des_init[38]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v869)+uint32(_c_F_des_init[39]))) = uint8(v846)
		v874 = v846 | int32(1)
		v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+uint32(_c_F_des_init[38]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v877)+uint32(_c_F_des_init[39]))) = uint8(v874)
		v882 = v846 | int32(2)
		v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+uint32(_c_F_des_init[38]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v885)+uint32(_c_F_des_init[39]))) = uint8(v882)
		v890 = v846 | int32(3)
		v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890)+uint32(_c_F_des_init[38]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v893)+uint32(_c_F_des_init[39]))) = uint8(v890)
		v898 = v846 + int32(4)
		if v898 != int32(56) {
			v846 = v898
			continue
		} else {
			break
		}
		break
	}
	v901 = v843
	for {
		v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+uint32(_c_F_des_init[40]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v925)+uint32(_c_F_des_init[41]))) = uint8(v901)
		v930 = v901 | int32(1)
		v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+uint32(_c_F_des_init[40]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v933)+uint32(_c_F_des_init[41]))) = uint8(v930)
		v938 = v901 | int32(2)
		v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+uint32(_c_F_des_init[40]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v941)+uint32(_c_F_des_init[41]))) = uint8(v938)
		v946 = v901 | int32(3)
		v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+uint32(_c_F_des_init[40]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v949)+uint32(_c_F_des_init[41]))) = uint8(v946)
		v954 = v901 + int32(4)
		if v954 != int32(48) {
			v901 = v954
			continue
		} else {
			break
		}
		break
	}
	v967 = v1
	for {
		v980 = v967 << (uint(int32(10)) % 32)
		v990 = v967 << (uint(int32(3)) % 32)
		v993 = int32(0)
		for {
			v1015 = v993 << (uint(int32(2)) % 32)
			v1016 = v980 + int32(_a_F_des_init_5) + v1015
			v1017 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1017
			v1019 = v1015 + (v980 + int32(_a_F_des_init_6))
			*(*int32)(unsafe.Add(mBase, uint32(v1019))) = v1017
			v1022 = v1015 + (v980 + int32(_a_F_des_init_7))
			*(*int32)(unsafe.Add(mBase, uint32(v1022))) = v1017
			v1025 = v1015 + (v980 + int32(_a_F_des_init_8))
			*(*int32)(unsafe.Add(mBase, uint32(v1025))) = v1017
			v1033 = v1017
			v1035 = v1017
			v1037 = v1017
			v1040 = v1017
			v1042 = v1017
			for {
				v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033)+uint32(_c_F_des_init[42]))))
				if v993&v1057 == int32(0) {
					v1098 = v1035
					v1099 = v1037
					v1101 = v1040
					v1102 = v1042
				} else {
					v1061 = v1033 + v990
					v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061)+uint32(_c_F_des_init[30]))))
					v1066 = v1064 << (uint(int32(2)) % 32)
					if base.Ui32(v1064) <= base.Ui32(int32(31)) {
						v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+uint32(_c_F_des_init[43])))
						v1072 = v1040 | v1071
						*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1072
						v1079 = v1072
						v1080 = v1042
					} else {
						v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1066+int32(_a_F_des_init_9)-int32(128))))
						v1077 = v1042 | v1076
						*(*int32)(unsafe.Add(mBase, uint32(v1019))) = v1077
						v1079 = v1040
						v1080 = v1077
					}
					v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061)+uint32(_c_F_des_init[29]))))
					v1085 = v1083 << (uint(int32(2)) % 32)
					if base.Ui32(v1083) <= base.Ui32(int32(31)) {
						v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+uint32(_c_F_des_init[43])))
						v1091 = v1037 | v1090
						*(*int32)(unsafe.Add(mBase, uint32(v1022))) = v1091
						v1098 = v1035
						v1099 = v1091
						v1101 = v1079
						v1102 = v1080
					} else {
						v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1085+int32(_a_F_des_init_9)-int32(128))))
						v1096 = v1035 | v1095
						*(*int32)(unsafe.Add(mBase, uint32(v1025))) = v1096
						v1098 = v1096
						v1099 = v1037
						v1101 = v1079
						v1102 = v1080
					}
				}
				v1106 = v1033 + int32(1)
				if v1106 != int32(8) {
					v1033 = v1106
					v1035 = v1098
					v1037 = v1099
					v1040 = v1101
					v1042 = v1102
					continue
				} else {
					break
				}
				break
			}
			v1110 = v993 + int32(1)
			if v1110 != int32(256) {
				v993 = v1110
				continue
			} else {
				break
			}
			break
		}
		v1114 = v967 << (uint(int32(9)) % 32)
		v1125 = v967 * int32(7)
		v1126 = int32(0)
		for {
			v1149 = v1126 << (uint(int32(2)) % 32)
			v1150 = v1114 + int32(_a_F_des_init_10) + v1149
			v1151 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1151
			v1153 = v1149 + (v1114 + int32(_a_F_des_init_11))
			*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1151
			v1159 = v1126 & int32(64)
			if v1159 == v1151 {
				v1180 = v1151
				v1181 = v1151
			} else {
				v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[27]))))
				if v1164 == int32(255) {
					v1180 = v1151
					v1181 = v1151
				} else {
					v1168 = v1164 << (uint(int32(2)) % 32)
					if base.Ui32(v1164) <= base.Ui32(int32(27)) {
						v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+uint32(_c_F_des_init[44])))
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1173
						v1180 = v1151
						v1181 = v1173
					} else {
						v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1168+int32(_a_F_des_init_12)-int32(112))))
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1177
						v1180 = v1177
						v1181 = int32(0)
					}
				}
			}
			v1184 = v1126 & int32(32)
			if v1184 == int32(0) {
				v1206 = v1180
				v1207 = v1181
			} else {
				v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[45]))))
				if v1189 == int32(255) {
					v1206 = v1180
					v1207 = v1181
				} else {
					v1193 = v1189 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1189) {
						v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1193+int32(_a_F_des_init_12)-int32(112))))
						v1201 = v1180 | v1200
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1201
						v1206 = v1201
						v1207 = v1181
					} else {
						v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+uint32(_c_F_des_init[44])))
						v1204 = v1181 | v1203
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1204
						v1206 = v1180
						v1207 = v1204
					}
				}
			}
			v1211 = v1126 & int32(16)
			if v1211 == int32(0) {
				v1233 = v1206
				v1234 = v1207
			} else {
				v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[46]))))
				if v1216 == int32(255) {
					v1233 = v1206
					v1234 = v1207
				} else {
					v1220 = v1216 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1216) {
						v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1220+int32(_a_F_des_init_12)-int32(112))))
						v1228 = v1206 | v1227
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1228
						v1233 = v1228
						v1234 = v1207
					} else {
						v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+uint32(_c_F_des_init[44])))
						v1231 = v1207 | v1230
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1231
						v1233 = v1206
						v1234 = v1231
					}
				}
			}
			v1238 = v1126 & int32(8)
			if v1238 == int32(0) {
				v1260 = v1233
				v1261 = v1234
			} else {
				v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[47]))))
				if v1243 == int32(255) {
					v1260 = v1233
					v1261 = v1234
				} else {
					v1247 = v1243 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1243) {
						v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1247+int32(_a_F_des_init_12)-int32(112))))
						v1255 = v1233 | v1254
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1255
						v1260 = v1255
						v1261 = v1234
					} else {
						v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+uint32(_c_F_des_init[44])))
						v1258 = v1234 | v1257
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1258
						v1260 = v1233
						v1261 = v1258
					}
				}
			}
			v1265 = v1126 & int32(4)
			if v1265 == int32(0) {
				v1287 = v1260
				v1288 = v1261
			} else {
				v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[48]))))
				if v1270 == int32(255) {
					v1287 = v1260
					v1288 = v1261
				} else {
					v1274 = v1270 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1270) {
						v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1274+int32(_a_F_des_init_12)-int32(112))))
						v1282 = v1260 | v1281
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1282
						v1287 = v1282
						v1288 = v1261
					} else {
						v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+uint32(_c_F_des_init[44])))
						v1285 = v1261 | v1284
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1285
						v1287 = v1260
						v1288 = v1285
					}
				}
			}
			v1292 = v1126 & int32(2)
			if v1292 == int32(0) {
				v1314 = v1287
				v1315 = v1288
			} else {
				v1297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[49]))))
				if v1297 == int32(255) {
					v1314 = v1287
					v1315 = v1288
				} else {
					v1301 = v1297 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1297) {
						v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1301+int32(_a_F_des_init_12)-int32(112))))
						v1309 = v1287 | v1308
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1309
						v1314 = v1309
						v1315 = v1288
					} else {
						v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+uint32(_c_F_des_init[44])))
						v1312 = v1288 | v1311
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1312
						v1314 = v1287
						v1315 = v1312
					}
				}
			}
			v1319 = v1126 & int32(1)
			if v1319 == int32(0) {
			} else {
				v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+uint32(_c_F_des_init[50]))))
				if v1324 == int32(255) {
				} else {
					v1328 = v1324 << (uint(int32(2)) % 32)
					if base.Ui32(int32(28)) <= base.Ui32(v1324) {
						v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1328+int32(_a_F_des_init_12)-int32(112))))
						*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1314 | v1335
					} else {
						v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+uint32(_c_F_des_init[44])))
						*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1315 | v1338
					}
				}
			}
			v1343 = int32(0)
			v1344 = v1149 + (v1114 + int32(_a_F_des_init_13))
			*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1343
			v1347 = v1149 + (v1114 + int32(_a_F_des_init_14))
			*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1343
			if v1159 == v1343 {
				v1370 = v1343
				v1373 = int32(0)
				v1374 = v1370
			} else {
				v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[37]))))
				if v1354 == int32(255) {
					v1370 = v1343
					v1373 = int32(0)
					v1374 = v1370
				} else {
					v1358 = v1354 << (uint(int32(2)) % 32)
					if base.Ui32(v1354) <= base.Ui32(int32(23)) {
						v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+uint32(_c_F_des_init[51])))
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1363
						v1373 = v1363
						v1374 = v1343
					} else {
						v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1358+int32(_a_F_des_init_15)-int32(96))))
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1367
						v1370 = v1367
						v1373 = int32(0)
						v1374 = v1370
					}
				}
			}
			if v1184 == int32(0) {
				v1397 = v1373
				v1398 = v1374
			} else {
				v1380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[52]))))
				if v1380 == int32(255) {
					v1397 = v1373
					v1398 = v1374
				} else {
					v1384 = v1380 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1380) {
						v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1384+int32(_a_F_des_init_15)-int32(96))))
						v1392 = v1374 | v1391
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1392
						v1397 = v1373
						v1398 = v1392
					} else {
						v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+uint32(_c_F_des_init[51])))
						v1395 = v1373 | v1394
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1395
						v1397 = v1395
						v1398 = v1374
					}
				}
			}
			if v1211 == int32(0) {
				v1422 = v1397
				v1423 = v1398
			} else {
				v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[53]))))
				if v1405 == int32(255) {
					v1422 = v1397
					v1423 = v1398
				} else {
					v1409 = v1405 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1405) {
						v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1409+int32(_a_F_des_init_15)-int32(96))))
						v1417 = v1398 | v1416
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1417
						v1422 = v1397
						v1423 = v1417
					} else {
						v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+uint32(_c_F_des_init[51])))
						v1420 = v1397 | v1419
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1420
						v1422 = v1420
						v1423 = v1398
					}
				}
			}
			if v1238 == int32(0) {
				v1447 = v1422
				v1448 = v1423
			} else {
				v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[54]))))
				if v1430 == int32(255) {
					v1447 = v1422
					v1448 = v1423
				} else {
					v1434 = v1430 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1430) {
						v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1434+int32(_a_F_des_init_15)-int32(96))))
						v1442 = v1423 | v1441
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1442
						v1447 = v1422
						v1448 = v1442
					} else {
						v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+uint32(_c_F_des_init[51])))
						v1445 = v1422 | v1444
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1445
						v1447 = v1445
						v1448 = v1423
					}
				}
			}
			if v1265 == int32(0) {
				v1472 = v1447
				v1473 = v1448
			} else {
				v1455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[55]))))
				if v1455 == int32(255) {
					v1472 = v1447
					v1473 = v1448
				} else {
					v1459 = v1455 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1455) {
						v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1459+int32(_a_F_des_init_15)-int32(96))))
						v1467 = v1448 | v1466
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1467
						v1472 = v1447
						v1473 = v1467
					} else {
						v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+uint32(_c_F_des_init[51])))
						v1470 = v1447 | v1469
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1470
						v1472 = v1470
						v1473 = v1448
					}
				}
			}
			if v1292 == int32(0) {
				v1497 = v1472
				v1498 = v1473
			} else {
				v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[56]))))
				if v1480 == int32(255) {
					v1497 = v1472
					v1498 = v1473
				} else {
					v1484 = v1480 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1480) {
						v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1484+int32(_a_F_des_init_15)-int32(96))))
						v1492 = v1473 | v1491
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1492
						v1497 = v1472
						v1498 = v1492
					} else {
						v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1484)+uint32(_c_F_des_init[51])))
						v1495 = v1472 | v1494
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1495
						v1497 = v1495
						v1498 = v1473
					}
				}
			}
			if v1319 == int32(0) {
			} else {
				v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+uint32(_c_F_des_init[57]))))
				if v1505 == int32(255) {
				} else {
					v1509 = v1505 << (uint(int32(2)) % 32)
					if base.Ui32(int32(24)) <= base.Ui32(v1505) {
						v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1509+int32(_a_F_des_init_15)-int32(96))))
						*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1498 | v1516
					} else {
						v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+uint32(_c_F_des_init[51])))
						*(*int32)(unsafe.Add(mBase, uint32(v1344))) = v1497 | v1519
					}
				}
			}
			v1525 = v1126 + int32(1)
			if v1525 != int32(128) {
				v1126 = v1525
				continue
			} else {
				break
			}
			break
		}
		v1529 = v967 + int32(1)
		if v1529 != int32(8) {
			v967 = v1529
			continue
		} else {
			break
		}
		break
	}
	v1533 = int32(0)
	for {
		v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1533)+uint32(_c_F_des_init[58]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1557)+uint32(_c_F_des_init[59]))) = uint8(v1533)
		v1562 = v1533 | int32(1)
		v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1562)+uint32(_c_F_des_init[58]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1565)+uint32(_c_F_des_init[59]))) = uint8(v1562)
		v1570 = v1533 | int32(2)
		v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1570)+uint32(_c_F_des_init[58]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1573)+uint32(_c_F_des_init[59]))) = uint8(v1570)
		v1578 = v1533 | int32(3)
		v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578)+uint32(_c_F_des_init[58]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v1581)+uint32(_c_F_des_init[59]))) = uint8(v1578)
		v1586 = v1533 + int32(4)
		if v1586 != int32(32) {
			v1533 = v1586
			continue
		} else {
			break
		}
		break
	}
	v1594 = int32(0)
	for {
		v1617 = v1594 << (uint(int32(3)) % 32)
		v1619 = int32(0)
		for {
			v1643 = v1594<<(uint(int32(10))%32) + int32(_a_F_des_init_16) + v1619<<(uint(int32(2))%32)
			v1644 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1644
			if v1619&int32(128) != 0 {
				v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[60]))))
				v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1651<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1656
				v1658 = v1656
			} else {
				v1658 = v1644
			}
			if v1619&int32(64) != 0 {
				v1663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[61]))))
				v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1663<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1669 = v1658 | v1668
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1669
				v1671 = v1669
			} else {
				v1671 = v1658
			}
			if v1619&int32(32) != 0 {
				v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[62]))))
				v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1676<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1682 = v1671 | v1681
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1682
				v1684 = v1682
			} else {
				v1684 = v1671
			}
			if v1619&int32(16) != 0 {
				v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[63]))))
				v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1689<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1695 = v1684 | v1694
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1695
				v1697 = v1695
			} else {
				v1697 = v1684
			}
			if v1619&int32(8) != 0 {
				v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[64]))))
				v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1702<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1708 = v1697 | v1707
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1708
				v1710 = v1708
			} else {
				v1710 = v1697
			}
			if v1619&int32(4) != 0 {
				v1715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[65]))))
				v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1715<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1721 = v1710 | v1720
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1721
				v1723 = v1721
			} else {
				v1723 = v1710
			}
			if v1619&int32(2) != 0 {
				v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[66]))))
				v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1728<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				v1734 = v1723 | v1733
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1734
				v1736 = v1734
			} else {
				v1736 = v1723
			}
			if v1619&int32(1) != 0 {
				v1741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1617)+uint32(_c_F_des_init[67]))))
				v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1741<<(uint(int32(2))%32))+uint32(_c_F_des_init[43])))
				*(*int32)(unsafe.Add(mBase, uint32(v1643))) = v1736 | v1746
			} else {
			}
			v1750 = v1619 + int32(1)
			if v1750 != int32(256) {
				v1619 = v1750
				continue
			} else {
				break
			}
			break
		}
		v1754 = v1594 + int32(1)
		if v1754 != int32(4) {
			v1594 = v1754
			continue
		} else {
			break
		}
		break
	}
	v1758 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_des_init[68])) = uint8(v1758)
	return
}
func F_discard_stack_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 != int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v38
	if v37 == v38 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	if v8 == v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v8 == v14 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v8 == v16 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v8 == v18 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v25 = l0 + int32(56)
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_pfree(m, v8)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v8 == v27 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v8 != v29 {
		v25 = v26
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	return
L15:
	;
	goto L1
L16:
	;
	return
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v37 == v42 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v37 == v44 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v49 = l0 + int32(56)
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_pfree(m, v37)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L27
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	if v37 == v53 {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
	if v37 != v55 {
		v49 = v52
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	goto L16
}
func F_dispose_chunk(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v57 int32
	_ = v57
	var __phi57 int32
	_ = __phi57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var __phi215 int32
	_ = __phi215
	var v217 int32
	_ = v217
	var __phi217 int32
	_ = __phi217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v411 int32
	_ = v411
	v10 = l0 + l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11&int32(1) != 0 {
		v128 = l0
		v129 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v137&int32(2) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	if v11&int32(2) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = v18 + l1
	v20 = l0 - v18
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v20 != v22 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if v38 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L27
	}
L6:
	;
	v90 = int32(0)
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v27
	v128 = v20
	v129 = v19
	goto L2
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(v18) <= base.Ui32(int32(255)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v71 = int32(3)
	if v70&v71 != v71 {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L26
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v24 != v27 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v20 != v24 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v29 = int32(_a_F_dispose_chunk_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v31 & base.I32_rotl(int32(-2), int32(base.Ui32(v18)>>(uint(int32(3))%32)))
	v128 = v20
	v129 = v19
	goto L2
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v40
	v90 = v24
	goto L5
L16:
	;
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v43
	v52 = v20 + int32(20)
	goto L20
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v46 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	__phi55 = v52
	__phi57 = v51
	v55 = __phi55
	v57 = __phi57
	goto L22
L21:
	;
	v51 = v46
	v52 = v20 + int32(16)
	goto L20
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v64 != 0 {
		__phi55 = v57 + int32(20)
		__phi57 = v64
		v55 = __phi55
		v57 = __phi57
		goto L22
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(0)
	v90 = v57
	goto L5
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v67 != 0 {
		__phi55 = v57 + int32(16)
		__phi57 = v67
		v55 = __phi55
		v57 = __phi57
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v70 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v19 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	return
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v100 = v98 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_dispose_chunk[3])))
	if v101 == v20 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v38
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v120 != 0 {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_dispose_chunk[3]))) = v90
	if v90 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v20 == v113 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v106 = int32(_a_F_dispose_chunk_1)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v108 & base.I32_rotl(int32(-2), v98)
	v128 = v20
	v129 = v19
	goto L2
L33:
	;
	if v90 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L37
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v90
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v90
	goto L33
L37:
	;
	goto L28
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v90
	goto L40
L39:
	;
	goto L40
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v123 == int32(0) {
		v128 = v20
		v129 = v19
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v123)+24)) = v90
	v128 = v20
	v129 = v19
	goto L2
L42:
	;
	if base.Ui32(v298) <= base.Ui32(int32(255)) {
		goto L89
	} else {
		goto L90
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v181 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v181))) = v181
	if v128 != v165 {
		v298 = v181
		goto L42
	} else {
		goto L88
	}
L44:
	;
	if v198 == int32(0) {
		goto L43
	} else {
		goto L73
	}
L45:
	;
	v242 = int32(0)
	goto L44
L46:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[5]))
	if v143 == v10 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v137 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v129 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v129))) = v129
	v298 = v129
	goto L42
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[5])) = v128
	v147 = int32(_a_F_dispose_chunk_2)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[6]))
	v150 = v149 + v129
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[6])) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v150 | int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v128 != v156 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0]))
	if v165 == v10 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0])) = v159
	return
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[0])) = v128
	v169 = int32(_a_F_dispose_chunk_3)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2]))
	v172 = v171 + v129
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v172 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128+v172))) = v172
	return
L54:
	;
	goto L55
L55:
	;
	v181 = v137&int32(-8) + v129
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if base.Ui32(v137) <= base.Ui32(int32(255)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v185 == v182 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v182 != v10 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v187 = int32(_a_F_dispose_chunk_0)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v189 & base.I32_rotl(int32(-2), int32(base.Ui32(v137)>>(uint(int32(3))%32)))
	goto L43
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v185
	goto L43
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v200
	v242 = v182
	goto L44
L63:
	;
	goto L64
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v203 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v211 = v203
	v212 = v10 + int32(20)
	goto L67
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v206 == int32(0) {
		goto L45
	} else {
		goto L68
	}
L67:
	;
	__phi215 = v212
	__phi217 = v211
	v215 = __phi215
	v217 = __phi217
	goto L69
L68:
	;
	v211 = v206
	v212 = v10 + int32(16)
	goto L67
L69:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	if v224 != 0 {
		__phi215 = v217 + int32(20)
		__phi217 = v224
		v215 = __phi215
		v217 = __phi217
		goto L69
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = int32(0)
	v242 = v217
	goto L44
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	if v227 != 0 {
		__phi215 = v217 + int32(16)
		__phi217 = v227
		v215 = __phi215
		v217 = __phi217
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v252 = v250 << (uint(int32(2)) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_dispose_chunk[3])))
	if v253 == v10 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = v198
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v272 != 0 {
		goto L84
	} else {
		goto L85
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_dispose_chunk[3]))) = v242
	if v242 != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	if v10 == v265 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v258 = int32(_a_F_dispose_chunk_1)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v260 & base.I32_rotl(int32(-2), v250)
	goto L43
L79:
	;
	if v242 == int32(0) {
		goto L43
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v242
	goto L79
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+20)) = v242
	goto L79
L83:
	;
	goto L74
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+16)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v272)+24)) = v242
	goto L86
L85:
	;
	goto L86
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v275 == int32(0) {
		goto L43
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v275)+24)) = v242
	goto L43
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[2])) = v181
	return
L89:
	;
	v309 = v298 & int32(248)
	v311 = v309 + int32(_a_F_dispose_chunk_4)
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1]))
	v317 = int32(1) << (uint(int32(base.Ui32(v298)>>(uint(int32(3))%32))) % 32)
	if v313&v317 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v298) <= base.Ui32(int32(16777215)) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+uint32(_c_F_dispose_chunk[7]))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v325)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v325
	return
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[1])) = v317 | v313
	v325 = v311
	goto L92
L94:
	;
	goto L95
L95:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v309)+uint32(_c_F_dispose_chunk[7])))
	v325 = v324
	goto L92
L96:
	;
	v336 = base.I32_clz(int32(base.Ui32(v298) >> (uint(int32(8)) % 32)))
	v339 = int32(1)
	v347 = int32(base.Ui32(v298)>>(uint(int32(38)-v336)%32))&v339 | v336<<(uint(v339)%32) ^ int32(62)
	goto L98
L97:
	;
	v347 = int32(31)
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+28)) = v347
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = int64(0)
	v352 = v347 << (uint(int32(2)) % 32)
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4]))
	v358 = int32(1) << (uint(v347) % 32)
	if v356&v358 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v411
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v128
	return
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dispose_chunk[4])) = v356 | v358
	*(*int32)(unsafe.Add(mBase, uint32(v352)+uint32(_c_F_dispose_chunk[3]))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v352 + int32(_a_F_dispose_chunk_5)
	goto L100
L102:
	;
	goto L103
L103:
	;
	if v347 != int32(31) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v374 = int32(25) - int32(base.Ui32(v347)>>(uint(int32(1))%32))
	goto L106
L105:
	;
	v374 = int32(0)
	goto L106
L106:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v352)+uint32(_c_F_dispose_chunk[3])))
	v379 = v376
	v380 = v298 << (uint(v374) % 32)
	goto L107
L107:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v386&int32(-8) == v298 {
		goto L99
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+16)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = v379
	goto L100
L109:
	;
	v396 = v379 + int32(base.Ui32(v380)>>(uint(int32(29))%32))&int32(4)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
	if v397 != 0 {
		v379 = v397
		v380 = v380 << (uint(int32(1)) % 32)
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
}
func F_dlgamma(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v9 float64
	_ = v9
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 float64
	_ = v60
	var v68 float64
	_ = v68
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v123 float64
	_ = v123
	var v139 float64
	_ = v139
	var v145 float64
	_ = v145
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v200 float64
	_ = v200
	var v217 float64
	_ = v217
	var v226 float64
	_ = v226
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v233 float64
	_ = v233
	var v249 float64
	_ = v249
	var v250 float64
	_ = v250
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v266 float64
	_ = v266
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v405 float64
	_ = v405
	var v406 int32
	_ = v406
	var v408 float64
	_ = v408
	var v451 float64
	_ = v451
	var v456 float64
	_ = v456
	var v460 float64
	_ = v460
	var v464 float64
	_ = v464
	var v468 float64
	_ = v468
	var v472 float64
	_ = v472
	var v474 float64
	_ = v474
	var v483 float64
	_ = v483
	var v484 float64
	_ = v484
	var v510 float64
	_ = v510
	var v515 float64
	_ = v515
	var v522 float64
	_ = v522
	var v524 float64
	_ = v524
	var v531 int32
	_ = v531
	var v540 int32
	_ = v540
	v2 = float64(0)
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_dlgamma[0])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_dlgamma[1])) = int32(1)
	v17 = base.I64_reinterpret_f64(v9)
	v22 = base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v22) {
		v522 = base.F64_mul(v9, v9)
	} else {
		if base.Ui32(v22) <= base.Ui32(int32(999292927)) {
			if v17 < int64(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_dlgamma[1])) = int32(-1)
				v33 = base.F64_neg(v9)
			} else {
				v33 = v9
			}
			v34 = F_log(m, v33)
			mBase = m.M
			v522 = base.F64_neg(v34)
		} else {
			if v17 < int64(0) {
				v41 = base.F64_neg(v9)
				v43 = base.F64_mul(v41, float64(0.5))
				v45 = base.F64_sub(v43, base.F64_floor(v43))
				v46 = base.F64_add(v45, v45)
				v50 = int32(1)
				v53 = base.I32_div_s(base.I32_trunc_sat_f64_s(base.F64_mul(v46, float64(4)))+v50, int32(2))
				v60 = base.F64_mul(base.F64_add(v46, base.F64_promote_f32(base.F32_mul(base.F32_convert_i32_s(v53), float32(-0.5)))), float64(3.141592653589793))
				switch v53 - v50 {
				case 0:
					v107 = float64(1)
					v108 = base.F64_mul(v60, v60)
					v110 = base.F64_mul(v108, float64(0.5))
					v111 = base.F64_sub(v107, v110)
					v123 = base.F64_mul(v108, v108)
					v217 = base.F64_add(v111, base.F64_add(base.F64_sub(base.F64_sub(v107, v111), v110), base.F64_sub(base.F64_mul(v108, base.F64_add(base.F64_mul(v108, base.F64_add(base.F64_mul(v108, base.F64_add(base.F64_mul(v108, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v123, v123), base.F64_add(base.F64_mul(v108, base.F64_add(base.F64_mul(v108, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v60, float64(0)))))
				case 1:
					v139 = base.F64_neg(v60)
					v145 = base.F64_mul(v139, v139)
					v217 = base.F64_add(base.F64_mul(base.F64_mul(v139, v145), base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(base.F64_mul(v145, base.F64_mul(v145, v145)), base.F64_add(base.F64_mul(v145, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v139)
				case 2:
					v184 = float64(1)
					v185 = base.F64_mul(v60, v60)
					v187 = base.F64_mul(v185, float64(0.5))
					v188 = base.F64_sub(v184, v187)
					v200 = base.F64_mul(v185, v185)
					v217 = base.F64_neg(base.F64_add(v188, base.F64_add(base.F64_sub(base.F64_sub(v184, v188), v187), base.F64_sub(base.F64_mul(v185, base.F64_add(base.F64_mul(v185, base.F64_add(base.F64_mul(v185, base.F64_add(base.F64_mul(v185, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v200, v200), base.F64_add(base.F64_mul(v185, base.F64_add(base.F64_mul(v185, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v60, float64(0))))))
				default:
					v68 = base.F64_mul(v60, v60)
					v217 = base.F64_add(base.F64_mul(base.F64_mul(v60, v68), base.F64_add(base.F64_mul(v68, base.F64_add(base.F64_mul(base.F64_mul(v68, base.F64_mul(v68, v68)), base.F64_add(base.F64_mul(v68, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v68, base.F64_add(base.F64_mul(v68, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), v60)
				}
				if base.F64_eq(v217, float64(0)) != 0 {
					v522 = base.F64_div(float64(1), base.F64_sub(v9, v9))
				} else {
					if base.F64_gt(v217, float64(0)) != 0 {
						*(*int32)(unsafe.Add(mBase, _c_F_dlgamma[1])) = int32(-1)
						v226 = v217
					} else {
						v226 = base.F64_neg(v217)
					}
					v230 = F_log(m, base.F64_div(float64(3.141592653589793), base.F64_mul(v226, v41)))
					mBase = m.M
					v231 = v41
					v233 = v230
					if base.I32_wrap_i64(v17) == int32(0) {
						if base.B2i32(v22 == int32(1072693248))|base.B2i32(v22 == int32(1073741824)) != 0 {
							v510 = float64(0)
						} else {
							if base.Ui32(v22) <= base.Ui32(int32(1073741823)) {
								if base.Ui32(v22) <= base.Ui32(int32(1072483532)) {
									v249 = F_log(m, v231)
									mBase = m.M
									v250 = base.F64_neg(v249)
									if base.Ui32(int32(1072130371)) < base.Ui32(v22) {
										v261 = v250
										v262 = float64(1)
										v263 = base.F64_sub(v262, v231)
										v266 = base.F64_mul(v263, v263)
										v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
									} else {
										if base.Ui32(v22) <= base.Ui32(int32(1070442080)) {
											v363 = v231
											v364 = v250
											v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
										} else {
											v308 = v250
											v309 = base.F64_add(v231, float64(-0.46163214496836225))
											v310 = base.F64_mul(v309, v309)
											v311 = base.F64_mul(v309, v310)
											v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
										}
									}
								} else {
									if base.Ui32(v22) <= base.Ui32(int32(1073460418)) {
										if base.Ui32(v22) < base.Ui32(int32(1072936132)) {
											v363 = base.F64_add(v231, float64(-1))
											v364 = v2
											v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
										} else {
											v308 = v2
											v309 = base.F64_add(v231, float64(-1.4616321449683622))
											v310 = base.F64_mul(v309, v309)
											v311 = base.F64_mul(v309, v310)
											v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
										}
									} else {
										v261 = v2
										v262 = float64(2)
										v263 = base.F64_sub(v262, v231)
										v266 = base.F64_mul(v263, v263)
										v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
									}
								}
							} else {
								if base.Ui32(v22) <= base.Ui32(int32(1075838975)) {
									v405 = float64(1)
									v406 = base.I32_trunc_sat_f64_s(v231)
									v408 = base.F64_sub(v231, base.F64_convert_i32_s(v406))
									v451 = base.F64_add(base.F64_mul(v408, float64(0.5)), base.F64_div(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), v405)))
									switch v406 - int32(3) {
									case 0:
										v468 = v405
										v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
										mBase = m.M
										v510 = base.F64_add(v451, v472)
									case 1:
										v464 = v405
										v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
										v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
										mBase = m.M
										v510 = base.F64_add(v451, v472)
									case 2:
										v460 = v405
										v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
										v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
										v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
										mBase = m.M
										v510 = base.F64_add(v451, v472)
									case 3:
										v456 = v405
										v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
										v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
										v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
										v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
										mBase = m.M
										v510 = base.F64_add(v451, v472)
									case 4:
										v456 = base.F64_add(v408, float64(6))
										v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
										v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
										v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
										v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
										mBase = m.M
										v510 = base.F64_add(v451, v472)
									default:
										v510 = v451
									}
								} else {
									v474 = F_log(m, v231)
									mBase = m.M
									if base.Ui32(v22) <= base.Ui32(int32(1133510655)) {
										v483 = base.F64_div(float64(1), v231)
										v484 = base.F64_mul(v483, v483)
										v510 = base.F64_add(base.F64_mul(base.F64_add(v231, float64(-0.5)), base.F64_add(v474, float64(-1))), base.F64_add(base.F64_mul(v483, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
									} else {
										v510 = base.F64_mul(v231, base.F64_add(v474, float64(-1)))
									}
								}
							}
						}
					} else {
						if base.Ui32(v22) <= base.Ui32(int32(1073741823)) {
							if base.Ui32(v22) <= base.Ui32(int32(1072483532)) {
								v249 = F_log(m, v231)
								mBase = m.M
								v250 = base.F64_neg(v249)
								if base.Ui32(int32(1072130371)) < base.Ui32(v22) {
									v261 = v250
									v262 = float64(1)
									v263 = base.F64_sub(v262, v231)
									v266 = base.F64_mul(v263, v263)
									v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								} else {
									if base.Ui32(v22) <= base.Ui32(int32(1070442080)) {
										v363 = v231
										v364 = v250
										v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v308 = v250
										v309 = base.F64_add(v231, float64(-0.46163214496836225))
										v310 = base.F64_mul(v309, v309)
										v311 = base.F64_mul(v309, v310)
										v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								}
							} else {
								if base.Ui32(v22) <= base.Ui32(int32(1073460418)) {
									if base.Ui32(v22) < base.Ui32(int32(1072936132)) {
										v363 = base.F64_add(v231, float64(-1))
										v364 = v2
										v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v308 = v2
										v309 = base.F64_add(v231, float64(-1.4616321449683622))
										v310 = base.F64_mul(v309, v309)
										v311 = base.F64_mul(v309, v310)
										v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								} else {
									v261 = v2
									v262 = float64(2)
									v263 = base.F64_sub(v262, v231)
									v266 = base.F64_mul(v263, v263)
									v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								}
							}
						} else {
							if base.Ui32(v22) <= base.Ui32(int32(1075838975)) {
								v405 = float64(1)
								v406 = base.I32_trunc_sat_f64_s(v231)
								v408 = base.F64_sub(v231, base.F64_convert_i32_s(v406))
								v451 = base.F64_add(base.F64_mul(v408, float64(0.5)), base.F64_div(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), v405)))
								switch v406 - int32(3) {
								case 0:
									v468 = v405
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 1:
									v464 = v405
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 2:
									v460 = v405
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 3:
									v456 = v405
									v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 4:
									v456 = base.F64_add(v408, float64(6))
									v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								default:
									v510 = v451
								}
							} else {
								v474 = F_log(m, v231)
								mBase = m.M
								if base.Ui32(v22) <= base.Ui32(int32(1133510655)) {
									v483 = base.F64_div(float64(1), v231)
									v484 = base.F64_mul(v483, v483)
									v510 = base.F64_add(base.F64_mul(base.F64_add(v231, float64(-0.5)), base.F64_add(v474, float64(-1))), base.F64_add(base.F64_mul(v483, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
								} else {
									v510 = base.F64_mul(v231, base.F64_add(v474, float64(-1)))
								}
							}
						}
					}
					if int64(0) <= v17 {
						v515 = v510
					} else {
						v515 = base.F64_sub(v233, v510)
					}
					v522 = v515
				}
			} else {
				v231 = v9
				v233 = v2
				if base.I32_wrap_i64(v17) == int32(0) {
					if base.B2i32(v22 == int32(1072693248))|base.B2i32(v22 == int32(1073741824)) != 0 {
						v510 = float64(0)
					} else {
						if base.Ui32(v22) <= base.Ui32(int32(1073741823)) {
							if base.Ui32(v22) <= base.Ui32(int32(1072483532)) {
								v249 = F_log(m, v231)
								mBase = m.M
								v250 = base.F64_neg(v249)
								if base.Ui32(int32(1072130371)) < base.Ui32(v22) {
									v261 = v250
									v262 = float64(1)
									v263 = base.F64_sub(v262, v231)
									v266 = base.F64_mul(v263, v263)
									v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								} else {
									if base.Ui32(v22) <= base.Ui32(int32(1070442080)) {
										v363 = v231
										v364 = v250
										v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v308 = v250
										v309 = base.F64_add(v231, float64(-0.46163214496836225))
										v310 = base.F64_mul(v309, v309)
										v311 = base.F64_mul(v309, v310)
										v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								}
							} else {
								if base.Ui32(v22) <= base.Ui32(int32(1073460418)) {
									if base.Ui32(v22) < base.Ui32(int32(1072936132)) {
										v363 = base.F64_add(v231, float64(-1))
										v364 = v2
										v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
									} else {
										v308 = v2
										v309 = base.F64_add(v231, float64(-1.4616321449683622))
										v310 = base.F64_mul(v309, v309)
										v311 = base.F64_mul(v309, v310)
										v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
									}
								} else {
									v261 = v2
									v262 = float64(2)
									v263 = base.F64_sub(v262, v231)
									v266 = base.F64_mul(v263, v263)
									v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
								}
							}
						} else {
							if base.Ui32(v22) <= base.Ui32(int32(1075838975)) {
								v405 = float64(1)
								v406 = base.I32_trunc_sat_f64_s(v231)
								v408 = base.F64_sub(v231, base.F64_convert_i32_s(v406))
								v451 = base.F64_add(base.F64_mul(v408, float64(0.5)), base.F64_div(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), v405)))
								switch v406 - int32(3) {
								case 0:
									v468 = v405
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 1:
									v464 = v405
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 2:
									v460 = v405
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 3:
									v456 = v405
									v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								case 4:
									v456 = base.F64_add(v408, float64(6))
									v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
									v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
									v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
									v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
									mBase = m.M
									v510 = base.F64_add(v451, v472)
								default:
									v510 = v451
								}
							} else {
								v474 = F_log(m, v231)
								mBase = m.M
								if base.Ui32(v22) <= base.Ui32(int32(1133510655)) {
									v483 = base.F64_div(float64(1), v231)
									v484 = base.F64_mul(v483, v483)
									v510 = base.F64_add(base.F64_mul(base.F64_add(v231, float64(-0.5)), base.F64_add(v474, float64(-1))), base.F64_add(base.F64_mul(v483, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
								} else {
									v510 = base.F64_mul(v231, base.F64_add(v474, float64(-1)))
								}
							}
						}
					}
				} else {
					if base.Ui32(v22) <= base.Ui32(int32(1073741823)) {
						if base.Ui32(v22) <= base.Ui32(int32(1072483532)) {
							v249 = F_log(m, v231)
							mBase = m.M
							v250 = base.F64_neg(v249)
							if base.Ui32(int32(1072130371)) < base.Ui32(v22) {
								v261 = v250
								v262 = float64(1)
								v263 = base.F64_sub(v262, v231)
								v266 = base.F64_mul(v263, v263)
								v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
							} else {
								if base.Ui32(v22) <= base.Ui32(int32(1070442080)) {
									v363 = v231
									v364 = v250
									v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
								} else {
									v308 = v250
									v309 = base.F64_add(v231, float64(-0.46163214496836225))
									v310 = base.F64_mul(v309, v309)
									v311 = base.F64_mul(v309, v310)
									v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
								}
							}
						} else {
							if base.Ui32(v22) <= base.Ui32(int32(1073460418)) {
								if base.Ui32(v22) < base.Ui32(int32(1072936132)) {
									v363 = base.F64_add(v231, float64(-1))
									v364 = v2
									v510 = base.F64_add(v364, base.F64_add(base.F64_mul(v363, float64(-0.5)), base.F64_div(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.013381091853678766)), float64(0.22896372806469245))), float64(0.9777175279633727))), float64(1.4549225013723477))), float64(0.6328270640250934))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, base.F64_add(base.F64_mul(v363, float64(0.003217092422824239)), float64(0.10422264559336913))), float64(0.7692851504566728))), float64(2.128489763798934))), float64(2.4559779371304113))), float64(1)))))
								} else {
									v308 = v2
									v309 = base.F64_add(v231, float64(-1.4616321449683622))
									v310 = base.F64_mul(v309, v309)
									v311 = base.F64_mul(v309, v310)
									v510 = base.F64_add(v308, base.F64_add(base.F64_sub(base.F64_mul(v310, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.00031563207090362595)), float64(-0.0014034646998923284))), float64(0.006100538702462913))), float64(-0.032788541075985965))), float64(0.48383612272381005))), base.F64_sub(float64(-3.638676997039505e-18), base.F64_mul(v311, base.F64_add(base.F64_mul(v309, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(0.0003355291926355191)), float64(-0.0005385953053567405))), float64(0.0022596478090061247))), float64(-0.010314224129834144))), float64(0.06462494023913339))), base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, base.F64_add(base.F64_mul(v311, float64(-0.00031275416837512086)), float64(0.000881081882437654))), float64(-0.0036845201678113826))), float64(0.01797067508118204))), float64(-0.1475877229945939)))))), float64(-0.12148629053584961)))
								}
							} else {
								v261 = v2
								v262 = float64(2)
								v263 = base.F64_sub(v262, v231)
								v266 = base.F64_mul(v263, v263)
								v510 = base.F64_add(v261, base.F64_add(base.F64_mul(v263, float64(-0.5)), base.F64_add(base.F64_mul(v263, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(2.5214456545125733e-05)), float64(0.00022086279071390839))), float64(0.0011927076318336207))), float64(0.007385550860814029))), float64(0.06735230105312927))), float64(0.07721566490153287))), base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, base.F64_add(base.F64_mul(v266, float64(4.4864094961891516e-05)), float64(0.00010801156724758394))), float64(0.0005100697921535113))), float64(0.0028905138367341563))), float64(0.020580808432516733))), float64(0.3224670334241136))))))
							}
						}
					} else {
						if base.Ui32(v22) <= base.Ui32(int32(1075838975)) {
							v405 = float64(1)
							v406 = base.I32_trunc_sat_f64_s(v231)
							v408 = base.F64_sub(v231, base.F64_convert_i32_s(v406))
							v451 = base.F64_add(base.F64_mul(v408, float64(0.5)), base.F64_div(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(3.194753265841009e-05)), float64(0.0018402845140733772))), float64(0.02664227030336386))), float64(0.14635047265246445))), float64(0.325778796408931))), float64(0.21498241596060885))), float64(-0.07721566490153287))), base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, base.F64_add(base.F64_mul(v408, float64(7.326684307446256e-06)), float64(0.0007779424963818936))), float64(0.01864591917156529))), float64(0.17193386563280308))), float64(0.7219355475671381))), float64(1.3920053346762105))), v405)))
							switch v406 - int32(3) {
							case 0:
								v468 = v405
								v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
								mBase = m.M
								v510 = base.F64_add(v451, v472)
							case 1:
								v464 = v405
								v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
								v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
								mBase = m.M
								v510 = base.F64_add(v451, v472)
							case 2:
								v460 = v405
								v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
								v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
								v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
								mBase = m.M
								v510 = base.F64_add(v451, v472)
							case 3:
								v456 = v405
								v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
								v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
								v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
								v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
								mBase = m.M
								v510 = base.F64_add(v451, v472)
							case 4:
								v456 = base.F64_add(v408, float64(6))
								v460 = base.F64_mul(base.F64_add(v408, float64(5)), v456)
								v464 = base.F64_mul(base.F64_add(v408, float64(4)), v460)
								v468 = base.F64_mul(base.F64_add(v408, float64(3)), v464)
								v472 = F_log(m, base.F64_mul(base.F64_add(v408, float64(2)), v468))
								mBase = m.M
								v510 = base.F64_add(v451, v472)
							default:
								v510 = v451
							}
						} else {
							v474 = F_log(m, v231)
							mBase = m.M
							if base.Ui32(v22) <= base.Ui32(int32(1133510655)) {
								v483 = base.F64_div(float64(1), v231)
								v484 = base.F64_mul(v483, v483)
								v510 = base.F64_add(base.F64_mul(base.F64_add(v231, float64(-0.5)), base.F64_add(v474, float64(-1))), base.F64_add(base.F64_mul(v483, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, base.F64_add(base.F64_mul(v484, float64(-0.0016309293409657527)), float64(0.0008363399189962821))), float64(-0.00059518755745034))), float64(0.0007936505586430196))), float64(-0.0027777777772877554))), float64(0.08333333333333297))), float64(0.4189385332046727)))
							} else {
								v510 = base.F64_mul(v231, base.F64_add(v474, float64(-1)))
							}
						}
					}
				}
				if int64(0) <= v17 {
					v515 = v510
				} else {
					v515 = base.F64_sub(v233, v510)
				}
				v522 = v515
			}
		}
	}
	v524 = math.Float64frombits(uint64(0x7ff0000000000000))
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_dlgamma[0]))
	if (base.F64_eq(base.F64_abs(v9), v524)|base.F64_ne(base.F64_abs(v522), v524))&base.B2i32(v531 != int32(68)) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v540 = m.ExcPending
		if v540 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I64_reinterpret_f64(v522)
	}
}
func F_dlog1(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_ne(v4, float64(0)) != 0 {
		if base.F64_lt(v4, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352583810))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_dlog1_0), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_dlog1_1), int32(1747), int32(_a_F_dlog1_2))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v9 = F_log(m, v4)
			mBase = m.M
			v11 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v9), v11)&base.F64_ne(base.F64_abs(v4), v11) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v9, float64(0))&base.F64_ne(v4, float64(1)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return base.I64_reinterpret_f64(v9)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(352583810))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_dlog1_3), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_dlog1_1), int32(1743), int32(_a_F_dlog1_2))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
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
func F_dlog10(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v19 int64
	_ = v19
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v117 float64
	_ = v117
	var v129 float64
	_ = v129
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_ne(v4, float64(0)) != 0 {
		if base.F64_lt(v4, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v187 = m.ExcPending
			if v187 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352583810))
				mBase = m.M
				v190 = m.ExcPending
				if v190 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_dlog10_0), int32(0))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_dlog10_1), int32(1780), int32(_a_F_dlog10_2))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v19 = base.I64_reinterpret_f64(v4)
			if v19 <= int64(4503599627370495) {
				if base.F64_eq(v4, float64(0)) != 0 {
					v151 = base.F64_div(float64(-1), base.F64_mul(v4, v4))
				} else {
					if int64(0) <= v19 {
						v46 = base.I64_reinterpret_f64(base.F64_mul(v4, float64(1.8014398509481984e+16)))
						v50 = v46
						v52 = int32(-1077)
						v53 = base.I32_wrap_i64(int64(base.Ui64(v46) >> (uint(int64(32)) % 64)))
						v55 = v53 + int32(_a_F_dlog10_3)
						v59 = base.F64_convert_i32_s(int32(base.Ui32(v55)>>(uint(int32(20))%32)) + v52)
						v61 = base.F64_mul(v59, float64(0.30102999566361177))
						v74 = base.F64_add(base.F64_reinterpret_i64(v50&int64(4294967295)|base.I64_extend_i32_u(v55&int32(_a_F_dlog10_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v77 = base.F64_mul(v74, base.F64_mul(v74, float64(0.5)))
						v82 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v74, v77)) & int64(-4294967296))
						v83 = float64(0.4342944818781689)
						v84 = base.F64_mul(v82, v83)
						v85 = base.F64_add(v61, v84)
						v90 = base.F64_div(v74, base.F64_add(v74, float64(2)))
						v91 = base.F64_mul(v90, v90)
						v92 = base.F64_mul(v91, v91)
						v117 = base.F64_add(base.F64_mul(v90, base.F64_add(v77, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v91, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v74, v82), v77))
						v129 = base.F64_add(v85, base.F64_add(base.F64_add(v84, base.F64_sub(v61, v85)), base.F64_add(base.F64_mul(v117, v83), base.F64_add(base.F64_mul(v59, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v117, v82), float64(2.5082946711645275e-11))))))
						v151 = v129
					} else {
						v151 = base.F64_div(base.F64_sub(v4, v4), float64(0))
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405311)) < base.Ui64(v19) {
					v129 = v4
					v151 = v129
				} else {
					v34 = int32(-1023)
					v36 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
					if v36 != int64(1072693248) {
						v50 = v19
						v52 = v34
						v53 = base.I32_wrap_i64(v36)
						v55 = v53 + int32(_a_F_dlog10_3)
						v59 = base.F64_convert_i32_s(int32(base.Ui32(v55)>>(uint(int32(20))%32)) + v52)
						v61 = base.F64_mul(v59, float64(0.30102999566361177))
						v74 = base.F64_add(base.F64_reinterpret_i64(v50&int64(4294967295)|base.I64_extend_i32_u(v55&int32(_a_F_dlog10_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
						v77 = base.F64_mul(v74, base.F64_mul(v74, float64(0.5)))
						v82 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v74, v77)) & int64(-4294967296))
						v83 = float64(0.4342944818781689)
						v84 = base.F64_mul(v82, v83)
						v85 = base.F64_add(v61, v84)
						v90 = base.F64_div(v74, base.F64_add(v74, float64(2)))
						v91 = base.F64_mul(v90, v90)
						v92 = base.F64_mul(v91, v91)
						v117 = base.F64_add(base.F64_mul(v90, base.F64_add(v77, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v91, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v74, v82), v77))
						v129 = base.F64_add(v85, base.F64_add(base.F64_add(v84, base.F64_sub(v61, v85)), base.F64_add(base.F64_mul(v117, v83), base.F64_add(base.F64_mul(v59, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v117, v82), float64(2.5082946711645275e-11))))))
						v151 = v129
					} else {
						if base.I32_wrap_i64(v19) != 0 {
							v50 = v19
							v52 = v34
							v53 = int32(1072693248)
							v55 = v53 + int32(_a_F_dlog10_3)
							v59 = base.F64_convert_i32_s(int32(base.Ui32(v55)>>(uint(int32(20))%32)) + v52)
							v61 = base.F64_mul(v59, float64(0.30102999566361177))
							v74 = base.F64_add(base.F64_reinterpret_i64(v50&int64(4294967295)|base.I64_extend_i32_u(v55&int32(_a_F_dlog10_4)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
							v77 = base.F64_mul(v74, base.F64_mul(v74, float64(0.5)))
							v82 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v74, v77)) & int64(-4294967296))
							v83 = float64(0.4342944818781689)
							v84 = base.F64_mul(v82, v83)
							v85 = base.F64_add(v61, v84)
							v90 = base.F64_div(v74, base.F64_add(v74, float64(2)))
							v91 = base.F64_mul(v90, v90)
							v92 = base.F64_mul(v91, v91)
							v117 = base.F64_add(base.F64_mul(v90, base.F64_add(v77, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v91, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v74, v82), v77))
							v129 = base.F64_add(v85, base.F64_add(base.F64_add(v84, base.F64_sub(v61, v85)), base.F64_add(base.F64_mul(v117, v83), base.F64_add(base.F64_mul(v59, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v117, v82), float64(2.5082946711645275e-11))))))
							v151 = v129
						} else {
							v151 = float64(0)
						}
					}
				}
			}
			v153 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v151), v153)&base.F64_ne(base.F64_abs(v4), v153) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v201 = m.ExcPending
				if v201 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v151, float64(0))&base.F64_ne(v4, float64(1)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v203 = m.ExcPending
					if v203 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return base.I64_reinterpret_f64(v151)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v171 = m.ExcPending
		if v171 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(352583810))
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_dlog10_5), int32(0))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_dlog10_1), int32(1776), int32(_a_F_dlog10_2))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
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
func F_doNegate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(72) {
		v37 = int32(0)
		v40 = F_makeSimpleA_Expr(m, v37, int32(_a_F_doNegate_0), v37, l0, l1)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			v42 = v40
			m.G0 = v6 + int32(16)
			return v42
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v12 - int32(473) {
		case 0:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0) - v16
			v42 = l0
			m.G0 = v6 + int32(16)
			return v42
		case 1:
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v23 = v19 + base.B2i32(v20 == int32(43))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			if v24 == int32(45) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23 + int32(1)
				v42 = l0
				m.G0 = v6 + int32(16)
				return v42
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
				v32 = F_psprintf(m, int32(_a_F_doNegate_1), v6)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32
					v42 = l0
					m.G0 = v6 + int32(16)
					return v42
				}
			}
		default:
			v37 = int32(0)
			v40 = F_makeSimpleA_Expr(m, v37, int32(_a_F_doNegate_0), v37, l0, l1)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = v40
				m.G0 = v6 + int32(16)
				return v42
			}
		}
	}
}
func F_do_getc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if int32(0) <= v4 {
		if v4 == int32(0) {
			v28 = l0 + int32(76)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			if v29 != 0 {
				v31 = v29
			} else {
				v31 = int32(1073741823)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v33 != v34 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 + int32(1)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
				v42 = v39
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
				return v42
			} else {
				v40 = F___uflow(m, l0)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
					return v42
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_do_getc[0]))
			if v10 != v4&int32(1073741823) {
				v28 = l0 + int32(76)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				if v29 != 0 {
					v31 = v29
				} else {
					v31 = int32(1073741823)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v33 != v34 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 + int32(1)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
					v42 = v39
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
					return v42
				} else {
					v40 = F___uflow(m, l0)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = v40
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
						return v42
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v14 != v15 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 + int32(1)
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
					return v20
				} else {
					v22 = F___uflow(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v22
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v14 != v15 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 + int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			return v20
		} else {
			v22 = F___uflow(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	}
}
func F_do_start_worker(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v87 int64
	_ = v87
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	v1 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[0]))
	v20 = F_LWLockAcquire(m, v16+int32(2816), int32(1))
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[1]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[2]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[3]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[0]))
	F_LWLockRelease(m, v32+int32(2816))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v30 - v28
	v38 = int32(0)
	if v38 < v37 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = v37
	goto L6
L5:
	;
	v41 = v38
	goto L6
L6:
	;
	if v41 < v26 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[4]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_do_start_worker_0), int32(0), int32(_a_F_do_start_worker_1), int32(_a_F_do_start_worker_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v407 = v1
	goto L9
L9:
	;
	return v407
L10:
	;
	v51 = int32(_a_F_do_start_worker_3)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[4])) = v49
	v55 = F_get_database_list(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v58 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = base.I32_wrap_i64(v58)
	*(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[5])) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[6]))
	v65 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[7])) = v65
	v68 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v73 = m.G0
	v74 = int32(16)
	v75 = v73 - v74
	m.G0 = v75
	F_gettimeofday(m, v75)
	mBase = m.M
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	v79 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75)+8)))
	m.G0 = v75 + v74
	v87 = v79 + v78*int64(1000000) - int64(946684800000000)
	goto L15
L15:
	;
	if v55 == int32(0) {
		v389 = v1
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[4])) = v52
	F_MemoryContextDelete(m, v49)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L87
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v90 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v300 != 0 {
		goto L75
	} else {
		goto L76
	}
L19:
	;
	v300 = int32(0)
	v306 = v1
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v68 == v65 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = int32(1)
	goto L24
L23:
	;
	v97 = v68 - v65
	goto L24
L24:
	;
	v98 = v60 - v63
	v99 = int32(3)
	if base.Ui32(v98) < base.Ui32(v99) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = v98 - v99
	goto L27
L26:
	;
	v103 = v98
	goto L27
L27:
	;
	v106 = int32(0)
	v108 = v106
	v110 = v106
	v111 = v1
	v115 = v1
	v116 = v1
	goto L28
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v115<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if base.B2i32(base.Ui32(v103) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v127) < base.Ui32(int32(3))) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v300 = v281
	v306 = v287
	goto L18
L30:
	;
	v295 = v115 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v295 < v296 {
		v108 = v293
		v110 = v281
		v111 = v282
		v115 = v295
		v116 = v287
		goto L28
	} else {
		goto L74
	}
L31:
	;
	if v110 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L32:
	;
	v137 = int32(1)
	if v108&v137 != 0 {
		v281 = v110
		v282 = v111
		v287 = v116
		v293 = v137
		goto L30
	} else {
		goto L38
	}
L33:
	;
	if int32(0) <= v127-v103 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(v127) < base.Ui32(v103) {
		goto L31
	} else {
		goto L37
	}
L36:
	;
	goto L31
L37:
	;
	goto L32
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	if v140+v97 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v281 = v110
	v282 = int32(1)
	v287 = v116
	v293 = int32(0)
	goto L30
L40:
	;
	if v110 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v111 != 0 {
		goto L39
	} else {
		goto L47
	}
L43:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if int32(0) <= v140-v144 {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v281 = v126
	v282 = int32(1)
	v287 = v116
	v293 = int32(0)
	goto L30
L46:
	;
	goto L45
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v151 = F_pgstat_fetch_stat_dbentry(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v151
	if v151 == int32(0) {
		v231 = v110
		v237 = v116
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v281 = v246
	v282 = v247
	v287 = v252
	v293 = int32(0)
	goto L30
L50:
	;
	v246 = v231
	v247 = int32(0)
	v252 = v237
	goto L49
L51:
	;
	v156 = int32(_a_F_do_start_worker_4)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[8]))
	if base.B2i32(v157 == int32(0))|base.B2i32(v157 == v156) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if v110 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v167 = v157
	goto L54
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v167-int32(20))))
	if v163 == v180 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L52
L56:
	;
	v183 = v167 - int32(12)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	goto L59
L57:
	;
	goto L58
L58:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v203 != int32(_a_F_do_start_worker_4) {
		v167 = v203
		goto L54
	} else {
		goto L63
	}
L59:
	;
	if base.I64_extend_i32_s(int32(0))*int64(1000) <= v87-v184 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[9]))
	goto L61
L61:
	;
	if base.I64_extend_i32_s(v194*int32(1000))*int64(1000) <= v192-v87 {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	v246 = v110
	v247 = int32(0)
	v252 = int32(1)
	goto L49
L63:
	;
	goto L55
L64:
	;
	v231 = v126
	v237 = int32(0)
	goto L50
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v222)+72))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v224)+72))
	if v223 < v225 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v231 = v110
	v237 = int32(0)
	goto L50
L67:
	;
	v281 = v126
	v282 = v111
	v287 = v116
	v293 = int32(1)
	goto L30
L68:
	;
	v263 = int32(3)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if base.B2i32(base.Ui32(v127) < base.Ui32(v263))|base.B2i32(base.Ui32(v265) < base.Ui32(v263)) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v127-v265 < int32(0) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v127) < base.Ui32(v265) {
		goto L67
	} else {
		goto L73
	}
L72:
	;
	v281 = v110
	v282 = v111
	v287 = v116
	v293 = int32(1)
	goto L30
L73:
	;
	v281 = v110
	v282 = v111
	v287 = v116
	v293 = int32(1)
	goto L30
L74:
	;
	goto L29
L75:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[0]))
	v317 = F_LWLockAcquire(m, v313+int32(2816), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v380 = int32(0)
	if v306 == v380 {
		v389 = v380
		goto L16
	} else {
		goto L85
	}
L78:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[1]))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v320)+20)) = v327 - int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v331
	v338 = m.G0
	v339 = int32(16)
	v340 = v338 - v339
	m.G0 = v340
	F_gettimeofday(m, v340)
	mBase = m.M
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	v344 = int64(*(*int32)(unsafe.Add(mBase, uint32(v340)+8)))
	m.G0 = v340 + v339
	goto L79
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v321)+24)) = v344 + v343*int64(1000000) - int64(946684800000000)
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+32)) = v321
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[0]))
	F_LWLockRelease(m, v358+int32(2816))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_start_worker[10])))
	if v365 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v389 = v379
	goto L16
L82:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v369+int32(20)))) = int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_do_start_worker[12]))
	v378 = F_pgmem_kill(m, v376, int32(10))
	mBase = m.M
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v389 = v380
	goto L16
L87:
	;
	v407 = v389
	goto L9
}
func F_do_tup_output(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	m.T0[v11].(func(*base.Module, int32))(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = v9 << (uint(int32(3)) % 32)
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			base.MemoryCopy(m, v16, l1, v15)
		} else {
		}
		if v9 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			base.MemoryCopy(m, v18, l2, v9)
		} else {
		}
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
		v22 = v20 & int32(_a_F_do_tup_output_0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)) = uint16(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)) = uint16(v25)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		v29 = m.T0[v28].(func(*base.Module, int32, int32) int32)(m, v7, v27)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			m.T0[v32].(func(*base.Module, int32))(m, v7)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_dotrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	v7 = int32(0)
	if base.B2i32(l1 <= v7)|base.B2i32(l3 <= v7) != 0 {
		v614 = l0
		v615 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v674 = F_cstring_to_text_with_len(m, v661, v662)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L12
	} else {
		goto L116
	}
L2:
	;
	v661 = v634
	v662 = v642
	goto L1
L3:
	;
	v634 = v614
	v642 = v615
	goto L2
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_dotrim[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28*int32(28))+uint32(_c_F_dotrim[1])))
	goto L8
L5:
	;
	F_pfree(m, v39)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L12
	} else {
		goto L111
	}
L6:
	;
	v444 = int32(0)
	if base.B2i32(l5 == v444)|base.B2i32(v430 <= v444) != 0 {
		v586 = v427
		v587 = v428
		goto L5
	} else {
		goto L79
	}
L7:
	;
	v427 = l0
	v428 = l1
	v430 = v74
	v442 = int32(-1)
	goto L6
L8:
	;
	if int32(2) <= v33 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = l1 << (uint(int32(2)) % 32)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if l4 == int32(0) {
		v320 = l0
		v321 = l1
		goto L56
	} else {
		goto L57
	}
L12:
	;
	return int32(0)
L13:
	;
	v43 = F_palloc(m, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v51 = l0
	v52 = l1
	v53 = v7
	goto L15
L15:
	;
	v66 = v53 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v39+v66))) = v51
	v70 = F_pg_mblen_range(m, v51, l0+l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v81 = l3 << (uint(int32(2)) % 32)
	v82 = F_palloc(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+v43))) = v70
	v74 = v53 + int32(1)
	v76 = v52 - v70
	if int32(0) < v76 {
		v51 = v51 + v70
		v52 = v76
		v53 = v74
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v84 = F_palloc(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v89 = l2
	v90 = l3
	v94 = int32(0)
	goto L21
L21:
	;
	v108 = v94 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v82+v108))) = v89
	v112 = F_pg_mblen_range(m, v89, l2+l3)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	if base.B2i32(l4 == int32(0))|base.B2i32(base.Ui32(int32(2147483646)) < base.Ui32(v53)) != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108+v84))) = v112
	v118 = v90 - v112
	if int32(0) < v118 {
		v89 = v89 + v112
		v90 = v118
		v94 = v94 + int32(1)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v132 = l0
	v133 = l1
	v135 = v74
	v144 = v7
	goto L26
L26:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v94) {
		goto L7
	} else {
		goto L28
	}
L27:
	;
	v586 = v257
	v587 = v256
	goto L5
L28:
	;
	v150 = v144 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v43+v150)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v39)))
	v158 = int32(0)
	goto L30
L29:
	;
	v254 = int32(1)
	v256 = v133 - v152
	v257 = v132 + v152
	if base.B2i32(v144 == v53) == int32(0) {
		v132 = v257
		v133 = v256
		v135 = v135 - v254
		v144 = v144 + v254
		goto L26
	} else {
		goto L55
	}
L30:
	;
	v177 = v158 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v84+v177)))
	if v179 == v152 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v427 = v132
	v428 = v133
	v430 = v135
	v442 = v144 - int32(1)
	goto L6
L32:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177+v82)))
	if base.Ui32(int32(4)) <= base.Ui32(v152) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	goto L34
L34:
	;
	if base.B2i32(v158 == v94) == int32(0) {
		v158 = v158 + int32(1)
		goto L30
	} else {
		goto L54
	}
L35:
	;
	if v244 == int32(0) {
		goto L29
	} else {
		goto L53
	}
L36:
	;
	v244 = int32(0)
	goto L35
L37:
	;
	v218 = v213
	v219 = v214
	v220 = v215
	goto L47
L38:
	;
	if (v154|v182)&int32(3) != 0 {
		v213 = v154
		v214 = v182
		v215 = v152
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v206 = v154
	v207 = v182
	v208 = v152
	goto L40
L40:
	;
	if v208 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v190 = v154
	v191 = v182
	v192 = v152
	goto L42
L42:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v195 != v196 {
		v213 = v190
		v214 = v191
		v215 = v192
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v206 = v201
	v207 = v199
	v208 = v203
	goto L40
L44:
	;
	v198 = int32(4)
	v199 = v191 + v198
	v201 = v190 + v198
	v203 = v192 - v198
	if base.Ui32(int32(3)) < base.Ui32(v203) {
		v190 = v201
		v191 = v199
		v192 = v203
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v213 = v206
	v214 = v207
	v215 = v208
	goto L37
L47:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v223 == v224 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v244 = v223 - v224
	goto L35
L49:
	;
	v226 = int32(1)
	v231 = v220 - v226
	if v231 != 0 {
		v218 = v218 + v226
		v219 = v219 + v226
		v220 = v231
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	goto L34
L54:
	;
	goto L31
L55:
	;
	goto L27
L56:
	;
	if l5 == int32(0) {
		v614 = v320
		v615 = v321
		goto L3
	} else {
		goto L66
	}
L57:
	;
	v266 = l0
	v267 = l1
	goto L58
L58:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v294 = int32(0)
	goto L61
L59:
	;
	v661 = l0 + l1
	v662 = v7
	goto L1
L60:
	;
	v314 = int32(1)
	if int32(2) <= v267 {
		v266 = v266 + v314
		v267 = v267 - v314
		goto L58
	} else {
		goto L65
	}
L61:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v294))))
	if v287 == v309 {
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v320 = v266
	v321 = v267
	goto L56
L63:
	;
	v312 = v294 + int32(1)
	if v312 != l3 {
		v294 = v312
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L59
L66:
	;
	v342 = int32(1)
	if l3 <= v342 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v345 = v342
	goto L69
L68:
	;
	v345 = l3
	goto L69
L69:
	;
	v347 = v321
	goto L70
L70:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v347-int32(1)))))
	v377 = int32(0)
	goto L72
L71:
	;
	v634 = v320
	v642 = v347
	goto L2
L72:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v377))))
	if v392 == v370 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L71
L74:
	;
	v398 = int32(0)
	if base.B2i32(v347 < int32(2)) == v398 {
		v347 = v347 - int32(1)
		goto L70
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v402 = v377 + int32(1)
	if v402 != v345 {
		v377 = v402
		goto L72
	} else {
		goto L78
	}
L77:
	;
	v634 = v320
	v642 = v398
	goto L2
L78:
	;
	goto L73
L79:
	;
	v457 = v430
	v458 = v428
	goto L80
L80:
	;
	if base.Ui32(int32(2147483646)) < base.Ui32(v94) {
		v586 = v427
		v587 = v428
		goto L5
	} else {
		goto L82
	}
L81:
	;
	v586 = v427
	v587 = v582
	goto L5
L82:
	;
	v473 = (v457 + v442) << (uint(int32(2)) % 32)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v43+v473)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473+v39)))
	v481 = int32(0)
	goto L84
L83:
	;
	goto L81
L84:
	;
	v500 = v481 << (uint(int32(2)) % 32)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v84+v500)))
	if v502 == v475 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v575 = v458 - v475
	v576 = int32(1)
	if v576 < v457 {
		v457 = v457 - v576
		v458 = v575
		goto L80
	} else {
		goto L110
	}
L86:
	;
	goto L85
L87:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v500+v82)))
	if base.Ui32(int32(4)) <= base.Ui32(v475) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	goto L89
L89:
	;
	if base.B2i32(v481 == v94) == int32(0) {
		v481 = v481 + int32(1)
		goto L84
	} else {
		goto L109
	}
L90:
	;
	if v567 == int32(0) {
		goto L86
	} else {
		goto L108
	}
L91:
	;
	v567 = int32(0)
	goto L90
L92:
	;
	v541 = v536
	v542 = v537
	v543 = v538
	goto L102
L93:
	;
	if (v477|v505)&int32(3) != 0 {
		v536 = v477
		v537 = v505
		v538 = v475
		goto L92
	} else {
		goto L96
	}
L94:
	;
	v529 = v477
	v530 = v505
	v531 = v475
	goto L95
L95:
	;
	if v531 == int32(0) {
		goto L91
	} else {
		goto L101
	}
L96:
	;
	v513 = v477
	v514 = v505
	v515 = v475
	goto L97
L97:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if v518 != v519 {
		v536 = v513
		v537 = v514
		v538 = v515
		goto L92
	} else {
		goto L99
	}
L98:
	;
	v529 = v524
	v530 = v522
	v531 = v526
	goto L95
L99:
	;
	v521 = int32(4)
	v522 = v514 + v521
	v524 = v513 + v521
	v526 = v515 - v521
	if base.Ui32(int32(3)) < base.Ui32(v526) {
		v513 = v524
		v514 = v522
		v515 = v526
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v536 = v529
	v537 = v530
	v538 = v531
	goto L92
L102:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	if v546 == v547 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v567 = v546 - v547
	goto L90
L104:
	;
	v549 = int32(1)
	v554 = v543 - v549
	if v554 != 0 {
		v541 = v541 + v549
		v542 = v542 + v549
		v543 = v554
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	goto L91
L108:
	;
	goto L89
L109:
	;
	v582 = v458
	goto L83
L110:
	;
	v582 = v575
	goto L83
L111:
	;
	F_pfree(m, v43)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L12
	} else {
		goto L112
	}
L112:
	;
	F_pfree(m, v82)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	F_pfree(m, v84)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	v611 = F_cstring_to_text_with_len(m, v586, v587)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	return v611
L116:
	;
	return v674
}
func F_downcase_truncate_identifier(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_downcase_identifier(m, l0, l1, l2, int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_drandom_normal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v124 int64
	_ = v124
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v267 float64
	_ = v267
	var v274 float64
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v294 float64
	_ = v294
	var v298 int32
	_ = v298
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v306 float64
	_ = v306
	var v307 float64
	_ = v307
	var v309 float64
	_ = v309
	var v311 float64
	_ = v311
	var v313 float64
	_ = v313
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_drandom_normal[0])))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(16)
	v15 = int32(0)
	v19 = m.G0
	v21 = v19 - v14
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v15
	v27 = F_open(m, int32(_a_F_drandom_normal_0), v15, v21)
	mBase = m.M
	if v27 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v137 = int32(_a_F_drandom_normal_1)
	v138 = int32(_a_F_drandom_normal_2)
	v139 = *(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[1]))
	v140 = int64(24)
	v143 = *(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[2]))
	v144 = v143 ^ v139
	v145 = int64(16)
	v148 = base.I64_rotl(v139, v140) ^ v144<<(uint(v145)%64) ^ v144
	v149 = int64(37)
	v151 = v148 ^ base.I64_rotl(v144, v149)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[2])) = base.I64_rotl(v151, v149)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[1])) = v151<<(uint(v145)%64) ^ base.I64_rotl(v148, v140) ^ v151
	goto L31
L4:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_drandom_normal[0])) = uint8(v135)
	goto L3
L5:
	;
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L10
L7:
	;
	v60 = v15
	goto L8
L8:
	;
	m.G0 = v21 + int32(16)
	goto L5
L9:
	;
	v55 = F_close(m, v27)
	mBase = m.M
	v60 = v53
	goto L8
L10:
	;
	v33 = int32(_a_F_drandom_normal_2)
	v34 = v14
	goto L11
L11:
	;
	v39 = F_read(m, v27, v33, v34)
	mBase = m.M
	if v39 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v53 = int32(1)
	goto L9
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_drandom_normal[3]))
	if v43 == int32(27) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v48 = v34 - v39
	if v48 != 0 {
		v33 = v33 + v39
		v34 = v48
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v53 = int32(0)
	goto L9
L17:
	;
	goto L12
L18:
	;
	v65 = int32(_a_F_drandom_normal_2)
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[1]))
	if v66 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v77 = int32(_a_F_drandom_normal_2)
	v81 = m.G0
	v82 = int32(16)
	v83 = v81 - v82
	m.G0 = v83
	F_gettimeofday(m, v83)
	mBase = m.M
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
	m.G0 = v83 + v82
	goto L26
L21:
	;
	goto L4
L22:
	;
	goto L21
L23:
	;
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[2]))
	if v69 != int64(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[2])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[1])) = int64(6364136223846793005)
	goto L22
L26:
	;
	v97 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_drandom_normal[4])))
	v100 = v87 + v86*int64(1000000) - int64(946684800000000) ^ v97<<(uint(int64(32))%64)
	v103 = v100 + int64(4354685564936845354)
	v104 = int64(30)
	v107 = int64(-4658895280553007687)
	v108 = (int64(base.Ui64(v103)>>(uint(v104)%64)) ^ v103) * v107
	v109 = int64(27)
	v112 = int64(-7723592293110705685)
	v113 = (int64(base.Ui64(v108)>>(uint(v109)%64)) ^ v108) * v112
	v114 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[2])) = int64(base.Ui64(v113)>>(uint(v114)%64)) ^ v113
	v119 = v100 - int64(7046029254386353131)
	v124 = (int64(base.Ui64(v119)>>(uint(v104)%64)) ^ v119) * v107
	v129 = (int64(base.Ui64(v124)>>(uint(v109)%64)) ^ v124) * v112
	*(*int64)(unsafe.Add(mBase, _c_F_drandom_normal[1])) = int64(base.Ui64(v129)>>(uint(v114)%64)) ^ v129
	goto L27
L27:
	;
	goto L4
L28:
	;
	goto L49
L29:
	;
	goto L28
L31:
	;
	goto L32
L32:
	;
	goto L29
L46:
	;
	v267 = F_log(m, base.F64_sub(float64(1), base.F64_mul(base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v139*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(971))<<(uint(int64(52))%64)))))
	mBase = m.M
	v274 = base.F64_mul(base.F64_sub(float64(1), base.F64_mul(base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v148*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(971))<<(uint(int64(52))%64)))), float64(6.283185307179586))
	v278 = m.G0
	v280 = v278 - int32(16)
	m.G0 = v280
	v287 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v274))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v287) <= base.Ui32(int32(1072243195)) {
		goto L66
	} else {
		goto L67
	}
L47:
	;
	goto L46
L49:
	;
	goto L50
L50:
	;
	goto L47
L64:
	;
	return base.I64_reinterpret_f64(base.F64_add(base.F64_mul(v7, base.F64_mul(base.F64_sqrt(base.F64_mul(v267, float64(-2))), v313)), v8))
L65:
	;
	m.G0 = v280 + int32(16)
	goto L64
L66:
	;
	if base.Ui32(v287) < base.Ui32(int32(1045430272)) {
		v313 = v274
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v287) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v294 = F___sin(m, v274, float64(0), int32(0))
	mBase = m.M
	v313 = v294
	goto L65
L70:
	;
	v313 = base.F64_sub(v274, v274)
	goto L65
L71:
	;
	goto L72
L72:
	;
	v298 = F___rem_pio2(m, v274, v280)
	mBase = m.M
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v280)+8))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v280)))
	switch v298&int32(3) - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L74
	case 2:
		goto L73
	default:
		goto L76
	}
L73:
	;
	v311 = F___cos(m, v300, v299)
	mBase = m.M
	v313 = base.F64_neg(v311)
	goto L65
L74:
	;
	v309 = F___sin(m, v300, v299, int32(1))
	mBase = m.M
	v313 = base.F64_neg(v309)
	goto L65
L75:
	;
	v307 = F___cos(m, v300, v299)
	mBase = m.M
	v313 = v307
	goto L65
L76:
	;
	v306 = F___sin(m, v300, v299, int32(1))
	mBase = m.M
	v313 = v306
	goto L65
}
func F_dsign(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v6 float64
	_ = v6
	var v9 float64
	_ = v9
	var v12 float64
	_ = v12
	v2 = float64(0)
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_lt(v6, v2) != 0 {
		v9 = float64(-1)
	} else {
		v9 = v2
	}
	if base.F64_gt(v6, float64(0)) != 0 {
		v12 = float64(1)
	} else {
		v12 = v9
	}
	return base.I64_reinterpret_f64(v12)
}
func F_dsnowball_init(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
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
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_palloc0(m, int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v351 != 0 {
		goto L98
	} else {
		goto L99
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v24 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = v2
	v39 = v2
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v46 = int32(_a_F_dsnowball_init_0)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsnowball_init[0])))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L13
	} else {
		goto L14
	}
L7:
	;
	goto L3
L8:
	;
	v337 = v37 + int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v337 < v338 {
		v37 = v337
		v39 = v335
		goto L6
	} else {
		goto L97
	}
L9:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v321 = m.T0[v320].(func(*base.Module) int32)(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L96
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L92
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L88
	}
L12:
	;
	if v70-v71 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v55 = v45
	v56 = v46
	goto L15
L15:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v70 = v60
	v71 = v59
	goto L13
L17:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v39 != 0 {
		goto L11
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v80 = int32(_a_F_dsnowball_init_1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsnowball_init[1])))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v75 = F_defGetString(m, v44)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_readstoplist(m, v75, v18+int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v335 = int32(1)
	goto L8
L25:
	;
	if v104-v105 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	v89 = v45
	v90 = v80
	goto L28
L28:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v104 = v94
	v105 = v93
	goto L26
L30:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v109 != 0 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L84
	}
L35:
	;
	v112 = F_defGetString(m, v44)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v115 = int32(_a_F_dsnowball_init_2)
	v120 = int32(_a_F_dsnowball_init_3)
	goto L37
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v177 = int32(_a_F_dsnowball_init_2)
	v181 = int32(_a_F_dsnowball_init_3)
	goto L60
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v174 != 0 {
		v115 = v115 + int32(20)
		v120 = v174
		goto L37
	} else {
		goto L59
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_init[2]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	goto L43
L41:
	;
	goto L42
L42:
	;
	v132 = v120
	v133 = v112
	goto L46
L43:
	;
	if v128 != v125 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v170 != 0 {
		goto L39
	} else {
		goto L58
	}
L46:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == v137 {
		v159 = v136
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v170 = int32(0)
	goto L45
L48:
	;
	v161 = int32(1)
	if v159 != 0 {
		v132 = v132 + v161
		v133 = v133 + v161
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v136-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v147 = v136 | int32(32)
	goto L52
L51:
	;
	v147 = v136
	goto L52
L52:
	;
	if base.Ui32((v137-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v156 = v137 | int32(32)
	goto L55
L54:
	;
	v156 = v137
	goto L55
L55:
	;
	if v147 == v156 {
		v159 = v147
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v170 = v147 - v156
	goto L45
L57:
	;
	goto L47
L58:
	;
	v308 = v115
	v311 = int32(0)
	goto L9
L59:
	;
	goto L38
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v188 != int32(6) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L79
	}
L62:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if v233 != 0 {
		v177 = v177 + int32(20)
		v181 = v233
		goto L60
	} else {
		goto L78
	}
L63:
	;
	v193 = v181
	v194 = v112
	goto L65
L64:
	;
	if v231 != 0 {
		goto L62
	} else {
		goto L77
	}
L65:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v197 == v198 {
		v220 = v197
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v231 = int32(0)
	goto L64
L67:
	;
	v222 = int32(1)
	if v220 != 0 {
		v193 = v193 + v222
		v194 = v194 + v222
		goto L65
	} else {
		goto L76
	}
L68:
	;
	if base.Ui32((v197-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v208 = v197 | int32(32)
	goto L71
L70:
	;
	v208 = v197
	goto L71
L71:
	;
	if base.Ui32((v198-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v217 = v198 | int32(32)
	goto L74
L73:
	;
	v217 = v198
	goto L74
L74:
	;
	if v208 == v217 {
		v220 = v208
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v231 = v208 - v217
	goto L64
L76:
	;
	goto L66
L77:
	;
	v308 = v177
	v311 = int32(1)
	goto L9
L78:
	;
	goto L61
L79:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_init[2]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v112
	F_errmsg(m, int32(_a_F_dsnowball_init_4), v14)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_dsnowball_init_5), int32(231), int32(_a_F_dsnowball_init_6))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v263
	F_errmsg(m, int32(_a_F_dsnowball_init_7), v14+int32(16))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_dsnowball_init_5), int32(270), int32(_a_F_dsnowball_init_8))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_dsnowball_init_9), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_dsnowball_init_5), int32(253), int32(_a_F_dsnowball_init_8))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_dsnowball_init_10), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_dsnowball_init_5), int32(262), int32(_a_F_dsnowball_init_8))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v321
	v335 = v39
	goto L8
L97:
	;
	goto L7
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_dsnowball_init[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v353
	m.G0 = v14 + int32(32)
	return base.I64_extend_i32_u(v18)
L99:
	;
	goto L100
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_dsnowball_init_11), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_dsnowball_init_5), int32(277), int32(_a_F_dsnowball_init_8))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_durable_unlink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	v9 = int32(-1)
	v10 = F_unlink(m, l0)
	mBase = m.M
	if v10 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(1040)
	return v216
L2:
	;
	v14 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v32 = v7 + int32(16)
	goto L14
L5:
	;
	return int32(0)
L6:
	;
	if v14 == int32(0) {
		v216 = v9
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(_a_F_durable_unlink_0), v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_durable_unlink_1), int32(885), int32(_a_F_durable_unlink_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v216 = v9
	goto L1
L11:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L12:
	;
	v149 = F_strlen(m, v138)
	mBase = m.M
	goto L11
L14:
	;
	goto L15
L15:
	;
	v39 = int32(1023)
	if (v32^l0)&int32(3) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v142)
	goto L12
L17:
	;
	v123 = v118
	v124 = v119
	v125 = v120
	goto L38
L18:
	;
	if v113 == int32(0) {
		v138 = v111
		v139 = v112
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v111 = l0
	v112 = v32
	v113 = v39
	goto L18
L20:
	;
	goto L21
L21:
	;
	v43 = int32(0)
	if base.B2i32(l0&int32(3) == v43)|int32(0) == v43 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v79 == int32(0) {
		v138 = v76
		v139 = v77
		goto L16
	} else {
		goto L31
	}
L23:
	;
	v55 = l0
	v56 = v32
	v57 = v39
	goto L26
L24:
	;
	goto L25
L25:
	;
	v76 = l0
	v77 = v32
	v78 = v39
	v79 = int32(1)
	goto L22
L26:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v59)
	if v59 == int32(0) {
		v118 = v55
		v119 = v56
		v120 = v57
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v76 = v70
	v77 = v64
	v78 = v66
	v79 = v68
	goto L22
L28:
	;
	v63 = int32(1)
	v64 = v56 + v63
	v66 = v57 - v63
	v67 = int32(0)
	v68 = base.B2i32(v66 != v67)
	v70 = v55 + v63
	if v70&int32(3) == v67 {
		v76 = v70
		v77 = v64
		v78 = v66
		v79 = v68
		goto L22
	} else {
		goto L29
	}
L29:
	;
	if v66 != 0 {
		v55 = v70
		v56 = v64
		v57 = v66
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if base.B2i32(v82 == int32(0))|base.B2i32(base.Ui32(v78) < base.Ui32(int32(4))) != 0 {
		v111 = v76
		v112 = v77
		v113 = v78
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v89 = v76
	v90 = v77
	v91 = v78
	goto L33
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v97 = int32(-2139062144)
	if (int32(16843008)-v94|v94)&v97 != v97 {
		v118 = v89
		v119 = v90
		v120 = v91
		goto L17
	} else {
		goto L35
	}
L34:
	;
	v111 = v105
	v112 = v103
	v113 = v107
	goto L18
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v94
	v102 = int32(4)
	v103 = v90 + v102
	v105 = v89 + v102
	v107 = v91 - v102
	if base.Ui32(int32(3)) < base.Ui32(v107) {
		v89 = v105
		v90 = v103
		v91 = v107
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v118 = v111
	v119 = v112
	v120 = v113
	goto L17
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v127)
	if v127 == int32(0) {
		v138 = v123
		v139 = v124
		goto L16
	} else {
		goto L40
	}
L39:
	;
	v138 = v134
	v139 = v132
	goto L16
L40:
	;
	v131 = int32(1)
	v132 = v124 + v131
	v134 = v123 + v131
	v136 = v125 - v131
	if v136 != 0 {
		v123 = v134
		v124 = v132
		v125 = v136
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
	if v202 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	v156 = F_strlen(m, v32)
	mBase = m.M
	v159 = v156 + v32
	goto L46
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v163 = v159 - int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if base.B2i32(v164 == int32(47))&base.B2i32(base.Ui32(v32) < base.Ui32(v163)) != 0 {
		v159 = v163
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v170 = v163
	goto L49
L48:
	;
	goto L47
L49:
	;
	if base.Ui32(v32) < base.Ui32(v170) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v182 = v170
	goto L55
L51:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v176 != int32(47) {
		v170 = v170 - int32(1)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	goto L53
L55:
	;
	if base.Ui32(v32) < base.Ui32(v182) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v32 == v182 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v186 = v182 - int32(1)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 == int32(47) {
		v182 = v186
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L59
L61:
	;
	v195 = v32 + base.B2i32(v155 == int32(47))
	goto L63
L62:
	;
	v195 = v182
	goto L63
L63:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v196)
	goto L45
L64:
	;
	v205 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)) = uint16(v205)
	goto L66
L65:
	;
	goto L66
L66:
	;
	v208 = int32(0)
	v213 = F_fsync_fname_ext(m, v7+int32(16), int32(1), v208, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	if v213 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v215 = int32(-1)
	goto L70
L69:
	;
	v215 = v208
	goto L70
L70:
	;
	v216 = v215
	goto L1
}
