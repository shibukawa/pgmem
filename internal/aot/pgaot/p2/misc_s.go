package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
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
	return v189
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
		v189 = v15
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
	v189 = v15
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
		v189 = v40
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
	v189 = v40
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
	if v10&int32(16) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	if v67 == int32(0) {
		v189 = v65
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
	v189 = v65
	goto L1
L27:
	;
	v111 = int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v112 != 0 {
		v189 = v111
		goto L1
	} else {
		goto L35
	}
L28:
	;
	if v60 == int32(-1) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = int32(0)
	v93 = F_errstart(m, l1, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	if v93 == int32(0) {
		v189 = v91
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_5), v6+int32(-32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(687), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v189 = v91
	goto L1
L35:
	;
	v114 = l0 + int32(96)
	if (l0^v114)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v189 = v111
	goto L1
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v168)
	if v168&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v167 = l0
	v168 = v120
	v169 = v114
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
	v124 = l0
	v126 = v114
	goto L45
L43:
	;
	v138 = l0
	v140 = v114
	goto L44
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v145 = int32(-2139062144)
	if (int32(16843008)-v142|v142)&v145 != v145 {
		v167 = v138
		v168 = v142
		v169 = v140
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v127)
	if v127 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v138 = v134
	v140 = v132
	goto L44
L47:
	;
	v131 = int32(1)
	v132 = v126 + v131
	v134 = v124 + v131
	if v134&int32(3) != 0 {
		v124 = v134
		v126 = v132
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v150 = v138
	v151 = v142
	v152 = v140
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v151
	v154 = int32(4)
	v155 = v152 + v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v158 = v150 + v154
	v162 = int32(-2139062144)
	if (v156|(int32(16843008)-v156))&v162 == v162 {
		v150 = v158
		v151 = v156
		v152 = v155
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v167 = v158
	v168 = v156
	v169 = v155
	goto L38
L52:
	;
	goto L51
L53:
	;
	v176 = v167
	v178 = v169
	goto L54
L54:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)) = uint8(v179)
	v181 = int32(1)
	if v179 != 0 {
		v176 = v176 + v181
		v178 = v178 + v181
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v7 = int32(-1)
	if l0&int32(3) == int32(0) {
		v31 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v112
L2:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v65) < base.Ui32(v64) {
		v112 = v7
		goto L1
	} else {
		goto L19
	}
L3:
	;
	v64 = v56 - l0
	goto L2
L4:
	;
	v35 = v31
	goto L13
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v64 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v20 = l0
	goto L9
L9:
	;
	v24 = v20 + int32(1)
	if v24&int32(3) == int32(0) {
		v31 = v24
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v56 = v24
	goto L3
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != 0 {
		v20 = v24
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v44 = int32(-2139062144)
	if (int32(16843008)-v41|v41)&v44 == v44 {
		v35 = v35 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v50 = v35
	goto L16
L15:
	;
	goto L14
L16:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		v50 = v50 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v56 = v50
	goto L3
L18:
	;
	goto L17
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v68 = m.T0[v67].(func(*base.Module, int32, int32) int32)(m, l0, v64)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v68 < int32(0) {
		v112 = v7
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v74 <= v68 {
		v112 = v7
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77+v68<<(uint(int32(1))%32)))))
	v83 = l0
	v85 = v76 + v81
	goto L24
L24:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v106 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(1)
	if base.Ui32((v89-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v103 = v89 | int32(32)
	goto L31
L30:
	;
	v103 = v89
	goto L31
L31:
	;
	if v90 == v103 {
		v83 = v83 + v91
		v85 = v85 + v91
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v112 = v7
	goto L1
L33:
	;
	v107 = int32(-1)
	goto L35
L34:
	;
	v107 = v68
	goto L35
L35:
	;
	v112 = v107
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
	var v43 int32
	_ = v43
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
	v42 = F_kill(m, v8, int32(23))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
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
L14:
	;
	return
L15:
	;
	goto L1
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
	var v40 int32
	_ = v40
	v3 = int32(0)
	if l0 == v3 {
		v40 = v3
		return v40
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
					v40 = v38
					return v40
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v12 != v13 {
					v40 = v3
					return v40
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v15 != v16 {
						v40 = v3
						return v40
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
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(368)
	m.G0 = v10
	v13 = v10 + int32(184)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = v10 + int32(200)
	v28 = F___memset(m, v10+int32(208), v2, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = int64(1)
	v41 = F___memcpy(m, v20, v23, int32(16))
	mBase = m.M
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v42
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v48
	v56 = F___syscall_ret(m, v2)
	mBase = m.M
	m.G0 = v20 + int32(16)
	F___gettimeofday(m, v10+int32(336))
	mBase = m.M
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+192))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+344))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
	if v64 < v66 {
		v68 = *(*int64)(unsafe.Add(mBase, uint32(v10)+336))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+336)) = v68 - int64(1)
		v74 = v64 + int32(_a_F_ShowUsage_0)
	} else {
		v74 = v64
	}
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+208))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
	if v63 < v78 {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v63 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v76 - int64(1)
	} else {
	}
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v10)+200))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
	if v75 < v88 {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = v75 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+200)) = v86 - int64(1)
	} else {
	}
	F_initStringInfo(m, v10+int32(352))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		return
	} else {
		F_appendStringInfoString(m, v10+int32(352), int32(_a_F_ShowUsage_1))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return
		} else {
			v105 = *(*int64)(unsafe.Add(mBase, uint32(v10)+336))
			v107 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[3]))
			v108 = v105 - v107
			*(*uint32)(unsafe.Add(mBase, uint32(v10)+176)) = uint32(v108)
			v111 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v74 - v111
			v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
			v116 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[4]))
			v117 = v114 - v116
			*(*uint32)(unsafe.Add(mBase, uint32(v10)+160)) = uint32(v117)
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v10)+192))
			v121 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v119 - v121
			v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+200))
			v126 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[5]))
			v127 = v124 - v126
			*(*uint32)(unsafe.Add(mBase, uint32(v10)+168)) = uint32(v127)
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+208))
			v131 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v129 - v131
			F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_2), v10+int32(160))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v75
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+152)) = uint32(v86)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v63
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+144)) = uint32(v76)
				F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_3), v10+int32(144))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+216))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v152
					F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_4), v10+int32(128))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return
					} else {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(v10)+244))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v161
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)+248))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v163
						v166 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[6]))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v161 - v166
						v170 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v163 - v170
						F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_5), v10+int32(112))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+240))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v180
							v183 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[8]))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v180 - v183
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v10)+236))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v186
							v188 = *(*int32)(unsafe.Add(mBase, uint32(v10)+232))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v188
							v191 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[9]))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v186 - v191
							v195 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[10]))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v188 - v195
							F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_6), v10+int32(80))
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
								return
							} else {
								v207 = *(*int32)(unsafe.Add(mBase, uint32(v10)+256))
								*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v207
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v10)+252))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v209
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v10)+260))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v211
								v214 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[11]))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v211 - v214
								v218 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[12]))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v207 - v218
								v222 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[13]))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v209 - v222
								F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_7), v10+int32(48))
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return
								} else {
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+264))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v232
									v234 = *(*int32)(unsafe.Add(mBase, uint32(v10)+268))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v234
									v237 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[14]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v232 - v237
									v241 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[15]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v234 - v241
									F_appendStringInfo(m, v10+int32(352), int32(_a_F_ShowUsage_8), v10+int32(32))
									mBase = m.M
									v250 = m.ExcPending
									if v250 != 0 {
										return
									} else {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+352))
										v252 = *(*int32)(unsafe.Add(mBase, uint32(v10)+356))
										v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v252-int32(1)))))
										if v256 == int32(10) {
											v260 = v252 - int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+356)) = v260
											v263 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v251+v260))) = uint8(v263)
										} else {
										}
										v268 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return
										} else {
											if v268 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
												F_errmsg_internal(m, int32(_a_F_ShowUsage_9), v10+int32(16))
												mBase = m.M
												v275 = m.ExcPending
												if v275 != 0 {
													return
												} else {
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v10)+352))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v276
													F_errdetail_internal(m, int32(_a_F_ShowUsage_9), v10)
													mBase = m.M
													v280 = m.ExcPending
													if v280 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ShowUsage_10), int32(_a_F_ShowUsage_11), int32(_a_F_ShowUsage_12))
														mBase = m.M
														v285 = m.ExcPending
														if v285 != 0 {
															return
														} else {
															v286 = *(*int32)(unsafe.Add(mBase, uint32(v10)+352))
															F_pfree(m, v286)
															mBase = m.M
															v288 = m.ExcPending
															if v288 != 0 {
																return
															} else {
																m.G0 = v10 + int32(368)
																return
															}
														}
													}
												}
											} else {
												v286 = *(*int32)(unsafe.Add(mBase, uint32(v10)+352))
												F_pfree(m, v286)
												mBase = m.M
												v288 = m.ExcPending
												if v288 != 0 {
													return
												} else {
													m.G0 = v10 + int32(368)
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v69 int32
	_ = v69
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
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
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
	v32 = v9
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
	v366 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v366)
	if v307&int32(3) == v366 {
		v391 = v307
		goto L117
	} else {
		goto L118
	}
L10:
	;
	return int32(0)
L11:
	;
	v311 = v305
	goto L103
L12:
	;
	if v34 == int32(0) {
		v66 = v32
		v69 = v32
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v74 = v32 + int32(1)
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
	if v32 == v69 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	if v34 == int32(44) {
		v66 = v32
		v69 = v32
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v42 = v32
	v45 = v32
	v46 = v34
	goto L18
L18:
	;
	v49 = v42 + int32(1)
	v50 = base.I32_extend8_s(v46)
	goto L20
L19:
	;
	v66 = v49
	v69 = v60
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
	v60 = v45
	goto L23
L22:
	;
	v60 = v49
	goto L23
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v61 == int32(0) {
		v66 = v49
		v69 = v60
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v61 != int32(44) {
		v42 = v49
		v45 = v60
		v46 = v61
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v305 = v66
	v307 = v32
	v308 = v69
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
	v88 = v82
	goto L32
L32:
	;
	v92 = v88 + int32(1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != int32(34) {
		v305 = v92
		v307 = v74
		v308 = v88
		goto L11
	} else {
		goto L34
	}
L33:
	;
	goto L10
L34:
	;
	if v88&int32(3) == int32(0) {
		v119 = v88
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v88 == v92 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v152 = v144 - v88
	goto L35
L37:
	;
	v123 = v119
	goto L46
L38:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v103 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v152 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v108 = v88
	goto L42
L42:
	;
	v112 = v108 + int32(1)
	if v112&int32(3) == int32(0) {
		v119 = v112
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v144 = v112
	goto L36
L44:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v117 != 0 {
		v108 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v132 = int32(-2139062144)
	if (int32(16843008)-v129|v129)&v132 == v132 {
		v123 = v123 + int32(4)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v138 = v123
	goto L49
L48:
	;
	goto L47
L49:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v142 != 0 {
		v138 = v138 + int32(1)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v144 = v138
	goto L36
L51:
	;
	goto L50
L52:
	;
	v297 = int32(34)
	v298 = F___strchrnul(m, v92, v297)
	mBase = m.M
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v300 == v297 {
		goto L99
	} else {
		goto L100
	}
L53:
	;
	goto L52
L54:
	;
	v156 = v88 + v152
	if base.Ui32(v92-v156) <= base.Ui32(int32(0)-v152<<(uint(int32(1))%32)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v163 = F___memcpy(m, v88, v92, v152)
	mBase = m.M
	goto L52
L56:
	;
	goto L57
L57:
	;
	v166 = (v88 ^ v92) & int32(3)
	if base.Ui32(v88) < base.Ui32(v92) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v268 == int32(0) {
		goto L53
	} else {
		goto L94
	}
L59:
	;
	if base.Ui32(v246) <= base.Ui32(int32(3)) {
		v267 = v245
		v268 = v246
		v269 = v247
		goto L58
	} else {
		goto L90
	}
L60:
	;
	if v166 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v166 != 0 {
		v228 = v152
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v267 = v92
	v268 = v152
	v269 = v88
	goto L58
L64:
	;
	goto L65
L65:
	;
	if v88&int32(3) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v245 = v92
	v246 = v152
	v247 = v88
	goto L59
L67:
	;
	goto L68
L68:
	;
	v173 = v92
	v174 = v152
	v175 = v88
	goto L69
L69:
	;
	if v174 == int32(0) {
		goto L53
	} else {
		goto L71
	}
L70:
	;
	v245 = v182
	v246 = v184
	v247 = v186
	goto L59
L71:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v179)
	v181 = int32(1)
	v182 = v173 + v181
	v184 = v174 - v181
	v186 = v175 + v181
	if v186&int32(3) != 0 {
		v173 = v182
		v174 = v184
		v175 = v186
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	if v228 == int32(0) {
		goto L53
	} else {
		goto L86
	}
L74:
	;
	if v156&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v193 = v152
	goto L78
L76:
	;
	v208 = v152
	goto L77
L77:
	;
	if base.Ui32(v208) <= base.Ui32(int32(3)) {
		v228 = v208
		goto L73
	} else {
		goto L82
	}
L78:
	;
	if v193 == int32(0) {
		goto L53
	} else {
		goto L80
	}
L79:
	;
	v208 = v199
	goto L77
L80:
	;
	v199 = v193 - int32(1)
	v200 = v88 + v199
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v202)
	if v200&int32(3) != 0 {
		v193 = v199
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v215 = v208
	goto L83
L83:
	;
	v219 = v215 - int32(4)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v92+v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v88+v219))) = v222
	if base.Ui32(int32(3)) < base.Ui32(v219) {
		v215 = v219
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v228 = v219
	goto L73
L85:
	;
	goto L84
L86:
	;
	v235 = v228
	goto L87
L87:
	;
	v239 = v235 - int32(1)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v239))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88+v239))) = uint8(v242)
	if v239 != 0 {
		v235 = v239
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L53
L89:
	;
	goto L88
L90:
	;
	v252 = v245
	v253 = v246
	v254 = v247
	goto L91
L91:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v256
	v258 = int32(4)
	v259 = v252 + v258
	v261 = v254 + v258
	v263 = v253 - v258
	if base.Ui32(int32(3)) < base.Ui32(v263) {
		v252 = v259
		v253 = v263
		v254 = v261
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v267 = v259
	v268 = v263
	v269 = v261
	goto L58
L93:
	;
	goto L92
L94:
	;
	v274 = v267
	v275 = v268
	v276 = v269
	goto L95
L95:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v278)
	v280 = int32(1)
	v285 = v275 - v280
	if v285 != 0 {
		v274 = v274 + v280
		v275 = v285
		v276 = v276 + v280
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L53
L97:
	;
	goto L96
L98:
	;
	if v304 != 0 {
		v88 = v304
		goto L32
	} else {
		goto L102
	}
L99:
	;
	v304 = v298
	goto L101
L100:
	;
	v304 = int32(0)
	goto L101
L101:
	;
	goto L98
L102:
	;
	goto L33
L103:
	;
	v319 = int32(*(*int8)(unsafe.Add(mBase, uint32(v311))))
	goto L105
L104:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v329 == int32(44) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if base.B2i32(v319 == int32(32))|base.B2i32(base.Ui32((v319-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v311 = v311 + int32(1)
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v337 = v311
	goto L110
L108:
	;
	goto L109
L109:
	;
	if v329 == int32(0) {
		v365 = v311
		goto L9
	} else {
		goto L114
	}
L110:
	;
	v339 = v337 + int32(1)
	v340 = int32(*(*int8)(unsafe.Add(mBase, uint32(v339))))
	goto L112
L112:
	;
	if base.B2i32(v340 == int32(32))|base.B2i32(base.Ui32((v340-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v337 = v339
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v365 = v339
	goto L9
L114:
	;
	goto L10
L115:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v424) {
		goto L132
	} else {
		goto L133
	}
L116:
	;
	v424 = v416 - v307
	goto L115
L117:
	;
	v395 = v391
	goto L126
L118:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v375 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v424 = int32(0)
	goto L115
L120:
	;
	goto L121
L121:
	;
	v380 = v307
	goto L122
L122:
	;
	v384 = v380 + int32(1)
	if v384&int32(3) == int32(0) {
		v391 = v384
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v416 = v384
	goto L116
L124:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v389 != 0 {
		v380 = v384
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v404 = int32(-2139062144)
	if (int32(16843008)-v401|v401)&v404 == v404 {
		v395 = v395 + int32(4)
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v410 = v395
	goto L129
L128:
	;
	goto L127
L129:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v414 != 0 {
		v410 = v410 + int32(1)
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v416 = v410
	goto L116
L131:
	;
	goto L130
L132:
	;
	v427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1023)) = uint8(v427)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v429 = F_pstrdup(m, v307)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	return int32(0)
L136:
	;
	F_canonicalize_path_enc(m, v429)
	mBase = m.M
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v435 = F_lappend(m, v434, v429)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v435
	if v329 != int32(44) {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v32 = v365
	v34 = v440
	goto L7
}
func F_StrategyNotifyBgWriter(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
		F_s_lock(m, v10, int32(_a_F_StrategyNotifyBgWriter_0), int32(438), int32(_a_F_StrategyNotifyBgWriter_1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
			return
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
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
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
				v40 = v15 + int64(1)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v43 == int64(0) {
					v52 = v42
				} else {
					v46 = v43 - v40
					if base.I64_extend_i32_s(v42-v41) <= v46 {
						v52 = v42
					} else {
						v52 = v41 + base.I32_wrap_i64(v46)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v52
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v40 + base.I64_extend_i32_s(v55-v41)
				if base.Ui32(v41) <= base.Ui32(v55) {
					*(*uint8)(unsafe.Add(mBase, uint32(v41-int32(1)))) = uint8(v20)
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	v77 = v40 + int32(24)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77+v22<<(uint(int32(2))%32)-int32(4))))
	v86 = v40 + v83&int32(_a_F_scanGetCandidate_1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v89)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v91)+6)))
	if v93&int32(32) != 0 {
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
	v97 = v22 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v97)
	v99 = int32(_a_F_scanGetCandidate_0)
	v100 = v97 & v99
	if base.Ui32(v49&v99) < base.Ui32(v100) {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v175 = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v175)
	goto L1
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100<<(uint(int32(2))%32)+v77-int32(4))))
	v112 = v40 + v109&int32(_a_F_scanGetCandidate_1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v115 = int32(16)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
	if v113|v114<<(uint(v115)%32) == v118|v119<<(uint(v115)%32) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v129 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	if v125 == v126 {
		v129 = int32(1)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v129 = int32(0)
	goto L27
L31:
	;
	goto L30
L32:
	;
	goto L33
L33:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v141 = v139 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v141)
	v143 = int32(_a_F_scanGetCandidate_0)
	v144 = v141 & v143
	if base.Ui32(v49&v143) < base.Ui32(v144) {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L1
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v144<<(uint(int32(2))%32)+v77-int32(4))))
	v156 = v40 + v153&int32(_a_F_scanGetCandidate_1)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v159 = int32(16)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+2)))
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156))))
	if v157|v158<<(uint(v159)%32) == v162|v163<<(uint(v159)%32) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v173 != 0 {
		goto L33
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	if v169 == v170 {
		v173 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v173 = int32(0)
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v117 int64
	_ = v117
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
	var v139 int32
	_ = v139
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v180 int64
	_ = v180
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v8 != v9 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	return v180
L2:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v165 < int64(0) {
		v180 = int64(-9223372036854775807 - 1)
		goto L1
	} else {
		goto L55
	}
L3:
	;
	if base.Ui32(v50) < base.Ui32(int32(-10)) {
		goto L2
	} else {
		goto L20
	}
L4:
	;
	v49 = v19
	v50 = v19 - int32(58)
	v51 = int32(0)
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v22 != v23 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	switch v19 - int32(43) {
	case 0, 2:
		goto L5
	default:
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = v14
	goto L6
L8:
	;
	goto L9
L9:
	;
	v15 = F___shgetc(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int64(0)
L11:
	;
	v19 = v15
	goto L6
L12:
	;
	v33 = base.B2i32(v19 == int32(45))
	v35 = v31 - int32(58)
	if l1 == int32(0) {
		v49 = v31
		v50 = v35
		v51 = v33
		goto L3
	} else {
		goto L17
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22 + int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v31 = v28
	goto L12
L14:
	;
	goto L15
L15:
	;
	v29 = F___shgetc(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v31 = v29
	goto L12
L17:
	;
	if base.Ui32(int32(-11)) < base.Ui32(v35) {
		v49 = v31
		v50 = v35
		v51 = v33
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v40 < int64(0) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - int32(1)
	goto L2
L20:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v49-int32(48)) {
		v150 = int64(0)
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if int64(0) <= v151 {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	v61 = v49
	v62 = int32(0)
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v69 != v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v88 = base.I64_extend_i32_s(v84)
	if base.Ui32(int32(10)) <= base.Ui32(v80) {
		v150 = v88
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v79 = int32(48)
	v80 = v78 - v79
	v84 = v61 + v62*int32(10) - v79
	if base.B2i32(base.Ui32(v80) <= base.Ui32(int32(9)))&base.B2i32(v84 < int32(214748364)) != 0 {
		v61 = v78
		v62 = v84
		goto L23
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69 + int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v78 = v75
	goto L25
L27:
	;
	goto L28
L28:
	;
	v76 = F___shgetc(m, l0)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v78 = v76
	goto L25
L30:
	;
	goto L24
L31:
	;
	v93 = v78
	v97 = v88
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v102 != v103 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v113) {
		v150 = v117
		goto L21
	} else {
		goto L40
	}
L34:
	;
	v113 = v111 - int32(48)
	v117 = base.I64_extend_i32_u(v93) + v97*int64(10) - int64(48)
	if base.B2i32(base.Ui32(v113) <= base.Ui32(int32(9)))&base.B2i32(v117 < int64(92233720368547758)) != 0 {
		v93 = v111
		v97 = v117
		goto L32
	} else {
		goto L39
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102 + int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v111 = v108
	goto L34
L36:
	;
	goto L37
L37:
	;
	v109 = F___shgetc(m, l0)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v111 = v109
	goto L34
L39:
	;
	goto L33
L40:
	;
	goto L41
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v130 != v131 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v150 = v117
	goto L21
L43:
	;
	if base.Ui32(v139-int32(48)) < base.Ui32(int32(10)) {
		goto L41
	} else {
		goto L48
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130 + int32(1)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v139 = v136
	goto L43
L45:
	;
	goto L46
L46:
	;
	v137 = F___shgetc(m, l0)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v139 = v137
	goto L43
L48:
	;
	goto L42
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - int32(1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v51 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v160 = int64(0) - v150
	goto L54
L53:
	;
	v160 = v150
	goto L54
L54:
	;
	v180 = v160
	goto L1
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168 - int32(1)
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
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
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
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L64
	} else {
		goto L66
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v232
L3:
	;
	v232 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 <= v18 {
		v232 = v18
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
	v223 = F_makeVarFromTargetEntry(m, l2, v39)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L64
	} else {
		goto L65
	}
L11:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v167 = int32(0)
	if v165 == v167 {
		goto L50
	} else {
		goto L51
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
	v232 = int32(0)
	goto L2
L14:
	;
	v162 = v28 + int32(1)
	if v162 != v25 {
		v28 = v162
		goto L12
	} else {
		goto L48
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
	v118 = base.B2i32(v109|v110 == v111)
	if v109 == v111 {
		v157 = v118
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
		v97 = v53
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = v97
	goto L20
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v63 < v62 {
		v97 = v53
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+(v52+v69))))
	v88 = v83 & (v85 ^ int32(-1))
	v90 = base.B2i32(v88 == int32(0))
	if v88 != 0 {
		v97 = v90
		goto L24
	} else {
		goto L32
	}
L31:
	;
	v97 = v90
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
		goto L47
	}
L36:
	;
	goto L35
L37:
	;
	if v110 == int32(0) {
		v157 = v118
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v124 != v125 {
		v157 = int32(0)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v127 = int32(1)
	if v124 <= v127 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = v127
	goto L42
L41:
	;
	v130 = v124
	goto L42
L42:
	;
	v131 = int32(8)
	v136 = int32(0)
	goto L43
L43:
	;
	v144 = v136 << (uint(int32(2)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v109+v131+v144)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v144+(v110+v131))))
	v149 = base.B2i32(v146 == v148)
	if v148 != v146 {
		v157 = v149
		goto L36
	} else {
		goto L45
	}
L44:
	;
	v157 = v149
	goto L36
L45:
	;
	v152 = v136 + int32(1)
	if v152 != v130 {
		v136 = v152
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L1
L48:
	;
	goto L13
L49:
	;
	if v220 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v220 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v166 == int32(0) {
		v211 = v167
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v220 = v211
	goto L49
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v177 < v176 {
		v211 = v167
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v179 = int32(1)
	if v176 <= v179 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v182 = v179
	goto L58
L57:
	;
	v182 = v176
	goto L58
L58:
	;
	v183 = int32(8)
	v188 = int32(0)
	goto L59
L59:
	;
	v195 = v188 << (uint(int32(2)) % 32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v165+v183+v195)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v166+v183))))
	v202 = v197 & (v199 ^ int32(-1))
	v204 = base.B2i32(v202 == int32(0))
	if v202 != 0 {
		v211 = v204
		goto L53
	} else {
		goto L61
	}
L60:
	;
	v211 = v204
	goto L53
L61:
	;
	v206 = v188 + int32(1)
	if v206 != v182 {
		v188 = v206
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L10
L64:
	;
	return int32(0)
L65:
	;
	v227 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+40)) = uint16(v227)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v227
	v232 = v223
	goto L2
L66:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v249 = F_bmsToString(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v252 = F_bmsToString(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v249
	F_errmsg_internal(m, int32(_a_F_search_indexed_tlist_for_phv_0), v12)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_search_indexed_tlist_for_phv_1), int32(2965), int32(_a_F_search_indexed_tlist_for_phv_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L64
	} else {
		goto L70
	}
L70:
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	var v180 int32
	_ = v180
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
	var v215 int32
	_ = v215
	var v221 float64
	_ = v221
	var v223 float64
	_ = v223
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 float64
	_ = v300
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = int32(12)
	v29 = v24*v25 + v25
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
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v23
	v36 = v30 + int32(12)
	if v23 == v32 {
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
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
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
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
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = base.F64_add(v300, v300)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v310 = F_fix_upper_expr_mutator(m, v299, v19)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L56
	}
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v293 = v4
	goto L27
L30:
	;
	goto L31
L31:
	;
	v154 = v4
	v157 = v4
	goto L32
L32:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v157<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v165 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v293 = v276
	goto L27
L34:
	;
	v273 = F_flatCopyTargetEntry(m, v163)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L53
	}
L35:
	;
	v251 = F_makeVarFromTargetEntry(m, int32(-2), v193)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L52
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v248 = F_fix_upper_expr_mutator(m, v235, v19)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L51
	}
L37:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v166 == int32(0) {
		v215 = v164
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v223 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v223
	v235 = v164
	goto L36
L40:
	;
	v221 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v221
	v235 = v215
	goto L36
L41:
	;
	v169 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v170 <= v169 {
		v215 = v164
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v177 = v169
	v180 = v170
	goto L43
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v177<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v165 == v194 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v215 = v204
	goto L40
L45:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v197 = F_equal(m, v164, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	v200 = v180
	goto L47
L47:
	;
	v202 = v177 + int32(1)
	if v202 < v200 {
		v177 = v202
		v180 = v200
		goto L43
	} else {
		goto L50
	}
L48:
	;
	if v197 != 0 {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v200 = v199
	goto L47
L50:
	;
	goto L44
L51:
	;
	v261 = v248
	goto L34
L52:
	;
	v253 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+40)) = uint16(v253)
	*(*int32)(unsafe.Add(mBase, uint32(v251)+36)) = v253
	v261 = v251
	goto L34
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+4)) = v261
	v276 = F_lappend(m, v154, v273)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v279 = v157 + int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v279 < v280 {
		v154 = v276
		v157 = v279
		goto L32
	} else {
		goto L55
	}
L55:
	;
	goto L33
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v310
	F_pfree(m, v30)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	m.G0 = v19 + int32(32)
	return
}
func F_setitimer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 float64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v41 int64
	_ = v41
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	v2 = int32(0)
	v5 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(v7 == v2)&base.B2i32(v10 == int64(0)) == v2 {
		v17 = int64(1000)
		v18 = v10 * v17
		v21 = int32(1000)
		v22 = base.I32_div_s(v7, v21)
		*(*float64)(unsafe.Add(mBase, _c_F_setitimer[0])) = base.F64_add(base.F64_add(v5, base.F64_convert_i64_s(v18)), base.F64_convert_i32_s(v22))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = base.I32_div_s(v27, v21)
		v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*float64)(unsafe.Add(mBase, _c_F_setitimer[1])) = base.F64_convert_i64_s(base.I64_extend_i32_s(v29) + v31*v17)
		v49 = base.F64_convert_i64_s(v18 + base.I64_extend_i32_s(v22))
	} else {
		v41 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_setitimer[0])) = v41
		*(*int64)(unsafe.Add(mBase, _c_F_setitimer[1])) = v41
		v49 = float64(0)
	}
	v50 = m.Env.X_setitimer_js(m, v2, v49)
	mBase = m.M
	return v50
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
	var v40 int32
	_ = v40
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
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v77 float64
	_ = v77
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v86 float64
	_ = v86
	var v90 int64
	_ = v90
	var v98 float64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
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
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L35
	}
L4:
	;
	v40 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pg_qsort(m, v20, l0, int32(32), int32(1477))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L34
	}
L7:
	;
	v45 = v20 + v40<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = v40
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v40))))
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
	v114 = v40 + int32(1)
	if v114 != l0 {
		v40 = v114
		goto L7
	} else {
		goto L33
	}
L10:
	;
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v51
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1+v40<<(uint(int32(2))%32))))
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	if base.F64_lt(v61, float64(0)) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if base.F64_gt(v61, float64(1)) != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if l4 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = base.F64_mul(v61, base.F64_convert_i64_s(l3-int64(1)))
	v72 = base.F64_floor(v71)
	*(*float64)(unsafe.Add(mBase, uint32(v45)+16)) = base.F64_sub(v71, v72)
	v77 = base.F64_add(base.F64_ceil(v71), float64(1))
	if base.F64_lt(base.F64_abs(v77), float64(9.223372036854776e+18)) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = int64(0)
	v98 = base.F64_ceil(base.F64_mul(v61, base.F64_convert_i64_s(l3)))
	if base.F64_lt(base.F64_abs(v98), float64(9.223372036854776e+18)) != 0 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v83
	v86 = base.F64_add(v72, float64(1))
	if base.F64_lt(base.F64_abs(v86), float64(9.223372036854776e+18)) != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v81 = base.I64_trunc_f64_s(v77)
	v83 = v81
	goto L19
