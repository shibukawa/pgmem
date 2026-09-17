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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
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
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		v152 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v14 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v179 = (v152 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	v28 = int32(8)
	v33 = v22
	v37 = v14 + v28
	v38 = v19 + v28
	v43 = v21
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v37 + v44
	v47 = v38 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v152 = v132
	goto L5
L9:
	;
	v132 = v33 - int32(1)
	if v33 < int32(2) {
		v152 = v132
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v51 = v48
	goto L12
L11:
	;
	v51 = v49
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v113 = int32(0)
	goto L13
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L25
L16:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L18
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L22:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L15
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = v92 - v93
	goto L13
L27:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v48 == v49 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v113 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(10)
	v179 = (v33*v117 + v117) * (v48 - v49)
	goto L4
L35:
	;
	v129 = int32(-10)
	goto L37
L36:
	;
	v129 = int32(10)
	goto L37
L37:
	;
	v179 = (v33 + int32(1)) * v129
	goto L4
L38:
	;
	v135 = int32(9)
	v137 = int32(_a_F_ltree_ge_0)
	v145 = int32(1)
	if v145 < v43 {
		v33 = v132
		v37 = v37 + (v48+v135)&v137
		v38 = v38 + (v49+v135)&v137
		v43 = v43 - v145
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v184 != v19 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v19)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	return int32(base.Ui32(v179^int32(-1)) >> (uint(int32(31)) % 32))
L47:
	;
	goto L46
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
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
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0])) = v17
	v22 = v17
	goto L3
L6:
	;
	return v169
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
	v169 = v66
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
		v169 = v66
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
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0]))
	v81 = F_pg_strfold(m, v73, v72, l0, l1, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v84 = v81 + int32(1)
	if base.Ui32(v72) < base.Ui32(v84) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v86 = F_repalloc(m, v73, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v92 = v73
	v93 = v81
	goto L33
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0]))
	v96 = F_pg_strfold(m, v77, v76, l2, l3, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0]))
	v90 = F_pg_strfold(m, v86, v84, l0, l1, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v92 = v86
	v93 = v90
	goto L33
L36:
	;
	v99 = v96 + int32(1)
	if base.Ui32(v76) < base.Ui32(v99) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = F_repalloc(m, v77, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v107 = v96
	v108 = v77
	goto L39
L39:
	;
	if base.Ui32(v93) <= base.Ui32(v107) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_prefix_eq_ci[0]))
	v105 = F_pg_strfold(m, v101, v99, l2, l3, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v107 = v105
	v108 = v101
	goto L39
L42:
	;
	if v93 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v158 = int32(0)
	goto L44
L44:
	;
	F_pfree(m, v92)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L58
	}
L45:
	;
	v158 = base.B2i32(v155 == int32(0))
	goto L44
L46:
	;
	v155 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v116 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v117 = v92
	v118 = v108
	v119 = v93
	v120 = v116
	goto L53
L50:
	;
	v143 = v108
	v147 = int32(0)
	goto L51
L51:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v155 = v147 - v148
	goto L45
L52:
	;
	v143 = v138
	v147 = v140
	goto L51
L53:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if base.B2i32(v120 != v122)|base.B2i32(v122 == int32(0)) != 0 {
		v138 = v118
		v140 = v120
		goto L52
	} else {
		goto L55
	}
L54:
	;
	v138 = v132
	v140 = int32(0)
	goto L52
