package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			if v25 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
					F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				if v28 != int32(3) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
						F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
					if v31 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
							F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
						if v32 != int32(701) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
								F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if v35 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
									F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
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
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								if v38 != int32(3) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
										*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
										F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
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
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
									if v41 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
											F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
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
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
										if v42 != int32(701) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
												F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_float8_combine_2), int32(2938), int32(_a_F_float8_combine_3))
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
											v45 = *(*float64)(unsafe.Add(mBase, uint32(v23)+40))
											v46 = *(*float64)(unsafe.Add(mBase, uint32(v23)+32))
											v47 = *(*float64)(unsafe.Add(mBase, uint32(v23)+24))
											v48 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
											if base.F64_eq(v48, float64(0)) != 0 {
												v130 = v45
												v131 = v47
												v132 = v46
												*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v132
												*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v131
												*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v130
												v138 = v130
												v140 = v131
												v141 = v132
												v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v146 == int32(0) {
													v174 = int32(0)
												} else {
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
													switch v149 - int32(429) {
													case 0:
														v174 = int32(1)
													case 1:
														v174 = int32(2)
													default:
														v174 = int32(0)
													}
												}
												if v174 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v141
													*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v140
													v193 = v18
													m.G0 = v15 - int32(-64)
													return v193
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
													v191 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int32(0)
													} else {
														v193 = v191
														m.G0 = v15 - int32(-64)
														return v193
													}
												}
											} else {
												v51 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
												v52 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
												if base.F64_eq(v47, float64(0)) != 0 {
													v130 = v51
													v131 = v48
													v132 = v52
													*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v132
													*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v131
													*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v130
													v138 = v130
													v140 = v131
													v141 = v132
													v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v146 == int32(0) {
														v174 = int32(0)
													} else {
														v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
														switch v149 - int32(429) {
														case 0:
															v174 = int32(1)
														case 1:
															v174 = int32(2)
														default:
															v174 = int32(0)
														}
													}
													if v174 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v141
														*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v140
														v193 = v18
														m.G0 = v15 - int32(-64)
														return v193
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
														v191 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return int32(0)
														} else {
															v193 = v191
															m.G0 = v15 - int32(-64)
															return v193
														}
													}
												} else {
													v55 = base.F64_add(v48, v47)
													*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v55
													v58 = math.Float64frombits(uint64(0x7ff0000000000000))
													v60 = base.F64_add(v52, v46)
													if base.B2i32(base.F64_eq(base.F64_abs(v52), v58)|base.F64_ne(base.F64_abs(v60), v58) == int32(0))&base.F64_ne(base.F64_abs(v46), v58) != 0 {
														F_float_overflow_error(m)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v60
														v75 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
														v80 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v75, base.F64_mul(v75, base.F64_mul(v48, v47))), v55))
														*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v80
														v83 = math.Float64frombits(uint64(0x7ff0000000000000))
														if base.F64_eq(base.F64_abs(v51), v83)|base.F64_ne(base.F64_abs(v80), v83) != 0 {
															v138 = v80
															v140 = v55
															v141 = v60
															v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															if v146 == int32(0) {
																v174 = int32(0)
															} else {
																v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
																switch v149 - int32(429) {
																case 0:
																	v174 = int32(1)
																case 1:
																	v174 = int32(2)
																default:
																	v174 = int32(0)
																}
															}
															if v174 != 0 {
																*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v138
																*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v141
																*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v140
																v193 = v18
																m.G0 = v15 - int32(-64)
																return v193
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																v191 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																mBase = m.M
																v192 = m.ExcPending
																if v192 != 0 {
																	return int32(0)
																} else {
																	v193 = v191
																	m.G0 = v15 - int32(-64)
																	return v193
																}
															}
														} else {
															if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v138 = v80
																v140 = v55
																v141 = v60
																v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v146 == int32(0) {
																	v174 = int32(0)
																} else {
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
																	switch v149 - int32(429) {
																	case 0:
																		v174 = int32(1)
																	case 1:
																		v174 = int32(2)
																	default:
																		v174 = int32(0)
																	}
																}
																if v174 != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v138
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v141
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v140
																	v193 = v18
																	m.G0 = v15 - int32(-64)
																	return v193
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																	v191 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																	mBase = m.M
																	v192 = m.ExcPending
																	if v192 != 0 {
																		return int32(0)
																	} else {
																		v193 = v191
																		m.G0 = v15 - int32(-64)
																		return v193
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
func F_float8_covar_samp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_covar_samp_0)
				F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v7)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2938), int32(_a_F_float8_covar_samp_3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_covar_samp_0)
					F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v7)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2938), int32(_a_F_float8_covar_samp_3))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_covar_samp_0)
						F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2938), int32(_a_F_float8_covar_samp_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_covar_samp_0)
							F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v7)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2938), int32(_a_F_float8_covar_samp_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(2)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v36 = int32(0)
							m.G0 = v7 + int32(16)
							return v36
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+64))
							v34 = F_Float8GetDatum(m, base.F64_div(v30, base.F64_add(v24, float64(-1))))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = v34
								m.G0 = v7 + int32(16)
								return v36
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_avgx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgx_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2938), int32(_a_F_float8_regr_avgx_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgx_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2938), int32(_a_F_float8_regr_avgx_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgx_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2938), int32(_a_F_float8_regr_avgx_3))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgx_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2938), int32(_a_F_float8_regr_avgx_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v34 = int32(0)
							m.G0 = v7 + int32(16)
							return v34
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+32))
							v32 = F_Float8GetDatum(m, base.F64_div(v30, v24))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = v32
								m.G0 = v7 + int32(16)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
