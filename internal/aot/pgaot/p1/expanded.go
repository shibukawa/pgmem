package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MakeExpandedObjectReadOnlyInternal(m *base.Module, l0 int64) int64 {
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
	var v14 int64
	_ = v14
	v3 = base.I32_wrap_i64(l0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v4 != int32(1) {
		v14 = l0
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		if v7 != int32(3) {
			v14 = l0
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+2))
			v14 = base.I64_extend_i32_u(v10 + int32(18))
		}
	}
	return v14
}
func F_build_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v20 = v18 + v19
	v23 = F_palloc0(m, v20*int32(24))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if int32(0) < v27 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = v5
	goto L6
L4:
	;
	v66 = v5
	goto L5
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if int32(0) < v74 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v47 = v23 + v37*int32(24)
	v50 = l2 + int32(40) + v37<<(uint(int32(4))%32)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)) = uint8(v54)
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v53
	v58 = v37 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v58 < v59 {
		v37 = v58
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v66 = v58
	goto L5
L8:
	;
	goto L7
L9:
	;
	v78 = l2 + int32(40)
	v84 = int32(0)
	v85 = v66
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l1
	F_qsort_arg(m, v23, v20, int32(24), int32(23), v16+int32(8))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v95 = v23 + v85*int32(24)
	v97 = v84 << (uint(int32(3)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v99 = int32(4)
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v97+(v78+v98<<(uint(v99)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v95))) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v78+v105<<(uint(v99)%32)+v97)))
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+16)) = uint8(v111)
	*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v110
	v117 = v84 + v111
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v117 < v118 {
		v84 = v117
		v85 = v85 + v111
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
	v141 = int32(1)
	if int32(2) <= v20 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v149 = v141
	v150 = int32(1)
	goto L19
L17:
	;
	v189 = v141
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v189
	m.G0 = v16 + int32(16)
	return v23
L19:
	;
	v158 = int32(24)
	v160 = v23 + v150*v158
	v165 = F_compare_expanded_ranges(m, v160-v158, v160, v16+int32(8))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v189 = v181
	goto L18
L21:
	;
	if v165 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v149 != v150 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v181 = v149
	goto L24
L24:
	;
	v183 = v150 + int32(1)
	if v183 != v20 {
		v149 = v181
		v150 = v183
		goto L19
	} else {
		goto L28
	}
L25:
	;
	v170 = v23 + v149*int32(24)
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v160)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = v171
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v160)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v173
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v175
	goto L27
L26:
	;
	goto L27
L27:
	;
	v181 = v149 + int32(1)
	goto L24
L28:
	;
	goto L20
}
func F_expanded_record_set_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v10&int32(64) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v94 = int32(0)
	if base.B2i32(l1 == v94)|base.B2i32(l3 == v94) != 0 {
		v131 = l1
		v132 = l3
		goto L22
	} else {
		goto L23
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	F_MemoryContextReset(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L21
	}
L5:
	;
	v16 = l0 + int32(96)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_build_dummy_expanded_header(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v32 = int32(_a_F_expanded_record_set_tuple_0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v31
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, int64(0), int32(1), v38, l0+int32(104), v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L15
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_expanded_record_set_tuple_1), int32(0), int32(1024), int32(_a_F_expanded_record_set_tuple_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_MemoryContextReset(m, v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v25
	v31 = v25
	goto L8
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v31 = v30
	goto L8
L15:
	;
	v81 = v16
	v83 = v33
	goto L4
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+84)) = l1
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+88)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v51 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+92)) = v48 + v50
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)))
	if v58&int32(4) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v51 | int32(17)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v64 = int32(_a_F_expanded_record_set_tuple_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v67
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, base.I64_extend_i32_u(v46+int32(18)), int32(0), v73, l0+int32(104), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v81 = l0 + int32(96)
	v83 = v65
	goto L4
L21:
	;
	goto L3
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v136 = v134 & int32(224)
	v137 = int32(0)
	if base.B2i32(l2 == v137)|base.B2i32(v131 == v137) != 0 {
		v158 = v131
		v160 = v136
		goto L32
	} else {
		goto L33
	}
L23:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)))
	if v101&int32(4) == v99 {
		v131 = l1
		v132 = v99
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v106 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v121 = int32(_a_F_expanded_record_set_tuple_0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v120
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v126 = F_toast_flatten_tuple(m, l1, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L31
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = F_AllocSetContextCreateInternal(m, v109, int32(_a_F_expanded_record_set_tuple_1), int32(0), int32(1024), int32(_a_F_expanded_record_set_tuple_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_MemoryContextReset(m, v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v114
	v120 = v114
	goto L25
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v120 = v119
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v122
	v131 = v126
	v132 = int32(1)
	goto L22
L32:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v158 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v142 = int32(_a_F_expanded_record_set_tuple_0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v145
	v147 = F_heap_copytuple(m, v131)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v143
	v152 = v136 | int32(2)
	if v132 == int32(0) {
		v158 = v147
		v160 = v152
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v158 = v147
	v160 = v152
	goto L32
L37:
	;
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v183
	if v134&int32(8) == v184 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v158
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v165 + v167
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+20)))
	if v173&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+84)) = int64(0)
	v183 = v160
	goto L37
L41:
	;
	v176 = int32(17)
	goto L43
L42:
	;
	v176 = int32(1)
	goto L43
L43:
	;
	v183 = v176 | v160
	goto L37
L44:
	;
	if v134&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v191 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v197 = int32(0)
	v200 = v191
	goto L47
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205+v197))))
	if v207 != 0 {
		v222 = v200
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L44
L49:
	;
	v224 = v197 + int32(1)
	if v224 < v222 {
		v197 = v224
		v200 = v222
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v209 = v197 << (uint(int32(3)) % 32)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v209)+32)))
	if v211 != 0 {
		v222 = v200
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v212+v209)))
	if base.B2i32(base.Ui32(v162) <= base.Ui32(v214))&base.B2i32(base.Ui32(v214) < base.Ui32(v161)) != 0 {
		v222 = v200
		goto L49
	} else {
		goto L52
	}
L52:
	;
	F_pfree(m, v214)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v222 = v220
	goto L49
L54:
	;
	goto L48
L55:
	;
	F_pfree(m, v163)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L12
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	return
L58:
	;
	goto L57
}
