package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_contain_point(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v106 float64
	_ = v106
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v9 = base.F64_mul(v7, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v7), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		v24 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v39 = v24
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v42 = base.F64_mul(v40, v41)
			v44 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v40), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
				v57 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v70 = v57
					v72 = math.Float64frombits(uint64(0x7ff0000000000000))
					v74 = base.F64_add(v39, v70)
					if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
						v86 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v88 = v86
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					} else {
						v88 = v74
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v91 = base.F64_abs(base.F64_add(v88, v89))
						if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v106 = v91
							return base.F64_le(v106, float64(1e-06))
						} else {
							v94 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
								v106 = v94
								return base.F64_le(v106, float64(1e-06))
							} else {
								v103 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v106 = base.F64_abs(v103)
									return base.F64_le(v106, float64(1e-06))
								}
							}
						}
					}
				}
			} else {
				v59 = float64(0)
				if base.F64_eq(v40, v59)|base.F64_ne(v42, v59)|base.F64_eq(v41, v59) != 0 {
					v70 = v42
					v72 = math.Float64frombits(uint64(0x7ff0000000000000))
					v74 = base.F64_add(v39, v70)
					if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
						v86 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v88 = v86
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					} else {
						v88 = v74
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v91 = base.F64_abs(base.F64_add(v88, v89))
						if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v106 = v91
							return base.F64_le(v106, float64(1e-06))
						} else {
							v94 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
								v106 = v94
								return base.F64_le(v106, float64(1e-06))
							} else {
								v103 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v106 = base.F64_abs(v103)
									return base.F64_le(v106, float64(1e-06))
								}
							}
						}
					}
				} else {
					v68 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v70 = v68
						v72 = math.Float64frombits(uint64(0x7ff0000000000000))
						v74 = base.F64_add(v39, v70)
						if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
							v86 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = v86
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								v91 = base.F64_abs(base.F64_add(v88, v89))
								if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v106 = v91
									return base.F64_le(v106, float64(1e-06))
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
										v106 = v94
										return base.F64_le(v106, float64(1e-06))
									} else {
										v103 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = base.F64_abs(v103)
											return base.F64_le(v106, float64(1e-06))
										}
									}
								}
							}
						} else {
							v88 = v74
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v28 = float64(0)
		if base.F64_eq(v7, v28)|base.F64_ne(v9, v28)|base.F64_eq(v8, v28) != 0 {
			v39 = v9
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v42 = base.F64_mul(v40, v41)
			v44 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v40), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
				v57 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v70 = v57
					v72 = math.Float64frombits(uint64(0x7ff0000000000000))
					v74 = base.F64_add(v39, v70)
					if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
						v86 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v88 = v86
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					} else {
						v88 = v74
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v91 = base.F64_abs(base.F64_add(v88, v89))
						if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v106 = v91
							return base.F64_le(v106, float64(1e-06))
						} else {
							v94 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
								v106 = v94
								return base.F64_le(v106, float64(1e-06))
							} else {
								v103 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v106 = base.F64_abs(v103)
									return base.F64_le(v106, float64(1e-06))
								}
							}
						}
					}
				}
			} else {
				v59 = float64(0)
				if base.F64_eq(v40, v59)|base.F64_ne(v42, v59)|base.F64_eq(v41, v59) != 0 {
					v70 = v42
					v72 = math.Float64frombits(uint64(0x7ff0000000000000))
					v74 = base.F64_add(v39, v70)
					if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
						v86 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v88 = v86
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					} else {
						v88 = v74
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v91 = base.F64_abs(base.F64_add(v88, v89))
						if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v106 = v91
							return base.F64_le(v106, float64(1e-06))
						} else {
							v94 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
								v106 = v94
								return base.F64_le(v106, float64(1e-06))
							} else {
								v103 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v106 = base.F64_abs(v103)
									return base.F64_le(v106, float64(1e-06))
								}
							}
						}
					}
				} else {
					v68 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v70 = v68
						v72 = math.Float64frombits(uint64(0x7ff0000000000000))
						v74 = base.F64_add(v39, v70)
						if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
							v86 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = v86
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								v91 = base.F64_abs(base.F64_add(v88, v89))
								if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v106 = v91
									return base.F64_le(v106, float64(1e-06))
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
										v106 = v94
										return base.F64_le(v106, float64(1e-06))
									} else {
										v103 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = base.F64_abs(v103)
											return base.F64_le(v106, float64(1e-06))
										}
									}
								}
							}
						} else {
							v88 = v74
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					}
				}
			}
		} else {
			v37 = F_float_underflow_error_ext(m, int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = v37
				v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				v42 = base.F64_mul(v40, v41)
				v44 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v40), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
					v57 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						v70 = v57
						v72 = math.Float64frombits(uint64(0x7ff0000000000000))
						v74 = base.F64_add(v39, v70)
						if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
							v86 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = v86
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								v91 = base.F64_abs(base.F64_add(v88, v89))
								if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v106 = v91
									return base.F64_le(v106, float64(1e-06))
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
										v106 = v94
										return base.F64_le(v106, float64(1e-06))
									} else {
										v103 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = base.F64_abs(v103)
											return base.F64_le(v106, float64(1e-06))
										}
									}
								}
							}
						} else {
							v88 = v74
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					}
				} else {
					v59 = float64(0)
					if base.F64_eq(v40, v59)|base.F64_ne(v42, v59)|base.F64_eq(v41, v59) != 0 {
						v70 = v42
						v72 = math.Float64frombits(uint64(0x7ff0000000000000))
						v74 = base.F64_add(v39, v70)
						if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
							v86 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								v88 = v86
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								v91 = base.F64_abs(base.F64_add(v88, v89))
								if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v106 = v91
									return base.F64_le(v106, float64(1e-06))
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
										v106 = v94
										return base.F64_le(v106, float64(1e-06))
									} else {
										v103 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = base.F64_abs(v103)
											return base.F64_le(v106, float64(1e-06))
										}
									}
								}
							}
						} else {
							v88 = v74
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
							v91 = base.F64_abs(base.F64_add(v88, v89))
							if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v106 = v91
								return base.F64_le(v106, float64(1e-06))
							} else {
								v94 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
									v106 = v94
									return base.F64_le(v106, float64(1e-06))
								} else {
									v103 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v106 = base.F64_abs(v103)
										return base.F64_le(v106, float64(1e-06))
									}
								}
							}
						}
					} else {
						v68 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							v70 = v68
							v72 = math.Float64frombits(uint64(0x7ff0000000000000))
							v74 = base.F64_add(v39, v70)
							if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
								v86 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v88 = v86
									v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
									v91 = base.F64_abs(base.F64_add(v88, v89))
									if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v106 = v91
										return base.F64_le(v106, float64(1e-06))
									} else {
										v94 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
											v106 = v94
											return base.F64_le(v106, float64(1e-06))
										} else {
											v103 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v106 = base.F64_abs(v103)
												return base.F64_le(v106, float64(1e-06))
											}
										}
									}
								}
							} else {
								v88 = v74
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
								v91 = base.F64_abs(base.F64_add(v88, v89))
								if base.F64_ne(v91, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v106 = v91
									return base.F64_le(v106, float64(1e-06))
								} else {
									v94 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v88), v94)|base.F64_eq(base.F64_abs(v89), v94) != 0 {
										v106 = v94
										return base.F64_le(v106, float64(1e-06))
									} else {
										v103 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v106 = base.F64_abs(v103)
											return base.F64_le(v106, float64(1e-06))
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
func F_line_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 float64
	_ = v176
	var v182 float64
	_ = v182
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 float64
	_ = v186
	var v189 int64
	_ = v189
	var v196 float64
	_ = v196
	var v203 int64
	_ = v203
	var v208 float64
	_ = v208
	var v226 int32
	_ = v226
	var v233 float64
	_ = v233
	var v237 int64
	_ = v237
	var v238 float64
	_ = v238
	var v241 int64
	_ = v241
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v285 float64
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v347 int32
	_ = v347
	var v361 int64
	_ = v361
	v12 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_palloc(m, int32(24))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v28 = v19
	goto L7
L3:
	;
	m.G0 = v16 + int32(48)
	return v361
L4:
	;
	v347 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v347)
	v361 = v12
	goto L3
