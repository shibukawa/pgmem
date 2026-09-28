package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DateTimeParseError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	v12 = F_errsave_start(m, l4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		switch l0 + int32(7) {
		case 0:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(22))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v87
					F_errmsg(m, int32(_a_F_DateTimeParseError_0), v10+int32(112))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v94
						v99 = F_errdetail(m, int32(_a_F_DateTimeParseError_1), v10+int32(96))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							v113 = int32(_a_F_DateTimeParseError_2)
							F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						}
					}
				}
			}
		case 1:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v74
					F_errmsg(m, int32(_a_F_DateTimeParseError_0), v10+int32(80))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v113 = int32(_a_F_DateTimeParseError_5)
						F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			}
		case 2:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(150995074))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l2
					F_errmsg(m, int32(_a_F_DateTimeParseError_6), v10-int32(-64))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v113 = int32(_a_F_DateTimeParseError_7)
						F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			}
		case 3:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(84148354))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l2
					F_errmsg(m, int32(_a_F_DateTimeParseError_8), v10+int32(48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v113 = int32(_a_F_DateTimeParseError_9)
						F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			}
		case 4:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l2
					F_errmsg(m, int32(_a_F_DateTimeParseError_10), v10+int32(32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_DateTimeParseError_11), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v113 = int32(_a_F_DateTimeParseError_12)
							F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						}
					}
				}
			}
		case 5:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
					F_errmsg(m, int32(_a_F_DateTimeParseError_10), v10+int32(16))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v113 = int32(_a_F_DateTimeParseError_13)
						F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			}
		default:
			if v12 == int32(0) {
				m.G0 = v10 + int32(128)
				return
			} else {
				F_errcode(m, int32(117440642))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l3
					F_errmsg(m, int32(_a_F_DateTimeParseError_14), v10)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						v113 = int32(_a_F_DateTimeParseError_15)
						F_errsave_finish(m, l4, int32(_a_F_DateTimeParseError_3), v113, int32(_a_F_DateTimeParseError_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			}
		}
	}
}
func F_ParseDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var __phi640 int32
	_ = __phi640
	var v642 int32
	_ = v642
	var __phi642 int32
	_ = __phi642
	var v646 int32
	_ = v646
	var __phi646 int32
	_ = __phi646
	var v647 int32
	_ = v647
	var __phi647 int32
	_ = __phi647
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v709 int32
	_ = v709
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v768 int32
	_ = v768
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v900 int32
	_ = v900
	v16 = l1 + l2
	v17 = l0
	v18 = l1
	v28 = int32(0)
	goto L2
L1:
	;
	return v900
L2:
	;
	v33 = v28 << (uint(int32(2)) % 32)
	v34 = l3 + v33
	v39 = v17
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v28
	v900 = int32(0)
	goto L1
