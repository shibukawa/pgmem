package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_avg(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(326331)
				F_errmsg_internal(m, int32(26146), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(326331)
					F_errmsg_internal(m, int32(26146), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(326331)
						F_errmsg_internal(m, int32(26146), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(326331)
							F_errmsg_internal(m, int32(26146), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_eq(v24, float64(0)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v34 = int32(0)
							m.G0 = v7 + int32(16)
							return v34
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+32))
							v32 = F_Float8GetDatum(m, base.F64_div(v30, v24))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = v32
								m.G0 = v7 + int32(16)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_cmp_internal(m *base.Module, l0 float64, l1 float64) int32 {
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	v7 = int64(9223372036854775807)
	v8 = base.I64_reinterpret_f64(l0) & v7
	v11 = base.I64_reinterpret_f64(l1) & v7
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v11) {
		v20 = base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313)))
		v28 = int32(0) - v20&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
	} else {
		v16 = int32(1)
		if base.F64_gt(l0, l1) != 0 {
			v28 = v16
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v8) {
				v28 = v16
			} else {
				v20 = v16
				v28 = int32(0) - v20&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
			}
		}
	}
	return v28
}
func F_float8_regr_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v104 float64
	_ = v104
	var v118 float64
	_ = v118
	var v122 float64
	_ = v122
	var v136 float64
	_ = v136
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v208 float64
	_ = v208
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v212 float64
	_ = v212
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v248 int32
	_ = v248
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	v20 = m.G0
	v22 = v20 - int32(112)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v32 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v220 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L5:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22)+96)) = v195
	*(*float64)(unsafe.Add(mBase, uint32(v22)+104)) = v193
	*(*float64)(unsafe.Add(mBase, uint32(v22)+88)) = v190
	*(*float64)(unsafe.Add(mBase, uint32(v22)+80)) = v192
	*(*float64)(unsafe.Add(mBase, uint32(v22)+72)) = v191
	*(*float64)(unsafe.Add(mBase, uint32(v22)+64)) = v189
	v206 = v189
	v207 = v190
	v208 = v191
	v209 = v192
	v210 = v193
	v212 = v195
	goto L4
