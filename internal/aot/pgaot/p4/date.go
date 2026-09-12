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
					F_errmsg(m, int32(417473), v10+int32(112))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v94
						F_errdetail(m, int32(625129), v10+int32(96))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							v113 = int32(4258)
							F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(417473), v10+int32(80))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v113 = int32(4250)
						F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(685492), v10-int32(-64))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v113 = int32(4244)
						F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(685563), v10+int32(48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v113 = int32(4238)
						F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(685603), v10+int32(32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errhint(m, int32(585572), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v113 = int32(4232)
							F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(685603), v10+int32(16))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v113 = int32(4224)
						F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
					F_errmsg(m, int32(683524), v10)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						v113 = int32(4265)
						F_errsave_finish(m, l4, int32(475778), v113, int32(202123))
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v634 int32
	_ = v634
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var __phi656 int32
	_ = __phi656
	var v657 int32
	_ = v657
	var __phi657 int32
	_ = __phi657
	var v658 int32
	_ = v658
	var __phi658 int32
	_ = __phi658
	var v662 int32
	_ = v662
	var __phi662 int32
	_ = __phi662
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v763 int32
	_ = v763
	var v773 int32
	_ = v773
	var v784 int32
	_ = v784
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v916 int32
	_ = v916
	v16 = l1 + l2
	v17 = l0
	v18 = l1
	v28 = int32(0)
	goto L2
L1:
	;
	return v916
L2:
	;
	v33 = v28 << (uint(int32(2)) % 32)
	v34 = l3 + v33
	v39 = v17
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v28
	v916 = int32(0)
	goto L1
L4:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32(v52-int32(9)) < base.Ui32(int32(5)) {
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
	v39 = v39 + int32(1)
	goto L4
L8:
	;
	if v52 == int32(32) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v52 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v61 = int32(-1)
	if int32(25) <= v28 {
		v916 = v61
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v18
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32((v63-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	if base.Ui32(v63-int32(33)) <= base.Ui32(int32(93)) {
		goto L181
	} else {
		goto L182
	}
L13:
	;
	v881 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v872))) = uint8(v881)
	v883 = int32(1)
	v17 = v866
	v18 = v872 + v883
	v28 = v28 + v883
	goto L2
L14:
	;
	if base.Ui32(v63|int32(32)-int32(97)) <= base.Ui32(int32(25)) {
		goto L81
	} else {
		goto L82
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(0)
	v866 = v399
	v872 = v405
	goto L13
L16:
	;
	v71 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v71) {
		v916 = v61
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v63 != int32(46) {
		goto L14
	} else {
		goto L74
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v63)
	v75 = v39 + int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if base.Ui32((v76-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v83 = v75
	v84 = v76
	v89 = v71
	goto L23
L21:
	;
	v111 = v75
	v112 = v76
	v113 = v39
	v117 = v71
	goto L22
L22:
	;
	v127 = v112 & int32(255)
	switch v127 - int32(45) {
	case 0, 1, 2:
		goto L27
	default:
		v399 = v111
		v405 = v117
		goto L15
	case 13:
		goto L28
	}
L23:
	;
	v99 = v89 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v99) {
		v916 = v61
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v111 = v103
	v112 = v104
	v113 = v83
	v117 = v99
	goto L22
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v84)
	v103 = v83 + int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if base.Ui32((v104-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v83 = v103
		v84 = v104
		v89 = v99
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v173 = v117 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v173) {
		v916 = v61
		goto L1
	} else {
		goto L35
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(3)
	v134 = v117 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v134) {
		v916 = v61
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v136)
	v140 = v113 + int32(2)
	v146 = v134
	goto L30
L30:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if base.Ui32((v155-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v167 = v146 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v167) {
		v916 = v61
		goto L1
	} else {
		goto L34
	}
L33:
	;
	switch v155&int32(255) - int32(46) {
	case 0, 12:
		goto L32
	default:
		v866 = v140
		v872 = v146
		goto L13
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v155)
	v140 = v140 + int32(1)
	v146 = v167
	goto L30
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v112)
	v177 = v113 + int32(2)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if base.Ui32((v178-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v185 = l4 + v33
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = base.B2i32(v127 != int32(46)) << (uint(int32(1)) % 32)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if base.Ui32((v191-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(2)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	goto L59
L39:
	;
	v198 = v177
	v200 = v191
	v205 = v173
	goto L42
L40:
	;
	v226 = v177
	v228 = v191
	v233 = v173
	goto L41
L41:
	;
	if v127 != v228&int32(255) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v214 = v205 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v214) {
		v916 = v61
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v226 = v218
	v228 = v219
	v233 = v214
	goto L41
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v200)
	v218 = v198 + int32(1)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if base.Ui32((v219-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v198 = v218
		v200 = v219
		v205 = v214
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v866 = v226
	v872 = v233
	goto L13
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(2)
	v247 = v233 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v247) {
		v916 = v61
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v249)
	v252 = v226 + int32(1)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v127 != v253 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(int32(9)) < base.Ui32((v253-int32(48))&int32(255)) {
		v866 = v252
		v872 = v247
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v262 = v252
	v263 = v253
	v264 = v247
	goto L54
L53:
	;
	goto L52
L54:
	;
	v278 = v264 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v278) {
		v916 = v61
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v866 = v282
	v872 = v278
	goto L13
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v264))) = uint8(v263)
	v282 = v262 + int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v283 == v127 {
		v262 = v282
		v263 = v283
		v264 = v278
		goto L54
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32((v283-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v262 = v282
		v263 = v283
		v264 = v278
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	if v294 != v127 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if base.B2i32(base.Ui32(v294-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v294|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		v866 = v177
		v872 = v173
		goto L13
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v310 = v177
	v312 = v294
	v317 = v173
	goto L64
L63:
	;
	goto L62
L64:
	;
	v326 = v317 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v326) {
		v916 = v61
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v866 = v341
	v872 = v326
	goto L13
L66:
	;
	v328 = int32(255)
	v329 = v312 & v328
	if base.Ui32((v329-int32(65))&v328) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v338)
	v341 = v310 + int32(1)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	goto L71
L68:
	;
	v338 = v329 | int32(32)
	goto L70
L69:
	;
	v338 = v329
	goto L70
L70:
	;
	goto L67
L71:
	;
	if v342 == v127 {
		v310 = v341
		v312 = v342
		v317 = v326
		goto L64
	} else {
		goto L72
	}
L72:
	;
	if base.B2i32(base.Ui32(v342-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v342|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		v310 = v341
		v312 = v342
		v317 = v326
		goto L64
	} else {
		goto L73
	}
L73:
	;
	goto L65
L74:
	;
	v358 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v358) {
		v916 = v61
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v360 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v360)
	v363 = v39 + int32(1)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if base.Ui32(int32(9)) < base.Ui32((v364-int32(48))&int32(255)) {
		v399 = v363
		v405 = v358
		goto L15
	} else {
		goto L76
	}
L76:
	;
	v371 = v363
	v373 = v364
	v378 = v358
	goto L77
L77:
	;
	v387 = v378 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v387) {
		v916 = v61
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v399 = v391
	v405 = v387
	goto L15
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v378))) = uint8(v373)
	v391 = v371 + int32(1)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if base.Ui32((v392-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v371 = v391
		v373 = v392
		v378 = v387
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v423 = l4 + v33
	v424 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v424
	v427 = v18 + v424
	if base.Ui32(v16) <= base.Ui32(v427) {
		v916 = v61
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	switch v63 - int32(43) {
	case 0, 2:
		goto L147
	default:
		goto L12
	}
L84:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32((v429-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v438)
	v441 = v39 + int32(1)
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	if base.Ui32((v442|int32(32)-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v438 = v429 | int32(32)
	goto L88
L87:
	;
	v438 = v429
	goto L88
L88:
	;
	goto L85
L89:
	;
	v451 = v441
	v453 = v442
	v457 = v427
	goto L92
L90:
	;
	v492 = v441
	v494 = v442
	v498 = v427
	goto L91
L91:
	;
	switch v494&int32(255) - int32(43) {
	case 0:
		goto L101
	default:
		goto L102
	case 2, 3, 4:
		goto L100
	}
L92:
	;
	v467 = v457 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v467) {
		v916 = v61
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v492 = v482
	v494 = v483
	v498 = v467
	goto L91
L94:
	;
	v469 = int32(255)
	v470 = v453 & v469
	if base.Ui32((v470-int32(65))&v469) < base.Ui32(int32(26)) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v479)
	v482 = v451 + int32(1)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if base.Ui32((v483|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v451 = v482
		v453 = v483
		v457 = v467
		goto L92
	} else {
		goto L99
	}
L96:
	;
	v479 = v470 | int32(32)
	goto L98
L97:
	;
	v479 = v470
	goto L98
L98:
	;
	goto L95
L99:
	;
	goto L93
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = int32(2)
	v653 = v498 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v653) {
		v916 = v61
		goto L1
	} else {
		goto L135
	}
L101:
	;
	v517 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v498))) = uint8(v517)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	goto L105
L102:
	;
	if base.Ui32(int32(9)) < base.Ui32((v494-int32(48))&int32(255)) {
		v866 = v492
		v872 = v498
		goto L13
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	if v634 != 0 {
		v866 = v492
		v872 = v498
		goto L13
	} else {
		goto L134
	}
L105:
	;
	v524 = int32(*(*int8)(unsafe.Add(mBase, uint32(v519))))
	v534 = int32(1617440)
	v535 = int32(1618576)
	goto L108
L107:
	;
	v634 = int32(0)
	goto L104
L108:
	;
	v545 = v534 + (v535-v534)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v546 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545))))
	v547 = v524 - v546
	if v547 != 0 {
		v593 = v547
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L107
L110:
	;
	v597 = base.B2i32(v593 < int32(0))
	if v593 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L111:
	;
	goto L114
L112:
	;
	if v586 != 0 {
		v593 = v586
		goto L110
	} else {
		goto L126
	}
L114:
	;
	goto L115
L115:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v554 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v555 = v519
	v556 = v545
	v557 = int32(10)
	v558 = v554
	goto L120
L117:
	;
	v580 = v545
	v584 = int32(0)
	goto L118
L118:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
	v586 = v584 - v585
	goto L112
L119:
	;
	v580 = v575
	v584 = v577
	goto L118
L120:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if v558 != v560 {
		v575 = v556
		v577 = v558
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v575 = v569
	v577 = int32(0)
	goto L119
L122:
	;
	if v560 == int32(0) {
		v575 = v556
		v577 = v558
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v565 = v557 - int32(1)
	if v565 == int32(0) {
		v575 = v556
		v577 = v558
		goto L119
	} else {
		goto L124
	}
L124:
	;
	v568 = int32(1)
	v569 = v556 + v568
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+1)))
	if v570 != 0 {
		v555 = v555 + v568
		v556 = v569
		v557 = v565
		v558 = v570
		goto L120
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	v634 = v545
	goto L104
L127:
	;
	v598 = v545 - int32(16)
	goto L129
L128:
	;
	v598 = v535
	goto L129
L129:
	;
	if v593 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v601 = v534
	goto L132
L131:
	;
	v601 = v545 + int32(16)
	goto L132
L132:
	;
	if base.Ui32(v601) <= base.Ui32(v598) {
		v534 = v601
		v535 = v598
		goto L108
	} else {
		goto L133
	}
L133:
	;
	goto L109
L134:
	;
	goto L100
L135:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	__phi656 = v492
	__phi657 = v655
	__phi658 = v653
	__phi662 = v498
	v656 = __phi656
	v657 = __phi657
	v658 = __phi658
	v662 = __phi662
	goto L136
L136:
	;
	v671 = int32(255)
	v672 = v657 & v671
	if base.Ui32((v672-int32(65))&v671) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v916 = v61
	goto L1
L138:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v662))) = uint8(v681)
	v684 = v656 + int32(1)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684))))
	switch v685 - int32(43) {
	case 0, 2, 3, 4, 15, 52:
		goto L142
	default:
		goto L143
	}
L139:
	;
	v681 = v672 | int32(32)
	goto L141
L140:
	;
	v681 = v672
	goto L141
L141:
	;
	goto L138
L142:
	;
	v700 = v658 + int32(1)
	if base.Ui32(v700) < base.Ui32(v16) {
		__phi656 = v684
		__phi657 = v685
		__phi658 = v700
		__phi662 = v658
		v656 = __phi656
		v657 = __phi657
		v658 = __phi658
		v662 = __phi662
		goto L136
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	if base.B2i32(base.Ui32(v685-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v685|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	v866 = v684
	v872 = v658
	goto L13
L146:
	;
	goto L137
L147:
	;
	v705 = v18 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v705) {
		v916 = v61
		goto L1
	} else {
		goto L148
	}
L148:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v63)
	v710 = v39
	goto L149
L149:
	;
	v724 = v710 + int32(1)
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	if base.Ui32(v725-int32(9)) < base.Ui32(int32(5)) {
		v710 = v724
		goto L149
	} else {
		goto L151
	}
L150:
	;
	if base.Ui32((v725-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	if v725 == int32(32) {
		v710 = v724
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(4)
	v742 = v18 + int32(2)
	if base.Ui32(v16) <= base.Ui32(v742) {
		v916 = v61
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if base.Ui32(int32(25)) < base.Ui32((v725|int32(32)-int32(97))&int32(255)) {
		v916 = v61
		goto L1
	} else {
		goto L165
	}
L156:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v705))) = uint8(v744)
	v748 = v710 + int32(2)
	v754 = v742
	goto L157
L157:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748))))
	if base.Ui32(int32(10)) <= base.Ui32((v763-int32(48))&int32(255)) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v773 = v763&int32(255) - int32(45)
	if base.Ui32(int32(13)) < base.Ui32(v773) {
		v866 = v748
		v872 = v754
		goto L13
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v784 = v754 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v784) {
		v916 = v61
		goto L1
	} else {
		goto L164
	}
