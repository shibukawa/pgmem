package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_WakeupRecovery(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _consts[183]))
	F_SetLatch(m, v2+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_wcstombs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	v4 = int32(0)
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v16 = v12 + int32(12)
	v17 = m.G0
	v19 = v17 - v11
	m.G0 = v19
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v539 = int32(16)
	m.G0 = v19 + v539
	m.G0 = v12 + v539
	return v535
L2:
	;
	v535 = l2 - v520
	goto L1
L3:
	;
	if v275 != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v143 = l2
	v144 = l0
	v148 = v142
	goto L37
L5:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 == int32(0) {
		v535 = v4
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v275 = l2
	v276 = l0
	goto L3
L9:
	;
	v27 = v24
	v28 = v23
	v30 = v4
	goto L10
L10:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v27) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v535 = v141
	goto L1
L12:
	;
	v41 = v19 + int32(12)
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v136 = int32(1)
	goto L14
L14:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v141 = v30 + v136
	if v138 != 0 {
		v27 = v138
		v28 = v28 + int32(4)
		v30 = v141
		goto L10
	} else {
		goto L36
	}
L15:
	;
	if v133 == int32(-1) {
		v535 = int32(-1)
		goto L1
	} else {
		goto L35
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v27)
	v133 = int32(1)
	goto L15
L17:
	;
	if base.Ui32(v27) <= base.Ui32(int32(127)) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v130 = int32(1)
	goto L19
L19:
	;
	v133 = v130
	goto L15
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v130 = int32(-1)
	goto L19
L22:
	;
	if v27&int32(-128) == int32(57216) {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v27) <= base.Ui32(int32(2047)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L21
L26:
	;
	v61 = v27&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v61)
	v66 = int32(base.Ui32(v27)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v66)
	v133 = int32(2)
	goto L15
L27:
	;
	goto L28
L28:
	;
	if base.B2i32(v27&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v27)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v78 = int32(63)
	v80 = int32(128)
	v81 = v27&v78 | v80
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)) = uint8(v81)
	v86 = int32(base.Ui32(v27)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v86)
	v93 = int32(base.Ui32(v27)>>(uint(int32(6))%32))&v78 | v80
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v93)
	v133 = int32(3)
	goto L15
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(v27-int32(65536)) <= base.Ui32(int32(1048575)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = int32(63)
	v102 = int32(128)
	v103 = v27&v100 | v102
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)) = uint8(v103)
	v108 = int32(base.Ui32(v27)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v108)
	v115 = int32(base.Ui32(v27)>>(uint(int32(6))%32))&v100 | v102
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)) = uint8(v115)
	v122 = int32(base.Ui32(v27)>>(uint(int32(12))%32))&v100 | v102
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)) = uint8(v122)
	v133 = int32(4)
	goto L15
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L21
L35:
	;
	v136 = v133
	goto L14
L36:
	;
	goto L11
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(v152-int32(128)) <= base.Ui32(int32(-128)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v275 = v265
	v276 = v268
	goto L3
L39:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v271 = v269 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v271
	if base.Ui32(int32(3)) < base.Ui32(v265) {
		v143 = v265
		v144 = v268
		v148 = v271
		goto L37
	} else {
		goto L67
	}
L40:
	;
	if v152 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v152)
	v261 = int32(1)
	v265 = v143 - v261
	v268 = v144 + v261
	goto L39
L43:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v159
	v520 = v143
	goto L2
L44:
	;
	goto L45
L45:
	;
	if v144 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v255 == int32(-1) {
		v535 = int32(-1)
		goto L1
	} else {
		goto L66
	}
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v152)
	v255 = int32(1)
	goto L46
L48:
	;
	if base.Ui32(v152) <= base.Ui32(int32(127)) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v252 = int32(1)
	goto L50
L50:
	;
	v255 = v252
	goto L46
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v168 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v252 = int32(-1)
	goto L50
L53:
	;
	if v152&int32(-128) == int32(57216) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v152) <= base.Ui32(int32(2047)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L52
L57:
	;
	v183 = v152&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)) = uint8(v183)
	v188 = int32(base.Ui32(v152)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v188)
	v255 = int32(2)
	goto L46
L58:
	;
	goto L59
L59:
	;
	if base.B2i32(v152&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v152)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v200 = int32(63)
	v202 = int32(128)
	v203 = v152&v200 | v202
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+2)) = uint8(v203)
	v208 = int32(base.Ui32(v152)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v208)
	v215 = int32(base.Ui32(v152)>>(uint(int32(6))%32))&v200 | v202
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)) = uint8(v215)
	v255 = int32(3)
	goto L46
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v152-int32(65536)) <= base.Ui32(int32(1048575)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v222 = int32(63)
	v224 = int32(128)
	v225 = v152&v222 | v224
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+3)) = uint8(v225)
	v230 = int32(base.Ui32(v152)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v230)
	v237 = int32(base.Ui32(v152)>>(uint(int32(6))%32))&v222 | v224
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+2)) = uint8(v237)
	v244 = int32(base.Ui32(v152)>>(uint(int32(12))%32))&v222 | v224
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)) = uint8(v244)
	v255 = int32(4)
	goto L46
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L52
L66:
	;
	v265 = v143 - v255
	v268 = v144 + v255
	goto L39
