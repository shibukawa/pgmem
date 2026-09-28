package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_shm_mq_receive(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v5
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v17 == v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L13
	} else {
		goto L108
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v362
L3:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if base.Ui32(int32(base.Ui32(v76)>>(uint(int32(2))%32))) < base.Ui32(v75) {
		goto L28
	} else {
		goto L29
	}
L6:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v71)
	goto L5
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)))
	if v21 != 0 {
		v37 = int32(2)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v54 = F_shm_mq_wait_internal(m, v14, v14+int32(8), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L13
	} else {
		goto L21
	}
L10:
	;
	v41 = base.AtomicRmwXchg32(m, v14, int32(0), int32(1))
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v22 = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 == int32(0) {
		v37 = v22
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v28 = F_GetBackgroundWorkerPid(m, v23, v12+int32(12))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if base.Ui32(v28) < base.Ui32(int32(2)) {
		v37 = v22
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v34)
	v37 = int32(2)
	goto L10
L16:
	;
	F_s_lock(m, v14, int32(_a_F_shm_mq_receive_0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v46 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14))), uint32(v46))
	if v45 == v46 {
		v362 = v37
		goto L2
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L6
L21:
	;
	if v54 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v58 = base.AtomicRmwXchg32(m, v14, int32(0), int32(1))
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_s_lock(m, v14, int32(_a_F_shm_mq_receive_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v63 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14))), uint32(v63))
	if v62 != 0 {
		goto L6
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v66)
	v362 = int32(2)
	goto L2
L28:
	;
	v80 = int32(0)
	v83 = base.AtomicRmwOr32(m, v80, int32(_a_F_shm_mq_receive_1), v80)
	v85 = int64(0)
	v87 = int32(16)
	v88 = base.AtomicRmwCmpxchg64(m, v14, v87, v85, v85)
	v91 = base.AtomicRmwXchg64(m, v14, v87, base.I64_extend_i32_u(v75)+v88)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v94 = v92 + int32(316)
	v98 = base.AtomicRmwOr32(m, v80, int32(_a_F_shm_mq_receive_2), v80)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v148 != 0 {
		goto L45
	} else {
		goto L46
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L30
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(1)
	v102 = int32(0)
	v105 = base.AtomicRmwOr32(m, v102, int32(_a_F_shm_mq_receive_2), v102)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v106 == v102 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v109 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[0]))
	if v113 == v109 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v115 = m.G0
	v117 = v115 - int32(16)
	m.G0 = v117
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[1]))
	if v120 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v143 = F_pgmem_kill(m, v109, int32(23))
	mBase = m.M
	goto L32
L39:
	;
	m.G0 = v117 + int32(16)
	goto L31
L40:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+15)) = uint8(v123)
	goto L41
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[2]))
	v131 = F_write(m, v127, v117+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v131 {
		goto L39
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[3]))
	if v135 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v258) {
		goto L1
	} else {
		goto L70
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v154 = v149
	goto L47
L47:
	;
	v165 = F_shm_mq_receive_bytes(m, l0, int32(4)-v154, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L49
	}
L48:
	;
	goto L45
L49:
	;
	if v165 != 0 {
		v362 = v165
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v168 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.B2i32(v167 == v168)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v170)) == v168 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v246 != int32(1) {
		v154 = v199
		goto L47
	} else {
		goto L69
	}
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v176 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v226 = (v220+int32(7))&int32(-8) + int32(8)
	if base.Ui32(v170) < base.Ui32(v226) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v181 = F_MemoryContextAlloc(m, v179, int32(_a_F_shm_mq_receive_3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L58
	}
L56:
	;
	v187 = v167
	v188 = v176
	goto L57
L57:
	;
	v189 = int32(4)
	if base.Ui32(v189) < base.Ui32(v187+v170) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(_a_F_shm_mq_receive_3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v187 = v186
	v188 = v181
	goto L57
L59:
	;
	v194 = v189 - v187
	goto L61
L60:
	;
	v194 = v170
	goto L61
L61:
	;
	if v194 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v187+v188, v196, v194)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v199 = v198 + v194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v201 + (v194+int32(7))&int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v170 - v194
	if base.Ui32(v199) < base.Ui32(int32(4)) {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L45
L66:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v220
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v232 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v231 + v232
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v170 - v232
	goto L45
L67:
	;
	goto L68
L68:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v238 + v226
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v220
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v219 + int32(8)
	v362 = int32(0)
	goto L2
L69:
	;
	goto L48
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v261 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v319 = v310
	goto L92
L72:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v310 = v262
	goto L71
L73:
	;
	goto L74
L74:
	;
	v267 = F_shm_mq_receive_bytes(m, l0, v258, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	if v267 != 0 {
		v362 = v267
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v258) <= base.Ui32(v269) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v271)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274 + (v258+int32(7))&int32(2147483640)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v258
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v282
	v362 = v271
	goto L2
L78:
	;
	goto L79
L79:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v258) <= base.Ui32(v284) {
		v310 = v269
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v287 = int32(1)
	if v258&(v258-v287) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v295 = v287 << (uint(int32(32)-base.I32_clz(v258)) % 32)
	goto L83
L82:
	;
	v295 = v258
	goto L83
L83:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v295) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v298 = int32(1073741823)
	goto L86
L85:
	;
	v298 = v295
	goto L86
L86:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v299 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v305 = F_MemoryContextAlloc(m, v304, v298)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L13
	} else {
		goto L91
	}
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	goto L89
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v305
	v310 = v269
	goto L71
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v319 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v258
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v350
	v352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v352
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v352)
	v362 = v352
	goto L2
L94:
	;
	if v319 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v329 = v321
	goto L96
L96:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v330 + (v319+int32(7))&int32(-8)
	if base.Ui32(v329) < base.Ui32(v258) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v322+v321, v324, v319)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v327 = v326 + v319
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v327
	v329 = v327
	goto L96
L100:
	;
	v338 = v258 - v329
	v343 = F_shm_mq_receive_bytes(m, l0, v338, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L13
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L93
L103:
	;
	if v343 != 0 {
		v362 = v343
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v345) < base.Ui32(v338) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v347 = v345
	goto L107
L106:
	;
	v347 = v338
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v347
	v319 = v347
	goto L92
L108:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v258
	F_errmsg(m, int32(_a_F_shm_mq_receive_4), v12)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_shm_mq_receive_5), int32(721), int32(_a_F_shm_mq_receive_6))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
