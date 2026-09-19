package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int64
	_ = v338
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int64
	_ = v391
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int64
	_ = v549
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int64
	_ = v635
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = m.G0
	v15 = v13 - int32(432)
	m.G0 = v15
	F_ReserveExternalFD(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[1]))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_pg_usleep(m, v20*int32(_a_F_BackendMain_0))
	mBase = m.M
	goto L5
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
	v35 = m.G0
	v37 = v35 + int32(-64)
	m.G0 = v37
	v40 = F_palloc0(m, int32(524))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3])) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[5])) = v40
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[6])) = int32(2)
	v284 = int32(_a_F_BackendMain_2)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+292)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v40)+276)) = v284
	v289 = int32(1132)
	v291 = m.G0
	v293 = v291 - int32(32)
	m.G0 = v293
	switch int32(1134) {
	case 0, 2:
		v303 = v289
		goto L72
	default:
		goto L73
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	base.MemoryCopy(m, v40+int32(144), v11+int32(4), v44)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+272)) = v50
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(12)))))
	if v56 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v59
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[7]))
	v68 = m.G0
	v70 = v68 - int32(16)
	m.G0 = v70
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v66
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v190 = int32(_a_F_BackendMain_3)
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[8])) = v190
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v196 = F_MemoryContextAlloc(m, v194, v190)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L58
	}
L14:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[9]))
	v99 = m.G0
	v101 = v99 - int32(16)
	m.G0 = v101
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v97
	if v40 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	m.G0 = v70 + int32(16)
	goto L14
L16:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if v75 == int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v40)+400))
	if v66 == v78 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v40)+384))
	if int32(0) < v80 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v66 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v83 = F_pq_getkeepalivesidle(m, v40)
	mBase = m.M
	if int32(0) <= v83 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v40)+384))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v88
	goto L24
L23:
	;
	goto L24
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+400)) = v91
	goto L15
L25:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[10]))
	v130 = m.G0
	v132 = v130 - int32(16)
	m.G0 = v132
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v128
	if v40 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	m.G0 = v101 + int32(16)
	goto L25
L27:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if v106 == int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+404))
	if v97 == v109 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v40)+388))
	if int32(0) < v111 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v97 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v114 = F_pq_getkeepalivesinterval(m, v40)
	mBase = m.M
	if int32(0) <= v114 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v40)+388))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v119
	goto L35
L34:
	;
	goto L35
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+404)) = v122
	goto L26
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[11]))
	v161 = m.G0
	v163 = v161 - int32(16)
	m.G0 = v163
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v159
	if v40 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L37:
	;
	m.G0 = v132 + int32(16)
	goto L36
L38:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if v137 == int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v40)+408))
	if v128 == v140 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v40)+392))
	if int32(0) < v142 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v128 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v145 = F_pq_getkeepalivescount(m, v40)
	mBase = m.M
	if int32(0) <= v145 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v40)+392))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v150
	goto L46
L45:
	;
	goto L46
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+408)) = v153
	goto L37
L47:
	;
	goto L13
L48:
	;
	m.G0 = v163 + int32(16)
	goto L47
L49:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if v168 == int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v40)+412))
	if v159 == v171 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v40)+396))
	if int32(0) < v173 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if v159 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v176 = F_pq_gettcpusertimeout(m, v40)
	mBase = m.M
	if int32(0) <= v176 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	goto L48
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v40)+396))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v181
	goto L57
L56:
	;
	goto L57
L57:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+412)) = v184
	goto L48
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[12])) = v196
	v200 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[13])) = v200
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[14])) = v200
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[15])) = v200
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[16])) = v200
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[17])) = uint8(v200)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[18])) = uint8(v200)
	F_on_proc_exit(m, int32(801))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v222 = m.G0
	v223 = int32(16)
	v224 = v222 - v223
	m.G0 = v224
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = int32(2048)
	m.G0 = v224 + v223
	goto L60
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = int32(1)
	v238 = F_CreateWaitEventSet(m, int32(0), int32(3))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[19])) = v238
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_AddWaitEventToSet(m, v238, int32(4), v242, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[19]))
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[20]))
	F_AddWaitEventToSet(m, v247, int32(1), int32(-1), v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[19]))
	F_AddWaitEventToSet(m, v255, int32(16), int32(-1), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v37 - int32(-64)
	goto L6
L71:
	;
	v321 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[21])) = v321
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[22])) = v321
	v331 = v321
	goto L78
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v303
	F_sigemptyset(m, v293+int32(16))
	mBase = m.M
	goto L75
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[23])) = v289
	v303 = int32(_a_F_BackendMain_4)
	goto L72
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+24)) = int32(268435456)
	v317 = F___sigaction(m, int32(15), v293+int32(12), int32(0))
	mBase = m.M
	m.G0 = v293 + int32(32)
	goto L71
L77:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_BackendMain_5), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L83
	}
