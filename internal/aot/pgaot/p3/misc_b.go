package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BTreeTupleGetHeapTID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v4&int32(_a_F_BTreeTupleGetHeapTID_0) == int32(0) {
		return l0
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if v10&int32(_a_F_BTreeTupleGetHeapTID_0) == int32(0) {
			v15 = int32(0)
			if v10&int32(_a_F_BTreeTupleGetHeapTID_1) == v15 {
				v32 = v15
				return v32
			} else {
				return l0 + v4&int32(_a_F_BTreeTupleGetHeapTID_2) - int32(6)
			}
		} else {
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v32 = v26 + (l0 + v27<<(uint(int32(16))%32))
			return v32
		}
	}
}
func F_BTreeTupleGetHeapTIDCareful(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v13&int32(_a_F_BTreeTupleGetHeapTIDCareful_0) != 0 {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v18 = v16 & int32(_a_F_BTreeTupleGetHeapTIDCareful_0)
		if v18|base.B2i32(l2 == int32(0)) != 0 {
			if v18 != 0 {
				v49 = l2
			} else {
				v49 = int32(1)
			}
			if v49 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v120 + int32(4)
						F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_1), v10+int32(16))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3557), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
				if v18 == int32(0) {
					if v16&int32(_a_F_BTreeTupleGetHeapTIDCareful_4) == int32(0) {
						v77 = int32(0)
						v80 = int32(1)
					} else {
						v71 = l1 + v13&int32(_a_F_BTreeTupleGetHeapTIDCareful_5) - int32(6)
						v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
						v77 = v71
						v80 = base.B2i32(v74 == int32(0))
					}
				} else {
					v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
					v71 = v65 + (l1 + v66<<(uint(int32(16))%32))
					v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
					v77 = v71
					v80 = base.B2i32(v74 == int32(0))
				}
				if v80 != 0 {
					v81 = l2
				} else {
					v81 = int32(0)
				}
				if v81 == int32(0) {
					m.G0 = v10 + int32(48)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(33557032))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96 + int32(4)
							F_errmsg(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_6), v10)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3565), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v32 + int32(4)
					F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_7), v10+int32(32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3550), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
		if l2 != 0 {
			v71 = l1
			v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
			v77 = v71
			v80 = base.B2i32(v74 == int32(0))
			if v80 != 0 {
				v81 = l2
			} else {
				v81 = int32(0)
			}
			if v81 == int32(0) {
				m.G0 = v10 + int32(48)
				return v77
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v97
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96 + int32(4)
						F_errmsg(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_6), v10)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3565), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
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
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v120 + int32(4)
					F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_1), v10+int32(16))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3557), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
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
func F_BackendMain(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = m.G0
	v14 = v12 - int32(432)
	m.G0 = v14
	F_ReserveExternalFD(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[1]))
	if int32(0) < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_pg_usleep(m, v19*int32(_a_F_BackendMain_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[2])) = uint8(v27)
	v29 = int32(_a_F_BackendMain_1)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3])) = v33
	v35 = F_pq_init(m, v10)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3])) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[5])) = v35
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[6])) = int32(2)
	v44 = int32(_a_F_BackendMain_2)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v35)+276)) = v44
	v49 = int32(1132)
	v51 = m.G0
	v53 = v51 - int32(32)
	m.G0 = v53
	switch int32(1134) {
	case 0, 2:
		v63 = v49
		goto L9
	default:
		goto L10
	}
L8:
	;
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[7])) = v94
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[8])) = v94
	v104 = v94
	goto L22
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v63
	F_sigemptyset(m, v53+int32(16))
	mBase = m.M
	goto L12
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[9])) = v49
	v63 = int32(_a_F_BackendMain_3)
	goto L9
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = int32(268435456)
	v75 = v53 + int32(12)
	goto L16
L14:
	;
	m.G0 = v53 + int32(32)
	goto L8
L16:
	;
	goto L17
L17:
	;
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = int32(300)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[10])) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackendMain[11])) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackendMain[12])) = v87
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L14
L21:
	;
	F_sigprocmask(m, int32(_a_F_BackendMain_4), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	v106 = int32(40)
	v107 = v104 * v106
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[13]))) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[14]))) = v104
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[15]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[16]))) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[17]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[18]))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_BackendMain[19]))) = uint8(v108)
	v122 = v104 | int32(1)
	v124 = v122 * v106
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[13]))) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[14]))) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[15]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[16]))) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[17]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[18]))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_BackendMain[19]))) = uint8(v108)
	v139 = v104 | int32(2)
	v141 = v139 * v106
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[13]))) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[14]))) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[15]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[16]))) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[17]))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[18]))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_BackendMain[19]))) = uint8(v108)
	if v104 != int32(20) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[20])) = uint8(v177)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L21
L24:
	;
	v158 = v104 | int32(3)
	v160 = v158 * int32(40)
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[13]))) = uint8(v161)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[14]))) = v158
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[15]))) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[16]))) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[17]))) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[18]))) = v161
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_BackendMain[19]))) = uint8(v161)
	v104 = v104 + int32(4)
	goto L22
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)) = uint8(v186)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+176)) = uint8(v186)
	v190 = int32(144)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[21])))
	v203 = F_pg_getnameinfo_all(m, v35+v190, v192, v14+int32(176), int32(255), v14+v190, int32(32), v200^int32(3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v257 = v14 + int32(176)
	v258 = F_MemoryContextStrdup(m, v255, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L45
	}
L29:
	;
	if v203 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v209 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v209 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v215 = int32(_a_F_BackendMain_5)
	v217 = v203 + int32(1)
	if v217 == int32(0) {
		v237 = v215
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v237 + base.B2i32(v239 == int32(0))
	F_errmsg_internal(m, int32(_a_F_BackendMain_6), v14+int32(112))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L43
	}
L34:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	goto L33
L35:
	;
	v221 = v215
	v222 = v217
	goto L36
L36:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v223 == int32(0) {
		v237 = v221
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v237 = v233
	goto L34
L38:
	;
	v227 = v221
	goto L39
L39:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
	if v231 != 0 {
		v227 = v227 + int32(1)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v233 = v227 + int32(2)
	v235 = v222 + int32(1)
	if v235 != 0 {
		v221 = v233
		v222 = v235
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(220), int32(_a_F_BackendMain_8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+276)) = v258
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v264 = v14 + int32(144)
	v265 = F_MemoryContextStrdup(m, v262, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+292)) = v265
	v269 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[22])))
	if v269&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v203 != 0 {
		goto L59
	} else {
		goto L60
	}
L48:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	v277 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v274 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), v301, int32(_a_F_BackendMain_8))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	if v277 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v277 == int32(0) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v257
	F_errmsg(m, int32(_a_F_BackendMain_9), v14+int32(96))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v301 = int32(236)
	goto L50
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v14 + int32(176)
	F_errmsg(m, int32(_a_F_BackendMain_10), v14+int32(80))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v301 = int32(240)
	goto L50
L58:
	;
	goto L47
L59:
	;
	F_RegisterTimeout(m, int32(0), int32(1133))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L103
	}
L60:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[21])))
	if v308&int32(1) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v314 = v14 + int32(176)
	v315 = int32(_a_F_BackendMain_11)
	v319 = m.G0
	v321 = v319 - int32(32)
	v322 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v321)+24)) = v322
	*(*int64)(unsafe.Add(mBase, uint32(v321)+16)) = v322
	*(*int64)(unsafe.Add(mBase, uint32(v321)+8)) = v322
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v322
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[23])))
	if v330 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v399 = F_strlen(m, v314)
	mBase = m.M
	if base.Ui32(v399) <= base.Ui32(v398) {
		goto L59
	} else {
		goto L81
	}
L63:
	;
	v398 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[24])))
	if v334 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v338 = v314
	goto L69
L67:
	;
	goto L68
L68:
	;
	v348 = v315
	v349 = v330
	goto L72
L69:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v344 == v330 {
		v338 = v338 + int32(1)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v398 = v338 - v314
	goto L62
L71:
	;
	goto L70
L72:
	;
	v356 = v321 + int32(base.Ui32(v349)>>(uint(int32(3))%32))&int32(28)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v357 | v358<<(uint(v349)%32)
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v362 != 0 {
		v348 = v348 + v358
		v349 = v362
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v365 == int32(0) {
		v388 = v314
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	v398 = v388 - v314
	goto L62
L76:
	;
	v369 = v314
	v370 = v365
	goto L77
L77:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v321+int32(base.Ui32(v370)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v378)>>(uint(v370)%32))&int32(1) == int32(0) {
		v388 = v369
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v388 = v386
	goto L75
L79:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+1)))
	v386 = v369 + int32(1)
	if v384 != 0 {
		v369 = v386
		v370 = v384
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v401 = int32(_a_F_BackendMain_12)
	v405 = m.G0
	v407 = v405 - int32(32)
	v408 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v407)+24)) = v408
	*(*int64)(unsafe.Add(mBase, uint32(v407)+16)) = v408
	*(*int64)(unsafe.Add(mBase, uint32(v407)+8)) = v408
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v408
	v416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[25])))
	if v416 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if base.Ui32(v399) <= base.Ui32(v484) {
		goto L59
	} else {
		goto L101
	}
L83:
	;
	v484 = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[26])))
	if v420 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v424 = v314
	goto L89
L87:
	;
	goto L88
L88:
	;
	v434 = v401
	v435 = v416
	goto L92
L89:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v430 == v416 {
		v424 = v424 + int32(1)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v484 = v424 - v314
	goto L82
L91:
	;
	goto L90
L92:
	;
	v442 = v407 + int32(base.Ui32(v435)>>(uint(int32(3))%32))&int32(28)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v443 | v444<<(uint(v435)%32)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)))
	if v448 != 0 {
		v434 = v434 + v444
		v435 = v448
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v451 == int32(0) {
		v474 = v314
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v484 = v474 - v314
	goto L82
L96:
	;
	v455 = v314
	v456 = v451
	goto L97
L97:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v407+int32(base.Ui32(v456)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v464)>>(uint(v456)%32))&int32(1) == int32(0) {
		v474 = v455
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v474 = v472
	goto L95
L99:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+1)))
	v472 = v455 + int32(1)
	if v470 != 0 {
		v455 = v472
		v456 = v470
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v488 = F_MemoryContextStrdup(m, v487, v314)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+280)) = v488
	goto L59
L103:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[27]))
	F_enable_timeout_after(m, int32(0), v499*int32(1000))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L108
L106:
	;
	v529 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[28])) = uint8(v529)
	goto L113
L107:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+uint32(_c_F_BackendMain[29]))))
	v527 = v526
	goto L106
L108:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[30]))
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[31]))
	if v515 < v517 {
		goto L107
	} else {
		goto L110
	}
L109:
	;
	v527 = int32(-1)
	goto L106
L110:
	;
	v519 = F_pq_recvbuf(m)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v519 == int32(0) {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	if v527 == int32(-1) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	F_InitProcess(m)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L197
	}
L115:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L193
	}
L116:
	;
	if v527 == int32(22) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[32])))
	if v536 != int32(1) {
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v554 = int32(0)
	v556 = F_ProcessStartupPacket(m, v35, v554, v554)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L125
	}
L120:
	;
	v541 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v541 == int32(0) {
		goto L115
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_BackendMain_13), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(477), int32(_a_F_BackendMain_14))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L115
L125:
	;
	if v556 != 0 {
		goto L115
	} else {
		goto L126
	}
L126:
	;
	switch v11 - int32(1) {
	case 0:
		goto L134
	case 1:
		goto L132
	case 2:
		goto L131
	case 3:
		goto L133
	case 4:
		goto L130
	default:
		goto L129
	}
L127:
	;
	F_errdetail(m, int32(_a_F_BackendMain_15), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L190
	}
L128:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L185
	}
L129:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L158
	}
L130:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L154
	}
L131:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L150
	}
L132:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L146
	}
L133:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[33])))
	if v577 == int32(0) {
		goto L128
	} else {
		goto L139
	}
L134:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_BackendMain_16), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(317), int32(_a_F_BackendMain_8))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[34])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_BackendMain_17), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v581 == int32(1) {
		goto L127
	} else {
		goto L143
	}
L143:
	;
	F_errdetail(m, int32(_a_F_BackendMain_18), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(336), int32(_a_F_BackendMain_8))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_BackendMain_19), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(341), int32(_a_F_BackendMain_8))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
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
	F_errcode(m, int32(50463173))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_BackendMain_20), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(346), int32(_a_F_BackendMain_8))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
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
	F_errcode(m, int32(_a_F_BackendMain_21))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_BackendMain_22), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(351), int32(_a_F_BackendMain_8))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_sigprocmask(m, int32(_a_F_BackendMain_23), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v662 = v14 + int32(128)
	F_initStringInfo(m, v662)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[35])))
	if v666 == int32(1) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	goto L166
L163:
	;
	goto L164
L164:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v35)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v683
	v686 = v14 + int32(128)
	F_appendStringInfo(m, v686, int32(_a_F_BackendMain_24), v14+int32(32))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L170
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v674
	F_appendStringInfo(m, v662, int32(_a_F_BackendMain_24), v14+int32(48))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L169
	}
L166:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[36]))
	goto L168
L168:
	;
	goto L165
L169:
	;
	goto L164
L170:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v35)+360))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692))))
	if v693 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v692
	F_appendStringInfo(m, v686, int32(_a_F_BackendMain_24), v14+int32(16))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v701 = v14 + int32(128)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v35)+276))
	F_appendStringInfoString(m, v701, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v35)+292))
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v706 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v705
	F_appendStringInfo(m, v701, int32(_a_F_BackendMain_25), v14)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v14)+128))
	if v711 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L178
L180:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v14)+128))
	F_pfree(m, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L184
	}
L181:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[37]))
	v716 = F_GetBackendTypeDesc(m, v715)
	mBase = m.M
	goto L183
L182:
	;
	goto L183
L183:
	;
	goto L180
L184:
	;
	m.G0 = v14 + int32(432)
	goto L114
L185:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errmsg(m, int32(_a_F_BackendMain_26), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_errdetail(m, int32(_a_F_BackendMain_27), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(324), int32(_a_F_BackendMain_8))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = int32(64)
	F_errhint(m, int32(_a_F_BackendMain_28), v14-int32(-64))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_BackendMain_7), int32(331), int32(_a_F_BackendMain_8))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_sigprocmask(m, int32(_a_F_BackendMain_23), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3])) = v775
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[5]))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)+360))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v778)+364))
	F_PostgresMain(m, v779, v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BackgroundWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int64
	_ = v389
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int64
	_ = v492
	var v494 int64
	_ = v494
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
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
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v14 = v3
	v15 = int32(-1)
	v16 = v3
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v15 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v812 = int32(m.ExcTag)
	v813 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v812 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L7:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v556 = v14
	v558 = v16
	goto L9
L9:
	;
	if v558 != 0 {
		goto L152
	} else {
		goto L153
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v14
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v14
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[0]))
	v42 = F_MemoryContextAlloc(m, v40, int32(1460))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v14
	F_errmsg_internal(m, int32(_a_F_BackgroundWorkerMain_0), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v14
	F_errfinish(m, int32(_a_F_BackgroundWorkerMain_1), int32(725), int32(_a_F_BackgroundWorkerMain_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	base.MemoryCopy(m, v42, l0, int32(1460))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[1]))
	if v47 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	F_MemoryContextDelete(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[2])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[3])) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	if v42 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[1])) = int32(0)
	goto L19
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[4]))
	if int32(0) < v66 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[2]))
	v64 = F_GetBackendTypeDesc(m, v63)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	F_pg_usleep(m, v66*int32(_a_F_BackgroundWorkerMain_3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v80 = v74 & int32(2)
	if v80 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v81 = int32(915)
	goto L31
L30:
	;
	v81 = int32(-2)
	goto L31
L31:
	;
	v83 = m.G0
	v85 = v83 - int32(32)
	m.G0 = v85
	switch v81 + int32(2) {
	case 0, 2:
		v95 = v81
		goto L33
	default:
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	if v80 != 0 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v95
	F_sigemptyset(m, v85+int32(16))
	mBase = m.M
	goto L36
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[5])) = v81
	v95 = int32(_a_F_BackgroundWorkerMain_4)
	goto L33
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = int32(268435456)
	v107 = v85 + int32(12)
	goto L40
L38:
	;
	m.G0 = v85 + int32(32)
	goto L32
L40:
	;
	goto L41
L41:
	;
	if v107 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(40)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[6])) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[7])) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[8])) = v119
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L38
L45:
	;
	v130 = int32(917)
	goto L47
