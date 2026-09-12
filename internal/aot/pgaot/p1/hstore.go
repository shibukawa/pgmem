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
	var v37 int32
	_ = v37
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
	v37 = base.I32_div_s(v30-v29, int32(2))
	v38 = v37 + v29
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
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
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v57 = v55 & int32(268435455)
	if v57 == v54 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v30&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v39 = v28
	goto L10
L9:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L10
L10:
	;
	if v30 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = v28
	goto L13
L12:
	;
	v42 = v39
	goto L13
L13:
	;
	v53 = v42
	goto L4
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
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
	v64 = v22
	goto L20
L19:
	;
	v64 = v19 + int32(4)
	goto L20
L20:
	;
	v66 = v14 + int32(8)
	v70 = v54
	v75 = v57
	goto L21
L21:
	;
	v84 = base.I32_div_s(v75-v70, int32(2))
	v85 = v84 + v70
	v87 = v85 << (uint(int32(3)) % 32)
	v88 = v66 + v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v91 = v89 & int32(1073741823)
	if int32(0) <= v89 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	return v198
L23:
	;
	goto L22
L24:
	;
	v188 = int32(0)
	v192 = base.B2i32(v186 < v188)
	if v186 < v188 {
		goto L55
	} else {
		goto L56
	}
L25:
	;
	v111 = v110 + (v66 + v57<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	if base.Ui32(v53) < base.Ui32(v103) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(4))))
	v98 = v96 & int32(1073741823)
	v99 = v91 - v98
	if v99 != v53 {
		v103 = v99
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v91 == v53 {
		v110 = int32(0)
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v110 = v98
	goto L25
L31:
	;
	v103 = v91
	goto L26
L32:
	;
	v108 = int32(1)
	goto L34
L33:
	;
	v108 = int32(-1)
	goto L34
L34:
	;
	v186 = v108
	goto L24
L35:
	;
	if v173 != 0 {
		v186 = v173
		goto L24
	} else {
		goto L53
	}
L36:
	;
	v173 = int32(0)
	goto L35
L37:
	;
	v147 = v142
	v148 = v143
	v149 = v144
	goto L47
L38:
	;
	if (v111|v64)&int32(3) != 0 {
		v142 = v111
		v143 = v64
		v144 = v53
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v135 = v111
	v136 = v64
	v137 = v53
	goto L40
L40:
	;
	if v137 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v119 = v111
	v120 = v64
	v121 = v53
	goto L42
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v124 != v125 {
		v142 = v119
		v143 = v120
		v144 = v121
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v135 = v130
	v136 = v128
	v137 = v132
	goto L40
L44:
	;
	v127 = int32(4)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L37
L47:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 == v153 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v173 = v152 - v153
	goto L35
L49:
	;
	v155 = int32(1)
	v160 = v149 - v155
	if v160 != 0 {
		v147 = v147 + v155
		v148 = v148 + v155
		v149 = v160
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	v174 = int32(0)
	if v85 < v174 {
		v198 = v174
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14+v87)+12))
	return int32(base.Ui32(v178^int32(-1))>>(uint(int32(30))%32)) & int32(1)
L55:
	;
	v193 = v85 + int32(1)
	goto L57
L56:
	;
	v193 = v70
	goto L57
L57:
	;
	if v186 < v188 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v194 = v75
	goto L60
L59:
	;
	v194 = v85
	goto L60
L60:
	;
	if v193 < v194 {
		v70 = v193
		v75 = v194
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v198 = v188
	goto L23
}
func F_hstore_delete_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v32 = F_palloc(m, int32(base.Ui32(v29)>>(uint(int32(2))%32)))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = F_hstoreArrayToPairs(m, v36, v22+int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v44 = v34 & int32(268435455)
	v46 = v44 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v42 & int32(-4)
	v52 = v32 + int32(8)
	v54 = v44 << (uint(int32(3)) % 32)
	v55 = v52 + v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	m.G0 = v22 + int32(16)
	return v32
L7:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v338&int32(268435455) != v324 {
		goto L75
	} else {
		goto L76
	}
L8:
	;
	v76 = v25 + int32(8)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v82 = v76 + v77<<(uint(int32(3))%32)&int32(2147483640)
	v84 = v52
	v85 = int32(0)
	v89 = v2
	v91 = v55
	v92 = v2
	goto L20
L9:
	;
	if v44 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v60 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v320 = int32(0)
	v324 = v2
	goto L7
L13:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v61 = F__emscripten_memcpy_bulkmem(m, v32, v25, v60)
	mBase = m.M
	v62 = v61
	goto L16
L15:
	;
	v62 = v32
	goto L16
L16:
	;
	goto L13
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(4))))
	v67 = v65
	goto L19
L18:
	;
	v67 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = (v54+v67)<<(uint(int32(2))%32) + int32(32)
	goto L6
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v103 <= v92 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v311 = v305 - v55
	if v304 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L22:
	;
	if v300 < v44 {
		v84 = v299
		v85 = v300
		v89 = v304
		v91 = v305
		v92 = v306
		goto L20
	} else {
		goto L71
	}
L23:
	;
	v221 = v76 + v85<<(uint(int32(3))%32)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v222 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L24:
	;
	v107 = v76 + v85<<(uint(int32(3))%32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v110 = v108 & int32(1073741823)
	if int32(0) <= v108 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v299 = v84
	v300 = v85
	v304 = v89
	v305 = v91
	v306 = v92 + int32(1)
	goto L22
L26:
	;
	v139 = v136 + v82
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if base.Ui32(int32(4)) <= base.Ui32(v135) {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	if base.Ui32(v130) <= base.Ui32(v133) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(4))))
	v117 = v115 & int32(1073741823)
	v118 = v110 - v117
	v121 = v40 + v92*int32(20)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v118 != v122 {
		v130 = v118
		v133 = v122
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v127 = v40 + v92*int32(20)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	if v110 == v128 {
		v135 = v110
		v136 = int32(0)
		v137 = v127
		goto L26
	} else {
		goto L32
	}
L31:
	;
	v135 = v118
	v136 = v117
	v137 = v121
	goto L26
L32:
	;
	v130 = v110
	v133 = v128
	goto L27
L33:
	;
	goto L25
L34:
	;
	if int32(0) < v202 {
		goto L25
	} else {
		goto L52
	}
L35:
	;
	v202 = int32(0)
	goto L34
L36:
	;
	v176 = v171
	v177 = v172
	v178 = v173
	goto L46
L37:
	;
	if (v139|v140)&int32(3) != 0 {
		v171 = v139
		v172 = v140
		v173 = v135
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v164 = v139
	v165 = v140
	v166 = v135
	goto L39
L39:
	;
	if v166 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v148 = v139
	v149 = v140
	v150 = v135
	goto L41
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v153 != v154 {
		v171 = v148
		v172 = v149
		v173 = v150
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v164 = v159
	v165 = v157
	v166 = v161
	goto L39
L43:
	;
	v156 = int32(4)
	v157 = v149 + v156
	v159 = v148 + v156
	v161 = v150 - v156
	if base.Ui32(int32(3)) < base.Ui32(v161) {
		v148 = v159
		v149 = v157
		v150 = v161
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v171 = v164
	v172 = v165
	v173 = v166
	goto L36
L46:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v181 == v182 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v202 = v181 - v182
	goto L34
L48:
	;
	v184 = int32(1)
	v189 = v178 - v184
	if v189 != 0 {
		v176 = v176 + v184
		v177 = v177 + v184
		v178 = v189
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
	if v202 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v205 = int32(1)
	v299 = v84
	v300 = v85 + v205
	v304 = v89
	v305 = v91
	v306 = v92 + v205
	goto L22
L54:
	;
	v243 = v76 + v85<<(uint(int32(3))%32) + int32(4)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v246 = v244 & int32(1073741823)
	if int32(0) <= v244 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v237 = v222 & int32(1073741823)
	v238 = v82
	goto L54
L56:
	;
	goto L57
L57:
	;
	v227 = int32(1073741823)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221-int32(4))))
	v233 = v231 & v227
	v237 = v222&v227 - v233
	v238 = v233 + v82
	goto L54
L58:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v243-int32(4))))
	v255 = v246 - v251&int32(1073741823)
	goto L60
