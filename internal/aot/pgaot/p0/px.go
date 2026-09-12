package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_crypt_des(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v391 int32
	_ = v391
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_crypt_des[0])))
	if v24 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_des_init(m)
	mBase = m.M
	goto L3
L2:
	;
	goto L3
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v29 = int32(1)
	v30 = v28 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v30)
	v32 = int32(0)
	v34 = l0 + base.B2i32(v28 != v32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v37 = v35 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v37)
	v41 = v34 + base.B2i32(v35 != v32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v44 = v42 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)) = uint8(v44)
	v48 = v41 + base.B2i32(v42 != v32)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v51 = v49 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)) = uint8(v51)
	v55 = v48 + base.B2i32(v49 != v32)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v58 = v56 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v58)
	v62 = v55 + base.B2i32(v56 != v32)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v65 = v63 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)) = uint8(v65)
	v69 = v62 + base.B2i32(v63 != v32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v72 = v70 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v72)
	v76 = v69 + base.B2i32(v70 != v32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v79 = v77 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v79)
	F_des_setkey(m, v21)
	mBase = m.M
	if l1&int32(3) == v32 {
		v105 = l1
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v139 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v139 == int32(95) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v138 = v130 - l1
	goto L4
L6:
	;
	v109 = v105
	goto L15
L7:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v89 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v138 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v94 = l1
	goto L11
L11:
	;
	v98 = v94 + int32(1)
	if v98&int32(3) == int32(0) {
		v105 = v98
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v130 = v98
	goto L5
L13:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v103 != 0 {
		v94 = v98
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v118 = int32(-2139062144)
	if (int32(16843008)-v115|v115)&v118 == v118 {
		v109 = v109 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v124 = v109
	goto L18
L17:
	;
	goto L16
L18:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v128 != 0 {
		v124 = v124 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v130 = v124
	goto L5
L20:
	;
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L128
	} else {
		goto L256
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L128
	} else {
		goto L252
	}
L23:
	;
	m.G0 = v21 + int32(16)
	return v1006
L24:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[1]))
	if v831 != v816 {
		goto L230
	} else {
		goto L231
	}
L25:
	;
	if base.Ui32(v138) < base.Ui32(int32(9)) {
		goto L22
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(v138) <= base.Ui32(int32(1)) {
		goto L21
	} else {
		goto L206
	}
L28:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if int32(122) < v145 {
		v168 = int32(0)
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
	if int32(122) < v169 {
		v191 = v3
		goto L37
	} else {
		goto L38
	}
L30:
	;
	if int32(97) <= v145 {
		v168 = v145 - int32(59)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if int32(90) < v145 {
		v168 = int32(0)
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if int32(65) <= v145 {
		v168 = v145 - int32(53)
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v160 = v145 - int32(46)
	if base.Ui32(v160&int32(255)) < base.Ui32(int32(12)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v166 = v160
	goto L36
L35:
	;
	v166 = int32(0)
	goto L36
L36:
	;
	v168 = v166
	goto L29
L37:
	;
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+3)))
	if int32(122) < v193 {
		v216 = int32(0)
		goto L49
	} else {
		goto L50
	}
L38:
	;
	if v169 <= int32(96) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if int32(90) < v169 {
		v191 = v3
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v191 = v169 - int32(59)
	goto L37
L42:
	;
	if v169 <= int32(64) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v179 = v169 - int32(46)
	if base.Ui32(v179&int32(255)) < base.Ui32(int32(12)) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v191 = v169 - int32(53)
	goto L37
L46:
	;
	v185 = v179
	goto L48
L47:
	;
	v185 = int32(0)
	goto L48
L48:
	;
	v191 = v185
	goto L37
L49:
	;
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+4)))
	if int32(122) < v217 {
		v239 = v3
		goto L61
	} else {
		goto L62
	}
L50:
	;
	if v193 <= int32(96) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(90) < v193 {
		v216 = int32(0)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v216 = v193 - int32(59)
	goto L49
L54:
	;
	if v193 <= int32(64) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v204 = v193 - int32(46)
	if base.Ui32(v204&int32(255)) < base.Ui32(int32(12)) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v216 = v193 - int32(53)
	goto L49
L58:
	;
	v210 = v204
	goto L60
L59:
	;
	v210 = int32(0)
	goto L60
L60:
	;
	v216 = v210
	goto L49
L61:
	;
	v241 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+5)))
	if int32(122) < v241 {
		v264 = int32(0)
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if v217 <= int32(96) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if int32(90) < v217 {
		v239 = v3
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v239 = v217 - int32(59)
	goto L61
L66:
	;
	if v217 <= int32(64) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v227 = v217 - int32(46)
	if base.Ui32(v227&int32(255)) < base.Ui32(int32(12)) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v239 = v217 - int32(53)
	goto L61
L70:
	;
	v233 = v227
	goto L72
L71:
	;
	v233 = int32(0)
	goto L72
L72:
	;
	v239 = v233
	goto L61
L73:
	;
	v265 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+6)))
	if int32(122) < v265 {
		v287 = v3
		goto L81
	} else {
		goto L82
	}
