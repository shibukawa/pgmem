package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreFindKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v166
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v159
	v166 = v154
	goto L1
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = v12
	goto L5
L4:
	;
	v14 = int32(0)
	goto L5
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = v15 & int32(268435455)
	if v14 < v17 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = l0 + int32(8)
	v29 = v14
	v30 = v17
	goto L9
L7:
	;
	v145 = v14
	goto L8
L8:
	;
	v151 = int32(-1)
	if l1 == int32(0) {
		v166 = v151
		goto L1
	} else {
		goto L49
	}
L9:
	;
	v38 = int32(base.Ui32(v30-v29)>>(uint(int32(1))%32)) + v29
	v41 = v20 + v38<<(uint(int32(3))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v44 = v42 & int32(1073741823)
	if int32(0) <= v42 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v145 = v137
	goto L8
L11:
	;
	v136 = base.B2i32(v131 < int32(0))
	if v131 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L12:
	;
	v64 = v63 + (v20 + v17<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(l3) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if base.Ui32(l3) < base.Ui32(v56) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(4))))
	v51 = v49 & int32(1073741823)
	v52 = v44 - v51
	if v52 != l3 {
		v56 = v52
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l3 == v44 {
		v63 = int32(0)
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v63 = v51
	goto L12
L18:
	;
	v56 = v44
	goto L13
L19:
	;
	v61 = int32(1)
	goto L21
L20:
	;
	v61 = int32(-1)
	goto L21
L21:
	;
	v131 = v61
	goto L11
L22:
	;
	if v126 != 0 {
		v131 = v126
		goto L11
	} else {
		goto L40
	}
L23:
	;
	v126 = int32(0)
	goto L22
L24:
	;
	v100 = v95
	v101 = v96
	v102 = v97
	goto L34
L25:
	;
	if (v64|l2)&int32(3) != 0 {
		v95 = v64
		v96 = l2
		v97 = l3
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v88 = v64
	v89 = l2
	v90 = l3
	goto L27
L27:
	;
	if v90 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v72 = v64
	v73 = l2
	v74 = l3
	goto L29
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v77 != v78 {
		v95 = v72
		v96 = v73
		v97 = v74
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v88 = v83
	v89 = v81
	v90 = v85
	goto L27
L31:
	;
	v80 = int32(4)
	v81 = v73 + v80
	v83 = v72 + v80
	v85 = v74 - v80
	if base.Ui32(int32(3)) < base.Ui32(v85) {
		v72 = v83
		v73 = v81
		v74 = v85
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v95 = v88
	v96 = v89
	v97 = v90
	goto L24
L34:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v105 == v106 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v126 = v105 - v106
	goto L22
L36:
	;
	v108 = int32(1)
	v113 = v102 - v108
	if v113 != 0 {
		v100 = v100 + v108
		v101 = v101 + v108
		v102 = v113
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L23
L40:
	;
	if l1 == int32(0) {
		v166 = v38
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v154 = v38
	v159 = v38 + int32(1)
	goto L2
L42:
	;
	v137 = v38 + int32(1)
	goto L44
L43:
	;
	v137 = v29
	goto L44
L44:
	;
	if v131 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v138 = v30
	goto L47
L46:
	;
	v138 = v38
	goto L47
L47:
	;
	if v137 < v138 {
		v29 = v137
		v30 = v138
		goto L9
	} else {
		goto L48
	}
L48:
	;
	goto L10
L49:
	;
	v154 = v151
	v159 = v145
	goto L2
}
func F_hstore_defined(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v53 = v51 & int32(268435455)
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v29 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v40 = int32(1)
	if v23 != 0 {
		v50 = int32(base.Ui32(v21)>>(uint(v40)%32)) - v40
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v32 = int32(16)
	goto L10
L9:
	;
	v32 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(4)
	goto L13
L12:
	;
	v39 = v32
	goto L13
L13:
	;
	v50 = v39
	goto L4
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v23 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	return int64(0)
L18:
	;
	v56 = int32(1)
	goto L20
L19:
	;
	v56 = int32(4)
	goto L20
L20:
	;
	v57 = v19 + v56
	v59 = v14 + int32(8)
	v64 = int32(0)
	v68 = v53
	goto L21
L21:
	;
	v79 = int32(base.Ui32(v68-v64)>>(uint(int32(1))%32)) + v64
	v81 = v79 << (uint(int32(3)) % 32)
	v82 = v59 + v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v85 = v83 & int32(1073741823)
	if int32(0) <= v83 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L17
L23:
	;
	v181 = base.B2i32(v176 < int32(0))
	if v176 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L24:
	;
	v105 = v104 + (v59 + v53<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if base.Ui32(v50) < base.Ui32(v97) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82-int32(4))))
	v92 = v90 & int32(1073741823)
	v93 = v85 - v92
	if v93 != v50 {
		v97 = v93
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v85 == v50 {
		v104 = int32(0)
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v104 = v92
	goto L24
L30:
	;
	v97 = v85
	goto L25
L31:
	;
	v102 = int32(1)
	goto L33
L32:
	;
	v102 = int32(-1)
	goto L33
L33:
	;
	v176 = v102
	goto L23
L34:
	;
	if v167 != 0 {
		v176 = v167
		goto L23
	} else {
		goto L52
	}
L35:
	;
	v167 = int32(0)
	goto L34
L36:
	;
	v141 = v136
	v142 = v137
	v143 = v138
	goto L46
L37:
	;
	if (v105|v57)&int32(3) != 0 {
		v136 = v105
		v137 = v57
		v138 = v50
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v129 = v105
	v130 = v57
	v131 = v50
	goto L39
L39:
	;
	if v131 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v113 = v105
	v114 = v57
	v115 = v50
	goto L41
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v118 != v119 {
		v136 = v113
		v137 = v114
		v138 = v115
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v129 = v124
	v130 = v122
	v131 = v126
	goto L39
L43:
	;
	v121 = int32(4)
	v122 = v114 + v121
	v124 = v113 + v121
	v126 = v115 - v121
	if base.Ui32(int32(3)) < base.Ui32(v126) {
		v113 = v124
		v114 = v122
		v115 = v126
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v136 = v129
	v137 = v130
	v138 = v131
	goto L36
L46:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 == v147 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v167 = v146 - v147
	goto L34
L48:
	;
	v149 = int32(1)
	v154 = v143 - v149
	if v154 != 0 {
		v141 = v141 + v149
		v142 = v142 + v149
		v143 = v154
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v81)+15)))
	return base.I64_extend_i32_u(base.B2i32(v169&int32(64) == int32(0)))
L53:
	;
	v182 = v79 + int32(1)
	goto L55
L54:
	;
	v182 = v64
	goto L55
L55:
	;
	if v176 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v183 = v68
	goto L58
L57:
	;
	v183 = v79
	goto L58
L58:
	;
	if v182 < v183 {
		v64 = v182
		v68 = v183
		goto L21
	} else {
		goto L59
	}
L59:
	;
	goto L22
}
func F_hstore_delete_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = F_hstoreUpgrade(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v31 = F_palloc(m, int32(base.Ui32(v28)>>(uint(int32(2))%32)))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = F_hstoreArrayToPairs(m, v35, v21+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v43 = v33 & int32(268435455)
	v45 = v43 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v41 & int32(-4)
	v51 = v31 + int32(8)
	v53 = v43 << (uint(int32(3)) % 32)
	v54 = v51 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	m.G0 = v21 + int32(16)
	return base.I64_extend_i32_u(v31)
L7:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v315 == v324&int32(268435455) {
		goto L73
	} else {
		goto L74
	}
L8:
	;
	v74 = v24 + int32(8)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v80 = v74 + v75<<(uint(int32(3))%32)&int32(2147483640)
	v85 = int32(0)
	v86 = v54
	v91 = v2
	v92 = v51
	v93 = v2
	goto L19
L9:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v59 = int32(base.Ui32(v57) >> (uint(int32(2)) % 32))
	if v59 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v309 = int32(0)
	v315 = v2
	goto L7
L13:
	;
	base.MemoryCopy(m, v31, v24, v59)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(4))))
	v65 = v63
	goto L18
L17:
	;
	v65 = int32(0)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = (v65+v53)<<(uint(int32(2))%32) + int32(32)
	goto L6
L19:
	;
	v101 = v85 << (uint(int32(3)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v102 <= v93 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v298 = v290 - v54
	if v293 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L21:
	;
	if v289 < v43 {
		v85 = v289
		v86 = v290
		v91 = v293
		v92 = v294
		v93 = v295
		goto L19
	} else {
		goto L69
	}
L22:
	;
	v216 = v101 + v74
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v217 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	v104 = v101 + v74
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = v105 & int32(1073741823)
	if int32(0) <= v105 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v289 = v85
	v290 = v86
	v293 = v91
	v294 = v92
	v295 = v93 + int32(1)
	goto L21
L25:
	;
	v136 = v134 + v80
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if base.Ui32(int32(4)) <= base.Ui32(v133) {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	if base.Ui32(v128) <= base.Ui32(v130) {
		goto L22
	} else {
		goto L32
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104-int32(4))))
	v114 = v112 & int32(1073741823)
	v115 = v107 - v114
	v118 = v39 + v93*int32(20)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	if v115 != v119 {
		v128 = v115
		v130 = v119
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v124 = v39 + v93*int32(20)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v107 == v125 {
		v132 = v124
		v133 = v107
		v134 = int32(0)
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v132 = v118
	v133 = v115
	v134 = v114
	goto L25
L31:
	;
	v128 = v107
	v130 = v125
	goto L26
L32:
	;
	goto L24
L33:
	;
	if int32(0) < v199 {
		goto L24
	} else {
		goto L51
	}
L34:
	;
	v199 = int32(0)
	goto L33
L35:
	;
	v173 = v168
	v174 = v169
	v175 = v170
	goto L45
L36:
	;
	if (v136|v137)&int32(3) != 0 {
		v168 = v136
		v169 = v137
		v170 = v133
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v161 = v136
	v162 = v137
	v163 = v133
	goto L38
L38:
	;
	if v163 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v145 = v136
	v146 = v137
	v147 = v133
	goto L40
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v150 != v151 {
		v168 = v145
		v169 = v146
		v170 = v147
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v161 = v156
	v162 = v154
	v163 = v158
	goto L38
L42:
	;
	v153 = int32(4)
	v154 = v146 + v153
	v156 = v145 + v153
	v158 = v147 - v153
	if base.Ui32(int32(3)) < base.Ui32(v158) {
		v145 = v156
		v146 = v154
		v147 = v158
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v168 = v161
	v169 = v162
	v170 = v163
	goto L35
L45:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v178 == v179 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v199 = v178 - v179
	goto L33
L47:
	;
	v181 = int32(1)
	v186 = v175 - v181
	if v186 != 0 {
		v173 = v173 + v181
		v174 = v174 + v181
		v175 = v186
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L34
L51:
	;
	if v199 != 0 {
		goto L22
	} else {
		goto L52
	}
L52:
	;
	v202 = int32(1)
	v289 = v85 + v202
	v290 = v86
	v293 = v91
	v294 = v92
	v295 = v93 + v202
	goto L21
L53:
	;
	v236 = v216 + int32(4)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v240 = int32(0)
	if v240 <= v237 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v221 = v217 & int32(1073741823)
	v231 = v221
	v232 = v221
	v234 = v80
	goto L53
L55:
	;
	goto L56
L56:
	;
	v222 = int32(1073741823)
	v223 = v217 & v222
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(4))))
	v228 = v226 & v222
	v231 = v223 - v228
	v232 = v223
	v234 = v228 + v80
	goto L53
L57:
	;
	v243 = v232
	goto L59
L58:
	;
	v243 = v240
	goto L59
L59:
	;
	v245 = v231 + (v237&int32(1073741823) - v243)
	if v245 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v86, v234, v245)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v249 = v247 & int32(1073741823)
	if int32(0) <= v247 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(4))))
	v258 = v249 - v254&int32(1073741823)
	goto L65
L64:
	;
	v258 = v249
	goto L65
L65:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v262 = int32(0)
	if v262 <= v259 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v265 = v249
	goto L68
L67:
	;
	v265 = v262
	goto L68
L68:
	;
	v266 = v259&int32(1073741823) - v265
	v268 = v266 + (v258 + v86)
	v269 = v268 - v54
	v271 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = (v269 - v266) & v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v274&int32(1073741824) | v269&v271
	v281 = int32(1)
	v289 = v85 + v281
	v290 = v268
	v293 = v91 + v281
	v294 = v92 + int32(8)
	v295 = v93
	goto L21
L69:
	;
	goto L20
L70:
	;
	v309 = v298
	v315 = int32(0)
	goto L7
L71:
	;
	goto L72
L72:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v302 | int32(-2147483648)
	v309 = v298
	v315 = v293
	goto L7
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = (v315<<(uint(int32(3))%32)+v309)<<(uint(int32(2))%32) + int32(32)
	goto L6
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v315 | int32(-2147483648)
	if v309 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	base.MemoryCopy(m, v51+v315<<(uint(int32(3))%32)&int32(2147483640), v54, v309)
	goto L73
}
func F_hstore_delete_hstore(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v509 int32
	_ = v509
	v2 = int32(0)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_hstoreUpgrade(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v30 = F_hstoreUpgrade(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v35 = F_palloc(m, int32(base.Ui32(v32)>>(uint(int32(2))%32)))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v39 & int32(-4)
	v43 = int32(268435455)
	v44 = v38 & v43
	v46 = v44 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v46
	v49 = v35 + int32(8)
	v51 = v44 << (uint(int32(3)) % 32)
	v52 = v49 + v51
	v54 = v37 & v43
	if v54 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v495 == v509&int32(268435455) {
		goto L121
	} else {
		goto L122
	}
L6:
	;
	v74 = int32(8)
	v75 = v30 + v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v77 = int32(3)
	v79 = int32(2147483640)
	v81 = v75 + v76<<(uint(v77)%32)&v79
	v83 = v25 + v74
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v89 = v83 + v84<<(uint(v77)%32)&v79
	v91 = v49
	v96 = v52
	v97 = int32(0)
	v100 = v2
	v103 = v2
	goto L17
L7:
	;
	if v44 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v58 = int32(base.Ui32(v56) >> (uint(int32(2)) % 32))
	if v58 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v492 = int32(0)
	v495 = v2
	goto L5
L11:
	;
	base.MemoryCopy(m, v35, v25, v58)
	goto L13
L12:
	;
	goto L13
L13:
	;
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52-int32(4))))
	v64 = v62
	goto L16
L15:
	;
	v64 = int32(0)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = (v64+v51)<<(uint(int32(2))%32) + int32(32)
	return base.I64_extend_i32_u(v35)
L17:
	;
	v115 = v97 << (uint(int32(3)) % 32)
	if v54 <= v103 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v478 = v468 - v52
	if v471 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L19:
	;
	if v469 < v44 {
		v91 = v463
		v96 = v468
		v97 = v469
		v100 = v471
		v103 = v472
		goto L17
	} else {
		goto L117
	}
L20:
	;
	v463 = v91
	v468 = v96
	v469 = v97
	v471 = v100
	v472 = v103 + int32(1)
	goto L19
L21:
	;
	v389 = v115 + v83
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	if v390 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L22:
	;
	v117 = v115 + v83
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v119 = int32(1073741823)
	v120 = v118 & v119
	v123 = v75 + v103<<(uint(int32(3))%32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v126 = v124 & v119
	if int32(0) <= v118 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
	v135 = v120 - v131&int32(1073741823)
	goto L25
L24:
	;
	v135 = v120
	goto L25
L25:
	;
	if int32(0) <= v124 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v123-int32(4))))
	v144 = v126 - v140&int32(1073741823)
	goto L28
L27:
	;
	v144 = v126
	goto L28
L28:
	;
	if v144 == v135 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if int32(0) <= v118 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v144 < v135 {
		goto L20
	} else {
		goto L100
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
	v154 = v150 & int32(1073741823)
	goto L34
L33:
	;
	v154 = int32(0)
	goto L34
L34:
	;
	v155 = v154 + v89
	if int32(0) <= v124 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v123-int32(4))))
	v164 = v160 & int32(1073741823)
	goto L37
L36:
	;
	v164 = int32(0)
	goto L37
L37:
	;
	v165 = v164 + v81
	if base.Ui32(int32(4)) <= base.Ui32(v135) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if int32(0) < v227 {
		goto L20
	} else {
		goto L56
	}
L39:
	;
	v227 = int32(0)
	goto L38
L40:
	;
	v201 = v196
	v202 = v197
	v203 = v198
	goto L50
L41:
	;
	if (v155|v165)&int32(3) != 0 {
		v196 = v155
		v197 = v165
		v198 = v135
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v189 = v155
	v190 = v165
	v191 = v135
	goto L43
L43:
	;
	if v191 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v173 = v155
	v174 = v165
	v175 = v135
	goto L45
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v178 != v179 {
		v196 = v173
		v197 = v174
		v198 = v175
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v189 = v184
	v190 = v182
	v191 = v186
	goto L43
L47:
	;
	v181 = int32(4)
	v182 = v174 + v181
	v184 = v173 + v181
	v186 = v175 - v181
	if base.Ui32(int32(3)) < base.Ui32(v186) {
		v173 = v184
		v174 = v182
		v175 = v186
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v196 = v189
	v197 = v190
	v198 = v191
	goto L40
L50:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v206 == v207 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v227 = v206 - v207
	goto L38
L52:
	;
	v209 = int32(1)
	v214 = v203 - v209
	if v214 != 0 {
		v201 = v201 + v209
		v202 = v202 + v209
		v203 = v214
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L39
L56:
	;
	if v227 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v231 = int32(1073741823)
	v234 = v118 & v231
	v235 = int32(0)
	v237 = base.B2i32(v235 <= v230)
	if v235 <= v230 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v238 = v234
	goto L60
L59:
	;
	v238 = v235
	goto L60
L60:
	;
	v239 = v230&v231 - v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v241 = int32(30)
	v246 = v230 & int32(1073741824)
	if int32(base.Ui32(v240)>>(uint(v241)%32))&int32(1) != int32(base.Ui32(v246)>>(uint(v241)%32)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v378 = int32(1)
	v463 = v372
	v468 = v375
	v469 = v97 + v378
	v471 = v376
	v472 = v103 + v378
	goto L19
L62:
	;
	if v118 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L63:
	;
	if v246 != 0 {
		v372 = v91
		v375 = v96
		v376 = v100
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v250 = int32(1073741823)
	v254 = int32(0)
	if v254 <= v240 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v257 = v124 & v250
	goto L67
L66:
	;
	v257 = v254
	goto L67
L67:
	;
	if v239 != v240&v250-v257 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v235 <= v230 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v261 = v234
	goto L71
L70:
	;
	v261 = int32(0)
	goto L71
L71:
	;
	v262 = v89 + v261
	v263 = v257 + v81
	if base.Ui32(int32(4)) <= base.Ui32(v239) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	if v325 == int32(0) {
		v372 = v91
		v375 = v96
		v376 = v100
		goto L61
	} else {
		goto L90
	}
L73:
	;
	v325 = int32(0)
	goto L72
L74:
	;
	v299 = v294
	v300 = v295
	v301 = v296
	goto L84
L75:
	;
	if (v262|v263)&int32(3) != 0 {
		v294 = v262
		v295 = v263
		v296 = v239
		goto L74
	} else {
		goto L78
	}
L76:
	;
	v287 = v262
	v288 = v263
	v289 = v239
	goto L77
L77:
	;
	if v289 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L78:
	;
	v271 = v262
	v272 = v263
	v273 = v239
	goto L79
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v276 != v277 {
		v294 = v271
		v295 = v272
		v296 = v273
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v287 = v282
	v288 = v280
	v289 = v284
	goto L77
L81:
	;
	v279 = int32(4)
	v280 = v272 + v279
	v282 = v271 + v279
	v284 = v273 - v279
	if base.Ui32(int32(3)) < base.Ui32(v284) {
		v271 = v282
		v272 = v280
		v273 = v284
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v294 = v287
	v295 = v288
	v296 = v289
	goto L74
L84:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v304 == v305 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v325 = v304 - v305
	goto L72
L86:
	;
	v307 = int32(1)
	v312 = v301 - v307
	if v312 != 0 {
		v299 = v299 + v307
		v300 = v300 + v307
		v301 = v312
		goto L84
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	goto L73
L90:
	;
	goto L62
L91:
	;
	v341 = v234
	v342 = v89
	goto L93
L92:
	;
	v331 = int32(1073741823)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
	v337 = v335 & v331
	v341 = v118&v331 - v337
	v342 = v337 + v89
	goto L93
L93:
	;
	v343 = v239 + v341
	if v343 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v96, v342, v343)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v347 = v345 & int32(1073741823)
	if int32(0) <= v345 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
	v356 = v347 - v352&int32(1073741823)
	goto L99
L98:
	;
	v356 = v347
	goto L99
L99:
	;
	v358 = v356 + v96 + v239
	v359 = v358 - v52
	v360 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v359&v360 | v246
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = (v359 - v239) & v360
	v372 = v91 + int32(8)
	v375 = v358
	v376 = v100 + int32(1)
	goto L61
L100:
	;
	goto L21
L101:
	;
	v409 = v389 + int32(4)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v413 = int32(0)
	if v413 <= v410 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v394 = v390 & int32(1073741823)
	v404 = v394
	v405 = v394
	v407 = v89
	goto L101
L103:
	;
	goto L104
L104:
	;
	v395 = int32(1073741823)
	v396 = v390 & v395
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v389-int32(4))))
	v401 = v399 & v395
	v404 = v396
	v405 = v396 - v401
	v407 = v401 + v89
	goto L101
L105:
	;
	v416 = v404
	goto L107
L106:
	;
	v416 = v413
	goto L107
L107:
	;
	v418 = v405 + (v410&int32(1073741823) - v416)
	if v418 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	base.MemoryCopy(m, v96, v407, v418)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v422 = v420 & int32(1073741823)
	if int32(0) <= v420 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v389-int32(4))))
	v431 = v422 - v427&int32(1073741823)
	goto L113
L112:
	;
	v431 = v422
	goto L113
L113:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v435 = int32(0)
	if v435 <= v432 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v438 = v422
	goto L116
L115:
	;
	v438 = v435
	goto L116
L116:
	;
	v439 = v432&int32(1073741823) - v438
	v441 = v439 + (v431 + v96)
	v442 = v441 - v52
	v444 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = (v442 - v439) & v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v447&int32(1073741824) | v442&v444
	v454 = int32(1)
	v463 = v91 + int32(8)
	v468 = v441
	v469 = v97 + v454
	v471 = v100 + v454
	v472 = v103
	goto L19
L117:
	;
	goto L18
L118:
	;
	v492 = v478
	v495 = int32(0)
	goto L5
L119:
	;
	goto L120
L120:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v482 | int32(-2147483648)
	v492 = v478
	v495 = v471
	goto L5
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = (v495<<(uint(int32(3))%32)+v492)<<(uint(int32(2))%32) + int32(32)
	return base.I64_extend_i32_u(v35)
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v495 | int32(-2147483648)
	if v492 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	base.MemoryCopy(m, v49+v495<<(uint(int32(3))%32)&int32(2147483640), v52, v492)
	goto L121
}
func F_hstore_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v404 int64
	_ = v404
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v12
	v16 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v16
	v20 = F_palloc_mul(m, int32(20), v16)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v12
	v33 = v12
	v34 = v24
	goto L4
L3:
	;
	m.G0 = v10 + int32(96)
	return v404
L4:
	;
	switch v34 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L11
	default:
		goto L15
	}
L5:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v393 = F_hstoreUniquePairs(m, v389, v390, v10+int32(56))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L113
	}
