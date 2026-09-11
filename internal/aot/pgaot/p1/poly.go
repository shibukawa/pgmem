package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_poly_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v79 float64
	_ = v79
	var v86 int64
	_ = v86
	var v91 float64
	_ = v91
	var v107 int32
	_ = v107
	var v114 int64
	_ = v114
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v121 int64
	_ = v121
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v175 float64
	_ = v175
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 float64
	_ = v179
	var v182 int64
	_ = v182
	var v189 float64
	_ = v189
	var v196 int64
	_ = v196
	var v201 float64
	_ = v201
	var v217 int32
	_ = v217
	var v224 int64
	_ = v224
	var v225 float64
	_ = v225
	var v228 float64
	_ = v228
	var v231 int64
	_ = v231
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v295 float64
	_ = v295
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v299 float64
	_ = v299
	var v302 int64
	_ = v302
	var v309 float64
	_ = v309
	var v316 int64
	_ = v316
	var v321 float64
	_ = v321
	var v337 int32
	_ = v337
	var v344 int64
	_ = v344
	var v345 float64
	_ = v345
	var v348 float64
	_ = v348
	var v351 int64
	_ = v351
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	v2 = int32(0)
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
	if v27 != v28 {
		v402 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v419 != v20 {
		goto L109
	} else {
		goto L110
	}
L5:
	;
	if v27 <= int32(0) {
		v402 = v2
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v32 = int32(40)
	v33 = v25 + v32
	v35 = v20 + v32
	v45 = v2
	goto L7
L7:
	;
	v58 = v33 + v45<<(uint(int32(4))%32)
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
	if base.Ui64(base.I64_reinterpret_f64(v59)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v402 = int32(0)
	goto L4
L9:
	;
	v398 = v45 + int32(1)
	if v398 != v27 {
		v45 = v398
		goto L7
	} else {
		goto L108
	}
L10:
	;
	v136 = int32(1)
	if v27 == v136 {
		v402 = v136
		goto L4
	} else {
		goto L37
	}
L11:
	;
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	v121 = base.I64_reinterpret_f64(v118) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v114) {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	if base.F64_ne(v59, v65) != 0 {
		goto L9
	} else {
		goto L29
	}
L13:
	;
	if base.F64_ne(v59, v65) != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v65 = *(*float64)(unsafe.Add(mBase, uint32(v35)))
	v67 = int64(9223372036854775807)
	v68 = base.I64_reinterpret_f64(v65) & v67
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v58)+8))
	v72 = base.I64_reinterpret_f64(v69) & v67
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v72) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	if base.Ui64(v86&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L9
	} else {
		goto L22
	}
L17:
	;
	v107 = base.B2i32(base.Ui64(v68) < base.Ui64(int64(9218868437227405313)))
	goto L12
L18:
	;
	goto L19
L19:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v68) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	if base.Ui64(base.I64_reinterpret_f64(v79)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v107 = int32(1)
	goto L12
L22:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v58)+8))
	v114 = base.I64_reinterpret_f64(v91) & int64(9223372036854775807)
	v115 = v91
	goto L11
L23:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v59, v65)), float64(1e-06)) == int32(0) {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if base.F64_eq(v69, v79) != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v69, v79)), float64(1e-06)) != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	goto L9
L29:
	;
	if v107 == int32(0) {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v114 = v72
	v115 = v69
	goto L11
L31:
	;
	if base.Ui64(v121) <= base.Ui64(int64(9218868437227405312)) {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if base.F64_ne(v115, v118) != 0 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L10
L35:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v121) {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L10
L37:
	;
	v142 = v136
	v144 = v45
	goto L38
L38:
	;
	v160 = v35 + v142<<(uint(int32(4))%32)
	v162 = v144 + int32(1)
	if v162 < v27 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v257 = int32(1)
	if v142 == v27 {
		v402 = v257
		goto L4
	} else {
		goto L72
	}
L40:
	;
	goto L39
L41:
	;
	v246 = int32(1)
	v248 = v142 + v246
	if v248 != v27 {
		v142 = v248
		v144 = v165
		goto L38
	} else {
		goto L71
	}
L42:
	;
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v160)+8))
	v231 = base.I64_reinterpret_f64(v228) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v224) {
		goto L65
	} else {
		goto L66
	}
L43:
	;
	if base.F64_ne(v169, v175) != 0 {
		goto L40
	} else {
		goto L63
	}