L21:
	;
	goto L22
L22:
	;
	v83 = int64(-9223372036854775807 - 1)
	goto L19
L23:
	;
	v90 = base.I64_trunc_f64_s(v86)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v90
	goto L9
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(-9223372036854775807 - 1)
	goto L9
L26:
	;
	if v104 <= int64(1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v102 = base.I64_trunc_f64_s(v98)
	v104 = v102
	goto L26
L28:
	;
	goto L29
L29:
	;
	v104 = int64(-9223372036854775807 - 1)
	goto L26
L30:
	;
	v107 = int64(1)
	goto L32
L31:
	;
	v107 = v104
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v107
	goto L9
L33:
	;
	goto L8
L34:
	;
	m.G0 = v16 + int32(16)
	return v20
L35:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16))) = v61
	F_errmsg(m, int32(_a_F_setup_pct_info_0), v16)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_setup_pct_info_1), int32(692), int32(_a_F_setup_pct_info_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
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
	var v56 int32
	_ = v56
	var v74 int32
	_ = v74
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
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
											v56 = int32(1)
											if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-v56) < base.Ui32(int32(255)) {
												v74 = v56
											} else {
												if v47&int32(127) == int32(0) {
													if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&int32(255)) {
														v74 = v56
													} else {
														v74 = int32(0)
													}
												} else {
													v74 = int32(0)
												}
											}
											if v74 != 0 {
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
														v160 = m.ExcPending
														if v160 != 0 {
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
															v147 = int32(101)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
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
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															v109 = int32(_a_F_shell_archive_file_8)
															if base.Ui32(v77-int32(65)) < base.Ui32(int32(-64)) {
																v126 = v109
															} else {
																v115 = v77
																v116 = v109
																for {
																	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
																	v120 = v116 + int32(1)
																	if v118 != 0 {
																		v116 = v120
																		continue
																	} else {
																	}
																	v122 = v115 - int32(1)
																	if v122 != 0 {
																		v115 = v122
																		v116 = v120
																		continue
																	} else {
																		break
																	}
																	break
																}
																v126 = v120
															}
															if v126 != 0 {
																v128 = v126
															} else {
																v128 = int32(_a_F_shell_archive_file_9)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v128
															*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
															F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																v147 = int32(117)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
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
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
															F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
															mBase = m.M
															v144 = m.ExcPending
															if v144 != 0 {
																return int32(0)
															} else {
																v147 = int32(126)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
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
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												v165 = F_errstart(m, int32(14), int32(0))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													if v165 == int32(0) {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
														F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
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
									v56 = int32(1)
									if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-v56) < base.Ui32(int32(255)) {
										v74 = v56
									} else {
										if v47&int32(127) == int32(0) {
											if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&int32(255)) {
												v74 = v56
											} else {
												v74 = int32(0)
											}
										} else {
											v74 = int32(0)
										}
									}
									if v74 != 0 {
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
												v160 = m.ExcPending
												if v160 != 0 {
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
													v147 = int32(101)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
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
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													v109 = int32(_a_F_shell_archive_file_8)
													if base.Ui32(v77-int32(65)) < base.Ui32(int32(-64)) {
														v126 = v109
													} else {
														v115 = v77
														v116 = v109
														for {
															v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
															v120 = v116 + int32(1)
															if v118 != 0 {
																v116 = v120
																continue
															} else {
															}
															v122 = v115 - int32(1)
															if v122 != 0 {
																v115 = v122
																v116 = v120
																continue
															} else {
																break
															}
															break
														}
														v126 = v120
													}
													if v126 != 0 {
														v128 = v126
													} else {
														v128 = int32(_a_F_shell_archive_file_9)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v128
													*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
													F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														v147 = int32(117)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
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
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
													F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
													mBase = m.M
													v144 = m.ExcPending
													if v144 != 0 {
														return int32(0)
													} else {
														v147 = int32(126)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
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
									v162 = m.ExcPending
									if v162 != 0 {
										return int32(0)
									} else {
										v165 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return int32(0)
										} else {
											if v165 == int32(0) {
												m.G0 = v8 + int32(112)
												return base.B2i32(v47 == int32(0))
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
												F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
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
										v56 = int32(1)
										if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-v56) < base.Ui32(int32(255)) {
											v74 = v56
										} else {
											if v47&int32(127) == int32(0) {
												if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&int32(255)) {
													v74 = v56
												} else {
													v74 = int32(0)
												}
											} else {
												v74 = int32(0)
											}
										}
										if v74 != 0 {
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
													v160 = m.ExcPending
													if v160 != 0 {
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
														v147 = int32(101)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
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
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														v109 = int32(_a_F_shell_archive_file_8)
														if base.Ui32(v77-int32(65)) < base.Ui32(int32(-64)) {
															v126 = v109
														} else {
															v115 = v77
															v116 = v109
															for {
																v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
																v120 = v116 + int32(1)
																if v118 != 0 {
																	v116 = v120
																	continue
																} else {
																}
																v122 = v115 - int32(1)
																if v122 != 0 {
																	v115 = v122
																	v116 = v120
																	continue
																} else {
																	break
																}
																break
															}
															v126 = v120
														}
														if v126 != 0 {
															v128 = v126
														} else {
															v128 = int32(_a_F_shell_archive_file_9)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v128
														*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
														F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															v147 = int32(117)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
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
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
														F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int32(0)
														} else {
															v147 = int32(126)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
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
										v162 = m.ExcPending
										if v162 != 0 {
											return int32(0)
										} else {
											v165 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												if v165 == int32(0) {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
													F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
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
								v56 = int32(1)
								if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-v56) < base.Ui32(int32(255)) {
									v74 = v56
								} else {
									if v47&int32(127) == int32(0) {
										if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&int32(255)) {
											v74 = v56
										} else {
											v74 = int32(0)
										}
									} else {
										v74 = int32(0)
									}
								}
								if v74 != 0 {
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
											v160 = m.ExcPending
											if v160 != 0 {
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
												v147 = int32(101)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
												F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v23)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
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
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												v109 = int32(_a_F_shell_archive_file_8)
												if base.Ui32(v77-int32(65)) < base.Ui32(int32(-64)) {
													v126 = v109
												} else {
													v115 = v77
													v116 = v109
													for {
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
														v120 = v116 + int32(1)
														if v118 != 0 {
															v116 = v120
															continue
														} else {
														}
														v122 = v115 - int32(1)
														if v122 != 0 {
															v115 = v122
															v116 = v120
															continue
														} else {
															break
														}
														break
													}
													v126 = v120
												}
												if v126 != 0 {
													v128 = v126
												} else {
													v128 = int32(_a_F_shell_archive_file_9)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v128
												*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
												F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													v147 = int32(117)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
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
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
												F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
												mBase = m.M
												v144 = m.ExcPending
												if v144 != 0 {
													return int32(0)
												} else {
													v147 = int32(126)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v147, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
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
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									v165 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int32(0)
									} else {
										if v165 == int32(0) {
											m.G0 = v8 + int32(112)
											return base.B2i32(v47 == int32(0))
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
											F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(134), int32(_a_F_shell_archive_file_4))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(41), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		v27 = F_LocalToUtf(m, v6, v10, v5, int32(_a_F_shift_jis_2004_to_utf8_0), int32(_a_F_shift_jis_2004_to_utf8_1), int32(25), v17, int32(41), base.B2i32(v7 != v17))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
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
	var v63 int32
	_ = v63
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
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
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
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
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
	return v380
L2:
	;
	v356 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v356)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v369 = F_FunctionCall8Coll(m, v358, v359, v360, v361, v362, v363, v364, l0+int32(87), v367, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	v380 = int32(2)
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
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v151 = l0 + int32(87)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v154 = F_FunctionCall8Coll(m, v143, v144, v145, v146, v147, v148, v149, v151, v152, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v63 = v2
	goto L19
L17:
	;
	v97 = v52
	goto L18
L18:
	;
	if v51 == int32(0) {
		goto L14
	} else {
		goto L22
	}
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
	v94 = v63 + v91
	if v94 != v42&int32(2147483644) {
		v59 = v92
		v63 = v94
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v97 = v92
	goto L18
L21:
	;
	goto L20
L22:
	;
	v109 = v97
	v112 = v52
	goto L23
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v13+v109<<(uint(int32(2))%32))))
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v122))) = uint8(v124)
	v126 = int32(1)
	v129 = v112 + v126
	if v129 != v51 {
		v109 = v109 + v126
		v112 = v129
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
	v159 = base.B2i32(v154 != int32(0))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)))
	v165 = v160
	goto L30
L28:
	;
	if v42 <= int32(0) {
		v380 = v261
		goto L1
	} else {
		goto L49
	}
L29:
	;
	if v165&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v171 = int32(0)
	if v42 <= int32(0) {
		v202 = v171
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v261 = int32(2)
	goto L28
L32:
	;
	if v42 == v202 {
		goto L29
	} else {
		goto L40
	}
L33:
	;
	v174 = v171
	goto L34
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v13+v174<<(uint(int32(2))%32))))
	v188 = v183 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v189 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L29
L36:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v192)
	v202 = v174
	goto L32
L37:
	;
	goto L38
L38:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v194)
	v197 = v174 + int32(1)
	if v197 != v42 {
		v174 = v197
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
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
	v221 = F_FunctionCall8Coll(m, v212, v213, v214, v215, v216, v217, v218, v151, v219, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)))
	v227 = int32(0)
	if v159^base.B2i32(v221 != v227) == v227 {
		v165 = base.B2i32(v223|v165&int32(1) != v227)
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L31
L43:
	;
	v248 = int32(2)
	goto L45
L44:
	;
	v248 = v159
	goto L45
L45:
	;
	if v154 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v250 = v248
	goto L48
L47:
	;
	v250 = int32(0)
	goto L48
L48:
	;
	v261 = v250
	goto L28
L49:
	;
	v265 = v42 & int32(3)
	v266 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v274 = v266
	v277 = int32(0)
	goto L53
L51:
	;
	v312 = v266
	goto L52
L52:
	;
	if v265 == int32(0) {
		v380 = v261
		goto L1
	} else {
		goto L56
	}
L53:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v284 = int32(2)
	v286 = v13 + v274<<(uint(v284)%32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	*(*uint8)(unsafe.Add(mBase, uint32(v283+v287))) = uint8(v284)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v291+v292))) = uint8(v284)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v296+v297))) = uint8(v284)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v301+v302))) = uint8(v284)
	v306 = int32(4)
	v307 = v274 + v306
	v309 = v277 + v306
	if v309 != v42&int32(2147483644) {
		v274 = v307
		v277 = v309
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v312 = v307
	goto L52
L55:
	;
	goto L54
L56:
	;
	v324 = v312
	v326 = v266
	goto L57
L57:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v334 = int32(2)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v13+v324<<(uint(v334)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v333+v337))) = uint8(v334)
	v341 = int32(1)
	v344 = v326 + v341
	if v344 != v265 {
		v324 = v324 + v341
		v326 = v344
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v380 = v261
	goto L1
L59:
	;
	goto L58
L60:
	;
	v380 = base.B2i32(v369 != int32(0))
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v3 = int32(_a_F_shim_popen_0)
	v4 = int32(_a_F_shim_popen_1)
	if (l0^v3)&int32(3) != 0 {
		v74 = l0
		v75 = v4
		v76 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[0])) = uint8(v114)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v117 == int32(119) {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v112 = F___memset(m, v109, int32(0), v108)
	mBase = m.M
	goto L1
L3:
	;
	v108 = int32(0)
	v109 = v103
	goto L2
L4:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L23
L5:
	;
	if v75 == int32(0) {
		v103 = v76
		goto L3
	} else {
		goto L22
	}
L6:
	;
	if l0&int32(3) == int32(0) {
		v40 = l0
		v41 = v4
		v42 = v3
		v43 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v43 == int32(0) {
		v103 = v42
		goto L3
	} else {
		goto L15
	}
L8:
	;
	goto L9
L9:
	;
	v19 = l0
	v20 = v4
	v21 = v3
	goto L10
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v23)
	if v23 == int32(0) {
		v108 = v20
		v109 = v21
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v40 = v34
	v41 = v30
	v42 = v28
	v43 = v32
	goto L7
L12:
	;
	v27 = int32(1)
	v28 = v21 + v27
	v30 = v20 - v27
	v31 = int32(0)
	v32 = base.B2i32(v30 != v31)
	v34 = v19 + v27
	if v34&int32(3) == v31 {
		v40 = v34
		v41 = v30
		v42 = v28
		v43 = v32
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v30 != 0 {
		v19 = v34
		v20 = v30
		v21 = v28
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v46 == int32(0) {
		v108 = v41
		v109 = v42
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v41) < base.Ui32(int32(4)) {
		v74 = v40
		v75 = v41
		v76 = v42
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v52 = v40
	v53 = v41
	v54 = v42
	goto L18
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v81 = v52
		v82 = v53
		v83 = v54
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v74 = v68
	v75 = v70
	v76 = v66
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v57
	v65 = int32(4)
	v66 = v54 + v65
	v68 = v52 + v65
	v70 = v53 - v65
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v52 = v68
		v53 = v70
		v54 = v66
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L4
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
	if v90 == int32(0) {
		v108 = v87
		v109 = v88
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v103 = v95
	goto L3
L25:
	;
	v94 = int32(1)
	v95 = v88 + v94
	v99 = v87 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v99
		v88 = v95
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[1])) = v141
	return v141
L28:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v121)
	v125 = F_fopen(m, int32(_a_F_shim_popen_2), int32(_a_F_shim_popen_3))
	mBase = m.M
	v141 = v125
	goto L27
L29:
	;
	goto L30
L30:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v127)
	v131 = int32(_a_F_shim_popen_4)
	v132 = m.Env.Pgmem_run(m, l0, int32(_a_F_shim_popen_5), v131)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[3])) = v132 << (uint(int32(8)) % 32) & int32(_a_F_shim_popen_6)
	v140 = F_fopen(m, v131, int32(_a_F_shim_popen_7))
	mBase = m.M
	v141 = v140
	goto L27
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
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
		v46 = base.B2i32(v43 == v44)
		m.G0 = v7 + int32(16)
		return v46
	case 2:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		switch v20 - int32(114) {
		case 0:
			v46 = int32(1)
		case 1:
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
			v25 = *(*int64)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[1]))
			v46 = base.B2i32(base.Ui64(v23) <= base.Ui64(v25))
		default:
			v46 = int32(0)
		}
		m.G0 = v7 + int32(16)
		return v46
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
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[2]))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v60
					F_errmsg(m, int32(_a_F_should_apply_changes_for_rel_3), v7)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(_a_F_should_apply_changes_for_rel_4), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(476), int32(_a_F_should_apply_changes_for_rel_2))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
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
			v46 = base.B2i32(v12 == int32(114))
			m.G0 = v7 + int32(16)
			return v46
		}
	default:
		v46 = int32(0)
		m.G0 = v7 + int32(16)
		return v46
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
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
	v47 = l1
	goto L6
