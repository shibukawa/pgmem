package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_start_xact_command(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[0])))
	if v4 == int32(0) {
		F_StartTransactionCommand(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[0])) = uint8(v10)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[1]))
			if v27 <= int32(0) {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
				if v48 == int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
					if v56 <= int32(0) {
						return
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
						if v60&int32(1) == int32(0) {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
							if v66 == int32(0) {
								return
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
								if v72 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
									F_enable_timeout_after(m, int32(11), v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
						if v56 <= int32(0) {
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
							if v60&int32(1) == int32(0) {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
								if v66 == int32(0) {
									return
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
									if v72 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
										F_enable_timeout_after(m, int32(11), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[7]))
				if v31 != 0 {
					v34 = base.B2i32(v31 <= v27)
				} else {
					v34 = int32(0)
				}
				if v34 != 0 {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
					if v48 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
						if v56 <= int32(0) {
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
							if v60&int32(1) == int32(0) {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
								if v66 == int32(0) {
									return
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
									if v72 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
										F_enable_timeout_after(m, int32(11), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
							if v56 <= int32(0) {
								return
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
								if v60&int32(1) == int32(0) {
									return
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
									if v66 == int32(0) {
										return
									} else {
										v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
										if v72 != 0 {
											return
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
											F_enable_timeout_after(m, int32(11), v75)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
					if v38 != 0 {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
						if v56 <= int32(0) {
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
							if v60&int32(1) == int32(0) {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
								if v66 == int32(0) {
									return
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
									if v72 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
										F_enable_timeout_after(m, int32(11), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[1]))
						F_enable_timeout_after(m, int32(3), v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
							if v56 <= int32(0) {
								return
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
								if v60&int32(1) == int32(0) {
									return
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
									if v66 == int32(0) {
										return
									} else {
										v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
										if v72 != 0 {
											return
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
											F_enable_timeout_after(m, int32(11), v75)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[8])))
		if v13&int32(8) == int32(0) {
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[9]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
			if v21 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(4)
			} else {
			}
		}
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[1]))
		if v27 <= int32(0) {
			v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
			if v48 == int32(0) {
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
				if v56 <= int32(0) {
					return
				} else {
					v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
					if v60&int32(1) == int32(0) {
						return
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
						if v66 == int32(0) {
							return
						} else {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
							if v72 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
								F_enable_timeout_after(m, int32(11), v75)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
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
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
					if v56 <= int32(0) {
						return
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
						if v60&int32(1) == int32(0) {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
							if v66 == int32(0) {
								return
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
								if v72 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
									F_enable_timeout_after(m, int32(11), v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[7]))
			if v31 != 0 {
				v34 = base.B2i32(v31 <= v27)
			} else {
				v34 = int32(0)
			}
			if v34 != 0 {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
				if v48 == int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
					if v56 <= int32(0) {
						return
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
						if v60&int32(1) == int32(0) {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
							if v66 == int32(0) {
								return
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
								if v72 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
									F_enable_timeout_after(m, int32(11), v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
						if v56 <= int32(0) {
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
							if v60&int32(1) == int32(0) {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
								if v66 == int32(0) {
									return
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
									if v72 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
										F_enable_timeout_after(m, int32(11), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
				v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[2])))
				if v38 != 0 {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
					if v56 <= int32(0) {
						return
					} else {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
						if v60&int32(1) == int32(0) {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
							if v66 == int32(0) {
								return
							} else {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
								if v72 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
									F_enable_timeout_after(m, int32(11), v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[1]))
					F_enable_timeout_after(m, int32(3), v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
						if v56 <= int32(0) {
							return
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[4])))
							if v60&int32(1) == int32(0) {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[5]))
								if v66 == int32(0) {
									return
								} else {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_start_xact_command[6])))
									if v72 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_start_xact_command[3]))
										F_enable_timeout_after(m, int32(11), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(28))+uint32(_c_F_xact_identify[0])))
	return v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int64
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int64
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int64
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int64
	_ = v684
	var v685 int64
	_ = v685
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int64
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
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
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v578 = v9 + int32(16)
	v579 = int32(0)
	base.MemoryFill(m, v578, v579, int32(288))
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v576)))
	*(*int64)(unsafe.Add(mBase, uint32(v578))) = v585
	if v579 <= base.I32_extend8_s(v12) {
		goto L139
	} else {
		goto L140
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L31
	} else {
		goto L135
	}
L4:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[0]))
	if v415 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v399 = F_LWLockAcquire(m, v395+int32(2304), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L31
	} else {
		goto L79
	}
L6:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v270 = v9 + int32(16)
	v271 = int32(0)
	base.MemoryFill(m, v270, v271, int32(264))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v268)))
	*(*int64)(unsafe.Add(mBase, uint32(v270))) = v277
	if v271 <= base.I32_extend8_s(v12) {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v161 = v9 + int32(16)
	v162 = int32(0)
	base.MemoryFill(m, v161, v162, int32(264))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = v168
	if v162 <= base.I32_extend8_s(v12) {
		goto L37
	} else {
		goto L38
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v21 = v9 + int32(16)
	v22 = int32(0)
	base.MemoryFill(m, v21, v22, int32(288))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v28
	if v22 <= base.I32_extend8_s(v12) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+56)))
	F_xact_redo_commit(m, v21, v136, v137, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v33
	if v33&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v39
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v48
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v60
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v74
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v88
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v104
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
	v115 = F_strlcpy(m, v9+int32(72), v107, int32(200))
	mBase = m.M
	v116 = F_strlen(m, v107)
	mBase = m.M
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v121 = v120
	v122 = v116 + v107 + int32(1)
	goto L27
L30:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v122)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+280)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v21)+272)) = v127
	goto L10
L31:
	;
	return
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v147 = F_LWLockAcquire(m, v143+int32(2304), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	F_PrepareRedoRemove(m, v149, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v154+int32(2304))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+36))
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+56)))
	F_xact_redo_abort(m, v161, v263, v264, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L31
	} else {
		goto L55
	}
L37:
	;
	goto L36
L38:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v173
	if v173&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+12)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v179
	v185 = v159 + int32(20)
	goto L41
L40:
	;
	v185 = v159 + int32(12)
	goto L41
L41:
	;
	if v173&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v190 = v185 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = v188
	v196 = v190 + v188<<(uint(int32(2))%32)
	goto L44
L43:
	;
	v196 = v185
	goto L44
L44:
	;
	if v173&int32(4) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v202 = v196 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+32)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v161)+28)) = v200
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v209 = v202 + v205*int32(12)
	goto L47
L46:
	;
	v209 = v196
	goto L47
L47:
	;
	if v173&int32(256) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v215 = int32(4)
	v216 = v209 + v215
	*(*int32)(unsafe.Add(mBase, uint32(v161)+40)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v161)+36)) = v214
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v223 = v216 + v219<<(uint(v215)%32)
	goto L50
L49:
	;
	v223 = v209
	goto L50
L50:
	;
	if v173&int32(16) == int32(0) {
		v247 = v173
		v248 = v223
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v247&int32(32) == int32(0) {
		goto L37
	} else {
		goto L54
	}
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+44)) = v230
	v233 = v223 + int32(4)
	if v173&int32(128) == int32(0) {
		v247 = v173
		v248 = v233
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v241 = F_strlcpy(m, v9+int32(64), v233, int32(200))
	mBase = m.M
	v242 = F_strlen(m, v233)
	mBase = m.M
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v247 = v246
	v248 = v242 + v233 + int32(1)
	goto L51
L54:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v248)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v161)+256)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v161)+248)) = v253
	goto L37
