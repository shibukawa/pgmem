package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IdleStatsUpdateTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1125])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[8])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_IncrementVarSublevelsUp_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	v3 = int32(0)
	if l0 == v3 {
		v94 = v3
		return v94
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 <= int32(57) {
			switch v7 - int32(6) {
			case 0:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v24) < base.Ui32(v25) {
					v94 = v3
					return v94
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v27 + v24
					return int32(0)
				}
			default:
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 3:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v48) < base.Ui32(v49) {
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v51 + v48
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 4:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v54) < base.Ui32(v55) {
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v57 + v54
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			}
		} else {
			switch v7 - int32(58) {
			case 0:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v32 != 0 {
					v94 = v3
					return v94
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(207332), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495848), int32(818), int32(221575))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 1, 2, 4, 5, 6, 7, 8:
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 3:
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v60 < v61 {
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v63 + v60
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 9:
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v82 + int32(1)
				v88 = F_query_tree_walker_impl(m, l0, int32(1050), l1, int32(16))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90 - int32(1)
					v94 = v88
					return v94
				}
			default:
				if v7 == int32(101) {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v71 != int32(6) {
						v94 = v3
						return v94
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if base.Ui32(v74) < base.Ui32(v75) {
							v94 = v3
							return v94
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v77 + v74
							return int32(0)
						}
					}
				} else {
					if v7 != int32(319) {
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if base.Ui32(v18) < base.Ui32(v19) {
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21 + v18
						}
					}
					v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						return v68
					}
				}
			}
		}
	}
}
func F_InitProcess(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v169 int32
	_ = v169
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(68632), int32(0))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errfinish(m, int32(499857), int32(402), int32(130129))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
			if v15 == int32(1) {
				F_RegisterPostmasterChildActive(m)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[198]))
					v24 = v22 - int32(3)
					if base.Ui32(v24) <= base.Ui32(int32(4)) {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[861])))
						v32 = v31
					} else {
						v32 = int32(20)
					}
					v34 = *(*int32)(unsafe.Add(mBase, _consts[615]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v37 = *(*int32)(unsafe.Add(mBase, _consts[150]))
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(1)
					v40 = v32 + v37
					if v35 != 0 {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[615]))
						F_s_lock(m, v42, int32(499857), int32(432), int32(130129))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, _consts[150]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
							*(*int32)(unsafe.Add(mBase, _consts[333])) = v50
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							if v53 != v40 {
								v56 = v53
							} else {
								v56 = int32(0)
							}
							if v56 != 0 {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
								*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
								v112 = int32(4438828)
								*(*int32)(unsafe.Add(mBase, _consts[137])) = v53
								v115 = *(*int32)(unsafe.Add(mBase, _consts[615]))
								v116 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116
								v118 = int32(4122272)
								v120 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								v122 = *(*int32)(unsafe.Add(mBase, _consts[150]))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
								v126 = base.I32_div_s(v120-v123, int32(640))
								*(*int32)(unsafe.Add(mBase, _consts[109])) = v126
								v128 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v120))) = v128
								v131 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+612)) = v116
								*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)) = uint8(v116)
								*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v116
								*(*int64)(unsafe.Add(mBase, uint32(v131)+36)) = v128
								v141 = *(*int32)(unsafe.Add(mBase, _consts[138]))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v141
								v144 = *(*int32)(unsafe.Add(mBase, _consts[109]))
								*(*int64)(unsafe.Add(mBase, uint32(v131)+56)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v131)+52)) = v144
								*(*int64)(unsafe.Add(mBase, uint32(v131-int32(-64)))) = v128
								v153 = *(*int32)(unsafe.Add(mBase, _consts[198]))
								*(*int32)(unsafe.Add(mBase, uint32(v131)+120)) = v116
								*(*int64)(unsafe.Add(mBase, uint32(v131)+112)) = v128
								*(*int64)(unsafe.Add(mBase, uint32(v131)+92)) = v128
								*(*uint16)(unsafe.Add(mBase, uint32(v131)+74)) = uint16(v116)
								*(*uint8)(unsafe.Add(mBase, uint32(v131)+124)) = uint8(base.B2i32(v153 == int32(4)))
								*(*uint8)(unsafe.Add(mBase, uint32(v131)+72)) = uint8(base.B2i32(v153 == int32(1)))
								v169 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								*(*int64)(unsafe.Add(mBase, uint32(v169)+560)) = v128
								*(*uint8)(unsafe.Add(mBase, uint32(v169)+536)) = uint8(v116)
								*(*uint8)(unsafe.Add(mBase, uint32(v169)+73)) = uint8(v116)
								*(*int64)(unsafe.Add(mBase, uint32(v169)+128)) = v128
								*(*int64)(unsafe.Add(mBase, uint32(v169)+544)) = v128
								*(*int64)(unsafe.Add(mBase, uint32(v169)+576)) = v128
								*(*int64)(unsafe.Add(mBase, uint32(v169)+568)) = int64(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v169)+136)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v169)+144)) = v116
								*(*uint8)(unsafe.Add(mBase, uint32(v169)+552)) = uint8(v116)
								F_OwnLatch(m, v169+int32(20))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									F_SwitchToSharedLatch(m)
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return
									} else {
										v197 = *(*int32)(unsafe.Add(mBase, _consts[137]))
										*(*int32)(unsafe.Add(mBase, _consts[39])) = v197 + int32(548)
										F_on_shmem_exit(m, int32(1113), int32(0))
										mBase = m.M
										v208 = m.ExcPending
										if v208 != 0 {
											return
										} else {
											v209 = int32(4515392)
											v210 = *(*int32)(unsafe.Add(mBase, _consts[0]))
											v213 = *(*int32)(unsafe.Add(mBase, _consts[12]))
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v213
											v217 = *(*int32)(unsafe.Add(mBase, _consts[111]))
											v220 = F_palloc(m, v217<<(uint(int32(2))%32))
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[862])) = v220
												v225 = *(*int32)(unsafe.Add(mBase, _consts[111]))
												v228 = F_palloc(m, v225*int32(24))
												mBase = m.M
												v229 = m.ExcPending
												if v229 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[863])) = v228
													v233 = *(*int32)(unsafe.Add(mBase, _consts[862]))
													*(*int32)(unsafe.Add(mBase, _consts[864])) = v233
													v237 = *(*int32)(unsafe.Add(mBase, _consts[111]))
													v240 = F_palloc(m, v237<<(uint(int32(2))%32))
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[865])) = v240
														v245 = *(*int32)(unsafe.Add(mBase, _consts[111]))
														v248 = F_palloc(m, v245<<(uint(int32(2))%32))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[866])) = v248
															v253 = *(*int32)(unsafe.Add(mBase, _consts[111]))
															v255 = base.I32_div_s(v253, int32(2))
															v258 = F_palloc(m, v255*int32(12))
															mBase = m.M
															v259 = m.ExcPending
															if v259 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[867])) = v258
																v263 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																v266 = F_palloc(m, v263<<(uint(int32(2))%32))
																mBase = m.M
																v267 = m.ExcPending
																if v267 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[868])) = v266
																	v271 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																	*(*int32)(unsafe.Add(mBase, _consts[869])) = v271
																	v276 = F_palloc(m, v271*int32(20))
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _consts[870])) = v276
																		v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																		*(*int32)(unsafe.Add(mBase, _consts[871])) = v281 << (uint(int32(2)) % 32)
																		v287 = F_palloc(m, v281*int32(80))
																		mBase = m.M
																		v288 = m.ExcPending
																		if v288 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _consts[0])) = v210
																			*(*int32)(unsafe.Add(mBase, _consts[872])) = v287
																			m.G0 = v8 + int32(16)
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
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _consts[615]))
								v59 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59
								v62 = *(*int32)(unsafe.Add(mBase, _consts[198]))
								F_errstart_cold(m, int32(22), v59)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errcode(m, int32(12485))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										if v62 == int32(6) {
											v297 = *(*int32)(unsafe.Add(mBase, _consts[308]))
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v297
											F_errmsg(m, int32(680057), v8)
											mBase = m.M
											v301 = m.ExcPending
											if v301 != 0 {
												return
											} else {
												F_errfinish(m, int32(499857), int32(454), int32(130129))
												mBase = m.M
												v306 = m.ExcPending
												if v306 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											F_errmsg(m, int32(22889), int32(0))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												F_errfinish(m, int32(499857), int32(457), int32(130129))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
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
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _consts[150]))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
						*(*int32)(unsafe.Add(mBase, _consts[333])) = v50
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						if v53 != v40 {
							v56 = v53
						} else {
							v56 = int32(0)
						}
						if v56 != 0 {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v108
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
							v112 = int32(4438828)
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v53
							v115 = *(*int32)(unsafe.Add(mBase, _consts[615]))
							v116 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116
							v118 = int32(4122272)
							v120 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							v122 = *(*int32)(unsafe.Add(mBase, _consts[150]))
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
							v126 = base.I32_div_s(v120-v123, int32(640))
							*(*int32)(unsafe.Add(mBase, _consts[109])) = v126
							v128 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = v128
							v131 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+612)) = v116
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)) = uint8(v116)
							*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v131)+36)) = v128
							v141 = *(*int32)(unsafe.Add(mBase, _consts[138]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v141
							v144 = *(*int32)(unsafe.Add(mBase, _consts[109]))
							*(*int64)(unsafe.Add(mBase, uint32(v131)+56)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v131)+52)) = v144
							*(*int64)(unsafe.Add(mBase, uint32(v131-int32(-64)))) = v128
							v153 = *(*int32)(unsafe.Add(mBase, _consts[198]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+120)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v131)+112)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v131)+92)) = v128
							*(*uint16)(unsafe.Add(mBase, uint32(v131)+74)) = uint16(v116)
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+124)) = uint8(base.B2i32(v153 == int32(4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+72)) = uint8(base.B2i32(v153 == int32(1)))
							v169 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							*(*int64)(unsafe.Add(mBase, uint32(v169)+560)) = v128
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+536)) = uint8(v116)
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+73)) = uint8(v116)
							*(*int64)(unsafe.Add(mBase, uint32(v169)+128)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+544)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+576)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+568)) = int64(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v169)+136)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v169)+144)) = v116
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+552)) = uint8(v116)
							F_OwnLatch(m, v169+int32(20))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, _consts[137]))
									*(*int32)(unsafe.Add(mBase, _consts[39])) = v197 + int32(548)
									F_on_shmem_exit(m, int32(1113), int32(0))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
										return
									} else {
										v209 = int32(4515392)
										v210 = *(*int32)(unsafe.Add(mBase, _consts[0]))
										v213 = *(*int32)(unsafe.Add(mBase, _consts[12]))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v213
										v217 = *(*int32)(unsafe.Add(mBase, _consts[111]))
										v220 = F_palloc(m, v217<<(uint(int32(2))%32))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[862])) = v220
											v225 = *(*int32)(unsafe.Add(mBase, _consts[111]))
											v228 = F_palloc(m, v225*int32(24))
											mBase = m.M
											v229 = m.ExcPending
											if v229 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[863])) = v228
												v233 = *(*int32)(unsafe.Add(mBase, _consts[862]))
												*(*int32)(unsafe.Add(mBase, _consts[864])) = v233
												v237 = *(*int32)(unsafe.Add(mBase, _consts[111]))
												v240 = F_palloc(m, v237<<(uint(int32(2))%32))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[865])) = v240
													v245 = *(*int32)(unsafe.Add(mBase, _consts[111]))
													v248 = F_palloc(m, v245<<(uint(int32(2))%32))
													mBase = m.M
													v249 = m.ExcPending
													if v249 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[866])) = v248
														v253 = *(*int32)(unsafe.Add(mBase, _consts[111]))
														v255 = base.I32_div_s(v253, int32(2))
														v258 = F_palloc(m, v255*int32(12))
														mBase = m.M
														v259 = m.ExcPending
														if v259 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[867])) = v258
															v263 = *(*int32)(unsafe.Add(mBase, _consts[111]))
															v266 = F_palloc(m, v263<<(uint(int32(2))%32))
															mBase = m.M
															v267 = m.ExcPending
															if v267 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[868])) = v266
																v271 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																*(*int32)(unsafe.Add(mBase, _consts[869])) = v271
																v276 = F_palloc(m, v271*int32(20))
																mBase = m.M
																v277 = m.ExcPending
																if v277 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[870])) = v276
																	v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																	*(*int32)(unsafe.Add(mBase, _consts[871])) = v281 << (uint(int32(2)) % 32)
																	v287 = F_palloc(m, v281*int32(80))
																	mBase = m.M
																	v288 = m.ExcPending
																	if v288 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v210
																		*(*int32)(unsafe.Add(mBase, _consts[872])) = v287
																		m.G0 = v8 + int32(16)
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
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _consts[615]))
							v59 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59
							v62 = *(*int32)(unsafe.Add(mBase, _consts[198]))
							F_errstart_cold(m, int32(22), v59)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errcode(m, int32(12485))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									if v62 == int32(6) {
										v297 = *(*int32)(unsafe.Add(mBase, _consts[308]))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v297
										F_errmsg(m, int32(680057), v8)
										mBase = m.M
										v301 = m.ExcPending
										if v301 != 0 {
											return
										} else {
											F_errfinish(m, int32(499857), int32(454), int32(130129))
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errmsg(m, int32(22889), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(499857), int32(457), int32(130129))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
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
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[198]))
				v24 = v22 - int32(3)
				if base.Ui32(v24) <= base.Ui32(int32(4)) {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[861])))
					v32 = v31
				} else {
					v32 = int32(20)
				}
				v34 = *(*int32)(unsafe.Add(mBase, _consts[615]))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
				v37 = *(*int32)(unsafe.Add(mBase, _consts[150]))
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(1)
				v40 = v32 + v37
				if v35 != 0 {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[615]))
					F_s_lock(m, v42, int32(499857), int32(432), int32(130129))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _consts[150]))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
						*(*int32)(unsafe.Add(mBase, _consts[333])) = v50
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						if v53 != v40 {
							v56 = v53
						} else {
							v56 = int32(0)
						}
						if v56 != 0 {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v108
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
							v112 = int32(4438828)
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v53
							v115 = *(*int32)(unsafe.Add(mBase, _consts[615]))
							v116 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116
							v118 = int32(4122272)
							v120 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							v122 = *(*int32)(unsafe.Add(mBase, _consts[150]))
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
							v126 = base.I32_div_s(v120-v123, int32(640))
							*(*int32)(unsafe.Add(mBase, _consts[109])) = v126
							v128 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = v128
							v131 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+612)) = v116
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)) = uint8(v116)
							*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v131)+36)) = v128
							v141 = *(*int32)(unsafe.Add(mBase, _consts[138]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v141
							v144 = *(*int32)(unsafe.Add(mBase, _consts[109]))
							*(*int64)(unsafe.Add(mBase, uint32(v131)+56)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v131)+52)) = v144
							*(*int64)(unsafe.Add(mBase, uint32(v131-int32(-64)))) = v128
							v153 = *(*int32)(unsafe.Add(mBase, _consts[198]))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+120)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v131)+112)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v131)+92)) = v128
							*(*uint16)(unsafe.Add(mBase, uint32(v131)+74)) = uint16(v116)
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+124)) = uint8(base.B2i32(v153 == int32(4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+72)) = uint8(base.B2i32(v153 == int32(1)))
							v169 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							*(*int64)(unsafe.Add(mBase, uint32(v169)+560)) = v128
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+536)) = uint8(v116)
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+73)) = uint8(v116)
							*(*int64)(unsafe.Add(mBase, uint32(v169)+128)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+544)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+576)) = v128
							*(*int64)(unsafe.Add(mBase, uint32(v169)+568)) = int64(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v169)+136)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v169)+144)) = v116
							*(*uint8)(unsafe.Add(mBase, uint32(v169)+552)) = uint8(v116)
							F_OwnLatch(m, v169+int32(20))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, _consts[137]))
									*(*int32)(unsafe.Add(mBase, _consts[39])) = v197 + int32(548)
									F_on_shmem_exit(m, int32(1113), int32(0))
									mBase = m.M
									v208 = m.ExcPending
									if v208 != 0 {
										return
									} else {
										v209 = int32(4515392)
										v210 = *(*int32)(unsafe.Add(mBase, _consts[0]))
										v213 = *(*int32)(unsafe.Add(mBase, _consts[12]))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v213
										v217 = *(*int32)(unsafe.Add(mBase, _consts[111]))
										v220 = F_palloc(m, v217<<(uint(int32(2))%32))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[862])) = v220
											v225 = *(*int32)(unsafe.Add(mBase, _consts[111]))
											v228 = F_palloc(m, v225*int32(24))
											mBase = m.M
											v229 = m.ExcPending
											if v229 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[863])) = v228
												v233 = *(*int32)(unsafe.Add(mBase, _consts[862]))
												*(*int32)(unsafe.Add(mBase, _consts[864])) = v233
												v237 = *(*int32)(unsafe.Add(mBase, _consts[111]))
												v240 = F_palloc(m, v237<<(uint(int32(2))%32))
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[865])) = v240
													v245 = *(*int32)(unsafe.Add(mBase, _consts[111]))
													v248 = F_palloc(m, v245<<(uint(int32(2))%32))
													mBase = m.M
													v249 = m.ExcPending
													if v249 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[866])) = v248
														v253 = *(*int32)(unsafe.Add(mBase, _consts[111]))
														v255 = base.I32_div_s(v253, int32(2))
														v258 = F_palloc(m, v255*int32(12))
														mBase = m.M
														v259 = m.ExcPending
														if v259 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[867])) = v258
															v263 = *(*int32)(unsafe.Add(mBase, _consts[111]))
															v266 = F_palloc(m, v263<<(uint(int32(2))%32))
															mBase = m.M
															v267 = m.ExcPending
															if v267 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[868])) = v266
																v271 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																*(*int32)(unsafe.Add(mBase, _consts[869])) = v271
																v276 = F_palloc(m, v271*int32(20))
																mBase = m.M
																v277 = m.ExcPending
																if v277 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[870])) = v276
																	v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																	*(*int32)(unsafe.Add(mBase, _consts[871])) = v281 << (uint(int32(2)) % 32)
																	v287 = F_palloc(m, v281*int32(80))
																	mBase = m.M
																	v288 = m.ExcPending
																	if v288 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v210
																		*(*int32)(unsafe.Add(mBase, _consts[872])) = v287
																		m.G0 = v8 + int32(16)
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
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _consts[615]))
							v59 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59
							v62 = *(*int32)(unsafe.Add(mBase, _consts[198]))
							F_errstart_cold(m, int32(22), v59)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errcode(m, int32(12485))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									if v62 == int32(6) {
										v297 = *(*int32)(unsafe.Add(mBase, _consts[308]))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v297
										F_errmsg(m, int32(680057), v8)
										mBase = m.M
										v301 = m.ExcPending
										if v301 != 0 {
											return
										} else {
											F_errfinish(m, int32(499857), int32(454), int32(130129))
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errmsg(m, int32(22889), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(499857), int32(457), int32(130129))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
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
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, _consts[150]))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
					*(*int32)(unsafe.Add(mBase, _consts[333])) = v50
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
					if v53 != v40 {
						v56 = v53
					} else {
						v56 = int32(0)
					}
					if v56 != 0 {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v108
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v108))) = v110
						v112 = int32(4438828)
						*(*int32)(unsafe.Add(mBase, _consts[137])) = v53
						v115 = *(*int32)(unsafe.Add(mBase, _consts[615]))
						v116 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v115))) = v116
						v118 = int32(4122272)
						v120 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						v122 = *(*int32)(unsafe.Add(mBase, _consts[150]))
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
						v126 = base.I32_div_s(v120-v123, int32(640))
						*(*int32)(unsafe.Add(mBase, _consts[109])) = v126
						v128 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v120))) = v128
						v131 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+612)) = v116
						*(*uint8)(unsafe.Add(mBase, uint32(v131)+608)) = uint8(v116)
						*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v116
						*(*int64)(unsafe.Add(mBase, uint32(v131)+36)) = v128
						v141 = *(*int32)(unsafe.Add(mBase, _consts[138]))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = v141
						v144 = *(*int32)(unsafe.Add(mBase, _consts[109]))
						*(*int64)(unsafe.Add(mBase, uint32(v131)+56)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v131)+52)) = v144
						*(*int64)(unsafe.Add(mBase, uint32(v131-int32(-64)))) = v128
						v153 = *(*int32)(unsafe.Add(mBase, _consts[198]))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+120)) = v116
						*(*int64)(unsafe.Add(mBase, uint32(v131)+112)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(v131)+92)) = v128
						*(*uint16)(unsafe.Add(mBase, uint32(v131)+74)) = uint16(v116)
						*(*uint8)(unsafe.Add(mBase, uint32(v131)+124)) = uint8(base.B2i32(v153 == int32(4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v131)+72)) = uint8(base.B2i32(v153 == int32(1)))
						v169 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						*(*int64)(unsafe.Add(mBase, uint32(v169)+560)) = v128
						*(*uint8)(unsafe.Add(mBase, uint32(v169)+536)) = uint8(v116)
						*(*uint8)(unsafe.Add(mBase, uint32(v169)+73)) = uint8(v116)
						*(*int64)(unsafe.Add(mBase, uint32(v169)+128)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(v169)+544)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(v169)+576)) = v128
						*(*int64)(unsafe.Add(mBase, uint32(v169)+568)) = int64(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v169)+136)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v169)+144)) = v116
						*(*uint8)(unsafe.Add(mBase, uint32(v169)+552)) = uint8(v116)
						F_OwnLatch(m, v169+int32(20))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							F_SwitchToSharedLatch(m)
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								*(*int32)(unsafe.Add(mBase, _consts[39])) = v197 + int32(548)
								F_on_shmem_exit(m, int32(1113), int32(0))
								mBase = m.M
								v208 = m.ExcPending
								if v208 != 0 {
									return
								} else {
									v209 = int32(4515392)
									v210 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v213 = *(*int32)(unsafe.Add(mBase, _consts[12]))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v213
									v217 = *(*int32)(unsafe.Add(mBase, _consts[111]))
									v220 = F_palloc(m, v217<<(uint(int32(2))%32))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[862])) = v220
										v225 = *(*int32)(unsafe.Add(mBase, _consts[111]))
										v228 = F_palloc(m, v225*int32(24))
										mBase = m.M
										v229 = m.ExcPending
										if v229 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[863])) = v228
											v233 = *(*int32)(unsafe.Add(mBase, _consts[862]))
											*(*int32)(unsafe.Add(mBase, _consts[864])) = v233
											v237 = *(*int32)(unsafe.Add(mBase, _consts[111]))
											v240 = F_palloc(m, v237<<(uint(int32(2))%32))
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[865])) = v240
												v245 = *(*int32)(unsafe.Add(mBase, _consts[111]))
												v248 = F_palloc(m, v245<<(uint(int32(2))%32))
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[866])) = v248
													v253 = *(*int32)(unsafe.Add(mBase, _consts[111]))
													v255 = base.I32_div_s(v253, int32(2))
													v258 = F_palloc(m, v255*int32(12))
													mBase = m.M
													v259 = m.ExcPending
													if v259 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[867])) = v258
														v263 = *(*int32)(unsafe.Add(mBase, _consts[111]))
														v266 = F_palloc(m, v263<<(uint(int32(2))%32))
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[868])) = v266
															v271 = *(*int32)(unsafe.Add(mBase, _consts[111]))
															*(*int32)(unsafe.Add(mBase, _consts[869])) = v271
															v276 = F_palloc(m, v271*int32(20))
															mBase = m.M
															v277 = m.ExcPending
															if v277 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[870])) = v276
																v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
																*(*int32)(unsafe.Add(mBase, _consts[871])) = v281 << (uint(int32(2)) % 32)
																v287 = F_palloc(m, v281*int32(80))
																mBase = m.M
																v288 = m.ExcPending
																if v288 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v210
																	*(*int32)(unsafe.Add(mBase, _consts[872])) = v287
																	m.G0 = v8 + int32(16)
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
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, _consts[615]))
						v59 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v59
						v62 = *(*int32)(unsafe.Add(mBase, _consts[198]))
						F_errstart_cold(m, int32(22), v59)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errcode(m, int32(12485))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								if v62 == int32(6) {
									v297 = *(*int32)(unsafe.Add(mBase, _consts[308]))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v297
									F_errmsg(m, int32(680057), v8)
									mBase = m.M
									v301 = m.ExcPending
									if v301 != 0 {
										return
									} else {
										F_errfinish(m, int32(499857), int32(454), int32(130129))
										mBase = m.M
										v306 = m.ExcPending
										if v306 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errmsg(m, int32(22889), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										F_errfinish(m, int32(499857), int32(457), int32(130129))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
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
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(438210), int32(0))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return
			} else {
				F_errfinish(m, int32(499857), int32(399), int32(130129))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
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
func F_InitializeSessionUserIdStandalone(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	v3 = int32(10)
	*(*int32)(unsafe.Add(mBase, _consts[130])) = v3
	*(*int32)(unsafe.Add(mBase, _consts[129])) = v3
	v11 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[133])) = uint8(v11)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[132])))
	if v14 != 0 {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[132])) = uint8(v16)
		v45 = v3
		v46 = int32(273054)
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
		*(*int32)(unsafe.Add(mBase, _consts[131])) = v45
		F_SetConfigOption(m, int32(217504), v46, int32(0), int32(1))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			return
		}
	} else {
		v20 = int32(10)
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v20
		*(*int32)(unsafe.Add(mBase, _consts[131])) = v20
		F_SetConfigOption(m, int32(217504), int32(273054), int32(0), int32(1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[132])) = uint8(v32)
			v35 = *(*int32)(unsafe.Add(mBase, _consts[130]))
			if v35 == v32 {
				return
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[133])))
				if v41&int32(1) != 0 {
					v44 = int32(273054)
				} else {
					v44 = int32(338926)
				}
				v45 = v35
				v46 = v44
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
				*(*int32)(unsafe.Add(mBase, _consts[131])) = v45
				F_SetConfigOption(m, int32(217504), v46, int32(0), int32(1))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_InvalidationCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v5)
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v8)
	return
}
func F_IoWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v661 int32
	_ = v661
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v728 int64
	_ = v728
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v771 int64
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v804 int32
	_ = v804
	var v810 int64
	_ = v810
	var v815 int64
	_ = v815
	var v817 int64
	_ = v817
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v878 int32
	_ = v878
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1085 int32
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int64
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v24 = v3
	v25 = v3
	v26 = v3
	v27 = v3
	v28 = v3
	v29 = v3
	v30 = int32(-1)
	v33 = v18
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v30 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1151 = int32(m.ExcTag)
	v1152 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1151 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L7:
	;
	v39 = v33 - int32(160)
	m.G0 = v39
	v41 = int32(16)
	v42 = v39 - v41
	m.G0 = v42
	v45 = v42 - v41
	m.G0 = v45
	v48 = v45 - v41
	m.G0 = v48
	v51 = v48 - int32(128)
	m.G0 = v51
	v54 = v51 - v41
	m.G0 = v54
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v56
	*(*int32)(unsafe.Add(mBase, _consts[198])) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v585 = v24
	v586 = v25
	v587 = v26
	v588 = v27
	v589 = v28
	v590 = v29
	v594 = v33
	goto L9
