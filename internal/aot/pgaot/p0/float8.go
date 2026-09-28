package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_combine(m *base.Module, l0 int32) int64 {
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
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v81 float64
	_ = v81
	var v86 float64
	_ = v86
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			if v25 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
					F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int64(0)
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
					v152 = m.ExcPending
					if v152 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
						F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int64(0)
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
						v152 = m.ExcPending
						if v152 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
							F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int64(0)
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
							v152 = m.ExcPending
							if v152 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_float8_combine_0)
								F_errmsg_internal(m, int32(_a_F_float8_combine_1), v13+int32(-48))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int64(0)
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
								v170 = m.ExcPending
								if v170 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
									F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
											return int64(0)
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
									v170 = m.ExcPending
									if v170 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
										*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
										F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return int64(0)
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
										v170 = m.ExcPending
										if v170 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
											F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return int64(0)
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
											v170 = m.ExcPending
											if v170 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_combine_0)
												F_errmsg_internal(m, int32(_a_F_float8_combine_1), v15)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_float8_combine_2), int32(2985), int32(_a_F_float8_combine_3))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return int64(0)
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
												v94 = v45
												v96 = v46
												v97 = v47
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v102 == int32(0) {
													v130 = int32(0)
												} else {
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
													switch v105 - int32(435) {
													case 0:
														v130 = int32(1)
													case 1:
														v130 = int32(2)
													default:
														v130 = int32(0)
													}
												}
												if v130 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
													*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
													*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
													v143 = v18
													m.G0 = v15 - int32(-64)
													return base.I64_extend_i32_u(v143)
												} else {
													*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
													*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
													*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
													v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int64(0)
													} else {
														v143 = v141
														m.G0 = v15 - int32(-64)
														return base.I64_extend_i32_u(v143)
													}
												}
											} else {
												v51 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
												v52 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
												if base.F64_eq(v47, float64(0)) != 0 {
													v94 = v51
													v96 = v52
													v97 = v48
													v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v102 == int32(0) {
														v130 = int32(0)
													} else {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
														switch v105 - int32(435) {
														case 0:
															v130 = int32(1)
														case 1:
															v130 = int32(2)
														default:
															v130 = int32(0)
														}
													}
													if v130 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
														*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
														*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
														v143 = v18
														m.G0 = v15 - int32(-64)
														return base.I64_extend_i32_u(v143)
													} else {
														*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
														*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
														*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
														v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int64(0)
														} else {
															v143 = v141
															m.G0 = v15 - int32(-64)
															return base.I64_extend_i32_u(v143)
														}
													}
												} else {
													v55 = base.F64_add(v48, v47)
													v57 = math.Float64frombits(uint64(0x7ff0000000000000))
													v59 = base.F64_add(v52, v46)
													if base.F64_eq(base.F64_abs(v52), v57)|base.F64_ne(base.F64_abs(v59), v57)|base.F64_eq(base.F64_abs(v46), v57) == int32(0) {
														v71 = F_float_overflow_error_ext(m, int32(0))
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int64(0)
														} else {
															v74 = float64(0)
															v76 = math.Float64frombits(uint64(0x7ff0000000000000))
															v81 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
															v86 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v81, base.F64_mul(v81, base.F64_mul(v48, v47))), v55))
															if base.F64_eq(base.F64_abs(v51), v76)|base.F64_ne(base.F64_abs(v86), v76) != 0 {
																v94 = v86
																v96 = v74
																v97 = v55
																v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v102 == int32(0) {
																	v130 = int32(0)
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
																	switch v105 - int32(435) {
																	case 0:
																		v130 = int32(1)
																	case 1:
																		v130 = int32(2)
																	default:
																		v130 = int32(0)
																	}
																}
																if v130 != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
																	v143 = v18
																	m.G0 = v15 - int32(-64)
																	return base.I64_extend_i32_u(v143)
																} else {
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
																	v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int64(0)
																	} else {
																		v143 = v141
																		m.G0 = v15 - int32(-64)
																		return base.I64_extend_i32_u(v143)
																	}
																}
															} else {
																if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v94 = v86
																	v96 = v74
																	v97 = v55
																	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	if v102 == int32(0) {
																		v130 = int32(0)
																	} else {
																		v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
																		switch v105 - int32(435) {
																		case 0:
																			v130 = int32(1)
																		case 1:
																			v130 = int32(2)
																		default:
																			v130 = int32(0)
																		}
																	}
																	if v130 != 0 {
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
																		v143 = v18
																		m.G0 = v15 - int32(-64)
																		return base.I64_extend_i32_u(v143)
																	} else {
																		*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
																		*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
																		*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
																		v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int64(0)
																		} else {
																			v143 = v141
																			m.G0 = v15 - int32(-64)
																			return base.I64_extend_i32_u(v143)
																		}
																	}
																}
															}
														}
													} else {
														v74 = v59
														v76 = math.Float64frombits(uint64(0x7ff0000000000000))
														v81 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
														v86 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v81, base.F64_mul(v81, base.F64_mul(v48, v47))), v55))
														if base.F64_eq(base.F64_abs(v51), v76)|base.F64_ne(base.F64_abs(v86), v76) != 0 {
															v94 = v86
															v96 = v74
															v97 = v55
															v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															if v102 == int32(0) {
																v130 = int32(0)
															} else {
																v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
																switch v105 - int32(435) {
																case 0:
																	v130 = int32(1)
																case 1:
																	v130 = int32(2)
																default:
																	v130 = int32(0)
																}
															}
															if v130 != 0 {
																*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
																*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
																*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
																v143 = v18
																m.G0 = v15 - int32(-64)
																return base.I64_extend_i32_u(v143)
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
																*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
																*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
																v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int64(0)
																} else {
																	v143 = v141
																	m.G0 = v15 - int32(-64)
																	return base.I64_extend_i32_u(v143)
																}
															}
														} else {
															if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v184 = m.ExcPending
																if v184 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v94 = v86
																v96 = v74
																v97 = v55
																v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v102 == int32(0) {
																	v130 = int32(0)
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
																	switch v105 - int32(435) {
																	case 0:
																		v130 = int32(1)
																	case 1:
																		v130 = int32(2)
																	default:
																		v130 = int32(0)
																	}
																}
																if v130 != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v94
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v96
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v97
																	v143 = v18
																	m.G0 = v15 - int32(-64)
																	return base.I64_extend_i32_u(v143)
																} else {
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v94
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v96
																	*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v97
																	v141 = F_construct_array_builtin(m, v13+int32(-32), int32(3), int32(701))
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int64(0)
																	} else {
																		v143 = v141
																		m.G0 = v15 - int32(-64)
																		return base.I64_extend_i32_u(v143)
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
func F_float8_covar_samp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v36 int64
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_samp_0)
				F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v8)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2985), int32(_a_F_float8_covar_samp_3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_samp_0)
					F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v8)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2985), int32(_a_F_float8_covar_samp_3))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_samp_0)
						F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v8)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2985), int32(_a_F_float8_covar_samp_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_samp_0)
							F_errmsg_internal(m, int32(_a_F_float8_covar_samp_1), v8)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_covar_samp_2), int32(2985), int32(_a_F_float8_covar_samp_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(2)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v36 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
							v36 = base.I64_reinterpret_f64(base.F64_div(v31, base.F64_add(v25, float64(-1))))
						}
						m.G0 = v8 + int32(16)
						return v36
					}
				}
			}
		}
	}
}
func F_float8_regr_avgx(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v35 float64
	_ = v35
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgx_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v8)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2985), int32(_a_F_float8_regr_avgx_3))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgx_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2985), int32(_a_F_float8_regr_avgx_3))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgx_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v8)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2985), int32(_a_F_float8_regr_avgx_3))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgx_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_avgx_1), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_avgx_2), int32(2985), int32(_a_F_float8_regr_avgx_3))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v38 = int64(0)
						} else {
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v11)+72))
							if base.Ui64(v30&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v38 = v30
							} else {
								v35 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
								v38 = base.I64_reinterpret_f64(base.F64_div(v35, v25))
							}
						}
						m.G0 = v8 + int32(16)
						return v38
					}
				}
			}
		}
	}
}
