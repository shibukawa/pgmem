package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ltree_ge(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v159 != v14 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v157 = v22 - v21
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v40 = v21
	v41 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113 != 0 {
		v157 = v113
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	if v48 != v49 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v48 - v49
	goto L4
L32:
	;
	goto L33
L33:
	;
	if v41 < int32(2) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(1)
	v120 = int32(9)
	v122 = int32(_a_F_ltree_ge_0)
	if v118 < v40 {
		v35 = v35 + (v48+v120)&v122
		v36 = v36 + (v49+v120)&v122
		v40 = v40 - v118
		v41 = v41 - v118
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v163 != v19 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v19)
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
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v157))
L43:
	;
	goto L42
}
func F_ltree_label_match(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	v7 = int32(0)
	if base.B2i32(l4&base.B2i32(base.Ui32(l1) < base.Ui32(l3)) == v7)&base.B2i32(l1 != l3) == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v179
L2:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0]))
	if v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	if l5 == int32(0) {
		v179 = v7
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v63 = base.B2i32(v61 == int32(0))
	if v61 == int32(0) {
		v179 = v63
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v61 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v23 = l0
	v24 = l2
	v25 = l1
	v26 = v22
	goto L14
L11:
	;
	v49 = l2
	v53 = int32(0)
	goto L12
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v61 = v53 - v54
	goto L6
L13:
	;
	v49 = v44
	v53 = v46
	goto L12
L14:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v26 != v28)|base.B2i32(v28 == int32(0)) != 0 {
		v44 = v24
		v46 = v26
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v44 = v38
	v46 = int32(0)
	goto L13
L16:
	;
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		v44 = v24
		v46 = v26
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v37 = int32(1)
	v38 = v24 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + v37
		v24 = v38
		v25 = v34
		v26 = v39
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	if l5 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v179 = v63
	goto L1
L21:
	;
	goto L2
L22:
	;
	v75 = F_pg_database_locale(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v81 = l1 + int32(1)
	v82 = F_palloc(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L27
	}
L25:
	;
	return int32(0)
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0])) = v75
	goto L24
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0]))
	v86 = F_pg_strfold(m, v82, v81, l0, l1, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(l1) < base.Ui32(v86) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v90 = v86 + int32(1)
	v91 = F_repalloc(m, v82, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L25
	} else {
		goto L32
	}
L30:
	;
	v97 = v82
	v98 = v86
	goto L31
L31:
	;
	v100 = l3 + int32(1)
	v101 = F_palloc(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L25
	} else {
		goto L35
	}
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0]))
	v95 = F_pg_strfold(m, v91, v90, l0, l1, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v97 = v91
	v98 = v95
	goto L31
L34:
	;
	F_pfree(m, v97)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L58
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0]))
	v105 = F_pg_strfold(m, v101, v100, l2, l3, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(l3) < base.Ui32(v105) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v109 = v105 + int32(1)
	v110 = F_repalloc(m, v101, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L25
	} else {
		goto L40
	}
L38:
	;
	v116 = v101
	v117 = v105
	goto L39
L39:
	;
	if base.B2i32(l4&base.B2i32(base.Ui32(v98) < base.Ui32(v117)) == int32(0))&base.B2i32(v117 != v98) != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_label_match[0]))
	v114 = F_pg_strfold(m, v110, v109, l2, l3, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v116 = v110
	v117 = v114
	goto L39
L42:
	;
	v171 = int32(0)
	goto L34
L43:
	;
	if v98 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v168 != 0 {
		goto L42
	} else {
		goto L57
	}
L45:
	;
	v168 = int32(0)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v129 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v130 = v97
	v131 = v116
	v132 = v98
	v133 = v129
	goto L52
L49:
	;
	v156 = v116
	v160 = int32(0)
	goto L50
L50:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v168 = v160 - v161
	goto L44
L51:
	;
	v156 = v151
	v160 = v153
	goto L50
L52:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if base.B2i32(v133 != v135)|base.B2i32(v135 == int32(0)) != 0 {
		v151 = v131
		v153 = v133
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v151 = v145
	v153 = int32(0)
	goto L51
L54:
	;
	v141 = v132 - int32(1)
	if v141 == int32(0) {
		v151 = v131
		v153 = v133
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v144 = int32(1)
	v145 = v131 + v144
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v146 != 0 {
		v130 = v130 + v144
		v131 = v145
		v132 = v141
		v133 = v146
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v171 = int32(1)
	goto L34
L58:
	;
	F_pfree(m, v116)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	v179 = v171
	goto L1
}
func F_ltree_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = base.I32_wrap_i64(v14)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 == v2 {
		v34 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v34&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v21 == int32(0) {
		v34 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v24 != int32(7) {
		v34 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v27 != int32(17) {
		v34 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)))
	v34 = v30 ^ int32(1)
	goto L2
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_get_fn_opclass_options(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v43 = int32(8)
	goto L9
L9:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v49 = int32(1)
	v50 = v48 & v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v50 != v51&v49 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int64(0)
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v43 = v42
	goto L9
L12:
	;
	return v14 & int64(4294967295)
L13:
	;
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v425)
	goto L12
L15:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)))
	if v56 != v57 {
		v425 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v150 = int32(2)
	v151 = v48 & v150
	if v151 != v51&v150 {
		goto L12
	} else {
		goto L34
	}
L18:
	;
	v59 = int32(8)
	v63 = int32(0)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v59)+4)))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v59)+4)))
	if base.B2i32(v70 == v63)|base.B2i32(v73 == v63) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v425 = base.B2i32(v147 == int32(0))
	goto L14
