package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetJsonPathVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l0 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v406
	m.G0 = v14 + int32(16)
	return v404
L2:
	;
	v404 = int32(0)
	v406 = int32(-1)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		v404 = v6
		v406 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = v19
	goto L7
L6:
	;
	v25 = v22
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v6
	v35 = int32(1)
	goto L8
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26+v33<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if l2 == v43 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v99 = F_palloc(m, int32(32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	goto L9
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if l2 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v93 = int32(1)
	v96 = v33 + v93
	if v25 != v96 {
		v33 = v96
		v35 = v35 + v93
		goto L8
	} else {
		goto L28
	}
L14:
	;
	if v90 == int32(0) {
		goto L10
	} else {
		goto L27
	}
L15:
	;
	v90 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = v45
	v53 = l1
	v54 = l2
	v55 = v51
	goto L22
L19:
	;
	v78 = l1
	v82 = int32(0)
	goto L20
L20:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v90 = v82 - v83
	goto L14
L21:
	;
	v78 = v73
	v82 = v75
	goto L20
L22:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if base.B2i32(v55 != v57)|base.B2i32(v57 == int32(0)) != 0 {
		v73 = v53
		v75 = v55
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v73 = v67
	v75 = int32(0)
	goto L21
L24:
	;
	v63 = v54 - int32(1)
	if v63 == int32(0) {
		v73 = v53
		v75 = v55
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v66 = int32(1)
	v67 = v53 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v68 != 0 {
		v52 = v52 + v66
		v53 = v67
		v54 = v63
		v55 = v68
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L13
L28:
	;
	goto L2
L29:
	;
	return int32(0)
L30:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v103 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v378
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v99)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v380
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v99)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v382
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v384
	v404 = v99
	v406 = v35
	goto L1
L32:
	;
	v106 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v106
	goto L31
L33:
	;
	goto L34
L34:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v42)+16))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v111 <= int32(1042) {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v320 = F_pg_detoast_datum(m, base.I32_wrap_i64(v110))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L29
	} else {
		goto L113
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+8)) = uint8(base.B2i32(v110 != int64(0)))
	goto L31
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L29
	} else {
		goto L108
	}
L38:
	;
	if v111 == int32(114) {
		goto L35
	} else {
		goto L107
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(18)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v261 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L40:
	;
	v215 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v110))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L29
	} else {
		goto L79
	}
L41:
	;
	v206 = F_DirectFunctionCall1Coll(m, int32(1541), int32(0), v110)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L29
	} else {
		goto L77
	}
L42:
	;
	v196 = F_DirectFunctionCall1Coll(m, int32(1540), int32(0), v110)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L29
	} else {
		goto L75
	}
L43:
	;
	v186 = F_DirectFunctionCall1Coll(m, int32(1539), int32(0), v110)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L29
	} else {
		goto L73
	}
L44:
	;
	v176 = F_DirectFunctionCall1Coll(m, int32(1538), int32(0), v110)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L29
	} else {
		goto L71
	}
L45:
	;
	v166 = F_DirectFunctionCall1Coll(m, int32(1537), int32(0), v110)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L29
	} else {
		goto L69
	}
