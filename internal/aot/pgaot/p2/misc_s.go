package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_SanityCheckBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v10&int32(1) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 - int32(-64)
	return v192
L2:
	;
	v15 = int32(0)
	v17 = F_errstart(m, l1, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v10&int32(2) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return int32(0)
L6:
	;
	if v17 == int32(0) {
		v192 = v15
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_0), v8)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(647), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v192 = v15
	goto L1
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if base.Ui32(v60-int32(86400001)) <= base.Ui32(int32(-86400003)) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v39 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v40 = int32(0)
	v42 = F_errstart(m, l1, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v42 == int32(0) {
		v192 = v40
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_3), v6+int32(-16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(658), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v192 = v40
	goto L1
L19:
	;
	v65 = int32(0)
	v67 = F_errstart(m, l1, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v87 = int32(0)
	if base.B2i32(v10&int32(16) == v87)|base.B2i32(v60 == int32(-1)) == v87 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	if v67 == int32(0) {
		v192 = v65
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_4), v6+int32(-48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(672), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v192 = v65
	goto L1
L27:
	;
	v94 = int32(0)
	v96 = F_errstart(m, l1, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v114 = int32(1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v115 != 0 {
		v192 = v114
		goto L1
	} else {
		goto L35
	}
L30:
	;
	if v96 == int32(0) {
		v192 = v94
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_5), v6+int32(-32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(687), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v192 = v94
	goto L1
L35:
	;
	v117 = l0 + int32(96)
	if (l0^v117)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v192 = v114
	goto L1
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v171)
	if v171&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v170 = l0
	v171 = v123
	v172 = v117
	goto L38
L40:
	;
	goto L41
L41:
	;
	if l0&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = l0
	v129 = v117
	goto L45
L43:
	;
	v141 = l0
	v143 = v117
	goto L44
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v148 = int32(-2139062144)
	if (int32(16843008)-v145|v145)&v148 != v148 {
		v170 = v141
		v171 = v145
		v172 = v143
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v130)
	if v130 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v141 = v137
	v143 = v135
	goto L44
L47:
	;
	v134 = int32(1)
	v135 = v129 + v134
	v137 = v127 + v134
	if v137&int32(3) != 0 {
		v127 = v137
		v129 = v135
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v153 = v141
	v154 = v145
	v155 = v143
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v154
	v157 = int32(4)
	v158 = v155 + v157
	v160 = v153 + v157
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v165 = int32(-2139062144)
	if (int32(16843008)-v162|v162)&v165 == v165 {
		v153 = v160
		v154 = v162
		v155 = v158
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v170 = v160
	v171 = v162
	v172 = v158
	goto L38
L52:
	;
	goto L51
L53:
	;
	v179 = v170
	v181 = v172
	goto L54
L54:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)) = uint8(v182)
	v184 = int32(1)
	if v182 != 0 {
		v179 = v179 + v184
		v181 = v181 + v184
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L37
L56:
	;
	goto L55
}
func F_ScanKeywordLookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v7 = int32(-1)
	v8 = F_strlen(m, l0)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v9) < base.Ui32(v8) {
		v56 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v56
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l0, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v12 < int32(0) {
		v56 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v18 <= v12 {
		v56 = v7
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v12<<(uint(int32(1))%32)))))
	v27 = l0
	v28 = v20 + v25
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v35 = int32(1)
	if base.Ui32((v33-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v47 = v33 | int32(32)
	goto L14
L13:
	;
	v47 = v33
	goto L14
L14:
	;
	if v34 == v47 {
		v27 = v27 + v35
		v28 = v28 + v35
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v56 = v7
	goto L1
L16:
	;
	v51 = int32(-1)
	goto L18
L17:
	;
	v51 = v12
	goto L18
L18:
	;
	v56 = v51
	goto L1
}
func F_SetLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[0]))
	if v12 == v8 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[1]))
	if v19 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v42 = F_pgmem_kill(m, v8, int32(23))
	mBase = m.M
	goto L1
L8:
	;
	m.G0 = v16 + int32(16)
	return
L9:
	;
	v22 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)) = uint8(v22)
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[2]))
	v30 = F_write(m, v26, v16+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v30 {
		goto L8
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[3]))
	if v34 == int32(27) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_SetVarReturningType_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(0)
	if l0 == v3 {
		v41 = v3
		return v41
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(67) {
			if v7 != int32(6) {
				v38 = F_expression_tree_walker_impl(m, l0, int32(1057), l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = v38
					return v41
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v12 != v13 {
					v41 = v3
					return v41
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v15 != v16 {
						v41 = v3
						return v41
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
						return int32(0)
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 + int32(1)
			v28 = F_query_tree_walker_impl(m, l0, int32(1057), l1, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v32 - int32(1)
				return v28
			}
		}
	}
}
func F_SetupHistoricSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_SetupHistoricSnapshot[0])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SetupHistoricSnapshot[1])) = l0
	return
}
func F_ShowUsage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	v9 = m.G0
	v11 = v9 - int32(368)
	m.G0 = v11
	v14 = v11 + int32(184)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1)
	v24 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, v11+int32(336))
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+344))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
	if v29 < v31 {
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+336))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+336)) = v33 - int64(1)
		v39 = v29 + int32(_a_F_ShowUsage_0)
	} else {
		v39 = v29
	}
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+208))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)+184))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
	if v28 < v43 {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v28 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+184)) = v41 - int64(1)
	} else {
	}
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
	if v40 < v53 {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v40 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+200)) = v51 - int64(1)
	} else {
	}
	v62 = v11 + int32(352)
	F_initStringInfo(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		return
	} else {
		F_appendStringInfoString(m, v62, int32(_a_F_ShowUsage_1))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			v68 = *(*int64)(unsafe.Add(mBase, uint32(v11)+336))
			v70 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[3]))
			v71 = v68 - v70
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+176)) = uint32(v71)
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+180)) = v39 - v74
			v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+184))
			v79 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[4]))
			v80 = v77 - v79
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+160)) = uint32(v80)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
			v84 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v82 - v84
			v87 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
			v89 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[5]))
			v90 = v87 - v89
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+168)) = uint32(v90)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+208))
			v94 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+172)) = v92 - v94
			F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_2), v11+int32(160))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v40
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+152)) = uint32(v51)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+148)) = v28
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+144)) = uint32(v41)
				F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_3), v11+int32(144))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v111
					F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_4), v11+int32(128))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+244))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = v118
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+248))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = v120
						v123 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[6]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v118 - v123
						v127 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v120 - v127
						F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_5), v11+int32(112))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+240))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v135
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[8]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v135 - v138
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+236))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v141
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v143
							v146 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[9]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v141 - v146
							v150 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[10]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v143 - v150
							F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_6), v11+int32(80))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+256))
								*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v160
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+252))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v162
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v11)+260))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v164
								v167 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[11]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v164 - v167
								v171 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[12]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v160 - v171
								v175 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[13]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v162 - v175
								F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_7), v11+int32(48))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return
								} else {
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+264))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v183
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+268))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v185
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[14]))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v183 - v188
									v192 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[15]))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v185 - v192
									F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_8), v11+int32(32))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return
									} else {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
										v201 = *(*int32)(unsafe.Add(mBase, uint32(v11)+356))
										v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v201-int32(1)))))
										if v205 == int32(10) {
											v209 = v201 - int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+356)) = v209
											v212 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v209+v200))) = uint8(v212)
										} else {
										}
										v217 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return
										} else {
											if v217 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
												F_errmsg_internal(m, int32(_a_F_ShowUsage_9), v11+int32(16))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v225
													F_errdetail_internal(m, int32(_a_F_ShowUsage_9), v11)
													mBase = m.M
													v229 = m.ExcPending
													if v229 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ShowUsage_10), int32(_a_F_ShowUsage_11), int32(_a_F_ShowUsage_12))
														mBase = m.M
														v234 = m.ExcPending
														if v234 != 0 {
															return
														} else {
															v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
															F_pfree(m, v235)
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
																return
															} else {
																m.G0 = v11 + int32(368)
																return
															}
														}
													}
												}
											} else {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
												F_pfree(m, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													m.G0 = v11 + int32(368)
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
	}
}
func F_SignalHandlerForCrashExit(m *base.Module, l0 int32) {
	F__Exit(m, int32(2))
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SplitDirectoriesString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v9 = l0
	goto L1
L1:
	;
	v17 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	goto L3
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v17 == int32(32))|base.B2i32(base.Ui32((v17-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v9 = v9 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	return int32(1)
L6:
	;
	v33 = v9
	v34 = v27
	goto L7
L7:
	;
	if v34 != int32(34) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v167)
	v169 = F_strlen(m, v109)
	mBase = m.M
	if base.Ui32(int32(1024)) <= base.Ui32(v169) {
		goto L55
	} else {
		goto L56
	}
L10:
	;
	return int32(0)
L11:
	;
	v112 = v106
	goto L43
L12:
	;
	if v34 == int32(0) {
		v66 = v33
		v71 = v33
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v74 = v33 + int32(1)
	v75 = int32(34)
	v76 = F___strchrnul(m, v74, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 == v75 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	if v33 == v71 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	if v34 == int32(44) {
		v66 = v33
		v71 = v33
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = v33
	v46 = v34
	v47 = v33
	goto L18
L18:
	;
	v49 = v44 + int32(1)
	v50 = base.I32_extend8_s(v46)
	goto L20
L19:
	;
	v66 = v49
	v71 = v60
	goto L15
L20:
	;
	if base.B2i32(v50 == int32(32))|base.B2i32(base.Ui32((v50-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v60 = v47
	goto L23
L22:
	;
	v60 = v49
	goto L23
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v61 == int32(0) {
		v66 = v49
		v71 = v60
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v61 != int32(44) {
		v44 = v49
		v46 = v61
		v47 = v60
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v106 = v66
	v109 = v33
	v111 = v71
	goto L11
L27:
	;
	if v82 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L28:
	;
	v82 = v76
	goto L30
L29:
	;
	v82 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v90 = v82
	goto L32
L32:
	;
	v92 = v90 + int32(1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v93 != int32(34) {
		v106 = v92
		v109 = v74
		v111 = v90
		goto L11
	} else {
		goto L34
	}
L33:
	;
	goto L10
L34:
	;
	v96 = F_strlen(m, v90)
	mBase = m.M
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryCopy(m, v90, v92, v96)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v98 = int32(34)
	v99 = F___strchrnul(m, v92, v98)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v101 == v98 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v105 != 0 {
		v90 = v105
		goto L32
	} else {
		goto L42
	}
L39:
	;
	v105 = v99
	goto L41
L40:
	;
	v105 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L33
L43:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	goto L45
L44:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v130 == int32(44) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if base.B2i32(v120 == int32(32))|base.B2i32(base.Ui32((v120-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v112 = v112 + int32(1)
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v135 = v112
	goto L50
L48:
	;
	goto L49
L49:
	;
	if v130 == int32(0) {
		v163 = v112
		goto L9
	} else {
		goto L54
	}
L50:
	;
	v140 = v135 + int32(1)
	v141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135)+1)))
	goto L52
L52:
	;
	if base.B2i32(v141 == int32(32))|base.B2i32(base.Ui32((v141-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v135 = v140
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v163 = v140
	goto L9
L54:
	;
	goto L10
L55:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+1023)) = uint8(v172)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v174 = F_pstrdup(m, v109)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	return int32(0)
L59:
	;
	F_canonicalize_path_enc(m, v174)
	mBase = m.M
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v180 = F_lappend(m, v179, v174)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v180
	if v130 != int32(44) {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v33 = v163
	v34 = v185
	goto L7
}
func F_StrategyNotifyBgWriter(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(0), int32(1))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
		F_s_lock(m, v9, int32(_a_F_StrategyNotifyBgWriter_0), int32(438), int32(_a_F_StrategyNotifyBgWriter_1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
			v18 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v18))
			return
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
		v18 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v18))
		return
	}
}
func F___shgetc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
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
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = v10 + base.I64_extend_i32_s(v11-v12)
	if base.B2i32(v7 != int64(0))&base.B2i32(v7 <= v15) == int32(0) {
		v20 = F___uflow(m, l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v20 {
				v41 = v15 + int64(1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v44 == int64(0) {
					v53 = v43
				} else {
					v47 = v44 - v41
					if base.I64_extend_i32_s(v43-v42) <= v47 {
						v53 = v43
					} else {
						v53 = v42 + base.I32_wrap_i64(v47)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v53
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v41 + base.I64_extend_i32_s(v56-v42)
				if base.Ui32(v42) <= base.Ui32(v56) {
					*(*uint8)(unsafe.Add(mBase, uint32(v42-int32(1)))) = uint8(v20)
				} else {
				}
				return v20
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v28 = v27
				v29 = v26
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v15 + base.I64_extend_i32_s(v29-v28)
				return int32(-1)
			}
		}
	} else {
		v28 = v11
		v29 = v12
		*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v28
		*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v15 + base.I64_extend_i32_s(v29-v28)
		return int32(-1)
	}
}
func F___stdio_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3 = m.Wasi_snapshot_preview1.Fd_close(m, v2)
	mBase = m.M
	if v3 == int32(0) {
		v10 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___stdio_close[0])) = v3
		v10 = int32(-1)
	}
	return v10
}
func F_scanGetCandidate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
	v13 = l1 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = v14
	goto L2
L1:
	;
	return base.B2i32(base.Ui32(v22) <= base.Ui32(v49&int32(_a_F_scanGetCandidate_0)))
L2:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v77 = v40 + int32(20)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v22<<(uint(int32(2))%32))))
	v84 = v40 + v81&int32(_a_F_scanGetCandidate_1)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v89)+6)))
	if v91&int32(32) != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v41) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_scanGetCandidate[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v17^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L4
L6:
	;
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_scanGetCandidate[1]))
	v40 = v34 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v49 = int32(base.Ui32(v41+int32(_a_F_scanGetCandidate_2)) >> (uint(int32(2)) % 32))
	goto L10
L9:
	;
	v49 = int32(0)
	goto L10
L10:
	;
	if base.Ui32(v49&int32(_a_F_scanGetCandidate_0)) < base.Ui32(v22) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40+v53)))
	if v55 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	goto L3
L14:
	;
	F_UnlockReleaseBuffer(m, v17)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = F_ReadBuffer(m, v64, v55)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L19:
	;
	F_LockBuffer(m, v65, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_UnlockReleaseBuffer(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v73 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v73)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v65
	v17 = v65
	goto L2
L22:
	;
	v95 = v22 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v95)
	v97 = int32(_a_F_scanGetCandidate_0)
	v98 = v95 & v97
	if base.Ui32(v49&v97) < base.Ui32(v98) {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v169 = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v169)
	goto L1
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v77+v98<<(uint(int32(2))%32))))
	v108 = v40 + v105&int32(_a_F_scanGetCandidate_1)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v111 = int32(16)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+2)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	if v109|v110<<(uint(v111)%32) == v114|v115<<(uint(v111)%32) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v125 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	if v121 == v122 {
		v125 = int32(1)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v125 = int32(0)
	goto L27
L31:
	;
	goto L30
L32:
	;
	goto L33
L33:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v137 = v135 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v137)
	v139 = int32(_a_F_scanGetCandidate_0)
	v140 = v137 & v139
	if base.Ui32(v49&v139) < base.Ui32(v140) {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L1
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v77+v140<<(uint(int32(2))%32))))
	v150 = v40 + v147&int32(_a_F_scanGetCandidate_1)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v153 = int32(16)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+2)))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150))))
	if v151|v152<<(uint(v153)%32) == v156|v157<<(uint(v153)%32) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v167 != 0 {
		goto L33
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	if v163 == v164 {
		v167 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v167 = int32(0)
	goto L37
L41:
	;
	goto L40
L42:
	;
	goto L34
}
func F_scanexp(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v115 int64
	_ = v115
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v156 int64
	_ = v156
	var v161 int64
	_ = v161
	var v164 int32
	_ = v164
	var v175 int64
	_ = v175
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v7 != v8 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	return v175
L2:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v161 < int64(0) {
		v175 = int64(-9223372036854775807 - 1)
		goto L1
	} else {
		goto L54
	}
L3:
	;
	if base.Ui32(v49) < base.Ui32(int32(-10)) {
		goto L2
	} else {
		goto L19
	}
L4:
	;
	v49 = v18 - int32(58)
	v50 = v18
	v51 = int32(0)
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v21 != v22 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	switch v18 - int32(43) {
	case 0, 2:
		goto L5
	default:
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + int32(1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v18 = v13
	goto L6
L8:
	;
	goto L9
L9:
	;
	v14 = F___shgetc(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int64(0)
L11:
	;
	v18 = v14
	goto L6
L12:
	;
	v36 = v30 - int32(58)
	if base.B2i32(l1 == int32(0))|base.B2i32(base.Ui32(int32(-11)) < base.Ui32(v36)) != 0 {
		v49 = v36
		v50 = v30
		v51 = base.B2i32(v18 == int32(45))
		goto L3
	} else {
		goto L17
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 + int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v30 = v27
	goto L12
L14:
	;
	goto L15
L15:
	;
	v28 = F___shgetc(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v30 = v28
	goto L12
L17:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v40 < int64(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - int32(1)
	goto L2
L19:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v50-int32(48)) {
		v146 = int64(0)
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if int64(0) <= v147 {
		goto L48
	} else {
		goto L49
	}
L21:
	;
	v61 = int32(0)
	v62 = v50
	goto L22
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v68 != v69 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v87 = base.I64_extend_i32_s(v83)
	if base.Ui32(int32(10)) <= base.Ui32(v79) {
		v146 = v87
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v78 = int32(48)
	v79 = v77 - v78
	v83 = v62 + v61*int32(10) - v78
	if base.B2i32(base.Ui32(v79) <= base.Ui32(int32(9)))&base.B2i32(v83 < int32(214748364)) != 0 {
		v61 = v83
		v62 = v77
		goto L22
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v77 = v74
	goto L24
L26:
	;
	goto L27
L27:
	;
	v75 = F___shgetc(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v77 = v75
	goto L24
L29:
	;
	goto L23
L30:
	;
	v93 = v77
	v95 = v87
	goto L31
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v100 != v101 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v111) {
		v146 = v115
		goto L20
	} else {
		goto L39
	}
L33:
	;
	v111 = v109 - int32(48)
	v115 = base.I64_extend_i32_u(v93) + v95*int64(10) - int64(48)
	if base.B2i32(base.Ui32(v111) <= base.Ui32(int32(9)))&base.B2i32(v115 < int64(92233720368547758)) != 0 {
		v93 = v109
		v95 = v115
		goto L31
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v109 = v106
	goto L33
L35:
	;
	goto L36
L36:
	;
	v107 = F___shgetc(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v109 = v107
	goto L33
L38:
	;
	goto L32
L39:
	;
	goto L40
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v127 != v128 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v146 = v115
	goto L20
L42:
	;
	if base.Ui32(v136-int32(48)) < base.Ui32(int32(10)) {
		goto L40
	} else {
		goto L47
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 + int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v136 = v133
	goto L42
L44:
	;
	goto L45
L45:
	;
	v134 = F___shgetc(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v136 = v134
	goto L42
L47:
	;
	goto L41
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150 - int32(1)
	goto L50
L49:
	;
	goto L50
L50:
	;
	if v51 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v156 = int64(0) - v146
	goto L53
L52:
	;
	v156 = v146
	goto L53
L53:
	;
	v175 = v156
	goto L1
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164 - int32(1)
	return int64(-9223372036854775807 - 1)
}
func F_search_indexed_tlist_for_phv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L63
	} else {
		goto L65
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v233
L3:
	;
	v233 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 <= v18 {
		v233 = v18
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = v19
	goto L9
L8:
	;
	v25 = v22
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v28 = v18
	goto L12
L10:
	;
	v224 = F_makeVarFromTargetEntry(m, l2, v39)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L63
	} else {
		goto L64
	}
L11:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v168 = int32(0)
	if v166 == v168 {
		goto L49
	} else {
		goto L50
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v233 = int32(0)
	goto L2
L14:
	;
	v163 = v28 + int32(1)
	if v163 != v25 {
		v28 = v163
		goto L12
	} else {
		goto L47
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v43 != int32(319) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v46 != v47 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	switch l3 - int32(1) {
	case 0:
		goto L11
	case 1:
		goto L19
	default:
		goto L18
	}
L18:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = int32(0)
	if base.B2i32(v109 == v111)|base.B2i32(v110 == v111) != 0 {
		v157 = base.B2i32(v109|v110 == v111)
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = int32(0)
	if v51 == v53 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v106 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	v106 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v52 == int32(0) {
		v99 = v53
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = v99
	goto L20
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v63 < v62 {
		v99 = v53
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(1)
	if v62 <= v65 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = v65
	goto L29
L28:
	;
	v68 = v62
	goto L29
L29:
	;
	v69 = int32(8)
	v74 = int32(0)
	goto L30
L30:
	;
	v81 = v74 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v51+v69+v81)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v52+v69+v81)))
	v88 = v83 & (v85 ^ int32(-1))
	v90 = base.B2i32(v88 == int32(0))
	if v88 != 0 {
		v99 = v90
		goto L24
	} else {
		goto L32
	}
L31:
	;
	v99 = v90
	goto L24
L32:
	;
	v92 = v74 + int32(1)
	if v92 != v68 {
		v74 = v92
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L10
L35:
	;
	if v157 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v125 != v126 {
		v157 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v128 = int32(1)
	if v125 <= v128 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v131 = v128
	goto L41
L40:
	;
	v131 = v125
	goto L41
L41:
	;
	v132 = int32(8)
	v137 = int32(0)
	goto L42
L42:
	;
	v145 = v137 << (uint(int32(2)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v109+v132+v145)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v110+v132+v145)))
	v150 = base.B2i32(v147 == v149)
	if v147 != v149 {
		v157 = v150
		goto L36
	} else {
		goto L44
	}
L43:
	;
	v157 = v150
	goto L36
L44:
	;
	v153 = v137 + int32(1)
	if v153 != v131 {
		v137 = v153
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L1
L47:
	;
	goto L13
L48:
	;
	if v221 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L49:
	;
	v221 = int32(1)
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v167 == int32(0) {
		v214 = v168
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v221 = v214
	goto L48
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v178 < v177 {
		v214 = v168
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v180 = int32(1)
	if v177 <= v180 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v180
	goto L57
L56:
	;
	v183 = v177
	goto L57
L57:
	;
	v184 = int32(8)
	v189 = int32(0)
	goto L58
L58:
	;
	v196 = v189 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v166+v184+v196)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v167+v184+v196)))
	v203 = v198 & (v200 ^ int32(-1))
	v205 = base.B2i32(v203 == int32(0))
	if v203 != 0 {
		v214 = v205
		goto L52
	} else {
		goto L60
	}
L59:
	;
	v214 = v205
	goto L52
L60:
	;
	v207 = v189 + int32(1)
	if v207 != v183 {
		v189 = v207
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L10
L63:
	;
	return int32(0)
L64:
	;
	v228 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+40)) = uint16(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+36)) = v228
	v233 = v224
	goto L2
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v250 = F_bmsToString(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v253 = F_bmsToString(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v250
	F_errmsg_internal(m, int32(_a_F_search_indexed_tlist_for_phv_0), v12)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_search_indexed_tlist_for_phv_1), int32(2965), int32(_a_F_search_indexed_tlist_for_phv_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_session_user(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_session_user[0]))
	v7 = F_GetUserNameFromId(m, v5, v3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_DirectFunctionCall1Coll(m, int32(500), v3, v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_set_upper_references(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v229 int32
	_ = v229
	var v237 float64
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = int32(12)
	v29 = v23*v24 + v24
	goto L3
L2:
	;
	v29 = int32(12)
	goto L3
L3:
	;
	v30 = F_palloc(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v32 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+8)) = uint16(v32)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v22
	v36 = v30 + int32(12)
	if v22 == v32 {
		v98 = v36
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v110 = base.I32_div_s(v98-v36, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v112 != int32(365) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v39 <= int32(0) {
		v98 = v36
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v46 = v4
	v48 = v36
	goto L9
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v46<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v63 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v98 = v86
	goto L6
L11:
	;
	v89 = v46 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v89 < v90 {
		v46 = v89
		v48 = v86
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+9)) = uint8(v84)
	v86 = v48
	goto L11
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v66 != int32(319) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v66 != int32(6) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+8)) = uint8(v81)
	v86 = v48
	goto L11
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v71
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)) = uint16(v73)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v77
	v86 = v48 + int32(12)
	goto L11
L18:
	;
	goto L10
L19:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v137 == int32(0) {
		v293 = v4
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v115 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v118 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v122 = F_bms_make_singleton(m, v115)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v125 = F_remove_nulling_relids(m, v121, v122, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v130 = F_bms_make_singleton(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v133 = F_remove_nulling_relids(m, v128, v130, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v133
	goto L19
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v293
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v298 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = base.F64_add(v298, v298)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v308 = F_fix_upper_expr_mutator(m, v297, v19)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L53
	}
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		v293 = v4
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v155 = v4
	v156 = v4
	goto L30
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v155<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v165 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v293 = v274
	goto L27
L32:
	;
	v271 = F_flatCopyTargetEntry(m, v163)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L50
	}
L33:
	;
	v249 = F_makeVarFromTargetEntry(m, int32(-2), v193)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L49
	}
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v166 == int32(0) {
		v213 = v164
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v229 = v164
	goto L36
L36:
	;
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v246 = F_fix_upper_expr_mutator(m, v229, v19)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L48
	}
L37:
	;
	v229 = v213
	goto L36
L38:
	;
	v169 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v170 <= v169 {
		v213 = v164
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v177 = v169
	v183 = v170
	goto L40
L40:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v177<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v165 == v194 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v213 = v204
	goto L37
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v197 = F_equal(m, v164, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	v200 = v183
	goto L44
L44:
	;
	v202 = v177 + int32(1)
	if v202 < v200 {
		v177 = v202
		v183 = v200
		goto L40
	} else {
		goto L47
	}
L45:
	;
	if v197 != 0 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v200 = v199
	goto L44
L47:
	;
	goto L41
L48:
	;
	v259 = v246
	goto L32
L49:
	;
	v251 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+40)) = uint16(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+36)) = v251
	v259 = v249
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = v259
	v274 = F_lappend(m, v156, v271)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v277 = v155 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v277 < v278 {
		v155 = v277
		v156 = v274
		goto L30
	} else {
		goto L52
	}
L52:
	;
	goto L31
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v308
	F_pfree(m, v30)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v19 + int32(32)
	return
}
func F_setitimer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v16 float64
	_ = v16
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v39 int32
	_ = v39
	v7 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v8), float64(1000)), base.F64_convert_i64_s(v12*int64(1000)))
	if v12 == int64(0) {
		if v8 == int32(0) {
			v33 = float64(0)
			v34 = float64(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v33 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v23), float64(1000)), base.F64_convert_i64_s(v27*int64(1000)))
			v34 = base.F64_add(v7, v16)
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v33 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v23), float64(1000)), base.F64_convert_i64_s(v27*int64(1000)))
		v34 = base.F64_add(v7, v16)
	}
	*(*float64)(unsafe.Add(mBase, _c_F_setitimer[0])) = v34
	*(*float64)(unsafe.Add(mBase, _c_F_setitimer[1])) = v33
	v39 = m.Env.X_setitimer_js(m, int32(0), v16)
	mBase = m.M
	return v39
}
func F_setup_parser_errposition_callback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l0
	v9 = int32(_a_F_setup_parser_errposition_callback_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_setup_parser_errposition_callback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	*(*int32)(unsafe.Add(mBase, _c_F_setup_parser_errposition_callback[0])) = l0 + int32(8)
	return
}
func F_setup_pct_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v78 float64
	_ = v78
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v101 int32
	_ = v101
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v20 = F_palloc(m, l0<<(uint(int32(5))%32))
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
	if int32(0) < l0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v36 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pg_qsort(m, v20, l0, int32(32), int32(1461))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L21
	}
L7:
	;
	v45 = v20 + v36<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = v36
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v36))))
	if v48 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v101 = v36 + int32(1)
	if v101 != l0 {
		v36 = v101
		goto L7
	} else {
		goto L20
	}
