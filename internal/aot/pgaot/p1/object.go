package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_alloc_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v158 int32
	_ = v158
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
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v495 int32
	_ = v495
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = v21 + l1<<(uint(int32(5))%32)
	v26 = v24 + int32(224)
	v28 = F_LWLockAcquire(m, v26, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+244))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L168
	}
L4:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	F_LWLockRelease(m, v825+v810<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L167
	}
L5:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_c_F_alloc_object[0]))))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v701 = l0
	v702 = l1
	v703 = v32
	v711 = v19
	goto L7
L7:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v701)+652))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+1468))
	if v717 != v719 {
		goto L141
	} else {
		goto L142
	}
L8:
	;
	v48 = l0 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = base.I32_div_u_s(int32(_a_F_alloc_object_0), v37)
	v46 = v41 - int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v45 = base.I32_div_u_s(int32(_a_F_alloc_object_1), v37)
	v46 = v45
	goto L8
L12:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v701 = l0
	v702 = l1
	v703 = v700
	v711 = v19
	goto L7
L13:
	;
	v56 = v49
	goto L16
L14:
	;
	goto L15
L15:
	;
	v514 = F_transfer_first_span(m, l0, v26, int32(2), int32(1))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L103
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+1468))
	if v68 != v70 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v495 != 0 {
		goto L12
	} else {
		goto L102
	}
L18:
	;
	v72 = int32(0)
	v76 = F_LWLockAcquire(m, v69+int32(1476), v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v158 = int32(base.Ui32(v56) >> (uint(int32(27)) % 32))
	v161 = l0 + v158*int32(20)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v162 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+1468))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v79 != v80 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = v72
	goto L25
L23:
	;
	v136 = v78
	goto L24
L24:
	;
	F_LWLockRelease(m, v136+int32(1476))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	v100 = v48 + v88*int32(20)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v101 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v79
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = v119
	goto L24
L27:
	;
	v115 = v88 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v115) <= base.Ui32(v116) {
		v88 = v115
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+24)))
	if v104 != int32(1) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	F_dsm_detach(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = int64(0)
	goto L27
L31:
	;
	goto L26
L32:
	;
	goto L20
L33:
	;
	v166 = v162
	goto L35
L34:
	;
	v163 = F_get_segment_by_index(m, l0, v158)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v169 = v166 + v56&int32(134217727)
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+24)))
	v174 = base.I32_div_u_s((v46-v170)*int32(3), v46)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	if v175 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v166 = v165
	goto L35
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+1468))
	if v176 != v178 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v295 = int32(0)
	goto L39
L39:
	;
	if v174 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	v180 = int32(0)
	v184 = F_LWLockAcquire(m, v177+int32(1476), v180)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v268 = int32(base.Ui32(v175) >> (uint(int32(27)) % 32))
	v271 = l0 + v268*int32(20)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	if v272 != 0 {
		goto L55
	} else {
		goto L56
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+1468))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v187 != v188 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v192 = v180
	goto L47
L45:
	;
	v244 = v186
	goto L46
L46:
	;
	F_LWLockRelease(m, v244+int32(1476))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	v208 = v48 + v192*int32(20)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	if v209 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v187
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = v227
	goto L46
L49:
	;
	v223 = v192 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v223) <= base.Ui32(v224) {
		v192 = v223
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+24)))
	if v212 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	F_dsm_detach(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(0)
	goto L49
L53:
	;
	goto L48
L54:
	;
	goto L42
L55:
	;
	v276 = v272
	goto L57
L56:
	;
	v273 = F_get_segment_by_index(m, l0, v268)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v295 = v276 + v175&int32(134217727)
	goto L39
L58:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v276 = v275
	goto L57
L59:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v56 == v298 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	goto L61
L61:
	;
	if v175 != 0 {
		v56 = v175
		goto L16
	} else {
		goto L101
	}
L62:
	;
	v352 = v24 + int32(240) + v174<<(uint(int32(2))%32)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	if v358 != 0 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = v344
	goto L62
L64:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v300
	if v295 == int32(0) {
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+1468))
	if v307 != v309 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = int32(0)
	goto L63
L68:
	;
	v314 = F_LWLockAcquire(m, v308+int32(1476), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v326 = int32(base.Ui32(v306) >> (uint(int32(27)) % 32))
	v329 = l0 + v326*int32(20)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	if v330 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v318+int32(1476))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v334 = v330
	goto L76
L75:
	;
	v331 = F_get_segment_by_index(m, l0, v326)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v334+v306&int32(134217727))+8)) = v336
	if v295 == int32(0) {
		goto L62
	} else {
		goto L78
	}
L77:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v334 = v333
	goto L76
L78:
	;
	goto L63
L79:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+1468))
	if v359 != v361 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+30)) = uint16(v174)
	goto L61
L82:
	;
	v363 = int32(0)
	v367 = F_LWLockAcquire(m, v360+int32(1476), v363)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v451 = int32(base.Ui32(v358) >> (uint(int32(27)) % 32))
	v454 = l0 + v451*int32(20)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	if v455 != 0 {
		goto L97
	} else {
		goto L98
	}
L85:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+1468))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v370 != v371 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v375 = v363
	goto L89
L87:
	;
	v427 = v369
	goto L88
L88:
	;
	F_LWLockRelease(m, v427+int32(1476))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L96
	}
L89:
	;
	v391 = v48 + v375*int32(20)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	if v392 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v370
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v427 = v410
	goto L88
L91:
	;
	v406 = v375 + int32(1)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v406) <= base.Ui32(v407) {
		v375 = v406
		goto L89
	} else {
		goto L95
	}
L92:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+24)))
	if v395 != int32(1) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_dsm_detach(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v391))) = int64(0)
	goto L91
L95:
	;
	goto L90
L96:
	;
	goto L84
L97:
	;
	v459 = v455
	goto L99
L98:
	;
	v456 = F_get_segment_by_index(m, l0, v451)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v459+v358&int32(134217727))+4)) = v56
	goto L81
L100:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v459 = v458
	goto L99
L101:
	;
	goto L17
L102:
	;
	goto L15
L103:
	;
	if v514 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v516 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v521 = F_transfer_first_span(m, l0, v26, int32(0), int32(1))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l1 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	if v521 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	v558 = F_FreePageManagerGet(m, v555, v534, v19+int32(12))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L124
	}
L111:
	;
	v809 = l0
	v810 = l1
	v813 = int32(0)
	v819 = v19
	goto L4
L112:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v539 = F_LWLockAcquire(m, v535+int32(1476), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L118
	}
L113:
	;
	v533 = int32(0)
	v534 = int32(1)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v528 = F_alloc_object(m, l0, int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v528 == int32(0) {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v533 = v528
	v534 = int32(16)
	goto L112
L118:
	;
	v541 = F_get_best_segment(m, l0, v534)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v541 != 0 {
		v554 = v541
		goto L110
	} else {
		goto L120
	}
L120:
	;
	v543 = F_make_new_segment(m, l0, v534)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v543 != 0 {
		v554 = v543
		goto L110
	} else {
		goto L122
	}
L122:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v545+int32(1476))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L111
L124:
	;
	if v558 == int32(0) {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v562+int32(1476))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v572 = base.I32_div_s(v554-v48, int32(20))
	v575 = v567<<(uint(int32(12))%32) | v572<<(uint(int32(27))%32)
	if l1 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v576 = v533
	goto L129
L128:
	;
	v576 = v575
	goto L129
L129:
	;
	F_init_span(m, l0, v576, v26, v575, v534, l1)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v579 = int32(0)
	if l1 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v588 = v579
	v597 = int32(0)
	goto L134
L132:
	;
	v641 = v579
	goto L133
L133:
	;
	v657 = v641
	v661 = v579
	goto L138
L134:
	;
	v602 = int32(2)
	v603 = v588 << (uint(v602) % 32)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v603+(v604+v605<<(uint(v602)%32))))) = v576
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v611+v612<<(uint(v602)%32)+v603)+4)) = v576
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v618+v619<<(uint(v602)%32)+v603)+8)) = v576
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v625+v626<<(uint(v602)%32)+v603)+12)) = v576
	v632 = int32(4)
	v633 = v588 + v632
	v635 = v597 + v632
	if v635 != v534&int32(16) {
		v588 = v633
		v597 = v635
		goto L134
	} else {
		goto L136
	}
