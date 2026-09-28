package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_close_lseg(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = F_point_sl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if base.F64_eq(v9, v15) != 0 {
				v18 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
				return int64(0)
			} else {
				v23 = F_palloc(m, int32(16))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = F_lseg_closept_lseg(m, v23, v5, v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v25)&int64(9223372036854775807)) {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v23)
						}
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
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 float64
	_ = v67
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 float64
	_ = v93
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
			v93 = float64(0)
			m.G0 = v9 + int32(16)
			return v93
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
					v22 = int64(9223372036854775807)
					v27 = int32(0)
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&v22))|base.B2i32(base.F64_gt(v15, v19) == v27)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v15)&v22) < base.Ui64(int64(9218868437227405313))) == v27 {
						if l0 != 0 {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v38
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v40
						} else {
						}
						v42 = v19
					} else {
						v42 = v15
					}
					v44 = F_lseg_closept_point(m, int32(0), l2, l1)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return float64(0)
					} else {
						v47 = int64(9223372036854775807)
						v52 = int32(0)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v44)&v47))|base.B2i32(base.F64_gt(v42, v44) == v52)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v42)&v47) < base.Ui64(int64(9218868437227405313))) == v52 {
							if l0 != 0 {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v63
								v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v65
							} else {
							}
							v67 = v44
						} else {
							v67 = v42
						}
						v70 = l1 + int32(16)
						v71 = F_lseg_closept_point(m, int32(0), l2, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return float64(0)
						} else {
							v74 = int64(9223372036854775807)
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v71)&v74))|base.B2i32(base.F64_gt(v67, v71) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v67)&v74) < base.Ui64(int64(9218868437227405313))) != 0 {
								v93 = v67
							} else {
								if l0 != 0 {
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v88
									v90 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v90
								} else {
								}
								v93 = v71
							}
							m.G0 = v9 + int32(16)
							return v93
						}
					}
				}
			}
		}
	}
}
func F_lseg_crossing(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) int32 {
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v18 float64
	_ = v18
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v64 int32
	_ = v64
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v106 float64
	_ = v106
	var v108 float64
	_ = v108
	var v120 float64
	_ = v120
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v139 float64
	_ = v139
	var v140 int32
	_ = v140
	var v141 float64
	_ = v141
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v163 float64
	_ = v163
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	var v169 float64
	_ = v169
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v192 float64
	_ = v192
	var v193 int32
	_ = v193
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v200 float64
	_ = v200
	var v207 float64
	_ = v207
	var v208 int32
	_ = v208
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v218 float64
	_ = v218
	var v237 int32
	_ = v237
	v8 = int32(0)
	v10 = base.F64_abs(l1)
	if base.F64_le(v10, float64(1e-06)) != 0 {
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
	v18 = base.F64_abs(l3)
	if base.F64_gt(l0, float64(1e-06)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if base.F64_le(v18, float64(1e-06)) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if base.F64_le(v18, float64(1e-06)) == int32(0) {
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
	v27 = int32(0)
	goto L15
L14:
	;
	v27 = int32(2147483647)
	goto L15
L15:
	;
	return v27
L16:
	;
	v35 = int32(1)
	goto L18
L17:
	;
	v35 = int32(-1)
	goto L18
L18:
	;
	return v35
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
	v49 = int32(0)
	goto L24
L23:
	;
	v49 = int32(2147483647)
	goto L24
L24:
	;
	return v49
L25:
	;
	v55 = int32(1)
	goto L27
L26:
	;
	v55 = int32(-1)
	goto L27
L27:
	;
	v56 = base.F64_abs(l3)
	if base.F64_le(v56, float64(1e-06)) != 0 {
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
	v64 = int32(0)
	goto L33
L32:
	;
	v64 = v55
	goto L33
L33:
	;
	return v64
L34:
	;
	v84 = float64(1e-06)
	v85 = base.F64_add(l0, v84)
	v88 = int32(0)
	if base.B2i32(base.F64_ge(v85, float64(0)) == v88)|base.B2i32(base.F64_gt(l2, v84) == v88) == v88 {
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
	return v55 << (uint(int32(1)) % 32)
L41:
	;
	goto L42
L42:
	;
	if base.F64_lt(v85, float64(0))&base.F64_le(l2, float64(1e-06)) != 0 {
		v237 = v8
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v237
L44:
	;
	v106 = math.Float64frombits(uint64(0x7ff0000000000000))
	v108 = base.F64_sub(l0, l2)
	if base.F64_eq(base.F64_abs(l0), v106)|base.F64_ne(base.F64_abs(v108), v106)|base.F64_eq(base.F64_abs(l2), v106) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v120 = F_float_overflow_error_ext(m, int32(0))
	v123 = m.ExcPending
	if v123 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v124 = v108
	goto L47
L47:
	;
	v125 = math.Float64frombits(uint64(0x7ff0000000000000))
	v127 = base.F64_mul(l1, v124)
	if base.F64_eq(v10, v125)|base.F64_ne(base.F64_abs(v127), v125)|base.F64_eq(base.F64_abs(v124), v125) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	return int32(0)
L49:
	;
	v124 = v120
	goto L47
L50:
	;
	v150 = math.Float64frombits(uint64(0x7ff0000000000000))
	v152 = base.F64_sub(l1, l3)
	if base.F64_eq(v10, v150)|base.F64_ne(base.F64_abs(v152), v150)|base.F64_eq(v56, v150) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v139 = F_float_overflow_error_ext(m, int32(0))
	v140 = m.ExcPending
	if v140 != 0 {
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v141 = float64(0)
	if base.F64_eq(v124, v141)|base.F64_ne(v127, v141) != 0 {
		v149 = v127
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v149 = v139
	goto L50
L55:
	;
	v147 = F_float_underflow_error_ext(m, int32(0))
	v148 = m.ExcPending
	if v148 != 0 {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v149 = v147
	goto L50
L57:
	;
	v163 = F_float_overflow_error_ext(m, int32(0))
	v164 = m.ExcPending
	if v164 != 0 {
		goto L48
	} else {
		goto L60
	}
L58:
	;
	v165 = v152
	goto L59
L59:
	;
	v167 = math.Float64frombits(uint64(0x7ff0000000000000))
	v169 = base.F64_mul(l0, v165)
	if base.F64_eq(base.F64_abs(l0), v167)|base.F64_ne(base.F64_abs(v169), v167)|base.F64_eq(base.F64_abs(v165), v167) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v165 = v163
	goto L59
L61:
	;
	v195 = base.F64_sub(v149, v194)
	v196 = base.F64_abs(v195)
	if base.F64_eq(v196, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L69
	} else {
		goto L70
	}
L62:
	;
	v181 = F_float_overflow_error_ext(m, int32(0))
	v182 = m.ExcPending
	if v182 != 0 {
		goto L48
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v183 = float64(0)
	if base.F64_eq(l0, v183)|base.F64_ne(v169, v183)|base.F64_eq(v165, v183) != 0 {
		v194 = v169
		goto L61
	} else {
		goto L66
	}
L65:
	;
	v194 = v181
	goto L61
L66:
	;
	v192 = F_float_underflow_error_ext(m, int32(0))
	v193 = m.ExcPending
	if v193 != 0 {
		goto L48
	} else {
		goto L67
	}
L67:
	;
	v194 = v192
	goto L61
L68:
	;
	if base.F64_gt(l1, float64(1e-06)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	v200 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v149), v200)|base.F64_eq(base.F64_abs(v194), v200) != 0 {
		v218 = v195
		goto L68
	} else {
		goto L72
	}
L70:
	;
	v210 = v195
	v211 = v196
	goto L71
L71:
	;
	if base.F64_le(v211, float64(1e-06)) == int32(0) {
		v218 = v210
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v207 = F_float_overflow_error_ext(m, int32(0))
	v208 = m.ExcPending
	if v208 != 0 {
		goto L48
	} else {
		goto L73
	}
L73:
	;
	v210 = v207
	v211 = base.F64_abs(v207)
	goto L71
L74:
	;
	return int32(2147483647)
L75:
	;
	v237 = v55 << (uint(int32(1)) % 32)
	goto L43
L76:
	;
	if base.F64_lt(base.F64_add(v218, float64(1e-06)), float64(0)) == int32(0) {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if base.F64_gt(v218, float64(1e-06)) != 0 {
		v237 = v8
		goto L43
	} else {
		goto L80
	}
L79:
	;
	v237 = v8
	goto L43
L80:
	;
	goto L75
}
func F_lseg_interpt(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_lseg_interpt_lseg(m, v8, v6, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			if v12 == int32(0) {
				v16 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v8)
			}
		}
	}
}
func F_lseg_vertical(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	return base.I64_extend_i32_u(base.F64_eq(v5, v6) | base.F64_le(base.F64_abs(base.F64_sub(v5, v6)), float64(1e-06)))
}
