package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14268(m, l0, int32(_a_F_gbt_bit_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_bool_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v10 != int64(0)))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v2)
	v20 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16 + v20
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v16
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v31)+12)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_gbt_num_consistent(m, v7+int32(4), v7+int32(15), v7+int32(12), v33&v20, int32(_a_F_gbt_bool_consistent_0), v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v38)
	}
}
func F_gbt_bool_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_bool_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_boolgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(base.Ui32(v5) < base.Ui32(v4))
}
func F_gbt_bpchar_consistent(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14283(m, l0, int32(_a_F_gbt_bpchar_consistent_0), int32(_a_F_gbt_bpchar_consistent_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_bpchar_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14284(m, l0, l1, l2, int32(2620))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_bytea_penalty(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14267(m, l0, int32(_a_F_gbt_bytea_penalty_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_bytea_union(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14271(m, l0, int32(_a_F_gbt_bytea_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_cash_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_cash_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_cashkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if v7 == v9 {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
		if v12 == v13 {
			v26 = int32(0)
			return v26
		} else {
			if v13 < v12 {
				v18 = int32(1)
			} else {
				v18 = int32(-1)
			}
			return v18
		}
	} else {
		if v9 < v7 {
			v23 = int32(1)
		} else {
			v23 = int32(-1)
		}
		v26 = v23
		return v26
	}
}
func F_gbt_date_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14278(m, l0, int32(_a_F_gbt_date_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_date_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_date_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_datele(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v6 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0))))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1))))
	v8 = F_DirectFunctionCall2Coll(m, int32(2627), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_enum_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 float32
	_ = v49
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_convert_i32_u(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = base.F64_convert_i32_u(v16)
	if base.F64_lt(v17, v13) != 0 {
		v21 = base.F64_sub(v13, v17)
	} else {
		v21 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v23 = base.F64_convert_i32_u(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v25 = base.F64_convert_i32_u(v24)
	if base.F64_lt(v25, v23) != 0 {
		v29 = base.F64_sub(v23, v25)
	} else {
		v29 = float64(0)
	}
	v30 = base.F64_add(v21, v29)
	if base.F64_gt(v30, float64(0)) != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		v49 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v30, base.F64_add(base.F64_sub(v25, v13), v30))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v42+int32(1))))
	} else {
		v49 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = v49
	return v8
}
func F_gbt_enum_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_enum_sortsupport_0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v9 = F_MemoryContextAlloc(m, v7, int32(28))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_fmgr_info_cxt(m, int32(3514), v9, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v9
			return int64(0)
		}
	}
}
func F_gbt_enum_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_enum_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_enumeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 == v5)
}
func F_gbt_float4_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_float4_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_float4_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14278(m, l0, int32(_a_F_gbt_float4_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_float4_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_float4_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_float4_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v7 float32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	v5 = *(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	v12 = int32(2147483647)
	v13 = base.I32_reinterpret_f32(v5) & v12
	v16 = base.I32_reinterpret_f32(v7) & v12
	if base.Ui32(int32(2139095041)) <= base.Ui32(v16) {
		v26 = base.B2i32(base.Ui32(v13) < base.Ui32(int32(2139095041)))
		v35 = int32(0) - v26&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v16))|base.F32_lt(v5, v7))
	} else {
		v21 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_gt(v5, v7) != 0 {
			v35 = v21
		} else {
			v26 = v21
			v35 = int32(0) - v26&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v16))|base.F32_lt(v5, v7))
		}
	}
	return v35
}
func F_gbt_float8_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v19 int32
	_ = v19
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v48 float64
	_ = v48
	var v54 int32
	_ = v54
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_sub(v9, v10)
	v12 = base.F64_abs(v11)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	v19 = int32(0)
	if base.B2i32(base.F64_ne(v12, v13)|base.F64_eq(base.F64_abs(v9), v13) == v19)&base.F64_ne(base.F64_abs(v10), v13) == v19 {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)) {
			v31 = int64(9223372036854775807)
			v32 = base.I64_reinterpret_f64(v10) & v31
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v9)&v31) {
				if base.Ui64(v32) <= base.Ui64(int64(9218868437227405312)) {
					v48 = math.Float64frombits(uint64(0x7ff0000000000000))
					return base.F64_abs(v48)
				} else {
					return float64(0)
				}
			} else {
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v32) {
					v48 = math.Float64frombits(uint64(0x7ff0000000000000))
					return base.F64_abs(v48)
				} else {
					return float64(0)
				}
			}
		} else {
			v48 = v11
			return base.F64_abs(v48)
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return float64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_gbt_float8_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_float8_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_float8_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	v5 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	v12 = int64(9223372036854775807)
	v13 = base.I64_reinterpret_f64(v5) & v12
	v16 = base.I64_reinterpret_f64(v7) & v12
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v16) {
		v26 = base.B2i32(base.Ui64(v13) < base.Ui64(int64(9218868437227405313)))
		v35 = int32(0) - v26&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v16))|base.F64_lt(v5, v7))
	} else {
		v21 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_gt(v5, v7) != 0 {
			v35 = v21
		} else {
			v26 = v21
			v35 = int32(0) - v26&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v16))|base.F64_lt(v5, v7))
		}
	}
	return v35
}
func F_gbt_inet_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+18)))
	if v11 != int32(1) {
		v42 = v10
		m.G0 = v8 + int32(16)
		return base.I64_extend_i32_u(v42)
	} else {
		v15 = F_palloc(m, int32(16))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v19)
			v22 = F_palloc(m, int32(24))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				v28 = F_convert_network_to_scalar(m, v24, int32(869), v8+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v28
					*(*float64)(unsafe.Add(mBase, uint32(v15))) = v28
					*(*int64)(unsafe.Add(mBase, uint32(v22))) = base.I64_extend_i32_u(v15)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v34
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v36
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)) = uint8(v39)
					*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)) = uint16(v38)
					v42 = v22
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v42)
				}
			}
		}
	}
}
func F_gbt_inet_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_inet_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_inetgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_gt(v4, v5)
}
func F_gbt_inetle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_le(v4, v5)
}
func F_gbt_int2_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_int2_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_int2eq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v4 == v5)
}
func F_gbt_int2lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v4 < v5)
}
func F_gbt_int4_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 float32
	_ = v49
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_convert_i32_s(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = base.F64_convert_i32_s(v16)
	if base.F64_lt(v17, v13) != 0 {
		v21 = base.F64_sub(v13, v17)
	} else {
		v21 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v23 = base.F64_convert_i32_s(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v25 = base.F64_convert_i32_s(v24)
	if base.F64_lt(v25, v23) != 0 {
		v29 = base.F64_sub(v23, v25)
	} else {
		v29 = float64(0)
	}
	v30 = base.F64_add(v21, v29)
	if base.F64_gt(v30, float64(0)) != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		v49 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v30, base.F64_add(base.F64_sub(v25, v13), v30))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v42+int32(1))))
	} else {
		v49 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = v49
	return v8
}
func F_gbt_int8_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14277(m, l0, int32(_a_F_gbt_int8_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_int8_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_int8_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macad8_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_macad8_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_macad8_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_macad8_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_macad8gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8gt_0), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_macad_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_macad_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_macad_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macaddr8_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v6 = int64(4294967295)
	v10 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macaddr8_ssup_cmp_0), int32(0), l0&v6, l1&v6)
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
	}
}
func F_gbt_num_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v7 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	switch v8 - int32(1) {
	case 0:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l3 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
			v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, l1, v18, l5)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, l1, v18, l5)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	case 1:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
		v13 = m.T0[v12].(func(*base.Module, int32, int32, int32) int32)(m, l1, v11, l5)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	case 2:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l3 != 0 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, l1, v27, l5)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
			v33 = m.T0[v32].(func(*base.Module, int32, int32, int32) int32)(m, v27, l1, l5)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v33 == int32(0) {
					v63 = v7
					return v63
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, l1, v66, l5)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						return v68
					}
				}
			}
		}
	case 3:
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
		v68 = m.T0[v67].(func(*base.Module, int32, int32, int32) int32)(m, l1, v66, l5)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			return v68
		}
	case 4:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l3 != 0 {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
			v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, l1, v37, l5)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
			v43 = m.T0[v42].(func(*base.Module, int32, int32, int32) int32)(m, l1, v37, l5)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				return v43
			}
		}
	case 5:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
		v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, l1, v46, l5)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			if l3 != 0 {
				return v48 ^ int32(1)
			} else {
				if v48 == int32(0) {
					v63 = int32(1)
					return v63
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
					v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, l1, v56, l5)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v63 = v58 ^ int32(1)
						return v63
					}
				}
			}
		}
	default:
		v63 = v7
		return v63
	}
}
func F_gbt_numeric_gt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2934), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_numeric_le(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2935), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_numeric_union(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14271(m, l0, int32(_a_F_gbt_numeric_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_oid_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_oid_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_oidle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v4) <= base.Ui32(v5))
}
func F_gbt_text_union(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14271(m, l0, int32(_a_F_gbt_text_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_time_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_time_sortsupport_0)
	return int64(0)
}
func F_gbt_tsge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_DirectFunctionCall2Coll(m, int32(2652), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_tskey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14285(m, l0, l1, l2, int32(1578))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_tslt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_DirectFunctionCall2Coll(m, int32(1711), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_uuidge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v77 = v42<<(uint(v7)%64) | v42&v9<<(uint(v11)%64) | (v42&v14<<(uint(v16)%64) | v42&v18<<(uint(v20)%64)) | (int64(base.Ui64(v42)>>(uint(v20)%64))&v18 | int64(base.Ui64(v42)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v42)>>(uint(v11)%64))&v9 | int64(base.Ui64(v42)>>(uint(v7)%64))))
	if v41 == v77 {
		v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v81 = int64(56)
		v83 = int64(65280)
		v85 = int64(40)
		v88 = int64(16711680)
		v90 = int64(24)
		v92 = int64(4278190080)
		v94 = int64(8)
		v115 = v80<<(uint(v81)%64) | v80&v83<<(uint(v85)%64) | (v80&v88<<(uint(v90)%64) | v80&v92<<(uint(v94)%64)) | (int64(base.Ui64(v80)>>(uint(v94)%64))&v92 | int64(base.Ui64(v80)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v80)>>(uint(v85)%64))&v83 | int64(base.Ui64(v80)>>(uint(v81)%64))))
		v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v151 = v116<<(uint(v81)%64) | v116&v83<<(uint(v85)%64) | (v116&v88<<(uint(v90)%64) | v116&v92<<(uint(v94)%64)) | (int64(base.Ui64(v116)>>(uint(v94)%64))&v92 | int64(base.Ui64(v116)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v116)>>(uint(v85)%64))&v83 | int64(base.Ui64(v116)>>(uint(v81)%64))))
		if v115 == v151 {
			v161 = int32(0)
		} else {
			v153 = v151
			v154 = v115
			if base.Ui64(v154) < base.Ui64(v153) {
				v158 = int32(-1)
			} else {
				v158 = int32(1)
			}
			v161 = v158
		}
	} else {
		v153 = v77
		v154 = v41
		if base.Ui64(v154) < base.Ui64(v153) {
			v158 = int32(-1)
		} else {
			v158 = int32(1)
		}
		v161 = v158
	}
	return int32(base.Ui32(v161^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_gbt_uuidlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v77 = v42<<(uint(v7)%64) | v42&v9<<(uint(v11)%64) | (v42&v14<<(uint(v16)%64) | v42&v18<<(uint(v20)%64)) | (int64(base.Ui64(v42)>>(uint(v20)%64))&v18 | int64(base.Ui64(v42)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v42)>>(uint(v11)%64))&v9 | int64(base.Ui64(v42)>>(uint(v7)%64))))
	if v41 == v77 {
		v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v81 = int64(56)
		v83 = int64(65280)
		v85 = int64(40)
		v88 = int64(16711680)
		v90 = int64(24)
		v92 = int64(4278190080)
		v94 = int64(8)
		v115 = v80<<(uint(v81)%64) | v80&v83<<(uint(v85)%64) | (v80&v88<<(uint(v90)%64) | v80&v92<<(uint(v94)%64)) | (int64(base.Ui64(v80)>>(uint(v94)%64))&v92 | int64(base.Ui64(v80)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v80)>>(uint(v85)%64))&v83 | int64(base.Ui64(v80)>>(uint(v81)%64))))
		v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		v151 = v116<<(uint(v81)%64) | v116&v83<<(uint(v85)%64) | (v116&v88<<(uint(v90)%64) | v116&v92<<(uint(v94)%64)) | (int64(base.Ui64(v116)>>(uint(v94)%64))&v92 | int64(base.Ui64(v116)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v116)>>(uint(v85)%64))&v83 | int64(base.Ui64(v116)>>(uint(v81)%64))))
		if v115 == v151 {
			v161 = int32(0)
		} else {
			v153 = v151
			v154 = v115
			if base.Ui64(v154) < base.Ui64(v153) {
				v158 = int32(-1)
			} else {
				v158 = int32(1)
			}
			v161 = v158
		}
	} else {
		v153 = v77
		v154 = v41
		if base.Ui64(v154) < base.Ui64(v153) {
			v158 = int32(-1)
		} else {
			v158 = int32(1)
		}
		v161 = v158
	}
	return int32(base.Ui32(v161) >> (uint(int32(31)) % 32))
}
func F_gbt_var_key_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = int32(2)
	v11 = int32(base.Ui32(v9) >> (uint(v10) % 32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v15 = int32(base.Ui32(v13) >> (uint(v10) % 32))
	v19 = (v15 + int32(3)) & int32(2147483644)
	v22 = v11 + v19 + int32(4)
	v23 = F_palloc0(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v28 = v23 + int32(4)
		if v15 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			base.MemoryCopy(m, v28, v29, v15)
		} else {
		}
		if v11 != 0 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			base.MemoryCopy(m, v28+v19, v32, v11)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = v22 << (uint(int32(2)) % 32)
		return v23
	}
}
func F_gbt_var_penalty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v355 float64
	_ = v355
	var v356 float32
	_ = v356
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v25 = int32(4)
	v26 = v22 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(2)
	v29 = int32(base.Ui32(v27) >> (uint(v28) % 32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui32(v29+v25) < base.Ui32(int32(base.Ui32(v37)>>(uint(v28)%32))) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v71 = int32(4)
	v72 = v21 + v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = int32(2)
	v75 = int32(base.Ui32(v73) >> (uint(v74) % 32))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if base.Ui32(v75+v71) < base.Ui32(int32(base.Ui32(v83)>>(uint(v74)%32))) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v41 = v26 + (v29+int32(3))&int32(2147483644)
	goto L4
L3:
	;
	v41 = v26
	goto L4
L4:
	;
	if v26 != v41 {
		v68 = v26
		v69 = v41
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l4)+36))
	if v43 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v68 = v26
	v69 = v26
	goto L1