L55:
	;
	goto L1
L56:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v372 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v373)+56)))
	F_xact_redo_abort(m, v270, v371, v372, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L31
	} else {
		goto L75
	}
L57:
	;
	goto L56
L58:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+8)) = v282
	if v282&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+16)) = v288
	v294 = v268 + int32(20)
	goto L61
L60:
	;
	v294 = v268 + int32(12)
	goto L61
L61:
	;
	if v282&int32(2) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v299 = v294 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v270)+24)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v270)+20)) = v297
	v305 = v299 + v297<<(uint(int32(2))%32)
	goto L64
L63:
	;
	v305 = v294
	goto L64
L64:
	;
	if v282&int32(4) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v311 = v305 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v270)+32)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v270)+28)) = v309
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v318 = v311 + v314*int32(12)
	goto L67
L66:
	;
	v318 = v305
	goto L67
L67:
	;
	if v282&int32(256) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v324 = int32(4)
	v325 = v318 + v324
	*(*int32)(unsafe.Add(mBase, uint32(v270)+40)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v270)+36)) = v323
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v332 = v325 + v328<<(uint(v324)%32)
	goto L70
L69:
	;
	v332 = v318
	goto L70
L70:
	;
	if v282&int32(16) == int32(0) {
		v356 = v282
		v357 = v332
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v356&int32(32) == int32(0) {
		goto L57
	} else {
		goto L74
	}
L72:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+44)) = v339
	v342 = v332 + int32(4)
	if v282&int32(128) == int32(0) {
		v356 = v282
		v357 = v342
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v350 = F_strlcpy(m, v9+int32(64), v342, int32(200))
	mBase = m.M
	v351 = F_strlen(m, v342)
	mBase = m.M
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	v356 = v355
	v357 = v351 + v342 + int32(1)
	goto L71
L74:
	;
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v357)))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v357)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v270)+256)) = v363
	*(*int64)(unsafe.Add(mBase, uint32(v270)+248)) = v362
	goto L57
