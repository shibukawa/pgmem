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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
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
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
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
	v365 = m.ExcPending
	if v365 != 0 {
		goto L13
	} else {
		goto L108
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v354
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))) < base.Ui32(v79) {
		goto L28
	} else {
		goto L29
	}
L6:
	;
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v75)
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = F_shm_mq_wait_internal(m, v14, v14+int32(8), v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	F_s_lock(m, v14, int32(_a_F_shm_mq_receive_0), int32(261), int32(_a_F_shm_mq_receive_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v48 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14))), uint32(v48))
	if v47 == v48 {
		v354 = v37
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
	if v56 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v60 = base.AtomicRmwXchg32(m, v14, int32(0), int32(1))
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_s_lock(m, v14, int32(_a_F_shm_mq_receive_0), int32(261), int32(_a_F_shm_mq_receive_1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v67 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14))), uint32(v67))
	if v66 != 0 {
		goto L6
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v70)
	v354 = int32(2)
	goto L2
L28:
	;
	v85 = int64(0)
	v87 = int32(16)
	v88 = base.AtomicRmwCmpxchg64(m, v14, v87, v85, v85)
	v91 = base.AtomicRmwXchg64(m, v14, v87, base.I64_extend_i32_u(v79)+v88)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v94 = v92 + int32(20)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v140 != 0 {
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
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v98 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v101 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[0]))
	if v105 == v101 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v107 = m.G0
	v109 = v107 - int32(16)
	m.G0 = v109
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[1]))
	if v112 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v135 = F_pgmem_kill(m, v101, int32(23))
	mBase = m.M
	goto L32
L39:
	;
	m.G0 = v109 + int32(16)
	goto L31
L40:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+15)) = uint8(v115)
	goto L41
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[2]))
	v123 = F_write(m, v119, v109+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v123 {
		goto L39
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[3]))
	if v127 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v250) {
		goto L1
	} else {
		goto L70
	}
L46:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v146 = v141
	goto L47
L47:
	;
	v157 = F_shm_mq_receive_bytes(m, l0, int32(4)-v146, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L49
	}
L48:
	;
	goto L45
L49:
	;
	if v157 != 0 {
		v354 = v157
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v160 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.B2i32(v159 == v160)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v162)) == v160 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v238 != int32(1) {
		v146 = v191
		goto L47
	} else {
		goto L69
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v168 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v218 = (v212+int32(7))&int32(-8) + int32(8)
	if base.Ui32(v162) < base.Ui32(v218) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v173 = F_MemoryContextAlloc(m, v171, int32(_a_F_shm_mq_receive_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L13
	} else {
		goto L58
	}
L56:
	;
	v179 = v159
	v180 = v168
	goto L57
L57:
	;
	v181 = int32(4)
	if base.Ui32(v181) < base.Ui32(v179+v162) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(_a_F_shm_mq_receive_2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v173
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v179 = v178
	v180 = v173
	goto L57
L59:
	;
	v186 = v181 - v179
	goto L61
L60:
	;
	v186 = v162
	goto L61
L61:
	;
	if v186 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v179+v180, v188, v186)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v191 = v190 + v186
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v193 + (v186+int32(7))&int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v162 - v186
	if base.Ui32(v191) < base.Ui32(int32(4)) {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L45
L66:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v220)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v212
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v224 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v162 - v224
	goto L45
L67:
	;
	goto L68
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v230 + v218
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v212
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211 + int32(8)
	v354 = int32(0)
	goto L2
L69:
	;
	goto L48
L70:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v253 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v311 = v302
	goto L92
L72:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v302 = v254
	goto L71
L73:
	;
	goto L74
L74:
	;
	v259 = F_shm_mq_receive_bytes(m, l0, v250, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	if v259 != 0 {
		v354 = v259
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v250) <= base.Ui32(v261) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v263 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v263)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v266 + (v250+int32(7))&int32(2147483640)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v250
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v274
	v354 = v263
	goto L2
L78:
	;
	goto L79
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v250) <= base.Ui32(v276) {
		v302 = v261
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v279 = int32(1)
	if v250&(v250-v279) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v287 = v279 << (uint(int32(32)-base.I32_clz(v250)) % 32)
	goto L83
L82:
	;
	v287 = v250
	goto L83
L83:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v287) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v290 = int32(1073741823)
	goto L86
L85:
	;
	v290 = v287
	goto L86
L86:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v291 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v297 = F_MemoryContextAlloc(m, v296, v290)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v297
	v302 = v261
	goto L71
L92:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v311 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v250
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v342
	v344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v344
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v344)
	v354 = v344
	goto L2
L94:
	;
	if v311 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v321 = v313
	goto L96
L96:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v322 + (v311+int32(7))&int32(-8)
	if base.Ui32(v321) < base.Ui32(v250) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v314+v313, v316, v311)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v319 = v318 + v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v319
	v321 = v319
	goto L96
L100:
	;
	v330 = v250 - v321
	v335 = F_shm_mq_receive_bytes(m, l0, v330, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
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
	if v335 != 0 {
		v354 = v335
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v337) < base.Ui32(v330) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v339 = v337
	goto L107
L106:
	;
	v339 = v330
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v339
	v311 = v339
	goto L92
L108:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v250
	F_errmsg(m, int32(_a_F_shm_mq_receive_3), v12)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_shm_mq_receive_0), int32(719), int32(_a_F_shm_mq_receive_4))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
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
