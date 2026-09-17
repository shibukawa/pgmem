package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayCastAndSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v95 int32
	_ = v95
	v1 = l0
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if int32(0) < l1 {
		if l2 != 0 {
			switch l1 - int32(1) {
			case 0:
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1)
				switch l3 - int32(99) {
				case 0:
					v95 = l1
				case 1:
					v95 = (l1 + int32(7)) & int32(-8)
				default:
					v95 = (l1 + int32(1)) & int32(-2)
				case 6:
					v95 = (l1 + int32(3)) & int32(-4)
				}
				m.G0 = v9 + int32(16)
				return v95
			case 1:
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v1)
				switch l3 - int32(99) {
				case 0:
					v95 = l1
				case 1:
					v95 = (l1 + int32(7)) & int32(-8)
				default:
					v95 = (l1 + int32(1)) & int32(-2)
				case 6:
					v95 = (l1 + int32(3)) & int32(-4)
				}
				m.G0 = v9 + int32(16)
				return v95
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_ArrayCastAndSet_0), v9)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ArrayCastAndSet_1), int32(230), int32(_a_F_ArrayCastAndSet_2))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1
				switch l3 - int32(99) {
				case 0:
					v95 = l1
				case 1:
					v95 = (l1 + int32(7)) & int32(-8)
				default:
					v95 = (l1 + int32(1)) & int32(-2)
				case 6:
					v95 = (l1 + int32(3)) & int32(-4)
				}
				m.G0 = v9 + int32(16)
				return v95
			}
		} else {
			if l1 == int32(0) {
			} else {
				base.MemoryCopy(m, l4, v1, l1)
			}
			switch l3 - int32(99) {
			case 0:
				v95 = l1
			case 1:
				v95 = (l1 + int32(7)) & int32(-8)
			default:
				v95 = (l1 + int32(1)) & int32(-2)
			case 6:
				v95 = (l1 + int32(3)) & int32(-4)
			}
			m.G0 = v9 + int32(16)
			return v95
		}
	} else {
		if l1 == int32(-1) {
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1))))
			if v52 == int32(1) {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1)+1)))
				if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v79 = int32(6)
				} else {
					v63 = int32(18)
					if v56 == v63 {
						v67 = v63
					} else {
						v67 = int32(2)
					}
					v79 = v67
				}
			} else {
				v68 = int32(1)
				if v52&v68 != 0 {
					v79 = int32(base.Ui32(v52) >> (uint(v68) % 32))
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v1)))
					v79 = int32(base.Ui32(v72) >> (uint(int32(2)) % 32))
				}
			}
		} else {
			v75 = F_strlen(m, v1)
			mBase = m.M
			v79 = v75 + int32(1)
		}
		if v79 != 0 {
			base.MemoryCopy(m, l4, v1, v79)
		} else {
		}
		switch l3 - int32(99) {
		case 0:
			v95 = v79
		case 1:
			v95 = (v79 + int32(7)) & int32(-8)
		default:
			v95 = (v79 + int32(1)) & int32(-2)
		case 6:
			v95 = (v79 + int32(3)) & int32(-4)
		}
		m.G0 = v9 + int32(16)
		return v95
	}
}
func F_array_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	v7 = int32(0)
	if l3 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_0))
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_1))
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	if int32(0) < l3 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L10:
	;
	m.T0[l4].(func(*base.Module, int32, int32, int32))(m, l0, l1+l2*v27, l5)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if v27 < l3-int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_array_desc_2))
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v38 = v27 + int32(1)
	if v38 != l3 {
		v27 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L11
L18:
	;
	return
}
func F_array_dims(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_DatumGetAnyArrayP(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(224)
	return v95
L2:
	;
	return int32(0)
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v20 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = int32(28)
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14+v23)))
	if base.Ui32(v25-int32(7)) <= base.Ui32(int32(-7)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
	v95 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	if v20 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v48 = v11 + int32(16)
	v51 = int32(0)
	goto L14
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v43 = v35
	v44 = v36
	goto L10
L12:
	;
	goto L13
L13:
	;
	v38 = v14 + int32(16)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v43 = v38
	v44 = v38 + v39<<(uint(int32(2))%32)
	goto L10
L14:
	;
	v57 = v51 << (uint(int32(2)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43+v57)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v61 + v59 - int32(1)
	v68 = F_pg_sprintf(m, v48, int32(_a_F_array_dims_0), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v85 = F_cstring_to_text(m, v11+int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L21
	}
L16:
	;
	v70 = F_strlen(m, v48)
	mBase = m.M
	v73 = v51 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v76 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v79 = int32(28)
	goto L19
L18:
	;
	v79 = int32(4)
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14+v79)))
	if v73 < v81 {
		v48 = v70 + v48
		v51 = v73
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v95 = v85
	goto L1
}
func F_array_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v88 int64
	_ = v88
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v580 int32
	_ = v580
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v962 int32
	_ = v962
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v985 int32
	_ = v985
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1228 int32
	_ = v1228
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1331 int32
	_ = v1331
	var v1357 int32
	_ = v1357
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1392 int32
	_ = v1392
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1487 int32
	_ = v1487
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1527 int32
	_ = v1527
	var v1541 int32
	_ = v1541
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1601 int32
	_ = v1601
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(576)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76)+8)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+4)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+7)))
	v82 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+512)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v27)+504)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v27)+496)) = v82
	v88 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+480)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v27)+472)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v27)+464)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v32
	v98 = v32
	v102 = v2
	goto L11
L2:
	;
	F_get_type_io_data(m, v31, int32(0), v52+int32(4), v52+int32(6), v52+int32(7), v52+int32(8), v52+int32(12), v52+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L9
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v39 = F_MemoryContextAlloc(m, v37, int32(48))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v50 == v31 {
		v76 = v34
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v39
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v31 ^ int32(-1)
	v52 = v46
	goto L2
L8:
	;
	v52 = v34
	goto L2
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	F_fmgr_info_cxt(m, v68, v52+int32(20), v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v31
	v76 = v52
	goto L1
L11:
	;
	v121 = v98 + int32(1)
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	goto L13
L12:
	;
	m.G0 = v27 + int32(576)
	return v1601
L13:
	;
	if base.B2i32(v122 == int32(32))|base.B2i32(base.Ui32((v122-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v98 = v121
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v98
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v133 == int32(91) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(496)+v102<<(uint(int32(2))%32)))) = v303
	v98 = v262
	v102 = v102 + int32(1)
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v121
	if v102 == int32(6) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L19
L19:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v102 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L20:
	;
	v1601 = int32(0)
	goto L15
L21:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), v324, int32(_a_F_array_in_1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L83
	}
L22:
	;
	v140 = F_errsave_start(m, v29)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v157 = F_ReadDimensionInt(m, v27+int32(528), v27+int32(524), v29)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	if v140 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_2), v27)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v324 = int32(432)
	goto L21
L29:
	;
	if v157 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v27)+528))
	if v121 == v161 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v163 = F_errsave_start(m, v29)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v181 == int32(58) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	if v163 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errdetail(m, int32(_a_F_array_in_4), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v324 = int32(441)
	goto L21
L39:
	;
	if v234&int32(255) != int32(93) {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v27)+524))
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(464)+v102<<(uint(int32(2))%32)))) = v189
	v192 = v161 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v192
	v198 = F_ReadDimensionInt(m, v27+int32(528), v27+int32(560), v29)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v223 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(464)+v102<<(uint(int32(2))%32)))) = v223
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v27)+524))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+560)) = v231
	v233 = v161
	v234 = v181
	v235 = v223
	goto L39
L43:
	;
	if v198 == int32(0) {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v27)+528))
	if v192 != v202 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v233 = v202
	v234 = v204
	v235 = v189
	goto L39
L46:
	;
	goto L47
L47:
	;
	v205 = F_errsave_start(m, v29)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v205 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errdetail(m, int32(_a_F_array_in_5), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v324 = int32(455)
	goto L21
L53:
	;
	v240 = F_errsave_start(m, v29)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v262 = v233 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+528)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v27)+560))
	if v264 < v235 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	if v240 == int32(0) {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(96))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = int32(_a_F_array_in_6)
	F_errdetail(m, int32(_a_F_array_in_7), v27+int32(80))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v324 = int32(468)
	goto L21
L61:
	;
	v266 = F_errsave_start(m, v29)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v264 == int32(2147483647) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v266 == int32(0) {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_array_in_8), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v324 = int32(481)
	goto L21
L68:
	;
	v280 = F_errsave_start(m, v29)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v295 = int32(0)
	v297 = v264 - v235
	if base.B2i32(v295 < v235)^base.B2i32(v297 < v264) == v295 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	if v280 == int32(0) {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = int32(2147483647)
	F_errmsg(m, int32(_a_F_array_in_9), v27+int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v324 = int32(487)
	goto L21
L75:
	;
	v303 = v297 + int32(1)
	if v297 <= v303 {
		goto L16
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v306 = F_errsave_start(m, v29)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v306 == int32(0) {
		goto L20
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_in_10), v27-int32(-64))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v324 = int32(495)
	goto L21
L83:
	;
	goto L20
L84:
	;
	v480 = F_palloc(m, int32(64))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L115
	}
L85:
	;
	if v332 == int32(123) {
		v450 = v98
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v332 == int32(61) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v337 = int32(0)
	v338 = F_errsave_start(m, v29)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	if v338 == int32(0) {
		v1601 = v337
		goto L15
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+400)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(400))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	F_errdetail(m, int32(_a_F_array_in_11), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(265), int32(_a_F_array_in_12))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v1601 = v337
	goto L15
L95:
	;
	v362 = v98
	goto L98
L96:
	;
	goto L97
L97:
	;
	v424 = int32(0)
	v425 = F_errsave_start(m, v29)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L109
	}
L98:
	;
	v387 = v362 + int32(1)
	v388 = int32(*(*int8)(unsafe.Add(mBase, uint32(v362)+1)))
	goto L100
L99:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v398 == int32(123) {
		v450 = v387
		goto L84
	} else {
		goto L102
	}
L100:
	;
	if base.B2i32(v388 == int32(32))|base.B2i32(base.Ui32((v388-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v362 = v387
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v401 = int32(0)
	v402 = F_errsave_start(m, v29)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	if v402 == int32(0) {
		v1601 = v401
		goto L15
	} else {
		goto L104
	}
L104:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+416)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(416))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	F_errdetail(m, int32(_a_F_array_in_13), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(285), int32(_a_F_array_in_12))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v1601 = v401
	goto L15
L109:
	;
	if v425 == int32(0) {
		v1601 = v424
		goto L15
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+448)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(448))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+432)) = int32(_a_F_array_in_14)
	F_errdetail(m, int32(_a_F_array_in_7), v27+int32(432))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(275), int32(_a_F_array_in_12))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v1601 = v424
	goto L15
L115:
	;
	v483 = F_palloc(m, int32(16))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_initStringInfo(m, v27+int32(560))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v490 = v450
	v494 = int32(0)
	v497 = base.B2i32(v102 != int32(0))
	v498 = v102
	v499 = v2
	v500 = v2
	v502 = v480
	v503 = v483
	v505 = int32(16)
	goto L121
L118:
	;
	if v1185 != 0 {
		goto L326
	} else {
		goto L327
	}
L119:
	;
	v1601 = int32(0)
	goto L15
L120:
	;
	v1289 = F_errsave_start(m, v29)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L6
	} else {
		goto L312
	}
L121:
	;
	v515 = v27 + int32(560)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v517 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v516))) = uint8(v517)
	*(*int32)(unsafe.Add(mBase, uint32(v515)+12)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v515)+4)) = v517
	goto L123
L122:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v27)+560))
	F_pfree(m, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L6
	} else {
		goto L300
	}
L123:
	;
	v523 = v490
	goto L129