L135:
	;
	if v534&int32(1) == int32(0) {
		goto L12
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	v641 = v633
	goto L133
L138:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v673 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v671+v672<<(uint(v673)%32)+v657<<(uint(v673)%32)))) = v576
	v680 = int32(1)
	if v661 != 0 {
		v657 = v657 + v680
		v661 = v661 + v680
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L12
L140:
	;
	goto L139
L141:
	;
	v724 = F_LWLockAcquire(m, v718+int32(1476), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v736 = int32(base.Ui32(v703) >> (uint(int32(27)) % 32))
	v739 = v701 + v736*int32(20)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+12))
	if v740 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	F_check_for_freed_segments_locked(m, v701)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	F_LWLockRelease(m, v728+int32(1476))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v743 = F_get_segment_by_index(m, v701, v736)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v746 = v740
	goto L149
L149:
	;
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v702<<(uint(int32(1))%32))+uint32(_c_F_alloc_object[0]))))
	v750 = v746 + v703&int32(134217727)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+12))
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+26)))
	if v752 != int32(_a_F_alloc_object_2) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v739)+12))
	v746 = v745
	goto L149
L151:
	;
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+24)))
	v801 = v799 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v750)+24)) = uint16(v801)
	if v801&int32(_a_F_alloc_object_2) != 0 {
		v809 = v701
		v810 = v702
		v813 = v796
		v819 = v711
		goto L4
	} else {
		goto L165
	}
L152:
	;
	v756 = v752*v749 + v751
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v701)+652))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+1468))
	if v757 != v759 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+22)))
	v790 = v788 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v750)+22)) = uint16(v790)
	v796 = v788*v749 + v751
	goto L151
L155:
	;
	v764 = F_LWLockAcquire(m, v758+int32(1476), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v776 = int32(base.Ui32(v756) >> (uint(int32(27)) % 32))
	v779 = v701 + v776*int32(20)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	if v780 != 0 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	F_check_for_freed_segments_locked(m, v701)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	F_LWLockRelease(m, v768+int32(1476))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	v784 = v780
	goto L163
L162:
	;
	v781 = F_get_segment_by_index(m, v701, v776)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v784+v756&int32(134217727)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v750)+26)) = uint16(v786)
	v796 = v756
	goto L151
L164:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v779)+12))
	v784 = v783
	goto L163
L165:
	;
	v807 = F_transfer_first_span(m, v701, v26, int32(1), int32(3))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v809 = v701
	v810 = v702
	v813 = v796
	v819 = v711
	goto L4
L167:
	;
	m.G0 = v819 + int32(16)
	return v813
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v534
	F_errmsg_internal(m, int32(_a_F_alloc_object_3), v19)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_alloc_object_4), int32(1719), int32(_a_F_alloc_object_5))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_free_object_addresses(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			F_pfree(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1462 int32
	_ = v1462
	var v1475 int32
	_ = v1475
	var v1487 int32
	_ = v1487
	var v1500 int32
	_ = v1500
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1745 int32
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1829 int32
	_ = v1829
	var v1842 int32
	_ = v1842
	var v1859 int32
	_ = v1859
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1909 int32
	_ = v1909
	var v1922 int32
	_ = v1922
	var v1939 int32
	_ = v1939
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int64
	_ = v1974
	var v1976 int32
	_ = v1976
	v7 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(448)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v31 = *(*int64)(unsafe.Add(mBase, _c_F_get_object_address[0]))
	v40 = v7
	v47 = v7
	v49 = v7
	v50 = v7
	v51 = v31
	goto L2
L1:
	;
	m.G0 = v24 + int32(448)
	return
L2:
	;
	switch l1 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 33, 36, 38, 42:
		goto L42
	case 1, 19, 29, 34:
		goto L41
	case 2, 3:
		goto L36
	case 4, 6:
		goto L45
	case 5:
		goto L34
	case 7:
		goto L39
	case 8:
		goto L38
	case 10:
		goto L44
	case 11:
		goto L25
	case 12, 49:
		goto L23
	case 13:
		goto L21
	case 18, 20, 23, 37, 41, 51:
		goto L46
	case 22:
		goto L35
	case 24, 26:
		goto L37
	case 25:
		goto L40
	case 28, 35, 40, 44:
		goto L43
	case 31:
		goto L27
	case 32:
		goto L26
	case 39:
		goto L24
	case 43:
		goto L33
	case 45:
		goto L29
	case 46:
		goto L31
	case 47:
		goto L32
	case 48:
		goto L30
	case 50:
		goto L28
	default:
		v1424 = v40
		goto L22
	}
L3:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L581
	}
L4:
	;
	if v1787 == int32(0) {
		goto L1
	} else {
		goto L508
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1772
	v1778 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1778
	v1787 = v1772
	v1788 = v1773
	v1789 = v1778
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1759
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1761
	v1787 = v1759
	v1788 = v1760
	v1789 = v1761
	goto L4
L7:
	;
	v1759 = v1751
	v1760 = int32(0)
	v1761 = v1753
	goto L6
L8:
	;
	v1666 = base.I32_extend16_s(v921)
	if l1 == int32(2) {
		goto L480
	} else {
		goto L481
	}
L9:
	;
	v1658 = int32(0)
	v1661 = v922
	v1662 = v1658
	v1663 = v1658
	v1665 = v1658
	goto L8
L10:
	;
	v1651 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1651
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1648
	v1655 = int32(826)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1655
	v1787 = v1648
	v1788 = v1651
	v1789 = v1655
	goto L4
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+408)) = v1361
	*(*int32)(unsafe.Add(mBase, uint32(v24)+404)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v1318
	F_errmsg(m, int32(_a_F_get_object_address_0), v24+int32(400))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L47
	} else {
		goto L476
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L47
	} else {
		goto L472
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L47
	} else {
		goto L468
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L47
	} else {
		goto L464
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L47
	} else {
		goto L459
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L47
	} else {
		goto L455
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L47
	} else {
		goto L450
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L47
	} else {
		goto L446
	}
L19:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1787 = v1500
	v1788 = int32(0)
	v1789 = v1487
	goto L4
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1475
	v1487 = v1462
	goto L19
L21:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1441)))
	F_get_object_address_type(m, v24+int32(436), int32(12), v1442, l5)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L47
	} else {
		goto L444
	}
L22:
	;
	if v1424 != 0 {
		v1487 = v1424
		goto L19
	} else {
		goto L440
	}
L23:
	;
	F_get_object_address_type(m, l0, l1, l2, l5)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L47
	} else {
		goto L439
	}
L24:
	;
	v1416 = int32(3381)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1416
	v1419 = F_get_statistics_object_oid(m, l2, l5)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L47
	} else {
		goto L438
	}
L25:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+4))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(3) <= v1312 {
		goto L404
	} else {
		goto L405
	}
L26:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1271)))
	v1273 = F_makeRangeVarFromNameList(m, v1272)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L47
	} else {
		goto L391
	}
L27:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+4))
	v1243 = int32(0)
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	v1247 = F_get_namespace_oid(m, v1246, l5)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L47
	} else {
		goto L383
	}
L28:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1106)))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1111 = int32(_a_F_get_object_address_1)
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_object_address[1])))
	if base.B2i32(v1114 == int32(0))|base.B2i32(v1114 != v1117) != 0 {
		v1135 = v1114
		v1136 = v1117
		goto L338
	} else {
		goto L339
	}
L29:
	;
	v1101 = int32(3602)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1101
	v1104 = F_get_ts_config_oid(m, l2, l5)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L47
	} else {
		goto L334
	}
L30:
	;
	v1096 = int32(3764)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1096
	v1099 = F_get_ts_template_oid(m, l2, l5)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L47
	} else {
		goto L333
	}
L31:
	;
	v1091 = int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1091
	v1094 = F_get_ts_dict_oid(m, l2, l5)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L47
	} else {
		goto L332
	}
L32:
	;
	v1086 = int32(3601)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1086
	v1089 = F_get_ts_parser_oid(m, l2, l5)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L47
	} else {
		goto L331
	}
L33:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+4))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1073)))
	v1078 = F_LookupTypeNameOid(m, v1077, l5)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L47
	} else {
		goto L328
	}
L34:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1031 = F_LookupTypeNameOid(m, v1030, l5)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L47
	} else {
		goto L316
	}
L35:
	;
	v968 = int32(2613)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v968
	v971 = m.G0
	v973 = v971 - int32(16)
	m.G0 = v973
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v975 - int32(465) {
	case 0:
		goto L303
	case 1:
		goto L305
	default:
		goto L304
	}
L36:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v865+v866<<(uint(int32(2))%32)-int32(4))))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)+4))
	v877 = v873
	goto L278
L37:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)+4))
	v835 = F_get_index_am_oid(m, v834)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L47
	} else {
		goto L267
	}
L38:
	;
	v670 = int32(2607)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v670
	v673 = m.G0
	v675 = v673 - int32(16)
	m.G0 = v675
	F_DeconstructQualifiedName(m, l2, v675+int32(12), v675+int32(8))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L47
	} else {
		goto L237
	}
