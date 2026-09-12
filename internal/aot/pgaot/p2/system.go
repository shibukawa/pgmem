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
	v8 = F_makeString(m, int32(342679))
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
				F_errmsg(m, int32(359560), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(515318), int32(185), int32(297063))
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
func F_system_time_nextsampleblock(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v25 int32
	_ = v25
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var __phi84 int32
	_ = __phi84
	var v88 int32
	_ = v88
	var __phi88 int32
	_ = __phi88
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 float64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v157 int64
	_ = v157
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v17 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v184
L2:
	;
	v184 = int32(-1)
	goto L1
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v132 = v17
	goto L5
L5:
	;
	v144 = v132 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v144
	v146 = int32(-1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	if base.Ui32(v147) < base.Ui32(v144) {
		v184 = v146
		goto L1
	} else {
		goto L42
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v121
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v125 + v126*int64(1000000000)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v132 = v131
	goto L5
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pg_prng_seed(m, v14, base.I64_extend_i32_u(v25))
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l1
	goto L13
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v45
	if base.Ui32(int32(2)) <= base.Ui32(v35) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v36 = base.F64_convert_i32_u(v35)
	v37 = base.F64_mul(v32, v36)
	if base.F64_lt(v37, float64(4.294967296e+09))&base.F64_ge(v37, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v32 = F_pg_prng_double(m, v14)
	mBase = m.M
	if base.F64_eq(v32, float64(0)) != 0 {
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	goto L14
L16:
	;
	v43 = base.I32_trunc_f64_u(v37)
	v45 = v43
	goto L11
L17:
	;
	goto L18
L18:
	;
	v45 = int32(0)
	goto L11
L19:
	;
	goto L22
L20:
	;
	v103 = int32(1)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v103
	goto L8
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v62 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v103 = v81
	goto L21
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	goto L31
L27:
	;
	return int32(0)
L28:
	;
	goto L26
L29:
	;
	if v81 == int32(0) {
		goto L22
	} else {
		goto L37
	}
L30:
	;
	v73 = base.F64_mul(v70, v36)
	if base.F64_lt(v73, float64(4.294967296e+09))&base.F64_ge(v73, float64(0)) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v70 = F_pg_prng_double(m, v14)
	mBase = m.M
	if base.F64_eq(v70, float64(0)) != 0 {
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	goto L32
L34:
	;
	v79 = base.I32_trunc_f64_u(v73)
	v81 = v79
	goto L29
L35:
	;
	goto L36
L36:
	;
	v81 = int32(0)
	goto L29
L37:
	;
	__phi84 = v81
	__phi88 = v35
	v84 = __phi84
	v88 = __phi88
	goto L38
L38:
	;
	v95 = base.I32_rem_u_s(v88, v84)
	if v95 != 0 {
		__phi84 = v95
		__phi88 = v84
		v84 = __phi84
		v88 = __phi88
		goto L38
	} else {
		goto L40
	}
L39:
	;
	if base.Ui32(int32(1)) < base.Ui32(v84) {
		goto L22
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L23
L42:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.F64_le(v151, base.F64_div(base.F64_convert_i64_s(v152+v153*int64(1000000000)-v157), float64(1e+06))) != 0 {
		v184 = v146
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+36)))
	v165 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+44)))
	v166 = v163
	goto L44
L44:
	;
	v179 = base.I64_rem_u_s(v165+base.I64_extend_i32_u(v166), v164)
	v180 = base.I32_wrap_i64(v179)
	if base.Ui32(l1) <= base.Ui32(v180) {
		v166 = v180
		goto L44
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v180
	v184 = v180
	goto L1
L46:
	;
	goto L45
}
