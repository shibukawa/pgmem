package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_poly_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v70 float64
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 float64
	_ = v74
	var v77 int64
	_ = v77
	var v84 float64
	_ = v84
	var v91 int64
	_ = v91
	var v96 float64
	_ = v96
	var v115 int32
	_ = v115
	var v120 float64
	_ = v120
	var v125 int64
	_ = v125
	var v126 float64
	_ = v126
	var v129 int64
	_ = v129
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 float64
	_ = v179
	var v185 float64
	_ = v185
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 float64
	_ = v189
	var v192 int64
	_ = v192
	var v199 float64
	_ = v199
	var v206 int64
	_ = v206
	var v211 float64
	_ = v211
	var v230 int32
	_ = v230
	var v235 float64
	_ = v235
	var v240 int64
	_ = v240
	var v241 float64
	_ = v241
	var v244 int64
	_ = v244
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 float64
	_ = v302
	var v308 float64
	_ = v308
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v312 float64
	_ = v312
	var v315 int64
	_ = v315
	var v322 float64
	_ = v322
	var v329 int64
	_ = v329
	var v334 float64
	_ = v334
	var v353 int32
	_ = v353
	var v358 float64
	_ = v358
	var v363 int64
	_ = v363
	var v364 float64
	_ = v364
	var v367 int64
	_ = v367
	var v384 int32
	_ = v384
	var v414 int32
	_ = v414
	var v453 int64
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v31 = int32(0)
	if base.B2i32(v28 != v29)|base.B2i32(v28 <= v31) == v31 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v455 != v21 {
		goto L92
	} else {
		goto L93
	}
L5:
	;
	v36 = int32(40)
	v37 = v26 + v36
	v39 = v21 + v36
	v55 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v453 = int64(0)
	goto L4
L8:
	;
	v63 = v37 + v55<<(uint(int32(4))%32)
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	goto L7
L10:
	;
	v414 = v55 + int32(1)
	if v414 != v28 {
		v55 = v414
		goto L8
	} else {
		goto L91
	}
L11:
	;
	v145 = int32(1)
	v146 = int64(1)
	if v28 == v145 {
		v453 = v146
		goto L4
	} else {
		goto L32
	}
L12:
	;
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	v129 = base.I64_reinterpret_f64(v126) & int64(9223372036854775807)
	if base.Ui64(v125) <= base.Ui64(int64(9218868437227405312)) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	if base.B2i32(v115 == int32(0))|base.F64_ne(v64, v70) != 0 {
		goto L10
	} else {
		goto L26
	}
L14:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v64, v70)), float64(1e-06)) == int32(0))&base.F64_ne(v64, v70) != 0 {
		goto L10
	} else {
		goto L24
	}
L15:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v39)))
	v72 = int64(9223372036854775807)
	v73 = base.I64_reinterpret_f64(v70) & v72
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	v77 = base.I64_reinterpret_f64(v74) & v72
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v77) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	if base.Ui64(v91&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L10
	} else {
		goto L23
	}
L18:
	;
	v115 = base.B2i32(base.Ui64(v73) < base.Ui64(int64(9218868437227405313)))
	goto L13
L19:
	;
	goto L20
L20:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v73) {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v21)+48))
	if base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v115 = int32(1)
	goto L13
L23:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	v120 = v96
	v125 = base.I64_reinterpret_f64(v96) & int64(9223372036854775807)
	goto L12
L24:
	;
	if base.F64_eq(v74, v84)|base.F64_le(base.F64_abs(base.F64_sub(v74, v84)), float64(1e-06)) != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	v120 = v74
	v125 = v77
	goto L12
L27:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v129))|base.F64_ne(v126, v120) != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui64(v129) < base.Ui64(int64(9218868437227405313)) {
		goto L10
	} else {
		goto L31
	}
L30:
	;
	goto L11
L31:
	;
	goto L11
L32:
	;
	v154 = v55
	v155 = v145
	goto L33
L33:
	;
	v170 = v39 + v155<<(uint(int32(4))%32)
	v172 = v154 + int32(1)
	if v172 < v28 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	if v155 == v28 {
		v453 = v146
		goto L4
	} else {
		goto L61
	}
L35:
	;
	goto L34
L36:
	;
	v261 = v155 + int32(1)
	if v261 != v28 {
		v154 = v175
		v155 = v261
		goto L33
	} else {
		goto L60
	}
L37:
	;
	v241 = *(*float64)(unsafe.Add(mBase, uint32(v170)+8))
	v244 = base.I64_reinterpret_f64(v241) & int64(9223372036854775807)
	if base.Ui64(v240) <= base.Ui64(int64(9218868437227405312)) {
		goto L55
	} else {
		goto L56
	}
