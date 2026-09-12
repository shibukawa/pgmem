package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseComplexProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v7 != int32(6) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = int32(0)
	if v31 == v32 {
		v104 = v32
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v29 = F_get_expr_result_tupdesc(m, l2, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L12
	}
L3:
	;
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v15 = F_GetNSItemByRangeTablePosn(m, l0, v13, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v23 != int32(2249) {
		goto L2
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v20 = F_scanNSItemForColumn(m, l0, v15, v19, l1, l3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	return v20
L10:
	;
	v26 = F_expandRecordVariable(m, l0, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v31 = v26
	goto L1
L12:
	;
	v31 = v29
	goto L1
L13:
	;
	return v104
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v35 <= int32(0) {
		v104 = v32
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(0)
	goto L16
L16:
	;
	v52 = v31 + v35<<(uint(int32(4))%32) + int32(20) + v44*int32(100)
	v54 = v52 + int32(4)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v58 == int32(0) {
		v77 = v57
		v78 = v58
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v104 = int32(0)
	goto L13
L18:
	;
	v98 = v44 + int32(1)
	if v98 != v35 {
		v44 = v98
		goto L16
	} else {
		goto L30
	}
L19:
	;
	if v78-v77 != 0 {
		goto L18
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	if v57 != v58 {
		v77 = v57
		v78 = v58
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v62 = l1
	v63 = v54
	goto L23
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v66
		v78 = v67
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v77 = v66
	v78 = v67
	goto L20
L25:
	;
	v70 = int32(1)
	if v66 == v67 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+91)))
	if v80 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v82 = F_palloc0(m, int32(24))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v85 = v44 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(25)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v52)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v52)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v94
	return v82
L30:
	;
	goto L17
}
func F_ParseLongOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	v5 = int32(536735)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*int8)(unsafe.Add(mBase, _consts[1436])))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v207 = v206
	goto L61
L2:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
	if v73 == int32(61) {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	m.G0 = v11 + int32(32)
	v71 = v66 - l0
	goto L2
L4:
	;
	v18 = F___memset(m, v11, int32(0), int32(32))
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1436])))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1437])))
	if v14 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v15 = F___strchrnul(m, l0, v13)
	mBase = m.M
	v66 = v15
	goto L3
L8:
	;
	goto L7
L9:
	;
	v21 = v5
	v22 = v19
	goto L12
L10:
	;
	goto L11
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v43 == int32(0) {
		v66 = l0
		goto L3
	} else {
		goto L15
	}
L12:
	;
	v29 = v11 + int32(base.Ui32(v22)>>(uint(int32(3))%32))&int32(28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30 | v31<<(uint(v22)%32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v35 != 0 {
		v21 = v21 + v31
		v22 = v35
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	v47 = l0
	v48 = v43
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v48)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v56)>>(uint(v48)%32))&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v66 = v62
	goto L3
L18:
	;
	v66 = v47
	goto L3
L19:
	;
	goto L20
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v62 = v47 + int32(1)
	if v60 != 0 {
		v47 = v62
		v48 = v60
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v77 = v71 + int32(1)
	v78 = F_palloc(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v198 = F_pstrdup(m, l0)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L25
	} else {
		goto L60
	}
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
	if v77 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v196 = F_pstrdup(m, l0+v77)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L25
	} else {
		goto L59
	}
L28:
	;
	v192 = F_strlen(m, v188)
	mBase = m.M
	goto L27
L29:
	;
	v188 = l0
	goto L28
L30:
	;
	goto L31
L31:
	;
	v86 = v77 - int32(1)
	if (v78^l0)&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v185)
	v188 = v181
	goto L28
L33:
	;
	v166 = v161
	v167 = v162
	v168 = v163
	goto L55
L34:
	;
	if v156 == int32(0) {
		v181 = v154
		v182 = v155
		goto L32
	} else {
		goto L54
	}
L35:
	;
	v154 = l0
	v155 = v78
	v156 = v86
	goto L34
L36:
	;
	goto L37
L37:
	;
	v90 = int32(0)
	if l0&int32(3) == v90 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v123 == int32(0) {
		v181 = v120
		v182 = v121
		goto L32
	} else {
		goto L47
	}
L39:
	;
	v120 = l0
	v121 = v78
	v122 = v86
	v123 = base.B2i32(v86 != v90)
	goto L38
L40:
	;
	if v86 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v99 = l0
	v100 = v78
	v101 = v86
	goto L42
L42:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
	if v103 == int32(0) {
		v161 = v99
		v162 = v100
		v163 = v101
		goto L33
	} else {
		goto L44
	}
L43:
	;
	v120 = v114
	v121 = v108
	v122 = v110
	v123 = v112
	goto L38
L44:
	;
	v107 = int32(1)
	v108 = v100 + v107
	v110 = v101 - v107
	v111 = int32(0)
	v112 = base.B2i32(v110 != v111)
	v114 = v99 + v107
	if v114&int32(3) == v111 {
		v120 = v114
		v121 = v108
		v122 = v110
		v123 = v112
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if v110 != 0 {
		v99 = v114
		v100 = v108
		v101 = v110
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v126 == int32(0) {
		v154 = v120
		v155 = v121
		v156 = v122
		goto L34
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(v122) < base.Ui32(int32(4)) {
		v154 = v120
		v155 = v121
		v156 = v122
		goto L34
	} else {
		goto L49
	}
L49:
	;
	v132 = v120
	v133 = v121
	v134 = v122
	goto L50
L50:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v140 = int32(-2139062144)
	if (int32(16843008)-v137|v137)&v140 != v140 {
		v161 = v132
		v162 = v133
		v163 = v134
		goto L33
	} else {
		goto L52
	}
L51:
	;
	v154 = v148
	v155 = v146
	v156 = v150
	goto L34
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v137
	v145 = int32(4)
	v146 = v133 + v145
	v148 = v132 + v145
	v150 = v134 - v145
	if base.Ui32(int32(3)) < base.Ui32(v150) {
		v132 = v148
		v133 = v146
		v134 = v150
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v161 = v154
	v162 = v155
	v163 = v156
	goto L33
L55:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v170)
	if v170 == int32(0) {
		v181 = v166
		v182 = v167
		goto L32
	} else {
		goto L57
	}
L56:
	;
	v181 = v177
	v182 = v175
	goto L32
L57:
	;
	v174 = int32(1)
	v175 = v167 + v174
	v177 = v166 + v174
	v179 = v168 - v174
	if v179 != 0 {
		v166 = v177
		v167 = v175
		v168 = v179
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v204 = v196
	goto L1
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v198
	v204 = int32(0)
	goto L1
L61:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v211 != int32(45) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v207 = v207 + int32(1)
	goto L61
L64:
	;
	if v211 != 0 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v214 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v214)
	goto L63
L67:
	;
	return
}
func F_PreCommit_CheckForSerializationFailure(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v11 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v149+int32(3584))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v119+int32(3584))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L29
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v17 = F_LWLockAcquire(m, v13+int32(3584), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	return
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v21&int32(2056) == int32(8) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1112]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+32))
	v96 = v94 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v99 | int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v104+int32(3584))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v30 = v20 + int32(40)
	if v26 == v30 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v34 = v26
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	if v42&int32(9) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v81 != v30 {
		v34 = v81
		goto L12
	} else {
		goto L27
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = v41 + int32(40)
	if v45 == v49 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v52 = v45
	goto L18
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v20 != v60 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L14
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v70 != v49 {
		v52 = v70
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+108)))
	if v62&int32(41) != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v42&int32(2) != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = v42 | int32(8)
	goto L14
L26:
	;
	goto L19
L27:
	;
	goto L13
L28:
	;
	goto L5
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(139603), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errdetail_internal(m, int32(554237), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errhint(m, int32(624848), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(489684), int32(4727), int32(357179))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(139603), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errdetail_internal(m, int32(554491), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errhint(m, int32(624848), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(489684), int32(4763), int32(357179))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PreCommit_Portals(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	F_hash_seq_init(m, v9+int32(12), v14)
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
	v21 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L42
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L39
	}
L5:
	;
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = v21
	v28 = v2
	goto L9
L7:
	;
	v90 = v2
	goto L8
L8:
	;
	m.G0 = v9 + int32(32)
	return v90
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+84)))
	if v30 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v90 = v80
	goto L8
L11:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+85)))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v36 == int32(3) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v83 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L37
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v39 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+76)))
	if v51&int32(32) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v40 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v46
	v80 = v28
	goto L15
L22:
	;
	F_UnregisterSnapshotFromOwner(m, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = int32(0)
	goto L21
L25:
	;
	goto L24
L26:
	;
	F_hash_seq_term(m, v9+int32(12))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	if v50 == int32(0) {
		v80 = v28
		goto L15
	} else {
		goto L33
	}
L28:
	;
	if v36 != int32(2) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v50 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_HoldPortal(m, v29)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	F_PortalDrop(m, v29, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	F_hash_seq_init(m, v9+int32(12), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v80 = int32(1)
	goto L15
L37:
	;
	if v83 != 0 {
		v24 = v83
		v28 = v80
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	F_errmsg_internal(m, int32(444857), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(488237), int32(695), int32(150854))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(533243), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(488237), int32(738), int32(150854))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PreCommit_on_commit_actions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v1 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[453]))
	if v12 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = v1
	v21 = v1
	v23 = v1
	goto L6
L4:
	;
	v54 = v1
	v56 = v1
	goto L5
L5:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v19<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v29 != 0 {
		v45 = v21
		v46 = v23
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v45
	v56 = v46
	goto L5
L8:
	;
	v48 = v19 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v48 < v49 {
		v19 = v48
		v21 = v45
		v23 = v46
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	switch v30 - int32(2) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		v45 = v21
		v46 = v23
		goto L8
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v43 = F_lappend_oid(m, v23, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L15
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[164])))
	if v34&int32(1) == int32(0) {
		v45 = v21
		v46 = v23
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v40 = F_lappend_oid(m, v21, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v45 = v40
	v46 = v23
	goto L8
L15:
	;
	v45 = v21
	v46 = v43
	goto L8
L16:
	;
	goto L7
L17:
	;
	v57 = int32(0)
	if v54 == v57 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L20:
	;
	goto L19
L21:
	;
	F_heap_truncate_check_FKs(m, int32(0), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if int32(0) < v65 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v68 = v57
	v70 = v57
	goto L28
L26:
	;
	v90 = v57
	goto L27
L27:
	;
	F_heap_truncate_check_FKs(m, v90, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v68<<(uint(int32(2))%32))))
	v80 = F_table_open(m, v78, int32(8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	v90 = v82
	goto L27
L30:
	;
	v82 = F_lappend(m, v70, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v85 = v68 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v85 < v86 {
		v68 = v85
		v70 = v82
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v90 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	goto L20
L35:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v100 <= v99 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v106 = v99
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v106<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+119)))
	if v115 == int32(112) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	F_sequence_close(m, v113, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L48
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+188))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
	m.T0[v119].(func(*base.Module, int32))(m, v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	F_RelationTruncateIndexes(m, v113)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+112))
	if v125 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v129 = F_table_open(m, v125, int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+188))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+116))
	m.T0[v132].(func(*base.Module, int32))(m, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_RelationTruncateIndexes(m, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_sequence_close(m, v129, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	v145 = v106 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v145 < v146 {
		v106 = v145
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v168 = F_new_object_addresses(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if int32(0) < v170 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	v204 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L59
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1259)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180+v175<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v186
	F_add_exact_object_address(m, v9+int32(4), v168)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v195 = v175 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v195 < v196 {
		v175 = v195
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_PushActiveSnapshot(m, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_performMultipleDeletions(m, v168, int32(1), int32(5))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	goto L1
}
func F_PrefetchSharedBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v18
	v24 = F_BufTableHashCode(m, v10+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, _consts[86]))
		v34 = v27 + v24&int32(127)<<(uint(int32(7))%32) + int32(6912)
		v36 = F_LWLockAcquire(m, v34, int32(1))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			v40 = F_BufTableLookup(m, v10+int32(12), v24)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_LWLockRelease(m, v34)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					if v40 < int32(0) {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[258])))
						if v47&int32(1) != 0 {
							m.G0 = v10 + int32(32)
							return
						} else {
							v51 = F_smgrprefetch(m, l1, l2, l3, int32(1))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 == int32(0) {
								} else {
									v55 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v55)
								}
								m.G0 = v10 + int32(32)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40 + int32(1)
						m.G0 = v10 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_PrepareInplaceInvalidationState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v14
		} else {
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, _consts[122])) = v18
			*(*int64)(unsafe.Add(mBase, _consts[121])) = v18
		}
		*(*int32)(unsafe.Add(mBase, _consts[117])) = v5
		return v5
	}
}
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = F_MemoryContextAlloc(m, v4, int32(64))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_fmgr_info_cxt(m, l0, v6, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v6
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+60)) = uint8(v15)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v15)
			v19 = int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+46)) = uint16(v19)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+44)) = uint8(v15)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(1834)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v6
			return
		}
	}
}
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v12 != 0 {
		v14 = int32(1)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)))
		v14 = v13
	}
	if v14&int32(1) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(256443), v5)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(483530), int32(429), int32(406992))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v8 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(252306), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(483530), int32(411), int32(19502))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v11) < base.Ui32(int32(2)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		if int32(2) <= v14 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errcode(m, int32(16777538))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
					F_errmsg(m, int32(251792), v7+int32(16))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_errfinish(m, int32(484613), int32(3668), int32(311642))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if l0 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16777538))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
						F_errmsg(m, int32(249975), v7+int32(32))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errfinish(m, int32(484613), int32(3677), int32(311642))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v19 = int32(4365524)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[164]))
				*(*int32)(unsafe.Add(mBase, _consts[164])) = v21 | int32(4)
				m.G0 = v7 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errcode(m, int32(16777538))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(311242), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errfinish(m, int32(484613), int32(3658), int32(311642))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_ProcessCheckpointerInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, _consts[706]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _consts[705]))
			if v6 == int32(0) {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
				if v35 != 0 {
					F_ProcessLogMemoryContextInterrupt(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_SyncRepUpdateSyncStandbysDefined(m)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_UpdateFullPageWrites(m)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v21 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								if v21 == int32(0) {
									v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
									if v35 != 0 {
										F_ProcessLogMemoryContextInterrupt(m)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											return
										}
									} else {
										return
									}
								} else {
									F_errmsg_internal(m, int32(155296), int32(0))
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return
									} else {
										F_errfinish(m, int32(486346), int32(1390), int32(331543))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
											if v35 != 0 {
												F_ProcessLogMemoryContextInterrupt(m)
												mBase = m.M
												v37 = m.ExcPending
												if v37 != 0 {
													return
												} else {
													return
												}
											} else {
												return
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
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[705]))
		if v6 == int32(0) {
			v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
			if v35 != 0 {
				F_ProcessLogMemoryContextInterrupt(m)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_SyncRepUpdateSyncStandbysDefined(m)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_UpdateFullPageWrites(m)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v21 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 == int32(0) {
								v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
								if v35 != 0 {
									F_ProcessLogMemoryContextInterrupt(m)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							} else {
								F_errmsg_internal(m, int32(155296), int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									F_errfinish(m, int32(486346), int32(1390), int32(331543))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, _consts[707]))
										if v35 != 0 {
											F_ProcessLogMemoryContextInterrupt(m)
											mBase = m.M
											v37 = m.ExcPending
											if v37 != 0 {
												return
											} else {
												return
											}
										} else {
											return
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
}
func F_ProcessCopyOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
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
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v988 int32
	_ = v988
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1044 int32
	_ = v1044
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1125 int32
	_ = v1125
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1167 int32
	_ = v1167
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1249 int32
	_ = v1249
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1272 int32
	_ = v1272
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1294 int32
	_ = v1294
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1338 int32
	_ = v1338
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1426 int64
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int64
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1829 int32
	_ = v1829
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2015 int32
	_ = v2015
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2116 int32
	_ = v2116
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2151 int32
	_ = v2151
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2213 int32
	_ = v2213
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2279 int32
	_ = v2279
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2335 int32
	_ = v2335
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2356 int32
	_ = v2356
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2437 int32
	_ = v2437
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2453 int32
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2472 int32
	_ = v2472
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2500 int32
	_ = v2500
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2598 int64
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2620 int32
	_ = v2620
	var v2625 int32
	_ = v2625
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2639 int32
	_ = v2639
	var v2644 int32
	_ = v2644
	var v2648 int32
	_ = v2648
	var v2651 int32
	_ = v2651
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2672 int32
	_ = v2672
	var v2681 int32
	_ = v2681
	var v2686 int32
	_ = v2686
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2700 int32
	_ = v2700
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2731 int32
	_ = v2731
	var v2735 int32
	_ = v2735
	var v2740 int32
	_ = v2740
	var v2744 int32
	_ = v2744
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2763 int32
	_ = v2763
	var v2774 int32
	_ = v2774
	var v2779 int32
	_ = v2779
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2795 int32
	_ = v2795
	var v2800 int32
	_ = v2800
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2817 int32
	_ = v2817
	var v2822 int32
	_ = v2822
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(528)
	m.G0 = v20
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_palloc0(m, int32(112))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(-1)
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	return
L5:
	;
	v27 = v25
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2807 = m.ExcPending
	if v2807 != 0 {
		goto L4
	} else {
		goto L967
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2783 = m.ExcPending
	if v2783 != 0 {
		goto L4
	} else {
		goto L963
	}
L8:
	;
	v2507 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1524))))
	v2508 = F___strchrnul(m, v1619, v2507)
	mBase = m.M
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2508))))
	if v2510 == v2507&int32(255) {
		goto L869
	} else {
		goto L870
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L4
	} else {
		goto L856
	}
L10:
	;
	if l2 != 0 {
		v2506 = v2478
		goto L8
	} else {
		goto L854
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L4
	} else {
		goto L850
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L4
	} else {
		goto L846
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L4
	} else {
		goto L842
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L4
	} else {
		goto L838
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L4
	} else {
		goto L834
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L4
	} else {
		goto L830
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L4
	} else {
		goto L826
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L4
	} else {
		goto L822
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L4
	} else {
		goto L818
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L4
	} else {
		goto L814
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L4
	} else {
		goto L810
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L4
	} else {
		goto L806
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L4
	} else {
		goto L801
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L4
	} else {
		goto L796
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L4
	} else {
		goto L792
	}
L26:
	;
	F_errorConflictingDefElem(m, v56, l0)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L4
	} else {
		goto L791
	}
L27:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v1490 != 0 {
		goto L525
	} else {
		goto L526
	}
L28:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v32 <= int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v43 = v5
	v44 = v5
	v45 = v5
	v46 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	goto L30
L30:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v43<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = int32(110235)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _consts[393])))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v62 == int32(0) {
		v81 = v61
		v82 = v62
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L27
L32:
	;
	v1469 = v43 + int32(1)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1469 < v1470 {
		v43 = v1469
		v44 = v1460
		v45 = v1461
		v46 = v1462
		v48 = v1464
		v49 = v1465
		v50 = v1466
		goto L30
	} else {
		goto L522
	}
L33:
	;
	if v82-v81 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	goto L33
L35:
	;
	if v61 != v62 {
		v81 = v61
		v82 = v62
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v66 = v57
	v67 = v58
	goto L37
L37:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v70
		v82 = v71
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v81 = v70
	v82 = v71
	goto L34
L39:
	;
	v74 = int32(1)
	if v70 == v71 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v86 = F_defGetString(m, v56)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v202 = int32(336347)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v206 == int32(0) {
		v225 = v205
		v226 = v206
		goto L85
	} else {
		goto L86
	}
L44:
	;
	if v44&int32(1) != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v90 = int32(62632)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[395])))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v94 == int32(0) {
		v113 = v93
		v114 = v94
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v114-v113 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	goto L46
L48:
	;
	if v93 != v94 {
		v113 = v93
		v114 = v94
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v98 = v86
	v99 = v90
	goto L50
L50:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v103 == int32(0) {
		v113 = v102
		v114 = v103
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v113 = v102
	v114 = v103
	goto L47
L52:
	;
	v106 = int32(1)
	if v102 == v103 {
		v98 = v98 + v106
		v99 = v99 + v106
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v1460 = int32(1)
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L55:
	;
	goto L56
L56:
	;
	v119 = int32(32035)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _consts[396])))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v123 == int32(0) {
		v142 = v122
		v143 = v123
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v143-v142 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	goto L57
L59:
	;
	if v122 != v123 {
		v142 = v122
		v143 = v123
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v127 = v86
	v128 = v119
	goto L61
L61:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	if v132 == int32(0) {
		v142 = v131
		v143 = v132
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v142 = v131
	v143 = v132
	goto L58
L63:
	;
	v135 = int32(1)
	if v131 == v132 {
		v127 = v127 + v135
		v128 = v128 + v135
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)) = uint8(v147)
	v1460 = v147
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L66:
	;
	goto L67
L67:
	;
	v150 = int32(17607)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[397])))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v154 == int32(0) {
		v173 = v153
		v174 = v154
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v174-v173 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	goto L68
L70:
	;
	if v153 != v154 {
		v173 = v153
		v174 = v154
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v158 = v86
	v159 = v150
	goto L72
L72:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v163 == int32(0) {
		v173 = v162
		v174 = v163
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v173 = v162
	v174 = v163
	goto L69
L74:
	;
	v166 = int32(1)
	if v162 == v163 {
		v158 = v158 + v166
		v159 = v159 + v166
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)) = uint8(v178)
	v1460 = v178
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+320)) = v86
	F_errmsg(m, int32(430234), v20+int32(320))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(483587), int32(576), int32(135850))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	if v226-v225 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	goto L84
L86:
	;
	if v205 != v206 {
		v225 = v205
		v226 = v206
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v210 = v57
	v211 = v202
	goto L88
L88:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	if v215 == int32(0) {
		v225 = v214
		v226 = v215
		goto L85
	} else {
		goto L90
	}
L89:
	;
	v225 = v214
	v226 = v215
	goto L85
L90:
	;
	v218 = int32(1)
	if v214 == v215 {
		v210 = v210 + v218
		v211 = v211 + v218
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v46&int32(1) != 0 {
		goto L26
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v236 = int32(212732)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, _consts[398])))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v240 == int32(0) {
		v259 = v239
		v260 = v240
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v232 = F_defGetBoolean(m, v56)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v232)
	v1460 = v44
	v1461 = v45
	v1462 = int32(1)
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L97:
	;
	if v260-v259 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	goto L97