L39:
	;
	v665 = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v665
	v668 = F_get_collation_oid(m, l2, l5)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L47
	} else {
		goto L236
	}
L40:
	;
	v660 = int32(2617)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v660
	v663 = F_LookupOperWithArgs(m, l2, l5)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L47
	} else {
		goto L235
	}
L41:
	;
	v655 = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v655
	v658 = F_LookupFuncWithArgs(m, l1, l2, l5)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L47
	} else {
		goto L234
	}
L42:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	switch l1 {
	case 0:
		goto L197
	default:
		goto L198
	case 9:
		goto L210
	case 14:
		goto L202
	case 15:
		goto L209
	case 16:
		goto L204
	case 17:
		goto L203
	case 21:
		goto L205
	case 27:
		goto L201
	case 30:
		goto L200
	case 33:
		goto L207
	case 36:
		goto L206
	case 38:
		goto L199
	case 42:
		goto L208
	}
L43:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v315 <= int32(1) {
		goto L14
	} else {
		goto L119
	}
L44:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L105
	}
L45:
	;
	if l2 == int32(0) {
		goto L18
	} else {
		goto L94
	}
L46:
	;
	v54 = F_makeRangeVarFromNameList(m, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	v56 = F_relation_openrv_extended(m, v54, l4, l5)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v56 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	switch l1 - int32(18) {
	case 0:
		goto L55
	default:
		goto L54
	case 2:
		goto L60
	case 5:
		goto L56
	case 19:
		goto L59
	case 23:
		goto L58
	case 33:
		goto L57
	}
L51:
	;
	v228 = int32(0)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v1772 = v228
	v1773 = v56
	goto L5
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	v228 = v227
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L47
	} else {
		goto L91
	}
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+119)))
	if v188 == int32(102) {
		goto L53
	} else {
		goto L86
	}
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+119)))
	if v163 == int32(109) {
		goto L53
	} else {
		goto L81
	}
L57:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+119)))
	if v138 == int32(118) {
		goto L53
	} else {
		goto L76
	}
L58:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+119)))
	switch v113 - int32(112) {
	case 0, 2:
		goto L53
	default:
		goto L71
	}
L59:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v88 == int32(83) {
		goto L53
	} else {
		goto L66
	}
L60:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	if v61|int32(32) == int32(105) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L47
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v73 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_2), v24+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L47
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1362), int32(_a_F_get_object_address_4))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L47
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L47
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v98 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_5), v24+int32(48))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1369), int32(_a_F_get_object_address_4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L47
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L47
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L47
	} else {
		goto L73
	}
L73:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v123 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_6), v24-int32(-64))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L47
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1377), int32(_a_F_get_object_address_4))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L47
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L47
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L47
	} else {
		goto L78
	}
L78:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v148 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_7), v24+int32(80))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L47
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1384), int32(_a_F_get_object_address_4))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L47
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L47
	} else {
		goto L83
	}
L83:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v173 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_8), v24+int32(96))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L47
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1391), int32(_a_F_get_object_address_4))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L47
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L47
	} else {
		goto L88
	}
L88:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v198 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_9), v24+int32(112))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L47
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1398), int32(_a_F_get_object_address_4))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L47
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_10), v24+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L47
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1401), int32(_a_F_get_object_address_4))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L47
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v233 <= int32(1) {
		goto L18
	} else {
		goto L95
	}
L95:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236+v233<<(uint(int32(2))%32)-int32(4))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v246 = F_list_copy_head(m, l2, v233-int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L47
	} else {
		goto L96
	}
L96:
	;
	v248 = F_makeRangeVarFromNameList(m, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L47
	} else {
		goto L97
	}
L97:
	;
	v250 = F_relation_openrv(m, v248, l4)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L47
	} else {
		goto L98
	}
L98:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	v253 = F_get_attnum(m, v252, v243)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L47
	} else {
		goto L99
	}
L99:
	;
	if v253 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if l5 == int32(0) {
		goto L17
	} else {
		goto L103
	}
L101:
	;
	v264 = v252
	v265 = v250
	v266 = v253
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v266
	v1772 = v264
	v1773 = v265
	goto L5
L103:
	;
	F_relation_close(m, v250, l4)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L47
	} else {
		goto L104
	}
L104:
	;
	v261 = int32(0)
	v264 = v261
	v265 = v261
	v266 = v261
	goto L102
L105:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v270 <= int32(1) {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+v270<<(uint(int32(2))%32)-int32(4))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v283 = F_list_copy_head(m, l2, v270-int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L47
	} else {
		goto L107
	}
L107:
	;
	v285 = F_makeRangeVarFromNameList(m, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L47
	} else {
		goto L108
	}
L108:
	;
	v287 = F_relation_openrv(m, v285, l4)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L47
	} else {
		goto L109
	}
L109:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v287)+52))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+56))
	v291 = F_get_attnum(m, v290, v280)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L47
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v307
	v312 = int32(2604)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v312
	v1787 = v307
	v1788 = v308
	v1789 = v312
	goto L4
L111:
	;
	if l5 == int32(0) {
		goto L15
	} else {
		goto L117
	}
L112:
	;
	if v291 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	if v295 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v298 = F_GetAttrDefaultOid(m, v290, v291)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L47
	} else {
		goto L115
	}
L115:
	;
	if v298 != 0 {
		v307 = v298
		v308 = v287
		goto L110
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	F_relation_close(m, v287, l4)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L47
	} else {
		goto L118
	}
L118:
	;
	v305 = int32(0)
	v307 = v305
	v308 = v305
	goto L110
L119:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+v315<<(uint(int32(2))%32)-int32(4))))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v328 = F_list_copy_head(m, l2, v315-int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L47
	} else {
		goto L120
	}
L120:
	;
	v330 = F_makeRangeVarFromNameList(m, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L47
	} else {
		goto L121
	}
L121:
	;
	v333 = F_table_openrv_extended(m, v330, int32(1), l5)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L47
	} else {
		goto L122
	}
L122:
	;
	if v333 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)+56))
	v337 = v335
	goto L125
L124:
	;
	v337 = int32(0)
	goto L125
L125:
	;
	switch l1 - int32(28) {
	case 0:
		goto L131
	default:
		goto L130
	case 7:
		goto L134
	case 12:
		goto L132
	case 16:
		goto L133
	}
L126:
	;
	v1759 = v561
	v1760 = v562
	v1761 = v563
	goto L6
L127:
	;
	v559 = int32(0)
	v561 = v559
	v562 = v559
	v563 = v553
	goto L126
L128:
	;
	if v548 != 0 {
		v561 = v548
		v562 = v333
		v563 = v542
		goto L126
	} else {
		goto L195
	}
L129:
	;
	v502 = m.G0
	v504 = v502 - int32(16)
	m.G0 = v504
	v507 = F_SearchSysCache2(m, int32(60), v337, v325)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L47
	} else {
		goto L184
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L47
	} else {
		goto L180
	}
L131:
	;
	if v333 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L132:
	;
	if v333 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L133:
	;
	if v333 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	if v333 != 0 {
		goto L129
	} else {
		goto L135
	}
L135:
	;
	v553 = int32(2618)
	goto L127
L136:
	;
	v553 = int32(2620)
	goto L127
L137:
	;
	goto L138
L138:
	;
	v344 = int32(2620)
	v346 = m.G0
	v348 = v346 - int32(112)
	m.G0 = v348
	v352 = F_table_open(m, v344, int32(1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L47
	} else {
		goto L139
	}
L139:
	;
	v355 = v348 + int32(16)
	F_ScanKeyInit(m, v355, int32(2), int32(3), int32(184), v337)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L47
	} else {
		goto L140
	}
L140:
	;
	F_ScanKeyInit(m, v348-int32(-64), int32(4), int32(3), int32(62), v325)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L47
	} else {
		goto L141
	}
L141:
	;
	v372 = F_systable_beginscan(m, v352, int32(2701), int32(1), int32(0), int32(2), v355)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L47
	} else {
		goto L143
	}
L142:
	;
	F_systable_endscan(m, v372)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L47
	} else {
		goto L154
	}
L143:
	;
	v374 = F_systable_getnext(m, v372)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L47
	} else {
		goto L144
	}
L144:
	;
	if v374 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if l5 != 0 {
		v401 = int32(0)
		goto L142
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v374)+16))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+22)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397+v398)))
	v401 = v400
	goto L142
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L47
	} else {
		goto L149
	}
L149:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L47
	} else {
		goto L150
	}