L10:
	;
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v51
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1+v36<<(uint(int32(2))%32))))
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	if base.F64_lt(v61, float64(0))|base.F64_gt(v61, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807))) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if l4 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = base.F64_mul(v61, base.F64_convert_i64_s(l3-int64(1)))
	v74 = base.F64_floor(v73)
	*(*float64)(unsafe.Add(mBase, uint32(v45)+16)) = base.F64_sub(v73, v74)
	v78 = float64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = base.I64_trunc_sat_f64_s(base.F64_add(base.F64_ceil(v73), v78))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = base.I64_trunc_sat_f64_s(base.F64_add(v74, v78))
	goto L9
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = int64(0)
	v88 = int64(1)
	v91 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v61, base.F64_convert_i64_s(l3))))
	if v91 <= v88 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v94 = v88
	goto L19
L18:
	;
	v94 = v91
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v94
	goto L9
L20:
	;
	goto L8
L21:
	;
	m.G0 = v16 + int32(16)
	return v20
L22:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16))) = v61
	F_errmsg(m, int32(_a_F_setup_pct_info_0), v16)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_setup_pct_info_1), int32(692), int32(_a_F_setup_pct_info_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shell_archive_file(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
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
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	if l2 != 0 {
		v10 = F_pstrdup(m, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = v10
			*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[0]))
			v23 = F_replace_percent_placeholders(m, v18, int32(_a_F_shell_archive_file_0), int32(_a_F_shell_archive_file_1), v8+int32(96))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v23
						F_errmsg_internal(m, int32(_a_F_shell_archive_file_2), v8+int32(80))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(76), int32(_a_F_shell_archive_file_4))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v41 = F_fflush(m, int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
									v47 = F_pgl_system(m, v23)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
										if v47 != 0 {
											v63 = int32(255)
											if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
												v75 = int32(22)
											} else {
												v75 = int32(15)
											}
											v77 = v47 & int32(127)
											if v77 == int32(0) {
												v81 = F_errstart(m, v75, int32(0))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													if v81 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
														F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															v145 = int32(101)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												}
											} else {
												v97 = F_errstart(m, v75, int32(0))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
														if v97 == int32(0) {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															v109 = int32(_a_F_shell_archive_file_8)
															if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
																v114 = v77
																v115 = v109
																for {
																	v118 = v115 + int32(1)
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
																	if v119 != 0 {
																		v115 = v118
																		continue
																	} else {
																	}
																	v121 = v114 - int32(1)
																	if v121 != 0 {
																		v114 = v121
																		v115 = v118
																		continue
																	} else {
																		break
																	}
																	break
																}
																v123 = v118
															} else {
																v123 = v109
															}
															if v123 != 0 {
																v126 = v123
															} else {
																v126 = int32(_a_F_shell_archive_file_9)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
															*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
															F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																v145 = int32(117)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v8 + int32(112)
																			return base.B2i32(v47 == int32(0))
																		}
																	}
																}
															}
														}
													} else {
														if v97 == int32(0) {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
															F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																v145 = int32(126)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v8 + int32(112)
																			return base.B2i32(v47 == int32(0))
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												v163 = F_errstart(m, int32(14), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													if v163 == int32(0) {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
														F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
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
						v41 = F_fflush(m, int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
							v47 = F_pgl_system(m, v23)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
								if v47 != 0 {
									v63 = int32(255)
									if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
										v75 = int32(22)
									} else {
										v75 = int32(15)
									}
									v77 = v47 & int32(127)
									if v77 == int32(0) {
										v81 = F_errstart(m, v75, int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											if v81 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
												F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													v145 = int32(101)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										}
									} else {
										v97 = F_errstart(m, v75, int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
												if v97 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													v109 = int32(_a_F_shell_archive_file_8)
													if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
														v114 = v77
														v115 = v109
														for {
															v118 = v115 + int32(1)
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
															if v119 != 0 {
																v115 = v118
																continue
															} else {
															}
															v121 = v114 - int32(1)
															if v121 != 0 {
																v114 = v121
																v115 = v118
																continue
															} else {
																break
															}
															break
														}
														v123 = v118
													} else {
														v123 = v109
													}
													if v123 != 0 {
														v126 = v123
													} else {
														v126 = int32(_a_F_shell_archive_file_9)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
													*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
													F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														v145 = int32(117)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											} else {
												if v97 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
													F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														v145 = int32(126)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									F_pfree(m, v23)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										v163 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												m.G0 = v8 + int32(112)
												return base.B2i32(v47 == int32(0))
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
												F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
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
		}
	} else {
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v14
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[0]))
		v23 = F_replace_percent_placeholders(m, v18, int32(_a_F_shell_archive_file_0), int32(_a_F_shell_archive_file_1), v8+int32(96))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v27 = F_errstart(m, int32(12), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v23
					F_errmsg_internal(m, int32(_a_F_shell_archive_file_2), v8+int32(80))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(76), int32(_a_F_shell_archive_file_4))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = F_fflush(m, int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
								v47 = F_pgl_system(m, v23)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
									if v47 != 0 {
										v63 = int32(255)
										if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
											v75 = int32(22)
										} else {
											v75 = int32(15)
										}
										v77 = v47 & int32(127)
										if v77 == int32(0) {
											v81 = F_errstart(m, v75, int32(0))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												if v81 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
													F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														v145 = int32(101)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											}
										} else {
											v97 = F_errstart(m, v75, int32(0))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
													if v97 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														v109 = int32(_a_F_shell_archive_file_8)
														if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
															v114 = v77
															v115 = v109
															for {
																v118 = v115 + int32(1)
																v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
																if v119 != 0 {
																	v115 = v118
																	continue
																} else {
																}
																v121 = v114 - int32(1)
																if v121 != 0 {
																	v114 = v121
																	v115 = v118
																	continue
																} else {
																	break
																}
																break
															}
															v123 = v118
														} else {
															v123 = v109
														}
														if v123 != 0 {
															v126 = v123
														} else {
															v126 = int32(_a_F_shell_archive_file_9)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
														*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
														F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															v145 = int32(117)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												} else {
													if v97 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
														F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															v145 = int32(126)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											v163 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												if v163 == int32(0) {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
													F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
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
					v41 = F_fflush(m, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
						v47 = F_pgl_system(m, v23)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
							if v47 != 0 {
								v63 = int32(255)
								if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
									v75 = int32(22)
								} else {
									v75 = int32(15)
								}
								v77 = v47 & int32(127)
								if v77 == int32(0) {
									v81 = F_errstart(m, v75, int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										if v81 == int32(0) {
											F_pfree(m, v23)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(112)
												return base.B2i32(v47 == int32(0))
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
											F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												v145 = int32(101)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
												F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													}
												}
											}
										}
									}
								} else {
									v97 = F_errstart(m, v75, int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
											if v97 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												v109 = int32(_a_F_shell_archive_file_8)
												if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
													v114 = v77
													v115 = v109
													for {
														v118 = v115 + int32(1)
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
														if v119 != 0 {
															v115 = v118
															continue
														} else {
														}
														v121 = v114 - int32(1)
														if v121 != 0 {
															v114 = v121
															v115 = v118
															continue
														} else {
															break
														}
														break
													}
													v123 = v118
												} else {
													v123 = v109
												}
												if v123 != 0 {
													v126 = v123
												} else {
													v126 = int32(_a_F_shell_archive_file_9)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
												*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
												F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													v145 = int32(117)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										} else {
											if v97 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
												F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													v145 = int32(126)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									v163 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										if v163 == int32(0) {
											m.G0 = v8 + int32(112)
											return base.B2i32(v47 == int32(0))
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
											F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
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
	}
}
func F_shell_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_shell_in_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_shell_in_1), int32(307), int32(_a_F_shell_in_2))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_shift_jis_2004_to_utf8(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13870(m, l0, int32(41), int32(0), int32(25), int32(_a_F_shift_jis_2004_to_utf8_0), int32(_a_F_shift_jis_2004_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_shimTriConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v368
L2:
	;
	v346 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v346)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v359 = F_FunctionCall8Coll(m, v348, v349, v350, v351, v352, v353, v354, l0+int32(87), v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L26
	} else {
		goto L60
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = v2
	v21 = v2
	goto L4
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v18))))
	if v30 == int32(2) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L6:
	;
	if int32(3) < v21 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v42 = v21
	goto L8
L8:
	;
	v44 = v20 + int32(1)
	if v44 != v15 {
		v20 = v44
		v21 = v42
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v368 = int32(2)
	goto L1
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+v21<<(uint(int32(2))%32)))) = v20
	v42 = v21 + int32(1)
	goto L8
L12:
	;
	goto L5
L13:
	;
	v49 = base.B2i32(v42 <= int32(0))
	if v42 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v142)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v152 = l0 + int32(87)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v155 = F_FunctionCall8Coll(m, v144, v145, v146, v147, v148, v149, v150, v152, v153, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v51 = v42 & int32(3)
	v52 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v59 = v52
	v62 = int32(0)
	goto L19
L17:
	;
	v99 = v52
	goto L18
L18:
	;
	v110 = v99
	v112 = int32(0)
	goto L23
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v71 = v13 + v59<<(uint(int32(2))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+v72))) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v76+v77))) = uint8(v74)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v81+v82))) = uint8(v74)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v86+v87))) = uint8(v74)
	v91 = int32(4)
	v92 = v59 + v91
	v94 = v62 + v91
	if v94 != v42&int32(2147483644) {
		v59 = v92
		v62 = v94
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v51 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v99 = v92
	goto L18
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v13+v110<<(uint(int32(2))%32))))
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v123))) = uint8(v125)
	v127 = int32(1)
	v130 = v112 + v127
	if v130 != v51 {
		v110 = v110 + v127
		v112 = v130
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	goto L24
L26:
	;
	return int32(0)
L27:
	;
	v160 = base.B2i32(v155 != int32(0))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)))
	v166 = v161
	goto L29
L28:
	;
	if v239&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L29:
	;
	v172 = int32(0)
	if v49 == v172 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v238 = int32(2)
	v239 = v228
	goto L28
L31:
	;
	v176 = v172
	goto L34
L32:
	;
	goto L33
L33:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v210)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v221 = F_FunctionCall8Coll(m, v212, v213, v214, v215, v216, v217, v218, v152, v219, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L41
	}
L34:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v13+v176<<(uint(int32(2))%32))))
	v190 = v185 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v191 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v197)
	if v176 == v42 {
		v238 = v160
		v239 = v166
		goto L28
	} else {
		goto L40
	}
L36:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v192)
	v195 = v176 + int32(1)
	if v195 != v42 {
		v176 = v195
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
	v238 = v160
	v239 = v166
	goto L28
L40:
	;
	goto L33
L41:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)))
	v227 = int32(0)
	v228 = base.B2i32(v223|v166&int32(1) != v227)
	if base.B2i32(v155 != v227) == base.B2i32(v221 != v227) {
		v166 = v228
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	v248 = int32(2)
	goto L45
L44:
	;
	v248 = v238
	goto L45
L45:
	;
	if v238 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v251 = v248
	goto L48
L47:
	;
	v251 = v238
	goto L48
L48:
	;
	if v42 <= int32(0) {
		v368 = v251
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v255 = v42 & int32(3)
	v256 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v264 = v256
	v267 = int32(0)
	goto L53
L51:
	;
	v304 = v256
	goto L52
L52:
	;
	v314 = v304
	v316 = v256
	goto L57
L53:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v274 = int32(2)
	v276 = v13 + v264<<(uint(v274)%32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	*(*uint8)(unsafe.Add(mBase, uint32(v273+v277))) = uint8(v274)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v281+v282))) = uint8(v274)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v286+v287))) = uint8(v274)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v291+v292))) = uint8(v274)
	v296 = int32(4)
	v297 = v264 + v296
	v299 = v267 + v296
	if v299 != v42&int32(2147483644) {
		v264 = v297
		v267 = v299
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v255 == int32(0) {
		v368 = v251
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v304 = v297
	goto L52
L57:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v324 = int32(2)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v13+v314<<(uint(v324)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v323+v327))) = uint8(v324)
	v331 = int32(1)
	v334 = v316 + v331
	if v334 != v255 {
		v314 = v314 + v331
		v316 = v334
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v368 = v251
	goto L1
L59:
	;
	goto L58
L60:
	;
	v368 = base.B2i32(v359 != int32(0))
	goto L1
}
func F_shim_pclose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = int32(-1)
	if l0 == int32(0) {
		v32 = v3
		return v32
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[0]))
		if l0 != v7 {
			v32 = v3
			return v32
		} else {
			v9 = F_fclose(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[0])) = int32(0)
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_shim_pclose[1])))
				if v17 != 0 {
					v19 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_shim_pclose[1])) = uint8(v19)
					v24 = m.Env.Pgmem_run(m, int32(_a_F_shim_pclose_0), int32(_a_F_shim_pclose_1), int32(_a_F_shim_pclose_2))
					mBase = m.M
					return v24 << (uint(int32(8)) % 32) & int32(_a_F_shim_pclose_3)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[2]))
					v32 = v31
					return v32
				}
			}
		}
	}
}
func F_shim_popen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	v3 = int32(_a_F_shim_popen_0)
	v4 = int32(_a_F_shim_popen_1)
	if (l0^v3)&int32(3) != 0 {
		v75 = l0
		v76 = v4
		v77 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[0])) = uint8(v115)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v118 == int32(119) {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	F___memset(m, v110, int32(0), v109)
	mBase = m.M
	goto L1
L3:
	;
	v109 = int32(0)
	v110 = v104
	goto L2
L4:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L22
L5:
	;
	if v76 == int32(0) {
		v104 = v77
		goto L3
	} else {
		goto L21
	}
L6:
	;
	if base.B2i32(l0&int32(3) == int32(0))|int32(0) != 0 {
		v41 = l0
		v42 = v4
		v43 = v3
		v44 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v44 == int32(0) {
		v104 = v43
		goto L3
	} else {
		goto L14
	}
L8:
	;
	v20 = l0
	v21 = v4
	v22 = v3
	goto L9
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
	if v24 == int32(0) {
		v109 = v21
		v110 = v22
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v41 = v35
	v42 = v31
	v43 = v29
	v44 = v33
	goto L7
L11:
	;
	v28 = int32(1)
	v29 = v22 + v28
	v31 = v21 - v28
	v32 = int32(0)
	v33 = base.B2i32(v31 != v32)
	v35 = v20 + v28
	if v35&int32(3) == v32 {
		v41 = v35
		v42 = v31
		v43 = v29
		v44 = v33
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v31 != 0 {
		v20 = v35
		v21 = v31
		v22 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v47 == int32(0) {
		v109 = v42
		v110 = v43
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v42) < base.Ui32(int32(4)) {
		v75 = v41
		v76 = v42
		v77 = v43
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v53 = v41
	v54 = v42
	v55 = v43
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v82 = v53
		v83 = v54
		v84 = v55
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v75 = v69
	v76 = v71
	v77 = v67
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v58
	v66 = int32(4)
	v67 = v55 + v66
	v69 = v53 + v66
	v71 = v54 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v53 = v69
		v54 = v71
		v55 = v67
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L4
L22:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v91)
	if v91 == int32(0) {
		v109 = v88
		v110 = v89
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v104 = v96
	goto L3
L24:
	;
	v95 = int32(1)
	v96 = v89 + v95
	v100 = v88 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v100
		v89 = v96
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[1])) = v142
	return v142
L27:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v122)
	v126 = F_fopen(m, int32(_a_F_shim_popen_2), int32(_a_F_shim_popen_3))
	mBase = m.M
	v142 = v126
	goto L26
L28:
	;
	goto L29
L29:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v128)
	v132 = int32(_a_F_shim_popen_4)
	v133 = m.Env.Pgmem_run(m, l0, int32(_a_F_shim_popen_5), v132)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[3])) = v133 << (uint(int32(8)) % 32) & int32(_a_F_shim_popen_6)
	v141 = F_fopen(m, v132, int32(_a_F_shim_popen_7))
	mBase = m.M
	v142 = v141
	goto L26
}
func F_should_apply_changes_for_rel(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v11 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_should_apply_changes_for_rel_0), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(487), int32(_a_F_should_apply_changes_for_rel_2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v47 = base.B2i32(v43 == v44)
		m.G0 = v7 + int32(16)
		return v47
	case 2:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		switch v20 - int32(114) {
		case 0:
			v47 = int32(1)
		case 1:
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
			v25 = *(*int64)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[1]))
			v47 = base.B2i32(base.Ui64(v23) <= base.Ui64(v25))
		default:
			v47 = int32(0)
		}
		m.G0 = v7 + int32(16)
		return v47
	case 3:
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v12 != int32(114) {
			v16 = v12
		} else {
			v16 = int32(0)
		}
		if v16 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[2]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v61
					F_errmsg(m, int32(_a_F_should_apply_changes_for_rel_3), v7)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(_a_F_should_apply_changes_for_rel_4), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(476), int32(_a_F_should_apply_changes_for_rel_2))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
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
		} else {
			v47 = base.B2i32(v12 == int32(114))
			m.G0 = v7 + int32(16)
			return v47
		}
	default:
		v47 = int32(0)
		m.G0 = v7 + int32(16)
		return v47
	}
}
func F_should_refetch_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l0 {
	case 0:
		v85 = int32(0)
		m.G0 = v7 + int32(16)
		return v85
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_should_refetch_tuple_0), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_refetch_tuple_1), int32(162), int32(_a_F_should_refetch_tuple_2))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_should_refetch_tuple_3), v7)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_refetch_tuple_1), int32(165), int32(_a_F_should_refetch_tuple_2))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9 != int32(_a_F_should_refetch_tuple_4) {
			v29 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 == int32(0) {
					v85 = int32(1)
					m.G0 = v7 + int32(16)
					return v85
				} else {
					v69 = int32(_a_F_should_refetch_tuple_5)
					v70 = int32(151)
					F_errcode(m, int32(16777220))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, v69, int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v85 = int32(1)
								m.G0 = v7 + int32(16)
								return v85
							}
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			if v12&v13 != int32(_a_F_should_refetch_tuple_6) {
				v29 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v29 == int32(0) {
						v85 = int32(1)
						m.G0 = v7 + int32(16)
						return v85
					} else {
						v69 = int32(_a_F_should_refetch_tuple_5)
						v70 = int32(151)
						F_errcode(m, int32(16777220))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, v69, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v85 = int32(1)
									m.G0 = v7 + int32(16)
									return v85
								}
							}
						}
					}
				}
			} else {
				v19 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v85 = int32(1)
						m.G0 = v7 + int32(16)
						return v85
					} else {
						v69 = int32(_a_F_should_refetch_tuple_7)
						v70 = int32(147)
						F_errcode(m, int32(16777220))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, v69, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v85 = int32(1)
									m.G0 = v7 + int32(16)
									return v85
								}
							}
						}
					}
				}
			}
		}
	case 4:
		v37 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			if v37 == int32(0) {
				v85 = int32(1)
				m.G0 = v7 + int32(16)
				return v85
			} else {
				v69 = int32(_a_F_should_refetch_tuple_8)
				v70 = int32(158)
				F_errcode(m, int32(16777220))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, v69, int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v85 = int32(1)
							m.G0 = v7 + int32(16)
							return v85
						}
					}
				}
			}
		}
	}
}
func F_show_archive_command(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_show_archive_command[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_show_archive_command[1]))
	if v5 <= int32(0) {
		v8 = int32(_a_F_show_archive_command_0)
	} else {
		v8 = v3
	}
	return v8
}
func F_sigaddset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = l1 - int32(1)
	if base.B2i32(base.Ui32(v5) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1-int32(32))) == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_sigaddset[0])) = int32(28)
		return
	} else {
		v22 = l0 + int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 | int32(1)<<(uint(v5)%32)
		return
	}
}
func F_sigdelset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = l1 - int32(1)
	if base.B2i32(base.Ui32(v5) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1-int32(32))) == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_sigdelset[0])) = int32(28)
		return
	} else {
		v22 = l0 + int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 & base.I32_rotl(int32(-2), v5)
		return
	}
}
func F_signValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v8 = int32(_a_F_signValue_0)
	v10 = base.I32_rem_s(l3, int32(2147483646))
	*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v10 + int32(1)
	v20 = l0 + l3<<(uint(int32(2))%32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+896))
	v22 = F_FunctionCall1Coll(m, l0+l3*int32(28), v21, l2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_signValue[0]))
		v28 = int32(_a_F_signValue_1)
		v29 = base.I32_div_s(v27, v28)
		v37 = v29*int32(-2836) + (v27-v29*v28)*int32(_a_F_signValue_2)
		if int32(0) <= v37 {
			v40 = int32(-1)
		} else {
			v40 = int32(2147483646)
		}
		v44 = base.I32_rem_s(v22^(v40+v37), int32(2147483646))
		v46 = v44 + int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v46
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(1032))))
		if int32(0) < v50 {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1028))
			v57 = int32(0)
			v60 = v46
			for {
				v64 = int32(_a_F_signValue_1)
				v65 = base.I32_div_s(v60, v64)
				v73 = v65*int32(-2836) + (v60-v65*v64)*int32(_a_F_signValue_2)
				if v73 < int32(0) {
					v78 = v73 + int32(2147483647)
				} else {
					v78 = v73
				}
				v79 = int32(1)
				v81 = base.I32_rem_s(v78-v79, v53<<(uint(int32(4))%32))
				v83 = base.I32_div_s(v81, int32(16))
				v86 = l1 + v83<<(uint(v79)%32)
				v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
				v92 = v87 | v79<<(uint(v81&int32(15))%32)
				*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v92)
				v95 = v57 + v79
				if v95 != v50 {
					v57 = v95
					v60 = v78
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v78
		} else {
		}
		return
	}
}
func F_similar_to_escape_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_similar_escape_internal(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_similarity_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_DirectFunctionCall2Coll(m, int32(_a_F_similarity_dist_0), int32(0), v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return base.I32_reinterpret_f32(base.F32_sub(float32(1), base.F32_reinterpret_i32(v7)))
	}
}
func F_sin(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v73 float64
	_ = v73
	var v89 float64
	_ = v89
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v127 float64
	_ = v127
	var v147 float64
	_ = v147
	var v163 float64
	_ = v163
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v202 float64
	_ = v202
	var v219 float64
	_ = v219
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v14) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v14) < base.Ui32(int32(1045430272)) {
			v219 = l0
		} else {
			v24 = base.F64_mul(l0, l0)
			v219 = base.F64_add(base.F64_mul(base.F64_mul(l0, v24), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(base.F64_mul(v24, base.F64_mul(v24, v24)), base.F64_add(base.F64_mul(v24, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(v24, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), l0)
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v14) {
			v219 = base.F64_sub(l0, l0)
		} else {
			v62 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v63 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v64 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v62&int32(3) - int32(1) {
			case 0:
				v111 = float64(1)
				v112 = base.F64_mul(v64, v64)
				v114 = base.F64_mul(v112, float64(0.5))
				v115 = base.F64_sub(v111, v114)
				v127 = base.F64_mul(v112, v112)
				v219 = base.F64_add(v115, base.F64_add(base.F64_sub(base.F64_sub(v111, v115), v114), base.F64_sub(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v127, v127), base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v64, v63))))
			case 1:
				v147 = base.F64_mul(v64, v64)
				v163 = base.F64_mul(v64, v147)
				v219 = base.F64_neg(base.F64_sub(v64, base.F64_add(base.F64_sub(base.F64_mul(v147, base.F64_sub(base.F64_mul(v63, float64(0.5)), base.F64_mul(v163, base.F64_add(base.F64_mul(base.F64_mul(v147, base.F64_mul(v147, v147)), base.F64_add(base.F64_mul(v147, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v147, base.F64_add(base.F64_mul(v147, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v63), base.F64_mul(v163, float64(0.16666666666666632)))))
			case 2:
				v186 = float64(1)
				v187 = base.F64_mul(v64, v64)
				v189 = base.F64_mul(v187, float64(0.5))
				v190 = base.F64_sub(v186, v189)
				v202 = base.F64_mul(v187, v187)
				v219 = base.F64_neg(base.F64_add(v190, base.F64_add(base.F64_sub(base.F64_sub(v186, v190), v189), base.F64_sub(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v202, v202), base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v64, v63)))))
			default:
				v73 = base.F64_mul(v64, v64)
				v89 = base.F64_mul(v64, v73)
				v219 = base.F64_sub(v64, base.F64_add(base.F64_sub(base.F64_mul(v73, base.F64_sub(base.F64_mul(v63, float64(0.5)), base.F64_mul(v89, base.F64_add(base.F64_mul(base.F64_mul(v73, base.F64_mul(v73, v73)), base.F64_add(base.F64_mul(v73, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v63), base.F64_mul(v89, float64(0.16666666666666632))))
			}
		}
	}
	m.G0 = v7 + int32(16)
	return v219
}
func F_single_bound_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(4)
	v8 = F_range_cmp_bounds(m, l2, l0+v4, l1+v4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_skip_b_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	if l3 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = l1
	v13 = l3
	goto L7
L5:
	;
	v48 = l1
	goto L6
L6:
	;
	return v48
L7:
	;
	if v11 <= l2 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v48 = v39
	goto L6
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v19 = v11 - int32(1)
	v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v19))))
	if base.B2i32(int32(0) <= v21)|base.B2i32(v19 <= l2) != 0 {
		v39 = v19
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(1)
	if v43 < v13 {
		v11 = v39
		v13 = v13 - v43
		goto L7
	} else {
		goto L18
	}
L13:
	;
	v27 = v19
	goto L14
L14:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v27))))
	if base.Ui32(int32(191)) < base.Ui32(v32) {
		v39 = v27
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v39 = l2
	goto L12
L16:
	;
	v36 = v27 - int32(1)
	if l2 < v36 {
		v27 = v36
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L8
}
func F_slice_from_v(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v8 = int32(-1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v9 < int32(0) {
		v31 = v8
		return v31
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 < v9 {
			v31 = v8
			return v31
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v14 < v12 {
				v31 = v8
				return v31
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v16 == int32(0) {
					v31 = v8
					return v31
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(4))))
					if v21 < v14 {
						v31 = v8
						return v31
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
						v27 = F_replace_s(m, l0, v9, v12, v25, l1, int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = v27
							return v31
						}
					}
				}
			}
		}
	}
}
func F_slotsync_failure_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[0]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(1))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
				if v11 != 0 {
					v13 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					v16 = base.AtomicRmwXchg32(m, v13, int32(16), int32(1))
					if v16 != 0 {
						v18 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						F_s_lock(m, v18+int32(16), int32(_a_F_slotsync_failure_callback_0), int32(1339), int32(_a_F_slotsync_failure_callback_1))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
							v28 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v28)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v27)+16)), uint32(v28))
							*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v28)
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
							m.T0[v39].(func(*base.Module, int32))(m, l1)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						v28 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v28)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v27)+16)), uint32(v28))
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v28)
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
						m.T0[v39].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
					m.T0[v39].(func(*base.Module, int32))(m, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
			if v11 != 0 {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
				v16 = base.AtomicRmwXchg32(m, v13, int32(16), int32(1))
				if v16 != 0 {
					v18 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					F_s_lock(m, v18+int32(16), int32(_a_F_slotsync_failure_callback_0), int32(1339), int32(_a_F_slotsync_failure_callback_1))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						v28 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v28)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v27)+16)), uint32(v28))
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v28)
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
						m.T0[v39].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					v28 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v28)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v27)+16)), uint32(v28))
					*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v28)
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
					m.T0[v39].(func(*base.Module, int32))(m, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
				m.T0[v39].(func(*base.Module, int32))(m, l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_smgrdounlinkall(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int64
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int64
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int64
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int64
	_ = v511
	var v536 int32
	_ = v536
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int64
	_ = v620
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v784 int32
	_ = v784
	var v803 int32
	_ = v803
	var v822 int32
	_ = v822
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int64
	_ = v870
	var v873 int32
	_ = v873
	var v874 int64
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v904 int32
	_ = v904
	var v916 int32
	_ = v916
	var v917 int64
	_ = v917
	var v919 int64
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v949 int64
	_ = v949
	var v951 int64
	_ = v951
	var v957 int32
	_ = v957
	var v958 int64
	_ = v958
	var v960 int64
	_ = v960
	var v966 int32
	_ = v966
	var v967 int64
	_ = v967
	var v969 int64
	_ = v969
	var v975 int32
	_ = v975
	var v976 int64
	_ = v976
	var v978 int64
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_smgrdounlinkall_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v24 + int32(1)
	v28 = m.G0
	v30 = v28 - int32(112)
	m.G0 = v30
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v20 + int32(80)
	return
L4:
	;
	v34 = F_palloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v30 + int32(112)
	v845 = F_palloc(m, l1<<(uint(int32(4))%32))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L7
	} else {
		goto L171
	}
L7:
	;
	return
L8:
	;
	if l1 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v34)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L7
	} else {
		goto L170
	}