L46:
	;
	v130 = int32(-2)
	goto L47
L47:
	;
	v132 = m.G0
	v134 = v132 - int32(32)
	m.G0 = v134
	switch v130 + int32(2) {
	case 0, 2:
		v144 = v130
		goto L49
	default:
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	if v80 != 0 {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v144
	F_sigemptyset(m, v134+int32(16))
	mBase = m.M
	goto L52
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[9])) = v130
	v144 = int32(_a_F_BackgroundWorkerMain_4)
	goto L49
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = int32(268435456)
	v156 = v134 + int32(12)
	goto L56
L54:
	;
	m.G0 = v134 + int32(32)
	goto L48
L56:
	;
	goto L57
L57:
	;
	if v156 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v163 = int32(200)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[10])) = v164
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v156)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[11])) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[12])) = v168
	goto L60
L59:
	;
	goto L60
L60:
	;
	goto L54
L61:
	;
	v179 = int32(919)
	goto L63
L62:
	;
	v179 = int32(-2)
	goto L63
L63:
	;
	v181 = m.G0
	v183 = v181 - int32(32)
	m.G0 = v183
	switch v179 + int32(2) {
	case 0, 2:
		v193 = v179
		goto L65
	default:
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v226 = int32(923)
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	switch int32(925) {
	case 0, 2:
		v240 = v226
		goto L78
	default:
		goto L79
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v193
	F_sigemptyset(m, v183+int32(16))
	mBase = m.M
	goto L68
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[13])) = v179
	v193 = int32(_a_F_BackgroundWorkerMain_4)
	goto L65
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+24)) = int32(268435456)
	v205 = v183 + int32(12)
	goto L72
L70:
	;
	m.G0 = v183 + int32(32)
	goto L64
L72:
	;
	goto L73
L73:
	;
	if v205 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v212 = int32(160)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[14])) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v205)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[15])) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v205)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[16])) = v217
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L70
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v273 = int32(-2)
	v275 = m.G0
	v277 = v275 - int32(32)
	m.G0 = v277
	switch int32(0) {
	case 0, 2:
		v287 = v273
		goto L91
	default:
		goto L92
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = v240
	F_sigemptyset(m, v230+int32(16))
	mBase = m.M
	goto L81
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[17])) = v226
	v240 = int32(_a_F_BackgroundWorkerMain_4)
	goto L78
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(268435456)
	v252 = v230 + int32(12)
	goto L85
L83:
	;
	m.G0 = v230 + int32(32)
	goto L77
L85:
	;
	goto L86
L86:
	;
	if v252 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v259 = int32(300)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[18])) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[19])) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[20])) = v264
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v319 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[21])) = v319
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[22])) = v319
	v329 = v319
	goto L104
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = v287
	F_sigemptyset(m, v277+int32(16))
	mBase = m.M
	goto L94
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[23])) = v273
	v287 = int32(_a_F_BackgroundWorkerMain_4)
	goto L91
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(268435456)
	v299 = v277 + int32(12)
	goto L98
L96:
	;
	m.G0 = v277 + int32(32)
	goto L90
L98:
	;
	goto L99
L99:
	;
	if v299 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v305 = int32(20)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v299)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[24])) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v299)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[25])) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[26])) = v311
	goto L102
L101:
	;
	goto L102
L102:
	;
	goto L96
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v409 = int32(-2)
	v411 = m.G0
	v413 = v411 - int32(32)
	m.G0 = v413
	switch int32(0) {
	case 0, 2:
		v423 = v409
		goto L110
	default:
		goto L111
	}
L104:
	;
	v331 = int32(40)
	v332 = v329 * v331
	v333 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[27]))) = uint8(v333)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[28]))) = v329
	v336 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[29]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[30]))) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[31]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[32]))) = v333
	*(*uint8)(unsafe.Add(mBase, uint32(v332)+uint32(_c_F_BackgroundWorkerMain[33]))) = uint8(v333)
	v347 = v329 | int32(1)
	v349 = v347 * v331
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[27]))) = uint8(v333)
	*(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[28]))) = v347
	*(*int64)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[29]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[30]))) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[31]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[32]))) = v333
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+uint32(_c_F_BackgroundWorkerMain[33]))) = uint8(v333)
	v364 = v329 | int32(2)
	v366 = v364 * v331
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[27]))) = uint8(v333)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[28]))) = v364
	*(*int64)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[29]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[30]))) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[31]))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[32]))) = v333
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+uint32(_c_F_BackgroundWorkerMain[33]))) = uint8(v333)
	if v329 != int32(20) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v402 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[34])) = uint8(v402)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L103
L106:
	;
	v383 = v329 | int32(3)
	v385 = v383 * int32(40)
	v386 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[27]))) = uint8(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[28]))) = v383
	v389 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[29]))) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[30]))) = v386
	*(*int64)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[31]))) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[32]))) = v386
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_BackgroundWorkerMain[33]))) = uint8(v386)
	v329 = v329 + int32(4)
	goto L104
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v456 = int32(-2)
	v458 = m.G0
	v460 = v458 - int32(32)
	m.G0 = v460
	switch int32(0) {
	case 0, 2:
		v470 = v456
		goto L123
	default:
		goto L124
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+12)) = v423
	F_sigemptyset(m, v413+int32(16))
	mBase = m.M
	goto L113
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[35])) = v409
	v423 = int32(_a_F_BackgroundWorkerMain_4)
	goto L110
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+24)) = int32(268435456)
	v435 = v413 + int32(12)
	goto L117
L115:
	;
	m.G0 = v413 + int32(32)
	goto L109
L117:
	;
	goto L118
L118:
	;
	if v435 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v442 = int32(260)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v435)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[36])) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v435)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[37])) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v435)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[38])) = v447
	goto L121
L120:
	;
	goto L121
L121:
	;
	goto L115
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v503 = int32(0)
	v505 = m.G0
	v507 = v505 - int32(32)
	m.G0 = v507
	switch int32(2) {
	case 0, 2:
		v517 = v503
		goto L136
	default:
		goto L137
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v460)+12)) = v470
	F_sigemptyset(m, v460+int32(16))
	mBase = m.M
	goto L126
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[39])) = v456
	v470 = int32(_a_F_BackgroundWorkerMain_4)
	goto L123
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v460)+24)) = int32(268435456)
	v482 = v460 + int32(12)
	goto L130
L128:
	;
	m.G0 = v460 + int32(32)
	goto L122
L130:
	;
	goto L131
L131:
	;
	if v482 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v489 = int32(240)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v482)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[40])) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v482)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[41])) = v492
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[42])) = v494
	goto L134
L133:
	;
	goto L134
L134:
	;
	goto L128
L135:
	;
	goto L148
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507)+12)) = v517
	F_sigemptyset(m, v507+int32(16))
	mBase = m.M
	goto L138
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[43])) = v503
	v517 = int32(_a_F_BackgroundWorkerMain_4)
	goto L136
L138:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507)+24)) = int32(268435457)
	v529 = v507 + int32(12)
	goto L143
L141:
	;
	m.G0 = v507 + int32(32)
	goto L135
L143:
	;
	goto L144
L144:
	;
	if v529 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v536 = int32(340)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v529)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[44])) = v537
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v529)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[45])) = v539
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v529)))
	*(*int64)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[46])) = v541
	goto L147
L146:
	;
	goto L147
L147:
	;
	goto L141
L148:
	;
	v549 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v9 + int32(12)
	goto L151
L149:
	;
	v556 = v42
	v558 = int32(0)
	goto L9
L151:
	;
	goto L149
L152:
	;
	v559 = int32(_a_F_BackgroundWorkerMain_5)
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[47]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[47])) = v561 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[48])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[49])) = v9 + int32(16)
	F_InitProcess(m)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L158
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	F_EmitErrorReport(m)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	F_proc_exit(m, int32(1))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	goto L3
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	F_BaseInit(m)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	v590 = v556 + int32(1228)
	v591 = m.G0
	v593 = v591 - int32(16)
	m.G0 = v593
	v596 = v556 + int32(204)
	v597 = int32(_a_F_BackgroundWorkerMain_6)
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	v603 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[50])))
	if base.B2i32(v600 == int32(0))|base.B2i32(v600 != v603) != 0 {
		v621 = v600
		v622 = v603
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v556)+1324))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	m.T0[v783].(func(*base.Module, int32))(m, v800)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L6
	} else {
		goto L225
	}
L161:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L222
	}
L162:
	;
	m.G0 = v593 + int32(16)
	goto L160
L163:
	;
	if v621-v622 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	goto L163
L165:
	;
	v606 = v596
	v607 = v597
	goto L166
L166:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)))
	if v611 == int32(0) {
		v621 = v611
		v622 = v610
		goto L164
	} else {
		goto L168
	}
L167:
	;
	v621 = v611
	v622 = v610
	goto L164
L168:
	;
	v614 = int32(1)
	if v611 == v610 {
		v606 = v606 + v614
		v607 = v607 + v614
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v626 = int32(_a_F_BackgroundWorkerMain_7)
	v629 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[51])))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v629 == int32(0))|base.B2i32(v629 != v632) != 0 {
		v650 = v629
		v651 = v632
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	v781 = F_load_external_function(m, v596, v590, int32(1), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L6
	} else {
		goto L221
	}
L173:
	;
	if v650-v651 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L174:
	;
	goto L173
L175:
	;
	v635 = v626
	v636 = v590
	goto L176
L176:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+1)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635)+1)))
	if v640 == int32(0) {
		v650 = v640
		v651 = v639
		goto L174
	} else {
		goto L178
	}
L177:
	;
	v650 = v640
	v651 = v639
	goto L174
L178:
	;
	v643 = int32(1)
	if v640 == v639 {
		v635 = v635 + v643
		v636 = v636 + v643
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[52]))
	v783 = v656
	goto L162
L181:
	;
	goto L182
L182:
	;
	v657 = int32(_a_F_BackgroundWorkerMain_8)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[53])))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v660 == int32(0))|base.B2i32(v660 != v663) != 0 {
		v681 = v660
		v682 = v663
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v681-v682 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L184:
	;
	goto L183
L185:
	;
	v666 = v657
	v667 = v590
	goto L186
L186:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+1)))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+1)))
	if v671 == int32(0) {
		v681 = v671
		v682 = v670
		goto L184
	} else {
		goto L188
	}
L187:
	;
	v681 = v671
	v682 = v670
	goto L184
L188:
	;
	v674 = int32(1)
	if v671 == v670 {
		v666 = v666 + v674
		v667 = v667 + v674
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[54]))
	v783 = v687
	goto L162
L191:
	;
	goto L192
L192:
	;
	v688 = int32(_a_F_BackgroundWorkerMain_9)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[55])))
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v691 == int32(0))|base.B2i32(v691 != v694) != 0 {
		v712 = v691
		v713 = v694
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v712-v713 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L194:
	;
	goto L193
L195:
	;
	v697 = v688
	v698 = v590
	goto L196
L196:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698)+1)))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+1)))
	if v702 == int32(0) {
		v712 = v702
		v713 = v701
		goto L194
	} else {
		goto L198
	}
L197:
	;
	v712 = v702
	v713 = v701
	goto L194
L198:
	;
	v705 = int32(1)
	if v702 == v701 {
		v697 = v697 + v705
		v698 = v698 + v705
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[56]))
	v783 = v718
	goto L162
L201:
	;
	goto L202
L202:
	;
	v719 = int32(_a_F_BackgroundWorkerMain_10)
	v722 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[57])))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v722 == int32(0))|base.B2i32(v722 != v725) != 0 {
		v743 = v722
		v744 = v725
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v743-v744 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L204:
	;
	goto L203
L205:
	;
	v728 = v719
	v729 = v590
	goto L206
L206:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+1)))
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+1)))
	if v733 == int32(0) {
		v743 = v733
		v744 = v732
		goto L204
	} else {
		goto L208
	}
L207:
	;
	v743 = v733
	v744 = v732
	goto L204
L208:
	;
	v736 = int32(1)
	if v733 == v732 {
		v728 = v728 + v736
		v729 = v729 + v736
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[58]))
	v783 = v749
	goto L162
L211:
	;
	goto L212
L212:
	;
	v750 = int32(_a_F_BackgroundWorkerMain_11)
	v753 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[59])))
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if base.B2i32(v753 == int32(0))|base.B2i32(v753 != v756) != 0 {
		v774 = v753
		v775 = v756
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if v774-v775 != 0 {
		goto L161
	} else {
		goto L220
	}
L214:
	;
	goto L213
L215:
	;
	v759 = v750
	v760 = v590
	goto L216
L216:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760)+1)))
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+1)))
	if v764 == int32(0) {
		v774 = v764
		v775 = v763
		goto L214
	} else {
		goto L218
	}
L217:
	;
	v774 = v764
	v775 = v763
	goto L214
L218:
	;
	v767 = int32(1)
	if v764 == v763 {
		v759 = v759 + v767
		v760 = v760 + v767
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[60]))
	v783 = v778
	goto L162
L221:
	;
	v783 = v781
	goto L162
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v590
	F_errmsg_internal(m, int32(_a_F_BackgroundWorkerMain_12), v593)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L6
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_BackgroundWorkerMain_1), int32(1355), int32(_a_F_BackgroundWorkerMain_13))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v556
	F_proc_exit(m, int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	goto L5
L227:
	;
	v817 = int32(v813)
	m.G0 = v9
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v820)))
	if v9+int32(12) == v823 {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	m.ExcPending = 1
	goto L236
L229:
	;
	if v827 != 0 {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v820)+4))
	v827 = v825
	goto L232
L231:
	;
	v827 = int32(0)
	goto L232
L232:
	;
	goto L229
L233:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v14 = v828
	v15 = v827
	v16 = v819
	goto L1
L234:
	;
	goto L235
