package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddFileToBackupManifest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	v8 = m.G0
	v10 = v8 - int32(1184)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1176)) = l4
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L8
	} else {
		goto L87
	}
L2:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(1184)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(489010)
	v24 = F_pg_snprintf(m, v10+int32(144), int32(1024), int32(176985), v10+int32(48))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v28 = l2
	goto L7
L7:
	;
	F_initStringInfo(m, v10+int32(128))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	v28 = v10 + int32(144)
	goto L7
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v33 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v48 = F_strlen(m, v28)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	F_appendStringInfoChar(m, v10+int32(128), int32(10))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(755941))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v41)
	goto L11
L16:
	;
	goto L11
L17:
	;
	F_appendStringInfoString(m, v10+int32(128), v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L39
	}
L18:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(731064))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L24
	}
L19:
	;
	v54 = F_pg_verify_mbstr(m, int32(6), v28, v48, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(746094))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_escape_json_with_len(m, v10+int32(128), v28, v48)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v173 = int32(746361)
	goto L17
L24:
	;
	F_enlargeStringInfo(m, v10+int32(128), v48<<(uint(int32(1))%32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v81 = v79 + v80
	v85 = v28 + v48
	if base.Ui32(v85) <= base.Ui32(v28) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v167 + base.I32_wrap_i64(base.I64_extend_i32_u(v48)<<(uint(int64(1))%64))
	v173 = int32(746360)
	goto L17
L27:
	;
	goto L26
L28:
	;
	v88 = v48 & int32(3)
	if v88 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v89 = v28
	v91 = v81
	v92 = int32(0)
	goto L32
L30:
	;
	v109 = v28
	v111 = v81
	goto L31
L31:
	;
	if base.Ui32(v48-int32(1)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L35
	}
L32:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v96 = int32(1)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95<<(uint(v96)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v91))) = uint16(v100)
	v103 = v91 + int32(2)
	v105 = v89 + v96
	v107 = v92 + v96
	if v107 != v88 {
		v89 = v105
		v91 = v103
		v92 = v107
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v105
	v111 = v103
	goto L31
L34:
	;
	goto L33
L35:
	;
	v119 = v109
	v121 = v111
	goto L36
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v126 = int32(1)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125<<(uint(v126)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v130)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132<<(uint(v126)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+2)) = uint16(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139<<(uint(v126)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+4)) = uint16(v144)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146<<(uint(v126)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v151)
	v156 = v119 + int32(4)
	if v156 != v85 {
		v119 = v156
		v121 = v121 + int32(8)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L27
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l3
	F_appendStringInfo(m, v10+int32(128), int32(746130), v10+int32(32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(731084))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v189 = int32(128)
	F_enlargeStringInfo(m, v10+v189, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v201 = F_pg_gmtime(m, v10+int32(1176))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v203 = F_pg_strftime(m, v194+v195, int32(128), int32(508515), v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v203 + v205
	F_appendStringInfoChar(m, v10+int32(128), int32(34))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v213 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v215 = v10 - int32(-64)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	switch v217 - int32(1) {
	case 0:
		goto L55
	case 1:
		goto L54
	case 2:
		goto L53
	case 3:
		goto L52
	case 4:
		goto L51
	default:
		v266 = int32(0)
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	F_appendStringInfoString(m, v10+int32(128), int32(6954))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L8
	} else {
		goto L84
	}
L49:
	;
	if v268 < int32(0) {
		goto L1
	} else {
		goto L64
	}
L50:
	;
	v268 = v266
	goto L49
L51:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v259 = F_pg_cryptohash_final(m, v257, v215, int32(64))
	mBase = m.M
	if v259 < int32(0) {
		v266 = int32(-1)
		goto L50
	} else {
		goto L62
	}
L52:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v249 = F_pg_cryptohash_final(m, v247, v215, int32(48))
	mBase = m.M
	if v249 < int32(0) {
		v266 = int32(-1)
		goto L50
	} else {
		goto L60
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v239 = F_pg_cryptohash_final(m, v237, v215, int32(32))
	mBase = m.M
	if v239 < int32(0) {
		v266 = int32(-1)
		goto L50
	} else {
		goto L58
	}
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v229 = F_pg_cryptohash_final(m, v227, v215, int32(28))
	mBase = m.M
	if v229 < int32(0) {
		v266 = int32(-1)
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v222 = v220 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v222
	v268 = int32(4)
	goto L49
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v268 = int32(28)
	goto L49
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v268 = int32(32)
	goto L49
L60:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	v268 = int32(48)
	goto L49
L62:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v266 = int32(64)
	goto L50
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v271) <= base.Ui32(int32(5)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v271<<(uint(int32(2))%32))+uint32(_consts[409])))
	v280 = v278
	goto L67
L66:
	;
	v280 = int32(546077)
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v280
	F_appendStringInfo(m, v10+int32(128), int32(731020), v10+int32(16))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_enlargeStringInfo(m, v10+int32(128), v268<<(uint(int32(1))%32))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v296 = v10 - int32(-64)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	v299 = v297 + v298
	v303 = v296 + v268
	if base.Ui32(v303) <= base.Ui32(v296) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v10)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v385 + base.I32_wrap_i64(base.I64_extend_i32_u(v268)<<(uint(int64(1))%64))
	F_appendStringInfoChar(m, v10+int32(128), int32(34))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L8
	} else {
		goto L83
	}
L71:
	;
	goto L70
L72:
	;
	v306 = v268 & int32(3)
	if v306 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v307 = v296
	v309 = v299
	v310 = int32(0)
	goto L76
L74:
	;
	v327 = v296
	v329 = v299
	goto L75
L75:
	;
	if base.Ui32(v268-int32(1)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L79
	}
L76:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	v314 = int32(1)
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313<<(uint(v314)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v309))) = uint16(v318)
	v321 = v309 + int32(2)
	v323 = v307 + v314
	v325 = v310 + v314
	if v325 != v306 {
		v307 = v323
		v309 = v321
		v310 = v325
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v327 = v323
	v329 = v321
	goto L75
L78:
	;
	goto L77
L79:
	;
	v337 = v327
	v339 = v329
	goto L80
L80:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v344 = int32(1)
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343<<(uint(v344)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v339))) = uint16(v348)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350<<(uint(v344)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v339)+2)) = uint16(v355)
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+2)))
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357<<(uint(v344)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v339)+4)) = uint16(v362)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+3)))
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364<<(uint(v344)%32))+uint32(_consts[408]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v339)+6)) = uint16(v369)
	v374 = v337 + int32(4)
	if v374 != v303 {
		v337 = v374
		v339 = v339 + int32(8)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L71
L82:
	;
	goto L81
L83:
	;
	goto L48
L84:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	F_AppendStringToManifest(m, l0, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	F_pfree(m, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	goto L4
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
	F_errmsg_internal(m, int32(717239), v10)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(493023), int32(187), int32(77774))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateLockFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int64
	_ = v487
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(2608)
	m.G0 = v14
	v16 = int32(543309)
	v22 = F___strchrnul(m, v16, int32(61))
	mBase = m.M
	if v16 == v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v64 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = v22 - v16
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1114]))))
	if v27 != 0 {
		v57 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = v57
	goto L1
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1115]))
	if v29 == int32(0) {
		v57 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 == int32(0) {
		v57 = v6
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v36 = v29
	v37 = v32
	goto L9
L9:
	;
	v40 = F_strncmp(m, v16, v37, v25)
	mBase = m.M
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v57 = v44 + int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v44 = v43 + v25
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 == int32(61) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v49 != 0 {
		v36 = v36 + int32(4)
		v37 = v49
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v57 = v6
	goto L5
L17:
	;
	v68 = v64
	goto L21
L18:
	;
	v113 = v6
	goto L19
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v115
	v120 = F_open(m, l0, int32(194), v14+int32(288))
	mBase = m.M
	if v120 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v113 = v112
	goto L19
L21:
	;
	v73 = v68 + int32(1)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68))))
	v75 = F___isspace(m, v74)
	mBase = m.M
	if v75 != 0 {
		v68 = v73
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v76 = int32(1)
	switch v74&int32(255) - int32(43) {
	case 0:
		v82 = v76
		goto L25
	default:
		v84 = v74
		v85 = v68
		v86 = v76
		goto L24
	case 2:
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v87 = int32(0)
	v89 = v84 - int32(48)
	if base.Ui32(v89) <= base.Ui32(int32(9)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v84 = v83
	v85 = v73
	v86 = v82
	goto L24
L26:
	;
	v82 = int32(0)
	goto L25
L27:
	;
	v92 = v87
	v93 = v89
	v94 = v85
	goto L30
L28:
	;
	v106 = v87
	goto L29
L29:
	;
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v96 = int32(10)
	v98 = v92*v96 - v93
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94)+1)))
	v103 = v99 - int32(48)
	if base.Ui32(v103) < base.Ui32(v96) {
		v92 = v98
		v93 = v103
		v94 = v94 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v106 = v98
	goto L29
L32:
	;
	goto L31
L33:
	;
	v112 = int32(0) - v106
	goto L35
L34:
	;
	v112 = v106
	goto L35
L35:
	;
	goto L20
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L62
	} else {
		goto L227
	}
