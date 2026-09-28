package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bit_sortsupport_0)
	return int64(0)
}
func F_gbt_biteq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2853), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_bool_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_bool_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_bool_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_bool_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_bool_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0)))))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)))))
	return v5 - v7
}
func F_gbt_boolle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(base.Ui32(v4) <= base.Ui32(v5))
}
func F_gbt_bytea_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14268(m, l0, int32(_a_F_gbt_bytea_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_byteage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(3060), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_cash_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14277(m, l0, int32(_a_F_gbt_cash_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_cash_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v22 int64
	_ = v22
	var v23 float64
	_ = v23
	var v24 int64
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
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_convert_i64_s(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v17 = base.F64_convert_i64_s(v16)
	if base.F64_lt(v17, v13) != 0 {
		v21 = base.F64_sub(v13, v17)
	} else {
		v21 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v23 = base.F64_convert_i64_s(v22)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v25 = base.F64_convert_i64_s(v24)
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
func F_gbt_cash_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_cash_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_casheq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 == v5)
}
func F_gbt_cashge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v5 <= v4)
}
func F_gbt_cashlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 < v5)
}
func F_gbt_date_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_date_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_date_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0)))))
	v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)))))
	v10 = F_DirectFunctionCall2Coll(m, int32(1576), int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
	}
}
func F_gbt_dategt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2628), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_enum_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_enum_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_enumge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	v7 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
	v8 = F_CallerFInfoFunctionCall2(m, int32(4009), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_float4_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v36 float64
	_ = v36
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v9 = base.F64_promote_f32(v8)
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_promote_f32(v10)
	v12 = base.F64_sub(v9, v11)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
		v19 = int64(9223372036854775807)
		v20 = base.I64_reinterpret_f64(v11) & v19
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v9)&v19) {
			if base.Ui64(v20) <= base.Ui64(int64(9218868437227405312)) {
				v36 = math.Float64frombits(uint64(0x7ff0000000000000))
				return base.F64_abs(v36)
			} else {
				return float64(0)
			}
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v20) {
				v36 = math.Float64frombits(uint64(0x7ff0000000000000))
				return base.F64_abs(v36)
			} else {
				return float64(0)
			}
		}
	} else {
		v36 = v12
		return base.F64_abs(v36)
	}
}
func F_gbt_float4gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v12 float32
	_ = v12
	var v22 int32
	_ = v22
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(base.I32_reinterpret_f32(v6)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
		v22 = base.F32_lt(v6, v12) | base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v12)&int32(2147483647)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_gbt_float4le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v12 float32
	_ = v12
	var v22 int32
	_ = v22
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(base.I32_reinterpret_f32(v6)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
		v22 = base.F32_ge(v6, v12) & base.B2i32(base.Ui32(base.I32_reinterpret_f32(v12)&int32(2147483647)) < base.Ui32(int32(2139095041)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_gbt_float8_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_float8_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_float8_consistent(m *base.Module, l0 int32) int64 {
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
	var v10 float64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v7)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+22)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v14 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v14
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v29)+12)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v31&int32(1), int32(_a_F_gbt_float8_consistent_0), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(32)
		return base.I64_extend_i32_u(v36)
	}
}
func F_gbt_float8gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
		v22 = base.F64_lt(v6, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_gbt_inet_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	v8 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	if base.F64_lt(v8, v10) != 0 {
		v13 = int32(-1)
	} else {
		v13 = base.F64_gt(v8, v10)
	}
	return v13
}
func F_gbt_int2_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_int2_union_0), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_int4_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14279(m, l0, int32(_a_F_gbt_int4_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_int4_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int4_sortsupport_0)
	return int64(0)
}
func F_gbt_int8_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_int8_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_intv_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v19 float64
	_ = v19
	var v21 int64
	_ = v21
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v42 float64
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v71 float64
	_ = v71
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 float32
	_ = v99
	v4 = float64(0)
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = float64(2.592e+06)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v19 = float64(86400)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v23 = float64(1e+06)
	v26 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v13), v15), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v17), v19), base.F64_div(base.F64_convert_i64_s(v21), v23)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v42 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v29), v15), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v33), v19), base.F64_div(base.F64_convert_i64_s(v37), v23)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	v57 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v44), v15), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v48), v19), base.F64_div(base.F64_convert_i64_s(v52), v23)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v71 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v58), v15), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v62), v19), base.F64_div(base.F64_convert_i64_s(v66), v23)))
	if base.F64_gt(v57, v71) != 0 {
		v77 = base.F64_add(base.F64_sub(v57, v71), v4)
	} else {
		v77 = v4
	}
	if base.F64_lt(v42, v26) != 0 {
		v80 = base.F64_add(base.F64_sub(v26, v42), v77)
	} else {
		v80 = v77
	}
	if base.F64_gt(v80, float64(0)) != 0 {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+52))
		v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
		v99 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v80, base.F64_add(base.F64_sub(v71, v26), v80))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v92+int32(1))))
	} else {
		v99 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v9)))) = v99
	return v9
}
func F_gbt_intv_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_intv_sortsupport_0)
	return int64(0)
}
func F_gbt_intvge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2659), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_intvkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14281(m, l0, l1, l2, int32(16), int32(2737))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_intvlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2657), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad8_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_macad8_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macad8le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8le_0), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_macad_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_macadgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2483), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macadle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2482), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_num_compress(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	if v11 != int32(1) {
		v89 = l0
		m.G0 = v9 + int32(16)
		return v89
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v15 = F_palloc0(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v19 - int32(1) {
			case 0:
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v28)
				v69 = v9 + int32(8)
			case 1:
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v32)
				v69 = v9 + int32(8)
			case 2:
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v36
				v69 = v9 + int32(8)
			case 3:
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v44)
				v69 = v9 + int32(8)
			case 4:
				v48 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
				*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v48
				v69 = v9 + int32(8)
			default:
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v69 = v68
			case 6:
				v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v60
				v69 = v9 + int32(8)
			case 7:
				v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v64
				v69 = v9 + int32(8)
			case 8, 21:
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v40)
				v69 = v9 + int32(8)
			case 9:
				v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v56
				v69 = v9 + int32(8)
			case 10:
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v52)
				v69 = v9 + int32(8)
			case 18:
				v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(base.B2i32(v22 != int64(0)))
				v69 = v9 + int32(8)
			}
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v70 != 0 {
				base.MemoryCopy(m, v15, v69, v70)
			} else {
			}
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v72 != 0 {
				base.MemoryCopy(m, v72+v15, v69, v72)
			} else {
			}
			v76 = F_palloc(m, int32(24))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v76))) = base.I64_extend_i32_u(v15)
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v80
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v82
				v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				v85 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v76)+18)) = uint8(v85)
				*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)) = uint16(v84)
				v89 = v76
				m.G0 = v9 + int32(16)
				return v89
			}
		}
	}
}
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v5 - int32(1) {
	case 0:
		v11 = int64(*(*int16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v11
	case 1:
		v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v13
	case 2:
		v15 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4))))
		v30 = v15
	case 3:
		v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v19
	case 4:
		v21 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4))))
		v30 = v21
	default:
		v30 = v4
	case 6:
		v27 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4))))
		v30 = v27
	case 7:
		v29 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4))))
		v30 = v29
	case 8, 21:
		v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v17
	case 9:
		v25 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4))))
		v30 = v25
	case 10:
		v23 = int64(*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v23
	case 18:
		v9 = int64(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))))
		v30 = v9
	}
	v32 = F_palloc(m, int32(24))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v32))) = v30
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v39
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v42 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v32)+18)) = uint8(v42)
		*(*uint16)(unsafe.Add(mBase, uint32(v32)+16)) = uint16(v41)
		return v32
	}
}
func F_gbt_num_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l0+v6, l1+v6, l3)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = v15
				return v18
			}
		} else {
			v18 = int32(0)
			return v18
		}
	}
}
func F_gbt_numeric_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14268(m, l0, int32(_a_F_gbt_numeric_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_oid_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v6 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	return base.B2i32(base.Ui32(v8) < base.Ui32(v6)) - base.B2i32(base.Ui32(v6) < base.Ui32(v8))
}
func F_gbt_oidgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v5) < base.Ui32(v4))
}
func F_gbt_text_penalty(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14267(m, l0, int32(_a_F_gbt_text_penalty_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_text_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14268(m, l0, int32(_a_F_gbt_text_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_text_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_text_sortsupport_0)
	return int64(0)
}
func F_gbt_textcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2325), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v8)
	}
}
func F_gbt_texteq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(1757), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_textge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2451), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_time_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14277(m, l0, int32(_a_F_gbt_time_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_time_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_time_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_time_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_time_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_timelt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2631), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_ts_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14277(m, l0, int32(_a_F_gbt_ts_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_ts_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_ts_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_tseq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2649), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_uuid_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_uuid_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_uuid_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_uuid_union_0), int32(32))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_uuidkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v237 int64
	_ = v237
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v311 int64
	_ = v311
	var v313 int32
	_ = v313
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = int64(56)
	v10 = int64(65280)
	v12 = int64(40)
	v15 = int64(16711680)
	v17 = int64(24)
	v19 = int64(4278190080)
	v21 = int64(8)
	v42 = v7<<(uint(v8)%64) | v7&v10<<(uint(v12)%64) | (v7&v15<<(uint(v17)%64) | v7&v19<<(uint(v21)%64)) | (int64(base.Ui64(v7)>>(uint(v21)%64))&v19 | int64(base.Ui64(v7)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v7)>>(uint(v12)%64))&v10 | int64(base.Ui64(v7)>>(uint(v8)%64))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v79 = v44<<(uint(v8)%64) | v44&v10<<(uint(v12)%64) | (v44&v15<<(uint(v17)%64) | v44&v19<<(uint(v21)%64)) | (int64(base.Ui64(v44)>>(uint(v21)%64))&v19 | int64(base.Ui64(v44)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v44)>>(uint(v12)%64))&v10 | int64(base.Ui64(v44)>>(uint(v8)%64))))
	if v42 == v79 {
		v82 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v83 = int64(56)
		v85 = int64(65280)
		v87 = int64(40)
		v90 = int64(16711680)
		v92 = int64(24)
		v94 = int64(4278190080)
		v96 = int64(8)
		v117 = v82<<(uint(v83)%64) | v82&v85<<(uint(v87)%64) | (v82&v90<<(uint(v92)%64) | v82&v94<<(uint(v96)%64)) | (int64(base.Ui64(v82)>>(uint(v96)%64))&v94 | int64(base.Ui64(v82)>>(uint(v92)%64))&v90 | (int64(base.Ui64(v82)>>(uint(v87)%64))&v85 | int64(base.Ui64(v82)>>(uint(v83)%64))))
		v118 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
		v153 = v118<<(uint(v83)%64) | v118&v85<<(uint(v87)%64) | (v118&v90<<(uint(v92)%64) | v118&v94<<(uint(v96)%64)) | (int64(base.Ui64(v118)>>(uint(v96)%64))&v94 | int64(base.Ui64(v118)>>(uint(v92)%64))&v90 | (int64(base.Ui64(v118)>>(uint(v87)%64))&v85 | int64(base.Ui64(v118)>>(uint(v83)%64))))
		if v117 == v153 {
			v163 = int32(0)
		} else {
			v155 = v153
			v156 = v117
			if base.Ui64(v156) < base.Ui64(v155) {
				v160 = int32(-1)
			} else {
				v160 = int32(1)
			}
			v163 = v160
		}
	} else {
		v155 = v79
		v156 = v42
		if base.Ui64(v156) < base.Ui64(v155) {
			v160 = int32(-1)
		} else {
			v160 = int32(1)
		}
		v163 = v160
	}
	if v163 == int32(0) {
		v166 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
		v167 = int64(56)
		v169 = int64(65280)
		v171 = int64(40)
		v174 = int64(16711680)
		v176 = int64(24)
		v178 = int64(4278190080)
		v180 = int64(8)
		v201 = v166<<(uint(v167)%64) | v166&v169<<(uint(v171)%64) | (v166&v174<<(uint(v176)%64) | v166&v178<<(uint(v180)%64)) | (int64(base.Ui64(v166)>>(uint(v180)%64))&v178 | int64(base.Ui64(v166)>>(uint(v176)%64))&v174 | (int64(base.Ui64(v166)>>(uint(v171)%64))&v169 | int64(base.Ui64(v166)>>(uint(v167)%64))))
		v202 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
		v237 = v202<<(uint(v167)%64) | v202&v169<<(uint(v171)%64) | (v202&v174<<(uint(v176)%64) | v202&v178<<(uint(v180)%64)) | (int64(base.Ui64(v202)>>(uint(v180)%64))&v178 | int64(base.Ui64(v202)>>(uint(v176)%64))&v174 | (int64(base.Ui64(v202)>>(uint(v171)%64))&v169 | int64(base.Ui64(v202)>>(uint(v167)%64))))
		if v201 != v237 {
			v318 = v237
			v319 = v201
			if base.Ui64(v319) < base.Ui64(v318) {
				v323 = int32(-1)
			} else {
				v323 = int32(1)
			}
			return v323
		} else {
			v240 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v241 = int64(56)
			v243 = int64(65280)
			v245 = int64(40)
			v248 = int64(16711680)
			v250 = int64(24)
			v252 = int64(4278190080)
			v254 = int64(8)
			v275 = v240<<(uint(v241)%64) | v240&v243<<(uint(v245)%64) | (v240&v248<<(uint(v250)%64) | v240&v252<<(uint(v254)%64)) | (int64(base.Ui64(v240)>>(uint(v254)%64))&v252 | int64(base.Ui64(v240)>>(uint(v250)%64))&v248 | (int64(base.Ui64(v240)>>(uint(v245)%64))&v243 | int64(base.Ui64(v240)>>(uint(v241)%64))))
			v276 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
			v311 = v276<<(uint(v241)%64) | v276&v243<<(uint(v245)%64) | (v276&v248<<(uint(v250)%64) | v276&v252<<(uint(v254)%64)) | (int64(base.Ui64(v276)>>(uint(v254)%64))&v252 | int64(base.Ui64(v276)>>(uint(v250)%64))&v248 | (int64(base.Ui64(v276)>>(uint(v245)%64))&v243 | int64(base.Ui64(v276)>>(uint(v241)%64))))
			if v275 != v311 {
				v318 = v311
				v319 = v275
				if base.Ui64(v319) < base.Ui64(v318) {
					v323 = int32(-1)
				} else {
					v323 = int32(1)
				}
				return v323
			} else {
				v313 = int32(0)
				return v313
			}
		}
	} else {
		v313 = v163
		return v313
	}
}
func F_gbt_var_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v66 int32
	_ = v66
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v179 int64
	_ = v179
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = v23 + v21
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = int32(2)
	v28 = int32(base.Ui32(v26) >> (uint(v27) % 32))
	v32 = (v28 + int32(3)) & int32(2147483644)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v28+v21) < base.Ui32(int32(base.Ui32(v36)>>(uint(v27)%32))) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = v25 + v32
	goto L3
