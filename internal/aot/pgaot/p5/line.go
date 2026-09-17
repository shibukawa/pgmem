package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_eq(m *base.Module, l0 int32) int32 {
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
	var v101 int32
	_ = v101
	var v109 float64
	_ = v109
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v130 int32
	_ = v130
	var v139 float64
	_ = v139
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v161 float64
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 float64
	_ = v165
	var v179 float64
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 float64
	_ = v183
	var v206 int32
	_ = v206
	v11 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
	if base.Ui64(base.I64_reinterpret_f64(v16)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v22)&int64(9223372036854775807)) {
			v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
			if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
				v206 = v11
				return v206
			} else {
				v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
				v163 = int64(9223372036854775807)
				v164 = base.I64_reinterpret_f64(v161) & v163
				v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
						v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
						v181 = int64(9223372036854775807)
						v182 = base.I64_reinterpret_f64(v179) & v181
						v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
						} else {
							return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
						}
					} else {
						return int32(0)
					}
				} else {
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
						v206 = v11
						return v206
					} else {
						v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
						v181 = int64(9223372036854775807)
						v182 = base.I64_reinterpret_f64(v179) & v181
						v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
						} else {
							return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
						}
					}
				}
			}
		} else {
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v28)&int64(9223372036854775807)) {
				v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
				if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
					v206 = v11
					return v206
				} else {
					v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
					v163 = int64(9223372036854775807)
					v164 = base.I64_reinterpret_f64(v161) & v163
					v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
							v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v181 = int64(9223372036854775807)
							v182 = base.I64_reinterpret_f64(v179) & v181
							v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
							} else {
								return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
							}
						} else {
							return int32(0)
						}
					} else {
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
							v206 = v11
							return v206
						} else {
							v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v181 = int64(9223372036854775807)
							v182 = base.I64_reinterpret_f64(v179) & v181
							v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
							} else {
								return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
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
						v206 = v11
						return v206
					} else {
						v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
						v163 = int64(9223372036854775807)
						v164 = base.I64_reinterpret_f64(v161) & v163
						v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
								v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
								v181 = int64(9223372036854775807)
								v182 = base.I64_reinterpret_f64(v179) & v181
								v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
								} else {
									return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
								}
							} else {
								return int32(0)
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
								v206 = v11
								return v206
							} else {
								v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
								v181 = int64(9223372036854775807)
								v182 = base.I64_reinterpret_f64(v179) & v181
								v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
								} else {
									return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
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
							v206 = v11
							return v206
						} else {
							v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
							v163 = int64(9223372036854775807)
							v164 = base.I64_reinterpret_f64(v161) & v163
							v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
									v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
									v181 = int64(9223372036854775807)
									v182 = base.I64_reinterpret_f64(v179) & v181
									v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
									} else {
										return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
									}
								} else {
									return int32(0)
								}
							} else {
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
									v206 = v11
									return v206
								} else {
									v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
									v181 = int64(9223372036854775807)
									v182 = base.I64_reinterpret_f64(v179) & v181
									v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
									} else {
										return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
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
									return int32(0)
								} else {
									v90 = v75
									v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
									v93 = math.Float64frombits(uint64(0x7ff0000000000000))
									v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
									v96 = base.F64_mul(v90, v95)
									v101 = int32(0)
									if base.B2i32(base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93) == v101)&base.F64_ne(base.F64_abs(v95), v93) == v101 {
										v109 = float64(0)
										if base.B2i32(base.F64_eq(v90, v109)|base.F64_ne(v96, v109) == int32(0))&base.F64_ne(v95, v109) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v96)), float64(1e-06)) == int32(0))&base.F64_ne(v96, v91) != 0 {
												v206 = v11
												return v206
											} else {
												v127 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
												v128 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
												v129 = F_float8_mul(m, v90, v128)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													if base.F64_ne(v127, v129)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v127, v129)), float64(1e-06)) == int32(0)) != 0 {
														v206 = v11
														return v206
													} else {
														v139 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
														v141 = F_float8_mul(m, v90, v140)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															v206 = base.F64_eq(v139, v141) | base.F64_le(base.F64_abs(base.F64_sub(v139, v141)), float64(1e-06))
															return v206
														}
													}
												}
											}
										}
									} else {
										F_float_overflow_error(m)
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
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
										return int32(0)
									} else {
										v90 = v83
										v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
										v93 = math.Float64frombits(uint64(0x7ff0000000000000))
										v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
										v96 = base.F64_mul(v90, v95)
										v101 = int32(0)
										if base.B2i32(base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93) == v101)&base.F64_ne(base.F64_abs(v95), v93) == v101 {
											v109 = float64(0)
											if base.B2i32(base.F64_eq(v90, v109)|base.F64_ne(v96, v109) == int32(0))&base.F64_ne(v95, v109) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v96)), float64(1e-06)) == int32(0))&base.F64_ne(v96, v91) != 0 {
													v206 = v11
													return v206
												} else {
													v127 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v128 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v129 = F_float8_mul(m, v90, v128)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														if base.F64_ne(v127, v129)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v127, v129)), float64(1e-06)) == int32(0)) != 0 {
															v206 = v11
															return v206
														} else {
															v139 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v140 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v141 = F_float8_mul(m, v90, v140)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																v206 = base.F64_eq(v139, v141) | base.F64_le(base.F64_abs(base.F64_sub(v139, v141)), float64(1e-06))
																return v206
															}
														}
													}
												}
											}
										} else {
											F_float_overflow_error(m)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
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
										v101 = int32(0)
										if base.B2i32(base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93) == v101)&base.F64_ne(base.F64_abs(v95), v93) == v101 {
											v109 = float64(0)
											if base.B2i32(base.F64_eq(v90, v109)|base.F64_ne(v96, v109) == int32(0))&base.F64_ne(v95, v109) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v96)), float64(1e-06)) == int32(0))&base.F64_ne(v96, v91) != 0 {
													v206 = v11
													return v206
												} else {
													v127 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
													v128 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
													v129 = F_float8_mul(m, v90, v128)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														if base.F64_ne(v127, v129)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v127, v129)), float64(1e-06)) == int32(0)) != 0 {
															v206 = v11
															return v206
														} else {
															v139 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
															v140 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
															v141 = F_float8_mul(m, v90, v140)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																v206 = base.F64_eq(v139, v141) | base.F64_le(base.F64_abs(base.F64_sub(v139, v141)), float64(1e-06))
																return v206
															}
														}
													}
												}
											}
										} else {
											F_float_overflow_error(m)
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										v88 = F_float8_div(m, v28, v44)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v90 = v88
											v91 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
											v93 = math.Float64frombits(uint64(0x7ff0000000000000))
											v95 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
											v96 = base.F64_mul(v90, v95)
											v101 = int32(0)
											if base.B2i32(base.F64_eq(base.F64_abs(v90), v93)|base.F64_ne(base.F64_abs(v96), v93) == v101)&base.F64_ne(base.F64_abs(v95), v93) == v101 {
												v109 = float64(0)
												if base.B2i32(base.F64_eq(v90, v109)|base.F64_ne(v96, v109) == int32(0))&base.F64_ne(v95, v109) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v96)), float64(1e-06)) == int32(0))&base.F64_ne(v96, v91) != 0 {
														v206 = v11
														return v206
													} else {
														v127 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
														v128 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
														v129 = F_float8_mul(m, v90, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															if base.F64_ne(v127, v129)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v127, v129)), float64(1e-06)) == int32(0)) != 0 {
																v206 = v11
																return v206
															} else {
																v139 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
																v140 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
																v141 = F_float8_mul(m, v90, v140)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	v206 = base.F64_eq(v139, v141) | base.F64_le(base.F64_abs(base.F64_sub(v139, v141)), float64(1e-06))
																	return v206
																}
															}
														}
													}
												}
											} else {
												F_float_overflow_error(m)
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
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
							v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
							if base.F64_ne(v16, v56)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807))) != 0 {
								v206 = v11
								return v206
							} else {
								v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
								v163 = int64(9223372036854775807)
								v164 = base.I64_reinterpret_f64(v161) & v163
								v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
										v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
										v181 = int64(9223372036854775807)
										v182 = base.I64_reinterpret_f64(v179) & v181
										v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
										} else {
											return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
										}
									} else {
										return int32(0)
									}
								} else {
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
										v206 = v11
										return v206
									} else {
										v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
										v181 = int64(9223372036854775807)
										v182 = base.I64_reinterpret_f64(v179) & v181
										v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
										} else {
											return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
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
			v161 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			v163 = int64(9223372036854775807)
			v164 = base.I64_reinterpret_f64(v161) & v163
			v165 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v165)&v163) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v164) {
					v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
					v181 = int64(9223372036854775807)
					v182 = base.I64_reinterpret_f64(v179) & v181
					v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
					} else {
						return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
					}
				} else {
					return int32(0)
				}
			} else {
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v164))|base.F64_ne(v165, v161) != 0 {
					v206 = v11
					return v206
				} else {
					v179 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
					v181 = int64(9223372036854775807)
					v182 = base.I64_reinterpret_f64(v179) & v181
					v183 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v183)&v181) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v182))
					} else {
						return base.B2i32(base.Ui64(v182) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v183, v179)
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_line_recv(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(24))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_pq_getmsgfloat8(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
			v12 = F_pq_getmsgfloat8(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
				v15 = F_pq_getmsgfloat8(m, v3)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v15
					v18 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
					if base.F64_le(base.F64_abs(v18), float64(1e-06)) == int32(0) {
						return v5
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
						if base.F64_le(base.F64_abs(v24), float64(1e-06)) == int32(0) {
							return v5
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50462850))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_line_recv_0), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_line_recv_1), int32(1052), int32(_a_F_line_recv_2))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
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
	}
}
func F_line_send(m *base.Module, l0 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				F_pq_sendfloat8(m, v5, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 << (uint(int32(2)) % 32)
					m.G0 = v5 + int32(16)
					return v22
				}
			}
		}
	}
}