L46:
	;
	switch v111 - int32(16) {
	case 0:
		goto L36
	case 1, 2, 3, 6, 8:
		goto L37
	case 4:
		goto L43
	case 5:
		goto L45
	case 7:
		goto L44
	case 9:
		goto L40
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v111 <= int32(1183) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	switch v111 - int32(700) {
	case 0:
		goto L42
	case 1:
		goto L41
	default:
		goto L38
	}
L50:
	;
	if v111 != int32(1700) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(32)
	goto L31
L52:
	;
	if base.Ui32(v111-int32(1082)) < base.Ui32(int32(2)) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if int32(1699) < v111 {
		goto L50
	} else {
		goto L58
	}
L55:
	;
	if v111 == int32(1043) {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	if v111 == int32(1114) {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L37
L58:
	;
	if v111 == int32(1184) {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	if v111 != int32(1266) {
		goto L37
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	if v111 != int32(3802) {
		goto L37
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v161 = F_pg_detoast_datum(m, base.I32_wrap_i64(v110))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L29
	} else {
		goto L68
	}
L64:
	;
	v147 = F_pg_detoast_datum(m, base.I32_wrap_i64(v110))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	v150 = v147 + int32(4)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+7)))
	if v151&int32(16) == int32(0) {
		goto L39
	} else {
		goto L66
	}
L66:
	;
	v156 = F_JsonbExtractScalar(m, v150, v99)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	goto L31
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v161
	goto L31
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v171 = F_pg_detoast_datum(m, base.I32_wrap_i64(v166))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v171
	goto L31
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v181 = F_pg_detoast_datum(m, base.I32_wrap_i64(v176))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L29
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v181
	goto L31
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v191 = F_pg_detoast_datum(m, base.I32_wrap_i64(v186))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L29
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v191
	goto L31
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v201 = F_pg_detoast_datum(m, base.I32_wrap_i64(v196))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L29
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v201
	goto L31
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(2)
	v211 = F_pg_detoast_datum(m, base.I32_wrap_i64(v206))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v211
	goto L31
L79:
	;
	v217 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v217
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v221&v217 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v224 = v217
	goto L82
L81:
	;
	v224 = int32(4)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v215 + v224
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v227 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v233 == int32(18) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v227&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v236 = int32(16)
	goto L88
L87:
	;
	v236 = int32(0)
	goto L88
L88:
	;
	if base.Ui32((v233-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v243 = int32(4)
	goto L91
L90:
	;
	v243 = v236
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v243
	goto L31
L92:
	;
	v247 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v227)>>(uint(v247)%32)) - v247
	goto L31
L93:
	;
	goto L94
L94:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v252)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L95:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v267 == int32(18) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	if v261&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v270 = int32(16)
	goto L100
L99:
	;
	v270 = int32(0)
	goto L100
L100:
	;
	if base.Ui32((v267-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v277 = int32(4)
	goto L103
L102:
	;
	v277 = v270
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v277
	goto L31
L104:
	;
	v281 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v261)>>(uint(v281)%32)) - v281
	goto L31
L105:
	;
	goto L106
L106:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v286)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L107:
	;
	goto L37
L108:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L29
	} else {
		goto L109
	}
L109:
	;
	v301 = F_format_type_be(m, v111)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L29
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v301
	F_errmsg(m, int32(_a_F_GetJsonPathVar_0), v14)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L29
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_GetJsonPathVar_1), int32(3394), int32(_a_F_GetJsonPathVar_2))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L29
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v322 = F_text_to_cstring(m, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L29
	} else {
		goto L114
	}
L114:
	;
	v325 = F_DirectFunctionCall1Coll(m, int32(521), int32(0), base.I64_extend_i32_u(v322))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L29
	} else {
		goto L115
	}
L115:
	;
	v328 = F_pg_detoast_datum(m, base.I32_wrap_i64(v325))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L29
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v322)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L29
	} else {
		goto L117
	}
L117:
	;
	v332 = F_pg_detoast_datum(m, v328)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L29
	} else {
		goto L118
	}
L118:
	;
	v335 = v332 + int32(4)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+7)))
	if v336&int32(16) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v339 = F_JsonbExtractScalar(m, v335, v99)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L29
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(18)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v344 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L31
L123:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v350 == int32(18) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	if v344&int32(1) != 0 {
		goto L132
	} else {
		goto L133
	}
L126:
	;
	v353 = int32(16)
	goto L128
L127:
	;
	v353 = int32(0)
	goto L128
L128:
	;
	if base.Ui32((v350-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v360 = int32(4)
	goto L131
L130:
	;
	v360 = v353
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v360
	goto L31
L132:
	;
	v364 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v344)>>(uint(v364)%32)) - v364
	goto L31
L133:
	;
	goto L134
