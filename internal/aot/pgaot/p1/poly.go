package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_poly_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 float64
	_ = v70
	var v73 int64
	_ = v73
	var v80 float64
	_ = v80
	var v87 int64
	_ = v87
	var v92 float64
	_ = v92
	var v111 int32
	_ = v111
	var v116 float64
	_ = v116
	var v120 int64
	_ = v120
	var v122 float64
	_ = v122
	var v125 int64
	_ = v125
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 float64
	_ = v174
	var v180 float64
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 float64
	_ = v184
	var v187 int64
	_ = v187
	var v194 float64
	_ = v194
	var v201 int64
	_ = v201
	var v206 float64
	_ = v206
	var v225 int32
	_ = v225
	var v230 float64
	_ = v230
	var v234 int64
	_ = v234
	var v236 float64
	_ = v236
	var v239 int64
	_ = v239
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v304 float64
	_ = v304
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v308 float64
	_ = v308
	var v311 int64
	_ = v311
	var v318 float64
	_ = v318
	var v325 int64
	_ = v325
	var v330 float64
	_ = v330
	var v349 int32
	_ = v349
	var v354 float64
	_ = v354
	var v358 int64
	_ = v358
	var v360 float64
	_ = v360
	var v363 int64
	_ = v363
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	v8 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.B2i32(v27 != v28)|base.B2i32(v27 <= int32(0)) != 0 {
		v421 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v432 != v20 {
		goto L90
	} else {
		goto L91
	}
L5:
	;
	v33 = int32(40)
	v34 = v25 + v33
	v36 = v20 + v33
	v54 = v8
	goto L6
L6:
	;
	v59 = v34 + v54<<(uint(int32(4))%32)
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v59)))
	if base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v421 = int32(0)
	goto L4
L8:
	;
	v411 = v54 + int32(1)
	if v411 != v27 {
		v54 = v411
		goto L6
	} else {
		goto L89
	}
L9:
	;
	v141 = int32(1)
	if v27 == v141 {
		v421 = v141
		goto L4
	} else {
		goto L30
	}
L10:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	v125 = base.I64_reinterpret_f64(v122) & int64(9223372036854775807)
	if base.Ui64(v120) <= base.Ui64(int64(9218868437227405312)) {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	if base.B2i32(v111 == int32(0))|base.F64_ne(v60, v66) != 0 {
		goto L8
	} else {
		goto L24
	}
L12:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v60, v66)), float64(1e-06)) == int32(0))&base.F64_ne(v60, v66) != 0 {
		goto L8
	} else {
		goto L22
	}
L13:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v68 = int64(9223372036854775807)
	v69 = base.I64_reinterpret_f64(v66) & v68
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v59)+8))
	v73 = base.I64_reinterpret_f64(v70) & v68
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v73) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	if base.Ui64(v87&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L8
	} else {
		goto L21
	}
L16:
	;
	v111 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313)))
	goto L11
L17:
	;
	goto L18
L18:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v69) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	if base.Ui64(base.I64_reinterpret_f64(v80)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(1)
	goto L11
L21:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v59)+8))
	v116 = v92
	v120 = base.I64_reinterpret_f64(v92) & int64(9223372036854775807)
	goto L10
L22:
	;
	if base.F64_eq(v70, v80)|base.F64_le(base.F64_abs(base.F64_sub(v70, v80)), float64(1e-06)) != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v116 = v70
	v120 = v73
	goto L10
L25:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v125))|base.F64_ne(v122, v116) != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui64(v125) < base.Ui64(int64(9218868437227405313)) {
		goto L8
	} else {
		goto L29
	}
L28:
	;
	goto L9
L29:
	;
	goto L9
L30:
	;
	v153 = v54
	v154 = v141
	goto L31
L31:
	;
	v165 = v36 + v154<<(uint(int32(4))%32)
	v167 = v153 + int32(1)
	if v167 < v27 {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v266 = int32(1)
	if v154 == v27 {
		v421 = v266
		goto L4
	} else {
		goto L59
	}
L33:
	;
	goto L32
L34:
	;
	v255 = int32(1)
	v257 = v154 + v255
	if v257 != v27 {
		v153 = v170
		v154 = v257
		goto L31
	} else {
		goto L58
	}
L35:
	;
	v236 = *(*float64)(unsafe.Add(mBase, uint32(v165)+8))
	v239 = base.I64_reinterpret_f64(v236) & int64(9223372036854775807)
	if base.Ui64(v234) <= base.Ui64(int64(9218868437227405312)) {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	if base.B2i32(v225 == int32(0))|base.F64_ne(v174, v180) != 0 {
		goto L33
	} else {
		goto L52
	}
L37:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v174, v180)), float64(1e-06)) == int32(0))&base.F64_ne(v174, v180) != 0 {
		goto L33
	} else {
		goto L50
	}
