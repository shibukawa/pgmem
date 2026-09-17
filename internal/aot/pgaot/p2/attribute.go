package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutText(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_strlen(m, l1)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = F_pg_server_to_any(m, l1, v12, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = l1
	goto L3
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v16 = v14
	goto L3
L6:
	;
	return
L7:
	;
	if base.Ui32(v287) <= base.Ui32(v288) {
		goto L6
	} else {
		goto L98
	}
L8:
	;
	v159 = v16
	v160 = v17
	v161 = v16
	goto L57
L9:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v26 = v16
	v27 = v16
	v28 = v17
	goto L14
L14:
	;
	if base.Ui32(v28) <= base.Ui32(int32(31)) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v287 = v155
	v288 = v152
	goto L7
L16:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v156 != 0 {
		v26 = v155
		v27 = v152
		v28 = v156
		goto L14
	} else {
		goto L56
	}
L17:
	;
	v34 = v28 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v34&int32(255)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	if base.B2i32(v28 != int32(92))&base.B2i32(v28 != v8) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	if base.Ui32(v27) < base.Ui32(v26) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	if v28 == v8 {
		v51 = v8
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v51 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v34<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L20
L24:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v42 = v26 + int32(1)
	if v40 != 0 {
		v26 = v42
		v28 = v40
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v287 = v42
	v288 = v27
	goto L7
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v53, v27, v26-v27)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v61 <= v58+int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v85 <= v82+int32(1) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	F_appendStringInfoChar(m, v57, int32(92))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v68 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v66+v58))) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v73 = v71 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v73))) = uint8(v77)
	goto L30
L34:
	;
	goto L30
L35:
	;
	v105 = v26 + int32(1)
	v152 = v105
	v155 = v105
	goto L16
L36:
	;
	F_appendStringInfoChar(m, v81, base.I32_extend8_s(v51))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v82))) = uint8(v51)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v96 = v94 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v96))) = uint8(v100)
	goto L35
L39:
	;
	goto L35
L40:
	;
	if base.Ui32(v27) < base.Ui32(v26) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if base.I32_extend8_s(v28) < int32(0) {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v113, v27, v26-v27)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v121 <= v118+int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v152 = v26
	v155 = v26 + int32(1)
	goto L16
L48:
	;
	F_appendStringInfoChar(m, v117, int32(92))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v128 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v118))) = uint8(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v133 = v131 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v133))) = uint8(v137)
	goto L47
L51:
	;
	goto L47
L52:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v147 = F_pg_encoding_mblen(m, v146, v26)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v152 = v27
	v155 = v26 + int32(1)
	goto L16
L55:
	;
	v152 = v27
	v155 = v147 + v26
	goto L16
L56:
	;
	goto L15
L57:
	;
	if base.Ui32(v160) <= base.Ui32(int32(31)) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v283 == int32(0) {
		v287 = v282
		v288 = v280
		goto L7
	} else {
		goto L97
	}
L60:
	;
	v166 = v160 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v166&int32(255)) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	if base.B2i32(v160 != int32(92))&base.B2i32(v160 != v8) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L63:
	;
	if base.Ui32(v159) < base.Ui32(v161) {
		goto L71
	} else {
		goto L72
	}
L64:
	;
	if v160 != v8 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v185 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v166<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L63
L67:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	v174 = v161 + int32(1)
	if v172 == int32(0) {
		v287 = v174
		v288 = v159
		goto L7
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v185 = v8
	goto L63
L70:
	;
	v160 = v172
	v161 = v174
	goto L57
L71:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v187, v159, v161-v159)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v195 <= v192+int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	if v219 <= v216+int32(1) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	F_appendStringInfoChar(m, v191, int32(92))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v202 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v200+v192))) = uint8(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v207 = v205 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v209+v207))) = uint8(v211)
	goto L75
L79:
	;
	goto L75
L80:
	;
	v239 = v161 + int32(1)
	v280 = v239
	v282 = v239
	goto L59
L81:
	;
	F_appendStringInfoChar(m, v215, base.I32_extend8_s(v185))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224+v216))) = uint8(v185)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v230 = v228 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v230))) = uint8(v234)
	goto L80
L84:
	;
	goto L80
