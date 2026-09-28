package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v56 float64
	_ = v56
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 float64
	_ = v129
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v162 float64
	_ = v162
	var v173 float64
	_ = v173
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v199 float64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v215 float64
	_ = v215
	var v223 float64
	_ = v223
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v239 float64
	_ = v239
	var v240 float64
	_ = v240
	var v242 float64
	_ = v242
	var v244 float64
	_ = v244
	var v251 float64
	_ = v251
	var v257 float64
	_ = v257
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = base.F64_abs(v13)
	if base.F64_le(v14, float64(1e-06)) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v125 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L32
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v45
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v71 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v72 = base.F64_mul(v45, v71)
	v74 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v72), v74)|base.F64_eq(base.F64_abs(v71), v74) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v65 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v65
	goto L1
L4:
	;
	if base.F64_ne(v45, float64(0)) != 0 {
		goto L2
	} else {
		goto L18
	}
L5:
	;
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v20 = base.F64_abs(v19)
	if base.F64_le(v20, float64(1e-06)) != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(-4616189618054758400)
	v56 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v56
	goto L1
L8:
	;
	v23 = math.Float64frombits(uint64(0x7ff0000000000000))
	v25 = base.F64_div(v19, v13)
	if base.F64_eq(v20, v23)|base.F64_ne(base.F64_abs(v25), v23) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L17
	}
L10:
	;
	v33 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if base.F64_eq(v14, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v25, float64(0)) != 0 {
		v45 = v25
		goto L9
	} else {
		goto L15
	}
L13:
	;
	return float64(0)
L14:
	;
	v45 = v33
	goto L9
L15:
	;
	v43 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v45 = v43
	goto L9
L17:
	;
	goto L7
L18:
	;
	goto L3
L19:
	;
	v95 = math.Float64frombits(uint64(0x7ff0000000000000))
	v97 = base.F64_sub(v70, v93)
	if base.F64_eq(base.F64_abs(v70), v95)|base.F64_ne(base.F64_abs(v97), v95)|base.F64_eq(base.F64_abs(v93), v95) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v83 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v85 = float64(0)
	if base.F64_eq(v71, v85)|base.F64_ne(v72, v85) != 0 {
		v93 = v72
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v93 = v83
	goto L19
L24:
	;
	v91 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v93 = v91
	goto L19
L26:
	;
	v109 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L29
	}
L27:
	;
	v111 = v97
	goto L28
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v111
	if base.F64_ne(v111, float64(0)) != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v111 = v109
	goto L28
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
	goto L1
L31:
	;
	m.G0 = v11 + int32(48)
	return v257
L32:
	;
	if v125 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v129 = math.Float64frombits(uint64(0x7ff8000000000000))
	if l0 == int32(0) {
		v257 = v129
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l0 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v134
	v257 = v129
	goto L31
L37:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v136
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v138
	goto L39
L38:
	;
	goto L39
L39:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
	v141 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v142 = base.F64_sub(v140, v141)
	v144 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v142), v144)|base.F64_eq(base.F64_abs(v140), v144)|base.F64_eq(base.F64_abs(v141), v144) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v157 = v142
	goto L42
L41:
	;
	v155 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L43
	}
L42:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v160 = base.F64_sub(v158, v159)
	v162 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v160), v162)|base.F64_eq(base.F64_abs(v158), v162)|base.F64_eq(base.F64_abs(v159), v162) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v157 = v155
	goto L42
L44:
	;
	v175 = v160
	goto L46
L45:
	;
	v173 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v184 = m.G0
	v186 = v184 - int32(32)
	m.G0 = v186
	v188 = base.F64_abs(v157)
	v189 = base.F64_abs(v175)
	v192 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)))
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v175 = v173
	goto L46
L48:
	;
	v257 = v251
	goto L31
L49:
	;
	m.G0 = v186 + int32(32)
	goto L48
L50:
	;
	v193 = v188
	goto L52
L51:
	;
	v193 = v189
	goto L52
L52:
	;
	v194 = base.I64_reinterpret_f64(v193)
	v196 = int64(base.Ui64(v194) >> (uint(int64(52)) % 64))
	if v196 == int64(2047) {
		v251 = v193
		goto L49
	} else {
		goto L53
	}