L10:
	;
	v42 = v4
	v45 = v4
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+v42<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v59 != int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v154 == int32(0) {
		goto L9
	} else {
		goto L31
	}
L13:
	;
	v165 = v42 + int32(1)
	if v165 != l1 {
		v42 = v165
		v45 = v154
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[1]))
	if v59 != v63 {
		v154 = v45
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+v45<<(uint(int32(2))%32)))) = v58
	v154 = v45 + int32(1)
	goto L13
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v65
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+72)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	if int32(0) < v70 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[3]))
	v76 = v30 + int32(72)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v84 = int32(0)
	v87 = v70
	v89 = v74
	goto L21
L19:
	;
	goto L20
L20:
	;
	v154 = v45
	goto L13
L21:
	;
	v100 = v89 + v84<<(uint(int32(6))%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
	if v101&int32(33554432) == int32(0) {
		v119 = v87
		v120 = v89
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v122 = v84 + int32(1)
	if v122 < v119 {
		v84 = v122
		v87 = v119
		v89 = v120
		goto L21
	} else {
		goto L29
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v106 != v79 {
		v119 = v87
		v120 = v89
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v108 != v78 {
		v119 = v87
		v120 = v89
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v110 != v77 {
		v119 = v87
		v120 = v89
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_InvalidateLocalBuffer(m, v100, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[3]))
	v119 = v116
	v120 = v118
	goto L23
L29:
	;
	goto L22
L30:
	;
	goto L12
L31:
	;
	v171 = F_palloc(m, v154<<(uint(int32(4))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if int32(0) < v154 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_pfree(m, v171)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L7
	} else {
		goto L169
	}
L34:
	;
	F_pfree(m, v171)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L7
	} else {
		goto L110
	}
L35:
	;
	v179 = int32(0)
	v192 = int64(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if base.Ui32(int32(63)) <= base.Ui32(v398+int32(31)) {
		goto L33
	} else {
		goto L109
	}
L38:
	;
	v195 = v171 + v179<<(uint(int32(4))%32)
	v198 = v34 + v179<<(uint(int32(2))%32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v202 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	v308 = base.I32_div_s(v306, int32(32))
	if base.Ui64(base.I64_extend_i32_s(v308)) <= base.Ui64(v300) {
		goto L34
	} else {
		goto L89
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v213
	if v213 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v199+int32(0))+20))
	if v208 != int32(-1) {
		v213 = v208
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v213 = int32(-1)
	goto L41
L45:
	;
	goto L44
L46:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v229 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v219 = F_smgrexists(m, v217, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v225 = v192 + base.I64_extend_i32_u(v213)
	goto L46
L50:
	;
	if v219 == int32(0) {
		v225 = v192
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L34
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v240
	if v240 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v226+int32(4))+20))
	if v235 != int32(-1) {
		v240 = v235
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v240 = int32(-1)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v254 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v250 = v225 + base.I64_extend_i32_u(v240)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v248 = F_smgrexists(m, v246, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	if v248 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	v250 = v225
	goto L58
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v265
	if v265 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	goto L64
L66:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(8))+20))
	if v260 != int32(-1) {
		v265 = v260
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v265 = int32(-1)
	goto L65
L69:
	;
	goto L68
L70:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v279 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v275 = v250 + base.I64_extend_i32_u(v265)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v273 = F_smgrexists(m, v271, int32(2))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	if v273 != 0 {
		goto L34
	} else {
		goto L75
	}
L75:
	;
	v275 = v250
	goto L70
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = v290
	if v290 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v276+int32(12))+20))
	if v285 != int32(-1) {
		v290 = v285
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v290 = int32(-1)
	goto L77
L81:
	;
	goto L80
L82:
	;
	v302 = v179 + int32(1)
	if v302 < v154 {
		v179 = v302
		v192 = v300
		goto L38
	} else {
		goto L88
	}
L83:
	;
	v300 = v275 + base.I64_extend_i32_u(v290)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v298 = F_smgrexists(m, v296, int32(3))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	if v298 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	v300 = v275
	goto L82
L88:
	;
	goto L39
L89:
	;
	v315 = int32(0)
	goto L90
L90:
	;
	v330 = v34 + v315<<(uint(int32(2))%32)
	v333 = v171 + v315<<(uint(int32(4))%32)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v334 != int32(-1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L33
L92:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v338
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v337)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v340
	v344 = int32(0)
	F_FindAndDropRelationBuffers(m, v30+int32(56), v344, v334, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v349 != int32(-1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v353
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v352)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v355
	F_FindAndDropRelationBuffers(m, v30+int32(40), int32(1), v349, int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	if v364 != int32(-1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v370
	F_FindAndDropRelationBuffers(m, v30+int32(24), int32(2), v364, int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	if v379 != int32(-1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v383
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v382)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v385
	F_FindAndDropRelationBuffers(m, v30+int32(8), int32(3), v379, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v395 = v315 + int32(1)
	if v395 != v154 {
		v315 = v395
		goto L90
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	goto L91
L109:
	;
	goto L34
L110:
	;
	v424 = F_palloc(m, v154*int32(12))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v426 = int32(0)
	if v154 <= v426 {
		v555 = v426
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if int32(0) < v557 {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	v429 = int32(0)
	if v154 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if base.Ui32(v154) < base.Ui32(int32(21)) {
		v555 = int32(0)
		goto L112
	} else {
		goto L122
	}
L115:
	;
	v440 = int32(0)
	v441 = v429
	goto L118
L116:
	;
	v489 = v429
	goto L117
L117:
	;
	v504 = v424 + v489*int32(12)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v34+v489<<(uint(int32(2))%32))))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+8)) = v509
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v508)))
	*(*int64)(unsafe.Add(mBase, uint32(v504))) = v511
	goto L114
L118:
	;
	v454 = int32(12)
	v456 = v424 + v441*v454
	v457 = int32(2)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v34+v441<<(uint(v457)%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v456)+8)) = v461
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v460)))
	*(*int64)(unsafe.Add(mBase, uint32(v456))) = v463
	v466 = v441 | int32(1)
	v469 = v424 + v466*v454
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v34+v466<<(uint(v457)%32))))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+8)) = v474
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v473)))
	*(*int64)(unsafe.Add(mBase, uint32(v469))) = v476
	v479 = v441 + v457
	v481 = v440 + v457
	if v481 != v154&int32(2147483646) {
		v440 = v481
		v441 = v479
		goto L118
	} else {
		goto L120
	}
L119:
	;
	if v154&int32(1) == int32(0) {
		goto L114
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v489 = v479
	goto L117
L122:
	;
	F_pg_qsort(m, v424, v154, int32(12), int32(1078))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	v555 = int32(1)
	goto L112
L124:
	;
	v569 = int32(0)
	goto L127
L125:
	;
	goto L126
L126:
	;
	F_pfree(m, v424)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L7
	} else {
		goto L168
	}
L127:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[6]))
	v582 = v579 + v569<<(uint(int32(6))%32)
	if v555 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L126
L129:
	;
	v762 = v569 + int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if v762 < v764 {
		v569 = v762
		goto L127
	} else {
		goto L167
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = int32(_a_F_smgrdounlinkall_1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = int32(_a_F_smgrdounlinkall_2)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+100)) = int32(_a_F_smgrdounlinkall_3)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+88)) = int64(0)
	v659 = int32(_a_F_smgrdounlinkall_4)
	v661 = base.AtomicRmwOr32(m, v582, int32(24), v659)
	if v661&v659 != 0 {
		goto L144
	} else {
		goto L145
	}
L131:
	;
	if v154 <= int32(0) {
		goto L129
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v582)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v621
	*(*int64)(unsafe.Add(mBase, uint32(v30)+88)) = v620
	v628 = F_bsearch(m, v30+int32(88), v424, v154, int32(12), int32(1078))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L7
	} else {
		goto L142
	}
L134:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v593 = int32(0)
	goto L135
L135:
	;
	v608 = v424 + v593*int32(12)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	if v587 != v609 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L129
L137:
	;
	v618 = v593 + int32(1)
	if v618 != v154 {
		v593 = v618
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v611 != v612 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v608)+8))
	if v614 == v615 {
		v638 = v608
		goto L130
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	goto L136
L142:
	;
	if v628 == int32(0) {
		goto L129
	} else {
		goto L143
	}
L143:
	;
	v638 = v628
	goto L130
L144:
	;
	goto L147
L145:
	;
	v694 = v661
	goto L146
L146:
	;
	v710 = int32(_a_F_smgrdounlinkall_5)
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(88))+8))
	if v713 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L147:
	;
	F_perform_spin_delay(m, v30+int32(88))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L7
	} else {
		goto L149
	}
L148:
	;
	v694 = v687
	goto L146
L149:
	;
	v685 = int32(_a_F_smgrdounlinkall_4)
	v687 = base.AtomicRmwOr32(m, v582, int32(24), v685)
	if v687&v685 != 0 {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	if v730 != v731 {
		goto L162
	} else {
		goto L163
	}
L152:
	;
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7])) = v728
	goto L152
L154:
	;
	if int32(999) < v711 {
		goto L152
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v711 < int32(11) {
		goto L152
	} else {
		goto L161
	}
L157:
	;
	v718 = int32(900)
	if v718 <= v711 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v721 = v718
	goto L160
L159:
	;
	v721 = v711
	goto L160
L160:
	;
	v728 = v721 + int32(100)
	goto L153
L161:
	;
	v728 = v711 - int32(1)
	goto L153
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+24)) = v694 & int32(-4194305)
	goto L129
L163:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v638)+4))
	if v733 != v734 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v582)+8))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v638)+8))
	if v736 != v737 {
		goto L162
	} else {
		goto L165
	}
L165:
	;
	F_InvalidateBuffer(m, v582)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L166
	}
L166:
	;
	goto L129
L167:
	;
	goto L128
L168:
	;
	goto L9
L169:
	;
	goto L9
L170:
	;
	goto L6
L171:
	;
	if int32(0) < l1 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v863 = v4
	goto L175
L173:
	;
	goto L174
L174:
	;
	F_pfree(m, v845)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L7
	} else {
		goto L193
	}
L175:
	;
	v868 = l0 + v863<<(uint(int32(2))%32)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v870 = *(*int64)(unsafe.Add(mBase, uint32(v869)))
	v873 = v845 + v863<<(uint(int32(4))%32)
	v874 = *(*int64)(unsafe.Add(mBase, uint32(v869)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v873)+8)) = v874
	*(*int64)(unsafe.Add(mBase, uint32(v873))) = v870
	v877 = int32(0)
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	F_mdclose(m, v878, v877)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L7
	} else {
		goto L177
	}
L176:
	;
	v904 = v877
	goto L182
L177:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	F_mdclose(m, v882, int32(1))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	F_mdclose(m, v886, int32(2))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L7
	} else {
		goto L179
	}
L179:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	F_mdclose(m, v890, int32(3))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	v895 = v863 + int32(1)
	if v895 != l1 {
		v863 = v895
		goto L175
	} else {
		goto L181
	}
L181:
	;
	goto L176
L182:
	;
	v916 = v845 + v904<<(uint(int32(4))%32)
	v917 = *(*int64)(unsafe.Add(mBase, uint32(v916)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+72)) = v917
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v916)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+64)) = v919
	F_CacheInvalidateSmgr(m, v20-int32(-64))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L7
	} else {
		goto L184
	}
L183:
	;
	v943 = int32(0)
	goto L186
L184:
	;
	v926 = v904 + int32(1)
	if v926 != l1 {
		v904 = v926
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v948 = v845 + v943<<(uint(int32(4))%32)
	v949 = *(*int64)(unsafe.Add(mBase, uint32(v948)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v949
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v951
	F_mdunlink(m, v20+int32(48), int32(0), l2)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L7
	} else {
		goto L188
	}
L187:
	;
	goto L174
L188:
	;
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v948)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v958
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v960
	F_mdunlink(m, v20+int32(32), int32(1), l2)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L7
	} else {
		goto L189
	}
L189:
	;
	v967 = *(*int64)(unsafe.Add(mBase, uint32(v948)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v967
	v969 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v969
	F_mdunlink(m, v20+int32(16), int32(2), l2)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v948)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v976
	v978 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v978
	F_mdunlink(m, v20, int32(3), l2)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L7
	} else {
		goto L191
	}
L191:
	;
	v984 = v943 + int32(1)
	if v984 != l1 {
		v943 = v984
		goto L186
	} else {
		goto L192
	}
L192:
	;
	goto L187
L193:
	;
	v1005 = int32(_a_F_smgrdounlinkall_0)
	v1007 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v1007 - int32(1)
	goto L3
}
func F_smgrpin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v3 != 0 {
		v11 = v3
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v5
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v11 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v11 + int32(1)
	return
}
func F_smgrwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	v17 = int32(_a_F_smgrwritev_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v19 + int32(1)
	v23 = m.G0
	v25 = v23 - int32(1040)
	m.G0 = v25
	v28 = F__mdfd_getseg(m, l0, l1, l2, l4, int32(9))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = int32(1)
	v34 = l2 & int32(_a_F_smgrwritev_1)
	v35 = int32(_a_F_smgrwritev_2) - v34
	if base.Ui32(v31) < base.Ui32(v35) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v269 = int32(_a_F_smgrwritev_0)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v271 - int32(1)
	return
L4:
	;
	v38 = v31
	goto L6
L5:
	;
	v38 = v35
	goto L6
L6:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v38) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = int32(128)
	goto L9
L8:
	;
	v41 = v38
	goto L9
L9:
	;
	if v41 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = base.I64_extend_i32_u(v34 << (uint(int32(13)) % 32))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v47
	v51 = int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v38) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L59
	}
L13:
	;
	v57 = v47
	v62 = v25 + int32(16)
	v64 = v51
	v68 = int32(1)
	v71 = int32(0)
	goto L16
L14:
	;
	v121 = v51
	goto L15
L15:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v134 = F_FileWriteV(m, v130, v25+int32(16), v121, v46, int32(167772184))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L28
	}
L16:
	;
	v75 = l3 + v68<<(uint(int32(2))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v76 == v57+v77 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v121 = v109
	goto L15
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v93 != v90+v94 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v77 - int32(-8192)
	v90 = v57
	v91 = v62
	v92 = v64
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v76
	v90 = v76
	v91 = v62 + int32(8)
	v92 = v64 + int32(1)
	goto L18
L22:
	;
	v110 = int32(2)
	v113 = v71 + v110
	if v113 != 0 {
		v57 = v107
		v62 = v108
		v64 = v109
		v68 = v68 + v110
		v71 = v113
		goto L16
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v93
	v107 = v93
	v108 = v91 + int32(8)
	v109 = v92 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v94 - int32(-8192)
	v107 = v90
	v108 = v91
	v109 = v92
	goto L22
L26:
	;
	goto L17
L27:
	;
	if l4 != 0 {
		goto L55
	} else {
		goto L56
	}
L28:
	;
	if int32(0) <= v134 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = int32(0)
	v144 = v134
	v146 = v121
	v154 = v46
	goto L32
L30:
	;
	goto L31
L31:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[1]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L46
	}
L32:
	;
	v155 = v139 + v144
	if v155 == int32(_a_F_smgrwritev_3) {
		goto L27
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v159 = v25 + int32(16)
	v162 = v159
	v163 = v146
	v164 = v144
	goto L37
L35:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v194 = v154 + base.I64_extend_i32_u(v144)
	v196 = F_FileWriteV(m, v192, v159, v191, v194, int32(167772184))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L44
	}
L36:
	;
	if v159 == v162 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if base.Ui32(v164) < base.Ui32(v166) {
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v191 = int32(0)
	goto L35
L39:
	;
	v172 = v163 - int32(1)
	if v172 != 0 {
		v162 = v162 + int32(8)
		v163 = v172
		v164 = v164 - v166
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v181 + v164
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v184 - v164
	v191 = v163
	goto L35
L42:
	;
	v176 = v163 << (uint(int32(3)) % 32)
	if v176 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	base.MemoryCopy(m, v159, v162, v176)
	goto L41
L44:
	;
	if int32(0) <= v196 {
		v139 = v155
		v144 = v196
		v146 = v191
		v154 = v194
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[2]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226+v224*int32(48))+32))
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l2
	F_errmsg(m, int32(_a_F_smgrwritev_4), v25)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v217 == int32(51) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errhint(m, int32(_a_F_smgrwritev_5), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_errfinish(m, int32(_a_F_smgrwritev_6), int32(1134), int32(_a_F_smgrwritev_7))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	m.G0 = v25 + int32(1040)
	goto L3
L56:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v248 != int32(-1) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_register_dirty_segment(m, l0, l1, v28)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	F_errmsg_internal(m, int32(_a_F_smgrwritev_8), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_smgrwritev_6), int32(1092), int32(_a_F_smgrwritev_7))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v11 = F_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v11
	}
}
func F_sort_asc(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14016(m, l0, int32(_a_F_sort_asc_0), int32(234), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_sort_checkpoint_bufferids(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
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
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v644 int32
	_ = v644
	var v646 int64
	_ = v646
	var v648 int64
	_ = v648
	var v650 int32
	_ = v650
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int64
	_ = v703
	var v705 int64
	_ = v705
	var v707 int32
	_ = v707
	var v709 int64
	_ = v709
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v715 int64
	_ = v715
	var v717 int64
	_ = v717
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int64
	_ = v772
	var v774 int64
	_ = v774
	var v776 int32
	_ = v776
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v782 int32
	_ = v782
	var v784 int64
	_ = v784
	var v786 int64
	_ = v786
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int64
	_ = v839
	var v841 int64
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int64
	_ = v846
	var v848 int64
	_ = v848
	var v850 int32
	_ = v850
	var v852 int64
	_ = v852
	var v854 int64
	_ = v854
	var v857 int32
	_ = v857
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int64
	_ = v904
	var v906 int64
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int64
	_ = v911
	var v913 int64
	_ = v913
	var v915 int32
	_ = v915
	var v917 int64
	_ = v917
	var v919 int64
	_ = v919
	var v922 int32
	_ = v922
	var v956 int32
	_ = v956
	var v958 int64
	_ = v958
	var v960 int64
	_ = v960
	var v962 int32
	_ = v962
	var v964 int64
	_ = v964
	var v966 int64
	_ = v966
	var v968 int32
	_ = v968
	var v970 int64
	_ = v970
	var v972 int64
	_ = v972
	var v974 int32
	_ = v974
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = l0
	v19 = l1
	goto L1
L1:
	;
	v32 = v18 + int32(20)
	v34 = v19
	goto L3
L2:
	;
	m.G0 = v16 + int32(32)
	return
L3:
	;
	v48 = v18 + v34*int32(20)
	if base.Ui32(v34) <= base.Ui32(int32(6)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	if base.Ui32(v34) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v142 = v32
	goto L26
L9:
	;
	v63 = v32
	goto L10
L10:
	;
	if base.Ui32(v63) <= base.Ui32(v18) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v138 = v63 + int32(20)
	if base.Ui32(v138) < base.Ui32(v48) {
		v63 = v138
		goto L10
	} else {
		goto L25
	}
L13:
	;
	v70 = v63
	goto L14
L14:
	;
	v81 = v70 - int32(20)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if base.Ui32(v82) < base.Ui32(v83) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	if base.Ui32(v83) < base.Ui32(v82) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v105
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v121
	if base.Ui32(v18) < base.Ui32(v81) {
		v70 = v81
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(16))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if base.Ui32(v88) < base.Ui32(v89) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v89) < base.Ui32(v88) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(12))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if v94 < v95 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v95 < v94 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(8))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if base.Ui32(v100) <= base.Ui32(v101) {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	goto L15
L25:
	;
	goto L11
L26:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(20))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if base.Ui32(v155) < base.Ui32(v156) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v187 = v18 + int32(base.Ui32(v34)>>(uint(int32(1))%32))*int32(20)
	if v34 != int32(7) {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	goto L27
L29:
	;
	v179 = v142 + int32(20)
	if base.Ui32(v179) < base.Ui32(v48) {
		v142 = v179
		goto L26
	} else {
		goto L37
	}
L30:
	;
	if base.Ui32(v156) < base.Ui32(v155) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(16))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if base.Ui32(v161) < base.Ui32(v162) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(v162) < base.Ui32(v161) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	if v167 < v168 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	if v168 < v167 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	if base.Ui32(v174) < base.Ui32(v173) {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	goto L5
L38:
	;
	v191 = v48 - int32(20)
	if base.Ui32(v34) < base.Ui32(int32(41)) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v633 = v187
	goto L40
L40:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v638
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v640
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v633)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v644
	v646 = *(*int64)(unsafe.Add(mBase, uint32(v633)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v646
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v633)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v648
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v633)+16)) = v650
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v633)+8)) = v652
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v633))) = v654
	v657 = v48 - int32(20)
	v660 = v657
	v661 = v32
	v663 = v32
	v665 = v657
	goto L233
L41:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if base.Ui32(v532) < base.Ui32(v533) {
		goto L191
	} else {
		goto L192
	}
L42:
	;
	v522 = v18
	v523 = v187
	v524 = v191
	goto L41