L9:
	;
	if v589 != 0 {
		goto L125
	} else {
		goto L126
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v80 = int32(914)
	v82 = m.G0
	v84 = v82 - int32(144)
	m.G0 = v84
	switch int32(916) {
	case 0, 2:
		v94 = v80
		goto L12
	default:
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v127 = int32(295)
	v129 = m.G0
	v131 = v129 - int32(144)
	m.G0 = v131
	switch int32(297) {
	case 0, 2:
		v141 = v127
		goto L25
	default:
		goto L26
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v94
	F_sigemptyset(m, v84+int32(8))
	mBase = m.M
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[625])) = v80
	v94 = int32(4730)
	goto L12
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+136)) = int32(268435456)
	v106 = v84 + int32(4)
	goto L19
L17:
	;
	m.G0 = v84 + int32(144)
	goto L11
L19:
	;
	goto L20
L20:
	;
	if v106 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v117 = F___memcpy(m, int32(4681004), v106, int32(140))
	mBase = m.M
	goto L23
L22:
	;
	goto L23
L23:
	;
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v174 = int32(-2)
	v176 = m.G0
	v178 = v176 - int32(144)
	m.G0 = v178
	switch int32(0) {
	case 0, 2:
		v188 = v174
		goto L38
	default:
		goto L39
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v141
	F_sigemptyset(m, v131+int32(8))
	mBase = m.M
	goto L28
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[619])) = v127
	v141 = int32(4730)
	goto L25
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+136)) = int32(268435456)
	v153 = v131 + int32(4)
	goto L32
L30:
	;
	m.G0 = v131 + int32(144)
	goto L24
L32:
	;
	goto L33
L33:
	;
	if v153 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v164 = F___memcpy(m, int32(4681144), v153, int32(140))
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L30
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v221 = int32(-2)
	v223 = m.G0
	v225 = v223 - int32(144)
	m.G0 = v225
	switch int32(0) {
	case 0, 2:
		v235 = v221
		goto L51
	default:
		goto L52
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v188
	F_sigemptyset(m, v178+int32(8))
	mBase = m.M
	goto L41
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[622])) = v174
	v188 = int32(4730)
	goto L38
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+136)) = int32(268435456)
	v200 = v178 + int32(4)
	goto L45
L43:
	;
	m.G0 = v178 + int32(144)
	goto L37
L45:
	;
	goto L46
L46:
	;
	if v200 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v211 = F___memcpy(m, int32(4682964), v200, int32(140))
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L43
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v268 = int32(-2)
	v270 = m.G0
	v272 = v270 - int32(144)
	m.G0 = v272
	switch int32(0) {
	case 0, 2:
		v282 = v268
		goto L64
	default:
		goto L65
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v235
	F_sigemptyset(m, v225+int32(8))
	mBase = m.M
	goto L54
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[779])) = v221
	v235 = int32(4730)
	goto L51
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+136)) = int32(268435456)
	v247 = v225 + int32(4)
	goto L58
L56:
	;
	m.G0 = v225 + int32(144)
	goto L50
L58:
	;
	goto L59
L59:
	;
	if v247 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v258 = F___memcpy(m, int32(4682824), v247, int32(140))
	mBase = m.M
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L56
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v315 = int32(917)
	v317 = m.G0
	v319 = v317 - int32(144)
	m.G0 = v319
	switch int32(919) {
	case 0, 2:
		v329 = v315
		goto L77
	default:
		goto L78
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+4)) = v282
	F_sigemptyset(m, v272+int32(8))
	mBase = m.M
	goto L67
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[634])) = v268
	v282 = int32(4730)
	goto L64
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+136)) = int32(268435456)
	v294 = v272 + int32(4)
	goto L71
L69:
	;
	m.G0 = v272 + int32(144)
	goto L63
L71:
	;
	goto L72
L72:
	;
	if v294 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v305 = F___memcpy(m, int32(4682684), v294, int32(140))
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v362 = int32(916)
	v364 = m.G0
	v366 = v364 - int32(144)
	m.G0 = v366
	switch int32(918) {
	case 0, 2:
		v376 = v362
		goto L90
	default:
		goto L91
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v329
	F_sigemptyset(m, v319+int32(8))
	mBase = m.M
	goto L80
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[620])) = v315
	v329 = int32(4730)
	goto L77
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+136)) = int32(268435456)
	v341 = v319 + int32(4)
	goto L84
L82:
	;
	m.G0 = v319 + int32(144)
	goto L76
L84:
	;
	goto L85
L85:
	;
	if v341 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v352 = F___memcpy(m, int32(4682264), v341, int32(140))
	mBase = m.M
	goto L88
L87:
	;
	goto L88
L88:
	;
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[780])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v412 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v416 = F_LWLockAcquire(m, v412+int32(6784), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L102
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v376
	F_sigemptyset(m, v366+int32(8))
	mBase = m.M
	goto L93
L91:
	;
	*(*int32)(unsafe.Add(mBase, _consts[635])) = v362
	v376 = int32(4730)
	goto L90
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+136)) = int32(268435456)
	v388 = v366 + int32(4)
	goto L97
L95:
	;
	m.G0 = v366 + int32(144)
	goto L89
L97:
	;
	goto L98
L98:
	;
	if v388 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v399 = F___memcpy(m, int32(4682544), v388, int32(140))
	mBase = m.M
	goto L101
L100:
	;
	goto L101
L101:
	;
	goto L95
L102:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v422 = v420 + int32(12)
	v430 = int32(0)
	goto L105
L103:
	;
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v420)))
	*(*int64)(unsafe.Add(mBase, uint32(v420))) = v513 | int64(1)<<(uint(base.I64_extend_i32_u(v511))%64)
	v523 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v420+v511<<(uint(int32(3))%32))+8)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v531 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v531+int32(6784))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L118
	}
L104:
	;
	v506 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v504))) = uint8(v506)
	*(*int32)(unsafe.Add(mBase, _consts[780])) = v505
	v511 = v505
	goto L103
L105:
	;
	v440 = v422 + v430<<(uint(int32(3))%32)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	if v441 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	if v473 != int32(-1) {
		v511 = v473
		goto L103
	} else {
		goto L114
	}
L107:
	;
	v504 = v440
	v505 = v430
	goto L104
L108:
	;
	goto L109