L53:
	;
	if base.Ui64(base.I64_reinterpret_f64(v188)) < base.Ui64(base.I64_reinterpret_f64(v189)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v199 = v189
	goto L56
L55:
	;
	v199 = v188
	goto L56
L56:
	;
	if v194 == int64(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v251 = v199
	goto L49
L58:
	;
	v202 = base.I64_reinterpret_f64(v199)
	v204 = int64(base.Ui64(v202) >> (uint(int64(52)) % 64))
	if v204 == int64(2047) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if int32(65) <= base.I32_wrap_i64(v204)-base.I32_wrap_i64(v196) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v251 = base.F64_add(v188, v189)
	goto L49
L61:
	;
	goto L62
L62:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v202) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	F_sq(m, v186+int32(24), v186+int32(16), v228)
	mBase = m.M
	F_sq(m, v186+int32(8), v186, v229)
	mBase = m.M
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v186)+16))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v186)+8))
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v186)+24))
	v251 = base.F64_mul(v230, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v239, v240), v242), v244)))
	goto L49
L64:
	;
	v215 = float64(1.90109156629516e-211)
	v228 = base.F64_mul(v199, v215)
	v229 = base.F64_mul(v193, v215)
	v230 = float64(5.260135901548374e+210)
	goto L63
L65:
	;
	goto L66
