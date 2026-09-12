package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gbt_bit_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_var_compress(m, v2, int32(4396080))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_bit_consistent(m *base.Module, l0 int32) int32 {
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+16)))
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47)+12)))
		if v49&int32(1) != 0 {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v60 = F_gbt_var_consistent(m, v11+int32(8), v15, v19&int32(65535), v56, int32(1), int32(4396080), v59)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v109 = v60
				m.G0 = v11 + int32(16)
				return v109
			}
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v64 = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
			v68 = (v64 - int32(1)) & int32(-4)
			v69 = F_palloc(m, v68)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v72 = v64 - int32(4)
				if v72 < v68 {
					v80 = F__emscripten_memset_bulkmem(m, v69+v72, base.I32_extend8_s(int32(0)), v68-v64+int32(4))
					mBase = m.M
				} else {
				}
				v81 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v69))) = v68 << (uint(v81) % 32)
				v86 = int32(8)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v92 = int32(base.Ui32(v88)>>(uint(v81)%32)) - v86
				if v92 != 0 {
					v93 = F__emscripten_memcpy_bulkmem(m, v69+int32(4), v15+v86, v92)
					mBase = m.M
				} else {
				}
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v103 = F_gbt_var_consistent(m, v11+int32(8), v69, v19&int32(65535), v99, int32(0), int32(4396080), v102)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					v109 = v103
					m.G0 = v11 + int32(16)
					return v109
				}
			}
		}
	}
}
func F_gbt_bitgt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2658), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2413), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_bytea_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_var_compress(m, v2, int32(4396160))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_byteale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2859), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_cash_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396200))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
func F_gbt_enum_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396280))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_enumle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3809), l2, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_float4_penalty(m *base.Module, l0 int32) int32 {
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
	var v15 float32
	_ = v15
	var v16 float32
	_ = v16
	var v27 float64
	_ = v27
	var v28 float32
	_ = v28
	var v29 float32
	_ = v29
	var v39 float64
	_ = v39
	var v44 float32
	_ = v44
	var v48 float32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	v15 = *(*float32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.F32_gt(v15, v16) != 0 {
		v27 = base.F64_add(base.F64_add(base.F64_mul(base.F64_promote_f32(v15), float64(0.49000000953674316)), base.F64_mul(base.F64_promote_f32(v16), float64(-0.49000000953674316))), float64(0))
	} else {
		v27 = float64(0)
	}
	v28 = *(*float32)(unsafe.Add(mBase, uint32(v11)))
	v29 = *(*float32)(unsafe.Add(mBase, uint32(v9)))
	if base.F32_gt(v28, v29) != 0 {
		v39 = base.F64_add(v27, base.F64_add(base.F64_mul(base.F64_promote_f32(v28), float64(0.49000000953674316)), base.F64_mul(base.F64_promote_f32(v29), float64(-0.49000000953674316))))
	} else {
		v39 = v27
	}
	if base.F64_gt(v39, float64(0)) != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(8388608)
		v44 = *(*float32)(unsafe.Add(mBase, uint32(v11)+4))
		v48 = *(*float32)(unsafe.Add(mBase, uint32(v11)))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v39, base.F64_add(v39, base.F64_add(base.F64_mul(base.F64_promote_f32(v44), float64(0.49000000953674316)), base.F64_mul(base.F64_promote_f32(v48), float64(-0.49000000953674316)))))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v61+int32(1))))
	} else {
	}
	return v12
}
func F_gbt_float4_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4396320), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_float4key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	var v12 float32
	_ = v12
	var v13 float32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(v8)))
	if base.F32_eq(v7, v9) != 0 {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
		v13 = *(*float32)(unsafe.Add(mBase, uint32(v8)+4))
		if base.F32_eq(v12, v13) != 0 {
			v26 = int32(0)
			return v26
		} else {
			if base.F32_gt(v12, v13) != 0 {
				v18 = int32(1)
			} else {
				v18 = int32(-1)
			}
			return v18
		}
	} else {
		if base.F32_gt(v7, v9) != 0 {
			v23 = int32(1)
		} else {
			v23 = int32(-1)
		}
		v26 = v23
		return v26
	}
}
func F_gbt_float8_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6981)
	return v3
}
func F_gbt_float8key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	if base.F64_eq(v7, v9) != 0 {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		if base.F64_eq(v12, v13) != 0 {
			v26 = int32(0)
			return v26
		} else {
			if base.F64_gt(v12, v13) != 0 {
				v18 = int32(1)
			} else {
				v18 = int32(-1)
			}
			return v18
		}
	} else {
		if base.F64_gt(v7, v9) != 0 {
			v23 = int32(1)
		} else {
			v23 = int32(-1)
		}
		v26 = v23
		return v26
	}
}
func F_gbt_inet_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4396464), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_inet_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6989)
	return v3
}
func F_gbt_int2_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396504))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int2_consistent(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v14 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v29)+12)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_gbt_num_consistent(m, v7+int32(4), v7+int32(14), v7+int32(12), v31&int32(1), int32(4396504), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v36
	}
}
func F_gbt_int2_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4396504))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int4_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396544))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
func F_gbt_int4_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4396544))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int4_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v5 < v4) - base.B2i32(v4 < v5)
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
func F_gbt_int8_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396584))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_int8_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4396584))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_intv_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4396624))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_macad8_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7035)
	return v3
}
func F_gbt_macad8ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(4192), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
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
	var v48 int64
	_ = v48
	var v49 int64
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
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v90 int64
	_ = v90
	var v105 int64
	_ = v105
	var v120 int64
	_ = v120
	var v131 float64
	_ = v131
	var v141 float64
	_ = v141
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
	v33 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	v34 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	v35 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
	v36 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v37 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)))
	v41 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	v42 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+11)))
	v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+10)))
	v48 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+9)))
	v49 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+11)))
	v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+10)))
	v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
	v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+8)))
	v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+7)))
	v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(0)
	v61 = int64(32)
	v63 = int64(40)
	v66 = int64(24)
	v69 = int64(16)
	v72 = int64(8)
	v75 = v40 | (v44<<(uint(v61)%64) | v45<<(uint(v63)%64) | v43<<(uint(v66)%64) | v42<<(uint(v69)%64) | v41<<(uint(v72)%64))
	v90 = v32 | (v36<<(uint(v61)%64) | v37<<(uint(v63)%64) | v35<<(uint(v66)%64) | v34<<(uint(v69)%64) | v33<<(uint(v72)%64))
	v105 = v46 | (v50<<(uint(v61)%64) | v51<<(uint(v63)%64) | v49<<(uint(v66)%64) | v48<<(uint(v69)%64) | v47<<(uint(v72)%64))
	v120 = v52 | (v56<<(uint(v61)%64) | v57<<(uint(v63)%64) | v55<<(uint(v66)%64) | v54<<(uint(v69)%64) | v53<<(uint(v72)%64))
	if base.Ui64(v120) < base.Ui64(v105) {
		v131 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v105), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v120), float64(-0.49000000953674316))), float64(0))
	} else {
		v131 = float64(0)
	}
	if base.Ui64(v75) < base.Ui64(v90) {
		v141 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v90), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v75), float64(-0.49000000953674316))), v131)
	} else {
		v141 = v131
	}
	if base.F64_gt(v141, float64(0)) != 0 {
		v157 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
		v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
		*(*float32)(unsafe.Add(mBase, uint32(v58))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v141, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v120), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v90), float64(-0.49000000953674316))), v141))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v159+int32(1))))
	} else {
	}
	return v58
}
func F_gbt_macaddr_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7028)
	return v3
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
	var v26 int32
	_ = v26
	var v27 int32
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v17 != 0 {
				v18 = F__emscripten_memcpy_bulkmem(m, v14, l1, v17)
				mBase = m.M
			} else {
			}
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v20 != 0 {
				v23 = F__emscripten_memcpy_bulkmem(m, v20+v21, v9, v20)
				mBase = m.M
			} else {
			}
			return
		}
	} else {
		v25 = v10 + v8
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v27 = m.T0[v26].(func(*base.Module, int32, int32, int32) int32)(m, v10, l1, l3)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			if v27 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v29 != 0 {
					v30 = F__emscripten_memcpy_bulkmem(m, v10, l1, v29)
					mBase = m.M
				} else {
				}
			} else {
			}
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
			v33 = m.T0[v32].(func(*base.Module, int32, int32, int32) int32)(m, v25, v9, l3)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				if v33 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v35 != 0 {
						v36 = F__emscripten_memcpy_bulkmem(m, v25, v9, v35)
						mBase = m.M
					} else {
					}
				} else {
				}
				return
			}
		}
	}
}
func F_gbt_num_picksplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = (v13 - int32(1)) & int32(65535)
	v22 = F_palloc(m, v17<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = v17<<(uint(int32(1))%32) + int32(4)
	v30 = F_palloc(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
	v33 = F_palloc(m, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v33
	v36 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+20)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = v36
	if v13&int32(65535) == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = int32(8)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v22+v44, v17, v44, v47, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v54 = int32(1)
	goto L9
L8:
	;
	return l1
L9:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v54<<(uint(int32(4))%32))))
	v72 = v22 + v54<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v69
	v78 = (v54 + int32(1)) & int32(65535)
	if base.Ui32(v78) <= base.Ui32(v17) {
		v54 = v78
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v80 = int32(8)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v22+v80, v17, v80, v83, l3)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v86 = int32(1)
	v89 = v86
	goto L13
