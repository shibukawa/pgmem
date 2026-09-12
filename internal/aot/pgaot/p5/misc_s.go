package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_SB_IMatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v729 int32
	_ = v729
	v6 = int32(0)
	if l3 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = base.B2i32(int32(0) < l1)
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v16 != int32(37) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	if l1 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v729
L7:
	;
	if v685 != 0 {
		v729 = v6
		goto L6
	} else {
		goto L275
	}
L8:
	;
	v682 = l2
	v685 = v22
	v686 = l3
	goto L7
L9:
	;
	goto L10
L10:
	;
	if l3 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v682 = l2
	v685 = v22
	v686 = l3
	goto L7
L12:
	;
	goto L13
L13:
	;
	v31 = l0
	v32 = l1
	v33 = l2
	v34 = l3
	goto L14
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	switch v44 - int32(92) {
	case 0:
		goto L24
	case 1, 2:
		goto L22
	case 3:
		v662 = v33
		v663 = v34
		goto L16
	default:
		goto L25
	}
L15:
	;
	v682 = v671
	v685 = v667
	v686 = v669
	goto L7
L16:
	;
	v666 = int32(1)
	v667 = base.B2i32(v666 < v32)
	v669 = v663 - v666
	v671 = v662 + v666
	if v32 < int32(2) {
		v682 = v671
		v685 = v667
		v686 = v669
		goto L7
	} else {
		goto L273
	}
L17:
	;
	v655 = F_pg_strncoll(m, v354, v459, v31, v32, l4)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L268
	}
L18:
	;
	v650 = F_pg_strncoll(m, v33, v642-v33, v31, v32, l4)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L267
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L263
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L259
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L255
	}
L22:
	;
	if l4 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L23:
	;
	v131 = v31
	v132 = v32
	v133 = v33
	v134 = v34
	goto L69
L24:
	;
	if v34 <= int32(1) {
		goto L21
	} else {
		goto L28
	}
L25:
	;
	if v44 != int32(37) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v49 = int32(1)
	if base.Ui32(v34) <= base.Ui32(v49) {
		v729 = v49
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v54 = int32(1)
	v55 = v33 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v57 == v54 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v93 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L30:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v69 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v89 = v68
	goto L29
L34:
	;
	v68 = v56 | int32(32)
	goto L36
L35:
	;
	v68 = v56
	goto L36
L36:
	;
	goto L33
L37:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v56-int32(65)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v89 = v80
	goto L29
L41:
	;
	v80 = v56 | int32(32)
	goto L43
L42:
	;
	v80 = v56
	goto L43
L43:
	;
	goto L40
L44:
	;
	v89 = v88
	goto L29
L45:
	;
	v88 = v56 | int32(32)
	goto L47
L46:
	;
	v88 = v56
	goto L47
L47:
	;
	goto L44
L48:
	;
	if v89&int32(255) != v125&int32(255) {
		v729 = v6
		goto L6
	} else {
		goto L67
	}
L49:
	;
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v105 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v125 = v104
	goto L48
L53:
	;
	v104 = v90 | int32(32)
	goto L55
L54:
	;
	v104 = v90
	goto L55
L55:
	;
	goto L52
L56:
	;
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v90-int32(65)) < base.Ui32(int32(26)) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v125 = v116
	goto L48
L60:
	;
	v116 = v90 | int32(32)
	goto L62
L61:
	;
	v116 = v90
	goto L62
L62:
	;
	goto L59
L63:
	;
	v125 = v124
	goto L48
L64:
	;
	v124 = v90 | int32(32)
	goto L66
L65:
	;
	v124 = v90
	goto L66
L66:
	;
	goto L63
L67:
	;
	v662 = v55
	v663 = v34 - int32(1)
	goto L16
L68:
	;
	if v132 <= int32(0) {
		goto L118
	} else {
		goto L119
	}
L69:
	;
	v144 = int32(1)
	v145 = v134 - v144
	v147 = v133 + v144
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	switch v148 - int32(92) {
	case 0:
		goto L74
	case 1, 2:
		goto L71
	case 3:
		goto L75
	default:
		goto L73
	}
L70:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v200 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L71:
	;
	goto L70
L72:
	;
	if base.Ui32(int32(2)) < base.Ui32(v134) {
		v131 = v196
		v132 = v197
		v133 = v147
		v134 = v145
		goto L69
	} else {
		goto L99
	}
L73:
	;
	if v148 != int32(37) {
		goto L71
	} else {
		goto L98
	}
L74:
	;
	if v145 == int32(1) {
		goto L20
	} else {
		goto L79
	}
L75:
	;
	if v132 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	return int32(-1)
L77:
	;
	goto L78
L78:
	;
	v155 = int32(1)
	v196 = v131 + v155
	v197 = v132 - v155
	goto L72
L79:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+2)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v162 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if base.Ui32((v161-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v174 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v233 = v173
	goto L68
L84:
	;
	v173 = v161 | int32(32)
	goto L86
L85:
	;
	v173 = v161
	goto L86
L86:
	;
	goto L83
L87:
	;
	if base.Ui32((v161-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v161-int32(65)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v233 = v185
	goto L68
L91:
	;
	v185 = v161 | int32(32)
	goto L93
L92:
	;
	v185 = v161
	goto L93
L93:
	;
	goto L90
L94:
	;
	v233 = v193
	goto L68
L95:
	;
	v193 = v161 | int32(32)
	goto L97
L96:
	;
	v193 = v161
	goto L97
L97:
	;
	goto L94
L98:
	;
	v196 = v131
	v197 = v132
	goto L72
L99:
	;
	v729 = v49
	goto L6
L100:
	;
	if base.Ui32((v148-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v212 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v233 = v211
	goto L68
L104:
	;
	v211 = v148 | int32(32)
	goto L106
L105:
	;
	v211 = v148
	goto L106
L106:
	;
	goto L103
L107:
	;
	if base.Ui32((v148-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	if base.Ui32(v148-int32(65)) < base.Ui32(int32(26)) {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v233 = v223
	goto L68
L111:
	;
	v223 = v148 | int32(32)
	goto L113
L112:
	;
	v223 = v148
	goto L113
L113:
	;
	goto L110
L114:
	;
	v233 = v231
	goto L68
L115:
	;
	v231 = v148 | int32(32)
	goto L117
L116:
	;
	v231 = v148
	goto L117
L117:
	;
	goto L114
L118:
	;
	return int32(-1)
L119:
	;
	goto L120
L120:
	;
	v240 = v131
	v241 = v132
	goto L121
L121:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v254 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v729 = int32(-1)
	goto L6
L123:
	;
	v295 = int32(1)
	if v295 < v241 {
		v240 = v240 + v295
		v241 = v241 - v295
		goto L121
	} else {
		goto L149
	}
L124:
	;
	if v233&int32(255) != v286&int32(255) {
		goto L143
	} else {
		goto L144
	}
L125:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v266 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v286 = v265
	goto L124
L129:
	;
	v265 = v253 | int32(32)
	goto L131
L130:
	;
	v265 = v253
	goto L131
L131:
	;
	goto L128
L132:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L134
L134:
	;
	if base.Ui32(v253-int32(65)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L135:
	;
	v286 = v277
	goto L124
L136:
	;
	v277 = v253 | int32(32)
	goto L138
L137:
	;
	v277 = v253
	goto L138
L138:
	;
	goto L135
L139:
	;
	v286 = v285
	goto L124
L140:
	;
	v285 = v253 | int32(32)
	goto L142
L141:
	;
	v285 = v253
	goto L142
L142:
	;
	goto L139
L143:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v290 != 0 {
		goto L123
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v291 = F_SB_IMatchText(m, v240, v241, v147, v145, l4)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	if v291 != 0 {
		v729 = v291
		goto L6
	} else {
		goto L148
	}
L148:
	;
	goto L123
L149:
	;
	goto L122
L150:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v516 == int32(1) {
		goto L217
	} else {
		goto L218
	}
L151:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v303 != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	if v34 == int32(0) {
		v642 = v33
		goto L18
	} else {
		goto L153
	}
L153:
	;
	v310 = v34
	v312 = int32(0)
	v313 = v33
	goto L157
L154:
	;
	v475 = v32
	v481 = v31
	goto L198
L155:
	;
	v463 = v33
	v464 = v310
	v466 = v313 - v33
	v467 = v313
	v469 = v6
	goto L154
L156:
	;
	v353 = v351 - v33
	v354 = F_palloc(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L168
	}
L157:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	switch v320 - int32(92) {
	case 0:
		goto L160
	case 1, 2:
		v336 = v310
		v337 = v312
		v338 = v313
		goto L159
	case 3:
		goto L161
	default:
		goto L162
	}
L158:
	;
	v343 = int32(1)
	if v337&v343 == int32(0) {
		v642 = v340
		goto L18
	} else {
		goto L167
	}
L159:
	;
	v339 = int32(1)
	v340 = v338 + v339
	v342 = v336 - v339
	if v342 != 0 {
		v310 = v342
		v312 = v337
		v313 = v340
		goto L157
	} else {
		goto L166
	}
L160:
	;
	v330 = v310 - int32(1)
	if v330 == int32(0) {
		goto L19
	} else {
		goto L165
	}
L161:
	;
	if v312&int32(1) == int32(0) {
		goto L155
	} else {
		goto L164
	}
L162:
	;
	if v320 != int32(37) {
		v336 = v310
		v337 = v312
		v338 = v313
		goto L159
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v349 = v310
	v351 = v313
	v352 = v6
	goto L156
L165:
	;
	v333 = int32(1)
	v336 = v330
	v337 = v333
	v338 = v313 + v333
	goto L159
L166:
	;
	goto L158
L167:
	;
	v349 = int32(0)
	v351 = v340
	v352 = v343
	goto L156
L168:
	;
	if base.Ui32(v351) <= base.Ui32(v33) {
		v451 = v354
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v459 = v451 - v354
	if v352 != 0 {
		goto L17
	} else {
		goto L197
	}
L170:
	;
	v358 = v353 & int32(3)
	if v358 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v33-v351) {
		v451 = v391
		goto L169
	} else {
		goto L181
	}
L172:
	;
	v391 = v354
	v393 = v33
	goto L171
L173:
	;
	goto L174
L174:
	;
	v366 = v354
	v368 = v33
	v370 = v6
	goto L175
L175:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v374 != int32(92) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v391 = v380
	v393 = v382
	goto L171
L177:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v374)
	v380 = v366 + int32(1)
	goto L179
L178:
	;
	v380 = v366
	goto L179
L179:
	;
	v381 = int32(1)
	v382 = v368 + v381
	v384 = v370 + v381
	if v384 != v358 {
		v366 = v380
		v368 = v382
		v370 = v384
		goto L175
	} else {
		goto L180
	}
L180:
	;
	goto L176
L181:
	;
	v407 = v391
	v409 = v393
	goto L182
L182:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v415 != int32(92) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v451 = v442
	goto L169
L184:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v407))) = uint8(v415)
	v421 = v407 + int32(1)
	goto L186
L185:
	;
	v421 = v407
	goto L186
L186:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	if v422 != int32(92) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v421))) = uint8(v422)
	v428 = v421 + int32(1)
	goto L189
L188:
	;
	v428 = v421
	goto L189
L189:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+2)))
	if v429 != int32(92) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v429)
	v435 = v428 + int32(1)
	goto L192
L191:
	;
	v435 = v428
	goto L192
L192:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+3)))
	if v436 != int32(92) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v435))) = uint8(v436)
	v442 = v435 + int32(1)
	goto L195
L194:
	;
	v442 = v435
	goto L195
L195:
	;
	v444 = v409 + int32(4)
	if v444 != v351 {
		v407 = v442
		v409 = v444
		goto L182
	} else {
		goto L196
	}
L196:
	;
	goto L183
L197:
	;
	v463 = v354
	v464 = v349
	v466 = v459
	v467 = v351
	v469 = v354
	goto L154
L198:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v488 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v492 = F_pg_strncoll(m, v463, v466, v31, v481-v31, l4)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L205
	}
L203:
	;
	goto L202
L204:
	;
	if v475 != 0 {
		goto L211
	} else {
		goto L212
	}
L205:
	;
	if v492 != 0 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v494 = F_SB_IMatchText(m, v481, v475, v467, v464, l4)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	if v494 != int32(1) {
		goto L204
	} else {
		goto L208
	}
L208:
	;
	if v469 == int32(0) {
		v729 = int32(1)
		goto L6
	} else {
		goto L209
	}
L209:
	;
	F_pfree(m, v469)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	return int32(1)
L211:
	;
	v505 = int32(1)
	v475 = v475 - v505
	v481 = v481 + v505
	goto L198
L212:
	;
	v509 = int32(0)
	if v469 == v509 {
		v729 = v509
		goto L6
	} else {
		goto L214
	}
L214:
	;
	F_pfree(m, v469)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L215
	}
L215:
	;
	return int32(0)
L216:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v552 == int32(1) {
		goto L236
	} else {
		goto L237
	}
