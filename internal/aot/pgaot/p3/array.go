package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayCastAndSet(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if int32(0) < l1 {
		if l2 != 0 {
			if l1&(l1-int32(1)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(_a_F_ArrayCastAndSet_0), v10)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ArrayCastAndSet_1), int32(474), int32(_a_F_ArrayCastAndSet_2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				switch base.I32_ctz(l1) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v1)
					v72 = l1
					v73 = l3
					m.G0 = v10 + int32(16)
					return (v72 + v73 - int32(1)) & (int32(0) - l3)
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v1)
					v72 = l1
					v73 = l3
					m.G0 = v10 + int32(16)
					return (v72 + v73 - int32(1)) & (int32(0) - l3)
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v1)
					v72 = l1
					v73 = l3
					m.G0 = v10 + int32(16)
					return (v72 + v73 - int32(1)) & (int32(0) - l3)
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(l4))) = v1
					v72 = l1
					v73 = l3
					m.G0 = v10 + int32(16)
					return (v72 + v73 - int32(1)) & (int32(0) - l3)
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(_a_F_ArrayCastAndSet_0), v10)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ArrayCastAndSet_1), int32(474), int32(_a_F_ArrayCastAndSet_2))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
		} else {
			if l1 != 0 {
				base.MemoryCopy(m, l4, base.I32_wrap_i64(v1), l1)
			} else {
			}
			v72 = l1
			v73 = l3
			m.G0 = v10 + int32(16)
			return (v72 + v73 - int32(1)) & (int32(0) - l3)
		}
	} else {
		v39 = base.I32_wrap_i64(v1)
		if l1 == int32(-1) {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
			if v42 == int32(1) {
				v46 = int32(18)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
				if v48 == v46 {
					v51 = v46
				} else {
					v51 = int32(2)
				}
				if base.Ui32((v48-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v58 = int32(6)
				} else {
					v58 = v51
				}
				v70 = v58
			} else {
				v59 = int32(1)
				if v42&v59 != 0 {
					v70 = int32(base.Ui32(v42) >> (uint(v59) % 32))
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v70 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
				}
			}
		} else {
			v66 = F_strlen(m, v39)
			mBase = m.M
			v70 = v66 + int32(1)
		}
		if v70 != 0 {
			base.MemoryCopy(m, l4, v39, v70)
		} else {
		}
		v72 = l3
		v73 = v70
		m.G0 = v10 + int32(16)
		return (v72 + v73 - int32(1)) & (int32(0) - l3)
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
func F_array_dims(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int64
	_ = v99
	v10 = m.G0
	v12 = v10 - int32(224)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_DatumGetAnyArrayP(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(224)
	return v99
L2:
	;
	return int64(0)
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v21 == int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(28)
	goto L6
L5:
	;
	v24 = int32(4)
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15+v24)))
	if base.Ui32(v26-int32(7)) <= base.Ui32(int32(-7)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v99 = int64(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	if v21 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v49 = v12 + int32(16)
	v52 = int32(0)
	goto L14
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v44 = v36
	v45 = v37
	goto L10
L12:
	;
	goto L13
L13:
	;
	v39 = v15 + int32(16)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v44 = v39
	v45 = v39 + v40<<(uint(int32(2))%32)
	goto L10
L14:
	;
	v59 = v52 << (uint(int32(2)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44+v59)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v63 + v61 - int32(1)
	v70 = F_pg_sprintf(m, v49, int32(_a_F_array_dims_0), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v87 = F_cstring_to_text(m, v12+int32(16))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L21
	}
L16:
	;
	v72 = F_strlen(m, v49)
	mBase = m.M
	v75 = v52 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v78 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v81 = int32(28)
	goto L19
L18:
	;
	v81 = int32(4)
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15+v81)))
	if v75 < v83 {
		v49 = v72 + v49
		v52 = v75
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v99 = base.I64_extend_i32_u(v87)
	goto L1
}
func F_array_in(m *base.Module, l0 int32) int64 {
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
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v986 int32
	_ = v986
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1009 int32
	_ = v1009
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1252 int32
	_ = v1252
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1340 int32
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1416 int32
	_ = v1416
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1542 int32
	_ = v1542
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1616 int32
	_ = v1616
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(592)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+4)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+7)))
	v80 = base.I32_extend8_s(v79)
	switch v79 - int32(99) {
	case 0:
		v100 = int32(1)
		goto L11
	case 1:
		goto L14
	default:
		goto L13
	case 6:
		goto L15
	case 16:
		goto L12
	}
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
	return int64(0)
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
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76)+8)))
	v103 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+528)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(v27)+520)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(v27)+512)) = v103
	v109 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+496)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v27)+488)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v27)+480)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v32
	v118 = v32
	v121 = v2
	goto L19
L12:
	;
	v100 = int32(2)
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	v100 = int32(8)
	goto L11
L15:
	;
	v100 = int32(4)
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v80
	F_errmsg_internal(m, int32(_a_F_array_in_0), v27)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_array_in_1), int32(322), int32(_a_F_array_in_2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v141 = v118 + int32(1)
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118))))
	goto L21
L20:
	;
	m.G0 = v27 + int32(592)
	return base.I64_extend_i32_u(v1616)
L21:
	;
	if base.B2i32(v142 == int32(32))|base.B2i32(base.Ui32((v142-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v118 = v141
		goto L19
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v118
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v153 == int32(91) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(512)+v121<<(uint(int32(2))%32)))) = v325
	v118 = v284
	v121 = v121 + int32(1)
	goto L19
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v141
	if v121 == int32(6) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L27
L27:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v121 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L28:
	;
	v1616 = int32(0)
	goto L23
L29:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), v346, int32(_a_F_array_in_4))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L91
	}
L30:
	;
	v160 = F_errsave_start(m, v29)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v179 = F_ReadDimensionInt(m, v27+int32(544), v27+int32(540), v29)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L37
	}
L33:
	;
	if v160 == int32(0) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_5), v27+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v346 = int32(436)
	goto L29
L37:
	;
	if v179 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v27)+544))
	if v141 == v183 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v185 = F_errsave_start(m, v29)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v203 == int32(58) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	if v185 == int32(0) {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v200 = F_errdetail(m, int32(_a_F_array_in_7), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v346 = int32(445)
	goto L29
L47:
	;
	if v256&int32(255) != int32(93) {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v27)+540))
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(480)+v121<<(uint(int32(2))%32)))) = v211
	v214 = v183 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v214
	v220 = F_ReadDimensionInt(m, v27+int32(544), v27+int32(576), v29)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v245 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(480)+v121<<(uint(int32(2))%32)))) = v245
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v27)+540))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+576)) = v253
	v255 = v183
	v256 = v203
	v257 = v245
	goto L47
L51:
	;
	if v220 == int32(0) {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v27)+544))
	if v214 != v224 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v255 = v224
	v256 = v226
	v257 = v211
	goto L47
L54:
	;
	goto L55
L55:
	;
	v227 = F_errsave_start(m, v29)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	if v227 == int32(0) {
		goto L28
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v242 = F_errdetail(m, int32(_a_F_array_in_8), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v346 = int32(459)
	goto L29
L61:
	;
	v262 = F_errsave_start(m, v29)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v284 = v255 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+544)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v27)+576))
	if v286 < v257 {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	if v262 == int32(0) {
		goto L28
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(112))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = int32(_a_F_array_in_9)
	v280 = F_errdetail(m, int32(_a_F_array_in_10), v27+int32(96))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v346 = int32(472)
	goto L29
L69:
	;
	v288 = F_errsave_start(m, v29)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v286 == int32(2147483647) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	if v288 == int32(0) {
		goto L28
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_array_in_11), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v346 = int32(485)
	goto L29
L76:
	;
	v302 = F_errsave_start(m, v29)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v317 = int32(0)
	v319 = v286 - v257
	if base.B2i32(v317 < v257)^base.B2i32(v319 < v286) == v317 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	if v302 == int32(0) {
		goto L28
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = int32(2147483647)
	F_errmsg(m, int32(_a_F_array_in_12), v27-int32(-64))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v346 = int32(491)
	goto L29
L83:
	;
	v325 = v319 + int32(1)
	if v319 <= v325 {
		goto L24
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v328 = F_errsave_start(m, v29)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	if v328 == int32(0) {
		goto L28
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_in_13), v27+int32(80))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	v346 = int32(499)
	goto L29
L91:
	;
	goto L28
L92:
	;
	v500 = int32(16)
	v503 = F_palloc_mul(m, int32(8), v500)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L6
	} else {
		goto L123
	}
L93:
	;
	if v354 == int32(123) {
		v472 = v118
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v354 == int32(61) {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v359 = int32(0)
	v360 = F_errsave_start(m, v29)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	if v360 == int32(0) {
		v1616 = v359
		goto L23
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+416)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(416))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	v375 = F_errdetail(m, int32(_a_F_array_in_14), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(269), int32(_a_F_array_in_15))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	v1616 = v359
	goto L23
L103:
	;
	v384 = v118
	goto L106
L104:
	;
	goto L105
L105:
	;
	v446 = int32(0)
	v447 = F_errsave_start(m, v29)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L6
	} else {
		goto L117
	}
L106:
	;
	v409 = v384 + int32(1)
	v410 = int32(*(*int8)(unsafe.Add(mBase, uint32(v384)+1)))
	goto L108
L107:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v420 == int32(123) {
		v472 = v409
		goto L92
	} else {
		goto L110
	}
L108:
	;
	if base.B2i32(v410 == int32(32))|base.B2i32(base.Ui32((v410-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v384 = v409
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v423 = int32(0)
	v424 = F_errsave_start(m, v29)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	if v424 == int32(0) {
		v1616 = v423
		goto L23
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+432)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(432))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v439 = F_errdetail(m, int32(_a_F_array_in_16), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(289), int32(_a_F_array_in_15))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v1616 = v423
	goto L23
L117:
	;
	if v447 == int32(0) {
		v1616 = v446
		goto L23
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+464)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(464))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+448)) = int32(_a_F_array_in_17)
	v465 = F_errdetail(m, int32(_a_F_array_in_10), v27+int32(448))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(279), int32(_a_F_array_in_15))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	v1616 = v446
	goto L23
L123:
	;
	v507 = F_palloc_mul(m, int32(1), int32(16))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_initStringInfo(m, v27+int32(576))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	v514 = v472
	v518 = int32(0)
	v521 = base.B2i32(v121 != int32(0))
	v522 = v121
	v523 = v2
	v524 = v2
	v526 = v503
	v527 = v507
	v531 = v500
	goto L129
L126:
	;
	if v1209 != 0 {
		goto L334
	} else {
		goto L335
	}
L127:
	;
	v1616 = int32(0)
	goto L23
L128:
	;
	v1313 = F_errsave_start(m, v29)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L6
	} else {
		goto L320
	}
L129:
	;
	v539 = v27 + int32(576)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v541 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v541)
	*(*int32)(unsafe.Add(mBase, uint32(v539)+12)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v539)+4)) = v541
	goto L131
L130:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v27)+576))
	F_pfree(m, v1225)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L6
	} else {
		goto L308
	}
L131:
	;
	v547 = v514
	goto L137
L132:
	;
	if int32(0) < v1203 {
		v514 = v1199
		v518 = v1203
		v521 = v1206
		v522 = v1207
		v523 = v1208
		v524 = v1209
		v526 = v1211
		v527 = v1212
		v531 = v1216
		goto L129
	} else {
		goto L307
	}
L133:
	;
	if v523 != 0 {
		goto L276
	} else {
		goto L277
	}
L134:
	;
	v1087 = v837
	v1094 = int32(0)
	goto L133
L135:
	;
	if v523 != 0 {
		goto L267
	} else {
		goto L268
	}