L162:
	;
	if int32(1)<<(uint(v773)%32)&int32(8195) == int32(0) {
		v866 = v748
		v872 = v754
		goto L13
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v754))) = uint8(v763)
	v748 = v748 + int32(1)
	v754 = v784
	goto L157
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v33))) = int32(6)
	v801 = v18 + int32(2)
	if base.Ui32(v16) <= base.Ui32(v801) {
		v916 = v61
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+1)))
	if base.Ui32((v803-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v705))) = uint8(v812)
	v815 = v710 + int32(2)
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	if base.Ui32(int32(25)) < base.Ui32((v816|int32(32)-int32(97))&int32(255)) {
		v866 = v815
		v872 = v801
		goto L13
	} else {
		goto L171
	}
L168:
	;
	v812 = v803 | int32(32)
	goto L170
L169:
	;
	v812 = v803
	goto L170
L170:
	;
	goto L167
L171:
	;
	v825 = v815
	v826 = v816
	v827 = v801
	goto L172
L172:
	;
	v841 = v827 + int32(1)
	if base.Ui32(v16) <= base.Ui32(v841) {
		v916 = v61
		goto L1
	} else {
		goto L174
	}
L173:
	;
	v866 = v856
	v872 = v841
	goto L13
L174:
	;
	v843 = int32(255)
	v844 = v826 & v843
	if base.Ui32((v844-int32(65))&v843) < base.Ui32(int32(26)) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v827))) = uint8(v853)
	v856 = v825 + int32(1)
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	if base.Ui32((v857|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v825 = v856
		v826 = v857
		v827 = v841
		goto L172
	} else {
		goto L179
	}
L176:
	;
	v853 = v844 | int32(32)
	goto L178
L177:
	;
	v853 = v844
	goto L178
L178:
	;
	goto L175
L179:
	;
	goto L173
L180:
	;
	if v895 == int32(0) {
		v916 = v61
		goto L1
	} else {
		goto L184
	}
L181:
	;
	v892 = F_isalnum(m, v63)
	mBase = m.M
	v895 = base.B2i32(v892 == int32(0))
	goto L183
L182:
	;
	v895 = int32(0)
	goto L183
L183:
	;
	goto L180
L184:
	;
	goto L7
}
func F_date_cmp_timestamp_internal(m *base.Module, l0 int32, l1 int64) int32 {
	var v5 int64
	_ = v5
	var v12 int64
	_ = v12
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	if l0 == int32(-2147483648) {
		v5 = int64(-9223372036854775807 - 1)
		return base.B2i32(l1 < v5) - base.B2i32(v5 < l1)
	} else {
		if l0 == int32(2147483647) {
			v12 = int64(9223372036854775807)
			return base.B2i32(l1 < v12) - base.B2i32(v12 < l1)
		} else {
			if l0 <= int32(106751982) {
				v21 = base.I64_extend_i32_s(l0) * int64(86400000000)
				return base.B2i32(l1 < v21) - base.B2i32(v21 < l1)
			} else {
				if l1 == int64(9223372036854775807) {
					v30 = int32(-1)
				} else {
					v30 = int32(1)
				}
				return v30
			}
		}
	}
}
func F_date_eq_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[450]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(v65 == int32(0))
}
func F_date_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 <= v3)
}
func F_date_lt_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[450]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return int32(base.Ui32(v65) >> (uint(int32(31)) % 32))
}
func F_date_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 != v3)
}
func F_date_pli(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v4-int32(2147483647)) < base.Ui32(int32(2)) {
		v37 = v4
		return v37
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = v9 + v4
		if int32(0) <= v9 {
			if v4 <= v10 {
				if base.Ui32(v10+int32(2451545)) < base.Ui32(int32(2147483494)) {
					v37 = v10
					return v37
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(383478), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(475541), int32(585), int32(304542))
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(383478), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(475541), int32(585), int32(304542))
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
			if v4 < v10 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(383478), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(475541), int32(585), int32(304542))
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
			} else {
				if base.Ui32(v10+int32(2451545)) < base.Ui32(int32(2147483494)) {
					v37 = v10
					return v37
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(383478), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(475541), int32(585), int32(304542))
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
			}
		}
	}
}
func F_extract_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v505 int32
	_ = v505
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v571 int32
	_ = v571
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v654 int64
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = v22
	goto L5
