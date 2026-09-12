package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAppendAsyncEventWait(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 float64
	_ = v306
	var v308 float64
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	v10 = m.G0
	v12 = v10 - int32(256)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	v16 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v19 = v17 + int32(2)
	v20 = F_CreateWaitEventSet(m, v16, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v20
	F_AddWaitEventToSet(m, v20, int32(32), int32(-1), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if int32(0) <= v85 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v85 = base.I32_ctz(v71) | v72<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v85 = int32(-2)
	goto L4
L7:
	;
	v38 = base.I32_div_s(int32(0), int32(32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 <= v38 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = v28 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32))))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v38
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v51 = v38 + int32(1)
	if v51 == v39 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v54 = v51
	goto L11
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+v54<<(uint(int32(2))%32))))
	if v61 != 0 {
		v71 = v61
		v72 = v54
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L6
L13:
	;
	v63 = v54 + int32(1)
	if v63 != v39 {
		v54 = v63
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v91 = v85
	goto L18
L16:
	;
	goto L17
L17:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v220 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v91<<(uint(int32(2))%32))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+12)))
	if v102 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v105 = m.G0
	v107 = v105 - int32(16)
	m.G0 = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	if v110 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v150 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	goto L22
L24:
	;
	F_InstrStartNode(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v114 = v109
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v115 == int32(418) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v114 = v113
	goto L26
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+128))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+176))
	m.T0[v120].(func(*base.Module, int32))(m, v101)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L36
	}
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	if v124 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_InstrStopNode(m, v124, float64(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	m.G0 = v107 + int32(16)
	goto L23
L35:
	;
	goto L34
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v136
	F_errmsg_internal(m, int32(484245), v107)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(498154), int32(76), int32(103345))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	if int32(0) <= v206 {
		v91 = v206
		goto L18
	} else {
		goto L50
	}
L40:
	;
	v206 = base.I32_ctz(v192) | v193<<(uint(int32(5))%32)
	goto L39
L41:
	;
	v206 = int32(-2)
	goto L39
L42:
	;
	v157 = v91 + int32(1)
	v159 = base.I32_div_s(v157, int32(32))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v160 <= v159 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v163 = v150 + int32(8)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v159<<(uint(int32(2))%32))))
	v170 = v167 & (int32(-1) << (uint(v157) % 32))
	if v170 != 0 {
		v192 = v170
		v193 = v159
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v172 = v159 + int32(1)
	if v172 == v160 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v175 = v172
	goto L46
L46:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v163+v175<<(uint(int32(2))%32))))
	if v182 != 0 {
		v192 = v182
		v193 = v175
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v184 = v175 + int32(1)
	if v184 != v160 {
		v175 = v184
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L19
L51:
	;
	m.G0 = v12 + int32(256)
	return
L52:
	;
	F_FreeWaitEventSet(m, v218)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[497]))
	F_AddWaitEventToSet(m, v218, int32(1), int32(-1), v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	goto L51
L56:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v237 = int32(16)
	if v237 <= v19 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v240 = v237
	goto L59
L58:
	;
	v240 = v19
	goto L59
L59:
	;
	v242 = F_WaitEventSetWait(m, v234, int32(0)-v14, v12, v240, int32(134217728))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_FreeWaitEventSet(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v247
	if v242 <= v247 {
		goto L51
	} else {
		goto L62
	}
L62:
	;
	v254 = int32(0)
	goto L63
L63:
	;
	v262 = v12 + v254<<(uint(int32(4))%32)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v263&int32(2) == int32(0) {
		v348 = v263
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L51
L65:
	;
	if v348&int32(1) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L66:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+12)))
	if v269 != int32(1) {
		v348 = v263
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v272 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+12)) = uint8(v272)
	v274 = m.G0
	v276 = v274 - int32(32)
	m.G0 = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	if v279 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v348 = v347
	goto L65
L69:
	;
	F_InstrStartNode(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	v283 = v278
	goto L71
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 == int32(418) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v283 = v282
	goto L71
L73:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+128))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+180))
	m.T0[v289].(func(*base.Module, int32))(m, v268)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L94
	}
L76:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v293 == int32(397) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_ExecAsyncAppendResponse(m, v268)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L91
	}
L80:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+20))
	if v299 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	if v300 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	m.G0 = v276 + int32(32)
	goto L68
L84:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+4)))
	if v303&int32(2) != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v308 = float64(0)
	goto L86
L86:
	;
	F_InstrStopNode(m, v299, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L90
	}
L87:
	;
	v306 = float64(0)
	goto L89
L88:
	;
	v306 = float64(1)
	goto L89
L89:
	;
	v308 = v306
	goto L86
L90:
	;
	goto L83
L91:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v320
	F_errmsg_internal(m, int32(484245), v276)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(498154), int32(127), int32(360214))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v335
	F_errmsg_internal(m, int32(484245), v276+int32(16))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(498154), int32(102), int32(20753))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v367 = v254 + int32(1)
	if v367 != v242 {
		v254 = v367
		goto L63
	} else {
		goto L102
	}
L98:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[497]))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = int32(0)
	goto L99
L99:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v361 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	goto L64
}
