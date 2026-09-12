package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_lseg_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_lseg_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_palloc(m, int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v24 = F_path_decode(m, v12, int32(1), int32(2), v16, v9+int32(15), int32(0), int32(335742), v12, v11)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
				v30 = int32(0)
			} else {
				v30 = v16
			}
			m.G0 = v9 + int32(16)
			return v30
		}
	}
}
func F_lseg_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v22 int32
	_ = v22
	var v30 float64
	_ = v30
	var v38 float64
	_ = v38
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v58 float64
	_ = v58
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 float64
	_ = v105
	var v111 float64
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 float64
	_ = v115
	var v118 int64
	_ = v118
	var v125 float64
	_ = v125
	var v132 float64
	_ = v132
	var v138 float64
	_ = v138
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v171 int64
	_ = v171
	var v172 float64
	_ = v172
	var v175 int64
	_ = v175
	var v184 float64
	_ = v184
	var v190 float64
	_ = v190
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v199 float64
	_ = v199
	var v202 int64
	_ = v202
	var v209 float64
	_ = v209
	var v221 float64
	_ = v221
	var v239 int32
	_ = v239
	var v245 float64
	_ = v245
	var v248 int64
	_ = v248
	var v249 float64
	_ = v249
	var v252 int64
	_ = v252
	var v274 int32
	_ = v274
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v288 int32
	_ = v288
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = l1 + int32(16)
	v19 = F_point_sl(m, l1, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return v288
L2:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v277
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v279
	v288 = v102
	goto L1
L3:
	;
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	if base.Ui64(base.I64_reinterpret_f64(v190)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L66
	} else {
		goto L67
	}
L4:
	;
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
	v175 = base.I64_reinterpret_f64(v172) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v171) {
		goto L56
	} else {
		goto L57
	}
L5:
	;
	if base.F64_ne(v111, v105) != 0 {
		v184 = v111
		goto L3
	} else {
		goto L54
	}
L6:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L53
	}
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L52
	}
L8:
	;
	v80 = F_line_interpt_line(m, v15+int32(32), v15+int32(8), l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L24
	}
L9:
	;
	return int32(0)
L10:
	;
	if base.F64_eq(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(-4616189618054758400)
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v30
	goto L8
L12:
	;
	goto L13
L13:
	;
	if base.F64_eq(v19, float64(0)) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(0)
	v38 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v38
	goto L8
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(-4616189618054758400)
	*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v19
	v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v45 = base.F64_mul(v19, v44)
	v46 = base.F64_abs(v45)
	v47 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v46, v47)&base.F64_ne(base.F64_abs(v44), v47) != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v53 = float64(0)
	if base.F64_eq(v45, v53)&base.F64_ne(v44, v53) != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v58 = base.F64_sub(v43, v45)
	if base.F64_ne(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v58
	if base.F64_ne(v58, float64(0)) != 0 {
		goto L8
	} else {
		goto L23
	}
L20:
	;
	if base.F64_eq(base.F64_abs(v43), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if base.F64_ne(v46, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = int64(0)
	goto L8
L24:
	;
	if v80 == int32(0) {
		v288 = v8
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v86 = F_point_dt(m, v15+int32(32), l1)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v90 = F_point_dt(m, v15+int32(32), v18)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v92 = base.F64_add(v86, v90)
	v93 = F_point_dt(m, l1, v18)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	if base.F64_ne(v92, v93) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v92, v93)), float64(1e-06)) == int32(0) {
		v288 = v8
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v102 = int32(1)
	if l0 == int32(0) {
		v288 = v102
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v105)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if base.F64_ne(v111, v105) != 0 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
	v113 = int64(9223372036854775807)
	v114 = base.I64_reinterpret_f64(v111) & v113
	v115 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v118 = base.I64_reinterpret_f64(v115) & v113
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v118) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
	if base.Ui64(base.I64_reinterpret_f64(v132)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v184 = v132
		goto L3
	} else {
		goto L43
	}
L38:
	;
	v161 = base.B2i32(base.Ui64(v114) < base.Ui64(int64(9218868437227405313)))
	goto L5
L39:
	;
	goto L40
L40:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v114) {
		v184 = v111
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
	if base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v161 = int32(1)
	goto L5
L43:
	;
	v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v167 = v132
	v168 = v138
	v171 = base.I64_reinterpret_f64(v138) & int64(9223372036854775807)
	goto L4
L44:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v105, v111)), float64(1e-06)) == int32(0) {
		v184 = v111
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.F64_eq(v115, v125) != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v274 = l1
	goto L2
L49:
	;
	goto L50
L50:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v115, v125)), float64(1e-06)) == int32(0) {
		v184 = v111
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v274 = l1
	goto L2
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	if v161 == int32(0) {
		v184 = v111
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v167 = v111
	v168 = v115
	v171 = v118
	goto L4
L56:
	;
	if base.Ui64(v175) <= base.Ui64(int64(9218868437227405312)) {
		v184 = v167
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if base.F64_ne(v172, v168) != 0 {
		v184 = v167
		goto L3
	} else {
		goto L60
	}
L59:
	;
	v274 = l1
	goto L2
L60:
	;
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v175) {
		v184 = v167
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v274 = l1
	goto L2
L62:
	;
	v274 = v15 + int32(32)
	goto L2
L63:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
	v252 = base.I64_reinterpret_f64(v249) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v248) {
		goto L83
	} else {
		goto L84
	}
L64:
	;
	if base.F64_ne(v184, v190) != 0 {
		goto L62
	} else {
		goto L81
	}
L65:
	;
	if base.F64_ne(v184, v190) != 0 {
		goto L75
	} else {
		goto L76
	}
L66:
	;
	v197 = int64(9223372036854775807)
	v198 = base.I64_reinterpret_f64(v184) & v197
	v199 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v202 = base.I64_reinterpret_f64(v199) & v197
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v202) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui64(base.I64_reinterpret_f64(v184)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L62
	} else {
		goto L74
	}
L69:
	;
	v239 = base.B2i32(base.Ui64(v198) < base.Ui64(int64(9218868437227405313)))
	goto L64
L70:
	;
	goto L71
L71:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v198) {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v209 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
	if base.Ui64(base.I64_reinterpret_f64(v209)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	v239 = int32(1)
	goto L64
L74:
	;
	v221 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v245 = v221
	v248 = base.I64_reinterpret_f64(v221) & int64(9223372036854775807)
	goto L63
L75:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v190, v184)), float64(1e-06)) == int32(0) {
		goto L62
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if base.F64_eq(v199, v209) != 0 {
		v274 = v18
		goto L2
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v199, v209)), float64(1e-06)) == int32(0) {
		goto L62
	} else {
		goto L80
	}
L80:
	;
	v274 = v18
	goto L2
L81:
	;
	if v239 == int32(0) {
		goto L62
	} else {
		goto L82
	}
L82:
	;
	v245 = v199
	v248 = v202
	goto L63
L83:
	;
	if base.Ui64(v252) <= base.Ui64(int64(9218868437227405312)) {
		goto L62
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if base.F64_ne(v249, v245) != 0 {
		goto L62
	} else {
		goto L87
	}
L86:
	;
	v274 = v18
	goto L2
L87:
	;
	if base.Ui64(v252) < base.Ui64(int64(9218868437227405313)) {
		v274 = v18
		goto L2
	} else {
		goto L88
	}
L88:
	;
	goto L62
}
func F_lseg_perp(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_point_invsl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return base.F64_eq(v9, v15) | base.F64_le(base.F64_abs(base.F64_sub(v9, v15)), float64(1e-06))
		}
	}
}
