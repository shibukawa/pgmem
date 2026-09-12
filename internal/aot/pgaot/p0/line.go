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
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v11 = base.F64_mul(v9, v10)
	v12 = base.F64_abs(v11)
	if base.F64_ne(v12, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return base.F64_le(v55, float64(1e-06))
L2:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L27
	} else {
		goto L29
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	if base.F64_ne(v11, float64(0)) != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v29 = base.F64_mul(v27, v28)
	v30 = base.F64_abs(v29)
	if base.F64_ne(v30, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.F64_eq(v9, float64(0)) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if base.F64_ne(v10, float64(0)) != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if base.F64_ne(v29, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v45 = base.F64_add(v11, v29)
	v46 = base.F64_abs(v45)
	if base.F64_ne(v46, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if base.F64_eq(v27, float64(0)) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(v28, float64(0)) != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = base.F64_abs(base.F64_add(v45, v53))
	if base.F64_ne(v55, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	if base.F64_eq(v12, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(v30, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if base.F64_eq(v46, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if base.F64_eq(base.F64_abs(v53), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	return int32(0)
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v48 int32
	_ = v48
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v127 float64
	_ = v127
	var v133 float64
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 float64
	_ = v176
	var v182 float64
	_ = v182
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 float64
	_ = v186
	var v189 int64
	_ = v189
	var v196 float64
	_ = v196
	var v203 int64
	_ = v203
	var v208 float64
	_ = v208
	var v224 int32
	_ = v224
	var v230 float64
	_ = v230
	var v234 int64
	_ = v234
	var v235 float64
	_ = v235
	var v238 int64
	_ = v238
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v283 float64
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
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
	return v338
L4:
	;
	v331 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v331)
	v338 = int32(0)
	goto L3
L5:
	;
	v300 = F_errsave_start(m, v18)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L83
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.Ui32(v38-int32(9)) < base.Ui32(int32(5)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v169 = F_path_decode(m, v28, int32(1), int32(2), v16+int32(16), v16+int32(15), int32(0), int32(373594), v19, v18)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L44
	}
L8:
	;
	goto L7
L9:
	;
	v28 = v28 + int32(1)
	goto L6
L10:
	;
	if v38 == int32(32) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v38 != int32(123) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v48 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v48
	v53 = F_float8in_internal(m, v48, v16+int32(16), int32(373594), v19, v18)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = v53
	if v18 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v64 = v62 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != int32(44) {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v58 != int32(447) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v61 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v72 = F_float8in_internal(m, v64, v16+int32(16), int32(373594), v19, v18)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+8)) = v72
	if v18 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v83 = v81 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 != int32(44) {
		goto L5
	} else {
		goto L24
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v77 != int32(447) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v91 = F_float8in_internal(m, v83, v16+int32(16), int32(373594), v19, v18)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+16)) = v91
	if v18 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v102 = v100 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v104 != int32(125) {
		goto L5
	} else {
		goto L30
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v96 != int32(447) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v99 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v110 = v102
	goto L31
L31:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if base.Ui32(v120-int32(9)) < base.Ui32(int32(5)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v157 = v110 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v157
	v110 = v157
	goto L31
L34:
	;
	if v120 == int32(32) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if v120 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	if base.F64_le(base.F64_abs(v127), float64(1e-06)) == int32(0) {
		v338 = v21
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	if base.F64_le(base.F64_abs(v133), float64(1e-06)) == int32(0) {
		v338 = v21
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v139 = int32(0)
	v140 = F_errsave_start(m, v18)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v140 == int32(0) {
		v338 = v139
		goto L3
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(240165), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v18, int32(493629), int32(998), int32(279616))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v338 = v139
	goto L3
L44:
	;
	if v169 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v173)
	v338 = int32(0)
	goto L3
L46:
	;
	goto L47
L47:
	;
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.Ui64(base.I64_reinterpret_f64(v176)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v278 = v16 + int32(16)
	v283 = F_point_sl(m, v278, v16+int32(32))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L81
	}
L49:
	;
	v253 = int32(0)
	v254 = F_errsave_start(m, v18)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L76
	}
L50:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	v238 = base.I64_reinterpret_f64(v235) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v234) {
		goto L70
	} else {
		goto L71
	}
L51:
	;
	if base.F64_ne(v176, v182) != 0 {
		goto L48
	} else {
		goto L68
	}
L52:
	;
	if base.F64_ne(v176, v182) != 0 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v184 = int64(9223372036854775807)
	v185 = base.I64_reinterpret_f64(v182) & v184
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v189 = base.I64_reinterpret_f64(v186) & v184
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v189) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.Ui64(v203&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L48
	} else {
		goto L61
	}
L56:
	;
	v224 = base.B2i32(base.Ui64(v185) < base.Ui64(int64(9218868437227405313)))
	goto L51
L57:
	;
	goto L58
L58:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v185) {
		goto L48
	} else {
		goto L59
	}
L59:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
	if base.Ui64(base.I64_reinterpret_f64(v196)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v224 = int32(1)
	goto L51
L61:
	;
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v230 = v208
	v234 = base.I64_reinterpret_f64(v208) & int64(9223372036854775807)
	goto L50
L62:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v176, v182)), float64(1e-06)) == int32(0) {
		goto L48
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.F64_eq(v186, v196) != 0 {
		goto L49
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v186, v196)), float64(1e-06)) != 0 {
		goto L49
	} else {
		goto L67
	}
L67:
	;
	goto L48
L68:
	;
	if v224 == int32(0) {
		goto L48
	} else {
		goto L69
	}
L69:
	;
	v230 = v186
	v234 = v189
	goto L50
L70:
	;
	if base.Ui64(v238) <= base.Ui64(int64(9218868437227405312)) {
		goto L48
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.F64_ne(v235, v230) != 0 {
		goto L48
	} else {
		goto L74
	}
L73:
	;
	goto L49
L74:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v238) {
		goto L48
	} else {
		goto L75
	}
L75:
	;
	goto L49
L76:
	;
	if v254 == int32(0) {
		v338 = v253
		goto L3
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(119191), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errsave_finish(m, v18, int32(493629), int32(1008), int32(279616))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v338 = v253
	goto L3
L81:
	;
	F_line_construct(m, v21, v278, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v338 = v21
	goto L3
L83:
	;
	if v300 == int32(0) {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(373594)
	F_errmsg(m, int32(724862), v16)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errsave_finish(m, v18, int32(493629), int32(975), int32(413638))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
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
