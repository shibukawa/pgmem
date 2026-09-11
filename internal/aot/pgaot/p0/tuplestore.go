package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_trim(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v12&int32(4) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v17 <= int32(0) {
		v91 = v16
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v101 = v91 - int32(1)
	if v101 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v20 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v17 == v20 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v17&v20 == int32(0) {
		v91 = v71
		goto L4
	} else {
		goto L25
	}
L7:
	;
	v70 = int32(0)
	v71 = v16
	goto L6
L8:
	;
	goto L9
L9:
	;
	v30 = int32(0)
	v31 = v16
	v35 = int32(0)
	goto L10
L10:
	;
	v42 = v22 + v30*int32(24)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v70 = v65
	v71 = v62
	goto L6
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v31 < v46 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v49 = v31
	goto L14
L14:
	;
	v55 = v22 + (v30|int32(1))*int32(24)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	if v56 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v48 = v31
	goto L17
L16:
	;
	v48 = v46
	goto L17
L17:
	;
	v49 = v48
	goto L14
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v49 < v59 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v62 = v49
	goto L20
L20:
	;
	v64 = int32(2)
	v65 = v30 + v64
	v67 = v35 + v64
	if v67 != v17&int32(2147483646) {
		v30 = v65
		v31 = v62
		v35 = v67
		goto L10
	} else {
		goto L24
	}
L21:
	;
	v61 = v49
	goto L23
L22:
	;
	v61 = v59
	goto L23
L23:
	;
	v62 = v61
	goto L20
L24:
	;
	goto L11
L25:
	;
	v84 = v22 + v70*int32(24)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v85 != 0 {
		v91 = v71
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v71 < v86 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v71
	goto L29
L28:
	;
	v88 = v86
	goto L29
L29:
	;
	v91 = v88
	goto L4
L30:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v107 = v105 - v106
	if v107 < v104 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = v104
	goto L33
L32:
	;
	v109 = v107
	goto L33
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v111 < v101 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v114 = v111
	goto L37
L35:
	;
	v153 = v16
	goto L36
L36:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v101
	v163 = base.I32_div_s(v153, int32(8))
	if v101 < v163 {
		goto L1
	} else {
		goto L43
	}
L37:
	;
	v125 = v114 << (uint(int32(2)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125+v126)))
	v129 = F_GetMemoryChunkSpace(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v153 = v147
	goto L36
L39:
	;
	return
L40:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v131 + base.I64_extend_i32_u(v129)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135+v125)))
	F_pfree(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v140+v125))) = int32(0)
	v145 = v114 + int32(1)
	if v145 != v101 {
		v114 = v145
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v168 = v165 + v101<<(uint(int32(2))%32)
	if v91 == v153 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v319 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v319
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v322 - v101
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v325 <= v319 {
		goto L1
	} else {
		goto L94
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v170
	goto L44
L46:
	;
	goto L47
L47:
	;
	v174 = (v153 - v101) << (uint(int32(2)) % 32)
	if v165 == v168 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L44
L49:
	;
	goto L48
L50:
	;
	v178 = v165 + v174
	if base.Ui32(v168-v178) <= base.Ui32(int32(0)-v174<<(uint(int32(1))%32)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v185 = F___memcpy(m, v165, v168, v174)
	mBase = m.M
	goto L48
L52:
	;
	goto L53
L53:
	;
	v188 = (v165 ^ v168) & int32(3)
	if base.Ui32(v165) < base.Ui32(v168) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v290 == int32(0) {
		goto L49
	} else {
		goto L90
	}
L55:
	;
	if base.Ui32(v268) <= base.Ui32(int32(3)) {
		v289 = v267
		v290 = v268
		v291 = v269
		goto L54
	} else {
		goto L86
	}
L56:
	;
	if v188 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v188 != 0 {
		v250 = v174
		goto L69
	} else {
		goto L70
	}
L59:
	;
	v289 = v168
	v290 = v174
	v291 = v165
	goto L54
L60:
	;
	goto L61
L61:
	;
	if v165&int32(3) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v267 = v168
	v268 = v174
	v269 = v165
	goto L55
L63:
	;
	goto L64
L64:
	;
	v195 = v168
	v196 = v174
	v197 = v165
	goto L65
L65:
	;
	if v196 == int32(0) {
		goto L49
	} else {
		goto L67
	}
L66:
	;
	v267 = v204
	v268 = v206
	v269 = v208
	goto L55
L67:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v201)
	v203 = int32(1)
	v204 = v195 + v203
	v206 = v196 - v203
	v208 = v197 + v203
	if v208&int32(3) != 0 {
		v195 = v204
		v196 = v206
		v197 = v208
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	if v250 == int32(0) {
		goto L49
	} else {
		goto L82
	}
L70:
	;
	if v178&int32(3) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v215 = v174
	goto L74
L72:
	;
	v230 = v174
	goto L73
L73:
	;
	if base.Ui32(v230) <= base.Ui32(int32(3)) {
		v250 = v230
		goto L69
	} else {
		goto L78
	}
L74:
	;
	if v215 == int32(0) {
		goto L49
	} else {
		goto L76
	}
L75:
	;
	v230 = v221
	goto L73
L76:
	;
	v221 = v215 - int32(1)
	v222 = v165 + v221
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v222))) = uint8(v224)
	if v222&int32(3) != 0 {
		v215 = v221
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v237 = v230
	goto L79
L79:
	;
	v241 = v237 - int32(4)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v168+v241)))
	*(*int32)(unsafe.Add(mBase, uint32(v165+v241))) = v244
	if base.Ui32(int32(3)) < base.Ui32(v241) {
		v237 = v241
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v250 = v241
	goto L69
L81:
	;
	goto L80
L82:
	;
	v257 = v250
	goto L83
L83:
	;
	v261 = v257 - int32(1)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165+v261))) = uint8(v264)
	if v261 != 0 {
		v257 = v261
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L49
L85:
	;
	goto L84
L86:
	;
	v274 = v267
	v275 = v268
	v276 = v269
	goto L87
L87:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v278
	v280 = int32(4)
	v281 = v274 + v280
	v283 = v276 + v280
	v285 = v275 - v280
	if base.Ui32(int32(3)) < base.Ui32(v285) {
		v274 = v281
		v275 = v285
		v276 = v283
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v289 = v281
	v290 = v285
	v291 = v283
	goto L54
L89:
	;
	goto L88
L90:
	;
	v296 = v289
	v297 = v290
	v298 = v291
	goto L91
L91:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v300)
	v302 = int32(1)
	v307 = v297 - v302
	if v307 != 0 {
		v296 = v296 + v302
		v297 = v307
		v298 = v298 + v302
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L49
L93:
	;
	goto L92
L94:
	;
	v329 = v319
	v332 = v325
	goto L95
L95:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v342 = v339 + v329*int32(24)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+4)))
	if v343 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L1
L97:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+8)) = v346 - v101
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v350 = v349
	goto L99
L98:
	;
	v350 = v332
	goto L99
L99:
	;
	v352 = v329 + int32(1)
	if v352 < v350 {
		v329 = v352
		v332 = v350
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
}
