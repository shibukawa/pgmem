package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gbt_bit_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6867)
	return v3
}
func F_gbt_biteq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2655), int32(0), l0, l1)
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
	v4 = F_gbt_num_compress(m, v2, int32(4378616))
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
	v4 = F_gbt_num_fetch(m, v2, int32(4378616))
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
	mBase := m.M
	_ = mBase
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v3, v4, v5, int32(4378656), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_byteage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2861), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_cash_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v7, v7+v14, v23&int32(1), int32(4378696), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v32
		}
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4378696), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
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
	mBase := m.M
	_ = mBase
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_num_same(m, v4, v5, int32(4378736), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
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
	v8 = F_DirectFunctionCall2Coll(m, int32(1444), int32(0), v6, v7)
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2431), int32(0), v6, v7)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4378776), v5)
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3810), l2, int32(0), v6, v7)
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
	v4 = F_gbt_num_compress(m, v2, int32(4378856))
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
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(4378856), v36)
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(4)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4379000), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_int4_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v13 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12 + v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v7+v13, v7+int32(12), v24&int32(1), int32(4379040), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v33 = F_Float8GetDatum(m, v29)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v33
		}
	}
}
func F_gbt_int4_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6943)
	return int32(0)
}
func F_gbt_int8_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4379080), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6959)
	return v3
}
func F_gbt_intvge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2462), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_intvkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = F_DirectFunctionCall2Coll(m, int32(2540), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v20 = v8
			return v20
		} else {
			v14 = int32(16)
			v18 = F_DirectFunctionCall2Coll(m, int32(2540), int32(0), v6+v14, v7+v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = v18
				return v20
			}
		}
	}
}
func F_gbt_intvlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2460), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_num_same(m, v4, v5, int32(4379200), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_macad8le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(4190), int32(0), l0, l1)
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
	v4 = F_gbt_num_fetch(m, v2, int32(4379160))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2279), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2278), int32(0), l0, l1)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v44 float32
	_ = v44
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v10 != int32(1) {
		v94 = l0
		m.G0 = v8 + int32(16)
		return v94
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = F_palloc0(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v18 - int32(1) {
			case 0:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)) = uint16(v27)
				v73 = v8 + int32(8)
			case 1:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v31
				v73 = v8 + int32(8)
			case 2:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v36
				v73 = v8 + int32(8)
			case 3:
				v44 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
				*(*float32)(unsafe.Add(mBase, uint32(v8)+8)) = v44
				v73 = v8 + int32(8)
			case 4:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
				*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v49
				v73 = v8 + int32(8)
			default:
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v73 = v72
			case 6:
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v63
				v73 = v8 + int32(8)
			case 7:
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v68
				v73 = v8 + int32(8)
			case 8, 21:
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v40
				v73 = v8 + int32(8)
			case 9:
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v58
				v73 = v8 + int32(8)
			case 10:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v53
				v73 = v8 + int32(8)
			case 18:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)) = uint8(base.B2i32(v21 != int32(0)))
				v73 = v8 + int32(8)
			}
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v74 != 0 {
				v75 = F__emscripten_memcpy_bulkmem(m, v14, v73, v74)
				mBase = m.M
				v76 = v75
			} else {
				v76 = v14
			}
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v77 != 0 {
				v79 = F__emscripten_memcpy_bulkmem(m, v76+v77, v73, v77)
				mBase = m.M
			} else {
			}
			v82 = F_palloc(m, int32(16))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v82))) = v76
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v87
				v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v82)+14)) = uint8(v90)
				*(*uint16)(unsafe.Add(mBase, uint32(v82)+12)) = uint16(v89)
				v94 = v82
				m.G0 = v8 + int32(16)
				return v94
			}
		}
	}
}
func F_gbt_num_fetch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
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
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v4 - int32(1) {
	case 0:
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9))))
		v42 = v10
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 1:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v42 = v12
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 2:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_Int64GetDatum(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v42 = v15
			v44 = F_palloc(m, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
				return v44
			}
		}
	case 3:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v42 = v22
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 4:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
		v25 = F_Float8GetDatum(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v42 = v25
			v44 = F_palloc(m, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
				return v44
			}
		}
	default:
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = v41
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 6:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
		v35 = F_Int64GetDatum(m, v34)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v42 = v35
			v44 = F_palloc(m, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
				return v44
			}
		}
	case 7:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
		v39 = F_Int64GetDatum(m, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v42 = v39
			v44 = F_palloc(m, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
				return v44
			}
		}
	case 8, 21:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v42 = v20
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 9:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
		v31 = F_Int64GetDatum(m, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v42 = v31
			v44 = F_palloc(m, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v52 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
				return v44
			}
		}
	case 10:
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		v42 = v28
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	case 18:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v42 = v8
		v44 = F_palloc(m, int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v47
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v52)
			*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v51)
			return v44
		}
	}
}
func F_gbt_num_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, l0+v7, l1+v7, l3)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = v16
				return v18
			}
		} else {
			v18 = int32(0)
			return v18
		}
	}
}
func F_gbt_numeric_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v3, v4, v5, int32(4379240), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v4
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
	v8 = F_gbt_var_penalty(m, v2, v3, v4, v5, int32(4379320), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_gbt_text_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v3, v4, v5, int32(4379320), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_text_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6996)
	return v3
}
func F_gbt_textcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2120), l2, l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1560), l2, l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2247), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v7, v7+v14, v23&int32(1), int32(4379400), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v32
		}
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4379400), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_time_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4379400), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_timelt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2434), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_ts_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v7, v7+v14, v23&int32(1), int32(4379440), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v32
		}
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4379440), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_tseq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2452), int32(0), l0, l1)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4379480), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_uuid_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(32)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4379480), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_uuidkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = int32(16)
	goto L4
