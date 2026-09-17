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
func F_hstore_defined(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v21 = int32(1)
	v22 = v19 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v25 = v23 & v21
	if v23 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v56 = v54 & int32(268435455)
	if v56 == v53 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v31 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v42 = int32(1)
	if v25 != 0 {
		v52 = int32(base.Ui32(v23)>>(uint(v42)%32)) - v42
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v34 = int32(16)
	goto L10
L9:
	;
	v34 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = int32(4)
	goto L13
L12:
	;
	v41 = v34
	goto L13
L13:
	;
	v52 = v41
	goto L4
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	return int32(0)
L16:
	;
	goto L17
L17:
	;
	if v25 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = v22
	goto L20
L19:
	;
	v63 = v19 + int32(4)
	goto L20
L20:
	;
	v65 = v14 + int32(8)
	v69 = v53
	v74 = v56
	goto L21
L21:
	;
	v84 = int32(base.Ui32(v74-v69)>>(uint(int32(1))%32)) + v69
	v86 = v84 << (uint(int32(3)) % 32)
	v87 = v65 + v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v90 = v88 & int32(1073741823)
	if int32(0) <= v88 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	return int32(0)
L23:
	;
	v187 = base.B2i32(v182 < int32(0))
	if v182 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L24:
	;
	v110 = v109 + (v65 + v56<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if base.Ui32(v52) < base.Ui32(v102) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87-int32(4))))
	v97 = v95 & int32(1073741823)
	v98 = v90 - v97
	if v98 != v52 {
		v102 = v98
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v90 == v52 {
		v109 = int32(0)
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v109 = v97
	goto L24
L30:
	;
	v102 = v90
	goto L25
L31:
	;
	v107 = int32(1)
	goto L33
L32:
	;
	v107 = int32(-1)
	goto L33
L33:
	;
	v182 = v107
	goto L23
L34:
	;
	if v172 != 0 {
		v182 = v172
		goto L23
	} else {
		goto L52
	}
L35:
	;
	v172 = int32(0)
	goto L34
L36:
	;
	v146 = v141
	v147 = v142
	v148 = v143
	goto L46
L37:
	;
	if (v110|v63)&int32(3) != 0 {
		v141 = v110
		v142 = v63
		v143 = v52
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v134 = v110
	v135 = v63
	v136 = v52
	goto L39
L39:
	;
	if v136 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v118 = v110
	v119 = v63
	v120 = v52
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v123 != v124 {
		v141 = v118
		v142 = v119
		v143 = v120
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v134 = v129
	v135 = v127
	v136 = v131
	goto L39
L43:
	;
	v126 = int32(4)
	v127 = v119 + v126
	v129 = v118 + v126
	v131 = v120 - v126
	if base.Ui32(int32(3)) < base.Ui32(v131) {
		v118 = v129
		v119 = v127
		v120 = v131
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v141 = v134
	v142 = v135
	v143 = v136
	goto L36
L46:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v151 == v152 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v172 = v151 - v152
	goto L34
L48:
	;
	v154 = int32(1)
	v159 = v148 - v154
	if v159 != 0 {
		v146 = v146 + v154
		v147 = v147 + v154
		v148 = v159
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
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v14+v86)+12))
	return int32(base.Ui32(v174^int32(-1))>>(uint(int32(30))%32)) & int32(1)
L53:
	;
	v188 = v84 + int32(1)
	goto L55
L54:
	;
	v188 = v69
	goto L55
L55:
	;
	if v182 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v189 = v74
	goto L58
L57:
	;
	v189 = v84
	goto L58
L58:
	;
	if v188 < v189 {
		v69 = v188
		v74 = v189
		goto L21
	} else {
		goto L59
	}
L59:
	;
	goto L22
}
func F_hstore_delete_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return v31
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
func F_hstore_delete_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v508 int32
	_ = v508
	v2 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v494 == v508&int32(268435455) {
		goto L121
	} else {
		goto L122
	}
L6:
	;
	v73 = int32(8)
	v74 = v30 + v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v76 = int32(3)
	v78 = int32(2147483640)
	v80 = v74 + v75<<(uint(v76)%32)&v78
	v82 = v25 + v73
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v88 = v82 + v83<<(uint(v76)%32)&v78
	v90 = v49
	v95 = v52
	v96 = int32(0)
	v99 = v2
	v102 = v2
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
	v491 = int32(0)
	v494 = v2
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
	return v35
L17:
	;
	v114 = v96 << (uint(int32(3)) % 32)
	if v54 <= v102 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v477 = v467 - v52
	if v470 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L19:
	;
	if v468 < v44 {
		v90 = v462
		v95 = v467
		v96 = v468
		v99 = v470
		v102 = v471
		goto L17
	} else {
		goto L117
	}
L20:
	;
	v462 = v90
	v467 = v95
	v468 = v96
	v470 = v99
	v471 = v102 + int32(1)
	goto L19
L21:
	;
	v388 = v114 + v82
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	if v389 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L22:
	;
	v116 = v114 + v82
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v118 = int32(1073741823)
	v119 = v117 & v118
	v122 = v74 + v102<<(uint(int32(3))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = v123 & v118
	if int32(0) <= v117 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v134 = v119 - v130&int32(1073741823)
	goto L25
L24:
	;
	v134 = v119
	goto L25
L25:
	;
	if int32(0) <= v123 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v122-int32(4))))
	v143 = v125 - v139&int32(1073741823)
	goto L28
