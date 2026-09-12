package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExplainSerializeDestReceiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(208))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(549)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(550)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(551)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(552)
		return v4
	}
}
func F_ExplainIndentText(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v5-int32(1)))))
		if v10 != int32(10) {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_appendStringInfoSpaces(m, v4, v13<<(uint(int32(1))%32))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_appendStringInfoSpaces(m, v4, v13<<(uint(int32(1))%32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExplainIndexScanDetails(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	if v11 != 0 {
		v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 != 0 {
				v19 = v12
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v20 == int32(0) {
					if l1 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						F_appendStringInfoString(m, v25, int32(419374))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v30 = F_quote_identifier(m, v19)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
								F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v30 = F_quote_identifier(m, v19)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
							F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					if l1 == int32(1) {
						v44 = int32(419283)
					} else {
						v44 = int32(542438)
					}
					if l1 == int32(-1) {
						v47 = int32(419375)
					} else {
						v47 = v44
					}
					F_ExplainPropertyText(m, int32(253537), v47, l2)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_ExplainPropertyText(m, int32(380012), v19, l2)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				v15 = F_get_rel_name(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if v15 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg_internal(m, int32(39997), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(493477), int32(4035), int32(376071))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v19 = v15
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						if v20 == int32(0) {
							if l1 == int32(-1) {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								F_appendStringInfoString(m, v25, int32(419374))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v30 = F_quote_identifier(m, v19)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
										F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v30 = F_quote_identifier(m, v19)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
									F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						} else {
							if l1 == int32(1) {
								v44 = int32(419283)
							} else {
								v44 = int32(542438)
							}
							if l1 == int32(-1) {
								v47 = int32(419375)
							} else {
								v47 = v44
							}
							F_ExplainPropertyText(m, int32(253537), v47, l2)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_ExplainPropertyText(m, int32(380012), v19, l2)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = F_get_rel_name(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(39997), v8)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errfinish(m, int32(493477), int32(4035), int32(376071))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = v15
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v20 == int32(0) {
					if l1 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						F_appendStringInfoString(m, v25, int32(419374))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v30 = F_quote_identifier(m, v19)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
								F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v30 = F_quote_identifier(m, v19)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
							F_appendStringInfo(m, v29, int32(185530), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					if l1 == int32(1) {
						v44 = int32(419283)
					} else {
						v44 = int32(542438)
					}
					if l1 == int32(-1) {
						v47 = int32(419375)
					} else {
						v47 = v44
					}
					F_ExplainPropertyText(m, int32(253537), v47, l2)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_ExplainPropertyText(m, int32(380012), v19, l2)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_ExplainOnePlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
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
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 float64
	_ = v525
	var v526 int32
	_ = v526
	var v527 float64
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v742 int32
	_ = v742
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int64
	_ = v841
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int64
	_ = v860
	var v864 int64
	_ = v864
	var v868 int64
	_ = v868
	var v871 int64
	_ = v871
	var v874 int32
	_ = v874
	var v875 int64
	_ = v875
	var v878 int64
	_ = v878
	var v881 int64
	_ = v881
	var v884 int64
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int64
	_ = v890
	var v893 int64
	_ = v893
	var v896 int32
	_ = v896
	var v897 int64
	_ = v897
	var v900 int64
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int64
	_ = v906
	var v909 int64
	_ = v909
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v916 int64
	_ = v916
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int64
	_ = v979
	var v980 int32
	_ = v980
	var v984 int64
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1035 int32
	_ = v1035
	var v1042 int64
	_ = v1042
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1085 int32
	_ = v1085
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1141 int32
	_ = v1141
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1168 int32
	_ = v1168
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1229 int32
	_ = v1229
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1280 int32
	_ = v1280
	var v1303 int32
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int64
	_ = v1333
	var v1334 int64
	_ = v1334
	var v1337 int64
	_ = v1337
	var v1338 int64
	_ = v1338
	var v1341 int64
	_ = v1341
	var v1342 int64
	_ = v1342
	var v1345 int64
	_ = v1345
	var v1346 int64
	_ = v1346
	var v1349 int64
	_ = v1349
	var v1350 int64
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int64
	_ = v1362
	var v1363 int64
	_ = v1363
	var v1366 int64
	_ = v1366
	var v1367 int64
	_ = v1367
	var v1370 int64
	_ = v1370
	var v1371 int64
	_ = v1371
	var v1374 int64
	_ = v1374
	var v1375 int64
	_ = v1375
	var v1378 int64
	_ = v1378
	var v1379 int64
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1412 int64
	_ = v1412
	var v1418 int64
	_ = v1418
	var v1427 int32
	_ = v1427
	var v1429 int64
	_ = v1429
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int64
	_ = v1446
	var v1450 int64
	_ = v1450
	var v1454 int64
	_ = v1454
	var v1457 int64
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1461 int64
	_ = v1461
	var v1464 int64
	_ = v1464
	var v1467 int64
	_ = v1467
	var v1470 int64
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int64
	_ = v1474
	var v1480 int64
	_ = v1480
	var v1483 int64
	_ = v1483
	var v1486 int64
	_ = v1486
	var v1489 int64
	_ = v1489
	var v1492 int64
	_ = v1492
	var v1495 int64
	_ = v1495
	var v1498 int64
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1520 int64
	_ = v1520
	var v1528 int32
	_ = v1528
	var v1531 int64
	_ = v1531
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1547 int32
	_ = v1547
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1564 int64
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1589 int64
	_ = v1589
	var v1591 int64
	_ = v1591
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	v10 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(256)
	m.G0 = v25
	v32 = F__emscripten_memset_bulkmem(m, v25-int32(-64), base.I32_extend8_s(v10), int32(144))
	mBase = m.M
	goto L1
L1:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v33 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v40 = v10
	goto L4
L4:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	F___clock_gettime(m, int32(1), v25+int32(208))
	mBase = m.M
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+216)))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v25)+208))
	v50 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	goto L8
L5:
	;
	v39 = int32(1)
	goto L7
L6:
	;
	v39 = int32(4)
	goto L7
L7:
	;
	v40 = v39
	goto L4
L8:
	;
	F_PushCopiedSnapshot(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = v40 | int32(2)
	goto L13
L12:
	;
	v56 = v40
	goto L13
L13:
	;
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v59 = v56 | int32(8)
	goto L16
L15:
	;
	v59 = v56
	goto L16
L16:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	goto L27
L19:
	;
	v62 = F_CreateIntoRelDestReceiver(m, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v64 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v69 = v62
	goto L18
L23:
	;
	v65 = F_CreateExplainSerializeDestReceiver(m, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[421]))
	v69 = v68
	goto L18
L26:
	;
	v69 = v65
	goto L18
L27:
	;
	v74 = F_CreateQueryDesc(m, l0, l3, v72, int32(0), v69, l4, l5, v59)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	v78 = v76 ^ int32(1)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = v78 | int32(2)
	goto L31
L30:
	;
	v82 = v78
	goto L31
L31:
	;
	if l1 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v126 != 0 {
		goto L49
	} else {
		goto L50
	}
L33:
	;
	F_ExecutorRun(m, v74, v103, int64(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L47
	}
L34:
	;
	v103 = int32(1)
	goto L33
L35:
	;
	F_ExecutorStart(m, v74, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v91 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v87 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v125 = float64(0)
	goto L32
L40:
	;
	F_ExecutorStart(m, v74, v92|v82)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L44
	}
L41:
	;
	v92 = int32(64)
	goto L43
L42:
	;
	v92 = int32(0)
	goto L43
L43:
	;
	goto L40
L44:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v97 != int32(1) {
		v125 = float64(0)
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v101 != 0 {
		v103 = int32(0)
		goto L33
	} else {
		goto L46
	}
L46:
	;
	goto L34
L47:
	;
	F_ExecutorFinish(m, v74)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F___clock_gettime(m, int32(1), v25+int32(208))
	mBase = m.M
	v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+216)))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v25)+208))
	v125 = base.F64_add(base.F64_div(base.F64_convert_i64_s(v113-v47+(v115-v48)*int64(1000000000)), float64(1e+09)), float64(0))
	goto L32
L49:
	;
	v128 = v25 - int32(-64)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v129 == int32(12) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	m.T0[v142].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L61
	}
L52:
	;
	goto L51
L53:
	;
	goto L57
L54:
	;
	goto L55
L55:
	;
	v140 = F__emscripten_memset_bulkmem(m, v128, base.I32_extend8_s(int32(0)), int32(144))
	mBase = m.M
	goto L60
L56:
	;
	goto L52
L57:
	;
	v135 = F__emscripten_memcpy_bulkmem(m, v128, v69-int32(-64), int32(144))
	mBase = m.M
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L52
L61:
	;
	F_ExplainOpenGroup(m, int32(17210), int32(0), int32(1), l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v150 = m.G0
	v152 = v150 - int32(80)
	m.G0 = v152
	*(*int32)(unsafe.Add(mBase, uint32(v152)+56)) = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	v163 = F_ExplainPreScanNode(m, v160, v152+int32(56))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v152)+56))
	v167 = m.G0
	v169 = v167 - int32(80)
	m.G0 = v169
	v176 = F__emscripten_memset_bulkmem(m, v169+int32(4), base.I32_extend8_s(int32(0)), int32(76))
	mBase = m.M
	goto L64
