package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistGetBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	v10 = F_spgGetCache(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui32(l2) < base.Ui32(int32(_a_F_SpGistGetBuffer_0)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v234)
	v242 = int32(3)
	v243 = l1 & v242
	v245 = base.B2i32(v243 == v242)
	if v243 == v242 {
		goto L65
	} else {
		goto L66
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L62
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v23 = base.I32_div_s(int32(_a_F_SpGistGetBuffer_1)-v18<<(uint(int32(13))%32), int32(100))
	v25 = v23
	goto L9
L8:
	;
	v25 = int32(1638)
	goto L9
L9:
	;
	v30 = v10 + l1&int32(7)<<(uint(int32(3))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	if v31 == int32(-1) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(_a_F_SpGistGetBuffer_2)
	v35 = l2 + v25
	if base.Ui32(v34) <= base.Ui32(v35) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v34
	goto L13
L12:
	;
	v38 = v35
	goto L13
L13:
	;
	v40 = v30 - int32(-64)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v38 <= v41 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = F_ReadBuffer(m, l0, v31)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L3
L17:
	;
	v45 = F_ConditionalLockBuffer(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v45 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_ReleaseBuffer(m, v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v43 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L3
L23:
	;
	v178 = int32(4)
	v180 = int32(0)
	v191 = int32(3)
	if base.B2i32(l1&v178 == v180)^base.B2i32(v74&int32(8) == v180)|base.B2i32(base.B2i32(v74&v178 == v180) == base.B2i32(l1&v191 == v191)) != 0 {
		goto L54
	} else {
		goto L55
	}
L24:
	;
	v87 = int32(3)
	if l1&v87 == v87 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+14)))
	if v69 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L26:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v43^int32(-1))<<(uint(int32(2))%32))))
	v68 = v60
	goto L25
L27:
	;
	goto L28
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[1]))
	v68 = v62 + v43<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+16)))
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v72))))
	if v74&int32(2) != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v77) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v91 = int32(4)
	goto L34
L33:
	;
	v91 = int32(0)
	goto L34
L34:
	;
	v92 = l1<<(uint(int32(1))%32)&int32(8) | v91
	if v43 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v111 = int32(_a_F_SpGistGetBuffer_3)
	v113 = int32(0)
	if v113|(v110&int32(3)|int32(1)) == v113 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+(v43^int32(-1))<<(uint(int32(2))%32))))
	v110 = v102
	goto L35
L37:
	;
	goto L38
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[1]))
	v110 = v104 + v43<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+16)))
	v162 = v110 + v161
	v163 = int32(_a_F_SpGistGetBuffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v162)+6)) = uint16(v163)
	*(*uint16)(unsafe.Add(mBase, uint32(v162))) = uint16(v92)
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+14)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	v168 = v166 - v167
	v169 = int32(0)
	if v169 < v168 {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+10)) = int32(_a_F_SpGistGetBuffer_5)
	v152 = int32(_a_F_SpGistGetBuffer_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+18)) = uint16(v152)
	v158 = int32(_a_F_SpGistGetBuffer_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+16)) = uint16(v158)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+14)) = uint16(v158)
	goto L39
L41:
	;
	goto L44
L42:
	;
	goto L43
L43:
	;
	goto L49
L44:
	;
	v129 = v110 + v111
	v131 = v110 + int32(4)
	if base.Ui32(v131) < base.Ui32(v129) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = v129
	goto L47
L46:
	;
	v133 = v131
	goto L47
L47:
	;
	v138 = (v110^int32(-1)+v133)&int32(-4) + int32(4)
	if v138 == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	base.MemoryFill(m, v110, int32(0), v138)
	goto L40
L49:
	;
	base.MemoryFill(m, v110, int32(0), v111)
	goto L40
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v172 - v38
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v175)
	return v43
L51:
	;
	v172 = v168
	goto L53
L52:
	;
	v172 = v169
	goto L53
L53:
	;
	goto L50
L54:
	;
	F_UnlockReleaseBuffer(m, v43)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L55:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+14)))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	v199 = v197 - v198
	v200 = int32(0)
	if v200 < v199 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v203 < v38 {
		goto L54
	} else {
		goto L60
	}
L57:
	;
	v203 = v199
	goto L59
L58:
	;
	v203 = v200
	goto L59
L59:
	;
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v203 - v38
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v207)
	return v43
L61:
	;
	goto L16
L62:
	;
	F_errmsg_internal(m, int32(_a_F_SpGistGetBuffer_8), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_SpGistGetBuffer_9), int32(576), int32(_a_F_SpGistGetBuffer_10))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v246 = int32(4)
	goto L67
