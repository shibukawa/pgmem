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
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
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
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	v4 = l3
	v5 = l4
	v6 = l5
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	v25 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_StartupDecodingContext_0), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		v29 = int32(_a_F_StartupDecodingContext_3)
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v25
		v34 = F_palloc0(m, int32(168))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v25
			if v5 == int32(0) {
				v42 = int32(0)
				v44 = F_load_external_function(m, v18+int32(137), int32(_a_F_StartupDecodingContext_4), v42, v42)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					if v44 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v350 = m.ExcPending
						if v350 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_5), int32(0))
							mBase = m.M
							v354 = m.ExcPending
							if v354 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(743), int32(_a_F_StartupDecodingContext_7))
								mBase = m.M
								v359 = m.ExcPending
								if v359 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.T0[v44].(func(*base.Module, int32))(m, v34+int32(24))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
							if v52 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v363 = m.ExcPending
								if v363 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_8), int32(0))
									mBase = m.M
									v367 = m.ExcPending
									if v367 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(749), int32(_a_F_StartupDecodingContext_7))
										mBase = m.M
										v372 = m.ExcPending
										if v372 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
								if v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v376 = m.ExcPending
									if v376 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_9), int32(0))
										mBase = m.M
										v380 = m.ExcPending
										if v380 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(751), int32(_a_F_StartupDecodingContext_7))
											mBase = m.M
											v385 = m.ExcPending
											if v385 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
									if v58 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v389 = m.ExcPending
										if v389 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_10), int32(0))
											mBase = m.M
											v393 = m.ExcPending
											if v393 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(753), int32(_a_F_StartupDecodingContext_7))
												mBase = m.M
												v398 = m.ExcPending
												if v398 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[2]))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
										if base.B2i32(v64 != int32(0)) == int32(0) {
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
											v74 = F_LWLockAcquire(m, v70+int32(512), int32(0))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
												v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+124)))
												v80 = v78 | int32(16)
												*(*uint8)(unsafe.Add(mBase, uint32(v77)+124)) = uint8(v80)
												v83 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[5]))
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
												*(*uint8)(unsafe.Add(mBase, uint32(v84+v85))) = uint8(v80)
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
												F_LWLockRelease(m, v89+int32(512))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v18
													v98 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
													v99 = F_XLogReaderAllocate(m, v98, l6, v34)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v99
														if v99 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v402 = m.ExcPending
															if v402 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(_a_F_StartupDecodingContext_11))
																mBase = m.M
																v405 = m.ExcPending
																if v405 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
																	mBase = m.M
																	v409 = m.ExcPending
																	if v409 != 0 {
																		return int32(0)
																	} else {
																		F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
																		mBase = m.M
																		v413 = m.ExcPending
																		if v413 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
																			mBase = m.M
																			v418 = m.ExcPending
																			if v418 != 0 {
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
															v104 = m.G0
															v106 = v104 - int32(48)
															m.G0 = v106
															v109 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
															v114 = F_AllocSetContextCreateInternal(m, v109, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v117 = F_MemoryContextAlloc(m, v114, int32(224))
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	v119 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v106)+40)) = v119
																	*(*int64)(unsafe.Add(mBase, uint32(v106)+32)) = v119
																	*(*int64)(unsafe.Add(mBase, uint32(v106)+24)) = v119
																	*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = v119
																	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v119
																	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v119
																	*(*int32)(unsafe.Add(mBase, uint32(v117)+120)) = v114
																	v135 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v117)+124)) = v135
																		v141 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v117)+128)) = v141
																			v145 = int32(_a_F_StartupDecodingContext_1)
																			v148 = F_GenerationContextCreate(m, v114, int32(_a_F_StartupDecodingContext_18), v145, v145, v145)
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+132)) = v148
																				*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = int64(34359738372)
																				v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)+120))
																				*(*int32)(unsafe.Add(mBase, uint32(v106)+40)) = v153
																				v158 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v106, int32(1064))
																				mBase = m.M
																				v159 = m.ExcPending
																				if v159 != 0 {
																					return int32(0)
																				} else {
																					v160 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v117)+152)) = v160
																					v162 = int64(0)
																					*(*int64)(unsafe.Add(mBase, uint32(v117)+144)) = v162
																					*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v162
																					*(*int32)(unsafe.Add(mBase, uint32(v117))) = v158
																					v169 = F_pairingheap_allocate(m, int32(1017), v160)
																					mBase = m.M
																					v170 = m.ExcPending
																					if v170 != 0 {
																						return int32(0)
																					} else {
																						v171 = int64(0)
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+160)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+136)) = v171
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+156)) = v169
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+168)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+176)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+184)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+192)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+200)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+208)) = v171
																						*(*int64)(unsafe.Add(mBase, uint32(v117)+216)) = v171
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = int32(0)
																						v193 = v117 + int32(20)
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v193
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v193
																						v197 = v117 + int32(12)
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v197
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v197
																						v201 = v117 + int32(4)
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v201
																						*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v201
																						v205 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																						F_ReorderBufferCleanupSerializedTXNs(m, v205+int32(24))
																						mBase = m.M
																						v209 = m.ExcPending
																						if v209 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v106 + int32(48)
																							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v117
																							v214 = *(*int64)(unsafe.Add(mBase, uint32(v18)+128))
																							v216 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																							v221 = F_AllocSetContextCreateInternal(m, v216, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																							mBase = m.M
																							v222 = m.ExcPending
																							if v222 != 0 {
																								return int32(0)
																							} else {
																								v223 = int32(_a_F_StartupDecodingContext_3)
																								v224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																								*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v221
																								v228 = F_palloc0(m, int32(88))
																								mBase = m.M
																								v229 = m.ExcPending
																								if v229 != 0 {
																									return int32(0)
																								} else {
																									*(*int64)(unsafe.Add(mBase, uint32(v228)+64)) = int64(549755813888)
																									*(*int32)(unsafe.Add(mBase, uint32(v228)+56)) = v117
																									*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v221
																									*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(-1)
																									v237 = F_palloc0(m, int32(512))
																									mBase = m.M
																									v238 = m.ExcPending
																									if v238 != 0 {
																										return int32(0)
																									} else {
																										*(*int64)(unsafe.Add(mBase, uint32(v228)+80)) = int64(0)
																										v241 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(v228)+72)) = uint8(v241)
																										*(*int32)(unsafe.Add(mBase, uint32(v228)+76)) = v237
																										*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = l2
																										*(*uint8)(unsafe.Add(mBase, uint32(v228)+37)) = uint8(v6)
																										*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = l1
																										*(*uint8)(unsafe.Add(mBase, uint32(v228)+36)) = uint8(v4)
																										*(*int64)(unsafe.Add(mBase, uint32(v228)+24)) = v214
																										*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v224
																										*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v228
																										v252 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v252)+112)) = v34
																										v254 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v254)+40)) = int32(994)
																										v257 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v257)+44)) = int32(995)
																										v260 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v260)+48)) = int32(996)
																										v263 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v263)+52)) = int32(997)
																										v266 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v266)+56)) = int32(998)
																										v271 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
																										if v271 != 0 {
																											v285 = v241
																										} else {
																											v273 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
																											if v273 != 0 {
																												v285 = int32(1)
																											} else {
																												v275 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
																												if v275 != 0 {
																													v285 = int32(1)
																												} else {
																													v277 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
																													if v277 != 0 {
																														v285 = int32(1)
																													} else {
																														v279 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
																														if v279 != 0 {
																															v285 = int32(1)
																														} else {
																															v281 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
																															if v281 != 0 {
																																v285 = int32(1)
																															} else {
																																v282 = *(*int32)(unsafe.Add(mBase, uint32(v34)+104))
																																v285 = base.B2i32(v282 != int32(0))
																															}
																														}
																													}
																												}
																											}
																										}
																										*(*uint8)(unsafe.Add(mBase, uint32(v34)+144)) = uint8(v285)
																										v287 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v287)+76)) = int32(999)
																										v290 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v290)+80)) = int32(1000)
																										v293 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v293)+84)) = int32(1001)
																										v296 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v296)+88)) = int32(1002)
																										v299 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v299)+92)) = int32(1003)
																										v302 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v302)+96)) = int32(1004)
																										v305 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v305)+100)) = int32(1005)
																										v308 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v308)+104)) = int32(1006)
																										v311 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
																										if v311 != 0 {
																											v319 = v241
																										} else {
																											v312 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
																											if v312 != 0 {
																												v319 = v241
																											} else {
																												v313 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
																												if v313 != 0 {
																													v319 = v241
																												} else {
																													v314 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
																													if v314 != 0 {
																														v319 = v241
																													} else {
																														v315 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
																														if v315 != 0 {
																															v319 = v241
																														} else {
																															v316 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
																															v319 = base.B2i32(v316 != int32(0))
																														}
																													}
																												}
																											}
																										}
																										*(*uint8)(unsafe.Add(mBase, uint32(v34)+145)) = uint8(v319)
																										v321 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v321)+60)) = int32(1007)
																										v324 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v324)+64)) = int32(1008)
																										v327 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v327)+68)) = int32(1009)
																										v330 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v330)+72)) = int32(1010)
																										v333 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																										*(*int32)(unsafe.Add(mBase, uint32(v333)+108)) = int32(1011)
																										v336 = F_makeStringInfo(m)
																										mBase = m.M
																										v337 = m.ExcPending
																										if v337 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = l9
																											*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = l8
																											*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l7
																											*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v336
																											*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = l0
																											*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)) = uint8(v5)
																											*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v30
																											return v34
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
											*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v18
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
											v99 = F_XLogReaderAllocate(m, v98, l6, v34)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v99
												if v99 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v402 = m.ExcPending
													if v402 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(_a_F_StartupDecodingContext_11))
														mBase = m.M
														v405 = m.ExcPending
														if v405 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
															mBase = m.M
															v409 = m.ExcPending
															if v409 != 0 {
																return int32(0)
															} else {
																F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
																mBase = m.M
																v413 = m.ExcPending
																if v413 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
																	mBase = m.M
																	v418 = m.ExcPending
																	if v418 != 0 {
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
													v104 = m.G0
													v106 = v104 - int32(48)
													m.G0 = v106
													v109 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
													v114 = F_AllocSetContextCreateInternal(m, v109, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = F_MemoryContextAlloc(m, v114, int32(224))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															v119 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v106)+40)) = v119
															*(*int64)(unsafe.Add(mBase, uint32(v106)+32)) = v119
															*(*int64)(unsafe.Add(mBase, uint32(v106)+24)) = v119
															*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = v119
															*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v119
															*(*int64)(unsafe.Add(mBase, uint32(v106))) = v119
															*(*int32)(unsafe.Add(mBase, uint32(v117)+120)) = v114
															v135 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v117)+124)) = v135
																v141 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v117)+128)) = v141
																	v145 = int32(_a_F_StartupDecodingContext_1)
																	v148 = F_GenerationContextCreate(m, v114, int32(_a_F_StartupDecodingContext_18), v145, v145, v145)
																	mBase = m.M
																	v149 = m.ExcPending
																	if v149 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v117)+132)) = v148
																		*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = int64(34359738372)
																		v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)+120))
																		*(*int32)(unsafe.Add(mBase, uint32(v106)+40)) = v153
																		v158 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v106, int32(1064))
																		mBase = m.M
																		v159 = m.ExcPending
																		if v159 != 0 {
																			return int32(0)
																		} else {
																			v160 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v117)+152)) = v160
																			v162 = int64(0)
																			*(*int64)(unsafe.Add(mBase, uint32(v117)+144)) = v162
																			*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v162
																			*(*int32)(unsafe.Add(mBase, uint32(v117))) = v158
																			v169 = F_pairingheap_allocate(m, int32(1017), v160)
																			mBase = m.M
																			v170 = m.ExcPending
																			if v170 != 0 {
																				return int32(0)
																			} else {
																				v171 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+160)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+136)) = v171
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+156)) = v169
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+168)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+176)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+184)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+192)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+200)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+208)) = v171
																				*(*int64)(unsafe.Add(mBase, uint32(v117)+216)) = v171
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = int32(0)
																				v193 = v117 + int32(20)
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v193
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v193
																				v197 = v117 + int32(12)
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v197
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v197
																				v201 = v117 + int32(4)
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v201
																				*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v201
																				v205 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																				F_ReorderBufferCleanupSerializedTXNs(m, v205+int32(24))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v106 + int32(48)
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v117
																					v214 = *(*int64)(unsafe.Add(mBase, uint32(v18)+128))
																					v216 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																					v221 = F_AllocSetContextCreateInternal(m, v216, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																					mBase = m.M
																					v222 = m.ExcPending
																					if v222 != 0 {
																						return int32(0)
																					} else {
																						v223 = int32(_a_F_StartupDecodingContext_3)
																						v224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																						*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v221
																						v228 = F_palloc0(m, int32(88))
																						mBase = m.M
																						v229 = m.ExcPending
																						if v229 != 0 {
																							return int32(0)
																						} else {
																							*(*int64)(unsafe.Add(mBase, uint32(v228)+64)) = int64(549755813888)
																							*(*int32)(unsafe.Add(mBase, uint32(v228)+56)) = v117
																							*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v221
																							*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(-1)
																							v237 = F_palloc0(m, int32(512))
																							mBase = m.M
																							v238 = m.ExcPending
																							if v238 != 0 {
																								return int32(0)
																							} else {
																								*(*int64)(unsafe.Add(mBase, uint32(v228)+80)) = int64(0)
																								v241 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v228)+72)) = uint8(v241)
																								*(*int32)(unsafe.Add(mBase, uint32(v228)+76)) = v237
																								*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = l2
																								*(*uint8)(unsafe.Add(mBase, uint32(v228)+37)) = uint8(v6)
																								*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = l1
																								*(*uint8)(unsafe.Add(mBase, uint32(v228)+36)) = uint8(v4)
																								*(*int64)(unsafe.Add(mBase, uint32(v228)+24)) = v214
																								*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v224
																								*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v228
																								v252 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v252)+112)) = v34
																								v254 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v254)+40)) = int32(994)
																								v257 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v257)+44)) = int32(995)
																								v260 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v260)+48)) = int32(996)
																								v263 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v263)+52)) = int32(997)
																								v266 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v266)+56)) = int32(998)
																								v271 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
																								if v271 != 0 {
																									v285 = v241
																								} else {
																									v273 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
																									if v273 != 0 {
																										v285 = int32(1)
																									} else {
																										v275 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
																										if v275 != 0 {
																											v285 = int32(1)
																										} else {
																											v277 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
																											if v277 != 0 {
																												v285 = int32(1)
																											} else {
																												v279 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
																												if v279 != 0 {
																													v285 = int32(1)
																												} else {
																													v281 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
																													if v281 != 0 {
																														v285 = int32(1)
																													} else {
																														v282 = *(*int32)(unsafe.Add(mBase, uint32(v34)+104))
																														v285 = base.B2i32(v282 != int32(0))
																													}
																												}
																											}
																										}
																									}
																								}
																								*(*uint8)(unsafe.Add(mBase, uint32(v34)+144)) = uint8(v285)
																								v287 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v287)+76)) = int32(999)
																								v290 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v290)+80)) = int32(1000)
																								v293 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v293)+84)) = int32(1001)
																								v296 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v296)+88)) = int32(1002)
																								v299 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v299)+92)) = int32(1003)
																								v302 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v302)+96)) = int32(1004)
																								v305 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v305)+100)) = int32(1005)
																								v308 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v308)+104)) = int32(1006)
																								v311 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
																								if v311 != 0 {
																									v319 = v241
																								} else {
																									v312 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
																									if v312 != 0 {
																										v319 = v241
																									} else {
																										v313 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
																										if v313 != 0 {
																											v319 = v241
																										} else {
																											v314 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
																											if v314 != 0 {
																												v319 = v241
																											} else {
																												v315 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
																												if v315 != 0 {
																													v319 = v241
																												} else {
																													v316 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
																													v319 = base.B2i32(v316 != int32(0))
																												}
																											}
																										}
																									}
																								}
																								*(*uint8)(unsafe.Add(mBase, uint32(v34)+145)) = uint8(v319)
																								v321 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v321)+60)) = int32(1007)
																								v324 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v324)+64)) = int32(1008)
																								v327 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v327)+68)) = int32(1009)
																								v330 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v330)+72)) = int32(1010)
																								v333 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																								*(*int32)(unsafe.Add(mBase, uint32(v333)+108)) = int32(1011)
																								v336 = F_makeStringInfo(m)
																								mBase = m.M
																								v337 = m.ExcPending
																								if v337 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = l9
																									*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = l8
																									*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l7
																									*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v336
																									*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = l0
																									*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)) = uint8(v5)
																									*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v30
																									return v34
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
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[2]))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
				if base.B2i32(v64 != int32(0)) == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
					v74 = F_LWLockAcquire(m, v70+int32(512), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+124)))
						v80 = v78 | int32(16)
						*(*uint8)(unsafe.Add(mBase, uint32(v77)+124)) = uint8(v80)
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[5]))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
						*(*uint8)(unsafe.Add(mBase, uint32(v84+v85))) = uint8(v80)
						v89 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
						F_LWLockRelease(m, v89+int32(512))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v18
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
							v99 = F_XLogReaderAllocate(m, v98, l6, v34)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v99
								if v99 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v402 = m.ExcPending
									if v402 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_StartupDecodingContext_11))
										mBase = m.M
										v405 = m.ExcPending
										if v405 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
											mBase = m.M
											v409 = m.ExcPending
											if v409 != 0 {
												return int32(0)
											} else {
												F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
												mBase = m.M
												v413 = m.ExcPending
												if v413 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
													mBase = m.M
													v418 = m.ExcPending
													if v418 != 0 {
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
									v104 = m.G0
									v106 = v104 - int32(48)
									m.G0 = v106
									v109 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
									v114 = F_AllocSetContextCreateInternal(m, v109, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v117 = F_MemoryContextAlloc(m, v114, int32(224))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v119 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v106)+40)) = v119
											*(*int64)(unsafe.Add(mBase, uint32(v106)+32)) = v119
											*(*int64)(unsafe.Add(mBase, uint32(v106)+24)) = v119
											*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = v119
											*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v119
											*(*int64)(unsafe.Add(mBase, uint32(v106))) = v119
											*(*int32)(unsafe.Add(mBase, uint32(v117)+120)) = v114
											v135 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v117)+124)) = v135
												v141 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v117)+128)) = v141
													v145 = int32(_a_F_StartupDecodingContext_1)
													v148 = F_GenerationContextCreate(m, v114, int32(_a_F_StartupDecodingContext_18), v145, v145, v145)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v117)+132)) = v148
														*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = int64(34359738372)
														v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)+120))
														*(*int32)(unsafe.Add(mBase, uint32(v106)+40)) = v153
														v158 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v106, int32(1064))
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v160 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v117)+152)) = v160
															v162 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v117)+144)) = v162
															*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v162
															*(*int32)(unsafe.Add(mBase, uint32(v117))) = v158
															v169 = F_pairingheap_allocate(m, int32(1017), v160)
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return int32(0)
															} else {
																v171 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v117)+160)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+136)) = v171
																*(*int32)(unsafe.Add(mBase, uint32(v117)+156)) = v169
																*(*int64)(unsafe.Add(mBase, uint32(v117)+168)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+176)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+184)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+192)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+200)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+208)) = v171
																*(*int64)(unsafe.Add(mBase, uint32(v117)+216)) = v171
																*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = int32(0)
																v193 = v117 + int32(20)
																*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v193
																*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v193
																v197 = v117 + int32(12)
																*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v197
																*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v197
																v201 = v117 + int32(4)
																*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v201
																*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v201
																v205 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
																F_ReorderBufferCleanupSerializedTXNs(m, v205+int32(24))
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v106 + int32(48)
																	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v117
																	v214 = *(*int64)(unsafe.Add(mBase, uint32(v18)+128))
																	v216 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																	v221 = F_AllocSetContextCreateInternal(m, v216, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
																	mBase = m.M
																	v222 = m.ExcPending
																	if v222 != 0 {
																		return int32(0)
																	} else {
																		v223 = int32(_a_F_StartupDecodingContext_3)
																		v224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v221
																		v228 = F_palloc0(m, int32(88))
																		mBase = m.M
																		v229 = m.ExcPending
																		if v229 != 0 {
																			return int32(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v228)+64)) = int64(549755813888)
																			*(*int32)(unsafe.Add(mBase, uint32(v228)+56)) = v117
																			*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v221
																			*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(-1)
																			v237 = F_palloc0(m, int32(512))
																			mBase = m.M
																			v238 = m.ExcPending
																			if v238 != 0 {
																				return int32(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v228)+80)) = int64(0)
																				v241 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v228)+72)) = uint8(v241)
																				*(*int32)(unsafe.Add(mBase, uint32(v228)+76)) = v237
																				*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = l2
																				*(*uint8)(unsafe.Add(mBase, uint32(v228)+37)) = uint8(v6)
																				*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = l1
																				*(*uint8)(unsafe.Add(mBase, uint32(v228)+36)) = uint8(v4)
																				*(*int64)(unsafe.Add(mBase, uint32(v228)+24)) = v214
																				*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v224
																				*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v228
																				v252 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v252)+112)) = v34
																				v254 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v254)+40)) = int32(994)
																				v257 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v257)+44)) = int32(995)
																				v260 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v260)+48)) = int32(996)
																				v263 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v263)+52)) = int32(997)
																				v266 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v266)+56)) = int32(998)
																				v271 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
																				if v271 != 0 {
																					v285 = v241
																				} else {
																					v273 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
																					if v273 != 0 {
																						v285 = int32(1)
																					} else {
																						v275 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
																						if v275 != 0 {
																							v285 = int32(1)
																						} else {
																							v277 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
																							if v277 != 0 {
																								v285 = int32(1)
																							} else {
																								v279 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
																								if v279 != 0 {
																									v285 = int32(1)
																								} else {
																									v281 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
																									if v281 != 0 {
																										v285 = int32(1)
																									} else {
																										v282 = *(*int32)(unsafe.Add(mBase, uint32(v34)+104))
																										v285 = base.B2i32(v282 != int32(0))
																									}
																								}
																							}
																						}
																					}
																				}
																				*(*uint8)(unsafe.Add(mBase, uint32(v34)+144)) = uint8(v285)
																				v287 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v287)+76)) = int32(999)
																				v290 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v290)+80)) = int32(1000)
																				v293 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v293)+84)) = int32(1001)
																				v296 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v296)+88)) = int32(1002)
																				v299 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v299)+92)) = int32(1003)
																				v302 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v302)+96)) = int32(1004)
																				v305 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v305)+100)) = int32(1005)
																				v308 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v308)+104)) = int32(1006)
																				v311 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
																				if v311 != 0 {
																					v319 = v241
																				} else {
																					v312 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
																					if v312 != 0 {
																						v319 = v241
																					} else {
																						v313 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
																						if v313 != 0 {
																							v319 = v241
																						} else {
																							v314 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
																							if v314 != 0 {
																								v319 = v241
																							} else {
																								v315 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
																								if v315 != 0 {
																									v319 = v241
																								} else {
																									v316 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
																									v319 = base.B2i32(v316 != int32(0))
																								}
																							}
																						}
																					}
																				}
																				*(*uint8)(unsafe.Add(mBase, uint32(v34)+145)) = uint8(v319)
																				v321 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v321)+60)) = int32(1007)
																				v324 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v324)+64)) = int32(1008)
																				v327 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v327)+68)) = int32(1009)
																				v330 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v330)+72)) = int32(1010)
																				v333 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v333)+108)) = int32(1011)
																				v336 = F_makeStringInfo(m)
																				mBase = m.M
																				v337 = m.ExcPending
																				if v337 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = l9
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = l8
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l7
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v336
																					*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = l0
																					*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)) = uint8(v5)
																					*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v30
																					return v34
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
					*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v18
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
					v99 = F_XLogReaderAllocate(m, v98, l6, v34)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v99
						if v99 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v402 = m.ExcPending
							if v402 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_StartupDecodingContext_11))
								mBase = m.M
								v405 = m.ExcPending
								if v405 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_StartupDecodingContext_12), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										F_errdetail(m, int32(_a_F_StartupDecodingContext_13), int32(0))
										mBase = m.M
										v413 = m.ExcPending
										if v413 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_StartupDecodingContext_6), int32(213), int32(_a_F_StartupDecodingContext_14))
											mBase = m.M
											v418 = m.ExcPending
											if v418 != 0 {
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
							v104 = m.G0
							v106 = v104 - int32(48)
							m.G0 = v106
							v109 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
							v114 = F_AllocSetContextCreateInternal(m, v109, int32(_a_F_StartupDecodingContext_15), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v117 = F_MemoryContextAlloc(m, v114, int32(224))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									v119 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v106)+40)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v106)+32)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v106)+24)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v106))) = v119
									*(*int32)(unsafe.Add(mBase, uint32(v117)+120)) = v114
									v135 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_16), int32(_a_F_StartupDecodingContext_1), int32(64))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v117)+124)) = v135
										v141 = F_SlabContextCreate(m, v114, int32(_a_F_StartupDecodingContext_17), int32(_a_F_StartupDecodingContext_1), int32(232))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v117)+128)) = v141
											v145 = int32(_a_F_StartupDecodingContext_1)
											v148 = F_GenerationContextCreate(m, v114, int32(_a_F_StartupDecodingContext_18), v145, v145, v145)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v117)+132)) = v148
												*(*int64)(unsafe.Add(mBase, uint32(v106)+16)) = int64(34359738372)
												v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)+120))
												*(*int32)(unsafe.Add(mBase, uint32(v106)+40)) = v153
												v158 = F_hash_create(m, int32(_a_F_StartupDecodingContext_19), int32(1000), v106, int32(1064))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v160 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v117)+152)) = v160
													v162 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v117)+144)) = v162
													*(*int64)(unsafe.Add(mBase, uint32(v117)+32)) = v162
													*(*int32)(unsafe.Add(mBase, uint32(v117))) = v158
													v169 = F_pairingheap_allocate(m, int32(1017), v160)
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														v171 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v117)+160)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+136)) = v171
														*(*int32)(unsafe.Add(mBase, uint32(v117)+156)) = v169
														*(*int64)(unsafe.Add(mBase, uint32(v117)+168)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+176)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+184)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+192)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+200)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+208)) = v171
														*(*int64)(unsafe.Add(mBase, uint32(v117)+216)) = v171
														*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = int32(0)
														v193 = v117 + int32(20)
														*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v193
														*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v193
														v197 = v117 + int32(12)
														*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v197
														*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v197
														v201 = v117 + int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v201
														*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v201
														v205 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
														F_ReorderBufferCleanupSerializedTXNs(m, v205+int32(24))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															m.G0 = v106 + int32(48)
															*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v117
															v214 = *(*int64)(unsafe.Add(mBase, uint32(v18)+128))
															v216 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
															v221 = F_AllocSetContextCreateInternal(m, v216, int32(_a_F_StartupDecodingContext_20), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
															mBase = m.M
															v222 = m.ExcPending
															if v222 != 0 {
																return int32(0)
															} else {
																v223 = int32(_a_F_StartupDecodingContext_3)
																v224 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
																*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v221
																v228 = F_palloc0(m, int32(88))
																mBase = m.M
																v229 = m.ExcPending
																if v229 != 0 {
																	return int32(0)
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v228)+64)) = int64(549755813888)
																	*(*int32)(unsafe.Add(mBase, uint32(v228)+56)) = v117
																	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v221
																	*(*int32)(unsafe.Add(mBase, uint32(v228))) = int32(-1)
																	v237 = F_palloc0(m, int32(512))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return int32(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v228)+80)) = int64(0)
																		v241 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v228)+72)) = uint8(v241)
																		*(*int32)(unsafe.Add(mBase, uint32(v228)+76)) = v237
																		*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = l2
																		*(*uint8)(unsafe.Add(mBase, uint32(v228)+37)) = uint8(v6)
																		*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v228)+36)) = uint8(v4)
																		*(*int64)(unsafe.Add(mBase, uint32(v228)+24)) = v214
																		*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v224
																		*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v228
																		v252 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v252)+112)) = v34
																		v254 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v254)+40)) = int32(994)
																		v257 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v257)+44)) = int32(995)
																		v260 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v260)+48)) = int32(996)
																		v263 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v263)+52)) = int32(997)
																		v266 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v266)+56)) = int32(998)
																		v271 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
																		if v271 != 0 {
																			v285 = v241
																		} else {
																			v273 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
																			if v273 != 0 {
																				v285 = int32(1)
																			} else {
																				v275 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
																				if v275 != 0 {
																					v285 = int32(1)
																				} else {
																					v277 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
																					if v277 != 0 {
																						v285 = int32(1)
																					} else {
																						v279 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
																						if v279 != 0 {
																							v285 = int32(1)
																						} else {
																							v281 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
																							if v281 != 0 {
																								v285 = int32(1)
																							} else {
																								v282 = *(*int32)(unsafe.Add(mBase, uint32(v34)+104))
																								v285 = base.B2i32(v282 != int32(0))
																							}
																						}
																					}
																				}
																			}
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v34)+144)) = uint8(v285)
																		v287 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v287)+76)) = int32(999)
																		v290 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v290)+80)) = int32(1000)
																		v293 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v293)+84)) = int32(1001)
																		v296 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v296)+88)) = int32(1002)
																		v299 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v299)+92)) = int32(1003)
																		v302 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v302)+96)) = int32(1004)
																		v305 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v305)+100)) = int32(1005)
																		v308 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v308)+104)) = int32(1006)
																		v311 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
																		if v311 != 0 {
																			v319 = v241
																		} else {
																			v312 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
																			if v312 != 0 {
																				v319 = v241
																			} else {
																				v313 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
																				if v313 != 0 {
																					v319 = v241
																				} else {
																					v314 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
																					if v314 != 0 {
																						v319 = v241
																					} else {
																						v315 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
																						if v315 != 0 {
																							v319 = v241
																						} else {
																							v316 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
																							v319 = base.B2i32(v316 != int32(0))
																						}
																					}
																				}
																			}
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v34)+145)) = uint8(v319)
																		v321 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v321)+60)) = int32(1007)
																		v324 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v324)+64)) = int32(1008)
																		v327 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v327)+68)) = int32(1009)
																		v330 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v330)+72)) = int32(1010)
																		v333 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v333)+108)) = int32(1011)
																		v336 = F_makeStringInfo(m)
																		mBase = m.M
																		v337 = m.ExcPending
																		if v337 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = l9
																			*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = l8
																			*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l7
																			*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v336
																			*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = l0
																			*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)) = uint8(v5)
																			*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v30
																			return v34
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int64
	_ = v220
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
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
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int64
	_ = v413
	var v415 int64
	_ = v415
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
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
			v17 = v15 - int32(32)
			m.G0 = v17
			switch int32(960) {
			case 0, 2:
				v27 = v13
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[1])) = v13
				v27 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v27
			F_sigemptyset(m, v17+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(268435456)
			v39 = v17 + int32(12)
			if v39 != 0 {
				v45 = int32(20)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[2])) = v47
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[3])) = v49
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[4])) = v51
			} else {
			}
			m.G0 = v17 + int32(32)
			v59 = int32(-2)
			v61 = m.G0
			v63 = v61 - int32(32)
			m.G0 = v63
			switch int32(0) {
			case 0, 2:
				v73 = v59
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[5])) = v59
				v73 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v73
			F_sigemptyset(m, v63+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(268435456)
			v85 = v63 + int32(12)
			if v85 != 0 {
				v92 = int32(40)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[6])) = v93
				v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[7])) = v95
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[8])) = v97
			} else {
			}
			m.G0 = v63 + int32(32)
			v105 = int32(959)
			v107 = m.G0
			v109 = v107 - int32(32)
			m.G0 = v109
			switch int32(961) {
			case 0, 2:
				v119 = v105
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[9])) = v105
				v119 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v119
			F_sigemptyset(m, v109+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(268435456)
			v131 = v109 + int32(12)
			if v131 != 0 {
				v138 = int32(300)
				v139 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[10])) = v139
				v141 = *(*int64)(unsafe.Add(mBase, uint32(v131)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[11])) = v141
				v143 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[12])) = v143
			} else {
			}
			m.G0 = v109 + int32(32)
			v150 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[13])) = v150
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[14])) = v150
			v160 = v150
			for {
				v162 = int32(40)
				v163 = v160 * v162
				v164 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[15]))) = uint8(v164)
				*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[16]))) = v160
				v167 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[17]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[18]))) = v164
				*(*int64)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[19]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[20]))) = v164
				*(*uint8)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_StartupProcessMain[21]))) = uint8(v164)
				v178 = v160 | int32(1)
				v180 = v178 * v162
				*(*uint8)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[15]))) = uint8(v164)
				*(*int32)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[16]))) = v178
				*(*int64)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[17]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[18]))) = v164
				*(*int64)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[19]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[20]))) = v164
				*(*uint8)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_StartupProcessMain[21]))) = uint8(v164)
				v195 = v160 | int32(2)
				v197 = v195 * v162
				*(*uint8)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[15]))) = uint8(v164)
				*(*int32)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[16]))) = v195
				*(*int64)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[17]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[18]))) = v164
				*(*int64)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[19]))) = v167
				*(*int32)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[20]))) = v164
				*(*uint8)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_StartupProcessMain[21]))) = uint8(v164)
				if v160 != int32(20) {
					v214 = v160 | int32(3)
					v216 = v214 * int32(40)
					v217 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[15]))) = uint8(v217)
					*(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[16]))) = v214
					v220 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[17]))) = v220
					*(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[18]))) = v217
					*(*int64)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[19]))) = v220
					*(*int32)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[20]))) = v217
					*(*uint8)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_StartupProcessMain[21]))) = uint8(v217)
					v160 = v160 + int32(4)
					continue
				} else {
					break
				}
				break
			}
			v233 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_StartupProcessMain[22])) = uint8(v233)
			F_pqsignal_be(m, int32(14), int32(1769))
			mBase = m.M
			v239 = int32(-2)
			v241 = m.G0
			v243 = v241 - int32(32)
			m.G0 = v243
			switch int32(0) {
			case 0, 2:
				v253 = v239
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[23])) = v239
				v253 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = v253
			F_sigemptyset(m, v243+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v243)+24)) = int32(268435456)
			v265 = v243 + int32(12)
			if v265 != 0 {
				v272 = int32(260)
				v273 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[24])) = v273
				v275 = *(*int64)(unsafe.Add(mBase, uint32(v265)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[25])) = v275
				v277 = *(*int64)(unsafe.Add(mBase, uint32(v265)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[26])) = v277
			} else {
			}
			m.G0 = v243 + int32(32)
			v285 = int32(917)
			v287 = m.G0
			v289 = v287 - int32(32)
			m.G0 = v289
			switch int32(919) {
			case 0, 2:
				v299 = v285
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[27])) = v285
				v299 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v289)+12)) = v299
			F_sigemptyset(m, v289+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v289)+24)) = int32(268435456)
			v311 = v289 + int32(12)
			if v311 != 0 {
				v318 = int32(200)
				v319 = *(*int32)(unsafe.Add(mBase, uint32(v311)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[28])) = v319
				v321 = *(*int64)(unsafe.Add(mBase, uint32(v311)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[29])) = v321
				v323 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[30])) = v323
			} else {
			}
			m.G0 = v289 + int32(32)
			v331 = int32(960)
			v333 = m.G0
			v335 = v333 - int32(32)
			m.G0 = v335
			switch int32(962) {
			case 0, 2:
				v345 = v331
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[31])) = v331
				v345 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v335)+12)) = v345
			F_sigemptyset(m, v335+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v335)+24)) = int32(268435456)
			v357 = v335 + int32(12)
			if v357 != 0 {
				v364 = int32(240)
				v365 = *(*int32)(unsafe.Add(mBase, uint32(v357)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[32])) = v365
				v367 = *(*int64)(unsafe.Add(mBase, uint32(v357)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[33])) = v367
				v369 = *(*int64)(unsafe.Add(mBase, uint32(v357)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[34])) = v369
			} else {
			}
			m.G0 = v335 + int32(32)
			v377 = int32(0)
			v379 = m.G0
			v381 = v379 - int32(32)
			m.G0 = v381
			switch int32(2) {
			case 0, 2:
				v391 = v377
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[35])) = v377
				v391 = int32(_a_F_StartupProcessMain_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v381)+12)) = v391
			F_sigemptyset(m, v381+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = int32(268435457)
			v403 = v381 + int32(12)
			if v403 != 0 {
				v410 = int32(340)
				v411 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[36])) = v411
				v413 = *(*int64)(unsafe.Add(mBase, uint32(v403)+8))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[37])) = v413
				v415 = *(*int64)(unsafe.Add(mBase, uint32(v403)))
				*(*int64)(unsafe.Add(mBase, _c_F_StartupProcessMain[38])) = v415
			} else {
			}
			m.G0 = v381 + int32(32)
			F_RegisterTimeout(m, int32(4), int32(961))
			mBase = m.M
			v425 = m.ExcPending
			if v425 != 0 {
				return
			} else {
				F_RegisterTimeout(m, int32(5), int32(962))
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return
				} else {
					F_RegisterTimeout(m, int32(6), int32(963))
					mBase = m.M
					v433 = m.ExcPending
					if v433 != 0 {
						return
					} else {
						F_sigprocmask(m, int32(_a_F_StartupProcessMain_1), int32(0))
						mBase = m.M
						v437 = m.ExcPending
						if v437 != 0 {
							return
						} else {
							F_StartupXLOG(m)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v442 = m.ExcPending
								if v442 != 0 {
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
