package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall1Coll(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
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
	var v23 int64
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
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+40)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	v18 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+26)) = uint16(v18)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = m.T0[v22].(func(*base.Module, int32) int64)(m, v7+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)))
		if v27 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
				F_errmsg_internal(m, int32(_a_F_FunctionCall1Coll_0), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall1Coll_1), int32(1145), int32(_a_F_FunctionCall1Coll_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(48)
			return v23
		}
	}
}
func F_FunctionCall3Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64) int64 {
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
	var v31 int64
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
	v9 = v7 - int32(80)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+72)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = l2
	v20 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+26)) = uint16(v20)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l0
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = m.T0[v30].(func(*base.Module, int32) int64)(m, v9+int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int64(0)
	} else {
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
		if v35 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v42
				F_errmsg_internal(m, int32(_a_F_FunctionCall3Coll_0), v9)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall3Coll_1), int32(1192), int32(_a_F_FunctionCall3Coll_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(80)
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
	var v14 int32
	_ = v14
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v204
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v204 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(_a_F_FunctionIsVisibleExt_0), v10)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_FunctionIsVisibleExt_1), int32(1770), int32(_a_F_FunctionIsVisibleExt_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L60
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_FunctionIsVisibleExt[0]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+104)))
	v89 = F_makeString(m, v39+int32(4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v195 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v77 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v77
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v77 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v77 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v89
	v96 = F_list_make1_impl(m, int32(1), v10+int32(4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v98 = int32(0)
	v105 = F_FuncnameGetCandidates(m, v96, v86, v98, v98, v98, v98, v98, v10+int32(12))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v105 == int32(0) {
		v195 = v3
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v110 = v86 << (uint(int32(2)) % 32)
	v112 = v39 + int32(136)
	v114 = v105
	goto L36
L36:
	;
	v121 = v114 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v110) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v195 = v3
	goto L14
L38:
	;
	if v183 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v183 = int32(0)
	goto L38
L40:
	;
	v157 = v152
	v158 = v153
	v159 = v154
	goto L50
L41:
	;
	if (v121|v112)&int32(3) != 0 {
		v152 = v121
		v153 = v112
		v154 = v110
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v145 = v121
	v146 = v112
	v147 = v110
	goto L43
L43:
	;
	if v147 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v129 = v121
	v130 = v112
	v131 = v110
	goto L45
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v134 != v135 {
		v152 = v129
		v153 = v130
		v154 = v131
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v145 = v140
	v146 = v138
	v147 = v142
	goto L43
L47:
	;
	v137 = int32(4)
	v138 = v130 + v137
	v140 = v129 + v137
	v142 = v131 - v137
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		v129 = v140
		v130 = v138
		v131 = v142
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v152 = v145
	v153 = v146
	v154 = v147
	goto L40
L50:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v162 == v163 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v183 = v162 - v163
	goto L38
L52:
	;
	v165 = int32(1)
	v170 = v159 - v165
	if v170 != 0 {
		v157 = v157 + v165
		v158 = v158 + v165
		v159 = v170
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
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v195 = base.B2i32(v186 == l0)
	goto L14
L57:
	;
	goto L58
L58:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v188 != 0 {
		v114 = v188
		goto L36
	} else {
		goto L59
	}
L59:
	;
	goto L37
L60:
	;
	v204 = v195
	goto L1
}
func F_FunctionNext(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
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
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	v2 = int32(0)
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+121)))
	if v15 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v14 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v31 = F_ExecMakeTableFunctionResult(m, v22, v23, v24, v25, int32(base.Ui32(v26&int32(8))>>(uint(int32(3))%32)))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v39 = v19
	goto L6
L6:
	;
	v43 = F_tuplestore_gettupleslot(m, v39, base.B2i32(v14 == int32(1)), int32(0), v12)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v31
	F_tuplestore_rescan(m, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = v31
	goto L6
L10:
	;
	return v12
L11:
	;
	v51 = int64(1)
	goto L13
L12:
	;
	v51 = int64(-1)
	goto L13
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+128)) = v46 + v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	m.T0[v55].(func(*base.Module, int32))(m, v12)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = v2
	v66 = v11
	v68 = v2
	goto L18
L16:
	;
	v220 = v2
	v222 = v11
	goto L17
L17:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v227 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v74 = v71 + v68<<(uint(int32(5))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v75 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v220 = v206
	v222 = v208
	goto L17
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v87 = F_ExecMakeTableFunctionResult(m, v78, v79, v80, v81, int32(base.Ui32(v82&int32(8))>>(uint(int32(3))%32)))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v74)+16))
	if base.B2i32(v93 == int64(-1))|base.B2i32(v46 <= v93) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v87
	F_tuplestore_rescan(m, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v113 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	m.T0[v102].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v110 = F_tuplestore_gettupleslot(m, v105, base.B2i32(v14 == int32(1)), int32(0), v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	v214 = v68 + int32(1)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v214 < v215 {
		v64 = v206
		v66 = v208
		v68 = v214
		goto L18
	} else {
		goto L52
	}
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+6)))
	if v158 < v157 {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+4)))
	if v114&int32(2) == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v14 != int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v126 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v127 <= v126 {
		v206 = v64
		v208 = v66
		goto L31
	} else {
		goto L40
	}