L6:
	;
	return v47
L7:
	;
	if v11 <= l2 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v47 = v38
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
	if int32(0) <= v21 {
		v38 = v19
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = int32(1)
	if v42 < v13 {
		v11 = v38
		v13 = v13 - v42
		goto L7
	} else {
		goto L19
	}
L13:
	;
	if v19 <= l2 {
		v38 = v19
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v26 = v19
	goto L15
L15:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v26))))
	if base.Ui32(int32(191)) < base.Ui32(v31) {
		v38 = v26
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v38 = l2
	goto L12
L17:
	;
	v35 = v26 - int32(1)
	if l2 < v35 {
		v26 = v35
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[0]))
	if v5 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(1))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
				if v12 != 0 {
					v14 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(1)
					if v15 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						F_s_lock(m, v19+int32(16), int32(_a_F_slotsync_failure_callback_0), int32(1339), int32(_a_F_slotsync_failure_callback_1))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
							v29 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v29
							*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v29)
							*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v29)
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
							m.T0[v40].(func(*base.Module, int32))(m, l1)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v29
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v29)
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
						m.T0[v40].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
					m.T0[v40].(func(*base.Module, int32))(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
			if v12 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(1)
				if v15 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					F_s_lock(m, v19+int32(16), int32(_a_F_slotsync_failure_callback_0), int32(1339), int32(_a_F_slotsync_failure_callback_1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v29
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v29)
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
						m.T0[v40].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v29
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v29)
					*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v29)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
					m.T0[v40].(func(*base.Module, int32))(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
				m.T0[v40].(func(*base.Module, int32))(m, l1)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int64
	_ = v386
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int64
	_ = v401
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int64
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int64
	_ = v631
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v711 int32
	_ = v711
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v807 int32
	_ = v807
	var v828 int32
	_ = v828
	var v849 int32
	_ = v849
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v905 int32
	_ = v905
	var v906 int64
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v955 int32
	_ = v955
	var v956 int64
	_ = v956
	var v958 int64
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int64
	_ = v997
	var v999 int64
	_ = v999
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1020 int64
	_ = v1020
	var v1022 int64
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1029 int64
	_ = v1029
	var v1031 int64
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_F_smgrdounlinkall_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v26 + int32(1)
	v30 = m.G0
	v32 = v30 - int32(112)
	m.G0 = v32
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
	m.G0 = v22 + int32(80)
	return
L4:
	;
	v36 = F_palloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v32 + int32(112)
	v874 = F_palloc(m, l1<<(uint(int32(4))%32))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L7
	} else {
		goto L172
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
	F_pfree(m, v36)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L7
	} else {
		goto L171
	}
L10:
	;
	v44 = v4
	v49 = v4
	goto L11
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+v44<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v63 != int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v164 == int32(0) {
		goto L9
	} else {
		goto L31
	}
L13:
	;
	v175 = v44 + int32(1)
	if v175 != l1 {
		v44 = v175
		v49 = v164
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[1]))
	if v63 != v67 {
		v164 = v49
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v49<<(uint(int32(2))%32)))) = v62
	v164 = v49 + int32(1)
	goto L13
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v71
	v74 = v32 + int32(72)
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	if int32(0) < v76 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[3]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v88 = int32(0)
	v96 = v80
	v100 = v76
	goto L21
L19:
	;
	goto L20
L20:
	;
	v164 = v49
	goto L13
L21:
	;
	v106 = v96 + v88<<(uint(int32(6))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v107&int32(33554432) == int32(0) {
		v125 = v96
		v126 = v100
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v128 = v88 + int32(1)
	if v128 < v126 {
		v88 = v128
		v96 = v125
		v100 = v126
		goto L21
	} else {
		goto L29
	}
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v112 != v83 {
		v125 = v96
		v126 = v100
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v114 != v82 {
		v125 = v96
		v126 = v100
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v116 != v81 {
		v125 = v96
		v126 = v100
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_InvalidateLocalBuffer(m, v106, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[3]))
	v125 = v124
	v126 = v122
	goto L23
L29:
	;
	goto L22
L30:
	;
	goto L12
L31:
	;
	v181 = F_palloc(m, v164<<(uint(int32(4))%32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if int32(0) < v164 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_pfree(m, v181)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L7
	} else {
		goto L170
	}
L34:
	;
	F_pfree(m, v181)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L110
	}
L35:
	;
	v189 = int32(0)
	v204 = int64(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if base.Ui32(int32(63)) <= base.Ui32(v414+int32(31)) {
		goto L33
	} else {
		goto L109
	}
L38:
	;
	v207 = v181 + v189<<(uint(int32(4))%32)
	v210 = v36 + v189<<(uint(int32(2))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v214 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	v320 = base.I32_div_s(v318, int32(32))
	if base.Ui64(base.I64_extend_i32_s(v320)) <= base.Ui64(v312) {
		goto L34
	} else {
		goto L89
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v225
	if v225 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(0))+20))
	if v220 != int32(-1) {
		v225 = v220
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v225 = int32(-1)
	goto L41
L45:
	;
	goto L44
L46:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v241 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v231 = F_smgrexists(m, v229, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v237 = v204 + base.I64_extend_i32_u(v225)
	goto L46
L50:
	;
	if v231 == int32(0) {
		v237 = v204
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L34
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v252
	if v252 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(4))+20))
	if v247 != int32(-1) {
		v252 = v247
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v252 = int32(-1)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v266 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v262 = v237 + base.I64_extend_i32_u(v252)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v260 = F_smgrexists(m, v258, int32(1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	if v260 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	v262 = v237
	goto L58
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v277
	if v277 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	goto L64
L66:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v263+int32(8))+20))
	if v272 != int32(-1) {
		v277 = v272
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v277 = int32(-1)
	goto L65
L69:
	;
	goto L68
L70:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v291 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v287 = v262 + base.I64_extend_i32_u(v277)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v285 = F_smgrexists(m, v283, int32(2))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	if v285 != 0 {
		goto L34
	} else {
		goto L75
	}
L75:
	;
	v287 = v262
	goto L70
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = v302
	if v302 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v288+int32(12))+20))
	if v297 != int32(-1) {
		v302 = v297
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v302 = int32(-1)
	goto L77
L81:
	;
	goto L80
L82:
	;
	v314 = v189 + int32(1)
	if v314 < v164 {
		v189 = v314
		v204 = v312
		goto L38
	} else {
		goto L88
	}
L83:
	;
	v312 = v287 + base.I64_extend_i32_u(v302)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v310 = F_smgrexists(m, v308, int32(3))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	if v310 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	v312 = v287
	goto L82
L88:
	;
	goto L39
L89:
	;
	v327 = int32(0)
	goto L90
L90:
	;
	v344 = v36 + v327<<(uint(int32(2))%32)
	v347 = v181 + v327<<(uint(int32(4))%32)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v348 != int32(-1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L33
L92:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32-int32(-64)))) = v354
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v353)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = v356
	v360 = int32(0)
	F_FindAndDropRelationBuffers(m, v32+int32(56), v360, v348, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v365 != int32(-1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v368)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v371
	F_FindAndDropRelationBuffers(m, v32+int32(40), int32(1), v365, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v380 != int32(-1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v384
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v386
	F_FindAndDropRelationBuffers(m, v32+int32(24), int32(2), v380, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	if v395 != int32(-1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v399
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v401
	F_FindAndDropRelationBuffers(m, v32+int32(8), int32(3), v395, int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v411 = v327 + int32(1)
	if v411 != v164 {
		v327 = v411
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
	v442 = F_palloc(m, v164*int32(12))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v444 = int32(0)
	if v164 <= v444 {
		v562 = v444
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if int32(0) < v564 {
		goto L125
	} else {
		goto L126
	}
L113:
	;
	v447 = int32(1)
	v449 = int32(0)
	if v164 != v447 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v458 = int32(0)
	v459 = v449
	goto L117
L115:
	;
	v507 = v449
	goto L116
L116:
	;
	if v164&v447 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v474 = int32(12)
	v476 = v442 + v459*v474
	v477 = int32(2)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v36+v459<<(uint(v477)%32))))
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v480)))
	*(*int64)(unsafe.Add(mBase, uint32(v476))) = v481
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v480)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v476)+8)) = v483
	v486 = v459 | int32(1)
	v489 = v442 + v486*v474
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v36+v486<<(uint(v477)%32))))
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v493)))
	*(*int64)(unsafe.Add(mBase, uint32(v489))) = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v493)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v489)+8)) = v496
	v499 = v459 + v477
	v501 = v458 + v477
	if v501 != v164&int32(2147483646) {
		v458 = v501
		v459 = v499
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v507 = v499
	goto L116
L119:
	;
	goto L118
L120:
	;
	v524 = v442 + v507*int32(12)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v36+v507<<(uint(int32(2))%32))))
	v529 = *(*int64)(unsafe.Add(mBase, uint32(v528)))
	*(*int64)(unsafe.Add(mBase, uint32(v524))) = v529
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v524)+8)) = v531
	goto L122
L121:
	;
	goto L122
L122:
	;
	if v164 < int32(21) {
		v562 = int32(0)
		goto L112
	} else {
		goto L123
	}
L123:
	;
	F_pg_qsort(m, v442, v164, int32(12), int32(1078))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	v562 = int32(1)
	goto L112
L125:
	;
	v579 = int32(0)
	goto L128
L126:
	;
	goto L127
L127:
	;
	F_pfree(m, v442)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L7
	} else {
		goto L169
	}
L128:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[6]))
	v591 = v588 + v579<<(uint(int32(6))%32)
	if v562 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	goto L127
L130:
	;
	v783 = v579 + int32(1)
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if v783 < v785 {
		v579 = v783
		goto L128
	} else {
		goto L168
	}
L131:
	;
	if v647 == int32(0) {
		goto L130
	} else {
		goto L144
	}
L132:
	;
	if v164 <= int32(0) {
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v591)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v632
	*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v631
	v639 = F_bsearch(m, v32+int32(88), v442, v164, int32(12), int32(1078))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L7
	} else {
		goto L143
	}
L135:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	v602 = int32(0)
	goto L136
L136:
	;
	v619 = v442 + v602*int32(12)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	if v596 != v620 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L130
L138:
	;
	v629 = v602 + int32(1)
	if v629 != v164 {
		v602 = v629
		goto L136
	} else {
		goto L142
	}
L139:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	if v622 != v623 {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	if v625 == v626 {
		v647 = v619
		goto L131
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	goto L137
L143:
	;
	v647 = v639
	goto L131
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = int32(_a_F_smgrdounlinkall_1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+104)) = int32(_a_F_smgrdounlinkall_2)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = int32(_a_F_smgrdounlinkall_3)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = int64(0)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v591)+24))
	v673 = int32(_a_F_smgrdounlinkall_4)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+24)) = v672 | v673
	if v672&v673 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	goto L148
L146:
	;
	v711 = v672
	goto L147
L147:
	;
	v729 = int32(_a_F_smgrdounlinkall_5)
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7]))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(88))+8))
	if v732 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L148:
	;
	F_perform_spin_delay(m, v32+int32(88))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L7
	} else {
		goto L150
	}
L149:
	;
	v711 = v701
	goto L147
L150:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v591)+24))
	v702 = int32(_a_F_smgrdounlinkall_4)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+24)) = v701 | v702
	if v701&v702 != 0 {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v647)))
	if v749 != v750 {
		goto L163
	} else {
		goto L164
	}
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7])) = v747
	goto L153
L155:
	;
	if int32(999) < v730 {
		goto L153
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if v730 < int32(11) {
		goto L153
	} else {
		goto L162
	}
L158:
	;
	v737 = int32(900)
	if v737 <= v730 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v740 = v737
	goto L161
L160:
	;
	v740 = v730
	goto L161
L161:
	;
	v747 = v740 + int32(100)
	goto L154
L162:
	;
	v747 = v730 - int32(1)
	goto L154
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v591)+24)) = v711 & int32(-4194305)
	goto L130
L164:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v647)+4))
	if v752 != v753 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v647)+8))
	if v755 != v756 {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	F_InvalidateBuffer(m, v591)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	goto L130
L168:
	;
	goto L129
L169:
	;
	goto L9
L170:
	;
	goto L9
L171:
	;
	goto L6
L172:
	;
	if int32(0) < l1 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v890 = v4
	goto L176
L174:
	;
	goto L175
L175:
	;
	F_pfree(m, v874)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L7
	} else {
		goto L194
	}
L176:
	;
	v899 = l0 + v890<<(uint(int32(2))%32)
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+36))
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v900)))
	v905 = v874 + v890<<(uint(int32(4))%32)
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v900)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v905)+8)) = v906
	*(*int64)(unsafe.Add(mBase, uint32(v905))) = v902
	v909 = int32(0)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v901*int32(80))+uint32(_c_F_smgrdounlinkall[8])))
	m.T0[v916].(func(*base.Module, int32, int32))(m, v910, v909)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L7
	} else {
		goto L178
	}
L177:
	;
	v937 = v909
	goto L183
L178:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	m.T0[v916].(func(*base.Module, int32, int32))(m, v919, int32(1))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L7
	} else {
		goto L179
	}
L179:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	m.T0[v916].(func(*base.Module, int32, int32))(m, v923, int32(2))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	m.T0[v916].(func(*base.Module, int32, int32))(m, v927, int32(3))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	v932 = v890 + int32(1)
	if v932 != l1 {
		v890 = v932
		goto L176
	} else {
		goto L182
	}
L182:
	;
	goto L177
L183:
	;
	v955 = v874 + v937<<(uint(int32(4))%32)
	v956 = *(*int64)(unsafe.Add(mBase, uint32(v955)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v956
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v955)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v958
	F_CacheInvalidateSmgr(m, v22-int32(-64))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L7
	} else {
		goto L185
	}
L184:
	;
	v980 = int32(0)
	goto L187
L185:
	;
	v965 = v937 + int32(1)
	if v965 != l1 {
		v937 = v965
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0+v980<<(uint(int32(2))%32))))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v990)+36))
	v994 = v874 + v980<<(uint(int32(4))%32)
	v996 = v994 + int32(8)
	v997 = *(*int64)(unsafe.Add(mBase, uint32(v996)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v997
	v999 = *(*int64)(unsafe.Add(mBase, uint32(v994)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v999
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v991*int32(80))+uint32(_c_F_smgrdounlinkall[9])))
	m.T0[v1008].(func(*base.Module, int32, int32, int32))(m, v22+int32(48), int32(0), l2)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L7
	} else {
		goto L189
	}
L188:
	;
	goto L175
L189:
	;
	v1011 = *(*int64)(unsafe.Add(mBase, uint32(v996)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v1011
	v1013 = *(*int64)(unsafe.Add(mBase, uint32(v994)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v1013
	m.T0[v1008].(func(*base.Module, int32, int32, int32))(m, v22+int32(32), int32(1), l2)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(v996)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v1020
	v1022 = *(*int64)(unsafe.Add(mBase, uint32(v994)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v1022
	m.T0[v1008].(func(*base.Module, int32, int32, int32))(m, v22+int32(16), int32(2), l2)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L7
	} else {
		goto L191
	}
L191:
	;
	v1029 = *(*int64)(unsafe.Add(mBase, uint32(v996)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v1029
	v1031 = *(*int64)(unsafe.Add(mBase, uint32(v994)))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v1031
	m.T0[v1008].(func(*base.Module, int32, int32, int32))(m, v22, int32(3), l2)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L7
	} else {
		goto L192
	}
L192:
	;
	v1037 = v980 + int32(1)
	if v1037 != l1 {
		v980 = v1037
		goto L187
	} else {
		goto L193
	}
L193:
	;
	goto L188
L194:
	;
	v1060 = int32(_a_F_smgrdounlinkall_0)
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v1062 - int32(1)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v6 = int32(_a_F_smgrwritev_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
	v9 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v8 + v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(80))+uint32(_c_F_smgrwritev[1])))
	m.T0[v18].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, v9, l4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(_a_F_smgrwritev_0)
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v23 - int32(1)
		return
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_copy(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		if v14 != 0 {
			v15 = F_array_contains_nulls(m, v10)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_sort_asc_0), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_sort_asc_1), int32(234), int32(_a_F_sort_asc_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v20 = F_ArrayGetNItems(m, v17, v10+int32(16))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v22)
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						if v24 != 0 {
							v32 = v24
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v32+v10, v20, v7+int32(15))
						mBase = m.M
						m.G0 = v7 + int32(16)
						return v10
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v20 = F_ArrayGetNItems(m, v17, v10+int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v22)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v24 != 0 {
					v32 = v24
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v32+v10, v20, v7+int32(15))
				mBase = m.M
				m.G0 = v7 + int32(16)
				return v10
			}
		}
	}
}
func F_sort_checkpoint_bufferids(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v63 int32
	_ = v63
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
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
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
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
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v649 int64
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v661 int32
	_ = v661
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int64
	_ = v728
	var v730 int64
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v740 int64
	_ = v740
	var v742 int32
	_ = v742
	var v744 int64
	_ = v744
	var v746 int64
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v778 int32
	_ = v778
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int64
	_ = v816
	var v818 int64
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int64
	_ = v826
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v832 int64
	_ = v832
	var v834 int64
	_ = v834
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v884 int32
	_ = v884
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v904 int64
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v915 int64
	_ = v915
	var v917 int32
	_ = v917
	var v919 int64
	_ = v919
	var v921 int64
	_ = v921
	var v924 int32
	_ = v924
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v963 int32
	_ = v963
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int64
	_ = v985
	var v987 int64
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int64
	_ = v996
	var v998 int64
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int64
	_ = v1002
	var v1004 int64
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int64
	_ = v1052
	var v1054 int64
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int64
	_ = v1068
	var v1070 int64
	_ = v1070
	var v1072 int32
	_ = v1072
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = l0
	v24 = l1
	goto L1
L1:
	;
	v42 = v23 + int32(20)
	v44 = v24
	goto L3
L2:
	;
	m.G0 = v21 + int32(32)
	return
L3:
	;
	v63 = v23 + v44*int32(20)
	if base.Ui32(v44) <= base.Ui32(int32(6)) {
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
	if base.Ui32(v63) <= base.Ui32(v42) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(v63) <= base.Ui32(v42) {
		goto L5
	} else {
		goto L26
	}
L9:
	;
	v81 = v42
	goto L10
L10:
	;
	if base.Ui32(v81) <= base.Ui32(v23) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v179 = v81 + int32(20)
	if base.Ui32(v179) < base.Ui32(v63) {
		v81 = v179
		goto L10
	} else {
		goto L25
	}
L13:
	;
	v89 = v81
	goto L14
L14:
	;
	v105 = v89 - int32(20)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if base.Ui32(v106) < base.Ui32(v107) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	if base.Ui32(v107) < base.Ui32(v106) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v130 = v21 + int32(24)
	v131 = int32(16)
	v132 = v89 + v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v133
	v136 = v21 + v131
	v137 = int32(8)
	v138 = v89 + v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v141
	v144 = v105 + v131
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v145
	v148 = v105 + v137
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v157
	if base.Ui32(v23) < base.Ui32(v105) {
		v89 = v105
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(16))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if base.Ui32(v112) < base.Ui32(v113) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v113) < base.Ui32(v112) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(12))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v118 < v119 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v119 < v118 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(8))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if base.Ui32(v124) <= base.Ui32(v125) {
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
	v184 = v42
	goto L27
L27:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(20))))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if base.Ui32(v202) < base.Ui32(v203) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v234 = v23 + int32(base.Ui32(v44)>>(uint(int32(1))%32))*int32(20)
	if v44 != int32(7) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	goto L28
L30:
	;
	v226 = v184 + int32(20)
	if base.Ui32(v226) < base.Ui32(v63) {
		v184 = v226
		goto L27
	} else {
		goto L38
	}
L31:
	;
	if base.Ui32(v203) < base.Ui32(v202) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(16))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if base.Ui32(v208) < base.Ui32(v209) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v209) < base.Ui32(v208) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(12))))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	if v214 < v215 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v215 < v214 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(8))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	if base.Ui32(v221) < base.Ui32(v220) {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	goto L5
L39:
	;
	v238 = v63 - int32(20)
	if base.Ui32(v44) < base.Ui32(int32(41)) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v631 = v234
	goto L41
L41:
	;
	v638 = v21 + int32(24)
	v639 = int32(16)
	v640 = v23 + v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v641
	v644 = v21 + v639
	v645 = int32(8)
	v646 = v23 + v645
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v646)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v647
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v649
	v652 = v631 + v639
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)))
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v653
	v656 = v631 + v645
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v656)))
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = v657
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v631)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v659
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v652))) = v661
	v663 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v656))) = v663
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v631))) = v665
	v668 = v63 - int32(20)
	v671 = v668
	v672 = v42
	v678 = v42
	v682 = v668
	goto L322
L42:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	if base.Ui32(v543) < base.Ui32(v544) {
		goto L256
	} else {
		goto L257
	}
L43:
	;
	v533 = v234
	v534 = v238
	v537 = v23
	goto L42
L44:
	;
	goto L45
L45:
	;
	v242 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	v244 = v242 * int32(20)
	v245 = v23 + v244
	v248 = v23 + v242*int32(40)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if base.Ui32(v253) < base.Ui32(v254) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v342 = v242 * int32(-20)
	v343 = v234 + v342
	v344 = v234 + v244
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if base.Ui32(v349) < base.Ui32(v350) {
		goto L118
	} else {
		goto L119
	}
L47:
	;
	v340 = v332
	goto L46