L217:
	;
	if base.Ui32((v44-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v528 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v548 = v527
	goto L216
L221:
	;
	v527 = v44 | int32(32)
	goto L223
L222:
	;
	v527 = v44
	goto L223
L223:
	;
	goto L220
L224:
	;
	if base.Ui32((v44-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	goto L226
L226:
	;
	if base.Ui32(v44-int32(65)) < base.Ui32(int32(26)) {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	v548 = v539
	goto L216
L228:
	;
	v539 = v44 | int32(32)
	goto L230
L229:
	;
	v539 = v44
	goto L230
L230:
	;
	goto L227
L231:
	;
	v548 = v547
	goto L216
L232:
	;
	v547 = v44 | int32(32)
	goto L234
L233:
	;
	v547 = v44
	goto L234
L234:
	;
	goto L231
L235:
	;
	if v548&int32(255) == v584&int32(255) {
		v662 = v33
		v663 = v34
		goto L16
	} else {
		goto L254
	}
L236:
	;
	if base.Ui32((v549-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	goto L238
L238:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v564 == int32(1) {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	v584 = v563
	goto L235
L240:
	;
	v563 = v549 | int32(32)
	goto L242
L241:
	;
	v563 = v549
	goto L242
L242:
	;
	goto L239
L243:
	;
	if base.Ui32((v549-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	goto L245
L245:
	;
	if base.Ui32(v549-int32(65)) < base.Ui32(int32(26)) {
		goto L251
	} else {
		goto L252
	}
L246:
	;
	v584 = v575
	goto L235
L247:
	;
	v575 = v549 | int32(32)
	goto L249
L248:
	;
	v575 = v549
	goto L249
L249:
	;
	goto L246
L250:
	;
	v584 = v583
	goto L235
L251:
	;
	v583 = v549 | int32(32)
	goto L253
L252:
	;
	v583 = v549
	goto L253
L253:
	;
	goto L250
L254:
	;
	v729 = v6
	goto L6
L255:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	F_errmsg(m, int32(217190), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(497926), int32(107), int32(63989))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	F_errmsg(m, int32(217190), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(497926), int32(169), int32(63989))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(217190), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(497926), int32(237), int32(63989))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	return base.B2i32(v650 == int32(0))
L268:
	;
	if v354 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	F_pfree(m, v354)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	return base.B2i32(v655 == int32(0))
L272:
	;
	goto L271
L273:
	;
	v674 = int32(1)
	if v674 < v663 {
		v31 = v31 + v674
		v32 = v32 - v674
		v33 = v671
		v34 = v669
		goto L14
	} else {
		goto L274
	}
L274:
	;
	goto L15
L275:
	;
	v693 = int32(1)
	if v686 <= int32(0) {
		v729 = v693
		goto L6
	} else {
		goto L276
	}
L276:
	;
	v698 = v682
	v702 = v686
	goto L277
L277:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	if v709 != int32(37) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v729 = v693
	goto L6
L279:
	;
	return int32(-1)
L280:
	;
	goto L281
L281:
	;
	v714 = int32(1)
	if v714 < v702 {
		v698 = v698 + v714
		v702 = v702 - v714
		goto L277
	} else {
		goto L282
	}
L282:
	;
	goto L278
}
func F_SearchNamedReplicationSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v3 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v12 = F_LWLockAcquire(m, v8+int32(4736), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	if v17 <= int32(0) {
		v69 = v3
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[560]))
	v25 = v3
	goto L8
L8:
	;
	v30 = v21 + v25*int32(288)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v31 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v69 = int32(0)
	goto L6
L10:
	;
	v35 = v30 + int32(24)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v39 == int32(0) {
		v58 = v38
		v59 = v39
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v64 = v25 + int32(1)
	if v64 != v17 {
		v25 = v64
		goto L8
	} else {
		goto L22
	}
L13:
	;
	if v59-v58 == int32(0) {
		v69 = v30
		goto L6
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v38 != v39 {
		v58 = v38
		v59 = v39
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = l0
	v44 = v35
	goto L17
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v47
		v59 = v48
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v58 = v47
	v59 = v48
	goto L14
L19:
	;
	v51 = int32(1)
	if v47 == v48 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L12
L22:
	;
	goto L9
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v74+int32(4736))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	return v69
L26:
	;
	goto L25
}
func F_SetUserIdAndSecContext(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[240])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[239])) = l0
	return
}
func F_ShmemAlloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[680]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	if v13 != 0 {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[680]))
		F_s_lock(m, v21, int32(497137), int32(208), int32(32540))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[681]))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
			v32 = v31 + (l0+int32(127))&int32(-128)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			if base.Ui32(v33) < base.Ui32(v32) {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[680]))
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(677708), v9)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(497137), int32(162), int32(489057))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[682]))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v32
				v43 = *(*int32)(unsafe.Add(mBase, _consts[680]))
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
				v46 = v31 + v40
				if v46 != 0 {
					m.G0 = v9 + int32(16)
					return v46
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(8389))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(677708), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497137), int32(162), int32(489057))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _consts[681]))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
		v32 = v31 + (l0+int32(127))&int32(-128)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
		if base.Ui32(v33) < base.Ui32(v32) {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[680]))
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg(m, int32(677708), v9)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497137), int32(162), int32(489057))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, _consts[682]))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v32
			v43 = *(*int32)(unsafe.Add(mBase, _consts[680]))
			*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
			v46 = v31 + v40
			if v46 != 0 {
				m.G0 = v9 + int32(16)
				return v46
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(677708), v9)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(497137), int32(162), int32(489057))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v19 = F_LWLockAcquire(m, v15+int32(128), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[683]))
		if v24 == int32(0) {
			v28 = *(*int32)(unsafe.Add(mBase, _consts[681]))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
			if v30 == int32(1) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
				v92 = v33
				v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				F_LWLockRelease(m, v98+int32(128))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					m.G0 = v12 + int32(48)
					return v92
				}
			} else {
				v36 = F_ShmemAlloc(m, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v36
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v39)
					v92 = v36
					v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					F_LWLockRelease(m, v98+int32(128))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(48)
						return v92
					}
				}
			}
		} else {
			v42 = F_hash_search(m, v24, l0, int32(3), l2)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v108 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					F_LWLockRelease(m, v108+int32(128))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
								F_errmsg(m, int32(713702), v12)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(497137), int32(437), int32(109077))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
					if v46 == int32(1) {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
						if v49 != l1 {
							v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							F_LWLockRelease(m, v130+int32(128))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v139
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
									F_errmsg(m, int32(37027), v12+int32(16))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497137), int32(453), int32(109077))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
							v92 = v51
							v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							F_LWLockRelease(m, v98+int32(128))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(48)
								return v92
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[680]))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(1)
						v60 = (l1 + int32(127)) & int32(-128)
						if v54 != 0 {
							v62 = *(*int32)(unsafe.Add(mBase, _consts[680]))
							F_s_lock(m, v62, int32(497137), int32(208), int32(32540))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, _consts[681]))
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
								v71 = v70 + v60
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
								if base.Ui32(v72) < base.Ui32(v71) {
									v75 = *(*int32)(unsafe.Add(mBase, _consts[680]))
									*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(0)
									v156 = *(*int32)(unsafe.Add(mBase, _consts[683]))
									v159 = F_hash_search(m, v156, l0, int32(2), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										F_LWLockRelease(m, v162+int32(128))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return int32(0)
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(8389))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
													F_errmsg(m, int32(677751), v12+int32(32))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(497137), int32(472), int32(109077))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, _consts[682]))
									*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v71
									v82 = *(*int32)(unsafe.Add(mBase, _consts[680]))
									v83 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = v83
									v85 = v70 + v79
									if v85 == v83 {
										v156 = *(*int32)(unsafe.Add(mBase, _consts[683]))
										v159 = F_hash_search(m, v156, l0, int32(2), int32(0))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											v162 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											F_LWLockRelease(m, v162+int32(128))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(8389))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
														F_errmsg(m, int32(677751), v12+int32(32))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(497137), int32(472), int32(109077))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v60
										*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v85
										v92 = v85
										v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										F_LWLockRelease(m, v98+int32(128))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(48)
											return v92
										}
									}
								}
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, _consts[681]))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
							v71 = v70 + v60
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
							if base.Ui32(v72) < base.Ui32(v71) {
								v75 = *(*int32)(unsafe.Add(mBase, _consts[680]))
								*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(0)
								v156 = *(*int32)(unsafe.Add(mBase, _consts[683]))
								v159 = F_hash_search(m, v156, l0, int32(2), int32(0))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									v162 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									F_LWLockRelease(m, v162+int32(128))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int32(0)
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(8389))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
												F_errmsg(m, int32(677751), v12+int32(32))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(497137), int32(472), int32(109077))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								}
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _consts[682]))
								*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v71
								v82 = *(*int32)(unsafe.Add(mBase, _consts[680]))
								v83 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = v83
								v85 = v70 + v79
								if v85 == v83 {
									v156 = *(*int32)(unsafe.Add(mBase, _consts[683]))
									v159 = F_hash_search(m, v156, l0, int32(2), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										F_LWLockRelease(m, v162+int32(128))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return int32(0)
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(8389))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
													F_errmsg(m, int32(677751), v12+int32(32))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(497137), int32(472), int32(109077))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v60
									*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v85
									v92 = v85
									v98 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									F_LWLockRelease(m, v98+int32(128))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(48)
										return v92
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_SimpleLruReadPage_ReadOnly(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v13 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = base.I64_rem_s(l1, v13)
	v15 = base.I32_wrap_i64(v14)
	v18 = v12 + v15<<(uint(int32(7))%32)
	v20 = F_LWLockAcquire(m, v18, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = v15 << (uint(int32(4)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v33 = v25
	goto L4
L3:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v86 = int32(2)
	v88 = v83 + v82>>(uint(int32(4))%32)<<(uint(v86)%32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v91 = v82 << (uint(v86) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)))
	if v89 != v94 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28+v33<<(uint(int32(2))%32))))
	if v42 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_LWLockRelease(m, v18)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v54 = v33 | int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28+v54<<(uint(int32(2))%32))))
	if v58 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v42 == int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v47+v33<<(uint(int32(3))%32))))
	if v51 != l1 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v82 = v33
	goto L3
L10:
	;
	if v33 != v25|int32(14) {
		v33 = v33 + int32(2)
		goto L4
	} else {
		goto L14
	}
L11:
	;
	if v58 == int32(1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v63+v54<<(uint(int32(3))%32))))
	if v67 == l1 {
		v82 = v54
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	v75 = F_LWLockAcquire(m, v18, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v78 = F_SimpleLruReadPage(m, l0, l1, int32(1), l2)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	return v78
L18:
	;
	v97 = v89 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v91))) = v97
	goto L20
L19:
	;
	goto L20
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[89])) = uint8(v105)
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v105)
	v111 = v103 << (uint(int32(6)) % 32)
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[91])))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[91]))) = v114 + int64(1)
	goto L21