L66:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v194) {
		v228 = v199
		v229 = v193
		v230 = float64(1)
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v223 = float64(5.260135901548374e+210)
	v228 = base.F64_mul(v199, v223)
	v229 = base.F64_mul(v193, v223)
	v230 = float64(1.90109156629516e-211)
	goto L63
}
func F_line_construct(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v21 float64
	_ = v21
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v30 float64
	_ = v30
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v47 float64
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4616189618054758400)
		v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v13
		return
	} else {
		if base.F64_eq(l2, float64(0)) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v21
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*float64)(unsafe.Add(mBase, uint32(l0))) = l2
			v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			v28 = base.F64_mul(l2, v27)
			v30 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v28), v30)|base.F64_eq(base.F64_abs(v27), v30) == int32(0) {
				v39 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v49 = v39
					v51 = math.Float64frombits(uint64(0x7ff0000000000000))
					v53 = base.F64_sub(v26, v49)
					if base.F64_eq(base.F64_abs(v26), v51)|base.F64_ne(base.F64_abs(v53), v51)|base.F64_eq(base.F64_abs(v49), v51) == int32(0) {
						v65 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							v67 = v65
							*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
							if base.F64_eq(v67, float64(0)) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							} else {
							}
							return
						}
					} else {
						v67 = v53
						*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
						if base.F64_eq(v67, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
						} else {
						}
						return
					}
				}
			} else {
				v41 = float64(0)
				if base.F64_eq(v27, v41)|base.F64_ne(v28, v41) != 0 {
					v49 = v28
					v51 = math.Float64frombits(uint64(0x7ff0000000000000))
					v53 = base.F64_sub(v26, v49)
					if base.F64_eq(base.F64_abs(v26), v51)|base.F64_ne(base.F64_abs(v53), v51)|base.F64_eq(base.F64_abs(v49), v51) == int32(0) {
						v65 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							v67 = v65
							*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
							if base.F64_eq(v67, float64(0)) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							} else {
							}
							return
						}
					} else {
						v67 = v53
						*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
						if base.F64_eq(v67, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
						} else {
						}
						return
					}
				} else {
					v47 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v49 = v47
						v51 = math.Float64frombits(uint64(0x7ff0000000000000))
						v53 = base.F64_sub(v26, v49)
						if base.F64_eq(base.F64_abs(v26), v51)|base.F64_ne(base.F64_abs(v53), v51)|base.F64_eq(base.F64_abs(v49), v51) == int32(0) {
							v65 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v67 = v65
								*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
								if base.F64_eq(v67, float64(0)) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								} else {
								}
								return
							}
						} else {
							v67 = v53
							*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v67
							if base.F64_eq(v67, float64(0)) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							} else {
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v22 float64
	_ = v22
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v130 float64
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v166 float64
	_ = v166
	var v179 float64
	_ = v179
	var v180 int32
	_ = v180
	var v181 float64
	_ = v181
	var v190 float64
	_ = v190
	var v191 int32
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v210 float64
	_ = v210
	var v211 int32
	_ = v211
	var v212 float64
	_ = v212
	var v221 float64
	_ = v221
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 float64
	_ = v241
	var v253 float64
	_ = v253
	var v254 int32
	_ = v254
	var v256 float64
	_ = v256
	var v258 float64
	_ = v258
	var v266 float64
	_ = v266
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v280 float64
	_ = v280
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v299 float64
	_ = v299
	var v308 float64
	_ = v308
	var v309 int32
	_ = v309
	var v310 float64
	_ = v310
	var v312 float64
	_ = v312
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v327 float64
	_ = v327
	var v328 int32
	_ = v328
	var v329 float64
	_ = v329
	var v335 float64
	_ = v335
	var v342 float64
	_ = v342
	var v343 int32
	_ = v343
	var v345 float64
	_ = v345
	var v348 float64
	_ = v348
	var v356 float64
	_ = v356
	var v357 int32
	_ = v357
	var v358 float64
	_ = v358
	var v368 float64
	_ = v368
	var v369 int32
	_ = v369
	var v370 float64
	_ = v370
	var v371 float64
	_ = v371
	var v374 float64
	_ = v374
	var v375 float64
	_ = v375
	var v376 float64
	_ = v376
	var v381 float64
	_ = v381
	var v382 int32
	_ = v382
	var v383 float64
	_ = v383
	var v392 float64
	_ = v392
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v396 float64
	_ = v396
	var v398 float64
	_ = v398
	var v410 float64
	_ = v410
	var v411 int32
	_ = v411
	var v412 float64
	_ = v412
	var v421 float64
	_ = v421
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v430 float64
	_ = v430
	var v431 float64
	_ = v431
	var v432 float64
	_ = v432
	var v434 float64
	_ = v434
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v449 float64
	_ = v449
	var v458 float64
	_ = v458
	var v459 int32
	_ = v459
	var v460 float64
	_ = v460
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v463 float64
	_ = v463
	var v465 float64
	_ = v465
	var v478 float64
	_ = v478
	var v479 int32
	_ = v479
	var v480 float64
	_ = v480
	var v489 float64
	_ = v489
	var v490 int32
	_ = v490
	var v491 float64
	_ = v491
	var v493 float64
	_ = v493
	var v495 float64
	_ = v495
	var v507 float64
	_ = v507
	var v508 int32
	_ = v508
	var v509 float64
	_ = v509
	var v510 float64
	_ = v510
	var v511 float64
	_ = v511
	var v512 float64
	_ = v512
	var v514 float64
	_ = v514
	var v527 float64
	_ = v527
	var v528 int32
	_ = v528
	var v529 float64
	_ = v529
	var v538 float64
	_ = v538
	var v539 int32
	_ = v539
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v542 float64
	_ = v542
	var v543 float64
	_ = v543
	var v545 float64
	_ = v545
	var v558 float64
	_ = v558
	var v559 int32
	_ = v559
	var v560 float64
	_ = v560
	var v569 float64
	_ = v569
	var v570 int32
	_ = v570
	var v571 float64
	_ = v571
	var v573 float64
	_ = v573
	var v575 float64
	_ = v575
	var v587 float64
	_ = v587
	var v588 int32
	_ = v588
	var v589 float64
	_ = v589
	var v601 float64
	_ = v601
	var v602 int32
	_ = v602
	var v604 float64
	_ = v604
	var v606 float64
	_ = v606
	var v614 float64
	_ = v614
	var v615 int32
	_ = v615
	var v616 float64
	_ = v616
	var v626 float64
	_ = v626
	var v627 int32
	_ = v627
	var v628 float64
	_ = v628
	var v630 float64
	_ = v630
	var v632 float64
	_ = v632
	var v633 float64
	_ = v633
	var v645 float64
	_ = v645
	var v646 int32
	_ = v646
	var v647 float64
	_ = v647
	var v656 float64
	_ = v656
	var v657 int32
	_ = v657
	var v658 float64
	_ = v658
	var v660 float64
	_ = v660
	var v662 float64
	_ = v662
	var v663 float64
	_ = v663
	var v675 float64
	_ = v675
	var v676 int32
	_ = v676
	var v677 float64
	_ = v677
	var v683 float64
	_ = v683
	var v690 float64
	_ = v690
	var v691 int32
	_ = v691
	var v693 float64
	_ = v693
	var v696 float64
	_ = v696
	var v704 float64
	_ = v704
	var v705 int32
	_ = v705
	var v706 float64
	_ = v706
	var v716 float64
	_ = v716
	var v717 int32
	_ = v717
	var v718 float64
	_ = v718
	var v720 float64
	_ = v720
	var v723 int32
	_ = v723
	var v726 float64
	_ = v726
	var v729 float64
	_ = v729
	var v731 float64
	_ = v731
	var v734 float64
	_ = v734
	var v741 int32
	_ = v741
	v10 = int32(0)
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = base.F64_abs(v11)
	if base.F64_le(v12, float64(1e-06)) == v10 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v741
L2:
	;
	v723 = int32(1)
	if l0 == int32(0) {
		v741 = v723
		goto L1
	} else {
		goto L174
	}
L3:
	;
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v20 = base.F64_div(v19, v11)
	v22 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v20), v22)|base.F64_eq(base.F64_abs(v19), v22) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v370 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v371 = base.F64_abs(v370)
	if base.F64_le(v371, float64(1e-06)) != 0 {
		v741 = v10
		goto L1
	} else {
		goto L90
	}
