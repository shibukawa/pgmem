package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateActiveSnapshotCommandId(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateActiveSnapshotCommandId[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	v8 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateActiveSnapshotCommandId[1]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
		if v13 != 0 {
			v16 = int32(1)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+76)))
			v16 = v15
		}
		v19 = int32(0)
		if base.B2i32(v16&int32(1) == v19)|base.B2i32(v8 == v6) == v19 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_UpdateActiveSnapshotCommandId_0), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_UpdateActiveSnapshotCommandId_1), int32(764), int32(_a_F_UpdateActiveSnapshotCommandId_2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_UpdateActiveSnapshotCommandId[0]))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v8
			return
		}
	}
}
func F_UtfToLocal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	if base.B2i32(l7 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l7)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v767 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v747))) = uint8(v767)
	m.G0 = v25 + int32(48)
	return v755 - l0
L2:
	;
	v55 = l1
	v56 = l2
	v64 = l0
	goto L12
L3:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v747 = l2
	v755 = l0
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l7
	F_errmsg(m, int32(_a_F_UtfToLocal_0), v25)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_UtfToLocal_1), int32(284), int32(_a_F_UtfToLocal_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v76 == int32(0) {
		v240 = v55
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v747 = v730
	v755 = v741
	goto L1
L14:
	;
	v742 = v731 - v733
	if int32(0) < v742 {
		v55 = v742
		v56 = v730
		v64 = v741
		goto L12
	} else {
		goto L197
	}
L15:
	;
	if l3 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L16:
	;
	v256 = int32(0)
	switch v238 - int32(1) {
	case 0:
		goto L84
	case 1:
		goto L85
	case 2:
		goto L86
	case 3:
		goto L87
	default:
		v305 = v256
		goto L81
	}
L17:
	;
	if l8|base.B2i32(v240 <= int32(0)) != 0 {
		v747 = v56
		v755 = v64
		goto L1
	} else {
		goto L78
	}
L18:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
	if int32(0) <= v79 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v55 < v103 {
		v240 = v55
		goto L17
	} else {
		goto L32
	}
L20:
	;
	v103 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v84 = v79 & int32(255)
	if v84&int32(224) == int32(192) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v103 = int32(2)
	goto L19
L24:
	;
	goto L25
L25:
	;
	if v84&int32(240) == int32(224) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v103 = int32(3)
	goto L19
L27:
	;
	goto L28
L28:
	;
	if v84&int32(248) == int32(240) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = int32(4)
	goto L31
L30:
	;
	v101 = int32(1)
	goto L31
L31:
	;
	v103 = v101
	goto L19
L32:
	;
	v105 = int32(0)
	switch v103 - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L38
	case 2:
		goto L39
	case 3:
		goto L40
	default:
		v154 = v105
		goto L34
	}
L33:
	;
	if v154 == int32(0) {
		v240 = v55
		goto L17
	} else {
		goto L54
	}
L34:
	;
	goto L33
L35:
	;
	v154 = base.B2i32(base.Ui32(v146&int32(255)) < base.Ui32(int32(245)))
	goto L34
L36:
	;
	if base.I32_extend8_s(v141) < int32(-62) {
		v154 = v105
		goto L34
	} else {
		goto L53
	}
L37:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v141 = v140
	goto L36
L38:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+1)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	switch v115 - int32(224) {
	case 0:
		goto L47
	default:
		goto L43
	case 13:
		goto L46
	case 16:
		goto L45
	case 20:
		goto L44
	}
L39:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+2)))
	if int32(-65) < v111 {
		v154 = v105
		goto L34
	} else {
		goto L42
	}
L40:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+3)))
	if int32(-65) < v108 {
		v154 = v105
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L38
L43:
	;
	if v114 <= int32(-65) {
		v141 = v115
		goto L36
	} else {
		goto L52
	}
L44:
	;
	if int32(-113) < v114 {
		v154 = v105
		goto L34
	} else {
		goto L51
	}
L45:
	;
	if base.Ui32((v114-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v154 = v105
		goto L34
	} else {
		goto L50
	}
L46:
	;
	if int32(-97) < v114 {
		v154 = v105
		goto L34
	} else {
		goto L49
	}
L47:
	;
	v118 = int32(224)
	if base.Ui32(v118) <= base.Ui32((v114-int32(-64))&int32(255)) {
		v146 = v118
		goto L35
	} else {
		goto L48
	}
L48:
	;
	v154 = v105
	goto L34
L49:
	;
	v146 = int32(237)
	goto L35
L50:
	;
	v146 = int32(240)
	goto L35
L51:
	;
	v146 = int32(244)
	goto L35
L52:
	;
	v154 = v105
	goto L34
L53:
	;
	v146 = v141
	goto L35
L54:
	;
	v157 = int32(0)
	v158 = int32(1)
	switch v103 - v158 {
	case 0:
		goto L59
	case 1:
		v192 = v64
		v193 = v158
		v194 = v157
		v195 = v157
		goto L55
	case 2:
		goto L56
	case 3:
		goto L58
	default:
		goto L57
	}
L55:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v193))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v207 = v197 | (v195<<(uint(int32(24))%32) | v194<<(uint(int32(16))%32) | v203<<(uint(int32(8))%32))
	v208 = v64 + v103
	if base.B2i32(l4 == int32(0))|base.B2i32(base.Ui32(v55) <= base.Ui32(v103)) != 0 {
		goto L15
	} else {
		goto L63
	}