L235:
	;
	F___wasm_longjmp(m, v820, v819)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	return
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BackgroundWorkerUnblockSignals(m *base.Module) {
	var v4 int32
	_ = v4
	F_sigprocmask(m, int32(_a_F_BackgroundWorkerUnblockSignals_0), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_BeginImplicitTransactionBlock(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_BeginImplicitTransactionBlock[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	if v4 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(4)
	} else {
	}
	return
}
func F_BitvecInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_BitvecInit[0])) = int32(_a_F_BitvecInit_0)
	*(*int32)(unsafe.Add(mBase, _c_F_BitvecInit[1])) = int32(_a_F_BitvecInit_1)
	return
}
func F_BlockSampler_Next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v51 float64
	_ = v51
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 float64
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = v9 - v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = v12 - v13
	if base.Ui32(v11) < base.Ui32(v14) {
		v17 = l0 + int32(16)
		for {
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
			v30 = v28 ^ v29
			*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = base.I64_rotl(v30, int64(37))
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v30<<(uint(int64(16))%64) ^ base.I64_rotl(v28, int64(24)) ^ v30
			v51 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v28*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
			mBase = m.M
			if base.F64_eq(v51, float64(0)) != 0 {
				continue
			} else {
				break
			}
			break
		}
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v56 = base.F64_convert_i32_s(v11)
		v59 = base.F64_sub(float64(1), base.F64_div(v56, base.F64_convert_i32_u(v14)))
		if base.F64_gt(v59, v51) != 0 {
			v62 = v54
			v64 = v14
			v67 = v59
			for {
				v69 = int32(1)
				v70 = v62 + v69
				v73 = v64 - v69
				v77 = base.F64_mul(v67, base.F64_sub(float64(1), base.F64_div(v56, base.F64_convert_i32_u(v73))))
				if base.F64_lt(v51, v77) != 0 {
					v62 = v70
					v64 = v73
					v67 = v77
					continue
				} else {
					break
				}
				break
			}
			v80 = v70
		} else {
			v80 = v54
		}
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v89 = v80
		v90 = v87
	} else {
		v89 = v13
		v90 = v10
	}
	v96 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v90 + v96
	return v89
}
func F_BogusGetChunkSpace(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13821(m, l0, int32(_a_F_BogusGetChunkSpace_0), int32(315), int32(_a_F_BogusGetChunkSpace_1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BuildIndex_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v157 float64
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 float64
	_ = v209
	var v212 int32
	_ = v212
	var v214 float64
	_ = v214
	var v217 float64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 float64
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 float64
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
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
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v734 int64
	_ = v734
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v936 float64
	_ = v936
	var v940 float64
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 float64
	_ = v954
	var v955 int32
	_ = v955
	var v973 float64
	_ = v973
	var v975 int32
	_ = v975
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1044 float64
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1108 int64
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1163 int64
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1183 int32
	_ = v1183
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1204 int64
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1264 int64
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1295 int64
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	F_InitBuildState_2(m, l3, l0, l1, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v29 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v63 == int32(0) {
		v86 = v62
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v33&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(_a_F_BuildIndex_2_0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v41 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v40 + v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v44 + v41
	*(*int64)(unsafe.Add(mBase, uint32(v29+int32(80))+232)) = int64(2)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v52 + v41
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v58 - v41
	goto L4
L7:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+164))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	v96 = F_mul_size(m, v86, (v91+int32(7))&int32(-8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	v67 = F_RelationGetNumberOfBlocksInFork(m, v63, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v67 == int32(0) {
		v86 = v62
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v73 = base.I64_extend_i32_u(v67) * int64(291)
	v74 = int32(_a_F_BuildIndex_2_1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v77 = v75 * int32(50)
	if v77 <= v74 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v80 = v74
	goto L13
L12:
	;
	v80 = v77
	goto L13
L13:
	;
	v81 = base.I64_extend_i32_u(v80)
	if base.Ui64(v73) < base.Ui64(v81) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v83 = v73
	goto L16
L15:
	;
	v83 = v81
	goto L16
L16:
	;
	v86 = base.I32_wrap_i64(v83)
	goto L7
L17:
	;
	v98 = F_add_size(m, int32(20), v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v100 = F_add_size(m, v89, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+164)) = v100
	F_IvfflatCheckMemoryUsage(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	v107 = F_VectorArrayInit(m, v86, v105, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v110 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)+164))
	F_IvfflatKmeans(m, v351, v352, v353, v354, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L73
	}
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v115 = F_RelationGetNumberOfBlocksInFork(m, v110, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+144)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+136)) = int64(0)
	v122 = l3 + int32(80)
	v124 = Fn13964(m, int64(32))
	mBase = m.M
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v122)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v115
	F_pg_prng_seed(m, l3+int32(96), base.I64_extend_i32_u(v124))
	mBase = m.M
	goto L27
L26:
	;
	v142 = l3 + int32(120)
	v143 = F_pg_prng_uint32(m)
	mBase = m.M
	F_pg_prng_seed(m, v142, base.I64_extend_i32_u(v143))
	mBase = m.M
	goto L31
L27:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if base.Ui32(v159) < base.Ui32(v160) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v150 = F_pg_prng_double(m, v142)
	mBase = m.M
	if base.F64_eq(v150, float64(0)) != 0 {
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v153 = F_log(m, v150)
	mBase = m.M
	v157 = F_exp(m, base.F64_div(base.F64_neg(v153), base.F64_convert_i32_s(v113)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(l3+int32(112)))) = v157
	goto L30
L33:
	;
	goto L32
L34:
	;
	if v166 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v166 = base.B2i32(v162 < v163)
	goto L37
L36:
	;
	v166 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L41
L39:
	;
	goto L40
L40:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v299 != 0 {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v194 = v192 - v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v197 = v195 - v196
	if base.Ui32(v194) < base.Ui32(v197) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L40
L43:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v263 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v260)+188))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+140))
	v271 = m.T0[v270].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v260, v261, v262, v263, v263, v263, v247, int32(1), int32(_a_F_BuildIndex_2_2), l3, v263)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L56
	}
L44:
	;
	goto L47
L45:
	;
	v247 = v196
	v248 = v193
	goto L46
L46:
	;
	v254 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v247 + v254
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v248 + v254
	goto L43
L47:
	;
	v209 = F_pg_prng_double(m, l3+int32(96))
	mBase = m.M
	if base.F64_eq(v209, float64(0)) != 0 {
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v214 = base.F64_convert_i32_s(v194)
	v217 = base.F64_sub(float64(1), base.F64_div(v214, base.F64_convert_i32_u(v197)))
	if base.F64_gt(v217, v209) != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v220 = v212
	v222 = v197
	v225 = v217
	goto L53
L51:
	;
	v238 = v212
	goto L52
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v247 = v238
	v248 = v245
	goto L46
L53:
	;
	v227 = int32(1)
	v228 = v220 + v227
	v231 = v222 - v227
	v235 = base.F64_mul(v225, base.F64_sub(float64(1), base.F64_div(v214, base.F64_convert_i32_u(v231))))
	if base.F64_lt(v209, v235) != 0 {
		v220 = v228
		v222 = v231
		v225 = v235
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v238 = v228
	goto L52
L55:
	;
	goto L54
L56:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if base.Ui32(v273) < base.Ui32(v274) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v280 != 0 {
		goto L41
	} else {
		goto L61
	}
L58:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v280 = base.B2i32(v276 < v277)
	goto L60
L59:
	;
	v280 = int32(0)
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L42
L62:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	F_IvfflatNormVectors(m, v300, v301, v302, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v308 <= v307 {
		goto L22
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v312 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v312 == int32(0) {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_BuildIndex_2_3), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errdetail(m, int32(_a_F_BuildIndex_2_4), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errhint(m, int32(_a_F_BuildIndex_2_5), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(471), int32(_a_F_BuildIndex_2_7))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L22
L73:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	F_VectorArrayFree(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v363 = F_HnswNewBuffer(m, l1, l4)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v363
	v367 = v19 + int32(-4)
	v369 = v19 + int32(-8)
	v371 = v19 + int32(-12)
	F_IvfflatInitRegisterPage(m, l1, v367, v369, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+34)) = uint16(v362)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+32)) = uint16(v361)
	*(*int64)(unsafe.Add(mBase, uint32(v374)+24)) = int64(4316983719)
	v379 = int32(36)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+12)) = uint16(v379)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	F_IvfflatCommitBuffer(m, v381, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+12))
	v389 = F_add_size(m, int32(8), v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v394 = (v389 + int32(7)) & int32(-8)
	v395 = F_palloc0(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v397 = F_HnswNewBuffer(m, l1, l4)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v397
	F_IvfflatInitRegisterPage(m, l1, v367, v369, v371)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if int32(0) < v385 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L311
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L1
	} else {
		goto L308
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L305
	}
L85:
	;
	v406 = v395 + v394
	v408 = v395 + int32(4)
	if base.Ui32(v408) < base.Ui32(v406) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	F_IvfflatCommitBuffer(m, v560, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L130
	}
L88:
	;
	v410 = v406
	goto L90
L89:
	;
	v410 = v408
	goto L90
L90:
	;
	v417 = v395 & int32(3)
	if v417 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v418 = v394
	goto L93
L92:
	;
	v418 = (v395^int32(-1)+v410)&int32(-4) + int32(4)
	goto L93
L93:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v394) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v421 = v394
	goto L96
L95:
	;
	v421 = v418
	goto L96
L96:
	;
	v428 = int32(0)
	goto L97
L97:
	;
	v444 = int32(0)
	if base.B2i32(v417|v394 == v444)|base.B2i32(v421 == v444) == v444 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L87
L99:
	;
	base.MemoryFill(m, v395, int32(0), v421)
	goto L101
L100:
	;
	goto L101
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v395))) = int64(-1)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	if v455 <= v428 {
		goto L84
	} else {
		goto L102
	}
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v387)+16))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v387)+12))
	v460 = v457 + v458*v428
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v461 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v485 != 0 {
		goto L112
	} else {
		goto L113
	}
L104:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	if base.Ui32((v465-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v485 = int32(6)
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v477 = int32(1)
	if v461&v477 != 0 {
		v485 = int32(base.Ui32(v461) >> (uint(v477) % 32))
		goto L103
	} else {
		goto L111
	}
L107:
	;
	v472 = int32(18)
	if v465 == v472 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v476 = v472
	goto L110
L109:
	;
	v476 = int32(2)
	goto L110
L110:
	;
	v485 = v476
	goto L103
L111:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v485 = int32(base.Ui32(v481) >> (uint(int32(2)) % 32))
	goto L103
L112:
	;
	base.MemoryCopy(m, v395+int32(8), v460, v485)
	goto L114
L113:
	;
	goto L114
L114:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v488 = int32(4)
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+14)))
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+12)))
	v491 = v489 - v490
	if v491 <= v488 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if base.Ui32(v494-int32(4)) < base.Ui32(v394) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v494 = v488
	goto L118
L117:
	;
	v494 = v491
	goto L118
L118:
	;
	goto L115
L119:
	;
	F_IvfflatAppendPage(m, l1, v19+int32(-4), v19+int32(-8), v19+int32(-12), l4)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v507 = int32(0)
	v509 = F_PageAddItemExtended(m, v506, v395, v394, v507, v507)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	if v509 == int32(0) {
		goto L83
	} else {
		goto L124
	}
L124:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v513 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v536 = v533 + v428<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v536)+4)) = uint16(v509)
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v532
	v540 = v428 + int32(1)
	if v540 != v385 {
		v428 = v540
		goto L97
	} else {
		goto L129
	}
L126:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v517+(v513^int32(-1))<<(uint(int32(6))%32))+16))
	v532 = v523
	goto L125
L127:
	;
	goto L128
L128:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v525+v513<<(uint(int32(6))%32)+int32(-64))+16))
	v532 = v531
	goto L125
L129:
	;
	goto L98
L130:
	;
	F_pfree(m, v395)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v570 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v603 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L132
L134:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v574&int32(1) == int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v579 = int32(_a_F_BuildIndex_2_0)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v582 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v581 + v582
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v585 + v582
	*(*int64)(unsafe.Add(mBase, uint32(v570+int32(80))+232)) = int64(3)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v593 + v582
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v599 - v582
	goto L133
L136:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v858 != 0 {
		goto L208
	} else {
		goto L209
	}
L137:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v603)+56))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+56))
	v609 = F_plan_create_index_workers(m, v606, v608)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	if v609 <= int32(0) {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v614 = v609 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+121)))
	v618 = F_palloc0(m, int32(24))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v622)+72)) = v623 + int32(1)
	goto L141
L141:
	;
	v630 = F_CreateParallelContext(m, int32(_a_F_BuildIndex_2_8), int32(_a_F_BuildIndex_2_9), v609)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v616 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v634 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	v638 = int32(_a_F_BuildIndex_2_10)
	goto L145
L145:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v641 = F_table_parallelscan_estimate(m, v640, v638)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v636 = F_RegisterSnapshot(m, v634)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v638 = v636
	goto L145
L148:
	;
	v643 = F_add_size(m, int32(64), v641)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v630)+36))
	v650 = F_add_size(m, v645, (v643+int32(31))&int32(-32))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+36)) = v650
	v653 = F_tuplesort_estimate_shared(m, v614)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v630)+36))
	v660 = F_add_size(m, v655, (v653+int32(31))&int32(-32))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+36)) = v660
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+4))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v663)+12))
	v666 = v664 * v665
	v671 = F_add_size(m, v660, (v666+int32(31))&int32(-32))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+36)) = v671
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v630)+40))
	v676 = F_add_size(m, v674, int32(3))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+40)) = v676
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	if v680 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v630)+36))
	v682 = F_strlen(m, v680)
	mBase = m.M
	v687 = F_add_size(m, v681, v682&int32(-32)+int32(32))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v699 = int32(1)
	goto L157
L157:
	;
	F_InitializeParallelDSM(m, v630)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+36)) = v687
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v630)+40))
	v692 = F_add_size(m, v690, int32(1))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+40)) = v692
	v699 = v682 + int32(1)
	goto L157
L160:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v630)+44))
	if v702 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	switch v705 {
	case 0, 5:
		goto L165
	default:
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v718 = F_shm_toc_allocate(m, v717, v643)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L169
	}
L164:
	;
	F_DestroyParallelContext(m, v630)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	F_UnregisterSnapshot(m, v638)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v712)+72)) = v713 - int32(1)
	goto L168
L168:
	;
	goto L136
L169:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v721
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v718)+12)) = v614
	*(*uint8)(unsafe.Add(mBase, uint32(v718)+8)) = uint8(v616)
	*(*int32)(unsafe.Add(mBase, uint32(v718)+4)) = v724
	v729 = v718 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v729)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v729))) = int64(-4294967296)
	goto L170
L170:
	;
	v734 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v718)+40)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v718)+28)) = v734
	*(*int64)(unsafe.Add(mBase, uint32(v718)+48)) = v734
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v740, v718-int32(-64), v638)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v746 = F_shm_toc_allocate(m, v745, v653)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v630)+44))
	F_tuplesort_initialize_shared(m, v746, v614, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v752 = F_shm_toc_allocate(m, v751, v666)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v666 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+16))
	base.MemoryCopy(m, v752, v755, v666)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v757, int64(-6917529027641081855), v718)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v761, int64(-6917529027641081854), v746)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v765, int64(-6917529027641081853), v752)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	if v770 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v772 = F_shm_toc_allocate(m, v771, v699)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	F_LaunchParallelWorkers(m, v630)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L189
	}
L184:
	;
	if v699 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	base.MemoryCopy(m, v772, v775, v699)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v777, int64(-6917529027641081852), v772)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	goto L183
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v630
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v618)+20)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v618)+16)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v618)+12)) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v618)+8)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v618)+4)) = v785 + int32(1)
	if v785 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_WaitForParallelWorkersToFinish(m, v630)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v813 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L199
	}
L193:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v618)+16))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	switch v798 {
	case 0, 5:
		goto L195
	default:
		goto L194
	}
L194:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	F_DestroyParallelContext(m, v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	F_UnregisterSnapshot(m, v797)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v806)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v806)+72)) = v807 - int32(1)
	goto L198
L198:
	;
	goto L136
L199:
	;
	if v813 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v815
	F_errmsg(m, int32(_a_F_BuildIndex_2_11), v19+int32(-32))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+172)) = v618
	v829 = F_palloc0(m, int32(12))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(955), int32(_a_F_BuildIndex_2_12))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v829)+4)) = v831
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v829)+8)) = v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v618)+8))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v618)+12))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v618)+20))
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	v841 = base.I32_div_s(v839, v840)
	F_IvfflatParallelScanAndSort(m, v829, v835, v836, v837, v841, int32(1))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_WaitForParallelWorkersToAttach(m, v630)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	goto L136
L208:
	;
	v860 = F_palloc0(m, int32(12))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v870 = int32(0)
	goto L210
L210:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v874 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+52)) = uint16(v874)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = int32(97)
	v878 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v878
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+48)) = uint8(v878)
	v892 = F_tuplesort_begin_heap(m, v873, v874, v19+int32(-12), v19+int32(-4), v19+int32(-8), v19+int32(-16), v872, v870, v878)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	v862 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v860))) = uint8(v862)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v860)+4)) = v865
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v860)+8)) = v867
	v870 = v860
	goto L210
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+152)) = v892
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v895 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v896 != 0 {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v994 = v892
	goto L215
L215:
	;
	F_tuplesort_performsort(m, v994)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L232
	}
L216:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v973
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	v994 = v975
	goto L215
L217:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+8))
	v901 = v897 + int32(28)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	goto L220
L218:
	;
	goto L219
L219:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v945 = int32(1)
	v946 = int32(0)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v895)+188))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)+140))
	v954 = m.T0[v953].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v895, v943, v944, v945, v946, v945, v946, int32(-1), int32(_a_F_BuildIndex_2_13), l3, v946)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L231
	}
L220:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v901)))
	*(*int32)(unsafe.Add(mBase, uint32(v901))) = int32(1)
	if v921 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v936 = *(*float64)(unsafe.Add(mBase, uint32(v897)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v897)+28)) = int32(0)
	v940 = *(*float64)(unsafe.Add(mBase, uint32(v897)+40))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L230
	}
