package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_dist_cpoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_point_dt(m, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v12 = base.F64_sub(v7, v11)
		if base.F64_ne(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v24 = float64(0)
			if base.F64_lt(v12, v24) != 0 {
				v27 = v24
			} else {
				v27 = v12
			}
			v28 = F_Float8GetDatum(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			if base.F64_eq(base.F64_abs(v7), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v24 = float64(0)
				if base.F64_lt(v12, v24) != 0 {
					v27 = v24
				} else {
					v27 = v12
				}
				v28 = F_Float8GetDatum(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v28
				}
			} else {
				if base.F64_eq(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v24 = float64(0)
					if base.F64_lt(v12, v24) != 0 {
						v27 = v24
					} else {
						v27 = v12
					}
					v28 = F_Float8GetDatum(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v28
					}
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
func F_dist_pc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_point_dt(m, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v12 = base.F64_sub(v7, v11)
		if base.F64_ne(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v24 = float64(0)
			if base.F64_lt(v12, v24) != 0 {
				v27 = v24
			} else {
				v27 = v12
			}
			v28 = F_Float8GetDatum(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			if base.F64_eq(base.F64_abs(v7), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v24 = float64(0)
				if base.F64_lt(v12, v24) != 0 {
					v27 = v24
				} else {
					v27 = v12
				}
				v28 = F_Float8GetDatum(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v28
				}
			} else {
				if base.F64_eq(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v24 = float64(0)
					if base.F64_lt(v12, v24) != 0 {
						v27 = v24
					} else {
						v27 = v12
					}
					v28 = F_Float8GetDatum(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v28
					}
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
func F_dist_polyc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_dist_ppoly_internal(m, v11, v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
			v15 = base.F64_sub(v12, v14)
			if base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v27 = float64(0)
				if base.F64_lt(v15, v27) != 0 {
					v30 = v27
				} else {
					v30 = v15
				}
				v31 = F_Float8GetDatum(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return v31
				}
			} else {
				if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v27 = float64(0)
					if base.F64_lt(v15, v27) != 0 {
						v30 = v27
					} else {
						v30 = v15
					}
					v31 = F_Float8GetDatum(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						return v31
					}
				} else {
					if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v27 = float64(0)
						if base.F64_lt(v15, v27) != 0 {
							v30 = v27
						} else {
							v30 = v15
						}
						v31 = F_Float8GetDatum(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							return v31
						}
					} else {
						F_float_overflow_error(m)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
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
func F_dist_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_dist_sl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v21 float64
	_ = v21
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_lseg_interpt_line(m, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v24 = float64(0)
			v25 = F_Float8GetDatum(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v25
			}
		} else {
			v13 = F_line_closept_point(m, int32(0), v7, v6)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v18 = F_line_closept_point(m, int32(0), v7, v6+int32(16))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(v13, v18) != 0 {
						v21 = v13
					} else {
						v21 = v18
					}
					v24 = v21
					v25 = F_Float8GetDatum(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v25
					}
				}
			}
		}
	}
}