L56:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v192 = v64 + int32(1)
	v193 = int32(2)
	v194 = v190
	v195 = v157
	goto L55
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v192 = v64 + int32(2)
	v193 = int32(3)
	v194 = v170
	v195 = v171
	goto L55
L59:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v162)
	v164 = int32(1)
	v730 = v56 + v164
	v731 = v55
	v733 = v158
	v741 = v64 + v164
	goto L14
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v103
	F_errmsg_internal(m, int32(_a_F_UtfToLocal_3), v25+int32(16))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_UtfToLocal_1), int32(332), int32(_a_F_UtfToLocal_2))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v213 = v55 - v103
	v214 = int32(*(*int8)(unsafe.Add(mBase, uint32(v208))))
	if int32(0) <= v214 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v238 <= v213 {
		goto L16
	} else {
		goto L77
	}
L65:
	;
	v238 = int32(1)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v219 = v214 & int32(255)
	if v219&int32(224) == int32(192) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v238 = int32(2)
	goto L64
L69:
	;
	goto L70
L70:
	;
	if v219&int32(240) == int32(224) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v238 = int32(3)
	goto L64
L72:
	;
	goto L73
L73:
	;
	if v219&int32(248) == int32(240) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v236 = int32(4)
	goto L76
L75:
	;
	v236 = int32(1)
	goto L76
L76:
	;
	v238 = v236
	goto L64
L77:
	;
	v240 = v213
	goto L17
L78:
	;
	F_report_invalid_encoding(m, int32(6), v64, v240)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	if v305 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L81:
	;
	goto L80
L82:
	;
	v305 = base.B2i32(base.Ui32(v297&int32(255)) < base.Ui32(int32(245)))
	goto L81
L83:
	;
	if base.I32_extend8_s(v292) < int32(-62) {
		v305 = v256
		goto L81
	} else {
		goto L100
	}
L84:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v292 = v291
	goto L83
L85:
	;
	v265 = int32(*(*int8)(unsafe.Add(mBase, uint32(v208)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	switch v266 - int32(224) {
	case 0:
		goto L94
	default:
		goto L90
	case 13:
		goto L93
	case 16:
		goto L92
	case 20:
		goto L91
	}
L86:
	;
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v208)+2)))
	if int32(-65) < v262 {
		v305 = v256
		goto L81
	} else {
		goto L89
	}