L75:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v382 = F_LWLockAcquire(m, v378+int32(2304), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L31
	} else {
		goto L76
	}
L76:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	F_PrepareRedoRemove(m, v384, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v389+int32(2304))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+56)))
	F_PrepareRedoAdd(m, v402, v403, v404, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v409+int32(2304))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L31
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v418 = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	v423 = v419 + int32(8)
	v427 = v421 - int32(1)
	if v427 < v418 {
		v495 = v420
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_RecordKnownAssignedTransactionIds(m, v495)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L31
	} else {
		goto L114
	}
L84:
	;
	goto L83
L85:
	;
	if v421&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v432 = int32(2)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v423+v427<<(uint(v432)%32))))
	if base.B2i32(base.Ui32(v432) < base.Ui32(v435))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v420)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v450 = v420
	v452 = v427
	goto L88
L88:
	;
	if v427 == int32(0) {
		v495 = v450
		goto L84
	} else {
		goto L96
	}
L89:
	;
	v450 = v447
	v452 = v421 - int32(2)
	goto L88
L90:
	;
	v447 = v435
	goto L89
L91:
	;
	if base.Ui32(v420) < base.Ui32(v435) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if int32(0) <= v420-v435 {
		v447 = v420
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v447 = v420
	goto L89
L95:
	;
	goto L90
L96:
	;
	v455 = v450
	v456 = v452
	goto L97
L97:
	;
	v460 = int32(3)
	v464 = v423 + v456<<(uint(int32(2))%32)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	if base.B2i32(base.Ui32(v455) < base.Ui32(v460))|base.B2i32(base.Ui32(v465) < base.Ui32(v460)) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v495 = v490
	goto L84
L99:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v464-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v478))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v475)) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L100:
	;
	v475 = v465
	goto L99
L101:
	;
	if v455-v465 < int32(0) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(v465) <= base.Ui32(v455) {
		v475 = v455
		goto L99
	} else {
		goto L105
	}
L104:
	;
	v475 = v455
	goto L99
L105:
	;
	goto L100
L106:
	;
	if int32(1) < v456 {
		v455 = v490
		v456 = v456 - int32(2)
		goto L97
	} else {
		goto L113
	}
L107:
	;
	v490 = v478
	goto L106
