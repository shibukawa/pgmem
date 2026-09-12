package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublication(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(51), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(47062), v9)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497842), int32(1078), int32(267895))
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v34 = F_palloc(m, int32(20))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = l0
				v37 = v31 + v32
				v40 = F_pstrdup(m, v37+int32(4))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v40
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+72)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)) = uint8(v43)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+73)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)) = uint8(v45)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+74)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+17)) = uint8(v47)
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+75)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+18)) = uint8(v49)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+76)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+19)) = uint8(v51)
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+77)))
					*(*uint8)(unsafe.Add(mBase, uint32(v34)+9)) = uint8(v53)
					v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v37)+78)))
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v55
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				}
			}
		}
	}
}
func F_PublicationAddTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
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
	var v217 int32
	_ = v217
	var v221 int64
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v422 int32
	_ = v422
	var v443 int32
	_ = v443
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int64
	_ = v693
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int64
	_ = v744
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v753 int64
	_ = v753
	var v756 int64
	_ = v756
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 - int32(-64)
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = v5
	goto L4
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v40<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
	v53 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v54 = F_object_ownercheck(m, int32(1259), v51, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v54 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59)+119)))
	switch v60 - int32(73) {
	case 0, 32:
		goto L17
	default:
		v70 = int32(41)
		goto L12
	case 10:
		goto L16
	case 29:
		goto L13
	case 36:
		goto L14
	case 45:
		goto L15
	}
L9:
	;
	goto L10
L10:
	;
	v79 = v18 + int32(-28)
	v80 = m.G0
	v82 = v80 - int32(144)
	m.G0 = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	v86 = F_GetPublication(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L19
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	F_aclcheck_error(m, int32(2), v72, v73+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L18
	}
L12:
	;
	v72 = v70
	goto L11
L13:
	;
	v70 = int32(18)
	goto L12
L14:
	;
	v72 = int32(23)
	goto L11
L15:
	;
	v72 = int32(51)
	goto L11
L16:
	;
	v72 = int32(37)
	goto L11
L17:
	;
	v72 = int32(20)
	goto L11
L18:
	;
	goto L10
L19:
	;
	v90 = F_table_open(m, int32(6106), int32(3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v93 = int32(0)
	v95 = F_SearchSysCacheExists(m, int32(53), v85, l0, v93, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L24
	}
L21:
	;
	v741 = v18 + int32(-8)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v741))) = v742
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v744
	if l3 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L160
	}
L23:
	;
	m.G0 = v82 + int32(144)
	goto L21
L24:
	;
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_sequence_close(m, v90, int32(3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+119)))
	switch v128 - int32(112) {
	case 0, 2:
		goto L36
	default:
		goto L37
	}
L28:
	;
	if l2 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = *(*int64)(unsafe.Add(mBase, _consts[283]))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v101
	v104 = *(*int32)(unsafe.Add(mBase, _consts[282]))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v104
	goto L23
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v113 + int32(4)
	F_errmsg(m, int32(711365), v82)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(497842), int32(463), int32(263972))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	goto L43
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v138 + int32(4)
	F_errmsg(m, int32(267343), v82+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+119)))
	F_errdetail_relkind_not_supported(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(497842), int32(65), int32(263966))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	if base.Ui32(v156) < base.Ui32(int32(12000)) {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+118)))
	switch v160 - int32(116) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L45
	}
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v215 = F_pub_collist_validate(m, v213, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L58
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L53
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+48)) = v170 + int32(4)
	F_errmsg(m, int32(267343), v82+int32(48))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errdetail(m, int32(602171), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(497842), int32(81), int32(263966))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v195 + int32(4)
	F_errmsg(m, int32(267343), v82-int32(-64))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	F_errdetail(m, int32(602881), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(497842), int32(87), int32(263966))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v217 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+128)) = v217
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+108)) = uint8(v217)
	v221 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v82)+120)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v82)+112)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v82)+104)) = v217
	v229 = F_GetNewOidWithIndex(m, v90, int32(6112), int32(1))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+120)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v82)+116)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v82)+112)) = v229
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v234 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v242 != 0 {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v235 = F_nodeToString(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+107)) = uint8(v240)
	goto L60