L78:
	;
	v333 = int32(40)
	v334 = v331 * v333
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[24]))) = uint8(v335)
	*(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[25]))) = v331
	v338 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[26]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[27]))) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[28]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[29]))) = v335
	*(*uint8)(unsafe.Add(mBase, uint32(v334)+uint32(_c_F_BackendMain[30]))) = uint8(v335)
	v349 = v331 | int32(1)
	v351 = v349 * v333
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[24]))) = uint8(v335)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[25]))) = v349
	*(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[26]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[27]))) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[28]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[29]))) = v335
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_BackendMain[30]))) = uint8(v335)
	v366 = v331 | int32(2)
	v368 = v366 * v333
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[24]))) = uint8(v335)
	*(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[25]))) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[26]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[27]))) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[28]))) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[29]))) = v335
	*(*uint8)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_BackendMain[30]))) = uint8(v335)
	if v331 != int32(20) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[31])) = uint8(v404)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L77
L80:
	;
	v385 = v331 | int32(3)
	v387 = v385 * int32(40)
	v388 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[24]))) = uint8(v388)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[25]))) = v385
	v391 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[26]))) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[27]))) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[28]))) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[29]))) = v388
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_BackendMain[30]))) = uint8(v388)
	v331 = v331 + int32(4)
	goto L78
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)) = uint8(v413)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+176)) = uint8(v413)
	v417 = int32(144)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v40)+272))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[32])))
	v430 = F_pg_getnameinfo_all(m, v40+v417, v419, v15+int32(176), int32(255), v15+v417, int32(32), v427^int32(3))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v484 = v15 + int32(176)
	v485 = F_MemoryContextStrdup(m, v482, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L101
	}
L85:
	;
	if v430 == int32(0) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v436 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v436 == int32(0) {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v442 = int32(_a_F_BackendMain_6)
	v444 = v430 + int32(1)
	if v444 == int32(0) {
		v464 = v442
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v464 + base.B2i32(v466 == int32(0))
	F_errmsg_internal(m, int32(_a_F_BackendMain_7), v15+int32(112))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L99
	}
L90:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	goto L89
L91:
	;
	v448 = v442
	v449 = v444
	goto L92
L92:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	if v450 == int32(0) {
		v464 = v448
		goto L90
	} else {
		goto L94
	}
L93:
	;
	v464 = v460
	goto L90
L94:
	;
	v454 = v448
	goto L95
L95:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+1)))
	if v458 != 0 {
		v454 = v454 + int32(1)
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v460 = v454 + int32(2)
	v462 = v449 + int32(1)
	if v462 != 0 {
		v448 = v460
		v449 = v462
		goto L92
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	goto L93
L99:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(220), int32(_a_F_BackendMain_9))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L84
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+276)) = v485
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v491 = v15 + int32(144)
	v492 = F_MemoryContextStrdup(m, v489, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+292)) = v492
	v496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[33])))
	if v496&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v430 != 0 {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	v504 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v501 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), v528, int32(_a_F_BackendMain_9))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L114
	}
L107:
	;
	if v504 == int32(0) {
		goto L103
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if v504 == int32(0) {
		goto L103
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v484
	F_errmsg(m, int32(_a_F_BackendMain_10), v15+int32(96))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v528 = int32(236)
	goto L106
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v15 + int32(176)
	F_errmsg(m, int32(_a_F_BackendMain_11), v15+int32(80))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v528 = int32(240)
	goto L106
L114:
	;
	goto L103
L115:
	;
	F_RegisterTimeout(m, int32(0), int32(1133))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L159
	}
L116:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[32])))
	if v535&int32(1) == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v541 = v15 + int32(176)
	v542 = int32(_a_F_BackendMain_12)
	v546 = m.G0
	v548 = v546 - int32(32)
	v549 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v548)+24)) = v549
	*(*int64)(unsafe.Add(mBase, uint32(v548)+16)) = v549
	*(*int64)(unsafe.Add(mBase, uint32(v548)+8)) = v549
	*(*int64)(unsafe.Add(mBase, uint32(v548))) = v549
	v557 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[34])))
	if v557 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v626 = F_strlen(m, v541)
	mBase = m.M
	if base.Ui32(v626) <= base.Ui32(v625) {
		goto L115
	} else {
		goto L137
	}
L119:
	;
	v625 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[35])))
	if v561 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v565 = v541
	goto L125
L123:
	;
	goto L124
L124:
	;
	v575 = v542
	v576 = v557
	goto L128
L125:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v571 == v557 {
		v565 = v565 + int32(1)
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v625 = v565 - v541
	goto L118
L127:
	;
	goto L126
L128:
	;
	v583 = v548 + int32(base.Ui32(v576)>>(uint(int32(3))%32))&int32(28)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	v585 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v584 | v585<<(uint(v576)%32)
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+1)))
	if v589 != 0 {
		v575 = v575 + v585
		v576 = v589
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v592 == int32(0) {
		v615 = v541
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L129
L131:
	;
	v625 = v615 - v541
	goto L118
L132:
	;
	v596 = v541
	v597 = v592
	goto L133
L133:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v548+int32(base.Ui32(v597)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v605)>>(uint(v597)%32))&int32(1) == int32(0) {
		v615 = v596
		goto L131
	} else {
		goto L135
	}
L134:
	;
	v615 = v613
	goto L131
L135:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	v613 = v596 + int32(1)
	if v611 != 0 {
		v596 = v613
		v597 = v611
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v628 = int32(_a_F_BackendMain_13)
	v632 = m.G0
	v634 = v632 - int32(32)
	v635 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v634)+24)) = v635
	*(*int64)(unsafe.Add(mBase, uint32(v634)+16)) = v635
	*(*int64)(unsafe.Add(mBase, uint32(v634)+8)) = v635
	*(*int64)(unsafe.Add(mBase, uint32(v634))) = v635
	v643 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[36])))
	if v643 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if base.Ui32(v626) <= base.Ui32(v711) {
		goto L115
	} else {
		goto L157
	}