L99:
	;
	if v239 != v240 {
		v259 = v239
		v260 = v240
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v244 = v57
	v245 = v236
	goto L101
L101:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v249 == int32(0) {
		v259 = v248
		v260 = v249
		goto L98
	} else {
		goto L103
	}
L102:
	;
	v259 = v248
	v260 = v249
	goto L98
L103:
	;
	v252 = int32(1)
	if v248 == v249 {
		v244 = v244 + v252
		v245 = v245 + v252
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v264 != 0 {
		goto L26
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v268 = int32(298338)
	v271 = int32(*(*uint8)(unsafe.Add(mBase, _consts[399])))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v272 == int32(0) {
		v291 = v271
		v292 = v272
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v265 = F_defGetString(m, v56)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v265
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L110:
	;
	if v292-v291 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	goto L110
L112:
	;
	if v271 != v272 {
		v291 = v271
		v292 = v272
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v276 = v57
	v277 = v268
	goto L114
L114:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	if v281 == int32(0) {
		v291 = v280
		v292 = v281
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v291 = v280
	v292 = v281
	goto L111
L116:
	;
	v284 = int32(1)
	if v280 == v281 {
		v276 = v276 + v284
		v277 = v277 + v284
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v296 != 0 {
		goto L26
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v300 = int32(96979)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _consts[400])))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v304 == int32(0) {
		v323 = v303
		v324 = v304
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v297 = F_defGetString(m, v56)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v297
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L123:
	;
	if v324-v323 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	goto L123
L125:
	;
	if v303 != v304 {
		v323 = v303
		v324 = v304
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v308 = v57
	v309 = v300
	goto L127
L127:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v313 == int32(0) {
		v323 = v312
		v324 = v313
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v323 = v312
	v324 = v313
	goto L124
L129:
	;
	v316 = int32(1)
	if v312 == v313 {
		v308 = v308 + v316
		v309 = v309 + v316
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v328 != 0 {
		goto L26
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v332 = int32(224167)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _consts[401])))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v336 == int32(0) {
		v355 = v335
		v356 = v336
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v329 = F_defGetString(m, v56)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v329
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L136:
	;
	if v356-v355 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	goto L136
L138:
	;
	if v335 != v336 {
		v355 = v335
		v356 = v336
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v340 = v57
	v341 = v332
	goto L140
L140:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if v345 == int32(0) {
		v355 = v344
		v356 = v345
		goto L137
	} else {
		goto L142
	}
L141:
	;
	v355 = v344
	v356 = v345
	goto L137
L142:
	;
	v348 = int32(1)
	if v344 == v345 {
		v340 = v340 + v348
		v341 = v341 + v348
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	if v45&int32(1) != 0 {
		goto L26
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v604 = int32(343046)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, _consts[402])))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v608 == int32(0) {
		v627 = v607
		v628 = v608
		goto L227
	} else {
		goto L228
	}
L147:
	;
	v362 = int32(1)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v364 == int32(0) {
		v600 = v362
		goto L148
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v600
	v1460 = v44
	v1461 = v362
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L149:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v367 == int32(465) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if base.Ui32(v370) < base.Ui32(int32(2)) {
		v600 = v370
		goto L148
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v373 = F_defGetString(m, v56)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L154
	}
L153:
	;
	goto L6
L154:
	;
	v378 = v373
	v379 = int32(338217)
	goto L156
L155:
	;
	if v416 == int32(0) {
		v600 = v362
		goto L148
	} else {
		goto L168
	}
L156:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == v383 {
		v405 = v382
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v416 = int32(0)
	goto L155
L158:
	;
	v407 = int32(1)
	if v405 != 0 {
		v378 = v378 + v407
		v379 = v379 + v407
		goto L156
	} else {
		goto L167
	}
L159:
	;
	if base.Ui32((v382-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v393 = v382 | int32(32)
	goto L162
L161:
	;
	v393 = v382
	goto L162
L162:
	;
	if base.Ui32((v383-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v402 = v383 | int32(32)
	goto L165
L164:
	;
	v402 = v383
	goto L165
L165:
	;
	if v393 == v402 {
		v405 = v393
		goto L158
	} else {
		goto L166
	}
L166:
	;
	v416 = v393 - v402
	goto L155
L167:
	;
	goto L157
L168:
	;
	v423 = v373
	v424 = int32(355036)
	goto L170
L169:
	;
	if v461 == int32(0) {
		v600 = int32(0)
		goto L148
	} else {
		goto L182
	}
L170:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v427 == v428 {
		v450 = v427
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v461 = int32(0)
	goto L169
L172:
	;
	v452 = int32(1)
	if v450 != 0 {
		v423 = v423 + v452
		v424 = v424 + v452
		goto L170
	} else {
		goto L181
	}
L173:
	;
	if base.Ui32((v427-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v438 = v427 | int32(32)
	goto L176
L175:
	;
	v438 = v427
	goto L176
L176:
	;
	if base.Ui32((v428-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v447 = v428 | int32(32)
	goto L179
L178:
	;
	v447 = v428
	goto L179
L179:
	;
	if v438 == v447 {
		v450 = v438
		goto L172
	} else {
		goto L180
	}
L180:
	;
	v461 = v438 - v447
	goto L169
L181:
	;
	goto L171
L182:
	;
	v468 = v373
	v469 = int32(268678)
	goto L184
L183:
	;
	if v506 == int32(0) {
		v600 = int32(1)
		goto L148
	} else {
		goto L196
	}
L184:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	if v472 == v473 {
		v495 = v472
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v506 = int32(0)
	goto L183
L186:
	;
	v497 = int32(1)
	if v495 != 0 {
		v468 = v468 + v497
		v469 = v469 + v497
		goto L184
	} else {
		goto L195
	}
L187:
	;
	if base.Ui32((v472-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v483 = v472 | int32(32)
	goto L190
L189:
	;
	v483 = v472
	goto L190
L190:
	;
	if base.Ui32((v473-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v492 = v473 | int32(32)
	goto L193
L192:
	;
	v492 = v473
	goto L193
L193:
	;
	if v483 == v492 {
		v495 = v483
		goto L186
	} else {
		goto L194
	}
L194:
	;
	v506 = v483 - v492
	goto L183
L195:
	;
	goto L185
L196:
	;
	v513 = v373
	v514 = int32(333024)
	goto L198
L197:
	;
	if v551 == int32(0) {
		v600 = int32(0)
		goto L148
	} else {
		goto L210
	}
L198:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if v517 == v518 {
		v540 = v517
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v551 = int32(0)
	goto L197
L200:
	;
	v542 = int32(1)
	if v540 != 0 {
		v513 = v513 + v542
		v514 = v514 + v542
		goto L198
	} else {
		goto L209
	}
L201:
	;
	if base.Ui32((v517-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v528 = v517 | int32(32)
	goto L204
L203:
	;
	v528 = v517
	goto L204
L204:
	;
	if base.Ui32((v518-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v537 = v518 | int32(32)
	goto L207
L206:
	;
	v537 = v518
	goto L207
L207:
	;
	if v528 == v537 {
		v540 = v528
		goto L200
	} else {
		goto L208
	}
L208:
	;
	v551 = v528 - v537
	goto L197
L209:
	;
	goto L199
L210:
	;
	v557 = v373
	v558 = int32(320052)
	goto L212
L211:
	;
	if l2|v595 == int32(0) {
		goto L25
	} else {
		goto L224
	}
L212:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if v561 == v562 {
		v584 = v561
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v595 = int32(0)
	goto L211
L214:
	;
	v586 = int32(1)
	if v584 != 0 {
		v557 = v557 + v586
		v558 = v558 + v586
		goto L212
	} else {
		goto L223
	}
L215:
	;
	if base.Ui32((v561-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v572 = v561 | int32(32)
	goto L218
L217:
	;
	v572 = v561
	goto L218
L218:
	;
	if base.Ui32((v562-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v581 = v562 | int32(32)
	goto L221
L220:
	;
	v581 = v562
	goto L221
L221:
	;
	if v572 == v581 {
		v584 = v572
		goto L214
	} else {
		goto L222
	}
L222:
	;
	v595 = v572 - v581
	goto L211
L223:
	;
	goto L213
L224:
	;
	if v595 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	v600 = int32(2)
	goto L148
L226:
	;
	if v628-v627 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L227:
	;
	goto L226
L228:
	;
	if v607 != v608 {
		v627 = v607
		v628 = v608
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v612 = v57
	v613 = v604
	goto L230
L230:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613)+1)))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v617 == int32(0) {
		v627 = v616
		v628 = v617
		goto L227
	} else {
		goto L232
	}
L231:
	;
	v627 = v616
	v628 = v617
	goto L227
L232:
	;
	v620 = int32(1)
	if v616 == v617 {
		v612 = v612 + v620
		v613 = v613 + v620
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v632 != 0 {
		goto L26
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v636 = int32(365495)
	v639 = int32(*(*uint8)(unsafe.Add(mBase, _consts[403])))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v640 == int32(0) {
		v659 = v639
		v660 = v640
		goto L240
	} else {
		goto L241
	}
L237:
	;
	v633 = F_defGetString(m, v56)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v633
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L239:
	;
	if v660-v659 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L240:
	;
	goto L239
L241:
	;
	if v639 != v640 {
		v659 = v639
		v660 = v640
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v644 = v57
	v645 = v636
	goto L243
L243:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+1)))
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+1)))
	if v649 == int32(0) {
		v659 = v648
		v660 = v649
		goto L240
	} else {
		goto L245
	}
L244:
	;
	v659 = v648
	v660 = v649
	goto L240
L245:
	;
	v652 = int32(1)
	if v648 == v649 {
		v644 = v644 + v652
		v645 = v645 + v652
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v664 != 0 {
		goto L26
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v668 = int32(343040)
	v671 = int32(*(*uint8)(unsafe.Add(mBase, _consts[404])))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v672 == int32(0) {
		v691 = v671
		v692 = v672
		goto L254
	} else {
		goto L255
	}
L250:
	;
	v665 = F_defGetString(m, v56)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v665
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v700
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L253:
	;
	if v692-v691 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L254:
	;
	goto L253
L255:
	;
	if v671 != v672 {
		v691 = v671
		v692 = v672
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v676 = v57
	v677 = v668
	goto L257
L257:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v681 == int32(0) {
		v691 = v680
		v692 = v681
		goto L254
	} else {
		goto L259
	}
L258:
	;
	v691 = v680
	v692 = v681
	goto L254
L259:
	;
	v684 = int32(1)
	if v680 == v681 {
		v676 = v676 + v684
		v677 = v677 + v684
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v696 != 0 {
		goto L26
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v733 = int32(296985)
	v736 = int32(*(*uint8)(unsafe.Add(mBase, _consts[405])))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v737 == int32(0) {
		v756 = v736
		v757 = v737
		goto L276
	} else {
		goto L277
	}
L264:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v697 == int32(1) {
		goto L26
	} else {
		goto L265
	}
L265:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v700 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L270
	}
L267:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	if v703 == int32(1) {
		goto L252
	} else {
		goto L268
	}
L268:
	;
	if v703 != int32(77) {
		goto L266
	} else {
		goto L269
	}
L269:
	;
	v708 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)) = uint8(v708)
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L270:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+368)) = v718
	F_errmsg(m, int32(160664), v20+int32(368))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(483587), int32(635), int32(135850))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	if v757-v756 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L276:
	;
	goto L275
L277:
	;
	if v736 != v737 {
		v756 = v736
		v757 = v737
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v741 = v57
	v742 = v733
	goto L279
L279:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+1)))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741)+1)))
	if v746 == int32(0) {
		v756 = v745
		v757 = v746
		goto L276
	} else {
		goto L281
	}
L280:
	;
	v756 = v745
	v757 = v746
	goto L276
L281:
	;
	v749 = int32(1)
	if v745 == v746 {
		v741 = v741 + v749
		v742 = v742 + v749
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v761 != 0 {
		goto L26
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v799 = int32(297033)
	v802 = int32(*(*uint8)(unsafe.Add(mBase, _consts[406])))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v803 == int32(0) {
		v822 = v802
		v823 = v803
		goto L300
	} else {
		goto L301
	}
L286:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v762 == int32(1) {
		goto L26
	} else {
		goto L287
	}
L287:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v765 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L294
	}
L289:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	if v768 != int32(1) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	if v768 != int32(77) {
		goto L288
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v765
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L293:
	;
	v773 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)) = uint8(v773)
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L294:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L295
	}
L295:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+384)) = v784
	F_errmsg(m, int32(160664), v20+int32(384))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L296
	}
L296:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(483587), int32(650), int32(135850))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	if v823-v822 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L300:
	;
	goto L299
L301:
	;
	if v802 != v803 {
		v822 = v802
		v823 = v803
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v807 = v57
	v808 = v799
	goto L303
L303:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808)+1)))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+1)))
	if v812 == int32(0) {
		v822 = v811
		v823 = v812
		goto L300
	} else {
		goto L305
	}
L304:
	;
	v822 = v811
	v823 = v812
	goto L300
L305:
	;
	v815 = int32(1)
	if v811 == v812 {
		v807 = v807 + v815
		v808 = v808 + v815
		goto L303
	} else {
		goto L306
	}
L306:
	;
	goto L304
L307:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v827 != 0 {
		goto L26
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v865 = int32(19958)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, _consts[407])))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v869 == int32(0) {
		v888 = v868
		v889 = v869
		goto L324
	} else {
		goto L325
	}
L310:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v828 == int32(1) {
		goto L26
	} else {
		goto L311
	}
L311:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v831 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L318
	}
L313:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	if v834 != int32(1) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	if v834 != int32(77) {
		goto L312
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v831
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L317:
	;
	v839 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)) = uint8(v839)
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L318:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+400)) = v850
	F_errmsg(m, int32(160664), v20+int32(400))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L320
	}
L320:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v857)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L4
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(483587), int32(665), int32(135850))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	if v889-v888 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L324:
	;
	goto L323
L325:
	;
	if v868 != v869 {
		v888 = v868
		v889 = v869
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v873 = v57
	v874 = v865
	goto L327
L327:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)))
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+1)))
	if v878 == int32(0) {
		v888 = v877
		v889 = v878
		goto L324
	} else {
		goto L329
	}
L328:
	;
	v888 = v877
	v889 = v878
	goto L324
L329:
	;
	v881 = int32(1)
	if v877 == v878 {
		v873 = v873 + v881
		v874 = v874 + v881
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)))
	if v893 == int32(1) {
		goto L26
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v925 = int32(330076)
	v928 = int32(*(*uint8)(unsafe.Add(mBase, _consts[408])))
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v929 == int32(0) {
		v948 = v928
		v949 = v929
		goto L346
	} else {
		goto L347
	}
L334:
	;
	v896 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)) = uint8(v896)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v898 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L4
	} else {
		goto L340
	}
L336:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	if v899 != int32(1) {
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v898
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = v50
	goto L32
L339:
	;
	goto L338
L340:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+416)) = v910
	F_errmsg(m, int32(160664), v20+int32(416))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(483587), int32(684), int32(135850))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L4
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	if v949-v948 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L346:
	;
	goto L345
L347:
	;
	if v928 != v929 {
		v948 = v928
		v949 = v929
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v933 = v57
	v934 = v925
	goto L349
L349:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+1)))
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)))
	if v938 == int32(0) {
		v948 = v937
		v949 = v938
		goto L346
	} else {
		goto L351
	}
L350:
	;
	v948 = v937
	v949 = v938
	goto L346
L351:
	;
	v941 = int32(1)
	if v937 == v938 {
		v933 = v933 + v941
		v934 = v934 + v941
		goto L349
	} else {
		goto L352
	}
L352:
	;
	goto L350
L353:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= v953 {
		goto L26
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1075 = int32(208580)
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, _consts[409])))
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v1079 == int32(0) {
		v1098 = v1078
		v1099 = v1079
		goto L391
	} else {
		goto L392
	}
L356:
	;
	v956 = F_defGetString(m, v56)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L4
	} else {
		goto L357
	}
L357:
	;
	v965 = m.G0
	v967 = v965 + int32(-64)
	m.G0 = v967
	v969 = int32(-1)
	if v956 == int32(0) {
		v1044 = v969
		goto L359
	} else {
		goto L360
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1044
	if int32(0) <= v1044 {
		v1460 = v44
		v1461 = v45
		v1462 = v46
		v1464 = v48
		v1465 = v49
		v1466 = v50
		goto L32
	} else {
		goto L384
	}
L359:
	;
	m.G0 = v967 - int32(-64)
	goto L358
L360:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	if v972 == int32(0) {
		v1044 = v969
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v975 = F_strlen(m, v956)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v975) {
		v1044 = v969
		goto L359
	} else {
		goto L362
	}
L362:
	;
	v978 = v956
	v979 = v972
	v980 = v967
	goto L363
L363:
	;
	v988 = F_isalnum(m, v979&int32(255))
	mBase = m.M
	if v988 != 0 {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	v1005 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1001))) = uint8(v1005)
	v1009 = int32(*(*int8)(unsafe.Add(mBase, uint32(v967))))
	v1010 = int32(1818896)
	v1011 = int32(1818256)
	goto L372
L365:
	;
	if base.Ui32((v979-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	v1001 = v980
	goto L367
L367:
	;
	v1003 = v978 + int32(1)
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003))))
	if v1004 != 0 {
		v978 = v1003
		v979 = v1004
		v980 = v1001
		goto L363
	} else {
		goto L371
	}
L368:
	;
	v997 = v979 | int32(32)
	goto L370
L369:
	;
	v997 = v979
	goto L370
L370:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v980))) = uint8(v997)
	v1001 = v980 + int32(1)
	goto L367
L371:
	;
	goto L364
L372:
	;
	v1023 = v1011 + (v1010-v1011)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v1025 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1024))))
	v1026 = v1009 - v1025
	if v1026 != 0 {
		v1029 = v1026
		goto L374
	} else {
		goto L375
	}
L373:
	;
	v1044 = v969
	goto L359
L374:
	;
	v1033 = base.B2i32(v1029 < int32(0))
	if v1029 < int32(0) {
		goto L377
	} else {
		goto L378
	}
L375:
	;
	v1027 = F_strcmp(m, v967, v1024)
	mBase = m.M
	if v1027 != 0 {
		v1029 = v1027
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+4))
	v1044 = v1028
	goto L359
L377:
	;
	v1034 = v1023 - int32(8)
	goto L379
L378:
	;
	v1034 = v1010
	goto L379
L379:
	;
	if v1029 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1037 = v1011
	goto L382
L381:
	;
	v1037 = v1023 + int32(8)
	goto L382
L382:
	;
	if base.Ui32(v1037) <= base.Ui32(v1034) {
		v1010 = v1034
		v1011 = v1037
		goto L372
	} else {
		goto L383
	}
L383:
	;
	goto L373
L384:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+432)) = v1060
	F_errmsg(m, int32(374944), v20+int32(432))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1067)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(483587), int32(696), int32(135850))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	if v1099-v1098 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L391:
	;
	goto L390
L392:
	;
	if v1078 != v1079 {
		v1098 = v1078
		v1099 = v1079
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v1083 = v57
	v1084 = v1075
	goto L394
L394:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084)+1)))
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+1)))
	if v1088 == int32(0) {
		v1098 = v1087
		v1099 = v1088
		goto L391
	} else {
		goto L396
	}
L395:
	;
	v1098 = v1087
	v1099 = v1088
	goto L391
L396:
	;
	v1091 = int32(1)
	if v1087 == v1088 {
		v1083 = v1083 + v1091
		v1084 = v1084 + v1091
		goto L394
	} else {
		goto L397
	}
L397:
	;
	goto L395
L398:
	;
	if v48 != 0 {
		goto L26
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1196 = int32(10827)
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[410])))
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v1200 == int32(0) {
		v1219 = v1199
		v1220 = v1200
		goto L435
	} else {
		goto L436
	}
L401:
	;
	v1103 = F_defGetString(m, v56)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L4
	} else {
		goto L402
	}
L402:
	;
	if l2 == int32(0) {
		goto L24
	} else {
		goto L403
	}
L403:
	;
	v1110 = v1103
	v1111 = int32(230447)
	goto L405
L404:
	;
	if v1148 != 0 {
		goto L417
	} else {
		goto L418
	}
L405:
	;
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	if v1114 == v1115 {
		v1137 = v1114
		goto L407
	} else {
		goto L408
	}
L406:
	;
	v1148 = int32(0)
	goto L404
L407:
	;
	v1139 = int32(1)
	if v1137 != 0 {
		v1110 = v1110 + v1139
		v1111 = v1111 + v1139
		goto L405
	} else {
		goto L416
	}
L408:
	;
	if base.Ui32((v1114-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1125 = v1114 | int32(32)
	goto L411
L410:
	;
	v1125 = v1114
	goto L411
L411:
	;
	if base.Ui32((v1115-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1134 = v1115 | int32(32)
	goto L414
L413:
	;
	v1134 = v1115
	goto L414
L414:
	;
	if v1125 == v1134 {
		v1137 = v1125
		goto L407
	} else {
		goto L415
	}
L415:
	;
	v1148 = v1125 - v1134
	goto L404
L416:
	;
	goto L406
L417:
	;
	v1152 = v1103
	v1153 = int32(358147)
	goto L421
L418:
	;
	v1193 = int32(0)
	goto L419
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v1193
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = int32(1)
	v1465 = v49
	v1466 = v50
	goto L32
L420:
	;
	if v1190 != 0 {
		goto L23
	} else {
		goto L433
	}
L421:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152))))
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153))))
	if v1156 == v1157 {
		v1179 = v1156
		goto L423
	} else {
		goto L424
	}
L422:
	;
	v1190 = int32(0)
	goto L420
L423:
	;
	v1181 = int32(1)
	if v1179 != 0 {
		v1152 = v1152 + v1181
		v1153 = v1153 + v1181
		goto L421
	} else {
		goto L432
	}
L424:
	;
	if base.Ui32((v1156-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1167 = v1156 | int32(32)
	goto L427
L426:
	;
	v1167 = v1156
	goto L427
L427:
	;
	if base.Ui32((v1157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1176 = v1157 | int32(32)
	goto L430
L429:
	;
	v1176 = v1157
	goto L430
L430:
	;
	if v1167 == v1176 {
		v1179 = v1167
		goto L423
	} else {
		goto L431
	}
L431:
	;
	v1190 = v1167 - v1176
	goto L420
L432:
	;
	goto L422
L433:
	;
	v1193 = int32(1)
	goto L419
L434:
	;
	if v1220-v1219 == int32(0) {
		goto L442
	} else {
		goto L443
	}
L435:
	;
	goto L434
L436:
	;
	if v1199 != v1200 {
		v1219 = v1199
		v1220 = v1200
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1204 = v57
	v1205 = v1196
	goto L438
L438:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205)+1)))
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1204)+1)))
	if v1209 == int32(0) {
		v1219 = v1208
		v1220 = v1209
		goto L435
	} else {
		goto L440
	}
L439:
	;
	v1219 = v1208
	v1220 = v1209
	goto L435
L440:
	;
	v1212 = int32(1)
	if v1208 == v1209 {
		v1204 = v1204 + v1212
		v1205 = v1205 + v1212
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	if v49 != 0 {
		goto L26
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v1390 = int32(99257)
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, _consts[411])))
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v1394 == int32(0) {
		v1413 = v1393
		v1414 = v1394
		goto L498
	} else {
		goto L499
	}
L445:
	;
	v1224 = m.G0
	v1226 = v1224 - int32(16)
	m.G0 = v1226
	v1229 = F_defGetString(m, v56)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L4
	} else {
		goto L449
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v1363
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = int32(1)
	v1466 = v50
	goto L32
L447:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L4
	} else {
		goto L492
	}
L448:
	;
	m.G0 = v1226 + int32(16)
	goto L446
L449:
	;
	v1234 = v1229
	v1235 = int32(94958)
	goto L451
L450:
	;
	if v1272 == int32(0) {
		v1363 = int32(-1)
		goto L448
	} else {
		goto L463
	}
L451:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1234))))
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	if v1238 == v1239 {
		v1261 = v1238
		goto L453
	} else {
		goto L454
	}
L452:
	;
	v1272 = int32(0)
	goto L450
L453:
	;
	v1263 = int32(1)
	if v1261 != 0 {
		v1234 = v1234 + v1263
		v1235 = v1235 + v1263
		goto L451
	} else {
		goto L462
	}
L454:
	;
	if base.Ui32((v1238-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1249 = v1238 | int32(32)
	goto L457
L456:
	;
	v1249 = v1238
	goto L457
L457:
	;
	if base.Ui32((v1239-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v1258 = v1239 | int32(32)
	goto L460
L459:
	;
	v1258 = v1239
	goto L460
L460:
	;
	if v1249 == v1258 {
		v1261 = v1249
		goto L453
	} else {
		goto L461
	}
L461:
	;
	v1272 = v1249 - v1258
	goto L450
L462:
	;
	goto L452
L463:
	;
	v1279 = v1229
	v1280 = int32(96979)
	goto L465
L464:
	;
	if v1317 == int32(0) {
		v1363 = int32(0)
		goto L448
	} else {
		goto L477
	}
L465:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279))))
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	if v1283 == v1284 {
		v1306 = v1283
		goto L467
	} else {
		goto L468
	}
L466:
	;
	v1317 = int32(0)
	goto L464
L467:
	;
	v1308 = int32(1)
	if v1306 != 0 {
		v1279 = v1279 + v1308
		v1280 = v1280 + v1308
		goto L465
	} else {
		goto L476
	}
L468:
	;
	if base.Ui32((v1283-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v1294 = v1283 | int32(32)
	goto L471
L470:
	;
	v1294 = v1283
	goto L471
L471:
	;
	if base.Ui32((v1284-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v1303 = v1284 | int32(32)
	goto L474
L473:
	;
	v1303 = v1284
	goto L474
L474:
	;
	if v1294 == v1303 {
		v1306 = v1294
		goto L467
	} else {
		goto L475
	}
L475:
	;
	v1317 = v1294 - v1303
	goto L464
L476:
	;
	goto L466
L477:
	;
	v1323 = v1229
	v1324 = int32(354917)
	goto L479
L478:
	;
	if v1361 != 0 {
		goto L447
	} else {
		goto L491
	}
L479:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1323))))
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1324))))
	if v1327 == v1328 {
		v1350 = v1327
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v1361 = int32(0)
	goto L478
L481:
	;
	v1352 = int32(1)
	if v1350 != 0 {
		v1323 = v1323 + v1352
		v1324 = v1324 + v1352
		goto L479
	} else {
		goto L490
	}
L482:
	;
	if base.Ui32((v1327-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v1338 = v1327 | int32(32)
	goto L485
L484:
	;
	v1338 = v1327
	goto L485
L485:
	;
	if base.Ui32((v1328-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1347 = v1328 | int32(32)
	goto L488
L487:
	;
	v1347 = v1328
	goto L488
L488:
	;
	if v1338 == v1347 {
		v1350 = v1338
		goto L481
	} else {
		goto L489
	}
L489:
	;
	v1361 = v1338 - v1347
	goto L478
L490:
	;
	goto L480
L491:
	;
	v1363 = int32(1)
	goto L448
L492:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+4)) = v1229
	*(*int32)(unsafe.Add(mBase, uint32(v1226))) = int32(499189)
	F_errmsg(m, int32(430266), v1226)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L4
	} else {
		goto L494
	}
L494:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1380)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L4
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(483587), int32(514), int32(410781))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L4
	} else {
		goto L496
	}
