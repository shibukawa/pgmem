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
	var v101 int32
	_ = v101
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
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
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
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
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
	v322 = m.ExcPending
	if v322 != 0 {
		goto L13
	} else {
		goto L97
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v310
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
	F_s_lock(m, v14, int32(495340), int32(261), int32(227152))
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
		v310 = v37
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
	F_s_lock(m, v14, int32(495340), int32(261), int32(227152))
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
	v310 = int32(2)
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
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v206) {
		goto L1
	} else {
		goto L58
	}
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v101 = v94
	goto L34
L34:
	;
	v110 = F_shm_mq_receive_bytes(m, l0, int32(4)-v101, l3, v12+int32(8), v12+int32(12))
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
		v310 = v110
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
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v192&int32(1) == int32(0) {
		v101 = v145
		goto L34
	} else {
		goto L57
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
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v172 = (v166+int32(7))&int32(-8) + int32(8)
	if base.Ui32(v115) < base.Ui32(v172) {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v126 = F_MemoryContextAlloc(m, v124, int32(8192))
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
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v136 = int32(4)
	if base.Ui32(v136) < base.Ui32(v132+v115) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v126
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v132 = v131
	v133 = v126
	goto L44
L46:
	;
	v141 = v136 - v132
	goto L48
L47:
	;
	v141 = v115
	goto L48
L48:
	;
	if v141 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v145 = v144 + v141
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v147 + (v141+int32(7))&int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v115 - v141
	if base.Ui32(v145) < base.Ui32(int32(4)) {
		goto L38
	} else {
		goto L53
	}
L50:
	;
	v142 = F__emscripten_memcpy_bulkmem(m, v132+v133, v135, v141)
	mBase = m.M
	goto L52
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L32
L54:
	;
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v174)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v166
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v178 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v177 + v178
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v115 - v178
	goto L32
L55:
	;
	goto L56
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v184 + v172
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v166
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v165 + int32(8)
	v310 = int32(0)
	goto L2
L57:
	;
	goto L35
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v209 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v265 = v258
	goto L80
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v258 = v210
	goto L59
L61:
	;
	goto L62
L62:
	;
	v215 = F_shm_mq_receive_bytes(m, l0, v206, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	if v215 != 0 {
		v310 = v215
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v206) <= base.Ui32(v217) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v219)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v222 + (v206+int32(7))&int32(2147483640)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v206
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v230
	v310 = v219
	goto L2
L66:
	;
	goto L67
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v206) <= base.Ui32(v232) {
		v258 = v217
		goto L59
	} else {
		goto L68
	}
L68:
	;
	v235 = int32(1)
	if v206&(v206-v235) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v243 = v235 << (uint(int32(32)-base.I32_clz(v206)) % 32)
	goto L71
L70:
	;
	v243 = v206
	goto L71
L71:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v243) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v246 = int32(1073741823)
	goto L74
L73:
	;
	v246 = v243
	goto L74
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_pfree(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L13
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v253 = F_MemoryContextAlloc(m, v252, v246)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L79
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v253
	v258 = v217
	goto L59
L80:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v265 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v206
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v299
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v301
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v301)
	v310 = v301
	goto L2
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v265 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v278 = v269
	goto L84
L84:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v279 + (v265+int32(7))&int32(-8)
	if base.Ui32(v278) < base.Ui32(v206) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v276 = v275 + v265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v276
	v278 = v276
	goto L84
L86:
	;
	v273 = F__emscripten_memcpy_bulkmem(m, v270+v269, v272, v265)
	mBase = m.M
	goto L88
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	v287 = v206 - v278
	v292 = F_shm_mq_receive_bytes(m, l0, v287, l3, v12+int32(8), v12+int32(12))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L81
L92:
	;
	if v292 != 0 {
		v310 = v292
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(v294) < base.Ui32(v287) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v296 = v294
	goto L96
L95:
	;
	v296 = v287
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v296
	v265 = v296
	goto L80
L97:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v206
	F_errmsg(m, int32(347503), v12)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(495340), int32(719), int32(343812))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
