package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_unicode_normalize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v828 int64
	_ = v828
	var v834 int64
	_ = v834
	var v838 int64
	_ = v838
	var v839 int64
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v867 int64
	_ = v867
	var v868 int64
	_ = v868
	var v870 int64
	_ = v870
	var v871 int64
	_ = v871
	var v873 int32
	_ = v873
	var v879 int64
	_ = v879
	var v880 int64
	_ = v880
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v973 int32
	_ = v973
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1051 int32
	_ = v1051
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v29 = l0 & int32(-2)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v30 == v3 {
		v207 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v223 = v207<<(uint(int32(2))%32) + int32(4)
	v224 = F_palloc(m, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v34 = base.B2i32(v29 == int32(2))
	v37 = l1
	v39 = v30
	v45 = v3
	goto L3
L3:
	;
	v64 = v39 - int32(_a_F_unicode_normalize_0)
	if base.Ui32(v64) <= base.Ui32(int32(_a_F_unicode_normalize_1)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v207 = v191
	goto L1
L5:
	;
	v191 = v190 + v45
	if base.Ui32(int32(268435456)) <= base.Ui32(v191) {
		v207 = v191
		goto L1
	} else {
		goto L26
	}
L6:
	;
	v72 = base.I32_rem_u_s(v64&int32(_a_F_unicode_normalize_2), int32(28))
	if v72 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v74 = int32(1)
	v75 = int32(16711935)
	v77 = int32(8)
	v78 = base.I32_rotr(v39&v75, v77)
	v79 = int32(24)
	v80 = base.I32_rotr(v39, v79)
	v82 = int32(255)
	v83 = (v78 | v80) & v82
	v84 = int32(127)
	v89 = int32(base.Ui32(v78)>>(uint(v77)%32)) & v82
	v96 = int32(base.Ui32(v80&v75) >> (uint(int32(16)) % 32))
	v101 = int32(base.Ui32(v78) >> (uint(v79) % 32))
	v105 = int32(_a_F_unicode_normalize_3)
	v106 = base.I32_rem_u_s(((v83*v84+v89)*v84+v96)*v84+v101+int32(260144641), v105)
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106<<(uint(v74)%32))+uint32(_c_F_unicode_normalize[0]))))
	v110 = int32(257)
	v120 = base.I32_rem_u_s(((v83*v110+v89)*v110+v96)*v110+v101, v105)
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v120<<(uint(v74)%32))+uint32(_c_F_unicode_normalize[0]))))
	v124 = v109 + v123
	if base.Ui32(int32(_a_F_unicode_normalize_4)) < base.Ui32(v124) {
		v179 = v74
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v73 = int32(3)
	goto L11
L10:
	;
	v73 = int32(2)
	goto L11
L11:
	;
	v190 = v73
	goto L5
L12:
	;
	v190 = v179
	goto L5
L13:
	;
	v128 = v124 << (uint(int32(3)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_c_F_unicode_normalize[1])))
	if v39 != v129 {
		v179 = v74
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+uint32(_c_F_unicode_normalize[2]))))
	v135 = v133 & int32(31)
	if v133&int32(32) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v141 = v34
	goto L17
L16:
	;
	v141 = int32(1)
	goto L17
L17:
	;
	if base.B2i32(v135 == int32(0))|base.B2i32(v141 == int32(0)) != 0 {
		v179 = v74
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+uint32(_c_F_unicode_normalize[3]))))
	if v133&int32(64) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v158 = int32(0)
	v160 = v158
	v163 = v158
	goto L23
L20:
	;
	v148 = int32(_a_F_unicode_normalize_5)
	*(*int32)(unsafe.Add(mBase, _c_F_unicode_normalize[4])) = v145
	v156 = int32(1)
	v157 = v148
	goto L19
L21:
	;
	goto L22
L22:
	;
	v156 = v135
	v157 = v145<<(uint(int32(2))%32) + int32(_a_F_unicode_normalize_6)
	goto L19
