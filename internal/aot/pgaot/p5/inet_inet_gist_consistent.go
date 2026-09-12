package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
		if v23 == v21 {
			return int32(1)
		} else {
			v28 = int32(1)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v33 = v31 & v28
			if v33 != 0 {
				v34 = v28
			} else {
				v34 = int32(4)
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v34))))
			if v36 != v23 {
				switch v18&int32(65535) - int32(19) {
				case 0:
					v787 = v28
					return v787
				case 1, 2:
					if base.Ui32(v36) <= base.Ui32(v23) {
						return int32(0)
					} else {
						v787 = v28
						return v787
					}
				case 3, 4:
					if base.Ui32(v36) < base.Ui32(v23) {
						v787 = v28
						return v787
					} else {
						return int32(0)
					}
				default:
					return int32(0)
				}
			} else {
				v44 = v14 + int32(1)
				v46 = v14 + int32(4)
				if v33 != 0 {
					v47 = v44
				} else {
					v47 = v46
				}
				switch v18&int32(65535) - int32(18) {
				case 0, 9:
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					if base.Ui32(v78) <= base.Ui32(v79) {
						v90 = v78
						v91 = v79
						v93 = v19 + int32(4)
						v95 = v47 + int32(2)
						v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						if base.Ui32(v96) < base.Ui32(v90) {
							v98 = v96
						} else {
							v98 = v90
						}
						if base.Ui32(v98) < base.Ui32(v91) {
							v100 = v98
						} else {
							v100 = v91
						}
						v105 = base.I32_div_s(v100, int32(8))
						v106 = F_memcmp(m, v93, v95, v105)
						mBase = m.M
						if v106 != 0 {
							v192 = v106
							v202 = v192
						} else {
							v107 = int32(0)
							v110 = v100 - v105<<(uint(int32(3))%32)
							if v110 <= v107 {
								v192 = v107
								v202 = v192
							} else {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v105))))
								v115 = int32(128)
								v116 = v114 & v115
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v105))))
								if v116 != v118&v115 {
									v193 = v116
									if v193 != 0 {
										v196 = int32(1)
									} else {
										v196 = int32(-1)
									}
									v202 = v196
								} else {
									if v110 == int32(1) {
										v192 = v107
										v202 = v192
									} else {
										v124 = int32(1)
										v126 = int32(128)
										v127 = v114 << (uint(v124) % 32) & v126
										if v127 != v118<<(uint(v124)%32)&v126 {
											v193 = v127
											if v193 != 0 {
												v196 = int32(1)
											} else {
												v196 = int32(-1)
											}
											v202 = v196
										} else {
											if v110 < int32(3) {
												v192 = v107
												v202 = v192
											} else {
												v135 = int32(2)
												v137 = int32(128)
												v138 = v114 << (uint(v135) % 32) & v137
												if v138 != v118<<(uint(v135)%32)&v137 {
													v193 = v138
													if v193 != 0 {
														v196 = int32(1)
													} else {
														v196 = int32(-1)
													}
													v202 = v196
												} else {
													if v110 == int32(3) {
														v192 = v107
														v202 = v192
													} else {
														v146 = int32(3)
														v148 = int32(128)
														v149 = v114 << (uint(v146) % 32) & v148
														if v149 != v118<<(uint(v146)%32)&v148 {
															v193 = v149
															if v193 != 0 {
																v196 = int32(1)
															} else {
																v196 = int32(-1)
															}
															v202 = v196
														} else {
															if v110 < int32(5) {
																v192 = v107
																v202 = v192
															} else {
																v157 = int32(4)
																v159 = int32(128)
																v160 = v114 << (uint(v157) % 32) & v159
																if v160 != v118<<(uint(v157)%32)&v159 {
																	v193 = v160
																	if v193 != 0 {
																		v196 = int32(1)
																	} else {
																		v196 = int32(-1)
																	}
																	v202 = v196
																} else {
																	if v110 == int32(5) {
																		v192 = v107
																		v202 = v192
																	} else {
																		v168 = int32(5)
																		v170 = int32(128)
																		v171 = v114 << (uint(v168) % 32) & v170
																		if v171 != v118<<(uint(v168)%32)&v170 {
																			v193 = v171
																			if v193 != 0 {
																				v196 = int32(1)
																			} else {
																				v196 = int32(-1)
																			}
																			v202 = v196
																		} else {
																			if v110 < int32(7) {
																				v192 = v107
																				v202 = v192
																			} else {
																				v179 = int32(6)
																				v181 = int32(128)
																				v182 = v114 << (uint(v179) % 32) & v181
																				if v182 != v118<<(uint(v179)%32)&v181 {
																					v193 = v182
																					if v193 != 0 {
																						v196 = int32(1)
																					} else {
																						v196 = int32(-1)
																					}
																					v202 = v196
																				} else {
																					v192 = v107
																					v202 = v192
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
								}
							}
						}
						switch v18&int32(65535) - int32(3) {
						case 0, 21, 22, 23, 24:
							v321 = v202
							return base.B2i32(v321 == int32(0))
						default:
							v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
							v638 = v635 & int32(1)
							if v638 != 0 {
								v642 = v44
							} else {
								v642 = v46
							}
							v644 = v642 + int32(2)
							v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
							if v647 == int32(3) {
								v650 = int32(128)
							} else {
								v650 = int32(32)
							}
							v655 = base.I32_div_s(v650, int32(8))
							v656 = F_memcmp(m, v93, v644, v655)
							mBase = m.M
							if v656 != 0 {
								v742 = v656
								v752 = v742
							} else {
								v657 = int32(0)
								v660 = v650 - v655<<(uint(int32(3))%32)
								if v660 <= v657 {
									v742 = v657
									v752 = v742
								} else {
									v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
									v665 = int32(128)
									v666 = v664 & v665
									v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
									if v666 != v668&v665 {
										v743 = v666
										if v743 != 0 {
											v746 = int32(1)
										} else {
											v746 = int32(-1)
										}
										v752 = v746
									} else {
										if v660 == int32(1) {
											v742 = v657
											v752 = v742
										} else {
											v674 = int32(1)
											v676 = int32(128)
											v677 = v664 << (uint(v674) % 32) & v676
											if v677 != v668<<(uint(v674)%32)&v676 {
												v743 = v677
												if v743 != 0 {
													v746 = int32(1)
												} else {
													v746 = int32(-1)
												}
												v752 = v746
											} else {
												if v660 < int32(3) {
													v742 = v657
													v752 = v742
												} else {
													v685 = int32(2)
													v687 = int32(128)
													v688 = v664 << (uint(v685) % 32) & v687
													if v688 != v668<<(uint(v685)%32)&v687 {
														v743 = v688
														if v743 != 0 {
															v746 = int32(1)
														} else {
															v746 = int32(-1)
														}
														v752 = v746
													} else {
														if v660 == int32(3) {
															v742 = v657
															v752 = v742
														} else {
															v696 = int32(3)
															v698 = int32(128)
															v699 = v664 << (uint(v696) % 32) & v698
															if v699 != v668<<(uint(v696)%32)&v698 {
																v743 = v699
																if v743 != 0 {
																	v746 = int32(1)
																} else {
																	v746 = int32(-1)
																}
																v752 = v746
															} else {
																if v660 < int32(5) {
																	v742 = v657
																	v752 = v742
																} else {
																	v707 = int32(4)
																	v709 = int32(128)
																	v710 = v664 << (uint(v707) % 32) & v709
																	if v710 != v668<<(uint(v707)%32)&v709 {
																		v743 = v710
																		if v743 != 0 {
																			v746 = int32(1)
																		} else {
																			v746 = int32(-1)
																		}
																		v752 = v746
																	} else {
																		if v660 == int32(5) {
																			v742 = v657
																			v752 = v742
																		} else {
																			v718 = int32(5)
																			v720 = int32(128)
																			v721 = v664 << (uint(v718) % 32) & v720
																			if v721 != v668<<(uint(v718)%32)&v720 {
																				v743 = v721
																				if v743 != 0 {
																					v746 = int32(1)
																				} else {
																					v746 = int32(-1)
																				}
																				v752 = v746
																			} else {
																				if v660 < int32(7) {
																					v742 = v657
																					v752 = v742
																				} else {
																					v729 = int32(6)
																					v731 = int32(128)
																					v732 = v664 << (uint(v729) % 32) & v731
																					if v732 != v668<<(uint(v729)%32)&v731 {
																						v743 = v732
																						if v743 != 0 {
																							v746 = int32(1)
																						} else {
																							v746 = int32(-1)
																						}
																						v752 = v746
																					} else {
																						v742 = v657
																						v752 = v742
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
									}
								}
							}
							switch v18&int32(65535) - int32(18) {
							case 0:
								v802 = v752
								return base.B2i32(v802 == int32(0))
							case 1:
								v795 = v752
								return base.B2i32(v795 != int32(0))
							case 2:
								return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
							case 3:
								return base.B2i32(v752 <= int32(0))
							case 4:
								return base.B2i32(int32(0) < v752)
							case 5:
								return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v774 = m.ExcPending
								if v774 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(518516), int32(0))
									mBase = m.M
									v778 = m.ExcPending
									if v778 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(494114), int32(327), int32(92161))
										mBase = m.M
										v783 = m.ExcPending
										if v783 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						case 15:
							if v202 != 0 {
								v787 = int32(0)
								return v787
							} else {
								v343 = int32(1)
								v344 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+16)))
								v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345)+12)))
								if v347&v343 == int32(0) {
									v787 = v343
									return v787
								} else {
									v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									if v354&int32(1) != 0 {
										v357 = v44
									} else {
										v357 = v46
									}
									v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
									if v353 != v358 {
										v787 = int32(0)
										return v787
									} else {
										v361 = v357 + int32(2)
										v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
										if v364 == int32(3) {
											v367 = int32(128)
										} else {
											v367 = int32(32)
										}
										v372 = base.I32_div_s(v367, int32(8))
										v373 = F_memcmp(m, v93, v361, v372)
										mBase = m.M
										if v373 != 0 {
											v459 = v373
											v469 = v459
										} else {
											v374 = int32(0)
											v377 = v367 - v372<<(uint(int32(3))%32)
											if v377 <= v374 {
												v459 = v374
												v469 = v459
											} else {
												v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v372))))
												v382 = int32(128)
												v383 = v381 & v382
												v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v372))))
												if v383 != v385&v382 {
													v460 = v383
													if v460 != 0 {
														v463 = int32(1)
													} else {
														v463 = int32(-1)
													}
													v469 = v463
												} else {
													if v377 == int32(1) {
														v459 = v374
														v469 = v459
													} else {
														v391 = int32(1)
														v393 = int32(128)
														v394 = v381 << (uint(v391) % 32) & v393
														if v394 != v385<<(uint(v391)%32)&v393 {
															v460 = v394
															if v460 != 0 {
																v463 = int32(1)
															} else {
																v463 = int32(-1)
															}
															v469 = v463
														} else {
															if v377 < int32(3) {
																v459 = v374
																v469 = v459
															} else {
																v402 = int32(2)
																v404 = int32(128)
																v405 = v381 << (uint(v402) % 32) & v404
																if v405 != v385<<(uint(v402)%32)&v404 {
																	v460 = v405
																	if v460 != 0 {
																		v463 = int32(1)
																	} else {
																		v463 = int32(-1)
																	}
																	v469 = v463
																} else {
																	if v377 == int32(3) {
																		v459 = v374
																		v469 = v459
																	} else {
																		v413 = int32(3)
																		v415 = int32(128)
																		v416 = v381 << (uint(v413) % 32) & v415
																		if v416 != v385<<(uint(v413)%32)&v415 {
																			v460 = v416
																			if v460 != 0 {
																				v463 = int32(1)
																			} else {
																				v463 = int32(-1)
																			}
																			v469 = v463
																		} else {
																			if v377 < int32(5) {
																				v459 = v374
																				v469 = v459
																			} else {
																				v424 = int32(4)
																				v426 = int32(128)
																				v427 = v381 << (uint(v424) % 32) & v426
																				if v427 != v385<<(uint(v424)%32)&v426 {
																					v460 = v427
																					if v460 != 0 {
																						v463 = int32(1)
																					} else {
																						v463 = int32(-1)
																					}
																					v469 = v463
																				} else {
																					if v377 == int32(5) {
																						v459 = v374
																						v469 = v459
																					} else {
																						v435 = int32(5)
																						v437 = int32(128)
																						v438 = v381 << (uint(v435) % 32) & v437
																						if v438 != v385<<(uint(v435)%32)&v437 {
																							v460 = v438
																							if v460 != 0 {
																								v463 = int32(1)
																							} else {
																								v463 = int32(-1)
																							}
																							v469 = v463
																						} else {
																							if v377 < int32(7) {
																								v459 = v374
																								v469 = v459
																							} else {
																								v446 = int32(6)
																								v448 = int32(128)
																								v449 = v381 << (uint(v446) % 32) & v448
																								if v449 != v385<<(uint(v446)%32)&v448 {
																									v460 = v449
																									if v460 != 0 {
																										v463 = int32(1)
																									} else {
																										v463 = int32(-1)
																									}
																									v469 = v463
																								} else {
																									v459 = v374
																									v469 = v459
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
												}
											}
										}
										v802 = v469
										return base.B2i32(v802 == int32(0))
									}
								}
							}
						case 16:
							if v202 != 0 {
								v787 = v28
								return v787
							} else {
								v502 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v502)+16)))
								v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v503)+12)))
								if v505&int32(1) == int32(0) {
									v787 = v28
									return v787
								} else {
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									if v511&int32(1) != 0 {
										v514 = v44
									} else {
										v514 = v46
									}
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
									if v510 != v515 {
										v787 = v28
										return v787
									} else {
										v518 = v514 + int32(2)
										v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
										if v521 == int32(3) {
											v524 = int32(128)
										} else {
											v524 = int32(32)
										}
										v529 = base.I32_div_s(v524, int32(8))
										v530 = F_memcmp(m, v93, v518, v529)
										mBase = m.M
										if v530 != 0 {
											v616 = v530
											v626 = v616
										} else {
											v531 = int32(0)
											v534 = v524 - v529<<(uint(int32(3))%32)
											if v534 <= v531 {
												v616 = v531
												v626 = v616
											} else {
												v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v529))))
												v539 = int32(128)
												v540 = v538 & v539
												v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v529))))
												if v540 != v542&v539 {
													v617 = v540
													if v617 != 0 {
														v620 = int32(1)
													} else {
														v620 = int32(-1)
													}
													v626 = v620
												} else {
													if v534 == int32(1) {
														v616 = v531
														v626 = v616
													} else {
														v548 = int32(1)
														v550 = int32(128)
														v551 = v538 << (uint(v548) % 32) & v550
														if v551 != v542<<(uint(v548)%32)&v550 {
															v617 = v551
															if v617 != 0 {
																v620 = int32(1)
															} else {
																v620 = int32(-1)
															}
															v626 = v620
														} else {
															if v534 < int32(3) {
																v616 = v531
																v626 = v616
															} else {
																v559 = int32(2)
																v561 = int32(128)
																v562 = v538 << (uint(v559) % 32) & v561
																if v562 != v542<<(uint(v559)%32)&v561 {
																	v617 = v562
																	if v617 != 0 {
																		v620 = int32(1)
																	} else {
																		v620 = int32(-1)
																	}
																	v626 = v620
																} else {
																	if v534 == int32(3) {
																		v616 = v531
																		v626 = v616
																	} else {
																		v570 = int32(3)
																		v572 = int32(128)
																		v573 = v538 << (uint(v570) % 32) & v572
																		if v573 != v542<<(uint(v570)%32)&v572 {
																			v617 = v573
																			if v617 != 0 {
																				v620 = int32(1)
																			} else {
																				v620 = int32(-1)
																			}
																			v626 = v620
																		} else {
																			if v534 < int32(5) {
																				v616 = v531
																				v626 = v616
																			} else {
																				v581 = int32(4)
																				v583 = int32(128)
																				v584 = v538 << (uint(v581) % 32) & v583
																				if v584 != v542<<(uint(v581)%32)&v583 {
																					v617 = v584
																					if v617 != 0 {
																						v620 = int32(1)
																					} else {
																						v620 = int32(-1)
																					}
																					v626 = v620
																				} else {
																					if v534 == int32(5) {
																						v616 = v531
																						v626 = v616
																					} else {
																						v592 = int32(5)
																						v594 = int32(128)
																						v595 = v538 << (uint(v592) % 32) & v594
																						if v595 != v542<<(uint(v592)%32)&v594 {
																							v617 = v595
																							if v617 != 0 {
																								v620 = int32(1)
																							} else {
																								v620 = int32(-1)
																							}
																							v626 = v620
																						} else {
																							if v534 < int32(7) {
																								v616 = v531
																								v626 = v616
																							} else {
																								v603 = int32(6)
																								v605 = int32(128)
																								v606 = v538 << (uint(v603) % 32) & v605
																								if v606 != v542<<(uint(v603)%32)&v605 {
																									v617 = v606
																									if v617 != 0 {
																										v620 = int32(1)
																									} else {
																										v620 = int32(-1)
																									}
																									v626 = v620
																								} else {
																									v616 = v531
																									v626 = v616
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
												}
											}
										}
										v795 = v626
										return base.B2i32(v795 != int32(0))
									}
								}
							}
						case 17, 18:
							v330 = int32(0)
							if v330 < v202 {
								v787 = v330
								return v787
							} else {
								v333 = int32(1)
								if v202 < int32(0) {
									v787 = v333
									return v787
								} else {
									v336 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+16)))
									v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v337)+12)))
									if v339&int32(1) != 0 {
										v483 = v333
										v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
										v486 = v484 & int32(1)
										if v486 != 0 {
											v487 = v44
										} else {
											v487 = v46
										}
										v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
										v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
										if v18&int32(65534) != int32(20) {
											v627 = int32(255)
											v628 = v489 & v627
											v630 = v488 & v627
											if base.Ui32(v630) < base.Ui32(v628) {
												v787 = v483
												return v787
											} else {
												if base.Ui32(v630) <= base.Ui32(v628) {
													v638 = v486
													if v638 != 0 {
														v642 = v44
													} else {
														v642 = v46
													}
													v644 = v642 + int32(2)
													v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v647 == int32(3) {
														v650 = int32(128)
													} else {
														v650 = int32(32)
													}
													v655 = base.I32_div_s(v650, int32(8))
													v656 = F_memcmp(m, v93, v644, v655)
													mBase = m.M
													if v656 != 0 {
														v742 = v656
														v752 = v742
													} else {
														v657 = int32(0)
														v660 = v650 - v655<<(uint(int32(3))%32)
														if v660 <= v657 {
															v742 = v657
															v752 = v742
														} else {
															v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
															v665 = int32(128)
															v666 = v664 & v665
															v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
															if v666 != v668&v665 {
																v743 = v666
																if v743 != 0 {
																	v746 = int32(1)
																} else {
																	v746 = int32(-1)
																}
																v752 = v746
															} else {
																if v660 == int32(1) {
																	v742 = v657
																	v752 = v742
																} else {
																	v674 = int32(1)
																	v676 = int32(128)
																	v677 = v664 << (uint(v674) % 32) & v676
																	if v677 != v668<<(uint(v674)%32)&v676 {
																		v743 = v677
																		if v743 != 0 {
																			v746 = int32(1)
																		} else {
																			v746 = int32(-1)
																		}
																		v752 = v746
																	} else {
																		if v660 < int32(3) {
																			v742 = v657
																			v752 = v742
																		} else {
																			v685 = int32(2)
																			v687 = int32(128)
																			v688 = v664 << (uint(v685) % 32) & v687
																			if v688 != v668<<(uint(v685)%32)&v687 {
																				v743 = v688
																				if v743 != 0 {
																					v746 = int32(1)
																				} else {
																					v746 = int32(-1)
																				}
																				v752 = v746
																			} else {
																				if v660 == int32(3) {
																					v742 = v657
																					v752 = v742
																				} else {
																					v696 = int32(3)
																					v698 = int32(128)
																					v699 = v664 << (uint(v696) % 32) & v698
																					if v699 != v668<<(uint(v696)%32)&v698 {
																						v743 = v699
																						if v743 != 0 {
																							v746 = int32(1)
																						} else {
																							v746 = int32(-1)
																						}
																						v752 = v746
																					} else {
																						if v660 < int32(5) {
																							v742 = v657
																							v752 = v742
																						} else {
																							v707 = int32(4)
																							v709 = int32(128)
																							v710 = v664 << (uint(v707) % 32) & v709
																							if v710 != v668<<(uint(v707)%32)&v709 {
																								v743 = v710
																								if v743 != 0 {
																									v746 = int32(1)
																								} else {
																									v746 = int32(-1)
																								}
																								v752 = v746
																							} else {
																								if v660 == int32(5) {
																									v742 = v657
																									v752 = v742
																								} else {
																									v718 = int32(5)
																									v720 = int32(128)
																									v721 = v664 << (uint(v718) % 32) & v720
																									if v721 != v668<<(uint(v718)%32)&v720 {
																										v743 = v721
																										if v743 != 0 {
																											v746 = int32(1)
																										} else {
																											v746 = int32(-1)
																										}
																										v752 = v746
																									} else {
																										if v660 < int32(7) {
																											v742 = v657
																											v752 = v742
																										} else {
																											v729 = int32(6)
																											v731 = int32(128)
																											v732 = v664 << (uint(v729) % 32) & v731
																											if v732 != v668<<(uint(v729)%32)&v731 {
																												v743 = v732
																												if v743 != 0 {
																													v746 = int32(1)
																												} else {
																													v746 = int32(-1)
																												}
																												v752 = v746
																											} else {
																												v742 = v657
																												v752 = v742
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
															}
														}
													}
													switch v18&int32(65535) - int32(18) {
													case 0:
														v802 = v752
														return base.B2i32(v802 == int32(0))
													case 1:
														v795 = v752
														return base.B2i32(v795 != int32(0))
													case 2:
														return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v752 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v752)
													case 5:
														return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v774 = m.ExcPending
														if v774 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(518516), int32(0))
															mBase = m.M
															v778 = m.ExcPending
															if v778 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494114), int32(327), int32(92161))
																mBase = m.M
																v783 = m.ExcPending
																if v783 != 0 {
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
													return int32(0)
												}
											}
										} else {
											v494 = int32(255)
											v495 = v489 & v494
											v497 = v488 & v494
											if base.Ui32(v495) < base.Ui32(v497) {
												v787 = v483
												return v787
											} else {
												if base.Ui32(v495) <= base.Ui32(v497) {
													v638 = v486
													if v638 != 0 {
														v642 = v44
													} else {
														v642 = v46
													}
													v644 = v642 + int32(2)
													v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v647 == int32(3) {
														v650 = int32(128)
													} else {
														v650 = int32(32)
													}
													v655 = base.I32_div_s(v650, int32(8))
													v656 = F_memcmp(m, v93, v644, v655)
													mBase = m.M
													if v656 != 0 {
														v742 = v656
														v752 = v742
													} else {
														v657 = int32(0)
														v660 = v650 - v655<<(uint(int32(3))%32)
														if v660 <= v657 {
															v742 = v657
															v752 = v742
														} else {
															v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
															v665 = int32(128)
															v666 = v664 & v665
															v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
															if v666 != v668&v665 {
																v743 = v666
																if v743 != 0 {
																	v746 = int32(1)
																} else {
																	v746 = int32(-1)
																}
																v752 = v746
															} else {
																if v660 == int32(1) {
																	v742 = v657
																	v752 = v742
																} else {
																	v674 = int32(1)
																	v676 = int32(128)
																	v677 = v664 << (uint(v674) % 32) & v676
																	if v677 != v668<<(uint(v674)%32)&v676 {
																		v743 = v677
																		if v743 != 0 {
																			v746 = int32(1)
																		} else {
																			v746 = int32(-1)
																		}
																		v752 = v746
																	} else {
																		if v660 < int32(3) {
																			v742 = v657
																			v752 = v742
																		} else {
																			v685 = int32(2)
																			v687 = int32(128)
																			v688 = v664 << (uint(v685) % 32) & v687
																			if v688 != v668<<(uint(v685)%32)&v687 {
																				v743 = v688
																				if v743 != 0 {
																					v746 = int32(1)
																				} else {
																					v746 = int32(-1)
																				}
																				v752 = v746
																			} else {
																				if v660 == int32(3) {
																					v742 = v657
																					v752 = v742
																				} else {
																					v696 = int32(3)
																					v698 = int32(128)
																					v699 = v664 << (uint(v696) % 32) & v698
																					if v699 != v668<<(uint(v696)%32)&v698 {
																						v743 = v699
																						if v743 != 0 {
																							v746 = int32(1)
																						} else {
																							v746 = int32(-1)
																						}
																						v752 = v746
																					} else {
																						if v660 < int32(5) {
																							v742 = v657
																							v752 = v742
																						} else {
																							v707 = int32(4)
																							v709 = int32(128)
																							v710 = v664 << (uint(v707) % 32) & v709
																							if v710 != v668<<(uint(v707)%32)&v709 {
																								v743 = v710
																								if v743 != 0 {
																									v746 = int32(1)
																								} else {
																									v746 = int32(-1)
																								}
																								v752 = v746
																							} else {
																								if v660 == int32(5) {
																									v742 = v657
																									v752 = v742
																								} else {
																									v718 = int32(5)
																									v720 = int32(128)
																									v721 = v664 << (uint(v718) % 32) & v720
																									if v721 != v668<<(uint(v718)%32)&v720 {
																										v743 = v721
																										if v743 != 0 {
																											v746 = int32(1)
																										} else {
																											v746 = int32(-1)
																										}
																										v752 = v746
																									} else {
																										if v660 < int32(7) {
																											v742 = v657
																											v752 = v742
																										} else {
																											v729 = int32(6)
																											v731 = int32(128)
																											v732 = v664 << (uint(v729) % 32) & v731
																											if v732 != v668<<(uint(v729)%32)&v731 {
																												v743 = v732
																												if v743 != 0 {
																													v746 = int32(1)
																												} else {
																													v746 = int32(-1)
																												}
																												v752 = v746
																											} else {
																												v742 = v657
																												v752 = v742
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
															}
														}
													}
													switch v18&int32(65535) - int32(18) {
													case 0:
														v802 = v752
														return base.B2i32(v802 == int32(0))
													case 1:
														v795 = v752
														return base.B2i32(v795 != int32(0))
													case 2:
														return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v752 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v752)
													case 5:
														return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v774 = m.ExcPending
														if v774 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(518516), int32(0))
															mBase = m.M
															v778 = m.ExcPending
															if v778 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494114), int32(327), int32(92161))
																mBase = m.M
																v783 = m.ExcPending
																if v783 != 0 {
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
													return int32(0)
												}
											}
										}
									} else {
										v787 = v333
										return v787
									}
								}
							}
						case 19, 20:
							v470 = int32(0)
							if v202 < v470 {
								v787 = v470
								return v787
							} else {
								v473 = int32(1)
								if v202 != 0 {
									v787 = v473
									return v787
								} else {
									v474 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+16)))
									v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474+v475)+12)))
									if v477&int32(1) == int32(0) {
										v787 = v473
										return v787
									} else {
										v483 = v473
										v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
										v486 = v484 & int32(1)
										if v486 != 0 {
											v487 = v44
										} else {
											v487 = v46
										}
										v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
										v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
										if v18&int32(65534) != int32(20) {
											v627 = int32(255)
											v628 = v489 & v627
											v630 = v488 & v627
											if base.Ui32(v630) < base.Ui32(v628) {
												v787 = v483
												return v787
											} else {
												if base.Ui32(v630) <= base.Ui32(v628) {
													v638 = v486
													if v638 != 0 {
														v642 = v44
													} else {
														v642 = v46
													}
													v644 = v642 + int32(2)
													v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v647 == int32(3) {
														v650 = int32(128)
													} else {
														v650 = int32(32)
													}
													v655 = base.I32_div_s(v650, int32(8))
													v656 = F_memcmp(m, v93, v644, v655)
													mBase = m.M
													if v656 != 0 {
														v742 = v656
														v752 = v742
													} else {
														v657 = int32(0)
														v660 = v650 - v655<<(uint(int32(3))%32)
														if v660 <= v657 {
															v742 = v657
															v752 = v742
														} else {
															v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
															v665 = int32(128)
															v666 = v664 & v665
															v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
															if v666 != v668&v665 {
																v743 = v666
																if v743 != 0 {
																	v746 = int32(1)
																} else {
																	v746 = int32(-1)
																}
																v752 = v746
															} else {
																if v660 == int32(1) {
																	v742 = v657
																	v752 = v742
																} else {
																	v674 = int32(1)
																	v676 = int32(128)
																	v677 = v664 << (uint(v674) % 32) & v676
																	if v677 != v668<<(uint(v674)%32)&v676 {
																		v743 = v677
																		if v743 != 0 {
																			v746 = int32(1)
																		} else {
																			v746 = int32(-1)
																		}
																		v752 = v746
																	} else {
																		if v660 < int32(3) {
																			v742 = v657
																			v752 = v742
																		} else {
																			v685 = int32(2)
																			v687 = int32(128)
																			v688 = v664 << (uint(v685) % 32) & v687
																			if v688 != v668<<(uint(v685)%32)&v687 {
																				v743 = v688
																				if v743 != 0 {
																					v746 = int32(1)
																				} else {
																					v746 = int32(-1)
																				}
																				v752 = v746
																			} else {
																				if v660 == int32(3) {
																					v742 = v657
																					v752 = v742
																				} else {
																					v696 = int32(3)
																					v698 = int32(128)
																					v699 = v664 << (uint(v696) % 32) & v698
																					if v699 != v668<<(uint(v696)%32)&v698 {
																						v743 = v699
																						if v743 != 0 {
																							v746 = int32(1)
																						} else {
																							v746 = int32(-1)
																						}
																						v752 = v746
																					} else {
																						if v660 < int32(5) {
																							v742 = v657
																							v752 = v742
																						} else {
																							v707 = int32(4)
																							v709 = int32(128)
																							v710 = v664 << (uint(v707) % 32) & v709
																							if v710 != v668<<(uint(v707)%32)&v709 {
																								v743 = v710
																								if v743 != 0 {
																									v746 = int32(1)
																								} else {
																									v746 = int32(-1)
																								}
																								v752 = v746
																							} else {
																								if v660 == int32(5) {
																									v742 = v657
																									v752 = v742
																								} else {
																									v718 = int32(5)
																									v720 = int32(128)
																									v721 = v664 << (uint(v718) % 32) & v720
																									if v721 != v668<<(uint(v718)%32)&v720 {
																										v743 = v721
																										if v743 != 0 {
																											v746 = int32(1)
																										} else {
																											v746 = int32(-1)
																										}
																										v752 = v746
																									} else {
																										if v660 < int32(7) {
																											v742 = v657
																											v752 = v742
																										} else {
																											v729 = int32(6)
																											v731 = int32(128)
																											v732 = v664 << (uint(v729) % 32) & v731
																											if v732 != v668<<(uint(v729)%32)&v731 {
																												v743 = v732
																												if v743 != 0 {
																													v746 = int32(1)
																												} else {
																													v746 = int32(-1)
																												}
																												v752 = v746
																											} else {
																												v742 = v657
																												v752 = v742
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
															}
														}
													}
													switch v18&int32(65535) - int32(18) {
													case 0:
														v802 = v752
														return base.B2i32(v802 == int32(0))
													case 1:
														v795 = v752
														return base.B2i32(v795 != int32(0))
													case 2:
														return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v752 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v752)
													case 5:
														return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v774 = m.ExcPending
														if v774 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(518516), int32(0))
															mBase = m.M
															v778 = m.ExcPending
															if v778 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494114), int32(327), int32(92161))
																mBase = m.M
																v783 = m.ExcPending
																if v783 != 0 {
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
													return int32(0)
												}
											}
										} else {
											v494 = int32(255)
											v495 = v489 & v494
											v497 = v488 & v494
											if base.Ui32(v495) < base.Ui32(v497) {
												v787 = v483
												return v787
											} else {
												if base.Ui32(v495) <= base.Ui32(v497) {
													v638 = v486
													if v638 != 0 {
														v642 = v44
													} else {
														v642 = v46
													}
													v644 = v642 + int32(2)
													v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
													if v647 == int32(3) {
														v650 = int32(128)
													} else {
														v650 = int32(32)
													}
													v655 = base.I32_div_s(v650, int32(8))
													v656 = F_memcmp(m, v93, v644, v655)
													mBase = m.M
													if v656 != 0 {
														v742 = v656
														v752 = v742
													} else {
														v657 = int32(0)
														v660 = v650 - v655<<(uint(int32(3))%32)
														if v660 <= v657 {
															v742 = v657
															v752 = v742
														} else {
															v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
															v665 = int32(128)
															v666 = v664 & v665
															v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
															if v666 != v668&v665 {
																v743 = v666
																if v743 != 0 {
																	v746 = int32(1)
																} else {
																	v746 = int32(-1)
																}
																v752 = v746
															} else {
																if v660 == int32(1) {
																	v742 = v657
																	v752 = v742
																} else {
																	v674 = int32(1)
																	v676 = int32(128)
																	v677 = v664 << (uint(v674) % 32) & v676
																	if v677 != v668<<(uint(v674)%32)&v676 {
																		v743 = v677
																		if v743 != 0 {
																			v746 = int32(1)
																		} else {
																			v746 = int32(-1)
																		}
																		v752 = v746
																	} else {
																		if v660 < int32(3) {
																			v742 = v657
																			v752 = v742
																		} else {
																			v685 = int32(2)
																			v687 = int32(128)
																			v688 = v664 << (uint(v685) % 32) & v687
																			if v688 != v668<<(uint(v685)%32)&v687 {
																				v743 = v688
																				if v743 != 0 {
																					v746 = int32(1)
																				} else {
																					v746 = int32(-1)
																				}
																				v752 = v746
																			} else {
																				if v660 == int32(3) {
																					v742 = v657
																					v752 = v742
																				} else {
																					v696 = int32(3)
																					v698 = int32(128)
																					v699 = v664 << (uint(v696) % 32) & v698
																					if v699 != v668<<(uint(v696)%32)&v698 {
																						v743 = v699
																						if v743 != 0 {
																							v746 = int32(1)
																						} else {
																							v746 = int32(-1)
																						}
																						v752 = v746
																					} else {
																						if v660 < int32(5) {
																							v742 = v657
																							v752 = v742
																						} else {
																							v707 = int32(4)
																							v709 = int32(128)
																							v710 = v664 << (uint(v707) % 32) & v709
																							if v710 != v668<<(uint(v707)%32)&v709 {
																								v743 = v710
																								if v743 != 0 {
																									v746 = int32(1)
																								} else {
																									v746 = int32(-1)
																								}
																								v752 = v746
																							} else {
																								if v660 == int32(5) {
																									v742 = v657
																									v752 = v742
																								} else {
																									v718 = int32(5)
																									v720 = int32(128)
																									v721 = v664 << (uint(v718) % 32) & v720
																									if v721 != v668<<(uint(v718)%32)&v720 {
																										v743 = v721
																										if v743 != 0 {
																											v746 = int32(1)
																										} else {
																											v746 = int32(-1)
																										}
																										v752 = v746
																									} else {
																										if v660 < int32(7) {
																											v742 = v657
																											v752 = v742
																										} else {
																											v729 = int32(6)
																											v731 = int32(128)
																											v732 = v664 << (uint(v729) % 32) & v731
																											if v732 != v668<<(uint(v729)%32)&v731 {
																												v743 = v732
																												if v743 != 0 {
																													v746 = int32(1)
																												} else {
																													v746 = int32(-1)
																												}
																												v752 = v746
																											} else {
																												v742 = v657
																												v752 = v742
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
															}
														}
													}
													switch v18&int32(65535) - int32(18) {
													case 0:
														v802 = v752
														return base.B2i32(v802 == int32(0))
													case 1:
														v795 = v752
														return base.B2i32(v795 != int32(0))
													case 2:
														return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
													case 3:
														return base.B2i32(v752 <= int32(0))
													case 4:
														return base.B2i32(int32(0) < v752)
													case 5:
														return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v774 = m.ExcPending
														if v774 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(518516), int32(0))
															mBase = m.M
															v778 = m.ExcPending
															if v778 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494114), int32(327), int32(92161))
																mBase = m.M
																v783 = m.ExcPending
																if v783 != 0 {
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
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						return int32(0)
					}
				default:
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v90 = v89
					v91 = v88
					v93 = v19 + int32(4)
					v95 = v47 + int32(2)
					v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
					if base.Ui32(v96) < base.Ui32(v90) {
						v98 = v96
					} else {
						v98 = v90
					}
					if base.Ui32(v98) < base.Ui32(v91) {
						v100 = v98
					} else {
						v100 = v91
					}
					v105 = base.I32_div_s(v100, int32(8))
					v106 = F_memcmp(m, v93, v95, v105)
					mBase = m.M
					if v106 != 0 {
						v192 = v106
						v202 = v192
					} else {
						v107 = int32(0)
						v110 = v100 - v105<<(uint(int32(3))%32)
						if v110 <= v107 {
							v192 = v107
							v202 = v192
						} else {
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v105))))
							v115 = int32(128)
							v116 = v114 & v115
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v105))))
							if v116 != v118&v115 {
								v193 = v116
								if v193 != 0 {
									v196 = int32(1)
								} else {
									v196 = int32(-1)
								}
								v202 = v196
							} else {
								if v110 == int32(1) {
									v192 = v107
									v202 = v192
								} else {
									v124 = int32(1)
									v126 = int32(128)
									v127 = v114 << (uint(v124) % 32) & v126
									if v127 != v118<<(uint(v124)%32)&v126 {
										v193 = v127
										if v193 != 0 {
											v196 = int32(1)
										} else {
											v196 = int32(-1)
										}
										v202 = v196
									} else {
										if v110 < int32(3) {
											v192 = v107
											v202 = v192
										} else {
											v135 = int32(2)
											v137 = int32(128)
											v138 = v114 << (uint(v135) % 32) & v137
											if v138 != v118<<(uint(v135)%32)&v137 {
												v193 = v138
												if v193 != 0 {
													v196 = int32(1)
												} else {
													v196 = int32(-1)
												}
												v202 = v196
											} else {
												if v110 == int32(3) {
													v192 = v107
													v202 = v192
												} else {
													v146 = int32(3)
													v148 = int32(128)
													v149 = v114 << (uint(v146) % 32) & v148
													if v149 != v118<<(uint(v146)%32)&v148 {
														v193 = v149
														if v193 != 0 {
															v196 = int32(1)
														} else {
															v196 = int32(-1)
														}
														v202 = v196
													} else {
														if v110 < int32(5) {
															v192 = v107
															v202 = v192
														} else {
															v157 = int32(4)
															v159 = int32(128)
															v160 = v114 << (uint(v157) % 32) & v159
															if v160 != v118<<(uint(v157)%32)&v159 {
																v193 = v160
																if v193 != 0 {
																	v196 = int32(1)
																} else {
																	v196 = int32(-1)
																}
																v202 = v196
															} else {
																if v110 == int32(5) {
																	v192 = v107
																	v202 = v192
																} else {
																	v168 = int32(5)
																	v170 = int32(128)
																	v171 = v114 << (uint(v168) % 32) & v170
																	if v171 != v118<<(uint(v168)%32)&v170 {
																		v193 = v171
																		if v193 != 0 {
																			v196 = int32(1)
																		} else {
																			v196 = int32(-1)
																		}
																		v202 = v196
																	} else {
																		if v110 < int32(7) {
																			v192 = v107
																			v202 = v192
																		} else {
																			v179 = int32(6)
																			v181 = int32(128)
																			v182 = v114 << (uint(v179) % 32) & v181
																			if v182 != v118<<(uint(v179)%32)&v181 {
																				v193 = v182
																				if v193 != 0 {
																					v196 = int32(1)
																				} else {
																					v196 = int32(-1)
																				}
																				v202 = v196
																			} else {
																				v192 = v107
																				v202 = v192
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
							}
						}
					}
					switch v18&int32(65535) - int32(3) {
					case 0, 21, 22, 23, 24:
						v321 = v202
						return base.B2i32(v321 == int32(0))
					default:
						v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
						v638 = v635 & int32(1)
						if v638 != 0 {
							v642 = v44
						} else {
							v642 = v46
						}
						v644 = v642 + int32(2)
						v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
						if v647 == int32(3) {
							v650 = int32(128)
						} else {
							v650 = int32(32)
						}
						v655 = base.I32_div_s(v650, int32(8))
						v656 = F_memcmp(m, v93, v644, v655)
						mBase = m.M
						if v656 != 0 {
							v742 = v656
							v752 = v742
						} else {
							v657 = int32(0)
							v660 = v650 - v655<<(uint(int32(3))%32)
							if v660 <= v657 {
								v742 = v657
								v752 = v742
							} else {
								v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
								v665 = int32(128)
								v666 = v664 & v665
								v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
								if v666 != v668&v665 {
									v743 = v666
									if v743 != 0 {
										v746 = int32(1)
									} else {
										v746 = int32(-1)
									}
									v752 = v746
								} else {
									if v660 == int32(1) {
										v742 = v657
										v752 = v742
									} else {
										v674 = int32(1)
										v676 = int32(128)
										v677 = v664 << (uint(v674) % 32) & v676
										if v677 != v668<<(uint(v674)%32)&v676 {
											v743 = v677
											if v743 != 0 {
												v746 = int32(1)
											} else {
												v746 = int32(-1)
											}
											v752 = v746
										} else {
											if v660 < int32(3) {
												v742 = v657
												v752 = v742
											} else {
												v685 = int32(2)
												v687 = int32(128)
												v688 = v664 << (uint(v685) % 32) & v687
												if v688 != v668<<(uint(v685)%32)&v687 {
													v743 = v688
													if v743 != 0 {
														v746 = int32(1)
													} else {
														v746 = int32(-1)
													}
													v752 = v746
												} else {
													if v660 == int32(3) {
														v742 = v657
														v752 = v742
													} else {
														v696 = int32(3)
														v698 = int32(128)
														v699 = v664 << (uint(v696) % 32) & v698
														if v699 != v668<<(uint(v696)%32)&v698 {
															v743 = v699
															if v743 != 0 {
																v746 = int32(1)
															} else {
																v746 = int32(-1)
															}
															v752 = v746
														} else {
															if v660 < int32(5) {
																v742 = v657
																v752 = v742
															} else {
																v707 = int32(4)
																v709 = int32(128)
																v710 = v664 << (uint(v707) % 32) & v709
																if v710 != v668<<(uint(v707)%32)&v709 {
																	v743 = v710
																	if v743 != 0 {
																		v746 = int32(1)
																	} else {
																		v746 = int32(-1)
																	}
																	v752 = v746
																} else {
																	if v660 == int32(5) {
																		v742 = v657
																		v752 = v742
																	} else {
																		v718 = int32(5)
																		v720 = int32(128)
																		v721 = v664 << (uint(v718) % 32) & v720
																		if v721 != v668<<(uint(v718)%32)&v720 {
																			v743 = v721
																			if v743 != 0 {
																				v746 = int32(1)
																			} else {
																				v746 = int32(-1)
																			}
																			v752 = v746
																		} else {
																			if v660 < int32(7) {
																				v742 = v657
																				v752 = v742
																			} else {
																				v729 = int32(6)
																				v731 = int32(128)
																				v732 = v664 << (uint(v729) % 32) & v731
																				if v732 != v668<<(uint(v729)%32)&v731 {
																					v743 = v732
																					if v743 != 0 {
																						v746 = int32(1)
																					} else {
																						v746 = int32(-1)
																					}
																					v752 = v746
																				} else {
																					v742 = v657
																					v752 = v742
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
								}
							}
						}
						switch v18&int32(65535) - int32(18) {
						case 0:
							v802 = v752
							return base.B2i32(v802 == int32(0))
						case 1:
							v795 = v752
							return base.B2i32(v795 != int32(0))
						case 2:
							return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
						case 3:
							return base.B2i32(v752 <= int32(0))
						case 4:
							return base.B2i32(int32(0) < v752)
						case 5:
							return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v774 = m.ExcPending
							if v774 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(518516), int32(0))
								mBase = m.M
								v778 = m.ExcPending
								if v778 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494114), int32(327), int32(92161))
									mBase = m.M
									v783 = m.ExcPending
									if v783 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					case 15:
						if v202 != 0 {
							v787 = int32(0)
							return v787
						} else {
							v343 = int32(1)
							v344 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+16)))
							v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v345)+12)))
							if v347&v343 == int32(0) {
								v787 = v343
								return v787
							} else {
								v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
								v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
								if v354&int32(1) != 0 {
									v357 = v44
								} else {
									v357 = v46
								}
								v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
								if v353 != v358 {
									v787 = int32(0)
									return v787
								} else {
									v361 = v357 + int32(2)
									v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
									if v364 == int32(3) {
										v367 = int32(128)
									} else {
										v367 = int32(32)
									}
									v372 = base.I32_div_s(v367, int32(8))
									v373 = F_memcmp(m, v93, v361, v372)
									mBase = m.M
									if v373 != 0 {
										v459 = v373
										v469 = v459
									} else {
										v374 = int32(0)
										v377 = v367 - v372<<(uint(int32(3))%32)
										if v377 <= v374 {
											v459 = v374
											v469 = v459
										} else {
											v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v372))))
											v382 = int32(128)
											v383 = v381 & v382
											v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v372))))
											if v383 != v385&v382 {
												v460 = v383
												if v460 != 0 {
													v463 = int32(1)
												} else {
													v463 = int32(-1)
												}
												v469 = v463
											} else {
												if v377 == int32(1) {
													v459 = v374
													v469 = v459
												} else {
													v391 = int32(1)
													v393 = int32(128)
													v394 = v381 << (uint(v391) % 32) & v393
													if v394 != v385<<(uint(v391)%32)&v393 {
														v460 = v394
														if v460 != 0 {
															v463 = int32(1)
														} else {
															v463 = int32(-1)
														}
														v469 = v463
													} else {
														if v377 < int32(3) {
															v459 = v374
															v469 = v459
														} else {
															v402 = int32(2)
															v404 = int32(128)
															v405 = v381 << (uint(v402) % 32) & v404
															if v405 != v385<<(uint(v402)%32)&v404 {
																v460 = v405
																if v460 != 0 {
																	v463 = int32(1)
																} else {
																	v463 = int32(-1)
																}
																v469 = v463
															} else {
																if v377 == int32(3) {
																	v459 = v374
																	v469 = v459
																} else {
																	v413 = int32(3)
																	v415 = int32(128)
																	v416 = v381 << (uint(v413) % 32) & v415
																	if v416 != v385<<(uint(v413)%32)&v415 {
																		v460 = v416
																		if v460 != 0 {
																			v463 = int32(1)
																		} else {
																			v463 = int32(-1)
																		}
																		v469 = v463
																	} else {
																		if v377 < int32(5) {
																			v459 = v374
																			v469 = v459
																		} else {
																			v424 = int32(4)
																			v426 = int32(128)
																			v427 = v381 << (uint(v424) % 32) & v426
																			if v427 != v385<<(uint(v424)%32)&v426 {
																				v460 = v427
																				if v460 != 0 {
																					v463 = int32(1)
																				} else {
																					v463 = int32(-1)
																				}
																				v469 = v463
																			} else {
																				if v377 == int32(5) {
																					v459 = v374
																					v469 = v459
																				} else {
																					v435 = int32(5)
																					v437 = int32(128)
																					v438 = v381 << (uint(v435) % 32) & v437
																					if v438 != v385<<(uint(v435)%32)&v437 {
																						v460 = v438
																						if v460 != 0 {
																							v463 = int32(1)
																						} else {
																							v463 = int32(-1)
																						}
																						v469 = v463
																					} else {
																						if v377 < int32(7) {
																							v459 = v374
																							v469 = v459
																						} else {
																							v446 = int32(6)
																							v448 = int32(128)
																							v449 = v381 << (uint(v446) % 32) & v448
																							if v449 != v385<<(uint(v446)%32)&v448 {
																								v460 = v449
																								if v460 != 0 {
																									v463 = int32(1)
																								} else {
																									v463 = int32(-1)
																								}
																								v469 = v463
																							} else {
																								v459 = v374
																								v469 = v459
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
											}
										}
									}
									v802 = v469
									return base.B2i32(v802 == int32(0))
								}
							}
						}
					case 16:
						if v202 != 0 {
							v787 = v28
							return v787
						} else {
							v502 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v502)+16)))
							v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v503)+12)))
							if v505&int32(1) == int32(0) {
								v787 = v28
								return v787
							} else {
								v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
								v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
								if v511&int32(1) != 0 {
									v514 = v44
								} else {
									v514 = v46
								}
								v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
								if v510 != v515 {
									v787 = v28
									return v787
								} else {
									v518 = v514 + int32(2)
									v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
									if v521 == int32(3) {
										v524 = int32(128)
									} else {
										v524 = int32(32)
									}
									v529 = base.I32_div_s(v524, int32(8))
									v530 = F_memcmp(m, v93, v518, v529)
									mBase = m.M
									if v530 != 0 {
										v616 = v530
										v626 = v616
									} else {
										v531 = int32(0)
										v534 = v524 - v529<<(uint(int32(3))%32)
										if v534 <= v531 {
											v616 = v531
											v626 = v616
										} else {
											v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v529))))
											v539 = int32(128)
											v540 = v538 & v539
											v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v529))))
											if v540 != v542&v539 {
												v617 = v540
												if v617 != 0 {
													v620 = int32(1)
												} else {
													v620 = int32(-1)
												}
												v626 = v620
											} else {
												if v534 == int32(1) {
													v616 = v531
													v626 = v616
												} else {
													v548 = int32(1)
													v550 = int32(128)
													v551 = v538 << (uint(v548) % 32) & v550
													if v551 != v542<<(uint(v548)%32)&v550 {
														v617 = v551
														if v617 != 0 {
															v620 = int32(1)
														} else {
															v620 = int32(-1)
														}
														v626 = v620
													} else {
														if v534 < int32(3) {
															v616 = v531
															v626 = v616
														} else {
															v559 = int32(2)
															v561 = int32(128)
															v562 = v538 << (uint(v559) % 32) & v561
															if v562 != v542<<(uint(v559)%32)&v561 {
																v617 = v562
																if v617 != 0 {
																	v620 = int32(1)
																} else {
																	v620 = int32(-1)
																}
																v626 = v620
															} else {
																if v534 == int32(3) {
																	v616 = v531
																	v626 = v616
																} else {
																	v570 = int32(3)
																	v572 = int32(128)
																	v573 = v538 << (uint(v570) % 32) & v572
																	if v573 != v542<<(uint(v570)%32)&v572 {
																		v617 = v573
																		if v617 != 0 {
																			v620 = int32(1)
																		} else {
																			v620 = int32(-1)
																		}
																		v626 = v620
																	} else {
																		if v534 < int32(5) {
																			v616 = v531
																			v626 = v616
																		} else {
																			v581 = int32(4)
																			v583 = int32(128)
																			v584 = v538 << (uint(v581) % 32) & v583
																			if v584 != v542<<(uint(v581)%32)&v583 {
																				v617 = v584
																				if v617 != 0 {
																					v620 = int32(1)
																				} else {
																					v620 = int32(-1)
																				}
																				v626 = v620
																			} else {
																				if v534 == int32(5) {
																					v616 = v531
																					v626 = v616
																				} else {
																					v592 = int32(5)
																					v594 = int32(128)
																					v595 = v538 << (uint(v592) % 32) & v594
																					if v595 != v542<<(uint(v592)%32)&v594 {
																						v617 = v595
																						if v617 != 0 {
																							v620 = int32(1)
																						} else {
																							v620 = int32(-1)
																						}
																						v626 = v620
																					} else {
																						if v534 < int32(7) {
																							v616 = v531
																							v626 = v616
																						} else {
																							v603 = int32(6)
																							v605 = int32(128)
																							v606 = v538 << (uint(v603) % 32) & v605
																							if v606 != v542<<(uint(v603)%32)&v605 {
																								v617 = v606
																								if v617 != 0 {
																									v620 = int32(1)
																								} else {
																									v620 = int32(-1)
																								}
																								v626 = v620
																							} else {
																								v616 = v531
																								v626 = v616
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
											}
										}
									}
									v795 = v626
									return base.B2i32(v795 != int32(0))
								}
							}
						}
					case 17, 18:
						v330 = int32(0)
						if v330 < v202 {
							v787 = v330
							return v787
						} else {
							v333 = int32(1)
							if v202 < int32(0) {
								v787 = v333
								return v787
							} else {
								v336 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+16)))
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v337)+12)))
								if v339&int32(1) != 0 {
									v483 = v333
									v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									v486 = v484 & int32(1)
									if v486 != 0 {
										v487 = v44
									} else {
										v487 = v46
									}
									v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
									v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									if v18&int32(65534) != int32(20) {
										v627 = int32(255)
										v628 = v489 & v627
										v630 = v488 & v627
										if base.Ui32(v630) < base.Ui32(v628) {
											v787 = v483
											return v787
										} else {
											if base.Ui32(v630) <= base.Ui32(v628) {
												v638 = v486
												if v638 != 0 {
													v642 = v44
												} else {
													v642 = v46
												}
												v644 = v642 + int32(2)
												v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v647 == int32(3) {
													v650 = int32(128)
												} else {
													v650 = int32(32)
												}
												v655 = base.I32_div_s(v650, int32(8))
												v656 = F_memcmp(m, v93, v644, v655)
												mBase = m.M
												if v656 != 0 {
													v742 = v656
													v752 = v742
												} else {
													v657 = int32(0)
													v660 = v650 - v655<<(uint(int32(3))%32)
													if v660 <= v657 {
														v742 = v657
														v752 = v742
													} else {
														v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
														v665 = int32(128)
														v666 = v664 & v665
														v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
														if v666 != v668&v665 {
															v743 = v666
															if v743 != 0 {
																v746 = int32(1)
															} else {
																v746 = int32(-1)
															}
															v752 = v746
														} else {
															if v660 == int32(1) {
																v742 = v657
																v752 = v742
															} else {
																v674 = int32(1)
																v676 = int32(128)
																v677 = v664 << (uint(v674) % 32) & v676
																if v677 != v668<<(uint(v674)%32)&v676 {
																	v743 = v677
																	if v743 != 0 {
																		v746 = int32(1)
																	} else {
																		v746 = int32(-1)
																	}
																	v752 = v746
																} else {
																	if v660 < int32(3) {
																		v742 = v657
																		v752 = v742
																	} else {
																		v685 = int32(2)
																		v687 = int32(128)
																		v688 = v664 << (uint(v685) % 32) & v687
																		if v688 != v668<<(uint(v685)%32)&v687 {
																			v743 = v688
																			if v743 != 0 {
																				v746 = int32(1)
																			} else {
																				v746 = int32(-1)
																			}
																			v752 = v746
																		} else {
																			if v660 == int32(3) {
																				v742 = v657
																				v752 = v742
																			} else {
																				v696 = int32(3)
																				v698 = int32(128)
																				v699 = v664 << (uint(v696) % 32) & v698
																				if v699 != v668<<(uint(v696)%32)&v698 {
																					v743 = v699
																					if v743 != 0 {
																						v746 = int32(1)
																					} else {
																						v746 = int32(-1)
																					}
																					v752 = v746
																				} else {
																					if v660 < int32(5) {
																						v742 = v657
																						v752 = v742
																					} else {
																						v707 = int32(4)
																						v709 = int32(128)
																						v710 = v664 << (uint(v707) % 32) & v709
																						if v710 != v668<<(uint(v707)%32)&v709 {
																							v743 = v710
																							if v743 != 0 {
																								v746 = int32(1)
																							} else {
																								v746 = int32(-1)
																							}
																							v752 = v746
																						} else {
																							if v660 == int32(5) {
																								v742 = v657
																								v752 = v742
																							} else {
																								v718 = int32(5)
																								v720 = int32(128)
																								v721 = v664 << (uint(v718) % 32) & v720
																								if v721 != v668<<(uint(v718)%32)&v720 {
																									v743 = v721
																									if v743 != 0 {
																										v746 = int32(1)
																									} else {
																										v746 = int32(-1)
																									}
																									v752 = v746
																								} else {
																									if v660 < int32(7) {
																										v742 = v657
																										v752 = v742
																									} else {
																										v729 = int32(6)
																										v731 = int32(128)
																										v732 = v664 << (uint(v729) % 32) & v731
																										if v732 != v668<<(uint(v729)%32)&v731 {
																											v743 = v732
																											if v743 != 0 {
																												v746 = int32(1)
																											} else {
																												v746 = int32(-1)
																											}
																											v752 = v746
																										} else {
																											v742 = v657
																											v752 = v742
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
														}
													}
												}
												switch v18&int32(65535) - int32(18) {
												case 0:
													v802 = v752
													return base.B2i32(v802 == int32(0))
												case 1:
													v795 = v752
													return base.B2i32(v795 != int32(0))
												case 2:
													return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v752 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v752)
												case 5:
													return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v774 = m.ExcPending
													if v774 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(518516), int32(0))
														mBase = m.M
														v778 = m.ExcPending
														if v778 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494114), int32(327), int32(92161))
															mBase = m.M
															v783 = m.ExcPending
															if v783 != 0 {
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
												return int32(0)
											}
										}
									} else {
										v494 = int32(255)
										v495 = v489 & v494
										v497 = v488 & v494
										if base.Ui32(v495) < base.Ui32(v497) {
											v787 = v483
											return v787
										} else {
											if base.Ui32(v495) <= base.Ui32(v497) {
												v638 = v486
												if v638 != 0 {
													v642 = v44
												} else {
													v642 = v46
												}
												v644 = v642 + int32(2)
												v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v647 == int32(3) {
													v650 = int32(128)
												} else {
													v650 = int32(32)
												}
												v655 = base.I32_div_s(v650, int32(8))
												v656 = F_memcmp(m, v93, v644, v655)
												mBase = m.M
												if v656 != 0 {
													v742 = v656
													v752 = v742
												} else {
													v657 = int32(0)
													v660 = v650 - v655<<(uint(int32(3))%32)
													if v660 <= v657 {
														v742 = v657
														v752 = v742
													} else {
														v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
														v665 = int32(128)
														v666 = v664 & v665
														v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
														if v666 != v668&v665 {
															v743 = v666
															if v743 != 0 {
																v746 = int32(1)
															} else {
																v746 = int32(-1)
															}
															v752 = v746
														} else {
															if v660 == int32(1) {
																v742 = v657
																v752 = v742
															} else {
																v674 = int32(1)
																v676 = int32(128)
																v677 = v664 << (uint(v674) % 32) & v676
																if v677 != v668<<(uint(v674)%32)&v676 {
																	v743 = v677
																	if v743 != 0 {
																		v746 = int32(1)
																	} else {
																		v746 = int32(-1)
																	}
																	v752 = v746
																} else {
																	if v660 < int32(3) {
																		v742 = v657
																		v752 = v742
																	} else {
																		v685 = int32(2)
																		v687 = int32(128)
																		v688 = v664 << (uint(v685) % 32) & v687
																		if v688 != v668<<(uint(v685)%32)&v687 {
																			v743 = v688
																			if v743 != 0 {
																				v746 = int32(1)
																			} else {
																				v746 = int32(-1)
																			}
																			v752 = v746
																		} else {
																			if v660 == int32(3) {
																				v742 = v657
																				v752 = v742
																			} else {
																				v696 = int32(3)
																				v698 = int32(128)
																				v699 = v664 << (uint(v696) % 32) & v698
																				if v699 != v668<<(uint(v696)%32)&v698 {
																					v743 = v699
																					if v743 != 0 {
																						v746 = int32(1)
																					} else {
																						v746 = int32(-1)
																					}
																					v752 = v746
																				} else {
																					if v660 < int32(5) {
																						v742 = v657
																						v752 = v742
																					} else {
																						v707 = int32(4)
																						v709 = int32(128)
																						v710 = v664 << (uint(v707) % 32) & v709
																						if v710 != v668<<(uint(v707)%32)&v709 {
																							v743 = v710
																							if v743 != 0 {
																								v746 = int32(1)
																							} else {
																								v746 = int32(-1)
																							}
																							v752 = v746
																						} else {
																							if v660 == int32(5) {
																								v742 = v657
																								v752 = v742
																							} else {
																								v718 = int32(5)
																								v720 = int32(128)
																								v721 = v664 << (uint(v718) % 32) & v720
																								if v721 != v668<<(uint(v718)%32)&v720 {
																									v743 = v721
																									if v743 != 0 {
																										v746 = int32(1)
																									} else {
																										v746 = int32(-1)
																									}
																									v752 = v746
																								} else {
																									if v660 < int32(7) {
																										v742 = v657
																										v752 = v742
																									} else {
																										v729 = int32(6)
																										v731 = int32(128)
																										v732 = v664 << (uint(v729) % 32) & v731
																										if v732 != v668<<(uint(v729)%32)&v731 {
																											v743 = v732
																											if v743 != 0 {
																												v746 = int32(1)
																											} else {
																												v746 = int32(-1)
																											}
																											v752 = v746
																										} else {
																											v742 = v657
																											v752 = v742
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
														}
													}
												}
												switch v18&int32(65535) - int32(18) {
												case 0:
													v802 = v752
													return base.B2i32(v802 == int32(0))
												case 1:
													v795 = v752
													return base.B2i32(v795 != int32(0))
												case 2:
													return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v752 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v752)
												case 5:
													return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v774 = m.ExcPending
													if v774 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(518516), int32(0))
														mBase = m.M
														v778 = m.ExcPending
														if v778 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494114), int32(327), int32(92161))
															mBase = m.M
															v783 = m.ExcPending
															if v783 != 0 {
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
												return int32(0)
											}
										}
									}
								} else {
									v787 = v333
									return v787
								}
							}
						}
					case 19, 20:
						v470 = int32(0)
						if v202 < v470 {
							v787 = v470
							return v787
						} else {
							v473 = int32(1)
							if v202 != 0 {
								v787 = v473
								return v787
							} else {
								v474 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+16)))
								v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474+v475)+12)))
								if v477&int32(1) == int32(0) {
									v787 = v473
									return v787
								} else {
									v483 = v473
									v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
									v486 = v484 & int32(1)
									if v486 != 0 {
										v487 = v44
									} else {
										v487 = v46
									}
									v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
									v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
									if v18&int32(65534) != int32(20) {
										v627 = int32(255)
										v628 = v489 & v627
										v630 = v488 & v627
										if base.Ui32(v630) < base.Ui32(v628) {
											v787 = v483
											return v787
										} else {
											if base.Ui32(v630) <= base.Ui32(v628) {
												v638 = v486
												if v638 != 0 {
													v642 = v44
												} else {
													v642 = v46
												}
												v644 = v642 + int32(2)
												v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v647 == int32(3) {
													v650 = int32(128)
												} else {
													v650 = int32(32)
												}
												v655 = base.I32_div_s(v650, int32(8))
												v656 = F_memcmp(m, v93, v644, v655)
												mBase = m.M
												if v656 != 0 {
													v742 = v656
													v752 = v742
												} else {
													v657 = int32(0)
													v660 = v650 - v655<<(uint(int32(3))%32)
													if v660 <= v657 {
														v742 = v657
														v752 = v742
													} else {
														v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
														v665 = int32(128)
														v666 = v664 & v665
														v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
														if v666 != v668&v665 {
															v743 = v666
															if v743 != 0 {
																v746 = int32(1)
															} else {
																v746 = int32(-1)
															}
															v752 = v746
														} else {
															if v660 == int32(1) {
																v742 = v657
																v752 = v742
															} else {
																v674 = int32(1)
																v676 = int32(128)
																v677 = v664 << (uint(v674) % 32) & v676
																if v677 != v668<<(uint(v674)%32)&v676 {
																	v743 = v677
																	if v743 != 0 {
																		v746 = int32(1)
																	} else {
																		v746 = int32(-1)
																	}
																	v752 = v746
																} else {
																	if v660 < int32(3) {
																		v742 = v657
																		v752 = v742
																	} else {
																		v685 = int32(2)
																		v687 = int32(128)
																		v688 = v664 << (uint(v685) % 32) & v687
																		if v688 != v668<<(uint(v685)%32)&v687 {
																			v743 = v688
																			if v743 != 0 {
																				v746 = int32(1)
																			} else {
																				v746 = int32(-1)
																			}
																			v752 = v746
																		} else {
																			if v660 == int32(3) {
																				v742 = v657
																				v752 = v742
																			} else {
																				v696 = int32(3)
																				v698 = int32(128)
																				v699 = v664 << (uint(v696) % 32) & v698
																				if v699 != v668<<(uint(v696)%32)&v698 {
																					v743 = v699
																					if v743 != 0 {
																						v746 = int32(1)
																					} else {
																						v746 = int32(-1)
																					}
																					v752 = v746
																				} else {
																					if v660 < int32(5) {
																						v742 = v657
																						v752 = v742
																					} else {
																						v707 = int32(4)
																						v709 = int32(128)
																						v710 = v664 << (uint(v707) % 32) & v709
																						if v710 != v668<<(uint(v707)%32)&v709 {
																							v743 = v710
																							if v743 != 0 {
																								v746 = int32(1)
																							} else {
																								v746 = int32(-1)
																							}
																							v752 = v746
																						} else {
																							if v660 == int32(5) {
																								v742 = v657
																								v752 = v742
																							} else {
																								v718 = int32(5)
																								v720 = int32(128)
																								v721 = v664 << (uint(v718) % 32) & v720
																								if v721 != v668<<(uint(v718)%32)&v720 {
																									v743 = v721
																									if v743 != 0 {
																										v746 = int32(1)
																									} else {
																										v746 = int32(-1)
																									}
																									v752 = v746
																								} else {
																									if v660 < int32(7) {
																										v742 = v657
																										v752 = v742
																									} else {
																										v729 = int32(6)
																										v731 = int32(128)
																										v732 = v664 << (uint(v729) % 32) & v731
																										if v732 != v668<<(uint(v729)%32)&v731 {
																											v743 = v732
																											if v743 != 0 {
																												v746 = int32(1)
																											} else {
																												v746 = int32(-1)
																											}
																											v752 = v746
																										} else {
																											v742 = v657
																											v752 = v742
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
														}
													}
												}
												switch v18&int32(65535) - int32(18) {
												case 0:
													v802 = v752
													return base.B2i32(v802 == int32(0))
												case 1:
													v795 = v752
													return base.B2i32(v795 != int32(0))
												case 2:
													return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v752 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v752)
												case 5:
													return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v774 = m.ExcPending
													if v774 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(518516), int32(0))
														mBase = m.M
														v778 = m.ExcPending
														if v778 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494114), int32(327), int32(92161))
															mBase = m.M
															v783 = m.ExcPending
															if v783 != 0 {
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
												return int32(0)
											}
										}
									} else {
										v494 = int32(255)
										v495 = v489 & v494
										v497 = v488 & v494
										if base.Ui32(v495) < base.Ui32(v497) {
											v787 = v483
											return v787
										} else {
											if base.Ui32(v495) <= base.Ui32(v497) {
												v638 = v486
												if v638 != 0 {
													v642 = v44
												} else {
													v642 = v46
												}
												v644 = v642 + int32(2)
												v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
												if v647 == int32(3) {
													v650 = int32(128)
												} else {
													v650 = int32(32)
												}
												v655 = base.I32_div_s(v650, int32(8))
												v656 = F_memcmp(m, v93, v644, v655)
												mBase = m.M
												if v656 != 0 {
													v742 = v656
													v752 = v742
												} else {
													v657 = int32(0)
													v660 = v650 - v655<<(uint(int32(3))%32)
													if v660 <= v657 {
														v742 = v657
														v752 = v742
													} else {
														v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v655))))
														v665 = int32(128)
														v666 = v664 & v665
														v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v655))))
														if v666 != v668&v665 {
															v743 = v666
															if v743 != 0 {
																v746 = int32(1)
															} else {
																v746 = int32(-1)
															}
															v752 = v746
														} else {
															if v660 == int32(1) {
																v742 = v657
																v752 = v742
															} else {
																v674 = int32(1)
																v676 = int32(128)
																v677 = v664 << (uint(v674) % 32) & v676
																if v677 != v668<<(uint(v674)%32)&v676 {
																	v743 = v677
																	if v743 != 0 {
																		v746 = int32(1)
																	} else {
																		v746 = int32(-1)
																	}
																	v752 = v746
																} else {
																	if v660 < int32(3) {
																		v742 = v657
																		v752 = v742
																	} else {
																		v685 = int32(2)
																		v687 = int32(128)
																		v688 = v664 << (uint(v685) % 32) & v687
																		if v688 != v668<<(uint(v685)%32)&v687 {
																			v743 = v688
																			if v743 != 0 {
																				v746 = int32(1)
																			} else {
																				v746 = int32(-1)
																			}
																			v752 = v746
																		} else {
																			if v660 == int32(3) {
																				v742 = v657
																				v752 = v742
																			} else {
																				v696 = int32(3)
																				v698 = int32(128)
																				v699 = v664 << (uint(v696) % 32) & v698
																				if v699 != v668<<(uint(v696)%32)&v698 {
																					v743 = v699
																					if v743 != 0 {
																						v746 = int32(1)
																					} else {
																						v746 = int32(-1)
																					}
																					v752 = v746
																				} else {
																					if v660 < int32(5) {
																						v742 = v657
																						v752 = v742
																					} else {
																						v707 = int32(4)
																						v709 = int32(128)
																						v710 = v664 << (uint(v707) % 32) & v709
																						if v710 != v668<<(uint(v707)%32)&v709 {
																							v743 = v710
																							if v743 != 0 {
																								v746 = int32(1)
																							} else {
																								v746 = int32(-1)
																							}
																							v752 = v746
																						} else {
																							if v660 == int32(5) {
																								v742 = v657
																								v752 = v742
																							} else {
																								v718 = int32(5)
																								v720 = int32(128)
																								v721 = v664 << (uint(v718) % 32) & v720
																								if v721 != v668<<(uint(v718)%32)&v720 {
																									v743 = v721
																									if v743 != 0 {
																										v746 = int32(1)
																									} else {
																										v746 = int32(-1)
																									}
																									v752 = v746
																								} else {
																									if v660 < int32(7) {
																										v742 = v657
																										v752 = v742
																									} else {
																										v729 = int32(6)
																										v731 = int32(128)
																										v732 = v664 << (uint(v729) % 32) & v731
																										if v732 != v668<<(uint(v729)%32)&v731 {
																											v743 = v732
																											if v743 != 0 {
																												v746 = int32(1)
																											} else {
																												v746 = int32(-1)
																											}
																											v752 = v746
																										} else {
																											v742 = v657
																											v752 = v742
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
														}
													}
												}
												switch v18&int32(65535) - int32(18) {
												case 0:
													v802 = v752
													return base.B2i32(v802 == int32(0))
												case 1:
													v795 = v752
													return base.B2i32(v795 != int32(0))
												case 2:
													return int32(base.Ui32(v752) >> (uint(int32(31)) % 32))
												case 3:
													return base.B2i32(v752 <= int32(0))
												case 4:
													return base.B2i32(int32(0) < v752)
												case 5:
													return int32(base.Ui32(v752^int32(-1)) >> (uint(int32(31)) % 32))
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v774 = m.ExcPending
													if v774 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(518516), int32(0))
														mBase = m.M
														v778 = m.ExcPending
														if v778 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494114), int32(327), int32(92161))
															mBase = m.M
															v783 = m.ExcPending
															if v783 != 0 {
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
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				case 6:
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+16)))
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v53)+12)))
					if v55&int32(1) == int32(0) {
						v209 = v19 + int32(4)
						v211 = v47 + int32(2)
						v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v212) < base.Ui32(v213) {
							v215 = v212
						} else {
							v215 = v213
						}
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v215) < base.Ui32(v216) {
							v218 = v215
						} else {
							v218 = v216
						}
						v223 = base.I32_div_s(v218, int32(8))
						v224 = F_memcmp(m, v209, v211, v223)
						mBase = m.M
						if v224 != 0 {
							v310 = v224
							v320 = v310
						} else {
							v225 = int32(0)
							v228 = v218 - v223<<(uint(int32(3))%32)
							if v228 <= v225 {
								v310 = v225
								v320 = v310
							} else {
								v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v223))))
								v233 = int32(128)
								v234 = v232 & v233
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v223))))
								if v234 != v236&v233 {
									v311 = v234
									if v311 != 0 {
										v314 = int32(1)
									} else {
										v314 = int32(-1)
									}
									v320 = v314
								} else {
									if v228 == int32(1) {
										v310 = v225
										v320 = v310
									} else {
										v242 = int32(1)
										v244 = int32(128)
										v245 = v232 << (uint(v242) % 32) & v244
										if v245 != v236<<(uint(v242)%32)&v244 {
											v311 = v245
											if v311 != 0 {
												v314 = int32(1)
											} else {
												v314 = int32(-1)
											}
											v320 = v314
										} else {
											if v228 < int32(3) {
												v310 = v225
												v320 = v310
											} else {
												v253 = int32(2)
												v255 = int32(128)
												v256 = v232 << (uint(v253) % 32) & v255
												if v256 != v236<<(uint(v253)%32)&v255 {
													v311 = v256
													if v311 != 0 {
														v314 = int32(1)
													} else {
														v314 = int32(-1)
													}
													v320 = v314
												} else {
													if v228 == int32(3) {
														v310 = v225
														v320 = v310
													} else {
														v264 = int32(3)
														v266 = int32(128)
														v267 = v232 << (uint(v264) % 32) & v266
														if v267 != v236<<(uint(v264)%32)&v266 {
															v311 = v267
															if v311 != 0 {
																v314 = int32(1)
															} else {
																v314 = int32(-1)
															}
															v320 = v314
														} else {
															if v228 < int32(5) {
																v310 = v225
																v320 = v310
															} else {
																v275 = int32(4)
																v277 = int32(128)
																v278 = v232 << (uint(v275) % 32) & v277
																if v278 != v236<<(uint(v275)%32)&v277 {
																	v311 = v278
																	if v311 != 0 {
																		v314 = int32(1)
																	} else {
																		v314 = int32(-1)
																	}
																	v320 = v314
																} else {
																	if v228 == int32(5) {
																		v310 = v225
																		v320 = v310
																	} else {
																		v286 = int32(5)
																		v288 = int32(128)
																		v289 = v232 << (uint(v286) % 32) & v288
																		if v289 != v236<<(uint(v286)%32)&v288 {
																			v311 = v289
																			if v311 != 0 {
																				v314 = int32(1)
																			} else {
																				v314 = int32(-1)
																			}
																			v320 = v314
																		} else {
																			if v228 < int32(7) {
																				v310 = v225
																				v320 = v310
																			} else {
																				v297 = int32(6)
																				v299 = int32(128)
																				v300 = v232 << (uint(v297) % 32) & v299
																				if v300 != v236<<(uint(v297)%32)&v299 {
																					v311 = v300
																					if v311 != 0 {
																						v314 = int32(1)
																					} else {
																						v314 = int32(-1)
																					}
																					v320 = v314
																				} else {
																					v310 = v225
																					v320 = v310
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
								}
							}
						}
						v321 = v320
						return base.B2i32(v321 == int32(0))
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v61) < base.Ui32(v60) {
							v209 = v19 + int32(4)
							v211 = v47 + int32(2)
							v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
							if base.Ui32(v212) < base.Ui32(v213) {
								v215 = v212
							} else {
								v215 = v213
							}
							v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
							if base.Ui32(v215) < base.Ui32(v216) {
								v218 = v215
							} else {
								v218 = v216
							}
							v223 = base.I32_div_s(v218, int32(8))
							v224 = F_memcmp(m, v209, v211, v223)
							mBase = m.M
							if v224 != 0 {
								v310 = v224
								v320 = v310
							} else {
								v225 = int32(0)
								v228 = v218 - v223<<(uint(int32(3))%32)
								if v228 <= v225 {
									v310 = v225
									v320 = v310
								} else {
									v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v223))))
									v233 = int32(128)
									v234 = v232 & v233
									v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v223))))
									if v234 != v236&v233 {
										v311 = v234
										if v311 != 0 {
											v314 = int32(1)
										} else {
											v314 = int32(-1)
										}
										v320 = v314
									} else {
										if v228 == int32(1) {
											v310 = v225
											v320 = v310
										} else {
											v242 = int32(1)
											v244 = int32(128)
											v245 = v232 << (uint(v242) % 32) & v244
											if v245 != v236<<(uint(v242)%32)&v244 {
												v311 = v245
												if v311 != 0 {
													v314 = int32(1)
												} else {
													v314 = int32(-1)
												}
												v320 = v314
											} else {
												if v228 < int32(3) {
													v310 = v225
													v320 = v310
												} else {
													v253 = int32(2)
													v255 = int32(128)
													v256 = v232 << (uint(v253) % 32) & v255
													if v256 != v236<<(uint(v253)%32)&v255 {
														v311 = v256
														if v311 != 0 {
															v314 = int32(1)
														} else {
															v314 = int32(-1)
														}
														v320 = v314
													} else {
														if v228 == int32(3) {
															v310 = v225
															v320 = v310
														} else {
															v264 = int32(3)
															v266 = int32(128)
															v267 = v232 << (uint(v264) % 32) & v266
															if v267 != v236<<(uint(v264)%32)&v266 {
																v311 = v267
																if v311 != 0 {
																	v314 = int32(1)
																} else {
																	v314 = int32(-1)
																}
																v320 = v314
															} else {
																if v228 < int32(5) {
																	v310 = v225
																	v320 = v310
																} else {
																	v275 = int32(4)
																	v277 = int32(128)
																	v278 = v232 << (uint(v275) % 32) & v277
																	if v278 != v236<<(uint(v275)%32)&v277 {
																		v311 = v278
																		if v311 != 0 {
																			v314 = int32(1)
																		} else {
																			v314 = int32(-1)
																		}
																		v320 = v314
																	} else {
																		if v228 == int32(5) {
																			v310 = v225
																			v320 = v310
																		} else {
																			v286 = int32(5)
																			v288 = int32(128)
																			v289 = v232 << (uint(v286) % 32) & v288
																			if v289 != v236<<(uint(v286)%32)&v288 {
																				v311 = v289
																				if v311 != 0 {
																					v314 = int32(1)
																				} else {
																					v314 = int32(-1)
																				}
																				v320 = v314
																			} else {
																				if v228 < int32(7) {
																					v310 = v225
																					v320 = v310
																				} else {
																					v297 = int32(6)
																					v299 = int32(128)
																					v300 = v232 << (uint(v297) % 32) & v299
																					if v300 != v236<<(uint(v297)%32)&v299 {
																						v311 = v300
																						if v311 != 0 {
																							v314 = int32(1)
																						} else {
																							v314 = int32(-1)
																						}
																						v320 = v314
																					} else {
																						v310 = v225
																						v320 = v310
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
									}
								}
							}
							v321 = v320
							return base.B2i32(v321 == int32(0))
						} else {
							return int32(0)
						}
					}
				case 7:
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v66)+12)))
					if v68&int32(1) == int32(0) {
						v209 = v19 + int32(4)
						v211 = v47 + int32(2)
						v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v212) < base.Ui32(v213) {
							v215 = v212
						} else {
							v215 = v213
						}
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v215) < base.Ui32(v216) {
							v218 = v215
						} else {
							v218 = v216
						}
						v223 = base.I32_div_s(v218, int32(8))
						v224 = F_memcmp(m, v209, v211, v223)
						mBase = m.M
						if v224 != 0 {
							v310 = v224
							v320 = v310
						} else {
							v225 = int32(0)
							v228 = v218 - v223<<(uint(int32(3))%32)
							if v228 <= v225 {
								v310 = v225
								v320 = v310
							} else {
								v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v223))))
								v233 = int32(128)
								v234 = v232 & v233
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v223))))
								if v234 != v236&v233 {
									v311 = v234
									if v311 != 0 {
										v314 = int32(1)
									} else {
										v314 = int32(-1)
									}
									v320 = v314
								} else {
									if v228 == int32(1) {
										v310 = v225
										v320 = v310
									} else {
										v242 = int32(1)
										v244 = int32(128)
										v245 = v232 << (uint(v242) % 32) & v244
										if v245 != v236<<(uint(v242)%32)&v244 {
											v311 = v245
											if v311 != 0 {
												v314 = int32(1)
											} else {
												v314 = int32(-1)
											}
											v320 = v314
										} else {
											if v228 < int32(3) {
												v310 = v225
												v320 = v310
											} else {
												v253 = int32(2)
												v255 = int32(128)
												v256 = v232 << (uint(v253) % 32) & v255
												if v256 != v236<<(uint(v253)%32)&v255 {
													v311 = v256
													if v311 != 0 {
														v314 = int32(1)
													} else {
														v314 = int32(-1)
													}
													v320 = v314
												} else {
													if v228 == int32(3) {
														v310 = v225
														v320 = v310
													} else {
														v264 = int32(3)
														v266 = int32(128)
														v267 = v232 << (uint(v264) % 32) & v266
														if v267 != v236<<(uint(v264)%32)&v266 {
															v311 = v267
															if v311 != 0 {
																v314 = int32(1)
															} else {
																v314 = int32(-1)
															}
															v320 = v314
														} else {
															if v228 < int32(5) {
																v310 = v225
																v320 = v310
															} else {
																v275 = int32(4)
																v277 = int32(128)
																v278 = v232 << (uint(v275) % 32) & v277
																if v278 != v236<<(uint(v275)%32)&v277 {
																	v311 = v278
																	if v311 != 0 {
																		v314 = int32(1)
																	} else {
																		v314 = int32(-1)
																	}
																	v320 = v314
																} else {
																	if v228 == int32(5) {
																		v310 = v225
																		v320 = v310
																	} else {
																		v286 = int32(5)
																		v288 = int32(128)
																		v289 = v232 << (uint(v286) % 32) & v288
																		if v289 != v236<<(uint(v286)%32)&v288 {
																			v311 = v289
																			if v311 != 0 {
																				v314 = int32(1)
																			} else {
																				v314 = int32(-1)
																			}
																			v320 = v314
																		} else {
																			if v228 < int32(7) {
																				v310 = v225
																				v320 = v310
																			} else {
																				v297 = int32(6)
																				v299 = int32(128)
																				v300 = v232 << (uint(v297) % 32) & v299
																				if v300 != v236<<(uint(v297)%32)&v299 {
																					v311 = v300
																					if v311 != 0 {
																						v314 = int32(1)
																					} else {
																						v314 = int32(-1)
																					}
																					v320 = v314
																				} else {
																					v310 = v225
																					v320 = v310
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
								}
							}
						}
						v321 = v320
						return base.B2i32(v321 == int32(0))
					} else {
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v74) <= base.Ui32(v73) {
							v209 = v19 + int32(4)
							v211 = v47 + int32(2)
							v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
							v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
							if base.Ui32(v212) < base.Ui32(v213) {
								v215 = v212
							} else {
								v215 = v213
							}
							v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
							if base.Ui32(v215) < base.Ui32(v216) {
								v218 = v215
							} else {
								v218 = v216
							}
							v223 = base.I32_div_s(v218, int32(8))
							v224 = F_memcmp(m, v209, v211, v223)
							mBase = m.M
							if v224 != 0 {
								v310 = v224
								v320 = v310
							} else {
								v225 = int32(0)
								v228 = v218 - v223<<(uint(int32(3))%32)
								if v228 <= v225 {
									v310 = v225
									v320 = v310
								} else {
									v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v223))))
									v233 = int32(128)
									v234 = v232 & v233
									v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v223))))
									if v234 != v236&v233 {
										v311 = v234
										if v311 != 0 {
											v314 = int32(1)
										} else {
											v314 = int32(-1)
										}
										v320 = v314
									} else {
										if v228 == int32(1) {
											v310 = v225
											v320 = v310
										} else {
											v242 = int32(1)
											v244 = int32(128)
											v245 = v232 << (uint(v242) % 32) & v244
											if v245 != v236<<(uint(v242)%32)&v244 {
												v311 = v245
												if v311 != 0 {
													v314 = int32(1)
												} else {
													v314 = int32(-1)
												}
												v320 = v314
											} else {
												if v228 < int32(3) {
													v310 = v225
													v320 = v310
												} else {
													v253 = int32(2)
													v255 = int32(128)
													v256 = v232 << (uint(v253) % 32) & v255
													if v256 != v236<<(uint(v253)%32)&v255 {
														v311 = v256
														if v311 != 0 {
															v314 = int32(1)
														} else {
															v314 = int32(-1)
														}
														v320 = v314
													} else {
														if v228 == int32(3) {
															v310 = v225
															v320 = v310
														} else {
															v264 = int32(3)
															v266 = int32(128)
															v267 = v232 << (uint(v264) % 32) & v266
															if v267 != v236<<(uint(v264)%32)&v266 {
																v311 = v267
																if v311 != 0 {
																	v314 = int32(1)
																} else {
																	v314 = int32(-1)
																}
																v320 = v314
															} else {
																if v228 < int32(5) {
																	v310 = v225
																	v320 = v310
																} else {
																	v275 = int32(4)
																	v277 = int32(128)
																	v278 = v232 << (uint(v275) % 32) & v277
																	if v278 != v236<<(uint(v275)%32)&v277 {
																		v311 = v278
																		if v311 != 0 {
																			v314 = int32(1)
																		} else {
																			v314 = int32(-1)
																		}
																		v320 = v314
																	} else {
																		if v228 == int32(5) {
																			v310 = v225
																			v320 = v310
																		} else {
																			v286 = int32(5)
																			v288 = int32(128)
																			v289 = v232 << (uint(v286) % 32) & v288
																			if v289 != v236<<(uint(v286)%32)&v288 {
																				v311 = v289
																				if v311 != 0 {
																					v314 = int32(1)
																				} else {
																					v314 = int32(-1)
																				}
																				v320 = v314
																			} else {
																				if v228 < int32(7) {
																					v310 = v225
																					v320 = v310
																				} else {
																					v297 = int32(6)
																					v299 = int32(128)
																					v300 = v232 << (uint(v297) % 32) & v299
																					if v300 != v236<<(uint(v297)%32)&v299 {
																						v311 = v300
																						if v311 != 0 {
																							v314 = int32(1)
																						} else {
																							v314 = int32(-1)
																						}
																						v320 = v314
																					} else {
																						v310 = v225
																						v320 = v310
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
									}
								}
							}
							v321 = v320
							return base.B2i32(v321 == int32(0))
						} else {
							return int32(0)
						}
					}
				case 8:
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
					if base.Ui32(v83) < base.Ui32(v84) {
						v209 = v19 + int32(4)
						v211 = v47 + int32(2)
						v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
						if base.Ui32(v212) < base.Ui32(v213) {
							v215 = v212
						} else {
							v215 = v213
						}
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						if base.Ui32(v215) < base.Ui32(v216) {
							v218 = v215
						} else {
							v218 = v216
						}
						v223 = base.I32_div_s(v218, int32(8))
						v224 = F_memcmp(m, v209, v211, v223)
						mBase = m.M
						if v224 != 0 {
							v310 = v224
							v320 = v310
						} else {
							v225 = int32(0)
							v228 = v218 - v223<<(uint(int32(3))%32)
							if v228 <= v225 {
								v310 = v225
								v320 = v310
							} else {
								v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v223))))
								v233 = int32(128)
								v234 = v232 & v233
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v223))))
								if v234 != v236&v233 {
									v311 = v234
									if v311 != 0 {
										v314 = int32(1)
									} else {
										v314 = int32(-1)
									}
									v320 = v314
								} else {
									if v228 == int32(1) {
										v310 = v225
										v320 = v310
									} else {
										v242 = int32(1)
										v244 = int32(128)
										v245 = v232 << (uint(v242) % 32) & v244
										if v245 != v236<<(uint(v242)%32)&v244 {
											v311 = v245
											if v311 != 0 {
												v314 = int32(1)
											} else {
												v314 = int32(-1)
											}
											v320 = v314
										} else {
											if v228 < int32(3) {
												v310 = v225
												v320 = v310
											} else {
												v253 = int32(2)
												v255 = int32(128)
												v256 = v232 << (uint(v253) % 32) & v255
												if v256 != v236<<(uint(v253)%32)&v255 {
													v311 = v256
													if v311 != 0 {
														v314 = int32(1)
													} else {
														v314 = int32(-1)
													}
													v320 = v314
												} else {
													if v228 == int32(3) {
														v310 = v225
														v320 = v310
													} else {
														v264 = int32(3)
														v266 = int32(128)
														v267 = v232 << (uint(v264) % 32) & v266
														if v267 != v236<<(uint(v264)%32)&v266 {
															v311 = v267
															if v311 != 0 {
																v314 = int32(1)
															} else {
																v314 = int32(-1)
															}
															v320 = v314
														} else {
															if v228 < int32(5) {
																v310 = v225
																v320 = v310
															} else {
																v275 = int32(4)
																v277 = int32(128)
																v278 = v232 << (uint(v275) % 32) & v277
																if v278 != v236<<(uint(v275)%32)&v277 {
																	v311 = v278
																	if v311 != 0 {
																		v314 = int32(1)
																	} else {
																		v314 = int32(-1)
																	}
																	v320 = v314
																} else {
																	if v228 == int32(5) {
																		v310 = v225
																		v320 = v310
																	} else {
																		v286 = int32(5)
																		v288 = int32(128)
																		v289 = v232 << (uint(v286) % 32) & v288
																		if v289 != v236<<(uint(v286)%32)&v288 {
																			v311 = v289
																			if v311 != 0 {
																				v314 = int32(1)
																			} else {
																				v314 = int32(-1)
																			}
																			v320 = v314
																		} else {
																			if v228 < int32(7) {
																				v310 = v225
																				v320 = v310
																			} else {
																				v297 = int32(6)
																				v299 = int32(128)
																				v300 = v232 << (uint(v297) % 32) & v299
																				if v300 != v236<<(uint(v297)%32)&v299 {
																					v311 = v300
																					if v311 != 0 {
																						v314 = int32(1)
																					} else {
																						v314 = int32(-1)
																					}
																					v320 = v314
																				} else {
																					v310 = v225
																					v320 = v310
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
								}
							}
						}
						v321 = v320
						return base.B2i32(v321 == int32(0))
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