L136:
	;
	v1035 = F_errsave_start(m, v29)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L6
	} else {
		goto L261
	}
L137:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	switch v571 - int32(123) {
	case 0:
		goto L141
	case 1:
		goto L139
	case 2:
		goto L140
	default:
		goto L142
	}
L138:
	;
	v837 = v547
	v844 = int32(0)
	v848 = int32(1)
	goto L212
L139:
	;
	v820 = base.I32_extend8_s(v571)
	if v820 == v102 {
		goto L135
	} else {
		goto L206
	}
L140:
	;
	v760 = v27 + int32(544) + v518<<(uint(int32(2))%32)
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v760-int32(4))))
	v764 = int32(0)
	if v523|base.B2i32(v763 <= v764) == v764 {
		goto L190
	} else {
		goto L191
	}
L141:
	;
	if v523 != 0 {
		goto L169
	} else {
		goto L170
	}
L142:
	;
	if v571 == int32(0) {
		goto L136
	} else {
		goto L143
	}
L143:
	;
	if v571 != int32(34) {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v580 = v547 + int32(1)
	goto L145
L145:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
	if v604 != int32(92) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v624 = v580
	goto L156
L147:
	;
	goto L146
L148:
	;
	F_appendStringInfoChar(m, v27+int32(576), base.I32_extend8_s(v616))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L155
	}
L149:
	;
	if v604 == int32(0) {
		goto L136
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)))
	if v612 == int32(0) {
		goto L136
	} else {
		goto L154
	}
L152:
	;
	if v604 == int32(34) {
		goto L147
	} else {
		goto L153
	}
L153:
	;
	v616 = v604
	v617 = int32(1)
	goto L148
L154:
	;
	v616 = v612
	v617 = int32(2)
	goto L148
L155:
	;
	v580 = v580 + v617
	goto L145
L156:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+1)))
	if v648 == int32(0) {
		goto L136
	} else {
		goto L158
	}
L157:
	;
	v669 = F_errsave_start(m, v29)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L6
	} else {
		goto L163
	}
L158:
	;
	v652 = v624 + int32(1)
	v653 = int32(0)
	v654 = base.I32_extend8_s(v648)
	if v654 == v102 {
		v1087 = v652
		v1094 = v653
		goto L133
	} else {
		goto L159
	}
L159:
	;
	switch v654&int32(255) - int32(123) {
	case 0, 2:
		v1087 = v652
		v1094 = v653
		goto L133
	default:
		goto L160
	}
L160:
	;
	goto L161
L161:
	;
	if base.B2i32(v654 == int32(32))|base.B2i32(base.Ui32((v654-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v624 = v652
		goto L156
	} else {
		goto L162
	}
L162:
	;
	goto L157
L163:
	;
	if v669 == int32(0) {
		goto L127
	} else {
		goto L164
	}
L164:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+400)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(400))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	v684 = F_errdetail(m, int32(_a_F_array_in_18), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(875), int32(_a_F_array_in_19))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	v1616 = int32(0)
	goto L23
L169:
	;
	v692 = F_errsave_start(m, v29)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(int32(6)) <= base.Ui32(v518) {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	if v692 == int32(0) {
		goto L127
	} else {
		goto L173
	}
L173:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+336)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(336))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+320)) = int32(123)
	v710 = F_errdetail(m, int32(_a_F_array_in_20), v27+int32(320))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(641), int32(_a_F_array_in_21))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v1616 = int32(0)
	goto L23
L178:
	;
	v720 = F_errsave_start(m, v29)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v740 = int32(1)
	v741 = v547 + v740
	v742 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(544)+v518<<(uint(int32(2))%32)))) = v742
	v751 = v518 + v740
	if v518 < v522 {
		goto L186
	} else {
		goto L187
	}
L181:
	;
	if v720 == int32(0) {
		goto L127
	} else {
		goto L182
	}
L182:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+352)) = int32(6)
	F_errmsg(m, int32(_a_F_array_in_5), v27+int32(352))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(648), int32(_a_F_array_in_21))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	v1616 = int32(0)
	goto L23
L186:
	;
	v1199 = v741
	v1203 = v751
	v1206 = v521
	v1207 = v522
	v1208 = v742
	v1209 = v524
	v1211 = v526
	v1212 = v527
	v1216 = v531
	goto L132
L187:
	;
	goto L188
L188:
	;
	if v521&int32(1) != 0 {
		goto L128
	} else {
		goto L189
	}
L189:
	;
	v1199 = v741
	v1203 = v751
	v1206 = int32(0)
	v1207 = v751
	v1208 = v742
	v1209 = v524
	v1211 = v526
	v1212 = v527
	v1216 = v531
	goto L132
L190:
	;
	v769 = F_errsave_start(m, v29)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L6
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v518) {
		goto L199
	} else {
		goto L200
	}
L193:
	;
	if v769 == int32(0) {
		goto L127
	} else {
		goto L194
	}
L194:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+384)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(384))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+368)) = int32(125)
	v787 = F_errdetail(m, int32(_a_F_array_in_20), v27+int32(368))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(674), int32(_a_F_array_in_21))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v1616 = int32(0)
	goto L23
L199:
	;
	v798 = v760 - int32(8)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	*(*int32)(unsafe.Add(mBase, uint32(v798))) = v799 + int32(1)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v804 = int32(1)
	v805 = v547 + v804
	v809 = v518 - v804
	v812 = v27 + int32(512) + v809<<(uint(int32(2))%32)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	if v813 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v812))) = v763
	v1199 = v805
	v1203 = v809
	v1206 = v521
	v1207 = v522
	v1208 = int32(1)
	v1209 = v524
	v1211 = v526
	v1212 = v527
	v1216 = v531
	goto L132
L203:
	;
	goto L204
L204:
	;
	if v813 != v763 {
		goto L128
	} else {
		goto L205
	}
L205:
	;
	v1199 = v805
	v1203 = v809
	v1206 = v521
	v1207 = v522
	v1208 = int32(1)
	v1209 = v524
	v1211 = v526
	v1212 = v527
	v1216 = v531
	goto L132
L206:
	;
	goto L208
L207:
	;
	goto L138
L208:
	;
	if base.B2i32(v820 == int32(32))|base.B2i32(base.Ui32((v820-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	goto L207
L210:
	;
	goto L211
L211:
	;
	v547 = v547 + int32(1)
	goto L137
L212:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if base.Ui32(v861) <= base.Ui32(int32(91)) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v27)+576))
	v959 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v957+v844))) = uint8(v959)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+580)) = v844
	v963 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_array_in[0])))
	if v963&v848 == v959 {
		goto L134
	} else {
		goto L246
	}
L214:
	;
	v932 = base.I32_extend8_s(v861)
	if base.B2i32(v932 == v102)|base.B2i32(v932 == int32(125)) == int32(0) {
		goto L238
	} else {
		goto L239
	}
L215:
	;
	if v861 == int32(0) {
		goto L136
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	if v861 != int32(92) {
		goto L226
	} else {
		goto L227
	}
L218:
	;
	if v861 != int32(34) {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	v868 = F_errsave_start(m, v29)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L6
	} else {
		goto L220
	}
L220:
	;
	if v868 == int32(0) {
		goto L127
	} else {
		goto L221
	}
L221:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L6
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(304))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L223
	}
L223:
	;
	v883 = F_errdetail(m, int32(_a_F_array_in_18), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(909), int32(_a_F_array_in_19))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	v1616 = int32(0)
	goto L23
L226:
	;
	if v861 != int32(123) {
		goto L214
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v921 = int32(*(*int8)(unsafe.Add(mBase, uint32(v837)+1)))
	if v921 == int32(0) {
		goto L136
	} else {
		goto L236
	}
L229:
	;
	v895 = F_errsave_start(m, v29)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	if v895 == int32(0) {
		goto L127
	} else {
		goto L231
	}
L231:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(288))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = int32(123)
	v913 = F_errdetail(m, int32(_a_F_array_in_20), v27+int32(272))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(903), int32(_a_F_array_in_19))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	v1616 = int32(0)
	goto L23
L236:
	;
	F_appendStringInfoChar(m, v27+int32(576), v921)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L6
	} else {
		goto L237
	}
L237:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v27)+580))
	v837 = v837 + int32(2)
	v844 = v931
	v848 = int32(0)
	goto L212
L238:
	;
	F_appendStringInfoChar(m, v27+int32(576), v932)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L6
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	goto L213
L241:
	;
	v943 = int32(*(*int8)(unsafe.Add(mBase, uint32(v837))))
	goto L242
L242:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v27)+580))
	if base.B2i32(v943 == int32(32))|base.B2i32(base.Ui32((v943-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v954 = v844
	goto L245
L244:
	;
	v954 = v953
	goto L245
L245:
	;
	v837 = v837 + int32(1)
	v844 = v954
	goto L212
L246:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v27)+576))
	v971 = v967
	v972 = int32(_a_F_array_in_22)
	goto L248
L247:
	;
	if v1009 != 0 {
		goto L134
	} else {
		goto L260
	}
L248:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	if v975 == v976 {
		v998 = v975
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v1009 = int32(0)
	goto L247
L250:
	;
	v1000 = int32(1)
	if v998 != 0 {
		v971 = v971 + v1000
		v972 = v972 + v1000
		goto L248
	} else {
		goto L259
	}
L251:
	;
	if base.Ui32((v975-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v986 = v975 | int32(32)
	goto L254
L253:
	;
	v986 = v975
	goto L254
L254:
	;
	if base.Ui32((v976-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v995 = v976 | int32(32)
	goto L257
L256:
	;
	v995 = v976
	goto L257
L257:
	;
	if v986 == v995 {
		v998 = v986
		goto L250
	} else {
		goto L258
	}
L258:
	;
	v1009 = v986 - v995
	goto L247
L259:
	;
	goto L249
L260:
	;
	v1087 = v837
	v1094 = int32(1)
	goto L133
L261:
	;
	if v1035 == int32(0) {
		goto L127
	} else {
		goto L262
	}
L262:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(256))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L6
	} else {
		goto L264
	}
L264:
	;
	v1050 = F_errdetail(m, int32(_a_F_array_in_23), int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L6
	} else {
		goto L265
	}
L265:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(946), int32(_a_F_array_in_19))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L6
	} else {
		goto L266
	}
L266:
	;
	v1616 = int32(0)
	goto L23
L267:
	;
	v1199 = v547 + int32(1)
	v1203 = v518
	v1206 = v521
	v1207 = v522
	v1208 = int32(0)
	v1209 = v524
	v1211 = v526
	v1212 = v527
	v1216 = v531
	goto L132
L268:
	;
	goto L269
L269:
	;
	v1061 = F_errsave_start(m, v29)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L6
	} else {
		goto L270
	}
L270:
	;
	if v1061 == int32(0) {
		goto L127
	} else {
		goto L271
	}
L271:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(176))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L6
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v102
	v1078 = F_errdetail(m, int32(_a_F_array_in_20), v27+int32(160))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L6
	} else {
		goto L274
	}
L274:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(709), int32(_a_F_array_in_21))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L6
	} else {
		goto L275
	}
L275:
	;
	v1616 = int32(0)
	goto L23
L276:
	;
	v1111 = F_errsave_start(m, v29)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L6
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	if v531 <= v524 {
		goto L285
	} else {
		goto L286
	}
L279:
	;
	if v1111 == int32(0) {
		goto L127
	} else {
		goto L280
	}
L280:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L6
	} else {
		goto L281
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(192))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L6
	} else {
		goto L282
	}
L282:
	;
	v1126 = F_errdetail(m, int32(_a_F_array_in_24), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L6
	} else {
		goto L283
	}
L283:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(723), int32(_a_F_array_in_21))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L6
	} else {
		goto L284
	}
