package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_picksplit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13882(m, l0, int32(_a_F_gbt_bit_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_bool_consistent(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v14 int32
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v10 != v2))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v2)
	v20 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16 + v20
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v16
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v31)+12)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_gbt_num_consistent(m, v7+int32(4), v7+int32(15), v7+int32(12), v33&v20, int32(_a_F_gbt_bool_consistent_0), v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v38
	}
}
func F_gbt_bool_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_bool_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
func F_gbt_bpchar_consistent(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13895(m, l0, int32(_a_F_gbt_bpchar_consistent_0), int32(_a_F_gbt_bpchar_consistent_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_bpchar_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13896(m, l0, l1, l2, int32(2407))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_bytea_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_penalty(m, v2, v3, v4, v5, int32(_a_F_gbt_bytea_penalty_0), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_gbt_bytea_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_gbt_var_union(m, v2, v3, v4, int32(_a_F_gbt_bytea_union_0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_cash_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_cash_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
func F_gbt_date_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13890(m, l0, int32(_a_F_gbt_date_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_date_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_date_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_datele(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_DirectFunctionCall2Coll(m, int32(2414), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_enum_penalty(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 float64
	_ = v39
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v16) < base.Ui32(v15) {
		v27 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v15), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_u(v16), float64(-0.49000000953674316))), float64(0))
	} else {
		v27 = float64(0)
	}
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui32(v29) < base.Ui32(v28) {
		v39 = base.F64_add(v27, base.F64_add(base.F64_mul(base.F64_convert_i32_u(v28), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_u(v29), float64(-0.49000000953674316))))
	} else {
		v39 = v27
	}
	if base.F64_gt(v39, float64(0)) != 0 {
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v39, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v16), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_u(v28), float64(-0.49000000953674316))), v39))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v57+int32(1))))
	} else {
	}
	return v12
}
func F_gbt_enum_sortsupport(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_enum_sortsupport_0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v9 = F_MemoryContextAlloc(m, v7, int32(28))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_fmgr_info_cxt(m, int32(3514), v9, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v9
			return int32(0)
		}
	}
}
func F_gbt_enum_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13886(m, l0, int32(_a_F_gbt_enum_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
func F_gbt_float4_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_float4_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_float4_consistent(m *base.Module, l0 int32) int32 {
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
	var v10 float32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*float32)(unsafe.Add(mBase, uint32(v7)+12)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+10)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+v27)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v7, v7+int32(12), v7+int32(10), v29&int32(1), int32(_a_F_gbt_float4_consistent_0), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v34
	}
}
func F_gbt_float4_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_float4_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_float4_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 float32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(2147483647)
	v9 = base.I32_reinterpret_f32(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	v13 = base.I32_reinterpret_f32(v10) & v8
	if base.Ui32(int32(2139095041)) <= base.Ui32(v13) {
		v23 = base.B2i32(base.Ui32(v9) < base.Ui32(int32(2139095041)))
		v32 = int32(0) - v23&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))
	} else {
		v18 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v9))|base.F32_gt(v6, v10) != 0 {
			v32 = v18
		} else {
			v23 = v18
			v32 = int32(0) - v23&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))
		}
	}
	return v32
}
func F_gbt_float8_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v10 = base.F64_abs(base.F64_sub(v7, v8))
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(v10, v11)|base.F64_eq(base.F64_abs(v7), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return float64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return v10
	}
}
func F_gbt_float8_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_float8_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_float8_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v10 = int64(9223372036854775807)
	v11 = base.I64_reinterpret_f64(v4) & v10
	v14 = base.I64_reinterpret_f64(v5) & v10
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v14) {
		v24 = base.B2i32(base.Ui64(v11) < base.Ui64(int64(9218868437227405313)))
		v33 = int32(0) - v24&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v14))|base.F64_lt(v4, v5))
	} else {
		v19 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_gt(v4, v5) != 0 {
			v33 = v19
		} else {
			v24 = v19
			v33 = int32(0) - v24&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v14))|base.F64_lt(v4, v5))
		}
	}
	return v33
}
func F_gbt_inet_compress(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
	if v11 != int32(1) {
		v41 = v10
		m.G0 = v8 + int32(16)
		return v41
	} else {
		v15 = F_palloc(m, int32(16))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v19)
			v22 = F_palloc(m, int32(16))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v28 = F_convert_network_to_scalar(m, v24, int32(869), v8+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v28
					*(*float64)(unsafe.Add(mBase, uint32(v15))) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v15
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v35
					v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+14)) = uint8(v38)
					*(*uint16)(unsafe.Add(mBase, uint32(v22)+12)) = uint16(v37)
					v41 = v22
					m.G0 = v8 + int32(16)
					return v41
				}
			}
		}
	}
}
func F_gbt_inet_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_inet_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_int2_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_int2_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
func F_gbt_int4_penalty(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 float64
	_ = v39
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v16 < v15 {
		v27 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v15), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_s(v16), float64(-0.49000000953674316))), float64(0))
	} else {
		v27 = float64(0)
	}
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v29 < v28 {
		v39 = base.F64_add(v27, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v28), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_s(v29), float64(-0.49000000953674316))))
	} else {
		v39 = v27
	}
	if base.F64_gt(v39, float64(0)) != 0 {
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v39, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v16), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_s(v28), float64(-0.49000000953674316))), v39))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v57+int32(1))))
	} else {
	}
	return v12
}
func F_gbt_int8_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13889(m, l0, int32(_a_F_gbt_int8_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_int8_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_int8_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macad8_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_macad8_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macad8_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_macad8_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macad8gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8gt_0), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_macad_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macad_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_macad_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_macaddr8_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macaddr8_ssup_cmp_0), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
					v60 = v7
					return v60
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, l1, v63, l5)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						return v65
					}
				}
			}
		}
	case 3:
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
		v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, l1, v63, l5)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			return v65
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
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
		v49 = m.T0[v48].(func(*base.Module, int32, int32, int32) int32)(m, l1, v47, l5)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			if v49 == int32(0) {
				v60 = int32(1)
				return v60
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
				v55 = m.T0[v54].(func(*base.Module, int32, int32, int32) int32)(m, l1, v53, l5)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v60 = v55 ^ int32(1)
					return v60
				}
			}
		}
	default:
		v60 = v7
		return v60
	}
}
func F_gbt_numeric_gt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2720), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_numeric_le(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2721), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_numeric_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_gbt_var_union(m, v2, v3, v4, int32(_a_F_gbt_numeric_union_0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_oid_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_oid_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
func F_gbt_text_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_gbt_var_union(m, v2, v3, v4, int32(_a_F_gbt_text_union_0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_time_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_time_sortsupport_0)
	return v4
}
func F_gbt_tsge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2439), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_tskey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13893(m, l0, l1, l2, int32(8), int32(1430))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_tslt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1498), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
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
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v358 float64
	_ = v358
	var v359 float32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
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
	if v73&int32(-4) != int32(16) {
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v100 = m.T0[v99].(func(*base.Module, int32, int32, int32, int32) int32)(m, v68, v72, l3, l5)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L23
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v92&int32(-4) != int32(16) {
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
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v264
	v268 = v19 + int32(12)
	F_gbt_var_bin_union(m, v268, v21, l3, l4, l5)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L69
	}
L23:
	;
	if v100 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v105 = int32(2)
	v106 = int32(base.Ui32(v104) >> (uint(v105) % 32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if base.Ui32(int32(base.Ui32(v107)>>(uint(v105)%32))) < base.Ui32(v106) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v181 = m.T0[v180].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v87, l3, l5)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L47
	}
L27:
	;
	v111 = int32(4)
	v112 = v68 + v111
	v114 = v21 + int32(8)
	v116 = v106 - v111
	if base.Ui32(v111) <= base.Ui32(v116) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v178 != 0 {
		goto L22
	} else {
		goto L46
	}
L29:
	;
	v178 = int32(0)
	goto L28
L30:
	;
	v152 = v147
	v153 = v148
	v154 = v149
	goto L40
L31:
	;
	if (v112|v114)&int32(3) != 0 {
		v147 = v112
		v148 = v114
		v149 = v116
		goto L30
	} else {
		goto L34
	}
L32:
	;
	v140 = v112
	v141 = v114
	v142 = v116
	goto L33
L33:
	;
	if v142 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L34:
	;
	v124 = v112
	v125 = v114
	v126 = v116
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v129 != v130 {
		v147 = v124
		v148 = v125
		v149 = v126
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v140 = v135
	v141 = v133
	v142 = v137
	goto L33
L37:
	;
	v132 = int32(4)
	v133 = v125 + v132
	v135 = v124 + v132
	v137 = v126 - v132
	if base.Ui32(int32(3)) < base.Ui32(v137) {
		v124 = v135
		v125 = v133
		v126 = v137
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v147 = v140
	v148 = v141
	v149 = v142
	goto L30
L40:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 == v158 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v178 = v157 - v158
	goto L28
L42:
	;
	v160 = int32(1)
	v165 = v154 - v160
	if v165 != 0 {
		v152 = v152 + v160
		v153 = v153 + v160
		v154 = v165
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
	if v181 <= int32(0) {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v186 = int32(2)
	v187 = int32(base.Ui32(v185) >> (uint(v186) % 32))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if base.Ui32(int32(base.Ui32(v188)>>(uint(v186)%32))) < base.Ui32(v187) {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v192 = int32(4)
	v193 = v69 + v192
	v195 = v87 + v192
	v197 = v187 - v192
	if base.Ui32(v192) <= base.Ui32(v197) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v259 == int32(0) {
		goto L18
	} else {
		goto L68
	}
L51:
	;
	v259 = int32(0)
	goto L50
L52:
	;
	v233 = v228
	v234 = v229
	v235 = v230
	goto L62
L53:
	;
	if (v193|v195)&int32(3) != 0 {
		v228 = v193
		v229 = v195
		v230 = v197
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v221 = v193
	v222 = v195
	v223 = v197
	goto L55
L55:
	;
	if v223 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v205 = v193
	v206 = v195
	v207 = v197
	goto L57
L57:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v210 != v211 {
		v228 = v205
		v229 = v206
		v230 = v207
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v221 = v216
	v222 = v214
	v223 = v218
	goto L55
L59:
	;
	v213 = int32(4)
	v214 = v206 + v213
	v216 = v205 + v213
	v218 = v207 - v213
	if base.Ui32(int32(3)) < base.Ui32(v218) {
		v205 = v216
		v206 = v214
		v207 = v218
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v228 = v221
	v229 = v222
	v230 = v223
	goto L52
L62:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v238 == v239 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v259 = v238 - v239
	goto L50
L64:
	;
	v241 = int32(1)
	v246 = v235 - v241
	if v246 != 0 {
		v233 = v233 + v241
		v234 = v234 + v241
		v235 = v246
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
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v272 = F_gbt_var_node_cp_len(m, v271, l4)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	F_gbt_var_bin_union(m, v268, v22, l3, l4, l5)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v277 = F_gbt_var_node_cp_len(m, v276, l4)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L9
	} else {
		goto L73
	}
L72:
	;
	v359 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v362 = int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+52))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	*(*float32)(unsafe.Add(mBase, uint32(l0))) = base.F32_mul(base.F32_add(base.F32_add(v359, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(v358, base.F64_convert_i32_s(v272+v362)))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v371+v362)))
	goto L18
L73:
	;
	if v277 < v272 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v358 = base.F64_convert_i32_s(v272 - v277)
	goto L72
L75:
	;
	goto L76
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v283 = int32(2)
	v285 = int32(4)
	v286 = v276 + v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v289 = int32(base.Ui32(v287) >> (uint(v283) % 32))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if base.Ui32(v277) < base.Ui32(int32(base.Ui32(v292)>>(uint(v283)%32))-v285) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v21)+8)))
	v300 = v299
	goto L79
L78:
	;
	v300 = v264
	goto L79
L79:
	;
	v301 = int32(0)
	if base.Ui32(v277) < base.Ui32(v289-int32(4)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v276)+8)))
	v308 = v307
	goto L82
L81:
	;
	v308 = v301
	goto L82
L82:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if base.Ui32(v277) < base.Ui32(int32(base.Ui32(v309)>>(uint(int32(2))%32))-int32(4)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v87)+4)))
	v317 = v316
	goto L85
L84:
	;
	v317 = v301
	goto L85
L85:
	;
	if base.Ui32(v289+v285) < base.Ui32(int32(base.Ui32(v282)>>(uint(v283)%32))) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v324 = (v289+int32(3))&int32(2147483644) + v286
	goto L88
L87:
	;
	v324 = v286
	goto L88
L88:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if base.Ui32(v277) < base.Ui32(int32(base.Ui32(v325)>>(uint(int32(2))%32))-int32(4)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v324)+4)))
	v334 = v332
	goto L91
L90:
	;
	v334 = int32(0)
	goto L91
L91:
	;
	v335 = v334 - v317
	v336 = int32(31)
	v337 = v335 >> (uint(v336) % 32)
	v340 = v300 - v308
	v342 = v340 >> (uint(v336) % 32)
	v358 = base.F64_mul(base.F64_convert_i32_u(v335^v337-v337+(v340^v342-v342)), float64(0.00390625))
	goto L72
}
