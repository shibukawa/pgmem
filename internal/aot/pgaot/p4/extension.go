package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v50
L2:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v39 != v40 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v50 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v50 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = F_equal(m, v44, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v50 = v46
	goto L1
}
func F_get_extension_control_directories(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	v10 = m.G0
	v12 = v10 - int32(1056)
	m.G0 = v12
	v15 = v12 + int32(32)
	F_get_share_path(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
	v24 = F_psprintf(m, int32(_a_F_get_extension_control_directories_0), v12+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_get_extension_control_directories[0]))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v12 + int32(1056)
	return v240
L5:
	;
	v32 = F_palloc(m, int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v42 = F_pstrdup(m, v27)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v35 = F_pstrdup(m, int32(_a_F_get_extension_control_directories_1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v35
	v40 = F_lappend(m, int32(0), v32)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v240 = v40
	goto L4
L11:
	;
	v44 = v42
	v48 = int32(0)
	goto L12
L12:
	;
	v54 = Fn14265(m, v44, int32(58))
	mBase = m.M
	goto L14
L13:
	;
	v240 = v232
	goto L4
L14:
	;
	v56 = F_palloc(m, int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v54 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v64 = v62 + int32(1)
	v65 = F_palloc(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v60 = F_strlen(m, v44)
	mBase = m.M
	v62 = v60
	goto L16
L18:
	;
	goto L19
L19:
	;
	v62 = v54 - v44
	goto L16
L20:
	;
	if v64 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v185 = int32(_a_F_get_extension_control_directories_1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_extension_control_directories[1])))
	if base.B2i32(v188 == int32(0))|base.B2i32(v188 != v191) != 0 {
		v209 = v188
		v210 = v191
		goto L54
	} else {
		goto L55
	}
L22:
	;
	v182 = F_strlen(m, v178)
	mBase = m.M
	goto L21
L23:
	;
	v178 = v44
	goto L22
L24:
	;
	goto L25
L25:
	;
	v72 = v64 - int32(1)
	if (v65^v44)&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v175)
	v178 = v171
	goto L22
L27:
	;
	v156 = v151
	v157 = v152
	v158 = v153
	goto L48
L28:
	;
	if v146 == int32(0) {
		v171 = v144
		v172 = v145
		goto L26
	} else {
		goto L47
	}
L29:
	;
	v144 = v44
	v145 = v65
	v146 = v72
	goto L28
L30:
	;
	goto L31
L31:
	;
	v76 = int32(0)
	if base.B2i32(v44&int32(3) == v76)|base.B2i32(v72 == v76) == v76 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v112 == int32(0) {
		v171 = v109
		v172 = v110
		goto L26
	} else {
		goto L41
	}
L33:
	;
	v88 = v44
	v89 = v65
	v90 = v72
	goto L36
L34:
	;
	goto L35
L35:
	;
	v109 = v44
	v110 = v65
	v111 = v72
	v112 = base.B2i32(v72 != v76)
	goto L32
L36:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v92)
	if v92 == int32(0) {
		v151 = v88
		v152 = v89
		v153 = v90
		goto L27
	} else {
		goto L38
	}
L37:
	;
	v109 = v103
	v110 = v97
	v111 = v99
	v112 = v101
	goto L32
L38:
	;
	v96 = int32(1)
	v97 = v89 + v96
	v99 = v90 - v96
	v100 = int32(0)
	v101 = base.B2i32(v99 != v100)
	v103 = v88 + v96
	if v103&int32(3) == v100 {
		v109 = v103
		v110 = v97
		v111 = v99
		v112 = v101
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v99 != 0 {
		v88 = v103
		v89 = v97
		v90 = v99
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.B2i32(v115 == int32(0))|base.B2i32(base.Ui32(v111) < base.Ui32(int32(4))) != 0 {
		v144 = v109
		v145 = v110
		v146 = v111
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v122 = v109
	v123 = v110
	v124 = v111
	goto L43
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v130 = int32(-2139062144)
	if (int32(16843008)-v127|v127)&v130 != v130 {
		v151 = v122
		v152 = v123
		v153 = v124
		goto L27
	} else {
		goto L45
	}
L44:
	;
	v144 = v138
	v145 = v136
	v146 = v140
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v127
	v135 = int32(4)
	v136 = v123 + v135
	v138 = v122 + v135
	v140 = v124 - v135
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		v122 = v138
		v123 = v136
		v124 = v140
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v151 = v144
	v152 = v145
	v153 = v146
	goto L27
L48:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v160)
	if v160 == int32(0) {
		v171 = v156
		v172 = v157
		goto L26
	} else {
		goto L50
	}
L49:
	;
	v171 = v167
	v172 = v165
	goto L26
L50:
	;
	v164 = int32(1)
	v165 = v157 + v164
	v167 = v156 + v164
	v169 = v158 - v164
	if v169 != 0 {
		v156 = v167
		v157 = v165
		v158 = v169
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	F_pfree(m, v65)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L66
	}
L53:
	;
	if v209-v210 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v194 = v65
	v195 = v185
	goto L56
L56:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v199 == int32(0) {
		v209 = v199
		v210 = v198
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v209 = v199
	v210 = v198
	goto L54
L58:
	;
	v202 = int32(1)
	if v199 == v198 {
		v194 = v194 + v202
		v195 = v195 + v202
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v214 = F_pstrdup(m, v65)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v65
	v224 = F_psprintf(m, int32(_a_F_get_extension_control_directories_0), v12)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v214
	v218 = F_substitute_path_macro(m, v65, int32(_a_F_get_extension_control_directories_1), v24)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v226 = v218
	goto L52
L65:
	;
	v226 = v224
	goto L52
L66:
	;
	F_canonicalize_path_enc(m, v226)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v226
	v232 = F_lappend(m, v48, v56)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v62))))
	if v235 != 0 {
		v44 = v44 + v64
		v48 = v232
		goto L12
	} else {
		goto L68
	}
L68:
	;
	goto L13
}
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v868 int32
	_ = v868
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(240)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v3
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L12
	} else {
		goto L265
	}
