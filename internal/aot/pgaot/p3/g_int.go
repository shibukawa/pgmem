package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_decompress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v176 int64
	_ = v176
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v2 {
		v31 = v2
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
		if v18 == int32(0) {
			v31 = v2
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			if v21 != int32(7) {
				v31 = v2
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				if v24 != int32(17) {
					v31 = v2
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
					v31 = v27 ^ int32(1)
				}
			}
		}
	}
	if v31&int32(1) != 0 {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = F_get_fn_opclass_options(m, v34)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
			v42 = v39 << (uint(int32(1)) % 32)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v44 = F_pg_detoast_datum(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				if v46 != 0 {
					v47 = F_array_contains_nulls(m, v44)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v306 = m.ExcPending
							if v306 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v309 = m.ExcPending
								if v309 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_g_int_decompress_0), int32(0))
									mBase = m.M
									v313 = m.ExcPending
									if v313 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(304), int32(_a_F_g_int_decompress_2))
										mBase = m.M
										v318 = m.ExcPending
										if v318 != 0 {
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
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v51 = v44 + int32(16)
							v52 = F_ArrayGetNItemsSafe(m, v49, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									if v44 != v56 {
										v281 = v44
										v291 = F_palloc(m, int32(16))
										mBase = m.M
										v292 = m.ExcPending
										if v292 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
											v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
											v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
											v299 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
											*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
											return v291
										}
									} else {
										return v12
									}
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v60 = F_ArrayGetNItemsSafe(m, v59, v51)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 < v42 {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											if v44 != v63 {
												v281 = v44
												v291 = F_palloc(m, int32(16))
												mBase = m.M
												v292 = m.ExcPending
												if v292 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
													v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
													v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
													v299 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
													*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
													return v291
												}
											} else {
												return v12
											}
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
											if v66 != 0 {
												v74 = v66
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
												v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											v75 = v74 + v44
											v77 = int32(0)
											if v60 <= v77 {
												v176 = int64(0)
											} else {
												v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75)+4)))
												v85 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75))))
												v88 = v84 - v85 + int64(1)
												if base.Ui32(v60) < base.Ui32(int32(3)) {
													v176 = v88
												} else {
													v94 = int32(base.Ui32(v60-int32(3)) >> (uint(int32(1)) % 32))
													if v94 == int32(0) {
														v153 = int32(2)
														v154 = v88
														v162 = v75 + v153<<(uint(int32(2))%32)
														v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
														if v163 == v166 {
															v176 = v154
														} else {
															v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
															v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
														}
													} else {
														v98 = int32(1)
														v99 = v94 + v98
														v106 = int32(2)
														v107 = v88
														v112 = v77
														for {
															v115 = v75 + v106<<(uint(int32(2))%32)
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
															if v116 != v119 {
																v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+4)))
																v127 = v107 + v121 - base.I64_extend_i32_s(v116) + int64(1)
															} else {
																v127 = v107
															}
															v128 = int32(2)
															v132 = v75 + (v106+v128)<<(uint(v128)%32)
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
															v136 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
															if v133 != v136 {
																v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+4)))
																v144 = v127 + v138 - base.I64_extend_i32_s(v133) + int64(1)
															} else {
																v144 = v127
															}
															v146 = v106 + int32(4)
															v148 = v112 + int32(2)
															if v148 != v99&int32(-2) {
																v106 = v146
																v107 = v144
																v112 = v148
																continue
															} else {
																break
															}
															break
														}
														if v99&v98 == int32(0) {
															v176 = v144
														} else {
															v153 = v146
															v154 = v144
															v162 = v75 + v153<<(uint(int32(2))%32)
															v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
															v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
															if v163 == v166 {
																v176 = v154
															} else {
																v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
																v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
															}
														}
													}
												}
											}
											if base.Ui64(v176-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
												v188 = int32(-1)
											} else {
												v188 = base.I32_wrap_i64(v176)
											}
											if base.Ui32(int32(134217725)) <= base.Ui32(v188) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v322 = m.ExcPending
												if v322 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(261))
													mBase = m.M
													v325 = m.ExcPending
													if v325 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
														mBase = m.M
														v329 = m.ExcPending
														if v329 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(338), int32(_a_F_g_int_decompress_2))
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
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
												v191 = F_new_intArrayType(m, v188)
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
													if v193 == int32(0) {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
														v203 = (v196<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													} else {
														v203 = v193
													}
													if int32(0) < v60 {
														v207 = v203 + v191
														v212 = v2
														for {
															v220 = v75 + v212<<(uint(int32(2))%32)
															v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
															v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
															if v221 <= v222 {
																v225 = v207
																v233 = v222
																v235 = base.I64_extend_i32_s(v221)
																for {
																	if v212 != 0 {
																		v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v225-int32(4)))))
																		if v235 == v238 {
																			v244 = v225
																			v245 = v233
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																			v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																			v244 = v225 + int32(4)
																			v245 = v243
																		}
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																		v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																		v244 = v225 + int32(4)
																		v245 = v243
																	}
																	if v235 < base.I64_extend_i32_s(v245) {
																		v225 = v244
																		v233 = v245
																		v235 = v235 + int64(1)
																		continue
																	} else {
																		break
																	}
																	break
																}
																v250 = v244
															} else {
																v250 = v207
															}
															v262 = v212 + int32(2)
															if v262 < v60 {
																v207 = v250
																v212 = v262
																continue
															} else {
																break
															}
															break
														}
													} else {
													}
													v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
													if v275 != v44 {
														F_pfree(m, v44)
														mBase = m.M
														v278 = m.ExcPending
														if v278 != 0 {
															return int32(0)
														} else {
															v281 = v191
															v291 = F_palloc(m, int32(16))
															mBase = m.M
															v292 = m.ExcPending
															if v292 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
																v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
																v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
																v299 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
																*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
																return v291
															}
														}
													} else {
														v281 = v191
														v291 = F_palloc(m, int32(16))
														mBase = m.M
														v292 = m.ExcPending
														if v292 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
															v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
															v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
															v299 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
															*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
															return v291
														}
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
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
					v51 = v44 + int32(16)
					v52 = F_ArrayGetNItemsSafe(m, v49, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							if v44 != v56 {
								v281 = v44
								v291 = F_palloc(m, int32(16))
								mBase = m.M
								v292 = m.ExcPending
								if v292 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
									v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
									v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
									v299 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
									*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
									return v291
								}
							} else {
								return v12
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v60 = F_ArrayGetNItemsSafe(m, v59, v51)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								if v60 < v42 {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									if v44 != v63 {
										v281 = v44
										v291 = F_palloc(m, int32(16))
										mBase = m.M
										v292 = m.ExcPending
										if v292 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
											v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
											v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
											v299 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
											*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
											return v291
										}
									} else {
										return v12
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
									if v66 != 0 {
										v74 = v66
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
										v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									v75 = v74 + v44
									v77 = int32(0)
									if v60 <= v77 {
										v176 = int64(0)
									} else {
										v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75)+4)))
										v85 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75))))
										v88 = v84 - v85 + int64(1)
										if base.Ui32(v60) < base.Ui32(int32(3)) {
											v176 = v88
										} else {
											v94 = int32(base.Ui32(v60-int32(3)) >> (uint(int32(1)) % 32))
											if v94 == int32(0) {
												v153 = int32(2)
												v154 = v88
												v162 = v75 + v153<<(uint(int32(2))%32)
												v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
												v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
												if v163 == v166 {
													v176 = v154
												} else {
													v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
													v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
												}
											} else {
												v98 = int32(1)
												v99 = v94 + v98
												v106 = int32(2)
												v107 = v88
												v112 = v77
												for {
													v115 = v75 + v106<<(uint(int32(2))%32)
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
													if v116 != v119 {
														v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+4)))
														v127 = v107 + v121 - base.I64_extend_i32_s(v116) + int64(1)
													} else {
														v127 = v107
													}
													v128 = int32(2)
													v132 = v75 + (v106+v128)<<(uint(v128)%32)
													v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
													v136 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
													if v133 != v136 {
														v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+4)))
														v144 = v127 + v138 - base.I64_extend_i32_s(v133) + int64(1)
													} else {
														v144 = v127
													}
													v146 = v106 + int32(4)
													v148 = v112 + int32(2)
													if v148 != v99&int32(-2) {
														v106 = v146
														v107 = v144
														v112 = v148
														continue
													} else {
														break
													}
													break
												}
												if v99&v98 == int32(0) {
													v176 = v144
												} else {
													v153 = v146
													v154 = v144
													v162 = v75 + v153<<(uint(int32(2))%32)
													v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
													if v163 == v166 {
														v176 = v154
													} else {
														v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
														v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
													}
												}
											}
										}
									}
									if base.Ui64(v176-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
										v188 = int32(-1)
									} else {
										v188 = base.I32_wrap_i64(v176)
									}
									if base.Ui32(int32(134217725)) <= base.Ui32(v188) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v322 = m.ExcPending
										if v322 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(261))
											mBase = m.M
											v325 = m.ExcPending
											if v325 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(338), int32(_a_F_g_int_decompress_2))
													mBase = m.M
													v334 = m.ExcPending
													if v334 != 0 {
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
										v191 = F_new_intArrayType(m, v188)
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int32(0)
										} else {
											v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
											if v193 == int32(0) {
												v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
												v203 = (v196<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v203 = v193
											}
											if int32(0) < v60 {
												v207 = v203 + v191
												v212 = v2
												for {
													v220 = v75 + v212<<(uint(int32(2))%32)
													v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
													if v221 <= v222 {
														v225 = v207
														v233 = v222
														v235 = base.I64_extend_i32_s(v221)
														for {
															if v212 != 0 {
																v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v225-int32(4)))))
																if v235 == v238 {
																	v244 = v225
																	v245 = v233
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																	v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																	v244 = v225 + int32(4)
																	v245 = v243
																}
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																v244 = v225 + int32(4)
																v245 = v243
															}
															if v235 < base.I64_extend_i32_s(v245) {
																v225 = v244
																v233 = v245
																v235 = v235 + int64(1)
																continue
															} else {
																break
															}
															break
														}
														v250 = v244
													} else {
														v250 = v207
													}
													v262 = v212 + int32(2)
													if v262 < v60 {
														v207 = v250
														v212 = v262
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											if v275 != v44 {
												F_pfree(m, v44)
												mBase = m.M
												v278 = m.ExcPending
												if v278 != 0 {
													return int32(0)
												} else {
													v281 = v191
													v291 = F_palloc(m, int32(16))
													mBase = m.M
													v292 = m.ExcPending
													if v292 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
														v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
														v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
														v299 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
														*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
														return v291
													}
												}
											} else {
												v281 = v191
												v291 = F_palloc(m, int32(16))
												mBase = m.M
												v292 = m.ExcPending
												if v292 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
													v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
													v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
													v299 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
													*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
													return v291
												}
											}
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
		v42 = int32(200)
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v44 = F_pg_detoast_datum(m, v43)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
			if v46 != 0 {
				v47 = F_array_contains_nulls(m, v44)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v47 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v306 = m.ExcPending
						if v306 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v309 = m.ExcPending
							if v309 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_g_int_decompress_0), int32(0))
								mBase = m.M
								v313 = m.ExcPending
								if v313 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(304), int32(_a_F_g_int_decompress_2))
									mBase = m.M
									v318 = m.ExcPending
									if v318 != 0 {
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
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v51 = v44 + int32(16)
						v52 = F_ArrayGetNItemsSafe(m, v49, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								if v44 != v56 {
									v281 = v44
									v291 = F_palloc(m, int32(16))
									mBase = m.M
									v292 = m.ExcPending
									if v292 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
										v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
										v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
										v299 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
										*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
										return v291
									}
								} else {
									return v12
								}
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
								v60 = F_ArrayGetNItemsSafe(m, v59, v51)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									if v60 < v42 {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										if v44 != v63 {
											v281 = v44
											v291 = F_palloc(m, int32(16))
											mBase = m.M
											v292 = m.ExcPending
											if v292 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
												v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
												v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
												v299 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
												*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
												return v291
											}
										} else {
											return v12
										}
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
										if v66 != 0 {
											v74 = v66
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
											v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v75 = v74 + v44
										v77 = int32(0)
										if v60 <= v77 {
											v176 = int64(0)
										} else {
											v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75)+4)))
											v85 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75))))
											v88 = v84 - v85 + int64(1)
											if base.Ui32(v60) < base.Ui32(int32(3)) {
												v176 = v88
											} else {
												v94 = int32(base.Ui32(v60-int32(3)) >> (uint(int32(1)) % 32))
												if v94 == int32(0) {
													v153 = int32(2)
													v154 = v88
													v162 = v75 + v153<<(uint(int32(2))%32)
													v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
													if v163 == v166 {
														v176 = v154
													} else {
														v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
														v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
													}
												} else {
													v98 = int32(1)
													v99 = v94 + v98
													v106 = int32(2)
													v107 = v88
													v112 = v77
													for {
														v115 = v75 + v106<<(uint(int32(2))%32)
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
														if v116 != v119 {
															v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+4)))
															v127 = v107 + v121 - base.I64_extend_i32_s(v116) + int64(1)
														} else {
															v127 = v107
														}
														v128 = int32(2)
														v132 = v75 + (v106+v128)<<(uint(v128)%32)
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
														v136 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
														if v133 != v136 {
															v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+4)))
															v144 = v127 + v138 - base.I64_extend_i32_s(v133) + int64(1)
														} else {
															v144 = v127
														}
														v146 = v106 + int32(4)
														v148 = v112 + int32(2)
														if v148 != v99&int32(-2) {
															v106 = v146
															v107 = v144
															v112 = v148
															continue
														} else {
															break
														}
														break
													}
													if v99&v98 == int32(0) {
														v176 = v144
													} else {
														v153 = v146
														v154 = v144
														v162 = v75 + v153<<(uint(int32(2))%32)
														v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
														if v163 == v166 {
															v176 = v154
														} else {
															v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
															v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
														}
													}
												}
											}
										}
										if base.Ui64(v176-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
											v188 = int32(-1)
										} else {
											v188 = base.I32_wrap_i64(v176)
										}
										if base.Ui32(int32(134217725)) <= base.Ui32(v188) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v322 = m.ExcPending
											if v322 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(261))
												mBase = m.M
												v325 = m.ExcPending
												if v325 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(338), int32(_a_F_g_int_decompress_2))
														mBase = m.M
														v334 = m.ExcPending
														if v334 != 0 {
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
											v191 = F_new_intArrayType(m, v188)
											mBase = m.M
											v192 = m.ExcPending
											if v192 != 0 {
												return int32(0)
											} else {
												v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
												if v193 == int32(0) {
													v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
													v203 = (v196<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v203 = v193
												}
												if int32(0) < v60 {
													v207 = v203 + v191
													v212 = v2
													for {
														v220 = v75 + v212<<(uint(int32(2))%32)
														v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
														v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
														if v221 <= v222 {
															v225 = v207
															v233 = v222
															v235 = base.I64_extend_i32_s(v221)
															for {
																if v212 != 0 {
																	v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v225-int32(4)))))
																	if v235 == v238 {
																		v244 = v225
																		v245 = v233
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																		v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																		v244 = v225 + int32(4)
																		v245 = v243
																	}
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																	v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																	v244 = v225 + int32(4)
																	v245 = v243
																}
																if v235 < base.I64_extend_i32_s(v245) {
																	v225 = v244
																	v233 = v245
																	v235 = v235 + int64(1)
																	continue
																} else {
																	break
																}
																break
															}
															v250 = v244
														} else {
															v250 = v207
														}
														v262 = v212 + int32(2)
														if v262 < v60 {
															v207 = v250
															v212 = v262
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
												if v275 != v44 {
													F_pfree(m, v44)
													mBase = m.M
													v278 = m.ExcPending
													if v278 != 0 {
														return int32(0)
													} else {
														v281 = v191
														v291 = F_palloc(m, int32(16))
														mBase = m.M
														v292 = m.ExcPending
														if v292 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
															v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
															v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
															v299 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
															*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
															return v291
														}
													}
												} else {
													v281 = v191
													v291 = F_palloc(m, int32(16))
													mBase = m.M
													v292 = m.ExcPending
													if v292 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
														v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
														v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
														v299 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
														*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
														return v291
													}
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
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v51 = v44 + int32(16)
				v52 = F_ArrayGetNItemsSafe(m, v49, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v44 != v56 {
							v281 = v44
							v291 = F_palloc(m, int32(16))
							mBase = m.M
							v292 = m.ExcPending
							if v292 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
								v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
								v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
								v299 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
								*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
								return v291
							}
						} else {
							return v12
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v60 = F_ArrayGetNItemsSafe(m, v59, v51)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							if v60 < v42 {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								if v44 != v63 {
									v281 = v44
									v291 = F_palloc(m, int32(16))
									mBase = m.M
									v292 = m.ExcPending
									if v292 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
										v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
										v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
										v299 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
										*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
										return v291
									}
								} else {
									return v12
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
								if v66 != 0 {
									v74 = v66
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v74 = (v67<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v75 = v74 + v44
								v77 = int32(0)
								if v60 <= v77 {
									v176 = int64(0)
								} else {
									v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75)+4)))
									v85 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75))))
									v88 = v84 - v85 + int64(1)
									if base.Ui32(v60) < base.Ui32(int32(3)) {
										v176 = v88
									} else {
										v94 = int32(base.Ui32(v60-int32(3)) >> (uint(int32(1)) % 32))
										if v94 == int32(0) {
											v153 = int32(2)
											v154 = v88
											v162 = v75 + v153<<(uint(int32(2))%32)
											v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
											if v163 == v166 {
												v176 = v154
											} else {
												v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
												v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
											}
										} else {
											v98 = int32(1)
											v99 = v94 + v98
											v106 = int32(2)
											v107 = v88
											v112 = v77
											for {
												v115 = v75 + v106<<(uint(int32(2))%32)
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
												if v116 != v119 {
													v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+4)))
													v127 = v107 + v121 - base.I64_extend_i32_s(v116) + int64(1)
												} else {
													v127 = v107
												}
												v128 = int32(2)
												v132 = v75 + (v106+v128)<<(uint(v128)%32)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
												v136 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
												if v133 != v136 {
													v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+4)))
													v144 = v127 + v138 - base.I64_extend_i32_s(v133) + int64(1)
												} else {
													v144 = v127
												}
												v146 = v106 + int32(4)
												v148 = v112 + int32(2)
												if v148 != v99&int32(-2) {
													v106 = v146
													v107 = v144
													v112 = v148
													continue
												} else {
													break
												}
												break
											}
											if v99&v98 == int32(0) {
												v176 = v144
											} else {
												v153 = v146
												v154 = v144
												v162 = v75 + v153<<(uint(int32(2))%32)
												v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
												v166 = *(*int32)(unsafe.Add(mBase, uint32(v162-int32(4))))
												if v163 == v166 {
													v176 = v154
												} else {
													v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v162)+4)))
													v176 = v154 + v168 - base.I64_extend_i32_s(v163) + int64(1)
												}
											}
										}
									}
								}
								if base.Ui64(v176-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
									v188 = int32(-1)
								} else {
									v188 = base.I32_wrap_i64(v176)
								}
								if base.Ui32(int32(134217725)) <= base.Ui32(v188) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v322 = m.ExcPending
									if v322 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v325 = m.ExcPending
										if v325 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
											mBase = m.M
											v329 = m.ExcPending
											if v329 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(338), int32(_a_F_g_int_decompress_2))
												mBase = m.M
												v334 = m.ExcPending
												if v334 != 0 {
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
									v191 = F_new_intArrayType(m, v188)
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
										if v193 == int32(0) {
											v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
											v203 = (v196<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v203 = v193
										}
										if int32(0) < v60 {
											v207 = v203 + v191
											v212 = v2
											for {
												v220 = v75 + v212<<(uint(int32(2))%32)
												v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
												if v221 <= v222 {
													v225 = v207
													v233 = v222
													v235 = base.I64_extend_i32_s(v221)
													for {
														if v212 != 0 {
															v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v225-int32(4)))))
															if v235 == v238 {
																v244 = v225
																v245 = v233
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
																v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
																v244 = v225 + int32(4)
																v245 = v243
															}
														} else {
															*(*uint32)(unsafe.Add(mBase, uint32(v225))) = uint32(v235)
															v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
															v244 = v225 + int32(4)
															v245 = v243
														}
														if v235 < base.I64_extend_i32_s(v245) {
															v225 = v244
															v233 = v245
															v235 = v235 + int64(1)
															continue
														} else {
															break
														}
														break
													}
													v250 = v244
												} else {
													v250 = v207
												}
												v262 = v212 + int32(2)
												if v262 < v60 {
													v207 = v250
													v212 = v262
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v275 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										if v275 != v44 {
											F_pfree(m, v44)
											mBase = m.M
											v278 = m.ExcPending
											if v278 != 0 {
												return int32(0)
											} else {
												v281 = v191
												v291 = F_palloc(m, int32(16))
												mBase = m.M
												v292 = m.ExcPending
												if v292 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
													v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
													v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
													v299 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
													*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
													return v291
												}
											}
										} else {
											v281 = v191
											v291 = F_palloc(m, int32(16))
											mBase = m.M
											v292 = m.ExcPending
											if v292 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v291))) = v281
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v294
												v296 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v296
												v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
												v299 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v291)+14)) = uint8(v299)
												*(*uint16)(unsafe.Add(mBase, uint32(v291)+12)) = uint16(v298)
												return v291
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_g_int_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v2 < v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L37
	}