L37:
	;
	v733 = int32(4680196)
	v734 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v735 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v734
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L62
	} else {
		goto L223
	}
L38:
	;
	v710 = int32(4680196)
	v711 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v712 = F_close(m, v472)
	mBase = m.M
	v713 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v711
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L62
	} else {
		goto L219
	}
L39:
	;
	v685 = int32(4680196)
	v686 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v687 = F_close(m, v472)
	mBase = m.M
	v688 = F_unlink(m, l0)
	mBase = m.M
	if v686 != 0 {
		goto L212
	} else {
		goto L213
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L62
	} else {
		goto L207
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L62
	} else {
		goto L202
	}
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L62
	} else {
		goto L199
	}
L43:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L62
	} else {
		goto L194
	}
L44:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L62
	} else {
		goto L190
	}
L45:
	;
	v132 = int32(0)
	goto L48
L46:
	;
	v472 = v120
	goto L47
L47:
	;
	v480 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = l2
	v484 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v484
	v487 = *(*int64)(unsafe.Add(mBase, _consts[1117]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v487
	if l1 != 0 {
		goto L162
	} else {
		goto L163
	}
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v136 != int32(20) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v472 = v463
	goto L47
L50:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v146
	v148 = int32(0)
	v151 = F_open(m, l0, v148, v14+int32(272))
	mBase = m.M
	if v151 < v148 {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	if v136 != int32(2) {
		goto L36
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v132) {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(v132) <= base.Ui32(int32(100)) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	goto L50
L57:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v456
	v463 = F_open(m, l0, int32(194), v14+int32(96))
	mBase = m.M
	if v463 < int32(0) {
		v132 = v132 + int32(1)
		goto L48
	} else {
		goto L161
	}
L58:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v155 == int32(44) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(167772190)
	v182 = F_read(m, v151, v14+int32(304), int32(2303))
	mBase = m.M
	if v182 < int32(0) {
		goto L44
	} else {
		goto L67
	}
L61:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = l0
	F_errmsg(m, int32(299036), v14+int32(112))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(493462), int32(1300), int32(390148))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v187
	v189 = F_close(m, v151)
	mBase = m.M
	if v182 == v187 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v193 = v14 + int32(304)
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193+v182))) = uint8(v195)
	v202 = v193
	goto L70
