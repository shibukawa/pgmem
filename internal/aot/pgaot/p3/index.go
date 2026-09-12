package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildIndexValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v22 = F_check_enable_rls(m, v19, v4, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return v192
L2:
	;
	return int32(0)
L3:
	;
	if v22 == int32(2) {
		v192 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v31 = F_pg_class_aclcheck(m, v19, v29, int64(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	F_initStringInfo(m, v14+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v18 <= int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v44 = int32(0)
	goto L9
L9:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17+int32(48)+v44<<(uint(int32(1))%32)))))
	if v54 == int32(0) {
		v192 = v4
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L5
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v60 = F_pg_attribute_aclcheck(m, v19, v54, v58, int64(2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v60 != 0 {
		v192 = v4
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v63 = v44 + int32(1)
	if v63 != v18 {
		v44 = v63
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v80 = int32(0)
	v82 = int32(1)
	v88 = F_pg_get_indexdef_worker(m, v16, v80, v80, v82, v82, v80, v80, int32(7), v80)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v88
	F_appendStringInfo(m, v14+int32(16), int32(687211), v14)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v18 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_appendStringInfoChar(m, v14+int32(16), int32(41))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L37
	}
L19:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v100 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v114 = int32(304161)
	goto L22
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	F_getTypeOutputInfo(m, v103, v14+int32(12), v14+int32(11))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L23
	}
L22:
	;
	F_appendStringInfoString(m, v14+int32(16), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v112 = F_OidOutputFunctionCall(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v114 = v112
	goto L22
L25:
	;
	v117 = int32(1)
	if v18 == v117 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v124 = v117
	goto L27
L27:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v124))))
	if v133 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L18
L29:
	;
	v137 = v124 << (uint(int32(2)) % 32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137+v138)))
	F_getTypeOutputInfo(m, v140, v14+int32(12), v14+int32(11))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	v152 = int32(304161)
	goto L31
L31:
	;
	F_appendStringInfoString(m, v14+int32(16), int32(747599))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1+v137)))
	v150 = F_OidOutputFunctionCall(m, v147, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v152 = v150
	goto L31
L34:
	;
	F_appendStringInfoString(m, v14+int32(16), v152)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v163 = v124 + int32(1)
	if v163 != v18 {
		v124 = v163
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v192 = v181
	goto L1
}
func F_CompareIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+116)))
	if v22 != v23 {
		v175 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L30
	} else {
		goto L44
	}
L2:
	;
	m.G0 = v20 + int32(16)
	return v175
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v25 != v26 {
		v175 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v28 != v29 {
		v175 = v8
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 != v32 {
		v175 = v8
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != v35 {
		v175 = v8
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v31 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = int32(12)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v54 = v8
	goto L11
L9:
	;
	goto L10
L10:
	;
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if base.B2i32(v122 == v121) == base.B2i32(v125 != v121) {
		v175 = v121
		goto L2
	} else {
		goto L26
	}
L11:
	;
	v62 = v54 << (uint(int32(1)) % 32)
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1+v39+v62))))
	if v43 < v64 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+(l0+v39)))))
	if (v67|v64)&int32(65535) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = int32(0)
	if v64 == v71 {
		v175 = v71
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v54 < v34 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v67 == int32(0) {
		v175 = v71
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76+v64<<(uint(int32(1))%32)-int32(2)))))
	if v82 != v67 {
		v175 = v71
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v86 = int32(0)
	v88 = v54 << (uint(int32(2)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2+v88)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3+v88)))
	if v90 != v92 {
		v175 = v86
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v102 = v54 + int32(1)
	if v102 != v31 {
		v54 = v102
		goto L11
	} else {
		goto L25
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l4+v88)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l5+v88)))
	if v95 != v97 {
		v175 = v86
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L12
L26:
	;
	if v125 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v133 = F_map_variable_attnos(m, v122, int32(1), l6, int32(0), v20+int32(15))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v145 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if base.B2i32(v144 == v145) == base.B2i32(v147 != v145) {
		v175 = v121
		goto L2
	} else {
		goto L35
	}
L30:
	;
	return int32(0)
L31:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	if v137 != 0 {
		v175 = v121
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v139 = F_equal(m, v138, v133)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v139 == int32(0) {
		v175 = v121
		goto L2
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	if v144 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v155 = F_map_variable_attnos(m, v147, int32(1), l6, int32(0), v20+int32(14))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L30
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v164 != 0 {
		v175 = v121
		goto L2
	} else {
		goto L43
	}
L39:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v157 != 0 {
		v175 = v121
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v159 = F_equal(m, v158, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	if v159 == int32(0) {
		v175 = v121
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v175 = base.B2i32(v165 == int32(0))
	goto L2
L44:
	;
	F_errmsg_internal(m, int32(239009), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L30
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(493606), int32(2571), int32(242661))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyIndexTuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v5 = v3 & int32(8191)
	v6 = F_palloc(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v10 = F__emscripten_memcpy_bulkmem(m, v6, l0, v5)
			mBase = m.M
			v11 = v10
		} else {
			v11 = v6
		}
		return v11
	}
}
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v634 int32
	_ = v634
	v9 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(176)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v39 == v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = v39
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = l1
	if int32(0) < v38 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v46 = v42
	goto L3
L6:
	;
	v51 = l1 + int32(28)
	v74 = v9
	v79 = v9
	goto L9
L7:
	;
	v634 = v9
	goto L8
L8:
	;
	m.G0 = v33 + int32(176)
	return v634
L9:
	;
	v83 = v74 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v37+v83)))
	if v85 == int32(0) {
		v601 = v79
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v634 = v601
	goto L8
L11:
	;
	v605 = v74 + int32(1)
	if v605 != v38 {
		v74 = v605
		v79 = v601
		goto L9
	} else {
		goto L125
	}
L12:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+v36)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+118)))
	if v90 != int32(1) {
		v601 = v79
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l7 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+123)))
	if v93 != int32(1) {
		v601 = v79
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+84))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_FormIndexDatum(m, v89, l1, l2, v33+int32(32), v33)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L27
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v99 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v102 = F_ExecPrepareQual(m, v96, l2)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v107 = v99
	goto L22