L67:
	;
	goto L38
L68:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v285 = v275
	v286 = v276
	v288 = v284
	goto L71
L69:
	;
	goto L70
L70:
	;
	v535 = l2
	goto L1
L71:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if base.Ui32(v294-int32(128)) <= base.Ui32(int32(-128)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L70
L73:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v509 = v507 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v509
	if v503 != 0 {
		v285 = v503
		v286 = v506
		v288 = v509
		goto L71
	} else {
		goto L122
	}
L74:
	;
	if v294 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v294)
	v499 = int32(1)
	v503 = v285 - v499
	v506 = v286 + v499
	goto L73
L77:
	;
	v301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v301)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v301
	v520 = v285
	goto L2
L78:
	;
	goto L79
L79:
	;
	v307 = v19 + int32(12)
	if v307 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	if v399 == int32(-1) {
		v535 = int32(-1)
		goto L1
	} else {
		goto L100
	}
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v294)
	v399 = int32(1)
	goto L80
L82:
	;
	if base.Ui32(v294) <= base.Ui32(int32(127)) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	v396 = int32(1)
	goto L84
L84:
	;
	v399 = v396
	goto L80
L85:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v312 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v396 = int32(-1)
	goto L84
L87:
	;
	if v294&int32(-128) == int32(57216) {
		goto L81
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v294) <= base.Ui32(int32(2047)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L86
L91:
	;
	v327 = v294&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)) = uint8(v327)
	v332 = int32(base.Ui32(v294)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v332)
	v399 = int32(2)
	goto L80
L92:
	;
	goto L93
L93:
	;
	if base.B2i32(v294&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v294)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v344 = int32(63)
	v346 = int32(128)
	v347 = v294&v344 | v346
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+2)) = uint8(v347)
	v352 = int32(base.Ui32(v294)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v352)
	v359 = int32(base.Ui32(v294)>>(uint(int32(6))%32))&v344 | v346
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)) = uint8(v359)
	v399 = int32(3)
	goto L80
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(v294-int32(65536)) <= base.Ui32(int32(1048575)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v366 = int32(63)
	v368 = int32(128)
	v369 = v294&v366 | v368
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+3)) = uint8(v369)
	v374 = int32(base.Ui32(v294)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v374)
	v381 = int32(base.Ui32(v294)>>(uint(int32(6))%32))&v366 | v368
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+2)) = uint8(v381)
	v388 = int32(base.Ui32(v294)>>(uint(int32(12))%32))&v366 | v368
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)) = uint8(v388)
	v399 = int32(4)
	goto L80
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L86
L100:
	;
	if base.Ui32(v285) < base.Ui32(v399) {
		v520 = v285
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v286 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v503 = v285 - v399
	v506 = v286 + v399
	goto L73
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v403)
	goto L102
L104:
	;
	if base.Ui32(v403) <= base.Ui32(int32(127)) {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L102
L107:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	if v408 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L106
L109:
	;
	if v403&int32(-128) == int32(57216) {
		goto L103
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if base.Ui32(v403) <= base.Ui32(int32(2047)) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L108
L113:
	;
	v423 = v403&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)) = uint8(v423)
	v428 = int32(base.Ui32(v403)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v428)
	goto L102
L114:
	;
	goto L115
L115:
	;
	if base.B2i32(v403&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v403)) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v440 = int32(63)
	v442 = int32(128)
	v443 = v403&v440 | v442
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+2)) = uint8(v443)
	v448 = int32(base.Ui32(v403)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v448)
	v455 = int32(base.Ui32(v403)>>(uint(int32(6))%32))&v440 | v442
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)) = uint8(v455)
	goto L102
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v403-int32(65536)) <= base.Ui32(int32(1048575)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v462 = int32(63)
	v464 = int32(128)
	v465 = v403&v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+3)) = uint8(v465)
	v470 = int32(base.Ui32(v403)>>(uint(int32(18))%32)) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v470)
	v477 = int32(base.Ui32(v403)>>(uint(int32(6))%32))&v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+2)) = uint8(v477)
	v484 = int32(base.Ui32(v403)>>(uint(int32(12))%32))&v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)) = uint8(v484)
	goto L102
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(25)
	goto L108