L23:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157+v160<<(uint(int32(2))%32))))
	v171 = F_get_decomposed_size(m, v170, v34)
	mBase = m.M
	v172 = v171 + v163
	v174 = v160 + int32(1)
	if v174 != v156 {
		v160 = v174
		v163 = v172
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v179 = v172
	goto L12
L25:
	;
	goto L24
L26:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v194 != 0 {
		v37 = v37 + int32(4)
		v39 = v194
		v45 = v191
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L4
L28:
	;
	return int32(0)
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v224
	if v224 == int32(0) {
		v1076 = v3
		goto L30
	} else {
		goto L31
	}
L30:
	;
	m.G0 = v26 + int32(16)
	return v1076
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v233 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v235 = base.B2i32(v29 == int32(2))
	v237 = l1
	v240 = v233
	goto L35
L33:
	;
	goto L34
L34:
	;
	v477 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v224+v207<<(uint(int32(2))%32)))) = v477
	if v207 == v477 {
		goto L58
	} else {
		goto L59
	}
L35:
	;
	v260 = v26 + int32(12)
	v262 = v26 + int32(8)
	v268 = v240 - int32(_a_F_unicode_normalize_0)
	if base.Ui32(v268) <= base.Ui32(int32(_a_F_unicode_normalize_1)) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	goto L34
L37:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v448 != 0 {
		v237 = v237 + int32(4)
		v240 = v448
		goto L35
	} else {
		goto L57
	}
L38:
	;
	goto L37
L39:
	;
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v367)+uint32(_c_F_unicode_normalize[3]))))
	if v372&int32(64) != 0 {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v398 + int32(1)
	goto L37
L41:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v273 = int32(2)
	v276 = int32(_a_F_unicode_normalize_2)
	v277 = v268 & v276
	v278 = int32(588)
	v279 = base.I32_div_u_s(v277, v278)
	*(*int32)(unsafe.Add(mBase, uint32(v271+v272<<(uint(v273)%32)))) = v279 | int32(_a_F_unicode_normalize_7)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v284 = int32(1)
	v285 = v283 + v284
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v285
	v295 = int32(28)
	v296 = base.I32_div_u_s((v268-v279*v278)&v276, v295)
	*(*int32)(unsafe.Add(mBase, uint32(v271+v285<<(uint(v273)%32)))) = v296 + int32(_a_F_unicode_normalize_8)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v302 = v300 + v284
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v302
	v305 = base.I32_rem_u_s(v277, v295)
	if v305 == int32(0) {
		goto L38
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v314 = int32(16711935)
	v316 = int32(8)
	v317 = base.I32_rotr(v240&v314, v316)
	v318 = int32(24)
	v319 = base.I32_rotr(v240, v318)
	v321 = int32(255)
	v322 = (v317 | v319) & v321
	v323 = int32(127)
	v328 = int32(base.Ui32(v317)>>(uint(v316)%32)) & v321
	v335 = int32(base.Ui32(v319&v314) >> (uint(int32(16)) % 32))
	v340 = int32(base.Ui32(v317) >> (uint(v318) % 32))
	v344 = int32(_a_F_unicode_normalize_3)
	v345 = base.I32_rem_u_s(((v322*v323+v328)*v323+v335)*v323+v340+int32(260144641), v344)
	v346 = int32(1)
	v348 = int32(*(*int16)(unsafe.Add(mBase, uint32(v345<<(uint(v346)%32))+uint32(_c_F_unicode_normalize[0]))))
	v349 = int32(257)
	v359 = base.I32_rem_u_s(((v322*v349+v328)*v349+v335)*v349+v340, v344)
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v359<<(uint(v346)%32))+uint32(_c_F_unicode_normalize[0]))))
	v363 = v348 + v362
	if base.Ui32(int32(_a_F_unicode_normalize_4)) < base.Ui32(v363) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271+v302<<(uint(int32(2))%32)))) = v305 + int32(_a_F_unicode_normalize_9)
	goto L40
L45:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	*(*int32)(unsafe.Add(mBase, uint32(v385+v386<<(uint(int32(2))%32)))) = v240
	goto L40
L46:
	;
	v367 = v363 << (uint(int32(3)) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+uint32(_c_F_unicode_normalize[1])))
	if v240 != v368 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+uint32(_c_F_unicode_normalize[2]))))
	v374 = v372 & int32(31)
	if v374 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	if v235|base.B2i32(v372&int32(32) == int32(0)) != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v416 = int32(0)
	goto L54