L55:
	;
	v128 = v119 - int32(1)
	if v128 == int32(0) {
		v138 = v118
		v140 = v120
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v131 = int32(1)
	v132 = v118 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v133 != 0 {
		v117 = v117 + v131
		v118 = v132
		v119 = v128
		v120 = v133
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	F_pfree(m, v108)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v169 = v158
	goto L6
}
func F_ltree_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
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
	var v133 int32
	_ = v133
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
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
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
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
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
	var v392 int32
	_ = v392
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v465 int32
	_ = v465
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v2 {
		v32 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v32&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v19 == int32(0) {
		v32 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(7) {
		v32 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v25 != int32(17) {
		v32 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v32 = v28 ^ int32(1)
	goto L2
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_get_fn_opclass_options(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v41 = int32(8)
	goto L9
L9:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v45 = int32(1)
	v46 = v44 & v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v46 != v47&v45 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v41 = v40
	goto L9
L12:
	;
	return v11
L13:
	;
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v452)
	goto L12
L15:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
	if v52 != v53 {
		v452 = int32(0)
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
	v158 = v44 & v157
	if v158 != v47&v157 {
		goto L12
	} else {
		goto L37
	}
L18:
	;
	v55 = int32(8)
	v59 = int32(0)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v55)+4)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v55)+4)))
	if base.B2i32(v66 == v59)|base.B2i32(v69 == v59) != 0 {
		v133 = v66
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v452 = base.B2i32(v154 == int32(0))
	goto L14
L20:
	;
	v154 = (v133 + int32(1)) * (v66 - v69) * int32(10)
	goto L19
L21:
	;
	v465 = int32(16)
	v77 = v13 + v465
	v78 = v12 + v465
	v81 = v66
	v84 = v69
	goto L22
L22:
	;
	v86 = int32(2)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	if base.Ui32(v90) < base.Ui32(v91) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v133 = v113
	goto L20
L24:
	;
	v113 = v81 - int32(1)
	if v81 < int32(2) {
		v133 = v113
		goto L20
	} else {
		goto L35
	}
L25:
	;
	v93 = v90
	goto L27
L26:
	;
	v93 = v91
	goto L27
L27:
	;
	v94 = F_memcmp(m, v77+v86, v78+v86, v93)
	mBase = m.M
	if v94 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v90 == v91 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v94 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v98 = int32(10)
	v154 = (v81*v98 + v98) * (v90 - v91)
	goto L19
L32:
	;
	v110 = int32(-10)
	goto L34
L33:
	;
	v110 = int32(10)
	goto L34
L34:
	;
	v154 = (v81 + int32(1)) * v110
	goto L19
L35:
	;
	v116 = int32(9)
	v118 = int32(_a_F_ltree_same_0)
	v126 = int32(1)
	if v126 < v84 {
		v77 = v77 + (v90+v116)&v118
		v78 = v78 + (v91+v116)&v118
		v81 = v113
		v84 = v84 - v126
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L23
L37:
	;
	v163 = v13 + int32(8)
	if v158 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v165 = int32(0)
	goto L40
L39:
	;
	v165 = v41
	goto L40
L40:
	;
	v166 = v163 + v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v169 = v12 + int32(8)
	v170 = v169 + v165
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	if v167 != v171 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v173 = int32(0)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	if base.B2i32(v180 == v173)|base.B2i32(v183 == v173) != 0 {
		v247 = v180
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v268 != 0 {
		goto L12
	} else {
		goto L60
	}
L43:
	;
	v268 = (v247 + int32(1)) * (v180 - v183) * int32(10)
	goto L42
L44:
	;
	v187 = int32(8)
	v191 = v166 + v187
	v192 = v170 + v187
	v195 = v180
	v198 = v183
	goto L45
L45:
	;
	v200 = int32(2)
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191))))
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	if base.Ui32(v204) < base.Ui32(v205) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v247 = v227
	goto L43
L47:
	;
	v227 = v195 - int32(1)
	if v195 < int32(2) {
		v247 = v227
		goto L43
	} else {
		goto L58
	}
L48:
	;
	v207 = v204
	goto L50
L49:
	;
	v207 = v205
	goto L50
L50:
	;
	v208 = F_memcmp(m, v191+v200, v192+v200, v207)
	mBase = m.M
	if v208 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v204 == v205 {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v208 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v212 = int32(10)
	v268 = (v195*v212 + v212) * (v204 - v205)
	goto L42
L55:
	;
	v224 = int32(-10)
	goto L57
L56:
	;
	v224 = int32(10)
	goto L57
L57:
	;
	v268 = (v195 + int32(1)) * v224
	goto L42
L58:
	;
	v230 = int32(9)
	v232 = int32(_a_F_ltree_same_0)
	v240 = int32(1)
	if v240 < v198 {
		v191 = v191 + (v204+v230)&v232
		v192 = v192 + (v205+v230)&v232
		v195 = v227
		v198 = v198 - v240
		goto L45
	} else {
		goto L59
	}
L59:
	;
	goto L46
L60:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v270&int32(2) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v273 = int32(0)
	goto L63
L62:
	;
	v273 = v41
	goto L63
L63:
	;
	v274 = v163 + v273
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v276&int32(2) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v279 = int32(0)
	goto L66
L65:
	;
	v279 = v41
	goto L66
L66:
	;
	v280 = v169 + v279
	v282 = v270 & int32(4)
	if v282 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v287 = v274
	goto L69
L68:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v287 = v274 + int32(base.Ui32(v283)>>(uint(int32(2))%32))
	goto L69
L69:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+4)))
	v290 = v276 & int32(4)
	if v290 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v295 = v280
	goto L72