L284:
	;
	v1616 = int32(0)
	goto L23
L285:
	;
	if base.Ui32(int32(134217727)) <= base.Ui32(v531) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	v1170 = v526
	v1171 = v527
	v1172 = v531
	goto L287
L287:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v27)+576))
	if v1094 != 0 {
		goto L301
	} else {
		goto L302
	}
L288:
	;
	v1137 = F_errsave_start(m, v29)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L6
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v1158 = int32(134217727)
	v1160 = v531 << (uint(int32(1)) % 32)
	if base.Ui32(v1158) <= base.Ui32(v1160) {
		goto L296
	} else {
		goto L297
	}
L291:
	;
	if v1137 == int32(0) {
		goto L127
	} else {
		goto L292
	}
L292:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L6
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_in_13), v27+int32(240))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(732), int32(_a_F_array_in_21))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L6
	} else {
		goto L295
	}
L295:
	;
	v1616 = int32(0)
	goto L23
L296:
	;
	v1163 = v1158
	goto L298
L297:
	;
	v1163 = v1160
	goto L298
L298:
	;
	v1164 = F_repalloc_mul(m, v526, int32(8), v1163)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L6
	} else {
		goto L299
	}
L299:
	;
	v1167 = F_repalloc_mul(m, v527, int32(1), v1163)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L6
	} else {
		goto L300
	}
L300:
	;
	v1170 = v1164
	v1171 = v1167
	v1172 = v1163
	goto L287
L301:
	;
	v1175 = int32(0)
	goto L303
L302:
	;
	v1175 = v1174
	goto L303
L303:
	;
	v1179 = F_InputFunctionCallSafe(m, v76+int32(20), v1175, v101, v30, v29, v1170+v524<<(uint(int32(3))%32))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L6
	} else {
		goto L304
	}
L304:
	;
	if v1179 == int32(0) {
		goto L127
	} else {
		goto L305
	}
L305:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v524+v1171))) = uint8(v1094)
	if v518 != v522 {
		goto L128
	} else {
		goto L306
	}
L306:
	;
	v1186 = int32(1)
	v1191 = v518<<(uint(int32(2))%32) + v27 + int32(540)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1191)))
	*(*int32)(unsafe.Add(mBase, uint32(v1191))) = v1192 + v1186
	v1199 = v1087
	v1203 = v518
	v1206 = v1186
	v1207 = v518
	v1208 = v1186
	v1209 = v524 + v1186
	v1211 = v1170
	v1212 = v1171
	v1216 = v1172
	goto L132
L307:
	;
	goto L130
L308:
	;
	v1228 = v1199
	goto L309
L309:
	;
	v1252 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1228))))
	if v1252 == int32(0) {
		goto L126
	} else {
		goto L311
	}
L310:
	;
	v1266 = int32(0)
	v1267 = F_errsave_start(m, v29)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L6
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	if base.B2i32(v1252 == int32(32))|base.B2i32(base.Ui32((v1252-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v1228 = v1228 + int32(1)
		goto L309
	} else {
		goto L313
	}
L313:
	;
	goto L310
L314:
	;
	if v1267 == int32(0) {
		v1616 = v1266
		goto L23
	} else {
		goto L315
	}
L315:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L6
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(144))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L6
	} else {
		goto L317
	}
L317:
	;
	v1282 = F_errdetail(m, int32(_a_F_array_in_25), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L6
	} else {
		goto L318
	}
L318:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(312), int32(_a_F_array_in_15))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L6
	} else {
		goto L319
	}
L319:
	;
	v1616 = v1266
	goto L23
L320:
	;
	if v121 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if v1313 == int32(0) {
		goto L127
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	if v1313 == int32(0) {
		goto L127
	} else {
		goto L329
	}
L324:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L6
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(208))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L6
	} else {
		goto L326
	}
L326:
	;
	v1328 = F_errdetail(m, int32(_a_F_array_in_26), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L6
	} else {
		goto L327
	}
L327:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(782), int32(_a_F_array_in_21))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L6
	} else {
		goto L328
	}
L328:
	;
	v1616 = int32(0)
	goto L23
L329:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L6
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = v32
	F_errmsg(m, int32(_a_F_array_in_6), v27+int32(224))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L6
	} else {
		goto L331
	}
L331:
	;
	v1349 = F_errdetail(m, int32(_a_F_array_in_27), int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L6
	} else {
		goto L332
	}
L332:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(787), int32(_a_F_array_in_21))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L6
	} else {
		goto L333
	}
L333:
	;
	goto L127
L334:
	;
	v1381 = int32(0)
	if v1209 <= v1381 {
		v1512 = v1381
		goto L338
	} else {
		goto L339
	}
L335:
	;
	goto L336
L336:
	;
	v1601 = F_palloc0(m, int32(16))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L6
	} else {
		goto L382
	}
L337:
	;
	v1565 = v1564 + v1542
	v1566 = F_palloc0(m, v1565)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L6
	} else {
		goto L372
	}
L338:
	;
	v1542 = v1512
	v1559 = v1381
	v1564 = (v1207<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L337
L339:
	;
	v1385 = int32(0)
	v1391 = v1385
	v1393 = v1381
	v1395 = v1385
	goto L340
L340:
	;
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391+v1212))))
	if v1416 != 0 {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	if v1489&int32(1) == int32(0) {
		v1512 = v1488
		goto L338
	} else {
		goto L371
	}
L342:
	;
	v1493 = v1391 + int32(1)
	if v1493 != v1209 {
		v1391 = v1493
		v1393 = v1488
		v1395 = v1489
		goto L340
	} else {
		goto L370
	}
L343:
	;
	v1488 = v1393
	v1489 = int32(1)
	goto L342
L344:
	;
	goto L345
L345:
	;
	if v78 != int32(-1) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1465 = (v1393 + (v100 - int32(1)) + v1461) & (v1385 - v100)
	if base.Ui32(v1465) < base.Ui32(int32(1073741824)) {
		v1488 = v1465
		v1489 = v1395
		goto L342
	} else {
		goto L364
	}
L347:
	;
	if int32(0) < v78 {
		v1461 = v78
		goto L346
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1431 = v1211 + v1391<<(uint(int32(3))%32)
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1431)))
	v1433 = F_pg_detoast_datum(m, v1432)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L6
	} else {
		goto L351
	}
L350:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1211+v1391<<(uint(int32(3))%32))))
	v1426 = F_strlen(m, v1425)
	mBase = m.M
	v1461 = v1426 + int32(1)
	goto L346
L351:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1431))) = base.I64_extend_i32_u(v1433)
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433))))
	if v1437 == int32(1) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1441 = int32(18)
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433)+1)))
	if v1443 == v1441 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	goto L354
L354:
	;
	if v1437&int32(1) != 0 {
		goto L361
	} else {
		goto L362
	}
L355:
	;
	v1446 = v1441
	goto L357
L356:
	;
	v1446 = int32(2)
	goto L357
L357:
	;
	if base.Ui32((v1443-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1453 = int32(6)
	goto L360
L359:
	;
	v1453 = v1446
	goto L360
L360:
	;
	v1461 = v1453
	goto L346
L361:
	;
	v1461 = int32(base.Ui32(v1437) >> (uint(int32(1)) % 32))
	goto L346
L362:
	;
	goto L363
L363:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1433)))
	v1461 = int32(base.Ui32(v1458) >> (uint(int32(2)) % 32))
	goto L346
L364:
	;
	v1468 = int32(0)
	v1469 = F_errsave_start(m, v29)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L6
	} else {
		goto L365
	}
L365:
	;
	if v1469 == int32(0) {
		v1616 = v1468
		goto L23
	} else {
		goto L366
	}
L366:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L6
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = int32(1073741823)
	F_errmsg(m, int32(_a_F_array_in_13), v27+int32(128))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L6
	} else {
		goto L368
	}
L368:
	;
	F_errsave_finish(m, v29, int32(_a_F_array_in_3), int32(340), int32(_a_F_array_in_15))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L6
	} else {
		goto L369
	}
L369:
	;
	v1616 = v1468
	goto L23
L370:
	;
	goto L341
L371:
	;
	v1502 = base.I32_div_s(v1209+int32(7), int32(8))
	v1509 = (v1502 + v1207<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v1542 = v1488
	v1559 = v1509
	v1564 = v1509
	goto L337
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1566)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1566)+8)) = v1559
	*(*int32)(unsafe.Add(mBase, uint32(v1566)+4)) = v1207
	v1571 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1566))) = v1565 << (uint(v1571) % 32)
	v1575 = v1566 + int32(16)
	v1577 = v1207 << (uint(v1571) % 32)
	v1578 = int32(0)
	v1579 = base.B2i32(v1577 == v1578)
	if v1579 == v1578 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	base.MemoryCopy(m, v1575, v27+int32(512), v1577)
	goto L375
L374:
	;
	goto L375
L375:
	;
	if v1579 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	base.MemoryCopy(m, v1575+v1577, v27+int32(480), v1577)
	goto L378
L377:
	;
	goto L378
L378:
	;
	v1591 = int32(1)
	F_CopyArrayEls(m, v1566, v1211, v1212, v1209, v78, v77&v1591, v80, v1591)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L6
	} else {
		goto L379
	}
L379:
	;
	F_pfree(m, v1211)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L6
	} else {
		goto L380
	}
L380:
	;
	F_pfree(m, v1212)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L6
	} else {
		goto L381
	}
L381:
	;
	v1616 = v1566
	goto L23
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v1601)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1601))) = int64(64)
	v1616 = v1601
	goto L23
}
func F_array_ndims(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
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
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_DatumGetAnyArrayP(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
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
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v15)
		}
	}
}
func F_array_ne(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2 ^ int64(1)
	}
}
func F_array_set_element(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
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
	var v186 int32
	_ = v186
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1245 int32
	_ = v1245
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1647 int32
	_ = v1647
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1884 int32
	_ = v1884
	var v1892 int32
	_ = v1892
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1909 int32
	_ = v1909
	var v1920 int32
	_ = v1920
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1945 int64
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v2006 int32
	_ = v2006
	var v2018 int32
	_ = v2018
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2091 int32
	_ = v2091
	var v2103 int32
	_ = v2103
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2207 int32
	_ = v2207
	var v2210 int32
	_ = v2210
	var v2215 int32
	_ = v2215
	var v2247 int32
	_ = v2247
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2287 int32
	_ = v2287
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2330 int32
	_ = v2330
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2390 int32
	_ = v2390
	var v2395 int32
	_ = v2395
	v5 = l4
	v10 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(144)
	m.G0 = v33
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+71)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+72)) = l3
	switch l8 - int32(99) {
	case 0:
		v58 = int32(1)
		goto L1
	case 1:
		goto L4
	default:
		goto L3
	case 6:
		goto L5
	case 16:
		goto L2
	}
L1:
	;
	if int32(0) < l5 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v58 = int32(2)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v58 = int32(8)
	goto L1
L5:
	;
	v58 = int32(4)
	goto L1
L6:
	;
	return int64(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = l8
	F_errmsg_internal(m, int32(_a_F_array_set_element_0), v33)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_array_set_element_1), int32(322), int32(_a_F_array_set_element_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L6
	} else {
		goto L477
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L6
	} else {
		goto L473
	}
L12:
	;
	m.G0 = v33 + int32(144)
	return base.I64_extend_i32_u(v2330)
L13:
	;
	v1721 = v33 + int32(112)
	v1723 = v33 + int32(80)
	v1724 = int32(0)
	v1733 = l1 - int32(1)
	if v1733 < v1724 {
		v1811 = v1724
		goto L385
	} else {
		goto L386
	}
