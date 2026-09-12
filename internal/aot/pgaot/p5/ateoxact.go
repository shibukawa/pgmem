package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v118 int32
	_ = v118
	var v133 int32
	_ = v133
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 float64
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1212 int32
	_ = v1212
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1288 int32
	_ = v1288
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1341 int32
	_ = v1341
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = l0
	v40 = l1
	v57 = v37
	v58 = int32(4547068)
	goto L4
L2:
	;
	v1341 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v1341 - int32(1)
	return
L4:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v76 = v57 - int32(16)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v1341 = v1306
	goto L3
L6:
	;
	if v1323 != 0 {
		v39 = v1305
		v40 = v1306
		v57 = v1323
		v58 = v1324
		goto L4
	} else {
		goto L292
	}
L7:
	;
	v1305 = v39
	v1306 = v40
	v1323 = v74
	v1324 = v57
	goto L6
L8:
	;
	goto L9
L9:
	;
	v80 = int32(72)
	v81 = v57 - v80
	v82 = int32(4)
	v83 = v57 + v82
	v84 = int32(44)
	v85 = v57 - v84
	v86 = int32(52)
	v88 = int32(24)
	v93 = v57 - v82
	v95 = v57 - int32(8)
	v97 = v57 - int32(40)
	v99 = v57 - int32(12)
	v101 = v57 - int32(48)
	v103 = v57 + v84
	v105 = v57 + v86
	v107 = v57 + v80
	v109 = v57 + int32(20)
	v111 = v57 + v88
	v118 = v77
	v133 = v57
	goto L10
L10:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v149 < v40 {
		v1305 = v39
		v1306 = v40
		v1323 = v74
		v1324 = v133
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v1305 = v39
	v1306 = v40
	v1323 = v74
	v1324 = v1288
	goto L6
L12:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v39 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v1304 != 0 {
		v118 = v1304
		v133 = v1288
		goto L10
	} else {
		goto L291
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v151
	if v151 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L15:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v989 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v989
	if v988 == v989 {
		goto L227
	} else {
		goto L228
	}
L16:
	;
	F_pfree(m, v918)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L38
	} else {
		goto L226
	}
L17:
	;
	v918 = v883
	v921 = v886
	goto L16
L18:
	;
	F_discard_stack_value(m, v81, v118+int32(32))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L38
	} else {
		goto L225
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v151
	F_pfree(m, v118)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L38
	} else {
		goto L224
	}
L20:
	;
	F_discard_stack_value(m, v81, v118+int32(32))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L38
	} else {
		goto L223
	}
L21:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235+v118)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v118+v238)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v236)))
	v247 = base.I32_wrap_i64(base.I64_reinterpret_f64(v245))
	v248 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v249 {
	case 0:
		goto L54
	case 1:
		goto L53
	case 2:
		goto L52
	case 3:
		goto L51
	case 4:
		goto L50
	default:
		v958 = v248
		goto L15
	}
L22:
	;
	v235 = int32(28)
	v236 = v118 + int32(48)
	v237 = int32(13)
	v238 = int32(20)
	v239 = v118 + int32(56)
	goto L21
L23:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v235 = int32(24)
	v236 = v118 + int32(32)
	v237 = v223
	v238 = int32(16)
	v239 = v118 + int32(40)
	goto L21
L24:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if v154 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v149 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	switch v154 - int32(1) {
	case 0:
		goto L18
	default:
		goto L23
	case 2:
		goto L22
	}
L27:
	;
	goto L28
L28:
	;
	if v151 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v149 - int32(1)
	v1288 = v133
	goto L13
L30:
	;
	goto L31
L31:
	;
	v167 = v149 - int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v168 < v167 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v167
	v1288 = v133
	goto L13
L33:
	;
	goto L34
L34:
	;
	switch v154 - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	default:
		goto L19
	}
L35:
	;
	F_discard_stack_value(m, v81, v118+int32(32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L38
	} else {
		goto L45
	}
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	if v186 != int32(1) {
		goto L20
	} else {
		goto L44
	}
L37:
	;
	F_discard_stack_value(m, v81, v118+int32(32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	if v177 == int32(3) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_discard_stack_value(m, v81, v151+int32(48))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(1)
	goto L19
L43:
	;
	goto L42
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v191
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+48)) = v193
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v118)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+56)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(3)
	goto L19