L4:
	;
	v28 = v17 + int32(4)
	goto L5
L5:
	;
	if v25 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v59 = F_downcase_truncate_identifier(m, v28, v56, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v31 = int32(4)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v33&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v46 = int32(1)
	if v27 != 0 {
		v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v42 = v31
	goto L12
L11:
	;
	v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
	goto L12
L12:
	;
	if v33 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = v31
	goto L15
L14:
	;
	v45 = v42
	goto L15
L15:
	;
	v56 = v45
	goto L6
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v62 = v12 + int32(76)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
	if v69 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v130 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = v113
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113)+11)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v121
	v130 = v120
	goto L18
L20:
	;
	v71 = F_strncmp(m, v59, v69, int32(10))
	mBase = m.M
	if v71 == int32(0) {
		v113 = v69
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v80 = int32(1618592)
	v82 = int32(1619552)
	goto L24
L23:
	;
	goto L22
L24:
	;
	v89 = v80 + (v82-v80)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v89))))
	v91 = v74 - v90
	if v91 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(0)
	v130 = int32(31)
	goto L18
L26:
	;
	v95 = F_strncmp(m, v59, v89, int32(10))
	mBase = m.M
	if v95 == int32(0) {
		v113 = v89
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v98 = v91
	goto L28
L28:
	;
	v102 = base.B2i32(v98 < int32(0))
	if v98 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v98 = v95
	goto L28
L30:
	;
	v103 = v89 - int32(16)
	goto L32
L31:
	;
	v103 = v82
	goto L32
L32:
	;
	if v98 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = v80
	goto L35
L34:
	;
	v106 = v89 + int32(16)
	goto L35
L35:
	;
	if base.Ui32(v106) <= base.Ui32(v103) {
		v80 = v106
		v82 = v103
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v134 = v12 + int32(76)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[1247]))
	if v141 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v203 = v130
	goto L39
