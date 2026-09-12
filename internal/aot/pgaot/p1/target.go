package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_target_list(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_initStringInfo(m, v15+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	F_pfree(m, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L76
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(746047)
	v34 = v3
	v37 = v3
	v39 = v3
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+26)))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoString(m, v17, v32)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v276 = v32
	v278 = v34
	v283 = v39
	goto L10
L10:
	;
	v285 = v37 + int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v285 < v286 {
		v32 = v276
		v34 = v278
		v37 = v285
		v39 = v283
		goto L6
	} else {
		goto L75
	}
L11:
	;
	v51 = v15 + int32(16)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v53
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15 + int32(16)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v62 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v80 = v34 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v81 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	F_get_rule_expr(m, v62, l1, int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v65 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v69 = F_get_variable(m, v62, int32(1), l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v78 = v69
	goto L13
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+33)))
	if v76 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v77 = int32(0)
	goto L21
L20:
	;
	v77 = int32(545471)
	goto L21
L21:
	;
	v78 = v77
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v138&int32(2) == int32(0) {
		v266 = v39
		goto L42
	} else {
		goto L43
	}
L23:
	;
	if v78 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v95 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v84 < v80 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v98 = v81 + v84<<(uint(int32(4))%32) + v34*int32(100) + int32(24)
	goto L23
L27:
	;
	v98 = v95
	goto L23
L28:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v103 == int32(0) {
		v122 = v102
		v123 = v103
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v127 = F_quote_identifier(m, v98)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L31:
	;
	if v123-v122 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	if v102 != v103 {
		v122 = v102
		v123 = v103
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v107 = v78
	v108 = v98
	goto L35
L35:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v111
		v123 = v112
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v122 = v111
	v123 = v112
	goto L32
L37:
	;
	v115 = int32(1)
	if v111 == v112 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v127
	F_appendStringInfo(m, v15+int32(16), int32(197946), v15)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L22
L42:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_appendBinaryStringInfo(m, v17, v267, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L74
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v143 < int32(0) {
		v266 = v39
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v146 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v245 = int32(10)
	v246 = F___strchrnul(m, v241+v233+int32(1), v245)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v248 == v245 {
		goto L71
	} else {
		goto L72
	}
L46:
	;
	v187 = int32(-1)
	if v80 < int32(2) {
		v233 = v187
		goto L45
	} else {
		goto L58
	}
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 != int32(10) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v153 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v233 = int32(0)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v159 = v153
	goto L52
L52:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v159-int32(1)))))
	if v173 == int32(32) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v233 = int32(0)
	goto L45
L54:
	;
	v177 = v159 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v177
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v177+v169))) = uint8(v179)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v179 < v183 {
		v159 = v183
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	v233 = v179
	goto L45
L58:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v194 = F_strlen(m, v190)
	mBase = m.M
	v201 = v194 + int32(1)
	goto L61
L59:
	;
	if v213 != 0 {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v203 = int32(0)
	if v201 == v203 {
		v213 = v203
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v213 = v208
	goto L60
L63:
	;
	v207 = v201 - int32(1)
	v208 = v190 + v207
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v209 != int32(10) {
		v201 = v207
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v216 = v213 + int32(1)
	goto L67
L66:
	;
	v216 = v190
	goto L67
L67:
	;
	v217 = F_strlen(m, v216)
	mBase = m.M
	if v39|base.B2i32(base.Ui32(v143) < base.Ui32(v217+v146)) == int32(0) {
		v233 = v187
		goto L45
	} else {
		goto L68
	}
L68:
	;
	F_appendContextKeyword(m, l1, int32(757108), int32(-8), int32(8), int32(4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v233 = v187
	goto L45
L70:
	;
	v266 = base.B2i32(v252 != int32(0))
	goto L42
L71:
	;
	v252 = v246
	goto L73
L72:
	;
	v252 = int32(0)
	goto L73
L73:
	;
	goto L70
L74:
	;
	v276 = int32(745866)
	v278 = v80
	v283 = v266
	goto L10
L75:
	;
	goto L7
L76:
	;
	m.G0 = v15 + int32(32)
	return
}
func F_makeTargetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = l1
	v4 = l3
	v7 = F_palloc0(m, int32(28))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)) = uint8(v4)
		v12 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+24)) = uint16(v12)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(62)
		return v7
	}
}
func F_markTargetListOrigins(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L47
	} else {
		goto L83
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L47
	} else {
		goto L80
	}
L3:
	;
	m.G0 = v13 + int32(32)
	return
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = v3
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v25<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v253 = v25 + int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v253 < v254 {
		v25 = v253
		goto L6
	} else {
		goto L79
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v38 != int32(6) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v43 = int32(0)
	if v42 <= v43 {
		v92 = l0
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+8)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	switch v107 {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L8
	case 6:
		goto L26
	}
L12:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+v41<<(uint(int32(2))%32)-int32(4))))
	goto L11
L13:
	;
	v49 = v42 & int32(7)
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if base.Ui32(v42) < base.Ui32(int32(8)) {
		v92 = v64
		goto L12
	} else {
		goto L21
	}
L15:
	;
	v64 = l0
	v67 = v42
	goto L14
L16:
	;
	goto L17
L17:
	;
	v52 = l0
	v55 = v42
	v56 = v43
	goto L18