L150:
	;
	v385 = F_get_rel_name(m, v337)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L47
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348)+4)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v325
	F_errmsg(m, int32(_a_F_get_object_address_11), v348)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L47
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_get_object_address_12), int32(1404), int32(_a_F_get_object_address_13))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L47
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_relation_close(m, v352, int32(1))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L47
	} else {
		goto L155
	}
L155:
	;
	m.G0 = v348 + int32(112)
	v542 = v344
	v548 = v401
	goto L128
L156:
	;
	v553 = int32(2606)
	goto L127
L157:
	;
	goto L158
L158:
	;
	v414 = F_get_relation_constraint_oid(m, v337, v325, l5)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L47
	} else {
		goto L159
	}
L159:
	;
	v542 = int32(2606)
	v548 = v414
	goto L128
L160:
	;
	v553 = int32(3256)
	goto L127
L161:
	;
	goto L162
L162:
	;
	v419 = int32(3256)
	v421 = m.G0
	v423 = v421 - int32(112)
	m.G0 = v423
	v427 = F_table_open(m, v419, int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L47
	} else {
		goto L163
	}
L163:
	;
	v430 = v423 + int32(16)
	v431 = int32(3)
	F_ScanKeyInit(m, v430, v431, v431, int32(184), v337)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L47
	} else {
		goto L164
	}
L164:
	;
	F_ScanKeyInit(m, v423-int32(-64), int32(2), int32(3), int32(62), v325)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L47
	} else {
		goto L165
	}
L165:
	;
	v447 = F_systable_beginscan(m, v427, int32(3258), int32(1), int32(0), int32(2), v430)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L47
	} else {
		goto L167
	}
L166:
	;
	F_systable_endscan(m, v447)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L47
	} else {
		goto L178
	}
L167:
	;
	v449 = F_systable_getnext(m, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L47
	} else {
		goto L168
	}
L168:
	;
	if v449 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if l5 != 0 {
		v476 = int32(0)
		goto L166
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v449)+16))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+22)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v472+v473)))
	v476 = v475
	goto L166
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L47
	} else {
		goto L173
	}
L173:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L47
	} else {
		goto L174
	}
L174:
	;
	v460 = F_get_rel_name(m, v337)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L47
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v325
	F_errmsg(m, int32(_a_F_get_object_address_14), v423)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L47
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_get_object_address_15), int32(1238), int32(_a_F_get_object_address_16))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L47
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_relation_close(m, v427, int32(1))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L47
	} else {
		goto L179
	}
L179:
	;
	m.G0 = v423 + int32(112)
	v542 = v419
	v548 = v476
	goto L128
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_10), v24+int32(160))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L47
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1477), int32(_a_F_get_object_address_17))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L47
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	m.G0 = v504 + int32(16)
	v542 = int32(2618)
	v548 = v537
	goto L128
L184:
	;
	if v507 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	if l5 != 0 {
		v537 = int32(0)
		goto L183
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+22)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v530+v531)))
	F_ReleaseCatCache(m, v507)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L47
	} else {
		goto L194
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L47
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L47
	} else {
		goto L190
	}
L190:
	;
	v518 = F_get_rel_name(m, v337)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L47
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+4)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v325
	F_errmsg(m, int32(_a_F_get_object_address_18), v504)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L47
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_get_object_address_19), int32(109), int32(_a_F_get_object_address_20))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L47
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
	v537 = v533
	goto L183
L195:
	;
	F_relation_close(m, v333, int32(1))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L47
	} else {
		goto L196
	}
L196:
	;
	v553 = v542
	goto L127
L197:
	;
	v653 = F_get_am_type_oid(m, v569, int32(0), l5)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L47
	} else {
		goto L233
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L47
	} else {
		goto L230
	}
L199:
	;
	v634 = F_get_subscription_oid(m, v569, l5)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L47
	} else {
		goto L229
	}
L200:
	;
	v631 = F_get_publication_oid(m, v569, l5)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L47
	} else {
		goto L228
	}
L201:
	;
	v628 = F_ParameterAclLookup(m, v569, l5)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L47
	} else {
		goto L227
	}
L202:
	;
	v595 = m.G0
	v597 = v595 - int32(16)
	m.G0 = v597
	v600 = int32(0)
	v603 = F_GetSysCacheOid(m, int32(25), v569, v600, v600, v600)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L47
	} else {
		goto L219
	}
L203:
	;
	v592 = F_get_foreign_server_oid(m, v569, l5)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L47
	} else {
		goto L218
	}
L204:
	;
	v589 = F_get_foreign_data_wrapper_oid(m, v569, l5)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L47
	} else {
		goto L217
	}
L205:
	;
	v586 = F_get_language_oid(m, v569, l5)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L47
	} else {
		goto L216
	}
L206:
	;
	v583 = F_get_namespace_oid(m, v569, l5)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L47
	} else {
		goto L215
	}
L207:
	;
	v580 = F_get_role_oid(m, v569, l5)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L47
	} else {
		goto L214
	}
L208:
	;
	v577 = F_get_tablespace_oid(m, v569, l5)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L47
	} else {
		goto L213
	}
L209:
	;
	v574 = F_get_extension_oid(m, v569, l5)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L47
	} else {
		goto L212
	}
L210:
	;
	v571 = F_get_database_oid(m, v569, l5)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L47
	} else {
		goto L211
	}
L211:
	;
	v1751 = v571
	v1753 = int32(1262)
	goto L7
L212:
	;
	v1751 = v574
	v1753 = int32(3079)
	goto L7
L213:
	;
	v1751 = v577
	v1753 = int32(1213)
	goto L7
L214:
	;
	v1751 = v580
	v1753 = int32(1260)
	goto L7
L215:
	;
	v1751 = v583
	v1753 = int32(2615)
	goto L7
L216:
	;
	v1751 = v586
	v1753 = int32(2612)
	goto L7
L217:
	;
	v1751 = v589
	v1753 = int32(2328)
	goto L7
L218:
	;
	v1751 = v592
	v1753 = int32(1417)
	goto L7
L219:
	;
	if v603|l5 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L47
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	m.G0 = v597 + int32(16)
	v1751 = v603
	v1753 = int32(3466)
	goto L7
L223:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L47
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = v569
	F_errmsg(m, int32(_a_F_get_object_address_21), v597)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L47
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_get_object_address_22), int32(588), int32(_a_F_get_object_address_23))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L47
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
	v1751 = v628
	v1753 = int32(_a_F_get_object_address_24)
	goto L7
L228:
	;
	v1751 = v631
	v1753 = int32(_a_F_get_object_address_25)
	goto L7
L229:
	;
	v1751 = v634
	v1753 = int32(_a_F_get_object_address_26)
	goto L7
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_10), v24+int32(176))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L47
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1324), int32(_a_F_get_object_address_27))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L47
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	v1751 = v653
	v1753 = int32(2601)
	goto L7
L234:
	;
	v1462 = v655
	v1475 = v658
	goto L20
L235:
	;
	v1462 = v660
	v1475 = v663
	goto L20
L236:
	;
	v1462 = v665
	v1475 = v668
	goto L20
L237:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v675)+12))
	if v683 != 0 {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	m.G0 = v675 + int32(16)
	v1462 = v670
	v1475 = v815
	goto L20
L239:
	;
	if v788|l5 != 0 {
		v815 = v788
		goto L238
	} else {
		goto L261
	}
L240:
	;
	v788 = int32(0)
	goto L239
L241:
	;
	v685 = F_LookupExplicitNamespace(m, v683, l5)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L47
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L47
	} else {
		goto L250
	}
L244:
	;
	if v685 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v687 = int32(0)
	goto L247
L246:
	;
	v687 = l5
	goto L247
L247:
	;
	if v687 != 0 {
		goto L240
	} else {
		goto L248
	}
L248:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v675)+8))
	v690 = int32(0)
	v692 = F_GetSysCacheOid(m, int32(18), v689, v685, v690, v690)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L47
	} else {
		goto L249
	}
L249:
	;
	v788 = v692
	goto L239
L250:
	;
	v696 = int32(0)
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_address[2]))
	if v698 == v696 {
		v788 = v696
		goto L239
	} else {
		goto L251
	}
L251:
	;
	v701 = int32(0)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v702 <= v701 {
		goto L240
	} else {
		goto L252
	}
L252:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v675)+8))
	v712 = v701
	goto L253
L253:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v698)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v712<<(uint(int32(2))%32))))
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_address[3]))
	if v731 != v733 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	goto L240
L255:
	;
	v736 = int32(0)
	v738 = F_GetSysCacheOid(m, int32(18), v705, v731, v736, v736)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L47
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v742 = v712 + int32(1)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	if v742 < v743 {
		v712 = v742
		goto L253
	} else {
		goto L260
	}