L74:
	;
	if int32(97) <= v241 {
		v264 = v241 - int32(59)
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if int32(90) < v241 {
		v264 = int32(0)
		goto L73
	} else {
		goto L76
	}
L76:
	;
	if int32(65) <= v241 {
		v264 = v241 - int32(53)
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v256 = v241 - int32(46)
	if base.Ui32(v256&int32(255)) < base.Ui32(int32(12)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v262 = v256
	goto L80
L79:
	;
	v262 = int32(0)
	goto L80
L80:
	;
	v264 = v262
	goto L73
L81:
	;
	v288 = int32(0)
	v291 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+7)))
	if int32(122) < v291 {
		v314 = v288
		goto L93
	} else {
		goto L94
	}
L82:
	;
	if v265 <= int32(96) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if int32(90) < v265 {
		v287 = v3
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v287 = v265 - int32(59)
	goto L81
L86:
	;
	if v265 <= int32(64) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v275 = v265 - int32(46)
	if base.Ui32(v275&int32(255)) < base.Ui32(int32(12)) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v287 = v265 - int32(53)
	goto L81
L90:
	;
	v281 = v275
	goto L92
L91:
	;
	v281 = int32(0)
	goto L92
L92:
	;
	v287 = v281
	goto L81
L93:
	;
	v315 = base.B2i32(v77 != v288) + v76
	v316 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+8)))
	if int32(122) < v316 {
		v338 = v3
		goto L105
	} else {
		goto L106
	}
L94:
	;
	if v291 <= int32(96) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if int32(90) < v291 {
		v314 = int32(0)
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v314 = v291 - int32(59)
	goto L93
L98:
	;
	if v291 <= int32(64) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v302 = v291 - int32(46)
	if base.Ui32(v302&int32(255)) < base.Ui32(int32(12)) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v314 = v291 - int32(53)
	goto L93
L102:
	;
	v308 = v302
	goto L104
L103:
	;
	v308 = int32(0)
	goto L104
L104:
	;
	v314 = v308
	goto L93
L105:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v339 != 0 {
		goto L117
	} else {
		goto L118
	}
L106:
	;
	if v316 <= int32(96) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if int32(90) < v316 {
		v338 = v3
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v338 = v316 - int32(59)
	goto L105
L110:
	;
	if v316 <= int32(64) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v326 = v316 - int32(46)
	if base.Ui32(v326&int32(255)) < base.Ui32(int32(12)) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v338 = v316 - int32(53)
	goto L105
L114:
	;
	v332 = v326
	goto L116
L115:
	;
	v332 = int32(0)
	goto L116
L116:
	;
	v338 = v332
	goto L105
L117:
	;
	v340 = v315
	goto L120
L118:
	;
	goto L119
L119:
	;
	v559 = int32(6)
	v562 = int32(12)
	v565 = int32(18)
	v577 = int32(_a_F_px_crypt_des_0)
	goto L160
L120:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_crypt_des[0])))
	if v359 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L119