L22:
	;
	v108 = int32(4520272)
	v109 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v111
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v116 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, v107, v46, v33+int32(175))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+88)) = v102
	if v102 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v107 = v102
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v109
	if v116 == int32(0) {
		v601 = v79
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if l4 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v202 = int32(0)
	if l3 == v202 {
		v502 = v202
		goto L59
	} else {
		goto L60
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	if l6 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v180 = int32(0)
	goto L31
L31:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+12)))
	if v184 != 0 {
		goto L51
	} else {
		goto L52
	}
L32:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+12)))
	if v136 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v140 = int32(0)
	if l6 == v140 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v137 = int32(2)
	goto L37
L36:
	;
	v137 = int32(0)
	goto L37
L37:
	;
	v198 = int32(1)
	v199 = v136
	v200 = v137
	v201 = v85 + int32(192)
	goto L28
L38:
	;
	v180 = v178
	goto L31
L39:
	;
	v178 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v146 <= int32(0) {
		v171 = v140
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v178 = v171
	goto L38
L43:
	;
	v149 = int32(0)
	if v149 < v146 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v152 = v146
	goto L46
L45:
	;
	v152 = v149
	goto L46
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v155 = int32(0)
	goto L47
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153+v155<<(uint(int32(2))%32))))
	v164 = base.B2i32(v163 == v139)
	if v163 == v139 {
		v171 = v164
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v171 = v164
	goto L42
L49:
	;
	v166 = v155 + int32(1)
	if v166 != v152 {
		v155 = v166
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v185 = int32(2)
	goto L53
L52:
	;
	v185 = int32(0)
	goto L53
L53:
	;
	v187 = v85 + int32(192)
	if v184 == int32(0) {
		v198 = v180
		v199 = v184
		v200 = v185
		v201 = v187
		goto L28
	} else {
		goto L54
	}
L54:
	;
	if v180 != 0 {
		v198 = v180
		v199 = v184
		v200 = v185
		v201 = v187
		goto L28
	} else {
		goto L55
	}
L55:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+16)))
	if v192 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v193 = int32(1)
	goto L58
L57:
	;
	v193 = int32(2)
	goto L58
L58:
	;
	v198 = int32(0)
	v199 = v192 ^ int32(1)
	v200 = v193
	v201 = v187
	goto L28
L59:
	;
	v528 = F_index_insert(m, v85, v33+int32(32), v33, v51, v35, v200, v502&int32(1), v89)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L107
	}
L60:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+119)))
	if v205 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+120)))
	v502 = v208
	goto L59
L62:
	;
	goto L63