L6:
	;
	v48 = math.Float64frombits(uint64(0x7ff0000000000000))
	v50 = base.F64_mul(v17, v46)
	if base.F64_eq(base.F64_abs(v17), v48)|base.F64_ne(base.F64_abs(v50), v48)|base.F64_eq(base.F64_abs(v46), v48) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v31 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v37 = float64(0)
	if base.F64_eq(v12, math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(v20, v37)|base.F64_eq(v19, v37) != 0 {
		v46 = v20
		goto L6
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v46 = v31
	goto L6
L12:
	;
	v44 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v46 = v44
	goto L6
L14:
	;
	if base.F64_eq(v75, v18)|base.F64_le(base.F64_abs(base.F64_sub(v18, v75)), float64(1e-06)) != 0 {
		v741 = v10
		goto L1
	} else {
		goto L21
	}
L15:
	;
	v62 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v64 = float64(0)
	if base.F64_eq(v17, v64)|base.F64_ne(v50, v64)|base.F64_eq(v46, v64) != 0 {
		v75 = v50
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v75 = v62
	goto L14
L19:
	;
	v73 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v75 = v73
	goto L14
L21:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v83 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v84 = base.F64_mul(v82, v83)
	v86 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v84), v86)|base.F64_eq(base.F64_abs(v82), v86)|base.F64_eq(base.F64_abs(v83), v86) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v115 = base.F64_mul(v113, v114)
	v117 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v115), v117)|base.F64_eq(base.F64_abs(v113), v117)|base.F64_eq(base.F64_abs(v114), v117) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v99 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v101 = float64(0)
	if base.F64_eq(v82, v101)|base.F64_ne(v84, v101)|base.F64_eq(v83, v101) != 0 {
		v112 = v84
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v112 = v99
	goto L22
L27:
	;
	v110 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v112 = v110
	goto L22
L29:
	;
	v145 = math.Float64frombits(uint64(0x7ff0000000000000))
	v147 = base.F64_sub(v112, v143)
	if base.F64_eq(base.F64_abs(v112), v145)|base.F64_ne(base.F64_abs(v147), v145)|base.F64_eq(base.F64_abs(v143), v145) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v130 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v132 = float64(0)
	if base.F64_eq(v113, v132)|base.F64_ne(v115, v132)|base.F64_eq(v114, v132) != 0 {
		v143 = v115
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v143 = v130
	goto L29
L34:
	;
	v141 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v143 = v141
	goto L29
L36:
	;
	v159 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	v161 = v147
	goto L38
L38:
	;
	v162 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v164 = base.F64_mul(v162, v163)
	v166 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v164), v166)|base.F64_eq(base.F64_abs(v162), v166)|base.F64_eq(base.F64_abs(v163), v166) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v161 = v159
	goto L38
L40:
	;
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v194 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v195 = base.F64_mul(v193, v194)
	v197 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v195), v197)|base.F64_eq(base.F64_abs(v193), v197)|base.F64_eq(base.F64_abs(v194), v197) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v179 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v181 = float64(0)
	if base.F64_eq(v162, v181)|base.F64_ne(v164, v181)|base.F64_eq(v163, v181) != 0 {
		v192 = v164
		goto L40
	} else {
		goto L45
	}
L44:
	;
	v192 = v179
	goto L40
L45:
	;
	v190 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v192 = v190
	goto L40
L47:
	;
	v225 = math.Float64frombits(uint64(0x7ff0000000000000))
	v227 = base.F64_sub(v192, v223)
	if base.F64_eq(base.F64_abs(v192), v225)|base.F64_ne(base.F64_abs(v227), v225)|base.F64_eq(base.F64_abs(v223), v225) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v210 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v212 = float64(0)
	if base.F64_eq(v193, v212)|base.F64_ne(v195, v212)|base.F64_eq(v194, v212) != 0 {
		v223 = v195
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v223 = v210
	goto L47
L52:
	;
	v221 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v223 = v221
	goto L47
L54:
	;
	v239 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L57
	}
L55:
	;
	v241 = v227
	goto L56