L14:
	;
	v1681 = v33 + int32(112)
	v1682 = F_ArrayGetNItemsSafe(m, l1, v1681)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L6
	} else {
		goto L382
	}
L15:
	;
	v1647 = int32(0)
	if v1624 == v1647 {
		v1702 = v1630
		v1704 = v10
		v1705 = v1633
		v1719 = v1647
		goto L13
	} else {
		goto L381
	}
L16:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	if v1573 <= v1572 {
		goto L372
	} else {
		goto L373
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1547))) = uint8(v1542)
	v2330 = v962
	goto L12
L18:
	;
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	v1469 = v1031
	v1470 = int32(1)
	v1473 = v1027
	v1474 = v1030
	v1480 = v709
	v1482 = v1466
	v1485 = v726
	goto L355
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L6
	} else {
		goto L351
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L6
	} else {
		goto L347
	}
L21:
	;
	v683 = v33 + int32(112)
	v684 = F_ArrayGetNItemsSafe(m, l1, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L160
	}
L22:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	if v606 <= v605 {
		goto L149
	} else {
		goto L150
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L144
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L140
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L136
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L132
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L128
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L6
	} else {
		goto L124
	}
L29:
	;
	if l1 != int32(1) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(l1-int32(7)) <= base.Ui32(int32(-7)) {
		goto L25
	} else {
		goto L41
	}
L32:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v63 < int32(0) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v66 = base.I32_div_s(l5, l6)
	if v66 <= v63 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v5 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v68 = F_palloc(m, l5)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if l5 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	base.MemoryCopy(m, v68, base.I32_wrap_i64(l0), l5)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v75 = F_ArrayCastAndSet(m, l3, l6, l7, v58, v68+v72*l6)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v2330 = v68
	goto L12
L41:
	;
	if v5|base.B2i32(l6 != int32(-1)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v87 = F_pg_detoast_datum(m, base.I32_wrap_i64(l3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	v91 = l3
	goto L44
L44:
	;
	v92 = base.I32_wrap_i64(l0)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v89 = base.I64_extend_i32_u(v87)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+72)) = v89
	v91 = v89
	goto L44
L46:
	;
	v238 = F_pg_detoast_datum(m, v92)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L85
	}
L47:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v96&int32(254) != int32(2) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v101 = F_DatumGetExpandedArray(m, l0)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v105 = v103 << (uint(int32(2)) % 32)
	v106 = int32(0)
	v107 = base.B2i32(v105 == v106)
	if v107 == v106 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v101)+32))
	base.MemoryCopy(m, v33+int32(112), v112, v105)
	goto L52
L51:
	;
	goto L52
L52:
	;
	if v107 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	base.MemoryCopy(m, v33+int32(80), v118, v105)
	goto L55
L54:
	;
	goto L55
L55:
	;
	if v103 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	F_deconstruct_expanded_array(m, v101)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L67
	}
L57:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v124 = l1 << (uint(int32(2)) % 32)
	v125 = F_MemoryContextAllocZero(m, v122, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if l1 != v103 {
		goto L24
	} else {
		goto L66
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+32)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v129 = F_MemoryContextAllocZero(m, v128, v124)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+36)) = v129
	v132 = int32(0)
	v133 = base.B2i32(v124 == v132)
	if v133 == v132 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryFill(m, v33+int32(112), int32(0), v124)
	goto L64
L63:
	;
	goto L64
L64:
	;
	if v124 == v132 {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	base.MemoryCopy(m, v33+int32(80), l2, v124)
	goto L56
L66:
	;
	goto L56
L67:
	;
	if v5 != 0 {
		v162 = v91
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v164 = int32(0)
	v165 = base.B2i32(v103 == v164)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v101)+52))
	v169 = v5 | base.B2i32(v166 != v164)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v101)+48))
	if l1 == int32(1) {
		goto L16
	} else {
		goto L72
	}
L69:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+46)))
	if v148&int32(1) != 0 {
		v162 = v91
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v151 = int32(_a_F_array_set_element_3)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v154
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+44)))
	v158 = F_datumCopy(m, v91, int32(0), v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_array_set_element[0])) = v152
	v162 = v158
	goto L68
L72:
	;
	v186 = v10
	goto L73
L73:
	;
	v204 = v186 << (uint(int32(2)) % 32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2+v204)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(80)+v204)))
	if v210 <= v206 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v1624 = v165
	v1630 = int32(0)
	v1633 = v169
	goto L15
L75:
	;
	v235 = v186 + int32(1)
	if v235 != l1 {
		v186 = v235
		goto L73
	} else {
		goto L84
	}
L76:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(112)+v204)))
	if v206 < v215+v210 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_array_set_element_4), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2658), int32(_a_F_array_set_element_6))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	goto L74
L85:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v240 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v245 = l1 << (uint(int32(2)) % 32)
	if v245 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if l1 != v240 {
		goto L23
	} else {
		goto L104
	}
L89:
	;
	base.MemoryCopy(m, v33+int32(80), l2, v245)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v249 = int32(0)
	if base.Ui32(int32(7)) <= base.Ui32(l1-int32(1)) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v414 = F_construct_md_array(m, v33+int32(72), v33+int32(71), l1, v33+int32(112), v33+int32(80), v243, l6, l7, l8)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L103
	}
L93:
	;
	v260 = v249
	v268 = v10
	goto L96
L94:
	;
	v309 = v249
	goto L95
L95:
	;
	v336 = v249
	v339 = v309
	goto L100
L96:
	;
	v289 = v33 + int32(112) + v260<<(uint(int32(2))%32)
	v290 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v289)+24)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v289)+16)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v289)+8)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = v290
	v298 = int32(8)
	v299 = v260 + v298
	v301 = v268 + v298
	if v301 != 0 {
		v260 = v299
		v268 = v301
		goto L96
	} else {
		goto L98
	}
L97:
	;
	if l1 == int32(0) {
		goto L92
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v309 = v299
	goto L95
L100:
	;
	v369 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(112)+v339<<(uint(int32(2))%32)))) = v369
	v374 = v336 + v369
	if v374 != l1 {
		v336 = v374
		v339 = v339 + v369
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L92
L102:
	;
	goto L101
L103:
	;
	v2330 = v414
	goto L12
L104:
	;
	v418 = v238 + int32(16)
	v420 = l1 << (uint(int32(2)) % 32)
	v421 = int32(0)
	v422 = base.B2i32(v420 == v421)
	if v422 == v421 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	base.MemoryCopy(m, v33+int32(112), v418, v420)
	goto L107
L106:
	;
	goto L107
L107:
	;
	if v422 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	base.MemoryCopy(m, v33+int32(80), v418+v432<<(uint(int32(2))%32), v420)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v437 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v441 = v5 | base.B2i32(v438 != v437)
	if l1 == int32(1) {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v449 = v437
	goto L112
L112:
	;
	v475 = v449 << (uint(int32(2)) % 32)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l2+v475)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(80)+v475)))
	if v481 <= v477 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v665 = v505
	v670 = v10
	v671 = v441
	goto L21
L114:
	;
	v505 = int32(1)
	v507 = v449 + v505
	if v507 != l1 {
		v449 = v507
		goto L112
	} else {
		goto L123
	}
L115:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(112)+v475)))
	if v477 < v486+v481 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_array_set_element_4), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2379), int32(_a_F_array_set_element_7))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	goto L113
L124:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_array_set_element_8), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2250), int32(_a_F_array_set_element_7))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
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
	v531 = m.ExcPending
	if v531 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_array_set_element_4), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2255), int32(_a_F_array_set_element_7))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L6
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_array_set_element_9), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2260), int32(_a_F_array_set_element_7))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
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
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_array_set_element_8), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2272), int32(_a_F_array_set_element_7))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L6
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_array_set_element_8), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2575), int32(_a_F_array_set_element_6))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_array_set_element_8), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2321), int32(_a_F_array_set_element_7))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v630 = v625 + v627
	if v605 < v630 {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v625 = v606
	v627 = v608
	v628 = v10
	v629 = v441
	goto L148
L150:
	;
	goto L151
L151:
	;
	v609 = v606 - v605
	if base.B2i32(v609 < v606)^base.B2i32(int32(0) < v605) != 0 {
		goto L20
	} else {
		goto L152
	}
L152:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v615 = v614 + v609
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v615
	if base.B2i32(v609 < int32(0)) != base.B2i32(v615 < v614) {
		goto L20
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v605
	v625 = v605
	v627 = v615
	v628 = v609
	v629 = base.B2i32(int32(1) < v609) | v441
	goto L148
L154:
	;
	v665 = int32(1)
	v670 = v628
	v671 = v629
	goto L21
L155:
	;
	goto L156
L156:
	;
	v635 = v605 - v630
	if base.B2i32(int32(0) < v630)^base.B2i32(v635 < v605) != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	v639 = v635 + int32(1)
	if v639 < v635 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	v641 = v639 + v627
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v641
	if base.B2i32(v639 < int32(0)) != base.B2i32(v641 < v627) {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	v665 = base.B2i32(v639 == int32(0))
	v670 = v628
	v671 = base.B2i32(int32(1) < v639) | v629
	goto L21
L160:
	;
	F_ArrayCheckBounds(m, l1, v683, v33+int32(80))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v691 = l1 << (uint(int32(3)) % 32)
	if v671&int32(1) != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v709 = F_ArrayGetNItemsSafe(m, l1, v418)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L6
	} else {
		goto L166
	}
L163:
	;
	v697 = base.I32_div_s(v684+int32(7), int32(8))
	v702 = (v691 + v697 + int32(23)) & int32(-8)
	v707 = v702
	v708 = v702
	goto L162
L164:
	;
	goto L165
L165:
	;
	v707 = v10
	v708 = (v691 + int32(23)) & int32(120)
	goto L162
L166:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v717 = v715 << (uint(int32(3)) % 32)
	if v714 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v722 = v714
	goto L169
L168:
	;
	v722 = (v717 + int32(23)) & int32(-8)
	goto L169
L169:
	;
	v723 = int32(base.Ui32(v711)>>(uint(int32(2))%32)) - v722
	v724 = v717 + v418
	if v714 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v726 = v724
	goto L172
L171:
	;
	v726 = int32(0)
	goto L172
L172:
	;
	if v670 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	if v5 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L174:
	;
	v907 = int32(0)
	v909 = v10
	v910 = v723
	v911 = v10
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v665 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v907 = v723
	v909 = v709
	v910 = int32(0)
	v911 = v10
	goto L173
L178:
	;
	goto L179
L179:
	;
	v732 = v33 + int32(112)
	v734 = v33 + int32(80)
	v735 = int32(0)
	v744 = l1 - int32(1)
	if v744 < v735 {
		v822 = v735
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v828 != 0 {
		goto L190
	} else {
		goto L191
	}
L181:
	;
	goto L180
L182:
	;
	v747 = int32(1)
	if v744 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v756 = v744
	v757 = v747
	v758 = v735
	v763 = v735
	goto L186
L184:
	;
	v799 = v744
	v800 = v747
	v801 = v735
	goto L185
L185:
	;
	v808 = v799 << (uint(int32(2)) % 32)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l2+v808)))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808+v734)))
	v822 = (v810-v812)*v800 + v801
	goto L181
L186:
	;
	v764 = int32(2)
	v765 = v756 << (uint(v764) % 32)
	v767 = v765 - int32(4)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l2+v767)))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v734+v767)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v765+v732)))
	v775 = v774 * v757
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v765+l2)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v765+v734)))
	v784 = (v769-v771)*v775 + ((v778-v780)*v757 + v758)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v732+v767)))
	v787 = v786 * v775
	v789 = v756 - v764
	v791 = v763 + v764
	if v791 != l1&int32(-2) {
		v756 = v789
		v757 = v787
		v758 = v784
		v763 = v791
		goto L186
	} else {
		goto L188
	}
