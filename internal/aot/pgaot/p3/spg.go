package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgDeformLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if l4 == int32(0) {
		F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(_a_F_spgDeformLeafTuple_0))>>(uint(int32(15))%32)))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v10 != int32(1) {
			F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(_a_F_spgDeformLeafTuple_0))>>(uint(int32(15))%32)))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
			return
		}
	}
}
func F_spgWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)) = uint8(v5)
	v44 = v5
	goto L3
L1:
	;
	if v798 != 0 {
		goto L144
	} else {
		goto L145
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L141
	}
L3:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v52 == int32(0) {
		v798 = v44
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L6
	} else {
		goto L138
	}
L5:
	;
	v55 = F_pairingheap_remove_first(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	if v55 == int32(0) {
		v798 = v44
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = v55 + int32(28)
	v74 = v44
	goto L9
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[0]))
	if v82 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+35)))
	if v85 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L13
L15:
	;
	goto L10
L16:
	;
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v752)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v247)+6))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v754
	v74 = v132
	goto L9
L17:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+35)))
	if v723 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+34)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+36)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+37)))
	m.T0[l3].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l1, v60, v88, v89, v90, v91, v92, v55+int32(40))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+30)))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+28)))
	v103 = v99 | v100<<(uint(int32(16))%32)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+32)))
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)) = uint8(v97)
	v716 = v74
	goto L17
L22:
	;
	if v132 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	if v74 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v127 = F_ReadBuffer(m, l0, v103)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L32
	}
L26:
	;
	if v123 == v103 {
		v132 = v74
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[1]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v74^int32(-1))<<(uint(int32(6))%32))+16))
	v123 = v114
	goto L26
L28:
	;
	goto L29
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[2]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+v74<<(uint(int32(6))%32)+int32(-64))+16))
	v123 = v122
	goto L26
L30:
	;
	F_UnlockReleaseBuffer(m, v74)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	F_LockBuffer(m, v127, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v132 = v127
	goto L22
L34:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+16)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151+v150))))
	v155 = v153 & int32(8)
	if v153&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[3]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v132^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L34
L36:
	;
	goto L37
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[4]))
	v150 = v144 + v132<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v158 = int32(1)
	if base.Ui32(v103-v158) <= base.Ui32(v158) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v150+v104<<(uint(int32(2))%32))+20))
	v247 = v150 + v244&int32(_a_F_spgWalk_0)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v250 = v248 & int32(3)
	if v250 == int32(1) {
		goto L16
	} else {
		goto L55
	}
