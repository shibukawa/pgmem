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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int64
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int64
	_ = v252
	var v253 int64
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int64
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int64
	_ = v361
	var v362 int64
	_ = v362
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
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
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int64
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int64
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
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
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v576 = v9 + int32(16)
	v577 = int32(0)
	base.MemoryFill(m, v576, v577, int32(288))
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*int64)(unsafe.Add(mBase, uint32(v576))) = v583
	if v577 <= base.I32_extend8_s(v12) {
		goto L139
	} else {
		goto L140
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L31
	} else {
		goto L135
	}
L4:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[0]))
	if v414 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L5:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v397 = F_LWLockAcquire(m, v393+int32(2304), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L31
	} else {
		goto L79
	}
L6:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v269 = v9 + int32(16)
	v270 = int32(0)
	base.MemoryFill(m, v269, v270, int32(264))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	*(*int64)(unsafe.Add(mBase, uint32(v269))) = v276
	if v270 <= base.I32_extend8_s(v12) {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v160 = v9 + int32(16)
	v161 = int32(0)
	base.MemoryFill(m, v160, v161, int32(264))
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = v167
	if v161 <= base.I32_extend8_s(v12) {
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
	F_PrepareRedoRemove(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v153+int32(2304))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+36))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+56)))
	F_xact_redo_abort(m, v160, v262, v263, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L31
	} else {
		goto L55
	}
L37:
	;
	goto L36
L38:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v172
	if v172&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v178
	v184 = v158 + int32(20)
	goto L41
L40:
	;
	v184 = v158 + int32(12)
	goto L41
L41:
	;
	if v172&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v189 = v184 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = v187
	v195 = v189 + v187<<(uint(int32(2))%32)
	goto L44
L43:
	;
	v195 = v184
	goto L44
L44:
	;
	if v172&int32(4) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v201 = v195 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+32)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v160)+28)) = v199
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v208 = v201 + v204*int32(12)
	goto L47
L46:
	;
	v208 = v195
	goto L47
L47:
	;
	if v172&int32(256) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v214 = int32(4)
	v215 = v208 + v214
	*(*int32)(unsafe.Add(mBase, uint32(v160)+40)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v160)+36)) = v213
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v222 = v215 + v218<<(uint(v214)%32)
	goto L50
L49:
	;
	v222 = v208
	goto L50
L50:
	;
	if v172&int32(16) == int32(0) {
		v246 = v172
		v247 = v222
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v246&int32(32) == int32(0) {
		goto L37
	} else {
		goto L54
	}
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+44)) = v229
	v232 = v222 + int32(4)
	if v172&int32(128) == int32(0) {
		v246 = v172
		v247 = v232
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v240 = F_strlcpy(m, v9+int32(64), v232, int32(200))
	mBase = m.M
	v241 = F_strlen(m, v232)
	mBase = m.M
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	v246 = v245
	v247 = v241 + v232 + int32(1)
	goto L51
L54:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v247)))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v247)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v160)+256)) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v160)+248)) = v252
	goto L37
L55:
	;
	goto L1
L56:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v371 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372)+56)))
	F_xact_redo_abort(m, v269, v370, v371, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L31
	} else {
		goto L75
	}
L57:
	;
	goto L56
L58:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+8)) = v281
	if v281&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+12)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+16)) = v287
	v293 = v267 + int32(20)
	goto L61
L60:
	;
	v293 = v267 + int32(12)
	goto L61
L61:
	;
	if v281&int32(2) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v298 = v293 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v269)+24)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v269)+20)) = v296
	v304 = v298 + v296<<(uint(int32(2))%32)
	goto L64
L63:
	;
	v304 = v293
	goto L64
L64:
	;
	if v281&int32(4) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v310 = v304 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v269)+32)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v269)+28)) = v308
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v317 = v310 + v313*int32(12)
	goto L67
L66:
	;
	v317 = v304
	goto L67
L67:
	;
	if v281&int32(256) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v323 = int32(4)
	v324 = v317 + v323
	*(*int32)(unsafe.Add(mBase, uint32(v269)+40)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v269)+36)) = v322
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v331 = v324 + v327<<(uint(v323)%32)
	goto L70
L69:
	;
	v331 = v317
	goto L70
L70:
	;
	if v281&int32(16) == int32(0) {
		v355 = v281
		v356 = v331
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v355&int32(32) == int32(0) {
		goto L57
	} else {
		goto L74
	}
L72:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	*(*int32)(unsafe.Add(mBase, uint32(v269)+44)) = v338
	v341 = v331 + int32(4)
	if v281&int32(128) == int32(0) {
		v355 = v281
		v356 = v341
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v349 = F_strlcpy(m, v9+int32(64), v341, int32(200))
	mBase = m.M
	v350 = F_strlen(m, v341)
	mBase = m.M
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	v355 = v354
	v356 = v350 + v341 + int32(1)
	goto L71
L74:
	;
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v356)))
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v356)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+256)) = v362
	*(*int64)(unsafe.Add(mBase, uint32(v269)+248)) = v361
	goto L57
L75:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v381 = F_LWLockAcquire(m, v377+int32(2304), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L31
	} else {
		goto L76
	}
L76:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	F_PrepareRedoRemove(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v387+int32(2304))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+64))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400)+56)))
	F_PrepareRedoAdd(m, int64(0), v401, v402, v403, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L31
	} else {
		goto L80
	}
L80:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v408+int32(2304))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L31
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v417 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v422 = v418 + int32(8)
	v426 = v420 - int32(1)
	if v426 < v417 {
		v494 = v419
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_RecordKnownAssignedTransactionIds(m, v494)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L31
	} else {
		goto L114
	}
L84:
	;
	goto L83