L139:
	;
	v711 = int32(0)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[37])))
	if v647 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v651 = v541
	goto L145
L143:
	;
	goto L144
L144:
	;
	v661 = v628
	v662 = v643
	goto L148
L145:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	if v657 == v643 {
		v651 = v651 + int32(1)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v711 = v651 - v541
	goto L138
L147:
	;
	goto L146
L148:
	;
	v669 = v634 + int32(base.Ui32(v662)>>(uint(int32(3))%32))&int32(28)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	v671 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v670 | v671<<(uint(v662)%32)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)))
	if v675 != 0 {
		v661 = v661 + v671
		v662 = v675
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	if v678 == int32(0) {
		v701 = v541
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	v711 = v701 - v541
	goto L138
L152:
	;
	v682 = v541
	v683 = v678
	goto L153
L153:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v634+int32(base.Ui32(v683)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v691)>>(uint(v683)%32))&int32(1) == int32(0) {
		v701 = v682
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v701 = v699
	goto L151
L155:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+1)))
	v699 = v682 + int32(1)
	if v697 != 0 {
		v682 = v699
		v683 = v697
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	v715 = F_MemoryContextStrdup(m, v714, v541)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+280)) = v715
	goto L115
L159:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[38]))
	F_enable_timeout_after(m, int32(0), v726*int32(1000))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L164
L162:
	;
	v757 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[18])) = uint8(v757)
	goto L169
L163:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+uint32(_c_F_BackendMain[39]))))
	v755 = v754
	goto L162
L164:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[14]))
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[13]))
	if v743 < v745 {
		goto L163
	} else {
		goto L166
	}
L165:
	;
	v755 = int32(-1)
	goto L162
L166:
	;
	v747 = F_pq_recvbuf(m)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	if v747 == int32(0) {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	if v755 == int32(-1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	F_InitProcess(m)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L253
	}
L171:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L249
	}
L172:
	;
	if v755 == int32(22) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[40])))
	if v764 != int32(1) {
		goto L171
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v782 = int32(0)
	v784 = F_ProcessStartupPacket(m, v40, v782, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L181
	}
L176:
	;
	v769 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	if v769 == int32(0) {
		goto L171
	} else {
		goto L178
	}
L178:
	;
	F_errmsg(m, int32(_a_F_BackendMain_14), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(477), int32(_a_F_BackendMain_15))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	goto L171
L181:
	;
	if v784 != 0 {
		goto L171
	} else {
		goto L182
	}
L182:
	;
	switch v12 - int32(1) {
	case 0:
		goto L190
	case 1:
		goto L188
	case 2:
		goto L187
	case 3:
		goto L189
	case 4:
		goto L186
	default:
		goto L185
	}
L183:
	;
	F_errdetail(m, int32(_a_F_BackendMain_16), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L246
	}
L184:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L241
	}
L185:
	;
	F_disable_timeout(m, int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L214
	}
L186:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L210
	}
L187:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L206
	}
L188:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L202
	}
L189:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[41])))
	if v805 == int32(0) {
		goto L184
	} else {
		goto L195
	}
L190:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errmsg(m, int32(_a_F_BackendMain_17), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(317), int32(_a_F_BackendMain_9))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[42])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_BackendMain_18), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if v809 == int32(1) {
		goto L183
	} else {
		goto L199
	}
L199:
	;
	F_errdetail(m, int32(_a_F_BackendMain_19), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(336), int32(_a_F_BackendMain_9))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(_a_F_BackendMain_20), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(341), int32(_a_F_BackendMain_9))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_errmsg(m, int32(_a_F_BackendMain_21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(346), int32(_a_F_BackendMain_9))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(_a_F_BackendMain_22))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errmsg(m, int32(_a_F_BackendMain_23), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(351), int32(_a_F_BackendMain_9))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_BackendMain_24), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v890 = v15 + int32(128)
	F_initStringInfo(m, v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackendMain[43])))
	if v894 == int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	goto L222
L219:
	;
	goto L220
L220:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v40)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v911
	v914 = v15 + int32(128)
	F_appendStringInfo(m, v914, int32(_a_F_BackendMain_25), v15+int32(32))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L226
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v902
	F_appendStringInfo(m, v890, int32(_a_F_BackendMain_25), v15+int32(48))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L225
	}
L222:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[44]))
	goto L224
L224:
	;
	goto L221
L225:
	;
	goto L220
