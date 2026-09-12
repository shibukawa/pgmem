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
		F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(32768))>>(uint(int32(15))%32)))
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
			F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(32768))>>(uint(int32(15))%32)))
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
	var v43 int32
	_ = v43
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
	var v73 int32
	_ = v73
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
	var v178 int32
	_ = v178
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
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
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
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
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)) = uint8(v5)
	v43 = v5
	goto L3
L1:
	;
	if v793 != 0 {
		goto L150
	} else {
		goto L151
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L6
	} else {
		goto L147
	}
L3:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v52 == int32(0) {
		v793 = v43
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L6
	} else {
		goto L144
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
		v793 = v43
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = v55 + int32(28)
	v73 = v43
	goto L9
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v247)+6))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v748
	v750 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v750)
	v73 = v132
	goto L9
L17:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+35)))
	if v719 == int32(1) {
		goto L125
	} else {
		goto L126
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
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)) = uint8(v97)
	v711 = v73
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
	if v73 < int32(0) {
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
		v132 = v73
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v73^int32(-1))<<(uint(int32(6))%32))+16))
	v123 = v114
	goto L26
L28:
	;
	goto L29
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+v73<<(uint(int32(6))%32)+int32(-64))+16))
	v123 = v122
	goto L26
L30:
	;
	F_UnlockReleaseBuffer(m, v73)
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
	v136 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v132^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L34
L36:
	;
	goto L37
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v104<<(uint(int32(2))%32)+v150)+20))
	v247 = v150 + v244&int32(32767)
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
		v711 = v132
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v212 = v104
	goto L50
L44:
	;
	v171 = int32(base.Ui32(v162+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v171 == int32(0) {
		v711 = v132
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v178 = int32(1)
	goto L46
L46:
	;
	v201 = F_spgTestLeafTuple(m, l1, v55, v150, v178&int32(65535), base.B2i32(v155 != int32(0)), int32(1), v23+int32(7), l3)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L48
	}
L47:
	;
	v711 = v132
	goto L17
L48:
	;
	v204 = v178 + int32(1)
	if base.Ui32(v204&int32(65535)) <= base.Ui32(v171) {
		v178 = v204
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v229 = v212 & int32(65535)
	if v229 == int32(0) {
		v711 = v132
		goto L17
	} else {
		goto L52
	}
L51:
	;
	v73 = v132
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
		v212 = v237
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
	v253 = int32(4470560)
	v254 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v23-int32(-64)))) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v259
	v268 = int32(base.Ui32(v258) >> (uint(int32(3)) % 32))
	v270 = v268 & int32(8191)
	if v155 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v483&int32(4) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L58:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v287
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+208)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+42)) = uint8(base.B2i32(base.Ui32(int32(65535)) < base.Ui32(v258)))
	v296 = int32(base.Ui32(v258)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+41)) = uint8(v296)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+40)) = uint8(v289)
	if base.Ui32(v258) < base.Ui32(int32(65536)) {
		v309 = int32(0)
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v270
	v326 = F_palloc(m, v270<<(uint(int32(2))%32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v309
	v312 = F_spgExtractNodeLabels(m, l1, v247)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	v303 = v247 + int32(8)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v304 != int32(1) {
		v309 = v303
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v309 = v307
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v320 = F_FunctionCall2Coll(m, l1+int32(132), v315, v23+int32(8), v23+int32(56))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v482 = v322
	goto L57
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v326
	v329 = int32(0)
	if v270 == v329 {
		v482 = v329
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v333 = v268 & int32(7)
	v334 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v270) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v345 = v334
	v350 = int32(0)
	goto L71
L69:
	;
	v416 = v334
	goto L70
L70:
	;
	if v333 == int32(0) {
		v482 = v270
		goto L57
	} else {
		goto L74
	}
L71:
	;
	v360 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v345<<(uint(v360)%32)))) = v345
	v365 = v345 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v365<<(uint(v360)%32)))) = v365
	v371 = v345 | v360
	*(*int32)(unsafe.Add(mBase, uint32(v326+v371<<(uint(v360)%32)))) = v371
	v377 = v345 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v377<<(uint(v360)%32)))) = v377
	v383 = v345 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v383<<(uint(v360)%32)))) = v383
	v389 = v345 | int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v389<<(uint(v360)%32)))) = v389
	v395 = v345 | int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v395<<(uint(v360)%32)))) = v395
	v401 = v345 | int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v326+v401<<(uint(v360)%32)))) = v401
	v406 = int32(8)
	v407 = v345 + v406
	v409 = v350 + v406
	if v409 != v270-v333 {
		v345 = v407
		v350 = v409
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v416 = v407
	goto L70
L73:
	;
	goto L72
L74:
	;
	v438 = v416
	v446 = v334
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326+v438<<(uint(int32(2))%32)))) = v438
	v457 = int32(1)
	v460 = v446 + v457
	if v460 != v333 {
		v438 = v438 + v457
		v446 = v460
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v482 = v270
	goto L57
L77:
	;
	goto L76
L78:
	;
	if v482 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	if v482 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	if v482 != v270 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v254
	v711 = v132
	goto L17
L83:
	;
	v495 = F_palloc(m, v270<<(uint(int32(2))%32))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v497&int32(65528) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v511 = v247 + int32(base.Ui32(v497)>>(uint(int32(16))%32)) + int32(8)
	v512 = int32(0)
	goto L88
L86:
	;
	goto L87
L87:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v563
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	if v565 <= int32(0) {
		goto L82
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v495+v512<<(uint(int32(2))%32)))) = v511
	v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+6)))
	v531 = int32(8191)
	v535 = v512 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if base.Ui32(v535) < base.Ui32(int32(base.Ui32(v536)>>(uint(int32(3))%32))&v531) {
		v511 = v511 + v530&v531
		v512 = v535
		goto L88
	} else {
		goto L90
	}
