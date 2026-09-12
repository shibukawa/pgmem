package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_start_xact_command(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[894])))
	if v3 == int32(0) {
		F_StartTransactionCommand(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[894])) = uint8(v9)
			v26 = *(*int32)(unsafe.Add(mBase, _consts[895]))
			if v26 <= int32(0) {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
				if v51 == int32(0) {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
					if v59 <= int32(0) {
						return
					} else {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
						if v63 != int32(1) {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
							if v67 == int32(0) {
								return
							} else {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
									F_enable_timeout_after(m, int32(11), v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					F_disable_timeout(m, int32(3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
						if v59 <= int32(0) {
							return
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
							if v63 != int32(1) {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
								if v67 == int32(0) {
									return
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
									if v75 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
										F_enable_timeout_after(m, int32(11), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[899]))
				if v30 != 0 {
					v33 = base.B2i32(v30 <= v26)
				} else {
					v33 = int32(0)
				}
				if v33 != 0 {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
					if v51 == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
						if v59 <= int32(0) {
							return
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
							if v63 != int32(1) {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
								if v67 == int32(0) {
									return
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
									if v75 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
										F_enable_timeout_after(m, int32(11), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						F_disable_timeout(m, int32(3))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
							if v59 <= int32(0) {
								return
							} else {
								v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
								if v63 != int32(1) {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
									if v67 == int32(0) {
										return
									} else {
										v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
										if v75 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
											F_enable_timeout_after(m, int32(11), v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
					if v39 != 0 {
						v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
						if v59 <= int32(0) {
							return
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
							if v63 != int32(1) {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
								if v67 == int32(0) {
									return
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
									if v75 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
										F_enable_timeout_after(m, int32(11), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[895]))
						F_enable_timeout_after(m, int32(3), v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
							if v59 <= int32(0) {
								return
							} else {
								v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
								if v63 != int32(1) {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
									if v67 == int32(0) {
										return
									} else {
										v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
										if v75 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
											F_enable_timeout_after(m, int32(11), v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												return
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
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[430])))
		if v12&int32(8) == int32(0) {
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[72]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			if v20 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(4)
			} else {
			}
		}
		v26 = *(*int32)(unsafe.Add(mBase, _consts[895]))
		if v26 <= int32(0) {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
			if v51 == int32(0) {
				v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
				if v59 <= int32(0) {
					return
				} else {
					v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
					if v63 != int32(1) {
						return
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
						if v67 == int32(0) {
							return
						} else {
							v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
							if v75 != 0 {
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
								F_enable_timeout_after(m, int32(11), v78)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				F_disable_timeout(m, int32(3))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
					if v59 <= int32(0) {
						return
					} else {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
						if v63 != int32(1) {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
							if v67 == int32(0) {
								return
							} else {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
									F_enable_timeout_after(m, int32(11), v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[899]))
			if v30 != 0 {
				v33 = base.B2i32(v30 <= v26)
			} else {
				v33 = int32(0)
			}
			if v33 != 0 {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
				if v51 == int32(0) {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
					if v59 <= int32(0) {
						return
					} else {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
						if v63 != int32(1) {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
							if v67 == int32(0) {
								return
							} else {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
									F_enable_timeout_after(m, int32(11), v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					F_disable_timeout(m, int32(3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
						if v59 <= int32(0) {
							return
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
							if v63 != int32(1) {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
								if v67 == int32(0) {
									return
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
									if v75 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
										F_enable_timeout_after(m, int32(11), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[896])))
				if v39 != 0 {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
					if v59 <= int32(0) {
						return
					} else {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
						if v63 != int32(1) {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
							if v67 == int32(0) {
								return
							} else {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
								if v75 != 0 {
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
									F_enable_timeout_after(m, int32(11), v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[895]))
					F_enable_timeout_after(m, int32(3), v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, _consts[897]))
						if v59 <= int32(0) {
							return
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
							if v63 != int32(1) {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, _consts[579]))
								if v67 == int32(0) {
									return
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[898])))
									if v75 != 0 {
										return
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _consts[897]))
										F_enable_timeout_after(m, int32(11), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											return
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
func F_xact_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(28))+uint32(_consts[79])))
	return v8
}
func F_xact_redo(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int64
	_ = v592
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int64
	_ = v691
	var v692 int64
	_ = v692
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int64
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	v7 = m.G0
	v9 = v7 - int32(304)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	v14 = v12 & int32(112)
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		goto L5
	case 1:
		goto L7
	case 2:
		goto L8
	case 3:
		goto L6
	case 4:
		goto L4
	case 5:
		goto L1
	case 6:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(304)
	return
L2:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v586 = int32(0)
	v591 = F___memset(m, v9+int32(16), v586, int32(288))
	mBase = m.M
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v583)))
	*(*int64)(unsafe.Add(mBase, uint32(v591))) = v592
	if v586 <= base.I32_extend8_s(v12) {
		goto L139
	} else {
		goto L140
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L31
	} else {
		goto L135
	}
L4:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if v421 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v405 = F_LWLockAcquire(m, v401+int32(2304), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L31
	} else {
		goto L79
	}
L6:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v275 = int32(0)
	v280 = F___memset(m, v9+int32(16), v275, int32(264))
	mBase = m.M
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v272)))
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v281
	if v275 <= base.I32_extend8_s(v12) {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v164 = int32(0)
	v169 = F___memset(m, v9+int32(16), v164, int32(264))
	mBase = m.M
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v170
	if v164 <= base.I32_extend8_s(v12) {
		goto L37
	} else {
		goto L38
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v22 = int32(0)
	v27 = F___memset(m, v9+int32(16), v22, int32(288))
	mBase = m.M
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v28
	if v22 <= base.I32_extend8_s(v12) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+56)))
	F_xact_redo_commit(m, v9+int32(16), v138, v139, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	goto L9
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v33
	if v33&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v39
	v45 = v19 + int32(20)
	goto L14
L13:
	;
	v45 = v19 + int32(12)
	goto L14
L14:
	;
	if v33&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v50 = v45 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v48
	v56 = v50 + v48<<(uint(int32(2))%32)
	goto L17
L16:
	;
	v56 = v45
	goto L17
L17:
	;
	if v33&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v62 = v56 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v60
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v69 = v62 + v65*int32(12)
	goto L20
L19:
	;
	v69 = v56
	goto L20
L20:
	;
	if v33&int32(256) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v75 = int32(4)
	v76 = v69 + v75
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v74
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v83 = v76 + v79<<(uint(v75)%32)
	goto L23
L22:
	;
	v83 = v69
	goto L23
L23:
	;
	if v33&int32(8) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = int32(4)
	v90 = v83 + v89
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v88
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v97 = v90 + v93<<(uint(v89)%32)
	goto L26
L25:
	;
	v97 = v83
	goto L26
L26:
	;
	if v33&int32(16) == int32(0) {
		v121 = v33
		v122 = v97
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v121&int32(32) == int32(0) {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v104
	v107 = v97 + int32(4)
	if v33&int32(128) == int32(0) {
		v121 = v33
		v122 = v107
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v115 = F_strlcpy(m, v27+int32(56), v107, int32(200))
	mBase = m.M
	v116 = F_strlen(m, v107)
	mBase = m.M
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v121 = v120
	v122 = v116 + v107 + int32(1)
	goto L27
L30:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v122)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+280)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v27)+272)) = v127
	goto L10
L31:
	;
	return
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v149 = F_LWLockAcquire(m, v145+int32(2304), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	F_PrepareRedoRemove(m, v151, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v156+int32(2304))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+56)))
	F_xact_redo_abort(m, v9+int32(16), v267, v268, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L31
	} else {
		goto L55
	}
L37:
	;
	goto L36
L38:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v175
	if v175&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+12)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+16)) = v181
	v187 = v161 + int32(20)
	goto L41
L40:
	;
	v187 = v161 + int32(12)
	goto L41
L41:
	;
	if v175&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v192 = v187 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+24)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v169)+20)) = v190
	v198 = v192 + v190<<(uint(int32(2))%32)
	goto L44
L43:
	;
	v198 = v187
	goto L44
L44:
	;
	if v175&int32(4) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v204 = v198 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+32)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v169)+28)) = v202
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v211 = v204 + v207*int32(12)
	goto L47
L46:
	;
	v211 = v198
	goto L47
L47:
	;
	if v175&int32(256) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v217 = int32(4)
	v218 = v211 + v217
	*(*int32)(unsafe.Add(mBase, uint32(v169)+40)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v169)+36)) = v216
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v225 = v218 + v221<<(uint(v217)%32)
	goto L50
L49:
	;
	v225 = v211
	goto L50
L50:
	;
	if v175&int32(16) == int32(0) {
		v249 = v175
		v250 = v225
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v249&int32(32) == int32(0) {
		goto L37
	} else {
		goto L54
	}
L52:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+44)) = v232
	v235 = v225 + int32(4)
	if v175&int32(128) == int32(0) {
		v249 = v175
		v250 = v235
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v243 = F_strlcpy(m, v169+int32(48), v235, int32(200))
	mBase = m.M
	v244 = F_strlen(m, v235)
	mBase = m.M
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v249 = v248
	v250 = v244 + v235 + int32(1)
	goto L51
L54:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v250)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v169)+256)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v169)+248)) = v255
	goto L37
L55:
	;
	goto L1
L56:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+56)))
	F_xact_redo_abort(m, v9+int32(16), v377, v378, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L31
	} else {
		goto L75
	}
L57:
	;
	goto L56
L58:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v286
	if v286&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v292
	v298 = v272 + int32(20)
	goto L61
L60:
	;
	v298 = v272 + int32(12)
	goto L61
L61:
	;
	if v286&int32(2) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v303 = v298 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+24)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v280)+20)) = v301
	v309 = v303 + v301<<(uint(int32(2))%32)
	goto L64
L63:
	;
	v309 = v298
	goto L64
L64:
	;
	if v286&int32(4) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v315 = v309 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+32)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v280)+28)) = v313
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v322 = v315 + v318*int32(12)
	goto L67