L134:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = int32(base.Ui32(v369)>>(uint(int32(2))%32)) - int32(4)
	goto L31
}
func F_JsonTableFetchRow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableFetchRow_0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v8 = F_JsonTablePlanNextRow(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F__equalJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v24 = v3
			return v24
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v24 = v3
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					if v18 != v19 {
						v24 = v3
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
						v24 = base.B2i32(v21 == v22)
					}
				}
				return v24
			}
		}
	}
}
func F_get_json_agg_constructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_get_json_constructor_options(m, l0, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			switch v18 - int32(6) {
			case 0:
				F_resolve_special_varno(m, v17, l1, int32(1701), l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
					F_errmsg_internal(m, int32(_a_F_get_json_agg_constructor_0), v9)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_get_json_agg_constructor_1), int32(_a_F_get_json_agg_constructor_2), int32(_a_F_get_json_agg_constructor_3))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 3:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				F_get_agg_expr_helper(m, v17, l1, v17, l2, v42, l3)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			case 5:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				F_get_windowfunc_expr_helper(m, v17, l1, l2, v21, l3)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			}
		}
	}
}
func F_get_json_agg_constructor_expr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(9) {
	case 0, 2:
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v25
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v27
		v29 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v29
		v31 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v31
		v33 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v33
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
		F_get_json_constructor(m, v7+int32(8), l1)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			m.G0 = v7 + int32(48)
			return
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_get_json_agg_constructor_expr_0), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_get_json_agg_constructor_expr_1), int32(_a_F_get_json_agg_constructor_expr_2), int32(_a_F_get_json_agg_constructor_expr_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_get_json_table_nested_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v11 - int32(50) {
	case 0:
		goto L4
	case 1:
		goto L5
	default:
		goto L1
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v55, int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v45, int32(44))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L4:
	;
	if l4 == int32(0) {
		v50 = l1
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_get_json_table_nested_columns(m, l0, v14, l2, l3, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	switch v18 - int32(50) {
	case 0:
		v40 = v17
		goto L3
	case 1:
		goto L8
	default:
		goto L1
	}
L8:
	;
	v22 = v17
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_get_json_table_nested_columns(m, l0, v27, l2, l3, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	if v32 == int32(50) {
		v40 = v31
		goto L3
	} else {
		goto L13
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(51) {
		v22 = v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L1
L14:
	;
	v40 = l1
	goto L3
L15:
	;
	v50 = v40
	goto L2
L16:
	;
	v60 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_json_table_nested_columns_0), v60, v60, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	F_get_const_expr(m, v66, l2, int32(-1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = F_quote_identifier(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
	F_appendStringInfo(m, v70, int32(_a_F_get_json_table_nested_columns_1), v9)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_get_json_table_columns(m, l0, v50, l2, l3)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L1
}
func F_json_array_elements(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(_a_F_json_array_elements_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_json_array_elements_text(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_elements_worker(m, l0, int32(_a_F_json_array_elements_text_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_json_build_array_noargs(m *base.Module, l0 int32) int64 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_cstring_to_text_with_len(m, int32(_a_F_json_build_array_noargs_0), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_json_each(m *base.Module, l0 int32) int64 {
	var v6 int32
	_ = v6
	F_each_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_json_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v11 = v8 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_cstring_to_text(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		F_makeJsonLexContext(m, v11, v13, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = F_pg_parse_json_or_errsave(m, v11, int32(_a_F_json_in_0), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				if v22 == int32(0) {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
					v30 = int64(0)
				} else {
					v30 = base.I64_extend_i32_u(v13)
				}
				m.G0 = v8 + int32(80)
				return v30
			}
		}
	}
}
func F_json_lex_number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = l1 - v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14) <= base.Ui32(v13) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != 0 {
		goto L74
	} else {
		goto L75
	}
L2:
	;
	if base.Ui32(v14) <= base.Ui32(v72) {
		v134 = v70
		v136 = v72
		v139 = v75
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v70 = l1
	v72 = v13
	v75 = v11
	goto L2
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 == int32(48) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v70 = v59
	v72 = v61
	v75 = int32(0)
	goto L2
L7:
	;
	v19 = int32(1)
	v59 = l1 + v19
	v61 = v13 + v19
	goto L6
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(8)) < base.Ui32((v16-int32(49))&int32(255)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = l1
	v72 = v13
	v75 = v11
	goto L2
L11:
	;
	goto L12
L12:
	;
	v33 = l1
	v38 = v13
	goto L13
L13:
	;
	if v38 == v14-int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v59 = v48
	v61 = v45
	goto L6
L15:
	;
	v302 = v12 + v14
	v304 = v14
	v305 = int32(0)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v44 = int32(1)
	v45 = v38 + v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v48 = v33 + v44
	if base.Ui32((v46-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v33 = v48
		v38 = v45
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if base.Ui32(v14) <= base.Ui32(v136) {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v77 != int32(46) {
		v134 = v70
		v136 = v72
		v139 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v80 = int32(1)
	v82 = v70 + v80
	v84 = v72 + v80
	if v14 == v84 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v302 = v82
	v304 = v14
	v305 = v80
	goto L1
L23:
	;
	goto L24
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.Ui32((v86-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v134 = v124
	v136 = v125
	v139 = v129
	goto L19
L26:
	;
	v124 = v82
	v125 = v84
	v129 = int32(1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v95 = v72 + int32(2)
	if base.Ui32(v95) < base.Ui32(v14) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v14
	goto L31
L30:
	;
	v97 = v95
	goto L31
L31:
	;
	v102 = v82
	v103 = v84
	goto L32
L32:
	;
	v108 = int32(1)
	v109 = v102 + v108
	v111 = v103 + v108
	if base.Ui32(v14) <= base.Ui32(v111) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v124 = v109
	v125 = v111
	v129 = v75
	goto L25
L34:
	;
	v302 = v109
	v304 = v97
	v305 = v75
	goto L1
L35:
	;
	goto L36
L36:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.Ui32((v113-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v102 = v109
		v103 = v111
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if base.Ui32(v14) <= base.Ui32(v212) {
		goto L62
	} else {
		goto L63
	}
L39:
	;
	v211 = v134
	v212 = v136
	v215 = v139
	goto L38
L40:
	;
	goto L41
L41:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v141|int32(32) != int32(101) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v211 = v134
	v212 = v136
	v215 = v139
	goto L38
L43:
	;
	goto L44
L44:
	;
	v146 = int32(1)
	v148 = v134 + v146
	v150 = v136 + v146
	if base.Ui32(v14) <= base.Ui32(v150) {
		v159 = v150
		v160 = v148
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v159 == v14 {
		v302 = v160
		v304 = v14
		v305 = v146
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	switch v152 - int32(43) {
	case 0, 2:
		goto L47
	default:
		v159 = v150
		v160 = v148
		goto L45
	}
L47:
	;
	v155 = int32(2)
	v159 = v136 + v155
	v160 = v134 + v155
	goto L45
L48:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if base.Ui32((v162-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v211 = v201
	v212 = v197
	v215 = v205
	goto L38
L50:
	;
	v197 = v159
	v201 = v160
	v205 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v171 = v159 + int32(1)
	if base.Ui32(v171) < base.Ui32(v14) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v173 = v14
	goto L55
L54:
	;
	v173 = v171
	goto L55
L55:
	;
	v175 = v159
	v179 = v160
	goto L56
L56:
	;
	v184 = int32(1)
	v185 = v179 + v184
	v187 = v175 + v184
	if base.Ui32(v14) <= base.Ui32(v187) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v197 = v187
	v201 = v185
	v205 = v139
	goto L49
L58:
	;
	v302 = v185
	v304 = v173
	v305 = v139
	goto L1
L59:
	;
	goto L60
L60:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32((v189-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v175 = v187
		v179 = v185
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	v302 = v211
	v304 = v212
	v305 = v215
	goto L1
L63:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v211))))
	v220 = int32(255)
	v236 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32((v217-int32(48))&v220) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v217&int32(-33)-int32(65))&v220) < base.Ui32(int32(26)))|base.B2i32(v217 == int32(95)) == v236)&base.B2i32(v236 <= v217) != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v243 = int32(1)
	v245 = v212 + v243
	if v14 != v245 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v248 = v245
	v252 = v211
	goto L68
L66:
	;
	goto L67
L67:
	;
	v302 = v211 + (v14 - v212)
	v304 = v14
	v305 = v243
	goto L1
L68:
	;
	v258 = v252 + int32(1)
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v252)+1)))
	v262 = int32(255)
	v277 = int32(0)
	if base.B2i32(base.Ui32((v259-int32(48))&v262) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v259&int32(-33)-int32(65))&v262) < base.Ui32(int32(26)))|(base.B2i32(v259 == int32(95))|base.B2i32(v259 < v277)) == v277 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v302 = v258
	v304 = v248
	v305 = v243
	goto L1
L71:
	;
	goto L72
L72:
	;
	v284 = v248 + int32(1)
	if v284 != v14 {
		v248 = v284
		v252 = v258
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v304
	goto L76
L75:
	;
	goto L76
L76:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v308 != int32(1) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	return v338
L78:
	;
	if l2 != 0 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v312 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v304) < base.Ui32(v313) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v311+int32(4), v317, v302-v317)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	return int32(0)
L83:
	;
	if l2 == int32(0) {
		v338 = int32(1)
		goto L77
	} else {
		goto L84
	}
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v305)
	return int32(1)
L85:
	;
	v338 = int32(0)
	goto L77
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v305)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v302
	if v305 != 0 {
		v338 = int32(15)
		goto L77
	} else {
		goto L89
	}
L89:
	;
	goto L85
}
func F_json_manifest_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 - int32(6) {
	case 0, 4:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_json_manifest_array_end_0)
		m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v18, int32(_a_F_json_manifest_array_end_1), v6)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_json_to_record(m *base.Module, l0 int32) int64 {
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_json_to_record_0), int32(1), v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_json_to_recordset(m *base.Module, l0 int32) int64 {
	var v8 int32
	_ = v8
	F_populate_recordset_worker(m, l0, int32(_a_F_json_to_recordset_0), int32(1), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_json_typeof(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = v6 + int32(12)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		F_makeJsonLexContext(m, v9, v11, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = F_json_lex(m, v9)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				if v18 != 0 {
					F_json_errsave_error(m, v18, v9, int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
						v24 = int32(1)
						v25 = v23 - v24
						if base.B2i32(base.Ui32(v25) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1815))>>(uint(v25)%32))&v24) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
								F_errmsg_internal(m, int32(_a_F_json_typeof_0), v6)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_json_typeof_1), int32(1879), int32(_a_F_json_typeof_2))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_json_typeof[0])))
							v52 = F_cstring_to_text(m, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								m.G0 = v6 + int32(80)
								return base.I64_extend_i32_u(v52)
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
					v24 = int32(1)
					v25 = v23 - v24
					if base.B2i32(base.Ui32(v25) <= base.Ui32(int32(10)))&(int32(base.Ui32(int32(1815))>>(uint(v25)%32))&v24) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
							F_errmsg_internal(m, int32(_a_F_json_typeof_0), v6)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_json_typeof_1), int32(1879), int32(_a_F_json_typeof_2))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_json_typeof[0])))
						v52 = F_cstring_to_text(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							m.G0 = v6 + int32(80)
							return base.I64_extend_i32_u(v52)
						}
					}
				}
			}
		}
	}
}
func F_json_unique_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v2 == int32(1) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		F_pfree(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
func F_json_unique_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v9 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v15 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = F_hash_search(m, v18, v7+int32(4), int32(1), v7+int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
	F_pfree(m, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_makeJsonTablePathSpec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v8 = F_palloc0(m, int32(20))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(123)
		v15 = F_palloc0(m, int32(20))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(2044404432968)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
			if l1 != 0 {
				v22 = F_pstrdup(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
					return v8
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
				return v8
			}
		}
	}
}
