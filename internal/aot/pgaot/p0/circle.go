package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_below(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		v26 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = v26
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_sub(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
			}
		}
	}
}
func F_circle_box(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v129 float64
	_ = v129
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v17 = base.F64_div(v15, float64(1.4142135623730951))
	v19 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v37 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v27 = F_float_overflow_error_ext(m, v14)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = float64(0)
	if base.F64_eq(v15, v29)|base.F64_ne(v17, v29) != 0 {
		v36 = v17
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v36 = v27
	goto L3
L8:
	;
	v34 = F_float_underflow_error_ext(m, v14)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v36 = v34
	goto L3
L10:
	;
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v146)
	return int64(0)
L11:
	;
	v45 = math.Float64frombits(uint64(0x7ff0000000000000))
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v48 = base.F64_add(v36, v47)
	if base.F64_eq(base.F64_abs(v36), v45)|base.F64_ne(base.F64_abs(v48), v45)|base.F64_eq(base.F64_abs(v47), v45) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v40 != int32(453) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
	if v43 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v59 = v48
	goto L17
L16:
	;
	v57 = F_float_overflow_error_ext(m, v37)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v59 = v57
	goto L17
L19:
	;
	v69 = math.Float64frombits(uint64(0x7ff0000000000000))
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v72 = base.F64_sub(v71, v36)
	if base.F64_eq(base.F64_abs(v36), v69)|base.F64_ne(base.F64_abs(v72), v69)|base.F64_eq(base.F64_abs(v71), v69) != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v64 != int32(453) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
	if v67 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v83 = v72
	goto L25
L24:
	;
	v81 = F_float_overflow_error_ext(m, v61)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v83 = v81
	goto L25
L27:
	;
	v93 = math.Float64frombits(uint64(0x7ff0000000000000))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v96 = base.F64_add(v36, v95)
	if base.F64_eq(base.F64_abs(v36), v93)|base.F64_ne(base.F64_abs(v96), v93)|base.F64_eq(base.F64_abs(v95), v93) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v88 != int32(453) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v91 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v107 = v96
	goto L33
L32:
	;
	v105 = F_float_overflow_error_ext(m, v85)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v109 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v107 = v105
	goto L33
L35:
	;
	v117 = math.Float64frombits(uint64(0x7ff0000000000000))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v120 = base.F64_sub(v119, v36)
	if base.F64_eq(base.F64_abs(v36), v117)|base.F64_ne(base.F64_abs(v120), v117)|base.F64_eq(base.F64_abs(v119), v117) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v112 != int32(453) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+4)))
	if v115 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v131 = v120
	goto L41
L40:
	;
	v129 = F_float_overflow_error_ext(m, v109)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v133 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v131 = v129
	goto L41
