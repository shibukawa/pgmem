package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_is_kind_built(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l1 - int32(100)
	v13 = v11 & int32(255)
	if base.B2i32(base.Ui32(int32(10)) <= base.Ui32(v13))|base.B2i32(int32(base.Ui32(int32(519))>>(uint(v13)%32))&int32(1) == v3) == v3 {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v11&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_statext_is_kind_built[0])))
		v31 = F_heap_attisnull(m, l0, v29, int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v31 ^ int32(1)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_statext_is_kind_built_0), v8)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_statext_is_kind_built_1), int32(444), int32(_a_F_statext_is_kind_built_2))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
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
func F_statext_mcv_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v402 int32
	_ = v402
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int64
	_ = v659
	var v661 int64
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int64
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int64
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int64
	_ = v775
	var v780 int32
	_ = v780
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v854 int32
	_ = v854
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1065 int32
	_ = v1065
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	if l0 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L23
	} else {
		goto L220
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L23
	} else {
		goto L204
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L23
	} else {
		goto L201
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L23
	} else {
		goto L198
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L23
	} else {
		goto L195
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L23
	} else {
		goto L192
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L23
	} else {
		goto L189
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L23
	} else {
		goto L186
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L23
	} else {
		goto L170
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v854 = v2
	goto L12
L12:
	;
	m.G0 = v23 + int32(128)
	return v854
L13:
	;
	if base.Ui32(v48) <= base.Ui32(int32(17)) {
		goto L9
	} else {
		goto L22
	}
L14:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = int32(1)
	if v25&v40 != 0 {
		v48 = int32(base.Ui32(v25) >> (uint(v40) % 32))
		goto L13
	} else {
		goto L21
	}
L17:
	;
	v35 = int32(18)
	if v28 == v35 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v39 = v35
	goto L20
L19:
	;
	v39 = int32(2)
	goto L20
L20:
	;
	v48 = v39
	goto L13
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = int32(base.Ui32(v44) >> (uint(int32(2)) % 32))
	goto L13
L22:
	;
	v52 = F_palloc0(m, int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v56 = int32(1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v58&v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = v56
	goto L27
L26:
	;
	v61 = int32(4)
	goto L27
L27:
	;
	v62 = l0 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v67
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+12)) = uint16(v69)
	if v63 != int32(-509193790) {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	if v65 != int32(1) {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	if v69 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v69) {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	if v67 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(int32(_a_F_statext_mcv_deserialize_0)) <= base.Ui32(v67) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v83 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v110 = v69 << (uint(int32(2)) % 32)
	v112 = v69 * int32(20)
	v121 = v110 + v112 + (v69*int32(3)+int32(16))*v67 + int32(18)
	if base.Ui32(v108) < base.Ui32(v121) {
		goto L2
	} else {
		goto L45
	}
L35:
	;
	v87 = int32(18)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v89 == v87 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v100 = int32(1)
	if v83&v100 != 0 {
		v108 = int32(base.Ui32(v83) >> (uint(v100) % 32))
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v92 = v87
	goto L40
L39:
	;
	v92 = int32(2)
	goto L40
L40:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v99 = int32(6)
	goto L43
L42:
	;
	v99 = v92
	goto L43
L43:
	;
	v108 = v99
	goto L34
L44:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = int32(base.Ui32(v104) >> (uint(int32(2)) % 32))
	goto L34
L45:
	;
	v124 = v62 + int32(14)
	if v110 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, v52+int32(16), v124, v110)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v128 = v124 + v110
	v129 = F_palloc(m, v112)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	if v112 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	base.MemoryCopy(m, v129, v128, v112)
	goto L52
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v69) < base.Ui32(int32(4)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v251 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v202 = v180
	v203 = v181
	v204 = int32(0)
	goto L62
L55:
	;
	v180 = int32(0)
	v181 = v121
	goto L54
L56:
	;
	goto L57
L57:
	;
	v140 = int32(0)
	v141 = v121
	v144 = v2
	goto L58
L58:
	;
	v161 = v129 + v140*int32(20)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+64))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)+44))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v169 = v162 + (v163 + (v164 + (v165 + v141)))
	v170 = int32(4)
	v171 = v140 + v170
	v173 = v144 + v170
	if v173 != v69&int32(12) {
		v140 = v171
		v141 = v169
		v144 = v173
		goto L58
	} else {
		goto L60
	}