L124:
	;
	if int32(0) < v1179 {
		v490 = v1175
		v494 = v1179
		v497 = v1182
		v498 = v1183
		v499 = v1184
		v500 = v1185
		v502 = v1187
		v503 = v1188
		v505 = v1190
		goto L121
	} else {
		goto L299
	}
L125:
	;
	if v499 != 0 {
		goto L268
	} else {
		goto L269
	}
L126:
	;
	v1063 = v813
	v1070 = int32(0)
	goto L125
L127:
	;
	if v499 != 0 {
		goto L259
	} else {
		goto L260
	}
L128:
	;
	v1011 = F_errsave_start(m, v29)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L6
	} else {
		goto L253
	}
L129:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	switch v547 - int32(123) {
	case 0:
		goto L133
	case 1:
		goto L131
	case 2:
		goto L132
	default:
		goto L134
	}
L130:
	;
	v813 = v523
	v820 = int32(0)
	v824 = int32(1)
	goto L204
L131:
	;
	v796 = base.I32_extend8_s(v547)
	if v796 == v78 {
		goto L127
	} else {
		goto L198
	}
L132:
	;
	v736 = v27 + int32(528) + v494<<(uint(int32(2))%32)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v736-int32(4))))
	v740 = int32(0)
	if v499|base.B2i32(v739 <= v740) == v740 {
		goto L182
	} else {
		goto L183
	}
L133:
	;
	if v499 != 0 {
		goto L161
	} else {
		goto L162
	}
L134:
	;
	if v547 == int32(0) {
		goto L128
	} else {
		goto L135
	}
L135:
	;
	if v547 != int32(34) {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v556 = v523 + int32(1)
	goto L137
L137:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if v580 != int32(92) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v600 = v556
	goto L148
L139:
	;
	goto L138
L140:
	;
	F_appendStringInfoChar(m, v27+int32(560), base.I32_extend8_s(v592))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L147
	}
L141:
	;
	if v580 == int32(0) {
		goto L128
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	if v588 == int32(0) {
		goto L128
	} else {
		goto L146
	}
L144:
	;
	if v580 == int32(34) {
		goto L139
	} else {
		goto L145
	}
L145:
	;
	v592 = v580
	v593 = int32(1)
	goto L140
L146:
	;
	v592 = v588
	v593 = int32(2)
	goto L140
L147:
	;
	v556 = v556 + v593
	goto L137
L148:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+1)))
	if v624 == int32(0) {
		goto L128
	} else {
		goto L150
	}
L149:
	;
	v645 = F_errsave_start(m, v29)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L6
	} else {
		goto L155
	}
L150:
	;
	v628 = v600 + int32(1)
	v629 = int32(0)
	v630 = base.I32_extend8_s(v624)
	if v630 == v78 {
		v1063 = v628
		v1070 = v629
		goto L125
	} else {
		goto L151
	}
L151:
	;
	switch v630&int32(255) - int32(123) {
	case 0, 2:
		v1063 = v628
		v1070 = v629
		goto L125
	default:
		goto L152
	}
L152:
	;
	goto L153
L153:
	;
	if base.B2i32(v630 == int32(32))|base.B2i32(base.Ui32((v630-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v600 = v628
		goto L148
	} else {
		goto L154
	}
L154:
	;
	goto L149
L155:
	;
	if v645 == int32(0) {
		goto L119
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+384)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(384))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	F_errdetail(m, int32(_a_F_array_in_15), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(871), int32(_a_F_array_in_16))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v1601 = int32(0)
	goto L15
L161:
	;
	v668 = F_errsave_start(m, v29)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L6
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(int32(6)) <= base.Ui32(v494) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	if v668 == int32(0) {
		goto L119
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+320)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(320))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(123)
	F_errdetail(m, int32(_a_F_array_in_17), v27+int32(304))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(637), int32(_a_F_array_in_18))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	v1601 = int32(0)
	goto L15
L170:
	;
	v696 = F_errsave_start(m, v29)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v716 = int32(1)
	v717 = v523 + v716
	v718 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(528)+v494<<(uint(int32(2))%32)))) = v718
	v727 = v494 + v716
	if v494 < v498 {
		goto L178
	} else {
		goto L179
	}
L173:
	;
	if v696 == int32(0) {
		goto L119
	} else {
		goto L174
	}
L174:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+336)) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_2), v27+int32(336))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(644), int32(_a_F_array_in_18))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v1601 = int32(0)
	goto L15
L178:
	;
	v1175 = v717
	v1179 = v727
	v1182 = v497
	v1183 = v498
	v1184 = v718
	v1185 = v500
	v1187 = v502
	v1188 = v503
	v1190 = v505
	goto L124
L179:
	;
	goto L180
L180:
	;
	if v497&int32(1) != 0 {
		goto L120
	} else {
		goto L181
	}
L181:
	;
	v1175 = v717
	v1179 = v727
	v1182 = int32(0)
	v1183 = v727
	v1184 = v718
	v1185 = v500
	v1187 = v502
	v1188 = v503
	v1190 = v505
	goto L124
L182:
	;
	v745 = F_errsave_start(m, v29)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L6
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v494) {
		goto L191
	} else {
		goto L192
	}
L185:
	;
	if v745 == int32(0) {
		goto L119
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+368)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(368))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+352)) = int32(125)
	F_errdetail(m, int32(_a_F_array_in_17), v27+int32(352))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(670), int32(_a_F_array_in_18))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v1601 = int32(0)
	goto L15
L191:
	;
	v774 = v736 - int32(8)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = v775 + int32(1)
	goto L193
L192:
	;
	goto L193
L193:
	;
	v780 = int32(1)
	v781 = v523 + v780
	v785 = v494 - v780
	v788 = v27 + int32(496) + v785<<(uint(int32(2))%32)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	if v789 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = v739
	v1175 = v781
	v1179 = v785
	v1182 = v497
	v1183 = v498
	v1184 = int32(1)
	v1185 = v500
	v1187 = v502
	v1188 = v503
	v1190 = v505
	goto L124
L195:
	;
	goto L196
L196:
	;
	if v789 != v739 {
		goto L120
	} else {
		goto L197
	}
L197:
	;
	v1175 = v781
	v1179 = v785
	v1182 = v497
	v1183 = v498
	v1184 = int32(1)
	v1185 = v500
	v1187 = v502
	v1188 = v503
	v1190 = v505
	goto L124
L198:
	;
	goto L200
L199:
	;
	goto L130
L200:
	;
	if base.B2i32(v796 == int32(32))|base.B2i32(base.Ui32((v796-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	goto L199
L202:
	;
	goto L203
L203:
	;
	v523 = v523 + int32(1)
	goto L129
L204:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813))))
	if base.Ui32(v837) <= base.Ui32(int32(91)) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v27)+560))
	v935 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v933+v820))) = uint8(v935)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+564)) = v820
	v939 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_array_in[0])))
	if v939&v824 == v935 {
		goto L126
	} else {
		goto L238
	}
L206:
	;
	v908 = base.I32_extend8_s(v837)
	if base.B2i32(v908 == v78)|base.B2i32(v908 == int32(125)) == int32(0) {
		goto L230
	} else {
		goto L231
	}
L207:
	;
	if v837 == int32(0) {
		goto L128
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	if v837 != int32(92) {
		goto L218
	} else {
		goto L219
	}
L210:
	;
	if v837 != int32(34) {
		goto L206
	} else {
		goto L211
	}
L211:
	;
	v844 = F_errsave_start(m, v29)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L6
	} else {
		goto L212
	}
L212:
	;
	if v844 == int32(0) {
		goto L119
	} else {
		goto L213
	}
L213:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(288))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	F_errdetail(m, int32(_a_F_array_in_15), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(905), int32(_a_F_array_in_16))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
	;
	v1601 = int32(0)
	goto L15
L218:
	;
	if v837 != int32(123) {
		goto L206
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v897 = int32(*(*int8)(unsafe.Add(mBase, uint32(v813)+1)))
	if v897 == int32(0) {
		goto L128
	} else {
		goto L228
	}
L221:
	;
	v871 = F_errsave_start(m, v29)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L6
	} else {
		goto L222
	}
L222:
	;
	if v871 == int32(0) {
		goto L119
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(272))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = int32(123)
	F_errdetail(m, int32(_a_F_array_in_17), v27+int32(256))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(899), int32(_a_F_array_in_16))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	v1601 = int32(0)
	goto L15
L228:
	;
	F_appendStringInfoChar(m, v27+int32(560), v897)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v27)+564))
	v813 = v813 + int32(2)
	v820 = v907
	v824 = int32(0)
	goto L204
L230:
	;
	F_appendStringInfoChar(m, v27+int32(560), v908)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L6
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	goto L205
L233:
	;
	v919 = int32(*(*int8)(unsafe.Add(mBase, uint32(v813))))
	goto L234
L234:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v27)+564))
	if base.B2i32(v919 == int32(32))|base.B2i32(base.Ui32((v919-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v930 = v820
	goto L237
L236:
	;
	v930 = v929
	goto L237
L237:
	;
	v813 = v813 + int32(1)
	v820 = v930
	goto L204
L238:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v27)+560))
	v947 = v943
	v948 = int32(_a_F_array_in_19)
	goto L240
L239:
	;
	if v985 != 0 {
		goto L126
	} else {
		goto L252
	}
L240:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947))))
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
	if v951 == v952 {
		v974 = v951
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v985 = int32(0)
	goto L239
L242:
	;
	v976 = int32(1)
	if v974 != 0 {
		v947 = v947 + v976
		v948 = v948 + v976
		goto L240
	} else {
		goto L251
	}
L243:
	;
	if base.Ui32((v951-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v962 = v951 | int32(32)
	goto L246
L245:
	;
	v962 = v951
	goto L246
L246:
	;
	if base.Ui32((v952-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v971 = v952 | int32(32)
	goto L249
L248:
	;
	v971 = v952
	goto L249
L249:
	;
	if v962 == v971 {
		v974 = v962
		goto L242
	} else {
		goto L250
	}
L250:
	;
	v985 = v962 - v971
	goto L239
L251:
	;
	goto L241
L252:
	;
	v1063 = v813
	v1070 = int32(1)
	goto L125
L253:
	;
	if v1011 == int32(0) {
		goto L119
	} else {
		goto L254
	}
L254:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L6
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(240))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L6
	} else {
		goto L256
	}
L256:
	;
	F_errdetail(m, int32(_a_F_array_in_20), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L6
	} else {
		goto L257
	}
L257:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(942), int32(_a_F_array_in_16))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L6
	} else {
		goto L258
	}
L258:
	;
	v1601 = int32(0)
	goto L15
L259:
	;
	v1175 = v523 + int32(1)
	v1179 = v494
	v1182 = v497
	v1183 = v498
	v1184 = int32(0)
	v1185 = v500
	v1187 = v502
	v1188 = v503
	v1190 = v505
	goto L124
L260:
	;
	goto L261
L261:
	;
	v1037 = F_errsave_start(m, v29)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L6
	} else {
		goto L262
	}
L262:
	;
	if v1037 == int32(0) {
		goto L119
	} else {
		goto L263
	}
L263:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L6
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(160))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L6
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v78
	F_errdetail(m, int32(_a_F_array_in_17), v27+int32(144))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L6
	} else {
		goto L266
	}
L266:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(705), int32(_a_F_array_in_18))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	v1601 = int32(0)
	goto L15
L268:
	;
	v1087 = F_errsave_start(m, v29)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L6
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	if v505 <= v500 {
		goto L277
	} else {
		goto L278
	}
L271:
	;
	if v1087 == int32(0) {
		goto L119
	} else {
		goto L272
	}
L272:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L6
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(176))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L6
	} else {
		goto L274
	}
L274:
	;
	F_errdetail(m, int32(_a_F_array_in_21), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L6
	} else {
		goto L275
	}
L275:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(719), int32(_a_F_array_in_18))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L6
	} else {
		goto L276
	}
L276:
	;
	v1601 = int32(0)
	goto L15