L64:
	;
	v177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+16)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v165
	F_set_rtable_names(m, v169, v177, v166)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v184 = int32(80)
	m.G0 = v169 + v184
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v183
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v189 = m.G0
	v191 = v189 - int32(16)
	m.G0 = v191
	v194 = F_palloc0(m, v184)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v194))) = v196
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v188)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188)+60))
	if v203 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_set_simple_column_names(m, v194)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L9
	} else {
		goto L80
	}
L68:
	;
	if v196 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+20)) = int32(0)
	goto L67
L71:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v210 = v204<<(uint(int32(2))%32) + int32(4)
	goto L73
L72:
	;
	v210 = int32(4)
	goto L73
L73:
	;
	v211 = F_palloc0(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+20)) = v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v188)+60))
	if v214 == int32(0) {
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v217 <= int32(0) {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v235 = int32(0)
	goto L77
L77:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v194)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v245 = int32(2)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v235<<(uint(v245)%32))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243+v249<<(uint(v245)%32)))) = v248
	v255 = v235 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v255 < v256 {
		v235 = v255
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L67
L79:
	;
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = v194
	v289 = F_list_make1_impl(m, int32(1), v191+int32(8))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	m.G0 = v191 + int32(16)
	v294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(l2)+44)) = v289
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v297 == v294 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	if v371 != int32(432) {
		v381 = v370
		goto L96
	} else {
		goto L97
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+56)) = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+56)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v304 <= int32(0) {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v307 = int32(0)
	if v307 < v304 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v310 = v304
	goto L89
L88:
	;
	v310 = v307
	goto L89
L89:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v322 = int32(0)
	goto L90
L90:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v311+v322<<(uint(int32(2))%32))))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v339 != int32(9) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+56)) = v302 - int32(1)
	goto L82