L63:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+119)) = uint8(v209)
	v211 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v213 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_ExecInitGenerated(m, l0, l2, int32(2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v220 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L67
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+120)) = uint8(v492)
	v502 = v492
	goto L59
L70:
	;
	v492 = int32(0)
	goto L69
L71:
	;
	v492 = int32(1)
	goto L69
L72:
	;
	v224 = v89 + int32(12)
	v234 = v202
	v250 = v220
	v252 = int32(0)
	goto L74
L73:
	;
	if v219 != 0 {
		goto L89
	} else {
		goto L90
	}
L74:
	;
	v259 = int32(*(*int16)(unsafe.Add(mBase, uint32(v224+v234<<(uint(int32(1))%32)))))
	if v259 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v328 == int32(0) {
		goto L71
	} else {
		goto L88
	}
L76:
	;
	v270 = v234
	goto L79
L77:
	;
	v310 = v234
	v312 = v259
	v328 = v252
	goto L78
L78:
	;
	v335 = (v312 + int32(7)) & int32(65535)
	v336 = F_bms_is_member(m, v335, v211)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L83
	}
L79:
	;
	v293 = v270 + int32(1)
	if v250 <= v293 {
		goto L73
	} else {
		goto L81
	}
L80:
	;
	v310 = v293
	v312 = v299
	v328 = v295
	goto L78
L81:
	;
	v295 = int32(1)
	v299 = int32(*(*int16)(unsafe.Add(mBase, uint32(v224+v293<<(uint(v295)%32)))))
	if v299 <= int32(0) {
		v270 = v293
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if v336 != 0 {
		goto L70
	} else {
		goto L84
	}
L84:
	;
	v338 = F_bms_is_member(m, v335, v219)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	if v338 != 0 {
		goto L70
	} else {
		goto L86
	}
L86:
	;
	v341 = v310 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v341 < v342 {
		v234 = v341
		v250 = v342
		v252 = v328
		goto L74
	} else {
		goto L87
	}
L87:
	;
	goto L75
L88:
	;
	goto L73
L89:
	;
	v376 = F_bms_union(m, v211, v219)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	v378 = v211
	goto L91
L91:
	;
	v380 = F_RelationGetIndexExpressions(m, v85)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	v378 = v376
	goto L91
L93:
	;
	F_list_free(m, v380)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L101
	}
L94:
	;
	if v380 == int32(0) {
		v395 = int32(0)
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if v384 == int32(6) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380)+8)))
	v390 = F_bms_is_member(m, v387+int32(7), v378)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v393 = F_expression_tree_walker_impl(m, v380, int32(625), v378)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L100
	}
L99:
	;
	v395 = v390
	goto L93
L100:
	;
	v395 = v393
	goto L93
L101:
	;
	if v219 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_bms_free(m, v378)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v395 != 0 {
		goto L70
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	goto L71
L107:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v89)+92))
	if v530 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if v198 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v549 = v528
	goto L110
L110:
	;
	if v199 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L111:
	;
	v547 = F_check_exclusion_or_unique_constraint(m, v35, v85, v89, v51, v33+int32(32), v33, l2, int32(0), v543, v542&int32(1), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L115
	}
L112:
	;
	v542 = int32(1)
	v543 = int32(2)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536)+16)))
	v538 = int32(1)
	v539 = v537 ^ v538
	v542 = v539
	v543 = v539 & v538
	goto L111
L115:
	;
	v549 = v547
	goto L110
L116:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	v562 = F_lappend_oid(m, v79, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L122
	}
L117:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v89)+92))
	v554 = int32(0)
	if (base.B2i32(v553 == v554)|v549)&int32(1) == v554 {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v549 != 0 {
		v601 = v79
		goto L11
	} else {
		goto L121
	}
L120:
	;
	v601 = v79
	goto L11
L121:
	;
	goto L116
L122:
	;
	if l5 == int32(0) {
		v601 = v562
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+16)))
	if v567&int32(1) == int32(0) {
		v601 = v562
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v572 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v572)
	v601 = v562
	goto L11