L38:
	;
	if base.B2i32(v230 == int32(0))|base.F64_ne(v179, v185) != 0 {
		goto L35
	} else {
		goto L54
	}
L39:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v179, v185)), float64(1e-06)) == int32(0))&base.F64_ne(v179, v185) != 0 {
		goto L35
	} else {
		goto L52
	}
L40:
	;
	v175 = v172
	goto L42
L41:
	;
	v175 = int32(0)
	goto L42
L42:
	;
	v178 = v37 + v175<<(uint(int32(4))%32)
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v178)))
	if base.Ui64(base.I64_reinterpret_f64(v179)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v170)))
	v187 = int64(9223372036854775807)
	v188 = base.I64_reinterpret_f64(v185) & v187
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v178)+8))
	v192 = base.I64_reinterpret_f64(v189) & v187
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v192) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	if base.Ui64(v206&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L35
	} else {
		goto L51
	}
L46:
	;
	v230 = base.B2i32(base.Ui64(v188) < base.Ui64(int64(9218868437227405313)))
	goto L38
L47:
	;
	goto L48
L48:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v188) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v170)+8))
	if base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v230 = int32(1)
	goto L38
L51:
	;
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v178)+8))
	v235 = v211
	v240 = base.I64_reinterpret_f64(v211) & int64(9223372036854775807)
	goto L37
L52:
	;
	if base.F64_eq(v189, v199)|base.F64_le(base.F64_abs(base.F64_sub(v189, v199)), float64(1e-06)) != 0 {
		goto L36
	} else {
		goto L53
	}
L53:
	;
	goto L35
L54:
	;
	v235 = v189
	v240 = v192
	goto L37
L55:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v244))|base.F64_ne(v241, v235) != 0 {
		goto L35
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui64(v244) < base.Ui64(int64(9218868437227405313)) {
		goto L35
	} else {
		goto L59
	}
L58:
	;
	goto L36
L59:
	;
	goto L36
L60:
	;
	v453 = v146
	goto L4
L61:
	;
	v277 = v55
	v281 = int32(1)
	goto L62
L62:
	;
	v293 = v39 + v281<<(uint(int32(4))%32)
	v295 = v277 - int32(1)
	if v295 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	if v28 == v281 {
		v453 = v146
		goto L4
	} else {
		goto L90
	}
L64:
	;
	goto L63
L65:
	;
	v384 = v281 + int32(1)
	if v384 != v28 {
		v277 = v298
		v281 = v384
		goto L62
	} else {
		goto L89
	}
L66:
	;
	v364 = *(*float64)(unsafe.Add(mBase, uint32(v293)+8))
	v367 = base.I64_reinterpret_f64(v364) & int64(9223372036854775807)
	if base.Ui64(v363) <= base.Ui64(int64(9218868437227405312)) {
		goto L84
	} else {
		goto L85
	}
L67:
	;
	if base.B2i32(v353 == int32(0))|base.F64_ne(v302, v308) != 0 {
		goto L64
	} else {
		goto L83
	}
L68:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v302, v308)), float64(1e-06)) == int32(0))&base.F64_ne(v302, v308) != 0 {
		goto L64
	} else {
		goto L81
	}
L69:
	;
	v298 = v28 - int32(1)
	goto L71
L70:
	;
	v298 = v295
	goto L71
L71:
	;
	v301 = v37 + v298<<(uint(int32(4))%32)
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v301)))
	if base.Ui64(base.I64_reinterpret_f64(v302)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v293)))
	v310 = int64(9223372036854775807)
	v311 = base.I64_reinterpret_f64(v308) & v310
	v312 = *(*float64)(unsafe.Add(mBase, uint32(v301)+8))
	v315 = base.I64_reinterpret_f64(v312) & v310
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v315) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	if base.Ui64(v329&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L64
	} else {
		goto L80
	}
L75:
	;
	v353 = base.B2i32(base.Ui64(v311) < base.Ui64(int64(9218868437227405313)))
	goto L67
L76:
	;
	goto L77
L77:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v311) {
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v322 = *(*float64)(unsafe.Add(mBase, uint32(v293)+8))
	if base.Ui64(base.I64_reinterpret_f64(v322)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L68
	} else {
		goto L79
	}
L79:
	;
	v353 = int32(1)
	goto L67
L80:
	;
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v301)+8))
	v358 = v334
	v363 = base.I64_reinterpret_f64(v334) & int64(9223372036854775807)
	goto L66
L81:
	;
	if base.F64_eq(v312, v322)|base.F64_le(base.F64_abs(base.F64_sub(v312, v322)), float64(1e-06)) != 0 {
		goto L65
	} else {
		goto L82
	}
L82:
	;
	goto L64
L83:
	;
	v358 = v312
	v363 = v315
	goto L66
