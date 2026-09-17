package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v878 int32
	_ = v878
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v937 int32
	_ = v937
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == v2 {
		v42 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v42&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v29 == int32(0) {
		v42 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 != int32(7) {
		v42 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 != int32(17) {
		v42 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	v42 = v38 ^ int32(1)
	goto L2
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = F_get_fn_opclass_options(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v53 = int32(128)
	goto L9
L9:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v54)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)))
	if v57&int32(4) != 0 {
		v955 = v54
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v53 = v50 << (uint(int32(3)) % 32)
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L10
	} else {
		goto L129
	}
L13:
	;
	m.G0 = v18 + int32(16)
	return v955
L14:
	;
	v61 = v23 + int32(8)
	switch v21 - int32(7) {
	case 0, 6:
		goto L18
	default:
		goto L12
	case 2:
		goto L17
	case 3:
		goto L15
	case 4:
		goto L16
	}
L15:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v761 = F_pg_detoast_datum(m, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L10
	} else {
		goto L106
	}
L16:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v568 = F_pg_detoast_datum(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L10
	} else {
		goto L85
	}
L17:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v403 = F_pg_detoast_datum_packed(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L10
	} else {
		goto L61
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = F_hstoreUpgrade(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = v67 & int32(268435455)
	if v69 == int32(0) {
		v955 = v54
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v73 = v65 + int32(8)
	v76 = v73 + v69<<(uint(int32(3))%32)
	v82 = v2
	goto L21
L21:
	;
	v94 = v73 + v82<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v955 = v385
	goto L13
L23:
	;
	if v110 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v110 = v95 & int32(1073741823)
	v111 = v76
	goto L23
L25:
	;
	goto L26
L26:
	;
	v100 = int32(1073741823)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v94-int32(4))))
	v106 = v104 & v100
	v110 = v95&v100 - v106
	v111 = v106 + v76
	goto L23
L27:
	;
	v113 = int32(-1)
	if v110 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v227 = int32(0)
	goto L29
L29:
	;
	v228 = base.I32_rem_u_s(v227, v53)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v228)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v232)>>(uint(v228&int32(7))%32))&int32(1) == int32(0) {
		v955 = int32(0)
		goto L13
	} else {
		goto L38
	}
L30:
	;
	v227 = v195 ^ int32(-1)
	goto L29
L31:
	;
	v121 = v111
	v122 = v113
	v123 = int32(0)
	goto L34
L32:
	;
	v167 = v111
	v168 = v113
	goto L33
L33:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32((v168^v182)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v195 = v190 ^ int32(base.Ui32(v168)>>(uint(int32(8))%32))
	goto L30
L34:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v138 = int32(255)
	v140 = int32(2)
	v144 = *(*int32)(unsafe.Add(mBase, uint32((v122^v136)&v138<<(uint(v140)%32))+uint32(_c_F_ghstore_consistent[0])))
	v145 = int32(8)
	v147 = v144 ^ int32(base.Ui32(v122)>>(uint(v145)%32))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32((v147^v148)&v138<<(uint(v140)%32))+uint32(_c_F_ghstore_consistent[0])))
	v159 = v156 ^ int32(base.Ui32(v147)>>(uint(v145)%32))
	v161 = v121 + v140
	v163 = v123 + v140
	if v163 != v110&int32(-2) {
		v121 = v161
		v122 = v159
		v123 = v163
		goto L34
	} else {
		goto L36
	}
L35:
	;
	if v110&int32(1) == int32(0) {
		v195 = v159
		goto L30
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v167 = v161
	v168 = v159
	goto L33
L38:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v241&int32(1073741824) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v246 = int32(1073741823)
	v247 = v95 & v246
	v252 = base.B2i32(int32(0) <= v241)
	if int32(0) <= v241 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v385 = int32(1)
	goto L41
L41:
	;
	if v385 == int32(0) {
		v955 = v385
		goto L13
	} else {
		goto L59
	}
L42:
	;
	v253 = v241 - v247
	goto L44
L43:
	;
	v253 = v241 & v246
	goto L44
L44:
	;
	if v253 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v254 = int32(-1)
	if int32(0) <= v241 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v371 = int32(0)
	goto L47
L47:
	;
	v372 = base.I32_rem_u_s(v371, v53)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v372)>>(uint(int32(3))%32))))))
	v385 = int32(base.Ui32(v376)>>(uint(v372&int32(7))%32)) & int32(1)
	goto L41
L48:
	;
	v256 = v247
	goto L50
L49:
	;
	v256 = int32(0)
	goto L50
L50:
	;
	v257 = v76 + v256
	if v253 != int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v371 = v339 ^ int32(-1)
	goto L47
L52:
	;
	v265 = v257
	v266 = v254
	v267 = int32(0)
	goto L55
L53:
	;
	v311 = v257
	v312 = v254
	goto L54
