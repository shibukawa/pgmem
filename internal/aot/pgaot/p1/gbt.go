package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_penalty(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14267(m, l0, int32(_a_F_gbt_bit_penalty_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_bit_union(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14271(m, l0, int32(_a_F_gbt_bit_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_bitcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(3062), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v9)
	}
}
func F_gbt_bitge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2855), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_bitlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2858), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_bpchargt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2609), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_bpcharle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2608), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_bytea_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bytea_sortsupport_0)
	return int64(0)
}
func F_gbt_byteaeq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(3056), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_bytealt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(3057), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_cash_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_cash_sortsupport_0)
	return int64(0)
}
func F_gbt_date_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_date_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
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
			v15 = v12
			v16 = v13
			v21 = F_CallerFInfoFunctionCall2(m, int32(4010), l2, int32(0), base.I64_extend_i32_u(v15), base.I64_extend_i32_u(v16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.I32_wrap_i64(v21)
				return v28
			}
		}
	} else {
		v15 = v7
		v16 = v9
		v21 = F_CallerFInfoFunctionCall2(m, int32(4010), l2, int32(0), base.I64_extend_i32_u(v15), base.I64_extend_i32_u(v16))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v28 = base.I32_wrap_i64(v21)
			return v28
		}
	}
}
func F_gbt_enumlt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(4006), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_float4_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_float4_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_float8_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_float8_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_float8le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
		v22 = base.F64_ge(v6, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_gbt_inet_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
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
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+30)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v2)
	v22 = F_convert_network_to_scalar(m, v11, int32(869), v9+int32(7))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v22
		v27 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v27)
		v29 = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v15 + v29
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v15
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v40)+12)))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = F_gbt_num_consistent(m, v9+int32(20), v9+v29, v9+int32(30), v42&v27, int32(_a_F_gbt_inet_consistent_0), v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			m.G0 = v9 + int32(32)
			return base.I64_extend_i32_u(v47)
		}
	}
}
func F_gbt_int2_distance(m *base.Module, l0 int32) int64 {
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
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v11)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v23)+12)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_gbt_num_distance(m, v8+int32(4), v8+int32(14), v25&int32(1), int32(_a_F_gbt_int2_distance_0), v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.I64_reinterpret_f64(v30)
	}
}
func F_gbt_int2_penalty(m *base.Module, l0 int32) int64 {
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
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11))))
	v13 = base.F64_convert_i32_s(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	v17 = base.F64_convert_i32_s(v16)
	if base.F64_lt(v17, v13) != 0 {
		v21 = base.F64_sub(v13, v17)
	} else {
		v21 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+2)))
	v23 = base.F64_convert_i32_s(v22)
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
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
func F_gbt_int2_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int2_sortsupport_0)
	return int64(0)
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
func F_gbt_int4_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_int4_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_int4_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_int4_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
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
func F_gbt_int8_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(_a_F_gbt_int8_sortsupport_0)
	return int64(0)
}
func F_gbt_intv_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13 + v8
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v9+int32(8), v11, v23&int32(1), int32(_a_F_gbt_intv_distance_0), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return base.I64_reinterpret_f64(v28)
	}
}
func F_gbt_intv_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_intv_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_intv_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_intv_union_0), int32(32))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_intveq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2655), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad8_consistent(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14280(m, l0, int32(_a_F_gbt_macad8_consistent_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_macad_consistent(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14280(m, l0, int32(_a_F_gbt_macad_consistent_0), int32(6))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_macaddr_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v6 = int64(4294967295)
	v10 = F_DirectFunctionCall2Coll(m, int32(2486), int32(0), l0&v6, l1&v6)
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
	}
}
func F_gbt_numeric_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v168 float32
	_ = v168
	var v171 float32
	_ = v171
	var v174 float32
	_ = v174
	var v182 float32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 float32
	_ = v197
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v21 = v15 + int32(16)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = int32(4)
	v26 = v23 + v25
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
		v42 = v26 + (v30+int32(3))&int32(2147483644)
	} else {
		v42 = v26
	}
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v42
	v44 = F_gbt_var_key_copy(m, v21)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		return int64(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = base.I64_extend_i32_u(v44)
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_gbt_var_bin_union(m, v15+int32(24), v19, v52, int32(_a_F_gbt_numeric_penalty_0), v54)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int64(0)
		} else {
			v58 = v15 + int32(8)
			v60 = int32(4)
			v61 = v23 + v60
			*(*int32)(unsafe.Add(mBase, uint32(v58))) = v61
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
			v64 = int32(2)
			v65 = int32(base.Ui32(v63) >> (uint(v64) % 32))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			if base.Ui32(v65+v60) < base.Ui32(int32(base.Ui32(v73)>>(uint(v64)%32))) {
				v77 = v61 + (v65+int32(3))&int32(2147483644)
			} else {
				v77 = v61
			}
			*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v77
			v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)))
			v80 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)))
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
			v83 = int32(4)
			v84 = v81 + v83
			*(*int32)(unsafe.Add(mBase, uint32(v58))) = v84
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
			v87 = int32(2)
			v88 = int32(base.Ui32(v86) >> (uint(v87) % 32))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
			if base.Ui32(v88+v83) < base.Ui32(int32(base.Ui32(v96)>>(uint(v87)%32))) {
				v100 = v84 + (v88+int32(3))&int32(2147483644)
			} else {
				v100 = v84
			}
			*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v100
			v104 = base.I32_wrap_i64(v17)
			v105 = int32(19)
			v106 = int32(0)
			v109 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)))
			v110 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)))
			v111 = F_DirectFunctionCall2Coll(m, v105, v106, v109, v110)
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int64(0)
			} else {
				v114 = F_pg_detoast_datum(m, base.I32_wrap_i64(v111))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int64(0)
				} else {
					v116 = base.I64_extend_i32_u(v114)
					v119 = F_DirectFunctionCall2Coll(m, int32(19), int32(0), v79, v80)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int64(0)
					} else {
						v122 = F_pg_detoast_datum(m, base.I32_wrap_i64(v119))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int64(0)
						} else {
							v125 = F_DirectFunctionCall2Coll(m, v105, v106, v116, base.I64_extend_i32_u(v122))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								v128 = F_pg_detoast_datum(m, base.I32_wrap_i64(v125))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int64(0)
								} else {
									v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
									if v130 == int32(_a_F_gbt_numeric_penalty_1) {
										v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
										if base.B2i32(v134 == int32(_a_F_gbt_numeric_penalty_1)) == int32(0) {
											v182 = float32(1)
											v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
											v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
											v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
											v197 = base.F32_mul(v182, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v187+int32(1))))
										} else {
											v197 = float32(0)
										}
										*(*float32)(unsafe.Add(mBase, uint32(v104))) = v197
										m.G0 = v15 + int32(32)
										return v17 & int64(4294967295)
									} else {
										v140 = F_int64_to_numeric(m, int64(0))
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int64(0)
										} else {
											v142 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v104))) = v142
											v146 = base.I64_extend_i32_u(v128)
											v148 = F_DirectFunctionCall2Coll(m, int32(2934), v142, v146, base.I64_extend_i32_u(v140))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int64(0)
											} else {
												v150 = *(*float32)(unsafe.Add(mBase, uint32(v104)))
												if v148 != int64(0) {
													*(*float32)(unsafe.Add(mBase, uint32(v104))) = base.F32_add(v150, float32(1.1754944e-38))
													v157 = int32(0)
													v160 = F_DirectFunctionCall2Coll(m, int32(1391), v157, v146, v116)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return int64(0)
													} else {
														v163 = F_pg_detoast_datum(m, base.I32_wrap_i64(v160))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return int64(0)
														} else {
															v166 = F_DirectFunctionCall1Coll(m, int32(1706), v157, base.I64_extend_i32_u(v163))
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int64(0)
															} else {
																v168 = *(*float32)(unsafe.Add(mBase, uint32(v104)))
																v171 = base.F32_add(v168, base.F32_demote_f64(base.F64_reinterpret_i64(v166)))
																*(*float32)(unsafe.Add(mBase, uint32(v104))) = v171
																v174 = v171
																if base.F32_gt(v174, float32(0)) == int32(0) {
																} else {
																	v182 = v174
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
																	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
																	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
																	v197 = base.F32_mul(v182, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v187+int32(1))))
																	*(*float32)(unsafe.Add(mBase, uint32(v104))) = v197
																}
																m.G0 = v15 + int32(32)
																return v17 & int64(4294967295)
															}
														}
													}
												} else {
													v174 = v150
													if base.F32_gt(v174, float32(0)) == int32(0) {
													} else {
														v182 = v174
														v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
														v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
														v197 = base.F32_mul(v182, base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v187+int32(1))))
														*(*float32)(unsafe.Add(mBase, uint32(v104))) = v197
													}
													m.G0 = v15 + int32(32)
													return v17 & int64(4294967295)
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
func F_gbt_numeric_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_numeric_sortsupport_0)
	return int64(0)
}
func F_gbt_oid_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_oid_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_oid_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14278(m, l0, int32(_a_F_gbt_oid_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
func F_gbt_oid_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_oid_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_textlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2448), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_time_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
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
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 float32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v21 = F_DirectFunctionCall2Coll(m, int32(2919), int32(0), v17, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v25 = base.I32_wrap_i64(v21)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
		v32 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
		v33 = F_DirectFunctionCall2Coll(m, int32(2919), int32(0), v31, v32)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			v35 = base.I32_wrap_i64(v33)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
			v39 = base.I32_wrap_i64(v12)
			*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(0)
			v52 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v26), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v27), float64(86400)), base.F64_div(base.F64_convert_i64_s(v28), float64(1e+06))))
			v53 = float64(0)
			if base.F64_gt(v52, v53) != 0 {
				v56 = v52
			} else {
				v56 = v53
			}
			v67 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v36), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v37), float64(86400)), base.F64_div(base.F64_convert_i64_s(v38), float64(1e+06))))
			v68 = float64(0)
			if base.F64_gt(v67, v68) != 0 {
				v71 = v67
			} else {
				v71 = v68
			}
			v72 = base.F64_add(v56, v71)
			if base.F64_gt(v72, float64(0)) != 0 {
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
				v79 = F_DirectFunctionCall2Coll(m, int32(2919), int32(0), v77, v78)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					v81 = *(*float32)(unsafe.Add(mBase, uint32(v39)))
					v84 = base.I32_wrap_i64(v79)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
					v93 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+52))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
					*(*float32)(unsafe.Add(mBase, uint32(v39))) = base.F32_mul(base.F32_add(base.F32_add(v81, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(v72, base.F64_add(v72, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v85), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v89), float64(86400)), base.F64_div(base.F64_convert_i64_s(v93), float64(1e+06)))))))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v107+int32(1))))
					return v12 & int64(4294967295)
				}
			} else {
				return v12 & int64(4294967295)
			}
		}
	}
}
func F_gbt_timeeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2644), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_timege(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2634), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_timekey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14285(m, l0, l1, l2, int32(1577))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_timekey_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
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
	v7 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	v10 = F_DirectFunctionCall2Coll(m, int32(1577), int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
	}
}
func F_gbt_ts_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_ts_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_ts_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_ts_sortsupport_0)
	return int64(0)
}
func F_gbt_tstz_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v11
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v9+v14, v9, v24&int32(1), int32(_a_F_gbt_tstz_distance_0), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return base.I64_reinterpret_f64(v29)
	}
}
func F_gbt_uuid_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v51 float64
	_ = v51
	var v53 int64
	_ = v53
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v132 int64
	_ = v132
	var v169 float64
	_ = v169
	var v171 int64
	_ = v171
	var v210 int64
	_ = v210
	var v247 float64
	_ = v247
	var v248 int64
	_ = v248
	var v287 int64
	_ = v287
	var v324 float64
	_ = v324
	var v328 float64
	_ = v328
	var v331 float64
	_ = v331
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v350 float32
	_ = v350
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = int64(56)
	v17 = int64(65280)
	v19 = int64(40)
	v22 = int64(16711680)
	v24 = int64(24)
	v26 = int64(4278190080)
	v28 = int64(8)
	v51 = float64(5.421010862427522e-20)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v90 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v14<<(uint(v15)%64)|v14&v17<<(uint(v19)%64)|(v14&v22<<(uint(v24)%64)|v14&v26<<(uint(v28)%64))|(int64(base.Ui64(v14)>>(uint(v28)%64))&v26|int64(base.Ui64(v14)>>(uint(v24)%64))&v22|(int64(base.Ui64(v14)>>(uint(v19)%64))&v17|int64(base.Ui64(v14)>>(uint(v15)%64))))), v51), base.F64_convert_i64_u(v53<<(uint(v15)%64)|v53&v17<<(uint(v19)%64)|(v53&v22<<(uint(v24)%64)|v53&v26<<(uint(v28)%64))|(int64(base.Ui64(v53)>>(uint(v28)%64))&v26|int64(base.Ui64(v53)>>(uint(v24)%64))&v22|(int64(base.Ui64(v53)>>(uint(v19)%64))&v17|int64(base.Ui64(v53)>>(uint(v15)%64))))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	v169 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v93<<(uint(v15)%64)|v93&v17<<(uint(v19)%64)|(v93&v22<<(uint(v24)%64)|v93&v26<<(uint(v28)%64))|(int64(base.Ui64(v93)>>(uint(v28)%64))&v26|int64(base.Ui64(v93)>>(uint(v24)%64))&v22|(int64(base.Ui64(v93)>>(uint(v19)%64))&v17|int64(base.Ui64(v93)>>(uint(v15)%64))))), v51), base.F64_convert_i64_u(v132<<(uint(v15)%64)|v132&v17<<(uint(v19)%64)|(v132&v22<<(uint(v24)%64)|v132&v26<<(uint(v28)%64))|(int64(base.Ui64(v132)>>(uint(v28)%64))&v26|int64(base.Ui64(v132)>>(uint(v24)%64))&v22|(int64(base.Ui64(v132)>>(uint(v19)%64))&v17|int64(base.Ui64(v132)>>(uint(v15)%64))))))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v92)+24))
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v92)+16))
	v247 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v171<<(uint(v15)%64)|v171&v17<<(uint(v19)%64)|(v171&v22<<(uint(v24)%64)|v171&v26<<(uint(v28)%64))|(int64(base.Ui64(v171)>>(uint(v28)%64))&v26|int64(base.Ui64(v171)>>(uint(v24)%64))&v22|(int64(base.Ui64(v171)>>(uint(v19)%64))&v17|int64(base.Ui64(v171)>>(uint(v15)%64))))), v51), base.F64_convert_i64_u(v210<<(uint(v15)%64)|v210&v17<<(uint(v19)%64)|(v210&v22<<(uint(v24)%64)|v210&v26<<(uint(v28)%64))|(int64(base.Ui64(v210)>>(uint(v28)%64))&v26|int64(base.Ui64(v210)>>(uint(v24)%64))&v22|(int64(base.Ui64(v210)>>(uint(v19)%64))&v17|int64(base.Ui64(v210)>>(uint(v15)%64))))))
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v324 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v248<<(uint(v15)%64)|v248&v17<<(uint(v19)%64)|(v248&v22<<(uint(v24)%64)|v248&v26<<(uint(v28)%64))|(int64(base.Ui64(v248)>>(uint(v28)%64))&v26|int64(base.Ui64(v248)>>(uint(v24)%64))&v22|(int64(base.Ui64(v248)>>(uint(v19)%64))&v17|int64(base.Ui64(v248)>>(uint(v15)%64))))), v51), base.F64_convert_i64_u(v287<<(uint(v15)%64)|v287&v17<<(uint(v19)%64)|(v287&v22<<(uint(v24)%64)|v287&v26<<(uint(v28)%64))|(int64(base.Ui64(v287)>>(uint(v28)%64))&v26|int64(base.Ui64(v287)>>(uint(v24)%64))&v22|(int64(base.Ui64(v287)>>(uint(v19)%64))&v17|int64(base.Ui64(v287)>>(uint(v15)%64))))))
	if base.F64_gt(v247, v324) != 0 {
		v328 = base.F64_sub(v247, v324)
	} else {
		v328 = float64(0)
	}
	if base.F64_lt(v169, v90) != 0 {
		v331 = base.F64_add(base.F64_sub(v90, v169), v328)
	} else {
		v331 = v328
	}
	if base.F64_gt(v331, float64(0)) != 0 {
		v341 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+52))
		v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
		v350 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v331, base.F64_add(base.F64_sub(v324, v90), v331))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v343+int32(1))))
	} else {
		v350 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v10)))) = v350
	return v10
}
func F_gbt_uuid_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_uuid_sortsupport_0)
	return int64(0)
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
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v107)
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v217 int32
	_ = v217
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
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
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(1)
	v26 = (v22 - v23) & int32(_a_F_gbt_var_picksplit_0)
	v28 = v26 + v23
	v29 = F_palloc_mul(m, int32(8), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v36 = v26<<(uint(int32(1))%32) + int32(4)
	v37 = F_palloc(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
	v40 = F_palloc(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47
	v52 = l1 + int32(32)
	v54 = l1 + int32(8)
	v56 = F_palloc_mul(m, int32(4), v28)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v22&int32(_a_F_gbt_var_picksplit_0) != int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v217 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L7:
	;
	v65 = int32(1)
	v76 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v193 = int32(8)
	F_qsort_arg(m, v29+v193, v26, v193, int32(_a_F_gbt_var_picksplit_1), v19+int32(4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L33
	}
L10:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v65*int32(24))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v87 = int32(base.Ui32(v85) >> (uint(int32(2)) % 32))
	if v87 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l3
	v128 = int32(8)
	F_qsort_arg(m, v29+v128, v26, v128, int32(_a_F_gbt_var_picksplit_1), v19+int32(4))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v65<<(uint(int32(3))%32)))) = v65
	v123 = (v65 + int32(1)) & int32(_a_F_gbt_var_picksplit_0)
	if base.Ui32(v123) <= base.Ui32(v26) {
		v65 = v123
		v76 = v114
		goto L10
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v65<<(uint(int32(3))%32))+4)) = v84
	v114 = v76
	goto L12
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if base.Ui32(v87+int32(4)) < base.Ui32(int32(base.Ui32(v90)>>(uint(int32(2))%32))) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v94 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v95 = m.T0[v94].(func(*base.Module, int32, int32) int32)(m, v84, l4)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v97 = v84
	goto L20
