package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_decompress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v6 == v11 {
				return v4
			} else {
				v16 = F_palloc(m, int32(16))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v6
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v21
					v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
					v24 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)) = uint8(v24)
					*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)) = uint16(v23)
					return v16
				}
			}
		}
	}
}
func F_g_cube_distance(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v69 float64
	_ = v69
	var v74 float64
	_ = v74
	var v78 float64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(15) {
			if v20 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_g_cube_distance_0), int32(0))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_g_cube_distance_1), int32(1423), int32(_a_F_g_cube_distance_2))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
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
				v27 = v20 >> (uint(int32(31)) % 32)
				v29 = v20 ^ v27 - v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if base.Ui32(v30<<(uint(int32(1))%32)) < base.Ui32(v29) {
					v74 = float64(0)
				} else {
					v35 = v29 - int32(1)
					v40 = v16 + v35<<(uint(int32(2))%32)&int32(-8)
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v40)+8))
					if v30 < int32(0) {
						v74 = v41
					} else {
						v51 = *(*float64)(unsafe.Add(mBase, uint32(v40+int32(8)+v30&int32(2147483647)<<(uint(int32(3))%32))))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+16)))
						v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v53)+12)))
						if v55&int32(1) != 0 {
							if base.F64_gt(v41, v51) != 0 {
								v59 = v41
							} else {
								v59 = v51
							}
							if v35&int32(1) != 0 {
								v74 = v59
							} else {
								if base.F64_lt(v41, v51) != 0 {
									v63 = v41
								} else {
									v63 = v51
								}
								v74 = v63
							}
						} else {
							if base.F64_lt(v41, v51) != 0 {
								v65 = v41
							} else {
								v65 = v51
							}
							if int32(0) <= v20 {
								v74 = v65
							} else {
								if base.F64_gt(v41, v51) != 0 {
									v69 = v41
								} else {
									v69 = v51
								}
								v74 = v69
							}
						}
					}
				}
				if v20 < int32(0) {
					v78 = base.F64_neg(v74)
				} else {
					v78 = v74
				}
				v117 = v78
				v118 = F_Float8GetDatum(m, v117)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(16)
					return v118
				}
			}
		} else {
			v79 = F_pg_detoast_datum(m, v20)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				switch v13 - int32(16) {
				case 0:
					v108 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_distance_3), int32(0), v16, v79)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*float64)(unsafe.Add(mBase, uint32(v108)))
						v117 = v110
						v118 = F_Float8GetDatum(m, v117)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return v118
						}
					}
				case 1:
					v85 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_distance_4), int32(0), v16, v79)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v87 = *(*float64)(unsafe.Add(mBase, uint32(v85)))
						v117 = v87
						v118 = F_Float8GetDatum(m, v117)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return v118
						}
					}
				case 2:
					v90 = F_DirectFunctionCall2Coll(m, int32(_a_F_g_cube_distance_5), int32(0), v16, v79)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
						v117 = v92
						v118 = F_Float8GetDatum(m, v117)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return v118
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
						F_errmsg_internal(m, int32(_a_F_g_cube_distance_6), v11)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_g_cube_distance_1), int32(1498), int32(_a_F_g_cube_distance_2))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
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
		}
	}
}
func F_g_cube_leaf_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v77 int32
	_ = v77
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v148 int32
	_ = v148
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v159 float64
	_ = v159
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 float64
	_ = v263
	var v265 int32
	_ = v265
	var v269 float64
	_ = v269
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 float64
	_ = v280
	var v283 float64
	_ = v283
	var v288 float64
	_ = v288
	var v290 float64
	_ = v290
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v332 int32
	_ = v332
	var v336 float64
	_ = v336
	var v339 float64
	_ = v339
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v343 int32
	_ = v343
	var v347 float64
	_ = v347
	var v350 float64
	_ = v350
	var v355 float64
	_ = v355
	var v357 float64
	_ = v357
	var v362 float64
	_ = v362
	var v364 float64
	_ = v364
	var v368 int32
	_ = v368
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v413 int32
	_ = v413
	var v414 float64
	_ = v414
	var v421 float64
	_ = v421
	var v423 int32
	_ = v423
	var v427 float64
	_ = v427
	var v433 float64
	_ = v433
	var v439 float64
	_ = v439
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v468 float64
	_ = v468
	var v475 float64
	_ = v475
	var v477 int32
	_ = v477
	var v481 float64
	_ = v481
	var v487 float64
	_ = v487
	var v493 float64
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v528 int32
	_ = v528
	var v529 float64
	_ = v529
	var v536 float64
	_ = v536
	var v538 int32
	_ = v538
	var v542 float64
	_ = v542
	var v554 float64
	_ = v554
	var v558 float64
	_ = v558
	var v562 int32
	_ = v562
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v587 float64
	_ = v587
	var v594 float64
	_ = v594
	var v596 int32
	_ = v596
	var v600 float64
	_ = v600
	var v612 float64
	_ = v612
	var v616 float64
	_ = v616
	var v621 int32
	_ = v621
	var v653 int32
	_ = v653
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v745 int32
	_ = v745
	var v746 float64
	_ = v746
	var v754 float64
	_ = v754
	var v758 int32
	_ = v758
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 float64
	_ = v802
	var v804 int32
	_ = v804
	var v808 float64
	_ = v808
	var v811 float64
	_ = v811
	var v812 int32
	_ = v812
	var v813 float64
	_ = v813
	var v815 int32
	_ = v815
	var v819 float64
	_ = v819
	var v822 float64
	_ = v822
	var v827 float64
	_ = v827
	var v829 float64
	_ = v829
	var v834 float64
	_ = v834
	var v836 float64
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v853 int32
	_ = v853
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v917 int32
	_ = v917
	var v918 float64
	_ = v918
	var v926 float64
	_ = v926
	var v930 int32
	_ = v930
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 float64
	_ = v974
	var v976 int32
	_ = v976
	var v980 float64
	_ = v980
	var v983 float64
	_ = v983
	var v984 int32
	_ = v984
	var v985 float64
	_ = v985
	var v987 int32
	_ = v987
	var v991 float64
	_ = v991
	var v994 float64
	_ = v994
	var v999 float64
	_ = v999
	var v1001 float64
	_ = v1001
	var v1006 float64
	_ = v1006
	var v1008 float64
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1025 int32
	_ = v1025
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	switch l2 - int32(3) {
	case 0:
		goto L5
	default:
		v1044 = int32(0)
		goto L1
	case 3:
		goto L4
	case 4, 10:
		goto L3
	case 5, 11:
		goto L2
	}