L108:
	;
	if base.Ui32(v475) < base.Ui32(v478) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if int32(0) <= v475-v478 {
		v490 = v475
		goto L106
	} else {
		goto L112
	}
L111:
	;
	v490 = v475
	goto L106
L112:
	;
	goto L107
L113:
	;
	goto L98
L114:
	;
	if int32(0) < v421 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v504 = v418
	goto L118
L116:
	;
	goto L117
L117:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[0]))
	if v526 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v423+v504<<(uint(int32(2))%32))))
	F_SubTransSetParent(m, v513, v420)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L31
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v517 = v504 + int32(1)
	if v517 != v421 {
		v504 = v517
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v534 = F_LWLockAcquire(m, v530+int32(512), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
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
	F_KnownAssignedXidsRemoveTree(m, int32(0), v421, v423)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L31
	} else {
		goto L126
	}
L126:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[2]))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v495))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v541)) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v553 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v553 = base.B2i32(base.Ui32(v541) < base.Ui32(v495))
	goto L127
L129:
	;
	goto L130
L130:
	;
	v553 = int32(base.Ui32(v541-v495) >> (uint(int32(31)) % 32))
	goto L127
L131:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v555)+24)) = v495
	goto L133
L132:
	;
	goto L133
L133:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v558+int32(512))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_xact_redo_0), v9)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L31
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_xact_redo_1), int32(_a_F_xact_redo_2), int32(_a_F_xact_redo_3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
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
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+36))
	v695 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v693)+56)))
	F_xact_redo_commit(m, v578, v694, v695, v696)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L31
	} else {
		goto L160
	}
L139:
	;
	goto L138
L140:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v576)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+8)) = v590
	if v590&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+12)) = v594
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v576)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+16)) = v596
	v602 = v576 + int32(20)
	goto L143
L142:
	;
	v602 = v576 + int32(12)
	goto L143
L143:
	;
	if v590&int32(2) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v607 = v602 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v578)+24)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v578)+20)) = v605
	v613 = v607 + v605<<(uint(int32(2))%32)
	goto L146
L145:
	;
	v613 = v602
	goto L146
L146:
	;
	if v590&int32(4) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v619 = v613 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v578)+32)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v578)+28)) = v617
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v613)))
	v626 = v619 + v622*int32(12)
	goto L149
L148:
	;
	v626 = v613
	goto L149
L149:
	;
	if v590&int32(256) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v632 = int32(4)
	v633 = v626 + v632
	*(*int32)(unsafe.Add(mBase, uint32(v578)+40)) = v633
	*(*int32)(unsafe.Add(mBase, uint32(v578)+36)) = v631
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v640 = v633 + v636<<(uint(v632)%32)
	goto L152
L151:
	;
	v640 = v626
	goto L152
L152:
	;
	if v590&int32(8) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	v646 = int32(4)
	v647 = v640 + v646
	*(*int32)(unsafe.Add(mBase, uint32(v578)+48)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v578)+44)) = v645
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	v654 = v647 + v650<<(uint(v646)%32)
	goto L155
L154:
	;
	v654 = v640
	goto L155
