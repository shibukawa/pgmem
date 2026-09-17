package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_before_server_shutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	v8 = m.G0
	v10 = v8 - int32(176)
	m.G0 = v10
	v13 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v10 + int32(176)
	return
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[0]))
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v17)
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[1])) = v20
	v24 = F_errstart(m, int32(13), v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(_a_F_pgstat_before_server_shutdown_0)
	F_errmsg_internal(m, int32(_a_F_pgstat_before_server_shutdown_1), v10-int32(-64))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v40 = F_AllocateFile(m, int32(_a_F_pgstat_before_server_shutdown_2), int32(_a_F_pgstat_before_server_shutdown_3))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	F_errfinish(m, int32(_a_F_pgstat_before_server_shutdown_4), int32(1587), int32(_a_F_pgstat_before_server_shutdown_5))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = int32(27638967)
	v68 = F_fwrite(m, v10+int32(172), int32(4), int32(1), v40)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_pgstat_before_server_shutdown_2)
	F_errmsg(m, int32(_a_F_pgstat_before_server_shutdown_6), v10)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_pgstat_before_server_shutdown_4), int32(1598), int32(_a_F_pgstat_before_server_shutdown_5))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v70 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v70
	v73 = v70
	goto L21
L21:
	;
	if base.Ui32(int32(13)) <= base.Ui32(v73) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v160 = v10 + int32(144)
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[2]))
	v163 = int32(0)
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v160)+4)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v160)+12)) = v164
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+24)) = uint8(v163)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = int32(-1)
	goto L48
L23:
	;
	v155 = v73 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v155
	if v155 != int32(33) {
		v73 = v155
		goto L21
	} else {
		goto L47
	}
L24:
	;
	v114 = v111 + v113
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[1]))
	if v116 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v83 = v73 - int32(24)
	if base.Ui32(int32(8)) < base.Ui32(v83) {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(v73) <= base.Ui32(int32(6)) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[3]))
	if v87 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87+v73<<(uint(int32(2))%32)-int32(96))))
	if v95 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v100 = int32(5)
	if v99&v100 == v100 {
		v111 = v83
		v112 = v95
		v113 = int32(_a_F_pgstat_before_server_shutdown_7)
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v111 = v73
	v112 = v73*int32(72) + int32(_a_F_pgstat_before_server_shutdown_8)
	v113 = int32(_a_F_pgstat_before_server_shutdown_9)
	goto L24
L33:
	;
	if base.Ui32(v73) <= base.Ui32(int32(12)) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)+64))
	m.T0[v122].(func(*base.Module))(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L39
	}
L35:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v119)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v121 != 0 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v125)
	goto L33
L40:
	;
	F_do_putc(m, int32(70), v40)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v137 = v129 + int32(_a_F_pgstat_before_server_shutdown_10)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v73<<(uint(int32(2))%32))+uint32(_c_F_pgstat_before_server_shutdown[4])))
	v137 = v136
	goto L40
L44:
	;
	v145 = F_fwrite(m, v10+int32(80), int32(4), int32(1), v40)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v149 = F_fwrite(m, v137, v147, int32(1), v40)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L23
L47:
	;
	goto L22
L48:
	;
	v172 = F_dshash_seq_next(m, v160)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v172 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v174 = v172
	goto L53
L51:
	;
	goto L52
L52:
	;
	F_dshash_seq_term(m, v10+int32(144))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L91
	}
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[5]))
	if v182 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+16)))
	if v185 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v303 = F_dshash_seq_next(m, v10+int32(144))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L89
	}
L60:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if base.B2i32(base.Ui32(v186-int32(1)) < base.Ui32(int32(12)))|base.B2i32(base.Ui32(v186-int32(24)) < base.Ui32(int32(9))) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v200 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[6]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v221 = F_dsa_get_address(m, v219, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L68
	}
L64:
	;
	if v200 == int32(0) {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v204
	F_errmsg_internal(m, int32(_a_F_pgstat_before_server_shutdown_11), v10+int32(48))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_pgstat_before_server_shutdown_4), int32(1667), int32(_a_F_pgstat_before_server_shutdown_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L59
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if base.Ui32(v223-int32(1)) <= base.Ui32(int32(11)) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v241&int32(4) == int32(0) {
		goto L59
	} else {
		goto L73
	}
L70:
	;
	v240 = v223*int32(72) + int32(_a_F_pgstat_before_server_shutdown_8)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[3]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233+v223<<(uint(int32(2))%32)-int32(96))))
	v240 = v239
	goto L69
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	if v246 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if base.Ui32(v272-int32(1)) <= base.Ui32(int32(11)) {
		goto L85
	} else {
		goto L86
	}
L75:
	;
	F_do_putc(m, int32(83), v40)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v257 = v10 + int32(80)
	m.T0[v246].(func(*base.Module, int32, int32, int32))(m, v174, v221, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v254 = F_fwrite(m, v174, int32(16), int32(1), v40)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	F_do_putc(m, int32(78), v40)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v265 = F_fwrite(m, v174, int32(4), int32(1), v40)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v269 = F_fwrite(m, v257, int32(64), int32(1), v40)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L74
L84:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
	v294 = F_fwrite(m, v221+v290, v292, int32(1), v40)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L88
	}
L85:
	;
	v289 = v272*int32(72) + int32(_a_F_pgstat_before_server_shutdown_8)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_before_server_shutdown[3]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282+v272<<(uint(int32(2))%32)-int32(96))))
	v289 = v288
	goto L84