L122:
	;
	F_des_init(m)
	mBase = m.M
	goto L124
L123:
	;
	goto L124
L124:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[1]))
	if v364 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v365 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[2])) = v365
	*(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[1])) = v365
	goto L127
L126:
	;
	goto L127
L127:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v376 = int32(24)
	v378 = int32(_a_F_px_crypt_des_1)
	v380 = int32(8)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v412 = F_do_des(m, v375<<(uint(v376)%32)|v375&v378<<(uint(v380)%32)|(int32(base.Ui32(v375)>>(uint(v380)%32))&v378|int32(base.Ui32(v375)>>(uint(v376)%32))), v391<<(uint(v376)%32)|v391&v378<<(uint(v380)%32)|(int32(base.Ui32(v391)>>(uint(v380)%32))&v378|int32(base.Ui32(v391)>>(uint(v376)%32))), v21+int32(12), v21+v380, int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	return int32(0)
L129:
	;
	if v412 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v1006 = int32(0)
	goto L23
L131:
	;
	goto L132
L132:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v418 = int32(24)
	v420 = int32(_a_F_px_crypt_des_1)
	v422 = int32(8)
	v424 = v417<<(uint(v418)%32) | v417&v420<<(uint(v422)%32)
	v432 = v424 | (int32(base.Ui32(v417)>>(uint(v422)%32))&v420 | int32(base.Ui32(v417)>>(uint(v418)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v441 = v434<<(uint(v418)%32) | v434&v420<<(uint(v422)%32)
	v449 = v441 | (int32(base.Ui32(v434)>>(uint(v422)%32))&v420 | int32(base.Ui32(v434)>>(uint(v418)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v449
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v451 == int32(0) {
		v535 = v340
		goto L133
	} else {
		goto L134
	}
L133:
	;
	F_des_setkey(m, v21)
	mBase = m.M
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v540 != 0 {
		v340 = v535
		goto L120
	} else {
		goto L156
	}
L134:
	;
	v456 = v451<<(uint(int32(1))%32) ^ v449
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v456)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if v458 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v535 = v340 + int32(1)
	goto L133
L136:
	;
	goto L137
L137:
	;
	v467 = v458<<(uint(int32(1))%32) ^ int32(base.Ui32(v449)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v467)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+2)))
	if v469 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v535 = v340 + int32(2)
	goto L133
L139:
	;
	goto L140
L140:
	;
	v478 = v469<<(uint(int32(1))%32) ^ int32(base.Ui32(v441)>>(uint(int32(16))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)) = uint8(v478)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+3)))
	if v480 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v535 = v340 + int32(3)
	goto L133
L142:
	;
	goto L143
L143:
	;
	v489 = v434&int32(255) ^ v480<<(uint(int32(1))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)) = uint8(v489)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+4)))
	if v491 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v535 = v340 + int32(4)
	goto L133
L145:
	;
	goto L146
L146:
	;
	v498 = v491<<(uint(int32(1))%32) ^ v432
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v498)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+5)))
	if v500 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v535 = v340 + int32(5)
	goto L133
L148:
	;
	goto L149
L149:
	;
	v509 = v500<<(uint(int32(1))%32) ^ int32(base.Ui32(v432)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)) = uint8(v509)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+6)))
	if v511 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v535 = v340 + int32(6)
	goto L133
L151:
	;
	goto L152
L152:
	;
	v520 = v511<<(uint(int32(1))%32) ^ int32(base.Ui32(v424)>>(uint(int32(16))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v520)
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+7)))
	if v522 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v535 = v340 + int32(7)
	goto L133
L154:
	;
	goto L155
L155:
	;
	v531 = v417&int32(255) ^ v522<<(uint(int32(1))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v531)
	v535 = v340 + int32(8)
	goto L133
L156:
	;
	goto L121
L157:
	;
	goto L191
