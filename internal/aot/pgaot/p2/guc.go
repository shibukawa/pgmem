package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeGUCOptionsFromEnvironment(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v259 int64
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(514029)
	v15 = F___strchrnul(m, v9, int32(61))
	mBase = m.M
	if v9 == v15 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v57 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v57 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = v15 - v9
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1229]))))
	if v20 != 0 {
		v50 = v1
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v57 = v50
	goto L1
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1230]))
	if v22 == int32(0) {
		v50 = v1
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v25 == int32(0) {
		v50 = v1
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v29 = v22
	v30 = v25
	goto L9
L9:
	;
	v33 = F_strncmp(m, v9, v30, v18)
	mBase = m.M
	if v33 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v50 = v37 + int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v37 = v36 + v18
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v38 == int32(61) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v42 != 0 {
		v29 = v29 + int32(4)
		v30 = v42
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v50 = v1
	goto L5
L17:
	;
	v59 = int32(0)
	v60 = int32(1)
	v67 = F_set_config_with_handle(m, int32(80658), v59, v57, v60, int32(2), int32(10), v59, v60, v59, v59)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v69 = int32(536708)
	v70 = int32(0)
	v75 = F___strchrnul(m, v69, int32(61))
	mBase = m.M
	if v69 == v75 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	return
L21:
	;
	goto L19
L22:
	;
	if v117 != 0 {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	v117 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v78 = v75 - v69
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_consts[1231]))))
	if v80 != 0 {
		v110 = v70
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v117 = v110
	goto L22
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[1230]))
	if v82 == int32(0) {
		v110 = v70
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v85 == int32(0) {
		v110 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v89 = v82
	v90 = v85
	goto L30
L30:
	;
	v93 = F_strncmp(m, v69, v90, v78)
	mBase = m.M
	if v93 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v110 = v97 + int32(1)
	goto L26
L32:
	;
	goto L31
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v97 = v96 + v78
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v98 == int32(61) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v102 != 0 {
		v89 = v89 + int32(4)
		v90 = v102
		goto L30
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v110 = v70
	goto L26
L38:
	;
	v119 = int32(0)
	v120 = int32(1)
	v127 = F_set_config_with_handle(m, int32(380124), v119, v117, v120, int32(2), int32(10), v119, v120, v119, v119)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v129 = int32(533294)
	v130 = int32(0)
	v135 = F___strchrnul(m, v129, int32(61))
	mBase = m.M
	if v129 == v135 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	if v177 != 0 {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	v177 = int32(0)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v138 = v135 - v129
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[1232]))))
	if v140 != 0 {
		v170 = v130
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v177 = v170
	goto L42
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[1230]))
	if v142 == int32(0) {
		v170 = v130
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v145 == int32(0) {
		v170 = v130
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v149 = v142
	v150 = v145
	goto L50
L50:
	;
	v153 = F_strncmp(m, v129, v150, v138)
	mBase = m.M
	if v153 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v170 = v157 + int32(1)
	goto L46
L52:
	;
	goto L51
L53:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v157 = v156 + v138
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v158 == int32(61) {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v162 != 0 {
		v149 = v149 + int32(4)
		v150 = v162
		goto L50
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v170 = v130
	goto L46
L58:
	;
	v179 = int32(0)
	v180 = int32(1)
	v187 = F_set_config_with_handle(m, int32(333774), v179, v177, v180, int32(2), int32(10), v179, v180, v179, v179)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v189 = m.G0
	v191 = v189 - int32(16)
	m.G0 = v191
	v194 = *(*int32)(unsafe.Add(mBase, _consts[1233]))
	if v194 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v201 = m.G0
	v203 = v201 - int32(16)
	m.G0 = v203
	if v191 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v268 = v194
	goto L64
L64:
	;
	m.G0 = v191 + int32(16)
	if int32(627712) <= v268 {
		goto L91
	} else {
		goto L92
	}
L65:
	;
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	if base.Ui64(int64(2147483646)) < base.Ui64(v259) {
		goto L85
	} else {
		goto L86
	}
L66:
	;
	v205 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v205
	goto L68
L67:
	;
	goto L68
L68:
	;
	v209 = int32(0)
	v210 = F___syscall_ret(m, v209)
	mBase = m.M
	if v210 == v209 {
		v251 = int32(0)
		goto L69
	} else {
		goto L70
	}
L69:
	;
	m.G0 = v203 + int32(16)
	goto L65
L70:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v214 != int32(52) {
		v251 = v210
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v218 = v203 + int32(8)
	v219 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v218))) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v218)+8)) = v219
	v223 = int32(0)
	v224 = F___syscall_ret(m, v223)
	mBase = m.M
	if v224 < v223 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v251 = int32(-1)
	goto L69
