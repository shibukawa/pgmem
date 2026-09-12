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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
		if v46 < v40 {
			v54 = F__emscripten_memset_bulkmem(m, v41+v46, base.I32_extend8_s(int32(0)), v40-v36+int32(4))
			mBase = m.M
		} else {
		}
		v55 = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(v55) % 32)
		v60 = int32(8)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
		v66 = int32(base.Ui32(v62)>>(uint(v55)%32)) - v60
		if v66 != 0 {
			v67 = F__emscripten_memcpy_bulkmem(m, v41+int32(4), v33+v60, v66)
			mBase = m.M
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
		v73 = F_gbt_var_key_copy(m, v9+int32(8))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v41)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v73
			}
		}
	}
}
func F_gbt_bit_same(m *base.Module, l0 int32) int32 {
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
	v9 = F_gbt_var_same(m, v4, v5, v6, int32(4378800), v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v9)
		return v3
	}
}
func F_gbt_bit_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
			v63 = F_DirectFunctionCall2Coll(m, int32(2863), int32(0), v61, v62)
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4378840), v5)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7021)
	return v3
}
func F_gbt_bpcharcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2423), l2, l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2409), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_bytea_consistent(m *base.Module, l0 int32) int32 {
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
		v55 = F_gbt_var_consistent(m, v9+int32(8), v13, v17, v46, v50&int32(1), int32(4378880), v54)
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
func F_gbt_bytea_same(m *base.Module, l0 int32) int32 {
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
	v9 = F_gbt_var_same(m, v4, v5, v6, int32(4378880), v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v9)
		return v3
	}
}
func F_gbt_byteagt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_DirectFunctionCall2Coll(m, int32(2860), int32(0), l0, l1)
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
	v4 = F_gbt_num_fetch(m, v2, int32(4378920))
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4378960), v5)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6921)
	return v3
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2432), int32(0), v6, v7)
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2429), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v8 != int32(0))
	}
}
func F_gbt_enum_consistent(m *base.Module, l0 int32) int32 {
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
	v34 = F_gbt_num_consistent(m, v7, v7+int32(12), v7+int32(10), v29&int32(1), int32(4379000), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v34
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
	v4 = F_gbt_num_fetch(m, v2, int32(4379000))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_enum_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4379000), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
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
	v9 = F_CallerFInfoFunctionCall2(m, int32(3811), v5, int32(0), v7, v8)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(6936)
	return v3
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
	v6 = F_gbt_num_picksplit(m, v2, v3, int32(4379080), v5)
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4379184), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
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
	v34 = F_gbt_num_consistent(m, v7, v7+int32(12), v7+int32(10), v29&int32(1), int32(4379264), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v34
	}
}
func F_gbt_int4_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4379264), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_intv_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4379344), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_intv_ssup_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2540), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_gbt_macad8_union(m *base.Module, l0 int32) int32 {
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
		v15 = F_gbt_num_union(m, v6, v4, int32(4379424), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_gbt_macad8eq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(4188), int32(0), l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_macad8key_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(4194), int32(0), v6, v7)
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
			v18 = F_DirectFunctionCall2Coll(m, int32(4194), int32(0), v6+v14, v7+v14)
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
func F_gbt_macadeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2276), int32(0), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2277), int32(0), l0, l1)
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
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
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
			F_errmsg_internal(m, int32(475273), v9)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(496466), int32(326), int32(416790))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
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
	v7 = F_DirectFunctionCall2Coll(m, int32(1344), int32(0), l0, l1)
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
	v7 = F_DirectFunctionCall2Coll(m, int32(1295), int32(0), l0, l1)
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(7012)
	return v3
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
	v9 = F_gbt_var_same(m, v4, v5, v6, int32(4379544), v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v9)
		return v3
	}
}
func F_gbt_textgt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2246), l2, l0, l1)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
	}
}
func F_gbt_time_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4379624), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
		return v3
	}
}
func F_gbt_timele(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall2Coll(m, int32(2435), int32(0), l0, l1)
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
	v44 = F_gbt_num_consistent(m, v10+v22, v10+int32(16), v10+int32(30), v39&v18, int32(4379624), v43)
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
	v37 = F_gbt_num_consistent(m, v7+int32(12), v7+int32(24), v7+int32(22), v32&int32(1), int32(4379664), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v37
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
	v4 = F_gbt_num_fetch(m, v2, int32(4379664))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gbt_ts_same(m *base.Module, l0 int32) int32 {
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
	v8 = F_gbt_num_same(m, v4, v5, int32(4379664), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(v8)
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
	v34 = F_gbt_num_consistent(m, v8+int32(4), v10, v8+int32(14), v29&int32(1), int32(4379704), v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v34
	}
}
func F_gbt_uuidle(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	return base.B2i32(v66 <= int32(0))
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
