package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_l2n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = v9 + int32(8)
	v14 = int32(4)
	v15 = l0 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v18 = int32(2)
	v19 = int32(base.Ui32(v17) >> (uint(v18) % 32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v19+v14) < base.Ui32(int32(base.Ui32(v27)>>(uint(v18)%32))) {
		v31 = v15 + (v19+int32(3))&int32(2147483644)
	} else {
		v31 = v15
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v36 = int32(base.Ui32(v34) >> (uint(int32(2)) % 32))
	v40 = (v36 - int32(1)) & int32(-4)
	v41 = F_palloc(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return int32(0)
	} else {
		v46 = v36 - int32(4)
		if v40 <= v46 {
		} else {
			v50 = v40 - v36 + int32(4)
			if v50 == int32(0) {
			} else {
				base.MemoryFill(m, v41+v46, int32(0), v50)
			}
		}
		v57 = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(v57) % 32)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
		v64 = int32(base.Ui32(v60)>>(uint(v57)%32)) - int32(8)
		if v64 != 0 {
			base.MemoryCopy(m, v41+int32(4), v33+int32(8), v64)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
		v74 = F_gbt_var_key_copy(m, v9+int32(8))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v41)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v74
			}
		}
	}
}
func F_gbt_bit_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13881(m, l0, int32(_a_F_gbt_bit_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_bit_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13882(m, l0, l1, l2, int32(2847))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_bool_penalty(m *base.Module, l0 int32) int32 {
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
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v15 = float64(0.49000000953674316)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	v19 = base.F64_mul(base.F64_convert_i32_u(v16), v15)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if base.Ui32(v16) < base.Ui32(v22) {
		v24 = base.F64_sub(v15, v19)
	} else {
		v24 = float64(0)
	}
	v25 = float64(0.49000000953674316)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.Ui32(v26) < base.Ui32(v32) {
		v34 = base.F64_sub(v25, base.F64_mul(base.F64_convert_i32_u(v26), v25))
	} else {
		v34 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v35 = base.F64_add(v24, v34)
	if base.F64_gt(v35, float64(0)) != 0 {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v35, base.F64_add(base.F64_add(v19, base.F64_mul(base.F64_convert_i32_u(v32), float64(-0.49000000953674316))), v35))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v50+int32(1))))
	} else {
	}
	return v12
}
func F_gbt_bool_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_bool_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_booleq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v4 == v5)
}
func F_gbt_boolge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(base.Ui32(v5) <= base.Ui32(v4))
}
func F_gbt_boollt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v5))
}
func F_gbt_bpchar_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bpchar_sortsupport_0)
	return v4
}
func F_gbt_bpcharcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2407), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_bpchareq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2393), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_bytea_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_bytea_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_bytea_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13881(m, l0, int32(_a_F_gbt_bytea_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_byteagt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2844), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_cash_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_cash_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_cashle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 <= v5)
}
func F_gbt_date_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_date_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_date_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_date_sortsupport_0)
	return v4
}
func F_gbt_datege(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2416), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_datelt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2413), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_enum_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13888(m, l0, int32(_a_F_gbt_enum_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_enum_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_enum_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_enum_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13883(m, l0, int32(_a_F_gbt_enum_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_enum_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = F_CallerFInfoFunctionCall2(m, int32(3795), v5, int32(0), v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_gbt_float4_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_float4_sortsupport_0)
	return v4
}
func F_gbt_float4eq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F32_eq(v4, v5)
}
func F_gbt_float4lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F32_lt(v4, v5)
}
func F_gbt_float8_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_float8_picksplit_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_float8eq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_eq(v4, v5)
}
func F_gbt_float8ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_ge(v4, v5)
}
func F_gbt_inet_union(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13884(m, l0, int32(_a_F_gbt_inet_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int2_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.F64_abs(base.F64_sub(base.F64_convert_i32_s(v4), base.F64_convert_i32_s(v6)))
}
func F_gbt_int2_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v5 < v4) - base.B2i32(v4 < v5)
}
func F_gbt_int2le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v4 <= v5)
}
func F_gbt_int4_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13888(m, l0, int32(_a_F_gbt_int4_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_int4_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13883(m, l0, int32(_a_F_gbt_int4_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_intv_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13883(m, l0, int32(_a_F_gbt_intv_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_intv_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2524), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macad8_union(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13892(m, l0, int32(_a_F_gbt_macad8_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_macad8eq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8eq_0), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13891(m, l0, l1, l2, int32(8), int32(_a_F_gbt_macad8key_cmp_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macadeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2260), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macadlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2261), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_num_distance(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
		v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, l1, v12, l4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return float64(0)
		} else {
			if v14 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
				v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) float64)(m, l1, v18, l4)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return float64(0)
				} else {
					v32 = v20
					m.G0 = v9 + int32(16)
					return v32
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
				v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, l1, v22, l4)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return float64(0)
				} else {
					if v24 == int32(0) {
						v32 = float64(0)
						m.G0 = v9 + int32(16)
						return v32
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
						v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) float64)(m, l1, v28, l4)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return float64(0)
						} else {
							v32 = v30
							m.G0 = v9 + int32(16)
							return v32
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return float64(0)
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41
			F_errmsg_internal(m, int32(_a_F_gbt_num_distance_0), v9)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(_a_F_gbt_num_distance_1), int32(326), int32(_a_F_gbt_num_distance_2))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_gbt_numeric_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(1328), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_numeric_ge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(1279), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_oid_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_oid_sortsupport_0)
	return v4
}
func F_gbt_oidge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v5) <= base.Ui32(v4))
}
func F_gbt_oidlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v5))
}
func F_gbt_text_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13881(m, l0, int32(_a_F_gbt_text_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_textgt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2230), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13883(m, l0, int32(_a_F_gbt_time_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_timele(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2419), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_timetz_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
	if v7 != int32(1) {
		return v6
	} else {
		v12 = F_palloc(m, int32(16))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v18 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+8)))
				v24 = v20 + v21*int64(1000000)
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v24
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v12
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v30
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v33)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v32)
				return v18
			}
		}
	}
}
func F_gbt_timetz_consistent(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+30)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
	v22 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v16 + v22
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v20 + v21*int64(1000000)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+v37)+12)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = F_gbt_num_consistent(m, v10+v22, v10+int32(16), v10+int32(30), v39&v18, int32(_a_F_gbt_timetz_consistent_0), v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(32)
		return v44
	}
}
func F_gbt_ts_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13886(m, l0, int32(_a_F_gbt_ts_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_ts_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_ts_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_ts_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13883(m, l0, int32(_a_F_gbt_ts_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_tstz_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)))
	if v6 != int32(1) {
		return v5
	} else {
		v11 = F_palloc(m, int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			v18 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v11))) = v16
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v11
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v25
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)))
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v27)
				return v18
			}
		}
	}
}
func F_gbt_uuid_consistent(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13890(m, l0, int32(_a_F_gbt_uuid_consistent_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_uuidle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	return base.B2i32(v161 <= int32(0))
}
