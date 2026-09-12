package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ltree_ge(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
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
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if v22 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v179 != v14 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v178 = (v149 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	if v21 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(8)
	v33 = v14 + v27
	v35 = v22
	v38 = v19 + v27
	v42 = v21
	goto L8
L8:
	;
	v43 = int32(2)
	v44 = v33 + v43
	v46 = v38 + v43
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v47) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v149 = v131
	goto L5
L10:
	;
	v131 = v35 - int32(1)
	if v35 < int32(2) {
		v149 = v131
		goto L5
	} else {
		goto L39
	}
L11:
	;
	v50 = v47
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v112 = int32(0)
	goto L14
L16:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L26
L17:
	;
	if (v44|v46)&int32(3) != 0 {
		v81 = v44
		v82 = v46
		v83 = v50
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v74 = v44
	v75 = v46
	v76 = v50
	goto L19
L19:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v58 = v44
	v59 = v46
	v60 = v50
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L19
L23:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L16
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = v91 - v92
	goto L14
L28:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	if v47 == v48 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v112 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v116 = int32(10)
	v178 = (v35*v116 + v116) * (v47 - v48)
	goto L4
L36:
	;
	v128 = int32(-10)
	goto L38
L37:
	;
	v128 = int32(10)
	goto L38
L38:
	;
	v178 = (v35 + int32(1)) * v128
	goto L4
L39:
	;
	v134 = int32(9)
	v136 = int32(131064)
	v144 = int32(1)
	if v144 < v42 {
		v33 = v33 + (v47+v134)&v136
		v35 = v131
		v38 = v38 + (v48+v134)&v136
		v42 = v42 - v144
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L9
L41:
	;
	F_pfree(m, v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 != v19 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return int32(base.Ui32(v178^int32(-1)) >> (uint(int32(31)) % 32))
L48:
	;
	goto L47
}
func F_ltree_prefix_eq_ci(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = F_pg_newlocale_from_collation(m, int32(100))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v22 = v12
	goto L3
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
	if v23 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1457])) = v17
	v22 = v17
	goto L3
L6:
	;
	return v173
L7:
	;
	if base.Ui32(l3) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v72 = l1 + int32(1)
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L28
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	if l1 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(1)
L14:
	;
	goto L15
L15:
	;
	v37 = int32(0)
	goto L16
L16:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v37))))
	if base.Ui32((v45-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v173 = v66
	goto L6
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v37))))
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v54 = v45 | int32(32)
	goto L21
L20:
	;
	v54 = v45
	goto L21
L21:
	;
	goto L18
L22:
	;
	v66 = base.B2i32(v54 == v65)
	if v54 != v65 {
		v173 = v66
		goto L6
	} else {
		goto L26
	}
L23:
	;
	v65 = v56 | int32(32)
	goto L25
L24:
	;
	v65 = v56
	goto L25
L25:
	;
	goto L22
