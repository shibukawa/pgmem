package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_unicode_normalize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v444 int32
	_ = v444
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
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
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
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
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
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
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int64
	_ = v820
	var v826 int64
	_ = v826
	var v830 int64
	_ = v830
	var v831 int64
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v859 int64
	_ = v859
	var v860 int64
	_ = v860
	var v862 int64
	_ = v862
	var v863 int64
	_ = v863
	var v865 int32
	_ = v865
	var v871 int64
	_ = v871
	var v872 int64
	_ = v872
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v965 int32
	_ = v965
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1043 int32
	_ = v1043
	var v1060 int32
	_ = v1060
	var v1068 int32
	_ = v1068
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v29 = l0 & int32(-2)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v30 != 0 {
		v32 = base.B2i32(v29 == int32(2))
		v35 = l1
		v38 = v30
		v43 = v3
		for {
			v62 = v38 - int32(_a_F_unicode_normalize_0)
			if base.Ui32(v62) <= base.Ui32(int32(_a_F_unicode_normalize_1)) {
				v70 = base.I32_rem_u_s(v62&int32(_a_F_unicode_normalize_2), int32(28))
				if v70 != 0 {
					v71 = int32(3)
				} else {
					v71 = int32(2)
				}
				v188 = v71
			} else {
				v72 = int32(1)
				v73 = int32(16711935)
				v75 = int32(8)
				v76 = base.I32_rotr(v38&v73, v75)
				v77 = int32(24)
				v78 = base.I32_rotr(v38, v77)
				v80 = int32(255)
				v81 = (v76 | v78) & v80
				v82 = int32(_a_F_unicode_normalize_3)
				v87 = int32(base.Ui32(v76)>>(uint(v75)%32)) & v80
				v94 = int32(base.Ui32(v78&v73) >> (uint(int32(16)) % 32))
				v99 = int32(base.Ui32(v76) >> (uint(v77) % 32))
				v103 = int32(_a_F_unicode_normalize_4)
				v104 = base.I32_rem_u_s(((v81*v82+v87)*v82+v94)*v82+v99+int32(402620417), v103)
				v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104<<(uint(v72)%32))+uint32(_c_F_unicode_normalize[0]))))
				v108 = int32(257)
				v118 = base.I32_rem_u_s(((v81*v108+v87)*v108+v94)*v108+v99, v103)
				v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118<<(uint(v72)%32))+uint32(_c_F_unicode_normalize[0]))))
				v122 = v107 + v121
				if base.Ui32(int32(_a_F_unicode_normalize_5)) < base.Ui32(v122) {
					v177 = v72
				} else {
					v126 = v122 << (uint(int32(3)) % 32)
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_unicode_normalize[1])))
					if v38 != v127 {
						v177 = v72
					} else {
						v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_unicode_normalize[2]))))
						v133 = v131 & int32(31)
						if v131&int32(32) != 0 {
							v139 = v32
						} else {
							v139 = int32(1)
						}
						if base.B2i32(v133 == int32(0))|base.B2i32(v139 == int32(0)) != 0 {
							v177 = v72
						} else {
							v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_unicode_normalize[3]))))
							if v131&int32(64) != 0 {
								v146 = int32(_a_F_unicode_normalize_6)
								*(*int32)(unsafe.Add(mBase, _c_F_unicode_normalize[4])) = v143
								v154 = int32(1)
								v155 = v146
							} else {
								v154 = v133
								v155 = v143<<(uint(int32(2))%32) + int32(_a_F_unicode_normalize_7)
							}
							v156 = int32(0)
							v158 = v156
							v161 = v156
							for {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v155+v158<<(uint(int32(2))%32))))
								v169 = F_get_decomposed_size(m, v168, v32)
								mBase = m.M
								v170 = v169 + v161
								v172 = v158 + int32(1)
								if v172 != v154 {
									v158 = v172
									v161 = v170
									continue
								} else {
									break
								}
								break
							}
							v177 = v170
						}
					}
				}
				v188 = v177
			}
			v189 = v188 + v43
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
			if v190 != 0 {
				v35 = v35 + int32(4)
				v38 = v190
				v43 = v189
				continue
			} else {
				break
			}
			break
		}
		v203 = v189
	} else {
		v203 = v3
	}
	v219 = v203<<(uint(int32(2))%32) + int32(4)
	v220 = F_palloc(m, v219)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v220
		if v220 == int32(0) {
			v1068 = v3
			m.G0 = v26 + int32(16)
			return v1068
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(0)
			v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v229 != 0 {
				v231 = base.B2i32(v29 == int32(2))
				v233 = l1
				v237 = v229
				for {
					v256 = v26 + int32(12)
					v258 = v26 + int32(8)
					v264 = v237 - int32(_a_F_unicode_normalize_0)
					if base.Ui32(v264) <= base.Ui32(int32(_a_F_unicode_normalize_1)) {
						v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
						v268 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v269 = int32(2)
						v272 = int32(_a_F_unicode_normalize_2)
						v273 = v264 & v272
						v274 = int32(588)
						v275 = base.I32_div_u_s(v273, v274)
						*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32)))) = v275 | int32(_a_F_unicode_normalize_8)
						v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v280 = int32(1)
						v281 = v279 + v280
						*(*int32)(unsafe.Add(mBase, uint32(v258))) = v281
						v291 = int32(28)
						v292 = base.I32_div_u_s((v264-v275*v274)&v272, v291)
						*(*int32)(unsafe.Add(mBase, uint32(v267+v281<<(uint(v269)%32)))) = v292 + int32(_a_F_unicode_normalize_9)
						v296 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v298 = v296 + v280
						*(*int32)(unsafe.Add(mBase, uint32(v258))) = v298
						v301 = base.I32_rem_u_s(v273, v291)
						if v301 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v267+v298<<(uint(int32(2))%32)))) = v301 + int32(_a_F_unicode_normalize_10)
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
						}
					} else {
						v310 = int32(16711935)
						v312 = int32(8)
						v313 = base.I32_rotr(v237&v310, v312)
						v314 = int32(24)
						v315 = base.I32_rotr(v237, v314)
						v317 = int32(255)
						v318 = (v313 | v315) & v317
						v319 = int32(_a_F_unicode_normalize_3)
						v324 = int32(base.Ui32(v313)>>(uint(v312)%32)) & v317
						v331 = int32(base.Ui32(v315&v310) >> (uint(int32(16)) % 32))
						v336 = int32(base.Ui32(v313) >> (uint(v314) % 32))
						v340 = int32(_a_F_unicode_normalize_4)
						v341 = base.I32_rem_u_s(((v318*v319+v324)*v319+v331)*v319+v336+int32(402620417), v340)
						v342 = int32(1)
						v344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341<<(uint(v342)%32))+uint32(_c_F_unicode_normalize[0]))))
						v345 = int32(257)
						v355 = base.I32_rem_u_s(((v318*v345+v324)*v345+v331)*v345+v336, v340)
						v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355<<(uint(v342)%32))+uint32(_c_F_unicode_normalize[0]))))
						v359 = v344 + v358
						if base.Ui32(int32(_a_F_unicode_normalize_5)) < base.Ui32(v359) {
							v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v237
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
						} else {
							v363 = v359 << (uint(int32(3)) % 32)
							v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_unicode_normalize[1])))
							if v237 != v364 {
								v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
								*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v237
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
								*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
							} else {
								v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_unicode_normalize[2]))))
								v370 = v368 & int32(31)
								if v370 == int32(0) {
									v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
									*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v237
									v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
									*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
								} else {
									if v231|base.B2i32(v368&int32(32) == int32(0)) != 0 {
										v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_unicode_normalize[3]))))
										if v368&int32(64) != 0 {
											v401 = int32(_a_F_unicode_normalize_6)
											*(*int32)(unsafe.Add(mBase, _c_F_unicode_normalize[4])) = v398
											v409 = int32(1)
											v410 = v401
										} else {
											v409 = v370
											v410 = v398<<(uint(int32(2))%32) + int32(_a_F_unicode_normalize_7)
										}
										v412 = int32(0)
										for {
											v423 = *(*int32)(unsafe.Add(mBase, uint32(v410+v412<<(uint(int32(2))%32))))
											F_decompose_code(m, v423, v231, v256, v258)
											mBase = m.M
											v426 = v412 + int32(1)
											if v426 != v409 {
												v412 = v426
												continue
											} else {
												break
											}
											break
										}
									} else {
										v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
										v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
										*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v237
										v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
										*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
									}
								}
							}
						}
					}
					v444 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
					if v444 != 0 {
						v233 = v233 + int32(4)
						v237 = v444
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v473 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v220+v203<<(uint(int32(2))%32)))) = v473
			if v203 == v473 {
				v1068 = v220
				m.G0 = v26 + int32(16)
				return v1068
			} else {
				if int32(2) <= v203 {
					v482 = int32(1)
					for {
						v505 = v220 + v482<<(uint(int32(2))%32)
						v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
						v507 = int32(0)
						v509 = v505 - int32(4)
						v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
						v511 = int32(16711935)
						v513 = int32(8)
						v514 = base.I32_rotr(v510&v511, v513)
						v515 = int32(24)
						v516 = base.I32_rotr(v510, v515)
						v518 = int32(255)
						v519 = (v514 | v516) & v518
						v520 = int32(_a_F_unicode_normalize_3)
						v525 = int32(base.Ui32(v514)>>(uint(v513)%32)) & v518
						v532 = int32(base.Ui32(v516&v511) >> (uint(int32(16)) % 32))
						v537 = int32(base.Ui32(v514) >> (uint(v515) % 32))
						v541 = int32(_a_F_unicode_normalize_4)
						v542 = base.I32_rem_u_s(((v519*v520+v525)*v520+v532)*v520+v537+int32(402620417), v541)
						v543 = int32(1)
						v545 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542<<(uint(v543)%32))+uint32(_c_F_unicode_normalize[0]))))
						v546 = int32(257)
						v556 = base.I32_rem_u_s(((v519*v546+v525)*v546+v532)*v546+v537, v541)
						v559 = int32(*(*int16)(unsafe.Add(mBase, uint32(v556<<(uint(v543)%32))+uint32(_c_F_unicode_normalize[0]))))
						v560 = v545 + v559
						if base.Ui32(int32(_a_F_unicode_normalize_5)) < base.Ui32(v560) {
							v571 = v507
						} else {
							v564 = v560 << (uint(int32(3)) % 32)
							v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+uint32(_c_F_unicode_normalize[1])))
							if v510 != v565 {
								v571 = v507
							} else {
								v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+uint32(_c_F_unicode_normalize[5]))))
								v571 = v569
							}
						}
						v572 = int32(16711935)
						v574 = int32(8)
						v575 = base.I32_rotr(v506&v572, v574)
						v576 = int32(24)
						v577 = base.I32_rotr(v506, v576)
						v579 = int32(255)
						v580 = (v575 | v577) & v579
						v581 = int32(_a_F_unicode_normalize_3)
						v586 = int32(base.Ui32(v575)>>(uint(v574)%32)) & v579
						v593 = int32(base.Ui32(v577&v572) >> (uint(int32(16)) % 32))
						v598 = int32(base.Ui32(v575) >> (uint(v576) % 32))
						v602 = int32(_a_F_unicode_normalize_4)
						v603 = base.I32_rem_u_s(((v580*v581+v586)*v581+v593)*v581+v598+int32(402620417), v602)
						v604 = int32(1)
						v606 = int32(*(*int16)(unsafe.Add(mBase, uint32(v603<<(uint(v604)%32))+uint32(_c_F_unicode_normalize[0]))))
						v607 = int32(257)
						v617 = base.I32_rem_u_s(((v580*v607+v586)*v607+v593)*v607+v598, v602)
						v620 = int32(*(*int16)(unsafe.Add(mBase, uint32(v617<<(uint(v604)%32))+uint32(_c_F_unicode_normalize[0]))))
						v621 = v606 + v620
						if base.Ui32(int32(_a_F_unicode_normalize_5)) < base.Ui32(v621) {
							v645 = v482
						} else {
							v627 = v621 << (uint(int32(3)) % 32)
							v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+uint32(_c_F_unicode_normalize[1])))
							if base.B2i32(v571 == int32(0))|base.B2i32(v506 != v628) != 0 {
								v645 = v482
							} else {
								v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+uint32(_c_F_unicode_normalize[5]))))
								if base.B2i32(v633 == int32(0))|base.B2i32(base.Ui32(v571) <= base.Ui32(v633)) != 0 {
									v645 = v482
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v509))) = v506
									*(*int32)(unsafe.Add(mBase, uint32(v505))) = v510
									if int32(1) < v482 {
										v644 = v482 - int32(2)
									} else {
										v644 = v482
									}
									v645 = v644
								}
							}
						}
						v648 = v645 + int32(1)
						if v648 < v203 {
							v482 = v648
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				if l0&int32(-3) != 0 {
					v1068 = v220
					m.G0 = v26 + int32(16)
					return v1068
				} else {
					v675 = F_palloc(m, v219)
					mBase = m.M
					v676 = m.ExcPending
					if v676 != 0 {
						return int32(0)
					} else {
						if v675 == int32(0) {
							v1043 = int32(0)
						} else {
							v680 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
							*(*int32)(unsafe.Add(mBase, uint32(v675))) = v680
							v682 = int32(1)
							if int32(2) <= v203 {
								v688 = v682
								v690 = v680
								v692 = int32(1)
								v697 = int32(0)
								v702 = int32(-1)
								for {
									v711 = int32(0)
									v715 = *(*int32)(unsafe.Add(mBase, uint32(v220+v692<<(uint(int32(2))%32))))
									v716 = int32(16711935)
									v718 = int32(8)
									v719 = base.I32_rotr(v715&v716, v718)
									v720 = int32(24)
									v721 = base.I32_rotr(v715, v720)
									v723 = int32(255)
									v724 = (v719 | v721) & v723
									v725 = int32(_a_F_unicode_normalize_3)
									v730 = int32(base.Ui32(v719)>>(uint(v718)%32)) & v723
									v737 = int32(base.Ui32(v721&v716) >> (uint(int32(16)) % 32))
									v742 = int32(base.Ui32(v719) >> (uint(v720) % 32))
									v746 = int32(_a_F_unicode_normalize_4)
									v747 = base.I32_rem_u_s(((v724*v725+v730)*v725+v737)*v725+v742+int32(402620417), v746)
									v748 = int32(1)
									v750 = int32(*(*int16)(unsafe.Add(mBase, uint32(v747<<(uint(v748)%32))+uint32(_c_F_unicode_normalize[0]))))
									v751 = int32(257)
									v761 = base.I32_rem_u_s(((v724*v751+v730)*v751+v737)*v751+v742, v746)
									v764 = int32(*(*int16)(unsafe.Add(mBase, uint32(v761<<(uint(v748)%32))+uint32(_c_F_unicode_normalize[0]))))
									v765 = v750 + v764
									if base.Ui32(int32(_a_F_unicode_normalize_5)) < base.Ui32(v765) {
										v775 = v711
									} else {
										v769 = v765 << (uint(int32(3)) % 32)
										v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+uint32(_c_F_unicode_normalize[1])))
										if v715 != v770 {
											v775 = v711
										} else {
											v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+uint32(_c_F_unicode_normalize[5]))))
											v775 = v774
										}
									}
									if v775 <= v702 {
										*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
										if v775 != 0 {
											v984 = v775
										} else {
											v984 = int32(-1)
										}
										if v775 != 0 {
											v985 = v697
										} else {
											v985 = v688
										}
										if v775 != 0 {
											v986 = v690
										} else {
											v986 = v715
										}
										v989 = v688 + int32(1)
										v990 = v986
										v993 = v985
										v996 = v984
									} else {
										if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v690-int32(_a_F_unicode_normalize_8)))|base.B2i32(base.Ui32(int32(20)) < base.Ui32(v715-int32(_a_F_unicode_normalize_9))) == int32(0) {
											v965 = (v715+v690*int32(21))*int32(28) - int32(_a_F_unicode_normalize_11)
											*(*int32)(unsafe.Add(mBase, uint32(v675+v697<<(uint(int32(2))%32)))) = v965
											v989 = v688
											v990 = v965
											v993 = v697
											v996 = v702
										} else {
											v800 = v690 - int32(_a_F_unicode_normalize_0)
											v806 = base.I32_rem_u_s(v800&int32(_a_F_unicode_normalize_2), int32(28))
											if base.B2i32(base.Ui32(int32(_a_F_unicode_normalize_1)) < base.Ui32(v800))|v806 != 0 {
												v814 = int32(24)
												v817 = int32(8)
												v820 = int64(16711680)
												v826 = int64(65280)
												v830 = base.I64_extend_i32_u(v690) << (uint(int64(32)) % 64)
												v831 = int64(56)
												v835 = base.I32_wrap_i64(base.I64_extend_i32_u(v690<<(uint(v814)%32)) | base.I64_extend_i32_u(v690<<(uint(v817)%32))&v820 | (base.I64_extend_i32_u(int32(base.Ui32(v690)>>(uint(v817)%32)))&v826 | int64(base.Ui64(v830)>>(uint(v831)%64))))
												v836 = int32(255)
												v837 = v835 & v836
												v838 = int32(17)
												v843 = int32(base.Ui32(v835)>>(uint(v817)%32)) & v836
												v850 = int32(base.Ui32(v835)>>(uint(int32(16))%32)) & v836
												v855 = int32(base.Ui32(v835) >> (uint(v814) % 32))
												v859 = base.I64_extend_i32_u(v715)
												v860 = v830 | v859
												v862 = v860 & int64(4278190080)
												v863 = int64(24)
												v865 = base.I32_wrap_i64(int64(base.Ui64(v862) >> (uint(v863) % 64)))
												v871 = int64(40)
												v872 = v859 & v826 << (uint(v871) % 64)
												v888 = base.I32_wrap_i64(int64(base.Ui64(v872|v859<<(uint(v831)%64)|(v860&v820<<(uint(v863)%64)|v862<<(uint(int64(8))%64)))>>(uint(v871)%64))) & v836
												v894 = base.I32_wrap_i64(int64(base.Ui64(v872) >> (uint(int64(48)) % 64)))
												v900 = base.I32_wrap_i64(v859 & int64(255))
												v902 = int32(1923)
												v903 = base.I32_rem_u_s(((((((v837*v838+v843)*v838+v850)*v838+v855)*v838+v865)*v838+v888)*v838+v894)*v838+v900, v902)
												v904 = int32(1)
												v906 = int32(*(*int16)(unsafe.Add(mBase, uint32(v903<<(uint(v904)%32))+uint32(_c_F_unicode_normalize[6]))))
												v907 = int32(257)
												v929 = base.I32_rem_u_s(((((((v837*v907+v843)*v907+v850)*v907+v855)*v907+v865)*v907+v888)*v907+v894)*v907+v900, v902)
												v932 = int32(*(*int16)(unsafe.Add(mBase, uint32(v929<<(uint(v904)%32))+uint32(_c_F_unicode_normalize[6]))))
												v933 = v906 + v932
												if base.Ui32(int32(960)) < base.Ui32(v933) {
													*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
													if v775 != 0 {
														v984 = v775
													} else {
														v984 = int32(-1)
													}
													if v775 != 0 {
														v985 = v697
													} else {
														v985 = v688
													}
													if v775 != 0 {
														v986 = v690
													} else {
														v986 = v715
													}
													v989 = v688 + int32(1)
													v990 = v986
													v993 = v985
													v996 = v984
												} else {
													v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v933<<(uint(int32(1))%32))+uint32(_c_F_unicode_normalize[7]))))
													v940 = v938 << (uint(int32(3)) % 32)
													v941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_unicode_normalize[3]))))
													v943 = v941 << (uint(int32(2)) % 32)
													v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_unicode_normalize[8])))
													if v690 != v944 {
														*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
														if v775 != 0 {
															v984 = v775
														} else {
															v984 = int32(-1)
														}
														if v775 != 0 {
															v985 = v697
														} else {
															v985 = v688
														}
														if v775 != 0 {
															v986 = v690
														} else {
															v986 = v715
														}
														v989 = v688 + int32(1)
														v990 = v986
														v993 = v985
														v996 = v984
													} else {
														v948 = *(*int32)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_unicode_normalize[9])))
														if v715 != v948 {
															*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
															if v775 != 0 {
																v984 = v775
															} else {
																v984 = int32(-1)
															}
															if v775 != 0 {
																v985 = v697
															} else {
																v985 = v688
															}
															if v775 != 0 {
																v986 = v690
															} else {
																v986 = v715
															}
															v989 = v688 + int32(1)
															v990 = v986
															v993 = v985
															v996 = v984
														} else {
															v952 = *(*int32)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_unicode_normalize[1])))
															v965 = v952
															*(*int32)(unsafe.Add(mBase, uint32(v675+v697<<(uint(int32(2))%32)))) = v965
															v989 = v688
															v990 = v965
															v993 = v697
															v996 = v702
														}
													}
												}
											} else {
												v809 = v715 - int32(_a_F_unicode_normalize_10)
												if base.Ui32(int32(27)) < base.Ui32(v809) {
													v814 = int32(24)
													v817 = int32(8)
													v820 = int64(16711680)
													v826 = int64(65280)
													v830 = base.I64_extend_i32_u(v690) << (uint(int64(32)) % 64)
													v831 = int64(56)
													v835 = base.I32_wrap_i64(base.I64_extend_i32_u(v690<<(uint(v814)%32)) | base.I64_extend_i32_u(v690<<(uint(v817)%32))&v820 | (base.I64_extend_i32_u(int32(base.Ui32(v690)>>(uint(v817)%32)))&v826 | int64(base.Ui64(v830)>>(uint(v831)%64))))
													v836 = int32(255)
													v837 = v835 & v836
													v838 = int32(17)
													v843 = int32(base.Ui32(v835)>>(uint(v817)%32)) & v836
													v850 = int32(base.Ui32(v835)>>(uint(int32(16))%32)) & v836
													v855 = int32(base.Ui32(v835) >> (uint(v814) % 32))
													v859 = base.I64_extend_i32_u(v715)
													v860 = v830 | v859
													v862 = v860 & int64(4278190080)
													v863 = int64(24)
													v865 = base.I32_wrap_i64(int64(base.Ui64(v862) >> (uint(v863) % 64)))
													v871 = int64(40)
													v872 = v859 & v826 << (uint(v871) % 64)
													v888 = base.I32_wrap_i64(int64(base.Ui64(v872|v859<<(uint(v831)%64)|(v860&v820<<(uint(v863)%64)|v862<<(uint(int64(8))%64)))>>(uint(v871)%64))) & v836
													v894 = base.I32_wrap_i64(int64(base.Ui64(v872) >> (uint(int64(48)) % 64)))
													v900 = base.I32_wrap_i64(v859 & int64(255))
													v902 = int32(1923)
													v903 = base.I32_rem_u_s(((((((v837*v838+v843)*v838+v850)*v838+v855)*v838+v865)*v838+v888)*v838+v894)*v838+v900, v902)
													v904 = int32(1)
													v906 = int32(*(*int16)(unsafe.Add(mBase, uint32(v903<<(uint(v904)%32))+uint32(_c_F_unicode_normalize[6]))))
													v907 = int32(257)
													v929 = base.I32_rem_u_s(((((((v837*v907+v843)*v907+v850)*v907+v855)*v907+v865)*v907+v888)*v907+v894)*v907+v900, v902)
													v932 = int32(*(*int16)(unsafe.Add(mBase, uint32(v929<<(uint(v904)%32))+uint32(_c_F_unicode_normalize[6]))))
													v933 = v906 + v932
													if base.Ui32(int32(960)) < base.Ui32(v933) {
														*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
														if v775 != 0 {
															v984 = v775
														} else {
															v984 = int32(-1)
														}
														if v775 != 0 {
															v985 = v697
														} else {
															v985 = v688
														}
														if v775 != 0 {
															v986 = v690
														} else {
															v986 = v715
														}
														v989 = v688 + int32(1)
														v990 = v986
														v993 = v985
														v996 = v984
													} else {
														v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v933<<(uint(int32(1))%32))+uint32(_c_F_unicode_normalize[7]))))
														v940 = v938 << (uint(int32(3)) % 32)
														v941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_unicode_normalize[3]))))
														v943 = v941 << (uint(int32(2)) % 32)
														v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_unicode_normalize[8])))
														if v690 != v944 {
															*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
															if v775 != 0 {
																v984 = v775
															} else {
																v984 = int32(-1)
															}
															if v775 != 0 {
																v985 = v697
															} else {
																v985 = v688
															}
															if v775 != 0 {
																v986 = v690
															} else {
																v986 = v715
															}
															v989 = v688 + int32(1)
															v990 = v986
															v993 = v985
															v996 = v984
														} else {
															v948 = *(*int32)(unsafe.Add(mBase, uint32(v943)+uint32(_c_F_unicode_normalize[9])))
															if v715 != v948 {
																*(*int32)(unsafe.Add(mBase, uint32(v675+v688<<(uint(int32(2))%32)))) = v715
																if v775 != 0 {
																	v984 = v775
																} else {
																	v984 = int32(-1)
																}
																if v775 != 0 {
																	v985 = v697
																} else {
																	v985 = v688
																}
																if v775 != 0 {
																	v986 = v690
																} else {
																	v986 = v715
																}
																v989 = v688 + int32(1)
																v990 = v986
																v993 = v985
																v996 = v984
															} else {
																v952 = *(*int32)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_unicode_normalize[1])))
																v965 = v952
																*(*int32)(unsafe.Add(mBase, uint32(v675+v697<<(uint(int32(2))%32)))) = v965
																v989 = v688
																v990 = v965
																v993 = v697
																v996 = v702
															}
														}
													}
												} else {
													v965 = v690 + v809
													*(*int32)(unsafe.Add(mBase, uint32(v675+v697<<(uint(int32(2))%32)))) = v965
													v989 = v688
													v990 = v965
													v993 = v697
													v996 = v702
												}
											}
										}
									}
									v1006 = v692 + int32(1)
									if v1006 != v203 {
										v688 = v989
										v690 = v990
										v692 = v1006
										v697 = v993
										v702 = v996
										continue
									} else {
										break
									}
									break
								}
								v1008 = v989
							} else {
								v1008 = v682
							}
							*(*int32)(unsafe.Add(mBase, uint32(v675+v1008<<(uint(int32(2))%32)))) = int32(0)
							v1043 = v675
						}
						F_pfree(m, v220)
						mBase = m.M
						v1060 = m.ExcPending
						if v1060 != 0 {
							return int32(0)
						} else {
							v1068 = v1043
							m.G0 = v26 + int32(16)
							return v1068
						}
					}
				}
			}
		}
	}
}