L1:
	;
	return v1044
L2:
	;
	v873 = int32(0)
	if base.B2i32(l1 == v873)|base.B2i32(l0 == v873) != 0 {
		v1025 = v873
		goto L202
	} else {
		goto L203
	}
L3:
	;
	v701 = int32(0)
	if base.B2i32(l0 == v701)|base.B2i32(l1 == v701) != 0 {
		v853 = v701
		goto L165
	} else {
		goto L166
	}
L4:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = int32(2147483647)
	v232 = v230 & v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v235 = v233 & v231
	if base.Ui32(v232) < base.Ui32(v235) {
		goto L58
	} else {
		goto L59
	}
L5:
	;
	v10 = int32(0)
	if base.B2i32(l0 == v10)|base.B2i32(l1 == v10) != 0 {
		v190 = v10
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v212
L7:
	;
	v212 = v190
	goto L6
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = int32(2147483647)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = base.B2i32(base.Ui32(v25&v26) < base.Ui32(v28&v26))
	if base.Ui32(v25&v26) < base.Ui32(v28&v26) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = l1
	goto L11
L10:
	;
	v32 = l0
	goto L11
L11:
	;
	if base.Ui32(v25&v26) < base.Ui32(v28&v26) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v120 = v104 & int32(2147483647)
	if base.Ui32(v120) <= base.Ui32(v36) {
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v33 = l0
	goto L15
L14:
	;
	v33 = l1
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v36 = v34 & int32(2147483647)
	if v36 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v104 = v39
	goto L12
L17:
	;
	goto L18
L18:
	;
	v40 = int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v47 = int32(0)
	goto L19
L19:
	;
	v62 = v47 << (uint(int32(3)) % 32)
	v63 = v32 + v40 + v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	v66 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		v73 = v64
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v104 = v44
	goto L12
L21:
	;
	v74 = v62 + (v33 + v40)
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v74)))
	v77 = base.B2i32(v34 < int32(0))
	if v34 < int32(0) {
		v84 = v75
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v63+v44<<(uint(int32(3))%32))))
	if base.F64_lt(v64, v70) != 0 {
		v73 = v64
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v73 = v70
	goto L21
L24:
	;
	if base.F64_gt(v73, v84) != 0 {
		v190 = v10
		goto L7
	} else {
		goto L27
	}
L25:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v74+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v75, v81) != 0 {
		v84 = v75
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v84 = v81
	goto L24
L27:
	;
	if v44 < int32(0) {
		v91 = v64
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v34 < int32(0) {
		v98 = v75
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v63+v44<<(uint(int32(3))%32))))
	if base.F64_gt(v64, v89) != 0 {
		v91 = v64
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v91 = v89
	goto L28
L31:
	;
	if base.F64_gt(v98, v91) != 0 {
		v190 = v10
		goto L7
	} else {
		goto L34
	}
L32:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v74+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v75, v96) != 0 {
		v98 = v75
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v98 = v96
	goto L31
L34:
	;
	v102 = v47 + int32(1)
	if v102 != v36 {
		v47 = v102
		goto L19
	} else {
		goto L35
	}
L35:
	;
	goto L20
L36:
	;
	v212 = int32(1)
	goto L6
L37:
	;
	goto L38
L38:
	;
	v132 = v36
	goto L39
L39:
	;
	v144 = v32 + int32(8) + v132<<(uint(int32(3))%32)
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v144)))
	if base.B2i32(v104 < int32(0)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v190 = int32(0)
	goto L7
L41:
	;
	goto L40
L42:
	;
	if base.F64_lt(v171, float64(0)) != 0 {
		goto L41
	} else {
		goto L52
	}
L43:
	;
	v148 = int32(0)
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v144+v120<<(uint(int32(3))%32))))
	if base.F64_lt(v145, v153) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if base.F64_gt(v145, float64(0)) != 0 {
		goto L41
	} else {
		goto L51
	}