L226:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v40)+360))
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920))))
	if v921 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v920
	F_appendStringInfo(m, v914, int32(_a_F_BackendMain_25), v15+int32(16))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v929 = v15 + int32(128)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v40)+276))
	F_appendStringInfoString(m, v929, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v40)+292))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	if v934 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v933
	F_appendStringInfo(m, v929, int32(_a_F_BackendMain_26), v15)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	if v939 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L234
L236:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	F_pfree(m, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L240
	}
L237:
	;
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[45]))
	v944 = F_GetBackendTypeDesc(m, v943)
	mBase = m.M
	goto L239
L238:
	;
	goto L239
L239:
	;
	goto L236
L240:
	;
	m.G0 = v15 + int32(432)
	goto L170
L241:
	;
	F_errcode(m, int32(50463173))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errmsg(m, int32(_a_F_BackendMain_27), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errdetail(m, int32(_a_F_BackendMain_28), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(324), int32(_a_F_BackendMain_9))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(64)
	F_errhint(m, int32(_a_F_BackendMain_29), v15-int32(-64))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_BackendMain_8), int32(331), int32(_a_F_BackendMain_9))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_BackendMain_24), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_check_on_shmem_exit_lists_are_empty(m)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackendMain[3])) = v1003
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_BackendMain[5]))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+360))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+364))
	F_PostgresMain(m, v1007, v1008)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
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
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int64
	_ = v270
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int64
	_ = v323
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int64
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
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
	v707 = int32(m.ExcTag)
	v708 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v707 == int32(0) {
		goto L170
	} else {
		goto L171
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
	v451 = v14
	v453 = v16
	goto L9
L9:
	;
	if v453 != 0 {
		goto L95
	} else {
		goto L96
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
	goto L27
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v75 = int32(2)
	v79 = v73 & v75
	if v79 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v80 = int32(915)
	goto L30
L29:
	;
	v80 = int32(-2)
	goto L30
L30:
	;
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	switch v80 + int32(2) {
	case 0, 2:
		v94 = v80
		goto L32
	default:
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	if v79 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v94
	F_sigemptyset(m, v84+int32(16))
	mBase = m.M
	goto L35
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[5])) = v80
	v94 = int32(_a_F_BackgroundWorkerMain_4)
	goto L32
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = int32(268435456)
	v108 = F___sigaction(m, v75, v84+int32(12), int32(0))
	mBase = m.M
	m.G0 = v84 + int32(32)
	goto L31
L37:
	;
	v116 = int32(917)
	goto L39
L38:
	;
	v116 = int32(-2)
	goto L39
L39:
	;
	v118 = m.G0
	v120 = v118 - int32(32)
	m.G0 = v120
	switch v116 + int32(2) {
	case 0, 2:
		v130 = v116
		goto L41
	default:
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	if v79 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v130
	F_sigemptyset(m, v120+int32(16))
	mBase = m.M
	goto L44
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[6])) = v116
	v130 = int32(_a_F_BackgroundWorkerMain_4)
	goto L41
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = int32(268435456)
	v144 = F___sigaction(m, int32(10), v120+int32(12), int32(0))
	mBase = m.M
	m.G0 = v120 + int32(32)
	goto L40
L46:
	;
	v152 = int32(919)
	goto L48
L47:
	;
	v152 = int32(-2)
	goto L48
L48:
	;
	v154 = m.G0
	v156 = v154 - int32(32)
	m.G0 = v156
	switch v152 + int32(2) {
	case 0, 2:
		v166 = v152
		goto L50
	default:
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v186 = int32(923)
	v188 = m.G0
	v190 = v188 - int32(32)
	m.G0 = v190
	switch int32(925) {
	case 0, 2:
		v200 = v186
		goto L56
	default:
		goto L57
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+12)) = v166
	F_sigemptyset(m, v156+int32(16))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[7])) = v152
	v166 = int32(_a_F_BackgroundWorkerMain_4)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+24)) = int32(268435456)
	v180 = F___sigaction(m, int32(8), v156+int32(12), int32(0))
	mBase = m.M
	m.G0 = v156 + int32(32)
	goto L49
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v220 = int32(-2)
	v222 = m.G0
	v224 = v222 - int32(32)
	m.G0 = v224
	switch int32(0) {
	case 0, 2:
		v234 = v220
		goto L62
	default:
		goto L63
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = v200
	F_sigemptyset(m, v190+int32(16))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[8])) = v186
	v200 = int32(_a_F_BackgroundWorkerMain_4)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+24)) = int32(268435456)
	v214 = F___sigaction(m, int32(15), v190+int32(12), int32(0))
	mBase = m.M
	m.G0 = v190 + int32(32)
	goto L55
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v253 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[9])) = v253
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[10])) = v253
	v263 = v253
	goto L68
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = v234
	F_sigemptyset(m, v224+int32(16))
	mBase = m.M
	goto L65
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[11])) = v220
	v234 = int32(_a_F_BackgroundWorkerMain_4)
	goto L62
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+24)) = int32(268435456)
	v248 = F___sigaction(m, int32(1), v224+int32(12), int32(0))
	mBase = m.M
	m.G0 = v224 + int32(32)
	goto L61
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v343 = int32(-2)
	v345 = m.G0
	v347 = v345 - int32(32)
	m.G0 = v347
	switch int32(0) {
	case 0, 2:
		v357 = v343
		goto L74
	default:
		goto L75
	}