L125:
	;
	goto L10
}
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != int32(403) {
		v19 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
			if v23 != 0 {
				v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l2)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v24
					v28 = v26
					m.G0 = v8 + int32(16)
					return v28
				}
			} else {
				v26 = int32(0)
				v28 = v26
				m.G0 = v8 + int32(16)
				return v28
			}
		}
	} else {
		if base.Ui32(int32(5)) <= base.Ui32((l0-int32(1))&int32(65535)) {
			v19 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
				if v23 != 0 {
					v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v24
						v28 = v26
						m.G0 = v8 + int32(16)
						return v28
					}
				} else {
					v26 = int32(0)
					v28 = v26
					m.G0 = v8 + int32(16)
					return v28
				}
			}
		} else {
			v28 = l0
			m.G0 = v8 + int32(16)
			return v28
		}
	}
}
func F_IndexNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = v14 * v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v20 != 0 {
		v43 = v20
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = F_index_getnext_slot(m, v43, v17, v18)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L11
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v28 = F_index_beginscan(m, v21, v22, v23, l0+int32(168), v26, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v28
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v34 != int32(1) {
		v43 = v28
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_index_rescan(m, v28, v37, v38, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v43 = v28
	goto L1
L10:
	;
	m.G0 = v11 + int32(16)
	return v18
L11:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	goto L15
L13:
	;
	goto L14
L14:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)) = uint8(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	m.T0[v101].(func(*base.Module, int32))(m, v18)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L34
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+72)))
	if v58 != int32(1) {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v68 = int32(4520272)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v76 = m.T0[v75].(func(*base.Module, int32, int32, int32) int32)(m, v62, v19, v11+int32(15))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L26
	}
L25:
	;
	goto L10
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v69
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v76 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v83)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v83)+248)) = base.F64_add(v84, float64(1))
	goto L31
L30:
	;
	goto L31
L31:
	;
	v88 = F_index_getnext_slot(m, v43, v17, v18)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v88 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	goto L10
}
func F_build_index_value_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	if l2 == int32(0) {
		v49 = int32(0)
		m.G0 = v8 + int32(160)
		return v49
	} else {
		v14 = F_index_open(m, l3, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			if v18 != int32(1619316) {
				v29 = l2
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				if v30 != 0 {
					v33 = v30
					*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
					v35 = F_BuildIndexInfo(m, v14)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_FormIndexDatum(m, v35, v29, l0, v8+int32(32), v8)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v43 = F_BuildIndexValueDescription(m, v14, v8+int32(32), v8)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_relation_close(m, v14, int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v49 = v43
									m.G0 = v8 + int32(160)
									return v49
								}
							}
						}
					}
				} else {
					v31 = F_MakePerTupleExprContext(m, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v31
						*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
						v35 = F_BuildIndexInfo(m, v14)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_FormIndexDatum(m, v35, v29, l0, v8+int32(32), v8)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v43 = F_BuildIndexValueDescription(m, v14, v8+int32(32), v8)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v14, int32(0))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v49 = v43
										m.G0 = v8 + int32(160)
										return v49
									}
								}
							}
						}
					}
				}
			} else {
				v23 = F_table_slot_create(m, l1, l0+int32(104))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
					m.T0[v26].(func(*base.Module, int32, int32))(m, v23, l2)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v23
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						if v30 != 0 {
							v33 = v30
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
							v35 = F_BuildIndexInfo(m, v14)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_FormIndexDatum(m, v35, v29, l0, v8+int32(32), v8)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v43 = F_BuildIndexValueDescription(m, v14, v8+int32(32), v8)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v14, int32(0))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											v49 = v43
											m.G0 = v8 + int32(160)
											return v49
										}
									}
								}
							}
						} else {
							v31 = F_MakePerTupleExprContext(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v31
								*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
								v35 = F_BuildIndexInfo(m, v14)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_FormIndexDatum(m, v35, v29, l0, v8+int32(32), v8)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v43 = F_BuildIndexValueDescription(m, v14, v8+int32(32), v8)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v14, int32(0))
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												v49 = v43
												m.G0 = v8 + int32(160)
												return v49
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
func F_check_index_is_clusterable(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = F_index_open(m, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+192))
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v47 = int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v46 + v47
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v45 + v47
					F_errmsg(m, int32(719077), v7)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errfinish(m, int32(496368), int32(509), int32(396272))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v14 != v15 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v47 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v46 + v47
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v45 + v47
						F_errmsg(m, int32(719077), v7)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errfinish(m, int32(496368), int32(509), int32(396272))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v68 + int32(4)
							F_errmsg(m, int32(331786), v5+int32(-16))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_errfinish(m, int32(496368), int32(516), int32(396272))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+196))
					v24 = F_heap_attisnull(m, v21, int32(21), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v89 + int32(4)
									F_errmsg(m, int32(695933), v5+int32(-32))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										F_errfinish(m, int32(496368), int32(528), int32(396272))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
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
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+192))
							v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
							if v29 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v110 + int32(4)
										F_errmsg(m, int32(698220), v5+int32(-48))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_errfinish(m, int32(496368), int32(542), int32(396272))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
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
								F_relation_close(m, v9, int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									m.G0 = v7 - int32(-64)
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
func F_create_index_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	v15 = F_palloc0(m, int32(112))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(280)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v21
		if l7 != 0 {
			v25 = int32(342)
		} else {
			v25 = int32(341)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v27
		v29 = F_get_baserel_parampathinfo(m, l0, v21, l8)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v31)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+26)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v34)
			F_cost_index(m, v15, l0, l9, l10)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_index_am_oid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_get_am_type_oid(m, l0, int32(105), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_get_index_isreplident(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	v4 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9)+22)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(1)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(1)
		}
	}
}
func F_index_beginscan_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v6 = l5
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v17 != v13 {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[47]))
		v21 = F_list_member_ptr(m, v20, v13)
		mBase = m.M
		v22 = v21
	} else {
		v22 = int32(1)
	}
	if v22 == int32(0) {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
		if v26 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(285056)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v71 + int32(4)
				F_errmsg_internal(m, int32(694518), v11+int32(16))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498364), int32(321), int32(312231))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+23)))
			if v29 == int32(0) {
				F_PredicateLockRelation(m, l0, l3)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_RelationIncrementReferenceCount(m, l0)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+92))
						v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v40)+29)) = uint8(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+88)) = l4
							m.G0 = v11 + int32(32)
							return v40
						}
					}
				}
			} else {
				F_RelationIncrementReferenceCount(m, l0)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+92))
					v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v40)+29)) = uint8(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+88)) = l4
						m.G0 = v11 + int32(32)
						return v40
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v55 + int32(4)
				F_errmsg(m, int32(439693), v11)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498364), int32(320), int32(312231))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