L87:
	;
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v208)+3)))
	if int32(-65) < v259 {
		v305 = v256
		goto L81
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	goto L85
L90:
	;
	if v265 <= int32(-65) {
		v292 = v266
		goto L83
	} else {
		goto L99
	}
L91:
	;
	if int32(-113) < v265 {
		v305 = v256
		goto L81
	} else {
		goto L98
	}
L92:
	;
	if base.Ui32((v265-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v305 = v256
		goto L81
	} else {
		goto L97
	}
L93:
	;
	if int32(-97) < v265 {
		v305 = v256
		goto L81
	} else {
		goto L96
	}
L94:
	;
	v269 = int32(224)
	if base.Ui32(v269) <= base.Ui32((v265-int32(-64))&int32(255)) {
		v297 = v269
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v305 = v256
	goto L81
L96:
	;
	v297 = int32(237)
	goto L82
L97:
	;
	v297 = int32(240)
	goto L82
L98:
	;
	v297 = int32(244)
	goto L82
L99:
	;
	v305 = v256
	goto L81
L100:
	;
	v297 = v292
	goto L82
L101:
	;
	if l8 != 0 {
		v747 = v56
		v755 = v64
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v238 < int32(2) {
		goto L15
	} else {
		goto L106
	}
L104:
	;
	F_report_invalid_encoding(m, int32(6), v208, v213)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	switch v238 - int32(2) {
	case 0:
		goto L108
	case 1:
		goto L111
	case 2:
		goto L110
	default:
		goto L109
	}
L107:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v359 | v358
	v367 = F_bsearch(m, v25+int32(40), l4, l5, int32(12), int32(1849))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L115
	}
L108:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v357 = v208 + int32(1)
	v358 = v354 << (uint(int32(8)) % 32)
	goto L107
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
	v357 = v208 + int32(3)
	v358 = v326<<(uint(int32(16))%32) | v329<<(uint(int32(24))%32) | v333<<(uint(int32(8))%32)
	goto L107
L111:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v357 = v208 + int32(2)
	v358 = v317<<(uint(int32(8))%32) | v320<<(uint(int32(16))%32)
	goto L107
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v238
	F_errmsg_internal(m, int32(_a_F_UtfToLocal_3), v25+int32(32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_UtfToLocal_1), int32(389), int32(_a_F_UtfToLocal_2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	if v367 == int32(0) {
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	if base.Ui32(int32(16777216)) <= base.Ui32(v371) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v375 = int32(base.Ui32(v371) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v375)
	v379 = v56 + int32(1)
	goto L119
L118:
	;
	v379 = v56
	goto L119
L119:
	;
	if v371&int32(16711680) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v383 = int32(base.Ui32(v371) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v379))) = uint8(v383)
	v387 = v379 + int32(1)
	goto L122
L121:
	;
	v387 = v379
	goto L122
L122:
	;
	if v371&int32(_a_F_UtfToLocal_4) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v391 = int32(base.Ui32(v371) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v391)
	v395 = v387 + int32(1)
	goto L125
L124:
	;
	v395 = v387
	goto L125
L125:
	;
	if v371&int32(255) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v395))) = uint8(v371)
	v401 = v395 + int32(1)
	goto L128
L127:
	;
	v401 = v395
	goto L128
L128:
	;
	v730 = v401
	v731 = v213
	v733 = v238
	v741 = v208 + v238
	goto L14
L129:
	;
	v730 = v727
	v731 = v55
	v733 = v103
	v741 = v208
	goto L14
L130:
	;
	if l6 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L131:
	;
	v409 = int32(0)
	switch v103 - int32(1) {
	case 0:
		goto L134
	case 1:
		goto L135
	case 2:
		goto L136
	case 3:
		goto L137
	default:
		v640 = v409
		goto L133
	}
L132:
	;
	if v651 == int32(0) {
		goto L130
	} else {
		goto L170
	}
L133:
	;
	v651 = v640
	goto L132
L134:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(v197) < base.Ui32(v611) {
		v640 = v409
		goto L133
	} else {
		goto L165
	}
L135:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if base.Ui32(v203) < base.Ui32(v566) {
		v640 = v409
		goto L133
	} else {
		goto L158
	}
L136:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if base.Ui32(v194) < base.Ui32(v501) {
		v640 = v409
		goto L133
	} else {
		goto L149
	}
L137:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	if base.Ui32(v195) < base.Ui32(v416) {
		v640 = v409
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+41)))
	if base.Ui32(v418) < base.Ui32(v195) {
		v640 = v409
		goto L133
	} else {
		goto L139
	}
L139:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+42)))
	if base.Ui32(v194) < base.Ui32(v420) {
		v640 = v409
		goto L133
	} else {
		goto L140
	}
L140:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+43)))
	if base.Ui32(v422) < base.Ui32(v194) {
		v640 = v409
		goto L133
	} else {
		goto L141
	}
L141:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+44)))
	if base.Ui32(v203) < base.Ui32(v424) {
		v640 = v409
		goto L133
	} else {
		goto L142
	}
L142:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if base.Ui32(v426) < base.Ui32(v203) {
		v640 = v409
		goto L133
	} else {
		goto L143
	}
L143:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if base.Ui32(v197) < base.Ui32(v428) {
		v640 = v409
		goto L133
	} else {
		goto L144
	}
L144:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+47)))
	if base.Ui32(v430) < base.Ui32(v197) {
		v640 = v409
		goto L133
	} else {
		goto L145
	}
L145:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v433 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v435 = int32(2)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v433+(v195-v416)<<(uint(v435)%32)+v432<<(uint(v435)%32))))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v433+(v194-v420)<<(uint(v435)%32)+v453<<(uint(v435)%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v433+(v203-v424)<<(uint(v435)%32)+v457<<(uint(v435)%32))))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v433+(v197-v428)<<(uint(v435)%32)+v461<<(uint(v435)%32))))
	v651 = v465
	goto L132
L147:
	;
	goto L148