L187:
	;
	if l1&int32(1) == int32(0) {
		v822 = v784
		goto L181
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v799 = v789
	v800 = v787
	v801 = v784
	goto L185
L190:
	;
	v836 = v828
	goto L192
L191:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v836 = (v829<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L192
L192:
	;
	v839 = F_array_seek(m, v836+v238, int32(0), v726, v822, l6, l8)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v841 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v849 = v841
	goto L196
L195:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v849 = (v842<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L196
L196:
	;
	v851 = v839 - (v849 + v238)
	if v714 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v907 = v851
	v909 = v822
	v910 = v723 - (v851 + v903)
	v911 = v903
	goto L173
L198:
	;
	if int32(0) < l6 {
		v895 = l6
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v855 = base.I32_div_s(v822, int32(8))
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+v855))))
	if int32(base.Ui32(v857)>>(uint(v822&int32(7))%32))&int32(1) != 0 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v903 = int32(0)
	goto L197
L201:
	;
	v903 = (v895 + v58 - int32(1)) & (int32(0) - v58)
	goto L197
L202:
	;
	if l6 == int32(-1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if v868 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v892 = F_strlen(m, v839)
	mBase = m.M
	v895 = v892 + int32(1)
	goto L201
L206:
	;
	v872 = int32(18)
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839)+1)))
	if v874 == v872 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L208
L208:
	;
	if v868&int32(1) != 0 {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	v877 = v872
	goto L211
L210:
	;
	v877 = int32(2)
	goto L211
L211:
	;
	if base.Ui32((v874-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v884 = int32(6)
	goto L214
L213:
	;
	v884 = v877
	goto L214
L214:
	;
	v895 = v884
	goto L201
L215:
	;
	v895 = int32(base.Ui32(v868) >> (uint(int32(1)) % 32))
	goto L201
L216:
	;
	goto L217
L217:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	v895 = int32(base.Ui32(v889) >> (uint(int32(2)) % 32))
	goto L201
L218:
	;
	if int32(0) < l6 {
		v948 = l6
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v958 = v10
	goto L220
L220:
	;
	v961 = v907 + v708 + v910 + v958
	v962 = F_palloc0(m, v961)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L6
	} else {
		goto L238
	}
L221:
	;
	v958 = (v948 + v58 - int32(1)) & (int32(0) - v58)
	goto L220
L222:
	;
	v916 = base.I32_wrap_i64(v91)
	if l6 == int32(-1) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	if v919 == int32(1) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v943 = F_strlen(m, v916)
	mBase = m.M
	v948 = v943 + int32(1)
	goto L221
L226:
	;
	v923 = int32(18)
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916)+1)))
	if v925 == v923 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	if v919&int32(1) != 0 {
		goto L235
	} else {
		goto L236
	}
L229:
	;
	v928 = v923
	goto L231
L230:
	;
	v928 = int32(2)
	goto L231
L231:
	;
	if base.Ui32((v925-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v935 = int32(6)
	goto L234
L233:
	;
	v935 = v928
	goto L234
L234:
	;
	v948 = v935
	goto L221
L235:
	;
	v948 = int32(base.Ui32(v919) >> (uint(int32(1)) % 32))
	goto L221
L236:
	;
	goto L237
L237:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v916)))
	v948 = int32(base.Ui32(v940) >> (uint(int32(2)) % 32))
	goto L221
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v962)+8)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v962)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v962))) = v961 << (uint(int32(2)) % 32)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v962)+12)) = v969
	v972 = v962 + int32(16)
	if v422 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	base.MemoryCopy(m, v972, v33+int32(112), v420)
	goto L241
L240:
	;
	goto L241
L241:
	;
	if v422 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	base.MemoryCopy(m, v972+v420, v33+int32(80), v420)
	goto L244
L243:
	;
	goto L244
L244:
	;
	v984 = v238 + v722
	v985 = v962 + v708
	if v907 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	base.MemoryCopy(m, v985, v984, v907)
	goto L247
L246:
	;
	goto L247
L247:
	;
	if v5 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v990 = F_ArrayCastAndSet(m, v91, l6, l7, v58, v985+v907)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L6
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	if v910 != 0 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L250
L252:
	;
	base.MemoryCopy(m, v985+v907+v958, v907+v984+v911, v910)
	goto L254
L253:
	;
	goto L254
L254:
	;
	if v671&int32(1) == int32(0) {
		v2330 = v962
		goto L12
	} else {
		goto L255
	}
L255:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	v1004 = v972 + v1001<<(uint(int32(3))%32)
	if v665 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1007 = v909
	goto L258
L257:
	;
	v1007 = v684 - int32(1)
	goto L258
L258:
	;
	v1009 = base.I32_div_s(v1007, int32(8))
	v1010 = v1004 + v1009
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010))))
	v1014 = v1007 & int32(7)
	if v5 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1020 = v1011 & base.I32_rotl(int32(-2), v1014)
	goto L261
L260:
	;
	v1020 = v1011 | int32(1)<<(uint(v1014)%32)
	goto L261
L261:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1010))) = uint8(v1020)
	if v670 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	if v709 <= int32(0) {
		v2330 = v962
		goto L12
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	if v909 <= int32(0) {
		goto L274
	} else {
		goto L275
	}
L265:
	;
	v1027 = int32(1) << (uint(v670&int32(7)) % 32)
	v1029 = base.I32_div_s(v670, int32(8))
	v1030 = v1004 + v1029
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030))))
	if v714 != 0 {
		goto L18
	} else {
		goto L266
	}
L266:
	;
	v1033 = v1031
	v1037 = v1027
	v1038 = v1030
	v1044 = v709
	goto L267
L267:
	;
	v1062 = v1033 | v1037
	v1063 = int32(1)
	v1064 = v1044 - v1063
	v1066 = v1037 << (uint(v1063) % 32)
	if v1066 == int32(256) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v1542 = v1062
	v1547 = v1038
	goto L17
L269:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1038))) = uint8(v1062)
	if v1064 == int32(0) {
		v2330 = v962
		goto L12
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1044) {
		v1033 = v1062
		v1037 = v1066
		v1044 = v1064
		goto L267
	} else {
		goto L273
	}
L272:
	;
	v1072 = int32(1)
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038)+1)))
	v1033 = v1073
	v1037 = v1072
	v1038 = v1038 + v1072
	v1044 = v1064
	goto L267
L273:
	;
	goto L268
L274:
	;
	if v665 == int32(0) {
		v2330 = v962
		goto L12
	} else {
		goto L315
	}
L275:
	;
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004))))
	if v714 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1245))) = uint8(v1232)
	goto L274
L277:
	;
	v1164 = v1080
	v1165 = v1004
	v1168 = v909
	goto L305
L278:
	;
	if v909 != int32(1) {
		goto L277
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	v1088 = int32(1)
	v1091 = v1080
	v1092 = v1088
	v1095 = v1088
	v1096 = v909
	v1097 = v726
	v1104 = v1004
	v1105 = v1087
	goto L282
L281:
	;
	v1232 = v1080 | int32(1)
	v1245 = v1004
	goto L276
L282:
	;
	if v1092&v1105 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	if v1140 != int32(1) {
		v1232 = v1139
		v1245 = v1141
		goto L276
	} else {
		goto L297
	}
L284:
	;
	v1125 = v1091 | v1095
	goto L286
L285:
	;
	v1125 = v1091 & (v1095 ^ int32(-1))
	goto L286
L286:
	;
	v1126 = int32(1)
	v1127 = v1096 - v1126
	v1129 = v1095 << (uint(v1126) % 32)
	if v1129 == int32(256) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1104))) = uint8(v1125)
	if v1127 == int32(0) {
		goto L274
	} else {
		goto L290
	}
L288:
	;
	v1139 = v1125
	v1140 = v1129
	v1141 = v1104
	goto L289
L289:
	;
	v1143 = v1092 << (uint(int32(1)) % 32)
	if v1143 == int32(256) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104)+1)))
	v1136 = int32(1)
	v1139 = v1135
	v1140 = v1136
	v1141 = v1104 + v1136
	goto L289
L291:
	;
	goto L283
L292:
	;
	if v1127 == int32(0) {
		goto L291
	} else {
		goto L295
	}
L293:
	;
	v1152 = v1143
	v1153 = v1097
	v1154 = v1105
	goto L294
L294:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1096) {
		v1091 = v1139
		v1092 = v1152
		v1095 = v1140
		v1096 = v1127
		v1097 = v1153
		v1104 = v1141
		v1105 = v1154
		goto L282
	} else {
		goto L296
	}
L295:
	;
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097)+1)))
	v1149 = int32(1)
	v1152 = v1149
	v1153 = v1097 + v1149
	v1154 = v1148
	goto L294
L296:
	;
	goto L291
L297:
	;
	goto L274
L298:
	;
	v1232 = v1230
	v1245 = v1165
	goto L276
L299:
	;
	v1230 = v1164 | int32(3)
	goto L298
L300:
	;
	v1230 = v1164 | int32(7)
	goto L298
L301:
	;
	v1230 = v1164 | int32(15)
	goto L298
L302:
	;
	v1230 = v1164 | int32(31)
	goto L298
L303:
	;
	v1230 = v1164 | int32(63)
	goto L298
L304:
	;
	v1230 = v1164 | int32(127)
	goto L298
L305:
	;
	if v1168 < int32(3) {
		goto L299
	} else {
		goto L307
	}
L306:
	;
	v1232 = v1211 | int32(1)
	v1245 = v1213
	goto L276
L307:
	;
	if v1168 == int32(3) {
		goto L300
	} else {
		goto L308
	}
L308:
	;
	if base.Ui32(v1168) < base.Ui32(int32(5)) {
		goto L301
	} else {
		goto L309
	}
L309:
	;
	if v1168 == int32(5) {
		goto L302
	} else {
		goto L310
	}
L310:
	;
	if base.Ui32(v1168) < base.Ui32(int32(7)) {
		goto L303
	} else {
		goto L311
	}
L311:
	;
	if v1168 == int32(7) {
		goto L304
	} else {
		goto L312
	}
L312:
	;
	v1205 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1165))) = uint8(v1205)
	v1208 = v1168 - int32(8)
	if v1208 == int32(0) {
		goto L274
	} else {
		goto L313
	}
L313:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165)+1)))
	v1212 = int32(1)
	v1213 = v1165 + v1212
	if v1208 != v1212 {
		v1164 = v1211
		v1165 = v1213
		v1168 = v1208
		goto L305
	} else {
		goto L314
	}
L314:
	;
	goto L306
L315:
	;
	v1295 = v909 + int32(1)
	v1298 = v709 + (v909 ^ int32(-1))
	if v1298 <= int32(0) {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v2330 = v962
	goto L12
L317:
	;
	goto L316
L318:
	;
	v1308 = int32(1) << (uint(v1295&int32(7)) % 32)
	v1310 = base.I32_div_s(v1295, int32(8))
	v1311 = v1004 + v1310
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311))))
	if v726 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1405))) = uint8(v1406)
	goto L317
L320:
	;
	v1315 = v1311
	v1316 = v1312
	v1319 = v1298
	v1320 = v1308
	goto L323
L321:
	;
	goto L322
L322:
	;
	v1350 = base.I32_div_s(v1295, int32(8))
	v1351 = v726 + v1350
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351))))
	v1353 = v1311
	v1354 = v1312
	v1356 = v1351
	v1357 = v1298
	v1358 = v1308
	v1359 = int32(1) << (uint(v1295&int32(7)) % 32)
	v1360 = v1352
	goto L331
