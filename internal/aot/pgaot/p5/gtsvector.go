package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v239 int64
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v273 int64
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v430 int64
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v437 int64
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v440 int32
	_ = v440
	var v441 int64
	_ = v441
	var v445 int64
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v464 int64
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v480 int64
	_ = v480
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v525 int64
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v532 int64
	_ = v532
	var v533 int32
	_ = v533
	var v534 int64
	_ = v534
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v540 int64
	_ = v540
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v559 int64
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int64
	_ = v659
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v674 int64
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int64
	_ = v681
	var v682 int32
	_ = v682
	var v683 int64
	_ = v683
	var v684 int32
	_ = v684
	var v685 int64
	_ = v685
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v691 int64
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int64
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int64
	_ = v706
	var v710 int32
	_ = v710
	var v711 int64
	_ = v711
	var v712 int64
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int64
	_ = v720
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v746 int64
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v754 int32
	_ = v754
	var v755 int64
	_ = v755
	var v756 int32
	_ = v756
	var v757 int64
	_ = v757
	var v758 int32
	_ = v758
	var v759 int64
	_ = v759
	var v763 int64
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int64
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int64
	_ = v778
	var v782 int32
	_ = v782
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v792 int64
	_ = v792
	var v802 int64
	_ = v802
	var v803 int64
	_ = v803
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v818 int64
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int64
	_ = v825
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v828 int32
	_ = v828
	var v829 int64
	_ = v829
	var v830 int32
	_ = v830
	var v831 int64
	_ = v831
	var v835 int64
	_ = v835
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v843 int64
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int64
	_ = v850
	var v854 int32
	_ = v854
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int64
	_ = v864
	var v874 int64
	_ = v874
	var v888 int64
	_ = v888
	var v905 int64
	_ = v905
	var v910 int32
	_ = v910
	var v940 int64
	_ = v940
	var v964 float32
	_ = v964
	var v967 int32
	_ = v967
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = base.I32_wrap_i64(v18)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == v2 {
		v38 = v2
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
		if v25 == int32(0) {
			v38 = v2
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			if v28 != int32(7) {
				v38 = v2
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				if v31 != int32(17) {
					v38 = v2
				} else {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
					v38 = v34 ^ int32(1)
				}
			}
		}
	}
	if v38&int32(1) != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = F_get_fn_opclass_options(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
			v47 = v46
			v49 = v18 & int64(4294967295)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
			v55 = v51 + int32(8)
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
			if v56&int32(1) != 0 {
				v59 = F_palloc(m, v47)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int64(0)
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v62 = int32(2)
					v67 = int32(base.Ui32(int32(base.Ui32(v61)>>(uint(v62)%32))-int32(8)) >> (uint(v62) % 32))
					v68 = int32(3)
					if v47&v68|(v59&v68|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v47))) == int32(0) {
						if v47 == int32(0) {
						} else {
							v82 = v59 + v47
							v84 = v59 + int32(4)
							if base.Ui32(v84) < base.Ui32(v82) {
								v86 = v82
							} else {
								v86 = v84
							}
							v92 = (v59^int32(-1)+v86)&int32(-4) + int32(4)
							if v92 == int32(0) {
							} else {
								base.MemoryFill(m, v59, int32(0), v92)
							}
						}
					} else {
						v92 = v47
						if v92 == int32(0) {
						} else {
							base.MemoryFill(m, v59, int32(0), v92)
						}
					}
					if v67 == int32(0) {
					} else {
						v103 = v50 + int32(8)
						v105 = v47 << (uint(int32(3)) % 32)
						v106 = int32(0)
						if v67 != int32(1) {
							v115 = v106
							v119 = int32(0)
							for {
								v129 = int32(2)
								v131 = v103 + v115<<(uint(v129)%32)
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
								v133 = base.I32_rem_u_s(v132, v105)
								v134 = int32(3)
								v136 = v59 + int32(base.Ui32(v133)>>(uint(v134)%32))
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
								v138 = int32(1)
								v139 = int32(7)
								v142 = v137 | v138<<(uint(v133&v139)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v142)
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
								v145 = base.I32_rem_u_s(v144, v105)
								v148 = v59 + int32(base.Ui32(v145)>>(uint(v134)%32))
								v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
								v154 = v149 | v138<<(uint(v145&v139)%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v154)
								v157 = v115 + v129
								v159 = v119 + v129
								if v159 != v67&int32(1073741822) {
									v115 = v157
									v119 = v159
									continue
								} else {
									break
								}
								break
							}
							if v67&int32(1) == int32(0) {
							} else {
								v164 = v157
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
								v182 = base.I32_rem_u_s(v181, v105)
								v185 = v59 + int32(base.Ui32(v182)>>(uint(int32(3))%32))
								v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
								v191 = v186 | int32(1)<<(uint(v182&int32(7))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v191)
							}
						} else {
							v164 = v106
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
							v182 = base.I32_rem_u_s(v181, v105)
							v185 = v59 + int32(base.Ui32(v182)>>(uint(int32(3))%32))
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
							v191 = v186 | int32(1)<<(uint(v182&int32(7))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v191)
						}
					}
					v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
					if v208&int32(4) != 0 {
						v212 = v47 << (uint(int32(3)) % 32)
						if int32(7) < v47 {
							v659 = int64(0)
							v660 = int32(0)
							if v47 == v660 {
								v730 = int64(0)
							} else {
								v667 = v47 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v47) {
									v672 = v59
									v674 = v659
									v677 = v660
									for {
										v678 = int32(4)
										v679 = v672 + v678
										v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+3)))
										v681 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v680)+uint32(_c_F_gtsvector_penalty[0]))))
										v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+2)))
										v683 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v682)+uint32(_c_F_gtsvector_penalty[0]))))
										v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)))
										v685 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v684)+uint32(_c_F_gtsvector_penalty[0]))))
										v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
										v687 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v686)+uint32(_c_F_gtsvector_penalty[0]))))
										v691 = v681 + (v683 + (v685 + (v674 + v687)))
										v693 = v677 + v678
										if v693 != v47&int32(-4) {
											v672 = v679
											v674 = v691
											v677 = v693
											continue
										} else {
											break
										}
										break
									}
									if v667 == int32(0) {
										v720 = v691
									} else {
										v697 = v679
										v699 = v691
										v704 = v697
										v705 = int32(0)
										v706 = v699
										for {
											v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
											v711 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v710)+uint32(_c_F_gtsvector_penalty[0]))))
											v712 = v706 + v711
											v713 = int32(1)
											v716 = v705 + v713
											if v716 != v667 {
												v704 = v704 + v713
												v705 = v716
												v706 = v712
												continue
											} else {
												break
											}
											break
										}
										v720 = v712
									}
								} else {
									v697 = v59
									v699 = v659
									v704 = v697
									v705 = int32(0)
									v706 = v699
									for {
										v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
										v711 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v710)+uint32(_c_F_gtsvector_penalty[0]))))
										v712 = v706 + v711
										v713 = int32(1)
										v716 = v705 + v713
										if v716 != v667 {
											v704 = v704 + v713
											v705 = v716
											v706 = v712
											continue
										} else {
											break
										}
										break
									}
									v720 = v712
								}
								v730 = v720
							}
							v940 = v730
						} else {
							if v47 == int32(0) {
								v940 = int64(0)
							} else {
								v219 = v47 & int32(3)
								v220 = int64(0)
								if base.Ui32(int32(4)) <= base.Ui32(v47) {
									v227 = v59
									v229 = int32(0)
									v239 = v220
									for {
										v241 = int32(4)
										v242 = v227 + v241
										v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+3)))
										v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_gtsvector_penalty[0]))))
										v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+2)))
										v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_gtsvector_penalty[0]))))
										v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
										v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_gtsvector_penalty[0]))))
										v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
										v250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_gtsvector_penalty[0]))))
										v254 = v244 + (v246 + (v248 + (v239 + v250)))
										v256 = v229 + v241
										if v256 != v47&int32(-4) {
											v227 = v242
											v229 = v256
											v239 = v254
											continue
										} else {
											break
										}
										break
									}
									if v219 == int32(0) {
										v940 = v254
									} else {
										v261 = v242
										v273 = v254
										v277 = v261
										v280 = int32(0)
										v289 = v273
										for {
											v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
											v292 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+uint32(_c_F_gtsvector_penalty[0]))))
											v293 = v289 + v292
											v294 = int32(1)
											v297 = v280 + v294
											if v297 != v219 {
												v277 = v277 + v294
												v280 = v297
												v289 = v293
												continue
											} else {
												break
											}
											break
										}
										v940 = v293
									}
								} else {
									v261 = v59
									v273 = v220
									v277 = v261
									v280 = int32(0)
									v289 = v273
									for {
										v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
										v292 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+uint32(_c_F_gtsvector_penalty[0]))))
										v293 = v289 + v292
										v294 = int32(1)
										v297 = v280 + v294
										if v297 != v219 {
											v277 = v277 + v294
											v280 = v297
											v289 = v293
											continue
										} else {
											break
										}
										break
									}
									v940 = v293
								}
							}
						}
						v964 = base.F32_div(base.F32_convert_i32_s(v212-base.I32_wrap_i64(v940)), base.F32_convert_i32_s(v212|int32(1)))
					} else {
						if v47 <= int32(0) {
							v964 = float32(0)
						} else {
							v302 = int32(0)
							if v47 != int32(1) {
								v312 = v302
								v314 = v302
								v316 = int32(0)
								for {
									v327 = v312 | int32(1)
									v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v327))))
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v327))))
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329^v331)+uint32(_c_F_gtsvector_penalty[0]))))
									v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312+v55))))
									v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v312))))
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335^v337)+uint32(_c_F_gtsvector_penalty[0]))))
									v341 = v333 + (v314 + v339)
									v342 = int32(2)
									v343 = v312 + v342
									v345 = v316 + v342
									if v345 != v47&int32(2147483646) {
										v312 = v343
										v314 = v341
										v316 = v345
										continue
									} else {
										break
									}
									break
								}
								if v47&int32(1) == int32(0) {
									v374 = v341
								} else {
									v350 = v343
									v352 = v341
									v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v55))))
									v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v350))))
									v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365^v367)+uint32(_c_F_gtsvector_penalty[0]))))
									v374 = v352 + v369
								}
							} else {
								v350 = v302
								v352 = v302
								v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v55))))
								v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v350))))
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365^v367)+uint32(_c_F_gtsvector_penalty[0]))))
								v374 = v352 + v369
							}
							v964 = base.F32_convert_i32_s(v374)
						}
					}
					*(*float32)(unsafe.Add(mBase, uint32(v19))) = v964
					F_pfree(m, v59)
					mBase = m.M
					v967 = m.ExcPending
					if v967 != 0 {
						return int64(0)
					} else {
						return v49
					}
				}
			} else {
				v387 = int32(4)
				v388 = v56 & v387
				v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
				if v389&v387 != 0 {
					if v388 != 0 {
						v910 = int32(0)
					} else {
						v393 = int32(8)
						v394 = v50 + v393
						v395 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						v397 = int32(base.Ui32(v395) >> (uint(int32(2)) % 32))
						v399 = v397 - v393
						if base.Ui32(int32(63)) < base.Ui32(v395) {
							v731 = int64(0)
							v732 = int32(0)
							if v399 == v732 {
								v802 = int64(0)
							} else {
								v739 = v399 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v399) {
									v744 = v394
									v746 = v731
									v749 = v732
									for {
										v750 = int32(4)
										v751 = v744 + v750
										v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
										v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v752)+uint32(_c_F_gtsvector_penalty[0]))))
										v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
										v755 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v754)+uint32(_c_F_gtsvector_penalty[0]))))
										v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
										v757 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v756)+uint32(_c_F_gtsvector_penalty[0]))))
										v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
										v759 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_gtsvector_penalty[0]))))
										v763 = v753 + (v755 + (v757 + (v746 + v759)))
										v765 = v749 + v750
										if v765 != v399&int32(-4) {
											v744 = v751
											v746 = v763
											v749 = v765
											continue
										} else {
											break
										}
										break
									}
									if v739 == int32(0) {
										v792 = v763
									} else {
										v769 = v751
										v771 = v763
										v776 = v769
										v777 = int32(0)
										v778 = v771
										for {
											v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
											v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
											v784 = v778 + v783
											v785 = int32(1)
											v788 = v777 + v785
											if v788 != v739 {
												v776 = v776 + v785
												v777 = v788
												v778 = v784
												continue
											} else {
												break
											}
											break
										}
										v792 = v784
									}
								} else {
									v769 = v394
									v771 = v731
									v776 = v769
									v777 = int32(0)
									v778 = v771
									for {
										v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
										v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
										v784 = v778 + v783
										v785 = int32(1)
										v788 = v777 + v785
										if v788 != v739 {
											v776 = v776 + v785
											v777 = v788
											v778 = v784
											continue
										} else {
											break
										}
										break
									}
									v792 = v784
								}
								v802 = v792
							}
							v905 = v802
						} else {
							if v399 == int32(0) {
								v905 = int64(0)
							} else {
								v407 = int32(3)
								v408 = v397 & v407
								v409 = int64(0)
								if base.Ui32(v407) <= base.Ui32(v397-int32(9)) {
									v417 = v394
									v421 = int32(0)
									v430 = v409
									for {
										v432 = int32(4)
										v433 = v417 + v432
										v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+3)))
										v435 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_gtsvector_penalty[0]))))
										v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+2)))
										v437 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v436)+uint32(_c_F_gtsvector_penalty[0]))))
										v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)))
										v439 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v438)+uint32(_c_F_gtsvector_penalty[0]))))
										v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
										v441 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v440)+uint32(_c_F_gtsvector_penalty[0]))))
										v445 = v435 + (v437 + (v439 + (v430 + v441)))
										v447 = v421 + v432
										if v447 != v399&int32(-4) {
											v417 = v433
											v421 = v447
											v430 = v445
											continue
										} else {
											break
										}
										break
									}
									if v408 == int32(0) {
										v905 = v445
									} else {
										v451 = v433
										v464 = v445
										v467 = v451
										v468 = int32(0)
										v480 = v464
										for {
											v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
											v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_gtsvector_penalty[0]))))
											v484 = v480 + v483
											v485 = int32(1)
											v488 = v468 + v485
											if v488 != v408 {
												v467 = v467 + v485
												v468 = v488
												v480 = v484
												continue
											} else {
												break
											}
											break
										}
										v905 = v484
									}
								} else {
									v451 = v394
									v464 = v409
									v467 = v451
									v468 = int32(0)
									v480 = v464
									for {
										v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
										v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_gtsvector_penalty[0]))))
										v484 = v480 + v483
										v485 = int32(1)
										v488 = v468 + v485
										if v488 != v408 {
											v467 = v467 + v485
											v468 = v488
											v480 = v484
											continue
										} else {
											break
										}
										break
									}
									v905 = v484
								}
							}
						}
						v910 = v399<<(uint(int32(3))%32) - base.I32_wrap_i64(v905)
					}
				} else {
					v490 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
					v492 = int32(base.Ui32(v490) >> (uint(int32(2)) % 32))
					v494 = v492 - int32(8)
					if v388 != 0 {
						if base.Ui32(int32(63)) < base.Ui32(v490) {
							v803 = int64(0)
							v804 = int32(0)
							if v494 == v804 {
								v874 = int64(0)
							} else {
								v811 = v494 & int32(3)
								if base.Ui32(int32(4)) <= base.Ui32(v494) {
									v816 = v55
									v818 = v803
									v821 = v804
									for {
										v822 = int32(4)
										v823 = v816 + v822
										v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+3)))
										v825 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v824)+uint32(_c_F_gtsvector_penalty[0]))))
										v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+2)))
										v827 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v826)+uint32(_c_F_gtsvector_penalty[0]))))
										v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
										v829 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v828)+uint32(_c_F_gtsvector_penalty[0]))))
										v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
										v831 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtsvector_penalty[0]))))
										v835 = v825 + (v827 + (v829 + (v818 + v831)))
										v837 = v821 + v822
										if v837 != v494&int32(-4) {
											v816 = v823
											v818 = v835
											v821 = v837
											continue
										} else {
											break
										}
										break
									}
									if v811 == int32(0) {
										v864 = v835
									} else {
										v841 = v823
										v843 = v835
										v848 = v841
										v849 = int32(0)
										v850 = v843
										for {
											v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
											v855 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_gtsvector_penalty[0]))))
											v856 = v850 + v855
											v857 = int32(1)
											v860 = v849 + v857
											if v860 != v811 {
												v848 = v848 + v857
												v849 = v860
												v850 = v856
												continue
											} else {
												break
											}
											break
										}
										v864 = v856
									}
								} else {
									v841 = v55
									v843 = v803
									v848 = v841
									v849 = int32(0)
									v850 = v843
									for {
										v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
										v855 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_gtsvector_penalty[0]))))
										v856 = v850 + v855
										v857 = int32(1)
										v860 = v849 + v857
										if v860 != v811 {
											v848 = v848 + v857
											v849 = v860
											v850 = v856
											continue
										} else {
											break
										}
										break
									}
									v864 = v856
								}
								v874 = v864
							}
							v888 = v874
						} else {
							if v494 == int32(0) {
								v888 = int64(0)
							} else {
								v502 = int32(3)
								v503 = v492 & v502
								v504 = int64(0)
								if base.Ui32(v502) <= base.Ui32(v492-int32(9)) {
									v513 = int32(0)
									v516 = v55
									v525 = v504
									for {
										v527 = int32(4)
										v528 = v516 + v527
										v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
										v530 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v529)+uint32(_c_F_gtsvector_penalty[0]))))
										v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
										v532 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v531)+uint32(_c_F_gtsvector_penalty[0]))))
										v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
										v534 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v533)+uint32(_c_F_gtsvector_penalty[0]))))
										v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
										v536 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v535)+uint32(_c_F_gtsvector_penalty[0]))))
										v540 = v530 + (v532 + (v534 + (v525 + v536)))
										v542 = v513 + v527
										if v542 != v494&int32(-4) {
											v513 = v542
											v516 = v528
											v525 = v540
											continue
										} else {
											break
										}
										break
									}
									if v503 == int32(0) {
										v888 = v540
									} else {
										v550 = v528
										v559 = v540
										v562 = int32(0)
										v566 = v550
										v575 = v559
										for {
											v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
											v578 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_c_F_gtsvector_penalty[0]))))
											v579 = v575 + v578
											v580 = int32(1)
											v583 = v562 + v580
											if v583 != v503 {
												v562 = v583
												v566 = v566 + v580
												v575 = v579
												continue
											} else {
												break
											}
											break
										}
										v888 = v579
									}
								} else {
									v550 = v55
									v559 = v504
									v562 = int32(0)
									v566 = v550
									v575 = v559
									for {
										v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
										v578 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_c_F_gtsvector_penalty[0]))))
										v579 = v575 + v578
										v580 = int32(1)
										v583 = v562 + v580
										if v583 != v503 {
											v562 = v583
											v566 = v566 + v580
											v575 = v579
											continue
										} else {
											break
										}
										break
									}
									v888 = v579
								}
							}
						}
						v910 = v494<<(uint(int32(3))%32) - base.I32_wrap_i64(v888)
					} else {
						if base.Ui32(v490) < base.Ui32(int32(36)) {
							v910 = int32(0)
						} else {
							v589 = v50 + int32(8)
							v590 = int32(0)
							if v492 != int32(9) {
								v599 = v590
								v600 = v590
								v604 = int32(0)
								for {
									v615 = v599 | int32(1)
									v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v615))))
									v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v615))))
									v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617^v619)+uint32(_c_F_gtsvector_penalty[0]))))
									v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v589))))
									v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v55))))
									v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623^v625)+uint32(_c_F_gtsvector_penalty[0]))))
									v629 = v621 + (v600 + v627)
									v630 = int32(2)
									v631 = v599 + v630
									v633 = v604 + v630
									if v633 != v494&int32(-2) {
										v599 = v631
										v600 = v629
										v604 = v633
										continue
									} else {
										break
									}
									break
								}
								if v492&int32(1) == int32(0) {
									v910 = v629
								} else {
									v637 = v631
									v638 = v629
									v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v589))))
									v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v55))))
									v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653^v655)+uint32(_c_F_gtsvector_penalty[0]))))
									v910 = v638 + v657
								}
							} else {
								v637 = v590
								v638 = v590
								v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v589))))
								v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v55))))
								v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653^v655)+uint32(_c_F_gtsvector_penalty[0]))))
								v910 = v638 + v657
							}
						}
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v19))) = base.F32_convert_i32_s(v910)
				return v49
			}
		}
	} else {
		v47 = int32(124)
		v49 = v18 & int64(4294967295)
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
		v55 = v51 + int32(8)
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
		if v56&int32(1) != 0 {
			v59 = F_palloc(m, v47)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int64(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v62 = int32(2)
				v67 = int32(base.Ui32(int32(base.Ui32(v61)>>(uint(v62)%32))-int32(8)) >> (uint(v62) % 32))
				v68 = int32(3)
				if v47&v68|(v59&v68|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v47))) == int32(0) {
					if v47 == int32(0) {
					} else {
						v82 = v59 + v47
						v84 = v59 + int32(4)
						if base.Ui32(v84) < base.Ui32(v82) {
							v86 = v82
						} else {
							v86 = v84
						}
						v92 = (v59^int32(-1)+v86)&int32(-4) + int32(4)
						if v92 == int32(0) {
						} else {
							base.MemoryFill(m, v59, int32(0), v92)
						}
					}
				} else {
					v92 = v47
					if v92 == int32(0) {
					} else {
						base.MemoryFill(m, v59, int32(0), v92)
					}
				}
				if v67 == int32(0) {
				} else {
					v103 = v50 + int32(8)
					v105 = v47 << (uint(int32(3)) % 32)
					v106 = int32(0)
					if v67 != int32(1) {
						v115 = v106
						v119 = int32(0)
						for {
							v129 = int32(2)
							v131 = v103 + v115<<(uint(v129)%32)
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
							v133 = base.I32_rem_u_s(v132, v105)
							v134 = int32(3)
							v136 = v59 + int32(base.Ui32(v133)>>(uint(v134)%32))
							v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
							v138 = int32(1)
							v139 = int32(7)
							v142 = v137 | v138<<(uint(v133&v139)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v142)
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
							v145 = base.I32_rem_u_s(v144, v105)
							v148 = v59 + int32(base.Ui32(v145)>>(uint(v134)%32))
							v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
							v154 = v149 | v138<<(uint(v145&v139)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v154)
							v157 = v115 + v129
							v159 = v119 + v129
							if v159 != v67&int32(1073741822) {
								v115 = v157
								v119 = v159
								continue
							} else {
								break
							}
							break
						}
						if v67&int32(1) == int32(0) {
						} else {
							v164 = v157
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
							v182 = base.I32_rem_u_s(v181, v105)
							v185 = v59 + int32(base.Ui32(v182)>>(uint(int32(3))%32))
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
							v191 = v186 | int32(1)<<(uint(v182&int32(7))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v191)
						}
					} else {
						v164 = v106
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
						v182 = base.I32_rem_u_s(v181, v105)
						v185 = v59 + int32(base.Ui32(v182)>>(uint(int32(3))%32))
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
						v191 = v186 | int32(1)<<(uint(v182&int32(7))%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v191)
					}
				}
				v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
				if v208&int32(4) != 0 {
					v212 = v47 << (uint(int32(3)) % 32)
					if int32(7) < v47 {
						v659 = int64(0)
						v660 = int32(0)
						if v47 == v660 {
							v730 = int64(0)
						} else {
							v667 = v47 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v47) {
								v672 = v59
								v674 = v659
								v677 = v660
								for {
									v678 = int32(4)
									v679 = v672 + v678
									v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+3)))
									v681 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v680)+uint32(_c_F_gtsvector_penalty[0]))))
									v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+2)))
									v683 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v682)+uint32(_c_F_gtsvector_penalty[0]))))
									v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)))
									v685 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v684)+uint32(_c_F_gtsvector_penalty[0]))))
									v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
									v687 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v686)+uint32(_c_F_gtsvector_penalty[0]))))
									v691 = v681 + (v683 + (v685 + (v674 + v687)))
									v693 = v677 + v678
									if v693 != v47&int32(-4) {
										v672 = v679
										v674 = v691
										v677 = v693
										continue
									} else {
										break
									}
									break
								}
								if v667 == int32(0) {
									v720 = v691
								} else {
									v697 = v679
									v699 = v691
									v704 = v697
									v705 = int32(0)
									v706 = v699
									for {
										v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
										v711 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v710)+uint32(_c_F_gtsvector_penalty[0]))))
										v712 = v706 + v711
										v713 = int32(1)
										v716 = v705 + v713
										if v716 != v667 {
											v704 = v704 + v713
											v705 = v716
											v706 = v712
											continue
										} else {
											break
										}
										break
									}
									v720 = v712
								}
							} else {
								v697 = v59
								v699 = v659
								v704 = v697
								v705 = int32(0)
								v706 = v699
								for {
									v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
									v711 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v710)+uint32(_c_F_gtsvector_penalty[0]))))
									v712 = v706 + v711
									v713 = int32(1)
									v716 = v705 + v713
									if v716 != v667 {
										v704 = v704 + v713
										v705 = v716
										v706 = v712
										continue
									} else {
										break
									}
									break
								}
								v720 = v712
							}
							v730 = v720
						}
						v940 = v730
					} else {
						if v47 == int32(0) {
							v940 = int64(0)
						} else {
							v219 = v47 & int32(3)
							v220 = int64(0)
							if base.Ui32(int32(4)) <= base.Ui32(v47) {
								v227 = v59
								v229 = int32(0)
								v239 = v220
								for {
									v241 = int32(4)
									v242 = v227 + v241
									v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+3)))
									v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_gtsvector_penalty[0]))))
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+2)))
									v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_gtsvector_penalty[0]))))
									v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+1)))
									v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_gtsvector_penalty[0]))))
									v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
									v250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_gtsvector_penalty[0]))))
									v254 = v244 + (v246 + (v248 + (v239 + v250)))
									v256 = v229 + v241
									if v256 != v47&int32(-4) {
										v227 = v242
										v229 = v256
										v239 = v254
										continue
									} else {
										break
									}
									break
								}
								if v219 == int32(0) {
									v940 = v254
								} else {
									v261 = v242
									v273 = v254
									v277 = v261
									v280 = int32(0)
									v289 = v273
									for {
										v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
										v292 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+uint32(_c_F_gtsvector_penalty[0]))))
										v293 = v289 + v292
										v294 = int32(1)
										v297 = v280 + v294
										if v297 != v219 {
											v277 = v277 + v294
											v280 = v297
											v289 = v293
											continue
										} else {
											break
										}
										break
									}
									v940 = v293
								}
							} else {
								v261 = v59
								v273 = v220
								v277 = v261
								v280 = int32(0)
								v289 = v273
								for {
									v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
									v292 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+uint32(_c_F_gtsvector_penalty[0]))))
									v293 = v289 + v292
									v294 = int32(1)
									v297 = v280 + v294
									if v297 != v219 {
										v277 = v277 + v294
										v280 = v297
										v289 = v293
										continue
									} else {
										break
									}
									break
								}
								v940 = v293
							}
						}
					}
					v964 = base.F32_div(base.F32_convert_i32_s(v212-base.I32_wrap_i64(v940)), base.F32_convert_i32_s(v212|int32(1)))
				} else {
					if v47 <= int32(0) {
						v964 = float32(0)
					} else {
						v302 = int32(0)
						if v47 != int32(1) {
							v312 = v302
							v314 = v302
							v316 = int32(0)
							for {
								v327 = v312 | int32(1)
								v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v327))))
								v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v327))))
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329^v331)+uint32(_c_F_gtsvector_penalty[0]))))
								v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312+v55))))
								v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v312))))
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335^v337)+uint32(_c_F_gtsvector_penalty[0]))))
								v341 = v333 + (v314 + v339)
								v342 = int32(2)
								v343 = v312 + v342
								v345 = v316 + v342
								if v345 != v47&int32(2147483646) {
									v312 = v343
									v314 = v341
									v316 = v345
									continue
								} else {
									break
								}
								break
							}
							if v47&int32(1) == int32(0) {
								v374 = v341
							} else {
								v350 = v343
								v352 = v341
								v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v55))))
								v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v350))))
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365^v367)+uint32(_c_F_gtsvector_penalty[0]))))
								v374 = v352 + v369
							}
						} else {
							v350 = v302
							v352 = v302
							v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v55))))
							v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v350))))
							v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365^v367)+uint32(_c_F_gtsvector_penalty[0]))))
							v374 = v352 + v369
						}
						v964 = base.F32_convert_i32_s(v374)
					}
				}
				*(*float32)(unsafe.Add(mBase, uint32(v19))) = v964
				F_pfree(m, v59)
				mBase = m.M
				v967 = m.ExcPending
				if v967 != 0 {
					return int64(0)
				} else {
					return v49
				}
			}
		} else {
			v387 = int32(4)
			v388 = v56 & v387
			v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
			if v389&v387 != 0 {
				if v388 != 0 {
					v910 = int32(0)
				} else {
					v393 = int32(8)
					v394 = v50 + v393
					v395 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v397 = int32(base.Ui32(v395) >> (uint(int32(2)) % 32))
					v399 = v397 - v393
					if base.Ui32(int32(63)) < base.Ui32(v395) {
						v731 = int64(0)
						v732 = int32(0)
						if v399 == v732 {
							v802 = int64(0)
						} else {
							v739 = v399 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v399) {
								v744 = v394
								v746 = v731
								v749 = v732
								for {
									v750 = int32(4)
									v751 = v744 + v750
									v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v752)+uint32(_c_F_gtsvector_penalty[0]))))
									v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v755 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v754)+uint32(_c_F_gtsvector_penalty[0]))))
									v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v757 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v756)+uint32(_c_F_gtsvector_penalty[0]))))
									v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v759 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_gtsvector_penalty[0]))))
									v763 = v753 + (v755 + (v757 + (v746 + v759)))
									v765 = v749 + v750
									if v765 != v399&int32(-4) {
										v744 = v751
										v746 = v763
										v749 = v765
										continue
									} else {
										break
									}
									break
								}
								if v739 == int32(0) {
									v792 = v763
								} else {
									v769 = v751
									v771 = v763
									v776 = v769
									v777 = int32(0)
									v778 = v771
									for {
										v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
										v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
										v784 = v778 + v783
										v785 = int32(1)
										v788 = v777 + v785
										if v788 != v739 {
											v776 = v776 + v785
											v777 = v788
											v778 = v784
											continue
										} else {
											break
										}
										break
									}
									v792 = v784
								}
							} else {
								v769 = v394
								v771 = v731
								v776 = v769
								v777 = int32(0)
								v778 = v771
								for {
									v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
									v783 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_gtsvector_penalty[0]))))
									v784 = v778 + v783
									v785 = int32(1)
									v788 = v777 + v785
									if v788 != v739 {
										v776 = v776 + v785
										v777 = v788
										v778 = v784
										continue
									} else {
										break
									}
									break
								}
								v792 = v784
							}
							v802 = v792
						}
						v905 = v802
					} else {
						if v399 == int32(0) {
							v905 = int64(0)
						} else {
							v407 = int32(3)
							v408 = v397 & v407
							v409 = int64(0)
							if base.Ui32(v407) <= base.Ui32(v397-int32(9)) {
								v417 = v394
								v421 = int32(0)
								v430 = v409
								for {
									v432 = int32(4)
									v433 = v417 + v432
									v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+3)))
									v435 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_gtsvector_penalty[0]))))
									v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+2)))
									v437 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v436)+uint32(_c_F_gtsvector_penalty[0]))))
									v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)))
									v439 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v438)+uint32(_c_F_gtsvector_penalty[0]))))
									v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
									v441 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v440)+uint32(_c_F_gtsvector_penalty[0]))))
									v445 = v435 + (v437 + (v439 + (v430 + v441)))
									v447 = v421 + v432
									if v447 != v399&int32(-4) {
										v417 = v433
										v421 = v447
										v430 = v445
										continue
									} else {
										break
									}
									break
								}
								if v408 == int32(0) {
									v905 = v445
								} else {
									v451 = v433
									v464 = v445
									v467 = v451
									v468 = int32(0)
									v480 = v464
									for {
										v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
										v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_gtsvector_penalty[0]))))
										v484 = v480 + v483
										v485 = int32(1)
										v488 = v468 + v485
										if v488 != v408 {
											v467 = v467 + v485
											v468 = v488
											v480 = v484
											continue
										} else {
											break
										}
										break
									}
									v905 = v484
								}
							} else {
								v451 = v394
								v464 = v409
								v467 = v451
								v468 = int32(0)
								v480 = v464
								for {
									v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
									v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_gtsvector_penalty[0]))))
									v484 = v480 + v483
									v485 = int32(1)
									v488 = v468 + v485
									if v488 != v408 {
										v467 = v467 + v485
										v468 = v488
										v480 = v484
										continue
									} else {
										break
									}
									break
								}
								v905 = v484
							}
						}
					}
					v910 = v399<<(uint(int32(3))%32) - base.I32_wrap_i64(v905)
				}
			} else {
				v490 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
				v492 = int32(base.Ui32(v490) >> (uint(int32(2)) % 32))
				v494 = v492 - int32(8)
				if v388 != 0 {
					if base.Ui32(int32(63)) < base.Ui32(v490) {
						v803 = int64(0)
						v804 = int32(0)
						if v494 == v804 {
							v874 = int64(0)
						} else {
							v811 = v494 & int32(3)
							if base.Ui32(int32(4)) <= base.Ui32(v494) {
								v816 = v55
								v818 = v803
								v821 = v804
								for {
									v822 = int32(4)
									v823 = v816 + v822
									v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+3)))
									v825 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v824)+uint32(_c_F_gtsvector_penalty[0]))))
									v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+2)))
									v827 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v826)+uint32(_c_F_gtsvector_penalty[0]))))
									v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
									v829 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v828)+uint32(_c_F_gtsvector_penalty[0]))))
									v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
									v831 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtsvector_penalty[0]))))
									v835 = v825 + (v827 + (v829 + (v818 + v831)))
									v837 = v821 + v822
									if v837 != v494&int32(-4) {
										v816 = v823
										v818 = v835
										v821 = v837
										continue
									} else {
										break
									}
									break
								}
								if v811 == int32(0) {
									v864 = v835
								} else {
									v841 = v823
									v843 = v835
									v848 = v841
									v849 = int32(0)
									v850 = v843
									for {
										v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
										v855 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_gtsvector_penalty[0]))))
										v856 = v850 + v855
										v857 = int32(1)
										v860 = v849 + v857
										if v860 != v811 {
											v848 = v848 + v857
											v849 = v860
											v850 = v856
											continue
										} else {
											break
										}
										break
									}
									v864 = v856
								}
							} else {
								v841 = v55
								v843 = v803
								v848 = v841
								v849 = int32(0)
								v850 = v843
								for {
									v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848))))
									v855 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v854)+uint32(_c_F_gtsvector_penalty[0]))))
									v856 = v850 + v855
									v857 = int32(1)
									v860 = v849 + v857
									if v860 != v811 {
										v848 = v848 + v857
										v849 = v860
										v850 = v856
										continue
									} else {
										break
									}
									break
								}
								v864 = v856
							}
							v874 = v864
						}
						v888 = v874
					} else {
						if v494 == int32(0) {
							v888 = int64(0)
						} else {
							v502 = int32(3)
							v503 = v492 & v502
							v504 = int64(0)
							if base.Ui32(v502) <= base.Ui32(v492-int32(9)) {
								v513 = int32(0)
								v516 = v55
								v525 = v504
								for {
									v527 = int32(4)
									v528 = v516 + v527
									v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
									v530 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v529)+uint32(_c_F_gtsvector_penalty[0]))))
									v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
									v532 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v531)+uint32(_c_F_gtsvector_penalty[0]))))
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
									v534 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v533)+uint32(_c_F_gtsvector_penalty[0]))))
									v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
									v536 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v535)+uint32(_c_F_gtsvector_penalty[0]))))
									v540 = v530 + (v532 + (v534 + (v525 + v536)))
									v542 = v513 + v527
									if v542 != v494&int32(-4) {
										v513 = v542
										v516 = v528
										v525 = v540
										continue
									} else {
										break
									}
									break
								}
								if v503 == int32(0) {
									v888 = v540
								} else {
									v550 = v528
									v559 = v540
									v562 = int32(0)
									v566 = v550
									v575 = v559
									for {
										v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
										v578 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_c_F_gtsvector_penalty[0]))))
										v579 = v575 + v578
										v580 = int32(1)
										v583 = v562 + v580
										if v583 != v503 {
											v562 = v583
											v566 = v566 + v580
											v575 = v579
											continue
										} else {
											break
										}
										break
									}
									v888 = v579
								}
							} else {
								v550 = v55
								v559 = v504
								v562 = int32(0)
								v566 = v550
								v575 = v559
								for {
									v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
									v578 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v577)+uint32(_c_F_gtsvector_penalty[0]))))
									v579 = v575 + v578
									v580 = int32(1)
									v583 = v562 + v580
									if v583 != v503 {
										v562 = v583
										v566 = v566 + v580
										v575 = v579
										continue
									} else {
										break
									}
									break
								}
								v888 = v579
							}
						}
					}
					v910 = v494<<(uint(int32(3))%32) - base.I32_wrap_i64(v888)
				} else {
					if base.Ui32(v490) < base.Ui32(int32(36)) {
						v910 = int32(0)
					} else {
						v589 = v50 + int32(8)
						v590 = int32(0)
						if v492 != int32(9) {
							v599 = v590
							v600 = v590
							v604 = int32(0)
							for {
								v615 = v599 | int32(1)
								v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v615))))
								v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v615))))
								v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617^v619)+uint32(_c_F_gtsvector_penalty[0]))))
								v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v589))))
								v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v55))))
								v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623^v625)+uint32(_c_F_gtsvector_penalty[0]))))
								v629 = v621 + (v600 + v627)
								v630 = int32(2)
								v631 = v599 + v630
								v633 = v604 + v630
								if v633 != v494&int32(-2) {
									v599 = v631
									v600 = v629
									v604 = v633
									continue
								} else {
									break
								}
								break
							}
							if v492&int32(1) == int32(0) {
								v910 = v629
							} else {
								v637 = v631
								v638 = v629
								v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v589))))
								v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v55))))
								v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653^v655)+uint32(_c_F_gtsvector_penalty[0]))))
								v910 = v638 + v657
							}
						} else {
							v637 = v590
							v638 = v590
							v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v589))))
							v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v55))))
							v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653^v655)+uint32(_c_F_gtsvector_penalty[0]))))
							v910 = v638 + v657
						}
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(v19))) = base.F32_convert_i32_s(v910)
			return v49
		}
	}
}