L258:
	;
	if v738 != 0 {
		v815 = v738
		goto L238
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	goto L254
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L47
	} else {
		goto L262
	}
L262:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L47
	} else {
		goto L263
	}
L263:
	;
	v797 = F_NameListToString(m, l2)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L47
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v675))) = v797
	F_errmsg(m, int32(_a_F_get_object_address_28), v675)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L47
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_get_object_address_29), int32(4088), int32(_a_F_get_object_address_30))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L47
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
	v838 = F_list_copy_tail(m, l2, int32(1))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L47
	} else {
		goto L268
	}
L268:
	;
	switch l1 - int32(24) {
	case 0:
		goto L269
	default:
		goto L270
	case 2:
		goto L271
	}
L269:
	;
	v861 = F_get_opclass_oid(m, v835, v838, l5)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L47
	} else {
		goto L276
	}
L270:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L47
	} else {
		goto L273
	}
L271:
	;
	v843 = F_get_opfamily_oid(m, v835, v838, l5)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L47
	} else {
		goto L272
	}
L272:
	;
	v1751 = v843
	v1753 = int32(2753)
	goto L7
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_10), v24+int32(192))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L47
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1669), int32(_a_F_get_object_address_31))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L47
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	v1751 = v861
	v1753 = int32(2616)
	goto L7
L277:
	;
	v922 = int32(0)
	v925 = F_list_copy_head(m, v864, v866-int32(1))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L47
	} else {
		goto L293
	}
L278:
	;
	v882 = v877 + int32(1)
	v883 = int32(*(*int8)(unsafe.Add(mBase, uint32(v877))))
	v884 = F___isspace(m, v883)
	mBase = m.M
	if v884 != 0 {
		v877 = v882
		goto L278
	} else {
		goto L280
	}
L279:
	;
	v885 = int32(1)
	switch v883&int32(255) - int32(43) {
	case 0:
		v891 = v885
		goto L282
	default:
		v893 = v883
		v894 = v877
		v895 = v885
		goto L281
	case 2:
		goto L283
	}
L280:
	;
	goto L279
L281:
	;
	v896 = int32(0)
	v898 = v893 - int32(48)
	if base.Ui32(v898) <= base.Ui32(int32(9)) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v892 = int32(*(*int8)(unsafe.Add(mBase, uint32(v882))))
	v893 = v892
	v894 = v882
	v895 = v891
	goto L281
L283:
	;
	v891 = int32(0)
	goto L282
L284:
	;
	v901 = v896
	v902 = v898
	v903 = v894
	goto L287
L285:
	;
	v915 = v896
	goto L286
L286:
	;
	if v895 != 0 {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	v905 = int32(10)
	v907 = v901*v905 - v902
	v908 = int32(*(*int8)(unsafe.Add(mBase, uint32(v903)+1)))
	v912 = v908 - int32(48)
	if base.Ui32(v912) < base.Ui32(v905) {
		v901 = v907
		v902 = v912
		v903 = v903 + int32(1)
		goto L287
	} else {
		goto L289
	}
L288:
	;
	v915 = v907
	goto L286
L289:
	;
	goto L288
L290:
	;
	v921 = int32(0) - v915
	goto L292
L291:
	;
	v921 = v915
	goto L292
L292:
	;
	goto L277
L293:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v925)+12))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v928)+4))
	v930 = F_get_index_am_oid(m, v929)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L47
	} else {
		goto L294
	}
L294:
	;
	v933 = F_list_copy_tail(m, v925, int32(1))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L47
	} else {
		goto L295
	}
L295:
	;
	v936 = F_get_opfamily_oid(m, v930, v933, int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L47
	} else {
		goto L296
	}
L296:
	;
	v938 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+444)) = v938
	*(*int32)(unsafe.Add(mBase, uint32(v24)+440)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v24)+436)) = int32(2753)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	if v944 == v938 {
		goto L9
	} else {
		goto L297
	}
L297:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	if v947 <= int32(0) {
		goto L9
	} else {
		goto L298
	}
L298:
	;
	v951 = v24 + int32(424)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v944)+12))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v953)))
	F_get_object_address_type(m, v951, int32(49), v954, l5)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L47
	} else {
		goto L299
	}
L299:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v24)+428))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	if v959 < int32(2) {
		v1661 = v922
		v1662 = v957
		v1663 = v954
		v1665 = int32(0)
		goto L8
	} else {
		goto L300
	}
L300:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v944)+12))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)+4))
	F_get_object_address_type(m, v951, int32(49), v964, l5)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L47
	} else {
		goto L301
	}
L301:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v24)+428))
	v1661 = v964
	v1662 = v957
	v1663 = v954
	v1665 = v967
	goto L8
L302:
	;
	m.G0 = v973 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v999
	v1006 = F_LargeObjectExists(m, v999)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L47
	} else {
		goto L310
	}
L303:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v999 = v998
	goto L302
L304:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L47
	} else {
		goto L307
	}
L305:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v979 = int32(0)
	v982 = F_uint32in_subr(m, v978, v979, int32(_a_F_get_object_address_32), v979)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L47
	} else {
		goto L306
	}
L306:
	;
	v999 = v982
	goto L302
L307:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v973))) = v988
	F_errmsg_internal(m, int32(_a_F_get_object_address_33), v973)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L47
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_get_object_address_34), int32(280), int32(_a_F_get_object_address_35))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L47
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	if v1006|l5 != 0 {
		v1487 = v968
		goto L19
	} else {
		goto L311
	}
L311:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L47
	} else {
		goto L312
	}
L312:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L47
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v999
	F_errmsg(m, int32(_a_F_get_object_address_36), v24+int32(240))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L47
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1056), int32(_a_F_get_object_address_37))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L47
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	v1033 = F_LookupTypeNameOid(m, v1028, l5)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L47
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2605)
	v1037 = m.G0
	v1039 = v1037 - int32(16)
	m.G0 = v1039
	v1042 = int32(0)
	v1044 = F_GetSysCacheOid(m, int32(12), v1031, v1033, v1042, v1042)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L47
	} else {
		goto L318
	}
L318:
	;
	if v1044|l5 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L47
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	m.G0 = v1039 + int32(16)
	v1462 = int32(2605)
	v1475 = v1044
	goto L20
L322:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L47
	} else {
		goto L323
	}
L323:
	;
	v1056 = F_format_type_be(m, v1031)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L47
	} else {
		goto L324
	}
L324:
	;
	v1058 = F_format_type_be(m, v1033)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L47
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1039)+4)) = v1058
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1056
	F_errmsg(m, int32(_a_F_get_object_address_38), v1039)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L47
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_get_object_address_39), int32(1111), int32(_a_F_get_object_address_40))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L47
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	v1080 = F_get_language_oid(m, v1075, l5)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L47
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3576)
	v1084 = F_get_transform_oid(m, v1078, v1080, l5)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L47
	} else {
		goto L330
	}
L330:
	;
	v1462 = int32(3576)
	v1475 = v1084
	goto L20
L331:
	;
	v1462 = v1086
	v1475 = v1089
	goto L20
L332:
	;
	v1462 = v1091
	v1475 = v1094
	goto L20
L333:
	;
	v1462 = v1096
	v1475 = v1099
	goto L20
L334:
	;
	v1462 = v1101
	v1475 = v1104
	goto L20
L335:
	;
	v1233 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1233
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1230
	v1237 = int32(1418)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1237
	v1787 = v1230
	v1788 = v1233
	v1789 = v1237
	goto L4
L336:
	;
	v1175 = F_GetForeignServerByName(m, v1108, int32(1))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L47
	} else {
		goto L359
	}
L337:
	;
	if v1135-v1136 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L338:
	;
	goto L337
L339:
	;
	v1120 = v1110
	v1121 = v1111
	goto L340
L340:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121)+1)))
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120)+1)))
	if v1125 == int32(0) {
		v1135 = v1125
		v1136 = v1124
		goto L338
	} else {
		goto L342
	}
L341:
	;
	v1135 = v1125
	v1136 = v1124
	goto L338
L342:
	;
	v1128 = int32(1)
	if v1125 == v1124 {
		v1120 = v1120 + v1128
		v1121 = v1121 + v1128
		goto L340
	} else {
		goto L343
	}
L343:
	;
	goto L341
L344:
	;
	v1172 = int32(0)
	goto L336
L345:
	;
	goto L346
L346:
	;
	v1142 = F_SearchSysCache1(m, int32(10), v1110)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L47
	} else {
		goto L347
	}