L277:
	;
	if base.Ui32(int32(268435455)) <= base.Ui32(v505) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	v1146 = v502
	v1147 = v503
	v1148 = v505
	goto L279
L279:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v27)+560))
	if v1070 != 0 {
		goto L293
	} else {
		goto L294
	}
L280:
	;
	v1113 = F_errsave_start(m, v29)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L6
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1133 = int32(268435455)
	v1135 = v505 << (uint(int32(1)) % 32)
	if base.Ui32(v1133) <= base.Ui32(v1135) {
		goto L288
	} else {
		goto L289
	}
L283:
	;
	if v1113 == int32(0) {
		goto L119
	} else {
		goto L284
	}
L284:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L6
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_in_10), v27+int32(224))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L6
	} else {
		goto L286
	}
L286:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(728), int32(_a_F_array_in_18))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L6
	} else {
		goto L287
	}
L287:
	;
	v1601 = int32(0)
	goto L15
L288:
	;
	v1138 = v1133
	goto L290
L289:
	;
	v1138 = v1135
	goto L290
L290:
	;
	v1141 = F_repalloc(m, v502, v1138<<(uint(int32(2))%32))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L6
	} else {
		goto L291
	}
L291:
	;
	v1143 = F_repalloc(m, v503, v1138)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	v1146 = v1141
	v1147 = v1143
	v1148 = v1138
	goto L279
L293:
	;
	v1151 = int32(0)
	goto L295
L294:
	;
	v1151 = v1150
	goto L295
L295:
	;
	v1155 = F_InputFunctionCallSafe(m, v76+int32(20), v1151, v77, v30, v29, v1146+v500<<(uint(int32(2))%32))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L6
	} else {
		goto L296
	}
L296:
	;
	if v1155 == int32(0) {
		goto L119
	} else {
		goto L297
	}
L297:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v500+v1147))) = uint8(v1070)
	if v494 != v498 {
		goto L120
	} else {
		goto L298
	}
L298:
	;
	v1162 = int32(1)
	v1167 = v494<<(uint(int32(2))%32) + v27 + int32(524)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1167)))
	*(*int32)(unsafe.Add(mBase, uint32(v1167))) = v1168 + v1162
	v1175 = v1063
	v1179 = v494
	v1182 = v1162
	v1183 = v494
	v1184 = v1162
	v1185 = v500 + v1162
	v1187 = v1146
	v1188 = v1147
	v1190 = v1148
	goto L124
L299:
	;
	goto L122
L300:
	;
	v1204 = v1175
	goto L301
L301:
	;
	v1228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1204))))
	if v1228 == int32(0) {
		goto L118
	} else {
		goto L303
	}
L302:
	;
	v1242 = int32(0)
	v1243 = F_errsave_start(m, v29)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L6
	} else {
		goto L306
	}
L303:
	;
	goto L304
L304:
	;
	if base.B2i32(v1228 == int32(32))|base.B2i32(base.Ui32((v1228-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v1204 = v1204 + int32(1)
		goto L301
	} else {
		goto L305
	}
L305:
	;
	goto L302
L306:
	;
	if v1243 == int32(0) {
		v1601 = v1242
		goto L15
	} else {
		goto L307
	}
L307:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L6
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(128))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L6
	} else {
		goto L309
	}
L309:
	;
	F_errdetail(m, int32(_a_F_array_in_22), int32(0))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L6
	} else {
		goto L310
	}
L310:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(308), int32(_a_F_array_in_12))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L6
	} else {
		goto L311
	}
L311:
	;
	v1601 = v1242
	goto L15
L312:
	;
	if v102 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	if v1289 == int32(0) {
		goto L119
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	if v1289 == int32(0) {
		goto L119
	} else {
		goto L321
	}
L316:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L6
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(192))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L6
	} else {
		goto L318
	}
L318:
	;
	F_errdetail(m, int32(_a_F_array_in_23), int32(0))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L6
	} else {
		goto L319
	}
L319:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(778), int32(_a_F_array_in_18))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L6
	} else {
		goto L320
	}
L320:
	;
	v1601 = int32(0)
	goto L15
L321:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L6
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v32
	F_errmsg(m, int32(_a_F_array_in_3), v27+int32(208))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L6
	} else {
		goto L323
	}
L323:
	;
	F_errdetail(m, int32(_a_F_array_in_24), int32(0))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L6
	} else {
		goto L324
	}
L324:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(783), int32(_a_F_array_in_18))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L6
	} else {
		goto L325
	}
L325:
	;
	goto L119
L326:
	;
	v1357 = int32(0)
	if v1185 <= v1357 {
		v1497 = v1357
		goto L330
	} else {
		goto L331
	}
L327:
	;
	goto L328
L328:
	;
	v1586 = F_palloc0(m, int32(16))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L6
	} else {
		goto L376
	}
L329:
	;
	v1550 = v1549 + v1527
	v1551 = F_palloc0(m, v1550)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L6
	} else {
		goto L366
	}
L330:
	;
	v1527 = v1497
	v1541 = v1357
	v1549 = (v1183<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L329
L331:
	;
	v1365 = int32(0)
	v1367 = v1365
	v1369 = v1357
	v1371 = v1365
	goto L332
L332:
	;
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367+v1188))))
	if v1392 != 0 {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	if v1474&int32(1) == int32(0) {
		v1497 = v1473
		goto L330
	} else {
		goto L365
	}
L334:
	;
	v1478 = v1367 + int32(1)
	if v1478 != v1185 {
		v1367 = v1478
		v1369 = v1473
		v1371 = v1474
		goto L332
	} else {
		goto L364
	}
L335:
	;
	v1473 = v1369
	v1474 = int32(1)
	goto L334
L336:
	;
	goto L337
L337:
	;
	if base.B2i32(v80 == int32(-1)) == int32(0) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v1437 = v1369 + v1436
	switch v81 - int32(99) {
	case 0:
		v1450 = v1437
		goto L354
	case 1:
		goto L356
	default:
		goto L355
	case 6:
		goto L357
	}
L339:
	;
	if int32(0) < v80 {
		v1436 = v80
		goto L338
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1407 = v1187 + v1367<<(uint(int32(2))%32)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	v1409 = F_pg_detoast_datum(m, v1408)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L6
	} else {
		goto L343
	}
L342:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1187+v1367<<(uint(int32(2))%32))))
	v1402 = F_strlen(m, v1401)
	mBase = m.M
	v1436 = v1402 + int32(1)
	goto L338
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1407))) = v1409
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409))))
	if v1412 == int32(1) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409)+1)))
	if base.Ui32((v1416-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v1436 = int32(6)
		goto L338
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	if v1412&int32(1) != 0 {
		goto L351
	} else {
		goto L352
	}
L347:
	;
	v1423 = int32(18)
	if v1416 == v1423 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1427 = v1423
	goto L350
L349:
	;
	v1427 = int32(2)
	goto L350
L350:
	;
	v1436 = v1427
	goto L338
L351:
	;
	v1436 = int32(base.Ui32(v1412) >> (uint(int32(1)) % 32))
	goto L338
L352:
	;
	goto L353
L353:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	v1436 = int32(base.Ui32(v1432) >> (uint(int32(2)) % 32))
	goto L338
L354:
	;
	if base.Ui32(v1450) < base.Ui32(int32(1073741824)) {
		v1473 = v1450
		v1474 = v1371
		goto L334
	} else {
		goto L358
	}
L355:
	;
	v1450 = (v1437 + int32(1)) & int32(-2)
	goto L354
L356:
	;
	v1450 = (v1437 + int32(7)) & int32(-8)
	goto L354
L357:
	;
	v1450 = (v1437 + int32(3)) & int32(-4)
	goto L354
L358:
	;
	v1453 = int32(0)
	v1454 = F_errsave_start(m, v29)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L6
	} else {
		goto L359
	}
L359:
	;
	if v1454 == int32(0) {
		v1601 = v1453
		goto L15
	} else {
		goto L360
	}
L360:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L6
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_in_10), v27+int32(112))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L6
	} else {
		goto L362
	}
L362:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_0), int32(336), int32(_a_F_array_in_12))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L6
	} else {
		goto L363
	}
L363:
	;
	v1601 = v1453
	goto L15
L364:
	;
	goto L333
L365:
	;
	v1487 = base.I32_div_s(v1185+int32(7), int32(8))
	v1494 = (v1487 + v1183<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v1527 = v1473
	v1541 = v1494
	v1549 = v1494
	goto L329
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+8)) = v1541
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+4)) = v1183
	v1556 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1551))) = v1550 << (uint(v1556) % 32)
	v1560 = v1551 + int32(16)
	v1562 = v1183 << (uint(v1556) % 32)
	v1563 = int32(0)
	v1564 = base.B2i32(v1562 == v1563)
	if v1564 == v1563 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	base.MemoryCopy(m, v1560, v27+int32(496), v1562)
	goto L369
L368:
	;
	goto L369
L369:
	;
	if v1564 == int32(0) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	base.MemoryCopy(m, v1560+v1562, v27+int32(464), v1562)
	goto L372
L371:
	;
	goto L372
L372:
	;
	v1576 = int32(1)
	F_CopyArrayEls(m, v1551, v1187, v1188, v1185, v80, v79&v1576, base.I32_extend8_s(v81), v1576)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L6
	} else {
		goto L373
	}
L373:
	;
	F_pfree(m, v1187)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L6
	} else {
		goto L374
	}
L374:
	;
	F_pfree(m, v1188)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L6
	} else {
		goto L375
	}
L375:
	;
	v1601 = v1551
	goto L15
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1586))) = int64(64)
	v1601 = v1586
	goto L15
}
func F_array_ndims(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_DatumGetAnyArrayP(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v10 == int32(-1) {
			v13 = int32(28)
		} else {
			v13 = int32(4)
		}
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4+v13)))
		if base.Ui32(v15-int32(7)) <= base.Ui32(int32(-7)) {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v23 = int32(0)
		} else {
			v23 = v15
		}
		return v23
	}
}
func F_array_ne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2 ^ int32(1)
	}
}
func F_array_set_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
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
	var v148 int32
	_ = v148
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1040 int32
	_ = v1040
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1565 int32
	_ = v1565
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1641 int32
	_ = v1641
	var v1648 int32
	_ = v1648
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1754 int32
	_ = v1754
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1893 int32
	_ = v1893
	var v1900 int32
	_ = v1900
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2173 int32
	_ = v2173
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2211 int32
	_ = v2211
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2254 int32
	_ = v2254
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	v5 = l4
	v10 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(128)
	m.G0 = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+59)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = l3
	if v10 < l5 {
		goto L20
	} else {
		goto L21
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L27
	} else {
		goto L465
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L27
	} else {
		goto L461
	}
L3:
	;
	m.G0 = v31 + int32(128)
	return v2254
L4:
	;
	v1664 = v31 + int32(96)
	v1666 = v31 - int32(-64)
	v1667 = int32(0)
	v1676 = l1 - int32(1)
	if v1676 < v1667 {
		v1754 = v1667
		goto L373
	} else {
		goto L374
	}
L5:
	;
	v1626 = v31 + int32(96)
	v1627 = F_ArrayGetNItemsSafe(m, l1, v1626)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L27
	} else {
		goto L370
	}
L6:
	;
	v1593 = int32(0)
	if v1578 == v1593 {
		v1634 = v1565
		v1641 = v1593
		v1648 = v1579
		v1662 = v1593
		goto L4
	} else {
		goto L369
	}
L7:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	if v1521 <= v1520 {
		goto L360
	} else {
		goto L361
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1497))) = uint8(v1492)
	v2254 = v924
	goto L3
L9:
	;
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	v1421 = v993
	v1422 = int32(1)
	v1423 = v1418
	v1425 = v989
	v1426 = v992
	v1432 = v671
	v1435 = v688
	goto L343
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L27
	} else {
		goto L339
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L27
	} else {
		goto L335
	}