L21:
	;
	return v82
}
func F_StatisticsObjIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(64), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v148
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v148 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(41842), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(499480), int32(2659), int32(64459))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v87 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = v3
	goto L14
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	if v108 != v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == v40 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v120 = v96 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v120 < v121 {
		v96 = v120
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v138 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v115 = int32(0)
	v117 = F_SearchSysCacheExists(m, int32(63), v39+int32(8), v108, v115, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v117 != 0 {
		v138 = v3
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v148 = v138
	goto L1
}
func F_StorePreparedStatement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v3 = l2
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = *(*int64)(unsafe.Add(mBase, _consts[291]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[292]))
	if v14 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(343597383744)
		v25 = F_hash_create(m, int32(168133), int32(32), v7+int32(-48), int32(24))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[292])) = v25
			v28 = v25
			v32 = F_hash_search(m, v28, l0, int32(1), v7+int32(-48))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
				if v34 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errcode(m, int32(84017284))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(116229), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(498746), int32(415), int32(95929))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v12
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l1
					F_SaveCachedPlan(m, l1)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						m.G0 = v9 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v28 = v14
		v32 = F_hash_search(m, v28, l0, int32(1), v7+int32(-48))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
			if v34 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errcode(m, int32(84017284))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(116229), v9)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(498746), int32(415), int32(95929))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v12
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+68)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l1
				F_SaveCachedPlan(m, l1)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					m.G0 = v9 - int32(-64)
					return
				}
			}
		}
	}
}
func F_StoreSingleInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_table_open(m, int32(2611), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v14
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
		v26 = F_heap_form_tuple(m, v21, v8+int32(16), v8+int32(12))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_CatalogTupleInsert(m, v12, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_pfree(m, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_sequence_close(m, v12, int32(3))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_SwitchBackToLocalLatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = int32(4513304)
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v3
	v8 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	if v8 != 0 {
		v9 = int32(1)
		F_ModifyWaitEvent(m, v8, v9, v9, int32(4513304))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[83]))
			v16 = v15
			F_SetLatch(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v16 = v3
		F_SetLatch(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F___strxfrm_l(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v4 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(v4) < base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if (l1^l0)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	return v4
L4:
	;
	goto L3
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v59)
	if v59&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v58 = l1
	v59 = v11
	v60 = l0
	goto L6
L8:
	;
	goto L9
L9:
	;
	if l1&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v15 = l1
	v17 = l0
	goto L13
L11:
	;
	v29 = l1
	v31 = l0
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v36 = int32(-2139062144)
	if (int32(16843008)-v33|v33)&v36 != v36 {
		v58 = v29
		v59 = v33
		v60 = v31
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
	if v18 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v29 = v25
	v31 = v23
	goto L12
L15:
	;
	v22 = int32(1)
	v23 = v17 + v22
	v25 = v15 + v22
	if v25&int32(3) != 0 {
		v15 = v25
		v17 = v23
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v41 = v29
	v42 = v33
	v43 = v31
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v42
	v45 = int32(4)
	v46 = v43 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v49 = v41 + v45
	v53 = int32(-2139062144)
	if (v47|(int32(16843008)-v47))&v53 == v53 {
		v41 = v49
		v42 = v47
		v43 = v46
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v58 = v49
	v59 = v47
	v60 = v46
	goto L6
L20:
	;
	goto L19
L21:
	;
	v67 = v58
	v69 = v60
	goto L22
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)) = uint8(v70)
	v72 = int32(1)
	if v70 != 0 {
		v67 = v67 + v72
		v69 = v69 + v72
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
}
func F___subtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	F___addtf3(m, v9, l1, l2, l3, l4^int64(-9223372036854775807-1))
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	m.G0 = v9 + v8
	return
}
func F_scalarineqsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 float64
	_ = v49
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v109 float64
	_ = v109
	var v112 float32
	_ = v112
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v137 float64
	_ = v137
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v20 == int32(0) {
		v23 = float64(0.3333333333333333)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
		if v24 == int32(0) {
			v137 = v23
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			if v27 != int32(6) {
				v137 = v23
			} else {
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
				if v30 != int32(65535) {
					v137 = v23
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
					if v34 == int32(0) {
						v137 = float64(1)
					} else {
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v33)+120))
						v40 = base.F64_add(base.F64_convert_i32_u(v34), float64(-0.5))
						v41 = base.F64_div(v37, v40)
						v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+2)))
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6))))
						v49 = base.F64_convert_i32_u(v44 | v45<<(uint(int32(16))%32))
						if base.F64_ge(v49, base.F64_convert_i32_u(v34-int32(1))) != 0 {
							v54 = base.F64_mul(v41, float64(0.5))
						} else {
							v54 = v41
						}
						if base.F64_gt(v54, float64(0)) != 0 {
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
							v59 = base.F64_div(base.F64_convert_i32_u(v57), v54)
							v60 = float64(1)
							if base.F64_lt(v59, v60) != 0 {
								v63 = v59
							} else {
								v63 = v60
							}
							v66 = base.F64_add(v63, v49)
						} else {
							v66 = v49
						}
						v67 = base.F64_div(v66, v40)
						if l2 != l3 {
							v76 = v67
						} else {
							if base.F64_ge(v37, float64(1)) == int32(0) {
								v76 = v67
							} else {
								v76 = base.F64_add(v67, base.F64_div(float64(-1), v37))
							}
						}
						if l2 != 0 {
							v80 = base.F64_sub(float64(1), v76)
						} else {
							v80 = v76
						}
						if base.F64_lt(v80, float64(0)) != 0 {
							v137 = float64(0)
						} else {
							if base.F64_gt(v80, float64(1)) != 0 {
								v137 = float64(1)
							} else {
								v137 = v80
							}
						}
					}
				}
			}
		}
		m.G0 = v18 + int32(48)
		return v137
	} else {
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
		v87 = F_get_opcode(m, l1)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return float64(0)
		} else {
			F_fmgr_info(m, v87, v18+int32(20))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return float64(0)
			} else {
				v100 = F_mcv_selectivity(m, l5, v18+int32(20), l4, l6, int32(1), v18+int32(8))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return float64(0)
				} else {
					v104 = F_ineq_histogram_selectivity(m, l0, l5, l1, v18+int32(20), l2, l3, l4, l6, l7)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return float64(0)
					} else {
						if base.F64_ge(v104, float64(0)) != 0 {
							v109 = v104
						} else {
							v109 = float64(0.5)
						}
						v112 = *(*float32)(unsafe.Add(mBase, uint32(v85+v86)+8))
						v115 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
						v118 = base.F64_add(v100, base.F64_mul(v109, base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v112)), v115)))
						if base.F64_lt(v118, float64(0)) != 0 {
							v137 = float64(0)
						} else {
							if base.F64_gt(v118, float64(1)) == int32(0) {
								v137 = v118
							} else {
								v137 = float64(1)
							}
						}
						m.G0 = v18 + int32(48)
						return v137
					}
				}
			}
		}
	}
}
func F_scalbn(m *base.Module, l0 float64, l1 int32) float64 {
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	if int32(1024) <= l1 {
		v6 = base.F64_mul(l0, float64(8.98846567431158e+307))
		if base.Ui32(l1) < base.Ui32(int32(2047)) {
			v35 = v6
			v36 = l1 - int32(1023)
		} else {
			v13 = int32(3069)
			if base.Ui32(v13) <= base.Ui32(l1) {
				v16 = v13
			} else {
				v16 = l1
			}
			v35 = base.F64_mul(v6, float64(8.98846567431158e+307))
			v36 = v16 - int32(2046)
		}
	} else {
		if int32(-1023) < l1 {
			v35 = l0
			v36 = l1
		} else {
			v22 = base.F64_mul(l0, float64(2.004168360008973e-292))
			if base.Ui32(int32(-1992)) < base.Ui32(l1) {
				v35 = v22
				v36 = l1 + int32(969)
			} else {
				v29 = int32(-2960)
				if base.Ui32(l1) <= base.Ui32(v29) {
					v32 = v29
				} else {
					v32 = l1
				}
				v35 = base.F64_mul(v22, float64(2.004168360008973e-292))
				v36 = v32 + int32(1938)
			}
		}
	}
	return base.F64_mul(v35, base.F64_reinterpret_i64(base.I64_extend_i32_u(v36+int32(1023))<<(uint(int64(52))%64)))
}
func F_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v7 = F_strlen(m, l0)
	mBase = m.M
	v9 = F_palloc(m, int32(100))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(48)
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(295595), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(314521), int32(1258), int32(99978))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v36 = F__emscripten_memset_bulkmem(m, v9+int32(4), base.I32_extend8_s(int32(0)), int32(96))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l2
			v41 = *(*int32)(unsafe.Add(mBase, _consts[272]))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v41
			v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v44)
			v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[274])))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)) = uint8(v47)
			v50 = v7 + int32(2)
			v51 = F_palloc(m, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v51
				if v7 != 0 {
					v55 = F__emscripten_memcpy_bulkmem(m, v51, l0, v7)
					mBase = m.M
				} else {
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v59 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v57+v7)+1)) = uint8(v59)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v61+v7))) = uint8(v59)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui32(v50) < base.Ui32(int32(2)) {
					v153 = int32(1024)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v153
					v156 = F_palloc(m, v153)
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
						return v9
					}
				} else {
					v69 = v50 - int32(2)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v69))))
					if v71 != 0 {
						v153 = int32(1024)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v153
						v156 = F_palloc(m, v153)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
							return v9
						}
					} else {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v65-int32(1)))))
						if v75 != 0 {
							v153 = int32(1024)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v153
							v156 = F_palloc(m, v153)
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
								return v9
							}
						} else {
							v77 = F_palloc(m, int32(48))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								if v77 != 0 {
									v79 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v79
									*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v69
									*(*int64)(unsafe.Add(mBase, uint32(v77)+40)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = int64(4294967296)
									*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v69
									*(*int32)(unsafe.Add(mBase, uint32(v77))) = v79
									F_core_yyensure_buffer_stack(m, v9)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94<<(uint(int32(2))%32))))
										if v98 == v77 {
										} else {
											if v98 != 0 {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
												*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v101)
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v105 = int32(2)
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104<<(uint(v105)%32))))
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
												*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v109
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(v105)%32))))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
												*(*int32)(unsafe.Add(mBase, uint32(v116)+16)) = v117
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v121 = v119
												v122 = v120
											} else {
												v121 = v93
												v122 = v94
											}
											v123 = int32(2)
											*(*int32)(unsafe.Add(mBase, uint32(v122<<(uint(v123)%32)+v121))) = v77
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v131 = v127 + v128<<(uint(v123)%32)
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
											v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v133
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
											v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v136
											*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v136
											v139 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v140
											v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v142)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(1)
										}
										v153 = int32(1024)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v153
										v156 = F_palloc(m, v153)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
											return v9
										}
									}
								} else {
									F_yy_fatal_error_2(m, int32(683681))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_schedule_alarm(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	if v11 <= int32(0) {
		m.G0 = v8 + int32(48)
		return
	} else {
		v14 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v14
		v23 = *(*int32)(unsafe.Add(mBase, _consts[1295]))
		if v23 == int32(0) {
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, _consts[1296]))
			if l0 <= v27+int64(10000) {
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1295])) = int32(0)
			}
		}
		v35 = *(*int32)(unsafe.Add(mBase, _consts[1297]))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
		if v36 < l0 {
			v38 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v38
			*(*int32)(unsafe.Add(mBase, _consts[1295])) = v38
			v70 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v70
			v74 = int32(0)
			v75 = v70
		} else {
			v49 = v36 - l0
			if v49 <= int64(0) {
				v61 = int32(0)
				v62 = int32(0)
			} else {
				v53 = int64(1000000)
				v54 = base.I64_div_u_s(v49, v53)
				v61 = base.I32_wrap_i64(v54)
				v62 = base.I32_wrap_i64(v49 - v54*v53)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12)))) = v61
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v62
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			if v65|v66 != 0 {
				v74 = v65
				v75 = v66
			} else {
				v70 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v70
				v74 = int32(0)
				v75 = v70
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[475])) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_s(v74)
		v83 = *(*int32)(unsafe.Add(mBase, _consts[1295]))
		if v83 != 0 {
			v85 = *(*int64)(unsafe.Add(mBase, _consts[1296]))
			if v85 <= v36 {
				m.G0 = v8 + int32(48)
				return
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[1296])) = v36
				*(*int32)(unsafe.Add(mBase, _consts[1295])) = int32(1)
				v94 = F_setitimer(m, v8+int32(16))
				mBase = m.M
				if v94 != 0 {
					v102 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1295])) = v102
					F_errstart_cold(m, int32(22), v102)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(293701), int32(0))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							F_errfinish(m, int32(492522), int32(347), int32(288301))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v8 + int32(48)
					return
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, _consts[1296])) = v36
			*(*int32)(unsafe.Add(mBase, _consts[1295])) = int32(1)
			v94 = F_setitimer(m, v8+int32(16))
			mBase = m.M
			if v94 != 0 {
				v102 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[1295])) = v102
				F_errstart_cold(m, int32(22), v102)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(293701), int32(0))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						F_errfinish(m, int32(492522), int32(347), int32(288301))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(48)
				return
			}
		}
	}
}
func F_search_indexed_tlist_for_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= v6 {
		v250 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L61
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v250
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v26 = l1
	v31 = v18
	goto L4
L4:
	;
	v39 = v26 + int32(12)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v250 = v6
	goto L2
L6:
	;
	v241 = int32(1)
	if base.Ui32(v241) < base.Ui32(v31) {
		v26 = v39
		v31 = v31 - v241
		goto L4
	} else {
		goto L60
	}
L7:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	if v42 != v22&int32(65535) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v45 = F_palloc(m, int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+32)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v55
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v59
	if v22 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = l2
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+8)) = uint16(v234)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v45)+36))
	if v236 == int32(0) {
		v250 = v45
		goto L2
	} else {
		goto L59
	}
L12:
	;
	switch l4 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L15
	default:
		goto L14
	}
L13:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v177 = int32(0)
	if v175 == v177 {
		goto L45
	} else {
		goto L46
	}
L14:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v125 = int32(0)
	v132 = base.B2i32(v123|v124 == v125)
	if v123 == v125 {
		v171 = v132
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v67 = int32(0)
	if v65 == v67 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v120 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v120 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v66 == int32(0) {
		v111 = v67
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v120 = v111
	goto L16
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v77 < v76 {
		v111 = v67
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v79 = int32(1)
	if v76 <= v79 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v79
	goto L25
L24:
	;
	v82 = v76
	goto L25
L25:
	;
	v83 = int32(8)
	v88 = int32(0)
	goto L26
L26:
	;
	v95 = v88 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65+v83+v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v66+v83))))
	v102 = v97 & (v99 ^ int32(-1))
	v104 = base.B2i32(v102 == int32(0))
	if v102 != 0 {
		v111 = v104
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v111 = v104
	goto L20
L28:
	;
	v106 = v88 + int32(1)
	if v106 != v82 {
		v88 = v106
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L11
L31:
	;
	if v171 != 0 {
		goto L11
	} else {
		goto L43
	}
L32:
	;
	goto L31
L33:
	;
	if v124 == int32(0) {
		v171 = v132
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v138 != v139 {
		v171 = int32(0)
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v141 = int32(1)
	if v138 <= v141 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v144 = v141
	goto L38
L37:
	;
	v144 = v138
	goto L38
L38:
	;
	v145 = int32(8)
	v150 = int32(0)
	goto L39
L39:
	;
	v158 = v150 << (uint(int32(2)) % 32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v123+v145+v158)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+(v124+v145))))
	v163 = base.B2i32(v160 == v162)
	if v162 != v160 {
		v171 = v163
		goto L32
	} else {
		goto L41
	}
L40:
	;
	v171 = v163
	goto L32
L41:
	;
	v166 = v150 + int32(1)
	if v166 != v144 {
		v150 = v166
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L1
L44:
	;
	if v230 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L45:
	;
	v230 = int32(1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	if v176 == int32(0) {
		v221 = v177
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v230 = v221
	goto L44
L49:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v187 < v186 {
		v221 = v177
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v189 = int32(1)
	if v186 <= v189 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v192 = v189
	goto L53
L52:
	;
	v192 = v186
	goto L53
L53:
	;
	v193 = int32(8)
	v198 = int32(0)
	goto L54
L54:
	;
	v205 = v198 << (uint(int32(2)) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v175+v193+v205)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+(v176+v193))))
	v212 = v207 & (v209 ^ int32(-1))
	v214 = base.B2i32(v212 == int32(0))
	if v212 != 0 {
		v221 = v214
		goto L48
	} else {
		goto L56
	}
L55:
	;
	v221 = v214
	goto L48
L56:
	;
	v216 = v198 + int32(1)
	if v216 != v192 {
		v198 = v216
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L11
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = l3 + v236
	v250 = v45
	goto L2
L60:
	;
	goto L5
L61:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v267 = F_bmsToString(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v270 = F_bmsToString(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v267
	F_errmsg_internal(m, int32(466553), v16)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(494045), int32(2909), int32(229317))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_search_plan_tree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
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
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	v4 = int32(0)
	if l0 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 - int32(394) {
	case 0, 43:
		v136 = int32(36)
		goto L7
	default:
		v246 = v4
		goto L4
	case 3:
		goto L11
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L10
	case 17:
		goto L8
	}
L4:
	;
	return v246
L5:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v239 != 0 {
		goto L99
	} else {
		goto L100
	}
L6:
	;
	if v226 == int32(0) {
		v246 = v4
		goto L4
	} else {
		goto L98
	}
L7:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0+v136)))
	v139 = int32(0)
	if v138 == v139 {
		goto L65
	} else {
		goto L66
	}
L8:
	;
	v136 = int32(116)
	goto L7
L9:
	;
	v32 = int32(0)
	v34 = v4
	goto L17
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v21 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if int32(0) < v16 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	return int32(0)
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v26 == l1 {
		v235 = l0
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v246 = v4
	goto L4
L17:
	;
	v36 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	if v43 == v36 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v226 = v130
	goto L6
L19:
	;
	if v32 != 0 {
		goto L57
	} else {
		goto L58
	}
L20:
	;
	if v127 != 0 {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	v127 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	switch v52 - int32(394) {
	case 0, 43:
		v90 = int32(36)
		goto L27
	default:
		v117 = v36
		goto L24
	case 3:
		goto L31
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L30
	case 17:
		goto L28
	}
L24:
	;
	v127 = v117
	goto L20
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	if v110 != 0 {
		goto L51
	} else {
		goto L52
	}
L26:
	;
	if v97 == int32(0) {
		v117 = v36
		goto L24
	} else {
		goto L50
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v43+v90)))
	v93 = F_search_plan_tree(m, v92, l1, l2)
	mBase = m.M
	v97 = v93
	goto L26
L28:
	;
	v90 = int32(116)
	goto L27
L29:
	;
	v69 = int32(0)
	v71 = v36
	goto L37
L30:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if v59 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+108))
	if int32(0) < v55 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v127 = int32(0)
	goto L20
L33:
	;
	v127 = int32(0)
	goto L20
L34:
	;
	goto L35
L35:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
	if v63 == l1 {
		v106 = v43
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v117 = v36
	goto L24
L37:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v71<<(uint(int32(2))%32))))
	v81 = F_search_plan_tree(m, v80, l1, l2)
	mBase = m.M
	if v81 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v97 = v84
	goto L26
L39:
	;
	if v69 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v82 = base.B2i32(v69 != int32(0))
	goto L42
L41:
	;
	v82 = int32(5)
	goto L42
L42:
	;
	switch v82 {
	case 0, 5:
		goto L39
	default:
		v117 = v36
		goto L24
	}
L43:
	;
	v83 = v69
	goto L45
L44:
	;
	v83 = v81
	goto L45
L45:
	;
	if v81 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v84 = v83
	goto L48
L47:
	;
	v84 = v69
	goto L48
L48:
	;
	v86 = v71 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v43)+108))
	if v86 < v87 {
		v69 = v84
		v71 = v86
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v106 = v97
	goto L25
L51:
	;
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v111)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v117 = v106
	goto L24
L54:
	;
	v128 = base.B2i32(v32 != v36)
	goto L56
L55:
	;
	v128 = int32(5)
	goto L56
L56:
	;
	switch v128 {
	case 0, 5:
		goto L19
	default:
		v246 = v4
		goto L4
	}
L57:
	;
	v129 = v32
	goto L59
L58:
	;
	v129 = v127
	goto L59
L59:
	;
	if v127 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v130 = v129
	goto L62
L61:
	;
	v130 = v32
	goto L62
L62:
	;
	v132 = v34 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v132 < v133 {
		v32 = v130
		v34 = v132
		goto L17
	} else {
		goto L63
	}
L63:
	;
	goto L18
L64:
	;
	v226 = v222
	goto L6
L65:
	;
	v222 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	switch v147 - int32(394) {
	case 0, 43:
		v185 = int32(36)
		goto L71
	default:
		v212 = v139
		goto L68
	case 3:
		goto L75
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L74
	case 17:
		goto L72
	}
L68:
	;
	v222 = v212
	goto L64
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v138)+52))
	if v205 != 0 {
		goto L95
	} else {
		goto L96
	}
L70:
	;
	if v192 == int32(0) {
		v212 = v139
		goto L68
	} else {
		goto L94
	}
L71:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v138+v185)))
	v188 = F_search_plan_tree(m, v187, l1, l2)
	mBase = m.M
	v192 = v188
	goto L70
L72:
	;
	v185 = int32(116)
	goto L71
L73:
	;
	v164 = int32(0)
	v166 = v139
	goto L81
L74:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v138)+104))
	if v154 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+108))
	if int32(0) < v150 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v222 = int32(0)
	goto L64
L77:
	;
	v222 = int32(0)
	goto L64
L78:
	;
	goto L79