L2:
	;
	v202 = F_strlen(m, v193)
	mBase = m.M
	v205 = F_pnstrdup(m, v193, v202-int32(10))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L56
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v56
	v77 = F_psprintf(m, int32(_a_F_parse_extension_control_file_0), v14+int32(208))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L26
	}
L4:
	;
	if v69 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L8:
	;
	v42 = F_palloc(m, int32(1024))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L19
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(47) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v40 = v24
	goto L8
L14:
	;
	v29 = F_pstrdup(m, v20)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v31
	v37 = F_psprintf(m, int32(_a_F_parse_extension_control_file_1), v14+int32(192))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v40 = v29
	goto L8
L18:
	;
	v40 = v37
	goto L8
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+180)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v40
	v52 = F_pg_snprintf(m, v42, int32(1024), int32(_a_F_parse_extension_control_file_2), v14+int32(176))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v69 = v42
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v57
	v65 = F_psprintf(m, int32(_a_F_parse_extension_control_file_3), v14+int32(224))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v69 = v65
	goto L4
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v193 = v72
	v196 = v69
	goto L2
L25:
	;
	if v131 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L26:
	;
	v79 = F_get_extension_control_directories(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v81 = int32(0)
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	if v79 == v81 {
		v131 = v81
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L44
	}
L29:
	;
	m.G0 = v84 + int32(32)
	goto L25
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v88 <= int32(0) {
		v131 = v81
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = v3
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v99<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v108 = F_pstrdup(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v131 = v81
	goto L29
L34:
	;
	F_canonicalize_path_enc(m, v108)
	mBase = m.M
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v111 != int32(47) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v108
	v117 = F_psprintf(m, int32(_a_F_parse_extension_control_file_1), v84)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v119 = F_pg_file_exists(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v119 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v131 = v117
	goto L29
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v108)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v117)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v126 = v99 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v126 < v127 {
		v99 = v126
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(_a_F_parse_extension_control_file_4)
	F_errmsg(m, int32(_a_F_parse_extension_control_file_5), v84+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(_a_F_parse_extension_control_file_7), int32(_a_F_parse_extension_control_file_8))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v167 = F_strlen(m, v131)
	mBase = m.M
	v174 = v167 + int32(1)
	goto L51
L49:
	;
	v188 = F_pnstrdup(m, v131, v186-v131)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L55
	}
L50:
	;
	goto L49
L51:
	;
	v176 = int32(0)
	if v174 == v176 {
		v186 = v176
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v186 = v181
	goto L50
L53:
	;
	v180 = v174 - int32(1)
	v181 = v131 + v180
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v182 != int32(47) {
		v174 = v180
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v188
	v193 = v188
	v196 = v131
	goto L2
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v205
	v209 = F_AllocateFile(m, v196, int32(_a_F_parse_extension_control_file_9))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L12
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v196)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L12
	} else {
		goto L264
	}
L58:
	;
	if v209 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if l1 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v240 = F_ParseConfigFp(m, v209, v196, int32(0), int32(21), v14+int32(236), v14+int32(232))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L70
	}
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_parse_extension_control_file[0]))
	if v214 == int32(44) {
		goto L57
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v196
	F_errmsg(m, int32(_a_F_parse_extension_control_file_10), v14+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(748), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v242 = F_FreeFile(m, v209)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	if v244 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v257 = v244
	goto L75
L73:
	;
	v831 = int32(0)
	goto L74
L74:
	;
	F_FreeConfigVariables(m, v831)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L12
	} else {
		goto L257
	}
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v267 = int32(_a_F_parse_extension_control_file_12)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[1])))
	if base.B2i32(v270 == int32(0))|base.B2i32(v270 != v273) != 0 {
		v291 = v270
		v292 = v273
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	v831 = v818
	goto L74
L77:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v257)+24))
	if v817 != 0 {
		v257 = v817
		goto L75
	} else {
		goto L256
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L12
	} else {
		goto L252
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L12
	} else {
		goto L248
	}
