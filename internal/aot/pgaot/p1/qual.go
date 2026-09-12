package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EvalPlanQualFetchRowMark(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+l1<<(uint(int32(2))%32)-int32(4))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if base.Ui32(int32(3)) < base.Ui32(v21) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		if v24 != v25 {
			v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+6)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+6)))
			if v29 < v27 {
				F_slot_getsomeattrs_int(m, v28, v27)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = v27 - int32(1)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v37))))
					if v39 != 0 {
						v194 = v4
						m.G0 = v11 + int32(16)
						return v194
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v36<<(uint(int32(2))%32))))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						if v44 != v45 {
							v194 = v4
							m.G0 = v11 + int32(16)
							return v194
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
							v48 = v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
							if v48 == int32(4) {
								v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
								if v51 < v54 {
									F_slot_getsomeattrs_int(m, v50, v54)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v59 = v54 - int32(1)
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
										if v62 != 0 {
											v194 = int32(0)
											m.G0 = v11 + int32(16)
											return v194
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
											if v71 == int32(102) {
												v74 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
												v77 = F_GetFdwRoutineForRelation(m, v69, v74)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
													if v79 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v220 = m.ExcPending
															if v220 != 0 {
																return int32(0)
															} else {
																v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
																v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
																F_errmsg(m, int32(719478), v11)
																mBase = m.M
																v228 = m.ExcPending
																if v228 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2870), int32(315362))
																	mBase = m.M
																	v233 = m.ExcPending
																	if v233 != 0 {
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
														v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															if l2 != 0 {
																v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
																if v87&int32(2) == int32(0) {
																	v194 = int32(1)
																	m.G0 = v11 + int32(16)
																	return v194
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v95 = m.ExcPending
																	if v95 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg_internal(m, int32(317992), int32(0))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(317992), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
														}
													}
												}
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
												if v106 != 0 {
													v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
													if v108&int32(1) == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(336384), int32(0))
															mBase = m.M
															v241 = m.ExcPending
															if v241 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(326828), int32(1264), int32(271151))
																mBase = m.M
																v246 = m.ExcPending
																if v246 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
														v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															if v116 != 0 {
																v194 = int32(1)
																m.G0 = v11 + int32(16)
																return v194
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(317992), int32(0))
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496951), int32(2893), int32(315362))
																		mBase = m.M
																		v130 = m.ExcPending
																		if v130 != 0 {
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
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														if v116 != 0 {
															v194 = int32(1)
															m.G0 = v11 + int32(16)
															return v194
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2893), int32(315362))
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
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
											}
										}
									}
								} else {
									v59 = v54 - int32(1)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
									if v62 != 0 {
										v194 = int32(0)
										m.G0 = v11 + int32(16)
										return v194
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
										if v71 == int32(102) {
											v74 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
											v77 = F_GetFdwRoutineForRelation(m, v69, v74)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
												if v79 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return int32(0)
														} else {
															v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
															v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
															F_errmsg(m, int32(719478), v11)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2870), int32(315362))
																mBase = m.M
																v233 = m.ExcPending
																if v233 != 0 {
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
													v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														if l2 != 0 {
															v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
															if v87&int32(2) == int32(0) {
																v194 = int32(1)
																m.G0 = v11 + int32(16)
																return v194
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(317992), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
													}
												}
											}
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
											if v106 != 0 {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
												if v108&int32(1) == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(336384), int32(0))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(326828), int32(1264), int32(271151))
															mBase = m.M
															v246 = m.ExcPending
															if v246 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														if v116 != 0 {
															v194 = int32(1)
															m.G0 = v11 + int32(16)
															return v194
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2893), int32(315362))
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
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
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													if v116 != 0 {
														v194 = int32(1)
														m.G0 = v11 + int32(16)
														return v194
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(317992), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2893), int32(315362))
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
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
										}
									}
								}
							} else {
								v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
								if v51 < v131 {
									F_slot_getsomeattrs_int(m, v50, v131)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										v136 = v131 - int32(1)
										v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
										v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
										if v139 != 0 {
											v194 = int32(0)
											m.G0 = v11 + int32(16)
											return v194
										} else {
											v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
											v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
											v146 = m.G0
											v148 = v146 - int32(32)
											m.G0 = v148
											v151 = v148 + int32(20)
											*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
											v154 = F_pg_detoast_datum(m, v145)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
												*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
												v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
												*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
												v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
												*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
												v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
												m.T0[v166].(func(*base.Module, int32))(m, l2)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int32(0)
												} else {
													v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
													F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
														v178 = v176 & int32(65533)
														*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
														v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
														*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
														m.G0 = v148 + int32(32)
														v194 = int32(1)
														m.G0 = v11 + int32(16)
														return v194
													}
												}
											}
										}
									}
								} else {
									v136 = v131 - int32(1)
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
									if v139 != 0 {
										v194 = int32(0)
										m.G0 = v11 + int32(16)
										return v194
									} else {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
										v146 = m.G0
										v148 = v146 - int32(32)
										m.G0 = v148
										v151 = v148 + int32(20)
										*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
										v154 = F_pg_detoast_datum(m, v145)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int32(0)
										} else {
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
											v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
											*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
											*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
											v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
											m.T0[v166].(func(*base.Module, int32))(m, l2)
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int32(0)
											} else {
												v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
													v178 = v176 & int32(65533)
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
													v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
													m.G0 = v148 + int32(32)
													v194 = int32(1)
													m.G0 = v11 + int32(16)
													return v194
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
				v36 = v27 - int32(1)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v37))))
				if v39 != 0 {
					v194 = v4
					m.G0 = v11 + int32(16)
					return v194
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v36<<(uint(int32(2))%32))))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
					if v44 != v45 {
						v194 = v4
						m.G0 = v11 + int32(16)
						return v194
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
						v48 = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
						if v48 == int32(4) {
							v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
							if v51 < v54 {
								F_slot_getsomeattrs_int(m, v50, v54)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v59 = v54 - int32(1)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
									if v62 != 0 {
										v194 = int32(0)
										m.G0 = v11 + int32(16)
										return v194
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
										if v71 == int32(102) {
											v74 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
											v77 = F_GetFdwRoutineForRelation(m, v69, v74)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
												if v79 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v220 = m.ExcPending
														if v220 != 0 {
															return int32(0)
														} else {
															v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
															v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
															F_errmsg(m, int32(719478), v11)
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2870), int32(315362))
																mBase = m.M
																v233 = m.ExcPending
																if v233 != 0 {
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
													v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														if l2 != 0 {
															v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
															if v87&int32(2) == int32(0) {
																v194 = int32(1)
																m.G0 = v11 + int32(16)
																return v194
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(317992), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
													}
												}
											}
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
											if v106 != 0 {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
												if v108&int32(1) == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(336384), int32(0))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(326828), int32(1264), int32(271151))
															mBase = m.M
															v246 = m.ExcPending
															if v246 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														if v116 != 0 {
															v194 = int32(1)
															m.G0 = v11 + int32(16)
															return v194
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2893), int32(315362))
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
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
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													if v116 != 0 {
														v194 = int32(1)
														m.G0 = v11 + int32(16)
														return v194
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(317992), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2893), int32(315362))
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
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
										}
									}
								}
							} else {
								v59 = v54 - int32(1)
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
								if v62 != 0 {
									v194 = int32(0)
									m.G0 = v11 + int32(16)
									return v194
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
									if v71 == int32(102) {
										v74 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
										v77 = F_GetFdwRoutineForRelation(m, v69, v74)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
											if v79 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return int32(0)
													} else {
														v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
														v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
														F_errmsg(m, int32(719478), v11)
														mBase = m.M
														v228 = m.ExcPending
														if v228 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496951), int32(2870), int32(315362))
															mBase = m.M
															v233 = m.ExcPending
															if v233 != 0 {
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
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													if l2 != 0 {
														v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
														if v87&int32(2) == int32(0) {
															v194 = int32(1)
															m.G0 = v11 + int32(16)
															return v194
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(317992), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(317992), int32(0))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
												}
											}
										}
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
										if v106 != 0 {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
											if v108&int32(1) == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(336384), int32(0))
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(326828), int32(1264), int32(271151))
														mBase = m.M
														v246 = m.ExcPending
														if v246 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													if v116 != 0 {
														v194 = int32(1)
														m.G0 = v11 + int32(16)
														return v194
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(317992), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496951), int32(2893), int32(315362))
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
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
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
											v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												if v116 != 0 {
													v194 = int32(1)
													m.G0 = v11 + int32(16)
													return v194
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(317992), int32(0))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496951), int32(2893), int32(315362))
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
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
									}
								}
							}
						} else {
							v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
							if v51 < v131 {
								F_slot_getsomeattrs_int(m, v50, v131)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return int32(0)
								} else {
									v136 = v131 - int32(1)
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
									if v139 != 0 {
										v194 = int32(0)
										m.G0 = v11 + int32(16)
										return v194
									} else {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
										v146 = m.G0
										v148 = v146 - int32(32)
										m.G0 = v148
										v151 = v148 + int32(20)
										*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
										v154 = F_pg_detoast_datum(m, v145)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int32(0)
										} else {
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
											v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
											*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
											*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
											v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
											m.T0[v166].(func(*base.Module, int32))(m, l2)
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int32(0)
											} else {
												v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
													v178 = v176 & int32(65533)
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
													v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
													m.G0 = v148 + int32(32)
													v194 = int32(1)
													m.G0 = v11 + int32(16)
													return v194
												}
											}
										}
									}
								}
							} else {
								v136 = v131 - int32(1)
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
								if v139 != 0 {
									v194 = int32(0)
									m.G0 = v11 + int32(16)
									return v194
								} else {
									v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
									v146 = m.G0
									v148 = v146 - int32(32)
									m.G0 = v148
									v151 = v148 + int32(20)
									*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
									v154 = F_pg_detoast_datum(m, v145)
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return int32(0)
									} else {
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
										*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
										v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
										*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
										*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
										v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
										m.T0[v166].(func(*base.Module, int32))(m, l2)
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int32(0)
										} else {
											v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int32(0)
											} else {
												v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
												v178 = v176 & int32(65533)
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
												v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
												m.G0 = v148 + int32(32)
												v194 = int32(1)
												m.G0 = v11 + int32(16)
												return v194
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
			v48 = v21
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
			if v48 == int32(4) {
				v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
				if v51 < v54 {
					F_slot_getsomeattrs_int(m, v50, v54)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v59 = v54 - int32(1)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
						if v62 != 0 {
							v194 = int32(0)
							m.G0 = v11 + int32(16)
							return v194
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
							if v71 == int32(102) {
								v74 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
								v77 = F_GetFdwRoutineForRelation(m, v69, v74)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
									if v79 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v217 = m.ExcPending
										if v217 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
												return int32(0)
											} else {
												v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
												F_errmsg(m, int32(719478), v11)
												mBase = m.M
												v228 = m.ExcPending
												if v228 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496951), int32(2870), int32(315362))
													mBase = m.M
													v233 = m.ExcPending
													if v233 != 0 {
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
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											if l2 != 0 {
												v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
												if v87&int32(2) == int32(0) {
													v194 = int32(1)
													m.G0 = v11 + int32(16)
													return v194
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(317992), int32(0))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(317992), int32(0))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
										}
									}
								}
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
								if v106 != 0 {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
									if v108&int32(1) == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(336384), int32(0))
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(326828), int32(1264), int32(271151))
												mBase = m.M
												v246 = m.ExcPending
												if v246 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
										v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											if v116 != 0 {
												v194 = int32(1)
												m.G0 = v11 + int32(16)
												return v194
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(317992), int32(0))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496951), int32(2893), int32(315362))
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
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
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
									v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										if v116 != 0 {
											v194 = int32(1)
											m.G0 = v11 + int32(16)
											return v194
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(317992), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496951), int32(2893), int32(315362))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
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
							}
						}
					}
				} else {
					v59 = v54 - int32(1)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v60))))
					if v62 != 0 {
						v194 = int32(0)
						m.G0 = v11 + int32(16)
						return v194
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
						if v71 == int32(102) {
							v74 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v74)
							v77 = F_GetFdwRoutineForRelation(m, v69, v74)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
								if v79 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
											return int32(0)
										} else {
											v221 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
											v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v222 + int32(4)
											F_errmsg(m, int32(719478), v11)
											mBase = m.M
											v228 = m.ExcPending
											if v228 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496951), int32(2870), int32(315362))
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
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
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									m.T0[v79].(func(*base.Module, int32, int32, int32, int32, int32))(m, v82, v20, v68, l2, v11+int32(15))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										if l2 != 0 {
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
											if v87&int32(2) == int32(0) {
												v194 = int32(1)
												m.G0 = v11 + int32(16)
												return v194
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(317992), int32(0))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(317992), int32(0))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496951), int32(2878), int32(315362))
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
									}
								}
							}
						} else {
							v106 = *(*int32)(unsafe.Add(mBase, _consts[114]))
							if v106 != 0 {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
								if v108&int32(1) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(336384), int32(0))
										mBase = m.M
										v241 = m.ExcPending
										if v241 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(326828), int32(1264), int32(271151))
											mBase = m.M
											v246 = m.ExcPending
											if v246 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
									v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										if v116 != 0 {
											v194 = int32(1)
											m.G0 = v11 + int32(16)
											return v194
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(317992), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496951), int32(2893), int32(315362))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
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
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)+188))
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
								v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(4173936), l2)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									if v116 != 0 {
										v194 = int32(1)
										m.G0 = v11 + int32(16)
										return v194
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(317992), int32(0))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496951), int32(2893), int32(315362))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
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
						}
					}
				}
			} else {
				v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
				if v51 < v131 {
					F_slot_getsomeattrs_int(m, v50, v131)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						v136 = v131 - int32(1)
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
						if v139 != 0 {
							v194 = int32(0)
							m.G0 = v11 + int32(16)
							return v194
						} else {
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
							v146 = m.G0
							v148 = v146 - int32(32)
							m.G0 = v148
							v151 = v148 + int32(20)
							*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
							v154 = F_pg_detoast_datum(m, v145)
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
								*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
								v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
								*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
								*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
								m.T0[v166].(func(*base.Module, int32))(m, l2)
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
									v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int32(0)
									} else {
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
										v178 = v176 & int32(65533)
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
										m.G0 = v148 + int32(32)
										v194 = int32(1)
										m.G0 = v11 + int32(16)
										return v194
									}
								}
							}
						}
					}
				} else {
					v136 = v131 - int32(1)
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
					v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v137))))
					if v139 != 0 {
						v194 = int32(0)
						m.G0 = v11 + int32(16)
						return v194
					} else {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v136<<(uint(int32(2))%32))))
						v146 = m.G0
						v148 = v146 - int32(32)
						m.G0 = v148
						v151 = v148 + int32(20)
						*(*int64)(unsafe.Add(mBase, uint32(v151))) = int64(0)
						v154 = F_pg_detoast_datum(m, v145)
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int32(0)
						} else {
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
							*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
							v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v160)
							v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v162
							*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v154
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
							m.T0[v166].(func(*base.Module, int32))(m, l2)
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return int32(0)
							} else {
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								F_heap_deform_tuple(m, v148+int32(12), v171, v172, v173)
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int32(0)
								} else {
									v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
									v178 = v176 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
									v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
									m.G0 = v148 + int32(32)
									v194 = int32(1)
									m.G0 = v11 + int32(16)
									return v194
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v204 = m.ExcPending
		if v204 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(153659), int32(0))
			mBase = m.M
			v208 = m.ExcPending
			if v208 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(496951), int32(2822), int32(315362))
				mBase = m.M
				v213 = m.ExcPending
				if v213 != 0 {
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
func F_ExecPrepareQual(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(4515600)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v7
	v9 = F_expression_planner(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = F_ExecInitQual(m, v9, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v5
			return v14
		}
	}
}
func F_cost_qual_eval_node(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(24)
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	v19 = F_cost_qual_eval_walker(m, l1, v8+int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
		m.G0 = v8 + int32(32)
		return
	}
}
func F_get_qual_for_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v37 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19))))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = F_makeVar(m, int32(1), v20, v23, v25, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = F_copyObjectImpl(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = v29
	goto L1
L8:
	;
	v36 = v34
	goto L1
L9:
	;
	m.G0 = v13 + int32(80)
	return v277
L10:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v250 != int32(1) {
		v277 = v249
		goto L9
	} else {
		goto L64
	}
L11:
	;
	v211 = F_palloc0(m, int32(20))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L58
	}
L12:
	;
	v161 = F_palloc0(m, int32(20))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L50
	}