L13:
	;
	v103 = v22 + v89<<(uint(int32(3))%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v106 = v104 + v105
	if base.Ui32(v89) <= base.Ui32(int32(base.Ui32(v17)>>(uint(v86)%32))) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return l1
L15:
	;
	v198 = (v89 + int32(1)) & int32(65535)
	if base.Ui32(v198) <= base.Ui32(v17) {
		v89 = v198
		goto L13
	} else {
		goto L73
	}
L16:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v108 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v150 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v142 = int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*uint16)(unsafe.Add(mBase, uint32(v140+v141<<(uint(v142)%32)))) = uint16(v145)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v141 + v142
	goto L15
L20:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v112 = F_palloc0(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v123 = v105 + v108
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v125 = m.T0[v124].(func(*base.Module, int32, int32, int32) int32)(m, v108, v104, l3)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L32
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v115 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v118 != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v116 = F__emscripten_memcpy_bulkmem(m, v112, v104, v115)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L19
L29:
	;
	v121 = F__emscripten_memcpy_bulkmem(m, v118+v119, v106, v118)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	if v125 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v127 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v131 = m.T0[v130].(func(*base.Module, int32, int32, int32) int32)(m, v123, v106, l3)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L40
	}
L36:
	;
	goto L35
L37:
	;
	v128 = F__emscripten_memcpy_bulkmem(m, v108, v104, v127)
	mBase = m.M
	goto L39
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	if v131 == int32(0) {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v135 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L19
L43:
	;
	v136 = F__emscripten_memcpy_bulkmem(m, v123, v106, v135)
	mBase = m.M
	goto L45
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v184 = int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*uint16)(unsafe.Add(mBase, uint32(v182+v183<<(uint(v184)%32)))) = uint16(v187)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v183 + v184
	goto L15
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v154 = F_palloc0(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v165 = v105 + v150
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32) int32)(m, v150, v104, l3)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L59
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v157 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v160 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v158 = F__emscripten_memcpy_bulkmem(m, v154, v104, v157)
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L46
L56:
	;
	v163 = F__emscripten_memcpy_bulkmem(m, v160+v161, v106, v160)
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	if v167 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v169 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v173 = m.T0[v172].(func(*base.Module, int32, int32, int32) int32)(m, v165, v106, l3)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L67
	}