L39:
	;
	if base.Ui32(v57-int32(2147483647)) <= base.Ui32(int32(1)) {
		goto L63
	} else {
		goto L64
	}
L40:
	;
	v203 = v202
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1247])) = v185
	v192 = int32(*(*int8)(unsafe.Add(mBase, uint32(v185)+11)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v193
	v202 = v192
	goto L40
L42:
	;
	v143 = F_strncmp(m, v59, v141, int32(10))
	mBase = m.M
	if v143 == int32(0) {
		v185 = v141
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v59))))
	v152 = int32(1617440)
	v154 = int32(1618576)
	goto L46
L45:
	;
	goto L44
L46:
	;
	v161 = v152 + (v154-v152)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v161))))
	v163 = v146 - v162
	if v163 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(0)
	v202 = int32(31)
	goto L40
L48:
	;
	v167 = F_strncmp(m, v59, v161, int32(10))
	mBase = m.M
	if v167 == int32(0) {
		v185 = v161
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v170 = v163
	goto L50
L50:
	;
	v174 = base.B2i32(v170 < int32(0))
	if v170 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v170 = v167
	goto L50
L52:
	;
	v175 = v161 - int32(16)
	goto L54
L53:
	;
	v175 = v154
	goto L54
L54:
	;
	if v170 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v178 = v152
	goto L57
L56:
	;
	v178 = v161 + int32(16)
	goto L57
L57:
	;
	if base.Ui32(v178) <= base.Ui32(v175) {
		v152 = v178
		v154 = v175
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	m.G0 = v12 + int32(80)
	return v660
L60:
	;
	v655 = F_int64_to_numeric(m, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L176
	}
L61:
	;
	v651 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+64)))
	v654 = v651
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L171
	}