L79:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+56))
	if v158 == l1 {
		v201 = v138
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v212 = v139
	goto L68
L81:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v138)+104))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v166<<(uint(int32(2))%32))))
	v176 = F_search_plan_tree(m, v175, l1, l2)
	mBase = m.M
	if v176 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v192 = v179
	goto L70
L83:
	;
	if v164 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v177 = base.B2i32(v164 != int32(0))
	goto L86
L85:
	;
	v177 = int32(5)
	goto L86
L86:
	;
	switch v177 {
	case 0, 5:
		goto L83
	default:
		v212 = v139
		goto L68
	}
L87:
	;
	v178 = v164
	goto L89
L88:
	;
	v178 = v176
	goto L89
L89:
	;
	if v176 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v179 = v178
	goto L92
L91:
	;
	v179 = v164
	goto L92
L92:
	;
	v181 = v166 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v138)+108))
	if v181 < v182 {
		v164 = v179
		v166 = v181
		goto L81
	} else {
		goto L93
	}
L93:
	;
	goto L82
L94:
	;
	v201 = v192
	goto L69
L95:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v206)
	goto L97
L96:
	;
	goto L97
L97:
	;
	v212 = v201
	goto L68
L98:
	;
	v235 = v226
	goto L5
L99:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v240)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v246 = v235
	goto L4
}
func F_select_rowmark_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(5)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != 0 {
		v29 = v10
		m.G0 = v8 + int32(16)
		return v29
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
		if v12 == int32(102) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v16 = F_GetFdwRoutineByRelId(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
				if v20 == int32(0) {
					v29 = v10
					m.G0 = v8 + int32(16)
					return v29
				} else {
					v23 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v29 = v23
						m.G0 = v8 + int32(16)
						return v29
					}
				}
			}
		} else {
			if base.Ui32(int32(5)) <= base.Ui32(l1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
					F_errmsg_internal(m, int32(474240), v8)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495271), int32(2554), int32(366536))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v29 = int32(4) - l1
				m.G0 = v8 + int32(16)
				return v29
			}
		}
	}
}
func F_self_join_candidates_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v4))
}
func F_sendAuthRequest(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v11 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_pq_beginmessage(m, v8, int32(82))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_enlargeStringInfo(m, v8, int32(4))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v23 = int32(24)
					v25 = int32(65280)
					v27 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = l0<<(uint(v23)%32) | l0&v25<<(uint(v27)%32) | (int32(base.Ui32(l0)>>(uint(v27)%32))&v25 | int32(base.Ui32(l0)>>(uint(v23)%32)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
					if int32(0) < l2 {
						F_pq_sendbytes(m, v8, l1, l2)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							F_pq_endmessage(m, v8)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								switch l0 {
								case 0, 12:
									v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
									if v54 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								default:
									v49 = *(*int32)(unsafe.Add(mBase, _consts[219]))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
									v51 = m.T0[v50].(func(*base.Module) int32)(m)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
										if v54 != 0 {
											F_ProcessInterrupts(m)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						}
					} else {
						F_pq_endmessage(m, v8)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							switch l0 {
							case 0, 12:
								v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								if v54 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							default:
								v49 = *(*int32)(unsafe.Add(mBase, _consts[219]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								v51 = m.T0[v50].(func(*base.Module) int32)(m)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
									if v54 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_pq_beginmessage(m, v8, int32(82))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_enlargeStringInfo(m, v8, int32(4))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v23 = int32(24)
				v25 = int32(65280)
				v27 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = l0<<(uint(v23)%32) | l0&v25<<(uint(v27)%32) | (int32(base.Ui32(l0)>>(uint(v27)%32))&v25 | int32(base.Ui32(l0)>>(uint(v23)%32)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
				if int32(0) < l2 {
					F_pq_sendbytes(m, v8, l1, l2)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_pq_endmessage(m, v8)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							switch l0 {
							case 0, 12:
								v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								if v54 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							default:
								v49 = *(*int32)(unsafe.Add(mBase, _consts[219]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								v51 = m.T0[v50].(func(*base.Module) int32)(m)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
									if v54 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					F_pq_endmessage(m, v8)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						switch l0 {
						case 0, 12:
							v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							if v54 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						default:
							v49 = *(*int32)(unsafe.Add(mBase, _consts[219]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							v51 = m.T0[v50].(func(*base.Module) int32)(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								if v54 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
										return
									}
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_setenv(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v449 int32
	_ = v449
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v107 != 0 {
		goto L53
	} else {
		goto L54
	}
L2:
	;
	return int32(-1)
L3:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(28)
	goto L2
L5:
	;
	goto L10
L6:
	;
	if v95 == l0 {
		goto L4
	} else {
		goto L30
	}
L7:
	;
	goto L6
L8:
	;
	v85 = v80
	goto L26
L9:
	;
	v80 = v72
	goto L8
L10:
	;
	if l0&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v19 = l0
	goto L16
L14:
	;
	v32 = l0
	goto L15
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 != v41 {
		v72 = v32
		goto L9
	} else {
		goto L21
	}
L16:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 == int32(0) {
		v95 = v19
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v32 = v29
	goto L15
L18:
	;
	if int32(61) == v24 {
		v95 = v19
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v29 = v19 + int32(1)
	if v29&int32(3) != 0 {
		v19 = v29
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v47 = v32
	v49 = v38
	goto L22
L22:
	;
	v53 = v49 ^ int32(1027423549)
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 != v56 {
		v72 = v47
		goto L9
	} else {
		goto L24
	}
L23:
	;
	v80 = v62
	goto L8
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v62 = v47 + int32(4)
	v66 = int32(-2139062144)
	if (v60|(int32(16843008)-v60))&v66 == v66 {
		v47 = v62
		v49 = v60
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v87 == int32(0) {
		v95 = v85
		goto L7
	} else {
		goto L28
	}
L27:
	;
	v95 = v85
	goto L7
L28:
	;
	if v87 != int32(61) {
		v85 = v85 + int32(1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v107 = v95 - l0
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v107))))
	if v109 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L4
L32:
	;
	v168 = F_strlen(m, l1)
	mBase = m.M
	v172 = F_emscripten_builtin_malloc(m, v107+v168+int32(2))
	mBase = m.M
	if v172 != 0 {
		goto L1
	} else {
		goto L51
	}
L33:
	;
	v116 = int32(0)
	v121 = F___strchrnul(m, l0, int32(61))
	mBase = m.M
	if l0 == v121 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v163 == int32(0) {
		goto L32
	} else {
		goto L50
	}
L35:
	;
	v163 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v124 = v121 - l0
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v124))))
	if v126 != 0 {
		v156 = v116
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v163 = v156
	goto L34
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v128 == int32(0) {
		v156 = v116
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v131 == int32(0) {
		v156 = v116
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v135 = v128
	v136 = v131
	goto L42
L42:
	;
	v139 = F_strncmp(m, l0, v136, v124)
	mBase = m.M
	if v139 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v156 = v143 + int32(1)
	goto L38
L44:
	;
	goto L43
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v143 = v142 + v124
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v144 == int32(61) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v148 != 0 {
		v135 = v135 + int32(4)
		v136 = v148
		goto L42
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v156 = v116
	goto L38
L50:
	;
	return int32(0)
L51:
	;
	goto L2
L52:
	;
	v180 = v172 + v107
	v181 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v181)
	v183 = int32(1)
	v186 = v168 + v183
	if v186 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v172, l0, v107)
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v190 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	v187 = F__emscripten_memcpy_bulkmem(m, v180+v183, l1, v186)
	mBase = m.M
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	return v449
L61:
	;
	v340 = v336 << (uint(int32(2)) % 32)
	v342 = v340 + int32(8)
	v344 = *(*int32)(unsafe.Add(mBase, _consts[1480]))
	if v344 == v337 {
		goto L106
	} else {
		goto L107
	}
L62:
	;
	v199 = v107 + int32(1)
	v204 = int32(0)
	v205 = v190
	v206 = v194
	goto L68
L63:
	;
	v336 = int32(0)
	v337 = v195
	goto L61
L64:
	;
	v195 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if v194 != 0 {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	v195 = v190
	goto L63
L68:
	;
	if v199 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v336 = v327
	v337 = v332
	goto L61
L70:
	;
	if v250 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L71:
	;
	v250 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v212 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v213 = v172
	v214 = v206
	v215 = v199
	v216 = v212
	goto L78
L75:
	;
	v238 = v206
	v242 = int32(0)
	goto L76
L76:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v250 = v242 - v243
	goto L70
L77:
	;
	v238 = v233
	v242 = v235
	goto L76
L78:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v216 != v218 {
		v233 = v214
		v235 = v216
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v233 = v227
	v235 = int32(0)
	goto L77
L80:
	;
	if v218 == int32(0) {
		v233 = v214
		v235 = v216
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v223 = v215 - int32(1)
	if v223 == int32(0) {
		v233 = v214
		v235 = v216
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v226 = int32(1)
	v227 = v214 + v226
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	if v228 != 0 {
		v213 = v213 + v226
		v214 = v227
		v215 = v223
		v216 = v228
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
L84:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v172
	v261 = *(*int32)(unsafe.Add(mBase, _consts[1481]))
	if v261 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L86
L86:
	;
	v327 = v204 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v328 != 0 {
		v204 = v327
		v205 = v205 + int32(4)
		v206 = v328
		goto L68
	} else {
		goto L103
	}
L87:
	;
	v449 = int32(0)
	goto L60
L88:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
	v265 = v172
	v266 = int32(0)
	goto L91
L89:
	;
	v287 = v172
	goto L90
L90:
	;
	if v287 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	v273 = v263 + v266<<(uint(int32(2))%32)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v253 == v274 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v287 = v282
	goto L90
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v265
	F_emscripten_builtin_free(m, v253)
	mBase = m.M
	goto L87
L94:
	;
	goto L95
L95:
	;
	if v274 != 0 {
		v282 = v265
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v284 = v266 + int32(1)
	if v284 != v261 {
		v265 = v282
		v266 = v284
		goto L91
	} else {
		goto L99
	}
L97:
	;
	if v265 == int32(0) {
		v282 = v265
		goto L96
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v265
	v282 = int32(0)
	goto L96
L99:
	;
	goto L92
L100:
	;
	goto L87
L101:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
	v301 = F_emscripten_builtin_realloc(m, v296, v261<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v301 == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1482])) = v301
	v306 = int32(4680644)
	v308 = *(*int32)(unsafe.Add(mBase, _consts[1481]))
	*(*int32)(unsafe.Add(mBase, _consts[1481])) = v308 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v301+v308<<(uint(int32(2))%32)))) = v287
	goto L100
L103:
	;
	goto L69
L104:
	;
	F_emscripten_builtin_free(m, v172)
	mBase = m.M
	v449 = int32(-1)
	goto L60
L105:
	;
	v360 = v357 + v336<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v357
	*(*int32)(unsafe.Add(mBase, _consts[1480])) = v357
	if v172 != 0 {
		goto L118
	} else {
		goto L119
	}
L106:
	;
	v346 = F_emscripten_builtin_realloc(m, v344, v342)
	mBase = m.M
	if v346 != 0 {
		v357 = v346
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v347 = F_emscripten_builtin_malloc(m, v342)
	mBase = m.M
	if v347 == int32(0) {
		goto L104
	} else {
		goto L110
	}
L109:
	;
	goto L104
L110:
	;
	if v336 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v340 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _consts[1480]))
	F_emscripten_builtin_free(m, v355)
	mBase = m.M
	v357 = v347
	goto L105
L114:
	;
	goto L113
L115:
	;
	v352 = F__emscripten_memcpy_bulkmem(m, v347, v351, v340)
	mBase = m.M
	goto L117
L116:
	;
	goto L117
L117:
	;
	goto L114
L118:
	;
	v368 = int32(0)
	v375 = *(*int32)(unsafe.Add(mBase, _consts[1481]))
	if v375 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	goto L120
L120:
	;
	v449 = int32(0)
	goto L60
L121:
	;
	goto L120
L122:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
	v379 = v172
	v380 = v368
	goto L125
L123:
	;
	v401 = v172
	goto L124
L124:
	;
	if v401 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L125:
	;
	v387 = v377 + v380<<(uint(int32(2))%32)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v368 == v388 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v401 = v396
	goto L124
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v379
	F_emscripten_builtin_free(m, v368)
	mBase = m.M
	goto L121
L128:
	;
	goto L129
L129:
	;
	if v388 != 0 {
		v396 = v379
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v398 = v380 + int32(1)
	if v398 != v375 {
		v379 = v396
		v380 = v398
		goto L125
	} else {
		goto L133
	}
L131:
	;
	if v379 == int32(0) {
		v396 = v379
		goto L130
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v379
	v396 = int32(0)
	goto L130
L133:
	;
	goto L126
L134:
	;
	goto L121
L135:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _consts[1482]))
	v415 = F_emscripten_builtin_realloc(m, v410, v375<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v415 == int32(0) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1482])) = v415
	v420 = int32(4680644)
	v422 = *(*int32)(unsafe.Add(mBase, _consts[1481]))
	*(*int32)(unsafe.Add(mBase, _consts[1481])) = v422 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v415+v422<<(uint(int32(2))%32)))) = v401
	goto L134
}
func F_setval3_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_do_setval(m, v3, v5, base.B2i32(v6 != int32(0)))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_Int64GetDatum(m, v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
func F_sha256_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(3), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_shim_system(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(757269)
	v4 = m.Env.Pgmem_run(m, l0, v2, v2)
	return v4 << (uint(int32(8)) % 32) & int32(65280)
}
func F_show_random_seed(m *base.Module) int32 {
	return int32(395450)
}
func F_sigismember(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v5 = l1 - int32(1)
	if base.Ui32(v5) <= base.Ui32(int32(63)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908))))
		v17 = int32(base.Ui32(v13)>>(uint(v5)%32)) & int32(1)
	} else {
		v17 = int32(0)
	}
	return v17
}
func F_signal_child(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v15) <= base.Ui32(int32(17)) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_consts[470])))
				v25 = v24
			} else {
				v25 = int32(367957)
			}
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[471])))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v34
			F_errmsg_internal(m, int32(478660), v8+int32(32))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_errfinish(m, int32(495164), int32(3457), int32(431219))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v47 = F_kill(m, v10, l1)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if int32(0) <= v47 {
							if base.Ui32(int32(15)) < base.Ui32(l1) {
								m.G0 = v8 + int32(48)
								return
							} else {
								if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
									m.G0 = v8 + int32(48)
									return
								} else {
									v78 = int32(0) - v10
									v79 = F_kill(m, v78, l1)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										if int32(0) <= v79 {
											m.G0 = v8 + int32(48)
											return
										} else {
											v85 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												if v85 == int32(0) {
													m.G0 = v8 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
													F_errmsg_internal(m, int32(295542), v8)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_errfinish(m, int32(495164), int32(3470), int32(431219))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v8 + int32(48)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v53 = F_errstart(m, int32(12), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								if v53 == int32(0) {
									if base.Ui32(int32(15)) < base.Ui32(l1) {
										m.G0 = v8 + int32(48)
										return
									} else {
										if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
											m.G0 = v8 + int32(48)
											return
										} else {
											v78 = int32(0) - v10
											v79 = F_kill(m, v78, l1)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												if int32(0) <= v79 {
													m.G0 = v8 + int32(48)
													return
												} else {
													v85 = F_errstart(m, int32(12), int32(0))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														if v85 == int32(0) {
															m.G0 = v8 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
															F_errmsg_internal(m, int32(295542), v8)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																F_errfinish(m, int32(495164), int32(3470), int32(431219))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
									F_errmsg_internal(m, int32(295542), v8+int32(16))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errfinish(m, int32(495164), int32(3460), int32(431219))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											if base.Ui32(int32(15)) < base.Ui32(l1) {
												m.G0 = v8 + int32(48)
												return
											} else {
												if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
													m.G0 = v8 + int32(48)
													return
												} else {
													v78 = int32(0) - v10
													v79 = F_kill(m, v78, l1)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														if int32(0) <= v79 {
															m.G0 = v8 + int32(48)
															return
														} else {
															v85 = F_errstart(m, int32(12), int32(0))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return
															} else {
																if v85 == int32(0) {
																	m.G0 = v8 + int32(48)
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
																	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
																	F_errmsg_internal(m, int32(295542), v8)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(495164), int32(3470), int32(431219))
																		mBase = m.M
																		v98 = m.ExcPending
																		if v98 != 0 {
																			return
																		} else {
																			m.G0 = v8 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v47 = F_kill(m, v10, l1)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if int32(0) <= v47 {
					if base.Ui32(int32(15)) < base.Ui32(l1) {
						m.G0 = v8 + int32(48)
						return
					} else {
						if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
							m.G0 = v8 + int32(48)
							return
						} else {
							v78 = int32(0) - v10
							v79 = F_kill(m, v78, l1)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								if int32(0) <= v79 {
									m.G0 = v8 + int32(48)
									return
								} else {
									v85 = F_errstart(m, int32(12), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										if v85 == int32(0) {
											m.G0 = v8 + int32(48)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
											F_errmsg_internal(m, int32(295542), v8)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_errfinish(m, int32(495164), int32(3470), int32(431219))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													m.G0 = v8 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v53 = F_errstart(m, int32(12), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 == int32(0) {
							if base.Ui32(int32(15)) < base.Ui32(l1) {
								m.G0 = v8 + int32(48)
								return
							} else {
								if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
									m.G0 = v8 + int32(48)
									return
								} else {
									v78 = int32(0) - v10
									v79 = F_kill(m, v78, l1)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										if int32(0) <= v79 {
											m.G0 = v8 + int32(48)
											return
										} else {
											v85 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												if v85 == int32(0) {
													m.G0 = v8 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
													F_errmsg_internal(m, int32(295542), v8)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														F_errfinish(m, int32(495164), int32(3470), int32(431219))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v8 + int32(48)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
							F_errmsg_internal(m, int32(295542), v8+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(495164), int32(3460), int32(431219))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									if base.Ui32(int32(15)) < base.Ui32(l1) {
										m.G0 = v8 + int32(48)
										return
									} else {
										if int32(1)<<(uint(l1)%32)&int32(33356) == int32(0) {
											m.G0 = v8 + int32(48)
											return
										} else {
											v78 = int32(0) - v10
											v79 = F_kill(m, v78, l1)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												if int32(0) <= v79 {
													m.G0 = v8 + int32(48)
													return
												} else {
													v85 = F_errstart(m, int32(12), int32(0))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														if v85 == int32(0) {
															m.G0 = v8 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
															F_errmsg_internal(m, int32(295542), v8)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																F_errfinish(m, int32(495164), int32(3470), int32(431219))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_similar_to_escape_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_similar_escape_internal(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_similarity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(1)
	v28 = v20 + v27
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v33 = v31 & v27
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v28
	goto L6
L5:
	;
	v34 = v20 + int32(4)
	goto L6
L6:
	;
	if v31 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v63 = int32(0)
	F_generate_trgm_only(m, v17+int32(4), v34, v62, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v37 = int32(4)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v39&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v52 = int32(1)
	if v33 != 0 {
		v62 = int32(base.Ui32(v31)>>(uint(v52)%32)) - v52
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v48 = v37
	goto L13
L12:
	;
	v48 = base.B2i32(v39 == int32(18)) << (uint(v37) % 32)
	goto L13
L13:
	;
	if v39 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = v37
	goto L16
L15:
	;
	v51 = v48
	goto L16
L16:
	;
	v62 = v51
	goto L7
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v62 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)) = uint8(v71)
	if int32(2) <= v69 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v76 = v70 + int32(5)
	F_pg_qsort(m, v76, v69, int32(3), int32(6915))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v139 = v69
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v139*int32(12) + int32(20)
	v145 = int32(4)
	v147 = int32(1)
	v148 = v25 + v147
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v153 = v151 & v147
	if v153 != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v83 = int32(1)
	v84 = v63
	goto L23
L23:
	;
	v96 = int32(3)
	v98 = v76 + v83*v96
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1455]))
	v104 = m.T0[v103].(func(*base.Module, int32, int32) int32)(m, v98, v76+v84*v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v139 = v118 + int32(1)
	goto L21
L25:
	;
	v121 = v83 + int32(1)
	if v121 != v69 {
		v83 = v121
		v84 = v118
		goto L23
	} else {
		goto L29
	}
L26:
	;
	if v104 == int32(0) {
		v118 = v84
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v109 = v84 + int32(1)
	if v83 == v109 {
		v118 = v83
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v113 = v76 + v109*int32(3)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98))))
	*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v114)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+2)) = uint8(v116)
	v118 = v109
	goto L25
L29:
	;
	goto L24
L30:
	;
	v154 = v148
	goto L32
L31:
	;
	v154 = v25 + v145
	goto L32
L32:
	;
	if v151 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_generate_trgm_only(m, v17+v145, v154, v182, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v157 = int32(4)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v159&int32(254) == int32(2) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v172 = int32(1)
	if v153 != 0 {
		v182 = int32(base.Ui32(v151)>>(uint(v172)%32)) - v172
		goto L33
	} else {
		goto L43
	}
L37:
	;
	v168 = v157
	goto L39
L38:
	;
	v168 = base.B2i32(v159 == int32(18)) << (uint(v157) % 32)
	goto L39
L39:
	;
	if v159 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v171 = v157
	goto L42
L41:
	;
	v171 = v168
	goto L42
L42:
	;
	v182 = v171
	goto L33
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v182 = int32(base.Ui32(v176)>>(uint(int32(2))%32)) - int32(4)
	goto L33
L44:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)) = uint8(v188)
	if int32(2) <= v186 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v193 = v187 + int32(5)
	F_pg_qsort(m, v193, v186, int32(3), int32(6915))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v257 = v186
	goto L47
L47:
	;
	v261 = v257*int32(12) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v261
	v263 = int32(2)
	v265 = int32(5)
	v266 = int32(base.Ui32(v261)>>(uint(v263)%32)) - v265
	v267 = int32(3)
	v268 = base.I32_div_u_s(v266, v267)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v273 = int32(base.Ui32(v269)>>(uint(v263)%32)) - v265
	v275 = base.I32_div_u_s(v273, v267)
	if base.Ui32(v273) < base.Ui32(v267) {
		v346 = v2
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v201 = int32(1)
	v202 = int32(0)
	goto L49
L49:
	;
	v214 = int32(3)
	v216 = v193 + v201*v214
	v221 = *(*int32)(unsafe.Add(mBase, _consts[1455]))
	v222 = m.T0[v221].(func(*base.Module, int32, int32) int32)(m, v216, v193+v202*v214)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v257 = v236 + int32(1)
	goto L47
L51:
	;
	v239 = v201 + int32(1)
	if v239 != v186 {
		v201 = v239
		v202 = v236
		goto L49
	} else {
		goto L55
	}
L52:
	;
	if v222 == int32(0) {
		v236 = v202
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v227 = v202 + int32(1)
	if v201 == v227 {
		v236 = v201
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v231 = v193 + v227*int32(3)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	*(*uint16)(unsafe.Add(mBase, uint32(v231))) = uint16(v232)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+2)) = uint8(v234)
	v236 = v227
	goto L51
L55:
	;
	goto L50
L56:
	;
	F_pfree(m, v70)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L73
	}
L57:
	;
	if base.Ui32(v266) < base.Ui32(int32(3)) {
		v346 = v2
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v280 = int32(5)
	v281 = v70 + v280
	v283 = v187 + v280
	v285 = v283
	v286 = v281
	v297 = v2
	goto L59
L59:
	;
	v300 = base.I32_div_s(v285-v283, int32(3))
	if v300 < v268 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v346 = base.I32_reinterpret_f32(base.F32_div(base.F32_convert_i32_s(v328), base.F32_convert_i32_s(v268+v275-v328)))
	goto L56
L61:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[1455]))
	v304 = m.T0[v303].(func(*base.Module, int32, int32) int32)(m, v286, v285)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L65
	}
L62:
	;
	v328 = v297
	goto L63
L63:
	;
	goto L60
L64:
	;
	v323 = base.I32_div_s(v319-v281, int32(3))
	if v323 < v275 {
		v285 = v318
		v286 = v319
		v297 = v320
		goto L59
	} else {
		goto L72
	}
L65:
	;
	if int32(0) <= v304 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v304 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v314 = v285
	v315 = v297
	goto L68
L68:
	;
	v318 = v314
	v319 = v286 + int32(3)
	v320 = v315
	goto L64
L69:
	;
	v318 = v285 + int32(3)
	v319 = v286
	v320 = v297
	goto L64
L70:
	;
	goto L71
L71:
	;
	v314 = v285 + int32(3)
	v315 = v297 + int32(1)
	goto L68
L72:
	;
	v328 = v320
	goto L63
L73:
	;
	F_pfree(m, v187)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v353 != v20 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_pfree(m, v20)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v357 != v25 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v25)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	m.G0 = v17 + int32(16)
	return v346
L82:
	;
	goto L81
}
func F_similarity_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 float64
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(5417), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, _consts[1456]))
		return base.F64_le(v11, base.F64_promote_f32(base.F32_reinterpret_i32(v6)))
	}
}
func F_size_box(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v11 float64
	_ = v11
	var v17 float64
	_ = v17
	var v20 int64
	_ = v20
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v39 float64
	_ = v39
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v64 float64
	_ = v64
	var v80 float64
	_ = v80
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v2 = float64(0)
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) {
		v80 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L25
	} else {
		goto L27
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L25
	} else {
		goto L26
	}
L3:
	;
	return v80
L4:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v20 = base.I64_reinterpret_f64(v17) & int64(9223372036854775807)
	if base.B2i32(base.Ui64(v20) < base.Ui64(int64(9218868437227405313)))&base.F64_ge(v11, v17) != 0 {
		v80 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v25)&int64(9223372036854775807)) {
		v80 = v2
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = base.I64_reinterpret_f64(v31) & int64(9223372036854775807)
	if base.B2i32(base.Ui64(v34) < base.Ui64(int64(9218868437227405313)))&base.F64_ge(v25, v31) != 0 {
		v80 = v2
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v39 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v20) {
		v80 = v39
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v34) {
		v80 = v39
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v44 = base.F64_sub(v17, v11)
	v45 = base.F64_abs(v44)
	if base.F64_ne(v45, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v54 = base.F64_sub(v31, v25)
	v55 = base.F64_abs(v54)
	if base.F64_ne(v55, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if base.F64_eq(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v64 = base.F64_mul(v44, v54)
	if base.F64_ne(base.F64_abs(v64), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if base.F64_eq(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.F64_ne(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	if base.F64_ne(v64, float64(0)) != 0 {
		v80 = v64
		goto L3
	} else {
		goto L22
	}
L19:
	;
	if base.F64_eq(v45, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.F64_ne(v55, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	if base.F64_eq(v44, float64(0)) != 0 {
		v80 = v64
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if base.F64_ne(v54, float64(0)) != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v80 = v64
	goto L3
L25:
	;
	return float64(0)
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_skeys(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_skeys(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_smgrcreate(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = int32(4510044)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v6 + int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(80))+uint32(_consts[801])))
	m.T0[v15].(func(*base.Module, int32, int32, int32))(m, l0, l1, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(4510044)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[414]))
		*(*int32)(unsafe.Add(mBase, _consts[414])) = v20 - int32(1)
		return
	}
}
func F_smgropen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = int32(4510044)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v11 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	if v16 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(360777252880)
		v26 = F_hash_create(m, int32(392985), int32(400), v5+int32(-56), int32(40))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = int32(4438872)
			*(*int32)(unsafe.Add(mBase, _consts[798])) = v31
			*(*int32)(unsafe.Add(mBase, _consts[799])) = v31
			*(*int32)(unsafe.Add(mBase, _consts[797])) = v26
			v38 = v26
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v39
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v41
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
			v49 = F_hash_search(m, v38, v5+int32(-56), int32(1), v5+int32(-1))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+63)))
				if v51 == int32(0) {
					v54 = int64(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v54
					*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = int64(4294967295)
					v63 = v49 + int32(76)
					v65 = *(*int32)(unsafe.Add(mBase, _consts[798]))
					if v65 != 0 {
						v67 = *(*int32)(unsafe.Add(mBase, _consts[799]))
						v72 = v67
					} else {
						v69 = int32(4438872)
						*(*int32)(unsafe.Add(mBase, _consts[798])) = v69
						v72 = v69
					}
					*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v72
					v74 = int32(4438872)
					*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v63
					*(*int32)(unsafe.Add(mBase, _consts[799])) = v63
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v49)+36))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v79*int32(80))+uint32(_consts[800])))
					m.T0[v84].(func(*base.Module, int32))(m, v49)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v89 = int32(4510044)
						v91 = *(*int32)(unsafe.Add(mBase, _consts[414]))
						*(*int32)(unsafe.Add(mBase, _consts[414])) = v91 - int32(1)
						m.G0 = v7 - int32(-64)
						return v49
					}
				} else {
					v89 = int32(4510044)
					v91 = *(*int32)(unsafe.Add(mBase, _consts[414]))
					*(*int32)(unsafe.Add(mBase, _consts[414])) = v91 - int32(1)
					m.G0 = v7 - int32(-64)
					return v49
				}
			}
		}
	} else {
		v38 = v16
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v39
		v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
		v49 = F_hash_search(m, v38, v5+int32(-56), int32(1), v5+int32(-1))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+63)))
			if v51 == int32(0) {
				v54 = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v54
				*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = int64(4294967295)
				v63 = v49 + int32(76)
				v65 = *(*int32)(unsafe.Add(mBase, _consts[798]))
				if v65 != 0 {
					v67 = *(*int32)(unsafe.Add(mBase, _consts[799]))
					v72 = v67
				} else {
					v69 = int32(4438872)
					*(*int32)(unsafe.Add(mBase, _consts[798])) = v69
					v72 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v72
				v74 = int32(4438872)
				*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v74
				*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v63
				*(*int32)(unsafe.Add(mBase, _consts[799])) = v63
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v49)+36))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v79*int32(80))+uint32(_consts[800])))
				m.T0[v84].(func(*base.Module, int32))(m, v49)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v89 = int32(4510044)
					v91 = *(*int32)(unsafe.Add(mBase, _consts[414]))
					*(*int32)(unsafe.Add(mBase, _consts[414])) = v91 - int32(1)
					m.G0 = v7 - int32(-64)
					return v49
				}
			} else {
				v89 = int32(4510044)
				v91 = *(*int32)(unsafe.Add(mBase, _consts[414]))
				*(*int32)(unsafe.Add(mBase, _consts[414])) = v91 - int32(1)
				m.G0 = v7 - int32(-64)
				return v49
			}
		}
	}
}
func F_smgrregistersync(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = int32(4510044)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v5 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(80))+uint32(_consts[803])))
	m.T0[v14].(func(*base.Module, int32, int32))(m, l0, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = int32(4510044)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[414]))
		*(*int32)(unsafe.Add(mBase, _consts[414])) = v19 - int32(1)
		return
	}
}
func F_smgrzeroextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v5 = int32(4510044)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v7 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12*int32(80))+uint32(_consts[802])))
	m.T0[v17].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v24 = l0 + l1<<(uint(int32(2))%32) + int32(20)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		if v27 != l2 {
			v29 = int32(-1)
		} else {
			v29 = l2 + l3
		}
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v29
		v31 = int32(4510044)
		v33 = *(*int32)(unsafe.Add(mBase, _consts[414]))
		*(*int32)(unsafe.Add(mBase, _consts[414])) = v33 - int32(1)
		return
	}
}
func F_sortins_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v8 < v11 {
		return int32(-1)
	} else {
		v15 = int32(1)
		if v11 < v8 {
			v30 = v15
			return v30
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v17 < v18 {
				return int32(-1)
			} else {
				if v18 < v17 {
					v30 = v15
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v24 < v25 {
						v30 = int32(-1)
					} else {
						v30 = base.B2i32(v25 < v24)
					}
				}
				return v30
			}
		}
	}
}
func F_spanish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 < v12 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v512 < v12 {
		goto L149
	} else {
		goto L150
	}