L46:
	;
	v155 = v148
	goto L48
L47:
	;
	v155 = v104
	goto L48
L48:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v144+v155<<(uint(int32(3))%32))))
	if base.F64_gt(v159, float64(0)) != 0 {
		v190 = v148
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v144+v104<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v165) == int32(0) {
		v171 = v165
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v171 = v145
	goto L42
L51:
	;
	v171 = v145
	goto L42
L52:
	;
	v175 = int32(1)
	v177 = v132 + v175
	if v120 != v177 {
		v132 = v177
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v190 = v175
	goto L7
L54:
	;
	return base.B2i32(v697 == int32(0))
L55:
	;
	v697 = int32(-1)
	goto L54
L56:
	;
	v697 = v653
	goto L54
L57:
	;
	v653 = int32(1)
	goto L56
L58:
	;
	v237 = v232
	goto L60
L59:
	;
	v237 = v235
	goto L60
L60:
	;
	if v237 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v238 = int32(8)
	v255 = int32(0)
	goto L64
L62:
	;
	goto L63
L63:
	;
	if base.Ui32(v235) < base.Ui32(v232) {
		goto L98
	} else {
		goto L99
	}
L64:
	;
	v261 = v255 << (uint(int32(3)) % 32)
	v262 = l0 + v238 + v261
	v263 = *(*float64)(unsafe.Add(mBase, uint32(v262)))
	v265 = base.B2i32(v230 < int32(0))
	if v230 < int32(0) {
		v272 = v263
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v304 = int32(8)
	v322 = int32(0)
	goto L81
L66:
	;
	v273 = v261 + (l1 + v238)
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v273)))
	v276 = base.B2i32(v233 < int32(0))
	if v233 < int32(0) {
		v283 = v274
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v269 = *(*float64)(unsafe.Add(mBase, uint32(v262+v230<<(uint(int32(3))%32))))
	if base.F64_lt(v263, v269) != 0 {
		v272 = v263
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v272 = v269
	goto L66
L69:
	;
	if base.F64_gt(v272, v283) != 0 {
		goto L57
	} else {
		goto L72
	}
L70:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v273+v233<<(uint(int32(3))%32))))
	if base.F64_lt(v274, v280) != 0 {
		v283 = v274
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v283 = v280
	goto L69
L72:
	;
	if v230 < int32(0) {
		v290 = v263
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v233 < int32(0) {
		v297 = v274
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v262+v230<<(uint(int32(3))%32))))
	if base.F64_lt(v263, v288) != 0 {
		v290 = v263
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v290 = v288
	goto L73
L76:
	;
	v299 = int32(-1)
	if base.F64_lt(v290, v297) != 0 {
		v653 = v299
		goto L56
	} else {
		goto L79
	}
L77:
	;
	v295 = *(*float64)(unsafe.Add(mBase, uint32(v273+v233<<(uint(int32(3))%32))))
	if base.F64_lt(v274, v295) != 0 {
		v297 = v274
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v297 = v295
	goto L76
L79:
	;
	v302 = v255 + int32(1)
	if v302 != v237 {
		v255 = v302
		goto L64
	} else {
		goto L80
	}
L80:
	;
	goto L65
L81:
	;
	v328 = v322 << (uint(int32(3)) % 32)
	v329 = l0 + v304 + v328
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v329)))
	v332 = base.B2i32(v230 < int32(0))
	if v230 < int32(0) {
		v339 = v330
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L63
L83:
	;
	v340 = v328 + (l1 + v304)
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v340)))
	v343 = base.B2i32(v233 < int32(0))
	if v233 < int32(0) {
		v350 = v341
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v329+v230<<(uint(int32(3))%32))))
	if base.F64_gt(v330, v336) != 0 {
		v339 = v330
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v339 = v336
	goto L83
L86:
	;
	if base.F64_gt(v339, v350) != 0 {
		goto L57
	} else {
		goto L89
	}
L87:
	;
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v340+v233<<(uint(int32(3))%32))))
	if base.F64_gt(v341, v347) != 0 {
		v350 = v341
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v350 = v347
	goto L86
L89:
	;
	if v230 < int32(0) {
		v357 = v330
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v233 < int32(0) {
		v364 = v341
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v355 = *(*float64)(unsafe.Add(mBase, uint32(v329+v230<<(uint(int32(3))%32))))
	if base.F64_gt(v330, v355) != 0 {
		v357 = v330
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v357 = v355
	goto L90
L93:
	;
	if base.F64_lt(v357, v364) != 0 {
		v653 = v299
		goto L56
	} else {
		goto L96
	}
L94:
	;
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v340+v233<<(uint(int32(3))%32))))
	if base.F64_gt(v341, v362) != 0 {
		v364 = v341
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v364 = v362
	goto L93
L96:
	;
	v368 = v322 + int32(1)
	if v368 != v237 {
		v322 = v368
		goto L81
	} else {
		goto L97
	}
L97:
	;
	goto L82
L98:
	;
	v390 = l0 + int32(8)
	v393 = v237
	goto L101
L99:
	;
	goto L100
L100:
	;
	if base.Ui32(v235) <= base.Ui32(v232) {
		goto L131
	} else {
		goto L132
	}
L101:
	;
	v413 = v390 + v393<<(uint(int32(3))%32)
	v414 = *(*float64)(unsafe.Add(mBase, uint32(v413)))
	if base.B2i32(v230 < int32(0)) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v456 = v237
	goto L115
L103:
	;
	if base.F64_lt(v439, float64(0)) != 0 {
		goto L55
	} else {
		goto L113
	}
L104:
	;
	v421 = *(*float64)(unsafe.Add(mBase, uint32(v413+v232<<(uint(int32(3))%32))))
	if base.F64_lt(v414, v421) != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if base.F64_gt(v414, float64(0)) != 0 {
		goto L57
	} else {
		goto L112
	}
L107:
	;
	v423 = int32(0)
	goto L109
L108:
	;
	v423 = v230
	goto L109
L109:
	;
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v413+v423<<(uint(int32(3))%32))))
	if base.F64_gt(v427, float64(0)) != 0 {
		goto L57
	} else {
		goto L110
	}
