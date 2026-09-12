package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_crypt_md5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l2 == v5 {
		v594 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v594
L2:
	;
	if base.Ui32(l3) < base.Ui32(int32(120)) {
		v594 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(36) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v93 = F_px_find_digest(m, int32(548705), v13+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v23 != int32(49) {
		v86 = l1
		v89 = v5
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v34 = int32(0)
	v35 = v20
	goto L7
L7:
	;
	v36 = l1 + v34
	v38 = v35 & int32(255)
	if v38 == int32(0) {
		v86 = v36
		v89 = v5
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v28 == int32(36) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = int32(3)
	goto L11
L10:
	;
	v31 = int32(0)
	goto L11
L11:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v31))))
	v34 = v31
	v35 = v33
	goto L7
L12:
	;
	if v38 == int32(36) {
		v86 = v36
		v89 = v5
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(1)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v44 == int32(0) {
		v86 = v36
		v89 = v43
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v44 == int32(36) {
		v86 = v36
		v89 = v43
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(2)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
	if v50 == int32(0) {
		v86 = v36
		v89 = v49
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v50 == int32(36) {
		v86 = v36
		v89 = v49
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(3)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
	if v56 == int32(0) {
		v86 = v36
		v89 = v55
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(36) {
		v86 = v36
		v89 = v55
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(4)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v62 == int32(0) {
		v86 = v36
		v89 = v61
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v62 == int32(36) {
		v86 = v36
		v89 = v61
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(5)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
	if v68 == int32(0) {
		v86 = v36
		v89 = v67
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v68 == int32(36) {
		v86 = v36
		v89 = v67
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v73 = int32(6)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)))
	if v74 == int32(0) {
		v86 = v36
		v89 = v73
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(36) {
		v86 = v36
		v89 = v73
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v79 = int32(7)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+7)))
	if v80 == int32(0) {
		v86 = v36
		v89 = v79
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v80 == int32(36) {
		v86 = v36
		v89 = v79
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v86 = v36
	v89 = int32(8)
	goto L4
L28:
	;
	return int32(0)
L29:
	;
	if v93 != 0 {
		v594 = v5
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v100 = F_px_find_digest(m, int32(548705), v13+int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v100 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v105 = F_strlen(m, l0)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	m.T0[v106].(func(*base.Module, int32, int32, int32))(m, v102, l0, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L28
	} else {
		goto L35
	}
L33:
	;
	v576 = v102
	v581 = v5
	goto L34
L34:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v576)+20))
	m.T0[v583].(func(*base.Module, int32))(m, v576)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L28
	} else {
		goto L108
	}
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	m.T0[v112].(func(*base.Module, int32, int32, int32))(m, v109, int32(672895), int32(3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v116].(func(*base.Module, int32, int32, int32))(m, v115, v86, v89)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v120 = F_strlen(m, l0)
	mBase = m.M
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	m.T0[v121].(func(*base.Module, int32, int32, int32))(m, v119, l0, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	m.T0[v125].(func(*base.Module, int32, int32, int32))(m, v124, v86, v89)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v129 = F_strlen(m, l0)
	mBase = m.M
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	m.T0[v130].(func(*base.Module, int32, int32, int32))(m, v128, l0, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	m.T0[v136].(func(*base.Module, int32, int32))(m, v133, v13+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	v139 = F_strlen(m, l0)
	mBase = m.M
	if v139 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v209 = int32(0)
	v210 = int32(16)
	v214 = F___memset(m, v13+v210, v209, v210)
	mBase = m.M
	goto L63
L43:
	;
	if (v139-int32(1))&int32(16) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v149 = int32(16)
	if base.Ui32(v149) <= base.Ui32(v139) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v160 = v139
	goto L46
L46:
	;
	if base.Ui32(v139) < base.Ui32(int32(17)) {
		goto L42
	} else {
		goto L51
	}
L47:
	;
	v154 = v149
	goto L49
L48:
	;
	v154 = v139
	goto L49
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	m.T0[v155].(func(*base.Module, int32, int32, int32))(m, v148, v13+v149, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v160 = v139 - int32(16)
	goto L46
L51:
	;
	v166 = v160
	goto L52
L52:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v174 = int32(16)
	if base.Ui32(v174) <= base.Ui32(v166) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L42
L54:
	;
	v179 = v174
	goto L56
L55:
	;
	v179 = v166
	goto L56
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	m.T0[v180].(func(*base.Module, int32, int32, int32))(m, v173, v13+v174, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L28
	} else {
		goto L57
	}
L57:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v184 = int32(16)
	v188 = v166 - v184
	if base.Ui32(v184) <= base.Ui32(v188) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v191 = v184
	goto L60
L59:
	;
	v191 = v188
	goto L60
L60:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	m.T0[v192].(func(*base.Module, int32, int32, int32))(m, v183, v13+v184, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L28
	} else {
		goto L61
	}
L61:
	;
	if base.Ui32(int32(16)) < base.Ui32(v188) {
		v166 = v166 - int32(32)
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
L63:
	;
	v215 = F_strlen(m, l0)
	mBase = m.M
	if v215 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v220 = v215
	goto L67
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2371876)
	v252 = F_strlen(m, l2)
	mBase = m.M
	v253 = v252 + l2
	if v89 == int32(0) {
		v284 = v253
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v220&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L66
L69:
	;
	v231 = v13 + int32(16)
	goto L71
L70:
	;
	v231 = l0
	goto L71
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	m.T0[v233].(func(*base.Module, int32, int32, int32))(m, v226, v231, int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L28
	} else {
		goto L72
	}
L72:
	;
	v236 = int32(1)
	if base.Ui32(v236) < base.Ui32(v220) {
		v220 = v220 >> (uint(v236) % 32)
		goto L67
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v284))) = uint8(v286)
	v288 = F_strlen(m, l2)
	mBase = m.M
	v290 = int32(36)
	*(*uint16)(unsafe.Add(mBase, uint32(v288+l2))) = uint16(v290)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	m.T0[v295].(func(*base.Module, int32, int32))(m, v292, v13+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L28
	} else {
		goto L80
	}
L75:
	;
	v260 = v86
	v262 = v89
	v264 = v253
	goto L76
L76:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v266 == int32(0) {
		v284 = v264
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v284 = v271
	goto L74
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v264))) = uint8(v266)
	v270 = int32(1)
	v271 = v264 + v270
	v275 = v262 - v270
	if v275 != 0 {
		v260 = v260 + v270
		v262 = v275
		v264 = v271
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v301 = v209
	goto L81
L81:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	m.T0[v309].(func(*base.Module, int32))(m, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L28
	} else {
		goto L83
	}
L82:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)))
	v363 = F_strlen(m, l2)
	mBase = m.M
	v364 = v363 + l2
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
	v366 = int32(2)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v365)>>(uint(v366)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+3)) = uint8(v370)
	v372 = int32(63)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v376)
	v378 = int32(8)
	v379 = v361 << (uint(v378) % 32)
	v380 = int32(16)
	v383 = int32(12)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v379|v365<<(uint(v380)%32))>>(uint(v383)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+2)) = uint8(v389)
	v392 = int32(6)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v379|v362)>>(uint(v392)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)) = uint8(v398)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+29)))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+17)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v402)>>(uint(v366)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+7)) = uint8(v407)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+4)) = uint8(v413)
	v416 = v400 << (uint(v378) % 32)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416|v402<<(uint(v380)%32))>>(uint(v383)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+6)) = uint8(v426)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416|v401)>>(uint(v392)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+5)) = uint8(v435)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+18)))
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v439)>>(uint(v366)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+11)) = uint8(v444)
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+8)) = uint8(v450)
	v453 = v437 << (uint(v378) % 32)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v453|v439<<(uint(v380)%32))>>(uint(v383)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+10)) = uint8(v463)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v453|v438)>>(uint(v392)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+9)) = uint8(v472)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+19)))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v476)>>(uint(v366)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+15)) = uint8(v481)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+12)) = uint8(v487)
	v490 = v474 << (uint(v378) % 32)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v490|v476<<(uint(v380)%32))>>(uint(v383)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+14)) = uint8(v500)
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v490|v475)>>(uint(v392)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+13)) = uint8(v509)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+21)))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v513)>>(uint(v366)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+19)) = uint8(v518)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+16)) = uint8(v524)
	v527 = v511 << (uint(v378) % 32)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v527|v513<<(uint(v380)%32))>>(uint(v383)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+18)) = uint8(v537)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v527|v512)>>(uint(v392)%32))&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+17)) = uint8(v546)
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	v549 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+22)) = uint8(v549)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v548)>>(uint(v392)%32)))+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+21)) = uint8(v555)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548&v372)+uint32(_consts[1527]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+20)) = uint8(v561)
	v567 = F___memset(m, v13+v380, v549, v380)
	mBase = m.M
	goto L106
