package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bitle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2659), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_bool_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7452)
	return v3
}
func F_gbt_bool_union(m *base.Module, l0 int32) int32 {
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
	v6 = F_palloc(m, int32(2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4434328), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_bpcharlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2410), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_cash_consistent(m *base.Module, l0 int32) int32 {
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
	var v11 int64
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
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v11
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
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(4434408), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v37
	}
}
func F_gbt_cash_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4434408), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_cashgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v5 < v4)
}
func F_gbt_date_distance(m *base.Module, l0 int32) int32 {
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
	v29 = F_gbt_num_distance(m, v7+v13, v7+int32(12), v24&int32(1), int32(4434448), v28)
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
func F_gbt_date_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = int32(2443)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = F_DirectFunctionCall2Coll(m, v8, int32(0), v12, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v23 = F_DirectFunctionCall2Coll(m, v8, int32(0), v21, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v25
			if v25 < v23 {
				v30 = v23
			} else {
				v30 = v25
			}
			v31 = int32(0)
			if v31 < v16 {
				v34 = v16
			} else {
				v34 = v31
			}
			v35 = v30 + v34
			if v35 != 0 {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v40 = F_DirectFunctionCall2Coll(m, int32(2443), int32(0), v38, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*float32)(unsafe.Add(mBase, uint32(v7)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					*(*float32)(unsafe.Add(mBase, uint32(v7))) = base.F32_mul(base.F32_add(base.F32_add(v42, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(base.F64_convert_i32_u(v35), base.F64_convert_i32_s(v40+v35)))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v55+int32(1))))
					return v7
				}
			} else {
				return v7
			}
		}
	}
}
func F_gbt_date_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4434448), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_dateeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2428), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_datekey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = F_DirectFunctionCall2Coll(m, int32(1444), int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v20 = v10
			return v20
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v18 = F_DirectFunctionCall2Coll(m, int32(1444), int32(0), v16, v17)
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
func F_gbt_enumgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3808), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_float4_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*float32)(unsafe.Add(mBase, uint32(v7)+12)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v13 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12 + v13
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v7+v13, v7+int32(12), v24&int32(1), int32(4434528), v28)
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
func F_gbt_float4_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4434528), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_float4ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	return base.F32_ge(v4, v5)
}
func F_gbt_float8_distance(m *base.Module, l0 int32) int32 {
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
	var v11 float64
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
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v7, v7+v14, v23&int32(1), int32(4434568), v27)
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
func F_gbt_float8_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
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
	var v16 float64
	_ = v16
	var v19 float64
	_ = v19
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v2 = float64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	if base.F64_gt(v16, v19) != 0 {
		v27 = base.F64_add(base.F64_add(base.F64_mul(v16, float64(0.49000000953674316)), base.F64_mul(v19, float64(-0.49000000953674316))), v2)
	} else {
		v27 = v2
	}
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.F64_gt(v28, v31) != 0 {
		v37 = base.F64_add(v27, base.F64_add(base.F64_mul(v28, float64(0.49000000953674316)), base.F64_mul(v31, float64(-0.49000000953674316))))
	} else {
		v37 = v27
	}
	if base.F64_gt(v37, float64(0)) != 0 {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
		*(*float32)(unsafe.Add(mBase, uint32(v13))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v37, base.F64_add(base.F64_add(base.F64_mul(v19, float64(0.49000000953674316)), base.F64_mul(v28, float64(-0.49000000953674316))), v37))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v53+int32(1))))
	} else {
	}
	return v13
}
func F_gbt_float8_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4434568), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_float8lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_lt(v4, v5)
}
func F_gbt_int2_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4434712), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_int2gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return base.B2i32(v5 < v4)
}
func F_gbt_int4gt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v5 < v4)
}
func F_gbt_int8_consistent(m *base.Module, l0 int32) int32 {
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
	var v11 int64
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
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v11
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
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(4434792), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v37
	}
}
func F_gbt_int8_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4434792), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_intv_compress(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
	if v7 != int32(1) {
		return v6
	} else {
		v12 = F_palloc(m, int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = F_palloc(m, int32(16))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v21
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v23
				if v19 != 0 {
					v27 = int32(0)
				} else {
					v27 = int32(16)
				}
				v28 = v20 + v27
				v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v29
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v12
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v36
				v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
				v39 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v39)
				*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v38)
				return v17
			}
		}
	}
}
func F_gbt_intv_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14 + v7
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+v27)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v8+int32(4), v10, v8+int32(14), v29&int32(1), int32(4434832), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v34
	}
}
func F_gbt_intv_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 float64
	_ = v6
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v12 int64
	_ = v12
	var v14 float64
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = float64(2.592e+06)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = float64(86400)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v14 = float64(1e+06)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_abs(base.F64_sub(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v4), v6), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v8), v10), base.F64_div(base.F64_convert_i64_s(v12), v14))), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v18), v6), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v22), v10), base.F64_div(base.F64_convert_i64_s(v26), v14)))))
}
func F_gbt_intvgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2463), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_intvle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2461), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v118 int64
	_ = v118
	var v139 int64
	_ = v139
	var v160 int64
	_ = v160
	var v171 float64
	_ = v171
	var v181 float64
	_ = v181
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	v41 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	v42 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)))
	v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+7)))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+6)))
	v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+5)))
	v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+3)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+15)))
	v59 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+14)))
	v60 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+13)))
	v61 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
	v62 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+11)))
	v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
	v64 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+9)))
	v65 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)))
	v66 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+15)))
	v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+14)))
	v68 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+13)))
	v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+12)))
	v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+11)))
	v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+10)))
	v72 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+9)))
	v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v77 = int64(48)
	v79 = int64(56)
	v82 = int64(40)
	v85 = int64(32)
	v88 = int64(24)
	v91 = int64(16)
	v94 = int64(8)
	v97 = v50 + (v56<<(uint(v77)%64) | v57<<(uint(v79)%64) | v55<<(uint(v82)%64) | v54<<(uint(v85)%64) | v53<<(uint(v88)%64) | v52<<(uint(v91)%64) | v51<<(uint(v94)%64))
	v118 = v40 + (v46<<(uint(v77)%64) | v47<<(uint(v79)%64) | v45<<(uint(v82)%64) | v44<<(uint(v85)%64) | v43<<(uint(v88)%64) | v42<<(uint(v91)%64) | v41<<(uint(v94)%64))
	v139 = v58 + (v64<<(uint(v77)%64) | v65<<(uint(v79)%64) | v63<<(uint(v82)%64) | v62<<(uint(v85)%64) | v61<<(uint(v88)%64) | v60<<(uint(v91)%64) | v59<<(uint(v94)%64))
	v160 = v66 + (v72<<(uint(v77)%64) | v73<<(uint(v79)%64) | v71<<(uint(v82)%64) | v70<<(uint(v85)%64) | v69<<(uint(v88)%64) | v68<<(uint(v91)%64) | v67<<(uint(v94)%64))
	if base.Ui64(v160) < base.Ui64(v139) {
		v171 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v139), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v160), float64(-0.49000000953674316))), float64(0))
	} else {
		v171 = float64(0)
	}
	if base.Ui64(v97) < base.Ui64(v118) {
		v181 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v118), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v97), float64(-0.49000000953674316))), v171)
	} else {
		v181 = v171
	}
	if base.F64_gt(v181, float64(0)) != 0 {
		v197 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
		v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+52))
		v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
		*(*float32)(unsafe.Add(mBase, uint32(v74))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v181, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v160), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v118), float64(-0.49000000953674316))), v181))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v199+int32(1))))
	} else {
	}
	return v74
}
func F_gbt_macad8_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4434912), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macad8lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(4189), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4434872), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macad_union(m *base.Module, l0 int32) int32 {
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
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v6, v4, int32(4434872), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_macadge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2280), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macadkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2282), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v20 = v8
			return v20
		} else {
			v14 = int32(6)
			v18 = F_DirectFunctionCall2Coll(m, int32(2282), int32(0), v6+v14, v7+v14)
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
func F_gbt_numeric_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(1343), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gbt_numeric_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_gbt_var_same(m, v4, v5, v6, int32(4434952), v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v9)
		return v3
	}
}
func F_gbt_oid_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4434992), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_oidkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
			if base.Ui32(v12) < base.Ui32(v11) {
				v17 = int32(1)
			} else {
				v17 = int32(-1)
			}
			return v17
		}
	} else {
		if base.Ui32(v8) < base.Ui32(v6) {
			v22 = int32(1)
		} else {
			v22 = int32(-1)
		}
		v25 = v22
		return v25
	}
}
func F_gbt_text_consistent(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v20 = v9 + int32(8)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v23 = int32(4)
		v24 = v21 + v23
		*(*int32)(unsafe.Add(mBase, uint32(v20))) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		v27 = int32(2)
		v28 = int32(base.Ui32(v26) >> (uint(v27) % 32))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if base.Ui32(v28+v23) < base.Ui32(int32(base.Ui32(v36)>>(uint(v27)%32))) {
			v40 = v24 + (v28+int32(3))&int32(2147483644)
		} else {
			v40 = v24
		}
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v40
		v42 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v42)
		v45 = *(*int32)(unsafe.Add(mBase, _consts[1102]))
		if v45 == v42 {
			v50 = *(*int32)(unsafe.Add(mBase, _consts[106]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_consts[862])))
			*(*int32)(unsafe.Add(mBase, _consts[1102])) = v56
		} else {
		}
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+v64)+12)))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v71 = F_gbt_var_consistent(m, v9+int32(8), v13, v17&int32(65535), v62, v66&int32(1), int32(4435032), v70)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			m.G0 = v9 + int32(16)
			return v71
		}
	}
}
func F_gbt_textle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2245), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_consistent(m *base.Module, l0 int32) int32 {
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
	var v11 int64
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
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v11
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
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(4435112), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v37
	}
}
func F_gbt_time_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	v6 = F_DirectFunctionCall2Coll(m, int32(2721), int32(0), l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return float64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		return base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v10), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v14), float64(86400)), base.F64_div(base.F64_convert_i64_s(v18), float64(1e+06)))))
	}
}
func F_gbt_timegt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func F_gbt_ts_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4435152))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_ts_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int64
	_ = v6
	var v11 int64
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v38 float64
	_ = v38
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v6-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v38 = v5
		return v38
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		if base.Ui64(v11-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			v38 = v5
			return v38
		} else {
			v18 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), l0, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return float64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				v38 = base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v22), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v26), float64(86400)), base.F64_div(base.F64_convert_i64_s(v30), float64(1e+06)))))
				return v38
			}
		}
	}
}
func F_gbt_tstz_consistent(m *base.Module, l0 int32) int32 {
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
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+30)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v2)
	v19 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v15 + v19
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v31)+12)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_gbt_num_consistent(m, v8+int32(20), v8+v19, v8+int32(30), v33&int32(1), int32(4435152), v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(32)
		return v38
	}
}
func F_gbt_uuid_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4435192))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_var_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	v8 = int32(0)
	switch l2 - int32(1) {
	case 0:
		goto L6
	case 1:
		goto L7
	case 2:
		goto L5
	case 3:
		goto L3
	case 4:
		goto L4
	case 5:
		goto L2
	default:
		v893 = v8
		goto L1
	}