L59:
	;
	if v69&int32(3) == int32(0) {
		v233 = v169
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v180 = v171
	v181 = v169
	goto L54
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v129+v202*int32(20))+4))
	v225 = v224 + v203
	v226 = int32(1)
	v229 = v204 + v226
	if v229 != v69&int32(3) {
		v202 = v202 + v226
		v203 = v225
		v204 = v229
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v233 = v225
	goto L53
L64:
	;
	goto L63
L65:
	;
	if v276 != v233 {
		goto L1
	} else {
		goto L76
	}
L66:
	;
	v255 = int32(18)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v257 == v255 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v268 = int32(1)
	if v251&v268 != 0 {
		v276 = int32(base.Ui32(v251) >> (uint(v268) % 32))
		goto L65
	} else {
		goto L75
	}
L69:
	;
	v260 = v255
	goto L71
L70:
	;
	v260 = int32(2)
	goto L71
L71:
	;
	if base.Ui32((v257-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v267 = int32(6)
	goto L74
L73:
	;
	v267 = v260
	goto L74
L74:
	;
	v276 = v267
	goto L65
L75:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v276 = int32(base.Ui32(v272) >> (uint(int32(2)) % 32))
	goto L65
L76:
	;
	v281 = F_palloc_mul(m, int32(4), v69)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L23
	} else {
		goto L77
	}
L77:
	;
	v284 = int32(0)
	v286 = int32(0)
	goto L78
L78:
	;
	v310 = v129 + v286*int32(20)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = F_palloc_mul(m, int32(8), v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L23
	} else {
		goto L80
	}
L79:
	;
	v320 = int32(7)
	v322 = int32(24)
	v323 = (v69 + v320) & v322
	v324 = v323 * v67
	v326 = v69 << (uint(int32(3)) % 32)
	v327 = v326 * v67
	v333 = (v67*v322 + int32(55)) & int32(_a_F_statext_mcv_deserialize_1)
	v341 = F_repalloc(m, v52, v324+(v327+v333)+(v316+v320)&int32(-8))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L23
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281+v286<<(uint(int32(2))%32)))) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	v316 = v315 + v284
	v318 = v286 + int32(1)
	if v318 != v69 {
		v284 = v316
		v286 = v318
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v343 = v341 + v333
	v344 = v343 + v327
	v348 = v112 + v128
	v350 = v344 + v324
	v354 = int32(0)
	goto L83
L83:
	;
	v369 = v129 + v354*int32(20)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+16)))
	if v370 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v615 = int32(1)
	if v67 <= v615 {
		goto L133
	} else {
		goto L134
	}
L85:
	;
	v613 = v354 + int32(1)
	if v613 != v69 {
		v348 = v593
		v350 = v595
		v354 = v613
		goto L83
	} else {
		goto L132
	}
L86:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v373 <= int32(0) {
		v593 = v348
		v595 = v350
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	if int32(0) < v440 {
		goto L106
	} else {
		goto L107
	}
L89:
	;
	v381 = v348
	v385 = int32(0)
	goto L90
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = int64(0)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	if v402 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v593 = v430
	v595 = v350
	goto L85
L92:
	;
	base.MemoryCopy(m, v23+int32(120), v381, v402)
	goto L94
L93:
	;
	goto L94
L94:
	;
	if base.I32_popcnt(v402) != int32(1) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v430 = v381 + v402
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v281+v354<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v431+v385<<(uint(int32(3))%32)))) = v429
	v437 = v385 + int32(1)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v437 < v438 {
		v381 = v430
		v385 = v437
		goto L90
	} else {
		goto L105
	}
L96:
	;
	v428 = int64(*(*int8)(unsafe.Add(mBase, uint32(v23)+120)))
	v429 = v428
	goto L95
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L23
	} else {
		goto L102
	}
L98:
	;
	switch base.I32_ctz(v402) {
	case 0:
		goto L96
	case 1:
		goto L101
	case 2:
		goto L100
	case 3:
		goto L99
	default:
		goto L97
	}
L99:
	;
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v23)+120))
	v429 = v412
	goto L95
L100:
	;
	v411 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+120)))
	v429 = v411
	goto L95