L347:
	;
	if v1142 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if l5 != 0 {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	goto L350
L350:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+16))
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+22)))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1166+v1167)))
	F_ReleaseCatCache(m, v1142)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L47
	} else {
		goto L358
	}
L351:
	;
	v1230 = int32(0)
	goto L335
L352:
	;
	goto L353
L353:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L47
	} else {
		goto L354
	}
L354:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L47
	} else {
		goto L355
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = v1110
	F_errmsg(m, int32(_a_F_get_object_address_41), v24+int32(288))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L47
	} else {
		goto L356
	}
L356:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1825), int32(_a_F_get_object_address_42))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L47
	} else {
		goto L357
	}
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	v1172 = v1169
	goto L336
L359:
	;
	if v1175 == int32(0) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	if l5 != 0 {
		goto L363
	} else {
		goto L364
	}
L361:
	;
	goto L362
L362:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	v1200 = F_SearchSysCache2(m, int32(84), v1172, v1199)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L47
	} else {
		goto L370
	}
L363:
	;
	v1230 = int32(0)
	goto L335
L364:
	;
	goto L365
L365:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L47
	} else {
		goto L366
	}
L366:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L47
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = v1108
	F_errmsg(m, int32(_a_F_get_object_address_43), v24+int32(256))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L47
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1839), int32(_a_F_get_object_address_42))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L47
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	if v1200 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	if l5 != 0 {
		goto L374
	} else {
		goto L375
	}
L372:
	;
	goto L373
L373:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+16))
	v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1224)+22)))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1224+v1225)))
	F_ReleaseCatCache(m, v1200)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L47
	} else {
		goto L381
	}
L374:
	;
	v1230 = int32(0)
	goto L335
L375:
	;
	goto L376
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L47
	} else {
		goto L377
	}
L377:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L47
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+276)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v1110
	F_errmsg(m, int32(_a_F_get_object_address_41), v24+int32(272))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L47
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1851), int32(_a_F_get_object_address_42))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L47
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	v1230 = v1227
	goto L335
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1263
	v1268 = int32(_a_F_get_object_address_44)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1268
	v1787 = v1263
	v1788 = v1243
	v1789 = v1268
	goto L4
L383:
	;
	if v1247 == int32(0) {
		v1263 = v1243
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1251 = F_GetPublicationByName(m, v1242, l5)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L47
	} else {
		goto L385
	}
L385:
	;
	if v1251 == int32(0) {
		v1263 = v1243
		goto L382
	} else {
		goto L386
	}
L386:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1251)))
	v1257 = int32(0)
	v1259 = F_GetSysCacheOid(m, int32(50), v1247, v1256, v1257, v1257)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L47
	} else {
		goto L387
	}
L387:
	;
	if l5 != 0 {
		v1263 = v1259
		goto L382
	} else {
		goto L388
	}
L388:
	;
	if v1259 == int32(0) {
		goto L13
	} else {
		goto L389
	}
L389:
	;
	v1263 = v1259
	goto L382
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1300
	v1306 = int32(_a_F_get_object_address_45)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1306
	v1787 = v1300
	v1788 = v1301
	v1789 = v1306
	goto L4
L391:
	;
	v1276 = F_relation_openrv_extended(m, v1273, int32(1), l5)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L47
	} else {
		goto L392
	}
L392:
	;
	if v1276 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+4))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+4))
	v1281 = F_GetPublicationByName(m, v1280, l5)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L47
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1298 = int32(0)
	v1300 = v1298
	v1301 = v1298
	goto L390
L396:
	;
	if v1281 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+56))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1281)))
	v1286 = int32(0)
	v1288 = F_GetSysCacheOid(m, int32(53), v1284, v1285, v1286, v1286)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L47
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	F_relation_close(m, v1276, int32(1))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L47
	} else {
		goto L403
	}
L400:
	;
	if v1288 != 0 {
		v1300 = v1288
		v1301 = v1276
		goto L390
	} else {
		goto L401
	}
L401:
	;
	if l5 == int32(0) {
		goto L12
	} else {
		goto L402
	}
L402:
	;
	goto L399
L403:
	;
	goto L395
L404:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+8))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	v1317 = v1316
	goto L406
L405:
	;
	v1317 = int32(0)
	goto L406
L406:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+4))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+4))
	v1322 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1321))))
	v1324 = v1322 & int32(255)
	switch v1324 - int32(76) {
	case 0:
		goto L410
	default:
		goto L409
	case 7:
		goto L408
	case 8:
		goto L412
	case 26:
		goto L413
	case 34:
		goto L411
	case 38:
		v1361 = int32(_a_F_get_object_address_46)
		goto L407
	}
L407:
	;
	v1363 = F_SearchSysCache1(m, int32(10), v1318)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L47
	} else {
		goto L420
	}
L408:
	;
	v1361 = int32(_a_F_get_object_address_47)
	goto L407
L409:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L47
	} else {
		goto L414
	}
L410:
	;
	v1361 = int32(_a_F_get_object_address_48)
	goto L407
L411:
	;
	v1361 = int32(_a_F_get_object_address_49)
	goto L407
L412:
	;
	v1361 = int32(_a_F_get_object_address_50)
	goto L407
L413:
	;
	v1361 = int32(_a_F_get_object_address_51)
	goto L407
L414:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L47
	} else {
		goto L415
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v1322
	F_errmsg(m, int32(_a_F_get_object_address_52), v24+int32(368))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L47
	} else {
		goto L416
	}
L416:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+352)) = int64(326417514606)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+344)) = int64(360777252966)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+336)) = int64(356482285682)
	F_errhint(m, int32(_a_F_get_object_address_53), v24+int32(336))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L47
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(2021), int32(_a_F_get_object_address_54))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L47
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	if l5 != 0 {
		v1648 = int32(0)
		goto L10
	} else {
		goto L432
	}
L420:
	;
	if v1363 == int32(0) {
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1363)+16))
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+22)))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1367+v1368)))
	F_ReleaseCatCache(m, v1363)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L47
	} else {
		goto L422
	}
L422:
	;
	if v1317 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	v1383 = F_SearchSysCache3(m, int32(22), v1370, v1381, v1324)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L47
	} else {
		goto L429
	}
L424:
	;
	v1381 = int32(0)
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1377 = F_get_namespace_oid(m, v1317, int32(1))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L47
	} else {
		goto L427
	}
L427:
	;
	if v1377 == int32(0) {
		goto L419
	} else {
		goto L428
	}
L428:
	;
	v1381 = v1377
	goto L423
L429:
	;
	if v1383 == int32(0) {
		goto L419
	} else {
		goto L430
	}
L430:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+16))
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387)+22)))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1387+v1388)))
	F_ReleaseCatCache(m, v1383)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L47
	} else {
		goto L431
	}
L431:
	;
	v1648 = v1390
	goto L10
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L47
	} else {
		goto L433
	}
L433:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L47
	} else {
		goto L434
	}
L434:
	;
	if v1317 != 0 {
		goto L11
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v1361
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = v1318
	F_errmsg(m, int32(_a_F_get_object_address_55), v24+int32(384))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L47
	} else {
		goto L436
	}
L436:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(2073), int32(_a_F_get_object_address_54))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L47
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	v1462 = v1416
	v1475 = v1419
	goto L20
L439:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1424 = v1423
	goto L22
L440:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L47
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_10), v24)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L47
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1134), int32(_a_F_get_object_address_37))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L47
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v24)+440))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+4))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+4))
	v1449 = int32(2606)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1449
	v1452 = F_get_domain_constraint_oid(m, v1445, v1448, l5)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L47
	} else {
		goto L445
	}
L445:
	;
	v1462 = v1449
	v1475 = v1452
	goto L20
L446:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L47
	} else {
		goto L447
	}
L447:
	;
	F_errmsg(m, int32(_a_F_get_object_address_56), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L47
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1514), int32(_a_F_get_object_address_57))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L47
	} else {
		goto L449
	}
L449:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L450:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L47
	} else {
		goto L451
	}
L451:
	;
	v1526 = F_NameListToString(m, v246)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L47
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v243
	F_errmsg(m, int32(_a_F_get_object_address_58), v24+int32(128))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L47
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1529), int32(_a_F_get_object_address_57))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L47
	} else {
		goto L454
	}
L454:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L455:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L47
	} else {
		goto L456
	}
L456:
	;
	F_errmsg(m, int32(_a_F_get_object_address_56), int32(0))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L47
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1567), int32(_a_F_get_object_address_59))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L47
	} else {
		goto L458
	}
L458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L459:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L47
	} else {
		goto L460
	}