L59:
	;
	v255 = v246
	goto L60
L60:
	;
	v256 = v255 + v237
	if v256 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v261 = v259 & int32(1073741823)
	if int32(0) <= v259 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v257 = F__emscripten_memcpy_bulkmem(m, v91, v238, v256)
	mBase = m.M
	v258 = v257
	goto L64
L63:
	;
	v258 = v91
	goto L64
L64:
	;
	goto L61
L65:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v221-int32(4))))
	v270 = v261 - v266&int32(1073741823)
	goto L67
L66:
	;
	v270 = v261
	goto L67
L67:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v274 = int32(0)
	if v274 <= v271 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v277 = v261
	goto L70
L69:
	;
	v277 = v274
	goto L70
L70:
	;
	v278 = v271&int32(1073741823) - v277
	v280 = v278 + (v270 + v258)
	v281 = v280 - v55
	v283 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = (v281 - v278) & v283
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v286&int32(1073741824) | v281&v283
	v293 = int32(1)
	v299 = v84 + int32(8)
	v300 = v85 + v293
	v304 = v89 + v293
	v305 = v280
	v306 = v92
	goto L22
L71:
	;
	goto L21
L72:
	;
	v320 = v311
	v324 = int32(0)
	goto L7
L73:
	;
	goto L74
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v315 | int32(-2147483648)
	v320 = v311
	v324 = v304
	goto L7
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v324 | int32(-2147483648)
	v349 = v52 + v324<<(uint(int32(3))%32)&int32(2147483640)
	if v349 == v55 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = (v324<<(uint(int32(3))%32)+v320)<<(uint(int32(2))%32) + int32(32)
	goto L6
L78:
	;
	goto L77
L79:
	;
	goto L78