L27:
	;
	v143 = v125
	goto L28
L28:
	;
	if v143 == v134 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if int32(0) <= v117 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v143 < v134 {
		goto L20
	} else {
		goto L100
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v153 = v149 & int32(1073741823)
	goto L34
L33:
	;
	v153 = int32(0)
	goto L34
L34:
	;
	v154 = v153 + v88
	if int32(0) <= v123 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v122-int32(4))))
	v163 = v159 & int32(1073741823)
	goto L37
L36:
	;
	v163 = int32(0)
	goto L37
L37:
	;
	v164 = v163 + v80
	if base.Ui32(int32(4)) <= base.Ui32(v134) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if int32(0) < v226 {
		goto L20
	} else {
		goto L56
	}
L39:
	;
	v226 = int32(0)
	goto L38
L40:
	;
	v200 = v195
	v201 = v196
	v202 = v197
	goto L50
L41:
	;
	if (v154|v164)&int32(3) != 0 {
		v195 = v154
		v196 = v164
		v197 = v134
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v188 = v154
	v189 = v164
	v190 = v134
	goto L43
L43:
	;
	if v190 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v172 = v154
	v173 = v164
	v174 = v134
	goto L45
L45:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v177 != v178 {
		v195 = v172
		v196 = v173
		v197 = v174
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v188 = v183
	v189 = v181
	v190 = v185
	goto L43
L47:
	;
	v180 = int32(4)
	v181 = v173 + v180
	v183 = v172 + v180
	v185 = v174 - v180
	if base.Ui32(int32(3)) < base.Ui32(v185) {
		v172 = v183
		v173 = v181
		v174 = v185
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v195 = v188
	v196 = v189
	v197 = v190
	goto L40
L50:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v205 == v206 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v226 = v205 - v206
	goto L38
L52:
	;
	v208 = int32(1)
	v213 = v202 - v208
	if v213 != 0 {
		v200 = v200 + v208
		v201 = v201 + v208
		v202 = v213
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
	if v226 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v230 = int32(1073741823)
	v233 = v117 & v230
	v234 = int32(0)
	v236 = base.B2i32(v234 <= v229)
	if v234 <= v229 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v237 = v233
	goto L60
L59:
	;
	v237 = v234
	goto L60
L60:
	;
	v238 = v229&v230 - v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v240 = int32(30)
	v245 = v229 & int32(1073741824)
	if int32(base.Ui32(v239)>>(uint(v240)%32))&int32(1) != int32(base.Ui32(v245)>>(uint(v240)%32)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v377 = int32(1)
	v462 = v371
	v467 = v374
	v468 = v96 + v377
	v470 = v375
	v471 = v102 + v377
	goto L19
L62:
	;
	if v117 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L63:
	;
	if v245 != 0 {
		v371 = v90
		v374 = v95
		v375 = v99
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v249 = int32(1073741823)
	v253 = int32(0)
	if v253 <= v239 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v256 = v123 & v249
	goto L67
L66:
	;
	v256 = v253
	goto L67
L67:
	;
	if v238 != v239&v249-v256 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v234 <= v229 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v260 = v233
	goto L71
L70:
	;
	v260 = int32(0)
	goto L71
L71:
	;
	v261 = v88 + v260
	v262 = v256 + v80
	if base.Ui32(int32(4)) <= base.Ui32(v238) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	if v324 == int32(0) {
		v371 = v90
		v374 = v95
		v375 = v99
		goto L61
	} else {
		goto L90
	}
L73:
	;
	v324 = int32(0)
	goto L72
L74:
	;
	v298 = v293
	v299 = v294
	v300 = v295
	goto L84
L75:
	;
	if (v261|v262)&int32(3) != 0 {
		v293 = v261
		v294 = v262
		v295 = v238
		goto L74
	} else {
		goto L78
	}
L76:
	;
	v286 = v261
	v287 = v262
	v288 = v238
	goto L77
L77:
	;
	if v288 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L78:
	;
	v270 = v261
	v271 = v262
	v272 = v238
	goto L79
L79:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v275 != v276 {
		v293 = v270
		v294 = v271
		v295 = v272
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v286 = v281
	v287 = v279
	v288 = v283
	goto L77
L81:
	;
	v278 = int32(4)
	v279 = v271 + v278
	v281 = v270 + v278
	v283 = v272 - v278
	if base.Ui32(int32(3)) < base.Ui32(v283) {
		v270 = v281
		v271 = v279
		v272 = v283
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v293 = v286
	v294 = v287
	v295 = v288
	goto L74
L84:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v303 == v304 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v324 = v303 - v304
	goto L72
L86:
	;
	v306 = int32(1)
	v311 = v300 - v306
	if v311 != 0 {
		v298 = v298 + v306
		v299 = v299 + v306
		v300 = v311
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
	v340 = v233
	v341 = v88
	goto L93
L92:
	;
	v330 = int32(1073741823)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v336 = v334 & v330
	v340 = v117&v330 - v336
	v341 = v336 + v88
	goto L93
L93:
	;
	v342 = v238 + v340
	if v342 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v95, v341, v342)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v346 = v344 & int32(1073741823)
	if int32(0) <= v344 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v355 = v346 - v351&int32(1073741823)
	goto L99
L98:
	;
	v355 = v346
	goto L99
L99:
	;
	v357 = v355 + v95 + v238
	v358 = v357 - v52
	v359 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v358&v359 | v245
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = (v358 - v238) & v359
	v371 = v90 + int32(8)
	v374 = v357
	v375 = v99 + int32(1)
	goto L61
L100:
	;
	goto L21
L101:
	;
	v408 = v388 + int32(4)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v412 = int32(0)
	if v412 <= v409 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v393 = v389 & int32(1073741823)
	v403 = v393
	v404 = v393
	v406 = v88
	goto L101
L103:
	;
	goto L104
L104:
	;
	v394 = int32(1073741823)
	v395 = v389 & v394
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v388-int32(4))))
	v400 = v398 & v394
	v403 = v395
	v404 = v395 - v400
	v406 = v400 + v88
	goto L101
L105:
	;
	v415 = v403
	goto L107
L106:
	;
	v415 = v412
	goto L107
L107:
	;
	v417 = v404 + (v409&int32(1073741823) - v415)
	if v417 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	base.MemoryCopy(m, v95, v406, v417)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v421 = v419 & int32(1073741823)
	if int32(0) <= v419 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v388-int32(4))))
	v430 = v421 - v426&int32(1073741823)
	goto L113
L112:
	;
	v430 = v421
	goto L113
L113:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v434 = int32(0)
	if v434 <= v431 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v437 = v421
	goto L116
L115:
	;
	v437 = v434
	goto L116
L116:
	;
	v438 = v431&int32(1073741823) - v437
	v440 = v438 + (v430 + v95)
	v441 = v440 - v52
	v443 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = (v441 - v438) & v443
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v446&int32(1073741824) | v441&v443
	v453 = int32(1)
	v462 = v90 + int32(8)
	v467 = v440
	v468 = v96 + v453
	v470 = v99 + v453
	v471 = v102
	goto L19
L117:
	;
	goto L18
L118:
	;
	v491 = v477
	v494 = int32(0)
	goto L5
L119:
	;
	goto L120
L120:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v481 | int32(-2147483648)
	v491 = v477
	v494 = v470
	goto L5
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = (v494<<(uint(int32(3))%32)+v491)<<(uint(int32(2))%32) + int32(32)
	return v35
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v494 | int32(-2147483648)
	if v491 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	base.MemoryCopy(m, v49+v494<<(uint(int32(3))%32)&int32(2147483640), v52, v491)
	goto L121
}
func F_hstore_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = int32(16)
	v18 = F_palloc(m, int32(320))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v11
	v31 = v11
	v32 = v22
	goto L4
L3:
	;
	m.G0 = v9 + int32(96)
	return v400
L4:
	;
	switch v32 - int32(1) {
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
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v390 = F_hstoreUniquePairs(m, v386, v387, v9+int32(56))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L113
	}
L6:
	;
	goto L5
L7:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	v382 = v380 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v382
	v31 = v382
	v32 = v379
	goto L4
L8:
	;
	v379 = int32(2)
	goto L7
L9:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v363 = v360 + v53*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+8)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v66
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+4)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v366
	goto L8
