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
	var v57 float64
	_ = v57
	var v71 float64
	_ = v71
	var v76 float64
	_ = v76
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v169 int32
	_ = v169
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(373435)
					F_errmsg_internal(m, int32(26146), v13+int32(-48))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				if v28 != int32(3) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(373435)
						F_errmsg_internal(m, int32(26146), v13+int32(-48))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
					if v31 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(373435)
							F_errmsg_internal(m, int32(26146), v13+int32(-48))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
						if v32 != int32(701) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(373435)
								F_errmsg_internal(m, int32(26146), v13+int32(-48))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if v35 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(373435)
									F_errmsg_internal(m, int32(26146), v15)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								if v38 != int32(3) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
										*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(373435)
										F_errmsg_internal(m, int32(26146), v15)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
									if v41 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(373435)
											F_errmsg_internal(m, int32(26146), v15)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
										if v42 != int32(701) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(3)
												*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(373435)
												F_errmsg_internal(m, int32(26146), v15)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
										} else {
											v45 = *(*float64)(unsafe.Add(mBase, uint32(v23)+40))
											v46 = *(*float64)(unsafe.Add(mBase, uint32(v23)+32))
											v47 = *(*float64)(unsafe.Add(mBase, uint32(v23)+24))
											v48 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
											if base.F64_eq(v48, float64(0)) != 0 {
												v125 = v45
												v126 = v47
												v127 = v46
												*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v127
												*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v126
												*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v125
												v133 = v125
												v135 = v126
												v136 = v127
												v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v141 == int32(0) {
													v169 = int32(0)
												} else {
													v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
													switch v144 - int32(429) {
													case 0:
														v169 = int32(1)
													case 1:
														v169 = int32(2)
													default:
														v169 = int32(0)
													}
												}
												if v169 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
													*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
													*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
													v188 = v18
													m.G0 = v15 - int32(-64)
													return v188
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
													v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
													mBase = m.M
													v187 = m.ExcPending
													if v187 != 0 {
														return int32(0)
													} else {
														v188 = v186
														m.G0 = v15 - int32(-64)
														return v188
													}
												}
											} else {
												v51 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
												v52 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
												if base.F64_eq(v47, float64(0)) != 0 {
													v125 = v51
													v126 = v48
													v127 = v52
													*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v127
													*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v126
													*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v125
													v133 = v125
													v135 = v126
													v136 = v127
													v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v141 == int32(0) {
														v169 = int32(0)
													} else {
														v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
														switch v144 - int32(429) {
														case 0:
															v169 = int32(1)
														case 1:
															v169 = int32(2)
														default:
															v169 = int32(0)
														}
													}
													if v169 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
														*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
														*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
														v188 = v18
														m.G0 = v15 - int32(-64)
														return v188
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
														v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
														mBase = m.M
														v187 = m.ExcPending
														if v187 != 0 {
															return int32(0)
														} else {
															v188 = v186
															m.G0 = v15 - int32(-64)
															return v188
														}
													}
												} else {
													v55 = base.F64_add(v48, v47)
													*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = v55
													v57 = base.F64_add(v52, v46)
													if base.F64_ne(base.F64_abs(v57), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v57
														v71 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
														v76 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v71, base.F64_mul(v71, base.F64_mul(v48, v47))), v55))
														*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v76
														if base.F64_ne(base.F64_abs(v76), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
															v133 = v76
															v135 = v55
															v136 = v57
															v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
															if v141 == int32(0) {
																v169 = int32(0)
															} else {
																v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																switch v144 - int32(429) {
																case 0:
																	v169 = int32(1)
																case 1:
																	v169 = int32(2)
																default:
																	v169 = int32(0)
																}
															}
															if v169 != 0 {
																*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																v188 = v18
																m.G0 = v15 - int32(-64)
																return v188
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																mBase = m.M
																v187 = m.ExcPending
																if v187 != 0 {
																	return int32(0)
																} else {
																	v188 = v186
																	m.G0 = v15 - int32(-64)
																	return v188
																}
															}
														} else {
															if base.F64_eq(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																v133 = v76
																v135 = v55
																v136 = v57
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v141 == int32(0) {
																	v169 = int32(0)
																} else {
																	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																	switch v144 - int32(429) {
																	case 0:
																		v169 = int32(1)
																	case 1:
																		v169 = int32(2)
																	default:
																		v169 = int32(0)
																	}
																}
																if v169 != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																	v188 = v18
																	m.G0 = v15 - int32(-64)
																	return v188
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																	v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		v188 = v186
																		m.G0 = v15 - int32(-64)
																		return v188
																	}
																}
															} else {
																if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v133 = v76
																	v135 = v55
																	v136 = v57
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	if v141 == int32(0) {
																		v169 = int32(0)
																	} else {
																		v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																		switch v144 - int32(429) {
																		case 0:
																			v169 = int32(1)
																		case 1:
																			v169 = int32(2)
																		default:
																			v169 = int32(0)
																		}
																	}
																	if v169 != 0 {
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																		v188 = v18
																		m.G0 = v15 - int32(-64)
																		return v188
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																		v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																		mBase = m.M
																		v187 = m.ExcPending
																		if v187 != 0 {
																			return int32(0)
																		} else {
																			v188 = v186
																			m.G0 = v15 - int32(-64)
																			return v188
																		}
																	}
																}
															}
														}
													} else {
														if base.F64_eq(base.F64_abs(v52), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
															*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v57
															v71 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
															v76 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v71, base.F64_mul(v71, base.F64_mul(v48, v47))), v55))
															*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v76
															if base.F64_ne(base.F64_abs(v76), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																v133 = v76
																v135 = v55
																v136 = v57
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																if v141 == int32(0) {
																	v169 = int32(0)
																} else {
																	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																	switch v144 - int32(429) {
																	case 0:
																		v169 = int32(1)
																	case 1:
																		v169 = int32(2)
																	default:
																		v169 = int32(0)
																	}
																}
																if v169 != 0 {
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																	*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																	v188 = v18
																	m.G0 = v15 - int32(-64)
																	return v188
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																	v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		v188 = v186
																		m.G0 = v15 - int32(-64)
																		return v188
																	}
																}
															} else {
																if base.F64_eq(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																	v133 = v76
																	v135 = v55
																	v136 = v57
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	if v141 == int32(0) {
																		v169 = int32(0)
																	} else {
																		v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																		switch v144 - int32(429) {
																		case 0:
																			v169 = int32(1)
																		case 1:
																			v169 = int32(2)
																		default:
																			v169 = int32(0)
																		}
																	}
																	if v169 != 0 {
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																		v188 = v18
																		m.G0 = v15 - int32(-64)
																		return v188
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																		v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																		mBase = m.M
																		v187 = m.ExcPending
																		if v187 != 0 {
																			return int32(0)
																		} else {
																			v188 = v186
																			m.G0 = v15 - int32(-64)
																			return v188
																		}
																	}
																} else {
																	if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																		F_float_overflow_error(m)
																		mBase = m.M
																		v124 = m.ExcPending
																		if v124 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v133 = v76
																		v135 = v55
																		v136 = v57
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																		if v141 == int32(0) {
																			v169 = int32(0)
																		} else {
																			v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																			switch v144 - int32(429) {
																			case 0:
																				v169 = int32(1)
																			case 1:
																				v169 = int32(2)
																			default:
																				v169 = int32(0)
																			}
																		}
																		if v169 != 0 {
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																			v188 = v18
																			m.G0 = v15 - int32(-64)
																			return v188
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																			v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																			mBase = m.M
																			v187 = m.ExcPending
																			if v187 != 0 {
																				return int32(0)
																			} else {
																				v188 = v186
																				m.G0 = v15 - int32(-64)
																				return v188
																			}
																		}
																	}
																}
															}
														} else {
															if base.F64_ne(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v57
																v71 = base.F64_sub(base.F64_div(v52, v48), base.F64_div(v46, v47))
																v76 = base.F64_add(base.F64_add(v51, v45), base.F64_div(base.F64_mul(v71, base.F64_mul(v71, base.F64_mul(v48, v47))), v55))
																*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v76
																if base.F64_ne(base.F64_abs(v76), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																	v133 = v76
																	v135 = v55
																	v136 = v57
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																	if v141 == int32(0) {
																		v169 = int32(0)
																	} else {
																		v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																		switch v144 - int32(429) {
																		case 0:
																			v169 = int32(1)
																		case 1:
																			v169 = int32(2)
																		default:
																			v169 = int32(0)
																		}
																	}
																	if v169 != 0 {
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																		*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																		v188 = v18
																		m.G0 = v15 - int32(-64)
																		return v188
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																		v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																		mBase = m.M
																		v187 = m.ExcPending
																		if v187 != 0 {
																			return int32(0)
																		} else {
																			v188 = v186
																			m.G0 = v15 - int32(-64)
																			return v188
																		}
																	}
																} else {
																	if base.F64_eq(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																		v133 = v76
																		v135 = v55
																		v136 = v57
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																		if v141 == int32(0) {
																			v169 = int32(0)
																		} else {
																			v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																			switch v144 - int32(429) {
																			case 0:
																				v169 = int32(1)
																			case 1:
																				v169 = int32(2)
																			default:
																				v169 = int32(0)
																			}
																		}
																		if v169 != 0 {
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																			*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																			v188 = v18
																			m.G0 = v15 - int32(-64)
																			return v188
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																			v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																			mBase = m.M
																			v187 = m.ExcPending
																			if v187 != 0 {
																				return int32(0)
																			} else {
																				v188 = v186
																				m.G0 = v15 - int32(-64)
																				return v188
																			}
																		}
																	} else {
																		if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v124 = m.ExcPending
																			if v124 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v133 = v76
																			v135 = v55
																			v136 = v57
																			v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																			if v141 == int32(0) {
																				v169 = int32(0)
																			} else {
																				v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
																				switch v144 - int32(429) {
																				case 0:
																					v169 = int32(1)
																				case 1:
																					v169 = int32(2)
																				default:
																					v169 = int32(0)
																				}
																			}
																			if v169 != 0 {
																				*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = v133
																				*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v136
																				*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v135
																				v188 = v18
																				m.G0 = v15 - int32(-64)
																				return v188
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v13 + int32(-16)
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v13 + int32(-8)
																				v186 = F_construct_array_builtin(m, v13+int32(-36), int32(3), int32(701))
																				mBase = m.M
																				v187 = m.ExcPending
																				if v187 != 0 {
																					return int32(0)
																				} else {
																					v188 = v186
																					m.G0 = v15 - int32(-64)
																					return v188
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
	var v32 float64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
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
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237133)
				F_errmsg_internal(m, int32(26146), v7)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
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
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237133)
					F_errmsg_internal(m, int32(26146), v7)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
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
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237133)
						F_errmsg_internal(m, int32(26146), v7)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
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
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237133)
							F_errmsg_internal(m, int32(26146), v7)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
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
							v38 = int32(0)
							m.G0 = v7 + int32(16)
							return v38
						} else {
							v32 = *(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64))))
							v36 = F_Float8GetDatum(m, base.F64_div(v32, base.F64_add(v24, float64(-1))))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = v36
								m.G0 = v7 + int32(16)
								return v38
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
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(27532)
				F_errmsg_internal(m, int32(26146), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(27532)
					F_errmsg_internal(m, int32(26146), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(27532)
						F_errmsg_internal(m, int32(26146), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(27532)
							F_errmsg_internal(m, int32(26146), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
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