L63:
	;
	switch v203 {
	case 0, 17:
		goto L66
	default:
		goto L62
	}
L64:
	;
	goto L65
L65:
	;
	switch v203 {
	case 0:
		goto L85
	default:
		goto L62
	case 17:
		goto L86
	}
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v210 = v208 - int32(11)
	if base.Ui32(int32(26)) < base.Ui32(v210) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L80
	}
L68:
	;
	v214 = int32(1) << (uint(v210) % 32)
	if v214&int32(34848769) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v214&int32(73415680) == int32(0) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v57 == int32(-2147483648) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v223)
	v660 = int32(0)
	goto L59
L73:
	;
	v229 = int32(0)
	v233 = F_DirectFunctionCall3Coll(m, int32(408), v229, int32(10332), v229, int32(-1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v238 = int32(0)
	v242 = F_DirectFunctionCall3Coll(m, int32(408), v238, int32(10343), v238, int32(-1))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	v235 = F_pg_detoast_datum(m, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v660 = v235
	goto L59
L78:
	;
	v244 = F_pg_detoast_datum(m, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v660 = v244
	goto L59
L80:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v255 = F_format_type_be(m, int32(1082))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v59
	F_errmsg(m, int32(179925), v12+int32(16))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(475541), int32(1188), int32(339804))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	if v606 == int32(11) {
		v654 = base.I64_extend_i32_s(v57)*int64(86400) + int64(946684800)
		goto L60
	} else {
		goto L165
	}
L86:
	;
	v270 = v57 + int32(2451545)
	v280 = v57 + int32(2483589)
	v281 = int32(146097)
	v282 = base.I32_div_u_s(v280, v281)
	v283 = int32(3)
	v289 = int32(2)
	v294 = base.I32_div_u_s((v282*int32(1073595727)+v280)<<(uint(v289)%32)|v283, v281)
	v297 = v270 + v282*v283 + v294 + int32(32104)
	v298 = int32(1461)
	v299 = base.I32_div_u_s(v297, v298)
	v302 = v299*int32(-1461) + v297
	v304 = v302 << (uint(v289) % 32)
	if base.Ui32(v298) <= base.Ui32(v304) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	switch v343 - int32(21) {
	case 0:
		goto L61
	case 1:
		goto L101
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L100
	case 5:
		goto L99
	case 6:
		goto L98
	case 7:
		goto L97
	default:
		goto L92
	case 10:
		goto L96
	case 11, 16:
		goto L94
	case 12:
		goto L93
	case 15:
		goto L95
	}
L88:
	;
	v317 = base.I32_div_u_s(v304, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v317 + v299<<(uint(int32(2))%32) - int32(4800)
	v325 = v315 + int32(123)
	v329 = int32(base.Ui32(v325*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v325 - int32(base.Ui32(v329*int32(7834))>>(uint(int32(8))%32))
	v339 = base.I32_rem_u_s(v329+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(68)))) = v339 + int32(1)
	goto L87
L89:
	;
	v310 = base.I32_rem_u_s(v302+int32(305), int32(365))
	v315 = v310
	goto L88
L90:
	;
	goto L91
L91:
	;
	v314 = base.I32_rem_u_s(v302+int32(306), int32(366))
	v315 = v314
	goto L88
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L160
	}
L93:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v516 = base.B2i32(int32(2) < v510)
	if int32(2) < v510 {
		goto L147
	} else {
		goto L148
	}
L94:
	;
	v495 = int32(7)
	v496 = base.I32_rem_s(v57+int32(2451546), v495)
	if v496 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L95:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v450 = F_date2j(m, v446, v447, v448)
	mBase = m.M
	v451 = int32(1)
	v453 = F_date2j(m, v446, v451, int32(4))
	mBase = m.M
	v456 = F_j2day(m, v453-v451)
	mBase = m.M
	if v450 < v453-v456 {
		goto L127
	} else {
		goto L128
	}
L96:
	;
	v654 = base.I64_extend_i32_s(v270)
	goto L60
L97:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if int32(0) < v432 {
		goto L123
	} else {
		goto L124
	}
L98:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if int32(0) < v419 {
		goto L120
	} else {
		goto L121
	}
L99:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if int32(0) <= v408 {
		goto L117
	} else {
		goto L118
	}
L100:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if int32(0) < v401 {
		goto L114
	} else {
		goto L115
	}
L101:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v359 = F_date2j(m, v355, v356, v357)
	mBase = m.M
	v360 = int32(1)
	v362 = F_date2j(m, v355, v360, int32(4))
	mBase = m.M
	v365 = F_j2day(m, v362-v360)
	mBase = m.M
	if v359 < v362-v365 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v348 = int32(1)
	v351 = base.I32_div_s(v347-v348, int32(3))
	v654 = base.I64_extend_i32_s(v351 + v348)
	goto L60
L103:
	;
	v346 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+68)))
	v654 = v346
	goto L60