L92:
	;
	v343 = v322 + int32(1)
	if v310 != v343 {
		v322 = v343
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	goto L91
L95:
	;
	goto L82
L96:
	;
	v382 = int32(0)
	F_ExplainNode(m, v381, v382, v382, v382, l2)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L9
	} else {
		goto L99
	}
L97:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+81)))
	if v375 != int32(1) {
		v381 = v370
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v370)+36))
	v379 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+52)) = uint8(v379)
	v381 = v378
	goto L96
L99:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
	if v387 != int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v837 != int32(1) {
		goto L181
	} else {
		goto L182
	}
L101:
	;
	v390 = m.G0
	v392 = v390 - int32(16)
	m.G0 = v392
	v395 = v152 + int32(76)
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = int32(0)
	v399 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)+412))
	if v403 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v471 = F_palloc(m, v468<<(uint(int32(2))%32))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L9
	} else {
		goto L106
	}
L103:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+376))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)+364))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v401)+352))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401)+340))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)+328))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v401)+316))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v401)+304))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)+292))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v401)+280))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v401)+268))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v401)+256))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v401)+244))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v401)+232))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v401)+220))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v401)+208))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v401)+196))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v401)+184))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v401)+172))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v401)+160))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v401)+148))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v401)+136))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v401)+124))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v401)+112))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v401)+100))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v401)+88))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v401)+76))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v401-int32(-64))))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v401)+52))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v401)+40))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v401)+28))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v401)+16))
	v468 = v404 + (v405 + (v406 + (v407 + (v408 + (v409 + (v410 + (v411 + (v412 + (v413 + (v414 + (v415 + (v416 + (v417 + (v418 + (v419 + (v420 + (v421 + (v422 + (v423 + (v424 + (v425 + (v426 + (v427 + (v428 + (v429 + (v432 + (v433 + (v434 + (v435 + (v436 + v402))))))))))))))))))))))))))))))
	goto L105
L104:
	;
	v468 = v402
	goto L105
L105:
	;
	goto L102
L106:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	if v474 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	m.G0 = v392 + int32(16)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v621 != 0 {
		goto L144
	} else {
		goto L145
	}
L108:
	;
	if v474 == int32(4484956) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v495 = v474
	goto L110
L110:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495-int32(44)))))
	if v503&int32(32) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L107
L112:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v593 != int32(4484956) {
		v495 = v593
		goto L110
	} else {
		goto L143
	}