L12:
	;
	v645 = v31 + int32(96)
	v646 = F_ArrayGetNItemsSafe(m, l1, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L27
	} else {
		goto L152
	}
L13:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	if v570 <= v569 {
		goto L141
	} else {
		goto L142
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L27
	} else {
		goto L136
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L27
	} else {
		goto L132
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L27
	} else {
		goto L128
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L27
	} else {
		goto L124
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L27
	} else {
		goto L120
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L27
	} else {
		goto L116
	}
L20:
	;
	if l1 != int32(1) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(l1-int32(7)) <= base.Ui32(int32(-7)) {
		goto L16
	} else {
		goto L33
	}
L23:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v39 < int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v42 = base.I32_div_s(l5, l6)
	if v42 <= v39 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	if v5 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v44 = F_palloc(m, l5)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	if l5 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	base.MemoryCopy(m, v44, l0, l5)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v52 = F_ArrayCastAndSet(m, l3, l6, l7, l8, v44+v49*l6)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v2254 = v44
	goto L3
L33:
	;
	if v5|base.B2i32(l6 != int32(-1)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v63 = F_pg_detoast_datum(m, l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L27
	} else {
		goto L37
	}
L35:
	;
	v66 = l3
	goto L36
L36:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v67 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v63
	v66 = v63
	goto L36
L38:
	;
	v211 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L27
	} else {
		goto L77
	}
L39:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v70&int32(254) != int32(2) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v75 = F_DatumGetExpandedArray(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v79 = v77 << (uint(int32(2)) % 32)
	v80 = int32(0)
	v81 = base.B2i32(v79 == v80)
	if v81 == v80 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	base.MemoryCopy(m, v31+int32(96), v86, v79)
	goto L44
L43:
	;
	goto L44
L44:
	;
	if v81 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
	base.MemoryCopy(m, v31-int32(-64), v92, v79)
	goto L47
L46:
	;
	goto L47
L47:
	;
	if v77 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	F_deconstruct_expanded_array(m, v75)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L27
	} else {
		goto L59
	}
L49:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v98 = l1 << (uint(int32(2)) % 32)
	v99 = F_MemoryContextAllocZero(m, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L27
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l1 != v77 {
		goto L15
	} else {
		goto L58
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v103 = F_MemoryContextAllocZero(m, v102, v98)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L27
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+36)) = v103
	v106 = int32(0)
	v107 = base.B2i32(v98 == v106)
	if v107 == v106 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	base.MemoryFill(m, v31+int32(96), int32(0), v98)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if v98 == v106 {
		goto L48
	} else {
		goto L57
	}
L57:
	;
	base.MemoryCopy(m, v31-int32(-64), l2, v98)
	goto L48
L58:
	;
	goto L48
L59:
	;
	if v5 != 0 {
		v137 = v66
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v138 = int32(0)
	v139 = base.B2i32(v77 == v138)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v75)+52))
	v144 = v5 | base.B2i32(v141 != v138)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	if l1 == int32(1) {
		goto L7
	} else {
		goto L64
	}
L61:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+46)))
	if v122&int32(1) != 0 {
		v137 = v66
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v125 = int32(_a_F_array_set_element_0)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v128
	v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+44)))
	v132 = F_datumCopy(m, v66, int32(0), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L27
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v126
	v137 = v132
	goto L60
L64:
	;
	v148 = v138
	goto L65
L65:
	;
	v177 = v148 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l2+v177)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(-64)+v177)))
	if v183 <= v179 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v1565 = int32(0)
	v1578 = v139
	v1579 = v144
	goto L6
L67:
	;
	v208 = v148 + int32(1)
	if v208 != l1 {
		v148 = v208
		goto L65
	} else {
		goto L76
	}
L68:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(96)+v177)))
	if v179 < v188+v183 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L27
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L27
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L27
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2653), int32(_a_F_array_set_element_3))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L27
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
	goto L66
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v213 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v218 = l1 << (uint(int32(2)) % 32)
	if v218 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if v213 != l1 {
		goto L14
	} else {
		goto L96
	}
L81:
	;
	base.MemoryCopy(m, v31-int32(-64), l2, v218)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v222 = int32(0)
	if base.Ui32(int32(7)) <= base.Ui32(l1-int32(1)) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v380 = F_construct_md_array(m, v31+int32(60), v31+int32(59), l1, v31+int32(96), v31-int32(-64), v216, l6, l7, l8)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L27
	} else {
		goto L95
	}
L85:
	;
	v229 = int32(0)
	v234 = v222
	goto L88
L86:
	;
	v281 = v222
	goto L87
L87:
	;
	v306 = v222
	v309 = v281
	goto L92
L88:
	;
	v261 = v31 + int32(96) + v234<<(uint(int32(2))%32)
	v262 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v261)+24)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v261)+16)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v262
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v262
	v270 = int32(8)
	v271 = v234 + v270
	v273 = v229 + v270
	if v273 != 0 {
		v229 = v273
		v234 = v271
		goto L88
	} else {
		goto L90
	}
L89:
	;
	if l1 == int32(0) {
		goto L84
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v281 = v271
	goto L87
L92:
	;
	v337 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(96)+v309<<(uint(int32(2))%32)))) = v337
	v342 = v306 + v337
	if v342 != l1 {
		v306 = v342
		v309 = v309 + v337
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L84
L94:
	;
	goto L93
L95:
	;
	v2254 = v380
	goto L3
L96:
	;
	v384 = v211 + int32(16)
	v386 = l1 << (uint(int32(2)) % 32)
	v387 = int32(0)
	v388 = base.B2i32(v386 == v387)
	if v388 == v387 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	base.MemoryCopy(m, v31+int32(96), v384, v386)
	goto L99
L98:
	;
	goto L99
L99:
	;
	if v388 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	base.MemoryCopy(m, v31-int32(-64), v384+v398<<(uint(int32(2))%32), v386)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v403 = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v407 = v5 | base.B2i32(v404 != v403)
	if l1 == int32(1) {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	v415 = v403
	goto L104
L104:
	;
	v439 = v415 << (uint(int32(2)) % 32)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l2+v439)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(-64)+v439)))
	if v445 <= v441 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v616 = v469
	v634 = v10
	v635 = v407
	goto L12
L106:
	;
	v469 = int32(1)
	v471 = v415 + v469
	if v471 != l1 {
		v415 = v471
		goto L104
	} else {
		goto L115
	}
L107:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(96)+v439)))
	if v441 < v450+v445 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L27
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L27
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L27
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2374), int32(_a_F_array_set_element_4))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L27
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
	goto L105
L116:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L27
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L27
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2245), int32(_a_F_array_set_element_4))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L27
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L27
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_array_set_element_1), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L27
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2250), int32(_a_F_array_set_element_4))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L27
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L27
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_array_set_element_6), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L27
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2255), int32(_a_F_array_set_element_4))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L27
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L27
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L27
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2267), int32(_a_F_array_set_element_4))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L27
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L27
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L27
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2570), int32(_a_F_array_set_element_3))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L27
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L27
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_array_set_element_5), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L27
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2316), int32(_a_F_array_set_element_4))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L27
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v594 = v589 + v590
	if v569 < v594 {
		goto L146
	} else {
		goto L147
	}
L141:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v589 = v572
	v590 = v570
	v592 = v10
	v593 = v407
	goto L140
L142:
	;
	goto L143
L143:
	;
	v573 = v570 - v569
	if base.B2i32(v573 < v570)^base.B2i32(int32(0) < v569) != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v579 = v578 + v573
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v579
	if base.B2i32(v573 < int32(0)) != base.B2i32(v579 < v578) {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v569
	v589 = v579
	v590 = v569
	v592 = v573
	v593 = base.B2i32(int32(1) < v573) | v407
	goto L140
L146:
	;
	v616 = int32(1)
	v634 = v592
	v635 = v593
	goto L12
L147:
	;
	goto L148
L148:
	;
	v599 = v569 - v594
	if base.B2i32(int32(0) < v594)^base.B2i32(v599 < v569) != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	v603 = v599 + int32(1)
	if v603 < v599 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	v605 = v589 + v603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v605
	if base.B2i32(v603 < int32(0)) != base.B2i32(v605 < v589) {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	v616 = base.B2i32(v603 == int32(0))
	v634 = v592
	v635 = base.B2i32(int32(1) < v603) | v593
	goto L12
L152:
	;
	F_ArrayCheckBounds(m, l1, v645, v31-int32(-64))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L27
	} else {
		goto L153
	}
L153:
	;
	v653 = l1 << (uint(int32(3)) % 32)
	if v635&int32(1) != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v671 = F_ArrayGetNItemsSafe(m, l1, v384)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L27
	} else {
		goto L158
	}
L155:
	;
	v659 = base.I32_div_s(v646+int32(7), int32(8))
	v664 = (v653 + v659 + int32(23)) & int32(-8)
	v669 = v664
	v670 = v664
	goto L154
L156:
	;
	goto L157
L157:
	;
	v669 = v10
	v670 = (v653 + int32(23)) & int32(120)
	goto L154
L158:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v679 = v677 << (uint(int32(3)) % 32)
	if v676 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v684 = v676
	goto L161
L160:
	;
	v684 = (v679 + int32(23)) & int32(-8)
	goto L161
L161:
	;
	v685 = int32(base.Ui32(v673)>>(uint(int32(2))%32)) - v684
	v686 = v679 + v384
	if v676 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v688 = v686
	goto L164
L163:
	;
	v688 = int32(0)
	goto L164
L164:
	;
	if v634 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if v5 != 0 {
		v920 = v10
		goto L208
	} else {
		goto L209
	}
L166:
	;
	v868 = int32(0)
	v869 = v10
	v871 = v685
	v872 = v10
	goto L165
L167:
	;
	goto L168
L168:
	;
	if v616 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v868 = v685
	v869 = v671
	v871 = int32(0)
	v872 = v10
	goto L165
L170:
	;
	goto L171
L171:
	;
	v694 = v31 + int32(96)
	v696 = v31 - int32(-64)
	v697 = int32(0)
	v706 = l1 - int32(1)
	if v706 < v697 {
		v784 = v697
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if v790 != 0 {
		goto L182
	} else {
		goto L183
	}
L173:
	;
	goto L172
L174:
	;
	v709 = int32(1)
	if v706 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v718 = v706
	v719 = v709
	v720 = v697
	v725 = v697
	goto L178
L176:
	;
	v761 = v706
	v762 = v709
	v763 = v697
	goto L177
L177:
	;
	v770 = v761 << (uint(int32(2)) % 32)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l2+v770)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v770+v696)))
	v784 = (v772-v774)*v762 + v763
	goto L173
L178:
	;
	v726 = int32(2)
	v727 = v718 << (uint(v726) % 32)
	v729 = v727 - int32(4)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l2+v729)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v696+v729)))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v727+v694)))
	v737 = v736 * v719
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v727+l2)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v727+v696)))
	v746 = (v731-v733)*v737 + ((v740-v742)*v719 + v720)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v694+v729)))
	v749 = v748 * v737
	v751 = v718 - v726
	v753 = v725 + v726
	if v753 != l1&int32(-2) {
		v718 = v751
		v719 = v749
		v720 = v746
		v725 = v753
		goto L178
	} else {
		goto L180
	}