L80:
	;
	v353 = v349 + v320
	if base.Ui32(v55-v353) <= base.Ui32(int32(0)-v320<<(uint(int32(1))%32)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v360 = F___memcpy(m, v349, v55, v320)
	mBase = m.M
	goto L78
L82:
	;
	goto L83
L83:
	;
	v363 = (v349 ^ v55) & int32(3)
	if base.Ui32(v349) < base.Ui32(v55) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v465 == int32(0) {
		goto L79
	} else {
		goto L120
	}
L85:
	;
	if base.Ui32(v443) <= base.Ui32(int32(3)) {
		v464 = v442
		v465 = v443
		v466 = v444
		goto L84
	} else {
		goto L116
	}
L86:
	;
	if v363 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v363 != 0 {
		v425 = v320
		goto L99
	} else {
		goto L100
	}
L89:
	;
	v464 = v55
	v465 = v320
	v466 = v349
	goto L84
L90:
	;
	goto L91
L91:
	;
	if v349&int32(3) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v442 = v55
	v443 = v320
	v444 = v349
	goto L85
L93:
	;
	goto L94
L94:
	;
	v370 = v55
	v371 = v320
	v372 = v349
	goto L95
L95:
	;
	if v371 == int32(0) {
		goto L79
	} else {
		goto L97
	}
L96:
	;
	v442 = v379
	v443 = v381
	v444 = v383
	goto L85
L97:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	*(*uint8)(unsafe.Add(mBase, uint32(v372))) = uint8(v376)
	v378 = int32(1)
	v379 = v370 + v378
	v381 = v371 - v378
	v383 = v372 + v378
	if v383&int32(3) != 0 {
		v370 = v379
		v371 = v381
		v372 = v383
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if v425 == int32(0) {
		goto L79
	} else {
		goto L112
	}
L100:
	;
	if v353&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v390 = v320
	goto L104
L102:
	;
	v405 = v320
	goto L103
L103:
	;
	if base.Ui32(v405) <= base.Ui32(int32(3)) {
		v425 = v405
		goto L99
	} else {
		goto L108
	}
L104:
	;
	if v390 == int32(0) {
		goto L79
	} else {
		goto L106
	}
L105:
	;
	v405 = v396
	goto L103
L106:
	;
	v396 = v390 - int32(1)
	v397 = v349 + v396
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v396))))
	*(*uint8)(unsafe.Add(mBase, uint32(v397))) = uint8(v399)
	if v397&int32(3) != 0 {
		v390 = v396
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v412 = v405
	goto L109
L109:
	;
	v416 = v412 - int32(4)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v55+v416)))
	*(*int32)(unsafe.Add(mBase, uint32(v349+v416))) = v419
	if base.Ui32(int32(3)) < base.Ui32(v416) {
		v412 = v416
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v425 = v416
	goto L99
L111:
	;
	goto L110
L112:
	;
	v432 = v425
	goto L113
L113:
	;
	v436 = v432 - int32(1)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v436))))
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v436))) = uint8(v439)
	if v436 != 0 {
		v432 = v436
		goto L113
	} else {
		goto L115
	}
L114:
	;
	goto L79
L115:
	;
	goto L114
L116:
	;
	v449 = v442
	v450 = v443
	v451 = v444
	goto L117
L117:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v453
	v455 = int32(4)
	v456 = v449 + v455
	v458 = v451 + v455
	v460 = v450 - v455
	if base.Ui32(int32(3)) < base.Ui32(v460) {
		v449 = v456
		v450 = v460
		v451 = v458
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v464 = v456
	v465 = v460
	v466 = v458
	goto L84
L119:
	;
	goto L118
L120:
	;
	v471 = v464
	v472 = v465
	v473 = v466
	goto L121
L121:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v475)
	v477 = int32(1)
	v482 = v472 - v477
	if v482 != 0 {
		v471 = v471 + v477
		v472 = v482
		v473 = v473 + v477
		goto L121
	} else {
		goto L123
	}
L122:
	;
	goto L79
L123:
	;
	goto L122
}
func F_hstore_delete_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	v2 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_hstoreUpgrade(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = F_hstoreUpgrade(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v33 = F_palloc(m, int32(base.Ui32(v30)>>(uint(int32(2))%32)))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v37 & int32(-4)
	v41 = int32(268435455)
	v42 = v36 & v41
	v44 = v42 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v44
	v47 = v33 + int32(8)
	v49 = v42 << (uint(int32(3)) % 32)
	v50 = v47 + v49
	v52 = v35 & v41
	if v52 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v544&int32(268435455) != v529 {
		goto L128
	} else {
		goto L129
	}
L6:
	;
	v72 = int32(8)
	v73 = v28 + v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v75 = int32(3)
	v77 = int32(2147483640)
	v79 = v73 + v74<<(uint(v75)%32)&v77
	v81 = v23 + v72
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v87 = v81 + v82<<(uint(v75)%32)&v77
	v89 = v47
	v92 = int32(0)
	v95 = v2
	v96 = v50
	v103 = v2
	goto L18
L7:
	;
	if v42 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v56 = int32(base.Ui32(v54) >> (uint(int32(2)) % 32))
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v526 = int32(0)
	v529 = v2
	goto L5
L11:
	;
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v57 = F__emscripten_memcpy_bulkmem(m, v33, v23, v56)
	mBase = m.M
	v58 = v57
	goto L14
L13:
	;
	v58 = v33
	goto L14
L14:
	;
	goto L11
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50-int32(4))))
	v63 = v61
	goto L17
L16:
	;
	v63 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = (v63+v49)<<(uint(int32(2))%32) + int32(32)
	return v58