L38:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v74)+16))
	if v121 != int64(-1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+16)) = v124
	goto L37
L40:
	;
	v132 = v126
	v133 = v64
	goto L41
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v140+v133<<(uint(int32(3))%32)))) = int64(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v148 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v146+v133))) = uint8(v148)
	v151 = v133 + v148
	v153 = v132 + v148
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v153 < v154 {
		v132 = v153
		v133 = v151
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v206 = v151
	v208 = v66
	goto L31
L43:
	;
	goto L42
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	m.T0[v161].(func(*base.Module, int32, int32))(m, v113, v157)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v164 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v166 <= v164 {
		v206 = v64
		v208 = v164
		goto L31
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v171 = v164
	v172 = v64
	goto L49
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v180 = int32(3)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v184+v171<<(uint(v180)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v179+v172<<(uint(v180)%32)))) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v171))))
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v172))) = uint8(v195)
	v197 = int32(1)
	v198 = v172 + v197
	v200 = v171 + v197
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v200 < v201 {
		v171 = v200
		v172 = v198
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v206 = v198
	v208 = v164
	goto L31
L51:
	;
	goto L50
L52:
	;
	goto L19
L53:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v234 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v230+v220<<(uint(int32(3))%32)))) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v236+v220))) = uint8(v238)
	goto L55
L54:
	;
	goto L55
L55:
	;
	if v222 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	v244 = v242 & int32(_a_F_FunctionNext_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v244)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)) = uint16(v247)
	goto L59
L57:
	;
	goto L58
L58:
	;
	return v12