L66:
	;
	v246 = int32(0)
	goto L67
L67:
	;
	v247 = l1<<(uint(v234)%32)&int32(8) | v246
	v248 = F_spgGetCache(m, l0)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L69
L69:
	;
	v261 = F_SpGistNewBuffer(m, l0)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	return v261
L71:
	;
	v283 = int32(_a_F_SpGistGetBuffer_3)
	v285 = int32(0)
	if v285|(v282&int32(3)|int32(1)) == v285 {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v263 = int32(0)
	v264 = base.B2i32(v263 <= v261)
	if v264 == v263 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[0]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268+(v261^int32(-1))<<(uint(int32(2))%32))))
	v282 = v274
	goto L71
L74:
	;
	goto L75
L75:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[1]))
	v282 = v276 + v261<<(uint(int32(13))%32) + int32(-8192)
	goto L71
L76:
	;
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v334 = v282 + v333
	v335 = int32(_a_F_SpGistGetBuffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v334)+6)) = uint16(v335)
	*(*uint16)(unsafe.Add(mBase, uint32(v334))) = uint16(v247)
	if v243 == v242 {
		goto L87
	} else {
		goto L88
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v282)+10)) = int32(_a_F_SpGistGetBuffer_5)
	v324 = int32(_a_F_SpGistGetBuffer_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+18)) = uint16(v324)
	v330 = int32(_a_F_SpGistGetBuffer_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)) = uint16(v330)
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+14)) = uint16(v330)
	goto L76
L78:
	;
	goto L81
L79:
	;
	goto L80
L80:
	;
	goto L86
L81:
	;
	v301 = v282 + v283
	v303 = v282 + int32(4)
	if base.Ui32(v303) < base.Ui32(v301) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v305 = v301
	goto L84
L83:
	;
	v305 = v303
	goto L84
L84:
	;
	v310 = (v282^int32(-1)+v305)&int32(-4) + int32(4)
	if v310 == int32(0) {
		goto L77
	} else {
		goto L85
	}
L85:
	;
	base.MemoryFill(m, v282, int32(0), v310)
	goto L77
L86:
	;
	base.MemoryFill(m, v282, int32(0), v283)
	goto L77
L87:
	;
	goto L70
L88:
	;
	if v261 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v358 = base.I32_rem_u_s(v356, int32(3))
	if v243 == v358 {
		goto L87
	} else {
		goto L93
	}
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[2]))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341+(v261^int32(-1))<<(uint(int32(6))%32))+16))
	v356 = v347
	goto L89
L91:
	;
	goto L92
L92:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[3]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v349+v261<<(uint(int32(6))%32)+int32(-64))+16))
	v356 = v355
	goto L89
L93:
	;
	if v247 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v362 = v358 | int32(4)
	goto L96
L95:
	;
	v362 = v358
	goto L96
L96:
	;
	v365 = v248 - int32(-64) + v362<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v356
	if v264 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+14)))
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+12)))
	v387 = v385 - v386
	v388 = int32(0)
	if v388 < v387 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[0]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v370+(v261^int32(-1))<<(uint(int32(2))%32))))
	v384 = v376
	goto L97
L99:
	;
	goto L100
L100:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistGetBuffer[1]))
	v384 = v378 + v261<<(uint(int32(13))%32) + int32(-8192)
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v391
	F_UnlockReleaseBuffer(m, v261)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L105
	}
L102:
	;
	v391 = v387
	goto L104
L103:
	;
	v391 = v388
	goto L104
L104:
	;
	goto L101
L105:
	;
	goto L69
}
func F_SpGistNewBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v85
L2:
	;
	return int32(0)
L3:
	;
	if v9 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v75
	v77 = int32(8)
	v79 = int32(0)
	v82 = F_ExtendBufferedRel(m, v7+v77, v79, v79, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L28
	}
L7:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v16) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v21 = F_ReadBuffer(m, l0, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v62 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L26
	}
L12:
	;
	v23 = F_ConditionalLockBuffer(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v21 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	F_ReleaseBuffer(m, v21)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L25
	}
L17:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)))
	if v43 == int32(0) {
		v85 = v21
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistNewBuffer[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v21^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L17
L19:
	;
	goto L20
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistNewBuffer[1]))
	v42 = v36 + v21<<(uint(int32(13))%32) + int32(-8192)
	goto L17
L21:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v46))))
	if v48&int32(2) != 0 {
		v85 = v21
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+12)))
	if base.Ui32(v51) < base.Ui32(int32(25)) {
		v85 = v21
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_LockBuffer(m, v21, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	goto L11
L26:
	;
	if v62 != int32(-1) {
		v16 = v62
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v85 = v82
	goto L1
}