L222:
	;
	F_s_lock(m, v901, int32(_a_F_BuildIndex_2_6), int32(630), int32(_a_F_BuildIndex_2_14))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v897)+32))
	if v902 != v929 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v901))) = int32(0)
	F_ConditionVariableSleep(m, v897+int32(16), int32(134217767))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	goto L221
L229:
	;
	goto L220
L230:
	;
	v973 = v940
	goto L216
L231:
	;
	v973 = v954
	goto L216
L232:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v998 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v998
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v1003 = F_MakeTupleTableSlot(m, v1001, int32(_a_F_BuildIndex_2_15))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1010 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1044 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1048 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	goto L234
L236:
	;
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1014&int32(1) == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1019 = int32(_a_F_BuildIndex_2_0)
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1022 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1021 + v1022
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = v1025 + v1022
	*(*int64)(unsafe.Add(mBase, uint32(v1010+int32(80))+232)) = int64(4)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = v1033 + v1022
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1039 - v1022
	goto L235
L238:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1081, v1005, v1003, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L242
	}
L239:
	;
	goto L238
L240:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1052&int32(1) == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1057 = int32(_a_F_BuildIndex_2_0)
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1060 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1059 + v1060
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	*(*int32)(unsafe.Add(mBase, uint32(v1048))) = v1063 + v1060
	*(*int64)(unsafe.Add(mBase, uint32(v1048+int32(88))+232)) = base.I64_trunc_sat_f64_s(v1044)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	*(*int32)(unsafe.Add(mBase, uint32(v1048))) = v1071 + v1060
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1077 - v1060
	goto L239
L242:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	if int32(0) < v1089 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1095 = v998
	v1108 = int64(0)
	goto L246
L244:
	;
	goto L245
L245:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_tuplesort_end(m, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L1
	} else {
		goto L287
	}
L246:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[8]))
	if v1112 != 0 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L245
L248:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1115 = F_HnswNewBuffer(m, v997, l4)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L252
	}
L251:
	;
	goto L250
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v1115
	F_IvfflatInitRegisterPage(m, v997, v19+int32(-12), v19+int32(-16), v19+int32(-20))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1126 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1095 == v1146 {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1130+(v1126^int32(-1))<<(uint(int32(6))%32))+16))
	v1145 = v1136
	goto L254
L256:
	;
	goto L257
L257:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1138+v1126<<(uint(int32(6))%32)+int32(-64))+16))
	v1145 = v1144
	goto L254
L258:
	;
	v1163 = v1108
	goto L261
L259:
	;
	v1264 = v1108
	goto L260
L260:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1267 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L261:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1166)+6)))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1169 = int32(4)
	v1170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1168)+14)))
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1168)+12)))
	v1172 = v1170 - v1171
	if v1172 <= v1169 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1264 = v1204
	goto L260
L263:
	;
	v1183 = (v1167&int32(_a_F_BuildIndex_2_16) + int32(7)) & int32(_a_F_BuildIndex_2_17)
	if base.Ui32(v1175-int32(4)) < base.Ui32(v1183) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1175 = v1169
	goto L266
L265:
	;
	v1175 = v1172
	goto L266
L266:
	;
	goto L263
L267:
	;
	F_IvfflatAppendPage(m, v997, v19+int32(-12), v19+int32(-16), v19+int32(-20), l4)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1194 = int32(0)
	v1196 = F_PageAddItemExtended(m, v1193, v1166, v1183, v1194, v1194)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	if v1196 == int32(0) {
		goto L82
	} else {
		goto L272
	}
L272:
	;
	F_pfree(m, v1166)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1204 = v1163 + int64(1)
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1207 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1240, v1005, v1003, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L278
	}
L275:
	;
	goto L274
L276:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1211&int32(1) == int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1216 = int32(_a_F_BuildIndex_2_0)
	v1218 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1219 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1218 + v1219
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	*(*int32)(unsafe.Add(mBase, uint32(v1207))) = v1222 + v1219
	*(*int64)(unsafe.Add(mBase, uint32(v1207+int32(96))+232)) = v1204
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	*(*int32)(unsafe.Add(mBase, uint32(v1207))) = v1230 + v1219
	v1236 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1236 - v1219
	goto L275
L278:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1247 == v1095 {
		v1163 = v1204
		goto L261
	} else {
		goto L279
	}
L279:
	;
	goto L262
L280:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	F_IvfflatCommitBuffer(m, v1287, v1288)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L284
	}
L281:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1271+(v1267^int32(-1))<<(uint(int32(6))%32))+16))
	v1286 = v1277
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1279+v1267<<(uint(int32(6))%32)+int32(-64))+16))
	v1286 = v1285
	goto L280
L284:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v1295 = *(*int64)(unsafe.Add(mBase, uint32(v1291+v1095<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v1295
	F_IvfflatUpdateList(m, v997, v19+int32(-40), v1286, int32(-1), v1145, l4)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1303 = v1095 + int32(1)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)))
	if v1303 < v1305 {
		v1095 = v1303
		v1108 = v1264
		goto L246
	} else {
		goto L286
	}
L286:
	;
	goto L247
L287:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v1328 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	F_WaitForParallelWorkersToFinish(m, v1329)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	if l4 == int32(3) {
		goto L297
	} else {
		goto L298
	}
L291:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+16))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	switch v1333 {
	case 0, 5:
		goto L293
	default:
		goto L292
	}
L292:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	F_DestroyParallelContext(m, v1336)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	F_UnregisterSnapshot(m, v1332)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1341)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1341)+72)) = v1342 - int32(1)
	goto L296
L296:
	;
	goto L290
L297:
	;
	v1349 = int32(3)
	v1351 = F_RelationGetNumberOfBlocksInFork(m, l1, v1349)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	F_VectorArrayFree(m, v1356)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	F_log_newpage_range(m, l1, v1349, v1351, int32(1))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	F_pfree(m, v1359)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	F_MemoryContextDelete(m, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	m.G0 = v21 - int32(-64)
	return
L305:
	;
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_18), int32(0))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_19), int32(326), int32(_a_F_BuildIndex_2_20))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1385 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_21), v21)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(546), int32(_a_F_BuildIndex_2_22))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v997)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1401 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_21), v19+int32(-48))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(315), int32(_a_F_BuildIndex_2_23))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___bswap_16(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(_a_F___bswap_16_0)
}
func F_basque_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_SN_close_env(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_basque_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bernoulli_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v15 float64
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0))|base.F32_gt(v7, float32(100)) == int32(0) {
		v15 = base.F64_promote_f32(v7)
		if base.Ui64(base.I64_reinterpret_f64(v15)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v39 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v38)+12)) = uint16(v39)
			*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = l3
			*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_trunc_sat_f64_u(base.F64_nearest(base.F64_div(base.F64_mul(v15, float64(4.294967296e+09)), float64(100))))
			v50 = base.F32_ge(v7, float32(25))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v50)
			v52 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v52)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_bernoulli_beginsamplescan_0), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_bernoulli_beginsamplescan_1), int32(148), int32(_a_F_bernoulli_beginsamplescan_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
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
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_bernoulli_beginsamplescan_0), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_bernoulli_beginsamplescan_1), int32(148), int32(_a_F_bernoulli_beginsamplescan_2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
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
func F_big5_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v17, v18, v19, int32(36), int32(7))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v172)
	m.G0 = v12 + int32(16)
	return v163 - v16
L4:
	;
	v163 = v16
	v164 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v16
	v29 = v15
	v30 = v19
	goto L7
L7:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if int32(0) <= v37 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v163 = v155
	v164 = v156
	goto L3
L9:
	;
	if int32(0) < v157 {
		v28 = v155
		v29 = v156
		v30 = v157
		goto L7
	} else {
		goto L59
	}
L10:
	;
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v53 = F_pg_encoding_verifymbchar(m, int32(36), v28, v30)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v37)
	v46 = int32(1)
	v155 = v28 + v46
	v156 = v29 + v46
	v157 = v30 - v46
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v53 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v65 = (v60 | v37<<(uint(int32(8))%32)) & int32(_a_F_big5_to_mic_0)
	v67 = v12 + int32(15)
	if base.Ui32(v65) <= base.Ui32(int32(_a_F_big5_to_mic_1)) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v130 != 0 {
		goto L51
	} else {
		goto L52
	}
L25:
	;
	v129 = v127 & int32(_a_F_big5_to_mic_0)
	goto L24
L26:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v122)
	v127 = int32(63)
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v116)
	v127 = v115 | int32(-32640)
	goto L25
L28:
	;
	v110 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v110)
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+2)))
	v127 = v112 | int32(-32640)
	goto L25
L29:
	;
	v109 = int32(_a_F_big5_to_mic_2)
	goto L28
L30:
	;
	v109 = int32(_a_F_big5_to_mic_3)
	goto L28
L31:
	;
	v109 = int32(_a_F_big5_to_mic_4)
	goto L28
L32:
	;
	v109 = int32(_a_F_big5_to_mic_5)
	goto L28
L33:
	;
	v109 = int32(_a_F_big5_to_mic_6)
	goto L28
L34:
	;
	v109 = int32(_a_F_big5_to_mic_7)
	goto L28
L35:
	;
	v99 = F_BinarySearchRange(m, int32(_a_F_big5_to_mic_8), int32(46), v65)
	mBase = m.M
	if v99 == int32(0) {
		goto L26
	} else {
		goto L50
	}
L36:
	;
	v115 = v94
	v116 = int32(149)
	goto L27
L37:
	;
	switch v65 - int32(_a_F_big5_to_mic_9) {
	case 0:
		v82 = int32(_a_F_big5_to_mic_10)
		goto L40
	case 1, 3:
		goto L44
	case 2:
		goto L43
	case 4:
		goto L42
	default:
		goto L45
	}
L38:
	;
	goto L39
L39:
	;
	switch v65 - int32(_a_F_big5_to_mic_11) {
	case 0:
		v109 = int32(_a_F_big5_to_mic_12)
		goto L28
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	case 5:
		goto L30
	case 6:
		goto L29
	default:
		goto L48
	}
L40:
	;
	v83 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v83)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+2)))
	v127 = v85 | int32(-32640)
	goto L25
L41:
	;
	v82 = int32(_a_F_big5_to_mic_13)
	goto L40
L42:
	;
	v82 = int32(_a_F_big5_to_mic_14)
	goto L40
L43:
	;
	v82 = int32(_a_F_big5_to_mic_15)
	goto L40
L44:
	;
	v78 = F_BinarySearchRange(m, int32(_a_F_big5_to_mic_16), int32(23), v65)
	mBase = m.M
	if v78 != 0 {
		v94 = v78
		goto L36
	} else {
		goto L47
	}
L45:
	;
	if v65 == int32(_a_F_big5_to_mic_17) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L26
L48:
	;
	if v65 != int32(_a_F_big5_to_mic_18) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v94 = int32(_a_F_big5_to_mic_19)
	goto L36
L50:
	;
	v115 = v99
	v116 = int32(150)
	goto L27
L51:
	;
	if v130&int32(254) == int32(246) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v14 != 0 {
		v163 = v28
		v164 = v29
		goto L3
	} else {
		goto L57
	}
L54:
	;
	v135 = int32(157)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	v140 = v29 + int32(1)
	v141 = v137
	goto L56
L55:
	;
	v140 = v29
	v141 = v130
	goto L56
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)) = uint8(v129)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v141)
	v145 = int32(base.Ui32(v129) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)) = uint8(v145)
	v155 = v28 + v53
	v156 = v140 + int32(3)
	v157 = v30 - v53
	goto L9
L57:
	;
	F_report_untranslatable_char(m, int32(36), int32(7), v28, v30)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	goto L8
}
func F_bitlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
L38:
	;
	goto L37
}
func F_bitncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v8 = base.I32_div_s(l2, int32(8))
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	if v158 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	return v156
L3:
	;
	if v70 != 0 {
		v156 = v70
		goto L2
	} else {
		goto L21
	}
L4:
	;
	v70 = int32(0)
	goto L3
L5:
	;
	v44 = v39
	v45 = v40
	v46 = v41
	goto L15
L6:
	;
	if (l0|l1)&int32(3) != 0 {
		v39 = l0
		v40 = l1
		v41 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v32 = l0
	v33 = l1
	v34 = v8
	goto L8
L8:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v16 = l0
	v17 = l1
	v18 = v8
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v21 != v22 {
		v39 = v16
		v40 = v17
		v41 = v18
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v32 = v27
	v33 = v25
	v34 = v29
	goto L8
L12:
	;
	v24 = int32(4)
	v25 = v17 + v24
	v27 = v16 + v24
	v29 = v18 - v24
	if base.Ui32(int32(3)) < base.Ui32(v29) {
		v16 = v27
		v17 = v25
		v18 = v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v32
	v40 = v33
	v41 = v34
	goto L5
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 == v50 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v70 = v49 - v50
	goto L3
L17:
	;
	v52 = int32(1)
	v57 = v46 - v52
	if v57 != 0 {
		v44 = v44 + v52
		v45 = v45 + v52
		v46 = v57
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v71 = int32(0)
	v74 = l2 - v8<<(uint(int32(3))%32)
	if v74 <= v71 {
		v156 = v71
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v8))))
	v79 = int32(128)
	v80 = v78 & v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v8))))
	if v80 != v82&v79 {
		v158 = v80
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v74 == int32(1) {
		v156 = v71
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(1)
	v90 = int32(128)
	v91 = v78 << (uint(v88) % 32) & v90
	if v91 != v82<<(uint(v88)%32)&v90 {
		v158 = v91
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v74 < int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(2)
	v101 = int32(128)
	v102 = v78 << (uint(v99) % 32) & v101
	if v102 != v82<<(uint(v99)%32)&v101 {
		v158 = v102
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v74 == int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(3)
	v112 = int32(128)
	v113 = v78 << (uint(v110) % 32) & v112
	if v113 != v82<<(uint(v110)%32)&v112 {
		v158 = v113
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v74 < int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v121 = int32(4)
	v123 = int32(128)
	v124 = v78 << (uint(v121) % 32) & v123
	if v124 != v82<<(uint(v121)%32)&v123 {
		v158 = v124
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v74 == int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(5)
	v134 = int32(128)
	v135 = v78 << (uint(v132) % 32) & v134
	if v135 != v82<<(uint(v132)%32)&v134 {
		v158 = v135
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v74 < int32(7) {
		v156 = v71
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v143 = int32(6)
	v145 = int32(128)
	v146 = v78 << (uint(v143) % 32) & v145
	if v146 != v82<<(uint(v143)%32)&v145 {
		v158 = v146
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v156 = v71
	goto L2
L36:
	;
	v161 = int32(1)
	goto L38
L37:
	;
	v161 = int32(-1)
	goto L38
L38:
	;
	return v161
}
func F_bittoint4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if base.Ui32(v11) < base.Ui32(int32(33)) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if base.Ui32(int32(36)) <= base.Ui32(v14) {
				v22 = v7 + int32(8)
				v25 = v2
				for {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					v30 = v27 | v25<<(uint(int32(8))%32)
					v32 = v22 + int32(1)
					if base.Ui32(v32) < base.Ui32(v7+v16) {
						v22 = v32
						v25 = v30
						continue
					} else {
						break
					}
					break
				}
				v37 = v30
			} else {
				v37 = v2
			}
			return int32(base.Ui32(v37) >> (uint(v16<<(uint(int32(3))%32)-v11+int32(-64)) % 32))
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bittoint4_0), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bittoint4_1), int32(1596), int32(_a_F_bittoint4_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
func F_bittoint8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if base.Ui32(v11) < base.Ui32(int32(65)) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if base.Ui32(int32(36)) <= base.Ui32(v14) {
				v22 = v7 + int32(8)
				v26 = v5
				for {
					v27 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					v30 = v27 | v26<<(uint(int64(8))%64)
					v32 = v22 + int32(1)
					if base.Ui32(v32) < base.Ui32(v7+v16) {
						v22 = v32
						v26 = v30
						continue
					} else {
						break
					}
					break
				}
				v38 = v30
			} else {
				v38 = v5
			}
			v46 = F_Int64GetDatum(m, int64(base.Ui64(v38)>>(uint(base.I64_extend_i32_u(v16<<(uint(int32(3))%32)-v11+int32(-64)))%64)))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				return v46
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bittoint8_0), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bittoint8_1), int32(1676), int32(_a_F_bittoint8_2))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
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
func F_bittypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(_a_F_bittypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_booland_statefunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		return base.B2i32(v7 != int32(0))
	}
}
func F_boolin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = v12
	goto L1
L1:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v20-int32(9)))&base.B2i32(v20 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v32 = F_strlen(m, v14)
	mBase = m.M
	if v32 == int32(0) {
		v63 = v2
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = v14 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v65 = v10 + int32(15)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	switch v67 - int32(48) {
	case 0:
		goto L22
	case 1:
		goto L23
	default:
		goto L20
	case 22, 54:
		goto L27
	case 30, 62:
		goto L25
	case 31, 63:
		goto L24
	case 36, 68:
		goto L28
	case 41, 73:
		goto L26
	}
L7:
	;
	v38 = v32
	goto L8
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v38-int32(1)))))
	if base.B2i32(base.Ui32(v45-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v45 == int32(32)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v63 = v2
	goto L6
L10:
	;
	v63 = v38
	goto L6
L11:
	;
	goto L12
L12:
	;
	v56 = v38 - int32(1)
	if v56 != 0 {
		v38 = v56
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	m.G0 = v10 + int32(16)
	return v153
L15:
	;
	if v128 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	v128 = v122
	goto L15
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v118)
	v122 = v120
	goto L16
L18:
	;
	if v65 == int32(0) {
		v122 = v113
		goto L16
	} else {
		goto L45
	}
L19:
	;
	v113 = int32(1)
	goto L18
L20:
	;
	v109 = int32(0)
	if v65 != 0 {
		v118 = v109
		v120 = v109
		goto L17
	} else {
		goto L44
	}
L21:
	;
	v118 = int32(0)
	v120 = v104
	goto L17
L22:
	;
	v99 = int32(1)
	if v63 != v99 {
		goto L20
	} else {
		goto L42
	}
L23:
	;
	v96 = int32(1)
	if v63 != v96 {
		goto L20
	} else {
		goto L41
	}
L24:
	;
	v85 = int32(2)
	if base.Ui32(v63) <= base.Ui32(v85) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v81 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_0), v63)
	mBase = m.M
	if v81 != 0 {
		goto L20
	} else {
		goto L33
	}