L2:
	;
	v24 = v2
	v25 = v2
	goto L5
L3:
	;
	v54 = v2
	goto L4
L4:
	;
	v61 = F_new_intArrayType(m, v54)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L15
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(4)+v24<<(uint(int32(4))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v54 = v46
	goto L4
L7:
	;
	v37 = F_array_contains_nulls(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v44 = F_ArrayGetNItemsSafe(m, v41, v35+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v46 = v44 + v25
	v48 = v24 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v48 < v49 {
		v24 = v48
		v25 = v46
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v63 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = (v66<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L17:
	;
	v73 = v63
	goto L18
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if int32(0) < v74 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v83 = int32(0)
	v84 = v73 + v61
	goto L22
L20:
	;
	goto L21
L21:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v133 = F_ArrayGetNItemsSafe(m, v130, v61+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L32
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(4)+v83<<(uint(int32(4))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v98 = F_ArrayGetNItemsSafe(m, v95, v94+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v100 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L27
L26:
	;
	v110 = v100
	goto L27
L27:
	;
	v112 = v98 << (uint(int32(2)) % 32)
	if v112 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	base.MemoryCopy(m, v84, v94+v110, v112)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v117 = v83 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v117 < v118 {
		v83 = v117
		v84 = v84 + v112
		goto L22
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v137 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v145 = v137
	goto L35
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v145 = (v138<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L35
L35:
	;
	F_isort(m, v145+v61, v133, v13+int32(15))
	mBase = m.M
	v150 = F__int_unique(m, v61)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(base.Ui32(v152) >> (uint(int32(2)) % 32))
	m.G0 = v13 + int32(16)
	return v150
L37:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_g_int_union_0), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_g_int_union_1), int32(136), int32(_a_F_g_int_union_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