L496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L497:
	;
	if v1414-v1413 == int32(0) {
		goto L505
	} else {
		goto L506
	}
L498:
	;
	goto L497
L499:
	;
	if v1393 != v1394 {
		v1413 = v1393
		v1414 = v1394
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v1398 = v57
	v1399 = v1390
	goto L501
L501:
	;
	v1402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399)+1)))
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1398)+1)))
	if v1403 == int32(0) {
		v1413 = v1402
		v1414 = v1403
		goto L498
	} else {
		goto L503
	}
L502:
	;
	v1413 = v1402
	v1414 = v1403
	goto L498
L503:
	;
	v1406 = int32(1)
	if v1402 == v1403 {
		v1398 = v1398 + v1406
		v1399 = v1399 + v1406
		goto L501
	} else {
		goto L504
	}
L504:
	;
	goto L502
L505:
	;
	if v50 != 0 {
		goto L26
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L4
	} else {
		goto L517
	}
L508:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v1418 == int32(0) {
		goto L22
	} else {
		goto L509
	}
L509:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1418)))
	if v1421 == int32(468) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	if v1430 <= int64(0) {
		goto L21
	} else {
		goto L516
	}
L511:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1418)+4))
	v1426 = F_pg_strtoint64_safe(m, v1424, int32(0))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L4
	} else {
		goto L514
	}
L512:
	;
	goto L513
L513:
	;
	v1428 = F_defGetInt64(m, v56)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L4
	} else {
		goto L515
	}
L514:
	;
	v1430 = v1426
	goto L510
L515:
	;
	v1430 = v1428
	goto L510
L516:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v1430
	v1460 = v44
	v1461 = v45
	v1462 = v46
	v1464 = v48
	v1465 = v49
	v1466 = int32(1)
	goto L32
L517:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L518
	}
L518:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+512)) = v1442
	F_errmsg(m, int32(430450), v20+int32(512))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1449)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L4
	} else {
		goto L520
	}
L520:
	;
	F_errfinish(m, int32(483587), int32(724), int32(135850))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L4
	} else {
		goto L521
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	goto L31
L523:
	;
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v1527 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L524:
	;
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v1521 != 0 {
		goto L538
	} else {
		goto L539
	}
L525:
	;
	if v1489 != 0 {
		goto L20
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v1517 = v27 + int32(32)
	if v1489 != 0 {
		v1524 = v1489
		v1525 = v1517
		goto L523
	} else {
		goto L537
	}
L528:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v1491 != 0 {
		goto L19
	} else {
		goto L529
	}
L529:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v1492 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v1518 = v27 + int32(32)
	goto L524
L531:
	;
	goto L532
L532:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L4
	} else {
		goto L533
	}
L533:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L4
	} else {
		goto L534
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(511294)
	F_errmsg(m, int32(406175), v20+int32(272))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L4
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(483587), int32(745), int32(135850))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L4
	} else {
		goto L536
	}
L536:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L537:
	;
	v1518 = v1517
	goto L524
L538:
	;
	v1522 = int32(644123)
	goto L540
L539:
	;
	v1522 = int32(731166)
	goto L540
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1522
	v1524 = v1522
	v1525 = v1518
	goto L523
L541:
	;
	if v1524&int32(3) == int32(0) {
		v1645 = v1524
		goto L577
	} else {
		goto L578
	}
L542:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v1608 == int32(0) {
		goto L571
	} else {
		goto L572
	}
L543:
	;
	v1530 = int32(0)
	v1534 = v1526 & int32(1)
	if v1534 != 0 {
		goto L546
	} else {
		goto L547
	}
L544:
	;
	goto L545
L545:
	;
	if v1527&int32(3) == int32(0) {
		v1564 = v1527
		goto L555
	} else {
		goto L556
	}
L546:
	;
	v1535 = v1530
	goto L548
L547:
	;
	v1535 = int32(2)
	goto L548
L548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1535
	if v1534 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v1539 = int32(731167)
	goto L551
L550:
	;
	v1539 = int32(518390)
	goto L551
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v1539
	if v1534 != 0 {
		v1605 = v1539
		v1607 = v1535
		goto L542
	} else {
		goto L552
	}
L552:
	;
	v1618 = v1530
	v1619 = v1539
	v1621 = v1535
	goto L541
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1597
	v1599 = int32(0)
	if v1526&int32(1) == v1599 {
		v1618 = v1599
		v1619 = v1527
		v1621 = v1597
		goto L541
	} else {
		goto L570
	}
L554:
	;
	v1597 = v1589 - v1527
	goto L553
L555:
	;
	v1568 = v1564
	goto L564
L556:
	;
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1527))))
	if v1548 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v1597 = int32(0)
	goto L553
L558:
	;
	goto L559
L559:
	;
	v1553 = v1527
	goto L560
L560:
	;
	v1557 = v1553 + int32(1)
	if v1557&int32(3) == int32(0) {
		v1564 = v1557
		goto L555
	} else {
		goto L562
	}
L561:
	;
	v1589 = v1557
	goto L554
L562:
	;
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557))))
	if v1562 != 0 {
		v1553 = v1557
		goto L560
	} else {
		goto L563
	}
L563:
	;
	goto L561
L564:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1568)))
	v1577 = int32(-2139062144)
	if (int32(16843008)-v1574|v1574)&v1577 == v1577 {
		v1568 = v1568 + int32(4)
		goto L564
	} else {
		goto L566
	}
L565:
	;
	v1583 = v1568
	goto L567
L566:
	;
	goto L565
L567:
	;
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583))))
	if v1587 != 0 {
		v1583 = v1583 + int32(1)
		goto L567
	} else {
		goto L569
	}
L568:
	;
	v1589 = v1583
	goto L554
L569:
	;
	goto L568
L570:
	;
	v1605 = v1527
	v1607 = v1597
	goto L542
L571:
	;
	v1611 = int32(704724)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v1611
	v1614 = v1611
	goto L573
L572:
	;
	v1614 = v1608
	goto L573
L573:
	;
	v1615 = int32(1)
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1616 != 0 {
		v1618 = v1615
		v1619 = v1605
		v1621 = v1607
		goto L541
	} else {
		goto L574
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1614
	v1618 = v1615
	v1619 = v1605
	v1621 = v1607
	goto L541
L575:
	;
	if v1678 != int32(1) {
		goto L18
	} else {
		goto L592
	}
L576:
	;
	v1678 = v1670 - v1524
	goto L575
L577:
	;
	v1649 = v1645
	goto L586
L578:
	;
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524))))
	if v1629 == int32(0) {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v1678 = int32(0)
	goto L575
L580:
	;
	goto L581
L581:
	;
	v1634 = v1524
	goto L582
L582:
	;
	v1638 = v1634 + int32(1)
	if v1638&int32(3) == int32(0) {
		v1645 = v1638
		goto L577
	} else {
		goto L584
	}
L583:
	;
	v1670 = v1638
	goto L576
L584:
	;
	v1643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638))))
	if v1643 != 0 {
		v1634 = v1638
		goto L582
	} else {
		goto L585
	}
L585:
	;
	goto L583
L586:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1649)))
	v1658 = int32(-2139062144)
	if (int32(16843008)-v1655|v1655)&v1658 == v1658 {
		v1649 = v1649 + int32(4)
		goto L586
	} else {
		goto L588
	}
L587:
	;
	v1664 = v1649
	goto L589
L588:
	;
	goto L587
L589:
	;
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664))))
	if v1668 != 0 {
		v1664 = v1664 + int32(1)
		goto L589
	} else {
		goto L591
	}
L590:
	;
	v1670 = v1664
	goto L576
L591:
	;
	goto L590
L592:
	;
	v1681 = int32(13)
	v1682 = F___strchrnul(m, v1524, v1681)
	mBase = m.M
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1682))))
	if v1684 == v1681 {
		goto L594
	} else {
		goto L595
	}
L593:
	;
	if v1688 != 0 {
		goto L17
	} else {
		goto L597
	}
L594:
	;
	v1688 = v1682
	goto L596
L595:
	;
	v1688 = int32(0)
	goto L596
L596:
	;
	goto L593
L597:
	;
	v1689 = int32(10)
	v1690 = F___strchrnul(m, v1524, v1689)
	mBase = m.M
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1690))))
	if v1692 == v1689 {
		goto L599
	} else {
		goto L600
	}
L598:
	;
	if v1696 != 0 {
		goto L17
	} else {
		goto L602
	}
L599:
	;
	v1696 = v1690
	goto L601
L600:
	;
	v1696 = int32(0)
	goto L601
L601:
	;
	goto L598
L602:
	;
	v1697 = int32(13)
	v1698 = F___strchrnul(m, v1619, v1697)
	mBase = m.M
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1698))))
	if v1700 == v1697 {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	if v1704 != 0 {
		goto L16
	} else {
		goto L607
	}
L604:
	;
	v1704 = v1698
	goto L606
L605:
	;
	v1704 = int32(0)
	goto L606
L606:
	;
	goto L603
L607:
	;
	v1705 = int32(10)
	v1706 = F___strchrnul(m, v1619, v1705)
	mBase = m.M
	v1708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1706))))
	if v1708 == v1705 {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	if v1712 != 0 {
		goto L16
	} else {
		goto L612
	}
L609:
	;
	v1712 = v1706
	goto L611
L610:
	;
	v1712 = int32(0)
	goto L611
L611:
	;
	goto L608
L612:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v1713 != 0 {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	if v1713&int32(3) == int32(0) {
		v1737 = v1713
		goto L618
	} else {
		goto L619
	}
L614:
	;
	goto L615
L615:
	;
	if v1618 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v1770
	v1772 = int32(13)
	v1773 = F___strchrnul(m, v1713, v1772)
	mBase = m.M
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1773))))
	if v1775 == v1772 {
		goto L634
	} else {
		goto L635
	}
L617:
	;
	v1770 = v1762 - v1713
	goto L616
L618:
	;
	v1741 = v1737
	goto L627
L619:
	;
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713))))
	if v1721 == int32(0) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v1770 = int32(0)
	goto L616
L621:
	;
	goto L622
L622:
	;
	v1726 = v1713
	goto L623
L623:
	;
	v1730 = v1726 + int32(1)
	if v1730&int32(3) == int32(0) {
		v1737 = v1730
		goto L618
	} else {
		goto L625
	}
L624:
	;
	v1762 = v1730
	goto L617
L625:
	;
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1730))))
	if v1735 != 0 {
		v1726 = v1730
		goto L623
	} else {
		goto L626
	}
L626:
	;
	goto L624
L627:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1741)))
	v1750 = int32(-2139062144)
	if (int32(16843008)-v1747|v1747)&v1750 == v1750 {
		v1741 = v1741 + int32(4)
		goto L627
	} else {
		goto L629
	}
L628:
	;
	v1756 = v1741
	goto L630
L629:
	;
	goto L628
L630:
	;
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1756))))
	if v1760 != 0 {
		v1756 = v1756 + int32(1)
		goto L630
	} else {
		goto L632
	}
L631:
	;
	v1762 = v1756
	goto L617
L632:
	;
	goto L631
L633:
	;
	if v1779 != 0 {
		goto L15
	} else {
		goto L637
	}
L634:
	;
	v1779 = v1773
	goto L636
L635:
	;
	v1779 = int32(0)
	goto L636
L636:
	;
	goto L633
L637:
	;
	v1780 = int32(10)
	v1781 = F___strchrnul(m, v1713, v1780)
	mBase = m.M
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1781))))
	if v1783 == v1780 {
		goto L639
	} else {
		goto L640
	}
L638:
	;
	if v1787 != 0 {
		goto L15
	} else {
		goto L642
	}
L639:
	;
	v1787 = v1781
	goto L641
L640:
	;
	v1787 = int32(0)
	goto L641
L641:
	;
	goto L638
L642:
	;
	goto L615
L643:
	;
	v1791 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1524))))
	goto L650
L644:
	;
	goto L645
L645:
	;
	if v1490 != 0 {
		goto L673
	} else {
		goto L674
	}
L646:
	;
	if v1895 != 0 {
		goto L14
	} else {
		goto L672
	}
L647:
	;
	v1895 = int32(0)
	goto L646
L648:
	;
	v1873 = v1866
	v1875 = v1868
	goto L666
L649:
	;
	if base.B2i32(v1813 != v1814) == int32(0) {
		goto L647
	} else {
		goto L657
	}
L650:
	;
	goto L651
L651:
	;
	v1805 = int32(537974)
	v1807 = int32(39)
	goto L652
L652:
	;
	v1810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1805))))
	if v1810 == v1791&int32(255) {
		v1866 = v1805
		v1868 = v1807
		goto L648
	} else {
		goto L654
	}
L653:
	;
	goto L649
L654:
	;
	v1812 = int32(1)
	v1813 = v1807 - v1812
	v1814 = int32(0)
	v1817 = v1805 + v1812
	if v1817&int32(3) == v1814 {
		goto L649
	} else {
		goto L655
	}
L655:
	;
	if v1813 != 0 {
		v1805 = v1817
		v1807 = v1813
		goto L652
	} else {
		goto L656
	}
L656:
	;
	goto L653
L657:
	;
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1817))))
	if v1829 == v1791&int32(255) {
		v1859 = v1817
		v1861 = v1813
		goto L658
	} else {
		goto L659
	}
L658:
	;
	if v1861 == int32(0) {
		goto L647
	} else {
		goto L665
	}
L659:
	;
	if base.Ui32(v1813) < base.Ui32(int32(4)) {
		v1859 = v1817
		v1861 = v1813
		goto L658
	} else {
		goto L660
	}
L660:
	;
	v1839 = v1817
	v1841 = v1813
	goto L661
L661:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1839)))
	v1846 = v1845 ^ v1791&int32(255)*int32(16843009)
	v1849 = int32(-2139062144)
	if (int32(16843008)-v1846|v1846)&v1849 != v1849 {
		v1866 = v1839
		v1868 = v1841
		goto L648
	} else {
		goto L663
	}
L662:
	;
	v1859 = v1854
	v1861 = v1856
	goto L658
L663:
	;
	v1853 = int32(4)
	v1854 = v1839 + v1853
	v1856 = v1841 - v1853
	if base.Ui32(int32(3)) < base.Ui32(v1856) {
		v1839 = v1854
		v1841 = v1856
		goto L661
	} else {
		goto L664
	}
L664:
	;
	goto L662
L665:
	;
	v1866 = v1859
	v1868 = v1861
	goto L648
L666:
	;
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1873))))
	if v1791&int32(255) == v1878 {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	goto L647
L668:
	;
	v1895 = v1873
	goto L646
L669:
	;
	goto L670
L670:
	;
	v1880 = int32(1)
	v1883 = v1875 - v1880
	if v1883 != 0 {
		v1873 = v1873 + v1880
		v1875 = v1883
		goto L666
	} else {
		goto L671
	}
L671:
	;
	goto L667
L672:
	;
	goto L645
L673:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v1896 != 0 {
		goto L13
	} else {
		goto L676
	}
L674:
	;
	goto L675
L675:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v1618 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L676:
	;
	goto L675
L677:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L790
	}
L678:
	;
	if l2 == int32(0) {
		goto L778
	} else {
		goto L779
	}
L679:
	;
	if l2 == int32(0) {
		goto L773
	} else {
		goto L774
	}
L680:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L771
	}
L681:
	;
	if l2 != 0 {
		goto L759
	} else {
		goto L760
	}
L682:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v2122 != 0 {
		goto L680
	} else {
		goto L752
	}
L683:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v2097 == int32(0) {
		goto L744
	} else {
		goto L745
	}
L684:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v2021&int32(3) == int32(0) {
		v2045 = v2021
		goto L724
	} else {
		goto L725
	}
L685:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1999 == int32(0) {
		goto L683
	} else {
		goto L717
	}
L686:
	;
	if v1897 == int32(0) {
		goto L685
	} else {
		goto L689
	}
L687:
	;
	goto L688
L688:
	;
	if v1897&int32(3) == int32(0) {
		v1944 = v1897
		goto L696
	} else {
		goto L697
	}
L689:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L4
	} else {
		goto L690
	}
L690:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L4
	} else {
		goto L691
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = int32(528430)
	F_errmsg(m, int32(406253), v20+int32(224))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L4
	} else {
		goto L692
	}
L692:
	;
	F_errfinish(m, int32(483587), int32(822), int32(135850))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L4
	} else {
		goto L693
	}
L693:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L694:
	;
	if v1977 != int32(1) {
		goto L12
	} else {
		goto L711
	}
L695:
	;
	v1977 = v1969 - v1897
	goto L694
L696:
	;
	v1948 = v1944
	goto L705
L697:
	;
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897))))
	if v1928 == int32(0) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v1977 = int32(0)
	goto L694
L699:
	;
	goto L700
L700:
	;
	v1933 = v1897
	goto L701
L701:
	;
	v1937 = v1933 + int32(1)
	if v1937&int32(3) == int32(0) {
		v1944 = v1937
		goto L696
	} else {
		goto L703
	}
L702:
	;
	v1969 = v1937
	goto L695
L703:
	;
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1937))))
	if v1942 != 0 {
		v1933 = v1937
		goto L701
	} else {
		goto L704
	}
L704:
	;
	goto L702
L705:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1948)))
	v1957 = int32(-2139062144)
	if (int32(16843008)-v1954|v1954)&v1957 == v1957 {
		v1948 = v1948 + int32(4)
		goto L705
	} else {
		goto L707
	}
L706:
	;
	v1963 = v1948
	goto L708
L707:
	;
	goto L706
L708:
	;
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1963))))
	if v1967 != 0 {
		v1963 = v1963 + int32(1)
		goto L708
	} else {
		goto L710
	}
L709:
	;
	v1969 = v1963
	goto L695
L710:
	;
	goto L709
L711:
	;
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524))))
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897))))
	if v1980 != v1981 {
		goto L684
	} else {
		goto L712
	}
L712:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L4
	} else {
		goto L713
	}
L713:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L4
	} else {
		goto L714
	}
L714:
	;
	F_errmsg(m, int32(91959), int32(0))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L4
	} else {
		goto L715
	}
L715:
	;
	F_errfinish(m, int32(483587), int32(832), int32(135850))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L4
	} else {
		goto L716
	}
L716:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L717:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L4
	} else {
		goto L718
	}
L718:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L4
	} else {
		goto L719
	}
L719:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = int32(530189)
	F_errmsg(m, int32(406253), v20+int32(208))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L4
	} else {
		goto L720
	}
L720:
	;
	F_errfinish(m, int32(483587), int32(839), int32(135850))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L4
	} else {
		goto L721
	}
L721:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L722:
	;
	if v2078 == int32(1) {
		goto L682
	} else {
		goto L739
	}
L723:
	;
	v2078 = v2070 - v2021
	goto L722
L724:
	;
	v2049 = v2045
	goto L733
L725:
	;
	v2029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2021))))
	if v2029 == int32(0) {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	v2078 = int32(0)
	goto L722
L727:
	;
	goto L728
L728:
	;
	v2034 = v2021
	goto L729
L729:
	;
	v2038 = v2034 + int32(1)
	if v2038&int32(3) == int32(0) {
		v2045 = v2038
		goto L724
	} else {
		goto L731
	}
L730:
	;
	v2070 = v2038
	goto L723
L731:
	;
	v2043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2038))))
	if v2043 != 0 {
		v2034 = v2038
		goto L729
	} else {
		goto L732
	}
L732:
	;
	goto L730
L733:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2049)))
	v2058 = int32(-2139062144)
	if (int32(16843008)-v2055|v2055)&v2058 == v2058 {
		v2049 = v2049 + int32(4)
		goto L733
	} else {
		goto L735
	}
L734:
	;
	v2064 = v2049
	goto L736
L735:
	;
	goto L734
L736:
	;
	v2068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2064))))
	if v2068 != 0 {
		v2064 = v2064 + int32(1)
		goto L736
	} else {
		goto L738
	}
L737:
	;
	v2070 = v2064
	goto L723
L738:
	;
	goto L737
L739:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L4
	} else {
		goto L740
	}
L740:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L4
	} else {
		goto L741
	}
L741:
	;
	F_errmsg(m, int32(213791), int32(0))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L4
	} else {
		goto L742
	}
L742:
	;
	F_errfinish(m, int32(483587), int32(844), int32(135850))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L4
	} else {
		goto L743
	}
L743:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L744:
	;
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v2100 != int32(1) {
		goto L681
	} else {
		goto L747
	}
L745:
	;
	goto L746
L746:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L4
	} else {
		goto L748
	}
L747:
	;
	goto L746
L748:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L4
	} else {
		goto L749
	}
L749:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = int32(528424)
	F_errmsg(m, int32(406253), v20+int32(160))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L4
	} else {
		goto L750
	}
L750:
	;
	F_errfinish(m, int32(483587), int32(851), int32(135850))
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L4
	} else {
		goto L751
	}
L751:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L752:
	;
	if l2 != 0 {
		goto L753
	} else {
		goto L754
	}
L753:
	;
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v2123&int32(1) != 0 {
		goto L11
	} else {
		goto L756
	}
L754:
	;
	goto L755
L755:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v2126 == int32(0) {
		goto L679
	} else {
		goto L757
	}
L756:
	;
	goto L755
L757:
	;
	if l2 != 0 {
		v2506 = v1897
		goto L8
	} else {
		goto L758
	}
L758:
	;
	goto L7
L759:
	;
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v2129&int32(1) != 0 {
		goto L11
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v2132 == int32(0) {
		goto L763
	} else {
		goto L764
	}
L762:
	;
	goto L761
L763:
	;
	v2135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v2135 != int32(1) {
		goto L678
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L4
	} else {
		goto L767
	}
L766:
	;
	goto L765
L767:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L4
	} else {
		goto L768
	}
L768:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = int32(522856)
	F_errmsg(m, int32(406253), v20+int32(176))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L4
	} else {
		goto L769
	}
L769:
	;
	F_errfinish(m, int32(483587), int32(866), int32(135850))
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L4
	} else {
		goto L770
	}
L770:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L771:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v2157 != 0 {
		goto L7
	} else {
		goto L772
	}
L772:
	;
	goto L679
L773:
	;
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v2160&int32(1) != 0 {
		goto L7
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v2163 != 0 {
		goto L677
	} else {
		goto L777
	}
L776:
	;
	goto L775
L777:
	;
	v2478 = v1897
	goto L10
L778:
	;
	v2166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v2166&int32(1) != 0 {
		goto L7
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v2169 == int32(0) {
		goto L782
	} else {
		goto L783
	}
L781:
	;
	goto L780
L782:
	;
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v2173 != int32(1) {
		v2478 = int32(0)
		goto L10
	} else {
		goto L785
	}
L783:
	;
	goto L784
L784:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L4
	} else {
		goto L786
	}
L785:
	;
	goto L784
L786:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L4
	} else {
		goto L787
	}
L787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = int32(522871)
	F_errmsg(m, int32(406253), v20+int32(192))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L4
	} else {
		goto L788
	}
L788:
	;
	F_errfinish(m, int32(483587), int32(882), int32(135850))
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L4
	} else {
		goto L789
	}
L789:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L790:
	;
	v2506 = v1897
	goto L8
L791:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L792:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L4
	} else {
		goto L793
	}
L793:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+352)) = v373
	F_errmsg(m, int32(517557), v20+int32(352))
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L4
	} else {
		goto L794
	}
L794:
	;
	F_errfinish(m, int32(483587), int32(415), int32(410834))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L4
	} else {
		goto L795
	}
L795:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L796:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L4
	} else {
		goto L797
	}
L797:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+468)) = int32(517626)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+464)) = int32(515668)
	F_errmsg(m, int32(183288), v20+int32(464))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L4
	} else {
		goto L798
	}
L798:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v2235)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L4
	} else {
		goto L799
	}
L799:
	;
	F_errfinish(m, int32(483587), int32(442), int32(410810))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L4
	} else {
		goto L800
	}
L800:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L801:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L4
	} else {
		goto L802
	}
L802:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+452)) = v1103
	*(*int32)(unsafe.Add(mBase, uint32(v20)+448)) = int32(515668)
	F_errmsg(m, int32(430266), v20+int32(448))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L4
	} else {
		goto L803
	}
L803:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v2258)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L4
	} else {
		goto L804
	}
L804:
	;
	F_errfinish(m, int32(483587), int32(456), int32(410810))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L4
	} else {
		goto L805
	}