L18:
	;
	if v52 <= v103 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v515 = v509 - v50
	if v508 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L20:
	;
	if v505 < v42 {
		v89 = v502
		v92 = v505
		v95 = v508
		v96 = v509
		v103 = v513
		goto L18
	} else {
		goto L124
	}
L21:
	;
	v502 = v89
	v505 = v92
	v508 = v95
	v509 = v96
	v513 = v103 + int32(1)
	goto L20
L22:
	;
	v419 = v81 + v92<<(uint(int32(3))%32)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v420 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L23:
	;
	v111 = int32(3)
	v113 = v81 + v92<<(uint(v111)%32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = int32(1073741823)
	v116 = v114 & v115
	v119 = v73 + v103<<(uint(v111)%32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v122 = v120 & v115
	if int32(0) <= v114 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(4))))
	v131 = v116 - v127&int32(1073741823)
	goto L26
L25:
	;
	v131 = v116
	goto L26
L26:
	;
	if int32(0) <= v120 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v119-int32(4))))
	v140 = v122 - v136&int32(1073741823)
	goto L29
L28:
	;
	v140 = v122
	goto L29
L29:
	;
	if v140 == v131 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v142 = int32(0)
	if v142 <= v114 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v140 < v131 {
		goto L21
	} else {
		goto L106
	}
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(4))))
	v151 = v148 & int32(1073741823)
	goto L35
L34:
	;
	v151 = v142
	goto L35
L35:
	;
	v152 = v151 + v87
	if int32(0) <= v120 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v119-int32(4))))
	v160 = v157 & int32(1073741823)
	goto L38
L37:
	;
	v160 = v142
	goto L38
L38:
	;
	v161 = v160 + v79
	if base.Ui32(int32(4)) <= base.Ui32(v131) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if int32(0) < v223 {
		goto L21
	} else {
		goto L57
	}
L40:
	;
	v223 = int32(0)
	goto L39
L41:
	;
	v197 = v192
	v198 = v193
	v199 = v194
	goto L51
L42:
	;
	if (v152|v161)&int32(3) != 0 {
		v192 = v152
		v193 = v161
		v194 = v131
		goto L41
	} else {
		goto L45
	}
L43:
	;
	v185 = v152
	v186 = v161
	v187 = v131
	goto L44
L44:
	;
	if v187 == int32(0) {
		goto L40
	} else {
		goto L50
	}
L45:
	;
	v169 = v152
	v170 = v161
	v171 = v131
	goto L46
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v174 != v175 {
		v192 = v169
		v193 = v170
		v194 = v171
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v185 = v180
	v186 = v178
	v187 = v182
	goto L44
L48:
	;
	v177 = int32(4)
	v178 = v170 + v177
	v180 = v169 + v177
	v182 = v171 - v177
	if base.Ui32(int32(3)) < base.Ui32(v182) {
		v169 = v180
		v170 = v178
		v171 = v182
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v192 = v185
	v193 = v186
	v194 = v187
	goto L41
L51:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v202 == v203 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v223 = v202 - v203
	goto L39
L53:
	;
	v205 = int32(1)
	v210 = v199 - v205
	if v210 != 0 {
		v197 = v197 + v205
		v198 = v198 + v205
		v199 = v210
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L40
L57:
	;
	if v223 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	v227 = v113 + int32(4)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v230 = v228 & int32(1073741823)
	if int32(0) <= v228 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227-int32(4))))
	v239 = v230 - v235&int32(1073741823)
	goto L61
L60:
	;
	v239 = v230
	goto L61
L61:
	;
	v241 = v119 + int32(4)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v243 = int32(30)
	v248 = v228 & int32(1073741824)
	if int32(base.Ui32(v242)>>(uint(v243)%32))&int32(1) != int32(base.Ui32(v248)>>(uint(v243)%32)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v405 = int32(1)
	v502 = v398
	v505 = v92 + v405
	v508 = v402
	v509 = v403
	v513 = v103 + v405
	goto L20
L63:
	;
	v350 = v81 + v92<<(uint(int32(3))%32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v351 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L64:
	;
	if v248 != 0 {
		v398 = v89
		v402 = v95
		v403 = v96
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v253 = v242 & int32(1073741823)
	if int32(0) <= v242 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v241-int32(4))))
	v262 = v253 - v258&int32(1073741823)
	goto L68
L67:
	;
	v262 = v253
	goto L68
L68:
	;
	if v262 != v239 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	v264 = int32(0)
	if v264 <= v228 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v227-int32(4))))
	v272 = v269 & int32(1073741823)
	goto L72
L71:
	;
	v272 = v264
	goto L72
L72:
	;
	v273 = v272 + v87
	if int32(0) <= v242 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v241-int32(4))))
	v281 = v278 & int32(1073741823)
	goto L75
L74:
	;
	v281 = v264
	goto L75
L75:
	;
	v282 = v281 + v79
	if base.Ui32(int32(4)) <= base.Ui32(v239) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v344 == int32(0) {
		v398 = v89
		v402 = v95
		v403 = v96
		goto L62
	} else {
		goto L94
	}
L77:
	;
	v344 = int32(0)
	goto L76
L78:
	;
	v318 = v313
	v319 = v314
	v320 = v315
	goto L88