L113:
	;
	v509 = v495 + int32(-64)
	v510 = F_ConfigOptionIsVisible(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	if v510 == int32(0) {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v515 = v495 - int32(40)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	switch v516 {
	case 0:
		goto L123
	case 1:
		goto L122
	case 2:
		goto L121
	case 3:
		goto L120
	case 4:
		goto L119
	default:
		goto L118
	}
L116:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	*(*int32)(unsafe.Add(mBase, uint32(v471+v581<<(uint(int32(2))%32)))) = v509
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v586 + int32(1)
	goto L112
L117:
	;
	if v530 == int32(0) {
		goto L116
	} else {
		goto L133
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L130
	}
L119:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	if v532 != v534 {
		goto L116
	} else {
		goto L129
	}
L120:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	if v531 != 0 {
		goto L117
	} else {
		goto L127
	}
L121:
	;
	v525 = *(*float64)(unsafe.Add(mBase, uint32(v495)+32))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v527 = *(*float64)(unsafe.Add(mBase, uint32(v526)))
	if base.F64_ne(v525, v527) != 0 {
		goto L116
	} else {
		goto L126
	}
L122:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v495)+32))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v521 != v523 {
		goto L116
	} else {
		goto L125
	}
L123:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+32)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v495)+28))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v517 != v519 {
		goto L116
	} else {
		goto L124
	}
L124:
	;
	goto L112
L125:
	;
	goto L112
L126:
	;
	goto L112
L127:
	;
	if v530 != 0 {
		goto L116
	} else {
		goto L128
	}
L128:
	;
	goto L112
L129:
	;
	goto L112
L130:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v392))) = v540
	F_errmsg_internal(m, int32(483107), v392)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L9
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(496556), int32(5420), int32(136714))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L9
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	if v555 == int32(0) {
		v574 = v554
		v575 = v555
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v575-v574 == int32(0) {
		goto L112
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	if v554 != v555 {
		v574 = v554
		v575 = v555
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v559 = v531
	v560 = v530
	goto L138
L138:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+1)))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+1)))
	if v564 == int32(0) {
		v574 = v563
		v575 = v564
		goto L135
	} else {
		goto L140
	}
L139:
	;
	v574 = v563
	v575 = v564
	goto L135
L140:
	;
	v567 = int32(1)
	if v563 == v564 {
		v559 = v559 + v567
		v560 = v560 + v567
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	goto L116
L143:
	;
	goto L111
L144:
	;
	v622 = int32(155215)
	F_ExplainOpenGroup(m, v622, v622, int32(1), l2)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L9
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v695 <= int32(0) {
		goto L100
	} else {
		goto L157
	}
L147:
	;
	v627 = int32(0)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v627 < v628 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v640 = v627
	goto L151
L149:
	;
	goto L150
L150:
	;
	F_ExplainCloseGroup(m, int32(155215), int32(1), l2)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L9
	} else {
		goto L156
	}
L151:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v471+v640<<(uint(int32(2))%32))))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v660 = F_GetConfigOptionByName(m, v657, int32(0), int32(1))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L9
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	F_ExplainPropertyText(m, v662, v660, l2)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L9
	} else {
		goto L154
	}
L154:
	;
	v666 = v640 + int32(1)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v666 < v667 {
		v640 = v666
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	goto L100
L157:
	;
	F_initStringInfo(m, v152+int32(60))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v702 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v152)+60))
	F_ExplainPropertyText(m, int32(155215), v812, l2)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L9
	} else {
		goto L180
	}
L160:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v709 = F_GetConfigOptionByName(m, v706, int32(0), int32(1))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	if v709 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v729 = int32(1)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v730 <= v729 {
		goto L159
	} else {
		goto L168
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+52)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v152)+48)) = v711
	F_appendStringInfo(m, v152+int32(60), int32(668599), v152+int32(48))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L9
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+32)) = v711
	F_appendStringInfo(m, v152+int32(60), int32(529775), v152+int32(32))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L9
	} else {
		goto L167
	}
L166:
	;
	goto L162
L167:
	;
	goto L162
L168:
	;
	v742 = v729
	goto L169
L169:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v471+v742<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v152+int32(60), int32(727439))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L171
	}
L170:
	;
	goto L159
L171:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v767 = F_GetConfigOptionByName(m, v764, int32(0), int32(1))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	if v767 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v786 = v742 + int32(1)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v786 < v787 {
		v742 = v786
		goto L169
	} else {
		goto L179
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+20)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v152)+16)) = v769
	F_appendStringInfo(m, v152+int32(60), int32(668599), v152+int32(16))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L9
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v769
	F_appendStringInfo(m, v152+int32(60), int32(529775), v152)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L9
	} else {
		goto L178
	}
