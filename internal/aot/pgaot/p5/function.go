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
				F_errmsg_internal(m, int32(524358), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487551), int32(1143), int32(299660))
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
				F_errmsg_internal(m, int32(524358), v9)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487551), int32(1190), int32(299612))
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
	var v75 int32
	_ = v75
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
	F_errmsg_internal(m, int32(44278), v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(491828), int32(1723), int32(63795))
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[250]))
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
		v75 = v44
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v82 = v75
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
		v75 = v68
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v75 = v68
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	if v61 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v230 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L16:
	;
	v221 = v2
	v224 = v12
	goto L15
L17:
	;
	goto L18
L18:
	;
	v66 = v2
	v69 = v12
	v71 = v2
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v78 = v75 + v71<<(uint(int32(5))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v79 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v221 = v206
	v224 = v209
	goto L15
L21:
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
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	if v97 == int64(-1) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v91
	F_tuplestore_rescan(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v111 = F_tuplestore_gettupleslot(m, v106, base.B2i32(v15 == int32(1)), int32(0), v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L31
	}
L28:
	;
	if v47 <= v97 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	m.T0[v103].(func(*base.Module, int32))(m, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	goto L26
L32:
	;
	v216 = v71 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v216 < v217 {
		v66 = v206
		v69 = v209
		v71 = v216
		goto L19
	} else {
		goto L53
	}
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+6)))
	if v160 < v159 {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
	if v115&int32(2) == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v15 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v127 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v128 <= v127 {
		v206 = v66
		v209 = v69
		goto L32
	} else {
		goto L41
	}
L39:
	;
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	if v122 != int64(-1) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+16)) = v125
	goto L38
L41:
	;
	v133 = v66
	v134 = v127
	goto L42
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v133<<(uint(int32(2))%32)))) = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v148+v133))) = uint8(v150)
	v153 = v133 + v150
	v155 = v134 + v150
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v155 < v156 {
		v133 = v153
		v134 = v155
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v206 = v153
	v209 = v69
	goto L32
L44:
	;
	goto L43
L45:
	;
	F_slot_getsomeattrs_int(m, v114, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v164 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v166 <= v164 {
		v206 = v66
		v209 = v164
		goto L32
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v171 = v66
	v172 = v164
	goto L50
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v181 = int32(2)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v172<<(uint(v181)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v180+v171<<(uint(v181)%32)))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v171))) = uint8(v196)
	v198 = int32(1)
	v199 = v171 + v198
	v201 = v172 + v198
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v201 < v202 {
		v171 = v199
		v172 = v201
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v206 = v199
	v209 = v164
	goto L32
L52:
	;
	goto L51
L53:
	;
	goto L20
L54:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v233+v221<<(uint(int32(2))%32)))) = v60
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v238+v221))) = uint8(v240)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if v224&int32(1) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v248 = v246 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v248)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+6)) = uint16(v251)
	goto L60
L58:
	;
	goto L59
L59:
	;
	return v13
L60:
	;
	goto L59
}
func F_ReceiveFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if l1 == v5 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v13 != 0 {
			v77 = v5
			m.G0 = v9 - int32(-64)
			return v77
		} else {
			v14 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+29)) = v14
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v14
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+60)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+52)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
			v27 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+38)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v7+int32(-44))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
				if l1 == int32(0) {
					if v37&int32(1) != 0 {
						v77 = v33
						m.G0 = v9 - int32(-64)
						return v77
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
							F_errmsg_internal(m, int32(524233), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(487551), int32(1722), int32(300693))
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
					}
				} else {
					if v37&int32(1) == int32(0) {
						v77 = v33
						m.G0 = v9 - int32(-64)
						return v77
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v64
							F_errmsg_internal(m, int32(524350), v7+int32(-48))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(487551), int32(1728), int32(300693))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
		}
	} else {
		v14 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+29)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v14
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+60)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+52)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
		v27 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+38)) = uint16(v27)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v7+int32(-44))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
			if l1 == int32(0) {
				if v37&int32(1) != 0 {
					v77 = v33
					m.G0 = v9 - int32(-64)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
						F_errmsg_internal(m, int32(524233), v9)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487551), int32(1722), int32(300693))
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
				}
			} else {
				if v37&int32(1) == int32(0) {
					v77 = v33
					m.G0 = v9 - int32(-64)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v64
						F_errmsg_internal(m, int32(524350), v7+int32(-48))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487551), int32(1728), int32(300693))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v5
	if l1 == v5 {
		v179 = v5
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L59
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L56
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L53
	}