L158:
	;
	v690 = F_strlen(m, v679)
	mBase = m.M
	goto L157
L160:
	;
	goto L161
L161:
	;
	v584 = int32(9)
	if (v577^l1)&int32(3) != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v680))) = uint8(v683)
	goto L158
L163:
	;
	v664 = v659
	v665 = v660
	v666 = v661
	goto L185
L164:
	;
	if v654 == int32(0) {
		v679 = v652
		v680 = v653
		goto L162
	} else {
		goto L184
	}
L165:
	;
	v652 = l1
	v653 = v577
	v654 = v584
	goto L164
L166:
	;
	goto L167
L167:
	;
	if l1&int32(3) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v621 == int32(0) {
		v679 = v618
		v680 = v619
		goto L162
	} else {
		goto L177
	}
L169:
	;
	v618 = l1
	v619 = v577
	v620 = v584
	v621 = int32(1)
	goto L168
L170:
	;
	goto L171
L171:
	;
	v597 = l1
	v598 = v577
	v599 = v584
	goto L172
L172:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	*(*uint8)(unsafe.Add(mBase, uint32(v598))) = uint8(v601)
	if v601 == int32(0) {
		v659 = v597
		v660 = v598
		v661 = v599
		goto L163
	} else {
		goto L174
	}
L173:
	;
	v618 = v612
	v619 = v606
	v620 = v608
	v621 = v610
	goto L168
L174:
	;
	v605 = int32(1)
	v606 = v598 + v605
	v608 = v599 - v605
	v609 = int32(0)
	v610 = base.B2i32(v608 != v609)
	v612 = v597 + v605
	if v612&int32(3) == v609 {
		v618 = v612
		v619 = v606
		v620 = v608
		v621 = v610
		goto L168
	} else {
		goto L175
	}
L175:
	;
	if v608 != 0 {
		v597 = v612
		v598 = v606
		v599 = v608
		goto L172
	} else {
		goto L176
	}
L176:
	;
	goto L173
L177:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	if v624 == int32(0) {
		v652 = v618
		v653 = v619
		v654 = v620
		goto L164
	} else {
		goto L178
	}
L178:
	;
	if base.Ui32(v620) < base.Ui32(int32(4)) {
		v652 = v618
		v653 = v619
		v654 = v620
		goto L164
	} else {
		goto L179
	}
L179:
	;
	v630 = v618
	v631 = v619
	v632 = v620
	goto L180
L180:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v638 = int32(-2139062144)
	if (int32(16843008)-v635|v635)&v638 != v638 {
		v659 = v630
		v660 = v631
		v661 = v632
		goto L163
	} else {
		goto L182
	}
