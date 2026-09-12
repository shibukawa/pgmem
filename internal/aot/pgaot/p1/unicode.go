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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
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
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
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
	var v236 int32
	_ = v236
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
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
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
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
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int64
	_ = v824
	var v830 int64
	_ = v830
	var v834 int64
	_ = v834
	var v835 int64
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v863 int64
	_ = v863
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v867 int64
	_ = v867
	var v869 int32
	_ = v869
	var v875 int64
	_ = v875
	var v876 int64
	_ = v876
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v979 int32
	_ = v979
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1056 int32
	_ = v1056
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v29 = l0 & int32(-2)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v30 != 0 {
		v32 = base.B2i32(v29 == int32(2))
		v36 = l1
		v37 = v30
		v40 = v3
		for {
			v62 = v37 - int32(44032)
			if base.Ui32(v62) <= base.Ui32(int32(11171)) {
				v70 = base.I32_rem_u_s(v62&int32(65535), int32(28))
				if v70 != 0 {
					v71 = int32(3)
				} else {
					v71 = int32(2)
				}
				v188 = v71
			} else {
				v72 = int32(1)
				v74 = v37 & int32(255)
				v75 = int32(8)
				v80 = int32(base.Ui32(v37<<(uint(v75)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
				v86 = int32(base.Ui32(int32(base.Ui32(v37)>>(uint(v75)%32))&int32(65280)) >> (uint(v75) % 32))
				v88 = int32(base.Ui32(v37) >> (uint(int32(24)) % 32))
				v89 = int32(8191)
				v100 = int32(13687)
				v101 = base.I32_rem_u_s(v74+(v80+(v86+v88*v89)*v89)*v89+int32(402620417), v100)
				v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101<<(uint(v72)%32))+uint32(_consts[1263]))))
				v107 = int32(257)
				v117 = base.I32_rem_u_s(((v88*v107+v86)*v107+v80)*v107+v74, v100)
				v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117<<(uint(v72)%32))+uint32(_consts[1263]))))
				v123 = v106 + v122
				if base.Ui32(int32(6842)) < base.Ui32(v123) {
					v176 = v72
				} else {
					v127 = v123 << (uint(int32(3)) % 32)
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[1264])))
					if v37 != v130 {
						v176 = v72
					} else {
						v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[1265]))))
						v134 = v132 & int32(31)
						if v134 == int32(0) {
							v176 = v72
						} else {
							if v132&int32(32) != 0 {
								v140 = v32
							} else {
								v140 = int32(1)
							}
							if v140 == int32(0) {
								v176 = v72
							} else {
								v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[1266]))))
								if v132&int32(64) != 0 {
									v146 = int32(4645616)
									*(*int32)(unsafe.Add(mBase, _consts[1267])) = v143
									v154 = int32(1)
									v155 = v146
								} else {
									v154 = v134
									v155 = v143<<(uint(int32(2))%32) + int32(2098752)
								}
								v156 = int32(0)
								v158 = v156
								v160 = v156
								for {
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v155+v158<<(uint(int32(2))%32))))
									v169 = F_get_decomposed_size(m, v168, v32)
									mBase = m.M
									v170 = v169 + v160
									v172 = v158 + int32(1)
									if v172 != v154 {
										v158 = v172
										v160 = v170
										continue
									} else {
										break
									}
									break
								}
								v176 = v170
							}
						}
					}
				}
				v188 = v176
			}
			v189 = v188 + v40
			v191 = v36 + int32(4)
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
			if v192 != 0 {
				v36 = v191
				v37 = v192
				v40 = v189
				continue
			} else {
				break
			}
			break
		}
		v200 = v189
	} else {
		v200 = v3
	}
	v219 = v200<<(uint(int32(2))%32) + int32(4)
	v220 = F_palloc(m, v219)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v220
		if v220 == int32(0) {
			v1081 = v3
			m.G0 = v26 + int32(16)
			return v1081
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(0)
			v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v229 != 0 {
				v231 = base.B2i32(v29 == int32(2))
				v233 = l1
				v236 = v229
				for {
					v256 = v26 + int32(12)
					v258 = v26 + int32(8)
					v264 = v236 - int32(44032)
					if base.Ui32(v264) <= base.Ui32(int32(11171)) {
						v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
						v268 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v269 = int32(2)
						v272 = int32(65535)
						v273 = v264 & v272
						v274 = int32(588)
						v275 = base.I32_div_u_s(v273, v274)
						*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32)))) = v275 | int32(4352)
						v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v280 = int32(1)
						v281 = v279 + v280
						*(*int32)(unsafe.Add(mBase, uint32(v258))) = v281
						v291 = int32(28)
						v292 = base.I32_div_u_s((v264-v275*v274)&v272, v291)
						*(*int32)(unsafe.Add(mBase, uint32(v267+v281<<(uint(v269)%32)))) = v292 + int32(4449)
						v296 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
						v298 = v296 + v280
						*(*int32)(unsafe.Add(mBase, uint32(v258))) = v298
						v301 = base.I32_rem_u_s(v273, v291)
						if v301 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v267+v298<<(uint(int32(2))%32)))) = v301 + int32(4519)
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
						}
					} else {
						v311 = v236 & int32(255)
						v312 = int32(8)
						v317 = int32(base.Ui32(v236<<(uint(v312)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
						v323 = int32(base.Ui32(int32(base.Ui32(v236)>>(uint(v312)%32))&int32(65280)) >> (uint(v312) % 32))
						v325 = int32(base.Ui32(v236) >> (uint(int32(24)) % 32))
						v326 = int32(8191)
						v337 = int32(13687)
						v338 = base.I32_rem_u_s(v311+(v317+(v323+v325*v326)*v326)*v326+int32(402620417), v337)
						v339 = int32(1)
						v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338<<(uint(v339)%32))+uint32(_consts[1263]))))
						v344 = int32(257)
						v354 = base.I32_rem_u_s(((v325*v344+v323)*v344+v317)*v344+v311, v337)
						v359 = int32(*(*int16)(unsafe.Add(mBase, uint32(v354<<(uint(v339)%32))+uint32(_consts[1263]))))
						v360 = v343 + v359
						if base.Ui32(int32(6842)) < base.Ui32(v360) {
							v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v236
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
							*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
						} else {
							v364 = v360 << (uint(int32(3)) % 32)
							v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1264])))
							if v236 != v367 {
								v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
								*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v236
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
								*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
							} else {
								v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1265]))))
								v371 = v369 & int32(31)
								if v371 == int32(0) {
									v381 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
									*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v236
									v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
									*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
								} else {
									if v29 == int32(2) {
										v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1266]))))
										if v369&int32(64) != 0 {
											v401 = int32(4645616)
											*(*int32)(unsafe.Add(mBase, _consts[1267])) = v398
											v409 = int32(1)
											v410 = v401
										} else {
											v409 = v371
											v410 = v398<<(uint(int32(2))%32) + int32(2098752)
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
										if v369&int32(32) == int32(0) {
											v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364)+uint32(_consts[1266]))))
											if v369&int32(64) != 0 {
												v401 = int32(4645616)
												*(*int32)(unsafe.Add(mBase, _consts[1267])) = v398
												v409 = int32(1)
												v410 = v401
											} else {
												v409 = v371
												v410 = v398<<(uint(int32(2))%32) + int32(2098752)
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
											*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(int32(2))%32)))) = v236
											v394 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
											*(*int32)(unsafe.Add(mBase, uint32(v258))) = v394 + int32(1)
										}
									}
								}
							}
						}
					}
					v445 = v233 + int32(4)
					v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
					if v446 != 0 {
						v233 = v445
						v236 = v446
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v473 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v220+v200<<(uint(int32(2))%32)))) = v473
			if v200 == v473 {
				v1081 = v220
				m.G0 = v26 + int32(16)
				return v1081
			} else {
				if int32(2) <= v200 {
					v485 = int32(1)
					for {
						v508 = v485 << (uint(int32(2)) % 32)
						v509 = v220 + v508
						v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
						v511 = int32(0)
						v512 = v508 + (v220 - int32(4))
						v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
						v515 = int32(base.Ui32(v513) >> (uint(int32(24)) % 32))
						v516 = int32(8191)
						v518 = int32(8)
						v520 = int32(65280)
						v523 = int32(base.Ui32(int32(base.Ui32(v513)>>(uint(v518)%32))&v520) >> (uint(v518) % 32))
						v532 = int32(base.Ui32(v513&v520<<(uint(v518)%32)) >> (uint(int32(16)) % 32))
						v537 = v513 & int32(255)
						v541 = int32(13687)
						v542 = base.I32_rem_u_s(((v515*v516+v523)*v516+v532)*v516+v537+int32(402620417), v541)
						v543 = int32(1)
						v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v542<<(uint(v543)%32))+uint32(_consts[1263]))))
						v548 = int32(257)
						v558 = base.I32_rem_u_s(((v515*v548+v523)*v548+v532)*v548+v537, v541)
						v563 = int32(*(*int16)(unsafe.Add(mBase, uint32(v558<<(uint(v543)%32))+uint32(_consts[1263]))))
						v564 = v547 + v563
						if base.Ui32(int32(6842)) < base.Ui32(v564) {
							v575 = v511
						} else {
							v568 = v564 << (uint(int32(3)) % 32)
							v571 = *(*int32)(unsafe.Add(mBase, uint32(v568)+uint32(_consts[1264])))
							if v513 != v571 {
								v575 = v511
							} else {
								v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+uint32(_consts[1268]))))
								v575 = v573
							}
						}
						v577 = v510 & int32(255)
						v578 = int32(8)
						v583 = int32(base.Ui32(v510<<(uint(v578)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
						v589 = int32(base.Ui32(int32(base.Ui32(v510)>>(uint(v578)%32))&int32(65280)) >> (uint(v578) % 32))
						v591 = int32(base.Ui32(v510) >> (uint(int32(24)) % 32))
						v592 = int32(8191)
						v603 = int32(13687)
						v604 = base.I32_rem_u_s(v577+(v583+(v589+v591*v592)*v592)*v592+int32(402620417), v603)
						v605 = int32(1)
						v609 = int32(*(*int16)(unsafe.Add(mBase, uint32(v604<<(uint(v605)%32))+uint32(_consts[1263]))))
						v610 = int32(257)
						v620 = base.I32_rem_u_s(((v591*v610+v589)*v610+v583)*v610+v577, v603)
						v625 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620<<(uint(v605)%32))+uint32(_consts[1263]))))
						v626 = v609 + v625
						if base.Ui32(int32(6842)) < base.Ui32(v626) {
							v650 = v485
						} else {
							v630 = v626 << (uint(int32(3)) % 32)
							v633 = *(*int32)(unsafe.Add(mBase, uint32(v630)+uint32(_consts[1264])))
							if v510 != v633 {
								v650 = v485
							} else {
								v636 = v575 & int32(255)
								if v636 == int32(0) {
									v650 = v485
								} else {
									v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+uint32(_consts[1268]))))
									if v639 == int32(0) {
										v650 = v485
									} else {
										if base.Ui32(v636) <= base.Ui32(v639) {
											v650 = v485
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v512))) = v510
											*(*int32)(unsafe.Add(mBase, uint32(v509))) = v513
											if int32(1) < v485 {
												v649 = v485 - int32(2)
											} else {
												v649 = v485
											}
											v650 = v649
										}
									}
								}
							}
						}
						v654 = v650 + int32(1)
						if v654 < v200 {
							v485 = v654
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				if l0&int32(-3) != 0 {
					v1081 = v220
					m.G0 = v26 + int32(16)
					return v1081
				} else {
					v679 = F_palloc(m, v219)
					mBase = m.M
					v680 = m.ExcPending
					if v680 != 0 {
						return int32(0)
					} else {
						if v679 == int32(0) {
							v1056 = int32(0)
						} else {
							v684 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
							*(*int32)(unsafe.Add(mBase, uint32(v679))) = v684
							if v200 < int32(2) {
								v1022 = int32(1)
							} else {
								v693 = int32(1)
								v696 = v684
								v698 = int32(1)
								v704 = int32(0)
								v707 = int32(-1)
								for {
									v716 = int32(0)
									v720 = *(*int32)(unsafe.Add(mBase, uint32(v220+v698<<(uint(int32(2))%32))))
									v722 = int32(base.Ui32(v720) >> (uint(int32(24)) % 32))
									v723 = int32(8191)
									v725 = int32(8)
									v727 = int32(65280)
									v730 = int32(base.Ui32(int32(base.Ui32(v720)>>(uint(v725)%32))&v727) >> (uint(v725) % 32))
									v739 = int32(base.Ui32(v720&v727<<(uint(v725)%32)) >> (uint(int32(16)) % 32))
									v744 = v720 & int32(255)
									v748 = int32(13687)
									v749 = base.I32_rem_u_s(((v722*v723+v730)*v723+v739)*v723+v744+int32(402620417), v748)
									v750 = int32(1)
									v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v749<<(uint(v750)%32))+uint32(_consts[1263]))))
									v755 = int32(257)
									v765 = base.I32_rem_u_s(((v722*v755+v730)*v755+v739)*v755+v744, v748)
									v770 = int32(*(*int16)(unsafe.Add(mBase, uint32(v765<<(uint(v750)%32))+uint32(_consts[1263]))))
									v771 = v754 + v770
									if base.Ui32(int32(6842)) < base.Ui32(v771) {
										v781 = v716
									} else {
										v775 = v771 << (uint(int32(3)) % 32)
										v778 = *(*int32)(unsafe.Add(mBase, uint32(v775)+uint32(_consts[1264])))
										if v720 != v778 {
											v781 = v716
										} else {
											v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+uint32(_consts[1268]))))
											v781 = v780
										}
									}
									v784 = v781 & int32(255)
									if v784 <= v707 {
										*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
										if v784 != 0 {
											v998 = v784
										} else {
											v998 = int32(-1)
										}
										if v784 != 0 {
											v999 = v704
										} else {
											v999 = v693
										}
										if v784 != 0 {
											v1000 = v696
										} else {
											v1000 = v720
										}
										v1003 = v693 + int32(1)
										v1005 = v1000
										v1007 = v999
										v1010 = v998
									} else {
										if base.Ui32(int32(18)) < base.Ui32(v696-int32(4352)) {
											v805 = v696 - int32(44032)
											if base.Ui32(int32(11171)) < base.Ui32(v805) {
												v818 = int32(24)
												v821 = int32(8)
												v824 = int64(16711680)
												v830 = int64(65280)
												v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
												v835 = int64(56)
												v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
												v840 = int32(255)
												v841 = v839 & v840
												v842 = int32(17)
												v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
												v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
												v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
												v863 = base.I64_extend_i32_u(v720)
												v864 = v834 | v863
												v866 = v864 & int64(4278190080)
												v867 = int64(24)
												v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
												v875 = int64(40)
												v876 = v863 & v830 << (uint(v875) % 64)
												v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
												v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
												v904 = base.I32_wrap_i64(v863 & int64(255))
												v906 = int32(1923)
												v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
												v908 = int32(1)
												v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
												v913 = int32(257)
												v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
												v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
												v941 = v912 + v940
												if base.Ui32(int32(960)) < base.Ui32(v941) {
													*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
													if v784 != 0 {
														v998 = v784
													} else {
														v998 = int32(-1)
													}
													if v784 != 0 {
														v999 = v704
													} else {
														v999 = v693
													}
													if v784 != 0 {
														v1000 = v696
													} else {
														v1000 = v720
													}
													v1003 = v693 + int32(1)
													v1005 = v1000
													v1007 = v999
													v1010 = v998
												} else {
													v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
													v950 = v948 << (uint(int32(3)) % 32)
													v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
													v955 = v953 << (uint(int32(2)) % 32)
													v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
													if v696 != v958 {
														*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
														if v784 != 0 {
															v998 = v784
														} else {
															v998 = int32(-1)
														}
														if v784 != 0 {
															v999 = v704
														} else {
															v999 = v693
														}
														if v784 != 0 {
															v1000 = v696
														} else {
															v1000 = v720
														}
														v1003 = v693 + int32(1)
														v1005 = v1000
														v1007 = v999
														v1010 = v998
													} else {
														v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
														if v720 != v962 {
															*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
															if v784 != 0 {
																v998 = v784
															} else {
																v998 = int32(-1)
															}
															if v784 != 0 {
																v999 = v704
															} else {
																v999 = v693
															}
															if v784 != 0 {
																v1000 = v696
															} else {
																v1000 = v720
															}
															v1003 = v693 + int32(1)
															v1005 = v1000
															v1007 = v999
															v1010 = v998
														} else {
															v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
															v979 = v966
															*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
															v1003 = v693
															v1005 = v979
															v1007 = v704
															v1010 = v707
														}
													}
												}
											} else {
												v811 = base.I32_rem_u_s(v805&int32(65535), int32(28))
												if v811 != 0 {
													v818 = int32(24)
													v821 = int32(8)
													v824 = int64(16711680)
													v830 = int64(65280)
													v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
													v835 = int64(56)
													v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
													v840 = int32(255)
													v841 = v839 & v840
													v842 = int32(17)
													v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
													v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
													v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
													v863 = base.I64_extend_i32_u(v720)
													v864 = v834 | v863
													v866 = v864 & int64(4278190080)
													v867 = int64(24)
													v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
													v875 = int64(40)
													v876 = v863 & v830 << (uint(v875) % 64)
													v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
													v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
													v904 = base.I32_wrap_i64(v863 & int64(255))
													v906 = int32(1923)
													v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
													v908 = int32(1)
													v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
													v913 = int32(257)
													v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
													v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
													v941 = v912 + v940
													if base.Ui32(int32(960)) < base.Ui32(v941) {
														*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
														if v784 != 0 {
															v998 = v784
														} else {
															v998 = int32(-1)
														}
														if v784 != 0 {
															v999 = v704
														} else {
															v999 = v693
														}
														if v784 != 0 {
															v1000 = v696
														} else {
															v1000 = v720
														}
														v1003 = v693 + int32(1)
														v1005 = v1000
														v1007 = v999
														v1010 = v998
													} else {
														v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
														v950 = v948 << (uint(int32(3)) % 32)
														v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
														v955 = v953 << (uint(int32(2)) % 32)
														v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
														if v696 != v958 {
															*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
															if v784 != 0 {
																v998 = v784
															} else {
																v998 = int32(-1)
															}
															if v784 != 0 {
																v999 = v704
															} else {
																v999 = v693
															}
															if v784 != 0 {
																v1000 = v696
															} else {
																v1000 = v720
															}
															v1003 = v693 + int32(1)
															v1005 = v1000
															v1007 = v999
															v1010 = v998
														} else {
															v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
															if v720 != v962 {
																*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																if v784 != 0 {
																	v998 = v784
																} else {
																	v998 = int32(-1)
																}
																if v784 != 0 {
																	v999 = v704
																} else {
																	v999 = v693
																}
																if v784 != 0 {
																	v1000 = v696
																} else {
																	v1000 = v720
																}
																v1003 = v693 + int32(1)
																v1005 = v1000
																v1007 = v999
																v1010 = v998
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
																v979 = v966
																*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
																v1003 = v693
																v1005 = v979
																v1007 = v704
																v1010 = v707
															}
														}
													}
												} else {
													v813 = v720 - int32(4519)
													if base.Ui32(int32(27)) < base.Ui32(v813) {
														v818 = int32(24)
														v821 = int32(8)
														v824 = int64(16711680)
														v830 = int64(65280)
														v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
														v835 = int64(56)
														v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
														v840 = int32(255)
														v841 = v839 & v840
														v842 = int32(17)
														v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
														v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
														v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
														v863 = base.I64_extend_i32_u(v720)
														v864 = v834 | v863
														v866 = v864 & int64(4278190080)
														v867 = int64(24)
														v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
														v875 = int64(40)
														v876 = v863 & v830 << (uint(v875) % 64)
														v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
														v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
														v904 = base.I32_wrap_i64(v863 & int64(255))
														v906 = int32(1923)
														v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
														v908 = int32(1)
														v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
														v913 = int32(257)
														v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
														v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
														v941 = v912 + v940
														if base.Ui32(int32(960)) < base.Ui32(v941) {
															*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
															if v784 != 0 {
																v998 = v784
															} else {
																v998 = int32(-1)
															}
															if v784 != 0 {
																v999 = v704
															} else {
																v999 = v693
															}
															if v784 != 0 {
																v1000 = v696
															} else {
																v1000 = v720
															}
															v1003 = v693 + int32(1)
															v1005 = v1000
															v1007 = v999
															v1010 = v998
														} else {
															v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
															v950 = v948 << (uint(int32(3)) % 32)
															v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
															v955 = v953 << (uint(int32(2)) % 32)
															v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
															if v696 != v958 {
																*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																if v784 != 0 {
																	v998 = v784
																} else {
																	v998 = int32(-1)
																}
																if v784 != 0 {
																	v999 = v704
																} else {
																	v999 = v693
																}
																if v784 != 0 {
																	v1000 = v696
																} else {
																	v1000 = v720
																}
																v1003 = v693 + int32(1)
																v1005 = v1000
																v1007 = v999
																v1010 = v998
															} else {
																v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
																if v720 != v962 {
																	*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																	if v784 != 0 {
																		v998 = v784
																	} else {
																		v998 = int32(-1)
																	}
																	if v784 != 0 {
																		v999 = v704
																	} else {
																		v999 = v693
																	}
																	if v784 != 0 {
																		v1000 = v696
																	} else {
																		v1000 = v720
																	}
																	v1003 = v693 + int32(1)
																	v1005 = v1000
																	v1007 = v999
																	v1010 = v998
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
																	v979 = v966
																	*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
																	v1003 = v693
																	v1005 = v979
																	v1007 = v704
																	v1010 = v707
																}
															}
														}
													} else {
														v979 = v813 + v696
														*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
														v1003 = v693
														v1005 = v979
														v1007 = v704
														v1010 = v707
													}
												}
											}
										} else {
											if base.Ui32(int32(20)) < base.Ui32(v720-int32(4449)) {
												v805 = v696 - int32(44032)
												if base.Ui32(int32(11171)) < base.Ui32(v805) {
													v818 = int32(24)
													v821 = int32(8)
													v824 = int64(16711680)
													v830 = int64(65280)
													v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
													v835 = int64(56)
													v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
													v840 = int32(255)
													v841 = v839 & v840
													v842 = int32(17)
													v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
													v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
													v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
													v863 = base.I64_extend_i32_u(v720)
													v864 = v834 | v863
													v866 = v864 & int64(4278190080)
													v867 = int64(24)
													v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
													v875 = int64(40)
													v876 = v863 & v830 << (uint(v875) % 64)
													v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
													v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
													v904 = base.I32_wrap_i64(v863 & int64(255))
													v906 = int32(1923)
													v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
													v908 = int32(1)
													v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
													v913 = int32(257)
													v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
													v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
													v941 = v912 + v940
													if base.Ui32(int32(960)) < base.Ui32(v941) {
														*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
														if v784 != 0 {
															v998 = v784
														} else {
															v998 = int32(-1)
														}
														if v784 != 0 {
															v999 = v704
														} else {
															v999 = v693
														}
														if v784 != 0 {
															v1000 = v696
														} else {
															v1000 = v720
														}
														v1003 = v693 + int32(1)
														v1005 = v1000
														v1007 = v999
														v1010 = v998
													} else {
														v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
														v950 = v948 << (uint(int32(3)) % 32)
														v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
														v955 = v953 << (uint(int32(2)) % 32)
														v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
														if v696 != v958 {
															*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
															if v784 != 0 {
																v998 = v784
															} else {
																v998 = int32(-1)
															}
															if v784 != 0 {
																v999 = v704
															} else {
																v999 = v693
															}
															if v784 != 0 {
																v1000 = v696
															} else {
																v1000 = v720
															}
															v1003 = v693 + int32(1)
															v1005 = v1000
															v1007 = v999
															v1010 = v998
														} else {
															v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
															if v720 != v962 {
																*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																if v784 != 0 {
																	v998 = v784
																} else {
																	v998 = int32(-1)
																}
																if v784 != 0 {
																	v999 = v704
																} else {
																	v999 = v693
																}
																if v784 != 0 {
																	v1000 = v696
																} else {
																	v1000 = v720
																}
																v1003 = v693 + int32(1)
																v1005 = v1000
																v1007 = v999
																v1010 = v998
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
																v979 = v966
																*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
																v1003 = v693
																v1005 = v979
																v1007 = v704
																v1010 = v707
															}
														}
													}
												} else {
													v811 = base.I32_rem_u_s(v805&int32(65535), int32(28))
													if v811 != 0 {
														v818 = int32(24)
														v821 = int32(8)
														v824 = int64(16711680)
														v830 = int64(65280)
														v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
														v835 = int64(56)
														v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
														v840 = int32(255)
														v841 = v839 & v840
														v842 = int32(17)
														v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
														v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
														v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
														v863 = base.I64_extend_i32_u(v720)
														v864 = v834 | v863
														v866 = v864 & int64(4278190080)
														v867 = int64(24)
														v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
														v875 = int64(40)
														v876 = v863 & v830 << (uint(v875) % 64)
														v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
														v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
														v904 = base.I32_wrap_i64(v863 & int64(255))
														v906 = int32(1923)
														v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
														v908 = int32(1)
														v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
														v913 = int32(257)
														v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
														v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
														v941 = v912 + v940
														if base.Ui32(int32(960)) < base.Ui32(v941) {
															*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
															if v784 != 0 {
																v998 = v784
															} else {
																v998 = int32(-1)
															}
															if v784 != 0 {
																v999 = v704
															} else {
																v999 = v693
															}
															if v784 != 0 {
																v1000 = v696
															} else {
																v1000 = v720
															}
															v1003 = v693 + int32(1)
															v1005 = v1000
															v1007 = v999
															v1010 = v998
														} else {
															v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
															v950 = v948 << (uint(int32(3)) % 32)
															v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
															v955 = v953 << (uint(int32(2)) % 32)
															v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
															if v696 != v958 {
																*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																if v784 != 0 {
																	v998 = v784
																} else {
																	v998 = int32(-1)
																}
																if v784 != 0 {
																	v999 = v704
																} else {
																	v999 = v693
																}
																if v784 != 0 {
																	v1000 = v696
																} else {
																	v1000 = v720
																}
																v1003 = v693 + int32(1)
																v1005 = v1000
																v1007 = v999
																v1010 = v998
															} else {
																v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
																if v720 != v962 {
																	*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																	if v784 != 0 {
																		v998 = v784
																	} else {
																		v998 = int32(-1)
																	}
																	if v784 != 0 {
																		v999 = v704
																	} else {
																		v999 = v693
																	}
																	if v784 != 0 {
																		v1000 = v696
																	} else {
																		v1000 = v720
																	}
																	v1003 = v693 + int32(1)
																	v1005 = v1000
																	v1007 = v999
																	v1010 = v998
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
																	v979 = v966
																	*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
																	v1003 = v693
																	v1005 = v979
																	v1007 = v704
																	v1010 = v707
																}
															}
														}
													} else {
														v813 = v720 - int32(4519)
														if base.Ui32(int32(27)) < base.Ui32(v813) {
															v818 = int32(24)
															v821 = int32(8)
															v824 = int64(16711680)
															v830 = int64(65280)
															v834 = base.I64_extend_i32_u(v696) << (uint(int64(32)) % 64)
															v835 = int64(56)
															v839 = base.I32_wrap_i64(base.I64_extend_i32_u(v696<<(uint(v818)%32)) | base.I64_extend_i32_u(v696<<(uint(v821)%32))&v824 | (base.I64_extend_i32_u(int32(base.Ui32(v696)>>(uint(v821)%32)))&v830 | int64(base.Ui64(v834)>>(uint(v835)%64))))
															v840 = int32(255)
															v841 = v839 & v840
															v842 = int32(17)
															v847 = int32(base.Ui32(v839)>>(uint(v821)%32)) & v840
															v854 = int32(base.Ui32(v839)>>(uint(int32(16))%32)) & v840
															v859 = int32(base.Ui32(v839) >> (uint(v818) % 32))
															v863 = base.I64_extend_i32_u(v720)
															v864 = v834 | v863
															v866 = v864 & int64(4278190080)
															v867 = int64(24)
															v869 = base.I32_wrap_i64(int64(base.Ui64(v866) >> (uint(v867) % 64)))
															v875 = int64(40)
															v876 = v863 & v830 << (uint(v875) % 64)
															v892 = base.I32_wrap_i64(int64(base.Ui64(v876|v863<<(uint(v835)%64)|(v864&v824<<(uint(v867)%64)|v866<<(uint(int64(8))%64)))>>(uint(v875)%64))) & v840
															v898 = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(48)) % 64)))
															v904 = base.I32_wrap_i64(v863 & int64(255))
															v906 = int32(1923)
															v907 = base.I32_rem_u_s(((((((v841*v842+v847)*v842+v854)*v842+v859)*v842+v869)*v842+v892)*v842+v898)*v842+v904, v906)
															v908 = int32(1)
															v912 = int32(*(*int16)(unsafe.Add(mBase, uint32(v907<<(uint(v908)%32))+uint32(_consts[1269]))))
															v913 = int32(257)
															v935 = base.I32_rem_u_s(((((((v841*v913+v847)*v913+v854)*v913+v859)*v913+v869)*v913+v892)*v913+v898)*v913+v904, v906)
															v940 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v908)%32))+uint32(_consts[1269]))))
															v941 = v912 + v940
															if base.Ui32(int32(960)) < base.Ui32(v941) {
																*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																if v784 != 0 {
																	v998 = v784
																} else {
																	v998 = int32(-1)
																}
																if v784 != 0 {
																	v999 = v704
																} else {
																	v999 = v693
																}
																if v784 != 0 {
																	v1000 = v696
																} else {
																	v1000 = v720
																}
																v1003 = v693 + int32(1)
																v1005 = v1000
																v1007 = v999
																v1010 = v998
															} else {
																v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v941<<(uint(int32(1))%32))+uint32(_consts[1270]))))
																v950 = v948 << (uint(int32(3)) % 32)
																v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1266]))))
																v955 = v953 << (uint(int32(2)) % 32)
																v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1271])))
																if v696 != v958 {
																	*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																	if v784 != 0 {
																		v998 = v784
																	} else {
																		v998 = int32(-1)
																	}
																	if v784 != 0 {
																		v999 = v704
																	} else {
																		v999 = v693
																	}
																	if v784 != 0 {
																		v1000 = v696
																	} else {
																		v1000 = v720
																	}
																	v1003 = v693 + int32(1)
																	v1005 = v1000
																	v1007 = v999
																	v1010 = v998
																} else {
																	v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+uint32(_consts[1272])))
																	if v720 != v962 {
																		*(*int32)(unsafe.Add(mBase, uint32(v679+v693<<(uint(int32(2))%32)))) = v720
																		if v784 != 0 {
																			v998 = v784
																		} else {
																			v998 = int32(-1)
																		}
																		if v784 != 0 {
																			v999 = v704
																		} else {
																			v999 = v693
																		}
																		if v784 != 0 {
																			v1000 = v696
																		} else {
																			v1000 = v720
																		}
																		v1003 = v693 + int32(1)
																		v1005 = v1000
																		v1007 = v999
																		v1010 = v998
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)+uint32(_consts[1264])))
																		v979 = v966
																		*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
																		v1003 = v693
																		v1005 = v979
																		v1007 = v704
																		v1010 = v707
																	}
																}
															}
														} else {
															v979 = v813 + v696
															*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
															v1003 = v693
															v1005 = v979
															v1007 = v704
															v1010 = v707
														}
													}
												}
											} else {
												v979 = (v720+v696*int32(21))*int32(28) - int32(2639516)
												*(*int32)(unsafe.Add(mBase, uint32(v679+v704<<(uint(int32(2))%32)))) = v979
												v1003 = v693
												v1005 = v979
												v1007 = v704
												v1010 = v707
											}
										}
									}
									v1020 = v698 + int32(1)
									if v1020 != v200 {
										v693 = v1003
										v696 = v1005
										v698 = v1020
										v704 = v1007
										v707 = v1010
										continue
									} else {
										break
									}
									break
								}
								v1022 = v1003
							}
							*(*int32)(unsafe.Add(mBase, uint32(v679+v1022<<(uint(int32(2))%32)))) = int32(0)
							v1056 = v679
						}
						F_pfree(m, v220)
						mBase = m.M
						v1074 = m.ExcPending
						if v1074 != 0 {
							return int32(0)
						} else {
							v1081 = v1056
							m.G0 = v26 + int32(16)
							return v1081
						}
					}
				}
			}
		}
	}
}
