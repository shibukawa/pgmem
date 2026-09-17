package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parseUnicode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	if l1 < int32(3) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v400
L2:
	;
	v400 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = int32(2)
	v29 = int32(-1)
	goto L7
L5:
	;
	F_jsonpath_yyerror(m, l2, l3, int32(_a_F_parseUnicode_0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L62
	} else {
		goto L97
	}
L6:
	;
	F_jsonpath_yyerror(m, l2, l3, int32(_a_F_parseUnicode_0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L62
	} else {
		goto L96
	}
L7:
	;
	v30 = l0 + v27
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v31 != int32(123) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if v226 == int32(-1) {
		goto L87
	} else {
		goto L88
	}
L9:
	;
	v354 = v231 + int32(2)
	if v354 < l1 {
		v27 = v354
		v29 = v226
		goto L7
	} else {
		goto L86
	}
L10:
	;
	v345 = F_addUnicodeChar(m, v340, l2, l3)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L62
	} else {
		goto L83
	}
L11:
	;
	v340 = v29<<(uint(int32(10))%32)&int32(_a_F_parseUnicode_1) | v226&int32(1023) + int32(_a_F_parseUnicode_2)
	v342 = v231
	goto L10
L12:
	;
	v304 = int32(0)
	v305 = F_errsave_start(m, l2)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L62
	} else {
		goto L77
	}