L84:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v367))|base.F64_ne(v364, v358) != 0 {
		goto L64
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui64(v367) < base.Ui64(int64(9218868437227405313)) {
		goto L64
	} else {
		goto L88
	}
L87:
	;
	goto L65
L88:
	;
	goto L65
L89:
	;
	v453 = v146
	goto L4
L90:
	;
	goto L10
L91:
	;
	goto L9
L92:
	;
	F_pfree(m, v21)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v459 != v26 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	F_pfree(m, v26)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	return v453
L99:
	;
	goto L98
}
func F_poly_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_enlargeStringInfo(m, v8, int32(4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v26 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v21+v22))) = base.I32_rotr(v17, int32(24))&v26 | base.I32_rotr(v17&v26, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v21 + int32(4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v37 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v43 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68 << (uint(int32(2)) % 32)
	goto L13
L8:
	;
	v50 = v11 + int32(40) + v43<<(uint(int32(4))%32)
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v50)))
	F_pq_sendfloat8(m, v8, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v50)+8))
	F_pq_sendfloat8(m, v8, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v58 = v43 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v58 < v59 {
		v43 = v58
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	m.G0 = v8 + int32(16)
	return base.I64_extend_i32_u(v67)
}
func F_poly_to_circle(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v112 int32
	_ = v112
	var v115 float64
	_ = v115
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v137 float64
	_ = v137
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
	var v175 float64
	_ = v175
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	var v182 float64
	_ = v182
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v203 float64
	_ = v203
	var v204 int32
	_ = v204
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v217 float64
	_ = v217
	var v218 int32
	_ = v218
	var v219 float64
	_ = v219
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v227 float64
	_ = v227
	var v232 int32
	_ = v232
	var v235 float64
	_ = v235
	var v247 int32
	_ = v247
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 float64
	_ = v259
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v273 float64
	_ = v273
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 float64
	_ = v284
	var v286 float64
	_ = v286
	var v287 float64
	_ = v287
	var v299 float64
	_ = v299
	var v300 int32
	_ = v300
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 float64
	_ = v312
	var v316 int32
	_ = v316
	var v327 float64
	_ = v327
	var v328 int32
	_ = v328
	var v330 float64
	_ = v330
	var v333 float64
	_ = v333
	var v340 float64
	_ = v340
	var v341 int32
	_ = v341
	var v342 float64
	_ = v342
	var v347 float64
	_ = v347
	var v348 int32
	_ = v348
	var v349 float64
	_ = v349
	v8 = int32(0)
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v11
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v146
	if l2 != 0 {
		goto L43
	} else {
		goto L44
	}
L3:
	;
	v112 = v17
	v115 = float64(0)
	goto L5
L4:
	;
	v31 = v8
	goto L6
L5:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v115)&int64(9223372036854775807)))|v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v36 = l1 + int32(40) + v31<<(uint(int32(4))%32)
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v38 = base.F64_add(v33, v37)
	v40 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v38), v40)|base.F64_eq(base.F64_abs(v33), v40)|base.F64_eq(base.F64_abs(v37), v40) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v112 = v102
	v115 = v104
	goto L5
L8:
	;
	v52 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v54 = v38
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return
L12:
	;
	v54 = v52
	goto L10
L13:
	;
	v101 = v31 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v101 < v102 {
		v31 = v101
		goto L6
	} else {
		goto L31
	}
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v79
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v54
	goto L13
L15:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v94 != 0 {
		goto L1
	} else {
		goto L30
	}
L16:
	;
	v61 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v36)+8))
	v63 = base.F64_add(v61, v62)
	v65 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v63), v65)|base.F64_eq(base.F64_abs(v61), v65)|base.F64_eq(base.F64_abs(v62), v65) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v57 != int32(453) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v60 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v77 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	v79 = v63
	goto L22
L22:
	;
	if l2 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	v79 = v77
	goto L22
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v82 == int32(453) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v85 != 0 {
		goto L15
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v79
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v54
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v88 != int32(453) {
		goto L13
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L15
L30:
	;
	goto L13
L31:
	;
	goto L7
L32:
	;
	v124 = F_float_zero_divide_error_ext(m, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v127 = math.Float64frombits(uint64(0x7ff0000000000000))
	v130 = base.F64_div(v115, base.F64_convert_i32_s(v112))
	if base.F64_eq(base.F64_abs(v115), v127)|base.F64_ne(base.F64_abs(v130), v127) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v146 = v124
	goto L2
L36:
	;
	v137 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v139 = float64(0)
	if base.F64_eq(v115, v139)|base.F64_ne(v130, v139) != 0 {
		v146 = v130
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v146 = v137
	goto L2
L40:
	;
	v144 = F_float_underflow_error_ext(m, l2)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v146 = v144
	goto L2
L42:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v232 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v148 == int32(453) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if v192|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v193)&int64(9223372036854775807))) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v151 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v153 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if v152|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807))) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v186 != int32(453) {
		goto L42
	} else {
		goto L61
	}
