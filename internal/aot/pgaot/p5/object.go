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
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
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
					F_appendStringInfoString(m, v8+int32(80), int32(438300))
					mBase = m.M
					v270 = m.ExcPending
					if v270 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39, 40, 41, 43, 44, 45:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v333 = m.ExcPending
					if v333 != 0 {
						return int32(0)
					} else {
						v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
						F_errmsg_internal(m, int32(62910), v8)
						mBase = m.M
						v338 = m.ExcPending
						if v338 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515466), int32(4674), int32(258685))
							mBase = m.M
							v343 = m.ExcPending
							if v343 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 34:
					F_appendStringInfoString(m, v8+int32(80), int32(388265))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 42:
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v95 = F_SearchSysCache1(m, int32(47), v94)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						if v95 != 0 {
							v349 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
							v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+22)))
							v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+v350)+96)))
							if v352 == int32(112) {
								v355 = int32(380686)
							} else {
								v355 = int32(266209)
							}
							if v352 == int32(97) {
								v358 = int32(371630)
							} else {
								v358 = v355
							}
							F_appendStringInfoString(m, v8+int32(80), v358)
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v95)
								mBase = m.M
								v362 = m.ExcPending
								if v362 != 0 {
									return int32(0)
								} else {
									v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v433
								}
							}
						} else {
							if l1 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v381 = m.ExcPending
								if v381 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v94
									F_errmsg_internal(m, int32(54391), v8+int32(32))
									mBase = m.M
									v387 = m.ExcPending
									if v387 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515466), int32(4798), int32(258737))
										mBase = m.M
										v392 = m.ExcPending
										if v392 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(390129))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v433
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
								v366 = m.ExcPending
								if v366 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
									F_errmsg_internal(m, int32(49777), v8+int32(16))
									mBase = m.M
									v372 = m.ExcPending
									if v372 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515466), int32(4698), int32(258710))
										mBase = m.M
										v377 = m.ExcPending
										if v377 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(275972))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v433
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
							switch v70 - int32(73) {
							case 0, 32:
								v81 = int32(30189)
							default:
								v81 = int32(275972)
							case 10:
								v81 = int32(434826)
							case 26:
								v81 = int32(386844)
							case 29:
								v81 = int32(410764)
							case 36:
								v81 = int32(33692)
							case 39, 41:
								v81 = int32(412760)
							case 43:
								v81 = int32(409834)
							case 45:
								v81 = int32(33762)
							}
							F_appendStringInfoString(m, v8+int32(80), v81)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									F_appendStringInfoString(m, v8+int32(80), int32(287169))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v433
										}
									}
								} else {
									F_ReleaseCatCache(m, v55)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
										m.G0 = v8 + int32(96)
										return v433
									}
								}
							}
						}
					}
				case 47:
					F_appendStringInfoString(m, v8+int32(80), int32(403090))
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 48:
					F_appendStringInfoString(m, v8+int32(80), int32(248760))
					mBase = m.M
					v260 = m.ExcPending
					if v260 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 49:
					F_appendStringInfoString(m, v8+int32(80), int32(379108))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				default:
					switch v16 - int32(1417) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(225158))
						mBase = m.M
						v282 = m.ExcPending
						if v282 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 1:
						F_appendStringInfoString(m, v8+int32(80), int32(349780))
						mBase = m.M
						v287 = m.ExcPending
						if v287 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					default:
						if v16 == int32(826) {
							F_appendStringInfoString(m, v8+int32(80), int32(322164))
							mBase = m.M
							v428 = m.ExcPending
							if v428 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v333 = m.ExcPending
							if v333 != 0 {
								return int32(0)
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
								F_errmsg_internal(m, int32(62910), v8)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515466), int32(4674), int32(258685))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
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
					F_appendStringInfoString(m, v8+int32(80), int32(441666))
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 1:
					F_appendStringInfoString(m, v8+int32(80), int32(441628))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 2:
					F_appendStringInfoString(m, v8+int32(80), int32(441654))
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 3:
					F_appendStringInfoString(m, v8+int32(80), int32(361848))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 4:
					F_appendStringInfoString(m, v8+int32(80), int32(83903))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
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
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v119
										F_errmsg_internal(m, int32(43580), v8+int32(48))
										mBase = m.M
										v402 = m.ExcPending
										if v402 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515466), int32(4762), int32(258656))
											mBase = m.M
											v407 = m.ExcPending
											if v407 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_sequence_close(m, v122, int32(1))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v8+int32(80), int32(96524))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v433
										}
									}
								}
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
								v144 = v142 + v143
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+80))
								if v145 != 0 {
									v151 = int32(96427)
									F_appendStringInfoString(m, v8+int32(80), v151)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										F_sequence_close(m, v122, int32(1))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
											m.G0 = v8 + int32(96)
											return v433
										}
									}
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+84))
									if v147 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v411 = m.ExcPending
										if v411 != 0 {
											return int32(0)
										} else {
											v412 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v412
											F_errmsg_internal(m, int32(43676), v8-int32(-64))
											mBase = m.M
											v418 = m.ExcPending
											if v418 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(515466), int32(4778), int32(258656))
												mBase = m.M
												v423 = m.ExcPending
												if v423 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v151 = int32(95920)
										F_appendStringInfoString(m, v8+int32(80), v151)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_sequence_close(m, v122, int32(1))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
												m.G0 = v8 + int32(96)
												return v433
											}
										}
									}
								}
							}
						}
					}
				case 6:
					F_appendStringInfoString(m, v8+int32(80), int32(283004))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 7, 8, 9, 10, 13, 18:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v333 = m.ExcPending
					if v333 != 0 {
						return int32(0)
					} else {
						v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
						F_errmsg_internal(m, int32(62910), v8)
						mBase = m.M
						v338 = m.ExcPending
						if v338 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515466), int32(4674), int32(258685))
							mBase = m.M
							v343 = m.ExcPending
							if v343 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 11:
					F_appendStringInfoString(m, v8+int32(80), int32(421232))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 12:
					F_appendStringInfoString(m, v8+int32(80), int32(118587))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 14:
					F_appendStringInfoString(m, v8+int32(80), int32(530412))
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 15:
					F_appendStringInfoString(m, v8+int32(80), int32(138812))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 16:
					F_appendStringInfoString(m, v8+int32(80), int32(219682))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 17:
					F_appendStringInfoString(m, v8+int32(80), int32(399918))
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				case 19:
					F_appendStringInfoString(m, v8+int32(80), int32(234877))
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						m.G0 = v8 + int32(96)
						return v433
					}
				default:
					if v16 != int32(2328) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v333 = m.ExcPending
						if v333 != 0 {
							return int32(0)
						} else {
							v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
							F_errmsg_internal(m, int32(62910), v8)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515466), int32(4674), int32(258685))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_appendStringInfoString(m, v8+int32(80), int32(228546))
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					}
				}
			}
		} else {
			if v16 <= int32(3575) {
				if v16 <= int32(3380) {
					if v16 == int32(2753) {
						F_appendStringInfoString(m, v8+int32(80), int32(20193))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					} else {
						if v16 == int32(3079) {
							F_appendStringInfoString(m, v8+int32(80), int32(283744))
							mBase = m.M
							v292 = m.ExcPending
							if v292 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
							}
						} else {
							if v16 != int32(3256) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v333 = m.ExcPending
								if v333 != 0 {
									return int32(0)
								} else {
									v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
									F_errmsg_internal(m, int32(62910), v8)
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515466), int32(4674), int32(258685))
										mBase = m.M
										v343 = m.ExcPending
										if v343 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_appendStringInfoString(m, v8+int32(80), int32(23676))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
									m.G0 = v8 + int32(96)
									return v433
								}
							}
						}
					}
				} else {
					switch v16 - int32(3456) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(274710))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v333 = m.ExcPending
						if v333 != 0 {
							return int32(0)
						} else {
							v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
							F_errmsg_internal(m, int32(62910), v8)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515466), int32(4674), int32(258685))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 10:
						F_appendStringInfoString(m, v8+int32(80), int32(234352))
						mBase = m.M
						v297 = m.ExcPending
						if v297 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					default:
						if v16 != int32(3381) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v333 = m.ExcPending
							if v333 != 0 {
								return int32(0)
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
								F_errmsg_internal(m, int32(62910), v8)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515466), int32(4674), int32(258685))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(80), int32(117828))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
							}
						}
					}
				}
			} else {
				if v16 <= int32(6099) {
					switch v16 - int32(3576) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(300775))
						mBase = m.M
						v327 = m.ExcPending
						if v327 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v333 = m.ExcPending
						if v333 != 0 {
							return int32(0)
						} else {
							v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
							F_errmsg_internal(m, int32(62910), v8)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515466), int32(4674), int32(258685))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 24:
						F_appendStringInfoString(m, v8+int32(80), int32(17790))
						mBase = m.M
						v238 = m.ExcPending
						if v238 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 25:
						F_appendStringInfoString(m, v8+int32(80), int32(228140))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 26:
						F_appendStringInfoString(m, v8+int32(80), int32(271018))
						mBase = m.M
						v250 = m.ExcPending
						if v250 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					default:
						if v16 != int32(3764) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v333 = m.ExcPending
							if v333 != 0 {
								return int32(0)
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
								F_errmsg_internal(m, int32(62910), v8)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515466), int32(4674), int32(258685))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_appendStringInfoString(m, v8+int32(80), int32(371213))
							mBase = m.M
							v245 = m.ExcPending
							if v245 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
							}
						}
					}
				} else {
					switch v16 - int32(6100) {
					case 0:
						F_appendStringInfoString(m, v8+int32(80), int32(259165))
						mBase = m.M
						v322 = m.ExcPending
						if v322 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 1, 2, 3, 5:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v333 = m.ExcPending
						if v333 != 0 {
							return int32(0)
						} else {
							v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
							F_errmsg_internal(m, int32(62910), v8)
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515466), int32(4674), int32(258685))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 4:
						F_appendStringInfoString(m, v8+int32(80), int32(279122))
						mBase = m.M
						v307 = m.ExcPending
						if v307 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					case 6:
						F_appendStringInfoString(m, v8+int32(80), int32(275611))
						mBase = m.M
						v317 = m.ExcPending
						if v317 != 0 {
							return int32(0)
						} else {
							v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							m.G0 = v8 + int32(96)
							return v433
						}
					default:
						switch v16 - int32(6237) {
						case 0:
							F_appendStringInfoString(m, v8+int32(80), int32(437449))
							mBase = m.M
							v312 = m.ExcPending
							if v312 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v333 = m.ExcPending
							if v333 != 0 {
								return int32(0)
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v334
								F_errmsg_internal(m, int32(62910), v8)
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515466), int32(4674), int32(258685))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 6:
							F_appendStringInfoString(m, v8+int32(80), int32(557778))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								m.G0 = v8 + int32(96)
								return v433
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+28)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[253])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[254])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(798656)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[254])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(798656)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[254])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(798656)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[254])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(798656)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(64596), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(515466), int32(2777), int32(527606))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
