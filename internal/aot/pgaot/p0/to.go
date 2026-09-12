package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddToDataDirLockFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16512)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v3
	v22 = F_open(m, int32(454328), int32(2), v14+int32(112))
	mBase = m.M
	if v22 < v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16512)
	return
L2:
	;
	v27 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v43 = int32(4164524)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(167772187)
	v50 = F_read(m, v22, v14+int32(8320), int32(8191))
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	if v50 < v53 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return
L6:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(454328)
	F_errmsg(m, int32(313019), v14)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(517023), int32(1587), int32(408789))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	v59 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v77 = v14 + int32(8320)
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v50))) = uint8(v79)
	v83 = int32(1)
	if l0 < int32(2) {
		v129 = v77
		v130 = v83
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v75 = F_close(m, v22)
	mBase = m.M
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(454328)
	F_errmsg(m, int32(313085), v14+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(517023), int32(1598), int32(408789))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = l1
	v242 = v14 + int32(8320)
	v247 = F_pg_snprintf(m, v231, v242-v231, int32(782775), v14+int32(96))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L67
	}
L22:
	;
	v138 = v14 + int32(8320)
	v141 = v129 - v138
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	v91 = v77
	v92 = v83
	goto L24
L24:
	;
	v97 = int32(10)
	v98 = F___strchrnul(m, v91, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == v97 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v115 = v14 + int32(8320)
	v118 = v108 - v115
	if v118 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v104 == int32(0) {
		v129 = v91
		v130 = v92
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v104 = v98
	goto L29
L28:
	;
	v104 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	v107 = int32(1)
	v108 = v104 + v107
	v110 = v92 + v107
	if v110 != l0 {
		v91 = v108
		v92 = v110
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	v231 = v14 + int32(128) + v118
	v234 = v108
	goto L21
L33:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v14+int32(128), v115, v118)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v146 = v14 + int32(128) + v141
	if l0 <= v130 {
		v231 = v146
		v234 = v129
		goto L21
	} else {
		goto L40
	}
L37:
	;
	v142 = F__emscripten_memcpy_bulkmem(m, v14+int32(128), v138, v141)
	mBase = m.M
	goto L39
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	v149 = v14 + int32(8320)
	v152 = (l0 - v130) & int32(3)
	if v152 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v130-l0) {
		v231 = v179
		v234 = v129
		goto L21
	} else {
		goto L51
	}
L42:
	;
	v179 = v146
	v184 = v130
	goto L41
L43:
	;
	goto L44
L44:
	;
	v157 = v146
	v162 = v130
	v165 = v3
	goto L45
L45:
	;
	if base.Ui32(v157) < base.Ui32(v149) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v179 = v171
	v184 = v173
	goto L41
L47:
	;
	v167 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v167)
	v171 = v157 + int32(1)
	goto L49
L48:
	;
	v171 = v157
	goto L49
L49:
	;
	v172 = int32(1)
	v173 = v162 + v172
	v175 = v165 + v172
	if v175 != v152 {
		v157 = v171
		v162 = v173
		v165 = v175
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v193 = v179
	v198 = v184
	goto L52
L52:
	;
	if base.Ui32(v193) < base.Ui32(v149) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v231 = v225
	v234 = v129
	goto L21
L54:
	;
	v203 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v203)
	v207 = v193 + int32(1)
	goto L56
L55:
	;
	v207 = v193
	goto L56
L56:
	;
	if base.Ui32(v207) < base.Ui32(v149) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v209 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v209)
	v213 = v207 + int32(1)
	goto L59
L58:
	;
	v213 = v207
	goto L59
L59:
	;
	if base.Ui32(v213) < base.Ui32(v149) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v215 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v215)
	v219 = v213 + int32(1)
	goto L62
L61:
	;
	v219 = v213
	goto L62
L62:
	;
	if base.Ui32(v219) < base.Ui32(v149) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v221 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v221)
	v225 = v219 + int32(1)
	goto L65
L64:
	;
	v225 = v219
	goto L65