L2:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+8)) = v498
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v292 < v12 {
		goto L83
	} else {
		goto L84
	}
L4:
	;
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v24 = v12
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	goto L9
L8:
	;
	v65 = v61
	goto L4
L9:
	;
	if v12 == v24 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = int32(0)
	goto L8
L11:
	;
	v65 = int32(-1)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v12))))
	if int32(252) < v39 {
		v61 = v36
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v47)>>(uint(v41&int32(7))%32))&int32(1) == int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L17
L17:
	;
	goto L10
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 < v66 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v180 < v66 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	if v115 != 0 {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v77 = v66
	goto L23
L22:
	;
	v77 = v75
	goto L23
L23:
	;
	goto L25
L24:
	;
	v115 = v112
	goto L20
L25:
	;
	if v66 == v77 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v112 = int32(0)
	goto L24
L27:
	;
	v115 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v66))))
	if int32(252) < v90 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(1)
	goto L34
L31:
	;
	v92 = v90 - int32(97)
	if v92 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v95 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v92)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v99)>>(uint(v92&int32(7))%32))&v95 != 0 {
		v112 = v95
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v124 < v123 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v164 < int32(0) {
		goto L19
	} else {
		goto L51
	}
L37:
	;
	v126 = v123
	goto L39
L38:
	;
	v126 = v124
	goto L39
L39:
	;
	v133 = v123
	goto L41
L40:
	;
	v164 = v144
	goto L36
L41:
	;
	if v133 == v126 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v133))))
	if int32(252) < v139 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v156 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v156
	v133 = v156
	goto L41
