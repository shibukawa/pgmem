package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_log_destination(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_destination[0])) = v4
	return
}
func F_assign_log_extension_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_extension_options[0])) = l1
	return
}
func F_check_log_min_messages(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(208)
	m.G0 = v17
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+96)) = uint16(v4)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v21
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_guc_strdup(m, int32(15), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(208)
	return v651
L2:
	;
	return int32(0)
L3:
	;
	if v27 == int32(0) {
		v651 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_SplitGUCList(m, v27, v17+int32(204))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v41
	v46 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	if v54 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[2])) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	F_list_free(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_bms_free(m, v27)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v651 = v4
	goto L1
L12:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_bms_free(m, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L2
	} else {
		goto L178
	}
L13:
	;
	v608 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v611
	v616 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_1), v608)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L175
	}
L14:
	;
	v57 = int32(-1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if int32(0) < v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = v57
	v72 = v4
	goto L18
L16:
	;
	v432 = v57
	goto L17
L17:
	;
	if v432 == int32(-1) {
		goto L13
	} else {
		goto L106
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v72<<(uint(int32(2))%32))))
	v80 = int32(58)
	v81 = F___strchrnul(m, v79, v80)
	mBase = m.M
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v83 == v80 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v432 = v414
	goto L17
L20:
	;
	v425 = v72 + int32(1)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v425 < v426 {
		v65 = v414
		v72 = v425
		goto L18
	} else {
		goto L105
	}
L21:
	;
	v408 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v408)
	v414 = v65
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[2])) = v396
	F_bms_free(m, v27)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L2
	} else {
		goto L103
	}
L23:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v376
	v380 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_2), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L102
	}
L24:
	;
	if v87 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v87 = v81
	goto L27
L26:
	;
	v87 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	if v65 != int32(-1) {
		goto L23
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v178)
	v181 = v87 + int32(1)
	v182 = int32(_a_F_check_log_min_messages_3)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[3]))
	if v184 != 0 {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v92 = int32(_a_F_check_log_min_messages_3)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[3]))
	if v94 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v414 = v177
	goto L20
L33:
	;
	v98 = v92
	v100 = v94
	goto L36
L34:
	;
	goto L35
L35:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v79
	v175 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_4), v17)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L53
	}
L36:
	;
	v111 = v100
	v112 = v79
	goto L39
L37:
	;
	goto L35
L38:
	;
	if v149 == int32(0) {
		goto L32
	} else {
		goto L51
	}
L39:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v115 == v116 {
		v138 = v115
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v149 = int32(0)
	goto L38
L41:
	;
	v140 = int32(1)
	if v138 != 0 {
		v111 = v111 + v140
		v112 = v112 + v140
		goto L39
	} else {
		goto L50
	}
L42:
	;
	if base.Ui32((v115-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v126 = v115 | int32(32)
	goto L45
L44:
	;
	v126 = v115
	goto L45
L45:
	;
	if base.Ui32((v116-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v135 = v116 | int32(32)
	goto L48
L47:
	;
	v135 = v116
	goto L48
L48:
	;
	if v126 == v135 {
		v138 = v126
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v149 = v126 - v135
	goto L38
L50:
	;
	goto L40
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if v152 != 0 {
		v98 = v98 + int32(12)
		v100 = v152
		goto L36
	} else {
		goto L52
	}
L52:
	;
	goto L37
L53:
	;
	v396 = v175
	goto L22
L54:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v271 = int32(0)
	v276 = v271
	v278 = v271
	goto L76
L55:
	;
	v188 = v182
	v190 = v184
	goto L58
L56:
	;
	goto L57
L57:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v181
	v268 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_5), v17+int32(32))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L75
	}
L58:
	;
	v201 = v190
	v202 = v181
	goto L61
L59:
	;
	goto L57
L60:
	;
	if v239 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L61:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v205 == v206 {
		v228 = v205
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v239 = int32(0)
	goto L60
L63:
	;
	v230 = int32(1)
	if v228 != 0 {
		v201 = v201 + v230
		v202 = v202 + v230
		goto L61
	} else {
		goto L72
	}
L64:
	;
	if base.Ui32((v205-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v216 = v205 | int32(32)
	goto L67
L66:
	;
	v216 = v205
	goto L67
L67:
	;
	if base.Ui32((v206-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v225 = v206 | int32(32)
	goto L70
L69:
	;
	v225 = v206
	goto L70
L70:
	;
	if v216 == v225 {
		v228 = v216
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v239 = v216 - v225
	goto L60
L72:
	;
	goto L62
L73:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if v242 != 0 {
		v188 = v188 + int32(12)
		v190 = v242
		goto L58
	} else {
		goto L74
	}
L74:
	;
	goto L59
L75:
	;
	v396 = v268
	goto L22
L76:
	;
	v288 = v276 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+uint32(_c_F_check_log_min_messages[4])))
	v292 = v289
	v293 = v79
	goto L79
L77:
	;
	if v278 != 0 {
		goto L21
	} else {
		goto L100
	}
L78:
	;
	if v330 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L79:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v296 == v297 {
		v319 = v296
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v330 = int32(0)
	goto L78
L81:
	;
	v321 = int32(1)
	if v319 != 0 {
		v292 = v292 + v321
		v293 = v293 + v321
		goto L79
	} else {
		goto L90
	}
L82:
	;
	if base.Ui32((v296-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v307 = v296 | int32(32)
	goto L85
L84:
	;
	v307 = v296
	goto L85
L85:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v316 = v297 | int32(32)
	goto L88
L87:
	;
	v316 = v297
	goto L88
L88:
	;
	if v307 == v316 {
		v319 = v307
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v330 = v307 - v316
	goto L78
L90:
	;
	goto L80
L91:
	;
	v335 = v17 + int32(80) + v276
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v336 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v361 = v276 + int32(1)
	if v361 != int32(18) {
		v276 = v361
		goto L76
	} else {
		goto L99
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(112)+v288))) = v270
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v335))) = uint8(v343)
	v347 = v276 + v343
	if v347 != int32(18) {
		v276 = v347
		v278 = v343
		goto L76
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v79
	v358 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_6), v17+int32(48))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L98
	}
L97:
	;
	goto L21
L98:
	;
	v396 = v358
	goto L22
L99:
	;
	goto L77
L100:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[1])) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v79
	v372 = F_format_elog_string(m, int32(_a_F_check_log_min_messages_7), v17-int32(-64))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v396 = v372
	goto L22
L102:
	;
	v396 = v380
	goto L22
L103:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	F_list_free(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v651 = int32(0)
	goto L1
L105:
	;
	goto L19
L106:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)))
	if v444 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v432
	goto L109