L179:
	;
	if l1&int32(1) == int32(0) {
		v784 = v746
		goto L173
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	v761 = v751
	v762 = v749
	v763 = v746
	goto L177
L182:
	;
	v798 = v790
	goto L184
L183:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v798 = (v791<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L184
L184:
	;
	v799 = v798 + v211
	v801 = F_array_seek(m, v799, int32(0), v688, v784, l6, l8)
	mBase = m.M
	v802 = v801 - v799
	if v676 != 0 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v868 = v802
	v869 = v784
	v871 = v685 - (v802 + v864)
	v872 = v864
	goto L165
L186:
	;
	v804 = base.I32_div_s(v784, int32(8))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686+v804))))
	if int32(base.Ui32(v806)>>(uint(v784&int32(7))%32))&int32(1) == int32(0) {
		v864 = v10
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if int32(0) < l6 {
		v847 = l6
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L188
L190:
	;
	switch l8 - int32(99) {
	case 0:
		v864 = v847
		goto L185
	case 1:
		goto L206
	default:
		goto L205
	case 6:
		goto L207
	}
L191:
	;
	if l6 == int32(-1) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	if v818 == int32(1) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v842 = F_strlen(m, v801)
	mBase = m.M
	v847 = v842 + int32(1)
	goto L190
L195:
	;
	v822 = int32(18)
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+1)))
	if v824 == v822 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	v835 = int32(1)
	if v818&v835 != 0 {
		v847 = int32(base.Ui32(v818) >> (uint(v835) % 32))
		goto L190
	} else {
		goto L204
	}
L198:
	;
	v827 = v822
	goto L200
L199:
	;
	v827 = int32(2)
	goto L200
L200:
	;
	if base.Ui32((v824-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v834 = int32(6)
	goto L203
L202:
	;
	v834 = v827
	goto L203
L203:
	;
	v847 = v834
	goto L190
L204:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	v847 = int32(base.Ui32(v839) >> (uint(int32(2)) % 32))
	goto L190
L205:
	;
	v864 = (v847 + int32(1)) & int32(-2)
	goto L185
L206:
	;
	v864 = (v847 + int32(7)) & int32(-8)
	goto L185
L207:
	;
	v864 = (v847 + int32(3)) & int32(-4)
	goto L185
L208:
	;
	v923 = v868 + v670 + v871 + v920
	v924 = F_palloc0(m, v923)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L27
	} else {
		goto L226
	}
L209:
	;
	if int32(0) < l6 {
		v904 = l6
		goto L210
	} else {
		goto L211
	}
L210:
	;
	switch l8 - int32(99) {
	case 0:
		v920 = v904
		goto L208
	case 1:
		goto L224
	default:
		goto L223
	case 6:
		goto L225
	}
L211:
	;
	if l6 == int32(-1) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v877 == int32(1) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v900 = F_strlen(m, v66)
	mBase = m.M
	v904 = v900 + int32(1)
	goto L210
L215:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if base.Ui32((v881-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v904 = int32(6)
		goto L210
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v893 = int32(1)
	if v877&v893 != 0 {
		v904 = int32(base.Ui32(v877) >> (uint(v893) % 32))
		goto L210
	} else {
		goto L222
	}
L218:
	;
	v888 = int32(18)
	if v881 == v888 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v892 = v888
	goto L221
L220:
	;
	v892 = int32(2)
	goto L221
L221:
	;
	v904 = v892
	goto L210
L222:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v904 = int32(base.Ui32(v897) >> (uint(int32(2)) % 32))
	goto L210
L223:
	;
	v920 = (v904 + int32(1)) & int32(-2)
	goto L208
L224:
	;
	v920 = (v904 + int32(7)) & int32(-8)
	goto L208
L225:
	;
	v920 = (v904 + int32(3)) & int32(-4)
	goto L208
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v924)+8)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v924)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = v923 << (uint(int32(2)) % 32)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v924)+12)) = v931
	v934 = v924 + int32(16)
	if v388 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	base.MemoryCopy(m, v934, v31+int32(96), v386)
	goto L229
L228:
	;
	goto L229
L229:
	;
	if v388 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	base.MemoryCopy(m, v934+v386, v31-int32(-64), v386)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v946 = v211 + v684
	v947 = v924 + v670
	if v868 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	base.MemoryCopy(m, v947, v946, v868)
	goto L235
L234:
	;
	goto L235
L235:
	;
	if v5 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v952 = F_ArrayCastAndSet(m, v66, l6, l7, l8, v947+v868)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L27
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	if v871 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L238
L240:
	;
	base.MemoryCopy(m, v947+v868+v920, v868+v946+v872, v871)
	goto L242
L241:
	;
	goto L242
L242:
	;
	if v635&int32(1) == int32(0) {
		v2254 = v924
		goto L3
	} else {
		goto L243
	}
L243:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v924)+4))
	v966 = v934 + v963<<(uint(int32(3))%32)
	if v616 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v969 = v869
	goto L246
L245:
	;
	v969 = v646 - int32(1)
	goto L246
L246:
	;
	v971 = base.I32_div_s(v969, int32(8))
	v972 = v966 + v971
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	v976 = v969 & int32(7)
	if v5 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v982 = v973 & base.I32_rotl(int32(-2), v976)
	goto L249
L248:
	;
	v982 = v973 | int32(1)<<(uint(v976)%32)
	goto L249
L249:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v972))) = uint8(v982)
	if v634 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	if v671 <= int32(0) {
		v2254 = v924
		goto L3
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if v869 <= int32(0) {
		goto L262
	} else {
		goto L263
	}
L253:
	;
	v989 = int32(1) << (uint(v634&int32(7)) % 32)
	v991 = base.I32_div_s(v634, int32(8))
	v992 = v966 + v991
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992))))
	if v676 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	v995 = v993
	v999 = v989
	v1000 = v992
	v1006 = v671
	goto L255
L255:
	;
	v1022 = v995 | v999
	v1023 = int32(1)
	v1024 = v1006 - v1023
	v1026 = v999 << (uint(v1023) % 32)
	if v1026 == int32(256) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1492 = v1022
	v1497 = v1000
	goto L8
L257:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1000))) = uint8(v1022)
	if v1024 == int32(0) {
		v2254 = v924
		goto L3
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1006) {
		v995 = v1022
		v999 = v1026
		v1006 = v1024
		goto L255
	} else {
		goto L261
	}
L260:
	;
	v1032 = int32(1)
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000)+1)))
	v995 = v1033
	v999 = v1032
	v1000 = v1000 + v1032
	v1006 = v1024
	goto L255
L261:
	;
	goto L256
L262:
	;
	if v616 == int32(0) {
		v2254 = v924
		goto L3
	} else {
		goto L303
	}
L263:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966))))
	if v676 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1190))) = uint8(v1188)
	goto L262
L265:
	;
	v1122 = v1040
	v1123 = v966
	v1126 = v869
	goto L293
L266:
	;
	if v869 != int32(1) {
		goto L265
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	v1048 = int32(1)
	v1051 = v1040
	v1052 = v1048
	v1053 = v966
	v1055 = v1048
	v1056 = v869
	v1057 = v1047
	v1063 = v688
	goto L270
L269:
	;
	v1188 = v1040 | int32(1)
	v1190 = v966
	goto L264
L270:
	;
	if v1052&v1057 != 0 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	if v1099 != int32(1) {
		v1188 = v1097
		v1190 = v1098
		goto L264
	} else {
		goto L285
	}
L272:
	;
	v1083 = v1051 | v1055
	goto L274
L273:
	;
	v1083 = v1051 & (v1055 ^ int32(-1))
	goto L274
L274:
	;
	v1084 = int32(1)
	v1085 = v1056 - v1084
	v1087 = v1055 << (uint(v1084) % 32)
	if v1087 == int32(256) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1053))) = uint8(v1083)
	if v1085 == int32(0) {
		goto L262
	} else {
		goto L278
	}
L276:
	;
	v1097 = v1083
	v1098 = v1053
	v1099 = v1087
	goto L277
L277:
	;
	v1101 = v1052 << (uint(int32(1)) % 32)
	if v1101 == int32(256) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	v1093 = int32(1)
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053)+1)))
	v1097 = v1094
	v1098 = v1053 + v1093
	v1099 = v1093
	goto L277
L279:
	;
	goto L271
L280:
	;
	if v1085 == int32(0) {
		goto L279
	} else {
		goto L283
	}
L281:
	;
	v1110 = v1101
	v1111 = v1057
	v1112 = v1063
	goto L282
L282:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1056) {
		v1051 = v1097
		v1052 = v1110
		v1053 = v1098
		v1055 = v1099
		v1056 = v1085
		v1057 = v1111
		v1063 = v1112
		goto L270
	} else {
		goto L284
	}
L283:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+1)))
	v1107 = int32(1)
	v1110 = v1107
	v1111 = v1106
	v1112 = v1063 + v1107
	goto L282
L284:
	;
	goto L279
L285:
	;
	goto L262
L286:
	;
	v1188 = v1186
	v1190 = v1123
	goto L264
L287:
	;
	v1186 = v1122 | int32(3)
	goto L286
L288:
	;
	v1186 = v1122 | int32(7)
	goto L286
L289:
	;
	v1186 = v1122 | int32(15)
	goto L286
L290:
	;
	v1186 = v1122 | int32(31)
	goto L286
L291:
	;
	v1186 = v1122 | int32(63)
	goto L286
L292:
	;
	v1186 = v1122 | int32(127)
	goto L286
L293:
	;
	if v1126 < int32(3) {
		goto L287
	} else {
		goto L295
	}
L294:
	;
	v1188 = v1167 | int32(1)
	v1190 = v1169
	goto L264
L295:
	;
	if v1126 == int32(3) {
		goto L288
	} else {
		goto L296
	}
L296:
	;
	if base.Ui32(v1126) < base.Ui32(int32(5)) {
		goto L289
	} else {
		goto L297
	}
L297:
	;
	if v1126 == int32(5) {
		goto L290
	} else {
		goto L298
	}
L298:
	;
	if base.Ui32(v1126) < base.Ui32(int32(7)) {
		goto L291
	} else {
		goto L299
	}
L299:
	;
	if v1126 == int32(7) {
		goto L292
	} else {
		goto L300
	}
L300:
	;
	v1161 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1123))) = uint8(v1161)
	v1164 = v1126 - int32(8)
	if v1164 == int32(0) {
		goto L262
	} else {
		goto L301
	}
L301:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123)+1)))
	v1168 = int32(1)
	v1169 = v1123 + v1168
	if v1164 != v1168 {
		v1122 = v1167
		v1123 = v1169
		v1126 = v1164
		goto L293
	} else {
		goto L302
	}
L302:
	;
	goto L294
L303:
	;
	v1247 = v869 + int32(1)
	v1250 = v671 + (v869 ^ int32(-1))
	if v1250 <= int32(0) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v2254 = v924
	goto L3
L305:
	;
	goto L304
L306:
	;
	v1260 = int32(1) << (uint(v1247&int32(7)) % 32)
	v1262 = base.I32_div_s(v1247, int32(8))
	v1263 = v966 + v1262
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1263))))
	if v688 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1357))) = uint8(v1358)
	goto L305
L308:
	;
	v1267 = v1263
	v1268 = v1264
	v1271 = v1250
	v1272 = v1260
	goto L311
L309:
	;
	goto L310
L310:
	;
	v1302 = base.I32_div_s(v1247, int32(8))
	v1303 = v688 + v1302
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303))))
	v1305 = v1263
	v1306 = v1264
	v1308 = v1303
	v1309 = v1250
	v1310 = v1260
	v1311 = int32(1) << (uint(v1247&int32(7)) % 32)
	v1312 = v1304
	goto L319
L311:
	;
	v1276 = v1268 | v1272
	v1277 = int32(1)
	v1278 = v1271 - v1277
	v1280 = v1272 << (uint(v1277) % 32)
	if v1280 == int32(256) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	if v1292 != int32(1) {
		v1357 = v1290
		v1358 = v1291
		goto L307
	} else {
		goto L318
	}
L313:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1267))) = uint8(v1276)
	if v1278 == int32(0) {
		goto L305
	} else {
		goto L316
	}
L314:
	;
	v1290 = v1267
	v1291 = v1276
	v1292 = v1280
	goto L315
L315:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1271) {
		v1267 = v1290
		v1268 = v1291
		v1271 = v1278
		v1272 = v1292
		goto L311
	} else {
		goto L317
	}