L73:
	;
	goto L74
L74:
	;
	v229 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v203)+8)))
	if v229 == int64(4294967295) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v232 = int64(-1)
	goto L77
L76:
	;
	v232 = v229
	goto L77
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	if v235 == int32(-1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v239 = int64(-1)
	goto L80
L79:
	;
	v239 = base.I64_extend_i32_u(v235)
	goto L80
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v239
	if v229 == int64(4294967295) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = int64(-1)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v245 = int32(0)
	if v235 != int32(-1) {
		v251 = v245
		goto L69
	} else {
		goto L84
	}
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = int64(-1)
	v251 = v245
	goto L69
L85:
	;
	v263 = int32(2147483647)
	goto L87
L86:
	;
	v263 = base.I32_wrap_i64(v259)
	goto L87
L87:
	;
	if v251 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v266 = int32(-1)
	goto L90
L89:
	;
	v266 = v263
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1233])) = v266
	v268 = v266
	goto L64
L91:
	;
	v281 = base.B2i32(base.Ui32(v268) < base.Ui32(int32(2621440)))
	if base.Ui32(v268) < base.Ui32(int32(2621440)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v7 + int32(32)
	return
L94:
	;
	v282 = int32(base.Ui32(v268-int32(524288)) >> (uint(int32(10)) % 32))
	goto L96
L95:
	;
	v282 = int32(2048)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v282
	v284 = int32(16)
	v288 = F_pg_snprintf(m, v7+v284, v284, int32(485191), v7)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v294 = int32(1)
	if base.Ui32(v268) < base.Ui32(int32(2621440)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v297 = int32(2)
	goto L100
L99:
	;
	v297 = v294
	goto L100
L100:
	;
	v299 = int32(0)
	v303 = F_set_config_with_handle(m, int32(318562), int32(0), v7+int32(16), v294, v297, int32(10), v299, int32(1), v299, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	goto L93
}
func F_MarkGUCPrefixReserved(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = F_strlen(m, l0)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1225]))
	F_hash_seq_init(m, v8+int32(28), v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v19
	goto L7
L5:
	;
	goto L6
L6:
	;
	v159 = int32(4487040)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v163 = *(*int32)(unsafe.Add(mBase, _consts[1235]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v163
	v166 = *(*int32)(unsafe.Add(mBase, _consts[1234]))
	v167 = F_pstrdup(m, l0)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)))
	if v27&int32(2) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v152 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v10 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v76 != 0 {
		goto L9
	} else {
		goto L25
	}
L12:
	;
	v76 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = l0
	v40 = v32
	v41 = v10
	v42 = v38
	goto L19
L16:
	;
	v64 = v32
	v68 = int32(0)
	goto L17
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v76 = v68 - v69
	goto L11
L18:
	;
	v64 = v59
	v68 = v61
	goto L17
L19:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 != v44 {
		v59 = v40
		v61 = v42
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v59 = v53
	v61 = int32(0)
	goto L18
L21:
	;
	if v44 == int32(0) {
		v59 = v40
		v61 = v42
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v49 = v41 - int32(1)
	if v49 == int32(0) {
		v59 = v40
		v61 = v42
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v52 = int32(1)
	v53 = v40 + v52
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v54 != 0 {
		v39 = v39 + v52
		v40 = v53
		v41 = v49
		v42 = v54
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v32))))
	if v78 != int32(46) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v83 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[1225]))
	v108 = F_hash_search(m, v105, v26, int32(2), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v88
	F_errmsg(m, int32(103576), v8+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errdetail(m, int32(558111), v8)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(496606), int32(5314), int32(437598))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v110 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v114
	goto L38
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	if v117 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = int32(4485084)
	goto L44
L40:
	;
	goto L41
L41:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)))
	if v131&int32(4) == int32(0) {
		goto L9
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	goto L42
L44:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v125 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v129
	goto L43
L46:
	;
	if v125 != v26+int32(72) {
		v122 = v125
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v140 = int32(4485076)
	goto L51
L49:
	;
	goto L9
L50:
	;
	goto L49
L51:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v143 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v147
	goto L50
L53:
	;
	if v143 != v26+int32(76) {
		v140 = v143
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v152 != 0 {
		v22 = v152
		goto L7
	} else {
		goto L56
	}
L56:
	;
	goto L8
L57:
	;
	v169 = F_lappend(m, v166, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v160
	*(*int32)(unsafe.Add(mBase, _consts[1234])) = v169
	m.G0 = v8 + int32(48)
	return
}