L26:
	;
	v69 = v37 + int32(1)
	if v69 != l1 {
		v37 = v69
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	v76 = l3 + int32(1)
	v77 = F_palloc(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v82 = F_pg_strfold(m, v73, v72, l0, l1, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v85 = v82 + int32(1)
	if base.Ui32(v72) < base.Ui32(v85) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v88 = F_repalloc(m, v73, v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v95 = v73
	v97 = v82
	goto L33
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v100 = F_pg_strfold(m, v77, v76, l2, l3, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v93 = F_pg_strfold(m, v88, v85, l0, l1, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v95 = v88
	v97 = v93
	goto L33
L36:
	;
	v103 = v100 + int32(1)
	if base.Ui32(v76) < base.Ui32(v103) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v106 = F_repalloc(m, v77, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v113 = v100
	v114 = v77
	goto L39
L39:
	;
	if base.Ui32(v97) <= base.Ui32(v113) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v111 = F_pg_strfold(m, v106, v103, l2, l3, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v113 = v111
	v114 = v106
	goto L39
L42:
	;
	if v97 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v163 = int32(0)
	goto L44
L44:
	;
	F_pfree(m, v95)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L59
	}
L45:
	;
	v163 = base.B2i32(v160 == int32(0))
	goto L44
L46:
	;
	v160 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v122 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v123 = v95
	v124 = v114
	v125 = v97
	v126 = v122
	goto L53
L50:
	;
	v148 = v114
	v152 = int32(0)
	goto L51
L51:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v160 = v152 - v153
	goto L45
L52:
	;
	v148 = v143
	v152 = v145
	goto L51
L53:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v126 != v128 {
		v143 = v124
		v145 = v126
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v143 = v137
	v145 = int32(0)
	goto L52
L55:
	;
	if v128 == int32(0) {
		v143 = v124
		v145 = v126
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v133 = v125 - int32(1)
	if v133 == int32(0) {
		v143 = v124
		v145 = v126
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v136 = int32(1)
	v137 = v124 + v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	if v138 != 0 {
		v123 = v123 + v136
		v124 = v137
		v125 = v133
		v126 = v138
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	F_pfree(m, v114)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v173 = v163
	goto L6
}
func F_ltree_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v453 int32
	_ = v453
	var v467 int32
	_ = v467
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == v2 {
		v33 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v33&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v20 == int32(0) {
		v33 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v23 != int32(7) {
		v33 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v26 != int32(17) {
		v33 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	v33 = v29 ^ int32(1)
	goto L2
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = F_get_fn_opclass_options(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v42 = int32(8)
	goto L9
L9:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v43)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v46 = int32(1)
	v47 = v45 & v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v47 != v48&v46 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = v41
	goto L9
L12:
	;
	return v12
L13:
	;
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v453)
	goto L12
L15:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)))
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
	if v53 != v54 {
		v453 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v157 = int32(2)
	v158 = v45 & v157
	if v158 != v48&v157 {
		goto L12
	} else {
		goto L38
	}
L18:
	;
	v56 = int32(8)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v56)+4)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v56)+4)))
	if v68 == int32(0) {
		v132 = v68
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v453 = base.B2i32(v154 == int32(0))
	goto L14
L20:
	;
	v154 = (v132 + int32(1)) * (v68 - v67) * int32(10)
	goto L19
L21:
	;
	if v67 == int32(0) {
		v132 = v68
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v467 = int32(16)
	v77 = v14 + v467
	v78 = v13 + v467
	v81 = v68
	v85 = v67
	goto L23
L23:
	;
	v86 = int32(2)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	if base.Ui32(v90) < base.Ui32(v91) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v132 = v113
	goto L20
L25:
	;
	v113 = v81 - int32(1)
	if v81 < int32(2) {
		v132 = v113
		goto L20
	} else {
		goto L36
	}
L26:
	;
	v93 = v90
	goto L28
L27:
	;
	v93 = v91
	goto L28
L28:
	;
	v94 = F_memcmp(m, v77+v86, v78+v86, v93)
	mBase = m.M
	if v94 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v90 == v91 {
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v94 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v98 = int32(10)
	v154 = (v81*v98 + v98) * (v90 - v91)
	goto L19
L33:
	;
	v110 = int32(-10)
	goto L35
L34:
	;
	v110 = int32(10)
	goto L35
L35:
	;
	v154 = (v81 + int32(1)) * v110
	goto L19
L36:
	;
	v116 = int32(9)
	v118 = int32(131064)
	v126 = int32(1)
	if v126 < v85 {
		v77 = v77 + (v90+v116)&v118
		v78 = v78 + (v91+v116)&v118
		v81 = v113
		v85 = v85 - v126
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L24
L38:
	;
	v163 = v14 + int32(8)
	if v158 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v165 = int32(0)
	goto L41
L40:
	;
	v165 = v42
	goto L41
L41:
	;
	v166 = v163 + v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v169 = v13 + int32(8)
	v170 = v169 + v165
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	if v167 != v171 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if v181 == int32(0) {
		v245 = v181
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v267 != 0 {
		goto L12
	} else {
		goto L62
	}
L44:
	;
	v267 = (v245 + int32(1)) * (v181 - v180) * int32(10)
	goto L43
L45:
	;
	if v180 == int32(0) {
		v245 = v181
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v186 = int32(8)
	v190 = v166 + v186
	v191 = v170 + v186
	v194 = v181
	v198 = v180
	goto L47
L47:
	;
	v199 = int32(2)
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191))))
	if base.Ui32(v203) < base.Ui32(v204) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v245 = v226
	goto L44
L49:
	;
	v226 = v194 - int32(1)
	if v194 < int32(2) {
		v245 = v226
		goto L44
	} else {
		goto L60
	}
L50:
	;
	v206 = v203
	goto L52
L51:
	;
	v206 = v204
	goto L52
L52:
	;
	v207 = F_memcmp(m, v190+v199, v191+v199, v206)
	mBase = m.M
	if v207 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v203 == v204 {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v207 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v211 = int32(10)
	v267 = (v194*v211 + v211) * (v203 - v204)
	goto L43
L57:
	;
	v223 = int32(-10)
	goto L59
L58:
	;
	v223 = int32(10)
	goto L59
L59:
	;
	v267 = (v194 + int32(1)) * v223
	goto L43
L60:
	;
	v229 = int32(9)
	v231 = int32(131064)
	v239 = int32(1)
	if v239 < v198 {
		v190 = v190 + (v203+v229)&v231
		v191 = v191 + (v204+v229)&v231
		v194 = v226
		v198 = v198 - v239
		goto L47
	} else {
		goto L61
	}
L61:
	;
	goto L48
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v269&int32(2) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v272 = int32(0)
	goto L65
L64:
	;
	v272 = v42
	goto L65
L65:
	;
	v273 = v163 + v272
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v275&int32(2) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v278 = int32(0)
	goto L68
L67:
	;
	v278 = v42
	goto L68
L68:
	;
	v279 = v169 + v278
	v281 = v269 & int32(4)
	if v281 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = v273
	goto L71
L70:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v286 = v273 + int32(base.Ui32(v282)>>(uint(int32(2))%32))
	goto L71
L71:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286)+4)))
	v289 = v275 & int32(4)
	if v289 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v294 = v279
	goto L74
L73:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v294 = v279 + int32(base.Ui32(v290)>>(uint(int32(2))%32))
	goto L74
L74:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+4)))
	if v295 != v287&int32(65535) {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	if v269&int32(2) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v302 = int32(0)
	goto L78
L77:
	;
	v302 = v42
	goto L78
L78:
	;
	v303 = v163 + v302
	if v275&int32(2) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v307 = int32(0)
	goto L81
L80:
	;
	v307 = v42
	goto L81
L81:
	;
	v308 = v169 + v307
	if v281 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v313 = v303
	goto L84
L83:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v313 = v303 + int32(base.Ui32(v309)>>(uint(int32(2))%32))
	goto L84
L84:
	;
	if v289 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v318 = v308
	goto L87
L86:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v318 = v308 + int32(base.Ui32(v314)>>(uint(int32(2))%32))
	goto L87
L87:
	;
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+4)))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+4)))
	if v327 == int32(0) {
		v391 = v327
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v413 != 0 {
		goto L12
	} else {
		goto L107
	}
L89:
	;
	v413 = (v391 + int32(1)) * (v327 - v326) * int32(10)
	goto L88
L90:
	;
	if v326 == int32(0) {
		v391 = v327
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v332 = int32(8)
	v336 = v313 + v332
	v337 = v318 + v332
	v340 = v327
	v344 = v326
	goto L92
L92:
	;
	v345 = int32(2)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336))))
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337))))
	if base.Ui32(v349) < base.Ui32(v350) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v391 = v372
	goto L89