L20:
	;
	v100 = v56 + v76<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v29+v65<<(uint(int32(3))%32))+4)) = v97
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v114 = v76 + base.B2i32(v106 != v84)
	goto L12
L21:
	;
	v97 = v95
	goto L20
L22:
	;
	goto L11
L23:
	;
	v136 = int32(1)
	v139 = v136
	goto L24
L24:
	;
	v157 = v29 + v139<<(uint(int32(3))%32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if base.Ui32(v139) <= base.Ui32(int32(base.Ui32(v26)>>(uint(v136)%32))) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L6
L26:
	;
	v188 = (v139 + int32(1)) & int32(_a_F_gbt_var_picksplit_0)
	if base.Ui32(v188) <= base.Ui32(v26) {
		v139 = v188
		goto L24
	} else {
		goto L32
	}
L27:
	;
	F_gbt_var_bin_union(m, v54, v158, l2, l3, l4)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_gbt_var_bin_union(m, v52, v158, l2, l3, l4)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v164 = int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*uint16)(unsafe.Add(mBase, uint32(v162+v163<<(uint(v164)%32)))) = uint16(v167)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v163 + v164
	goto L26
L31:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v176 = int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*uint16)(unsafe.Add(mBase, uint32(v174+v175<<(uint(v176)%32)))) = uint16(v179)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v175 + v176
	goto L26