L88:
	;
	goto L59
L89:
	;
	if v303 != 0 {
		v174 = v303
		goto L53
	} else {
		goto L90
	}
L90:
	;
	goto L54
L91:
	;
	F_do_putc(m, int32(69), v40)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	goto L93
L93:
	;
	if int32(base.Ui32(v319)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v346 = F_FreeFile(m, v40)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L105
	}
L97:
	;
	if v326 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v342 = F_FreeFile(m, v40)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L104
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_pgstat_before_server_shutdown_2)
	F_errmsg(m, int32(_a_F_pgstat_before_server_shutdown_12), v10+int32(32))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_pgstat_before_server_shutdown_4), int32(1719), int32(_a_F_pgstat_before_server_shutdown_5))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v345 = F_unlink(m, int32(_a_F_pgstat_before_server_shutdown_2))
	mBase = m.M
	goto L3
L105:
	;
	if v346 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v352 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v373 = F_durable_rename(m, int32(_a_F_pgstat_before_server_shutdown_2), int32(_a_F_pgstat_before_server_shutdown_0), int32(15))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L116
	}
L109:
	;
	if v352 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v369 = F_unlink(m, int32(_a_F_pgstat_before_server_shutdown_2))
	mBase = m.M
	goto L3
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_pgstat_before_server_shutdown_2)
	F_errmsg(m, int32(_a_F_pgstat_before_server_shutdown_13), v10+int32(16))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_pgstat_before_server_shutdown_4), int32(1728), int32(_a_F_pgstat_before_server_shutdown_5))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	goto L112