L43:
	;
	goto L44
L44:
	;
	v195 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v197 = v195 * int32(20)
	v198 = v18 + v197
	v201 = v18 + v195*int32(40)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if base.Ui32(v206) < base.Ui32(v207) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v307 = v195 * int32(-20)
	v308 = v187 + v307
	v309 = v187 + v197
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if base.Ui32(v314) < base.Ui32(v315) {
		goto L97
	} else {
		goto L98
	}
L46:
	;
	v305 = v18
	goto L45
L47:
	;
	v305 = v201
	goto L45
L48:
	;
	v305 = v285
	goto L45
L49:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if base.Ui32(v207) < base.Ui32(v255) {
		goto L75
	} else {
		goto L76
	}
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if base.Ui32(v207) < base.Ui32(v223) {
		v285 = v198
		goto L48
	} else {
		goto L58
	}
L51:
	;
	if base.Ui32(v207) < base.Ui32(v206) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if base.Ui32(v210) < base.Ui32(v211) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v211) < base.Ui32(v210) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	if v214 < v215 {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if v215 < v214 {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if base.Ui32(v219) <= base.Ui32(v218) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	if base.Ui32(v223) < base.Ui32(v207) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if base.Ui32(v206) < base.Ui32(v223) {
		goto L47
	} else {
		goto L66
	}
L60:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if base.Ui32(v226) < base.Ui32(v227) {
		v285 = v198
		goto L48
	} else {
		goto L61
	}
L61:
	;
	if base.Ui32(v227) < base.Ui32(v226) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v230 < v231 {
		v285 = v198
		goto L48
	} else {
		goto L63
	}
L63:
	;
	if v231 < v230 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if base.Ui32(v234) < base.Ui32(v235) {
		v285 = v198
		goto L48
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	if base.Ui32(v223) < base.Ui32(v206) {
		goto L46
	} else {
		goto L67
	}
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if base.Ui32(v241) < base.Ui32(v242) {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v242) < base.Ui32(v241) {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v245 < v246 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	if v246 < v245 {
		goto L46
	} else {
		goto L71
	}
L71:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if base.Ui32(v249) < base.Ui32(v250) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v252 = v201
	goto L74
L73:
	;
	v252 = v18
	goto L74
L74:
	;
	v305 = v252
	goto L45
L75:
	;
	if base.Ui32(v206) < base.Ui32(v255) {
		goto L46
	} else {
		goto L83
	}
L76:
	;
	if base.Ui32(v255) < base.Ui32(v207) {
		v285 = v198
		goto L48
	} else {
		goto L77
	}
L77:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if base.Ui32(v258) < base.Ui32(v259) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v259) < base.Ui32(v258) {
		v285 = v198
		goto L48
	} else {
		goto L79
	}
L79:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v262 < v263 {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	if v263 < v262 {
		v285 = v198
		goto L48
	} else {
		goto L81
	}
L81:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if base.Ui32(v267) < base.Ui32(v266) {
		v285 = v198
		goto L48
	} else {
		goto L82
	}
L82:
	;
	goto L75
L83:
	;
	if base.Ui32(v255) < base.Ui32(v206) {
		goto L47
	} else {
		goto L84
	}
L84:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if base.Ui32(v273) < base.Ui32(v274) {
		goto L46
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32(v274) < base.Ui32(v273) {
		goto L47
	} else {
		goto L86
	}
L86:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	if v277 < v278 {
		goto L46
	} else {
		goto L87
	}
L87:
	;
	if v278 < v277 {
		goto L47
	} else {
		goto L88
	}
L88:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	if base.Ui32(v281) < base.Ui32(v282) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v284 = v18
	goto L91
L90:
	;
	v284 = v201
	goto L91
L91:
	;
	v285 = v284
	goto L48
L92:
	;
	v416 = v191 + v195*int32(-40)
	v417 = v191 + v307
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	if base.Ui32(v422) < base.Ui32(v423) {
		goto L144
	} else {
		goto L145
	}
L93:
	;
	v413 = v308
	goto L92
L94:
	;
	v413 = v309
	goto L92
L95:
	;
	v413 = v393
	goto L92
L96:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if base.Ui32(v315) < base.Ui32(v363) {
		goto L122
	} else {
		goto L123
	}
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if base.Ui32(v315) < base.Ui32(v331) {
		v393 = v187
		goto L95
	} else {
		goto L105
	}
L98:
	;
	if base.Ui32(v315) < base.Ui32(v314) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if base.Ui32(v318) < base.Ui32(v319) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	if base.Ui32(v319) < base.Ui32(v318) {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v322 < v323 {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	if v323 < v322 {
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	if base.Ui32(v327) <= base.Ui32(v326) {
		goto L96
	} else {
		goto L104
	}
L104:
	;
	goto L97
L105:
	;
	if base.Ui32(v331) < base.Ui32(v315) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if base.Ui32(v314) < base.Ui32(v331) {
		goto L94
	} else {
		goto L113
	}
L107:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if base.Ui32(v334) < base.Ui32(v335) {
		v393 = v187
		goto L95
	} else {
		goto L108
	}
L108:
	;
	if base.Ui32(v335) < base.Ui32(v334) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v338 < v339 {
		v393 = v187
		goto L95
	} else {
		goto L110
	}
L110:
	;
	if v339 < v338 {
		goto L106
	} else {
		goto L111
	}
L111:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if base.Ui32(v342) < base.Ui32(v343) {
		v393 = v187
		goto L95
	} else {
		goto L112
	}
L112:
	;
	goto L106
L113:
	;
	if base.Ui32(v331) < base.Ui32(v314) {
		goto L93
	} else {
		goto L114
	}
L114:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if base.Ui32(v349) < base.Ui32(v350) {
		goto L94
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v350) < base.Ui32(v349) {
		goto L93
	} else {
		goto L116
	}
L116:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v353 < v354 {
		goto L94
	} else {
		goto L117
	}
L117:
	;
	if v354 < v353 {
		goto L93
	} else {
		goto L118
	}
L118:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if base.Ui32(v357) < base.Ui32(v358) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v360 = v309
	goto L121
L120:
	;
	v360 = v308
	goto L121
L121:
	;
	v413 = v360
	goto L92
L122:
	;
	if base.Ui32(v314) < base.Ui32(v363) {
		goto L93
	} else {
		goto L130
	}
L123:
	;
	if base.Ui32(v363) < base.Ui32(v315) {
		v393 = v187
		goto L95
	} else {
		goto L124
	}
L124:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if base.Ui32(v366) < base.Ui32(v367) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	if base.Ui32(v367) < base.Ui32(v366) {
		v393 = v187
		goto L95
	} else {
		goto L126
	}
L126:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v370 < v371 {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	if v371 < v370 {
		v393 = v187
		goto L95
	} else {
		goto L128
	}
L128:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if base.Ui32(v375) < base.Ui32(v374) {
		v393 = v187
		goto L95
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	if base.Ui32(v363) < base.Ui32(v314) {
		goto L94
	} else {
		goto L131
	}
L131:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if base.Ui32(v381) < base.Ui32(v382) {
		goto L93
	} else {
		goto L132
	}
L132:
	;
	if base.Ui32(v382) < base.Ui32(v381) {
		goto L94
	} else {
		goto L133
	}
L133:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v385 < v386 {
		goto L93
	} else {
		goto L134
	}
L134:
	;
	if v386 < v385 {
		goto L94
	} else {
		goto L135
	}
L135:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if base.Ui32(v389) < base.Ui32(v390) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v392 = v308
	goto L138
L137:
	;
	v392 = v309
	goto L138
L138:
	;
	v393 = v392
	goto L95
L139:
	;
	v522 = v305
	v523 = v413
	v524 = v521
	goto L41
L140:
	;
	v521 = v416
	goto L139
L141:
	;
	v521 = v191
	goto L139
L142:
	;
	v521 = v501
	goto L139
L143:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if base.Ui32(v423) < base.Ui32(v471) {
		goto L169
	} else {
		goto L170
	}
L144:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if base.Ui32(v423) < base.Ui32(v439) {
		v501 = v417
		goto L142
	} else {
		goto L152
	}
L145:
	;
	if base.Ui32(v423) < base.Ui32(v422) {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if base.Ui32(v426) < base.Ui32(v427) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	if base.Ui32(v427) < base.Ui32(v426) {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	if v430 < v431 {
		goto L144
	} else {
		goto L149
	}
L149:
	;
	if v431 < v430 {
		goto L143
	} else {
		goto L150
	}
L150:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	if base.Ui32(v435) <= base.Ui32(v434) {
		goto L143
	} else {
		goto L151
	}
L151:
	;
	goto L144
L152:
	;
	if base.Ui32(v439) < base.Ui32(v423) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v422) < base.Ui32(v439) {
		goto L141
	} else {
		goto L160
	}
L154:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(v442) < base.Ui32(v443) {
		v501 = v417
		goto L142
	} else {
		goto L155
	}
L155:
	;
	if base.Ui32(v443) < base.Ui32(v442) {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v446 < v447 {
		v501 = v417
		goto L142
	} else {
		goto L157
	}
L157:
	;
	if v447 < v446 {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if base.Ui32(v450) < base.Ui32(v451) {
		v501 = v417
		goto L142
	} else {
		goto L159
	}
L159:
	;
	goto L153
L160:
	;
	if base.Ui32(v439) < base.Ui32(v422) {
		goto L140
	} else {
		goto L161
	}
L161:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(v457) < base.Ui32(v458) {
		goto L141
	} else {
		goto L162
	}
L162:
	;
	if base.Ui32(v458) < base.Ui32(v457) {
		goto L140
	} else {
		goto L163
	}
L163:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v461 < v462 {
		goto L141
	} else {
		goto L164
	}
L164:
	;
	if v462 < v461 {
		goto L140
	} else {
		goto L165
	}
L165:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if base.Ui32(v465) < base.Ui32(v466) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v468 = v191
	goto L168
L167:
	;
	v468 = v416
	goto L168
L168:
	;
	v521 = v468
	goto L139
L169:
	;
	if base.Ui32(v422) < base.Ui32(v471) {
		goto L140
	} else {
		goto L177
	}
L170:
	;
	if base.Ui32(v471) < base.Ui32(v423) {
		v501 = v417
		goto L142
	} else {
		goto L171
	}
L171:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(v474) < base.Ui32(v475) {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	if base.Ui32(v475) < base.Ui32(v474) {
		v501 = v417
		goto L142
	} else {
		goto L173
	}
L173:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v478 < v479 {
		goto L169
	} else {
		goto L174
	}
L174:
	;
	if v479 < v478 {
		v501 = v417
		goto L142
	} else {
		goto L175
	}
L175:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if base.Ui32(v483) < base.Ui32(v482) {
		v501 = v417
		goto L142
	} else {
		goto L176
	}
L176:
	;
	goto L169
L177:
	;
	if base.Ui32(v471) < base.Ui32(v422) {
		goto L141
	} else {
		goto L178
	}
L178:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(v489) < base.Ui32(v490) {
		goto L140
	} else {
		goto L179
	}
L179:
	;
	if base.Ui32(v490) < base.Ui32(v489) {
		goto L141
	} else {
		goto L180
	}
L180:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if v493 < v494 {
		goto L140
	} else {
		goto L181
	}
L181:
	;
	if v494 < v493 {
		goto L141
	} else {
		goto L182
	}
L182:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if base.Ui32(v497) < base.Ui32(v498) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v500 = v416
	goto L185
L184:
	;
	v500 = v191
	goto L185
L185:
	;
	v501 = v500
	goto L142
L186:
	;
	v633 = v631
	goto L40
L187:
	;
	v631 = v522
	goto L186
L188:
	;
	v631 = v524
	goto L186
L189:
	;
	v631 = v611
	goto L186
L190:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if base.Ui32(v533) < base.Ui32(v581) {
		goto L216
	} else {
		goto L217
	}
L191:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if base.Ui32(v533) < base.Ui32(v549) {
		v611 = v523
		goto L189
	} else {
		goto L199
	}
L192:
	;
	if base.Ui32(v533) < base.Ui32(v532) {
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if base.Ui32(v536) < base.Ui32(v537) {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	if base.Ui32(v537) < base.Ui32(v536) {
		goto L190
	} else {
		goto L195
	}
L195:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	if v540 < v541 {
		goto L191
	} else {
		goto L196
	}
L196:
	;
	if v541 < v540 {
		goto L190
	} else {
		goto L197
	}
L197:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	if base.Ui32(v545) <= base.Ui32(v544) {
		goto L190
	} else {
		goto L198
	}
L198:
	;
	goto L191
L199:
	;
	if base.Ui32(v549) < base.Ui32(v533) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if base.Ui32(v532) < base.Ui32(v549) {
		goto L188
	} else {
		goto L207
	}
L201:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if base.Ui32(v552) < base.Ui32(v553) {
		v611 = v523
		goto L189
	} else {
		goto L202
	}
L202:
	;
	if base.Ui32(v553) < base.Ui32(v552) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if v556 < v557 {
		v611 = v523
		goto L189
	} else {
		goto L204
	}
L204:
	;
	if v557 < v556 {
		goto L200
	} else {
		goto L205
	}
L205:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	if base.Ui32(v560) < base.Ui32(v561) {
		v611 = v523
		goto L189
	} else {
		goto L206
	}
L206:
	;
	goto L200
L207:
	;
	if base.Ui32(v549) < base.Ui32(v532) {
		goto L187
	} else {
		goto L208
	}
L208:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if base.Ui32(v567) < base.Ui32(v568) {
		goto L188
	} else {
		goto L209
	}
L209:
	;
	if base.Ui32(v568) < base.Ui32(v567) {
		goto L187
	} else {
		goto L210
	}
L210:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if v571 < v572 {
		goto L188
	} else {
		goto L211
	}
L211:
	;
	if v572 < v571 {
		goto L187
	} else {
		goto L212
	}
L212:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	if base.Ui32(v575) < base.Ui32(v576) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v578 = v524
	goto L215
L214:
	;
	v578 = v522
	goto L215
L215:
	;
	v631 = v578
	goto L186
L216:
	;
	if base.Ui32(v532) < base.Ui32(v581) {
		goto L187
	} else {
		goto L224
	}
L217:
	;
	if base.Ui32(v581) < base.Ui32(v533) {
		v611 = v523
		goto L189
	} else {
		goto L218
	}
L218:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if base.Ui32(v584) < base.Ui32(v585) {
		goto L216
	} else {
		goto L219
	}
L219:
	;
	if base.Ui32(v585) < base.Ui32(v584) {
		v611 = v523
		goto L189
	} else {
		goto L220
	}
L220:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if v588 < v589 {
		goto L216
	} else {
		goto L221
	}
L221:
	;
	if v589 < v588 {
		v611 = v523
		goto L189
	} else {
		goto L222
	}
L222:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	if base.Ui32(v593) < base.Ui32(v592) {
		v611 = v523
		goto L189
	} else {
		goto L223
	}
L223:
	;
	goto L216
L224:
	;
	if base.Ui32(v581) < base.Ui32(v532) {
		goto L188
	} else {
		goto L225
	}
L225:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if base.Ui32(v599) < base.Ui32(v600) {
		goto L187
	} else {
		goto L226
	}
L226:
	;
	if base.Ui32(v600) < base.Ui32(v599) {
		goto L188
	} else {
		goto L227
	}
L227:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if v603 < v604 {
		goto L187
	} else {
		goto L228
	}
L228:
	;
	if v604 < v603 {
		goto L188
	} else {
		goto L229
	}
L229:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	if base.Ui32(v607) < base.Ui32(v608) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v610 = v522
	goto L232
L231:
	;
	v610 = v524
	goto L232
L232:
	;
	v611 = v610
	goto L189
L233:
	;
	if base.Ui32(v660) < base.Ui32(v661) {
		v730 = v661
		v732 = v663
		goto L235
	} else {
		goto L236
	}
L235:
	;
	if base.Ui32(v730) <= base.Ui32(v660) {
		goto L250
	} else {
		goto L251
	}
L236:
	;
	v675 = v661
	v677 = v663
	goto L237
L237:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.Ui32(v685) < base.Ui32(v686) {
		v722 = v677
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v730 = v725
	v732 = v722
	goto L235
L239:
	;
	v725 = v675 + int32(20)
	if base.Ui32(v725) <= base.Ui32(v660) {
		v675 = v725
		v677 = v722
		goto L237
	} else {
		goto L248
	}
L240:
	;
	if base.Ui32(v686) < base.Ui32(v685) {
		v730 = v675
		v732 = v677
		goto L235
	} else {
		goto L241
	}
L241:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if base.Ui32(v689) < base.Ui32(v690) {
		v722 = v677
		goto L239
	} else {
		goto L242
	}
L242:
	;
	if base.Ui32(v690) < base.Ui32(v689) {
		v730 = v675
		v732 = v677
		goto L235
	} else {
		goto L243
	}
L243:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v675)+8))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v693 < v694 {
		v722 = v677
		goto L239
	} else {
		goto L244
	}
L244:
	;
	if v694 < v693 {
		v730 = v675
		v732 = v677
		goto L235
	} else {
		goto L245
	}
L245:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v675)+12))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if base.Ui32(v697) < base.Ui32(v698) {
		v722 = v677
		goto L239
	} else {
		goto L246
	}
L246:
	;
	if base.Ui32(v698) < base.Ui32(v697) {
		v730 = v675
		v732 = v677
		goto L235
	} else {
		goto L247
	}
L247:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v677)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v701
	v703 = *(*int64)(unsafe.Add(mBase, uint32(v677)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v703
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v677)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v705
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v675)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v677)+16)) = v707
	v709 = *(*int64)(unsafe.Add(mBase, uint32(v675)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v677)+8)) = v709
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v675)))
	*(*int64)(unsafe.Add(mBase, uint32(v677))) = v711
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v675)+16)) = v713
	v715 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v675)+8)) = v715
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v675))) = v717
	v722 = v677 + int32(20)
	goto L239
L248:
	;
	goto L238
L249:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v730)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v956
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v730)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v958
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v730)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v960
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v743)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+16)) = v962
	v964 = *(*int64)(unsafe.Add(mBase, uint32(v743)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v730)+8)) = v964
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v743)))
	*(*int64)(unsafe.Add(mBase, uint32(v730))) = v966
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v743)+16)) = v968
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v743)+8)) = v970
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v743))) = v972
	v974 = int32(20)
	v660 = v743 - v974
	v661 = v730 + v974
	v663 = v732
	v665 = v748
	goto L233
L250:
	;
	v743 = v660
	v748 = v665
	goto L253
L251:
	;
	v798 = v660
	v803 = v665
	goto L252
L252:
	;
	v810 = int32(20)
	v811 = base.I32_div_s(v732-v18, v810)
	v814 = base.I32_div_s(v730-v732, v810)
	if v811 < v814 {
		goto L265
	} else {
		goto L266
	}
L253:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.Ui32(v754) < base.Ui32(v755) {
		goto L249
	} else {
		goto L255
	}
L254:
	;
	v798 = v794
	v803 = v792
	goto L252
L255:
	;
	if base.Ui32(v755) < base.Ui32(v754) {
		v792 = v748
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v794 = v743 - int32(20)
	if base.Ui32(v730) <= base.Ui32(v794) {
		v743 = v794
		v748 = v792
		goto L253
	} else {
		goto L264
	}
L257:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if base.Ui32(v758) < base.Ui32(v759) {
		goto L249
	} else {
		goto L258
	}
L258:
	;
	if base.Ui32(v759) < base.Ui32(v758) {
		v792 = v748
		goto L256
	} else {
		goto L259
	}
L259:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v743)+8))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v762 < v763 {
		goto L249
	} else {
		goto L260
	}
L260:
	;
	if v763 < v762 {
		v792 = v748
		goto L256
	} else {
		goto L261
	}
L261:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v743)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if base.Ui32(v766) < base.Ui32(v767) {
		goto L249
	} else {
		goto L262
	}
L262:
	;
	if base.Ui32(v767) < base.Ui32(v766) {
		v792 = v748
		goto L256
	} else {
		goto L263
	}
L263:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v743)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v770
	v772 = *(*int64)(unsafe.Add(mBase, uint32(v743)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v772
	v774 = *(*int64)(unsafe.Add(mBase, uint32(v743)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v774
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v748)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v743)+16)) = v776
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v748)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v743)+8)) = v778
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v748)))
	*(*int64)(unsafe.Add(mBase, uint32(v743))) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v748)+16)) = v782
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v748)+8)) = v784
	v786 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v748))) = v786
	v792 = v748 - int32(20)
	goto L256
L264:
	;
	goto L254
L265:
	;
	v816 = v811
	goto L267
L266:
	;
	v816 = v814
	goto L267
L267:
	;
	if v816 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v826 = int32(0)
	goto L271
L269:
	;
	goto L270
L270:
	;
	v873 = int32(20)
	v874 = base.I32_div_s(v803-v798, v873)
	v877 = base.I32_div_s(v48-v803, v873)
	v879 = v877 - int32(1)
	if v874 < v879 {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	v835 = v826 * int32(20)
	v836 = v18 + v835
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v837
	v839 = *(*int64)(unsafe.Add(mBase, uint32(v836)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v839
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v836)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v841
	v843 = v835 + (v730 + v816*int32(-20))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v843)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v836)+16)) = v844
	v846 = *(*int64)(unsafe.Add(mBase, uint32(v843)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v836)+8)) = v846
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v843)))
	*(*int64)(unsafe.Add(mBase, uint32(v836))) = v848
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+16)) = v850
	v852 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v843)+8)) = v852
	v854 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v843))) = v854
	v857 = v826 + int32(1)
	if v857 != v816 {
		v826 = v857
		goto L271
	} else {
		goto L273
	}
L272:
	;
	goto L270
L273:
	;
	goto L272
L274:
	;
	v881 = v874
	goto L276
L275:
	;
	v881 = v879
	goto L276
L276:
	;
	if v881 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v892 = int32(0)
	goto L280
L278:
	;
	goto L279
L279:
	;
	if base.Ui32(v814) <= base.Ui32(v874) {
		goto L283
	} else {
		goto L284
	}
L280:
	;
	v900 = v892 * int32(20)
	v901 = v730 + v900
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v901)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v902
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v901)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v904
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v901)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v906
	v908 = v900 + (v48 + v881*int32(-20))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v901)+16)) = v909
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v908)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v901)+8)) = v911
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v908)))
	*(*int64)(unsafe.Add(mBase, uint32(v901))) = v913
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v908)+16)) = v915
	v917 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v908)+8)) = v917
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v908))) = v919
	v922 = v892 + int32(1)
	if v922 != v881 {
		v892 = v922
		goto L280
	} else {
		goto L282
	}
L281:
	;
	goto L279
L282:
	;
	goto L281
L283:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v814) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v874) {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	F_sort_checkpoint_bufferids(m, v18, v814)
	mBase = m.M
	goto L288
L287:
	;
	goto L288
L288:
	;
	if base.Ui32(v874) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	v18 = v48 + v874*int32(-20)
	v19 = v874
	goto L1
L290:
	;
	F_sort_checkpoint_bufferids(m, v48+v874*int32(-20), v874)
	mBase = m.M
	goto L292
L291:
	;
	goto L292
L292:
	;
	if base.Ui32(int32(1)) < base.Ui32(v814) {
		v34 = v814
		goto L3
	} else {
		goto L293
	}
L293:
	;
	goto L5
}
func F_sort_checkpoint_bufferids_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v9) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return l0
L2:
	;
	return l2
L3:
	;
	return v88
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v58) {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v25) {
		v88 = l1
		goto L3
	} else {
		goto L13
	}
L6:
	;
	if base.Ui32(v9) < base.Ui32(v8) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v12) < base.Ui32(v13) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(v13) < base.Ui32(v12) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 < v17 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v17 < v16 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v21) <= base.Ui32(v20) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	if base.Ui32(v25) < base.Ui32(v9) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui32(v8) < base.Ui32(v25) {
		goto L2
	} else {
		goto L21
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v28) < base.Ui32(v29) {
		v88 = l1
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v29) < base.Ui32(v28) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v32 < v33 {
		v88 = l1
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v33 < v32 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v36) < base.Ui32(v37) {
		v88 = l1
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	if base.Ui32(v25) < base.Ui32(v8) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v43) < base.Ui32(v44) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v44) < base.Ui32(v43) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v47 < v48 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	if v48 < v47 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v51) < base.Ui32(v52) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v54 = l2
	goto L29
L28:
	;
	v54 = l0
	goto L29
L29:
	;
	return v54
L30:
	;
	if base.Ui32(v8) < base.Ui32(v58) {
		goto L1
	} else {
		goto L38
	}
L31:
	;
	if base.Ui32(v58) < base.Ui32(v9) {
		v88 = l1
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v61) < base.Ui32(v62) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v62) < base.Ui32(v61) {
		v88 = l1
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v65 < v66 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v66 < v65 {
		v88 = l1
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v70) < base.Ui32(v69) {
		v88 = l1
		goto L3
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	if base.Ui32(v58) < base.Ui32(v8) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v76) < base.Ui32(v77) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v77) < base.Ui32(v76) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v80 < v81 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v81 < v80 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v84) < base.Ui32(v85) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v87 = l0
	goto L46
L45:
	;
	v87 = l2
	goto L46