L44:
	;
	if base.F64_ne(v169, v175) != 0 {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v165 = v162
	goto L47
L46:
	;
	v165 = int32(0)
	goto L47
L47:
	;
	v168 = v33 + v165<<(uint(int32(4))%32)
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v168)))
	if base.Ui64(base.I64_reinterpret_f64(v169)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v160)))
	v177 = int64(9223372036854775807)
	v178 = base.I64_reinterpret_f64(v175) & v177
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v168)+8))
	v182 = base.I64_reinterpret_f64(v179) & v177
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v182) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	if base.Ui64(v196&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L40
	} else {
		goto L56
	}
L51:
	;
	v217 = base.B2i32(base.Ui64(v178) < base.Ui64(int64(9218868437227405313)))
	goto L43
L52:
	;
	goto L53
L53:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v178) {
		goto L40
	} else {
		goto L54
	}
L54:
	;
	v189 = *(*float64)(unsafe.Add(mBase, uint32(v160)+8))
	if base.Ui64(base.I64_reinterpret_f64(v189)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	v217 = int32(1)
	goto L43
L56:
	;
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v168)+8))
	v224 = base.I64_reinterpret_f64(v201) & int64(9223372036854775807)
	v225 = v201
	goto L42
L57:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v169, v175)), float64(1e-06)) == int32(0) {
		goto L40
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if base.F64_eq(v179, v189) != 0 {
		goto L41
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v179, v189)), float64(1e-06)) != 0 {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	goto L40
L63:
	;
	if v217 == int32(0) {
		goto L40
	} else {
		goto L64
	}
L64:
	;
	v224 = v182
	v225 = v179
	goto L42
L65:
	;
	if base.Ui64(v231) <= base.Ui64(int64(9218868437227405312)) {
		goto L40
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if base.F64_ne(v228, v225) != 0 {
		goto L40
	} else {
		goto L69
	}
L68:
	;
	goto L41
L69:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v231) {
		goto L40
	} else {
		goto L70
	}
L70:
	;
	goto L41
L71:
	;
	v402 = v246
	goto L4
L72:
	;
	v264 = v257
	v265 = v45
	goto L73
L73:
	;
	v280 = v35 + v264<<(uint(int32(4))%32)
	v282 = v265 - int32(1)
	if v282 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	if v27 == v264 {
		v402 = int32(1)
		goto L4
	} else {
		goto L107
	}
L75:
	;
	goto L74
L76:
	;
	v366 = int32(1)
	v368 = v264 + v366
	if v368 != v27 {
		v264 = v368
		v265 = v285
		goto L73
	} else {
		goto L106
	}
L77:
	;
	v348 = *(*float64)(unsafe.Add(mBase, uint32(v280)+8))
	v351 = base.I64_reinterpret_f64(v348) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v344) {
		goto L100
	} else {
		goto L101
	}
L78:
	;
	if base.F64_ne(v289, v295) != 0 {
		goto L75
	} else {
		goto L98
	}
L79:
	;
	if base.F64_ne(v289, v295) != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v285 = v27 - int32(1)
	goto L82
L81:
	;
	v285 = v282
	goto L82
L82:
	;
	v288 = v33 + v285<<(uint(int32(4))%32)
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.Ui64(base.I64_reinterpret_f64(v289)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v295 = *(*float64)(unsafe.Add(mBase, uint32(v280)))
	v297 = int64(9223372036854775807)
	v298 = base.I64_reinterpret_f64(v295) & v297
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v288)+8))
	v302 = base.I64_reinterpret_f64(v299) & v297
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v302) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	if base.Ui64(v316&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L75
	} else {
		goto L91
	}
L86:
	;
	v337 = base.B2i32(base.Ui64(v298) < base.Ui64(int64(9218868437227405313)))
	goto L78
L87:
	;
	goto L88
L88:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v298) {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v280)+8))
	if base.Ui64(base.I64_reinterpret_f64(v309)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L79
	} else {
		goto L90
	}
L90:
	;
	v337 = int32(1)
	goto L78
L91:
	;
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v288)+8))
	v344 = base.I64_reinterpret_f64(v321) & int64(9223372036854775807)
	v345 = v321
	goto L77
L92:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v289, v295)), float64(1e-06)) == int32(0) {
		goto L75
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if base.F64_eq(v299, v309) != 0 {
		goto L76
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v299, v309)), float64(1e-06)) != 0 {
		goto L76
	} else {
		goto L97
	}
L97:
	;
	goto L75
L98:
	;
	if v337 == int32(0) {
		goto L75
	} else {
		goto L99
	}
L99:
	;
	v344 = v302
	v345 = v299
	goto L77