L109:
	;
	v444 = int32(1)
	v445 = v430 | v444
	v448 = v422 + v445<<(uint(int32(3))%32)
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	if v449 != v444 {
		v504 = v448
		v505 = v445
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v453 = v430 | int32(2)
	v456 = v422 + v453<<(uint(int32(3))%32)
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v457 != int32(1) {
		v504 = v456
		v505 = v453
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v460 = int32(3)
	v461 = v430 | v460
	v464 = v422 + v461<<(uint(v460)%32)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	if v465 != int32(1) {
		v504 = v464
		v505 = v461
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v469 = v430 + int32(4)
	if v469 != int32(32) {
		v430 = v469
		goto L105
	} else {
		goto L113
	}
L113:
	;
	goto L106
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	F_errmsg_internal(m, int32(85113), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	F_errfinish(m, int32(495374), int32(355), int32(215129))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L3
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	F_on_shmem_exit(m, int32(1069), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v551 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v551
	v556 = F_pg_sprintf(m, v51, int32(488195), v16+int32(-32))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		v1148 = v54
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	v563 = F_strlen(m, v51)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(1070)
	v572 = int32(4508296)
	v573 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v573
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v45
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v16 + int32(-24)
	goto L124
L122:
	;
	v585 = v54
	v586 = v45
	v587 = v42
	v588 = v48
	v589 = int32(0)
	v590 = v39
	v594 = v54
	goto L9
L124:
	;
	goto L122
L125:
	;
	v597 = int32(4510044)
	v599 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v599 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[49])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_EmitErrorReport(m)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_sigprocmask(m, int32(4422616), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L135
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_LWLockReleaseAll(m)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L129
	}
L129:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if v620 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v621
	v629 = int32(4510052)
	v631 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v631 + int32(1)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_pgaio_io_process_completion(m, v635, int32(0)-v636)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_proc_exit(m, int32(1))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L134
	}
L133:
	;
	v646 = int32(4510052)
	v648 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v648 - int32(1)
	goto L132
L134:
	;
	goto L3
L135:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _consts[339]))
	if v674 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	goto L139
L137:
	;
	goto L138
L138:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L204
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v698 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v702 = F_LWLockAcquire(m, v698+int32(6784), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+12))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v705)+8))
	if v706 != v707 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v1085 != 0 {
		goto L195
	} else {
		goto L196
	}
L143:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v765)))
	v769 = int64(*(*uint32)(unsafe.Add(mBase, _consts[780])))
	v771 = v766 & base.I64_rotl(int64(-2), v769)
	*(*int64)(unsafe.Add(mBase, uint32(v765))) = v771
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v705)+8))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v705)+12))
	if base.Ui32(v773) < base.Ui32(v774) {
		goto L157
	} else {
		goto L158
	}
L144:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v705+v706<<(uint(int32(2))%32))+16))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v714 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v705)+12)) = (v713 - v714) & (v706 + v714)
	if v712 != int32(-1) {
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v725 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	v728 = int64(*(*uint32)(unsafe.Add(mBase, _consts[780])))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v725 | int64(1)<<(uint(v728)%64)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v738 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v738+int32(6784))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v749 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v753 = F_WaitLatch(m, v749, int32(33), int32(-1), int32(83886086))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v761 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = int32(0)
	goto L150
L150:
	;
	goto L142
L151:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+24))
	v920 = v917 + v712<<(uint(int32(7))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v586)+8)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v930 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L175
	}
L152:
	;
	v878 = int32(0)
	goto L171
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v859 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v859+int32(6784))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L169
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v848 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v848+int32(6784))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L168
	}
L155:
	;
	v804 = int32(0)
	v810 = v771
	goto L163
L156:
	;
	if v789 <= int32(0) {
		goto L154
	} else {
		goto L162
	}
L157:
	;
	v776 = int32(2)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	if base.Ui32(v776) <= base.Ui32(v777+(v773-v774)) {
		v792 = v776
		goto L155
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v784 = int32(2)
	v785 = v773 - v774
	if base.Ui32(v784) <= base.Ui32(v785) {
		v792 = v784
		goto L155
	} else {
		goto L161
	}
L160:
	;
	v789 = v773 + v777 - v774
	goto L156
L161:
	;
	v789 = v785
	goto L156
L162:
	;
	v792 = v789
	goto L155
L163:
	;
	if v810 == int64(0) {
		goto L153
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v837 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v837+int32(6784))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L167
	}
L165:
	;
	v815 = base.I64_ctz(v810)
	v817 = v810 & base.I64_rotl(int64(-2), v815)
	*(*int64)(unsafe.Add(mBase, uint32(v765))) = v817
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v765+int32(8)+base.I32_wrap_i64(v815)<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v585+v804<<(uint(int32(2))%32)))) = v826
	v829 = v804 + int32(1)
	if v829 != v792 {
		v804 = v829
		v810 = v817
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v866 = v792
	goto L152
L168:
	;
	goto L151
L169:
	;
	if v804 == int32(0) {
		goto L151
	} else {
		goto L170
	}
L170:
	;
	v866 = v804
	goto L152
L171:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v585+v878<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_SetLatch(m, v889)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L173
	}
L172:
	;
	goto L151
L173:
	;
	v898 = v878 + int32(1)
	if v898 != v866 {
		v878 = v898
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	if v930 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_errhidestmt(m)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v1028 = int32(4510044)
	v1030 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v1030 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v588))) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+1)))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1041<<(uint(int32(2))%32))+uint32(_consts[782])))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)))
	m.T0[v1047].(func(*base.Module, int32))(m, v920)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L193
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_errhidecontext(m)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v952 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)+24))
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+2)))
	if base.Ui32(v963) <= base.Ui32(int32(2)) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+1)))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v977<<(uint(int32(2))%32))+uint32(_consts[782])))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	goto L186
L183:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v963<<(uint(int32(2))%32))+uint32(_consts[783])))
	v971 = v970
	goto L185
L184:
	;
	v971 = int32(0)
	goto L185
L185:
	;
	goto L182
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920))))
	if base.Ui32(v990) <= base.Ui32(int32(7)) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	v1005 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1005
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v983
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = (v920 - v953) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(527095), v18)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L191
	}
L188:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v990<<(uint(int32(2))%32))+uint32(_consts[784])))
	v998 = v997
	goto L190
L189:
	;
	v998 = int32(0)
	goto L190
L190:
	;
	goto L187
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_errfinish(m, int32(495374), int32(513), int32(278567))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L192
	}
L192:
	;
	goto L178
L193:
	;
	v1050 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v588))) = v1050
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v1050
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_pgaio_io_perform_synchronously(m, v920)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L194
	}
L194:
	;
	v1061 = int32(4510044)
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	*(*int32)(unsafe.Add(mBase, _consts[115])) = v1063 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v586)+8)) = int32(0)
	goto L142
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_ProcessInterrupts(m)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if v1094 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v585
	*(*int32)(unsafe.Add(mBase, _consts[346])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v590
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		v1148 = v594
		goto L6
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _consts[339]))
	if v1107 == int32(0) {
		goto L139
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	goto L140
L204:
	;
	goto L5
L205:
	;
	v1156 = int32(v1152)
	m.G0 = v1148
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+4))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1156)))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	if v16+int32(-24) == v1163 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	m.ExcPending = 1
	goto L214
L207:
	;
	if v1166 != 0 {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+4))
	v1166 = v1165
	goto L210
L209:
	;
	v1166 = int32(0)
	goto L210
L210:
	;
	goto L207
L211:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v24 = v1171
	v25 = v1169
	v26 = v1168
	v27 = v1170
	v28 = v1158
	v29 = v1167
	v30 = v1166
	v33 = v1148
	goto L1
L212:
	;
	goto L213
L213:
	;
	F___wasm_longjmp(m, v1159, v1158)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	return
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IsPinnedObject(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	v3 = int32(0)
	if l0 == int32(2613) {
		v16 = v3
	} else {
		if base.Ui32(int32(11999)) < base.Ui32(l1) {
			v16 = v3
		} else {
			v16 = (base.B2i32(l0 != int32(2615)) | base.B2i32(l1 != int32(2200))) & base.B2i32(l0 != int32(1262))
		}
	}
	return v16
}
func F_IsReservedName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = int32(0)
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v3 != int32(112) {
		v12 = v2
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v6 != int32(103) {
			v12 = v2
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = base.B2i32(v9 == int32(95))
		}
	}
	return v12
}
func F_IsSquashableConstant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	v4 = l0
	goto L4
L1:
	;
	return v243
L2:
	;
	v243 = int32(1)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v20 = int32(1)
	if base.Ui32(v20) < base.Ui32(v19-v20) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui32(int32(2)) <= base.Ui32(v7-int32(27)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	return base.B2i32(v15 == int32(0))
L6:
	;
	switch v7 - int32(7) {
	case 0:
		goto L2
	case 1:
		goto L9
	default:
		v243 = int32(0)
		goto L1
	case 8:
		goto L3
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v14
	goto L4
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if base.Ui32(int32(10000)) < base.Ui32(v26) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	if v31 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 <= int32(0) {
		v243 = int32(1)
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 == int32(7) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v141 < int32(2) {
		goto L2
	} else {
		goto L56
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[593]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[594]))
	v49 = m.G0
	v52 = v48 - (v49 - int32(1))
	v54 = v52 >> (uint(int32(31)) % 32)
	goto L20
L20:
	;
	if base.B2i32(v46 < v52^v54-v54)&base.B2i32(v48 != int32(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	goto L23
L23:
	;
	v65 = v39
	goto L28
L24:
	;
	if v137 != 0 {
		goto L18
	} else {
		goto L55
	}
L25:
	;
	v137 = v133
	goto L24
L26:
	;
	v133 = int32(1)
	goto L25
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v80 = int32(1)
	if base.Ui32(v80) < base.Ui32(v79-v80) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if base.Ui32(int32(2)) <= base.Ui32(v68-int32(27)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v137 = base.B2i32(v76 == int32(0))
	goto L24
L30:
	;
	switch v68 - int32(7) {
	case 0:
		goto L26
	case 1:
		goto L33
	default:
		v133 = int32(0)
		goto L25
	case 8:
		goto L27
	}
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v65 = v75
	goto L28
L32:
	;
	goto L29
L33:
	;
	goto L32
L34:
	;
	v137 = int32(0)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if base.Ui32(int32(10000)) < base.Ui32(v85) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = int32(0)
	goto L24
L38:
	;
	goto L39
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	if v89 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v93 <= int32(0) {
		v133 = int32(1)
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 == int32(7) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v106 < int32(2) {
		goto L26
	} else {
		goto L48
	}
L43:
	;
	v101 = F_stack_is_too_deep(m)
	mBase = m.M
	if v101 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v137 = int32(0)
	goto L24
L45:
	;
	goto L46
L46:
	;
	v103 = F_IsSquashableConstant(m, v97)
	mBase = m.M
	if v103 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(0)
	goto L24
L48:
	;
	v109 = int32(1)
	goto L49
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v109<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 == int32(7) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v133 = v122
	goto L25
L51:
	;
	v122 = int32(1)
	v124 = v109 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v124 < v125 {
		v109 = v124
		goto L49
	} else {
		goto L54
	}
L52:
	;
	v120 = F_IsSquashableConstant(m, v116)
	mBase = m.M
	if v120 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v137 = int32(0)
	goto L24
L54:
	;
	goto L50
L55:
	;
	return int32(0)
L56:
	;
	v144 = int32(1)
	goto L57
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v144<<(uint(int32(2))%32))))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v152 == int32(7) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v243 = v232
	goto L1
L59:
	;
	v232 = int32(1)
	v234 = v144 + v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v234 < v235 {
		v144 = v234
		goto L57
	} else {
		goto L93
	}
L60:
	;
	v157 = v151
	goto L65
L61:
	;
	if v229 != 0 {
		goto L59
	} else {
		goto L92
	}
L62:
	;
	v229 = v225
	goto L61
L63:
	;
	v225 = int32(1)
	goto L62
L64:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v172 = int32(1)
	if base.Ui32(v172) < base.Ui32(v171-v172) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if base.Ui32(int32(2)) <= base.Ui32(v160-int32(27)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v229 = base.B2i32(v168 == int32(0))
	goto L61
L67:
	;
	switch v160 - int32(7) {
	case 0:
		goto L63
	case 1:
		goto L70
	default:
		v225 = int32(0)
		goto L62
	case 8:
		goto L64
	}
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v157 = v167
	goto L65
L69:
	;
	goto L66
L70:
	;
	goto L69
L71:
	;
	v229 = int32(0)
	goto L61
L72:
	;
	goto L73
L73:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if base.Ui32(int32(10000)) < base.Ui32(v177) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v229 = int32(0)
	goto L61
L75:
	;
	goto L76
L76:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	if v181 == int32(0) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v185 <= int32(0) {
		v225 = int32(1)
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v190 == int32(7) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v198 < int32(2) {
		goto L63
	} else {
		goto L85
	}
L80:
	;
	v193 = F_stack_is_too_deep(m)
	mBase = m.M
	if v193 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v229 = int32(0)
	goto L61
L82:
	;
	goto L83
L83:
	;
	v195 = F_IsSquashableConstant(m, v189)
	mBase = m.M
	if v195 != 0 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v229 = int32(0)
	goto L61
L85:
	;
	v201 = int32(1)
	goto L86
L86:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v201<<(uint(int32(2))%32))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v209 == int32(7) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v225 = v214
	goto L62
L88:
	;
	v214 = int32(1)
	v216 = v201 + v214
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v216 < v217 {
		v201 = v216
		goto L86
	} else {
		goto L91
	}
L89:
	;
	v212 = F_IsSquashableConstant(m, v208)
	mBase = m.M
	if v212 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v229 = int32(0)
	goto L61
L91:
	;
	goto L87
L92:
	;
	return int32(0)
L93:
	;
	goto L58
}
func F_IsThereCollationInNamespace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v14 = F_SearchSysCacheExists(m, int32(15), l0, v12, l1, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			v21 = F_SearchSysCacheExists(m, int32(15), l0, int32(-1), l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errcode(m, int32(290948))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = F_get_namespace_name(m, l1)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
								F_errmsg(m, int32(723652), v7+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(494295), int32(417), int32(418665))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
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
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[460]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v36 = F_get_namespace_name(m, l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(723787), v7)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errfinish(m, int32(494295), int32(407), int32(418665))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
func F_IsThereOpClassInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_SearchSysCacheExists(m, int32(13), l1, l0, l2, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = F_get_am_name(m, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = F_get_namespace_name(m, l2)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v20
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(724020), v7)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_errfinish(m, int32(494221), int32(1843), int32(418612))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_IsThereOpFamilyInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_SearchSysCacheExists(m, int32(41), l1, l0, l2, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errcode(m, int32(290948))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = F_get_am_name(m, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = F_get_namespace_name(m, l2)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v22
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v20
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(723946), v7)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_errfinish(m, int32(494221), int32(1866), int32(418585))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_icregexeqsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v2, v3, v4, v5, v6, v7, int32(3), v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_ind_fetch_func(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v5 = v4 * l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v6))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v5<<(uint(int32(2))%32))))
	return v14
}
func F_indonesian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v500 int32
	_ = v500
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v639 int32
	_ = v639
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = v9
	goto L3
L1:
	;
	if int32(0) <= v126 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v126 = v98
	goto L1
L3:
	;
	if v22 <= v31 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v126 = int32(-1)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v23))))
	if base.Ui32(v40) < base.Ui32(int32(192)) {
		v97 = v40
		v98 = v38
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int32(117) < v97 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v44 = v31 + int32(1)
	if v44 == v22 {
		v97 = v40
		v98 = v38
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v23))))
	v49 = v47 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v40) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v23))))
	v65 = v63 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v40) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v53 = v31 + int32(2)
	if v53 != v22 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v97 = v40<<(uint(int32(6))%32)&int32(1984) | v49
	v98 = int32(2)
	goto L8
L15:
	;
	goto L14
L16:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v69))))
	v97 = v82&int32(63) | (v40<<(uint(int32(18))%32)&int32(1835008) | v49<<(uint(int32(12))%32) | v65<<(uint(int32(6))%32))
	v98 = int32(4)
	goto L8
L17:
	;
	v69 = v31 + int32(3)
	if v69 != v22 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = v40<<(uint(int32(12))%32)&int32(61440) | v49<<(uint(int32(6))%32) | v65
	v98 = int32(3)
	goto L8
L20:
	;
	goto L19
L21:
	;
	v115 = v98 + v31
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	v31 = v115
	goto L3
L22:
	;
	v102 = v97 - int32(97)
	if v102 < int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v102)>>(uint(int32(3))%32)))+uint32(_consts[1294]))))
	if int32(base.Ui32(v108)>>(uint(v102&int32(7))%32))&int32(1) != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L21