L69:
	;
	v248 = v246 >> (uint(int32(31)) % 32)
	v250 = v246 ^ v248 - v248
	if v250 <= int32(0) {
		goto L42
	} else {
		goto L85
	}
L70:
	;
	v207 = v202 + int32(1)
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v202))))
	v209 = F___isspace(m, v208)
	mBase = m.M
	if v209 != 0 {
		v202 = v207
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v210 = int32(1)
	switch v208&int32(255) - int32(43) {
	case 0:
		v216 = v210
		goto L74
	default:
		v218 = v208
		v219 = v202
		v220 = v210
		goto L73
	case 2:
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v221 = int32(0)
	v223 = v218 - int32(48)
	if base.Ui32(v223) <= base.Ui32(int32(9)) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v207))))
	v218 = v217
	v219 = v207
	v220 = v216
	goto L73
L75:
	;
	v216 = int32(0)
	goto L74
L76:
	;
	v226 = v221
	v227 = v223
	v228 = v219
	goto L79
L77:
	;
	v240 = v221
	goto L78
L78:
	;
	if v220 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v230 = int32(10)
	v232 = v226*v230 - v227
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v228)+1)))
	v237 = v233 - int32(48)
	if base.Ui32(v237) < base.Ui32(v230) {
		v226 = v232
		v227 = v237
		v228 = v228 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v240 = v232
	goto L78
L81:
	;
	goto L80
L82:
	;
	v246 = int32(0) - v240
	goto L84
L83:
	;
	v246 = v240
	goto L84
L84:
	;
	goto L69
L85:
	;
	if v250 == int32(42) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if l3 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L87:
	;
	if v250 == int32(1) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if v250 == v113 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v259 = F_kill(m, v250, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L62
	} else {
		goto L91
	}
L90:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L62
	} else {
		goto L93
	}