L51:
	;
	v405 = int32(_a_F_unicode_normalize_5)
	*(*int32)(unsafe.Add(mBase, _c_F_unicode_normalize[4])) = v402
	v413 = int32(1)
	v414 = v405
	goto L50
L52:
	;
	goto L53
L53:
	;
	v413 = v374
	v414 = v402<<(uint(int32(2))%32) + int32(_a_F_unicode_normalize_6)
	goto L50
L54:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v414+v416<<(uint(int32(2))%32))))
	F_decompose_code(m, v427, v235, v260, v262)
	mBase = m.M
	v430 = v416 + int32(1)
	if v430 != v413 {
		v416 = v430
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L38
L56:
	;
	goto L55
L57:
	;
	goto L36
L58:
	;
	v1076 = v224
	goto L30
L59:
	;
	goto L60
L60:
	;
	if int32(2) <= v207 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v486 = int32(1)
	goto L64
L62:
	;
	goto L63
L63:
	;
	if l0&int32(-3) != 0 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v509 = v224 + v486<<(uint(int32(2))%32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v511 = int32(0)
	v513 = v509 - int32(4)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v515 = int32(16711935)
	v517 = int32(8)
	v518 = base.I32_rotr(v514&v515, v517)
	v519 = int32(24)
	v520 = base.I32_rotr(v514, v519)
	v522 = int32(255)
	v523 = (v518 | v520) & v522
	v524 = int32(127)
	v529 = int32(base.Ui32(v518)>>(uint(v517)%32)) & v522
	v536 = int32(base.Ui32(v520&v515) >> (uint(int32(16)) % 32))
	v541 = int32(base.Ui32(v518) >> (uint(v519) % 32))
	v545 = int32(_a_F_unicode_normalize_3)
	v546 = base.I32_rem_u_s(((v523*v524+v529)*v524+v536)*v524+v541+int32(260144641), v545)
	v547 = int32(1)
	v549 = int32(*(*int16)(unsafe.Add(mBase, uint32(v546<<(uint(v547)%32))+uint32(_c_F_unicode_normalize[0]))))
	v550 = int32(257)
	v560 = base.I32_rem_u_s(((v523*v550+v529)*v550+v536)*v550+v541, v545)
	v563 = int32(*(*int16)(unsafe.Add(mBase, uint32(v560<<(uint(v547)%32))+uint32(_c_F_unicode_normalize[0]))))
	v564 = v549 + v563
	if base.Ui32(int32(_a_F_unicode_normalize_4)) < base.Ui32(v564) {
		v575 = v511
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L63
L66:
	;
	v576 = int32(16711935)
	v578 = int32(8)
	v579 = base.I32_rotr(v510&v576, v578)
	v580 = int32(24)
	v581 = base.I32_rotr(v510, v580)
	v583 = int32(255)
	v584 = (v579 | v581) & v583
	v585 = int32(127)
	v590 = int32(base.Ui32(v579)>>(uint(v578)%32)) & v583
	v597 = int32(base.Ui32(v581&v576) >> (uint(int32(16)) % 32))
	v602 = int32(base.Ui32(v579) >> (uint(v580) % 32))
	v606 = int32(_a_F_unicode_normalize_3)
	v607 = base.I32_rem_u_s(((v584*v585+v590)*v585+v597)*v585+v602+int32(260144641), v606)
	v608 = int32(1)
	v610 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607<<(uint(v608)%32))+uint32(_c_F_unicode_normalize[0]))))
	v611 = int32(257)
	v621 = base.I32_rem_u_s(((v584*v611+v590)*v611+v597)*v611+v602, v606)
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(v621<<(uint(v608)%32))+uint32(_c_F_unicode_normalize[0]))))
	v625 = v610 + v624
	if base.Ui32(int32(_a_F_unicode_normalize_4)) < base.Ui32(v625) {
		v649 = v486
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v568 = v564 << (uint(int32(3)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+uint32(_c_F_unicode_normalize[1])))
	if v514 != v569 {
		v575 = v511
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+uint32(_c_F_unicode_normalize[5]))))
	v575 = v573
	goto L66
L69:
	;
	v652 = v649 + int32(1)
	if v652 < v207 {
		v486 = v652
		goto L64
	} else {
		goto L76
	}
