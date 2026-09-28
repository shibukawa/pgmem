package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_horizontal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.I64_extend_i32_u(base.F64_eq(v5, v6) | base.F64_le(base.F64_abs(base.F64_sub(v5, v6)), float64(1e-06)))
}
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v44 float64
	_ = v44
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = F_point_sl(m, l2, l2+int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if base.F64_eq(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(-4616189618054758400)
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v27
			v91 = v12 + int32(32)
			v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				if v94 == int32(0) {
					v109 = v4
					m.G0 = v12 + int32(48)
					return v109
				} else {
					v98 = F_lseg_contain_point(m, l2, v91)
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						if v98 == int32(0) {
							v109 = v4
						} else {
							v102 = int32(1)
							if l0 == int32(0) {
								v109 = v102
							} else {
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
								v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
								v109 = v102
							}
						}
						m.G0 = v12 + int32(48)
						return v109
					}
				}
			}
		} else {
			if base.F64_eq(v16, float64(0)) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
				v35 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v35
				v91 = v12 + int32(32)
				v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					if v94 == int32(0) {
						v109 = v4
						m.G0 = v12 + int32(48)
						return v109
					} else {
						v98 = F_lseg_contain_point(m, l2, v91)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							if v98 == int32(0) {
								v109 = v4
							} else {
								v102 = int32(1)
								if l0 == int32(0) {
									v109 = v102
								} else {
									v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
									v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
									v109 = v102
								}
							}
							m.G0 = v12 + int32(48)
							return v109
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(-4616189618054758400)
				*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
				v40 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v42 = base.F64_mul(v16, v41)
				v44 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
					v53 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v63 = v53
						v65 = math.Float64frombits(uint64(0x7ff0000000000000))
						v67 = base.F64_sub(v40, v63)
						if base.F64_eq(base.F64_abs(v40), v65)|base.F64_ne(base.F64_abs(v67), v65)|base.F64_eq(base.F64_abs(v63), v65) == int32(0) {
							v79 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = v79
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
								if base.F64_ne(v81, float64(0)) != 0 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
								}
								v91 = v12 + int32(32)
								v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									if v94 == int32(0) {
										v109 = v4
										m.G0 = v12 + int32(48)
										return v109
									} else {
										v98 = F_lseg_contain_point(m, l2, v91)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											if v98 == int32(0) {
												v109 = v4
											} else {
												v102 = int32(1)
												if l0 == int32(0) {
													v109 = v102
												} else {
													v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
													v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
													v109 = v102
												}
											}
											m.G0 = v12 + int32(48)
											return v109
										}
									}
								}
							}
						} else {
							v81 = v67
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
							if base.F64_ne(v81, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
							}
							v91 = v12 + int32(32)
							v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								if v94 == int32(0) {
									v109 = v4
									m.G0 = v12 + int32(48)
									return v109
								} else {
									v98 = F_lseg_contain_point(m, l2, v91)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										if v98 == int32(0) {
											v109 = v4
										} else {
											v102 = int32(1)
											if l0 == int32(0) {
												v109 = v102
											} else {
												v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
												v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
												v109 = v102
											}
										}
										m.G0 = v12 + int32(48)
										return v109
									}
								}
							}
						}
					}
				} else {
					v55 = float64(0)
					if base.F64_eq(v41, v55)|base.F64_ne(v42, v55) != 0 {
						v63 = v42
						v65 = math.Float64frombits(uint64(0x7ff0000000000000))
						v67 = base.F64_sub(v40, v63)
						if base.F64_eq(base.F64_abs(v40), v65)|base.F64_ne(base.F64_abs(v67), v65)|base.F64_eq(base.F64_abs(v63), v65) == int32(0) {
							v79 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v81 = v79
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
								if base.F64_ne(v81, float64(0)) != 0 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
								}
								v91 = v12 + int32(32)
								v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									if v94 == int32(0) {
										v109 = v4
										m.G0 = v12 + int32(48)
										return v109
									} else {
										v98 = F_lseg_contain_point(m, l2, v91)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											if v98 == int32(0) {
												v109 = v4
											} else {
												v102 = int32(1)
												if l0 == int32(0) {
													v109 = v102
												} else {
													v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
													v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
													v109 = v102
												}
											}
											m.G0 = v12 + int32(48)
											return v109
										}
									}
								}
							}
						} else {
							v81 = v67
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
							if base.F64_ne(v81, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
							}
							v91 = v12 + int32(32)
							v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								if v94 == int32(0) {
									v109 = v4
									m.G0 = v12 + int32(48)
									return v109
								} else {
									v98 = F_lseg_contain_point(m, l2, v91)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										if v98 == int32(0) {
											v109 = v4
										} else {
											v102 = int32(1)
											if l0 == int32(0) {
												v109 = v102
											} else {
												v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
												v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
												v109 = v102
											}
										}
										m.G0 = v12 + int32(48)
										return v109
									}
								}
							}
						}
					} else {
						v61 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = v61
							v65 = math.Float64frombits(uint64(0x7ff0000000000000))
							v67 = base.F64_sub(v40, v63)
							if base.F64_eq(base.F64_abs(v40), v65)|base.F64_ne(base.F64_abs(v67), v65)|base.F64_eq(base.F64_abs(v63), v65) == int32(0) {
								v79 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v81 = v79
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
									if base.F64_ne(v81, float64(0)) != 0 {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
									}
									v91 = v12 + int32(32)
									v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										if v94 == int32(0) {
											v109 = v4
											m.G0 = v12 + int32(48)
											return v109
										} else {
											v98 = F_lseg_contain_point(m, l2, v91)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												if v98 == int32(0) {
													v109 = v4
												} else {
													v102 = int32(1)
													if l0 == int32(0) {
														v109 = v102
													} else {
														v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
														v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
														v109 = v102
													}
												}
												m.G0 = v12 + int32(48)
												return v109
											}
										}
									}
								}
							} else {
								v81 = v67
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v81
								if base.F64_ne(v81, float64(0)) != 0 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
								}
								v91 = v12 + int32(32)
								v94 = F_lseg_interpt_line(m, v91, l1, v12+int32(8))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									if v94 == int32(0) {
										v109 = v4
										m.G0 = v12 + int32(48)
										return v109
									} else {
										v98 = F_lseg_contain_point(m, l2, v91)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											if v98 == int32(0) {
												v109 = v4
											} else {
												v102 = int32(1)
												if l0 == int32(0) {
													v109 = v102
												} else {
													v105 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
													v107 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
													v109 = v102
												}
											}
											m.G0 = v12 + int32(48)
											return v109
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
func F_lseg_length(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v30 float64
	_ = v30
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v67 float64
	_ = v67
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v119 float64
	_ = v119
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) != 0 {
		v25 = v8
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
		v27 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
		v28 = base.F64_sub(v26, v27)
		v30 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v28), v30)|base.F64_eq(base.F64_abs(v26), v30)|base.F64_eq(base.F64_abs(v27), v30) != 0 {
			v43 = v28
			v52 = m.G0
			v54 = v52 - int32(32)
			m.G0 = v54
			v56 = base.F64_abs(v25)
			v57 = base.F64_abs(v43)
			v60 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)))
			if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
				v61 = v56
			} else {
				v61 = v57
			}
			v62 = base.I64_reinterpret_f64(v61)
			v64 = int64(base.Ui64(v62) >> (uint(int64(52)) % 64))
			if v64 == int64(2047) {
				v119 = v61
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
					v67 = v57
				} else {
					v67 = v56
				}
				if v62 == int64(0) {
					v119 = v67
				} else {
					v70 = base.I64_reinterpret_f64(v67)
					v72 = int64(base.Ui64(v70) >> (uint(int64(52)) % 64))
					if v72 == int64(2047) {
						v119 = v67
					} else {
						if int32(65) <= base.I32_wrap_i64(v72)-base.I32_wrap_i64(v64) {
							v119 = base.F64_add(v56, v57)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v70) {
								v83 = float64(1.90109156629516e-211)
								v96 = base.F64_mul(v67, v83)
								v97 = base.F64_mul(v61, v83)
								v98 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v62) {
									v96 = v67
									v97 = v61
									v98 = float64(1)
								} else {
									v91 = float64(5.260135901548374e+210)
									v96 = base.F64_mul(v67, v91)
									v97 = base.F64_mul(v61, v91)
									v98 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v54+int32(24), v54+int32(16), v96)
							mBase = m.M
							F_sq(m, v54+int32(8), v54, v97)
							mBase = m.M
							v107 = *(*float64)(unsafe.Add(mBase, uint32(v54)))
							v108 = *(*float64)(unsafe.Add(mBase, uint32(v54)+16))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v54)+8))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v54)+24))
							v119 = base.F64_mul(v98, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v107, v108), v110), v112)))
						}
					}
				}
			}
			m.G0 = v54 + int32(32)
			return base.I64_reinterpret_f64(v119)
		} else {
			v41 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				v43 = v41
				v52 = m.G0
				v54 = v52 - int32(32)
				m.G0 = v54
				v56 = base.F64_abs(v25)
				v57 = base.F64_abs(v43)
				v60 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)))
				if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
					v61 = v56
				} else {
					v61 = v57
				}
				v62 = base.I64_reinterpret_f64(v61)
				v64 = int64(base.Ui64(v62) >> (uint(int64(52)) % 64))
				if v64 == int64(2047) {
					v119 = v61
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
						v67 = v57
					} else {
						v67 = v56
					}
					if v62 == int64(0) {
						v119 = v67
					} else {
						v70 = base.I64_reinterpret_f64(v67)
						v72 = int64(base.Ui64(v70) >> (uint(int64(52)) % 64))
						if v72 == int64(2047) {
							v119 = v67
						} else {
							if int32(65) <= base.I32_wrap_i64(v72)-base.I32_wrap_i64(v64) {
								v119 = base.F64_add(v56, v57)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v70) {
									v83 = float64(1.90109156629516e-211)
									v96 = base.F64_mul(v67, v83)
									v97 = base.F64_mul(v61, v83)
									v98 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v62) {
										v96 = v67
										v97 = v61
										v98 = float64(1)
									} else {
										v91 = float64(5.260135901548374e+210)
										v96 = base.F64_mul(v67, v91)
										v97 = base.F64_mul(v61, v91)
										v98 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v54+int32(24), v54+int32(16), v96)
								mBase = m.M
								F_sq(m, v54+int32(8), v54, v97)
								mBase = m.M
								v107 = *(*float64)(unsafe.Add(mBase, uint32(v54)))
								v108 = *(*float64)(unsafe.Add(mBase, uint32(v54)+16))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v54)+8))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v54)+24))
								v119 = base.F64_mul(v98, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v107, v108), v110), v112)))
							}
						}
					}
				}
				m.G0 = v54 + int32(32)
				return base.I64_reinterpret_f64(v119)
			}
		}
	} else {
		v21 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = v21
			v26 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
			v28 = base.F64_sub(v26, v27)
			v30 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v28), v30)|base.F64_eq(base.F64_abs(v26), v30)|base.F64_eq(base.F64_abs(v27), v30) != 0 {
				v43 = v28
				v52 = m.G0
				v54 = v52 - int32(32)
				m.G0 = v54
				v56 = base.F64_abs(v25)
				v57 = base.F64_abs(v43)
				v60 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)))
				if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
					v61 = v56
				} else {
					v61 = v57
				}
				v62 = base.I64_reinterpret_f64(v61)
				v64 = int64(base.Ui64(v62) >> (uint(int64(52)) % 64))
				if v64 == int64(2047) {
					v119 = v61
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
						v67 = v57
					} else {
						v67 = v56
					}
					if v62 == int64(0) {
						v119 = v67
					} else {
						v70 = base.I64_reinterpret_f64(v67)
						v72 = int64(base.Ui64(v70) >> (uint(int64(52)) % 64))
						if v72 == int64(2047) {
							v119 = v67
						} else {
							if int32(65) <= base.I32_wrap_i64(v72)-base.I32_wrap_i64(v64) {
								v119 = base.F64_add(v56, v57)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v70) {
									v83 = float64(1.90109156629516e-211)
									v96 = base.F64_mul(v67, v83)
									v97 = base.F64_mul(v61, v83)
									v98 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v62) {
										v96 = v67
										v97 = v61
										v98 = float64(1)
									} else {
										v91 = float64(5.260135901548374e+210)
										v96 = base.F64_mul(v67, v91)
										v97 = base.F64_mul(v61, v91)
										v98 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v54+int32(24), v54+int32(16), v96)
								mBase = m.M
								F_sq(m, v54+int32(8), v54, v97)
								mBase = m.M
								v107 = *(*float64)(unsafe.Add(mBase, uint32(v54)))
								v108 = *(*float64)(unsafe.Add(mBase, uint32(v54)+16))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v54)+8))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v54)+24))
								v119 = base.F64_mul(v98, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v107, v108), v110), v112)))
							}
						}
					}
				}
				m.G0 = v54 + int32(32)
				return base.I64_reinterpret_f64(v119)
			} else {
				v41 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					v43 = v41
					v52 = m.G0
					v54 = v52 - int32(32)
					m.G0 = v54
					v56 = base.F64_abs(v25)
					v57 = base.F64_abs(v43)
					v60 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)))
					if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
						v61 = v56
					} else {
						v61 = v57
					}
					v62 = base.I64_reinterpret_f64(v61)
					v64 = int64(base.Ui64(v62) >> (uint(int64(52)) % 64))
					if v64 == int64(2047) {
						v119 = v61
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v56)) < base.Ui64(base.I64_reinterpret_f64(v57)) {
							v67 = v57
						} else {
							v67 = v56
						}
						if v62 == int64(0) {
							v119 = v67
						} else {
							v70 = base.I64_reinterpret_f64(v67)
							v72 = int64(base.Ui64(v70) >> (uint(int64(52)) % 64))
							if v72 == int64(2047) {
								v119 = v67
							} else {
								if int32(65) <= base.I32_wrap_i64(v72)-base.I32_wrap_i64(v64) {
									v119 = base.F64_add(v56, v57)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v70) {
										v83 = float64(1.90109156629516e-211)
										v96 = base.F64_mul(v67, v83)
										v97 = base.F64_mul(v61, v83)
										v98 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v62) {
											v96 = v67
											v97 = v61
											v98 = float64(1)
										} else {
											v91 = float64(5.260135901548374e+210)
											v96 = base.F64_mul(v67, v91)
											v97 = base.F64_mul(v61, v91)
											v98 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v54+int32(24), v54+int32(16), v96)
									mBase = m.M
									F_sq(m, v54+int32(8), v54, v97)
									mBase = m.M
									v107 = *(*float64)(unsafe.Add(mBase, uint32(v54)))
									v108 = *(*float64)(unsafe.Add(mBase, uint32(v54)+16))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v54)+8))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v54)+24))
									v119 = base.F64_mul(v98, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v107, v108), v110), v112)))
								}
							}
						}
					}
					m.G0 = v54 + int32(32)
					return base.I64_reinterpret_f64(v119)
				}
			}
		}
	}
}
func F_lseg_recv(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(32))
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
					v18 = F_pq_getmsgfloat8(m, v3)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v5)+24)) = v18
						return base.I64_extend_i32_u(v5)
					}
				}
			}
		}
	}
}