L48:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if base.Ui32(v254) < base.Ui32(v302) {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if base.Ui32(v254) < base.Ui32(v270) {
		v332 = v245
		goto L47
	} else {
		goto L57
	}
L50:
	;
	if base.Ui32(v254) < base.Ui32(v253) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if base.Ui32(v257) < base.Ui32(v258) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(v258) < base.Ui32(v257) {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	if v261 < v262 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if v262 < v261 {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if base.Ui32(v266) <= base.Ui32(v265) {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	if base.Ui32(v270) < base.Ui32(v254) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v253) < base.Ui32(v270) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if base.Ui32(v273) < base.Ui32(v274) {
		v332 = v245
		goto L47
	} else {
		goto L60
	}
L60:
	;
	if base.Ui32(v274) < base.Ui32(v273) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	if v277 < v278 {
		v332 = v245
		goto L47
	} else {
		goto L62
	}
L62:
	;
	if v278 < v277 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	if base.Ui32(v281) < base.Ui32(v282) {
		v332 = v245
		goto L47
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	v340 = v248
	goto L46
L66:
	;
	goto L67
L67:
	;
	if base.Ui32(v270) < base.Ui32(v253) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v340 = v23
	goto L46
L69:
	;
	goto L70
L70:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if base.Ui32(v288) < base.Ui32(v289) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v340 = v248
	goto L46
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(v289) < base.Ui32(v288) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v340 = v23
	goto L46
L75:
	;
	goto L76
L76:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	if v292 < v293 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v340 = v248
	goto L46
L78:
	;
	goto L79
L79:
	;
	if v293 < v292 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v340 = v23
	goto L46
L81:
	;
	goto L82
L82:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	if base.Ui32(v296) < base.Ui32(v297) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v299 = v248
	goto L85
L84:
	;
	v299 = v23
	goto L85
L85:
	;
	v340 = v299
	goto L46
L86:
	;
	if base.Ui32(v253) < base.Ui32(v302) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	if base.Ui32(v302) < base.Ui32(v254) {
		v332 = v245
		goto L47
	} else {
		goto L88
	}
L88:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if base.Ui32(v305) < base.Ui32(v306) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(v306) < base.Ui32(v305) {
		v332 = v245
		goto L47
	} else {
		goto L90
	}
L90:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	if v309 < v310 {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	if v310 < v309 {
		v332 = v245
		goto L47
	} else {
		goto L92
	}
L92:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	if base.Ui32(v314) < base.Ui32(v313) {
		v332 = v245
		goto L47
	} else {
		goto L93
	}
L93:
	;
	goto L86
L94:
	;
	v340 = v23
	goto L46
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(v302) < base.Ui32(v253) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v340 = v248
	goto L46
L98:
	;
	goto L99
L99:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if base.Ui32(v320) < base.Ui32(v321) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v340 = v23
	goto L46
L101:
	;
	goto L102
L102:
	;
	if base.Ui32(v321) < base.Ui32(v320) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v340 = v248
	goto L46
L104:
	;
	goto L105
L105:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	if v324 < v325 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v340 = v23
	goto L46
L107:
	;
	goto L108
L108:
	;
	if v325 < v324 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v340 = v248
	goto L46
L110:
	;
	goto L111
L111:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	if base.Ui32(v328) < base.Ui32(v329) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v331 = v23
	goto L114
L113:
	;
	v331 = v248
	goto L114
L114:
	;
	v332 = v331
	goto L47
L115:
	;
	v439 = v238 + v242*int32(-40)
	v440 = v238 + v342
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if base.Ui32(v445) < base.Ui32(v446) {
		goto L187
	} else {
		goto L188
	}
L116:
	;
	v436 = v428
	goto L115
L117:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if base.Ui32(v350) < base.Ui32(v398) {
		goto L155
	} else {
		goto L156
	}
L118:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if base.Ui32(v350) < base.Ui32(v366) {
		v428 = v234
		goto L116
	} else {
		goto L126
	}
L119:
	;
	if base.Ui32(v350) < base.Ui32(v349) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if base.Ui32(v353) < base.Ui32(v354) {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	if base.Ui32(v354) < base.Ui32(v353) {
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	if v357 < v358 {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	if v358 < v357 {
		goto L117
	} else {
		goto L124
	}
L124:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	if base.Ui32(v362) <= base.Ui32(v361) {
		goto L117
	} else {
		goto L125
	}
L125:
	;
	goto L118
L126:
	;
	if base.Ui32(v366) < base.Ui32(v350) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if base.Ui32(v349) < base.Ui32(v366) {
		goto L134
	} else {
		goto L135
	}
L128:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if base.Ui32(v369) < base.Ui32(v370) {
		v428 = v234
		goto L116
	} else {
		goto L129
	}
L129:
	;
	if base.Ui32(v370) < base.Ui32(v369) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v373 < v374 {
		v428 = v234
		goto L116
	} else {
		goto L131
	}
L131:
	;
	if v374 < v373 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	if base.Ui32(v377) < base.Ui32(v378) {
		v428 = v234
		goto L116
	} else {
		goto L133
	}
L133:
	;
	goto L127
L134:
	;
	v436 = v344
	goto L115
L135:
	;
	goto L136
L136:
	;
	if base.Ui32(v366) < base.Ui32(v349) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v436 = v343
	goto L115
L138:
	;
	goto L139
L139:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if base.Ui32(v384) < base.Ui32(v385) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v436 = v344
	goto L115
L141:
	;
	goto L142
L142:
	;
	if base.Ui32(v385) < base.Ui32(v384) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v436 = v343
	goto L115
L144:
	;
	goto L145
L145:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v388 < v389 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v436 = v344
	goto L115
L147:
	;
	goto L148
L148:
	;
	if v389 < v388 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v436 = v343
	goto L115
L150:
	;
	goto L151
L151:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	if base.Ui32(v392) < base.Ui32(v393) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v395 = v344
	goto L154
L153:
	;
	v395 = v343
	goto L154
L154:
	;
	v436 = v395
	goto L115
L155:
	;
	if base.Ui32(v349) < base.Ui32(v398) {
		goto L163
	} else {
		goto L164
	}
L156:
	;
	if base.Ui32(v398) < base.Ui32(v350) {
		v428 = v234
		goto L116
	} else {
		goto L157
	}
L157:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if base.Ui32(v401) < base.Ui32(v402) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	if base.Ui32(v402) < base.Ui32(v401) {
		v428 = v234
		goto L116
	} else {
		goto L159
	}
L159:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v405 < v406 {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	if v406 < v405 {
		v428 = v234
		goto L116
	} else {
		goto L161
	}
L161:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	if base.Ui32(v410) < base.Ui32(v409) {
		v428 = v234
		goto L116
	} else {
		goto L162
	}
L162:
	;
	goto L155
L163:
	;
	v436 = v343
	goto L115
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v398) < base.Ui32(v349) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v436 = v344
	goto L115
L167:
	;
	goto L168
L168:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if base.Ui32(v416) < base.Ui32(v417) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v436 = v343
	goto L115
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(v417) < base.Ui32(v416) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v436 = v344
	goto L115
L173:
	;
	goto L174
L174:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
	if v420 < v421 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v436 = v343
	goto L115
L176:
	;
	goto L177
L177:
	;
	if v421 < v420 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v436 = v344
	goto L115
L179:
	;
	goto L180
L180:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	if base.Ui32(v424) < base.Ui32(v425) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v427 = v343
	goto L183
L182:
	;
	v427 = v344
	goto L183
L183:
	;
	v428 = v427
	goto L116
L184:
	;
	v533 = v436
	v534 = v532
	v537 = v340
	goto L42
L185:
	;
	v532 = v524
	goto L184
L186:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if base.Ui32(v446) < base.Ui32(v494) {
		goto L224
	} else {
		goto L225
	}
L187:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if base.Ui32(v446) < base.Ui32(v462) {
		v524 = v440
		goto L185
	} else {
		goto L195
	}
L188:
	;
	if base.Ui32(v446) < base.Ui32(v445) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if base.Ui32(v449) < base.Ui32(v450) {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	if base.Ui32(v450) < base.Ui32(v449) {
		goto L186
	} else {
		goto L191
	}
L191:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	if v453 < v454 {
		goto L187
	} else {
		goto L192
	}
L192:
	;
	if v454 < v453 {
		goto L186
	} else {
		goto L193
	}
L193:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	if base.Ui32(v458) <= base.Ui32(v457) {
		goto L186
	} else {
		goto L194
	}
L194:
	;
	goto L187
L195:
	;
	if base.Ui32(v462) < base.Ui32(v446) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	if base.Ui32(v445) < base.Ui32(v462) {
		goto L203
	} else {
		goto L204
	}
L197:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v465) < base.Ui32(v466) {
		v524 = v440
		goto L185
	} else {
		goto L198
	}
L198:
	;
	if base.Ui32(v466) < base.Ui32(v465) {
		goto L196
	} else {
		goto L199
	}
L199:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v469 < v470 {
		v524 = v440
		goto L185
	} else {
		goto L200
	}
L200:
	;
	if v470 < v469 {
		goto L196
	} else {
		goto L201
	}
L201:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	if base.Ui32(v473) < base.Ui32(v474) {
		v524 = v440
		goto L185
	} else {
		goto L202
	}
L202:
	;
	goto L196
L203:
	;
	v532 = v238
	goto L184
L204:
	;
	goto L205
L205:
	;
	if base.Ui32(v462) < base.Ui32(v445) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v532 = v439
	goto L184
L207:
	;
	goto L208
L208:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v480) < base.Ui32(v481) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v532 = v238
	goto L184
L210:
	;
	goto L211
L211:
	;
	if base.Ui32(v481) < base.Ui32(v480) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v532 = v439
	goto L184
L213:
	;
	goto L214
L214:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v484 < v485 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v532 = v238
	goto L184
L216:
	;
	goto L217
L217:
	;
	if v485 < v484 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v532 = v439
	goto L184
L219:
	;
	goto L220
L220:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	if base.Ui32(v488) < base.Ui32(v489) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v491 = v238
	goto L223
L222:
	;
	v491 = v439
	goto L223
L223:
	;
	v532 = v491
	goto L184
L224:
	;
	if base.Ui32(v445) < base.Ui32(v494) {
		goto L232
	} else {
		goto L233
	}
L225:
	;
	if base.Ui32(v494) < base.Ui32(v446) {
		v524 = v440
		goto L185
	} else {
		goto L226
	}
L226:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v497) < base.Ui32(v498) {
		goto L224
	} else {
		goto L227
	}
L227:
	;
	if base.Ui32(v498) < base.Ui32(v497) {
		v524 = v440
		goto L185
	} else {
		goto L228
	}
L228:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v501 < v502 {
		goto L224
	} else {
		goto L229
	}
L229:
	;
	if v502 < v501 {
		v524 = v440
		goto L185
	} else {
		goto L230
	}
L230:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	if base.Ui32(v506) < base.Ui32(v505) {
		v524 = v440
		goto L185
	} else {
		goto L231
	}
L231:
	;
	goto L224
L232:
	;
	v532 = v439
	goto L184
L233:
	;
	goto L234
L234:
	;
	if base.Ui32(v494) < base.Ui32(v445) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v532 = v238
	goto L184
L236:
	;
	goto L237
L237:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if base.Ui32(v512) < base.Ui32(v513) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v532 = v439
	goto L184
L239:
	;
	goto L240
L240:
	;
	if base.Ui32(v513) < base.Ui32(v512) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v532 = v238
	goto L184
L242:
	;
	goto L243
L243:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v516 < v517 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v532 = v439
	goto L184
L245:
	;
	goto L246
L246:
	;
	if v517 < v516 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v532 = v238
	goto L184
L248:
	;
	goto L249
L249:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	if base.Ui32(v520) < base.Ui32(v521) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v523 = v439
	goto L252
L251:
	;
	v523 = v238
	goto L252
L252:
	;
	v524 = v523
	goto L185
L253:
	;
	v631 = v630
	goto L41
L254:
	;
	v630 = v622
	goto L253
L255:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	if base.Ui32(v544) < base.Ui32(v592) {
		goto L293
	} else {
		goto L294
	}
L256:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	if base.Ui32(v544) < base.Ui32(v560) {
		v622 = v533
		goto L254
	} else {
		goto L264
	}
L257:
	;
	if base.Ui32(v544) < base.Ui32(v543) {
		goto L255
	} else {
		goto L258
	}
L258:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if base.Ui32(v547) < base.Ui32(v548) {
		goto L256
	} else {
		goto L259
	}
L259:
	;
	if base.Ui32(v548) < base.Ui32(v547) {
		goto L255
	} else {
		goto L260
	}
L260:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	if v551 < v552 {
		goto L256
	} else {
		goto L261
	}
L261:
	;
	if v552 < v551 {
		goto L255
	} else {
		goto L262
	}
L262:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	if base.Ui32(v556) <= base.Ui32(v555) {
		goto L255
	} else {
		goto L263
	}
L263:
	;
	goto L256
L264:
	;
	if base.Ui32(v560) < base.Ui32(v544) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	if base.Ui32(v543) < base.Ui32(v560) {
		goto L272
	} else {
		goto L273
	}
L266:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if base.Ui32(v563) < base.Ui32(v564) {
		v622 = v533
		goto L254
	} else {
		goto L267
	}
L267:
	;
	if base.Ui32(v564) < base.Ui32(v563) {
		goto L265
	} else {
		goto L268
	}
L268:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v567 < v568 {
		v622 = v533
		goto L254
	} else {
		goto L269
	}
L269:
	;
	if v568 < v567 {
		goto L265
	} else {
		goto L270
	}
L270:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if base.Ui32(v571) < base.Ui32(v572) {
		v622 = v533
		goto L254
	} else {
		goto L271
	}
L271:
	;
	goto L265
L272:
	;
	v630 = v534
	goto L253
L273:
	;
	goto L274
L274:
	;
	if base.Ui32(v560) < base.Ui32(v543) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v630 = v537
	goto L253
L276:
	;
	goto L277
L277:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if base.Ui32(v578) < base.Ui32(v579) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v630 = v534
	goto L253
L279:
	;
	goto L280
L280:
	;
	if base.Ui32(v579) < base.Ui32(v578) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v630 = v537
	goto L253
L282:
	;
	goto L283
L283:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v582 < v583 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v630 = v534
	goto L253
L285:
	;
	goto L286
L286:
	;
	if v583 < v582 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v630 = v537
	goto L253
L288:
	;
	goto L289
L289:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if base.Ui32(v586) < base.Ui32(v587) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v589 = v534
	goto L292
L291:
	;
	v589 = v537
	goto L292
L292:
	;
	v630 = v589
	goto L253
L293:
	;
	if base.Ui32(v543) < base.Ui32(v592) {
		goto L301
	} else {
		goto L302
	}
L294:
	;
	if base.Ui32(v592) < base.Ui32(v544) {
		v622 = v533
		goto L254
	} else {
		goto L295
	}
L295:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if base.Ui32(v595) < base.Ui32(v596) {
		goto L293
	} else {
		goto L296
	}
L296:
	;
	if base.Ui32(v596) < base.Ui32(v595) {
		v622 = v533
		goto L254
	} else {
		goto L297
	}
L297:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v599 < v600 {
		goto L293
	} else {
		goto L298
	}
L298:
	;
	if v600 < v599 {
		v622 = v533
		goto L254
	} else {
		goto L299
	}
L299:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if base.Ui32(v604) < base.Ui32(v603) {
		v622 = v533
		goto L254
	} else {
		goto L300
	}
L300:
	;
	goto L293
L301:
	;
	v630 = v537
	goto L253
L302:
	;
	goto L303
L303:
	;
	if base.Ui32(v592) < base.Ui32(v543) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v630 = v534
	goto L253
L305:
	;
	goto L306
L306:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if base.Ui32(v610) < base.Ui32(v611) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v630 = v537
	goto L253
L308:
	;
	goto L309
L309:
	;
	if base.Ui32(v611) < base.Ui32(v610) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v630 = v534
	goto L253
L311:
	;
	goto L312
L312:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v614 < v615 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v630 = v537
	goto L253
L314:
	;
	goto L315
L315:
	;
	if v615 < v614 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v630 = v534
	goto L253
L317:
	;
	goto L318
L318:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	if base.Ui32(v618) < base.Ui32(v619) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v621 = v537
	goto L321
L320:
	;
	v621 = v534
	goto L321
L321:
	;
	v622 = v621
	goto L254
L322:
	;
	if base.Ui32(v671) < base.Ui32(v672) {
		v760 = v672
		v766 = v678
		goto L324
	} else {
		goto L325
	}
L324:
	;
	if base.Ui32(v760) <= base.Ui32(v671) {
		goto L339
	} else {
		goto L340
	}
L325:
	;
	v691 = v672
	v697 = v678
	goto L326
L326:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v706) < base.Ui32(v707) {
		v752 = v697
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v760 = v755
	v766 = v752
	goto L324
L328:
	;
	v755 = v691 + int32(20)
	if base.Ui32(v755) <= base.Ui32(v671) {
		v691 = v755
		v697 = v752
		goto L326
	} else {
		goto L337
	}
L329:
	;
	if base.Ui32(v707) < base.Ui32(v706) {
		v760 = v691
		v766 = v697
		goto L324
	} else {
		goto L330
	}
L330:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v710) < base.Ui32(v711) {
		v752 = v697
		goto L328
	} else {
		goto L331
	}
L331:
	;
	if base.Ui32(v711) < base.Ui32(v710) {
		v760 = v691
		v766 = v697
		goto L324
	} else {
		goto L332
	}
L332:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v691)+8))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v714 < v715 {
		v752 = v697
		goto L328
	} else {
		goto L333
	}
L333:
	;
	if v715 < v714 {
		v760 = v691
		v766 = v697
		goto L324
	} else {
		goto L334
	}
L334:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if base.Ui32(v718) < base.Ui32(v719) {
		v752 = v697
		goto L328
	} else {
		goto L335
	}
L335:
	;
	if base.Ui32(v719) < base.Ui32(v718) {
		v760 = v691
		v766 = v697
		goto L324
	} else {
		goto L336
	}
L336:
	;
	v722 = int32(16)
	v723 = v697 + v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v724
	v726 = int32(8)
	v727 = v697 + v726
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v728
	v730 = *(*int64)(unsafe.Add(mBase, uint32(v697)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v730
	v733 = v691 + v722
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	*(*int32)(unsafe.Add(mBase, uint32(v723))) = v734
	v737 = v691 + v726
	v738 = *(*int64)(unsafe.Add(mBase, uint32(v737)))
	*(*int64)(unsafe.Add(mBase, uint32(v727))) = v738
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v691)))
	*(*int64)(unsafe.Add(mBase, uint32(v697))) = v740
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v742
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v737))) = v744
	v746 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v691))) = v746
	v752 = v697 + int32(20)
	goto L328
L337:
	;
	goto L327
L338:
	;
	v1046 = int32(16)
	v1047 = v760 + v1046
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v1048
	v1050 = int32(8)
	v1051 = v760 + v1050
	v1052 = *(*int64)(unsafe.Add(mBase, uint32(v1051)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v1052
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v760)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v1054
	v1057 = v778 + v1046
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)))
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1058
	v1061 = v778 + v1050
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1061)))
	*(*int64)(unsafe.Add(mBase, uint32(v1051))) = v1062
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v778)))
	*(*int64)(unsafe.Add(mBase, uint32(v760))) = v1064
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v1057))) = v1066
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v1061))) = v1068
	v1070 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v778))) = v1070
	v1072 = int32(20)
	v671 = v778 - v1072
	v672 = v760 + v1072
	v678 = v766
	v682 = v789
	goto L322
L339:
	;
	v778 = v671
	v789 = v682
	goto L342
L340:
	;
	v847 = v671
	v858 = v682
	goto L341
L341:
	;
	v864 = int32(20)
	v865 = base.I32_div_s(v766-v23, v864)
	v868 = base.I32_div_s(v760-v766, v864)
	if v865 < v868 {
		goto L354
	} else {
		goto L355
	}
L342:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v794) < base.Ui32(v795) {
		goto L338
	} else {
		goto L344
	}
L343:
	;
	v847 = v843
	v858 = v841
	goto L341
L344:
	;
	if base.Ui32(v795) < base.Ui32(v794) {
		v841 = v789
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v843 = v778 - int32(20)
	if base.Ui32(v760) <= base.Ui32(v843) {
		v778 = v843
		v789 = v841
		goto L342
	} else {
		goto L353
	}
L346:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v798) < base.Ui32(v799) {
		goto L338
	} else {
		goto L347
	}
L347:
	;
	if base.Ui32(v799) < base.Ui32(v798) {
		v841 = v789
		goto L345
	} else {
		goto L348
	}
L348:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v778)+8))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v802 < v803 {
		goto L338
	} else {
		goto L349
	}
L349:
	;
	if v803 < v802 {
		v841 = v789
		goto L345
	} else {
		goto L350
	}
L350:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if base.Ui32(v806) < base.Ui32(v807) {
		goto L338
	} else {
		goto L351
	}
L351:
	;
	if base.Ui32(v807) < base.Ui32(v806) {
		v841 = v789
		goto L345
	} else {
		goto L352
	}
L352:
	;
	v810 = int32(16)
	v811 = v778 + v810
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v812
	v814 = int32(8)
	v815 = v778 + v814
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v815)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v816
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v778)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v818
	v821 = v789 + v810
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	*(*int32)(unsafe.Add(mBase, uint32(v811))) = v822
	v825 = v789 + v814
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v825)))
	*(*int64)(unsafe.Add(mBase, uint32(v815))) = v826
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v789)))
	*(*int64)(unsafe.Add(mBase, uint32(v778))) = v828
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = v830
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v825))) = v832
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v789))) = v834
	v841 = v789 - int32(20)
	goto L345
L353:
	;
	goto L343
L354:
	;
	v870 = v865
	goto L356
L355:
	;
	v870 = v868
	goto L356
L356:
	;
	if v870 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v884 = int32(0)
	goto L360
L358:
	;
	goto L359
L359:
	;
	v945 = int32(20)
	v946 = base.I32_div_s(v858-v847, v945)
	v949 = base.I32_div_s(v63-v858, v945)
	v951 = v949 - int32(1)
	if v946 < v951 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	v894 = v884 * int32(20)
	v895 = v23 + v894
	v896 = int32(16)
	v897 = v895 + v896
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v898
	v900 = int32(8)
	v901 = v895 + v900
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v901)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v902
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v895)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v904
	v906 = v894 + (v760 + v870*int32(-20))
	v908 = v906 + v896
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	*(*int32)(unsafe.Add(mBase, uint32(v897))) = v909
	v912 = v906 + v900
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v912)))
	*(*int64)(unsafe.Add(mBase, uint32(v901))) = v913
	v915 = *(*int64)(unsafe.Add(mBase, uint32(v906)))
	*(*int64)(unsafe.Add(mBase, uint32(v895))) = v915
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v908))) = v917
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v912))) = v919
	v921 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v906))) = v921
	v924 = v884 + int32(1)
	if v924 != v870 {
		v884 = v924
		goto L360
	} else {
		goto L362
	}
L361:
	;
	goto L359
L362:
	;
	goto L361
L363:
	;
	v953 = v946
	goto L365
L364:
	;
	v953 = v951
	goto L365
L365:
	;
	if v953 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v963 = int32(0)
	goto L369
L367:
	;
	goto L368