L13:
	;
	v233 = v226 & int32(-1024)
	if v233 != int32(_a_F_parseUnicode_3) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30))))
	if base.Ui32((v103-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v124 = int32(-48)
		goto L31
	} else {
		goto L32
	}
L15:
	;
	if v27 < l1 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v38 = int32(0)
	v40 = v27 + int32(1)
	if l1 <= v40 {
		v95 = v38
		v97 = v27
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v29 != int32(-1) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v340 = int32(0)
	v342 = v27
	goto L10
L20:
	;
	v226 = v95
	v231 = v97 + int32(2)
	goto L13
L21:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v40))))
	if v43 == int32(125) {
		v95 = v38
		v97 = v27
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v50 = v40
	v51 = v38
	v52 = v43
	goto L23
L23:
	;
	if base.Ui32((v52-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v77 = int32(-48)
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v95 = v82
	v97 = v50
	goto L20
L25:
	;
	v82 = v77 + base.I32_extend8_s(v52) | v51<<(uint(int32(4))%32)
	v84 = v50 + int32(1)
	if l1 <= v84 {
		v95 = v82
		v97 = v50
		goto L20
	} else {
		goto L29
	}
L26:
	;
	if base.Ui32((v52-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v77 = int32(-87)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v52-int32(65))&int32(255)) {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v77 = int32(-55)
	goto L25
L29:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v84))))
	if v87 != int32(125) {
		v50 = v84
		v51 = v82
		v52 = v87
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	v125 = v124 + v103
	v127 = v27 + int32(1)
	if l1 <= v127 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if base.Ui32((v103-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v124 = int32(-87)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v103-int32(65))&int32(255)) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v124 = int32(-55)
	goto L31
L35:
	;
	v226 = v125
	v231 = v127
	goto L13
L36:
	;
	goto L37
L37:
	;
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v127))))
	if base.Ui32((v131-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v152 = int32(-48)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v156 = v131 + v152 | v125<<(uint(int32(4))%32)
	v158 = v27 + int32(2)
	if l1 <= v158 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if base.Ui32((v131-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v152 = int32(-87)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(int32(5)) < base.Ui32((v131-int32(65))&int32(255)) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v152 = int32(-55)
	goto L38
L42:
	;
	v226 = v156
	v231 = v158
	goto L13
L43:
	;
	goto L44
L44:
	;
	v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v158))))
	if base.Ui32((v162-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v183 = int32(-48)
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v187 = v162 + v183 | v156<<(uint(int32(4))%32)
	v189 = v27 + int32(3)
	if l1 <= v189 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if base.Ui32((v162-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v183 = int32(-87)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(5)) < base.Ui32((v162-int32(65))&int32(255)) {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v183 = int32(-55)
	goto L45
L49:
	;
	v226 = v187
	v231 = v189
	goto L13
L50:
	;
	goto L51
L51:
	;
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v189))))
	if base.Ui32((v193-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v214 = int32(-48)
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v216 = int32(4)
	v226 = v193 + v214 | v187<<(uint(v216)%32)
	v231 = v27 + v216
	goto L13
L53:
	;
	if base.Ui32((v193-int32(97))&int32(255)) < base.Ui32(int32(6)) {
		v214 = int32(-87)
		goto L52
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(int32(5)) < base.Ui32((v193-int32(65))&int32(255)) {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v214 = int32(-55)
	goto L52
L56:
	;
	if v29 == int32(-1) {
		v340 = v226
		v342 = v231
		goto L10
	} else {
		goto L76
	}
L57:
	;
	if v233 != int32(_a_F_parseUnicode_4) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v29 != int32(-1) {
		goto L11
	} else {
		goto L69
	}
L60:
	;
	if v29 == int32(-1) {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v240 = int32(0)
	v241 = F_errsave_start(m, l2)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return int32(0)
L63:
	;
	if v241 == int32(0) {
		v400 = v240
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_parseUnicode_5)
	F_errmsg(m, int32(_a_F_parseUnicode_6), v11+int32(-32))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	F_errdetail(m, int32(_a_F_parseUnicode_7), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, l2, int32(_a_F_parseUnicode_8), int32(619), int32(_a_F_parseUnicode_9))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v400 = v240
	goto L1
L69:
	;
	v268 = int32(0)
	v269 = F_errsave_start(m, l2)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	if v269 == int32(0) {
		v400 = v268
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_parseUnicode_5)
	F_errmsg(m, int32(_a_F_parseUnicode_6), v11+int32(-16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	F_errdetail(m, int32(_a_F_parseUnicode_10), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L62
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, l2, int32(_a_F_parseUnicode_8), int32(630), int32(_a_F_parseUnicode_9))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L62
	} else {
		goto L75
	}
L75:
	;
	v400 = v268
	goto L1
L76:
	;
	goto L12
L77:
	;
	if v305 == int32(0) {
		v400 = v304
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L62
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_parseUnicode_5)
	F_errmsg(m, int32(_a_F_parseUnicode_6), v13)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L62
	} else {
		goto L80
	}
L80:
	;
	F_errdetail(m, int32(_a_F_parseUnicode_10), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L62
	} else {
		goto L81
	}
L81:
	;
	F_errsave_finish(m, l2, int32(_a_F_parseUnicode_8), int32(640), int32(_a_F_parseUnicode_9))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L62
	} else {
		goto L82
	}
L82:
	;
	v400 = v304
	goto L1
L83:
	;
	if v345 == int32(0) {
		v400 = v345
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v351 = v342 + int32(2)
	if v351 < l1 {
		v27 = v351
		v29 = int32(-1)
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v400 = v345
	goto L1
L86:
	;
	goto L8
L87:
	;
	v400 = int32(1)
	goto L1
L88:
	;
	goto L89
L89:
	;
	v359 = int32(0)
	v360 = F_errsave_start(m, l2)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L62
	} else {
		goto L90
	}
L90:
	;
	if v360 == int32(0) {
		v400 = v359
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L62
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_parseUnicode_5)
	F_errmsg(m, int32(_a_F_parseUnicode_6), v11+int32(-48))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L62
	} else {
		goto L93
	}
L93:
	;
	F_errdetail(m, int32(_a_F_parseUnicode_10), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L62
	} else {
		goto L94
	}
L94:
	;
	F_errsave_finish(m, l2, int32(_a_F_parseUnicode_8), int32(692), int32(_a_F_parseUnicode_11))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L62
	} else {
		goto L95
	}
L95:
	;
	v400 = v359
	goto L1
L96:
	;
	v400 = int32(0)
	goto L1
L97:
	;
	v400 = int32(0)
	goto L1
}
func F_unicode_assigned(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v93 int32
	_ = v93
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_unicode_assigned[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	goto L3
L3:
	;
	if v16 == int32(6) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int32(1)
	v20 = v10 + v19
	v22 = v10 + int32(4)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L7:
	;
	v26 = v20
	goto L9
L8:
	;
	v26 = v22
	goto L9
L9:
	;
	if v23 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v54 = F_pg_mbstrlen_with_len(m, v26, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v35 = int32(16)
	goto L16
L15:
	;
	v35 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = int32(4)
	goto L19
L18:
	;
	v42 = v35
	goto L19
L19:
	;
	v53 = v42
	goto L10
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if int32(0) < v54 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v58&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	return int32(1)
L25:
	;
	v61 = v20
	goto L27
L26:
	;
	v61 = v22
	goto L27
L27:
	;
	v62 = v61
	v69 = int32(0)
	goto L28
L28:
	;
	v70 = int32(0)
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v73 = v71 & int32(255)
	if v70 <= v71 {
		v130 = v73
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L24
L30:
	;
	if v175&int32(255) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L31:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+uint32(_c_F_unicode_assigned[1]))))
	v175 = v166
	goto L30
L32:
	;
	if base.Ui32(int32(127)) < base.Ui32(v130) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	if v73&int32(224) == int32(192) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v62))))
	v130 = v125&int32(63) | v122
	goto L32