L70:
	;
	v631 = v625 << (uint(int32(3)) % 32)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+uint32(_c_F_unicode_normalize[1])))
	if base.B2i32(v575 == int32(0))|base.B2i32(v510 != v632) != 0 {
		v649 = v486
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631)+uint32(_c_F_unicode_normalize[5]))))
	if base.B2i32(v637 == int32(0))|base.B2i32(base.Ui32(v575) <= base.Ui32(v637)) != 0 {
		v649 = v486
		goto L69
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = v514
	if int32(1) < v486 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v648 = v486 - int32(2)
	goto L75
L74:
	;
	v648 = v486
	goto L75
L75:
	;
	v649 = v648
	goto L69
L76:
	;
	goto L65
L77:
	;
	v1076 = v224
	goto L30
L78:
	;
	goto L79
L79:
	;
	v679 = F_palloc(m, v223)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L28
	} else {
		goto L81
	}
L80:
	;
	F_pfree(m, v224)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L28
	} else {
		goto L112
	}
L81:
	;
	if v679 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v1051 = int32(0)
	goto L80
L83:
	;
	goto L84
L84:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v679))) = v684
	v686 = int32(1)
	if int32(2) <= v207 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v692 = v686
	v694 = v684
	v695 = int32(1)
	v701 = int32(0)
	v706 = int32(-1)
	goto L88
L86:
	;
	v1016 = v686
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679+v1016<<(uint(int32(2))%32)))) = int32(0)
	v1051 = v679
	goto L80
L88:
	;
	v715 = int32(0)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v224+v695<<(uint(int32(2))%32))))
	v720 = int32(16711935)
	v722 = int32(8)
	v723 = base.I32_rotr(v719&v720, v722)
	v724 = int32(24)
	v725 = base.I32_rotr(v719, v724)
	v727 = int32(255)
	v728 = (v723 | v725) & v727
	v729 = int32(127)
	v734 = int32(base.Ui32(v723)>>(uint(v722)%32)) & v727
	v741 = int32(base.Ui32(v725&v720) >> (uint(int32(16)) % 32))
	v746 = int32(base.Ui32(v723) >> (uint(v724) % 32))
	v750 = int32(_a_F_unicode_normalize_3)
	v751 = base.I32_rem_u_s(((v728*v729+v734)*v729+v741)*v729+v746+int32(260144641), v750)
	v752 = int32(1)
	v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v751<<(uint(v752)%32))+uint32(_c_F_unicode_normalize[0]))))
	v755 = int32(257)
	v765 = base.I32_rem_u_s(((v728*v755+v734)*v755+v741)*v755+v746, v750)
	v768 = int32(*(*int16)(unsafe.Add(mBase, uint32(v765<<(uint(v752)%32))+uint32(_c_F_unicode_normalize[0]))))
	v769 = v754 + v768
	if base.Ui32(int32(_a_F_unicode_normalize_4)) < base.Ui32(v769) {
		v779 = v715
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v1016 = v997
	goto L87
L90:
	;
	if v779 <= v706 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v773 = v769 << (uint(int32(3)) % 32)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)+uint32(_c_F_unicode_normalize[1])))
	if v719 != v774 {
		v779 = v715
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+uint32(_c_F_unicode_normalize[5]))))
	v779 = v778
	goto L90
L93:
	;
	v1014 = v695 + int32(1)
	if v1014 != v207 {
		v692 = v997
		v694 = v998
		v695 = v1014
		v701 = v1001
		v706 = v1004
		goto L88
	} else {
		goto L111
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679+v692<<(uint(int32(2))%32)))) = v719
	if v779 != 0 {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v694-int32(_a_F_unicode_normalize_7)))|base.B2i32(base.Ui32(int32(20)) < base.Ui32(v719-int32(_a_F_unicode_normalize_8))) == int32(0) {
		v973 = (v719+v694*int32(21))*int32(28) - int32(_a_F_unicode_normalize_10)
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679+v701<<(uint(int32(2))%32)))) = v973
	v997 = v692
	v998 = v973
	v1001 = v701
	v1004 = v706
	goto L93