L148:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v468 = int32(1)
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466+(v195-v416)<<(uint(v468)%32)+v432&int32(_a_F_UtfToLocal_5)<<(uint(v468)%32)))))
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466+(v194-v420)<<(uint(v468)%32)+v488<<(uint(v468)%32)))))
	v496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466+(v203-v424)<<(uint(v468)%32)+v492<<(uint(v468)%32)))))
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466+(v197-v428)<<(uint(v468)%32)+v496<<(uint(v468)%32)))))
	v651 = v500
	goto L132
L149:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)))
	if base.Ui32(v503) < base.Ui32(v194) {
		v640 = v409
		goto L133
	} else {
		goto L150
	}
L150:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+30)))
	if base.Ui32(v203) < base.Ui32(v505) {
		v640 = v409
		goto L133
	} else {
		goto L151
	}
L151:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+31)))
	if base.Ui32(v507) < base.Ui32(v203) {
		v640 = v409
		goto L133
	} else {
		goto L152
	}
L152:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if base.Ui32(v197) < base.Ui32(v509) {
		v640 = v409
		goto L133
	} else {
		goto L153
	}
L153:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+33)))
	if base.Ui32(v511) < base.Ui32(v197) {
		v640 = v409
		goto L133
	} else {
		goto L154
	}
L154:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v514 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v516 = int32(2)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v514+(v194-v501)<<(uint(v516)%32)+v513<<(uint(v516)%32))))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v514+(v203-v505)<<(uint(v516)%32)+v530<<(uint(v516)%32))))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v514+(v197-v509)<<(uint(v516)%32)+v534<<(uint(v516)%32))))
	v651 = v538
	goto L132
L156:
	;
	goto L157
L157:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v541 = int32(1)
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539+(v194-v501)<<(uint(v541)%32)+v513&int32(_a_F_UtfToLocal_5)<<(uint(v541)%32)))))
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539+(v203-v505)<<(uint(v541)%32)+v557<<(uint(v541)%32)))))
	v565 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539+(v197-v509)<<(uint(v541)%32)+v561<<(uint(v541)%32)))))
	v651 = v565
	goto L132
L158:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if base.Ui32(v568) < base.Ui32(v203) {
		v640 = v409
		goto L133
	} else {
		goto L159
	}
L159:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)))
	if base.Ui32(v197) < base.Ui32(v570) {
		v640 = v409
		goto L133
	} else {
		goto L160
	}
L160:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)))
	if base.Ui32(v572) < base.Ui32(v197) {
		v640 = v409
		goto L133
	} else {
		goto L161
	}
L161:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v575 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v577 = int32(2)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v575+(v203-v566)<<(uint(v577)%32)+v574<<(uint(v577)%32))))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v575+(v197-v570)<<(uint(v577)%32)+v587<<(uint(v577)%32))))
	v651 = v591
	goto L132
L163:
	;
	goto L164
L164:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v594 = int32(1)
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592+(v203-v566)<<(uint(v594)%32)+v574&int32(_a_F_UtfToLocal_5)<<(uint(v594)%32)))))
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592+(v197-v570)<<(uint(v594)%32)+v606<<(uint(v594)%32)))))
	v651 = v610
	goto L132
L165:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(v613) < base.Ui32(v197) {
		v640 = v409
		goto L133
	} else {
		goto L166
	}
L166:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v615 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v617 = int32(2)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v615+(v197-v611)<<(uint(v617)%32)+v620<<(uint(v617)%32))))
	v651 = v624
	goto L132
L168:
	;
	goto L169
L169:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v627 = int32(1)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625+(v197-v611)<<(uint(v627)%32)+v630<<(uint(v627)%32)))))
	v640 = v634
	goto L133
L170:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v651) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v657 = int32(base.Ui32(v651) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v657)
	v661 = v56 + int32(1)
	goto L173
L172:
	;
	v661 = v56
	goto L173
L173:
	;
	if v651&int32(16711680) != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v665 = int32(base.Ui32(v651) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v661))) = uint8(v665)
	v669 = v661 + int32(1)
	goto L176
L175:
	;
	v669 = v661
	goto L176
L176:
	;
	if v651&int32(_a_F_UtfToLocal_4) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v673 = int32(base.Ui32(v651) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v669))) = uint8(v673)
	v677 = v669 + int32(1)
	goto L179
L178:
	;
	v677 = v669
	goto L179
L179:
	;
	if v651&int32(255) == int32(0) {
		v727 = v677
		goto L129
	} else {
		goto L180
	}
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v677))) = uint8(v651)
	v727 = v677 + int32(1)
	goto L129
L181:
	;
	if l8 != 0 {
		v747 = v56
		v755 = v64
		goto L1
	} else {
		goto L195
	}
