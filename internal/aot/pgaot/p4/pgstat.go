package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_before_server_shutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v15 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
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
	m.G0 = v12 + int32(176)
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v19)
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1219])) = v22
	v26 = F_errstart(m, int32(13), v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(110836)
	F_errmsg_internal(m, int32(694137), v12-int32(-64))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v42 = F_AllocateFile(m, int32(233153), int32(32371))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	F_errfinish(m, int32(488059), int32(1587), int32(382000))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+172)) = int32(27638967)
	v70 = F_fwrite(m, v12+int32(172), int32(4), int32(1), v42)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	if v48 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(233153)
	F_errmsg(m, int32(294797), v12)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(488059), int32(1598), int32(382000))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v72
	v75 = v72
	goto L21
L21:
	;
	v87 = base.B2i32(base.Ui32(int32(11)) < base.Ui32(v75-int32(1)))
	if base.Ui32(int32(11)) < base.Ui32(v75-int32(1)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v196 = v12 + int32(144)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[1220]))
	v199 = int32(0)
	v200 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v196)+4)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v198
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+24)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v196)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v196)+12)) = v200
	goto L53
L23:
	;
	v191 = v185 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v191
	if base.Ui32(v191) < base.Ui32(int32(33)) {
		v75 = v191
		goto L21
	} else {
		goto L52
	}
L24:
	;
	if base.Ui32(int32(8)) < base.Ui32(v75-int32(24)) {
		v185 = v75
		goto L23
	} else {
		goto L27
	}
L25:
	;
	v107 = v75*int32(72) + int32(1630560)
	goto L26
L26:
	;
	if v107 == int32(0) {
		v185 = v75
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	if v93 == int32(0) {
		v185 = v75
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93+v75<<(uint(int32(2))%32)-int32(96))))
	v107 = v101
	goto L26
L29:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110&int32(4) == int32(0) {
		v185 = v75
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v110&int32(1) == int32(0) {
		v185 = v75
		goto L23
	} else {
		goto L31
	}
L31:
	;
	if v87 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v147 = v143 + v144
	v149 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
	if v149 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v143 = v75
	v144 = int32(4400064)
	v146 = v75*int32(72) + int32(1630560)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v126 = int32(4452736)
	v129 = v75 - int32(24)
	if base.Ui32(int32(8)) < base.Ui32(v129) {
		v143 = v129
		v144 = v126
		v146 = int32(0)
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v132 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	if v134 == v132 {
		v143 = v129
		v144 = v126
		v146 = v132
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134+v75<<(uint(int32(2))%32)-int32(96))))
	v143 = v129
	v144 = v126
	v146 = v142
	goto L32
L38:
	;
	if v87 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
	m.T0[v155].(func(*base.Module))(m)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L44
	}
L40:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v152)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v154 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v158)
	goto L38
L45:
	;
	F_fputc(m, int32(70), v42)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v170 = v162 + int32(4400048)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_consts[1221])))
	v170 = v169
	goto L45
L49:
	;
	v178 = F_fwrite(m, v12+int32(80), int32(4), int32(1), v42)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v182 = F_fwrite(m, v170, v180, int32(1), v42)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	v185 = v184
	goto L23
L52:
	;
	goto L22
L53:
	;
	v210 = F_dshash_seq_next(m, v12+int32(144))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v210 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v212 = v210
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_dshash_seq_term(m, v12+int32(144))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L96
	}
L58:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v222 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+16)))
	if v225 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v346 = F_dshash_seq_next(m, v12+int32(144))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L94
	}
L65:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if base.Ui32(v226-int32(1)) < base.Ui32(int32(12)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[1222]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v212)+28))
	v258 = F_dsa_get_address(m, v256, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L73
	}
L67:
	;
	if base.Ui32(v226-int32(24)) < base.Ui32(int32(9)) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v237 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v237 == int32(0) {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v212)))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v212)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v241
	F_errmsg_internal(m, int32(37600), v12+int32(48))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(488059), int32(1667), int32(382000))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L64
L73:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if base.Ui32(v260-int32(1)) <= base.Ui32(int32(11)) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v278&int32(4) == int32(0) {
		goto L64
	} else {
		goto L78
	}