L97:
	;
	v807 = v694 - int32(_a_F_unicode_normalize_0)
	v813 = base.I32_rem_u_s(v807&int32(_a_F_unicode_normalize_2), int32(28))
	if base.B2i32(base.Ui32(int32(_a_F_unicode_normalize_1)) < base.Ui32(v807))|v813|base.B2i32(base.Ui32(int32(26)) < base.Ui32(v719-int32(_a_F_unicode_normalize_11))) == int32(0) {
		v973 = v694 + v719 - int32(_a_F_unicode_normalize_9)
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v822 = int32(24)
	v825 = int32(8)
	v828 = int64(16711680)
	v834 = int64(65280)
	v838 = base.I64_extend_i32_u(v694) << (uint(int64(32)) % 64)
	v839 = int64(56)
	v843 = base.I32_wrap_i64(base.I64_extend_i32_u(v694<<(uint(v822)%32)) | base.I64_extend_i32_u(v694<<(uint(v825)%32))&v828 | (base.I64_extend_i32_u(int32(base.Ui32(v694)>>(uint(v825)%32)))&v834 | int64(base.Ui64(v838)>>(uint(v839)%64))))
	v844 = int32(255)
	v845 = v843 & v844
	v846 = int32(17)
	v851 = int32(base.Ui32(v843)>>(uint(v825)%32)) & v844
	v858 = int32(base.Ui32(v843)>>(uint(int32(16))%32)) & v844
	v863 = int32(base.Ui32(v843) >> (uint(v822) % 32))
	v867 = base.I64_extend_i32_u(v719)
	v868 = v838 | v867
	v870 = v868 & int64(4278190080)
	v871 = int64(24)
	v873 = base.I32_wrap_i64(int64(base.Ui64(v870) >> (uint(v871) % 64)))
	v879 = int64(40)
	v880 = v867 & v834 << (uint(v879) % 64)
	v896 = base.I32_wrap_i64(int64(base.Ui64(v880|v867<<(uint(v839)%64)|(v868&v828<<(uint(v871)%64)|v870<<(uint(int64(8))%64)))>>(uint(v879)%64))) & v844
	v902 = base.I32_wrap_i64(int64(base.Ui64(v880) >> (uint(int64(48)) % 64)))
	v908 = base.I32_wrap_i64(v867 & int64(255))
	v910 = int32(1923)
	v911 = base.I32_rem_u_s(((((((v845*v846+v851)*v846+v858)*v846+v863)*v846+v873)*v846+v896)*v846+v902)*v846+v908, v910)
	v912 = int32(1)
	v914 = int32(*(*int16)(unsafe.Add(mBase, uint32(v911<<(uint(v912)%32))+uint32(_c_F_unicode_normalize[6]))))
	v915 = int32(257)
	v937 = base.I32_rem_u_s(((((((v845*v915+v851)*v915+v858)*v915+v863)*v915+v873)*v915+v896)*v915+v902)*v915+v908, v910)
	v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v937<<(uint(v912)%32))+uint32(_c_F_unicode_normalize[6]))))
	v941 = v914 + v940
	if base.Ui32(int32(960)) < base.Ui32(v941) {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_c_F_unicode_normalize[7]))))
	v948 = v946 << (uint(int32(3)) % 32)
	v949 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+uint32(_c_F_unicode_normalize[3]))))
	v951 = v949 << (uint(int32(2)) % 32)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+uint32(_c_F_unicode_normalize[8])))
	if v694 != v952 {
		goto L94
	} else {
		goto L100
	}
L100:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v951)+uint32(_c_F_unicode_normalize[9])))
	if v719 != v956 {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v948)+uint32(_c_F_unicode_normalize[1])))
	v973 = v960
	goto L96
L102:
	;
	v992 = v779
	goto L104
L103:
	;
	v992 = int32(-1)
	goto L104
L104:
	;
	if v779 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v993 = v701
	goto L107
L106:
	;
	v993 = v692
	goto L107
L107:
	;
	if v779 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v994 = v694
	goto L110
L109:
	;
	v994 = v719
	goto L110
L110:
	;
	v997 = v692 + int32(1)
	v998 = v994
	v1001 = v993
	v1004 = v992
	goto L93
L111:
	;
	goto L89
L112:
	;
	v1076 = v1051
	goto L30
}