L182:
	;
	v688 = m.T0[l6].(func(*base.Module, int32) int32)(m, v207)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L7
	} else {
		goto L183
	}
L183:
	;
	if v688 == int32(0) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v688) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v695 = int32(base.Ui32(v688) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v695)
	v699 = v56 + int32(1)
	goto L187
L186:
	;
	v699 = v56
	goto L187
L187:
	;
	if v688&int32(16711680) != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v703 = int32(base.Ui32(v688) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v699))) = uint8(v703)
	v707 = v699 + int32(1)
	goto L190
L189:
	;
	v707 = v699
	goto L190
L190:
	;
	if v688&int32(_a_F_UtfToLocal_4) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v711 = int32(base.Ui32(v688) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v707))) = uint8(v711)
	v715 = v707 + int32(1)
	goto L193
L192:
	;
	v715 = v707
	goto L193
L193:
	;
	if v688&int32(255) == int32(0) {
		v727 = v715
		goto L129
	} else {
		goto L194
	}
L194:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v715))) = uint8(v688)
	v727 = v715 + int32(1)
	goto L129
L195:
	;
	F_report_untranslatable_char(m, int32(6), l7, v64, v55)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	goto L13
}
func F___uflow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(-1)
	v9 = F___toread(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v22 = v8
			m.G0 = v6 + int32(16)
			return v22
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = m.T0[v16].(func(*base.Module, int32, int32, int32) int32)(m, l0, v6+int32(15), int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != int32(1) {
					v22 = v8
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
					v22 = v21
				}
				m.G0 = v6 + int32(16)
				return v22
			}
		}
	}
}
func F_uniq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		if v7 != 0 {
			v8 = F_array_contains_nulls(m, v3)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return int64(0)
			} else {
				if v8 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_uniq_0), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_uniq_1), int32(255), int32(_a_F_uniq_2))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
					v13 = F_ArrayGetNItemsSafe(m, v10, v3+int32(16))
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return int64(0)
					} else {
						if int32(2) <= v13 {
							v17 = F__int_unique(m, v3)
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return int64(0)
							} else {
								v19 = v17
								return base.I64_extend_i32_u(v19)
							}
						} else {
							v19 = v3
							return base.I64_extend_i32_u(v19)
						}
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			v13 = F_ArrayGetNItemsSafe(m, v10, v3+int32(16))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				if int32(2) <= v13 {
					v17 = F__int_unique(m, v3)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = v17
						return base.I64_extend_i32_u(v19)
					}
				} else {
					v19 = v3
					return base.I64_extend_i32_u(v19)
				}
			}
		}
	}
}
func F_update_frameheadpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int64
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v152 int64
	_ = v152
	var v157 int64
	_ = v157
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v245 int32
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
	var v254 int32
	_ = v254
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
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v278 int32
	_ = v278
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v304 int64
	_ = v304
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v413 int32
	_ = v413
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v447 int32
	_ = v447
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+388)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L22
	} else {
		goto L144
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L22
	} else {
		goto L141
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L22
	} else {
		goto L138
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = int32(_a_F_update_frameheadpos_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v28
	if v22&int32(32) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v17 + int32(16)
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v25
	goto L6
L8:
	;
	v447 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+388)) = uint8(v447)
	goto L7
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v22&int32(512) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v22&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v22&int32(_a_F_update_frameheadpos_1) == int32(0) {
		goto L7
	} else {
		goto L47
	}
L15:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v38
	goto L8
L16:
	;
	goto L17
L17:
	;
	if v22&int32(10) == int32(0) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	if v44 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = int64(0)
	goto L8
L20:
	;
	goto L21
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v53 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	goto L32
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v57&int32(2) == int32(0) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v63 = int32(1)
	v65 = F_tuplestore_gettupleslot(m, v62, v63, v63, v56)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v65 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v84 == int32(0) {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	goto L8
L34:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v87&int32(2) != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+96))
	if v91 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v84
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v98 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	F_MemoryContextReset(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L22
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v104 = int32(_a_F_update_frameheadpos_0)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	v112 = m.T0[v111].(func(*base.Module, int32, int32, int32) int64)(m, v98, v94, v17+int32(14))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L22
	} else {
		goto L41
	}