L4:
	;
	m.G0 = v15 + int32(48)
	return v179
L5:
	;
	if l2 == int32(0) {
		v179 = v5
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 != int32(1) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v30 < int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v33 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v34 != int32(26) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v37 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v39 != int32(1) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v42 != v30 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v44 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v45 != int32(18) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = F_pg_detoast_datum(m, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v30 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v50 != int32(1) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v53 != v30 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v55 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v56 != int32(25) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_deconstruct_array_builtin(m, v48, int32(25), v15+int32(44), int32(0), v15+int32(40))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	v179 = int32(0)
	goto L4
L28:
	;
	goto L29
L29:
	;
	v71 = int32(24)
	v77 = v30 << (uint(int32(2)) % 32)
	v78 = F_palloc(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v80 = F_palloc(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v84 = int32(0)
	v85 = int32(0)
	goto L32
L32:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+(v37+v71)))))
	switch v96 - int32(105) {
	case 0, 13:
		v132 = v85
		goto L34
	default:
		goto L35
	}
L33:
	;
	if l0 == int32(112) {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v137 = v84 + int32(1)
	if v137 != v30 {
		v84 = v137
		v85 = v132
		goto L32
	} else {
		goto L43
	}
L35:
	;
	v99 = int32(2)
	v100 = v85 << (uint(v99) % 32)
	v103 = v84 << (uint(v99) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v23+v71+v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v78+v100))) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v107 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80+v100))) = v129
	v132 = v128
	goto L34
L37:
	;
	v123 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v123
	v126 = F_psprintf(m, int32(458919), v15)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L42
	}
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v103+v107)))
	v112 = F_text_to_cstring(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v112 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 == int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v128 = v85 + int32(1)
	v129 = v112
	goto L36
L42:
	;
	v128 = v123
	v129 = v126
	goto L36
L43:
	;
	goto L33
L44:
	;
	v144 = F_CreateTemplateTupleDesc(m, v132)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	if int32(2) <= v132 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v179 = int32(0)
	goto L4
L47:
	;
	if v132 <= int32(0) {
		v179 = v144
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v150 = int32(0)
	goto L49
L49:
	;
	v162 = v150 << (uint(int32(2)) % 32)
	v164 = v150 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162+v80)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162+v78)))
	F_TupleDescInitEntry(m, v144, base.I32_extend16_s(v164), v167, v169, int32(-1), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L51
	}
L50:
	;
	v179 = v144
	goto L4
