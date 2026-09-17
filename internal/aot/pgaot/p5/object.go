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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
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
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(80))
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
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_0))
					mBase = m.M
					v268 = m.ExcPending
					if v268 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v331 = m.ExcPending
					if v331 != 0 {
						return int32(0)
					} else {
						v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
						F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
							mBase = m.M
							v341 = m.ExcPending
							if v341 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 34:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_5))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 42:
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v93 = F_SearchSysCache1(m, int32(47), v92)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						if v93 != 0 {
							v347 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
							v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+22)))
							v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v348)+96)))
							if v350 == int32(112) {
								v353 = int32(_a_F_getObjectTypeDescription_6)
							} else {
								v353 = int32(_a_F_getObjectTypeDescription_7)
							}
							if v350 == int32(97) {
								v356 = int32(_a_F_getObjectTypeDescription_8)
							} else {
								v356 = v353
							}
							F_appendStringInfoString(m, v8+int32(80), v356)
							mBase = m.M
							v358 = m.ExcPending
							if v358 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v93)
								mBase = m.M
								v360 = m.ExcPending
								if v360 != 0 {
									return int32(0)
								} else {
									v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v431
								}
							}
						} else {
							if l1 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v379 = m.ExcPending
								if v379 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v92
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_9), v8+int32(32))
									mBase = m.M
									v385 = m.ExcPending
									if v385 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_10), int32(_a_F_getObjectTypeDescription_11))
										mBase = m.M
										v390 = m.ExcPending
										if v390 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_12))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v431
								}
							}
						}
					}
				case 46:
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v55 = F_SearchSysCache1(m, int32(57), v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 == int32(0) {
							if l1 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v364 = m.ExcPending
								if v364 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_13), v8+int32(16))
									mBase = m.M
									v370 = m.ExcPending
									if v370 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_14), int32(_a_F_getObjectTypeDescription_15))
										mBase = m.M
										v375 = m.ExcPending
										if v375 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_16))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v431
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
							switch v70 - int32(73) {
							case 0, 32:
								v81 = int32(_a_F_getObjectTypeDescription_17)
							default:
								v81 = int32(_a_F_getObjectTypeDescription_16)
							case 10:
								v81 = int32(_a_F_getObjectTypeDescription_18)
							case 26:
								v81 = int32(_a_F_getObjectTypeDescription_19)
							case 29:
								v81 = int32(_a_F_getObjectTypeDescription_20)
							case 36:
								v81 = int32(_a_F_getObjectTypeDescription_21)
							case 39, 41:
								v81 = int32(_a_F_getObjectTypeDescription_22)
							case 43:
								v81 = int32(_a_F_getObjectTypeDescription_23)
							case 45:
								v81 = int32(_a_F_getObjectTypeDescription_24)
							}
							v83 = v8 + int32(80)
							F_appendStringInfoString(m, v83, v81)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									F_appendStringInfoString(m, v83, int32(_a_F_getObjectTypeDescription_25))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v431
										}
									}
								} else {
									F_ReleaseCatCache(m, v55)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
										m.G0 = v8 + int32(96)
										return v431
									}
								}
							}
						}
					}
				case 47:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_26))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 48:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_27))
					mBase = m.M
					v258 = m.ExcPending
					if v258 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 49:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_28))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				default:
					switch v16 - int32(1417) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_29))
						mBase = m.M
						v280 = m.ExcPending
						if v280 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 1:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_30))
						mBase = m.M
						v285 = m.ExcPending
						if v285 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					default:
						if v16 == int32(826) {
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_31))
							mBase = m.M
							v426 = m.ExcPending
							if v426 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v331 = m.ExcPending
							if v331 != 0 {
								return int32(0)
							} else {
								v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v341 = m.ExcPending
									if v341 != 0 {
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
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_32))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 1:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_33))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 2:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_34))
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 3:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_35))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 4:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_36))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 5:
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v120 = F_table_open(m, int32(2606), int32(1))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						v124 = F_get_catalog_object_by_oid_extended(m, v120, int32(1), v117, int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							if v124 == int32(0) {
								if l1 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v394 = m.ExcPending
									if v394 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v117
										F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_37), v8+int32(48))
										mBase = m.M
										v400 = m.ExcPending
										if v400 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_38), int32(_a_F_getObjectTypeDescription_39))
											mBase = m.M
											v405 = m.ExcPending
											if v405 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_relation_close(m, v120, int32(1))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_40))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v431
										}
									}
								}
							} else {
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+22)))
								v142 = v140 + v141
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+80))
								if v143 != 0 {
									v149 = int32(_a_F_getObjectTypeDescription_41)
									F_appendStringInfoString(m, v8+int32(80), v149)
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v120, int32(1))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v431
										}
									}
								} else {
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+84))
									if v145 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v409 = m.ExcPending
										if v409 != 0 {
											return int32(0)
										} else {
											v410 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v410
											F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_42), v8-int32(-64))
											mBase = m.M
											v416 = m.ExcPending
											if v416 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_43), int32(_a_F_getObjectTypeDescription_39))
												mBase = m.M
												v421 = m.ExcPending
												if v421 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v149 = int32(_a_F_getObjectTypeDescription_44)
										F_appendStringInfoString(m, v8+int32(80), v149)
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v120, int32(1))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
												m.G0 = v8 + int32(96)
												return v431
											}
										}
									}
								}
							}
						}
					}
				case 6:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_45))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 7, 8, 9, 10, 13, 18:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v331 = m.ExcPending
					if v331 != 0 {
						return int32(0)
					} else {
						v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
						F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
							mBase = m.M
							v341 = m.ExcPending
							if v341 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 11:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_46))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 12:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_47))
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 14:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_48))
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 15:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_49))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 16:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_50))
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 17:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_51))
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				case 19:
					F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_52))
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return int32(0)
					} else {
						v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v431
					}
				default:
					if v16 != int32(2328) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v331 = m.ExcPending
						if v331 != 0 {
							return int32(0)
						} else {
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_53))
						mBase = m.M
						v275 = m.ExcPending
						if v275 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					}
				}
			}
		} else {
			if v16 <= int32(3575) {
				if v16 <= int32(3380) {
					if v16 == int32(2753) {
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_54))
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					} else {
						if v16 == int32(3079) {
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_55))
							mBase = m.M
							v290 = m.ExcPending
							if v290 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
							}
						} else {
							if v16 != int32(3256) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v331 = m.ExcPending
								if v331 != 0 {
									return int32(0)
								} else {
									v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
									F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
									mBase = m.M
									v336 = m.ExcPending
									if v336 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
										mBase = m.M
										v341 = m.ExcPending
										if v341 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_56))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v431
								}
							}
						}
					}
				} else {
					switch v16 - int32(3456) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_57))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v331 = m.ExcPending
						if v331 != 0 {
							return int32(0)
						} else {
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 10:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_58))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					default:
						if v16 != int32(3381) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v331 = m.ExcPending
							if v331 != 0 {
								return int32(0)
							} else {
								v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v341 = m.ExcPending
									if v341 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_59))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
							}
						}
					}
				}
			} else {
				if v16 <= int32(_a_F_getObjectTypeDescription_60) {
					switch v16 - int32(3576) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_61))
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v331 = m.ExcPending
						if v331 != 0 {
							return int32(0)
						} else {
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 24:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_62))
						mBase = m.M
						v236 = m.ExcPending
						if v236 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 25:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_63))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 26:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_64))
						mBase = m.M
						v248 = m.ExcPending
						if v248 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					default:
						if v16 != int32(3764) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v331 = m.ExcPending
							if v331 != 0 {
								return int32(0)
							} else {
								v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v341 = m.ExcPending
									if v341 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_65))
							mBase = m.M
							v243 = m.ExcPending
							if v243 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
							}
						}
					}
				} else {
					switch v16 - int32(_a_F_getObjectTypeDescription_66) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_67))
						mBase = m.M
						v320 = m.ExcPending
						if v320 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 1, 2, 3, 5:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v331 = m.ExcPending
						if v331 != 0 {
							return int32(0)
						} else {
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
							F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 4:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_68))
						mBase = m.M
						v305 = m.ExcPending
						if v305 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					case 6:
						F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_69))
						mBase = m.M
						v315 = m.ExcPending
						if v315 != 0 {
							return int32(0)
						} else {
							v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v431
						}
					default:
						switch v16 - int32(_a_F_getObjectTypeDescription_70) {
						case 0:
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_71))
							mBase = m.M
							v310 = m.ExcPending
							if v310 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v331 = m.ExcPending
							if v331 != 0 {
								return int32(0)
							} else {
								v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v332
								F_errmsg_internal(m, int32(_a_F_getObjectTypeDescription_1), v8)
								mBase = m.M
								v336 = m.ExcPending
								if v336 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_getObjectTypeDescription_2), int32(_a_F_getObjectTypeDescription_3), int32(_a_F_getObjectTypeDescription_4))
									mBase = m.M
									v341 = m.ExcPending
									if v341 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							F_appendStringInfoString(m, v8+int32(80), int32(_a_F_getObjectTypeDescription_72))
							mBase = m.M
							v300 = m.ExcPending
							if v300 != 0 {
								return int32(0)
							} else {
								v431 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v431
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
	F_errfinish(m, int32(_a_F_get_object_attnum_acl_5), int32(2777), int32(_a_F_get_object_attnum_acl_6))
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