L56:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)))|base.F64_ne(v241, float64(0)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v241 = v239
	goto L56
L58:
	;
	v282 = math.Float64frombits(uint64(0x7ff0000000000000))
	v284 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v285 = base.F64_mul(v280, v284)
	if base.F64_eq(base.F64_abs(v280), v282)|base.F64_ne(base.F64_abs(v285), v282)|base.F64_eq(base.F64_abs(v284), v282) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	v253 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v256 = math.Float64frombits(uint64(0x7ff0000000000000))
	v258 = base.F64_div(v161, v241)
	if base.F64_eq(base.F64_abs(v161), v256)|base.F64_ne(base.F64_abs(v258), v256) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v280 = v253
	goto L58
L63:
	;
	v266 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v268 = float64(0)
	if base.F64_eq(v161, v268)|base.F64_ne(v258, v268)|base.F64_eq(base.F64_abs(v241), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v280 = v258
		goto L58
	} else {
		goto L67
	}
L66:
	;
	v280 = v266
	goto L58
L67:
	;
	v278 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v280 = v278
	goto L58
L69:
	;
	v312 = math.Float64frombits(uint64(0x7ff0000000000000))
	v314 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v315 = base.F64_add(v310, v314)
	if base.F64_eq(base.F64_abs(v310), v312)|base.F64_ne(base.F64_abs(v315), v312)|base.F64_eq(base.F64_abs(v314), v312) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v297 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v299 = float64(0)
	if base.F64_eq(v280, v299)|base.F64_ne(v285, v299)|base.F64_eq(v284, v299) != 0 {
		v310 = v285
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v310 = v297
	goto L69
L74:
	;
	v308 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v310 = v308
	goto L69
L76:
	;
	v327 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L10
	} else {
		goto L79
	}
L77:
	;
	v329 = v315
	goto L78
L78:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v329)&int64(9223372036854775807)))|base.F64_ne(v335, float64(0)) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v329 = v327
	goto L78
L80:
	;
	v342 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v345 = math.Float64frombits(uint64(0x7ff0000000000000))
	v348 = base.F64_div(base.F64_neg(v329), v335)
	if base.F64_eq(base.F64_abs(v329), v345)|base.F64_ne(base.F64_abs(v348), v345) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v718 = v342
	v720 = v280
	goto L2
L84:
	;
	v356 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L10
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v358 = float64(0)
	if base.F64_eq(v329, v358)|base.F64_ne(v348, v358)|base.F64_eq(base.F64_abs(v335), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v718 = v348
		v720 = v280
		goto L2
	} else {
		goto L88
	}
L87:
	;
	v718 = v356
	v720 = v280
	goto L2
L88:
	;
	v368 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	v718 = v368
	v720 = v280
	goto L2
L90:
	;
	v374 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v375 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v376 = base.F64_div(v11, v370)
	if base.F64_eq(base.F64_abs(v376), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v396 = math.Float64frombits(uint64(0x7ff0000000000000))
	v398 = base.F64_mul(v374, v394)
	if base.F64_eq(base.F64_abs(v374), v396)|base.F64_ne(base.F64_abs(v398), v396)|base.F64_eq(base.F64_abs(v394), v396) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	v381 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L10
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v383 = float64(0)
	if base.F64_eq(v11, v383)|base.F64_ne(v376, v383)|base.F64_eq(v371, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v394 = v376
		goto L91
	} else {
		goto L96
	}
L95:
	;
	v394 = v381
	goto L91
L96:
	;
	v392 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v394 = v392
	goto L91
L98:
	;
	if base.F64_eq(v423, v375)|base.F64_le(base.F64_abs(base.F64_sub(v375, v423)), float64(1e-06)) != 0 {
		v741 = v10
		goto L1
	} else {
		goto L105
	}
L99:
	;
	v410 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v412 = float64(0)
	if base.F64_eq(v374, v412)|base.F64_ne(v398, v412)|base.F64_eq(v394, v412) != 0 {
		v423 = v398
		goto L98
	} else {
		goto L103
	}
L102:
	;
	v423 = v410
	goto L98
L103:
	;
	v421 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v423 = v421
	goto L98
L105:
	;
	v430 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v431 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v432 = base.F64_mul(v430, v431)
	v434 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v432), v434)|base.F64_eq(base.F64_abs(v430), v434)|base.F64_eq(base.F64_abs(v431), v434) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v461 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v462 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v463 = base.F64_mul(v461, v462)
	v465 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v463), v465)|base.F64_eq(base.F64_abs(v461), v465)|base.F64_eq(base.F64_abs(v462), v465) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L107:
	;
	v447 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v449 = float64(0)
	if base.F64_eq(v430, v449)|base.F64_ne(v432, v449)|base.F64_eq(v431, v449) != 0 {
		v460 = v432
		goto L106
	} else {
		goto L111
	}
L110:
	;
	v460 = v447
	goto L106
L111:
	;
	v458 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	v460 = v458
	goto L106
L113:
	;
	v493 = math.Float64frombits(uint64(0x7ff0000000000000))
	v495 = base.F64_sub(v460, v491)
	if base.F64_eq(base.F64_abs(v460), v493)|base.F64_ne(base.F64_abs(v495), v493)|base.F64_eq(base.F64_abs(v491), v493) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v478 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L10
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v480 = float64(0)
	if base.F64_eq(v461, v480)|base.F64_ne(v463, v480)|base.F64_eq(v462, v480) != 0 {
		v491 = v463
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v491 = v478
	goto L113
L118:
	;
	v489 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	v491 = v489
	goto L113
L120:
	;
	v507 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L10
	} else {
		goto L123
	}