L85:
	;
	if v420&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v431 = int32(2)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v422+v426<<(uint(v431)%32))))
	if base.B2i32(base.Ui32(v431) < base.Ui32(v434))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v419)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v449 = v419
	v451 = v426
	goto L88
L88:
	;
	if v426 == int32(0) {
		v494 = v449
		goto L84
	} else {
		goto L96
	}
L89:
	;
	v449 = v446
	v451 = v420 - int32(2)
	goto L88
L90:
	;
	v446 = v434
	goto L89
L91:
	;
	if base.Ui32(v419) < base.Ui32(v434) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if int32(0) <= v419-v434 {
		v446 = v419
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v446 = v419
	goto L89
L95:
	;
	goto L90
L96:
	;
	v454 = v449
	v455 = v451
	goto L97
L97:
	;
	v459 = int32(3)
	v463 = v422 + v455<<(uint(int32(2))%32)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	if base.B2i32(base.Ui32(v454) < base.Ui32(v459))|base.B2i32(base.Ui32(v464) < base.Ui32(v459)) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v494 = v489
	goto L84
L99:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v463-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v477))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v474)) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L100:
	;
	v474 = v464
	goto L99
L101:
	;
	if v454-v464 < int32(0) {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(v464) <= base.Ui32(v454) {
		v474 = v454
		goto L99
	} else {
		goto L105
	}
L104:
	;
	v474 = v454
	goto L99
L105:
	;
	goto L100
L106:
	;
	if int32(1) < v455 {
		v454 = v489
		v455 = v455 - int32(2)
		goto L97
	} else {
		goto L113
	}
L107:
	;
	v489 = v477
	goto L106
L108:
	;
	if base.Ui32(v474) < base.Ui32(v477) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if int32(0) <= v474-v477 {
		v489 = v474
		goto L106
	} else {
		goto L112
	}
L111:
	;
	v489 = v474
	goto L106
L112:
	;
	goto L107
L113:
	;
	goto L98
L114:
	;
	if int32(0) < v420 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v506 = v417
	goto L118
L116:
	;
	goto L117
L117:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[0]))
	if v525 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v422+v506<<(uint(int32(2))%32))))
	F_SubTransSetParent(m, v512, v419)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L31
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	v516 = v506 + int32(1)
	if v516 != v420 {
		v506 = v516
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	v533 = F_LWLockAcquire(m, v529+int32(512), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
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
	F_KnownAssignedXidsRemoveTree(m, int32(0), v420, v422)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L31
	} else {
		goto L126
	}
L126:
	;
	v538 = int32(3)
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[2]))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+24))
	if base.B2i32(base.Ui32(v494) < base.Ui32(v538))|base.B2i32(base.Ui32(v542) < base.Ui32(v538)) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo[1]))
	F_LWLockRelease(m, v554+int32(512))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L31
	} else {
		goto L134
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541)+24)) = v494
	goto L127
L129:
	;
	if v542-v494 < int32(0) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if base.Ui32(v494) <= base.Ui32(v542) {
		goto L127
	} else {
		goto L133
	}
L132:
	;
	goto L127
L133:
	;
	goto L128
L134:
	;
	goto L124
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_errmsg_internal(m, int32(_a_F_xact_redo_0), v9)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L31
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_xact_redo_1), int32(_a_F_xact_redo_2), int32(_a_F_xact_redo_3))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
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
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+36))
	v693 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v691)+56)))
	F_xact_redo_commit(m, v576, v692, v693, v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L31
	} else {
		goto L160
	}
L139:
	;
	goto L138
L140:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+8)) = v588
	if v588&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v574)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+12)) = v592
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v574)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+16)) = v594
	v600 = v574 + int32(20)
	goto L143
L142:
	;
	v600 = v574 + int32(12)
	goto L143
L143:
	;
	if v588&int32(2) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	v605 = v600 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v576)+24)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v576)+20)) = v603
	v611 = v605 + v603<<(uint(int32(2))%32)
	goto L146
L145:
	;
	v611 = v600
	goto L146
L146:
	;
	if v588&int32(4) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v617 = v611 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v576)+32)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v576)+28)) = v615
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v624 = v617 + v620*int32(12)
	goto L149
L148:
	;
	v624 = v611
	goto L149
L149:
	;
	if v588&int32(256) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	v630 = int32(4)
	v631 = v624 + v630
	*(*int32)(unsafe.Add(mBase, uint32(v576)+40)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v576)+36)) = v629
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	v638 = v631 + v634<<(uint(v630)%32)
	goto L152
L151:
	;
	v638 = v624
	goto L152
L152:
	;
	if v588&int32(8) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	v644 = int32(4)
	v645 = v638 + v644
	*(*int32)(unsafe.Add(mBase, uint32(v576)+48)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v576)+44)) = v643
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	v652 = v645 + v648<<(uint(v644)%32)
	goto L155
L154:
	;
	v652 = v638
	goto L155
L155:
	;
	if v588&int32(16) == int32(0) {
		v676 = v588
		v677 = v652
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v676&int32(32) == int32(0) {
		goto L139
	} else {
		goto L159
	}
L157:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+52)) = v659
	v662 = v652 + int32(4)
	if v588&int32(128) == int32(0) {
		v676 = v588
		v677 = v662
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v670 = F_strlcpy(m, v9+int32(72), v662, int32(200))
	mBase = m.M
	v671 = F_strlen(m, v662)
	mBase = m.M
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v576)+8))
	v676 = v675
	v677 = v671 + v662 + int32(1)
	goto L156
L159:
	;
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v677)))
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v677)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v576)+280)) = v683
	*(*int64)(unsafe.Add(mBase, uint32(v576)+272)) = v682
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
