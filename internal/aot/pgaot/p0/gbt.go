package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bitle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2857), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_bool_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_bool_sortsupport_0)
	return int64(0)
}
func F_gbt_bool_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_bool_union_0), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_bpcharlt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2607), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_bytea_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14270(m, l0, l1, l2, int32(3062))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_cash_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14276(m, l0, int32(_a_F_gbt_cash_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_cash_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_cash_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_cash_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	v5 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
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
func F_gbt_date_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14279(m, l0, int32(_a_F_gbt_date_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_date_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+4)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+4)))
	v17 = F_DirectFunctionCall2Coll(m, int32(2640), int32(0), v13, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v23 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15))))
		v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12))))
		v25 = F_DirectFunctionCall2Coll(m, int32(2640), int32(0), v23, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = base.I32_wrap_i64(v8)
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = v28
			v30 = base.I32_wrap_i64(v25)
			if v28 < v30 {
				v34 = v30
			} else {
				v34 = v28
			}
			v35 = base.I32_wrap_i64(v17)
			v36 = int32(0)
			if v36 < v35 {
				v39 = v35
			} else {
				v39 = v36
			}
			v40 = v34 + v39
			if v40 != 0 {
				v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+4)))
				v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15))))
				v45 = F_DirectFunctionCall2Coll(m, int32(2640), int32(0), v43, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					v47 = *(*float32)(unsafe.Add(mBase, uint32(v27)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					*(*float32)(unsafe.Add(mBase, uint32(v27))) = base.F32_mul(base.F32_add(base.F32_add(v47, float32(1.1754944e-38)), base.F32_demote_f64(base.F64_div(base.F64_convert_i32_u(v40), base.F64_convert_i32_s(v40+base.I32_wrap_i64(v45))))), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v61+int32(1))))
					return v8 & int64(4294967295)
				}
			} else {
				return v8 & int64(4294967295)
			}
		}
	}
}
func F_gbt_date_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_date_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_dateeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2625), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_datekey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6))))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8))))
	v10 = F_DirectFunctionCall2Coll(m, int32(1576), int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = base.I32_wrap_i64(v10)
		if v14 != 0 {
			v22 = v14
			return v22
		} else {
			v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+4)))
			v19 = F_DirectFunctionCall2Coll(m, int32(1576), int32(0), v17, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = base.I32_wrap_i64(v19)
				return v22
			}
		}
	}
}
func F_gbt_enumgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(4007), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_float4_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14279(m, l0, int32(_a_F_gbt_float4_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_float4_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_float4_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_float4ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v12 float32
	_ = v12
	var v22 int32
	_ = v22
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(base.I32_reinterpret_f32(v6)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
		v22 = base.F32_ge(v6, v12) & base.B2i32(base.Ui32(base.I32_reinterpret_f32(v12)&int32(2147483647)) < base.Ui32(int32(2139095041)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_gbt_float8_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v8, v8+v14, v23&int32(1), int32(_a_F_gbt_float8_distance_0), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.I64_reinterpret_f64(v28)
	}
}
func F_gbt_float8_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v20 int64
	_ = v20
	var v23 float64
	_ = v23
	var v26 int64
	_ = v26
	var v33 float64
	_ = v33
	var v43 float64
	_ = v43
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v66 int64
	_ = v66
	var v70 float64
	_ = v70
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 float64
	_ = v98
	var v106 int64
	_ = v106
	var v113 float64
	_ = v113
	var v118 float64
	_ = v118
	var v124 float64
	_ = v124
	var v141 float32
	_ = v141
	v2 = float64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v20 = base.I64_reinterpret_f64(v17) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v20) {
		v52 = v2
	} else {
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
		v26 = base.I64_reinterpret_f64(v23) & int64(9223372036854775807)
		if base.B2i32(base.F64_gt(v23, v17) == int32(0))&base.B2i32(base.Ui64(v26) < base.Ui64(int64(9218868437227405313))) != 0 {
			v52 = v2
		} else {
			v33 = base.F64_sub(v23, v17)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v26) {
					v43 = float64(3.4028234663852886e+38)
				} else {
					v43 = float64(0)
				}
				v49 = v43
			} else {
				if base.F64_gt(v33, float64(3.4028234663852886e+38)) == int32(0) {
					v49 = v33
				} else {
					v49 = float64(3.4028234663852886e+38)
				}
			}
			v52 = base.F64_add(v49, float64(0))
		}
	}
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) {
		v89 = v52
	} else {
		v66 = base.I64_reinterpret_f64(v55) & int64(9223372036854775807)
		if base.B2i32(base.F64_lt(v14, v55) == int32(0))&base.B2i32(base.Ui64(v66) < base.Ui64(int64(9218868437227405313))) != 0 {
			v89 = v52
		} else {
			v70 = base.F64_sub(v55, v14)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&int64(9223372036854775807)) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v66) {
					v80 = float64(3.4028234663852886e+38)
				} else {
					v80 = float64(0)
				}
				v86 = v80
			} else {
				if base.F64_gt(v70, float64(3.4028234663852886e+38)) == int32(0) {
					v86 = v70
				} else {
					v86 = float64(3.4028234663852886e+38)
				}
			}
			v89 = base.F64_add(v52, v86)
		}
	}
	v91 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	if base.F64_gt(v89, float64(0)) != 0 {
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
		v98 = base.F64_sub(v17, v55)
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) {
			v106 = base.I64_reinterpret_f64(v55) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v20) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v106) {
					v113 = float64(0)
				} else {
					v113 = float64(3.4028234663852886e+38)
				}
				v124 = v113
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v106) {
					v118 = float64(3.4028234663852886e+38)
				} else {
					v118 = float64(0)
				}
				v124 = v118
			}
		} else {
			if base.F64_gt(v98, float64(3.4028234663852886e+38)) == int32(0) {
				v124 = v98
			} else {
				v124 = float64(3.4028234663852886e+38)
			}
		}
		v141 = base.F32_mul(base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v97+int32(1))), base.F32_add(base.F32_demote_f64(base.F64_div(v89, base.F64_add(v89, v124))), float32(1.1754944e-38)))
	} else {
		v141 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v91)))) = v141
	return v91
}
func F_gbt_float8_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_float8_union_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_float8lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		v22 = base.F64_lt(v6, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_gbt_int2_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_int2_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
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
func F_gbt_int8_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14276(m, l0, int32(_a_F_gbt_int8_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_int8_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_int8_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_intv_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+18)))
	if v7 != int32(1) {
		return base.I64_extend_i32_u(v6)
	} else {
		v13 = F_palloc(m, int32(32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v18 = F_palloc(m, int32(24))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+18)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v24
				if v20 != 0 {
					v28 = int32(0)
				} else {
					v28 = int32(16)
				}
				v29 = v21 + v28
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v30
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v32
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = base.I64_extend_i32_u(v13)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v38
				v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+18)) = uint8(v41)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v40)
				return base.I64_extend_i32_u(v18)
			}
		}
	}
}
func F_gbt_intv_consistent(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14280(m, l0, int32(_a_F_gbt_intv_consistent_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
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
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2660), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_intvle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2658), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad8_penalty(m *base.Module, l0 int32) int64 {
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
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v107 float64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v137 float64
	_ = v137
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 float32
	_ = v161
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v13 = int64(48)
	v15 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v16 = int64(56)
	v19 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	v20 = int64(40)
	v23 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)))
	v24 = int64(32)
	v27 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v28 = int64(24)
	v31 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)))
	v32 = int64(16)
	v35 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
	v36 = int64(8)
	v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
	v41 = base.F64_convert_i64_u(v12<<(uint(v13)%64) | v15<<(uint(v16)%64) | v19<<(uint(v20)%64) | v23<<(uint(v24)%64) | v27<<(uint(v28)%64) | v31<<(uint(v32)%64) | v35<<(uint(v36)%64) + v39)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+2)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+3)))
	v59 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+4)))
	v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+5)))
	v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+6)))
	v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+7)))
	v73 = base.F64_convert_i64_u(v44<<(uint(v13)%64) | v47<<(uint(v16)%64) | v51<<(uint(v20)%64) | v55<<(uint(v24)%64) | v59<<(uint(v28)%64) | v63<<(uint(v32)%64) | v67<<(uint(v36)%64) + v71)
	if base.F64_lt(v73, v41) != 0 {
		v77 = base.F64_sub(v41, v73)
	} else {
		v77 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)))
	v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
	v80 = int64(48)
	v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+8)))
	v83 = int64(56)
	v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+10)))
	v87 = int64(40)
	v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+11)))
	v91 = int64(32)
	v94 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+12)))
	v95 = int64(24)
	v98 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+13)))
	v99 = int64(16)
	v102 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+14)))
	v103 = int64(8)
	v107 = base.F64_convert_i64_u(v78 + (v79<<(uint(v80)%64) | v82<<(uint(v83)%64) | v86<<(uint(v87)%64) | v90<<(uint(v91)%64) | v94<<(uint(v95)%64) | v98<<(uint(v99)%64) | v102<<(uint(v103)%64)))
	v108 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	v109 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+9)))
	v112 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)))
	v116 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	v124 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
	v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	v137 = base.F64_convert_i64_u(v108 + (v109<<(uint(v80)%64) | v112<<(uint(v83)%64) | v116<<(uint(v87)%64) | v120<<(uint(v91)%64) | v124<<(uint(v95)%64) | v128<<(uint(v99)%64) | v132<<(uint(v103)%64)))
	if base.F64_lt(v137, v107) != 0 {
		v141 = base.F64_sub(v107, v137)
	} else {
		v141 = float64(0)
	}
	v142 = base.F64_add(v77, v141)
	if base.F64_gt(v142, float64(0)) != 0 {
		v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
		v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
		v161 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v142, base.F64_add(base.F64_sub(v137, v41), v142))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v154+int32(1))))
	} else {
		v161 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = v161
	return v8
}
func F_gbt_macad8_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_macad8_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macad8lt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8lt_0), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_macad_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macad_union(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14282(m, l0, int32(_a_F_gbt_macad_union_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_macadge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2484), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macadkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14281(m, l0, l1, l2, int32(6), int32(2486))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_numeric_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(1467), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v9)
	}
}
func F_gbt_numeric_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14269(m, l0, int32(_a_F_gbt_numeric_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_oid_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_oid_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
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
func F_gbt_text_consistent(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14283(m, l0, int32(_a_F_gbt_text_consistent_0), int32(_a_F_gbt_text_consistent_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_textle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2449), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_time_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14276(m, l0, int32(_a_F_gbt_time_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_time_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_DirectFunctionCall2Coll(m, int32(2919), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return float64(0)
	} else {
		v12 = base.I32_wrap_i64(v8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		return base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v13), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v17), float64(86400)), base.F64_div(base.F64_convert_i64_s(v21), float64(1e+06)))))
	}
}
func F_gbt_timegt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2633), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_ts_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_ts_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_ts_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v41 float64
	_ = v41
	v7 = math.Float64frombits(uint64(0x7ff0000000000000))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v8-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v41 = v7
		return v41
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		if base.Ui64(v13-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			v41 = v7
			return v41
		} else {
			v20 = F_DirectFunctionCall2Coll(m, int32(1714), int32(0), v8, v13)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return float64(0)
			} else {
				v24 = base.I32_wrap_i64(v20)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
				v41 = base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v25), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v29), float64(86400)), base.F64_div(base.F64_convert_i64_s(v33), float64(1e+06)))))
				return v41
			}
		}
	}
}
func F_gbt_tstz_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+30)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	v18 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v14 + v18
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v10
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v30)+12)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = F_gbt_num_consistent(m, v8+int32(20), v8+v18, v8+int32(30), v32&int32(1), int32(_a_F_gbt_tstz_consistent_0), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(32)
		return base.I64_extend_i32_u(v37)
	}
}
func F_gbt_uuid_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_uuid_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
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
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v609 int32
	_ = v609
	v8 = int32(0)
	switch l2 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L9
	case 2:
		goto L7
	case 3:
		goto L5
	case 4:
		goto L6
	case 5:
		goto L4
	default:
		v529 = v8
		goto L3
	}