L26:
	;
	v130 = v126
	goto L29
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v268 < int32(3) {
		v762 = v2
		goto L57
	} else {
		goto L58
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134 + v130
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v138 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v163 = v153
	goto L33
L30:
	;
	goto L28
L31:
	;
	if int32(0) <= v258 {
		v130 = v258
		goto L29
	} else {
		goto L56
	}
L32:
	;
	v258 = v230
	goto L31
L33:
	;
	if v154 <= v163 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v258 = int32(-1)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v170 = int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v155))))
	if base.Ui32(v172) < base.Ui32(int32(192)) {
		v229 = v172
		v230 = v170
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(117) < v229 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v176 = v163 + int32(1)
	if v176 == v154 {
		v229 = v172
		v230 = v170
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v155))))
	v181 = v179 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v172) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v155))))
	v197 = v195 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v172) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v185 = v163 + int32(2)
	if v185 != v154 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v229 = v172<<(uint(int32(6))%32)&int32(1984) | v181
	v230 = int32(2)
	goto L38
L45:
	;
	goto L44
L46:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v201))))
	v229 = v214&int32(63) | (v172<<(uint(int32(18))%32)&int32(1835008) | v181<<(uint(int32(12))%32) | v197<<(uint(int32(6))%32))
	v230 = int32(4)
	goto L38
L47:
	;
	v201 = v163 + int32(3)
	if v201 != v154 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v229 = v172<<(uint(int32(12))%32)&int32(61440) | v181<<(uint(int32(6))%32) | v197
	v230 = int32(3)
	goto L38
L50:
	;
	goto L49
L51:
	;
	v247 = v230 + v163
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v247
	v163 = v247
	goto L33
L52:
	;
	v234 = v229 - int32(97)
	if v234 < int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v234)>>(uint(int32(3))%32)))+uint32(_consts[1294]))))
	if int32(base.Ui32(v240)>>(uint(v234&int32(7))%32))&int32(1) != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	goto L51
L56:
	;
	goto L30
L57:
	;
	return v762
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v275
	if v275-int32(2) <= v273 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v316 < int32(3) {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	v311 = v267
	v313 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v282 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v275-int32(1)))))
	switch v287 - int32(104) {
	case 0, 6:
		goto L63
	default:
		v311 = v267
		v313 = v282
		goto L59
	}
L63:
	;
	v292 = F_find_among_b(m, l0, int32(4299968), int32(3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	if v292 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = v298
	v313 = v282
	goto L59
L67:
	;
	goto L68
L68:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
	v301 = F_slice_del(m, l0)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	if v301 < int32(0) {
		v762 = v301
		goto L57
	} else {
		goto L70
	}
L70:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	v307 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v306 - v307
	v311 = v305
	v313 = v307
	goto L59
L71:
	;
	return int32(0)
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v314
	v323 = v314 - int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v323 <= v324 {
		v352 = v311
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	if v357 < int32(3) {
		v762 = int32(0)
		goto L57
	} else {
		goto L83
	}
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v323))))
	if base.B2i32(v328 != int32(117))&base.B2i32(v328 != int32(97)) != 0 {
		v352 = v311
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v336 = F_find_among_b(m, l0, int32(4300032), int32(3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L64
	} else {
		goto L77
	}
L77:
	;
	if v336 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = v340
	goto L74
L79:
	;
	goto L80
L80:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v341
	v343 = F_slice_del(m, l0)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L64
	} else {
		goto L81
	}
L81:
	;
	if v343 < int32(0) {
		v762 = v343
		goto L57
	} else {
		goto L82
	}
L82:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v348 - int32(1)
	v352 = v347
	goto L74
L83:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v354
	v363 = v354 + int32(1)
	if v360 <= v363 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v762 = int32(1)
	goto L57
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v734 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L64
	} else {
		goto L186
	}
L86:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365+v363))))
	switch v367 - int32(101) {
	case 0, 4:
		goto L87
	default:
		goto L85
	}
L87:
	;
	v372 = F_find_among(m, l0, int32(4300096), int32(12))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L64
	} else {
		goto L88
	}
L88:
	;
	if v372 == int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v376
	switch v372 - int32(1) {
	case 0:
		goto L98
	case 1:
		goto L97
	case 2:
		goto L96
	case 3:
		goto L95
	case 4:
		goto L92
	case 5:
		goto L91
	default:
		goto L90
	}
L90:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	if v701 < int32(3) {
		goto L84
	} else {
		goto L169
	}
L91:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = int32(3)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+4)) = v562 - int32(1)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L140
L92:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v421 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v423 - v421
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L109
L93:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v414)+4)) = v416 - int32(1)
	goto L90
L94:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v414 = v413
	goto L93
L95:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = int32(3)
	v408 = F_slice_from_s(m, l0, int32(1), int32(2204519))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L64
	} else {
		goto L105
	}
L96:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v395 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v395
	v399 = F_slice_from_s(m, l0, v395, int32(2204518))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L64
	} else {
		goto L103
	}
L97:
	;
	v387 = F_slice_del(m, l0)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L64
	} else {
		goto L101
	}
L98:
	;
	v380 = F_slice_del(m, l0)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L64
	} else {
		goto L99
	}
L99:
	;
	if v380 < int32(0) {
		v762 = v380
		goto L57
	} else {
		goto L100
	}
L100:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = int32(1)
	v414 = v384
	goto L93
L101:
	;
	if v387 < int32(0) {
		v762 = v387
		goto L57
	} else {
		goto L102
	}
L102:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = int32(3)
	v414 = v391
	goto L93
L103:
	;
	if int32(0) <= v399 {
		goto L94
	} else {
		goto L104
	}
L104:
	;
	v762 = v399
	goto L57
L105:
	;
	if v408 < int32(0) {
		v762 = v408
		goto L57
	} else {
		goto L106
	}
L106:
	;
	goto L94
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v427
	if v545 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L108:
	;
	v545 = v538
	goto L107
L109:
	;
	if v440 <= v427 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v538 = int32(0)
	goto L108
L111:
	;
	v545 = int32(-1)
	goto L107
L112:
	;
	goto L113
L113:
	;
	v456 = int32(1)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427+v441))))
	if base.Ui32(v458) < base.Ui32(int32(192)) {
		v515 = v458
		v516 = v456
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if int32(117) < v515 {
		v538 = v516
		goto L108
	} else {
		goto L127
	}
L115:
	;
	v462 = v427 + int32(1)
	if v462 == v440 {
		v515 = v458
		v516 = v456
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+v441))))
	v467 = v465 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v458) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471+v441))))
	v483 = v481 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v458) {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	v471 = v427 + int32(2)
	if v471 != v440 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v515 = v458<<(uint(int32(6))%32)&int32(1984) | v467
	v516 = int32(2)
	goto L114
L121:
	;
	goto L120
L122:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441+v487))))
	v515 = v500&int32(63) | (v458<<(uint(int32(18))%32)&int32(1835008) | v467<<(uint(int32(12))%32) | v483<<(uint(int32(6))%32))
	v516 = int32(4)
	goto L114
L123:
	;
	v487 = v427 + int32(3)
	if v487 != v440 {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v515 = v458<<(uint(int32(12))%32)&int32(61440) | v467<<(uint(int32(6))%32) | v483
	v516 = int32(3)
	goto L114
L126:
	;
	goto L125
L127:
	;
	v520 = v515 - int32(97)
	if v520 < int32(0) {
		v538 = v516
		goto L108
	} else {
		goto L128
	}
L128:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v520)>>(uint(int32(3))%32)))+uint32(_consts[1294]))))
	if int32(base.Ui32(v526)>>(uint(v520&int32(7))%32))&int32(1) == int32(0) {
		v538 = v516
		goto L108
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v516 + v427
	goto L130
L130:
	;
	goto L110
L131:
	;
	v551 = F_slice_from_s(m, l0, int32(1), int32(2204520))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L64
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v555 = F_slice_del(m, l0)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L64
	} else {
		goto L136
	}
L134:
	;
	if int32(0) <= v551 {
		goto L90
	} else {
		goto L135
	}
L135:
	;
	v762 = v551
	goto L57
L136:
	;
	if int32(0) <= v555 {
		goto L90
	} else {
		goto L137
	}
L137:
	;
	v762 = v555
	goto L57
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v566
	if v684 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L139:
	;
	v684 = v677
	goto L138
L140:
	;
	if v579 <= v566 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v677 = int32(0)
	goto L139
L142:
	;
	v684 = int32(-1)
	goto L138
L143:
	;
	goto L144
L144:
	;
	v595 = int32(1)
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566+v580))))
	if base.Ui32(v597) < base.Ui32(int32(192)) {
		v654 = v597
		v655 = v595
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if int32(117) < v654 {
		v677 = v655
		goto L139
	} else {
		goto L158
	}
L146:
	;
	v601 = v566 + int32(1)
	if v601 == v579 {
		v654 = v597
		v655 = v595
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601+v580))))
	v606 = v604 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v597) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v580))))
	v622 = v620 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v597) {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	v610 = v566 + int32(2)
	if v610 != v579 {
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v654 = v597<<(uint(int32(6))%32)&int32(1984) | v606
	v655 = int32(2)
	goto L145
L152:
	;
	goto L151
L153:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580+v626))))
	v654 = v639&int32(63) | (v597<<(uint(int32(18))%32)&int32(1835008) | v606<<(uint(int32(12))%32) | v622<<(uint(int32(6))%32))
	v655 = int32(4)
	goto L145
L154:
	;
	v626 = v566 + int32(3)
	if v626 != v579 {
		goto L153
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v654 = v597<<(uint(int32(12))%32)&int32(61440) | v606<<(uint(int32(6))%32) | v622
	v655 = int32(3)
	goto L145
L157:
	;
	goto L156
L158:
	;
	v659 = v654 - int32(97)
	if v659 < int32(0) {
		v677 = v655
		goto L139
	} else {
		goto L159
	}
L159:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v659)>>(uint(int32(3))%32)))+uint32(_consts[1294]))))
	if int32(base.Ui32(v665)>>(uint(v659&int32(7))%32))&int32(1) == int32(0) {
		v677 = v655
		goto L139
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v655 + v566
	goto L161
L161:
	;
	goto L141
L162:
	;
	v690 = F_slice_from_s(m, l0, int32(1), int32(2204521))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L64
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v694 = F_slice_del(m, l0)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L64
	} else {
		goto L167
	}
L165:
	;
	if int32(0) <= v690 {
		goto L90
	} else {
		goto L166
	}
L166:
	;
	v762 = v690
	goto L57
L167:
	;
	if v694 < int32(0) {
		v762 = v694
		goto L57
	} else {
		goto L168
	}
L168:
	;
	goto L90
L169:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v704
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v706
	v708 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L64
	} else {
		goto L170
	}
L170:
	;
	if v708 == int32(0) {
		goto L84
	} else {
		goto L171
	}
L171:
	;
	if v708 < int32(0) {
		v762 = v708
		goto L57
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v704
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	if v716 < int32(3) {
		goto L84
	} else {
		goto L173
	}
L173:
	;
	v719 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L64
	} else {
		goto L175
	}
L174:
	;
	if int32(0) <= v719 {
		goto L84
	} else {
		goto L179
	}
L175:
	;
	if v719 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v724 = int32(base.Ui32(v719) >> (uint(int32(31)) % 32))
	goto L178
L177:
	;
	v724 = int32(6)
	goto L178
L178:
	;
	switch v724 {
	case 0, 6:
		goto L84
	default:
		goto L174
	}
L179:
	;
	if v719 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v729 = v719
	goto L182
L181:
	;
	v729 = v313
	goto L182
L182:
	;
	if v719 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v730 = v729
	goto L185
L184:
	;
	v730 = v313
	goto L185
L185:
	;
	return v730
L186:
	;
	if v734 < int32(0) {
		v762 = v734
		goto L57
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+4))
	if v740 < int32(3) {
		goto L84
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v744
	v746 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L64
	} else {
		goto L190
	}
L189:
	;
	if int32(0) <= v746 {
		goto L194
	} else {
		goto L195
	}
L190:
	;
	if v746 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v751 = int32(base.Ui32(v746) >> (uint(int32(31)) % 32))
	goto L193
L192:
	;
	v751 = int32(8)
	goto L193
L193:
	;
	switch v751 {
	case 0, 8:
		goto L84
	default:
		goto L189
	}
L194:
	;
	v755 = int32(1)
	goto L196
L195:
	;
	v755 = v746
	goto L196