L64:
	;
	v237 = F_cstring_to_text(m, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+124)) = v237
	goto L60
L66:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
	v467 = F_heap_form_tuple(m, v462, v82+int32(112), v82+int32(104))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L112
	}
L67:
	;
	v243 = int32(0)
	if v215 == v243 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+108)) = uint8(v443)
	goto L66
L70:
	;
	v281 = F_buildint2vector(m, v243, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L83
	}
L71:
	;
	v280 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v252 = int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v253 <= v252 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v256 = v252
	goto L76
L75:
	;
	v256 = v253
	goto L76
L76:
	;
	v260 = int32(0)
	v262 = v243
	goto L77
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v215+int32(8)+v260<<(uint(int32(2))%32))))
	if v268 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v280 = v271
	goto L70
L79:
	;
	v271 = v262 + base.I32_popcnt(v268)
	goto L81
L80:
	;
	v271 = v262
	goto L81
L81:
	;
	v273 = v260 + int32(1)
	if v273 != v256 {
		v260 = v273
		v262 = v271
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if v215 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if int32(0) <= v339 {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	v339 = base.I32_ctz(v325) | v326<<(uint(int32(5))%32)
	goto L84
L86:
	;
	v339 = int32(-2)
	goto L84
L87:
	;
	v292 = base.I32_div_s(int32(0), int32(32))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v293 <= v292 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v296 = v215 + int32(8)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v296+v292<<(uint(int32(2))%32))))
	v303 = v300 & int32(-1)
	if v303 != 0 {
		v325 = v303
		v326 = v292
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v305 = v292 + int32(1)
	if v305 == v293 {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v308 = v305
	goto L91
L91:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v296+v308<<(uint(int32(2))%32))))
	if v315 != 0 {
		v325 = v315
		v326 = v308
		goto L85
	} else {
		goto L93
	}
L92:
	;
	goto L86
L93:
	;
	v317 = v308 + int32(1)
	if v317 != v293 {
		v308 = v317
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v349 = v243
	v352 = v339
	goto L98
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+128)) = v281
	goto L66
L98:
	;
	v361 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v281+int32(24)+v349<<(uint(v361)%32)))) = uint16(v352)
	if v215 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L97
L100:
	;
	if int32(0) <= v422 {
		v349 = v349 + v361
		v352 = v422
		goto L98
	} else {
		goto L111
	}
L101:
	;
	v422 = base.I32_ctz(v408) | v409<<(uint(int32(5))%32)
	goto L100
L102:
	;
	v422 = int32(-2)
	goto L100
L103:
	;
	v373 = v352 + int32(1)
	v375 = base.I32_div_s(v373, int32(32))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v376 <= v375 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v379 = v215 + int32(8)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379+v375<<(uint(int32(2))%32))))
	v386 = v383 & (int32(-1) << (uint(v373) % 32))
	if v386 != 0 {
		v408 = v386
		v409 = v375
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v388 = v375 + int32(1)
	if v388 == v376 {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v391 = v388
	goto L107
L107:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v379+v391<<(uint(int32(2))%32))))
	if v398 != 0 {
		v408 = v398
		v409 = v391
		goto L101
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v400 = v391 + int32(1)
	if v400 != v376 {
		v391 = v400
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L99
L112:
	;
	F_CatalogTupleInsert(m, v90, v467)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_pfree(m, v467)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+100)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v82)+96)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v82)+92)) = int32(6106)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+88)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v82)+84)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = int32(6104)
	F_recordDependencyOn(m, v82+int32(92), v82+int32(80), int32(97))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+84)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = int32(1259)
	F_recordDependencyOn(m, v82+int32(92), v82+int32(80), int32(97))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v502 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_recordDependencyOnSingleRelExpr(m, v82+int32(92), v502, v85, int32(110), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v215 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L119
L121:
	;
	if int32(0) <= v565 {
		goto L132
	} else {
		goto L133
	}
L122:
	;
	v565 = base.I32_ctz(v551) | v552<<(uint(int32(5))%32)
	goto L121
L123:
	;
	v565 = int32(-2)
	goto L121
L124:
	;
	v518 = base.I32_div_s(int32(0), int32(32))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v519 <= v518 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v522 = v215 + int32(8)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522+v518<<(uint(int32(2))%32))))
	v529 = v526 & int32(-1)
	if v529 != 0 {
		v551 = v529
		v552 = v518
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v531 = v518 + int32(1)
	if v531 == v519 {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v534 = v531
	goto L128
L128:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v522+v534<<(uint(int32(2))%32))))
	if v541 != 0 {
		v551 = v541
		v552 = v534
		goto L122
	} else {
		goto L130
	}