L4:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.B2i32(base.Ui32(v52-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v52 == int32(32)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	if v52 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v39 = v39 + int32(1)
	goto L4
L10:
	;
	v64 = int32(-1)
	if int32(25) <= v28 {
		v900 = v64
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v18
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32((v66-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	if base.Ui32(v66-int32(33)) <= base.Ui32(int32(93)) {
		goto L177
	} else {
		goto L178
	}
L13:
	;
	v865 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v857))) = uint8(v865)
	v867 = int32(1)
	v17 = v850
	v18 = v857 + v867
	v28 = v28 + v867
	goto L2
L14:
	;
	if base.Ui32(v66|int32(32)-int32(97)) <= base.Ui32(int32(25)) {
		goto L83
	} else {
		goto L84
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(0)
	v850 = v411
	v857 = v418
	goto L13
L16:
	;
	v74 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v74) {
		v900 = v64
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v66 != int32(46) {
		goto L14
	} else {
		goto L76
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v66)
	v78 = v39 + int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32((v79-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = v78
	v92 = v79
	v93 = v74
	goto L23
L21:
	;
	v114 = v78
	v116 = v39
	v120 = v79
	v121 = v74
	goto L22
L22:
	;
	v130 = v120 & int32(255)
	switch v130 - int32(45) {
	case 0, 1, 2:
		goto L27
	default:
		v411 = v114
		v418 = v121
		goto L15
	case 13:
		goto L28
	}
L23:
	;
	v102 = v93 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v102) {
		v900 = v64
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v114 = v106
	v116 = v86
	v120 = v107
	v121 = v102
	goto L22
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v92)
	v106 = v86 + int32(1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if base.Ui32((v107-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v86 = v106
		v92 = v107
		v93 = v102
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v180 = v121 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v180) {
		v900 = v64
		goto L1
	} else {
		goto L36
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(3)
	v137 = v121 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v137) {
		v900 = v64
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v139)
	v143 = v116 + int32(2)
	v150 = v137
	goto L30
L30:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if base.Ui32((v158-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v174 = v150 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v174) {
		v900 = v64
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v166 = v158 - int32(46)
	if base.B2i32(v166 == int32(0))|base.B2i32(v166 == int32(12)) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v850 = v143
	v857 = v150
	goto L13
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v158)
	v143 = v143 + int32(1)
	v150 = v174
	goto L30
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v120)
	v184 = v116 + int32(2)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if base.Ui32((v185-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v192 = l4 + v33
	if v130 != int32(46) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(2)
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	goto L62
L40:
	;
	v197 = int32(2)
	goto L42
L41:
	;
	v197 = int32(0)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v197
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if base.Ui32((v199-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v207 = v180
	v208 = v199
	v212 = v184
	goto L46
L44:
	;
	v235 = v180
	v236 = v199
	v240 = v184
	goto L45
L45:
	;
	if v130 != v236&int32(255) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v222 = v207 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v222) {
		v900 = v64
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v235 = v222
	v236 = v225
	v240 = v227
	goto L45
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v208)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
	v227 = v212 + int32(1)
	if base.Ui32((v225-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v207 = v222
		v208 = v225
		v212 = v227
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v850 = v240
	v857 = v235
	goto L13
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = int32(2)
	v255 = v235 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v255) {
		v900 = v64
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v257)
	v260 = v240 + int32(1)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v130 != v261 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.Ui32(int32(9)) < base.Ui32((v261-int32(48))&int32(255)) {
		v850 = v260
		v857 = v255
		goto L13
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v272 = v260
	v273 = v255
	v277 = v261
	goto L58
L57:
	;
	goto L56
L58:
	;
	v287 = v273 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v287) {
		v900 = v64
		goto L1
	} else {
		goto L60
	}
L59:
	;
	v850 = v292
	v857 = v287
	goto L13
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v277)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v292 = v272 + int32(1)
	if base.B2i32(v290 == v130)|base.B2i32(base.Ui32((v290-int32(48))&int32(255)) < base.Ui32(int32(10))) != 0 {
		v272 = v292
		v273 = v287
		v277 = v290
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	if v304 != v130 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if base.B2i32(base.Ui32(v304-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v304|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		v850 = v184
		v857 = v180
		goto L13
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v322 = v180
	v323 = v304
	v327 = v184
	goto L67
L66:
	;
	goto L65
L67:
	;
	v337 = v322 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v337) {
		v900 = v64
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v850 = v364
	v857 = v337
	goto L13
L69:
	;
	v339 = int32(255)
	v340 = v323 & v339
	if base.Ui32((v340-int32(65))&v339) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v322))) = uint8(v349)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+1)))
	goto L74
L71:
	;
	v349 = v340 | int32(32)
	goto L73
L72:
	;
	v349 = v340
	goto L73
L73:
	;
	goto L70
L74:
	;
	v364 = v327 + int32(1)
	if base.B2i32(v351 == v130)|(base.B2i32(base.Ui32(v351-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v351|int32(32)-int32(97)) < base.Ui32(int32(26)))) != 0 {
		v322 = v337
		v323 = v351
		v327 = v364
		goto L67
	} else {
		goto L75
	}
L75:
	;
	goto L68
L76:
	;
	v370 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v370) {
		v900 = v64
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v372 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v372)
	v375 = v39 + int32(1)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v376-int32(48))&int32(255)) {
		v411 = v375
		v418 = v370
		goto L15
	} else {
		goto L78
	}
L78:
	;
	v384 = v375
	v385 = v376
	v389 = v370
	goto L79
L79:
	;
	v399 = v389 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v399) {
		v900 = v64
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v411 = v404
	v418 = v399
	goto L15
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v385)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	v404 = v384 + int32(1)
	if base.Ui32((v402-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v384 = v404
		v385 = v402
		v389 = v399
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v435 = l4 + v33
	v436 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v436
	v439 = v18 + v436
	if base.Ui32(v16) <= base.Ui32(v439) {
		v900 = v64
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	switch v66 - int32(43) {
	case 0, 2:
		goto L145
	default:
		goto L12
	}
L86:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32((v441-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v450)
	v453 = v39 + int32(1)
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if base.Ui32((v454|int32(32)-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v450 = v441 | int32(32)
	goto L90
L89:
	;
	v450 = v441
	goto L90
L90:
	;
	goto L87
L91:
	;
	v463 = v453
	v465 = v454
	v470 = v439
	goto L94
L92:
	;
	v504 = v453
	v506 = v454
	v511 = v439
	goto L93
L93:
	;
	switch v506&int32(255) - int32(43) {
	case 0:
		goto L103
	default:
		goto L104
	case 2, 3, 4:
		goto L102
	}
L94:
	;
	v479 = v470 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v479) {
		v900 = v64
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v504 = v495
	v506 = v493
	v511 = v479
	goto L93
L96:
	;
	v481 = int32(255)
	v482 = v465 & v481
	if base.Ui32((v482-int32(65))&v481) < base.Ui32(int32(26)) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v470))) = uint8(v491)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	v495 = v463 + int32(1)
	if base.Ui32((v493|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v463 = v495
		v465 = v493
		v470 = v479
		goto L94
	} else {
		goto L101
	}
L98:
	;
	v491 = v482 | int32(32)
	goto L100
L99:
	;
	v491 = v482
	goto L100
L100:
	;
	goto L97
L101:
	;
	goto L95
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = int32(2)
	v637 = v511 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v637) {
		v900 = v64
		goto L1
	} else {
		goto L133
	}
L103:
	;
	v529 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v511))) = uint8(v529)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v534 = int32(*(*int8)(unsafe.Add(mBase, uint32(v533))))
	v541 = int32(_a_F_ParseDateTime_0)
	v544 = int32(_a_F_ParseDateTime_1)
	goto L107
L104:
	;
	if base.Ui32(int32(9)) < base.Ui32((v506-int32(48))&int32(255)) {
		v850 = v504
		v857 = v511
		goto L13
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	if v618 != 0 {
		v850 = v504
		v857 = v511
		goto L13
	} else {
		goto L132
	}
L107:
	;
	v555 = v541 + (v544-v541)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v556 = int32(*(*int8)(unsafe.Add(mBase, uint32(v555))))
	v557 = v534 - v556
	if v557 != 0 {
		v604 = v557
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v618 = int32(0)
	goto L106
L109:
	;
	v608 = base.B2i32(v604 < int32(0))
	if v604 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L110:
	;
	goto L113
L111:
	;
	if v597 != 0 {
		v604 = v597
		goto L109
	} else {
		goto L124
	}
L113:
	;
	goto L114
L114:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v564 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v565 = v533
	v566 = v555
	v567 = int32(10)
	v568 = v564
	goto L119
L116:
	;
	v591 = v555
	v595 = int32(0)
	goto L117
L117:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591))))
	v597 = v595 - v596
	goto L111