L75:
	;
	v277 = v260*int32(72) + int32(1630560)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270+v260<<(uint(int32(2))%32)-int32(96))))
	v277 = v276
	goto L74
L78:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v277)+44))
	if v283 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if base.Ui32(v310-int32(1)) <= base.Ui32(int32(11)) {
		goto L90
	} else {
		goto L91
	}
L80:
	;
	F_fputc(m, int32(83), v42)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	m.T0[v283].(func(*base.Module, int32, int32, int32))(m, v212, v258, v12+int32(80))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v291 = F_fwrite(m, v212, int32(16), int32(1), v42)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L79
L85:
	;
	F_fputc(m, int32(78), v42)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v302 = F_fwrite(m, v212, int32(4), int32(1), v42)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v308 = F_fwrite(m, v12+int32(80), int32(64), int32(1), v42)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L79
L89:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v331)+20))
	v338 = F_fwrite(m, v258+v334, v336, int32(1), v42)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L93
	}
L90:
	;
	v316 = v310 * int32(72)
	v331 = v316 + int32(1630560)
	v333 = v316 + int32(1630576)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322+v310<<(uint(int32(2))%32)-int32(96))))
	v331 = v328
	v333 = v328 + int32(16)
	goto L89
L93:
	;
	goto L64
L94:
	;
	if v346 != 0 {
		v212 = v346
		goto L58
	} else {
		goto L95
	}
L95:
	;
	goto L59
L96:
	;
	F_fputc(m, int32(69), v42)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	if v364 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	if int32(base.Ui32(v369)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	goto L98
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v369 = v367
	goto L99
L101:
	;
	goto L102
L102:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v369 = v368
	goto L99
L103:
	;
	v376 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v396 = F_FreeFile(m, v42)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L114
	}
L106:
	;
	if v376 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v392 = F_FreeFile(m, v42)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L113
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(233153)
	F_errmsg(m, int32(294847), v12+int32(32))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(488059), int32(1719), int32(382000))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v395 = F_unlink(m, int32(233153))
	mBase = m.M
	goto L3
L114:
	;
	if v396 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v402 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v423 = F_durable_rename(m, int32(233153), int32(110836), int32(15))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L125
	}