L110:
	;
	v433 = *(*float64)(unsafe.Add(mBase, uint32(v413+v230<<(uint(int32(3))%32))))
	if base.F64_lt(v414, v433) == int32(0) {
		v439 = v433
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v439 = v414
	goto L103
L112:
	;
	v439 = v414
	goto L103
L113:
	;
	v443 = v393 + int32(1)
	if v443 != v232 {
		v393 = v443
		goto L101
	} else {
		goto L114
	}
L114:
	;
	goto L102
L115:
	;
	v467 = v390 + v456<<(uint(int32(3))%32)
	v468 = *(*float64)(unsafe.Add(mBase, uint32(v467)))
	if base.B2i32(v230 < int32(0)) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L55
L117:
	;
	if base.F64_lt(v493, float64(0)) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L118:
	;
	v475 = *(*float64)(unsafe.Add(mBase, uint32(v467+v232<<(uint(int32(3))%32))))
	if base.F64_gt(v468, v475) != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	if base.F64_gt(v468, float64(0)) != 0 {
		goto L57
	} else {
		goto L126
	}
L121:
	;
	v477 = int32(0)
	goto L123
L122:
	;
	v477 = v230
	goto L123
L123:
	;
	v481 = *(*float64)(unsafe.Add(mBase, uint32(v467+v477<<(uint(int32(3))%32))))
	if base.F64_gt(v481, float64(0)) != 0 {
		goto L57
	} else {
		goto L124
	}
L124:
	;
	v487 = *(*float64)(unsafe.Add(mBase, uint32(v467+v230<<(uint(int32(3))%32))))
	if base.F64_gt(v468, v487) == int32(0) {
		v493 = v487
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v493 = v468
	goto L117
L126:
	;
	v493 = v468
	goto L117
L127:
	;
	v498 = int32(1)
	v500 = v456 + v498
	if v500 == v232 {
		v653 = v498
		goto L56
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	goto L116
L130:
	;
	v456 = v500
	goto L115
L131:
	;
	v697 = int32(0)
	goto L54
L132:
	;
	goto L133
L133:
	;
	v505 = l1 + int32(8)
	v515 = v232
	goto L134
L134:
	;
	v528 = v505 + v515<<(uint(int32(3))%32)
	v529 = *(*float64)(unsafe.Add(mBase, uint32(v528)))
	if base.B2i32(v233 < int32(0)) == int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v575 = v237
	goto L149
L136:
	;
	if base.F64_lt(v558, float64(0)) != 0 {
		goto L57
	} else {
		goto L147
	}
L137:
	;
	v554 = *(*float64)(unsafe.Add(mBase, uint32(v528+v233<<(uint(int32(3))%32))))
	if base.F64_lt(v529, v554) == int32(0) {
		v558 = v554
		goto L136
	} else {
		goto L146
	}
L138:
	;
	v536 = *(*float64)(unsafe.Add(mBase, uint32(v528+v235<<(uint(int32(3))%32))))
	if base.F64_lt(v529, v536) != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	if base.F64_gt(v529, float64(0)) == int32(0) {
		v558 = v529
		goto L136
	} else {
		goto L145
	}
L141:
	;
	v538 = int32(0)
	goto L143
L142:
	;
	v538 = v233
	goto L143
L143:
	;
	v542 = *(*float64)(unsafe.Add(mBase, uint32(v528+v538<<(uint(int32(3))%32))))
	if base.F64_gt(v542, float64(0)) == int32(0) {
		goto L137
	} else {
		goto L144
	}
L144:
	;
	goto L55
L145:
	;
	goto L55
L146:
	;
	v558 = v529
	goto L136
L147:
	;
	v562 = v515 + int32(1)
	if v562 != v235 {
		v515 = v562
		goto L134
	} else {
		goto L148
	}
L148:
	;
	goto L135
L149:
	;
	v586 = v505 + v575<<(uint(int32(3))%32)
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v586)))
	if base.B2i32(v233 < int32(0)) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v653 = int32(-1)
	goto L56