L26:
	;
	v77 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_1), v63)
	mBase = m.M
	if v77 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L27:
	;
	v73 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_2), v63)
	mBase = m.M
	if v73 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v71 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_3), v63)
	mBase = m.M
	if v71 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	if v65 != 0 {
		v104 = int32(1)
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v128 = int32(1)
	goto L15
L32:
	;
	goto L20
L33:
	;
	if v65 != 0 {
		v104 = int32(1)
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v128 = int32(1)
	goto L15
L35:
	;
	v88 = v85
	goto L37
L36:
	;
	v88 = v63
	goto L37
L37:
	;
	v89 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_4), v88)
	mBase = m.M
	if v89 == int32(0) {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v93 = F_pg_strncasecmp(m, v14, int32(_a_F_boolin_5), v88)
	mBase = m.M
	if v93 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v65 != 0 {
		v104 = int32(1)
		goto L21
	} else {
		goto L40
	}
L40:
	;
	v128 = int32(1)
	goto L15
L41:
	;
	v113 = v96
	goto L18
L42:
	;
	if v65 != 0 {
		v104 = v99
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v128 = int32(1)
	goto L15
L44:
	;
	v122 = v109
	goto L16
L45:
	;
	v118 = v113
	v120 = int32(1)
	goto L17
L46:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	v153 = v129
	goto L14
L47:
	;
	goto L48
L48:
	;
	v130 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = F_errsave_start(m, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	if v132 == int32(0) {
		v153 = v130
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_boolin_6)
	F_errmsg(m, int32(_a_F_boolin_7), v10)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, v131, int32(_a_F_boolin_8), int32(151), int32(_a_F_boolin_9))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v153 = v130
	goto L14
}
func F_boolle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 == v3) | base.B2i32(v5 != v3)
}
func F_boolne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 != v3) ^ base.B2i32(v5 != v3)
}
func F_boolop(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 != 0 {
				v20 = F_array_contains_nulls(m, v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_boolop_0), int32(0))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_boolop_1), int32(424), int32(_a_F_boolop_2))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
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
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v25 = F_ArrayGetNItemsSafe(m, v22, v12+int32(16))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v29 != 0 {
								v37 = v29
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v37+v12, v25, v9+int32(7))
							mBase = m.M
							v42 = F__int_unique(m, v12)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								if v44 != 0 {
									v52 = v44
								} else {
									v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v53 = v42 + v52
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
								v57 = F_ArrayGetNItemsSafe(m, v45, v42+int32(16))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(_a_F_boolop_3))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v42)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v76 != v17 {
												F_pfree(m, v17)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(16)
													return v72
												}
											} else {
												m.G0 = v9 + int32(16)
												return v72
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v25 = F_ArrayGetNItemsSafe(m, v22, v12+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v29 != 0 {
						v37 = v29
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v37+v12, v25, v9+int32(7))
					mBase = m.M
					v42 = F__int_unique(m, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						if v44 != 0 {
							v52 = v44
						} else {
							v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						v53 = v42 + v52
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
						v57 = F_ArrayGetNItemsSafe(m, v45, v42+int32(16))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(_a_F_boolop_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v42)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v76 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v72
										}
									} else {
										m.G0 = v9 + int32(16)
										return v72
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
func F_booltext(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 != 0 {
		v5 = int32(_a_F_booltext_0)
	} else {
		v5 = int32(_a_F_booltext_1)
	}
	v6 = F_cstring_to_text(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_boot_yylex_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_boot_yylex_init[0])) = int32(28)
		return int32(1)
	} else {
		v11 = F_palloc(m, int32(96))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
			if v11 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_boot_yylex_init[0])) = int32(48)
				return int32(1)
			} else {
				v23 = int32(0)
				base.MemoryFill(m, v11, v23, int32(96))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v23
				v29 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v26)+52)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v26)+36)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v26)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v23
				return v23
			}
		}
	}
}
func F_bpchargt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v17
	goto L6
L5:
	;
	v26 = v12 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v53
	goto L18
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v58 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = int32(1)
	v78 = v19 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v76 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v70 = v58 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v70))))
	if v72 == int32(32) {
		v58 = v70
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = v58
	goto L20
L25:
	;
	v84 = v78
	goto L27
L26:
	;
	v84 = v19 + int32(4)
	goto L27
L27:
	;
	if v81 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v116 = v111
	goto L39
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v93 = int32(16)
	goto L34
L33:
	;
	v93 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = int32(4)
	goto L37
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v111 = v100
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v116 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(1)
	if v23&v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v134 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v128 = v116 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v128))))
	if v130 == int32(32) {
		v116 = v128
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v116
	goto L41
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v139 = int32(4)
	goto L48
L48:
	;
	v141 = int32(1)
	if v81&v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = v141
	goto L51
L50:
	;
	v145 = int32(4)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = F_varstr_cmp(m, v12+v139, v76, v19+v145, v134, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v150 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v154 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.B2i32(int32(0) < v148)
L60:
	;
	goto L59
}
func F_bpcharlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
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
	v11 = int32(1)
	v12 = v7 + v11
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v17 = v15 & v11
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = v12
	goto L5
L4:
	;
	v18 = v7 + int32(4)
	goto L5
L5:
	;
	if v15 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v49 = v45
	goto L17
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v35 = int32(1)
	if v17 != 0 {
		v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v27 = int32(16)
	goto L12
L11:
	;
	v27 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = int32(4)
	goto L15
L14:
	;
	v34 = v27
	goto L15
L15:
	;
	v45 = v34
	goto L6
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if v49 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_bpcharlen[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66*int32(28))+uint32(_c_F_bpcharlen[1])))
	goto L24
L19:
	;
	goto L18
L20:
	;
	v63 = v45 & (v45 >> (uint(int32(31)) % 32))
	goto L19
L21:
	;
	goto L22
L22:
	;
	v57 = v49 - int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v57))))
	if v59 == int32(32) {
		v49 = v57
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v63 = v49
	goto L19
L24:
	;
	if v71 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v76&v74 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v83 = v63
	goto L27
L27:
	;
	return v83
L28:
	;
	v79 = v74
	goto L30
L29:
	;
	v79 = int32(4)
	goto L30
L30:
	;
	v81 = F_pg_mbstrlen_with_len(m, v7+v79, v63)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v83 = v81
	goto L27
}
func F_bqarr_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != 0 {
			v15 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 + v14<<(uint(int32(3))%32)
			v22 = F_palloc(m, v15)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v22
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v26)
				F_infix_3(m, v7, int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v31 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							m.G0 = v7 + int32(16)
							return v35
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
						m.G0 = v7 + int32(16)
						return v35
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bqarr_out_0), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bqarr_out_1), int32(650), int32(_a_F_bqarr_out_2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
func F_brinbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(12))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_brinRevmapInitialize(m, l0, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
				v14 = F_brin_build_desc(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
					return v4
				}
			}
		}
	}
}
func F_brinendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	F_brinRevmapTerminate(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		F_MemoryContextDelete(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, v2)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_brinhandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v29 int64
	_ = v29
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = v7
		v9 = int32(5)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(4222124650660278)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+13)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(3)
		v23 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = v9
		v29 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+76)) = v29
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(6)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = int32(10)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(11)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(13)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v7
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+25)) = int32(16777217)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+17)) = int64(4311744769)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v29
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v29
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v29
		return v3
	}
}
func F_brinoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(1024), int32(12), int32(_a_F_brinoptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_btbpchar_pattern_sortsupport(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13853(m, l0, int32(1042))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_btbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = F_smgr_bulk_start_rel(m, l0, int32(3))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = F_smgr_bulk_get_buf(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = int32(0)
				F_PageInit(m, v10, int32(_a_F_btbuildempty_0), int32(16))
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+64)) = uint8(v5)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(-4616189618054758400)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v12
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(17180209506)
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
				v30 = int32(8)
				*(*uint16)(unsafe.Add(mBase, uint32(v10+v28)+12)) = uint16(v30)
				v32 = int32(72)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v32)
				F_smgr_bulk_write(m, v8, int32(0), v10, int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_smgr_bulk_finish(m, v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_btbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int64
	_ = v319
	var v325 int32
	_ = v325
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(208)
	m.G0 = v23
	v31 = int32(-1)
	v32 = v5
	v33 = v5
	v34 = v5
	v35 = v5
	v36 = v5
	v37 = v5
	v38 = v5
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v31 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L2
L5:
	;
	v407 = int32(m.ExcTag)
	v408 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v407 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L6:
	;
	if v87 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v81 = v32
	v82 = v33
	v83 = v34
	v84 = v35
	v85 = v36
	v86 = v37
	v87 = v38
	goto L6
L8:
	;
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v58 = v37
	v59 = l1
	goto L12
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v48
	v56 = F_palloc0(m, int32(40))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v48
	F_before_shmem_exit(m, int32(237), v48)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v58 = v56
	v59 = v56
	goto L12
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0]))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1]))
	goto L15
L15:
	;
	v75 = v23 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v23 + int32(12)
	goto L18
L16:
	;
	v81 = v48
	v82 = v73
	v83 = v71
	v84 = v48
	v85 = v59
	v86 = v58
	v87 = int32(0)
	goto L6
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	v92 = int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v23 + v92
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	v100 = int32(0)
	v101 = m.G0
	v103 = v101 - v92
	m.G0 = v103
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	v110 = F_LWLockAcquire(m, v106+int32(2560), v100)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0])) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	F_cancel_before_shmem_exit(m, int32(237), v81)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L61
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	F_btvacuumscan(m, l0, v85, l2, l3, v120&int32(_a_F_btbulkdelete_0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L49
	}
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v114 = int32(1)
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
	if base.Ui32(int32(_a_F_btbulkdelete_1)) < base.Ui32(v115) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v120 = v114
	goto L26
L25:
	;
	v120 = v115 + v114
	goto L26
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v120)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if int32(0) < v122 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v225+int32(2560))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L45
	}
L28:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v203+int32(2560))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L41
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
	v140 = v100
	goto L32
L30:
	;
	goto L31
L31:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	if v179 <= v122 {
		goto L27
	} else {
		goto L39
	}
L32:
	;
	v150 = v113 + int32(12) + v140*int32(12)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v127 == v151 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v84)+64))
	if v153 == v154 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v157 = v140 + int32(1)
	if v157 != v122 {
		v140 = v157
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v84)+60))
	v184 = v113 + v122*int32(12)
	*(*uint16)(unsafe.Add(mBase, uint32(v184)+20)) = uint16(v120)
	*(*int64)(unsafe.Add(mBase, uint32(v184)+12)) = v181
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v187 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v192+int32(2560))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	m.G0 = v103 + int32(16)
	goto L22
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v212 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbulkdelete_2), v103)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_btbulkdelete_3), int32(3578), int32(_a_F_btbulkdelete_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_btbulkdelete_5), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_btbulkdelete_3), int32(3586), int32(_a_F_btbulkdelete_4))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	F_cancel_before_shmem_exit(m, int32(237), v81)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0])) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	v270 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	v276 = F_LWLockAcquire(m, v272+int32(2560), v270)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v280 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v348+int32(2560))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L60
	}
L53:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
	v291 = v270
	goto L54
L54:
	;
	v308 = v279 + int32(12) + v291*int32(12)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	if v309 != v285 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L52
L56:
	;
	v325 = v291 + int32(1)
	if v325 != v280 {
		v291 = v325
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v84)+64))
	if v311 != v312 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v316 = v279 + v280*int32(12)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = v317
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v316)))
	*(*int64)(unsafe.Add(mBase, uint32(v308))) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v280 - int32(1)
	goto L52
L59:
	;
	goto L55
L60:
	;
	m.G0 = v23 + int32(208)
	return v85
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	F__bt_end_vacuum_callback(m, int32(0), v81)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v84
	F_pg_re_throw(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L4
L64:
	;
	v412 = int32(v408)
	m.G0 = v23
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	if v23+int32(12) == v418 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	m.ExcPending = 1
	goto L73
L66:
	;
	if v422 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v422 = v420
	goto L69
L68:
	;
	v422 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v23)+184))
	v31 = v422
	v32 = v426
	v33 = v427
	v34 = v428
	v35 = v423
	v36 = v425
	v37 = v424
	v38 = v414
	goto L1
L71:
	;
	goto L72
L72:
	;
	F___wasm_longjmp(m, v415, v414)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	return int32(0)
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = int64(0)
	goto L1
L1:
	;
	v21 = F__bt_first(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v62
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	return int64(0)
L5:
	;
	if v21 == int32(0) {
		v62 = v19
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v30 = l0 + int32(60)
	v33 = v19
	goto L7
L7:
	;
	F_tbm_add_tuples(m, l1, v30, int32(1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v40 = v38 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v40
	v43 = v33 + int64(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	if v44 < v40 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = F__bt_next(m, l0, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v52 = v40
	goto L12
L12:
	;
	v30 = v10 + int32(104) + v52*int32(10)
	v33 = v43
	goto L7
L13:
	;
	if v47 == int32(0) {
		v62 = v43
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v52 = v51
	goto L12
L15:
	;
	v64 = F__bt_start_prim_scan(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L2
L18:
	;
	if v64 != 0 {
		v19 = v62
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func F_btgettreeheight(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F__bt_getrootheight(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_btint82cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v6 < v5) - base.B2i32(v5 < v6)
}
func F_btint84cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v6 < v5) - base.B2i32(v5 < v6)
}
func F_btint8cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
}
func F_btint8skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(201)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(202)
	v8 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v2))) = v8
		v14 = F_Int64GetDatum(m, int64(9223372036854775807))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v2)+4)) = v14
			return int32(0)
		}
	}
}
func F_btmarkpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0])))
	if v4 != 0 {
		F_ReleaseBuffer(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int32(0)
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
			if v9 != int32(-1) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
		if v9 != int32(-1) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int64(-4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
			return
		}
	}
}
func F_btnamecmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	goto L3
L3:
	;
	v56 = F_strlen(m, v5)
	mBase = m.M
	v57 = F_strlen(m, v4)
	mBase = m.M
	v58 = F_varstr_cmp(m, v5, v56, v4, v57, v6)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	return v46 - v47