L118:
	;
	if v402 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v419 = F_unlink(m, int32(233153))
	mBase = m.M
	goto L3
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(233153)
	F_errmsg(m, int32(294898), v12+int32(16))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(488059), int32(1728), int32(382000))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	if int32(0) <= v423 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v428 = F_unlink(m, int32(233153))
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[746]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v33
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v35
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
func F_pgstat_clip_activity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v3 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v6 = F_pnstrdup(m, l0, v3-int32(1))
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
	if v6&int32(3) == int32(0) {
		v33 = v6
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v71 = F_pg_mbcliplen(m, v6, v66, v68-int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v66 = v58 - v6
	goto L3
L5:
	;
	v37 = v33
	goto L14
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v66 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v22 = v6
	goto L10
L10:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v58 = v26
	goto L4
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v52 = v37
	goto L17
L16:
	;
	goto L15
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v58 = v52
	goto L4
L19:
	;
	goto L18
L20:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v6))) = uint8(v74)
	return v6
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
				v17 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
							v25 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
						v25 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
		v17 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
					v25 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
				v25 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v127 int64
	_ = v127
	var v132 int32
	_ = v132
	var v148 int32
	_ = v148
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
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
				v30 = int32(4455872)
				v32 = *(*int64)(unsafe.Add(mBase, _consts[1235]))
				v34 = base.I64_div_s(v23, int64(1000))
				*(*int64)(unsafe.Add(mBase, _consts[1235])) = v32 + v34
				switch l0 {
				case 0:
					v37 = int32(4374496)
					v39 = *(*int64)(unsafe.Add(mBase, _consts[16]))
					*(*int64)(unsafe.Add(mBase, _consts[16])) = v39 + v23
				case 1:
					v42 = int32(4374512)
					v44 = *(*int64)(unsafe.Add(mBase, _consts[18]))
					*(*int64)(unsafe.Add(mBase, _consts[18])) = v44 + v23
				default:
				}
			} else {
				if l2 != int32(6) {
				} else {
					v49 = int32(4455864)
					v51 = *(*int64)(unsafe.Add(mBase, _consts[1236]))
					v53 = base.I64_div_s(v23, int64(1000))
					*(*int64)(unsafe.Add(mBase, _consts[1236])) = v51 + v53
					switch l0 {
					case 0:
						v56 = int32(4374488)
						v58 = *(*int64)(unsafe.Add(mBase, _consts[15]))
						*(*int64)(unsafe.Add(mBase, _consts[15])) = v58 + v23
					case 1:
						v61 = int32(4374504)
						v63 = *(*int64)(unsafe.Add(mBase, _consts[17]))
						*(*int64)(unsafe.Add(mBase, _consts[17])) = v63 + v23
					default:
					}
				}
			}
		}
		v73 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
		v76 = *(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[1237])))
		*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[1237]))) = v76 + v23
		v80 = *(*int32)(unsafe.Add(mBase, _consts[264]))
		if base.Ui32(int32(16)) < base.Ui32(v80) {
		} else {
			if int32(1)<<(uint(v80)%32)&int32(115186) == int32(0) {
			} else {
				v96 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v96)+uint32(_consts[1238])))
				*(*int64)(unsafe.Add(mBase, uint32(v96)+uint32(_consts[1238]))) = v99 + v23
				v103 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v103)
				*(*uint8)(unsafe.Add(mBase, _consts[877])) = uint8(v103)
			}
		}
	} else {
	}
	v118 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v118)+uint32(_consts[911])))
	*(*int64)(unsafe.Add(mBase, uint32(v118)+uint32(_consts[911]))) = v121 + base.I64_extend_i32_u(l4)
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v118)+uint32(_consts[912])))
	*(*int64)(unsafe.Add(mBase, uint32(v118)+uint32(_consts[912]))) = v127 + l5
	v132 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	if base.Ui32(int32(16)) < base.Ui32(v132) {
	} else {
		if int32(1)<<(uint(v132)%32)&int32(115186) == int32(0) {
		} else {
			v148 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
			v151 = *(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_consts[1239])))
			*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_consts[1239]))) = v151 + base.I64_extend_i32_u(l4)
			v157 = *(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_consts[1240])))
			*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_consts[1240]))) = v157 + l5
			v161 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v161)
			*(*uint8)(unsafe.Add(mBase, _consts[877])) = uint8(v161)
		}
	}
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v169)
	*(*uint8)(unsafe.Add(mBase, _consts[874])) = uint8(v169)
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
				v16 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
							v24 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
						v24 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
		v16 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
					v24 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
				v24 = *(*int32)(unsafe.Add(mBase, _consts[144]))
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
	var v132 int32
	_ = v132
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
				v132 = F__emscripten_memset_bulkmem(m, v6, base.I32_extend8_s(int32(0)), int32(264))
				mBase = m.M
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = l0 - int32(16777216)
		if base.Ui32(int32(184549375)) < base.Ui32(v7) {
			return int32(539709)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v7)>>(uint(int32(22))%32))&int32(1020))+uint32(_consts[1246])))
			return v18
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1234]))
	if v6 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[36]))
		v16 = F_AllocSetContextCreateInternal(m, v11, int32(333089), int32(0), int32(1024), int32(8192))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1234])) = v16
			v22 = F_pgstat_get_entry_ref(m, l0, l1, l2, int32(1), l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				if v24 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[1234]))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v45 = l0*int32(72) + int32(1630560)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
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
						v53 = *(*int32)(unsafe.Add(mBase, _consts[1223]))
						if v53 != 0 {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[1224]))
							v60 = v55
						} else {
							v57 = int32(4094672)
							*(*int32)(unsafe.Add(mBase, _consts[1223])) = v57
							v60 = v57
						}
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v60
						v62 = int32(4094672)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
						*(*int32)(unsafe.Add(mBase, _consts[1224])) = v51
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
				v28 = *(*int32)(unsafe.Add(mBase, _consts[1234]))
				if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
					v45 = l0*int32(72) + int32(1630560)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
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
					v53 = *(*int32)(unsafe.Add(mBase, _consts[1223]))
					if v53 != 0 {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[1224]))
						v60 = v55
					} else {
						v57 = int32(4094672)
						*(*int32)(unsafe.Add(mBase, _consts[1223])) = v57
						v60 = v57
					}
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v60
					v62 = int32(4094672)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
					*(*int32)(unsafe.Add(mBase, _consts[1224])) = v51
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1226])))
	if v4 != 0 {
		v6 = int64(0)
		*(*int64)(unsafe.Add(mBase, _consts[1227])) = v6
		*(*int64)(unsafe.Add(mBase, _consts[1228])) = v6
		*(*int64)(unsafe.Add(mBase, _consts[1229])) = v6
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[1230])) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, _consts[1231])) = v15
		*(*int32)(unsafe.Add(mBase, _consts[1232])) = v15
		v24 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
		if v24 != 0 {
			F_MemoryContextDelete(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1233])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v33 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[1226])) = uint8(v33)
					v37 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
					if v37 == int32(0) {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
						if v41 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
							if v43 == int32(0) {
								v48 = *(*int32)(unsafe.Add(mBase, _consts[36]))
								v53 = F_AllocSetContextCreateInternal(m, v48, int32(86617), int32(0), int32(1024), int32(8192))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[1233])) = v53
									v56 = v53
									v58 = F_MemoryContextAllocZero(m, v56, int32(32))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
										v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
											*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
											*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
											*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
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
									v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
										*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
										*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
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
				*(*uint8)(unsafe.Add(mBase, _consts[1226])) = uint8(v33)
				v37 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
				if v37 == int32(0) {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
						if v43 == int32(0) {
							v48 = *(*int32)(unsafe.Add(mBase, _consts[36]))
							v53 = F_AllocSetContextCreateInternal(m, v48, int32(86617), int32(0), int32(1024), int32(8192))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[1233])) = v53
								v56 = v53
								v58 = F_MemoryContextAllocZero(m, v56, int32(32))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
									v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
										*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
										*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
										*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
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
								v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
									*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
									*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
		if v37 == int32(0) {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
			if v41 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
				if v43 == int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, _consts[36]))
					v53 = F_AllocSetContextCreateInternal(m, v48, int32(86617), int32(0), int32(1024), int32(8192))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1233])) = v53
						v56 = v53
						v58 = F_MemoryContextAllocZero(m, v56, int32(32))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v56
							v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
								*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
								*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
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
						v65 = F_MemoryContextAllocExtended(m, v56, int32(24576), int32(5))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v58)+12)) = int64(3955664880639)
							*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(1024)
							*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v65
							*(*int32)(unsafe.Add(mBase, _consts[1231])) = v58
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	v5 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(0) <= v5 {
		F_initStringInfo(m, int32(4399972))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_pq_beginmessage(m, int32(4399972), int32(80))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_enlargeStringInfo(m, int32(4399972), int32(4))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = int32(4399972)
					v20 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
					v21 = int32(4399976)
					v22 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
					v24 = int32(24)
					v26 = int32(65280)
					v28 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v20+v22))) = l0<<(uint(v24)%32) | l0&v26<<(uint(v28)%32) | (int32(base.Ui32(l0)>>(uint(v28)%32))&v26 | int32(base.Ui32(l0)>>(uint(v24)%32)))
					v42 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
					*(*int32)(unsafe.Add(mBase, _consts[1218])) = v42 + int32(4)
					F_enlargeStringInfo(m, v19, v28)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = int32(4399972)
						v51 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
						v52 = int32(4399976)
						v53 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
						v55 = int64(56)
						v57 = int64(65280)
						v59 = int64(40)
						v62 = int64(16711680)
						v64 = int64(24)
						v66 = int64(4278190080)
						v68 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v51+v53))) = l1<<(uint(v55)%64) | l1&v57<<(uint(v59)%64) | (l1&v62<<(uint(v64)%64) | l1&v66<<(uint(v68)%64)) | (int64(base.Ui64(l1)>>(uint(v68)%64))&v66 | int64(base.Ui64(l1)>>(uint(v64)%64))&v62 | (int64(base.Ui64(l1)>>(uint(v59)%64))&v57 | int64(base.Ui64(l1)>>(uint(v55)%64))))
						v93 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
						*(*int32)(unsafe.Add(mBase, _consts[1218])) = v93 + int32(8)
						F_pq_endmessage(m, v50)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, _consts[55]))
		if v101 == int32(0) {
		} else {
			v105 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
			if v105 != int32(1) {
			} else {
				v108 = int32(4470804)
				v110 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				v111 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[26])) = v110 + v111
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
				*(*int32)(unsafe.Add(mBase, uint32(v101))) = v114 + v111
				v122 = v101 + l0<<(uint(int32(3))%32) + int32(232)
				v123 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
				*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123 + l1
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
				*(*int32)(unsafe.Add(mBase, uint32(v101))) = v126 + v111
				v132 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				*(*int32)(unsafe.Add(mBase, _consts[26])) = v132 - v111
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v211 int64
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v229 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v228 + v229
	v232 = int32(4470804)
	v234 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v234 - v229
	goto L1
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v21 == int32(7) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v54 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v24 = int32(4470804)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v26 + v27
	v31 = *(*int32)(unsafe.Add(mBase, _consts[120]))
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
	goto L3