L66:
	;
	v322 = v309
	goto L67
L67:
	;
	if v286&int32(256) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v328 = int32(4)
	v329 = v322 + v328
	*(*int32)(unsafe.Add(mBase, uint32(v280)+40)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v280)+36)) = v327
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v336 = v329 + v332<<(uint(v328)%32)
	goto L70
L69:
	;
	v336 = v322
	goto L70
L70:
	;
	if v286&int32(16) == int32(0) {
		v360 = v286
		v361 = v336
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v360&int32(32) == int32(0) {
		goto L57
	} else {
		goto L74
	}
L72:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+44)) = v343
	v346 = v336 + int32(4)
	if v286&int32(128) == int32(0) {
		v360 = v286
		v361 = v346
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v354 = F_strlcpy(m, v280+int32(48), v346, int32(200))
	mBase = m.M
	v355 = F_strlen(m, v346)
	mBase = m.M
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v360 = v359
	v361 = v355 + v346 + int32(1)
	goto L71
L74:
	;
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v361)))
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v361)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+256)) = v367
	*(*int64)(unsafe.Add(mBase, uint32(v280)+248)) = v366
	goto L57
L75:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v388 = F_LWLockAcquire(m, v384+int32(2304), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L31
	} else {
		goto L76
	}
L76:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	F_PrepareRedoRemove(m, v390, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v395+int32(2304))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+64))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+56)))
	F_PrepareRedoAdd(m, v408, v409, v410, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v415+int32(2304))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L31
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v424 = int32(0)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v429 = v425 + int32(8)
	v434 = v427 - int32(1)
	if v434 < v424 {
		v501 = v426
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_RecordKnownAssignedTransactionIds(m, v501)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L31
	} else {
		goto L114
	}
