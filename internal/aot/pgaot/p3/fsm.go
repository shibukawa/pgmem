package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fsm_search(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v25 = int32(0)
	v27 = int64(2)
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v27
	v32 = base.I32_wrap_i64(int64(base.Ui64(v27) >> (uint(int64(32)) % 64)))
	v33 = base.I32_wrap_i64(v27)
	v34 = int32(0)
	v38 = F_fsm_readbuf(m, l0, v14+int32(24), v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	m.G0 = v14 + int32(48)
	return v435
L3:
	;
	goto L2
L4:
	;
	v327 = v321 + int32(28)
	v329 = v32 - v82*int32(4069) + int32(4095)
	v330 = v327 + v329
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v331 != v73 {
		goto L87
	} else {
		goto L88
	}
L5:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[0]))
	v321 = v315 + v90<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L6:
	;
	F_LockBuffer(m, v38, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L24
	}
L7:
	;
	return int32(0)
L8:
	;
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_LockBuffer(m, v38, int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v73 = v34
	goto L11
L11:
	;
	v74 = int32(-1)
	if v33 == int32(2) {
		v435 = v74
		goto L3
	} else {
		goto L20
	}
L12:
	;
	v45 = int32(0)
	v48 = F_fsm_search_avail(m, v38, l1, base.B2i32(v33 == v45), v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v48 != int32(-1) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v38 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+28)))
	F_UnlockReleaseBuffer(m, v38)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L19
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[1]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v38^int32(-1))<<(uint(int32(2))%32))))
	v69 = v61
	goto L15
L17:
	;
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[0]))
	v69 = v63 + v38<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	v73 = v70
	goto L11
L20:
	;
	v82 = base.I32_div_u_s(v32, int32(4069))
	v86 = (v27+int64(1))&int64(4294967295) | base.I64_extend_i32_u(v82)<<(uint(int64(32))%64)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v86
	v90 = F_fsm_readbuf(m, l0, v14, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if int32(0) <= v90 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[1]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+(v90^int32(-1))<<(uint(int32(2))%32))))
	v321 = v107
	goto L4
L24:
	;
	v112 = v48 & int32(_a_F_fsm_search_0)
	if v33 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v25 = v306
	v27 = base.I64_extend_i32_u(v308+v112)<<(uint(int64(32))%64) | v307
	goto L1
L26:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v117 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	F_ReleaseBuffer(m, v38)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L85
	}
L29:
	;
	if v38 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v143 = v117
	goto L32
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v121
	v125 = F_smgropen(m, v14+int32(8), v118)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L33
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v147 = v32*int32(4069) + v112
	if base.B2i32(v144 != int32(-1))&base.B2i32(base.Ui32(v147) < base.Ui32(v144)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v125
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+72))
	if v129 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v143 = v141
	goto L32
L35:
	;
	v137 = v129
	goto L37
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+72))
	v137 = v135
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+72)) = v137 + int32(1)
	goto L34
L38:
	;
	v153 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_ReleaseBuffer(m, v38)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v153) <= base.Ui32(v147) {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v435 = v147
	goto L3
L44:
	;
	F_LockBuffer(m, v38, int32(2))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L48
	}
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[1]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v38^int32(-1))<<(uint(int32(2))%32))))
	v175 = v167
	goto L44
L46:
	;
	goto L47
L47:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_search[0]))
	v175 = v169 + v38<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L48:
	;
	v179 = int32(0)
	v185 = v175 + int32(28)
	v187 = v48 + int32(4095)
	v188 = v185 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v189 != v179 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	F_MarkBufferDirtyHint(m, v38, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L80
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v179)
	v198 = v187
	goto L53
L51:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32(v191) < base.Ui32(v179) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v202 = int32(1)
	v203 = v198 - v202
	v204 = int32(2)
	v205 = base.I32_div_s(v203, v204)
	v207 = v205 << (uint(v202) % 32)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v207)+1)))
	v211 = v207 + v204
	if base.Ui32(v211) <= base.Ui32(int32(_a_F_fsm_search_1)) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32(v230) < base.Ui32(v179) {
		goto L65
	} else {
		goto L66
	}