L68:
	;
	v265 = int32(40)
	v266 = v263 * v265
	v267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[13]))) = v263
	v270 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[14]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[15]))) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[16]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[17]))) = v267
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v267)
	v281 = v263 | int32(1)
	v283 = v281 * v265
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[13]))) = v281
	*(*int64)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[14]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[15]))) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[16]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[17]))) = v267
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v267)
	v298 = v263 | int32(2)
	v300 = v298 * v265
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[13]))) = v298
	*(*int64)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[14]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[15]))) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[16]))) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[17]))) = v267
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v267)
	if v263 != int32(20) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[19])) = uint8(v336)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L67
L70:
	;
	v317 = v263 | int32(3)
	v319 = v317 * int32(40)
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[12]))) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[13]))) = v317
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[14]))) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[15]))) = v320
	*(*int64)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[16]))) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[17]))) = v320
	*(*uint8)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_BackgroundWorkerMain[18]))) = uint8(v320)
	v263 = v263 + int32(4)
	goto L68
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v377 = int32(-2)
	v379 = m.G0
	v381 = v379 - int32(32)
	m.G0 = v381
	switch int32(0) {
	case 0, 2:
		v391 = v377
		goto L80
	default:
		goto L81
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+12)) = v357
	F_sigemptyset(m, v347+int32(16))
	mBase = m.M
	goto L77
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[20])) = v343
	v357 = int32(_a_F_BackgroundWorkerMain_4)
	goto L74
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+24)) = int32(268435456)
	v371 = F___sigaction(m, int32(13), v347+int32(12), int32(0))
	mBase = m.M
	m.G0 = v347 + int32(32)
	goto L73
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v42
	v411 = int32(0)
	v413 = m.G0
	v415 = v413 - int32(32)
	m.G0 = v415
	switch int32(2) {
	case 0, 2:
		v425 = v411
		goto L86
	default:
		goto L87
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+12)) = v391
	F_sigemptyset(m, v381+int32(16))
	mBase = m.M
	goto L83
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[21])) = v377
	v391 = int32(_a_F_BackgroundWorkerMain_4)
	goto L80
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = int32(268435456)
	v405 = F___sigaction(m, int32(12), v381+int32(12), int32(0))
	mBase = m.M
	m.G0 = v381 + int32(32)
	goto L79
L85:
	;
	goto L91
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415)+12)) = v425
	F_sigemptyset(m, v415+int32(16))
	mBase = m.M
	goto L88
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[22])) = v411
	v425 = int32(_a_F_BackgroundWorkerMain_4)
	goto L86
L88:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415)+24)) = int32(268435457)
	v439 = F___sigaction(m, int32(17), v415+int32(12), int32(0))
	mBase = m.M
	m.G0 = v415 + int32(32)
	goto L85
L91:
	;
	v444 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v444)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v9 + int32(12)
	goto L94
L92:
	;
	v451 = v42
	v453 = int32(0)
	goto L9
L94:
	;
	goto L92
L95:
	;
	v454 = int32(_a_F_BackgroundWorkerMain_5)
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[23])) = v456 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[24])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[25])) = v9 + int32(16)
	F_InitProcess(m)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	F_EmitErrorReport(m)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	F_proc_exit(m, int32(1))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	goto L3
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	F_BaseInit(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	v485 = v451 + int32(1228)
	v486 = m.G0
	v488 = v486 - int32(16)
	m.G0 = v488
	v491 = v451 + int32(204)
	v492 = int32(_a_F_BackgroundWorkerMain_6)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[26])))
	if base.B2i32(v495 == int32(0))|base.B2i32(v495 != v498) != 0 {
		v516 = v495
		v517 = v498
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v451)+1324))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	m.T0[v678].(func(*base.Module, int32))(m, v695)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L168
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L165
	}
L105:
	;
	m.G0 = v488 + int32(16)
	goto L103
L106:
	;
	if v516-v517 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L107:
	;
	goto L106
L108:
	;
	v501 = v491
	v502 = v492
	goto L109
L109:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+1)))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+1)))
	if v506 == int32(0) {
		v516 = v506
		v517 = v505
		goto L107
	} else {
		goto L111
	}
L110:
	;
	v516 = v506
	v517 = v505
	goto L107
L111:
	;
	v509 = int32(1)
	if v506 == v505 {
		v501 = v501 + v509
		v502 = v502 + v509
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v521 = int32(_a_F_BackgroundWorkerMain_7)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[27])))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if base.B2i32(v524 == int32(0))|base.B2i32(v524 != v527) != 0 {
		v545 = v524
		v546 = v527
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	v676 = F_load_external_function(m, v491, v485, int32(1), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L164
	}
L116:
	;
	if v545-v546 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	v530 = v521
	v531 = v485
	goto L119
L119:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+1)))
	if v535 == int32(0) {
		v545 = v535
		v546 = v534
		goto L117
	} else {
		goto L121
	}