L80:
	;
	if v291-v292 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	v276 = v266
	v277 = v267
	goto L83
L83:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	if v281 == int32(0) {
		v291 = v281
		v292 = v280
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v291 = v281
	v292 = v280
	goto L81
L85:
	;
	v284 = int32(1)
	if v281 == v280 {
		v276 = v276 + v284
		v277 = v277 + v284
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	if l1 != 0 {
		goto L79
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v300 = int32(_a_F_parse_extension_control_file_13)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[2])))
	if base.B2i32(v303 == int32(0))|base.B2i32(v303 != v306) != 0 {
		v324 = v303
		v325 = v306
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v297 = F_pstrdup(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v297
	goto L77
L92:
	;
	if v324-v325 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	goto L92
L94:
	;
	v309 = v266
	v310 = v300
	goto L95
L95:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v314 == int32(0) {
		v324 = v314
		v325 = v313
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v324 = v314
	v325 = v313
	goto L93
L97:
	;
	v317 = int32(1)
	if v314 == v313 {
		v309 = v309 + v317
		v310 = v310 + v317
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if l1 != 0 {
		goto L78
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v333 = int32(_a_F_parse_extension_control_file_14)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[3])))
	if base.B2i32(v336 == int32(0))|base.B2i32(v336 != v339) != 0 {
		v357 = v336
		v358 = v339
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v330 = F_pstrdup(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v330
	goto L77
L104:
	;
	if v357-v358 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	v342 = v266
	v343 = v333
	goto L107
L107:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	if v347 == int32(0) {
		v357 = v347
		v358 = v346
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v357 = v347
	v358 = v346
	goto L105
L109:
	;
	v350 = int32(1)
	if v347 == v346 {
		v342 = v342 + v350
		v343 = v343 + v350
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v363 = F_pstrdup(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v366 = int32(_a_F_parse_extension_control_file_15)
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[4])))
	if base.B2i32(v369 == int32(0))|base.B2i32(v369 != v372) != 0 {
		v390 = v369
		v391 = v372
		goto L116
	} else {
		goto L117
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v363
	goto L77
L115:
	;
	if v390-v391 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L116:
	;
	goto L115
L117:
	;
	v375 = v266
	v376 = v366
	goto L118
L118:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+1)))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+1)))
	if v380 == int32(0) {
		v390 = v380
		v391 = v379
		goto L116
	} else {
		goto L120
	}
L119:
	;
	v390 = v380
	v391 = v379
	goto L116
