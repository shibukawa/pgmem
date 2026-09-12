package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_at_local(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
	v7 = F_cstring_to_text(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_DirectFunctionCall2Coll(m, int32(1290), int32(0), v7, v2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_timetz_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v9 = int64(1000000)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v12 = base.I64_extend_i32_s(v7)*v9 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v19 = base.I64_extend_i32_s(v14)*v9 + v18
	if v19 < v12 {
		return int32(1)
	} else {
		if v12 < v19 {
			return int32(0)
		} else {
			return base.B2i32(v14 < v7)
		}
	}
}
func F_timetz_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 float64
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
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
	v24 = int32(1)
	v25 = v20 + v24
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v30 = v28 & v24
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = v25
	goto L5
L4:
	;
	v31 = v20 + int32(4)
	goto L5
L5:
	;
	if v28 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v62 = F_downcase_truncate_identifier(m, v31, v59, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	v34 = int32(4)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v36&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v49 = int32(1)
	if v30 != 0 {
		v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v45 = v34
	goto L12
L11:
	;
	v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
	goto L12
L12:
	;
	if v36 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = v34
	goto L15
L14:
	;
	v48 = v45
	goto L15
L15:
	;
	v59 = v48
	goto L6
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	m.G0 = v15 + int32(32)
	return v359
L18:
	;
	v65 = v15 + int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v133 == int32(31) {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1067])) = v116
	v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v116)+11)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v124
	v133 = v123
	goto L19
L21:
	;
	v74 = F_strncmp(m, v62, v72, int32(10))
	mBase = m.M
	if v74 == int32(0) {
		v116 = v72
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v83 = int32(1663648)
	v85 = int32(1664608)
	goto L25
L24:
	;
	goto L23
L25:
	;
	v92 = v83 + (v85-v83)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92))))
	v94 = v77 - v93
	if v94 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(0)
	v133 = int32(31)
	goto L19