L100:
	;
	if base.Ui64(v351) <= base.Ui64(int64(9218868437227405312)) {
		goto L75
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if base.F64_ne(v348, v345) != 0 {
		goto L75
	} else {
		goto L104
	}
L103:
	;
	goto L76
L104:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v351) {
		goto L75
	} else {
		goto L105
	}
L105:
	;
	goto L76
L106:
	;
	v402 = v366
	goto L4
L107:
	;
	goto L9
L108:
	;
	goto L8
L109:
	;
	F_pfree(m, v20)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v423 != v25 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	F_pfree(m, v25)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	return v402
L116:
	;
	goto L115
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	v24 = int32(24)
	v26 = int32(65280)
	v28 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v21+v22))) = v17<<(uint(v24)%32) | v17&v26<<(uint(v28)%32) | (int32(base.Ui32(v17)>>(uint(v28)%32))&v26 | int32(base.Ui32(v17)>>(uint(v24)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v21 + int32(4)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v43 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v49 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 << (uint(int32(2)) % 32)
	goto L13
L8:
	;
	v56 = v11 + int32(40) + v49<<(uint(int32(4))%32)
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
	F_pq_sendfloat8(m, v8, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
	F_pq_sendfloat8(m, v8, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v64 = v49 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v64 < v65 {
		v49 = v64
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
	return v73
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
	var v24 float64
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v67 int32
	_ = v67
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v92 int32
	_ = v92
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v109 float64
	_ = v109
	var v116 int32
	_ = v116
	var v124 float64
	_ = v124
	var v129 int32
	_ = v129
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v136 float64
	_ = v136
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v176 float64
	_ = v176
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	v3 = float64(0)
	v7 = int32(0)
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 <= v7 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L37
	} else {
		goto L51
	}
L2:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L37
	} else {
		goto L50
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L37
	} else {
		goto L49
	}
L4:
	;
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v63 = v3
	v64 = v3
	v67 = v16
	goto L4
L6:
	;
	goto L7
L7:
	;
	v23 = v3
	v24 = v3
	v28 = v7
	goto L8
L8:
	;
	v32 = l1 + int32(40) + v28<<(uint(int32(4))%32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v34 = base.F64_add(v23, v33)
	if base.F64_ne(base.F64_abs(v34), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v63 = v34
	v64 = v45
	v67 = v59
	goto L4
L10:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v45 = base.F64_add(v24, v44)
	if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if base.F64_eq(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.F64_ne(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v45
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v34
	v58 = v28 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v58 < v59 {
		v23 = v34
		v24 = v45
		v28 = v58
		goto L8
	} else {
		goto L18
	}
L15:
	;
	if base.F64_eq(base.F64_abs(v24), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.F64_ne(base.F64_abs(v44), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L9
L19:
	;
	if base.Ui64(base.I64_reinterpret_f64(v63)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v78 = base.F64_div(v63, base.F64_convert_i32_s(v67))
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v78), v80)&base.F64_ne(base.F64_abs(v63), v80) != 0 {
		goto L3
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v86 = float64(0)
	if base.F64_eq(v78, v86)&base.F64_ne(v63, v86) != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v78
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v101 = base.F64_div(v64, base.F64_convert_i32_s(v92))
	v103 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v101), v103)&base.F64_ne(base.F64_abs(v64), v103) != 0 {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v109 = float64(0)
	if base.F64_eq(v101, v109)&base.F64_ne(v64, v109) != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v101
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v116 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v124 = v109
	v129 = int32(0)
	goto L34
L32:
	;
	v153 = v109
	v157 = v116
	goto L33
L33:
	;
	if v157 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v134 = F_point_dt(m, l1+int32(40)+v129<<(uint(int32(4))%32), l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v153 = v136
	v157 = v149
	goto L33
L36:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v136
	v148 = v129 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v148 < v149 {
		v124 = v136
		v129 = v148
		goto L34
	} else {
		goto L42
	}
L37:
	;
	return
L38:
	;
	v136 = base.F64_add(v124, v134)
	if base.F64_ne(base.F64_abs(v136), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if base.F64_eq(base.F64_abs(v124), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	if base.F64_ne(base.F64_abs(v134), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	goto L35
L43:
	;
	if base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v168 = base.F64_div(v153, base.F64_convert_i32_s(v157))
	v170 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v168), v170)&base.F64_ne(base.F64_abs(v153), v170) != 0 {
		goto L3
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v176 = float64(0)
	if base.F64_eq(v168, v176)&base.F64_ne(v153, v176) != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v168
	return
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