L104:
	;
	v654 = base.I64_extend_i32_s(v397 + int32(1))
	goto L60
L105:
	;
	v368 = int32(1)
	v372 = F_date2j(m, v355-v368, v368, int32(4))
	mBase = m.M
	v375 = F_j2day(m, v372-v368)
	mBase = m.M
	v376 = v372
	v377 = v375
	goto L107
L106:
	;
	v376 = v362
	v377 = v365
	goto L107
L107:
	;
	v379 = v377 - v376 + v359
	if int32(357) <= v379 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v382 = int32(1)
	v386 = F_date2j(m, v355+v382, v382, int32(4))
	mBase = m.M
	v389 = F_j2day(m, v386-v382)
	mBase = m.M
	v390 = v386 - v389
	if v359 < v390 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v395 = v379
	goto L110
L110:
	;
	v397 = base.I32_div_s(v395, int32(7))
	goto L104
L111:
	;
	v393 = v379
	goto L113
L112:
	;
	v393 = v359 - v390
	goto L113
L113:
	;
	v395 = v393
	goto L110
L114:
	;
	v654 = base.I64_extend_i32_u(v401)
	goto L60
L115:
	;
	goto L116
L116:
	;
	v654 = base.I64_extend_i32_s(v401 - int32(1))
	goto L60