L7:
	;
	goto L8
L8:
	;
	v46 = m.T0[v43].(func(*base.Module, int32, int32) int32)(m, v22, l5)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v46 == v22 {
		v68 = v26
		v69 = v26
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v51 = int32(4)
	v52 = v46 + v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(2)
	v55 = int32(base.Ui32(v53) >> (uint(v54) % 32))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if base.Ui32(v55+v51) < base.Ui32(int32(base.Ui32(v63)>>(uint(v54)%32))) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v67 = v52 + (v55+int32(3))&int32(2147483644)
	goto L14
L13:
	;
	v67 = v52
	goto L14
L14:
	;
	v68 = v52
	v69 = v67
	goto L1
L15:
	;
	v87 = v72 + (v75+int32(3))&int32(2147483644)
	goto L17
L16:
	;
	v87 = v72
	goto L17
L17:
	;
	if v75 != int32(4) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	m.G0 = v19 + int32(16)
	return l0
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v98 = m.T0[v97].(func(*base.Module, int32, int32, int32, int32) int32)(m, v68, v72, l3, l5)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L23
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v90&int32(-4) != int32(16) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	goto L18
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
	v265 = v19 + int32(8)
	F_gbt_var_bin_union(m, v265, v21, l3, l4, l5)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L69
	}