L118:
	;
	v591 = v586
	v595 = v588
	goto L117
L119:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if base.B2i32(v568 != v570)|base.B2i32(v570 == int32(0)) != 0 {
		v586 = v566
		v588 = v568
		goto L118
	} else {
		goto L121
	}
L120:
	;
	v586 = v580
	v588 = int32(0)
	goto L118
L121:
	;
	v576 = v567 - int32(1)
	if v576 == int32(0) {
		v586 = v566
		v588 = v568
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v579 = int32(1)
	v580 = v566 + v579
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	if v581 != 0 {
		v565 = v565 + v579
		v566 = v580
		v567 = v576
		v568 = v581
		goto L119
	} else {
		goto L123
	}
L123:
	;
	goto L120
L124:
	;
	v618 = v555
	goto L106
L125:
	;
	v609 = v555 - int32(16)
	goto L127
L126:
	;
	v609 = v544
	goto L127
L127:
	;
	if v604 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v612 = v541
	goto L130
L129:
	;
	v612 = v555 + int32(16)
	goto L130
L130:
	;
	if base.Ui32(v612) <= base.Ui32(v609) {
		v541 = v612
		v544 = v609
		goto L107
	} else {
		goto L131
	}
L131:
	;
	goto L108
L132:
	;
	goto L102
L133:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	__phi640 = v504
	__phi642 = v639
	__phi646 = v637
	__phi647 = v511
	v640 = __phi640
	v642 = __phi642
	v646 = __phi646
	v647 = __phi647
	goto L134
L134:
	;
	v655 = int32(255)
	v656 = v642 & v655
	if base.Ui32((v656-int32(65))&v655) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v900 = v64
	goto L1
L136:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v647))) = uint8(v665)
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+1)))
	v669 = v640 + int32(1)
	switch v667 - int32(43) {
	case 0, 2, 3, 4, 15, 52:
		goto L140
	default:
		goto L141
	}