L316:
	;
	v1286 = int32(1)
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267)+1)))
	v1290 = v1267 + v1286
	v1291 = v1287
	v1292 = v1286
	goto L315
L317:
	;
	goto L312
L318:
	;
	goto L305
L319:
	;
	if v1311&v1312 != 0 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	if v1335 == int32(1) {
		goto L305
	} else {
		goto L334
	}
L321:
	;
	v1319 = v1306 | v1310
	goto L323
L322:
	;
	v1319 = v1306 & (v1310 ^ int32(-1))
	goto L323
L323:
	;
	v1320 = int32(1)
	v1321 = v1309 - v1320
	v1323 = v1310 << (uint(v1320) % 32)
	if v1323 == int32(256) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1305))) = uint8(v1319)
	if v1321 == int32(0) {
		goto L305
	} else {
		goto L327
	}
L325:
	;
	v1333 = v1305
	v1334 = v1319
	v1335 = v1323
	goto L326
L326:
	;
	v1337 = v1311 << (uint(int32(1)) % 32)
	if v1337 == int32(256) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	v1329 = int32(1)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305)+1)))
	v1333 = v1305 + v1329
	v1334 = v1330
	v1335 = v1329
	goto L326
L328:
	;
	goto L320
L329:
	;
	if v1321 == int32(0) {
		goto L328
	} else {
		goto L332
	}
L330:
	;
	v1346 = v1308
	v1347 = v1337
	v1348 = v1312
	goto L331
L331:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1309) {
		v1305 = v1333
		v1306 = v1334
		v1308 = v1346
		v1309 = v1321
		v1310 = v1335
		v1311 = v1347
		v1312 = v1348
		goto L319
	} else {
		goto L333
	}
L332:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308)+1)))
	v1343 = int32(1)
	v1346 = v1308 + v1343
	v1347 = v1343
	v1348 = v1342
	goto L331
L333:
	;
	goto L328
L334:
	;
	v1357 = v1333
	v1358 = v1334
	goto L307
L335:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L27
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v31+int32(32))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L27
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2342), int32(_a_F_array_set_element_4))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L27
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L27
	} else {
		goto L340
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v31+int32(48))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L27
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2357), int32(_a_F_array_set_element_4))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L27
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	if v1422&v1423 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if v1468 == int32(1) {
		v2254 = v924
		goto L3
	} else {
		goto L358
	}
L345:
	;
	v1453 = v1421 | v1425
	goto L347
L346:
	;
	v1453 = v1421 & (v1425 ^ int32(-1))
	goto L347
L347:
	;
	v1454 = int32(1)
	v1455 = v1432 - v1454
	v1457 = v1425 << (uint(v1454) % 32)
	if v1457 == int32(256) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1426))) = uint8(v1453)
	if v1455 == int32(0) {
		v2254 = v924
		goto L3
	} else {
		goto L351
	}
L349:
	;
	v1467 = v1453
	v1468 = v1457
	v1469 = v1426
	goto L350
L350:
	;
	v1471 = v1422 << (uint(int32(1)) % 32)
	if v1471 == int32(256) {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426)+1)))
	v1464 = int32(1)
	v1467 = v1463
	v1468 = v1464
	v1469 = v1426 + v1464
	goto L350
L352:
	;
	goto L344
L353:
	;
	if v1455 == int32(0) {
		goto L352
	} else {
		goto L356
	}
L354:
	;
	v1480 = v1471
	v1481 = v1423
	v1482 = v1435
	goto L355
L355:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1432) {
		v1421 = v1467
		v1422 = v1480
		v1423 = v1481
		v1425 = v1468
		v1426 = v1469
		v1432 = v1455
		v1435 = v1482
		goto L343
	} else {
		goto L357
	}
L356:
	;
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1435)+1)))
	v1477 = int32(1)
	v1480 = v1477
	v1481 = v1476
	v1482 = v1435 + v1477
	goto L355
L357:
	;
	goto L352
L358:
	;
	v1492 = v1467
	v1497 = v1469
	goto L8
L359:
	;
	v1546 = v1542 + v1543
	if v1520 < v1546 {
		v1565 = v1541
		v1578 = v1544
		v1579 = v1545
		goto L6
	} else {
		goto L365
	}
L360:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v1541 = v138
	v1542 = v1523
	v1543 = v1521
	v1544 = v139
	v1545 = v144
	goto L359
L361:
	;
	goto L362
L362:
	;
	v1524 = v1521 - v1520
	if base.B2i32(v1524 < v1521)^base.B2i32(int32(0) < v1520) != 0 {
		goto L2
	} else {
		goto L363
	}
L363:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v1530 = v1529 + v1524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1530
	if base.B2i32(v1524 < int32(0)) != base.B2i32(v1530 < v1529) {
		goto L2
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1520
	v1537 = int32(1)
	v1541 = v1524
	v1542 = v1530
	v1543 = v1520
	v1544 = v1537
	v1545 = base.B2i32(v1537 < v1524) | v144
	goto L359
L365:
	;
	v1550 = v1520 - v1546
	if base.B2i32(int32(0) < v1546)^base.B2i32(v1550 < v1520) != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v1554 = v1550 + int32(1)
	if v1554 < v1550 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	v1556 = v1542 + v1554
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1556
	if base.B2i32(v1554 < int32(0)) != base.B2i32(v1556 < v1542) {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	v1597 = v1541
	v1609 = v1554
	v1611 = base.B2i32(int32(1) < v1554) | v1545
	goto L5
L369:
	;
	v1597 = v1565
	v1609 = v10
	v1611 = v1579
	goto L5
L370:
	;
	F_ArrayCheckBounds(m, l1, v1626, v31-int32(-64))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L27
	} else {
		goto L371
	}
L371:
	;
	v1634 = v1597
	v1641 = v1609
	v1648 = v1611
	v1662 = int32(1)
	goto L4
L372:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v75)+56))
	if v1761 < v1760 {
		goto L382
	} else {
		goto L383
	}
L373:
	;
	goto L372
L374:
	;
	v1679 = int32(1)
	if v1676 != 0 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1688 = v1676
	v1689 = v1679
	v1690 = v1667
	v1695 = v1667
	goto L378
L376:
	;
	v1731 = v1676
	v1732 = v1679
	v1733 = v1667
	goto L377
L377:
	;
	v1740 = v1731 << (uint(int32(2)) % 32)
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1740)))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1740+v1666)))
	v1754 = (v1742-v1744)*v1732 + v1733
	goto L373
L378:
	;
	v1696 = int32(2)
	v1697 = v1688 << (uint(v1696) % 32)
	v1699 = v1697 - int32(4)
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1699)))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1666+v1699)))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1697+v1664)))
	v1707 = v1706 * v1689
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1697+l2)))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1697+v1666)))
	v1716 = (v1701-v1703)*v1707 + ((v1710-v1712)*v1689 + v1690)
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1664+v1699)))
	v1719 = v1718 * v1707
	v1721 = v1688 - v1696
	v1723 = v1695 + v1696
	if v1723 != l1&int32(-2) {
		v1688 = v1721
		v1689 = v1719
		v1690 = v1716
		v1695 = v1723
		goto L378
	} else {
		goto L380
	}
L379:
	;
	if l1&int32(1) == int32(0) {
		v1754 = v1716
		goto L373
	} else {
		goto L381
	}
L380:
	;
	goto L379
L381:
	;
	v1731 = v1721
	v1732 = v1719
	v1733 = v1716
	goto L377
L382:
	;
	v1764 = base.I32_div_s(v1760, int32(8))
	v1765 = v1764 + v1760
	if v1760 < v1765 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	v1781 = v141
	v1782 = v1761
	v1783 = v145
	goto L384
L384:
	;
	if base.B2i32(v1781 == int32(0))&v1648 != 0 {
		goto L394
	} else {
		goto L395
	}
L385:
	;
	v1767 = v1765
	goto L387
L386:
	;
	v1767 = v1760
	goto L387
L387:
	;
	v1770 = F_repalloc(m, v145, v1767<<(uint(int32(2))%32))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L27
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+48)) = v1770
	if v141 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+56)) = v1767
	v1781 = v1779
	v1782 = v1767
	v1783 = v1770
	goto L384
L390:
	;
	v1779 = int32(0)
	goto L389
L391:
	;
	goto L392
L392:
	;
	v1776 = F_repalloc(m, v141, v1767)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L27
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+52)) = v1776
	v1779 = v1776
	goto L389
L394:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v1788 = F_MemoryContextAllocZero(m, v1787, v1782)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L27
	} else {
		goto L397
	}
L395:
	;
	v1791 = v1781
	goto L396
L396:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+64)) = int64(0)
	if v1662 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+52)) = v1788
	v1791 = v1788
	goto L396
L398:
	;
	if int32(0) < v1634 {
		goto L404
	} else {
		goto L405
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+28)) = l1
	v1798 = l1 << (uint(int32(2)) % 32)
	v1799 = int32(0)
	v1800 = base.B2i32(v1798 == v1799)
	if v1800 == v1799 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	base.MemoryCopy(m, v1803, v31+int32(96), v1798)
	goto L402
L401:
	;
	goto L402
L402:
	;
	if v1798 == v1799 {
		goto L398
	} else {
		goto L403
	}
L403:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
	base.MemoryCopy(m, v1807, v31-int32(-64), v1798)
	goto L398
L404:
	;
	v1815 = int32(2)
	v1816 = v1634 << (uint(v1815) % 32)
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v1819 = v1817 << (uint(v1815) % 32)
	if v1819 != 0 {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	goto L406
L406:
	;
	if int32(0) < v1641 {
		goto L419
	} else {
		goto L420
	}
L407:
	;
	base.MemoryCopy(m, v1816+v1783, v1783, v1819)
	goto L409
L408:
	;
	goto L409
L409:
	;
	if v1816 != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	base.MemoryFill(m, v1783, int32(0), v1816)
	goto L412
L411:
	;
	goto L412
L412:
	;
	if v1791 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+60)) = v1834 + v1634
	goto L406
L414:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	if v1826 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	base.MemoryCopy(m, v1634+v1791, v1791, v1826)
	goto L417
L416:
	;
	goto L417
L417:
	;
	if v1634 == int32(0) {
		goto L413
	} else {
		goto L418
	}
L418:
	;
	base.MemoryFill(m, v1791, int32(1), v1634)
	goto L413
L419:
	;
	v1842 = v1641 & int32(3)
	v1843 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1641) {
		goto L423
	} else {
		goto L424
	}
L420:
	;
	goto L421
L421:
	;
	v2204 = int32(0)
	v2205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+46)))
	if v2205 == v2204 {
		goto L447
	} else {
		goto L448
	}
L422:
	;
	if v1791 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L423:
	;
	v1849 = int32(0)
	v1850 = v1843
	goto L426
L424:
	;
	v1915 = v1843
	goto L425
L425:
	;
	v1944 = v1915
	v1945 = int32(0)
	goto L430
L426:
	;
	v1877 = int32(2)
	v1878 = v1850 << (uint(v1877) % 32)
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v1884 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1878+(v1783+v1879<<(uint(v1877)%32))))) = v1884
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1886<<(uint(v1877)%32)+v1878)+4)) = v1884
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1893<<(uint(v1877)%32)+v1878)+8)) = v1884
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1900<<(uint(v1877)%32)+v1878)+12)) = v1884
	v1907 = int32(4)
	v1908 = v1850 + v1907
	v1910 = v1849 + v1907
	if v1910 != v1641&int32(2147483644) {
		v1849 = v1910
		v1850 = v1908
		goto L426
	} else {
		goto L428
	}
L427:
	;
	if v1842 == int32(0) {
		goto L422
	} else {
		goto L429
	}
L428:
	;
	goto L427
L429:
	;
	v1915 = v1908
	goto L425