L38:
	;
	v170 = v167
	goto L40
L39:
	;
	v170 = int32(0)
	goto L40
L40:
	;
	v173 = v34 + v170<<(uint(int32(4))%32)
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v173)))
	if base.Ui64(base.I64_reinterpret_f64(v174)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v165)))
	v182 = int64(9223372036854775807)
	v183 = base.I64_reinterpret_f64(v180) & v182
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v173)+8))
	v187 = base.I64_reinterpret_f64(v184) & v182
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v187) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	if base.Ui64(v201&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L33
	} else {
		goto L49
	}
L44:
	;
	v225 = base.B2i32(base.Ui64(v183) < base.Ui64(int64(9218868437227405313)))
	goto L36
L45:
	;
	goto L46
L46:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v183) {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	v194 = *(*float64)(unsafe.Add(mBase, uint32(v165)+8))
	if base.Ui64(base.I64_reinterpret_f64(v194)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	v225 = int32(1)
	goto L36
L49:
	;
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v173)+8))
	v230 = v206
	v234 = base.I64_reinterpret_f64(v206) & int64(9223372036854775807)
	goto L35
L50:
	;
	if base.F64_eq(v184, v194)|base.F64_le(base.F64_abs(base.F64_sub(v184, v194)), float64(1e-06)) != 0 {
		goto L34
	} else {
		goto L51
	}
L51:
	;
	goto L33
L52:
	;
	v230 = v184
	v234 = v187
	goto L35
L53:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v239))|base.F64_ne(v236, v230) != 0 {
		goto L33
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui64(v239) < base.Ui64(int64(9218868437227405313)) {
		goto L33
	} else {
		goto L57
	}
L56:
	;
	goto L34
L57:
	;
	goto L34
L58:
	;
	v421 = v255
	goto L4
L59:
	;
	v277 = v266
	v279 = v54
	goto L60
L60:
	;
	v289 = v36 + v277<<(uint(int32(4))%32)
	v291 = v279 - int32(1)
	if v291 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	if v277 == v27 {
		v421 = int32(1)
		goto L4
	} else {
		goto L88
	}
L62:
	;
	goto L61
L63:
	;
	v379 = int32(1)
	v381 = v277 + v379
	if v381 != v27 {
		v277 = v381
		v279 = v294
		goto L60
	} else {
		goto L87
	}
L64:
	;
	v360 = *(*float64)(unsafe.Add(mBase, uint32(v289)+8))
	v363 = base.I64_reinterpret_f64(v360) & int64(9223372036854775807)
	if base.Ui64(v358) <= base.Ui64(int64(9218868437227405312)) {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	if base.B2i32(v349 == int32(0))|base.F64_ne(v298, v304) != 0 {
		goto L62
	} else {
		goto L81
	}
L66:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v298, v304)), float64(1e-06)) == int32(0))&base.F64_ne(v298, v304) != 0 {
		goto L62
	} else {
		goto L79
	}
L67:
	;
	v294 = v27 - int32(1)
	goto L69
L68:
	;
	v294 = v291
	goto L69
L69:
	;
	v297 = v34 + v294<<(uint(int32(4))%32)
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v297)))
	if base.Ui64(base.I64_reinterpret_f64(v298)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v289)))
	v306 = int64(9223372036854775807)
	v307 = base.I64_reinterpret_f64(v304) & v306
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v297)+8))
	v311 = base.I64_reinterpret_f64(v308) & v306
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v311) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
	if base.Ui64(v325&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L62
	} else {
		goto L78
	}
L73:
	;
	v349 = base.B2i32(base.Ui64(v307) < base.Ui64(int64(9218868437227405313)))
	goto L65
L74:
	;
	goto L75
L75:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v307) {
		goto L62
	} else {
		goto L76
	}
L76:
	;
	v318 = *(*float64)(unsafe.Add(mBase, uint32(v289)+8))
	if base.Ui64(base.I64_reinterpret_f64(v318)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L66
	} else {
		goto L77
	}
