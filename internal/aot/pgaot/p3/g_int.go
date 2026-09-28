package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_decompress(m *base.Module, l0 int32) int64 {
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v178 int64
	_ = v178
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v240 int64
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)))
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
			return int64(0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
			v42 = v39 << (uint(int32(1)) % 32)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v44 = F_pg_detoast_datum(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				if v46 != 0 {
					v47 = F_array_contains_nulls(m, v44)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						if v47 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v310 = m.ExcPending
							if v310 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v313 = m.ExcPending
								if v313 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_g_int_decompress_0), int32(0))
									mBase = m.M
									v317 = m.ExcPending
									if v317 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(305), int32(_a_F_g_int_decompress_2))
										mBase = m.M
										v322 = m.ExcPending
										if v322 != 0 {
											return int64(0)
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
								return int64(0)
							} else {
								if v52 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									if v44 != v56 {
										v283 = v44
										v293 = F_palloc(m, int32(24))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
											v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
											v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
											v302 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
											*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
											return base.I64_extend_i32_u(v293)
										}
									} else {
										return base.I64_extend_i32_u(v12)
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v61 = F_ArrayGetNItemsSafe(m, v60, v51)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										if v61 < v42 {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											if v44 != v64 {
												v283 = v44
												v293 = F_palloc(m, int32(24))
												mBase = m.M
												v294 = m.ExcPending
												if v294 != 0 {
													return int64(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
													v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
													v302 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
													*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
													return base.I64_extend_i32_u(v293)
												}
											} else {
												return base.I64_extend_i32_u(v12)
											}
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
											if v68 != 0 {
												v76 = v68
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
												v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											v77 = v76 + v44
											v79 = int32(0)
											if v61 <= v79 {
												v178 = int64(0)
											} else {
												v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
												v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77))))
												v90 = v86 - v87 + int64(1)
												if base.Ui32(v61) < base.Ui32(int32(3)) {
													v178 = v90
												} else {
													v96 = int32(base.Ui32(v61-int32(3)) >> (uint(int32(1)) % 32))
													if v96 == int32(0) {
														v155 = int32(2)
														v156 = v90
														v164 = v77 + v155<<(uint(int32(2))%32)
														v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
														if v165 == v168 {
															v178 = v156
														} else {
															v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
															v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
														}
													} else {
														v100 = int32(1)
														v101 = v96 + v100
														v108 = int32(2)
														v109 = v90
														v114 = v79
														for {
															v117 = v77 + v108<<(uint(int32(2))%32)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
															if v118 != v121 {
																v123 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
																v129 = v109 + v123 - base.I64_extend_i32_s(v118) + int64(1)
															} else {
																v129 = v109
															}
															v130 = int32(2)
															v134 = v77 + (v108+v130)<<(uint(v130)%32)
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v134-int32(4))))
															if v135 != v138 {
																v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v134)+4)))
																v146 = v129 + v140 - base.I64_extend_i32_s(v135) + int64(1)
															} else {
																v146 = v129
															}
															v148 = v108 + int32(4)
															v150 = v114 + int32(2)
															if v150 != v101&int32(-2) {
																v108 = v148
																v109 = v146
																v114 = v150
																continue
															} else {
																break
															}
															break
														}
														if v101&v100 == int32(0) {
															v178 = v146
														} else {
															v155 = v148
															v156 = v146
															v164 = v77 + v155<<(uint(int32(2))%32)
															v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
															v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
															if v165 == v168 {
																v178 = v156
															} else {
																v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
																v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
															}
														}
													}
												}
											}
											if base.Ui64(v178-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
												v190 = int32(-1)
											} else {
												v190 = base.I32_wrap_i64(v178)
											}
											if base.Ui32(int32(67108864)) <= base.Ui32(v190) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v326 = m.ExcPending
												if v326 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(261))
													mBase = m.M
													v329 = m.ExcPending
													if v329 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
														mBase = m.M
														v333 = m.ExcPending
														if v333 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(339), int32(_a_F_g_int_decompress_2))
															mBase = m.M
															v338 = m.ExcPending
															if v338 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v193 = F_new_intArrayType(m, v190)
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int64(0)
												} else {
													v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
													if v195 == int32(0) {
														v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
														v205 = (v198<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													} else {
														v205 = v195
													}
													if int32(0) < v61 {
														v209 = v205 + v193
														v214 = v2
														for {
															v222 = v77 + v214<<(uint(int32(2))%32)
															v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
															v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
															if v223 <= v224 {
																v227 = v209
																v235 = v224
																v237 = base.I64_extend_i32_s(v223)
																for {
																	if v214 != 0 {
																		v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227-int32(4)))))
																		if v237 == v240 {
																			v246 = v227
																			v247 = v235
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																			v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																			v246 = v227 + int32(4)
																			v247 = v245
																		}
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																		v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																		v246 = v227 + int32(4)
																		v247 = v245
																	}
																	if v237 < base.I64_extend_i32_s(v247) {
																		v227 = v246
																		v235 = v247
																		v237 = v237 + int64(1)
																		continue
																	} else {
																		break
																	}
																	break
																}
																v252 = v246
															} else {
																v252 = v209
															}
															v264 = v214 + int32(2)
															if v264 < v61 {
																v209 = v252
																v214 = v264
																continue
															} else {
																break
															}
															break
														}
													} else {
													}
													v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
													if v277 != v44 {
														F_pfree(m, v44)
														mBase = m.M
														v280 = m.ExcPending
														if v280 != 0 {
															return int64(0)
														} else {
															v283 = v193
															v293 = F_palloc(m, int32(24))
															mBase = m.M
															v294 = m.ExcPending
															if v294 != 0 {
																return int64(0)
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
																v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
																v302 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
																*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
																return base.I64_extend_i32_u(v293)
															}
														}
													} else {
														v283 = v193
														v293 = F_palloc(m, int32(24))
														mBase = m.M
														v294 = m.ExcPending
														if v294 != 0 {
															return int64(0)
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
															v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
															v302 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
															*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
															return base.I64_extend_i32_u(v293)
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
						return int64(0)
					} else {
						if v52 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							if v44 != v56 {
								v283 = v44
								v293 = F_palloc(m, int32(24))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
									v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
									v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
									v302 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
									*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
									return base.I64_extend_i32_u(v293)
								}
							} else {
								return base.I64_extend_i32_u(v12)
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v61 = F_ArrayGetNItemsSafe(m, v60, v51)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								if v61 < v42 {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									if v44 != v64 {
										v283 = v44
										v293 = F_palloc(m, int32(24))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
											v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
											v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
											v302 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
											*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
											return base.I64_extend_i32_u(v293)
										}
									} else {
										return base.I64_extend_i32_u(v12)
									}
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
									if v68 != 0 {
										v76 = v68
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
										v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									v77 = v76 + v44
									v79 = int32(0)
									if v61 <= v79 {
										v178 = int64(0)
									} else {
										v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
										v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77))))
										v90 = v86 - v87 + int64(1)
										if base.Ui32(v61) < base.Ui32(int32(3)) {
											v178 = v90
										} else {
											v96 = int32(base.Ui32(v61-int32(3)) >> (uint(int32(1)) % 32))
											if v96 == int32(0) {
												v155 = int32(2)
												v156 = v90
												v164 = v77 + v155<<(uint(int32(2))%32)
												v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
												if v165 == v168 {
													v178 = v156
												} else {
													v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
													v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
												}
											} else {
												v100 = int32(1)
												v101 = v96 + v100
												v108 = int32(2)
												v109 = v90
												v114 = v79
												for {
													v117 = v77 + v108<<(uint(int32(2))%32)
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
													if v118 != v121 {
														v123 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
														v129 = v109 + v123 - base.I64_extend_i32_s(v118) + int64(1)
													} else {
														v129 = v109
													}
													v130 = int32(2)
													v134 = v77 + (v108+v130)<<(uint(v130)%32)
													v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v134-int32(4))))
													if v135 != v138 {
														v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v134)+4)))
														v146 = v129 + v140 - base.I64_extend_i32_s(v135) + int64(1)
													} else {
														v146 = v129
													}
													v148 = v108 + int32(4)
													v150 = v114 + int32(2)
													if v150 != v101&int32(-2) {
														v108 = v148
														v109 = v146
														v114 = v150
														continue
													} else {
														break
													}
													break
												}
												if v101&v100 == int32(0) {
													v178 = v146
												} else {
													v155 = v148
													v156 = v146
													v164 = v77 + v155<<(uint(int32(2))%32)
													v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
													v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
													if v165 == v168 {
														v178 = v156
													} else {
														v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
														v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
													}
												}
											}
										}
									}
									if base.Ui64(v178-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
										v190 = int32(-1)
									} else {
										v190 = base.I32_wrap_i64(v178)
									}
									if base.Ui32(int32(67108864)) <= base.Ui32(v190) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v326 = m.ExcPending
										if v326 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(261))
											mBase = m.M
											v329 = m.ExcPending
											if v329 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
												mBase = m.M
												v333 = m.ExcPending
												if v333 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(339), int32(_a_F_g_int_decompress_2))
													mBase = m.M
													v338 = m.ExcPending
													if v338 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v193 = F_new_intArrayType(m, v190)
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int64(0)
										} else {
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
											if v195 == int32(0) {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
												v205 = (v198<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v205 = v195
											}
											if int32(0) < v61 {
												v209 = v205 + v193
												v214 = v2
												for {
													v222 = v77 + v214<<(uint(int32(2))%32)
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
													v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
													if v223 <= v224 {
														v227 = v209
														v235 = v224
														v237 = base.I64_extend_i32_s(v223)
														for {
															if v214 != 0 {
																v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227-int32(4)))))
																if v237 == v240 {
																	v246 = v227
																	v247 = v235
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																	v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																	v246 = v227 + int32(4)
																	v247 = v245
																}
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																v246 = v227 + int32(4)
																v247 = v245
															}
															if v237 < base.I64_extend_i32_s(v247) {
																v227 = v246
																v235 = v247
																v237 = v237 + int64(1)
																continue
															} else {
																break
															}
															break
														}
														v252 = v246
													} else {
														v252 = v209
													}
													v264 = v214 + int32(2)
													if v264 < v61 {
														v209 = v252
														v214 = v264
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											if v277 != v44 {
												F_pfree(m, v44)
												mBase = m.M
												v280 = m.ExcPending
												if v280 != 0 {
													return int64(0)
												} else {
													v283 = v193
													v293 = F_palloc(m, int32(24))
													mBase = m.M
													v294 = m.ExcPending
													if v294 != 0 {
														return int64(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
														v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
														v302 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
														*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
														return base.I64_extend_i32_u(v293)
													}
												}
											} else {
												v283 = v193
												v293 = F_palloc(m, int32(24))
												mBase = m.M
												v294 = m.ExcPending
												if v294 != 0 {
													return int64(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
													v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
													v302 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
													*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
													return base.I64_extend_i32_u(v293)
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
			return int64(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
			if v46 != 0 {
				v47 = F_array_contains_nulls(m, v44)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					if v47 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v310 = m.ExcPending
						if v310 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v313 = m.ExcPending
							if v313 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_g_int_decompress_0), int32(0))
								mBase = m.M
								v317 = m.ExcPending
								if v317 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(305), int32(_a_F_g_int_decompress_2))
									mBase = m.M
									v322 = m.ExcPending
									if v322 != 0 {
										return int64(0)
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
							return int64(0)
						} else {
							if v52 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								if v44 != v56 {
									v283 = v44
									v293 = F_palloc(m, int32(24))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
										v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
										v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
										v302 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
										*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
										return base.I64_extend_i32_u(v293)
									}
								} else {
									return base.I64_extend_i32_u(v12)
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
								v61 = F_ArrayGetNItemsSafe(m, v60, v51)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									if v61 < v42 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										if v44 != v64 {
											v283 = v44
											v293 = F_palloc(m, int32(24))
											mBase = m.M
											v294 = m.ExcPending
											if v294 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
												v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
												v302 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
												*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
												return base.I64_extend_i32_u(v293)
											}
										} else {
											return base.I64_extend_i32_u(v12)
										}
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
										if v68 != 0 {
											v76 = v68
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
											v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v77 = v76 + v44
										v79 = int32(0)
										if v61 <= v79 {
											v178 = int64(0)
										} else {
											v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
											v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77))))
											v90 = v86 - v87 + int64(1)
											if base.Ui32(v61) < base.Ui32(int32(3)) {
												v178 = v90
											} else {
												v96 = int32(base.Ui32(v61-int32(3)) >> (uint(int32(1)) % 32))
												if v96 == int32(0) {
													v155 = int32(2)
													v156 = v90
													v164 = v77 + v155<<(uint(int32(2))%32)
													v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
													v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
													if v165 == v168 {
														v178 = v156
													} else {
														v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
														v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
													}
												} else {
													v100 = int32(1)
													v101 = v96 + v100
													v108 = int32(2)
													v109 = v90
													v114 = v79
													for {
														v117 = v77 + v108<<(uint(int32(2))%32)
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
														if v118 != v121 {
															v123 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
															v129 = v109 + v123 - base.I64_extend_i32_s(v118) + int64(1)
														} else {
															v129 = v109
														}
														v130 = int32(2)
														v134 = v77 + (v108+v130)<<(uint(v130)%32)
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v134-int32(4))))
														if v135 != v138 {
															v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v134)+4)))
															v146 = v129 + v140 - base.I64_extend_i32_s(v135) + int64(1)
														} else {
															v146 = v129
														}
														v148 = v108 + int32(4)
														v150 = v114 + int32(2)
														if v150 != v101&int32(-2) {
															v108 = v148
															v109 = v146
															v114 = v150
															continue
														} else {
															break
														}
														break
													}
													if v101&v100 == int32(0) {
														v178 = v146
													} else {
														v155 = v148
														v156 = v146
														v164 = v77 + v155<<(uint(int32(2))%32)
														v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
														if v165 == v168 {
															v178 = v156
														} else {
															v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
															v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
														}
													}
												}
											}
										}
										if base.Ui64(v178-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
											v190 = int32(-1)
										} else {
											v190 = base.I32_wrap_i64(v178)
										}
										if base.Ui32(int32(67108864)) <= base.Ui32(v190) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v326 = m.ExcPending
											if v326 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(261))
												mBase = m.M
												v329 = m.ExcPending
												if v329 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
													mBase = m.M
													v333 = m.ExcPending
													if v333 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(339), int32(_a_F_g_int_decompress_2))
														mBase = m.M
														v338 = m.ExcPending
														if v338 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v193 = F_new_intArrayType(m, v190)
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int64(0)
											} else {
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
												if v195 == int32(0) {
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
													v205 = (v198<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v205 = v195
												}
												if int32(0) < v61 {
													v209 = v205 + v193
													v214 = v2
													for {
														v222 = v77 + v214<<(uint(int32(2))%32)
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
														v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
														if v223 <= v224 {
															v227 = v209
															v235 = v224
															v237 = base.I64_extend_i32_s(v223)
															for {
																if v214 != 0 {
																	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227-int32(4)))))
																	if v237 == v240 {
																		v246 = v227
																		v247 = v235
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																		v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																		v246 = v227 + int32(4)
																		v247 = v245
																	}
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																	v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																	v246 = v227 + int32(4)
																	v247 = v245
																}
																if v237 < base.I64_extend_i32_s(v247) {
																	v227 = v246
																	v235 = v247
																	v237 = v237 + int64(1)
																	continue
																} else {
																	break
																}
																break
															}
															v252 = v246
														} else {
															v252 = v209
														}
														v264 = v214 + int32(2)
														if v264 < v61 {
															v209 = v252
															v214 = v264
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
												if v277 != v44 {
													F_pfree(m, v44)
													mBase = m.M
													v280 = m.ExcPending
													if v280 != 0 {
														return int64(0)
													} else {
														v283 = v193
														v293 = F_palloc(m, int32(24))
														mBase = m.M
														v294 = m.ExcPending
														if v294 != 0 {
															return int64(0)
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
															v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
															v302 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
															*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
															return base.I64_extend_i32_u(v293)
														}
													}
												} else {
													v283 = v193
													v293 = F_palloc(m, int32(24))
													mBase = m.M
													v294 = m.ExcPending
													if v294 != 0 {
														return int64(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
														v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
														v302 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
														*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
														return base.I64_extend_i32_u(v293)
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
					return int64(0)
				} else {
					if v52 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v44 != v56 {
							v283 = v44
							v293 = F_palloc(m, int32(24))
							mBase = m.M
							v294 = m.ExcPending
							if v294 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
								v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
								v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
								v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
								v302 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
								*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
								return base.I64_extend_i32_u(v293)
							}
						} else {
							return base.I64_extend_i32_u(v12)
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v61 = F_ArrayGetNItemsSafe(m, v60, v51)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							if v61 < v42 {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								if v44 != v64 {
									v283 = v44
									v293 = F_palloc(m, int32(24))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
										v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
										v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
										v302 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
										*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
										return base.I64_extend_i32_u(v293)
									}
								} else {
									return base.I64_extend_i32_u(v12)
								}
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
								if v68 != 0 {
									v76 = v68
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
									v76 = (v69<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v77 = v76 + v44
								v79 = int32(0)
								if v61 <= v79 {
									v178 = int64(0)
								} else {
									v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
									v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77))))
									v90 = v86 - v87 + int64(1)
									if base.Ui32(v61) < base.Ui32(int32(3)) {
										v178 = v90
									} else {
										v96 = int32(base.Ui32(v61-int32(3)) >> (uint(int32(1)) % 32))
										if v96 == int32(0) {
											v155 = int32(2)
											v156 = v90
											v164 = v77 + v155<<(uint(int32(2))%32)
											v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
											if v165 == v168 {
												v178 = v156
											} else {
												v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
												v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
											}
										} else {
											v100 = int32(1)
											v101 = v96 + v100
											v108 = int32(2)
											v109 = v90
											v114 = v79
											for {
												v117 = v77 + v108<<(uint(int32(2))%32)
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
												if v118 != v121 {
													v123 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
													v129 = v109 + v123 - base.I64_extend_i32_s(v118) + int64(1)
												} else {
													v129 = v109
												}
												v130 = int32(2)
												v134 = v77 + (v108+v130)<<(uint(v130)%32)
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v134-int32(4))))
												if v135 != v138 {
													v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v134)+4)))
													v146 = v129 + v140 - base.I64_extend_i32_s(v135) + int64(1)
												} else {
													v146 = v129
												}
												v148 = v108 + int32(4)
												v150 = v114 + int32(2)
												if v150 != v101&int32(-2) {
													v108 = v148
													v109 = v146
													v114 = v150
													continue
												} else {
													break
												}
												break
											}
											if v101&v100 == int32(0) {
												v178 = v146
											} else {
												v155 = v148
												v156 = v146
												v164 = v77 + v155<<(uint(int32(2))%32)
												v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
												if v165 == v168 {
													v178 = v156
												} else {
													v170 = int64(*(*int32)(unsafe.Add(mBase, uint32(v164)+4)))
													v178 = v156 + v170 - base.I64_extend_i32_s(v165) + int64(1)
												}
											}
										}
									}
								}
								if base.Ui64(v178-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
									v190 = int32(-1)
								} else {
									v190 = base.I32_wrap_i64(v178)
								}
								if base.Ui32(int32(67108864)) <= base.Ui32(v190) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v326 = m.ExcPending
									if v326 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v329 = m.ExcPending
										if v329 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_g_int_decompress_3), int32(0))
											mBase = m.M
											v333 = m.ExcPending
											if v333 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_g_int_decompress_1), int32(339), int32(_a_F_g_int_decompress_2))
												mBase = m.M
												v338 = m.ExcPending
												if v338 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v193 = F_new_intArrayType(m, v190)
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int64(0)
									} else {
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
										if v195 == int32(0) {
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
											v205 = (v198<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v205 = v195
										}
										if int32(0) < v61 {
											v209 = v205 + v193
											v214 = v2
											for {
												v222 = v77 + v214<<(uint(int32(2))%32)
												v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
												v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
												if v223 <= v224 {
													v227 = v209
													v235 = v224
													v237 = base.I64_extend_i32_s(v223)
													for {
														if v214 != 0 {
															v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v227-int32(4)))))
															if v237 == v240 {
																v246 = v227
																v247 = v235
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
																v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
																v246 = v227 + int32(4)
																v247 = v245
															}
														} else {
															*(*uint32)(unsafe.Add(mBase, uint32(v227))) = uint32(v237)
															v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
															v246 = v227 + int32(4)
															v247 = v245
														}
														if v237 < base.I64_extend_i32_s(v247) {
															v227 = v246
															v235 = v247
															v237 = v237 + int64(1)
															continue
														} else {
															break
														}
														break
													}
													v252 = v246
												} else {
													v252 = v209
												}
												v264 = v214 + int32(2)
												if v264 < v61 {
													v209 = v252
													v214 = v264
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										if v277 != v44 {
											F_pfree(m, v44)
											mBase = m.M
											v280 = m.ExcPending
											if v280 != 0 {
												return int64(0)
											} else {
												v283 = v193
												v293 = F_palloc(m, int32(24))
												mBase = m.M
												v294 = m.ExcPending
												if v294 != 0 {
													return int64(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
													v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
													v302 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
													*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
													return base.I64_extend_i32_u(v293)
												}
											}
										} else {
											v283 = v193
											v293 = F_palloc(m, int32(24))
											mBase = m.M
											v294 = m.ExcPending
											if v294 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v293))) = base.I64_extend_i32_u(v283)
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v297
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v299
												v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)))
												v302 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v293)+18)) = uint8(v302)
												*(*uint16)(unsafe.Add(mBase, uint32(v293)+16)) = uint16(v301)
												return base.I64_extend_i32_u(v293)
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
func F_g_int_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v165 = m.ExcPending
	if v165 != 0 {
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8)+v24*int32(24))))
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
	return int64(0)
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
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8)+v83*int32(24))))
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
	v151 = F__int_unique(m, v61)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v15)))) = int32(base.Ui32(v153) >> (uint(int32(2)) % 32))
	m.G0 = v13 + int32(16)
	return base.I64_extend_i32_u(v151)
L37:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_g_int_union_0), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_g_int_union_1), int32(137), int32(_a_F_g_int_union_2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