L55:
	;
	v215 = v209 & int32(255)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v211))))
	if base.Ui32(v217) < base.Ui32(v215) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v220 = v209
	goto L57
L57:
	;
	v222 = v205 + v185
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v223 != v220&int32(255) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v219 = v215
	goto L60
L59:
	;
	v219 = v217
	goto L60
L60:
	;
	v220 = v219
	goto L57
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v222))) = uint8(v220)
	if int32(1) < v203 {
		v198 = v205
		goto L53
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L54
L64:
	;
	goto L63
L65:
	;
	v236 = int32(4094)
	goto L68
L66:
	;
	goto L67
L67:
	;
	goto L49
L68:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v236) {
		v257 = int32(0)
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v258 = v236 + v185
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v259 != v257&int32(255) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v246 = v236 << (uint(int32(1)) % 32)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(29)+v246))))
	if v236 == int32(4081) {
		v257 = v248
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v185)+2)))
	if base.Ui32(v252) < base.Ui32(v248) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v254 = v248
	goto L75
L74:
	;
	v254 = v252
	goto L75
L75:
	;
	v257 = v254
	goto L70
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v257)
	goto L78
L77:
	;
	goto L78
L78:
	;
	if v236 != 0 {
		v236 = v236 - int32(1)
		goto L68
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	F_UnlockReleaseBuffer(m, v38)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	if int32(_a_F_fsm_search_2) < v25 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v435 = int32(-1)
	goto L3
L83:
	;
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(2)
	v306 = v25 + int32(1)
	v307 = int64(1)
	v308 = int32(0)
	goto L25
L85:
	;
	v306 = v25
	v307 = (v27 - int64(1)) & int64(4294967295)
	v308 = v32 * int32(4069)
	goto L25
L86:
	;
	if v423 != 0 {
		goto L117
	} else {
		goto L118
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v73)
	v340 = v329
	goto L90
L88:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if base.Ui32(v333) < base.Ui32(v73) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v423 = int32(0)
	goto L86
L90:
	;
	v344 = int32(1)
	v345 = v340 - v344
	v346 = int32(2)
	v347 = base.I32_div_s(v345, v346)
	v349 = v347 << (uint(v344) % 32)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327+v349)+1)))
	v353 = v349 + v346
	if base.Ui32(v353) <= base.Ui32(int32(_a_F_fsm_search_1)) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if base.Ui32(v372) < base.Ui32(v73) {
		goto L102
	} else {
		goto L103
	}
L92:
	;
	v357 = v351 & int32(255)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327+v353))))
	if base.Ui32(v359) < base.Ui32(v357) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v362 = v351
	goto L94
L94:
	;
	v364 = v347 + v327
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v365 != v362&int32(255) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v361 = v357
	goto L97
L96:
	;
	v361 = v359
	goto L97
L97:
	;
	v362 = v361
	goto L94
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v362)
	if int32(1) < v345 {
		v340 = v347
		goto L90
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	goto L91
L101:
	;
	goto L100
L102:
	;
	v378 = int32(4094)
	goto L105
L103:
	;
	goto L104
L104:
	;
	v423 = int32(1)
	goto L86
L105:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v378) {
		v399 = int32(0)
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L104
L107:
	;
	v400 = v378 + v327
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v401 != v399&int32(255) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v388 = v378 << (uint(int32(1)) % 32)
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+int32(29)+v388))))
	if v378 == int32(4081) {
		v399 = v390
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388+v327)+2)))
	if base.Ui32(v394) < base.Ui32(v390) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v396 = v390
	goto L112
L111:
	;
	v396 = v394
	goto L112
L112:
	;
	v399 = v396
	goto L107
L113:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v400))) = uint8(v399)
	goto L115
L114:
	;
	goto L115
L115:
	;
	if v378 != 0 {
		v378 = v378 - int32(1)
		goto L105
	} else {
		goto L116
	}
L116:
	;
	goto L106
L117:
	;
	F_MarkBufferDirtyHint(m, v90, int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	F_UnlockReleaseBuffer(m, v90)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	if int32(_a_F_fsm_search_2) < v25 {
		v435 = v74
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v25 = v25 + int32(1)
	v27 = int64(2)
	goto L1
}