L460:
	;
	v1564 = F_NameListToString(m, v283)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L47
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v1564
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v280
	F_errmsg(m, int32(_a_F_get_object_address_60), v24+int32(144))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L47
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1587), int32(_a_F_get_object_address_59))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L47
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L47
	} else {
		goto L465
	}
L465:
	;
	F_errmsg(m, int32(_a_F_get_object_address_61), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L47
	} else {
		goto L466
	}
L466:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1438), int32(_a_F_get_object_address_17))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L47
	} else {
		goto L467
	}
L467:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L468:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L47
	} else {
		goto L469
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+308)) = v1242
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v1246
	F_errmsg(m, int32(_a_F_get_object_address_62), v24+int32(304))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L47
	} else {
		goto L470
	}
L470:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1954), int32(_a_F_get_object_address_63))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L47
	} else {
		goto L471
	}
L471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L472:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L47
	} else {
		goto L473
	}
L473:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+324)) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v1620 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_64), v24+int32(320))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L47
	} else {
		goto L474
	}
L474:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1907), int32(_a_F_get_object_address_65))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L47
	} else {
		goto L475
	}
L475:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L476:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(2068), int32(_a_F_get_object_address_54))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L47
	} else {
		goto L477
	}
L477:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L478:
	;
	if l5 != 0 {
		goto L498
	} else {
		goto L499
	}
L479:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+16))
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713)+22)))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1713+v1714)))
	F_ReleaseCatCache(m, v1711)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L47
	} else {
		goto L497
	}
L480:
	;
	v1671 = F_SearchSysCache4(m, int32(4), v936, v1662, v1665, v1666)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L47
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	v1706 = F_SearchSysCache4(m, int32(5), v936, v1662, v1665, v1666)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L47
	} else {
		goto L495
	}
L483:
	;
	if v1671 != 0 {
		v1711 = v1671
		v1712 = int32(2602)
		goto L479
	} else {
		goto L484
	}
L484:
	;
	if l5 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v1751 = int32(0)
	v1753 = int32(2602)
	goto L7
L486:
	;
	goto L487
L487:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L47
	} else {
		goto L488
	}
L488:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L47
	} else {
		goto L489
	}
L489:
	;
	v1682 = F_TypeNameToString(m, v1663)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L47
	} else {
		goto L490
	}
L490:
	;
	v1684 = F_TypeNameToString(m, v1661)
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L47
	} else {
		goto L491
	}
L491:
	;
	v1689 = F_getObjectDescription(m, v24+int32(436), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L47
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+220)) = v1689
	*(*int32)(unsafe.Add(mBase, uint32(v24)+216)) = v1684
	*(*int32)(unsafe.Add(mBase, uint32(v24)+212)) = v1682
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v921
	F_errmsg(m, int32(_a_F_get_object_address_66), v24+int32(208))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L47
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1746), int32(_a_F_get_object_address_67))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L47
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	if v1706 == int32(0) {
		goto L478
	} else {
		goto L496
	}
L496:
	;
	v1711 = v1706
	v1712 = int32(2603)
	goto L479
L497:
	;
	v1751 = v1716
	v1753 = v1712
	goto L7
L498:
	;
	v1751 = int32(0)
	v1753 = int32(2603)
	goto L7
L499:
	;
	goto L500
L500:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L47
	} else {
		goto L501
	}
L501:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L47
	} else {
		goto L502
	}
L502:
	;
	v1728 = F_TypeNameToString(m, v1663)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L47
	} else {
		goto L503
	}
L503:
	;
	v1730 = F_TypeNameToString(m, v1661)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L47
	} else {
		goto L504
	}
L504:
	;
	v1735 = F_getObjectDescription(m, v24+int32(436), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L47
	} else {
		goto L505
	}
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+236)) = v1735
	*(*int32)(unsafe.Add(mBase, uint32(v24)+232)) = v1730
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v1728
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = v921
	F_errmsg(m, int32(_a_F_get_object_address_68), v24+int32(224))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L47
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(_a_F_get_object_address_3), int32(1777), int32(_a_F_get_object_address_67))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L47
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L508:
	;
	if v47 == int32(0) {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	goto L3
L510:
	;
	if v1789 == int32(1259) {
		goto L547
	} else {
		goto L548
	}
L511:
	;
	if base.B2i32(v1789 != v47)|base.B2i32(v1787 != v49) == int32(0) {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v50 == v1811 {
		goto L509
	} else {
		goto L515
	}
L513:
	;
	goto L514
L514:
	;
	if v47 == int32(1259) {
		goto L510
	} else {
		goto L516
	}
L515:
	;
	goto L514
L516:
	;
	v1817 = int32(1)
	if v47 <= int32(3591) {
		goto L521
	} else {
		goto L522
	}
L517:
	;
	if v1888 != 0 {
		goto L542
	} else {
		goto L543
	}
L518:
	;
	goto L517
L519:
	;
	v1888 = int32(0)
	goto L518
L520:
	;
	if base.B2i32(base.Ui32(v47-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v47-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v1888 = v1817
		goto L518
	} else {
		goto L541
	}
L521:
	;
	if v47 <= int32(2670) {
		goto L524
	} else {
		goto L525
	}
L522:
	;
	goto L523
L523:
	;
	if v47 <= int32(_a_F_get_object_address_69) {
		goto L531
	} else {
		goto L532
	}
L524:
	;
	switch v47 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v1888 = v1817
		goto L518
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L519
	default:
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v1829 = v47 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v1829))|base.B2i32(int32(1)<<(uint(v1829)%32)&int32(226492515) == int32(0)) != 0 {
		goto L520
	} else {
		goto L529
	}
L527:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v47-int32(2396)) {
		goto L519
	} else {
		goto L528
	}
L528:
	;
	v1888 = v1817
	goto L518
L529:
	;
	v1888 = v1817
	goto L518
L530:
	;
	if base.Ui32(v47-int32(3592)) < base.Ui32(int32(2)) {
		v1888 = v1817
		goto L518
	} else {
		goto L539
	}
L531:
	;
	v1842 = v47 - int32(_a_F_get_object_address_70)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v1842))|base.B2i32(int32(1)<<(uint(v1842)%32)&int32(963) == int32(0)) != 0 {
		goto L530
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	switch v47 - int32(_a_F_get_object_address_24) {
	case 0, 1, 2, 3, 4, 59, 60:
		v1888 = v1817
		goto L518
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L519
	default:
		goto L535
	}
L534:
	;
	v1888 = v1817
	goto L518
L535:
	;
	if base.Ui32(v47-int32(_a_F_get_object_address_71)) < base.Ui32(int32(3)) {
		v1888 = v1817
		goto L518
	} else {
		goto L536
	}
L536:
	;
	v1859 = v47 - int32(_a_F_get_object_address_26)
	if base.Ui32(int32(15)) < base.Ui32(v1859) {
		goto L519
	} else {
		goto L537
	}
L537:
	;
	if int32(1)<<(uint(v1859)%32)&int32(_a_F_get_object_address_72) != 0 {
		v1888 = v1817
		goto L518
	} else {
		goto L538
	}
L538:
	;
	goto L519
L539:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v47-int32(4060)) {
		goto L519
	} else {
		goto L540
	}
L540:
	;
	v1888 = v1817
	goto L518
L541:
	;
	goto L519
L542:
	;
	F_UnlockSharedObject(m, v47, v49, l4)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L47
	} else {
		goto L545
	}
L543:
	;
	goto L544
L544:
	;
	F_UnlockDatabaseObject(m, v47, v49, l4)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L47
	} else {
		goto L546
	}
L545:
	;
	goto L510
L546:
	;
	goto L510
L547:
	;
	if v1788 != 0 {
		goto L509
	} else {
		goto L579
	}
L548:
	;
	v1897 = int32(1)
	if v1789 <= int32(3591) {
		goto L553
	} else {
		goto L554
	}
L549:
	;
	if v1968 != 0 {
		goto L574
	} else {
		goto L575
	}
L550:
	;
	goto L549
L551:
	;
	v1968 = int32(0)
	goto L550
L552:
	;
	if base.B2i32(base.Ui32(v1789-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v1789-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v1968 = v1897
		goto L550
	} else {
		goto L573
	}
L553:
	;
	if v1789 <= int32(2670) {
		goto L556
	} else {
		goto L557
	}
L554:
	;
	goto L555
L555:
	;
	if v1789 <= int32(_a_F_get_object_address_69) {
		goto L563
	} else {
		goto L564
	}
L556:
	;
	switch v1789 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v1968 = v1897
		goto L550
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L551
	default:
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v1909 = v1789 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v1909))|base.B2i32(int32(1)<<(uint(v1909)%32)&int32(226492515) == int32(0)) != 0 {
		goto L552
	} else {
		goto L561
	}