L177:
	;
	goto L173
L178:
	;
	goto L173
L179:
	;
	goto L170
L180:
	;
	goto L100
L181:
	;
	m.G0 = v152 + int32(80)
	if l7 != 0 {
		goto L189
	} else {
		goto L190
	}
L182:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v840)+8))
	if v841 == int64(0) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v845 = *(*int32)(unsafe.Add(mBase, _consts[424]))
	if v845 == int32(3) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	F_ExplainPropertyInteger(m, int32(221367), int32(0), v841, l2)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	if l6 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L187:
	;
	v952 = int32(333131)
	F_ExplainOpenGroup(m, v952, v952, int32(1), l2)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L9
	} else {
		goto L216
	}
L188:
	;
	v946 = base.B2i32(l8 != int32(0))
	goto L187
L189:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v857 != 0 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	v927 = int32(0)
	goto L191
L191:
	;
	v933 = base.B2i32(l8 != int32(0))
	if l8 != 0 {
		v946 = v933
		goto L187
	} else {
		goto L214
	}
L192:
	;
	v858 = int32(1)
	v860 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	if int64(0) < v860 {
		v874 = v858
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v875 = *(*int64)(unsafe.Add(mBase, uint32(l7)+32))
	if int64(0) < v875 {
		v887 = v858
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v864 = *(*int64)(unsafe.Add(mBase, uint32(l7)+8))
	if int64(0) < v864 {
		v874 = int32(1)
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v868 = *(*int64)(unsafe.Add(mBase, uint32(l7)+16))
	if int64(0) < v868 {
		v874 = int32(1)
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v871 = *(*int64)(unsafe.Add(mBase, uint32(l7)+24))
	v874 = base.B2i32(int64(0) < v871)
	goto L193
L197:
	;
	v888 = int32(1)
	v890 = *(*int64)(unsafe.Add(mBase, uint32(l7)+64))
	if v890 <= int64(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	v878 = *(*int64)(unsafe.Add(mBase, uint32(l7)+40))
	if int64(0) < v878 {
		v887 = v858
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v881 = *(*int64)(unsafe.Add(mBase, uint32(l7)+48))
	if int64(0) < v881 {
		v887 = v858
		goto L197
	} else {
		goto L200
	}
L200:
	;
	v884 = *(*int64)(unsafe.Add(mBase, uint32(l7)+56))
	v887 = base.B2i32(int64(0) < v884)
	goto L197
L201:
	;
	v893 = *(*int64)(unsafe.Add(mBase, uint32(l7)+72))
	v896 = base.B2i32(int64(0) < v893)
	goto L203
L202:
	;
	v896 = v888
	goto L203
L203:
	;
	v897 = *(*int64)(unsafe.Add(mBase, uint32(l7)+80))
	if v897 == int64(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v900 = *(*int64)(unsafe.Add(mBase, uint32(l7)+88))
	v903 = base.B2i32(v900 != int64(0))
	goto L206
L205:
	;
	v903 = v888
	goto L206
L206:
	;
	v904 = int32(1)
	v906 = *(*int64)(unsafe.Add(mBase, uint32(l7)+96))
	if v906 == int64(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v909 = *(*int64)(unsafe.Add(mBase, uint32(l7)+104))
	v912 = base.B2i32(v909 != int64(0))
	goto L209
L208:
	;
	v912 = v904
	goto L209
L209:
	;
	v913 = *(*int64)(unsafe.Add(mBase, uint32(l7)+112))
	if v913 == int64(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v916 = *(*int64)(unsafe.Add(mBase, uint32(l7)+120))
	v919 = base.B2i32(v916 != int64(0))
	goto L212
L211:
	;
	v919 = v904
	goto L212
L212:
	;
	if (v874|v887|v896|v903)&int32(1) != 0 {
		goto L188
	} else {
		goto L213
	}
L213:
	;
	v927 = v919 | v912
	goto L191
L214:
	;
	if v927&int32(1) == int32(0) {
		goto L186
	} else {
		goto L215
	}
L215:
	;
	v946 = v933
	goto L187
L216:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v957 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L9
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	if l7 != 0 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoString(m, v962, int32(735508))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L9
	} else {
		goto L221
	}
L221:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v966 + int32(1)
	goto L219
L222:
	;
	F_show_buffer_usage(m, l2, l7)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L9
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v946 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v1013 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L227:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	v976 = v974 + int32(1023)
	v977 = int32(10)
	v979 = base.I64_extend_i32_u(int32(base.Ui32(v976) >> (uint(v977) % 32)))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v984 = base.I64_extend_i32_u(int32(base.Ui32(v976-v980) >> (uint(v977) % 32)))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v985 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L9
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	F_ExplainPropertyInteger(m, int32(447369), int32(541476), v984, l2)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L9
	} else {
		goto L234
	}
L231:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v979
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v984
	F_appendStringInfo(m, v990, int32(541301), v25+int32(48))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L9
	} else {
		goto L232
	}
L232:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v998, int32(10))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	goto L226
L234:
	;
	F_ExplainPropertyInteger(m, int32(446110), int32(541476), v979, l2)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L9
	} else {
		goto L235
	}
L235:
	;
	goto L226
L236:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v1016 - int32(1)
	goto L238
L237:
	;
	goto L238
L238:
	;
	F_ExplainCloseGroup(m, int32(333131), int32(1), l2)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L9
	} else {
		goto L239
	}
L239:
	;
	goto L186
L240:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v1051 == int32(1) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v1035&int32(1) == int32(0) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(l6)))
	F_ExplainPropertyFloat(m, int32(373836), int32(150882), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v1042), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L9
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	v1054 = int32(0)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+84))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+80))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+72))
	v1059 = int32(134190)
	F_ExplainOpenGroup(m, v1059, v1059, v1054, l2)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L9
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v1303 != int32(1) {
		goto L275
	} else {
		goto L276
	}