L59:
	;
	goto L58
}
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14217(m, l0, l1, l2, l3, int32(_a_F_ReceiveFunctionCall_0), int32(1729), int32(_a_F_ReceiveFunctionCall_1), int32(1723), int32(_a_F_ReceiveFunctionCall_2))
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
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
func F_build_function_result_tupdesc_d(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v5
	v21 = int64(0)
	if base.B2i32(l1 == v21)|base.B2i32(l2 == v21) != 0 {
		v305 = v5
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L80
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L77
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L74
	}
L4:
	;
	m.G0 = v17 + int32(48)
	return v305
L5:
	;
	v27 = F_pg_detoast_datum(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v31 != int32(1) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v34 < int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v37 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v38 != int32(26) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v42 = F_pg_detoast_datum(m, base.I32_wrap_i64(l2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v44 != int32(1) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v47 != v34 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v49 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v50 != int32(18) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if l3 != int64(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v56 = F_pg_detoast_datum(m, base.I32_wrap_i64(l3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v34 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v58 != int32(1) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v61 != v34 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v64 != int32(25) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_deconstruct_array_builtin(m, v56, int32(25), v17+int32(44), int32(0), v17+int32(40))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v305 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v79 = int32(24)
	v85 = v34 << (uint(int32(2)) % 32)
	v86 = F_palloc(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v88 = F_palloc(m, v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v95 = int32(0)
	v96 = int32(0)
	goto L31
L31:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+(v42+v79)))))
	v108 = v106 - int32(105)
	v109 = int32(0)
	if base.B2i32(v108 == v109)|base.B2i32(v108 == int32(13)) == v109 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(l0 == int32(112))|base.B2i32(int32(2) <= v151) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v116 = int32(2)
	v117 = v95 << (uint(v116) % 32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v27+v79+v96<<(uint(v116)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v86+v117))) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v124 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v151 = v95
	goto L35
L35:
	;
	v155 = v96 + int32(1)
	if v155 != v34 {
		v95 = v151
		v96 = v155
		goto L31
	} else {
		goto L43
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88+v117))) = v148
	v151 = v147
	goto L35
L37:
	;
	v142 = v95 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v142
	v145 = F_psprintf(m, int32(_a_F_build_function_result_tupdesc_d_0), v17)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L42
	}
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+v96<<(uint(int32(3))%32))))
	v131 = F_text_to_cstring(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v131 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v135 == int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v147 = v95 + int32(1)
	v148 = v131
	goto L36
L42:
	;
	v147 = v142
	v148 = v145
	goto L36
L43:
	;
	goto L32
L44:
	;
	v305 = int32(0)
	goto L4
L45:
	;
	goto L46
L46:
	;
	v165 = F_CreateTemplateTupleDesc(m, v151)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if int32(0) < v151 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v212 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v212 < v221 {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v185 = v175 << (uint(int32(2)) % 32)
	v187 = v175 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185+v88)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185+v86)))
	F_TupleDescInitEntry(m, v165, base.I32_extend16_s(v187), v190, v192, int32(-1), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	if v151 != v187 {
		v175 = v187
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v305 = v165
	goto L4
L56:
	;
	v225 = v165 + int32(28)
	v232 = v212
	v233 = v221
	v235 = v212
	goto L60
L57:
	;
	v289 = v212
	v296 = v221
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v289
	goto L55
L59:
	;
	v289 = v283
	v296 = v262
	goto L58
L60:
	;
	v241 = v225 + v221<<(uint(int32(3))%32) + v232*int32(100)
	v244 = v225 + v232<<(uint(int32(3))%32)
	if v221 != v233 {
		v262 = v233
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v283 = v221
	goto L59
L62:
	;
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244)+2)))
	if v263 <= int32(0) {
		v283 = v232
		goto L59
	} else {
		goto L70
	}
L63:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+7)))
	if v246 != int32(118) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v262 = v232
	goto L62
L65:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+4)))
	if v249 != int32(1) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+6)))
	if v252&int32(6) != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244)+2)))
	if v255 <= int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+90)))
	if v258 != int32(118) {
		v262 = v221
		goto L62
	} else {
		goto L69
	}
L69:
	;
	goto L64
L70:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+90)))
	if v266 == int32(118) {
		v283 = v232
		goto L59
	} else {
		goto L71
	}