L805:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L806:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L4
	} else {
		goto L807
	}
L807:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+480)) = v2273
	F_errmsg(m, int32(340521), v20+int32(480))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L4
	} else {
		goto L808
	}
L808:
	;
	F_errfinish(m, int32(483587), int32(476), int32(243713))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L4
	} else {
		goto L809
	}
L809:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L810:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L4
	} else {
		goto L811
	}
L811:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+496)) = v1430
	F_errmsg(m, int32(236023), v20+int32(496))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L4
	} else {
		goto L812
	}
L812:
	;
	F_errfinish(m, int32(483587), int32(486), int32(243713))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L4
	} else {
		goto L813
	}
L813:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L814:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L4
	} else {
		goto L815
	}
L815:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+304)) = int32(516038)
	F_errmsg(m, int32(406175), v20+int32(304))
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L4
	} else {
		goto L816
	}
L816:
	;
	F_errfinish(m, int32(483587), int32(735), int32(135850))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L4
	} else {
		goto L817
	}
L817:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L818:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L4
	} else {
		goto L819
	}
L819:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = int32(523681)
	F_errmsg(m, int32(406175), v20+int32(288))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L4
	} else {
		goto L820
	}
L820:
	;
	F_errfinish(m, int32(483587), int32(740), int32(135850))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L4
	} else {
		goto L821
	}
L821:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L822:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L4
	} else {
		goto L823
	}
L823:
	;
	F_errmsg(m, int32(213693), int32(0))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L4
	} else {
		goto L824
	}
L824:
	;
	F_errfinish(m, int32(483587), int32(767), int32(135850))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L4
	} else {
		goto L825
	}
L825:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L826:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L4
	} else {
		goto L827
	}
L827:
	;
	F_errmsg(m, int32(240913), int32(0))
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L4
	} else {
		goto L828
	}
L828:
	;
	F_errfinish(m, int32(483587), int32(774), int32(135850))
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L4
	} else {
		goto L829
	}
L829:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L830:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L4
	} else {
		goto L831
	}
L831:
	;
	F_errmsg(m, int32(240850), int32(0))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L4
	} else {
		goto L832
	}
L832:
	;
	F_errfinish(m, int32(483587), int32(780), int32(135850))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L4
	} else {
		goto L833
	}
L833:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L834:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L4
	} else {
		goto L835
	}
L835:
	;
	F_errmsg(m, int32(240784), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L4
	} else {
		goto L836
	}
L836:
	;
	F_errfinish(m, int32(483587), int32(790), int32(135850))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L4
	} else {
		goto L837
	}
L837:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L838:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L4
	} else {
		goto L839
	}
L839:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v1525)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v2412
	F_errmsg(m, int32(696161), v20+int32(256))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L4
	} else {
		goto L840
	}
L840:
	;
	F_errfinish(m, int32(483587), int32(808), int32(135850))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L4
	} else {
		goto L841
	}
L841:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L842:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L4
	} else {
		goto L843
	}
L843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = int32(516624)
	F_errmsg(m, int32(406175), v20+int32(240))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L4
	} else {
		goto L844
	}
L844:
	;
	F_errfinish(m, int32(483587), int32(815), int32(135850))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L4
	} else {
		goto L845
	}
L845:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L846:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L4
	} else {
		goto L847
	}
L847:
	;
	F_errmsg(m, int32(213744), int32(0))
	mBase = m.M
	v2453 = m.ExcPending
	if v2453 != 0 {
		goto L4
	} else {
		goto L848
	}
L848:
	;
	F_errfinish(m, int32(483587), int32(827), int32(135850))
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L4
	} else {
		goto L849
	}
L849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L850:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L4
	} else {
		goto L851
	}
L851:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(522205)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(528424)
	F_errmsg(m, int32(183288), v20)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L4
	} else {
		goto L852
	}
L852:
	;
	F_errfinish(m, int32(483587), int32(858), int32(135850))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L4
	} else {
		goto L853
	}
L853:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L854:
	;
	v2479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v2479&int32(1) == int32(0) {
		v2506 = v2478
		goto L8
	} else {
		goto L855
	}
L855:
	;
	goto L9
L856:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L4
	} else {
		goto L857
	}
L857:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = int32(517626)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = int32(522871)
	F_errmsg(m, int32(183288), v20+int32(128))
	mBase = m.M
	v2500 = m.ExcPending
	if v2500 != 0 {
		goto L4
	} else {
		goto L858
	}
L858:
	;
	F_errfinish(m, int32(483587), int32(891), int32(135850))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L4
	} else {
		goto L859
	}
L859:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L860:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L4
	} else {
		goto L959
	}
L861:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L4
	} else {
		goto L955
	}
L862:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L4
	} else {
		goto L951
	}
L863:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		goto L4
	} else {
		goto L947
	}
L864:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L4
	} else {
		goto L943
	}
L865:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L4
	} else {
		goto L939
	}
L866:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L4
	} else {
		goto L935
	}
L867:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L4
	} else {
		goto L931
	}
L868:
	;
	if v2514 == int32(0) {
		goto L872
	} else {
		goto L873
	}
L869:
	;
	v2514 = v2508
	goto L871
L870:
	;
	v2514 = int32(0)
	goto L871
L871:
	;
	goto L868
L872:
	;
	if v1618 != 0 {
		goto L875
	} else {
		goto L876
	}
L873:
	;
	goto L874
L874:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L4
	} else {
		goto L927
	}
L875:
	;
	v2517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2506))))
	v2518 = F___strchrnul(m, v1619, v2517)
	mBase = m.M
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2518))))
	if v2520 == v2517&int32(255) {
		goto L879
	} else {
		goto L880
	}
L876:
	;
	goto L877
L877:
	;
	if l2 == int32(0) {
		goto L883
	} else {
		goto L884
	}
L878:
	;
	if v2524 != 0 {
		goto L867
	} else {
		goto L882
	}
L879:
	;
	v2524 = v2518
	goto L881
L880:
	;
	v2524 = int32(0)
	goto L881
L881:
	;
	goto L878
L882:
	;
	goto L877
L883:
	;
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)))
	if v2527&int32(1) != 0 {
		goto L866
	} else {
		goto L886
	}
L884:
	;
	goto L885
L885:
	;
	if v1713 == int32(0) {
		goto L887
	} else {
		goto L888
	}
L886:
	;
	goto L885
L887:
	;
	if v1490 != 0 {
		goto L919
	} else {
		goto L920
	}
L888:
	;
	if l2 == int32(0) {
		goto L865
	} else {
		goto L889
	}
L889:
	;
	v2534 = F___strchrnul(m, v1713, v2507)
	mBase = m.M
	v2536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2534))))
	if v2536 == v2507&int32(255) {
		goto L891
	} else {
		goto L892
	}
L890:
	;
	if v2540 != 0 {
		goto L864
	} else {
		goto L894
	}
L891:
	;
	v2540 = v2534
	goto L893
L892:
	;
	v2540 = int32(0)
	goto L893
L893:
	;
	goto L890
L894:
	;
	if v1618 != 0 {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	v2541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2506))))
	v2542 = F___strchrnul(m, v1713, v2541)
	mBase = m.M
	v2544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2542))))
	if v2544 == v2541&int32(255) {
		goto L899
	} else {
		goto L900
	}
L896:
	;
	goto L897
L897:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v1621 != v2549 {
		goto L887
	} else {
		goto L903
	}
L898:
	;
	if v2548 != 0 {
		goto L863
	} else {
		goto L902
	}
L899:
	;
	v2548 = v2542
	goto L901
L900:
	;
	v2548 = int32(0)
	goto L901
L901:
	;
	goto L898
L902:
	;
	goto L897
L903:
	;
	if v1621 == int32(0) {
		goto L905
	} else {
		goto L906
	}
L904:
	;
	if v2594 == int32(0) {
		goto L862
	} else {
		goto L918
	}
L905:
	;
	v2594 = int32(0)
	goto L904
L906:
	;
	goto L907
L907:
	;
	v2556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619))))
	if v2556 != 0 {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v2557 = v1619
	v2558 = v1713
	v2559 = v1621
	v2560 = v2556
	goto L912
L909:
	;
	v2582 = v1713
	v2586 = int32(0)
	goto L910
L910:
	;
	v2587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2582))))
	v2594 = v2586 - v2587
	goto L904
L911:
	;
	v2582 = v2577
	v2586 = v2579
	goto L910
L912:
	;
	v2562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2558))))
	if v2560 != v2562 {
		v2577 = v2558
		v2579 = v2560
		goto L911
	} else {
		goto L914
	}
L913:
	;
	v2577 = v2571
	v2579 = int32(0)
	goto L911
L914:
	;
	if v2562 == int32(0) {
		v2577 = v2558
		v2579 = v2560
		goto L911
	} else {
		goto L915
	}
L915:
	;
	v2567 = v2559 - int32(1)
	if v2567 == int32(0) {
		v2577 = v2558
		v2579 = v2560
		goto L911
	} else {
		goto L916
	}
L916:
	;
	v2570 = int32(1)
	v2571 = v2558 + v2570
	v2572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2557)+1)))
	if v2572 != 0 {
		v2557 = v2557 + v2570
		v2558 = v2571
		v2559 = v2567
		v2560 = v2572
		goto L912
	} else {
		goto L917
	}
L917:
	;
	goto L913
L918:
	;
	goto L887
L919:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v2597 != 0 {
		goto L861
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v2598 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	if v2598 != int64(0) {
		goto L923
	} else {
		goto L924
	}
L922:
	;
	goto L921
L923:
	;
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v2601 == int32(0) {
		goto L860
	} else {
		goto L926
	}
L924:
	;
	goto L925
L925:
	;
	m.G0 = v20 + int32(528)
	return
L926:
	;
	goto L925
L927:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2613 = m.ExcPending
	if v2613 != 0 {
		goto L4
	} else {
		goto L928
	}
L928:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = int32(523681)
	F_errmsg(m, int32(263307), v20+int32(112))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L4
	} else {
		goto L929
	}
L929:
	;
	F_errfinish(m, int32(483587), int32(899), int32(135850))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L4
	} else {
		goto L930
	}
L930:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L931:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L4
	} else {
		goto L932
	}
L932:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(523681)
	F_errmsg(m, int32(263372), v20+int32(96))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L4
	} else {
		goto L933
	}
L933:
	;
	F_errfinish(m, int32(483587), int32(908), int32(135850))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L4
	} else {
		goto L934
	}
L934:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L935:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L4
	} else {
		goto L936
	}
L936:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(517626)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(527918)
	F_errmsg(m, int32(183288), v20+int32(80))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L4
	} else {
		goto L937
	}
L937:
	;
	F_errfinish(m, int32(483587), int32(917), int32(135850))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L4
	} else {
		goto L938
	}
L938:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L939:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L4
	} else {
		goto L940
	}
L940:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(517626)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(511294)
	F_errmsg(m, int32(183288), v20-int32(-64))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L4
	} else {
		goto L941
	}
L941:
	;
	F_errfinish(m, int32(483587), int32(927), int32(135850))
	mBase = m.M
	v2686 = m.ExcPending
	if v2686 != 0 {
		goto L4
	} else {
		goto L942
	}
L942:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L943:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L4
	} else {
		goto L944
	}
L944:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(511294)
	F_errmsg(m, int32(263307), v20+int32(48))
	mBase = m.M
	v2700 = m.ExcPending
	if v2700 != 0 {
		goto L4
	} else {
		goto L945
	}
L945:
	;
	F_errfinish(m, int32(483587), int32(935), int32(135850))
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L4
	} else {
		goto L946
	}
L946:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L947:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L4
	} else {
		goto L948
	}
L948:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(511294)
	F_errmsg(m, int32(263372), v20+int32(32))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L4
	} else {
		goto L949
	}
L949:
	;
	F_errfinish(m, int32(483587), int32(944), int32(135850))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L4
	} else {
		goto L950
	}
L950:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L951:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L4
	} else {
		goto L952
	}
L952:
	;
	F_errmsg(m, int32(370700), int32(0))
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L4
	} else {
		goto L953
	}
L953:
	;
	F_errfinish(m, int32(483587), int32(952), int32(135850))
	mBase = m.M
	v2740 = m.ExcPending
	if v2740 != 0 {
		goto L4
	} else {
		goto L954
	}
L954:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L955:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L4
	} else {
		goto L956
	}
L956:
	;
	F_errmsg(m, int32(406208), int32(0))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L4
	} else {
		goto L957
	}
L957:
	;
	F_errfinish(m, int32(483587), int32(958), int32(135850))
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L4
	} else {
		goto L958
	}
L958:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L959:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L4
	} else {
		goto L960
	}
L960:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(529968)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(515668)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(511652)
	F_errmsg(m, int32(179620), v20+int32(16))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L4
	} else {
		goto L961
	}
L961:
	;
	F_errfinish(m, int32(483587), int32(966), int32(135850))
	mBase = m.M
	v2779 = m.ExcPending
	if v2779 != 0 {
		goto L4
	} else {
		goto L962
	}
L962:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L963:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L4
	} else {
		goto L964
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+148)) = int32(517626)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = int32(522856)
	F_errmsg(m, int32(183288), v20+int32(144))
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L4
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(483587), int32(874), int32(135850))
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L4
	} else {
		goto L966
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L4
	} else {
		goto L968
	}
L968:
	;
	v2811 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+336)) = v2811
	F_errmsg(m, int32(702758), v20+int32(336))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L4
	} else {
		goto L969
	}
L969:
	;
	F_errfinish(m, int32(483587), int32(424), int32(410834))
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L4
	} else {
		goto L970
	}
L970:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessPendingWrites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	goto L2
L1:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L48
	}
L2:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	F_SetLatch(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L47
	}
L4:
	;
	return
L5:
	;
	v12 = *(*int64)(unsafe.Add(mBase, _consts[886]))
	if v12 <= int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v65 = m.T0[v64].(func(*base.Module) int32)(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	if v16 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _consts[888]))
	if base.I64_extend_i32_u(v16)*int64(1000)+v12 <= v20 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[889])))
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errmsg(m, int32(65182), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(486641), int32(2789), int32(66742))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _consts[888]))
	if v44 < base.I64_extend_i32_u(int32(base.Ui32(v16)>>(uint(int32(1))%32)))*int64(1000)+v12 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = m.T0[v58].(func(*base.Module) int32)(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = m.G0
	v71 = int32(16)
	v72 = v70 - v71
	m.G0 = v72
	F___gettimeofday(m, v72)
	mBase = m.M
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	v76 = int64(*(*int32)(unsafe.Add(mBase, uint32(v72)+8)))
	m.G0 = v72 + v71
	v84 = v76 + v75*int64(1000000) - int64(946684800000000)
	goto L25
L23:
	;
	goto L24
L24:
	;
	goto L3
L25:
	;
	v85 = int32(10000)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[887]))
	if v87 <= int32(0) {
		v122 = v85
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_WalSndWait(m, int32(6), v122, int32(100663304))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L34
	}
L27:
	;
	v91 = *(*int64)(unsafe.Add(mBase, _consts[886]))
	if v91 <= int64(0) {
		v122 = v85
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _consts[889])))
	v104 = base.I64_extend_i32_u(int32(base.Ui32(v87)>>(uint((v95^int32(-1))&int32(1))%32)))*int64(1000) + v91
	if v104 <= v84 {
		v121 = int32(0)
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v122 = v121
	goto L26
L30:
	;
	goto L29
L31:
	;
	v107 = int32(2147483647)
	v110 = v104 - v84
	if base.B2i32(int64(0) < v84)^base.B2i32(v110 < v104) != 0 {
		v121 = v107
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if int64(2147483646000) < v110 {
		v121 = v107
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v118 = base.I64_div_s(v110+int64(999), int64(1000))
	v121 = base.I32_wrap_i64(v118)
	goto L30
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = int32(0)
	goto L35
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v133 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v137 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v149 = m.T0[v148].(func(*base.Module) int32)(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v149 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L1
L47:
	;
	return
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PushCopiedSnapshot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	v10 = l0 + int32(24)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v13 = l0 + int32(16)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v16 = int32(2)
	v18 = int32(72)
	v23 = v14<<(uint(v16)%32) + v18
	if int32(0) < v11 {
		v26 = (v11+v14)<<(uint(v16)%32) + v18
	} else {
		v26 = v23
	}
	v27 = F_MemoryContextAlloc(m, v8, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v30 = v27 + int32(48)
		v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
		v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v33
		v35 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = v37
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v39
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v41
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v43
		v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v27))) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = int64(0)
		v49 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v30))) = v49
		*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v49
		v53 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v27)+30)) = uint8(v53)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v55 != 0 {
			v57 = v27 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v57
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v62 = v60 << (uint(int32(2)) % 32)
			if v62 != 0 {
				v63 = F__emscripten_memcpy_bulkmem(m, v57, v59, v62)
				mBase = m.M
			} else {
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
		}
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v68 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(0)
		} else {
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v71 == int32(1) {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				if v74 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(0)
				} else {
					v77 = v27 + v23
					*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v77
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v82 = v80 << (uint(int32(2)) % 32)
					if v82 != 0 {
						v83 = F__emscripten_memcpy_bulkmem(m, v77, v79, v82)
						mBase = m.M
					} else {
					}
				}
			} else {
				v77 = v27 + v23
				*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v77
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v82 = v80 << (uint(int32(2)) % 32)
				if v82 != 0 {
					v83 = F__emscripten_memcpy_bulkmem(m, v77, v79, v82)
					mBase = m.M
				} else {
				}
			}
		}
		v89 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
		F_PushActiveSnapshotWithLevel(m, v27, v90)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return
		} else {
			return
		}
	}
}
func F_p_isalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
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
	var v64 int32
	_ = v64
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
			if base.Ui32(int32(127)) < base.Ui32(v13) {
				v64 = int32(1)
				return v64
			} else {
				return base.B2i32(base.Ui32(v13-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v13|int32(32)-int32(97)) < base.Ui32(int32(26)))
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
			if base.Ui32(int32(10)) <= base.Ui32(v34-int32(48)) {
				v41 = F_iswalpha(m, v34)
				mBase = m.M
				v44 = base.B2i32(v41 != int32(0))
			} else {
				v44 = int32(1)
			}
			return v44
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v48))))
		v64 = base.B2i32(base.Ui32(v50-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v50|int32(32)-int32(97)) < base.Ui32(int32(26)))
		return v64
	}
}
func F_p_isspecial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12*int32(28))+uint32(_consts[1215])))
	v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v6+v8)
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
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(1)
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	if v28 != int32(6) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v31 != int32(1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v34
	goto L12
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = v35
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+v38<<(uint(int32(2))%32))))
	v45 = int32(1617664)
	v46 = int32(1618576)
	goto L13
L13:
	;
	v55 = v45 + (v46-v45)>>(uint(int32(3))%32)<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v42 == v56 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L6
L15:
	;
	return int32(1)
L16:
	;
	goto L17
L17:
	;
	v62 = base.B2i32(base.Ui32(v56) < base.Ui32(v42))
	if base.Ui32(v56) < base.Ui32(v42) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = v55 + int32(4)
	goto L20
L19:
	;
	v63 = v45
	goto L20
L20:
	;
	if base.Ui32(v56) < base.Ui32(v42) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v64 = v46
	goto L23
L22:
	;
	v64 = v55
	goto L23
L23:
	;
	if base.Ui32(v63) < base.Ui32(v64) {
		v45 = v63
		v46 = v64
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L14
}
func F_p_isurlchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != int32(1) {
		v24 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9))))
		if base.Ui32((v11-int32(127))&int32(255)) < base.Ui32(int32(162)) {
			v24 = v2
		} else {
			switch v11 - int32(60) {
			case 0, 2, 32, 34, 36, 63, 64, 65:
				v24 = v2
			case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62:
				v24 = int32(1)
			default:
				if v11 == int32(34) {
					v24 = v2
				} else {
					v24 = int32(1)
				}
			}
		}
	}
	return v24
}
func F_pairingheap_GISTSearchItem_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v4 < v11 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v106
L2:
	;
	v106 = int32(0) - v71
	goto L1
L3:
	;
	v16 = int32(32)
	v24 = v4
	v25 = v11
	goto L6
L4:
	;
	goto L5
L5:
	;
	v88 = int32(-1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v90 == v88 {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	v31 = v24 << (uint(int32(4)) % 32)
	v32 = l0 + v16 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+8)))
	if v33 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v76 = v24 + int32(1)
	if v76 < v74 {
		v24 = v76
		v25 = v74
		goto L6
	} else {
		goto L25
	}
L9:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+(l1+int32(40))))))
	if v37 != 0 {
		v74 = v25
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = v31 + (l1 + v16)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(-1)
L13:
	;
	return int32(1)
L14:
	;
	goto L15
L15:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
	v50 = int64(9223372036854775807)
	v51 = base.I64_reinterpret_f64(v44) & v50
	v54 = base.I64_reinterpret_f64(v45) & v50
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v54) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v71 != 0 {
		goto L2
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v71 = int32(0) - v63&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v54))|base.F64_lt(v44, v45))
	goto L17
L19:
	;
	v63 = base.B2i32(base.Ui64(v51) < base.Ui64(int64(9218868437227405313)))
	goto L18
L20:
	;
	goto L21