L51:
	;
	v162 = F_float_zero_divide_error_ext(m, l2)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L11
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v165 = math.Float64frombits(uint64(0x7ff0000000000000))
	v168 = base.F64_div(v153, base.F64_convert_i32_s(v152))
	if base.F64_eq(base.F64_abs(v153), v165)|base.F64_ne(base.F64_abs(v168), v165) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v184 = v162
	goto L50
L55:
	;
	v175 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L11
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v177 = float64(0)
	if base.F64_eq(v153, v177)|base.F64_ne(v168, v177) != 0 {
		v184 = v168
		goto L50
	} else {
		goto L59
	}
L58:
	;
	v184 = v175
	goto L50
L59:
	;
	v182 = F_float_underflow_error_ext(m, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	v184 = v182
	goto L50
L61:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v189 == int32(0) {
		goto L42
	} else {
		goto L62
	}
L62:
	;
	goto L1
L63:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v227
	goto L42
L64:
	;
	v203 = F_float_zero_divide_error_ext(m, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v206 = math.Float64frombits(uint64(0x7ff0000000000000))
	v209 = base.F64_div(v193, base.F64_convert_i32_s(v192))
	if base.F64_eq(base.F64_abs(v193), v206)|base.F64_ne(base.F64_abs(v209), v206) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v227 = v203
	goto L63
L68:
	;
	v217 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v219 = float64(0)
	if base.F64_eq(v193, v219)|base.F64_ne(v209, v219) != 0 {
		v227 = v209
		goto L63
	} else {
		goto L72
	}
L71:
	;
	v227 = v217
	goto L63
L72:
	;
	v225 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v227 = v225
	goto L63
L74:
	;
	if v316|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v312)&int64(9223372036854775807))) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L75:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v312 = v235
	v316 = v232
	goto L74
L76:
	;
	goto L77
L77:
	;
	v247 = int32(0)
	goto L78
L78:
	;
	v252 = F_point_dt(m, l1+int32(40)+v247<<(uint(int32(4))%32), l0, l2)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L80
	}
L79:
	;
	v312 = v303
	v316 = v307
	goto L74
L80:
	;
	if l2 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v306 = v247 + int32(1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v306 < v307 {
		v247 = v306
		goto L78
	} else {
		goto L99
	}
L82:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v254 == int32(453) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v284 = math.Float64frombits(uint64(0x7ff0000000000000))
	v286 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v287 = base.F64_add(v252, v286)
	if base.F64_eq(base.F64_abs(v252), v284)|base.F64_ne(base.F64_abs(v287), v284)|base.F64_eq(base.F64_abs(v286), v284) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v257 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v259 = math.Float64frombits(uint64(0x7ff0000000000000))
	v261 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v262 = base.F64_add(v252, v261)
	if base.F64_eq(base.F64_abs(v252), v259)|base.F64_ne(base.F64_abs(v262), v259)|base.F64_eq(base.F64_abs(v261), v259) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	v273 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L11
	} else {
		goto L92
	}
L90:
	;
	v275 = v262
	goto L91
L91:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v277 != int32(453) {
		v303 = v275
		goto L81
	} else {
		goto L93
	}
L92:
	;
	v275 = v273
	goto L91
L93:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v280 == int32(0) {
		v303 = v275
		goto L81
	} else {
		goto L94
	}
L94:
	;
	goto L1
L95:
	;
	v299 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L11
	} else {
		goto L98
	}
L96:
	;
	v301 = v287
	goto L97
L97:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v301
	v303 = v301
	goto L81
L98:
	;
	v301 = v299
	goto L97
L99:
	;
	goto L79
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v349
	goto L1
L101:
	;
	v327 = F_float_zero_divide_error_ext(m, l2)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L11
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v330 = math.Float64frombits(uint64(0x7ff0000000000000))
	v333 = base.F64_div(v312, base.F64_convert_i32_s(v316))
	if base.F64_eq(base.F64_abs(v312), v330)|base.F64_ne(base.F64_abs(v333), v330) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v349 = v327
	goto L100
L105:
	;
	v340 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v342 = float64(0)
	if base.F64_eq(v312, v342)|base.F64_ne(v333, v342) != 0 {
		v349 = v333
		goto L100
	} else {
		goto L109
	}
L108:
	;
	v349 = v340
	goto L100
L109:
	;
	v347 = F_float_underflow_error_ext(m, l2)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	v349 = v347
	goto L100
}