L65:
	;
	v227 = v198 + int32(4)
	if v227 != l0 {
		v193 = v225
		v198 = v227
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v249 = int32(10)
	v250 = F___strchrnul(m, v234, v249)
	mBase = m.M
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v252 == v249 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v256 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v256 = v250
	goto L71
L70:
	;
	v256 = int32(0)
	goto L71
L71:
	;
	goto L68
L72:
	;
	v257 = F_strlen(m, v231)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v256 + int32(1)
	v261 = v231 + v257
	v266 = F_pg_snprintf(m, v261, v242-v261, int32(216894), v14+int32(80))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v271 = v14 + int32(128)
	v272 = F_strlen(m, v271)
	mBase = m.M
	v274 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v274
	v276 = int32(4164524)
	v277 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(167772189)
	v283 = F_pwrite(m, v22, v271, v272, int64(0))
	mBase = m.M
	v285 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v274
	if v272 != v283 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v290 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = int32(167772188)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, _consts[159])))
	if v321 != int32(1) {
		v335 = int32(0)
		goto L91
	} else {
		goto L92
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(51)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v298 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if v298 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v314 = F_close(m, v22)
	mBase = m.M
	goto L1
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(454328)
	F_errmsg(m, int32(312702), v14-int32(-64))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(517023), int32(1662), int32(408789))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v360 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360
	v362 = F_close(m, v22)
	mBase = m.M
	if v362 == v360 {
		goto L1
	} else {
		goto L103
	}
L90:
	;
	if v335 == int32(0) {
		goto L89
	} else {
		goto L97
	}
L91:
	;
	goto L90
L92:
	;
	goto L93
L93:
	;
	v326 = F_fsync(m, v22)
	mBase = m.M
	if v326 != int32(-1) {
		v335 = v326
		goto L91
	} else {
		goto L95
	}
L94:
	;
	v335 = int32(-1)
	goto L91
L95:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v330 == int32(27) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v340 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if v340 == int32(0) {
		goto L89
	} else {
		goto L99
	}
L99:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(454328)
	F_errmsg(m, int32(312702), v14+int32(48))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(517023), int32(1673), int32(408789))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	goto L89
L103:
	;
	v367 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	if v367 == int32(0) {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(454328)
	F_errmsg(m, int32(312702), v14+int32(32))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(517023), int32(1681), int32(408789))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	goto L1
}
func F_CopyToBinaryEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v9, v5+int32(14), int32(2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_CopySendEndOfRow(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_CopyToCSVOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L28
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v49 = int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v50 <= v49 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = F_OutputFunctionCall(m, v18+v22*int32(28), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v45 = F_strlen(m, v44)
	mBase = m.M
	F_appendBinaryStringInfo(m, v43, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v22))))
	F_CopyAttributeOutCSV(m, l0, v36, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	goto L4
L12:
	;
	v55 = v49
	goto L13
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v65 = int32(2)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v55<<(uint(v65)%32))))
	v69 = int32(1)
	v70 = v68 - v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v71))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v70<<(uint(v65)%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v85 <= v82+v69 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L1
L15:
	;
	if v73&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	F_appendStringInfoChar(m, v81, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v89+v82))) = uint8(v80)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v95 = v93 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97+v95))) = uint8(v99)
	goto L15
L19:
	;
	goto L15
L20:
	;
	v122 = v55 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v122 < v123 {
		v55 = v122
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v107 = F_strlen(m, v106)
	mBase = m.M
	F_appendBinaryStringInfo(m, v105, v106, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v113 = F_OutputFunctionCall(m, v18+v70*int32(28), v78)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v70))))
	F_CopyAttributeOutCSV(m, l0, v113, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L14
L28:
	;
	return
}
func F_CopyToTextLikeOutFunc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_getTypeOutputInfo(m, l1, v6+int32(12), v6+int32(11))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		F_fmgr_info(m, v14, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_CopyToTextLikeStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = F_pg_server_to_any(m, v13, v14, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v16
	goto L3
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L33
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = l1 - int32(76)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v36 = v27 + v28<<(uint(int32(4))%32) + v33*int32(100)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v37&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v48 <= v47 {
		goto L9
	} else {
		goto L18
	}
L13:
	;
	F_CopyAttributeOutText(m, l0, v36)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_CopyAttributeOutCSV(m, l0, v36, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	goto L12
L18:
	;
	v54 = v47
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v54<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v71 <= v68+int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L9
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v95 = v27 + v89<<(uint(int32(4))%32) + v64*int32(100)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+54)))
	if v96 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_appendStringInfoChar(m, v67, v66)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v68))) = uint8(v66)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = v79 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v81))) = uint8(v85)
	goto L21