L181:
	;
	v652 = v646
	v653 = v644
	v654 = v648
	goto L164
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631))) = v635
	v643 = int32(4)
	v644 = v631 + v643
	v646 = v630 + v643
	v648 = v632 - v643
	if base.Ui32(int32(3)) < base.Ui32(v648) {
		v630 = v646
		v631 = v644
		v632 = v648
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v659 = v652
	v660 = v653
	v661 = v654
	goto L163
L185:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	*(*uint8)(unsafe.Add(mBase, uint32(v665))) = uint8(v668)
	if v668 == int32(0) {
		v679 = v664
		v680 = v665
		goto L162
	} else {
		goto L187
	}
L186:
	;
	v679 = v675
	v680 = v673
	goto L162
L187:
	;
	v672 = int32(1)
	v673 = v665 + v672
	v675 = v664 + v672
	v677 = v666 - v672
	if v677 != 0 {
		v664 = v675
		v665 = v673
		v666 = v677
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v816 = v287<<(uint(v559)%32) | v264 | v314<<(uint(v562)%32) | v338<<(uint(v565)%32)
	v820 = v191<<(uint(v559)%32) | v168 | v216<<(uint(v562)%32) | v239<<(uint(v565)%32)
	v829 = v735 - v577 + v577
	goto L24
L190:
	;
	goto L189
L191:
	;
	v720 = v577
	goto L200
L200:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v729 = int32(-2139062144)
	if (int32(16843008)-v726|v726)&v729 == v729 {
		v720 = v720 + int32(4)
		goto L200
	} else {
		goto L202
	}
L201:
	;
	v735 = v720
	goto L203
L202:
	;
	goto L201
L203:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	if v739 != 0 {
		v735 = v735 + int32(1)
		goto L203
	} else {
		goto L205
	}
L204:
	;
	goto L190
L205:
	;
	goto L204
L206:
	;
	v753 = int32(0)
	v755 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if int32(122) < v755 {
		v778 = v753
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if int32(122) < v139 {
		v802 = v753
		goto L215
	} else {
		goto L216
	}
L208:
	;
	if int32(97) <= v755 {
		v778 = v755 - int32(59)
		goto L207
	} else {
		goto L209
	}
L209:
	;
	if int32(90) < v755 {
		v778 = int32(0)
		goto L207
	} else {
		goto L210
	}
L210:
	;
	if int32(65) <= v755 {
		v778 = v755 - int32(53)
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v770 = v755 - int32(46)
	if base.Ui32(v770&int32(255)) < base.Ui32(int32(12)) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v776 = v770
	goto L214
L213:
	;
	v776 = int32(0)
	goto L214
L214:
	;
	v778 = v776
	goto L207
L215:
	;
	v803 = int32(_a_F_px_crypt_des_0)
	if v755 != 0 {
		goto L227
	} else {
		goto L228
	}
L216:
	;
	if int32(97) <= v139 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v802 = v139 - int32(59)
	goto L215
L218:
	;
	goto L219
L219:
	;
	if int32(90) < v139 {
		v802 = v753
		goto L215
	} else {
		goto L220
	}
L220:
	;
	if int32(65) <= v139 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v802 = v139 - int32(53)
	goto L215
L222:
	;
	goto L223
L223:
	;
	v794 = v139 - int32(46)
	if base.Ui32(v794&int32(255)) < base.Ui32(int32(12)) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v800 = v794
	goto L226
L225:
	;
	v800 = int32(0)
	goto L226
L226:
	;
	v802 = v800
	goto L215
L227:
	;
	v804 = v755
	goto L229
L228:
	;
	v804 = v139
	goto L229
L229:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_px_crypt_des[3])) = uint8(v804)
	*(*uint8)(unsafe.Add(mBase, _c_F_px_crypt_des[4])) = uint8(v139)
	v816 = v802 + v778<<(uint(int32(6))%32)
	v820 = int32(25)
	v829 = int32(_a_F_px_crypt_des_2)
	goto L24
L230:
	;
	v833 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[2])) = v833
	*(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[1])) = v816
	v845 = int32(_a_F_px_crypt_des_3)
	v847 = int32(1)
	v849 = v833
	v852 = v833
	goto L233
L231:
	;
	goto L232
L232:
	;
	v910 = int32(0)
	v917 = F_do_des(m, v910, v910, v21+int32(12), v21+int32(8), v820)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L128
	} else {
		goto L250
	}
L233:
	;
	v864 = v847 & v816
	if v864 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	goto L232
L235:
	;
	v865 = v845 | v849
	goto L237
L236:
	;
	v865 = v849
	goto L237
L237:
	;
	v866 = int32(1)
	v871 = v847 << (uint(v866) % 32) & v816
	if v871 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v872 = v865 | int32(base.Ui32(v845)>>(uint(v866)%32))
	goto L240
L239:
	;
	v872 = v865
	goto L240
L240:
	;
	v873 = int32(2)
	v878 = v847 << (uint(v873) % 32) & v816
	if v878 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v879 = v872 | int32(base.Ui32(v845)>>(uint(v873)%32))
	goto L243
L242:
	;
	v879 = v872
	goto L243