L84:
	;
	goto L83
L85:
	;
	if v427&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v439 = int32(2)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v429+v434<<(uint(v439)%32))))
	if base.B2i32(base.Ui32(v439) < base.Ui32(v442))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v426)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v457 = v426
	v459 = v434
	goto L88
L88:
	;
	if v434 == int32(0) {
		v501 = v457
		goto L84
	} else {
		goto L96
	}
L89:
	;
	v457 = v454
	v459 = v427 - int32(2)
	goto L88
L90:
	;
	v454 = v442
	goto L89
L91:
	;
	if base.Ui32(v426) < base.Ui32(v442) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if int32(0) <= v426-v442 {
		v454 = v426
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v454 = v426
	goto L89
L95:
	;
	goto L90
L96:
	;
	v464 = v457
	v465 = v459
	goto L97
L97:
	;
	v471 = v465 << (uint(int32(2)) % 32)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v429+v471)))
	if base.Ui32(v464) < base.Ui32(int32(3)) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v501 = v496
	goto L84
L99:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v471+(v429-int32(4)))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v484))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v482)) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L100:
	;
	v482 = v473
	goto L99
L101:
	;
	if base.Ui32(v473) <= base.Ui32(v464) {
		v482 = v464
		goto L99
	} else {
		goto L105
	}
