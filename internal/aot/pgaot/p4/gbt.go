package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gbt_bit_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_bit_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_bit_consistent(m *base.Module, l0 int32) int64 {
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)))
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47)+12)))
		if v49&int32(1) != 0 {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v56 = F_gbt_var_consistent(m, v22, v15, v19, v52, int32(1), int32(_a_F_gbt_bit_consistent_0), v55)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				v104 = v56
				m.G0 = v11 + int32(16)
				return base.I64_extend_i32_u(v104)
			}
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v60 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
			v64 = (v60 - int32(1)) & int32(-4)
			v65 = F_palloc(m, v64)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int64(0)
			} else {
				v68 = v60 - int32(4)
				if v64 <= v68 {
				} else {
					v72 = v64 - v60 + int32(4)
					if v72 == int32(0) {
					} else {
						base.MemoryFill(m, v65+v68, int32(0), v72)
					}
				}
				v79 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v65))) = v64 << (uint(v79) % 32)
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v86 = int32(base.Ui32(v82)>>(uint(v79)%32)) - int32(8)
				if v86 != 0 {
					base.MemoryCopy(m, v65+int32(4), v15+int32(8), v86)
				} else {
				}
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v98 = F_gbt_var_consistent(m, v11+int32(8), v65, v19, v94, int32(0), int32(_a_F_gbt_bit_consistent_0), v97)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int64(0)
				} else {
					v104 = v98
					m.G0 = v11 + int32(16)
					return base.I64_extend_i32_u(v104)
				}
			}
		}
	}
}
func F_gbt_bitgt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(2856), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_boolkey_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v6 == v8 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
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
func F_gbt_bpcharge(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(2610), l2, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_bytea_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_bytea_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_byteale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(3058), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_cash_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_cash_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_cash_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_abs(base.F64_sub(base.F64_convert_i64_s(v4), base.F64_convert_i64_s(v6)))
}
func F_gbt_enum_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_enum_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_enumle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(4008), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_float4_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 float32
	_ = v18
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 float32
	_ = v22
	var v23 float32
	_ = v23
	var v24 float64
	_ = v24
	var v27 int64
	_ = v27
	var v30 float32
	_ = v30
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v41 float64
	_ = v41
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v64 float64
	_ = v64
	var v75 int64
	_ = v75
	var v79 float64
	_ = v79
	var v89 float64
	_ = v89
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v115 int64
	_ = v115
	var v122 float64
	_ = v122
	var v127 float64
	_ = v127
	var v133 float64
	_ = v133
	var v150 float32
	_ = v150
	v2 = float64(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*float32)(unsafe.Add(mBase, uint32(v17)))
	v19 = base.F64_promote_f32(v18)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*float32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*float32)(unsafe.Add(mBase, uint32(v21)+4))
	v24 = base.F64_promote_f32(v23)
	v27 = base.I64_reinterpret_f64(v24) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v27) {
		v60 = v2
	} else {
		v30 = *(*float32)(unsafe.Add(mBase, uint32(v17)+4))
		v31 = base.F64_promote_f32(v30)
		v34 = base.I64_reinterpret_f64(v31) & int64(9223372036854775807)
		if base.B2i32(base.F32_lt(v23, v30) == int32(0))&base.B2i32(base.Ui64(v34) < base.Ui64(int64(9218868437227405313))) != 0 {
			v60 = v2
		} else {
			v41 = base.F64_sub(v31, v24)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v41)&int64(9223372036854775807)) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v34) {
					v51 = float64(3.4028234663852886e+38)
				} else {
					v51 = float64(0)
				}
				v57 = v51
			} else {
				if base.F64_gt(v41, float64(3.4028234663852886e+38)) == int32(0) {
					v57 = v41
				} else {
					v57 = float64(3.4028234663852886e+38)
				}
			}
			v60 = base.F64_add(v57, float64(0))
		}
	}
	v64 = base.F64_promote_f32(v22)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
		v98 = v60
	} else {
		v75 = base.I64_reinterpret_f64(v64) & int64(9223372036854775807)
		if base.B2i32(base.F32_gt(v22, v18) == int32(0))&base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) != 0 {
			v98 = v60
		} else {
			v79 = base.F64_sub(v64, v19)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v79)&int64(9223372036854775807)) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v75) {
					v89 = float64(3.4028234663852886e+38)
				} else {
					v89 = float64(0)
				}
				v95 = v89
			} else {
				if base.F64_gt(v79, float64(3.4028234663852886e+38)) == int32(0) {
					v95 = v79
				} else {
					v95 = float64(3.4028234663852886e+38)
				}
			}
			v98 = base.F64_add(v60, v95)
		}
	}
	v100 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	if base.F64_gt(v98, float64(0)) != 0 {
		v104 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
		v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+52))
		v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
		v107 = base.F64_sub(v24, v64)
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) {
			v115 = base.I64_reinterpret_f64(v64) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v27) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v115) {
					v122 = float64(0)
				} else {
					v122 = float64(3.4028234663852886e+38)
				}
				v133 = v122
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v115) {
					v127 = float64(3.4028234663852886e+38)
				} else {
					v127 = float64(0)
				}
				v133 = v127
			}
		} else {
			if base.F64_gt(v107, float64(3.4028234663852886e+38)) == int32(0) {
				v133 = v107
			} else {
				v133 = float64(3.4028234663852886e+38)
			}
		}
		v150 = base.F32_mul(base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v106+int32(1))), base.F32_add(base.F32_demote_f64(base.F64_div(v98, base.F64_add(v98, v133))), float32(1.1754944e-38)))
	} else {
		v150 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v100)))) = v150
	return v100
}
func F_gbt_float4_union(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14274(m, l0, int32(_a_F_gbt_float4_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_gbt_float4key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 float32
	_ = v36
	var v37 float32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)))
	v12 = int32(2147483647)
	v13 = base.I32_reinterpret_f32(v5) & v12
	v16 = base.I32_reinterpret_f32(v7) & v12
	if base.Ui32(int32(2139095041)) <= base.Ui32(v16) {
		v26 = base.B2i32(base.Ui32(v13) < base.Ui32(int32(2139095041)))
		v35 = int32(0) - v26&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v16))|base.F32_lt(v5, v7))
	} else {
		v21 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_gt(v5, v7) != 0 {
			v35 = v21
		} else {
			v26 = v21
			v35 = int32(0) - v26&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v16))|base.F32_lt(v5, v7))
		}
	}
	if v35 != 0 {
		v66 = v35
	} else {
		v36 = *(*float32)(unsafe.Add(mBase, uint32(v4)+4))
		v37 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
		v42 = int32(2147483647)
		v43 = base.I32_reinterpret_f32(v36) & v42
		v46 = base.I32_reinterpret_f32(v37) & v42
		if base.Ui32(int32(2139095041)) <= base.Ui32(v46) {
			v56 = base.B2i32(base.Ui32(v43) < base.Ui32(int32(2139095041)))
			v65 = int32(0) - v56&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v46))|base.F32_lt(v36, v37))
		} else {
			v51 = int32(1)
			if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v43))|base.F32_gt(v36, v37) != 0 {
				v65 = v51
			} else {
				v56 = v51
				v65 = int32(0) - v56&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v46))|base.F32_lt(v36, v37))
			}
		}
		v66 = v65
	}
	return v66
}
func F_gbt_float8_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_float8_sortsupport_0)
	return int64(0)
}
func F_gbt_float8key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v12 = int64(9223372036854775807)
	v13 = base.I64_reinterpret_f64(v5) & v12
	v16 = base.I64_reinterpret_f64(v7) & v12
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v16) {
		v26 = base.B2i32(base.Ui64(v13) < base.Ui64(int64(9218868437227405313)))
		v35 = int32(0) - v26&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v16))|base.F64_lt(v5, v7))
	} else {
		v21 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_gt(v5, v7) != 0 {
			v35 = v21
		} else {
			v26 = v21
			v35 = int32(0) - v26&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v16))|base.F64_lt(v5, v7))
		}
	}
	if v35 != 0 {
		v66 = v35
	} else {
		v36 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
		v37 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v42 = int64(9223372036854775807)
		v43 = base.I64_reinterpret_f64(v36) & v42
		v46 = base.I64_reinterpret_f64(v37) & v42
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v46) {
			v56 = base.B2i32(base.Ui64(v43) < base.Ui64(int64(9218868437227405313)))
			v65 = int32(0) - v56&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v46))|base.F64_lt(v36, v37))
		} else {
			v51 = int32(1)
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v43))|base.F64_gt(v36, v37) != 0 {
				v65 = v51
			} else {
				v56 = v51
				v65 = int32(0) - v56&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v46))|base.F64_lt(v36, v37))
			}
		}
		v66 = v65
	}
	return v66
}
func F_gbt_inet_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_inet_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_inet_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_inet_sortsupport_0)
	return int64(0)
}
func F_gbt_inetge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func F_gbt_int2_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int2_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_int2_consistent(m *base.Module, l0 int32) int64 {
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
	var v10 int64
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
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v10)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v14 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v29)+12)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_gbt_num_consistent(m, v7+int32(4), v7+int32(14), v7+int32(12), v31&int32(1), int32(_a_F_gbt_int2_consistent_0), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v36)
	}
}
func F_gbt_int2_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int2_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_int4_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int4_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_int4_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.F64_abs(base.F64_sub(base.F64_convert_i32_s(v4), base.F64_convert_i32_s(v6)))
}
func F_gbt_int4_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int4_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_int4_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v6 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1))))
	return base.B2i32(v8 < v6) - base.B2i32(v6 < v8)
}
func F_gbt_int4le(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v4 <= v5)
}
func F_gbt_int8_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int8_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_int8_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int8_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_intv_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_intv_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_macad8_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_macad8_sortsupport_0)
	return int64(0)
}
func F_gbt_macad8ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8ge_0), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_macad_penalty(m *base.Module, l0 int32) int64 {
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
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 float64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v100 int64
	_ = v100
	var v105 float64
	_ = v105
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 float32
	_ = v129
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v13 = int64(32)
	v15 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v16 = int64(40)
	v19 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	v20 = int64(24)
	v23 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)))
	v24 = int64(16)
	v27 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v28 = int64(8)
	v31 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+5)))
	v33 = base.F64_convert_i64_u(v12<<(uint(v13)%64) | v15<<(uint(v16)%64) | v19<<(uint(v20)%64) | v23<<(uint(v24)%64) | v27<<(uint(v28)%64) | v31)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	v39 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
	v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+3)))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+5)))
	v57 = base.F64_convert_i64_u(v36<<(uint(v13)%64) | v39<<(uint(v16)%64) | v43<<(uint(v20)%64) | v47<<(uint(v24)%64) | v51<<(uint(v28)%64) | v55)
	if base.F64_lt(v57, v33) != 0 {
		v61 = base.F64_sub(v33, v57)
	} else {
		v61 = math.Float64frombits(uint64(0x8000000000000000))
	}
	v62 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+11)))
	v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+7)))
	v64 = int64(32)
	v66 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+6)))
	v67 = int64(40)
	v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	v71 = int64(24)
	v74 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+9)))
	v75 = int64(16)
	v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+10)))
	v79 = int64(8)
	v83 = base.F64_convert_i64_u(v62 | (v63<<(uint(v64)%64) | v66<<(uint(v67)%64) | v70<<(uint(v71)%64) | v74<<(uint(v75)%64) | v78<<(uint(v79)%64)))
	v84 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
	v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+6)))
	v92 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)))
	v96 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+9)))
	v100 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	v105 = base.F64_convert_i64_u(v84 | (v85<<(uint(v64)%64) | v88<<(uint(v67)%64) | v92<<(uint(v71)%64) | v96<<(uint(v75)%64) | v100<<(uint(v79)%64)))
	if base.F64_lt(v105, v83) != 0 {
		v109 = base.F64_sub(v83, v105)
	} else {
		v109 = float64(0)
	}
	v110 = base.F64_add(v61, v109)
	if base.F64_gt(v110, float64(0)) != 0 {
		v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+52))
		v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
		v129 = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v110, base.F64_add(base.F64_sub(v105, v33), v110))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v122+int32(1))))
	} else {
		v129 = float32(0)
	}
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = v129
	return v8
}
func F_gbt_macaddr_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_macaddr_sortsupport_0)
	return int64(0)
}
func F_gbt_num_bin_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v9 = l1 + v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v14 = F_palloc0(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v14)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v18 != 0 {
				base.MemoryCopy(m, v14, l1, v18)
			} else {
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v20 == int32(0) {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, v23+v20, v9, v20)
				return
			}
		}
	} else {
		v26 = v10 + v8
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v28 = m.T0[v27].(func(*base.Module, int32, int32, int32) int32)(m, v10, l1, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v32 == int32(0) {
				} else {
					base.MemoryCopy(m, v10, l1, v32)
				}
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
			v38 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v26, v9, l3)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 == int32(0) {
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v42 == int32(0) {
					} else {
						base.MemoryCopy(m, v26, v9, v42)
					}
				}
				return
			}
		}
	}
}
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = (v15 - int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	v24 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = v19<<(uint(int32(1))%32) + int32(4)
	v32 = F_palloc(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
	v35 = F_palloc(m, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v35
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
	if v15&int32(_a_F_gbt_num_picksplit_0) == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v50 = int32(8)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v24+v50, v19, v50, v53, l3)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v60 = int32(1)
	goto L9
L8:
	;
	return l1
L9:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)+v60*int32(24))))
	v80 = v24 + v60<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v60
	*(*uint32)(unsafe.Add(mBase, uint32(v80)+4)) = uint32(v77)
	v86 = (v60 + int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	if base.Ui32(v86) <= base.Ui32(v19) {
		v60 = v86
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v88 = int32(8)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v24+v88, v19, v88, v91, l3)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v94 = int32(1)
	v97 = v94
	goto L13
L13:
	;
	v113 = v24 + v97<<(uint(int32(3))%32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v116 = v114 + v115
	if base.Ui32(v97) <= base.Ui32(int32(base.Ui32(v19)>>(uint(v94)%32))) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return l1
L15:
	;
	v225 = (v97 + int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	if base.Ui32(v225) <= base.Ui32(v19) {
		v97 = v225
		goto L13
	} else {
		goto L51
	}
L16:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v118 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v168 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v160 = int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*uint16)(unsafe.Add(mBase, uint32(v158+v159<<(uint(v160)%32)))) = uint16(v163)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v159 + v160
	goto L15
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v122 = F_palloc0(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v134 = v115 + v118
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v136 = m.T0[v135].(func(*base.Module, int32, int32, int32) int32)(m, v118, v114, l3)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = base.I64_extend_i32_u(v122)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v126 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	base.MemoryCopy(m, v122, v114, v126)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v128 == int32(0) {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	base.MemoryCopy(m, v131+v128, v116, v128)
	goto L19
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v146 = m.T0[v145].(func(*base.Module, int32, int32, int32) int32)(m, v134, v116, l3)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	if v136 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v140 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	base.MemoryCopy(m, v118, v114, v140)
	goto L28
L32:
	;
	if v146 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 == int32(0) {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.MemoryCopy(m, v134, v116, v150)
	goto L19
L35:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v210 = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*uint16)(unsafe.Add(mBase, uint32(v208+v209<<(uint(v210)%32)))) = uint16(v213)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v209 + v210
	goto L15
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v172 = F_palloc0(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v184 = v115 + v168
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v186 = m.T0[v185].(func(*base.Module, int32, int32, int32) int32)(m, v168, v114, l3)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = base.I64_extend_i32_u(v172)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v176 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryCopy(m, v172, v114, v176)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v178 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	base.MemoryCopy(m, v181+v178, v116, v178)
	goto L35
L44:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v196 = m.T0[v195].(func(*base.Module, int32, int32, int32) int32)(m, v184, v116, l3)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	if v186 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v190 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	base.MemoryCopy(m, v168, v114, v190)
	goto L44
L48:
	;
	if v196 == int32(0) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v200 == int32(0) {
		goto L35
	} else {
		goto L50
	}
L50:
	;
	base.MemoryCopy(m, v184, v116, v200)
	goto L35
L51:
	;
	goto L14
}
func F_gbt_num_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v14 = v12 << (uint(int32(1)) % 32)
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	base.MemoryCopy(m, l0, v15, v14)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if int32(2) <= v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = l0 + v12
	v24 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return l0
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v24*int32(24))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v38 = v36 + v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int32)(m, l0, v36, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v21, v38, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	if v40 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	base.MemoryCopy(m, l0, v36, v46)
	goto L9
L14:
	;
	v62 = v24 + int32(1)
	if v62 != v11 {
		v24 = v62
		goto L7
	} else {
		goto L18
	}
L15:
	;
	if v52 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v56 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.MemoryCopy(m, v21, v38, v56)
	goto L14
L18:
	;
	goto L8
}
func F_gbt_numeric_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_numeric_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_numeric_consistent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14275(m, l0, int32(_a_F_gbt_numeric_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_numeric_lt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_DirectFunctionCall2Coll(m, int32(1403), int32(0), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v9 != int64(0))
	}
}
func F_gbt_numeric_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14270(m, l0, l1, l2, int32(1467))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_oid_distance(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14279(m, l0, int32(_a_F_gbt_oid_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_oid_picksplit(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14272(m, l0, int32(_a_F_gbt_oid_picksplit_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_text_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[0]))
	if v5 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[1]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11*int32(28))+uint32(_c_F_gbt_text_compress[2])))
		*(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[0])) = v16
	} else {
	}
	v20 = F_gbt_var_compress(m, base.I32_wrap_i64(v3), int32(_a_F_gbt_text_compress_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v20)
	}
}
func F_gbt_text_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14284(m, l0, l1, l2, int32(2325))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_time_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_time_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_time_fetch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_time_fetch_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_gbt_ts_ssup_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
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
	v10 = F_DirectFunctionCall2Coll(m, int32(1578), int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
	}
}
func F_gbt_tsgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(1710), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_tsle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2651), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int64(0))
	}
}
func F_gbt_uuid_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+18)))
	if v6 != int32(1) {
		return base.I64_extend_i32_u(v5)
	} else {
		v12 = F_palloc(m, int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v18 = F_palloc(m, int32(24))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v20
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v24
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = base.I64_extend_i32_u(v12)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v32
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)))
				v35 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+18)) = uint8(v35)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v34)
				return base.I64_extend_i32_u(v18)
			}
		}
	}
}
func F_gbt_uuid_same(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14273(m, l0, int32(_a_F_gbt_uuid_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_gbt_uuidgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	return base.B2i32(int32(0) < v161)
}
func F_gbt_var_compress(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	if v6 != int32(1) {
		return l0
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v17 = int32(base.Ui32(v15) >> (uint(int32(2)) % 32))
			v19 = v17 + int32(4)
			v20 = F_palloc(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v17 != 0 {
					base.MemoryCopy(m, v20+int32(4), v11, v17)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v19 << (uint(int32(2)) % 32)
				v29 = F_palloc(m, int32(24))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = base.I64_extend_i32_u(v20)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v35
					v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v29)+18)) = uint8(v38)
					*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)) = uint16(v37)
					return v29
				}
			}
		}
	}
}
func F_gbt_vsrt_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v17 = int32(4)
	v18 = v14 + v17
	v20 = v11 + v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v25 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32) int32)(m, v18, v20, v21, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if v25 != 0 {
			v67 = v25
			return v67
		} else {
			v29 = int32(2)
			v30 = int32(base.Ui32(v15) >> (uint(v29) % 32))
			v36 = int32(4)
			if base.Ui32(v30+v36) < base.Ui32(int32(base.Ui32(v16)>>(uint(v29)%32))) {
				v43 = v14 + (v30+int32(3))&int32(2147483644) + v36
			} else {
				v43 = v18
			}
			v44 = int32(2)
			v45 = int32(base.Ui32(v12) >> (uint(v44) % 32))
			v51 = int32(4)
			if base.Ui32(v45+v51) < base.Ui32(int32(base.Ui32(v13)>>(uint(v44)%32))) {
				v58 = v11 + (v45+int32(3))&int32(2147483644) + v51
			} else {
				v58 = v20
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
			v63 = m.T0[v62].(func(*base.Module, int32, int32, int32, int32) int32)(m, v43, v58, v59, v60)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v67 = v63
				return v67
			}
		}
	}
}