L8:
	;
	if l1&int32(3) == int32(0) {
		v78 = l1
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v121 = int32(0)
	goto L10
L10:
	;
	v125 = m.G0
	v126 = int32(16)
	v127 = v125 - v126
	m.G0 = v127
	F___gettimeofday(m, v127)
	mBase = m.M
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
	v131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v127)+8)))
	m.G0 = v127 + v126
	v139 = v131 + v130*int64(1000000) - int64(946684800000000)
	goto L31
L11:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v115 = v113 - int32(1)
	if base.Ui32(v111) < base.Ui32(v115) {
		goto L28
	} else {
		goto L29
	}
L12:
	;
	v111 = v103 - l1
	goto L11
L13:
	;
	v82 = v78
	goto L22
L14:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v62 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v111 = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v67 = l1
	goto L18
L18:
	;
	v71 = v67 + int32(1)
	if v71&int32(3) == int32(0) {
		v78 = v71
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v103 = v71
	goto L12
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v76 != 0 {
		v67 = v71
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 == v91 {
		v82 = v82 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v97 = v82
	goto L25
L24:
	;
	goto L23
L25:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v101 != 0 {
		v97 = v97 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v103 = v97
	goto L12
L27:
	;
	goto L26
L28:
	;
	v117 = v111
	goto L30
L29:
	;
	v117 = v115
	goto L30
L30:
	;
	v121 = v117
	goto L10
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v140 == int32(3) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v197 = int32(4470804)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v200 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v199 + v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v203 + v200
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v139
	if l0 == int32(3) {
		goto L50
	} else {
		goto L51
	}
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if l0 == v152 {
		goto L32
	} else {
		goto L38
	}
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v143 == int32(5) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v146 == int32(4) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v149 != int32(6) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	v161 = v139 - v154
	if v161 <= int64(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v178 != int32(3) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v174
	goto L39
L41:
	;
	v173 = int32(0)
	v174 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v165 = int64(1000000)
	v166 = base.I64_div_u_s(v161, v165)
	v173 = base.I32_wrap_i64(v166)
	v174 = base.I32_wrap_i64(v161 - v166*v165)
	goto L40
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	if v183 == int32(5) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v187 = int32(4455880)
	goto L46
L46:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v187)))
	v189 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = v188 + (v189 + v190*int64(1000000))
	goto L32
L47:
	;
	v186 = int32(4455880)
	goto L49
L48:
	;
	v186 = int32(4455888)
	goto L49
L49:
	;
	v187 = v186
	goto L46
L50:
	;
	v211 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+392)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v14)+400)) = v211
	goto L52
