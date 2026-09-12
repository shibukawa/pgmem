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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
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
	v12 = int32(689799)
	goto L6
L3:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	if v328 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L4:
	;
	if v50-v51 == int32(0) {
		v327 = int32(4395264)
		goto L3
	} else {
		goto L18
	}
L6:
	;
	goto L7
L7:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v21 = l1
	v22 = v12
	v23 = int32(4)
	v24 = v20
	goto L12
L9:
	;
	v46 = v12
	v50 = int32(0)
	goto L10
L10:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	goto L4
L11:
	;
	v46 = v41
	v50 = v43
	goto L10
L12:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v24 != v26 {
		v41 = v22
		v43 = v24
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v41 = v35
	v43 = int32(0)
	goto L11
L14:
	;
	if v26 == int32(0) {
		v41 = v22
		v43 = v24
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v31 = v23 - int32(1)
	if v31 == int32(0) {
		v41 = v22
		v43 = v24
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v34 = int32(1)
	v35 = v22 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v36 != 0 {
		v21 = v21 + v34
		v22 = v35
		v23 = v31
		v24 = v36
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v62 = int32(689767)
	goto L21
L19:
	;
	if v100-v101 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	goto L22
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v70 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v71 = l1
	v72 = v62
	v73 = int32(4)
	v74 = v70
	goto L27
L24:
	;
	v96 = v62
	v100 = int32(0)
	goto L25
L25:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	goto L19
L26:
	;
	v96 = v91
	v100 = v93
	goto L25
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 != v76 {
		v91 = v72
		v93 = v74
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v91 = v85
	v93 = int32(0)
	goto L26
L29:
	;
	if v76 == int32(0) {
		v91 = v72
		v93 = v74
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v81 = v73 - int32(1)
	if v81 == int32(0) {
		v91 = v72
		v93 = v74
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v84 = int32(1)
	v85 = v72 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v86 != 0 {
		v71 = v71 + v84
		v72 = v85
		v73 = v81
		v74 = v86
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v327 = int32(4395276)
	goto L3
L34:
	;
	goto L35
L35:
	;
	v114 = int32(689812)
	goto L38
L36:
	;
	if v152-v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	goto L39
L39:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v122 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v123 = l1
	v124 = v114
	v125 = int32(3)
	v126 = v122
	goto L44
L41:
	;
	v148 = v114
	v152 = int32(0)
	goto L42
L42:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	goto L36
L43:
	;
	v148 = v143
	v152 = v145
	goto L42
L44:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v126 != v128 {
		v143 = v124
		v145 = v126
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v143 = v137
	v145 = int32(0)
	goto L43
L46:
	;
	if v128 == int32(0) {
		v143 = v124
		v145 = v126
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v133 = v125 - int32(1)
	if v133 == int32(0) {
		v143 = v124
		v145 = v126
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v136 = int32(1)
	v137 = v124 + v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if v138 != 0 {
		v123 = v123 + v136
		v124 = v137
		v125 = v133
		v126 = v138
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v327 = int32(4395288)
	goto L3
L51:
	;
	goto L52
L52:
	;
	v166 = int32(689816)
	goto L55
L53:
	;
	if v204-v205 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L55:
	;
	goto L56
L56:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v174 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v175 = l1
	v176 = v166
	v177 = int32(3)
	v178 = v174
	goto L61
L58:
	;
	v200 = v166
	v204 = int32(0)
	goto L59
L59:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	goto L53
L60:
	;
	v200 = v195
	v204 = v197
	goto L59
L61:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v178 != v180 {
		v195 = v176
		v197 = v178
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v195 = v189
	v197 = int32(0)
	goto L60
L63:
	;
	if v180 == int32(0) {
		v195 = v176
		v197 = v178
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v185 = v177 - int32(1)
	if v185 == int32(0) {
		v195 = v176
		v197 = v178
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v188 = int32(1)
	v189 = v176 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v190 != 0 {
		v175 = v175 + v188
		v176 = v189
		v177 = v185
		v178 = v190
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v327 = int32(4395300)
	goto L3
L68:
	;
	goto L69
L69:
	;
	v218 = int32(689808)
	goto L72
L70:
	;
	if v256-v257 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L72:
	;
	goto L73
L73:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v226 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v227 = l1
	v228 = v218
	v229 = int32(3)
	v230 = v226
	goto L78
L75:
	;
	v252 = v218
	v256 = int32(0)
	goto L76
L76:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	goto L70
L77:
	;
	v252 = v247
	v256 = v249
	goto L76
L78:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v230 != v232 {
		v247 = v228
		v249 = v230
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v247 = v241
	v249 = int32(0)
	goto L77
L80:
	;
	if v232 == int32(0) {
		v247 = v228
		v249 = v230
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v237 = v229 - int32(1)
	if v237 == int32(0) {
		v247 = v228
		v249 = v230
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v240 = int32(1)
	v241 = v228 + v240
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	if v242 != 0 {
		v227 = v227 + v240
		v228 = v241
		v229 = v237
		v230 = v242
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	v327 = int32(4395312)
	goto L3
L85:
	;
	goto L86
L86:
	;
	v270 = int32(689804)
	goto L89
L87:
	;
	if v308-v309 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L89:
	;
	goto L90
L90:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v278 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v279 = l1
	v280 = v270
	v281 = int32(3)
	v282 = v278
	goto L95
L92:
	;
	v304 = v270
	v308 = int32(0)
	goto L93
L93:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	goto L87
L94:
	;
	v304 = v299
	v308 = v301
	goto L93
L95:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v282 != v284 {
		v299 = v280
		v301 = v282
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v299 = v293
	v301 = int32(0)
	goto L94
L97:
	;
	if v284 == int32(0) {
		v299 = v280
		v301 = v282
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v289 = v281 - int32(1)
	if v289 == int32(0) {
		v299 = v280
		v301 = v282
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v292 = int32(1)
	v293 = v280 + v292
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v294 != 0 {
		v279 = v279 + v292
		v280 = v293
		v281 = v289
		v282 = v294
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	v327 = int32(4395324)
	goto L3
L102:
	;
	goto L103
L103:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v323 == int32(95) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v326 = int32(4395336)
	goto L106
L105:
	;
	v326 = int32(4395348)
	goto L106
L106:
	;
	v327 = v326
	goto L3
L107:
	;
	return int32(0)
L108:
	;
	goto L109
L109:
	;
	v333 = m.T0[v328].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	return v333
}