L25:
	;
	goto L21
L26:
	;
	v105 = v54 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v105 < v106 {
		v54 = v105
		goto L19
	} else {
		goto L32
	}
L27:
	;
	F_CopyAttributeOutCSV(m, l0, v95, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_CopyAttributeOutText(m, l0, v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	goto L26
L32:
	;
	goto L20
L33:
	;
	goto L8
}
func F_coerce_to_boolean(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_exprType(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(16) {
			v26 = l1
			v27 = F_expression_returns_set(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg(m, int32(114084), v9)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = F_exprLocation(m, v26)
								mBase = m.M
								F_parser_errposition(m, l0, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(524516), int32(1190), int32(297530))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					m.G0 = v9 + int32(32)
					return v26
				}
			}
		} else {
			v18 = int32(-1)
			v22 = F_coerce_to_target_type(m, l0, l1, v11, int32(16), v18, int32(1), int32(2), v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_format_type_be(m, v11)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(297776)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
								F_errmsg(m, int32(199400), v9+int32(16))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = F_exprLocation(m, l1)
									mBase = m.M
									F_parser_errposition(m, l0, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(524516), int32(1180), int32(297530))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
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
				} else {
					v26 = v22
					v27 = F_expression_returns_set(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
									F_errmsg(m, int32(114084), v9)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = F_exprLocation(m, v26)
										mBase = m.M
										F_parser_errposition(m, l0, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(524516), int32(1190), int32(297530))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							m.G0 = v9 + int32(32)
							return v26
						}
					}
				}
			}
		}
	}
}
func F_encode_to_ascii(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(160)
	switch l1 - int32(8) {
	case 0:
		v51 = v15
		v52 = int32(8169)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 1:
		v51 = v15
		v52 = int32(606035)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if base.Ui32(l1) <= base.Ui32(int32(41)) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(3))%32))+uint32(_consts[768])))
					v40 = v39
				} else {
					v40 = int32(791891)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v40
				F_errmsg(m, int32(465203), v12)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(522732), int32(78), int32(336014))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 8:
		v51 = v15
		v52 = int32(8072)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	case 21:
		v51 = int32(128)
		v52 = int32(765821)
		v54 = l0 + int32(4)
		v56 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v57 = l0 + v56
		if base.Ui32(v57) <= base.Ui32(v54) {
		} else {
			if v14&int32(4) != 0 {
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				v62 = base.I32_extend8_s(v61)
				if int32(0) <= v62 {
					v70 = v62
				} else {
					if base.Ui32(v61) < base.Ui32(v51) {
						v70 = int32(32)
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v61-v51)))))
						v70 = v69
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v70)
				v74 = l0 + int32(5)
			} else {
				v74 = v54
			}
			if v56 == int32(5) {
			} else {
				v79 = v74
				for {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
					v88 = base.I32_extend8_s(v87)
					if int32(0) <= v88 {
						v96 = v88
					} else {
						if base.Ui32(v87) < base.Ui32(v51) {
							v96 = int32(32)
						} else {
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v87-v51)))))
							v96 = v95
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v96)
					v99 = v79 + int32(1)
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
					v101 = base.I32_extend8_s(v100)
					if int32(0) <= v101 {
						v109 = v101
					} else {
						if base.Ui32(v100) < base.Ui32(v51) {
							v109 = int32(32)
						} else {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v100-v51)))))
							v109 = v108
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v109)
					v112 = v79 + int32(2)
					if v112 != v57 {
						v79 = v112
						continue
					} else {
						break
					}
					break
				}
			}
		}
		m.G0 = v12 + int32(16)
		return l0
	}
}
func F_to_oct64(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v14 = v9 - v8
	v15 = v14
	v16 = v12
	goto L1
L1:
	;
	v22 = v15 - int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)&int32(7))+uint32(_consts[863]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v28)
	if base.Ui64(v16) < base.Ui64(int64(8)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v14 - v22
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v22) {
		v15 = v22
		v16 = int64(base.Ui64(v16) >> (uint(int64(3)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v39
L9:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v39+int32(4), v22, v36)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_regclass(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[842]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[843]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1500), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regtype(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[842]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[843]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1257), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_tsquery_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v17 = F_text_to_cstring(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v22 = int32(0)
			v24 = F_parse_tsquery(m, v17, int32(1177), v6+int32(8), v22, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func F_to_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1175), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