L46:
	;
	v88 = v87
	goto L3
}
func F_sort_desc(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14016(m, l0, int32(_a_F_sort_desc_0), int32(244), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_sort_order_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 float32
	_ = v15
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
	v10 = *(*float32)(unsafe.Add(mBase, uint32(v7+v8)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v15 = *(*float32)(unsafe.Add(mBase, uint32(v12+v13)+8))
	if base.F32_lt(v10, v15) != 0 {
		v18 = int32(-1)
	} else {
		v18 = base.F32_gt(v10, v15)
	}
	return v18
}
func F_sortins(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(2) <= v10 {
		v13 = int32(2)
		v16 = F_palloc_extended(m, v10<<(uint(v13)%32), v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(101)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				if v24 != 0 {
					v26 = v24
				} else {
					v26 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v26
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v28 != 0 {
					v29 = v28
					v33 = int32(0)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v16+v33<<(uint(int32(2))%32)))) = v29
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						if v44 != 0 {
							v29 = v44
							v33 = v33 + int32(1)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				F_pg_qsort(m, v16, v10, int32(4), int32(970))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v60
					if v10 == int32(2) {
						v140 = int32(1)
					} else {
						v67 = int32(1)
						v69 = v10 - v67
						if v10 != int32(3) {
							v80 = int32(0)
							v83 = v67
							for {
								v88 = int32(2)
								v90 = v16 + v83<<(uint(v88)%32)
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v92 = int32(4)
								v93 = v90 + v92
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = v94
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v90-v92)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v98
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								v102 = v83 + v88
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v16+v102<<(uint(v88)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+24)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+28)) = v108
								if v80 != v10&int32(2147483646)-int32(4) {
									v80 = v80 + v88
									v83 = v102
									continue
								} else {
									break
								}
								break
							}
							if v10&int32(1) == int32(0) {
								v140 = v69
							} else {
								v119 = v102
								v126 = v16 + v119<<(uint(int32(2))%32)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v128
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v132
								v140 = v69
							}
						} else {
							v119 = v67
							v126 = v16 + v119<<(uint(int32(2))%32)
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v128
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v132
							v140 = v69
						}
					}
					v145 = v16 + v140<<(uint(int32(2))%32)
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = int32(0)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v145-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = v151
					F_pfree(m, v16)
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_spgadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(0)
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	v29 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+24)) = uint16(v29)
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 < v33 {
		v18 = v32
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(16)
	return
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v43 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(0)
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v48<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui32(int32(5)) <= base.Ui32(v58-int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L7
L12:
	;
	v92 = v48 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v92 < v93 {
		v48 = v92
		goto L10
	} else {
		goto L24
	}
L13:
	;
	if base.Ui32(v58-int32(6)) < base.Ui32(int32(2)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+24)) = uint8(v89)
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = l0
	v68 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)) = uint16(v68)
	goto L12
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgadjustmembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg(m, int32(_a_F_spgadjustmembers_1), v9)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_spgadjustmembers_2), int32(379), int32(_a_F_spgadjustmembers_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L11
}
func F_spgbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v287 int32
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	v11 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_palloc0(m, int32(_a_F_spgbeginscan_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) < l1 {
				v22 = F_palloc(m, l1*int32(48))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v25 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					F_initSpGistState(m, v16, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
						v36 = F_AllocSetContextCreateInternal(m, v31, int32(_a_F_spgbeginscan_1), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v36
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
							v45 = F_AllocSetContextCreateInternal(m, v40, int32(_a_F_spgbeginscan_4), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v45
								v49 = v16 + int32(20)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(4))%32))+88))
								if v50 != v56 {
									v58 = F_CreateTupleDescCopy(m, v51)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
										v64 = v58 + v61<<(uint(int32(4))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v64)+96)) = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v60
										v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v64)+92)) = uint16(v68)
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+6)))
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+102)) = uint8(v70)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+7)))
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+103)) = uint8(v72)
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)))
										v75 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v64)+116)) = v75
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+105)) = uint8(v75)
										*(*uint8)(unsafe.Add(mBase, uint32(v64)+104)) = uint8(v74)
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
										if v80 < int32(2) {
										} else {
											v84 = v58 + int32(20)
											v85 = int32(1)
											v86 = v80 - v85
											v87 = int32(7)
											v88 = v86 & v87
											if base.Ui32(v87) <= base.Ui32(v80-int32(2)) {
												v98 = v85
												v104 = int32(0)
												for {
													v109 = v84 + v98<<(uint(int32(4))%32)
													v110 = int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+80)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+64)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110
													v126 = int32(8)
													v127 = v98 + v126
													v129 = v104 + v126
													if v129 != v86&int32(-8) {
														v98 = v127
														v104 = v129
														continue
													} else {
														break
													}
													break
												}
												if v88 == int32(0) {
												} else {
													v134 = v127
													v145 = v134
													v147 = int32(0)
													for {
														*(*int32)(unsafe.Add(mBase, uint32(v84+v145<<(uint(int32(4))%32)))) = int32(-1)
														v159 = int32(1)
														v162 = v147 + v159
														if v162 != v88 {
															v145 = v145 + v159
															v147 = v162
															continue
														} else {
															break
														}
														break
													}
												}
											} else {
												v134 = v85
												v145 = v134
												v147 = int32(0)
												for {
													*(*int32)(unsafe.Add(mBase, uint32(v84+v145<<(uint(int32(4))%32)))) = int32(-1)
													v159 = int32(1)
													v162 = v147 + v159
													if v162 != v88 {
														v145 = v145 + v159
														v147 = v162
														continue
													} else {
														break
													}
													break
												}
											}
										}
										F_populate_compact_attribute(m, v58, int32(0))
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return int32(0)
										} else {
											v182 = v58
											*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v182
											*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v182
											v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											if v189 <= int32(0) {
												v278 = v16 + int32(132)
												v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
												mBase = m.M
												v282 = m.ExcPending
												if v282 != 0 {
													return int32(0)
												} else {
													v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
													v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
													v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
													v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
													v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
													*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
													*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
													*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
													v297 = v16 + int32(160)
													v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
													mBase = m.M
													v301 = m.ExcPending
													if v301 != 0 {
														return int32(0)
													} else {
														v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
														v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
														v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
														v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
														v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
														*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
														*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
														*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
														v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
														v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
														*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
														*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
														return v11
													}
												}
											} else {
												v194 = F_palloc(m, v189<<(uint(int32(2))%32))
												mBase = m.M
												v195 = m.ExcPending
												if v195 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v194
													v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v200 = F_palloc(m, v197<<(uint(int32(2))%32))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v200
														v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v206 = F_palloc(m, v203<<(uint(int32(3))%32))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v206
															v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															v212 = F_palloc(m, v209<<(uint(int32(3))%32))
															mBase = m.M
															v213 = m.ExcPending
															if v213 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v212
																v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																if int32(0) < v215 {
																	v220 = int32(0)
																	for {
																		v230 = v220 << (uint(int32(3)) % 32)
																		v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																		*(*int64)(unsafe.Add(mBase, uint32(v230+v231))) = int64(0)
																		v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																		*(*int64)(unsafe.Add(mBase, uint32(v235+v230))) = int64(9218868437227405312)
																		v240 = v220 + int32(1)
																		v241 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																		if v240 < v241 {
																			v220 = v240
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v246 = v241
																} else {
																	v246 = v215
																}
																v255 = F_palloc0(m, v246<<(uint(int32(2))%32))
																mBase = m.M
																v256 = m.ExcPending
																if v256 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v255
																	v258 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	v259 = F_palloc(m, v258)
																	mBase = m.M
																	v260 = m.ExcPending
																	if v260 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v259
																		v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																		if v262 == int32(0) {
																		} else {
																			base.MemoryFill(m, v259, int32(1), v262)
																		}
																		v278 = v16 + int32(132)
																		v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																		mBase = m.M
																		v282 = m.ExcPending
																		if v282 != 0 {
																			return int32(0)
																		} else {
																			v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																			v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
																			v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
																			*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
																			v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
																			*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
																			v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
																			*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
																			*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
																			*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
																			v297 = v16 + int32(160)
																			v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																			mBase = m.M
																			v301 = m.ExcPending
																			if v301 != 0 {
																				return int32(0)
																			} else {
																				v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																				v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
																				*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
																				v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
																				*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
																				v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
																				*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
																				v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
																				*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
																				*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
																				*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
																				v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																				v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
																				*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
																				return v11
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
								} else {
									v182 = v51
									*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v182
									*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v182
									v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									if v189 <= int32(0) {
										v278 = v16 + int32(132)
										v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
										mBase = m.M
										v282 = m.ExcPending
										if v282 != 0 {
											return int32(0)
										} else {
											v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
											v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
											v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
											v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
											v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
											*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
											*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
											*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
											v297 = v16 + int32(160)
											v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
											mBase = m.M
											v301 = m.ExcPending
											if v301 != 0 {
												return int32(0)
											} else {
												v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
												v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
												v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
												v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
												v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
												*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
												*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
												*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
												v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
												v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
												*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
												*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
												return v11
											}
										}
									} else {
										v194 = F_palloc(m, v189<<(uint(int32(2))%32))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v194
											v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v200 = F_palloc(m, v197<<(uint(int32(2))%32))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v200
												v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v206 = F_palloc(m, v203<<(uint(int32(3))%32))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v206
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v212 = F_palloc(m, v209<<(uint(int32(3))%32))
													mBase = m.M
													v213 = m.ExcPending
													if v213 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v212
														v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														if int32(0) < v215 {
															v220 = int32(0)
															for {
																v230 = v220 << (uint(int32(3)) % 32)
																v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																*(*int64)(unsafe.Add(mBase, uint32(v230+v231))) = int64(0)
																v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																*(*int64)(unsafe.Add(mBase, uint32(v235+v230))) = int64(9218868437227405312)
																v240 = v220 + int32(1)
																v241 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																if v240 < v241 {
																	v220 = v240
																	continue
																} else {
																	break
																}
																break
															}
															v246 = v241
														} else {
															v246 = v215
														}
														v255 = F_palloc0(m, v246<<(uint(int32(2))%32))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v255
															v258 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															v259 = F_palloc(m, v258)
															mBase = m.M
															v260 = m.ExcPending
															if v260 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v259
																v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																if v262 == int32(0) {
																} else {
																	base.MemoryFill(m, v259, int32(1), v262)
																}
																v278 = v16 + int32(132)
																v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																mBase = m.M
																v282 = m.ExcPending
																if v282 != 0 {
																	return int32(0)
																} else {
																	v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																	v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
																	v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
																	v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
																	v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
																	*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
																	*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
																	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
																	v297 = v16 + int32(160)
																	v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																	mBase = m.M
																	v301 = m.ExcPending
																	if v301 != 0 {
																		return int32(0)
																	} else {
																		v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																		v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
																		*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
																		v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
																		*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
																		v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
																		*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
																		v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
																		*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
																		*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
																		*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
																		v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																		v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
																		return v11
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
						}
					}
				}
			} else {
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v25
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				F_initSpGistState(m, v16, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
					v36 = F_AllocSetContextCreateInternal(m, v31, int32(_a_F_spgbeginscan_1), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v36
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
						v45 = F_AllocSetContextCreateInternal(m, v40, int32(_a_F_spgbeginscan_4), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v45
							v49 = v16 + int32(20)
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(4))%32))+88))
							if v50 != v56 {
								v58 = F_CreateTupleDescCopy(m, v51)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
									v64 = v58 + v61<<(uint(int32(4))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v64)+96)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v60
									v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v64)+92)) = uint16(v68)
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+6)))
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+102)) = uint8(v70)
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+7)))
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+103)) = uint8(v72)
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)))
									v75 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v64)+116)) = v75
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+105)) = uint8(v75)
									*(*uint8)(unsafe.Add(mBase, uint32(v64)+104)) = uint8(v74)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
									if v80 < int32(2) {
									} else {
										v84 = v58 + int32(20)
										v85 = int32(1)
										v86 = v80 - v85
										v87 = int32(7)
										v88 = v86 & v87
										if base.Ui32(v87) <= base.Ui32(v80-int32(2)) {
											v98 = v85
											v104 = int32(0)
											for {
												v109 = v84 + v98<<(uint(int32(4))%32)
												v110 = int32(-1)
												*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+80)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+64)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110
												v126 = int32(8)
												v127 = v98 + v126
												v129 = v104 + v126
												if v129 != v86&int32(-8) {
													v98 = v127
													v104 = v129
													continue
												} else {
													break
												}
												break
											}
											if v88 == int32(0) {
											} else {
												v134 = v127
												v145 = v134
												v147 = int32(0)
												for {
													*(*int32)(unsafe.Add(mBase, uint32(v84+v145<<(uint(int32(4))%32)))) = int32(-1)
													v159 = int32(1)
													v162 = v147 + v159
													if v162 != v88 {
														v145 = v145 + v159
														v147 = v162
														continue
													} else {
														break
													}
													break
												}
											}
										} else {
											v134 = v85
											v145 = v134
											v147 = int32(0)
											for {
												*(*int32)(unsafe.Add(mBase, uint32(v84+v145<<(uint(int32(4))%32)))) = int32(-1)
												v159 = int32(1)
												v162 = v147 + v159
												if v162 != v88 {
													v145 = v145 + v159
													v147 = v162
													continue
												} else {
													break
												}
												break
											}
										}
									}
									F_populate_compact_attribute(m, v58, int32(0))
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return int32(0)
									} else {
										v182 = v58
										*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v182
										*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v182
										v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										if v189 <= int32(0) {
											v278 = v16 + int32(132)
											v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
											mBase = m.M
											v282 = m.ExcPending
											if v282 != 0 {
												return int32(0)
											} else {
												v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
												v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
												v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
												v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
												v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
												*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
												*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
												*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
												v297 = v16 + int32(160)
												v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
												mBase = m.M
												v301 = m.ExcPending
												if v301 != 0 {
													return int32(0)
												} else {
													v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
													v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
													v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
													v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
													v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
													*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
													*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
													*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
													v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
													v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
													*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
													*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
													return v11
												}
											}
										} else {
											v194 = F_palloc(m, v189<<(uint(int32(2))%32))
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v194
												v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v200 = F_palloc(m, v197<<(uint(int32(2))%32))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v200
													v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v206 = F_palloc(m, v203<<(uint(int32(3))%32))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v206
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v212 = F_palloc(m, v209<<(uint(int32(3))%32))
														mBase = m.M
														v213 = m.ExcPending
														if v213 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v212
															v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															if int32(0) < v215 {
																v220 = int32(0)
																for {
																	v230 = v220 << (uint(int32(3)) % 32)
																	v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																	*(*int64)(unsafe.Add(mBase, uint32(v230+v231))) = int64(0)
																	v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																	*(*int64)(unsafe.Add(mBase, uint32(v235+v230))) = int64(9218868437227405312)
																	v240 = v220 + int32(1)
																	v241 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	if v240 < v241 {
																		v220 = v240
																		continue
																	} else {
																		break
																	}
																	break
																}
																v246 = v241
															} else {
																v246 = v215
															}
															v255 = F_palloc0(m, v246<<(uint(int32(2))%32))
															mBase = m.M
															v256 = m.ExcPending
															if v256 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v255
																v258 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																v259 = F_palloc(m, v258)
																mBase = m.M
																v260 = m.ExcPending
																if v260 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v259
																	v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	if v262 == int32(0) {
																	} else {
																		base.MemoryFill(m, v259, int32(1), v262)
																	}
																	v278 = v16 + int32(132)
																	v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																	mBase = m.M
																	v282 = m.ExcPending
																	if v282 != 0 {
																		return int32(0)
																	} else {
																		v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																		v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
																		*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
																		v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
																		*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
																		v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
																		*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
																		v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
																		*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
																		*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
																		*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
																		v297 = v16 + int32(160)
																		v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																		mBase = m.M
																		v301 = m.ExcPending
																		if v301 != 0 {
																			return int32(0)
																		} else {
																			v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																			v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
																			v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
																			*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
																			v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
																			*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
																			v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
																			*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
																			*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
																			*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
																			v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																			v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
																			*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
																			return v11
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
							} else {
								v182 = v51
								*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v182
								*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v182
								v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
								if v189 <= int32(0) {
									v278 = v16 + int32(132)
									v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
									mBase = m.M
									v282 = m.ExcPending
									if v282 != 0 {
										return int32(0)
									} else {
										v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
										v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
										v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
										v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
										v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
										*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
										*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
										*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
										v297 = v16 + int32(160)
										v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
										mBase = m.M
										v301 = m.ExcPending
										if v301 != 0 {
											return int32(0)
										} else {
											v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
											v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
											v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
											v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
											v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
											*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
											*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
											*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
											v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
											v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
											*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
											return v11
										}
									}
								} else {
									v194 = F_palloc(m, v189<<(uint(int32(2))%32))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v194
										v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										v200 = F_palloc(m, v197<<(uint(int32(2))%32))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v200
											v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v206 = F_palloc(m, v203<<(uint(int32(3))%32))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v206
												v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v212 = F_palloc(m, v209<<(uint(int32(3))%32))
												mBase = m.M
												v213 = m.ExcPending
												if v213 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v212
													v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													if int32(0) < v215 {
														v220 = int32(0)
														for {
															v230 = v220 << (uint(int32(3)) % 32)
															v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
															*(*int64)(unsafe.Add(mBase, uint32(v230+v231))) = int64(0)
															v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
															*(*int64)(unsafe.Add(mBase, uint32(v235+v230))) = int64(9218868437227405312)
															v240 = v220 + int32(1)
															v241 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															if v240 < v241 {
																v220 = v240
																continue
															} else {
																break
															}
															break
														}
														v246 = v241
													} else {
														v246 = v215
													}
													v255 = F_palloc0(m, v246<<(uint(int32(2))%32))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v255
														v258 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v259 = F_palloc(m, v258)
														mBase = m.M
														v260 = m.ExcPending
														if v260 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v259
															v262 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															if v262 == int32(0) {
															} else {
																base.MemoryFill(m, v259, int32(1), v262)
															}
															v278 = v16 + int32(132)
															v281 = F_index_getprocinfo(m, l0, int32(1), int32(4))
															mBase = m.M
															v282 = m.ExcPending
															if v282 != 0 {
																return int32(0)
															} else {
																v284 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																v285 = *(*int64)(unsafe.Add(mBase, uint32(v281)+16))
																*(*int64)(unsafe.Add(mBase, uint32(v278)+16)) = v285
																v287 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v287
																v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
																*(*int64)(unsafe.Add(mBase, uint32(v278)+8)) = v289
																v291 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
																*(*int64)(unsafe.Add(mBase, uint32(v278))) = v291
																*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v284
																*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = int32(0)
																v297 = v16 + int32(160)
																v300 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																mBase = m.M
																v301 = m.ExcPending
																if v301 != 0 {
																	return int32(0)
																} else {
																	v303 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																	v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v304
																	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v306
																	v308 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v308
																	v310 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
																	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v310
																	*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v303
																	*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = int32(0)
																	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v316
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
																	return v11
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
					}
				}
			}
		}
	}
}
func F_spgoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(256), int32(8), int32(_a_F_spgoptions_0), int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_spgvacuumscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 float64
	_ = v265
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v784 int32
	_ = v784
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	v19 = m.G0
	v21 = v19 - int32(880)
	m.G0 = v21
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	F_initSpGistState(m, l0+int32(16), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)) = uint8(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = v39
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	if v47 == v39 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v53 = base.B2i32(v50 == int32(0))
	goto L6
L5:
	;
	v53 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v61 = int32(0)
	v66 = F_read_stream_begin_relation(m, int32(13), v60, v26, v61, int32(120), v21+int32(40), v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v68 = l0
	v71 = v21
	v77 = v26
	v81 = l0 + int32(100)
	v83 = v66
	v85 = v53
	goto L8
L8:
	;
	if v85 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	if base.Ui32(v101) < base.Ui32(v100) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v89 = F_RelationGetNumberOfBlocksInFork(m, v77, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LockRelationForExtension(m, v77, int32(7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v100 = v89
	goto L10
L15:
	;
	v95 = F_RelationGetNumberOfBlocksInFork(m, v77, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_UnlockRelationForExtension(m, v77, int32(7))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v100 = v95
	goto L10
L18:
	;
	F_read_stream_reset(m, v119)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L189
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+44)) = v100
	v104 = v68
	v107 = v71
	v113 = v77
	v117 = v81
	v119 = v83
	v121 = v85
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_read_stream_end(m, v83)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L183
	}
L22:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L180
	}
L24:
	;
	v126 = F_read_stream_next_buffer(m, v119, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v126 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v126 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBuffer(m, v126, int32(2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+(v126^int32(-1))<<(uint(int32(6))%32))+16))
	v150 = v141
	goto L27
L29:
	;
	goto L30
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[2]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+v126<<(uint(int32(6))%32)+int32(-64))+16))
	v150 = v149
	goto L27
L31:
	;
	v154 = int32(0)
	v155 = base.B2i32(v154 <= v126)
	if v155 == v154 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v126)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+14)))
	if v420 != 0 {
		goto L86
	} else {
		goto L87
	}
L34:
	;
	if base.Ui32(v150-int32(1)) < base.Ui32(int32(2)) {
		goto L32
	} else {
		goto L84
	}
L35:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+14)))
	if v174 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+(v126^int32(-1))<<(uint(int32(2))%32))))
	v173 = v165
	goto L35
L37:
	;
	goto L38
L38:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v173 = v167 + v126<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)))
	if base.Ui32(v177) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+16)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v180))))
	if v182&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v185 = int32(1)
	if base.Ui32(v150-v185) <= base.Ui32(v185) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	F_vacuumRedirectAndPlaceholder(m, v131, v377, v126)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L83
	}
L44:
	;
	if v155 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_vacuumLeafPage(m, v104, v131, v126, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L81
	}
L47:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+12)))
	v208 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)) = uint16(v208)
	if base.Ui32(v207) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L51
	}
L48:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v126^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L47
L49:
	;
	goto L50
L50:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v206 = v200 + v126<<(uint(int32(13))%32) + int32(-8192)
	goto L47
L51:
	;
	v217 = int32(base.Ui32(v207+int32(_a_F_spgvacuumscan_0))>>(uint(int32(2))%32)) & int32(_a_F_spgvacuumscan_1)
	if v217 == int32(0) {
		goto L34
	} else {
		goto L52
	}
L52:
	;
	v228 = int32(1)
	v230 = int32(0)
	goto L53
L53:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(20)+v228&int32(_a_F_spgvacuumscan_1)<<(uint(int32(2))%32))))
	v250 = v206 + v247&int32(_a_F_spgvacuumscan_2)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v251&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v309 = v302 & int32(_a_F_spgvacuumscan_1)
	if v309 == int32(0) {
		goto L34
	} else {
		goto L66
	}
L55:
	;
	v304 = v228 + int32(1)
	if base.Ui32(v304&int32(_a_F_spgvacuumscan_1)) <= base.Ui32(v217) {
		v228 = v304
		v230 = v302
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v262)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v262)+8)) = base.F64_add(v298, float64(1))
	v302 = v230
	goto L55
L57:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v260 = m.T0[v259].(func(*base.Module, int32, int32) int32)(m, v250+int32(6), v258)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v260 == int32(0) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v262)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v262)+16)) = base.F64_add(v265, float64(1))
	v273 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v107+int32(48)+v230&int32(_a_F_spgvacuumscan_1)<<(uint(v273)%32)))) = uint16(v228)
	v278 = v230 + v273
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)) = uint16(v278)
	v302 = v278
	goto L55
L62:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+32)) = v284 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v107+int32(32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(445), int32(_a_F_spgvacuumscan_5))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	goto L54
L66:
	;
	v312 = int32(_a_F_spgvacuumscan_6)
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v314 + int32(1)
	F_PageIndexMultiDelete(m, v206, v107+int32(48), v309)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_MarkBufferDirty(m, v126)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+118)))
	if v325 != int32(112) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v363 = int32(_a_F_spgvacuumscan_6)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v365 - int32(1)
	goto L34
L70:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[6]))
	if v329 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	if v332 != 0 {
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v131)+40))
	if v333 != 0 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v104)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+872)) = v336
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+876)) = uint8(v338)
	F_XLogRegisterData(m, v107+int32(868), int32(12))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)))
	F_XLogRegisterData(m, v107+int32(48), v347<<(uint(int32(1))%32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_XLogRegisterBuffer(m, int32(0), v126, int32(8))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v358 = F_XLogInsert(m, int32(16), int32(112))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v206))) = base.I64_rotr(v358, int64(32))
	goto L69
L81:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	F_vacuumRedirectAndPlaceholder(m, v131, v373, v126)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L33
L83:
	;
	goto L34
L84:
	;
	goto L33
L85:
	;
	F_SpGistSetLastUsedPage(m, v131, v126)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v421) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_RecordFreeIndexPage(m, v131, v150)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+28)) = v427 + int32(1)
	goto L32
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+108)) = v150
	goto L32
L92:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v454 == int32(0) {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v459 = v104
	v462 = v107
	v464 = v458
	v465 = v454
	v468 = v113
	v472 = v117
	v474 = v119
	v476 = v121
	goto L95
L94:
	;
	goto L23
L95:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)))
	if v477 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	if v905 != 0 {
		goto L173
	} else {
		goto L174
	}
L97:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	v886 = v459
	v889 = v462
	v891 = v464
	v895 = v468
	v899 = v472
	v901 = v474
	v903 = v476
	goto L99
L99:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v904 != 0 {
		v459 = v886
		v462 = v889
		v464 = v891
		v465 = v904
		v468 = v895
		v472 = v899
		v474 = v901
		v476 = v903
		goto L95
	} else {
		goto L172
	}
L100:
	;
	v483 = int32(0)
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+2)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465))))
	v488 = v484 | v485<<(uint(int32(16))%32)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+24))
	v492 = F_ReadBufferExtended(m, v464, v483, v488, v483, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_LockBuffer(m, v492, int32(2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v492 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	F_UnlockReleaseBuffer(m, v492)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L171
	}
L104:
	;
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514)+14)))
	if v515 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L105:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v500+(v492^int32(-1))<<(uint(int32(2))%32))))
	v514 = v506
	goto L104
L106:
	;
	goto L107
L107:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v514 = v508 + v492<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514)+16)))
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+v518))))
	if v520&int32(2) != 0 {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	if v520&int32(4) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v540 = v465
	goto L113
L111:
	;
	goto L112
L112:
	;
	v821 = int32(1)
	if base.Ui32(v488-v821) <= base.Ui32(v821) {
		goto L94
	} else {
		goto L160
	}
L113:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+6)))
	if v547 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L157
	}
L115:
	;
	goto L114
L116:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if v804 != 0 {
		v540 = v804
		goto L113
	} else {
		goto L156
	}
