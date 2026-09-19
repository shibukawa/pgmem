package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gbt_bit_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bit_sortsupport_0)
	return v4
}
func F_gbt_biteq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2639), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_bool_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_bool_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_bool_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_bool_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_bool_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return v4 - v5
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
func F_gbt_bytea_picksplit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13902(m, l0, int32(_a_F_gbt_bytea_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_byteage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2845), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_cash_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13909(m, l0, int32(_a_F_gbt_cash_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_cash_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v28 float64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v40 float64
	_ = v40
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v17 < v16 {
		v28 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v16), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_s(v17), float64(-0.49000000953674316))), float64(0))
	} else {
		v28 = float64(0)
	}
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v30 < v29 {
		v40 = base.F64_add(v28, base.F64_add(base.F64_mul(base.F64_convert_i64_s(v29), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_s(v30), float64(-0.49000000953674316))))
	} else {
		v40 = v28
	}
	if base.F64_gt(v40, float64(0)) != 0 {
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
		*(*float32)(unsafe.Add(mBase, uint32(v13))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v40, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v17), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_s(v29), float64(-0.49000000953674316))), v40))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v58+int32(1))))
	} else {
	}
	return v13
}
func F_gbt_cash_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_cash_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
func F_gbt_date_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13905(m, l0, int32(_a_F_gbt_date_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_date_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(1428), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_gbt_dategt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2415), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_enum_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_enum_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_enumge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3794), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_float4_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v6 float32
	_ = v6
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_abs(base.F64_sub(base.F64_promote_f32(v4), base.F64_promote_f32(v6)))
}
func F_gbt_float4gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F32_gt(v4, v5)
}
func F_gbt_float4le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F32_le(v4, v5)
}
func F_gbt_float8_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_float8_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_float8_consistent(m *base.Module, l0 int32) int32 {
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
	var v11 float64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	*(*float64)(unsafe.Add(mBase, uint32(v7)+24)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+22)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v15
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v30)+12)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(_a_F_gbt_float8_consistent_0), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v37
	}
}
func F_gbt_float8gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func F_gbt_inet_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_gt(v7, v8)
	}
	return v11
}
func F_gbt_int2_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_int2_union_0), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int4_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13911(m, l0, int32(_a_F_gbt_int4_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_int4_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int4_sortsupport_0)
	return int32(0)
}
func F_gbt_int8_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_int8_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_intv_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v46 float64
	_ = v46
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v76 float64
	_ = v76
	var v89 float64
	_ = v89
	var v97 float64
	_ = v97
	var v100 float64
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v12 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	v40 = float64(2.592e+06)
	v43 = float64(86400)
	v46 = float64(1e+06)
	v49 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v22), v40), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v23), v43), base.F64_div(base.F64_convert_i64_s(v24), v46)))
	v50 = float64(0.49000000953674316)
	v62 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v27), v40), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v28), v43), base.F64_div(base.F64_convert_i64_s(v29), v46)))
	v63 = float64(-0.49000000953674316)
	v76 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v30), v40), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v31), v43), base.F64_div(base.F64_convert_i64_s(v32), v46)))
	v89 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v33), v40), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v34), v43), base.F64_div(base.F64_convert_i64_s(v35), v46)))
	if base.F64_gt(v76, v89) != 0 {
		v97 = base.F64_add(base.F64_add(base.F64_mul(v76, v50), base.F64_mul(v89, v63)), v12)
	} else {
		v97 = v12
	}
	if base.F64_lt(v62, v49) != 0 {
		v100 = base.F64_add(base.F64_add(base.F64_mul(v49, v50), base.F64_mul(v62, v63)), v97)
	} else {
		v100 = v97
	}
	if base.F64_gt(v100, float64(0)) != 0 {
		v114 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+52))
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
		*(*float32)(unsafe.Add(mBase, uint32(v36))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v100, base.F64_add(base.F64_add(base.F64_mul(v89, float64(0.49000000953674316)), base.F64_mul(v49, float64(-0.49000000953674316))), v100))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v116+int32(1))))
	} else {
	}
	return v36
}
func F_gbt_intv_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_intv_sortsupport_0)
	return v4
}
func F_gbt_intvge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2446), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_intvkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13913(m, l0, l1, l2, int32(16), int32(2524))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_intvlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2444), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13905(m, l0, int32(_a_F_gbt_macad8_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_macad8le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8le_0), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_macad_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macadgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2263), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macadle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2262), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v45 float32
	_ = v45
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v11 != int32(1) {
		v93 = l0
		m.G0 = v9 + int32(16)
		return v93
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v28)
				v74 = v9 + int32(8)
			case 1:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32
				v74 = v9 + int32(8)
			case 2:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v37
				v74 = v9 + int32(8)
			case 3:
				v45 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
				*(*float32)(unsafe.Add(mBase, uint32(v9)+8)) = v45
				v74 = v9 + int32(8)
			case 4:
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v49)))
				*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v50
				v74 = v9 + int32(8)
			default:
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v74 = v73
			case 6:
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v64
				v74 = v9 + int32(8)
			case 7:
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v69
				v74 = v9 + int32(8)
			case 8, 21:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
				v74 = v9 + int32(8)
			case 9:
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v59
				v74 = v9 + int32(8)
			case 10:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v54
				v74 = v9 + int32(8)
			case 18:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(base.B2i32(v22 != int32(0)))
				v74 = v9 + int32(8)
			}
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v75 != 0 {
				base.MemoryCopy(m, v15, v74, v75)
			} else {
			}
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v77 != 0 {
				base.MemoryCopy(m, v77+v15, v74, v77)
			} else {
			}
			v81 = F_palloc(m, int32(16))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v81))) = v15
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v84
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v86
				v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v89 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v81)+14)) = uint8(v89)
				*(*uint16)(unsafe.Add(mBase, uint32(v81)+12)) = uint16(v88)
				v93 = v81
				m.G0 = v9 + int32(16)
				return v93
			}
		}
	}
}
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v5 - int32(1) {
	case 0:
		v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4))))
		v30 = v9
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
	case 1:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v30 = v10
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
	case 2:
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v12 = F_Int64GetDatum(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v30 = v12
			v32 = F_palloc(m, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
				*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
				return v32
			}
		}
	case 3, 8, 21:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v30 = v16
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
	case 4:
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		v18 = F_Float8GetDatum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v30 = v18
			v32 = F_palloc(m, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
				*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
				return v32
			}
		}
	default:
		v30 = v4
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
	case 6:
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v25 = F_Int64GetDatum(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v30 = v25
			v32 = F_palloc(m, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
				*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
				return v32
			}
		}
	case 7:
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v28 = F_Int64GetDatum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = v28
			v32 = F_palloc(m, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
				*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
				return v32
			}
		}
	case 9:
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v22 = F_Int64GetDatum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v30 = v22
			v32 = F_palloc(m, int32(16))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
				*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
				return v32
			}
		}
	case 10:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v30 = v20
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
	case 18:
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		v30 = v8
		v32 = F_palloc(m, int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v37
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v40 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v40)
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)) = uint16(v39)
			return v32
		}
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
func F_gbt_numeric_picksplit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13902(m, l0, int32(_a_F_gbt_numeric_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_oid_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v5) < base.Ui32(v4)) - base.B2i32(base.Ui32(v4) < base.Ui32(v5))
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
func F_gbt_text_penalty(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_var_penalty(m, v2, v3, v4, v5, int32(_a_F_gbt_text_penalty_0), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_gbt_text_picksplit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13902(m, l0, int32(_a_F_gbt_text_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_text_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_text_sortsupport_0)
	return v4
}
func F_gbt_textcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2104), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_texteq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1544), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_textge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2231), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13909(m, l0, int32(_a_F_gbt_time_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_time_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_time_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_time_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_time_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_timelt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2418), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_ts_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13909(m, l0, int32(_a_F_gbt_ts_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_ts_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v2 = float64(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v24 = base.F64_convert_i64_s(v15)
	v25 = float64(0.49000000953674316)
	v27 = base.F64_convert_i64_s(v18)
	v28 = float64(-0.49000000953674316)
	v31 = base.F64_convert_i64_s(v19)
	v34 = base.F64_convert_i64_s(v20)
	if base.F64_gt(v31, v34) != 0 {
		v42 = base.F64_add(base.F64_add(base.F64_mul(v31, v25), base.F64_mul(v34, v28)), v2)
	} else {
		v42 = v2
	}
	if base.F64_lt(v27, v24) != 0 {
		v45 = base.F64_add(base.F64_add(base.F64_mul(v24, v25), base.F64_mul(v27, v28)), v42)
	} else {
		v45 = v42
	}
	if base.F64_gt(v45, float64(0)) != 0 {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
		*(*float32)(unsafe.Add(mBase, uint32(v21))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v45, base.F64_add(base.F64_add(base.F64_mul(v34, float64(0.49000000953674316)), base.F64_mul(v24, float64(-0.49000000953674316))), v45))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v61+int32(1))))
	} else {
	}
	return v21
}
func F_gbt_ts_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_ts_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_tseq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2436), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_uuid_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_uuid_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_uuid_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_uuid_union_0), int32(32))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = v22 + v20
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v31 = (v27 + int32(3)) & int32(2147483644)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui32(v27+v20) < base.Ui32(int32(base.Ui32(v35)>>(uint(v26)%32))) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = v24 + v31
	goto L3