L47:
	;
	v141 = v139 - int32(97)
	if v141 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v144 = int32(1)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v141)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v148)>>(uint(v141&int32(7))%32))&v144 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v167 + v164
	goto L2
L52:
	;
	if v223 != 0 {
		goto L3
	} else {
		goto L66
	}
L53:
	;
	v182 = v66
	goto L55
L54:
	;
	v182 = v180
	goto L55
L55:
	;
	goto L57
L56:
	;
	v223 = v219
	goto L52
L57:
	;
	if v66 == v182 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v219 = int32(0)
	goto L56
L59:
	;
	v223 = int32(-1)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v66))))
	if int32(252) < v197 {
		v219 = v194
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v199 = v197 - int32(97)
	if v199 < int32(0) {
		v219 = v194
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v199)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v205)>>(uint(v199&int32(7))%32))&int32(1) == int32(0) {
		v219 = v194
		goto L56
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(1)
	goto L65
L65:
	;
	goto L58
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v233 < v232 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v276 < int32(0) {
		goto L3
	} else {
		goto L81
	}
L68:
	;
	v235 = v232
	goto L70
L69:
	;
	v235 = v233
	goto L70
L70:
	;
	v242 = v232
	goto L72
L71:
	;
	v276 = int32(1)
	goto L67
L72:
	;
	if v242 == v235 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v276 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+v242))))
	if int32(252) < v250 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v252 = v250 - int32(97)
	if v252 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v252)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v258)>>(uint(v252&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v267 = v242 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v267
	v242 = v267
	goto L72
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v279 + v276
	goto L2
L82:
	;
	if v332 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v294 = v12
	goto L85
L84:
	;
	v294 = v292
	goto L85
L85:
	;
	goto L87
L86:
	;
	v332 = v329
	goto L82
L87:
	;
	if v12 == v294 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v329 = int32(0)
	goto L86
L89:
	;
	v332 = int32(-1)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v12))))
	if int32(252) < v307 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L96
L93:
	;
	v309 = v307 - int32(97)
	if v309 < int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v312 = int32(1)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v309)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v316)>>(uint(v309&int32(7))%32))&v312 != 0 {
		v329 = v312
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L88
L97:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v342 < v333 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v447 < v333 {
		goto L132
	} else {
		goto L133
	}
L99:
	;
	if v382 != 0 {
		goto L98
	} else {
		goto L114
	}
L100:
	;
	v344 = v333
	goto L102
L101:
	;
	v344 = v342
	goto L102
L102:
	;
	goto L104
L103:
	;
	v382 = v379
	goto L99
L104:
	;
	if v333 == v344 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v379 = int32(0)
	goto L103
L106:
	;
	v382 = int32(-1)
	goto L99
L107:
	;
	goto L108
L108:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v333))))
	if int32(252) < v357 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333 + int32(1)
	goto L113
L110:
	;
	v359 = v357 - int32(97)
	if v359 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v362 = int32(1)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v359)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v366)>>(uint(v359&int32(7))%32))&v362 != 0 {
		v379 = v362
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v391 < v390 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v431 < int32(0) {
		goto L98
	} else {
		goto L130
	}
L116:
	;
	v393 = v390
	goto L118
L117:
	;
	v393 = v391
	goto L118
L118:
	;
	v400 = v390
	goto L120
L119:
	;
	v431 = v411
	goto L115
L120:
	;
	if v400 == v393 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v431 = int32(-1)
	goto L115
L123:
	;
	goto L124
L124:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v400))))
	if int32(252) < v406 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v423 = v400 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v423
	v400 = v423
	goto L120
L126:
	;
	v408 = v406 - int32(97)
	if v408 < int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v411 = int32(1)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v408)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v415)>>(uint(v408&int32(7))%32))&v411 != 0 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	goto L125
L130:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v434 + v431
	goto L2
L131:
	;
	if v490 != 0 {
		goto L1
	} else {
		goto L145
	}
L132:
	;
	v449 = v333
	goto L134
L133:
	;
	v449 = v447
	goto L134
L134:
	;
	goto L136
L135:
	;
	v490 = v486
	goto L131
L136:
	;
	if v333 == v449 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v486 = int32(0)
	goto L135
L138:
	;
	v490 = int32(-1)
	goto L131
L139:
	;
	goto L140
L140:
	;
	v461 = int32(1)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+v333))))
	if int32(252) < v464 {
		v486 = v461
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v466 = v464 - int32(97)
	if v466 < int32(0) {
		v486 = v461
		goto L135
	} else {
		goto L142
	}
L142:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v466)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v472)>>(uint(v466&int32(7))%32))&int32(1) == int32(0) {
		v486 = v461
		goto L135
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333 + int32(1)
	goto L144
L144:
	;
	goto L137
L145:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v492 <= v491 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v498 = v491 + int32(1)
	goto L2
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v733
	v737 = v733 - int32(1)
	if v737 <= v12 {
		goto L211
	} else {
		goto L212
	}
L148:
	;
	if v552 < int32(0) {
		goto L147
	} else {
		goto L163
	}
L149:
	;
	v514 = v12
	goto L151
L150:
	;
	v514 = v512
	goto L151
L151:
	;
	v521 = v12
	goto L153
L152:
	;
	v552 = v532
	goto L148
L153:
	;
	if v521 == v514 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v552 = int32(-1)
	goto L148
L156:
	;
	goto L157
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525+v521))))
	if int32(252) < v527 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v544 = v521 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v544
	v521 = v544
	goto L153
L159:
	;
	v529 = v527 - int32(97)
	if v529 < int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v532 = int32(1)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v529)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v536)>>(uint(v529&int32(7))%32))&v532 != 0 {
		goto L152
	} else {
		goto L161
	}
L161:
	;
	goto L158
L163:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v556 = v555 + v552
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v556
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v567 < v556 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v610 < int32(0) {
		goto L147
	} else {
		goto L178
	}
L165:
	;
	v569 = v556
	goto L167
L166:
	;
	v569 = v567
	goto L167
L167:
	;
	v576 = v556
	goto L169
L168:
	;
	v610 = int32(1)
	goto L164
L169:
	;
	if v576 == v569 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v610 = int32(-1)
	goto L164
L172:
	;
	goto L173
L173:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+v576))))
	if int32(252) < v584 {
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v586 = v584 - int32(97)
	if v586 < int32(0) {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v586)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v592)>>(uint(v586&int32(7))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v601 = v576 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601
	v576 = v601
	goto L169
L178:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v614 = v613 + v610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v614
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v616)+4)) = v614
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v626 < v625 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v666 < int32(0) {
		goto L147
	} else {
		goto L194
	}
L180:
	;
	v628 = v625
	goto L182
L181:
	;
	v628 = v626
	goto L182
L182:
	;
	v635 = v625
	goto L184
L183:
	;
	v666 = v646
	goto L179
L184:
	;
	if v635 == v628 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v666 = int32(-1)
	goto L179
L187:
	;
	goto L188
L188:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639+v635))))
	if int32(252) < v641 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v658 = v635 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v635 = v658
	goto L184
L190:
	;
	v643 = v641 - int32(97)
	if v643 < int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v646 = int32(1)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v643)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v650)>>(uint(v643&int32(7))%32))&v646 != 0 {
		goto L183
	} else {
		goto L192
	}
L192:
	;
	goto L189
L194:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v670 = v669 + v666
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v670
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v681 < v670 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v724 < int32(0) {
		goto L147
	} else {
		goto L209
	}
L196:
	;
	v683 = v670
	goto L198
L197:
	;
	v683 = v681
	goto L198
L198:
	;
	v690 = v670
	goto L200
L199:
	;
	v724 = int32(1)
	goto L195
L200:
	;
	if v690 == v683 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v724 = int32(-1)
	goto L195
L203:
	;
	goto L204
L204:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v690))))
	if int32(252) < v698 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v700 = v698 - int32(97)
	if v700 < int32(0) {
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v700)>>(uint(int32(3))%32)))+uint32(_consts[1443]))))
	if int32(base.Ui32(v706)>>(uint(v700&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v715 = v690 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v715
	v690 = v715
	goto L200
L209:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v727))) = v728 + v724
	goto L147
L210:
	;
	return v1370
L211:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v840
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v840-int32(2) <= v843 {
		goto L247
	} else {
		goto L248
	}
L212:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739+v737))))
	if v741&int32(224) != int32(96) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	if int32(1)<<(uint(v741)%32)&int32(557090) == int32(0) {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	v754 = F_find_among_b(m, l0, int32(4218704), int32(13))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	return int32(0)
L216:
	;
	if v754 == int32(0) {
		goto L211
	} else {
		goto L217
	}
L217:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v760
	v763 = v760 - int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v763 <= v764 {
		goto L211
	} else {
		goto L218
	}
L218:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766+v763))))
	switch v768 - int32(111) {
	case 0, 3:
		goto L219
	default:
		goto L211
	}
L219:
	;
	v773 = F_find_among_b(m, l0, int32(4218976), int32(11))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L215
	} else {
		goto L220
	}
L220:
	;
	if v773 == int32(0) {
		goto L211
	} else {
		goto L221
	}
L221:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+8))
	if v777 < v779 {
		goto L211
	} else {
		goto L222
	}
L222:
	;
	switch v773 - int32(1) {
	case 0:
		goto L229
	case 1:
		goto L228
	case 2:
		goto L227
	case 3:
		goto L226
	case 4:
		goto L225
	case 5:
		goto L224
	case 6:
		goto L223
	default:
		goto L211
	}
L223:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v777 <= v822 {
		goto L211
	} else {
		goto L242
	}
L224:
	;
	v818 = F_slice_del(m, l0)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L215
	} else {
		goto L240
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v814 = F_slice_from_s(m, l0, int32(2), int32(2181681))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L215
	} else {
		goto L238
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v807 = F_slice_from_s(m, l0, int32(2), int32(2181679))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L215
	} else {
		goto L236
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v800 = F_slice_from_s(m, l0, int32(2), int32(2181677))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L215
	} else {
		goto L234
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v793 = F_slice_from_s(m, l0, int32(4), int32(2181673))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L215
	} else {
		goto L232
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v786 = F_slice_from_s(m, l0, int32(5), int32(2181668))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L215
	} else {
		goto L230
	}
L230:
	;
	if int32(0) <= v786 {
		goto L211
	} else {
		goto L231
	}
L231:
	;
	v1370 = v786
	goto L210
L232:
	;
	if int32(0) <= v793 {
		goto L211
	} else {
		goto L233
	}
L233:
	;
	v1370 = v793
	goto L210
L234:
	;
	if int32(0) <= v800 {
		goto L211
	} else {
		goto L235
	}
L235:
	;
	v1370 = v800
	goto L210
L236:
	;
	if int32(0) <= v807 {
		goto L211
	} else {
		goto L237
	}
L237:
	;
	v1370 = v807
	goto L210
L238:
	;
	if int32(0) <= v814 {
		goto L211
	} else {
		goto L239
	}
L239:
	;
	v1370 = v814
	goto L210
L240:
	;
	if int32(0) <= v818 {
		goto L211
	} else {
		goto L241
	}
L241:
	;
	v1370 = v818
	goto L210
L242:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824+v777-int32(1)))))
	if v828 != int32(117) {
		goto L211
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v777 - int32(1)
	v834 = F_slice_del(m, l0)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L215
	} else {
		goto L244
	}
L244:
	;
	if v834 < int32(0) {
		v1370 = v834
		goto L210
	} else {
		goto L245
	}
L245:
	;
	goto L211
L246:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1234
	v1239 = F_find_among_b(m, l0, int32(4222496), int32(8))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L215
	} else {
		goto L376
	}
L247:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1128
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+8))
	if v1128 < v1131 {
		goto L337
	} else {
		goto L338
	}
L248:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847+v840-int32(1)))))
	if v851&int32(224) != int32(96) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	if int32(1)<<(uint(v851)%32)&int32(835634) == int32(0) {
		goto L247
	} else {
		goto L250
	}
L250:
	;
	v864 = F_find_among_b(m, l0, int32(4219200), int32(46))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L215
	} else {
		goto L251
	}