L117:
	;
	v548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+2)))
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540))))
	if v548|v549<<(uint(int32(16))%32) != v488 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v514+int32(20)+v554<<(uint(int32(2))%32))))
	v561 = v514 + v558&int32(_a_F_spgvacuumscan_2)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	switch v562 & int32(3) {
	case 0:
		goto L121
	case 1:
		goto L120
	default:
		goto L115
	}
L119:
	;
	v784 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+6)) = uint8(v784)
	goto L116
L120:
	;
	v696 = v561 + int32(6)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v697 != 0 {
		goto L142
	} else {
		goto L143
	}
L121:
	;
	if v562&int32(_a_F_spgvacuumscan_7) == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v577 = v561 + int32(base.Ui32(v562)>>(uint(int32(16))%32)) + int32(8)
	v582 = int32(0)
	goto L123
L123:
	;
	v593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	if v593 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L119
L125:
	;
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+6)))
	v684 = int32(_a_F_spgvacuumscan_8)
	v688 = v582 + int32(1)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	if base.Ui32(v688) < base.Ui32(int32(base.Ui32(v689)>>(uint(int32(3))%32))&v684) {
		v577 = v577 + v683&v684
		v582 = v688
		goto L123
	} else {
		goto L141
	}
L126:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v596 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v598 = v596
	goto L130
L128:
	;
	v639 = v472
	goto L129
L129:
	;
	v654 = F_palloc(m, int32(12))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L140
	}
L130:
	;
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+2)))
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577))))
	v617 = int32(16)
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+2)))
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	if v615|v616<<(uint(v617)%32) == v620|v621<<(uint(v617)%32) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v639 = v598 + int32(8)
	goto L129
L132:
	;
	if v631 != 0 {
		goto L125
	} else {
		goto L138
	}
L133:
	;
	goto L132
L134:
	;
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+4)))
	if v627 == v628 {
		v631 = int32(1)
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v631 = int32(0)
	goto L133
L137:
	;
	goto L136
L138:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v598)+8))
	if v632 != 0 {
		v598 = v632
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v654)+4)) = uint16(v656)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	*(*int32)(unsafe.Add(mBase, uint32(v654))) = v658
	v660 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v654)+8)) = v660
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+6)) = uint8(v660)
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = v654
	goto L125
L141:
	;
	goto L124
L142:
	;
	v699 = v697
	goto L145
L143:
	;
	v740 = v472
	goto L144
L144:
	;
	v755 = F_palloc(m, int32(12))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L155
	}
L145:
	;
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+2)))
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696))))
	v718 = int32(16)
	v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699)+2)))
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699))))
	if v716|v717<<(uint(v718)%32) == v721|v722<<(uint(v718)%32) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v740 = v699 + int32(8)
	goto L144
L147:
	;
	if v732 != 0 {
		goto L119
	} else {
		goto L153
	}
L148:
	;
	goto L147
L149:
	;
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+4)))
	v729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699)+4)))
	if v728 == v729 {
		v732 = int32(1)
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v732 = int32(0)
	goto L148
L152:
	;
	goto L151
L153:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	if v733 != 0 {
		v699 = v733
		goto L145
	} else {
		goto L154
	}
L154:
	;
	goto L146
L155:
	;
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v755)+4)) = uint16(v757)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v759
	v761 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v755)+8)) = v761
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+6)) = uint8(v761)
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = v755
	goto L119
L156:
	;
	goto L103
L157:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v809 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v462)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(783), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_vacuumLeafPage(m, v459, v464, v492, int32(1))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	F_vacuumRedirectAndPlaceholder(m, v464, v829, v492)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_SpGistSetLastUsedPage(m, v464, v492)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v834 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)) = uint8(v834)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v836 == int32(0) {
		goto L103
	} else {
		goto L164
	}
L164:
	;
	v840 = v836
	goto L165
L165:
	;
	v857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v840)+2)))
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v840))))
	if v488 == v857|v858<<(uint(int32(16))%32) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L103
L167:
	;
	v863 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v840)+6)) = uint8(v863)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v840)+8))
	if v865 != 0 {
		v840 = v865
		goto L165
	} else {
		goto L170
	}
L170:
	;
	goto L166
L171:
	;
	v886 = v459
	v889 = v462
	v891 = v464
	v895 = v468
	v899 = v472
	v901 = v474
	v903 = v476
	goto L99
L172:
	;
	goto L96
L173:
	;
	v907 = v905
	goto L176
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = int32(0)
	v104 = v886
	v107 = v889
	v113 = v895
	v117 = v899
	v119 = v901
	v121 = v903
	goto L22
L176:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v907)+8))
	F_pfree(m, v907)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	goto L175
L178:
	;
	if v924 != 0 {
		v907 = v924
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v464)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+16)) = v951 + int32(4)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_10), v462+int32(16))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(722), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_SpGistUpdateMetaPage(m, v77)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+28))
	if v970 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_FreeSpaceMapVacuum(m, v77)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	v974 = v969
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = v100
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v976)+24)) = v977
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v979)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+32)) = v980
	m.G0 = v71 + int32(880)
	return
L188:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v974 = v973
	goto L187
L189:
	;
	v68 = v104
	v71 = v107
	v77 = v113
	v81 = v117
	v83 = v119
	v85 = v121
	goto L8
}
func F_spgvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
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
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v240 int32
	_ = v240
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v652 int32
	_ = v652
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
	var v663 int64
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v928 int32
	_ = v928
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(336)
	m.G0 = v25
	v28 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if int32(0) < v453 {
		goto L84
	} else {
		goto L85
	}
L2:
	;
	return int32(0)
L3:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	v38 = F_get_opfamily_name(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L81
	}
L7:
	;
	v42 = int32(0)
	v44 = F_SearchSysCacheList(m, int32(4), int32(1), v37, v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v48 = int32(0)
	v50 = F_SearchSysCacheList(m, int32(5), int32(1), v37, v48, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v52 = F_identify_opfamily_groups(m, v44, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	if v55 <= int32(0) {
		v431 = v54
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v60 = v54
	v65 = v2
	v66 = v2
	v76 = v2
	v78 = v2
	goto L12
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(48)+v65<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
	v88 = v86 + v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v89 == v90 {
		v119 = v60
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v431 = v392
	goto L1
L14:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+16)))
	switch v120 - int32(1) {
	case 0:
		goto L30
	case 1, 2, 3:
		goto L25
	case 4:
		goto L26
	case 5:
		goto L27
	case 6:
		goto L28
	default:
		goto L29
	}
L15:
	;
	v92 = int32(0)
	v95 = F_errstart(m, int32(17), v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v95 == int32(0) {
		v119 = v92
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v103 = F_format_procedure(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+296)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v25)+292)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+288)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_1), v25+int32(288))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(96), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v119 = v92
	goto L14
L22:
	;
	v415 = v65 + int32(1)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	if v415 < v416 {
		v60 = v392
		v65 = v415
		v66 = v398
		v76 = v408
		v78 = v410
		goto L12
	} else {
		goto L80
	}
L23:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L76
	}
L24:
	;
	v339 = int32(0)
	v342 = F_errstart(m, int32(17), v339)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L74
	}
L25:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+240)) = int64(9796820404457)
	v311 = int32(2)
	v315 = F_check_amproc_signature(m, v306, int32(2278), int32(1), v311, v311, v25+int32(240))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L72
	}
L26:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+256)) = int64(9796820404457)
	v298 = int32(2)
	v302 = F_check_amproc_signature(m, v293, int32(16), int32(1), v298, v298, v25+int32(256))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L70
	}
L27:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v76 != v278 {
		v323 = v66
		v333 = v76
		v335 = v78
		goto L24
	} else {
		goto L66
	}
L28:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v274 = F_check_amoptsproc_signature(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L2
	} else {
		goto L64
	}
L29:
	;
	v264 = int32(0)
	v267 = F_errstart(m, int32(17), v264)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L62
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+224)) = int64(9796820404457)
	v128 = int32(2)
	v132 = F_check_amproc_signature(m, v123, int32(2278), int32(1), v128, v128, v25+int32(224))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+312)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v25)+332)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v25)+320)) = v135
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v146 = F_OidFunctionCall2Coll(m, v140, int32(0), v25+int32(332), v25+int32(312))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v25)+320))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v35 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v153 = v35
	goto L35
L34:
	;
	v153 = v152
	goto L35
L35:
	;
	if base.B2i32(v149 == int32(0))|base.B2i32(v149 == v153) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v160 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	v185 = v119
	v186 = v153
	goto L38
L38:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v25)+332))
	if base.B2i32(v52 == int32(0))|base.B2i32(v186 != v189) != 0 {
		goto L48
	} else {
		goto L49
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+320))
	v185 = int32(0)
	v186 = v183
	goto L38
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v25)+320))
	v166 = F_format_type_be(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v168 = F_format_type_be(m, v153)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+212)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v25)+208)) = v166
	F_errmsg(m, int32(_a_F_spgvalidate_4), v25+int32(208))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(130), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	if v132 != 0 {
		v392 = v185
		v398 = v186
		v408 = v152
		v410 = v148
		goto L22
	} else {
		goto L61
	}
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v192 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v195 = int32(0)
	if v195 < v192 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v198 = v192
	goto L53
L52:
	;
	v198 = v195
	goto L53
L53:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v205 = int32(0)
	goto L54
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v200+v205<<(uint(int32(2))%32))))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v228 != v199 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L48
L56:
	;
	v240 = v205 + int32(1)
	if v240 != v198 {
		v205 = v240
		goto L54
	} else {
		goto L60
	}
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v230 != v231 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v227)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v227)+16)) = v233 | int64(64)
	if v132 == int32(0) {
		v323 = v186
		v333 = v152
		v335 = v148
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v392 = v185
	v398 = v186
	v408 = v152
	v410 = v148
	goto L22
L60:
	;
	goto L55
L61:
	;
	v323 = v186
	v333 = v152
	v335 = v148
	goto L24
L62:
	;
	if v267 == int32(0) {
		v392 = v264
		v398 = v66
		v408 = v76
		v410 = v78
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v348 = int32(184)
	v354 = v66
	v364 = v76
	v366 = v78
	v370 = int32(_a_F_spgvalidate_5)
	goto L23
L64:
	;
	if v274 == int32(0) {
		v323 = v66
		v333 = v76
		v335 = v78
		goto L24
	} else {
		goto L65
	}
L65:
	;
	v392 = v119
	v398 = v66
	v408 = v76
	v410 = v78
	goto L22
L66:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v78 != v280 {
		v323 = v66
		v333 = v76
		v335 = v78
		goto L24
	} else {
		goto L67
	}
L67:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+272)) = v76
	v284 = int32(1)
	v289 = F_check_amproc_signature(m, v282, v66, v284, v284, v284, v25+int32(272))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v289 == int32(0) {
		v323 = v66
		v333 = v76
		v335 = v78
		goto L24
	} else {
		goto L69
	}
L69:
	;
	v392 = v119
	v398 = v66
	v408 = v76
	v410 = v78
	goto L22
L70:
	;
	if v302 == int32(0) {
		v323 = v66
		v333 = v76
		v335 = v78
		goto L24
	} else {
		goto L71
	}
L71:
	;
	v392 = v119
	v398 = v66
	v408 = v76
	v410 = v78
	goto L22
L72:
	;
	if v315 != 0 {
		v392 = v119
		v398 = v66
		v408 = v76
		v410 = v78
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v323 = v66
	v333 = v76
	v335 = v78
	goto L24
L74:
	;
	if v342 == int32(0) {
		v392 = v339
		v398 = v323
		v408 = v333
		v410 = v335
		goto L22
	} else {
		goto L75
	}
L75:
	;
	v348 = int32(196)
	v354 = v323
	v364 = v333
	v366 = v335
	v370 = int32(_a_F_spgvalidate_6)
	goto L23
L76:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v375 = F_format_procedure(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v377 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+204)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v25)+200)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v25)+196)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+192)) = v38
	F_errmsg(m, v370, v25+int32(192))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), v348, int32(_a_F_spgvalidate_3))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v392 = int32(0)
	v398 = v354
	v408 = v364
	v410 = v366
	goto L22
L80:
	;
	goto L13
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
	F_errmsg_internal(m, int32(_a_F_spgvalidate_7), v25)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(63), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
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
	v459 = v431
	v464 = int32(0)
	goto L87
L85:
	;
	v602 = v431
	goto L86
L86:
	;
	if v52 != 0 {
		goto L119
	} else {
		goto L120
	}
L87:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(48)+v464<<(uint(int32(2))%32))))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+56))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+22)))
	v487 = v485 + v486
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+16)))
	if base.Ui32(int32(_a_F_spgvalidate_8)) < base.Ui32((v488+int32(-64))&int32(_a_F_spgvalidate_9)) {
		v524 = v459
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v602 = v597
	goto L86
L89:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+18)))
	if v527 == int32(115) {
		v563 = v524
		v564 = int32(16)
		goto L97
	} else {
		goto L98
	}
L90:
	;
	v495 = int32(0)
	v498 = F_errstart(m, int32(17), v495)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v498 == int32(0) {
		v524 = v495
		goto L89
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v506 = F_format_operator(m, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v508 = int32(*(*int16)(unsafe.Add(mBase, uint32(v487)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+188)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v25)+184)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v25)+180)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+176)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_10), v25+int32(176))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(216), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v524 = v495
	goto L89
L97:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v568 = F_check_amop_signature(m, v565, v564, v566, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L2
	} else {
		goto L109
	}
L98:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v531 = F_get_op_rettype(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v487)+28))
	v534 = F_opfamily_can_sort_type(m, v533, v531)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v534 != 0 {
		v563 = v524
		v564 = v531
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v536 = int32(0)
	v539 = F_errstart(m, int32(17), v536)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	if v539 == int32(0) {
		v563 = v536
		v564 = v531
		goto L97
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v547 = F_format_operator(m, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+168)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v25)+164)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+160)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_11), v25+int32(160))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(231), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	v563 = v536
	v564 = v531
	goto L97
L108:
	;
	v599 = v464 + int32(1)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v599 < v600 {
		v459 = v597
		v464 = v599
		goto L87
	} else {
		goto L117
	}
L109:
	;
	if v568 != 0 {
		v597 = v563
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v570 = int32(0)
	v573 = F_errstart(m, int32(17), v570)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	if v573 == int32(0) {
		v597 = v570
		goto L108
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v487)+20))
	v581 = F_format_operator(m, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+152)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v25)+148)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+144)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_12), v25+int32(144))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(247), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	v597 = v570
	goto L108
L117:
	;
	goto L88
L118:
	;
	F_ReleaseCatCacheList(m, v50)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L2
	} else {
		goto L198
	}
L119:
	;
	v624 = int32(0)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v625 <= v624 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v953 = int32(0)
	v956 = F_errstart(m, int32(17), v953)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L2
	} else {
		goto L193
	}
L122:
	;
	v906 = v602
	v928 = int32(1)
	goto L124
L123:
	;
	v630 = v602
	v633 = v624
	v635 = int32(0)
	goto L125
L124:
	;
	if v928 == int32(0) {
		v978 = v906
		goto L118
	} else {
		goto L192
	}
L125:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v652+v635<<(uint(int32(2))%32))))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	if v36 == v657 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v906 = v899
	v928 = base.B2i32(v662 == int32(0))
	goto L124
L127:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	if v659 == v36 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v662 = v633
	goto L129
L129:
	;
	v663 = *(*int64)(unsafe.Add(mBase, uint32(v656)+8))
	if v663 != int64(0) {
		v697 = v630
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v661 = v656
	goto L132
L131:
	;
	v661 = v633
	goto L132
L132:
	;
	v662 = v661
	goto L129
L133:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	if v699 != v700 {
		v899 = v697
		goto L142
	} else {
		goto L143
	}
L134:
	;
	v666 = int32(0)
	v669 = F_errstart(m, int32(17), v666)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	if v669 == int32(0) {
		v697 = v666
		goto L133
	} else {
		goto L136
	}
L136:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v677 = F_format_type_be(m, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v680 = F_format_type_be(m, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+140)) = v680
	*(*int32)(unsafe.Add(mBase, uint32(v25)+136)) = v677
	*(*int32)(unsafe.Add(mBase, uint32(v25)+132)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+128)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_13), v25+int32(128))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(275), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v697 = v666
	goto L133
L142:
	;
	v901 = v635 + int32(1)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v901 < v902 {
		v630 = v899
		v633 = v662
		v635 = v901
		goto L125
	} else {
		goto L191
	}
L143:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v702&int32(2) != 0 {
		v734 = v697
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v735&int32(4) != 0 {
		v767 = v734
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v705 = int32(0)
	v708 = F_errstart(m, int32(17), v705)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	if v708 == int32(0) {
		v734 = v705
		goto L144
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v716 = F_format_type_be(m, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+124)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v25)+120)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+116)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25+int32(112))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L2
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	v734 = v705
	goto L144
L152:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v768&int32(8) != 0 {
		v800 = v767
		goto L160
	} else {
		goto L161
	}
L153:
	;
	v738 = int32(0)
	v741 = F_errstart(m, int32(17), v738)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	if v741 == int32(0) {
		v767 = v738
		goto L152
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v749 = F_format_type_be(m, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+96)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25+int32(96))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	v767 = v738
	goto L152
L160:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v801&int32(16) != 0 {
		v833 = v800
		goto L168
	} else {
		goto L169
	}
L161:
	;
	v771 = int32(0)
	v774 = F_errstart(m, int32(17), v771)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	if v774 == int32(0) {
		v800 = v771
		goto L160
	} else {
		goto L163
	}
L163:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v782 = F_format_type_be(m, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v25)+88)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25+int32(80))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	v800 = v771
	goto L160
L168:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v834&int32(32) != 0 {
		v866 = v833
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v804 = int32(0)
	v807 = F_errstart(m, int32(17), v804)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	if v807 == int32(0) {
		v833 = v804
		goto L168
	} else {
		goto L171
	}
L171:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L2
	} else {
		goto L172
	}
L172:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v815 = F_format_type_be(m, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = v815
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25-int32(-64))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	v833 = v804
	goto L168
L176:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v867&int32(64) != 0 {
		v899 = v866
		goto L142
	} else {
		goto L184
	}
L177:
	;
	v837 = int32(0)
	v840 = F_errstart(m, int32(17), v837)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	if v840 == int32(0) {
		v866 = v837
		goto L176
	} else {
		goto L179
	}
L179:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v848 = F_format_type_be(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25+int32(48))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L2
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L2
	} else {
		goto L183
	}
L183:
	;
	v866 = v837
	goto L176
L184:
	;
	v870 = int32(0)
	v873 = F_errstart(m, int32(17), v870)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	if v873 == int32(0) {
		v899 = v870
		goto L142
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v881 = F_format_type_be(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v38
	F_errmsg(m, int32(_a_F_spgvalidate_14), v25+int32(32))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L2
	} else {
		goto L190
	}
L190:
	;
	v899 = v870
	goto L142
L191:
	;
	goto L126
L192:
	;
	goto L121
L193:
	;
	if v956 == int32(0) {
		v978 = v953
		goto L118
	} else {
		goto L194
	}
L194:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v34 + int32(8)
	F_errmsg(m, int32(_a_F_spgvalidate_15), v25+int32(16))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L2
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(308), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L2
	} else {
		goto L197
	}
L197:
	;
	v978 = v953
	goto L118
L198:
	;
	F_ReleaseCatCacheList(m, v44)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L2
	} else {
		goto L199
	}
L199:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L2
	} else {
		goto L200
	}
L200:
	;
	m.G0 = v25 + int32(336)
	return v978 & int32(1)
}
func F_split_part(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	v12 = m.G0
	v14 = v12 - int32(1072)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v14 + int32(1072)
	return v310
L5:
	;
	v296 = v285 - v288
	v298 = v296 + int32(4)
	v299 = F_palloc(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L96
	}
L6:
	;
	v285 = v243
	v288 = v223 + v236
	goto L5
L7:
	;
	if v249 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L8:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1052))
	if v211 <= int32(1) {
		v285 = v220
		v288 = v212
		goto L5
	} else {
		goto L80
	}
L9:
	;
	v201 = int32(1)
	v202 = v17 + v201
	v204 = v17 + int32(4)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v205&v201 != 0 {
		goto L77
	} else {
		goto L78
	}
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v25 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L73
	}
L13:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v55 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v31 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v34 = int32(16)
	goto L19
L18:
	;
	v34 = int32(0)
	goto L19
L19:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v41 = int32(4)
	goto L22
L21:
	;
	v41 = v34
	goto L22
L22:
	;
	v54 = v41
	goto L13
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	if v54 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v61 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v72 = int32(1)
	if v55&v72 != 0 {
		v84 = int32(base.Ui32(v55)>>(uint(v72)%32)) - v72
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v64 = int32(16)
	goto L30
L29:
	;
	v64 = int32(0)
	goto L30
L30:
	;
	if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v71 = int32(4)
	goto L33
L32:
	;
	v71 = v64
	goto L33
L33:
	;
	v84 = v71
	goto L24
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v88 = F_palloc(m, int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v84 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(16)
	v310 = v88
	goto L4
L39:
	;
	switch v24 + int32(1) {
	case 0, 2:
		v310 = v17
		goto L4
	default:
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_text_position_setup(m, v17, v22, v101, v14)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v97 = F_palloc(m, int32(4))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(16)
	v310 = v97
	goto L4
L44:
	;
	v104 = F_text_position_next(m, v14)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v104 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	switch v24 + int32(1) {
	case 0, 2:
		v310 = v17
		goto L4
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if int32(0) <= v24 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	v111 = F_palloc(m, int32(4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(16)
	v310 = v111
	goto L4
L51:
	;
	v118 = int32(2)
	goto L52
L52:
	;
	v131 = F_text_position_next(m, v14)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	if v24 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v131 != 0 {
		v118 = v118 + int32(1)
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v135 = int32(1)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v137&v135 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v161 = v24 + v118 + int32(1)
	if v161 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v140 = v135
	goto L61
L60:
	;
	v140 = int32(4)
	goto L61
L61:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1052))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1056))
	v145 = v143 + v144
	v146 = v17 + v140 + v54 - v145
	v148 = v146 + int32(4)
	v149 = F_palloc(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v148 << (uint(int32(2)) % 32)
	if v146 == int32(0) {
		v310 = v149
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.MemoryCopy(m, v149+int32(4), v145, v146)
	v310 = v149
	goto L4
L64:
	;
	v165 = F_cstring_to_text(m, int32(_a_F_split_part_0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1052)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1068)) = v167
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1064)) = v171
	v173 = F_text_position_next(m, v14)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v310 = v165
	goto L4
L68:
	;
	v175 = int32(1)
	v176 = v17 + v175
	v178 = v17 + int32(4)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v179&v175 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v182 = v176
	goto L71
L70:
	;
	v182 = v178
	goto L71
L71:
	;
	if v173 == int32(0) {
		v249 = v161
		v254 = v176
		v255 = v178
		v258 = v182
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v211 = v161
	v212 = v182
	v216 = v176
	v217 = v178
	goto L8
L73:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_split_part_1), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_split_part_2), int32(_a_F_split_part_3), int32(_a_F_split_part_4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v208 = v202
	goto L79
L78:
	;
	v208 = v204
	goto L79
L79:
	;
	v211 = v24
	v212 = v208
	v216 = v202
	v217 = v204
	goto L8
L80:
	;
	v223 = v220
	v225 = v211
	goto L81
L81:
	;
	v235 = v225 - int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1056))
	v237 = F_text_position_next(m, v14)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v249 = v235
	v254 = v216
	v255 = v217
	v258 = v223 + v236
	goto L7
L83:
	;
	goto L82
L84:
	;
	if v237 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1052))
	if base.B2i32(v225 < int32(3)) == int32(0) {
		v223 = v243
		v225 = v235
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L6
L87:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v261&int32(1) != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v280 = F_palloc(m, int32(4))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L95
	}
L90:
	;
	v264 = v254
	goto L92
L91:
	;
	v264 = v255
	goto L92
L92:
	;
	v266 = v264 - v258 + v54
	v268 = v266 + int32(4)
	v269 = F_palloc(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v268 << (uint(int32(2)) % 32)
	if v266 == int32(0) {
		v310 = v269
		goto L4
	} else {
		goto L94
	}
L94:
	;
	base.MemoryCopy(m, v269+int32(4), v258, v266)
	v310 = v269
	goto L4
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(16)
	v310 = v280
	goto L4
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v298 << (uint(int32(2)) % 32)
	if v296 == int32(0) {
		v310 = v299
		goto L4
	} else {
		goto L97
	}
L97:
	;
	base.MemoryCopy(m, v299+int32(4), v288, v296)
	v310 = v299
	goto L4
}
func F_sqlfunction_receive(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5 != 0 {
		v6 = F_ExecFilterJunk(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			F_tuplestore_puttupleslot(m, v10, v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
		if v16&int32(2) != 0 {
			v19 = F_ExecFilterJunk(m, v4, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
				m.T0[v22].(func(*base.Module, int32))(m, v19)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			}
		} else {
			return int32(1)
		}
	}
}
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v11 = m.G0
	v12 = int32(144)
	v13 = v11 - v12
	m.G0 = v13
	base.MemoryFill(m, v13, int32(0), v12)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_sscanf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = l0
	v24 = F_vfscanf(m, v13, l1, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v13 + int32(144)
		m.G0 = v8 + int32(16)
		return v24
	}
}
func F_storeOperators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v54 int64
	_ = v54
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v20 = F_table_open(m, int32(2602), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L61
	}
L4:
	;
	F_relation_close(m, v20, int32(3))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L60
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(0)
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+8)))
	v49 = F_SearchSysCacheExists(m, int32(4), l2, v46, v47, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v52
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v54
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v52)
	v68 = F_GetNewOidWithIndex(m, v20, int32(2756), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v49 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v74
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+8)))
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = int32(111)
	goto L17