L35:
	;
	v122 = v73 << (uint(int32(6)) % 32) & int32(1984)
	v123 = int32(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v73&int32(240) == int32(224) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v122 = v73<<(uint(int32(12))%32)&int32(_a_F_unicode_assigned_0) | v93&int32(63)<<(uint(int32(6))%32)
	v123 = int32(2)
	goto L34
L39:
	;
	goto L40
L40:
	;
	if v73&int32(248) != int32(240) {
		v130 = int32(-1)
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v110 = int32(63)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	v122 = v73<<(uint(int32(18))%32)&int32(_a_F_unicode_assigned_1) | v109&v110<<(uint(int32(12))%32) | v115&v110<<(uint(int32(6))%32)
	v123 = int32(3)
	goto L34
L42:
	;
	v135 = int32(3367)
	v136 = v70
	goto L45
L43:
	;
	goto L44
L44:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130<<(uint(int32(1))%32))+uint32(_c_F_unicode_assigned[2]))))
	v175 = v165
	goto L30
L45:
	;
	v144 = base.I32_div_s(v135+v136, int32(2))
	v146 = v144 * int32(12)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+uint32(_c_F_unicode_assigned[3])))
	if base.Ui32(v149) < base.Ui32(v130) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v175 = int32(0)
	goto L30
L47:
	;
	if v160 <= v159 {
		v135 = v159
		v136 = v160
		goto L45
	} else {
		goto L52
	}
L48:
	;
	v159 = v135
	v160 = v144 + int32(1)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+uint32(_c_F_unicode_assigned[4])))
	if base.Ui32(v155) <= base.Ui32(v130) {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v159 = v144 - int32(1)
	v160 = v136
	goto L47
L52:
	;
	goto L46
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	v182 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	if int32(0) <= v182 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v209 = v69 + int32(1)
	if v209 != v54 {
		v62 = v206 + v62
		v69 = v209
		goto L28
	} else {
		goto L69
	}
L57:
	;
	v206 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v187 = v182 & int32(255)
	if v187&int32(224) == int32(192) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v206 = int32(2)
	goto L56
L61:
	;
	goto L62
L62:
	;
	if v187&int32(240) == int32(224) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v206 = int32(3)
	goto L56
L64:
	;
	goto L65
L65:
	;
	if v187&int32(248) == int32(240) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v204 = int32(4)
	goto L68
L67:
	;
	v204 = int32(1)
	goto L68
L68:
	;
	v206 = v204
	goto L56
L69:
	;
	goto L29
L70:
	;
	F_errmsg(m, int32(_a_F_unicode_assigned_2), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_unicode_assigned_3), int32(_a_F_unicode_assigned_4), int32(_a_F_unicode_assigned_5))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unicode_norm_form_from_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_unicode_norm_form_from_string[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L63
	} else {
		goto L68
	}
L2:
	;
	if v9 == int32(6) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = l0
	v17 = int32(_a_F_unicode_norm_form_from_string_0)
	goto L8
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L63
	} else {
		goto L64
	}