L45:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	if v207 == int32(3) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_discard_stack_value(m, v81, v151+int32(48))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L38
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v118)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+48)) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v118)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v151)+56)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = int32(3)
	goto L19
L49:
	;
	goto L48
L50:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	if v247 == v800 {
		goto L197
	} else {
		goto L198
	}
L51:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	if v463 != v247 {
		goto L134
	} else {
		goto L135
	}
L52:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v392)))
	if base.F64_eq(v245, v393) != 0 {
		goto L107
	} else {
		goto L108
	}
L53:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if v247 == v323 {
		goto L81
	} else {
		goto L82
	}
L54:
	;
	v251 = v247 & int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v251 == v253 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v255 == v244 {
		v958 = v248
		goto L15
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v257 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	m.T0[v257].(func(*base.Module, int32, int32))(m, v251, v244)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L38
	} else {
		goto L62
	}
L60:
	;
	v261 = v252
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v251)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v244
	v265 = int32(1)
	if v263 == int32(0) {
		v958 = v265
		goto L15
	} else {
		goto L63
	}
L62:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v261 = v260
	goto L61
L63:
	;
	if v263 == v244 {
		v958 = v265
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v269 {
	case 0:
		goto L70
	case 1:
		goto L69
	case 2:
		goto L68
	case 3:
		goto L67
	case 4:
		goto L66
	default:
		goto L65
	}
L65:
	;
	v283 = v76
	goto L76
L66:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v263 == v278 {
		v958 = v265
		goto L15
	} else {
		goto L75
	}
L67:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v263 != v276 {
		goto L65
	} else {
		goto L74
	}
L68:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v263 != v274 {
		goto L65
	} else {
		goto L73
	}
L69:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v263 != v272 {
		goto L65
	} else {
		goto L72
	}
L70:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v263 != v270 {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v958 = v265
	goto L15
L72:
	;
	v958 = v265
	goto L15
L73:
	;
	v958 = v265
	goto L15
L74:
	;
	v958 = v265
	goto L15
L75:
	;
	goto L65
L76:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v315 == int32(0) {
		v883 = v263
		v886 = v265
		goto L17
	} else {
		goto L78
	}
L77:
	;
	v958 = v265
	goto L15
L78:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v315)+40))
	if v263 == v318 {
		v958 = v265
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)+56))
	if v263 != v320 {
		v283 = v315
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v325 == v244 {
		v958 = v248
		goto L15
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v327 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	m.T0[v327].(func(*base.Module, int32, int32))(m, v247, v244)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L38
	} else {
		goto L88
	}
L86:
	;
	v331 = v322
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v247
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v244
	v335 = int32(1)
	if v333 == int32(0) {
		v958 = v335
		goto L15
	} else {
		goto L89
	}
L88:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v331 = v330
	goto L87
L89:
	;
	if v333 == v244 {
		v958 = v335
		goto L15
	} else {
		goto L90
	}
L90:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v339 {
	case 0:
		goto L96
	case 1:
		goto L95
	case 2:
		goto L94
	case 3:
		goto L93
	case 4:
		goto L92
	default:
		goto L91
	}
L91:
	;
	v353 = v76
	goto L102
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v333 == v348 {
		v958 = v335
		goto L15
	} else {
		goto L101
	}
L93:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v333 != v346 {
		goto L91
	} else {
		goto L100
	}
L94:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v333 != v344 {
		goto L91
	} else {
		goto L99
	}
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v333 != v342 {
		goto L91
	} else {
		goto L98
	}
L96:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v333 != v340 {
		goto L91
	} else {
		goto L97
	}
L97:
	;
	v958 = v335
	goto L15
L98:
	;
	v958 = v335
	goto L15
L99:
	;
	v958 = v335
	goto L15
L100:
	;
	v958 = v335
	goto L15
L101:
	;
	goto L91
L102:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v385 == int32(0) {
		v883 = v333
		v886 = v335
		goto L17
	} else {
		goto L104
	}
L103:
	;
	v958 = v335
	goto L15
L104:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v385)+40))
	if v333 == v388 {
		v958 = v335
		goto L15
	} else {
		goto L105
	}