L16:
	;
	v79 = int32(115)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v76
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v92 = F_heap_form_tuple(m, v87, v16-int32(-64), v16+int32(48))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_CatalogTupleInsert(m, v20, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(2602)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(2617)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v105
	v110 = v16 + int32(36)
	v112 = v16 + int32(24)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v115 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v116 = int32(110)
	goto L23
L22:
	;
	v116 = int32(97)
	goto L23
L23:
	;
	F_recordDependencyOn(m, v110, v112, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+25)))
	if v121 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v122 = int32(2753)
	goto L27
L26:
	;
	v122 = int32(2616)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v124
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v130 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v131 = int32(105)
	goto L30
L29:
	;
	v131 = int32(97)
	goto L30
L30:
	;
	F_recordDependencyOn(m, v110, v112, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v135 = F_typeDepNeeded(m, v134, v44)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v135 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(1247)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v139
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v145 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v150 == v151 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v146 = int32(110)
	goto L38
L37:
	;
	v146 = int32(97)
	goto L38
L38:
	;
	F_recordDependencyOn(m, v110, v112, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	if v174 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v153 = F_typeDepNeeded(m, v150, v44)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v153 == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(1247)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v159
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v169 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v170 = int32(110)
	goto L46
L45:
	;
	v170 = int32(97)
	goto L46
L46:
	;
	F_recordDependencyOn(m, v16+int32(36), v16+int32(24), v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(2753)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v186 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_storeOperators[0]))
	if v191 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v187 = int32(110)
	goto L53
L52:
	;
	v187 = int32(97)
	goto L53
L53:
	;
	F_recordDependencyOn(m, v16+int32(36), v16+int32(24), v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v193 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2602), v68, v193, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v198 = v38 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v198 < v199 {
		v38 = v198
		goto L7
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L8
L60:
	;
	m.G0 = v16 + int32(112)
	return
L61:
	;
	F_errcode(m, int32(_a_F_storeOperators_0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v229 = F_format_type_be(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v232 = F_format_type_be(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v234 = F_NameListToString(m, l0)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v227
	F_errmsg(m, int32(_a_F_storeOperators_1), v16)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_storeOperators_2), int32(1489), int32(_a_F_storeOperators_3))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_store_flush_position(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[0]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
	if v7 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v10 == int32(3) {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v15
			v18 = F_palloc(m, int32(24))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = l1
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3]))
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4]))
					v30 = v25
				} else {
					v27 = int32(_a_F_store_flush_position_0)
					*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3])) = v27
					v30 = v27
				}
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v30
				v32 = int32(_a_F_store_flush_position_0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v18
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4])) = v18
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[5]))
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v39
				return
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v15
		v18 = F_palloc(m, int32(24))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = l1
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3]))
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4]))
				v30 = v25
			} else {
				v27 = int32(_a_F_store_flush_position_0)
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3])) = v27
				v30 = v27
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v30
			v32 = int32(_a_F_store_flush_position_0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v18
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4])) = v18
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[5]))
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v39
			return
		}
	}
}
func F_strchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v7 = l1 & int32(255)
	if v7 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v99 == l1&int32(255) {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v97 = v87
	goto L1
L3:
	;
	v77 = v72
	goto L20
L4:
	;
	v72 = v64
	goto L3
L5:
	;
	if l0&int32(3) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v62 = F_strlen(m, l0)
	mBase = m.M
	v97 = v62 + l0
	goto L1
L8:
	;
	v10 = l0
	goto L11
L9:
	;
	v24 = l0
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v33 = int32(-2139062144)
	if (int32(16843008)-v30|v30)&v33 != v33 {
		v64 = v24
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v7 == v15) != 0 {
		v87 = v10
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v24 = v21
	goto L10
L13:
	;
	v21 = v10 + int32(1)
	if v21&int32(3) != 0 {
		v10 = v21
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v39 = v24
	v41 = v30
	goto L16
L16:
	;
	v45 = v41 ^ v7*int32(16843009)
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 != v48 {
		v64 = v39
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v72 = v54
	goto L3
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v54 = v39 + int32(4)
	v58 = int32(-2139062144)
	if (v52|(int32(16843008)-v52))&v58 == v58 {
		v39 = v54
		v41 = v52
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == int32(0) {
		v87 = v77
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v87 = v77
	goto L2
L22:
	;
	if v79 != l1&int32(255) {
		v77 = v77 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v103 = v97
	goto L26
L25:
	;
	v103 = int32(0)
	goto L26
L26:
	;
	return v103
}
func F_strdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v4 = F_strlen(m, l0)
	mBase = m.M
	v6 = v4 + int32(1)
	v7 = F_emscripten_builtin_malloc(m, v6)
	mBase = m.M
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v6) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v7
L5:
	;
	if v6 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v18 = v7 + v6
	if (v7^l0)&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	base.MemoryCopy(m, v7, l0, v6)
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v150) < base.Ui32(v18) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v7&int32(3) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v18) < base.Ui32(int32(4)) {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v54 = v18 & int32(-4)
	if base.Ui32(v18) < base.Ui32(int32(64)) {
		v104 = v48
		v105 = v49
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v48 = l0
	v49 = v7
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v6 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = l0
	v49 = v7
	goto L15
L20:
	;
	goto L21
L21:
	;
	v31 = l0
	v32 = v7
	goto L22
L22:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v36)
	v38 = int32(1)
	v39 = v31 + v38
	v41 = v32 + v38
	if v41&int32(3) == int32(0) {
		v48 = v39
		v49 = v41
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v48 = v39
	v49 = v41
	goto L15
L24:
	;
	if base.Ui32(v41) < base.Ui32(v18) {
		v31 = v39
		v32 = v41
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if base.Ui32(v54) <= base.Ui32(v105) {
		v149 = v104
		v150 = v105
		goto L11
	} else {
		goto L32
	}
L27:
	;
	v58 = v54 + int32(-64)
	if base.Ui32(v58) < base.Ui32(v49) {
		v104 = v48
		v105 = v49
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v61 = v48
	v62 = v49
	goto L29
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+28)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+32)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+36)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+48)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+56)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+60)) = v96
	v98 = int32(-64)
	v99 = v61 - v98
	v101 = v62 - v98
	if base.Ui32(v101) <= base.Ui32(v58) {
		v61 = v99
		v62 = v101
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v104 = v99
	v105 = v101
	goto L26
L31:
	;
	goto L30
L32:
	;
	v111 = v104
	v112 = v105
	goto L33
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v116
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v112 + v118
	if base.Ui32(v121) < base.Ui32(v54) {
		v111 = v119
		v112 = v121
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v149 = v119
	v150 = v121
	goto L11
L35:
	;
	goto L34
L36:
	;
	v149 = l0
	v150 = v7
	goto L11
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v6) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v149 = l0
	v150 = v7
	goto L11
L40:
	;
	goto L41
L41:
	;
	v130 = l0
	v131 = v7
	goto L42
L42:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+2)) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+3)) = uint8(v141)
	v143 = int32(4)
	v144 = v130 + v143
	v146 = v131 + v143
	if base.Ui32(v146) <= base.Ui32(v18-int32(4)) {
		v130 = v144
		v131 = v146
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v149 = v144
	v150 = v146
	goto L11
L44:
	;
	goto L43
L45:
	;
	v156 = v149
	v157 = v150
	goto L48
L46:
	;
	goto L47
L47:
	;
	goto L4
L48:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v161)
	v163 = int32(1)
	v166 = v157 + v163
	if v166 != v18 {
		v156 = v156 + v163
		v157 = v166
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
}
func F_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v542 int32
	_ = v542
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v728 int32
	_ = v728
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v773 int32
	_ = v773
	v3 = int32(0)
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v15 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	goto L3
L3:
	;
	v19 = F___strchrnul(m, l0, v15)
	mBase = m.M
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == v15&int32(255) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v773
L5:
	;
	if v25 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L9
	}
L6:
	;
	v25 = v19
	goto L8
L7:
	;
	v25 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v28 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return v25
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v32 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v39 = int32(0)
	v40 = base.B2i32(v38 != v39)
	if v38 == v39 {
		v82 = v25
		v85 = v40
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	if v99 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L27
	}
L17:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v44 = int32(8)
	v46 = v43<<(uint(v44)%32) | v38
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v51 = v47 | v48<<(uint(v44)%32)
	if v46 == v51 {
		v82 = v25
		v85 = v40
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = v25 + int32(1)
	v60 = v46
	goto L20
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v70 = int32(0)
	v71 = base.B2i32(v69 != v70)
	if v69 == v70 {
		v82 = v56
		v85 = v71
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v82 = v56
	v85 = v71
	goto L17
L22:
	;
	v80 = v60<<(uint(int32(8))%32)&int32(_a_F_strstr_0) | v69
	if v80 != v51 {
		v56 = v56 + int32(1)
		v60 = v80
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v97 = v82
	goto L26
L25:
	;
	v97 = int32(0)
	goto L26
L26:
	;
	return v97
L27:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	if v102 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = v25 + int32(2)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v108 = int32(0)
	if v107 == v108 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
	if v180 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L42
	}
L31:
	;
	if v163 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v162 = v106
	v163 = base.B2i32(v107 != v108)
	goto L31
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v113 = int32(16)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v116 = int32(24)
	v119 = int32(8)
	v121 = v112<<(uint(v113)%32) | v115<<(uint(v116)%32) | v107<<(uint(v119)%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v132 = v122<<(uint(v113)%32) | v125<<(uint(v116)%32) | v129<<(uint(v119)%32)
	if v121 == v132 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v137 = v106
	v138 = v121
	goto L35
L35:
	;
	v149 = v137 + int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v151 = int32(0)
	v152 = base.B2i32(v150 != v151)
	if v150 == v151 {
		v162 = v149
		v163 = v152
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v162 = v149
	v163 = v152
	goto L31
L37:
	;
	v157 = (v150 | v138) << (uint(int32(8)) % 32)
	if v157 != v132 {
		v137 = v149
		v138 = v157
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v178 = v162 - int32(2)
	goto L41
L40:
	;
	v178 = int32(0)
	goto L41
L41:
	;
	return v178
L42:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v183 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v187 = v25 + int32(3)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
	v189 = int32(0)
	if v188 == v189 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	v263 = int32(0)
	v264 = m.G0
	v266 = v264 - int32(1056)
	m.G0 = v266
	v268 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1048)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1040)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1032)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1024)) = v268
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v276 == v263 {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	if v249 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v246 = v187
	v249 = base.B2i32(v188 != v189)
	goto L46
L48:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v197 = int32(24)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v201 = int32(8)
	v204 = v193<<(uint(int32(16))%32) | v196<<(uint(v197)%32) | v200<<(uint(v201)%32) | v188
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v206 = int32(16711935)
	v214 = base.I32_rotr(v205&v206, v201) | base.I32_rotr(v205, v197)&v206
	if v204 == v214 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v219 = v187
	v220 = v204
	goto L50
L50:
	;
	v231 = v219 + int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v233 = int32(0)
	v234 = base.B2i32(v232 != v233)
	if v232 == v233 {
		v246 = v231
		v249 = v234
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v246 = v231
	v249 = v234
	goto L46
L52:
	;
	v239 = v220<<(uint(int32(8))%32) | v232
	if v239 != v214 {
		v219 = v231
		v220 = v239
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v261 = v246 - int32(3)
	goto L56
L55:
	;
	v261 = int32(0)
	goto L56
L56:
	;
	return v261
L57:
	;
	m.G0 = v266 + int32(1056)
	v773 = v756
	goto L4
L58:
	;
	v439 = int32(1)
	v443 = base.B2i32(base.Ui32(v430+v439) < base.Ui32(v429+v439))
	if base.Ui32(v430+v439) < base.Ui32(v429+v439) {
		goto L97
	} else {
		goto L98
	}
L59:
	;
	v342 = int32(1)
	v344 = v321
	v345 = v342
	v346 = v263
	v349 = v322
	v354 = v342
	goto L70
L60:
	;
	v756 = int32(0)
	goto L57
L61:
	;
	v425 = v325
	v429 = int32(-1)
	v430 = v330
	v431 = v331
	v432 = int32(1)
	goto L58
L62:
	;
	v325 = int32(1)
	v330 = int32(-1)
	v331 = v3
	goto L61
L63:
	;
	goto L64
L64:
	;
	v282 = v276
	v287 = v3
	goto L65
L65:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v287))))
	if v296 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v321 = int32(1)
	v322 = int32(-1)
	if base.Ui32(v321) < base.Ui32(v305) {
		goto L59
	} else {
		goto L69
	}
L67:
	;
	v304 = int32(1)
	v305 = v287 + v304
	*(*int32)(unsafe.Add(mBase, uint32(v266+v282&int32(255)<<(uint(int32(2))%32)))) = v305
	v313 = v266 + int32(1024) + int32(base.Ui32(v282)>>(uint(int32(3))%32))&int32(28)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v314 | v304<<(uint(v282)%32)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+l1))))
	if v320 != 0 {
		v282 = v320
		v287 = v305
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v325 = v321
	v330 = v322
	v331 = v305
	goto L61
L70:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+l1+v345))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+l1))))
	if v360 == v362 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v384 = int32(1)
	v387 = int32(0)
	v388 = v384
	v389 = v384
	v391 = int32(-1)
	v394 = v384
	goto L83
L72:
	;
	v380 = v379 + v376
	if base.Ui32(v380) < base.Ui32(v305) {
		v344 = v380
		v345 = v379
		v346 = v376
		v349 = v377
		v354 = v378
		goto L70
	} else {
		goto L82
	}
L73:
	;
	if v345 == v354 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v362) < base.Ui32(v360) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v376 = v346 + v354
	v377 = v349
	v378 = v354
	v379 = int32(1)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v376 = v346
	v377 = v349
	v378 = v354
	v379 = v345 + int32(1)
	goto L72
L79:
	;
	v376 = v344
	v377 = v349
	v378 = v344 - v349
	v379 = int32(1)
	goto L72
L80:
	;
	goto L81
L81:
	;
	v372 = int32(1)
	v376 = v346 + v372
	v377 = v346
	v378 = v372
	v379 = v372
	goto L72
L82:
	;
	goto L71
L83:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391+l1+v388))))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+l1))))
	if v403 == v405 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v425 = v378
	v429 = v420
	v430 = v377
	v431 = v305
	v432 = v421
	goto L58
L85:
	;
	v423 = v422 + v419
	if base.Ui32(v423) < base.Ui32(v305) {
		v387 = v419
		v388 = v422
		v389 = v423
		v391 = v420
		v394 = v421
		goto L83
	} else {
		goto L95
	}
L86:
	;
	if v388 == v394 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(v403) < base.Ui32(v405) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v419 = v387 + v394
	v420 = v391
	v421 = v394
	v422 = int32(1)
	goto L85
L90:
	;
	goto L91
L91:
	;
	v419 = v387
	v420 = v391
	v421 = v394
	v422 = v388 + int32(1)
	goto L85
L92:
	;
	v419 = v389
	v420 = v391
	v421 = v389 - v391
	v422 = int32(1)
	goto L85
L93:
	;
	goto L94
L94:
	;
	v415 = int32(1)
	v419 = v387 + v415
	v420 = v387
	v421 = v415
	v422 = v415
	goto L85
L95:
	;
	goto L84
L96:
	;
	v524 = v431 | int32(63)
	v527 = int32(0)
	v529 = v25
	v530 = v25
	goto L127
L97:
	;
	v444 = v432
	goto L99
L98:
	;
	v444 = v425
	goto L99
L99:
	;
	v445 = l1 + v444
	if base.Ui32(v430+v439) < base.Ui32(v429+v439) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v446 = v429
	goto L102
L101:
	;
	v446 = v430
	goto L102
L102:
	;
	v448 = v446 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v448) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	if v510 != 0 {
		goto L121
	} else {
		goto L122
	}
L104:
	;
	v510 = int32(0)
	goto L103
L105:
	;
	v484 = v479
	v485 = v480
	v486 = v481
	goto L115
L106:
	;
	if (l1|v445)&int32(3) != 0 {
		v479 = l1
		v480 = v445
		v481 = v448
		goto L105
	} else {
		goto L109
	}
L107:
	;
	v472 = l1
	v473 = v445
	v474 = v448
	goto L108
L108:
	;
	if v474 == int32(0) {
		goto L104
	} else {
		goto L114
	}
L109:
	;
	v456 = l1
	v457 = v445
	v458 = v448
	goto L110
L110:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v461 != v462 {
		v479 = v456
		v480 = v457
		v481 = v458
		goto L105
	} else {
		goto L112
	}
L111:
	;
	v472 = v467
	v473 = v465
	v474 = v469
	goto L108
L112:
	;
	v464 = int32(4)
	v465 = v457 + v464
	v467 = v456 + v464
	v469 = v458 - v464
	if base.Ui32(int32(3)) < base.Ui32(v469) {
		v456 = v467
		v457 = v465
		v458 = v469
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v479 = v472
	v480 = v473
	v481 = v474
	goto L105
L115:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v489 == v490 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v510 = v489 - v490
	goto L103
L117:
	;
	v492 = int32(1)
	v497 = v486 - v492
	if v497 != 0 {
		v484 = v484 + v492
		v485 = v485 + v492
		v486 = v497
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	goto L104
L121:
	;
	v513 = v431 + (v446 ^ int32(-1))
	if base.Ui32(v513) < base.Ui32(v446) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v521 = v444
	v522 = v431 - v444
	goto L96
L124:
	;
	v515 = v446
	goto L126
L125:
	;
	v515 = v513
	goto L126
L126:
	;
	v521 = v515 + int32(1)
	v522 = int32(0)
	goto L96
L127:
	;
	if base.Ui32(v431) <= base.Ui32(v530-v529) {
		v657 = v530
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v658 = int32(0)
	v661 = v529 + v431
	v662 = int32(1)
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661-v662))))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(1024)+int32(base.Ui32(v664)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v670)>>(uint(v664)%32))&v662 == v658 {
		v527 = v658
		v529 = v661
		v530 = v657
		goto L127
	} else {
		goto L161
	}
L130:
	;
	v542 = int32(0)
	if base.B2i32(v530&int32(3) == v542)|base.B2i32(v524 == v542) != 0 {
		v574 = v530
		v576 = v524
		v577 = base.B2i32(v524 != v542)
		goto L134
	} else {
		goto L135
	}
L131:
	;
	if v648 != 0 {
		goto L156
	} else {
		goto L157
	}
L132:
	;
	v648 = int32(0)
	goto L131
L133:
	;
	v626 = v619
	v628 = v621
	goto L150
L134:
	;
	if v577 == int32(0) {
		goto L132
	} else {
		goto L141
	}
L135:
	;
	v557 = v530
	v559 = v524
	goto L136
L136:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v562 == int32(0) {
		v619 = v557
		v621 = v559
		goto L133
	} else {
		goto L138
	}
L137:
	;
	v574 = v569
	v576 = v565
	v577 = v567
	goto L134
L138:
	;
	v564 = int32(1)
	v565 = v559 - v564
	v566 = int32(0)
	v567 = base.B2i32(v565 != v566)
	v569 = v557 + v564
	if v569&int32(3) == v566 {
		v574 = v569
		v576 = v565
		v577 = v567
		goto L134
	} else {
		goto L139
	}
L139:
	;
	if v565 != 0 {
		v557 = v569
		v559 = v565
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v582 = int32(0)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if base.B2i32(v582 == v583)|base.B2i32(base.Ui32(v576) < base.Ui32(int32(4))) == v582 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v592 = v574
	v594 = v576
	goto L145
L143:
	;
	v612 = v574
	v614 = v576
	goto L144
L144:
	;
	if v614 == int32(0) {
		goto L132
	} else {
		goto L149
	}
L145:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	v599 = v598 ^ int32(0)
	v602 = int32(-2139062144)
	if (int32(16843008)-v599|v599)&v602 != v602 {
		v619 = v592
		v621 = v594
		goto L133
	} else {
		goto L147
	}
L146:
	;
	v612 = v607
	v614 = v609
	goto L144
L147:
	;
	v606 = int32(4)
	v607 = v592 + v606
	v609 = v594 - v606
	if base.Ui32(int32(3)) < base.Ui32(v609) {
		v592 = v607
		v594 = v609
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v619 = v612
	v621 = v614
	goto L133
L150:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if int32(0) == v631 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L132
L152:
	;
	v648 = v626
	goto L131
L153:
	;
	goto L154
L154:
	;
	v633 = int32(1)
	v636 = v628 - v633
	if v636 != 0 {
		v626 = v626 + v633
		v628 = v636
		goto L150
	} else {
		goto L155
	}
L155:
	;
	goto L151
L156:
	;
	v650 = v648
	goto L158
L157:
	;
	v650 = v530 + v524
	goto L158
L158:
	;
	if v648 == int32(0) {
		v657 = v650
		goto L129
	} else {
		goto L159
	}
L159:
	;
	if base.Ui32(v648-v529) < base.Ui32(v431) {
		v756 = v542
		goto L57
	} else {
		goto L160
	}
L160:
	;
	v657 = v650
	goto L129
L161:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v266+v664<<(uint(int32(2))%32))))
	if v679 != v431 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v681 = v431 - v679
	if base.Ui32(v527) < base.Ui32(v681) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	if base.Ui32(v527) < base.Ui32(v448) {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v683 = v681
	goto L167
L166:
	;
	v683 = v527
	goto L167
L167:
	;
	v527 = int32(0)
	v529 = v529 + v683
	v530 = v657
	goto L127
L168:
	;
	v527 = int32(0)
	v529 = v529 + (v693 - v446)
	v530 = v657
	goto L127
L169:
	;
	v687 = v448
	goto L171
L170:
	;
	v687 = v527
	goto L171
L171:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v687))))
	if v689 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v692 = v689
	v693 = v687
	goto L175
L173:
	;
	goto L174
L174:
	;
	v728 = v448
	goto L179
L175:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v693))))
	if v705 != v692&int32(255) {
		goto L168
	} else {
		goto L177
	}
L176:
	;
	goto L174
L177:
	;
	v710 = v693 + int32(1)
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v710))))
	if v712 != 0 {
		v692 = v712
		v693 = v710
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	if base.Ui32(v728) <= base.Ui32(v527) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v527 = v522
	v529 = v529 + v521
	v530 = v657
	goto L127
L181:
	;
	v756 = v529
	goto L57
L182:
	;
	goto L183
L183:
	;
	v743 = v728 - int32(1)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v743))))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v743))))
	if v745 == v747 {
		v728 = v743
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
}
func F_strtof(m *base.Module, l0 int32, l1 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float32(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v23 = m.G0
		v25 = v23 - int32(32)
		m.G0 = v25
		v28 = v15 & int64(281474976710655)
		v32 = int64(base.Ui64(v15)>>(uint(int64(48))%64)) & int64(32767)
		v33 = base.I32_wrap_i64(v32)
		if base.Ui32(v33-int32(_a_F_strtof_0)) <= base.Ui32(int32(253)) {
			v40 = base.I32_wrap_i64(int64(base.Ui64(v28) >> (uint(int64(25)) % 64)))
			v44 = v15 & int64(33554431)
			v45 = int64(16777216)
			if v44 == v45 {
				v49 = base.B2i32(v14 == int64(0))
			} else {
				v49 = base.B2i32(base.Ui64(v44) < base.Ui64(v45))
			}
			if v49 == int32(0) {
				v62 = v40 + int32(1)
			} else {
				if v14|(v44^int64(16777216)) != int64(0) {
					v62 = v40
				} else {
					v62 = v40&int32(1) + v40
				}
			}
			v65 = base.B2i32(base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62))
			if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62) {
				v66 = int32(0)
			} else {
				v66 = v62
			}
			if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62) {
				v69 = int32(-16255)
			} else {
				v69 = int32(-16256)
			}
			v150 = v66
			v151 = v69 + v33
		} else {
			if base.B2i32(v14|v28 == int64(0))|base.B2i32(v32 != int64(32767)) == int32(0) {
				v150 = base.I32_wrap_i64(int64(base.Ui64(v28)>>(uint(int64(25))%64))) | int32(_a_F_strtof_2)
				v151 = int32(255)
			} else {
				if base.Ui32(int32(_a_F_strtof_3)) < base.Ui32(v33) {
					v150 = int32(0)
					v151 = int32(255)
				} else {
					v91 = base.B2i32(v32 == int64(0))
					if v32 == int64(0) {
						v92 = int32(_a_F_strtof_4)
					} else {
						v92 = int32(_a_F_strtof_0)
					}
					v93 = v92 - v33
					if int32(112) < v93 {
						v96 = int32(0)
						v150 = v96
						v151 = v96
					} else {
						if v32 == int64(0) {
							v100 = v28
						} else {
							v100 = v28 | int64(281474976710656)
						}
						if v33 != v92 {
							F___ashlti3(m, v25+int32(16), v14, v100, int32(128)-v93)
							mBase = m.M
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
							v113 = base.B2i32(v108|v109 != int64(0))
						} else {
							v113 = int32(0)
						}
						F___lshrti3(m, v25, v14, v100, v93)
						mBase = m.M
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
						v118 = base.I32_wrap_i64(int64(base.Ui64(v115) >> (uint(int64(25)) % 64)))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
						v121 = v119 | base.I64_extend_i32_u(v113)
						v125 = v115 & int64(33554431)
						v126 = int64(16777216)
						if v125 == v126 {
							v130 = base.B2i32(v121 == int64(0))
						} else {
							v130 = base.B2i32(base.Ui64(v125) < base.Ui64(v126))
						}
						if v130 == int32(0) {
							v143 = v118 + int32(1)
						} else {
							if v121|(v125^int64(16777216)) != int64(0) {
								v143 = v118
							} else {
								v143 = v118&int32(1) + v118
							}
						}
						v147 = base.B2i32(base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v143))
						if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v143) {
							v148 = v143 ^ int32(_a_F_strtof_5)
						} else {
							v148 = v143
						}
						v150 = v148
						v151 = v147
					}
				}
			}
		}
		m.G0 = v25 + int32(32)
		m.G0 = v7 + int32(16)
		return base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(32))%64)))&int32(-2147483648) | v151<<(uint(int32(23))%32) | v150)
	}
}
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(2147483648))
	return base.I32_wrap_i64(v5)
}
func F_subcolorcvec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v230 int32
	_ = v230
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v300 int32
	_ = v300
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v816 int32
	_ = v816
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1075 int32
	_ = v1075
	var v1095 int32
	_ = v1095
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1143 int32
	_ = v1143
	var v1159 int32
	_ = v1159
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v25 = int32(_a_F_subcolorcvec_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v27 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v1159 + int32(16)
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v35 = v27
	v36 = v30
	goto L5
L3:
	;
	goto L4
L4:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v81 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_subcoloronechr(m, l0, v50, l2, l3, v22+int32(14))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v55 != 0 {
		v1159 = v22
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v58 = int32(1)
	if v58 < v35 {
		v35 = v35 - v58
		v36 = v36 + int32(4)
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v100 = v84
	v101 = v81
	goto L14
L12:
	;
	v662 = v22
	goto L13
L13:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v667 < int32(0) {
		v1159 = v662
		goto L1
	} else {
		goto L133
	}
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if base.Ui32(int32(2047)) < base.Ui32(v105) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v662 = v22
	goto L13
L16:
	;
	if base.Ui32(v300) < base.Ui32(v104) {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	v300 = v105
	goto L16
L18:
	;
	goto L19
L19:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
	if base.Ui32(v104) < base.Ui32(v105) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v279)
	v300 = v280
	goto L16
L21:
	;
	v279 = v108
	v280 = v105
	goto L20
L22:
	;
	goto L23
L23:
	;
	v110 = int32(2047)
	if base.Ui32(v110) <= base.Ui32(v104) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v113 = v110
	goto L26
L25:
	;
	v113 = v104
	goto L26
L26:
	;
	v119 = v105
	v120 = v108
	goto L27
L27:
	;
	v133 = F_subcolor(m, v24, v119)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	v279 = v257
	v280 = v272
	goto L20
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v135 != 0 {
		v1159 = v22
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v136 = int32(_a_F_subcolorcvec_0)
	v137 = v133 & v136
	if v137 != v120&v136 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_subcolorcvec[0]))
	if v143 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v257 = v120
	goto L33
L33:
	;
	v272 = v119 + int32(1)
	if base.Ui32(v119) < base.Ui32(v113) {
		v119 = v272
		v120 = v257
		goto L27
	} else {
		goto L61
	}
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v146 <= v147 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L36
L38:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v250 != 0 {
		v1159 = v22
		goto L1
	} else {
		goto L60
	}
L39:
	;
	F_createarc(m, v141, int32(112), v133, l2, l3)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L59
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v149 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v179 == int32(0) {
		goto L39
	} else {
		goto L51
	}
L43:
	;
	v156 = v149
	goto L44
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	if v171 != l3 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L39
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v178 != 0 {
		v156 = v178
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	if v173 != v137 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v175 == int32(112) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L45
L51:
	;
	v186 = v179
	goto L52
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v201 != l2 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L39
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	if v208 != 0 {
		v186 = v208
		goto L52
	} else {
		goto L58
	}
L55:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+4)))
	if v203 != v137 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if v205 == int32(112) {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L53
L59:
	;
	goto L38
L60:
	;
	v257 = v133
	goto L33
L61:
	;
	goto L28
L62:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v641 != 0 {
		v1159 = v22
		goto L1
	} else {
		goto L131
	}
L63:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	v320 = F_palloc_extended(m, v314*int32(24)+int32(12), int32(2))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v300 != v104 {
		goto L62
	} else {
		goto L129
	}
L66:
	;
	if v320 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+24)) = int32(101)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	if v328 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v313)+88))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v333 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v330 = v328
	goto L72