L79:
	;
	if (v273|v282)&int32(3) != 0 {
		v313 = v273
		v314 = v282
		v315 = v239
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v306 = v273
	v307 = v282
	v308 = v239
	goto L81
L81:
	;
	if v308 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v290 = v273
	v291 = v282
	v292 = v239
	goto L83
L83:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v295 != v296 {
		v313 = v290
		v314 = v291
		v315 = v292
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v306 = v301
	v307 = v299
	v308 = v303
	goto L81
L85:
	;
	v298 = int32(4)
	v299 = v291 + v298
	v301 = v290 + v298
	v303 = v292 - v298
	if base.Ui32(int32(3)) < base.Ui32(v303) {
		v290 = v301
		v291 = v299
		v292 = v303
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v313 = v306
	v314 = v307
	v315 = v308
	goto L78
L88:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v323 == v324 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v344 = v323 - v324
	goto L76
L90:
	;
	v326 = int32(1)
	v331 = v320 - v326
	if v331 != 0 {
		v318 = v318 + v326
		v319 = v319 + v326
		v320 = v331
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	goto L63
L95:
	;
	v368 = v239 + v365
	if v368 != 0 {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v365 = v351 & int32(1073741823)
	v367 = v87
	goto L95
L97:
	;
	goto L98
L98:
	;
	v356 = int32(1073741823)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v350-int32(4))))
	v362 = v360 & v356
	v365 = v351&v356 - v362
	v367 = v362 + v87
	goto L95
L99:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v373 = v371 & int32(1073741823)
	if int32(0) <= v371 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v369 = F__emscripten_memcpy_bulkmem(m, v96, v367, v368)
	mBase = m.M
	v370 = v369
	goto L102
L101:
	;
	v370 = v96
	goto L102
L102:
	;
	goto L99
L103:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v350-int32(4))))
	v382 = v373 - v378&int32(1073741823)
	goto L105
L104:
	;
	v382 = v373
	goto L105
L105:
	;
	v384 = v382 + v370 + v239
	v385 = v384 - v50
	v386 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v385&v386 | v248
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = (v385 - v239) & v386
	v398 = v89 + int32(8)
	v402 = v95 + int32(1)
	v403 = v384
	goto L62
L106:
	;
	goto L22
L107:
	;
	v441 = v81 + v92<<(uint(int32(3))%32) + int32(4)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v444 = v442 & int32(1073741823)
	if int32(0) <= v442 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v435 = v420 & int32(1073741823)
	v436 = v87
	goto L107
L109:
	;
	goto L110
L110:
	;
	v425 = int32(1073741823)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v419-int32(4))))
	v431 = v429 & v425
	v435 = v420&v425 - v431
	v436 = v431 + v87
	goto L107
L111:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v441-int32(4))))
	v453 = v444 - v449&int32(1073741823)
	goto L113
L112:
	;
	v453 = v444
	goto L113
L113:
	;
	v454 = v453 + v435
	if v454 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v459 = v457 & int32(1073741823)
	if int32(0) <= v457 {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v455 = F__emscripten_memcpy_bulkmem(m, v96, v436, v454)
	mBase = m.M
	v456 = v455
	goto L117
L116:
	;
	v456 = v96
	goto L117
L117:
	;
	goto L114
L118:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v419-int32(4))))
	v468 = v459 - v464&int32(1073741823)
	goto L120
L119:
	;
	v468 = v459
	goto L120
L120:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v472 = int32(0)
	if v472 <= v469 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v475 = v459
	goto L123
L122:
	;
	v475 = v472
	goto L123
L123:
	;
	v476 = v469&int32(1073741823) - v475
	v478 = v476 + (v468 + v456)
	v479 = v478 - v50
	v481 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = (v479 - v476) & v481
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v484&int32(1073741824) | v479&v481
	v491 = int32(1)
	v502 = v89 + int32(8)
	v505 = v92 + v491
	v508 = v95 + v491
	v509 = v478
	v513 = v103
	goto L20
L124:
	;
	goto L19
L125:
	;
	v526 = v515
	v529 = int32(0)
	goto L5
L126:
	;
	goto L127
L127:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v519 | int32(-2147483648)
	v526 = v515
	v529 = v508
	goto L5
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v529 | int32(-2147483648)
	v555 = v47 + v529<<(uint(int32(3))%32)&int32(2147483640)
	if v555 == v50 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = (v529<<(uint(int32(3))%32)+v526)<<(uint(int32(2))%32) + int32(32)
	return v33
L131:
	;
	goto L130
L132:
	;
	goto L131