L6:
	;
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L12
L9:
	;
	v42 = v4
	v46 = int32(0)
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L4
L11:
	;
	v42 = v37
	v46 = v39
	goto L10
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v37 = v31
	v39 = int32(0)
	goto L11
L14:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	return int32(0)
L18:
	;
	return v58
}
func F_btparallelrescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	v5 = v3 + v4
	v7 = v5 + int32(12)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(-1)
		F_LWLockRelease(m, v7)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_btvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v216 int32
	_ = v216
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 != 0 {
		v216 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v216
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_ReadBuffer(m, v13, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v82 = l1
	goto L5
L5:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	v89 = v87 - v88
	v90 = m.G0
	v92 = v90 - int32(32)
	m.G0 = v92
	v95 = F_ReadBuffer(m, v86, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L31
	}
L6:
	;
	if v69 == int32(0) {
		v216 = v12
		goto L1
	} else {
		goto L28
	}
L7:
	;
	return int32(0)
L8:
	;
	F_LockBuffer(m, v15, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F__bt_checkpage(m, v13, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v15 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	if base.Ui32(v42) <= base.Ui32(int32(2)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v15^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L11
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[1]))
	v41 = v35 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	F_LockBuffer(m, v15, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	F_LockBuffer(m, v15, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	F_ReleaseBuffer(m, v15)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v69 = int32(1)
	goto L6
L20:
	;
	F_ReleaseBuffer(m, v15)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v51 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v69 = v66
	goto L6
L23:
	;
	v59 = F_RelationGetNumberOfBlocksInFork(m, v13, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v66 = int32(0)
	goto L22
L26:
	;
	v62 = base.I32_div_u_s(v59, int32(20))
	if base.Ui32(v62) < base.Ui32(v51) {
		v66 = int32(1)
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v73 = F_palloc0(m, int32(40))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v75 = int32(0)
	F_btvacuumscan(m, l0, v73, v75, v75, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)) = uint8(v80)
	v82 = v73
	goto L5
L31:
	;
	F_LockBuffer(m, v95, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F__bt_checkpage(m, v86, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	if v95 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_LockBuffer(m, v95, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L59
	}
L35:
	;
	F_LockBuffer(m, v95, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L42
	}
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if base.Ui32(v120) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L40
	}
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105+(v95^int32(-1))<<(uint(int32(2))%32))))
	v119 = v111
	goto L36
L38:
	;
	goto L39
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[1]))
	v119 = v113 + v95<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
	if v123 != v89 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	F_LockBuffer(m, v95, int32(2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v131 = int32(_a_F_btvacuumcleanup_0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2])) = v133 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if base.Ui32(v137) <= base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+64)) = uint8(v140)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = int32(3)
	v144 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+12)) = uint16(v144)
	goto L46
L45:
	;
	goto L46
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+48)) = v89
	F_MarkBufferDirty(m, v95)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v86)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+118)))
	if v152 != int32(112) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v194 = int32(_a_F_btvacuumcleanup_0)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2])) = v196 - int32(1)
	goto L34
L49:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[3]))
	if v156 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v86)+32))
	if v159 != 0 {
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v160 != 0 {
		goto L48
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	F_XLogRegisterBuffer(m, int32(0), v95, int32(14))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v119)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v119)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v119)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v175
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+64)))
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+28)) = uint8(v178)
	F_XLogRegisterBufData(m, int32(0), v92+int32(4), int32(28))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v188 = F_XLogInsert(m, int32(11), int32(224))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = base.I64_rotr(v188, int64(32))
	goto L48
L59:
	;
	F_ReleaseBuffer(m, v95)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v92 + int32(32)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v209 != 0 {
		v216 = v82
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v210 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
	if base.F64_lt(v210, v211) == int32(0) {
		v216 = v82
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v82)+8)) = v210
	v216 = v82
	goto L1
}
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v10 = F_build_attrmap_by_name(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 != v15 {
		v76 = v10
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v76
L4:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v17 < v18 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L6:
	;
	v21 = int32(20)
	v27 = v17
	goto L9
L7:
	;
	goto L8
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v63 = v58
	goto L5
L9:
	;
	v35 = v27 << (uint(int32(4)) % 32)
	v36 = l0 + v21 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != 0 {
		v76 = v10
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v63 = v40
	goto L5
L11:
	;
	v38 = int32(1)
	v39 = v27 + v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40+v27<<(uint(v38)%32)))))
	if v39 != v44 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v44 != 0 {
		v76 = v10
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v18 != v39 {
		v27 = v39
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+9)))
	if v46 != int32(1) {
		v76 = v10
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	v50 = l1 + v21 + v35
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	if v49 != v51 {
		v76 = v10
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+12)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+12)))
	if v53 != v54 {
		v76 = v10
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	goto L10
L20:
	;
	F_pfree(m, v10)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(0)
	goto L3
}
func F_build_colinfo_names_hash(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 < int32(32) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(274877907008)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_build_colinfo_names_hash[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_hash_create(m, int32(_a_F_build_colinfo_names_hash_0), v20+v11, v9, int32(1048))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v26 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = int32(0)
	v32 = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v62 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v31<<(uint(int32(2))%32))))
	if v40 == int32(0) {
		v51 = v32
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v54 = v31 + int32(1)
	if v54 < v51 {
		v31 = v54
		v32 = v51
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v43 == int32(0) {
		v51 = v32
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v48 = F_hash_search(m, v43, v40, int32(1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = v50
	goto L10
L14:
	;
	goto L9
L15:
	;
	v67 = int32(0)
	v68 = v62
	goto L18
L16:
	;
	goto L17
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v98 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v67<<(uint(int32(2))%32))))
	if v76 == int32(0) {
		v87 = v68
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v90 = v67 + int32(1)
	if v90 < v87 {
		v67 = v90
		v68 = v87
		goto L18
	} else {
		goto L24
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v79 == int32(0) {
		v87 = v68
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v84 = F_hash_search(m, v79, v76, int32(1), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = v86
	goto L20
L24:
	;
	goto L19
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v101 <= int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v106 = int32(0)
	goto L27
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L1
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v106<<(uint(int32(2))%32))))
	v119 = F_hash_search(m, v111, v116, int32(1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v122 = v106 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v122 < v123 {
		v106 = v122
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
}
func F_build_joinrel_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
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
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int64
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
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
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
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
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v622 int64
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v641 int64
	_ = v641
	var v643 int64
	_ = v643
	var v646 int64
	_ = v646
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+32)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == v7 {
		v641 = v21
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L14
	} else {
		goto L159
	}
L2:
	;
	v643 = int64(1073741823)
	if v643 <= v641 {
		goto L156
	} else {
		goto L157
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		v641 = v21
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = v7
	v43 = v21
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v641 = v622
	goto L2
L7:
	;
	v625 = v41 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v625 < v626 {
		v41 = v625
		v43 = v622
		goto L5
	} else {
		goto L154
	}
L8:
	;
	if v50 == int32(319) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v398 == int32(-4) {
		goto L104
	} else {
		goto L105
	}
L11:
	;
	v55 = F_find_placeholder_info(m, l0, v49)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L100
	}
L14:
	;
	return
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v57 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v112 == int32(0) {
		v622 = v43
		goto L7
	} else {
		goto L30
	}
L17:
	;
	v112 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v65 = int32(1)
	if v29 == int32(0) {
		v102 = v65
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v112 = v102
	goto L16
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v69 < v68 {
		v102 = v65
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(1)
	if v68 <= v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = v71
	goto L25
L24:
	;
	v74 = v68
	goto L25
L25:
	;
	v75 = int32(8)
	v80 = int32(0)
	goto L26
L26:
	;
	v87 = v80 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57+v75+v87)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v29+v75+v87)))
	v94 = v89 & (v91 ^ int32(-1))
	v96 = base.B2i32(v94 != int32(0))
	if v94 != 0 {
		v102 = v96
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v102 = v96
	goto L20
L28:
	;
	v98 = v80 + int32(1)
	if v98 != v74 {
		v80 = v98
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if l5 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L34
	}
L32:
	;
	v363 = v49
	goto L33
L33:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v378 = F_lappend(m, v377, v363)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L14
	} else {
		goto L99
	}
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v117 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if l4 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L36:
	;
	v120 = F_bms_is_member(m, v117, v29)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	if v120 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v126 = int32(0)
	if v124 == v126 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v179 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v179 = int32(1)
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v125 == int32(0) {
		v172 = v126
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v179 = v172
	goto L39
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v136 < v135 {
		v172 = v126
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v138 = int32(1)
	if v135 <= v138 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v141 = v138
	goto L48
L47:
	;
	v141 = v135
	goto L48
L48:
	;
	v142 = int32(8)
	v147 = int32(0)
	goto L49
L49:
	;
	v154 = v147 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v124+v142+v154)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v125+v142+v154)))
	v161 = v156 & (v158 ^ int32(-1))
	v163 = base.B2i32(v161 == int32(0))
	if v161 != 0 {
		v172 = v163
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v172 = v163
	goto L43
L51:
	;
	v165 = v147 + int32(1)
	if v165 != v141 {
		v147 = v165
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v182 != int32(2) {
		goto L35
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v245 = F_bms_add_member(m, v243, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L14
	} else {
		goto L72
	}
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v187 = int32(0)
	if v185 == v187 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v240 == int32(0) {
		goto L35
	} else {
		goto L71
	}
L58:
	;
	v240 = int32(1)
	goto L57
L59:
	;
	goto L60
L60:
	;
	if v186 == int32(0) {
		v233 = v187
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v240 = v233
	goto L57
L62:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v197 < v196 {
		v233 = v187
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v199 = int32(1)
	if v196 <= v199 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = v199
	goto L66
L65:
	;
	v202 = v196
	goto L66
L66:
	;
	v203 = int32(8)
	v208 = int32(0)
	goto L67
L67:
	;
	v215 = v208 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v185+v203+v215)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v186+v203+v215)))
	v222 = v217 & (v219 ^ int32(-1))
	v224 = base.B2i32(v222 == int32(0))
	if v222 != 0 {
		v233 = v224
		goto L61
	} else {
		goto L69
	}
L68:
	;
	v233 = v224
	goto L61
L69:
	;
	v226 = v208 + int32(1)
	if v226 != v202 {
		v208 = v226
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L55
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v245
	goto L35
L73:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v356 = F_bms_intersect(m, v355, v29)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L14
	} else {
		goto L97
	}
L74:
	;
	v250 = int32(0)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v251 <= v250 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v260 = v250
	goto L76
L76:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v260<<(uint(int32(2))%32))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v276 = int32(0)
	if v269 == v276 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L73
L78:
	;
	if v329 != 0 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v329 = int32(1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v275 == int32(0) {
		v322 = v276
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v329 = v322
	goto L78
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v286 < v285 {
		v322 = v276
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v288 = int32(1)
	if v285 <= v288 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v291 = v288
	goto L87
L86:
	;
	v291 = v285
	goto L87
L87:
	;
	v292 = int32(8)
	v297 = int32(0)
	goto L88
L88:
	;
	v304 = v297 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v269+v292+v304)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v275+v292+v304)))
	v311 = v306 & (v308 ^ int32(-1))
	v313 = base.B2i32(v311 == int32(0))
	if v311 != 0 {
		v322 = v313
		goto L82
	} else {
		goto L90
	}
L89:
	;
	v322 = v313
	goto L82
L90:
	;
	v315 = v297 + int32(1)
	if v315 != v291 {
		v297 = v315
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v274)+24))
	v332 = F_bms_add_member(m, v330, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L14
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v336 = v260 + int32(1)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v336 < v337 {
		v260 = v336
		goto L76
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v332
	goto L94
L96:
	;
	goto L77
L97:
	;
	v358 = F_bms_join(m, v354, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v358
	v363 = v115
	goto L33
L99:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v378
	v382 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+24)))
	v622 = v43 + v382
	goto L7
L100:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v388
	F_errmsg_internal(m, int32(_a_F_build_joinrel_tlist_0), v18)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L14
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_tlist_1), int32(1172), int32(_a_F_build_joinrel_tlist_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v491 = int64(*(*int32)(unsafe.Add(mBase, uint32(v490))))
	if l5 == int32(0) {
		v589 = v49
		goto L124
	} else {
		goto L125
	}
L104:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v402+v403<<(uint(int32(2))%32)-int32(4))))
	v490 = v409 + int32(8)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v412) <= base.Ui32(v398) {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v398<<(uint(int32(2))%32))))
	if v418 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v421 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v422 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418)+80)))
	v425 = (v421 - v422) << (uint(int32(2)) % 32)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v418)+84))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425+v426)))
	if v428 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v483 == int32(0) {
		v622 = v43
		goto L7
	} else {
		goto L123
	}
L110:
	;
	v483 = int32(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v436 = int32(1)
	if v29 == int32(0) {
		v473 = v436
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v483 = v473
	goto L109
L114:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v440 < v439 {
		v473 = v436
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v442 = int32(1)
	if v439 <= v442 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v445 = v442
	goto L118
L117:
	;
	v445 = v439
	goto L118
L118:
	;
	v446 = int32(8)
	v451 = int32(0)
	goto L119
L119:
	;
	v458 = v451 << (uint(int32(2)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v428+v446+v458)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v29+v446+v458)))
	v465 = v460 & (v462 ^ int32(-1))
	v467 = base.B2i32(v465 != int32(0))
	if v465 != 0 {
		v473 = v467
		goto L113
	} else {
		goto L121
	}
L120:
	;
	v473 = v467
	goto L113
L121:
	;
	v469 = v451 + int32(1)
	if v469 != v445 {
		v451 = v469
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v418)+88))
	v490 = v486 + v425
	goto L103
L124:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v605 = F_lappend(m, v604, v589)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L14
	} else {
		goto L153
	}
L125:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v494 == int32(-4) {
		v589 = v49
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v497 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v499 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if l4 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L129:
	;
	v502 = F_bms_is_member(m, v499, v29)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	if v502 == int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v508 = F_bms_is_member(m, v506, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	if v508 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v512 != int32(2) {
		goto L128
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v523 = F_bms_add_member(m, v521, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L14
	} else {
		goto L139
	}
L136:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v517 = F_bms_is_member(m, v515, v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	if v517 == int32(0) {
		goto L128
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v523
	goto L128
L140:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v582 = F_bms_intersect(m, v581, v29)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L14
	} else {
		goto L151
	}
L141:
	;
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v529 <= v528 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v538 = v528
	goto L143
L143:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548+v538<<(uint(int32(2))%32))))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+16))
	v554 = F_bms_is_member(m, v547, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L14
	} else {
		goto L145
	}
L144:
	;
	goto L140
L145:
	;
	if v554 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v552)+24))
	v558 = F_bms_add_member(m, v556, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L14
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v562 = v538 + int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v562 < v563 {
		v538 = v562
		goto L143
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v558
	goto L148
L150:
	;
	goto L144
L151:
	;
	v584 = F_bms_join(m, v580, v582)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L14
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v584
	v589 = v497
	goto L124
L153:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v607)+4)) = v605
	v622 = v43 + v491
	goto L7
L154:
	;
	goto L6
L155:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v648)+32)) = base.I32_wrap_i64(v646)
	m.G0 = v18 + int32(32)
	return
L156:
	;
	v646 = v643
	goto L158
L157:
	;
	v646 = v641
	goto L158