L71:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+5)))
	v275 = (v235 + v269 - int32(1)) & (int32(0) - v269)
	if int32(_a_F_build_function_result_tupdesc_d_1) < v275 {
		v283 = v232
		goto L59
	} else {
		goto L72
	}
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v244))) = uint16(v275)
	v281 = v232 + int32(1)
	if v281 != v221 {
		v232 = v281
		v233 = v262
		v235 = v275 + v263
		goto L60
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_2), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_3), int32(1787), int32(_a_F_build_function_result_tupdesc_d_4))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L6
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v34
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_5), v17+int32(32))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_3), int32(1795), int32(_a_F_build_function_result_tupdesc_d_4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v34
	F_errmsg_internal(m, int32(_a_F_build_function_result_tupdesc_d_6), v17+int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_build_function_result_tupdesc_d_3), int32(1805), int32(_a_F_build_function_result_tupdesc_d_4))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
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
func F_has_function_privilege_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14306(m, l0, int32(_a_F_has_function_privilege_name_0), int32(1255), int32(_a_F_has_function_privilege_name_1), int32(3589), int32(_a_F_has_function_privilege_name_2), int32(52461700), int32(1365))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_print_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int64
	_ = v239
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v315 int32
	_ = v315
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	v5 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(128)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v36 = F_get_func_arg_info(m, l1, v25+int32(68), v25-int32(-64), v25+int32(60))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = int32(-1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+96)))
	if v71 == int32(97) {
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
		v67 = v5
		v68 = v5
		v69 = v36
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+106)))
	if v42 <= int32(0) {
		v67 = v5
		v68 = v5
		v69 = v36
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v49 = F_SysCacheGetAttr(m, int32(47), l1, int32(24), v25+int32(72))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+72)))
	if v51 != 0 {
		v67 = v5
		v68 = v5
		v69 = v36
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v53 = F_text_to_cstring(m, base.I32_wrap_i64(v49))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v55 = F_stringToNode(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_pfree(m, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+104)))
	if v55 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v63 = v60
	v64 = v61
	goto L13
L12:
	;
	v63 = v5
	v64 = int32(0)
	goto L13
L13:
	;
	v67 = v55
	v68 = v63
	v69 = v59 - v64
	goto L1
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L76
	}
L15:
	;
	v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v29))))
	v76 = F_SearchSysCache1(m, int32(0), v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	v92 = v70
	goto L17
L17:
	;
	if int32(0) < v36 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	if v76 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v82 = v80 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v83 != int32(110) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+6)))
	v87 = v86
	goto L22
L21:
	;
	v87 = v70
	goto L22
L22:
	;
	F_ReleaseCatCache(m, v76)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v92 = v87
	goto L17
L24:
	;
	v98 = v25 + int32(76)
	v101 = int32(0)
	v103 = l3
	v113 = v68
	v115 = v5
	v117 = v5
	goto L27
L25:
	;
	v315 = v5
	goto L26
L26:
	;
	m.G0 = v25 + int32(128)
	return v315
L27:
	;
	v123 = v101 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	if v127 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v315 = v294
	goto L26
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+v127)))
	v130 = v129
	goto L31
L30:
	;
	v130 = int32(0)
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v123+v124)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
	if v132 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v298 = v287 + int32(1)
	if v298 < v36 {
		v101 = v298
		v103 = v288
		v113 = v293
		v115 = v294
		v117 = v296
		goto L27
	} else {
		goto L75
	}
L33:
	;
	if v92 == v115 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	if l2 == int32(0) {
		v287 = v101
		v288 = v103
		v293 = v113
		v294 = v115
		v296 = v117
		goto L32
	} else {
		goto L48
	}
L35:
	;
	if l2 == int32(0) {
		v187 = v177
		v188 = v178
		v190 = v180
		goto L33
	} else {
		goto L47
	}
L36:
	;
	v173 = int32(1)
	v177 = v169
	v178 = v173
	v180 = v117 + v173
	goto L35
L37:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+96)))
	if v165 == int32(112) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v136 = int32(0)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v132))))
	switch v139 - int32(98) {
	case 0:
		v169 = int32(_a_F_print_function_arguments_0)
		goto L36
	default:
		goto L39
	case 7:
		goto L37
	case 13:
		v177 = int32(_a_F_print_function_arguments_1)
		v178 = v136
		v180 = v117
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
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L41
	}