L91:
	;
	if v259 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	switch v264 - int32(63) {
	case 0, 8:
		goto L86
	default:
		goto L90
	}
L93:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L62
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = l0
	F_errmsg(m, int32(117065), v14+int32(256))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L62
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v250
	if l3 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v284 = int32(546113)
	goto L98
L97:
	;
	v284 = int32(546237)
	goto L98
L98:
	;
	if l3 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v287 = int32(546174)
	goto L101
L100:
	;
	v287 = int32(546290)
	goto L101
L101:
	;
	if v246 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v290 = v284
	goto L104
L103:
	;
	v290 = v287
	goto L104
L104:
	;
	F_errhint(m, v290, v14+int32(240))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L62
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(493462), int32(1372), int32(390148))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L62
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v449 = F_unlink(m, l0)
	mBase = m.M
	if v449 < int32(0) {
		goto L40
	} else {
		goto L160
	}
L108:
	;
	v304 = int32(10)
	v305 = F___strchrnul(m, v14+int32(304), v304)
	mBase = m.M
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v307 == v304 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v311 == int32(0) {
		goto L107
	} else {
		goto L113
	}
L110:
	;
	v311 = v305
	goto L112
L111:
	;
	v311 = int32(0)
	goto L112
L112:
	;
	goto L109
L113:
	;
	v316 = int32(10)
	v317 = F___strchrnul(m, v311+int32(1), v316)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v319 == v316 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v323 == int32(0) {
		goto L107
	} else {
		goto L118
	}
L115:
	;
	v323 = v317
	goto L117
L116:
	;
	v323 = int32(0)
	goto L117
L117:
	;
	goto L114
L118:
	;
	v328 = int32(10)
	v329 = F___strchrnul(m, v323+int32(1), v328)
	mBase = m.M
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	if v331 == v328 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v335 == int32(0) {
		goto L107
	} else {
		goto L123
	}
L120:
	;
	v335 = v329
	goto L122
L121:
	;
	v335 = int32(0)
	goto L122
L122:
	;
	goto L119
L123:
	;
	v340 = int32(10)
	v341 = F___strchrnul(m, v335+int32(1), v340)
	mBase = m.M
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v343 == v340 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v347 == int32(0) {
		goto L107
	} else {
		goto L128
	}
L125:
	;
	v347 = v341
	goto L127
L126:
	;
	v347 = int32(0)
	goto L127
L127:
	;
	goto L124
L128:
	;
	v352 = int32(10)
	v353 = F___strchrnul(m, v347+int32(1), v352)
	mBase = m.M
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v355 == v352 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v359 == int32(0) {
		goto L107
	} else {
		goto L133
	}
L130:
	;
	v359 = v353
	goto L132
L131:
	;
	v359 = int32(0)
	goto L132
L132:
	;
	goto L129
L133:
	;
	v364 = int32(10)
	v365 = F___strchrnul(m, v359+int32(1), v364)
	mBase = m.M
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v367 == v364 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v371 == int32(0) {
		goto L107
	} else {
		goto L138
	}
L135:
	;
	v371 = v365
	goto L137
L136:
	;
	v371 = int32(0)
	goto L137
L137:
	;
	goto L134
L138:
	;
	v375 = v371 + int32(1)
	if v375 == int32(0) {
		goto L107
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v14 + int32(296)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v14 + int32(300)
	v387 = F_sscanf(m, v375, int32(38355), v14+int32(224))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L62
	} else {
		goto L140
	}
L140:
	;
	if v387 != int32(2) {
		goto L107
	} else {
		goto L141
	}
L141:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	v393 = m.G0
	v395 = v393 - int32(16)
	m.G0 = v395
	v399 = F_PGSharedMemoryAttach(m, v392, v395+int32(12))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L62
	} else {
		goto L142
	}
L142:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	if v401 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	m.G0 = v395 + int32(16)
	if base.Ui32(v399) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L159
	}
