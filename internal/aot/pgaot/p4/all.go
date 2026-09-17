package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetAllOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
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
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	v1 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0]))
	if base.B2i32(v10 == v1)|base.B2i32(v10 == int32(_a_F_ResetAllOptions_0)) == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = v10
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(60))))
	if base.Ui32(int32(1)) < base.Ui32(v29-int32(5)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if v26 != int32(_a_F_ResetAllOptions_0) {
		v23 = v26
		goto L4
	} else {
		goto L135
	}
L7:
	;
	v35 = v23 - int32(44)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36&int32(16) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v40 = v23 - int32(32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui32(v41) < base.Ui32(int32(11)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_push_old_value(m, v23+int32(-64), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v50 = v23 - int32(40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v51 {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		goto L12
	}
L12:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(28))))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v320 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L13:
	;
	F_pfree(m, v301)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L10
	} else {
		goto L122
	}
L14:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v258 != 0 {
		goto L103
	} else {
		goto L104
	}
L15:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v180 != 0 {
		goto L74
	} else {
		goto L75
	}
L16:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	if v137 != 0 {
		goto L55
	} else {
		goto L56
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	if v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	m.T0[v52].(func(*base.Module, int32, int32))(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
	v61 = v23 - int32(4)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v63
	if base.B2i32(v62 == int32(0))|base.B2i32(v63 == v62) != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v69 - int32(1) {
	case 0:
		goto L27
	case 1:
		goto L26
	default:
		goto L24
	case 3:
		goto L25
	}
L24:
	;
	v81 = v23 - int32(8)
	goto L31
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if v62 == v76 {
		goto L12
	} else {
		goto L30
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v62 != v74 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	if v62 != v72 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L12
L29:
	;
	goto L12
L30:
	;
	goto L24
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v88 == int32(0) {
		v301 = v62
		goto L13
	} else {
		goto L33
	}
L32:
	;
	goto L12
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+40))
	if v62 == v91 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)+56))
	if v62 != v93 {
		v81 = v88
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	m.T0[v95].(func(*base.Module, int32, int32))(m, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v101
	v104 = v23 - int32(4)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v106
	if base.B2i32(v105 == int32(0))|base.B2i32(v106 == v105) != 0 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v112 {
	case 0:
		goto L45
	default:
		goto L41
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	}
L41:
	;
	v123 = v23 - int32(8)
	goto L50
L42:
	;
	if v105 == v101 {
		goto L12
	} else {
		goto L49
	}
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	if v105 != v117 {
		goto L41
	} else {
		goto L48
	}
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v105 != v115 {
		goto L41
	} else {
		goto L47
	}
L45:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	if v105 != v113 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L12
L47:
	;
	goto L12
L48:
	;
	goto L12
L49:
	;
	goto L41
L50:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v130 == int32(0) {
		v301 = v105
		goto L13
	} else {
		goto L52
	}
L51:
	;
	goto L12
L52:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	if v105 == v133 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+56))
	if v105 != v135 {
		v123 = v130
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v23)+72))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	m.T0[v137].(func(*base.Module, float64, int32))(m, v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v23)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v142))) = v143
	v146 = v23 - int32(4)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v148
	if base.B2i32(v147 == int32(0))|base.B2i32(v148 == v147) != 0 {
		goto L12
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v154 {
	case 0:
		goto L64
	case 1:
		goto L63
	default:
		goto L60
	case 3:
		goto L62
	case 4:
		goto L61
	}
L60:
	;
	v166 = v23 - int32(8)
	goto L69
L61:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if v147 == v161 {
		goto L12
	} else {
		goto L68
	}
L62:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	if v147 != v159 {
		goto L60
	} else {
		goto L67
	}
L63:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	if v147 != v157 {
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	if v147 != v155 {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L12
L66:
	;
	goto L12
L67:
	;
	goto L12
L68:
	;
	goto L60
L69:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v173 == int32(0) {
		v301 = v147
		goto L13
	} else {
		goto L71
	}
L70:
	;
	goto L12
L71:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v173)+40))
	if v147 == v176 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
	if v147 != v178 {
		v166 = v173
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	m.T0[v180].(func(*base.Module, int32, int32))(m, v181, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v187
	if v186 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v224 = v23 - int32(4)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v226
	if base.B2i32(v225 == int32(0))|base.B2i32(v226 == v225) != 0 {
		goto L12
	} else {
		goto L90
	}
L79:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if base.B2i32(v186 == v192)|base.B2i32(v187 == v186) != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v186 == v196 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v202 = v23 - int32(8)
	goto L82
L82:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v208 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_pfree(m, v186)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L89
	}
L84:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+32))
	if v186 == v209 {
		goto L78
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v208)+48))
	if v186 != v211 {
		v202 = v208
		goto L82
	} else {
		goto L88
	}