L117:
	;
	v412 = base.I32_div_u_s(v408, int32(10))
	v654 = base.I64_extend_i32_u(v412)
	goto L60
L118:
	;
	goto L119
L119:
	;
	v417 = base.I32_div_s(int32(9)-v408, int32(-10))
	v654 = base.I64_extend_i32_s(v417)
	goto L60
L120:
	;
	v425 = base.I32_div_s(v419+int32(99), int32(100))
	v654 = base.I64_extend_i32_s(v425)
	goto L60
L121:
	;
	goto L122
L122:
	;
	v430 = base.I32_div_s(int32(100)-v419, int32(-100))
	v654 = base.I64_extend_i32_s(v430)
	goto L60
L123:
	;
	v438 = base.I32_div_s(v432+int32(999), int32(1000))
	v654 = base.I64_extend_i32_s(v438)
	goto L60
L124:
	;
	goto L125
L125:
	;
	v443 = base.I32_div_s(int32(1000)-v432, int32(-1000))
	v654 = base.I64_extend_i32_s(v443)
	goto L60
L126:
	;
	v654 = base.I64_extend_i32_s(v487) - base.I64_extend_i32_u(base.B2i32(v487 <= int32(0)))
	goto L60
L127:
	;
	v459 = int32(1)
	v460 = v446 - v459
	v463 = F_date2j(m, v460, v459, int32(4))
	mBase = m.M
	v466 = F_j2day(m, v463-v459)
	mBase = m.M
	v467 = v460
	v468 = v463
	v469 = v466
	goto L129