L144:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _consts[1118]))
	if v406 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if int32(0) <= v423 {
		goto L143
	} else {
		goto L154
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	v423 = int32(-1)
	goto L145
L147:
	;
	v410 = v406
	goto L148
L148:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	if v401 != v411 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v423 = int32(0)
	goto L145
L150:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+20))
	if v413 != 0 {
		v410 = v413
		goto L148
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	goto L146
L154:
	;
	v428 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L62
	} else {
		goto L155
	}
L155:
	;
	if v428 == int32(0) {
		goto L143
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v401
	F_errmsg_internal(m, int32(295780), v395)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L62
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(497468), int32(324), int32(362842))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L62
	} else {
		goto L158
	}
L158:
	;
	goto L143
L159:
	;
	goto L107
L160:
	;
	goto L57
L161:
	;
	goto L49
L162:
	;
	v491 = int32(42)
	goto L164
L163:
	;
	v491 = int32(-42)
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v491
	v499 = F_pg_snprintf(m, v14+int32(304), int32(2304), int32(748485), v14+int32(48))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L62
	} else {
		goto L165
	}
L165:
	;
	if l1 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v519 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = int32(167772192)
	v523 = v14 + int32(304)
	v526 = F_strlen(m, v523)
	mBase = m.M
	v527 = F_write(m, v472, v523, v526)
	mBase = m.M
	v530 = F_strlen(m, v523)
	mBase = m.M
	if v527 != v530 {
		goto L39
	} else {
		goto L174
	}
L167:
	;
	if l3 == int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v504 = v14 + int32(304)
	v505 = int32(757453)
	v506 = int32(2304)
	v508 = F_pg_ascii_verifystr(m, v504, v506)
	mBase = m.M
	if v508 == v506 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L166
L170:
	;
	goto L169
L171:
	;
	v510 = F_strlen(m, v505)
	mBase = m.M
	goto L170
L172:
	;
	goto L173
L173:
	;
	v513 = F_strlcpy(m, v504+v508, v505, v506-v508)
	mBase = m.M
	goto L170
L174:
	;
	v532 = int32(4122220)
	v533 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v534 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v534
	v537 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = int32(167772191)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v542 != int32(1) {
		v556 = v534
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v556 != 0 {
		goto L38
	} else {
		goto L182
	}
L176:
	;
	goto L175
L177:
	;
	goto L178
L178:
	;
	v547 = F_fsync(m, v472)
	mBase = m.M
	if v547 != int32(-1) {
		v556 = v547
		goto L176
	} else {
		goto L180
	}
L179:
	;
	v556 = int32(-1)
	goto L176
L180:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v551 == int32(27) {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = int32(0)
	v561 = F_close(m, v472)
	mBase = m.M
	if v561 != 0 {
		goto L37
	} else {
		goto L183
	}
L183:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[1119]))
	if v563 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	F_on_proc_exit(m, int32(1639))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L62
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v570 = F_pstrdup(m, l0)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L62
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _consts[1119]))
	v574 = F_lcons(m, v570, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L62
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1119])) = v574
	m.G0 = v14 + int32(2608)
	return
L190:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L62
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = l0
	F_errmsg(m, int32(299181), v14+int32(128))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L62
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(493462), int32(1307), int32(390148))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L62
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L62
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = l0
	F_errmsg(m, int32(8837), v14+int32(144))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L62
	} else {
		goto L196
	}
L196:
	;
	F_errhint(m, int32(624651), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L62
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(493462), int32(1316), int32(390148))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L62
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v14 + int32(304)
	F_errmsg_internal(m, int32(728410), v14+int32(160))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L62
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(493462), int32(1327), int32(390148))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L62
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L62
	} else {
		goto L203
	}
L203:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v14)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v646
	F_errmsg(m, int32(360376), v14+int32(208))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L62
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = l4
	F_errhint(m, int32(664797), v14+int32(192))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L62
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(493462), int32(1410), int32(390148))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L62
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L62
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = l0
	F_errmsg(m, int32(299141), v14+int32(176))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L62
	} else {
		goto L209
	}
L209:
	;
	F_errhint(m, int32(619726), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L62
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(493462), int32(1426), int32(390148))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L62
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	v691 = v686
	goto L214
L213:
	;
	v691 = int32(51)
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v691
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L62
	} else {
		goto L215
	}