L6:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L48
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v35 != int32(6) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v38 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v39 != int32(701) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v42 != int32(1) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v45 != int32(6) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v48 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v49 != int32(701) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v30-int32(-64))))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v30)+56))
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v30)+48))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v30)+40))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v30)+32))
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v25)+24))
	if base.F64_eq(v60, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v189 = v54
	v190 = v57
	v191 = v55
	v192 = v56
	v193 = v59
	v195 = v58
	goto L5
L18:
	;
	goto L19
L19:
	;
	v65 = *(*float64)(unsafe.Add(mBase, uint32(v25-int32(-64))))
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v25)+56))
	v67 = *(*float64)(unsafe.Add(mBase, uint32(v25)+48))
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v25)+40))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v25)+32))
	if base.F64_eq(v59, float64(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v189 = v65
	v190 = v68
	v191 = v66
	v192 = v67
	v193 = v60
	v195 = v69
	goto L5
L21:
	;
	goto L22
L22:
	;
	v72 = base.F64_add(v60, v59)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+104)) = v72
	v74 = base.F64_add(v69, v58)
	if base.F64_ne(base.F64_abs(v74), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22)+96)) = v74
	v88 = base.F64_sub(base.F64_div(v69, v60), base.F64_div(v58, v59))
	v89 = base.F64_mul(v60, v59)
	v90 = base.F64_mul(v88, v89)
	v93 = base.F64_add(base.F64_add(v68, v57), base.F64_div(base.F64_mul(v88, v90), v72))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+88)) = v93
	if base.F64_ne(base.F64_abs(v93), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if base.F64_eq(base.F64_abs(v69), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.F64_ne(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v104 = base.F64_add(v67, v56)
	if base.F64_ne(base.F64_abs(v104), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if base.F64_eq(base.F64_abs(v68), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if base.F64_ne(base.F64_abs(v57), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22)+80)) = v104
	v118 = base.F64_sub(base.F64_div(v67, v60), base.F64_div(v56, v59))
	v122 = base.F64_add(base.F64_add(v66, v55), base.F64_div(base.F64_mul(v118, base.F64_mul(v89, v118)), v72))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+72)) = v122
	if base.F64_ne(base.F64_abs(v122), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if base.F64_ne(base.F64_abs(v56), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v136 = base.F64_add(base.F64_add(v65, v54), base.F64_div(base.F64_mul(v90, v118), v72))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+64)) = v136
	if base.F64_ne(base.F64_abs(v136), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v206 = v136
		v207 = v93
		v208 = v122
		v209 = v104
		v210 = v72
		v212 = v74
		goto L4
	} else {
		goto L39
	}
L36:
	;
	if base.F64_eq(base.F64_abs(v66), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if base.F64_ne(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	if base.F64_eq(base.F64_abs(v65), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v206 = v136
		v207 = v93
		v208 = v122
		v209 = v104
		v210 = v72
		v212 = v74
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if base.F64_ne(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v206 = v136
	v207 = v93
	v208 = v122
	v209 = v104
	v210 = v72
	v212 = v74
	goto L4
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(373287)
	F_errmsg_internal(m, int32(26146), v22+int32(16))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(492922), int32(2938), int32(24686))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(373287)
	F_errmsg_internal(m, int32(26146), v22)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(492922), int32(2938), int32(24686))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	m.G0 = v22 + int32(112)
	return v281
L50:
	;
	if v248 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v248 = int32(0)
	goto L50
L53:
	;
	goto L51
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	switch v223 - int32(429) {
	case 0:
		goto L56
	case 1:
		goto L55
	default:
		goto L53
	}
L55:
	;
	goto L60
L56:
	;
	goto L57
L57:
	;
	v248 = int32(1)
	goto L50
L60:
	;
	v248 = int32(2)
	goto L50
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v25-int32(-64)))) = v206
	*(*float64)(unsafe.Add(mBase, uint32(v25)+56)) = v208
	*(*float64)(unsafe.Add(mBase, uint32(v25)+48)) = v209
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = v207
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v212
	*(*float64)(unsafe.Add(mBase, uint32(v25)+24)) = v210
	v281 = v25
	goto L49
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v22 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v22 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v22 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v22 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v22 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v22 + int32(104)
	v279 = F_construct_array_builtin(m, v22+int32(32), int32(6), int32(701))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v281 = v279
	goto L49
}
func F_float8_regr_sxx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(27056)
				F_errmsg_internal(m, int32(26146), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			if v16 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(27056)
					F_errmsg_internal(m, int32(26146), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(27056)
						F_errmsg_internal(m, int32(26146), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v20 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(27056)
							F_errmsg_internal(m, int32(26146), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						if base.F64_lt(v23, float64(1)) != 0 {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v32 = int32(0)
							m.G0 = v6 + int32(16)
							return v32
						} else {
							v29 = *(*float64)(unsafe.Add(mBase, uint32(v9)+40))
							v30 = F_Float8GetDatum(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v30
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_syy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8027)
				F_errmsg_internal(m, int32(26146), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			if v16 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8027)
					F_errmsg_internal(m, int32(26146), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8027)
						F_errmsg_internal(m, int32(26146), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v20 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8027)
							F_errmsg_internal(m, int32(26146), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						if base.F64_lt(v23, float64(1)) != 0 {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v32 = int32(0)
							m.G0 = v6 + int32(16)
							return v32
						} else {
							v29 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
							v30 = F_Float8GetDatum(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v30
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_stddev_pop(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(234232)
				F_errmsg_internal(m, int32(26146), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(234232)
					F_errmsg_internal(m, int32(26146), v7)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(234232)
						F_errmsg_internal(m, int32(26146), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(234232)
							F_errmsg_internal(m, int32(26146), v7)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_eq(v24, float64(0)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v35 = int32(0)
							m.G0 = v7 + int32(16)
							return v35
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v33 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_div(v30, v24)))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = v33
								m.G0 = v7 + int32(16)
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_stddev_samp(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237114)
				F_errmsg_internal(m, int32(26146), v7)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492922), int32(2938), int32(24686))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237114)
					F_errmsg_internal(m, int32(26146), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492922), int32(2938), int32(24686))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237114)
						F_errmsg_internal(m, int32(26146), v7)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492922), int32(2938), int32(24686))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237114)
							F_errmsg_internal(m, int32(26146), v7)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492922), int32(2938), int32(24686))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_le(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v37 = int32(0)
							m.G0 = v7 + int32(16)
							return v37
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v35 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_div(v30, base.F64_add(v24, float64(-1)))))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = v35
								m.G0 = v7 + int32(16)
								return v37
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 float64
	_ = v76
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 float64
	_ = v153
	var v156 float64
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		if v23 == int32(1) {
			v26 = int32(4)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
			if v28&int32(254) == int32(2) {
				v37 = v26
			} else {
				v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
			}
			if v28 == int32(1) {
				v40 = v26
			} else {
				v40 = v37
			}
			v53 = v40
		} else {
			v41 = int32(1)
			if v23&v41 != 0 {
				v53 = int32(base.Ui32(v23)>>(uint(v41)%32)) - v41
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v53-int32(1)) {
			v59 = F_cstring_to_text(m, int32(741336))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v253 = v59
				m.G0 = v14 + int32(96)
				return v253
			}
		} else {
			v65 = F_palloc0(m, v53<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v71 = F_NUM_cache(m, v53, v14+int32(60), v19, v14+int32(59))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v73&int32(1024) != 0 {
						v76 = base.F64_nearest(v17)
						if base.F64_lt(base.F64_abs(v76), float64(2.147483648e+09)) != 0 {
							v85 = base.I32_trunc_f64_s(v76)
							v87 = v85
						} else {
							v87 = int32(-2147483648)
						}
						if base.Ui64(base.I64_reinterpret_f64(v76)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v89 = v87
						} else {
							v89 = int32(2147483647)
						}
						if base.F64_ge(v76, float64(-2.147483648e+09)) != 0 {
							v93 = v89
						} else {
							v93 = int32(2147483647)
						}
						if base.F64_lt(v76, float64(2.147483648e+09)) != 0 {
							v97 = v93
						} else {
							v97 = int32(2147483647)
						}
						v98 = F_int_to_roman(m, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v220 = v98
							v222 = v2
							v223 = v2
							v231 = v65 + int32(4)
							F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
							mBase = m.M
							v235 = m.ExcPending
							if v235 != 0 {
								return int32(0)
							} else {
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v236 == int32(1) {
									F_pfree(m, v71)
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										v241 = F_strlen(m, v231)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
										v253 = v65
										m.G0 = v14 + int32(96)
										return v253
									}
								} else {
									v241 = F_strlen(m, v231)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
									v253 = v65
									m.G0 = v14 + int32(96)
									return v253
								}
							}
						}
					} else {
						if v73&int32(16384) != 0 {
							v102 = base.F64_abs(v17)
							if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v102)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v102, math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v113 = v111 + v112
								v116 = F_palloc(m, v113+int32(7))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v120 = v113 + int32(6)
									v122 = F__emscripten_memset_bulkmem(m, v116, base.I32_extend8_s(int32(35)), v120)
									mBase = m.M
									v124 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v122+v120))) = uint8(v124)
									v126 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v126)
									v129 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v111+v122)+1)) = uint8(v129)
									v220 = v116
									v222 = v124
									v223 = v2
									v231 = v65 + int32(4)
									F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int32(0)
									} else {
										v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v236 == int32(1) {
											F_pfree(m, v71)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v241 = F_strlen(m, v231)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
												v253 = v65
												m.G0 = v14 + int32(96)
												return v253
											}
										} else {
											v241 = F_strlen(m, v231)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
											v253 = v65
											m.G0 = v14 + int32(96)
											return v253
										}
									}
								}
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v132
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = v17
								v138 = F_psprintf(m, int32(419362), v14+int32(32))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
									if v140 != int32(43) {
										v220 = v138
										v222 = v2
										v223 = v2
									} else {
										v143 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v143)
										v220 = v138
										v222 = v2
										v223 = v2
									}
									v231 = v65 + int32(4)
									F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int32(0)
									} else {
										v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v236 == int32(1) {
											F_pfree(m, v71)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v241 = F_strlen(m, v231)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
												v253 = v65
												m.G0 = v14 + int32(96)
												return v253
											}
										} else {
											v241 = F_strlen(m, v231)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
											v253 = v65
											m.G0 = v14 + int32(96)
											return v253
										}
									}
								}
							}
						} else {
							if v73&int32(2048) != 0 {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v147 + v148
								v153 = F_pow(m, float64(10), base.F64_convert_i32_s(v147))
								mBase = m.M
								v156 = base.F64_mul(v17, v153)
							} else {
								v156 = v17
							}
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v156)
							v162 = F_psprintf(m, int32(339588), v14+int32(16))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								v164 = F_strlen(m, v162)
								mBase = m.M
								if v164 <= int32(14) {
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if v167+v164 < int32(16) {
										v175 = v167
									} else {
										v173 = int32(15) - v164
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v173
										v175 = v173
									}
								} else {
									v173 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v173
									v175 = v173
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v156
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v175
								v179 = F_psprintf(m, int32(339593), v14)
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
									v183 = base.B2i32(v181 == int32(45))
									v184 = v179 + v183
									v185 = int32(46)
									v186 = F___strchrnul(m, v184, v185)
									mBase = m.M
									v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
									if v188 == v185 {
										v192 = v186
									} else {
										v192 = int32(0)
									}
									if v192 != 0 {
										v195 = v192 - v184
									} else {
										v194 = F_strlen(m, v184)
										mBase = m.M
										v195 = v194
									}
									if v181 == int32(45) {
										v198 = int32(45)
									} else {
										v198 = int32(43)
									}
									v199 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if v195 < v199 {
										v220 = v184
										v222 = v198
										v223 = v199 - v195
										v231 = v65 + int32(4)
										F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
										mBase = m.M
										v235 = m.ExcPending
										if v235 != 0 {
											return int32(0)
										} else {
											v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v236 == int32(1) {
												F_pfree(m, v71)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													v241 = F_strlen(m, v231)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
													v253 = v65
													m.G0 = v14 + int32(96)
													return v253
												}
											} else {
												v241 = F_strlen(m, v231)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
												v253 = v65
												m.G0 = v14 + int32(96)
												return v253
											}
										}
									} else {
										v202 = int32(0)
										if v195 <= v199 {
											v220 = v184
											v222 = v198
											v223 = v202
											v231 = v65 + int32(4)
											F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
											mBase = m.M
											v235 = m.ExcPending
											if v235 != 0 {
												return int32(0)
											} else {
												v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v236 == int32(1) {
													F_pfree(m, v71)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														v241 = F_strlen(m, v231)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
														v253 = v65
														m.G0 = v14 + int32(96)
														return v253
													}
												} else {
													v241 = F_strlen(m, v231)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
													v253 = v65
													m.G0 = v14 + int32(96)
													return v253
												}
											}
										} else {
											v204 = v175 + v199
											v207 = F_palloc(m, v204+int32(2))
											mBase = m.M
											v208 = m.ExcPending
											if v208 != 0 {
												return int32(0)
											} else {
												v211 = v204 + int32(1)
												v213 = F__emscripten_memset_bulkmem(m, v207, base.I32_extend8_s(int32(35)), v211)
												mBase = m.M
												v215 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v213+v211))) = uint8(v215)
												v218 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v213+v199))) = uint8(v218)
												v220 = v207
												v222 = v198
												v223 = v202
												v231 = v65 + int32(4)
												F_NUM_processor(m, v71, v14+int32(60), v231, v220, int32(0), v223, v222, int32(1))
												mBase = m.M
												v235 = m.ExcPending
												if v235 != 0 {
													return int32(0)
												} else {
													v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v236 == int32(1) {
														F_pfree(m, v71)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int32(0)
														} else {
															v241 = F_strlen(m, v231)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
															v253 = v65
															m.G0 = v14 + int32(96)
															return v253
														}
													} else {
														v241 = F_strlen(m, v231)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v65))) = v241<<(uint(int32(2))%32) + int32(16)
														v253 = v65
														m.G0 = v14 + int32(96)
														return v253
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
