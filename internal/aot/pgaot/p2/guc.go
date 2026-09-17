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
	var v51 int32
	_ = v51
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
	var v111 int32
	_ = v111
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
	var v171 int32
	_ = v171
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
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v218 int32
	_ = v218
	var v222 int64
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(_a_F_InitializeGUCOptionsFromEnvironment_0)
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
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_InitializeGUCOptionsFromEnvironment[0]))))
	if v20 != 0 {
		v51 = v1
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v57 = v51
	goto L1
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptionsFromEnvironment[1]))
	if v22 == int32(0) {
		v51 = v1
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v25 == int32(0) {
		v51 = v1
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
	v51 = v37 + int32(1)
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
	v51 = v1
	goto L5
L17:
	;
	v59 = int32(0)
	v60 = int32(1)
	v67 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptionsFromEnvironment_1), v59, v57, v60, int32(2), int32(10), v59, v60, v59, v59)
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
	v69 = int32(_a_F_InitializeGUCOptionsFromEnvironment_2)
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
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_InitializeGUCOptionsFromEnvironment[2]))))
	if v80 != 0 {
		v111 = v70
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v117 = v111
	goto L22
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptionsFromEnvironment[1]))
	if v82 == int32(0) {
		v111 = v70
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v85 == int32(0) {
		v111 = v70
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
	v111 = v97 + int32(1)
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
	v111 = v70
	goto L26
L38:
	;
	v119 = int32(0)
	v120 = int32(1)
	v127 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptionsFromEnvironment_3), v119, v117, v120, int32(2), int32(10), v119, v120, v119, v119)
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
	v129 = int32(_a_F_InitializeGUCOptionsFromEnvironment_4)
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
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_InitializeGUCOptionsFromEnvironment[3]))))
	if v140 != 0 {
		v171 = v130
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v177 = v171
	goto L42
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptionsFromEnvironment[1]))
	if v142 == int32(0) {
		v171 = v130
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v145 == int32(0) {
		v171 = v130
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
	v171 = v157 + int32(1)
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
	v171 = v130
	goto L46
L58:
	;
	v179 = int32(0)
	v180 = int32(1)
	v187 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptionsFromEnvironment_5), v179, v177, v180, int32(2), int32(10), v179, v180, v179, v179)
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
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptionsFromEnvironment[4]))
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
	if v191 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v231 = v194
	goto L64
L64:
	;
	m.G0 = v191 + int32(16)
	if int32(_a_F_InitializeGUCOptionsFromEnvironment_6) <= v231 {
		goto L79
	} else {
		goto L80
	}
L65:
	;
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	if base.Ui64(int64(2147483646)) < base.Ui64(v222) {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	switch int32(0) {
	case 0:
		goto L71
	default:
		goto L70
	case 4:
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v218 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L65
L69:
	;
	goto L68
L70:
	;
	v211 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v211
	goto L69
L71:
	;
	v205 = m.G2
	v206 = m.G1
	v208 = base.I64_extend_i32_u(v205 - v206)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v208
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v208
	goto L69
L72:
	;
	v201 = int64(4096)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v201
	goto L69
L73:
	;
	v226 = int32(2147483647)
	goto L75
L74:
	;
	v226 = base.I32_wrap_i64(v222)
	goto L75
L75:
	;
	if v218 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v229 = int32(-1)
	goto L78
L77:
	;
	v229 = v226
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptionsFromEnvironment[4])) = v229
	v231 = v229
	goto L64
L79:
	;
	v244 = base.B2i32(base.Ui32(v231) < base.Ui32(int32(_a_F_InitializeGUCOptionsFromEnvironment_7)))
	if base.Ui32(v231) < base.Ui32(int32(_a_F_InitializeGUCOptionsFromEnvironment_7)) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	m.G0 = v7 + int32(32)
	return
L82:
	;
	v245 = int32(base.Ui32(v231-int32(_a_F_InitializeGUCOptionsFromEnvironment_8)) >> (uint(int32(10)) % 32))
	goto L84
L83:
	;
	v245 = int32(2048)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v245
	v247 = int32(16)
	v248 = v7 + v247
	v251 = F_pg_snprintf(m, v248, v247, int32(_a_F_InitializeGUCOptionsFromEnvironment_9), v7)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	v255 = int32(1)
	if base.Ui32(v231) < base.Ui32(int32(_a_F_InitializeGUCOptionsFromEnvironment_7)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v258 = int32(2)
	goto L88
L87:
	;
	v258 = v255
	goto L88
L88:
	;
	v260 = int32(0)
	v264 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptionsFromEnvironment_10), int32(0), v248, v255, v258, int32(10), v260, int32(1), v260, v260)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L20
	} else {
		goto L89
	}
L89:
	;
	goto L81
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
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
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = F_strlen(m, l0)
	mBase = m.M
	v12 = v8 + int32(28)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[0]))
	F_hash_seq_init(m, v12, v14)
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
	v17 = F_hash_seq_search(m, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v17
	goto L7
L5:
	;
	goto L6
L6:
	;
	v158 = int32(_a_F_MarkGUCPrefixReserved_0)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[1]))
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[1])) = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[3]))
	v166 = F_pstrdup(m, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L56
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	if v25&int32(2) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v151 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v10 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v75 != 0 {
		goto L9
	} else {
		goto L24
	}
L12:
	;
	v75 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = l0
	v38 = v30
	v39 = v10
	v40 = v36
	goto L19
L16:
	;
	v63 = v30
	v67 = int32(0)
	goto L17
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v75 = v67 - v68
	goto L11
L18:
	;
	v63 = v58
	v67 = v60
	goto L17
L19:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if base.B2i32(v40 != v42)|base.B2i32(v42 == int32(0)) != 0 {
		v58 = v38
		v60 = v40
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v58 = v52
	v60 = int32(0)
	goto L18
L21:
	;
	v48 = v39 - int32(1)
	if v48 == int32(0) {
		v58 = v38
		v60 = v40
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v51 = int32(1)
	v52 = v38 + v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v53 != 0 {
		v37 = v37 + v51
		v38 = v52
		v39 = v48
		v40 = v53
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v10))))
	if v77 != int32(46) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v82 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[0]))
	v107 = F_hash_search(m, v104, v24, int32(2), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v87
	F_errmsg(m, int32(_a_F_MarkGUCPrefixReserved_1), v8+int32(16))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errdetail(m, int32(_a_F_MarkGUCPrefixReserved_2), v8)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_MarkGUCPrefixReserved_3), int32(_a_F_MarkGUCPrefixReserved_4), int32(_a_F_MarkGUCPrefixReserved_5))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	if v109 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v113
	goto L37
L36:
	;
	goto L37
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	if v116 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v121 = int32(_a_F_MarkGUCPrefixReserved_6)
	goto L43
L39:
	;
	goto L40
L40:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+28)))
	if v130&int32(4) == int32(0) {
		goto L9
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	goto L41
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v124 == int32(0) {
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v128
	goto L42
L45:
	;
	if v124 != v24+int32(72) {
		v121 = v124
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v139 = int32(_a_F_MarkGUCPrefixReserved_7)
	goto L50
L48:
	;
	goto L9
L49:
	;
	goto L48
L50:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v142 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v146
	goto L49
L52:
	;
	if v142 != v24+int32(76) {
		v139 = v142
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v151 != 0 {
		v20 = v151
		goto L7
	} else {
		goto L55
	}
L55:
	;
	goto L8
L56:
	;
	v168 = F_lappend(m, v165, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[1])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_MarkGUCPrefixReserved[3])) = v168
	m.G0 = v8 + int32(48)
	return
}
