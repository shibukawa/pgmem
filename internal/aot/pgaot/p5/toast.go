package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(304)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+112))
	if v18 != 0 {
		v326 = v8
		m.G0 = v15 + int32(304)
		return v326
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[79])))
		if v21 == int32(0) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
			if v24 == int32(112) {
				v326 = v8
				m.G0 = v15 + int32(304)
				return v326
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
				if v27 == int32(1) {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[231]))
					if v31 != 0 {
						v326 = v8
						m.G0 = v15 + int32(304)
						return v326
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if base.Ui32(v32) < base.Ui32(int32(12000)) {
							v36 = *(*int32)(unsafe.Add(mBase, _consts[231]))
							if v36 != 0 {
								v326 = v8
								m.G0 = v15 + int32(304)
								return v326
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
								v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, l0)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v39 != 0 {
										if l4 != int32(8) {
											v50 = l5
										} else {
											v50 = int32(0)
										}
										if v50 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v337 = m.ExcPending
											if v337 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(619417), int32(0))
												mBase = m.M
												v341 = m.ExcPending
												if v341 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(493186), int32(193), int32(388317))
													mBase = m.M
													v346 = m.ExcPending
													if v346 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
											v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(38255), v15+int32(48))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
												v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(27939), v15+int32(32))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v71 = F_CreateTemplateTupleDesc(m, int32(3))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														F_TupleDescInitEntry(m, v71, int32(1), int32(431902), int32(26), int32(-1), int32(0))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															F_TupleDescInitEntry(m, v71, int32(2), int32(228934), int32(23), int32(-1), int32(0))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return int32(0)
															} else {
																F_TupleDescInitEntry(m, v71, int32(3), int32(499016), int32(17), int32(-1), int32(0))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																	v95 = int32(4)
																	v97 = v71 + v94<<(uint(v95)%32)
																	v98 = int32(112)
																	*(*uint8)(unsafe.Add(mBase, uint32(v97)+304)) = uint8(v98)
																	*(*uint8)(unsafe.Add(mBase, uint32(v97)+204)) = uint8(v98)
																	*(*uint8)(unsafe.Add(mBase, uint32(v97)+104)) = uint8(v98)
																	v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																	v107 = v71 + v104<<(uint(v95)%32)
																	v108 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v107)+305)) = uint8(v108)
																	*(*uint8)(unsafe.Add(mBase, uint32(v107)+205)) = uint8(v108)
																	*(*uint8)(unsafe.Add(mBase, uint32(v107)+105)) = uint8(v108)
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
																	v120 = *(*int32)(unsafe.Add(mBase, _consts[249]))
																	if v120 != 0 {
																		v121 = int32(1)
																		if v116 == v120 {
																			v128 = v121
																		} else {
																			v124 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																			if v124 == v116 {
																				v128 = v121
																			} else {
																				v128 = int32(0)
																			}
																		}
																	} else {
																		v128 = int32(0)
																	}
																	if v128 != 0 {
																		v130 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																		v131 = v130
																	} else {
																		v131 = int32(99)
																	}
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+117)))
																	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
																	switch v134 - int32(83) {
																	case 0, 22, 26, 31, 33:
																		v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
																		v140 = base.B2i32(v137 == int32(0))
																	default:
																		v140 = int32(0)
																	}
																	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
																	v144 = int32(0)
																	v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+156))
																	v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return int32(0)
																	} else {
																		v151 = int32(0)
																		v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																		v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+118)))
																		v155 = int32(1)
																		v162 = F_heap_create_with_catalog(m, v15+int32(224), v131, v143, l1, v144, v144, v146, v149, v71, v151, int32(116), v154, v133&v155, v140, v151, l3, v151, v155, v155, l6, v151)
																		mBase = m.M
																		v163 = m.ExcPending
																		if v163 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v165 = m.ExcPending
																			if v165 != 0 {
																				return int32(0)
																			} else {
																				v167 = F_table_open(m, v162, int32(5))
																				mBase = m.M
																				v168 = m.ExcPending
																				if v168 != 0 {
																					return int32(0)
																				} else {
																					v170 = F_palloc0(m, int32(144))
																					mBase = m.M
																					v171 = m.ExcPending
																					if v171 != 0 {
																						return int32(0)
																					} else {
																						v172 = int64(0)
																						*(*int64)(unsafe.Add(mBase, uint32(v170)+76)) = v172
																						*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(562954248388610)
																						*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(8589934973)
																						v178 = int32(0)
																						*(*int32)(unsafe.Add(mBase, uint32(v170)+128)) = v178
																						v180 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(v170)+118)) = uint8(v180)
																						*(*uint16)(unsafe.Add(mBase, uint32(v170)+116)) = uint16(v180)
																						*(*int64)(unsafe.Add(mBase, uint32(v170)+132)) = int64(403)
																						*(*int32)(unsafe.Add(mBase, uint32(v170)+119)) = v178
																						*(*int64)(unsafe.Add(mBase, uint32(v170)+84)) = v172
																						*(*int64)(unsafe.Add(mBase, uint32(v170)+92)) = v172
																						*(*int32)(unsafe.Add(mBase, uint32(v170)+100)) = v178
																						v195 = *(*int32)(unsafe.Add(mBase, _consts[9]))
																						*(*int32)(unsafe.Add(mBase, uint32(v170)+140)) = v195
																						*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
																						*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v172
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v178
																						v203 = int32(431902)
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v203
																						v205 = int32(228934)
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v205
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v203
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v205
																						v220 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
																						mBase = m.M
																						v221 = m.ExcPending
																						if v221 != 0 {
																							return int32(0)
																						} else {
																							v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																							v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
																							v229 = int32(0)
																							v234 = int32(1)
																							v239 = F_index_create(m, v167, v15+int32(160), l2, v178, v178, v178, v170, v220, int32(403), v224, v15+int32(152), v15+int32(144), v229, v15+int32(140), v229, v229, v234, v229, v234, v234, v229)
																							mBase = m.M
																							v240 = m.ExcPending
																							if v240 != 0 {
																								return int32(0)
																							} else {
																								F_sequence_close(m, v167, int32(0))
																								mBase = m.M
																								v243 = m.ExcPending
																								if v243 != 0 {
																									return int32(0)
																								} else {
																									v246 = F_table_open(m, int32(1259), int32(3))
																									mBase = m.M
																									v247 = m.ExcPending
																									if v247 != 0 {
																										return int32(0)
																									} else {
																										v249 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																										if v249 != 0 {
																											v252 = F_SearchSysCacheCopy(m, int32(57), v19, int32(0))
																											mBase = m.M
																											v253 = m.ExcPending
																											if v253 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v252
																												if v252 == int32(0) {
																													F_errstart_cold(m, int32(21), int32(0))
																													mBase = m.M
																													v350 = m.ExcPending
																													if v350 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
																														F_errmsg_internal(m, int32(46032), v15+int32(16))
																														mBase = m.M
																														v356 = m.ExcPending
																														if v356 != 0 {
																															return int32(0)
																														} else {
																															F_errfinish(m, int32(493186), int32(342), int32(388317))
																															mBase = m.M
																															v361 = m.ExcPending
																															if v361 != 0 {
																																return int32(0)
																															} else {
																																base.Wasm_trap_unreachable()
																																for {
																																}
																															}
																														}
																													}
																												} else {
																													v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
																													v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
																													*(*int32)(unsafe.Add(mBase, uint32(v257+v258)+112)) = v162
																													F_CatalogTupleUpdate(m, v246, v252+int32(4), v252)
																													mBase = m.M
																													v264 = m.ExcPending
																													if v264 != 0 {
																														return int32(0)
																													} else {
																														v293 = v252
																														F_pfree(m, v293)
																														mBase = m.M
																														v296 = m.ExcPending
																														if v296 != 0 {
																															return int32(0)
																														} else {
																															F_sequence_close(m, v246, int32(3))
																															mBase = m.M
																															v299 = m.ExcPending
																															if v299 != 0 {
																																return int32(0)
																															} else {
																																v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																																if v301 != 0 {
																																	v302 = int32(0)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																	v305 = int32(1259)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																	F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																	mBase = m.M
																																	v318 = m.ExcPending
																																	if v318 != 0 {
																																		return int32(0)
																																	} else {
																																		F_CommandCounterIncrement(m)
																																		mBase = m.M
																																		v320 = m.ExcPending
																																		if v320 != 0 {
																																			return int32(0)
																																		} else {
																																			v326 = int32(1)
																																			m.G0 = v15 + int32(304)
																																			return v326
																																		}
																																	}
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
																																	}
																																}
																															}
																														}
																													}
																												}
																											}
																										} else {
																											F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v19)
																											mBase = m.M
																											v271 = m.ExcPending
																											if v271 != 0 {
																												return int32(0)
																											} else {
																												F_systable_inplace_update_begin(m, v246, int32(2662), v15-int32(-64), v15+int32(300), v15+int32(128))
																												mBase = m.M
																												v280 = m.ExcPending
																												if v280 != 0 {
																													return int32(0)
																												} else {
																													v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																													if v281 == int32(0) {
																														F_errstart_cold(m, int32(21), int32(0))
																														mBase = m.M
																														v365 = m.ExcPending
																														if v365 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
																															F_errmsg_internal(m, int32(46032), v15)
																															mBase = m.M
																															v369 = m.ExcPending
																															if v369 != 0 {
																																return int32(0)
																															} else {
																																F_errfinish(m, int32(493186), int32(362), int32(388317))
																																mBase = m.M
																																v374 = m.ExcPending
																																if v374 != 0 {
																																	return int32(0)
																																} else {
																																	base.Wasm_trap_unreachable()
																																	for {
																																	}
																																}
																															}
																														}
																													} else {
																														v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
																														v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+22)))
																														*(*int32)(unsafe.Add(mBase, uint32(v284+v285)+112)) = v162
																														v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
																														v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																														F_systable_inplace_update_finish(m, v288, v289)
																														mBase = m.M
																														v291 = m.ExcPending
																														if v291 != 0 {
																															return int32(0)
																														} else {
																															v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																															v293 = v292
																															F_pfree(m, v293)
																															mBase = m.M
																															v296 = m.ExcPending
																															if v296 != 0 {
																																return int32(0)
																															} else {
																																F_sequence_close(m, v246, int32(3))
																																mBase = m.M
																																v299 = m.ExcPending
																																if v299 != 0 {
																																	return int32(0)
																																} else {
																																	v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																																	if v301 != 0 {
																																		v302 = int32(0)
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																		v305 = int32(1259)
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																		*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																		F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																		mBase = m.M
																																		v318 = m.ExcPending
																																		if v318 != 0 {
																																			return int32(0)
																																		} else {
																																			F_CommandCounterIncrement(m)
																																			mBase = m.M
																																			v320 = m.ExcPending
																																			if v320 != 0 {
																																				return int32(0)
																																			} else {
																																				v326 = int32(1)
																																				m.G0 = v15 + int32(304)
																																				return v326
																																			}
																																		}
																																	} else {
																																		F_CommandCounterIncrement(m)
																																		mBase = m.M
																																		v320 = m.ExcPending
																																		if v320 != 0 {
																																			return int32(0)
																																		} else {
																																			v326 = int32(1)
																																			m.G0 = v15 + int32(304)
																																			return v326
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
													}
												}
											}
										}
									} else {
										v326 = v8
										m.G0 = v15 + int32(304)
										return v326
									}
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
							v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								if v39 != 0 {
									if l4 != int32(8) {
										v50 = l5
									} else {
										v50 = int32(0)
									}
									if v50 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v337 = m.ExcPending
										if v337 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(619417), int32(0))
											mBase = m.M
											v341 = m.ExcPending
											if v341 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493186), int32(193), int32(388317))
												mBase = m.M
												v346 = m.ExcPending
												if v346 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
										v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(38255), v15+int32(48))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
											v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(27939), v15+int32(32))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v71 = F_CreateTemplateTupleDesc(m, int32(3))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													F_TupleDescInitEntry(m, v71, int32(1), int32(431902), int32(26), int32(-1), int32(0))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														F_TupleDescInitEntry(m, v71, int32(2), int32(228934), int32(23), int32(-1), int32(0))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_TupleDescInitEntry(m, v71, int32(3), int32(499016), int32(17), int32(-1), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																v95 = int32(4)
																v97 = v71 + v94<<(uint(v95)%32)
																v98 = int32(112)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+304)) = uint8(v98)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+204)) = uint8(v98)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+104)) = uint8(v98)
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																v107 = v71 + v104<<(uint(v95)%32)
																v108 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+305)) = uint8(v108)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+205)) = uint8(v108)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+105)) = uint8(v108)
																v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
																v120 = *(*int32)(unsafe.Add(mBase, _consts[249]))
																if v120 != 0 {
																	v121 = int32(1)
																	if v116 == v120 {
																		v128 = v121
																	} else {
																		v124 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																		if v124 == v116 {
																			v128 = v121
																		} else {
																			v128 = int32(0)
																		}
																	}
																} else {
																	v128 = int32(0)
																}
																if v128 != 0 {
																	v130 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																	v131 = v130
																} else {
																	v131 = int32(99)
																}
																v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+117)))
																v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
																switch v134 - int32(83) {
																case 0, 22, 26, 31, 33:
																	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
																	v140 = base.B2i32(v137 == int32(0))
																default:
																	v140 = int32(0)
																}
																v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
																v144 = int32(0)
																v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
																v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+156))
																v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return int32(0)
																} else {
																	v151 = int32(0)
																	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+118)))
																	v155 = int32(1)
																	v162 = F_heap_create_with_catalog(m, v15+int32(224), v131, v143, l1, v144, v144, v146, v149, v71, v151, int32(116), v154, v133&v155, v140, v151, l3, v151, v155, v155, l6, v151)
																	mBase = m.M
																	v163 = m.ExcPending
																	if v163 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v165 = m.ExcPending
																		if v165 != 0 {
																			return int32(0)
																		} else {
																			v167 = F_table_open(m, v162, int32(5))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return int32(0)
																			} else {
																				v170 = F_palloc0(m, int32(144))
																				mBase = m.M
																				v171 = m.ExcPending
																				if v171 != 0 {
																					return int32(0)
																				} else {
																					v172 = int64(0)
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+76)) = v172
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(562954248388610)
																					*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(8589934973)
																					v178 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+128)) = v178
																					v180 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v170)+118)) = uint8(v180)
																					*(*uint16)(unsafe.Add(mBase, uint32(v170)+116)) = uint16(v180)
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+132)) = int64(403)
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+119)) = v178
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+84)) = v172
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+92)) = v172
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+100)) = v178
																					v195 = *(*int32)(unsafe.Add(mBase, _consts[9]))
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+140)) = v195
																					*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
																					*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v172
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v178
																					v203 = int32(431902)
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v203
																					v205 = int32(228934)
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v205
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v203
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v205
																					v220 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																						v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
																						v229 = int32(0)
																						v234 = int32(1)
																						v239 = F_index_create(m, v167, v15+int32(160), l2, v178, v178, v178, v170, v220, int32(403), v224, v15+int32(152), v15+int32(144), v229, v15+int32(140), v229, v229, v234, v229, v234, v234, v229)
																						mBase = m.M
																						v240 = m.ExcPending
																						if v240 != 0 {
																							return int32(0)
																						} else {
																							F_sequence_close(m, v167, int32(0))
																							mBase = m.M
																							v243 = m.ExcPending
																							if v243 != 0 {
																								return int32(0)
																							} else {
																								v246 = F_table_open(m, int32(1259), int32(3))
																								mBase = m.M
																								v247 = m.ExcPending
																								if v247 != 0 {
																									return int32(0)
																								} else {
																									v249 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																									if v249 != 0 {
																										v252 = F_SearchSysCacheCopy(m, int32(57), v19, int32(0))
																										mBase = m.M
																										v253 = m.ExcPending
																										if v253 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v252
																											if v252 == int32(0) {
																												F_errstart_cold(m, int32(21), int32(0))
																												mBase = m.M
																												v350 = m.ExcPending
																												if v350 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
																													F_errmsg_internal(m, int32(46032), v15+int32(16))
																													mBase = m.M
																													v356 = m.ExcPending
																													if v356 != 0 {
																														return int32(0)
																													} else {
																														F_errfinish(m, int32(493186), int32(342), int32(388317))
																														mBase = m.M
																														v361 = m.ExcPending
																														if v361 != 0 {
																															return int32(0)
																														} else {
																															base.Wasm_trap_unreachable()
																															for {
																															}
																														}
																													}
																												}
																											} else {
																												v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
																												v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
																												*(*int32)(unsafe.Add(mBase, uint32(v257+v258)+112)) = v162
																												F_CatalogTupleUpdate(m, v246, v252+int32(4), v252)
																												mBase = m.M
																												v264 = m.ExcPending
																												if v264 != 0 {
																													return int32(0)
																												} else {
																													v293 = v252
																													F_pfree(m, v293)
																													mBase = m.M
																													v296 = m.ExcPending
																													if v296 != 0 {
																														return int32(0)
																													} else {
																														F_sequence_close(m, v246, int32(3))
																														mBase = m.M
																														v299 = m.ExcPending
																														if v299 != 0 {
																															return int32(0)
																														} else {
																															v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																															if v301 != 0 {
																																v302 = int32(0)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																v305 = int32(1259)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																mBase = m.M
																																v318 = m.ExcPending
																																if v318 != 0 {
																																	return int32(0)
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
																																	}
																																}
																															} else {
																																F_CommandCounterIncrement(m)
																																mBase = m.M
																																v320 = m.ExcPending
																																if v320 != 0 {
																																	return int32(0)
																																} else {
																																	v326 = int32(1)
																																	m.G0 = v15 + int32(304)
																																	return v326
																																}
																															}
																														}
																													}
																												}
																											}
																										}
																									} else {
																										F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v19)
																										mBase = m.M
																										v271 = m.ExcPending
																										if v271 != 0 {
																											return int32(0)
																										} else {
																											F_systable_inplace_update_begin(m, v246, int32(2662), v15-int32(-64), v15+int32(300), v15+int32(128))
																											mBase = m.M
																											v280 = m.ExcPending
																											if v280 != 0 {
																												return int32(0)
																											} else {
																												v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																												if v281 == int32(0) {
																													F_errstart_cold(m, int32(21), int32(0))
																													mBase = m.M
																													v365 = m.ExcPending
																													if v365 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
																														F_errmsg_internal(m, int32(46032), v15)
																														mBase = m.M
																														v369 = m.ExcPending
																														if v369 != 0 {
																															return int32(0)
																														} else {
																															F_errfinish(m, int32(493186), int32(362), int32(388317))
																															mBase = m.M
																															v374 = m.ExcPending
																															if v374 != 0 {
																																return int32(0)
																															} else {
																																base.Wasm_trap_unreachable()
																																for {
																																}
																															}
																														}
																													}
																												} else {
																													v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
																													v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+22)))
																													*(*int32)(unsafe.Add(mBase, uint32(v284+v285)+112)) = v162
																													v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
																													v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																													F_systable_inplace_update_finish(m, v288, v289)
																													mBase = m.M
																													v291 = m.ExcPending
																													if v291 != 0 {
																														return int32(0)
																													} else {
																														v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																														v293 = v292
																														F_pfree(m, v293)
																														mBase = m.M
																														v296 = m.ExcPending
																														if v296 != 0 {
																															return int32(0)
																														} else {
																															F_sequence_close(m, v246, int32(3))
																															mBase = m.M
																															v299 = m.ExcPending
																															if v299 != 0 {
																																return int32(0)
																															} else {
																																v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																																if v301 != 0 {
																																	v302 = int32(0)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																	v305 = int32(1259)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																	F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																	mBase = m.M
																																	v318 = m.ExcPending
																																	if v318 != 0 {
																																		return int32(0)
																																	} else {
																																		F_CommandCounterIncrement(m)
																																		mBase = m.M
																																		v320 = m.ExcPending
																																		if v320 != 0 {
																																			return int32(0)
																																		} else {
																																			v326 = int32(1)
																																			m.G0 = v15 + int32(304)
																																			return v326
																																		}
																																	}
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
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
												}
											}
										}
									}
								} else {
									v326 = v8
									m.G0 = v15 + int32(304)
									return v326
								}
							}
						}
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if base.Ui32(v32) < base.Ui32(int32(12000)) {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[231]))
						if v36 != 0 {
							v326 = v8
							m.G0 = v15 + int32(304)
							return v326
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
							v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, l0)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								if v39 != 0 {
									if l4 != int32(8) {
										v50 = l5
									} else {
										v50 = int32(0)
									}
									if v50 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v337 = m.ExcPending
										if v337 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(619417), int32(0))
											mBase = m.M
											v341 = m.ExcPending
											if v341 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493186), int32(193), int32(388317))
												mBase = m.M
												v346 = m.ExcPending
												if v346 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
										v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(38255), v15+int32(48))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
											v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(27939), v15+int32(32))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v71 = F_CreateTemplateTupleDesc(m, int32(3))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													F_TupleDescInitEntry(m, v71, int32(1), int32(431902), int32(26), int32(-1), int32(0))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														F_TupleDescInitEntry(m, v71, int32(2), int32(228934), int32(23), int32(-1), int32(0))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_TupleDescInitEntry(m, v71, int32(3), int32(499016), int32(17), int32(-1), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																v95 = int32(4)
																v97 = v71 + v94<<(uint(v95)%32)
																v98 = int32(112)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+304)) = uint8(v98)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+204)) = uint8(v98)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+104)) = uint8(v98)
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
																v107 = v71 + v104<<(uint(v95)%32)
																v108 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+305)) = uint8(v108)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+205)) = uint8(v108)
																*(*uint8)(unsafe.Add(mBase, uint32(v107)+105)) = uint8(v108)
																v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
																v120 = *(*int32)(unsafe.Add(mBase, _consts[249]))
																if v120 != 0 {
																	v121 = int32(1)
																	if v116 == v120 {
																		v128 = v121
																	} else {
																		v124 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																		if v124 == v116 {
																			v128 = v121
																		} else {
																			v128 = int32(0)
																		}
																	}
																} else {
																	v128 = int32(0)
																}
																if v128 != 0 {
																	v130 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																	v131 = v130
																} else {
																	v131 = int32(99)
																}
																v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+117)))
																v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
																switch v134 - int32(83) {
																case 0, 22, 26, 31, 33:
																	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
																	v140 = base.B2i32(v137 == int32(0))
																default:
																	v140 = int32(0)
																}
																v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
																v144 = int32(0)
																v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
																v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+156))
																v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return int32(0)
																} else {
																	v151 = int32(0)
																	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+118)))
																	v155 = int32(1)
																	v162 = F_heap_create_with_catalog(m, v15+int32(224), v131, v143, l1, v144, v144, v146, v149, v71, v151, int32(116), v154, v133&v155, v140, v151, l3, v151, v155, v155, l6, v151)
																	mBase = m.M
																	v163 = m.ExcPending
																	if v163 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v165 = m.ExcPending
																		if v165 != 0 {
																			return int32(0)
																		} else {
																			v167 = F_table_open(m, v162, int32(5))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return int32(0)
																			} else {
																				v170 = F_palloc0(m, int32(144))
																				mBase = m.M
																				v171 = m.ExcPending
																				if v171 != 0 {
																					return int32(0)
																				} else {
																					v172 = int64(0)
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+76)) = v172
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(562954248388610)
																					*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(8589934973)
																					v178 = int32(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+128)) = v178
																					v180 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v170)+118)) = uint8(v180)
																					*(*uint16)(unsafe.Add(mBase, uint32(v170)+116)) = uint16(v180)
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+132)) = int64(403)
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+119)) = v178
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+84)) = v172
																					*(*int64)(unsafe.Add(mBase, uint32(v170)+92)) = v172
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+100)) = v178
																					v195 = *(*int32)(unsafe.Add(mBase, _consts[9]))
																					*(*int32)(unsafe.Add(mBase, uint32(v170)+140)) = v195
																					*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
																					*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v172
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v178
																					v203 = int32(431902)
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v203
																					v205 = int32(228934)
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v205
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v203
																					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v205
																					v220 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																						v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
																						v229 = int32(0)
																						v234 = int32(1)
																						v239 = F_index_create(m, v167, v15+int32(160), l2, v178, v178, v178, v170, v220, int32(403), v224, v15+int32(152), v15+int32(144), v229, v15+int32(140), v229, v229, v234, v229, v234, v234, v229)
																						mBase = m.M
																						v240 = m.ExcPending
																						if v240 != 0 {
																							return int32(0)
																						} else {
																							F_sequence_close(m, v167, int32(0))
																							mBase = m.M
																							v243 = m.ExcPending
																							if v243 != 0 {
																								return int32(0)
																							} else {
																								v246 = F_table_open(m, int32(1259), int32(3))
																								mBase = m.M
																								v247 = m.ExcPending
																								if v247 != 0 {
																									return int32(0)
																								} else {
																									v249 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																									if v249 != 0 {
																										v252 = F_SearchSysCacheCopy(m, int32(57), v19, int32(0))
																										mBase = m.M
																										v253 = m.ExcPending
																										if v253 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v252
																											if v252 == int32(0) {
																												F_errstart_cold(m, int32(21), int32(0))
																												mBase = m.M
																												v350 = m.ExcPending
																												if v350 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
																													F_errmsg_internal(m, int32(46032), v15+int32(16))
																													mBase = m.M
																													v356 = m.ExcPending
																													if v356 != 0 {
																														return int32(0)
																													} else {
																														F_errfinish(m, int32(493186), int32(342), int32(388317))
																														mBase = m.M
																														v361 = m.ExcPending
																														if v361 != 0 {
																															return int32(0)
																														} else {
																															base.Wasm_trap_unreachable()
																															for {
																															}
																														}
																													}
																												}
																											} else {
																												v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
																												v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
																												*(*int32)(unsafe.Add(mBase, uint32(v257+v258)+112)) = v162
																												F_CatalogTupleUpdate(m, v246, v252+int32(4), v252)
																												mBase = m.M
																												v264 = m.ExcPending
																												if v264 != 0 {
																													return int32(0)
																												} else {
																													v293 = v252
																													F_pfree(m, v293)
																													mBase = m.M
																													v296 = m.ExcPending
																													if v296 != 0 {
																														return int32(0)
																													} else {
																														F_sequence_close(m, v246, int32(3))
																														mBase = m.M
																														v299 = m.ExcPending
																														if v299 != 0 {
																															return int32(0)
																														} else {
																															v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																															if v301 != 0 {
																																v302 = int32(0)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																v305 = int32(1259)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																mBase = m.M
																																v318 = m.ExcPending
																																if v318 != 0 {
																																	return int32(0)
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
																																	}
																																}
																															} else {
																																F_CommandCounterIncrement(m)
																																mBase = m.M
																																v320 = m.ExcPending
																																if v320 != 0 {
																																	return int32(0)
																																} else {
																																	v326 = int32(1)
																																	m.G0 = v15 + int32(304)
																																	return v326
																																}
																															}
																														}
																													}
																												}
																											}
																										}
																									} else {
																										F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v19)
																										mBase = m.M
																										v271 = m.ExcPending
																										if v271 != 0 {
																											return int32(0)
																										} else {
																											F_systable_inplace_update_begin(m, v246, int32(2662), v15-int32(-64), v15+int32(300), v15+int32(128))
																											mBase = m.M
																											v280 = m.ExcPending
																											if v280 != 0 {
																												return int32(0)
																											} else {
																												v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																												if v281 == int32(0) {
																													F_errstart_cold(m, int32(21), int32(0))
																													mBase = m.M
																													v365 = m.ExcPending
																													if v365 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
																														F_errmsg_internal(m, int32(46032), v15)
																														mBase = m.M
																														v369 = m.ExcPending
																														if v369 != 0 {
																															return int32(0)
																														} else {
																															F_errfinish(m, int32(493186), int32(362), int32(388317))
																															mBase = m.M
																															v374 = m.ExcPending
																															if v374 != 0 {
																																return int32(0)
																															} else {
																																base.Wasm_trap_unreachable()
																																for {
																																}
																															}
																														}
																													}
																												} else {
																													v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
																													v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+22)))
																													*(*int32)(unsafe.Add(mBase, uint32(v284+v285)+112)) = v162
																													v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
																													v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																													F_systable_inplace_update_finish(m, v288, v289)
																													mBase = m.M
																													v291 = m.ExcPending
																													if v291 != 0 {
																														return int32(0)
																													} else {
																														v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																														v293 = v292
																														F_pfree(m, v293)
																														mBase = m.M
																														v296 = m.ExcPending
																														if v296 != 0 {
																															return int32(0)
																														} else {
																															F_sequence_close(m, v246, int32(3))
																															mBase = m.M
																															v299 = m.ExcPending
																															if v299 != 0 {
																																return int32(0)
																															} else {
																																v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																																if v301 != 0 {
																																	v302 = int32(0)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																	v305 = int32(1259)
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																	F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																	mBase = m.M
																																	v318 = m.ExcPending
																																	if v318 != 0 {
																																		return int32(0)
																																	} else {
																																		F_CommandCounterIncrement(m)
																																		mBase = m.M
																																		v320 = m.ExcPending
																																		if v320 != 0 {
																																			return int32(0)
																																		} else {
																																			v326 = int32(1)
																																			m.G0 = v15 + int32(304)
																																			return v326
																																		}
																																	}
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
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
												}
											}
										}
									}
								} else {
									v326 = v8
									m.G0 = v15 + int32(304)
									return v326
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
						v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, l0)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								if l4 != int32(8) {
									v50 = l5
								} else {
									v50 = int32(0)
								}
								if v50 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v337 = m.ExcPending
									if v337 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(619417), int32(0))
										mBase = m.M
										v341 = m.ExcPending
										if v341 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(493186), int32(193), int32(388317))
											mBase = m.M
											v346 = m.ExcPending
											if v346 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
									v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(38255), v15+int32(48))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
										v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(27939), v15+int32(32))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v71 = F_CreateTemplateTupleDesc(m, int32(3))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												F_TupleDescInitEntry(m, v71, int32(1), int32(431902), int32(26), int32(-1), int32(0))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_TupleDescInitEntry(m, v71, int32(2), int32(228934), int32(23), int32(-1), int32(0))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_TupleDescInitEntry(m, v71, int32(3), int32(499016), int32(17), int32(-1), int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
															v95 = int32(4)
															v97 = v71 + v94<<(uint(v95)%32)
															v98 = int32(112)
															*(*uint8)(unsafe.Add(mBase, uint32(v97)+304)) = uint8(v98)
															*(*uint8)(unsafe.Add(mBase, uint32(v97)+204)) = uint8(v98)
															*(*uint8)(unsafe.Add(mBase, uint32(v97)+104)) = uint8(v98)
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
															v107 = v71 + v104<<(uint(v95)%32)
															v108 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v107)+305)) = uint8(v108)
															*(*uint8)(unsafe.Add(mBase, uint32(v107)+205)) = uint8(v108)
															*(*uint8)(unsafe.Add(mBase, uint32(v107)+105)) = uint8(v108)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
															v120 = *(*int32)(unsafe.Add(mBase, _consts[249]))
															if v120 != 0 {
																v121 = int32(1)
																if v116 == v120 {
																	v128 = v121
																} else {
																	v124 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																	if v124 == v116 {
																		v128 = v121
																	} else {
																		v128 = int32(0)
																	}
																}
															} else {
																v128 = int32(0)
															}
															if v128 != 0 {
																v130 = *(*int32)(unsafe.Add(mBase, _consts[220]))
																v131 = v130
															} else {
																v131 = int32(99)
															}
															v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+117)))
															v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
															switch v134 - int32(83) {
															case 0, 22, 26, 31, 33:
																v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
																v140 = base.B2i32(v137 == int32(0))
															default:
																v140 = int32(0)
															}
															v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
															v144 = int32(0)
															v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
															v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
															v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+156))
															v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return int32(0)
															} else {
																v151 = int32(0)
																v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+118)))
																v155 = int32(1)
																v162 = F_heap_create_with_catalog(m, v15+int32(224), v131, v143, l1, v144, v144, v146, v149, v71, v151, int32(116), v154, v133&v155, v140, v151, l3, v151, v155, v155, l6, v151)
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return int32(0)
																	} else {
																		v167 = F_table_open(m, v162, int32(5))
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return int32(0)
																		} else {
																			v170 = F_palloc0(m, int32(144))
																			mBase = m.M
																			v171 = m.ExcPending
																			if v171 != 0 {
																				return int32(0)
																			} else {
																				v172 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v170)+76)) = v172
																				*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(562954248388610)
																				*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(8589934973)
																				v178 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v170)+128)) = v178
																				v180 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v170)+118)) = uint8(v180)
																				*(*uint16)(unsafe.Add(mBase, uint32(v170)+116)) = uint16(v180)
																				*(*int64)(unsafe.Add(mBase, uint32(v170)+132)) = int64(403)
																				*(*int32)(unsafe.Add(mBase, uint32(v170)+119)) = v178
																				*(*int64)(unsafe.Add(mBase, uint32(v170)+84)) = v172
																				*(*int64)(unsafe.Add(mBase, uint32(v170)+92)) = v172
																				*(*int32)(unsafe.Add(mBase, uint32(v170)+100)) = v178
																				v195 = *(*int32)(unsafe.Add(mBase, _consts[9]))
																				*(*int32)(unsafe.Add(mBase, uint32(v170)+140)) = v195
																				*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
																				*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v172
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v178
																				v203 = int32(431902)
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v203
																				v205 = int32(228934)
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v205
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v203
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v205
																				v220 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																					v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
																					v229 = int32(0)
																					v234 = int32(1)
																					v239 = F_index_create(m, v167, v15+int32(160), l2, v178, v178, v178, v170, v220, int32(403), v224, v15+int32(152), v15+int32(144), v229, v15+int32(140), v229, v229, v234, v229, v234, v234, v229)
																					mBase = m.M
																					v240 = m.ExcPending
																					if v240 != 0 {
																						return int32(0)
																					} else {
																						F_sequence_close(m, v167, int32(0))
																						mBase = m.M
																						v243 = m.ExcPending
																						if v243 != 0 {
																							return int32(0)
																						} else {
																							v246 = F_table_open(m, int32(1259), int32(3))
																							mBase = m.M
																							v247 = m.ExcPending
																							if v247 != 0 {
																								return int32(0)
																							} else {
																								v249 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																								if v249 != 0 {
																									v252 = F_SearchSysCacheCopy(m, int32(57), v19, int32(0))
																									mBase = m.M
																									v253 = m.ExcPending
																									if v253 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v252
																										if v252 == int32(0) {
																											F_errstart_cold(m, int32(21), int32(0))
																											mBase = m.M
																											v350 = m.ExcPending
																											if v350 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
																												F_errmsg_internal(m, int32(46032), v15+int32(16))
																												mBase = m.M
																												v356 = m.ExcPending
																												if v356 != 0 {
																													return int32(0)
																												} else {
																													F_errfinish(m, int32(493186), int32(342), int32(388317))
																													mBase = m.M
																													v361 = m.ExcPending
																													if v361 != 0 {
																														return int32(0)
																													} else {
																														base.Wasm_trap_unreachable()
																														for {
																														}
																													}
																												}
																											}
																										} else {
																											v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
																											v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
																											*(*int32)(unsafe.Add(mBase, uint32(v257+v258)+112)) = v162
																											F_CatalogTupleUpdate(m, v246, v252+int32(4), v252)
																											mBase = m.M
																											v264 = m.ExcPending
																											if v264 != 0 {
																												return int32(0)
																											} else {
																												v293 = v252
																												F_pfree(m, v293)
																												mBase = m.M
																												v296 = m.ExcPending
																												if v296 != 0 {
																													return int32(0)
																												} else {
																													F_sequence_close(m, v246, int32(3))
																													mBase = m.M
																													v299 = m.ExcPending
																													if v299 != 0 {
																														return int32(0)
																													} else {
																														v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																														if v301 != 0 {
																															v302 = int32(0)
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																															v305 = int32(1259)
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																															*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																															F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																															mBase = m.M
																															v318 = m.ExcPending
																															if v318 != 0 {
																																return int32(0)
																															} else {
																																F_CommandCounterIncrement(m)
																																mBase = m.M
																																v320 = m.ExcPending
																																if v320 != 0 {
																																	return int32(0)
																																} else {
																																	v326 = int32(1)
																																	m.G0 = v15 + int32(304)
																																	return v326
																																}
																															}
																														} else {
																															F_CommandCounterIncrement(m)
																															mBase = m.M
																															v320 = m.ExcPending
																															if v320 != 0 {
																																return int32(0)
																															} else {
																																v326 = int32(1)
																																m.G0 = v15 + int32(304)
																																return v326
																															}
																														}
																													}
																												}
																											}
																										}
																									}
																								} else {
																									F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v19)
																									mBase = m.M
																									v271 = m.ExcPending
																									if v271 != 0 {
																										return int32(0)
																									} else {
																										F_systable_inplace_update_begin(m, v246, int32(2662), v15-int32(-64), v15+int32(300), v15+int32(128))
																										mBase = m.M
																										v280 = m.ExcPending
																										if v280 != 0 {
																											return int32(0)
																										} else {
																											v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																											if v281 == int32(0) {
																												F_errstart_cold(m, int32(21), int32(0))
																												mBase = m.M
																												v365 = m.ExcPending
																												if v365 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
																													F_errmsg_internal(m, int32(46032), v15)
																													mBase = m.M
																													v369 = m.ExcPending
																													if v369 != 0 {
																														return int32(0)
																													} else {
																														F_errfinish(m, int32(493186), int32(362), int32(388317))
																														mBase = m.M
																														v374 = m.ExcPending
																														if v374 != 0 {
																															return int32(0)
																														} else {
																															base.Wasm_trap_unreachable()
																															for {
																															}
																														}
																													}
																												}
																											} else {
																												v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
																												v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+22)))
																												*(*int32)(unsafe.Add(mBase, uint32(v284+v285)+112)) = v162
																												v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
																												v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																												F_systable_inplace_update_finish(m, v288, v289)
																												mBase = m.M
																												v291 = m.ExcPending
																												if v291 != 0 {
																													return int32(0)
																												} else {
																													v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																													v293 = v292
																													F_pfree(m, v293)
																													mBase = m.M
																													v296 = m.ExcPending
																													if v296 != 0 {
																														return int32(0)
																													} else {
																														F_sequence_close(m, v246, int32(3))
																														mBase = m.M
																														v299 = m.ExcPending
																														if v299 != 0 {
																															return int32(0)
																														} else {
																															v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																															if v301 != 0 {
																																v302 = int32(0)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																																v305 = int32(1259)
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																																*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																																F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																																mBase = m.M
																																v318 = m.ExcPending
																																if v318 != 0 {
																																	return int32(0)
																																} else {
																																	F_CommandCounterIncrement(m)
																																	mBase = m.M
																																	v320 = m.ExcPending
																																	if v320 != 0 {
																																		return int32(0)
																																	} else {
																																		v326 = int32(1)
																																		m.G0 = v15 + int32(304)
																																		return v326
																																	}
																																}
																															} else {
																																F_CommandCounterIncrement(m)
																																mBase = m.M
																																v320 = m.ExcPending
																																if v320 != 0 {
																																	return int32(0)
																																} else {
																																	v326 = int32(1)
																																	m.G0 = v15 + int32(304)
																																	return v326
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
											}
										}
									}
								}
							} else {
								v326 = v8
								m.G0 = v15 + int32(304)
								return v326
							}
						}
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[261]))
			if v44 == int32(0) {
				v326 = v8
				m.G0 = v15 + int32(304)
				return v326
			} else {
				if l4 != int32(8) {
					v50 = l5
				} else {
					v50 = int32(0)
				}
				if v50 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v337 = m.ExcPending
					if v337 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(619417), int32(0))
						mBase = m.M
						v341 = m.ExcPending
						if v341 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493186), int32(193), int32(388317))
							mBase = m.M
							v346 = m.ExcPending
							if v346 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
					v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(38255), v15+int32(48))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
						v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(27939), v15+int32(32))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v71 = F_CreateTemplateTupleDesc(m, int32(3))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v71, int32(1), int32(431902), int32(26), int32(-1), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_TupleDescInitEntry(m, v71, int32(2), int32(228934), int32(23), int32(-1), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_TupleDescInitEntry(m, v71, int32(3), int32(499016), int32(17), int32(-1), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
											v95 = int32(4)
											v97 = v71 + v94<<(uint(v95)%32)
											v98 = int32(112)
											*(*uint8)(unsafe.Add(mBase, uint32(v97)+304)) = uint8(v98)
											*(*uint8)(unsafe.Add(mBase, uint32(v97)+204)) = uint8(v98)
											*(*uint8)(unsafe.Add(mBase, uint32(v97)+104)) = uint8(v98)
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
											v107 = v71 + v104<<(uint(v95)%32)
											v108 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v107)+305)) = uint8(v108)
											*(*uint8)(unsafe.Add(mBase, uint32(v107)+205)) = uint8(v108)
											*(*uint8)(unsafe.Add(mBase, uint32(v107)+105)) = uint8(v108)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
											v120 = *(*int32)(unsafe.Add(mBase, _consts[249]))
											if v120 != 0 {
												v121 = int32(1)
												if v116 == v120 {
													v128 = v121
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, _consts[220]))
													if v124 == v116 {
														v128 = v121
													} else {
														v128 = int32(0)
													}
												}
											} else {
												v128 = int32(0)
											}
											if v128 != 0 {
												v130 = *(*int32)(unsafe.Add(mBase, _consts[220]))
												v131 = v130
											} else {
												v131 = int32(99)
											}
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+117)))
											v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
											switch v134 - int32(83) {
											case 0, 22, 26, 31, 33:
												v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
												v140 = base.B2i32(v137 == int32(0))
											default:
												v140 = int32(0)
											}
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
											v144 = int32(0)
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+156))
											v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												v151 = int32(0)
												v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v153)+118)))
												v155 = int32(1)
												v162 = F_heap_create_with_catalog(m, v15+int32(224), v131, v143, l1, v144, v144, v146, v149, v71, v151, int32(116), v154, v133&v155, v140, v151, l3, v151, v155, v155, l6, v151)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int32(0)
													} else {
														v167 = F_table_open(m, v162, int32(5))
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															v170 = F_palloc0(m, int32(144))
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																v172 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v170)+76)) = v172
																*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(562954248388610)
																*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(8589934973)
																v178 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v170)+128)) = v178
																v180 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v170)+118)) = uint8(v180)
																*(*uint16)(unsafe.Add(mBase, uint32(v170)+116)) = uint16(v180)
																*(*int64)(unsafe.Add(mBase, uint32(v170)+132)) = int64(403)
																*(*int32)(unsafe.Add(mBase, uint32(v170)+119)) = v178
																*(*int64)(unsafe.Add(mBase, uint32(v170)+84)) = v172
																*(*int64)(unsafe.Add(mBase, uint32(v170)+92)) = v172
																*(*int32)(unsafe.Add(mBase, uint32(v170)+100)) = v178
																v195 = *(*int32)(unsafe.Add(mBase, _consts[9]))
																*(*int32)(unsafe.Add(mBase, uint32(v170)+140)) = v195
																*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
																*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v172
																*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v178
																v203 = int32(431902)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v203
																v205 = int32(228934)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v205
																*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v203
																*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v205
																v220 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
																	return int32(0)
																} else {
																	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
																	v229 = int32(0)
																	v234 = int32(1)
																	v239 = F_index_create(m, v167, v15+int32(160), l2, v178, v178, v178, v170, v220, int32(403), v224, v15+int32(152), v15+int32(144), v229, v15+int32(140), v229, v229, v234, v229, v234, v234, v229)
																	mBase = m.M
																	v240 = m.ExcPending
																	if v240 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v167, int32(0))
																		mBase = m.M
																		v243 = m.ExcPending
																		if v243 != 0 {
																			return int32(0)
																		} else {
																			v246 = F_table_open(m, int32(1259), int32(3))
																			mBase = m.M
																			v247 = m.ExcPending
																			if v247 != 0 {
																				return int32(0)
																			} else {
																				v249 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																				if v249 != 0 {
																					v252 = F_SearchSysCacheCopy(m, int32(57), v19, int32(0))
																					mBase = m.M
																					v253 = m.ExcPending
																					if v253 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v252
																						if v252 == int32(0) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v350 = m.ExcPending
																							if v350 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
																								F_errmsg_internal(m, int32(46032), v15+int32(16))
																								mBase = m.M
																								v356 = m.ExcPending
																								if v356 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(493186), int32(342), int32(388317))
																									mBase = m.M
																									v361 = m.ExcPending
																									if v361 != 0 {
																										return int32(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
																							v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
																							*(*int32)(unsafe.Add(mBase, uint32(v257+v258)+112)) = v162
																							F_CatalogTupleUpdate(m, v246, v252+int32(4), v252)
																							mBase = m.M
																							v264 = m.ExcPending
																							if v264 != 0 {
																								return int32(0)
																							} else {
																								v293 = v252
																								F_pfree(m, v293)
																								mBase = m.M
																								v296 = m.ExcPending
																								if v296 != 0 {
																									return int32(0)
																								} else {
																									F_sequence_close(m, v246, int32(3))
																									mBase = m.M
																									v299 = m.ExcPending
																									if v299 != 0 {
																										return int32(0)
																									} else {
																										v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																										if v301 != 0 {
																											v302 = int32(0)
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																											v305 = int32(1259)
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																											*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																											F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																											mBase = m.M
																											v318 = m.ExcPending
																											if v318 != 0 {
																												return int32(0)
																											} else {
																												F_CommandCounterIncrement(m)
																												mBase = m.M
																												v320 = m.ExcPending
																												if v320 != 0 {
																													return int32(0)
																												} else {
																													v326 = int32(1)
																													m.G0 = v15 + int32(304)
																													return v326
																												}
																											}
																										} else {
																											F_CommandCounterIncrement(m)
																											mBase = m.M
																											v320 = m.ExcPending
																											if v320 != 0 {
																												return int32(0)
																											} else {
																												v326 = int32(1)
																												m.G0 = v15 + int32(304)
																												return v326
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				} else {
																					F_ScanKeyInit(m, v15-int32(-64), int32(1), int32(3), int32(184), v19)
																					mBase = m.M
																					v271 = m.ExcPending
																					if v271 != 0 {
																						return int32(0)
																					} else {
																						F_systable_inplace_update_begin(m, v246, int32(2662), v15-int32(-64), v15+int32(300), v15+int32(128))
																						mBase = m.M
																						v280 = m.ExcPending
																						if v280 != 0 {
																							return int32(0)
																						} else {
																							v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																							if v281 == int32(0) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v365 = m.ExcPending
																								if v365 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
																									F_errmsg_internal(m, int32(46032), v15)
																									mBase = m.M
																									v369 = m.ExcPending
																									if v369 != 0 {
																										return int32(0)
																									} else {
																										F_errfinish(m, int32(493186), int32(362), int32(388317))
																										mBase = m.M
																										v374 = m.ExcPending
																										if v374 != 0 {
																											return int32(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
																								v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+22)))
																								*(*int32)(unsafe.Add(mBase, uint32(v284+v285)+112)) = v162
																								v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
																								v289 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																								F_systable_inplace_update_finish(m, v288, v289)
																								mBase = m.M
																								v291 = m.ExcPending
																								if v291 != 0 {
																									return int32(0)
																								} else {
																									v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
																									v293 = v292
																									F_pfree(m, v293)
																									mBase = m.M
																									v296 = m.ExcPending
																									if v296 != 0 {
																										return int32(0)
																									} else {
																										F_sequence_close(m, v246, int32(3))
																										mBase = m.M
																										v299 = m.ExcPending
																										if v299 != 0 {
																											return int32(0)
																										} else {
																											v301 = *(*int32)(unsafe.Add(mBase, _consts[231]))
																											if v301 != 0 {
																												v302 = int32(0)
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v302
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
																												v305 = int32(1259)
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v305
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v302
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v162
																												*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v305
																												F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
																												mBase = m.M
																												v318 = m.ExcPending
																												if v318 != 0 {
																													return int32(0)
																												} else {
																													F_CommandCounterIncrement(m)
																													mBase = m.M
																													v320 = m.ExcPending
																													if v320 != 0 {
																														return int32(0)
																													} else {
																														v326 = int32(1)
																														m.G0 = v15 + int32(304)
																														return v326
																													}
																												}
																											} else {
																												F_CommandCounterIncrement(m)
																												mBase = m.M
																												v320 = m.ExcPending
																												if v320 != 0 {
																													return int32(0)
																												} else {
																													v326 = int32(1)
																													m.G0 = v15 + int32(304)
																													return v326
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
							}
						}
					}
				}
			}
		}
	}
}
func F_toast_fetch_datum_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(149253), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(487850), int32(405), int32(414171))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v12 != int32(18) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(149253), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487850), int32(405), int32(414171))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+14))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+10))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+6))
			v19 = v17 & int32(1073741823)
			v21 = base.B2i32(base.Ui32(l1) < base.Ui32(v19))
			if base.Ui32(l1) < base.Ui32(v19) {
				v22 = l1
			} else {
				v22 = int32(0)
			}
			v23 = v19 - v22
			if base.Ui32(l1) < base.Ui32(v19) {
				v25 = l2
			} else {
				v25 = int32(0)
			}
			if int32(0) < v25 {
				v30 = v25 + int32(4)
			} else {
				v30 = v25
			}
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v34 = base.B2i32(base.Ui32(v19) < base.Ui32(v31-int32(4)))
			if base.Ui32(v19) < base.Ui32(v31-int32(4)) {
				v35 = v30
			} else {
				v35 = v25
			}
			if v19 < v35+v22 {
				v38 = v23
			} else {
				v38 = v35
			}
			if v35 < int32(0) {
				v41 = v23
			} else {
				v41 = v38
			}
			v43 = v41 + int32(4)
			v44 = F_palloc(m, v43)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v48 = int32(2)
				v49 = v43 << (uint(v48) % 32)
				if base.Ui32(v19) < base.Ui32(v31-int32(4)) {
					v52 = v49 | v48
				} else {
					v52 = v49
				}
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v52
				if v41 != 0 {
					v55 = F_table_open(m, v15, int32(1))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+188))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+160))
						m.T0[v58].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v55, v16, v19, v22, v41, v44)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v55, int32(1))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								return v44
							}
						}
					}
				} else {
					return v44
				}
			}
		}
	}
}
func F_toast_tuple_try_compression(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v6 + l1<<(uint(int32(2))%32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v11 + l1*int32(12)
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+9)))
	v16 = F_toast_compress_datum(m, v10, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
		if v16 != 0 {
			if v18&int32(2) != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				F_pfree(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
					v26 = int32(2)
					v27 = v25 | v26
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v27)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(base.Ui32(v30) >> (uint(v26) % 32))
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
					v36 = v34 | int32(10)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v36)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
				v26 = int32(2)
				v27 = v25 | v26
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v27)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(base.Ui32(v30) >> (uint(v26) % 32))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				v36 = v34 | int32(10)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v36)
				return
			}
		} else {
			v39 = v18 | int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v39)
			return
		}
	}
}