L85:
	;
	if base.Ui32(v159) < base.Ui32(v161) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v280 = v159
	v282 = v161 + int32(1)
	goto L59
L88:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v247, v159, v161-v159)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	if v255 <= v252+int32(1) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L90
L92:
	;
	v280 = v161
	v282 = v161 + int32(1)
	goto L59
L93:
	;
	F_appendStringInfoChar(m, v251, int32(92))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v262 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v260+v252))) = uint8(v262)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v267 = v265 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v267))) = uint8(v271)
	goto L92
L96:
	;
	goto L92
L97:
	;
	v159 = v280
	v160 = v283
	v161 = v282
	goto L57
L98:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v293, v288, v287-v288)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	goto L6
}
func F_attribute_reloptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(64), int32(24), int32(_a_F_attribute_reloptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_attribute_options(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1556)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(51539607560)
		v24 = F_hash_create(m, int32(_a_F_get_attribute_options_0), int32(256), v6+int32(-48), int32(72))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0])) = v24
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
			if v30 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(7), int32(1557), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
						v42 = v41
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
						v47 = int32(0)
						v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
								v99 = v51
								v100 = v49
								if v99 == int32(0) {
									v118 = int32(0)
									m.G0 = v8 - int32(-64)
									return v118
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
										v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
										if v112 == int32(0) {
											v118 = v107
										} else {
											base.MemoryCopy(m, v107, v109, v112)
											v118 = v107
										}
										m.G0 = v8 - int32(-64)
										return v118
									}
								}
							} else {
								v52 = int32(0)
								v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									if v55 != 0 {
										v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
											if v63 != 0 {
												v81 = v52
												F_ReleaseCatCache(m, v55)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v81
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
													v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
														v99 = v86
														v100 = v94
														if v99 == int32(0) {
															v118 = int32(0)
															m.G0 = v8 - int32(-64)
															return v118
														} else {
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
															v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return int32(0)
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
																v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
																if v112 == int32(0) {
																	v118 = v107
																} else {
																	base.MemoryCopy(m, v107, v109, v112)
																	v118 = v107
																}
																m.G0 = v8 - int32(-64)
																return v118
															}
														}
													}
												}
											} else {
												v65 = F_attribute_reloptions(m, v61, int32(0))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
													v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
														v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
														if v76 == int32(0) {
															v81 = v72
														} else {
															base.MemoryCopy(m, v72, v65, v76)
															v81 = v72
														}
														F_ReleaseCatCache(m, v55)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v86 = v81
															v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
															v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
																v99 = v86
																v100 = v94
																if v99 == int32(0) {
																	v118 = int32(0)
																	m.G0 = v8 - int32(-64)
																	return v118
																} else {
																	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
																	v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
																		v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
																		v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
																		if v112 == int32(0) {
																			v118 = v107
																		} else {
																			base.MemoryCopy(m, v107, v109, v112)
																			v118 = v107
																		}
																		m.G0 = v8 - int32(-64)
																		return v118
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v86 = v52
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
										v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
											v99 = v86
											v100 = v94
											if v99 == int32(0) {
												v118 = int32(0)
												m.G0 = v8 - int32(-64)
												return v118
											} else {
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
												v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
													v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
													if v112 == int32(0) {
														v118 = v107
													} else {
														base.MemoryCopy(m, v107, v109, v112)
														v118 = v107
													}
													m.G0 = v8 - int32(-64)
													return v118
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(7), int32(1557), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
					v42 = v41
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v47 = int32(0)
					v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
							v99 = v51
							v100 = v49
							if v99 == int32(0) {
								v118 = int32(0)
								m.G0 = v8 - int32(-64)
								return v118
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
									v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
									if v112 == int32(0) {
										v118 = v107
									} else {
										base.MemoryCopy(m, v107, v109, v112)
										v118 = v107
									}
									m.G0 = v8 - int32(-64)
									return v118
								}
							}
						} else {
							v52 = int32(0)
							v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 != 0 {
									v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
										if v63 != 0 {
											v81 = v52
											F_ReleaseCatCache(m, v55)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v81
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
												v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
													v99 = v86
													v100 = v94
													if v99 == int32(0) {
														v118 = int32(0)
														m.G0 = v8 - int32(-64)
														return v118
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
														v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int32(0)
														} else {
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
															v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
															if v112 == int32(0) {
																v118 = v107
															} else {
																base.MemoryCopy(m, v107, v109, v112)
																v118 = v107
															}
															m.G0 = v8 - int32(-64)
															return v118
														}
													}
												}
											}
										} else {
											v65 = F_attribute_reloptions(m, v61, int32(0))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
												v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
													v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
													if v76 == int32(0) {
														v81 = v72
													} else {
														base.MemoryCopy(m, v72, v65, v76)
														v81 = v72
													}
													F_ReleaseCatCache(m, v55)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = v81
														v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
														v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
															v99 = v86
															v100 = v94
															if v99 == int32(0) {
																v118 = int32(0)
																m.G0 = v8 - int32(-64)
																return v118
															} else {
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
																v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
																	v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
																	if v112 == int32(0) {
																		v118 = v107
																	} else {
																		base.MemoryCopy(m, v107, v109, v112)
																		v118 = v107
																	}
																	m.G0 = v8 - int32(-64)
																	return v118
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v86 = v52
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
									v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
										v99 = v86
										v100 = v94
										if v99 == int32(0) {
											v118 = int32(0)
											m.G0 = v8 - int32(-64)
											return v118
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
												v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
												if v112 == int32(0) {
													v118 = v107
												} else {
													base.MemoryCopy(m, v107, v109, v112)
													v118 = v107
												}
												m.G0 = v8 - int32(-64)
												return v118
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
	} else {
		v42 = v11
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
		v47 = int32(0)
		v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			if v49 != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
				v99 = v51
				v100 = v49
				if v99 == int32(0) {
					v118 = int32(0)
					m.G0 = v8 - int32(-64)
					return v118
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
					v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
						if v112 == int32(0) {
							v118 = v107
						} else {
							base.MemoryCopy(m, v107, v109, v112)
							v118 = v107
						}
						m.G0 = v8 - int32(-64)
						return v118
					}
				}
			} else {
				v52 = int32(0)
				v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 != 0 {
						v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v63 != 0 {
								v81 = v52
								F_ReleaseCatCache(m, v55)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v86 = v81
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
									v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
										v99 = v86
										v100 = v94
										if v99 == int32(0) {
											v118 = int32(0)
											m.G0 = v8 - int32(-64)
											return v118
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
												v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
												if v112 == int32(0) {
													v118 = v107
												} else {
													base.MemoryCopy(m, v107, v109, v112)
													v118 = v107
												}
												m.G0 = v8 - int32(-64)
												return v118
											}
										}
									}
								}
							} else {
								v65 = F_attribute_reloptions(m, v61, int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
									v72 = F_MemoryContextAlloc(m, v68, int32(base.Ui32(v69)>>(uint(int32(2))%32)))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										v76 = int32(base.Ui32(v74) >> (uint(int32(2)) % 32))
										if v76 == int32(0) {
											v81 = v72
										} else {
											base.MemoryCopy(m, v72, v65, v76)
											v81 = v72
										}
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v81
											v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
											v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
												v99 = v86
												v100 = v94
												if v99 == int32(0) {
													v118 = int32(0)
													m.G0 = v8 - int32(-64)
													return v118
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
													v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return int32(0)
													} else {
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
														v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
														if v112 == int32(0) {
															v118 = v107
														} else {
															base.MemoryCopy(m, v107, v109, v112)
															v118 = v107
														}
														m.G0 = v8 - int32(-64)
														return v118
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v86 = v52
						v89 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
						v94 = F_hash_search(m, v89, v6+int32(-48), int32(1), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v86
							v99 = v86
							v100 = v94
							if v99 == int32(0) {
								v118 = int32(0)
								m.G0 = v8 - int32(-64)
								return v118
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								v107 = F_palloc(m, int32(base.Ui32(v104)>>(uint(int32(2))%32)))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
									v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
									if v112 == int32(0) {
										v118 = v107
									} else {
										base.MemoryCopy(m, v107, v109, v112)
										v118 = v107
									}
									m.G0 = v8 - int32(-64)
									return v118
								}
							}
						}
					}
				}
			}
		}
	}
}