L27:
	;
	v98 = F_strncmp(m, v62, v92, int32(10))
	mBase = m.M
	if v98 == int32(0) {
		v116 = v92
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v101 = v94
	goto L29
L29:
	;
	v105 = base.B2i32(v101 < int32(0))
	if v101 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v101 = v98
	goto L29
L31:
	;
	v106 = v92 - int32(16)
	goto L33
L32:
	;
	v106 = v85
	goto L33
L33:
	;
	if v101 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v109 = v83
	goto L36
L35:
	;
	v109 = v92 + int32(16)
	goto L36
L36:
	;
	if base.Ui32(v109) <= base.Ui32(v106) {
		v83 = v109
		v85 = v106
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	v137 = v15 + int32(28)
	v144 = *(*int32)(unsafe.Add(mBase, _consts[1068]))
	if v144 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v206 = v133
	goto L40
L40:
	;
	if v206 == int32(17) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	v206 = v205
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1068])) = v188
	v195 = int32(*(*int8)(unsafe.Add(mBase, uint32(v188)+11)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v196
	v205 = v195
	goto L41
L43:
	;
	v146 = F_strncmp(m, v62, v144, int32(10))
	mBase = m.M
	if v146 == int32(0) {
		v188 = v144
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v149 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v155 = int32(1662496)
	v157 = int32(1663632)
	goto L47
L46:
	;
	goto L45
L47:
	;
	v164 = v155 + (v157-v155)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v165 = int32(*(*int8)(unsafe.Add(mBase, uint32(v164))))
	v166 = v149 - v165
	if v166 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(0)
	v205 = int32(31)
	goto L41
L49:
	;
	v170 = F_strncmp(m, v62, v164, int32(10))
	mBase = m.M
	if v170 == int32(0) {
		v188 = v164
		goto L42
	} else {
		goto L52
	}
L50:
	;
	v173 = v166
	goto L51
L51:
	;
	v177 = base.B2i32(v173 < int32(0))
	if v173 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v173 = v170
	goto L51
L53:
	;
	v178 = v164 - int32(16)
	goto L55
L54:
	;
	v178 = v157
	goto L55
L55:
	;
	if v173 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v181 = v155
	goto L58
L57:
	;
	v181 = v164 + int32(16)
	goto L58
L58:
	;
	if base.Ui32(v181) <= base.Ui32(v178) {
		v155 = v181
		v157 = v178
		goto L47
	} else {
		goto L59
	}
L59:
	;
	goto L48
L60:
	;
	if l1 != 0 {
		goto L100
	} else {
		goto L101
	}
L61:
	;
	v346 = base.I64_extend_i32_s(int32(0) - v229)
	goto L60
L62:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	v211 = base.I64_div_s(v209, int64(3600000000))
	v212 = base.I64_extend32_s(v211)
	v215 = v212*int64(-3600000000) + v209
	v217 = base.I64_div_s(v215, int64(60000000))
	v218 = base.I64_extend32_s(v217)
	v221 = v218*int64(-60000000) + v215
	v223 = base.I64_div_s(v221, int64(1000000))
	v226 = v223*int64(4293967296) + v221
	v227 = base.I32_wrap_i64(v226)
	v228 = base.I32_wrap_i64(v223)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	switch v230 - int32(4) {
	case 0:
		goto L61
	default:
		goto L65
	case 14:
		goto L67
	case 15:
		v346 = v218
		goto L60
	case 16:
		goto L66
	case 25:
		goto L68
	case 26:
		goto L69
	case 30:
		goto L70
	case 31:
		goto L71
	}
L63:
	;
	goto L64
L64:
	;
	if v206 != 0 {
		goto L87
	} else {
		goto L88
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L82
	}
L66:
	;
	v346 = v212
	goto L60
L67:
	;
	if l1 != 0 {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	if l1 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v346 = base.I64_extend32_s(v226) + base.I64_extend32_s(v223)*int64(1000000)
	goto L60
L70:
	;
	v243 = base.I32_div_s(int32(0)-v229, int32(3600))
	v346 = base.I64_extend_i32_s(v243)
	goto L60
L71:
	;
	v235 = int32(60)
	v236 = base.I32_div_s(int32(0)-v229, v235)
	v238 = base.I32_rem_s(v236, v235)
	v346 = base.I64_extend_i32_s(v238)
	goto L60
L72:
	;
	v256 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v226)+base.I64_extend32_s(v223)*int64(1000000), int32(3))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v259 = float64(1000)
	v265 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v228), v259), base.F64_div(base.F64_convert_i32_s(v227), v259)))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v359 = v256
	goto L17
L76:
	;
	v359 = v265
	goto L17
L77:
	;
	v273 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v226)+base.I64_extend32_s(v223)*int64(1000000), int32(6))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v280 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v227), float64(1e+06)), base.F64_convert_i32_s(v228)))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v359 = v273
	goto L17
L81:
	;
	v359 = v280
	goto L17
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v290 = F_format_type_be(m, int32(1266))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v62
	F_errmsg(m, int32(190497), v15)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(499833), int32(3077), int32(246150))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L95
	}
L88:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v302 != int32(11) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	if l1 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v306 = int64(*(*int32)(unsafe.Add(mBase, uint32(v60)+8)))
	v311 = F_int64_div_fast_to_numeric(m, v306*int64(1000000)+v305, int32(6))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v319 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i64_s(v305), float64(1e+06)), base.F64_convert_i32_s(v316)))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v359 = v311
	goto L17
L94:
	;
	v359 = v319
	goto L17
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v329 = F_format_type_be(m, int32(1266))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v62
	F_errmsg(m, int32(190460), v15+int32(16))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(499833), int32(3097), int32(246150))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v347 = F_int64_to_numeric(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v350 = F_Float8GetDatum(m, base.F64_convert_i64_s(v346))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v359 = v347
	goto L17
L104:
	;
	v359 = v350
	goto L17
}
func F_timetz_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
