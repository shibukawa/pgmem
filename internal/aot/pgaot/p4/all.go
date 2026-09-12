package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetAllOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	v12 = *(*int32)(unsafe.Add(mBase, _consts[432]))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 == int32(4547052) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v12
	goto L4
L4:
	;
	v28 = v24 + int32(4)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(60))))
	if base.Ui32(int32(1)) < base.Ui32(v32-int32(5)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	if v29 != int32(4547052) {
		v24 = v29
		goto L4
	} else {
		goto L141
	}
L7:
	;
	v38 = v24 - int32(44)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39&int32(16) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = v24 - int32(32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if base.Ui32(v44) < base.Ui32(int32(11)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_push_old_value(m, v24+int32(-64), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v53 = v24 - int32(40)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v54 {
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
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(28))))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v339 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L13:
	;
	F_pfree(m, v315)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L128
	}
L14:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v272 != 0 {
		goto L108
	} else {
		goto L109
	}
L15:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v188 != 0 {
		goto L77
	} else {
		goto L78
	}
L16:
	;
	v143 = v24 + int32(60)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v144 != 0 {
		goto L57
	} else {
		goto L58
	}
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v99 != 0 {
		goto L37
	} else {
		goto L38
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+48)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	m.T0[v55].(func(*base.Module, int32, int32))(m, v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v61)
	v64 = v24 - int32(4)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v66
	if v65 == int32(0) {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if v65 == v66 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v71 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L25
	case 3:
		goto L26
	}
L25:
	;
	v84 = v24 - int32(8)
	goto L32
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	if v65 == v78 {
		goto L12
	} else {
		goto L31
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	if v65 != v76 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	if v65 != v74 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L12
L30:
	;
	goto L12
L31:
	;
	goto L25
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v92 == int32(0) {
		v315 = v65
		goto L13
	} else {
		goto L34
	}
L33:
	;
	goto L12
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+40))
	if v65 == v95 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+56))
	if v65 != v97 {
		v84 = v92
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	m.T0[v99].(func(*base.Module, int32, int32))(m, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v105
	v108 = v24 - int32(4)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
	if v109 == int32(0) {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	if v109 == v110 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v115 {
	case 0:
		goto L47
	default:
		goto L43
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	}
L43:
	;
	v127 = v24 - int32(8)
	goto L52
L44:
	;
	if v109 == v105 {
		goto L12
	} else {
		goto L51
	}
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	if v109 != v120 {
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	if v109 != v118 {
		goto L43
	} else {
		goto L49
	}
L47:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	if v109 != v116 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L12
L49:
	;
	goto L12
L50:
	;
	goto L12
L51:
	;
	goto L43
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v135 == int32(0) {
		v315 = v109
		goto L13
	} else {
		goto L54
	}
L53:
	;
	goto L12
L54:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	if v109 == v138 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+56))
	if v109 != v140 {
		v127 = v135
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v24)+72))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	m.T0[v144].(func(*base.Module, float64, int32))(m, v145, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L10
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v24)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v149))) = v150
	v153 = v24 - int32(4)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v155
	if v154 == int32(0) {
		goto L12
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	if v154 == v155 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v160 {
	case 0:
		goto L67
	case 1:
		goto L66
	default:
		goto L63
	case 3:
		goto L65
	case 4:
		goto L64
	}
L63:
	;
	v173 = v24 - int32(8)
	goto L72
L64:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	if v154 == v167 {
		goto L12
	} else {
		goto L71
	}
L65:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	if v154 != v165 {
		goto L63
	} else {
		goto L70
	}
L66:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v154 != v163 {
		goto L63
	} else {
		goto L69
	}
L67:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	if v154 != v161 {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	goto L12
L69:
	;
	goto L12
L70:
	;
	goto L12
L71:
	;
	goto L63
L72:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v181 == int32(0) {
		v315 = v154
		goto L13
	} else {
		goto L74
	}
L73:
	;
	goto L12
L74:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+40))
	if v154 == v184 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+56))
	if v154 != v186 {
		v173 = v181
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	m.T0[v188].(func(*base.Module, int32, int32))(m, v189, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v194 = v24 + int32(28)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v197
	if v196 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v237 = v24 - int32(4)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v239
	if v238 == int32(0) {
		goto L12
	} else {
		goto L94
	}
L82:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v196 == v202 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if v196 == v197 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	if v196 == v205 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v210 = v24 - int32(8)
	goto L86
L86:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v219 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_pfree(m, v196)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L93
	}
L88:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+32))
	if v196 == v220 {
		goto L81
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	goto L87
L91:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	if v196 != v222 {
		v210 = v219
		goto L86
	} else {
		goto L92
	}
L92:
	;
	goto L81
L93:
	;
	goto L81
L94:
	;
	if v238 == v239 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v244 - int32(1) {
	case 0:
		goto L99
	case 1:
		goto L98
	default:
		goto L96
	case 3:
		goto L97
	}
L96:
	;
	v257 = v24 - int32(8)
	goto L103
L97:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	if v238 == v251 {
		goto L12
	} else {
		goto L102
	}
L98:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	if v238 != v249 {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	if v238 != v247 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	goto L12
L101:
	;
	goto L12
L102:
	;
	goto L96
L103:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v265 == int32(0) {
		v315 = v238
		goto L13
	} else {
		goto L105
	}
L104:
	;
	goto L12
L105:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+40))
	if v238 == v268 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+56))
	if v238 != v270 {
		v257 = v265
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	m.T0[v272].(func(*base.Module, int32, int32))(m, v273, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v278
	v281 = v24 - int32(4)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v283
	if v282 == int32(0) {
		goto L12
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	if v282 == v283 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v288 {
	case 0:
		goto L118
	case 1:
		goto L117
	case 2:
		goto L116
	case 3:
		goto L115
	default:
		goto L114
	}
L114:
	;
	v299 = v24 - int32(8)
	goto L123
L115:
	;
	if v282 == v278 {
		goto L12
	} else {
		goto L122
	}
L116:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	if v282 != v292 {
		goto L114
	} else {
		goto L121
	}
L117:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	if v282 != v290 {
		goto L114
	} else {
		goto L120
	}
L118:
	;
	if v282 != v278 {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L12
L120:
	;
	goto L12
L121:
	;
	goto L12
L122:
	;
	goto L114
L123:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v307 == int32(0) {
		v315 = v282
		goto L13
	} else {
		goto L125
	}
L124:
	;
	goto L12
L125:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v307)+40))
	if v282 == v310 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v307)+56))
	if v282 != v312 {
		v299 = v307
		goto L123
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	goto L12
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v338
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v24-int32(24)))) = v371
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v24-int32(16)))) = v377
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v379&int32(64) == int32(0) {
		goto L6
	} else {
		goto L139
	}
L130:
	;
	if v338 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v338 != 0 {
		goto L129
	} else {
		goto L138
	}
L133:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _consts[432]))
	if v345 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v352
	v354 = int32(4547052)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v24
	*(*int32)(unsafe.Add(mBase, _consts[1439])) = v24
	goto L129
L135:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	v352 = v347
	goto L134
L136:
	;
	goto L137
L137:
	;
	v349 = int32(4547052)
	*(*int32)(unsafe.Add(mBase, _consts[432])) = v349
	v352 = v349
	goto L134
L138:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v362
	goto L129
L139:
	;
	v385 = v24 - int32(36)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v386&int32(4) != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v386 | int32(4)
	v393 = v24 + int32(12)
	v394 = int32(4547060)
	v395 = *(*int32)(unsafe.Add(mBase, _consts[1440]))
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v395
	*(*int32)(unsafe.Add(mBase, _consts[1440])) = v393
	goto L6
L141:
	;
	goto L5
}