L155:
	;
	if v590&int32(16) == int32(0) {
		v678 = v590
		v679 = v654
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v678&int32(32) == int32(0) {
		goto L139
	} else {
		goto L159
	}
L157:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+52)) = v661
	v664 = v654 + int32(4)
	if v590&int32(128) == int32(0) {
		v678 = v590
		v679 = v664
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v672 = F_strlcpy(m, v9+int32(72), v664, int32(200))
	mBase = m.M
	v673 = F_strlen(m, v664)
	mBase = m.M
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	v678 = v677
	v679 = v673 + v664 + int32(1)
	goto L156
L159:
	;
	v684 = *(*int64)(unsafe.Add(mBase, uint32(v679)))
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v679)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v578)+280)) = v685
	*(*int64)(unsafe.Add(mBase, uint32(v578)+272)) = v684
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
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
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
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = v6 - int32(1)
	if v11 < int32(0) {
		v79 = l1
	} else {
		if v6&int32(1) != 0 {
			v16 = int32(2)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7+v11<<(uint(v16)%32))))
			if base.B2i32(base.Ui32(v16) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				if base.Ui32(l1) < base.Ui32(v19) {
					v31 = v19
				} else {
					v31 = l1
				}
			} else {
				if int32(0) <= l1-v19 {
					v31 = l1
				} else {
					v31 = v19
				}
			}
			v34 = v31
			v36 = v6 - int32(2)
		} else {
			v34 = l1
			v36 = v11
		}
		if v11 == int32(0) {
			v79 = v34
		} else {
			v39 = v34
			v40 = v36
			for {
				v44 = int32(3)
				v48 = v7 + v40<<(uint(int32(2))%32)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				if base.B2i32(base.Ui32(v39) < base.Ui32(v44))|base.B2i32(base.Ui32(v49) < base.Ui32(v44)) == int32(0) {
					if v39-v49 < int32(0) {
						v59 = v49
					} else {
						v59 = v39
					}
				} else {
					if base.Ui32(v49) <= base.Ui32(v39) {
						v59 = v39
					} else {
						v59 = v49
					}
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(4))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v62))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v59)) == int32(0) {
					if base.Ui32(v59) < base.Ui32(v62) {
						v74 = v62
					} else {
						v74 = v59
					}
				} else {
					if int32(0) <= v59-v62 {
						v74 = v59
					} else {
						v74 = v62
					}
				}
				if int32(1) < v40 {
					v39 = v74
					v40 = v40 - int32(2)
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
	v85 = m.ExcPending
	if v85 != 0 {
		return
	} else {
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v96 = *(*int64)(unsafe.Add(mBase, uint32(l0+v88<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(280))))
		F_TransactionTreeSetCommitTsData(m, l1, v86, v87, v96, l3)
		mBase = m.M
		v98 = m.ExcPending
		if v98 != 0 {
			return
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo_commit[0]))
			if v100 == int32(0) {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_TransactionIdCommitTree(m, l1, v103, v104)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
					if v137&int32(32) != 0 {
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
						v141 = int32(0)
						F_replorigin_advance(m, l3, v140, l2, v141, v141)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return
						} else {
							v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if int32(0) < v145 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return
								} else {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									F_DropRelationFiles(m, v150, v151, int32(1))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return
									} else {
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v155 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return
											} else {
												v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v160, v161)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return
												} else {
													v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v164 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return
														} else {
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v170 = v169
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													} else {
														v170 = v164
														if v170&int32(536870912) != 0 {
															v174 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
														} else {
														}
														return
													}
												}
											}
										} else {
											v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v164 < int32(0) {
												F_XLogFlush(m, l2)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v170 = v169
													if v170&int32(536870912) != 0 {
														v174 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
													} else {
													}
													return
												}
											} else {
												v170 = v164
												if v170&int32(536870912) != 0 {
													v174 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
												} else {
												}
												return
											}
										}
									}
								}
							} else {
								v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if int32(0) < v155 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return
									} else {
										v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										F_pgstat_execute_transactional_drops(m, v160, v161)
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return
										} else {
											v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v164 < int32(0) {
												F_XLogFlush(m, l2)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v170 = v169
													if v170&int32(536870912) != 0 {
														v174 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
													} else {
													}
													return
												}
											} else {
												v170 = v164
												if v170&int32(536870912) != 0 {
													v174 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
												} else {
												}
												return
											}
										}
									}
								} else {
									v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v164 < int32(0) {
										F_XLogFlush(m, l2)
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v170 = v169
											if v170&int32(536870912) != 0 {
												v174 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
											} else {
											}
											return
										}
									} else {
										v170 = v164
										if v170&int32(536870912) != 0 {
											v174 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
										} else {
										}
										return
									}
								}
							}
						}
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if int32(0) < v145 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return
							} else {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_DropRelationFiles(m, v150, v151, int32(1))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return
								} else {
									v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v155 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v160, v161)
											mBase = m.M
											v163 = m.ExcPending
											if v163 != 0 {
												return
											} else {
												v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v164 < int32(0) {
													F_XLogFlush(m, l2)
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return
													} else {
														v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v170 = v169
														if v170&int32(536870912) != 0 {
															v174 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
														} else {
														}
														return
													}
												} else {
													v170 = v164
													if v170&int32(536870912) != 0 {
														v174 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
													} else {
													}
													return
												}
											}
										}
									} else {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v164 < int32(0) {
											F_XLogFlush(m, l2)
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v170 = v169
												if v170&int32(536870912) != 0 {
													v174 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
												} else {
												}
												return
											}
										} else {
											v170 = v164
											if v170&int32(536870912) != 0 {
												v174 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
											} else {
											}
											return
										}
									}
								}
							}
						} else {
							v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if int32(0) < v155 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return
								} else {
									v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_pgstat_execute_transactional_drops(m, v160, v161)
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return
									} else {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v164 < int32(0) {
											F_XLogFlush(m, l2)
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v170 = v169
												if v170&int32(536870912) != 0 {
													v174 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
												} else {
												}
												return
											}
										} else {
											v170 = v164
											if v170&int32(536870912) != 0 {
												v174 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
											} else {
											}
											return
										}
									}
								}
							} else {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v164 < int32(0) {
									F_XLogFlush(m, l2)
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v170 = v169
										if v170&int32(536870912) != 0 {
											v174 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
										} else {
										}
										return
									}
								} else {
									v170 = v164
									if v170&int32(536870912) != 0 {
										v174 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
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
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_TransactionIdAsyncCommitTree(m, l1, v109, v110, l2)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_ExpireTreeKnownAssignedTransactionIds(m, l1, v113, v114, v79)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_ProcessCommittedInvalidationMessages(m, v117, v118, int32(base.Ui32(v119&int32(1073741824))>>(uint(int32(30))%32)), v124, v125)
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return
							} else {
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v128&int32(64) == int32(0) {
									v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
									if v137&int32(32) != 0 {
										v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
										v141 = int32(0)
										F_replorigin_advance(m, l3, v140, l2, v141, v141)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if int32(0) < v145 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_DropRelationFiles(m, v150, v151, int32(1))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if int32(0) < v155 {
															F_XLogFlush(m, l2)
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return
															} else {
																v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																F_pgstat_execute_transactional_drops(m, v160, v161)
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return
																} else {
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	if v164 < int32(0) {
																		F_XLogFlush(m, l2)
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			v170 = v169
																			if v170&int32(536870912) != 0 {
																				v174 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																			} else {
																			}
																			return
																		}
																	} else {
																		v170 = v164
																		if v170&int32(536870912) != 0 {
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																		} else {
																		}
																		return
																	}
																}
															}
														} else {
															v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v164 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v170 = v169
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															} else {
																v170 = v164
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														}
													}
												}
											} else {
												v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v155 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return
													} else {
														v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v160, v161)
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return
														} else {
															v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v164 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v170 = v169
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															} else {
																v170 = v164
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														}
													}
												} else {
													v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v164 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return
														} else {
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v170 = v169
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													} else {
														v170 = v164
														if v170&int32(536870912) != 0 {
															v174 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
														} else {
														}
														return
													}
												}
											}
										}
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if int32(0) < v145 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_DropRelationFiles(m, v150, v151, int32(1))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v155 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return
														} else {
															v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v160, v161)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return
															} else {
																v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v164 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v170 = v169
																		if v170&int32(536870912) != 0 {
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																		} else {
																		}
																		return
																	}
																} else {
																	v170 = v164
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															}
														}
													} else {
														v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v164 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v170 = v169
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														} else {
															v170 = v164
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													}
												}
											}
										} else {
											v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v155 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return
												} else {
													v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v160, v161)
													mBase = m.M
													v163 = m.ExcPending
													if v163 != 0 {
														return
													} else {
														v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v164 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v170 = v169
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														} else {
															v170 = v164
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													}
												}
											} else {
												v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v164 < int32(0) {
													F_XLogFlush(m, l2)
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return
													} else {
														v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v170 = v169
														if v170&int32(536870912) != 0 {
															v174 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
														} else {
														}
														return
													}
												} else {
													v170 = v164
													if v170&int32(536870912) != 0 {
														v174 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
													} else {
													}
													return
												}
											}
										}
									}
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									F_StandbyReleaseLockTree(m, l1, v133, v134)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
										if v137&int32(32) != 0 {
											v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+272))
											v141 = int32(0)
											F_replorigin_advance(m, l3, v140, l2, v141, v141)
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return
											} else {
												v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if int32(0) < v145 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return
													} else {
														v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														F_DropRelationFiles(m, v150, v151, int32(1))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if int32(0) < v155 {
																F_XLogFlush(m, l2)
																mBase = m.M
																v159 = m.ExcPending
																if v159 != 0 {
																	return
																} else {
																	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																	F_pgstat_execute_transactional_drops(m, v160, v161)
																	mBase = m.M
																	v163 = m.ExcPending
																	if v163 != 0 {
																		return
																	} else {
																		v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		if v164 < int32(0) {
																			F_XLogFlush(m, l2)
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return
																			} else {
																				v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				v170 = v169
																				if v170&int32(536870912) != 0 {
																					v174 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																				} else {
																				}
																				return
																			}
																		} else {
																			v170 = v164
																			if v170&int32(536870912) != 0 {
																				v174 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																			} else {
																			}
																			return
																		}
																	}
																}
															} else {
																v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v164 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v170 = v169
																		if v170&int32(536870912) != 0 {
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																		} else {
																		}
																		return
																	}
																} else {
																	v170 = v164
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															}
														}
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v155 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return
														} else {
															v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v160, v161)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return
															} else {
																v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																if v164 < int32(0) {
																	F_XLogFlush(m, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		v170 = v169
																		if v170&int32(536870912) != 0 {
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																		} else {
																		}
																		return
																	}
																} else {
																	v170 = v164
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															}
														}
													} else {
														v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														if v164 < int32(0) {
															F_XLogFlush(m, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return
															} else {
																v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v170 = v169
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														} else {
															v170 = v164
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													}
												}
											}
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if int32(0) < v145 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_DropRelationFiles(m, v150, v151, int32(1))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if int32(0) < v155 {
															F_XLogFlush(m, l2)
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return
															} else {
																v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																F_pgstat_execute_transactional_drops(m, v160, v161)
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return
																} else {
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	if v164 < int32(0) {
																		F_XLogFlush(m, l2)
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			v170 = v169
																			if v170&int32(536870912) != 0 {
																				v174 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																			} else {
																			}
																			return
																		}
																	} else {
																		v170 = v164
																		if v170&int32(536870912) != 0 {
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																		} else {
																		}
																		return
																	}
																}
															}
														} else {
															v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v164 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v170 = v169
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															} else {
																v170 = v164
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														}
													}
												}
											} else {
												v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v155 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return
													} else {
														v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v160, v161)
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return
														} else {
															v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															if v164 < int32(0) {
																F_XLogFlush(m, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return
																} else {
																	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v170 = v169
																	if v170&int32(536870912) != 0 {
																		v174 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																	} else {
																	}
																	return
																}
															} else {
																v170 = v164
																if v170&int32(536870912) != 0 {
																	v174 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
																} else {
																}
																return
															}
														}
													}
												} else {
													v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													if v164 < int32(0) {
														F_XLogFlush(m, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return
														} else {
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v170 = v169
															if v170&int32(536870912) != 0 {
																v174 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
															} else {
															}
															return
														}
													} else {
														v170 = v164
														if v170&int32(536870912) != 0 {
															v174 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_xact_redo_commit[1])) = uint8(v174)
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