L6:
	;
	goto L5
L7:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v385 = v383 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v385
	v33 = v385
	v34 = v382
	goto L4
L8:
	;
	v382 = int32(2)
	goto L7
L9:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	v366 = v363 + v56*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = v69
	v369 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v369
	goto L8
L10:
	;
	v360 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v360)
	v404 = int64(0)
	goto L3
L11:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v315 == int32(44) {
		v382 = int32(0)
		goto L7
	} else {
		goto L103
	}
L12:
	;
	v199 = F_get_val(m, v10+int32(60), int32(1), v10+int32(56))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v147 == int32(62) {
		v382 = int32(1)
		goto L7
	} else {
		goto L50
	}
L14:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v91 == int32(61) {
		v382 = int32(3)
		goto L7
	} else {
		goto L33
	}
L15:
	;
	v45 = F_get_val(m, v10+int32(60), int32(0), v10+int32(56))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v45 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v49 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	if v57 <= v56 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v52 != int32(453) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if v55 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v57 << (uint(int32(1)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	v65 = F_repalloc(m, v62, v57*int32(40))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	v70 = v68 - v69
	if base.Ui32(v70) < base.Ui32(int32(1073741824)) {
		goto L9
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v65
	goto L25
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v74 = F_errsave_start(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v74 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_hstore_in_0), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, v73, int32(_a_F_hstore_in_1), int32(423), int32(_a_F_hstore_in_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	if v91 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v97 = F_errsave_start(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v113 = base.I32_extend8_s(v91)
	goto L42
L37:
	;
	if v97 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errsave_finish(m, v96, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L10
L42:
	;
	if base.B2i32(v113 == int32(32))|base.B2i32(base.Ui32((v113-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v124 = F_errsave_start(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v124 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v131 = F_pg_mblen_cstr(m, v33)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v33
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v33 - v135
	F_errmsg(m, int32(_a_F_hstore_in_5), v10)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errsave_finish(m, v123, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L10
L50:
	;
	if v147 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v153 = F_errsave_start(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v170 = F_errsave_start(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	if v153 == int32(0) {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, v152, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	if v170 == int32(0) {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v177 = F_pg_mblen_cstr(m, v33)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v33
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33 - v181
	F_errmsg(m, int32(_a_F_hstore_in_5), v10+int32(16))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errsave_finish(m, v169, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	if v199 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v203 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	v228 = v226 - v227
	if base.Ui32(int32(1073741824)) <= base.Ui32(v228) {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	v210 = F_errsave_start(m, v203)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v206 != int32(453) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)))
	if v209 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v210 == int32(0) {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errsave_finish(m, v203, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L10
L78:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v232 = F_errsave_start(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v252 = v248 + v249*int32(20)
	v253 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v252)+16)) = uint16(v253)
	*(*int32)(unsafe.Add(mBase, uint32(v252)+12)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v252)+4)) = v227
	if v228 != int32(4) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	if v232 == int32(0) {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_hstore_in_7), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errsave_finish(m, v231, int32(_a_F_hstore_in_1), int32(443), int32(_a_F_hstore_in_8))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L10
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v249 + int32(1)
	v382 = int32(4)
	goto L7
L87:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)))
	if v259&int32(1) != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+4)) = uint8(v262)
	v267 = v227
	v268 = int32(_a_F_hstore_in_9)
	goto L90
L89:
	;
	if v305 != 0 {
		goto L86
	} else {
		goto L102
	}
L90:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v271 == v272 {
		v294 = v271
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v305 = int32(0)
	goto L89
L92:
	;
	v296 = int32(1)
	if v294 != 0 {
		v267 = v267 + v296
		v268 = v268 + v296
		goto L90
	} else {
		goto L101
	}
L93:
	;
	if base.Ui32((v271-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v282 = v271 | int32(32)
	goto L96
L95:
	;
	v282 = v271
	goto L96
L96:
	;
	if base.Ui32((v272-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v291 = v272 | int32(32)
	goto L99
L98:
	;
	v291 = v272
	goto L99
L99:
	;
	if v282 == v291 {
		v294 = v282
		goto L92
	} else {
		goto L100
	}
L100:
	;
	v305 = v282 - v291
	goto L89
L101:
	;
	goto L91
L102:
	;
	v306 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v252)+16)) = uint8(v306)
	goto L86
L103:
	;
	if v315 == int32(0) {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v321 = base.I32_extend8_s(v315)
	goto L105
L105:
	;
	if base.B2i32(v321 == int32(32))|base.B2i32(base.Ui32((v321-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v382 = int32(4)
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v332 = F_errsave_start(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v332 == int32(0) {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v339 = F_pg_mblen_cstr(m, v33)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v33
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v33 - v343
	F_errmsg(m, int32(_a_F_hstore_in_5), v10+int32(32))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errsave_finish(m, v331, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L10
L113:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v396 = F_hstorePairs(m, v389, v393, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v404 = base.I64_extend_i32_u(v396)
	goto L3
}
func F_hstore_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_le_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.I32_wrap_i64(v6) <= int32(0)))
	}
}
func F_hstore_skeys(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = F_hstoreUpgrade(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = int32(_a_F_hstore_skeys_0)
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_hstore_skeys[0]))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_hstore_skeys[0])) = v21
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v26 = F_palloc(m, int32(base.Ui32(v23)>>(uint(int32(2))%32)))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v30 = int32(base.Ui32(v28) >> (uint(int32(2)) % 32))
					if v30 != 0 {
						base.MemoryCopy(m, v26, v12, v30)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v26
					*(*int32)(unsafe.Add(mBase, _c_F_hstore_skeys[0])) = v19
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v45 = v43 & int32(268435455)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					if base.Ui32(v46) < base.Ui32(v45) {
						v49 = v42 + int32(8)
						v50 = int32(3)
						v52 = v49 + v45<<(uint(v50)%32)
						v55 = v49 + v46<<(uint(v50)%32)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						if v56 < int32(0) {
							v70 = v52
							v71 = v56 & int32(1073741823)
						} else {
							v61 = int32(1073741823)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
							v67 = v65 & v61
							v70 = v52 + v67
							v71 = v56&v61 - v67
						}
						v73 = F_cstring_to_text_with_len(m, v70, v71)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
						} else {
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
							*(*int64)(unsafe.Add(mBase, uint32(v41))) = v75 + int64(1)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = int32(1)
							return base.I64_extend_i32_u(v73)
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(2)
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
		v45 = v43 & int32(268435455)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
		if base.Ui32(v46) < base.Ui32(v45) {
			v49 = v42 + int32(8)
			v50 = int32(3)
			v52 = v49 + v45<<(uint(v50)%32)
			v55 = v49 + v46<<(uint(v50)%32)
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
			if v56 < int32(0) {
				v70 = v52
				v71 = v56 & int32(1073741823)
			} else {
				v61 = int32(1073741823)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
				v67 = v65 & v61
				v70 = v52 + v67
				v71 = v56&v61 - v67
			}
			v73 = F_cstring_to_text_with_len(m, v70, v71)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int64(0)
			} else {
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v75 + int64(1)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = int32(1)
				return base.I64_extend_i32_u(v73)
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int64(0)
			} else {
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(2)
				v89 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
				return int64(0)
			}
		}
	}
}
func F_hstore_to_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_hstoreUpgrade(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_hstore_to_array_internal(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_hstore_to_json_loose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_hstoreUpgrade(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_u(v169)
L2:
	;
	return int64(0)
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v21 = v19 & int32(268435455)
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = F_cstring_to_text_with_len(m, int32(_a_F_hstore_to_json_loose_0), int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = v15 + int32(8)
	v32 = v29 + v21<<(uint(int32(3))%32)
	F_initStringInfo(m, v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v169 = v26
	goto L1
L8:
	;
	F_appendStringInfoChar(m, v12, int32(123))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v39 = int32(0)
	goto L10
L10:
	;
	v50 = v29 + v39<<(uint(int32(3))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_appendStringInfoChar(m, v12, int32(125))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L56
	}
L12:
	;
	F_escape_json_with_len(m, v12, v67, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v65 = v51 & int32(1073741823)
	v67 = v32
	goto L12
L14:
	;
	goto L15
L15:
	;
	v56 = int32(1073741823)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50-int32(4))))
	v62 = v60 & v56
	v65 = v51&v56 - v62
	v67 = v62 + v32
	goto L12
L16:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_loose_1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v73&int32(1073741824) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v147 = v39 + int32(1)
	if v21 != v147 {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_loose_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if int32(0) <= v73 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L18
L23:
	;
	v137 = F_IsValidJsonNumber(m, v136, v134)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L45
	}
L24:
	;
	if v73 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v114))))
	if v116 != int32(102) {
		goto L24
	} else {
		goto L40
	}
L26:
	;
	if v109 != int32(1) {
		v134 = v109
		v136 = v32
		goto L23
	} else {
		goto L39
	}
L27:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_loose_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L38
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v83 = v81 & int32(1073741823)
	if v73-v83 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v98 = v73 & int32(1073741823)
	if v98 != int32(1) {
		v109 = v98
		goto L26
	} else {
		goto L36
	}
L31:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v32))))
	if v88 == int32(116) {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v93 = v91 & int32(1073741823)
	if v73-v93 == int32(1) {
		v114 = v93
		goto L25
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L24
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v102 != int32(116) {
		v109 = int32(1)
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	goto L18
L39:
	;
	v114 = int32(0)
	goto L25
L40:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_loose_4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	goto L18
L42:
	;
	v134 = v73 & int32(1073741823)
	v136 = v32
	goto L23
L43:
	;
	goto L44
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v130 = v128 & int32(1073741823)
	v134 = v73 - v130
	v136 = v130 + v32
	goto L23
L45:
	;
	if v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_appendBinaryStringInfo(m, v12, v136, v134)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_escape_json_with_len(m, v12, v136, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L50
	}
L49:
	;
	goto L18
L50:
	;
	goto L18
L51:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_loose_5))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v147 != v21 {
		v39 = v147
		goto L10
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	goto L11
L56:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v158 = F_cstring_to_text_with_len(m, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v169 = v158
	goto L1
}
func F_hstore_to_jsonb_loose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = int32(0)
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v21
	F_initStringInfo(m, v11+int32(72))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_pushJsonbValue(m, v11+int32(88), int32(6), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = v18 & int32(268435455)
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v38 = v14 + int32(8)
	v41 = v38 + v36<<(uint(int32(3))%32)
	v47 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_pushJsonbValue(m, v11+int32(88), int32(7), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(1)
	v54 = v38 + v47<<(uint(int32(3))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v57 = v55 & int32(1073741823)
	if v55 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v73 + v41
	F_pushJsonbValue(m, v11+int32(88), int32(1), v11+int32(40))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v57
	v73 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v63 = v54 - int32(4)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v57 - v64&v65
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v73 = v69 & v65
	goto L10
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v83&int32(1073741824) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	F_pushJsonbValue(m, v11+int32(88), int32(2), v11+int32(8))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L49
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if int32(0) <= v83 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v133 = v11 + int32(72)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v135)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v135
	goto L33
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v41))))
	if v123 != int32(102) {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(3)
	goto L15
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v92 = v90 & int32(1073741823)
	if v83-v92 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v83&int32(1073741823) != int32(1) {
		goto L19
	} else {
		goto L30
	}
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v92))))
	if v97 == int32(116) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v102 = v100 & int32(1073741823)
	if v83-v102 != int32(1) {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v120 = v102
	goto L20
L30:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v111 != int32(116) {
		v120 = int32(0)
		goto L20
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v126)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(3)
	goto L15
L33:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v141 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_appendBinaryStringInfo(m, v133, v155, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v153 = v141 & int32(1073741823)
	v155 = v41
	goto L34
L36:
	;
	goto L37
L37:
	;
	v146 = int32(1073741823)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v150 = v148 & v146
	v153 = v141&v146 - v150
	v155 = v41 + v150
	goto L34
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v160 = F_IsValidJsonNumber(m, v158, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v160 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(2)
	v166 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+72)))
	v169 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), v166, int64(0), int64(-1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v179 = v177 & int32(1073741823)
	if v177 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v172 = F_pg_detoast_datum(m, base.I32_wrap_i64(v169))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v172
	goto L15
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v192 + v41
	goto L15
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v179
	v192 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v185 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v179 - v184&v185
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v192 = v189 & v185
	goto L45
L49:
	;
	v205 = v47 + int32(1)
	if v205 != v36 {
		v47 = v205
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L9
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
	v222 = F_JsonbValueToJsonb(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	m.G0 = v11 + int32(112)
	return base.I64_extend_i32_u(v222)
}
