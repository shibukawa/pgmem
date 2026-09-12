package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_int4(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int4(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_int8(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_int8(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_generate_series_step_numeric(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_generate_series_numeric_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int64
	_ = v372
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 float64
	_ = v419
	var v420 int32
	_ = v420
	var v425 float64
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 != int32(460) {
		v440 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return v440
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v18 == int32(0) {
		v440 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(15) {
		v440 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = F_estimate_expression_value(m, v24, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v35 = F_estimate_expression_value(m, v32, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(3) <= v38 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = F_estimate_expression_value(m, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v47 == int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v46 = v44
	goto L10
L12:
	;
	v440 = v432
	goto L1
L13:
	;
	v81 = *(*int64)(unsafe.Add(mBase, _consts[1072]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v81
	v84 = *(*int64)(unsafe.Add(mBase, _consts[1073]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v84
	v87 = *(*int64)(unsafe.Add(mBase, _consts[1074]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v90 = F_pg_detoast_datum(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L33
	}
L14:
	;
	v74 = int32(0)
	if v47 != int32(7) {
		v440 = v74
		goto L1
	} else {
		goto L31
	}
L15:
	;
	v66 = int32(0)
	if v47 != int32(7) {
		v440 = v66
		goto L1
	} else {
		goto L28
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(0)
	v432 = v14
	goto L12
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v50 != 0 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v51 == int32(7) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v54 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v46 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v57 != int32(7) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+24)))
	if v60 != int32(1) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L16
L28:
	;
	if v51 != int32(7) {
		v440 = v66
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v71 == int32(7) {
		v79 = v66
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v440 = v66
	goto L1
L31:
	;
	if v51 != int32(7) {
		v440 = v74
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v79 = v74
	goto L13
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v93 = F_pg_detoast_datum(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)))
	if base.Ui32(int32(49151)) < base.Ui32(v95) {
		v440 = v79
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	if base.Ui32(int32(49151)) < base.Ui32(v98) {
		v440 = v79
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v46 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v102 = F_pg_detoast_datum(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v180 = v12 + int32(72)
	v181 = int32(1770548)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[1075]))
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1076]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v190 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L40:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	if base.Ui32(int32(49151)) < base.Ui32(v104) {
		v440 = v79
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v108 = v12 + int32(72)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
	if int32(0) <= v116 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L39
L43:
	;
	v119 = int32(-8)
	goto L45
L44:
	;
	v119 = int32(-6)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(base.Ui32(int32(base.Ui32(v111)>>(uint(int32(2))%32))+v119) >> (uint(int32(1)) % 32))
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
	if v124 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v140
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
	v143 = int32(49152)
	v144 = v142 & v143
	if v144 != v143 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v128 = v124 & int32(65535)
	v140 = v128<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v128&int32(63)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+6)))
	v140 = v138
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v155
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
	if v157 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v144 != int32(32768) {
		v155 = v144
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v155 = v142 & int32(61440)
	goto L50
L54:
	;
	v155 = v142 << (uint(int32(1)) % 32) & int32(16384)
	goto L50
L55:
	;
	v166 = int32(base.Ui32(v157)>>(uint(int32(7))%32)) & int32(63)
	goto L57
L56:
	;
	v166 = v157 & int32(16383)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+12)) = v166
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v169
	if v168 < v169 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = int32(6)
	goto L60
L59:
	;
	v175 = int32(8)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v102 + v175
	goto L42
L61:
	;
	if v226 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	if v189 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	if v189 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v226 = int32(0)
	goto L61
L66:
	;
	goto L67
L67:
	;
	if v188 == int32(16384) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v200 = int32(1)
	goto L70
L69:
	;
	v200 = int32(-1)
	goto L70
L70:
	;
	v226 = v200
	goto L61
L71:
	;
	if v201 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[1077]))
	v208 = *(*int32)(unsafe.Add(mBase, _consts[1078]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	if v201 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v206 = int32(-1)
	goto L76
L75:
	;
	v206 = int32(1)
	goto L76
L76:
	;
	v226 = v206
	goto L61
L77:
	;
	if v188 == int32(16384) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if v188 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v226 = int32(1)
	goto L61
L81:
	;
	goto L82
L82:
	;
	v216 = F_cmp_abs_common(m, v210, v190, v209, v208, v189, v207)
	mBase = m.M
	v226 = v216
	goto L61
L83:
	;
	v226 = int32(-1)
	goto L61
L84:
	;
	goto L85
L85:
	;
	v220 = F_cmp_abs_common(m, v208, v189, v207, v210, v190, v209)
	mBase = m.M
	v226 = v220
	goto L61
L86:
	;
	v432 = int32(0)
	goto L12
L87:
	;
	goto L88
L88:
	;
	v231 = v12 + int32(48)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v239 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
	if int32(0) <= v239 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v302 = v12 + int32(24)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v310 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+4)))
	if int32(0) <= v310 {
		goto L109
	} else {
		goto L110
	}
L90:
	;
	v242 = int32(-8)
	goto L92
L91:
	;
	v242 = int32(-6)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = int32(base.Ui32(int32(base.Ui32(v234)>>(uint(int32(2))%32))+v242) >> (uint(int32(1)) % 32))
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
	if v247 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v263
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)))
	v266 = int32(49152)
	v267 = v265 & v266
	if v267 != v266 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v251 = v247 & int32(65535)
	v263 = v251<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v251&int32(63)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+6)))
	v263 = v261
	goto L93
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = v278
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
	if v280 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	if v267 != int32(32768) {
		v278 = v267
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v278 = v265 & int32(61440)
	goto L97
L101:
	;
	v278 = v265 << (uint(int32(1)) % 32) & int32(16384)
	goto L97
L102:
	;
	v289 = int32(base.Ui32(v280)>>(uint(int32(7))%32)) & int32(63)
	goto L104
L103:
	;
	v289 = v280 & int32(16383)
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+12)) = v289
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
	v292 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v292
	if v291 < v292 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v298 = int32(6)
	goto L107
L106:
	;
	v298 = int32(8)
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = v90 + v298
	goto L89
L108:
	;
	v372 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v372
	v375 = v12 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v375))) = v372
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v372
	F_sub_var(m, v12+int32(24), v12+int32(48), v12)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L127
	}
