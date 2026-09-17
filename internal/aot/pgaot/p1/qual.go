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
	var v49 int32
	_ = v49
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
							v49 = v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
							if v49 == int32(4) {
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
														v215 = m.ExcPending
														if v215 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v218 = m.ExcPending
															if v218 != 0 {
																return int32(0)
															} else {
																v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
																v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
																F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
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
																		F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
																	F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
												v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
												if v106 != 0 {
													v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
													if v108&int32(1) == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v235 = m.ExcPending
														if v235 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
															mBase = m.M
															v239 = m.ExcPending
															if v239 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
																mBase = m.M
																v244 = m.ExcPending
																if v244 != 0 {
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
														v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
																	F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
													v215 = m.ExcPending
													if v215 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v218 = m.ExcPending
														if v218 != 0 {
															return int32(0)
														} else {
															v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
															v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
															F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
																	F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
											v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
											if v106 != 0 {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
												if v108&int32(1) == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v235 = m.ExcPending
													if v235 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
															mBase = m.M
															v244 = m.ExcPending
															if v244 != 0 {
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
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
															F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
											*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
											v152 = F_pg_detoast_datum(m, v145)
											mBase = m.M
											v153 = m.ExcPending
											if v153 != 0 {
												return int32(0)
											} else {
												v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
												*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
												v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
												*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
												*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
												v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
												m.T0[v164].(func(*base.Module, int32))(m, l2)
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
													F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
														v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
														*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
														v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
														*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
										*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
										v152 = F_pg_detoast_datum(m, v145)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
											v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
											*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
											*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
											m.T0[v164].(func(*base.Module, int32))(m, l2)
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
													v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
													v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
						v49 = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
						if v49 == int32(4) {
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
													v215 = m.ExcPending
													if v215 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v218 = m.ExcPending
														if v218 != 0 {
															return int32(0)
														} else {
															v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
															v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
															F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
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
																	F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
											v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
											if v106 != 0 {
												v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
												if v108&int32(1) == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v235 = m.ExcPending
													if v235 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
															mBase = m.M
															v244 = m.ExcPending
															if v244 != 0 {
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
													v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
															F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
												v215 = m.ExcPending
												if v215 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return int32(0)
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
														v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
														F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
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
																F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
															F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
										v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
										if v106 != 0 {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
											if v108&int32(1) == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v235 = m.ExcPending
												if v235 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
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
												v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
															F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
											v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
														F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
										*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
										v152 = F_pg_detoast_datum(m, v145)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
											v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
											*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
											*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
											m.T0[v164].(func(*base.Module, int32))(m, l2)
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
													v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
													v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
													*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
									*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
									v152 = F_pg_detoast_datum(m, v145)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
										*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
										v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
										*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
										v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
										m.T0[v164].(func(*base.Module, int32))(m, l2)
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return int32(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
												v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
												v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
												*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
			v49 = v21
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+6)))
			if v49 == int32(4) {
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
										v215 = m.ExcPending
										if v215 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v218 = m.ExcPending
											if v218 != 0 {
												return int32(0)
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
												v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
												F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
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
														F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
													F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
								v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
								if v106 != 0 {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
									if v108&int32(1) == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v235 = m.ExcPending
										if v235 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
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
										v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
													F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
									v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
												F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
									v215 = m.ExcPending
									if v215 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return int32(0)
										} else {
											v219 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
											v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v220 + int32(4)
											F_errmsg(m, int32(_a_F_EvalPlanQualFetchRowMark_0), v11)
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2870), int32(_a_F_EvalPlanQualFetchRowMark_2))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
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
													F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
												F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2878), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
							v106 = *(*int32)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[0]))
							if v106 != 0 {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EvalPlanQualFetchRowMark[1])))
								if v108&int32(1) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_4), int32(0))
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_5), int32(1264), int32(_a_F_EvalPlanQualFetchRowMark_6))
											mBase = m.M
											v244 = m.ExcPending
											if v244 != 0 {
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
									v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
												F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
								v116 = m.T0[v115].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v68, int32(_a_F_EvalPlanQualFetchRowMark_7), l2)
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
											F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_3), int32(0))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2893), int32(_a_F_EvalPlanQualFetchRowMark_2))
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
							*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
							v152 = F_pg_detoast_datum(m, v145)
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int32(0)
							} else {
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
								*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
								v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
								*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
								*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
								v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
								m.T0[v164].(func(*base.Module, int32))(m, l2)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return int32(0)
									} else {
										v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
										v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
										*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
						*(*int64)(unsafe.Add(mBase, uint32(v148)+20)) = int64(0)
						v152 = F_pg_detoast_datum(m, v145)
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
							*(*int32)(unsafe.Add(mBase, uint32(v148)+12)) = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
							v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = v158
							v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)) = uint16(v160)
							*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v152
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
							m.T0[v164].(func(*base.Module, int32))(m, l2)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								F_heap_deform_tuple(m, v148+int32(12), v169, v170, v171)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
									v176 = v174 & int32(_a_F_EvalPlanQualFetchRowMark_8)
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v176)
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
									*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v179)
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
		v202 = m.ExcPending
		if v202 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_EvalPlanQualFetchRowMark_9), int32(0))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_EvalPlanQualFetchRowMark_1), int32(2822), int32(_a_F_EvalPlanQualFetchRowMark_2))
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
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
	v4 = int32(_a_F_ExecPrepareQual_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareQual[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareQual[0])) = v7
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
			*(*int32)(unsafe.Add(mBase, _c_F_ExecPrepareQual[0])) = v5
			return v14
		}
	}
}
func F_cost_qual_eval_node(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v9
	v16 = F_cost_qual_eval_walker(m, l1, v7+int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v20
		m.G0 = v7 + int32(32)
		return
	}
}
func F_get_qual_for_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v36 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18))))
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = F_makeVar(m, int32(1), v19, v22, v24, v26, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = F_copyObjectImpl(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v35 = v28
	goto L1
L8:
	;
	v35 = v33
	goto L1
L9:
	;
	m.G0 = v12 + int32(80)
	return v269
L10:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v243 != int32(1) {
		v269 = v242
		goto L9
	} else {
		goto L62
	}
L11:
	;
	v205 = F_palloc0(m, int32(20))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L56
	}
L12:
	;
	v156 = F_palloc0(m, int32(20))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L48
	}