L158:
	;
	goto L155
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v398
	F_errmsg_internal(m, int32(_a_F_build_joinrel_tlist_3), v18+int32(16))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_tlist_1), int32(426), int32(_a_F_build_joinrel_tlist_4))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L14
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_paths_for_OR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v17 == v5 {
		v201 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(144)
	return v201
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v201 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = v5
	v33 = v5
	v34 = v5
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+110)))
	if v40 == int32(0) {
		v184 = v33
		v185 = v34
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v201 = v185
	goto L1
L6:
	;
	v187 = v32 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v187 < v188 {
		v32 = v187
		v33 = v184
		v34 = v185
		goto L4
	} else {
		goto L36
	}
L7:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v45 == v43 {
		v68 = v33
		v69 = v43
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v72 = int32(0)
	base.MemoryFill(m, v15+int32(12), v72, int32(132))
	if l2 == v72 {
		v110 = v43
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+100)))
	if v49 != 0 {
		v68 = v33
		v69 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v33
	v56 = v45
	goto L13
L12:
	;
	v50 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v58 = F_predicate_implied_by(m, v56, v55, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v55 = v50
	v56 = v54
	goto L13
L16:
	;
	if v58 == int32(0) {
		v184 = v55
		v185 = v34
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v64 = F_predicate_implied_by(m, v62, l3, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v68 = v55
	v69 = v64 ^ int32(1)
	goto L8
L19:
	;
	if (v110|v69)&int32(1) == int32(0) {
		v184 = v68
		v185 = v34
		goto L6
	} else {
		goto L26
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v77 <= int32(0) {
		v110 = v43
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v84 = v43
	goto L22
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v84<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v96, v39, v15+int32(12))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)))
	v110 = v105
	goto L19
L24:
	;
	v102 = v84 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v102 < v103 {
		v84 = v102
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v170 = F_build_index_paths(m, l0, l1, v39, v15+int32(12), v69, int32(1), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L34
	}
L28:
	;
	v125 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v126 <= v125 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v133 = v125
	goto L30
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v133<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v145, v39, v15+int32(12))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L32
	}
L31:
	;
	goto L27
L32:
	;
	v151 = v133 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v151 < v152 {
		v133 = v151
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v172 = F_list_concat(m, v34, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v184 = v68
	v185 = v172
	goto L6
L36:
	;
	goto L5
}
func F_build_pgstattuple_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v13 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(3024)
	m.G0 = v19
	v24 = F_get_call_result_type(m, l1, int32(0), v19+int32(140))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		if v24 == int32(1) {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
			v31 = F_TupleDescGetAttInMetadata(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				if v33 == int64(0) {
					v54 = v13
					v55 = v13
					v56 = float64(0)
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v39 = float64(100)
					v41 = base.F64_convert_i64_u(v33)
					v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					v54 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v37), v39), v41)
					v55 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v43), v39), v41)
					v56 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v48), v39), v41)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v33
				v59 = v19 + int32(2656)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3008)) = v59
				v62 = v19 + int32(2342)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3004)) = v62
				v65 = v19 + int32(2028)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3000)) = v65
				v68 = v19 + int32(1714)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2996)) = v68
				v71 = v19 + int32(1400)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2992)) = v71
				v74 = v19 + int32(1086)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2988)) = v74
				v77 = v19 + int32(772)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2984)) = v77
				v80 = v19 + int32(458)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2980)) = v80
				v83 = v19 + int32(144)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2976)) = v83
				v89 = F_pg_snprintf(m, v83, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(128))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v91
					v97 = F_pg_snprintf(m, v80, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(112))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v99
						v105 = F_pg_snprintf(m, v77, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(96))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v19)+80)) = v56
							v112 = F_pg_snprintf(m, v74, int32(314), int32(_a_F_build_pgstattuple_type_1), v19+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v114
								v120 = F_pg_snprintf(m, v71, int32(314), int32(_a_F_build_pgstattuple_type_0), v19-int32(-64))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v122 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v122
									v128 = F_pg_snprintf(m, v68, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(48))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v55
										v135 = F_pg_snprintf(m, v65, int32(314), int32(_a_F_build_pgstattuple_type_1), v19+int32(32))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v137
											v143 = F_pg_snprintf(m, v62, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(16))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												*(*float64)(unsafe.Add(mBase, uint32(v19))) = v54
												v148 = F_pg_snprintf(m, v59, int32(314), int32(_a_F_build_pgstattuple_type_1), v19)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													v152 = F_BuildTupleFromCStrings(m, v31, v19+int32(2976))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
														v155 = F_HeapTupleHeaderGetDatum(m, v154)
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															m.G0 = v19 + int32(3024)
															return v155
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v164 = m.ExcPending
			if v164 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_pgstattuple_type_2), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_build_pgstattuple_type_3), int32(109), int32(_a_F_build_pgstattuple_type_4))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
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
func F_build_sorted_items(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	v6 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = v17 * l3
	v24 = F_palloc0(m, v18*int32(5)+v17*int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v28 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v24 + v28*int32(12)
	v43 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = F_palloc(m, v84<<(uint(int32(2))%32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v56 = v24 + v43*int32(12)
	v57 = l3 * v43
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v33 + v18<<(uint(int32(2))%32) + v57
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v33 + v57<<(uint(int32(2))%32)
	v65 = v43 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v65 < v66 {
		v43 = v65
		goto L6
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	goto L7
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v89 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v139 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v110 = v98 << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112+v110)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v116 = F_get_typlen(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87+v110))) = v116
	v120 = v98 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 < v121 {
		v98 = v120
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_qsort_interruptible(m, v24, v296, int32(12), int32(1062), l2)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L46
	}
L18:
	;
	F_pfree(m, v24)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L45
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v155 = v6
	v159 = v6
	goto L22
L22:
	;
	if base.B2i32(l3 <= int32(0)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v296
	if v296 != 0 {
		goto L17
	} else {
		goto L44
	}
L24:
	;
	v304 = v159 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v304 < v305 {
		v155 = v296
		v159 = v304
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v166 = v24 + v155*int32(12)
	v178 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v296 = v155 + int32(1)
	goto L24
L28:
	;
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v185 <= v184 {
		v222 = v184
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v233 = int32(2)
	v234 = v222 << (uint(v233) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234+v235)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v159<<(uint(v233)%32))))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242+v234)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v159))))
	if v246 != 0 {
		v257 = v241
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4+v178<<(uint(int32(1))%32)))))
	v198 = v184
	goto L32
L32:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v198<<(uint(int32(1))%32)))))
	if v192 == v212 {
		v222 = v198
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v222 = v185
	goto L30
L34:
	;
	v215 = v198 + int32(1)
	if v215 != v185 {
		v198 = v215
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v258+v178<<(uint(int32(2))%32)))) = v257
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v263+v178))) = uint8(v246)
	v267 = v178 + int32(1)
	if v267 != l3 {
		v178 = v267
		goto L28
	} else {
		goto L42
	}