L133:
	;
	v559 = v555 + v526
	if base.Ui32(v50-v559) <= base.Ui32(int32(0)-v526<<(uint(int32(1))%32)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v566 = F___memcpy(m, v555, v50, v526)
	mBase = m.M
	goto L131
L135:
	;
	goto L136
L136:
	;
	v569 = (v555 ^ v50) & int32(3)
	if base.Ui32(v555) < base.Ui32(v50) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	if v671 == int32(0) {
		goto L132
	} else {
		goto L173
	}
L138:
	;
	if base.Ui32(v649) <= base.Ui32(int32(3)) {
		v670 = v648
		v671 = v649
		v672 = v650
		goto L137
	} else {
		goto L169
	}
L139:
	;
	if v569 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	if v569 != 0 {
		v631 = v526
		goto L152
	} else {
		goto L153
	}
L142:
	;
	v670 = v50
	v671 = v526
	v672 = v555
	goto L137
L143:
	;
	goto L144
L144:
	;
	if v555&int32(3) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v648 = v50
	v649 = v526
	v650 = v555
	goto L138
L146:
	;
	goto L147
L147:
	;
	v576 = v50
	v577 = v526
	v578 = v555
	goto L148
L148:
	;
	if v577 == int32(0) {
		goto L132
	} else {
		goto L150
	}
L149:
	;
	v648 = v585
	v649 = v587
	v650 = v589
	goto L138
L150:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	*(*uint8)(unsafe.Add(mBase, uint32(v578))) = uint8(v582)
	v584 = int32(1)
	v585 = v576 + v584
	v587 = v577 - v584
	v589 = v578 + v584
	if v589&int32(3) != 0 {
		v576 = v585
		v577 = v587
		v578 = v589
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	if v631 == int32(0) {
		goto L132
	} else {
		goto L165
	}
L153:
	;
	if v559&int32(3) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v596 = v526
	goto L157
L155:
	;
	v611 = v526
	goto L156
L156:
	;
	if base.Ui32(v611) <= base.Ui32(int32(3)) {
		v631 = v611
		goto L152
	} else {
		goto L161
	}
L157:
	;
	if v596 == int32(0) {
		goto L132
	} else {
		goto L159
	}
L158:
	;
	v611 = v602
	goto L156
L159:
	;
	v602 = v596 - int32(1)
	v603 = v555 + v602
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v602))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v605)
	if v603&int32(3) != 0 {
		v596 = v602
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v618 = v611
	goto L162
L162:
	;
	v622 = v618 - int32(4)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v50+v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v555+v622))) = v625
	if base.Ui32(int32(3)) < base.Ui32(v622) {
		v618 = v622
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v631 = v622
	goto L152
L164:
	;
	goto L163
L165:
	;
	v638 = v631
	goto L166
L166:
	;
	v642 = v638 - int32(1)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v642))))
	*(*uint8)(unsafe.Add(mBase, uint32(v555+v642))) = uint8(v645)
	if v642 != 0 {
		v638 = v642
		goto L166
	} else {
		goto L168
	}
L167:
	;
	goto L132
L168:
	;
	goto L167
L169:
	;
	v655 = v648
	v656 = v649
	v657 = v650
	goto L170
L170:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	*(*int32)(unsafe.Add(mBase, uint32(v657))) = v659
	v661 = int32(4)
	v662 = v655 + v661
	v664 = v657 + v661
	v666 = v656 - v661
	if base.Ui32(int32(3)) < base.Ui32(v666) {
		v655 = v662
		v656 = v666
		v657 = v664
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v670 = v662
	v671 = v666
	v672 = v664
	goto L137
L172:
	;
	goto L171
L173:
	;
	v677 = v670
	v678 = v671
	v679 = v672
	goto L174
L174:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677))))
	*(*uint8)(unsafe.Add(mBase, uint32(v679))) = uint8(v681)
	v683 = int32(1)
	v688 = v678 - v683
	if v688 != 0 {
		v677 = v677 + v683
		v678 = v688
		v679 = v679 + v683
		goto L174
	} else {
		goto L176
	}
L175:
	;
	goto L132
L176:
	;
	goto L175
}
func F_hstore_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
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
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v2)
	v20 = F_palloc(m, int32(320))
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
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v11
	v33 = v24
	v34 = v11
	goto L4
L3:
	;
	m.G0 = v9 + int32(96)
	return v430
L4:
	;
	switch v33 - int32(1) {
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
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v421 = F_hstoreUniquePairs(m, v417, v418, v9+int32(56))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L113
	}
L6:
	;
	goto L5
L7:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	v414 = v412 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v414
	v33 = v411
	v34 = v414
	goto L4
L8:
	;
	v411 = int32(2)
	goto L7
L9:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v395 = v392 + v55*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v68
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v398
	goto L8
L10:
	;
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v389)
	v430 = int32(0)
	goto L3
L11:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v340 == int32(44) {
		v411 = int32(0)
		goto L7
	} else {
		goto L103
	}
L12:
	;
	v218 = F_get_val(m, v9+int32(60), int32(1), v9+int32(56))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v158 == int32(62) {
		v411 = int32(1)
		goto L7
	} else {
		goto L50
	}
L14:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v94 == int32(61) {
		v411 = int32(3)
		goto L7
	} else {
		goto L33
	}