L151:
	;
	if base.F64_lt(v616, float64(0)) != 0 {
		goto L57
	} else {
		goto L162
	}
L152:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(v586+v233<<(uint(int32(3))%32))))
	if base.F64_gt(v587, v612) == int32(0) {
		v616 = v612
		goto L151
	} else {
		goto L161
	}
L153:
	;
	v594 = *(*float64)(unsafe.Add(mBase, uint32(v586+v235<<(uint(int32(3))%32))))
	if base.F64_gt(v587, v594) != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	if base.F64_gt(v587, float64(0)) == int32(0) {
		v616 = v587
		goto L151
	} else {
		goto L160
	}
L156:
	;
	v596 = int32(0)
	goto L158
L157:
	;
	v596 = v233
	goto L158
L158:
	;
	v600 = *(*float64)(unsafe.Add(mBase, uint32(v586+v596<<(uint(int32(3))%32))))
	if base.F64_gt(v600, float64(0)) == int32(0) {
		goto L152
	} else {
		goto L159
	}
L159:
	;
	goto L55
L160:
	;
	goto L55
L161:
	;
	v616 = v587
	goto L151
L162:
	;
	v621 = v575 + int32(1)
	if v235 != v621 {
		v575 = v621
		goto L149
	} else {
		goto L163
	}