L129:
	;
	goto L123
L130:
	;
	v543 = v534 + int32(1)
	if v543 != v519 {
		v534 = v543
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v573 = v565
	goto L135
L133:
	;
	goto L134
L134:
	;
	F_sequence_close(m, v90, int32(3))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L6
	} else {
		goto L150
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+88)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v82)+84)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v82)+80)) = int32(1259)
	F_recordDependencyOn(m, v82+int32(92), v82+int32(80), int32(110))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L137
	}
L136:
	;
	goto L134
L137:
	;
	if v215 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	if int32(0) <= v651 {
		v573 = v651
		goto L135
	} else {
		goto L149
	}
L139:
	;
	v651 = base.I32_ctz(v637) | v638<<(uint(int32(5))%32)
	goto L138
L140:
	;
	v651 = int32(-2)
	goto L138
L141:
	;
	v602 = v573 + int32(1)
	v604 = base.I32_div_s(v602, int32(32))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v605 <= v604 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v608 = v215 + int32(8)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608+v604<<(uint(int32(2))%32))))
	v615 = v612 & (int32(-1) << (uint(v602) % 32))
	if v615 != 0 {
		v637 = v615
		v638 = v604
		goto L139
	} else {
		goto L143
	}
L143:
	;
	v617 = v604 + int32(1)
	if v617 == v605 {
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v620 = v617
	goto L145
L145:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v608+v620<<(uint(int32(2))%32))))
	if v627 != 0 {
		v637 = v627
		v638 = v620
		goto L139
	} else {
		goto L147
	}
L146:
	;
	goto L140
L147:
	;
	v629 = v620 + int32(1)
	if v629 != v605 {
		v620 = v629
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L136
L150:
	;
	v674 = F_get_rel_relkind(m, v85)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L152
	}
L151:
	;
	F_InvalidatePublicationRels(m, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L159
	}
L152:
	;
	if v674 == int32(112) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v678 = int32(0)
	v681 = F_find_all_inheritors(m, v85, v678, v678)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L6
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v686 = F_lappend_oid(m, int32(0), v85)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	v683 = F_list_concat(m, v678, v681)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	v688 = v683
	goto L151
L158:
	;
	v688 = v686
	goto L151
L159:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v82)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v691
	v693 = *(*int64)(unsafe.Add(mBase, uint32(v82)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v693
	goto L23
L160:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v722 + int32(4)
	F_errmsg(m, int32(267343), v82+int32(32))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	F_errdetail(m, int32(602594), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(497842), int32(73), int32(263966))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	v775 = v40 + int32(1)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v775 < v776 {
		v40 = v775
		goto L4
	} else {
		goto L170
	}
L166:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v748
	v751 = *(*int32)(unsafe.Add(mBase, _consts[282]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v751
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v753
	v756 = *(*int64)(unsafe.Add(mBase, _consts[283]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v756
	F_EventTriggerCollectSimpleCommand(m, v18+int32(-40), v18+int32(-56), l3)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	if v765 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v770 = int32(0)
	F_RunObjectPostCreateHook(m, int32(6106), v769, v770, v770)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L165
L170:
	;
	goto L5
}