L40:
	;
	goto L8
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v105
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	F_MemoryContextReset(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	if v112 != int64(0) {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v123 = v121 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v123
	F_spool_tuples(m, l0, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v128 = int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v131 = F_tuplestore_gettupleslot(m, v127, v128, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	if v131 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	goto L33
L47:
	;
	if v22&int32(4) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	if v22&int32(2048) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v22&int32(2) != 0 {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	v145 = int64(0) - v141
	goto L53
L52:
	;
	v145 = v141
	goto L53
L53:
	;
	v146 = v139 + v145
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v146
	if base.B2i32(v145 < int64(0)) != base.B2i32(v146 < v139) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v169
	goto L8
L55:
	;
	if v157 <= v139+int64(1) {
		goto L8
	} else {
		goto L60
	}
L56:
	;
	v152 = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v152
	v157 = v152
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v146 < int64(0) {
		v169 = int64(0)
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v157 = v146
	goto L55
L60:
	;
	F_spool_tuples(m, l0, v157-int64(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v166 <= v165 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v169 = v165
	goto L54
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+316)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v176, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v22&int32(8) == int32(0) {
		goto L7
	} else {
		goto L101
	}
L66:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v180 != int64(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v199 = int32(1)
	v202 = int64(1)
	v213 = v174 - v199
	v215 = v213 << (uint(int32(3)) % 32)
	goto L75
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v183 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
	if v184&int32(2) == int32(0) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v190 = int32(1)
	v192 = F_tuplestore_gettupleslot(m, v189, v190, v190, v183)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L22
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	if v192 == int32(0) {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	goto L67
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v230 == int32(0) {
		goto L8
	} else {
		goto L77
	}
L76:
	;
	goto L8
L77:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+4)))
	if v233&int32(2) != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v230)+6)))
	if v236 < v174 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	m.T0[v239].(func(*base.Module, int32, int32))(m, v230, v174)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L22
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v242+v215)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v213))))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+6)))
	if v249 < v174 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	m.T0[v252].(func(*base.Module, int32, int32))(m, v248, v174)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L22
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255+v213))))
	if v247|v257 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L85
L87:
	;
	v282 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v284 = v282 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v284
	F_spool_tuples(m, l0, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L98
	}
L88:
	;
	v259 = int32(1)
	v260 = v247 ^ v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v261 == v259 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v273+v215)))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	v277 = F_FunctionCall5Coll(m, l0+int32(256), v272, v244, v275, v276, base.I64_extend_i32_u(base.B2i32(v22&int32(2048) == int32(0))^v175)&v202, base.I64_extend_i32_u(v175^v199)&v202)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L22
	} else {
		goto L96
	}
L91:
	;
	if (v260|v257)&int32(1) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v260&v257&int32(1) != 0 {
		goto L87
	} else {
		goto L95
	}
L94:
	;
	goto L8
L95:
	;
	goto L8
L96:
	;
	if v277 != int64(0) {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	goto L87
L98:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v289 = int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v292 = F_tuplestore_gettupleslot(m, v288, v289, v289, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	if v292 != 0 {
		goto L75
	} else {
		goto L100
	}
L100:
	;
	goto L76
L101:
	;
	v298 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	v299 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	if v22&int32(2048) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v311 = v298 - v299
	goto L104
L103:
	;
	v304 = v298 + v299
	if base.B2i32(v299 < int64(0))^base.B2i32(v304 < v298) != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_tuplestore_select_read_pointer(m, v312, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L108
	}
L105:
	;
	v309 = int64(9223372036854775807)
	goto L107
L106:
	;
	v309 = v304
	goto L107
L107:
	;
	v311 = v309
	goto L104
L108:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	if v316 != int64(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v333 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L110:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v319 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+4)))
	if v320&int32(2) == int32(0) {
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v326 = int32(1)
	v328 = F_tuplestore_gettupleslot(m, v325, v326, v326, v319)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L22
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	if v328 == int32(0) {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L109
L117:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	m.T0[v430].(func(*base.Module, int32))(m, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L22
	} else {
		goto L137
	}
L118:
	;
	v337 = v333
	goto L119
L119:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+4)))
	if v350&int32(2) != 0 {
		goto L117
	} else {
		goto L121
	}
L120:
	;
	goto L117
L121:
	;
	v353 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	if v311 <= v353 {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+32))
	m.T0[v357].(func(*base.Module, int32, int32))(m, v355, v337)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	v362 = v360 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v362
	F_spool_tuples(m, l0, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v367 = int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v370 = F_tuplestore_gettupleslot(m, v366, v367, v367, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	if v370 == int32(0) {
		goto L117
	} else {
		goto L126
	}
L126:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+96))
	if v375 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v413 != 0 {
		v337 = v413
		goto L119
	} else {
		goto L136
	}
L128:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = v378
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v383 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	F_MemoryContextReset(m, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L22
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v389 = int32(_a_F_update_frameheadpos_0)
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v392
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	v397 = m.T0[v396].(func(*base.Module, int32, int32, int32) int64)(m, v383, v379, v17+int32(15))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L22
	} else {
		goto L133
	}