L40:
	;
	v169 = int32(_a_F_print_function_arguments_2)
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = base.I32_extend8_s(v139)
	F_errmsg_internal(m, int32(_a_F_print_function_arguments_3), v25+int32(48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_print_function_arguments_4), int32(3402), int32(_a_F_print_function_arguments_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v168 = int32(_a_F_print_function_arguments_6)
	goto L46
L45:
	;
	v168 = int32(_a_F_print_function_arguments_7)
	goto L46
L46:
	;
	v169 = v168
	goto L36
L47:
	;
	v287 = v101
	v288 = v103
	v293 = v113
	v294 = v115
	v296 = v180
	goto L32
L48:
	;
	v187 = int32(_a_F_print_function_arguments_7)
	v188 = v136
	v190 = v117
	goto L33
L49:
	;
	F_appendStringInfoString(m, l0, v187)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L58
	}
L50:
	;
	F_appendStringInfoString(m, l0, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L57
	}
L51:
	;
	v192 = int32(_a_F_print_function_arguments_8)
	if v92 == int32(0) {
		v201 = v192
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v115 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v201 = v192
	goto L50
L56:
	;
	v201 = int32(_a_F_print_function_arguments_9)
	goto L50
L57:
	;
	goto L49
L58:
	;
	if v130 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v220 = F_format_type_be(m, v131)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L64
	}
L60:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v209 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v212 = F_quote_identifier(m, v130)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v212
	F_appendStringInfo(m, l0, int32(_a_F_print_function_arguments_10), v25+int32(32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_appendStringInfoString(m, l0, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v225 = int32(0)
	if base.B2i32(v103&v188 == v225)|base.B2i32(v190 <= v69) == v225 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v235 = v25 + int32(112)
	F_initStringInfo(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	v277 = v113
	goto L68
L68:
	;
	v279 = v115 + int32(1)
	v282 = base.B2i32(v279 == v92) & base.B2i32(v101 == v36-int32(1))
	v287 = v101 - v282
	v288 = base.B2i32(v282 == int32(0)) & v103
	v293 = v277
	v294 = v279
	v296 = v190
	goto L32
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = v235
	v239 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+21)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v98))) = v239
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = v247
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+107)) = uint8(v247)
	v251 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+105)) = uint16(v251)
	F_get_rule_expr(m, v233, v25+int32(72), v247)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v25)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v258
	F_appendStringInfo(m, l0, int32(_a_F_print_function_arguments_11), v25+int32(16))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v266 = v113 + int32(4)
	if base.Ui32(v266) < base.Ui32(v231+v232<<(uint(int32(2))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v272 = v266
	goto L74
L73:
	;
	v272 = int32(0)
	goto L74
L74:
	;
	v277 = v272
	goto L68
L75:
	;
	goto L28
L76:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v330
	F_errmsg_internal(m, int32(_a_F_print_function_arguments_12), v25)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_print_function_arguments_4), int32(3354), int32(_a_F_print_function_arguments_5))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
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
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
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
	v40 = F_text_to_cstring(m, base.I32_wrap_i64(v37))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v42 = F_stringToNode(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	F_appendStringInfoString(m, l0, int32(_a_F_print_function_sqlbody_0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v108 = int32(0)
	F_AcquireRewriteLocks(m, v42, v108, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_print_function_sqlbody_1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v54 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v59 = int32(0)
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v59<<(uint(int32(2))%32))))
	v69 = int32(0)
	F_AcquireRewriteLocks(m, v68, v69, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v74 = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v74
	v80 = F_list_make1_impl(m, int32(1), v9+int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v82 = int32(0)
	F_get_query_def(m, v68, l0, v80, v82, v82, int32(2), v82, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_appendStringInfoChar(m, l0, int32(59))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoChar(m, l0, int32(10))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v96 = v59 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v96 < v97 {
		v59 = v96
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
	v113 = v9 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v113
	v119 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v121 = int32(0)
	F_get_query_def(m, v42, l0, v119, v121, v121, v121, v121, v121)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L4
}