L368:
	;
	if base.Ui32(v868) <= base.Ui32(v946) {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	v977 = v963 * int32(20)
	v978 = v760 + v977
	v979 = int32(16)
	v980 = v978 + v979
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v981
	v983 = int32(8)
	v984 = v978 + v983
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v984)))
	*(*int64)(unsafe.Add(mBase, uint32(v644))) = v985
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v978)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v987
	v989 = v977 + (v63 + v953*int32(-20))
	v991 = v989 + v979
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	*(*int32)(unsafe.Add(mBase, uint32(v980))) = v992
	v995 = v989 + v983
	v996 = *(*int64)(unsafe.Add(mBase, uint32(v995)))
	*(*int64)(unsafe.Add(mBase, uint32(v984))) = v996
	v998 = *(*int64)(unsafe.Add(mBase, uint32(v989)))
	*(*int64)(unsafe.Add(mBase, uint32(v978))) = v998
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v991))) = v1000
	v1002 = *(*int64)(unsafe.Add(mBase, uint32(v644)))
	*(*int64)(unsafe.Add(mBase, uint32(v995))) = v1002
	v1004 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v989))) = v1004
	v1007 = v963 + int32(1)
	if v1007 != v953 {
		v963 = v1007
		goto L369
	} else {
		goto L371
	}
L370:
	;
	goto L368
L371:
	;
	goto L370
L372:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v868) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	goto L374
L374:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v946) {
		goto L379
	} else {
		goto L380
	}
L375:
	;
	F_sort_checkpoint_bufferids(m, v23, v868)
	mBase = m.M
	goto L377
L376:
	;
	goto L377
L377:
	;
	if base.Ui32(v946) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L378
	}
L378:
	;
	v23 = v63 + v946*int32(-20)
	v24 = v946
	goto L1
L379:
	;
	F_sort_checkpoint_bufferids(m, v63+v946*int32(-20), v946)
	mBase = m.M
	goto L381
L380:
	;
	goto L381
L381:
	;
	if base.Ui32(int32(1)) < base.Ui32(v868) {
		v44 = v868
		goto L3
	} else {
		goto L382
	}
L382:
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v9) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v100
L2:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v64) {
		goto L40
	} else {
		goto L41
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v25) {
		v100 = l1
		goto L1
	} else {
		goto L11
	}
L4:
	;
	if base.Ui32(v9) < base.Ui32(v8) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v12) < base.Ui32(v13) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(v13) < base.Ui32(v12) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 < v17 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v17 < v16 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v21) <= base.Ui32(v20) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	if base.Ui32(v25) < base.Ui32(v9) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.Ui32(v8) < base.Ui32(v25) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v28) < base.Ui32(v29) {
		v100 = l1
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v29) < base.Ui32(v28) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v32 < v33 {
		v100 = l1
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v33 < v32 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v36) < base.Ui32(v37) {
		v100 = l1
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	return l2
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v25) < base.Ui32(v8) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return l0
L23:
	;
	goto L24
L24:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v45) < base.Ui32(v46) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return l2
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(v46) < base.Ui32(v45) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return l0
L29:
	;
	goto L30
L30:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v51 < v52 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return l2
L32:
	;
	goto L33
L33:
	;
	if v52 < v51 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return l0
L35:
	;
	goto L36
L36:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v57) < base.Ui32(v58) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v60 = l2
	goto L39
L38:
	;
	v60 = l0
	goto L39
L39:
	;
	return v60
L40:
	;
	if base.Ui32(v8) < base.Ui32(v64) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32(v64) < base.Ui32(v9) {
		v100 = l1
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v67) < base.Ui32(v68) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v68) < base.Ui32(v67) {
		v100 = l1
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v71 < v72 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	if v72 < v71 {
		v100 = l1
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v76) < base.Ui32(v75) {
		v100 = l1
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	return l0
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v64) < base.Ui32(v8) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return l2
L52:
	;
	goto L53
L53:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v84) < base.Ui32(v85) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return l0
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v85) < base.Ui32(v84) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return l2
L58:
	;
	goto L59
L59:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v90 < v91 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	return l0
L61:
	;
	goto L62
L62:
	;
	if v91 < v90 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return l2
L64:
	;
	goto L65
L65:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v96) < base.Ui32(v97) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v99 = l0
	goto L68
L67:
	;
	v99 = l2
	goto L68
L68:
	;
	v100 = v99
	goto L1
}
func F_sort_desc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_copy(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		if v14 != 0 {
			v15 = F_array_contains_nulls(m, v10)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_sort_desc_0), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_sort_desc_1), int32(244), int32(_a_F_sort_desc_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v20 = F_ArrayGetNItems(m, v17, v10+int32(16))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v22)
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						if v24 != 0 {
							v32 = v24
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v32+v10, v20, v7+int32(15))
						mBase = m.M
						m.G0 = v7 + int32(16)
						return v10
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v20 = F_ArrayGetNItems(m, v17, v10+int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v22)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v24 != 0 {
					v32 = v24
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v32+v10, v20, v7+int32(15))
				mBase = m.M
				m.G0 = v7 + int32(16)
				return v10
			}
		}
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
	var v31 int32
	_ = v31
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
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
					v31 = int32(0)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v16+v31<<(uint(int32(2))%32)))) = v29
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						if v44 != 0 {
							v29 = v44
							v31 = v31 + int32(1)
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
						v138 = int32(1)
					} else {
						v67 = int32(1)
						v68 = int32(3)
						if v10 <= v68 {
							v71 = v68
						} else {
							v71 = v10
						}
						if int32(4) <= v10 {
							v82 = int32(0)
							v83 = v67
							for {
								v90 = int32(2)
								v92 = v16 + v83<<(uint(v90)%32)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
								v94 = int32(4)
								v95 = v92 + v94
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v96
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v92-v94)))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v100
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
								v104 = v83 + v90
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v16+v104<<(uint(v90)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v102)+24)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
								*(*int32)(unsafe.Add(mBase, uint32(v102)+28)) = v110
								if v82 != v71&int32(2147483646)-int32(4) {
									v82 = v82 + v90
									v83 = v104
									continue
								} else {
									break
								}
								break
							}
							v117 = v104
						} else {
							v117 = v67
						}
						v125 = v71 - int32(1)
						if v71&int32(1) == int32(0) {
							v138 = v125
						} else {
							v130 = v16 + v117<<(uint(int32(2))%32)
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+24)) = v132
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v130-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v131)+28)) = v136
							v138 = v125
						}
					}
					v149 = v16 + v138<<(uint(int32(2))%32)
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
					*(*int32)(unsafe.Add(mBase, uint32(v150)+24)) = int32(0)
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v149-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v150)+28)) = v155
					F_pfree(m, v16)
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
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
	var v14 int32
	_ = v14
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
	var v44 int32
	_ = v44
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
	v13 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v14 <= v13 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v13
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
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v44 <= v43 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = v43
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
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
													*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109-int32(-64)))) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+80)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = v110
													*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v110
													v128 = int32(8)
													v129 = v98 + v128
													v131 = v104 + v128
													if v131 != v86&int32(-8) {
														v98 = v129
														v104 = v131
														continue
													} else {
														break
													}
													break
												}
												v134 = v129
											} else {
												v134 = v85
											}
											if v88 == int32(0) {
											} else {
												v147 = v134
												v149 = int32(0)
												for {
													*(*int32)(unsafe.Add(mBase, uint32(v84+v147<<(uint(int32(4))%32)))) = int32(-1)
													v161 = int32(1)
													v164 = v149 + v161
													if v164 != v88 {
														v147 = v147 + v161
														v149 = v164
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
										v178 = m.ExcPending
										if v178 != 0 {
											return int32(0)
										} else {
											v184 = v58
											*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v184
											*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v184
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											if int32(0) < v191 {
												v196 = F_palloc(m, v191<<(uint(int32(2))%32))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v196
													v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v202 = F_palloc(m, v199<<(uint(int32(2))%32))
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v202
														v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v208 = F_palloc(m, v205<<(uint(int32(3))%32))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v208
															v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															v214 = F_palloc(m, v211<<(uint(int32(3))%32))
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v214
																v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																if int32(0) < v217 {
																	v222 = int32(0)
																	for {
																		v232 = v222 << (uint(int32(3)) % 32)
																		v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																		*(*int64)(unsafe.Add(mBase, uint32(v232+v233))) = int64(0)
																		v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																		*(*int64)(unsafe.Add(mBase, uint32(v237+v232))) = int64(9218868437227405312)
																		v242 = v222 + int32(1)
																		v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																		if v242 < v243 {
																			v222 = v242
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v248 = v243
																} else {
																	v248 = v217
																}
																v257 = F_palloc0(m, v248<<(uint(int32(2))%32))
																mBase = m.M
																v258 = m.ExcPending
																if v258 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v257
																	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	v261 = F_palloc(m, v260)
																	mBase = m.M
																	v262 = m.ExcPending
																	if v262 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v261
																		v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																		v267 = F__emscripten_memset_bulkmem(m, v261, base.I32_extend8_s(int32(1)), v265)
																		mBase = m.M
																		v279 = v16 + int32(132)
																		v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																		mBase = m.M
																		v283 = m.ExcPending
																		if v283 != 0 {
																			return int32(0)
																		} else {
																			v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																			v288 = v16 + int32(148)
																			v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
																			v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
																			*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
																			v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
																			*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
																			v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
																			*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
																			*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
																			*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
																			v301 = v16 + int32(160)
																			v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																			mBase = m.M
																			v305 = m.ExcPending
																			if v305 != 0 {
																				return int32(0)
																			} else {
																				v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																				v310 = v16 + int32(176)
																				v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
																				*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
																				v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
																				*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
																				v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
																				*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
																				v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
																				*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
																				*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
																				*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
																				v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																				v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
											} else {
												v279 = v16 + int32(132)
												v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
												mBase = m.M
												v283 = m.ExcPending
												if v283 != 0 {
													return int32(0)
												} else {
													v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
													v288 = v16 + int32(148)
													v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
													v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
													*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
													v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
													v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
													*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
													*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
													v301 = v16 + int32(160)
													v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
													mBase = m.M
													v305 = m.ExcPending
													if v305 != 0 {
														return int32(0)
													} else {
														v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
														v310 = v16 + int32(176)
														v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
														v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
														*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
														v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
														v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
														*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
														*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
														v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
														v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
														*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
														*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
														return v11
													}
												}
											}
										}
									}
								} else {
									v184 = v51
									*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v184
									*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v184
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									if int32(0) < v191 {
										v196 = F_palloc(m, v191<<(uint(int32(2))%32))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v196
											v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v202 = F_palloc(m, v199<<(uint(int32(2))%32))
											mBase = m.M
											v203 = m.ExcPending
											if v203 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v202
												v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v208 = F_palloc(m, v205<<(uint(int32(3))%32))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v208
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v214 = F_palloc(m, v211<<(uint(int32(3))%32))
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v214
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														if int32(0) < v217 {
															v222 = int32(0)
															for {
																v232 = v222 << (uint(int32(3)) % 32)
																v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																*(*int64)(unsafe.Add(mBase, uint32(v232+v233))) = int64(0)
																v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																*(*int64)(unsafe.Add(mBase, uint32(v237+v232))) = int64(9218868437227405312)
																v242 = v222 + int32(1)
																v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																if v242 < v243 {
																	v222 = v242
																	continue
																} else {
																	break
																}
																break
															}
															v248 = v243
														} else {
															v248 = v217
														}
														v257 = F_palloc0(m, v248<<(uint(int32(2))%32))
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v257
															v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															v261 = F_palloc(m, v260)
															mBase = m.M
															v262 = m.ExcPending
															if v262 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v261
																v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																v267 = F__emscripten_memset_bulkmem(m, v261, base.I32_extend8_s(int32(1)), v265)
																mBase = m.M
																v279 = v16 + int32(132)
																v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																mBase = m.M
																v283 = m.ExcPending
																if v283 != 0 {
																	return int32(0)
																} else {
																	v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																	v288 = v16 + int32(148)
																	v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
																	v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
																	*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
																	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
																	v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
																	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
																	*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
																	v301 = v16 + int32(160)
																	v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																	mBase = m.M
																	v305 = m.ExcPending
																	if v305 != 0 {
																		return int32(0)
																	} else {
																		v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																		v310 = v16 + int32(176)
																		v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
																		*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
																		v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
																		*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
																		v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
																		*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
																		v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
																		*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
																		*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
																		*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
																		v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																		v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
																		*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
									} else {
										v279 = v16 + int32(132)
										v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
										mBase = m.M
										v283 = m.ExcPending
										if v283 != 0 {
											return int32(0)
										} else {
											v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
											v288 = v16 + int32(148)
											v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
											v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
											*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
											v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
											v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
											*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
											*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
											v301 = v16 + int32(160)
											v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
											mBase = m.M
											v305 = m.ExcPending
											if v305 != 0 {
												return int32(0)
											} else {
												v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
												v310 = v16 + int32(176)
												v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
												v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
												*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
												v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
												v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
												*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
												*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
												v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
												v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
												*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
												*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+48)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109-int32(-64)))) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+80)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+96)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v109)+112)) = v110
												v128 = int32(8)
												v129 = v98 + v128
												v131 = v104 + v128
												if v131 != v86&int32(-8) {
													v98 = v129
													v104 = v131
													continue
												} else {
													break
												}
												break
											}
											v134 = v129
										} else {
											v134 = v85
										}
										if v88 == int32(0) {
										} else {
											v147 = v134
											v149 = int32(0)
											for {
												*(*int32)(unsafe.Add(mBase, uint32(v84+v147<<(uint(int32(4))%32)))) = int32(-1)
												v161 = int32(1)
												v164 = v149 + v161
												if v164 != v88 {
													v147 = v147 + v161
													v149 = v164
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
									v178 = m.ExcPending
									if v178 != 0 {
										return int32(0)
									} else {
										v184 = v58
										*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v184
										*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v184
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										if int32(0) < v191 {
											v196 = F_palloc(m, v191<<(uint(int32(2))%32))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v196
												v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v202 = F_palloc(m, v199<<(uint(int32(2))%32))
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v202
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v208 = F_palloc(m, v205<<(uint(int32(3))%32))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v208
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v214 = F_palloc(m, v211<<(uint(int32(3))%32))
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v214
															v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															if int32(0) < v217 {
																v222 = int32(0)
																for {
																	v232 = v222 << (uint(int32(3)) % 32)
																	v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
																	*(*int64)(unsafe.Add(mBase, uint32(v232+v233))) = int64(0)
																	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
																	*(*int64)(unsafe.Add(mBase, uint32(v237+v232))) = int64(9218868437227405312)
																	v242 = v222 + int32(1)
																	v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	if v242 < v243 {
																		v222 = v242
																		continue
																	} else {
																		break
																	}
																	break
																}
																v248 = v243
															} else {
																v248 = v217
															}
															v257 = F_palloc0(m, v248<<(uint(int32(2))%32))
															mBase = m.M
															v258 = m.ExcPending
															if v258 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v257
																v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																v261 = F_palloc(m, v260)
																mBase = m.M
																v262 = m.ExcPending
																if v262 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v261
																	v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																	v267 = F__emscripten_memset_bulkmem(m, v261, base.I32_extend8_s(int32(1)), v265)
																	mBase = m.M
																	v279 = v16 + int32(132)
																	v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
																	mBase = m.M
																	v283 = m.ExcPending
																	if v283 != 0 {
																		return int32(0)
																	} else {
																		v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																		v288 = v16 + int32(148)
																		v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
																		*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
																		v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
																		*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
																		v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
																		*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
																		v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
																		*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
																		*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
																		*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
																		v301 = v16 + int32(160)
																		v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																		mBase = m.M
																		v305 = m.ExcPending
																		if v305 != 0 {
																			return int32(0)
																		} else {
																			v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																			v310 = v16 + int32(176)
																			v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
																			*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
																			v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
																			*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
																			v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
																			*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
																			v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
																			*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
																			*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
																			*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
																			v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																			v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
																			*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
										} else {
											v279 = v16 + int32(132)
											v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
											mBase = m.M
											v283 = m.ExcPending
											if v283 != 0 {
												return int32(0)
											} else {
												v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
												v288 = v16 + int32(148)
												v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
												v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
												*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
												v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
												v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
												*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
												*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
												v301 = v16 + int32(160)
												v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
												mBase = m.M
												v305 = m.ExcPending
												if v305 != 0 {
													return int32(0)
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
													v310 = v16 + int32(176)
													v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
													v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
													*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
													v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
													v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
													*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
													*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
													v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
													v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
													*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
													*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v16
													return v11
												}
											}
										}
									}
								}
							} else {
								v184 = v51
								*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v184
								*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v184
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
								if int32(0) < v191 {
									v196 = F_palloc(m, v191<<(uint(int32(2))%32))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v196
										v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										v202 = F_palloc(m, v199<<(uint(int32(2))%32))
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v202
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v208 = F_palloc(m, v205<<(uint(int32(3))%32))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+188)) = v208
												v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v214 = F_palloc(m, v211<<(uint(int32(3))%32))
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v214
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													if int32(0) < v217 {
														v222 = int32(0)
														for {
															v232 = v222 << (uint(int32(3)) % 32)
															v233 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
															*(*int64)(unsafe.Add(mBase, uint32(v232+v233))) = int64(0)
															v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
															*(*int64)(unsafe.Add(mBase, uint32(v237+v232))) = int64(9218868437227405312)
															v242 = v222 + int32(1)
															v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															if v242 < v243 {
																v222 = v242
																continue
															} else {
																break
															}
															break
														}
														v248 = v243
													} else {
														v248 = v217
													}
													v257 = F_palloc0(m, v248<<(uint(int32(2))%32))
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v257
														v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v261 = F_palloc(m, v260)
														mBase = m.M
														v262 = m.ExcPending
														if v262 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v261
															v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
															v267 = F__emscripten_memset_bulkmem(m, v261, base.I32_extend8_s(int32(1)), v265)
															mBase = m.M
															v279 = v16 + int32(132)
															v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
															mBase = m.M
															v283 = m.ExcPending
															if v283 != 0 {
																return int32(0)
															} else {
																v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																v288 = v16 + int32(148)
																v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
																*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
																v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
																*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
																v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
																v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
																*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
																*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
																*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
																v301 = v16 + int32(160)
																v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
																mBase = m.M
																v305 = m.ExcPending
																if v305 != 0 {
																	return int32(0)
																} else {
																	v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
																	v310 = v16 + int32(176)
																	v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
																	v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
																	*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
																	v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
																	*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
																	v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
																	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
																	*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
																	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
																	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
								} else {
									v279 = v16 + int32(132)
									v282 = F_index_getprocinfo(m, l0, int32(1), int32(4))
									mBase = m.M
									v283 = m.ExcPending
									if v283 != 0 {
										return int32(0)
									} else {
										v285 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
										v288 = v16 + int32(148)
										v289 = *(*int64)(unsafe.Add(mBase, uint32(v282)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v288))) = v289
										v291 = *(*int64)(unsafe.Add(mBase, uint32(v282)))
										*(*int64)(unsafe.Add(mBase, uint32(v279))) = v291
										v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v293
										v295 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v295
										*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v285
										*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(0)
										v301 = v16 + int32(160)
										v304 = F_index_getprocinfo(m, l0, int32(1), int32(5))
										mBase = m.M
										v305 = m.ExcPending
										if v305 != 0 {
											return int32(0)
										} else {
											v307 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
											v310 = v16 + int32(176)
											v311 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v310))) = v311
											v313 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
											*(*int64)(unsafe.Add(mBase, uint32(v301))) = v313
											v315 = *(*int32)(unsafe.Add(mBase, uint32(v304)+24))
											*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v315
											v317 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v317
											*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v307
											*(*int32)(unsafe.Add(mBase, uint32(v310))) = int32(0)
											v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
											v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
											*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v323
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 float64
	_ = v266
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 float64
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v789 int32
	_ = v789
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
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
	v65 = F_read_stream_begin_relation(m, int32(13), v60, v26, int32(120), v21+int32(40), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v67 = l0
	v70 = v21
	v76 = v26
	v81 = l0 + int32(100)
	v82 = v65
	v83 = v53
	goto L8
L8:
	;
	if v83 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	if base.Ui32(v100) < base.Ui32(v99) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v88 = F_RelationGetNumberOfBlocksInFork(m, v76, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LockRelationForExtension(m, v76, int32(7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v99 = v88
	goto L10
L15:
	;
	v94 = F_RelationGetNumberOfBlocksInFork(m, v76, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_UnlockRelationForExtension(m, v76, int32(7))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v99 = v94
	goto L10
L18:
	;
	F_read_stream_reset(m, v118)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L190
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+44)) = v99
	v103 = v67
	v106 = v70
	v112 = v76
	v117 = v81
	v118 = v82
	v119 = v83
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_read_stream_end(m, v82)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L184
	}
L22:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L181
	}
L24:
	;
	v125 = F_read_stream_next_buffer(m, v118, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v125 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v125 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBuffer(m, v125, int32(2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[1]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134+(v125^int32(-1))<<(uint(int32(6))%32))+16))
	v149 = v140
	goto L27
L29:
	;
	goto L30
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[2]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+v125<<(uint(int32(6))%32)+int32(-64))+16))
	v149 = v148
	goto L27
L31:
	;
	v153 = int32(0)
	v154 = base.B2i32(v153 <= v125)
	if v154 == v153 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v125)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+14)))
	if v421 != 0 {
		goto L86
	} else {
		goto L87
	}
L34:
	;
	if base.Ui32(v149-int32(1)) < base.Ui32(int32(2)) {
		goto L32
	} else {
		goto L84
	}
L35:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+14)))
	if v173 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158+(v125^int32(-1))<<(uint(int32(2))%32))))
	v172 = v164
	goto L35
L37:
	;
	goto L38
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v172 = v166 + v125<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+12)))
	if base.Ui32(v176) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+16)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+v179))))
	if v181&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v184 = int32(1)
	if base.Ui32(v149-v184) <= base.Ui32(v184) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	F_vacuumRedirectAndPlaceholder(m, v130, v378, v125)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L83
	}
L44:
	;
	if v154 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_vacuumLeafPage(m, v103, v130, v125, int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L81
	}
L47:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+12)))
	v207 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+868)) = uint16(v207)
	if base.Ui32(v206) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L51
	}
L48:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191+(v125^int32(-1))<<(uint(int32(2))%32))))
	v205 = v197
	goto L47
L49:
	;
	goto L50
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v205 = v199 + v125<<(uint(int32(13))%32) + int32(-8192)
	goto L47
L51:
	;
	v216 = int32(base.Ui32(v206+int32(_a_F_spgvacuumscan_0))>>(uint(int32(2))%32)) & int32(_a_F_spgvacuumscan_1)
	if v216 == int32(0) {
		goto L34
	} else {
		goto L52
	}