L108:
	;
	goto L109
L109:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+81)))
	if v448 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = v432
	goto L112
L111:
	;
	goto L112
L112:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+82)))
	if v452 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v432
	goto L115
L114:
	;
	goto L115
L115:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	if v456 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v432
	goto L118
L117:
	;
	goto L118
L118:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+84)))
	if v460 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v432
	goto L121
L120:
	;
	goto L121
L121:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+85)))
	if v464 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = v432
	goto L124
L123:
	;
	goto L124
L124:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+86)))
	if v468 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v432
	goto L127
L126:
	;
	goto L127
L127:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)))
	if v472 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v432
	goto L130
L129:
	;
	goto L130
L130:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)))
	if v476 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v432
	goto L133
L132:
	;
	goto L133
L133:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+89)))
	if v480 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+148)) = v432
	goto L136
L135:
	;
	goto L136
L136:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+90)))
	if v484 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v432
	goto L139
L138:
	;
	goto L139
L139:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+91)))
	if v488 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v432
	goto L142
L141:
	;
	goto L142
L142:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+92)))
	if v492 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v432
	goto L145
L144:
	;
	goto L145
L145:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+93)))
	if v496 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = v432
	goto L148
L147:
	;
	goto L148
L148:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+94)))
	if v500 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+168)) = v432
	goto L151
L150:
	;
	goto L151
L151:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+95)))
	if v504 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+172)) = v432
	goto L154
L153:
	;
	goto L154
L154:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+96)))
	if v508 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v432
	goto L157
L156:
	;
	goto L157
L157:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+97)))
	if v512 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v432
	goto L160
L159:
	;
	goto L160
L160:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	F_list_sort(m, v516, int32(1825))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	v521 = v17 + int32(188)
	v522 = F_strlen(m, v27)
	mBase = m.M
	F_initStringInfoExt(m, v521, v522+int32(1))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	if v527 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
	v588 = F_guc_strdup(m, int32(15), v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L172
	}
L164:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v530 <= int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	F_appendStringInfoString(m, v521, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	v537 = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v538 <= v537 {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v544 = v537
	goto L168
L168:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v544<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v559
	F_appendStringInfo(m, v17+int32(188), int32(_a_F_check_log_min_messages_8), v17+int32(16))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L2
	} else {
		goto L170
	}
L169:
	;
	goto L163
L170:
	;
	v569 = v544 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v569 < v570 {
		v544 = v569
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	if v588 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
	F_pfree(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	v651 = int32(0)
	goto L1
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_min_messages[2])) = v616
	F_bms_free(m, v27)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L2
	} else {
		goto L176
	}
L176:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	F_list_free(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L2
	} else {
		goto L177
	}
L177:
	;
	v651 = v608
	goto L1
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v588
	F_bms_free(m, v27)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	F_list_free(m, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
	F_pfree(m, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	v637 = F_guc_malloc(m, int32(72))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L2
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v637
	if v637 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v651 = int32(0)
	goto L1
L184:
	;
	goto L185
L185:
	;
	base.MemoryCopy(m, v637, v17+int32(112), int32(72))
	v651 = int32(1)
	goto L1
}
func F_check_log_timezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_pg_tzset(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v51 = v4
			m.G0 = v8 + int32(16)
			return v51
		} else {
			v17 = F_pg_tz_acceptable(m, v11)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[1])) = v22
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
					v29 = F_format_elog_string(m, int32(_a_F_check_log_timezone_0), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[2])) = v29
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[1])) = v33
						v39 = F_format_elog_string(m, int32(_a_F_check_log_timezone_1), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_check_log_timezone[3])) = v39
							v51 = v4
							m.G0 = v8 + int32(16)
							return v51
						}
					}
				} else {
					v43 = F_guc_malloc(m, int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
						if v43 == int32(0) {
							v51 = v4
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v11
							v51 = int32(1)
						}
						m.G0 = v8 + int32(16)
						return v51
					}
				}
			}
		}
	}
}