L120:
	;
	v545 = v535
	v546 = v534
	goto L117
L121:
	;
	v538 = int32(1)
	if v535 == v534 {
		v530 = v530 + v538
		v531 = v531 + v538
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[28]))
	v678 = v551
	goto L105
L124:
	;
	goto L125
L125:
	;
	v552 = int32(_a_F_BackgroundWorkerMain_8)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[29])))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if base.B2i32(v555 == int32(0))|base.B2i32(v555 != v558) != 0 {
		v576 = v555
		v577 = v558
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v576-v577 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	goto L126
L128:
	;
	v561 = v552
	v562 = v485
	goto L129
L129:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+1)))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+1)))
	if v566 == int32(0) {
		v576 = v566
		v577 = v565
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v576 = v566
	v577 = v565
	goto L127
L131:
	;
	v569 = int32(1)
	if v566 == v565 {
		v561 = v561 + v569
		v562 = v562 + v569
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[30]))
	v678 = v582
	goto L105
L134:
	;
	goto L135
L135:
	;
	v583 = int32(_a_F_BackgroundWorkerMain_9)
	v586 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[31])))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if base.B2i32(v586 == int32(0))|base.B2i32(v586 != v589) != 0 {
		v607 = v586
		v608 = v589
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v607-v608 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L137:
	;
	goto L136
L138:
	;
	v592 = v583
	v593 = v485
	goto L139
L139:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+1)))
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592)+1)))
	if v597 == int32(0) {
		v607 = v597
		v608 = v596
		goto L137
	} else {
		goto L141
	}
L140:
	;
	v607 = v597
	v608 = v596
	goto L137
L141:
	;
	v600 = int32(1)
	if v597 == v596 {
		v592 = v592 + v600
		v593 = v593 + v600
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[32]))
	v678 = v613
	goto L105
L144:
	;
	goto L145
L145:
	;
	v614 = int32(_a_F_BackgroundWorkerMain_10)
	v617 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[33])))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if base.B2i32(v617 == int32(0))|base.B2i32(v617 != v620) != 0 {
		v638 = v617
		v639 = v620
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v638-v639 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L147:
	;
	goto L146
L148:
	;
	v623 = v614
	v624 = v485
	goto L149
L149:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+1)))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+1)))
	if v628 == int32(0) {
		v638 = v628
		v639 = v627
		goto L147
	} else {
		goto L151
	}
L150:
	;
	v638 = v628
	v639 = v627
	goto L147
L151:
	;
	v631 = int32(1)
	if v628 == v627 {
		v623 = v623 + v631
		v624 = v624 + v631
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[34]))
	v678 = v644
	goto L105
L154:
	;
	goto L155
L155:
	;
	v645 = int32(_a_F_BackgroundWorkerMain_11)
	v648 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[35])))
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if base.B2i32(v648 == int32(0))|base.B2i32(v648 != v651) != 0 {
		v669 = v648
		v670 = v651
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if v669-v670 != 0 {
		goto L104
	} else {
		goto L163
	}
L157:
	;
	goto L156
L158:
	;
	v654 = v645
	v655 = v485
	goto L159
L159:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+1)))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654)+1)))
	if v659 == int32(0) {
		v669 = v659
		v670 = v658
		goto L157
	} else {
		goto L161
	}
L160:
	;
	v669 = v659
	v670 = v658
	goto L157
L161:
	;
	v662 = int32(1)
	if v659 == v658 {
		v654 = v654 + v662
		v655 = v655 + v662
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerMain[36]))
	v678 = v673
	goto L105
L164:
	;
	v678 = v676
	goto L105
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v485
	F_errmsg_internal(m, int32(_a_F_BackgroundWorkerMain_12), v488)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_BackgroundWorkerMain_1), int32(1355), int32(_a_F_BackgroundWorkerMain_13))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v451
	F_proc_exit(m, int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L5
L170:
	;
	v712 = int32(v708)
	m.G0 = v9
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	if v9+int32(12) == v718 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	m.ExcPending = 1
	goto L179
L172:
	;
	if v722 != 0 {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	v722 = v720
	goto L175
L174:
	;
	v722 = int32(0)
	goto L175
L175:
	;
	goto L172
L176:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v14 = v723
	v15 = v722
	v16 = v714
	goto L1
L177:
	;
	goto L178
L178:
	;
	F___wasm_longjmp(m, v715, v714)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	return
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BackgroundWorkerUnblockSignals(m *base.Module) {
	var v4 int32
	_ = v4
	F_pgmem_sigprocmask(m, int32(_a_F_BackgroundWorkerUnblockSignals_0), int32(0))
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
	v5 = Fn13822(m, l0, int32(_a_F_BogusGetChunkSpace_0), int32(315), int32(_a_F_BogusGetChunkSpace_1))
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
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v738 int64
	_ = v738
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 float64
	_ = v941
	var v943 float64
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 float64
	_ = v960
	var v961 int32
	_ = v961
	var v979 float64
	_ = v979
	var v981 int32
	_ = v981
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
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
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1050 float64
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1114 int64
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1169 int64
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int64
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1236 int32
	_ = v1236
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1270 int64
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int64
	_ = v1301
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
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
	v124 = Fn13966(m, int64(32))
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
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L311
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L308
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
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
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v862 != 0 {
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
	v730 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v729))), uint32(v730))
	*(*int64)(unsafe.Add(mBase, uint32(v729)+4)) = int64(-1)
	goto L170
