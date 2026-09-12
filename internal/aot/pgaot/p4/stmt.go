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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1473]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1474])))
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
	v27 = int32(4)
	v28 = v17 & v27
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v31 == int32(0) {
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
	v34 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 == v34 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v100 = int32(2)
	v103 = int32(base.Ui32(v19&v27) >> (uint(v100) % 32))
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	if v106 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v87)
	goto L8
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v43 = int32(0)
	if v43 < v40 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = v40
	goto L14
L13:
	;
	v46 = v43
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v51 = int32(0)
	goto L15
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47+v51<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	switch v64 - int32(158) {
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
	v74 = v51 + int32(1)
	if v74 != v46 {
		v51 = v74
		goto L15
	} else {
		goto L22
	}
L18:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v71)
	goto L9
L19:
	;
	if v64 == int32(191) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v64 != int32(103) {
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = v20
	v109 = v107
	goto L25
L24:
	;
	v109 = v104
	goto L25
L25:
	;
	v110 = int32(base.Ui32(v28)>>(uint(v100)%32)) | v103
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v111 != int32(1) {
		v124 = v104
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v127 = F_SPI_execute_plan_with_paramlist(m, v125, v109, v126, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L42
	}
L27:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v115 != 0 {
		v124 = int32(2)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v116 = int32(2)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v119 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v120 = v116
	goto L31
L30:
	;
	v120 = int32(1)
	goto L31
L31:
	;
	if v110&int32(1) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v123 = v116
	goto L34
L33:
	;
	v123 = v120
	goto L34
L34:
	;
	v124 = v123
	goto L26
L35:
	;
	v220 = *(*int64)(unsafe.Add(mBase, _consts[511]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v220
	v223 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+18)))
	if v224 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L36:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(int32(2))%32))))
	v212 = *(*int64)(unsafe.Add(mBase, _consts[511]))
	v215 = int32(0)
	F_assign_simple_var(m, l0, v210, base.B2i32(v212 != int64(0)), v215, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L57
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L53
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L49
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145+v146<<(uint(int32(2))%32))))
	v151 = int32(0)
	F_assign_simple_var(m, l0, v150, v151, v151, v151)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L44
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131+v132<<(uint(int32(2))%32))))
	v138 = *(*int64)(unsafe.Add(mBase, _consts[511]))
	v141 = int32(0)
	F_assign_simple_var(m, l0, v136, base.B2i32(v138 != int64(0)), v141, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	switch v127 + int32(8) {
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
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(533639), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(501532), int32(4334), int32(301773))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
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
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(533678), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(501532), int32(4340), int32(301773))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v193 = F_SPI_result_code_string(m, v127)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v192
	F_errmsg_internal(m, int32(205624), v14)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(501532), int32(4345), int32(301773))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L133
	}
L60:
	;
	if v223 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v223 == int32(0) {
		goto L58
	} else {
		goto L124
	}
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229+v231<<(uint(int32(2))%32))))
	if base.Ui64(v220) <= base.Ui64(int64(1)) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	F_exec_move_row(m, l0, v235, v344, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L114
	}
L65:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v344 = v341
	goto L64
L66:
	;
	if base.I32_wrap_i64(v220) == int32(1) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v275 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L69:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v242 != int32(1) {
		v344 = int32(0)
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+492)))
	if v247 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v250 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	v252 = int32(0)
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L75
	}
L74:
	;
	v252 = v250
	goto L73
L75:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(114049), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	if v252 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v252
	F_errdetail_internal(m, int32(200676), v14+int32(32))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_errfinish(m, int32(501532), int32(4387), int32(301773))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v110|v278)&int32(1) == int32(0) {
		goto L65
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+492)))
	if v285 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v312 = F_errstart(m, v310, int32(557877))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L104
	}
L88:
	;
	v301 = int32(21)
	if v103 != 0 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v288 = F_format_expr_params(m, l0, v20)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v296 = int32(0)
	if v275 == v296 {
		v300 = v296
		goto L88
	} else {
		goto L94
	}
L92:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+19)))
	if v290&int32(1) == int32(0) {
		v300 = v288
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v309 = v288
	v310 = int32(21)
	goto L87
L94:
	;
	v309 = v296
	v310 = int32(21)
	goto L87
L95:
	;
	v305 = int32(19)
	goto L97
L96:
	;
	v305 = int32(0)
	goto L97
L97:
	;
	if v28 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v306 = v301
	goto L100
L99:
	;
	v306 = v305
	goto L100
L100:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v307 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v308 = v301
	goto L103
L102:
	;
	v308 = v306
	goto L103
L103:
	;
	v309 = v300
	v310 = v308
	goto L87
L104:
	;
	if v312 == int32(0) {
		goto L65
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(30795), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	if v309 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v309
	F_errdetail_internal(m, int32(200676), v14+int32(16))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_errhint(m, int32(662098), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_errfinish(m, int32(501532), int32(4410), int32(301773))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L65
L114:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v348 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_SPI_freetuptable(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
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
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v353 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+20))
	F_MemoryContextReset(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	F_SPI_freetuptable(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
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
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(505915), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	if v127 == int32(5) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_errhint(m, int32(654398), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_errfinish(m, int32(501532), int32(4427), int32(301773))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
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
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(506102), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(501532), int32(4363), int32(301773))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
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