L6:
	;
	m.G0 = v5 + int32(16)
	return v190
L7:
	;
	if v54 == int32(0) {
		v190 = int32(0)
		goto L6
	} else {
		goto L20
	}
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v20 == v21 {
		v43 = v20
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v54 = int32(0)
	goto L7
L10:
	;
	v45 = int32(1)
	if v43 != 0 {
		v16 = v16 + v45
		v17 = v17 + v45
		goto L8
	} else {
		goto L19
	}
L11:
	;
	if base.Ui32((v20-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v20 | int32(32)
	goto L14
L13:
	;
	v31 = v20
	goto L14
L14:
	;
	if base.Ui32((v21-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = v21 | int32(32)
	goto L17
L16:
	;
	v40 = v21
	goto L17
L17:
	;
	if v31 == v40 {
		v43 = v31
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v54 = v31 - v40
	goto L7
L19:
	;
	goto L9
L20:
	;
	v61 = l0
	v62 = int32(_a_F_unicode_norm_form_from_string_1)
	goto L22
L21:
	;
	if v99 == int32(0) {
		v190 = int32(1)
		goto L6
	} else {
		goto L34
	}
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == v66 {
		v88 = v65
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v99 = int32(0)
	goto L21
L24:
	;
	v90 = int32(1)
	if v88 != 0 {
		v61 = v61 + v90
		v62 = v62 + v90
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v65-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v76 = v65 | int32(32)
	goto L28
L27:
	;
	v76 = v65
	goto L28
L28:
	;
	if base.Ui32((v66-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v85 = v66 | int32(32)
	goto L31
L30:
	;
	v85 = v66
	goto L31
L31:
	;
	if v76 == v85 {
		v88 = v76
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v99 = v76 - v85
	goto L21
L33:
	;
	goto L23
L34:
	;
	v106 = l0
	v107 = int32(_a_F_unicode_norm_form_from_string_2)
	goto L36
L35:
	;
	if v144 == int32(0) {
		v190 = int32(2)
		goto L6
	} else {
		goto L48
	}
L36:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 == v111 {
		v133 = v110
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v144 = int32(0)
	goto L35
L38:
	;
	v135 = int32(1)
	if v133 != 0 {
		v106 = v106 + v135
		v107 = v107 + v135
		goto L36
	} else {
		goto L47
	}
L39:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = v110 | int32(32)
	goto L42
L41:
	;
	v121 = v110
	goto L42
L42:
	;
	if base.Ui32((v111-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v111 | int32(32)
	goto L45
L44:
	;
	v130 = v111
	goto L45
L45:
	;
	if v121 == v130 {
		v133 = v121
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v144 = v121 - v130
	goto L35
L47:
	;
	goto L37
L48:
	;
	v150 = l0
	v151 = int32(_a_F_unicode_norm_form_from_string_3)
	goto L50
L49:
	;
	if v188 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 == v155 {
		v177 = v154
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v188 = int32(0)
	goto L49
L52:
	;
	v179 = int32(1)
	if v177 != 0 {
		v150 = v150 + v179
		v151 = v151 + v179
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v154-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v165 = v154 | int32(32)
	goto L56
L55:
	;
	v165 = v154
	goto L56
L56:
	;
	if base.Ui32((v155-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v174 = v155 | int32(32)
	goto L59
L58:
	;
	v174 = v155
	goto L59
L59:
	;
	if v165 == v174 {
		v177 = v165
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v188 = v165 - v174
	goto L49
L61:
	;
	goto L51
L62:
	;
	v190 = int32(3)
	goto L6
L63:
	;
	return int32(0)
L64:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_unicode_norm_form_from_string_4), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_unicode_norm_form_from_string_5), int32(_a_F_unicode_norm_form_from_string_6), int32(_a_F_unicode_norm_form_from_string_7))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
	F_errmsg(m, int32(_a_F_unicode_norm_form_from_string_8), v5)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_unicode_norm_form_from_string_5), int32(_a_F_unicode_norm_form_from_string_9), int32(_a_F_unicode_norm_form_from_string_7))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L63
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