L323:
	;
	v1324 = v1316 | v1320
	v1325 = int32(1)
	v1326 = v1319 - v1325
	v1328 = v1320 << (uint(v1325) % 32)
	if v1328 == int32(256) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	if v1340 != int32(1) {
		v1405 = v1338
		v1406 = v1339
		goto L319
	} else {
		goto L330
	}
L325:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1315))) = uint8(v1324)
	if v1326 == int32(0) {
		goto L317
	} else {
		goto L328
	}
L326:
	;
	v1338 = v1315
	v1339 = v1324
	v1340 = v1328
	goto L327
L327:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1319) {
		v1315 = v1338
		v1316 = v1339
		v1319 = v1326
		v1320 = v1340
		goto L323
	} else {
		goto L329
	}
L328:
	;
	v1334 = int32(1)
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315)+1)))
	v1338 = v1315 + v1334
	v1339 = v1335
	v1340 = v1334
	goto L327
L329:
	;
	goto L324
L330:
	;
	goto L317
L331:
	;
	if v1359&v1360 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	if v1383 == int32(1) {
		goto L317
	} else {
		goto L346
	}
L333:
	;
	v1367 = v1354 | v1358
	goto L335
L334:
	;
	v1367 = v1354 & (v1358 ^ int32(-1))
	goto L335
L335:
	;
	v1368 = int32(1)
	v1369 = v1357 - v1368
	v1371 = v1358 << (uint(v1368) % 32)
	if v1371 == int32(256) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1353))) = uint8(v1367)
	if v1369 == int32(0) {
		goto L317
	} else {
		goto L339
	}
L337:
	;
	v1381 = v1353
	v1382 = v1367
	v1383 = v1371
	goto L338
L338:
	;
	v1385 = v1359 << (uint(int32(1)) % 32)
	if v1385 == int32(256) {
		goto L341
	} else {
		goto L342
	}
L339:
	;
	v1377 = int32(1)
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1353)+1)))
	v1381 = v1353 + v1377
	v1382 = v1378
	v1383 = v1377
	goto L338
L340:
	;
	goto L332
L341:
	;
	if v1369 == int32(0) {
		goto L340
	} else {
		goto L344
	}
L342:
	;
	v1394 = v1356
	v1395 = v1385
	v1396 = v1360
	goto L343
L343:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1357) {
		v1353 = v1381
		v1354 = v1382
		v1356 = v1394
		v1357 = v1369
		v1358 = v1383
		v1359 = v1395
		v1360 = v1396
		goto L331
	} else {
		goto L345
	}
L344:
	;
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1356)+1)))
	v1391 = int32(1)
	v1394 = v1356 + v1391
	v1395 = v1391
	v1396 = v1390
	goto L343
L345:
	;
	goto L340
L346:
	;
	v1405 = v1381
	v1406 = v1382
	goto L319
L347:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L6
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_set_element_10), v33+int32(48))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L6
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2347), int32(_a_F_array_set_element_7))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L6
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L351:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L6
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_set_element_10), v33-int32(-64))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L6
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2362), int32(_a_F_array_set_element_7))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L6
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	if v1470&v1482 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v1518 == int32(1) {
		v2330 = v962
		goto L12
	} else {
		goto L370
	}
L357:
	;
	v1503 = v1469 | v1473
	goto L359
L358:
	;
	v1503 = v1469 & (v1473 ^ int32(-1))
	goto L359
L359:
	;
	v1504 = int32(1)
	v1505 = v1480 - v1504
	v1507 = v1473 << (uint(v1504) % 32)
	if v1507 == int32(256) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1474))) = uint8(v1503)
	if v1505 == int32(0) {
		v2330 = v962
		goto L12
	} else {
		goto L363
	}
L361:
	;
	v1517 = v1503
	v1518 = v1507
	v1519 = v1474
	goto L362
L362:
	;
	v1521 = v1470 << (uint(int32(1)) % 32)
	if v1521 == int32(256) {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1474)+1)))
	v1514 = int32(1)
	v1517 = v1513
	v1518 = v1514
	v1519 = v1474 + v1514
	goto L362
L364:
	;
	goto L356
L365:
	;
	if v1505 == int32(0) {
		goto L364
	} else {
		goto L368
	}
L366:
	;
	v1530 = v1521
	v1531 = v1482
	v1532 = v1485
	goto L367
L367:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1480) {
		v1469 = v1517
		v1470 = v1530
		v1473 = v1518
		v1474 = v1519
		v1480 = v1505
		v1482 = v1531
		v1485 = v1532
		goto L355
	} else {
		goto L369
	}
L368:
	;
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485)+1)))
	v1527 = int32(1)
	v1530 = v1527
	v1531 = v1526
	v1532 = v1485 + v1527
	goto L367
L369:
	;
	goto L364
L370:
	;
	v1542 = v1517
	v1547 = v1519
	goto L17
L371:
	;
	v1598 = v1594 + v1596
	if v1572 < v1598 {
		v1624 = v1593
		v1630 = v1595
		v1633 = v1597
		goto L15
	} else {
		goto L377
	}
L372:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v1593 = v165
	v1594 = v1575
	v1595 = v10
	v1596 = v1573
	v1597 = v169
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1576 = v1573 - v1572
	if base.B2i32(v1576 < v1573)^base.B2i32(int32(0) < v1572) != 0 {
		goto L11
	} else {
		goto L375
	}
L375:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v1582 = v1581 + v1576
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1582
	if base.B2i32(v1576 < int32(0)) != base.B2i32(v1582 < v1581) {
		goto L11
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1572
	v1589 = int32(1)
	v1593 = v1589
	v1594 = v1582
	v1595 = v1576
	v1596 = v1572
	v1597 = base.B2i32(v1589 < v1576) | v169
	goto L371
L377:
	;
	v1602 = v1572 - v1598
	if base.B2i32(int32(0) < v1598)^base.B2i32(v1602 < v1572) != 0 {
		goto L10
	} else {
		goto L378
	}
L378:
	;
	v1606 = v1602 + int32(1)
	if v1606 < v1602 {
		goto L10
	} else {
		goto L379
	}
L379:
	;
	v1608 = v1594 + v1606
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1608
	if base.B2i32(v1606 < int32(0)) != base.B2i32(v1608 < v1594) {
		goto L10
	} else {
		goto L380
	}
L380:
	;
	v1662 = v1606
	v1663 = v1595
	v1666 = base.B2i32(int32(1) < v1606) | v1597
	goto L14
L381:
	;
	v1662 = v1647
	v1663 = v1630
	v1666 = v1633
	goto L14
L382:
	;
	F_ArrayCheckBounds(m, l1, v1681, v33+int32(80))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L6
	} else {
		goto L383
	}
L383:
	;
	v1702 = v1663
	v1704 = v1662
	v1705 = v1666
	v1719 = int32(1)
	goto L13
L384:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v101)+56))
	if v1818 < v1817 {
		goto L394
	} else {
		goto L395
	}
L385:
	;
	goto L384
L386:
	;
	v1736 = int32(1)
	if v1733 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1745 = v1733
	v1746 = v1736
	v1747 = v1724
	v1752 = v1724
	goto L390
L388:
	;
	v1788 = v1733
	v1789 = v1736
	v1790 = v1724
	goto L389
L389:
	;
	v1797 = v1788 << (uint(int32(2)) % 32)
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1797)))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1797+v1723)))
	v1811 = (v1799-v1801)*v1789 + v1790
	goto L385
L390:
	;
	v1753 = int32(2)
	v1754 = v1745 << (uint(v1753) % 32)
	v1756 = v1754 - int32(4)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1756)))
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1723+v1756)))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1754+v1721)))
	v1764 = v1763 * v1746
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1754+l2)))
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1754+v1723)))
	v1773 = (v1758-v1760)*v1764 + ((v1767-v1769)*v1746 + v1747)
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1721+v1756)))
	v1776 = v1775 * v1764
	v1778 = v1745 - v1753
	v1780 = v1752 + v1753
	if v1780 != l1&int32(-2) {
		v1745 = v1778
		v1746 = v1776
		v1747 = v1773
		v1752 = v1780
		goto L390
	} else {
		goto L392
	}
L391:
	;
	if l1&int32(1) == int32(0) {
		v1811 = v1773
		goto L385
	} else {
		goto L393
	}
L392:
	;
	goto L391
L393:
	;
	v1788 = v1778
	v1789 = v1776
	v1790 = v1773
	goto L389
L394:
	;
	v1821 = base.I32_div_s(v1817, int32(8))
	v1822 = v1821 + v1817
	if v1817 < v1822 {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	v1838 = v166
	v1840 = v170
	v1841 = v1818
	goto L396
L396:
	;
	if base.B2i32(v1838 == int32(0))&v1705 != 0 {
		goto L406
	} else {
		goto L407
	}
L397:
	;
	v1824 = v1822
	goto L399
L398:
	;
	v1824 = v1817
	goto L399
L399:
	;
	v1827 = F_repalloc(m, v170, v1824<<(uint(int32(3))%32))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L6
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+48)) = v1827
	if v166 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+56)) = v1824
	v1838 = v1836
	v1840 = v1827
	v1841 = v1824
	goto L396
L402:
	;
	v1836 = int32(0)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v1833 = F_repalloc(m, v166, v1824)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L6
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+52)) = v1833
	v1836 = v1833
	goto L401
L406:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v1846 = F_MemoryContextAllocZero(m, v1845, v1841)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L6
	} else {
		goto L409
	}
L407:
	;
	v1849 = v1838
	goto L408
L408:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v101)+64)) = int64(0)
	if v1719 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+52)) = v1846
	v1849 = v1846
	goto L408
L410:
	;
	if int32(0) < v1702 {
		goto L416
	} else {
		goto L417
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+28)) = l1
	v1856 = l1 << (uint(int32(2)) % 32)
	v1857 = int32(0)
	v1858 = base.B2i32(v1856 == v1857)
	if v1858 == v1857 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v101)+32))
	base.MemoryCopy(m, v1861, v33+int32(112), v1856)
	goto L414
L413:
	;
	goto L414
L414:
	;
	if v1856 == v1857 {
		goto L410
	} else {
		goto L415
	}
L415:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	base.MemoryCopy(m, v1865, v33+int32(80), v1856)
	goto L410
L416:
	;
	v1873 = int32(3)
	v1874 = v1702 << (uint(v1873) % 32)
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	v1877 = v1875 << (uint(v1873) % 32)
	if v1877 != 0 {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	goto L418
L418:
	;
	if int32(0) < v1704 {
		goto L431
	} else {
		goto L432
	}
L419:
	;
	base.MemoryCopy(m, v1874+v1840, v1840, v1877)
	goto L421
L420:
	;
	goto L421
L421:
	;
	if v1874 != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	base.MemoryFill(m, v1840, int32(0), v1874)
	goto L424
L423:
	;
	goto L424
L424:
	;
	if v1849 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+60)) = v1892 + v1702
	goto L418
L426:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	if v1884 != 0 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	base.MemoryCopy(m, v1849+v1702, v1849, v1884)
	goto L429
L428:
	;
	goto L429
L429:
	;
	if v1702 == int32(0) {
		goto L425
	} else {
		goto L430
	}
L430:
	;
	base.MemoryFill(m, v1849, int32(1), v1702)
	goto L425
L431:
	;
	v1900 = v1704 & int32(3)
	v1901 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1704) {
		goto L435
	} else {
		goto L436
	}
L432:
	;
	goto L433
L433:
	;
	v2280 = int32(0)
	v2281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+46)))
	if v2281 == v2280 {
		goto L459
	} else {
		goto L460
	}
L434:
	;
	if v1849 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L435:
	;
	v1909 = v1901
	v1920 = int32(0)
	goto L438
L436:
	;
	v1976 = v1901
	goto L437