L52:
	;
	v225 = int32(1)
	v228 = int32(0)
	goto L53
L53:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v225&int32(_a_F_spgvacuumscan_1)<<(uint(int32(2))%32)+(v205+int32(24))-int32(4))))
	v251 = v205 + v248&int32(_a_F_spgvacuumscan_2)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v252&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v310 = v303 & int32(_a_F_spgvacuumscan_1)
	if v310 == int32(0) {
		goto L34
	} else {
		goto L66
	}
L55:
	;
	v305 = v225 + int32(1)
	if base.Ui32(v305&int32(_a_F_spgvacuumscan_1)) <= base.Ui32(v216) {
		v225 = v305
		v228 = v303
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v263)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v263)+8)) = base.F64_add(v299, float64(1))
	v303 = v228
	goto L55
L57:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v261 = m.T0[v260].(func(*base.Module, int32, int32) int32)(m, v251+int32(6), v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v261 == int32(0) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v266 = *(*float64)(unsafe.Add(mBase, uint32(v263)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v263)+16)) = base.F64_add(v266, float64(1))
	v274 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v106+int32(48)+v228&int32(_a_F_spgvacuumscan_1)<<(uint(v274)%32)))) = uint16(v225)
	v279 = v228 + v274
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+868)) = uint16(v279)
	v303 = v279
	goto L55
L62:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v285 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v106+int32(32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(445), int32(_a_F_spgvacuumscan_5))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
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
	v313 = int32(_a_F_spgvacuumscan_6)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v315 + int32(1)
	F_PageIndexMultiDelete(m, v205, v106+int32(48), v310)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_MarkBufferDirty(m, v125)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+118)))
	if v326 != int32(112) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v364 = int32(_a_F_spgvacuumscan_6)
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v366 - int32(1)
	goto L34
L70:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[6]))
	if v330 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v130)+32))
	if v333 != 0 {
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	if v334 != 0 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v103)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+872)) = v337
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+876)) = uint8(v339)
	F_XLogRegisterData(m, v106+int32(868), int32(12))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+868)))
	F_XLogRegisterData(m, v106+int32(48), v348<<(uint(int32(1))%32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_XLogRegisterBuffer(m, int32(0), v125, int32(8))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v359 = F_XLogInsert(m, int32(16), int32(112))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = base.I64_rotr(v359, int64(32))
	goto L69
L81:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	F_vacuumRedirectAndPlaceholder(m, v130, v374, v125)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
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
	F_SpGistSetLastUsedPage(m, v130, v125)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v422) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_RecordFreeIndexPage(m, v130, v149)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v427)+28)) = v428 + int32(1)
	goto L32
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+108)) = v149
	goto L32
L92:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v455 == int32(0) {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	v460 = v103
	v463 = v106
	v465 = v455
	v469 = v112
	v472 = v459
	v474 = v117
	v475 = v118
	v476 = v119
	goto L95
L94:
	;
	goto L23
L95:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)))
	if v478 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	if v910 != 0 {
		goto L174
	} else {
		goto L175
	}
L97:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	v891 = v460
	v894 = v463
	v900 = v469
	v903 = v472
	v905 = v474
	v906 = v475
	v907 = v476
	goto L99
L99:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v909 != 0 {
		v460 = v891
		v463 = v894
		v465 = v909
		v469 = v900
		v472 = v903
		v474 = v905
		v475 = v906
		v476 = v907
		goto L95
	} else {
		goto L173
	}
L100:
	;
	v484 = int32(0)
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+2)))
	v486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465))))
	v489 = v485 | v486<<(uint(int32(16))%32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+24))
	v493 = F_ReadBufferExtended(m, v472, v484, v489, v484, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_LockBuffer(m, v493, int32(2))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v493 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	F_UnlockReleaseBuffer(m, v493)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L172
	}
L104:
	;
	v516 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+14)))
	if v516 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L105:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v501+(v493^int32(-1))<<(uint(int32(2))%32))))
	v515 = v507
	goto L104
L106:
	;
	goto L107
L107:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v515 = v509 + v493<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+16)))
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515+v519))))
	if v521&int32(2) != 0 {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	if v521&int32(4) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v537 = v465
	goto L113
L111:
	;
	goto L112
L112:
	;
	v826 = int32(1)
	if base.Ui32(v489-v826) <= base.Ui32(v826) {
		goto L94
	} else {
		goto L161
	}
L113:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+6)))
	if v548 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L158
	}
L115:
	;
	goto L114
L116:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	if v809 != 0 {
		v537 = v809
		goto L113
	} else {
		goto L157
	}
L117:
	;
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v537)+2)))
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v537))))
	if v549|v550<<(uint(int32(16))%32) != v489 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v537)+4)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v555<<(uint(int32(2))%32)+(v515+int32(24))-int32(4))))
	v564 = v515 + v561&int32(_a_F_spgvacuumscan_2)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	switch v565 & int32(3) {
	case 0:
		goto L121
	case 1:
		goto L120
	default:
		goto L115
	}
L119:
	;
	v789 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v537)+6)) = uint8(v789)
	goto L116
L120:
	;
	v701 = v564 + int32(6)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	if v702 != 0 {
		goto L143
	} else {
		goto L144
	}
L121:
	;
	if v565&int32(_a_F_spgvacuumscan_7) == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v579 = v564 + int32(base.Ui32(v565)>>(uint(int32(16))%32)) + int32(8)
	v586 = int32(0)
	goto L123
L123:
	;
	if v579 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L119
L125:
	;
	v688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+6)))
	v689 = int32(_a_F_spgvacuumscan_8)
	v693 = v586 + int32(1)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	if base.Ui32(v693) < base.Ui32(int32(base.Ui32(v694)>>(uint(int32(3))%32))&v689) {
		v579 = v579 + v688&v689
		v586 = v693
		goto L123
	} else {
		goto L142
	}
L126:
	;
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)))
	if v598 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	if v601 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v604 = v601
	goto L131
L129:
	;
	v644 = v474
	goto L130
L130:
	;
	v659 = F_palloc(m, int32(12))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L141
	}
L131:
	;
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+2)))
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579))))
	v622 = int32(16)
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+2)))
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604))))
	if v620|v621<<(uint(v622)%32) == v625|v626<<(uint(v622)%32) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v644 = v604 + int32(8)
	goto L130
L133:
	;
	if v636 != 0 {
		goto L125
	} else {
		goto L139
	}
L134:
	;
	goto L133
L135:
	;
	v632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)))
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	if v632 == v633 {
		v636 = int32(1)
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v636 = int32(0)
	goto L134
L138:
	;
	goto L137
L139:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v604)+8))
	if v637 != 0 {
		v604 = v637
		goto L131
	} else {
		goto L140
	}
L140:
	;
	goto L132
L141:
	;
	v661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v659)+4)) = uint16(v661)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v659))) = v663
	v665 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v659)+8)) = v665
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+6)) = uint8(v665)
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v659
	goto L125
L142:
	;
	goto L124
L143:
	;
	v704 = v702
	goto L146
L144:
	;
	v743 = v474
	goto L145
L145:
	;
	v760 = F_palloc(m, int32(12))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L156
	}
L146:
	;
	v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+2)))
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701))))
	v723 = int32(16)
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+2)))
	v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704))))
	if v721|v722<<(uint(v723)%32) == v726|v727<<(uint(v723)%32) {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v743 = v704 + int32(8)
	goto L145
L148:
	;
	if v737 != 0 {
		goto L119
	} else {
		goto L154
	}
L149:
	;
	goto L148
L150:
	;
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+4)))
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+4)))
	if v733 == v734 {
		v737 = int32(1)
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v737 = int32(0)
	goto L149
L153:
	;
	goto L152
L154:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v704)+8))
	if v738 != 0 {
		v704 = v738
		goto L146
	} else {
		goto L155
	}
L155:
	;
	goto L147
L156:
	;
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v760)+4)) = uint16(v762)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	*(*int32)(unsafe.Add(mBase, uint32(v760))) = v764
	v766 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = v766
	*(*uint8)(unsafe.Add(mBase, uint32(v760)+6)) = uint8(v766)
	*(*int32)(unsafe.Add(mBase, uint32(v743))) = v760
	goto L119
L157:
	;
	goto L103
L158:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v814 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v463)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(783), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_vacuumLeafPage(m, v460, v472, v493, int32(1))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)+4))
	F_vacuumRedirectAndPlaceholder(m, v472, v834, v493)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_SpGistSetLastUsedPage(m, v472, v493)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v839 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)) = uint8(v839)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v841 == int32(0) {
		goto L103
	} else {
		goto L165
	}
L165:
	;
	v845 = v841
	goto L166
L166:
	;
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v845)+2)))
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v845))))
	if v489 == v862|v863<<(uint(int32(16))%32) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L103
L168:
	;
	v868 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v845)+6)) = uint8(v868)
	goto L170
L169:
	;
	goto L170
L170:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v845)+8))
	if v870 != 0 {
		v845 = v870
		goto L166
	} else {
		goto L171
	}
L171:
	;
	goto L167
L172:
	;
	v891 = v460
	v894 = v463
	v900 = v469
	v903 = v472
	v905 = v474
	v906 = v475
	v907 = v476
	goto L99
L173:
	;
	goto L96
L174:
	;
	v912 = v910
	goto L177
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v905))) = int32(0)
	v103 = v891
	v106 = v894
	v112 = v900
	v117 = v905
	v118 = v906
	v119 = v907
	goto L22
L177:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v912)+8))
	F_pfree(m, v912)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	goto L176
L179:
	;
	if v929 != 0 {
		v912 = v929
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v472)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+16)) = v956 + int32(4)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_10), v463+int32(16))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(722), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_SpGistUpdateMetaPage(m, v76)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)+28))
	if v975 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_FreeSpaceMapVacuum(m, v76)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	v979 = v974
	goto L188
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979))) = v99
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v981)+24)) = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v984)+32)) = v985
	m.G0 = v70 + int32(880)
	return
L189:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v979 = v978
	goto L188
L190:
	;
	v67 = v103
	v70 = v106
	v76 = v112
	v81 = v117
	v82 = v118
	v83 = v119
	goto L8
}
func F_spgvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v244 int32
	_ = v244
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
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
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v750 int64
	_ = v750
	var v753 int64
	_ = v753
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int64
	_ = v789
	var v795 int32
	_ = v795
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v846 int32
	_ = v846
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(256)
	m.G0 = v27
	v30 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if int32(0) < v471 {
		goto L87
	} else {
		goto L88
	}
L2:
	;
	return int32(0)
L3:
	;
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = v34 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	v40 = F_get_opfamily_name(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v437 = m.ExcPending
	if v437 != 0 {
		goto L2
	} else {
		goto L84
	}
L7:
	;
	v44 = int32(0)
	v46 = F_SearchSysCacheList(m, int32(4), int32(1), v39, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v50 = int32(0)
	v52 = F_SearchSysCacheList(m, int32(5), int32(1), v39, v50, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v54 = F_identify_opfamily_groups(m, v46, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v56 = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	if v57 <= int32(0) {
		v450 = v56
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v63 = v27 + int32(240)
	v67 = v56
	v68 = v2
	v72 = v2
	v82 = v2
	v83 = v2
	goto L12
L12:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(48)+v68<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+56))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v94 = v92 + v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v95 == v96 {
		v125 = v67
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v450 = v409
	goto L1
L14:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+16)))
	switch v126 - int32(1) {
	case 0:
		goto L31
	case 1, 2, 3:
		goto L30
	case 4:
		goto L29
	case 5:
		goto L28
	case 6:
		goto L27
	default:
		goto L26
	}
L15:
	;
	v98 = int32(0)
	v101 = F_errstart(m, int32(17), v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v101 == int32(0) {
		v125 = v98
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v109 = F_format_procedure(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v27)+212)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_1), v27+int32(208))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(96), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v125 = v98
	goto L14
L22:
	;
	v431 = v68 + int32(1)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	if v431 < v432 {
		v67 = v409
		v68 = v431
		v72 = v414
		v82 = v424
		v83 = v425
		goto L12
	} else {
		goto L83
	}
L23:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L2
	} else {
		goto L79
	}
L24:
	;
	v351 = int32(0)
	v354 = F_errstart(m, int32(17), v351)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L2
	} else {
		goto L77
	}
L25:
	;
	if v138 != 0 {
		v409 = v188
		v414 = v189
		v424 = v154
		v425 = v156
		goto L22
	} else {
		goto L76
	}
L26:
	;
	v318 = int32(0)
	v321 = F_errstart(m, int32(17), v318)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L2
	} else {
		goto L74
	}
L27:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v314 = F_check_amoptsproc_signature(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L72
	}
L28:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v82 != v298 {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L68
	}
L29:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+176)) = int64(9796820404457)
	v290 = int32(2)
	v294 = F_check_amproc_signature(m, v285, int32(16), int32(1), v290, v290, v27+int32(176))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L66
	}
L30:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+160)) = int64(9796820404457)
	v277 = int32(2)
	v281 = F_check_amproc_signature(m, v272, int32(2278), int32(1), v277, v277, v27+int32(160))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L64
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+144)) = int64(9796820404457)
	v134 = int32(2)
	v138 = F_check_amproc_signature(m, v129, int32(2278), int32(1), v134, v134, v27+int32(144))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v141 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v27)+252)) = v140
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v152 = F_OidFunctionCall2Coll(m, v146, int32(0), v27+int32(252), v27+int32(232))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v37 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v155 = v37
	goto L36
L35:
	;
	v155 = v154
	goto L36
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v157 == int32(0) {
		v188 = v125
		v189 = v155
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v27)+252))
	if v189 != v190 {
		goto L25
	} else {
		goto L49
	}
L38:
	;
	if v157 == v155 {
		v188 = v125
		v189 = v155
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v163 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v163 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v27)+240))
	v188 = int32(0)
	v189 = v187
	goto L37
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v27)+240))
	v169 = F_format_type_be(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v171 = F_format_type_be(m, v155)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+132)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v169
	F_errmsg(m, int32(_a_F_spgvalidate_4), v27+int32(128))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(130), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	if v54 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v138 == int32(0) {
		v335 = v189
		v345 = v154
		v346 = v156
		goto L24
	} else {
		goto L63
	}
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v194 <= int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v197 = int32(0)
	if v197 < v194 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v200 = v194
	goto L55
L54:
	;
	v200 = v197
	goto L55
L55:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v206 = int32(0)
	goto L56
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(int32(2))%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v232 != v201 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L50
L58:
	;
	v244 = v206 + int32(1)
	if v244 != v200 {
		v206 = v244
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v234 != v235 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v231)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+16)) = v237 | int64(64)
	if v138 == int32(0) {
		v335 = v189
		v345 = v154
		v346 = v156
		goto L24
	} else {
		goto L61
	}
L61:
	;
	v409 = v188
	v414 = v189
	v424 = v154
	v425 = v156
	goto L22
L62:
	;
	goto L57
L63:
	;
	v409 = v188
	v414 = v189
	v424 = v154
	v425 = v156
	goto L22
L64:
	;
	if v281 == int32(0) {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L65
	}
L65:
	;
	v409 = v125
	v414 = v72
	v424 = v82
	v425 = v83
	goto L22
L66:
	;
	if v294 == int32(0) {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L67
	}
L67:
	;
	v409 = v125
	v414 = v72
	v424 = v82
	v425 = v83
	goto L22
L68:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v83 != v300 {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L69
	}
L69:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v82
	v304 = int32(1)
	v309 = F_check_amproc_signature(m, v302, v72, v304, v304, v304, v27+int32(192))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	if v309 == int32(0) {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L71
	}
L71:
	;
	v409 = v125
	v414 = v72
	v424 = v82
	v425 = v83
	goto L22
L72:
	;
	if v314 == int32(0) {
		v335 = v72
		v345 = v82
		v346 = v83
		goto L24
	} else {
		goto L73
	}
L73:
	;
	v409 = v125
	v414 = v72
	v424 = v82
	v425 = v83
	goto L22
L74:
	;
	if v321 == int32(0) {
		v409 = v318
		v414 = v72
		v424 = v82
		v425 = v83
		goto L22
	} else {
		goto L75
	}
L75:
	;
	v363 = int32(184)
	v368 = v72
	v378 = v82
	v379 = v83
	v384 = int32(_a_F_spgvalidate_5)
	goto L23
L76:
	;
	v335 = v189
	v345 = v154
	v346 = v156
	goto L24
L77:
	;
	if v354 == int32(0) {
		v409 = v351
		v414 = v335
		v424 = v345
		v425 = v346
		goto L22
	} else {
		goto L78
	}
L78:
	;
	v363 = int32(196)
	v368 = v335
	v378 = v345
	v379 = v346
	v384 = int32(_a_F_spgvalidate_6)
	goto L23
L79:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v389 = F_format_procedure(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+124)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v27)+116)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v40
	F_errmsg(m, v384, v27+int32(112))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), v363, int32(_a_F_spgvalidate_3))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v409 = int32(0)
	v414 = v368
	v424 = v378
	v425 = v379
	goto L22
L83:
	;
	goto L13
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
	F_errmsg_internal(m, int32(_a_F_spgvalidate_7), v27)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(63), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v480 = v450
	v481 = int32(0)
	goto L90
L88:
	;
	v625 = v450
	goto L89
L89:
	;
	if v54 != 0 {
		goto L122
	} else {
		goto L123
	}
L90:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v481<<(uint(int32(2))%32))))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+22)))
	v507 = v505 + v506
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v507)+16)))
	if base.Ui32(int32(_a_F_spgvalidate_8)) < base.Ui32((v508+int32(-64))&int32(_a_F_spgvalidate_9)) {
		v545 = v480
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v625 = v617
	goto L89
L92:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+18)))
	if v547 == int32(115) {
		v583 = int32(16)
		v584 = v545
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v515 = int32(0)
	v518 = F_errstart(m, int32(17), v515)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	if v518 == int32(0) {
		v545 = v515
		goto L92
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v526 = F_format_operator(m, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v528 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_10), v27+int32(96))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(216), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v545 = v515
	goto L92
L100:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v507)+12))
	v588 = F_check_amop_signature(m, v585, v583, v586, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L112
	}
L101:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v551 = F_get_op_rettype(m, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v507)+28))
	v554 = F_opfamily_can_sort_type(m, v553, v551)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	if v554 != 0 {
		v583 = v551
		v584 = v545
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v556 = int32(0)
	v559 = F_errstart(m, int32(17), v556)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	if v559 == int32(0) {
		v583 = v551
		v584 = v556
		goto L100
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v567 = F_format_operator(m, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_11), v27+int32(80))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(231), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v583 = v551
	v584 = v556
	goto L100
L111:
	;
	v619 = v481 + int32(1)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v619 < v620 {
		v480 = v617
		v481 = v619
		goto L90
	} else {
		goto L120
	}
L112:
	;
	if v588 != 0 {
		v617 = v584
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v590 = int32(0)
	v593 = F_errstart(m, int32(17), v590)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	if v593 == int32(0) {
		v617 = v590
		goto L111
	} else {
		goto L115
	}
L115:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	v601 = F_format_operator(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v601
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_12), v27-int32(-64))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(247), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	v617 = v590
	goto L111
L120:
	;
	goto L91
L121:
	;
	F_ReleaseCatCacheList(m, v52)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L2
	} else {
		goto L167
	}
L122:
	;
	v646 = int32(0)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v647 <= v646 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	v873 = int32(0)
	v876 = F_errstart(m, int32(17), v873)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L2
	} else {
		goto L162
	}
L125:
	;
	v825 = v625
	v846 = int32(1)
	goto L127
L126:
	;
	v654 = int32(0)
	v655 = v625
	v656 = v646
	goto L128
L127:
	;
	if v846 == int32(0) {
		v901 = v825
		goto L121
	} else {
		goto L161
	}
L128:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676+v656<<(uint(int32(2))%32))))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	if v38 == v681 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v825 = v795
	v846 = base.B2i32(v686 == int32(0))
	goto L127
L130:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v683 == v38 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v686 = v654
	goto L132
L132:
	;
	v687 = *(*int64)(unsafe.Add(mBase, uint32(v680)+8))
	if v687 != int64(0) {
		v721 = v655
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v685 = v680
	goto L135
L134:
	;
	v685 = v654
	goto L135
L135:
	;
	v686 = v685
	goto L132
L136:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v723 == v724 {
		goto L145
	} else {
		goto L146
	}
L137:
	;
	v690 = int32(0)
	v693 = F_errstart(m, int32(17), v690)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	if v693 == int32(0) {
		v721 = v690
		goto L136
	} else {
		goto L139
	}
L139:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v701 = F_format_type_be(m, v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	v704 = F_format_type_be(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v701
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_13), v27+int32(48))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(275), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	v721 = v690
	goto L136
L145:
	;
	v730 = v721
	v750 = int64(1)
	goto L148
L146:
	;
	v795 = v721
	goto L147
L147:
	;
	v817 = v656 + int32(1)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v817 < v818 {
		v654 = v686
		v655 = v795
		v656 = v817
		goto L128
	} else {
		goto L160
	}
L148:
	;
	if v750 == int64(7) {
		v787 = v730
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v795 = v787
	goto L147
L150:
	;
	v789 = v750 + int64(1)
	if v789 != int64(8) {
		v730 = v787
		v750 = v789
		goto L148
	} else {
		goto L159
	}
L151:
	;
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v680)+16))
	if v753&(int64(1)<<(uint(v750)%64)) != int64(0) {
		v787 = v730
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v759 = int32(0)
	v762 = F_errstart(m, int32(17), v759)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	if v762 == int32(0) {
		v787 = v759
		goto L150
	} else {
		goto L154
	}
L154:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v770 = F_format_type_be(m, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v770
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+40)) = uint32(v750)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v40
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(32))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	v787 = v759
	goto L150
L159:
	;
	goto L149
L160:
	;
	goto L129
L161:
	;
	goto L124
L162:
	;
	if v876 == int32(0) {
		v901 = v873
		goto L121
	} else {
		goto L163
	}
L163:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v36 + int32(8)
	F_errmsg(m, int32(_a_F_spgvalidate_15), v27+int32(16))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(308), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	v901 = v873
	goto L121
L167:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	m.G0 = v27 + int32(256)
	return v901 & int32(1)
}
func F_split_part(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	v10 = m.G0
	v12 = v10 - int32(1072)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v12 + int32(1072)
	return v297
L5:
	;
	v284 = v275 - v277
	v286 = v284 + int32(4)
	v287 = F_palloc(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	v275 = v234
	v277 = v220 + v229
	goto L5
L7:
	;
	if v243 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L8:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	if v209 < int32(2) {
		v275 = v215
		v277 = v208
		goto L5
	} else {
		goto L83
	}
L9:
	;
	v198 = int32(1)
	v199 = v15 + v198
	v201 = v15 + int32(4)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v202&v198 != 0 {
		goto L80
	} else {
		goto L81
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v23 == int32(1) {
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
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L76
	}
L13:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v54 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v28&int32(254) == int32(2) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v41 = int32(1)
	if v23&v41 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v41)%32)) - v41
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v37 = v26
	goto L19