L215:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L62
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	F_errmsg(m, int32(299070), v14+int32(32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L62
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(493462), int32(1461), int32(390148))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L62
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L62
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(299070), v14+int32(16))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L62
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(493462), int32(1475), int32(390148))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L62
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L62
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(299070), v14)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L62
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(493462), int32(1486), int32(390148))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L62
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L62
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l0
	F_errmsg(m, int32(299105), v14+int32(80))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L62
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(493462), int32(1286), int32(390148))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L62
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = F_FileAccess(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v16 != int32(1) {
		v38 = int32(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(0)
	return v38
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0*int32(48))))
	goto L8
L8:
	;
	v28 = F_fsync(m, v24)
	mBase = m.M
	if v28 != int32(-1) {
		v38 = v28
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(-1)
	goto L6
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v32 == int32(27) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
func F_load_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v90 = F_expand_dynamic_library_name(m, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L34
	}
L2:
	;
	v10 = int32(570889)
	goto L5
L3:
	;
	if v47-v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = l0
	v19 = v10
	v20 = int32(16)
	v21 = v17
	goto L11
L8:
	;
	v43 = v10
	v47 = int32(0)
	goto L9
L9:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	goto L3
L10:
	;
	v43 = v38
	v47 = v40
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 != v23 {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v38 = v32
	v40 = int32(0)
	goto L10
L13:
	;
	if v23 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v28 = v20 - int32(1)
	if v28 == int32(0) {
		v38 = v19
		v40 = v21
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(1)
	v32 = v19 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v33 != 0 {
		v18 = v18 + v31
		v19 = v32
		v20 = v28
		v21 = v33
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v61 = l0 + int32(16)
	goto L21
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	if v71 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v63 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	goto L22
L24:
	;
	v71 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v63 == int32(47) {
		v71 = v61
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v61 = v61 + int32(1)
	goto L21
L28:
	;
	goto L19
L29:
	;
	return
L30:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(439463), v6)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(495425), int32(528), int32(378478))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v92 = F_internal_load_library(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_sendFileWithContent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v18 = F_pg_checksum_init(m, v13+int32(24), v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L41
	}
L2:
	;
	return
L3:
	;
	if int32(0) <= v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = F_strlen(m, l2)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(123)
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v25
	v27 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = base.I64_extend_i32_s(v22)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v32
	F__tarWriteHeader(m, l0, l1, v25, v13+int32(32), v25)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L38
	}
L7:
	;
	v42 = F_pg_checksum_update(m, v13+int32(24), l2, v22)
	mBase = m.M
	if v42 < int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if int32(0) < v22 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = l2
	v54 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v85 = (v22+int32(511))&int32(-512) - v22
	if int32(0) < v85 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v59 = v22 - v54
	if base.Ui32(v58) < base.Ui32(v59) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v61 = v58
	goto L16
L15:
	;
	v61 = v59
	goto L16
L16:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	m.T0[v65].(func(*base.Module, int32, int32))(m, l0, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L21
	}
L18:
	;
	v62 = F__emscripten_memcpy_bulkmem(m, v57, v49, v61)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v69 = v61 + v54
	if v69 < v22 {
		v49 = v49 + v61
		v54 = v69
		goto L12
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v88&int32(3) != 0 {
		v110 = v85
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	F_AddFileToBackupManifest(m, l3, int32(0), l1, v22, v27, v13+int32(24))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L37
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	m.T0[v117].(func(*base.Module, int32, int32))(m, l0, v85)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L36
	}
L27:
	;
	v113 = F__emscripten_memset_bulkmem(m, v88, base.I32_extend8_s(int32(0)), v110)
	mBase = m.M
	goto L35
L28:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v85) {
		v110 = v85
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v85&int32(3) != 0 {
		v110 = v85
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v95 = v85 + v88
	if base.Ui32(v95) <= base.Ui32(v88) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v100 = v88 + int32(4)
	if base.Ui32(v100) < base.Ui32(v95) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = v95
	goto L34
L33:
	;
	v102 = v100
	goto L34
L34:
	;
	v110 = (v88^int32(-1)+v102)&int32(-4) + int32(4)
	goto L27
L35:
	;
	goto L26
L36:
	;
	goto L25
L37:
	;
	m.G0 = v13 + int32(128)
	return
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg_internal(m, int32(717280), v13)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(495980), int32(1084), int32(93412))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg_internal(m, int32(717323), v13+int32(16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(495980), int32(1109), int32(93412))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
