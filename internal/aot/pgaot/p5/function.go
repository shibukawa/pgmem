package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v18 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+22)) = uint16(v18)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = m.T0[v22].(func(*base.Module, int32) int32)(m, v7+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
		if v27 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
				F_errmsg_internal(m, int32(_a_F_FunctionCall1Coll_0), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall1Coll_1), int32(1143), int32(_a_F_FunctionCall1Coll_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(32)
			return v23
		}
	}
}
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
	v20 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v20)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, v9+int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
		if v35 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v42
				F_errmsg_internal(m, int32(_a_F_FunctionCall3Coll_0), v9)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall3Coll_1), int32(1190), int32(_a_F_FunctionCall3Coll_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(48)
			return v31
		}
	}
}
func F_FunctionIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v201
L2:
	;
	return int32(0)
L3:
	;
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v19)
	v201 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(_a_F_FunctionIsVisibleExt_0), v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_FunctionIsVisibleExt_1), int32(1723), int32(_a_F_FunctionIsVisibleExt_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v38 = v34 + v35
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v39 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v13)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L60
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_FunctionIsVisibleExt[0]))
	v44 = int32(0)
	if v43 == v44 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+104)))
	v88 = F_makeString(m, v38+int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v82 == int32(0) {
		v192 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v82 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v50 <= int32(0) {
		v76 = v44
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v82 = v76
	goto L18
L23:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v56 = v50
	goto L26
L25:
	;
	v56 = v53
	goto L26
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v59 = int32(0)
	goto L27
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59<<(uint(int32(2))%32))))
	v68 = base.B2i32(v67 == v39)
	if v67 == v39 {
		v76 = v68
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v68
	goto L22
L29:
	;
	v70 = v59 + int32(1)
	if v70 != v56 {
		v59 = v70
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v88
	v95 = F_list_make1_impl(m, int32(1), v10+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v97 = int32(0)
	v102 = F_FuncnameGetCandidates(m, v95, v85, v97, v97, v97, v97, v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v102 == int32(0) {
		v192 = v3
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v107 = v85 << (uint(int32(2)) % 32)
	v109 = v38 + int32(136)
	v111 = v102
	goto L36
L36:
	;
	v118 = v111 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v107) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v192 = v3
	goto L14
L38:
	;
	if v180 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v180 = int32(0)
	goto L38
L40:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L50
L41:
	;
	if (v118|v109)&int32(3) != 0 {
		v149 = v118
		v150 = v109
		v151 = v107
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v142 = v118
	v143 = v109
	v144 = v107
	goto L43
L43:
	;
	if v144 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v126 = v118
	v127 = v109
	v128 = v107
	goto L45
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v131 != v132 {
		v149 = v126
		v150 = v127
		v151 = v128
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v142 = v137
	v143 = v135
	v144 = v139
	goto L43
L47:
	;
	v134 = int32(4)
	v135 = v127 + v134
	v137 = v126 + v134
	v139 = v128 - v134
	if base.Ui32(int32(3)) < base.Ui32(v139) {
		v126 = v137
		v127 = v135
		v128 = v139
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v149 = v142
	v150 = v143
	v151 = v144
	goto L40
L50:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v159 == v160 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v180 = v159 - v160
	goto L38
L52:
	;
	v162 = int32(1)
	v167 = v156 - v162
	if v167 != 0 {
		v154 = v154 + v162
		v155 = v155 + v162
		v156 = v167
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
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v192 = base.B2i32(v183 == l0)
	goto L14
L57:
	;
	goto L58
L58:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v185 != 0 {
		v111 = v185
		goto L36
	} else {
		goto L59
	}
L59:
	;
	goto L37
L60:
	;
	v201 = v192
	goto L1
}
func F_FunctionNext(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	v2 = int32(0)
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+121)))
	if v16 == v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v15 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v32 = F_ExecMakeTableFunctionResult(m, v23, v24, v25, v26, int32(base.Ui32(v27&int32(8))>>(uint(int32(3))%32)))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v40 = v20
	goto L6
L6:
	;
	v44 = F_tuplestore_gettupleslot(m, v40, base.B2i32(v15 == int32(1)), int32(0), v13)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v32
	F_tuplestore_rescan(m, v32)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = v32
	goto L6
L10:
	;
	return v13
L11:
	;
	v52 = int64(1)
	goto L13
L12:
	;
	v52 = int64(-1)
	goto L13
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v47 + v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	m.T0[v56].(func(*base.Module, int32))(m, v13)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v60 = l0 + int32(128)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v61 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = v2
	v69 = v12
	v72 = v2
	goto L18
L16:
	;
	v225 = v2
	v227 = v12
	goto L17
L17:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v233 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v78 = v75 + v72<<(uint(int32(5))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v79 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v225 = v210
	v227 = v212
	goto L17
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v91 = F_ExecMakeTableFunctionResult(m, v82, v83, v84, v85, int32(base.Ui32(v86&int32(8))>>(uint(int32(3))%32)))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	if base.B2i32(v97 == int64(-1))|base.B2i32(v47 <= v97) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v91
	F_tuplestore_rescan(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	if v117 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	m.T0[v106].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v114 = F_tuplestore_gettupleslot(m, v109, base.B2i32(v15 == int32(1)), int32(0), v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	goto L25
L31:
	;
	v219 = v72 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v219 < v220 {
		v67 = v210
		v69 = v212
		v72 = v219
		goto L18
	} else {
		goto L52
	}
L32:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v163 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117)+6)))
	if v163 < v162 {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	if v118&int32(2) == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v15 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v130 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v131 <= v130 {
		v210 = v67
		v212 = v69
		goto L31
	} else {
		goto L40
	}
L38:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	if v125 != int64(-1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+16)) = v128
	goto L37
L40:
	;
	v136 = v130
	v137 = v67
	goto L41
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v137<<(uint(int32(2))%32)))) = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v137))) = uint8(v153)
	v156 = v137 + v153
	v158 = v136 + v153
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v158 < v159 {
		v136 = v158
		v137 = v156
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v210 = v156
	v212 = v69
	goto L31
L43:
	;
	goto L42
L44:
	;
	F_slot_getsomeattrs_int(m, v117, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v167 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v169 <= v167 {
		v210 = v67
		v212 = v167
		goto L31
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v174 = v167
	v175 = v67
	goto L49
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v184 = int32(2)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+16))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+v174<<(uint(v184)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v183+v175<<(uint(v184)%32)))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v174))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194+v175))) = uint8(v199)
	v201 = int32(1)
	v202 = v175 + v201
	v204 = v174 + v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v204 < v205 {
		v174 = v204
		v175 = v202
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v210 = v202
	v212 = v167
	goto L31
L51:
	;
	goto L50
L52:
	;
	goto L19
L53:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v225<<(uint(int32(2))%32)))) = v60
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241+v225))) = uint8(v243)
	goto L55
