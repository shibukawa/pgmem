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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v2 {
		v32 = v2
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		if v19 == int32(0) {
			v32 = v2
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v22 != int32(7) {
				v32 = v2
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				if v25 != int32(17) {
					v32 = v2
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
					v32 = v28 ^ int32(1)
				}
			}
		}
	}
	if v32&int32(1) != 0 {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v36 = F_get_fn_opclass_options(m, v35)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v43 = v40 << (uint(int32(1)) % 32)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v45 = F_pg_detoast_datum(m, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				if v47 != 0 {
					v48 = F_array_contains_nulls(m, v45)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v255 = m.ExcPending
								if v255 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(161982), int32(0))
									mBase = m.M
									v261 = m.ExcPending
									if v261 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(514759), int32(304), int32(134450))
										mBase = m.M
										v268 = m.ExcPending
										if v268 != 0 {
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
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							v52 = v45 + int32(16)
							v53 = F_ArrayGetNItems(m, v50, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v53 == int32(0) {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									if v45 != v57 {
										v225 = v45
										v237 = F_palloc(m, int32(16))
										mBase = m.M
										v238 = m.ExcPending
										if v238 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
											v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
											v245 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
											return v237
										}
									} else {
										return v13
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									v61 = F_ArrayGetNItems(m, v60, v52)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 < v43 {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											if v45 != v64 {
												v225 = v45
												v237 = F_palloc(m, int32(16))
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
													v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
													v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
													v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
													v245 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
													return v237
												}
											} else {
												return v13
											}
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
											if v67 != 0 {
												v75 = v67
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											v76 = v75 + v45
											if v61 <= int32(0) {
												v117 = int64(0)
											} else {
												v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+4)))
												v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76))))
												v87 = v83 - v84 + int64(1)
												if base.Ui32(v61) < base.Ui32(int32(3)) {
													v117 = v87
												} else {
													v93 = v87
													v94 = int32(2)
													for {
														v99 = v76 + v94<<(uint(int32(2))%32)
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
														if v100 != v103 {
															v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v99)+4)))
															v111 = v93 + v105 - base.I64_extend_i32_s(v100) + int64(1)
														} else {
															v111 = v93
														}
														v113 = v94 + int32(2)
														if v113 < v61 {
															v93 = v111
															v94 = v113
															continue
														} else {
															break
														}
														break
													}
													v117 = v111
												}
											}
											if base.Ui64(v117-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
												v127 = int32(-1)
											} else {
												v127 = base.I32_wrap_i64(v117)
											}
											if base.Ui32(int32(134217725)) <= base.Ui32(v127) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v272 = m.ExcPending
												if v272 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(261))
													mBase = m.M
													v275 = m.ExcPending
													if v275 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(484333), int32(0))
														mBase = m.M
														v281 = m.ExcPending
														if v281 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(514759), int32(338), int32(134450))
															mBase = m.M
															v288 = m.ExcPending
															if v288 != 0 {
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
												v130 = F_new_intArrayType(m, v127)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
													if v132 == int32(0) {
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
														v142 = (v135<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													} else {
														v142 = v132
													}
													if int32(0) < v61 {
														v146 = v142 + v130
														v150 = v2
														for {
															v160 = v76 + v150<<(uint(int32(2))%32)
															v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
															v163 = v160 + int32(4)
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
															if v161 <= v164 {
																v167 = v146
																v174 = v164
																v178 = base.I64_extend_i32_s(v161)
																for {
																	if v150 != 0 {
																		v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167-int32(4)))))
																		if v178 == v181 {
																			v187 = v167
																			v188 = v174
																		} else {
																			*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																			v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																			v187 = v167 + int32(4)
																			v188 = v186
																		}
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																		v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																		v187 = v167 + int32(4)
																		v188 = v186
																	}
																	if v178 < base.I64_extend_i32_s(v188) {
																		v167 = v187
																		v174 = v188
																		v178 = v178 + int64(1)
																		continue
																	} else {
																		break
																	}
																	break
																}
																v193 = v187
															} else {
																v193 = v146
															}
															v206 = v150 + int32(2)
															if v206 < v61 {
																v146 = v193
																v150 = v206
																continue
															} else {
																break
															}
															break
														}
													} else {
													}
													v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													if v220 != v45 {
														F_pfree(m, v45)
														mBase = m.M
														v223 = m.ExcPending
														if v223 != 0 {
															return int32(0)
														} else {
															v225 = v130
															v237 = F_palloc(m, int32(16))
															mBase = m.M
															v238 = m.ExcPending
															if v238 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
																v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
																v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
																v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
																v245 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
																*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
																return v237
															}
														}
													} else {
														v225 = v130
														v237 = F_palloc(m, int32(16))
														mBase = m.M
														v238 = m.ExcPending
														if v238 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
															v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
															v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
															v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
															v245 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
															*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
															return v237
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
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					v52 = v45 + int32(16)
					v53 = F_ArrayGetNItems(m, v50, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						if v53 == int32(0) {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							if v45 != v57 {
								v225 = v45
								v237 = F_palloc(m, int32(16))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
									v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
									v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
									v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
									v245 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
									*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
									return v237
								}
							} else {
								return v13
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							v61 = F_ArrayGetNItems(m, v60, v52)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								if v61 < v43 {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									if v45 != v64 {
										v225 = v45
										v237 = F_palloc(m, int32(16))
										mBase = m.M
										v238 = m.ExcPending
										if v238 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
											v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
											v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
											v245 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
											return v237
										}
									} else {
										return v13
									}
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
									if v67 != 0 {
										v75 = v67
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									v76 = v75 + v45
									if v61 <= int32(0) {
										v117 = int64(0)
									} else {
										v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+4)))
										v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76))))
										v87 = v83 - v84 + int64(1)
										if base.Ui32(v61) < base.Ui32(int32(3)) {
											v117 = v87
										} else {
											v93 = v87
											v94 = int32(2)
											for {
												v99 = v76 + v94<<(uint(int32(2))%32)
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
												if v100 != v103 {
													v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v99)+4)))
													v111 = v93 + v105 - base.I64_extend_i32_s(v100) + int64(1)
												} else {
													v111 = v93
												}
												v113 = v94 + int32(2)
												if v113 < v61 {
													v93 = v111
													v94 = v113
													continue
												} else {
													break
												}
												break
											}
											v117 = v111
										}
									}
									if base.Ui64(v117-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
										v127 = int32(-1)
									} else {
										v127 = base.I32_wrap_i64(v117)
									}
									if base.Ui32(int32(134217725)) <= base.Ui32(v127) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v272 = m.ExcPending
										if v272 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(261))
											mBase = m.M
											v275 = m.ExcPending
											if v275 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(484333), int32(0))
												mBase = m.M
												v281 = m.ExcPending
												if v281 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(514759), int32(338), int32(134450))
													mBase = m.M
													v288 = m.ExcPending
													if v288 != 0 {
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
										v130 = F_new_intArrayType(m, v127)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
											if v132 == int32(0) {
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
												v142 = (v135<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v142 = v132
											}
											if int32(0) < v61 {
												v146 = v142 + v130
												v150 = v2
												for {
													v160 = v76 + v150<<(uint(int32(2))%32)
													v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
													v163 = v160 + int32(4)
													v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
													if v161 <= v164 {
														v167 = v146
														v174 = v164
														v178 = base.I64_extend_i32_s(v161)
														for {
															if v150 != 0 {
																v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167-int32(4)))))
																if v178 == v181 {
																	v187 = v167
																	v188 = v174
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																	v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																	v187 = v167 + int32(4)
																	v188 = v186
																}
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																v187 = v167 + int32(4)
																v188 = v186
															}
															if v178 < base.I64_extend_i32_s(v188) {
																v167 = v187
																v174 = v188
																v178 = v178 + int64(1)
																continue
															} else {
																break
															}
															break
														}
														v193 = v187
													} else {
														v193 = v146
													}
													v206 = v150 + int32(2)
													if v206 < v61 {
														v146 = v193
														v150 = v206
														continue
													} else {
														break
													}
													break
												}
											} else {
											}
											v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											if v220 != v45 {
												F_pfree(m, v45)
												mBase = m.M
												v223 = m.ExcPending
												if v223 != 0 {
													return int32(0)
												} else {
													v225 = v130
													v237 = F_palloc(m, int32(16))
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
														v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
														v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
														v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
														v245 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
														*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
														return v237
													}
												}
											} else {
												v225 = v130
												v237 = F_palloc(m, int32(16))
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
													v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
													v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
													v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
													v245 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
													return v237
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
		v43 = int32(200)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v45 = F_pg_detoast_datum(m, v44)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
			if v47 != 0 {
				v48 = F_array_contains_nulls(m, v45)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					if v48 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(161982), int32(0))
								mBase = m.M
								v261 = m.ExcPending
								if v261 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(514759), int32(304), int32(134450))
									mBase = m.M
									v268 = m.ExcPending
									if v268 != 0 {
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
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						v52 = v45 + int32(16)
						v53 = F_ArrayGetNItems(m, v50, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							if v53 == int32(0) {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								if v45 != v57 {
									v225 = v45
									v237 = F_palloc(m, int32(16))
									mBase = m.M
									v238 = m.ExcPending
									if v238 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
										v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
										v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
										v245 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
										return v237
									}
								} else {
									return v13
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								v61 = F_ArrayGetNItems(m, v60, v52)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 < v43 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										if v45 != v64 {
											v225 = v45
											v237 = F_palloc(m, int32(16))
											mBase = m.M
											v238 = m.ExcPending
											if v238 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
												v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
												v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
												v245 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
												return v237
											}
										} else {
											return v13
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
										if v67 != 0 {
											v75 = v67
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v76 = v75 + v45
										if v61 <= int32(0) {
											v117 = int64(0)
										} else {
											v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+4)))
											v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76))))
											v87 = v83 - v84 + int64(1)
											if base.Ui32(v61) < base.Ui32(int32(3)) {
												v117 = v87
											} else {
												v93 = v87
												v94 = int32(2)
												for {
													v99 = v76 + v94<<(uint(int32(2))%32)
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
													if v100 != v103 {
														v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v99)+4)))
														v111 = v93 + v105 - base.I64_extend_i32_s(v100) + int64(1)
													} else {
														v111 = v93
													}
													v113 = v94 + int32(2)
													if v113 < v61 {
														v93 = v111
														v94 = v113
														continue
													} else {
														break
													}
													break
												}
												v117 = v111
											}
										}
										if base.Ui64(v117-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
											v127 = int32(-1)
										} else {
											v127 = base.I32_wrap_i64(v117)
										}
										if base.Ui32(int32(134217725)) <= base.Ui32(v127) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(261))
												mBase = m.M
												v275 = m.ExcPending
												if v275 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(484333), int32(0))
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(514759), int32(338), int32(134450))
														mBase = m.M
														v288 = m.ExcPending
														if v288 != 0 {
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
											v130 = F_new_intArrayType(m, v127)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
												if v132 == int32(0) {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
													v142 = (v135<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v142 = v132
												}
												if int32(0) < v61 {
													v146 = v142 + v130
													v150 = v2
													for {
														v160 = v76 + v150<<(uint(int32(2))%32)
														v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
														v163 = v160 + int32(4)
														v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
														if v161 <= v164 {
															v167 = v146
															v174 = v164
															v178 = base.I64_extend_i32_s(v161)
															for {
																if v150 != 0 {
																	v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167-int32(4)))))
																	if v178 == v181 {
																		v187 = v167
																		v188 = v174
																	} else {
																		*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																		v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																		v187 = v167 + int32(4)
																		v188 = v186
																	}
																} else {
																	*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																	v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																	v187 = v167 + int32(4)
																	v188 = v186
																}
																if v178 < base.I64_extend_i32_s(v188) {
																	v167 = v187
																	v174 = v188
																	v178 = v178 + int64(1)
																	continue
																} else {
																	break
																}
																break
															}
															v193 = v187
														} else {
															v193 = v146
														}
														v206 = v150 + int32(2)
														if v206 < v61 {
															v146 = v193
															v150 = v206
															continue
														} else {
															break
														}
														break
													}
												} else {
												}
												v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												if v220 != v45 {
													F_pfree(m, v45)
													mBase = m.M
													v223 = m.ExcPending
													if v223 != 0 {
														return int32(0)
													} else {
														v225 = v130
														v237 = F_palloc(m, int32(16))
														mBase = m.M
														v238 = m.ExcPending
														if v238 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
															v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
															v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
															v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
															v245 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
															*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
															return v237
														}
													}
												} else {
													v225 = v130
													v237 = F_palloc(m, int32(16))
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
														v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
														v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
														v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
														v245 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
														*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
														return v237
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
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
				v52 = v45 + int32(16)
				v53 = F_ArrayGetNItems(m, v50, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					if v53 == int32(0) {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						if v45 != v57 {
							v225 = v45
							v237 = F_palloc(m, int32(16))
							mBase = m.M
							v238 = m.ExcPending
							if v238 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
								v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
								v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
								v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
								v245 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
								*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
								return v237
							}
						} else {
							return v13
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						v61 = F_ArrayGetNItems(m, v60, v52)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 < v43 {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								if v45 != v64 {
									v225 = v45
									v237 = F_palloc(m, int32(16))
									mBase = m.M
									v238 = m.ExcPending
									if v238 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
										v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
										v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
										v245 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
										return v237
									}
								} else {
									return v13
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
								if v67 != 0 {
									v75 = v67
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v76 = v75 + v45
								if v61 <= int32(0) {
									v117 = int64(0)
								} else {
									v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+4)))
									v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76))))
									v87 = v83 - v84 + int64(1)
									if base.Ui32(v61) < base.Ui32(int32(3)) {
										v117 = v87
									} else {
										v93 = v87
										v94 = int32(2)
										for {
											v99 = v76 + v94<<(uint(int32(2))%32)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
											if v100 != v103 {
												v105 = int64(*(*int32)(unsafe.Add(mBase, uint32(v99)+4)))
												v111 = v93 + v105 - base.I64_extend_i32_s(v100) + int64(1)
											} else {
												v111 = v93
											}
											v113 = v94 + int32(2)
											if v113 < v61 {
												v93 = v111
												v94 = v113
												continue
											} else {
												break
											}
											break
										}
										v117 = v111
									}
								}
								if base.Ui64(v117-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
									v127 = int32(-1)
								} else {
									v127 = base.I32_wrap_i64(v117)
								}
								if base.Ui32(int32(134217725)) <= base.Ui32(v127) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v275 = m.ExcPending
										if v275 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(484333), int32(0))
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(514759), int32(338), int32(134450))
												mBase = m.M
												v288 = m.ExcPending
												if v288 != 0 {
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
									v130 = F_new_intArrayType(m, v127)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
										if v132 == int32(0) {
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
											v142 = (v135<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v142 = v132
										}
										if int32(0) < v61 {
											v146 = v142 + v130
											v150 = v2
											for {
												v160 = v76 + v150<<(uint(int32(2))%32)
												v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
												v163 = v160 + int32(4)
												v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
												if v161 <= v164 {
													v167 = v146
													v174 = v164
													v178 = base.I64_extend_i32_s(v161)
													for {
														if v150 != 0 {
															v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v167-int32(4)))))
															if v178 == v181 {
																v187 = v167
																v188 = v174
															} else {
																*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
																v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																v187 = v167 + int32(4)
																v188 = v186
															}
														} else {
															*(*uint32)(unsafe.Add(mBase, uint32(v167))) = uint32(v178)
															v186 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
															v187 = v167 + int32(4)
															v188 = v186
														}
														if v178 < base.I64_extend_i32_s(v188) {
															v167 = v187
															v174 = v188
															v178 = v178 + int64(1)
															continue
														} else {
															break
														}
														break
													}
													v193 = v187
												} else {
													v193 = v146
												}
												v206 = v150 + int32(2)
												if v206 < v61 {
													v146 = v193
													v150 = v206
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										if v220 != v45 {
											F_pfree(m, v45)
											mBase = m.M
											v223 = m.ExcPending
											if v223 != 0 {
												return int32(0)
											} else {
												v225 = v130
												v237 = F_palloc(m, int32(16))
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
													v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
													v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
													v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
													v245 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
													return v237
												}
											}
										} else {
											v225 = v130
											v237 = F_palloc(m, int32(16))
											mBase = m.M
											v238 = m.ExcPending
											if v238 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v237))) = v225
												v240 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v240
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = v242
												v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
												v245 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v237)+14)) = uint8(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v244)
												return v237
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
	var v23 int32
	_ = v23
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
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
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 <= v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L11
	} else {
		goto L39
	}