L120:
	;
	v383 = int32(1)
	if v380 == v379 {
		v375 = v375 + v383
		v376 = v376 + v383
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v396 = F_pstrdup(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L12
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v399 = int32(_a_F_parse_extension_control_file_16)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[5])))
	if base.B2i32(v402 == int32(0))|base.B2i32(v402 != v405) != 0 {
		v423 = v402
		v424 = v405
		goto L127
	} else {
		goto L128
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v396
	goto L77
L126:
	;
	if v423-v424 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	goto L126
L128:
	;
	v408 = v266
	v409 = v399
	goto L129
L129:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+1)))
	if v413 == int32(0) {
		v423 = v413
		v424 = v412
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v423 = v413
	v424 = v412
	goto L127
L131:
	;
	v416 = int32(1)
	if v413 == v412 {
		v408 = v408 + v416
		v409 = v409 + v416
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v429 = F_pstrdup(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L12
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v432 = int32(_a_F_parse_extension_control_file_17)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[6])))
	if base.B2i32(v435 == int32(0))|base.B2i32(v435 != v438) != 0 {
		v456 = v435
		v457 = v438
		goto L138
	} else {
		goto L139
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v429
	goto L77
L137:
	;
	if v456-v457 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	v441 = v266
	v442 = v432
	goto L140
L140:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	if v446 == int32(0) {
		v456 = v446
		v457 = v445
		goto L138
	} else {
		goto L142
	}
L141:
	;
	v456 = v446
	v457 = v445
	goto L138
L142:
	;
	v449 = int32(1)
	if v446 == v445 {
		v441 = v441 + v449
		v442 = v442 + v449
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v462 = F_strlen(m, v461)
	mBase = m.M
	v463 = F_parse_bool_with_len(m, v461, v462, l0+int32(32))
	mBase = m.M
	goto L147
L145:
	;
	goto L146
L146:
	;
	v483 = int32(_a_F_parse_extension_control_file_18)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[7])))
	if base.B2i32(v486 == int32(0))|base.B2i32(v486 != v489) != 0 {
		v507 = v486
		v508 = v489
		goto L154
	} else {
		goto L155
	}
L147:
	;
	if v463 != 0 {
		goto L77
	} else {
		goto L148
	}
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v471
	F_errmsg(m, int32(_a_F_parse_extension_control_file_19), v14-int32(-64))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(803), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	if v507-v508 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	goto L153
L155:
	;
	v492 = v266
	v493 = v483
	goto L156
L156:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	if v497 == int32(0) {
		v507 = v497
		v508 = v496
		goto L154
	} else {
		goto L158
	}
L157:
	;
	v507 = v497
	v508 = v496
	goto L154
L158:
	;
	v500 = int32(1)
	if v497 == v496 {
		v492 = v492 + v500
		v493 = v493 + v500
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v513 = F_strlen(m, v512)
	mBase = m.M
	v514 = F_parse_bool_with_len(m, v512, v513, l0+int32(33))
	mBase = m.M
	goto L163
L161:
	;
	goto L162
L162:
	;
	v534 = int32(_a_F_parse_extension_control_file_20)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[8])))
	if base.B2i32(v537 == int32(0))|base.B2i32(v537 != v540) != 0 {
		v558 = v537
		v559 = v540
		goto L170
	} else {
		goto L171
	}
L163:
	;
	if v514 != 0 {
		goto L77
	} else {
		goto L164
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v522
	F_errmsg(m, int32(_a_F_parse_extension_control_file_19), v14+int32(80))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(811), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	if v558-v559 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L170:
	;
	goto L169
L171:
	;
	v543 = v266
	v544 = v534
	goto L172
L172:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+1)))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+1)))
	if v548 == int32(0) {
		v558 = v548
		v559 = v547
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v558 = v548
	v559 = v547
	goto L170
L174:
	;
	v551 = int32(1)
	if v548 == v547 {
		v543 = v543 + v551
		v544 = v544 + v551
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v564 = F_strlen(m, v563)
	mBase = m.M
	v565 = F_parse_bool_with_len(m, v563, v564, l0+int32(34))
	mBase = m.M
	goto L179
L177:
	;
	goto L178
L178:
	;
	v585 = int32(_a_F_parse_extension_control_file_21)
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[9])))
	if base.B2i32(v588 == int32(0))|base.B2i32(v588 != v591) != 0 {
		v609 = v588
		v610 = v591
		goto L186
	} else {
		goto L187
	}
L179:
	;
	if v565 != 0 {
		goto L77
	} else {
		goto L180
	}
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L12
	} else {
		goto L181
	}
L181:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L12
	} else {
		goto L182
	}
L182:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v573
	F_errmsg(m, int32(_a_F_parse_extension_control_file_19), v14+int32(96))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(819), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L12
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	if v609-v610 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L186:
	;
	goto L185
L187:
	;
	v594 = v266
	v595 = v585
	goto L188
L188:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+1)))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	if v599 == int32(0) {
		v609 = v599
		v610 = v598
		goto L186
	} else {
		goto L190
	}
L189:
	;
	v609 = v599
	v610 = v598
	goto L186