L21:
	;
	v59 = int32(1)
	if base.F64_gt(v44, v45) != 0 {
		v71 = v59
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v51) {
		v71 = v59
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v63 = v59
	goto L18
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v74 = v72
	goto L8
L25:
	;
	goto L7
L26:
	;
	return int32(0)
L27:
	;
	if v89 == int32(-1) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v89 == int32(-1) {
		v106 = v88
		goto L1
	} else {
		goto L31
	}
L30:
	;
	return int32(1)
L31:
	;
	goto L26
}
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, v8, l1, v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = base.B2i32(v11 < int32(0))
			if v11 < int32(0) {
				v15 = v8
			} else {
				v15 = l1
			}
			if v11 < int32(0) {
				v16 = l1
			} else {
				v16 = v8
			}
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v15
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v16
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15
			v23 = v16
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
			return
		}
	} else {
		v23 = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
		v28 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
		return
	}
}
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if base.Ui32(l1) <= base.Ui32(int32(8)) {
		v10 = F_MemoryContextAllocExtended(m, v6, l0, int32(4))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	} else {
		v17 = F_MemoryContextAllocExtended(m, v6, l0+l1, int32(4))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v24 = (v17 + l1 + int32(7)) & (int32(0) - l1)
			v26 = v24 - int32(8)
			*(*int64)(unsafe.Add(mBase, uint32(v26))) = base.I64_extend_i32_u(l1)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v26-v17)<<(uint(int64(34))%64) | int64(6)
			return v24
		}
	}
}
func F_parse_filename_for_nontemp_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v20-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v139 = v5
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
		v34 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(4294967295))
		mBase = m.M
		v35 = base.I32_wrap_i64(v34)
		v37 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		if v37 != 0 {
			v139 = v5
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if l0 == v38 {
				v139 = v5
			} else {
				if v35 == int32(0) {
					v139 = v5
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
					if v42 != int32(95) {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
						v93 = v42
						v94 = v38
						if v93&int32(255) == int32(46) {
							v100 = v94 + int32(1)
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
							if base.Ui32((v101-int32(58))&int32(255)) < base.Ui32(int32(247)) {
								v139 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
								v115 = F_strtox_2(m, v100, v12+int32(8), int32(10), int64(4294967295))
								mBase = m.M
								v116 = base.I32_wrap_i64(v115)
								v118 = *(*int32)(unsafe.Add(mBase, _consts[140]))
								if v118 != 0 {
									v139 = v5
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if v100 == v119 {
										v139 = v5
									} else {
										if v116 == int32(0) {
											v139 = v5
										} else {
											v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
											v126 = v116
											v127 = v123
											if v127&int32(255) != 0 {
												v139 = v5
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v126
												v139 = int32(1)
											}
										}
									}
								}
							}
						} else {
							v126 = v5
							v127 = v93
							if v127&int32(255) != 0 {
								v139 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v126
								v139 = int32(1)
							}
						}
					} else {
						v48 = v38 + int32(1)
						v50 = v12 + int32(12)
						v53 = int32(3)
						v56 = F_strncmp(m, int32(282818), v48, v53)
						mBase = m.M
						if v56 == int32(0) {
							v78 = v53
							v79 = int32(1)
							if v50 == int32(0) {
								v85 = v78
							} else {
								v82 = v78
								v83 = v79
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = v83
								v85 = v82
							}
						} else {
							v60 = int32(2)
							v64 = F_strncmp(m, int32(280832), v48, v60)
							mBase = m.M
							if v64 == int32(0) {
								v78 = v60
								v79 = v60
								if v50 == int32(0) {
									v85 = v78
								} else {
									v82 = v78
									v83 = v79
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v83
									v85 = v82
								}
							} else {
								v67 = int32(4)
								v70 = F_strncmp(m, int32(98702), v48, v67)
								mBase = m.M
								if v70 == int32(0) {
									if v50 != 0 {
										v82 = v67
										v83 = int32(3)
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = v83
										v85 = v82
									} else {
										v85 = v67
									}
								} else {
									v75 = int32(0)
									if v50 == v75 {
										v85 = v75
									} else {
										v82 = v75
										v83 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = v83
										v85 = v82
									}
								}
							}
						}
						if v85 <= int32(0) {
							v139 = v5
						} else {
							v91 = v85 + v38 + int32(1)
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
							v93 = v92
							v94 = v91
							if v93&int32(255) == int32(46) {
								v100 = v94 + int32(1)
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
								if base.Ui32((v101-int32(58))&int32(255)) < base.Ui32(int32(247)) {
									v139 = v5
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
									v115 = F_strtox_2(m, v100, v12+int32(8), int32(10), int64(4294967295))
									mBase = m.M
									v116 = base.I32_wrap_i64(v115)
									v118 = *(*int32)(unsafe.Add(mBase, _consts[140]))
									if v118 != 0 {
										v139 = v5
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										if v100 == v119 {
											v139 = v5
										} else {
											if v116 == int32(0) {
												v139 = v5
											} else {
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
												v126 = v116
												v127 = v123
												if v127&int32(255) != 0 {
													v139 = v5
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v126
													v139 = int32(1)
												}
											}
										}
									}
								}
							} else {
								v126 = v5
								v127 = v93
								if v127&int32(255) != 0 {
									v139 = v5
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v131
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v126
									v139 = int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	return v139
}
func F_parse_new_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
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
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(1)
	v16 = F_pullf_read_fixed(m, l0, v12, v10+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 < int32(0) {
			v104 = v16
			m.G0 = v10 + int32(16)
			return v104
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if base.Ui32(v22) < base.Ui32(int32(192)) {
				v98 = v22
				v100 = v12
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
				v104 = v100
				m.G0 = v10 + int32(16)
				return v104
			} else {
				if base.Ui32(v22) <= base.Ui32(int32(223)) {
					v30 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(14))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 < int32(0) {
							v104 = v30
							m.G0 = v10 + int32(16)
							return v104
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
							v87 = v34 | v22<<(uint(int32(8))%32) - int32(48960)
							v89 = v12
							if base.Ui32(v87) < base.Ui32(int32(16777217)) {
								v98 = v87
								v100 = v89
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
								v104 = v100
								m.G0 = v10 + int32(16)
								return v104
							} else {
								F_px_debug(m, int32(315553), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v104 = int32(-100)
									m.G0 = v10 + int32(16)
									return v104
								}
							}
						}
					}
				} else {
					if v22 == int32(255) {
						v45 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(13))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 < int32(0) {
								v104 = v45
								m.G0 = v10 + int32(16)
								return v104
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
								v53 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(12))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									if v53 < int32(0) {
										v104 = v53
										m.G0 = v10 + int32(16)
										return v104
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
										v61 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(11))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 < int32(0) {
												v104 = v61
												m.G0 = v10 + int32(16)
												return v104
											} else {
												v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
												v69 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(10))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													if v69 < int32(0) {
														v104 = v69
														m.G0 = v10 + int32(16)
														return v104
													} else {
														v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
														v74 = int32(8)
														v87 = v73 | (v57<<(uint(v74)%32)|v49<<(uint(int32(16))%32)|v65)<<(uint(v74)%32)
														v89 = int32(1)
														if base.Ui32(v87) < base.Ui32(int32(16777217)) {
															v98 = v87
															v100 = v89
															*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
															v104 = v100
															m.G0 = v10 + int32(16)
															return v104
														} else {
															F_px_debug(m, int32(315553), int32(0))
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																v104 = int32(-100)
																m.G0 = v10 + int32(16)
																return v104
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
					} else {
						v87 = int32(1) << (uint(v22) % 32)
						v89 = int32(2)
						if base.Ui32(v87) < base.Ui32(int32(16777217)) {
							v98 = v87
							v100 = v89
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
							v104 = v100
							m.G0 = v10 + int32(16)
							return v104
						} else {
							F_px_debug(m, int32(315553), int32(0))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v104 = int32(-100)
								m.G0 = v10 + int32(16)
								return v104
							}
						}
					}
				}
			}
		}
	}
}
func F_parser_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return
	} else {
		if l1 < int32(0) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v7 == int32(0) {
				return
			} else {
				v10 = F_pg_mbstrlen_with_len(m, v7, l1)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v14 = F_errposition(m, v10+int32(1))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v5
	v19 = F_lookup_ts_config_cache(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v22 = F_lookup_ts_parser_cache(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_FunctionCall2Coll(m, v22+int32(28), int32(0), l2, l3)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+44)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	goto L5
L5:
	;
	v57 = F_FunctionCall3Coll(m, v22+int32(56), int32(0), v27, v11+int32(-8), v11+int32(-4))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v301 = F_FunctionCall1Coll(m, v22+int32(84), int32(0), v27)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L64
	}
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v57 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if int32(0) < v57 {
		goto L5
	} else {
		goto L63
	}
L9:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v89 = F_palloc(m, int32(16))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	if v59 < int32(2047) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v66 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v66 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(431220), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2047)
	F_errdetail(m, int32(623311), v13)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(489727), int32(389), int32(61897))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v57
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v94 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v89
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v100
	v105 = F_LexizeExec(m, v11+int32(-56), v98)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v89
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v89
	goto L19
L23:
	;
	if v105 == int32(0) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v114 = v105
	goto L25
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v119 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v123 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L8
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v127 = v114
	v129 = v126
	v130 = v114 + int32(4)
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v114)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L60
	}
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v137 == v129 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v129 << (uint(int32(1)) % 32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v145 = F_repalloc(m, v142, v129<<(uint(int32(5))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	if v148&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v145
	goto L34
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v151 + int32(1)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v160&int32(3) == int32(0) {
		v184 = v160
		goto L41
	} else {
		goto L42
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v155+v156<<(uint(int32(4))%32))+2)) = uint16(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v221 = int32(4)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v219+v220<<(uint(v221)%32))+12)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	*(*uint16)(unsafe.Add(mBase, uint32(v226+v227<<(uint(v221)%32))+4)) = uint16(v231)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+2)))
	v240 = v238 & int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v233+v234<<(uint(v221)%32)))) = uint16(v240)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v247 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v242+v243<<(uint(v221)%32))+6)) = uint16(v247)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v254 = int32(16383)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v254 <= v255 {
		goto L56
	} else {
		goto L57
	}
L40:
	;
	v217 = v209 - v160
	goto L39
L41:
	;
	v188 = v184
	goto L50
L42:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v168 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v217 = int32(0)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v173 = v160
	goto L46
L46:
	;
	v177 = v173 + int32(1)
	if v177&int32(3) == int32(0) {
		v184 = v177
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v209 = v177
	goto L40
L48:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v182 != 0 {
		v173 = v177
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v197 = int32(-2139062144)
	if (int32(16843008)-v194|v194)&v197 == v197 {
		v188 = v188 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v203 = v188
	goto L53
L52:
	;
	goto L51
L53:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v207 != 0 {
		v203 = v203 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v209 = v203
	goto L40
L55:
	;
	goto L54
L56:
	;
	v258 = v254
	goto L58
L57:
	;
	v258 = v255
	goto L58
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v249+v250<<(uint(v221)%32))+8)) = uint16(v258)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v262 = v260 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v262
	v265 = v127 + int32(12)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	if v268 != 0 {
		v127 = v127 + int32(8)
		v129 = v262
		v130 = v265
		goto L30
	} else {
		goto L59
	}
L59:
	;
	goto L31
L60:
	;
	v284 = F_LexizeExec(m, v11+int32(-56), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v284 != 0 {
		v114 = v284
		goto L25
	} else {
		goto L62
	}
L62:
	;
	goto L26
L63:
	;
	goto L6
L64:
	;
	m.G0 = v13 - int32(-64)
	return
}
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
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
	var v83 int32
	_ = v83
	var v85 float32
	_ = v85
	var v89 float64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 float64
	_ = v176
	var v177 int32
	_ = v177
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v202 float64
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v207 int32
	_ = v207
	var v210 float64
	_ = v210
	var v220 float64
	_ = v220
	var v222 float64
	_ = v222
	var v230 float64
	_ = v230
	var v236 float64
	_ = v236
	var v237 int32
	_ = v237
	var v240 float64
	_ = v240
	var v250 float64
	_ = v250
	var v253 float64
	_ = v253
	var v261 float64
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v280 float64
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v301 float64
	_ = v301
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = int64(0)
	if l7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = float64(0.995)
	goto L3
L2:
	;
	v29 = float64(0.005)
	goto L3
L3:
	;
	v36 = F_get_restriction_variable(m, l0, l3, l4, v21-int32(-64), v21+int32(60), v21+int32(59))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v21 + int32(96)
	return v301
L5:
	;
	return float64(0)
L6:
	;
	if v36 == int32(0) {
		v301 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+59)))
	if v42 != int32(1) {
		v280 = v29
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v283 == int32(0) {
		v301 = v280
		goto L4
	} else {
		goto L76
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(7) {
		v280 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+24)))
	if v49 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v280 = float64(0)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v53&int32(-9) != int32(17) {
		v280 = v29
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	switch v62 - int32(17) {
	case 0:
		goto L18
	case 1, 3, 4, 5, 6, 7:
		v280 = v29
		goto L8
	case 2:
		goto L16
	case 8:
		v77 = v62
		v78 = int32(667)
		v79 = int32(98)
		v80 = int32(664)
		goto L15
	default:
		goto L17
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v81 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v77 = int32(25)
	v78 = int32(257)
	v79 = int32(254)
	v80 = int32(255)
	goto L15
L17:
	;
	if v62 != int32(1042) {
		v280 = v29
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v77 = v62
	v78 = int32(1960)
	v79 = int32(1955)
	v80 = int32(1957)
	goto L15
L19:
	;
	v77 = v62
	v78 = int32(1061)
	v79 = int32(1054)
	v80 = int32(1058)
	goto L15
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v85 = *(*float32)(unsafe.Add(mBase, uint32(v82+v83)+8))
	v89 = base.F64_promote_f32(v85)
	goto L22
L21:
	;
	v89 = float64(0)
	goto L22
L22:
	;
	v94 = F_pattern_fixed_prefix(m, v45, l6, l5, v21+int32(52), v21+int32(40))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v96 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v94 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 == v77 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v77
	goto L24
L27:
	;
	if v96 == int32(0) {
		v280 = v261
		goto L8
	} else {
		goto L73
	}
L28:
	;
	if l7 != 0 {
		goto L68
	} else {
		goto L69
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v109 = int32(0)
	v112 = F_var_eq_const(m, v21-int32(-64), v79, l5, v108, v109, int32(1), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if l2 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v250 = v112
	goto L28
L33:
	;
	v116 = l2
	goto L35
L34:
	;
	v114 = F_get_opcode(m, l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	F_fmgr_info(m, v116, v21+int32(8))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v116 = v114
	goto L35
L37:
	;
	v128 = F_histogram_selectivity(m, v21-int32(-64), v21+int32(8), l5, v58, int32(1), v21+int32(36))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if int32(99) < v130 {
		v220 = v128
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v222 = float64(0.0001)
	if base.F64_lt(v220, v222) != 0 {
		v230 = v222
		goto L64
	} else {
		goto L65
	}
L40:
	;
	if v94 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = v21 - int32(-64)
	v137 = m.G0
	v139 = v137 - int32(32)
	m.G0 = v139
	v141 = F_get_opcode(m, v78)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	v202 = float64(1)
	goto L43
L43:
	;
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v21)+40))
	v204 = base.F64_mul(v202, v203)
	if base.F64_lt(v128, float64(0)) != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	F_fmgr_info(m, v141, v139+int32(4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v150 = int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v154 = F_ineq_histogram_selectivity(m, l0, v136, v78, v139+int32(4), v150, v150, l5, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if base.F64_lt(v154, float64(0)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v160 = F_get_opcode(m, v80)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v191 = float64(0.005)
	goto L49
L49:
	;
	m.G0 = v139 + int32(32)
	v202 = v191
	goto L43
L50:
	;
	F_fmgr_info(m, v160, v139+int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v168 = F_make_greater_string(m, v96, v139+int32(4), l5)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v168 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v172 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v176 = F_ineq_histogram_selectivity(m, l0, v136, v80, v139+int32(4), v172, v172, l5, v174, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	v181 = v154
	goto L55
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v183 = int32(0)
	v186 = F_var_eq_const(m, v136, v79, l5, v182, v183, int32(1), v183)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	v181 = base.F64_add(base.F64_add(v154, v176), float64(-1))
	goto L55
L57:
	;
	if base.F64_lt(v186, v181) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v189 = v181
	goto L60
L59:
	;
	v189 = v186
	goto L60
L60:
	;
	v191 = v189
	goto L49
L61:
	;
	v220 = v204
	goto L39
L62:
	;
	goto L63
L63:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v210 = base.F64_div(base.F64_convert_i32_s(v207), float64(100))
	v220 = base.F64_add(base.F64_mul(v128, v210), base.F64_mul(v204, base.F64_sub(float64(1), v210)))
	goto L39
L64:
	;
	v236 = F_mcv_selectivity(m, v21-int32(-64), v21+int32(8), l5, v58, int32(1), v21)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	if base.F64_gt(v220, float64(0.9999)) == int32(0) {
		v230 = v220
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v230 = float64(0.9999)
	goto L64
L67:
	;
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v250 = base.F64_add(v236, base.F64_mul(v230, base.F64_sub(base.F64_sub(float64(1), v89), v240)))
	goto L28
L68:
	;
	v253 = base.F64_sub(base.F64_sub(float64(1), v250), v89)
	goto L70
L69:
	;
	v253 = v250
	goto L70
L70:
	;
	if base.F64_lt(v253, float64(0)) != 0 {
		v261 = float64(0)
		goto L27
	} else {
		goto L71
	}
L71:
	;
	if base.F64_gt(v253, float64(1)) == int32(0) {
		v261 = v253
		goto L27
	} else {
		goto L72
	}
L72:
	;
	v261 = float64(1)
	goto L27
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	F_pfree(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v96)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v280 = v261
	goto L8
L76:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	m.T0[v286].(func(*base.Module, int32))(m, v283)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v301 = v280
	goto L4
}
func F_pcb_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	v3 = F_geterrcode(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(67371461) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v7 == int32(0) {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v10 < int32(0) {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					if v13 == int32(0) {
						return
					} else {
						v16 = F_pg_mbstrlen_with_len(m, v13, v10)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							v20 = F_errposition(m, v16+int32(1))
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_consts[1445])))
	m.T0[v11].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_pgsql_version(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text(m, int32(101477))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_plainto_tsquery_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v17 = F_text_to_cstring(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v24 = F_parse_tsquery(m, v17, int32(1173), v6+int32(8), int32(1), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func F_portuguese_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v1259
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= v9 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v47
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v71 < v61 {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	goto L3
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v54
	goto L2
L6:
	;
	if v47 <= v46 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v46 = v9
	v47 = v15
	goto L6
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v9))))
	switch v19 - int32(227) {
	case 0, 18:
		goto L9
	default:
		goto L7
	}
L9:
	;
	v24 = F_find_among(m, l0, int32(4181824), int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28
	switch v24 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L5
	}
L12:
	;
	v41 = F_slice_from_s(m, l0, int32(2), int32(2152356))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v35 = F_slice_from_s(m, l0, int32(2), int32(2152354))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = v28
	v47 = v32
	goto L6
L15:
	;
	if int32(0) <= v35 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v1259 = v35
	goto L1
L17:
	;
	if int32(0) <= v41 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v1259 = v41
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 + int32(1)
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v561 < v61 {
		goto L168
	} else {
		goto L169
	}
L21:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = v547
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v341 < v61 {
		goto L102
	} else {
		goto L103
	}
L23:
	;
	if v114 != 0 {
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v73 = v61
	goto L26
L25:
	;
	v73 = v71
	goto L26
L26:
	;
	goto L28
L27:
	;
	v114 = v110
	goto L23
L28:
	;
	if v61 == v73 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v110 = int32(0)
	goto L27
L30:
	;
	v114 = int32(-1)
	goto L23
L31:
	;
	goto L32
L32:
	;
	v85 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v61))))
	if int32(250) < v88 {
		v110 = v85
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v90 = v88 - int32(97)
	if v90 < int32(0) {
		v110 = v85
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v90)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v96)>>(uint(v90&int32(7))%32))&int32(1) == int32(0) {
		v110 = v85
		goto L27
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 + int32(1)
	goto L36
L36:
	;
	goto L29
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v124 < v115 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v229 < v115 {
		goto L72
	} else {
		goto L73
	}
L39:
	;
	if v164 != 0 {
		goto L38
	} else {
		goto L54
	}
L40:
	;
	v126 = v115
	goto L42
L41:
	;
	v126 = v124
	goto L42
L42:
	;
	goto L44
L43:
	;
	v164 = v161
	goto L39
L44:
	;
	if v115 == v126 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v161 = int32(0)
	goto L43
L46:
	;
	v164 = int32(-1)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v115))))
	if int32(250) < v139 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115 + int32(1)
	goto L53
L50:
	;
	v141 = v139 - int32(97)
	if v141 < int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v144 = int32(1)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v141)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v148)>>(uint(v141&int32(7))%32))&v144 != 0 {
		v161 = v144
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	goto L45
L54:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v173 < v172 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v213 < int32(0) {
		goto L38
	} else {
		goto L70
	}
L56:
	;
	v175 = v172
	goto L58
L57:
	;
	v175 = v173
	goto L58
L58:
	;
	v182 = v172
	goto L60
L59:
	;
	v213 = v193
	goto L55
L60:
	;
	if v182 == v175 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v213 = int32(-1)
	goto L55
L63:
	;
	goto L64
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v182))))
	if int32(250) < v188 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v205 = v182 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v205
	v182 = v205
	goto L60
L66:
	;
	v190 = v188 - int32(97)
	if v190 < int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v193 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v190)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v197)>>(uint(v190&int32(7))%32))&v193 != 0 {
		goto L59
	} else {
		goto L68
	}
L68:
	;
	goto L65
L70:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v547 = v216 + v213
	goto L21
L71:
	;
	if v272 != 0 {
		goto L22
	} else {
		goto L85
	}
L72:
	;
	v231 = v115
	goto L74
L73:
	;
	v231 = v229
	goto L74
L74:
	;
	goto L76
L75:
	;
	v272 = v268
	goto L71
L76:
	;
	if v115 == v231 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v268 = int32(0)
	goto L75
L78:
	;
	v272 = int32(-1)
	goto L71
L79:
	;
	goto L80
L80:
	;
	v243 = int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v115))))
	if int32(250) < v246 {
		v268 = v243
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v248 = v246 - int32(97)
	if v248 < int32(0) {
		v268 = v243
		goto L75
	} else {
		goto L82
	}
L82:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v248)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v254)>>(uint(v248&int32(7))%32))&int32(1) == int32(0) {
		v268 = v243
		goto L75
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115 + int32(1)
	goto L84
L84:
	;
	goto L77
L85:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v282 < v281 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v325 < int32(0) {
		goto L22
	} else {
		goto L100
	}
L87:
	;
	v284 = v281
	goto L89
L88:
	;
	v284 = v282
	goto L89
L89:
	;
	v291 = v281
	goto L91
L90:
	;
	v325 = int32(1)
	goto L86
L91:
	;
	if v291 == v284 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v325 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v291))))
	if int32(250) < v299 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v301 = v299 - int32(97)
	if v301 < int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v301)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v307)>>(uint(v301&int32(7))%32))&int32(1) == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v316 = v291 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v316
	v291 = v316
	goto L91
L100:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v547 = v328 + v325
	goto L21
L101:
	;
	if v381 != 0 {
		goto L20
	} else {
		goto L116
	}
L102:
	;
	v343 = v61
	goto L104
L103:
	;
	v343 = v341
	goto L104
L104:
	;
	goto L106
L105:
	;
	v381 = v378
	goto L101
L106:
	;
	if v61 == v343 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v378 = int32(0)
	goto L105
L108:
	;
	v381 = int32(-1)
	goto L101
L109:
	;
	goto L110
L110:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+v61))))
	if int32(250) < v356 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 + int32(1)
	goto L115
L112:
	;
	v358 = v356 - int32(97)
	if v358 < int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v361 = int32(1)
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v358)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v365)>>(uint(v358&int32(7))%32))&v361 != 0 {
		v378 = v361
		goto L105
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	goto L107
L116:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v391 < v382 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v382
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v496 < v382 {
		goto L151
	} else {
		goto L152
	}
L118:
	;
	if v431 != 0 {
		goto L117
	} else {
		goto L133
	}
L119:
	;
	v393 = v382
	goto L121
L120:
	;
	v393 = v391
	goto L121
L121:
	;
	goto L123
L122:
	;
	v431 = v428
	goto L118
L123:
	;
	if v382 == v393 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v428 = int32(0)
	goto L122
L125:
	;
	v431 = int32(-1)
	goto L118
L126:
	;
	goto L127
L127:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v382))))
	if int32(250) < v406 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v382 + int32(1)
	goto L132
L129:
	;
	v408 = v406 - int32(97)
	if v408 < int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v411 = int32(1)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v408)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v415)>>(uint(v408&int32(7))%32))&v411 != 0 {
		v428 = v411
		goto L122
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	goto L124
L133:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v440 < v439 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v480 < int32(0) {
		goto L117
	} else {
		goto L149
	}
L135:
	;
	v442 = v439
	goto L137
L136:
	;
	v442 = v440
	goto L137
L137:
	;
	v449 = v439
	goto L139
L138:
	;
	v480 = v460
	goto L134
L139:
	;
	if v449 == v442 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v480 = int32(-1)
	goto L134
L142:
	;
	goto L143
L143:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453+v449))))
	if int32(250) < v455 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v472 = v449 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v472
	v449 = v472
	goto L139
L145:
	;
	v457 = v455 - int32(97)
	if v457 < int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v460 = int32(1)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v464)>>(uint(v457&int32(7))%32))&v460 != 0 {
		goto L138
	} else {
		goto L147
	}
L147:
	;
	goto L144
L149:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v547 = v483 + v480
	goto L21
L150:
	;
	if v539 != 0 {
		goto L20
	} else {
		goto L164
	}
L151:
	;
	v498 = v382
	goto L153
L152:
	;
	v498 = v496
	goto L153
L153:
	;
	goto L155
L154:
	;
	v539 = v535
	goto L150
L155:
	;
	if v382 == v498 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v535 = int32(0)
	goto L154
L157:
	;
	v539 = int32(-1)
	goto L150
L158:
	;
	goto L159
L159:
	;
	v510 = int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v382))))
	if int32(250) < v513 {
		v535 = v510
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v515 = v513 - int32(97)
	if v515 < int32(0) {
		v535 = v510
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v515)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v521)>>(uint(v515&int32(7))%32))&int32(1) == int32(0) {
		v535 = v510
		goto L154
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v382 + int32(1)
	goto L163
L163:
	;
	goto L156
L164:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v541 <= v540 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v547 = v540 + int32(1)
	goto L21
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v61
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v782
	if v782-int32(2) <= v61 {
		goto L231
	} else {
		goto L232
	}
L167:
	;
	if v601 < int32(0) {
		goto L166
	} else {
		goto L182
	}
L168:
	;
	v563 = v61
	goto L170
L169:
	;
	v563 = v561
	goto L170
L170:
	;
	v570 = v61
	goto L172
L171:
	;
	v601 = v581
	goto L167
L172:
	;
	if v570 == v563 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v601 = int32(-1)
	goto L167
L175:
	;
	goto L176
L176:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574+v570))))
	if int32(250) < v576 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v593 = v570 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v593
	v570 = v593
	goto L172
L178:
	;
	v578 = v576 - int32(97)
	if v578 < int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v581 = int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v578)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v585)>>(uint(v578&int32(7))%32))&v581 != 0 {
		goto L171
	} else {
		goto L180
	}
L180:
	;
	goto L177
L182:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v605 = v604 + v601
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v605
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v616 < v605 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v659 < int32(0) {
		goto L166
	} else {
		goto L197
	}
L184:
	;
	v618 = v605
	goto L186
L185:
	;
	v618 = v616
	goto L186
L186:
	;
	v625 = v605
	goto L188
L187:
	;
	v659 = int32(1)
	goto L183
L188:
	;
	if v625 == v618 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v659 = int32(-1)
	goto L183
L191:
	;
	goto L192