L430:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v1972 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1971<<(uint(v1972)%32)+v1944<<(uint(v1972)%32)))) = int32(0)
	v1980 = int32(1)
	v1983 = v1945 + v1980
	if v1983 != v1842 {
		v1944 = v1944 + v1980
		v1945 = v1983
		goto L430
	} else {
		goto L432
	}
L431:
	;
	goto L422
L432:
	;
	goto L431
L433:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+60)) = v2173 + v1641
	goto L421
L434:
	;
	v2016 = v1641 & int32(3)
	v2017 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1641) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2024 = int32(0)
	v2025 = v2017
	goto L438
L436:
	;
	v2080 = v2017
	goto L437
L437:
	;
	v2108 = v2080
	v2109 = v2017
	goto L442
L438:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v2055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1791+v2052+v2025))) = uint8(v2055)
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1791+v2057+v2025)+1)) = uint8(v2055)
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1791+v2062+v2025)+2)) = uint8(v2055)
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1791+v2067+v2025)+3)) = uint8(v2055)
	v2072 = int32(4)
	v2073 = v2025 + v2072
	v2075 = v2024 + v2072
	if v2075 != v1641&int32(2147483644) {
		v2024 = v2075
		v2025 = v2073
		goto L438
	} else {
		goto L440
	}
L439:
	;
	if v2016 == int32(0) {
		goto L433
	} else {
		goto L441
	}
L440:
	;
	goto L439
L441:
	;
	v2080 = v2073
	goto L437
L442:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v2138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1791+v2135+v2108))) = uint8(v2138)
	v2143 = v2109 + v2138
	if v2143 != v2016 {
		v2108 = v2108 + v2138
		v2109 = v2143
		goto L442
	} else {
		goto L444
	}
L443:
	;
	goto L433
L444:
	;
	goto L443
L445:
	;
	if v2232 == int32(0) {
		goto L454
	} else {
		goto L455
	}
L446:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1754+v1791))) = uint8(v5)
	v2232 = v2229
	goto L445
L447:
	;
	if v1791 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	v2222 = v2204
	goto L449
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1754<<(uint(int32(2))%32)))) = v137
	if v1791 == int32(0) {
		v2232 = v2222
		goto L445
	} else {
		goto L453
	}
L450:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v1783+v1754<<(uint(int32(2))%32))))
	v2222 = v2221
	goto L449
L451:
	;
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754+v1791))))
	if v2211 != int32(1) {
		goto L450
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783+v1754<<(uint(int32(2))%32)))) = v137
	v2229 = v2204
	goto L446
L453:
	;
	v2229 = v2222
	goto L446
L454:
	;
	v2254 = v75 + int32(12)
	goto L3
L455:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
	if base.Ui32(v2235) <= base.Ui32(v2232) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v75)+76))
	if base.Ui32(v2232) < base.Ui32(v2237) {
		goto L454
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	F_pfree(m, v2232)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L27
	} else {
		goto L460
	}
L459:
	;
	goto L458
L460:
	;
	goto L454
L461:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L27
	} else {
		goto L462
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v31)
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L27
	} else {
		goto L463
	}
L463:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2619), int32(_a_F_array_set_element_3))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L27
	} else {
		goto L464
	}
L464:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L465:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L27
	} else {
		goto L466
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_set_element_7), v31+int32(16))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L27
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_array_set_element_2), int32(2635), int32(_a_F_array_set_element_3))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L27
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_subscript_check_subscripts(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v144
L2:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v136)
	v144 = int32(0)
	goto L1
L3:
	;
	v19 = int32(0)
	v20 = v11
	goto L6
L4:
	;
	goto L5
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v74 <= int32(0) {
		v144 = v8
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v19))))
	if v26 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v19))))
	if v31 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v63 = v20
	goto L10
L10:
	;
	v65 = v19 + int32(1)
	if v65 < v63 {
		v19 = v65
		v20 = v63
		goto L6
	} else {
		goto L20
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v34 != int32(1) {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v56 = v19 << (uint(int32(2)) % 32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)+v56))) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v63 = v62
	goto L10
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_array_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(199), int32(_a_F_array_subscript_check_subscripts_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	goto L7
L21:
	;
	v82 = int32(0)
	v83 = v74
	goto L22
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v82))))
	if v89 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v144 = v8
	goto L1
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v82))))
	if v94 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v124 = v83
	goto L26
L26:
	;
	v126 = v82 + int32(1)
	if v126 < v124 {
		v82 = v126
		v83 = v124
		goto L22
	} else {
		goto L35
	}
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v97 != int32(1) {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v117 = v82 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119+v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(36)+v117))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v124 = v123
	goto L26
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_array_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(218), int32(_a_F_array_subscript_check_subscripts_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	goto L23
}
func F_array_to_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = F_pg_detoast_datum_packed(m, v13)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v17 == int32(1) {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
					if v23 == int32(18) {
						v26 = int32(16)
					} else {
						v26 = int32(0)
					}
					if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v33 = int32(4)
					} else {
						v33 = v26
					}
					v46 = v33
				} else {
					v34 = int32(1)
					if v17&v34 != 0 {
						v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
						v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v49 = F_palloc(m, v46+int32(1))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v46 != 0 {
						v51 = int32(1)
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if v53&v51 != 0 {
							v56 = v51
						} else {
							v56 = int32(4)
						}
						base.MemoryCopy(m, v49, v15+v56, v46)
					} else {
					}
					v60 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v46+v49))) = uint8(v60)
					if v15 != v13 {
						F_pfree(m, v15)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v66 = F_array_to_text_internal(m, l0, v8, v49, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								return v66
							}
						}
					} else {
						v66 = F_array_to_text_internal(m, l0, v8, v49, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							return v66
						}
					}
				}
			}
		}
	}
}
func F_array_to_vector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 float32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 float32
	_ = v141
	var v144 int32
	_ = v144
	var v147 float32
	_ = v147
	var v150 int32
	_ = v150
	var v153 float32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 float32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 float64
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 float64
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 float64
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 float64
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v420 float32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L82
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L78
	}
L4:
	;
	return int32(0)
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L74
	}
L9:
	;
	v26 = F_array_contains_nulls(m, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_get_typlenbyvalalign(m, v28, v14+int32(30), v14+int32(29), v14+int32(28))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	if v26 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+30)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+29)))
	v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+28)))
	F_deconstruct_array(m, v17, v38, v39, v40, v14+int32(24), int32(0), v14+int32(20))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_CheckDim_3(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if base.B2i32(v24 != int32(-1))&base.B2i32(v53 != v24) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v58 = F_mul_size(m, int32(4), v53)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v60 = F_add_size(m, int32(8), v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v62 = F_palloc0(m, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v60 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	switch v68 - int32(700) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L25
	}
L21:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	F_pfree(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L66
	}
L22:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v293 <= int32(0) {
		goto L21
	} else {
		goto L55
	}
L23:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v195 <= int32(0) {
		goto L21
	} else {
		goto L44
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v107 <= int32(0) {
		goto L21
	} else {
		goto L33
	}
L25:
	;
	if v68 == int32(23) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v68 != int32(1700) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v75 <= int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v81 = int32(0)
	goto L29
L29:
	;
	v93 = v81 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97+v93)))
	v100 = F_DirectFunctionCall1Coll(m, int32(1319), int32(0), v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62+int32(8)+v93))) = v100
	v104 = v81 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v104 < v105 {
		v81 = v104
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v111 = v107 & int32(3)
	v113 = v62 + int32(8)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v107) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = v115
	v126 = v2
	goto L37
L35:
	;
	v162 = v115
	goto L36
L36:
	;
	v173 = v162
	v183 = v2
	goto L41
L37:
	;
	v132 = v120 << (uint(int32(2)) % 32)
	v135 = *(*float32)(unsafe.Add(mBase, uint32(v132+v114)))
	*(*float32)(unsafe.Add(mBase, uint32(v113+v132))) = v135
	v137 = int32(4)
	v138 = v132 | v137
	v141 = *(*float32)(unsafe.Add(mBase, uint32(v114+v138)))
	*(*float32)(unsafe.Add(mBase, uint32(v113+v138))) = v141
	v144 = v132 | int32(8)
	v147 = *(*float32)(unsafe.Add(mBase, uint32(v114+v144)))
	*(*float32)(unsafe.Add(mBase, uint32(v113+v144))) = v147
	v150 = v132 | int32(12)
	v153 = *(*float32)(unsafe.Add(mBase, uint32(v150+v114)))
	*(*float32)(unsafe.Add(mBase, uint32(v113+v150))) = v153
	v156 = v120 + v137
	v158 = v126 + v137
	if v158 != v107&int32(2147483644) {
		v120 = v156
		v126 = v158
		goto L37
	} else {
		goto L39
	}
L38:
	;
	if v111 == int32(0) {
		goto L21
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v162 = v156
	goto L36
L41:
	;
	v185 = v173 << (uint(int32(2)) % 32)
	v188 = *(*float32)(unsafe.Add(mBase, uint32(v114+v185)))
	*(*float32)(unsafe.Add(mBase, uint32(v113+v185))) = v188
	v190 = int32(1)
	v193 = v183 + v190
	if v193 != v111 {
		v173 = v173 + v190
		v183 = v193
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L21
L43:
	;
	goto L42
L44:
	;
	v199 = v195 & int32(3)
	v201 = v62 + int32(8)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v203 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v195) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v208 = v203
	v214 = v2
	goto L48
L46:
	;
	v258 = v203
	goto L47
L47:
	;
	v269 = v258
	v279 = v2
	goto L52
L48:
	;
	v220 = v208 << (uint(int32(2)) % 32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v202)))
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v223)))
	*(*float32)(unsafe.Add(mBase, uint32(v201+v220))) = base.F32_demote_f64(v224)
	v227 = int32(4)
	v228 = v220 | v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v202+v228)))
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v231)))
	*(*float32)(unsafe.Add(mBase, uint32(v201+v228))) = base.F32_demote_f64(v232)
	v236 = v220 | int32(8)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v202+v236)))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v239)))
	*(*float32)(unsafe.Add(mBase, uint32(v201+v236))) = base.F32_demote_f64(v240)
	v244 = v220 | int32(12)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244+v202)))
	v248 = *(*float64)(unsafe.Add(mBase, uint32(v247)))
	*(*float32)(unsafe.Add(mBase, uint32(v201+v244))) = base.F32_demote_f64(v248)
	v252 = v208 + v227
	v254 = v214 + v227
	if v254 != v195&int32(2147483644) {
		v208 = v252
		v214 = v254
		goto L48
	} else {
		goto L50
	}
L49:
	;
	if v199 == int32(0) {
		goto L21
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v258 = v252
	goto L47
L52:
	;
	v281 = v269 << (uint(int32(2)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v202+v281)))
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v284)))
	*(*float32)(unsafe.Add(mBase, uint32(v201+v281))) = base.F32_demote_f64(v285)
	v288 = int32(1)
	v291 = v279 + v288
	if v291 != v199 {
		v269 = v269 + v288
		v279 = v291
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L21
L54:
	;
	goto L53
L55:
	;
	v297 = v293 & int32(3)
	v299 = v62 + int32(8)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v301 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v293) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v306 = v301
	v312 = v2
	goto L59
L57:
	;
	v352 = v301
	goto L58
L58:
	;
	v363 = v352
	v373 = v2
	goto L63
L59:
	;
	v318 = v306 << (uint(int32(2)) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318+v300)))
	*(*float32)(unsafe.Add(mBase, uint32(v299+v318))) = base.F32_convert_i32_s(v321)
	v324 = int32(4)
	v325 = v318 | v324
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v300+v325)))
	*(*float32)(unsafe.Add(mBase, uint32(v299+v325))) = base.F32_convert_i32_s(v328)
	v332 = v318 | int32(8)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v300+v332)))
	*(*float32)(unsafe.Add(mBase, uint32(v299+v332))) = base.F32_convert_i32_s(v335)
	v339 = v318 | int32(12)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v339+v300)))
	*(*float32)(unsafe.Add(mBase, uint32(v299+v339))) = base.F32_convert_i32_s(v342)
	v346 = v306 + v324
	v348 = v312 + v324
	if v348 != v293&int32(2147483644) {
		v306 = v346
		v312 = v348
		goto L59
	} else {
		goto L61
	}