L94:
	;
	v372 = v340 - int32(1)
	if v340 < int32(2) {
		v391 = v372
		goto L89
	} else {
		goto L105
	}
L95:
	;
	v352 = v349
	goto L97
L96:
	;
	v352 = v350
	goto L97
L97:
	;
	v353 = F_memcmp(m, v336+v345, v337+v345, v352)
	mBase = m.M
	if v353 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v349 == v350 {
		goto L94
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v353 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v357 = int32(10)
	v413 = (v340*v357 + v357) * (v349 - v350)
	goto L88
L102:
	;
	v369 = int32(-10)
	goto L104
L103:
	;
	v369 = int32(10)
	goto L104
L104:
	;
	v413 = (v340 + int32(1)) * v369
	goto L88
L105:
	;
	v375 = int32(9)
	v377 = int32(131064)
	v385 = int32(1)
	if v385 < v344 {
		v336 = v336 + (v349+v375)&v377
		v337 = v337 + (v350+v375)&v377
		v340 = v372
		v344 = v344 - v385
		goto L92
	} else {
		goto L106
	}
L106:
	;
	goto L93
L107:
	;
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v414)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	if v416&int32(2) != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	if v42 <= int32(0) {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	v422 = int32(0)
	goto L110
L110:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+v163))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+v169))))
	if v434 == v436 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v453 = int32(0)
	goto L14
L112:
	;
	v439 = v422 + int32(1)
	if v42 != v439 {
		v422 = v439
		goto L110
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	goto L12
}
func F_ltree_send(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v20 = F_palloc(m, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
			if v22 == int32(0) {
				v73 = v20
			} else {
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				if v27 != 0 {
					v28 = F__emscripten_memcpy_bulkmem(m, v20, v13+int32(10), v27)
					mBase = m.M
					v29 = v28
				} else {
					v29 = v20
				}
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				v31 = v29 + v30
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
				if base.Ui32(v32) < base.Ui32(int32(2)) {
					v73 = v31
				} else {
					v43 = v13 + int32(8) + (v30+int32(9))&int32(131064)
					v46 = v31
					v48 = int32(1)
					for {
						v50 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
						v53 = v46 + int32(1)
						v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
						if v56 != 0 {
							v57 = F__emscripten_memcpy_bulkmem(m, v53, v43+int32(2), v56)
							mBase = m.M
							v58 = v57
						} else {
							v58 = v53
						}
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
						v60 = v58 + v59
						v67 = v48 + int32(1)
						v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
						if base.Ui32(v67) < base.Ui32(v68) {
							v43 = v43 + (v59+int32(9))&int32(131064)
							v46 = v60
							v48 = v67
							continue
						} else {
							break
						}
						break
					}
					v73 = v60
				}
			}
			v77 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v77)
			F_pq_begintypsend(m, v10)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v10, int32(1))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v87 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v84+v85))) = uint8(v87)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v84 + v87
					v92 = F_strlen(m, v20)
					mBase = m.M
					F_pq_sendtext(m, v10, v20, v92)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
							m.G0 = v10 + int32(16)
							return v98
						}
					}
				}
			}
		}
	}
}
func F_ltree_textadd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_text_to_cstring(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_DirectFunctionCall1Coll(m, int32(5636), int32(0), v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v16)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_ltree_concat(m, v18, v7)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v18)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v26 != v7 {
									F_pfree(m, v7)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v30 != v14 {
											F_pfree(m, v14)
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											return v22
										}
									}
								} else {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v30 != v14 {
										F_pfree(m, v14)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										return v22
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