L23:
	;
	if v98 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v103 = int32(2)
	v104 = int32(base.Ui32(v102) >> (uint(v103) % 32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if base.Ui32(int32(base.Ui32(v105)>>(uint(v103)%32))) < base.Ui32(v104) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v179 = m.T0[v178].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v87, l3, l5)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L47
	}
L27:
	;
	v109 = int32(4)
	v110 = v68 + v109
	v112 = v21 + int32(8)
	v114 = v104 - v109
	if base.Ui32(v109) <= base.Ui32(v114) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v176 != 0 {
		goto L22
	} else {
		goto L46
	}
L29:
	;
	v176 = int32(0)
	goto L28
L30:
	;
	v150 = v145
	v151 = v146
	v152 = v147
	goto L40
L31:
	;
	if (v110|v112)&int32(3) != 0 {
		v145 = v110
		v146 = v112
		v147 = v114
		goto L30
	} else {
		goto L34
	}
L32:
	;
	v138 = v110
	v139 = v112
	v140 = v114
	goto L33
L33:
	;
	if v140 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L34:
	;
	v122 = v110
	v123 = v112
	v124 = v114
	goto L35
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v127 != v128 {
		v145 = v122
		v146 = v123
		v147 = v124
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v138 = v133
	v139 = v131
	v140 = v135
	goto L33
L37:
	;
	v130 = int32(4)
	v131 = v123 + v130
	v133 = v122 + v130
	v135 = v124 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v122 = v133
		v123 = v131
		v124 = v135
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v145 = v138
	v146 = v139
	v147 = v140
	goto L30
L40:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 == v156 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v176 = v155 - v156
	goto L28
L42:
	;
	v158 = int32(1)
	v163 = v152 - v158
	if v163 != 0 {
		v150 = v150 + v158
		v151 = v151 + v158
		v152 = v163
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	goto L29
L46:
	;
	goto L26
L47:
	;
	if v179 <= int32(0) {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v184 = int32(2)
	v185 = int32(base.Ui32(v183) >> (uint(v184) % 32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if base.Ui32(int32(base.Ui32(v186)>>(uint(v184)%32))) < base.Ui32(v185) {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v190 = int32(4)
	v191 = v69 + v190
	v193 = v87 + v190
	v195 = v185 - v190
	if base.Ui32(v190) <= base.Ui32(v195) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v257 == int32(0) {
		goto L18
	} else {
		goto L68
	}
L51:
	;
	v257 = int32(0)
	goto L50
L52:
	;
	v231 = v226
	v232 = v227
	v233 = v228
	goto L62
L53:
	;
	if (v191|v193)&int32(3) != 0 {
		v226 = v191
		v227 = v193
		v228 = v195
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v219 = v191
	v220 = v193
	v221 = v195
	goto L55
L55:
	;
	if v221 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v203 = v191
	v204 = v193
	v205 = v195
	goto L57
L57:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v208 != v209 {
		v226 = v203
		v227 = v204
		v228 = v205
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v219 = v214
	v220 = v212
	v221 = v216
	goto L55
L59:
	;
	v211 = int32(4)
	v212 = v204 + v211
	v214 = v203 + v211
	v216 = v205 - v211
	if base.Ui32(int32(3)) < base.Ui32(v216) {
		v203 = v214
		v204 = v212
		v205 = v216
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v226 = v219
	v227 = v220
	v228 = v221
	goto L52
L62:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v236 == v237 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v257 = v236 - v237
	goto L50
L64:
	;
	v239 = int32(1)
	v244 = v233 - v239
	if v244 != 0 {
		v231 = v231 + v239
		v232 = v232 + v239
		v233 = v244
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L51
L68:
	;
	goto L22
L69:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v269 = F_gbt_var_node_cp_len(m, v268, l4)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	F_gbt_var_bin_union(m, v265, v22, l3, l4, l5)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v274 = F_gbt_var_node_cp_len(m, v273, l4)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L73
	}
L72:
	;
	v356 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v359 = int32(1)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+52))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = base.F32_mul(base.F32_add(base.F32_add(v356, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(v355, base.F64_convert_i32_s(v269+v359)))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v368+v359)))
	goto L18
L73:
	;
	if v274 < v269 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v355 = base.F64_convert_i32_s(v269 - v274)
	goto L72
L75:
	;
	goto L76
L76:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v280 = int32(2)
	v282 = int32(4)
	v283 = v273 + v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v286 = int32(base.Ui32(v284) >> (uint(v280) % 32))
	v289 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if base.Ui32(v274) < base.Ui32(int32(base.Ui32(v291)>>(uint(v280)%32))-v282) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v21)+8)))
	v299 = v298
	goto L79
L78:
	;
	v299 = v289
	goto L79
L79:
	;
	if base.Ui32(v274) < base.Ui32(v286-int32(4)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v273)+8)))
	v305 = v304
	goto L82
L81:
	;
	v305 = v289
	goto L82
L82:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if base.Ui32(v274) < base.Ui32(int32(base.Ui32(v307)>>(uint(int32(2))%32))-int32(4)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v87)+4)))
	v315 = v314
	goto L85
L84:
	;
	v315 = int32(0)
	goto L85
L85:
	;
	if base.Ui32(v286+v282) < base.Ui32(int32(base.Ui32(v279)>>(uint(v280)%32))) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v322 = (v286+int32(3))&int32(2147483644) + v283
	goto L88
L87:
	;
	v322 = v283
	goto L88
L88:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if base.Ui32(v274) < base.Ui32(int32(base.Ui32(v323)>>(uint(int32(2))%32))-int32(4)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v322)+4)))
	v332 = v330
	goto L91
L90:
	;
	v332 = int32(0)
	goto L91
L91:
	;
	v333 = v332 - v315
	v334 = int32(31)
	v335 = v333 >> (uint(v334) % 32)
	v338 = v299 - v305
	v340 = v338 >> (uint(v334) % 32)
	v355 = base.F64_mul(base.F64_convert_i32_u(v333^v335-v335+(v338^v340-v340)), float64(0.00390625))
	goto L72
}