L132:
	;
	goto L127
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frameheadpos[0])) = v390
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	F_MemoryContextReset(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L22
	} else {
		goto L134
	}
L134:
	;
	if v397 != int64(0) {
		goto L127
	} else {
		goto L135
	}
L135:
	;
	v406 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v406 + int64(1)
	goto L127
L136:
	;
	goto L120
L137:
	;
	goto L8
L138:
	;
	F_errmsg_internal(m, int32(_a_F_update_frameheadpos_2), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L22
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_update_frameheadpos_3), int32(1671), int32(_a_F_update_frameheadpos_4))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L22
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errmsg_internal(m, int32(_a_F_update_frameheadpos_2), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L22
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_update_frameheadpos_3), int32(1761), int32(_a_F_update_frameheadpos_4))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L22
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errmsg_internal(m, int32(_a_F_update_frameheadpos_2), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L22
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_update_frameheadpos_3), int32(1846), int32(_a_F_update_frameheadpos_4))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L22
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_update_frametailpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v40 int64
	_ = v40
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v166 int64
	_ = v166
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int64
	_ = v218
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v289 int32
	_ = v289
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v314 int64
	_ = v314
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v362 int64
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int64
	_ = v415
	var v422 int32
	_ = v422
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v454 int32
	_ = v454
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+389)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L12
	} else {
		goto L150
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L12
	} else {
		goto L147
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L12
	} else {
		goto L144
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = int32(_a_F_update_frametailpos_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v27
	if v21&int32(256) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v16 + int32(16)
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v24
	goto L6
L8:
	;
	v454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+389)) = uint8(v454)
	goto L7
L9:
	;
	F_spool_tuples(m, l0, int64(-1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v21&int32(1024) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v34
	goto L8
L14:
	;
	if v21&int32(4) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v21&int32(_a_F_update_frametailpos_1) == int32(0) {
		goto L7
	} else {
		goto L51
	}
L17:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v40 + int64(1)
	goto L8
L18:
	;
	goto L19
L19:
	;
	if v21&int32(10) == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	if v48 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_spool_tuples(m, l0, int64(-1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v54
	goto L8
L25:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v60 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	goto L34
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v63 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v64&int32(2) == int32(0) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v70 = int32(1)
	v72 = F_tuplestore_gettupleslot(m, v69, v70, v70, v63)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v72 == int32(0) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v90 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L35:
	;
	goto L8
L36:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v93&int32(2) != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	if v96 <= v97 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v135 = v133 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v135
	F_spool_tuples(m, l0, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L48
	}
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	if v100 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = v90
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v107 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	F_MemoryContextReset(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v113 = int32(_a_F_update_frametailpos_0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v107)+24))
	v121 = m.T0[v120].(func(*base.Module, int32, int32, int32) int64)(m, v107, v103, v16+int32(14))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L45
	}
L44:
	;
	goto L38
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v114
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	F_MemoryContextReset(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	if v121 == int64(0) {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	goto L38
L48:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v140 = int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v143 = F_tuplestore_gettupleslot(m, v139, v140, v140, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if v143 != 0 {
		goto L34
	} else {
		goto L50
	}
L50:
	;
	goto L35
L51:
	;
	if v21&int32(4) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v21&int32(_a_F_update_frametailpos_2) != 0 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	goto L54
L54:
	;
	if v21&int32(2) != 0 {
		goto L69
	} else {
		goto L70
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v189
	goto L8
L56:
	;
	if v176 <= v159+int64(1) {
		goto L8
	} else {
		goto L66
	}
L57:
	;
	v173 = int64(0)
	if v166 < v173 {
		v189 = v173
		goto L55
	} else {
		goto L65
	}
L58:
	;
	v156 = int64(0) - v152
	goto L60
L59:
	;
	v156 = v152
	goto L60
L60:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
	v160 = v159 + v156
	if base.B2i32(v156 < int64(0))^base.B2i32(v160 < v159) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v166 = v160 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v166
	if v160 <= v166 {
		goto L57
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v170 = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v170
	v176 = v170
	goto L56
L64:
	;
	goto L63
L65:
	;
	v176 = v166
	goto L56
L66:
	;
	F_spool_tuples(m, l0, v176-int64(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v186 <= v185 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v189 = v185
	goto L55
L69:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v194 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+316)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v196, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L12
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v21&int32(8) == int32(0) {
		goto L7
	} else {
		goto L107
	}
L72:
	;
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v200 != int64(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v218 = int64(1)
	v231 = v194 - int32(1)
	v233 = v231 << (uint(int32(3)) % 32)
	goto L81
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v203 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)))
	if v204&int32(2) == int32(0) {
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v210 = int32(1)
	v212 = F_tuplestore_gettupleslot(m, v209, v210, v210, v203)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v212 == int32(0) {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	goto L73
L81:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v247 == int32(0) {
		goto L8
	} else {
		goto L83
	}
L82:
	;
	goto L8
L83:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+4)))
	if v250&int32(2) != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v247)+6)))
	if v253 < v194 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	m.T0[v256].(func(*base.Module, int32, int32))(m, v247, v194)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L12
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v259+v233)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v231))))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(v265)+6)))
	if v266 < v194 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	m.T0[v269].(func(*base.Module, int32, int32))(m, v265, v194)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v231))))
	if v264|v274 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L91