L101:
	;
	v410 = int64(*(*int16)(unsafe.Add(mBase, uint32(v23)+120)))
	v429 = v410
	goto L95
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v402
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_2), v23-int32(-64))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L23
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_3), int32(123), int32(_a_F_statext_mcv_deserialize_4))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L23
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	goto L91
L106:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v443 <= int32(0) {
		v593 = v348
		v595 = v350
		goto L85
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	switch v440 + int32(2) {
	case 0:
		goto L117
	case 1:
		goto L116
	default:
		v593 = v348
		v595 = v350
		goto L85
	}
L109:
	;
	v451 = v348
	v452 = int32(0)
	v453 = v350
	v455 = v440
	goto L110
L110:
	;
	if v455 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v593 = v478
	v595 = v484
	goto L85
L112:
	;
	base.MemoryCopy(m, v453, v451, v455)
	goto L114
L113:
	;
	goto L114
L114:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v281+v354<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v472+v452<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v453)
	v478 = v451 + v471
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v484 = v453 + (v479+int32(7))&int32(-8)
	v486 = v452 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v486 < v487 {
		v451 = v478
		v452 = v486
		v453 = v484
		v455 = v479
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v538 <= int32(0) {
		v593 = v348
		v595 = v350
		goto L85
	} else {
		goto L125
	}
L117:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v491 <= int32(0) {
		v593 = v348
		v595 = v350
		goto L85
	} else {
		goto L118
	}
L118:
	;
	v499 = v348
	v501 = v350
	v503 = int32(0)
	goto L119
L119:
	;
	v519 = v499 + int32(4)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	if v520 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v593 = v528
	v595 = v533
	goto L85
L121:
	;
	base.MemoryCopy(m, v501, v519, v520)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v281+v354<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v522+v503<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v501)
	v528 = v520 + v519
	v533 = v501 + (v520+int32(7))&int32(-8)
	v535 = v503 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v535 < v536 {
		v499 = v528
		v501 = v533
		v503 = v535
		goto L119
	} else {
		goto L124
	}
L124:
	;
	goto L120
L125:
	;
	v546 = v348
	v548 = v350
	v550 = int32(0)
	goto L126
L126:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = v565<<(uint(int32(2))%32) + int32(16)
	v572 = v546 + int32(4)
	if v565 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v593 = v582
	v595 = v587
	goto L85
L128:
	;
	base.MemoryCopy(m, v548+int32(4), v572, v565)
	goto L130
L129:
	;
	goto L130
L130:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v281+v354<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v576+v550<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v548)
	v582 = v572 + v565
	v587 = v548 + (v565+int32(11))&int32(-8)
	v589 = v550 + int32(1)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v589 < v590 {
		v546 = v582
		v548 = v587
		v550 = v589
		goto L126
	} else {
		goto L131
	}
L131:
	;
	goto L127
L132:
	;
	goto L84
L133:
	;
	v618 = v615
	goto L135
L134:
	;
	v618 = v67
	goto L135
L135:
	;
	v621 = int32(1)
	if v69 <= v621 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v624 = v621
	goto L138
L137:
	;
	v624 = v69
	goto L138
L138:
	;
	v633 = v593
	v639 = int32(0)
	v643 = v344
	v645 = v343
	goto L139
L139:
	;
	v654 = v341 + int32(48) + v639*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v654)+16)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v654)+20)) = v645
	if v69 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v804 = int32(1)
	if v69 <= v804 {
		goto L162
	} else {
		goto L163
	}
L141:
	;
	base.MemoryCopy(m, v643, v633, v69)
	goto L143
L142:
	;
	goto L143
L143:
	;
	v658 = v633 + v69
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v658)))
	*(*int64)(unsafe.Add(mBase, uint32(v654))) = v659
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v658)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v654)+8)) = v661
	v664 = v658 + int32(16)
	v665 = int32(0)
	if base.B2i32(v69 < int32(2)) == v665 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v802 = v639 + int32(1)
	if v802 != v618 {
		v633 = v780
		v639 = v802
		v643 = v643 + v323
		v645 = v645 + v326
		goto L139
	} else {
		goto L161
	}
L145:
	;
	v670 = v664
	v671 = v665
	v672 = v665
	goto L148
L146:
	;
	v739 = v664
	v740 = v665
	goto L147
L147:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v654)+16))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758+v740))))
	if v760 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v654)+16))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689+v671))))
	if v691 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	if v624&int32(1) == int32(0) {
		v780 = v730
		goto L144
	} else {
		goto L157
	}
L150:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v654)+20))
	v695 = int32(3)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v281+v671<<(uint(int32(2))%32))))
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670))))
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v701+v702<<(uint(v695)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v694+v671<<(uint(v695)%32)))) = v706
	goto L152
L151:
	;
	goto L152
