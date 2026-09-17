package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_avg(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_avg_0)
				F_errmsg_internal(m, int32(_a_F_float8_avg_1), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_avg_2), int32(2938), int32(_a_F_float8_avg_3))
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
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_avg_0)
					F_errmsg_internal(m, int32(_a_F_float8_avg_1), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_avg_2), int32(2938), int32(_a_F_float8_avg_3))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_avg_0)
						F_errmsg_internal(m, int32(_a_F_float8_avg_1), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_avg_2), int32(2938), int32(_a_F_float8_avg_3))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_avg_0)
							F_errmsg_internal(m, int32(_a_F_float8_avg_1), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_avg_2), int32(2938), int32(_a_F_float8_avg_3))
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
						if base.F64_eq(v24, float64(0)) != 0 {
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
func F_float8_cmp_internal(m *base.Module, l0 float64, l1 float64) int32 {
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	v7 = int64(9223372036854775807)
	v8 = base.I64_reinterpret_f64(l0) & v7
	v11 = base.I64_reinterpret_f64(l1) & v7
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v11) {
		v21 = base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313)))
		v30 = int32(0) - v21&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
	} else {
		v16 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v8))|base.F64_gt(l0, l1) != 0 {
			v30 = v16
		} else {
			v21 = v16
			v30 = int32(0) - v21&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
		}
	}
	return v30
}
func F_float8_regr_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v97 float64
	_ = v97
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v127 float64
	_ = v127
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v208 float64
	_ = v208
	var v219 float64
	_ = v219
	var v220 float64
	_ = v220
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v261 int32
	_ = v261
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v31 = F_pg_detoast_datum(m, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			if v33 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_float8_regr_combine_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23+int32(16))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				if v36 != int32(6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_float8_regr_combine_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23+int32(16))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					if v39 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_float8_regr_combine_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23+int32(16))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
						if v40 != int32(701) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(6)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_float8_regr_combine_0)
								F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23+int32(16))
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v43 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(6)
									*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(_a_F_float8_regr_combine_0)
									F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
								if v46 != int32(6) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(6)
										*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(_a_F_float8_regr_combine_0)
										F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
									if v49 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(6)
											*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(_a_F_float8_regr_combine_0)
											F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23)
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
										if v50 != int32(701) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(6)
												*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(_a_F_float8_regr_combine_0)
												F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v23)
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2938), int32(_a_F_float8_regr_combine_3))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v53 = *(*float64)(unsafe.Add(mBase, uint32(v31)+64))
											v54 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
											v55 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
											v56 = *(*float64)(unsafe.Add(mBase, uint32(v31)+40))
											v57 = *(*float64)(unsafe.Add(mBase, uint32(v31)+32))
											v58 = *(*float64)(unsafe.Add(mBase, uint32(v31)+24))
											v59 = *(*float64)(unsafe.Add(mBase, uint32(v26)+24))
											if base.F64_eq(v59, float64(0)) != 0 {
												v203 = v53
												v204 = v56
												v205 = v54
												v206 = v55
												v207 = v58
												v208 = v57
												*(*float64)(unsafe.Add(mBase, uint32(v23)+96)) = v208
												*(*float64)(unsafe.Add(mBase, uint32(v23)+104)) = v207
												*(*float64)(unsafe.Add(mBase, uint32(v23)+88)) = v204
												*(*float64)(unsafe.Add(mBase, uint32(v23)+80)) = v206
												*(*float64)(unsafe.Add(mBase, uint32(v23)+72)) = v205
												*(*float64)(unsafe.Add(mBase, uint32(v23)+64)) = v203
												v219 = v203
												v220 = v204
												v221 = v205
												v222 = v206
												v223 = v207
												v224 = v208
												v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v233 == int32(0) {
													v261 = int32(0)
												} else {
													v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
													switch v236 - int32(429) {
													case 0:
														v261 = int32(1)
													case 1:
														v261 = int32(2)
													default:
														v261 = int32(0)
													}
												}
												if v261 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v219
													*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v221
													*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v222
													*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v220
													*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v224
													*(*float64)(unsafe.Add(mBase, uint32(v26)+24)) = v223
													v292 = v26
													m.G0 = v23 + int32(112)
													return v292
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v23 - int32(-64)
													*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v23 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v23 + int32(80)
													*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v23 + int32(88)
													*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v23 + int32(96)
													*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(104)
													v290 = F_construct_array_builtin(m, v23+int32(32), int32(6), int32(701))
													mBase = m.M
													v291 = m.ExcPending
													if v291 != 0 {
														return int32(0)
													} else {
														v292 = v290
														m.G0 = v23 + int32(112)
														return v292
													}
												}
											} else {
												v62 = *(*float64)(unsafe.Add(mBase, uint32(v26)+64))
												v63 = *(*float64)(unsafe.Add(mBase, uint32(v26)+56))
												v64 = *(*float64)(unsafe.Add(mBase, uint32(v26)+48))
												v65 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
												v66 = *(*float64)(unsafe.Add(mBase, uint32(v26)+32))
												if base.F64_eq(v58, float64(0)) != 0 {
													v203 = v62
													v204 = v65
													v205 = v63
													v206 = v64
													v207 = v59
													v208 = v66
													*(*float64)(unsafe.Add(mBase, uint32(v23)+96)) = v208
													*(*float64)(unsafe.Add(mBase, uint32(v23)+104)) = v207
													*(*float64)(unsafe.Add(mBase, uint32(v23)+88)) = v204
													*(*float64)(unsafe.Add(mBase, uint32(v23)+80)) = v206
													*(*float64)(unsafe.Add(mBase, uint32(v23)+72)) = v205
													*(*float64)(unsafe.Add(mBase, uint32(v23)+64)) = v203
													v219 = v203
													v220 = v204
													v221 = v205
													v222 = v206
													v223 = v207
													v224 = v208
													v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v233 == int32(0) {
														v261 = int32(0)
													} else {
														v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
														switch v236 - int32(429) {
														case 0:
															v261 = int32(1)
														case 1:
															v261 = int32(2)
														default:
															v261 = int32(0)
														}
													}
													if v261 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v219
														*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v221
														*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v222
														*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v220
														*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v224
														*(*float64)(unsafe.Add(mBase, uint32(v26)+24)) = v223
														v292 = v26
														m.G0 = v23 + int32(112)
														return v292
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v23 - int32(-64)
														*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v23 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v23 + int32(80)
														*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v23 + int32(88)
														*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v23 + int32(96)
														*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(104)
														v290 = F_construct_array_builtin(m, v23+int32(32), int32(6), int32(701))
														mBase = m.M
														v291 = m.ExcPending
														if v291 != 0 {
															return int32(0)
														} else {
															v292 = v290
															m.G0 = v23 + int32(112)
															return v292
														}
													}
												} else {
													v69 = base.F64_add(v59, v58)
													*(*float64)(unsafe.Add(mBase, uint32(v23)+104)) = v69
													v72 = math.Float64frombits(uint64(0x7ff0000000000000))
													v74 = base.F64_add(v66, v57)
													if base.B2i32(base.F64_eq(base.F64_abs(v66), v72)|base.F64_ne(base.F64_abs(v74), v72) == int32(0))&base.F64_ne(base.F64_abs(v57), v72) != 0 {
														F_float_overflow_error(m)
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*float64)(unsafe.Add(mBase, uint32(v23)+96)) = v74
														v89 = base.F64_sub(base.F64_div(v66, v59), base.F64_div(v57, v58))
														v90 = base.F64_mul(v59, v58)
														v91 = base.F64_mul(v89, v90)
														v94 = base.F64_add(base.F64_add(v65, v56), base.F64_div(base.F64_mul(v89, v91), v69))
														*(*float64)(unsafe.Add(mBase, uint32(v23)+88)) = v94
														v97 = math.Float64frombits(uint64(0x7ff0000000000000))
														if base.B2i32(base.F64_eq(base.F64_abs(v65), v97)|base.F64_ne(base.F64_abs(v94), v97) == int32(0))&base.F64_ne(base.F64_abs(v56), v97) != 0 {
															F_float_overflow_error(m)
															mBase = m.M
															v202 = m.ExcPending
															if v202 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v110 = math.Float64frombits(uint64(0x7ff0000000000000))
															v112 = base.F64_add(v64, v55)
															if base.B2i32(base.F64_eq(base.F64_abs(v64), v110)|base.F64_ne(base.F64_abs(v112), v110) == int32(0))&base.F64_ne(base.F64_abs(v55), v110) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v202 = m.ExcPending
																if v202 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(v23)+80)) = v112
																v127 = base.F64_sub(base.F64_div(v64, v59), base.F64_div(v55, v58))
																v131 = base.F64_add(base.F64_add(v63, v54), base.F64_div(base.F64_mul(v127, base.F64_mul(v90, v127)), v69))
																*(*float64)(unsafe.Add(mBase, uint32(v23)+72)) = v131
																v134 = math.Float64frombits(uint64(0x7ff0000000000000))
																if base.B2i32(base.F64_eq(base.F64_abs(v63), v134)|base.F64_ne(base.F64_abs(v131), v134) == int32(0))&base.F64_ne(base.F64_abs(v54), v134) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v149 = base.F64_add(base.F64_add(v62, v53), base.F64_div(base.F64_mul(v91, v127), v69))
																	*(*float64)(unsafe.Add(mBase, uint32(v23)+64)) = v149
																	v152 = math.Float64frombits(uint64(0x7ff0000000000000))
																	if base.F64_eq(base.F64_abs(v62), v152)|base.F64_ne(base.F64_abs(v149), v152) != 0 {
																		v219 = v149
																		v220 = v94
																		v221 = v131
																		v222 = v112
																		v223 = v69
																		v224 = v74
																		v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																		if v233 == int32(0) {
																			v261 = int32(0)
																		} else {
																			v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
																			switch v236 - int32(429) {
																			case 0:
																				v261 = int32(1)
																			case 1:
																				v261 = int32(2)
																			default:
																				v261 = int32(0)
																			}
																		}
																		if v261 != 0 {
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v219
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v221
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v222
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v220
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v224
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+24)) = v223
																			v292 = v26
																			m.G0 = v23 + int32(112)
																			return v292
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v23 - int32(-64)
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v23 + int32(72)
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v23 + int32(80)
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v23 + int32(88)
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v23 + int32(96)
																			*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(104)
																			v290 = F_construct_array_builtin(m, v23+int32(32), int32(6), int32(701))
																			mBase = m.M
																			v291 = m.ExcPending
																			if v291 != 0 {
																				return int32(0)
																			} else {
																				v292 = v290
																				m.G0 = v23 + int32(112)
																				return v292
																			}
																		}
																	} else {
																		if base.F64_ne(base.F64_abs(v53), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v202 = m.ExcPending
																			if v202 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v219 = v149
																			v220 = v94
																			v221 = v131
																			v222 = v112
																			v223 = v69
																			v224 = v74
																			v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																			if v233 == int32(0) {
																				v261 = int32(0)
																			} else {
																				v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
																				switch v236 - int32(429) {
																				case 0:
																					v261 = int32(1)
																				case 1:
																					v261 = int32(2)
																				default:
																					v261 = int32(0)
																				}
																			}
																			if v261 != 0 {
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v219
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v221
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v222
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v220
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v224
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+24)) = v223
																				v292 = v26
																				m.G0 = v23 + int32(112)
																				return v292
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v23 - int32(-64)
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v23 + int32(72)
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v23 + int32(80)
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v23 + int32(88)
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v23 + int32(96)
																				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v23 + int32(104)
																				v290 = F_construct_array_builtin(m, v23+int32(32), int32(6), int32(701))
																				mBase = m.M
																				v291 = m.ExcPending
																				if v291 != 0 {
																					return int32(0)
																				} else {
																					v292 = v290
																					m.G0 = v23 + int32(112)
																					return v292
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
func F_float8_regr_sxx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_sxx_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2938), int32(_a_F_float8_regr_sxx_3))
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
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			if v16 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_sxx_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2938), int32(_a_F_float8_regr_sxx_3))
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
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_sxx_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2938), int32(_a_F_float8_regr_sxx_3))
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
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v20 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_sxx_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2938), int32(_a_F_float8_regr_sxx_3))
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
					} else {
						v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						if base.F64_lt(v23, float64(1)) != 0 {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v32 = int32(0)
							m.G0 = v6 + int32(16)
							return v32
						} else {
							v29 = *(*float64)(unsafe.Add(mBase, uint32(v9)+40))
							v30 = F_Float8GetDatum(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v30
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_syy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_syy_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2938), int32(_a_F_float8_regr_syy_3))
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
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			if v16 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_syy_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2938), int32(_a_F_float8_regr_syy_3))
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
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_syy_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2938), int32(_a_F_float8_regr_syy_3))
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
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v20 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_float8_regr_syy_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2938), int32(_a_F_float8_regr_syy_3))
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
					} else {
						v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						if base.F64_lt(v23, float64(1)) != 0 {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v32 = int32(0)
							m.G0 = v6 + int32(16)
							return v32
						} else {
							v29 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
							v30 = F_Float8GetDatum(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v30
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_stddev_pop(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_pop_0)
				F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2938), int32(_a_F_float8_stddev_pop_3))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_pop_0)
					F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v7)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2938), int32(_a_F_float8_stddev_pop_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_pop_0)
						F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2938), int32(_a_F_float8_stddev_pop_3))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_pop_0)
							F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v7)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2938), int32(_a_F_float8_stddev_pop_3))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
						if base.F64_eq(v24, float64(0)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v35 = int32(0)
							m.G0 = v7 + int32(16)
							return v35
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v33 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_div(v30, v24)))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = v33
								m.G0 = v7 + int32(16)
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_stddev_samp(m *base.Module, l0 int32) int32 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_samp_0)
				F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v7)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2938), int32(_a_F_float8_stddev_samp_3))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
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
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_samp_0)
					F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2938), int32(_a_F_float8_stddev_samp_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
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
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_samp_0)
						F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v7)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2938), int32(_a_F_float8_stddev_samp_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_stddev_samp_0)
							F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v7)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2938), int32(_a_F_float8_stddev_samp_3))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
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
						if base.F64_le(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v37 = int32(0)
							m.G0 = v7 + int32(16)
							return v37
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v35 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_div(v30, base.F64_add(v24, float64(-1)))))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = v35
								m.G0 = v7 + int32(16)
								return v37
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 float64
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 float64
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		if v23 == int32(1) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
			if v29 == int32(18) {
				v32 = int32(16)
			} else {
				v32 = int32(0)
			}
			if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v39 = int32(4)
			} else {
				v39 = v32
			}
			v52 = v39
		} else {
			v40 = int32(1)
			if v23&v40 != 0 {
				v52 = int32(base.Ui32(v23)>>(uint(v40)%32)) - v40
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v52-int32(1)) {
			v58 = F_cstring_to_text(m, int32(_a_F_float8_to_char_0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v244 = v58
				m.G0 = v14 + int32(96)
				return v244
			}
		} else {
			v64 = F_palloc0(m, v52<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = F_NUM_cache(m, v52, v14+int32(60), v19, v14+int32(59))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v72&int32(1024) != 0 {
						v76 = base.F64_nearest(v17)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v76)&int64(9223372036854775807)) {
							v83 = int32(2147483647)
						} else {
							v83 = base.I32_trunc_sat_f64_s(v76)
						}
						if base.F64_ge(v76, float64(-2.147483648e+09)) != 0 {
							v87 = v83
						} else {
							v87 = int32(2147483647)
						}
						if base.F64_lt(v76, float64(2.147483648e+09)) != 0 {
							v91 = v87
						} else {
							v91 = int32(2147483647)
						}
						v92 = F_int_to_roman(m, v91)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							v213 = v92
							v216 = v2
							v218 = v2
							v224 = v64 + int32(4)
							F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int32(0)
							} else {
								v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v229 == int32(1) {
									F_pfree(m, v70)
									mBase = m.M
									v233 = m.ExcPending
									if v233 != 0 {
										return int32(0)
									} else {
										v234 = F_strlen(m, v224)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
										v244 = v64
										m.G0 = v14 + int32(96)
										return v244
									}
								} else {
									v234 = F_strlen(m, v224)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
									v244 = v64
									m.G0 = v14 + int32(96)
									return v244
								}
							}
						}
					} else {
						if v72&int32(_a_F_float8_to_char_1) != 0 {
							v96 = base.F64_abs(v17)
							if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v96)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v96, math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v107 = v105 + v106
								v110 = F_palloc(m, v107+int32(7))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v113 = v107 + int32(6)
									if v113 != 0 {
										base.MemoryFill(m, v110, int32(35), v113)
									} else {
									}
									v116 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v110+v113))) = uint8(v116)
									v120 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v120)
									v123 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v110+v105)+1)) = uint8(v123)
									v213 = v110
									v216 = v116
									v218 = v2
									v224 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v229 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int32(0)
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return v244
											}
										} else {
											v234 = F_strlen(m, v224)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
											v244 = v64
											m.G0 = v14 + int32(96)
											return v244
										}
									}
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v125
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = v17
								v131 = F_psprintf(m, int32(_a_F_float8_to_char_2), v14+int32(32))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
									if v133 != int32(43) {
										v213 = v131
										v216 = v2
										v218 = v2
									} else {
										v136 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v136)
										v213 = v131
										v216 = v2
										v218 = v2
									}
									v224 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v229 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int32(0)
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return v244
											}
										} else {
											v234 = F_strlen(m, v224)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
											v244 = v64
											m.G0 = v14 + int32(96)
											return v244
										}
									}
								}
							}
						} else {
							if v72&int32(2048) != 0 {
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v140 + v141
								v146 = F_pow(m, float64(10), base.F64_convert_i32_s(v140))
								mBase = m.M
								v149 = base.F64_mul(v17, v146)
							} else {
								v149 = v17
							}
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v149)
							v155 = F_psprintf(m, int32(_a_F_float8_to_char_3), v14+int32(16))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								v157 = F_strlen(m, v155)
								mBase = m.M
								if v157 <= int32(14) {
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if v160+v157 < int32(16) {
										v168 = v160
									} else {
										v166 = int32(15) - v157
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v166
										v168 = v166
									}
								} else {
									v166 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v166
									v168 = v166
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v149
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
								v172 = F_psprintf(m, int32(_a_F_float8_to_char_4), v14)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
									v176 = base.B2i32(v174 == int32(45))
									v177 = v172 + v176
									v178 = int32(46)
									v179 = F___strchrnul(m, v177, v178)
									mBase = m.M
									v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
									if v181 == v178 {
										v185 = v179
									} else {
										v185 = int32(0)
									}
									if v185 != 0 {
										v188 = v185 - v177
									} else {
										v187 = F_strlen(m, v177)
										mBase = m.M
										v188 = v187
									}
									if v174 == int32(45) {
										v191 = int32(45)
									} else {
										v191 = int32(43)
									}
									v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if v188 < v192 {
										v213 = v177
										v216 = v192 - v188
										v218 = v191
										v224 = v64 + int32(4)
										F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return int32(0)
										} else {
											v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v229 == int32(1) {
												F_pfree(m, v70)
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int32(0)
												} else {
													v234 = F_strlen(m, v224)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
													v244 = v64
													m.G0 = v14 + int32(96)
													return v244
												}
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return v244
											}
										}
									} else {
										if v188 <= v192 {
											v213 = v177
											v216 = int32(0)
											v218 = v191
											v224 = v64 + int32(4)
											F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
											mBase = m.M
											v228 = m.ExcPending
											if v228 != 0 {
												return int32(0)
											} else {
												v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v229 == int32(1) {
													F_pfree(m, v70)
													mBase = m.M
													v233 = m.ExcPending
													if v233 != 0 {
														return int32(0)
													} else {
														v234 = F_strlen(m, v224)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
														v244 = v64
														m.G0 = v14 + int32(96)
														return v244
													}
												} else {
													v234 = F_strlen(m, v224)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
													v244 = v64
													m.G0 = v14 + int32(96)
													return v244
												}
											}
										} else {
											v197 = v168 + v192
											v200 = F_palloc(m, v197+int32(2))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												v203 = v197 + int32(1)
												if v203 != 0 {
													base.MemoryFill(m, v200, int32(35), v203)
												} else {
												}
												v206 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v200+v203))) = uint8(v206)
												v211 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v200+v192))) = uint8(v211)
												v213 = v200
												v216 = v206
												v218 = v191
												v224 = v64 + int32(4)
												F_NUM_processor(m, v70, v14+int32(60), v224, v213, int32(0), v216, v218, int32(1))
												mBase = m.M
												v228 = m.ExcPending
												if v228 != 0 {
													return int32(0)
												} else {
													v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v229 == int32(1) {
														F_pfree(m, v70)
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
															return int32(0)
														} else {
															v234 = F_strlen(m, v224)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
															v244 = v64
															m.G0 = v14 + int32(96)
															return v244
														}
													} else {
														v234 = F_strlen(m, v224)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
														v244 = v64
														m.G0 = v14 + int32(96)
														return v244
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
