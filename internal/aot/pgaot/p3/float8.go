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
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(321315)
				F_errmsg_internal(m, int32(25787), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(321315)
					F_errmsg_internal(m, int32(25787), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(321315)
						F_errmsg_internal(m, int32(25787), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(321315)
							F_errmsg_internal(m, int32(25787), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(367457)
	F_errmsg_internal(m, int32(25787), v22+int32(16))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(367457)
	F_errmsg_internal(m, int32(25787), v22)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(26674)
				F_errmsg_internal(m, int32(25787), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(26674)
					F_errmsg_internal(m, int32(25787), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(26674)
						F_errmsg_internal(m, int32(25787), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(26674)
							F_errmsg_internal(m, int32(25787), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
				F_errmsg_internal(m, int32(25787), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
					F_errmsg_internal(m, int32(25787), v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
						F_errmsg_internal(m, int32(25787), v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
							F_errmsg_internal(m, int32(25787), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(230839)
				F_errmsg_internal(m, int32(25787), v7)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(230839)
					F_errmsg_internal(m, int32(25787), v7)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(230839)
						F_errmsg_internal(m, int32(25787), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(230839)
							F_errmsg_internal(m, int32(25787), v7)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233653)
				F_errmsg_internal(m, int32(25787), v7)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233653)
					F_errmsg_internal(m, int32(25787), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233653)
						F_errmsg_internal(m, int32(25787), v7)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233653)
							F_errmsg_internal(m, int32(25787), v7)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485076), int32(2938), int32(24327))
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
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v421 int32
	_ = v421
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
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(96)
	return v421
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v53-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int32(0)
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23&v41 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v41)%32)) - v41
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v53 = v40
	goto L2
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v59 = F_cstring_to_text(m, int32(731620))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v65 = F_palloc0(m, v53<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v421 = v59
	goto L1
L19:
	;
	v71 = F_NUM_cache(m, v53, v14+int32(60), v19, v14+int32(59))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if v73&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v343 = v65 + int32(4)
	F_NUM_processor(m, v71, v14+int32(60), v343, v332, int32(0), v335, v334, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L112
	}
L22:
	;
	v76 = base.F64_nearest(v17)
	if base.F64_lt(base.F64_abs(v76), float64(2.147483648e+09)) != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if v73&int32(16384) != 0 {
		goto L39
	} else {
		goto L40
	}
L25:
	;
	if base.Ui64(base.I64_reinterpret_f64(v76)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v85 = base.I32_trunc_f64_s(v76)
	v87 = v85
	goto L25
L27:
	;
	goto L28
L28:
	;
	v87 = int32(-2147483648)
	goto L25
L29:
	;
	v89 = v87
	goto L31
L30:
	;
	v89 = int32(2147483647)
	goto L31
L31:
	;
	if base.F64_ge(v76, float64(-2.147483648e+09)) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = v89
	goto L34
L33:
	;
	v93 = int32(2147483647)
	goto L34
L34:
	;
	if base.F64_lt(v76, float64(2.147483648e+09)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v97 = v93
	goto L37
L36:
	;
	v97 = int32(2147483647)
	goto L37
L37:
	;
	v98 = F_int_to_roman(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v332 = v98
	v334 = v2
	v335 = v2
	goto L21
L39:
	;
	v102 = base.F64_abs(v17)
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v102)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v102, math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v73&int32(2048) != 0 {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v113 = v111 + v112
	v116 = F_palloc(m, v113+int32(7))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v132
	*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = v17
	v138 = F_psprintf(m, int32(412924), v14+int32(32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	v120 = v113 + int32(6)
	v122 = F__emscripten_memset_bulkmem(m, v116, base.I32_extend8_s(int32(35)), v120)
	mBase = m.M
	goto L46
L46:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122+v120))) = uint8(v124)
	v126 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v126)
	v129 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v111+v122)+1)) = uint8(v129)
	v332 = v116
	v334 = v124
	v335 = v2
	goto L21
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v140 != int32(43) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v332 = v138
	v334 = v2
	v335 = v2
	goto L21
L49:
	;
	goto L50
L50:
	;
	v143 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v143)
	v332 = v138
	v334 = v2
	v335 = v2
	goto L21
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v147 + v148
	v153 = F_pow(m, float64(10), base.F64_convert_i32_s(v147))
	mBase = m.M
	v156 = base.F64_mul(v17, v153)
	goto L53
L52:
	;
	v156 = v17
	goto L53
L53:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v156)
	v162 = F_psprintf(m, int32(334303), v14+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L55
	}
L54:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v231
	v235 = F_psprintf(m, int32(334308), v14)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L78
	}
L55:
	;
	if v162&int32(3) == int32(0) {
		v187 = v162
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v220 <= int32(14) {
		goto L73
	} else {
		goto L74
	}
L57:
	;
	v220 = v212 - v162
	goto L56
L58:
	;
	v191 = v187
	goto L67
L59:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v171 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v220 = int32(0)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v176 = v162
	goto L63
L63:
	;
	v180 = v176 + int32(1)
	if v180&int32(3) == int32(0) {
		v187 = v180
		goto L58
	} else {
		goto L65
	}
L64:
	;
	v212 = v180
	goto L57
L65:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v185 != 0 {
		v176 = v180
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v200 = int32(-2139062144)
	if (int32(16843008)-v197|v197)&v200 == v200 {
		v191 = v191 + int32(4)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v206 = v191
	goto L70
L69:
	;
	goto L68
L70:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v210 != 0 {
		v206 = v206 + int32(1)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v212 = v206
	goto L57
L72:
	;
	goto L71
L73:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if v223+v220 < int32(16) {
		v231 = v223
		goto L54
	} else {
		goto L76
	}
L74:
	;
	v229 = v2
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v229
	v231 = v229
	goto L54
L76:
	;
	v229 = int32(15) - v220
	goto L75
L77:
	;
	if v237 == int32(45) {
		goto L103
	} else {
		goto L104
	}
L78:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	v239 = base.B2i32(v237 == int32(45))
	v240 = v235 + v239
	v241 = int32(46)
	v242 = F___strchrnul(m, v240, v241)
	mBase = m.M
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v244 == v241 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v248 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v248 = v242
	goto L82
L81:
	;
	v248 = int32(0)
	goto L82
L82:
	;
	goto L79
L83:
	;
	v307 = v248 - v240
	goto L77
L84:
	;
	goto L85
L85:
	;
	if v240&int32(3) == int32(0) {
		v273 = v240
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v307 = v306
	goto L77
L87:
	;
	v306 = v298 - v240
	goto L86
L88:
	;
	v277 = v273
	goto L97
L89:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v257 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v306 = int32(0)
	goto L86
L91:
	;
	goto L92
L92:
	;
	v262 = v240
	goto L93
L93:
	;
	v266 = v262 + int32(1)
	if v266&int32(3) == int32(0) {
		v273 = v266
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v298 = v266
	goto L87
L95:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v271 != 0 {
		v262 = v266
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v286 = int32(-2139062144)
	if (int32(16843008)-v283|v283)&v286 == v286 {
		v277 = v277 + int32(4)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v292 = v277
	goto L100
L99:
	;
	goto L98
L100:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v296 != 0 {
		v292 = v292 + int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v298 = v292
	goto L87
L102:
	;
	goto L101
L103:
	;
	v310 = int32(45)
	goto L105
L104:
	;
	v310 = int32(43)
	goto L105
L105:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v307 < v311 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v332 = v240
	v334 = v310
	v335 = v311 - v307
	goto L21
L107:
	;
	goto L108
L108:
	;
	v314 = int32(0)
	if v307 <= v311 {
		v332 = v240
		v334 = v310
		v335 = v314
		goto L21
	} else {
		goto L109
	}
L109:
	;
	v316 = v231 + v311
	v319 = F_palloc(m, v316+int32(2))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	v323 = v316 + int32(1)
	v325 = F__emscripten_memset_bulkmem(m, v319, base.I32_extend8_s(int32(35)), v323)
	mBase = m.M
	goto L111
L111:
	;
	v327 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v325+v323))) = uint8(v327)
	v330 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v325+v311))) = uint8(v330)
	v332 = v319
	v334 = v310
	v335 = v314
	goto L21
L112:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
	if v348 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_pfree(m, v71)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v343&int32(3) == int32(0) {
		v376 = v343
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v409<<(uint(int32(2))%32) + int32(16)
	v421 = v65
	goto L1
L118:
	;
	v409 = v401 - v343
	goto L117
L119:
	;
	v380 = v376
	goto L128
L120:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if v360 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v409 = int32(0)
	goto L117
L122:
	;
	goto L123
L123:
	;
	v365 = v343
	goto L124
L124:
	;
	v369 = v365 + int32(1)
	if v369&int32(3) == int32(0) {
		v376 = v369
		goto L119
	} else {
		goto L126
	}
L125:
	;
	v401 = v369
	goto L118
L126:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	if v374 != 0 {
		v365 = v369
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v389 = int32(-2139062144)
	if (int32(16843008)-v386|v386)&v389 == v389 {
		v380 = v380 + int32(4)
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v395 = v380
	goto L131
L130:
	;
	goto L129
L131:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v399 != 0 {
		v395 = v395 + int32(1)
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v401 = v395
	goto L118
L133:
	;
	goto L132
}
