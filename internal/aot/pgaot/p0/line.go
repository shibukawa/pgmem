package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_contain_point(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v25 float64
	_ = v25
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v51 float64
	_ = v51
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_mul(v9, v10)
	v12 = base.F64_abs(v11)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(v12, v13)|base.F64_eq(base.F64_abs(v9), v13) == int32(0))&base.F64_ne(base.F64_abs(v10), v13) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v25 = float64(0)
		if base.B2i32(base.F64_eq(v9, v25)|base.F64_ne(v11, v25) == int32(0))&base.F64_ne(v10, v25) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v37 = base.F64_mul(v35, v36)
			v38 = base.F64_abs(v37)
			v39 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(v38, v39)|base.F64_eq(base.F64_abs(v35), v39) == int32(0))&base.F64_ne(base.F64_abs(v36), v39) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v51 = float64(0)
				if base.B2i32(base.F64_eq(v35, v51)|base.F64_ne(v37, v51) == int32(0))&base.F64_ne(v36, v51) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v61 = math.Float64frombits(uint64(0x7ff0000000000000))
					v63 = base.F64_add(v11, v37)
					v64 = base.F64_abs(v63)
					if base.B2i32(base.F64_eq(v12, v61)|base.F64_ne(v64, v61) == int32(0))&base.F64_ne(v38, v61) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v73 = math.Float64frombits(uint64(0x7ff0000000000000))
						v75 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						v77 = base.F64_abs(base.F64_add(v63, v75))
						if base.F64_eq(v64, v73)|base.F64_ne(v77, v73)|base.F64_eq(base.F64_abs(v75), v73) != 0 {
							return base.F64_le(v77, float64(1e-06))
						} else {
							F_float_overflow_error(m)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
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
	}
}
func F_line_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 float64
	_ = v178
	var v184 float64
	_ = v184
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 float64
	_ = v188
	var v191 int64
	_ = v191
	var v198 float64
	_ = v198
	var v205 int64
	_ = v205
	var v210 float64
	_ = v210
	var v228 int32
	_ = v228
	var v235 float64
	_ = v235
	var v239 int64
	_ = v239
	var v240 float64
	_ = v240
	var v243 int64
	_ = v243
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v287 float64
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_palloc(m, int32(24))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = v19
	goto L6
L3:
	;
	m.G0 = v16 + int32(48)
	return v342
L4:
	;
	v335 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v335)
	v342 = int32(0)
	goto L3
L5:
	;
	v304 = F_errsave_start(m, v18)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L78
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(base.Ui32(v38-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v38 == int32(32)) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v171 = F_path_decode(m, v28, int32(1), int32(2), v16+int32(16), v16+int32(15), int32(0), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	v28 = v28 + int32(1)
	goto L6
L9:
	;
	if v38 == int32(123) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v51 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v51
	v56 = F_float8in_internal(m, v51, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = v56
	if v18 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v67 = v65 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 != int32(44) {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v61 != int32(447) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v64 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v75 = F_float8in_internal(m, v67, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+8)) = v75
	if v18 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v86 = v84 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != int32(44) {
		goto L5
	} else {
		goto L25
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v80 != int32(447) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v83 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v94 = F_float8in_internal(m, v86, v16+int32(16), int32(_a_F_line_in_0), v19, v18)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+16)) = v94
	if v18 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v105 = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v107 != int32(125) {
		goto L5
	} else {
		goto L31
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v99 != int32(447) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v102 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v113 = v105
	goto L32
L32:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if base.B2i32(base.Ui32(v123-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v123 == int32(32)) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v132 = v113 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v132
	v113 = v132
	goto L32
L35:
	;
	if v123 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	if base.F64_le(base.F64_abs(v134), float64(1e-06)) == int32(0) {
		v342 = v21
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	if base.F64_le(base.F64_abs(v140), float64(1e-06)) == int32(0) {
		v342 = v21
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v146 = int32(0)
	v147 = F_errsave_start(m, v18)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v147 == int32(0) {
		v342 = v146
		goto L3
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_line_in_4), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(998), int32(_a_F_line_in_5))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v342 = v146
	goto L3
L45:
	;
	if v171 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
	v342 = int32(0)
	goto L3
L47:
	;
	goto L48
L48:
	;
	v178 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.Ui64(base.I64_reinterpret_f64(v178)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v284 = v16 + int32(16)
	v287 = F_point_sl(m, v284, v16+int32(32))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L76
	}
L50:
	;
	v259 = int32(0)
	v260 = F_errsave_start(m, v18)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L71
	}
L51:
	;
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v243 = base.I64_reinterpret_f64(v240) & int64(9223372036854775807)
	if base.Ui64(v239) <= base.Ui64(int64(9218868437227405312)) {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	if base.B2i32(v228 == int32(0))|base.F64_ne(v178, v184) != 0 {
		goto L49
	} else {
		goto L65
	}
L53:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v178, v184)), float64(1e-06)) == int32(0))&base.F64_ne(v178, v184) != 0 {
		goto L49
	} else {
		goto L63
	}
L54:
	;
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v186 = int64(9223372036854775807)
	v187 = base.I64_reinterpret_f64(v184) & v186
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v191 = base.I64_reinterpret_f64(v188) & v186
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v191) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.Ui64(v205&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L49
	} else {
		goto L62
	}
L57:
	;
	v228 = base.B2i32(base.Ui64(v187) < base.Ui64(int64(9218868437227405313)))
	goto L52
L58:
	;
	goto L59
L59:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v187) {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v198)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v228 = int32(1)
	goto L52
L62:
	;
	v210 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v235 = v210
	v239 = base.I64_reinterpret_f64(v210) & int64(9223372036854775807)
	goto L51
L63:
	;
	if base.F64_eq(v188, v198)|base.F64_le(base.F64_abs(base.F64_sub(v188, v198)), float64(1e-06)) != 0 {
		goto L50
	} else {
		goto L64
	}
L64:
	;
	goto L49
L65:
	;
	v235 = v188
	v239 = v191
	goto L51
L66:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v243))|base.F64_ne(v240, v235) != 0 {
		goto L49
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui64(v243) < base.Ui64(int64(9218868437227405313)) {
		goto L49
	} else {
		goto L70
	}
L69:
	;
	goto L50
L70:
	;
	goto L50
L71:
	;
	if v260 == int32(0) {
		v342 = v259
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_line_in_6), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(1008), int32(_a_F_line_in_5))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v342 = v259
	goto L3
L76:
	;
	F_line_construct(m, v21, v284, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v342 = v21
	goto L3
L78:
	;
	if v304 == int32(0) {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_line_in_0)
	F_errmsg(m, int32(_a_F_line_in_1), v16)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errsave_finish(m, v18, int32(_a_F_line_in_2), int32(975), int32(_a_F_line_in_3))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L4
}
func F_line_parallel(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_line_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5 ^ int32(1)
	}
}