L251:
	;
	if v864 == int32(0) {
		goto L247
	} else {
		goto L252
	}
L252:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v868
	switch v864 - int32(1) {
	case 0:
		goto L261
	case 1:
		goto L260
	case 2:
		goto L259
	case 3:
		goto L258
	case 4:
		goto L257
	case 5:
		goto L256
	case 6:
		goto L255
	case 7:
		goto L254
	case 8:
		goto L253
	default:
		goto L246
	}
L253:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	if v868 < v1091 {
		goto L247
	} else {
		goto L326
	}
L254:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	if v868 < v1050 {
		goto L247
	} else {
		goto L315
	}
L255:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	if v868 < v1015 {
		goto L247
	} else {
		goto L305
	}
L256:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)+4))
	if v868 < v943 {
		goto L247
	} else {
		goto L285
	}
L257:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	if v868 < v934 {
		goto L247
	} else {
		goto L282
	}
L258:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)))
	if v868 < v925 {
		goto L247
	} else {
		goto L279
	}
L259:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	if v868 < v916 {
		goto L247
	} else {
		goto L276
	}
L260:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	if v868 < v880 {
		goto L247
	} else {
		goto L265
	}
L261:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	if v868 < v873 {
		goto L247
	} else {
		goto L262
	}
L262:
	;
	v875 = F_slice_del(m, l0)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L215
	} else {
		goto L263
	}
L263:
	;
	if int32(0) <= v875 {
		goto L246
	} else {
		goto L264
	}
L264:
	;
	v1370 = v875
	goto L210
L265:
	;
	v882 = F_slice_del(m, l0)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L215
	} else {
		goto L266
	}
L266:
	;
	if v882 < int32(0) {
		v1370 = v882
		goto L210
	} else {
		goto L267
	}
L267:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v886
	v888 = int32(2)
	v890 = int32(0)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v886-v893 < v888 {
		v903 = v890
		goto L269
	} else {
		goto L270
	}
L268:
	;
	if v903 == int32(0) {
		goto L246
	} else {
		goto L272
	}
L269:
	;
	goto L268
L270:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v899 = F_memcmp(m, v896+v886-v888, int32(2181758), v888)
	mBase = m.M
	if v899 != 0 {
		v903 = v890
		goto L269
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v886 - v888
	v903 = int32(1)
	goto L269
L272:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v906
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	if v906 < v909 {
		goto L246
	} else {
		goto L273
	}
L273:
	;
	v911 = F_slice_del(m, l0)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L215
	} else {
		goto L274
	}
L274:
	;
	if int32(0) <= v911 {
		goto L246
	} else {
		goto L275
	}
L275:
	;
	v1370 = v911
	goto L210
L276:
	;
	v920 = F_slice_from_s(m, l0, int32(3), int32(2181760))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L215
	} else {
		goto L277
	}
L277:
	;
	if int32(0) <= v920 {
		goto L246
	} else {
		goto L278
	}
L278:
	;
	v1370 = v920
	goto L210
L279:
	;
	v929 = F_slice_from_s(m, l0, int32(1), int32(2181763))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L215
	} else {
		goto L280
	}
L280:
	;
	if int32(0) <= v929 {
		goto L246
	} else {
		goto L281
	}
L281:
	;
	v1370 = v929
	goto L210
L282:
	;
	v938 = F_slice_from_s(m, l0, int32(4), int32(2181764))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L215
	} else {
		goto L283
	}
L283:
	;
	if int32(0) <= v938 {
		goto L246
	} else {
		goto L284
	}
L284:
	;
	v1370 = v938
	goto L210
L285:
	;
	v945 = F_slice_del(m, l0)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L215
	} else {
		goto L286
	}
L286:
	;
	if v945 < int32(0) {
		v1370 = v945
		goto L210
	} else {
		goto L287
	}
L287:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v949
	v952 = v949 - int32(1)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v952 <= v953 {
		goto L246
	} else {
		goto L288
	}
L288:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955+v952))))
	if v957&int32(224) != int32(96) {
		goto L246
	} else {
		goto L289
	}
L289:
	;
	if int32(1)<<(uint(v957)%32)&int32(4718616) == int32(0) {
		goto L246
	} else {
		goto L290
	}
L290:
	;
	v970 = F_find_among_b(m, l0, int32(4220128), int32(4))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L215
	} else {
		goto L291
	}
L291:
	;
	if v970 == int32(0) {
		goto L246
	} else {
		goto L292
	}
L292:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v974
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	if v974 < v977 {
		goto L246
	} else {
		goto L293
	}
L293:
	;
	v979 = F_slice_del(m, l0)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L215
	} else {
		goto L294
	}
L294:
	;
	if v979 < int32(0) {
		v1370 = v979
		goto L210
	} else {
		goto L295
	}
L295:
	;
	if v970 != int32(1) {
		goto L246
	} else {
		goto L296
	}
L296:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v985
	v987 = int32(2)
	v989 = int32(0)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v985-v992 < v987 {
		v1002 = v989
		goto L298
	} else {
		goto L299
	}
L297:
	;
	if v1002 == int32(0) {
		goto L246
	} else {
		goto L301
	}
L298:
	;
	goto L297
L299:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v998 = F_memcmp(m, v995+v985-v987, int32(2181768), v987)
	mBase = m.M
	if v998 != 0 {
		v1002 = v989
		goto L298
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v985 - v987
	v1002 = int32(1)
	goto L298
L301:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1005
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)))
	if v1005 < v1008 {
		goto L246
	} else {
		goto L302
	}
L302:
	;
	v1010 = F_slice_del(m, l0)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L215
	} else {
		goto L303
	}
L303:
	;
	if int32(0) <= v1010 {
		goto L246
	} else {
		goto L304
	}
L304:
	;
	v1370 = v1010
	goto L210
L305:
	;
	v1017 = F_slice_del(m, l0)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L215
	} else {
		goto L306
	}
L306:
	;
	if v1017 < int32(0) {
		v1370 = v1017
		goto L210
	} else {
		goto L307
	}
L307:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1021
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1021-int32(3) <= v1023 {
		goto L246
	} else {
		goto L308
	}
L308:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027+v1021-int32(1)))))
	if v1031 != int32(101) {
		goto L246
	} else {
		goto L309
	}
L309:
	;
	v1036 = F_find_among_b(m, l0, int32(4220208), int32(3))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L215
	} else {
		goto L310
	}
L310:
	;
	if v1036 == int32(0) {
		goto L246
	} else {
		goto L311
	}
L311:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1040
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1042)))
	if v1040 < v1043 {
		goto L246
	} else {
		goto L312
	}
L312:
	;
	v1045 = F_slice_del(m, l0)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L215
	} else {
		goto L313
	}
L313:
	;
	if int32(0) <= v1045 {
		goto L246
	} else {
		goto L314
	}
L314:
	;
	v1370 = v1045
	goto L210
L315:
	;
	v1052 = F_slice_del(m, l0)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L215
	} else {
		goto L316
	}
L316:
	;
	if v1052 < int32(0) {
		v1370 = v1052
		goto L210
	} else {
		goto L317
	}
L317:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1056
	v1059 = v1056 - int32(1)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1059 <= v1060 {
		goto L246
	} else {
		goto L318
	}
L318:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062+v1059))))
	if v1064&int32(224) != int32(96) {
		goto L246
	} else {
		goto L319
	}
L319:
	;
	if int32(1)<<(uint(v1064)%32)&int32(4198408) == int32(0) {
		goto L246
	} else {
		goto L320
	}
L320:
	;
	v1077 = F_find_among_b(m, l0, int32(4220272), int32(3))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L215
	} else {
		goto L321
	}
L321:
	;
	if v1077 == int32(0) {
		goto L246
	} else {
		goto L322
	}
L322:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1081
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	if v1081 < v1084 {
		goto L246
	} else {
		goto L323
	}
L323:
	;
	v1086 = F_slice_del(m, l0)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L215
	} else {
		goto L324
	}
L324:
	;
	if int32(0) <= v1086 {
		goto L246
	} else {
		goto L325
	}
L325:
	;
	v1370 = v1086
	goto L210
L326:
	;
	v1093 = F_slice_del(m, l0)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L215
	} else {
		goto L327
	}
L327:
	;
	if v1093 < int32(0) {
		v1370 = v1093
		goto L210
	} else {
		goto L328
	}
L328:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1097
	v1099 = int32(2)
	v1101 = int32(0)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1097-v1104 < v1099 {
		v1114 = v1101
		goto L330
	} else {
		goto L331
	}
L329:
	;
	if v1114 == int32(0) {
		goto L246
	} else {
		goto L333
	}
L330:
	;
	goto L329
L331:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1110 = F_memcmp(m, v1107+v1097-v1099, int32(2181770), v1099)
	mBase = m.M
	if v1110 != 0 {
		v1114 = v1101
		goto L330
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1097 - v1099
	v1114 = int32(1)
	goto L330
L333:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1117
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)))
	if v1117 < v1120 {
		goto L246
	} else {
		goto L334
	}
L334:
	;
	v1122 = F_slice_del(m, l0)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L215
	} else {
		goto L335
	}
L335:
	;
	if int32(0) <= v1122 {
		goto L246
	} else {
		goto L336
	}
L336:
	;
	v1370 = v1122
	goto L210
L337:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1179
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1181)+8))
	if v1179 < v1182 {
		goto L246
	} else {
		goto L357
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1128
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1131
	v1138 = F_find_among_b(m, l0, int32(4220336), int32(12))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L215
	} else {
		goto L339
	}
L339:
	;
	if v1138 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1134
	goto L337
L341:
	;
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1134
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1144
	if v1144 <= v1134 {
		goto L337
	} else {
		goto L343
	}
L343:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147+v1144-int32(1)))))
	if v1151 != int32(117) {
		goto L337
	} else {
		goto L344
	}
L344:
	;
	v1154 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1144 - v1154
	v1158 = F_slice_del(m, l0)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L215
	} else {
		goto L346
	}
L345:
	;
	v1170 = int32(0)
	v1171 = base.B2i32(v1165 < v1170)
	if v1171 == v1170 {
		goto L246
	} else {
		goto L353
	}
L346:
	;
	if int32(0) <= v1158 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1165 = v1154
	goto L349
L348:
	;
	v1165 = v1158 >> (uint(int32(31)) % 32) & v1158
	goto L349
L349:
	;
	if v1165 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1169 = int32(base.Ui32(v1165) >> (uint(int32(31)) % 32))
	goto L352
L351:
	;
	v1169 = int32(4)
	goto L352
L352:
	;
	switch v1169 {
	case 0:
		goto L246
	default:
		goto L345
	case 4:
		goto L337
	}
L353:
	;
	if v1165 < v1170 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1175 = v1165
	goto L356
L355:
	;
	v1175 = int32(1)
	goto L356
L356:
	;
	return v1175
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1179
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1182
	v1189 = F_find_among_b(m, l0, int32(4220576), int32(96))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L215
	} else {
		goto L358
	}
L358:
	;
	if v1189 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1185
	goto L246
L360:
	;
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1185
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1195
	switch v1189 - int32(1) {
	case 0:
		goto L363
	case 1:
		goto L362
	default:
		goto L246
	}
L362:
	;
	v1226 = F_slice_del(m, l0)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L215
	} else {
		goto L373
	}
L363:
	;
	if v1195 <= v1185 {
		v1218 = v1195
		goto L364
	} else {
		goto L365
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1218
	v1222 = F_slice_del(m, l0)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L215
	} else {
		goto L371
	}
L365:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1201 = v1200 + v1195
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201-int32(1)))))
	if v1204 != int32(117) {
		v1218 = v1195
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v1208 = v1195 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1208
	if v1208 <= v1185 {
		v1218 = v1195
		goto L364
	} else {
		goto L367
	}
L367:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201-int32(2)))))
	if v1213 == int32(103) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1216 = v1208
	goto L370
L369:
	;
	v1216 = v1195
	goto L370
L370:
	;
	v1218 = v1216
	goto L364
L371:
	;
	if int32(0) <= v1222 {
		goto L246
	} else {
		goto L372
	}
L372:
	;
	v1370 = v1222
	goto L210
L373:
	;
	if v1226 < int32(0) {
		v1370 = v1226
		goto L210
	} else {
		goto L374
	}
L374:
	;
	goto L246
L375:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1293
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1297 = v1293
	v1298 = v1295
	goto L393
L376:
	;
	if v1239 == int32(0) {
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1243
	switch v1239 - int32(1) {
	case 0:
		goto L379
	case 1:
		goto L378
	default:
		goto L375
	}
L378:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+8))
	if v1243 < v1255 {
		goto L375
	} else {
		goto L383
	}
L379:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+8))
	if v1243 < v1248 {
		goto L375
	} else {
		goto L380
	}
L380:
	;
	v1250 = F_slice_del(m, l0)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L215
	} else {
		goto L381
	}
L381:
	;
	if int32(0) <= v1250 {
		goto L375
	} else {
		goto L382
	}
L382:
	;
	v1370 = v1250
	goto L210
L383:
	;
	v1257 = F_slice_del(m, l0)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L215
	} else {
		goto L384
	}
L384:
	;
	if v1257 < int32(0) {
		v1370 = v1257
		goto L210
	} else {
		goto L385
	}
L385:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1261
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1261 <= v1263 {
		goto L375
	} else {
		goto L386
	}
L386:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1266 = v1265 + v1261
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266-int32(1)))))
	if v1269 != int32(117) {
		goto L375
	} else {
		goto L387
	}
L387:
	;
	v1273 = v1261 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1273
	if v1273 <= v1263 {
		goto L375
	} else {
		goto L388
	}
L388:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266-int32(2)))))
	if v1279 != int32(103) {
		goto L375
	} else {
		goto L389
	}
L389:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+8))
	if v1261 <= v1283 {
		goto L375
	} else {
		goto L390
	}
L390:
	;
	v1285 = F_slice_del(m, l0)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L215
	} else {
		goto L391
	}
L391:
	;
	if v1285 < int32(0) {
		v1370 = v1285
		goto L210
	} else {
		goto L392
	}
L392:
	;
	goto L375
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1297
	if v1298 <= v1297 {
		goto L399
	} else {
		goto L400
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1293
	v1370 = int32(1)
	goto L210
L395:
	;
	goto L394
L396:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1297 = v1366
	v1298 = v1365
	goto L393
L397:
	;
	if v1356 <= v1355 {
		goto L395
	} else {
		goto L420
	}
L398:
	;
	v1318 = F_find_among(m, l0, int32(4222656), int32(6))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L215
	} else {
		goto L403
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1297
	v1355 = v1297
	v1356 = v1298
	goto L397
L400:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303+v1297))))
	v1306 = int32(224)
	if v1305&v1306 != v1306 {
		goto L399
	} else {
		goto L401
	}
L401:
	;
	if int32(1)<<(uint(v1305)%32)&int32(67641858) != 0 {
		goto L398
	} else {
		goto L402
	}
L402:
	;
	goto L399
L403:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1320
	switch v1318 - int32(1) {
	case 0:
		goto L409
	case 1:
		goto L408
	case 2:
		goto L407
	case 3:
		goto L406
	case 4:
		goto L405
	case 5:
		goto L404
	default:
		goto L396
	}
L404:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1355 = v1320
	v1356 = v1354
	goto L397