L32:
	;
	goto L25
L33:
	;
	goto L6
L34:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v221 = F_gbt_var_node_cp_len(m, v220, l3)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
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
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v224 = F_gbt_var_node_cp_len(m, v223, l3)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v227 = int32(4)
	v228 = v226 + v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v230 = int32(2)
	v231 = int32(base.Ui32(v229) >> (uint(v230) % 32))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if base.Ui32(v231+v227) < base.Ui32(int32(base.Ui32(v239)>>(uint(v230)%32))) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v243 = v228 + (v231+int32(3))&int32(2147483644)
	goto L41
L40:
	;
	v243 = v228
	goto L41
L41:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v248 = int32(base.Ui32(v244)>>(uint(int32(2))%32)) - int32(4)
	if v224 < v221 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v250 = v221
	goto L44
L43:
	;
	v250 = v224
	goto L44
L44:
	;
	v252 = v250 + int32(2)
	if v248 < v252 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v254 = v248
	goto L47
L46:
	;
	v254 = v252
	goto L47
L47:
	;
	v256 = v231 - int32(4)
	if v256 < v252 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v258 = v256
	goto L50
L49:
	;
	v258 = v252
	goto L50
L50:
	;
	v262 = (v258 + int32(7)) & int32(-4)
	v265 = v254 + v262 + int32(8)
	v266 = F_palloc0(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v265 << (uint(int32(2)) % 32)
	v271 = int32(4)
	v272 = v266 + v271
	v274 = v258 + v271
	if v274 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	base.MemoryCopy(m, v272, v228, v274)
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v274 << (uint(int32(2)) % 32)
	v279 = v272 + v262
	v281 = v254 + int32(4)
	if v281 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v279, v243, v281)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v283 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v281 << (uint(v283) % 32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v287 = int32(4)
	v288 = v286 + v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v291 = int32(base.Ui32(v289) >> (uint(v283) % 32))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	if base.Ui32(v291+v287) < base.Ui32(int32(base.Ui32(v299)>>(uint(v283)%32))) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v303 = v288 + (v291+int32(3))&int32(2147483644)
	goto L60
L59:
	;
	v303 = v288
	goto L60
L60:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v308 = int32(base.Ui32(v304)>>(uint(int32(2))%32)) - int32(4)
	if v308 < v252 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v310 = v308
	goto L63
L62:
	;
	v310 = v252
	goto L63
L63:
	;
	v312 = v291 - int32(4)
	if v312 < v252 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v314 = v312
	goto L66
L65:
	;
	v314 = v252
	goto L66
L66:
	;
	v318 = (v314 + int32(7)) & int32(-4)
	v321 = v310 + v318 + int32(8)
	v322 = F_palloc0(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v321 << (uint(int32(2)) % 32)
	v327 = int32(4)
	v328 = v322 + v327
	v330 = v314 + v327
	if v330 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v328, v288, v330)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v330 << (uint(int32(2)) % 32)
	v335 = v328 + v318
	v337 = v310 + int32(4)
	if v337 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v335, v303, v337)
	goto L73
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v337 << (uint(int32(2)) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = base.I64_extend_i32_u(v266)
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = base.I64_extend_i32_u(v322)
	goto L36
}