L437:
	;
	v2006 = v1976
	v2018 = v1901
	goto L442
L438:
	;
	v1938 = int32(3)
	v1939 = v1909 << (uint(v1938) % 32)
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	v1945 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1939+(v1840+v1940<<(uint(v1938)%32))))) = v1945
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v1947<<(uint(v1938)%32)+v1939)+8)) = v1945
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v1954<<(uint(v1938)%32)+v1939)+16)) = v1945
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v1961<<(uint(v1938)%32)+v1939)+24)) = v1945
	v1968 = int32(4)
	v1969 = v1909 + v1968
	v1971 = v1920 + v1968
	if v1971 != v1704&int32(2147483644) {
		v1909 = v1969
		v1920 = v1971
		goto L438
	} else {
		goto L440
	}
L439:
	;
	if v1900 == int32(0) {
		goto L434
	} else {
		goto L441
	}
L440:
	;
	goto L439
L441:
	;
	v1976 = v1969
	goto L437
L442:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	v2036 = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v2035<<(uint(v2036)%32)+v2006<<(uint(v2036)%32)))) = int64(0)
	v2044 = int32(1)
	v2047 = v2018 + v2044
	if v2047 != v1900 {
		v2006 = v2006 + v2044
		v2018 = v2047
		goto L442
	} else {
		goto L444
	}
L443:
	;
	goto L434
L444:
	;
	goto L443
L445:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+60)) = v2247 + v1704
	goto L433
L446:
	;
	v2082 = v1704 & int32(3)
	v2083 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1704) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2091 = v2083
	v2103 = int32(0)
	goto L450
L448:
	;
	v2148 = v2083
	goto L449
L449:
	;
	v2178 = v2148
	v2179 = v2083
	goto L454
L450:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	v2123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1849+v2120+v2091))) = uint8(v2123)
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1849+v2125+v2091)+1)) = uint8(v2123)
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1849+v2130+v2091)+2)) = uint8(v2123)
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v1849+v2135+v2091)+3)) = uint8(v2123)
	v2140 = int32(4)
	v2141 = v2091 + v2140
	v2143 = v2103 + v2140
	if v2143 != v1704&int32(2147483644) {
		v2091 = v2141
		v2103 = v2143
		goto L450
	} else {
		goto L452
	}
L451:
	;
	if v2082 == int32(0) {
		goto L445
	} else {
		goto L453
	}
L452:
	;
	goto L451
L453:
	;
	v2148 = v2141
	goto L449
L454:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	v2210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1849+v2207+v2178))) = uint8(v2210)
	v2215 = v2179 + v2210
	if v2215 != v2082 {
		v2178 = v2178 + v2210
		v2179 = v2215
		goto L454
	} else {
		goto L456
	}
L455:
	;
	goto L445
L456:
	;
	goto L455
L457:
	;
	if v2308 == int32(0) {
		goto L466
	} else {
		goto L467
	}
L458:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1811+v1849))) = uint8(v5)
	v2308 = v2305
	goto L457
L459:
	;
	if v1849 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L460:
	;
	v2298 = v2280
	goto L461
L461:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v1811<<(uint(int32(3))%32)))) = v162
	if v1849 == int32(0) {
		v2308 = v2298
		goto L457
	} else {
		goto L465
	}
L462:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v1840+v1811<<(uint(int32(3))%32))))
	v2298 = v2297
	goto L461
L463:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1811+v1849))))
	if v2287 != int32(1) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1840+v1811<<(uint(int32(3))%32)))) = v162
	v2305 = v2280
	goto L458
L465:
	;
	v2305 = v2298
	goto L458
L466:
	;
	v2330 = v101 + int32(12)
	goto L12
L467:
	;
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v101)+72))
	if base.Ui32(v2311) <= base.Ui32(v2308) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v101)+76))
	if base.Ui32(v2308) < base.Ui32(v2313) {
		goto L466
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	F_pfree(m, v2308)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L6
	} else {
		goto L472
	}
L471:
	;
	goto L470
L472:
	;
	goto L466
L473:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L6
	} else {
		goto L474
	}
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_set_element_10), v33+int32(16))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L6
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2624), int32(_a_F_array_set_element_6))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L6
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L6
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_set_element_10), v33+int32(32))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L6
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_array_set_element_5), int32(2640), int32(_a_F_array_set_element_6))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L6
	} else {
		goto L480
	}
L480:
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
	var v17 int32
	_ = v17
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
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
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
	return v148
L2:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v140)
	v148 = int32(0)
	goto L1
L3:
	;
	v17 = int32(0)
	v20 = v11
	goto L6
L4:
	;
	goto L5
L5:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v76 <= int32(0) {
		v148 = v8
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v17))))
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v17))))
	if v31 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v65 = v20
	goto L10
L10:
	;
	v67 = v17 + int32(1)
	if v67 < v65 {
		v17 = v67
		v20 = v65
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v58+v17<<(uint(int32(3))%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(12)+v17<<(uint(int32(2))%32)))) = uint32(v62)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v65 = v64
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
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(200), int32(_a_F_array_subscript_check_subscripts_2))
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
	v85 = v76
	goto L22
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v82))))
	if v91 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v148 = v8
	goto L1
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v82))))
	if v96 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v128 = v85
	goto L26
L26:
	;
	v130 = v82 + int32(1)
	if v130 < v128 {
		v82 = v130
		v85 = v128
		goto L22
	} else {
		goto L35
	}
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v99 != int32(1) {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v121+v82<<(uint(int32(3))%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v10+int32(36)+v82<<(uint(int32(2))%32)))) = uint32(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v128 = v127
	goto L26
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_array_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_array_subscript_check_subscripts_1), int32(219), int32(_a_F_array_subscript_check_subscripts_2))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
func F_array_to_text(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = F_pg_detoast_datum_packed(m, v13)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
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
					return int64(0)
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
							return int64(0)
						} else {
							v66 = F_array_to_text_internal(m, l0, v8, v49, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v66)
							}
						}
					} else {
						v66 = F_array_to_text_internal(m, l0, v8, v49, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v66)
						}
					}
				}
			}
		}
	}
}
func F_array_to_vector(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 float64
	_ = v239
	var v243 int32
	_ = v243
	var v250 float64
	_ = v250
	var v254 int32
	_ = v254
	var v261 float64
	_ = v261
	var v265 int32
	_ = v265
	var v272 float64
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v308 float64
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v455 float32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L82
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L78
	}
L4:
	;
	return int64(0)
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v24 != 0 {
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
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L74
	}
L9:
	;
	v25 = F_array_contains_nulls(m, v16)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_get_typlenbyvalalign(m, v27, v13+int32(30), v13+int32(29), v13+int32(28))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	if v25 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+30)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+29)))
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+28)))
	F_deconstruct_array(m, v16, v37, v38, v39, v13+int32(24), int32(0), v13+int32(20))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_CheckDim_3(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if base.B2i32(v23 != int32(-1))&base.B2i32(v52 != v23) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v57 = F_mul_size(m, int32(4), v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v59 = F_add_size(m, int32(8), v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v61 = F_palloc0(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)) = uint16(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v59 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	switch v67 - int32(700) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L25
	}
L21:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	F_pfree(m, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L66
	}
L22:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v316 <= int32(0) {
		goto L21
	} else {
		goto L55
	}
L23:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v209 <= int32(0) {
		goto L21
	} else {
		goto L44
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v107 <= int32(0) {
		goto L21
	} else {
		goto L33
	}
L25:
	;
	if v67 == int32(23) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v67 != int32(1700) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v74 <= int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v80 = int32(0)
	goto L29
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v95+v80<<(uint(int32(3))%32))))
	v100 = F_DirectFunctionCall1Coll(m, int32(1459), int32(0), v99)
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
	*(*uint32)(unsafe.Add(mBase, uint32(v61+int32(8)+v80<<(uint(int32(2))%32)))) = uint32(v100)
	v104 = v80 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v104 < v105 {
		v80 = v104
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
	v113 = v61 + int32(8)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v115 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v107) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v121 = v115
	v122 = int32(0)
	goto L37
L35:
	;
	v176 = v115
	goto L36
L36:
	;
	v186 = v176
	v195 = v2
	goto L41
L37:
	;
	v131 = int32(2)
	v134 = int32(3)
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v114+v121<<(uint(v134)%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v113+v121<<(uint(v131)%32)))) = uint32(v137)
	v140 = v121 | int32(1)
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v114+v140<<(uint(v134)%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v113+v140<<(uint(v131)%32)))) = uint32(v147)
	v150 = v121 | v131
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v114+v150<<(uint(v134)%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v113+v150<<(uint(v131)%32)))) = uint32(v157)
	v160 = v121 | v134
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v114+v160<<(uint(v134)%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v113+v160<<(uint(v131)%32)))) = uint32(v167)
	v169 = int32(4)
	v170 = v121 + v169
	v172 = v122 + v169
	if v172 != v107&int32(2147483644) {
		v121 = v170
		v122 = v172
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
	v176 = v170
	goto L36
L41:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v114+v186<<(uint(int32(3))%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v113+v186<<(uint(int32(2))%32)))) = uint32(v202)
	v204 = int32(1)
	v207 = v195 + v204
	if v207 != v111 {
		v186 = v186 + v204
		v195 = v207
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
	v213 = v209 & int32(3)
	v215 = v61 + int32(8)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v217 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v209) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v223 = v217
	v224 = int32(0)
	goto L48
L46:
	;
	v282 = v217
	goto L47
L47:
	;
	v292 = v282
	v301 = v2
	goto L52
L48:
	;
	v233 = int32(2)
	v236 = int32(3)
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v216+v223<<(uint(v236)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v215+v223<<(uint(v233)%32)))) = base.F32_demote_f64(v239)
	v243 = v223 | int32(1)
	v250 = *(*float64)(unsafe.Add(mBase, uint32(v216+v243<<(uint(v236)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v215+v243<<(uint(v233)%32)))) = base.F32_demote_f64(v250)
	v254 = v223 | v233
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v216+v254<<(uint(v236)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v215+v254<<(uint(v233)%32)))) = base.F32_demote_f64(v261)
	v265 = v223 | v236
	v272 = *(*float64)(unsafe.Add(mBase, uint32(v216+v265<<(uint(v236)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v215+v265<<(uint(v233)%32)))) = base.F32_demote_f64(v272)
	v275 = int32(4)
	v276 = v223 + v275
	v278 = v224 + v275
	if v278 != v209&int32(2147483644) {
		v223 = v276
		v224 = v278
		goto L48
	} else {
		goto L50
	}
L49:
	;
	if v213 == int32(0) {
		goto L21
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v282 = v276
	goto L47
L52:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v216+v292<<(uint(int32(3))%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v215+v292<<(uint(int32(2))%32)))) = base.F32_demote_f64(v308)
	v311 = int32(1)
	v314 = v301 + v311
	if v314 != v213 {
		v292 = v292 + v311
		v301 = v314
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
	v320 = v316 & int32(3)
	v322 = v61 + int32(8)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v324 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v316) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v330 = v324
	v331 = int32(0)
	goto L59
L57:
	;
	v389 = v324
	goto L58
L58:
	;
	v399 = v389
	v408 = v2
	goto L63
L59:
	;
	v340 = int32(2)
	v343 = int32(3)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v323+v330<<(uint(v343)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v322+v330<<(uint(v340)%32)))) = base.F32_convert_i32_s(v346)
	v350 = v330 | int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v323+v350<<(uint(v343)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v322+v350<<(uint(v340)%32)))) = base.F32_convert_i32_s(v357)
	v361 = v330 | v340
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v323+v361<<(uint(v343)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v322+v361<<(uint(v340)%32)))) = base.F32_convert_i32_s(v368)
	v372 = v330 | v343
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v323+v372<<(uint(v343)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v322+v372<<(uint(v340)%32)))) = base.F32_convert_i32_s(v379)
	v382 = int32(4)
	v383 = v330 + v382
	v385 = v331 + v382
	if v385 != v316&int32(2147483644) {
		v330 = v383
		v331 = v385
		goto L59
	} else {
		goto L61
	}