L116:
	;
	if int32(0) <= v373 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	v378 = F_unlink(m, int32(_a_F_pgstat_before_server_shutdown_2))
	mBase = m.M
	goto L3
}
func F_pgstat_bgwriter_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_reset_all_cb[0]))
	v10 = v8 + int32(320)
	v12 = F_LWLockAcquire(m, v10, int32(0))
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
	v15 = v8 + int32(344)
	v17 = v8 + int32(376)
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+336))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bgwriter_reset_all_cb[1]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+368)) = l0
	F_LWLockRelease(m, v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v33
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v35
	if v24&int32(1) != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+336))
	if v24 != v39 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	return
}
func F_pgstat_btree_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int64
	_ = v126
	var v137 int32
	_ = v137
	v5 = int32(0)
	v9 = F_ReadBufferExtended(m, l1, v5, l2, v5, l3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_LockBuffer(m, v9, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v9 < int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_btree_page[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v9^int32(-1))<<(uint(int32(2))%32))))
				v31 = v23
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_btree_page[1]))
				v31 = v25 + v9<<(uint(int32(13))%32) + int32(-8192)
			}
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
			if v32 == int32(0) {
				v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v126 - int64(-8192)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+19)))
				v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+16)))
				if (v35<<(uint(int32(8))%32)-v38)&int32(_a_F_pgstat_btree_page_0) != int32(16) {
				} else {
					v44 = v31 + v38
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)))
					if v45&int32(20) != 0 {
						v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v126 - int64(-8192)
					} else {
						if v45&int32(1) == int32(0) {
						} else {
							v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+12)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
							v56 = v54 - v52
							v57 = int32(0)
							if v57 < v56 {
								v60 = v56
							} else {
								v60 = v57
							}
							v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v61 + base.I64_extend_i32_u(v60)
							if v53 != 0 {
								v67 = int32(2)
							} else {
								v67 = int32(1)
							}
							if base.Ui32(int32(25)) <= base.Ui32(v52) {
								v75 = int32(base.Ui32(v52+int32(_a_F_pgstat_btree_page_1)) >> (uint(int32(2)) % 32))
							} else {
								v75 = int32(0)
							}
							if base.Ui32(v75&int32(_a_F_pgstat_btree_page_0)) < base.Ui32(v67) {
							} else {
								v88 = v67
								for {
									v93 = v31 + int32(20) + v88<<(uint(int32(2))%32)
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
									v95 = int32(_a_F_pgstat_btree_page_2)
									if v94&v95 == v95 {
										v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v99 + int64(1)
										v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v103 + base.I64_extend_i32_u(int32(base.Ui32(v104)>>(uint(int32(17))%32)))
									} else {
										v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v110 + int64(1)
										v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v114 + base.I64_extend_i32_u(int32(base.Ui32(v115)>>(uint(int32(17))%32)))
									}
									v122 = v88 + int32(1)
									if v122 != (v75+int32(1))&int32(_a_F_pgstat_btree_page_0) {
										v88 = v122
										continue
									} else {
										break
									}
									break
								}
							}
						}
					}
				}
			}
			F__bt_relbuf(m, v9)
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pgstat_clip_activity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_clip_activity[0]))
	v6 = F_pnstrdup(m, l0, v3-int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_strlen(m, v6)
		mBase = m.M
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_clip_activity[0]))
		v15 = F_pg_mbcliplen(m, v6, v10, v12-int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15+v6))) = uint8(v18)
			return v6
		}
	}
}
func F_pgstat_count_heap_insert(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v6 == int32(0) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
		if v9 != int32(1) {
			return
		} else {
			F_pgstat_assoc_relation(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v15 = v14
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[0]))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v19 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
					if v20 == v18 {
						v37 = v19
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
						return
					} else {
						v22 = F_pgstat_get_xact_stack_level(m, v18)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[1]))
							v27 = F_MemoryContextAllocZero(m, v25, int32(72))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v18
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v30
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v33
								*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v27
								*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v27
								v37 = v27
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
								*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
								return
							}
						}
					}
				} else {
					v22 = F_pgstat_get_xact_stack_level(m, v18)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[1]))
						v27 = F_MemoryContextAllocZero(m, v25, int32(72))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v18
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v15
							*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v30
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v27
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v27
							v37 = v27
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
							return
						}
					}
				}
			}
		}
	} else {
		v15 = v6
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
			if v20 == v18 {
				v37 = v19
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
				*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
				return
			} else {
				v22 = F_pgstat_get_xact_stack_level(m, v18)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[1]))
					v27 = F_MemoryContextAllocZero(m, v25, int32(72))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v18
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v33
						*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v27
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v27
						v37 = v27
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
						return
					}
				}
			}
		} else {
			v22 = F_pgstat_get_xact_stack_level(m, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_heap_insert[1]))
				v27 = F_MemoryContextAllocZero(m, v25, int32(72))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v18
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v15
					*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v30
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v27
					v37 = v27
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
					*(*int64)(unsafe.Add(mBase, uint32(v37))) = v40 + l1
					return
				}
			}
		}
	}
}
func F_pgstat_count_io_op_time(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l3 != int64(0) {
		F___clock_gettime(m, int32(1), v12)
		mBase = m.M
		v18 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		v23 = v18 + (v19*int64(1000000000) - l3)
		if l0 == int32(2) {
		} else {
			if l2&int32(-3) == int32(5) {
				v30 = int32(_a_F_pgstat_count_io_op_time_0)
				v32 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[0]))
				v34 = base.I64_div_s(v23, int64(1000))
				*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[0])) = v32 + v34
				switch l0 {
				case 0:
					v37 = int32(_a_F_pgstat_count_io_op_time_1)
					v39 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[1]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[1])) = v39 + v23
				case 1:
					v42 = int32(_a_F_pgstat_count_io_op_time_2)
					v44 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[2]))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[2])) = v44 + v23
				default:
				}
			} else {
				if l2 != int32(6) {
				} else {
					v49 = int32(_a_F_pgstat_count_io_op_time_3)
					v51 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[3]))
					v53 = base.I64_div_s(v23, int64(1000))
					*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[3])) = v51 + v53
					switch l0 {
					case 0:
						v56 = int32(_a_F_pgstat_count_io_op_time_4)
						v58 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[4]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[4])) = v58 + v23
					case 1:
						v61 = int32(_a_F_pgstat_count_io_op_time_5)
						v63 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[5]))
						*(*int64)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[5])) = v63 + v23
					default:
					}
				}
			}
		}
		v73 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
		v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pgstat_count_io_op_time[6])))
		*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pgstat_count_io_op_time[6]))) = v74 + v23
		v78 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[7]))
		v85 = int32(0)
		if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v78))|base.B2i32(int32(1)<<(uint(v78)%32)&int32(_a_F_pgstat_count_io_op_time_6) == v85) == v85 {
			v90 = *(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pgstat_count_io_op_time[8])))
			*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pgstat_count_io_op_time[8]))) = v90 + v23
			v94 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[9])) = uint8(v94)
			*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[10])) = uint8(v94)
		} else {
		}
	} else {
	}
	v110 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_pgstat_count_io_op_time[11])))
	*(*int64)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_pgstat_count_io_op_time[11]))) = v111 + base.I64_extend_i32_u(l4)
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_pgstat_count_io_op_time[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_pgstat_count_io_op_time[12]))) = v115 + l5
	v118 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[7]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v120))|base.B2i32(int32(1)<<(uint(v120)%32)&int32(_a_F_pgstat_count_io_op_time_6) == v118) == v118 {
		v139 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
		v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pgstat_count_io_op_time[13])))
		*(*int64)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pgstat_count_io_op_time[13]))) = v140 + base.I64_extend_i32_u(l4)
		v144 = *(*int64)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pgstat_count_io_op_time[14])))
		*(*int64)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pgstat_count_io_op_time[14]))) = v144 + l5
		v148 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[9])) = uint8(v148)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[10])) = uint8(v148)
	} else {
	}
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[9])) = uint8(v155)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op_time[15])) = uint8(v155)
	m.G0 = v12 + int32(16)
	return
}
func F_pgstat_count_truncate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v5 == int32(0) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
		if v8 != int32(1) {
			return
		} else {
			F_pgstat_assoc_relation(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v14 = v13
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[0]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
					if v19 == v17 {
						v36 = v18
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
						if v39 != 0 {
							v49 = v36
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
							v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v49 = v48
						}
						v50 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
						return
					} else {
						v21 = F_pgstat_get_xact_stack_level(m, v17)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[1]))
							v26 = F_MemoryContextAllocZero(m, v24, int32(72))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
								*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
								*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
								v36 = v26
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
								if v39 != 0 {
									v49 = v36
								} else {
									v40 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
									v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
									v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
									*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
									v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
									v49 = v48
								}
								v50 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
								return
							}
						}
					}
				} else {
					v21 = F_pgstat_get_xact_stack_level(m, v17)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[1]))
						v26 = F_MemoryContextAllocZero(m, v24, int32(72))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
							*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
							v36 = v26
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
							if v39 != 0 {
								v49 = v36
							} else {
								v40 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
								*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								v49 = v48
							}
							v50 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
							return
						}
					}
				}
			}
		}
	} else {
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
			if v19 == v17 {
				v36 = v18
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
				if v39 != 0 {
					v49 = v36
				} else {
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v49 = v48
				}
				v50 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
				return
			} else {
				v21 = F_pgstat_get_xact_stack_level(m, v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[1]))
					v26 = F_MemoryContextAllocZero(m, v24, int32(72))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
						*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
						v36 = v26
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
						if v39 != 0 {
							v49 = v36
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
							v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v49 = v48
						}
						v50 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
						return
					}
				}
			}
		} else {
			v21 = F_pgstat_get_xact_stack_level(m, v17)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_truncate[1]))
				v26 = F_MemoryContextAllocZero(m, v24, int32(72))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
					v36 = v26
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
					if v39 != 0 {
						v49 = v36
					} else {
						v40 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v40)
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v42
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v44
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+40)) = v46
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						v49 = v48
					}
					v50 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v50
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v50
					return
				}
			}
		}
	}
}
func F_pgstat_database_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = F_pgstat_lock_entry(m, l0, l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v11 + v12
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v15 + v16
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v5)+40))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+40)) = v19 + v20
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v5)+48))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+48)) = v23 + v24
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v5)+56))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+56)) = v27 + v28
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v5)+64))
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+64)) = v31 + v32
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v6)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+72)) = v35 + v36
			v39 = *(*int64)(unsafe.Add(mBase, uint32(v5)+80))
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+80)) = v39 + v40
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v5)+88))
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v6)+64))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+88)) = v43 + v44
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v5)+104))
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v6)+80))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+104)) = v47 + v48
			v51 = *(*int64)(unsafe.Add(mBase, uint32(v5)+112))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)+88))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+112)) = v51 + v52
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)+120))
			v56 = *(*int64)(unsafe.Add(mBase, uint32(v6)+96))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+120)) = v55 + v56
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v5)+128))
			v60 = *(*int64)(unsafe.Add(mBase, uint32(v6)+104))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+128)) = v59 + v60
			v63 = *(*int64)(unsafe.Add(mBase, uint32(v5)+136))
			v64 = *(*int64)(unsafe.Add(mBase, uint32(v6)+112))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+136)) = v63 + v64
			v67 = *(*int64)(unsafe.Add(mBase, uint32(v5)+144))
			v68 = *(*int64)(unsafe.Add(mBase, uint32(v6)+120))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+144)) = v67 + v68
			v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)+160))
			v72 = *(*int64)(unsafe.Add(mBase, uint32(v6)+136))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+160)) = v71 + v72
			v75 = *(*int64)(unsafe.Add(mBase, uint32(v5)+152))
			v76 = *(*int64)(unsafe.Add(mBase, uint32(v6)+128))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+152)) = v75 + v76
			v79 = *(*int64)(unsafe.Add(mBase, uint32(v5)+168))
			v80 = *(*int64)(unsafe.Add(mBase, uint32(v6)+144))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+168)) = v79 + v80
			v83 = *(*int64)(unsafe.Add(mBase, uint32(v5)+192))
			v84 = *(*int64)(unsafe.Add(mBase, uint32(v6)+168))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+192)) = v83 + v84
			v87 = *(*int64)(unsafe.Add(mBase, uint32(v5)+200))
			v88 = *(*int64)(unsafe.Add(mBase, uint32(v6)+176))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+200)) = v87 + v88
			v91 = *(*int64)(unsafe.Add(mBase, uint32(v5)+208))
			v92 = *(*int64)(unsafe.Add(mBase, uint32(v6)+184))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+208)) = v91 + v92
			v95 = *(*int64)(unsafe.Add(mBase, uint32(v5)+216))
			v96 = *(*int64)(unsafe.Add(mBase, uint32(v6)+192))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+216)) = v95 + v96
			v99 = *(*int64)(unsafe.Add(mBase, uint32(v5)+224))
			v100 = *(*int64)(unsafe.Add(mBase, uint32(v6)+200))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+224)) = v99 + v100
			v103 = *(*int64)(unsafe.Add(mBase, uint32(v5)+232))
			v104 = *(*int64)(unsafe.Add(mBase, uint32(v6)+208))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+232)) = v103 + v104
			v107 = *(*int64)(unsafe.Add(mBase, uint32(v5)+240))
			v108 = *(*int64)(unsafe.Add(mBase, uint32(v6)+216))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+240)) = v107 + v108
			v111 = *(*int64)(unsafe.Add(mBase, uint32(v5)+248))
			v112 = *(*int64)(unsafe.Add(mBase, uint32(v6)+224))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+248)) = v111 + v112
			v115 = *(*int64)(unsafe.Add(mBase, uint32(v5)+256))
			v116 = *(*int64)(unsafe.Add(mBase, uint32(v6)+232))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+256)) = v115 + v116
			v119 = *(*int64)(unsafe.Add(mBase, uint32(v5)+264))
			v120 = *(*int64)(unsafe.Add(mBase, uint32(v6)+240))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+264)) = v119 + v120
			v123 = *(*int64)(unsafe.Add(mBase, uint32(v5)+272))
			v124 = *(*int64)(unsafe.Add(mBase, uint32(v6)+248))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+272)) = v123 + v124
			F_pgstat_unlock_entry(m, l0)
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return int32(0)
			} else {
				base.MemoryFill(m, v6, int32(0), int32(264))
				return v7
			}
		} else {
			return v7
		}
	}
}
func F_pgstat_drop_relation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_relation[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_relation[1]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+117)))
	if v12 != 0 {
		v13 = int32(0)
	} else {
		v13 = v10
	}
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_pgstat_drop_transactional(m, int32(2), v13, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		if v17 == int32(0) {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
			if v20 != int32(1) {
				return
			} else {
				F_pgstat_assoc_relation(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
					v26 = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					if v27 == int32(0) {
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
						if v30 != v6 {
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)) = uint8(v32)
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
							*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v34
							v36 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v36
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v38
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							v41 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v40))) = v41
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v41
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v41
						}
					}
					return
				}
			}
		} else {
			v26 = v17
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			if v27 == int32(0) {
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
				if v30 != v6 {
				} else {
					v32 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)) = uint8(v32)
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v34
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v36
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					v41 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v40))) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v41
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v41
				}
			}
			return
		}
	}
}
func F_pgstat_get_entry_ref_locked(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v5 = int32(0)
	v8 = F_pgstat_get_entry_ref(m, l0, l1, l2, int32(1), v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v14 = v12 + int32(4)
		if l3 == int32(0) {
			v18 = F_LWLockAcquire(m, v14, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v25 = v8
				return v25
			}
		} else {
			v21 = F_LWLockConditionalAcquire(m, v14, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					v25 = v5
				} else {
					v25 = v8
				}
				return v25
			}
		}
	}
}
func F_pgstat_get_wait_event_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = l0 - int32(16777216)
		if base.Ui32(int32(184549375)) < base.Ui32(v7) {
			return int32(_a_F_pgstat_get_wait_event_type_0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v7)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pgstat_get_wait_event_type[0])))
			return v16
		}
	}
}
func F_pgstat_hash_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v128 int32
	_ = v128
	v5 = int32(0)
	v9 = F_ReadBufferExtended(m, l1, v5, l2, v5, l3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_LockBuffer(m, v9, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v9 < int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_hash_page[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v9^int32(-1))<<(uint(int32(2))%32))))
				v31 = v23
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_hash_page[1]))
				v31 = v25 + v9<<(uint(int32(13))%32) + int32(-8192)
			}
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
			if v32 == int32(0) {
				v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v117 - int64(-8192)
			} else {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+19)))
				v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+16)))
				if (v35<<(uint(int32(8))%32)-v38)&int32(_a_F_pgstat_hash_page_0) != int32(16) {
				} else {
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+v38)+12)))
					switch v45 & int32(15) {
					case 0:
						v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v117 - int64(-8192)
					case 1, 2:
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+12)))
						v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+14)))
						v51 = v49 - v48
						v52 = int32(0)
						if v52 < v51 {
							v55 = v51
						} else {
							v55 = v52
						}
						v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v56 + base.I64_extend_i32_u(v55)
						if base.Ui32(v48) < base.Ui32(int32(25)) {
						} else {
							v63 = v48 + int32(_a_F_pgstat_hash_page_1)
							if v63&int32(_a_F_pgstat_hash_page_2) == int32(0) {
							} else {
								v72 = int32(1)
								v80 = v72
								for {
									v85 = v31 + int32(20) + v80<<(uint(int32(2))%32)
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
									v87 = int32(_a_F_pgstat_hash_page_3)
									if v86&v87 == v87 {
										v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v91 + int64(1)
										v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v95 + base.I64_extend_i32_u(int32(base.Ui32(v96)>>(uint(int32(17))%32)))
									} else {
										v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v102 + int64(1)
										v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v106 + base.I64_extend_i32_u(int32(base.Ui32(v107)>>(uint(int32(17))%32)))
									}
									v114 = v80 + int32(1)
									if v114 != (int32(base.Ui32(v63)>>(uint(int32(2))%32))+v72)&int32(_a_F_pgstat_hash_page_0) {
										v80 = v114
										continue
									} else {
										break
									}
									break
								}
							}
						}
					default:
					}
				}
			}
			F_UnlockReleaseBuffer(m, v9)
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pgstat_lock_entry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = v3 + int32(4)
	if l1 != 0 {
		v7 = F_LWLockConditionalAcquire(m, v5, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v13 = F_LWLockAcquire(m, v5, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_pgstat_prep_pending_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[0]))
	if v6 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[1]))
		v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_pgstat_prep_pending_entry_0), int32(0), int32(1024), int32(_a_F_pgstat_prep_pending_entry_1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[0])) = v16
			v22 = F_pgstat_get_entry_ref(m, l0, l1, l2, int32(1), l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				if v24 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[0]))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v45 = l0*int32(72) + int32(_a_F_pgstat_prep_pending_entry_2)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[2]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+l0<<(uint(int32(2))%32)-int32(96))))
						v45 = v44
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
					v47 = F_MemoryContextAllocZero(m, v28, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v47
						v51 = v22 + int32(16)
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[3]))
						if v53 != 0 {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[4]))
							v60 = v55
						} else {
							v57 = int32(_a_F_pgstat_prep_pending_entry_3)
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[3])) = v57
							v60 = v57
						}
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v60
						v62 = int32(_a_F_pgstat_prep_pending_entry_3)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[4])) = v51
						return v22
					}
				} else {
					return v22
				}
			}
		}
	} else {
		v22 = F_pgstat_get_entry_ref(m, l0, l1, l2, int32(1), l3)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
			if v24 == int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[0]))
				if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
					v45 = l0*int32(72) + int32(_a_F_pgstat_prep_pending_entry_2)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[2]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+l0<<(uint(int32(2))%32)-int32(96))))
					v45 = v44
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
				v47 = F_MemoryContextAllocZero(m, v28, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v47
					v51 = v22 + int32(16)
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[3]))
					if v53 != 0 {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[4]))
						v60 = v55
					} else {
						v57 = int32(_a_F_pgstat_prep_pending_entry_3)
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[3])) = v57
						v60 = v57
					}
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v60
					v62 = int32(_a_F_pgstat_prep_pending_entry_3)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_pending_entry[4])) = v51
					return v22
				}
			} else {
				return v22
			}
		}
	}
}
func F_pgstat_prep_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[0])))
	if v4 != 0 {
		v6 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[1])) = v6
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[2])) = v6
		*(*int64)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[3])) = v6
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[4])) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v15
		*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[6])) = v15
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7]))
		if v24 != 0 {
			F_MemoryContextDelete(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v33 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[0])) = uint8(v33)
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[8]))
					if v37 == int32(0) {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5]))
						if v41 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7]))
							if v43 == int32(0) {
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[9]))
								v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_pgstat_prep_snapshot_0), int32(0), int32(1024), int32(_a_F_pgstat_prep_snapshot_1))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7])) = v53
									v56 = v53
									v58 = F_MemoryContextAllocZero(m, v56, int32(32))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
										v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
											*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
											*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
											*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
											return
										}
									}
								}
							} else {
								v56 = v43
								v58 = F_MemoryContextAllocZero(m, v56, int32(32))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
									v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
										*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
										*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[0])) = uint8(v33)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[8]))
				if v37 == int32(0) {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5]))
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7]))
						if v43 == int32(0) {
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[9]))
							v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_pgstat_prep_snapshot_0), int32(0), int32(1024), int32(_a_F_pgstat_prep_snapshot_1))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7])) = v53
								v56 = v53
								v58 = F_MemoryContextAllocZero(m, v56, int32(32))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
									v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
										*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
										*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
										return
									}
								}
							}
						} else {
							v56 = v43
							v58 = F_MemoryContextAllocZero(m, v56, int32(32))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
								v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
									*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
									*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[8]))
		if v37 == int32(0) {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5]))
			if v41 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7]))
				if v43 == int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[9]))
					v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_pgstat_prep_snapshot_0), int32(0), int32(1024), int32(_a_F_pgstat_prep_snapshot_1))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[7])) = v53
						v56 = v53
						v58 = F_MemoryContextAllocZero(m, v56, int32(32))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
							v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
								*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
								*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
								return
							}
						}
					}
				} else {
					v56 = v43
					v58 = F_MemoryContextAllocZero(m, v56, int32(32))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
						v65 = F_MemoryContextAllocExtended(m, v56, int32(_a_F_pgstat_prep_snapshot_2), int32(5))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
							*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
							*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
							*(*int32)(unsafe.Add(mBase, _c_F_pgstat_prep_snapshot[5])) = v58
							return
						}
					}
				}
			}
		}
	}
}
func F_pgstat_prepare_io_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		F___clock_gettime(m, int32(1), v6)
		mBase = m.M
		v10 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v15 = v10 + v11*int64(1000000000)
	} else {
		v15 = int64(0)
	}
	m.G0 = v6 + int32(16)
	return v15
}
func F_pgstat_progress_parallel_incr_param(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[0]))
	if int32(0) <= v5 {
		F_initStringInfo(m, int32(_a_F_pgstat_progress_parallel_incr_param_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_pq_beginmessage(m, int32(_a_F_pgstat_progress_parallel_incr_param_0), int32(80))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_enlargeStringInfo(m, int32(_a_F_pgstat_progress_parallel_incr_param_0), int32(4))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = int32(_a_F_pgstat_progress_parallel_incr_param_0)
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[1]))
					v21 = int32(_a_F_pgstat_progress_parallel_incr_param_1)
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2]))
					v26 = int32(16711935)
					v30 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v20+v22))) = base.I32_rotr(l0, int32(24))&v26 | base.I32_rotr(l0&v26, v30)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2])) = v36 + int32(4)
					F_enlargeStringInfo(m, v19, v30)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = int32(_a_F_pgstat_progress_parallel_incr_param_0)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[1]))
						v46 = int32(_a_F_pgstat_progress_parallel_incr_param_1)
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2]))
						v49 = int64(56)
						v51 = int64(65280)
						v53 = int64(40)
						v56 = int64(16711680)
						v58 = int64(24)
						v60 = int64(4278190080)
						v62 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v45+v47))) = l1<<(uint(v49)%64) | l1&v51<<(uint(v53)%64) | (l1&v56<<(uint(v58)%64) | l1&v60<<(uint(v62)%64)) | (int64(base.Ui64(l1)>>(uint(v62)%64))&v60 | int64(base.Ui64(l1)>>(uint(v58)%64))&v56 | (int64(base.Ui64(l1)>>(uint(v53)%64))&v51 | int64(base.Ui64(l1)>>(uint(v49)%64))))
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[2])) = v87 + int32(8)
						F_pq_endmessage(m, v44)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v95 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[3]))
		if v95 == int32(0) {
		} else {
			v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[4])))
			if v99&int32(1) == int32(0) {
			} else {
				v104 = int32(_a_F_pgstat_progress_parallel_incr_param_2)
				v106 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[5]))
				v107 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[5])) = v106 + v107
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = v110 + v107
				v118 = v95 + l0<<(uint(int32(3))%32) + int32(232)
				v119 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
				*(*int64)(unsafe.Add(mBase, uint32(v118))) = v119 + l1
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = v122 + v107
				v128 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[5]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_parallel_incr_param[5])) = v128 - v107
			}
		}
		return
	}
}
func F_pgstat_report_activity(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[0]))
	if v14 == int32(0) {
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_activity[1])))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
			if v21 == int32(7) {
			} else {
				v24 = int32(_a_F_pgstat_report_activity_0)
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2]))
				v27 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2])) = v26 + v27
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[3]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v32 + v27
				*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = int32(7)
				v38 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+216))
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v14)+392)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v14)+400)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v41
				v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v173 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v172 + v173
				v176 = int32(_a_F_pgstat_report_activity_0)
				v178 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2])) = v178 - v173
			}
		} else {
			v54 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_report_activity[4]))
			if l1 != 0 {
				v55 = F_strlen(m, l1)
				mBase = m.M
				v57 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[5]))
				v59 = v57 - int32(1)
				if base.Ui32(v55) < base.Ui32(v59) {
					v61 = v55
				} else {
					v61 = v59
				}
				v65 = v61
			} else {
				v65 = int32(0)
			}
			v69 = m.G0
			v70 = int32(16)
			v71 = v69 - v70
			m.G0 = v71
			F_gettimeofday(m, v71)
			mBase = m.M
			v74 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
			v75 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+8)))
			m.G0 = v71 + v70
			v83 = v75 + v74*int64(1000000) - int64(946684800000000)
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
			if v84 == int32(3) {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
				if l0 == v96 {
				} else {
					v98 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
					v105 = v83 - v98
					if v105 <= int64(0) {
						v117 = int32(0)
						v118 = int32(0)
					} else {
						v109 = int64(1000000)
						v110 = base.I64_div_u_s(v105, v109)
						v117 = base.I32_wrap_i64(v110)
						v118 = base.I32_wrap_i64(v105 - v110*v109)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v117
					*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v118
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
					if v122 != int32(3) {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
						if v127 == int32(5) {
							v130 = int32(_a_F_pgstat_report_activity_1)
						} else {
							v130 = int32(_a_F_pgstat_report_activity_2)
						}
						v131 = v130
					} else {
						v131 = int32(_a_F_pgstat_report_activity_1)
					}
					v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
					v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
					v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+12)))
					*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132 + (v133 + v134*int64(1000000))
				}
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
				if v87 == int32(5) {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
					if l0 == v96 {
					} else {
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
						v105 = v83 - v98
						if v105 <= int64(0) {
							v117 = int32(0)
							v118 = int32(0)
						} else {
							v109 = int64(1000000)
							v110 = base.I64_div_u_s(v105, v109)
							v117 = base.I32_wrap_i64(v110)
							v118 = base.I32_wrap_i64(v105 - v110*v109)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v117
						*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v118
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
						if v122 != int32(3) {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
							if v127 == int32(5) {
								v130 = int32(_a_F_pgstat_report_activity_1)
							} else {
								v130 = int32(_a_F_pgstat_report_activity_2)
							}
							v131 = v130
						} else {
							v131 = int32(_a_F_pgstat_report_activity_1)
						}
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
						v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
						v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+12)))
						*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132 + (v133 + v134*int64(1000000))
					}
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
					if v90 == int32(4) {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
						if l0 == v96 {
						} else {
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							v105 = v83 - v98
							if v105 <= int64(0) {
								v117 = int32(0)
								v118 = int32(0)
							} else {
								v109 = int64(1000000)
								v110 = base.I64_div_u_s(v105, v109)
								v117 = base.I32_wrap_i64(v110)
								v118 = base.I32_wrap_i64(v105 - v110*v109)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v117
							*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v118
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
							if v122 != int32(3) {
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
								if v127 == int32(5) {
									v130 = int32(_a_F_pgstat_report_activity_1)
								} else {
									v130 = int32(_a_F_pgstat_report_activity_2)
								}
								v131 = v130
							} else {
								v131 = int32(_a_F_pgstat_report_activity_1)
							}
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
							v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
							v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+12)))
							*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132 + (v133 + v134*int64(1000000))
						}
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
						if v93 != int32(6) {
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
							if l0 == v96 {
							} else {
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								v105 = v83 - v98
								if v105 <= int64(0) {
									v117 = int32(0)
									v118 = int32(0)
								} else {
									v109 = int64(1000000)
									v110 = base.I64_div_u_s(v105, v109)
									v117 = base.I32_wrap_i64(v110)
									v118 = base.I32_wrap_i64(v105 - v110*v109)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v117
								*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v118
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
								if v122 != int32(3) {
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
									if v127 == int32(5) {
										v130 = int32(_a_F_pgstat_report_activity_1)
									} else {
										v130 = int32(_a_F_pgstat_report_activity_2)
									}
									v131 = v130
								} else {
									v131 = int32(_a_F_pgstat_report_activity_1)
								}
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
								v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
								v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+12)))
								*(*int64)(unsafe.Add(mBase, uint32(v131))) = v132 + (v133 + v134*int64(1000000))
							}
						}
					}
				}
			}
			v141 = int32(_a_F_pgstat_report_activity_0)
			v143 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2]))
			v144 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2])) = v143 + v144
			v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v147 + v144
			*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v83
			if l0 == int32(3) {
				v155 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+392)) = v155
				*(*int64)(unsafe.Add(mBase, uint32(v14)+400)) = v155
			} else {
			}
			if l1 != 0 {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+216))
				if v65 != 0 {
					base.MemoryCopy(m, v159, l1, v65)
				} else {
				}
				v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+216))
				v163 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v161+v65))) = uint8(v163)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v54
			} else {
			}
			v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v173 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v172 + v173
			v176 = int32(_a_F_pgstat_report_activity_0)
			v178 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_activity[2])) = v178 - v173
		}
	}
	m.G0 = v11 + int32(16)
	return
}
func F_pgstat_reset_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v5 = int32(0)
	v10 = F_pgstat_get_entry_ref(m, l0, l1, l2, v5, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 == int32(0) {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v20 = F_LWLockAcquire(m, v16+int32(4), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v51 = l0*int32(72) + int32(_a_F_pgstat_reset_entry_0)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v49 = int32(0)
						} else {
							v37 = int32(0)
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_entry[0]))
							if v39 == v37 {
								v49 = v37
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v39+l0<<(uint(int32(2))%32)-int32(96))))
								v49 = v47
							}
						}
						v51 = v49
					}
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v80 = l0*int32(72) + int32(_a_F_pgstat_reset_entry_0)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v78 = int32(0)
						} else {
							v66 = int32(0)
							v68 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_entry[0]))
							if v68 == v66 {
								v78 = v66
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v68+l0<<(uint(int32(2))%32)-int32(96))))
								v78 = v76
							}
						}
						v80 = v78
					}
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v110 = l0*int32(72) + int32(_a_F_pgstat_reset_entry_0)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v108 = int32(0)
						} else {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_entry[0]))
							if v98 == v96 {
								v108 = v96
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v98+l0<<(uint(int32(2))%32)-int32(96))))
								v108 = v106
							}
						}
						v110 = v108
					}
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
					if v111 != 0 {
						base.MemoryFill(m, v22+v81, int32(0), v111)
					} else {
					}
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
					if v115 != 0 {
						m.T0[v115].(func(*base.Module, int32, int64))(m, v22, l3)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							F_LWLockRelease(m, v118+int32(4))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						F_LWLockRelease(m, v118+int32(4))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_pgstat_shutdown_hook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v35 int64
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v241 int32
	_ = v241
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[0]))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[1]))
	if v15 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v44 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	goto L3
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[0]))
	v23 = F_pgstat_prep_pending_entry(m, int32(1), v20, int64(0), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[2]))
	v27 = int32(2)
	v28 = v26 - v27
	if base.Ui32(v27) < base.Ui32(v28) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v34 = v31 + v28<<(uint(int32(3))%32)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+216)) = v35 + int64(1)
	goto L4
