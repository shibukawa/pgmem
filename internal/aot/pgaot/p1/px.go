package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_crypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	F_CheckBuiltinCryptoMode(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(_a_F_px_crypt_0)
	goto L6
L3:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	if v317 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L4:
	;
	if v49-v50 == int32(0) {
		v316 = int32(_a_F_px_crypt_1)
		goto L3
	} else {
		goto L17
	}
L6:
	;
	goto L7
L7:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v19 = l1
	v20 = v11
	v21 = int32(4)
	v22 = v18
	goto L12
L9:
	;
	v45 = v11
	v49 = int32(0)
	goto L10
L10:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L4
L11:
	;
	v45 = v40
	v49 = v42
	goto L10
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v40 = v34
	v42 = int32(0)
	goto L11
L14:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v61 = int32(_a_F_px_crypt_2)
	goto L20
L18:
	;
	if v99-v100 == int32(0) {
		v316 = int32(_a_F_px_crypt_3)
		goto L3
	} else {
		goto L31
	}
L20:
	;
	goto L21
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = l1
	v70 = v61
	v71 = int32(4)
	v72 = v68
	goto L26
L23:
	;
	v95 = v61
	v99 = int32(0)
	goto L24
L24:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	goto L18
L25:
	;
	v95 = v90
	v99 = v92
	goto L24
L26:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.B2i32(v72 != v74)|base.B2i32(v74 == int32(0)) != 0 {
		v90 = v70
		v92 = v72
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v90 = v84
	v92 = int32(0)
	goto L25
L28:
	;
	v80 = v71 - int32(1)
	if v80 == int32(0) {
		v90 = v70
		v92 = v72
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v83 = int32(1)
	v84 = v70 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v85 != 0 {
		v69 = v69 + v83
		v70 = v84
		v71 = v80
		v72 = v85
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v111 = int32(_a_F_px_crypt_4)
	goto L34
L32:
	;
	if v149-v150 == int32(0) {
		v316 = int32(_a_F_px_crypt_5)
		goto L3
	} else {
		goto L45
	}
L34:
	;
	goto L35
L35:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v118 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = l1
	v120 = v111
	v121 = int32(3)
	v122 = v118
	goto L40
L37:
	;
	v145 = v111
	v149 = int32(0)
	goto L38
L38:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	goto L32
L39:
	;
	v145 = v140
	v149 = v142
	goto L38
L40:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if base.B2i32(v122 != v124)|base.B2i32(v124 == int32(0)) != 0 {
		v140 = v120
		v142 = v122
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v140 = v134
	v142 = int32(0)
	goto L39
L42:
	;
	v130 = v121 - int32(1)
	if v130 == int32(0) {
		v140 = v120
		v142 = v122
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v133 = int32(1)
	v134 = v120 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v135 != 0 {
		v119 = v119 + v133
		v120 = v134
		v121 = v130
		v122 = v135
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v161 = int32(_a_F_px_crypt_6)
	goto L48
L46:
	;
	if v199-v200 == int32(0) {
		v316 = int32(_a_F_px_crypt_7)
		goto L3
	} else {
		goto L59
	}
L48:
	;
	goto L49
L49:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v168 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v169 = l1
	v170 = v161
	v171 = int32(3)
	v172 = v168
	goto L54
L51:
	;
	v195 = v161
	v199 = int32(0)
	goto L52
L52:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	goto L46
L53:
	;
	v195 = v190
	v199 = v192
	goto L52
L54:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if base.B2i32(v172 != v174)|base.B2i32(v174 == int32(0)) != 0 {
		v190 = v170
		v192 = v172
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v190 = v184
	v192 = int32(0)
	goto L53
L56:
	;
	v180 = v171 - int32(1)
	if v180 == int32(0) {
		v190 = v170
		v192 = v172
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v183 = int32(1)
	v184 = v170 + v183
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v185 != 0 {
		v169 = v169 + v183
		v170 = v184
		v171 = v180
		v172 = v185
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v211 = int32(_a_F_px_crypt_8)
	goto L62
L60:
	;
	if v249-v250 == int32(0) {
		v316 = int32(_a_F_px_crypt_9)
		goto L3
	} else {
		goto L73
	}
L62:
	;
	goto L63
L63:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v218 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v219 = l1
	v220 = v211
	v221 = int32(3)
	v222 = v218
	goto L68
L65:
	;
	v245 = v211
	v249 = int32(0)
	goto L66
L66:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	goto L60
L67:
	;
	v245 = v240
	v249 = v242
	goto L66
L68:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if base.B2i32(v222 != v224)|base.B2i32(v224 == int32(0)) != 0 {
		v240 = v220
		v242 = v222
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v240 = v234
	v242 = int32(0)
	goto L67
L70:
	;
	v230 = v221 - int32(1)
	if v230 == int32(0) {
		v240 = v220
		v242 = v222
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v233 = int32(1)
	v234 = v220 + v233
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v235 != 0 {
		v219 = v219 + v233
		v220 = v234
		v221 = v230
		v222 = v235
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v261 = int32(_a_F_px_crypt_10)
	goto L76
L74:
	;
	if v299-v300 == int32(0) {
		v316 = int32(_a_F_px_crypt_11)
		goto L3
	} else {
		goto L87
	}
L76:
	;
	goto L77
L77:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v268 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v269 = l1
	v270 = v261
	v271 = int32(3)
	v272 = v268
	goto L82
L79:
	;
	v295 = v261
	v299 = int32(0)
	goto L80
L80:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	goto L74
L81:
	;
	v295 = v290
	v299 = v292
	goto L80
L82:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if base.B2i32(v272 != v274)|base.B2i32(v274 == int32(0)) != 0 {
		v290 = v270
		v292 = v272
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v290 = v284
	v292 = int32(0)
	goto L81
L84:
	;
	v280 = v271 - int32(1)
	if v280 == int32(0) {
		v290 = v270
		v292 = v272
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v283 = int32(1)
	v284 = v270 + v283
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v285 != 0 {
		v269 = v269 + v283
		v270 = v284
		v271 = v280
		v272 = v285
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v312 == int32(95) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v315 = int32(_a_F_px_crypt_12)
	goto L90
L89:
	;
	v315 = int32(_a_F_px_crypt_13)
	goto L90
L90:
	;
	v316 = v315
	goto L3
L91:
	;
	return int32(0)
L92:
	;
	goto L93
L93:
	;
	v322 = m.T0[v317].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	return v322
}