L196:
	;
	return v755
}
func F_init_degree_constants(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 float64
	_ = v14
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v48 float64
	_ = v48
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v67 float64
	_ = v67
	var v71 float64
	_ = v71
	var v80 float64
	_ = v80
	var v89 float64
	_ = v89
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v102 float64
	_ = v102
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v127 float64
	_ = v127
	var v138 float64
	_ = v138
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v157 float64
	_ = v157
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v169 float64
	_ = v169
	var v175 float64
	_ = v175
	var v180 float64
	_ = v180
	var v184 float64
	_ = v184
	var v188 float64
	_ = v188
	var v195 int64
	_ = v195
	var v200 int32
	_ = v200
	var v210 float64
	_ = v210
	var v216 float64
	_ = v216
	var v247 float64
	_ = v247
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v250 float64
	_ = v250
	var v264 float64
	_ = v264
	var v281 float64
	_ = v281
	var v288 int32
	_ = v288
	var v291 float64
	_ = v291
	var v296 float64
	_ = v296
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v304 float64
	_ = v304
	var v316 float64
	_ = v316
	var v320 float64
	_ = v320
	var v322 float64
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v342 float64
	_ = v342
	var v346 int32
	_ = v346
	var v347 float64
	_ = v347
	var v348 float64
	_ = v348
	var v354 float64
	_ = v354
	var v355 float64
	_ = v355
	var v357 float64
	_ = v357
	var v359 float64
	_ = v359
	var v361 float64
	_ = v361
	var v371 float64
	_ = v371
	var v373 float64
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v393 float64
	_ = v393
	var v397 int32
	_ = v397
	var v398 float64
	_ = v398
	var v399 float64
	_ = v399
	var v404 float64
	_ = v404
	var v406 float64
	_ = v406
	var v408 float64
	_ = v408
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v419 float64
	_ = v419
	var v422 float64
	_ = v422
	var v426 float64
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v446 float64
	_ = v446
	var v450 int32
	_ = v450
	var v451 float64
	_ = v451
	var v452 float64
	_ = v452
	var v458 float64
	_ = v458
	var v459 float64
	_ = v459
	var v461 float64
	_ = v461
	var v463 float64
	_ = v463
	var v465 float64
	_ = v465
	var v480 float64
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v500 float64
	_ = v500
	var v504 int32
	_ = v504
	var v505 float64
	_ = v505
	var v506 float64
	_ = v506
	var v511 float64
	_ = v511
	var v513 float64
	_ = v513
	var v515 float64
	_ = v515
	var v518 float64
	_ = v518
	var v522 float64
	_ = v522
	var v526 float64
	_ = v526
	var v534 float64
	_ = v534
	var v539 float64
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v559 float64
	_ = v559
	var v563 int32
	_ = v563
	var v564 float64
	_ = v564
	var v565 float64
	_ = v565
	var v570 float64
	_ = v570
	var v572 float64
	_ = v572
	var v574 float64
	_ = v574
	var v577 float64
	_ = v577
	var v581 float64
	_ = v581
	var v585 float64
	_ = v585
	var v593 float64
	_ = v593
	var v603 float64
	_ = v603
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v623 float64
	_ = v623
	var v627 int32
	_ = v627
	var v628 float64
	_ = v628
	var v629 float64
	_ = v629
	var v635 float64
	_ = v635
	var v636 float64
	_ = v636
	var v638 float64
	_ = v638
	var v640 float64
	_ = v640
	var v642 float64
	_ = v642
	var v653 float64
	_ = v653
	var v658 float64
	_ = v658
	var v660 float64
	_ = v660
	var v667 float64
	_ = v667
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v687 float64
	_ = v687
	var v691 int32
	_ = v691
	var v692 float64
	_ = v692
	var v693 float64
	_ = v693
	var v699 float64
	_ = v699
	var v700 float64
	_ = v700
	var v702 float64
	_ = v702
	var v704 float64
	_ = v704
	var v706 float64
	_ = v706
	var v721 float64
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v741 float64
	_ = v741
	var v745 int32
	_ = v745
	var v746 float64
	_ = v746
	var v747 float64
	_ = v747
	var v752 float64
	_ = v752
	var v754 float64
	_ = v754
	var v756 float64
	_ = v756
	var v759 float64
	_ = v759
	var v763 float64
	_ = v763
	var v767 float64
	_ = v767
	var v775 float64
	_ = v775
	var v777 int32
	_ = v777
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*float64)(unsafe.Add(mBase, _consts[954]))
	v20 = base.I64_reinterpret_f64(v14)
	v25 = base.I32_wrap_i64(int64(base.Ui64(v20)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v25) {
		if base.I32_wrap_i64(v20)|(v25-int32(1072693248)) == int32(0) {
			v102 = base.F64_add(base.F64_mul(v14, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		} else {
			v102 = base.F64_div(float64(0), base.F64_sub(v14, v14))
		}
	} else {
		if base.Ui32(v25) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v25+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v94 = v14
				v102 = v94
			} else {
				v48 = F_R(m, base.F64_mul(v14, v14))
				mBase = m.M
				v102 = base.F64_add(base.F64_mul(v14, v48), v14)
			}
		} else {
			v55 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v14)), float64(0.5))
			v56 = base.F64_sqrt(v55)
			v57 = F_R(m, v55)
			mBase = m.M
			if base.Ui32(int32(1072640819)) <= base.Ui32(v25) {
				v62 = base.F64_add(base.F64_mul(v56, v57), v56)
				v89 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v62, v62), float64(-6.123233995736766e-17)))
			} else {
				v67 = float64(0.7853981633974483)
				v71 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v56) & int64(-4294967296))
				v80 = base.F64_div(base.F64_sub(v55, base.F64_mul(v71, v71)), base.F64_add(v56, v71))
				v89 = base.F64_add(base.F64_sub(base.F64_sub(v67, base.F64_add(v71, v71)), base.F64_sub(base.F64_mul(base.F64_add(v56, v56), v57), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v80, v80)))), v67)
			}
			if v20 < int64(0) {
				v93 = base.F64_neg(v89)
			} else {
				v93 = v89
			}
			v94 = v93
			v102 = v94
		}
	}
	*(*float64)(unsafe.Add(mBase, _consts[955])) = v102
	v109 = base.I64_reinterpret_f64(v14)
	v114 = base.I32_wrap_i64(int64(base.Ui64(v109)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v114) {
		if base.I32_wrap_i64(v109)|(v114-int32(1072693248)) == int32(0) {
			if int64(0) <= v109 {
				v127 = float64(0)
			} else {
				v127 = float64(3.141592653589793)
			}
			v184 = v127
		} else {
			v184 = base.F64_div(float64(0), base.F64_sub(v14, v14))
		}
	} else {
		if base.Ui32(v114) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v114) < base.Ui32(int32(1012924417)) {
				v180 = float64(1.5707963267948966)
				v184 = v180
			} else {
				v138 = F_R(m, base.F64_mul(v14, v14))
				mBase = m.M
				v184 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v14, v138)), v14), float64(1.5707963267948966))
			}
		} else {
			if v109 < int64(0) {
				v150 = base.F64_mul(base.F64_add(v14, float64(1)), float64(0.5))
				v151 = base.F64_sqrt(v150)
				v152 = F_R(m, v150)
				mBase = m.M
				v157 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v151, base.F64_add(base.F64_mul(v151, v152), float64(-6.123233995736766e-17))))
				v184 = base.F64_add(v157, v157)
			} else {
				v162 = base.F64_mul(base.F64_sub(float64(1), v14), float64(0.5))
				v163 = base.F64_sqrt(v162)
				v164 = F_R(m, v162)
				mBase = m.M
				v169 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v163) & int64(-4294967296))
				v175 = base.F64_add(base.F64_add(base.F64_mul(v163, v164), base.F64_div(base.F64_sub(v162, base.F64_mul(v169, v169)), base.F64_add(v163, v169))), v169)
				v180 = base.F64_add(v175, v175)
				v184 = v180
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, _consts[956])) = v184
	v188 = *(*float64)(unsafe.Add(mBase, _consts[957]))
	v195 = base.I64_reinterpret_f64(v188)
	v200 = base.I32_wrap_i64(int64(base.Ui64(v195)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1141899264)) <= base.Ui32(v200) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v188)&int64(9223372036854775807)) {
			v210 = v188
		} else {
			v210 = base.F64_copysign(float64(1.5707963267948966), v188)
		}
		v316 = v210
	} else {
		if base.Ui32(v200) <= base.Ui32(int32(1071382527)) {
			if base.Ui32(int32(1044381696)) <= base.Ui32(v200) {
				v247 = v188
				v248 = int32(-1)
				v249 = base.F64_mul(v247, v247)
				v250 = base.F64_mul(v249, v249)
				v264 = base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v281 = base.F64_mul(v249, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v200) <= base.Ui32(int32(1071382527)) {
					v316 = base.F64_sub(v247, base.F64_mul(v247, base.F64_add(v264, v281)))
				} else {
					v288 = v248 << (uint(int32(3)) % 32)
					v291 = *(*float64)(unsafe.Add(mBase, uint32(v288)+uint32(_consts[958])))
					v296 = *(*float64)(unsafe.Add(mBase, uint32(v288)+uint32(_consts[959])))
					v299 = base.F64_sub(v291, base.F64_sub(base.F64_sub(base.F64_mul(v247, base.F64_add(v264, v281)), v296), v247))
					if v195 < int64(0) {
						v303 = base.F64_neg(v299)
					} else {
						v303 = v299
					}
					v304 = v303
					v316 = v304
				}
			} else {
				v304 = v188
				v316 = v304
			}
		} else {
			v216 = base.F64_abs(v188)
			if base.Ui32(v200) <= base.Ui32(int32(1072889855)) {
				if base.Ui32(v200) <= base.Ui32(int32(1072037887)) {
					v247 = base.F64_div(base.F64_add(base.F64_add(v216, v216), float64(-1)), base.F64_add(v216, float64(2)))
					v248 = int32(0)
				} else {
					v247 = base.F64_div(base.F64_add(v216, float64(-1)), base.F64_add(v216, float64(1)))
					v248 = int32(1)
				}
			} else {
				if base.Ui32(v200) <= base.Ui32(int32(1073971199)) {
					v247 = base.F64_div(base.F64_add(v216, float64(-1.5)), base.F64_add(base.F64_mul(v216, float64(1.5)), float64(1)))
					v248 = int32(2)
				} else {
					v247 = base.F64_div(float64(-1), v216)
					v248 = int32(3)
				}
			}
			v249 = base.F64_mul(v247, v247)
			v250 = base.F64_mul(v249, v249)
			v264 = base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
			v281 = base.F64_mul(v249, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, base.F64_add(base.F64_mul(v250, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
			if base.Ui32(v200) <= base.Ui32(int32(1071382527)) {
				v316 = base.F64_sub(v247, base.F64_mul(v247, base.F64_add(v264, v281)))
			} else {
				v288 = v248 << (uint(int32(3)) % 32)
				v291 = *(*float64)(unsafe.Add(mBase, uint32(v288)+uint32(_consts[958])))
				v296 = *(*float64)(unsafe.Add(mBase, uint32(v288)+uint32(_consts[959])))
				v299 = base.F64_sub(v291, base.F64_sub(base.F64_sub(base.F64_mul(v247, base.F64_add(v264, v281)), v296), v247))
				if v195 < int64(0) {
					v303 = base.F64_neg(v299)
				} else {
					v303 = v299
				}
				v304 = v303
				v316 = v304
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, _consts[960])) = v316
	v320 = *(*float64)(unsafe.Add(mBase, _consts[961]))
	v322 = base.F64_mul(v320, float64(0.017453292519943295))
	v326 = m.G0
	v328 = v326 - int32(16)
	m.G0 = v328
	v335 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v322))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v335) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v335) < base.Ui32(int32(1045430272)) {
			v361 = v322
		} else {
			v342 = F___sin(m, v322, float64(0), int32(0))
			mBase = m.M
			v361 = v342
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v335) {
			v361 = base.F64_sub(v322, v322)
		} else {
			v346 = F___rem_pio2(m, v322, v328)
			mBase = m.M
			v347 = *(*float64)(unsafe.Add(mBase, uint32(v328)+8))
			v348 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
			switch v346&int32(3) - int32(1) {
			case 0:
				v355 = F___cos(m, v348, v347)
				mBase = m.M
				v361 = v355
			case 1:
				v357 = F___sin(m, v348, v347, int32(1))
				mBase = m.M
				v361 = base.F64_neg(v357)
			case 2:
				v359 = F___cos(m, v348, v347)
				mBase = m.M
				v361 = base.F64_neg(v359)
			default:
				v354 = F___sin(m, v348, v347, int32(1))
				mBase = m.M
				v361 = v354
			}
		}
	}
	m.G0 = v328 + int32(16)
	*(*float64)(unsafe.Add(mBase, _consts[962])) = v361
	v371 = *(*float64)(unsafe.Add(mBase, _consts[963]))
	v373 = base.F64_mul(v371, float64(0.017453292519943295))
	v377 = m.G0
	v379 = v377 - int32(16)
	m.G0 = v379
	v386 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v373))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v386) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v386) < base.Ui32(int32(1044816030)) {
			v415 = float64(1)
		} else {
			v393 = F___cos(m, v373, float64(0))
			mBase = m.M
			v415 = v393
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v386) {
			v415 = base.F64_sub(v373, v373)
		} else {
			v397 = F___rem_pio2(m, v373, v379)
			mBase = m.M
			v398 = *(*float64)(unsafe.Add(mBase, uint32(v379)+8))
			v399 = *(*float64)(unsafe.Add(mBase, uint32(v379)))
			switch v397&int32(3) - int32(1) {
			case 0:
				v406 = F___sin(m, v399, v398, int32(1))
				mBase = m.M
				v415 = base.F64_neg(v406)
			case 1:
				v408 = F___cos(m, v399, v398)
				mBase = m.M
				v415 = base.F64_neg(v408)
			case 2:
				v411 = F___sin(m, v399, v398, int32(1))
				mBase = m.M
				v415 = v411
			default:
				v404 = F___cos(m, v399, v398)
				mBase = m.M
				v415 = v404
			}
		}
	}
	m.G0 = v379 + int32(16)
	v419 = base.F64_sub(float64(1), v415)
	*(*float64)(unsafe.Add(mBase, _consts[964])) = v419
	v422 = *(*float64)(unsafe.Add(mBase, _consts[965]))
	if base.F64_le(v422, float64(30)) != 0 {
		v426 = base.F64_mul(v422, float64(0.017453292519943295))
		v430 = m.G0
		v432 = v430 - int32(16)
		m.G0 = v432
		v439 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v426))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v439) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v439) < base.Ui32(int32(1045430272)) {
				v465 = v426
			} else {
				v446 = F___sin(m, v426, float64(0), int32(0))
				mBase = m.M
				v465 = v446
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v439) {
				v465 = base.F64_sub(v426, v426)
			} else {
				v450 = F___rem_pio2(m, v426, v432)
				mBase = m.M
				v451 = *(*float64)(unsafe.Add(mBase, uint32(v432)+8))
				v452 = *(*float64)(unsafe.Add(mBase, uint32(v432)))
				switch v450&int32(3) - int32(1) {
				case 0:
					v459 = F___cos(m, v452, v451)
					mBase = m.M
					v465 = v459
				case 1:
					v461 = F___sin(m, v452, v451, int32(1))
					mBase = m.M
					v465 = base.F64_neg(v461)
				case 2:
					v463 = F___cos(m, v452, v451)
					mBase = m.M
					v465 = base.F64_neg(v463)
				default:
					v458 = F___sin(m, v452, v451, int32(1))
					mBase = m.M
					v465 = v458
				}
			}
		}
		m.G0 = v432 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v465
		v534 = base.F64_mul(base.F64_div(v465, v361), float64(0.5))
	} else {
		v480 = base.F64_mul(base.F64_sub(float64(90), v422), float64(0.017453292519943295))
		v484 = m.G0
		v486 = v484 - int32(16)
		m.G0 = v486
		v493 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v480))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v493) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v493) < base.Ui32(int32(1044816030)) {
				v522 = float64(1)
			} else {
				v500 = F___cos(m, v480, float64(0))
				mBase = m.M
				v522 = v500
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v493) {
				v522 = base.F64_sub(v480, v480)
			} else {
				v504 = F___rem_pio2(m, v480, v486)
				mBase = m.M
				v505 = *(*float64)(unsafe.Add(mBase, uint32(v486)+8))
				v506 = *(*float64)(unsafe.Add(mBase, uint32(v486)))
				switch v504&int32(3) - int32(1) {
				case 0:
					v513 = F___sin(m, v506, v505, int32(1))
					mBase = m.M
					v522 = base.F64_neg(v513)
				case 1:
					v515 = F___cos(m, v506, v505)
					mBase = m.M
					v522 = base.F64_neg(v515)
				case 2:
					v518 = F___sin(m, v506, v505, int32(1))
					mBase = m.M
					v522 = v518
				default:
					v511 = F___cos(m, v506, v505)
					mBase = m.M
					v522 = v511
				}
			}
		}
		m.G0 = v486 + int32(16)
		v526 = base.F64_sub(float64(1), v522)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v526
		v534 = base.F64_add(base.F64_mul(base.F64_div(v526, v419), float64(-0.5)), float64(1))
	}
	if base.F64_le(v422, float64(60)) != 0 {
		v539 = base.F64_mul(v422, float64(0.017453292519943295))
		v543 = m.G0
		v545 = v543 - int32(16)
		m.G0 = v545
		v552 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v539))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v552) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v552) < base.Ui32(int32(1044816030)) {
				v581 = float64(1)
			} else {
				v559 = F___cos(m, v539, float64(0))
				mBase = m.M
				v581 = v559
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v552) {
				v581 = base.F64_sub(v539, v539)
			} else {
				v563 = F___rem_pio2(m, v539, v545)
				mBase = m.M
				v564 = *(*float64)(unsafe.Add(mBase, uint32(v545)+8))
				v565 = *(*float64)(unsafe.Add(mBase, uint32(v545)))
				switch v563&int32(3) - int32(1) {
				case 0:
					v572 = F___sin(m, v565, v564, int32(1))
					mBase = m.M
					v581 = base.F64_neg(v572)
				case 1:
					v574 = F___cos(m, v565, v564)
					mBase = m.M
					v581 = base.F64_neg(v574)
				case 2:
					v577 = F___sin(m, v565, v564, int32(1))
					mBase = m.M
					v581 = v577
				default:
					v570 = F___cos(m, v565, v564)
					mBase = m.M
					v581 = v570
				}
			}
		}
		m.G0 = v545 + int32(16)
		v585 = base.F64_sub(float64(1), v581)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v585
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v585
		v593 = base.F64_sub(float64(1), base.F64_mul(base.F64_div(v585, v419), float64(0.5)))
		v658 = v593
		v660 = v593
	} else {
		v603 = base.F64_mul(base.F64_sub(float64(90), v422), float64(0.017453292519943295))
		v607 = m.G0
		v609 = v607 - int32(16)
		m.G0 = v609
		v616 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v603))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v616) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v616) < base.Ui32(int32(1045430272)) {
				v642 = v603
			} else {
				v623 = F___sin(m, v603, float64(0), int32(0))
				mBase = m.M
				v642 = v623
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v616) {
				v642 = base.F64_sub(v603, v603)
			} else {
				v627 = F___rem_pio2(m, v603, v609)
				mBase = m.M
				v628 = *(*float64)(unsafe.Add(mBase, uint32(v609)+8))
				v629 = *(*float64)(unsafe.Add(mBase, uint32(v609)))
				switch v627&int32(3) - int32(1) {
				case 0:
					v636 = F___cos(m, v629, v628)
					mBase = m.M
					v642 = v636
				case 1:
					v638 = F___sin(m, v629, v628, int32(1))
					mBase = m.M
					v642 = base.F64_neg(v638)
				case 2:
					v640 = F___cos(m, v629, v628)
					mBase = m.M
					v642 = base.F64_neg(v640)
				default:
					v635 = F___sin(m, v629, v628, int32(1))
					mBase = m.M
					v642 = v635
				}
			}
		}
		m.G0 = v609 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v642
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v642
		v653 = base.F64_mul(base.F64_div(v642, v361), float64(0.5))
		v658 = v653
		v660 = v653
	}
	*(*float64)(unsafe.Add(mBase, _consts[966])) = base.F64_div(v534, v658)
	if base.F64_le(v422, float64(30)) != 0 {
		v667 = base.F64_mul(v422, float64(0.017453292519943295))
		v671 = m.G0
		v673 = v671 - int32(16)
		m.G0 = v673
		v680 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v667))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v680) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v680) < base.Ui32(int32(1045430272)) {
				v706 = v667
			} else {
				v687 = F___sin(m, v667, float64(0), int32(0))
				mBase = m.M
				v706 = v687
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v680) {
				v706 = base.F64_sub(v667, v667)
			} else {
				v691 = F___rem_pio2(m, v667, v673)
				mBase = m.M
				v692 = *(*float64)(unsafe.Add(mBase, uint32(v673)+8))
				v693 = *(*float64)(unsafe.Add(mBase, uint32(v673)))
				switch v691&int32(3) - int32(1) {
				case 0:
					v700 = F___cos(m, v693, v692)
					mBase = m.M
					v706 = v700
				case 1:
					v702 = F___sin(m, v693, v692, int32(1))
					mBase = m.M
					v706 = base.F64_neg(v702)
				case 2:
					v704 = F___cos(m, v693, v692)
					mBase = m.M
					v706 = base.F64_neg(v704)
				default:
					v699 = F___sin(m, v693, v692, int32(1))
					mBase = m.M
					v706 = v699
				}
			}
		}
		m.G0 = v673 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v706
		v775 = base.F64_mul(base.F64_div(v706, v361), float64(0.5))
	} else {
		v721 = base.F64_mul(base.F64_sub(float64(90), v422), float64(0.017453292519943295))
		v725 = m.G0
		v727 = v725 - int32(16)
		m.G0 = v727
		v734 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v721))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v734) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v734) < base.Ui32(int32(1044816030)) {
				v763 = float64(1)
			} else {
				v741 = F___cos(m, v721, float64(0))
				mBase = m.M
				v763 = v741
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v734) {
				v763 = base.F64_sub(v721, v721)
			} else {
				v745 = F___rem_pio2(m, v721, v727)
				mBase = m.M
				v746 = *(*float64)(unsafe.Add(mBase, uint32(v727)+8))
				v747 = *(*float64)(unsafe.Add(mBase, uint32(v727)))
				switch v745&int32(3) - int32(1) {
				case 0:
					v754 = F___sin(m, v747, v746, int32(1))
					mBase = m.M
					v763 = base.F64_neg(v754)
				case 1:
					v756 = F___cos(m, v747, v746)
					mBase = m.M
					v763 = base.F64_neg(v756)
				case 2:
					v759 = F___sin(m, v747, v746, int32(1))
					mBase = m.M
					v763 = v759
				default:
					v752 = F___cos(m, v747, v746)
					mBase = m.M
					v763 = v752
				}
			}
		}
		m.G0 = v727 + int32(16)
		v767 = base.F64_sub(float64(1), v763)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v767
		v775 = base.F64_add(base.F64_mul(base.F64_div(v767, v419), float64(-0.5)), float64(1))
	}
	v777 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[967])) = uint8(v777)
	*(*float64)(unsafe.Add(mBase, _consts[968])) = base.F64_div(v660, v775)
	m.G0 = v10 + int32(16)
	return
}
func F_init_execution_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v655 int32
	_ = v655
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ReleaseCachedPlan(m, v21, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L3
L6:
	;
	m.G0 = v19 + int32(32)
	return v881
L7:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+68))
	if v476 < v478 {
		goto L116
	} else {
		goto L117
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L112
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L107
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L103
	}