L89:
	;
	goto L87
L90:
	;
	goto L89
L91:
	;
	v569 = int32(base.Ui32(v155) >> (uint(int32(3)) % 32))
	v575 = v565
	v576 = int32(0)
	goto L92
L92:
	;
	v591 = int32(2)
	v592 = v576 << (uint(v591) % 32)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v592+v593)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v495+v595<<(uint(v591)%32))))
	if v599 == int32(0) {
		v670 = v575
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L82
L94:
	;
	v675 = v576 + int32(1)
	if v675 < v670 {
		v575 = v670
		v576 = v675
		goto L92
	} else {
		goto L122
	}
L95:
	;
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v599)+4)))
	if v602 == int32(0) {
		v670 = v575
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if v155 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	*(*int32)(unsafe.Add(mBase, uint32(v630)+28)) = v633
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v599)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v630)+32)) = uint16(v635)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v638 != 0 {
		goto L111
	} else {
		goto L112
	}
L98:
	;
	v606 = F_palloc(m, int32(40))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	if v609 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v606)+34)) = uint8(v569)
	v630 = v606
	goto L97
L102:
	;
	v611 = v609 + v592
	goto L104
L103:
	;
	v611 = l1 + int32(192)
	goto L104
L104:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v618 = F_palloc(m, v613<<(uint(int32(3))%32)+int32(40))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v618)+34)) = uint8(v569)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v621 <= int32(0) {
		v630 = v618
		goto L97
	} else {
		goto L106
	}
L106:
	;
	v627 = v621 << (uint(int32(3)) % 32)
	if v627 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v630 = v618
	goto L97
L108:
	;
	v628 = F__emscripten_memcpy_bulkmem(m, v618+int32(40), v612, v627)
	mBase = m.M
	goto L110
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v638+v592)))
	v642 = v640 + v637
	goto L113
L112:
	;
	v642 = v637
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+24)) = v642
	v644 = int32(0)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	if v646 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v592+v646)))
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	v650 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	v651 = F_datumCopy(m, v648, v649, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	v653 = v644
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v630)+12)) = v653
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	if v657 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v653 = v651
	goto L116
L118:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v657+v592)))
	v660 = v659
	goto L120
L119:
	;
	v660 = v644
	goto L120
L120:
	;
	v661 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v630)+37)) = uint8(v661)
	*(*uint16)(unsafe.Add(mBase, uint32(v630)+35)) = uint16(v661)
	*(*int32)(unsafe.Add(mBase, uint32(v630)+20)) = v660
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	F_pairingheap_add(m, v666, v630)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v670 = v669
	goto L94
L122:
	;
	goto L93
L123:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v732 != 0 {
		goto L132
	} else {
		goto L133
	}
L124:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v726 == int32(0) {
		goto L123
	} else {
		goto L130
	}
L125:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v722 == int32(0) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v725 != 0 {
		goto L123
	} else {
		goto L129
	}
L128:
	;
	goto L123
L129:
	;
	goto L124
L130:
	;
	F_pfree(m, v726)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	goto L123
L132:
	;
	F_pfree(m, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v735 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	F_pfree(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L6
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_pfree(m, v55)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	F_MemoryContextReset(m, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	if l2 != 0 {
		v43 = v711
		goto L3
	} else {
		goto L142
	}
L142:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+7)))
	if v743&int32(1) == int32(0) {
		v43 = v711
		goto L3
	} else {
		goto L143
	}
L143:
	;
	v793 = v711
	goto L1
L144:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v756 & int32(3)
	F_errmsg_internal(m, int32(475646), v23)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(488373), int32(911), int32(310700))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L6
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
	F_errmsg_internal(m, int32(377357), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(488373), int32(700), int32(76809))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L6
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
	F_UnlockReleaseBuffer(m, v793)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L6
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	m.G0 = v23 + int32(80)
	return
L153:
	;
	goto L152
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
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = l0 - int32(16)
	if base.Ui32(v4) <= base.Ui32(int32(127)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v4)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_consts[78])))
		v14 = v13
	} else {
		v14 = int32(0)
	}
	return v14
}