L128:
	;
	v467 = v446
	v468 = v453
	v469 = v456
	goto L129
L129:
	;
	if int32(357) <= v450+v469-v468 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v474 = int32(1)
	v475 = v467 + v474
	v478 = F_date2j(m, v475, v474, int32(4))
	mBase = m.M
	v481 = F_j2day(m, v478-v474)
	mBase = m.M
	if v450 < v478-v481 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v487 = v467
	goto L132
L132:
	;
	goto L126
L133:
	;
	v484 = v467
	goto L135
L134:
	;
	v484 = v475
	goto L135
L135:
	;
	v487 = v484
	goto L132
L136:
	;
	v502 = base.I64_extend_i32_s(v501)
	if v501 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v501 = v496 + v495
	goto L139
L138:
	;
	v501 = v496
	goto L139
L139:
	;
	goto L136
L140:
	;
	v504 = v502
	goto L142
L141:
	;
	v504 = int64(7)
	goto L142
L142:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	if v505 == int32(37) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v508 = v504
	goto L145
L144:
	;
	v508 = v502
	goto L145
L145:
	;
	v654 = v508
	goto L60
L146:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	goto L155
L147:
	;
	v517 = int32(4800)
	goto L149
L148:
	;
	v517 = int32(4799)
	goto L149
L149:
	;
	v518 = v517 + v509
	v523 = base.I32_div_s(v518, int32(4))
	v526 = base.I32_div_s(v518, int32(-100))
	v529 = base.I32_div_s(v518, int32(400))
	if int32(2) < v510 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v533 = int32(1)
	goto L152
L151:
	;
	v533 = int32(13)
	goto L152
L152:
	;
	v538 = base.I32_div_s((v533+v510)*int32(7834), int32(256))
	goto L146
L153:
	;
	v654 = base.I64_extend_i32_s(v511 + v518*int32(365) + v523 + v526 + v529 + v538 - int32(32167) - (int32(1) + v551*int32(365) + v556 + v559 + v562 + v571 - int32(32167)) + int32(1))
	goto L60
L155:
	;
	goto L156
L156:
	;
	v551 = int32(4799) + v542
	v556 = base.I32_div_s(v551, int32(4))
	v559 = base.I32_div_s(v551, int32(-100))
	v562 = base.I32_div_s(v551, int32(400))
	goto L158
L158:
	;
	goto L159
L159:
	;
	v571 = base.I32_div_s(int32(109676), int32(256))
	goto L153
L160:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v587 = F_format_type_be(m, int32(1082))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v59
	F_errmsg(m, int32(179925), v12+int32(32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(475541), int32(1271), int32(339804))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v617 = F_format_type_be(m, int32(1082))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v59
	F_errmsg(m, int32(179925), v12+int32(48))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(475541), int32(1287), int32(339804))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v639 = F_format_type_be(m, int32(1082))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v59
	F_errmsg(m, int32(179888), v12)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(475541), int32(1296), int32(339804))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v660 = v655
	goto L59
}