L2:
	;
	v40 = v25
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v43 = int32(base.Ui32(v41) >> (uint(int32(2)) % 32))
	v46 = v43 + v32 + int32(4)
	v47 = F_palloc0(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v52 = v47 + int32(4)
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	base.MemoryCopy(m, v52, v25, v28)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	base.MemoryCopy(m, v32+v52, v40, v43)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v56 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46 << (uint(v56) % 32)
	v59 = base.I64_extend_i32_u(v47)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v59
	if v56 <= v20 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v66 = int32(1)
	goto L15
L13:
	;
	v107 = v59
	goto L14
L14:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v108 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v66*int32(24))))
	F_gbt_var_bin_union(m, v18+int32(8), v86, l2, l3, l4)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v107 = v92
	goto L14
L17:
	;
	v90 = v66 + int32(1)
	if v90 != v20 {
		v66 = v90
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v111 = base.I32_wrap_i64(v107)
	v112 = F_gbt_var_node_cp_len(m, v111, l3)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v179 = v107
	goto L21
L21:
	;
	m.G0 = v18 + int32(16)
	return base.I32_wrap_i64(v179)
L22:
	;
	v114 = int32(4)
	v115 = v111 + v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = int32(2)
	v118 = int32(base.Ui32(v116) >> (uint(v117) % 32))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if base.Ui32(v118+v114) < base.Ui32(int32(base.Ui32(v126)>>(uint(v117)%32))) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v130 = v115 + (v118+int32(3))&int32(2147483644)
	goto L25
L24:
	;
	v130 = v115
	goto L25
L25:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = int32(2)
	v135 = int32(base.Ui32(v131)>>(uint(v132)%32)) - int32(4)
	v137 = v112 + v132
	if v135 < v137 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v139 = v135
	goto L28
L27:
	;
	v139 = v137
	goto L28
L28:
	;
	v141 = v118 - int32(4)
	if v141 < v137 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v143 = v141
	goto L31
L30:
	;
	v143 = v137
	goto L31
L31:
	;
	v147 = (v143 + int32(7)) & int32(-4)
	v150 = v139 + v147 + int32(8)
	v151 = F_palloc0(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v150 << (uint(int32(2)) % 32)
	v156 = int32(4)
	v157 = v151 + v156
	v159 = v143 + v156
	if v159 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	base.MemoryCopy(m, v157, v115, v159)
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v159 << (uint(int32(2)) % 32)
	v164 = v157 + v147
	v166 = v139 + int32(4)
	if v166 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	base.MemoryCopy(m, v164, v130, v166)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v166 << (uint(int32(2)) % 32)
	v179 = base.I64_extend_i32_u(v151)
	goto L21
}