L121:
	;
	v509 = v495
	goto L122
L122:
	;
	v510 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v511 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v512 = base.F64_mul(v510, v511)
	v514 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v512), v514)|base.F64_eq(base.F64_abs(v510), v514)|base.F64_eq(base.F64_abs(v511), v514) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v509 = v507
	goto L122
L124:
	;
	v541 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v542 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v543 = base.F64_mul(v541, v542)
	v545 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v543), v545)|base.F64_eq(base.F64_abs(v541), v545)|base.F64_eq(base.F64_abs(v542), v545) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	v527 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L10
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v529 = float64(0)
	if base.F64_eq(v510, v529)|base.F64_ne(v512, v529)|base.F64_eq(v511, v529) != 0 {
		v540 = v512
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v540 = v527
	goto L124
L129:
	;
	v538 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
	;
	v540 = v538
	goto L124
L131:
	;
	v573 = math.Float64frombits(uint64(0x7ff0000000000000))
	v575 = base.F64_sub(v540, v571)
	if base.F64_eq(base.F64_abs(v540), v573)|base.F64_ne(base.F64_abs(v575), v573)|base.F64_eq(base.F64_abs(v571), v573) == int32(0) {
		goto L138
	} else {
		goto L139
	}
L132:
	;
	v558 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L10
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v560 = float64(0)
	if base.F64_eq(v541, v560)|base.F64_ne(v543, v560)|base.F64_eq(v542, v560) != 0 {
		v571 = v543
		goto L131
	} else {
		goto L136
	}
L135:
	;
	v571 = v558
	goto L131
L136:
	;
	v569 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	v571 = v569
	goto L131
L138:
	;
	v587 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L10
	} else {
		goto L141
	}
L139:
	;
	v589 = v575
	goto L140
L140:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v509)&int64(9223372036854775807)))|base.F64_ne(v589, float64(0)) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v589 = v587
	goto L140
L142:
	;
	v630 = math.Float64frombits(uint64(0x7ff0000000000000))
	v632 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v633 = base.F64_mul(v628, v632)
	if base.F64_eq(base.F64_abs(v628), v630)|base.F64_ne(base.F64_abs(v633), v630)|base.F64_eq(base.F64_abs(v632), v630) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L143:
	;
	v601 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L10
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v604 = math.Float64frombits(uint64(0x7ff0000000000000))
	v606 = base.F64_div(v509, v589)
	if base.F64_eq(base.F64_abs(v509), v604)|base.F64_ne(base.F64_abs(v606), v604) == int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v628 = v601
	goto L142
L147:
	;
	v614 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L10
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v616 = float64(0)
	if base.F64_eq(v509, v616)|base.F64_ne(v606, v616)|base.F64_eq(base.F64_abs(v589), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v628 = v606
		goto L142
	} else {
		goto L151
	}
L150:
	;
	v628 = v614
	goto L142
L151:
	;
	v626 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	v628 = v626
	goto L142
L153:
	;
	v660 = math.Float64frombits(uint64(0x7ff0000000000000))
	v662 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v663 = base.F64_add(v658, v662)
	if base.F64_eq(base.F64_abs(v658), v660)|base.F64_ne(base.F64_abs(v663), v660)|base.F64_eq(base.F64_abs(v662), v660) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v645 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L10
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v647 = float64(0)
	if base.F64_eq(v628, v647)|base.F64_ne(v633, v647)|base.F64_eq(v632, v647) != 0 {
		v658 = v633
		goto L153
	} else {
		goto L158
	}
L157:
	;
	v658 = v645
	goto L153
L158:
	;
	v656 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	v658 = v656
	goto L153
L160:
	;
	v675 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L10
	} else {
		goto L163
	}
L161:
	;
	v677 = v663
	goto L162
L162:
	;
	v683 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v677)&int64(9223372036854775807)))|base.F64_ne(v683, float64(0)) == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v677 = v675
	goto L162
L164:
	;
	v690 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L10
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v693 = math.Float64frombits(uint64(0x7ff0000000000000))
	v696 = base.F64_div(base.F64_neg(v677), v683)
	if base.F64_eq(base.F64_abs(v677), v693)|base.F64_ne(base.F64_abs(v696), v693) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v718 = v690
	v720 = v628
	goto L2