L152:
	;
	v709 = v671 | int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v654)+16))
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709+v710))))
	if v712 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v654)+20))
	v716 = int32(3)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v281+v709<<(uint(int32(2))%32))))
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+2)))
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v722+v723<<(uint(v716)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v715+v709<<(uint(v716)%32)))) = v727
	goto L155
L154:
	;
	goto L155
L155:
	;
	v730 = v670 + int32(4)
	v731 = int32(2)
	v732 = v671 + v731
	v734 = v672 + v731
	if v734 != v624&int32(14) {
		v670 = v730
		v671 = v732
		v672 = v734
		goto L148
	} else {
		goto L156
	}
L156:
	;
	goto L149
L157:
	;
	v739 = v730
	v740 = v732
	goto L147
L158:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v654)+20))
	v764 = int32(3)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v281+v740<<(uint(int32(2))%32))))
	v771 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v739))))
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v770+v771<<(uint(v764)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v763+v740<<(uint(v764)%32)))) = v775
	goto L160
L159:
	;
	goto L160
L160:
	;
	v780 = v739 + int32(2)
	goto L144
L161:
	;
	goto L140
L162:
	;
	v807 = v804
	goto L164
L163:
	;
	v807 = v69
	goto L164
L164:
	;
	v810 = int32(0)
	goto L165
L165:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v281+v810<<(uint(int32(2))%32))))
	F_pfree(m, v832)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L23
	} else {
		goto L167
	}
L166:
	;
	F_pfree(m, v281)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L23
	} else {
		goto L169
	}
L167:
	;
	v836 = v810 + int32(1)
	if v836 != v807 {
		v810 = v836
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	v854 = v341
	goto L12
L170:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v870 == int32(1) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v895
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_5), v23)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L23
	} else {
		goto L184
	}
L172:
	;
	v874 = int32(18)
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v876 == v874 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	if v870&int32(1) != 0 {
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v879 = v874
	goto L177
L176:
	;
	v879 = int32(2)
	goto L177
L177:
	;
	if base.Ui32((v876-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v886 = int32(6)
	goto L180
L179:
	;
	v886 = v879
	goto L180
L180:
	;
	v895 = v886
	goto L171
L181:
	;
	v895 = int32(base.Ui32(v870) >> (uint(int32(1)) % 32))
	goto L171
L182:
	;
	goto L183
L183:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v895 = int32(base.Ui32(v891) >> (uint(int32(2)) % 32))
	goto L171
L184:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1030), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L23
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = int32(-509193790)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v911
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_8), v23+int32(112))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L23
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1055), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L23
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v929
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_9), v23+int32(96))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L23
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1059), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L23
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_10), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L23
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1062), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L23
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v960 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v960
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_11), v23+int32(16))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L23
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1066), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L23
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_12), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L23
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1069), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L23
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v989
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_13), v23+int32(32))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L23
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1072), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L23
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1006 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1031
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_14), v23+int32(48))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L23
	} else {
		goto L218
	}
L206:
	;
	v1010 = int32(18)
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v1012 == v1010 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L208
L208:
	;
	if v1006&int32(1) != 0 {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	v1015 = v1010
	goto L211
L210:
	;
	v1015 = int32(2)
	goto L211
L211:
	;
	if base.Ui32((v1012-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1022 = int32(6)
	goto L214
L213:
	;
	v1022 = v1015
	goto L214
L214:
	;
	v1031 = v1022
	goto L205
L215:
	;
	v1031 = int32(base.Ui32(v1006) >> (uint(int32(1)) % 32))
	goto L205
L216:
	;
	goto L217
L217:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1031 = int32(base.Ui32(v1027) >> (uint(int32(2)) % 32))
	goto L205
L218:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1091), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L23
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1049 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1074
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_14), v23+int32(80))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L23
	} else {
		goto L234
	}
L222:
	;
	v1053 = int32(18)
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v1055 == v1053 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	if v1049&int32(1) != 0 {
		goto L231
	} else {
		goto L232
	}
L225:
	;
	v1058 = v1053
	goto L227
L226:
	;
	v1058 = int32(2)
	goto L227
L227:
	;
	if base.Ui32((v1055-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1065 = int32(6)
	goto L230
L229:
	;
	v1065 = v1058
	goto L230
L230:
	;
	v1074 = v1065
	goto L221
L231:
	;
	v1074 = int32(base.Ui32(v1049) >> (uint(int32(1)) % 32))
	goto L221
L232:
	;
	goto L233
L233:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1074 = int32(base.Ui32(v1070) >> (uint(int32(2)) % 32))
	goto L221
L234:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_6), int32(1123), int32(_a_F_statext_mcv_deserialize_7))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L23
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_statext_ndistinct_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L37
	} else {
		goto L71
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L37
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L37
	} else {
		goto L65
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L37
	} else {
		goto L62
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L37
	} else {
		goto L46
	}
