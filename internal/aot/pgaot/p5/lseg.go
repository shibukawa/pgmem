package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v27 float64
	_ = v27
	var v35 float64
	_ = v35
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v50 float64
	_ = v50
	var v55 float64
	_ = v55
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l1 + int32(16)
	v16 = F_point_invsl(m, l1, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return float64(0)
	} else {
		if base.F64_eq(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(-4616189618054758400)
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v27
			v77 = F_lseg_interpt_line(m, v12+int32(32), l1, v12+int32(8))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return float64(0)
			} else {
				if v77 == int32(0) {
					v84 = F_line_closept_point(m, int32(0), v12+int32(8), l1)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return float64(0)
					} else {
						v89 = F_line_closept_point(m, int32(0), v12+int32(8), v15)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return float64(0)
						} else {
							if base.F64_lt(v84, v89) != 0 {
								v92 = l1
							} else {
								v92 = v15
							}
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v93
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
							if l0 != 0 {
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
							} else {
							}
							v104 = F_point_dt(m, v12+int32(32), l2)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return float64(0)
							} else {
								m.G0 = v12 + int32(48)
								return v104
							}
						}
					}
				} else {
					if l0 != 0 {
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
					} else {
					}
					v104 = F_point_dt(m, v12+int32(32), l2)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return float64(0)
					} else {
						m.G0 = v12 + int32(48)
						return v104
					}
				}
			}
		} else {
			if base.F64_eq(v16, float64(0)) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
				v35 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v35
				v77 = F_lseg_interpt_line(m, v12+int32(32), l1, v12+int32(8))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return float64(0)
				} else {
					if v77 == int32(0) {
						v84 = F_line_closept_point(m, int32(0), v12+int32(8), l1)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return float64(0)
						} else {
							v89 = F_line_closept_point(m, int32(0), v12+int32(8), v15)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return float64(0)
							} else {
								if base.F64_lt(v84, v89) != 0 {
									v92 = l1
								} else {
									v92 = v15
								}
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v93
								v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
								*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
								if l0 != 0 {
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
									v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
								} else {
								}
								v104 = F_point_dt(m, v12+int32(32), l2)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return float64(0)
								} else {
									m.G0 = v12 + int32(48)
									return v104
								}
							}
						}
					} else {
						if l0 != 0 {
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
						} else {
						}
						v104 = F_point_dt(m, v12+int32(32), l2)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return float64(0)
						} else {
							m.G0 = v12 + int32(48)
							return v104
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
				*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
				v40 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v42 = base.F64_mul(v16, v41)
				v43 = base.F64_abs(v42)
				v44 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v43, v44)&base.F64_ne(base.F64_abs(v41), v44) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v50 = float64(0)
					if base.F64_eq(v42, v50)&base.F64_ne(v41, v50) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v55 = base.F64_sub(v40, v42)
						if base.F64_ne(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v55
							if base.F64_ne(v55, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
							}
							v77 = F_lseg_interpt_line(m, v12+int32(32), l1, v12+int32(8))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return float64(0)
							} else {
								if v77 == int32(0) {
									v84 = F_line_closept_point(m, int32(0), v12+int32(8), l1)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return float64(0)
									} else {
										v89 = F_line_closept_point(m, int32(0), v12+int32(8), v15)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return float64(0)
										} else {
											if base.F64_lt(v84, v89) != 0 {
												v92 = l1
											} else {
												v92 = v15
											}
											v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v93
											v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
											*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
											if l0 != 0 {
												v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
												v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
											} else {
											}
											v104 = F_point_dt(m, v12+int32(32), l2)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return float64(0)
											} else {
												m.G0 = v12 + int32(48)
												return v104
											}
										}
									}
								} else {
									if l0 != 0 {
										v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
										v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
									} else {
									}
									v104 = F_point_dt(m, v12+int32(32), l2)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return float64(0)
									} else {
										m.G0 = v12 + int32(48)
										return v104
									}
								}
							}
						} else {
							if base.F64_eq(base.F64_abs(v40), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v55
								if base.F64_ne(v55, float64(0)) != 0 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
								}
								v77 = F_lseg_interpt_line(m, v12+int32(32), l1, v12+int32(8))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return float64(0)
								} else {
									if v77 == int32(0) {
										v84 = F_line_closept_point(m, int32(0), v12+int32(8), l1)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return float64(0)
										} else {
											v89 = F_line_closept_point(m, int32(0), v12+int32(8), v15)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return float64(0)
											} else {
												if base.F64_lt(v84, v89) != 0 {
													v92 = l1
												} else {
													v92 = v15
												}
												v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v93
												v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
												if l0 != 0 {
													v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
													v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
												} else {
												}
												v104 = F_point_dt(m, v12+int32(32), l2)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return float64(0)
												} else {
													m.G0 = v12 + int32(48)
													return v104
												}
											}
										}
									} else {
										if l0 != 0 {
											v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
											v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
										} else {
										}
										v104 = F_point_dt(m, v12+int32(32), l2)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return float64(0)
										} else {
											m.G0 = v12 + int32(48)
											return v104
										}
									}
								}
							} else {
								if base.F64_ne(v43, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v55
									if base.F64_ne(v55, float64(0)) != 0 {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
									}
									v77 = F_lseg_interpt_line(m, v12+int32(32), l1, v12+int32(8))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return float64(0)
									} else {
										if v77 == int32(0) {
											v84 = F_line_closept_point(m, int32(0), v12+int32(8), l1)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return float64(0)
											} else {
												v89 = F_line_closept_point(m, int32(0), v12+int32(8), v15)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return float64(0)
												} else {
													if base.F64_lt(v84, v89) != 0 {
														v92 = l1
													} else {
														v92 = v15
													}
													v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v93
													v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
													*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
													if l0 != 0 {
														v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
														v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
													} else {
													}
													v104 = F_point_dt(m, v12+int32(32), l2)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return float64(0)
													} else {
														m.G0 = v12 + int32(48)
														return v104
													}
												}
											}
										} else {
											if l0 != 0 {
												v98 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v98
												v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v100
											} else {
											}
											v104 = F_point_dt(m, v12+int32(32), l2)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return float64(0)
											} else {
												m.G0 = v12 + int32(48)
												return v104
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v76 int64
	_ = v76
	var v79 float64
	_ = v79
	var v82 int64
	_ = v82
	var v89 int32
	_ = v89
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 float64
	_ = v111
	var v114 int64
	_ = v114
	var v119 int32
	_ = v119
	var v122 float64
	_ = v122
	var v130 int64
	_ = v130
	var v135 float64
	_ = v135
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 float64
	_ = v160
	var v164 int64
	_ = v164
	var v166 float64
	_ = v166
	var v169 int64
	_ = v169
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
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
			v67 = base.B2i32(base.Ui64(v22) < base.Ui64(int64(9218868437227405313)))
			v68 = int32(1)
			if base.F64_ne(v13, v19) != 0 {
				v195 = v68
				return v195
			} else {
				if v67 == int32(0) {
					v195 = v68
					return v195
				} else {
					v72 = v23
					v76 = v26
					v79 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
							v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v109 = int64(9223372036854775807)
								v110 = base.I64_reinterpret_f64(v107) & v109
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v114 = base.I64_reinterpret_f64(v111) & v109
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
									v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
									v155 = int32(0)
									if base.F64_ne(v101, v107) != 0 {
										v185 = v155
									} else {
										if v152 == int32(0) {
											v185 = v155
										} else {
											v160 = v111
											v164 = v114
											v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
												v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
											} else {
												v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
											}
										}
									}
								} else {
									v119 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
										v185 = v119
									} else {
										v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.F64_ne(v101, v107) != 0 {
												if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
													v185 = v119
												} else {
													v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
												}
											} else {
												v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
											}
										} else {
											v152 = int32(1)
											v155 = int32(0)
											if base.F64_ne(v101, v107) != 0 {
												v185 = v155
											} else {
												if v152 == int32(0) {
													v185 = v155
												} else {
													v160 = v111
													v164 = v114
													v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
														v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
													} else {
														v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
													}
												}
											}
										}
									}
								}
							} else {
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v185 = int32(0)
								} else {
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v160 = v135
									v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
									v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
										v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
									} else {
										v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
									}
								}
							}
							v195 = v185 ^ int32(1)
							return v195
						} else {
							return int32(1)
						}
					} else {
						v89 = int32(1)
						if base.F64_ne(v79, v72) != 0 {
							v195 = v89
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
								v195 = v89
							} else {
								v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v109 = int64(9223372036854775807)
									v110 = base.I64_reinterpret_f64(v107) & v109
									v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v114 = base.I64_reinterpret_f64(v111) & v109
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
										v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
										v155 = int32(0)
										if base.F64_ne(v101, v107) != 0 {
											v185 = v155
										} else {
											if v152 == int32(0) {
												v185 = v155
											} else {
												v160 = v111
												v164 = v114
												v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
													v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
												} else {
													v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
												}
											}
										}
									} else {
										v119 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
											v185 = v119
										} else {
											v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.F64_ne(v101, v107) != 0 {
													if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
														v185 = v119
													} else {
														v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
													}
												} else {
													v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
												}
											} else {
												v152 = int32(1)
												v155 = int32(0)
												if base.F64_ne(v101, v107) != 0 {
													v185 = v155
												} else {
													if v152 == int32(0) {
														v185 = v155
													} else {
														v160 = v111
														v164 = v114
														v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
															v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
														} else {
															v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
														}
													}
												}
											}
										}
									}
								} else {
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v185 = int32(0)
									} else {
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v160 = v135
										v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
										v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
											v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
										} else {
											v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
										}
									}
								}
								v195 = v185 ^ int32(1)
							}
						}
						return v195
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
					if base.F64_ne(v13, v19) != 0 {
						if base.F64_le(base.F64_abs(base.F64_sub(v13, v19)), float64(1e-06)) == int32(0) {
							v195 = v53
						} else {
							if base.F64_eq(v23, v35) != 0 {
								v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v109 = int64(9223372036854775807)
									v110 = base.I64_reinterpret_f64(v107) & v109
									v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v114 = base.I64_reinterpret_f64(v111) & v109
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
										v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
										v155 = int32(0)
										if base.F64_ne(v101, v107) != 0 {
											v185 = v155
										} else {
											if v152 == int32(0) {
												v185 = v155
											} else {
												v160 = v111
												v164 = v114
												v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
													v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
												} else {
													v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
												}
											}
										}
									} else {
										v119 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
											v185 = v119
										} else {
											v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.F64_ne(v101, v107) != 0 {
													if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
														v185 = v119
													} else {
														v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
													}
												} else {
													v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
												}
											} else {
												v152 = int32(1)
												v155 = int32(0)
												if base.F64_ne(v101, v107) != 0 {
													v185 = v155
												} else {
													if v152 == int32(0) {
														v185 = v155
													} else {
														v160 = v111
														v164 = v114
														v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
															v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
														} else {
															v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
														}
													}
												}
											}
										}
									}
								} else {
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v185 = int32(0)
									} else {
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v160 = v135
										v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
										v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
											v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
										} else {
											v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
										}
									}
								}
								v195 = v185 ^ int32(1)
							} else {
								if base.F64_le(base.F64_abs(base.F64_sub(v23, v35)), float64(1e-06)) != 0 {
									v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
										v109 = int64(9223372036854775807)
										v110 = base.I64_reinterpret_f64(v107) & v109
										v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v114 = base.I64_reinterpret_f64(v111) & v109
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
											v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
											v155 = int32(0)
											if base.F64_ne(v101, v107) != 0 {
												v185 = v155
											} else {
												if v152 == int32(0) {
													v185 = v155
												} else {
													v160 = v111
													v164 = v114
													v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
														v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
													} else {
														v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
													}
												}
											}
										} else {
											v119 = int32(0)
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
												v185 = v119
											} else {
												v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
													if base.F64_ne(v101, v107) != 0 {
														if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
															v185 = v119
														} else {
															v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
														}
													} else {
														v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
													}
												} else {
													v152 = int32(1)
													v155 = int32(0)
													if base.F64_ne(v101, v107) != 0 {
														v185 = v155
													} else {
														if v152 == int32(0) {
															v185 = v155
														} else {
															v160 = v111
															v164 = v114
															v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
															v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
															if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
																v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
															} else {
																v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
															}
														}
													}
												}
											}
										}
									} else {
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
										if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v185 = int32(0)
										} else {
											v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v160 = v135
											v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
											v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
												v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
											} else {
												v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
											}
										}
									}
									v195 = v185 ^ int32(1)
								} else {
									v195 = v53
								}
							}
						}
					} else {
						if base.F64_eq(v23, v35) != 0 {
							v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v109 = int64(9223372036854775807)
								v110 = base.I64_reinterpret_f64(v107) & v109
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v114 = base.I64_reinterpret_f64(v111) & v109
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
									v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
									v155 = int32(0)
									if base.F64_ne(v101, v107) != 0 {
										v185 = v155
									} else {
										if v152 == int32(0) {
											v185 = v155
										} else {
											v160 = v111
											v164 = v114
											v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
												v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
											} else {
												v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
											}
										}
									}
								} else {
									v119 = int32(0)
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
										v185 = v119
									} else {
										v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.F64_ne(v101, v107) != 0 {
												if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
													v185 = v119
												} else {
													v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
												}
											} else {
												v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
											}
										} else {
											v152 = int32(1)
											v155 = int32(0)
											if base.F64_ne(v101, v107) != 0 {
												v185 = v155
											} else {
												if v152 == int32(0) {
													v185 = v155
												} else {
													v160 = v111
													v164 = v114
													v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
														v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
													} else {
														v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
													}
												}
											}
										}
									}
								}
							} else {
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v185 = int32(0)
								} else {
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v160 = v135
									v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
									v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
										v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
									} else {
										v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
									}
								}
							}
							v195 = v185 ^ int32(1)
						} else {
							if base.F64_le(base.F64_abs(base.F64_sub(v23, v35)), float64(1e-06)) != 0 {
								v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v109 = int64(9223372036854775807)
									v110 = base.I64_reinterpret_f64(v107) & v109
									v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v114 = base.I64_reinterpret_f64(v111) & v109
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
										v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
										v155 = int32(0)
										if base.F64_ne(v101, v107) != 0 {
											v185 = v155
										} else {
											if v152 == int32(0) {
												v185 = v155
											} else {
												v160 = v111
												v164 = v114
												v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
													v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
												} else {
													v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
												}
											}
										}
									} else {
										v119 = int32(0)
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
											v185 = v119
										} else {
											v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.F64_ne(v101, v107) != 0 {
													if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
														v185 = v119
													} else {
														v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
													}
												} else {
													v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
												}
											} else {
												v152 = int32(1)
												v155 = int32(0)
												if base.F64_ne(v101, v107) != 0 {
													v185 = v155
												} else {
													if v152 == int32(0) {
														v185 = v155
													} else {
														v160 = v111
														v164 = v114
														v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
															v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
														} else {
															v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
														}
													}
												}
											}
										}
									}
								} else {
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v185 = int32(0)
									} else {
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v160 = v135
										v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
										v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
											v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
										} else {
											v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
										}
									}
								}
								v195 = v185 ^ int32(1)
							} else {
								v195 = v53
							}
						}
					}
					return v195
				} else {
					v67 = int32(1)
					v68 = int32(1)
					if base.F64_ne(v13, v19) != 0 {
						v195 = v68
						return v195
					} else {
						if v67 == int32(0) {
							v195 = v68
							return v195
						} else {
							v72 = v23
							v76 = v26
							v79 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
							v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
									v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
										v109 = int64(9223372036854775807)
										v110 = base.I64_reinterpret_f64(v107) & v109
										v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v114 = base.I64_reinterpret_f64(v111) & v109
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
											v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
											v155 = int32(0)
											if base.F64_ne(v101, v107) != 0 {
												v185 = v155
											} else {
												if v152 == int32(0) {
													v185 = v155
												} else {
													v160 = v111
													v164 = v114
													v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
														v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
													} else {
														v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
													}
												}
											}
										} else {
											v119 = int32(0)
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
												v185 = v119
											} else {
												v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
													if base.F64_ne(v101, v107) != 0 {
														if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
															v185 = v119
														} else {
															v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
														}
													} else {
														v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
													}
												} else {
													v152 = int32(1)
													v155 = int32(0)
													if base.F64_ne(v101, v107) != 0 {
														v185 = v155
													} else {
														if v152 == int32(0) {
															v185 = v155
														} else {
															v160 = v111
															v164 = v114
															v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
															v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
															if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
																v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
															} else {
																v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
															}
														}
													}
												}
											}
										}
									} else {
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
										if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v185 = int32(0)
										} else {
											v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v160 = v135
											v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
											v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
												v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
											} else {
												v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
											}
										}
									}
									v195 = v185 ^ int32(1)
									return v195
								} else {
									return int32(1)
								}
							} else {
								v89 = int32(1)
								if base.F64_ne(v79, v72) != 0 {
									v195 = v89
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
										v195 = v89
									} else {
										v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
										if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
											v109 = int64(9223372036854775807)
											v110 = base.I64_reinterpret_f64(v107) & v109
											v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v114 = base.I64_reinterpret_f64(v111) & v109
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
												v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
												v155 = int32(0)
												if base.F64_ne(v101, v107) != 0 {
													v185 = v155
												} else {
													if v152 == int32(0) {
														v185 = v155
													} else {
														v160 = v111
														v164 = v114
														v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
															v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
														} else {
															v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
														}
													}
												}
											} else {
												v119 = int32(0)
												if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
													v185 = v119
												} else {
													v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
														if base.F64_ne(v101, v107) != 0 {
															if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
																v185 = v119
															} else {
																v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
															}
														} else {
															v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
														}
													} else {
														v152 = int32(1)
														v155 = int32(0)
														if base.F64_ne(v101, v107) != 0 {
															v185 = v155
														} else {
															if v152 == int32(0) {
																v185 = v155
															} else {
																v160 = v111
																v164 = v114
																v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
																v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
																if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
																	v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
																} else {
																	v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
																}
															}
														}
													}
												}
											}
										} else {
											v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
											if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
												v185 = int32(0)
											} else {
												v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
												v160 = v135
												v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
												v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
													v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
												} else {
													v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
												}
											}
										}
										v195 = v185 ^ int32(1)
									}
								}
								return v195
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
			v72 = v49
			v76 = base.I64_reinterpret_f64(v49) & int64(9223372036854775807)
			v79 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v76) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
					v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
					if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
						v109 = int64(9223372036854775807)
						v110 = base.I64_reinterpret_f64(v107) & v109
						v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
						v114 = base.I64_reinterpret_f64(v111) & v109
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
							v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
							v155 = int32(0)
							if base.F64_ne(v101, v107) != 0 {
								v185 = v155
							} else {
								if v152 == int32(0) {
									v185 = v155
								} else {
									v160 = v111
									v164 = v114
									v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
										v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
									} else {
										v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
									}
								}
							}
						} else {
							v119 = int32(0)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
								v185 = v119
							} else {
								v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									if base.F64_ne(v101, v107) != 0 {
										if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
											v185 = v119
										} else {
											v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
										}
									} else {
										v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
									}
								} else {
									v152 = int32(1)
									v155 = int32(0)
									if base.F64_ne(v101, v107) != 0 {
										v185 = v155
									} else {
										if v152 == int32(0) {
											v185 = v155
										} else {
											v160 = v111
											v164 = v114
											v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
												v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
											} else {
												v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
											}
										}
									}
								}
							}
						}
					} else {
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
						if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v185 = int32(0)
						} else {
							v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v160 = v135
							v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
							v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
							v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
								v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
							} else {
								v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
							}
						}
					}
					v195 = v185 ^ int32(1)
					return v195
				} else {
					return int32(1)
				}
			} else {
				v89 = int32(1)
				if base.F64_ne(v79, v72) != 0 {
					v195 = v89
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v82) {
						v195 = v89
					} else {
						v101 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v107 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v109 = int64(9223372036854775807)
							v110 = base.I64_reinterpret_f64(v107) & v109
							v111 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v114 = base.I64_reinterpret_f64(v111) & v109
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
								v152 = base.B2i32(base.Ui64(v110) < base.Ui64(int64(9218868437227405313)))
								v155 = int32(0)
								if base.F64_ne(v101, v107) != 0 {
									v185 = v155
								} else {
									if v152 == int32(0) {
										v185 = v155
									} else {
										v160 = v111
										v164 = v114
										v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
											v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
										} else {
											v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
										}
									}
								}
							} else {
								v119 = int32(0)
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v110) {
									v185 = v119
								} else {
									v122 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v122)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.F64_ne(v101, v107) != 0 {
											if base.F64_le(base.F64_abs(base.F64_sub(v101, v107)), float64(1e-06)) == int32(0) {
												v185 = v119
											} else {
												v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
											}
										} else {
											v185 = base.F64_eq(v111, v122) | base.F64_le(base.F64_abs(base.F64_sub(v111, v122)), float64(1e-06))
										}
									} else {
										v152 = int32(1)
										v155 = int32(0)
										if base.F64_ne(v101, v107) != 0 {
											v185 = v155
										} else {
											if v152 == int32(0) {
												v185 = v155
											} else {
												v160 = v111
												v164 = v114
												v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
													v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
												} else {
													v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
												}
											}
										}
									}
								}
							}
						} else {
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v130&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v185 = int32(0)
							} else {
								v135 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v160 = v135
								v164 = base.I64_reinterpret_f64(v135) & int64(9223372036854775807)
								v166 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v169 = base.I64_reinterpret_f64(v166) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v164) {
									v185 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v169))
								} else {
									v185 = base.B2i32(base.Ui64(v169) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v166, v160)
								}
							}
						}
						v195 = v185 ^ int32(1)
					}
				}
				return v195
			}
		}
	}
}