L60:
	;
	if v320 == int32(0) {
		goto L21
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v389 = v383
	goto L58
L63:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v323+v399<<(uint(int32(3))%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v322+v399<<(uint(int32(2))%32)))) = base.F32_convert_i32_s(v415)
	v418 = int32(1)
	v421 = v408 + v418
	if v421 != v320 {
		v399 = v399 + v418
		v408 = v421
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
	v436 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+4)))
	if int32(0) < v436 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v442 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	m.G0 = v13 + int32(32)
	return base.I64_extend_i32_u(v61)
L70:
	;
	v455 = *(*float32)(unsafe.Add(mBase, uint32(v61+int32(8)+v442<<(uint(int32(2))%32))))
	F_CheckElement_3(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	v459 = v442 + int32(1)
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+4)))
	if v459 < v460 {
		v442 = v459
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
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_0), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(459), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
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
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_3), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(464), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
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
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	F_errmsg(m, int32(_a_F_array_to_vector_4), v13)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(88), int32(_a_F_array_to_vector_5))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
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
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_array_to_vector_6), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_array_to_vector_1), int32(498), int32(_a_F_array_to_vector_2))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
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
func F_array_typanalyze(m *base.Module, l0 int32) int64 {
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
	var v23 int64
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
	var v67 int64
	_ = v67
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_std_typanalyze(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v67 = int64(0)
			m.G0 = v9 + int32(16)
			return v67
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = F_get_base_element_type(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v76
						F_errmsg_internal(m, int32(_a_F_array_typanalyze_0), v9)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_array_typanalyze_1), int32(118), int32(_a_F_array_typanalyze_2))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = int64(1)
					v25 = F_lookup_type_cache(m, v19, int32(193))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
						if v27 == int32(0) {
							v67 = v23
							m.G0 = v9 + int32(16)
							return v67
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+108))
							if v30 == int32(0) {
								v67 = v23
								m.G0 = v9 + int32(16)
								return v67
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+136))
								if v33 == int32(0) {
									v67 = v23
									m.G0 = v9 + int32(16)
									return v67
								} else {
									v37 = F_palloc(m, int32(36))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int64(0)
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
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1372)
										v67 = v23
										m.G0 = v9 + int32(16)
										return v67
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
func F_array_unnest(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		v19 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = int32(_a_F_array_unnest_0)
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v26
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			v29 = F_DatumGetAnyArrayP(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v32 = F_palloc(m, int32(36))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v34 == int32(-1) {
						v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+44)))
						*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)) = uint16(v37)
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+46)))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v39)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+47)))
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v41)
						v55 = v37
						v56 = v39
						v57 = v41
						F_array_iter_setup(m, v32, v29, base.I32_extend16_s(v55), v56&int32(1), base.I32_extend8_s(v57))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = int32(0)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
							if v68 == int32(-1) {
								v71 = int32(28)
							} else {
								v71 = int32(4)
							}
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v29+v71)))
							if v68 == int32(-1) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
								v79 = v76
							} else {
								v79 = v29 + int32(16)
							}
							v80 = F_ArrayGetNItemsSafe(m, v73, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v80
								*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v32
								*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v24
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
								if v96 < v97 {
									*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v96 + int32(1)
									v104 = F_array_iter_next(m, v95, l0+int32(16), v96)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										v106 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
										*(*int64)(unsafe.Add(mBase, uint32(v94))) = v106 + int64(1)
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = int32(1)
										v120 = v104
										m.G0 = v13 + int32(16)
										return v120
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = int32(2)
										v118 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
										v120 = int64(0)
										m.G0 = v13 + int32(16)
										return v120
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						F_get_typlenbyvalalign(m, v43, v13+int32(14), v13+int32(13), v13+int32(12))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
							v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+12)))
							v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
							v55 = v54
							v56 = v52
							v57 = v53
							F_array_iter_setup(m, v32, v29, base.I32_extend16_s(v55), v56&int32(1), base.I32_extend8_s(v57))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = int32(0)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
								if v68 == int32(-1) {
									v71 = int32(28)
								} else {
									v71 = int32(4)
								}
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v29+v71)))
								if v68 == int32(-1) {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
									v79 = v76
								} else {
									v79 = v29 + int32(16)
								}
								v80 = F_ArrayGetNItemsSafe(m, v73, v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v80
									*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v32
									*(*int32)(unsafe.Add(mBase, _c_F_array_unnest[0])) = v24
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
									if v96 < v97 {
										*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v96 + int32(1)
										v104 = F_array_iter_next(m, v95, l0+int32(16), v96)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											v106 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
											*(*int64)(unsafe.Add(mBase, uint32(v94))) = v106 + int64(1)
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = int32(1)
											v120 = v104
											m.G0 = v13 + int32(16)
											return v120
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = int32(2)
											v118 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
											v120 = int64(0)
											m.G0 = v13 + int32(16)
											return v120
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
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
		if v96 < v97 {
			*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v96 + int32(1)
			v104 = F_array_iter_next(m, v95, l0+int32(16), v96)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int64(0)
			} else {
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
				*(*int64)(unsafe.Add(mBase, uint32(v94))) = v106 + int64(1)
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = int32(1)
				v120 = v104
				m.G0 = v13 + int32(16)
				return v120
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return int64(0)
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = int32(2)
				v118 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v118)
				v120 = int64(0)
				m.G0 = v13 + int32(16)
				return v120
			}
		}
	}
}
func F_array_unnest_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.I32_wrap_i64(v5)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != int32(468) {
		v31 = v3
		return v31
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		if v10 == int32(0) {
			v31 = v3
			return v31
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(15) {
				v31 = v3
				return v31
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v22 = F_estimate_expression_value(m, v18, v21)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					v27 = F_estimate_array_length(m, v26, v22)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v27
						v31 = v5 & int64(4294967295)
						return v31
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v17 = F_MemoryContextAlloc(m, v15, int32(48))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v17
			v25 = v17
			v27 = v9 + int32(12)
			v28 = int32(0)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v29 == v28 {
				v46 = int32(0)
				if v27 == v46 {
					v54 = v46
				} else {
					v49 = v46
					v50 = v28
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
					v54 = v50
				}
				v57 = v54
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				switch v32 - int32(435) {
				case 0:
					if v27 == int32(0) {
						v57 = int32(1)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+168))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
						v49 = v39
						v50 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
						v54 = v50
						v57 = v54
					}
				case 1:
					if v27 == int32(0) {
						v57 = int32(2)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+376))
						v49 = v44
						v50 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
						v54 = v50
						v57 = v54
					}
				default:
					v46 = int32(0)
					if v27 == v46 {
						v54 = v46
					} else {
						v49 = v46
						v50 = v28
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
						v54 = v50
					}
					v57 = v54
				}
			}
			if v57 == int32(0) {
				v61 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v61
			} else {
			}
			v65 = l0 + l1<<(uint(int32(4))%32)
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+32)))
			if v66 == int32(0) {
				v69 = int32(_a_F_fetch_array_arg_replace_nulls_0)
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v72
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
				v75 = base.I32_wrap_i64(v74)
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
				if v76 != int32(1) {
					v95 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v96 = F_expand_array(m, v74, v95, v25)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v96))+2))
						v101 = v99
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
						v129 = v101
						m.G0 = v9 + int32(16)
						return v129
					}
				} else {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
					if v79 != int32(3) {
						v95 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
						v96 = F_expand_array(m, v74, v95, v25)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v96))+2))
							v101 = v99
							*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
							v129 = v101
							m.G0 = v9 + int32(16)
							return v129
						}
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v74))+2))
						if v25 == int32(0) {
							v101 = v83
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v86
							v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+44)))
							*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v88)
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v90)
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+47)))
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)) = uint8(v92)
							v101 = v83
						}
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
						v129 = v101
						m.G0 = v9 + int32(16)
						return v129
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v105 = F_get_fn_expr_argtype(m, v104, l1)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					if v105 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(119), int32(_a_F_fetch_array_arg_replace_nulls_3))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
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
						v109 = F_get_element_type(m, v105)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							if v109 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(124), int32(_a_F_fetch_array_arg_replace_nulls_3))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
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
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								v115 = F_palloc0(m, int32(16))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v115))) = int64(64)
									v123 = F_expand_array(m, base.I64_extend_i32_u(v115), v113, v25)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v115)
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v128 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v123))+2))
											v129 = v128
											m.G0 = v9 + int32(16)
											return v129
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
		v25 = v12
		v27 = v9 + int32(12)
		v28 = int32(0)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v29 == v28 {
			v46 = int32(0)
			if v27 == v46 {
				v54 = v46
			} else {
				v49 = v46
				v50 = v28
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
				v54 = v50
			}
			v57 = v54
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			switch v32 - int32(435) {
			case 0:
				if v27 == int32(0) {
					v57 = int32(1)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+168))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					v49 = v39
					v50 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
					v54 = v50
					v57 = v54
				}
			case 1:
				if v27 == int32(0) {
					v57 = int32(2)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+376))
					v49 = v44
					v50 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
					v54 = v50
					v57 = v54
				}
			default:
				v46 = int32(0)
				if v27 == v46 {
					v54 = v46
				} else {
					v49 = v46
					v50 = v28
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v49
					v54 = v50
				}
				v57 = v54
			}
		}
		if v57 == int32(0) {
			v61 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v61
		} else {
		}
		v65 = l0 + l1<<(uint(int32(4))%32)
		v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+32)))
		if v66 == int32(0) {
			v69 = int32(_a_F_fetch_array_arg_replace_nulls_0)
			v70 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v72
			v74 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
			v75 = base.I32_wrap_i64(v74)
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
			if v76 != int32(1) {
				v95 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
				v96 = F_expand_array(m, v74, v95, v25)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v96))+2))
					v101 = v99
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
					v129 = v101
					m.G0 = v9 + int32(16)
					return v129
				}
			} else {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
				if v79 != int32(3) {
					v95 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0]))
					v96 = F_expand_array(m, v74, v95, v25)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v96))+2))
						v101 = v99
						*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
						v129 = v101
						m.G0 = v9 + int32(16)
						return v129
					}
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v74))+2))
					if v25 == int32(0) {
						v101 = v83
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v86
						v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+44)))
						*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v88)
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+46)))
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v90)
						v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+47)))
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)) = uint8(v92)
						v101 = v83
					}
					*(*int32)(unsafe.Add(mBase, _c_F_fetch_array_arg_replace_nulls[0])) = v70
					v129 = v101
					m.G0 = v9 + int32(16)
					return v129
				}
			}
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v105 = F_get_fn_expr_argtype(m, v104, l1)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				if v105 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_1), int32(0))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(119), int32(_a_F_fetch_array_arg_replace_nulls_3))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
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
					v109 = F_get_element_type(m, v105)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						if v109 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_fetch_array_arg_replace_nulls_4), int32(0))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_fetch_array_arg_replace_nulls_2), int32(124), int32(_a_F_fetch_array_arg_replace_nulls_3))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
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
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							v115 = F_palloc0(m, int32(16))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v115))) = int64(64)
								v123 = F_expand_array(m, base.I64_extend_i32_u(v115), v113, v25)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v115)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v123))+2))
										v129 = v128
										m.G0 = v9 + int32(16)
										return v129
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
				v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(3))%32))
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
			v23 = F_MemoryContextAlloc(m, v14, l3<<(uint(int32(3))%32))
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
