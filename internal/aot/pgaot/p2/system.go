package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SystemTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_makeString(m, int32(327481))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v8
		v13 = F_makeString(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v17
			v21 = F_list_make2_impl(m, v5+int32(4), v5)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_makeTypeNameFromNameList(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v23
				}
			}
		}
	}
}
func F_system_rows_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v7 < int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg(m, int32(343719), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(493837), int32(185), int32(284723))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v31 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v31
		*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)) = uint16(v31)
		*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v30))) = l3
		v37 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v37)
		return
	}
}
func F_system_rows_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int64
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 float64
	_ = v26
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v46 float64
	_ = v46
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v83 float64
	_ = v83
	var v87 float64
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v10 = int64(1000)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = F_estimate_expression_value(m, l0, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v15 != int32(7) {
			v25 = v10
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
			if v18 != 0 {
				v25 = v10
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				if v21 < int64(0) {
					v24 = int64(1000)
				} else {
					v24 = v21
				}
				v25 = v24
			}
		}
		v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		if base.F64_lt(base.F64_abs(v26), float64(9.223372036854776e+18)) != 0 {
			v32 = base.I64_trunc_f64_s(v26)
			v34 = v32
		} else {
			v34 = int64(-9223372036854775807 - 1)
		}
		if base.F64_lt(v26, base.F64_convert_i64_s(v25)) != 0 {
			v35 = v34
		} else {
			v35 = v25
		}
		v36 = base.F64_convert_i64_s(v35)
		v38 = float64(1e+100)
		if base.F64_gt(v36, v38) != 0 {
			v50 = v38
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) {
				v50 = v38
			} else {
				v46 = float64(1)
				if base.F64_le(v36, v46) != 0 {
					v50 = v46
				} else {
					v50 = base.F64_nearest(v36)
				}
			}
		}
		v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		if base.F64_lt(base.F64_abs(v50), float64(9.223372036854776e+18)) != 0 {
			v57 = base.I64_trunc_f64_s(v50)
			v59 = v57
		} else {
			v59 = int64(-9223372036854775807 - 1)
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.F64_gt(v51, float64(0)) != 0 {
			if v60 != 0 {
				v65 = base.F64_convert_i64_s(v59)
				v66 = base.F64_convert_i32_u(v60)
				v69 = v66
				v70 = v65
				v71 = base.F64_div(v65, base.F64_div(v51, v66))
			} else {
				v63 = float64(0)
				v64 = base.F64_convert_i64_s(v59)
				v69 = v63
				v70 = v64
				v71 = v64
			}
		} else {
			v63 = base.F64_convert_i32_u(v60)
			v64 = base.F64_convert_i64_s(v59)
			v69 = v63
			v70 = v64
			v71 = v64
		}
		if base.F64_lt(v69, v71) != 0 {
			v73 = v69
		} else {
			v73 = v71
		}
		v75 = float64(1e+100)
		if base.F64_gt(v73, v75) != 0 {
			v87 = v75
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v73)&int64(9223372036854775807)) {
				v87 = v75
			} else {
				v83 = float64(1)
				if base.F64_le(v73, v83) != 0 {
					v87 = v83
				} else {
					v87 = base.F64_nearest(v73)
				}
			}
		}
		if base.F64_lt(v87, float64(4.294967296e+09))&base.F64_ge(v87, float64(0)) != 0 {
			v93 = base.I32_trunc_f64_u(v87)
			v95 = v93
		} else {
			v95 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v95
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v70
		return
	}
}
