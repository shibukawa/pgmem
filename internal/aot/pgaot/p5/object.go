package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_getObjectTypeDescription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(96))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v16 <= int32(2752) {
			if v16 <= int32(2327) {
				switch v16 - int32(1213) {
				case 0:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_0))
					mBase = m.M
					v270 = m.ExcPending
					if v270 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v355 = m.ExcPending
					if v355 != 0 {
						return int32(0)
					} else {
						v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
						F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
						mBase = m.M
						v360 = m.ExcPending
						if v360 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
							mBase = m.M
							v365 = m.ExcPending
							if v365 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 34:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_5))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 42:
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v95 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v93))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						if v95 != 0 {
							v371 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
							v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+22)))
							v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v372)+96)))
							if v374 == int32(112) {
								v377 = int32(_a_F_getObjectTypeDescription_6)
							} else {
								v377 = int32(_a_F_getObjectTypeDescription_7)
							}
							if v374 == int32(97) {
								v380 = int32(_a_F_getObjectTypeDescription_8)
							} else {
								v380 = v377
							}
							F_appendStringInfoString(m, v8+int32(96), v380)
							mBase = m.M
							v382 = m.ExcPending
							if v382 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v95)
								mBase = m.M
								v384 = m.ExcPending
								if v384 != 0 {
									return int32(0)
								} else {
									v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
									m.G0 = v8 + int32(112)
									return v470
								}
							}
						} else {
							if l1 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v403 = m.ExcPending
								if v403 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v93
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_9), v8+int32(32))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_10), int32(_a_F_getObjectTypeDescription_11))
										mBase = m.M
										v414 = m.ExcPending
										if v414 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_12))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
									m.G0 = v8 + int32(112)
									return v470
								}
							}
						}
					}
				case 46:
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v54))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						if v56 == int32(0) {
							if l1 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v388 = m.ExcPending
								if v388 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_13), v8+int32(16))
									mBase = m.M
									v394 = m.ExcPending
									if v394 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_14), int32(_a_F_getObjectTypeDescription_15))
										mBase = m.M
										v399 = m.ExcPending
										if v399 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_16))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
									m.G0 = v8 + int32(112)
									return v470
								}
							}
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v69)+119)))
							switch v71 - int32(73) {
							case 0, 32:
								v82 = int32(_a_F_getObjectTypeDescription_17)
							default:
								v82 = int32(_a_F_getObjectTypeDescription_16)
							case 10:
								v82 = int32(_a_F_getObjectTypeDescription_18)
							case 26:
								v82 = int32(_a_F_getObjectTypeDescription_19)
							case 29:
								v82 = int32(_a_F_getObjectTypeDescription_20)
							case 36:
								v82 = int32(_a_F_getObjectTypeDescription_21)
							case 39, 41:
								v82 = int32(_a_F_getObjectTypeDescription_22)
							case 43:
								v82 = int32(_a_F_getObjectTypeDescription_23)
							case 45:
								v82 = int32(_a_F_getObjectTypeDescription_24)
							}
							v84 = v8 + int32(96)
							F_appendStringInfoString(m, v84, v82)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									F_appendStringInfoString(m, v84, int32(_a_F_getObjectTypeDescription_25))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_ReleaseCatCache(m, v56)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
											m.G0 = v8 + int32(112)
											return v470
										}
									}
								} else {
									F_ReleaseCatCache(m, v56)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
										m.G0 = v8 + int32(112)
										return v470
									}
								}
							}
						}
					}
				case 47:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_26))
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 48:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_27))
					mBase = m.M
					v260 = m.ExcPending
					if v260 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 49:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_28))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				default:
					switch v16 - int32(1417) {
					case 0:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_29))
						mBase = m.M
						v282 = m.ExcPending
						if v282 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 1:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_30))
						mBase = m.M
						v287 = m.ExcPending
						if v287 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					default:
						if v16 == int32(826) {
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_31))
							mBase = m.M
							v465 = m.ExcPending
							if v465 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v355 = m.ExcPending
							if v355 != 0 {
								return int32(0)
							} else {
								v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v360 = m.ExcPending
								if v360 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
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
				switch v16 - int32(2601) {
				case 0:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_32))
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 1:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_33))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 2:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_34))
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 3:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_35))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 4:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_36))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 5:
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v122 = F_table_open(m, int32(2606), int32(1))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						v126 = F_get_catalog_object_by_oid_extended(m, v122, int32(1), v119, int32(0))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							if v126 == int32(0) {
								if l1 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v418 = m.ExcPending
									if v418 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v119
										F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_37), v8+int32(48))
										mBase = m.M
										v424 = m.ExcPending
										if v424 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_38), int32(_a_F_getObjectTypeDescription_39))
											mBase = m.M
											v429 = m.ExcPending
											if v429 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_relation_close(m, v122, int32(1))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_40))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
											m.G0 = v8 + int32(112)
											return v470
										}
									}
								}
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
								v144 = v142 + v143
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+80))
								if v145 != 0 {
									v151 = int32(_a_F_getObjectTypeDescription_41)
									F_appendStringInfoString(m, v8+int32(96), v151)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v122, int32(1))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
											m.G0 = v8 + int32(112)
											return v470
										}
									}
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+84))
									if v147 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v433 = m.ExcPending
										if v433 != 0 {
											return int32(0)
										} else {
											v434 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v434
											F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_42), v8-int32(-64))
											mBase = m.M
											v440 = m.ExcPending
											if v440 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_43), int32(_a_F_getObjectTypeDescription_39))
												mBase = m.M
												v445 = m.ExcPending
												if v445 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v151 = int32(_a_F_getObjectTypeDescription_44)
										F_appendStringInfoString(m, v8+int32(96), v151)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v122, int32(1))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
												m.G0 = v8 + int32(112)
												return v470
											}
										}
									}
								}
							}
						}
					}
				case 6:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_45))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 7, 8, 9, 10, 13, 18:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v355 = m.ExcPending
					if v355 != 0 {
						return int32(0)
					} else {
						v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
						F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
						mBase = m.M
						v360 = m.ExcPending
						if v360 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
							mBase = m.M
							v365 = m.ExcPending
							if v365 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 11:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_46))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 12:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_47))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 14:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_48))
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 15:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_49))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 16:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_50))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 17:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_51))
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				case 19:
					F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_52))
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
						m.G0 = v8 + int32(112)
						return v470
					}
				default:
					if v16 != int32(2328) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v355 = m.ExcPending
						if v355 != 0 {
							return int32(0)
						} else {
							v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_53))
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					}
				}
			}
		} else {
			if v16 <= int32(3575) {
				if v16 <= int32(3380) {
					if v16 == int32(2753) {
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_54))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					} else {
						if v16 == int32(3079) {
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_55))
							mBase = m.M
							v292 = m.ExcPending
							if v292 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						} else {
							if v16 != int32(3256) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v355 = m.ExcPending
								if v355 != 0 {
									return int32(0)
								} else {
									v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
									mBase = m.M
									v360 = m.ExcPending
									if v360 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
										mBase = m.M
										v365 = m.ExcPending
										if v365 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_56))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
									m.G0 = v8 + int32(112)
									return v470
								}
							}
						}
					}
				} else {
					switch v16 - int32(3456) {
					case 0:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_57))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v355 = m.ExcPending
						if v355 != 0 {
							return int32(0)
						} else {
							v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 10:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_58))
						mBase = m.M
						v297 = m.ExcPending
						if v297 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					default:
						if v16 != int32(3381) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v355 = m.ExcPending
							if v355 != 0 {
								return int32(0)
							} else {
								v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v360 = m.ExcPending
								if v360 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_59))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						}
					}
				}
			} else {
				if v16 <= int32(_a_F_getObjectTypeDescription_60) {
					switch v16 - int32(3576) {
					case 0:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_61))
						mBase = m.M
						v349 = m.ExcPending
						if v349 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v355 = m.ExcPending
						if v355 != 0 {
							return int32(0)
						} else {
							v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 24:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_62))
						mBase = m.M
						v238 = m.ExcPending
						if v238 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 25:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_63))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 26:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_64))
						mBase = m.M
						v250 = m.ExcPending
						if v250 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					default:
						if v16 != int32(3764) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v355 = m.ExcPending
							if v355 != 0 {
								return int32(0)
							} else {
								v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v360 = m.ExcPending
								if v360 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_65))
							mBase = m.M
							v245 = m.ExcPending
							if v245 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						}
					}
				} else {
					switch v16 - int32(_a_F_getObjectTypeDescription_66) {
					case 0:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_67))
						mBase = m.M
						v344 = m.ExcPending
						if v344 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 1, 2, 3, 5:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v355 = m.ExcPending
						if v355 != 0 {
							return int32(0)
						} else {
							v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 4:
						F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_68))
						mBase = m.M
						v307 = m.ExcPending
						if v307 != 0 {
							return int32(0)
						} else {
							v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
							m.G0 = v8 + int32(112)
							return v470
						}
					case 6:
						v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v316 = F_SearchSysCache1(m, int32(52), base.I64_extend_i32_u(v314))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							if v316 == int32(0) {
								if l1 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v449 = m.ExcPending
									if v449 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v314
										F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_69), v8+int32(80))
										mBase = m.M
										v455 = m.ExcPending
										if v455 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_70), int32(_a_F_getObjectTypeDescription_71))
											mBase = m.M
											v460 = m.ExcPending
											if v460 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_72))
									mBase = m.M
									v326 = m.ExcPending
									if v326 != 0 {
										return int32(0)
									} else {
										v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
										m.G0 = v8 + int32(112)
										return v470
									}
								}
							} else {
								v331 = *(*int32)(unsafe.Add(mBase, uint32(v316)+16))
								v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+22)))
								v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v332)+12)))
								if v334 != 0 {
									v335 = int32(_a_F_getObjectTypeDescription_73)
								} else {
									v335 = int32(_a_F_getObjectTypeDescription_72)
								}
								F_appendStringInfoString(m, v8+int32(96), v335)
								mBase = m.M
								v337 = m.ExcPending
								if v337 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v316)
									mBase = m.M
									v339 = m.ExcPending
									if v339 != 0 {
										return int32(0)
									} else {
										v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
										m.G0 = v8 + int32(112)
										return v470
									}
								}
							}
						}
					default:
						switch v16 - int32(_a_F_getObjectTypeDescription_74) {
						case 0:
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_75))
							mBase = m.M
							v312 = m.ExcPending
							if v312 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v355 = m.ExcPending
							if v355 != 0 {
								return int32(0)
							} else {
								v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v356
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v360 = m.ExcPending
								if v360 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							F_appendStringInfoString(m, v8+int32(96), int32(_a_F_getObjectTypeDescription_76))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								v470 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
								m.G0 = v8 + int32(112)
								return v470
							}
						}
					}
				}
			}
		}
	}
}
func F_get_object_attnum_acl(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_acl[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+28)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_acl[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_acl_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_acl_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_acl_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_acl[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_acl_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_acl[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_acl[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_acl[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_acl_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_acl_5), int32(2827), int32(_a_F_get_object_attnum_acl_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
