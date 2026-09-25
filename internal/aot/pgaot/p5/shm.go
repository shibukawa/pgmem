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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L108
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v366
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
		v366 = v37
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
	v366 = int32(2)
	goto L2
L28:
	;
	v84 = int32(0)
	v87 = base.AtomicRmwOr32(m, v84, int32(_a_F_shm_mq_receive_2), v84)
	v89 = int64(0)
	v91 = int32(16)
	v92 = base.AtomicRmwCmpxchg64(m, v14, v91, v89, v89)
	v95 = base.AtomicRmwXchg64(m, v14, v91, base.I64_extend_i32_u(v79)+v92)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v98 = v96 + int32(20)
	v102 = base.AtomicRmwOr32(m, v84, int32(_a_F_shm_mq_receive_3), v84)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v103 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v152 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = int32(1)
	v106 = int32(0)
	v109 = base.AtomicRmwOr32(m, v106, int32(_a_F_shm_mq_receive_3), v106)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v110 == v106 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if v113 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[0]))
	if v117 == v113 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = m.G0
	v121 = v119 - int32(16)
	m.G0 = v121
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[1]))
	if v124 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v147 = F_pgmem_kill(m, v113, int32(23))
	mBase = m.M
	goto L32
L39:
	;
	m.G0 = v121 + int32(16)
	goto L31
L40:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+15)) = uint8(v127)
	goto L41
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[2]))
	v135 = F_write(m, v131, v121+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v135 {
		goto L39
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_receive[3]))
	if v139 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v262) {
		goto L1
	} else {
		goto L70
	}
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v158 = v153
	goto L47
L47:
	;
	v169 = F_shm_mq_receive_bytes(m, l0, int32(4)-v158, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L49
	}
L48:
	;
	goto L45
L49:
	;
	if v169 != 0 {
		v366 = v169
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v172 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.B2i32(v171 == v172)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v174)) == v172 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v250 != int32(1) {
		v158 = v203
		goto L47
	} else {
		goto L69
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v180 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v230 = (v224+int32(7))&int32(-8) + int32(8)
	if base.Ui32(v174) < base.Ui32(v230) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v185 = F_MemoryContextAlloc(m, v183, int32(_a_F_shm_mq_receive_4))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L58
	}
L56:
	;
	v191 = v171
	v192 = v180
	goto L57
L57:
	;
	v193 = int32(4)
	if base.Ui32(v193) < base.Ui32(v191+v174) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(_a_F_shm_mq_receive_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v191 = v190
	v192 = v185
	goto L57
L59:
	;
	v198 = v193 - v191
	goto L61
L60:
	;
	v198 = v174
	goto L61
L61:
	;
	if v198 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v191+v192, v200, v198)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v203 = v202 + v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v205 + (v198+int32(7))&int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v174 - v198
	if base.Ui32(v203) < base.Ui32(int32(4)) {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v218)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L45
L66:
	;
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v224
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v236 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v235 + v236
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v174 - v236
	goto L45
L67:
	;
	goto L68
L68:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v242 + v230
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v224
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v223 + int32(8)
	v366 = int32(0)
	goto L2
L69:
	;
	goto L48
L70:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v265 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v323 = v314
	goto L92
L72:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v314 = v266
	goto L71
L73:
	;
	goto L74
L74:
	;
	v271 = F_shm_mq_receive_bytes(m, l0, v262, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	if v271 != 0 {
		v366 = v271
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v262) <= base.Ui32(v273) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v275)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v278 + (v262+int32(7))&int32(2147483640)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v262
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v286
	v366 = v275
	goto L2
L78:
	;
	goto L79
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v262) <= base.Ui32(v288) {
		v314 = v273
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v291 = int32(1)
	if v262&(v262-v291) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v299 = v291 << (uint(int32(32)-base.I32_clz(v262)) % 32)
	goto L83
L82:
	;
	v299 = v262
	goto L83
L83:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v299) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v302 = int32(1073741823)
	goto L86
L85:
	;
	v302 = v299
	goto L86
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v303 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v309 = F_MemoryContextAlloc(m, v308, v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v309
	v314 = v273
	goto L71
L92:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v323 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v262
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v354
	v356 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v356
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v356)
	v366 = v356
	goto L2
L94:
	;
	if v323 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v333 = v325
	goto L96
L96:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v334 + (v323+int32(7))&int32(-8)
	if base.Ui32(v333) < base.Ui32(v262) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v326+v325, v328, v323)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v331 = v330 + v323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v331
	v333 = v331
	goto L96
L100:
	;
	v342 = v262 - v333
	v347 = F_shm_mq_receive_bytes(m, l0, v342, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
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
	if v347 != 0 {
		v366 = v347
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v349) < base.Ui32(v342) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v351 = v349
	goto L107
L106:
	;
	v351 = v342
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v351
	v323 = v351
	goto L92
L108:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v262
	F_errmsg(m, int32(_a_F_shm_mq_receive_5), v12)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_shm_mq_receive_0), int32(719), int32(_a_F_shm_mq_receive_6))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
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