L190:
	;
	v602 = int32(1)
	if v599 == v598 {
		v594 = v594 + v602
		v595 = v595 + v602
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v615 = int32(-1)
	v618 = F_pg_char_to_encoding_private(m, v614)
	mBase = m.M
	if v618 == int32(7) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v650 = int32(_a_F_parse_extension_control_file_22)
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[10])))
	if base.B2i32(v653 == int32(0))|base.B2i32(v653 != v656) != 0 {
		v674 = v653
		v675 = v656
		goto L211
	} else {
		goto L212
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v627
	if int32(0) <= v627 {
		goto L77
	} else {
		goto L205
	}
L196:
	;
	v621 = v615
	goto L198
L197:
	;
	v621 = v618
	goto L198
L198:
	;
	if int32(34) < v618 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v624 = v615
	goto L201
L200:
	;
	v624 = v621
	goto L201
L201:
	;
	if v618 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v627 = v615
	goto L204
L203:
	;
	v627 = v624
	goto L204
L204:
	;
	goto L195
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L12
	} else {
		goto L206
	}
L206:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L12
	} else {
		goto L207
	}
L207:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v638
	F_errmsg(m, int32(_a_F_parse_extension_control_file_23), v14+int32(112))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L12
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(828), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	if v674-v675 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L211:
	;
	goto L210
L212:
	;
	v659 = v266
	v660 = v650
	goto L213
L213:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660)+1)))
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659)+1)))
	if v664 == int32(0) {
		v674 = v664
		v675 = v663
		goto L211
	} else {
		goto L215
	}
L214:
	;
	v674 = v664
	v675 = v663
	goto L211
L215:
	;
	v667 = int32(1)
	if v664 == v663 {
		v659 = v659 + v667
		v660 = v660 + v667
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v680 = F_pstrdup(m, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L12
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v704 = int32(_a_F_parse_extension_control_file_24)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v710 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_extension_control_file[11])))
	if base.B2i32(v707 == int32(0))|base.B2i32(v707 != v710) != 0 {
		v728 = v707
		v729 = v710
		goto L228
	} else {
		goto L229
	}
L220:
	;
	v683 = F_SplitIdentifierString(m, v680, int32(44), l0+int32(40))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L12
	} else {
		goto L221
	}
L221:
	;
	if v683 != 0 {
		goto L77
	} else {
		goto L222
	}
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L12
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L12
	} else {
		goto L224
	}
L224:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v692
	F_errmsg(m, int32(_a_F_parse_extension_control_file_25), v14+int32(128))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L12
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(842), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L12
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	if v728-v729 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L228:
	;
	goto L227
L229:
	;
	v713 = v266
	v714 = v704
	goto L230
L230:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+1)))
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)))
	if v718 == int32(0) {
		v728 = v718
		v729 = v717
		goto L228
	} else {
		goto L232
	}
L231:
	;
	v728 = v718
	v729 = v717
	goto L228
L232:
	;
	v721 = int32(1)
	if v718 == v717 {
		v713 = v713 + v721
		v714 = v714 + v721
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v734 = F_pstrdup(m, v733)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L12
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L12
	} else {
		goto L244
	}
L237:
	;
	v737 = F_SplitIdentifierString(m, v734, int32(44), l0+int32(44))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L12
	} else {
		goto L238
	}
L238:
	;
	if v737 != 0 {
		goto L77
	} else {
		goto L239
	}
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L12
	} else {
		goto L240
	}
L240:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L12
	} else {
		goto L241
	}
L241:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v746
	F_errmsg(m, int32(_a_F_parse_extension_control_file_25), v14+int32(144))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(857), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L12
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v765
	F_errmsg(m, int32(_a_F_parse_extension_control_file_26), v14+int32(160))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(864), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v785
	F_errmsg(m, int32(_a_F_parse_extension_control_file_27), v14+int32(32))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L12
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(771), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L12
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L12
	} else {
		goto L253
	}
L253:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v804
	F_errmsg(m, int32(_a_F_parse_extension_control_file_27), v14+int32(48))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L12
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(781), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	goto L76
L257:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v834 != int32(1) {
		goto L57
	} else {
		goto L258
	}
L258:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v837 == int32(0) {
		goto L57
	} else {
		goto L259
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L12
	} else {
		goto L260
	}
L260:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L12
	} else {
		goto L261
	}
L261:
	;
	F_errmsg(m, int32(_a_F_parse_extension_control_file_28), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(872), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L12
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	m.G0 = v14 + int32(240)
	return
L265:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L12
	} else {
		goto L266
	}
L266:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v890
	F_errmsg(m, int32(_a_F_parse_extension_control_file_29), v14)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L12
	} else {
		goto L267
	}
L267:
	;
	F_errhint(m, int32(_a_F_parse_extension_control_file_30), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L12
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_parse_extension_control_file_6), int32(725), int32(_a_F_parse_extension_control_file_11))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L12
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