func F_index_expression_changed_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(6) {
			v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v13 = F_bms_is_member(m, v10+int32(7), l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		} else {
			v19 = F_expression_tree_walker_impl(m, l0, int32(625), l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	}
}
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = l1 - int32(1)
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v15 {
		v22 = l2 + v14<<(uint(int32(4))%32) + int32(20)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int32(0) {
			v69 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v75 = v69
				m.G0 = v9 + int32(16)
				return v75
			}
		} else {
			v28 = l0 + v23 + int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
			if v29 != int32(1) {
				v75 = v28
				m.G0 = v9 + int32(16)
				return v75
			} else {
				v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
				switch v32&int32(65535) - int32(1) {
				case 0:
					v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
					v75 = v37
					m.G0 = v9 + int32(16)
					return v75
				case 1:
					v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28))))
					v75 = v38
					m.G0 = v9 + int32(16)
					return v75
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
						F_errmsg_internal(m, int32(484426), v9)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(327351), int32(70), int32(68101))
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
				case 3:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v75 = v39
					m.G0 = v9 + int32(16)
					return v75
				}
			}
		}
	} else {
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v14>>(uint(int32(3))%32))+8)))
		if int32(base.Ui32(v58)>>(uint(v14&int32(7))%32))&int32(1) != 0 {
			v69 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v75 = v69
				m.G0 = v9 + int32(16)
				return v75
			}
		} else {
			v64 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
			v75 = int32(0)
			m.G0 = v9 + int32(16)
			return v75
		}
	}
}
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	if v15 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
			m.T0[v19].(func(*base.Module, int32))(m, v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = v22
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v24)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v24)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
				m.T0[v29].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		} else {
			v23 = v13
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v24)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v24)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
			m.T0[v29].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(285110)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v40 + int32(4)
			F_errmsg_internal(m, int32(694518), v11)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(498364), int32(361), int32(285128))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
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
func F_index_set_state_flags(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_SearchSysCacheCopy(m, int32(34), l0, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v20 = v18 + v19
				switch l1 {
				case 0:
					v21 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v21)
				case 1:
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+18)) = uint8(v23)
				case 2:
					v25 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)) = uint8(v25)
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+17)) = uint16(v25)
				case 3:
					v29 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+20)) = uint16(v29)
				default:
				}
				F_CatalogTupleUpdate(m, v12, v16+int32(4), v16)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_sequence_close(m, v12, int32(3))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(40367), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(493606), int32(3515), int32(157223))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
}