L51:
	;
	goto L52
L52:
	;
	if l1 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v14)+216))
	if v121 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	goto L3
L56:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v14)+216))
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218+v121))) = uint8(v220)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v54
	goto L55
L57:
	;
	v216 = F__emscripten_memcpy_bulkmem(m, v215, l1, v121)
	mBase = m.M
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
}
func F_pgstat_reset_entry(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v5 = int32(0)
	v9 = F_pgstat_get_entry_ref(m, l0, l1, l2, v5, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 == int32(0) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v19 = F_LWLockAcquire(m, v15+int32(4), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v50 = l0*int32(72) + int32(1630560)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v48 = int32(0)
						} else {
							v36 = int32(0)
							v38 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
							if v38 == v36 {
								v48 = v36
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v38+l0<<(uint(int32(2))%32)-int32(96))))
								v48 = v46
							}
						}
						v50 = v48
					}
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v79 = l0*int32(72) + int32(1630560)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v77 = int32(0)
						} else {
							v65 = int32(0)
							v67 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
							if v67 == v65 {
								v77 = v65
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v67+l0<<(uint(int32(2))%32)-int32(96))))
								v77 = v75
							}
						}
						v79 = v77
					}
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v111 = l0*int32(72) + int32(1630560)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v109 = int32(0)
						} else {
							v97 = int32(0)
							v99 = *(*int32)(unsafe.Add(mBase, _consts[1142]))
							if v99 == v97 {
								v109 = v97
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v99+l0<<(uint(int32(2))%32)-int32(96))))
								v109 = v107
							}
						}
						v111 = v109
					}
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
					v114 = F__emscripten_memset_bulkmem(m, v21+v80, base.I32_extend8_s(int32(0)), v112)
					mBase = m.M
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
					if v115 != 0 {
						m.T0[v115].(func(*base.Module, int32, int64))(m, v21, l3)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
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
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
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
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	v13 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[264]))
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[1212]))
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
	v47 = int32(4094672)
	*(*int32)(unsafe.Add(mBase, _consts[1223])) = v47
	*(*int32)(unsafe.Add(mBase, _consts[1224])) = v47
	v55 = int64(*(*int32)(unsafe.Add(mBase, _consts[157])))
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
	v62 = *(*int32)(unsafe.Add(mBase, _consts[746]))
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
	v72 = *(*int32)(unsafe.Add(mBase, _consts[1225]))
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
		v103 = int32(-1)
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1220]))
	F_pfree(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L35
	}