L54:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v334 = *(*int32)(unsafe.Add(mBase, uint32((v312^v326)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v339 = v334 ^ int32(base.Ui32(v312)>>(uint(int32(8))%32))
	goto L51
L55:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v282 = int32(255)
	v284 = int32(2)
	v288 = *(*int32)(unsafe.Add(mBase, uint32((v266^v280)&v282<<(uint(v284)%32))+uint32(_c_F_ghstore_consistent[0])))
	v289 = int32(8)
	v291 = v288 ^ int32(base.Ui32(v266)>>(uint(v289)%32))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32((v291^v292)&v282<<(uint(v284)%32))+uint32(_c_F_ghstore_consistent[0])))
	v303 = v300 ^ int32(base.Ui32(v291)>>(uint(v289)%32))
	v305 = v265 + v284
	v307 = v267 + v284
	if v307 != v253&int32(-2) {
		v265 = v305
		v266 = v303
		v267 = v307
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v253&int32(1) == int32(0) {
		v339 = v303
		goto L51
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v311 = v305
	v312 = v303
	goto L54
L59:
	;
	v400 = v82 + int32(1)
	if base.Ui32(v400) < base.Ui32(v69) {
		v82 = v400
		goto L21
	} else {
		goto L60
	}
L60:
	;
	goto L22
L61:
	;
	v405 = int32(1)
	v406 = v403 + v405
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v409 = v407 & v405
	if v407 == v405 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v557 = base.I32_rem_u_s(v556, v53)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v557)>>(uint(int32(3))%32))))))
	v955 = int32(base.Ui32(v561)>>(uint(v557&int32(7))%32)) & int32(1)
	goto L13
L63:
	;
	if v409 != 0 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	if v436 != 0 {
		v438 = v436
		goto L63
	} else {
		goto L73
	}
L65:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if base.Ui32((v413-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v438 = int32(4)
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v425 = int32(1)
	if v409 != 0 {
		v436 = int32(base.Ui32(v407)>>(uint(v425)%32)) - v425
		goto L64
	} else {
		goto L72
	}
L68:
	;
	if v413 == int32(18) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v424 = int32(16)
	goto L71
L70:
	;
	v424 = int32(0)
	goto L71
L71:
	;
	v436 = v424
	goto L64
L72:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v436 = int32(base.Ui32(v429)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L73:
	;
	v556 = int32(0)
	goto L62
L74:
	;
	v442 = v406
	goto L76
L75:
	;
	v442 = v403 + int32(4)
	goto L76
L76:
	;
	v443 = int32(-1)
	if v438 != int32(1) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v556 = v525 ^ int32(-1)
	goto L62
L78:
	;
	v451 = v442
	v452 = v443
	v453 = int32(0)
	goto L81
L79:
	;
	v497 = v442
	v498 = v443
	goto L80
L80:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v520 = *(*int32)(unsafe.Add(mBase, uint32((v498^v512)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v525 = v520 ^ int32(base.Ui32(v498)>>(uint(int32(8))%32))
	goto L77
L81:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451))))
	v468 = int32(255)
	v470 = int32(2)
	v474 = *(*int32)(unsafe.Add(mBase, uint32((v452^v466)&v468<<(uint(v470)%32))+uint32(_c_F_ghstore_consistent[0])))
	v475 = int32(8)
	v477 = v474 ^ int32(base.Ui32(v452)>>(uint(v475)%32))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+1)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32((v477^v478)&v468<<(uint(v470)%32))+uint32(_c_F_ghstore_consistent[0])))
	v489 = v486 ^ int32(base.Ui32(v477)>>(uint(v475)%32))
	v491 = v451 + v470
	v493 = v453 + v470
	if v493 != v438&int32(-2) {
		v451 = v491
		v452 = v489
		v453 = v493
		goto L81
	} else {
		goto L83
	}
L82:
	;
	if v438&int32(1) == int32(0) {
		v525 = v489
		goto L77
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v497 = v491
	v498 = v489
	goto L80
L85:
	;
	F_deconstruct_array_builtin(m, v568, int32(25), v18+int32(12), v18+int32(8), v18+int32(4))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v579 <= int32(0) {
		v955 = v54
		goto L13
	} else {
		goto L87
	}
L87:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v589 = v2
	goto L88
L88:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v583))))
	if v601 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v955 = v743
	goto L13
L90:
	;
	v604 = int32(2)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v582+v589<<(uint(v604)%32))))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v610 = int32(base.Ui32(v608) >> (uint(v604) % 32))
	v612 = v610 - int32(4)
	if v612 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v743 = int32(1)
	goto L92
L92:
	;
	if v743 == int32(0) {
		v955 = v743
		goto L13
	} else {
		goto L104
	}
L93:
	;
	v614 = v607 + int32(4)
	v615 = int32(-1)
	if v610 != int32(5) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v729 = int32(0)
	goto L95
