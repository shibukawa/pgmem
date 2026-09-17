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
	v8 = F_makeString(m, int32(_a_F_SystemTypeName_0))
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
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
				F_errmsg(m, int32(_a_F_system_rows_beginsamplescan_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_system_rows_beginsamplescan_1), int32(185), int32(_a_F_system_rows_beginsamplescan_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v27 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v27
		*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)) = uint16(v27)
		*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = l3
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v33)
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
	var v30 int64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v41 float64
	_ = v41
	var v45 float64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
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
		if base.F64_lt(v26, base.F64_convert_i64_s(v25)) != 0 {
			v30 = base.I64_trunc_sat_f64_s(v26)
		} else {
			v30 = v25
		}
		v31 = base.F64_convert_i64_s(v30)
		v32 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)&int64(9223372036854775807)))|base.F64_gt(v31, v32) != 0 {
			v45 = v32
		} else {
			v41 = float64(1)
			if base.F64_le(v31, v41) != 0 {
				v45 = v41
			} else {
				v45 = base.F64_nearest(v31)
			}
		}
		v46 = base.I64_trunc_sat_f64_s(v45)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		v48 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		if base.F64_gt(v48, float64(0)) == int32(0) {
			v60 = base.F64_convert_i32_u(v47)
			v61 = base.F64_convert_i64_s(v46)
			v62 = v61
			v63 = v60
			v64 = v61
		} else {
			if v47 == int32(0) {
				v60 = float64(0)
				v61 = base.F64_convert_i64_s(v46)
				v62 = v61
				v63 = v60
				v64 = v61
			} else {
				v56 = base.F64_convert_i64_s(v46)
				v57 = base.F64_convert_i32_u(v47)
				v62 = v56
				v63 = v57
				v64 = base.F64_div(v56, base.F64_div(v48, v57))
			}
		}
		if base.F64_gt(v64, v63) != 0 {
			v66 = v63
		} else {
			v66 = v64
		}
		v67 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)))|base.F64_gt(v66, v67) != 0 {
			v80 = v67
		} else {
			v76 = float64(1)
			if base.F64_le(v66, v76) != 0 {
				v80 = v76
			} else {
				v80 = base.F64_nearest(v66)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.I32_trunc_sat_f64_u(v80)
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v62
		return
	}
}
func F_system_time_nextsampleblock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 float64
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var __phi68 int32
	_ = __phi68
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 float64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v163
L2:
	;
	v163 = int32(-1)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v113 = v16
	goto L5
L5:
	;
	v124 = v113 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v124
	v126 = int32(-1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if base.Ui32(v127) < base.Ui32(v124) {
		v163 = v126
		goto L1
	} else {
		goto L34
	}
L6:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v102
	F___clock_gettime(m, int32(1), v13)
	mBase = m.M
	v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+8)))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v106 + v107*int64(1000000000)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v113 = v112
	goto L5
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	F_pg_prng_seed(m, v13, base.I64_extend_i32_u(v24))
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l1
	goto L12
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v35 = base.F64_convert_i32_u(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = base.I32_trunc_sat_f64_u(base.F64_mul(v31, v35))
	if base.Ui32(int32(2)) <= base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v31 = F_pg_prng_double(m, v13)
	mBase = m.M
	if base.F64_eq(v31, float64(0)) != 0 {
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	goto L18
L16:
	;
	v86 = int32(1)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v86
	goto L8
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_system_time_nextsampleblock[0]))
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v86 = v65
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L26
L23:
	;
	return int32(0)
L24:
	;
	goto L22
L25:
	;
	v65 = base.I32_trunc_sat_f64_u(base.F64_mul(v61, v35))
	if v65 == int32(0) {
		goto L18
	} else {
		goto L29
	}
L26:
	;
	v61 = F_pg_prng_double(m, v13)
	mBase = m.M
	if base.F64_eq(v61, float64(0)) != 0 {
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	goto L27
L29:
	;
	__phi68 = v34
	__phi74 = v65
	v68 = __phi68
	v74 = __phi74
	goto L30
L30:
	;
	v78 = base.I32_rem_u_s(v68, v74)
	if v78 != 0 {
		__phi68 = v74
		__phi74 = v78
		v68 = __phi68
		v74 = __phi74
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(int32(1)) < base.Ui32(v74) {
		goto L18
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L19
L34:
	;
	F___clock_gettime(m, int32(1), v13)
	mBase = m.M
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	v132 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+8)))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	if base.F64_le(v131, base.F64_div(base.F64_convert_i64_s(v132+v133*int64(1000000000)-v137), float64(1e+06))) != 0 {
		v163 = v126
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v144 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)))
	v145 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)))
	v146 = v143
	goto L36
L36:
	;
	v158 = base.I64_rem_u_s(v145+base.I64_extend_i32_u(v146), v144)
	v159 = base.I32_wrap_i64(v158)
	if base.Ui32(l1) <= base.Ui32(v159) {
		v146 = v159
		goto L36
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v159
	v163 = v159
	goto L1
L38:
	;
	goto L37
}