L18:
	;
	v112 = v103
	v115 = v72
	v118 = int32(0)
	goto L24
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v81 = int32(0)
	goto L20
L20:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v81*int32(24))+16)))
	if v93 != int32(1) {
		v103 = v81
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v103 = int32(-1)
	goto L18
L22:
	;
	v97 = v81 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v97)) < base.Ui64(v74) {
		v81 = v97
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	v124 = v112
	v130 = v118
	v131 = v118
	goto L27
L25:
	;
	F_pfree(m, v123)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v131&int32(1) != 0 {
		goto L26
	} else {
		goto L29
	}
L28:
	;
	if v148 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v138 = int32(1)
	v139 = v124 - v138
	v143 = base.B2i32(v137&(v139^v103) == int32(0))
	v144 = v143 | v130
	v147 = v139 & v137
	v148 = v123 + v124*int32(24)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+16)))
	if v149 != v138 {
		v124 = v147
		v130 = v144
		v131 = v143
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v157
	F_pgstat_release_entry_ref(m, v69, v154, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[1225]))
	v112 = v147
	v115 = v163
	v118 = v144
	goto L24
L33:
	;
	F_pfree(m, v115)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1225])) = int32(0)
	goto L17
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1220])) = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, _consts[1222]))
	F_dsa_detach(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v199 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v204 = v202 + int32(1476)
	v206 = F_LWLockAcquire(m, v204, v199)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202)+1460))
	v210 = v208 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v202)+1460)) = v210
	if v210 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v218 = v199
	goto L41
L39:
	;
	goto L40
L40:
	;
	F_LWLockRelease(m, v204)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L48
	}
L41:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(32)+v218<<(uint(int32(2))%32))))
	if v230 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	F_dsm_unpin_segment(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v234 = v218 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v202)+1456))
	if base.Ui32(v234) <= base.Ui32(v235) {
		v218 = v234
		goto L41
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	goto L42
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1222])) = int32(0)
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
	var v13 int64
	_ = v13
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v7 = v5 + int32(53272)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v13 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[1241]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[1242]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[1243]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[1244]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[1245]))) = l0
		F_LWLockRelease(m, v7)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			return
		}
	}
}