L95:
	;
	v730 = base.I32_rem_u_s(v729, v53)
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v730)>>(uint(int32(3))%32))))))
	v743 = int32(base.Ui32(v734)>>(uint(v730&int32(7))%32)) & int32(1)
	goto L92
L96:
	;
	v729 = v697 ^ int32(-1)
	goto L95
L97:
	;
	v623 = v614
	v624 = v615
	v625 = int32(0)
	goto L100
L98:
	;
	v669 = v614
	v670 = v615
	goto L99
L99:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669))))
	v692 = *(*int32)(unsafe.Add(mBase, uint32((v670^v684)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v697 = v692 ^ int32(base.Ui32(v670)>>(uint(int32(8))%32))
	goto L96
L100:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
	v640 = int32(255)
	v642 = int32(2)
	v646 = *(*int32)(unsafe.Add(mBase, uint32((v624^v638)&v640<<(uint(v642)%32))+uint32(_c_F_ghstore_consistent[0])))
	v647 = int32(8)
	v649 = v646 ^ int32(base.Ui32(v624)>>(uint(v647)%32))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+1)))
	v658 = *(*int32)(unsafe.Add(mBase, uint32((v649^v650)&v640<<(uint(v642)%32))+uint32(_c_F_ghstore_consistent[0])))
	v661 = v658 ^ int32(base.Ui32(v649)>>(uint(v647)%32))
	v663 = v623 + v642
	v665 = v625 + v642
	if v665 != v612&int32(-2) {
		v623 = v663
		v624 = v661
		v625 = v665
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if v610&int32(1) == int32(0) {
		v697 = v661
		goto L96
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v669 = v663
	v670 = v661
	goto L99
L104:
	;
	v758 = v589 + int32(1)
	if v758 < v579 {
		v589 = v758
		goto L88
	} else {
		goto L105
	}
L105:
	;
	goto L89
L106:
	;
	F_deconstruct_array_builtin(m, v761, int32(25), v18+int32(12), v18+int32(8), v18+int32(4))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v772 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v955 = int32(0)
	goto L13
L109:
	;
	goto L110
L110:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v783 = v2
	goto L111
L111:
	;
	v793 = int32(0)
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v777))))
	if v795 == v793 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v955 = v937
	goto L13
L113:
	;
	v798 = int32(2)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v776+v783<<(uint(v798)%32))))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	v804 = int32(base.Ui32(v802) >> (uint(v798) % 32))
	v806 = v804 - int32(4)
	if v806 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v937 = v793
	goto L115
L115:
	;
	if v937 != 0 {
		v955 = v937
		goto L13
	} else {
		goto L127
	}
L116:
	;
	v808 = v801 + int32(4)
	v809 = int32(-1)
	if v804 != int32(5) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v923 = int32(0)
	goto L118
L118:
	;
	v924 = base.I32_rem_u_s(v923, v53)
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v924)>>(uint(int32(3))%32))))))
	v937 = int32(base.Ui32(v928)>>(uint(v924&int32(7))%32)) & int32(1)
	goto L115
L119:
	;
	v923 = v891 ^ int32(-1)
	goto L118
L120:
	;
	v817 = v808
	v818 = v809
	v819 = int32(0)
	goto L123
L121:
	;
	v863 = v808
	v864 = v809
	goto L122
L122:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	v886 = *(*int32)(unsafe.Add(mBase, uint32((v864^v878)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v891 = v886 ^ int32(base.Ui32(v864)>>(uint(int32(8))%32))
	goto L119
L123:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	v834 = int32(255)
	v836 = int32(2)
	v840 = *(*int32)(unsafe.Add(mBase, uint32((v818^v832)&v834<<(uint(v836)%32))+uint32(_c_F_ghstore_consistent[0])))
	v841 = int32(8)
	v843 = v840 ^ int32(base.Ui32(v818)>>(uint(v841)%32))
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+1)))
	v852 = *(*int32)(unsafe.Add(mBase, uint32((v843^v844)&v834<<(uint(v836)%32))+uint32(_c_F_ghstore_consistent[0])))
	v855 = v852 ^ int32(base.Ui32(v843)>>(uint(v841)%32))
	v857 = v817 + v836
	v859 = v819 + v836
	if v859 != v806&int32(-2) {
		v817 = v857
		v818 = v855
		v819 = v859
		goto L123
	} else {
		goto L125
	}
L124:
	;
	if v804&int32(1) == int32(0) {
		v891 = v855
		goto L119
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v863 = v857
	v864 = v855
	goto L122
L127:
	;
	v950 = v783 + int32(1)
	if v950 < v772 {
		v783 = v950
		goto L111
	} else {
		goto L128
	}
L128:
	;
	goto L112
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
	F_errmsg_internal(m, int32(_a_F_ghstore_consistent_0), v18)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ghstore_consistent_1), int32(608), int32(_a_F_ghstore_consistent_2))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