L54:
	;
	goto L55
L55:
	;
	if v227 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v249 = v247 & int32(_a_F_FunctionNext_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)) = uint16(v252)
	goto L59
L57:
	;
	goto L58
L58:
	;
	return v13
L59:
	;
	goto L58
}
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn13834(m, l0, l1, l2, l3, int32(_a_F_ReceiveFunctionCall_0), int32(1728), int32(_a_F_ReceiveFunctionCall_1), int32(1722), int32(_a_F_ReceiveFunctionCall_2))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_RunFunctionExecuteHook(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RunFunctionExecuteHook[0]))
	m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(4), int32(1255), l0, v4, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_build_function_result_tupdesc_d(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v5
	if base.B2i32(l1 == v5)|base.B2i32(l2 == v5) != 0 {
		v190 = v5
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L59
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L56
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L53
	}
L4:
	;
	m.G0 = v15 + int32(48)
	return v190
L5:
	;
	v24 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v28 != int32(1) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v31 < int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v34 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v35 != int32(26) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v38 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v40 != int32(1) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v43 != v31 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v45 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v46 != int32(18) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if l3 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = F_pg_detoast_datum(m, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v31 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v51 != int32(1) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v54 != v31 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v56 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v57 != int32(25) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_deconstruct_array_builtin(m, v49, int32(25), v15+int32(44), int32(0), v15+int32(40))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v190 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v72 = int32(24)
	v78 = v31 << (uint(int32(2)) % 32)
	v79 = F_palloc(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v81 = F_palloc(m, v78)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v85 = int32(0)
	v86 = int32(0)
	goto L31
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+(v38+v72)))))
	v99 = v97 - int32(105)
	v100 = int32(0)
	if base.B2i32(v99 == v100)|base.B2i32(v99 == int32(13)) == v100 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(l0 == int32(112))|base.B2i32(int32(2) <= v140) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v107 = int32(2)
	v108 = v86 << (uint(v107) % 32)
	v111 = v85 << (uint(v107) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24+v72+v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v79+v108))) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v115 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v140 = v86
	goto L35