L137:
	;
	v665 = v656 | int32(32)
	goto L139
L138:
	;
	v665 = v656
	goto L139
L139:
	;
	goto L136
L140:
	;
	v684 = v646 + int32(1)
	if base.Ui32(v684) < base.Ui32(v16) {
		__phi640 = v669
		__phi642 = v667
		__phi646 = v684
		__phi647 = v646
		v640 = __phi640
		v642 = __phi642
		v646 = __phi646
		v647 = __phi647
		goto L134
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	if base.B2i32(base.Ui32(v667-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v667|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v850 = v669
	v857 = v646
	goto L13
L144:
	;
	goto L135
L145:
	;
	v689 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v689) {
		v900 = v64
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v66)
	v694 = v39
	goto L147
L147:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	if base.B2i32(base.Ui32(v709-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v709 == int32(32)) != 0 {
		v694 = v694 + int32(1)
		goto L147
	} else {
		goto L149
	}
L148:
	;
	if base.Ui32((v709-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(4)
	v727 = v18 + int32(2)
	if base.Ui32(v16) <= base.Ui32(v727) {
		v900 = v64
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	if base.Ui32(int32(25)) < base.Ui32((v709|int32(32)-int32(97))&int32(255)) {
		v900 = v64
		goto L1
	} else {
		goto L161
	}
L153:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v729)
	v733 = v694 + int32(2)
	v740 = v727
	goto L154
L154:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	if base.Ui32(int32(10)) <= base.Ui32((v748-int32(48))&int32(255)) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v756 = v748 - int32(45)
	if base.B2i32(base.Ui32(int32(13)) < base.Ui32(v756))|base.B2i32(int32(1)<<(uint(v756)%32)&int32(_a_F_ParseDateTime_2) == int32(0)) != 0 {
		v850 = v733
		v857 = v740
		goto L13
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v768 = v740 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v768) {
		v900 = v64
		goto L1
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v740))) = uint8(v748)
	v733 = v733 + int32(1)
	v740 = v768
	goto L154
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(6)
	v785 = v18 + int32(2)
	if base.Ui32(v16) <= base.Ui32(v785) {
		v900 = v64
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	if base.Ui32((v787-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v796)
	v799 = v694 + int32(2)
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+2)))
	if base.Ui32(int32(25)) < base.Ui32((v800|int32(32)-int32(97))&int32(255)) {
		v850 = v799
		v857 = v785
		goto L13
	} else {
		goto L167
	}
L164:
	;
	v796 = v787 | int32(32)
	goto L166
L165:
	;
	v796 = v787
	goto L166
L166:
	;
	goto L163
L167:
	;
	v810 = v799
	v811 = v785
	v815 = v800
	goto L168
L168:
	;
	v825 = v811 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v825) {
		v900 = v64
		goto L1
	} else {
		goto L170
	}
L169:
	;
	v850 = v841
	v857 = v825
	goto L13
L170:
	;
	v827 = int32(255)
	v828 = v815 & v827
	if base.Ui32((v828-int32(65))&v827) < base.Ui32(int32(26)) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v811))) = uint8(v837)
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810)+1)))
	v841 = v810 + int32(1)
	if base.Ui32((v839|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v810 = v841
		v811 = v825
		v815 = v839
		goto L168
	} else {
		goto L175
	}
L172:
	;
	v837 = v828 | int32(32)
	goto L174
L173:
	;
	v837 = v828
	goto L174
L174:
	;
	goto L171
L175:
	;
	goto L169
L176:
	;
	if base.B2i32(v877 == int32(0)) == int32(0) {
		v900 = v64
		goto L1
	} else {
		goto L180
	}
L177:
	;
	v875 = F_isalnum(m, v66)
	mBase = m.M
	v877 = v875
	goto L179
L178:
	;
	v877 = int32(1)
	goto L179
L179:
	;
	goto L176
L180:
	;
	goto L9
}
func F_date_cmp_timestamp_internal(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_date_cmp_timestamp_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_date_cmp_timestamp_internal[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
	if l0 == int32(-2147483648) {
		v45 = int64(-9223372036854775807 - 1)
		v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v46 == int32(1) {
			if l1 == int64(9223372036854775807) {
				v53 = int32(-1)
			} else {
				v53 = int32(1)
			}
			v57 = v53
		} else {
			v57 = base.B2i32(l1 < v45) - base.B2i32(v45 < l1)
		}
		m.G0 = v7 + int32(16)
		return v57
	} else {
		if l0 == int32(2147483647) {
			v45 = int64(9223372036854775807)
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
			if v46 == int32(1) {
				if l1 == int64(9223372036854775807) {
					v53 = int32(-1)
				} else {
					v53 = int32(1)
				}
				v57 = v53
			} else {
				v57 = base.B2i32(l1 < v45) - base.B2i32(v45 < l1)
			}
			m.G0 = v7 + int32(16)
			return v57
		} else {
			if int32(106751983) <= l0 {
				v23 = int64(9223372036854775807)
				v24 = F_errsave_start(m, v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v45 = v23
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
						if v46 == int32(1) {
							if l1 == int64(9223372036854775807) {
								v53 = int32(-1)
							} else {
								v53 = int32(1)
							}
							v57 = v53
						} else {
							v57 = base.B2i32(l1 < v45) - base.B2i32(v45 < l1)
						}
						m.G0 = v7 + int32(16)
						return v57
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_date_cmp_timestamp_internal_0), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v7, int32(_a_F_date_cmp_timestamp_internal_1), int32(646), int32(_a_F_date_cmp_timestamp_internal_2))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v45 = v23
									v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
									if v46 == int32(1) {
										if l1 == int64(9223372036854775807) {
											v53 = int32(-1)
										} else {
											v53 = int32(1)
										}
										v57 = v53
									} else {
										v57 = base.B2i32(l1 < v45) - base.B2i32(v45 < l1)
									}
									m.G0 = v7 + int32(16)
									return v57
								}
							}
						}
					}
				}
			} else {
				v45 = base.I64_extend_i32_s(l0) * int64(86400000000)
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
				if v46 == int32(1) {
					if l1 == int64(9223372036854775807) {
						v53 = int32(-1)
					} else {
						v53 = int32(1)
					}
					v57 = v53
				} else {
					v57 = base.B2i32(l1 < v45) - base.B2i32(v45 < l1)
				}
				m.G0 = v7 + int32(16)
				return v57
			}
		}
	}
}
func F_date_eq_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v35 int64
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_date_eq_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_date_eq_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 == int32(1) {
			if base.Ui64(v17-int64(9223372036854775807)) < base.Ui64(int64(2)) {
				v35 = int64(0)
			} else {
				v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) == int32(0)))
			}
		} else {
			v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) == int32(0)))
		}
		m.G0 = v7 + int32(16)
		return v35
	}
}
func F_date_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 <= v3))
}
func F_date_lt_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_date_lt_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_date_lt_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(int32(base.Ui32(v41) >> (uint(int32(31)) % 32)))
	}
}
func F_date_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_date_pli(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.I32_wrap_i64(v5)
	if base.Ui32(v6-int32(2147483647)) <= base.Ui32(int32(1)) {
		return base.I64_extend32_s(v5)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = v13 + v6
		if int32(0) <= v13 {
			if v6 <= v14 {
				if base.Ui32(int32(2147483494)) <= base.Ui32(v14+int32(_a_F_date_pli_0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_date_pli_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_date_pli_2), int32(582), int32(_a_F_date_pli_3))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					return base.I64_extend_i32_s(v14)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_date_pli_1), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_date_pli_2), int32(582), int32(_a_F_date_pli_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
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
			if v6 < v14 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_date_pli_1), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_date_pli_2), int32(582), int32(_a_F_date_pli_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(2147483494)) <= base.Ui32(v14+int32(_a_F_date_pli_0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_date_pli_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_date_pli_2), int32(582), int32(_a_F_date_pli_3))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					return base.I64_extend_i32_s(v14)
				}
			}
		}
	}
}
func F_extract_date(m *base.Module, l0 int32) int64 {
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
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = int32(1)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		v21 = v19 & v17
		if v21 != 0 {
			v22 = v17
		} else {
			v22 = int32(4)
		}
		if v19 == int32(1) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v29 == int32(18) {
				v32 = int32(16)
			} else {
				v32 = int32(0)
			}
			if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v39 = int32(4)
			} else {
				v39 = v32
			}
			v50 = v39
		} else {
			v40 = int32(1)
			if v21 != 0 {
				v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v52 = F_downcase_truncate_identifier(m, v12+v22, v50, int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			v55 = v9 + int32(76)
			v59 = Fn14210(m, v52, v55, int32(_a_F_extract_date_0), int32(_a_F_extract_date_1), int32(_a_F_extract_date_2))
			mBase = m.M
			if v59 == int32(31) {
				v65 = Fn14210(m, v52, v55, int32(_a_F_extract_date_3), int32(_a_F_extract_date_4), int32(_a_F_extract_date_5))
				mBase = m.M
				v66 = v65
			} else {
				v66 = v59
			}
			v67 = base.I32_wrap_i64(v16)
			if base.Ui32(v67-int32(2147483647)) <= base.Ui32(int32(1)) {
				if v66 != int32(17) {
					v75 = v66
				} else {
					v75 = int32(0)
				}
				if v75 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v506 = m.ExcPending
					if v506 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v509 = m.ExcPending
						if v509 != 0 {
							return int64(0)
						} else {
							v511 = F_format_type_be(m, int32(1082))
							mBase = m.M
							v512 = m.ExcPending
							if v512 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v511
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v52
								F_errmsg(m, int32(_a_F_extract_date_6), v9)
								mBase = m.M
								v517 = m.ExcPending
								if v517 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_extract_date_7), int32(1258), int32(_a_F_extract_date_8))
									mBase = m.M
									v522 = m.ExcPending
									if v522 != 0 {
										return int64(0)
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
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
					v78 = v76 - int32(11)
					if base.Ui32(int32(26)) < base.Ui32(v78) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								v125 = F_format_type_be(m, int32(1082))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v125
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v52
									F_errmsg(m, int32(_a_F_extract_date_9), v9+int32(16))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_extract_date_7), int32(1150), int32(_a_F_extract_date_8))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int64(0)
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
						v82 = int32(1) << (uint(v78) % 32)
						if v82&int32(34848769) == int32(0) {
							if v82&int32(73415680) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int64(0)
									} else {
										v125 = F_format_type_be(m, int32(1082))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v125
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v52
											F_errmsg(m, int32(_a_F_extract_date_9), v9+int32(16))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_extract_date_7), int32(1150), int32(_a_F_extract_date_8))
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int64(0)
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
								v91 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v91)
								v532 = int32(0)
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						} else {
							if v67 == int32(-2147483648) {
								v101 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11926), int64(0), int64(-1))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int64(0)
								} else {
									v104 = F_pg_detoast_datum(m, base.I32_wrap_i64(v101))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										v532 = v104
										m.G0 = v9 + int32(80)
										return base.I64_extend_i32_u(v532)
									}
								}
							} else {
								v111 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11937), int64(0), int64(-1))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v114 = F_pg_detoast_datum(m, base.I32_wrap_i64(v111))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int64(0)
									} else {
										v532 = v114
										m.G0 = v9 + int32(80)
										return base.I64_extend_i32_u(v532)
									}
								}
							}
						}
					}
				}
			} else {
				if v66 != 0 {
					if v66 != int32(17) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v506 = m.ExcPending
						if v506 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v509 = m.ExcPending
							if v509 != 0 {
								return int64(0)
							} else {
								v511 = F_format_type_be(m, int32(1082))
								mBase = m.M
								v512 = m.ExcPending
								if v512 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v511
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v52
									F_errmsg(m, int32(_a_F_extract_date_6), v9)
									mBase = m.M
									v517 = m.ExcPending
									if v517 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_extract_date_7), int32(1258), int32(_a_F_extract_date_8))
										mBase = m.M
										v522 = m.ExcPending
										if v522 != 0 {
											return int64(0)
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
						v142 = v67 + int32(_a_F_extract_date_10)
						v152 = v67 + int32(_a_F_extract_date_11)
						v153 = int32(_a_F_extract_date_12)
						v154 = base.I32_div_u_s(v152, v153)
						v155 = int32(3)
						v161 = int32(2)
						v166 = base.I32_div_u_s((v154*int32(1073595727)+v152)<<(uint(v161)%32)|v155, v153)
						v169 = v142 + v154*v155 + v166 + int32(_a_F_extract_date_13)
						v170 = int32(1461)
						v171 = base.I32_div_u_s(v169, v170)
						v174 = v171*int32(-1461) + v169
						v176 = v174 << (uint(v161) % 32)
						if base.Ui32(v170) <= base.Ui32(v176) {
							v182 = base.I32_rem_u_s(v174+int32(305), int32(365))
							v187 = v182
						} else {
							v186 = base.I32_rem_u_s(v174+int32(306), int32(366))
							v187 = v186
						}
						v189 = base.I32_div_u_s(v176, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(72)))) = v189 + v171<<(uint(int32(2))%32) - int32(_a_F_extract_date_14)
						v197 = v187 + int32(123)
						v201 = int32(base.Ui32(v197*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = v197 - int32(base.Ui32(v201*int32(_a_F_extract_date_15))>>(uint(int32(8))%32))
						v211 = base.I32_rem_u_s(v201+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v9+int32(68)))) = v211 + int32(1)
						v215 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
						switch v215 - int32(21) {
						case 0:
							v523 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+64)))
							v526 = v523
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 1:
							v227 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							v228 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
							v229 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
							v231 = F_date2j(m, v227, v228, v229)
							mBase = m.M
							v232 = int32(1)
							v234 = F_date2j(m, v227, v232, int32(4))
							mBase = m.M
							v237 = F_j2day(m, v234-v232)
							mBase = m.M
							if v231 < v234-v237 {
								v240 = int32(1)
								v244 = F_date2j(m, v227-v240, v240, int32(4))
								mBase = m.M
								v247 = F_j2day(m, v244-v240)
								mBase = m.M
								v248 = v244
								v249 = v247
							} else {
								v248 = v234
								v249 = v237
							}
							v251 = v249 - v248 + v231
							if int32(357) <= v251 {
								v254 = int32(1)
								v258 = F_date2j(m, v227+v254, v254, int32(4))
								mBase = m.M
								v261 = F_j2day(m, v258-v254)
								mBase = m.M
								v262 = v258 - v261
								if v231 < v262 {
									v265 = v251
								} else {
									v265 = v231 - v262
								}
								v267 = v265
							} else {
								v267 = v251
							}
							v269 = base.I32_div_s(v267, int32(7))
							v526 = base.I64_extend_i32_s(v269 + int32(1))
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 2:
							v218 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+68)))
							v526 = v218
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 3:
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
							v220 = int32(1)
							v223 = base.I32_div_s(v219-v220, int32(3))
							v526 = base.I64_extend_i32_s(v223 + v220)
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 4:
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							if int32(0) < v273 {
								v526 = base.I64_extend_i32_u(v273)
							} else {
								v526 = base.I64_extend_i32_s(v273 - int32(1))
							}
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 5:
							v280 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							if int32(0) <= v280 {
								v284 = base.I32_div_u_s(v280, int32(10))
								v526 = base.I64_extend_i32_u(v284)
							} else {
								v289 = base.I32_div_s(int32(9)-v280, int32(-10))
								v526 = base.I64_extend_i32_s(v289)
							}
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 6:
							v291 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							if int32(0) < v291 {
								v297 = base.I32_div_s(v291+int32(99), int32(100))
								v526 = base.I64_extend_i32_s(v297)
							} else {
								v302 = base.I32_div_s(int32(100)-v291, int32(-100))
								v526 = base.I64_extend_i32_s(v302)
							}
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 7:
							v304 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							if int32(0) < v304 {
								v310 = base.I32_div_s(v304+int32(999), int32(1000))
								v526 = base.I64_extend_i32_s(v310)
							} else {
								v315 = base.I32_div_s(int32(1000)-v304, int32(-1000))
								v526 = base.I64_extend_i32_s(v315)
							}
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v454 = m.ExcPending
							if v454 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v457 = m.ExcPending
								if v457 != 0 {
									return int64(0)
								} else {
									v459 = F_format_type_be(m, int32(1082))
									mBase = m.M
									v460 = m.ExcPending
									if v460 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v459
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v52
										F_errmsg(m, int32(_a_F_extract_date_9), v9+int32(32))
										mBase = m.M
										v467 = m.ExcPending
										if v467 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_extract_date_7), int32(1233), int32(_a_F_extract_date_8))
											mBase = m.M
											v472 = m.ExcPending
											if v472 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						case 10:
							v526 = base.I64_extend_i32_s(v142)
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 11, 16:
							v367 = int32(7)
							v368 = base.I32_rem_s(v67+int32(_a_F_extract_date_16), v367)
							if v368 < int32(0) {
								v373 = v368 + v367
							} else {
								v373 = v368
							}
							v374 = base.I64_extend_i32_s(v373)
							if v373 != 0 {
								v376 = v374
							} else {
								v376 = int64(7)
							}
							v377 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
							if v377 == int32(37) {
								v380 = v376
							} else {
								v380 = v374
							}
							v526 = v380
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 12:
							v381 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
							v383 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
							v388 = base.B2i32(int32(2) < v382)
							if int32(2) < v382 {
								v389 = int32(_a_F_extract_date_14)
							} else {
								v389 = int32(_a_F_extract_date_17)
							}
							v390 = v389 + v381
							v395 = base.I32_div_s(v390, int32(4))
							v398 = base.I32_div_s(v390, int32(-100))
							v401 = base.I32_div_s(v390, int32(400))
							if int32(2) < v382 {
								v405 = int32(1)
							} else {
								v405 = int32(13)
							}
							v410 = base.I32_div_s((v405+v382)*int32(_a_F_extract_date_15), int32(256))
							v414 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							v423 = int32(_a_F_extract_date_17) + v414
							v428 = base.I32_div_s(v423, int32(4))
							v431 = base.I32_div_s(v423, int32(-100))
							v434 = base.I32_div_s(v423, int32(400))
							v443 = base.I32_div_s(int32(_a_F_extract_date_18), int32(256))
							v526 = base.I64_extend_i32_s(v383 + v390*int32(365) + v395 + v398 + v401 + v410 - int32(_a_F_extract_date_19) - (int32(1) + v423*int32(365) + v428 + v431 + v434 + v443 - int32(_a_F_extract_date_19)) + int32(1))
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						case 15:
							v318 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
							v319 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
							v320 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
							v322 = F_date2j(m, v318, v319, v320)
							mBase = m.M
							v323 = int32(1)
							v325 = F_date2j(m, v318, v323, int32(4))
							mBase = m.M
							v328 = F_j2day(m, v325-v323)
							mBase = m.M
							if v322 < v325-v328 {
								v331 = int32(1)
								v332 = v318 - v331
								v335 = F_date2j(m, v332, v331, int32(4))
								mBase = m.M
								v338 = F_j2day(m, v335-v331)
								mBase = m.M
								v339 = v332
								v340 = v335
								v341 = v338
							} else {
								v339 = v318
								v340 = v325
								v341 = v328
							}
							if int32(357) <= v341+(v322-v340) {
								v346 = int32(1)
								v347 = v339 + v346
								v350 = F_date2j(m, v347, v346, int32(4))
								mBase = m.M
								v353 = F_j2day(m, v350-v346)
								mBase = m.M
								if v322 < v350-v353 {
									v356 = v339
								} else {
									v356 = v347
								}
								v359 = v356
							} else {
								v359 = v339
							}
							v526 = base.I64_extend_i32_s(v359) - base.I64_extend_i32_u(base.B2i32(v359 <= int32(0)))
							v527 = F_int64_to_numeric(m, v526)
							mBase = m.M
							v528 = m.ExcPending
							if v528 != 0 {
								return int64(0)
							} else {
								v532 = v527
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v532)
							}
						}
					}
				} else {
					v478 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
					if v478 == int32(11) {
						v526 = base.I64_extend32_s(v16)*int64(86400) + int64(946684800)
						v527 = F_int64_to_numeric(m, v526)
						mBase = m.M
						v528 = m.ExcPending
						if v528 != 0 {
							return int64(0)
						} else {
							v532 = v527
							m.G0 = v9 + int32(80)
							return base.I64_extend_i32_u(v532)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v484 = m.ExcPending
						if v484 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v487 = m.ExcPending
							if v487 != 0 {
								return int64(0)
							} else {
								v489 = F_format_type_be(m, int32(1082))
								mBase = m.M
								v490 = m.ExcPending
								if v490 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v489
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v52
									F_errmsg(m, int32(_a_F_extract_date_9), v9+int32(48))
									mBase = m.M
									v497 = m.ExcPending
									if v497 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_extract_date_7), int32(1249), int32(_a_F_extract_date_8))
										mBase = m.M
										v502 = m.ExcPending
										if v502 != 0 {
											return int64(0)
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
		}
	}
}