L243:
	;
	if v864 != 0 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v884 = int32(3)
	v889 = v852 + v884
	if v889 != int32(24) {
		v845 = int32(base.Ui32(v845) >> (uint(v884) % 32))
		v847 = v847 << (uint(v884) % 32)
		v849 = v879
		v852 = v889
		goto L233
	} else {
		goto L249
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_px_crypt_des[2])) = v879
	goto L244
L246:
	;
	if v871 != 0 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	if v878 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	goto L245
L249:
	;
	goto L234
L250:
	;
	if v917 != 0 {
		v1006 = v910
		goto L23
	} else {
		goto L251
	}
L251:
	;
	v919 = int32(0)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v922)>>(uint(int32(26))%32)))+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829))) = uint8(v926)
	v930 = int32(63)
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v922)>>(uint(int32(8))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+3)) = uint8(v933)
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v922)>>(uint(int32(14))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+2)) = uint8(v940)
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v922)>>(uint(int32(20))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+1)) = uint8(v947)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+11)) = uint8(v919)
	v952 = int32(2)
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v922)>>(uint(v952)%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+4)) = uint8(v957)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949<<(uint(v952)%32)&int32(60))+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+10)) = uint8(v964)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v949)>>(uint(int32(4))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+9)) = uint8(v971)
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v949)>>(uint(int32(10))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+8)) = uint8(v978)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v949)>>(uint(int32(22))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+6)) = uint8(v985)
	v987 = int32(16)
	v991 = v922<<(uint(v987)%32) | int32(base.Ui32(v949)>>(uint(v987)%32))
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+7)) = uint8(v995)
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v991)>>(uint(int32(12))%32))&v930)+uint32(_c_F_px_crypt_des[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v829)+5)) = uint8(v1002)
	v1006 = int32(_a_F_px_crypt_des_0)
	goto L23
L252:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L128
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_px_crypt_des_4), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L128
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_px_crypt_des_5), int32(697), int32(_a_F_px_crypt_des_6))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L128
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L128
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_px_crypt_des_4), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L128
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_px_crypt_des_5), int32(745), int32(_a_F_px_crypt_des_6))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L128
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_px_debug(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(528)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+524)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_px_debug[0]))
	if v10 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+524))
		v13 = F_pg_vsnprintf(m, v6, int32(512), l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_px_debug[0]))
			m.T0[v16].(func(*base.Module, int32))(m, v6)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(528)
				return
			}
		}
	} else {
		m.G0 = v6 + int32(528)
		return
	}
}
func F_px_find_hmac(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v62 int32
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_px_find_digest(m, l0, v8+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v62 = v12
			m.G0 = v8 + int32(16)
			return v62
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if base.Ui32(v18) <= base.Ui32(int32(1)) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
					m.T0[v23].(func(*base.Module, int32))(m, v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v62 = int32(-9)
						m.G0 = v8 + int32(16)
						return v62
					}
				} else {
					v28 = F_palloc(m, int32(40))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = F_palloc(m, v18)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v30
							v33 = F_palloc(m, v18)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v33
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = int32(_a_F_px_find_hmac_0)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = int32(_a_F_px_find_hmac_1)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(_a_F_px_find_hmac_2)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(_a_F_px_find_hmac_3)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(_a_F_px_find_hmac_4)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(_a_F_px_find_hmac_5)
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(_a_F_px_find_hmac_6)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v36
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
								v62 = int32(0)
								m.G0 = v8 + int32(16)
								return v62
							}
						}
					}
				}
			}
		}
	}
}
func F_px_set_debug_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_px_set_debug_handler[0])) = l0
	return
}
func F_px_strerror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(_a_F_px_strerror_0)
L2:
	;
	goto L3
L3:
	;
	v13 = int32(_a_F_px_strerror_1)
	goto L5
L4:
	;
	return v32
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if l0 != v16 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v32 = v29
	goto L4
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v19 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	return int32(_a_F_px_strerror_2)
L11:
	;
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if l0 != v25 {
		v13 = v13 + int32(16)
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v32 = v19
	goto L4
}