L35:
	;
	v145 = v85 + int32(1)
	if v145 != v31 {
		v85 = v145
		v86 = v140
		goto L31
	} else {
		goto L43
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81+v108))) = v137
	v140 = v136
	goto L35
L37:
	;
	v131 = v86 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v131
	v134 = F_psprintf(m, int32(_a_F_build_function_result_tupdesc_d_0), v15)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L42
	}
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v111+v115)))
	v120 = F_text_to_cstring(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v120 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v136 = v86 + int32(1)
	v137 = v120
	goto L36
L42:
	;
	v136 = v131
	v137 = v134
	goto L36
L43:
	;
	goto L32
L44:
	;
	v190 = int32(0)
	goto L4
L45:
	;
	goto L46
L46:
	;
	v155 = F_CreateTemplateTupleDesc(m, v140)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if v140 <= int32(0) {
		v190 = v155
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v161 = int32(0)
	goto L49
L49:
	;
	v173 = v161 << (uint(int32(2)) % 32)
	v175 = v161 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173+v81)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173+v79)))
	F_TupleDescInitEntry(m, v155, base.I32_extend16_s(v175), v178, v180, int32(-1), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	v190 = v155
	goto L4
L51:
	;
	if v175 != v140 {
		v161 = v175
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_1), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_2), int32(1784), int32(_a_F_build_function_result_tupdesc_d_3))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v31
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_4), v15+int32(32))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_2), int32(1792), int32(_a_F_build_function_result_tupdesc_d_3))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v31
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_5), v15+int32(16))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_2), int32(1802), int32(_a_F_build_function_result_tupdesc_d_3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_function_parse_error_transpose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	v2 = int32(0)
	v13 = F_geterrposition(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v277
L2:
	;
	return int32(0)
L3:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = F_getinternalerrposition(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v23 = v13
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_function_parse_error_transpose[0]))
	if v25 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v19 <= int32(0) {
		v277 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = v19
	goto L6
L9:
	;
	v269 = F_errposition(m, v260)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L69
	}
L10:
	;
	v257 = v23
	v260 = int32(0)
	v268 = l0
	goto L9
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	if v28 != int32(3) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v31 = F_strlen(m, l0)
	mBase = m.M
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v33 = F_strlen(m, v32)
	mBase = m.M
	v34 = v33 - v31
	if v34 <= int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v40 = v2
	v41 = v2
	goto L14
L14:
	;
	v51 = v40 + int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v32))))
	switch v52 - int32(36) {
	case 0:
		goto L19
	default:
		v230 = v41
		goto L16
	case 3:
		goto L18
	}
L15:
	;
	v239 = int32(0)
	if v239 < v230 {
		v257 = v239
		v260 = v230
		v268 = v239
		goto L9
	} else {
		goto L68
	}
L16:
	;
	if v51 != v34 {
		v40 = v51
		v41 = v230
		goto L14
	} else {
		goto L67
	}