L83:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v315 = v301 & int32(1)
	if v315 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v325 = v301 & int32(65535)
	v327 = base.I32_rem_u_s(v325, int32(3))
	if v327 != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v316 = F_strlen(m, l0)
	mBase = m.M
	m.T0[v313].(func(*base.Module, int32, int32, int32))(m, v312, l0, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L28
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v319 = int32(16)
	m.T0[v313].(func(*base.Module, int32, int32, int32))(m, v312, v13+v319, v319)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L28
	} else {
		goto L89
	}
L88:
	;
	goto L84
L89:
	;
	goto L84
L90:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	m.T0[v329].(func(*base.Module, int32, int32, int32))(m, v328, v86, v89)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L28
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v334 = base.I32_rem_u_s(v325, int32(7))
	if v334 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v336 = F_strlen(m, l0)
	mBase = m.M
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	m.T0[v337].(func(*base.Module, int32, int32, int32))(m, v335, l0, v336)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L28
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	if v315 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L96
L98:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+16))
	m.T0[v354].(func(*base.Module, int32, int32))(m, v351, v13+int32(16))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L28
	} else {
		goto L104
	}
L99:
	;
	v343 = int32(16)
	m.T0[v342].(func(*base.Module, int32, int32, int32))(m, v341, v13+v343, v343)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L28
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v348 = F_strlen(m, l0)
	mBase = m.M
	m.T0[v342].(func(*base.Module, int32, int32, int32))(m, v341, l0, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L28
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	goto L98
L104:
	;
	v358 = v301 + int32(1)
	if v358 != int32(1000) {
		v301 = v358
		goto L81
	} else {
		goto L105
	}
L105:
	;
	goto L82
L106:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+20))
	m.T0[v569].(func(*base.Module, int32))(m, v568)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L28
	} else {
		goto L107
	}