L105:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v385)+56))
	if v333 != v390 {
		v353 = v385
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v395 == v244 {
		v958 = v248
		goto L15
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v397 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	m.T0[v397].(func(*base.Module, float64, int32))(m, v245, v244)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L38
	} else {
		goto L114
	}
L112:
	;
	v401 = v392
	goto L113
L113:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v401))) = v245
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v244
	v405 = int32(1)
	if v403 == int32(0) {
		v958 = v405
		goto L15
	} else {
		goto L115
	}
L114:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v401 = v400
	goto L113
L115:
	;
	if v403 == v244 {
		v958 = v405
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v409 {
	case 0:
		goto L122
	case 1:
		goto L121
	case 2:
		goto L120
	case 3:
		goto L119
	case 4:
		goto L118
	default:
		goto L117
	}
L117:
	;
	v423 = v76
	goto L128
L118:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v403 == v418 {
		v958 = v405
		goto L15
	} else {
		goto L127
	}
L119:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v403 != v416 {
		goto L117
	} else {
		goto L126
	}
L120:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v403 != v414 {
		goto L117
	} else {
		goto L125
	}
L121:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v403 != v412 {
		goto L117
	} else {
		goto L124
	}
L122:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v403 != v410 {
		goto L117
	} else {
		goto L123
	}
L123:
	;
	v958 = v405
	goto L15
L124:
	;
	v958 = v405
	goto L15
L125:
	;
	v958 = v405
	goto L15
L126:
	;
	v958 = v405
	goto L15
L127:
	;
	goto L117
L128:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	if v455 == int32(0) {
		v883 = v403
		v886 = v405
		goto L17
	} else {
		goto L130
	}
L129:
	;
	v958 = v405
	goto L15
L130:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v455)+40))
	if v403 == v458 {
		v958 = v405
		goto L15
	} else {
		goto L131
	}
L131:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v455)+56))
	if v403 != v460 {
		v423 = v455
		goto L128
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	v657 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+32)) = v657
	if v656 == v657 {
		goto L175
	} else {
		goto L176
	}
L134:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if v468 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v465 != v244 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v626 = int32(0)
	goto L133
L137:
	;
	m.T0[v468].(func(*base.Module, int32, int32))(m, v247, v244)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L38
	} else {
		goto L140
	}
L138:
	;
	v473 = v463
	v474 = v462
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v247
	if v473 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v473 = v472
	v474 = v471
	goto L139
L141:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v244
	v564 = int32(1)
	if v562 == int32(0) {
		v626 = v564
		goto L133
	} else {
		goto L154
	}
L142:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	if v473 == v479 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v473 == v481 {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v473 == v483 {
		goto L141
	} else {
		goto L145
	}
L145:
	;
	v487 = v76
	goto L146
L146:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v520 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	F_pfree(m, v473)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L38
	} else {
		goto L153
	}
L148:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+32))
	if v473 == v521 {
		goto L141
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+48))
	if v473 != v523 {
		v487 = v520
		goto L146
	} else {
		goto L152
	}
L152:
	;
	goto L141
L153:
	;
	goto L141
L154:
	;
	if v244 == v562 {
		v626 = v564
		goto L133
	} else {
		goto L155
	}
L155:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v568 {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	case 3:
		goto L158
	case 4:
		goto L157
	default:
		goto L156
	}
L156:
	;
	v581 = v76
	goto L167
L157:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v562 == v577 {
		v626 = v564
		goto L133
	} else {
		goto L166
	}
L158:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v562 != v575 {
		goto L156
	} else {
		goto L165
	}
L159:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v562 != v573 {
		goto L156
	} else {
		goto L164
	}
L160:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v562 != v571 {
		goto L156
	} else {
		goto L163
	}
L161:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v562 != v569 {
		goto L156
	} else {
		goto L162
	}
L162:
	;
	v626 = v564
	goto L133
L163:
	;
	v626 = v564
	goto L133
L164:
	;
	v626 = v564
	goto L133
L165:
	;
	v626 = v564
	goto L133
L166:
	;
	goto L156
L167:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	if v614 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	F_pfree(m, v562)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L38
	} else {
		goto L174
	}
L169:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+40))
	if v562 == v615 {
		v626 = v564
		goto L133
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	goto L168
L172:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614)+56))
	if v562 != v617 {
		v581 = v614
		goto L167
	} else {
		goto L173
	}
L173:
	;
	v626 = v564
	goto L133
L174:
	;
	v626 = v564
	goto L133