L13:
	;
	if v47 == int32(-1) {
		v204 = v40
		goto L11
	} else {
		goto L49
	}
L14:
	;
	if v131 != 0 {
		goto L43
	} else {
		goto L44
	}
L15:
	;
	v40 = int32(0)
	v42 = F_RelationGetPartitionDesc(m, l0, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v95 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v44 == int32(0) {
		v277 = v40
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	v49 = base.B2i32(v47 != int32(-1))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v49|v50 == int32(0) {
		v277 = v40
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v50 <= int32(0) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(0)
	v61 = v40
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v61<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v83 = F_datumCopy(m, v80, v82, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v131 = v90
	v138 = v49
	goto L14
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v88 = F_makeConst(m, v68, v70, v72, v74, v83, int32(0), v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v90 = F_lappend(m, v57, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v93 = v61 + int32(1)
	if v93 != v50 {
		v57 = v90
		v61 = v93
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v204 = int32(0)
	goto L11
L29:
	;
	goto L30
L30:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v100 <= v99 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v131 = v99
	v138 = v3
	goto L14
L32:
	;
	goto L33
L33:
	;
	v104 = v99
	v108 = int32(0)
	v111 = v3
	goto L34
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v108<<(uint(int32(2))%32))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+24)))
	if v119 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v131 = v125
	v138 = v126
	goto L14
L36:
	;
	v128 = v108 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v128 < v129 {
		v104 = v125
		v108 = v128
		v111 = v126
		goto L34
	} else {
		goto L42
	}
L37:
	;
	v125 = v104
	v126 = int32(1)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v121 = F_copyObjectImpl(m, v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v123 = F_lappend(m, v104, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v125 = v123
	v126 = v111
	goto L36
L42:
	;
	goto L35
L43:
	;
	v143 = F_make_partition_op_expr(m, v15, int32(0), int32(3), v36, v131)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v147 = int32(0)
	if v138 != 0 {
		v154 = v147
		goto L12
	} else {
		goto L48
	}
L46:
	;
	if v138 == int32(0) {
		v204 = v143
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v154 = v143
	goto L12
L48:
	;
	v204 = v147
	goto L11
L49:
	;
	v154 = v40
	goto L12
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = int32(-1)
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+12)) = uint8(v165)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = int32(52)
	if v154 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v154
	v181 = F_list_make2_impl(m, v13+int32(28), v13+int32(24))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v161
	v198 = F_list_make1_impl(m, int32(1), v13+int32(16))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L57
	}
L54:
	;
	v184 = F_makeBoolExpr(m, int32(1), v181, int32(-1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v184
	v191 = F_list_make1_impl(m, int32(1), v13+int32(20))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v249 = v191
	goto L10
L57:
	;
	v249 = v198
	goto L10
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+16)) = int32(-1)
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+12)) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = int32(52)
	if v204 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v204
	v230 = F_list_make2_impl(m, v13+int32(40), v13+int32(36))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v211
	v237 = F_list_make1_impl(m, int32(1), v13+int32(32))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L63
	}
L62:
	;
	v249 = v230
	goto L10
L63:
	;
	v249 = v237
	goto L10
L64:
	;
	v253 = F_make_ands_explicit(m, v249)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v253
	v261 = F_list_make1_impl(m, int32(1), v13+int32(12))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v264 = F_makeBoolExpr(m, int32(2), v261, int32(-1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v264
	v271 = F_list_make1_impl(m, int32(1), v13+int32(8))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v277 = v271
	goto L9
}