L5:
	;
	v316 = F_errsave_start(m, v18)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L79
	}
L6:
	;
	v361 = base.I64_extend_i32_u(v21)
	goto L3
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(base.Ui32(v38-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v38 == int32(32)) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v170 = F_path_decode(m, v28, int32(1), int32(2), v16+int32(16), v16+int32(15), int32(0), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L46
	}
L9:
	;
	v28 = v28 + int32(1)
	goto L7
L10:
	;
	if v38 == int32(123) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	v51 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v51
	v56 = F_float8in_internal(m, v51, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = v56
	if v18 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v67 = v65 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 != int32(44) {
		goto L5
	} else {
		goto L20
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v61 != int32(453) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v64 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v75 = F_float8in_internal(m, v67, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+8)) = v75
	if v18 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v86 = v84 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != int32(44) {
		goto L5
	} else {
		goto L26
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v80 != int32(453) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v94 = F_float8in_internal(m, v86, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+16)) = v94
	if v18 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v105 = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v107 != int32(125) {
		goto L5
	} else {
		goto L32
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v99 != int32(453) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v102 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v113 = v105
	goto L33
L33:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if base.B2i32(base.Ui32(v123-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v123 == int32(32)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = v113 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v132
	v113 = v132
	goto L33
L36:
	;
	if v123 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	if base.F64_le(base.F64_abs(v134), float64(1e-06)) == int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	if base.F64_le(base.F64_abs(v140), float64(1e-06)) == int32(0) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v146 = F_errsave_start(m, v18)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v146 == int32(0) {
		v361 = v12
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_line_in_4), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(1044), int32(_a_F_line_in_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v361 = v12
	goto L3
L46:
	;
	if v170 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v174)
	v361 = v12
	goto L3
L48:
	;
	goto L49
L49:
	;
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.Ui64(base.I64_reinterpret_f64(v176)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v282 = v16 + int32(16)
	v285 = F_point_sl(m, v282, v16+int32(32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L77
	}
L51:
	;
	v257 = int64(0)
	v258 = F_errsave_start(m, v18)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L72
	}
L52:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v241 = base.I64_reinterpret_f64(v238) & int64(9223372036854775807)
	if base.Ui64(v237) <= base.Ui64(int64(9218868437227405312)) {
		goto L67
	} else {
		goto L68
	}
L53:
	;
	if base.B2i32(v226 == int32(0))|base.F64_ne(v176, v182) != 0 {
		goto L50
	} else {
		goto L66
	}
L54:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v176, v182)), float64(1e-06)) == int32(0))&base.F64_ne(v176, v182) != 0 {
		goto L50
	} else {
		goto L64
	}
L55:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v184 = int64(9223372036854775807)
	v185 = base.I64_reinterpret_f64(v182) & v184
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v189 = base.I64_reinterpret_f64(v186) & v184
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v189) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.Ui64(v203&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L50
	} else {
		goto L63
	}
L58:
	;
	v226 = base.B2i32(base.Ui64(v185) < base.Ui64(int64(9218868437227405313)))
	goto L53
L59:
	;
	goto L60
L60:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v185) {
		goto L50
	} else {
		goto L61
	}
L61:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v196)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v226 = int32(1)
	goto L53
L63:
	;
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v233 = v208
	v237 = base.I64_reinterpret_f64(v208) & int64(9223372036854775807)
	goto L52
L64:
	;
	if base.F64_eq(v186, v196)|base.F64_le(base.F64_abs(base.F64_sub(v186, v196)), float64(1e-06)) != 0 {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	goto L50
L66:
	;
	v233 = v186
	v237 = v189
	goto L52
L67:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v241))|base.F64_ne(v238, v233) != 0 {
		goto L50
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if base.Ui64(v241) < base.Ui64(int64(9218868437227405313)) {
		goto L50
	} else {
		goto L71
	}
L70:
	;
	goto L51
L71:
	;
	goto L51
L72:
	;
	if v258 == int32(0) {
		v361 = v257
		goto L3
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_line_in_6), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(1054), int32(_a_F_line_in_5))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v361 = v257
	goto L3
L77:
	;
	F_line_construct(m, v21, v282, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L6
L79:
	;
	if v316 == int32(0) {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_line_in_0)
	F_errmsg(m, int32(_a_F_line_in_1), v16)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(1021), int32(_a_F_line_in_3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L4
}
func F_line_parallel(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_line_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5 ^ int32(1))
	}
}