L175:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
	v746 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+48)) = v746
	if v745 == v746 {
		v958 = v626
		goto L15
	} else {
		goto L188
	}
L176:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	if v656 == v662 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v656 == v664 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v656 == v666 {
		goto L175
	} else {
		goto L179
	}
L179:
	;
	v670 = v76
	goto L180
L180:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	if v703 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	F_pfree(m, v656)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L38
	} else {
		goto L187
	}
L182:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)+32))
	if v656 == v704 {
		goto L175
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	goto L181
L185:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+48))
	if v656 != v706 {
		v670 = v703
		goto L180
	} else {
		goto L186
	}
L186:
	;
	goto L175
L187:
	;
	goto L175
L188:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	if v745 == v751 {
		v958 = v626
		goto L15
	} else {
		goto L189
	}
L189:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v745 == v753 {
		v958 = v626
		goto L15
	} else {
		goto L190
	}
L190:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v745 == v755 {
		v958 = v626
		goto L15
	} else {
		goto L191
	}
L191:
	;
	v760 = v76
	goto L192
L192:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	if v792 == int32(0) {
		v918 = v745
		v921 = v626
		goto L16
	} else {
		goto L194
	}
L193:
	;
	v958 = v626
	goto L15
L194:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v792)+32))
	if v745 == v795 {
		v958 = v626
		goto L15
	} else {
		goto L195
	}
L195:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v792)+48))
	if v745 != v797 {
		v760 = v792
		goto L192
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v802 == v244 {
		v958 = v248
		goto L15
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v57+int32(36))))
	if v804 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L199
L201:
	;
	m.T0[v804].(func(*base.Module, int32, int32))(m, v247, v244)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L38
	} else {
		goto L204
	}
L202:
	;
	v808 = v799
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v808))) = v247
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v244
	v812 = int32(1)
	if v810 == int32(0) {
		v958 = v812
		goto L15
	} else {
		goto L205
	}
L204:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v808 = v807
	goto L203
L205:
	;
	if v810 == v244 {
		v958 = v812
		goto L15
	} else {
		goto L206
	}
L206:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v816 {
	case 0:
		goto L212
	case 1:
		goto L211
	case 2:
		goto L210
	case 3:
		goto L209
	case 4:
		goto L208
	default:
		goto L207
	}
L207:
	;
	v830 = v76
	goto L218
L208:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v810 == v825 {
		v958 = v812
		goto L15
	} else {
		goto L217
	}
L209:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v810 != v823 {
		goto L207
	} else {
		goto L216
	}
L210:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v810 != v821 {
		goto L207
	} else {
		goto L215
	}
L211:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v810 != v819 {
		goto L207
	} else {
		goto L214
	}
L212:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v810 != v817 {
		goto L207
	} else {
		goto L213
	}
L213:
	;
	v958 = v812
	goto L15
L214:
	;
	v958 = v812
	goto L15
L215:
	;
	v958 = v812
	goto L15
L216:
	;
	v958 = v812
	goto L15
L217:
	;
	goto L207
L218:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	if v862 == int32(0) {
		v883 = v810
		v886 = v812
		goto L17
	} else {
		goto L220
	}
L219:
	;
	v958 = v812
	goto L15
L220:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v862)+40))
	if v810 == v865 {
		v958 = v812
		goto L15
	} else {
		goto L221
	}
L221:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v862)+56))
	if v810 != v867 {
		v830 = v862
		goto L218
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	goto L19
L224:
	;
	v1288 = v133
	goto L13
L225:
	;
	v1212 = int32(0)
	goto L14
L226:
	;
	v958 = v921
	goto L15
L227:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v118)+56))
	v1084 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+56)) = v1084
	if v1083 == v1084 {
		goto L249
	} else {
		goto L250
	}
L228:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v988 == v993 {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v995 {
	case 0:
		goto L235
	case 1:
		goto L234
	case 2:
		goto L233
	case 3:
		goto L232
	case 4:
		goto L231
	default:
		goto L230
	}
L230:
	;
	v1008 = v76
	goto L241
L231:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v988 == v1004 {
		goto L227
	} else {
		goto L240
	}
L232:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v988 != v1002 {
		goto L230
	} else {
		goto L239
	}
L233:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v988 != v1000 {
		goto L230
	} else {
		goto L238
	}
L234:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v988 != v998 {
		goto L230
	} else {
		goto L237
	}
L235:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v988 != v996 {
		goto L230
	} else {
		goto L236
	}
L236:
	;
	goto L227
L237:
	;
	goto L227
L238:
	;
	goto L227
L239:
	;
	goto L227
L240:
	;
	goto L230
L241:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	if v1041 != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	F_pfree(m, v988)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L38
	} else {
		goto L248
	}