L163:
	;
	goto L150
L164:
	;
	return v871
L165:
	;
	v871 = v853
	goto L164
L166:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v719 = int32(2147483647)
	v720 = v718 & v719
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v723 = v721 & v719
	if base.Ui32(v720) < base.Ui32(v723) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v731 = v720
	goto L170
L168:
	;
	goto L169
L169:
	;
	if base.Ui32(v720) < base.Ui32(v723) {
		goto L178
	} else {
		goto L179
	}
L170:
	;
	v745 = l1 + int32(8) + v731<<(uint(int32(3))%32)
	v746 = *(*float64)(unsafe.Add(mBase, uint32(v745)))
	if base.F64_ne(v746, float64(0)) != 0 {
		v853 = v701
		goto L165
	} else {
		goto L172
	}
L171:
	;
	goto L169
L172:
	;
	if base.B2i32(v721 < int32(0)) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v754 = *(*float64)(unsafe.Add(mBase, uint32(v745+v723<<(uint(int32(3))%32))))
	if base.F64_ne(v754, float64(0)) != 0 {
		v853 = v701
		goto L165
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v758 = v731 + int32(1)
	if v758 != v723 {
		v731 = v758
		goto L170
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	goto L171
L178:
	;
	v775 = v720
	goto L180
L179:
	;
	v775 = v723
	goto L180
L180:
	;
	if v775 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v871 = int32(1)
	goto L164
L182:
	;
	goto L183
L183:
	;
	v779 = int32(8)
	v786 = int32(0)
	goto L184
L184:
	;
	v798 = int32(0)
	v800 = v786 << (uint(int32(3)) % 32)
	v801 = l0 + v779 + v800
	v802 = *(*float64)(unsafe.Add(mBase, uint32(v801)))
	v804 = base.B2i32(v718 < v798)
	if v718 < v798 {
		v811 = v802
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v853 = v839
	goto L165
L186:
	;
	v812 = v800 + (l1 + v779)
	v813 = *(*float64)(unsafe.Add(mBase, uint32(v812)))
	v815 = base.B2i32(v721 < int32(0))
	if v721 < int32(0) {
		v822 = v813
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v808 = *(*float64)(unsafe.Add(mBase, uint32(v801+v718<<(uint(int32(3))%32))))
	if base.F64_lt(v802, v808) != 0 {
		v811 = v802
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v811 = v808
	goto L186
L189:
	;
	if base.F64_gt(v811, v822) != 0 {
		v853 = v798
		goto L165
	} else {
		goto L192
	}
L190:
	;
	v819 = *(*float64)(unsafe.Add(mBase, uint32(v812+v721<<(uint(int32(3))%32))))
	if base.F64_lt(v813, v819) != 0 {
		v822 = v813
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v822 = v819
	goto L189
L192:
	;
	if v718 < v798 {
		v829 = v802
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if v721 < int32(0) {
		v836 = v813
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v827 = *(*float64)(unsafe.Add(mBase, uint32(v801+v718<<(uint(int32(3))%32))))
	if base.F64_gt(v802, v827) != 0 {
		v829 = v802
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v829 = v827
	goto L193
L196:
	;
	if base.F64_gt(v836, v829) != 0 {
		v853 = v798
		goto L165
	} else {
		goto L199
	}
L197:
	;
	v834 = *(*float64)(unsafe.Add(mBase, uint32(v812+v721<<(uint(int32(3))%32))))
	if base.F64_gt(v813, v834) != 0 {
		v836 = v813
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v836 = v834
	goto L196
L199:
	;
	v839 = int32(1)
	v841 = v786 + v839
	if v841 != v775 {
		v786 = v841
		goto L184
	} else {
		goto L200
	}
L200:
	;
	goto L185
L201:
	;
	v1044 = v1043
	goto L1
L202:
	;
	v1043 = v1025
	goto L201
L203:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v891 = int32(2147483647)
	v892 = v890 & v891
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v895 = v893 & v891
	if base.Ui32(v892) < base.Ui32(v895) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v903 = v892
	goto L207
L205:
	;
	goto L206
L206:
	;
	if base.Ui32(v892) < base.Ui32(v895) {
		goto L215
	} else {
		goto L216
	}
L207:
	;
	v917 = l0 + int32(8) + v903<<(uint(int32(3))%32)
	v918 = *(*float64)(unsafe.Add(mBase, uint32(v917)))
	if base.F64_ne(v918, float64(0)) != 0 {
		v1025 = v873
		goto L202
	} else {
		goto L209
	}
L208:
	;
	goto L206
L209:
	;
	if base.B2i32(v893 < int32(0)) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v926 = *(*float64)(unsafe.Add(mBase, uint32(v917+v895<<(uint(int32(3))%32))))
	if base.F64_ne(v926, float64(0)) != 0 {
		v1025 = v873
		goto L202
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v930 = v903 + int32(1)
	if v930 != v895 {
		v903 = v930
		goto L207
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	goto L208
L215:
	;
	v947 = v892
	goto L217
L216:
	;
	v947 = v895
	goto L217
L217:
	;
	if v947 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1043 = int32(1)
	goto L201
L219:
	;
	goto L220
L220:
	;
	v951 = int32(8)
	v958 = int32(0)
	goto L221
L221:
	;
	v970 = int32(0)
	v972 = v958 << (uint(int32(3)) % 32)
	v973 = l1 + v951 + v972
	v974 = *(*float64)(unsafe.Add(mBase, uint32(v973)))
	v976 = base.B2i32(v890 < v970)
	if v890 < v970 {
		v983 = v974
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v1025 = v1011
	goto L202
L223:
	;
	v984 = v972 + (l0 + v951)
	v985 = *(*float64)(unsafe.Add(mBase, uint32(v984)))
	v987 = base.B2i32(v893 < int32(0))
	if v893 < int32(0) {
		v994 = v985
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v980 = *(*float64)(unsafe.Add(mBase, uint32(v973+v890<<(uint(int32(3))%32))))
	if base.F64_lt(v974, v980) != 0 {
		v983 = v974
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v983 = v980
	goto L223
L226:
	;
	if base.F64_gt(v983, v994) != 0 {
		v1025 = v970
		goto L202
	} else {
		goto L229
	}
L227:
	;
	v991 = *(*float64)(unsafe.Add(mBase, uint32(v984+v893<<(uint(int32(3))%32))))
	if base.F64_lt(v985, v991) != 0 {
		v994 = v985
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v994 = v991
	goto L226
L229:
	;
	if v890 < v970 {
		v1001 = v974
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if v893 < int32(0) {
		v1008 = v985
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v999 = *(*float64)(unsafe.Add(mBase, uint32(v973+v890<<(uint(int32(3))%32))))
	if base.F64_gt(v974, v999) != 0 {
		v1001 = v974
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1001 = v999
	goto L230
L233:
	;
	if base.F64_gt(v1008, v1001) != 0 {
		v1025 = v970
		goto L202
	} else {
		goto L236
	}
L234:
	;
	v1006 = *(*float64)(unsafe.Add(mBase, uint32(v984+v893<<(uint(int32(3))%32))))
	if base.F64_gt(v985, v1006) != 0 {
		v1008 = v985
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1008 = v1006
	goto L233
L236:
	;
	v1011 = int32(1)
	v1013 = v958 + v1011
	if v1013 != v947 {
		v958 = v1013
		goto L221
	} else {
		goto L237
	}
L237:
	;
	goto L222
}