L20:
	;
	v147 = v137
	goto L19
L21:
	;
	v137 = v70 - v73
	goto L20
L22:
	;
	v439 = int32(16)
	v81 = v13 + v439
	v82 = v12 + v439
	v85 = v73
	v86 = v70
	goto L23
L23:
	;
	v90 = int32(2)
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81))))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82))))
	if base.Ui32(v94) < base.Ui32(v95) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L21
L25:
	;
	v97 = v94
	goto L27
L26:
	;
	v97 = v95
	goto L27
L27:
	;
	v98 = F_memcmp(m, v81+v90, v82+v90, v97)
	mBase = m.M
	if v98 != 0 {
		v137 = v98
		goto L20
	} else {
		goto L28
	}
L28:
	;
	if v94 != v95 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v147 = v94 - v95
	goto L19
L30:
	;
	goto L31
L31:
	;
	if v86 < int32(2) {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v103 = int32(1)
	v105 = int32(9)
	v107 = int32(_a_F_ltree_same_0)
	if v103 < v85 {
		v81 = v81 + (v94+v105)&v107
		v82 = v82 + (v95+v105)&v107
		v85 = v85 - v103
		v86 = v86 - v103
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L24
L34:
	;
	v156 = v13 + int32(8)
	if v151 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v158 = int32(0)
	goto L37
L36:
	;
	v158 = v43
	goto L37
L37:
	;
	v159 = v156 + v158
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159)+4)))
	v162 = v12 + int32(8)
	v163 = v162 + v158
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+4)))
	if v160 != v164 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	v166 = int32(0)
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159)+4)))
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+4)))
	if base.B2i32(v173 == v166)|base.B2i32(v176 == v166) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v250 != 0 {
		goto L12
	} else {
		goto L54
	}
L40:
	;
	v250 = v240
	goto L39
L41:
	;
	v240 = v173 - v176
	goto L40
L42:
	;
	v180 = int32(8)
	v184 = v159 + v180
	v185 = v163 + v180
	v188 = v176
	v189 = v173
	goto L43
L43:
	;
	v193 = int32(2)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32(v197) < base.Ui32(v198) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L41
L45:
	;
	v200 = v197
	goto L47
L46:
	;
	v200 = v198
	goto L47
L47:
	;
	v201 = F_memcmp(m, v184+v193, v185+v193, v200)
	mBase = m.M
	if v201 != 0 {
		v240 = v201
		goto L40
	} else {
		goto L48
	}
L48:
	;
	if v197 != v198 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v250 = v197 - v198
	goto L39
L50:
	;
	goto L51
L51:
	;
	if v189 < int32(2) {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v206 = int32(1)
	v208 = int32(9)
	v210 = int32(_a_F_ltree_same_0)
	if v206 < v188 {
		v184 = v184 + (v197+v208)&v210
		v185 = v185 + (v198+v208)&v210
		v188 = v188 - v206
		v189 = v189 - v206
		goto L43
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v252&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v255 = int32(0)
	goto L57
L56:
	;
	v255 = v43
	goto L57
L57:
	;
	v256 = v156 + v255
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v258&int32(2) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v261 = int32(0)
	goto L60
L59:
	;
	v261 = v43
	goto L60
L60:
	;
	v262 = v162 + v261
	v264 = v252 & int32(4)
	if v264 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v269 = v256
	goto L63
L62:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v269 = v256 + int32(base.Ui32(v265)>>(uint(int32(2))%32))
	goto L63
L63:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+4)))
	v272 = v258 & int32(4)
	if v272 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v277 = v262
	goto L66
