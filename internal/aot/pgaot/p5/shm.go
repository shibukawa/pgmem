package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
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
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L95
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v307
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if base.Ui32(int32(base.Ui32(v78)>>(uint(int32(2))%32))) < base.Ui32(v77) {
		goto L28
	} else {
		goto L29
	}
L6:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v73)
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
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = F_shm_mq_wait_internal(m, v14, v14+int32(8), v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L21
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	if v39 != 0 {
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
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v49 == v47 {
		v307 = v37
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
	if v55 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	if v57 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_s_lock(m, v14, int32(_a_F_shm_mq_receive_0), int32(261), int32(_a_F_shm_mq_receive_1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v67 != 0 {
		goto L6
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v68)
	v307 = int32(2)
	goto L2
L28:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82 + base.I64_extend_i32_u(v77)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	F_SetLatch(m, v86+int32(20))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v93 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L30
L32:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v203) {
		goto L1
	} else {
		goto L57
	}
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v99 = v94
	goto L34
L34:
	;
	v110 = F_shm_mq_receive_bytes(m, l0, int32(4)-v99, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L36
	}
L35:
	;
	goto L32
L36:
	;
	if v110 != 0 {
		v307 = v110
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v113 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.B2i32(v112 == v113)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v115)) == v113 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v191 != int32(1) {
		v99 = v144
		goto L34
	} else {
		goto L56
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v121 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v171 = (v165+int32(7))&int32(-8) + int32(8)
	if base.Ui32(v115) < base.Ui32(v171) {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v126 = F_MemoryContextAlloc(m, v124, int32(_a_F_shm_mq_receive_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L45
	}
L43:
	;
	v132 = v112
	v133 = v121
	goto L44
L44:
	;
	v134 = int32(4)
	if base.Ui32(v134) < base.Ui32(v132+v115) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(_a_F_shm_mq_receive_2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v132 = v131
	v133 = v126
	goto L44
L46:
	;
	v139 = v134 - v132
	goto L48
L47:
	;
	v139 = v115
	goto L48
L48:
	;
	if v139 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v132+v133, v141, v139)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v144 = v143 + v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146 + (v139+int32(7))&int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v115 - v139
	if base.Ui32(v144) < base.Ui32(int32(4)) {
		goto L38
	} else {
		goto L52
	}
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L32
L53:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v173)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v165
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v177 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v176 + v177
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v115 - v177
	goto L32
L54:
	;
	goto L55
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v183 + v171
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v165
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v164 + int32(8)
	v307 = int32(0)
	goto L2
L56:
	;
	goto L35
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v206 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v264 = v255
	goto L79
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v255 = v207
	goto L58
L60:
	;
	goto L61
L61:
	;
	v212 = F_shm_mq_receive_bytes(m, l0, v203, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	if v212 != 0 {
		v307 = v212
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v203) <= base.Ui32(v214) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v216)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v219 + (v203+int32(7))&int32(2147483640)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v203
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v227
	v307 = v216
	goto L2
L65:
	;
	goto L66
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v203) <= base.Ui32(v229) {
		v255 = v214
		goto L58
	} else {
		goto L67
	}
L67:
	;
	v232 = int32(1)
	if v203&(v203-v232) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v240 = v232 << (uint(int32(32)-base.I32_clz(v203)) % 32)
	goto L70
L69:
	;
	v240 = v203
	goto L70
L70:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v240) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v243 = int32(1073741823)
	goto L73
L72:
	;
	v243 = v240
	goto L73
L73:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v244 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L13
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v250 = F_MemoryContextAlloc(m, v249, v243)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L78
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v250
	v255 = v214
	goto L58
L79:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v264 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v203
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v295
	v297 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v297
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v297)
	v307 = v297
	goto L2
L81:
	;
	if v264 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v274 = v266
	goto L83
L83:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v275 + (v264+int32(7))&int32(-8)
	if base.Ui32(v274) < base.Ui32(v203) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	base.MemoryCopy(m, v267+v266, v269, v264)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v272 = v271 + v264
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v272
	v274 = v272
	goto L83
L87:
	;
	v283 = v203 - v274
	v288 = F_shm_mq_receive_bytes(m, l0, v283, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	goto L80
L90:
	;
	if v288 != 0 {
		v307 = v288
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v290) < base.Ui32(v283) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v292 = v290
	goto L94
L93:
	;
	v292 = v283
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v292
	v264 = v292
	goto L79
L95:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v203
	F_errmsg(m, int32(_a_F_shm_mq_receive_3), v12)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_shm_mq_receive_0), int32(719), int32(_a_F_shm_mq_receive_4))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