L192:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v625))))
	if int32(250) < v633 {
		goto L187
	} else {
		goto L193
	}
L193:
	;
	v635 = v633 - int32(97)
	if v635 < int32(0) {
		goto L187
	} else {
		goto L194
	}
L194:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v635)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v641)>>(uint(v635&int32(7))%32))&int32(1) == int32(0) {
		goto L187
	} else {
		goto L195
	}
L195:
	;
	v650 = v625 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650
	v625 = v650
	goto L188
L197:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v663 = v662 + v659
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v663
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v665)+4)) = v663
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v675 < v674 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v715 < int32(0) {
		goto L166
	} else {
		goto L213
	}
L199:
	;
	v677 = v674
	goto L201
L200:
	;
	v677 = v675
	goto L201
L201:
	;
	v684 = v674
	goto L203
L202:
	;
	v715 = v695
	goto L198
L203:
	;
	if v684 == v677 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v715 = int32(-1)
	goto L198
L206:
	;
	goto L207
L207:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688+v684))))
	if int32(250) < v690 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v707 = v684 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v707
	v684 = v707
	goto L203
L209:
	;
	v692 = v690 - int32(97)
	if v692 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v695 = int32(1)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v692)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v699)>>(uint(v692&int32(7))%32))&v695 != 0 {
		goto L202
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v719 = v718 + v715
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v719
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v730 < v719 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v773 < int32(0) {
		goto L166
	} else {
		goto L228
	}
L215:
	;
	v732 = v719
	goto L217
L216:
	;
	v732 = v730
	goto L217
L217:
	;
	v739 = v719
	goto L219
L218:
	;
	v773 = int32(1)
	goto L214
L219:
	;
	if v739 == v732 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v773 = int32(-1)
	goto L214
L222:
	;
	goto L223
L223:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745+v739))))
	if int32(250) < v747 {
		goto L218
	} else {
		goto L224
	}
L224:
	;
	v749 = v747 - int32(97)
	if v749 < int32(0) {
		goto L218
	} else {
		goto L225
	}
L225:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v749)>>(uint(int32(3))%32)))+uint32(_consts[1469]))))
	if int32(base.Ui32(v755)>>(uint(v749&int32(7))%32))&int32(1) == int32(0) {
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v764 = v739 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v764
	v739 = v764
	goto L219
L228:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v776))) = v777 + v773
	goto L166
L229:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1130
	v1135 = F_find_among_b(m, l0, int32(4185552), int32(4))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L10
	} else {
		goto L336
	}
L230:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1097
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1097
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1097 <= v1100 {
		goto L229
	} else {
		goto L328
	}
L231:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1054
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+8))
	if v1057 <= v1054 {
		goto L316
	} else {
		goto L317
	}
L232:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788+v782-int32(1)))))
	if v792&int32(224) != int32(96) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	if int32(1)<<(uint(v792)%32)&int32(823330) == int32(0) {
		goto L231
	} else {
		goto L234
	}
L234:
	;
	v805 = F_find_among_b(m, l0, int32(4181888), int32(45))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	if v805 == int32(0) {
		goto L231
	} else {
		goto L236
	}
L236:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v809
	switch v805 - int32(1) {
	case 0:
		goto L245
	case 1:
		goto L244
	case 2:
		goto L243
	case 3:
		goto L242
	case 4:
		goto L241
	case 5:
		goto L240
	case 6:
		goto L239
	case 7:
		goto L238
	case 8:
		goto L237
	default:
		goto L230
	}
L237:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+8))
	if v809 < v1032 {
		goto L231
	} else {
		goto L310
	}
L238:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	if v809 < v996 {
		goto L231
	} else {
		goto L299
	}
L239:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)))
	if v809 < v955 {
		goto L231
	} else {
		goto L288
	}
L240:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	if v809 < v920 {
		goto L231
	} else {
		goto L278
	}
L241:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v809 < v848 {
		goto L231
	} else {
		goto L258
	}
L242:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	if v809 < v839 {
		goto L231
	} else {
		goto L255
	}
L243:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	if v809 < v830 {
		goto L231
	} else {
		goto L252
	}
L244:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)))
	if v809 < v821 {
		goto L231
	} else {
		goto L249
	}
L245:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	if v809 < v814 {
		goto L231
	} else {
		goto L246
	}
L246:
	;
	v816 = F_slice_del(m, l0)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	if int32(0) <= v816 {
		goto L230
	} else {
		goto L248
	}
L248:
	;
	v1259 = v816
	goto L1
L249:
	;
	v825 = F_slice_from_s(m, l0, int32(3), int32(2152388))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	if int32(0) <= v825 {
		goto L230
	} else {
		goto L251
	}
L251:
	;
	v1259 = v825
	goto L1
L252:
	;
	v834 = F_slice_from_s(m, l0, int32(1), int32(2152391))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	if int32(0) <= v834 {
		goto L230
	} else {
		goto L254
	}
L254:
	;
	v1259 = v834
	goto L1
L255:
	;
	v843 = F_slice_from_s(m, l0, int32(4), int32(2152392))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L10
	} else {
		goto L256
	}
L256:
	;
	if int32(0) <= v843 {
		goto L230
	} else {
		goto L257
	}
L257:
	;
	v1259 = v843
	goto L1
L258:
	;
	v850 = F_slice_del(m, l0)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	if v850 < int32(0) {
		v1259 = v850
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v854
	v857 = v854 - int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v857 <= v858 {
		goto L230
	} else {
		goto L261
	}
L261:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+v857))))
	if v862&int32(224) != int32(96) {
		goto L230
	} else {
		goto L262
	}
L262:
	;
	if int32(1)<<(uint(v862)%32)&int32(4718616) == int32(0) {
		goto L230
	} else {
		goto L263
	}
L263:
	;
	v875 = F_find_among_b(m, l0, int32(4182800), int32(4))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L264
	}
L264:
	;
	if v875 == int32(0) {
		goto L230
	} else {
		goto L265
	}
L265:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v879
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v879 < v882 {
		goto L230
	} else {
		goto L266
	}
L266:
	;
	v884 = F_slice_del(m, l0)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L10
	} else {
		goto L267
	}
L267:
	;
	if v884 < int32(0) {
		v1259 = v884
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v875 != int32(1) {
		goto L230
	} else {
		goto L269
	}
L269:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v890
	v892 = int32(2)
	v894 = int32(0)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v890-v897 < v892 {
		v907 = v894
		goto L271
	} else {
		goto L272
	}
L270:
	;
	if v907 == int32(0) {
		goto L230
	} else {
		goto L274
	}
L271:
	;
	goto L270
L272:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v903 = F_memcmp(m, v900+v890-v892, int32(2152396), v892)
	mBase = m.M
	if v903 != 0 {
		v907 = v894
		goto L271
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v890 - v892
	v907 = int32(1)
	goto L271
L274:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v910
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	if v910 < v913 {
		goto L230
	} else {
		goto L275
	}
L275:
	;
	v915 = F_slice_del(m, l0)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L10
	} else {
		goto L276
	}
L276:
	;
	if int32(0) <= v915 {
		goto L230
	} else {
		goto L277
	}
L277:
	;
	v1259 = v915
	goto L1
L278:
	;
	v922 = F_slice_del(m, l0)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L10
	} else {
		goto L279
	}
L279:
	;
	if v922 < int32(0) {
		v1259 = v922
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v926
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v926-int32(3) <= v928 {
		goto L230
	} else {
		goto L281
	}
L281:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932+v926-int32(1)))))
	switch v936 - int32(101) {
	case 0, 7:
		goto L282
	default:
		goto L230
	}
L282:
	;
	v941 = F_find_among_b(m, l0, int32(4182880), int32(3))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L10
	} else {
		goto L283
	}
L283:
	;
	if v941 == int32(0) {
		goto L230
	} else {
		goto L284
	}
L284:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v945
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)))
	if v945 < v948 {
		goto L230
	} else {
		goto L285
	}
L285:
	;
	v950 = F_slice_del(m, l0)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L10
	} else {
		goto L286
	}
L286:
	;
	if int32(0) <= v950 {
		goto L230
	} else {
		goto L287
	}
L287:
	;
	v1259 = v950
	goto L1
L288:
	;
	v957 = F_slice_del(m, l0)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L10
	} else {
		goto L289
	}
L289:
	;
	if v957 < int32(0) {
		v1259 = v957
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v961
	v964 = v961 - int32(1)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v964 <= v965 {
		goto L230
	} else {
		goto L291
	}
L291:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967+v964))))
	if v969&int32(224) != int32(96) {
		goto L230
	} else {
		goto L292
	}
L292:
	;
	if int32(1)<<(uint(v969)%32)&int32(4198408) == int32(0) {
		goto L230
	} else {
		goto L293
	}
L293:
	;
	v982 = F_find_among_b(m, l0, int32(4182944), int32(3))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L10
	} else {
		goto L294
	}
L294:
	;
	if v982 == int32(0) {
		goto L230
	} else {
		goto L295
	}
L295:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v986
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v988)))
	if v986 < v989 {
		goto L230
	} else {
		goto L296
	}
L296:
	;
	v991 = F_slice_del(m, l0)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L10
	} else {
		goto L297
	}
L297:
	;
	if int32(0) <= v991 {
		goto L230
	} else {
		goto L298
	}
L298:
	;
	v1259 = v991
	goto L1
L299:
	;
	v998 = F_slice_del(m, l0)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L10
	} else {
		goto L300
	}
L300:
	;
	if v998 < int32(0) {
		v1259 = v998
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1002
	v1004 = int32(2)
	v1006 = int32(0)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1002-v1009 < v1004 {
		v1019 = v1006
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if v1019 == int32(0) {
		goto L230
	} else {
		goto L306
	}
L303:
	;
	goto L302
L304:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1015 = F_memcmp(m, v1012+v1002-v1004, int32(2152398), v1004)
	mBase = m.M
	if v1015 != 0 {
		v1019 = v1006
		goto L303
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1002 - v1004
	v1019 = int32(1)
	goto L303
L306:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1022
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	if v1022 < v1025 {
		goto L230
	} else {
		goto L307
	}
L307:
	;
	v1027 = F_slice_del(m, l0)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L10
	} else {
		goto L308
	}
L308:
	;
	if int32(0) <= v1027 {
		goto L230
	} else {
		goto L309
	}
L309:
	;
	v1259 = v1027
	goto L1
L310:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v809 <= v1034 {
		goto L231
	} else {
		goto L311
	}
L311:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036+v809-int32(1)))))
	if v1040 != int32(101) {
		goto L231
	} else {
		goto L312
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	v1048 = F_slice_from_s(m, l0, int32(2), int32(2152400))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L10
	} else {
		goto L313
	}
L313:
	;
	if int32(0) <= v1048 {
		goto L230
	} else {
		goto L314
	}
L314:
	;
	v1259 = v1048
	goto L1
L315:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1087
	v1089 = F_slice_del(m, l0)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L10
	} else {
		goto L326
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1054
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1057
	v1064 = F_find_among_b(m, l0, int32(4183008), int32(120))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L10
	} else {
		goto L319
	}
L317:
	;
	v1068 = v1054
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1068
	v1074 = F_find_among_b(m, l0, int32(4185408), int32(7))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L10
	} else {
		goto L321
	}
L319:
	;
	if v1064 != 0 {
		goto L315
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1060
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1068 = v1067
	goto L318
L321:
	;
	if v1074 == int32(0) {
		goto L229
	} else {
		goto L322
	}
L322:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1078
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+8))
	if v1078 < v1081 {
		goto L229
	} else {
		goto L323
	}
L323:
	;
	v1083 = F_slice_del(m, l0)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L10
	} else {
		goto L324
	}
L324:
	;
	if int32(0) <= v1083 {
		goto L229
	} else {
		goto L325
	}
L325:
	;
	v1259 = v1083
	goto L1
L326:
	;
	if v1089 < int32(0) {
		v1259 = v1089
		goto L1
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1060
	goto L230
L328:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1103 = v1102 + v1097
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103-int32(1)))))
	if v1106 != int32(105) {
		goto L229
	} else {
		goto L329
	}
L329:
	;
	v1110 = v1097 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1110
	if v1110 <= v1100 {
		goto L229
	} else {
		goto L330
	}
L330:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103-int32(2)))))
	if v1116 != int32(99) {
		goto L229
	} else {
		goto L331
	}
L331:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	if v1097 <= v1120 {
		goto L229
	} else {
		goto L332
	}
L332:
	;
	v1122 = F_slice_del(m, l0)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L10
	} else {
		goto L333
	}
L333:
	;
	if v1122 < int32(0) {
		v1259 = v1122
		goto L1
	} else {
		goto L334
	}
L334:
	;
	goto L229
L335:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1206
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1210 = v1206
	v1211 = v1208
	goto L357
L336:
	;
	if v1135 == int32(0) {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1139
	switch v1135 - int32(1) {
	case 0:
		goto L339
	case 1:
		goto L338
	default:
		goto L335
	}
L338:
	;
	v1197 = F_slice_from_s(m, l0, int32(1), int32(2153162))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L10
	} else {
		goto L355
	}
L339:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+8))
	if v1139 < v1144 {
		goto L335
	} else {
		goto L340
	}
L340:
	;
	v1146 = F_slice_del(m, l0)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L10
	} else {
		goto L341
	}
L341:
	;
	if v1146 < int32(0) {
		v1259 = v1146
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1150
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1150 <= v1152 {
		goto L335
	} else {
		goto L343
	}
L343:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1155 = v1154 + v1150
	v1157 = v1155 - int32(1)
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if v1158 != int32(117) {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1186
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+8))
	if v1186 < v1189 {
		goto L335
	} else {
		goto L352
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1150
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if v1173 != int32(105) {
		goto L335
	} else {
		goto L349
	}
L346:
	;
	v1162 = v1150 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1162
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1162
	if v1162 <= v1152 {
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155-int32(2)))))
	if v1168 == int32(103) {
		v1186 = v1162
		goto L344
	} else {
		goto L348
	}
L348:
	;
	goto L345
L349:
	;
	v1177 = v1150 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1177
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1177
	if v1177 <= v1152 {
		goto L335
	} else {
		goto L350
	}
L350:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155-int32(2)))))
	if v1183 != int32(99) {
		goto L335
	} else {
		goto L351
	}
L351:
	;
	v1186 = v1177
	goto L344
L352:
	;
	v1191 = F_slice_del(m, l0)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L10
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v1191 {
		goto L335
	} else {
		goto L354
	}
L354:
	;
	v1259 = v1191
	goto L1
L355:
	;
	if v1197 < int32(0) {
		v1259 = v1197
		goto L1
	} else {
		goto L356
	}
L356:
	;
	goto L335
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1210
	v1217 = v1210 + int32(1)
	if v1217 < v1211 {
		goto L363
	} else {
		goto L364
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1206
	v1259 = int32(1)
	goto L1
L359:
	;
	goto L358
L360:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1210 = v1255
	v1211 = v1254
	goto L357
L361:
	;
	if v1247 <= v1246 {
		goto L359
	} else {
		goto L375
	}
L362:
	;
	v1227 = F_find_among(m, l0, int32(4185632), int32(3))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L10
	} else {
		goto L367
	}
L363:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219+v1217))))
	if v1221 == int32(126) {
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1210
	v1246 = v1210
	v1247 = v1211
	goto L361
L366:
	;
	goto L365
L367:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1229
	switch v1227 - int32(1) {
	case 0:
		goto L370
	case 1:
		goto L369
	case 2:
		goto L368
	default:
		goto L360
	}
L368:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1246 = v1229
	v1247 = v1245
	goto L361
L369:
	;
	v1241 = F_slice_from_s(m, l0, int32(1), int32(2153168))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L10
	} else {
		goto L373
	}
L370:
	;
	v1235 = F_slice_from_s(m, l0, int32(1), int32(2153167))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L10
	} else {
		goto L371
	}
L371:
	;
	if int32(0) <= v1235 {
		goto L360
	} else {
		goto L372
	}
L372:
	;
	v1259 = v1235
	goto L1
L373:
	;
	if int32(0) <= v1241 {
		goto L360
	} else {
		goto L374
	}