L11:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+76))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v278 + int32(1)
	v286 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v292 = F_GetCachedPlan(m, v282, v290, v288, v286)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L58
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v37 = v35
	goto L14
L13:
	;
	v37 = int32(0)
	goto L14
L14:
	;
	if v37 <= v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	if v39 <= v32 {
		v881 = v29
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v255 + int32(1)
	goto L11
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v41 + int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+72)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v52 = v50
	goto L21
L20:
	;
	v52 = int32(0)
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48+v52<<(uint(int32(2))%32))))
	v57 = F_copyObjectImpl(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
	if v45 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v209 = int32(0)
	v212 = base.B2i32(v52+int32(1) < v46)
	if v212 == v209 {
		goto L49
	} else {
		goto L50
	}
L24:
	;
	if v153 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L25:
	;
	v62 = F_CreateCommandTag(m, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v137 = F_CreateCommandTag(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v70 = F_AllocSetContextCreateInternal(m, v65, int32(414447), int32(0), int32(1024), int32(8388608))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v72 = int32(4515392)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v70
	v77 = F_palloc0(m, int32(144))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(195726186)
	v82 = F_copyObjectImpl(m, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v82
	v87 = F_pstrdup(m, v59)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v87
	goto L33
L33:
	;
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+52)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v62
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+20)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+28)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+36)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+41)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+60)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v77)+56)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v77)+68)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+76)) = v94
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+84)) = uint16(v91)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+88)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v77)+96)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v77)+120)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+112)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+128)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v94
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v70
	v125 = F_copyObjectImpl(m, v57)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v125
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v73
	F_AcquireRewriteLocks(m, v57, int32(1), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v134 = F_pg_rewrite_query(m, v57)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v150 = v77
	v153 = v134
	goto L24
L37:
	;
	v139 = F_CreateCachedPlan(m, v57, v59, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v145 = F_pg_analyze_and_rewrite_withcb(m, v57, v141, int32(474), v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v150 = v139
	v153 = v145
	goto L24
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v156 <= int32(0) {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v162 = int32(0)
	goto L42
L42:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v159+v162<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v181 != int32(6) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L23
L44:
	;
	v191 = v162 + int32(1)
	if v156 != v191 {
		v162 = v191
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v185 != int32(213) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	if v188 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L43
L49:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+58)))
	v219 = F_check_sql_stmt_retval(m, v153, v215, v216, v217, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v222 = v209
	goto L51
L51:
	;
	v223 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	F_CompleteCachedPlan(m, v150, v153, v223, v223, v223, int32(474), v227, int32(2052), v223)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+56)) = uint8(v219)
	v222 = v33
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+40)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v150)+36)) = int32(687)
	v235 = int32(4515392)
	v236 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v241 = F_lappend(m, v240, v150)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v241
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v236
	F_SaveCachedPlan(m, v150)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v52+int32(1) < v46 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	F_MemoryContextDelete(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(0)
	goto L11
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v292
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v295 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v297 = v296
	goto L61
L60:
	;
	v297 = v286
	goto L61
L61:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v298 < v297 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v300 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v318 = v295
	goto L64
L64:
	;
	if v318 != 0 {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v318 = v316
	goto L64
L66:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v306 = F_MemoryContextAlloc(m, v303, v297*int32(20))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v310 = F_repalloc(m, v300, v297*int32(20))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	v312 = v306
	goto L65
L70:
	;
	v312 = v310
	goto L65
L71:
	;
	v325 = int32(0)
	v329 = v319
	v331 = v325
	v333 = v325
	goto L76
L72:
	;
	v319 = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v319 < v320 {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v466 = int32(0)
	goto L7
L75:
	;
	goto L74
L76:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v329<<(uint(int32(2))%32))))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v348 != int32(6) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v466 = v405
	goto L7
L78:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+57)))
	if v384 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L79:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347)+88))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	if v352 != int32(157) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v352 != int32(225) {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v351)+20))
	if v378 == int32(0) {
		goto L8
	} else {
		goto L89
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v347)+88))
	v365 = F_CreateCommandName(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v365
	F_errmsg(m, int32(254138), v19+int32(16))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(493767), int32(744), int32(351688))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	goto L78
L90:
	;
	v387 = F_CommandIsReadOnly(m, v347)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v394 = v391 + v329*int32(20)
	if v331 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	if v387 == int32(0) {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = int64(0)
	v399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+16)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = v347
	*(*uint16)(unsafe.Add(mBase, uint32(v394)+8)) = uint16(v399)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+26)))
	if v404 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v394
	goto L95
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v394
	goto L95
L99:
	;
	v405 = v394
	goto L101
L100:
	;
	v405 = v333
	goto L101
L101:
	;
	v407 = v329 + int32(1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v407 < v408 {
		v329 = v407
		v331 = v394
		v333 = v405
		goto L76
	} else {
		goto L102
	}
L102:
	;
	goto L77
L103:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(141262), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(493767), int32(2075), int32(95457))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v433 = F_CreateCommandName(m, v347)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v433
	F_errmsg(m, int32(253660), v19)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(493767), int32(752), int32(351688))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(254092), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(493767), int32(737), int32(351688))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v881 = int32(1)
	goto L6
L117:
	;
	goto L118
L118:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)+48))
	if v481 == int32(2278) {
		v835 = v477
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v849 = int32(1)
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+55)))
	if v850 != v849 {
		goto L205
	} else {
		goto L206
	}
L120:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v484 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+24))
	if v485 == v487 {
		v835 = v477
		goto L119
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v489 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L123
L125:
	;
	v504 = int32(4515392)
	v505 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v503
	v510 = F_MakeSingleTupleTableSlot(m, int32(0), int32(1617852))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L131
	}
L126:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v497 = F_AllocSetContextCreateInternal(m, v492, int32(215546), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_MemoryContextReset(m, v489)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v497
	v503 = v497
	goto L125
L130:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v503 = v502
	goto L125
L131:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v282)+60))
	if v512 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)+60))
	if v671 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L133:
	;
	v655 = int32(0)
	goto L132
L134:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v515 <= int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v519 = v515 & int32(3)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v512)+12))
	if base.Ui32(v515) < base.Ui32(int32(4)) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if v519 != 0 {
		goto L155
	} else {
		goto L156
	}
L137:
	;
	v524 = int32(0)
	v568 = v524
	v569 = v524
	goto L136
L138:
	;
	goto L139
L139:
	;
	v528 = int32(0)
	v532 = v528
	v533 = v528
	v538 = v528
	goto L140
L140:
	;
	v549 = v520 + v532<<(uint(int32(2))%32)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+12))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v549)+8))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+24)))
	if v554 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v568 = v563
	v569 = v561
	goto L136
L142:
	;
	v555 = v553
	goto L144
L143:
	;
	v555 = v533
	goto L144
L144:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+24)))
	if v556 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v557 = v552
	goto L147
L146:
	;
	v557 = v555
	goto L147
L147:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+24)))
	if v558 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v559 = v551
	goto L150
L149:
	;
	v559 = v557
	goto L150
L150:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+24)))
	if v560 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v561 = v550
	goto L153
L152:
	;
	v561 = v559
	goto L153
L153:
	;
	v562 = int32(4)
	v563 = v532 + v562
	v565 = v538 + v562
	if v565 != v515&int32(2147483644) {
		v532 = v563
		v533 = v561
		v538 = v565
		goto L140
	} else {
		goto L154
	}
L154:
	;
	goto L141
L155:
	;
	v584 = v568
	v585 = v569
	v587 = int32(0)
	goto L158
L156:
	;
	v612 = v569
	goto L157
L157:
	;
	if v612 == int32(0) {
		goto L133
	} else {
		goto L164
	}
L158:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v520+v584<<(uint(int32(2))%32))))
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+24)))
	if v603 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v612 = v604
	goto L157
L160:
	;
	v604 = v602
	goto L162
L161:
	;
	v604 = v585
	goto L162
L162:
	;
	v605 = int32(1)
	v608 = v587 + v605
	if v608 != v519 {
		v584 = v584 + v605
		v585 = v604
		v587 = v608
		goto L158
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if v628 == int32(1) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v612)+76))
	v655 = v631
	goto L132
L166:
	;
	goto L167
L167:
	;
	if base.Ui32(int32(3)) < base.Ui32(v628-int32(2)) {
		goto L133
	} else {
		goto L168
	}
L168:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v612)+96))
	if v636 != 0 {
		v655 = v636
		goto L132
	} else {
		goto L169
	}
L169:
	;
	goto L133
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v814)+4)) = int32(0)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+56)))
	if v819 == int32(1) {
		goto L201
	} else {
		goto L202
	}
L171:
	;
	v796 = F_ExecInitJunkFilter(m, v655, v510)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L200
	}
L172:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+56)))
	if v674 != int32(1) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v677 = int32(0)
	if v510 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	if int32(0) < v686 {
		goto L180
	} else {
		goto L181
	}
L175:
	;
	F_ExecSetSlotDescriptor(m, v510, v671)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v683 = F_MakeSingleTupleTableSlot(m, v671, int32(1617748))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L179
	}
L178:
	;
	v685 = v510
	goto L174
L179:
	;
	v685 = v683
	goto L174
L180:
	;
	v691 = F_palloc0(m, v686<<(uint(int32(1))%32))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L183
	}
L181:
	;
	v776 = v677
	goto L182
L182:
	;
	v788 = F_palloc0(m, int32(20))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L199
	}
L183:
	;
	if v655 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v655)+12))
	v694 = v693
	goto L186
L185:
	;
	v694 = v677
	goto L186
L186:
	;
	v699 = v677
	v701 = v694
	goto L187
L187:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+int32(29)+v699<<(uint(int32(4))%32)))))
	if v716 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v776 = v691
	goto L182
L189:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v655)+12))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v728 = v701
	goto L192
L190:
	;
	v756 = v701
	goto L191
L191:
	;
	v769 = v699 + int32(1)
	if v769 != v686 {
		v699 = v769
		v701 = v756
		goto L187
	} else {
		goto L198
	}
L192:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v742 = v728 + int32(4)
	if base.Ui32(v742) < base.Ui32(v719+v720<<(uint(int32(2))%32)) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v750 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v691+v699<<(uint(int32(1))%32)))) = uint16(v750)
	v756 = v745
	goto L191
L194:
	;
	v745 = v742
	goto L196
L195:
	;
	v745 = int32(0)
	goto L196
L196:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+26)))
	if v746 != 0 {
		v728 = v745
		goto L192
	} else {
		goto L197
	}
L197:
	;
	goto L193
L198:
	;
	goto L188
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v788)+16)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v788)+12)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v788)+8)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v788)+4)) = v655
	*(*int32)(unsafe.Add(mBase, uint32(v788))) = int32(385)
	v814 = v788
	goto L170
L200:
	;
	v814 = v796
	goto L170
L201:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+16))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+12))
	v825 = F_BlessTupleDesc(m, v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v828
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v505
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v835 = v832
	goto L119
L204:
	;
	goto L203
L205:
	;
	if v466 == int32(0) {
		v881 = v849
		goto L6
	} else {
		goto L210
	}
L206:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+56)))
	if v853 != 0 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v835)+48))
	v855 = F_type_is_rowtype(m, v854)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	if v855 == int32(0) {
		goto L205
	} else {
		goto L209
	}
L209:
	;
	v859 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v859)
	goto L205
L210:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v863 == int32(0) {
		v881 = v849
		goto L6
	} else {
		goto L211
	}
L211:
	;
	v866 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+8)) = uint8(v866)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v868 != v866 {
		v881 = v849
		goto L6
	} else {
		goto L212
	}
L212:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v466)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	if v872 != int32(1) {
		v881 = v849
		goto L6
	} else {
		goto L213
	}
L213:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871)+25)))
	if v875 != 0 {
		v881 = v849
		goto L6
	} else {
		goto L214
	}