L10:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v357)
	v400 = int32(0)
	goto L3
L11:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v312 == int32(44) {
		v379 = int32(0)
		goto L7
	} else {
		goto L103
	}
L12:
	;
	v196 = F_get_val(m, v9+int32(60), int32(1), v9+int32(56))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v144 == int32(62) {
		v379 = int32(1)
		goto L7
	} else {
		goto L50
	}
L14:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v88 == int32(61) {
		v379 = int32(3)
		goto L7
	} else {
		goto L33
	}
L15:
	;
	v42 = F_get_val(m, v9+int32(60), int32(0), v9+int32(56))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v42 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	if v46 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
	if v54 <= v53 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 != int32(447) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v52 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v54 << (uint(int32(1)) % 32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v62 = F_repalloc(m, v59, v54*int32(40))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	v67 = v65 - v66
	if base.Ui32(v67) < base.Ui32(int32(1073741824)) {
		goto L9
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v62
	goto L25
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v71 = F_errsave_start(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v71 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_hstore_in_0), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, v70, int32(_a_F_hstore_in_1), int32(423), int32(_a_F_hstore_in_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	if v88 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v94 = F_errsave_start(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v110 = base.I32_extend8_s(v88)
	goto L42
L37:
	;
	if v94 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errsave_finish(m, v93, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L10
L42:
	;
	if base.B2i32(v110 == int32(32))|base.B2i32(base.Ui32((v110-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v121 = F_errsave_start(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v121 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v128 = F_pg_mblen_cstr(m, v31)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v31
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31 - v132
	F_errmsg(m, int32(_a_F_hstore_in_5), v9)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errsave_finish(m, v120, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L10
L50:
	;
	if v144 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v150 = F_errsave_start(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v167 = F_errsave_start(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	if v150 == int32(0) {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, v149, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	if v167 == int32(0) {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v174 = F_pg_mblen_cstr(m, v31)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v31
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v31 - v178
	F_errmsg(m, int32(_a_F_hstore_in_5), v9+int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errsave_finish(m, v166, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	if v196 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	if v200 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	v225 = v223 - v224
	if base.Ui32(int32(1073741824)) <= base.Ui32(v225) {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	v207 = F_errsave_start(m, v200)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v203 != int32(447) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+4)))
	if v206 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v207 == int32(0) {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_hstore_in_3), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errsave_finish(m, v200, int32(_a_F_hstore_in_1), int32(83), int32(_a_F_hstore_in_4))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L10
L78:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v229 = F_errsave_start(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v249 = v245 + v246*int32(20)
	v250 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+16)) = uint16(v250)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+12)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v224
	if v225 != int32(4) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	if v229 == int32(0) {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_hstore_in_7), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errsave_finish(m, v228, int32(_a_F_hstore_in_1), int32(443), int32(_a_F_hstore_in_8))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L10
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v246 + int32(1)
	v379 = int32(4)
	goto L7
L87:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)))
	if v256&int32(1) != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+4)) = uint8(v259)
	v264 = v224
	v265 = int32(_a_F_hstore_in_9)
	goto L90
L89:
	;
	if v302 != 0 {
		goto L86
	} else {
		goto L102
	}
L90:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v268 == v269 {
		v291 = v268
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v302 = int32(0)
	goto L89
L92:
	;
	v293 = int32(1)
	if v291 != 0 {
		v264 = v264 + v293
		v265 = v265 + v293
		goto L90
	} else {
		goto L101
	}
L93:
	;
	if base.Ui32((v268-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v279 = v268 | int32(32)
	goto L96
L95:
	;
	v279 = v268
	goto L96
L96:
	;
	if base.Ui32((v269-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v288 = v269 | int32(32)
	goto L99
L98:
	;
	v288 = v269
	goto L99
L99:
	;
	if v279 == v288 {
		v291 = v279
		goto L92
	} else {
		goto L100
	}
L100:
	;
	v302 = v279 - v288
	goto L89
L101:
	;
	goto L91
L102:
	;
	v303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v249)+16)) = uint8(v303)
	goto L86
L103:
	;
	if v312 == int32(0) {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v318 = base.I32_extend8_s(v312)
	goto L105
L105:
	;
	if base.B2i32(v318 == int32(32))|base.B2i32(base.Ui32((v318-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v379 = int32(4)
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v329 = F_errsave_start(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v329 == int32(0) {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v336 = F_pg_mblen_cstr(m, v31)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v31
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v31 - v340
	F_errmsg(m, int32(_a_F_hstore_in_5), v9+int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errsave_finish(m, v328, int32(_a_F_hstore_in_1), int32(71), int32(_a_F_hstore_in_6))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L10
L113:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v393 = F_hstorePairs(m, v386, v390, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v400 = v393
	goto L3
}
func F_hstore_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_le_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 <= int32(0))
	}
}
func F_hstore_skeys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = F_hstoreUpgrade(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
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
					return int32(0)
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
							return int32(0)
						} else {
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
							*(*int64)(unsafe.Add(mBase, uint32(v41))) = v75 + int64(1)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = int32(1)
							return v73
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = int32(2)
							v88 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v88)
							return int32(0)
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
				return int32(0)
			} else {
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v75 + int64(1)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = int32(1)
				return v73
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = int32(2)
				v88 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v88)
				return int32(0)
			}
		}
	}
}
func F_hstore_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_hstoreUpgrade(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_hstore_to_array_internal(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_hstore_to_json_loose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return v169
L2:
	;
	return int32(0)
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
func F_hstore_to_jsonb_loose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
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
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	F_initStringInfo(m, v9+int32(-20))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_pushJsonbValue(m, v9+int32(-4), int32(6), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = v18 & int32(268435455)
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = v14 + int32(8)
	v37 = v34 + v32<<(uint(int32(3))%32)
	v43 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v214 = F_pushJsonbValue(m, v9+int32(-4), int32(7), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1)
	v50 = v34 + v43<<(uint(int32(3))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v53 = v51 & int32(1073741823)
	if v51 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v69 + v37
	v77 = F_pushJsonbValue(m, v9+int32(-4), int32(1), v9+int32(-40))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
	v69 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v59 = v50 - int32(4)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53 - v60&v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v69 = v65 & v61
	goto L10
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v79&int32(1073741824) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v197 = F_pushJsonbValue(m, v9+int32(-4), int32(2), v9+int32(-60))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L49
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if int32(0) <= v79 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v129 = v9 + int32(-20)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v131
	goto L33
L20:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v37))))
	if v119 != int32(102) {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(3)
	goto L15
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v88 = v86 & int32(1073741823)
	if v79-v88 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v79&int32(1073741823) != int32(1) {
		goto L19
	} else {
		goto L30
	}
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v88))))
	if v93 == int32(116) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v98 = v96 & int32(1073741823)
	if v79-v98 != int32(1) {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v116 = v98
	goto L20
L30:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v107 != int32(116) {
		v116 = int32(0)
		goto L20
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(3)
	goto L15
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v137 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_appendBinaryStringInfo(m, v129, v151, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v149 = v137 & int32(1073741823)
	v151 = v37
	goto L34
L36:
	;
	goto L37
L37:
	;
	v142 = int32(1073741823)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v146 = v144 & v142
	v149 = v137&v142 - v146
	v151 = v37 + v146
	goto L34
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v156 = F_IsValidJsonNumber(m, v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v156 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(2)
	v161 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v165 = F_DirectFunctionCall3Coll(m, int32(408), v161, v162, v161, int32(-1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v174 = v172 & int32(1073741823)
	if v172 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v167 = F_pg_detoast_datum(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v167
	goto L15
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v187 + v37
	goto L15
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v174
	v187 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v180 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v174 - v179&v180
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v187 = v184 & v180
	goto L45
L49:
	;
	v200 = v43 + int32(1)
	if v200 != v32 {
		v43 = v200
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L9
L51:
	;
	v216 = F_JsonbValueToJsonb(m, v214)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	m.G0 = v11 - int32(-64)
	return v216
}