L18:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L19
L19:
	;
	if v28 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v40 = v26
	goto L22
L21:
	;
	v40 = v37
	goto L22
L22:
	;
	v53 = v40
	goto L13
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	if v53 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v57 = int32(4)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v59&int32(254) == int32(2) {
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
	if v54&v72 != 0 {
		v84 = int32(base.Ui32(v54)>>(uint(v72)%32)) - v72
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v68 = v57
	goto L30
L29:
	;
	v68 = base.B2i32(v59 == int32(18)) << (uint(v57) % 32)
	goto L30
L30:
	;
	if v59 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v71 = v57
	goto L33
L32:
	;
	v71 = v68
	goto L33
L33:
	;
	v84 = v71
	goto L24
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
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
	v297 = v88
	goto L4
L39:
	;
	switch v22 + int32(1) {
	case 0, 2:
		v297 = v15
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
	F_text_position_setup(m, v15, v20, v101, v12)
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
	v297 = v97
	goto L4
L44:
	;
	v104 = F_text_position_next(m, v12)
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
	switch v22 + int32(1) {
	case 0, 2:
		v297 = v15
		goto L4
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if int32(0) <= v22 {
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
	v297 = v111
	goto L4
L51:
	;
	v118 = int32(2)
	goto L52
L52:
	;
	v129 = F_text_position_next(m, v12)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	if v22 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v129 != 0 {
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
	v133 = int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v135&v133 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v158 = v118 + v22 + int32(1)
	if v158 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v138 = v133
	goto L61
L60:
	;
	v138 = int32(4)
	goto L61
L61:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1056))
	v143 = v141 + v142
	v144 = v15 + v138 + v53 - v143
	v146 = v144 + int32(4)
	v147 = F_palloc(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v146 << (uint(int32(2)) % 32)
	if v144 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v297 = v147
	goto L4
L64:
	;
	v154 = F__emscripten_memcpy_bulkmem(m, v147+int32(4), v143, v144)
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	v162 = F_cstring_to_text(m, int32(_a_F_split_part_0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v164 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1052)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1068)) = v164
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1064)) = v168
	v170 = F_text_position_next(m, v12)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	v297 = v162
	goto L4
L71:
	;
	v172 = int32(1)
	v173 = v15 + v172
	v175 = v15 + int32(4)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v176&v172 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v179 = v173
	goto L74
L73:
	;
	v179 = v175
	goto L74
L74:
	;
	if v170 == int32(0) {
		v243 = v158
		v247 = v173
		v248 = v175
		v249 = v179
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v208 = v179
	v209 = v158
	v213 = v173
	v214 = v175
	goto L8
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_split_part_1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_split_part_2), int32(_a_F_split_part_3), int32(_a_F_split_part_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v205 = v199
	goto L82
L81:
	;
	v205 = v201
	goto L82
L82:
	;
	v208 = v205
	v209 = v22
	v213 = v199
	v214 = v201
	goto L8
L83:
	;
	v220 = v215
	v223 = v209 - int32(1)
	goto L84
L84:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1056))
	v230 = F_text_position_next(m, v12)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v243 = v223
	v247 = v213
	v248 = v214
	v249 = v220 + v229
	goto L7
L86:
	;
	goto L85
L87:
	;
	if v230 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	v236 = v223 - int32(1)
	if int32(0) < v236 {
		v220 = v234
		v223 = v236
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L6
L90:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v252&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v270 = F_palloc(m, int32(4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L101
	}
L93:
	;
	v255 = v247
	goto L95
L94:
	;
	v255 = v248
	goto L95
L95:
	;
	v257 = v255 - v249 + v53
	v259 = v257 + int32(4)
	v260 = F_palloc(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v259 << (uint(int32(2)) % 32)
	if v257 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v297 = v260
	goto L4
L98:
	;
	v267 = F__emscripten_memcpy_bulkmem(m, v260+int32(4), v249, v257)
	mBase = m.M
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = int32(16)
	v297 = v270
	goto L4
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v286 << (uint(int32(2)) % 32)
	if v284 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v297 = v287
	goto L4
L104:
	;
	v294 = F__emscripten_memcpy_bulkmem(m, v287+int32(4), v277, v284)
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L103
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
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v11 = m.G0
	v12 = int32(144)
	v13 = v11 - v12
	m.G0 = v13
	v18 = F__emscripten_memset_bulkmem(m, v13, base.I32_extend8_s(int32(0)), v12)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(_a_F_sscanf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = l0
	v25 = F_vfscanf(m, v18, l1, l2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		m.G0 = v18 + int32(144)
		m.G0 = v8 + int32(16)
		return v25
	}
}
func F_storeOperators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v22 = F_table_open(m, int32(2602), int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L61
	}
L4:
	;
	F_sequence_close(m, v22, int32(3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L60
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v26 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v30 = v18 + int32(96)
	v32 = v18 + int32(88)
	v34 = v18 + int32(80)
	v50 = int32(0)
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v50<<(uint(int32(2))%32))))
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+8)))
	v61 = F_SearchSysCacheExists(m, int32(4), l2, v58, v59, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v64
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(56)))) = uint8(v64)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v66
	v80 = F_GetNewOidWithIndex(m, v22, int32(2756), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v61 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v86
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v88
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v92 = int32(111)
	goto L17
L16:
	;
	v92 = int32(115)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = l1
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v104 = F_heap_form_tuple(m, v99, v18-int32(-64), v18+int32(48))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_CatalogTupleInsert(m, v22, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v104)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(2602)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(2617)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v117
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v127 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v128 = int32(110)
	goto L23
L22:
	;
	v128 = int32(97)
	goto L23
L23:
	;
	F_recordDependencyOn(m, v18+int32(36), v18+int32(24), v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+25)))
	if v133 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v134 = int32(2753)
	goto L27
L26:
	;
	v134 = int32(2616)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v136
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v146 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v147 = int32(105)
	goto L30
L29:
	;
	v147 = int32(97)
	goto L30
L30:
	;
	F_recordDependencyOn(m, v18+int32(36), v18+int32(24), v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v151 = F_typeDepNeeded(m, v150, v56)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v151 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(1247)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v155
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v165 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v170 == v171 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v166 = int32(110)
	goto L38
L37:
	;
	v166 = int32(97)
	goto L38
L38:
	;
	F_recordDependencyOn(m, v18+int32(36), v18+int32(24), v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v194 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v173 = F_typeDepNeeded(m, v170, v56)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v173 == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(1247)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v179
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v189 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v190 = int32(110)
	goto L46
L45:
	;
	v190 = int32(97)
	goto L46
L46:
	;
	F_recordDependencyOn(m, v18+int32(36), v18+int32(24), v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(2753)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
	if v206 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_storeOperators[0]))
	if v211 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v207 = int32(110)
	goto L53
L52:
	;
	v207 = int32(97)
	goto L53
L53:
	;
	F_recordDependencyOn(m, v18+int32(36), v18+int32(24), v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v213 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2602), v80, v213, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v218 = v50 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v218 < v219 {
		v50 = v218
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
	m.G0 = v18 + int32(112)
	return
L61:
	;
	F_errcode(m, int32(_a_F_storeOperators_0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v251 = F_format_type_be(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v254 = F_format_type_be(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v256 = F_NameListToString(m, l0)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v249
	F_errmsg(m, int32(_a_F_storeOperators_1), v18)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_storeOperators_2), int32(1489), int32(_a_F_storeOperators_3))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v7 = l1 & int32(255)
	if v7 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == l1&int32(255) {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v98 = v88
	goto L1
L3:
	;
	v78 = v73
	goto L21
L4:
	;
	v73 = v65
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
	v63 = F_strlen(m, l0)
	mBase = m.M
	v98 = v63 + l0
	goto L1
L8:
	;
	v12 = l0
	goto L11
L9:
	;
	v25 = l0
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v34 = int32(-2139062144)
	if (int32(16843008)-v31|v31)&v34 != v34 {
		v65 = v25
		goto L4
	} else {
		goto L16
	}
L11:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v17 == int32(0) {
		v88 = v12
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v25 = v22
	goto L10
L13:
	;
	if l1&int32(255) == v17 {
		v88 = v12
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v22 = v12 + int32(1)
	if v22&int32(3) != 0 {
		v12 = v22
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v40 = v25
	v42 = v31
	goto L17
L17:
	;
	v46 = v42 ^ v7*int32(16843009)
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 != v49 {
		v65 = v40
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v73 = v55
	goto L3
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v55 = v40 + int32(4)
	v59 = int32(-2139062144)
	if (v53|(int32(16843008)-v53))&v59 == v59 {
		v40 = v55
		v42 = v53
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v80 == int32(0) {
		v88 = v78
		goto L2
	} else {
		goto L23
	}
L22:
	;
	v88 = v78
	goto L2
L23:
	;
	if v80 != l1&int32(255) {
		v78 = v78 + int32(1)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v104 = v98
	goto L27
L26:
	;
	v104 = int32(0)
	goto L27
L27:
	;
	return v104
}
func F_strdup(m *base.Module, l0 int32) int32 {
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v62 = v60 + int32(1)
	v63 = F_emscripten_builtin_malloc(m, v62)
	mBase = m.M
	if v63 == int32(0) {
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
	goto L20
L20:
	;
	if v62 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v69
L22:
	;
	v68 = F__emscripten_memcpy_bulkmem(m, v63, l0, v62)
	mBase = m.M
	v69 = v68
	goto L24
L23:
	;
	v69 = v63
	goto L24
L24:
	;
	goto L21
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
	var v58 int32
	_ = v58
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
	var v89 int32
	_ = v89
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v168 int32
	_ = v168
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
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
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
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int64
	_ = v276
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
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
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
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
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v736 int32
	_ = v736
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v781 int32
	_ = v781
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
	return v781
L5:
	;
	if v25 == int32(0) {
		v781 = v3
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
		v781 = v3
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
		v89 = v40
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
		v781 = v3
		goto L4
	} else {
		goto L27
	}
L17:
	;
	if v89 != 0 {
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
		v89 = v40
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = v25 + int32(1)
	v58 = v46
	goto L20
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v70 = int32(0)
	v71 = base.B2i32(v69 != v70)
	if v69 == v70 {
		v82 = v56
		v89 = v71
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v82 = v56
	v89 = v71
	goto L17
L22:
	;
	v80 = v58<<(uint(int32(8))%32)&int32(_a_F_strstr_0) | v69
	if v80 != v51 {
		v56 = v56 + int32(1)
		v58 = v80
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
		v781 = v3
		goto L4
	} else {
		goto L42
	}
L31:
	;
	if v168 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v162 = v106
	v168 = base.B2i32(v107 != v108)
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
	v136 = v106
	v137 = v121
	goto L35
L35:
	;
	v149 = v136 + int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v151 = int32(0)
	v152 = base.B2i32(v150 != v151)
	if v150 == v151 {
		v162 = v149
		v168 = v152
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v162 = v149
	v168 = v152
	goto L31
L37:
	;
	v157 = (v150 | v137) << (uint(int32(8)) % 32)
	if v157 != v132 {
		v136 = v149
		v137 = v157
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
	v269 = int32(0)
	v270 = m.G0
	v272 = v270 - int32(1056)
	m.G0 = v272
	v276 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v272+int32(1048)))) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v272+int32(1040)))) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v272)+1032)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v272)+1024)) = v276
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v286 == v269 {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	if v257 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v252 = v187
	v257 = base.B2i32(v188 != v189)
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
	v208 = int32(_a_F_strstr_0)
	v220 = v205<<(uint(v197)%32) | v205&v208<<(uint(v201)%32) | (int32(base.Ui32(v205)>>(uint(v201)%32))&v208 | int32(base.Ui32(v205)>>(uint(v197)%32)))
	if v204 == v220 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v225 = v187
	v227 = v204
	goto L50
L50:
	;
	v237 = v225 + int32(1)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	v239 = int32(0)
	v240 = base.B2i32(v238 != v239)
	if v238 == v239 {
		v252 = v237
		v257 = v240
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v252 = v237
	v257 = v240
	goto L46
L52:
	;
	v245 = v227<<(uint(int32(8))%32) | v238
	if v245 != v220 {
		v225 = v237
		v227 = v245
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v267 = v252 - int32(3)
	goto L56
L55:
	;
	v267 = int32(0)
	goto L56
L56:
	;
	return v267
L57:
	;
	m.G0 = v272 + int32(1056)
	v781 = v761
	goto L4
L58:
	;
	v449 = int32(1)
	v453 = base.B2i32(base.Ui32(v441+v449) < base.Ui32(v438+v449))
	if base.Ui32(v441+v449) < base.Ui32(v438+v449) {
		goto L97
	} else {
		goto L98
	}
L59:
	;
	v352 = int32(1)
	v354 = v331
	v355 = v352
	v356 = v269
	v360 = v332
	v365 = v352
	goto L70
L60:
	;
	v761 = int32(0)
	goto L57
L61:
	;
	v435 = v335
	v438 = int32(-1)
	v439 = v339
	v441 = v341
	v444 = int32(1)
	goto L58
L62:
	;
	v335 = int32(1)
	v339 = v3
	v341 = int32(-1)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v292 = v286
	v295 = v3
	goto L65
L65:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v25))))
	if v306 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v331 = int32(1)
	v332 = int32(-1)
	if base.Ui32(v331) < base.Ui32(v315) {
		goto L59
	} else {
		goto L69
	}
L67:
	;
	v314 = int32(1)
	v315 = v295 + v314
	*(*int32)(unsafe.Add(mBase, uint32(v272+v292&int32(255)<<(uint(int32(2))%32)))) = v315
	v323 = v272 + int32(1024) + int32(base.Ui32(v292)>>(uint(int32(3))%32))&int32(28)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v324 | v314<<(uint(v292)%32)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315+l1))))
	if v330 != 0 {
		v292 = v330
		v295 = v315
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v335 = v331
	v339 = v315
	v341 = v332
	goto L61
L70:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360+l1+v355))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+l1))))
	if v370 == v372 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v394 = int32(1)
	v397 = int32(0)
	v398 = v394
	v399 = v394
	v400 = int32(-1)
	v406 = v394
	goto L83
L72:
	;
	v390 = v389 + v386
	if base.Ui32(v390) < base.Ui32(v315) {
		v354 = v390
		v355 = v389
		v356 = v386
		v360 = v387
		v365 = v388
		goto L70
	} else {
		goto L82
	}
L73:
	;
	if v355 == v365 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v372) < base.Ui32(v370) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v386 = v356 + v365
	v387 = v360
	v388 = v365
	v389 = int32(1)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v386 = v356
	v387 = v360
	v388 = v365
	v389 = v355 + int32(1)
	goto L72
L79:
	;
	v386 = v354
	v387 = v360
	v388 = v354 - v360
	v389 = int32(1)
	goto L72
L80:
	;
	goto L81
L81:
	;
	v382 = int32(1)
	v386 = v356 + v382
	v387 = v356
	v388 = v382
	v389 = v382
	goto L72
L82:
	;
	goto L71
L83:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+l1+v398))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+l1))))
	if v413 == v415 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v435 = v388
	v438 = v430
	v439 = v315
	v441 = v387
	v444 = v431
	goto L58
L85:
	;
	v433 = v432 + v429
	if base.Ui32(v433) < base.Ui32(v315) {
		v397 = v429
		v398 = v432
		v399 = v433
		v400 = v430
		v406 = v431
		goto L83
	} else {
		goto L95
	}
L86:
	;
	if v398 == v406 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(v413) < base.Ui32(v415) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v429 = v397 + v406
	v430 = v400
	v431 = v406
	v432 = int32(1)
	goto L85
L90:
	;
	goto L91
L91:
	;
	v429 = v397
	v430 = v400
	v431 = v406
	v432 = v398 + int32(1)
	goto L85
L92:
	;
	v429 = v399
	v430 = v400
	v431 = v399 - v400
	v432 = int32(1)
	goto L85
L93:
	;
	goto L94
L94:
	;
	v425 = int32(1)
	v429 = v397 + v425
	v430 = v397
	v431 = v425
	v432 = v425
	goto L85
L95:
	;
	goto L84
L96:
	;
	v534 = v439 | int32(63)
	v536 = v25
	v537 = int32(0)
	v541 = v25
	goto L127
L97:
	;
	v454 = v444
	goto L99
L98:
	;
	v454 = v435
	goto L99
L99:
	;
	v455 = l1 + v454
	if base.Ui32(v441+v449) < base.Ui32(v438+v449) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v456 = v438
	goto L102
L101:
	;
	v456 = v441
	goto L102
L102:
	;
	v458 = v456 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v458) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	if v520 != 0 {
		goto L121
	} else {
		goto L122
	}
L104:
	;
	v520 = int32(0)
	goto L103
L105:
	;
	v494 = v489
	v495 = v490
	v496 = v491
	goto L115
L106:
	;
	if (l1|v455)&int32(3) != 0 {
		v489 = l1
		v490 = v455
		v491 = v458
		goto L105
	} else {
		goto L109
	}
L107:
	;
	v482 = l1
	v483 = v455
	v484 = v458
	goto L108
L108:
	;
	if v484 == int32(0) {
		goto L104
	} else {
		goto L114
	}
L109:
	;
	v466 = l1
	v467 = v455
	v468 = v458
	goto L110
L110:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	if v471 != v472 {
		v489 = v466
		v490 = v467
		v491 = v468
		goto L105
	} else {
		goto L112
	}
L111:
	;
	v482 = v477
	v483 = v475
	v484 = v479
	goto L108
L112:
	;
	v474 = int32(4)
	v475 = v467 + v474
	v477 = v466 + v474
	v479 = v468 - v474
	if base.Ui32(int32(3)) < base.Ui32(v479) {
		v466 = v477
		v467 = v475
		v468 = v479
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v489 = v482
	v490 = v483
	v491 = v484
	goto L105
L115:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	if v499 == v500 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v520 = v499 - v500
	goto L103
L117:
	;
	v502 = int32(1)
	v507 = v496 - v502
	if v507 != 0 {
		v494 = v494 + v502
		v495 = v495 + v502
		v496 = v507
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
	v523 = v439 + (v456 ^ int32(-1))
	if base.Ui32(v523) < base.Ui32(v456) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v531 = v454
	v532 = v439 - v454
	goto L96
L124:
	;
	v525 = v456
	goto L126
L125:
	;
	v525 = v523
	goto L126
L126:
	;
	v531 = v525 + int32(1)
	v532 = int32(0)
	goto L96
L127:
	;
	if base.Ui32(v439) <= base.Ui32(v541-v536) {
		v665 = v541
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v666 = int32(0)
	v669 = v536 + v439
	v670 = int32(1)
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669-v670))))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(1024)+int32(base.Ui32(v672)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v678)>>(uint(v672)%32))&v670 == v666 {
		v536 = v669
		v537 = v666
		v541 = v665
		goto L127
	} else {
		goto L162
	}
L130:
	;
	v552 = int32(0)
	v557 = base.B2i32(v534 != v552)
	if v541&int32(3) == v552 {
		v583 = v541
		v585 = v534
		v586 = v557
		goto L134
	} else {
		goto L135
	}
L131:
	;
	if v656 != 0 {
		goto L157
	} else {
		goto L158
	}
L132:
	;
	v656 = int32(0)
	goto L131
L133:
	;
	v634 = v627
	v636 = v629
	goto L151
L134:
	;
	if v586 == int32(0) {
		goto L132
	} else {
		goto L142
	}
L135:
	;
	if v534 == int32(0) {
		v583 = v541
		v585 = v534
		v586 = v557
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v566 = v541
	v568 = v534
	goto L137
L137:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if v571 == int32(0) {
		v627 = v566
		v629 = v568
		goto L133
	} else {
		goto L139
	}
L138:
	;
	v583 = v578
	v585 = v574
	v586 = v576
	goto L134
L139:
	;
	v573 = int32(1)
	v574 = v568 - v573
	v575 = int32(0)
	v576 = base.B2i32(v574 != v575)
	v578 = v566 + v573
	if v578&int32(3) == v575 {
		v583 = v578
		v585 = v574
		v586 = v576
		goto L134
	} else {
		goto L140
	}
L140:
	;
	if v574 != 0 {
		v566 = v578
		v568 = v574
		goto L137
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	if v590 == int32(0) {
		v620 = v583
		v622 = v585
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if v622 == int32(0) {
		goto L132
	} else {
		goto L150
	}
L144:
	;
	if base.Ui32(v585) < base.Ui32(int32(4)) {
		v620 = v583
		v622 = v585
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v600 = v583
	v602 = v585
	goto L146
L146:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	v607 = v606 ^ int32(0)
	v610 = int32(-2139062144)
	if (int32(16843008)-v607|v607)&v610 != v610 {
		v627 = v600
		v629 = v602
		goto L133
	} else {
		goto L148
	}
L147:
	;
	v620 = v615
	v622 = v617
	goto L143
L148:
	;
	v614 = int32(4)
	v615 = v600 + v614
	v617 = v602 - v614
	if base.Ui32(int32(3)) < base.Ui32(v617) {
		v600 = v615
		v602 = v617
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v627 = v620
	v629 = v622
	goto L133
L151:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if int32(0) == v639 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L132
L153:
	;
	v656 = v634
	goto L131
L154:
	;
	goto L155
L155:
	;
	v641 = int32(1)
	v644 = v636 - v641
	if v644 != 0 {
		v634 = v634 + v641
		v636 = v644
		goto L151
	} else {
		goto L156
	}
L156:
	;
	goto L152
L157:
	;
	v658 = v656
	goto L159
L158:
	;
	v658 = v541 + v534
	goto L159
L159:
	;
	if v656 == int32(0) {
		v665 = v658
		goto L129
	} else {
		goto L160
	}
L160:
	;
	if base.Ui32(v656-v536) < base.Ui32(v439) {
		v761 = v552
		goto L57
	} else {
		goto L161
	}
L161:
	;
	v665 = v658
	goto L129
L162:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v272+v672<<(uint(int32(2))%32))))
	if v687 != v439 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v689 = v439 - v687
	if base.Ui32(v537) < base.Ui32(v689) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v537) < base.Ui32(v458) {
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v691 = v689
	goto L168
L167:
	;
	v691 = v537
	goto L168
L168:
	;
	v536 = v536 + v691
	v537 = int32(0)
	v541 = v665
	goto L127
L169:
	;
	v536 = v536 + (v698 - v456)
	v537 = int32(0)
	v541 = v665
	goto L127
L170:
	;
	v695 = v458
	goto L172
L171:
	;
	v695 = v537
	goto L172
L172:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v695))))
	if v697 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v698 = v695
	v704 = v697
	goto L176