L214:
	;
	v876 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+9)) = uint8(v876)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v876)
	v881 = v849
	goto L6
}
func F_init_sexpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L3
L2:
	;
	v19 = int32(0)
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v24 = F_object_aclcheck(m, int32(1255), l0, v22, int64(128))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = F_get_func_name(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_aclcheck_error(m, v24, int32(19), v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_RunFunctionExecuteHook(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v35 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L48
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(101) <= v36 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v40 = l3 + int32(16)
	F_fmgr_info_cxt(m, l0, v40, l5)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
	v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
	if l6 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l7 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	if v66&int32(1) == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(106837), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if l4 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v83 = F_exprLocation(m, l2)
	mBase = m.M
	F_executor_errposition(m, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errfinish(m, int32(500106), int32(740), int32(206822))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v146)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
	m.G0 = v14 + int32(16)
	return
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L33
L35:
	;
	if v66&int32(1) == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v102 = F_get_expr_result_type(m, v97, v14+int32(12), v14+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v104 = int32(4515392)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = l5
	v108 = int32(1)
	if base.Ui32(v102-v108) <= base.Ui32(v108) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v105
	goto L33
L39:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v113 = F_CreateTupleDescCopy(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	switch v102 {
	case 0:
		goto L45
	default:
		goto L43
	case 3:
		goto L44
	}
L42:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v113
	goto L38
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L38
L44:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v133)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L38
L45:
	;
	v119 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v119
	v123 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_TupleDescInitEntry(m, v119, int32(1), v123, v124, int32(-1), v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v130)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v129
	goto L38
L48:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v160 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v160
	F_errmsg_plural(m, int32(253904), int32(253952), v160, v14)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(500106), int32(721), int32(206822))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_instrumentSortedGroup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9 + int64(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v17 == v3 {
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)))
		v33 = v23
		v34 = v20 - v21
		if v33&int32(255) != base.B2i32(v17 != int32(0)) {
			if v33&int32(1) != 0 {
				v54 = v3
			} else {
				v54 = int32(1)
			}
		} else {
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
			if v34 <= v40 {
				if v33&int32(1) != 0 {
					v54 = v3
				} else {
					v54 = int32(1)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v33)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v34
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v44
				if v33&int32(1) == int32(0) {
					v54 = int32(1)
				} else {
					v54 = v3
				}
			}
		}
	} else {
		v24 = F_LogicalTapeSetBlocks(m, v17)
		mBase = m.M
		v26 = v24 << (uint(int64(13)) % 64)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)))
		if v27 != 0 {
			v33 = v27
			v34 = v26
			if v33&int32(255) != base.B2i32(v17 != int32(0)) {
				if v33&int32(1) != 0 {
					v54 = v3
				} else {
					v54 = int32(1)
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
				if v34 <= v40 {
					if v33&int32(1) != 0 {
						v54 = v3
					} else {
						v54 = int32(1)
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v33)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v34
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v44
					if v33&int32(1) == int32(0) {
						v54 = int32(1)
					} else {
						v54 = v3
					}
				}
			}
		} else {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v28)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v26
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v31
			v54 = v3
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v54
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	v61 = base.I64_div_s(v57+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	switch v63 - int32(3) {
	case 0:
		v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+69)))
		if v68 != 0 {
			v69 = int32(1)
		} else {
			v69 = int32(2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v69
	case 1:
		v74 = v63
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v74
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(8)
	default:
		v74 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v74
	}
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	switch v77 {
	case 0:
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v78 + v79
		v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v78 <= v82 {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v78
		}
	case 1:
		v85 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v85 + v86
		v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		if v85 <= v89 {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v85
		}
	default:
	}
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v93 | v94
	m.G0 = v7 + int32(16)
	return
}
func F_int24ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 <= v3)
}
func F_int24lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 < v2)
}
func F_int28lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 < v3)
}
func F_int28mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v8 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(1136), int32(319824))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int2abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2&int32(65535) == int32(32768) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(401996), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493025), int32(1242), int32(174102))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = base.I32_extend16_s(v2)
		v27 = v25 >> (uint(int32(31)) % 32)
		return v25 ^ v27 - v27
	}
}
func F_int2eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 == v3)
}
func F_int2hashfast(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v2 = base.I32_extend16_s(l0)
	v3 = int32(16)
	v7 = (int32(base.Ui32(v2)>>(uint(v3)%32)) ^ v2) * int32(-2048144789)
	v12 = (int32(base.Ui32(v7)>>(uint(int32(13))%32)) ^ v7) * int32(-1028477387)
	return int32(base.Ui32(v12)>>(uint(v3)%32)) ^ v12
}
func F_int2lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 < v3)
}
func F_int2recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pq_getmsgint(m, v2, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.I32_extend16_s(v4)
	}
}
func F_int2smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_int2vectorin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_palloc0(m, int32(88))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = v14
	v23 = v16
	v24 = int32(32)
	v26 = int32(0)
	goto L3
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.Ui32(v29-int32(9)) < base.Ui32(int32(5)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = v21 + int32(1)
	goto L3
L6:
	;
	if v29 == int32(32) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v11 + int32(48)
	return v145
L9:
	;
	if v26 < v24 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(21)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26<<(uint(int32(3))%32) + int32(96)
	v145 = v23
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v54 = F_strtox_2(m, v21, v11+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	v55 = base.I32_wrap_i64(v54)
	goto L17
L13:
	;
	v45 = v23
	v46 = v24
	goto L12
L14:
	;
	goto L15
L15:
	;
	v43 = F_repalloc(m, v23, v24<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v45 = v43
	v46 = v24 << (uint(int32(1)) % 32)
	goto L12
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v57 == v21 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v126 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45+v26<<(uint(v126)%32))+24)) = uint16(v55)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v21 = v132
	v23 = v45
	v24 = v46
	v26 = v26 + v126
	goto L3
L19:
	;
	v145 = int32(0)
	goto L8
L20:
	;
	F_errsave_finish(m, v13, int32(493025), v121, int32(275360))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v59 = F_errsave_start(m, v13)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if base.B2i32(v74 != int32(68))&base.B2i32(base.Ui32(int32(-65537)) < base.Ui32(v55-int32(32768))) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v59 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(89105)
	F_errmsg(m, int32(724989), v11)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v121 = int32(199)
	goto L20