L247:
	;
	if v1058 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	if v1056 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L249:
	;
	v1157 = int32(0)
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1158 <= v1157 {
		v1200 = v1141
		goto L248
	} else {
		goto L262
	}
L250:
	;
	if v1057 != 0 {
		v1141 = int32(1)
		goto L249
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1070 = int32(0)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	v1075 = base.B2i32(v1056|v1057 != v1070) | base.B2i32(int32(1) < v1072)
	if v1070 < v1072 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1200 = base.B2i32(v1056 != int32(0))
	goto L248
L254:
	;
	v1085 = v1054
	goto L257
L255:
	;
	goto L256
L256:
	;
	if v1057 == int32(0) {
		v1200 = v1075
		goto L248
	} else {
		goto L261
	}
L257:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+12))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1100+v1085<<(uint(int32(2))%32))))
	F_report_triggers(m, v1104, v1075, l2)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L9
	} else {
		goto L259
	}
L258:
	;
	goto L256
L259:
	;
	v1108 = v1085 + int32(1)
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	if v1108 < v1109 {
		v1085 = v1108
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1141 = v1075
	goto L249
L262:
	;
	v1168 = v1157
	goto L263
L263:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1183+v1168<<(uint(int32(2))%32))))
	F_report_triggers(m, v1187, v1141, l2)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L9
	} else {
		goto L265
	}
L264:
	;
	v1200 = v1141
	goto L248
L265:
	;
	v1191 = v1168 + int32(1)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1191 < v1192 {
		v1168 = v1191
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	F_ExplainCloseGroup(m, int32(134190), int32(0), l2)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L9
	} else {
		goto L274
	}
L268:
	;
	v1218 = int32(0)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	if v1219 <= v1218 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v1229 = v1218
	goto L270
L270:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+12))
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1244+v1229<<(uint(int32(2))%32))))
	F_report_triggers(m, v1248, v1200, l2)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L9
	} else {
		goto L272
	}
L271:
	;
	goto L267
L272:
	;
	v1252 = v1229 + int32(1)
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	if v1252 < v1253 {
		v1229 = v1252
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	goto L246
L275:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1391 != 0 {
		goto L287
	} else {
		goto L288
	}
L276:
	;
	v1306 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+248)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v25)+240)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v25)+232)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v25)+224)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v25)+216)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v25)+208)) = v1306
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1318)+176)))
	if v1319&int32(1) == int32(0) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+180))
	if v1324 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1326 = v25 + int32(208)
	v1328 = v1324 + int32(8)
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1328)))
	*(*int32)(unsafe.Add(mBase, uint32(v1326))) = v1329 + v1330
	v1333 = *(*int64)(unsafe.Add(mBase, uint32(v1326)+8))
	v1334 = *(*int64)(unsafe.Add(mBase, uint32(v1328)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1326)+8)) = v1333 + v1334
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v1326)+16))
	v1338 = *(*int64)(unsafe.Add(mBase, uint32(v1328)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1326)+16)) = v1337 + v1338
	v1341 = *(*int64)(unsafe.Add(mBase, uint32(v1326)+24))
	v1342 = *(*int64)(unsafe.Add(mBase, uint32(v1328)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1326)+24)) = v1341 + v1342
	v1345 = *(*int64)(unsafe.Add(mBase, uint32(v1326)+32))
	v1346 = *(*int64)(unsafe.Add(mBase, uint32(v1328)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1326)+32)) = v1345 + v1346
	v1349 = *(*int64)(unsafe.Add(mBase, uint32(v1326)+40))
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v1328)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1326)+40)) = v1349 + v1350
	goto L281
