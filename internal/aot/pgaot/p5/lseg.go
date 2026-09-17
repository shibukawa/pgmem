package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v28 float64
	_ = v28
	var v36 float64
	_ = v36
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = l1 + int32(16)
	v17 = F_point_invsl(m, l1, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return float64(0)
	} else {
		if base.F64_eq(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(-4616189618054758400)
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v28
			v81 = v13 + int32(8)
			v82 = F_lseg_interpt_line(m, v13+int32(32), l1, v81)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return float64(0)
			} else {
				if v82 == int32(0) {
					v87 = F_line_closept_point(m, int32(0), v81, l1)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return float64(0)
					} else {
						v90 = F_line_closept_point(m, int32(0), v81, v16)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return float64(0)
						} else {
							if base.F64_lt(v87, v90) != 0 {
								v93 = l1
							} else {
								v93 = v16
							}
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v96
							if l0 != 0 {
								v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
								v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
							} else {
							}
							v105 = F_point_dt(m, v13+int32(32), l2)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return float64(0)
							} else {
								m.G0 = v13 + int32(48)
								return v105
							}
						}
					}
				} else {
					if l0 != 0 {
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
					} else {
					}
					v105 = F_point_dt(m, v13+int32(32), l2)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return float64(0)
					} else {
						m.G0 = v13 + int32(48)
						return v105
					}
				}
			}
		} else {
			if base.F64_eq(v17, float64(0)) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(-4616189618054758400)
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
				v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v36
				v81 = v13 + int32(8)
				v82 = F_lseg_interpt_line(m, v13+int32(32), l1, v81)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return float64(0)
				} else {
					if v82 == int32(0) {
						v87 = F_line_closept_point(m, int32(0), v81, l1)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return float64(0)
						} else {
							v90 = F_line_closept_point(m, int32(0), v81, v16)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return float64(0)
							} else {
								if base.F64_lt(v87, v90) != 0 {
									v93 = l1
								} else {
									v93 = v16
								}
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v96
								if l0 != 0 {
									v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
									v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
								} else {
								}
								v105 = F_point_dt(m, v13+int32(32), l2)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return float64(0)
								} else {
									m.G0 = v13 + int32(48)
									return v105
								}
							}
						}
					} else {
						if l0 != 0 {
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
							v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
						} else {
						}
						v105 = F_point_dt(m, v13+int32(32), l2)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return float64(0)
						} else {
							m.G0 = v13 + int32(48)
							return v105
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(-4616189618054758400)
				*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v17
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v43 = base.F64_mul(v17, v42)
				v44 = base.F64_abs(v43)
				v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v44, v45)&base.F64_ne(base.F64_abs(v42), v45) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v51 = float64(0)
					if base.F64_eq(v43, v51)&base.F64_ne(v42, v51) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v57 = math.Float64frombits(uint64(0x7ff0000000000000))
						v59 = base.F64_sub(v41, v43)
						if base.B2i32(base.F64_eq(base.F64_abs(v41), v57)|base.F64_ne(base.F64_abs(v59), v57) == int32(0))&base.F64_ne(v44, v57) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v59
							if base.F64_ne(v59, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
							}
							v81 = v13 + int32(8)
							v82 = F_lseg_interpt_line(m, v13+int32(32), l1, v81)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return float64(0)
							} else {
								if v82 == int32(0) {
									v87 = F_line_closept_point(m, int32(0), v81, l1)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return float64(0)
									} else {
										v90 = F_line_closept_point(m, int32(0), v81, v16)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return float64(0)
										} else {
											if base.F64_lt(v87, v90) != 0 {
												v93 = l1
											} else {
												v93 = v16
											}
											v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v94
											v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
											*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v96
											if l0 != 0 {
												v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
												v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
											} else {
											}
											v105 = F_point_dt(m, v13+int32(32), l2)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return float64(0)
											} else {
												m.G0 = v13 + int32(48)
												return v105
											}
										}
									}
								} else {
									if l0 != 0 {
										v99 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v99
										v101 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v101
									} else {
									}
									v105 = F_point_dt(m, v13+int32(32), l2)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return float64(0)
									} else {
										m.G0 = v13 + int32(48)
										return v105
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
func F_lseg_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_point_dt(m, v4, v4+int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = F_point_dt(m, v3, v3+int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return base.F64_gt(v7, base.F64_add(v13, float64(1e-06)))
		}
	}
}
func F_lseg_intersect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_lseg_interpt_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_lseg_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_point_dt(m, v4, v4+int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = F_point_dt(m, v3, v3+int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return base.F64_le(v7, base.F64_add(v13, float64(1e-06)))
		}
	}
}
func F_lseg_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 float64
	_ = v23
	var v26 int64
	_ = v26
	var v35 float64
	_ = v35
	var v42 int64
	_ = v42
	var v49 float64
	_ = v49
	var v53 int32
	_ = v53
	var v70 int32
	_ = v70
	var v76 float64
	_ = v76
	var v80 int64
	_ = v80
	var v83 float64
	_ = v83
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v105 float64
	_ = v105
	var v111 float64
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 float64
	_ = v115
	var v118 int64
	_ = v118
	var v123 int32
	_ = v123
	var v126 float64
	_ = v126
	var v134 int64
	_ = v134
	var v139 float64
	_ = v139
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 float64
	_ = v166
	var v170 int64
	_ = v170
	var v172 float64
	_ = v172
	var v175 int64
	_ = v175
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = int64(9223372036854775807)
		v22 = base.I64_reinterpret_f64(v19) & v21
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		v26 = base.I64_reinterpret_f64(v23) & v21
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v26) {
			v70 = base.B2i32(base.Ui64(v22) < base.Ui64(int64(9218868437227405313)))
			if base.B2i32(v70 == int32(0))|base.F64_ne(v13, v19) != 0 {
				v201 = int32(1)
				return v201
			} else {
				v76 = v23
				v80 = v26
				v83 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v86 = base.I64_reinterpret_f64(v83) & int64(9223372036854775807)
				if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
					v89 = int32(1)
					if base.F64_ne(v83, v76) != 0 {
						v201 = v89
					} else {
						if base.Ui64(v86) < base.Ui64(int64(9218868437227405313)) {
							v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v113 = int64(9223372036854775807)
								v114 = base.I64_reinterpret_f64(v111) & v113
								v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v118 = base.I64_reinterpret_f64(v115) & v113
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
									v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
									v160 = int32(0)
									if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
										v191 = v160
									} else {
										v166 = v115
										v170 = v118
										v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
											v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
										} else {
											v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
										}
									}
								} else {
									v123 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
										v191 = v123
									} else {
										v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
												v191 = v123
											} else {
												v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
											}
										} else {
											v157 = int32(1)
											v160 = int32(0)
											if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
												v191 = v160
											} else {
												v166 = v115
												v170 = v118
												v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
													v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
												} else {
													v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
												}
											}
										}
									}
								}
							} else {
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v191 = int32(0)
								} else {
									v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v166 = v139
									v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
									v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
										v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
									} else {
										v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
									}
								}
							}
							v201 = v191 ^ int32(1)
						} else {
							v201 = v89
						}
					}
					return v201
				} else {
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v86) {
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v113 = int64(9223372036854775807)
							v114 = base.I64_reinterpret_f64(v111) & v113
							v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v118 = base.I64_reinterpret_f64(v115) & v113
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
								v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
								v160 = int32(0)
								if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
									v191 = v160
								} else {
									v166 = v115
									v170 = v118
									v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
										v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
									} else {
										v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
									}
								}
							} else {
								v123 = int32(0)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
									v191 = v123
								} else {
									v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
											v191 = v123
										} else {
											v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
										}
									} else {
										v157 = int32(1)
										v160 = int32(0)
										if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
											v191 = v160
										} else {
											v166 = v115
											v170 = v118
											v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
												v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
											} else {
												v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
											}
										}
									}
								}
							}
						} else {
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v191 = int32(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v166 = v139
								v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
								v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
									v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
								} else {
									v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
								}
							}
						}
						v201 = v191 ^ int32(1)
						return v201
					} else {
						return int32(1)
					}
				}
			}
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v22) {
				return int32(1)
			} else {
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				if base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					v53 = int32(1)
					if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v13, v19)), float64(1e-06)) == int32(0))&base.F64_ne(v13, v19) != 0 {
						v201 = v53
					} else {
						if base.F64_eq(v23, v35) != 0 {
							v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v113 = int64(9223372036854775807)
								v114 = base.I64_reinterpret_f64(v111) & v113
								v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v118 = base.I64_reinterpret_f64(v115) & v113
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
									v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
									v160 = int32(0)
									if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
										v191 = v160
									} else {
										v166 = v115
										v170 = v118
										v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
											v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
										} else {
											v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
										}
									}
								} else {
									v123 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
										v191 = v123
									} else {
										v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
												v191 = v123
											} else {
												v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
											}
										} else {
											v157 = int32(1)
											v160 = int32(0)
											if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
												v191 = v160
											} else {
												v166 = v115
												v170 = v118
												v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
													v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
												} else {
													v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
												}
											}
										}
									}
								}
							} else {
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v191 = int32(0)
								} else {
									v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v166 = v139
									v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
									v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
										v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
									} else {
										v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
									}
								}
							}
							v201 = v191 ^ int32(1)
						} else {
							if base.F64_le(base.F64_abs(base.F64_sub(v23, v35)), float64(1e-06)) == int32(0) {
								v201 = v53
							} else {
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v113 = int64(9223372036854775807)
									v114 = base.I64_reinterpret_f64(v111) & v113
									v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v118 = base.I64_reinterpret_f64(v115) & v113
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
										v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
										v160 = int32(0)
										if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
											v191 = v160
										} else {
											v166 = v115
											v170 = v118
											v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
												v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
											} else {
												v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
											}
										}
									} else {
										v123 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
											v191 = v123
										} else {
											v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
													v191 = v123
												} else {
													v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
												}
											} else {
												v157 = int32(1)
												v160 = int32(0)
												if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
													v191 = v160
												} else {
													v166 = v115
													v170 = v118
													v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
														v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
													} else {
														v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
													}
												}
											}
										}
									}
								} else {
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v191 = int32(0)
									} else {
										v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v166 = v139
										v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
										v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
											v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
										} else {
											v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
										}
									}
								}
								v201 = v191 ^ int32(1)
							}
						}
					}
					return v201
				} else {
					v70 = int32(1)
					if base.B2i32(v70 == int32(0))|base.F64_ne(v13, v19) != 0 {
						v201 = int32(1)
						return v201
					} else {
						v76 = v23
						v80 = v26
						v83 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v86 = base.I64_reinterpret_f64(v83) & int64(9223372036854775807)
						if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
							v89 = int32(1)
							if base.F64_ne(v83, v76) != 0 {
								v201 = v89
							} else {
								if base.Ui64(v86) < base.Ui64(int64(9218868437227405313)) {
									v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
										v113 = int64(9223372036854775807)
										v114 = base.I64_reinterpret_f64(v111) & v113
										v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v118 = base.I64_reinterpret_f64(v115) & v113
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
											v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
											v160 = int32(0)
											if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
												v191 = v160
											} else {
												v166 = v115
												v170 = v118
												v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
													v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
												} else {
													v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
												}
											}
										} else {
											v123 = int32(0)
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
												v191 = v123
											} else {
												v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
														v191 = v123
													} else {
														v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
													}
												} else {
													v157 = int32(1)
													v160 = int32(0)
													if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
														v191 = v160
													} else {
														v166 = v115
														v170 = v118
														v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
															v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
														} else {
															v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
														}
													}
												}
											}
										}
									} else {
										v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
										if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v191 = int32(0)
										} else {
											v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v166 = v139
											v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
											v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
												v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
											} else {
												v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
											}
										}
									}
									v201 = v191 ^ int32(1)
								} else {
									v201 = v89
								}
							}
							return v201
						} else {
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v86) {
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v113 = int64(9223372036854775807)
									v114 = base.I64_reinterpret_f64(v111) & v113
									v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v118 = base.I64_reinterpret_f64(v115) & v113
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
										v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
										v160 = int32(0)
										if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
											v191 = v160
										} else {
											v166 = v115
											v170 = v118
											v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
												v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
											} else {
												v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
											}
										}
									} else {
										v123 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
											v191 = v123
										} else {
											v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
													v191 = v123
												} else {
													v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
												}
											} else {
												v157 = int32(1)
												v160 = int32(0)
												if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
													v191 = v160
												} else {
													v166 = v115
													v170 = v118
													v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
														v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
													} else {
														v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
													}
												}
											}
										}
									}
								} else {
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v191 = int32(0)
									} else {
										v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v166 = v139
										v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
										v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
											v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
										} else {
											v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
										}
									}
								}
								v201 = v191 ^ int32(1)
								return v201
							} else {
								return int32(1)
							}
						}
					}
				}
			}
		}
	} else {
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if base.Ui64(v42&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			return int32(1)
		} else {
			v49 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v76 = v49
			v80 = base.I64_reinterpret_f64(v49) & int64(9223372036854775807)
			v83 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v86 = base.I64_reinterpret_f64(v83) & int64(9223372036854775807)
			if base.Ui64(v80) <= base.Ui64(int64(9218868437227405312)) {
				v89 = int32(1)
				if base.F64_ne(v83, v76) != 0 {
					v201 = v89
				} else {
					if base.Ui64(v86) < base.Ui64(int64(9218868437227405313)) {
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v113 = int64(9223372036854775807)
							v114 = base.I64_reinterpret_f64(v111) & v113
							v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v118 = base.I64_reinterpret_f64(v115) & v113
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
								v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
								v160 = int32(0)
								if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
									v191 = v160
								} else {
									v166 = v115
									v170 = v118
									v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
										v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
									} else {
										v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
									}
								}
							} else {
								v123 = int32(0)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
									v191 = v123
								} else {
									v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
											v191 = v123
										} else {
											v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
										}
									} else {
										v157 = int32(1)
										v160 = int32(0)
										if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
											v191 = v160
										} else {
											v166 = v115
											v170 = v118
											v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
												v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
											} else {
												v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
											}
										}
									}
								}
							}
						} else {
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v191 = int32(0)
							} else {
								v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v166 = v139
								v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
								v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
									v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
								} else {
									v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
								}
							}
						}
						v201 = v191 ^ int32(1)
					} else {
						v201 = v89
					}
				}
				return v201
			} else {
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v86) {
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
					if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						v111 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
						v113 = int64(9223372036854775807)
						v114 = base.I64_reinterpret_f64(v111) & v113
						v115 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
						v118 = base.I64_reinterpret_f64(v115) & v113
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
							v157 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
							v160 = int32(0)
							if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
								v191 = v160
							} else {
								v166 = v115
								v170 = v118
								v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
									v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
								} else {
									v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
								}
							}
						} else {
							v123 = int32(0)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
								v191 = v123
							} else {
								v126 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								if base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0))&base.F64_ne(v105, v111) != 0 {
										v191 = v123
									} else {
										v191 = base.F64_eq(v115, v126) | base.F64_le(base.F64_abs(base.F64_sub(v115, v126)), float64(1e-06))
									}
								} else {
									v157 = int32(1)
									v160 = int32(0)
									if base.B2i32(v157 == v160)|base.F64_ne(v105, v111) != 0 {
										v191 = v160
									} else {
										v166 = v115
										v170 = v118
										v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
											v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
										} else {
											v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
										}
									}
								}
							}
						}
					} else {
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
						if base.Ui64(v134&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v191 = int32(0)
						} else {
							v139 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v166 = v139
							v170 = base.I64_reinterpret_f64(v139) & int64(9223372036854775807)
							v172 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
							v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v170) {
								v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v175))
							} else {
								v191 = base.B2i32(base.Ui64(v175) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v172, v166)
							}
						}
					}
					v201 = v191 ^ int32(1)
					return v201
				} else {
					return int32(1)
				}
			}
		}
	}
}
