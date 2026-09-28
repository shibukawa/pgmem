package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StatsShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_StatsShmemRequest_0)
	v12 = F_add_size(m, int32(_a_F_StatsShmemRequest_1), int32(_a_F_StatsShmemRequest_2))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L6
L3:
	;
	goto L17
L4:
	;
	if v41 == int32(0) {
		v58 = v12
		goto L3
	} else {
		goto L11
	}
L6:
	;
	goto L7
L7:
	;
	goto L9
L8:
	;
	goto L4
L9:
	;
	v29 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v31 == v29 {
		v41 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = int32(96)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31+v35-v35)))
	v41 = v39
	goto L8
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v46&int32(1) == int32(0) {
		v58 = v12
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v56 = F_add_size(m, v12, (v51+int32(7))&int32(-8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v58 = v56
	goto L3
L14:
	;
	goto L28
L15:
	;
	if v86 == int32(0) {
		v103 = v58
		goto L14
	} else {
		goto L22
	}
L17:
	;
	goto L18
L18:
	;
	goto L20
L19:
	;
	goto L15
L20:
	;
	v74 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v76 == v74 {
		v86 = v74
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(100)-int32(96))))
	v86 = v84
	goto L19
L22:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v91&int32(1) == int32(0) {
		v103 = v58
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v101 = F_add_size(m, v58, (v96+int32(7))&int32(-8))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v103 = v101
	goto L14
L25:
	;
	goto L39
L26:
	;
	if v131 == int32(0) {
		v148 = v103
		goto L25
	} else {
		goto L33
	}
L28:
	;
	goto L29
L29:
	;
	goto L31
L30:
	;
	goto L26
L31:
	;
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v121 == v119 {
		v131 = v119
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(104)-int32(96))))
	v131 = v129
	goto L30
L33:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v136&int32(1) == int32(0) {
		v148 = v103
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v146 = F_add_size(m, v103, (v141+int32(7))&int32(-8))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v148 = v146
	goto L25
L36:
	;
	goto L50
L37:
	;
	if v176 == int32(0) {
		v193 = v148
		goto L36
	} else {
		goto L44
	}
L39:
	;
	goto L40
L40:
	;
	goto L42
L41:
	;
	goto L37
L42:
	;
	v164 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v166 == v164 {
		v176 = v164
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166+int32(108)-int32(96))))
	v176 = v174
	goto L41
L44:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v181&int32(1) == int32(0) {
		v193 = v148
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v191 = F_add_size(m, v148, (v186+int32(7))&int32(-8))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v193 = v191
	goto L36
L47:
	;
	goto L61
L48:
	;
	if v221 == int32(0) {
		v238 = v193
		goto L47
	} else {
		goto L55
	}
L50:
	;
	goto L51
L51:
	;
	goto L53
L52:
	;
	goto L48
L53:
	;
	v209 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v211 == v209 {
		v221 = v209
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(112)-int32(96))))
	v221 = v219
	goto L52
L55:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v226&int32(1) == int32(0) {
		v238 = v193
		goto L47
	} else {
		goto L56
	}
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v236 = F_add_size(m, v193, (v231+int32(7))&int32(-8))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v238 = v236
	goto L47
L58:
	;
	goto L72
L59:
	;
	if v266 == int32(0) {
		v283 = v238
		goto L58
	} else {
		goto L66
	}
L61:
	;
	goto L62
L62:
	;
	goto L64
L63:
	;
	goto L59
L64:
	;
	v254 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v256 == v254 {
		v266 = v254
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256+int32(116)-int32(96))))
	v266 = v264
	goto L63
L66:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v271&int32(1) == int32(0) {
		v283 = v238
		goto L58
	} else {
		goto L67
	}
L67:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v281 = F_add_size(m, v238, (v276+int32(7))&int32(-8))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v283 = v281
	goto L58
L69:
	;
	goto L83
L70:
	;
	if v311 == int32(0) {
		v328 = v283
		goto L69
	} else {
		goto L77
	}
L72:
	;
	goto L73
L73:
	;
	goto L75
L74:
	;
	goto L70
L75:
	;
	v299 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v301 == v299 {
		v311 = v299
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v301+int32(120)-int32(96))))
	v311 = v309
	goto L74
L77:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v316&int32(1) == int32(0) {
		v328 = v283
		goto L69
	} else {
		goto L78
	}
L78:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v326 = F_add_size(m, v283, (v321+int32(7))&int32(-8))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v328 = v326
	goto L69
L80:
	;
	goto L94
L81:
	;
	if v356 == int32(0) {
		v373 = v328
		goto L80
	} else {
		goto L88
	}
L83:
	;
	goto L84
L84:
	;
	goto L86
L85:
	;
	goto L81
L86:
	;
	v344 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v346 == v344 {
		v356 = v344
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v346+int32(124)-int32(96))))
	v356 = v354
	goto L85
L88:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v361&int32(1) == int32(0) {
		v373 = v328
		goto L80
	} else {
		goto L89
	}
L89:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v371 = F_add_size(m, v328, (v366+int32(7))&int32(-8))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v373 = v371
	goto L80
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_StatsShmemRequest_3)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v418
	F_ShmemRequestStructWithOpts(m, v6)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L102
	}
L92:
	;
	if v401 == int32(0) {
		v418 = v373
		goto L91
	} else {
		goto L99
	}
L94:
	;
	goto L95
L95:
	;
	goto L97
L96:
	;
	goto L92
L97:
	;
	v389 = int32(0)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemRequest[0]))
	if v391 == v389 {
		v401 = v389
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v391+int32(128)-int32(96))))
	v401 = v399
	goto L96
L99:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v406&int32(1) == int32(0) {
		v418 = v373
		goto L91
	} else {
		goto L100
	}
L100:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v416 = F_add_size(m, v373, (v411+int32(7))&int32(-8))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v418 = v416
	goto L91
L102:
	;
	m.G0 = v6 + int32(16)
	return
}