L279:
	;
	v1354 = v1318
	goto L280
L280:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+184))
	if v1355 != 0 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v1354 = v1353
	goto L280
L282:
	;
	v1357 = v25 + int32(208)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1357)))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	*(*int32)(unsafe.Add(mBase, uint32(v1357))) = v1358 + v1359
	v1362 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+8))
	v1363 = *(*int64)(unsafe.Add(mBase, uint32(v1355)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1357)+8)) = v1362 + v1363
	v1366 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+16))
	v1367 = *(*int64)(unsafe.Add(mBase, uint32(v1355)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1357)+16)) = v1366 + v1367
	v1370 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+24))
	v1371 = *(*int64)(unsafe.Add(mBase, uint32(v1355)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1357)+24)) = v1370 + v1371
	v1374 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+32))
	v1375 = *(*int64)(unsafe.Add(mBase, uint32(v1355)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1357)+32)) = v1374 + v1375
	v1378 = *(*int64)(unsafe.Add(mBase, uint32(v1357)+40))
	v1379 = *(*int64)(unsafe.Add(mBase, uint32(v1355)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1357)+40)) = v1378 + v1379
	goto L285
L283:
	;
	v1383 = v1354
	goto L284
L284:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+176))
	F_ExplainPrintJIT(m, l2, v1384, v25+int32(208))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L286
	}
L285:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v1383 = v1382
	goto L284
L286:
	;
	goto L275
L287:
	;
	v1392 = int32(257430)
	F_ExplainOpenGroup(m, v1392, v1392, int32(1), l2)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L9
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	if v1557 != 0 {
		goto L335
	} else {
		goto L336
	}
L290:
	;
	if v1391 == int32(1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1401 = int32(63226)
	goto L293
L292:
	;
	v1401 = int32(17797)
	goto L293
L293:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v1402 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	F_ExplainCloseGroup(m, int32(257430), int32(1), l2)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L9
	} else {
		goto L334
	}
L295:
	;
	F_ExplainIndentText(m, l2)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L9
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v1515 == int32(1) {
		goto L326
	} else {
		goto L327
	}
L298:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v1408 == int32(1) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v1440 != int32(1) {
		goto L294
	} else {
		goto L305
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v1401
	v1412 = *(*int64)(unsafe.Add(mBase, uint32(v25)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(base.Ui64(v1412+int64(1023)) >> (uint(int64(10)) % 64))
	v1418 = *(*int64)(unsafe.Add(mBase, uint32(v25)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v25))) = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v1418), float64(1e+09)), float64(1000))
	F_appendStringInfo(m, v1407, int32(729220), v25)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L9
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v1401
	v1429 = *(*int64)(unsafe.Add(mBase, uint32(v25)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = int64(base.Ui64(v1429+int64(1023)) >> (uint(int64(10)) % 64))
	F_appendStringInfo(m, v1407, int32(729179), v25+int32(32))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L9
	} else {
		goto L304
	}
L303:
	;
	goto L299
L304:
	;
	goto L299
L305:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v1443 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v1503 + int32(1)
	F_show_buffer_usage(m, l2, v25+int32(80))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L9
	} else {
		goto L325
	}