L17:
	;
	v223 = F_pg_mbstrlen_with_len(m, v32, v51)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L66
	}
L18:
	;
	v107 = v51 + v32
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v108 != 0 {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v55 = v51 + v32
	if v31 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v100 != 0 {
		v230 = v41
		goto L16
	} else {
		goto L33
	}
L21:
	;
	v100 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v61 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = l0
	v63 = v55
	v64 = v31
	v65 = v61
	goto L28
L25:
	;
	v88 = v55
	v92 = int32(0)
	goto L26
L26:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v100 = v92 - v93
	goto L20
L27:
	;
	v88 = v83
	v92 = v85
	goto L26
L28:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if base.B2i32(v65 != v67)|base.B2i32(v67 == int32(0)) != 0 {
		v83 = v63
		v85 = v65
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v83 = v77
	v85 = int32(0)
	goto L27
L30:
	;
	v73 = v64 - int32(1)
	if v73 == int32(0) {
		v83 = v63
		v85 = v65
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v76 = int32(1)
	v77 = v63 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v78 != 0 {
		v62 = v62 + v76
		v63 = v77
		v64 = v73
		v65 = v78
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v31))))
	if v102 != int32(36) {
		v230 = v41
		goto L16
	} else {
		goto L34
	}
L34:
	;
	if v41 == int32(0) {
		v213 = v23
		goto L17
	} else {
		goto L35
	}
L35:
	;
	goto L10
L36:
	;
	v110 = v107
	v111 = v23
	v116 = l0
	v117 = v23
	goto L39
L37:
	;
	v194 = v107
	v195 = v23
	goto L38
L38:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v205 != int32(39) {
		v230 = v41
		goto L16
	} else {
		goto L63
	}
L39:
	;
	v122 = v117 - int32(1)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v123 != int32(39) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v194 = v190
	v195 = v142
	goto L38
L41:
	;
	v143 = F_pg_mblen_cstr(m, v116)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L47
	}
L42:
	;
	if v123 != int32(92) {
		v141 = v110
		v142 = v111
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v133 != int32(39) {
		v230 = v41
		goto L16
	} else {
		goto L46
	}
L45:
	;
	v141 = v110 + int32(1)
	v142 = v111 + base.B2i32(int32(0) < v122)
	goto L41
L46:
	;
	v141 = v110 + int32(1)
	v142 = v111 + base.B2i32(int32(0) < v122)
	goto L41
L47:
	;
	if v143 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v189 != 0 {
		v230 = v41
		goto L16
	} else {
		goto L61
	}
L49:
	;
	v189 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v150 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v151 = v116
	v152 = v141
	v153 = v143
	v154 = v150
	goto L56
L53:
	;
	v177 = v141
	v181 = int32(0)
	goto L54
L54:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v189 = v181 - v182
	goto L48
L55:
	;
	v177 = v172
	v181 = v174
	goto L54
L56:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if base.B2i32(v154 != v156)|base.B2i32(v156 == int32(0)) != 0 {
		v172 = v152
		v174 = v154
		goto L55
	} else {
		goto L58
	}
L57:
	;
	v172 = v166
	v174 = int32(0)
	goto L55
L58:
	;
	v162 = v153 - int32(1)
	if v162 == int32(0) {
		v172 = v152
		v174 = v154
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v165 = int32(1)
	v166 = v152 + v165
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v167 != 0 {
		v151 = v151 + v165
		v152 = v166
		v153 = v162
		v154 = v167
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v190 = v141 + v143
	v191 = v116 + v143
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v192 != 0 {
		v110 = v190
		v111 = v142
		v116 = v191
		v117 = v122
		goto L39
	} else {
		goto L62
	}
L62:
	;
	goto L40
L63:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v208 == int32(39) {
		v230 = v41
		goto L16
	} else {
		goto L64
	}
L64:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v213 = v195
	goto L17
L66:
	;
	v230 = v223 + v213
	goto L16
L67:
	;
	goto L15
L68:
	;
	goto L10
L69:
	;
	F_internalerrposition(m, v257)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v273 = F_internalerrquery(m, v268)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v277 = int32(1)
	goto L1
}
func F_has_function_privilege_name(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13919(m, l0, int32(_a_F_has_function_privilege_name_0), int32(1255), int32(_a_F_has_function_privilege_name_1), int32(3565), int32(_a_F_has_function_privilege_name_2), int32(52461700), int32(1237))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_print_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
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
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int64
	_ = v236
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(128)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v35 = F_get_func_arg_info(m, l1, v24+int32(68), v24-int32(-64), v24+int32(60))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v68 = int32(-1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+96)))
	if v69 == int32(97) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	if l3 == int32(0) {
		v65 = v5
		v66 = v5
		v67 = v35
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+106)))
	if v41 <= int32(0) {
		v65 = v5
		v66 = v5
		v67 = v35
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = F_SysCacheGetAttr(m, int32(47), l1, int32(24), v24+int32(72))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+72)))
	if v50 != 0 {
		v65 = v5
		v66 = v5
		v67 = v35
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v51 = F_text_to_cstring(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v53 = F_stringToNode(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_pfree(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+104)))
	if v53 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v61 = v58
	v62 = v59
	goto L13
L12:
	;
	v61 = v5
	v62 = int32(0)
	goto L13
L13:
	;
	v65 = v53
	v66 = v61
	v67 = v57 - v62
	goto L1
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L76
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v74 = F_SearchSysCache1(m, int32(0), v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	v90 = v68
	goto L17
L17:
	;
	if int32(0) < v35 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	if v74 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
	v80 = v78 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
	if v81 != int32(110) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+6)))
	v85 = v84
	goto L22