L374:
	;
	v1259 = v1241
	goto L1
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1246 + int32(1)
	goto L360
}
func F_pow(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v92 int32
	_ = v92
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 float64
	_ = v167
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v191 float64
	_ = v191
	var v201 float64
	_ = v201
	var v204 float64
	_ = v204
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v217 float64
	_ = v217
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v233 int32
	_ = v233
	var v236 float64
	_ = v236
	var v240 int64
	_ = v240
	var v245 float64
	_ = v245
	var v248 float64
	_ = v248
	var v251 float64
	_ = v251
	var v254 float64
	_ = v254
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v279 float64
	_ = v279
	var v282 float64
	_ = v282
	var v285 float64
	_ = v285
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v297 float64
	_ = v297
	var v300 float64
	_ = v300
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v307 float64
	_ = v307
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v342 float64
	_ = v342
	var v344 float64
	_ = v344
	var v356 float64
	_ = v356
	var v358 float64
	_ = v358
	var v360 int32
	_ = v360
	var v362 float64
	_ = v362
	var v365 float64
	_ = v365
	var v366 float64
	_ = v366
	var v367 float64
	_ = v367
	var v369 float64
	_ = v369
	var v372 float64
	_ = v372
	var v376 float64
	_ = v376
	var v377 float64
	_ = v377
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v387 float64
	_ = v387
	var v390 float64
	_ = v390
	var v393 int64
	_ = v393
	var v398 int32
	_ = v398
	var v401 float64
	_ = v401
	var v404 float64
	_ = v404
	var v407 int64
	_ = v407
	var v412 int64
	_ = v412
	var v421 float64
	_ = v421
	var v427 int64
	_ = v427
	var v428 float64
	_ = v428
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v434 float64
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v454 float64
	_ = v454
	var v455 float64
	_ = v455
	var v462 float64
	_ = v462
	var v465 float64
	_ = v465
	var v469 float64
	_ = v469
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v491 float64
	_ = v491
	var v495 float64
	_ = v495
	v11 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v24 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
	v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l1)) >> (uint(int64(52)) % 64)))
	v29 = int32(2047)
	v30 = v28 & v29
	v32 = v30 - int32(1086)
	v33 = base.I64_reinterpret_f64(l1)
	v34 = base.I64_reinterpret_f64(l0)
	if base.B2i32(base.Ui32(int32(-129)) < base.Ui32(v32))&base.B2i32(base.Ui32(int32(-2046)) <= base.Ui32(v24-v29)) != 0 {
		v212 = v34
		v214 = v11
		v215 = int64(-134217728)
		v217 = base.F64_reinterpret_i64(v33 & v215)
		v219 = v212 - int64(4604531861337669632)
		v220 = int64(52)
		v223 = base.F64_convert_i32_s(base.I32_wrap_i64(v219 >> (uint(v220) % 64)))
		v225 = *(*float64)(unsafe.Add(mBase, _consts[1571]))
		v233 = base.I32_wrap_i64(int64(base.Ui64(v219)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
		v236 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1572])))
		v240 = v212 - v219&int64(-4503599627370496)
		v245 = base.F64_reinterpret_i64((v240 + int64(2147483648)) & int64(-4294967296))
		v248 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1573])))
		v251 = base.F64_add(base.F64_mul(v245, v248), float64(-1))
		v254 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v240), v245), v248)
		v255 = base.F64_add(v251, v254)
		v257 = *(*float64)(unsafe.Add(mBase, _consts[1574]))
		v261 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1575])))
		v262 = base.F64_add(base.F64_mul(v223, v257), v261)
		v263 = base.F64_add(v255, v262)
		v268 = *(*float64)(unsafe.Add(mBase, _consts[1576]))
		v269 = base.F64_mul(v255, v268)
		v270 = base.F64_mul(v251, v268)
		v274 = base.F64_mul(v251, v270)
		v275 = base.F64_add(v263, v274)
		v279 = base.F64_mul(v255, v269)
		v282 = *(*float64)(unsafe.Add(mBase, _consts[1577]))
		v285 = *(*float64)(unsafe.Add(mBase, _consts[1578]))
		v289 = *(*float64)(unsafe.Add(mBase, _consts[1579]))
		v292 = *(*float64)(unsafe.Add(mBase, _consts[1580]))
		v297 = *(*float64)(unsafe.Add(mBase, _consts[1581]))
		v300 = *(*float64)(unsafe.Add(mBase, _consts[1582]))
		v304 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v236), base.F64_add(v255, base.F64_sub(v262, v263))), base.F64_mul(v254, base.F64_add(v269, v270))), base.F64_add(v274, base.F64_sub(v263, v275))), base.F64_mul(base.F64_mul(v255, v279), base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v255, v282), v285)), base.F64_add(base.F64_mul(v255, v289), v292))), base.F64_add(base.F64_mul(v255, v297), v300))))
		v305 = base.F64_add(v275, v304)
		v307 = base.F64_add(v304, base.F64_sub(v275, v305))
		*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v307
		v312 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v305) & v215)
		v313 = base.F64_mul(v217, v312)
		v326 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v313))>>(uint(v220)%64))) & int32(2047)
		v331 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
		if base.Ui32(v326-v331) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v331) {
			v360 = v326
			v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
			v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
			v366 = base.F64_add(base.F64_mul(v313, v362), v365)
			v367 = base.F64_sub(v366, v365)
			v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
			v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
			v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
			v377 = base.F64_mul(v376, v376)
			v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
			v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
			v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
			v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
			v393 = base.I64_reinterpret_f64(v366)
			v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
			v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
			v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
			v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
			v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
			if v360 == int32(0) {
				if v393&int64(2147483648) == int64(0) {
					v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
					v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
				} else {
					v427 = v412 + int64(4602678819172646912)
					v428 = base.F64_reinterpret_i64(v427)
					v429 = base.F64_mul(v428, v404)
					v430 = base.F64_add(v429, v428)
					if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
						v434 = float64(2.2250738585072014e-308)
						v436 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
						v443 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
						if base.F64_lt(v430, float64(0)) != 0 {
							v454 = float64(-1)
						} else {
							v454 = float64(1)
						}
						v455 = base.F64_add(v430, v454)
						v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
						if base.F64_eq(v462, float64(0)) != 0 {
							v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
						} else {
							v465 = v462
						}
						v469 = v465
					} else {
						v469 = v430
					}
					v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
				}
				v491 = v478
			} else {
				v479 = base.F64_reinterpret_i64(v412)
				v491 = base.F64_add(base.F64_mul(v479, v404), v479)
			}
		} else {
			if base.Ui32(v326) < base.Ui32(v331) {
				v342 = base.F64_add(v313, float64(1))
				if v214 != 0 {
					v344 = base.F64_neg(v342)
				} else {
					v344 = v342
				}
				v491 = v344
			} else {
				if base.Ui32(v326) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
					v360 = int32(0)
					v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
					v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
					v366 = base.F64_add(base.F64_mul(v313, v362), v365)
					v367 = base.F64_sub(v366, v365)
					v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
					v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
					v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
					v377 = base.F64_mul(v376, v376)
					v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
					v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
					v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
					v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
					v393 = base.I64_reinterpret_f64(v366)
					v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
					v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
					v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
					v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
					v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
					if v360 == int32(0) {
						if v393&int64(2147483648) == int64(0) {
							v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
							v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
						} else {
							v427 = v412 + int64(4602678819172646912)
							v428 = base.F64_reinterpret_i64(v427)
							v429 = base.F64_mul(v428, v404)
							v430 = base.F64_add(v429, v428)
							if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
								v434 = float64(2.2250738585072014e-308)
								v436 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
								v443 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
								if base.F64_lt(v430, float64(0)) != 0 {
									v454 = float64(-1)
								} else {
									v454 = float64(1)
								}
								v455 = base.F64_add(v430, v454)
								v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
								if base.F64_eq(v462, float64(0)) != 0 {
									v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
								} else {
									v465 = v462
								}
								v469 = v465
							} else {
								v469 = v430
							}
							v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
						}
						v491 = v478
					} else {
						v479 = base.F64_reinterpret_i64(v412)
						v491 = base.F64_add(base.F64_mul(v479, v404), v479)
					}
				} else {
					if base.I64_reinterpret_f64(v313) < int64(0) {
						v356 = F___math_xflow(m, v214, float64(1.2882297539194267e-231))
						mBase = m.M
						v491 = v356
					} else {
						v358 = F___math_xflow(m, v214, float64(3.105036184601418e+231))
						mBase = m.M
						v491 = v358
					}
				}
			}
		}
		v495 = v491
	} else {
		if base.Ui64(v33<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
			v48 = float64(1)
			if v34 == int64(4607182418800017408) {
				v495 = v48
			} else {
				v52 = v33 << (uint(int64(1)) % 64)
				if v52 == int64(0) {
					v495 = v48
				} else {
					v58 = v34 << (uint(int64(1)) % 64)
					if base.B2i32(base.Ui64(v52) < base.Ui64(int64(-9007199254740991)))&base.B2i32(base.Ui64(v58) <= base.Ui64(int64(-9007199254740992))) == int32(0) {
						v495 = base.F64_add(l0, l1)
					} else {
						if v58 == int64(9214364837600034816) {
							v495 = v48
						} else {
							if base.B2i32(v33 < int64(0))^base.B2i32(base.Ui64(v58) < base.Ui64(int64(9214364837600034816))) != 0 {
								v74 = float64(0)
							} else {
								v74 = base.F64_mul(l1, l1)
							}
							v495 = v74
						}
					}
				}
			}
		} else {
			if base.Ui64(v34<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
				v81 = base.F64_mul(l0, l0)
				if v34 < int64(0) {
					v92 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v92) < base.Ui32(int32(1023)) {
						v116 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v92) {
							v116 = int32(2)
						} else {
							v99 = int64(1)
							v103 = v99 << (uint(base.I64_extend_i32_u(int32(1075)-v92)) % 64)
							if (v103-v99)&v33 != int64(0) {
								v116 = int32(0)
							} else {
								if v33&v103 == int64(0) {
									v114 = int32(2)
								} else {
									v114 = int32(1)
								}
								v116 = v114
							}
						}
					}
					if v116 == int32(1) {
						v119 = base.F64_neg(v81)
					} else {
						v119 = v81
					}
					v120 = v119
				} else {
					v120 = v81
				}
				if int64(0) <= v33 {
					v495 = v120
				} else {
					v124 = base.F64_div(float64(1), v120)
					v126 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v126-int32(16))+8)) = v124
					v495 = v124
				}
			} else {
				if v34 < int64(0) {
					v140 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v140) < base.Ui32(int32(1023)) {
						v164 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v140) {
							v164 = int32(2)
						} else {
							v147 = int64(1)
							v151 = v147 << (uint(base.I64_extend_i32_u(int32(1075)-v140)) % 64)
							if (v151-v147)&v33 != int64(0) {
								v164 = int32(0)
							} else {
								if v33&v151 == int64(0) {
									v162 = int32(2)
								} else {
									v162 = int32(1)
								}
								v164 = v162
							}
						}
					}
					if v164 == int32(0) {
						v167 = base.F64_sub(l0, l0)
						v495 = base.F64_div(v167, v167)
					} else {
						v178 = base.I64_reinterpret_f64(l0) & int64(9223372036854775807)
						v179 = v24 & int32(2047)
						v180 = base.B2i32(v164 == int32(1)) << (uint(int32(18)) % 32)
						if base.Ui32(v32) <= base.Ui32(int32(-129)) {
							if v178 == int64(4607182418800017408) {
								v495 = float64(1)
							} else {
								if base.Ui32(v30) <= base.Ui32(int32(957)) {
									if base.Ui64(int64(4607182418800017408)) < base.Ui64(v178) {
										v191 = l1
									} else {
										v191 = base.F64_neg(l1)
									}
									v495 = base.F64_add(v191, float64(1))
								} else {
									if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v178)) {
										v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
										mBase = m.M
										v495 = v201
									} else {
										v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
										mBase = m.M
										v495 = v204
									}
								}
							}
						} else {
							if v179 != 0 {
								v212 = v178
								v214 = v180
							} else {
								v212 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
								v214 = v180
							}
							v215 = int64(-134217728)
							v217 = base.F64_reinterpret_i64(v33 & v215)
							v219 = v212 - int64(4604531861337669632)
							v220 = int64(52)
							v223 = base.F64_convert_i32_s(base.I32_wrap_i64(v219 >> (uint(v220) % 64)))
							v225 = *(*float64)(unsafe.Add(mBase, _consts[1571]))
							v233 = base.I32_wrap_i64(int64(base.Ui64(v219)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
							v236 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1572])))
							v240 = v212 - v219&int64(-4503599627370496)
							v245 = base.F64_reinterpret_i64((v240 + int64(2147483648)) & int64(-4294967296))
							v248 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1573])))
							v251 = base.F64_add(base.F64_mul(v245, v248), float64(-1))
							v254 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v240), v245), v248)
							v255 = base.F64_add(v251, v254)
							v257 = *(*float64)(unsafe.Add(mBase, _consts[1574]))
							v261 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1575])))
							v262 = base.F64_add(base.F64_mul(v223, v257), v261)
							v263 = base.F64_add(v255, v262)
							v268 = *(*float64)(unsafe.Add(mBase, _consts[1576]))
							v269 = base.F64_mul(v255, v268)
							v270 = base.F64_mul(v251, v268)
							v274 = base.F64_mul(v251, v270)
							v275 = base.F64_add(v263, v274)
							v279 = base.F64_mul(v255, v269)
							v282 = *(*float64)(unsafe.Add(mBase, _consts[1577]))
							v285 = *(*float64)(unsafe.Add(mBase, _consts[1578]))
							v289 = *(*float64)(unsafe.Add(mBase, _consts[1579]))
							v292 = *(*float64)(unsafe.Add(mBase, _consts[1580]))
							v297 = *(*float64)(unsafe.Add(mBase, _consts[1581]))
							v300 = *(*float64)(unsafe.Add(mBase, _consts[1582]))
							v304 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v236), base.F64_add(v255, base.F64_sub(v262, v263))), base.F64_mul(v254, base.F64_add(v269, v270))), base.F64_add(v274, base.F64_sub(v263, v275))), base.F64_mul(base.F64_mul(v255, v279), base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v255, v282), v285)), base.F64_add(base.F64_mul(v255, v289), v292))), base.F64_add(base.F64_mul(v255, v297), v300))))
							v305 = base.F64_add(v275, v304)
							v307 = base.F64_add(v304, base.F64_sub(v275, v305))
							*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v307
							v312 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v305) & v215)
							v313 = base.F64_mul(v217, v312)
							v326 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v313))>>(uint(v220)%64))) & int32(2047)
							v331 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
							if base.Ui32(v326-v331) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v331) {
								v360 = v326
								v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
								v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
								v366 = base.F64_add(base.F64_mul(v313, v362), v365)
								v367 = base.F64_sub(v366, v365)
								v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
								v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
								v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
								v377 = base.F64_mul(v376, v376)
								v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
								v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
								v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
								v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
								v393 = base.I64_reinterpret_f64(v366)
								v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
								v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
								v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
								v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
								v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
								if v360 == int32(0) {
									if v393&int64(2147483648) == int64(0) {
										v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
										v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
									} else {
										v427 = v412 + int64(4602678819172646912)
										v428 = base.F64_reinterpret_i64(v427)
										v429 = base.F64_mul(v428, v404)
										v430 = base.F64_add(v429, v428)
										if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
											v434 = float64(2.2250738585072014e-308)
											v436 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
											v443 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
											if base.F64_lt(v430, float64(0)) != 0 {
												v454 = float64(-1)
											} else {
												v454 = float64(1)
											}
											v455 = base.F64_add(v430, v454)
											v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
											if base.F64_eq(v462, float64(0)) != 0 {
												v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
											} else {
												v465 = v462
											}
											v469 = v465
										} else {
											v469 = v430
										}
										v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
									}
									v491 = v478
								} else {
									v479 = base.F64_reinterpret_i64(v412)
									v491 = base.F64_add(base.F64_mul(v479, v404), v479)
								}
							} else {
								if base.Ui32(v326) < base.Ui32(v331) {
									v342 = base.F64_add(v313, float64(1))
									if v214 != 0 {
										v344 = base.F64_neg(v342)
									} else {
										v344 = v342
									}
									v491 = v344
								} else {
									if base.Ui32(v326) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
										v360 = int32(0)
										v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
										v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
										v366 = base.F64_add(base.F64_mul(v313, v362), v365)
										v367 = base.F64_sub(v366, v365)
										v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
										v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
										v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
										v377 = base.F64_mul(v376, v376)
										v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
										v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
										v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
										v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
										v393 = base.I64_reinterpret_f64(v366)
										v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
										v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
										v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
										v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
										v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
										if v360 == int32(0) {
											if v393&int64(2147483648) == int64(0) {
												v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
												v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
											} else {
												v427 = v412 + int64(4602678819172646912)
												v428 = base.F64_reinterpret_i64(v427)
												v429 = base.F64_mul(v428, v404)
												v430 = base.F64_add(v429, v428)
												if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
													v434 = float64(2.2250738585072014e-308)
													v436 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
													v443 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
													if base.F64_lt(v430, float64(0)) != 0 {
														v454 = float64(-1)
													} else {
														v454 = float64(1)
													}
													v455 = base.F64_add(v430, v454)
													v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
													if base.F64_eq(v462, float64(0)) != 0 {
														v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
													} else {
														v465 = v462
													}
													v469 = v465
												} else {
													v469 = v430
												}
												v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
											}
											v491 = v478
										} else {
											v479 = base.F64_reinterpret_i64(v412)
											v491 = base.F64_add(base.F64_mul(v479, v404), v479)
										}
									} else {
										if base.I64_reinterpret_f64(v313) < int64(0) {
											v356 = F___math_xflow(m, v214, float64(1.2882297539194267e-231))
											mBase = m.M
											v491 = v356
										} else {
											v358 = F___math_xflow(m, v214, float64(3.105036184601418e+231))
											mBase = m.M
											v491 = v358
										}
									}
								}
							}
							v495 = v491
						}
					}
				} else {
					v178 = v34
					v179 = v24
					v180 = v11
					if base.Ui32(v32) <= base.Ui32(int32(-129)) {
						if v178 == int64(4607182418800017408) {
							v495 = float64(1)
						} else {
							if base.Ui32(v30) <= base.Ui32(int32(957)) {
								if base.Ui64(int64(4607182418800017408)) < base.Ui64(v178) {
									v191 = l1
								} else {
									v191 = base.F64_neg(l1)
								}
								v495 = base.F64_add(v191, float64(1))
							} else {
								if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v178)) {
									v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
									mBase = m.M
									v495 = v201
								} else {
									v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
									mBase = m.M
									v495 = v204
								}
							}
						}
					} else {
						if v179 != 0 {
							v212 = v178
							v214 = v180
						} else {
							v212 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
							v214 = v180
						}
						v215 = int64(-134217728)
						v217 = base.F64_reinterpret_i64(v33 & v215)
						v219 = v212 - int64(4604531861337669632)
						v220 = int64(52)
						v223 = base.F64_convert_i32_s(base.I32_wrap_i64(v219 >> (uint(v220) % 64)))
						v225 = *(*float64)(unsafe.Add(mBase, _consts[1571]))
						v233 = base.I32_wrap_i64(int64(base.Ui64(v219)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
						v236 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1572])))
						v240 = v212 - v219&int64(-4503599627370496)
						v245 = base.F64_reinterpret_i64((v240 + int64(2147483648)) & int64(-4294967296))
						v248 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1573])))
						v251 = base.F64_add(base.F64_mul(v245, v248), float64(-1))
						v254 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v240), v245), v248)
						v255 = base.F64_add(v251, v254)
						v257 = *(*float64)(unsafe.Add(mBase, _consts[1574]))
						v261 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[1575])))
						v262 = base.F64_add(base.F64_mul(v223, v257), v261)
						v263 = base.F64_add(v255, v262)
						v268 = *(*float64)(unsafe.Add(mBase, _consts[1576]))
						v269 = base.F64_mul(v255, v268)
						v270 = base.F64_mul(v251, v268)
						v274 = base.F64_mul(v251, v270)
						v275 = base.F64_add(v263, v274)
						v279 = base.F64_mul(v255, v269)
						v282 = *(*float64)(unsafe.Add(mBase, _consts[1577]))
						v285 = *(*float64)(unsafe.Add(mBase, _consts[1578]))
						v289 = *(*float64)(unsafe.Add(mBase, _consts[1579]))
						v292 = *(*float64)(unsafe.Add(mBase, _consts[1580]))
						v297 = *(*float64)(unsafe.Add(mBase, _consts[1581]))
						v300 = *(*float64)(unsafe.Add(mBase, _consts[1582]))
						v304 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v236), base.F64_add(v255, base.F64_sub(v262, v263))), base.F64_mul(v254, base.F64_add(v269, v270))), base.F64_add(v274, base.F64_sub(v263, v275))), base.F64_mul(base.F64_mul(v255, v279), base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v279, base.F64_add(base.F64_mul(v255, v282), v285)), base.F64_add(base.F64_mul(v255, v289), v292))), base.F64_add(base.F64_mul(v255, v297), v300))))
						v305 = base.F64_add(v275, v304)
						v307 = base.F64_add(v304, base.F64_sub(v275, v305))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v307
						v312 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v305) & v215)
						v313 = base.F64_mul(v217, v312)
						v326 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v313))>>(uint(v220)%64))) & int32(2047)
						v331 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
						if base.Ui32(v326-v331) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v331) {
							v360 = v326
							v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
							v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
							v366 = base.F64_add(base.F64_mul(v313, v362), v365)
							v367 = base.F64_sub(v366, v365)
							v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
							v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
							v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
							v377 = base.F64_mul(v376, v376)
							v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
							v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
							v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
							v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
							v393 = base.I64_reinterpret_f64(v366)
							v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
							v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
							v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
							v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
							v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
							if v360 == int32(0) {
								if v393&int64(2147483648) == int64(0) {
									v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
									v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
								} else {
									v427 = v412 + int64(4602678819172646912)
									v428 = base.F64_reinterpret_i64(v427)
									v429 = base.F64_mul(v428, v404)
									v430 = base.F64_add(v429, v428)
									if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
										v434 = float64(2.2250738585072014e-308)
										v436 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
										v443 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
										if base.F64_lt(v430, float64(0)) != 0 {
											v454 = float64(-1)
										} else {
											v454 = float64(1)
										}
										v455 = base.F64_add(v430, v454)
										v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
										if base.F64_eq(v462, float64(0)) != 0 {
											v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
										} else {
											v465 = v462
										}
										v469 = v465
									} else {
										v469 = v430
									}
									v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
								}
								v491 = v478
							} else {
								v479 = base.F64_reinterpret_i64(v412)
								v491 = base.F64_add(base.F64_mul(v479, v404), v479)
							}
						} else {
							if base.Ui32(v326) < base.Ui32(v331) {
								v342 = base.F64_add(v313, float64(1))
								if v214 != 0 {
									v344 = base.F64_neg(v342)
								} else {
									v344 = v342
								}
								v491 = v344
							} else {
								if base.Ui32(v326) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
									v360 = int32(0)
									v362 = *(*float64)(unsafe.Add(mBase, _consts[1583]))
									v365 = *(*float64)(unsafe.Add(mBase, _consts[1584]))
									v366 = base.F64_add(base.F64_mul(v313, v362), v365)
									v367 = base.F64_sub(v366, v365)
									v369 = *(*float64)(unsafe.Add(mBase, _consts[1585]))
									v372 = *(*float64)(unsafe.Add(mBase, _consts[1586]))
									v376 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v312), base.F64_mul(l1, base.F64_add(v307, base.F64_sub(v305, v312)))), base.F64_add(base.F64_mul(v367, v369), base.F64_add(base.F64_mul(v367, v372), v313)))
									v377 = base.F64_mul(v376, v376)
									v380 = *(*float64)(unsafe.Add(mBase, _consts[1587]))
									v383 = *(*float64)(unsafe.Add(mBase, _consts[1588]))
									v387 = *(*float64)(unsafe.Add(mBase, _consts[1589]))
									v390 = *(*float64)(unsafe.Add(mBase, _consts[1590]))
									v393 = base.I64_reinterpret_f64(v366)
									v398 = base.I32_wrap_i64(v393) << (uint(int32(4)) % 32) & int32(2032)
									v401 = *(*float64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[620])))
									v404 = base.F64_add(base.F64_mul(base.F64_mul(v377, v377), base.F64_add(base.F64_mul(v376, v380), v383)), base.F64_add(base.F64_mul(v377, base.F64_add(base.F64_mul(v376, v387), v390)), base.F64_add(v401, v376)))
									v407 = *(*int64)(unsafe.Add(mBase, uint32(v398)+uint32(_consts[621])))
									v412 = v407 + (v393+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
									if v360 == int32(0) {
										if v393&int64(2147483648) == int64(0) {
											v421 = base.F64_reinterpret_i64(v412 - int64(4544132024016830464))
											v478 = base.F64_mul(base.F64_add(base.F64_mul(v421, v404), v421), float64(5.486124068793689e+303))
										} else {
											v427 = v412 + int64(4602678819172646912)
											v428 = base.F64_reinterpret_i64(v427)
											v429 = base.F64_mul(v428, v404)
											v430 = base.F64_add(v429, v428)
											if base.F64_lt(base.F64_abs(v430), float64(1)) != 0 {
												v434 = float64(2.2250738585072014e-308)
												v436 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v436-int32(16))+8)) = v434
												v443 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v443-int32(16))+8)) = base.F64_mul(v434, float64(2.2250738585072014e-308))
												if base.F64_lt(v430, float64(0)) != 0 {
													v454 = float64(-1)
												} else {
													v454 = float64(1)
												}
												v455 = base.F64_add(v430, v454)
												v462 = base.F64_sub(base.F64_add(v455, base.F64_add(base.F64_add(v429, base.F64_sub(v428, v430)), base.F64_add(v430, base.F64_sub(v454, v455)))), v454)
												if base.F64_eq(v462, float64(0)) != 0 {
													v465 = base.F64_reinterpret_i64(v427 & int64(-9223372036854775807-1))
												} else {
													v465 = v462
												}
												v469 = v465
											} else {
												v469 = v430
											}
											v478 = base.F64_mul(v469, float64(2.2250738585072014e-308))
										}
										v491 = v478
									} else {
										v479 = base.F64_reinterpret_i64(v412)
										v491 = base.F64_add(base.F64_mul(v479, v404), v479)
									}
								} else {
									if base.I64_reinterpret_f64(v313) < int64(0) {
										v356 = F___math_xflow(m, v214, float64(1.2882297539194267e-231))
										mBase = m.M
										v491 = v356
									} else {
										v358 = F___math_xflow(m, v214, float64(3.105036184601418e+231))
										mBase = m.M
										v491 = v358
									}
								}
							}
						}
						v495 = v491
					}
				}
			}
		}
	}
	m.G0 = v19 + int32(16)
	return v495
}
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v17 = m.Wasi_snapshot_preview1.Fd_pread(m, l0, v8+int32(8), int32(1), l3, v8+int32(4))
	mBase = m.M
	if v17 == int32(0) {
		v24 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v17
		v24 = int32(-1)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	m.G0 = v8 + int32(16)
	if v24 != 0 {
		v30 = int32(-1)
	} else {
		v30 = v25
	}
	return v30
}
func F_printtup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v17 == v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v138 = l1 + int32(40)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v141 < v140 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v21 == v18 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	F_pfree(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v17
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v32
	if v18 <= v32 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	v38 = F_palloc0(m, v18*int32(40))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v38
	v49 = int32(0)
	goto L13
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v60 = v57 + v49*int32(40)
	if v24 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L1
L15:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	F_fmgr_info(m, v117, v60+int32(12))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L28
	}
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v17+int32(88))+v49*int32(100))))
	F_getTypeOutputInfo(m, v110, v60, v60+int32(8))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L27
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v64 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+10)) = uint16(v64)
	v104 = v63 << (uint(int32(4)) % 32)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+v49<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+10)) = uint16(v72)
	v75 = v68 << (uint(int32(4)) % 32)
	switch v72 {
	case 0:
		v104 = v75
		goto L16
	case 1:
		goto L21
	default:
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v17+int32(20))+v49*int32(100))+68))
	v82 = v60 + int32(4)
	F_getTypeBinaryOutputInfo(m, v80, v82, v60+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v115 = v82
	goto L15
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = base.I32_extend16_s(v72)
	F_errmsg(m, int32(478108), v15)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(486754), int32(292), int32(237769))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v115 = v60
	goto L15
L28:
	;
	v123 = v49 + int32(1)
	if v123 != v18 {
		v49 = v123
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	F_slot_getsomeattrs_int(m, l0, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v145 = int32(4470400)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v148
	F_resetStringInfo(m, v138)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = int32(68)
	goto L34
L33:
	;
	goto L32
L34:
	;
	F_enlargeStringInfo(m, v138, int32(2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v159 = int32(8)
	v165 = v18<<(uint(v159)%32) | int32(base.Ui32(v18&int32(65280))>>(uint(v159)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v156+v157))) = uint16(v165)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v156 + int32(2)
	if int32(0) < v18 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v176 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_pq_endmessage_reuse(m, v138)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
	} else {
		goto L72
	}
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v176))))
	if v187 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	v321 = v176 + int32(1)
	if v321 != v18 {
		v176 = v321
		goto L39
	} else {
		goto L71
	}
L42:
	;
	F_enlargeStringInfo(m, v138, int32(4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v176<<(uint(int32(2))%32))))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v209 = v206 + v176*int32(40)
	v211 = v209 + int32(12)
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+10)))
	if v212 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v194))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v193 + int32(4)
	goto L41
L46:
	;
	v215 = F_OutputFunctionCall(m, v211, v205)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v276 = F_SendFunctionCall(m, v211, v205)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L9
	} else {
		goto L68
	}
