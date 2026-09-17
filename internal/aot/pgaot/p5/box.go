package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_above(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v89 int32
	_ = v89
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_add(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_add(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
				v51 = base.F64_add(v49, v50)
				v53 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.B2i32(base.F64_ne(base.F64_abs(v51), v53)|base.F64_eq(base.F64_abs(v49), v53) == int32(0))&base.F64_ne(base.F64_abs(v50), v53) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v65 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
					v66 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v67 = base.F64_add(v65, v66)
					v69 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.B2i32(base.F64_ne(base.F64_abs(v67), v69)|base.F64_eq(base.F64_abs(v65), v69) == int32(0))&base.F64_ne(base.F64_abs(v66), v69) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v67
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v51
						return v11
					}
				}
			}
		}
	}
}
func F_box_below(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == v2 {
		v32 = v2
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v32 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v32 = v2
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
				v32 = base.F64_le(v27, base.F64_add(v28, float64(1e-06)))
			}
		}
	}
	return v32
}
func F_box_gt(m *base.Module, l0 int32) int32 {
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
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_box_ar(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return base.F64_gt(v5, base.F64_add(v9, float64(1e-06)))
		}
	}
}
func F_box_height(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v26 = F_Float8GetDatum(m, v8)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			return v26
		}
	}
}
func F_box_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v38 float64
	_ = v38
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v69 int32
	_ = v69
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_palloc(m, int32(32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v25 = F_path_decode(m, v13, int32(0), int32(2), v17, v10+int32(15), int32(0), int32(_a_F_box_in_0), v13, v12)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v25 == int32(0) {
				v29 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
				v69 = int32(0)
			} else {
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)) {
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
					if base.B2i32(base.F64_lt(v32, v38) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v32
						*(*float64)(unsafe.Add(mBase, uint32(v17))) = v38
					}
				}
				v51 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v51)&int64(9223372036854775807)) {
					v69 = v17
				} else {
					v57 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
					if base.B2i32(base.F64_lt(v51, v57) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
						v69 = v17
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v17)+24)) = v51
						*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v57
						v69 = v17
					}
				}
			}
			m.G0 = v10 + int32(16)
			return v69
		}
	}
}
func F_box_intersect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	if base.F64_le(v7, base.F64_add(v9, float64(1e-06))) == int32(0) {
		v34 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		if base.F64_le(v15, base.F64_add(v16, float64(1e-06))) == int32(0) {
			v34 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
			return int32(0)
		} else {
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			if base.F64_le(v22, base.F64_add(v23, float64(1e-06))) == int32(0) {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
				v30 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
				if base.F64_le(v29, base.F64_add(v30, float64(1e-06))) != 0 {
					v39 = F_palloc(m, int32(32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
						if base.Ui64(base.I64_reinterpret_f64(v44)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807)) {
								v55 = v44
							} else {
								v55 = v43
							}
							if base.F64_gt(v43, v44) != 0 {
								v57 = v44
							} else {
								v57 = v55
							}
							v58 = v57
						} else {
							v58 = v43
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39))) = v58
						v60 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
						if base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v66 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
								v72 = v66
							} else {
								v72 = v60
							}
							if base.F64_lt(v60, v66) != 0 {
								v74 = v66
							} else {
								v74 = v72
							}
							v76 = v74
						} else {
							v76 = v60
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+16)) = v76
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
						if base.Ui64(base.I64_reinterpret_f64(v79)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v78)&int64(9223372036854775807)) {
								v90 = v79
							} else {
								v90 = v78
							}
							if base.F64_gt(v78, v79) != 0 {
								v92 = v79
							} else {
								v92 = v90
							}
							v93 = v92
						} else {
							v93 = v78
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+8)) = v93
						v95 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						if base.Ui64(base.I64_reinterpret_f64(v95)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v101 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) {
								v107 = v101
							} else {
								v107 = v95
							}
							if base.F64_lt(v95, v101) != 0 {
								v109 = v101
							} else {
								v109 = v107
							}
							v111 = v109
						} else {
							v111 = v95
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+24)) = v111
						return v39
					}
				} else {
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					return int32(0)
				}
			}
		}
	}
}
func F_box_overleft(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