L168:
	;
	v704 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L10
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v706 = float64(0)
	if base.F64_eq(v677, v706)|base.F64_ne(v696, v706)|base.F64_eq(base.F64_abs(v683), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v718 = v696
		v720 = v628
		goto L2
	} else {
		goto L172
	}
L171:
	;
	v718 = v704
	v720 = v628
	goto L2
L172:
	;
	v716 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	v718 = v716
	v720 = v628
	goto L2
L174:
	;
	v726 = float64(0)
	if base.F64_eq(v718, v726) != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v729 = v726
	goto L177
L176:
	;
	v729 = v718
	goto L177
L177:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v729
	v731 = float64(0)
	if base.F64_eq(v720, v731) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v734 = v731
	goto L180
L179:
	;
	v734 = v720
	goto L180
L180:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v734
	v741 = v723
	goto L1
}
func F_line_intersect(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_line_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_line_perp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v33 float64
	_ = v33
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v50 float64
	_ = v50
	var v53 int32
	_ = v53
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v142 int32
	_ = v142
	v7 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = base.F64_abs(v11)
	if base.F64_le(v12, float64(1e-06)) != 0 {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
		return base.I64_extend_i32_u(base.F64_le(base.F64_abs(v15), float64(1e-06)))
	} else {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
		v22 = base.F64_abs(v21)
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v24 = base.F64_abs(v23)
		if base.F64_le(v24, float64(1e-06)) != 0 {
			return base.I64_extend_i32_u(base.F64_le(v22, float64(1e-06)))
		} else {
			if base.F64_le(v22, float64(1e-06)) != 0 {
				v142 = v7
				return base.I64_extend_i32_u(v142)
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				if base.F64_le(base.F64_abs(v33), float64(1e-06)) != 0 {
					v142 = v7
					return base.I64_extend_i32_u(v142)
				} else {
					v37 = math.Float64frombits(uint64(0x7ff0000000000000))
					v39 = base.F64_mul(v11, v23)
					if base.F64_eq(v12, v37)|base.F64_ne(base.F64_abs(v39), v37)|base.F64_eq(v24, v37) == int32(0) {
						v50 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v59 = v50
							v60 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v61 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v62 = base.F64_mul(v60, v61)
							v64 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v62), v64)|base.F64_eq(base.F64_abs(v60), v64)|base.F64_eq(base.F64_abs(v61), v64) == int32(0) {
								v77 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v90 = v77
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
										v102 = F_float_zero_divide_error_ext(m, int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int64(0)
										} else {
											v129 = v102
											v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
											return base.I64_extend_i32_u(v142)
										}
									} else {
										v105 = math.Float64frombits(uint64(0x7ff0000000000000))
										v107 = base.F64_div(v59, v90)
										if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
											v115 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												v129 = v115
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v117 = float64(0)
											if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v129 = v107
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											} else {
												v127 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int64(0)
												} else {
													v129 = v127
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											}
										}
									}
								}
							} else {
								v79 = float64(0)
								if base.F64_eq(v60, v79)|base.F64_ne(v62, v79)|base.F64_eq(v61, v79) != 0 {
									v90 = v62
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
										v102 = F_float_zero_divide_error_ext(m, int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int64(0)
										} else {
											v129 = v102
											v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
											return base.I64_extend_i32_u(v142)
										}
									} else {
										v105 = math.Float64frombits(uint64(0x7ff0000000000000))
										v107 = base.F64_div(v59, v90)
										if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
											v115 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												v129 = v115
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v117 = float64(0)
											if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v129 = v107
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											} else {
												v127 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int64(0)
												} else {
													v129 = v127
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											}
										}
									}
								} else {
									v88 = F_float_underflow_error_ext(m, int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int64(0)
									} else {
										v90 = v88
										if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
											v102 = F_float_zero_divide_error_ext(m, int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v129 = v102
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v105 = math.Float64frombits(uint64(0x7ff0000000000000))
											v107 = base.F64_div(v59, v90)
											if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
												v115 = F_float_overflow_error_ext(m, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int64(0)
												} else {
													v129 = v115
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											} else {
												v117 = float64(0)
												if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													v129 = v107
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												} else {
													v127 = F_float_underflow_error_ext(m, int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int64(0)
													} else {
														v129 = v127
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_ne(v39, float64(0)) != 0 {
							v59 = v39
							v60 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v61 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v62 = base.F64_mul(v60, v61)
							v64 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v62), v64)|base.F64_eq(base.F64_abs(v60), v64)|base.F64_eq(base.F64_abs(v61), v64) == int32(0) {
								v77 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v90 = v77
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
										v102 = F_float_zero_divide_error_ext(m, int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int64(0)
										} else {
											v129 = v102
											v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
											return base.I64_extend_i32_u(v142)
										}
									} else {
										v105 = math.Float64frombits(uint64(0x7ff0000000000000))
										v107 = base.F64_div(v59, v90)
										if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
											v115 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												v129 = v115
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v117 = float64(0)
											if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v129 = v107
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											} else {
												v127 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int64(0)
												} else {
													v129 = v127
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											}
										}
									}
								}
							} else {
								v79 = float64(0)
								if base.F64_eq(v60, v79)|base.F64_ne(v62, v79)|base.F64_eq(v61, v79) != 0 {
									v90 = v62
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
										v102 = F_float_zero_divide_error_ext(m, int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int64(0)
										} else {
											v129 = v102
											v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
											return base.I64_extend_i32_u(v142)
										}
									} else {
										v105 = math.Float64frombits(uint64(0x7ff0000000000000))
										v107 = base.F64_div(v59, v90)
										if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
											v115 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												v129 = v115
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v117 = float64(0)
											if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v129 = v107
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											} else {
												v127 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int64(0)
												} else {
													v129 = v127
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											}
										}
									}
								} else {
									v88 = F_float_underflow_error_ext(m, int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int64(0)
									} else {
										v90 = v88
										if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
											v102 = F_float_zero_divide_error_ext(m, int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v129 = v102
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v105 = math.Float64frombits(uint64(0x7ff0000000000000))
											v107 = base.F64_div(v59, v90)
											if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
												v115 = F_float_overflow_error_ext(m, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int64(0)
												} else {
													v129 = v115
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											} else {
												v117 = float64(0)
												if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													v129 = v107
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												} else {
													v127 = F_float_underflow_error_ext(m, int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int64(0)
													} else {
														v129 = v127
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v57 = F_float_underflow_error_ext(m, int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								v59 = v57
								v60 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v61 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								v62 = base.F64_mul(v60, v61)
								v64 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v62), v64)|base.F64_eq(base.F64_abs(v60), v64)|base.F64_eq(base.F64_abs(v61), v64) == int32(0) {
									v77 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int64(0)
									} else {
										v90 = v77
										if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
											v102 = F_float_zero_divide_error_ext(m, int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v129 = v102
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v105 = math.Float64frombits(uint64(0x7ff0000000000000))
											v107 = base.F64_div(v59, v90)
											if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
												v115 = F_float_overflow_error_ext(m, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int64(0)
												} else {
													v129 = v115
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											} else {
												v117 = float64(0)
												if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													v129 = v107
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												} else {
													v127 = F_float_underflow_error_ext(m, int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int64(0)
													} else {
														v129 = v127
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													}
												}
											}
										}
									}
								} else {
									v79 = float64(0)
									if base.F64_eq(v60, v79)|base.F64_ne(v62, v79)|base.F64_eq(v61, v79) != 0 {
										v90 = v62
										if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
											v102 = F_float_zero_divide_error_ext(m, int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v129 = v102
												v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
												return base.I64_extend_i32_u(v142)
											}
										} else {
											v105 = math.Float64frombits(uint64(0x7ff0000000000000))
											v107 = base.F64_div(v59, v90)
											if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
												v115 = F_float_overflow_error_ext(m, int32(0))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int64(0)
												} else {
													v129 = v115
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											} else {
												v117 = float64(0)
												if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													v129 = v107
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												} else {
													v127 = F_float_underflow_error_ext(m, int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int64(0)
													} else {
														v129 = v127
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													}
												}
											}
										}
									} else {
										v88 = F_float_underflow_error_ext(m, int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int64(0)
										} else {
											v90 = v88
											if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)))|base.F64_ne(v90, float64(0)) == int32(0) {
												v102 = F_float_zero_divide_error_ext(m, int32(0))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int64(0)
												} else {
													v129 = v102
													v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
													return base.I64_extend_i32_u(v142)
												}
											} else {
												v105 = math.Float64frombits(uint64(0x7ff0000000000000))
												v107 = base.F64_div(v59, v90)
												if base.F64_eq(base.F64_abs(v59), v105)|base.F64_ne(base.F64_abs(v107), v105) == int32(0) {
													v115 = F_float_overflow_error_ext(m, int32(0))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int64(0)
													} else {
														v129 = v115
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													}
												} else {
													v117 = float64(0)
													if base.F64_eq(v59, v117)|base.F64_ne(v107, v117)|base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														v129 = v107
														v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
														return base.I64_extend_i32_u(v142)
													} else {
														v127 = F_float_underflow_error_ext(m, int32(0))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int64(0)
														} else {
															v129 = v127
															v142 = base.F64_eq(v129, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v129, float64(1))), float64(1e-06))
															return base.I64_extend_i32_u(v142)
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
}