L51:
	;
	if v164 != v132 {
		v150 = v164
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	F_errmsg_internal(m, int32(149859), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(490081), int32(1784), int32(457897))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v30
	F_errmsg_internal(m, int32(150044), v15+int32(32))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(490081), int32(1792), int32(457897))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v30
	F_errmsg_internal(m, int32(149974), v15+int32(16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(490081), int32(1802), int32(457897))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L7
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
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
	return v387
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	if v25 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v19 <= int32(0) {
		v387 = v2
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
	v379 = F_errposition(m, v369)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L2
	} else {
		goto L105
	}
L10:
	;
	v367 = v23
	v369 = int32(0)
	v378 = l0
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
	if l0&int32(3) == int32(0) {
		v54 = l0
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v88&int32(3) == int32(0) {
		v112 = v88
		goto L32
	} else {
		goto L33
	}
L14:
	;
	v87 = v79 - l0
	goto L13
L15:
	;
	v58 = v54
	goto L24
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v38 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v87 = int32(0)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v43 = l0
	goto L20
L20:
	;
	v47 = v43 + int32(1)
	if v47&int32(3) == int32(0) {
		v54 = v47
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v79 = v47
	goto L14
L22:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v52 != 0 {
		v43 = v47
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v67 = int32(-2139062144)
	if (int32(16843008)-v64|v64)&v67 == v67 {
		v58 = v58 + int32(4)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v73 = v58
	goto L27
L26:
	;
	goto L25
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != 0 {
		v73 = v73 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v79 = v73
	goto L14
L29:
	;
	goto L28
L30:
	;
	v146 = v145 - v87
	if v146 <= int32(0) {
		goto L10
	} else {
		goto L47
	}
L31:
	;
	v145 = v137 - v88
	goto L30
L32:
	;
	v116 = v112
	goto L41
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v96 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v145 = int32(0)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v101 = v88
	goto L37
L37:
	;
	v105 = v101 + int32(1)
	if v105&int32(3) == int32(0) {
		v112 = v105
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v137 = v105
	goto L31
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v110 != 0 {
		v101 = v105
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v125 = int32(-2139062144)
	if (int32(16843008)-v122|v122)&v125 == v125 {
		v116 = v116 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v131 = v116
	goto L44
L43:
	;
	goto L42
L44:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v135 != 0 {
		v131 = v131 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v137 = v131
	goto L31
L46:
	;
	goto L45
L47:
	;
	v151 = v2
	v152 = v2
	goto L48
L48:
	;
	v163 = v151 + int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v88))))
	switch v164 - int32(36) {
	case 0:
		goto L53
	default:
		v339 = v152
		goto L50
	case 3:
		goto L52
	}
L49:
	;
	v349 = int32(0)
	if v349 < v339 {
		v367 = v349
		v369 = v339
		v378 = v349
		goto L9
	} else {
		goto L104
	}
L50:
	;
	if v163 != v146 {
		v151 = v163
		v152 = v339
		goto L48
	} else {
		goto L103
	}
L51:
	;
	v333 = F_pg_mbstrlen_with_len(m, v88, v163)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L102
	}
L52:
	;
	v218 = v163 + v88
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v219 != 0 {
		goto L71
	} else {
		goto L72
	}
L53:
	;
	v167 = v163 + v88
	if v87 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v211 != 0 {
		v339 = v152
		goto L50
	} else {
		goto L68
	}
L55:
	;
	v211 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v173 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v174 = l0
	v175 = v167
	v176 = v87
	v177 = v173
	goto L62
L59:
	;
	v199 = v167
	v203 = int32(0)
	goto L60
L60:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v211 = v203 - v204
	goto L54
L61:
	;
	v199 = v194
	v203 = v196
	goto L60
L62:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v177 != v179 {
		v194 = v175
		v196 = v177
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v194 = v188
	v196 = int32(0)
	goto L61
L64:
	;
	if v179 == int32(0) {
		v194 = v175
		v196 = v177
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v184 = v176 - int32(1)
	if v184 == int32(0) {
		v194 = v175
		v196 = v177
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v187 = int32(1)
	v188 = v175 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v189 != 0 {
		v174 = v174 + v187
		v175 = v188
		v176 = v184
		v177 = v189
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v87))))
	if v213 != int32(36) {
		v339 = v152
		goto L50
	} else {
		goto L69
	}
L69:
	;
	if v152 == int32(0) {
		v325 = v23
		goto L51
	} else {
		goto L70
	}
L70:
	;
	goto L10
L71:
	;
	v221 = v218
	v224 = v23
	v225 = l0
	v228 = v23
	goto L74
L72:
	;
	v304 = v218
	v307 = v23
	goto L73
L73:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v315 != int32(39) {
		v339 = v152
		goto L50
	} else {
		goto L99
	}
L74:
	;
	v233 = v228 - int32(1)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v234 != int32(39) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v304 = v300
	v307 = v253
	goto L73
L76:
	;
	v254 = F_pg_mblen_cstr(m, v225)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L82
	}
L77:
	;
	if v234 != int32(92) {
		v252 = v221
		v253 = v224
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v245 = v221 + int32(1)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v246 != int32(39) {
		v339 = v152
		goto L50
	} else {
		goto L81
	}
L80:
	;
	v252 = v221 + int32(1)
	v253 = v224 + base.B2i32(int32(0) < v233)
	goto L76
L81:
	;
	v252 = v245
	v253 = v224 + base.B2i32(int32(0) < v233)
	goto L76
L82:
	;
	if v254 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v299 != 0 {
		v339 = v152
		goto L50
	} else {
		goto L97
	}
L84:
	;
	v299 = int32(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v261 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v262 = v225
	v263 = v252
	v264 = v254
	v265 = v261
	goto L91
L88:
	;
	v287 = v252
	v291 = int32(0)
	goto L89
L89:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v299 = v291 - v292
	goto L83
L90:
	;
	v287 = v282
	v291 = v284
	goto L89
L91:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v265 != v267 {
		v282 = v263
		v284 = v265
		goto L90
	} else {
		goto L93
	}
L92:
	;
	v282 = v276
	v284 = int32(0)
	goto L90
L93:
	;
	if v267 == int32(0) {
		v282 = v263
		v284 = v265
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v272 = v264 - int32(1)
	if v272 == int32(0) {
		v282 = v263
		v284 = v265
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v275 = int32(1)
	v276 = v263 + v275
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if v277 != 0 {
		v262 = v262 + v275
		v263 = v276
		v264 = v272
		v265 = v277
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	v300 = v252 + v254
	v301 = v225 + v254
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v302 != 0 {
		v221 = v300
		v224 = v253
		v225 = v301
		v228 = v233
		goto L74
	} else {
		goto L98
	}
L98:
	;
	goto L75
L99:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+1)))
	if v318 == int32(39) {
		v339 = v152
		goto L50
	} else {
		goto L100
	}
L100:
	;
	if v152 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v325 = v307
	goto L51
L102:
	;
	v339 = v333 + v325
	goto L50
L103:
	;
	goto L49
L104:
	;
	goto L10
L105:
	;
	F_internalerrposition(m, v367)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v383 = F_internalerrquery(m, v378)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	v387 = int32(1)
	goto L1
}
func F_has_function_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[239]))
			v22 = F_text_to_cstring(m, v11)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_DirectFunctionCall1Coll(m, int32(1252), int32(0), v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
								F_errmsg(m, int32(70476), v8)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(489865), int32(3565), int32(372981))
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
						}
					} else {
						v46 = F_convert_any_priv_string(m, v16, int32(1632528))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = F_object_aclcheck(m, int32(1255), v24, v19, v46)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return base.B2i32(v48 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_print_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v320 int32
	_ = v320
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(128)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v31 = v29 + v30
	v38 = F_get_func_arg_info(m, l1, v27+int32(68), v27-int32(-64), v27+int32(60))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v71 = int32(-1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+96)))
	if v72 == int32(97) {
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
		v68 = v5
		v69 = v5
		v70 = v38
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+106)))
	if v44 <= int32(0) {
		v68 = v5
		v69 = v5
		v70 = v38
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v51 = F_SysCacheGetAttr(m, int32(47), l1, int32(24), v27+int32(72))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v53 != 0 {
		v68 = v5
		v69 = v5
		v70 = v38
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v54 = F_text_to_cstring(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v56 = F_stringToNode(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_pfree(m, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
	if v56 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v64 = v61
	v65 = v62
	goto L13
L12:
	;
	v64 = v5
	v65 = int32(0)
	goto L13
L13:
	;
	v68 = v56
	v69 = v64
	v70 = v60 - v65
	goto L1
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L2
	} else {
		goto L76
	}
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v77 = F_SearchSysCache1(m, int32(0), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	v93 = v71
	goto L17
L17:
	;
	if int32(0) < v38 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+22)))
	v83 = v81 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+4)))
	if v84 != int32(110) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+6)))
	v88 = v87
	goto L22