L93:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v294 = v292 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v294
	F_spool_tuples(m, l0, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L104
	}
L94:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v276 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v284+v233)))
	v287 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	v288 = F_FunctionCall5Coll(m, l0+int32(284), v283, v261, v286, v287, base.I64_extend_i32_u(base.B2i32(v21&int32(_a_F_update_frametailpos_2) == int32(0))^v195)&v218, base.I64_extend_i32_u(v195)&v218)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L102
	}
L97:
	;
	if v264&int32(1) != 0 {
		goto L93
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	if v274&int32(1) != 0 {
		goto L93
	} else {
		goto L101
	}
L100:
	;
	goto L8
L101:
	;
	goto L8
L102:
	;
	if v288 == int64(0) {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	goto L93
L104:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v299 = int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v302 = F_tuplestore_gettupleslot(m, v298, v299, v299, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	if v302 != 0 {
		goto L81
	} else {
		goto L106
	}
L106:
	;
	goto L82
L107:
	;
	v308 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v21&int32(_a_F_update_frametailpos_2) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v321 = v308 - v309
	goto L110
L109:
	;
	v314 = v308 + v309
	if base.B2i32(v309 < int64(0))^base.B2i32(v314 < v308) != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_tuplestore_select_read_pointer(m, v322, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L114
	}
L111:
	;
	v319 = int64(9223372036854775807)
	goto L113
L112:
	;
	v319 = v314
	goto L113
L113:
	;
	v321 = v319
	goto L110
L114:
	;
	v326 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v326 != int64(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v343 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L116:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v329 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+4)))
	if v330&int32(2) == int32(0) {
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v336 = int32(1)
	v338 = F_tuplestore_gettupleslot(m, v335, v336, v336, v329)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	if v338 == int32(0) {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L115
L123:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	m.T0[v438].(func(*base.Module, int32))(m, v436)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L12
	} else {
		goto L143
	}
L124:
	;
	v347 = v343
	goto L125
L125:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+4)))
	if v359&int32(2) != 0 {
		goto L123
	} else {
		goto L127
	}
L126:
	;
	goto L123
L127:
	;
	v362 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
	if v321 < v362 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+32))
	m.T0[v366].(func(*base.Module, int32, int32))(m, v364, v347)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	v369 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v371 = v369 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v371
	F_spool_tuples(m, l0, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v376 = int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	v379 = F_tuplestore_gettupleslot(m, v375, v376, v376, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	if v379 == int32(0) {
		goto L123
	} else {
		goto L132
	}
L132:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+96))
	if v384 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	if v422 != 0 {
		v347 = v422
		goto L125
	} else {
		goto L142
	}
L134:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+412))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+8)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v388)+12)) = v387
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v392 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v388)+20))
	F_MemoryContextReset(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L12
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v398 = int32(_a_F_update_frametailpos_0)
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0]))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v388)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v401
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v392)+24))
	v406 = m.T0[v405].(func(*base.Module, int32, int32, int32) int64)(m, v392, v388, v16+int32(15))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L12
	} else {
		goto L139
	}
L138:
	;
	goto L133
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_update_frametailpos[0])) = v399
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v388)+20))
	F_MemoryContextReset(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	if v406 != int64(0) {
		goto L133
	} else {
		goto L141
	}
L141:
	;
	v415 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v415 + int64(1)
	goto L133
L142:
	;
	goto L126
L143:
	;
	goto L8
L144:
	;
	F_errmsg_internal(m, int32(_a_F_update_frametailpos_3), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_update_frametailpos_4), int32(1942), int32(_a_F_update_frametailpos_5))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errmsg_internal(m, int32(_a_F_update_frametailpos_3), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_update_frametailpos_4), int32(2035), int32(_a_F_update_frametailpos_5))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errmsg_internal(m, int32(_a_F_update_frametailpos_3), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_update_frametailpos_4), int32(2120), int32(_a_F_update_frametailpos_5))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