L13:
	;
	if v46 == int32(-1) {
		v198 = v39
		goto L11
	} else {
		goto L47
	}
L14:
	;
	if v128 != 0 {
		goto L41
	} else {
		goto L42
	}
L15:
	;
	v39 = int32(0)
	v41 = F_RelationGetPartitionDesc(m, l0, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v93 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v43 == int32(0) {
		v269 = v39
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	v48 = base.B2i32(v46 != int32(-1))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v48|v49 == int32(0) {
		v269 = v39
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v49 <= int32(0) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v56 = int32(0)
	v59 = v39
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v59<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v81 = F_datumCopy(m, v78, v80, v72)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v128 = v88
	v136 = v48
	goto L14
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v86 = F_makeConst(m, v66, v68, v70, v72, v81, int32(0), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v88 = F_lappend(m, v56, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v91 = v59 + int32(1)
	if v91 != v49 {
		v56 = v88
		v59 = v91
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v198 = int32(0)
	goto L11
L29:
	;
	goto L30
L30:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v98 <= v97 {
		v128 = v97
		v136 = v3
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v102 = v97
	v105 = int32(0)
	v110 = v3
	goto L32
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v105<<(uint(int32(2))%32))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+24)))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v128 = v122
	v136 = v123
	goto L14
L34:
	;
	v125 = v105 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v125 < v126 {
		v102 = v122
		v105 = v125
		v110 = v123
		goto L32
	} else {
		goto L40
	}
L35:
	;
	v122 = v102
	v123 = int32(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v118 = F_copyObjectImpl(m, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v120 = F_lappend(m, v102, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v122 = v120
	v123 = v110
	goto L34
L40:
	;
	goto L33
L41:
	;
	v139 = F_make_partition_op_expr(m, v14, int32(0), int32(3), v35, v128)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v143 = int32(0)
	if v136 != 0 {
		v149 = v143
		goto L12
	} else {
		goto L46
	}
L44:
	;
	if v136 == int32(0) {
		v198 = v139
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v149 = v139
	goto L12
L46:
	;
	v198 = v143
	goto L11
L47:
	;
	v149 = v39
	goto L12
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = int32(-1)
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+12)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v156)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = int32(52)
	if v149 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v149
	v176 = F_list_make2_impl(m, v12+int32(28), v12+int32(24))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v156
	v193 = F_list_make1_impl(m, int32(1), v12+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L55
	}
L52:
	;
	v179 = F_makeBoolExpr(m, int32(1), v176, int32(-1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v179
	v186 = F_list_make1_impl(m, int32(1), v12+int32(20))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v242 = v186
	goto L10
L55:
	;
	v242 = v193
	goto L10
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+16)) = int32(-1)
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+12)) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v205)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = int32(52)
	if v198 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v198
	v224 = F_list_make2_impl(m, v12+int32(40), v12+int32(36))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v205
	v231 = F_list_make1_impl(m, int32(1), v12+int32(32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L61
	}
L60:
	;
	v242 = v224
	goto L10
L61:
	;
	v242 = v231
	goto L10
L62:
	;
	v246 = F_make_ands_explicit(m, v242)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v246
	v254 = F_list_make1_impl(m, int32(1), v12+int32(12))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v257 = F_makeBoolExpr(m, int32(2), v254, int32(-1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v257
	v264 = F_list_make1_impl(m, int32(1), v12+int32(8))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v269 = v264
	goto L9
}