L21:
	;
	v88 = v71
	goto L22
L22:
	;
	F_ReleaseCatCache(m, v77)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v93 = v88
	goto L17
L24:
	;
	v108 = int32(0)
	v110 = l3
	v120 = v5
	v121 = v69
	v124 = v5
	goto L27
L25:
	;
	v320 = v5
	goto L26
L26:
	;
	m.G0 = v27 + int32(128)
	return v320
L27:
	;
	v132 = v108 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	if v136 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v320 = v300
	goto L26
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+v136)))
	v139 = v138
	goto L31
L30:
	;
	v139 = int32(0)
	goto L31
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v141 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v305 = v294 + int32(1)
	if v305 < v38 {
		v108 = v305
		v110 = v295
		v120 = v300
		v121 = v301
		v124 = v303
		goto L27
	} else {
		goto L75
	}
L33:
	;
	if v93 == v120 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	if l2 == int32(0) {
		v294 = v108
		v295 = v110
		v300 = v120
		v301 = v121
		v303 = v124
		goto L32
	} else {
		goto L48
	}
L35:
	;
	if l2 == int32(0) {
		v196 = v186
		v197 = v187
		v199 = v189
		goto L33
	} else {
		goto L47
	}
L36:
	;
	v182 = int32(1)
	v186 = v178
	v187 = v182
	v189 = v124 + v182
	goto L35
