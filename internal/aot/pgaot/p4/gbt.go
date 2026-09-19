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
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_bit_compress_0))
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
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
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v58 = F_gbt_var_consistent(m, v22, v15, v19&int32(_a_F_gbt_bit_consistent_0), v54, int32(1), int32(_a_F_gbt_bit_consistent_1), v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v108 = v58
				m.G0 = v11 + int32(16)
				return v108
			}
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v62 = int32(base.Ui32(v60) >> (uint(int32(2)) % 32))
			v66 = (v62 - int32(1)) & int32(-4)
			v67 = F_palloc(m, v66)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				v70 = v62 - int32(4)
				if v66 <= v70 {
				} else {
					v74 = v66 - v62 + int32(4)
					if v74 == int32(0) {
					} else {
						base.MemoryFill(m, v67+v70, int32(0), v74)
					}
				}
				v81 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66 << (uint(v81) % 32)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v88 = int32(base.Ui32(v84)>>(uint(v81)%32)) - int32(8)
				if v88 != 0 {
					base.MemoryCopy(m, v67+int32(4), v15+int32(8), v88)
				} else {
				}
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v102 = F_gbt_var_consistent(m, v11+int32(8), v67, v19&int32(_a_F_gbt_bit_consistent_0), v98, int32(0), int32(_a_F_gbt_bit_consistent_1), v101)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v108 = v102
					m.G0 = v11 + int32(16)
					return v108
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2642), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2397), l2, l0, l1)
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
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_bytea_compress_0))
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
	v7 = F_DirectFunctionCall2Coll(m, int32(2843), int32(0), l0, l1)
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_cash_compress_0))
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_enum_compress_0))
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
	v8 = F_CallerFInfoFunctionCall2(m, int32(3793), l2, int32(0), v6, v7)
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
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_gbt_float4_penalty_0)
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13886(m, l0, int32(_a_F_gbt_float4_union_0), int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_float8_sortsupport_0)
	return v4
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_inet_picksplit_0), v5)
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
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_inet_sortsupport_0)
	return v4
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int2_compress_0))
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
	v36 = F_gbt_num_consistent(m, v7+int32(4), v7+int32(14), v7+int32(12), v31&int32(1), int32(_a_F_gbt_int2_consistent_0), v35)
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
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int2_fetch_0))
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int4_compress_0))
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
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int4_fetch_0))
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_int8_compress_0))
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
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_int8_fetch_0))
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
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_intv_fetch_0))
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
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_macad8_sortsupport_0)
	return v4
}
func F_gbt_macad8ge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_gbt_macad8ge_0), int32(0), l0, l1)
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
	var v101 float64
	_ = v101
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v131 int64
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
	v75 = v46 | (v50<<(uint(v61)%64) | v51<<(uint(v63)%64) | v49<<(uint(v66)%64) | v48<<(uint(v69)%64) | v47<<(uint(v72)%64))
	v90 = v52 | (v56<<(uint(v61)%64) | v57<<(uint(v63)%64) | v55<<(uint(v66)%64) | v54<<(uint(v69)%64) | v53<<(uint(v72)%64))
	if base.Ui64(v90) < base.Ui64(v75) {
		v101 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v75), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v90), float64(-0.49000000953674316))), float64(0))
	} else {
		v101 = float64(0)
	}
	v102 = int64(32)
	v104 = int64(40)
	v107 = int64(24)
	v110 = int64(16)
	v113 = int64(8)
	v116 = v40 | (v44<<(uint(v102)%64) | v45<<(uint(v104)%64) | v43<<(uint(v107)%64) | v42<<(uint(v110)%64) | v41<<(uint(v113)%64))
	v131 = v32 | (v36<<(uint(v102)%64) | v37<<(uint(v104)%64) | v35<<(uint(v107)%64) | v34<<(uint(v110)%64) | v33<<(uint(v113)%64))
	if base.Ui64(v116) < base.Ui64(v131) {
		v141 = base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v131), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v116), float64(-0.49000000953674316))), v101)
	} else {
		v141 = v101
	}
	if base.F64_gt(v141, float64(0)) != 0 {
		v157 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
		v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
		*(*float32)(unsafe.Add(mBase, uint32(v58))) = base.F32_mul(base.F32_add(base.F32_demote_f64(base.F64_div(v141, base.F64_add(base.F64_add(base.F64_mul(base.F64_convert_i64_u(v90), float64(0.49000000953674316)), base.F64_mul(base.F64_convert_i64_u(v131), float64(-0.49000000953674316))), v141))), float32(1.1754944e-38)), base.F32_div(float32(3.4028235e+38), base.F32_convert_i32_s(v159+int32(1))))
	} else {
	}
	return v58
}
func F_gbt_macaddr_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(_a_F_gbt_macaddr_sortsupport_0)
	return v4
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
				base.MemoryCopy(m, v14, l1, v17)
			} else {
			}
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v19 == int32(0) {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, v22+v19, v9, v19)
				return
			}
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
			if v27 == int32(0) {
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v31 == int32(0) {
				} else {
					base.MemoryCopy(m, v10, l1, v31)
				}
			}
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
			v37 = m.T0[v36].(func(*base.Module, int32, int32, int32) int32)(m, v25, v9, l3)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				if v37 == int32(0) {
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v41 == int32(0) {
					} else {
						base.MemoryCopy(m, v25, v9, v41)
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = (v14 - int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	v23 = F_palloc(m, v18<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = v18<<(uint(int32(1))%32) + int32(4)
	v31 = F_palloc(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
	v34 = F_palloc(m, v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v34
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+20)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = v37
	if v14&int32(_a_F_gbt_num_picksplit_0) == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v45 = int32(8)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v23+v45, v18, v45, v48, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v55 = int32(1)
	goto L9
L8:
	;
	return l1
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(4)+v55<<(uint(int32(4))%32))))
	v74 = v23 + v55<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v71
	v80 = (v55 + int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	if base.Ui32(v80) <= base.Ui32(v18) {
		v55 = v80
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v82 = int32(8)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	F_qsort_arg(m, v23+v82, v18, v82, v85, l3)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v88 = int32(1)
	v91 = v88
	goto L13
L13:
	;
	v106 = v23 + v91<<(uint(int32(3))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v109 = v107 + v108
	if base.Ui32(v91) <= base.Ui32(int32(base.Ui32(v18)>>(uint(v88)%32))) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return l1
L15:
	;
	v216 = (v91 + int32(1)) & int32(_a_F_gbt_num_picksplit_0)
	if base.Ui32(v216) <= base.Ui32(v18) {
		v91 = v216
		goto L13
	} else {
		goto L51
	}
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v111 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v160 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v152 = int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*uint16)(unsafe.Add(mBase, uint32(v150+v151<<(uint(v152)%32)))) = uint16(v155)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v151 + v152
	goto L15
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v115 = F_palloc0(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v126 = v111 + v108
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v128 = m.T0[v127].(func(*base.Module, int32, int32, int32) int32)(m, v111, v107, l3)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v118 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	base.MemoryCopy(m, v115, v107, v118)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v120 == int32(0) {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	base.MemoryCopy(m, v123+v120, v109, v120)
	goto L19
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v138 = m.T0[v137].(func(*base.Module, int32, int32, int32) int32)(m, v126, v109, l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	if v128 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v132 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	base.MemoryCopy(m, v111, v107, v132)
	goto L28
L32:
	;
	if v138 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v142 == int32(0) {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.MemoryCopy(m, v126, v109, v142)
	goto L19
L35:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v201 = int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*uint16)(unsafe.Add(mBase, uint32(v199+v200<<(uint(v201)%32)))) = uint16(v204)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v200 + v201
	goto L15
L36:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v164 = F_palloc0(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v175 = v160 + v108
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v177 = m.T0[v176].(func(*base.Module, int32, int32, int32) int32)(m, v160, v107, l3)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v164
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v167 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryCopy(m, v164, v107, v167)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v169 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	base.MemoryCopy(m, v172+v169, v109, v169)
	goto L35
L44:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v187 = m.T0[v186].(func(*base.Module, int32, int32, int32) int32)(m, v175, v109, l3)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	if v177 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v181 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	base.MemoryCopy(m, v160, v107, v181)
	goto L44
L48:
	;
	if v187 == int32(0) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v191 == int32(0) {
		goto L35
	} else {
		goto L50
	}
L50:
	;
	base.MemoryCopy(m, v175, v109, v191)
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)+v24<<(uint(int32(4))%32))))
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
	v4 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_numeric_compress_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_numeric_consistent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13887(m, l0, int32(_a_F_gbt_numeric_consistent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_gbt_numeric_lt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(1275), int32(0), l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v7 != int32(0))
	}
}
func F_gbt_numeric_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13884(m, l0, l1, l2, int32(1327))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_gbt_oid_distance(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13891(m, l0, int32(_a_F_gbt_oid_distance_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(_a_F_gbt_oid_picksplit_0), v5)
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[0]))
	if v4 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_gbt_text_compress[2])))
		*(*int32)(unsafe.Add(mBase, _c_F_gbt_text_compress[0])) = v15
	} else {
	}
	v18 = F_gbt_var_compress(m, v2, int32(_a_F_gbt_text_compress_0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		return v18
	}
}
func F_gbt_text_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13896(m, l0, l1, l2, int32(2104))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	v4 = F_gbt_num_compress(m, v2, int32(_a_F_gbt_time_compress_0))
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
	v4 = F_gbt_num_fetch(m, v2, int32(_a_F_gbt_time_fetch_0))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1430), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1497), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2438), int32(0), l0, l1)
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)))
	if v6 != int32(1) {
		return v5
	} else {
		v11 = F_palloc(m, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v17 = F_palloc(m, int32(16))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v19
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				*(*int64)(unsafe.Add(mBase, uint32(v11))) = v21
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v23
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v11
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v30
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)))
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v33)
				*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v32)
				return v17
			}
		}
	}
}
func F_gbt_uuid_same(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13885(m, l0, int32(_a_F_gbt_uuid_same_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
					base.MemoryCopy(m, v20+int32(4), v11, v17)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v19 << (uint(int32(2)) % 32)
				v29 = F_palloc(m, int32(16))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v29))) = v20
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v32
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v34
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v37 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v29)+14)) = uint8(v37)
					*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v36)
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