L1:
	;
	return v893
L2:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v877 = m.T0[v876].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v875, l3, l6)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L11
	} else {
		goto L262
	}
L3:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l4 != 0 {
		goto L213
	} else {
		goto L214
	}
L4:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l4 != 0 {
		goto L164
	} else {
		goto L165
	}
L5:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L111
	} else {
		goto L112
	}
L6:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L60
	} else {
		goto L61
	}
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v13 = m.T0[v12].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v11, l3, l6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v11, l3, l6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L11
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	return v13
L13:
	;
	if int32(0) <= v19 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v25 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = int32(2)
	v33 = int32(base.Ui32(v31) >> (uint(v32) % 32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = int32(base.Ui32(v34) >> (uint(v32) % 32))
	if base.Ui32(v36) < base.Ui32(v33) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v111 = int32(base.Ui32(v109) >> (uint(int32(2)) % 32))
	if base.Ui32(v36) < base.Ui32(v111) {
		v893 = v8
		goto L1
	} else {
		goto L41
	}
L21:
	;
	v38 = int32(4)
	v39 = l1 + v38
	v41 = v30 + v38
	v43 = v33 - v38
	if base.Ui32(v38) <= base.Ui32(v43) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if v105 != 0 {
		goto L20
	} else {
		goto L40
	}
L23:
	;
	v105 = int32(0)
	goto L22
L24:
	;
	v79 = v74
	v80 = v75
	v81 = v76
	goto L34
L25:
	;
	if (v39|v41)&int32(3) != 0 {
		v74 = v39
		v75 = v41
		v76 = v43
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v67 = v39
	v68 = v41
	v69 = v43
	goto L27
L27:
	;
	if v69 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v51 = v39
	v52 = v41
	v53 = v43
	goto L29
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v56 != v57 {
		v74 = v51
		v75 = v52
		v76 = v53
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v67 = v62
	v68 = v60
	v69 = v64
	goto L27
L31:
	;
	v59 = int32(4)
	v60 = v52 + v59
	v62 = v51 + v59
	v64 = v53 - v59
	if base.Ui32(int32(3)) < base.Ui32(v64) {
		v51 = v62
		v52 = v60
		v53 = v64
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v74 = v67
	v75 = v68
	v76 = v69
	goto L24
L34:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v84 == v85 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v105 = v84 - v85
	goto L22
L36:
	;
	v87 = int32(1)
	v92 = v81 - v87
	if v92 != 0 {
		v79 = v79 + v87
		v80 = v80 + v87
		v81 = v92
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
L40:
	;
	return int32(1)
L41:
	;
	v113 = int32(4)
	v114 = l1 + v113
	v116 = v108 + v113
	v118 = v111 - v113
	if base.Ui32(v113) <= base.Ui32(v118) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	return base.B2i32(v180 == int32(0))
L43:
	;
	v180 = int32(0)
	goto L42
L44:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L54
L45:
	;
	if (v114|v116)&int32(3) != 0 {
		v149 = v114
		v150 = v116
		v151 = v118
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v142 = v114
	v143 = v116
	v144 = v118
	goto L47
L47:
	;
	if v144 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v126 = v114
	v127 = v116
	v128 = v118
	goto L49
L49:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v131 != v132 {
		v149 = v126
		v150 = v127
		v151 = v128
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v142 = v137
	v143 = v135
	v144 = v139
	goto L47
L51:
	;
	v134 = int32(4)
	v135 = v127 + v134
	v137 = v126 + v134
	v139 = v128 - v134
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		v126 = v137
		v127 = v135
		v128 = v139
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v149 = v142
	v150 = v143
	v151 = v144
	goto L44
L54:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 == v160 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v180 = v159 - v160
	goto L42
L56:
	;
	v162 = int32(1)
	v167 = v156 - v162
	if v167 != 0 {
		v154 = v154 + v162
		v155 = v155 + v162
		v156 = v167
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L43
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v186 = m.T0[v185].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v184, l3, l6)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L11
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v190 = m.T0[v189].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v184, l3, l6)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L11
	} else {
		goto L64
	}
L63:
	;
	return v186
L64:
	;
	if int32(0) <= v190 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	return int32(1)
L66:
	;
	goto L67
L67:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v196 != int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	return int32(0)
L69:
	;
	goto L70
L70:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v203 = int32(2)
	v204 = int32(base.Ui32(v202) >> (uint(v203) % 32))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v207 = int32(base.Ui32(v205) >> (uint(v203) % 32))
	if base.Ui32(v207) < base.Ui32(v204) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v282 = int32(base.Ui32(v280) >> (uint(int32(2)) % 32))
	if base.Ui32(v207) < base.Ui32(v282) {
		v893 = v8
		goto L1
	} else {
		goto L92
	}
L72:
	;
	v209 = int32(4)
	v210 = l1 + v209
	v212 = v201 + v209
	v214 = v204 - v209
	if base.Ui32(v209) <= base.Ui32(v214) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v276 != 0 {
		goto L71
	} else {
		goto L91
	}
L74:
	;
	v276 = int32(0)
	goto L73
L75:
	;
	v250 = v245
	v251 = v246
	v252 = v247
	goto L85
L76:
	;
	if (v210|v212)&int32(3) != 0 {
		v245 = v210
		v246 = v212
		v247 = v214
		goto L75
	} else {
		goto L79
	}
L77:
	;
	v238 = v210
	v239 = v212
	v240 = v214
	goto L78
L78:
	;
	if v240 == int32(0) {
		goto L74
	} else {
		goto L84
	}
L79:
	;
	v222 = v210
	v223 = v212
	v224 = v214
	goto L80
L80:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v227 != v228 {
		v245 = v222
		v246 = v223
		v247 = v224
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v238 = v233
	v239 = v231
	v240 = v235
	goto L78
L82:
	;
	v230 = int32(4)
	v231 = v223 + v230
	v233 = v222 + v230
	v235 = v224 - v230
	if base.Ui32(int32(3)) < base.Ui32(v235) {
		v222 = v233
		v223 = v231
		v224 = v235
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v245 = v238
	v246 = v239
	v247 = v240
	goto L75
L85:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v255 == v256 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v276 = v255 - v256
	goto L73
L87:
	;
	v258 = int32(1)
	v263 = v252 - v258
	if v263 != 0 {
		v250 = v250 + v258
		v251 = v251 + v258
		v252 = v263
		goto L85
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	goto L74
L91:
	;
	return int32(1)
L92:
	;
	v284 = int32(4)
	v285 = l1 + v284
	v287 = v279 + v284
	v289 = v282 - v284
	if base.Ui32(v284) <= base.Ui32(v289) {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	return base.B2i32(v351 == int32(0))
L94:
	;
	v351 = int32(0)
	goto L93
L95:
	;
	v325 = v320
	v326 = v321
	v327 = v322
	goto L105
L96:
	;
	if (v285|v287)&int32(3) != 0 {
		v320 = v285
		v321 = v287
		v322 = v289
		goto L95
	} else {
		goto L99
	}
L97:
	;
	v313 = v285
	v314 = v287
	v315 = v289
	goto L98
L98:
	;
	if v315 == int32(0) {
		goto L94
	} else {
		goto L104
	}
L99:
	;
	v297 = v285
	v298 = v287
	v299 = v289
	goto L100
L100:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v302 != v303 {
		v320 = v297
		v321 = v298
		v322 = v299
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v313 = v308
	v314 = v306
	v315 = v310
	goto L98
L102:
	;
	v305 = int32(4)
	v306 = v298 + v305
	v308 = v297 + v305
	v310 = v299 - v305
	if base.Ui32(int32(3)) < base.Ui32(v310) {
		v297 = v308
		v298 = v306
		v299 = v310
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v320 = v313
	v321 = v314
	v322 = v315
	goto L95
L105:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v330 == v331 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v351 = v330 - v331
	goto L93
L107:
	;
	v333 = int32(1)
	v338 = v327 - v333
	if v338 != 0 {
		v325 = v325 + v333
		v326 = v326 + v333
		v327 = v338
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	goto L94
L111:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v357 = m.T0[v356].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v355, l3, l6)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v361 = m.T0[v360].(func(*base.Module, int32, int32, int32, int32) int32)(m, v355, l1, l3, l6)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L115
	}
L114:
	;
	return v357
L115:
	;
	if v361 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v368 = m.T0[v367].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v366, l3, l6)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v373 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	if v368 <= int32(0) {
		v893 = int32(1)
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	return int32(0)
L122:
	;
	goto L123
L123:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v380 = int32(2)
	v381 = int32(base.Ui32(v379) >> (uint(v380) % 32))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v384 = int32(base.Ui32(v382) >> (uint(v380) % 32))
	if base.Ui32(v384) < base.Ui32(v381) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v460 = int32(base.Ui32(v458) >> (uint(int32(2)) % 32))
	if base.Ui32(v384) < base.Ui32(v460) {
		v893 = int32(0)
		goto L1
	} else {
		goto L145
	}
L125:
	;
	v386 = int32(4)
	v387 = l1 + v386
	v389 = v378 + v386
	v391 = v381 - v386
	if base.Ui32(v386) <= base.Ui32(v391) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	if v453 != 0 {
		goto L124
	} else {
		goto L144
	}
L127:
	;
	v453 = int32(0)
	goto L126
L128:
	;
	v427 = v422
	v428 = v423
	v429 = v424
	goto L138
L129:
	;
	if (v387|v389)&int32(3) != 0 {
		v422 = v387
		v423 = v389
		v424 = v391
		goto L128
	} else {
		goto L132
	}
L130:
	;
	v415 = v387
	v416 = v389
	v417 = v391
	goto L131
L131:
	;
	if v417 == int32(0) {
		goto L127
	} else {
		goto L137
	}
L132:
	;
	v399 = v387
	v400 = v389
	v401 = v391
	goto L133
L133:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v404 != v405 {
		v422 = v399
		v423 = v400
		v424 = v401
		goto L128
	} else {
		goto L135
	}
L134:
	;
	v415 = v410
	v416 = v408
	v417 = v412
	goto L131
L135:
	;
	v407 = int32(4)
	v408 = v400 + v407
	v410 = v399 + v407
	v412 = v401 - v407
	if base.Ui32(int32(3)) < base.Ui32(v412) {
		v399 = v410
		v400 = v408
		v401 = v412
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v422 = v415
	v423 = v416
	v424 = v417
	goto L128
L138:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v432 == v433 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v453 = v432 - v433
	goto L126
L140:
	;
	v435 = int32(1)
	v440 = v429 - v435
	if v440 != 0 {
		v427 = v427 + v435
		v428 = v428 + v435
		v429 = v440
		goto L138
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	goto L127
L144:
	;
	return int32(1)
L145:
	;
	v462 = int32(4)
	v463 = l1 + v462
	v465 = v457 + v462
	v467 = v460 - v462
	if base.Ui32(v462) <= base.Ui32(v467) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	return base.B2i32(v529 == int32(0))
L147:
	;
	v529 = int32(0)
	goto L146
L148:
	;
	v503 = v498
	v504 = v499
	v505 = v500
	goto L158
L149:
	;
	if (v463|v465)&int32(3) != 0 {
		v498 = v463
		v499 = v465
		v500 = v467
		goto L148
	} else {
		goto L152
	}
L150:
	;
	v491 = v463
	v492 = v465
	v493 = v467
	goto L151
L151:
	;
	if v493 == int32(0) {
		goto L147
	} else {
		goto L157
	}
L152:
	;
	v475 = v463
	v476 = v465
	v477 = v467
	goto L153
L153:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	if v480 != v481 {
		v498 = v475
		v499 = v476
		v500 = v477
		goto L148
	} else {
		goto L155
	}
L154:
	;
	v491 = v486
	v492 = v484
	v493 = v488
	goto L151
L155:
	;
	v483 = int32(4)
	v484 = v476 + v483
	v486 = v475 + v483
	v488 = v477 - v483
	if base.Ui32(int32(3)) < base.Ui32(v488) {
		v475 = v486
		v476 = v484
		v477 = v488
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v498 = v491
	v499 = v492
	v500 = v493
	goto L148
L158:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	if v508 == v509 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v529 = v508 - v509
	goto L146
L160:
	;
	v511 = int32(1)
	v516 = v505 - v511
	if v516 != 0 {
		v503 = v503 + v511
		v504 = v504 + v511
		v505 = v516
		goto L158
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	goto L159
L163:
	;
	goto L147
L164:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v535 = m.T0[v534].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v533, l3, l6)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v540 = m.T0[v539].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v533, l3, l6)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L11
	} else {
		goto L168
	}
L167:
	;
	return v535
L168:
	;
	if v540 <= int32(0) {
		v893 = int32(1)
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v544 != int32(1) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	return int32(0)
L171:
	;
	goto L172
L172:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v551 = int32(2)
	v552 = int32(base.Ui32(v550) >> (uint(v551) % 32))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v555 = int32(base.Ui32(v553) >> (uint(v551) % 32))
	if base.Ui32(v555) < base.Ui32(v552) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	v631 = int32(base.Ui32(v629) >> (uint(int32(2)) % 32))
	if base.Ui32(v555) < base.Ui32(v631) {
		v893 = int32(0)
		goto L1
	} else {
		goto L194
	}
L174:
	;
	v557 = int32(4)
	v558 = l1 + v557
	v560 = v549 + v557
	v562 = v552 - v557
	if base.Ui32(v557) <= base.Ui32(v562) {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	if v624 != 0 {
		goto L173
	} else {
		goto L193
	}
L176:
	;
	v624 = int32(0)
	goto L175
L177:
	;
	v598 = v593
	v599 = v594
	v600 = v595
	goto L187
L178:
	;
	if (v558|v560)&int32(3) != 0 {
		v593 = v558
		v594 = v560
		v595 = v562
		goto L177
	} else {
		goto L181
	}
L179:
	;
	v586 = v558
	v587 = v560
	v588 = v562
	goto L180
L180:
	;
	if v588 == int32(0) {
		goto L176
	} else {
		goto L186
	}
L181:
	;
	v570 = v558
	v571 = v560
	v572 = v562
	goto L182
L182:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	if v575 != v576 {
		v593 = v570
		v594 = v571
		v595 = v572
		goto L177
	} else {
		goto L184
	}
L183:
	;
	v586 = v581
	v587 = v579
	v588 = v583
	goto L180
L184:
	;
	v578 = int32(4)
	v579 = v571 + v578
	v581 = v570 + v578
	v583 = v572 - v578
	if base.Ui32(int32(3)) < base.Ui32(v583) {
		v570 = v581
		v571 = v579
		v572 = v583
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v593 = v586
	v594 = v587
	v595 = v588
	goto L177
L187:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	if v603 == v604 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v624 = v603 - v604
	goto L175
L189:
	;
	v606 = int32(1)
	v611 = v600 - v606
	if v611 != 0 {
		v598 = v598 + v606
		v599 = v599 + v606
		v600 = v611
		goto L187
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L188
L192:
	;
	goto L176
L193:
	;
	return int32(1)
L194:
	;
	v633 = int32(4)
	v634 = l1 + v633
	v636 = v628 + v633
	v638 = v631 - v633
	if base.Ui32(v633) <= base.Ui32(v638) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	return base.B2i32(v700 == int32(0))
L196:
	;
	v700 = int32(0)
	goto L195
L197:
	;
	v674 = v669
	v675 = v670
	v676 = v671
	goto L207
L198:
	;
	if (v634|v636)&int32(3) != 0 {
		v669 = v634
		v670 = v636
		v671 = v638
		goto L197
	} else {
		goto L201
	}
L199:
	;
	v662 = v634
	v663 = v636
	v664 = v638
	goto L200
L200:
	;
	if v664 == int32(0) {
		goto L196
	} else {
		goto L206
	}
L201:
	;
	v646 = v634
	v647 = v636
	v648 = v638
	goto L202
L202:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	if v651 != v652 {
		v669 = v646
		v670 = v647
		v671 = v648
		goto L197
	} else {
		goto L204
	}
L203:
	;
	v662 = v657
	v663 = v655
	v664 = v659
	goto L200
L204:
	;
	v654 = int32(4)
	v655 = v647 + v654
	v657 = v646 + v654
	v659 = v648 - v654
	if base.Ui32(int32(3)) < base.Ui32(v659) {
		v646 = v657
		v647 = v655
		v648 = v659
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v669 = v662
	v670 = v663
	v671 = v664
	goto L197
L207:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v679 == v680 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v700 = v679 - v680
	goto L195
L209:
	;
	v682 = int32(1)
	v687 = v676 - v682
	if v687 != 0 {
		v674 = v674 + v682
		v675 = v675 + v682
		v676 = v687
		goto L207
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	goto L208
L212:
	;
	goto L196
L213:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v706 = m.T0[v705].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v704, l3, l6)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L11
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v711 = m.T0[v710].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v704, l3, l6)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L11
	} else {
		goto L217
	}
L216:
	;
	return v706
L217:
	;
	if v711 <= int32(0) {
		v893 = int32(1)
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v715 != int32(1) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	return int32(0)
L220:
	;
	goto L221
L221:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v722 = int32(2)
	v723 = int32(base.Ui32(v721) >> (uint(v722) % 32))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v726 = int32(base.Ui32(v724) >> (uint(v722) % 32))
	if base.Ui32(v726) < base.Ui32(v723) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v802 = int32(base.Ui32(v800) >> (uint(int32(2)) % 32))
	if base.Ui32(v726) < base.Ui32(v802) {
		v893 = int32(0)
		goto L1
	} else {
		goto L243
	}
L223:
	;
	v728 = int32(4)
	v729 = l1 + v728
	v731 = v720 + v728
	v733 = v723 - v728
	if base.Ui32(v728) <= base.Ui32(v733) {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	if v795 != 0 {
		goto L222
	} else {
		goto L242
	}
L225:
	;
	v795 = int32(0)
	goto L224
L226:
	;
	v769 = v764
	v770 = v765
	v771 = v766
	goto L236
L227:
	;
	if (v729|v731)&int32(3) != 0 {
		v764 = v729
		v765 = v731
		v766 = v733
		goto L226
	} else {
		goto L230
	}
L228:
	;
	v757 = v729
	v758 = v731
	v759 = v733
	goto L229
L229:
	;
	if v759 == int32(0) {
		goto L225
	} else {
		goto L235
	}
L230:
	;
	v741 = v729
	v742 = v731
	v743 = v733
	goto L231
L231:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	if v746 != v747 {
		v764 = v741
		v765 = v742
		v766 = v743
		goto L226
	} else {
		goto L233
	}
L232:
	;
	v757 = v752
	v758 = v750
	v759 = v754
	goto L229
L233:
	;
	v749 = int32(4)
	v750 = v742 + v749
	v752 = v741 + v749
	v754 = v743 - v749
	if base.Ui32(int32(3)) < base.Ui32(v754) {
		v741 = v752
		v742 = v750
		v743 = v754
		goto L231
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	v764 = v757
	v765 = v758
	v766 = v759
	goto L226
L236:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769))))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v774 == v775 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v795 = v774 - v775
	goto L224
L238:
	;
	v777 = int32(1)
	v782 = v771 - v777
	if v782 != 0 {
		v769 = v769 + v777
		v770 = v770 + v777
		v771 = v782
		goto L236
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	goto L237
L241:
	;
	goto L225
L242:
	;
	return int32(1)
L243:
	;
	v804 = int32(4)
	v805 = l1 + v804
	v807 = v799 + v804
	v809 = v802 - v804
	if base.Ui32(v804) <= base.Ui32(v809) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	return base.B2i32(v871 == int32(0))
L245:
	;
	v871 = int32(0)
	goto L244
L246:
	;
	v845 = v840
	v846 = v841
	v847 = v842
	goto L256
L247:
	;
	if (v805|v807)&int32(3) != 0 {
		v840 = v805
		v841 = v807
		v842 = v809
		goto L246
	} else {
		goto L250
	}
L248:
	;
	v833 = v805
	v834 = v807
	v835 = v809
	goto L249
L249:
	;
	if v835 == int32(0) {
		goto L245
	} else {
		goto L255
	}
L250:
	;
	v817 = v805
	v818 = v807
	v819 = v809
	goto L251
L251:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	if v822 != v823 {
		v840 = v817
		v841 = v818
		v842 = v819
		goto L246
	} else {
		goto L253
	}
L252:
	;
	v833 = v828
	v834 = v826
	v835 = v830
	goto L249
L253:
	;
	v825 = int32(4)
	v826 = v818 + v825
	v828 = v817 + v825
	v830 = v819 - v825
	if base.Ui32(int32(3)) < base.Ui32(v830) {
		v817 = v828
		v818 = v826
		v819 = v830
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v840 = v833
	v841 = v834
	v842 = v835
	goto L246
L256:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846))))
	if v850 == v851 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v871 = v850 - v851
	goto L244
L258:
	;
	v853 = int32(1)
	v858 = v847 - v853
	if v858 != 0 {
		v845 = v845 + v853
		v846 = v846 + v853
		v847 = v858
		goto L256
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	goto L257
L261:
	;
	goto L245
L262:
	;
	if v877 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	return int32(1)
L264:
	;
	goto L265
L265:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v885 = m.T0[v884].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v883, l3, l6)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L11
	} else {
		goto L266
	}
L266:
	;
	v893 = v885 ^ int32(1)
	goto L1
}
func F_gbt_var_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = F_palloc(m, int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v6 + int32(4)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v21)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v20)
			return v11
		}
	}
}
func F_gbt_var_key_readable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = int32(4)
	v5 = l1 + v4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = int32(2)
	v9 = int32(base.Ui32(v7) >> (uint(v8) % 32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v9+v4) < base.Ui32(int32(base.Ui32(v17)>>(uint(v8)%32))) {
		v21 = v5 + (v9+int32(3))&int32(2147483644)
	} else {
		v21 = v5
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
	return
}
func F_gbt_var_node_cp_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = int32(2)
	v12 = int32(base.Ui32(v10) >> (uint(v11) % 32))
	v17 = l0 + (v12+int32(3))&int32(2147483644)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = base.B2i32(base.Ui32(v12+int32(4)) < base.Ui32(int32(base.Ui32(v20)>>(uint(v11)%32))))
	if base.Ui32(v12+int32(4)) < base.Ui32(int32(base.Ui32(v20)>>(uint(v11)%32))) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v17
	goto L3
L2:
	;
	v24 = l0
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = int32(base.Ui32(v25) >> (uint(int32(2)) % 32))
	if base.Ui32(v12) < base.Ui32(v27) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = v12
	goto L6
L5:
	;
	v29 = v27
	goto L6
L6:
	;
	v31 = v29 - int32(4)
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(int32(5)) <= base.Ui32(v29) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(4)
	v41 = l0 + v38
	if base.Ui32(v12+int32(4)) < base.Ui32(int32(base.Ui32(v20)>>(uint(v11)%32))) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	return v31
L13:
	;
	v42 = v17 + v38
	goto L15
L14:
	;
	v42 = v41
	goto L15
L15:
	;
	v49 = int32(0)
	v52 = l0 + int32(8)
	v54 = v42 + int32(4)
	v55 = v49
	v56 = v49
	v57 = v49
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v61 < int32(2) {
		v72 = v56
		v73 = v57
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L12
L18:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v74 != v75 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	if v56 != 0 {
		v72 = v56
		v73 = v57
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = F_pg_mblen_range(m, v52, v12+v41)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v68 = F_pg_mblen_range(m, v54, v42+v27)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v68 == v64 {
		v72 = v64
		v73 = v64
		goto L18
	} else {
		goto L24
	}
L24:
	;
	return v55
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v77 < int32(2) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v84 = int32(1)
	v91 = v55 + v84
	if v91 != v31 {
		v52 = v52 + v84
		v54 = v54 + v84
		v55 = v91
		v56 = v72 - v84
		v57 = v73
		goto L16
	} else {
		goto L31
	}
L28:
	;
	return v55
L29:
	;
	goto L30
L30:
	;
	return v55 - v73 + v72
L31:
	;
	goto L17
}
func F_gbt_var_same(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = int32(4)
	v18 = l0 + v17
	v20 = l1 + v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v22 = m.T0[v21].(func(*base.Module, int32, int32, int32, int32) int32)(m, v18, v20, l2, l4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			v63 = int32(0)
			return v63
		} else {
			v26 = int32(2)
			v27 = int32(base.Ui32(v15) >> (uint(v26) % 32))
			v33 = int32(4)
			if base.Ui32(v27+v33) < base.Ui32(int32(base.Ui32(v16)>>(uint(v26)%32))) {
				v40 = l0 + (v27+int32(3))&int32(2147483644) + v33
			} else {
				v40 = v18
			}
			v41 = int32(2)
			v42 = int32(base.Ui32(v13) >> (uint(v41) % 32))
			v48 = int32(4)
			if base.Ui32(v42+v48) < base.Ui32(int32(base.Ui32(v14)>>(uint(v41)%32))) {
				v55 = l1 + (v42+int32(3))&int32(2147483644) + v48
			} else {
				v55 = v20
			}
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
			v57 = m.T0[v56].(func(*base.Module, int32, int32, int32, int32) int32)(m, v40, v55, l2, l4)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v63 = base.B2i32(v57 == int32(0))
				return v63
			}
		}
	}
}