L102:
	;
	if base.Ui32(v473) < base.Ui32(int32(3)) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	if v464-v473 < int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v482 = v464
	goto L99
L105:
	;
	goto L100
L106:
	;
	if int32(1) < v465 {
		v464 = v496
		v465 = v465 - int32(2)
		goto L97
	} else {
		goto L113
	}
L107:
	;
	v496 = v484
	goto L106
L108:
	;
	if base.Ui32(v482) < base.Ui32(v484) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if int32(0) <= v482-v484 {
		v496 = v482
		goto L106
	} else {
		goto L112
	}
L111:
	;
	v496 = v482
	goto L106
L112:
	;
	goto L107
L113:
	;
	goto L98
L114:
	;
	if int32(0) < v427 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v514 = v424
	goto L118
L116:
	;
	goto L117
L117:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if v533 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v429+v514<<(uint(int32(2))%32))))
	F_SubTransSetParent(m, v520, v426)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L31
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v524 = v514 + int32(1)
	if v524 != v427 {
		v514 = v524
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v541 = F_LWLockAcquire(m, v537+int32(512), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L31
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L1
L125:
	;
	F_KnownAssignedXidsRemoveTree(m, int32(0), v427, v429)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L31
	} else {
		goto L126
	}
L126:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v501))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v548)) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v560 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v560 = base.B2i32(base.Ui32(v548) < base.Ui32(v501))
	goto L127
L129:
	;
	goto L130
L130:
	;
	v560 = int32(base.Ui32(v548-v501) >> (uint(int32(31)) % 32))
	goto L127
L131:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	*(*int32)(unsafe.Add(mBase, uint32(v562)+24)) = v501
	goto L133
L132:
	;
	goto L133
L133:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v565+int32(512))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L31
	} else {
		goto L134
	}
L134:
	;
	goto L124
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(53328), v9)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L31
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(504485), int32(6445), int32(247780))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L31
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+36))
	v704 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v702)+56)))
	F_xact_redo_commit(m, v9+int32(16), v703, v704, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L31
	} else {
		goto L160
	}
L139:
	;
	goto L138
L140:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v583)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+8)) = v597
	if v597&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+12)) = v601
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+16)) = v603
	v609 = v583 + int32(20)
	goto L143
L142:
	;
	v609 = v583 + int32(12)
	goto L143
L143:
	;
	if v597&int32(2) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v614 = v609 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+24)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v591)+20)) = v612
	v620 = v614 + v612<<(uint(int32(2))%32)
	goto L146
L145:
	;
	v620 = v609
	goto L146
L146:
	;
	if v597&int32(4) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v626 = v620 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+32)) = v626
	*(*int32)(unsafe.Add(mBase, uint32(v591)+28)) = v624
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v633 = v626 + v629*int32(12)
	goto L149
L148:
	;
	v633 = v620
	goto L149
L149:
	;
	if v597&int32(256) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v639 = int32(4)
	v640 = v633 + v639
	*(*int32)(unsafe.Add(mBase, uint32(v591)+40)) = v640
	*(*int32)(unsafe.Add(mBase, uint32(v591)+36)) = v638
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v647 = v640 + v643<<(uint(v639)%32)
	goto L152
L151:
	;
	v647 = v633
	goto L152
L152:
	;
	if v597&int32(8) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	v653 = int32(4)
	v654 = v647 + v653
	*(*int32)(unsafe.Add(mBase, uint32(v591)+48)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v591)+44)) = v652
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	v661 = v654 + v657<<(uint(v653)%32)
	goto L155
L154:
	;
	v661 = v647
	goto L155