L37:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v234+v87)))
	if v248 != int32(-1) {
		v257 = v241
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v251 = F_toast_raw_datum_size(m, v241)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v251) {
		v296 = v155
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v255 = F_pg_detoast_datum(m, v241)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v257 = v255
	goto L36
L42:
	;
	goto L29
L43:
	;
	goto L23
L44:
	;
	goto L18
L45:
	;
	return int32(0)
L46:
	;
	return v24
}
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13857(m, l0, l1, int32(26), int32(2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_byteacmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	v88 = base.B2i32(v45 < v75)
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v151 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v162 = v151
	goto L63
L62:
	;
	v162 = base.B2i32(v75 < v45) - v88
	goto L63
L63:
	;
	return v162
}
func F_byteaeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_toast_raw_datum_size(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v109
L2:
	;
	return int32(0)
L3:
	;
	v12 = F_toast_raw_datum_size(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v8 != v12 {
		v109 = int32(0)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v15 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v17 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v19 = int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21&v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = v19
	goto L10
L9:
	;
	v24 = int32(4)
	goto L10
L10:
	;
	v25 = v15 + v24
	v26 = int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v26
	goto L13
L12:
	;
	v31 = int32(4)
	goto L13
L13:
	;
	v32 = v17 + v31
	v33 = int32(4)
	v34 = v8 - v33
	if base.Ui32(v33) <= base.Ui32(v34) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v97 != v15 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v96 = int32(0)
	goto L14
L16:
	;
	v70 = v65
	v71 = v66
	v72 = v67
	goto L26
L17:
	;
	if (v25|v32)&int32(3) != 0 {
		v65 = v25
		v66 = v32
		v67 = v34
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v58 = v25
	v59 = v32
	v60 = v34
	goto L19
L19:
	;
	if v60 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v42 = v25
	v43 = v32
	v44 = v34
	goto L21
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v47 != v48 {
		v65 = v42
		v66 = v43
		v67 = v44
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v58 = v53
	v59 = v51
	v60 = v55
	goto L19
L23:
	;
	v50 = int32(4)
	v51 = v43 + v50
	v53 = v42 + v50
	v55 = v44 - v50
	if base.Ui32(int32(3)) < base.Ui32(v55) {
		v42 = v53
		v43 = v51
		v44 = v55
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v65 = v58
	v66 = v59
	v67 = v60
	goto L16
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v75 == v76 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v96 = v75 - v76
	goto L14
L28:
	;
	v78 = int32(1)
	v83 = v72 - v78
	if v83 != 0 {
		v70 = v70 + v78
		v71 = v71 + v78
		v72 = v83
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	F_pfree(m, v15)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v102 = base.B2i32(v96 == int32(0))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 == v103 {
		v109 = v102
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v17)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v109 = v102
	goto L1
}
func F_bytealt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	v88 = base.B2i32(v45 < v75)
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v160 = int32(0)
	return base.B2i32(v151 == v160)&v88 | base.B2i32(v151 < v160)
L60:
	;
	goto L59
}
func F_byteaout(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v196 int32
	_ = v196
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v224 int64
	_ = v224
	var v228 int32
	_ = v228
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v249 int64
	_ = v249
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_byteaout[0]))
		switch v21 {
		case 0:
			v122 = int32(4)
			v124 = v16 + v122
			v125 = int32(1)
			v126 = v16 + v125
			v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v129 = v127 & v125
			if v127 == v125 {
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
				if base.Ui32((v132-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v157 = v122
					if v129 != 0 {
						v158 = v126
					} else {
						v158 = v124
					}
					if v157 == int32(1) {
						v215 = v158
						v224 = int64(1)
						v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
						if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v235 = int64(4)
						} else {
							v235 = int64(1)
						}
						if v228 == int32(92) {
							v238 = int64(2)
						} else {
							v238 = v235
						}
						v249 = v224 + v238
					} else {
						v168 = v158
						v170 = int32(0)
						v177 = int64(1)
						for {
							v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
							if base.Ui32((v181-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v188 = int64(4)
							} else {
								v188 = int64(1)
							}
							if v181 == int32(92) {
								v191 = int64(2)
							} else {
								v191 = v188
							}
							v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
							if base.Ui32((v196-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v203 = int64(4)
							} else {
								v203 = int64(1)
							}
							if v196 == int32(92) {
								v206 = int64(2)
							} else {
								v206 = v203
							}
							v207 = v177 + v191 + v206
							v208 = int32(2)
							v209 = v168 + v208
							v211 = v170 + v208
							if v211 != v157&int32(-2) {
								v168 = v209
								v170 = v211
								v177 = v207
								continue
							} else {
								break
							}
							break
						}
						if v157&int32(1) == int32(0) {
							v249 = v207
						} else {
							v215 = v209
							v224 = v207
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
							if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v235 = int64(4)
							} else {
								v235 = int64(1)
							}
							if v228 == int32(92) {
								v238 = int64(2)
							} else {
								v238 = v235
							}
							v249 = v224 + v238
						}
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v249) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v372 = m.ExcPending
							if v372 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
								mBase = m.M
								v376 = m.ExcPending
								if v376 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_byteaout_1), int32(438), int32(_a_F_byteaout_2))
									mBase = m.M
									v381 = m.ExcPending
									if v381 != 0 {
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
						v261 = v249
						v263 = F_palloc(m, base.I32_wrap_i64(v261))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v266 = int32(1)
							v267 = v265 & v266
							if v265 == v266 {
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
								if base.Ui32((v271-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v296 = int32(4)
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									if v271 == int32(18) {
										v282 = int32(16)
									} else {
										v282 = int32(0)
									}
									v294 = v282
									if v294 != 0 {
										v296 = v294
										if v267 != 0 {
											v297 = v126
										} else {
											v297 = v124
										}
										v298 = v263
										v300 = v297
										v302 = v296
										for {
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
											if v308 == int32(92) {
												v311 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
												v345 = v298 + int32(2)
											} else {
												if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v321 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
													v323 = int32(7)
													v325 = int32(48)
													v326 = v308&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
													v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
													v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
													v345 = v298 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
													v345 = v298 + int32(1)
												}
											}
											v346 = int32(1)
											v349 = v302 - v346
											if v349 != 0 {
												v298 = v345
												v300 = v300 + v346
												v302 = v349
												continue
											} else {
												break
											}
											break
										}
										v350 = v345
										v353 = v263
									} else {
										v350 = v263
										v353 = v263
									}
								}
							} else {
								v283 = int32(1)
								if v267 != 0 {
									v294 = int32(base.Ui32(v265)>>(uint(v283)%32)) - v283
								} else {
									v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v294 = int32(base.Ui32(v287)>>(uint(int32(2))%32)) - int32(4)
								}
								if v294 != 0 {
									v296 = v294
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									v350 = v263
									v353 = v263
								}
							}
							v360 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
							m.G0 = v13 + int32(16)
							return v353
						}
					}
				} else {
					if v132 == int32(18) {
						v143 = int32(16)
					} else {
						v143 = int32(0)
					}
					v154 = v143
					if v154 != 0 {
						v157 = v154
						if v129 != 0 {
							v158 = v126
						} else {
							v158 = v124
						}
						if v157 == int32(1) {
							v215 = v158
							v224 = int64(1)
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
							if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v235 = int64(4)
							} else {
								v235 = int64(1)
							}
							if v228 == int32(92) {
								v238 = int64(2)
							} else {
								v238 = v235
							}
							v249 = v224 + v238
						} else {
							v168 = v158
							v170 = int32(0)
							v177 = int64(1)
							for {
								v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
								if base.Ui32((v181-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v188 = int64(4)
								} else {
									v188 = int64(1)
								}
								if v181 == int32(92) {
									v191 = int64(2)
								} else {
									v191 = v188
								}
								v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
								if base.Ui32((v196-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v203 = int64(4)
								} else {
									v203 = int64(1)
								}
								if v196 == int32(92) {
									v206 = int64(2)
								} else {
									v206 = v203
								}
								v207 = v177 + v191 + v206
								v208 = int32(2)
								v209 = v168 + v208
								v211 = v170 + v208
								if v211 != v157&int32(-2) {
									v168 = v209
									v170 = v211
									v177 = v207
									continue
								} else {
									break
								}
								break
							}
							if v157&int32(1) == int32(0) {
								v249 = v207
							} else {
								v215 = v209
								v224 = v207
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
								if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v235 = int64(4)
								} else {
									v235 = int64(1)
								}
								if v228 == int32(92) {
									v238 = int64(2)
								} else {
									v238 = v235
								}
								v249 = v224 + v238
							}
						}
						if base.Ui64(int64(1073741824)) <= base.Ui64(v249) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v369 = m.ExcPending
							if v369 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v372 = m.ExcPending
								if v372 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
									mBase = m.M
									v376 = m.ExcPending
									if v376 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_byteaout_1), int32(438), int32(_a_F_byteaout_2))
										mBase = m.M
										v381 = m.ExcPending
										if v381 != 0 {
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
							v261 = v249
							v263 = F_palloc(m, base.I32_wrap_i64(v261))
							mBase = m.M
							v264 = m.ExcPending
							if v264 != 0 {
								return int32(0)
							} else {
								v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
								v266 = int32(1)
								v267 = v265 & v266
								if v265 == v266 {
									v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
									if base.Ui32((v271-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v296 = int32(4)
										if v267 != 0 {
											v297 = v126
										} else {
											v297 = v124
										}
										v298 = v263
										v300 = v297
										v302 = v296
										for {
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
											if v308 == int32(92) {
												v311 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
												v345 = v298 + int32(2)
											} else {
												if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v321 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
													v323 = int32(7)
													v325 = int32(48)
													v326 = v308&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
													v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
													v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
													v345 = v298 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
													v345 = v298 + int32(1)
												}
											}
											v346 = int32(1)
											v349 = v302 - v346
											if v349 != 0 {
												v298 = v345
												v300 = v300 + v346
												v302 = v349
												continue
											} else {
												break
											}
											break
										}
										v350 = v345
										v353 = v263
									} else {
										if v271 == int32(18) {
											v282 = int32(16)
										} else {
											v282 = int32(0)
										}
										v294 = v282
										if v294 != 0 {
											v296 = v294
											if v267 != 0 {
												v297 = v126
											} else {
												v297 = v124
											}
											v298 = v263
											v300 = v297
											v302 = v296
											for {
												v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
												if v308 == int32(92) {
													v311 = int32(_a_F_byteaout_3)
													*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
													v345 = v298 + int32(2)
												} else {
													if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
														v321 = int32(92)
														*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
														v323 = int32(7)
														v325 = int32(48)
														v326 = v308&v323 | v325
														*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
														v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
														*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
														v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
														*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
														v345 = v298 + int32(4)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
														v345 = v298 + int32(1)
													}
												}
												v346 = int32(1)
												v349 = v302 - v346
												if v349 != 0 {
													v298 = v345
													v300 = v300 + v346
													v302 = v349
													continue
												} else {
													break
												}
												break
											}
											v350 = v345
											v353 = v263
										} else {
											v350 = v263
											v353 = v263
										}
									}
								} else {
									v283 = int32(1)
									if v267 != 0 {
										v294 = int32(base.Ui32(v265)>>(uint(v283)%32)) - v283
									} else {
										v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
										v294 = int32(base.Ui32(v287)>>(uint(int32(2))%32)) - int32(4)
									}
									if v294 != 0 {
										v296 = v294
										if v267 != 0 {
											v297 = v126
										} else {
											v297 = v124
										}
										v298 = v263
										v300 = v297
										v302 = v296
										for {
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
											if v308 == int32(92) {
												v311 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
												v345 = v298 + int32(2)
											} else {
												if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v321 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
													v323 = int32(7)
													v325 = int32(48)
													v326 = v308&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
													v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
													v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
													v345 = v298 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
													v345 = v298 + int32(1)
												}
											}
											v346 = int32(1)
											v349 = v302 - v346
											if v349 != 0 {
												v298 = v345
												v300 = v300 + v346
												v302 = v349
												continue
											} else {
												break
											}
											break
										}
										v350 = v345
										v353 = v263
									} else {
										v350 = v263
										v353 = v263
									}
								}
								v360 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
								m.G0 = v13 + int32(16)
								return v353
							}
						}
					} else {
						v261 = int64(1)
						v263 = F_palloc(m, base.I32_wrap_i64(v261))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v266 = int32(1)
							v267 = v265 & v266
							if v265 == v266 {
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
								if base.Ui32((v271-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v296 = int32(4)
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									if v271 == int32(18) {
										v282 = int32(16)
									} else {
										v282 = int32(0)
									}
									v294 = v282
									if v294 != 0 {
										v296 = v294
										if v267 != 0 {
											v297 = v126
										} else {
											v297 = v124
										}
										v298 = v263
										v300 = v297
										v302 = v296
										for {
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
											if v308 == int32(92) {
												v311 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
												v345 = v298 + int32(2)
											} else {
												if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v321 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
													v323 = int32(7)
													v325 = int32(48)
													v326 = v308&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
													v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
													v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
													v345 = v298 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
													v345 = v298 + int32(1)
												}
											}
											v346 = int32(1)
											v349 = v302 - v346
											if v349 != 0 {
												v298 = v345
												v300 = v300 + v346
												v302 = v349
												continue
											} else {
												break
											}
											break
										}
										v350 = v345
										v353 = v263
									} else {
										v350 = v263
										v353 = v263
									}
								}
							} else {
								v283 = int32(1)
								if v267 != 0 {
									v294 = int32(base.Ui32(v265)>>(uint(v283)%32)) - v283
								} else {
									v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v294 = int32(base.Ui32(v287)>>(uint(int32(2))%32)) - int32(4)
								}
								if v294 != 0 {
									v296 = v294
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									v350 = v263
									v353 = v263
								}
							}
							v360 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
							m.G0 = v13 + int32(16)
							return v353
						}
					}
				}
			} else {
				v144 = int32(1)
				if v129 != 0 {
					v154 = int32(base.Ui32(v127)>>(uint(v144)%32)) - v144
				} else {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v154 = int32(base.Ui32(v148)>>(uint(int32(2))%32)) - int32(4)
				}
				if v154 != 0 {
					v157 = v154
					if v129 != 0 {
						v158 = v126
					} else {
						v158 = v124
					}
					if v157 == int32(1) {
						v215 = v158
						v224 = int64(1)
						v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
						if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v235 = int64(4)
						} else {
							v235 = int64(1)
						}
						if v228 == int32(92) {
							v238 = int64(2)
						} else {
							v238 = v235
						}
						v249 = v224 + v238
					} else {
						v168 = v158
						v170 = int32(0)
						v177 = int64(1)
						for {
							v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
							if base.Ui32((v181-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v188 = int64(4)
							} else {
								v188 = int64(1)
							}
							if v181 == int32(92) {
								v191 = int64(2)
							} else {
								v191 = v188
							}
							v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
							if base.Ui32((v196-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v203 = int64(4)
							} else {
								v203 = int64(1)
							}
							if v196 == int32(92) {
								v206 = int64(2)
							} else {
								v206 = v203
							}
							v207 = v177 + v191 + v206
							v208 = int32(2)
							v209 = v168 + v208
							v211 = v170 + v208
							if v211 != v157&int32(-2) {
								v168 = v209
								v170 = v211
								v177 = v207
								continue
							} else {
								break
							}
							break
						}
						if v157&int32(1) == int32(0) {
							v249 = v207
						} else {
							v215 = v209
							v224 = v207
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
							if base.Ui32((v228-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v235 = int64(4)
							} else {
								v235 = int64(1)
							}
							if v228 == int32(92) {
								v238 = int64(2)
							} else {
								v238 = v235
							}
							v249 = v224 + v238
						}
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v249) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v372 = m.ExcPending
							if v372 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
								mBase = m.M
								v376 = m.ExcPending
								if v376 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_byteaout_1), int32(438), int32(_a_F_byteaout_2))
									mBase = m.M
									v381 = m.ExcPending
									if v381 != 0 {
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
						v261 = v249
						v263 = F_palloc(m, base.I32_wrap_i64(v261))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
							v266 = int32(1)
							v267 = v265 & v266
							if v265 == v266 {
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
								if base.Ui32((v271-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v296 = int32(4)
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									if v271 == int32(18) {
										v282 = int32(16)
									} else {
										v282 = int32(0)
									}
									v294 = v282
									if v294 != 0 {
										v296 = v294
										if v267 != 0 {
											v297 = v126
										} else {
											v297 = v124
										}
										v298 = v263
										v300 = v297
										v302 = v296
										for {
											v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
											if v308 == int32(92) {
												v311 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
												v345 = v298 + int32(2)
											} else {
												if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v321 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
													v323 = int32(7)
													v325 = int32(48)
													v326 = v308&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
													v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
													v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
													*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
													v345 = v298 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
													v345 = v298 + int32(1)
												}
											}
											v346 = int32(1)
											v349 = v302 - v346
											if v349 != 0 {
												v298 = v345
												v300 = v300 + v346
												v302 = v349
												continue
											} else {
												break
											}
											break
										}
										v350 = v345
										v353 = v263
									} else {
										v350 = v263
										v353 = v263
									}
								}
							} else {
								v283 = int32(1)
								if v267 != 0 {
									v294 = int32(base.Ui32(v265)>>(uint(v283)%32)) - v283
								} else {
									v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v294 = int32(base.Ui32(v287)>>(uint(int32(2))%32)) - int32(4)
								}
								if v294 != 0 {
									v296 = v294
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									v350 = v263
									v353 = v263
								}
							}
							v360 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
							m.G0 = v13 + int32(16)
							return v353
						}
					}
				} else {
					v261 = int64(1)
					v263 = F_palloc(m, base.I32_wrap_i64(v261))
					mBase = m.M
					v264 = m.ExcPending
					if v264 != 0 {
						return int32(0)
					} else {
						v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
						v266 = int32(1)
						v267 = v265 & v266
						if v265 == v266 {
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
							if base.Ui32((v271-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v296 = int32(4)
								if v267 != 0 {
									v297 = v126
								} else {
									v297 = v124
								}
								v298 = v263
								v300 = v297
								v302 = v296
								for {
									v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
									if v308 == int32(92) {
										v311 = int32(_a_F_byteaout_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
										v345 = v298 + int32(2)
									} else {
										if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v321 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
											v323 = int32(7)
											v325 = int32(48)
											v326 = v308&v323 | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
											v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
											v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
											v345 = v298 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
											v345 = v298 + int32(1)
										}
									}
									v346 = int32(1)
									v349 = v302 - v346
									if v349 != 0 {
										v298 = v345
										v300 = v300 + v346
										v302 = v349
										continue
									} else {
										break
									}
									break
								}
								v350 = v345
								v353 = v263
							} else {
								if v271 == int32(18) {
									v282 = int32(16)
								} else {
									v282 = int32(0)
								}
								v294 = v282
								if v294 != 0 {
									v296 = v294
									if v267 != 0 {
										v297 = v126
									} else {
										v297 = v124
									}
									v298 = v263
									v300 = v297
									v302 = v296
									for {
										v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
										if v308 == int32(92) {
											v311 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
											v345 = v298 + int32(2)
										} else {
											if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v321 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
												v323 = int32(7)
												v325 = int32(48)
												v326 = v308&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
												v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
												v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
												*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
												v345 = v298 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
												v345 = v298 + int32(1)
											}
										}
										v346 = int32(1)
										v349 = v302 - v346
										if v349 != 0 {
											v298 = v345
											v300 = v300 + v346
											v302 = v349
											continue
										} else {
											break
										}
										break
									}
									v350 = v345
									v353 = v263
								} else {
									v350 = v263
									v353 = v263
								}
							}
						} else {
							v283 = int32(1)
							if v267 != 0 {
								v294 = int32(base.Ui32(v265)>>(uint(v283)%32)) - v283
							} else {
								v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v294 = int32(base.Ui32(v287)>>(uint(int32(2))%32)) - int32(4)
							}
							if v294 != 0 {
								v296 = v294
								if v267 != 0 {
									v297 = v126
								} else {
									v297 = v124
								}
								v298 = v263
								v300 = v297
								v302 = v296
								for {
									v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
									if v308 == int32(92) {
										v311 = int32(_a_F_byteaout_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v311)
										v345 = v298 + int32(2)
									} else {
										if base.Ui32((v308-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v321 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v321)
											v323 = int32(7)
											v325 = int32(48)
											v326 = v308&v323 | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+3)) = uint8(v326)
											v331 = int32(base.Ui32(v308)>>(uint(int32(6))%32)) | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)) = uint8(v331)
											v338 = int32(base.Ui32(v308)>>(uint(int32(3))%32))&v323 | v325
											*(*uint8)(unsafe.Add(mBase, uint32(v298)+2)) = uint8(v338)
											v345 = v298 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v308)
											v345 = v298 + int32(1)
										}
									}
									v346 = int32(1)
									v349 = v302 - v346
									if v349 != 0 {
										v298 = v345
										v300 = v300 + v346
										v302 = v349
										continue
									} else {
										break
									}
									break
								}
								v350 = v345
								v353 = v263
							} else {
								v350 = v263
								v353 = v263
							}
						}
						v360 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
						m.G0 = v13 + int32(16)
						return v353
					}
				}
			}
		case 1:
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v22 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v51 = v38
			} else {
				v39 = int32(1)
				if v22&v39 != 0 {
					v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v56 = F_palloc(m, v51<<(uint(int32(1))%32)+int32(3))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = int32(_a_F_byteaout_4)
				*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v58)
				v61 = v56 + int32(2)
				v62 = int32(1)
				v63 = v16 + v62
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				v68 = v66 & v62
				if v68 != 0 {
					v69 = v63
				} else {
					v69 = v16 + int32(4)
				}
				if v66 == int32(1) {
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
					if v75 == int32(18) {
						v78 = int32(16)
					} else {
						v78 = int32(0)
					}
					if base.Ui32((v75-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v85 = int32(4)
					} else {
						v85 = v78
					}
					v96 = v85
				} else {
					v86 = int32(1)
					if v68 != 0 {
						v96 = int32(base.Ui32(v66)>>(uint(v86)%32)) - v86
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						v96 = int32(base.Ui32(v90)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				if v96 != 0 {
					v99 = v69
					v101 = v61
					for {
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
						v104 = int32(1)
						v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103<<(uint(v104)%32))+uint32(_c_F_byteaout[1]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v101))) = uint16(v106)
						v111 = v99 + v104
						if base.Ui32(v111) < base.Ui32(v69+v96) {
							v99 = v111
							v101 = v101 + int32(2)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v350 = v61 + base.I32_wrap_i64(base.I64_extend_i32_u(v96)<<(uint(int64(1))%64))
				v353 = v56
				v360 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v360)
				m.G0 = v13 + int32(16)
				return v353
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v385 = m.ExcPending
			if v385 != 0 {
				return int32(0)
			} else {
				v387 = *(*int32)(unsafe.Add(mBase, _c_F_byteaout[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v387
				F_errmsg_internal(m, int32(_a_F_byteaout_5), v13)
				mBase = m.M
				v391 = m.ExcPending
				if v391 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_byteaout_1), int32(469), int32(_a_F_byteaout_2))
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
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
func F_byteapos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v193 int32
	_ = v193
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == v46 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	return v193
L16:
	;
	if v79 <= v45 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	if v76 <= int32(0) {
		v193 = v46
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v79 = int32(4)
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v63 = int32(1)
	if v47&v63 != 0 {
		v76 = int32(base.Ui32(v47)>>(uint(v63)%32)) - v63
		goto L17
	} else {
		goto L25
	}
L21:
	;
	if v51 == int32(18) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(16)
	goto L24
L23:
	;
	v62 = int32(0)
	goto L24
L24:
	;
	v76 = v62
	goto L17
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v76 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L26:
	;
	v79 = v76
	goto L16
L27:
	;
	v83 = int32(1)
	if v16&v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v193 = int32(0)
	goto L15
L30:
	;
	v89 = v83
	goto L32
L31:
	;
	v89 = int32(4)
	goto L32
L32:
	;
	v92 = int32(1)
	if v47&v92 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v92
	goto L35
L34:
	;
	v96 = int32(4)
	goto L35
L35:
	;
	v97 = v14 + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v99 = int32(0)
	v100 = v9 + v89
	goto L36
L36:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v98 != v106 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L29
L38:
	;
	v173 = int32(1)
	v176 = v99 + v173
	if v176 != v45-v79+v83 {
		v99 = v176
		v100 = v100 + v173
		goto L36
	} else {
		goto L59
	}
L39:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v79) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if v169 != 0 {
		goto L38
	} else {
		goto L58
	}
L41:
	;
	v169 = int32(0)
	goto L40
L42:
	;
	v143 = v138
	v144 = v139
	v145 = v140
	goto L52
L43:
	;
	if (v100|v97)&int32(3) != 0 {
		v138 = v100
		v139 = v97
		v140 = v79
		goto L42
	} else {
		goto L46
	}
L44:
	;
	v131 = v100
	v132 = v97
	v133 = v79
	goto L45
L45:
	;
	if v133 == int32(0) {
		goto L41
	} else {
		goto L51
	}
L46:
	;
	v115 = v100
	v116 = v97
	v117 = v79
	goto L47
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v120 != v121 {
		v138 = v115
		v139 = v116
		v140 = v117
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v131 = v126
	v132 = v124
	v133 = v128
	goto L45
L49:
	;
	v123 = int32(4)
	v124 = v116 + v123
	v126 = v115 + v123
	v128 = v117 - v123
	if base.Ui32(int32(3)) < base.Ui32(v128) {
		v115 = v126
		v116 = v124
		v117 = v128
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v138 = v131
	v139 = v132
	v140 = v133
	goto L42
L52:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 == v149 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v169 = v148 - v149
	goto L40
L54:
	;
	v151 = int32(1)
	v156 = v145 - v151
	if v156 != 0 {
		v143 = v143 + v151
		v144 = v144 + v151
		v145 = v156
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L41
L58:
	;
	return v99 + int32(1)
L59:
	;
	goto L37
}
func F_byteartrim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v12 = F_dobyteatrim(m, v3, v8, int32(0), int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