L405:
	;
	v1350 = F_slice_from_s(m, l0, int32(1), int32(2182478))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L215
	} else {
		goto L418
	}
L406:
	;
	v1344 = F_slice_from_s(m, l0, int32(1), int32(2182477))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L215
	} else {
		goto L416
	}
L407:
	;
	v1338 = F_slice_from_s(m, l0, int32(1), int32(2182476))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L215
	} else {
		goto L414
	}
L408:
	;
	v1332 = F_slice_from_s(m, l0, int32(1), int32(2182475))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L215
	} else {
		goto L412
	}
L409:
	;
	v1326 = F_slice_from_s(m, l0, int32(1), int32(2182474))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L215
	} else {
		goto L410
	}
L410:
	;
	if int32(0) <= v1326 {
		goto L396
	} else {
		goto L411
	}
L411:
	;
	v1370 = v1326
	goto L210
L412:
	;
	if int32(0) <= v1332 {
		goto L396
	} else {
		goto L413
	}
L413:
	;
	v1370 = v1332
	goto L210
L414:
	;
	if int32(0) <= v1338 {
		goto L396
	} else {
		goto L415
	}
L415:
	;
	v1370 = v1338
	goto L210
L416:
	;
	if int32(0) <= v1344 {
		goto L396
	} else {
		goto L417
	}
L417:
	;
	v1370 = v1344
	goto L210
L418:
	;
	if int32(0) <= v1350 {
		goto L396
	} else {
		goto L419
	}
L419:
	;
	v1370 = v1350
	goto L210
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1355 + int32(1)
	goto L396
}
func F_spgbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v4 = F_smgr_bulk_start_rel(m, l0, int32(3))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = F_smgr_bulk_get_buf(m, v4)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_PageInit(m, v6, int32(8192), int32(8))
			mBase = m.M
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
			v13 = v6 + v12
			v14 = int32(65410)
			*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)) = uint16(v14)
			v16 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v13))) = uint16(v16)
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+80)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6-int32(-64)))) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v6)+88)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = int64(-1173640210)
			v38 = int32(92)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)) = uint16(v38)
			v40 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v40
			F_smgr_bulk_write(m, v4, int32(0), v6, int32(1))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				v58 = F_smgr_bulk_get_buf(m, v4)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = int32(4)
					F_PageInit(m, v58, int32(8192), int32(8))
					mBase = m.M
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
					v65 = v58 + v64
					v66 = int32(65410)
					*(*uint16)(unsafe.Add(mBase, uint32(v65)+6)) = uint16(v66)
					*(*uint16)(unsafe.Add(mBase, uint32(v65))) = uint16(v60)
					v69 = int32(1)
					F_smgr_bulk_write(m, v4, v69, v58, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v73 = F_smgr_bulk_get_buf(m, v4)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v75 = int32(12)
							F_PageInit(m, v73, int32(8192), int32(8))
							mBase = m.M
							v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
							v80 = v73 + v79
							v81 = int32(65410)
							*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)) = uint16(v81)
							*(*uint16)(unsafe.Add(mBase, uint32(v80))) = uint16(v75)
							F_smgr_bulk_write(m, v4, int32(2), v73, int32(1))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_smgr_bulk_finish(m, v4)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
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
	}
}
func F_spginsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v19 = F_AllocSetContextCreateInternal(m, v14, int32(59674), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(4515392)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v19
	F_initSpGistState(m, v11+int32(12), l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = F_spgdoinsert(m, l0, v11+int32(12), l3, l1, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v33 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_SpGistUpdateMetaPage(m, l0)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	F_MemoryContextReset(m, v19)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	F_initSpGistState(m, v11+int32(12), l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v53 = F_spgdoinsert(m, l0, v11+int32(12), l3, l1, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v53 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v24
	F_MemoryContextDelete(m, v19)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	m.G0 = v11 + int32(96)
	return int32(0)
}
func F_spgproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v21 = base.B2i32(l2 == int32(6)) & base.B2i32(l1 != int32(0))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = F_get_index_column_opclass(m, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v14 + int32(16)
	return v21
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v112)
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	if v22 == int32(0) {
		v112 = v16
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v32 = F_get_opclass_opfamily_and_input_type(m, v22, v14+int32(12), v14+int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v32 == int32(0) {
		v112 = v16
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v42 = F_SearchSysCacheList(m, int32(4), int32(1), v39, v36, v36)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v44)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v46 <= v44 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_ReleaseCatCacheList(m, v42)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L26
	}
L12:
	;
	v52 = v46
	v53 = int32(0)
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(48)+v53<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+18)))
	if v70 != int32(111) {
		v85 = v52
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v90)
	goto L11
L15:
	;
	goto L14
L16:
	;
	v88 = v53 + int32(1)
	if v88 < v85 {
		v52 = v85
		v53 = v88
		goto L13
	} else {
		goto L25
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v73 != v74 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v76 != v73 {
		v85 = v52
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v80 = F_get_op_rettype(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v82 = F_opfamily_can_sort_type(m, v78, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v82 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	v85 = v84
	goto L16
L25:
	;
	goto L11
L26:
	;
	v112 = v36
	goto L4
}
func F_stack_is_too_deep(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1293]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1294]))
	v7 = m.G0
	v10 = v6 - (v7 - int32(1))
	v12 = v10 >> (uint(int32(31)) % 32)
	return base.B2i32(v4 < v10^v12-v12) & base.B2i32(v6 != int32(0))
}
func F_statistic_proc_security_check(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v10 != 0 {
		v36 = v9
		m.G0 = v7 + int32(16)
		return v36
	} else {
		if l1 == int32(0) {
			v36 = int32(0)
			m.G0 = v7 + int32(16)
			return v36
		} else {
			v14 = F_get_func_leakproof(m, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v36 = v9
					m.G0 = v7 + int32(16)
					return v36
				} else {
					v18 = int32(0)
					v21 = F_errstart(m, int32(13), v18)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						if v21 == int32(0) {
							v36 = v18
							m.G0 = v7 + int32(16)
							return v36
						} else {
							v25 = F_get_func_name(m, l1)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v25
								F_errmsg_internal(m, int32(338345), v7)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494745), int32(6242), int32(318029))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										v36 = v18
										m.G0 = v7 + int32(16)
										return v36
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_storeGettuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	v4 = l3
	v6 = l5
	v7 = l6
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v19 = l0 + v16*int32(6)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+228)) = uint16(v20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v24+int32(2672)))) = uint8(v6)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v29+int32(3080)))) = uint8(v7)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if int32(0) < v34 {
		v37 = int32(0)
		if v4 != 0 {
			v94 = v37
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			*(*int32)(unsafe.Add(mBase, uint32(l0+v96<<(uint(int32(2))%32))+uint32(_consts[64]))) = v94
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			if v114 == int32(1) {
				v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				if int32(2) <= v118 {
					F_spgDeformLeafTuple(m, l4, v117, v14+int32(32), v14, v4)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
						v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
							m.G0 = v14 + int32(160)
							return
						}
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
					v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
						m.G0 = v14 + int32(160)
						return
					}
				}
			} else {
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
				m.G0 = v14 + int32(160)
				return
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			if v38 <= int32(0) {
				v94 = v37
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				*(*int32)(unsafe.Add(mBase, uint32(l0+v96<<(uint(int32(2))%32))+uint32(_consts[64]))) = v94
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
				if v114 == int32(1) {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
					if int32(2) <= v118 {
						F_spgDeformLeafTuple(m, l4, v117, v14+int32(32), v14, v4)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
							v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
								m.G0 = v14 + int32(160)
								return
							}
						}
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
						v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
							m.G0 = v14 + int32(160)
							return
						}
					}
				} else {
					v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
					m.G0 = v14 + int32(160)
					return
				}
			} else {
				v43 = F_palloc(m, v34<<(uint(int32(4))%32))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					if v45 <= int32(0) {
						v94 = v43
					} else {
						v50 = int32(0)
						for {
							v62 = v43 + v50<<(uint(int32(4))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v50<<(uint(int32(2))%32))))
							if v67 < int32(0) {
								v77 = float64(0)
								v78 = int32(1)
							} else {
								v75 = *(*float64)(unsafe.Add(mBase, uint32(l7+v67<<(uint(int32(3))%32))))
								v77 = v75
								v78 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v78)
							*(*float64)(unsafe.Add(mBase, uint32(v62))) = v77
							v82 = v50 + int32(1)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
							if v82 < v83 {
								v50 = v82
								continue
							} else {
								break
							}
							break
						}
						v94 = v43
					}
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0+v96<<(uint(int32(2))%32))+uint32(_consts[64]))) = v94
					v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
					if v114 == int32(1) {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
						if int32(2) <= v118 {
							F_spgDeformLeafTuple(m, l4, v117, v14+int32(32), v14, v4)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
								v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
									v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
									m.G0 = v14 + int32(160)
									return
								}
							}
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
							v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
								m.G0 = v14 + int32(160)
								return
							}
						}
					} else {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
						m.G0 = v14 + int32(160)
						return
					}
				}
			}
		}
	} else {
		v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
		if v114 == int32(1) {
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
			if int32(2) <= v118 {
				F_spgDeformLeafTuple(m, l4, v117, v14+int32(32), v14, v4)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
					v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
						m.G0 = v14 + int32(160)
						return
					}
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
				v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
				v130 = F_heap_form_tuple(m, v127, v14+int32(32), v14)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0+v132<<(uint(int32(2))%32)+int32(3488)))) = v130
					v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
					m.G0 = v14 + int32(160)
					return
				}
			}
		} else {
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v140 + int32(1)
			m.G0 = v14 + int32(160)
			return
		}
	}
}
func F_strict_word_similarity_dist_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v22 = int32(1)
			v23 = v21 & v22
			if v21 == v22 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v17
			} else {
				v52 = v10 + int32(4)
			}
			v53 = int32(1)
			v54 = v19 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v19 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(2))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
					}
				}
			}
		}
	}
}
func F_strtoint(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, int32(10), int64(2147483648))
	return base.I32_wrap_i64(v5)
}
func F_strtoull(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-1))
	return v5
}
func F_subcolor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	v8 = int32(65535)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(1))%32)))))
	v17 = v9 + v14*int32(24)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
	if v18 != v8 {
		v43 = v18
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
		if v45 != 0 {
			v88 = v8
			return base.I32_extend16_s(v88)
		} else {
			v46 = int32(65535)
			if v14&v46 == v43&v46 {
				return base.I32_extend16_s(v14)
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v54 = int32(24)
				v56 = v53 + v14*v54
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 - int32(1)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v64 = base.I32_extend16_s(v43) * v54
				v65 = v61 + v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				if v66 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = l1
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v64)))
					v73 = v70
					v74 = v72
				} else {
					v73 = v61
					v74 = v66
				}
				v76 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v73+v64))) = v74 + v76
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*uint16)(unsafe.Add(mBase, uint32(v79+l1<<(uint(v76)%32)))) = uint16(v43)
				v88 = v43
				return base.I32_extend16_s(v88)
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v21+v22 == int32(1) {
			v43 = v14
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
			if v45 != 0 {
				v88 = v8
				return base.I32_extend16_s(v88)
			} else {
				v46 = int32(65535)
				if v14&v46 == v43&v46 {
					return base.I32_extend16_s(v14)
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v54 = int32(24)
					v56 = v53 + v14*v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 - int32(1)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v64 = base.I32_extend16_s(v43) * v54
					v65 = v61 + v64
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					if v66 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = l1
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v64)))
						v73 = v70
						v74 = v72
					} else {
						v73 = v61
						v74 = v66
					}
					v76 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v73+v64))) = v74 + v76
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*uint16)(unsafe.Add(mBase, uint32(v79+l1<<(uint(v76)%32)))) = uint16(v43)
					v88 = v43
					return base.I32_extend16_s(v88)
				}
			}
		} else {
			v26 = F_newcolor(m, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 == int32(-1) {
					v43 = int32(65535)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v34 = int32(24)
					*(*uint16)(unsafe.Add(mBase, uint32(v33+v14*v34)+8)) = uint16(v26)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*uint16)(unsafe.Add(mBase, uint32(v38+v26*v34)+8)) = uint16(v26)
					v43 = v26
				}
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
				if v45 != 0 {
					v88 = v8
					return base.I32_extend16_s(v88)
				} else {
					v46 = int32(65535)
					if v14&v46 == v43&v46 {
						return base.I32_extend16_s(v14)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v54 = int32(24)
						v56 = v53 + v14*v54
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 - int32(1)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v64 = base.I32_extend16_s(v43) * v54
						v65 = v61 + v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						if v66 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = l1
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v64)))
							v73 = v70
							v74 = v72
						} else {
							v73 = v61
							v74 = v66
						}
						v76 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v73+v64))) = v74 + v76
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						*(*uint16)(unsafe.Add(mBase, uint32(v79+l1<<(uint(v76)%32)))) = uint16(v43)
						v88 = v43
						return base.I32_extend16_s(v88)
					}
				}
			}
		}
	}
}
func F_subcoloronerow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v6 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	if v13 <= v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v29 = v16 + l1*v13<<(uint(int32(1))%32)
	v31 = v6
	goto L3
L3:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29))))
	v34 = v32 * int32(24)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v36 = v34 + v35
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != int32(65535) {
		v58 = v37
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v61 != 0 {
		v88 = int32(65535)
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v40+v41 == int32(1) {
		v58 = v32
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = F_newcolor(m, v12)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v45 == int32(-1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v58 = int32(65535)
	goto L5
L11:
	;
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v50+v34)+8)) = uint16(v45)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v53+v45*int32(24))+8)) = uint16(v45)
	v58 = v45
	goto L5
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v62 = int32(65535)
	if v32&v62 == v58&v62 {
		v88 = v32
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v69 = int32(4)
	v70 = v67 + v34 + v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 - v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v81 = v75 + base.I32_extend16_s(v58)*int32(24) + v69
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 + v72
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v58)
	v88 = v58
	goto L13
L16:
	;
	v91 = v88 & int32(65535)
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	if v91 != v92 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v96 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v96 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v188 = v31 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	if v188 < v189 {
		v29 = v29 + int32(2)
		v31 = v188
		goto L3
	} else {
		goto L47
	}
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v99 <= v100 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v172 != 0 {
		goto L1
	} else {
		goto L46
	}
L25:
	;
	F_createarc(m, v94, int32(112), base.I32_extend16_s(v88), l2, l3)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L45
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v102 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v124 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L29:
	;
	v106 = v102
	goto L30
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	if v116 != l3 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L25
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if v123 != 0 {
		v106 = v123
		goto L30
	} else {
		goto L36
	}
L33:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)))
	if v118 != v91 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v120 == int32(112) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	goto L31
L37:
	;
	v128 = v124
	goto L38
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v138 != l2 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L25
L40:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v128)+24))
	if v145 != 0 {
		v128 = v145
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+4)))
	if v140 != v91 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v142 == int32(112) {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L39
L45:
	;
	goto L24
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v88)
	goto L19
L47:
	;
	goto L4
}
func F_svals(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_svals(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_syncrep_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(16384)
			v15 = F_palloc(m, int32(16386))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_4(m, int32(683723))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[86])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_4(m, int32(683723))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_syncrep_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_4(m, int32(684000))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_4(m, int32(684000))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
