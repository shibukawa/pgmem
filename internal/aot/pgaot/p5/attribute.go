package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutCSV(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v18 = base.B2i32(v14 == int32(1))
	goto L3
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if l2 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v29 == int32(0) {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v53 = int32(1)
	goto L6
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v54 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v53 = base.B2i32(v49-v48 == int32(0))
	goto L6
L8:
	;
	goto L7
L9:
	;
	if v28 != v29 {
		v48 = v28
		v49 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = l1
	v34 = v25
	goto L11
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v38 == int32(0) {
		v48 = v37
		v49 = v38
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v48 = v37
	v49 = v38
	goto L8
L13:
	;
	v41 = int32(1)
	if v37 == v38 {
		v33 = v33 + v41
		v34 = v34 + v41
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if l1&int32(3) == int32(0) {
		v80 = l1
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v117 = l1
	goto L17
L17:
	;
	if v53 != 0 {
		goto L37
	} else {
		goto L38
	}
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v115 = F_pg_server_to_any(m, l1, v113, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v113 = v105 - l1
	goto L18
L20:
	;
	v84 = v80
	goto L29
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v64 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v113 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v69 = l1
	goto L25
L25:
	;
	v73 = v69 + int32(1)
	if v73&int32(3) == int32(0) {
		v80 = v73
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v105 = v73
	goto L19
L27:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v78 != 0 {
		v69 = v73
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 == v93 {
		v84 = v84 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v99 = v84
	goto L32
L31:
	;
	goto L30
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v103 != 0 {
		v99 = v99 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v105 = v99
	goto L19
L34:
	;
	goto L33
L35:
	;
	return
L36:
	;
	v117 = v115
	goto L17
L37:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	if v243 <= v240+int32(1) {
		goto L76
	} else {
		goto L77
	}
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v118 == int32(92))&v18 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v117&int32(3) == int32(0) {
		v194 = v117
		goto L59
	} else {
		goto L60
	}
L40:
	;
	v133 = v118
	v134 = v117
	goto L47
L41:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 != int32(46) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v118 == int32(0) {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)))
	if v125 != 0 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	goto L40
L47:
	;
	v140 = v133 & int32(255)
	if v140 == v21&int32(255) {
		goto L37
	} else {
		goto L49
	}
L48:
	;
	goto L39
L49:
	;
	if v140 == v20&int32(255) {
		goto L37
	} else {
		goto L50
	}
L50:
	;
	switch v140 - int32(10) {
	case 0, 3:
		goto L37
	default:
		goto L51
	}
L51:
	;
	if int32(0) <= base.I32_extend8_s(v133) {
		v158 = int32(1)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = v158 + v134
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 != 0 {
		v133 = v160
		v134 = v159
		goto L47
	} else {
		goto L56
	}
L53:
	;
	v151 = int32(1)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v152 != v151 {
		v158 = v151
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v156 = F_pg_encoding_mblen(m, v155, v134)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L35
	} else {
		goto L55
	}
L55:
	;
	v158 = v156
	goto L52
L56:
	;
	goto L48
L57:
	;
	F_appendBinaryStringInfo(m, v170, v117, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L35
	} else {
		goto L74
	}
L58:
	;
	v227 = v219 - v117
	goto L57
L59:
	;
	v198 = v194
	goto L68
L60:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v178 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v227 = int32(0)
	goto L57
L62:
	;
	goto L63
L63:
	;
	v183 = v117
	goto L64
L64:
	;
	v187 = v183 + int32(1)
	if v187&int32(3) == int32(0) {
		v194 = v187
		goto L59
	} else {
		goto L66
	}
L65:
	;
	v219 = v187
	goto L58
L66:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v192 != 0 {
		v183 = v187
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v207 = int32(-2139062144)
	if (int32(16843008)-v204|v204)&v207 == v207 {
		v198 = v198 + int32(4)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v213 = v198
	goto L71
L70:
	;
	goto L69
L71:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v217 != 0 {
		v213 = v213 + int32(1)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v219 = v213
	goto L58
L73:
	;
	goto L72
L74:
	;
	return
L75:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v261 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	F_appendStringInfoChar(m, v239, v20)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L35
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v240))) = uint8(v20)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v253 = v251 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v255+v253))) = uint8(v257)
	goto L75
L79:
	;
	goto L75
L80:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+8))
	if v345 <= v342+int32(1) {
		goto L103
	} else {
		goto L104
	}
L81:
	;
	v265 = v117
	v267 = v261
	v268 = v117
	goto L82
L82:
	;
	v273 = int32(255)
	v274 = v267 & v273
	if base.B2i32(v274 != v20&v273)&base.B2i32(v274 != v19&v273) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if base.Ui32(v325) <= base.Ui32(v312) {
		goto L80
	} else {
		goto L101
	}
L84:
	;
	if base.Ui32(v268) < base.Ui32(v265) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v312 = v268
	goto L86
L86:
	;
	if int32(0) <= base.I32_extend8_s(v267) {
		v324 = int32(1)
		goto L96
	} else {
		goto L97
	}
L87:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v285, v268, v265-v268)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L35
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	if v293 <= v290+int32(1) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	v312 = v265
	goto L86
L92:
	;
	F_appendStringInfoChar(m, v289, v19)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L35
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*uint8)(unsafe.Add(mBase, uint32(v297+v290))) = uint8(v19)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v303 = v301 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305+v303))) = uint8(v307)
	goto L91
L95:
	;
	goto L91
L96:
	;
	v325 = v324 + v265
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v326 != 0 {
		v265 = v325
		v267 = v326
		v268 = v312
		goto L82
	} else {
		goto L100
	}
L97:
	;
	v317 = int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v318 != v317 {
		v324 = v317
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v322 = F_pg_encoding_mblen(m, v321, v265)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L35
	} else {
		goto L99
	}
L99:
	;
	v324 = v322
	goto L96
L100:
	;
	goto L83
L101:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v328, v312, v325-v312)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L35
	} else {
		goto L102
	}
L102:
	;
	goto L80
L103:
	;
	F_appendStringInfoChar(m, v341, v20)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L35
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v342))) = uint8(v20)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v355 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357+v355))) = uint8(v359)
	return
L106:
	;
	return
}