L2:
	;
	v39 = v24
	goto L3
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
	v45 = v42 + v31 + int32(4)
	v46 = F_palloc0(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v51 = v46 + int32(4)
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	base.MemoryCopy(m, v51, v24, v27)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	base.MemoryCopy(m, v31+v51, v39, v42)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v55 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v45 << (uint(v55) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v46
	if v55 <= v19 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v65 = int32(1)
	goto L15
L13:
	;
	v91 = v46
	goto L14
L14:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v104 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v65<<(uint(int32(4))%32))))
	F_gbt_var_bin_union(m, v17+int32(12), v83, l2, l3, l4)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v91 = v89
	goto L14
L17:
	;
	v87 = v65 + int32(1)
	if v87 != v19 {
		v65 = v87
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v107 = F_gbt_var_node_cp_len(m, v91, l3)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v167 = v91
	goto L21
L21:
	;
	m.G0 = v17 + int32(16)
	return v167
L22:
	;
	v109 = int32(4)
	v110 = v91 + v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = int32(2)
	v113 = int32(base.Ui32(v111) >> (uint(v112) % 32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if base.Ui32(v113+v109) < base.Ui32(int32(base.Ui32(v121)>>(uint(v112)%32))) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v125 = v110 + (v113+int32(3))&int32(2147483644)
	goto L25
L24:
	;
	v125 = v110
	goto L25
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v127 = int32(2)
	v130 = int32(base.Ui32(v126)>>(uint(v127)%32)) - int32(4)
	v132 = v107 + v127
	if v130 < v132 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v134 = v130
	goto L28
L27:
	;
	v134 = v132
	goto L28
L28:
	;
	v136 = v113 - int32(4)
	if v136 < v132 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v138 = v136
	goto L31
L30:
	;
	v138 = v132
	goto L31
L31:
	;
	v142 = (v138 + int32(7)) & int32(-4)
	v145 = v134 + v142 + int32(8)
	v146 = F_palloc0(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v145 << (uint(int32(2)) % 32)
	v151 = int32(4)
	v152 = v146 + v151
	v154 = v138 + v151
	if v154 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	base.MemoryCopy(m, v152, v110, v154)
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v154 << (uint(int32(2)) % 32)
	v159 = v152 + v142
	v161 = v134 + int32(4)
	if v161 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	base.MemoryCopy(m, v159, v125, v161)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161 << (uint(int32(2)) % 32)
	v167 = v146
	goto L21
}