L71:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v295 = v280 + int32(base.Ui32(v291)>>(uint(int32(2))%32))
	goto L72
L72:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+4)))
	if v288 != v296 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	if v270&int32(2) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v301 = int32(0)
	goto L76
L75:
	;
	v301 = v41
	goto L76
L76:
	;
	v302 = v163 + v301
	if v276&int32(2) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v306 = int32(0)
	goto L79
L78:
	;
	v306 = v41
	goto L79
L79:
	;
	v307 = v169 + v306
	if v282 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v312 = v302
	goto L82
L81:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v312 = v302 + int32(base.Ui32(v308)>>(uint(int32(2))%32))
	goto L82
L82:
	;
	if v290 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v317 = v307
	goto L85
L84:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v317 = v307 + int32(base.Ui32(v313)>>(uint(int32(2))%32))
	goto L85
L85:
	;
	v318 = int32(0)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312)+4)))
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v317)+4)))
	if base.B2i32(v325 == v318)|base.B2i32(v328 == v318) != 0 {
		v392 = v325
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v413 != 0 {
		goto L12
	} else {
		goto L104
	}
L87:
	;
	v413 = (v392 + int32(1)) * (v325 - v328) * int32(10)
	goto L86
L88:
	;
	v332 = int32(8)
	v336 = v312 + v332
	v337 = v317 + v332
	v340 = v325
	v343 = v328
	goto L89
L89:
	;
	v345 = int32(2)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336))))
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337))))
	if base.Ui32(v349) < base.Ui32(v350) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v392 = v372
	goto L87
L91:
	;
	v372 = v340 - int32(1)
	if v340 < int32(2) {
		v392 = v372
		goto L87
	} else {
		goto L102
	}
L92:
	;
	v352 = v349
	goto L94
L93:
	;
	v352 = v350
	goto L94
L94:
	;
	v353 = F_memcmp(m, v336+v345, v337+v345, v352)
	mBase = m.M
	if v353 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v349 == v350 {
		goto L91
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v353 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v357 = int32(10)
	v413 = (v340*v357 + v357) * (v349 - v350)
	goto L86
L99:
	;
	v369 = int32(-10)
	goto L101
L100:
	;
	v369 = int32(10)
	goto L101
L101:
	;
	v413 = (v340 + int32(1)) * v369
	goto L86
L102:
	;
	v375 = int32(9)
	v377 = int32(_a_F_ltree_same_0)
	v385 = int32(1)
	if v385 < v343 {
		v336 = v336 + (v349+v375)&v377
		v337 = v337 + (v350+v375)&v377
		v340 = v372
		v343 = v343 - v385
		goto L89
	} else {
		goto L103
	}
L103:
	;
	goto L90
L104:
	;
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v414)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	if v416&int32(2)|base.B2i32(v41 <= int32(0)) != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v423 = int32(0)
	goto L106
L106:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v163))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v169))))
	if v434 == v436 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v452 = int32(0)
	goto L14
L108:
	;
	v439 = v423 + int32(1)
	if v41 != v439 {
		v423 = v439
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
				v70 = v20
			} else {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				if v25 != 0 {
					base.MemoryCopy(m, v20, v13+int32(10), v25)
				} else {
				}
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
				v30 = v20 + v29
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
				if base.Ui32(v31) < base.Ui32(int32(2)) {
					v70 = v30
				} else {
					v42 = v13 + int32(8) + (v29+int32(9))&int32(_a_F_ltree_send_0)
					v44 = v30
					v48 = int32(1)
					for {
						v49 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v49)
						v52 = v44 + int32(1)
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
						if v53 != 0 {
							base.MemoryCopy(m, v52, v42+int32(2), v53)
						} else {
						}
						v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
						v58 = v52 + v57
						v65 = v48 + int32(1)
						v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
						if base.Ui32(v65) < base.Ui32(v66) {
							v42 = v42 + (v57+int32(9))&int32(_a_F_ltree_send_0)
							v44 = v58
							v48 = v65
							continue
						} else {
							break
						}
						break
					}
					v70 = v58
				}
			}
			v75 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v75)
			F_pq_begintypsend(m, v10)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v10, int32(1))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v85 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v82+v83))) = uint8(v85)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v82 + v85
					v90 = F_strlen(m, v20)
					mBase = m.M
					F_pq_sendtext(m, v10, v20, v90)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 << (uint(int32(2)) % 32)
							m.G0 = v10 + int32(16)
							return v96
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
				v18 = F_DirectFunctionCall1Coll(m, int32(_a_F_ltree_textadd_0), int32(0), v16)
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