L41:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+12)))
	if base.Ui32(v162) < base.Ui32(int32(25)) {
		v716 = v132
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v214 = v104
	goto L50
L44:
	;
	v171 = int32(base.Ui32(v162+int32(_a_F_spgWalk_1))>>(uint(int32(2))%32)) & int32(_a_F_spgWalk_2)
	if v171 == int32(0) {
		v716 = v132
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v180 = int32(1)
	goto L46
L46:
	;
	v201 = F_spgTestLeafTuple(m, l1, v55, v150, v180&int32(_a_F_spgWalk_2), base.B2i32(v155 != int32(0)), int32(1), v23+int32(7), l3)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L48
	}
L47:
	;
	v716 = v132
	goto L17
L48:
	;
	v204 = v180 + int32(1)
	if base.Ui32(v204&int32(_a_F_spgWalk_2)) <= base.Ui32(v171) {
		v180 = v204
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v229 = v214 & int32(_a_F_spgWalk_2)
	if v229 == int32(0) {
		v716 = v132
		goto L17
	} else {
		goto L52
	}
L51:
	;
	v74 = v132
	goto L9
L52:
	;
	v232 = int32(0)
	v237 = F_spgTestLeafTuple(m, l1, v55, v150, v229, base.B2i32(v155 != v232), v232, v23+int32(7), l3)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	if v237 != int32(2049) {
		v214 = v237
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	if v250 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v253 = int32(_a_F_spgWalk_3)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v259
	v266 = int32(base.Ui32(v258) >> (uint(int32(3)) % 32))
	v268 = v266 & int32(_a_F_spgWalk_4)
	if v155 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v485 = int32(0)
	if base.B2i32(base.B2i32(v482&int32(4) == v485)|base.B2i32(v481 == v485) == v485)&base.B2i32(v481 != v268) != 0 {
		goto L2
	} else {
		goto L78
	}
L58:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v285
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+208)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+42)) = uint8(base.B2i32(base.Ui32(int32(_a_F_spgWalk_2)) < base.Ui32(v258)))
	v294 = int32(base.Ui32(v258)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+41)) = uint8(v294)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)) = uint8(v287)
	if base.Ui32(v258) < base.Ui32(int32(_a_F_spgWalk_5)) {
		v307 = int32(0)
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v268
	v324 = F_palloc(m, v268<<(uint(int32(2))%32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v307
	v310 = F_spgExtractNodeLabels(m, l1, v247)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	v301 = v247 + int32(8)
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v302 != int32(1) {
		v307 = v301
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v307 = v305
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v310
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v318 = F_FunctionCall2Coll(m, l1+int32(132), v313, v23+int32(8), v23+int32(56))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v481 = v320
	goto L57
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v324
	v327 = int32(0)
	if v268 == v327 {
		v481 = v327
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v331 = v266 & int32(7)
	v332 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v268) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v343 = v332
	v351 = int32(0)
	goto L71
L69:
	;
	v416 = v332
	goto L70
L70:
	;
	v436 = v416
	v443 = v332
	goto L75
L71:
	;
	v359 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v343<<(uint(v359)%32)))) = v343
	v364 = v343 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v364<<(uint(v359)%32)))) = v364
	v370 = v343 | v359
	*(*int32)(unsafe.Add(mBase, uint32(v324+v370<<(uint(v359)%32)))) = v370
	v376 = v343 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v376<<(uint(v359)%32)))) = v376
	v382 = v343 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v382<<(uint(v359)%32)))) = v382
	v388 = v343 | int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v388<<(uint(v359)%32)))) = v388
	v394 = v343 | int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v394<<(uint(v359)%32)))) = v394
	v400 = v343 | int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v324+v400<<(uint(v359)%32)))) = v400
	v405 = int32(8)
	v406 = v343 + v405
	v408 = v351 + v405
	if v408 != v266&int32(_a_F_spgWalk_6) {
		v343 = v406
		v351 = v408
		goto L71
	} else {
		goto L73
	}
L72:
	;
	if v331 == int32(0) {
		v481 = v268
		goto L57
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v416 = v406
	goto L70
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v436<<(uint(int32(2))%32)))) = v436
	v456 = int32(1)
	v459 = v443 + v456
	if v459 != v331 {
		v436 = v436 + v456
		v443 = v459
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v481 = v268
	goto L57
L77:
	;
	goto L76
L78:
	;
	if v481 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v254
	v716 = v132
	goto L17
L80:
	;
	v498 = F_palloc(m, v268<<(uint(int32(2))%32))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v500&int32(_a_F_spgWalk_7) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v513 = v247 + int32(base.Ui32(v500)>>(uint(int32(16))%32)) + int32(8)
	v516 = int32(0)
	goto L85
L83:
	;
	goto L84
L84:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if v568 <= int32(0) {
		goto L79
	} else {
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v498+v516<<(uint(int32(2))%32)))) = v513
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+6)))
	v534 = int32(_a_F_spgWalk_4)
	v538 = v516 + int32(1)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if base.Ui32(v538) < base.Ui32(int32(base.Ui32(v539)>>(uint(int32(3))%32))&v534) {
		v513 = v513 + v533&v534
		v516 = v538
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L84
L87:
	;
	goto L86
L88:
	;
	v572 = int32(base.Ui32(v155) >> (uint(int32(3)) % 32))
	v578 = int32(0)
	v580 = v568
	goto L89
L89:
	;
	v594 = int32(2)
	v595 = v578 << (uint(v594) % 32)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v595+v596)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v498+v598<<(uint(v594)%32))))
	if v602 == int32(0) {
		v674 = v580
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L79
L91:
	;
	v679 = v578 + int32(1)
	if v679 < v674 {
		v578 = v679
		v580 = v674
		goto L89
	} else {
		goto L116
	}
L92:
	;
	v605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602)+4)))
	if v605 == int32(0) {
		v674 = v580
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if v155 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+32)) = uint16(v637)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+28)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v642 != 0 {
		goto L105
	} else {
		goto L106
	}