L9:
	;
	v47 = int32(_a_F_pgstat_shutdown_hook_0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[3])) = v47
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[4])) = v47
	v55 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[5])))
	v56 = F_pgstat_drop_entry(m, int32(6), int32(0), v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v56 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[6]))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+16)) = v63 + int64(1)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v67 = m.G0
	v69 = v67 - int32(16)
	m.G0 = v69
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[7]))
	if v72 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	if v74 == int64(0) {
		v102 = int32(-1)
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[8]))
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L34
	}
L18:
	;
	v112 = v102
	v114 = v72
	v118 = int32(0)
	goto L24
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v80 = int32(0)
	goto L20
L20:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v80*int32(24))+16)))
	if v93 != int32(1) {
		v102 = v80
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v102 = int32(-1)
	goto L18
L22:
	;
	v97 = v80 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v97)) < base.Ui64(v74) {
		v80 = v97
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v124 = v112
	v127 = v118
	v130 = v118
	goto L27
L25:
	;
	F_pfree(m, v123)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	if v127&int32(1) != 0 {
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v155
	F_pgstat_release_entry_ref(m, v69, v152, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v138 = int32(1)
	v139 = v124 - v138
	v143 = base.B2i32(v137&(v139^v102) == int32(0))
	v144 = v143 | v130
	v147 = v137 & v139
	v148 = v123 + v124*int32(24)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+16)))
	if v149 != v138 {
		v124 = v147
		v127 = v143
		v130 = v144
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[7]))
	v112 = v147
	v114 = v161
	v118 = v144
	goto L24
L32:
	;
	F_pfree(m, v114)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[7])) = int32(0)
	goto L17
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[8])) = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[9]))
	F_dsa_detach(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v191 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[6]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v196 = v194 + int32(1476)
	v198 = F_LWLockAcquire(m, v196, v191)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)+1460))
	v202 = v200 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+1460)) = v202
	if v202 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v209 = v191
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_LWLockRelease(m, v196)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L47
	}
L40:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v194+int32(32)+v209<<(uint(int32(2))%32))))
	if v222 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	F_dsm_unpin_segment(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v226 = v209 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v194)+1456))
	if base.Ui32(v226) <= base.Ui32(v227) {
		v209 = v226
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_shutdown_hook[9])) = int32(0)
	m.G0 = v69 + int32(16)
	return
}
func F_pgstat_wal_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_wal_reset_all_cb[0]))
	v7 = v5 + int32(_a_F_pgstat_wal_reset_all_cb_0)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_wal_reset_all_cb[1]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_wal_reset_all_cb[2]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_wal_reset_all_cb[3]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_wal_reset_all_cb[4]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_wal_reset_all_cb[5]))) = l0
		F_LWLockRelease(m, v7)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