L21:
	;
	v85 = v68
	goto L22
L22:
	;
	F_ReleaseCatCache(m, v74)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v90 = v85
	goto L17
L24:
	;
	v96 = v24 + int32(76)
	v99 = int32(0)
	v101 = l3
	v111 = v66
	v113 = v5
	v115 = v5
	goto L27
L25:
	;
	v312 = v5
	goto L26
L26:
	;
	m.G0 = v24 + int32(128)
	return v312
L27:
	;
	v120 = v99 << (uint(int32(2)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	if v124 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v312 = v291
	goto L26
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+v124)))
	v127 = v126
	goto L31
L30:
	;
	v127 = int32(0)
	goto L31
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	if v129 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v295 = v284 + int32(1)
	if v295 < v35 {
		v99 = v295
		v101 = v285
		v111 = v290
		v113 = v291
		v115 = v293
		goto L27
	} else {
		goto L75
	}
L33:
	;
	if v90 == v113 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	if l2 == int32(0) {
		v284 = v99
		v285 = v101
		v290 = v111
		v291 = v113
		v293 = v115
		goto L32
	} else {
		goto L48
	}
L35:
	;
	if l2 == int32(0) {
		v184 = v174
		v185 = v175
		v187 = v177
		goto L33
	} else {
		goto L47
	}
L36:
	;
	v170 = int32(1)
	v174 = v166
	v175 = v170
	v177 = v115 + v170
	goto L35
L37:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+96)))
	if v162 == int32(112) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v133 = int32(0)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v129))))
	switch v136 - int32(98) {
	case 0:
		v166 = int32(_a_F_print_function_arguments_0)
		goto L36
	default:
		goto L39
	case 7:
		goto L37
	case 13:
		v174 = int32(_a_F_print_function_arguments_1)
		v175 = v133
		v177 = v115
		goto L35
	case 18:
		goto L34
	case 20:
		goto L40
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L41
	}
L40:
	;
	v166 = int32(_a_F_print_function_arguments_2)
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = base.I32_extend8_s(v136)
	F_errmsg_internal(m, int32(_a_F_print_function_arguments_3), v24+int32(48))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_print_function_arguments_4), int32(3396), int32(_a_F_print_function_arguments_5))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v165 = int32(_a_F_print_function_arguments_6)
	goto L46