L243:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+40))
	if v988 == v1042 {
		goto L227
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	goto L242
L246:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+56))
	if v988 != v1044 {
		v1008 = v1041
		goto L241
	} else {
		goto L247
	}
L247:
	;
	goto L227
L248:
	;
	goto L227
L249:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v1178 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L250:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v1083 == v1088 {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	switch v1090 {
	case 0:
		goto L257
	case 1:
		goto L256
	case 2:
		goto L255
	case 3:
		goto L254
	case 4:
		goto L253
	default:
		goto L252
	}
L252:
	;
	v1103 = v76
	goto L263
L253:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if v1083 == v1099 {
		goto L249
	} else {
		goto L262
	}
L254:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v1083 != v1097 {
		goto L252
	} else {
		goto L261
	}
L255:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v1083 != v1095 {
		goto L252
	} else {
		goto L260
	}
L256:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v1083 != v1093 {
		goto L252
	} else {
		goto L259
	}
L257:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v1083 != v1091 {
		goto L252
	} else {
		goto L258
	}
L258:
	;
	goto L249
L259:
	;
	goto L249
L260:
	;
	goto L249
L261:
	;
	goto L249
L262:
	;
	goto L252
L263:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1103)))
	if v1136 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	F_pfree(m, v1083)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L38
	} else {
		goto L270
	}
L265:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+40))
	if v1083 == v1137 {
		goto L249
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	goto L264
L268:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+56))
	if v1083 != v1139 {
		v1103 = v1136
		goto L263
	} else {
		goto L269
	}
L269:
	;
	goto L249
L270:
	;
	goto L249
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v57-int32(32)))) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v57-v88))) = v241
	v1212 = v958
	goto L14
L272:
	;
	if v237 == int32(0) {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	if v237 != 0 {
		goto L271
	} else {
		goto L280
	}
L275:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, _consts[1277]))
	if v1184 != 0 {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1192 = int32(4547052)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v1192
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v1191
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+4)) = v95
	*(*int32)(unsafe.Add(mBase, _consts[1278])) = v95
	goto L271
L277:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
	v1191 = v1186
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1188 = int32(4547052)
	*(*int32)(unsafe.Add(mBase, _consts[1277])) = v1188
	v1191 = v1188
	goto L276
L280:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+4)) = v1199
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v1199))) = v1201
	goto L271
L281:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57-v86))))
	if v1253&int32(64) == int32(0) {
		v1288 = v1252
		goto L13
	} else {
		goto L289
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v74
	F_pfree(m, v118)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L38
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_pfree(m, v118)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L38
	} else {
		goto L287
	}
L285:
	;
	if v1212 != 0 {
		v1252 = v58
		goto L281
	} else {
		goto L286
	}
L286:
	;
	v1288 = v58
	goto L13
L287:
	;
	if v1212 == int32(0) {
		v1288 = v133
		goto L13
	} else {
		goto L288
	}
L288:
	;
	v1252 = v133
	goto L281
L289:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v1258&int32(4) != 0 {
		v1288 = v1252
		goto L13
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v1258 | int32(4)
	v1264 = int32(4547060)
	v1265 = *(*int32)(unsafe.Add(mBase, _consts[1280]))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v1265
	*(*int32)(unsafe.Add(mBase, _consts[1280])) = v83
	v1288 = v1252
	goto L13
L291:
	;
	goto L11
L292:
	;
	goto L5
}
func F_AtEOXact_MultiXact(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, _consts[73]))
	v3 = int32(4155456)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	v5 = int32(2)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(v5)%32)))) = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	*(*int32)(unsafe.Add(mBase, uint32(v11+v13<<(uint(v5)%32)))) = v8
	v20 = int32(4146096)
	*(*int32)(unsafe.Add(mBase, _consts[76])) = v20
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v20
	*(*int32)(unsafe.Add(mBase, _consts[78])) = v8
	*(*int32)(unsafe.Add(mBase, _consts[79])) = v8
	return
}