L559:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1789-int32(2396)) {
		goto L551
	} else {
		goto L560
	}
L560:
	;
	v1968 = v1897
	goto L550
L561:
	;
	v1968 = v1897
	goto L550
L562:
	;
	if base.Ui32(v1789-int32(3592)) < base.Ui32(int32(2)) {
		v1968 = v1897
		goto L550
	} else {
		goto L571
	}
L563:
	;
	v1922 = v1789 - int32(_a_F_get_object_address_70)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v1922))|base.B2i32(int32(1)<<(uint(v1922)%32)&int32(963) == int32(0)) != 0 {
		goto L562
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	switch v1789 - int32(_a_F_get_object_address_24) {
	case 0, 1, 2, 3, 4, 59, 60:
		v1968 = v1897
		goto L550
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L551
	default:
		goto L567
	}
L566:
	;
	v1968 = v1897
	goto L550
L567:
	;
	if base.Ui32(v1789-int32(_a_F_get_object_address_71)) < base.Ui32(int32(3)) {
		v1968 = v1897
		goto L550
	} else {
		goto L568
	}
L568:
	;
	v1939 = v1789 - int32(_a_F_get_object_address_26)
	if base.Ui32(int32(15)) < base.Ui32(v1939) {
		goto L551
	} else {
		goto L569
	}
L569:
	;
	if int32(1)<<(uint(v1939)%32)&int32(_a_F_get_object_address_72) != 0 {
		v1968 = v1897
		goto L550
	} else {
		goto L570
	}
L570:
	;
	goto L551
L571:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1789-int32(4060)) {
		goto L551
	} else {
		goto L572
	}
L572:
	;
	v1968 = v1897
	goto L550
L573:
	;
	goto L551
L574:
	;
	F_LockSharedObject(m, v1789, v1787, l4)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L47
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	F_LockDatabaseObject(m, v1789, v1787, l4)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L47
	} else {
		goto L578
	}
L577:
	;
	goto L547
L578:
	;
	goto L547
L579:
	;
	v1974 = *(*int64)(unsafe.Add(mBase, _c_F_get_object_address[0]))
	if v51 == v1974 {
		goto L509
	} else {
		goto L580
	}
L580:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = v1789
	v47 = v1789
	v49 = v1787
	v50 = v1976
	v51 = v1974
	goto L2
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1788
	goto L1
}
func F_get_object_attnum_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_oid[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+20)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_oid[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_oid_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_oid_5), int32(2777), int32(_a_F_get_object_attnum_oid_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_catcache_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_name[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_name[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_catcache_name_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_catcache_name_5), int32(2777), int32(_a_F_get_object_catcache_name_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_field_end(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(1)
	v16 = v11 - v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v16))))
	if v18 != v15 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
	if v27 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v53-v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v38 = l1
	v39 = v27
	goto L9
L9:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v53 = v43
	v54 = v42
	goto L7
L11:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v11 < v12 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v14))) = uint8(v58)
	return v58
L15:
	;
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v62 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	goto L1
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v66 != 0 {
		v73 = int32(0)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v69 = F_cstring_to_text_with_len(m, v62, v67-v62)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	return int32(0)
L24:
	;
	v73 = v69
	goto L18
}
func F_get_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 != 0 {
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v5 != 0 {
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		}
	}
	return int32(0)
}
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_type[0]))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L27
	} else {
		goto L29
	}
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
	if v55 != int32(41) {
		v72 = v55
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v13 == l0 {
		v51 = v12
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v20 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_type[0])) = v47
	v51 = v47
	goto L2
L8:
	;
	v47 = v23 + int32(_a_F_get_object_type_0)
	goto L7
L9:
	;
	v47 = v23 + int32(_a_F_get_object_type_1)
	goto L7
L10:
	;
	v47 = v23 + int32(_a_F_get_object_type_2)
	goto L7
L11:
	;
	v23 = v20 * int32(40)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[1])))
	if l0 != v24 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v47 = v23 + int32(_a_F_get_object_type_3)
	goto L7
L13:
	;
	if v20 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[2])))
	if v30 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[3])))
	if v32 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[4])))
	if v34 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v20 = v20 + int32(4)
	goto L11
L20:
	;
	m.G0 = v9 + int32(16)
	return v72
L21:
	;
	v59 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v72 = int32(41)
	goto L20
L23:
	;
	v72 = int32(18)
	goto L20
L24:
	;
	v72 = int32(23)
	goto L20
L25:
	;
	v72 = int32(51)
	goto L20
L26:
	;
	v72 = int32(37)
	goto L20
L27:
	;
	return int32(0)
L28:
	;
	switch v59&int32(255) - int32(73) {
	case 0, 32:
		v72 = int32(20)
		goto L20
	default:
		goto L22
	case 10:
		goto L26
	case 29:
		goto L23
	case 36:
		goto L24
	case 45:
		goto L25
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_type_4), v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_get_object_type_5), int32(2777), int32(_a_F_get_object_type_6))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = F_superuser_arg(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			if l0 == int32(2613) {
				v22 = int32(2995)
			} else {
				v22 = l0
			}
			v23 = F_get_object_catcache_oid(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != int32(-1) {
					v27 = F_SearchSysCache1(m, v23, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = F_get_object_class_descr(m, v22)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v88
									F_errmsg_internal(m, int32(_a_F_object_ownercheck_0), v10+int32(16))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_object_ownercheck_1), int32(_a_F_object_ownercheck_2), int32(_a_F_object_ownercheck_3))
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
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
							v31 = F_get_object_attnum_owner(m, v22)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = F_SysCacheGetAttrNotNull(m, v23, v27, v31)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v27)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v71 = v33
										v74 = F_has_privs_of_role(m, l2, v71)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v77 = v74
											m.G0 = v10 + int32(80)
											return v77
										}
									}
								}
							}
						}
					}
				} else {
					v38 = F_table_open(m, v22, int32(1))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = v10 + int32(32)
						v42 = F_get_object_attnum_oid(m, v22)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_ScanKeyInit(m, v41, v42, int32(3), int32(184), l1)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = F_get_object_oid_index(m, v22)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = int32(1)
									v53 = F_systable_beginscan(m, v38, v48, v50, int32(0), v50, v41)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = F_systable_getnext(m, v53)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											if v55 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													v106 = F_get_object_class_descr(m, v22)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v106
														F_errmsg_internal(m, int32(_a_F_object_ownercheck_4), v10)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_object_ownercheck_1), int32(_a_F_object_ownercheck_5), int32(_a_F_object_ownercheck_3))
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
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
												v59 = F_get_object_attnum_owner(m, v22)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
													v64 = F_heap_getattr_2(m, v55, v59, v61, v10+int32(31))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v53)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v38, int32(1))
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v71 = v64
																v74 = F_has_privs_of_role(m, l2, v71)
																mBase = m.M
																v75 = m.ExcPending
																if v75 != 0 {
																	return int32(0)
																} else {
																	v77 = v74
																	m.G0 = v10 + int32(80)
																	return v77
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
			v77 = int32(1)
			m.G0 = v10 + int32(80)
			return v77
		}
	}
}
func F_parse_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v76
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = m.T0[v6].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15 + int32(1)
	v19 = F_json_lex(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if v12 != 0 {
		v76 = v12
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if v19 != 0 {
		v76 = v19
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v21 - int32(1) {
	case 0:
		goto L13
	default:
		goto L12
	case 3:
		goto L11
	}
L11:
	;
	v63 = F_json_lex(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L38
	}
L12:
	;
	v50 = int32(11)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v53 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v24 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v24 != 0 {
		v76 = v24
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v30 != int32(7) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v76 = v46
	goto L3
L18:
	;
	if v30 == int32(4) {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = F_json_lex(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	v35 = int32(11)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v39 = int32(13)
	goto L24
L23:
	;
	v39 = v35
	goto L24
L24:
	;
	if v30 == int32(12) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v42 = v35
	goto L27
L26:
	;
	v42 = v39
	goto L27
L27:
	;
	return v42
L28:
	;
	if v44 != 0 {
		v76 = v44
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v46 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v46 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	v54 = int32(12)
	goto L34
L33:
	;
	v54 = v50
	goto L34
L34:
	;
	if v21 == int32(12) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v57 = v50
	goto L37
L36:
	;
	v57 = v54
	goto L37
L37:
	;
	return v57
L38:
	;
	if v63 != 0 {
		v76 = v63
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65 - int32(1)
	if v5 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v70 = m.T0[v5].(func(*base.Module, int32) int32)(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v76 = int32(0)
	goto L3
L43:
	;
	if v70 != 0 {
		v76 = v70
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L42
}
