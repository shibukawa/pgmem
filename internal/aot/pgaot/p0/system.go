package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSystemIdentifier(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetSystemIdentifier[0]))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_SystemAttributeDefinition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if base.Ui32(l0) <= base.Ui32(int32(-7)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg_internal(m, int32(_a_F_SystemAttributeDefinition_0), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_SystemAttributeDefinition_1), int32(239), int32(_a_F_SystemAttributeDefinition_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32((l0^int32(-1))<<(uint(int32(2))%32))+uint32(_c_F_SystemAttributeDefinition[0])))
		m.G0 = v5 + int32(16)
		return v28
	}
}
func F_system_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v15 float64
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0))|base.F32_gt(v7, float32(100)) == int32(0) {
		v15 = base.F64_promote_f32(v7)
		if base.Ui64(base.I64_reinterpret_f64(v15)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v39 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)) = uint16(v39)
			*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = l3
			*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_trunc_sat_f64_u(base.F64_nearest(base.F64_div(base.F64_mul(v15, float64(4.294967296e+09)), float64(100))))
			v51 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v51)
			v54 = base.F32_ge(v7, float32(1))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v54)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_system_beginsamplescan_0), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_system_beginsamplescan_1), int32(151), int32(_a_F_system_beginsamplescan_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_system_beginsamplescan_0), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_system_beginsamplescan_1), int32(151), int32(_a_F_system_beginsamplescan_2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_system_rows_nextsampleblock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 float64
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v77 int32
	_ = v77
	var __phi77 int32
	_ = __phi77
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v17 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v155
L2:
	;
	v155 = int32(-1)
	goto L1
L3:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	if v21 == v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v114 = v17
	goto L5
L5:
	;
	v123 = v114 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v123
	v125 = int32(-1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if base.Ui32(v126) < base.Ui32(v123) {
		v155 = v125
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
	v101 = v20
	goto L8
L8:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v109
	v114 = v101
	goto L5
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pg_prng_seed(m, v14, base.I64_extend_i32_u(v26))
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l1
	goto L12
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v37 = base.F64_convert_i32_u(v36)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = base.I32_trunc_sat_f64_u(base.F64_mul(v33, v37))
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v33 = F_pg_prng_double(m, v14)
	mBase = m.M
	if base.F64_eq(v33, float64(0)) != 0 {
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
	v92 = int32(1)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v101 = v97
	goto L8
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_system_rows_nextsampleblock[0]))
	if v56 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v92 = v68
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	v68 = base.I32_trunc_sat_f64_u(base.F64_mul(v64, v37))
	if v68 == int32(0) {
		goto L18
	} else {
		goto L29
	}
L26:
	;
	v64 = F_pg_prng_double(m, v14)
	mBase = m.M
	if base.F64_eq(v64, float64(0)) != 0 {
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
	__phi74 = v36
	__phi77 = v68
	v74 = __phi74
	v77 = __phi77
	goto L30
L30:
	;
	v82 = base.I32_rem_u_s(v74, v77)
	if v82 != 0 {
		__phi74 = v77
		__phi77 = v82
		v74 = __phi74
		v77 = __phi77
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(int32(1)) < base.Ui32(v77) {
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
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	if v129 <= v128 {
		v155 = v125
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+36)))
	v137 = v131
	goto L36
L36:
	;
	v147 = base.I64_rem_u_s(v133+base.I64_extend_i32_u(v137), base.I64_extend_i32_u(v126))
	v148 = base.I32_wrap_i64(v147)
	if base.Ui32(l1) <= base.Ui32(v148) {
		v137 = v148
		goto L36
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v148
	v155 = v148
	goto L1
L38:
	;
	goto L37
}
func F_system_time_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v10 = int32(0)
	if base.B2i32(base.F64_lt(v7, float64(0)) == v10)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) == v10 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_system_time_beginsamplescan_0), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_system_time_beginsamplescan_1), int32(201), int32(_a_F_system_time_beginsamplescan_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
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
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v37 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v37
		*(*uint16)(unsafe.Add(mBase, uint32(v36)+24)) = uint16(v37)
		*(*float64)(unsafe.Add(mBase, uint32(v36)+8)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v36))) = l3
		return
	}
}
func F_system_time_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 float64
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v61 float64
	_ = v61
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v90 float64
	_ = v90
	var v94 float64
	_ = v94
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = float64(1000)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_estimate_expression_value(m, l0, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v19 != int32(7) {
			v36 = v14
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
			if v22 != 0 {
				v36 = v14
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
				if base.B2i32(base.F64_lt(v24, float64(0)) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					v36 = v24
				} else {
					v36 = float64(1000)
				}
			}
		}
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
		F_get_tablespace_page_costs(m, v37, v12+int32(8), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v44 = base.F64_convert_i32_u(v43)
			v45 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			if base.F64_gt(v45, float64(0)) != 0 {
				v49 = base.F64_div(v36, v45)
			} else {
				v49 = v36
			}
			if base.F64_gt(v49, v44) != 0 {
				v51 = v44
			} else {
				v51 = v49
			}
			v52 = float64(1e+100)
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v51)&int64(9223372036854775807)))|base.F64_gt(v51, v52) != 0 {
				v65 = v52
			} else {
				v61 = float64(1)
				if base.F64_le(v51, v61) != 0 {
					v65 = v61
				} else {
					v65 = base.F64_nearest(v51)
				}
			}
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
			if base.F64_gt(v66, float64(0)) == int32(0) {
				v78 = v65
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
				if v71 == int32(0) {
					v78 = v65
				} else {
					v78 = base.F64_mul(v65, base.F64_div(v66, base.F64_convert_i32_u(v71)))
				}
			}
			if base.F64_gt(v78, v66) != 0 {
				v80 = v66
			} else {
				v80 = v78
			}
			v81 = float64(1e+100)
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v80)&int64(9223372036854775807)))|base.F64_gt(v80, v81) != 0 {
				v94 = v81
			} else {
				v90 = float64(1)
				if base.F64_le(v80, v90) != 0 {
					v94 = v90
				} else {
					v94 = base.F64_nearest(v80)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = base.I32_trunc_sat_f64_u(v65)
			*(*float64)(unsafe.Add(mBase, uint32(l4))) = v94
			m.G0 = v12 + int32(16)
			return
		}
	}
}