L60:
	;
	if v297 == int32(0) {
		goto L21
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v352 = v346
	goto L58
L63:
	;
	v375 = v363 << (uint(int32(2)) % 32)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v300+v375)))
	*(*float32)(unsafe.Add(mBase, uint32(v299+v375))) = base.F32_convert_i32_s(v378)
	v381 = int32(1)
	v384 = v373 + v381
	if v384 != v297 {
		v363 = v363 + v381
		v373 = v384
		goto L63
	} else {
		goto L65
	}
L64:
	;
	goto L21
L65:
	;
	goto L64
L66:
	;
	v400 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+4)))
	if int32(0) < v400 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v406 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	m.G0 = v14 + int32(32)
	return v62
L70:
	;
	v420 = *(*float32)(unsafe.Add(mBase, uint32(v62+int32(8)+v406<<(uint(int32(2))%32))))
	F_CheckElement_3(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	v424 = v406 + int32(1)
	v425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+4)))
	if v424 < v425 {
		v406 = v424
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_0), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(459), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_3), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(464), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v24
	F_errmsg(m, int32(_a_F_array_to_vector_4), v14)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(88), int32(_a_F_array_to_vector_5))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
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
	F_errcode(m, int32(130))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_6), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(498), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_typanalyze(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_std_typanalyze(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v66 = int32(0)
			m.G0 = v9 + int32(16)
			return v66
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = F_get_base_element_type(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v76
						F_errmsg_internal(m, int32(_a_F_array_typanalyze_0), v9)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_typanalyze_1), int32(118), int32(_a_F_array_typanalyze_2))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = int32(1)
					v25 = F_lookup_type_cache(m, v19, int32(193))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
						if v27 == int32(0) {
							v66 = v23
							m.G0 = v9 + int32(16)
							return v66
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+108))
							if v30 == int32(0) {
								v66 = v23
								m.G0 = v9 + int32(16)
								return v66
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+136))
								if v33 == int32(0) {
									v66 = v23
									m.G0 = v9 + int32(16)
									return v66
								} else {
									v37 = F_palloc(m, int32(36))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
										*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v41
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v43
										v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+10)))
										*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)) = uint8(v45)
										v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v37)+14)) = uint16(v47)
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+11)))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = v25 + int32(132)
										*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v25 + int32(104)
										*(*uint8)(unsafe.Add(mBase, uint32(v37)+16)) = uint8(v49)
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v57
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v59
										*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v37
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1244)
										v66 = v23
										m.G0 = v9 + int32(16)
										return v66
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
func F_array_unnest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(_a_F_array_unnest_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_DatumGetAnyArrayP(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = F_palloc(m, int32(32))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v27 == int32(-1) {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
						if v30 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v30
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
							v33 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
							v87 = v33
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
							if v40 != 0 {
								v48 = v40
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v48 + v39
							v51 = int32(0)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
							if v53 == v51 {
								v87 = v51
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
								v87 = v52 + v56<<(uint(int32(3))%32) + int32(16)
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(0)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
						if v64 != 0 {
							v72 = v64
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							v72 = (v65<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v72 + v22
						v75 = int32(0)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
						if v76 == v75 {
							v87 = v75
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							v87 = v22 + v79<<(uint(int32(3))%32) + int32(16)
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = int64(1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v87
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v93 == int32(-1) {
						v96 = int32(28)
					} else {
						v96 = int32(4)
					}
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v22+v96)))
					if v93 == int32(-1) {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
						v104 = v101
					} else {
						v104 = v22 + int32(16)
					}
					v105 = F_ArrayGetNItemsSafe(m, v98, v104)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v105
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						if v108 == int32(-1) {
							v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)))
							*(*uint16)(unsafe.Add(mBase, uint32(v25)+28)) = uint16(v111)
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+30)) = uint8(v113)
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+47)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+31)) = uint8(v115)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v25
							*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v17
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
							if v138 < v139 {
								*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v138 + int32(1)
								v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+28)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+30)))
								v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137)+31)))
								v149 = F_array_iter_next(m, v137, l0+int32(16), v138, v146, v147, v148)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v151 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
									*(*int64)(unsafe.Add(mBase, uint32(v136))) = v151 + int64(1)
									v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v155)+20)) = int32(1)
									return v149
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = int32(2)
									v164 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v164)
									return int32(0)
								}
							}
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							F_get_typlenbyvalalign(m, v117, v25+int32(28), v25+int32(30), v25+int32(31))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v25
								*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v17
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
								if v138 < v139 {
									*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v138 + int32(1)
									v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+28)))
									v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+30)))
									v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137)+31)))
									v149 = F_array_iter_next(m, v137, l0+int32(16), v138, v146, v147, v148)
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										v151 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
										*(*int64)(unsafe.Add(mBase, uint32(v136))) = v151 + int64(1)
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v155)+20)) = int32(1)
										return v149
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = int32(2)
										v164 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v164)
										return int32(0)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
		v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
		if v138 < v139 {
			*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v138 + int32(1)
			v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+28)))
			v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+30)))
			v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137)+31)))
			v149 = F_array_iter_next(m, v137, l0+int32(16), v138, v146, v147, v148)
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				v151 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
				*(*int64)(unsafe.Add(mBase, uint32(v136))) = v151 + int64(1)
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v155)+20)) = int32(1)
				return v149
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = int32(2)
				v164 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v164)
				return int32(0)
			}
		}
	}
}
func F_array_unnest_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(460) {
		v26 = v2
		return v26
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		if v8 == int32(0) {
			v26 = v2
			return v26
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v11 != int32(15) {
				v26 = v2
				return v26
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = F_estimate_expression_value(m, v14, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					v23 = F_estimate_array_length(m, v22, v18)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v4)+16)) = v23
						v26 = v4
						return v26
					}
				}
			}
		}
	}
}
func F_construct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v13
	v21 = F_construct_md_array(m, l0, int32(0), v13, v10+int32(12), v10+int32(8), l2, l3, l4, l5)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(16)
		return v21
	}
}
func F_fetch_array_arg_replace_nulls(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		v16 = F_MemoryContextAlloc(m, v14, int32(48))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v16
			v24 = v16
			v26 = v8 + int32(12)
			v27 = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v28 == v27 {
				v45 = int32(0)
				if v26 == v45 {
					v53 = v45
				} else {
					v48 = v45
					v49 = v27
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
				}
				v56 = v53
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				switch v31 - int32(429) {
				case 0:
					if v26 == int32(0) {
						v56 = int32(1)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
						v48 = v38
						v49 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
						v56 = v53
					}
				case 1:
					if v26 == int32(0) {
						v56 = int32(2)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+368))
						v48 = v43
						v49 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
						v56 = v53
					}
				default:
					v45 = int32(0)
					if v26 == v45 {
						v53 = v45
					} else {
						v48 = v45
						v49 = v27
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
						v53 = v49
					}
					v56 = v53
				}
			}
			if v56 == int32(0) {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v60
			} else {
			}
			v64 = l0 + l1<<(uint(int32(3))%32)
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
			if v65 == int32(0) {
				v68 = int32(_a_F_fetch_array_arg_replace_nulls_0)
				v69 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v71
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
				if v74 != int32(1) {
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v93 = F_expand_array(m, v73, v92, v24)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
						v97 = v95
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				} else {
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
					if v77 != int32(3) {
						v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
						v93 = F_expand_array(m, v73, v92, v24)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
							v97 = v95
							*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
							v123 = v97
							m.G0 = v8 + int32(16)
							return v123
						}
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2))
						if v24 == int32(0) {
							v97 = v80
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v24))) = v83
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)))
							*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v85)
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v87)
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+47)))
							*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v89)
							v97 = v80
						}
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				}
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v101 = F_get_fn_expr_argtype(m, v100, l1)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					if v101 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(118), int32(_a_F_fetch_array_arg_replace_nulls_3))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
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
						v105 = F_get_element_type(m, v101)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							if v105 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(123), int32(_a_F_fetch_array_arg_replace_nulls_3))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
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
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v111 = F_palloc0(m, int32(16))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v105
									*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v111))) = int64(64)
									v118 = F_expand_array(m, v111, v109, v24)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+2))
											v123 = v122
											m.G0 = v8 + int32(16)
											return v123
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
		v24 = v11
		v26 = v8 + int32(12)
		v27 = int32(0)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v28 == v27 {
			v45 = int32(0)
			if v26 == v45 {
				v53 = v45
			} else {
				v48 = v45
				v49 = v27
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
				v53 = v49
			}
			v56 = v53
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			switch v31 - int32(429) {
			case 0:
				if v26 == int32(0) {
					v56 = int32(1)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
					v48 = v38
					v49 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
					v56 = v53
				}
			case 1:
				if v26 == int32(0) {
					v56 = int32(2)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+368))
					v48 = v43
					v49 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
					v56 = v53
				}
			default:
				v45 = int32(0)
				if v26 == v45 {
					v53 = v45
				} else {
					v48 = v45
					v49 = v27
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v48
					v53 = v49
				}
				v56 = v53
			}
		}
		if v56 == int32(0) {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v60
		} else {
		}
		v64 = l0 + l1<<(uint(int32(3))%32)
		v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
		if v65 == int32(0) {
			v68 = int32(_a_F_fetch_array_arg_replace_nulls_0)
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
			if v74 != int32(1) {
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v93 = F_expand_array(m, v73, v92, v24)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
					v97 = v95
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
					v123 = v97
					m.G0 = v8 + int32(16)
					return v123
				}
			} else {
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
				if v77 != int32(3) {
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v93 = F_expand_array(m, v73, v92, v24)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+2))
						v97 = v95
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
						v123 = v97
						m.G0 = v8 + int32(16)
						return v123
					}
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2))
					if v24 == int32(0) {
						v97 = v80
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v24))) = v83
						v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+44)))
						*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v85)
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+46)))
						*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v87)
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+47)))
						*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v89)
						v97 = v80
					}
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v69
					v123 = v97
					m.G0 = v8 + int32(16)
					return v123
				}
			}
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v101 = F_get_fn_expr_argtype(m, v100, l1)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				if v101 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(118), int32(_a_F_fetch_array_arg_replace_nulls_3))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
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
					v105 = F_get_element_type(m, v101)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						if v105 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(123), int32(_a_F_fetch_array_arg_replace_nulls_3))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
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
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							v111 = F_palloc0(m, int32(16))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v105
								*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v111))) = int64(64)
								v118 = F_expand_array(m, v111, v109, v24)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v111)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+2))
										v123 = v122
										m.G0 = v8 + int32(16)
										return v123
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
func F_initArrayResultWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	v3 = l2
	if v3 != 0 {
		v10 = F_AllocSetContextCreateInternal(m, l1, int32(_a_F_initArrayResultWithSize_0), int32(0), int32(_a_F_initArrayResultWithSize_1), int32(_a_F_initArrayResultWithSize_2))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = v10
			v16 = F_MemoryContextAlloc(m, v14, int32(32))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
				v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(2))%32))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v27 = F_MemoryContextAlloc(m, v14, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v27
						F_get_typlenbyvalalign(m, l0, v16+int32(24), v16+int32(26), v16+int32(27))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							return v16
						}
					}
				}
			}
		}
	} else {
		v14 = l1
		v16 = F_MemoryContextAlloc(m, v14, int32(32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
			v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(2))%32))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v27 = F_MemoryContextAlloc(m, v14, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v27
					F_get_typlenbyvalalign(m, l0, v16+int32(24), v16+int32(26), v16+int32(27))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			}
		}
	}
}