L1:
	;
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v68 = int32(0)
	goto L1
L3:
	;
	v42 = v37
	v43 = v38
	v44 = v39
	goto L13
L4:
	;
	if (v4|v5)&int32(3) != 0 {
		v37 = v4
		v38 = v5
		v39 = v6
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v27 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v14 = v4
	v15 = v5
	v16 = v6
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v19 != v20 {
		v37 = v14
		v38 = v15
		v39 = v16
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v22 = int32(4)
	v23 = v15 + v22
	v25 = v14 + v22
	v27 = v16 - v22
	if base.Ui32(int32(3)) < base.Ui32(v27) {
		v14 = v25
		v15 = v23
		v16 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = v25
	v38 = v23
	v39 = v27
	goto L3
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 == v48 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v68 = v47 - v48
	goto L1
L15:
	;
	v50 = int32(1)
	v55 = v44 - v50
	if v55 != 0 {
		v42 = v42 + v50
		v43 = v43 + v50
		v44 = v55
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
L19:
	;
	v136 = v68
	goto L21
L20:
	;
	v69 = int32(16)
	v70 = v4 + v69
	v72 = v5 + v69
	goto L25
L21:
	;
	return v136
L22:
	;
	v136 = v135
	goto L21
L23:
	;
	v135 = int32(0)
	goto L22
L24:
	;
	v109 = v104
	v110 = v105
	v111 = v106
	goto L34
L25:
	;
	if (v70|v72)&int32(3) != 0 {
		v104 = v70
		v105 = v72
		v106 = v69
		goto L24
	} else {
		goto L28
	}
L27:
	;
	if v94 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v81 = v70
	v82 = v72
	v83 = v69
	goto L29
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v86 != v87 {
		v104 = v81
		v105 = v82
		v106 = v83
		goto L24
	} else {
		goto L31
	}
L30:
	;
	goto L27
L31:
	;
	v89 = int32(4)
	v90 = v82 + v89
	v92 = v81 + v89
	v94 = v83 - v89
	if base.Ui32(int32(3)) < base.Ui32(v94) {
		v81 = v92
		v82 = v90
		v83 = v94
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = v92
	v105 = v90
	v106 = v94
	goto L24
L34:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == v115 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v135 = v114 - v115
	goto L22
L36:
	;
	v117 = int32(1)
	v122 = v111 - v117
	if v122 != 0 {
		v109 = v109 + v117
		v110 = v110 + v117
		v111 = v122
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L23
}
func F_gbt_var_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = v21 + v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = int32(2)
	v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
	v30 = (v26 + int32(3)) & int32(2147483644)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if base.Ui32(v26+v19) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = v23 + v30
	goto L3
L2:
	;
	v38 = v23
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(2)) % 32))
	v44 = v41 + v30 + int32(4)
	v45 = F_palloc0(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v50 = v45 + int32(4)
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v51 = F__emscripten_memcpy_bulkmem(m, v50, v23, v26)
	mBase = m.M
	v52 = v51
	goto L9
L8:
	;
	v52 = v50
	goto L9
L9:
	;
	goto L6
L10:
	;
	v56 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v44 << (uint(v56) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v45
	if v56 <= v18 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v52+v30, v38, v41)
	mBase = m.M
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	v66 = int32(1)
	goto L17
L15:
	;
	v91 = v45
	goto L16
L16:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v103 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v66<<(uint(int32(4))%32))))
	F_gbt_var_bin_union(m, v16+int32(12), v83, l2, l3, l4)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v91 = v89
	goto L16
L19:
	;
	v87 = v66 + int32(1)
	if v87 != v18 {
		v66 = v87
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v106 = F_gbt_var_node_cp_len(m, v91, l3)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v168 = v91
	goto L23
L23:
	;
	m.G0 = v16 + int32(16)
	return v168
L24:
	;
	v108 = int32(4)
	v109 = v91 + v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = int32(2)
	v112 = int32(base.Ui32(v110) >> (uint(v111) % 32))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if base.Ui32(v112+v108) < base.Ui32(int32(base.Ui32(v120)>>(uint(v111)%32))) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v109 + (v112+int32(3))&int32(2147483644)
	goto L27
L26:
	;
	v124 = v109
	goto L27
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(2)
	v129 = int32(base.Ui32(v125)>>(uint(v126)%32)) - int32(4)
	v131 = v106 + v126
	if v129 < v131 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v133 = v129
	goto L30
L29:
	;
	v133 = v131
	goto L30
L30:
	;
	v135 = v112 - int32(4)
	if v135 < v131 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v137 = v135
	goto L33
L32:
	;
	v137 = v131
	goto L33
L33:
	;
	v141 = (v137 + int32(7)) & int32(-4)
	v144 = v133 + v141 + int32(8)
	v145 = F_palloc0(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v144 << (uint(int32(2)) % 32)
	v150 = int32(4)
	v151 = v145 + v150
	v153 = v137 + v150
	if v153 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v153 << (uint(int32(2)) % 32)
	v159 = v155 + v141
	v161 = v133 + int32(4)
	if v161 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v154 = F__emscripten_memcpy_bulkmem(m, v151, v109, v153)
	mBase = m.M
	v155 = v154
	goto L38
L37:
	;
	v155 = v151
	goto L38
L38:
	;
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v161 << (uint(int32(2)) % 32)
	v168 = v145
	goto L23
L40:
	;
	v162 = F__emscripten_memcpy_bulkmem(m, v159, v124, v161)
	mBase = m.M
	v163 = v162
	goto L42
L41:
	;
	v163 = v159
	goto L42
L42:
	;
	goto L39
}
