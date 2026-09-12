package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetStartupBufferPinWaitBufId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetStartupBufferPinWaitBufId[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+72))
	return v3
}
func F_SetStartupBufferPinWaitBufId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SetStartupBufferPinWaitBufId[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = l0
	return
}
func F_StartupDecodingContext(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	v4 = l3
	v5 = l4
	v6 = l5
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	v27 = F_AllocSetContextCreateInternal(m, v22, int32(_a_F_StartupDecodingContext_0), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(_a_F_StartupDecodingContext_3)
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v27
		v36 = F_palloc0(m, int32(168))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v27
			if v5 == int32(0) {
				v44 = int32(0)
				v46 = F_load_external_function(m, v20+int32(137), int32(_a_F_StartupDecodingContext_4), v44, v44)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					if v46 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v356 = m.ExcPending
						if v356 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_5), int32(0))
							mBase = m.M
							v360 = m.ExcPending
							if v360 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(743), int32(_a_F_StartupDecodingContext_7))
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
						m.T0[v46].(func(*base.Module, int32))(m, v36+int32(24))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
							if v54 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_8), int32(0))
									mBase = m.M
									v373 = m.ExcPending
									if v373 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(749), int32(_a_F_StartupDecodingContext_7))
										mBase = m.M
										v378 = m.ExcPending
										if v378 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
								if v57 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v382 = m.ExcPending
									if v382 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_9), int32(0))
										mBase = m.M
										v386 = m.ExcPending
										if v386 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(751), int32(_a_F_StartupDecodingContext_7))
											mBase = m.M
											v391 = m.ExcPending
											if v391 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
									if v60 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v395 = m.ExcPending
										if v395 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_10), int32(0))
											mBase = m.M
											v399 = m.ExcPending
											if v399 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(753), int32(_a_F_StartupDecodingContext_7))
												mBase = m.M
												v404 = m.ExcPending
												if v404 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[2]))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
										if base.B2i32(v66 != int32(0)) == int32(0) {
											v72 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
											v76 = F_LWLockAcquire(m, v72+int32(512), int32(0))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
												v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+124)))
												v82 = v80 | int32(16)
												*(*uint8)(unsafe.Add(mBase, uint32(v79)+124)) = uint8(v82)
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[5]))
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
												*(*uint8)(unsafe.Add(mBase, uint32(v86+v87))) = uint8(v82)
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
												F_LWLockRelease(m, v91+int32(512))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v20
													v100 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
													v101 = F_XLogReaderAllocate(m, v100, l6, v36)
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v101
														if v101 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v408 = m.ExcPending
															if v408 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(_a_F_StartupDecodingContext_11))
																mBase = m.M
																v411 = m.ExcPending
																if v411 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
																	mBase = m.M
																	v415 = m.ExcPending
																	if v415 != 0 {
																		return int32(0)
																	} else {
																		F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
																		mBase = m.M
																		v419 = m.ExcPending
																		if v419 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
																			mBase = m.M
																			v424 = m.ExcPending
																			if v424 != 0 {
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
														} else {
															v106 = m.G0
															v108 = v106 - int32(48)
															m.G0 = v108
															v111 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
															v116 = F_AllocSetContextCreateInternal(m, v111, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return int32(0)
															} else {
																v119 = F_MemoryContextAlloc(m, v116, int32(224))
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return int32(0)
																} else {
																	v122 = v108 + int32(40)
																	v123 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123
																	*(*int64)(unsafe.Add(mBase, uint32(v108)+32)) = v123
																	*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v123
																	v130 = v108 + int32(16)
																	*(*int64)(unsafe.Add(mBase, uint32(v130))) = v123
																	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v123
																	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v123
																	*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v116
																	v141 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v119)+124)) = v141
																		v147 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v119)+128)) = v147
																			v151 = int32(_a_F_StartupDecodingContext_1)
																			v154 = F_GenerationContextCreate(m, v116, int32(_a_F_StartupDecodingContext_18), v151, v151, v151)
																			mBase = m.M
																			v155 = m.ExcPending
																			if v155 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+132)) = v154
																				*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(34359738372)
																				v159 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
																				*(*int32)(unsafe.Add(mBase, uint32(v122))) = v159
																				v164 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v108, int32(1064))
																				mBase = m.M
																				v165 = m.ExcPending
																				if v165 != 0 {
																					return int32(0)
																				} else {
																					v166 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v119)+152)) = v166
																					v168 = int64(0)
																					*(*int64)(unsafe.Add(mBase, uint32(v119)+144)) = v168
																					*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = v168
																					*(*int32)(unsafe.Add(mBase, uint32(v119))) = v164
																					v175 = F_pairingheap_allocate(m, int32(1017), v166)
																					mBase = m.M
																					v176 = m.ExcPending
																					if v176 != 0 {
																						return int32(0)
																					} else {
																						v177 = int64(0)
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+160)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+136)) = v177
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+156)) = v175
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+168)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+176)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+184)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+192)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+200)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+208)) = v177
																						*(*int64)(unsafe.Add(mBase, uint32(v119)+216)) = v177
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = int32(0)
																						v199 = v119 + int32(20)
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v199
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v199
																						v203 = v119 + int32(12)
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v203
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v203
																						v207 = v119 + int32(4)
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v207
																						*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v207
																						v211 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																						F_ReorderBufferCleanupSerializedTXNs(m, v211+int32(24))
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v108 + int32(48)
																							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v119
																							v220 = *(*int64)(unsafe.Add(mBase, uint32(v20)+128))
																							v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																							v227 = F_AllocSetContextCreateInternal(m, v222, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																							mBase = m.M
																							v228 = m.ExcPending
																							if v228 != 0 {
																								return int32(0)
																							} else {
																								v229 = int32(_a_F_StartupDecodingContext_3)
																								v230 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																								*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v227
																								v234 = F_palloc0(m, int32(88))
																								mBase = m.M
																								v235 = m.ExcPending
																								if v235 != 0 {
																									return int32(0)
																								} else {
																									*(*int64)(unsafe.Add(mBase, uint32(v234)+64)) = int64(549755813888)
																									*(*int32)(unsafe.Add(mBase, uint32(v234)+56)) = v119
																									*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v227
																									*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(-1)
																									v243 = F_palloc0(m, int32(512))
																									mBase = m.M
																									v244 = m.ExcPending
																									if v244 != 0 {
																										return int32(0)
																									} else {
																										*(*int64)(unsafe.Add(mBase, uint32(v234)+80)) = int64(0)
																										v247 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(v234)+72)) = uint8(v247)
																										*(*int32)(unsafe.Add(mBase, uint32(v234)+76)) = v243
																										*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = l2
																										*(*uint8)(unsafe.Add(mBase, uint32(v234)+37)) = uint8(v6)
																										*(*int64)(unsafe.Add(mBase, uint32(v234)+16)) = l1
																										*(*uint8)(unsafe.Add(mBase, uint32(v234)+36)) = uint8(v4)
																										*(*int64)(unsafe.Add(mBase, uint32(v234)+24)) = v220
																										*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v230
																										*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v234
																										v258 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v258)+112)) = v36
																										v260 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v260)+40)) = int32(994)
																										v263 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v263)+44)) = int32(995)
																										v266 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = int32(996)
																										v269 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v269)+52)) = int32(997)
																										v272 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v272)+56)) = int32(998)
																										v277 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
																										if v277 != 0 {
																											v291 = v247
																										} else {
																											v279 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																											if v279 != 0 {
																												v291 = int32(1)
																											} else {
																												v281 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
																												if v281 != 0 {
																													v291 = int32(1)
																												} else {
																													v283 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
																													if v283 != 0 {
																														v291 = int32(1)
																													} else {
																														v285 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
																														if v285 != 0 {
																															v291 = int32(1)
																														} else {
																															v287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																															if v287 != 0 {
																																v291 = int32(1)
																															} else {
																																v288 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
																																v291 = base.B2i32(v288 != int32(0))
																															}
																														}
																													}
																												}
																											}
																										}
																										*(*uint8)(unsafe.Add(mBase, uint32(v36)+144)) = uint8(v291)
																										v293 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v293)+76)) = int32(999)
																										v296 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v296)+80)) = int32(1000)
																										v299 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v299)+84)) = int32(1001)
																										v302 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v302)+88)) = int32(1002)
																										v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v305)+92)) = int32(1003)
																										v308 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v308)+96)) = int32(1004)
																										v311 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v311)+100)) = int32(1005)
																										v314 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v314)+104)) = int32(1006)
																										v317 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
																										if v317 != 0 {
																											v325 = v247
																										} else {
																											v318 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
																											if v318 != 0 {
																												v325 = v247
																											} else {
																												v319 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
																												if v319 != 0 {
																													v325 = v247
																												} else {
																													v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
																													if v320 != 0 {
																														v325 = v247
																													} else {
																														v321 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
																														if v321 != 0 {
																															v325 = v247
																														} else {
																															v322 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
																															v325 = base.B2i32(v322 != int32(0))
																														}
																													}
																												}
																											}
																										}
																										*(*uint8)(unsafe.Add(mBase, uint32(v36)+145)) = uint8(v325)
																										v327 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v327)+60)) = int32(1007)
																										v330 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = int32(1008)
																										v333 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v333)+68)) = int32(1009)
																										v336 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v336)+72)) = int32(1010)
																										v339 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v339)+108)) = int32(1011)
																										v342 = F_makeStringInfo(m)
																										mBase = m.M
																										v343 = m.ExcPending
																										if v343 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = l9
																											*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = l8
																											*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = l7
																											*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v342
																											*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = l0
																											*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)) = uint8(v5)
																											*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v32
																											return v36
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
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v20
											v100 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
											v101 = F_XLogReaderAllocate(m, v100, l6, v36)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v101
												if v101 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v408 = m.ExcPending
													if v408 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(_a_F_StartupDecodingContext_11))
														mBase = m.M
														v411 = m.ExcPending
														if v411 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
															mBase = m.M
															v415 = m.ExcPending
															if v415 != 0 {
																return int32(0)
															} else {
																F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
																mBase = m.M
																v419 = m.ExcPending
																if v419 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
																	mBase = m.M
																	v424 = m.ExcPending
																	if v424 != 0 {
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
												} else {
													v106 = m.G0
													v108 = v106 - int32(48)
													m.G0 = v108
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
													v116 = F_AllocSetContextCreateInternal(m, v111, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														v119 = F_MemoryContextAlloc(m, v116, int32(224))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v122 = v108 + int32(40)
															v123 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123
															*(*int64)(unsafe.Add(mBase, uint32(v108)+32)) = v123
															*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v123
															v130 = v108 + int32(16)
															*(*int64)(unsafe.Add(mBase, uint32(v130))) = v123
															*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v123
															*(*int64)(unsafe.Add(mBase, uint32(v108))) = v123
															*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v116
															v141 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v119)+124)) = v141
																v147 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v119)+128)) = v147
																	v151 = int32(_a_F_StartupDecodingContext_1)
																	v154 = F_GenerationContextCreate(m, v116, int32(_a_F_StartupDecodingContext_18), v151, v151, v151)
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v119)+132)) = v154
																		*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(34359738372)
																		v159 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
																		*(*int32)(unsafe.Add(mBase, uint32(v122))) = v159
																		v164 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v108, int32(1064))
																		mBase = m.M
																		v165 = m.ExcPending
																		if v165 != 0 {
																			return int32(0)
																		} else {
																			v166 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v119)+152)) = v166
																			v168 = int64(0)
																			*(*int64)(unsafe.Add(mBase, uint32(v119)+144)) = v168
																			*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = v168
																			*(*int32)(unsafe.Add(mBase, uint32(v119))) = v164
																			v175 = F_pairingheap_allocate(m, int32(1017), v166)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return int32(0)
																			} else {
																				v177 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+160)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+136)) = v177
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+156)) = v175
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+168)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+176)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+184)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+192)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+200)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+208)) = v177
																				*(*int64)(unsafe.Add(mBase, uint32(v119)+216)) = v177
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = int32(0)
																				v199 = v119 + int32(20)
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v199
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v199
																				v203 = v119 + int32(12)
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v203
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v203
																				v207 = v119 + int32(4)
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v207
																				*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v207
																				v211 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																				F_ReorderBufferCleanupSerializedTXNs(m, v211+int32(24))
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v108 + int32(48)
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v119
																					v220 = *(*int64)(unsafe.Add(mBase, uint32(v20)+128))
																					v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																					v227 = F_AllocSetContextCreateInternal(m, v222, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																					mBase = m.M
																					v228 = m.ExcPending
																					if v228 != 0 {
																						return int32(0)
																					} else {
																						v229 = int32(_a_F_StartupDecodingContext_3)
																						v230 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																						*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v227
																						v234 = F_palloc0(m, int32(88))
																						mBase = m.M
																						v235 = m.ExcPending
																						if v235 != 0 {
																							return int32(0)
																						} else {
																							*(*int64)(unsafe.Add(mBase, uint32(v234)+64)) = int64(549755813888)
																							*(*int32)(unsafe.Add(mBase, uint32(v234)+56)) = v119
																							*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v227
																							*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(-1)
																							v243 = F_palloc0(m, int32(512))
																							mBase = m.M
																							v244 = m.ExcPending
																							if v244 != 0 {
																								return int32(0)
																							} else {
																								*(*int64)(unsafe.Add(mBase, uint32(v234)+80)) = int64(0)
																								v247 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v234)+72)) = uint8(v247)
																								*(*int32)(unsafe.Add(mBase, uint32(v234)+76)) = v243
																								*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = l2
																								*(*uint8)(unsafe.Add(mBase, uint32(v234)+37)) = uint8(v6)
																								*(*int64)(unsafe.Add(mBase, uint32(v234)+16)) = l1
																								*(*uint8)(unsafe.Add(mBase, uint32(v234)+36)) = uint8(v4)
																								*(*int64)(unsafe.Add(mBase, uint32(v234)+24)) = v220
																								*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v230
																								*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v234
																								v258 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v258)+112)) = v36
																								v260 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v260)+40)) = int32(994)
																								v263 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v263)+44)) = int32(995)
																								v266 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = int32(996)
																								v269 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v269)+52)) = int32(997)
																								v272 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v272)+56)) = int32(998)
																								v277 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
																								if v277 != 0 {
																									v291 = v247
																								} else {
																									v279 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																									if v279 != 0 {
																										v291 = int32(1)
																									} else {
																										v281 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
																										if v281 != 0 {
																											v291 = int32(1)
																										} else {
																											v283 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
																											if v283 != 0 {
																												v291 = int32(1)
																											} else {
																												v285 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
																												if v285 != 0 {
																													v291 = int32(1)
																												} else {
																													v287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																													if v287 != 0 {
																														v291 = int32(1)
																													} else {
																														v288 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
																														v291 = base.B2i32(v288 != int32(0))
																													}
																												}
																											}
																										}
																									}
																								}
																								*(*uint8)(unsafe.Add(mBase, uint32(v36)+144)) = uint8(v291)
																								v293 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v293)+76)) = int32(999)
																								v296 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v296)+80)) = int32(1000)
																								v299 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v299)+84)) = int32(1001)
																								v302 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v302)+88)) = int32(1002)
																								v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v305)+92)) = int32(1003)
																								v308 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v308)+96)) = int32(1004)
																								v311 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v311)+100)) = int32(1005)
																								v314 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v314)+104)) = int32(1006)
																								v317 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
																								if v317 != 0 {
																									v325 = v247
																								} else {
																									v318 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
																									if v318 != 0 {
																										v325 = v247
																									} else {
																										v319 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
																										if v319 != 0 {
																											v325 = v247
																										} else {
																											v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
																											if v320 != 0 {
																												v325 = v247
																											} else {
																												v321 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
																												if v321 != 0 {
																													v325 = v247
																												} else {
																													v322 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
																													v325 = base.B2i32(v322 != int32(0))
																												}
																											}
																										}
																									}
																								}
																								*(*uint8)(unsafe.Add(mBase, uint32(v36)+145)) = uint8(v325)
																								v327 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v327)+60)) = int32(1007)
																								v330 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = int32(1008)
																								v333 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v333)+68)) = int32(1009)
																								v336 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v336)+72)) = int32(1010)
																								v339 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v339)+108)) = int32(1011)
																								v342 = F_makeStringInfo(m)
																								mBase = m.M
																								v343 = m.ExcPending
																								if v343 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = l9
																									*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = l8
																									*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = l7
																									*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v342
																									*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = l0
																									*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)) = uint8(v5)
																									*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v32
																									return v36
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
									}
								}
							}
						}
					}
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[2]))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
				if base.B2i32(v66 != int32(0)) == int32(0) {
					v72 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
					v76 = F_LWLockAcquire(m, v72+int32(512), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+124)))
						v82 = v80 | int32(16)
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+124)) = uint8(v82)
						v85 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[5]))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
						*(*uint8)(unsafe.Add(mBase, uint32(v86+v87))) = uint8(v82)
						v91 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
						F_LWLockRelease(m, v91+int32(512))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v20
							v100 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
							v101 = F_XLogReaderAllocate(m, v100, l6, v36)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v101
								if v101 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v408 = m.ExcPending
									if v408 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_StartupDecodingContext_11))
										mBase = m.M
										v411 = m.ExcPending
										if v411 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
											mBase = m.M
											v415 = m.ExcPending
											if v415 != 0 {
												return int32(0)
											} else {
												F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
												mBase = m.M
												v419 = m.ExcPending
												if v419 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
													mBase = m.M
													v424 = m.ExcPending
													if v424 != 0 {
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
								} else {
									v106 = m.G0
									v108 = v106 - int32(48)
									m.G0 = v108
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
									v116 = F_AllocSetContextCreateInternal(m, v111, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v119 = F_MemoryContextAlloc(m, v116, int32(224))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v122 = v108 + int32(40)
											v123 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123
											*(*int64)(unsafe.Add(mBase, uint32(v108)+32)) = v123
											*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v123
											v130 = v108 + int32(16)
											*(*int64)(unsafe.Add(mBase, uint32(v130))) = v123
											*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v123
											*(*int64)(unsafe.Add(mBase, uint32(v108))) = v123
											*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v116
											v141 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v119)+124)) = v141
												v147 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v119)+128)) = v147
													v151 = int32(_a_F_StartupDecodingContext_1)
													v154 = F_GenerationContextCreate(m, v116, int32(_a_F_StartupDecodingContext_18), v151, v151, v151)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v119)+132)) = v154
														*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(34359738372)
														v159 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
														*(*int32)(unsafe.Add(mBase, uint32(v122))) = v159
														v164 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v108, int32(1064))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															v166 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v119)+152)) = v166
															v168 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v119)+144)) = v168
															*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = v168
															*(*int32)(unsafe.Add(mBase, uint32(v119))) = v164
															v175 = F_pairingheap_allocate(m, int32(1017), v166)
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																v177 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v119)+160)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+136)) = v177
																*(*int32)(unsafe.Add(mBase, uint32(v119)+156)) = v175
																*(*int64)(unsafe.Add(mBase, uint32(v119)+168)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+176)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+184)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+192)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+200)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+208)) = v177
																*(*int64)(unsafe.Add(mBase, uint32(v119)+216)) = v177
																*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = int32(0)
																v199 = v119 + int32(20)
																*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v199
																*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v199
																v203 = v119 + int32(12)
																*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v203
																*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v203
																v207 = v119 + int32(4)
																*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v207
																*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v207
																v211 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																F_ReorderBufferCleanupSerializedTXNs(m, v211+int32(24))
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v108 + int32(48)
																	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v119
																	v220 = *(*int64)(unsafe.Add(mBase, uint32(v20)+128))
																	v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																	v227 = F_AllocSetContextCreateInternal(m, v222, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																	mBase = m.M
																	v228 = m.ExcPending
																	if v228 != 0 {
																		return int32(0)
																	} else {
																		v229 = int32(_a_F_StartupDecodingContext_3)
																		v230 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v227
																		v234 = F_palloc0(m, int32(88))
																		mBase = m.M
																		v235 = m.ExcPending
																		if v235 != 0 {
																			return int32(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v234)+64)) = int64(549755813888)
																			*(*int32)(unsafe.Add(mBase, uint32(v234)+56)) = v119
																			*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v227
																			*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(-1)
																			v243 = F_palloc0(m, int32(512))
																			mBase = m.M
																			v244 = m.ExcPending
																			if v244 != 0 {
																				return int32(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v234)+80)) = int64(0)
																				v247 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v234)+72)) = uint8(v247)
																				*(*int32)(unsafe.Add(mBase, uint32(v234)+76)) = v243
																				*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = l2
																				*(*uint8)(unsafe.Add(mBase, uint32(v234)+37)) = uint8(v6)
																				*(*int64)(unsafe.Add(mBase, uint32(v234)+16)) = l1
																				*(*uint8)(unsafe.Add(mBase, uint32(v234)+36)) = uint8(v4)
																				*(*int64)(unsafe.Add(mBase, uint32(v234)+24)) = v220
																				*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v230
																				*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v234
																				v258 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v258)+112)) = v36
																				v260 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v260)+40)) = int32(994)
																				v263 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v263)+44)) = int32(995)
																				v266 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = int32(996)
																				v269 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v269)+52)) = int32(997)
																				v272 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v272)+56)) = int32(998)
																				v277 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
																				if v277 != 0 {
																					v291 = v247
																				} else {
																					v279 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																					if v279 != 0 {
																						v291 = int32(1)
																					} else {
																						v281 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
																						if v281 != 0 {
																							v291 = int32(1)
																						} else {
																							v283 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
																							if v283 != 0 {
																								v291 = int32(1)
																							} else {
																								v285 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
																								if v285 != 0 {
																									v291 = int32(1)
																								} else {
																									v287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																									if v287 != 0 {
																										v291 = int32(1)
																									} else {
																										v288 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
																										v291 = base.B2i32(v288 != int32(0))
																									}
																								}
																							}
																						}
																					}
																				}
																				*(*uint8)(unsafe.Add(mBase, uint32(v36)+144)) = uint8(v291)
																				v293 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v293)+76)) = int32(999)
																				v296 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v296)+80)) = int32(1000)
																				v299 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v299)+84)) = int32(1001)
																				v302 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v302)+88)) = int32(1002)
																				v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v305)+92)) = int32(1003)
																				v308 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v308)+96)) = int32(1004)
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v311)+100)) = int32(1005)
																				v314 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v314)+104)) = int32(1006)
																				v317 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
																				if v317 != 0 {
																					v325 = v247
																				} else {
																					v318 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
																					if v318 != 0 {
																						v325 = v247
																					} else {
																						v319 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
																						if v319 != 0 {
																							v325 = v247
																						} else {
																							v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
																							if v320 != 0 {
																								v325 = v247
																							} else {
																								v321 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
																								if v321 != 0 {
																									v325 = v247
																								} else {
																									v322 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
																									v325 = base.B2i32(v322 != int32(0))
																								}
																							}
																						}
																					}
																				}
																				*(*uint8)(unsafe.Add(mBase, uint32(v36)+145)) = uint8(v325)
																				v327 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v327)+60)) = int32(1007)
																				v330 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = int32(1008)
																				v333 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v333)+68)) = int32(1009)
																				v336 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v336)+72)) = int32(1010)
																				v339 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v339)+108)) = int32(1011)
																				v342 = F_makeStringInfo(m)
																				mBase = m.M
																				v343 = m.ExcPending
																				if v343 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = l9
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = l8
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = l7
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v342
																					*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = l0
																					*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)) = uint8(v5)
																					*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v32
																					return v36
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
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v20
					v100 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
					v101 = F_XLogReaderAllocate(m, v100, l6, v36)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v101
						if v101 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v408 = m.ExcPending
							if v408 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_StartupDecodingContext_11))
								mBase = m.M
								v411 = m.ExcPending
								if v411 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
									mBase = m.M
									v415 = m.ExcPending
									if v415 != 0 {
										return int32(0)
									} else {
										F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
										mBase = m.M
										v419 = m.ExcPending
										if v419 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
											mBase = m.M
											v424 = m.ExcPending
											if v424 != 0 {
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
						} else {
							v106 = m.G0
							v108 = v106 - int32(48)
							m.G0 = v108
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
							v116 = F_AllocSetContextCreateInternal(m, v111, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v119 = F_MemoryContextAlloc(m, v116, int32(224))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = v108 + int32(40)
									v123 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v122))) = v123
									*(*int64)(unsafe.Add(mBase, uint32(v108)+32)) = v123
									*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = v123
									v130 = v108 + int32(16)
									*(*int64)(unsafe.Add(mBase, uint32(v130))) = v123
									*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v123
									*(*int64)(unsafe.Add(mBase, uint32(v108))) = v123
									*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v116
									v141 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v119)+124)) = v141
										v147 = F_SlabContextCreate(m, v116, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v119)+128)) = v147
											v151 = int32(_a_F_StartupDecodingContext_1)
											v154 = F_GenerationContextCreate(m, v116, int32(_a_F_StartupDecodingContext_18), v151, v151, v151)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v119)+132)) = v154
												*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(34359738372)
												v159 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
												*(*int32)(unsafe.Add(mBase, uint32(v122))) = v159
												v164 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v108, int32(1064))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int32(0)
												} else {
													v166 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v119)+152)) = v166
													v168 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v119)+144)) = v168
													*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = v168
													*(*int32)(unsafe.Add(mBase, uint32(v119))) = v164
													v175 = F_pairingheap_allocate(m, int32(1017), v166)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return int32(0)
													} else {
														v177 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v119)+160)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+136)) = v177
														*(*int32)(unsafe.Add(mBase, uint32(v119)+156)) = v175
														*(*int64)(unsafe.Add(mBase, uint32(v119)+168)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+176)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+184)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+192)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+200)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+208)) = v177
														*(*int64)(unsafe.Add(mBase, uint32(v119)+216)) = v177
														*(*int32)(unsafe.Add(mBase, uint32(v119)+28)) = int32(0)
														v199 = v119 + int32(20)
														*(*int32)(unsafe.Add(mBase, uint32(v119)+24)) = v199
														*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v199
														v203 = v119 + int32(12)
														*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v203
														*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v203
														v207 = v119 + int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v207
														*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v207
														v211 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
														F_ReorderBufferCleanupSerializedTXNs(m, v211+int32(24))
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return int32(0)
														} else {
															m.G0 = v108 + int32(48)
															*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v119
															v220 = *(*int64)(unsafe.Add(mBase, uint32(v20)+128))
															v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
															v227 = F_AllocSetContextCreateInternal(m, v222, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return int32(0)
															} else {
																v229 = int32(_a_F_StartupDecodingContext_3)
																v230 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v227
																v234 = F_palloc0(m, int32(88))
																mBase = m.M
																v235 = m.ExcPending
																if v235 != 0 {
																	return int32(0)
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v234)+64)) = int64(549755813888)
																	*(*int32)(unsafe.Add(mBase, uint32(v234)+56)) = v119
																	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v227
																	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(-1)
																	v243 = F_palloc0(m, int32(512))
																	mBase = m.M
																	v244 = m.ExcPending
																	if v244 != 0 {
																		return int32(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v234)+80)) = int64(0)
																		v247 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v234)+72)) = uint8(v247)
																		*(*int32)(unsafe.Add(mBase, uint32(v234)+76)) = v243
																		*(*int32)(unsafe.Add(mBase, uint32(v234)+32)) = l2
																		*(*uint8)(unsafe.Add(mBase, uint32(v234)+37)) = uint8(v6)
																		*(*int64)(unsafe.Add(mBase, uint32(v234)+16)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v234)+36)) = uint8(v4)
																		*(*int64)(unsafe.Add(mBase, uint32(v234)+24)) = v220
																		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v230
																		*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v234
																		v258 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v258)+112)) = v36
																		v260 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v260)+40)) = int32(994)
																		v263 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v263)+44)) = int32(995)
																		v266 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = int32(996)
																		v269 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v269)+52)) = int32(997)
																		v272 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v272)+56)) = int32(998)
																		v277 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
																		if v277 != 0 {
																			v291 = v247
																		} else {
																			v279 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
																			if v279 != 0 {
																				v291 = int32(1)
																			} else {
																				v281 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
																				if v281 != 0 {
																					v291 = int32(1)
																				} else {
																					v283 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
																					if v283 != 0 {
																						v291 = int32(1)
																					} else {
																						v285 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
																						if v285 != 0 {
																							v291 = int32(1)
																						} else {
																							v287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																							if v287 != 0 {
																								v291 = int32(1)
																							} else {
																								v288 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
																								v291 = base.B2i32(v288 != int32(0))
																							}
																						}
																					}
																				}
																			}
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v36)+144)) = uint8(v291)
																		v293 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v293)+76)) = int32(999)
																		v296 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v296)+80)) = int32(1000)
																		v299 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v299)+84)) = int32(1001)
																		v302 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v302)+88)) = int32(1002)
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v305)+92)) = int32(1003)
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v308)+96)) = int32(1004)
																		v311 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v311)+100)) = int32(1005)
																		v314 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v314)+104)) = int32(1006)
																		v317 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
																		if v317 != 0 {
																			v325 = v247
																		} else {
																			v318 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
																			if v318 != 0 {
																				v325 = v247
																			} else {
																				v319 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
																				if v319 != 0 {
																					v325 = v247
																				} else {
																					v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
																					if v320 != 0 {
																						v325 = v247
																					} else {
																						v321 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
																						if v321 != 0 {
																							v325 = v247
																						} else {
																							v322 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
																							v325 = base.B2i32(v322 != int32(0))
																						}
																					}
																				}
																			}
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v36)+145)) = uint8(v325)
																		v327 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v327)+60)) = int32(1007)
																		v330 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v330)+64)) = int32(1008)
																		v333 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v333)+68)) = int32(1009)
																		v336 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v336)+72)) = int32(1010)
																		v339 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v339)+108)) = int32(1011)
																		v342 = F_makeStringInfo(m)
																		mBase = m.M
																		v343 = m.ExcPending
																		if v343 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = l9
																			*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = l8
																			*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = l7
																			*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v342
																			*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = l0
																			*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)) = uint8(v5)
																			*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v32
																			return v36
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
			}
		}
	}
}
func F_StartupProcExit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcExit[0]))
	if v4 != 0 {
		F_ShutdownRecoveryTransactionEnvironment(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_StartupProcShutdownHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcShutdownHandler[0]))
	if v3 != 0 {
		F_proc_exit(m, int32(1))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_StartupProcShutdownHandler[1])) = int32(1)
		F_WakeupRecovery(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_StartupProcSigHupHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, _c_F_StartupProcSigHupHandler[0])) = int32(1)
	F_WakeupRecovery(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_StartupProcessMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int64
	_ = v161
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int64
	_ = v258
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[0])) = int32(13)
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_on_shmem_exit(m, int32(957), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v13 = int32(958)
			v15 = m.G0
			v17 = v15 - int32(144)
			m.G0 = v17
			switch int32(960) {
			case 0, 2:
				v27 = v13
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[1])) = v13
				v27 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v27
			F_sigemptyset(m, v17+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = int32(268435456)
			v39 = v17 + int32(4)
			if v39 != 0 {
				v50 = F___memcpy(m, int32(_a_F_StartupProcessMain_1), v39, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v17 + int32(144)
			v55 = int32(-2)
			v57 = m.G0
			v59 = v57 - int32(144)
			m.G0 = v59
			switch int32(0) {
			case 0, 2:
				v69 = v55
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[2])) = v55
				v69 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v69
			F_sigemptyset(m, v59+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v59)+136)) = int32(268435456)
			v81 = v59 + int32(4)
			if v81 != 0 {
				v92 = F___memcpy(m, int32(_a_F_StartupProcessMain_2), v81, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v59 + int32(144)
			v97 = int32(959)
			v99 = m.G0
			v101 = v99 - int32(144)
			m.G0 = v101
			switch int32(961) {
			case 0, 2:
				v111 = v97
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[3])) = v97
				v111 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v111
			F_sigemptyset(m, v101+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v101)+136)) = int32(268435456)
			v123 = v101 + int32(4)
			if v123 != 0 {
				v134 = F___memcpy(m, int32(_a_F_StartupProcessMain_3), v123, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v101 + int32(144)
			v138 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[4])) = v138
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[5])) = v138
			v148 = v138
			for {
				v150 = int32(40)
				v151 = v148 * v150
				v154 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[6]))) = uint8(v154)
				*(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[7]))) = v148
				v161 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[8]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[9]))) = v154
				*(*int64)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[10]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[11]))) = v154
				*(*uint8)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_StartupProcessMain[12]))) = uint8(v154)
				v180 = v148 | int32(1)
				v182 = v180 * v150
				*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[6]))) = uint8(v154)
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[7]))) = v180
				*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[8]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[9]))) = v154
				*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[10]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[11]))) = v154
				*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[12]))) = uint8(v154)
				v211 = v148 | int32(2)
				v213 = v211 * v150
				*(*uint8)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[6]))) = uint8(v154)
				*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[7]))) = v211
				*(*int64)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[8]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[9]))) = v154
				*(*int64)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[10]))) = v161
				*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[11]))) = v154
				*(*uint8)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_StartupProcessMain[12]))) = uint8(v154)
				if base.B2i32(v148 == int32(20)) == v154 {
					v246 = v148 | int32(3)
					v248 = v246 * int32(40)
					v251 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[6]))) = uint8(v251)
					*(*int32)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[7]))) = v246
					v258 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[8]))) = v258
					*(*int32)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[9]))) = v251
					*(*int64)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[10]))) = v258
					*(*int32)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[11]))) = v251
					*(*uint8)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_StartupProcessMain[12]))) = uint8(v251)
					v148 = v148 + int32(4)
					continue
				} else {
					break
				}
				break
			}
			v279 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_StartupProcessMain[13])) = uint8(v279)
			F_pqsignal_be(m, int32(14), int32(1788))
			mBase = m.M
			v285 = int32(-2)
			v287 = m.G0
			v289 = v287 - int32(144)
			m.G0 = v289
			switch int32(0) {
			case 0, 2:
				v299 = v285
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[14])) = v285
				v299 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v299
			F_sigemptyset(m, v289+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v289)+136)) = int32(268435456)
			v311 = v289 + int32(4)
			if v311 != 0 {
				v322 = F___memcpy(m, int32(_a_F_StartupProcessMain_4), v311, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v289 + int32(144)
			v327 = int32(917)
			v329 = m.G0
			v331 = v329 - int32(144)
			m.G0 = v331
			switch int32(919) {
			case 0, 2:
				v341 = v327
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[15])) = v327
				v341 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v331)+4)) = v341
			F_sigemptyset(m, v331+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v331)+136)) = int32(268435456)
			v353 = v331 + int32(4)
			if v353 != 0 {
				v364 = F___memcpy(m, int32(_a_F_StartupProcessMain_5), v353, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v331 + int32(144)
			v369 = int32(960)
			v371 = m.G0
			v373 = v371 - int32(144)
			m.G0 = v373
			switch int32(962) {
			case 0, 2:
				v383 = v369
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[16])) = v369
				v383 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = v383
			F_sigemptyset(m, v373+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v373)+136)) = int32(268435456)
			v395 = v373 + int32(4)
			if v395 != 0 {
				v406 = F___memcpy(m, int32(_a_F_StartupProcessMain_6), v395, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v373 + int32(144)
			v411 = int32(0)
			v413 = m.G0
			v415 = v413 - int32(144)
			m.G0 = v415
			switch int32(2) {
			case 0, 2:
				v425 = v411
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[17])) = v411
				v425 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v415)+4)) = v425
			F_sigemptyset(m, v415+int32(8))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v415)+136)) = int32(268435457)
			v437 = v415 + int32(4)
			if v437 != 0 {
				v448 = F___memcpy(m, int32(_a_F_StartupProcessMain_7), v437, int32(140))
				mBase = m.M
			} else {
			}
			m.G0 = v415 + int32(144)
			F_RegisterTimeout(m, int32(4), int32(961))
			mBase = m.M
			v455 = m.ExcPending
			if v455 != 0 {
				return
			} else {
				F_RegisterTimeout(m, int32(5), int32(962))
				mBase = m.M
				v459 = m.ExcPending
				if v459 != 0 {
					return
				} else {
					F_RegisterTimeout(m, int32(6), int32(963))
					mBase = m.M
					v463 = m.ExcPending
					if v463 != 0 {
						return
					} else {
						F_sigprocmask(m, int32(_a_F_StartupProcessMain_8), int32(0))
						mBase = m.M
						v467 = m.ExcPending
						if v467 != 0 {
							return
						} else {
							F_StartupXLOG(m)
							mBase = m.M
							v469 = m.ExcPending
							if v469 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v472 = m.ExcPending
								if v472 != 0 {
									return
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
		}
	}
}