L174:
	;
	goto L175
L175:
	;
	v736 = v458
	goto L180
L176:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698+v536))))
	if v713 != v704&int32(255) {
		goto L169
	} else {
		goto L178
	}
L177:
	;
	goto L175
L178:
	;
	v718 = v698 + int32(1)
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v718))))
	if v720 != 0 {
		v698 = v718
		v704 = v720
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	if base.Ui32(v736) <= base.Ui32(v537) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v536 = v536 + v531
	v537 = v532
	v541 = v665
	goto L127
L182:
	;
	v761 = v536
	goto L57
L183:
	;
	goto L184
L184:
	;
	v751 = v736 - int32(1)
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v751))))
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v536))))
	if v753 == v755 {
		v736 = v751
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
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
	var v122 int32
	_ = v122
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
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v299 int32
	_ = v299
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
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int64
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
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
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int64
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v822 int32
	_ = v822
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1040 int32
	_ = v1040
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1085 int32
	_ = v1085
	var v1105 int32
	_ = v1105
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1153 int32
	_ = v1153
	var v1169 int32
	_ = v1169
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
	m.G0 = v1169 + int32(16)
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
		v1169 = v22
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
	v666 = v22
	goto L13
L13:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v671 < int32(0) {
		v1169 = v666
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
	v666 = v22
	goto L13
L16:
	;
	if base.Ui32(v299) < base.Ui32(v104) {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	v299 = v105
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
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v281)
	v299 = v279
	goto L16
L21:
	;
	v279 = v105
	v281 = v108
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
	v122 = v108
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
	v279 = v272
	v281 = v259
	goto L20
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v135 != 0 {
		v1169 = v22
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v136 = int32(_a_F_subcolorcvec_0)
	v137 = v133 & v136
	if v137 != v122&v136 {
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
	v259 = v122
	goto L33
L33:
	;
	v272 = v119 + int32(1)
	if base.Ui32(v119) < base.Ui32(v113) {
		v119 = v272
		v122 = v259
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
		v1169 = v22
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
	v259 = v133
	goto L33
L61:
	;
	goto L28
L62:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v645 != 0 {
		v1169 = v22
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
	if v299 != v104 {
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
	if v387 <= v381 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v380 = v332
	v381 = int32(0)
	v387 = v333
	goto L73
L75:
	;
	goto L76
L76:
	;
	v337 = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if base.Ui32(v299) <= base.Ui32(v338) {
		v380 = v332
		v381 = v337
		v387 = v333
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v345 = v337
	v347 = v332
	goto L78
L78:
	;
	v359 = int32(12)
	v361 = v320 + v345*v359
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v361))) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v364
	v367 = v347 + v359
	v369 = v345 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v370 <= v369 {
		v380 = v367
		v381 = v369
		v387 = v370
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v380 = v367
	v381 = v369
	v387 = v370
	goto L73
L80:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v347+int32(16))))
	if base.Ui32(v374) < base.Ui32(v299) {
		v345 = v369
		v347 = v367
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if base.Ui32(v529) <= base.Ui32(v104) {
		goto L114
	} else {
		goto L115
	}
L83:
	;
	v527 = v380
	v528 = v381
	v529 = v299
	v530 = v381
	v534 = v387
	goto L82
L84:
	;
	goto L85
L85:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if base.Ui32(v104) < base.Ui32(v396) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v527 = v380
	v528 = v381
	v529 = v299
	v530 = v381
	v534 = v387
	goto L82
L87:
	;
	goto L88
L88:
	;
	v404 = v380
	v405 = v381
	v406 = v299
	v407 = v381
	v413 = v396
	goto L89
L89:
	;
	if base.Ui32(v406) < base.Ui32(v413) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v527 = v514
	v528 = v507
	v529 = v516
	v530 = v518
	v534 = v519
	goto L82
L91:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	F_subcoloronerow(m, l0, v504, l2, l3, v22+int32(14))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L7
	} else {
		goto L111
	}
L92:
	;
	v473 = v320 + v468*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v469
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if base.Ui32(v104) < base.Ui32(v475) {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v459 = v320 + v405*int32(12)
	v460 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v406 - v460
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v413
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+8)) = v464
	v468 = v405 + v460
	v469 = v406
	goto L92
L94:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if base.Ui32(v104) < base.Ui32(v443) {
		v468 = v440
		v469 = v441
		goto L92
	} else {
		goto L101
	}
L95:
	;
	v422 = v320 + v405*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v406
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v424 - int32(1)
	v429 = F_newhicolorrow(m, v313, int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if base.Ui32(v413) < base.Ui32(v406) {
		goto L93
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+8)) = v429
	F_subcoloronerow(m, l0, v429, l2, l3, v22+int32(14))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v440 = v405 + int32(1)
	v441 = v438
	goto L94
L100:
	;
	v440 = v405
	v441 = v406
	goto L94
L101:
	;
	v447 = v320 + v440*int32(12)
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v404)))
	*(*int64)(unsafe.Add(mBase, uint32(v447))) = v448
	v451 = v404 + int32(8)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v447)+8)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v504 = v454
	v507 = v440 + int32(1)
	goto L91
L102:
	;
	v477 = v104
	goto L104
L103:
	;
	v477 = v475
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v480 = F_newhicolorrow(m, v313, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+8)) = v480
	v484 = v468 + int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if base.Ui32(v485) <= base.Ui32(v104) {
		v504 = v480
		v507 = v484
		goto L91
	} else {
		goto L106
	}
L106:
	;
	v489 = v320 + v484*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v104 + int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v489)+4)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if base.Ui32(v494) < base.Ui32(v469) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v496 = F_newhicolorrow(m, v313, v493)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L110
	}
L108:
	;
	v498 = v493
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489)+8)) = v498
	v504 = v480
	v507 = v468 + int32(2)
	goto L91
L110:
	;
	v498 = v496
	goto L109
L111:
	;
	v514 = v404 + int32(12)
	v515 = int32(1)
	v516 = v508 + v515
	v518 = v407 + v515
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v519 <= v518 {
		v527 = v514
		v528 = v507
		v529 = v516
		v530 = v518
		v534 = v519
		goto L82
	} else {
		goto L112
	}
L112:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	if base.Ui32(v521) <= base.Ui32(v104) {
		v404 = v514
		v405 = v507
		v406 = v516
		v407 = v518
		v413 = v521
		goto L89
	} else {
		goto L113
	}
L113:
	;
	goto L90
L114:
	;
	v545 = v320 + v528*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v545)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v529
	v549 = F_newhicolorrow(m, v313, int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L117
	}
L115:
	;
	v559 = v528
	v561 = v534
	goto L116
L116:
	;
	if v530 < v561 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545)+8)) = v549
	F_subcoloronerow(m, l0, v549, l2, l3, v22+int32(14))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	v559 = v528 + int32(1)
	v561 = v556
	goto L116
L119:
	;
	v567 = v527
	v568 = v559
	v570 = v530
	goto L122
L120:
	;
	v602 = v559
	goto L121
L121:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v313)+88))
	if v616 != 0 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v582 = int32(12)
	v584 = v320 + v568*v582
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v567)))
	*(*int64)(unsafe.Add(mBase, uint32(v584))) = v585
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+8)) = v587
	v591 = int32(1)
	v592 = v568 + v591
	v594 = v570 + v591
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v313)+84))
	if v594 < v595 {
		v567 = v567 + v582
		v568 = v592
		v570 = v594
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v602 = v592
	goto L121
L124:
	;
	goto L123
L125:
	;
	F_pfree(m, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+84)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v313)+88)) = v320
	goto L62
L128:
	;
	goto L127
L129:
	;
	F_subcoloronechr(m, l0, v104, l2, l3, v22+int32(14))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	goto L62
L131:
	;
	v648 = int32(1)
	if v648 < v101 {
		v100 = v100 + int32(8)
		v101 = v101 - v648
		goto L14
	} else {
		goto L132
	}
L132:
	;
	goto L15
L133:
	;
	v675 = v24 + int32(28)
	v678 = v675 + v671<<(uint(int32(2))%32)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v679 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v678))) = v682
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v24)+96))
	v688 = base.I32_div_s(int32(2147483647), v685<<(uint(int32(1))%32))
	if v688 <= v682 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v876 = v679
	goto L136
L136:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	if v883 <= int32(0) {
		v1169 = v666
		goto L1
	} else {
		goto L164
	}
L137:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v858 != 0 {
		v1169 = v666
		goto L1
	} else {
		goto L163
	}
L138:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v690)+24)) = int32(101)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+12))
	if v694 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v702 = F_repalloc_extended(m, v698, v685*v682<<(uint(int32(2))%32))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L7
	} else {
		goto L144
	}
L141:
	;
	v696 = v694
	goto L143
L142:
	;
	v696 = int32(12)
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+12)) = v696
	goto L137
L144:
	;
	if v702 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v706)+24)) = int32(101)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	if v710 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v702
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v718 = v716 - int32(1)
	if int32(0) <= v718 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v712 = v710
	goto L150
L149:
	;
	v712 = int32(12)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709)+12)) = v712
	goto L137
L151:
	;
	v725 = v718
	v726 = v715
	goto L154
L152:
	;
	v822 = v715
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v822 << (uint(int32(1)) % 32)
	goto L137
L154:
	;
	if int32(0) < v726 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v822 = v799
	goto L153
L156:
	;
	v742 = v725 * v726
	v745 = v702 + v742<<(uint(int32(2))%32)
	v746 = int32(1)
	v757 = int32(0)
	goto L159
L157:
	;
	v799 = v726
	goto L158
L158:
	;
	if int32(0) < v725 {
		v725 = v725 - int32(1)
		v726 = v799
		goto L154
	} else {
		goto L162
	}
L159:
	;
	v772 = int32(1)
	v773 = v757 << (uint(v772) % 32)
	v776 = int32(*(*int16)(unsafe.Add(mBase, uint32(v773+(v702+v742<<(uint(v746)%32))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v745+v726<<(uint(v746)%32)+v773))) = uint16(v776)
	*(*uint16)(unsafe.Add(mBase, uint32(v773+v745))) = uint16(v776)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v785 = v780 + v776*int32(24) + int32(4)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = v786 + v772
	v791 = v757 + v772
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v791 < v792 {
		v757 = v791
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v799 = v792
	goto L158
L161:
	;
	goto L160
L162:
	;
	goto L155
L163:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v675+v859<<(uint(int32(2))%32))))
	v876 = v863
	goto L136
L164:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v893 = v886
	v895 = v887
	v896 = v883
	v904 = int32(0)
	goto L165
L165:
	;
	v908 = int32(0)
	if v908 < v893 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v1169 = v666
	goto L1
L167:
	;
	v916 = v908
	v917 = v895
	goto L170
L168:
	;
	v1137 = v893
	v1139 = v895
	v1140 = v896
	goto L169
L169:
	;
	v1153 = v904 + int32(1)
	if v1153 < v1140 {
		v893 = v1137
		v895 = v1139
		v896 = v1140
		v904 = v1153
		goto L165
	} else {
		goto L214
	}
L170:
	;
	if v916&v876 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v1137 = v1130
	v1139 = v1127
	v1140 = v1132
	goto L169
L172:
	;
	v1127 = v917 + int32(2)
	v1129 = v916 + int32(1)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v1129 < v1130 {
		v916 = v1129
		v917 = v1127
		goto L170
	} else {
		goto L213
	}
L173:
	;
	v933 = int32(*(*int16)(unsafe.Add(mBase, uint32(v917))))
	v935 = v933 * int32(24)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v937 = v935 + v936
	v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+8)))
	if v938 != int32(_a_F_subcolorcvec_0) {
		v959 = v938
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v961)+12))
	if v962 != 0 {
		v989 = int32(_a_F_subcolorcvec_0)
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v937)+4))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v941+v942 == int32(1) {
		v959 = v933
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v946 = F_newcolor(m, v24)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	if v946 == int32(-1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v959 = int32(_a_F_subcolorcvec_0)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v951+v935)+8)) = uint16(v946)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v954+v946*int32(24))+8)) = uint16(v946)
	v959 = v946
	goto L174
L181:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v990 != 0 {
		v1169 = v666
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v963 = int32(_a_F_subcolorcvec_0)
	if v933&v963 == v959&v963 {
		v989 = v933
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v970 = int32(4)
	v971 = v968 + v935 + v970
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v971)))
	v973 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v971))) = v972 - v973
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v982 = v976 + base.I32_extend16_s(v959)*int32(24) + v970
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)))
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = v983 + v973
	*(*uint16)(unsafe.Add(mBase, uint32(v917))) = uint16(v959)
	v989 = v959
	goto L181
L184:
	;
	v992 = v989 & int32(_a_F_subcolorcvec_0)
	v993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v666)+14)))
	if v992 == v993 {
		goto L172
	} else {
		goto L185
	}
L185:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_subcolorcvec[0]))
	if v997 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L7
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v1000 <= v1001 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L188
L190:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1105 != 0 {
		v1169 = v666
		goto L1
	} else {
		goto L212
	}
L191:
	;
	F_createarc(m, v995, int32(112), base.I32_extend16_s(v989), l2, l3)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L7
	} else {
		goto L211
	}
L192:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v1003 == int32(0) {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v1033 == int32(0) {
		goto L191
	} else {
		goto L203
	}
L195:
	;
	v1010 = v1003
	goto L196
L196:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+12))
	if v1025 != l3 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L191
L198:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+16))
	if v1032 != 0 {
		v1010 = v1032
		goto L196
	} else {
		goto L202
	}
L199:
	;
	v1027 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1010)+4)))
	if v1027 != v992 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	if v1029 == int32(112) {
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
	v1040 = v1033
	goto L204
L204:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+8))
	if v1055 != l2 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L191
L206:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+24))
	if v1062 != 0 {
		v1040 = v1062
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v1057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1040)+4)))
	if v1057 != v992 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	if v1059 == int32(112) {
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
	*(*uint16)(unsafe.Add(mBase, uint32(v666)+14)) = uint16(v989)
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
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v156 int32
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
	var v173 int32
	_ = v173
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
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
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
		goto L45
	} else {
		goto L46
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
	F_subcoloronerow(m, l0, v265, l2, l3, l4)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L70
	}
L43:
	;
	if v149 == v164 {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v181 = v120 + v173*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = l1
	v185 = F_newhicolorrow(m, v14, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L55
	}
L45:
	;
	v171 = v132
	v173 = v6
	goto L44
L46:
	;
	goto L47
L47:
	;
	v141 = v132
	v143 = v6
	goto L48
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if base.Ui32(v149) < base.Ui32(l1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v164) <= base.Ui32(l1) {
		goto L43
	} else {
		goto L54
	}
L50:
	;
	v151 = int32(12)
	v153 = v120 + v143*v151
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v153))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v156
	v159 = v141 + v151
	v161 = v143 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v161 < v162 {
		v141 = v159
		v143 = v161
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v171 = v159
	v173 = v161
	goto L44
L54:
	;
	v171 = v141
	v173 = v143
	goto L44
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+8)) = v185
	v260 = v171
	v261 = v173 + int32(1)
	v262 = v173
	v265 = v185
	goto L42
L56:
	;
	v191 = int32(12)
	v193 = v120 + v143*v191
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = v194
	v197 = v141 + int32(8)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v198
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v204 = v143 + int32(1)
	v260 = v141 + v191
	v261 = v204
	v262 = v204
	v265 = v202
	goto L42
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v164) < base.Ui32(l1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v208 = v120 + v143*int32(12)
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = l1 - v209
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v164
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v213
	v218 = v143 + v209
	goto L61
L60:
	;
	v218 = v143
	goto L61
L61:
	;
	v221 = v120 + v218*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = l1
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v225 = F_newhicolorrow(m, v14, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v225
	v229 = v218 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if base.Ui32(l1) < base.Ui32(v230) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v234 = v120 + v229*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = l1 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v241) < base.Ui32(l1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v249 = v229
	goto L65
L65:
	;
	v260 = v141 + int32(12)
	v261 = v249
	v262 = v143 + int32(1)
	v265 = v225
	goto L42
L66:
	;
	v243 = F_newhicolorrow(m, v14, v240)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	v245 = v240
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v245
	v249 = v218 + int32(2)
	goto L65
L69:
	;
	v245 = v243
	goto L68
L70:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v262 < v270 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v277 = v260
	v278 = v261
	v279 = v262
	goto L74
L72:
	;
	v306 = v261
	goto L73
L73:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v313 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v285 = int32(12)
	v287 = v120 + v278*v285
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v277)))
	*(*int64)(unsafe.Add(mBase, uint32(v287))) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+8)) = v290
	v294 = int32(1)
	v295 = v278 + v294
	v297 = v279 + v294
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v297 < v298 {
		v277 = v277 + v285
		v278 = v295
		v279 = v297
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v306 = v295
	goto L73
L76:
	;
	goto L75
L77:
	;
	F_pfree(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v120
	goto L1
L80:
	;
	goto L79
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = v10 + int32(3)
	if v8 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v133 < v136 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v10 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v64 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v26 = v10
	goto L6
L5:
	;
	v26 = v24
	goto L6
L6:
	;
	v33 = v10
	goto L8
L7:
	;
	v64 = v44
	goto L3
L8:
	;
	if v33 == v26 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v64 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v33))))
	if int32(246) < v39 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
	v33 = v56
	goto L8
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v48)>>(uint(v41&int32(7))%32))&v44 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 < v75 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v119 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v78 = v75
	goto L22
L21:
	;
	v78 = v76
	goto L22
L22:
	;
	v85 = v75
	goto L24
L23:
	;
	v119 = int32(1)
	goto L19
L24:
	;
	if v85 == v78 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v119 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v85))))
	if int32(246) < v93 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v95 = v93 - int32(97)
	if v95 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v95)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v110 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
	v85 = v110
	goto L24
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v122 + v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126 < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v128 = v123
	goto L36
L35:
	;
	v128 = v126
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v128
	goto L1
L37:
	;
	return v392
L38:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v236 <= v233 {
		goto L65
	} else {
		goto L66
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v136
	if v133 <= v136 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	goto L38
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v133-int32(1)))))
	if v145&int32(224) != int32(96) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(1)<<(uint(v145)%32)&int32(_a_F_swedish_ISO_8859_1_stem_0) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v158 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_1), int32(37))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v158 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v165
	switch v158 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L47
	default:
		goto L38
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L53
L48:
	;
	v169 = F_slice_del(m, l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	if int32(0) <= v169 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v392 = v169
	goto L37
L51:
	;
	if v225 != 0 {
		goto L38
	} else {
		goto L62
	}
L52:
	;
	v225 = v221
	goto L51
L53:
	;
	if v181 <= v182 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v221 = int32(0)
	goto L52
L55:
	;
	v225 = int32(-1)
	goto L51
L56:
	;
	goto L57
L57:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v181-v194))))
	if int32(121) < v199 {
		v221 = v194
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v201 = v199 - int32(98)
	if v201 < int32(0) {
		v221 = v194
		goto L52
	} else {
		goto L59
	}
L59:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v201)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v207)>>(uint(v201&int32(7))%32))&int32(1) == int32(0) {
		v221 = v194
		goto L52
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v181 - int32(1)
	goto L61
L61:
	;
	goto L54
L62:
	;
	v226 = F_slice_del(m, l0)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	if int32(0) <= v226 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v392 = v226
	goto L37
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v236
	v241 = v233 - int32(1)
	if v241 <= v236 {
		v275 = v233
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v279 = v233
	v281 = v235
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v279
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v279 < v284 {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v238
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v279 = v275
	v281 = v278
	goto L67
L69:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v241))))
	if v245&int32(224) != int32(96) {
		v275 = v233
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if int32(1)<<(uint(v245)%32)&int32(_a_F_swedish_ISO_8859_1_stem_2) == int32(0) {
		v275 = v233
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v258 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_3), int32(7))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L44
	} else {
		goto L72
	}
L72:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v258 == int32(0) {
		v275 = v260
		goto L68
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v260
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v260 <= v264 {
		v275 = v260
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v267 = v260 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v267
	v270 = F_slice_del(m, l0)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	if v270 < int32(0) {
		v392 = v270
		goto L37
	} else {
		goto L76
	}
L76:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v275 = v274
	goto L68
L77:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v389
	v392 = int32(1)
	goto L37
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v279
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v284
	v290 = v279 - int32(1)
	if v290 <= v284 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v287
	goto L77
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v290))))
	if v294&int32(224) != int32(96) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if int32(1)<<(uint(v294)%32)&int32(_a_F_swedish_ISO_8859_1_stem_4) == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v307 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_5), int32(5))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L44
	} else {
		goto L83
	}
L83:
	;
	if v307 == int32(0) {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v287
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v312
	switch v307 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	case 2:
		goto L85
	default:
		goto L77
	}
L85:
	;
	v381 = F_slice_from_s(m, l0, int32(4), int32(_a_F_swedish_ISO_8859_1_stem_6))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L44
	} else {
		goto L104
	}
L86:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L92
L87:
	;
	v316 = F_slice_del(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L44
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v316 {
		goto L77
	} else {
		goto L89
	}
L89:
	;
	v392 = v316
	goto L37
L90:
	;
	if v372 != 0 {
		goto L77
	} else {
		goto L101
	}
L91:
	;
	v372 = v368
	goto L90
L92:
	;
	if v328 <= v329 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v368 = int32(0)
	goto L91
L94:
	;
	v372 = int32(-1)
	goto L90
L95:
	;
	goto L96
L96:
	;
	v341 = int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342+v328-v341))))
	if int32(118) < v346 {
		v368 = v341
		goto L91
	} else {
		goto L97
	}
L97:
	;
	v348 = v346 - int32(105)
	if v348 < int32(0) {
		v368 = v341
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v348)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v354)>>(uint(v348&int32(7))%32))&int32(1) == int32(0) {
		v368 = v341
		goto L91
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v328 - int32(1)
	goto L100
L100:
	;
	goto L93
L101:
	;
	v375 = F_slice_from_s(m, l0, int32(2), int32(_a_F_swedish_ISO_8859_1_stem_7))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L44
	} else {
		goto L102
	}
L102:
	;
	if int32(0) <= v375 {
		goto L77
	} else {
		goto L103
	}
L103:
	;
	v392 = v375
	goto L37
L104:
	;
	if int32(0) <= v381 {
		goto L77
	} else {
		goto L105
	}
L105:
	;
	v392 = v381
	goto L37
}
