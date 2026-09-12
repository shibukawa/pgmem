package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generic_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v14 = v12 + v13
	if base.Ui32(v12) < base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+2)))
	v20 = v12 + v17 + int32(4)
	if base.Ui32(v20) < base.Ui32(v14) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	v23 = v20
	v25 = v16
	v26 = v17
	goto L7
L5:
	;
	v46 = v16
	v47 = v17
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46 & int32(65535)
	F_appendStringInfo(m, l0, int32(47472), v9)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v25 & int32(65535)
	F_appendStringInfo(m, l0, int32(719658), v9+int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v46 = v37
	v47 = v38
	goto L6
L9:
	;
	return
L10:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+2)))
	v41 = v23 + v38 + int32(4)
	if base.Ui32(v41) < base.Ui32(v14) {
		v23 = v41
		v25 = v37
		v26 = v38
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L3
}
func F_transformGenericOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = F_untransformRelOptions(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 == int32(0) {
		v201 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v204 = int32(0)
	if v201 == v204 {
		v396 = v204
		goto L59
	} else {
		goto L60
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v21 <= int32(0) {
		v201 = v15
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = v15
	v32 = int32(0)
	goto L9
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32<<(uint(int32(2))%32))))
	if v31 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	switch v104 {
	case 0, 2:
		goto L34
	case 1:
		goto L35
	case 3:
		goto L36
	default:
		goto L6
	}
L12:
	;
	v98 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v43 = int32(0)
	if v43 < v42 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = v42
	goto L17
L16:
	;
	v46 = v43
	goto L17
L17:
	;
	v49 = int32(0)
	goto L18
L18:
	;
	if v49 == v46 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v98 = v65
	goto L11
L20:
	;
	v98 = int32(0)
	goto L11
L21:
	;
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v65 = v49<<(uint(int32(2))%32) + v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(0) {
		v91 = v71
		v92 = v72
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v92-v91 != 0 {
		v49 = v49 + int32(1)
		goto L18
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if v71 != v72 {
		v91 = v71
		v92 = v72
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v76 = v67
	v77 = v68
	goto L27
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v80
		v92 = v81
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v91 = v80
	v92 = v81
	goto L24
L29:
	;
	v84 = int32(1)
	if v80 == v81 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L19
L32:
	;
	goto L10
L33:
	;
	v116 = v32 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v116 < v117 {
		v31 = v114
		v32 = v116
		goto L9
	} else {
		goto L42
	}
L34:
	;
	if v98 != 0 {
		goto L7
	} else {
		goto L40
	}
L35:
	;
	if v98 == int32(0) {
		goto L8
	} else {
		goto L39
	}
L36:
	;
	if v98 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v107 = F_list_delete_cell(m, v31, v98)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v114 = v107
	goto L33
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v38
	v114 = v31
	goto L33
L40:
	;
	v112 = F_lappend(m, v31, v38)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v114 = v112
	goto L33
L42:
	;
	v201 = v114
	goto L3
L43:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v126
	F_errmsg(m, int32(416914), v13+int32(48))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(485946), int32(160), int32(136224))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v145
	F_errmsg(m, int32(416914), v13-int32(-64))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(485946), int32(169), int32(136224))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v164
	F_errmsg(m, int32(408986), v13+int32(80))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(485946), int32(179), int32(136224))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v180
	F_errmsg_internal(m, int32(677196), v13+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(485946), int32(185), int32(136224))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L118
	}
L59:
	;
	if l3 != 0 {
		goto L110
	} else {
		goto L111
	}
L60:
	;
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v208 <= v207 {
		v396 = v207
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v216 = v207
	v220 = int32(0)
	goto L62
L62:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v216<<(uint(int32(2))%32))))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	v228 = F_defGetString(m, v226)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v379 = int32(0)
	if v373 == v379 {
		v396 = v379
		goto L59
	} else {
		goto L108
	}
L64:
	;
	v230 = int32(61)
	v231 = F___strchrnul(m, v227, v230)
	mBase = m.M
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v233 == v230 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v237 != 0 {
		goto L58
	} else {
		goto L69
	}
L66:
	;
	v237 = v231
	goto L68
L67:
	;
	v237 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	if v227&int32(3) == int32(0) {
		v261 = v227
		goto L72
	} else {
		goto L73
	}
L70:
	;
	if v228&int32(3) == int32(0) {
		v318 = v228
		goto L89
	} else {
		goto L90
	}
L71:
	;
	v294 = v286 - v227
	goto L70
L72:
	;
	v265 = v261
	goto L81
L73:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v245 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v294 = int32(0)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v250 = v227
	goto L77
L77:
	;
	v254 = v250 + int32(1)
	if v254&int32(3) == int32(0) {
		v261 = v254
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v286 = v254
	goto L71
L79:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v259 != 0 {
		v250 = v254
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v274 = int32(-2139062144)
	if (int32(16843008)-v271|v271)&v274 == v274 {
		v265 = v265 + int32(4)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v280 = v265
	goto L84
L83:
	;
	goto L82
L84:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v284 != 0 {
		v280 = v280 + int32(1)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v286 = v280
	goto L71
L86:
	;
	goto L85
L87:
	;
	v352 = v294 + v351
	v355 = F_palloc(m, v352+int32(6))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L104
	}
L88:
	;
	v351 = v343 - v228
	goto L87
L89:
	;
	v322 = v318
	goto L98
L90:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v302 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v351 = int32(0)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v307 = v228
	goto L94
L94:
	;
	v311 = v307 + int32(1)
	if v311&int32(3) == int32(0) {
		v318 = v311
		goto L89
	} else {
		goto L96
	}
L95:
	;
	v343 = v311
	goto L88
L96:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v316 != 0 {
		v307 = v311
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v331 = int32(-2139062144)
	if (int32(16843008)-v328|v328)&v331 == v331 {
		v322 = v322 + int32(4)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v337 = v322
	goto L101
L100:
	;
	goto L99
L101:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if v341 != 0 {
		v337 = v337 + int32(1)
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v343 = v337
	goto L88
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v352<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227
	v367 = F_pg_sprintf(m, v355+int32(4), int32(173528), v13)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v373 = F_accumArrayResult(m, v220, v355, int32(0), int32(25), v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v376 = v216 + int32(1)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v376 < v377 {
		v216 = v376
		v220 = v373
		goto L62
	} else {
		goto L107
	}
L107:
	;
	goto L63
L108:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v384 = F_makeArrayResult(m, v373, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v396 = v384
	goto L59
L110:
	;
	if v396 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	m.G0 = v13 + int32(96)
	return v396
L113:
	;
	v401 = v396
	goto L115
L114:
	;
	v399 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	v402 = F_OidFunctionCall2Coll(m, l3, int32(0), v401, l0)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	v401 = v399
	goto L115
L117:
	;
	goto L112
L118:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v227
	F_errmsg(m, int32(704711), v13+int32(16))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(485946), int32(87), int32(25890))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