L2:
	;
	v61 = F_new_intArrayType(m, v54)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L16
	}
L3:
	;
	v54 = v2
	goto L2
L4:
	;
	goto L5
L5:
	;
	v23 = v2
	v25 = v2
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(4)+v23<<(uint(int32(4))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v46
	goto L2
L8:
	;
	v37 = F_array_contains_nulls(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v44 = F_ArrayGetNItems(m, v41, v35+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	if v37 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v46 = v44 + v25
	v48 = v23 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v48 < v49 {
		v23 = v48
		v25 = v46
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v63 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = (v66<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L19
L18:
	;
	v73 = v63
	goto L19
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if int32(0) < v74 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v82 = int32(0)
	v86 = v73 + v61
	goto L23
L21:
	;
	goto L22
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v132 = F_ArrayGetNItems(m, v129, v61+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L34
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(4)+v82<<(uint(int32(4))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v98 = F_ArrayGetNItems(m, v95, v94+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v100 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v108 = v100
	goto L28
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v108 = (v101<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L28
L28:
	;
	v111 = v98 << (uint(int32(2)) % 32)
	if v111 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v116 = v82 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v116 < v117 {
		v82 = v116
		v86 = v113 + v111
		goto L23
	} else {
		goto L33
	}
L30:
	;
	v112 = F__emscripten_memcpy_bulkmem(m, v86, v108+v94, v111)
	mBase = m.M
	v113 = v112
	goto L32
L31:
	;
	v113 = v86
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L24
L34:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v134)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v136 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v144 = v136
	goto L37
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v144 = (v137<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L37:
	;
	F_isort(m, v144+v61, v132, v13+int32(15))
	mBase = m.M
	v149 = F__int_unique(m, v61)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(base.Ui32(v151) >> (uint(int32(2)) % 32))
	m.G0 = v13 + int32(16)
	return v149
L39:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(161982), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(514759), int32(136), int32(284758))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