L170:
	;
	v735 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v718)+28)), uint32(v735))
	v738 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v718)+40)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v718)+32)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v718)+48)) = v738
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v744, v718-int32(-64), v638)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v750 = F_shm_toc_allocate(m, v749, v653)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v630)+44))
	F_tuplesort_initialize_shared(m, v750, v614, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v756 = F_shm_toc_allocate(m, v755, v666)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
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
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v758)+16))
	base.MemoryCopy(m, v756, v759, v666)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v761, int64(-6917529027641081855), v718)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v765, int64(-6917529027641081854), v750)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v769, int64(-6917529027641081853), v756)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	if v774 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	v776 = F_shm_toc_allocate(m, v775, v699)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
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
	v787 = m.ExcPending
	if v787 != 0 {
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
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	base.MemoryCopy(m, v776, v779, v699)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v630)+52))
	F_shm_toc_insert(m, v781, int64(-6917529027641081852), v776)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
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
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v618)+20)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v618)+16)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v618)+12)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v618)+8)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v618)+4)) = v789 + int32(1)
	if v789 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_WaitForParallelWorkersToFinish(m, v630)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v817 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L199
	}
L193:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v618)+16))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	switch v802 {
	case 0, 5:
		goto L195
	default:
		goto L194
	}
L194:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	F_DestroyParallelContext(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	F_UnregisterSnapshot(m, v801)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+72)) = v811 - int32(1)
	goto L198
L198:
	;
	goto L136
L199:
	;
	if v817 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v819
	F_errmsg(m, int32(_a_F_BuildIndex_2_11), v19+int32(-32))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
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
	v833 = F_palloc0(m, int32(12))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(955), int32(_a_F_BuildIndex_2_12))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v833)+4)) = v835
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v833)+8)) = v837
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v618)+8))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v618)+12))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v618)+20))
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	v845 = base.I32_div_s(v843, v844)
	F_IvfflatParallelScanAndSort(m, v833, v839, v840, v841, v845, int32(1))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_WaitForParallelWorkersToAttach(m, v630)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	goto L136
L208:
	;
	v864 = F_palloc0(m, int32(12))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v874 = int32(0)
	goto L210
L210:
	;
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v878 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+52)) = uint16(v878)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = int32(97)
	v882 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v882
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+48)) = uint8(v882)
	v896 = F_tuplesort_begin_heap(m, v877, v878, v19+int32(-12), v19+int32(-4), v19+int32(-8), v19+int32(-16), v876, v874, v882)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	v866 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v864))) = uint8(v866)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v864)+4)) = v869
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v868)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v864)+8)) = v871
	v874 = v864
	goto L210
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+152)) = v896
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v899 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v900 != 0 {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v1000 = v896
	goto L215
L215:
	;
	F_tuplesort_performsort(m, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L232
	}
L216:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v979
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	v1000 = v981
	goto L215
L217:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+8))
	v905 = v901 + int32(28)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
	goto L220
L218:
	;
	goto L219
L219:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v951 = int32(1)
	v952 = int32(0)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v899)+188))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)+140))
	v960 = m.T0[v959].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v899, v949, v950, v951, v952, v951, v952, int32(-1), int32(_a_F_BuildIndex_2_13), l3, v952)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L231
	}
L220:
	;
	v927 = base.AtomicRmwXchg32(m, v905, int32(0), int32(1))
	if v927 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v901)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v941
	v943 = *(*float64)(unsafe.Add(mBase, uint32(v901)+40))
	v944 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v901)+28)), uint32(v944))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L230
	}
L222:
	;
	F_s_lock(m, v905, int32(_a_F_BuildIndex_2_6), int32(630), int32(_a_F_BuildIndex_2_14))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v901)+32))
	if v906 != v933 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	v935 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v905))), uint32(v935))
	F_ConditionVariableSleep(m, v901+int32(16), int32(134217767))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
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
	v979 = v943
	goto L216
L231:
	;
	v979 = v960
	goto L216
L232:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1004 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v1004
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v1009 = F_MakeTupleTableSlot(m, v1007, int32(_a_F_BuildIndex_2_15))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1016 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1016 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1050 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1054 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	goto L234
L236:
	;
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1020&int32(1) == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1025 = int32(_a_F_BuildIndex_2_0)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1028 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1027 + v1028
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1031 + v1028
	*(*int64)(unsafe.Add(mBase, uint32(v1016+int32(80))+232)) = int64(4)
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	*(*int32)(unsafe.Add(mBase, uint32(v1016))) = v1039 + v1028
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1045 - v1028
	goto L235
L238:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1087, v1011, v1009, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L242
	}
L239:
	;
	goto L238