L71:
	;
	v330 = int32(12)
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+12)) = v330
	goto L62
L73:
	;
	if v385 <= v379 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v378 = v332
	v379 = int32(0)
	v385 = v333
	goto L73
L75:
	;
	goto L76
L76:
	;
	v337 = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if base.Ui32(v300) <= base.Ui32(v338) {
		v378 = v332
		v379 = v337
		v385 = v333
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v345 = v337
	v346 = v332
	goto L78
L78:
	;
	v359 = int32(12)
	v361 = v320 + v345*v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v362
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v346)))
	*(*int64)(unsafe.Add(mBase, uint32(v361))) = v364
	v367 = v346 + v359
	v369 = v345 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v370 <= v369 {
		v378 = v367
		v379 = v369
		v385 = v370
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v378 = v367
	v379 = v369
	v385 = v370
	goto L73
L80:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v346)+16))
	if base.Ui32(v372) < base.Ui32(v300) {
		v345 = v369
		v346 = v367
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if base.Ui32(v526) <= base.Ui32(v104) {
		goto L114
	} else {
		goto L115
	}
L83:
	;
	v523 = v378
	v524 = v379
	v525 = v379
	v526 = v300
	v530 = v385
	goto L82
L84:
	;
	goto L85
L85:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if base.Ui32(v104) < base.Ui32(v394) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v523 = v378
	v524 = v379
	v525 = v379
	v526 = v300
	v530 = v385
	goto L82
L87:
	;
	goto L88
L88:
	;
	v402 = v378
	v403 = v379
	v404 = v379
	v405 = v300
	v411 = v394
	goto L89
L89:
	;
	if base.Ui32(v405) < base.Ui32(v411) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v523 = v510
	v524 = v503
	v525 = v514
	v526 = v512
	v530 = v515
	goto L82
L91:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	F_subcoloronerow(m, l0, v500, l2, l3, v22+int32(14))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L7
	} else {
		goto L111
	}
L92:
	;
	v469 = v320 + v464*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v465
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v104) < base.Ui32(v471) {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v455 = v320 + v403*int32(12)
	v456 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v405 - v456
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = v411
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+8)) = v460
	v464 = v403 + v456
	v465 = v405
	goto L92
L94:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v104) < base.Ui32(v441) {
		v464 = v438
		v465 = v439
		goto L92
	} else {
		goto L101
	}
L95:
	;
	v420 = v320 + v403*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v405
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v422 - int32(1)
	v427 = F_newhicolorrow(m, v313, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if base.Ui32(v411) < base.Ui32(v405) {
		goto L93
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = v427
	F_subcoloronerow(m, l0, v427, l2, l3, v22+int32(14))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v438 = v403 + int32(1)
	v439 = v436
	goto L94
L100:
	;
	v438 = v403
	v439 = v405
	goto L94
L101:
	;
	v445 = v320 + v438*int32(12)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v446
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v402)))
	*(*int64)(unsafe.Add(mBase, uint32(v445))) = v448
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v500 = v450
	v503 = v438 + int32(1)
	goto L91
L102:
	;
	v473 = v104
	goto L104
L103:
	;
	v473 = v471
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v473
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v476 = F_newhicolorrow(m, v313, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+8)) = v476
	v480 = v464 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v481) <= base.Ui32(v104) {
		v500 = v476
		v503 = v480
		goto L91
	} else {
		goto L106
	}
L106:
	;
	v485 = v320 + v480*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v104 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+4)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if base.Ui32(v490) < base.Ui32(v465) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v492 = F_newhicolorrow(m, v313, v489)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L110
	}
L108:
	;
	v494 = v489
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+8)) = v494
	v500 = v476
	v503 = v464 + int32(2)
	goto L91
L110:
	;
	v494 = v492
	goto L109
L111:
	;
	v510 = v402 + int32(12)
	v511 = int32(1)
	v512 = v504 + v511
	v514 = v404 + v511
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v515 <= v514 {
		v523 = v510
		v524 = v503
		v525 = v514
		v526 = v512
		v530 = v515
		goto L82
	} else {
		goto L112
	}
L112:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	if base.Ui32(v517) <= base.Ui32(v104) {
		v402 = v510
		v403 = v503
		v404 = v514
		v405 = v512
		v411 = v517
		goto L89
	} else {
		goto L113
	}
L113:
	;
	goto L90
L114:
	;
	v541 = v320 + v524*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v541)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v526
	v545 = F_newhicolorrow(m, v313, int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L117
	}
L115:
	;
	v555 = v524
	v557 = v530
	goto L116
L116:
	;
	if v525 < v557 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541)+8)) = v545
	F_subcoloronerow(m, l0, v545, l2, l3, v22+int32(14))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	v555 = v524 + int32(1)
	v557 = v552
	goto L116
L119:
	;
	v563 = v523
	v564 = v555
	v565 = v525
	goto L122
L120:
	;
	v598 = v555
	goto L121
L121:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v313)+88))
	if v612 != 0 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v578 = int32(12)
	v580 = v320 + v564*v578
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+8)) = v581
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v563)))
	*(*int64)(unsafe.Add(mBase, uint32(v580))) = v583
	v587 = int32(1)
	v588 = v564 + v587
	v590 = v565 + v587
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v590 < v591 {
		v563 = v563 + v578
		v564 = v588
		v565 = v590
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v598 = v588
	goto L121
L124:
	;
	goto L123
L125:
	;
	F_pfree(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L7
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+84)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v313)+88)) = v320
	goto L62
L128:
	;
	goto L127
L129:
	;
	F_subcoloronechr(m, l0, v104, l2, l3, v22+int32(14))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	goto L62
L131:
	;
	v644 = int32(1)
	if v644 < v101 {
		v100 = v100 + int32(8)
		v101 = v101 - v644
		goto L14
	} else {
		goto L132
	}
L132:
	;
	goto L15
L133:
	;
	v671 = v24 + int32(28)
	v674 = v671 + v667<<(uint(int32(2))%32)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if v675 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v678
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v24)+96))
	v684 = base.I32_div_s(int32(2147483647), v681<<(uint(int32(1))%32))
	if v684 <= v678 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v868 = v675
	goto L136
L136:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	if v877 <= int32(0) {
		v1159 = v662
		goto L1
	} else {
		goto L164
	}
L137:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v852 != 0 {
		v1159 = v662
		goto L1
	} else {
		goto L163
	}
L138:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v686)+24)) = int32(101)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+12))
	if v690 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v698 = F_repalloc_extended(m, v694, v678*v681<<(uint(int32(2))%32))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L7
	} else {
		goto L144
	}
L141:
	;
	v692 = v690
	goto L143
L142:
	;
	v692 = int32(12)
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+12)) = v692
	goto L137
L144:
	;
	if v698 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+24)) = int32(101)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+12))
	if v706 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v698
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v714 = v712 - int32(1)
	if int32(0) <= v714 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v708 = v706
	goto L150
L149:
	;
	v708 = int32(12)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v705)+12)) = v708
	goto L137
L151:
	;
	v721 = v714
	v722 = v711
	goto L154
L152:
	;
	v816 = v711
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v816 << (uint(int32(1)) % 32)
	goto L137
L154:
	;
	if int32(0) < v722 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v816 = v793
	goto L153
L156:
	;
	v738 = v721 * v722
	v741 = v698 + v738<<(uint(int32(2))%32)
	v742 = int32(1)
	v753 = int32(0)
	goto L159
L157:
	;
	v793 = v722
	goto L158
L158:
	;
	if int32(0) < v721 {
		v721 = v721 - int32(1)
		v722 = v793
		goto L154
	} else {
		goto L162
	}
L159:
	;
	v768 = int32(1)
	v769 = v753 << (uint(v768) % 32)
	v772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v769+(v698+v738<<(uint(v742)%32))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v741+v722<<(uint(v742)%32)+v769))) = uint16(v772)
	*(*uint16)(unsafe.Add(mBase, uint32(v769+v741))) = uint16(v772)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v779 = v776 + v772*int32(24)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v779)+4)) = v780 + v768
	v785 = v753 + v768
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v785 < v786 {
		v753 = v785
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v793 = v786
	goto L158
L161:
	;
	goto L160
L162:
	;
	goto L155
L163:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v671+v853<<(uint(int32(2))%32))))
	v868 = v857
	goto L136
L164:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v887 = v880
	v889 = v877
	v890 = v881
	v898 = int32(0)
	goto L165
L165:
	;
	v902 = int32(0)
	if v902 < v887 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v1159 = v662
	goto L1
L167:
	;
	v910 = v902
	v912 = v890
	goto L170
L168:
	;
	v1127 = v887
	v1129 = v889
	v1130 = v890
	goto L169
L169:
	;
	v1143 = v898 + int32(1)
	if v1143 < v1129 {
		v887 = v1127
		v889 = v1129
		v890 = v1130
		v898 = v1143
		goto L165
	} else {
		goto L214
	}
L170:
	;
	if v910&v868 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v1127 = v1120
	v1129 = v1122
	v1130 = v1117
	goto L169
L172:
	;
	v1117 = v912 + int32(2)
	v1119 = v910 + int32(1)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v1119 < v1120 {
		v910 = v1119
		v912 = v1117
		goto L170
	} else {
		goto L213
	}
L173:
	;
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v912))))
	v929 = v927 * int32(24)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v931 = v929 + v930
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v931)+8)))
	if v932 != int32(_a_F_subcolorcvec_0) {
		v953 = v932
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+12))
	if v956 != 0 {
		v979 = int32(_a_F_subcolorcvec_0)
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v931)))
	if v935+v936 == int32(1) {
		v953 = v927
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v940 = F_newcolor(m, v24)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	if v940 == int32(-1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v953 = int32(_a_F_subcolorcvec_0)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v945+v929)+8)) = uint16(v940)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v948+v940*int32(24))+8)) = uint16(v940)
	v953 = v940
	goto L174
L181:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v980 != 0 {
		v1159 = v662
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v957 = int32(_a_F_subcolorcvec_0)
	if v927&v957 == v953&v957 {
		v979 = v927
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v963 = v962 + v929
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)+4))
	v965 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v963)+4)) = v964 - v965
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v972 = v968 + base.I32_extend16_s(v953)*int32(24)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v972)+4)) = v973 + v965
	*(*uint16)(unsafe.Add(mBase, uint32(v912))) = uint16(v953)
	v979 = v953
	goto L181
L184:
	;
	v982 = v979 & int32(_a_F_subcolorcvec_0)
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662)+14)))
	if v982 == v983 {
		goto L172
	} else {
		goto L185
	}
L185:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v987 = *(*int32)(unsafe.Add(mBase, _c_F_subcolorcvec[0]))
	if v987 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L7
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v990 <= v991 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L188
L190:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1095 != 0 {
		v1159 = v662
		goto L1
	} else {
		goto L212
	}
L191:
	;
	F_createarc(m, v985, int32(112), base.I32_extend16_s(v979), l2, l3)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L7
	} else {
		goto L211
	}
L192:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v993 == int32(0) {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v1023 == int32(0) {
		goto L191
	} else {
		goto L203
	}
L195:
	;
	v1000 = v993
	goto L196
L196:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+12))
	if v1015 != l3 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L191
L198:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+16))
	if v1022 != 0 {
		v1000 = v1022
		goto L196
	} else {
		goto L202
	}
L199:
	;
	v1017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1000)+4)))
	if v1017 != v982 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	if v1019 == int32(112) {
		goto L190
	} else {
		goto L201
	}
L201:
	;
	goto L198
L202:
	;
	goto L197
L203:
	;
	v1030 = v1023
	goto L204
L204:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+8))
	if v1045 != l2 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L191
L206:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+24))
	if v1052 != 0 {
		v1030 = v1052
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v1047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1030)+4)))
	if v1047 != v982 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1030)))
	if v1049 == int32(112) {
		goto L190
	} else {
		goto L209
	}
L209:
	;
	goto L206
L210:
	;
	goto L205
L211:
	;
	goto L190
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v662)+14)) = uint16(v979)
	goto L172
L213:
	;
	goto L171
L214:
	;
	goto L166
}
func F_subcoloronechr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	v6 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if base.Ui32(l1) <= base.Ui32(int32(2047)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v17 = F_subcolor(m, v14, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v120 = F_palloc_extended(m, v114*int32(12)+int32(24), int32(2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L35
	}
L5:
	;
	return
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	if v20 == v17&int32(_a_F_subcoloronechr_0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_subcoloronechr[0]))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v29 <= v30 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L11
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v17)
	return
L14:
	;
	F_createarc(m, v24, int32(112), v17, l2, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L34
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v32 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v58 == int32(0) {
		goto L14
	} else {
		goto L26
	}
L18:
	;
	v42 = v32
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v50 != l3 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L14
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v57 != 0 {
		v42 = v57
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	if v52 != v17&int32(_a_F_subcoloronechr_0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v54 == int32(112) {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L20
L26:
	;
	v68 = v58
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v76 != l2 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L14
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if v83 != 0 {
		v68 = v83
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	if v78 != v17&int32(_a_F_subcoloronechr_0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v80 == int32(112) {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L28
L34:
	;
	goto L13
L35:
	;
	if v120 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+24)) = int32(101)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v133 <= int32(0) {
		v171 = v132
		v172 = v6
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v130 = v128
	goto L41
L40:
	;
	v130 = int32(12)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v130
	return
L42:
	;
	F_subcoloronerow(m, l0, v263, l2, l3, l4)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L68
	}
L43:
	;
	if v164 == v149 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v181 = v120 + v172*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = l1
	v185 = F_newhicolorrow(m, v14, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L53
	}
L45:
	;
	v141 = v132
	v142 = v6
	goto L46
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if base.Ui32(v149) < base.Ui32(l1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v164) <= base.Ui32(l1) {
		goto L43
	} else {
		goto L52
	}
L48:
	;
	v151 = int32(12)
	v153 = v120 + v142*v151
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v154
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v153))) = v156
	v159 = v141 + v151
	v161 = v142 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v161 < v162 {
		v141 = v159
		v142 = v161
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
	v171 = v159
	v172 = v161
	goto L44
L52:
	;
	v171 = v141
	v172 = v142
	goto L44
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+8)) = v185
	v257 = v171
	v258 = v172
	v259 = v172 + int32(1)
	v263 = v185
	goto L42
L54:
	;
	v191 = int32(12)
	v193 = v120 + v142*v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v202 = v142 + int32(1)
	v257 = v141 + v191
	v258 = v202
	v259 = v202
	v263 = v198
	goto L42
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v164) < base.Ui32(l1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v206 = v120 + v142*int32(12)
	v207 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = l1 - v207
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v164
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+8)) = v211
	v215 = v142 + v207
	goto L59
L58:
	;
	v215 = v142
	goto L59
L59:
	;
	v218 = v120 + v215*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = l1
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v222 = F_newhicolorrow(m, v14, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v222
	v226 = v215 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if base.Ui32(l1) < base.Ui32(v227) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v231 = v120 + v226*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = l1 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v238) < base.Ui32(l1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v246 = v226
	goto L63
L63:
	;
	v257 = v141 + int32(12)
	v258 = v142 + int32(1)
	v259 = v246
	v263 = v222
	goto L42
L64:
	;
	v240 = F_newhicolorrow(m, v14, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	v242 = v237
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = v242
	v246 = v215 + int32(2)
	goto L63
L67:
	;
	v242 = v240
	goto L66
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v258 < v267 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v274 = v257
	v275 = v258
	v276 = v259
	goto L72
L70:
	;
	v304 = v259
	goto L71
L71:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v310 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v282 = int32(12)
	v284 = v120 + v276*v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+8)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(v284))) = v287
	v291 = int32(1)
	v292 = v276 + v291
	v294 = v275 + v291
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v294 < v295 {
		v274 = v274 + v282
		v275 = v294
		v276 = v292
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v304 = v292
	goto L71
L74:
	;
	goto L73
L75:
	;
	F_pfree(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v120
	goto L1
L78:
	;
	goto L77
}
func F_substitute_phv_relids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(319) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v11 != v12 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = v32
				if v34 == int32(67) {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
					v43 = F_query_tree_walker_impl(m, l0, int32(852), l1, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
						return v43
					}
				} else {
					v51 = F_expression_tree_walker_impl(m, l0, int32(852), l1)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						return v51
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v16 = F_bms_is_member(m, v14, v15)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v16 == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v34 = v32
						if v34 == int32(67) {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
							v43 = F_query_tree_walker_impl(m, l0, int32(852), l1, int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
								return v43
							}
						} else {
							v51 = F_expression_tree_walker_impl(m, l0, int32(852), l1)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								return v51
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v24 = F_bms_union(m, v22, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v28 = F_bms_del_member(m, v24, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v34 = v32
								if v34 == int32(67) {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
									v43 = F_query_tree_walker_impl(m, l0, int32(852), l1, int32(0))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
										return v43
									}
								} else {
									v51 = F_expression_tree_walker_impl(m, l0, int32(852), l1)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										return v51
									}
								}
							}
						}
					}
				}
			}
		} else {
			v34 = v8
			if v34 == int32(67) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
				v43 = F_query_tree_walker_impl(m, l0, int32(852), l1, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
					return v43
				}
			} else {
				v51 = F_expression_tree_walker_impl(m, l0, int32(852), l1)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			}
		}
	}
}
func F_swedish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
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
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v11 + int32(3)
	if v9 < v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v134 < v137 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 < v11 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v65 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v27 = v11
	goto L6
L5:
	;
	v27 = v25
	goto L6
L6:
	;
	v34 = v11
	goto L8
L7:
	;
	v65 = v45
	goto L3
L8:
	;
	if v34 == v27 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v34))))
	if int32(246) < v40 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = v34 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
	v34 = v57
	goto L8
L14:
	;
	v42 = v40 - int32(97)
	if v42 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v42)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v49)>>(uint(v42&int32(7))%32))&v45 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 < v76 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v120 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v79 = v76
	goto L22
L21:
	;
	v79 = v77
	goto L22
L22:
	;
	v85 = v76
	goto L24
L23:
	;
	v120 = int32(1)
	goto L19
L24:
	;
	if v85 == v79 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v120 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v85))))
	if int32(246) < v94 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v96 = v94 - int32(97)
	if v96 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v96)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v102)>>(uint(v96&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v111 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	v85 = v111
	goto L24
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = v123 + v120
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 < v124 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = v124
	goto L36
L35:
	;
	v129 = v127
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v129
	goto L1
L37:
	;
	return v395
L38:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v238 <= v235 {
		goto L64
	} else {
		goto L65
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	if v134 <= v137 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	goto L38
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v134-v144))))
	if base.B2i32(v146&int32(224) != int32(96))|base.B2i32(v144<<(uint(v146)%32)&int32(_a_F_swedish_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v160 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_1), int32(37))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v160 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	switch v160 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L52
L47:
	;
	v171 = F_slice_del(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v171 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v395 = v171
	goto L37
L50:
	;
	if v227 != 0 {
		goto L38
	} else {
		goto L61
	}
L51:
	;
	v227 = v223
	goto L50
L52:
	;
	if v183 <= v184 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v223 = int32(0)
	goto L51
L54:
	;
	v227 = int32(-1)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v196 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v183-v196))))
	if int32(121) < v201 {
		v223 = v196
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v203 = v201 - int32(98)
	if v203 < int32(0) {
		v223 = v196
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v203)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v209)>>(uint(v203&int32(7))%32))&int32(1) == int32(0) {
		v223 = v196
		goto L51
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v183 - int32(1)
	goto L60
L60:
	;
	goto L53
L61:
	;
	v228 = F_slice_del(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	if int32(0) <= v228 {
		goto L38
	} else {
		goto L63
	}
L63:
	;
	v395 = v228
	goto L37
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
	v243 = v235 - int32(1)
	if v243 <= v238 {
		v278 = v235
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v282 = v235
	v284 = v237
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v282 < v286 {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v240
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v282 = v278
	v284 = v281
	goto L66
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))))
	if base.B2i32(v247&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v247)%32)&int32(_a_F_swedish_ISO_8859_1_stem_2) == int32(0)) != 0 {
		v278 = v235
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v261 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_3), int32(7))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L43
	} else {
		goto L70
	}
L70:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v261 == int32(0) {
		v278 = v263
		goto L67
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v263
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v263 <= v267 {
		v278 = v263
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v270 = v263 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v270
	v273 = F_slice_del(m, l0)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L43
	} else {
		goto L73
	}
L73:
	;
	if v273 < int32(0) {
		v395 = v273
		goto L37
	} else {
		goto L74
	}
L74:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v278 = v277
	goto L67
L75:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v395 = int32(1)
	goto L37
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v282
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v286
	v292 = v282 - int32(1)
	if v292 <= v286 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v289
	goto L75
L78:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v292))))
	if base.B2i32(v296&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v296)%32)&int32(_a_F_swedish_ISO_8859_1_stem_4) == int32(0)) != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v310 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_5), int32(5))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L43
	} else {
		goto L80
	}
L80:
	;
	if v310 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v289
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v315
	switch v310 - int32(1) {
	case 0:
		goto L84
	case 1:
		goto L83
	case 2:
		goto L82
	default:
		goto L75
	}
L82:
	;
	v384 = F_slice_from_s(m, l0, int32(4), int32(_a_F_swedish_ISO_8859_1_stem_6))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L43
	} else {
		goto L101
	}
L83:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L89
L84:
	;
	v319 = F_slice_del(m, l0)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L43
	} else {
		goto L85
	}
L85:
	;
	if int32(0) <= v319 {
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v395 = v319
	goto L37
L87:
	;
	if v375 != 0 {
		goto L75
	} else {
		goto L98
	}
L88:
	;
	v375 = v371
	goto L87
L89:
	;
	if v331 <= v332 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v371 = int32(0)
	goto L88
L91:
	;
	v375 = int32(-1)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v344 = int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+v331-v344))))
	if int32(118) < v349 {
		v371 = v344
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v351 = v349 - int32(105)
	if v351 < int32(0) {
		v371 = v344
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v357)>>(uint(v351&int32(7))%32))&int32(1) == int32(0) {
		v371 = v344
		goto L88
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 - int32(1)
	goto L97
L97:
	;
	goto L90
L98:
	;
	v378 = F_slice_from_s(m, l0, int32(2), int32(_a_F_swedish_ISO_8859_1_stem_7))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L43
	} else {
		goto L99
	}
L99:
	;
	if int32(0) <= v378 {
		goto L75
	} else {
		goto L100
	}
L100:
	;
	v395 = v378
	goto L37
L101:
	;
	if int32(0) <= v384 {
		goto L75
	} else {
		goto L102
	}
L102:
	;
	v395 = v384
	goto L37
}
