package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_close_lseg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_point_sl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if base.F64_eq(v9, v15) != 0 {
				v18 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
				v34 = int32(0)
				return v34
			} else {
				v21 = F_palloc(m, int32(16))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_lseg_closept_lseg(m, v21, v5, v6)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v34 = v21
						} else {
							v30 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
							v34 = int32(0)
						}
						return v34
					}
				}
			}
		}
	}
}
func F_lseg_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
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
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int32
	_ = v64
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 float64
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_lseg_interpt_lseg(m, l0, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return float64(0)
	} else {
		if v11 != 0 {
			v83 = float64(0)
			m.G0 = v9 + int32(16)
			return v83
		} else {
			v15 = F_lseg_closept_point(m, l0, l1, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return float64(0)
			} else {
				v19 = F_lseg_closept_point(m, v9, l1, l2+int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return float64(0)
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
						v38 = v15
					} else {
						if base.F64_gt(v15, v19) == int32(0) {
							if base.Ui64(base.I64_reinterpret_f64(v15)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v38 = v15
							} else {
								if l0 != 0 {
									v34 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v34
									v36 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v36
								} else {
								}
								v38 = v19
							}
						} else {
							if l0 != 0 {
								v34 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v34
								v36 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v36
							} else {
							}
							v38 = v19
						}
					}
					v40 = F_lseg_closept_point(m, int32(0), l2, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return float64(0)
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)) {
							v59 = v38
						} else {
							if base.F64_gt(v38, v40) == int32(0) {
								if base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v59 = v38
								} else {
									if l0 != 0 {
										v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v55
										v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v57
									} else {
									}
									v59 = v40
								}
							} else {
								if l0 != 0 {
									v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v55
									v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v57
								} else {
								}
								v59 = v40
							}
						}
						v62 = l1 + int32(16)
						v63 = F_lseg_closept_point(m, int32(0), l2, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return float64(0)
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v63)&int64(9223372036854775807)) {
								v83 = v59
							} else {
								if base.F64_gt(v59, v63) == int32(0) {
									if base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v83 = v59
									} else {
										if l0 != 0 {
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v78
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v80
										} else {
										}
										v83 = v63
									}
								} else {
									if l0 != 0 {
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v78
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v80
									} else {
									}
									v83 = v63
								}
							}
							m.G0 = v9 + int32(16)
							return v83
						}
					}
				}
			}
		}
	}
}
func F_lseg_crossing(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) int32 {
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v20 float64
	_ = v20
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v66 int32
	_ = v66
	var v87 float64
	_ = v87
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v122 float64
	_ = v122
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	v10 = int32(0)
	v12 = base.F64_abs(l1)
	if base.F64_le(v12, float64(1e-06)) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if base.F64_le(base.F64_abs(l0), float64(1e-06)) != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if base.F64_gt(l1, float64(1e-06)) != 0 {
		goto L25
	} else {
		goto L26
	}
L4:
	;
	return int32(2147483647)
L5:
	;
	goto L6
L6:
	;
	v20 = base.F64_abs(l3)
	if base.F64_gt(l0, float64(1e-06)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if base.F64_le(v20, float64(1e-06)) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if base.F64_le(v20, float64(1e-06)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	if base.F64_gt(l2, float64(1e-06)) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if base.F64_lt(base.F64_add(l3, float64(1e-06)), float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v29 = int32(0)
	goto L15
L14:
	;
	v29 = int32(2147483647)
	goto L15
L15:
	;
	return v29
L16:
	;
	v37 = int32(1)
	goto L18
L17:
	;
	v37 = int32(-1)
	goto L18
L18:
	;
	return v37
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	if base.F64_lt(base.F64_add(l2, float64(1e-06)), float64(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v51 = int32(0)
	goto L24
L23:
	;
	v51 = int32(2147483647)
	goto L24
L24:
	;
	return v51
L25:
	;
	v57 = int32(1)
	goto L27
L26:
	;
	v57 = int32(-1)
	goto L27
L27:
	;
	v58 = base.F64_abs(l3)
	if base.F64_le(v58, float64(1e-06)) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.F64_lt(base.F64_add(l2, float64(1e-06)), float64(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if base.F64_gt(l1, float64(1e-06)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v66 = int32(0)
	goto L33
L32:
	;
	v66 = v57
	goto L33
L33:
	;
	return v66
L34:
	;
	v87 = base.F64_add(l0, float64(1e-06))
	if base.F64_ge(v87, float64(0)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	if base.F64_lt(base.F64_add(l3, float64(1e-06)), float64(0)) == int32(0) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.F64_gt(l3, float64(1e-06)) == int32(0) {
		goto L34
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	return int32(0)
L40:
	;
	if base.F64_lt(v87, float64(0))&base.F64_le(l2, float64(1e-06)) != 0 {
		v183 = v10
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if base.F64_gt(l2, float64(1e-06)) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	return v57 << (uint(int32(1)) % 32)
L43:
	;
	F_float_underflow_error(m)
	v201 = m.ExcPending
	if v201 != 0 {
		goto L81
	} else {
		goto L83
	}
L44:
	;
	F_float_overflow_error(m)
	v195 = m.ExcPending
	if v195 != 0 {
		goto L81
	} else {
		goto L82
	}
L45:
	;
	return v183
L46:
	;
	v104 = base.F64_sub(l0, l2)
	v105 = base.F64_abs(v104)
	if base.F64_ne(v105, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v114 = base.F64_mul(l1, v104)
	v115 = base.F64_abs(v114)
	if base.F64_ne(v115, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	if base.F64_eq(base.F64_abs(l0), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if base.F64_ne(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v122 = float64(0)
	if base.F64_eq(v114, v122)&base.F64_ne(v104, v122) != 0 {
		goto L43
	} else {
		goto L55
	}
L52:
	;
	if base.F64_eq(v12, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if base.F64_ne(v105, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v127 = base.F64_sub(l1, l3)
	v128 = base.F64_abs(v127)
	if base.F64_ne(v128, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v135 = base.F64_mul(l0, v127)
	v136 = base.F64_abs(v135)
	if base.F64_ne(v136, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if base.F64_eq(v12, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if base.F64_ne(v58, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	if base.F64_ne(v135, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if base.F64_eq(base.F64_abs(l0), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if base.F64_ne(v128, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v150 = base.F64_sub(v114, v135)
	v151 = base.F64_abs(v150)
	if base.F64_ne(v151, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if base.F64_eq(l0, float64(0)) != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if base.F64_ne(v127, float64(0)) != 0 {
		goto L43
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	if base.F64_le(v151, float64(1e-06)) != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if base.F64_eq(v115, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if base.F64_ne(v136, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	return int32(2147483647)
L73:
	;
	goto L74
L74:
	;
	if base.F64_gt(l1, float64(1e-06)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v183 = v57 << (uint(int32(1)) % 32)
	goto L45
L76:
	;
	if base.F64_lt(base.F64_add(v150, float64(1e-06)), float64(0)) == int32(0) {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if base.F64_gt(v150, float64(1e-06)) != 0 {
		v183 = v10
		goto L45
	} else {
		goto L80
	}
L79:
	;
	v183 = v10
	goto L45
L80:
	;
	goto L75
L81:
	;
	return int32(0)
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lseg_interpt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_lseg_interpt_lseg(m, v8, v6, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v17 = v8
			} else {
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
				v17 = int32(0)
			}
			return v17
		}
	}
}
func F_lseg_vertical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	return base.F64_eq(v5, v6) | base.F64_le(base.F64_abs(base.F64_sub(v5, v6)), float64(1e-06))
}