L240:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1058&int32(1) == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1063 = int32(_a_F_BuildIndex_2_0)
	v1065 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1066 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1065 + v1066
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	*(*int32)(unsafe.Add(mBase, uint32(v1054))) = v1069 + v1066
	*(*int64)(unsafe.Add(mBase, uint32(v1054+int32(88))+232)) = base.I64_trunc_sat_f64_s(v1050)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	*(*int32)(unsafe.Add(mBase, uint32(v1054))) = v1077 + v1066
	v1083 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1083 - v1066
	goto L239
L242:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)))
	if int32(0) < v1095 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1101 = v1004
	v1114 = int64(0)
	goto L246
L244:
	;
	goto L245
L245:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_tuplesort_end(m, v1331)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L287
	}
L246:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[8]))
	if v1118 != 0 {
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
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1121 = F_HnswNewBuffer(m, v1003, l4)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L252
	}
L251:
	;
	goto L250
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v1121
	F_IvfflatInitRegisterPage(m, v1003, v19+int32(-12), v19+int32(-16), v19+int32(-20))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1132 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1101 == v1152 {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1136+(v1132^int32(-1))<<(uint(int32(6))%32))+16))
	v1151 = v1142
	goto L254
L256:
	;
	goto L257
L257:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1144+v1132<<(uint(int32(6))%32)+int32(-64))+16))
	v1151 = v1150
	goto L254
L258:
	;
	v1169 = v1114
	goto L261
L259:
	;
	v1270 = v1114
	goto L260
L260:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1273 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L261:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v1173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1172)+6)))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1175 = int32(4)
	v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174)+14)))
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174)+12)))
	v1178 = v1176 - v1177
	if v1178 <= v1175 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1270 = v1210
	goto L260
L263:
	;
	v1189 = (v1173&int32(_a_F_BuildIndex_2_16) + int32(7)) & int32(_a_F_BuildIndex_2_17)
	if base.Ui32(v1181-int32(4)) < base.Ui32(v1189) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1181 = v1175
	goto L266
L265:
	;
	v1181 = v1178
	goto L266
L266:
	;
	goto L263
L267:
	;
	F_IvfflatAppendPage(m, v1003, v19+int32(-12), v19+int32(-16), v19+int32(-20), l4)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1200 = int32(0)
	v1202 = F_PageAddItemExtended(m, v1199, v1172, v1189, v1200, v1200)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	if v1202 == int32(0) {
		goto L82
	} else {
		goto L272
	}
L272:
	;
	F_pfree(m, v1172)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1210 = v1169 + int64(1)
	v1213 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1213 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1246, v1011, v1009, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L278
	}
L275:
	;
	goto L274
L276:
	;
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1217&int32(1) == int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1222 = int32(_a_F_BuildIndex_2_0)
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1225 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1224 + v1225
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	*(*int32)(unsafe.Add(mBase, uint32(v1213))) = v1228 + v1225
	*(*int64)(unsafe.Add(mBase, uint32(v1213+int32(96))+232)) = v1210
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	*(*int32)(unsafe.Add(mBase, uint32(v1213))) = v1236 + v1225
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1242 - v1225
	goto L275
L278:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1253 == v1101 {
		v1169 = v1210
		goto L261
	} else {
		goto L279
	}
L279:
	;
	goto L262
L280:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	F_IvfflatCommitBuffer(m, v1293, v1294)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L284
	}
L281:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1277+(v1273^int32(-1))<<(uint(int32(6))%32))+16))
	v1292 = v1283
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1285+v1273<<(uint(int32(6))%32)+int32(-64))+16))
	v1292 = v1291
	goto L280
L284:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v1301 = *(*int64)(unsafe.Add(mBase, uint32(v1297+v1101<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v1301
	F_IvfflatUpdateList(m, v1003, v19+int32(-40), v1292, int32(-1), v1151, l4)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1309 = v1101 + int32(1)
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1310)))
	if v1309 < v1311 {
		v1101 = v1309
		v1114 = v1270
		goto L246
	} else {
		goto L286
	}
L286:
	;
	goto L247
L287:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v1334 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	F_WaitForParallelWorkersToFinish(m, v1335)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
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
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1334)+16))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	switch v1339 {
	case 0, 5:
		goto L293
	default:
		goto L292
	}
L292:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	F_DestroyParallelContext(m, v1342)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	F_UnregisterSnapshot(m, v1338)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+72)) = v1348 - int32(1)
	goto L296
L296:
	;
	goto L290
L297:
	;
	v1355 = int32(3)
	v1357 = F_RelationGetNumberOfBlocksInFork(m, l1, v1355)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	F_VectorArrayFree(m, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	F_log_newpage_range(m, l1, v1355, v1357, int32(1))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	F_pfree(m, v1365)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	F_MemoryContextDelete(m, v1368)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
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
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_19), int32(326), int32(_a_F_BuildIndex_2_20))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
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
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1391 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_21), v21)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(546), int32(_a_F_BuildIndex_2_22))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
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
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1407 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_21), v19+int32(-48))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_6), int32(315), int32(_a_F_BuildIndex_2_23))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
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
	v3 = Fn13855(m, l0, int32(1042))
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
	v5 = Fn13859(m, l0, l1, int32(26), int32(2))
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