L155:
	;
	if v597&int32(16) == int32(0) {
		v685 = v597
		v686 = v661
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v685&int32(32) == int32(0) {
		goto L139
	} else {
		goto L159
	}
L157:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+52)) = v668
	v671 = v661 + int32(4)
	if v597&int32(128) == int32(0) {
		v685 = v597
		v686 = v671
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v679 = F_strlcpy(m, v591+int32(56), v671, int32(200))
	mBase = m.M
	v680 = F_strlen(m, v671)
	mBase = m.M
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	v685 = v684
	v686 = v680 + v671 + int32(1)
	goto L156
L159:
	;
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v686)))
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v686)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v591)+280)) = v692
	*(*int64)(unsafe.Add(mBase, uint32(v591)+272)) = v691
	goto L139
L160:
	;
	goto L1
}
func F_xact_redo_commit(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = v6 - int32(1)
	if v12 < int32(0) {
		v79 = l1
	} else {
		if v6&int32(1) != 0 {
			v17 = int32(2)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7+v12<<(uint(v17)%32))))
			if base.B2i32(base.Ui32(v17) < base.Ui32(v20))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				if base.Ui32(l1) < base.Ui32(v20) {
					v32 = v20
				} else {
					v32 = l1
				}
			} else {
				if int32(0) <= l1-v20 {
					v32 = l1
				} else {
					v32 = v20
				}
			}
			v35 = v32
			v37 = v6 - int32(2)
		} else {
			v35 = l1
			v37 = v12
		}
		if v12 == int32(0) {
			v79 = v35
		} else {
			v42 = v35
			v43 = v37
			for {
				v49 = v43 << (uint(int32(2)) % 32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7+v49)))
				if base.Ui32(v42) < base.Ui32(int32(3)) {
					if base.Ui32(v51) <= base.Ui32(v42) {
						v60 = v42
					} else {
						v60 = v51
					}
				} else {
					if base.Ui32(v51) < base.Ui32(int32(3)) {
						if base.Ui32(v51) <= base.Ui32(v42) {
							v60 = v42
						} else {
							v60 = v51
						}
					} else {
						if v42-v51 < int32(0) {
							v60 = v51
						} else {
							v60 = v42
						}
					}
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v7-int32(4)))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v62))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v60)) == int32(0) {
					if base.Ui32(v60) < base.Ui32(v62) {
						v74 = v62
					} else {
						v74 = v60
					}
				} else {
					if int32(0) <= v60-v62 {
						v74 = v60
					} else {
						v74 = v62
					}
				}
				if int32(1) < v43 {
					v42 = v74
					v43 = v43 - int32(2)
					continue
				} else {
					break
				}
				break
			}
			v79 = v74
		}
	}
	F_AdvanceNextFullTransactionIdPastXid(m, v79)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		return
	} else {
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v97 = *(*int64)(unsafe.Add(mBase, uint32(l0+v89<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(280))))
		F_TransactionTreeSetCommitTsData(m, l1, v87, v88, v97, l3)
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return
		} else {
			v101 = *(*int32)(unsafe.Add(mBase, _consts[33]))
			if v101 == int32(0) {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_TransactionIdCommitTree(m, l1, v104, v105)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
					if v138&int32(32) != 0 {
						v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
						v142 = int32(0)
						F_replorigin_advance(m, l3, v141, l2, v142, v142)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return
						} else {
							v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if int32(0) < v146 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return
								} else {
									v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_DropRelationFiles(m, v151, v152, int32(1))
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return
									} else {
										v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v156 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v161, v162)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v165 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v171 = v170
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													} else {
														v171 = v165
														if v171&int32(536870912) != 0 {
															v175 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
														} else {
														}
														return
													}
												}
											}
										} else {
											v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v165 < int32(0) {
												F_XLogFlush(m, l2)
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return
												} else {
													v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v171 = v170
													if v171&int32(536870912) != 0 {
														v175 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
													} else {
													}
													return
												}
											} else {
												v171 = v165
												if v171&int32(536870912) != 0 {
													v175 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
												} else {
												}
												return
											}
										}
									}
								}
							} else {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if int32(0) < v156 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										F_pgstat_execute_transactional_drops(m, v161, v162)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v165 < int32(0) {
												F_XLogFlush(m, l2)
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return
												} else {
													v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v171 = v170
													if v171&int32(536870912) != 0 {
														v175 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
													} else {
													}
													return
												}
											} else {
												v171 = v165
												if v171&int32(536870912) != 0 {
													v175 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
												} else {
												}
												return
											}
										}
									}
								} else {
									v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v165 < int32(0) {
										F_XLogFlush(m, l2)
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return
										} else {
											v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v171 = v170
											if v171&int32(536870912) != 0 {
												v175 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
											} else {
											}
											return
										}
									} else {
										v171 = v165
										if v171&int32(536870912) != 0 {
											v175 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
										} else {
										}
										return
									}
								}
							}
						}
					} else {
						v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if int32(0) < v146 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_DropRelationFiles(m, v151, v152, int32(1))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return
								} else {
									v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v156 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v161, v162)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v165 < int32(0) {
													F_XLogFlush(m, l2)
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v171 = v170
														if v171&int32(536870912) != 0 {
															v175 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
														} else {
														}
														return
													}
												} else {
													v171 = v165
													if v171&int32(536870912) != 0 {
														v175 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
													} else {
													}
													return
												}
											}
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v165 < int32(0) {
											F_XLogFlush(m, l2)
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v171 = v170
												if v171&int32(536870912) != 0 {
													v175 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
												} else {
												}
												return
											}
										} else {
											v171 = v165
											if v171&int32(536870912) != 0 {
												v175 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
											} else {
											}
											return
										}
									}
								}
							}
						} else {
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if int32(0) < v156 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_pgstat_execute_transactional_drops(m, v161, v162)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v165 < int32(0) {
											F_XLogFlush(m, l2)
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v171 = v170
												if v171&int32(536870912) != 0 {
													v175 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
												} else {
												}
												return
											}
										} else {
											v171 = v165
											if v171&int32(536870912) != 0 {
												v175 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
											} else {
											}
											return
										}
									}
								}
							} else {
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v165 < int32(0) {
									F_XLogFlush(m, l2)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return
									} else {
										v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v171 = v170
										if v171&int32(536870912) != 0 {
											v175 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
										} else {
										}
										return
									}
								} else {
									v171 = v165
									if v171&int32(536870912) != 0 {
										v175 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
									} else {
									}
									return
								}
							}
						}
					}
				}
			} else {
				F_RecordKnownAssignedTransactionIds(m, v79)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_TransactionIdAsyncCommitTree(m, l1, v110, v111, l2)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_ExpireTreeKnownAssignedTransactionIds(m, l1, v114, v115, v79)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_ProcessCommittedInvalidationMessages(m, v118, v119, int32(base.Ui32(v120&int32(1073741824))>>(uint(int32(30))%32)), v125, v126)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v129&int32(64) == int32(0) {
									v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
									if v138&int32(32) != 0 {
										v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
										v142 = int32(0)
										F_replorigin_advance(m, l3, v141, l2, v142, v142)
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
											return
										} else {
											v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if int32(0) < v146 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_DropRelationFiles(m, v151, v152, int32(1))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if int32(0) < v156 {
															F_XLogFlush(m, l2)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																F_pgstat_execute_transactional_drops(m, v161, v162)
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return
																} else {
																	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	if v165 < int32(0) {
																		F_XLogFlush(m, l2)
																		mBase = m.M
																		v169 = m.ExcPending
																		if v169 != 0 {
																			return
																		} else {
																			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			v171 = v170
																			if v171&int32(536870912) != 0 {
																				v175 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																			} else {
																			}
																			return
																		}
																	} else {
																		v171 = v165
																		if v171&int32(536870912) != 0 {
																			v175 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																		} else {
																		}
																		return
																	}
																}
															}
														} else {
															v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v165 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v171 = v170
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															} else {
																v171 = v165
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														}
													}
												}
											} else {
												v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v156 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v161, v162)
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v165 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v171 = v170
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															} else {
																v171 = v165
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														}
													}
												} else {
													v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v165 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v171 = v170
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													} else {
														v171 = v165
														if v171&int32(536870912) != 0 {
															v175 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
														} else {
														}
														return
													}
												}
											}
										}
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if int32(0) < v146 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return
											} else {
												v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_DropRelationFiles(m, v151, v152, int32(1))
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v156 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v161, v162)
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return
															} else {
																v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v165 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return
																	} else {
																		v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v171 = v170
																		if v171&int32(536870912) != 0 {
																			v175 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																		} else {
																		}
																		return
																	}
																} else {
																	v171 = v165
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															}
														}
													} else {
														v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v165 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v171 = v170
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														} else {
															v171 = v165
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													}
												}
											}
										} else {
											v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v156 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v161, v162)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v165 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v171 = v170
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														} else {
															v171 = v165
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													}
												}
											} else {
												v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v165 < int32(0) {
													F_XLogFlush(m, l2)
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v171 = v170
														if v171&int32(536870912) != 0 {
															v175 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
														} else {
														}
														return
													}
												} else {
													v171 = v165
													if v171&int32(536870912) != 0 {
														v175 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
													} else {
													}
													return
												}
											}
										}
									}
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									F_StandbyReleaseLockTree(m, l1, v134, v135)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
										if v138&int32(32) != 0 {
											v141 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
											v142 = int32(0)
											F_replorigin_advance(m, l3, v141, l2, v142, v142)
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return
											} else {
												v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if int32(0) < v146 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return
													} else {
														v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														F_DropRelationFiles(m, v151, v152, int32(1))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if int32(0) < v156 {
																F_XLogFlush(m, l2)
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return
																} else {
																	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																	F_pgstat_execute_transactional_drops(m, v161, v162)
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return
																	} else {
																		v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		if v165 < int32(0) {
																			F_XLogFlush(m, l2)
																			mBase = m.M
																			v169 = m.ExcPending
																			if v169 != 0 {
																				return
																			} else {
																				v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				v171 = v170
																				if v171&int32(536870912) != 0 {
																					v175 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																				} else {
																				}
																				return
																			}
																		} else {
																			v171 = v165
																			if v171&int32(536870912) != 0 {
																				v175 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																			} else {
																			}
																			return
																		}
																	}
																}
															} else {
																v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v165 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return
																	} else {
																		v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v171 = v170
																		if v171&int32(536870912) != 0 {
																			v175 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																		} else {
																		}
																		return
																	}
																} else {
																	v171 = v165
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															}
														}
													}
												} else {
													v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v156 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v161, v162)
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return
															} else {
																v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v165 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return
																	} else {
																		v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v171 = v170
																		if v171&int32(536870912) != 0 {
																			v175 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																		} else {
																		}
																		return
																	}
																} else {
																	v171 = v165
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															}
														}
													} else {
														v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v165 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v171 = v170
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														} else {
															v171 = v165
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													}
												}
											}
										} else {
											v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if int32(0) < v146 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_DropRelationFiles(m, v151, v152, int32(1))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if int32(0) < v156 {
															F_XLogFlush(m, l2)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																F_pgstat_execute_transactional_drops(m, v161, v162)
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return
																} else {
																	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	if v165 < int32(0) {
																		F_XLogFlush(m, l2)
																		mBase = m.M
																		v169 = m.ExcPending
																		if v169 != 0 {
																			return
																		} else {
																			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			v171 = v170
																			if v171&int32(536870912) != 0 {
																				v175 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																			} else {
																			}
																			return
																		}
																	} else {
																		v171 = v165
																		if v171&int32(536870912) != 0 {
																			v175 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																		} else {
																		}
																		return
																	}
																}
															}
														} else {
															v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v165 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v171 = v170
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															} else {
																v171 = v165
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														}
													}
												}
											} else {
												v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v156 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v161, v162)
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v165 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return
																} else {
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v171 = v170
																	if v171&int32(536870912) != 0 {
																		v175 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																	} else {
																	}
																	return
																}
															} else {
																v171 = v165
																if v171&int32(536870912) != 0 {
																	v175 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
																} else {
																}
																return
															}
														}
													}
												} else {
													v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v165 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v171 = v170
															if v171&int32(536870912) != 0 {
																v175 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
															} else {
															}
															return
														}
													} else {
														v171 = v165
														if v171&int32(536870912) != 0 {
															v175 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[186])) = uint8(v175)
														} else {
														}
														return
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