L307:
	;
	v1444 = int32(1)
	v1446 = *(*int64)(unsafe.Add(mBase, uint32(v25)+80))
	if int64(0) < v1446 {
		v1460 = v1444
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1461 = *(*int64)(unsafe.Add(mBase, uint32(v25)+112))
	if int64(0) < v1461 {
		v1473 = v1444
		goto L312
	} else {
		goto L313
	}
L309:
	;
	v1450 = *(*int64)(unsafe.Add(mBase, uint32(v25)+88))
	if int64(0) < v1450 {
		v1460 = int32(1)
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1454 = *(*int64)(unsafe.Add(mBase, uint32(v25)+96))
	if int64(0) < v1454 {
		v1460 = int32(1)
		goto L308
	} else {
		goto L311
	}
L311:
	;
	v1457 = *(*int64)(unsafe.Add(mBase, uint32(v25)+104))
	v1460 = base.B2i32(int64(0) < v1457)
	goto L308
L312:
	;
	v1474 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	if v1474 != int64(0) {
		goto L306
	} else {
		goto L316
	}
L313:
	;
	v1464 = *(*int64)(unsafe.Add(mBase, uint32(v25)+120))
	if int64(0) < v1464 {
		v1473 = v1444
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1467 = *(*int64)(unsafe.Add(mBase, uint32(v25)+128))
	if int64(0) < v1467 {
		v1473 = v1444
		goto L312
	} else {
		goto L315
	}
L315:
	;
	v1470 = *(*int64)(unsafe.Add(mBase, uint32(v25)+136))
	v1473 = base.B2i32(int64(0) < v1470)
	goto L312
L316:
	;
	if (v1460|v1473)&int32(1) != 0 {
		goto L306
	} else {
		goto L317
	}
L317:
	;
	v1480 = *(*int64)(unsafe.Add(mBase, uint32(v25)+144))
	if int64(0) < v1480 {
		goto L306
	} else {
		goto L318
	}
L318:
	;
	v1483 = *(*int64)(unsafe.Add(mBase, uint32(v25)+152))
	if int64(0) < v1483 {
		goto L306
	} else {
		goto L319
	}
L319:
	;
	v1486 = *(*int64)(unsafe.Add(mBase, uint32(v25)+160))
	if v1486 != int64(0) {
		goto L306
	} else {
		goto L320
	}
L320:
	;
	v1489 = *(*int64)(unsafe.Add(mBase, uint32(v25)+168))
	if v1489 != int64(0) {
		goto L306
	} else {
		goto L321
	}
L321:
	;
	v1492 = *(*int64)(unsafe.Add(mBase, uint32(v25)+176))
	if v1492 != int64(0) {
		goto L306
	} else {
		goto L322
	}
L322:
	;
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v25)+184))
	if v1495 != int64(0) {
		goto L306
	} else {
		goto L323
	}
L323:
	;
	v1498 = *(*int64)(unsafe.Add(mBase, uint32(v25)+200))
	if v1498 == int64(0) {
		goto L294
	} else {
		goto L324
	}
L324:
	;
	goto L306
L325:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v1511 - int32(1)
	goto L294
L326:
	;
	v1520 = *(*int64)(unsafe.Add(mBase, uint32(v25)+72))
	F_ExplainPropertyFloat(m, int32(373968), int32(150882), base.F64_mul(base.F64_div(base.F64_convert_i64_s(v1520), float64(1e+09)), float64(1000)), int32(3), l2)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L9
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v1531 = *(*int64)(unsafe.Add(mBase, uint32(v25)+64))
	F_ExplainPropertyUInteger(m, int32(372171), int32(541476), int64(base.Ui64(v1531+int64(1023))>>(uint(int64(10))%64)), l2)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L9
	} else {
		goto L330
	}
L329:
	;
	goto L328
L330:
	;
	F_ExplainPropertyText(m, int32(111706), v1401, l2)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L9
	} else {
		goto L331
	}
L331:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v1541 != int32(1) {
		goto L294
	} else {
		goto L332
	}
L332:
	;
	F_show_buffer_usage(m, l2, v25+int32(80))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L9
	} else {
		goto L333
	}
L333:
	;
	goto L294
L334:
	;
	goto L289
L335:
	;
	m.T0[v1557].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L9
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	F___clock_gettime(m, int32(1), v25+int32(208))
	mBase = m.M
	v1564 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+216)))
	v1565 = *(*int64)(unsafe.Add(mBase, uint32(v25)+208))
	F_ExecutorEnd(m, v74)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L9
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	F_FreeQueryDesc(m, v74)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L9
	} else {
		goto L340
	}
L340:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L9
	} else {
		goto L341
	}
L341:
	;
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v1572 == int32(1) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L9
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1577 = int32(1)
	F___clock_gettime(m, v1577, v25+int32(208))
	mBase = m.M
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
	if v1581 != v1577 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	goto L344
L346:
	;
	F_ExplainCloseGroup(m, int32(17210), int32(1), l2)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L9
	} else {
		goto L350
	}
L347:
	;
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v1584 != int32(1) {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1589 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+216)))
	v1591 = *(*int64)(unsafe.Add(mBase, uint32(v25)+208))
	F_ExplainPropertyFloat(m, int32(373803), int32(150882), base.F64_mul(base.F64_add(v125, base.F64_div(base.F64_convert_i64_s(v1589-v1564+(v1591-v1565)*int64(1000000000)), float64(1e+09))), float64(1000)), int32(3), l2)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L9
	} else {
		goto L349
	}
L349:
	;
	goto L346
L350:
	;
	m.G0 = v25 + int32(256)
	return
}