L95:
	;
	v609 = F_palloc(m, int32(40))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L6
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	if v612 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v609)+34)) = uint8(v572)
	v634 = v609
	goto L94
L99:
	;
	v614 = v612 + v595
	goto L101
L100:
	;
	v614 = l1 + int32(192)
	goto L101
L101:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v621 = F_palloc(m, v616<<(uint(int32(3))%32)+int32(40))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+34)) = uint8(v572)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v624 <= int32(0) {
		v634 = v621
		goto L94
	} else {
		goto L103
	}
L103:
	;
	v628 = v624 << (uint(int32(3)) % 32)
	if v628 == int32(0) {
		v634 = v621
		goto L94
	} else {
		goto L104
	}
L104:
	;
	base.MemoryCopy(m, v621+int32(40), v615, v628)
	v634 = v621
	goto L94
L105:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v595+v642)))
	v646 = v644 + v641
	goto L107
L106:
	;
	v646 = v641
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+24)) = v646
	v648 = int32(0)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	if v650 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v595+v650)))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	v654 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	v655 = F_datumCopy(m, v652, v653, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L111
	}
L109:
	;
	v657 = v648
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v634)+12)) = v657
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	if v661 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v657 = v655
	goto L110
L112:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v595+v661)))
	v664 = v663
	goto L114
L113:
	;
	v664 = v648
	goto L114
L114:
	;
	v665 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v634)+37)) = uint8(v665)
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+35)) = uint16(v665)
	*(*int32)(unsafe.Add(mBase, uint32(v634)+20)) = v664
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	F_pairingheap_add(m, v670, v634)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v674 = v673
	goto L91
L116:
	;
	goto L90
L117:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v736 != 0 {
		goto L126
	} else {
		goto L127
	}
L118:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v730 == int32(0) {
		goto L117
	} else {
		goto L124
	}
L119:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v726 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v729 != 0 {
		goto L117
	} else {
		goto L123
	}
L122:
	;
	goto L117
L123:
	;
	goto L118
L124:
	;
	F_pfree(m, v730)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	goto L117
L126:
	;
	F_pfree(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v739 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	F_pfree(m, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L6
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_pfree(m, v55)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L6
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	F_MemoryContextReset(m, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	if l2 != 0 {
		v44 = v716
		goto L3
	} else {
		goto L136
	}
L136:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
	if v747&int32(1) == int32(0) {
		v44 = v716
		goto L3
	} else {
		goto L137
	}
L137:
	;
	v798 = v716
	goto L1
L138:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v760 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgWalk_8), v23)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_spgWalk_9), int32(911), int32(_a_F_spgWalk_10))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L6
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
	F_errmsg_internal(m, int32(_a_F_spgWalk_11), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_spgWalk_9), int32(700), int32(_a_F_spgWalk_12))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L6
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
	F_UnlockReleaseBuffer(m, v798)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L6
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	m.G0 = v23 + int32(80)
	return
L147:
	;
	goto L146
}
func F_spg_bbox_quad_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(603)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500891)
	return v3
}
func F_spg_box_quad_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500891)
	return int32(0)
}
func F_spg_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = l0 - int32(16)
	if base.Ui32(v3) <= base.Ui32(int32(127)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_c_F_spg_identify[0])))
		v12 = v10
	} else {
		v12 = int32(0)
	}
	return v12
}