L77:
	;
	v349 = int32(1)
	goto L65
L78:
	;
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v297)+8))
	v354 = v330
	v358 = base.I64_reinterpret_f64(v330) & int64(9223372036854775807)
	goto L64
L79:
	;
	if base.F64_eq(v308, v318)|base.F64_le(base.F64_abs(base.F64_sub(v308, v318)), float64(1e-06)) != 0 {
		goto L63
	} else {
		goto L80
	}
L80:
	;
	goto L62
L81:
	;
	v354 = v308
	v358 = v311
	goto L64
L82:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v363))|base.F64_ne(v360, v354) != 0 {
		goto L62
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if base.Ui64(v363) < base.Ui64(int64(9218868437227405313)) {
		goto L62
	} else {
		goto L86
	}
L85:
	;
	goto L63
L86:
	;
	goto L63
L87:
	;
	v421 = v379
	goto L4
L88:
	;
	goto L8
L89:
	;
	goto L7
L90:
	;
	F_pfree(m, v20)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v436 != v25 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	F_pfree(m, v25)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	return v421
L97:
	;
	goto L96
}
func F_poly_send(m *base.Module, l0 int32) int32 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v67
}
func F_poly_to_circle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v95 float64
	_ = v95
	var v101 int32
	_ = v101
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v119 float64
	_ = v119
	var v126 int32
	_ = v126
	var v134 float64
	_ = v134
	var v139 int32
	_ = v139
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 float64
	_ = v167
	var v171 int32
	_ = v171
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	v3 = float64(0)
	v7 = int32(0)
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7 < v16 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L23
	} else {
		goto L32
	}
L2:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L31
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L23
	} else {
		goto L30
	}
L4:
	;
	v23 = v3
	v25 = v3
	v28 = v7
	goto L7
L5:
	;
	v71 = v3
	v73 = v3
	v75 = v16
	goto L6
L6:
	;
	if base.B2i32(v75 == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v71)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v32 = l1 + int32(40) + v28<<(uint(int32(4))%32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v34 = base.F64_add(v23, v33)
	v36 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v23), v36) == int32(0))&base.F64_ne(base.F64_abs(v33), v36) != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v71 = v34
	v73 = v52
	v75 = v67
	goto L6
L9:
	;
	v49 = math.Float64frombits(uint64(0x7ff0000000000000))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v52 = base.F64_add(v25, v51)
	if base.B2i32(base.F64_eq(base.F64_abs(v25), v49)|base.F64_ne(base.F64_abs(v52), v49) == int32(0))&base.F64_ne(base.F64_abs(v51), v49) != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v52
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v34
	v66 = v28 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 < v67 {
		v23 = v34
		v25 = v52
		v28 = v66
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v87 = base.F64_div(v71, base.F64_convert_i32_s(v75))
	v89 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v87), v89)&base.F64_ne(base.F64_abs(v71), v89) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v95 = float64(0)
	if base.F64_eq(v87, v95)&base.F64_ne(v71, v95) != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v87
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v101 == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v73)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v111 = base.F64_div(v73, base.F64_convert_i32_s(v101))
	v113 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v111), v113)&base.F64_ne(base.F64_abs(v73), v113) != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v119 = float64(0)
	if base.F64_eq(v111, v119)&base.F64_ne(v73, v119) != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v126 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v134 = v119
	v139 = int32(0)
	goto L21
L19:
	;
	v167 = v119
	v171 = v126
	goto L20
L20:
	;
	if base.B2i32(v171 == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v167)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
		goto L2
	} else {
		goto L27
	}
L21:
	;
	v144 = F_point_dt(m, l1+int32(40)+v139<<(uint(int32(4))%32), l0)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v167 = v146
	v171 = v163
	goto L20
L23:
	;
	return
L24:
	;
	v146 = base.F64_add(v134, v144)
	v148 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v146), v148)|base.F64_eq(base.F64_abs(v134), v148) == int32(0))&base.F64_ne(base.F64_abs(v144), v148) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v146
	v162 = v139 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v162 < v163 {
		v134 = v146
		v139 = v162
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v183 = base.F64_div(v167, base.F64_convert_i32_s(v171))
	v185 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v183), v185)&base.F64_ne(base.F64_abs(v167), v185) != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v191 = float64(0)
	if base.F64_eq(v183, v191)&base.F64_ne(v167, v191) != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v183
	return
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