L88:
	;
	goto L78
L89:
	;
	goto L78
L90:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v232 - int32(1) {
	case 0:
		goto L94
	case 1:
		goto L93
	default:
		goto L91
	case 3:
		goto L92
	}
L91:
	;
	v244 = v23 - int32(8)
	goto L98
L92:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if v225 == v239 {
		goto L12
	} else {
		goto L97
	}
L93:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v225 != v237 {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	if v225 != v235 {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L12
L96:
	;
	goto L12
L97:
	;
	goto L91
L98:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	if v251 == int32(0) {
		v301 = v225
		goto L13
	} else {
		goto L100
	}
L99:
	;
	goto L12
L100:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+40))
	if v225 == v254 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)+56))
	if v225 != v256 {
		v244 = v251
		goto L98
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	m.T0[v258].(func(*base.Module, int32, int32))(m, v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v264
	v267 = v23 - int32(4)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v269
	if base.B2i32(v268 == int32(0))|base.B2i32(v269 == v268) != 0 {
		goto L12
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v275 {
	case 0:
		goto L112
	case 1:
		goto L111
	case 2:
		goto L110
	case 3:
		goto L109
	default:
		goto L108
	}
L108:
	;
	v285 = v23 - int32(8)
	goto L117
L109:
	;
	if v268 == v264 {
		goto L12
	} else {
		goto L116
	}
L110:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v268 != v279 {
		goto L108
	} else {
		goto L115
	}
L111:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	if v268 != v277 {
		goto L108
	} else {
		goto L114
	}
L112:
	;
	if v268 != v264 {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L12
L114:
	;
	goto L12
L115:
	;
	goto L12
L116:
	;
	goto L108
L117:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	if v292 == int32(0) {
		v301 = v268
		goto L13
	} else {
		goto L119
	}
L118:
	;
	goto L12
L119:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)+40))
	if v268 == v295 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)+56))
	if v268 != v297 {
		v285 = v292
		goto L117
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	goto L12
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v319
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(24)))) = v351
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(16)))) = v357
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v359&int32(64) == int32(0) {
		goto L6
	} else {
		goto L133
	}
L124:
	;
	if v319 == int32(0) {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	if v319 != 0 {
		goto L123
	} else {
		goto L132
	}
L127:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0]))
	if v326 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v333
	v335 = int32(_a_F_ResetAllOptions_0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v333)+4)) = v23
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[1])) = v23
	goto L123
L129:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[1]))
	v333 = v328
	goto L128
L130:
	;
	goto L131
L131:
	;
	v330 = int32(_a_F_ResetAllOptions_0)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0])) = v330
	v333 = v330
	goto L128
L132:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+4)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v343
	goto L123
L133:
	;
	v365 = v23 - int32(36)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v366&int32(4) != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v366 | int32(4)
	v372 = int32(_a_F_ResetAllOptions_1)
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v373
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[2])) = v23 + int32(12)
	goto L6
L135:
	;
	goto L5
}