L122:
	;
	goto L72
}
func F_weak_input_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1123])))
	return v3
}
func F_websearch_to_tsquery(m *base.Module, l0 int32) int32 {
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
			v11 = F_DirectFunctionCall2Coll(m, int32(1178), int32(0), v9, v3)
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
func F_width_bucket_float8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 float64
	_ = v65
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(0) < v8 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v175 = m.ExcPending
			if v175 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(386138242))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(552407), int32(0))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515956), int32(4090), int32(580385))
						mBase = m.M
						v187 = m.ExcPending
						if v187 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(386138242))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(552407), int32(0))
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515956), int32(4090), int32(580385))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v26 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v26)&int64(9223372036854775807)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(386138242))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(552407), int32(0))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515956), int32(4090), int32(580385))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.F64_eq(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(386138242))
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(367215), int32(0))
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515956), int32(4096), int32(580385))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.F64_eq(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(386138242))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(367215), int32(0))
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515956), int32(4096), int32(580385))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if base.F64_lt(v19, v26) != 0 {
								if base.F64_gt(v19, v12) != 0 {
									return int32(0)
								} else {
									if base.F64_le(v26, v12) != 0 {
										v44 = v8 + int32(1)
										if v8 <= v44 {
											v149 = v44
											return v149
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(420724), int32(0))
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515956), int32(4107), int32(580385))
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
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
										v65 = base.F64_sub(v26, v19)
										if base.F64_ne(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v81 = base.F64_div(base.F64_sub(v12, v19), v65)
										} else {
											v71 = float64(0.5)
											v74 = base.F64_mul(v19, v71)
											v81 = base.F64_div(base.F64_sub(base.F64_mul(v12, v71), v74), base.F64_sub(base.F64_mul(v26, v71), v74))
										}
										v82 = base.F64_mul(v81, base.F64_convert_i32_u(v8))
										if base.F64_lt(base.F64_abs(v82), float64(2.147483648e+09)) != 0 {
											v86 = base.I32_trunc_f64_s(v82)
											v88 = v86
										} else {
											v88 = int32(-2147483648)
										}
										if v88 < v8 {
											v92 = v88 + int32(1)
										} else {
											v92 = v8
										}
										return v92
									}
								}
							} else {
								if base.F64_gt(v19, v26) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(444258), int32(0))
											mBase = m.M
											v214 = m.ExcPending
											if v214 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(515956), int32(4162), int32(580385))
												mBase = m.M
												v219 = m.ExcPending
												if v219 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									if base.F64_lt(v19, v12) != 0 {
										return int32(0)
									} else {
										if base.F64_ge(v26, v12) != 0 {
											v102 = v8 + int32(1)
											if v8 <= v102 {
												v149 = v102
												return v149
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(420724), int32(0))
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(515956), int32(4145), int32(580385))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
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
											v121 = base.F64_sub(v19, v26)
											if base.F64_ne(base.F64_abs(v121), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v137 = base.F64_div(base.F64_sub(v19, v12), v121)
											} else {
												v127 = float64(0.5)
												v128 = base.F64_mul(v19, v127)
												v137 = base.F64_div(base.F64_sub(v128, base.F64_mul(v12, v127)), base.F64_sub(v128, base.F64_mul(v26, v127)))
											}
											v138 = base.F64_mul(v137, base.F64_convert_i32_u(v8))
											if base.F64_lt(base.F64_abs(v138), float64(2.147483648e+09)) != 0 {
												v142 = base.I32_trunc_f64_s(v138)
												v144 = v142
											} else {
												v144 = int32(-2147483648)
											}
											if v144 < v8 {
												v148 = v144 + int32(1)
											} else {
												v148 = v8
											}
											v149 = v148
											return v149
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v157 = m.ExcPending
		if v157 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(386138242))
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(251475), int32(0))
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515956), int32(4085), int32(580385))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
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
func F_win866_to_koi8r(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(20), int32(22))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(20), int32(22), int32(2266720), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win866_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(20), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_latin2mic_with_table(m, v6, v5, v10, int32(139), int32(20), int32(2266720), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_write_stderr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	v14 = m.G0
	v16 = v14 - int32(1056)
	m.G0 = v16
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(28)
		m.G0 = v16 + int32(1056)
		v65 = F_fflush(m, v13)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	} else {
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+1052)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1048)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1044)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1040)) = v16 + int32(1024)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1036)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v16)+1032)) = v16
		F_dopr(m, v16+int32(1032), l0, l1)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1052)))
			if v37 == int32(0) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1032))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1036))
				if v40 != v41 {
					v47 = v40 - v41
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1044))
					v49 = F_fwrite(m, v41, int32(1), v47, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						if v49 != v47 {
						} else {
							if v37 != 0 {
							} else {
							}
						}
						m.G0 = v16 + int32(1056)
						v65 = F_fflush(m, v13)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					if v37 != 0 {
					} else {
					}
					m.G0 = v16 + int32(1056)
					v65 = F_fflush(m, v13)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			} else {
				if v37 != 0 {
				} else {
				}
				m.G0 = v16 + int32(1056)
				v65 = F_fflush(m, v13)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	}
}