L6:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v141 = int32(0)
	goto L8
L8:
	;
	m.G0 = v11 - int32(-64)
	return v141
L9:
	;
	if base.Ui32(v40) <= base.Ui32(int32(11)) {
		goto L5
	} else {
		goto L18
	}
L10:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v16-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v28 = int32(1)
	if v13&v28 != 0 {
		v40 = int32(base.Ui32(v13)>>(uint(v28)%32)) - v28
		goto L9
	} else {
		goto L17
	}
L13:
	;
	if v16 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v27 = int32(16)
	goto L16
L15:
	;
	v27 = int32(0)
	goto L16
L16:
	;
	v40 = v27
	goto L9
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L18:
	;
	v43 = int32(1)
	if v13&v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v47 = v43
	goto L21
L20:
	;
	v47 = int32(4)
	goto L21
L21:
	;
	v48 = l0 + v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != int32(-1554858076) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v52 != int32(1) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v55 == int32(0) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if v13 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v88 = v55 << (uint(int32(4)) % 32)
	v90 = v88 | int32(12)
	if base.Ui32(v86) < base.Ui32(v90) {
		goto L1
	} else {
		goto L36
	}
L26:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v63 == int32(18) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v74 = int32(1)
	if v13&v74 != 0 {
		v86 = int32(base.Ui32(v13)>>(uint(v74)%32)) - v74
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v66 = int32(16)
	goto L31
L30:
	;
	v66 = int32(0)
	goto L31
L31:
	;
	if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v73 = int32(4)
	goto L34
L33:
	;
	v73 = v66
	goto L34
L34:
	;
	v86 = v73
	goto L25
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v96 = F_palloc0(m, v88+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = int64(7035076516)
	v106 = v48 + int32(12)
	v110 = int32(0)
	goto L39
L39:
	;
	v116 = v96 + int32(16) + v110<<(uint(int32(4))%32)
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v119
	v123 = F_palloc(m, v119<<(uint(int32(1))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v141 = v96
	goto L8
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v123
	v127 = v106 + int32(12)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v130 = v128 << (uint(int32(1)) % 32)
	if v130 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	base.MemoryCopy(m, v123, v127, v130)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v133 = int32(1)
	v137 = v110 + v133
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if base.Ui32(v137) < base.Ui32(v138) {
		v106 = v127 + v132<<(uint(v133)%32)
		v110 = v137
		goto L39
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v158 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v187
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_0), v11)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L37
	} else {
		goto L60
	}
L48:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v164 == int32(18) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v158&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v167 = int32(16)
	goto L53
L52:
	;
	v167 = int32(0)
	goto L53
L53:
	;
	if base.Ui32((v164-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v174 = int32(4)
	goto L56
L55:
	;
	v174 = v167
	goto L56
L56:
	;
	v187 = v174
	goto L47
L57:
	;
	v177 = int32(1)
	v187 = int32(base.Ui32(v158)>>(uint(v177)%32)) - v177
	goto L47
L58:
	;
	goto L59
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = int32(base.Ui32(v181)>>(uint(int32(2))%32)) - int32(4)
	goto L47
L60:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(261), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L37
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(-1554858076)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v49
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_3), v9+int32(-16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L37
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(276), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L37
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v52
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_4), v9+int32(-32))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L37
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(279), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L37
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_5), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L37
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(281), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L37
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v251 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v280
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_0), v9+int32(-48))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L37
	} else {
		goto L85
	}
L73:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v257 == int32(18) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if v251&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v260 = int32(16)
	goto L78
L77:
	;
	v260 = int32(0)
	goto L78
L78:
	;
	if base.Ui32((v257-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v267 = int32(4)
	goto L81
L80:
	;
	v267 = v260
	goto L81
L81:
	;
	v280 = v267
	goto L72
L82:
	;
	v270 = int32(1)
	v280 = int32(base.Ui32(v251)>>(uint(v270)%32)) - v270
	goto L72
L83:
	;
	goto L84
L84:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = int32(base.Ui32(v274)>>(uint(int32(2))%32)) - int32(4)
	goto L72
L85:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(287), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L37
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