L107:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v576 = v572
	v581 = l2
	goto L34
L108:
	;
	v594 = v581
	goto L1
}
func F_px_gen_salt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_CheckBuiltinCryptoMode(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(170981)
	v22 = l0
	goto L6
L3:
	;
	m.G0 = v9 + int32(16)
	return v378
L4:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	if v302 != 0 {
		goto L99
	} else {
		goto L100
	}
L5:
	;
	if v59 == int32(0) {
		v301 = int32(4377792)
		goto L4
	} else {
		goto L18
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == v26 {
		v48 = v25
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v59 = int32(0)
	goto L5
L8:
	;
	v50 = int32(1)
	if v48 != 0 {
		v21 = v21 + v50
		v22 = v22 + v50
		goto L6
	} else {
		goto L17
	}
L9:
	;
	if base.Ui32((v25-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v25 | int32(32)
	goto L12
L11:
	;
	v36 = v25
	goto L12
L12:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = v26 | int32(32)
	goto L15
L14:
	;
	v45 = v26
	goto L15
L15:
	;
	if v36 == v45 {
		v48 = v36
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v59 = v36 - v45
	goto L5
L17:
	;
	goto L7
L18:
	;
	v67 = int32(548705)
	v68 = l0
	goto L20
L19:
	;
	if v105 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v71 == v72 {
		v94 = v71
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v105 = int32(0)
	goto L19
L22:
	;
	v96 = int32(1)
	if v94 != 0 {
		v67 = v67 + v96
		v68 = v68 + v96
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v71-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = v71 | int32(32)
	goto L26
L25:
	;
	v82 = v71
	goto L26
L26:
	;
	if base.Ui32((v72-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = v72 | int32(32)
	goto L29
L28:
	;
	v91 = v72
	goto L29
L29:
	;
	if v82 == v91 {
		v94 = v82
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v105 = v82 - v91
	goto L19
L31:
	;
	goto L21
L32:
	;
	v301 = int32(4377816)
	goto L4
L33:
	;
	goto L34
L34:
	;
	v115 = int32(170786)
	v116 = l0
	goto L36
L35:
	;
	if v153 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v119 == v120 {
		v142 = v119
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v153 = int32(0)
	goto L35
L38:
	;
	v144 = int32(1)
	if v142 != 0 {
		v115 = v115 + v144
		v116 = v116 + v144
		goto L36
	} else {
		goto L47
	}
L39:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = v119 | int32(32)
	goto L42
L41:
	;
	v130 = v119
	goto L42
L42:
	;
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v139 = v120 | int32(32)
	goto L45
L44:
	;
	v139 = v120
	goto L45
L45:
	;
	if v130 == v139 {
		v142 = v130
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v153 = v130 - v139
	goto L35
L47:
	;
	goto L37
L48:
	;
	v301 = int32(4377840)
	goto L4
L49:
	;
	goto L50
L50:
	;
	v163 = int32(338860)
	v164 = l0
	goto L52
L51:
	;
	if v201 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v167 == v168 {
		v190 = v167
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v201 = int32(0)
	goto L51
L54:
	;
	v192 = int32(1)
	if v190 != 0 {
		v163 = v163 + v192
		v164 = v164 + v192
		goto L52
	} else {
		goto L63
	}
L55:
	;
	if base.Ui32((v167-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v178 = v167 | int32(32)
	goto L58
L57:
	;
	v178 = v167
	goto L58
L58:
	;
	if base.Ui32((v168-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v187 = v168 | int32(32)
	goto L61
L60:
	;
	v187 = v168
	goto L61
L61:
	;
	if v178 == v187 {
		v190 = v178
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v201 = v178 - v187
	goto L51
L63:
	;
	goto L53
L64:
	;
	v301 = int32(4377864)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v211 = int32(83187)
	v212 = l0
	goto L68
L67:
	;
	if v249 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v215 == v216 {
		v238 = v215
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v249 = int32(0)
	goto L67
L70:
	;
	v240 = int32(1)
	if v238 != 0 {
		v211 = v211 + v240
		v212 = v212 + v240
		goto L68
	} else {
		goto L79
	}
L71:
	;
	if base.Ui32((v215-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v226 = v215 | int32(32)
	goto L74
L73:
	;
	v226 = v215
	goto L74
L74:
	;
	if base.Ui32((v216-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v235 = v216 | int32(32)
	goto L77
L76:
	;
	v235 = v216
	goto L77
L77:
	;
	if v226 == v235 {
		v238 = v226
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v249 = v226 - v235
	goto L67
L79:
	;
	goto L69
L80:
	;
	v301 = int32(4377888)
	goto L4
L81:
	;
	goto L82
L82:
	;
	v259 = int32(83199)
	v260 = l0
	goto L84
L83:
	;
	if v297 != 0 {
		goto L96
	} else {
		goto L97
	}
L84:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v263 == v264 {
		v286 = v263
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v297 = int32(0)
	goto L83
L86:
	;
	v288 = int32(1)
	if v286 != 0 {
		v259 = v259 + v288
		v260 = v260 + v288
		goto L84
	} else {
		goto L95
	}
L87:
	;
	if base.Ui32((v263-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v274 = v263 | int32(32)
	goto L90
L89:
	;
	v274 = v263
	goto L90
L90:
	;
	if base.Ui32((v264-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v283 = v264 | int32(32)
	goto L93
L92:
	;
	v283 = v264
	goto L93
L93:
	;
	if v274 == v283 {
		v286 = v274
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v297 = v274 - v283
	goto L83
L95:
	;
	goto L85
L96:
	;
	v378 = int32(-14)
	goto L3
L97:
	;
	goto L98
L98:
	;
	v301 = int32(4377912)
	goto L4
L99:
	;
	v303 = int32(-15)
	if l2 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v309 = l2
	goto L101
L101:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v312 = int32(0)
	v316 = m.G0
	v318 = v316 - int32(16)
	m.G0 = v318
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v312
	v324 = F_open(m, int32(287595), v312, v318)
	mBase = m.M
	if v324 != int32(-1) {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	v304 = l2
	goto L104
L103:
	;
	v304 = v302
	goto L104
L104:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+16))
	if v304 < v305 {
		v378 = v303
		goto L3
	} else {
		goto L105
	}
L105:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	if v307 < v304 {
		v378 = v303
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v309 = v304
	goto L101
L107:
	;
	if v357 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L108:
	;
	v327 = int32(1)
	if v311 == int32(0) {
		v350 = v327
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v357 = v312
	goto L110
L110:
	;
	m.G0 = v318 + int32(16)
	goto L107
L111:
	;
	v352 = F_close(m, v324)
	mBase = m.M
	v357 = v350
	goto L110
L112:
	;
	v330 = v9
	v331 = v311
	goto L113
L113:
	;
	v336 = F_read(m, v324, v330, v331)
	mBase = m.M
	if v336 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v350 = v327
	goto L111
L115:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v340 == int32(27) {
		goto L113
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v345 = v331 - v336
	if v345 != 0 {
		v330 = v330 + v336
		v331 = v345
		goto L113
	} else {
		goto L119
	}
L118:
	;
	v350 = int32(0)
	goto L111
L119:
	;
	goto L114
L120:
	;
	v378 = int32(-17)
	goto L3
L121:
	;
	goto L122
L122:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v367 = m.T0[v366].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v309, v9, v311, l1, int32(128))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v371 = F___memset(m, v9, int32(0), int32(16))
	mBase = m.M
	goto L124
L124:
	;
	if v367 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v378 = int32(-15)
	goto L3
L126:
	;
	goto L127
L127:
	;
	v375 = F_strlen(m, v367)
	mBase = m.M
	v378 = v375
	goto L3
}
func F_px_memset(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	v5 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(l1), l2)
	return
}