L1:
	;
	v542 = int32(4)
	v543 = l1 + v542
	v545 = v537 + v542
	v547 = v538 - v542
	if base.Ui32(v542) <= base.Ui32(v547) {
		goto L181
	} else {
		goto L182
	}
L2:
	;
	return int32(1)
L3:
	;
	return v529
L4:
	;
	if l4 != 0 {
		goto L170
	} else {
		goto L171
	}
L5:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l4 != 0 {
		goto L139
	} else {
		goto L140
	}
L6:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l4 != 0 {
		goto L108
	} else {
		goto L109
	}
L7:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L73
	} else {
		goto L74
	}
L8:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L42
	} else {
		goto L43
	}
L9:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v13 = m.T0[v12].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v11, l3, l6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v11, l3, l6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	return v13
L15:
	;
	if int32(0) <= v19 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v23 != int32(1) {
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = int32(2)
	v31 = int32(base.Ui32(v29) >> (uint(v30) % 32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = int32(base.Ui32(v32) >> (uint(v30) % 32))
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = int32(base.Ui32(v105) >> (uint(int32(2)) % 32))
	if base.Ui32(v34) < base.Ui32(v107) {
		v529 = v8
		goto L3
	} else {
		goto L41
	}
L21:
	;
	v36 = int32(4)
	v37 = l1 + v36
	v39 = v28 + v36
	v41 = v31 - v36
	if base.Ui32(v36) <= base.Ui32(v41) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	if v103 != 0 {
		goto L20
	} else {
		goto L40
	}
L23:
	;
	v103 = int32(0)
	goto L22
L24:
	;
	v77 = v72
	v78 = v73
	v79 = v74
	goto L34
L25:
	;
	if (v37|v39)&int32(3) != 0 {
		v72 = v37
		v73 = v39
		v74 = v41
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v65 = v37
	v66 = v39
	v67 = v41
	goto L27
L27:
	;
	if v67 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v49 = v37
	v50 = v39
	v51 = v41
	goto L29
L29:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v54 != v55 {
		v72 = v49
		v73 = v50
		v74 = v51
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v65 = v60
	v66 = v58
	v67 = v62
	goto L27
L31:
	;
	v57 = int32(4)
	v58 = v50 + v57
	v60 = v49 + v57
	v62 = v51 - v57
	if base.Ui32(int32(3)) < base.Ui32(v62) {
		v49 = v60
		v50 = v58
		v51 = v62
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v72 = v65
	v73 = v66
	v74 = v67
	goto L24
L34:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 == v83 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v103 = v82 - v83
	goto L22
L36:
	;
	v85 = int32(1)
	v90 = v79 - v85
	if v90 != 0 {
		v77 = v77 + v85
		v78 = v78 + v85
		v79 = v90
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
	goto L2
L41:
	;
	v537 = v104
	v538 = v107
	goto L1
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v111 = m.T0[v110].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v109, l3, l6)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v109, l3, l6)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L46
	}
L45:
	;
	return v111
L46:
	;
	if int32(0) <= v115 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v119 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	goto L50
L50:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(2)
	v127 = int32(base.Ui32(v125) >> (uint(v126) % 32))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v130 = int32(base.Ui32(v128) >> (uint(v126) % 32))
	if base.Ui32(v130) < base.Ui32(v127) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v203 = int32(base.Ui32(v201) >> (uint(int32(2)) % 32))
	if base.Ui32(v130) < base.Ui32(v203) {
		v529 = v8
		goto L3
	} else {
		goto L72
	}
L52:
	;
	v132 = int32(4)
	v133 = l1 + v132
	v135 = v124 + v132
	v137 = v127 - v132
	if base.Ui32(v132) <= base.Ui32(v137) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v199 != 0 {
		goto L51
	} else {
		goto L71
	}
L54:
	;
	v199 = int32(0)
	goto L53
L55:
	;
	v173 = v168
	v174 = v169
	v175 = v170
	goto L65
L56:
	;
	if (v133|v135)&int32(3) != 0 {
		v168 = v133
		v169 = v135
		v170 = v137
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v161 = v133
	v162 = v135
	v163 = v137
	goto L58
L58:
	;
	if v163 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	v145 = v133
	v146 = v135
	v147 = v137
	goto L60
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v150 != v151 {
		v168 = v145
		v169 = v146
		v170 = v147
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v161 = v156
	v162 = v154
	v163 = v158
	goto L58
L62:
	;
	v153 = int32(4)
	v154 = v146 + v153
	v156 = v145 + v153
	v158 = v147 - v153
	if base.Ui32(int32(3)) < base.Ui32(v158) {
		v145 = v156
		v146 = v154
		v147 = v158
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v168 = v161
	v169 = v162
	v170 = v163
	goto L55
L65:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v178 == v179 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v199 = v178 - v179
	goto L53
L67:
	;
	v181 = int32(1)
	v186 = v175 - v181
	if v186 != 0 {
		v173 = v173 + v181
		v174 = v174 + v181
		v175 = v186
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L54
L71:
	;
	goto L2
L72:
	;
	v537 = v200
	v538 = v203
	goto L1
L73:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v207 = m.T0[v206].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v205, l3, l6)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v211 = m.T0[v210].(func(*base.Module, int32, int32, int32, int32) int32)(m, v205, l1, l3, l6)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L77
	}
L76:
	;
	return v207
L77:
	;
	if v211 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v218 = m.T0[v217].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v216, l3, l6)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v223 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v218 <= int32(0) {
		v529 = int32(1)
		goto L3
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	return int32(0)
L84:
	;
	goto L85
L85:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v230 = int32(2)
	v231 = int32(base.Ui32(v229) >> (uint(v230) % 32))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v234 = int32(base.Ui32(v232) >> (uint(v230) % 32))
	if base.Ui32(v234) < base.Ui32(v231) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v308 = int32(base.Ui32(v306) >> (uint(int32(2)) % 32))
	if base.Ui32(v234) < base.Ui32(v308) {
		v529 = int32(0)
		goto L3
	} else {
		goto L107
	}
L87:
	;
	v236 = int32(4)
	v237 = l1 + v236
	v239 = v228 + v236
	v241 = v231 - v236
	if base.Ui32(v236) <= base.Ui32(v241) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if v303 != 0 {
		goto L86
	} else {
		goto L106
	}
L89:
	;
	v303 = int32(0)
	goto L88
L90:
	;
	v277 = v272
	v278 = v273
	v279 = v274
	goto L100
L91:
	;
	if (v237|v239)&int32(3) != 0 {
		v272 = v237
		v273 = v239
		v274 = v241
		goto L90
	} else {
		goto L94
	}
L92:
	;
	v265 = v237
	v266 = v239
	v267 = v241
	goto L93
L93:
	;
	if v267 == int32(0) {
		goto L89
	} else {
		goto L99
	}
L94:
	;
	v249 = v237
	v250 = v239
	v251 = v241
	goto L95
L95:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v254 != v255 {
		v272 = v249
		v273 = v250
		v274 = v251
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v265 = v260
	v266 = v258
	v267 = v262
	goto L93
L97:
	;
	v257 = int32(4)
	v258 = v250 + v257
	v260 = v249 + v257
	v262 = v251 - v257
	if base.Ui32(int32(3)) < base.Ui32(v262) {
		v249 = v260
		v250 = v258
		v251 = v262
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v272 = v265
	v273 = v266
	v274 = v267
	goto L90
L100:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v282 == v283 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v303 = v282 - v283
	goto L88
L102:
	;
	v285 = int32(1)
	v290 = v279 - v285
	if v290 != 0 {
		v277 = v277 + v285
		v278 = v278 + v285
		v279 = v290
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	goto L89
L106:
	;
	goto L2
L107:
	;
	v537 = v305
	v538 = v308
	goto L1
L108:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v312 = m.T0[v311].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v310, l3, l6)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L13
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v317 = m.T0[v316].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v310, l3, l6)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L112
	}
L111:
	;
	return v312
L112:
	;
	if v317 <= int32(0) {
		v529 = int32(1)
		goto L3
	} else {
		goto L113
	}
L113:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v321 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	return int32(0)
L115:
	;
	goto L116
L116:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v328 = int32(2)
	v329 = int32(base.Ui32(v327) >> (uint(v328) % 32))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v332 = int32(base.Ui32(v330) >> (uint(v328) % 32))
	if base.Ui32(v332) < base.Ui32(v329) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v406 = int32(base.Ui32(v404) >> (uint(int32(2)) % 32))
	if base.Ui32(v332) < base.Ui32(v406) {
		v529 = int32(0)
		goto L3
	} else {
		goto L138
	}
L118:
	;
	v334 = int32(4)
	v335 = l1 + v334
	v337 = v326 + v334
	v339 = v329 - v334
	if base.Ui32(v334) <= base.Ui32(v339) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	if v401 != 0 {
		goto L117
	} else {
		goto L137
	}
L120:
	;
	v401 = int32(0)
	goto L119
L121:
	;
	v375 = v370
	v376 = v371
	v377 = v372
	goto L131
L122:
	;
	if (v335|v337)&int32(3) != 0 {
		v370 = v335
		v371 = v337
		v372 = v339
		goto L121
	} else {
		goto L125
	}
L123:
	;
	v363 = v335
	v364 = v337
	v365 = v339
	goto L124
L124:
	;
	if v365 == int32(0) {
		goto L120
	} else {
		goto L130
	}
L125:
	;
	v347 = v335
	v348 = v337
	v349 = v339
	goto L126
L126:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	if v352 != v353 {
		v370 = v347
		v371 = v348
		v372 = v349
		goto L121
	} else {
		goto L128
	}
L127:
	;
	v363 = v358
	v364 = v356
	v365 = v360
	goto L124
L128:
	;
	v355 = int32(4)
	v356 = v348 + v355
	v358 = v347 + v355
	v360 = v349 - v355
	if base.Ui32(int32(3)) < base.Ui32(v360) {
		v347 = v358
		v348 = v356
		v349 = v360
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v370 = v363
	v371 = v364
	v372 = v365
	goto L121
L131:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v380 == v381 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v401 = v380 - v381
	goto L119
L133:
	;
	v383 = int32(1)
	v388 = v377 - v383
	if v388 != 0 {
		v375 = v375 + v383
		v376 = v376 + v383
		v377 = v388
		goto L131
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	goto L120
L137:
	;
	goto L2
L138:
	;
	v537 = v403
	v538 = v406
	goto L1
L139:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v410 = m.T0[v409].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v408, l3, l6)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v415 = m.T0[v414].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v408, l3, l6)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L13
	} else {
		goto L143
	}
L142:
	;
	return v410
L143:
	;
	if v415 <= int32(0) {
		v529 = int32(1)
		goto L3
	} else {
		goto L144
	}
L144:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v419 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	return int32(0)
L146:
	;
	goto L147
L147:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = int32(2)
	v427 = int32(base.Ui32(v425) >> (uint(v426) % 32))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v430 = int32(base.Ui32(v428) >> (uint(v426) % 32))
	if base.Ui32(v430) < base.Ui32(v427) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	v504 = int32(base.Ui32(v502) >> (uint(int32(2)) % 32))
	if base.Ui32(v430) < base.Ui32(v504) {
		v529 = int32(0)
		goto L3
	} else {
		goto L169
	}
L149:
	;
	v432 = int32(4)
	v433 = l1 + v432
	v435 = v424 + v432
	v437 = v427 - v432
	if base.Ui32(v432) <= base.Ui32(v437) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	if v499 != 0 {
		goto L148
	} else {
		goto L168
	}
L151:
	;
	v499 = int32(0)
	goto L150
L152:
	;
	v473 = v468
	v474 = v469
	v475 = v470
	goto L162
L153:
	;
	if (v433|v435)&int32(3) != 0 {
		v468 = v433
		v469 = v435
		v470 = v437
		goto L152
	} else {
		goto L156
	}
L154:
	;
	v461 = v433
	v462 = v435
	v463 = v437
	goto L155
L155:
	;
	if v463 == int32(0) {
		goto L151
	} else {
		goto L161
	}
L156:
	;
	v445 = v433
	v446 = v435
	v447 = v437
	goto L157
L157:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	if v450 != v451 {
		v468 = v445
		v469 = v446
		v470 = v447
		goto L152
	} else {
		goto L159
	}
L158:
	;
	v461 = v456
	v462 = v454
	v463 = v458
	goto L155
L159:
	;
	v453 = int32(4)
	v454 = v446 + v453
	v456 = v445 + v453
	v458 = v447 - v453
	if base.Ui32(int32(3)) < base.Ui32(v458) {
		v445 = v456
		v446 = v454
		v447 = v458
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v468 = v461
	v469 = v462
	v470 = v463
	goto L152
L162:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v478 == v479 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v499 = v478 - v479
	goto L150
L164:
	;
	v481 = int32(1)
	v486 = v475 - v481
	if v486 != 0 {
		v473 = v473 + v481
		v474 = v474 + v481
		v475 = v486
		goto L162
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	goto L163
L167:
	;
	goto L151
L168:
	;
	goto L2
L169:
	;
	v537 = v501
	v538 = v504
	goto L1
L170:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v508 = m.T0[v507].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v506, l3, l6)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L13
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v513 = int32(1)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	if v514 != 0 {
		v529 = v513
		goto L3
	} else {
		goto L174
	}
L173:
	;
	return v508 ^ int32(1)
L174:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v517 = m.T0[v516].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v515, l3, l6)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	if v517 != 0 {
		v529 = v513
		goto L3
	} else {
		goto L176
	}
L176:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v521 = m.T0[v520].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v519, l3, l6)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L13
	} else {
		goto L177
	}
L177:
	;
	v529 = base.B2i32(v521 != int32(0))
	goto L3
L178:
	;
	return base.B2i32(v609 == int32(0))
L179:
	;
	v609 = int32(0)
	goto L178
L180:
	;
	v583 = v578
	v584 = v579
	v585 = v580
	goto L190
L181:
	;
	if (v543|v545)&int32(3) != 0 {
		v578 = v543
		v579 = v545
		v580 = v547
		goto L180
	} else {
		goto L184
	}
L182:
	;
	v571 = v543
	v572 = v545
	v573 = v547
	goto L183
L183:
	;
	if v573 == int32(0) {
		goto L179
	} else {
		goto L189
	}
L184:
	;
	v555 = v543
	v556 = v545
	v557 = v547
	goto L185
L185:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	if v560 != v561 {
		v578 = v555
		v579 = v556
		v580 = v557
		goto L180
	} else {
		goto L187
	}
L186:
	;
	v571 = v566
	v572 = v564
	v573 = v568
	goto L183
L187:
	;
	v563 = int32(4)
	v564 = v556 + v563
	v566 = v555 + v563
	v568 = v557 - v563
	if base.Ui32(int32(3)) < base.Ui32(v568) {
		v555 = v566
		v556 = v564
		v557 = v568
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v578 = v571
	v579 = v572
	v580 = v573
	goto L180
L190:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	if v588 == v589 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v609 = v588 - v589
	goto L178
L192:
	;
	v591 = int32(1)
	v596 = v585 - v591
	if v596 != 0 {
		v583 = v583 + v591
		v584 = v584 + v591
		v585 = v596
		goto L190
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	goto L191
L195:
	;
	goto L179
}
func F_gbt_var_fetch(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v11 = F_palloc(m, int32(24))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = base.I64_extend_i32_u(v6 + int32(4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)))
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)) = uint8(v22)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v21)
			return base.I64_extend_i32_u(v11)
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = int32(2)
	v12 = int32(base.Ui32(v10) >> (uint(v11) % 32))
	v13 = int32(4)
	v14 = v12 - v13
	v19 = l0 + (v12+int32(3))&int32(2147483644)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = base.B2i32(base.Ui32(v12+v13) < base.Ui32(int32(base.Ui32(v22)>>(uint(v11)%32))))
	if base.Ui32(v12+v13) < base.Ui32(int32(base.Ui32(v22)>>(uint(v11)%32))) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v19
	goto L3
L2:
	;
	v26 = l0
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
	v31 = v29 - int32(4)
	if v14 < v31 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = v14
	goto L6
L5:
	;
	v33 = v31
	goto L6
L6:
	;
	if v33 == int32(0) {
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
	if int32(0) < v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = int32(4)
	if base.Ui32(v12+v13) < base.Ui32(int32(base.Ui32(v22)>>(uint(v11)%32))) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	return v33
L13:
	;
	v44 = v19 + v40
	goto L15
L14:
	;
	v44 = l0 + v40
	goto L15
L15:
	;
	v47 = l0 + int32(8)
	v51 = int32(0)
	v54 = v47
	v56 = v51
	v57 = v51
	v58 = v44 + int32(4)
	v59 = v51
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v63 < int32(2))|v57 != 0 {
		v75 = v57
		v76 = v59
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L12
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v77 != v78 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v67 = F_pg_mblen_range(m, v54, v47+v14)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v71 = F_pg_mblen_range(m, v58, v44+v29)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v71 == v67 {
		v75 = v67
		v76 = v67
		goto L18
	} else {
		goto L23
	}
L23:
	;
	return v56
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v80 < int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v87 = int32(1)
	v94 = v56 + v87
	if v94 != v33 {
		v54 = v54 + v87
		v56 = v94
		v57 = v75 - v87
		v58 = v58 + v87
		v59 = v76
		goto L16
	} else {
		goto L30
	}
L27:
	;
	return v56
L28:
	;
	goto L29
L29:
	;
	return v56 - v76 + v75
L30:
	;
	goto L17
}
func F_gbt_var_same(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v14 = base.I32_wrap_i64(l1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v17 = base.I32_wrap_i64(l0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v20 = int32(4)
	v21 = v17 + v20
	v23 = v14 + v20
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v25 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, v23, l2, l4)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if v25 != 0 {
			v65 = int32(1)
			return base.B2i32(v65 == int32(0))
		} else {
			v30 = int32(2)
			v31 = int32(base.Ui32(v18) >> (uint(v30) % 32))
			v37 = int32(4)
			if base.Ui32(v31+v37) < base.Ui32(int32(base.Ui32(v19)>>(uint(v30)%32))) {
				v44 = v17 + (v31+int32(3))&int32(2147483644) + v37
			} else {
				v44 = v21
			}
			v45 = int32(2)
			v46 = int32(base.Ui32(v15) >> (uint(v45) % 32))
			v52 = int32(4)
			if base.Ui32(v46+v52) < base.Ui32(int32(base.Ui32(v16)>>(uint(v45)%32))) {
				v59 = v14 + (v46+int32(3))&int32(2147483644) + v52
			} else {
				v59 = v23
			}
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
			v61 = m.T0[v60].(func(*base.Module, int32, int32, int32, int32) int32)(m, v44, v59, l2, l4)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v65 = v61
				return base.B2i32(v65 == int32(0))
			}
		}
	}
}