L63:
	;
	goto L62
L64:
	;
	v170 = F__emscripten_memcpy_bulkmem(m, v150, v104, v169)
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	if v173 == int32(0) {
		goto L46
	} else {
		goto L68
	}
L68:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v177 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L46
L70:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v165, v106, v177)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	goto L14
}
func F_gbt_num_union(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v14 = v12 << (uint(int32(1)) % 32)
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(2) <= v10 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v15 = F__emscripten_memcpy_bulkmem(m, l0, v11, v14)
	mBase = m.M
	v16 = v15
	goto L4
L3:
	;
	v16 = l0
	goto L4
L4:
	;
	goto L1
L5:
	;
	v21 = v12 + v16
	v24 = int32(1)
	goto L8
L6:
	;
	goto L7
L7:
	;
	return v16
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)+v24<<(uint(int32(4))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v37 = v35 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, v16, v35, l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	return int32(0)
L11:
	;
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v47 = m.T0[v46].(func(*base.Module, int32, int32, int32) int32)(m, v21, v37, l3)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	goto L14
L16:
	;
	v44 = F__emscripten_memcpy_bulkmem(m, v16, v35, v43)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v49 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v53 = v24 + int32(1)
	if v53 != v10 {
		v24 = v53
		goto L8
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v50 = F__emscripten_memcpy_bulkmem(m, v21, v37, v49)
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L9
}
func F_gbt_numeric_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_var_compress(m, v2, int32(4396744))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_numeric_consistent(m *base.Module, l0 int32) int32 {
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
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
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v48)+12)))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v55 = F_gbt_var_consistent(m, v9+int32(8), v13, v17, v46, v50&int32(1), int32(4396744), v54)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			m.G0 = v9 + int32(16)
			return v55
		}
	}
}
func F_gbt_numeric_lt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(1291), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_numeric_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = v9 + int32(8)
			v20 = int32(4)
			v21 = v11 + v20
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v24 = int32(2)
			v25 = int32(base.Ui32(v23) >> (uint(v24) % 32))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui32(v25+v20) < base.Ui32(int32(base.Ui32(v33)>>(uint(v24)%32))) {
				v37 = v21 + (v25+int32(3))&int32(2147483644)
			} else {
				v37 = v21
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v37
			v40 = int32(4)
			v41 = v15 + v40
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v44 = int32(2)
			v45 = int32(base.Ui32(v43) >> (uint(v44) % 32))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if base.Ui32(v45+v40) < base.Ui32(int32(base.Ui32(v53)>>(uint(v44)%32))) {
				v57 = v41 + (v45+int32(3))&int32(2147483644)
			} else {
				v57 = v41
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v57
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v63 = F_DirectFunctionCall2Coll(m, int32(1343), int32(0), v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v63
							}
						} else {
							m.G0 = v9 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v63
						}
					} else {
						m.G0 = v9 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func F_gbt_oid_distance(m *base.Module, l0 int32) int32 {
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
	v29 = F_gbt_num_distance(m, v7+v13, v7+int32(12), v24&int32(1), int32(4396784), v28)
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
func F_gbt_oid_picksplit(m *base.Module, l0 int32) int32 {
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4396784), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_text_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1568]))
	if v4 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[495]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_consts[1304])))
		*(*int32)(unsafe.Add(mBase, _consts[1568])) = v15
	} else {
	}
	v18 = F_gbt_var_compress(m, v2, int32(4396824))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		return v18
	}
}
func F_gbt_text_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = v9 + int32(8)
			v20 = int32(4)
			v21 = v11 + v20
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v24 = int32(2)
			v25 = int32(base.Ui32(v23) >> (uint(v24) % 32))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui32(v25+v20) < base.Ui32(int32(base.Ui32(v33)>>(uint(v24)%32))) {
				v37 = v21 + (v25+int32(3))&int32(2147483644)
			} else {
				v37 = v21
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v37
			v40 = int32(4)
			v41 = v15 + v40
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v44 = int32(2)
			v45 = int32(base.Ui32(v43) >> (uint(v44) % 32))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if base.Ui32(v45+v40) < base.Ui32(int32(base.Ui32(v53)>>(uint(v44)%32))) {
				v57 = v41 + (v45+int32(3))&int32(2147483644)
			} else {
				v57 = v41
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v57
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v63 = F_DirectFunctionCall2Coll(m, int32(2120), v60, v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v63
							}
						} else {
							m.G0 = v9 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v63
						}
					} else {
						m.G0 = v9 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func F_gbt_time_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_compress(m, v2, int32(4396904))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_time_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_gbt_num_fetch(m, v2, int32(4396904))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_ts_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1446), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_tsgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(1513), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_tsle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2454), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_uuid_compress(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v18 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = v16 + int32(8)
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v24
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v26
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v12
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v33
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v36)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v35)
				return v18
			}
		}
	}
}
func F_gbt_uuid_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4396984), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_uuidgt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	return base.B2i32(int32(0) < v66)
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
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
					v24 = F__emscripten_memcpy_bulkmem(m, v20+int32(4), v11, v17)
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v19 << (uint(int32(2)) % 32)
				v30 = F_palloc(m, int32(16))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v20
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v35
					v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v38 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v30)+14)) = uint8(v38)
					*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)) = uint16(v37)
					return v30
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