L37:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+96)))
	if v174 == int32(112) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v145 = int32(0)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v141))))
	switch v148 - int32(98) {
	case 0:
		v178 = int32(719934)
		goto L36
	default:
		goto L39
	case 7:
		goto L37
	case 13:
		v186 = int32(719936)
		v187 = v145
		v189 = v124
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
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L41
	}
L40:
	;
	v178 = int32(721025)
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = base.I32_extend8_s(v148)
	F_errmsg_internal(m, int32(665892), v27+int32(48))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(486359), int32(3396), int32(118877))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
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
	v177 = int32(720431)
	goto L46
L45:
	;
	v177 = int32(733277)
	goto L46
L46:
	;
	v178 = v177
	goto L36
L47:
	;
	v294 = v108
	v295 = v110
	v300 = v120
	v301 = v121
	v303 = v189
	goto L32
L48:
	;
	v196 = int32(733277)
	v197 = v145
	v199 = v124
	goto L33
L49:
	;
	F_appendStringInfoString(m, l0, v196)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L58
	}
L50:
	;
	F_appendStringInfoString(m, l0, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L57
	}
L51:
	;
	v201 = int32(719816)
	if v93 == int32(0) {
		v210 = v201
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v120 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v210 = v201
	goto L50
L56:
	;
	v210 = int32(722041)
	goto L50
L57:
	;
	goto L49
L58:
	;
	if v139 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v229 = F_format_type_be(m, v140)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L64
	}
L60:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v218 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v221 = F_quote_identifier(m, v139)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v221
	F_appendStringInfo(m, l0, int32(714219), v27+int32(32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_appendStringInfoString(m, l0, v229)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if v110&v197 != int32(1) {
		v284 = v121
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v285 = int32(1)
	v286 = v120 + v285
	v289 = base.B2i32(v286 == v93) & base.B2i32(v108 == v38-int32(1))
	v294 = v108 - v289
	v295 = (v289 ^ v285) & v110
	v300 = v286
	v301 = v284
	v303 = v199
	goto L32
L67:
	;
	if v199 <= v70 {
		v284 = v121
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	F_initStringInfo(m, v27+int32(112))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v27 + int32(112)
	v247 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(97)))) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(92)))) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(84)))) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(76)))) = v247
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v255
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+107)) = uint8(v255)
	v259 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+105)) = uint16(v259)
	F_get_rule_expr(m, v239, v27+int32(72), v255)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v266
	F_appendStringInfo(m, l0, int32(195305), v27+int32(16))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v274 = v121 + int32(4)
	if base.Ui32(v274) < base.Ui32(v237+v238<<(uint(int32(2))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v280 = v274
	goto L74
L73:
	;
	v280 = int32(0)
	goto L74
L74:
	;
	v284 = v280
	goto L66
L75:
	;
	goto L28
L76:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v339
	F_errmsg_internal(m, int32(48804), v27)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(486359), int32(3348), int32(118877))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v15 = F__emscripten_memset_bulkmem(m, v8+int32(20), base.I32_extend8_s(int32(0)), int32(68))
	mBase = m.M
	goto L1
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v21 = F_pstrdup(m, v16+v17+int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v21
	v30 = F_get_func_arg_info(m, l1, v8+int32(108), v8+int32(104), v8+int32(100))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v33
	v37 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(28))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v8 + int32(112)
	return
L6:
	;
	v39 = F_text_to_cstring(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v41 = F_stringToNode(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v43 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	F_appendStringInfoString(m, l0, int32(728181))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v107 = int32(0)
	F_AcquireRewriteLocks(m, v41, v107, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L25
	}
L12:
	;
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(534511))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L24
	}
L14:
	;
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v54 <= v53 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v58 = v53
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v58<<(uint(int32(2))%32))))
	v67 = int32(0)
	F_AcquireRewriteLocks(m, v66, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v72 = v8 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v72
	v80 = F_list_make1_impl(m, int32(1), v8+int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v82 = int32(0)
	F_get_query_def(m, v66, l0, v80, v82, v82, int32(2), v82, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoChar(m, l0, int32(59))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoChar(m, l0, int32(10))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v96 = v58 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v96 < v97 {
		v58 = v96
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	goto L5
L25:
	;
	v112 = v8 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v112
	v120 = F_list_make1_impl(m, int32(1), v8+int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v122 = int32(0)
	F_get_query_def(m, v41, l0, v120, v122, v122, v122, v122, v122)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	goto L5
}