L18:
	;
	v58 = int32(1)
	v59 = v55 - v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v62 = v56 + v58
	if v62 != v49 {
		v52 = v60
		v55 = v59
		v56 = v62
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v64 = v60
	v67 = v59
	goto L14
L20:
	;
	goto L19
L21:
	;
	v72 = v64
	v75 = v67
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if base.Ui32(v75-int32(9)) < base.Ui32(int32(-2)) {
		v72 = v87
		v75 = v75 - int32(8)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v92 = v87
	goto L12
L24:
	;
	goto L23
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)) = uint16(v242)
	goto L8
L26:
	;
	if v106 == int32(0) {
		goto L8
	} else {
		goto L45
	}
L27:
	;
	if v106 == int32(0) {
		goto L8
	} else {
		goto L29
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v108
	v242 = v106
	goto L25
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+36))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+76))
	if v113 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v151 == int32(0) {
		goto L2
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v117 <= int32(0) {
		v151 = int32(0)
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v151 = int32(0)
	goto L31
L35:
	;
	v120 = int32(0)
	if v120 < v117 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = v117
	goto L38
L37:
	;
	v123 = v120
	goto L38
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v128 = int32(0)
	goto L39
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124+v128<<(uint(int32(2))%32))))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+8)))
	if v137 == v106&int32(65535) {
		v151 = v136
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v140 = v128 + int32(1)
	if v140 != v123 {
		v128 = v140
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+26)))
	if v155 == int32(1) {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v158
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+24)))
	v242 = v160
	goto L25
L45:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+92)))
	if v163 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v164 = F_GetCTEForRTE(m, l0, v105, v42)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v169 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v172 = int32(76)
	goto L51
L50:
	;
	v172 = int32(96)
	goto L51
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166+v172)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v177 = base.B2i32(v175 != int32(0))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v164)+24))
	if v180 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v174 != 0 {
		goto L66
	} else {
		goto L67
	}
L53:
	;
	v181 = v177 | int32(2)
	goto L55
L54:
	;
	v181 = v177
	goto L55
L55:
	;
	if v181 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if v174 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v106 <= v181+v191 {
		goto L8
	} else {
		goto L63
	}
L58:
	;
	v186 = int32(0)
	if v186 < v106 {
		v191 = v186
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v106 <= v189 {
		goto L52
	} else {
		goto L62
	}
L61:
	;
	goto L52
L62:
	;
	v191 = v189
	goto L57
L63:
	;
	goto L52
L64:
	;
	if v232 == int32(0) {
		goto L1
	} else {
		goto L77
	}
L65:
	;
	goto L64
L66:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v198 <= int32(0) {
		v232 = int32(0)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v232 = int32(0)
	goto L65
L69:
	;
	v201 = int32(0)
	if v201 < v198 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v204 = v198
	goto L72
L71:
	;
	v204 = v201
	goto L72
L72:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v209 = int32(0)
	goto L73
L73:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v205+v209<<(uint(int32(2))%32))))
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+8)))
	if v218 == v106&int32(65535) {
		v232 = v217
		goto L65
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v221 = v209 + int32(1)
	if v221 != v204 {
		v209 = v221
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+26)))
	if v236 == int32(1) {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v239
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+24)))
	v242 = v241
	goto L25
L79:
	;
	goto L7
L80:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v274
	F_errmsg_internal(m, int32(474944), v13)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(493195), int32(372), int32(276684))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v290
	F_errmsg_internal(m, int32(474983), v13+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L47
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(493195), int32(418), int32(276684))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setTargetTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v16 = F_name_matches_visible_ENR(m, l0, v15)
		mBase = m.M
		if v16 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v52
					F_errmsg(m, int32(95607), v10)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498506), int32(192), int32(396894))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v17 != 0 {
				F_sequence_close(m, v17, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v24 = F_parserOpenTable(m, l0, l1, int32(3))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v24
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v30 = F_addRangeTableEntryForRelation(m, l0, v24, int32(3), v28, l2, int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v30
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
							*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = l4
							if l3 != 0 {
								v35 = int32(1)
								F_addNSItemToQuery(m, l0, v30, v35, v35, v35)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
									m.G0 = v10 + int32(16)
									return v40
								}
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
								m.G0 = v10 + int32(16)
								return v40
							}
						}
					}
				}
			} else {
				v24 = F_parserOpenTable(m, l0, l1, int32(3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v24
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v30 = F_addRangeTableEntryForRelation(m, l0, v24, int32(3), v28, l2, int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = l4
						if l3 != 0 {
							v35 = int32(1)
							F_addNSItemToQuery(m, l0, v30, v35, v35, v35)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
								m.G0 = v10 + int32(16)
								return v40
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							m.G0 = v10 + int32(16)
							return v40
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v17 != 0 {
			F_sequence_close(m, v17, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = F_parserOpenTable(m, l0, l1, int32(3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v24
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v30 = F_addRangeTableEntryForRelation(m, l0, v24, int32(3), v28, l2, int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
						*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = l4
						if l3 != 0 {
							v35 = int32(1)
							F_addNSItemToQuery(m, l0, v30, v35, v35, v35)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
								m.G0 = v10 + int32(16)
								return v40
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							m.G0 = v10 + int32(16)
							return v40
						}
					}
				}
			}
		} else {
			v24 = F_parserOpenTable(m, l0, l1, int32(3))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v24
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v30 = F_addRangeTableEntryForRelation(m, l0, v24, int32(3), v28, l2, int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v30
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
					*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = l4
					if l3 != 0 {
						v35 = int32(1)
						F_addNSItemToQuery(m, l0, v30, v35, v35, v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							m.G0 = v10 + int32(16)
							return v40
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						m.G0 = v10 + int32(16)
						return v40
					}
				}
			}
		}
	}
}
