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
	v8 = F_gbt_var_penalty(m, v2, v3, v4, v5, int32(4429344), v7)
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
	v7 = F_gbt_var_union(m, v2, v3, v4, int32(4429344), v6)
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2863), int32(0), l0, l1)
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2657), int32(0), l0, l1)
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2660), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2412), l2, l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2411), l2, l0, l1)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7376)
	return v3
}
func F_gbt_byteaeq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2857), int32(0), l0, l1)
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2858), int32(0), l0, l1)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7383)
	return v3
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
	v4 = F_gbt_num_compress(m, v2, int32(4429504))
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
			v17 = F_CallerFInfoFunctionCall2(m, int32(3811), l2, int32(0), v12, v13)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v24 = F_CallerFInfoFunctionCall2(m, int32(3811), l2, int32(0), v7, v9)
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3807), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_float4_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4429584), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
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
	v4 = F_gbt_num_fetch(m, v2, int32(4429624))
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
		v47 = F_gbt_num_consistent(m, v9+int32(20), v9+v29, v9+int32(30), v42&v27, int32(4429728), v46)
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
	v29 = F_gbt_num_distance(m, v7+int32(4), v7+int32(14), v24&int32(1), int32(4429768), v28)
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
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7429)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4429808), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_int4_union(m *base.Module, l0 int32) int32 {
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
	v6 = F_palloc(m, int32(8))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4429808), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
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
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7445)
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
	v27 = F_gbt_num_distance(m, v8+int32(8), v10, v22&int32(1), int32(4429888), v26)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4429888), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_intv_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4429888), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_intveq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2458), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+v27)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v8+int32(4), v10, v8+int32(14), v29&int32(1), int32(4429968), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v34
	}
}
func F_gbt_macad_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14 + int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+v27)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v8+int32(4), v10, v8+int32(14), v29&int32(1), int32(4429928), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v34
	}
}
func F_gbt_macaddr_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2282), int32(0), l0, l1)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 float32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 float32
	_ = v152
	var v153 float64
	_ = v153
	var v155 float32
	_ = v155
	var v158 float32
	_ = v158
	var v165 float32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 float32
	_ = v179
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v18 = v12 + int32(20)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = int32(4)
	v23 = v20 + v22
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui32(v27+v22) < base.Ui32(int32(base.Ui32(v35)>>(uint(v26)%32))) {
		v39 = v23 + (v27+int32(3))&int32(2147483644)
	} else {
		v39 = v23
	}
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v39
	v43 = F_gbt_var_key_copy(m, v12+int32(20))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v43
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_gbt_var_bin_union(m, v12+int32(28), v16, v50, int32(4430008), v52)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v56 = v12 + int32(12)
			v58 = int32(4)
			v59 = v20 + v58
			*(*int32)(unsafe.Add(mBase, uint32(v56))) = v59
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v62 = int32(2)
			v63 = int32(base.Ui32(v61) >> (uint(v62) % 32))
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if base.Ui32(v63+v58) < base.Ui32(int32(base.Ui32(v71)>>(uint(v62)%32))) {
				v75 = v59 + (v63+int32(3))&int32(2147483644)
			} else {
				v75 = v59
			}
			*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v75
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v80 = v12 + int32(12)
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			v83 = int32(4)
			v84 = v81 + v83
			*(*int32)(unsafe.Add(mBase, uint32(v80))) = v84
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
			v87 = int32(2)
			v88 = int32(base.Ui32(v86) >> (uint(v87) % 32))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
			if base.Ui32(v88+v83) < base.Ui32(int32(base.Ui32(v96)>>(uint(v87)%32))) {
				v100 = v84 + (v88+int32(3))&int32(2147483644)
			} else {
				v100 = v84
			}
			*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v100
			v102 = int32(18)
			v103 = int32(0)
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v107 = F_DirectFunctionCall2Coll(m, v102, v103, v105, v106)
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				v109 = F_pg_detoast_datum(m, v107)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v112 = F_DirectFunctionCall2Coll(m, v102, int32(0), v77, v78)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						v114 = F_pg_detoast_datum(m, v112)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v116 = F_DirectFunctionCall2Coll(m, v102, v103, v109, v114)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v118 = F_pg_detoast_datum(m, v116)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+4)))
									if v120 == int32(49152) {
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
										if base.B2i32(v124 == int32(49152)) == int32(0) {
											v165 = float32(1)
											v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
											v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
											v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
											v179 = base.F32_mul(v165, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v170+int32(1))))
										} else {
											v179 = float32(0)
										}
										*(*float32)(unsafe.Add(mBase, uint32(v14))) = v179
										m.G0 = v12 + int32(32)
										return v14
									} else {
										v130 = F_int64_to_numeric(m, int64(0))
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v14))) = v132
											v136 = F_DirectFunctionCall2Coll(m, int32(2736), v132, v118, v130)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												v138 = *(*float32)(unsafe.Add(mBase, uint32(v14)))
												if v136 != 0 {
													*(*float32)(unsafe.Add(mBase, uint32(v14))) = base.F32_add(v138, float32(1.1754944e-38))
													v143 = int32(0)
													v146 = F_DirectFunctionCall2Coll(m, int32(1276), v143, v118, v109)
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return int32(0)
													} else {
														v148 = F_pg_detoast_datum(m, v146)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															v150 = F_DirectFunctionCall1Coll(m, int32(1508), v143, v148)
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																v152 = *(*float32)(unsafe.Add(mBase, uint32(v14)))
																v153 = *(*float64)(unsafe.Add(mBase, uint32(v150)))
																v155 = base.F32_add(v152, base.F32_demote_f64(v153))
																*(*float32)(unsafe.Add(mBase, uint32(v14))) = v155
																v158 = v155
																if base.F32_gt(v158, float32(0)) == int32(0) {
																} else {
																	v165 = v158
																	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
																	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
																	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
																	v179 = base.F32_mul(v165, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v170+int32(1))))
																	*(*float32)(unsafe.Add(mBase, uint32(v14))) = v179
																}
																m.G0 = v12 + int32(32)
																return v14
															}
														}
													}
												} else {
													v158 = v138
													if base.F32_gt(v158, float32(0)) == int32(0) {
													} else {
														v165 = v158
														v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
														v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
														v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
														v179 = base.F32_mul(v165, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v170+int32(1))))
														*(*float32)(unsafe.Add(mBase, uint32(v14))) = v179
													}
													m.G0 = v12 + int32(32)
													return v14
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7475)
	return v3
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
	v4 = F_gbt_num_compress(m, v2, int32(4430048))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_oid_consistent(m *base.Module, l0 int32) int32 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10
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
	v34 = F_gbt_num_consistent(m, v7, v7+int32(12), v7+int32(10), v29&int32(1), int32(4430048), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v34
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
	v4 = F_gbt_num_fetch(m, v2, int32(4430048))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2244), l2, l0, l1)
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
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = int32(2721)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = v21 + v18
	v24 = F_DirectFunctionCall2Coll(m, v14, int32(0), v17+v18, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
		v32 = F_DirectFunctionCall2Coll(m, v14, int32(0), v21, v17)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
			v49 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v28), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v29), float64(86400)), base.F64_div(base.F64_convert_i64_s(v30), float64(1e+06))))
			v50 = float64(0)
			if base.F64_gt(v49, v50) != 0 {
				v53 = v49
			} else {
				v53 = v50
			}
			v64 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v34), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v35), float64(86400)), base.F64_div(base.F64_convert_i64_s(v36), float64(1e+06))))
			v65 = float64(0)
			if base.F64_gt(v64, v65) != 0 {
				v68 = v64
			} else {
				v68 = v65
			}
			v69 = base.F64_add(v53, v68)
			if base.F64_gt(v69, float64(0)) != 0 {
				v74 = F_DirectFunctionCall2Coll(m, int32(2721), int32(0), v23, v21)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = *(*float32)(unsafe.Add(mBase, uint32(v13)))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
					v87 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+52))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
					*(*float32)(unsafe.Add(mBase, uint32(v13))) = base.F32_mul(base.F32_add(base.F32_add(v76, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(v69, base.F64_add(v69, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v79), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v83), float64(86400)), base.F64_div(base.F64_convert_i64_s(v87), float64(1e+06)))))))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v101+int32(1))))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2447), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2437), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_timekey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(1445), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v20 = v8
			return v20
		} else {
			v14 = int32(8)
			v18 = F_DirectFunctionCall2Coll(m, int32(1445), int32(0), v6+v14, v7+v14)
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
func F_gbt_timekey_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1445), int32(0), l0, l1)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4430208), v5)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7512)
	return v3
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
	v29 = F_gbt_num_distance(m, v8+v14, v8, v24&int32(1), int32(4430208), v28)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7521)
	return v3
}
func F_gbt_uuideq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(v66 == int32(0))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v35 = l0
		v36 = l1
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = l0
	v13 = l1
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
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
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
			v96 = int32(2)
			v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
			v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
			v104 = (v100 + int32(3)) & int32(2147483644)
			v107 = v97 + v104 + int32(4)
			v108 = F_palloc0(m, v107)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return
			} else {
				v111 = v108 + int32(4)
				if v100 != 0 {
					v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
					mBase = m.M
					v113 = v112
				} else {
					v113 = v111
				}
				if v97 != 0 {
					v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						v96 = int32(2)
						v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
						v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
						v104 = (v100 + int32(3)) & int32(2147483644)
						v107 = v97 + v104 + int32(4)
						v108 = F_palloc0(m, v107)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							v111 = v108 + int32(4)
							if v100 != 0 {
								v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
								mBase = m.M
								v113 = v112
							} else {
								v113 = v111
							}
							if v97 != 0 {
								v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
							return
						}
					} else {
						if int32(0) <= v80 {
							return
						} else {
							v89 = v61
							v91 = v54
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
							v96 = int32(2)
							v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
							v104 = (v100 + int32(3)) & int32(2147483644)
							v107 = v97 + v104 + int32(4)
							v108 = F_palloc0(m, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v111 = v108 + int32(4)
								if v100 != 0 {
									v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
									mBase = m.M
									v113 = v112
								} else {
									v113 = v111
								}
								if v97 != 0 {
									v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
									mBase = m.M
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v96 = int32(2)
				v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
				v104 = (v100 + int32(3)) & int32(2147483644)
				v107 = v97 + v104 + int32(4)
				v108 = F_palloc0(m, v107)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					v111 = v108 + int32(4)
					if v100 != 0 {
						v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
						mBase = m.M
						v113 = v112
					} else {
						v113 = v111
					}
					if v97 != 0 {
						v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
							v96 = int32(2)
							v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
							v104 = (v100 + int32(3)) & int32(2147483644)
							v107 = v97 + v104 + int32(4)
							v108 = F_palloc0(m, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v111 = v108 + int32(4)
								if v100 != 0 {
									v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
									mBase = m.M
									v113 = v112
								} else {
									v113 = v111
								}
								if v97 != 0 {
									v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
									mBase = m.M
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
								return
							}
						} else {
							if int32(0) <= v80 {
								return
							} else {
								v89 = v61
								v91 = v54
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								v96 = int32(2)
								v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
								v104 = (v100 + int32(3)) & int32(2147483644)
								v107 = v97 + v104 + int32(4)
								v108 = F_palloc0(m, v107)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v111 = v108 + int32(4)
									if v100 != 0 {
										v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
										mBase = m.M
										v113 = v112
									} else {
										v113 = v111
									}
									if v97 != 0 {
										v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
										mBase = m.M
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
				if v32 == l1 {
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
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
					v96 = int32(2)
					v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
					v104 = (v100 + int32(3)) & int32(2147483644)
					v107 = v97 + v104 + int32(4)
					v108 = F_palloc0(m, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v111 = v108 + int32(4)
						if v100 != 0 {
							v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
							mBase = m.M
							v113 = v112
						} else {
							v113 = v111
						}
						if v97 != 0 {
							v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								v96 = int32(2)
								v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
								v104 = (v100 + int32(3)) & int32(2147483644)
								v107 = v97 + v104 + int32(4)
								v108 = F_palloc0(m, v107)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v111 = v108 + int32(4)
									if v100 != 0 {
										v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
										mBase = m.M
										v113 = v112
									} else {
										v113 = v111
									}
									if v97 != 0 {
										v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
										mBase = m.M
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
									return
								}
							} else {
								if int32(0) <= v80 {
									return
								} else {
									v89 = v61
									v91 = v54
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
									v96 = int32(2)
									v97 = int32(base.Ui32(v95) >> (uint(v96) % 32))
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
									v100 = int32(base.Ui32(v98) >> (uint(v96) % 32))
									v104 = (v100 + int32(3)) & int32(2147483644)
									v107 = v97 + v104 + int32(4)
									v108 = F_palloc0(m, v107)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v111 = v108 + int32(4)
										if v100 != 0 {
											v112 = F__emscripten_memcpy_bulkmem(m, v111, v89, v100)
											mBase = m.M
											v113 = v112
										} else {
											v113 = v111
										}
										if v97 != 0 {
											v115 = F__emscripten_memcpy_bulkmem(m, v113+v104, v91, v97)
											mBase = m.M
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v108))) = v107 << (uint(int32(2)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v108
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
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v218 int32
	_ = v218
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
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
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(1)
	v25 = (v21 - v22) & int32(65535)
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
	if v21&int32(65535) != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v218 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L7:
	;
	v64 = int32(1)
	v73 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v194 = int32(8)
	F_qsort_arg(m, v30+v194, v25, v194, int32(7520), v19+int32(4))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L33
	}
L10:
	;
	v80 = int32(4)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v64<<(uint(v80)%32))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if base.Ui32(v80) <= base.Ui32(v84) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v129 = int32(8)
	F_qsort_arg(m, v30+v129, v25, v129, int32(7520), v19+int32(4))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v64<<(uint(int32(3))%32)))) = v64
	v124 = (v64 + int32(1)) & int32(65535)
	if base.Ui32(v124) <= base.Ui32(v25) {
		v64 = v124
		v73 = v115
		goto L10
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v64<<(uint(int32(3))%32))+4)) = v83
	v115 = v73
	goto L12
L14:
	;
	v87 = int32(2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if base.Ui32(int32(base.Ui32(v84)>>(uint(v87)%32))+int32(4)) < base.Ui32(int32(base.Ui32(v91)>>(uint(v87)%32))) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v95 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v96 = m.T0[v95].(func(*base.Module, int32, int32) int32)(m, v83, l4)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v98 = v83
	goto L20
L20:
	;
	v101 = v54 + v73<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v30+v64<<(uint(int32(3))%32))+4)) = v98
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v115 = v73 + base.B2i32(v107 != v83)
	goto L12
L21:
	;
	v98 = v96
	goto L20
L22:
	;
	goto L11
L23:
	;
	v137 = int32(1)
	v140 = v137
	goto L24
L24:
	;
	v158 = v30 + v140<<(uint(int32(3))%32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if base.Ui32(v140) <= base.Ui32(int32(base.Ui32(v25)>>(uint(v137)%32))) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L6
L26:
	;
	v189 = (v140 + int32(1)) & int32(65535)
	if base.Ui32(v189) <= base.Ui32(v25) {
		v140 = v189
		goto L24
	} else {
		goto L32
	}
L27:
	;
	F_gbt_var_bin_union(m, v51, v159, l2, l3, l4)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_gbt_var_bin_union(m, v49, v159, l2, l3, l4)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v165 = int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*uint16)(unsafe.Add(mBase, uint32(v163+v164<<(uint(v165)%32)))) = uint16(v168)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v164 + v165
	goto L26
L31:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v177 = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*uint16)(unsafe.Add(mBase, uint32(v175+v176<<(uint(v177)%32)))) = uint16(v180)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v176 + v177
	goto L26
L32:
	;
	goto L25
L33:
	;
	goto L6
L34:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v222 = F_gbt_var_node_cp_len(m, v221, l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v225 = F_gbt_var_node_cp_len(m, v224, l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v228 = int32(4)
	v229 = v227 + v228
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v231 = int32(2)
	v232 = int32(base.Ui32(v230) >> (uint(v231) % 32))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if base.Ui32(v232+v228) < base.Ui32(int32(base.Ui32(v240)>>(uint(v231)%32))) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v244 = v229 + (v232+int32(3))&int32(2147483644)
	goto L41
L40:
	;
	v244 = v229
	goto L41
L41:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v249 = int32(base.Ui32(v245)>>(uint(int32(2))%32)) - int32(4)
	if v225 < v222 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v251 = v222
	goto L44
L43:
	;
	v251 = v225
	goto L44
L44:
	;
	v253 = v251 + int32(2)
	if v249 < v253 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v255 = v249
	goto L47
L46:
	;
	v255 = v253
	goto L47
L47:
	;
	v257 = v232 - int32(4)
	if v257 < v253 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v259 = v257
	goto L50
L49:
	;
	v259 = v253
	goto L50
L50:
	;
	v263 = (v259 + int32(7)) & int32(-4)
	v266 = v255 + v263 + int32(8)
	v267 = F_palloc0(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v266 << (uint(int32(2)) % 32)
	v272 = int32(4)
	v273 = v267 + v272
	v275 = v259 + v272
	if v275 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+4)) = v275 << (uint(int32(2)) % 32)
	v281 = v277 + v263
	v283 = v255 + int32(4)
	if v283 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v276 = F__emscripten_memcpy_bulkmem(m, v273, v229, v275)
	mBase = m.M
	v277 = v276
	goto L55
L54:
	;
	v277 = v273
	goto L55
L55:
	;
	goto L52
L56:
	;
	v286 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v283 << (uint(v286) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v290 = int32(4)
	v291 = v289 + v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v294 = int32(base.Ui32(v292) >> (uint(v286) % 32))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	if base.Ui32(v294+v290) < base.Ui32(int32(base.Ui32(v302)>>(uint(v286)%32))) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v284 = F__emscripten_memcpy_bulkmem(m, v281, v244, v283)
	mBase = m.M
	v285 = v284
	goto L59
L58:
	;
	v285 = v281
	goto L59
L59:
	;
	goto L56
L60:
	;
	v306 = v291 + (v294+int32(3))&int32(2147483644)
	goto L62
L61:
	;
	v306 = v291
	goto L62
L62:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v311 = int32(base.Ui32(v307)>>(uint(int32(2))%32)) - int32(4)
	if v311 < v253 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v313 = v311
	goto L65
L64:
	;
	v313 = v253
	goto L65
L65:
	;
	v315 = v294 - int32(4)
	if v315 < v253 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v317 = v315
	goto L68
L67:
	;
	v317 = v253
	goto L68
L68:
	;
	v321 = (v317 + int32(7)) & int32(-4)
	v324 = v313 + v321 + int32(8)
	v325 = F_palloc0(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v324 << (uint(int32(2)) % 32)
	v330 = int32(4)
	v331 = v325 + v330
	v333 = v317 + v330
	if v333 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325)+4)) = v333 << (uint(int32(2)) % 32)
	v339 = v335 + v321
	v341 = v313 + int32(4)
	if v341 != 0 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v334 = F__emscripten_memcpy_bulkmem(m, v331, v291, v333)
	mBase = m.M
	v335 = v334
	goto L73
L72:
	;
	v335 = v331
	goto L73
L73:
	;
	goto L70
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = v341 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v325
	goto L36
L75:
	;
	v342 = F__emscripten_memcpy_bulkmem(m, v339, v306, v341)
	mBase = m.M
	v343 = v342
	goto L77
L76:
	;
	v343 = v339
	goto L77
L77:
	;
	goto L74
}