L45:
	;
	v165 = int32(_a_F_print_function_arguments_7)
	goto L46
L46:
	;
	v166 = v165
	goto L36
L47:
	;
	v284 = v99
	v285 = v101
	v290 = v111
	v291 = v113
	v293 = v177
	goto L32
L48:
	;
	v184 = int32(_a_F_print_function_arguments_7)
	v185 = v133
	v187 = v115
	goto L33
L49:
	;
	F_appendStringInfoString(m, l0, v184)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L58
	}
L50:
	;
	F_appendStringInfoString(m, l0, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L57
	}
L51:
	;
	v189 = int32(_a_F_print_function_arguments_8)
	if v90 == int32(0) {
		v198 = v189
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v113 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v198 = v189
	goto L50
L56:
	;
	v198 = int32(_a_F_print_function_arguments_9)
	goto L50
L57:
	;
	goto L49
L58:
	;
	if v127 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = F_format_type_be(m, v128)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L64
	}
L60:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v206 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v209 = F_quote_identifier(m, v127)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v209
	F_appendStringInfo(m, l0, int32(_a_F_print_function_arguments_10), v24+int32(32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_appendStringInfoString(m, l0, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v222 = int32(0)
	if base.B2i32(v101&v185 == v222)|base.B2i32(v187 <= v67) == v222 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v232 = v24 + int32(112)
	F_initStringInfo(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	v274 = v111
	goto L68
L68:
	;
	v276 = v113 + int32(1)
	v279 = base.B2i32(v276 == v90) & base.B2i32(v99 == v35-int32(1))
	v284 = v99 - v279
	v285 = base.B2i32(v279 == int32(0)) & v101
	v290 = v274
	v291 = v276
	v293 = v187
	goto L32
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = v232
	v236 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+21)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v236
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+108)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+107)) = uint8(v244)
	v248 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+105)) = uint16(v248)
	F_get_rule_expr(m, v230, v24+int32(72), v244)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v24)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v255
	F_appendStringInfo(m, l0, int32(_a_F_print_function_arguments_11), v24+int32(16))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v263 = v111 + int32(4)
	if base.Ui32(v263) < base.Ui32(v228+v229<<(uint(int32(2))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v269 = v263
	goto L74
L73:
	;
	v269 = int32(0)
	goto L74
L74:
	;
	v274 = v269
	goto L68
L75:
	;
	goto L28
L76:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v326
	F_errmsg_internal(m, int32(_a_F_print_function_arguments_12), v24)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_print_function_arguments_4), int32(3348), int32(_a_F_print_function_arguments_5))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_print_function_sqlbody(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	base.MemoryFill(m, v9+int32(20), int32(0), int32(68))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v21 = F_pstrdup(m, v16+v17+int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v21
	v30 = F_get_func_arg_info(m, l1, v9+int32(108), v9+int32(104), v9+int32(100))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v33
	v37 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(28))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v9 + int32(112)
	return
L5:
	;
	v39 = F_text_to_cstring(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = F_stringToNode(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v43 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	F_appendStringInfoString(m, l0, int32(_a_F_print_function_sqlbody_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v107 = int32(0)
	F_AcquireRewriteLocks(m, v41, v107, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	if v47 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_print_function_sqlbody_1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v53 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v58 = int32(0)
	goto L15
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v58<<(uint(int32(2))%32))))
	v68 = int32(0)
	F_AcquireRewriteLocks(m, v67, v68, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v73 = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
	v79 = F_list_make1_impl(m, int32(1), v9+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(0)
	F_get_query_def(m, v67, l0, v79, v81, v81, int32(2), v81, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_appendStringInfoChar(m, l0, int32(59))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoChar(m, l0, int32(10))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v95 = v58 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v95 < v96 {
		v58 = v95
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	goto L4
L24:
	;
	v112 = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v112
	v118 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v120 = int32(0)
	F_get_query_def(m, v41, l0, v118, v120, v120, v120, v120, v120)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L4
}