L65:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v277 = v262 + int32(base.Ui32(v273)>>(uint(int32(2))%32))
	goto L66
L66:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)))
	if v270 != v278 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if v252&int32(2) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v283 = int32(0)
	goto L70
L69:
	;
	v283 = v43
	goto L70
L70:
	;
	v284 = v156 + v283
	if v258&int32(2) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v288 = int32(0)
	goto L73
L72:
	;
	v288 = v43
	goto L73
L73:
	;
	v289 = v162 + v288
	if v264 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v294 = v284
	goto L76
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v294 = v284 + int32(base.Ui32(v290)>>(uint(int32(2))%32))
	goto L76
L76:
	;
	if v272 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v299 = v289
	goto L79
L78:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v299 = v289 + int32(base.Ui32(v295)>>(uint(int32(2))%32))
	goto L79
L79:
	;
	v300 = int32(0)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+4)))
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+4)))
	if base.B2i32(v307 == v300)|base.B2i32(v310 == v300) != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	if v384 != 0 {
		goto L12
	} else {
		goto L95
	}
L81:
	;
	v384 = v374
	goto L80
L82:
	;
	v374 = v307 - v310
	goto L81
L83:
	;
	v314 = int32(8)
	v318 = v294 + v314
	v319 = v299 + v314
	v322 = v310
	v323 = v307
	goto L84
L84:
	;
	v327 = int32(2)
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318))))
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319))))
	if base.Ui32(v331) < base.Ui32(v332) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L82
L86:
	;
	v334 = v331
	goto L88
L87:
	;
	v334 = v332
	goto L88
L88:
	;
	v335 = F_memcmp(m, v318+v327, v319+v327, v334)
	mBase = m.M
	if v335 != 0 {
		v374 = v335
		goto L81
	} else {
		goto L89
	}
L89:
	;
	if v331 != v332 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v384 = v331 - v332
	goto L80
L91:
	;
	goto L92
L92:
	;
	if v323 < int32(2) {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	v340 = int32(1)
	v342 = int32(9)
	v344 = int32(_a_F_ltree_same_0)
	if v340 < v322 {
		v318 = v318 + (v331+v342)&v344
		v319 = v319 + (v332+v342)&v344
		v322 = v322 - v340
		v323 = v323 - v340
		goto L84
	} else {
		goto L94
	}
L94:
	;
	goto L85
L95:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v385)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	if v387&int32(2)|base.B2i32(v43 <= int32(0)) != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	v394 = int32(0)
	goto L97
L97:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+v156))))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+v162))))
	if v406 == v408 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v425 = int32(0)
	goto L14
L99:
	;
	v411 = v394 + int32(1)
	if v43 != v411 {
		v394 = v411
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	goto L12
}
func F_ltree_send(m *base.Module, l0 int32) int64 {
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v20 = F_palloc(m, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
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
				return int64(0)
			} else {
				F_enlargeStringInfo(m, v10, int32(1))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
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
						return int64(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int64(0)
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 << (uint(int32(2)) % 32)
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v96)
						}
					}
				}
			}
		}
	}
}
func F_ltree_textadd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = F_text_to_cstring(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v20 = F_DirectFunctionCall1Coll(m, int32(_a_F_ltree_textadd_0), int32(0), base.I64_extend_i32_u(v17))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v17)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v24 = base.I32_wrap_i64(v20)
						v25 = F_ltree_concat(m, v24, v8)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v24)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v29 != v8 {
									F_pfree(m, v8)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int64(0)
									} else {
										v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v33 != v15 {
											F_pfree(m, v15)
											mBase = m.M
											v36 = m.ExcPending
											if v36 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v25)
											}
										} else {
											return base.I64_extend_i32_u(v25)
										}
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v33 != v15 {
										F_pfree(m, v15)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v25)
										}
									} else {
										return base.I64_extend_i32_u(v25)
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