L28:
	;
	v84 = F_errsave_start(m, v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v101 = int32(32)
	if v100|v101 == v101 {
		goto L18
	} else {
		goto L35
	}
L31:
	;
	if v84 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(89105)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v21
	F_errmsg(m, int32(189891), v11+int32(16))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v121 = int32(205)
	goto L20
L35:
	;
	v105 = F_errsave_start(m, v13)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v105 == int32(0) {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(89105)
	F_errmsg(m, int32(724989), v11+int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v121 = int32(211)
	goto L20
L40:
	;
	goto L19
}
func F_int42ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v2)
}
func F_int42lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 < v3)
}
func F_int42mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402308), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493025), int32(1101), int32(319889))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return v7
	}
}
func F_int42pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402308), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493025), int32(1087), int32(301312))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return v7
	}
}
func F_int48lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 < v3)
}
func F_int48mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v8 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(994), int32(319806))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int48pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v8 + v5
	if base.B2i32(v5 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(980), int32(301229))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int4abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 == int32(-2147483648) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402308), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493025), int32(1228), int32(174084))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = v3 >> (uint(int32(31)) % 32)
		return v3 ^ v25 - v25
	}
}
func F_int4eqfast(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(l0 == l1)
}
func F_int4hashfast(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(16)
	v6 = (int32(base.Ui32(l0)>>(uint(v2)%32)) ^ l0) * int32(-2048144789)
	v11 = (int32(base.Ui32(v6)>>(uint(int32(13))%32)) ^ v6) * int32(-1028477387)
	return int32(base.Ui32(v11)>>(uint(v2)%32)) ^ v11
}
func F_int4lcm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var __phi32 int32
	_ = __phi32
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == v2 {
		v66 = v2
		return v66
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v9 == int32(0) {
			v66 = v2
			return v66
		} else {
			v12 = int32(31)
			v13 = v6 >> (uint(v12) % 32)
			v17 = v9 >> (uint(v12) % 32)
			v20 = base.B2i32(base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)))
			if base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)) {
				v21 = v6
			} else {
				v21 = v9
			}
			if base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)) {
				v22 = v9
			} else {
				v22 = v6
			}
			if v22 != int32(-2147483648) {
				__phi32 = v21
				__phi33 = v22
				v32 = __phi32
				v33 = __phi33
				for {
					v37 = base.I32_rem_s(v33, v32)
					if v37 != 0 {
						__phi32 = v37
						__phi33 = v32
						v32 = __phi32
						v33 = __phi33
						continue
					} else {
						break
					}
					break
				}
				v39 = v32 >> (uint(int32(31)) % 32)
				v47 = v32 ^ v39 - v39
				v48 = base.I32_div_s(v6, v47)
				v51 = base.I64_extend_i32_s(v48) * base.I64_extend_i32_s(v9)
				v55 = base.I32_wrap_i64(v51)
				if base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(32))%64))) != v55>>(uint(int32(31))%32) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(402308), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493025), int32(1360), int32(291676))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
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
					if v55 == int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(402308), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493025), int32(1366), int32(291676))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
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
						v62 = v55 >> (uint(int32(31)) % 32)
						v66 = v55 ^ v62 - v62
						return v66
					}
				}
			} else {
				if v21&int32(2147483647) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(402308), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493025), int32(1292), int32(312334))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					if v21 != int32(-1) {
						__phi32 = v21
						__phi33 = v22
						v32 = __phi32
						v33 = __phi33
						for {
							v37 = base.I32_rem_s(v33, v32)
							if v37 != 0 {
								__phi32 = v37
								__phi33 = v32
								v32 = __phi32
								v33 = __phi33
								continue
							} else {
								break
							}
							break
						}
						v39 = v32 >> (uint(int32(31)) % 32)
						v47 = v32 ^ v39 - v39
					} else {
						v47 = int32(1)
					}
					v48 = base.I32_div_s(v6, v47)
					v51 = base.I64_extend_i32_s(v48) * base.I64_extend_i32_s(v9)
					v55 = base.I32_wrap_i64(v51)
					if base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(32))%64))) != v55>>(uint(int32(31))%32) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(402308), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493025), int32(1360), int32(291676))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
						if v55 == int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(402308), int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493025), int32(1366), int32(291676))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
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
							v62 = v55 >> (uint(int32(31)) % 32)
							v66 = v55 ^ v62 - v62
							return v66
						}
					}
				}
			}
		}
	}
}
func F_int4mod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v4 + int32(1) {
	case 0:
		v27 = int32(0)
		return v27
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(239536), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493025), int32(1168), int32(422343))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = base.I32_rem_s(v25, v4)
		v27 = v26
		return v27
	}
}
func F_int4range_canonical(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v39 == int32(1) {
						v108 = v12
						m.G0 = v9 + int32(32)
						return v108
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
						if v42 != 0 {
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							if v70 != 0 {
								v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v108 = v105
									m.G0 = v9 + int32(32)
									return v108
								}
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
								if v71 != int32(1) {
									v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v108 = v105
										m.G0 = v9 + int32(32)
										return v108
									}
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									if v74 == int32(2147483647) {
										v77 = int32(0)
										v78 = F_errsave_start(m, v16)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											if v78 == int32(0) {
												v108 = v77
												m.G0 = v9 + int32(32)
												return v108
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(402308), int32(0))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															v108 = v77
															m.G0 = v9 + int32(32)
															return v108
														}
													}
												}
											}
										}
									} else {
										v94 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									}
								}
							}
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
							if v43 != 0 {
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v70 != 0 {
									v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v108 = v105
										m.G0 = v9 + int32(32)
										return v108
									}
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
									if v71 != int32(1) {
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										if v74 == int32(2147483647) {
											v77 = int32(0)
											v78 = F_errsave_start(m, v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												if v78 == int32(0) {
													v108 = v77
													m.G0 = v9 + int32(32)
													return v108
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402308), int32(0))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																v108 = v77
																m.G0 = v9 + int32(32)
																return v108
															}
														}
													}
												}
											}
										} else {
											v94 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										}
									}
								}
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
								if v44 == int32(2147483647) {
									v47 = int32(0)
									v48 = F_errsave_start(m, v16)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 == int32(0) {
											v108 = v47
											m.G0 = v9 + int32(32)
											return v108
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(402308), int32(0))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v16, int32(494106), int32(1552), int32(314272))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														v108 = v47
														m.G0 = v9 + int32(32)
														return v108
													}
												}
											}
										}
									}
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v64)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v44 + v64
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v70 != 0 {
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v71 != int32(1) {
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v74 == int32(2147483647) {
												v77 = int32(0)
												v78 = F_errsave_start(m, v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													if v78 == int32(0) {
														v108 = v77
														m.G0 = v9 + int32(32)
														return v108
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(402308), int32(0))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	v108 = v77
																	m.G0 = v9 + int32(32)
																	return v108
																}
															}
														}
													}
												}
											} else {
												v94 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
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
				v23 = F_lookup_type_cache(m, v17, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(370382), v9)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(494106), int32(1776), int32(398868))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
							if v39 == int32(1) {
								v108 = v12
								m.G0 = v9 + int32(32)
								return v108
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
								if v42 != 0 {
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v70 != 0 {
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v71 != int32(1) {
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v74 == int32(2147483647) {
												v77 = int32(0)
												v78 = F_errsave_start(m, v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													if v78 == int32(0) {
														v108 = v77
														m.G0 = v9 + int32(32)
														return v108
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(402308), int32(0))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	v108 = v77
																	m.G0 = v9 + int32(32)
																	return v108
																}
															}
														}
													}
												}
											} else {
												v94 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
												}
											}
										}
									}
								} else {
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
									if v43 != 0 {
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v70 != 0 {
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
											if v71 != int32(1) {
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
												}
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												if v74 == int32(2147483647) {
													v77 = int32(0)
													v78 = F_errsave_start(m, v16)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														if v78 == int32(0) {
															v108 = v77
															m.G0 = v9 + int32(32)
															return v108
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(402308), int32(0))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		v108 = v77
																		m.G0 = v9 + int32(32)
																		return v108
																	}
																}
															}
														}
													}
												} else {
													v94 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
													v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v108 = v105
														m.G0 = v9 + int32(32)
														return v108
													}
												}
											}
										}
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										if v44 == int32(2147483647) {
											v47 = int32(0)
											v48 = F_errsave_start(m, v16)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 == int32(0) {
													v108 = v47
													m.G0 = v9 + int32(32)
													return v108
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402308), int32(0))
														mBase = m.M
														v58 = m.ExcPending
														if v58 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(494106), int32(1552), int32(314272))
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return int32(0)
															} else {
																v108 = v47
																m.G0 = v9 + int32(32)
																return v108
															}
														}
													}
												}
											}
										} else {
											v64 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v64)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v44 + v64
											v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
											if v70 != 0 {
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
												}
											} else {
												v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
												if v71 != int32(1) {
													v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v108 = v105
														m.G0 = v9 + int32(32)
														return v108
													}
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
													if v74 == int32(2147483647) {
														v77 = int32(0)
														v78 = F_errsave_start(m, v16)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															if v78 == int32(0) {
																v108 = v77
																m.G0 = v9 + int32(32)
																return v108
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(402308), int32(0))
																	mBase = m.M
																	v88 = m.ExcPending
																	if v88 != 0 {
																		return int32(0)
																	} else {
																		F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int32(0)
																		} else {
																			v108 = v77
																			m.G0 = v9 + int32(32)
																			return v108
																		}
																	}
																}
															}
														}
													} else {
														v94 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
														v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v108 = v105
															m.G0 = v9 + int32(32)
															return v108
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
			v23 = F_lookup_type_cache(m, v17, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(370382), v9)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494106), int32(1776), int32(398868))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v39 == int32(1) {
							v108 = v12
							m.G0 = v9 + int32(32)
							return v108
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
							if v42 != 0 {
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v70 != 0 {
									v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v108 = v105
										m.G0 = v9 + int32(32)
										return v108
									}
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
									if v71 != int32(1) {
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										if v74 == int32(2147483647) {
											v77 = int32(0)
											v78 = F_errsave_start(m, v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												if v78 == int32(0) {
													v108 = v77
													m.G0 = v9 + int32(32)
													return v108
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402308), int32(0))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																v108 = v77
																m.G0 = v9 + int32(32)
																return v108
															}
														}
													}
												}
											}
										} else {
											v94 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										}
									}
								}
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
								if v43 != 0 {
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v70 != 0 {
										v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v108 = v105
											m.G0 = v9 + int32(32)
											return v108
										}
									} else {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v71 != int32(1) {
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v74 == int32(2147483647) {
												v77 = int32(0)
												v78 = F_errsave_start(m, v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													if v78 == int32(0) {
														v108 = v77
														m.G0 = v9 + int32(32)
														return v108
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(402308), int32(0))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int32(0)
																} else {
																	v108 = v77
																	m.G0 = v9 + int32(32)
																	return v108
																}
															}
														}
													}
												}
											} else {
												v94 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
												}
											}
										}
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									if v44 == int32(2147483647) {
										v47 = int32(0)
										v48 = F_errsave_start(m, v16)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											if v48 == int32(0) {
												v108 = v47
												m.G0 = v9 + int32(32)
												return v108
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(402308), int32(0))
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(494106), int32(1552), int32(314272))
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return int32(0)
														} else {
															v108 = v47
															m.G0 = v9 + int32(32)
															return v108
														}
													}
												}
											}
										}
									} else {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v64)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v44 + v64
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v70 != 0 {
											v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v108 = v105
												m.G0 = v9 + int32(32)
												return v108
											}
										} else {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
											if v71 != int32(1) {
												v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v108 = v105
													m.G0 = v9 + int32(32)
													return v108
												}
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												if v74 == int32(2147483647) {
													v77 = int32(0)
													v78 = F_errsave_start(m, v16)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														if v78 == int32(0) {
															v108 = v77
															m.G0 = v9 + int32(32)
															return v108
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(402308), int32(0))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v16, int32(494106), int32(1565), int32(314272))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		v108 = v77
																		m.G0 = v9 + int32(32)
																		return v108
																	}
																}
															}
														}
													}
												} else {
													v94 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v94)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74 + int32(1)
													v105 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v108 = v105
														m.G0 = v9 + int32(32)
														return v108
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
func F_int64_to_numeric(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v16 = F_palloc(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v16
		v21 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16))) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v16 + int32(2)
		if l0 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v36 = int64(0) - l0
			v39 = v36
			v41 = v2
			v43 = v16 + int32(12)
			for {
				v48 = v43 - int32(2)
				v50 = base.I64_div_u_s(v39, int64(10000))
				v53 = v50*int64(55536) + v39
				*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
				v56 = v41 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v39) {
					v39 = v50
					v41 = v56
					v43 = v48
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v48
			v62 = v56
			v65 = v41
		} else {
			v32 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v32
			if l0 == v32 {
				v62 = v2
				v65 = v2
			} else {
				v36 = l0
				v39 = v36
				v41 = v2
				v43 = v16 + int32(12)
				for {
					v48 = v43 - int32(2)
					v50 = base.I64_div_u_s(v39, int64(10000))
					v53 = v50*int64(55536) + v39
					*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
					v56 = v41 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v39) {
						v39 = v50
						v41 = v56
						v43 = v48
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v48
				v62 = v56
				v65 = v41
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v65
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v62
		v73 = F_make_result_opt_error(m, v11+int32(8), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v16)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(32)
				return v73
			}
		}
	}
}
func F_int82eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 == v4)
}
func F_int82pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v8 + v4
	if base.B2i32(v4 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(1041), int32(301304))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int84(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(v4-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402308), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(1256), int32(558576))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		return base.I32_wrap_i64(v4)
	}
}
func F_int84lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v4)
}
func F_int84mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v8 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(913), int32(319848))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int8abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	if v5 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(554), int32(174066))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v27 = v5 >> (uint(int64(63)) % 64)
		v30 = F_Int64GetDatum(m, v5^v27-v27)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int8range_canonical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v21 == v18 {
				v31 = v20
				F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v40 == int32(1) {
						v117 = v13
						m.G0 = v10 + int32(32)
						return v117
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v43 != 0 {
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
							if v75 != 0 {
								v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v117 = v114
									m.G0 = v10 + int32(32)
									return v117
								}
							} else {
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
								if v76 != int32(1) {
									v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v117 = v114
										m.G0 = v10 + int32(32)
										return v117
									}
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
									if v80 == int64(9223372036854775807) {
										v83 = int32(0)
										v84 = F_errsave_start(m, v17)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											if v84 == int32(0) {
												v117 = v83
												m.G0 = v10 + int32(32)
												return v117
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(402018), int32(0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															v117 = v83
															m.G0 = v10 + int32(32)
															return v117
														}
													}
												}
											}
										}
									} else {
										v102 = F_Int64GetDatum(m, v80+int64(1))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v104 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
											v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v117 = v114
												m.G0 = v10 + int32(32)
												return v117
											}
										}
									}
								}
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
							if v44 != 0 {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
								if v75 != 0 {
									v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v117 = v114
										m.G0 = v10 + int32(32)
										return v117
									}
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
									if v76 != int32(1) {
										v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v117 = v114
											m.G0 = v10 + int32(32)
											return v117
										}
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
										if v80 == int64(9223372036854775807) {
											v83 = int32(0)
											v84 = F_errsave_start(m, v17)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												if v84 == int32(0) {
													v117 = v83
													m.G0 = v10 + int32(32)
													return v117
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402018), int32(0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																v117 = v83
																m.G0 = v10 + int32(32)
																return v117
															}
														}
													}
												}
											}
										} else {
											v102 = F_Int64GetDatum(m, v80+int64(1))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v104 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
												v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v117 = v114
													m.G0 = v10 + int32(32)
													return v117
												}
											}
										}
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
								if v46 == int64(9223372036854775807) {
									v49 = int32(0)
									v50 = F_errsave_start(m, v17)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										if v50 == int32(0) {
											v117 = v49
											m.G0 = v10 + int32(32)
											return v117
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(402018), int32(0))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v17, int32(494106), int32(1599), int32(314252))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														v117 = v49
														m.G0 = v10 + int32(32)
														return v117
													}
												}
											}
										}
									}
								} else {
									v68 = F_Int64GetDatum(m, v46+int64(1))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v70)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v68
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
										if v75 != 0 {
											v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v117 = v114
												m.G0 = v10 + int32(32)
												return v117
											}
										} else {
											v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
											if v76 != int32(1) {
												v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v117 = v114
													m.G0 = v10 + int32(32)
													return v117
												}
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
												if v80 == int64(9223372036854775807) {
													v83 = int32(0)
													v84 = F_errsave_start(m, v17)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														if v84 == int32(0) {
															v117 = v83
															m.G0 = v10 + int32(32)
															return v117
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(402018), int32(0))
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		v117 = v83
																		m.G0 = v10 + int32(32)
																		return v117
																	}
																}
															}
														}
													}
												} else {
													v102 = F_Int64GetDatum(m, v80+int64(1))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v104 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
														v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v117 = v114
															m.G0 = v10 + int32(32)
															return v117
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
				v24 = F_lookup_type_cache(m, v18, int32(2048))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(370382), v10)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(494106), int32(1776), int32(398868))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
						v31 = v24
						F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v40 == int32(1) {
								v117 = v13
								m.G0 = v10 + int32(32)
								return v117
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
								if v43 != 0 {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
									if v75 != 0 {
										v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v117 = v114
											m.G0 = v10 + int32(32)
											return v117
										}
									} else {
										v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
										if v76 != int32(1) {
											v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v117 = v114
												m.G0 = v10 + int32(32)
												return v117
											}
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
											if v80 == int64(9223372036854775807) {
												v83 = int32(0)
												v84 = F_errsave_start(m, v17)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													if v84 == int32(0) {
														v117 = v83
														m.G0 = v10 + int32(32)
														return v117
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(402018), int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	v117 = v83
																	m.G0 = v10 + int32(32)
																	return v117
																}
															}
														}
													}
												}
											} else {
												v102 = F_Int64GetDatum(m, v80+int64(1))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													v104 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
													v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = v114
														m.G0 = v10 + int32(32)
														return v117
													}
												}
											}
										}
									}
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
									if v44 != 0 {
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
										if v75 != 0 {
											v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v117 = v114
												m.G0 = v10 + int32(32)
												return v117
											}
										} else {
											v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
											if v76 != int32(1) {
												v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v117 = v114
													m.G0 = v10 + int32(32)
													return v117
												}
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
												if v80 == int64(9223372036854775807) {
													v83 = int32(0)
													v84 = F_errsave_start(m, v17)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														if v84 == int32(0) {
															v117 = v83
															m.G0 = v10 + int32(32)
															return v117
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(402018), int32(0))
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		v117 = v83
																		m.G0 = v10 + int32(32)
																		return v117
																	}
																}
															}
														}
													}
												} else {
													v102 = F_Int64GetDatum(m, v80+int64(1))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v104 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
														v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v117 = v114
															m.G0 = v10 + int32(32)
															return v117
														}
													}
												}
											}
										}
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
										v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
										if v46 == int64(9223372036854775807) {
											v49 = int32(0)
											v50 = F_errsave_start(m, v17)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												if v50 == int32(0) {
													v117 = v49
													m.G0 = v10 + int32(32)
													return v117
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402018), int32(0))
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(494106), int32(1599), int32(314252))
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return int32(0)
															} else {
																v117 = v49
																m.G0 = v10 + int32(32)
																return v117
															}
														}
													}
												}
											}
										} else {
											v68 = F_Int64GetDatum(m, v46+int64(1))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v70)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v68
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
												if v75 != 0 {
													v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = v114
														m.G0 = v10 + int32(32)
														return v117
													}
												} else {
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
													if v76 != int32(1) {
														v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v117 = v114
															m.G0 = v10 + int32(32)
															return v117
														}
													} else {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
														v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
														if v80 == int64(9223372036854775807) {
															v83 = int32(0)
															v84 = F_errsave_start(m, v17)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																if v84 == int32(0) {
																	v117 = v83
																	m.G0 = v10 + int32(32)
																	return v117
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(402018), int32(0))
																		mBase = m.M
																		v94 = m.ExcPending
																		if v94 != 0 {
																			return int32(0)
																		} else {
																			F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																			mBase = m.M
																			v99 = m.ExcPending
																			if v99 != 0 {
																				return int32(0)
																			} else {
																				v117 = v83
																				m.G0 = v10 + int32(32)
																				return v117
																			}
																		}
																	}
																}
															}
														} else {
															v102 = F_Int64GetDatum(m, v80+int64(1))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v104 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
																*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
																v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	v117 = v114
																	m.G0 = v10 + int32(32)
																	return v117
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
			v24 = F_lookup_type_cache(m, v18, int32(2048))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
						F_errmsg_internal(m, int32(370382), v10)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494106), int32(1776), int32(398868))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
					v31 = v24
					F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v40 == int32(1) {
							v117 = v13
							m.G0 = v10 + int32(32)
							return v117
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v43 != 0 {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
								if v75 != 0 {
									v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v117 = v114
										m.G0 = v10 + int32(32)
										return v117
									}
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
									if v76 != int32(1) {
										v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v117 = v114
											m.G0 = v10 + int32(32)
											return v117
										}
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
										if v80 == int64(9223372036854775807) {
											v83 = int32(0)
											v84 = F_errsave_start(m, v17)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												if v84 == int32(0) {
													v117 = v83
													m.G0 = v10 + int32(32)
													return v117
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(402018), int32(0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																v117 = v83
																m.G0 = v10 + int32(32)
																return v117
															}
														}
													}
												}
											}
										} else {
											v102 = F_Int64GetDatum(m, v80+int64(1))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v104 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
												v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v117 = v114
													m.G0 = v10 + int32(32)
													return v117
												}
											}
										}
									}
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
								if v44 != 0 {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
									if v75 != 0 {
										v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v117 = v114
											m.G0 = v10 + int32(32)
											return v117
										}
									} else {
										v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
										if v76 != int32(1) {
											v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v117 = v114
												m.G0 = v10 + int32(32)
												return v117
											}
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
											if v80 == int64(9223372036854775807) {
												v83 = int32(0)
												v84 = F_errsave_start(m, v17)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													if v84 == int32(0) {
														v117 = v83
														m.G0 = v10 + int32(32)
														return v117
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(402018), int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	v117 = v83
																	m.G0 = v10 + int32(32)
																	return v117
																}
															}
														}
													}
												}
											} else {
												v102 = F_Int64GetDatum(m, v80+int64(1))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													v104 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
													v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = v114
														m.G0 = v10 + int32(32)
														return v117
													}
												}
											}
										}
									}
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
									if v46 == int64(9223372036854775807) {
										v49 = int32(0)
										v50 = F_errsave_start(m, v17)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											if v50 == int32(0) {
												v117 = v49
												m.G0 = v10 + int32(32)
												return v117
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(402018), int32(0))
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v17, int32(494106), int32(1599), int32(314252))
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															v117 = v49
															m.G0 = v10 + int32(32)
															return v117
														}
													}
												}
											}
										}
									} else {
										v68 = F_Int64GetDatum(m, v46+int64(1))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v70)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v68
											v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
											if v75 != 0 {
												v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v117 = v114
													m.G0 = v10 + int32(32)
													return v117
												}
											} else {
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
												if v76 != int32(1) {
													v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = v114
														m.G0 = v10 + int32(32)
														return v117
													}
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
													v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
													if v80 == int64(9223372036854775807) {
														v83 = int32(0)
														v84 = F_errsave_start(m, v17)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															if v84 == int32(0) {
																v117 = v83
																m.G0 = v10 + int32(32)
																return v117
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v90 = m.ExcPending
																if v90 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(402018), int32(0))
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int32(0)
																	} else {
																		F_errsave_finish(m, v17, int32(494106), int32(1612), int32(314252))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			v117 = v83
																			m.G0 = v10 + int32(32)
																			return v117
																		}
																	}
																}
															}
														}
													} else {
														v102 = F_Int64GetDatum(m, v80+int64(1))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v104 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v104)
															*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v102
															v114 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v117 = v114
																m.G0 = v10 + int32(32)
																return v117
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
func F_int8um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500141), int32(448), int32(287519))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v27 = F_Int64GetDatum(m, int64(0)-v4)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_intarray_push_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_intarray_concat_arrays(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_inter_lb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v16
	v22 = int32(1)
	v24 = F_lseg_interpt_line(m, int32(0), v11, v13)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		if v24 != 0 {
			v55 = v22
			m.G0 = v11 + int32(32)
			return v55
		} else {
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v28
			v35 = F_lseg_interpt_line(m, int32(0), v11, v13)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v35 != 0 {
					v55 = v22
					m.G0 = v11 + int32(32)
					return v55
				} else {
					v37 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v38
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v28
					v44 = F_lseg_interpt_line(m, int32(0), v11, v13)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 != 0 {
							v55 = v22
							m.G0 = v11 + int32(32)
							return v55
						} else {
							v46 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v47 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v38
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v37
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v47
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v46
							v53 = F_lseg_interpt_line(m, int32(0), v11, v13)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = v53
								m.G0 = v11 + int32(32)
								return v55
							}
						}
					}
				}
			}
		}
	}
}
func F_internalerrposition(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1095])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(453924), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(498091), int32(1489), int32(249343))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_consts[1105]))) = l0
		return
	}
}
func F_intervaltypmodleastfield(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 < v2 {
		v56 = v2
		m.G0 = v7 + int32(16)
		return v56
	} else {
		v12 = int32(base.Ui32(l0) >> (uint(int32(16)) % 32))
		if base.Ui32(v12) <= base.Ui32(int32(3071)) {
			switch v12 - int32(2) {
			case 0, 4:
				v56 = int32(4)
				m.G0 = v7 + int32(16)
				return v56
			case 1, 3, 5:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(29868), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495818), int32(1248), int32(431417))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 2:
				v56 = int32(5)
				m.G0 = v7 + int32(16)
				return v56
			case 6:
				v56 = int32(3)
				m.G0 = v7 + int32(16)
				return v56
			default:
				switch v12 - int32(1024) {
				case 0, 8:
					v56 = int32(2)
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(29868), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495818), int32(1248), int32(431417))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 == int32(2048) {
						v56 = int32(1)
						m.G0 = v7 + int32(16)
						return v56
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(29868), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495818), int32(1248), int32(431417))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
			if base.Ui32(v12) <= base.Ui32(int32(6143)) {
				switch v12 - int32(3072) {
				case 0, 8:
					v56 = int32(1)
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(29868), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495818), int32(1248), int32(431417))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 != int32(4096) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(29868), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495818), int32(1248), int32(431417))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v56 = v2
						m.G0 = v7 + int32(16)
						return v56
					}
				}
			} else {
				switch v12 - int32(7168) {
				case 0, 8:
					v56 = v2
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(29868), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495818), int32(1248), int32(431417))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 == int32(6144) {
						v56 = v2
						m.G0 = v7 + int32(16)
						return v56
					} else {
						if v12 != int32(32767) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
								F_errmsg_internal(m, int32(29868), v7)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495818), int32(1248), int32(431417))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v56 = v2
							m.G0 = v7 + int32(16)
							return v56
						}
					}
				}
			}
		}
	}
}
func F_is_encoding_supported_by_icu(m *base.Module, l0 int32) int32 {
	var v12 int32
	_ = v12
	if base.Ui32(l0) < base.Ui32(int32(35)) {
		v12 = base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(l0))%64)&int64(34357509982) != int64(0))
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_is_member_of_role_nosuper(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = int32(0)
	v9 = F_roles_is_member_of(m, l0, v6, v6, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v52 = int32(1)
	goto L3
L3:
	;
	return v52
L4:
	;
	return int32(0)
L5:
	;
	v13 = int32(0)
	if v9 == v13 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v52 = v51
	goto L3
L7:
	;
	v51 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v19 <= int32(0) {
		v44 = v13
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v51 = v44
	goto L6
L11:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v25 = v19
	goto L14
L13:
	;
	v25 = v22
	goto L14
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v28 = int32(0)
	goto L15
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28<<(uint(int32(2))%32))))
	v37 = base.B2i32(v36 == l1)
	if v36 == l1 {
		v44 = v37
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v44 = v37
	goto L10
L17:
	;
	v39 = v28 + int32(1)
	if v39 != v25 {
		v28 = v39
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_isn_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	F_ean2string(m, v8, v5, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_pstrdup(m, v5)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(32)
			return v14
		}
	}
}
func F_iso8859_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(-1), int32(6))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v25 = v15 - int32(9)
		if base.Ui32(int32(20)) <= base.Ui32(v25) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
					F_errmsg(m, int32(124191), v10)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(500116), int32(133), int32(555478))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
			if int32(base.Ui32(int32(983551))>>(uint(v25)%32))&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
						F_errmsg(m, int32(124191), v10)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(500116), int32(133), int32(555478))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
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
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_consts[1305])))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v40 = int32(0)
				v45 = F_LocalToUtf(m, v14, v17, v13, v39, v40, v40, v40, v15, base.B2i32(v12 != v40))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v45
				}
			}
		}
	}
}
func F_iso_to_koi8r(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(25), int32(22))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(25), int32(22), int32(2230288), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_iso_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(25), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_latin2mic_with_table(m, v6, v5, v10, int32(139), int32(25), int32(2230288), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_iso_to_win1251(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(25), int32(23))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(25), int32(23), int32(2230416), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_iswpunct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	if base.Ui32(l0) <= base.Ui32(int32(131071)) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[1377]))))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v13<<(uint(int32(5))%32))+uint32(_consts[1377]))))
		v25 = int32(base.Ui32(v19)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		v25 = int32(0)
	}
	return v25
}
func F_iswupper(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = F_casemap(m, l0, int32(0))
	return base.B2i32(v3 != l0)
}
func F_iteratorFromContainer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = F_palloc0(m, int32(40))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = l0 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v15
		v18 = v13 & int32(268435455)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = v20 & int32(1610612736)
		if v22 != int32(536870912) {
			if v22 == int32(1073741824) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15 + v18<<(uint(int32(2))%32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v35 = int32(base.Ui32(v31)>>(uint(int32(28))%32)) & int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(0)
				return v7
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(218765), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(497387), int32(1043), int32(218797))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15 + v18<<(uint(int32(3))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(2)
			return v7
		}
	}
}