L49:
	;
	if v215&int32(3) == int32(0) {
		v240 = v215
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_pq_sendcountedtext(m, v138, v215, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L67
	}
L51:
	;
	v273 = v265 - v215
	goto L50
L52:
	;
	v244 = v240
	goto L61
L53:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v224 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v273 = int32(0)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v229 = v215
	goto L57
L57:
	;
	v233 = v229 + int32(1)
	if v233&int32(3) == int32(0) {
		v240 = v233
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v265 = v233
	goto L51
L59:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v238 != 0 {
		v229 = v233
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v253 = int32(-2139062144)
	if (int32(16843008)-v250|v250)&v253 == v253 {
		v244 = v244 + int32(4)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v259 = v244
	goto L64
L63:
	;
	goto L62
L64:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v263 != 0 {
		v259 = v259 + int32(1)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v265 = v259
	goto L51
L66:
	;
	goto L65
L67:
	;
	goto L41
L68:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	F_enlargeStringInfo(m, v138, int32(4))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v285 = int32(2)
	v287 = int32(4)
	v288 = int32(base.Ui32(v278)>>(uint(v285)%32)) - v287
	v289 = int32(24)
	v291 = int32(65280)
	v293 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v282+v283))) = v288<<(uint(v289)%32) | v288&v291<<(uint(v293)%32) | (int32(base.Ui32(v288)>>(uint(v293)%32))&v291 | int32(base.Ui32(v288)>>(uint(v289)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v282 + v287
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	F_pq_sendbytes(m, v138, v276+v287, int32(base.Ui32(v310)>>(uint(v285)%32))-v287)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	goto L41
L71:
	;
	goto L40
L72:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v146
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_MemoryContextReset(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	m.G0 = v15 + int32(16)
	return int32(1)
}
func F_printtup_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = l0 + int32(40)
	F_initStringInfo(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v16 = F_AllocSetContextCreateInternal(m, v11, int32(227851), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v16
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v19 == int32(1) {
				v22 = F_FetchPortalTargetList(m, v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
					F_SendRowDescriptionMessage(m, v7, l2, v22, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				return
			}
		}
	}
}
func F_proclock_hash(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1104]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_get_hash_value(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6 ^ v10<<(uint(int32(4))%32)
	}
}
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v14 = F_palloc(m, int32(128))
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
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v22 = F_pvsnprintf(m, v14, int32(128), l0, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v22) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v14
	v29 = v22
	goto L7
L5:
	;
	v44 = v14
	goto L6
L6:
	;
	m.G0 = v9 + int32(16)
	return v44
L7:
	;
	F_pfree(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v44 = v34
	goto L6
L9:
	;
	v34 = F_palloc(m, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v39 = F_pvsnprintf(m, v34, v29, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(v29) <= base.Ui32(v39) {
		v28 = v34
		v29 = v39
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_pstrdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v63)
	v66 = v60 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v62, v66, v63)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v60 = v52 - l0
	goto L1
L3:
	;
	v31 = v27
	goto L12
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v16 = l0
	goto L8
L8:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v52 = v20
	goto L2
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v46 = v31
	goto L15
L14:
	;
	goto L13
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	if v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	return v75
L21:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v70, l0, v66)
	mBase = m.M
	v75 = v74
	goto L23
L22:
	;
	v75 = v70
	goto L23
L23:
	;
	goto L20
}
func F_pull_exec_paramids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = int32(0)
	if l0 == v3 {
		v24 = v3
		return v24
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(8) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v9 != int32(1) {
				v24 = v3
				return v24
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v14 = F_bms_add_member(m, v12, v13)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					return int32(0)
				}
			}
		} else {
			v22 = F_expression_tree_walker_impl(m, l0, int32(908), l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				return v24
			}
		}
	}
}
func F_pull_up_sublinks_jointree_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(32)
	return v181
L4:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v181 = v19
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v22 - int32(63) {
	case 0:
		goto L8
	case 1:
		goto L10
	case 2:
		goto L11
	default:
		goto L9
	}
L7:
	;
	v167 = F_makeFromExpr(m, v164, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L46
	}
L8:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v155 = F_bms_make_singleton(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L45
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L10:
	;
	v58 = F_palloc(m, int32(40))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 == int32(0) {
		v163 = v4
		v164 = v4
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v163 = v4
	v164 = v4
	goto L7
L14:
	;
	goto L15
L15:
	;
	v34 = v4
	v36 = v4
	v37 = v4
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v46 = F_pull_up_sublinks_jointree_recurse(m, l0, v43, v11+int32(28))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v163 = v51
	v164 = v48
	goto L7
L18:
	;
	v48 = F_lappend(m, v37, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v51 = F_bms_join(m, v36, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v54 = v34 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v54 < v55 {
		v34 = v54
		v36 = v51
		v37 = v48
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+32)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+24)) = v62
	v65 = v58 + int32(16)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
	v68 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v58
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v76 = F_pull_up_sublinks_jointree_recurse(m, l0, v73, v11+int32(28))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v82 = F_pull_up_sublinks_jointree_recurse(m, l0, v79, v11+int32(24))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	switch v85 {
	case 0:
		goto L26
	case 1:
		goto L29
	case 2:
		goto L25
	case 3:
		goto L28
	default:
		goto L27
	}
L25:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v132 = F_bms_join(m, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v123 = F_bms_union(m, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v97 = int32(0)
	v99 = F_pull_up_sublinks_qual_recurse(m, l0, v93, v58+int32(12), v96, v97, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v88 = int32(0)
	v90 = F_pull_up_sublinks_qual_recurse(m, l0, v86, v65, v87, v88, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v90
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v99
	goto L25
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v106
	F_errmsg_internal(m, int32(476536), v11+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(490403), int32(613), int32(354246))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v125 = int32(0)
	v127 = F_pull_up_sublinks_qual_recurse(m, l0, v118, v11+int32(20), v123, v125, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v127
	goto L25
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v58)+36))
	if v135 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v136 = F_bms_add_member(m, v132, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v181 = v139
	goto L3
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v136
	goto L40
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v144
	F_errmsg_internal(m, int32(477341), v11)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(490403), int32(632), int32(354246))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v155
	v181 = l1
	goto L3
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v167
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v173 = int32(0)
	v175 = F_pull_up_sublinks_qual_recurse(m, l0, v170, v11+int32(28), v163, v173, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v163
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v181 = v179
	goto L3
}
func F_pull_up_subqueries_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
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
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int64
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v764 int32
	_ = v764
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1026 int32
	_ = v1026
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	F_check_stack_depth(m)
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v26 - int32(63) {
	case 0:
		goto L13
	case 1:
		goto L9
	case 2:
		goto L12
	default:
		goto L11
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(80)
	return v1066
L8:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v568 = F_copyObjectImpl(m, v47)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L136
	}
L9:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v496 {
	case 0:
		goto L120
	case 1, 4, 5:
		goto L121
	case 2:
		goto L122
	case 3:
		goto L123
	default:
		goto L124
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L117
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L114
	}
L12:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v437 == int32(0) {
		v1066 = l1
		goto L7
	} else {
		goto L108
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(int32(2))%32)-int32(4))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v39 != int32(1) {
		v123 = v39
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v360 != int32(3) {
		v1066 = l1
		goto L7
	} else {
		goto L89
	}
L15:
	;
	v319 = F_palloc0(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L84
	}
L16:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v318 = v313<<(uint(int32(2))%32) + int32(4)
	goto L15
L17:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+52))
	if v229 != 0 {
		goto L70
	} else {
		goto L71
	}
L18:
	;
	if l3 != 0 {
		goto L14
	} else {
		goto L48
	}
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v43 = F_is_simple_subquery(m, l0, v42, v38, l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v97 != int32(1) {
		v123 = v97
		goto L18
	} else {
		goto L37
	}
L21:
	;
	if v43 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	if l3 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v51 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v54 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v61 = v50
	goto L28
L27:
	;
	goto L26
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v70 != int32(65) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v70 != int32(63) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v75 != 0 {
		goto L20
	} else {
		goto L34
	}
L33:
	;
	goto L8
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 != int32(1) {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v61 = v83
	goto L28
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != int32(67) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v104 != int32(1) {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+144))
	if v107 == int32(0) {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v100)+124))
	if v110 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+128))
	if v111 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+132))
	if v112 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v100)+140))
	if v113 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	if v114 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v116 = F_is_simple_union_all_recurse(m, v107, v100, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v116 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v123 = v118
	goto L18
L48:
	;
	if l2 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	if v123 != int32(5) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	if v134 == int32(0) {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v137 != int32(1) {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v140 = F_expression_returns_set(m, v134)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v140 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v143 = F_contain_volatile_functions(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v143 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+52))
	if v146 == int32(0) {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v149 != int32(1) {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v38 != v153 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v156 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = F_copyObjectImpl(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v213 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v145 + int32(39)
	if v207 != 0 {
		goto L16
	} else {
		goto L69
	}
L61:
	;
	if v160 == int32(0) {
		v207 = v156
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v164 <= int32(0) {
		v207 = v156
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v173 = int32(1)
	v174 = int32(0)
	v176 = v156
	goto L64
L64:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v174<<(uint(int32(2))%32))))
	v188 = int32(0)
	v190 = F_makeTargetEntry(m, v186, base.I32_extend16_s(v173), v188, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v207 = v192
	goto L60
L66:
	;
	v192 = F_lappend(m, v176, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v194 = int32(1)
	v197 = v174 + v194
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v197 < v198 {
		v173 = v173 + v194
		v174 = v197
		v176 = v192
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v318 = int32(4)
	goto L15
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v232 = v230
	goto L72
L71:
	;
	v232 = int32(0)
	goto L72
L72:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v227)+52))
	v235 = F_copyObjectImpl(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v237 = m.G0
	v238 = int32(16)
	v239 = v237 - v238
	m.G0 = v239
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+8)) = int32(-1)
	v249 = F_range_table_walker_impl(m, v235, int32(1049), v239+int32(8), v238)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v239 + int32(16)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if v254 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v227)+56))
	F_CombineRangeTables(m, v300+int32(52), v300+int32(56), v235, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L82
	}
L76:
	;
	if v235 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v259 <= int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v267 = int32(0)
	goto L79
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v267<<(uint(int32(2))%32))))
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+124)) = uint8(v281)
	v284 = v267 + v281
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v284 < v285 {
		v267 = v284
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L75
L81:
	;
	goto L80
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v227)+144))
	F_pull_up_union_leaf_queries(m, v308, l0, v233, v227, v232)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+20)) = uint8(v311)
	v1066 = l1
	goto L7
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v319
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v328 = F_palloc0(m, int32(136))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+12)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(101)
	v336 = F_makeAlias(m, int32(644142), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+8)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v328
	v344 = F_list_make1_impl(m, int32(1), v16+int32(12))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+52)) = v344
	v1066 = l1
	goto L7
L89:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+72)))
	if v363 != 0 {
		v1066 = l1
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v364 == int32(0) {
		v1066 = l1
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v367 != int32(1) {
		v1066 = l1
		goto L7
	} else {
		goto L92
	}
L92:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	if v373 != int32(7) {
		v1066 = l1
		goto L7
	} else {
		goto L93
	}
L93:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	if v376 != int32(1) {
		v1066 = l1
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	if v379 != 0 {
		v1066 = l1
		goto L7
	} else {
		goto L95
	}
L95:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v385 = F_get_expr_result_type(m, v372, v16+int32(68), v16-int32(-64))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v385 != 0 {
		v1066 = l1
		goto L7
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	v390 = int32(0)
	v392 = F_makeTargetEntry(m, v388, int32(1), v390, v390)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v392
	v399 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v380 + int32(39)
	v404 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v399
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v410
	if v399 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v420 = v414<<(uint(int32(2))%32) + int32(4)
	goto L102
L101:
	;
	v420 = int32(4)
	goto L102
L102:
	;
	v421 = F_palloc0(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v421
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v380)+108))
	if v424 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	goto L106
L105:
	;
	goto L106
L106:
	;
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), l3)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v431 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)) = uint8(v431)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = int32(8)
	v1066 = l1
	goto L7
L108:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v440 <= int32(0) {
		v1066 = l1
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v447 = int32(0)
	goto L110
L110:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v459 = v456 + v447<<(uint(int32(2))%32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v462 = F_pull_up_subqueries_recurse(m, l0, v460, l2, int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v1066 = l1
	goto L7
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v462
	v466 = v447 + int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v466 < v467 {
		v447 = v466
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v473
	F_errmsg_internal(m, int32(477341), v16)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(490403), int32(1255), int32(354152))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errmsg_internal(m, int32(113082), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(490403), int32(2222), int32(299900))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v545 = F_pull_up_subqueries_recurse(m, l0, v543, l2, int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L134
	}
L121:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v535 = F_pull_up_subqueries_recurse(m, l0, v533, l1, int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v525 = F_pull_up_subqueries_recurse(m, l0, v523, l1, int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L130
	}
L123:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v515 = F_pull_up_subqueries_recurse(m, l0, v513, l1, int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L128
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v501
	F_errmsg_internal(m, int32(476536), v16+int32(16))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(490403), int32(1249), int32(354152))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v515
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v520 = F_pull_up_subqueries_recurse(m, l0, v518, l1, int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v520
	v1066 = l1
	goto L7
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v525
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v530 = F_pull_up_subqueries_recurse(m, l0, v528, l1, int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v530
	v1066 = l1
	goto L7
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v535
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v540 = F_pull_up_subqueries_recurse(m, l0, v538, l1, int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v540
	v1066 = l1
	goto L7
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v545
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v550 = F_pull_up_subqueries_recurse(m, l0, v548, l2, int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v550
	v1066 = l1
	goto L7
L136:
	;
	v571 = F_palloc0(m, int32(384))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = int32(267)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+8)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+12)) = v578
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v581 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v571)+20)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v571)+16)) = v580
	v585 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v586 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v571)+321)) = uint16(v586)
	*(*int32)(unsafe.Add(mBase, uint32(v571)+312)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v571)+280)) = v585
	*(*int64)(unsafe.Add(mBase, uint32(v571)+72)) = v581
	*(*int64)(unsafe.Add(mBase, uint32(v571)+80)) = v581
	*(*int64)(unsafe.Add(mBase, uint32(v571)+85)) = v581
	*(*int64)(unsafe.Add(mBase, uint32(v571)+116)) = v581
	*(*int64)(unsafe.Add(mBase, uint32(v571)+124)) = v581
	*(*int64)(unsafe.Add(mBase, uint32(v571)+132)) = v581
	v608 = F__emscripten_memset_bulkmem(m, v571+int32(192), base.I32_extend8_s(v586), int32(88))
	mBase = m.M
	goto L138
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v571)+344)) = int64(4294967295)
	F_replace_empty_jointree(m, v568)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+39)))
	if v613 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+60))
	v620 = F_pull_up_sublinks_jointree_recurse(m, v571, v617, v16+int32(24))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+52))
	if v640 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v622 != int32(65) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v620
	v630 = F_list_make1_impl(m, int32(1), v16+int32(4))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	v635 = v620
	goto L146
L146:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+60)) = v635
	goto L142
L147:
	;
	v633 = F_makeFromExpr(m, v630, int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v635 = v633
	goto L146
L149:
	;
	v701 = F_expand_virtual_generated_columns(m, v571)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L160
	}
L150:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v643 <= int32(0) {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v652 = int32(0)
	goto L152
L152:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660+v652<<(uint(int32(2))%32))))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v664)+12))
	if v665 != int32(3) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L149
L154:
	;
	v685 = v652 + int32(1)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v685 < v686 {
		v652 = v685
		goto L152
	} else {
		goto L159
	}
L155:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v664)+68))
	v669 = F_eval_const_expressions(m, v571, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664)+68)) = v669
	v672 = F_inline_set_returning_function(m, v571, v664)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	if v672 == int32(0) {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v676 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+72)) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+40)) = uint8(v676)
	*(*int32)(unsafe.Add(mBase, uint32(v664)+36)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v664)+12)) = int32(1)
	goto L154
L159:
	;
	goto L153
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v701
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)+60))
	v705 = int32(0)
	v707 = F_pull_up_subqueries_recurse(m, v571, v704, v705, v705)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v709)+60)) = v707
	v711 = F_is_simple_subquery(m, l0, v701, v38, l2)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v711 == int32(0) {
		v1066 = l1
		goto L7
	} else {
		goto L163
	}
L163:
	;
	if l3 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v701)+60))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	if v716 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v701)+76))
	v782 = F_flatten_join_alias_vars(m, v571, v780, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L183
	}
L167:
	;
	if v764 == int32(0) {
		v1066 = l1
		goto L7
	} else {
		goto L182
	}
L168:
	;
	v722 = v715
	goto L172
L169:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v715)+8))
	if v717 != 0 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v764 = int32(1)
	goto L167
L171:
	;
	v764 = v750
	goto L167
L172:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if v733 != int32(65) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v750 = int32(0)
	goto L171
L174:
	;
	goto L173
L175:
	;
	if v733 == int32(63) {
		v750 = int32(1)
		goto L171
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v722)+8))
	if v738 != 0 {
		goto L174
	} else {
		goto L179
	}
L178:
	;
	goto L174
L179:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v739 == int32(0) {
		goto L174
	} else {
		goto L180
	}
L180:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v739)+4))
	if v742 != int32(1) {
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v739)+12))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v722 = v746
	goto L172
L182:
	;
	goto L166
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701)+76)) = v782
	v785 = int32(0)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v567)+52))
	if v787 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	v789 = v788
	goto L186
L185:
	;
	v789 = v785
	goto L186
L186:
	;
	F_OffsetVarNodes(m, v701, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v571)+128))
	F_OffsetVarNodes(m, v792, v789)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_IncrementVarSublevelsUp(m, v701, int32(-1), int32(1))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v571)+128))
	F_IncrementVarSublevelsUp(m, v799, int32(-1), int32(1))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v701)+76))
	v806 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if v811 == int32(1) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v701)+60))
	v815 = int32(1)
	v817 = F_get_relids_in_jointree(m, v814, v815, v815)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	v839 = v785
	v840 = v806
	v841 = v805
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v567 + int32(39)
	if v841 != 0 {
		goto L201
	} else {
		goto L202
	}
L194:
	;
	v821 = F_palloc(m, int32(8))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v567)+52))
	if v823 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+4))
	v825 = v824
	goto L198
L197:
	;
	v825 = int32(0)
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v821)+4)) = v825
	v831 = F_palloc0(m, v825<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = v831
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v567)+60))
	F_get_nullingrels_recurse(m, v834, int32(0), v821)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v701)+76))
	v839 = v817
	v840 = v821
	v841 = v838
	goto L193
L201:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v841)+4))
	v857 = v851<<(uint(int32(2))%32) + int32(4)
	goto L203
L202:
	;
	v857 = int32(4)
	goto L203
L203:
	;
	v858 = F_palloc0(m, v857)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v858
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v567)+108))
	if v861 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	goto L207
L206:
	;
	goto L207
L207:
	;
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), l3)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if v870 != int32(1) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v701)+52))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v701)+56))
	F_CombineRangeTables(m, v567+int32(52), v567+int32(56), v923, v924)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L220
	}
L210:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v701)+52))
	if v873 == int32(0) {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v876 <= int32(0) {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	v885 = int32(0)
	goto L213
L213:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v893+v885<<(uint(int32(2))%32))))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+12))
	switch v898 {
	case 0:
		goto L217
	case 1, 3, 4, 5:
		goto L216
	default:
		goto L215
	}
L214:
	;
	goto L209
L215:
	;
	v905 = v885 + int32(1)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v905 < v906 {
		v885 = v905
		goto L213
	} else {
		goto L219
	}
L216:
	;
	v902 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v897)+124)) = uint8(v902)
	goto L215
L217:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v897)+32))
	if v899 == int32(0) {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	goto L214
L220:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v567)+140))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v701)+140))
	v929 = F_list_concat(m, v927, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+140)) = v929
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)+68))
	if v933 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v571)+128))
	v1042 = F_list_concat(m, v1040, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L248
	}
L223:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v701)+60))
	v939 = F_get_relids_in_jointree(m, v936, int32(1), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v934 != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1040 = int32(0)
	goto L222
L226:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v941)+68))
	if v942 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v939
	v944 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v566
	v951 = F_query_or_expression_tree_walker_impl(m, v567, int32(851), v16+int32(68), v944)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v953 = int32(0)
	v954 = m.G0
	v956 = v954 - int32(16)
	m.G0 = v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v958 == v953 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L229
L231:
	;
	m.G0 = v956 + int32(16)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1040 = v1026
	goto L222
L232:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v961 <= int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v968 = int32(-1)
	v974 = v953
	goto L234
L234:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v958)+12))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v978+v974<<(uint(int32(2))%32))))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+8))
	if v566 == v983 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	goto L231
L236:
	;
	if v968 < int32(0) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v991 = v968
	goto L238
L238:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)+68))
	if v993 != 0 {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	v987 = F_bms_singleton_member(m, v939)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	v989 = v968
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v982)+8)) = v989
	v991 = v989
	goto L238
L242:
	;
	v989 = v987
	goto L241
L243:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v982)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v956)+12)) = v939
	v996 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v956)+8)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v956)+4)) = v566
	v1003 = F_query_or_expression_tree_walker_impl(m, v994, int32(851), v956+int32(4), v996)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1007 = v974 + int32(1)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v1007 < v1008 {
		v968 = v991
		v974 = v1007
		goto L234
	} else {
		goto L247
	}
L246:
	;
	goto L245
L247:
	;
	goto L235
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = int32(0)
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+39)))
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+39)))
	v1049 = v1047 | v1048
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+39)) = uint8(v1049)
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+44)))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+44)))
	v1053 = v1051 | v1052
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+44)) = uint8(v1053)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v701)+60))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+8))
	if v1056 != 0 {
		v1066 = v1055
		goto L7
	} else {
		goto L249
	}
L249:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+4))
	if v1057 == int32(0) {
		v1066 = v1055
		goto L7
	} else {
		goto L250
	}
L250:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	if v1060 != int32(1) {
		v1066 = v1055
		goto L7
	} else {
		goto L251
	}
L251:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	v1066 = v1064
	goto L7
}
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		v14 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(12), l2, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v14 {
				v22 = v14
				v24 = F_palloc0(m, int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
					if v22 != 0 {
						v32 = F_palloc(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
							v42 = v35
							m.G0 = v9 + int32(16)
							return v42
						}
					} else {
						v34 = int32(0)
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
						v42 = v35
						m.G0 = v9 + int32(16)
						return v42
					}
				}
			} else {
				v42 = v14
				m.G0 = v9 + int32(16)
				return v42
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		v22 = int32(0)
		v24 = F_palloc0(m, int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
			if v22 != 0 {
				v32 = F_palloc(m, v22)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v42 = v35
					m.G0 = v9 + int32(16)
					return v42
				}
			} else {
				v34 = int32(0)
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v42 = v35
				m.G0 = v9 + int32(16)
				return v42
			}
		}
	}
}
func F_push_into_mbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	if int32(0) < l3 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
		if v9 == int32(1) {
			F_px_debug(m, int32(343498), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(-12)
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v20) < base.Ui32(v21+l3) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = v20 - v24 + (l3+int32(32767))&int32(2147467264)
				v31 = F_repalloc(m, v24, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31 + v30
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v39 = v31 + (v37 - v35)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31 + (v41 - v35)
					v45 = v39
					if l3 != 0 {
						v48 = F__emscripten_memcpy_bulkmem(m, v45, l2, l3)
						mBase = m.M
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50 + l3
					return int32(0)
				}
			} else {
				v45 = v21
				if l3 != 0 {
					v48 = F__emscripten_memcpy_bulkmem(m, v45, l2, l3)
					mBase = m.M
				} else {
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50 + l3
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_pushval_morph(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4)
	v23 = F_palloc(m, int32(64))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_parsetext(m, v26, v16, l2, l3)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if int32(0) < v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v16 + int32(16)
	return
L5:
	;
	v36 = int32(0)
	v40 = v29
	v41 = v7
	v43 = v7
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pfree(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	if v95 <= v36 {
		v223 = v36
		v227 = v95
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v36<<(uint(int32(4))%32))+8)))
	v95 = v40
	v96 = v52
	v98 = v43
	goto L10
L12:
	;
	goto L13
L13:
	;
	v54 = v41 + int32(1)
	v56 = v36 << (uint(int32(4)) % 32)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v56)+8)))
	if base.Ui32(v58) <= base.Ui32(v54) {
		v95 = v40
		v96 = v58
		v98 = v43
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v62 = v54
	v70 = v43
	goto L15
L15:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v95 = v87
	v96 = v85
	v98 = v80
	goto L10
L17:
	;
	if v70 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v75, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v79 = int32(1)
	v80 = v70 + v79
	v82 = v62 + v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+v56)+8)))
	if base.Ui32(v82) < base.Ui32(v85) {
		v62 = v82
		v70 = v80
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L16
L23:
	;
	if v98 != 0 {
		goto L50
	} else {
		goto L51
	}
L24:
	;
	v106 = v36
	v110 = v95
	v114 = int32(0)
	goto L25
L25:
	;
	v117 = v106 << (uint(int32(4)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v119 = v117 + v118
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+8)))
	if v96 != v120 {
		v223 = v106
		v227 = v110
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v223 = v201
	v227 = v216
	goto L23
L27:
	;
	if v110 <= v106 {
		v201 = v106
		v205 = v110
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v114 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+2)))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119))))
	F_pushValue(m, l1, v124, v125, l4, l5|int32(base.Ui32(v126&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134+v117)+12))
	F_pfree(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v140 = v106 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v141 <= v140 {
		v201 = v140
		v205 = v141
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v147 = v144 + v140<<(uint(int32(4))%32)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+8)))
	if v96 != v148 {
		v201 = v140
		v205 = v141
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v152 = v147
	v153 = v140
	v157 = v141
	v159 = int32(1)
	goto L34
L34:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)))
	if v123 != v163 {
		v201 = v153
		v205 = v157
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v201 = v187
	v205 = v188
	goto L28
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	F_pushValue(m, l1, v165, v166, l4, l5|int32(base.Ui32(v167&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v153<<(uint(int32(4))%32))+12))
	F_pfree(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pushOperator(m, l1, int32(2), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v187 = v153 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v188 <= v187 {
		v201 = v187
		v205 = v188
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v195 = v192 + v187<<(uint(int32(4))%32)
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+8)))
	if v96 == v196 {
		v152 = v195
		v153 = v187
		v157 = v188
		v159 = v159 + int32(1)
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	F_pushOperator(m, l1, int32(3), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v216 = v205
	goto L47
L47:
	;
	if v201 < v216 {
		v106 = v201
		v110 = v216
		v114 = v114 + int32(1)
		goto L25
	} else {
		goto L49
	}
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v216 = v215
	goto L47
L49:
	;
	goto L26
L50:
	;
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v233, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v238 = v227
	goto L52
L52:
	;
	if v223 < v238 {
		v36 = v223
		v40 = v238
		v41 = v96
		v43 = v98 + int32(1)
		goto L8
	} else {
		goto L54
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v238 = v237
	goto L52
L54:
	;
	goto L9
L55:
	;
	goto L4
L56:
	;
	goto L4
}
func F_puts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	v5 = F_fputs(m, l0, int32(4364616))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v5 < int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[1591]))
			if v10 == int32(10) {
				F___overflow(m, int32(4364616), int32(10))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					return
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[1592]))
				v16 = *(*int32)(unsafe.Add(mBase, _consts[1593]))
				if v14 == v16 {
					F___overflow(m, int32(4364616), int32(10))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[1592])) = v14 + int32(1)
					v22 = int32(10)
					*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v22)
					return
				}
			}
		}
	}
}
func F_pvsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_pg_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v10 {
			if base.Ui32(l1) <= base.Ui32(v10) {
				if base.Ui32(int32(1073741823)) <= base.Ui32(v10) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(13796), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(489331), int32(140), int32(332368))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
					v21 = v10 + int32(1)
					m.G0 = v8 + int32(16)
					return v21
				}
			} else {
				v21 = v10
				m.G0 = v8 + int32(16)
				return v21
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
				F_errmsg_internal(m, int32(686717), v8)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(489331), int32(113), int32(332368))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
