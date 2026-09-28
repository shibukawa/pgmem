package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v28 float64
	_ = v28
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v56 float64
	_ = v56
	var v64 int64
	_ = v64
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v121 float64
	_ = v121
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v133 int32
	_ = v133
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v160 float64
	_ = v160
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 float64
	_ = v164
	var v178 float64
	_ = v178
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 float64
	_ = v182
	var v207 int32
	_ = v207
	v11 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
	if base.Ui64(base.I64_reinterpret_f64(v16)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v22)&int64(9223372036854775807)) {
			v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
			if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
				v207 = v11
				return base.I64_extend_i32_u(v207)
			} else {
				v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
				v162 = int64(9223372036854775807)
				v163 = base.I64_reinterpret_f64(v160) & v162
				v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
						v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
						v180 = int64(9223372036854775807)
						v181 = base.I64_reinterpret_f64(v178) & v180
						v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
						} else {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
						}
					} else {
						return int64(0)
					}
				} else {
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
						v207 = v11
						return base.I64_extend_i32_u(v207)
					} else {
						v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
						v180 = int64(9223372036854775807)
						v181 = base.I64_reinterpret_f64(v178) & v180
						v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
						} else {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
						}
					}
				}
			}
		} else {
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v28)&int64(9223372036854775807)) {
				v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
				if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
					v207 = v11
					return base.I64_extend_i32_u(v207)
				} else {
					v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
					v162 = int64(9223372036854775807)
					v163 = base.I64_reinterpret_f64(v160) & v162
					v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
							v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v180 = int64(9223372036854775807)
							v181 = base.I64_reinterpret_f64(v178) & v180
							v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
							} else {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
							}
						} else {
							return int64(0)
						}
					} else {
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
							v207 = v11
							return base.I64_extend_i32_u(v207)
						} else {
							v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v180 = int64(9223372036854775807)
							v181 = base.I64_reinterpret_f64(v178) & v180
							v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
							} else {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
							}
						}
					}
				}
			} else {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
				v35 = base.F64_abs(v34)
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)) {
					v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
					if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
						v207 = v11
						return base.I64_extend_i32_u(v207)
					} else {
						v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
						v162 = int64(9223372036854775807)
						v163 = base.I64_reinterpret_f64(v160) & v162
						v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
								v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
								v180 = int64(9223372036854775807)
								v181 = base.I64_reinterpret_f64(v178) & v180
								v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
								} else {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
								}
							} else {
								return int64(0)
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
								v207 = v11
								return base.I64_extend_i32_u(v207)
							} else {
								v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
								v180 = int64(9223372036854775807)
								v181 = base.I64_reinterpret_f64(v178) & v180
								v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
								} else {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
								}
							}
						}
					}
				} else {
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
					v40 = base.F64_abs(v39)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)) {
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
						if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
							v207 = v11
							return base.I64_extend_i32_u(v207)
						} else {
							v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
							v162 = int64(9223372036854775807)
							v163 = base.I64_reinterpret_f64(v160) & v162
							v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
									v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
									v180 = int64(9223372036854775807)
									v181 = base.I64_reinterpret_f64(v178) & v180
									v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
									} else {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
									}
								} else {
									return int64(0)
								}
							} else {
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
									v207 = v11
									return base.I64_extend_i32_u(v207)
								} else {
									v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
									v180 = int64(9223372036854775807)
									v181 = base.I64_reinterpret_f64(v178) & v180
									v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
									} else {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
									}
								}
							}
						}
					} else {
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
						v45 = base.F64_abs(v44)
						if base.Ui64(base.I64_reinterpret_f64(v45)) <= base.Ui64(int64(9218868437227405312)) {
							if base.F64_le(v35, float64(1e-06)) == int32(0) {
								v75 = F_float8_div(m, v16, v34)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v90 = v75
									v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
									v93 = math.Float64frombits(uint64(0x7ff0000000000000))
									v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
									v96 = base.F64_mul(v90, v95)
									if base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93)|base.F64_eq(base.F64_abs(v95), v93) == int32(0) {
										v108 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											v121 = v108
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
												v207 = v11
												return base.I64_extend_i32_u(v207)
											} else {
												v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
												v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
												v132 = F_float8_mul(m, v90, v131)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
														v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
														v144 = F_float8_mul(m, v90, v143)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int64(0)
														} else {
															v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
															return base.I64_extend_i32_u(v207)
														}
													}
												}
											}
										}
									} else {
										v110 = float64(0)
										if base.F64_eq(v90, v110)|base.F64_ne(v96, v110)|base.F64_eq(v95, v110) != 0 {
											v121 = v96
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
												v207 = v11
												return base.I64_extend_i32_u(v207)
											} else {
												v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
												v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
												v132 = F_float8_mul(m, v90, v131)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
														v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
														v144 = F_float8_mul(m, v90, v143)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int64(0)
														} else {
															v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
															return base.I64_extend_i32_u(v207)
														}
													}
												}
											}
										} else {
											v119 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int64(0)
											} else {
												v121 = v119
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
													v207 = v11
													return base.I64_extend_i32_u(v207)
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v132 = F_float8_mul(m, v90, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v144 = F_float8_mul(m, v90, v143)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int64(0)
															} else {
																v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																return base.I64_extend_i32_u(v207)
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								if base.F64_le(v40, float64(1e-06)) == int32(0) {
									v83 = F_float8_div(m, v22, v39)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int64(0)
									} else {
										v90 = v83
										v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
										v93 = math.Float64frombits(uint64(0x7ff0000000000000))
										v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
										v96 = base.F64_mul(v90, v95)
										if base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93)|base.F64_eq(base.F64_abs(v95), v93) == int32(0) {
											v108 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v121 = v108
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
													v207 = v11
													return base.I64_extend_i32_u(v207)
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v132 = F_float8_mul(m, v90, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v144 = F_float8_mul(m, v90, v143)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int64(0)
															} else {
																v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																return base.I64_extend_i32_u(v207)
															}
														}
													}
												}
											}
										} else {
											v110 = float64(0)
											if base.F64_eq(v90, v110)|base.F64_ne(v96, v110)|base.F64_eq(v95, v110) != 0 {
												v121 = v96
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
													v207 = v11
													return base.I64_extend_i32_u(v207)
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v132 = F_float8_mul(m, v90, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v144 = F_float8_mul(m, v90, v143)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int64(0)
															} else {
																v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																return base.I64_extend_i32_u(v207)
															}
														}
													}
												}
											} else {
												v119 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int64(0)
												} else {
													v121 = v119
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
														v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
														v132 = F_float8_mul(m, v90, v131)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int64(0)
														} else {
															if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
																v207 = v11
																return base.I64_extend_i32_u(v207)
															} else {
																v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																v144 = F_float8_mul(m, v90, v143)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int64(0)
																} else {
																	v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																	return base.I64_extend_i32_u(v207)
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									if base.F64_le(v45, float64(1e-06)) != 0 {
										v90 = float64(1)
										v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
										v93 = math.Float64frombits(uint64(0x7ff0000000000000))
										v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
										v96 = base.F64_mul(v90, v95)
										if base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93)|base.F64_eq(base.F64_abs(v95), v93) == int32(0) {
											v108 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v121 = v108
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
													v207 = v11
													return base.I64_extend_i32_u(v207)
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v132 = F_float8_mul(m, v90, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v144 = F_float8_mul(m, v90, v143)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int64(0)
															} else {
																v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																return base.I64_extend_i32_u(v207)
															}
														}
													}
												}
											}
										} else {
											v110 = float64(0)
											if base.F64_eq(v90, v110)|base.F64_ne(v96, v110)|base.F64_eq(v95, v110) != 0 {
												v121 = v96
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
													v207 = v11
													return base.I64_extend_i32_u(v207)
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v132 = F_float8_mul(m, v90, v131)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v144 = F_float8_mul(m, v90, v143)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int64(0)
															} else {
																v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																return base.I64_extend_i32_u(v207)
															}
														}
													}
												}
											} else {
												v119 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int64(0)
												} else {
													v121 = v119
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
														v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
														v132 = F_float8_mul(m, v90, v131)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int64(0)
														} else {
															if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
																v207 = v11
																return base.I64_extend_i32_u(v207)
															} else {
																v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																v144 = F_float8_mul(m, v90, v143)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int64(0)
																} else {
																	v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																	return base.I64_extend_i32_u(v207)
																}
															}
														}
													}
												}
											}
										}
									} else {
										v88 = F_float8_div(m, v28, v44)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int64(0)
										} else {
											v90 = v88
											v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
											v93 = math.Float64frombits(uint64(0x7ff0000000000000))
											v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
											v96 = base.F64_mul(v90, v95)
											if base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93)|base.F64_eq(base.F64_abs(v95), v93) == int32(0) {
												v108 = F_float_overflow_error_ext(m, int32(0))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return int64(0)
												} else {
													v121 = v108
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
														v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
														v132 = F_float8_mul(m, v90, v131)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int64(0)
														} else {
															if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
																v207 = v11
																return base.I64_extend_i32_u(v207)
															} else {
																v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																v144 = F_float8_mul(m, v90, v143)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int64(0)
																} else {
																	v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																	return base.I64_extend_i32_u(v207)
																}
															}
														}
													}
												}
											} else {
												v110 = float64(0)
												if base.F64_eq(v90, v110)|base.F64_ne(v96, v110)|base.F64_eq(v95, v110) != 0 {
													v121 = v96
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
														v207 = v11
														return base.I64_extend_i32_u(v207)
													} else {
														v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
														v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
														v132 = F_float8_mul(m, v90, v131)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int64(0)
														} else {
															if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
																v207 = v11
																return base.I64_extend_i32_u(v207)
															} else {
																v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																v144 = F_float8_mul(m, v90, v143)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int64(0)
																} else {
																	v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																	return base.I64_extend_i32_u(v207)
																}
															}
														}
													}
												} else {
													v119 = F_float_underflow_error_ext(m, int32(0))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int64(0)
													} else {
														v121 = v119
														if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v121)), float64(1e-06)) == int32(0))&base.F64_ne(v121, v91) != 0 {
															v207 = v11
															return base.I64_extend_i32_u(v207)
														} else {
															v130 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
															v131 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
															v132 = F_float8_mul(m, v90, v131)
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int64(0)
															} else {
																if base.F64_ne(v130, v132)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v130, v132)), float64(1e-06)) == int32(0)) != 0 {
																	v207 = v11
																	return base.I64_extend_i32_u(v207)
																} else {
																	v142 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																	v143 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																	v144 = F_float8_mul(m, v90, v143)
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int64(0)
																	} else {
																		v207 = base.F64_eq(v142, v144) | base.F64_le(base.F64_abs(base.F64_sub(v142, v144)), float64(1e-06))
																		return base.I64_extend_i32_u(v207)
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
							v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
							if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
								v207 = v11
								return base.I64_extend_i32_u(v207)
							} else {
								v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
								v162 = int64(9223372036854775807)
								v163 = base.I64_reinterpret_f64(v160) & v162
								v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
										v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
										v180 = int64(9223372036854775807)
										v181 = base.I64_reinterpret_f64(v178) & v180
										v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
											return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
										} else {
											return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
										}
									} else {
										return int64(0)
									}
								} else {
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
										v207 = v11
										return base.I64_extend_i32_u(v207)
									} else {
										v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
										v180 = int64(9223372036854775807)
										v181 = base.I64_reinterpret_f64(v178) & v180
										v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
											return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
										} else {
											return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
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
		v64 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v64&int64(9223372036854775807)) {
			v160 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			v162 = int64(9223372036854775807)
			v163 = base.I64_reinterpret_f64(v160) & v162
			v164 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v164)&v162) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v163) {
					v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
					v180 = int64(9223372036854775807)
					v181 = base.I64_reinterpret_f64(v178) & v180
					v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
					} else {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
					}
				} else {
					return int64(0)
				}
			} else {
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v163))|base.F64_ne(v164, v160) != 0 {
					v207 = v11
					return base.I64_extend_i32_u(v207)
				} else {
					v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
					v180 = int64(9223372036854775807)
					v181 = base.I64_reinterpret_f64(v178) & v180
					v182 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v182)&v180) {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v181)))
					} else {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(v181) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v182, v178))
					}
				}
			}
		} else {
			return int64(0)
		}
	}
}
func F_line_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(24))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_pq_getmsgfloat8(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
			v12 = F_pq_getmsgfloat8(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
				v15 = F_pq_getmsgfloat8(m, v3)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v15
					v18 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
					if base.F64_le(base.F64_abs(v18), float64(1e-06)) == int32(0) {
						return base.I64_extend_i32_u(v5)
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
						if base.F64_le(base.F64_abs(v24), float64(1e-06)) == int32(0) {
							return base.I64_extend_i32_u(v5)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50462850))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_line_recv_0), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_line_recv_1), int32(1098), int32(_a_F_line_recv_2))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int64(0)
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
	}
}
func F_line_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				F_pq_sendfloat8(m, v5, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 << (uint(int32(2)) % 32)
					m.G0 = v5 + int32(16)
					return base.I64_extend_i32_u(v22)
				}
			}
		}
	}
}