L43:
	;
	return base.I64_extend_i32_u(v10)
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v136 != int32(453) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)))
	if v139 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	goto L43
}
func F_circle_contain_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_sub(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12)|base.F64_eq(base.F64_abs(v9), v12) != 0 {
		v27 = v10
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
			v45 = v30
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			v58 = base.F64_abs(v27)
			v59 = base.F64_abs(v45)
			v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
			if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
				v63 = v58
			} else {
				v63 = v59
			}
			v64 = base.I64_reinterpret_f64(v63)
			v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
			if v66 == int64(2047) {
				v121 = v63
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v69 = v59
				} else {
					v69 = v58
				}
				if v64 == int64(0) {
					v121 = v69
				} else {
					v72 = base.I64_reinterpret_f64(v69)
					v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
					if v74 == int64(2047) {
						v121 = v69
					} else {
						if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
							v121 = base.F64_add(v58, v59)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
								v85 = float64(1.90109156629516e-211)
								v98 = base.F64_mul(v69, v85)
								v99 = base.F64_mul(v63, v85)
								v100 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
									v98 = v69
									v99 = v63
									v100 = float64(1)
								} else {
									v93 = float64(5.260135901548374e+210)
									v98 = base.F64_mul(v69, v93)
									v99 = base.F64_mul(v63, v93)
									v100 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v56+int32(24), v56+int32(16), v98)
							mBase = m.M
							F_sq(m, v56+int32(8), v56, v99)
							mBase = m.M
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
							v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
						}
					}
				}
			}
			m.G0 = v56 + int32(32)
			v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
			return base.I64_extend_i32_u(base.F64_le(v121, v127))
		} else {
			v43 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = v43
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				return base.I64_extend_i32_u(base.F64_le(v121, v127))
			}
		}
	} else {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
				v45 = v30
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				return base.I64_extend_i32_u(base.F64_le(v121, v127))
			} else {
				v43 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = v43
					v54 = m.G0
					v56 = v54 - int32(32)
					m.G0 = v56
					v58 = base.F64_abs(v27)
					v59 = base.F64_abs(v45)
					v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v63 = v58
					} else {
						v63 = v59
					}
					v64 = base.I64_reinterpret_f64(v63)
					v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
					if v66 == int64(2047) {
						v121 = v63
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
							v69 = v59
						} else {
							v69 = v58
						}
						if v64 == int64(0) {
							v121 = v69
						} else {
							v72 = base.I64_reinterpret_f64(v69)
							v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
							if v74 == int64(2047) {
								v121 = v69
							} else {
								if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
									v121 = base.F64_add(v58, v59)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
										v85 = float64(1.90109156629516e-211)
										v98 = base.F64_mul(v69, v85)
										v99 = base.F64_mul(v63, v85)
										v100 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
											v98 = v69
											v99 = v63
											v100 = float64(1)
										} else {
											v93 = float64(5.260135901548374e+210)
											v98 = base.F64_mul(v69, v93)
											v99 = base.F64_mul(v63, v93)
											v100 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v56+int32(24), v56+int32(16), v98)
									mBase = m.M
									F_sq(m, v56+int32(8), v56, v99)
									mBase = m.M
									v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
									v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
								}
							}
						}
					}
					m.G0 = v56 + int32(32)
					v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
					return base.I64_extend_i32_u(base.F64_le(v121, v127))
				}
			}
		}
	}
}
func F_circle_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = base.F64_mul(v8, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = math.Float64frombits(uint64(0x7ff0000000000000))
	v37 = base.F64_mul(v32, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v32), v34)|base.F64_ne(base.F64_abs(v37), v34) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v20 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v8, v24)|base.F64_ne(v9, v24) != 0 {
		v32 = v9
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return int64(0)
L6:
	;
	v32 = v20
	goto L1
L7:
	;
	v30 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = v30
	goto L1
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v57 = base.F64_mul(v56, v56)
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v45 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = float64(0)
	if base.F64_eq(v32, v47)|base.F64_ne(v37, v47) != 0 {
		v55 = v37
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v55 = v45
	goto L9
L14:
	;
	v53 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = v53
	goto L9
L16:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	v83 = base.F64_mul(v78, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v78), v80)|base.F64_ne(base.F64_abs(v83), v80) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v68 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v70 = float64(0)
	if base.F64_eq(v56, v70)|base.F64_ne(v57, v70) != 0 {
		v78 = v57
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v78 = v68
	goto L16
L21:
	;
	v76 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v78 = v76
	goto L16
L23:
	;
	return base.I64_extend_i32_u(base.F64_gt(v55, base.F64_add(v101, float64(1e-06))))
L24:
	;
	v91 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v93 = float64(0)
	if base.F64_eq(v78, v93)|base.F64_ne(v83, v93) != 0 {
		v101 = v83
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v101 = v91
	goto L23
L28:
	;
	v99 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v101 = v99
	goto L23
}
func F_circle_left(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		v26 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = v26
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_sub(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_gt(v48, base.F64_add(v30, float64(1e-06))))
			}
		}
	}
}
func F_circle_overabove(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		v26 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = v26
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_sub(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
			}
		}
	}
}
func F_circle_overbelow(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		v26 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = v26
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v33 = base.F64_add(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_add(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
			}
		}
	}
}
func F_circle_overlap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 float64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v111 float64
	_ = v111
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v147 float64
	_ = v147
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = base.F64_sub(v9, v11)
	v14 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v12), v14)|base.F64_eq(base.F64_abs(v9), v14)|base.F64_eq(base.F64_abs(v11), v14) == int32(0) {
		v27 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			v31 = v27
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v34 = base.F64_sub(v32, v33)
			v36 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36)|base.F64_eq(base.F64_abs(v33), v36) == int32(0) {
				v49 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					v51 = v49
					v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
					v54 = base.F64_add(v52, v53)
					v56 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
						v69 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							v71 = v69
							v80 = m.G0
							v82 = v80 - int32(32)
							m.G0 = v82
							v84 = base.F64_abs(v31)
							v85 = base.F64_abs(v51)
							v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v89 = v84
							} else {
								v89 = v85
							}
							v90 = base.I64_reinterpret_f64(v89)
							v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
							if v92 == int64(2047) {
								v147 = v89
							} else {
								if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
									v95 = v85
								} else {
									v95 = v84
								}
								if v90 == int64(0) {
									v147 = v95
								} else {
									v98 = base.I64_reinterpret_f64(v95)
									v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
									if v100 == int64(2047) {
										v147 = v95
									} else {
										if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
											v147 = base.F64_add(v84, v85)
										} else {
											if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
												v111 = float64(1.90109156629516e-211)
												v124 = base.F64_mul(v95, v111)
												v125 = base.F64_mul(v89, v111)
												v126 = float64(5.260135901548374e+210)
											} else {
												if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
													v124 = v95
													v125 = v89
													v126 = float64(1)
												} else {
													v119 = float64(5.260135901548374e+210)
													v124 = base.F64_mul(v95, v119)
													v125 = base.F64_mul(v89, v119)
													v126 = float64(1.90109156629516e-211)
												}
											}
											F_sq(m, v82+int32(24), v82+int32(16), v124)
											mBase = m.M
											F_sq(m, v82+int32(8), v82, v125)
											mBase = m.M
											v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
											v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
											v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
											v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
										}
									}
								}
							}
							m.G0 = v82 + int32(32)
							return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
						}
					} else {
						v71 = v54
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				}
			} else {
				v51 = v34
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v54 = base.F64_add(v52, v53)
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
					v69 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v71 = v69
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				} else {
					v71 = v54
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			}
		}
	} else {
		v31 = v12
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
		v34 = base.F64_sub(v32, v33)
		v36 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36)|base.F64_eq(base.F64_abs(v33), v36) == int32(0) {
			v49 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				v51 = v49
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v54 = base.F64_add(v52, v53)
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
					v69 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v71 = v69
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				} else {
					v71 = v54
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			}
		} else {
			v51 = v34
			v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
			v54 = base.F64_add(v52, v53)
			v56 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
				v69 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					v71 = v69
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			} else {
				v71 = v54
				v80 = m.G0
				v82 = v80 - int32(32)
				m.G0 = v82
				v84 = base.F64_abs(v31)
				v85 = base.F64_abs(v51)
				v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
				if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
					v89 = v84
				} else {
					v89 = v85
				}
				v90 = base.I64_reinterpret_f64(v89)
				v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
				if v92 == int64(2047) {
					v147 = v89
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v95 = v85
					} else {
						v95 = v84
					}
					if v90 == int64(0) {
						v147 = v95
					} else {
						v98 = base.I64_reinterpret_f64(v95)
						v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
						if v100 == int64(2047) {
							v147 = v95
						} else {
							if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
								v147 = base.F64_add(v84, v85)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
									v111 = float64(1.90109156629516e-211)
									v124 = base.F64_mul(v95, v111)
									v125 = base.F64_mul(v89, v111)
									v126 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
										v124 = v95
										v125 = v89
										v126 = float64(1)
									} else {
										v119 = float64(5.260135901548374e+210)
										v124 = base.F64_mul(v95, v119)
										v125 = base.F64_mul(v89, v119)
										v126 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v82+int32(24), v82+int32(16), v124)
								mBase = m.M
								F_sq(m, v82+int32(8), v82, v125)
								mBase = m.M
								v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
								v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
								v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
								v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
							}
						}
					}
				}
				m.G0 = v82 + int32(32)
				return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
			}
		}
	}
}
func F_circle_sub_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19)|base.F64_eq(base.F64_abs(v16), v19) == int32(0) {
			v32 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = v32
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v37 = base.F64_sub(v35, v36)
				v39 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
					v52 = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					return base.I64_extend_i32_u(v11)
				} else {
					v50 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						v52 = v50
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
						v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
						return base.I64_extend_i32_u(v11)
					}
				}
			}
		} else {
			v34 = v17
			v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v37 = base.F64_sub(v35, v36)
			v39 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
				v52 = v37
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
				v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
				return base.I64_extend_i32_u(v11)
			} else {
				v50 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = v50
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					return base.I64_extend_i32_u(v11)
				}
			}
		}
	}
}
