package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gbt_bit_penalty(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_var_penalty(m, v2, v3, v4, v5, int32(_a_F_gbt_bit_penalty_0), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_gbt_bit_union(m *base.Module, l0 int32) int32 {
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
	v7 = F_gbt_var_union(m, v2, v3, v4, int32(_a_F_gbt_bit_union_0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_bitcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2847), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_bitge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2641), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_bitlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2644), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_bpchargt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2396), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_bpcharle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2395), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_bytea_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bytea_sortsupport_0)
	return v4
}
func F_gbt_byteaeq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2841), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_bytealt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2842), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_cash_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_cash_sortsupport_0)
	return v4
}
func F_gbt_date_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_date_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_enumkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v7 == v9 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		if v12 == v13 {
			v28 = int32(0)
			return v28
		} else {
			v17 = F_CallerFInfoFunctionCall2(m, int32(3795), l2, int32(0), v12, v13)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v24 = F_CallerFInfoFunctionCall2(m, int32(3795), l2, int32(0), v7, v9)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v28 = v24
			return v28
		}
	}
}
func F_gbt_enumlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3791), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_float4_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13905(m, l0, int32(_a_F_gbt_float4_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_float8_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_float8_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_float8le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func F_gbt_inet_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 float64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+30)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v2)
	v22 = F_convert_network_to_scalar(m, v11, int32(869), v9+int32(7))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v22
		v27 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v27)
		v29 = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v15 + v29
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v15
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v40)+12)))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = F_gbt_num_consistent(m, v9+int32(20), v9+v29, v9+int32(30), v42&v27, int32(_a_F_gbt_inet_consistent_0), v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			m.G0 = v9 + int32(32)
			return v47
		}
	}
}
func F_gbt_int2_distance(m *base.Module, l0 int32) int32 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v7+int32(4), v7+int32(14), v24&int32(1), int32(_a_F_gbt_int2_distance_0), v28)
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
func F_gbt_int2_penalty(m *base.Module, l0 int32) int32 {
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
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+2)))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
	if v16 < v15 {
		v27 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v15), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i32_s(v16), float64(-0.49000000953674316))), float64(0))
	} else {
		v27 = float64(0)
	}
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11))))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9))))
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
func F_gbt_int2_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int2_sortsupport_0)
	return int32(0)
}
func F_gbt_int2ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v5 <= v4)
}
func F_gbt_int2key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5))))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7))))
	if v6 == v8 {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+2)))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+2)))
		if v11 == v12 {
			v29 = int32(0)
			return v29
		} else {
			if base.I32_extend16_s(v12) < base.I32_extend16_s(v11) {
				v19 = int32(1)
			} else {
				v19 = int32(-1)
			}
			return v19
		}
	} else {
		if base.I32_extend16_s(v8) < base.I32_extend16_s(v6) {
			v26 = int32(1)
		} else {
			v26 = int32(-1)
		}
		v29 = v26
		return v29
	}
}
func F_gbt_int4_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_int4_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_int4_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_int4_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int4ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v5 <= v4)
}
func F_gbt_int4key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v6 == v8 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		if v11 == v12 {
			v25 = int32(0)
			return v25
		} else {
			if v12 < v11 {
				v17 = int32(1)
			} else {
				v17 = int32(-1)
			}
			return v17
		}
	} else {
		if v8 < v6 {
			v22 = int32(1)
		} else {
			v22 = int32(-1)
		}
		v25 = v22
		return v25
	}
}
func F_gbt_int4lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 < v5)
}
func F_gbt_int8_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int8_sortsupport_0)
	return int32(0)
}
func F_gbt_intv_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v12 + v7
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v20)+12)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_gbt_num_distance(m, v8+int32(8), v10, v22&int32(1), int32(_a_F_gbt_intv_distance_0), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = F_Float8GetDatum(m, v27)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v31
		}
	}
}
func F_gbt_intv_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_intv_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_intv_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13906(m, l0, int32(_a_F_gbt_intv_union_0), int32(32))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_intveq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2442), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8_consistent(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13912(m, l0, int32(_a_F_gbt_macad8_consistent_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macad_consistent(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13912(m, l0, int32(_a_F_gbt_macad_consistent_0), int32(6))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macaddr_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2266), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_numeric_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 float32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 float32
	_ = v149
	var v150 float64
	_ = v150
	var v152 float32
	_ = v152
	var v155 float32
	_ = v155
	var v162 float32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 float32
	_ = v176
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v17 = v11 + int32(20)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v21 = int32(4)
	v22 = v19 + v21
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = int32(2)
	v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if base.Ui32(v26+v21) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
		v38 = v22 + (v26+int32(3))&int32(2147483644)
	} else {
		v38 = v22
	}
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v38
	v40 = F_gbt_var_key_copy(m, v17)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v40
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_gbt_var_bin_union(m, v11+int32(28), v15, v47, int32(_a_F_gbt_numeric_penalty_0), v49)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v53 = v11 + int32(12)
			v55 = int32(4)
			v56 = v19 + v55
			*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v59 = int32(2)
			v60 = int32(base.Ui32(v58) >> (uint(v59) % 32))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if base.Ui32(v60+v55) < base.Ui32(int32(base.Ui32(v68)>>(uint(v59)%32))) {
				v72 = v56 + (v60+int32(3))&int32(2147483644)
			} else {
				v72 = v56
			}
			*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			v78 = int32(4)
			v79 = v76 + v78
			*(*int32)(unsafe.Add(mBase, uint32(v53))) = v79
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
			v82 = int32(2)
			v83 = int32(base.Ui32(v81) >> (uint(v82) % 32))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
			if base.Ui32(v83+v78) < base.Ui32(int32(base.Ui32(v91)>>(uint(v82)%32))) {
				v95 = v79 + (v83+int32(3))&int32(2147483644)
			} else {
				v95 = v79
			}
			*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v95
			v97 = int32(18)
			v98 = int32(0)
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v103 = F_DirectFunctionCall2Coll(m, v97, v98, v101, v102)
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v105 = F_pg_detoast_datum(m, v103)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					v109 = F_DirectFunctionCall2Coll(m, int32(18), int32(0), v74, v75)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v111 = F_pg_detoast_datum(m, v109)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v113 = F_DirectFunctionCall2Coll(m, v97, v98, v105, v111)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v115 = F_pg_detoast_datum(m, v113)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
									if v117 == int32(_a_F_gbt_numeric_penalty_1) {
										v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+4)))
										if base.B2i32(v121 == int32(_a_F_gbt_numeric_penalty_1)) == int32(0) {
											v162 = float32(1)
											v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
											v176 = base.F32_mul(v162, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v167+int32(1))))
										} else {
											v176 = float32(0)
										}
										*(*float32)(unsafe.Add(mBase, uint32(v13))) = v176
										m.G0 = v11 + int32(32)
										return v13
									} else {
										v127 = F_int64_to_numeric(m, int64(0))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											v129 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = v129
											v133 = F_DirectFunctionCall2Coll(m, int32(2720), v129, v115, v127)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int32(0)
											} else {
												v135 = *(*float32)(unsafe.Add(mBase, uint32(v13)))
												if v133 != 0 {
													*(*float32)(unsafe.Add(mBase, uint32(v13))) = base.F32_add(v135, float32(1.1754944e-38))
													v140 = int32(0)
													v143 = F_DirectFunctionCall2Coll(m, int32(1260), v140, v115, v105)
													mBase = m.M
													v144 = m.ExcPending
													if v144 != 0 {
														return int32(0)
													} else {
														v145 = F_pg_detoast_datum(m, v143)
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return int32(0)
														} else {
															v147 = F_DirectFunctionCall1Coll(m, int32(1492), v140, v145)
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																v149 = *(*float32)(unsafe.Add(mBase, uint32(v13)))
																v150 = *(*float64)(unsafe.Add(mBase, uint32(v147)))
																v152 = base.F32_add(v149, base.F32_demote_f64(v150))
																*(*float32)(unsafe.Add(mBase, uint32(v13))) = v152
																v155 = v152
																if base.F32_gt(v155, float32(0)) == int32(0) {
																} else {
																	v162 = v155
																	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
																	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
																	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
																	v176 = base.F32_mul(v162, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v167+int32(1))))
																	*(*float32)(unsafe.Add(mBase, uint32(v13))) = v176
																}
																m.G0 = v11 + int32(32)
																return v13
															}
														}
													}
												} else {
													v155 = v135
													if base.F32_gt(v155, float32(0)) == int32(0) {
													} else {
														v162 = v155
														v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
														v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
														v176 = base.F32_mul(v162, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v167+int32(1))))
														*(*float32)(unsafe.Add(mBase, uint32(v13))) = v176
													}
													m.G0 = v11 + int32(32)
													return v13
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
}
func F_gbt_numeric_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_numeric_sortsupport_0)
	return v4
}
func F_gbt_oid_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_oid_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_oid_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13910(m, l0, int32(_a_F_gbt_oid_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_oid_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v5) < base.Ui32(v4) {
		v9 = v4 - v5
	} else {
		v9 = v5 - v4
	}
	return base.F64_convert_i32_u(v9)
}
func F_gbt_oid_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_oid_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_textlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2228), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = v21 + v18
	v24 = F_DirectFunctionCall2Coll(m, int32(2705), int32(0), v17+v18, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
		v33 = F_DirectFunctionCall2Coll(m, int32(2705), int32(0), v21, v17)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
			v50 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v28), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v29), float64(86400)), base.F64_div(base.F64_convert_i64_s(v30), float64(1e+06))))
			v51 = float64(0)
			if base.F64_gt(v50, v51) != 0 {
				v54 = v50
			} else {
				v54 = v51
			}
			v65 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v35), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v36), float64(86400)), base.F64_div(base.F64_convert_i64_s(v37), float64(1e+06))))
			v66 = float64(0)
			if base.F64_gt(v65, v66) != 0 {
				v69 = v65
			} else {
				v69 = v66
			}
			v70 = base.F64_add(v54, v69)
			if base.F64_gt(v70, float64(0)) != 0 {
				v75 = F_DirectFunctionCall2Coll(m, int32(2705), int32(0), v23, v21)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v77 = *(*float32)(unsafe.Add(mBase, uint32(v13)))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
					*(*float32)(unsafe.Add(mBase, uint32(v13))) = base.F32_mul(base.F32_add(base.F32_add(v77, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(v70, base.F64_add(v70, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v80), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v84), float64(86400)), base.F64_div(base.F64_convert_i64_s(v88), float64(1e+06)))))))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v102+int32(1))))
					return v13
				}
			} else {
				return v13
			}
		}
	}
}
func F_gbt_timeeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2431), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_timege(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2421), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_timekey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13913(m, l0, l1, l2, int32(8), int32(1429))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_timekey_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1429), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_ts_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_ts_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_ts_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_ts_sortsupport_0)
	return v4
}
func F_gbt_tstz_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v11
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v8+v14, v8, v24&int32(1), int32(_a_F_gbt_tstz_distance_0), v28)
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
			m.G0 = v8 + int32(16)
			return v33
		}
	}
}
func F_gbt_uuid_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v68 float64
	_ = v68
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v261 float64
	_ = v261
	var v338 float64
	_ = v338
	var v346 float64
	_ = v346
	var v349 float64
	_ = v349
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	v10 = float64(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	v32 = int64(56)
	v34 = int64(65280)
	v36 = int64(40)
	v39 = int64(16711680)
	v41 = int64(24)
	v43 = int64(4278190080)
	v45 = int64(8)
	v68 = float64(5.421010862427522e-20)
	v106 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v20<<(uint(v32)%64)|v20&v34<<(uint(v36)%64)|(v20&v39<<(uint(v41)%64)|v20&v43<<(uint(v45)%64))|(int64(base.Ui64(v20)>>(uint(v45)%64))&v43|int64(base.Ui64(v20)>>(uint(v41)%64))&v39|(int64(base.Ui64(v20)>>(uint(v36)%64))&v34|int64(base.Ui64(v20)>>(uint(v32)%64))))), v68), base.F64_convert_i64_u(v19<<(uint(v32)%64)|v19&v34<<(uint(v36)%64)|(v19&v39<<(uint(v41)%64)|v19&v43<<(uint(v45)%64))|(int64(base.Ui64(v19)>>(uint(v45)%64))&v43|int64(base.Ui64(v19)>>(uint(v41)%64))&v39|(int64(base.Ui64(v19)>>(uint(v36)%64))&v34|int64(base.Ui64(v19)>>(uint(v32)%64))))))
	v107 = float64(0.49000000953674316)
	v183 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v24<<(uint(v32)%64)|v24&v34<<(uint(v36)%64)|(v24&v39<<(uint(v41)%64)|v24&v43<<(uint(v45)%64))|(int64(base.Ui64(v24)>>(uint(v45)%64))&v43|int64(base.Ui64(v24)>>(uint(v41)%64))&v39|(int64(base.Ui64(v24)>>(uint(v36)%64))&v34|int64(base.Ui64(v24)>>(uint(v32)%64))))), v68), base.F64_convert_i64_u(v23<<(uint(v32)%64)|v23&v34<<(uint(v36)%64)|(v23&v39<<(uint(v41)%64)|v23&v43<<(uint(v45)%64))|(int64(base.Ui64(v23)>>(uint(v45)%64))&v43|int64(base.Ui64(v23)>>(uint(v41)%64))&v39|(int64(base.Ui64(v23)>>(uint(v36)%64))&v34|int64(base.Ui64(v23)>>(uint(v32)%64))))))
	v184 = float64(-0.49000000953674316)
	v261 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v26<<(uint(v32)%64)|v26&v34<<(uint(v36)%64)|(v26&v39<<(uint(v41)%64)|v26&v43<<(uint(v45)%64))|(int64(base.Ui64(v26)>>(uint(v45)%64))&v43|int64(base.Ui64(v26)>>(uint(v41)%64))&v39|(int64(base.Ui64(v26)>>(uint(v36)%64))&v34|int64(base.Ui64(v26)>>(uint(v32)%64))))), v68), base.F64_convert_i64_u(v25<<(uint(v32)%64)|v25&v34<<(uint(v36)%64)|(v25&v39<<(uint(v41)%64)|v25&v43<<(uint(v45)%64))|(int64(base.Ui64(v25)>>(uint(v45)%64))&v43|int64(base.Ui64(v25)>>(uint(v41)%64))&v39|(int64(base.Ui64(v25)>>(uint(v36)%64))&v34|int64(base.Ui64(v25)>>(uint(v32)%64))))))
	v338 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v28<<(uint(v32)%64)|v28&v34<<(uint(v36)%64)|(v28&v39<<(uint(v41)%64)|v28&v43<<(uint(v45)%64))|(int64(base.Ui64(v28)>>(uint(v45)%64))&v43|int64(base.Ui64(v28)>>(uint(v41)%64))&v39|(int64(base.Ui64(v28)>>(uint(v36)%64))&v34|int64(base.Ui64(v28)>>(uint(v32)%64))))), v68), base.F64_convert_i64_u(v27<<(uint(v32)%64)|v27&v34<<(uint(v36)%64)|(v27&v39<<(uint(v41)%64)|v27&v43<<(uint(v45)%64))|(int64(base.Ui64(v27)>>(uint(v45)%64))&v43|int64(base.Ui64(v27)>>(uint(v41)%64))&v39|(int64(base.Ui64(v27)>>(uint(v36)%64))&v34|int64(base.Ui64(v27)>>(uint(v32)%64))))))
	if base.F64_gt(v261, v338) != 0 {
		v346 = base.F64_add(base.F64_add(base.F64_mul(v261, v107), base.F64_mul(v338, v184)), v10)
	} else {
		v346 = v10
	}
	if base.F64_lt(v183, v106) != 0 {
		v349 = base.F64_add(base.F64_add(base.F64_mul(v106, v107), base.F64_mul(v183, v184)), v346)
	} else {
		v349 = v346
	}
	if base.F64_gt(v349, float64(0)) != 0 {
		v363 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+52))
		v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
		*(*float32)(unsafe.Add(mBase, uint32(v29))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v349, base.F64_add(base.F64_add(base.F64_mul(v338, float64(0.49000000953674316)), base.F64_mul(v106, float64(-0.49000000953674316))), v349))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v365+int32(1))))
	} else {
	}
	return v29
}
func F_gbt_uuid_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_uuid_sortsupport_0)
	return v4
}
func F_gbt_uuideq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	return base.B2i32(v4^v5|(v7^v8) == int64(0))
}
func F_gbt_var_bin_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	v11 = int32(4)
	v12 = l1 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = int32(2)
	v15 = int32(base.Ui32(v13) >> (uint(v14) % 32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v15+v11) < base.Ui32(int32(base.Ui32(v23)>>(uint(v14)%32))) {
		v27 = v12 + (v15+int32(3))&int32(2147483644)
	} else {
		v27 = v12
	}
	if v12 != v27 {
		v53 = v12
		v54 = v27
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v56 == int32(0) {
			v89 = v53
			v91 = v54
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
			v95 = int32(2)
			v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
			v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
			v103 = (v99 + int32(3)) & int32(2147483644)
			v106 = v96 + v103 + int32(4)
			v107 = F_palloc0(m, v106)
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return
			} else {
				v110 = v107 + int32(4)
				if v99 != 0 {
					base.MemoryCopy(m, v110, v89, v99)
				} else {
				}
				if v96 != 0 {
					base.MemoryCopy(m, v110+v103, v91, v96)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
				return
			}
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v61 = v56 + int32(4)
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
			v64 = m.T0[v63].(func(*base.Module, int32, int32, int32, int32) int32)(m, v61, v53, l2, l4)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				v66 = int32(2)
				v67 = int32(base.Ui32(v62) >> (uint(v66) % 32))
				if base.Ui32(v67+int32(4)) < base.Ui32(int32(base.Ui32(v59)>>(uint(v66)%32))) {
					v78 = v61 + (v67+int32(3))&int32(2147483644)
				} else {
					v78 = v61
				}
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
				v80 = m.T0[v79].(func(*base.Module, int32, int32, int32, int32) int32)(m, v78, v54, l2, l4)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					if int32(0) < v64 {
						if v80 < int32(0) {
							v86 = v54
						} else {
							v86 = v78
						}
						v89 = v53
						v91 = v86
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						v95 = int32(2)
						v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
						v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
						v103 = (v99 + int32(3)) & int32(2147483644)
						v106 = v96 + v103 + int32(4)
						v107 = F_palloc0(m, v106)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							v110 = v107 + int32(4)
							if v99 != 0 {
								base.MemoryCopy(m, v110, v89, v99)
							} else {
							}
							if v96 != 0 {
								base.MemoryCopy(m, v110+v103, v91, v96)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
							return
						}
					} else {
						if int32(0) <= v80 {
							return
						} else {
							v89 = v61
							v91 = v54
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
							v95 = int32(2)
							v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
							v103 = (v99 + int32(3)) & int32(2147483644)
							v106 = v96 + v103 + int32(4)
							v107 = F_palloc0(m, v106)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v110 = v107 + int32(4)
								if v99 != 0 {
									base.MemoryCopy(m, v110, v89, v99)
								} else {
								}
								if v96 != 0 {
									base.MemoryCopy(m, v110+v103, v91, v96)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
								return
							}
						}
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
		if v29 == int32(0) {
			v53 = v12
			v54 = v12
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v56 == int32(0) {
				v89 = v53
				v91 = v54
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v95 = int32(2)
				v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
				v103 = (v99 + int32(3)) & int32(2147483644)
				v106 = v96 + v103 + int32(4)
				v107 = F_palloc0(m, v106)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					v110 = v107 + int32(4)
					if v99 != 0 {
						base.MemoryCopy(m, v110, v89, v99)
					} else {
					}
					if v96 != 0 {
						base.MemoryCopy(m, v110+v103, v91, v96)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
					return
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v61 = v56 + int32(4)
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
				v64 = m.T0[v63].(func(*base.Module, int32, int32, int32, int32) int32)(m, v61, v53, l2, l4)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = int32(2)
					v67 = int32(base.Ui32(v62) >> (uint(v66) % 32))
					if base.Ui32(v67+int32(4)) < base.Ui32(int32(base.Ui32(v59)>>(uint(v66)%32))) {
						v78 = v61 + (v67+int32(3))&int32(2147483644)
					} else {
						v78 = v61
					}
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
					v80 = m.T0[v79].(func(*base.Module, int32, int32, int32, int32) int32)(m, v78, v54, l2, l4)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						if int32(0) < v64 {
							if v80 < int32(0) {
								v86 = v54
							} else {
								v86 = v78
							}
							v89 = v53
							v91 = v86
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
							v95 = int32(2)
							v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
							v103 = (v99 + int32(3)) & int32(2147483644)
							v106 = v96 + v103 + int32(4)
							v107 = F_palloc0(m, v106)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v110 = v107 + int32(4)
								if v99 != 0 {
									base.MemoryCopy(m, v110, v89, v99)
								} else {
								}
								if v96 != 0 {
									base.MemoryCopy(m, v110+v103, v91, v96)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
								return
							}
						} else {
							if int32(0) <= v80 {
								return
							} else {
								v89 = v61
								v91 = v54
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								v95 = int32(2)
								v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
								v103 = (v99 + int32(3)) & int32(2147483644)
								v106 = v96 + v103 + int32(4)
								v107 = F_palloc0(m, v106)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = v107 + int32(4)
									if v99 != 0 {
										base.MemoryCopy(m, v110, v89, v99)
									} else {
									}
									if v96 != 0 {
										base.MemoryCopy(m, v110+v103, v91, v96)
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
									return
								}
							}
						}
					}
				}
			}
		} else {
			v32 = m.T0[v29].(func(*base.Module, int32, int32) int32)(m, l1, l4)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if l1 == v32 {
					v53 = v12
					v54 = v12
				} else {
					v35 = int32(4)
					v36 = v32 + v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v38 = int32(2)
					v39 = int32(base.Ui32(v37) >> (uint(v38) % 32))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					if base.Ui32(v39+v35) < base.Ui32(int32(base.Ui32(v47)>>(uint(v38)%32))) {
						v51 = v36 + (v39+int32(3))&int32(2147483644)
					} else {
						v51 = v36
					}
					v53 = v36
					v54 = v51
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v56 == int32(0) {
					v89 = v53
					v91 = v54
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
					v95 = int32(2)
					v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
					v103 = (v99 + int32(3)) & int32(2147483644)
					v106 = v96 + v103 + int32(4)
					v107 = F_palloc0(m, v106)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return
					} else {
						v110 = v107 + int32(4)
						if v99 != 0 {
							base.MemoryCopy(m, v110, v89, v99)
						} else {
						}
						if v96 != 0 {
							base.MemoryCopy(m, v110+v103, v91, v96)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
						return
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v61 = v56 + int32(4)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
					v64 = m.T0[v63].(func(*base.Module, int32, int32, int32, int32) int32)(m, v61, v53, l2, l4)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = int32(2)
						v67 = int32(base.Ui32(v62) >> (uint(v66) % 32))
						if base.Ui32(v67+int32(4)) < base.Ui32(int32(base.Ui32(v59)>>(uint(v66)%32))) {
							v78 = v61 + (v67+int32(3))&int32(2147483644)
						} else {
							v78 = v61
						}
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
						v80 = m.T0[v79].(func(*base.Module, int32, int32, int32, int32) int32)(m, v78, v54, l2, l4)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							if int32(0) < v64 {
								if v80 < int32(0) {
									v86 = v54
								} else {
									v86 = v78
								}
								v89 = v53
								v91 = v86
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								v95 = int32(2)
								v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
								v103 = (v99 + int32(3)) & int32(2147483644)
								v106 = v96 + v103 + int32(4)
								v107 = F_palloc0(m, v106)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v110 = v107 + int32(4)
									if v99 != 0 {
										base.MemoryCopy(m, v110, v89, v99)
									} else {
									}
									if v96 != 0 {
										base.MemoryCopy(m, v110+v103, v91, v96)
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
									return
								}
							} else {
								if int32(0) <= v80 {
									return
								} else {
									v89 = v61
									v91 = v54
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									v95 = int32(2)
									v96 = int32(base.Ui32(v94) >> (uint(v95) % 32))
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
									v99 = int32(base.Ui32(v97) >> (uint(v95) % 32))
									v103 = (v99 + int32(3)) & int32(2147483644)
									v106 = v96 + v103 + int32(4)
									v107 = F_palloc0(m, v106)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v110 = v107 + int32(4)
										if v99 != 0 {
											base.MemoryCopy(m, v110, v89, v99)
										} else {
										}
										if v96 != 0 {
											base.MemoryCopy(m, v110+v103, v91, v96)
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
										return
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
func F_gbt_var_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(1)
	v25 = (v21 - v22) & int32(_a_F_gbt_var_picksplit_0)
	v27 = v25 + v22
	v30 = F_palloc(m, v27<<(uint(int32(3))%32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = v25<<(uint(int32(1))%32) + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
	v41 = F_palloc(m, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v41
	v44 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+20)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = v44
	v49 = l1 + int32(24)
	v51 = l1 + int32(8)
	v54 = F_palloc(m, v27<<(uint(int32(2))%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v21&int32(_a_F_gbt_var_picksplit_0) != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v215 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L7:
	;
	v63 = int32(1)
	v74 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v191 = int32(8)
	F_qsort_arg(m, v30+v191, v25, v191, int32(_a_F_gbt_var_picksplit_1), v19+int32(4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L33
	}
L10:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v63<<(uint(int32(4))%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v85 = int32(base.Ui32(v83) >> (uint(int32(2)) % 32))
	if v85 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v126 = int32(8)
	F_qsort_arg(m, v30+v126, v25, v126, int32(_a_F_gbt_var_picksplit_1), v19+int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v63<<(uint(int32(3))%32)))) = v63
	v121 = (v63 + int32(1)) & int32(_a_F_gbt_var_picksplit_0)
	if base.Ui32(v121) <= base.Ui32(v25) {
		v63 = v121
		v74 = v112
		goto L10
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v63<<(uint(int32(3))%32))+4)) = v82
	v112 = v74
	goto L12
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if base.Ui32(v85+int32(4)) < base.Ui32(int32(base.Ui32(v88)>>(uint(int32(2))%32))) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v92 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v93 = m.T0[v92].(func(*base.Module, int32, int32) int32)(m, v82, l4)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v95 = v82
	goto L20
L20:
	;
	v98 = v54 + v74<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v30+v63<<(uint(int32(3))%32))+4)) = v95
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v112 = v74 + base.B2i32(v104 != v82)
	goto L12
L21:
	;
	v95 = v93
	goto L20
L22:
	;
	goto L11
L23:
	;
	v134 = int32(1)
	v137 = v134
	goto L24
L24:
	;
	v155 = v30 + v137<<(uint(int32(3))%32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if base.Ui32(v137) <= base.Ui32(int32(base.Ui32(v25)>>(uint(v134)%32))) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L6
L26:
	;
	v186 = (v137 + int32(1)) & int32(_a_F_gbt_var_picksplit_0)
	if base.Ui32(v186) <= base.Ui32(v25) {
		v137 = v186
		goto L24
	} else {
		goto L32
	}
L27:
	;
	F_gbt_var_bin_union(m, v51, v156, l2, l3, l4)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_gbt_var_bin_union(m, v49, v156, l2, l3, l4)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v162 = int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*uint16)(unsafe.Add(mBase, uint32(v160+v161<<(uint(v162)%32)))) = uint16(v165)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v161 + v162
	goto L26
L31:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v174 = int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*uint16)(unsafe.Add(mBase, uint32(v172+v173<<(uint(v174)%32)))) = uint16(v177)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v173 + v174
	goto L26
L32:
	;
	goto L25
L33:
	;
	goto L6
L34:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v219 = F_gbt_var_node_cp_len(m, v218, l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v19 + int32(16)
	return l1
L37:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v222 = F_gbt_var_node_cp_len(m, v221, l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v225 = int32(4)
	v226 = v224 + v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = int32(2)
	v229 = int32(base.Ui32(v227) >> (uint(v228) % 32))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if base.Ui32(v229+v225) < base.Ui32(int32(base.Ui32(v237)>>(uint(v228)%32))) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v241 = v226 + (v229+int32(3))&int32(2147483644)
	goto L41
L40:
	;
	v241 = v226
	goto L41
L41:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v246 = int32(base.Ui32(v242)>>(uint(int32(2))%32)) - int32(4)
	if v222 < v219 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v248 = v219
	goto L44
L43:
	;
	v248 = v222
	goto L44
L44:
	;
	v250 = v248 + int32(2)
	if v246 < v250 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v252 = v246
	goto L47
L46:
	;
	v252 = v250
	goto L47
L47:
	;
	v254 = v229 - int32(4)
	if v254 < v250 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v256 = v254
	goto L50
L49:
	;
	v256 = v250
	goto L50
L50:
	;
	v260 = (v256 + int32(7)) & int32(-4)
	v263 = v252 + v260 + int32(8)
	v264 = F_palloc0(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264))) = v263 << (uint(int32(2)) % 32)
	v269 = int32(4)
	v270 = v264 + v269
	v272 = v256 + v269
	if v272 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	base.MemoryCopy(m, v270, v226, v272)
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v272 << (uint(int32(2)) % 32)
	v277 = v270 + v260
	v279 = v252 + int32(4)
	if v279 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v277, v241, v279)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v281 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v279 << (uint(v281) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v285 = int32(4)
	v286 = v284 + v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v289 = int32(base.Ui32(v287) >> (uint(v281) % 32))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if base.Ui32(v289+v285) < base.Ui32(int32(base.Ui32(v297)>>(uint(v281)%32))) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v301 = v286 + (v289+int32(3))&int32(2147483644)
	goto L60
L59:
	;
	v301 = v286
	goto L60
L60:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v306 = int32(base.Ui32(v302)>>(uint(int32(2))%32)) - int32(4)
	if v306 < v250 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v308 = v306
	goto L63
L62:
	;
	v308 = v250
	goto L63
L63:
	;
	v310 = v289 - int32(4)
	if v310 < v250 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v312 = v310
	goto L66
L65:
	;
	v312 = v250
	goto L66
L66:
	;
	v316 = (v312 + int32(7)) & int32(-4)
	v319 = v308 + v316 + int32(8)
	v320 = F_palloc0(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v319 << (uint(int32(2)) % 32)
	v325 = int32(4)
	v326 = v320 + v325
	v328 = v312 + v325
	if v328 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v326, v286, v328)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v328 << (uint(int32(2)) % 32)
	v333 = v326 + v316
	v335 = v308 + int32(4)
	if v335 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v333, v301, v335)
	goto L73
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v335 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v320
	goto L36
}