L15:
	;
	v44 = F_get_val(m, v9+int32(60), int32(0), v9+int32(56))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v44 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	if v48 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
	if v56 <= v55 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v51 != int32(447) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v54 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v56 << (uint(int32(1)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v64 = F_repalloc(m, v61, v56*int32(40))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	v69 = v67 - v68
	if base.Ui32(v69) < base.Ui32(int32(1073741824)) {
		goto L9
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v64
	goto L25
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v73 = F_errsave_start(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v73 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(22339), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, v72, int32(495860), int32(423), int32(320942))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	if v94 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v100 = F_errsave_start(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v120 = base.I32_extend8_s(v94)
	goto L42
L37:
	;
	if v100 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(330195), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errsave_finish(m, v99, int32(495860), int32(83), int32(338286))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L10
L42:
	;
	if base.B2i32(v120 == int32(32))|base.B2i32(base.Ui32((v120-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v131 = F_errsave_start(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v131 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v138 = F_pg_mblen_cstr(m, v34)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v34
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v34 - v142
	F_errmsg(m, int32(472452), v9)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errsave_finish(m, v130, int32(495860), int32(71), int32(210085))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L10
L50:
	;
	if v158 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v164 = F_errsave_start(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v185 = F_errsave_start(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	if v164 == int32(0) {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(330195), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, v163, int32(495860), int32(83), int32(338286))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	if v185 == int32(0) {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v192 = F_pg_mblen_cstr(m, v34)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v34
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v34 - v196
	F_errmsg(m, int32(472452), v9+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errsave_finish(m, v184, int32(495860), int32(71), int32(210085))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	if v218 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	if v222 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	v251 = v249 - v250
	if base.Ui32(int32(1073741824)) <= base.Ui32(v251) {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	v229 = F_errsave_start(m, v222)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v225 != int32(447) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
	if v228 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	if v229 == int32(0) {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(330195), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errsave_finish(m, v222, int32(495860), int32(83), int32(338286))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L10
L78:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v255 = F_errsave_start(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	v279 = v275 + v276*int32(20)
	v280 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v279)+16)) = uint16(v280)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v250
	if v251 != int32(4) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	if v255 == int32(0) {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(345807), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errsave_finish(m, v254, int32(495860), int32(443), int32(320963))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v276 + int32(1)
	v411 = int32(4)
	goto L7
L87:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)))
	if v286 != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+4)) = uint8(v287)
	v292 = v250
	v293 = int32(303260)
	goto L90
L89:
	;
	if v330 != 0 {
		goto L86
	} else {
		goto L102
	}
L90:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v296 == v297 {
		v319 = v296
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v330 = int32(0)
	goto L89
L92:
	;
	v321 = int32(1)
	if v319 != 0 {
		v292 = v292 + v321
		v293 = v293 + v321
		goto L90
	} else {
		goto L101
	}
L93:
	;
	if base.Ui32((v296-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v307 = v296 | int32(32)
	goto L96
L95:
	;
	v307 = v296
	goto L96
L96:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v316 = v297 | int32(32)
	goto L99
L98:
	;
	v316 = v297
	goto L99
L99:
	;
	if v307 == v316 {
		v319 = v307
		goto L92
	} else {
		goto L100
	}
L100:
	;
	v330 = v307 - v316
	goto L89
L101:
	;
	goto L91
L102:
	;
	v331 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v279)+16)) = uint8(v331)
	goto L86
L103:
	;
	if v340 == int32(0) {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v346 = base.I32_extend8_s(v340)
	goto L105
L105:
	;
	if base.B2i32(v346 == int32(32))|base.B2i32(base.Ui32((v346-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v411 = int32(4)
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
	v357 = F_errsave_start(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v357 == int32(0) {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v364 = F_pg_mblen_cstr(m, v34)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v34
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v34 - v368
	F_errmsg(m, int32(472452), v9+int32(32))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errsave_finish(m, v356, int32(495860), int32(71), int32(210085))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L10
L113:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v424 = F_hstorePairs(m, v417, v421, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v430 = v424
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5467), int32(0), v4, v5)
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = int32(4515120)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = F_hstoreUpgrade(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v20
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v25 = F_palloc(m, int32(base.Ui32(v22)>>(uint(int32(2))%32)))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
					if v29 != 0 {
						v30 = F__emscripten_memcpy_bulkmem(m, v25, v13, v29)
						mBase = m.M
						v31 = v30
					} else {
						v31 = v25
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v31
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v19
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
					v43 = v41 & int32(268435455)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					if base.Ui32(v44) < base.Ui32(v43) {
						v47 = v40 + int32(8)
						v48 = int32(3)
						v50 = v47 + v43<<(uint(v48)%32)
						v53 = v47 + v44<<(uint(v48)%32)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						if v54 < int32(0) {
							v68 = v50
							v69 = v54 & int32(1073741823)
						} else {
							v59 = int32(1073741823)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v53-int32(4))))
							v65 = v63 & v59
							v68 = v50 + v65
							v69 = v54&v59 - v65
						}
						v71 = F_cstring_to_text_with_len(m, v68, v69)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
							*(*int64)(unsafe.Add(mBase, uint32(v39))) = v73 + int64(1)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = int32(1)
							return v71
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(2)
							v86 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
		v43 = v41 & int32(268435455)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
		if base.Ui32(v44) < base.Ui32(v43) {
			v47 = v40 + int32(8)
			v48 = int32(3)
			v50 = v47 + v43<<(uint(v48)%32)
			v53 = v47 + v44<<(uint(v48)%32)
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
			if v54 < int32(0) {
				v68 = v50
				v69 = v54 & int32(1073741823)
			} else {
				v59 = int32(1073741823)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v53-int32(4))))
				v65 = v63 & v59
				v68 = v50 + v65
				v69 = v54&v59 - v65
			}
			v71 = F_cstring_to_text_with_len(m, v68, v69)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
				*(*int64)(unsafe.Add(mBase, uint32(v39))) = v73 + int64(1)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = int32(1)
				return v71
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(2)
				v86 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
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
	var v179 int32
	_ = v179
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
	return v179
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
	v26 = F_cstring_to_text_with_len(m, int32(4103), int32(2))
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
	v179 = v26
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
	v165 = m.ExcPending
	if v165 != 0 {
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
	v67 = v32 + v62
	goto L12
L16:
	;
	F_appendStringInfoString(m, v12, int32(745451))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v77 = v29 + v39<<(uint(int32(3))%32) + int32(4)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78&int32(1073741824) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v157 = v39 + int32(1)
	if v21 != v157 {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	F_appendStringInfoString(m, v12, int32(303260))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if int32(0) <= v78 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L18
L23:
	;
	v147 = F_IsValidJsonNumber(m, v146, v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L45
	}
L24:
	;
	if v78 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v122))))
	if v124 != int32(102) {
		goto L24
	} else {
		goto L40
	}
L26:
	;
	if v117 != int32(1) {
		v143 = v117
		v146 = v32
		goto L23
	} else {
		goto L39
	}
L27:
	;
	F_appendStringInfoString(m, v12, int32(344091))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L38
	}
L28:
	;
	v87 = v77 - int32(4)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v90 = v88 & int32(1073741823)
	if v78-v90 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v105 = v78 & int32(1073741823)
	if v105 != int32(1) {
		v117 = v105
		goto L26
	} else {
		goto L36
	}
L31:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v90))))
	if v95 == int32(116) {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v100 = v98 & int32(1073741823)
	if v78-v100 == int32(1) {
		v122 = v100
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
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v109 != int32(116) {
		v117 = int32(1)
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
	v122 = int32(0)
	goto L25
L40:
	;
	F_appendStringInfoString(m, v12, int32(361107))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	goto L18
L42:
	;
	v143 = v78 & int32(1073741823)
	v146 = v32
	goto L23
L43:
	;
	goto L44
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v77-int32(4))))
	v140 = v138 & int32(1073741823)
	v143 = v78 - v140
	v146 = v32 + v140
	goto L23
L45:
	;
	if v147 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_appendBinaryStringInfo(m, v12, v146, v143)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_escape_json_with_len(m, v12, v146, v143)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	F_appendStringInfoString(m, v12, int32(745694))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v157 != v21 {
		v39 = v157
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
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v168 = F_cstring_to_text_with_len(m, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v179 = v168
	goto L1
}
func F_hstore_to_jsonb_loose(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_hstoreUpgrade(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(0)
	F_initStringInfo(m, v10+int32(-20))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = F_pushJsonbValue(m, v10+int32(-4), int32(6), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = v19 & int32(268435455)
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v35 = v15 + int32(8)
	v38 = v35 + v33<<(uint(int32(3))%32)
	v40 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v233 = F_pushJsonbValue(m, v10+int32(-4), int32(7), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1)
	v53 = v35 + v40<<(uint(int32(3))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v56 = v54 & int32(1073741823)
	if v54 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v72 + v38
	v80 = F_pushJsonbValue(m, v10+int32(-4), int32(1), v10+int32(-40))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v56
	v72 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v62 = v53 - int32(4)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v56 - v63&v64
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v72 = v68 & v64
	goto L10
L14:
	;
	v86 = v35 + v40<<(uint(int32(3))%32) + int32(4)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87&int32(1073741824) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v215 = F_pushJsonbValue(m, v10+int32(-4), int32(2), v10+int32(-60))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L49
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if int32(0) <= v87 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v139 = v10 + int32(-20)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v141)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v141
	goto L33
L20:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v38))))
	if v129 != int32(102) {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
	goto L15
L22:
	;
	v95 = v86 - int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v98 = v96 & int32(1073741823)
	if v87-v98 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v87&int32(1073741823) != int32(1) {
		goto L19
	} else {
		goto L30
	}
L25:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v98))))
	if v103 == int32(116) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v108 = v106 & int32(1073741823)
	if v87-v108 != int32(1) {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v126 = v108
	goto L20
L30:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v117 != int32(116) {
		v126 = int32(0)
		goto L20
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v132 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v132)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
	goto L15
L33:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v149 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_appendBinaryStringInfo(m, v10+int32(-20), v165, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v163 = v149 & int32(1073741823)
	v165 = v38
	goto L34
L36:
	;
	goto L37
L37:
	;
	v154 = int32(1073741823)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v86-int32(4))))
	v160 = v158 & v154
	v163 = v149&v154 - v160
	v165 = v160 + v38
	goto L34
L38:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v170 = F_IsValidJsonNumber(m, v168, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v170 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(2)
	v175 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v179 = F_DirectFunctionCall3Coll(m, int32(408), v175, v176, v175, int32(-1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v188 = v186 & int32(1073741823)
	if v186 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v181 = F_pg_detoast_datum(m, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v181
	goto L15
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v204 + v38
	goto L15
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v188
	v204 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v194 = v86 - int32(4)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v188 - v195&v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v204 = v200 & v196
	goto L45
L49:
	;
	v218 = v40 + int32(1)
	if v218 != v33 {
		v40 = v218
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L9
L51:
	;
	v235 = F_JsonbValueToJsonb(m, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	m.G0 = v12 - int32(-64)
	return v235
}
