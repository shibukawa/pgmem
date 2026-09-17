package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_exec_stmt_execsql(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
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
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[0]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[1])))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_exec_prepare_plan(m, l0, v20, int32(2048))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 == v30 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v98 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v83)
	goto L8
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v36
	goto L14
L13:
	;
	v42 = v39
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v47 = int32(0)
	goto L15
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43+v47<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	switch v60 - int32(158) {
	case 0, 5:
		goto L18
	case 1, 2, 3, 4:
		goto L17
	default:
		goto L19
	}
L16:
	;
	goto L9
L17:
	;
	v70 = v47 + int32(1)
	if v70 != v42 {
		v47 = v70
		goto L15
	} else {
		goto L22
	}
L18:
	;
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v67)
	goto L9
L19:
	;
	if v60 == int32(191) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v60 != int32(103) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L16
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v20
	v101 = v99
	goto L25
L24:
	;
	v101 = v96
	goto L25
L25:
	;
	v102 = int32(4)
	v103 = v17 & v102
	v105 = v19 & v102
	v108 = base.B2i32(v103|v105 != int32(0))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v109 != int32(1) {
		v120 = v96
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v123 = F_SPI_execute_plan_with_paramlist(m, v121, v101, v122, v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L42
	}
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v113 != 0 {
		v120 = int32(2)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v114 = int32(2)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v117 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v114
	goto L31
L30:
	;
	v118 = int32(1)
	goto L31
L31:
	;
	if v103|v105 != int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v119 = v114
	goto L34
L33:
	;
	v119 = v118
	goto L34
L34:
	;
	v120 = v119
	goto L26
L35:
	;
	v216 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[3]))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v220 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L36:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201+v202<<(uint(int32(2))%32))))
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2]))
	v211 = int32(0)
	F_assign_simple_var(m, l0, v206, base.B2i32(v208 != int64(0)), v211, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L57
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L53
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L49
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142<<(uint(int32(2))%32))))
	v147 = int32(0)
	F_assign_simple_var(m, l0, v146, v147, v147, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L44
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32))))
	v134 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[2]))
	v137 = int32(0)
	F_assign_simple_var(m, l0, v132, base.B2i32(v134 != int64(0)), v137, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	switch v123 + int32(8) {
	case 0:
		goto L38
	default:
		goto L37
	case 6:
		goto L39
	case 12, 14:
		goto L35
	case 13:
		goto L36
	case 15, 16, 17, 19, 20, 21, 26, 27:
		goto L41
	case 22:
		goto L40
	}
L43:
	;
	goto L35
L44:
	;
	goto L35
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_1), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_3), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_5), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_6), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v189 = F_SPI_result_code_string(m, v123)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v188
	F_errmsg_internal(m, int32(_a_F_exec_stmt_execsql_7), v14)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_8), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	goto L35
L58:
	;
	m.G0 = v14 + int32(48)
	return
L59:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L133
	}
L60:
	;
	if v219 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v219 == int32(0) {
		goto L58
	} else {
		goto L124
	}
L63:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+v227<<(uint(int32(2))%32))))
	if base.Ui64(v216) <= base.Ui64(int64(1)) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	F_exec_move_row(m, l0, v231, v340, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L114
	}
L65:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v340 = v337
	goto L64
L66:
	;
	if base.I32_wrap_i64(v216) == int32(1) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v271 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L69:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v238 != int32(1) {
		v340 = int32(0)
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+492)))
	if v243 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v246 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	v248 = int32(0)
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L75
	}
L74:
	;
	v248 = v246
	goto L73
L75:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_9), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	if v248 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v248
	F_errdetail_internal(m, int32(_a_F_exec_stmt_execsql_10), v14+int32(32))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_11), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v108|v274)&int32(1) == int32(0) {
		goto L65
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+492)))
	if v281 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v308 = F_errstart(m, v306, int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L104
	}
L88:
	;
	v295 = int32(21)
	if int32(base.Ui32(v105)>>(uint(int32(2))%32)) != 0 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v284 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v290 = int32(0)
	if v271 == v290 {
		v294 = v290
		goto L88
	} else {
		goto L94
	}
L92:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v286 == int32(0) {
		v294 = v284
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v305 = v284
	v306 = int32(21)
	goto L87
L94:
	;
	v305 = v290
	v306 = int32(21)
	goto L87
L95:
	;
	v301 = int32(19)
	goto L97
L96:
	;
	v301 = int32(0)
	goto L97
L97:
	;
	if v103 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v302 = v295
	goto L100
L99:
	;
	v302 = v301
	goto L100
L100:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v303 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v304 = v295
	goto L103
L102:
	;
	v304 = v302
	goto L103
L103:
	;
	v305 = v294
	v306 = v304
	goto L87
L104:
	;
	if v308 == int32(0) {
		goto L65
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_12), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	if v305 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v305
	F_errdetail_internal(m, int32(_a_F_exec_stmt_execsql_10), v14+int32(16))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_errhint(m, int32(_a_F_exec_stmt_execsql_13), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_14), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L65
L114:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v344 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_SPI_freetuptable(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v349 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+20))
	F_MemoryContextReset(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmt_execsql[3]))
	F_SPI_freetuptable(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	goto L58
L124:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmt_execsql_0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_15), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	if v123 == int32(5) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_errhint(m, int32(_a_F_exec_stmt_execsql_16), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_17), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_exec_stmt_execsql_18), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_exec_stmt_execsql_2), int32(_a_F_exec_stmt_execsql_19), int32(_a_F_exec_stmt_execsql_4))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