L109:
	;
	v313 = int32(-8)
	goto L111
L110:
	;
	v313 = int32(-6)
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = int32(base.Ui32(int32(base.Ui32(v305)>>(uint(int32(2))%32))+v313) >> (uint(int32(1)) % 32))
	v318 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+4)))
	if v318 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+4)) = v334
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	v337 = int32(49152)
	v338 = v336 & v337
	if v338 != v337 {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v322 = v318 & int32(65535)
	v334 = v322<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v322&int32(63)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v332 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+6)))
	v334 = v332
	goto L112
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+8)) = v349
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+4)))
	if v351 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	if v338 != int32(32768) {
		v349 = v338
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v349 = v336 & int32(61440)
	goto L116
L120:
	;
	v349 = v336 << (uint(int32(1)) % 32) & int32(16384)
	goto L116
L121:
	;
	v360 = int32(base.Ui32(v351)>>(uint(int32(7))%32)) & int32(63)
	goto L123
L122:
	;
	v360 = v351 & int32(16383)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+12)) = v360
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+4)))
	v363 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v302)+16)) = v363
	if v362 < v363 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v369 = int32(6)
	goto L126
L125:
	;
	v369 = int32(8)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+20)) = v93 + v369
	goto L108
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	if v386 == v387 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v46 != 0 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v425 = float64(0)
	goto L130
L130:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v427 == int32(0) {
		v432 = v14
		goto L12
	} else {
		goto L143
	}
L131:
	;
	v419 = F_numericvar_to_double_no_overflow(m, v12)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L5
	} else {
		goto L142
	}
L132:
	;
	v391 = int32(0)
	F_div_var(m, v12, v12+int32(72), v12, v391, v391, v391)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v397 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v397
	v400 = v396 << (uint(int32(2)) % 32)
	if v400+int32(4) <= v397 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L131
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(0)
	goto L131
L137:
	;
	goto L138
L138:
	;
	v412 = base.I32_div_s(v400+int32(7), int32(4))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v412 < v413 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v415 = v412
	goto L141
L140:
	;
	v415 = v413
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v415
	goto L131
L142:
	;
	v425 = base.F64_add(v419, float64(1))
	goto L130
L143:
	;
	F_pfree(m, v427)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	v432 = v14
	goto L12
}
