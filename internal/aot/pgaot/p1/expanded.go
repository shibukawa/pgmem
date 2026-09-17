package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MakeExpandedObjectReadOnlyInternal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2 != int32(1) {
		v11 = l0
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v5 != int32(3) {
			v11 = l0
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v11 = v8 + int32(18)
		}
	}
	return v11
}
func F_build_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v19 = v17 + v18
	v22 = F_palloc0(m, v19*int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if int32(0) < v26 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = v5
	goto L6
L4:
	;
	v64 = v5
	goto L5
L5:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if int32(0) < v71 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v45 = v22 + v36*int32(12)
	v48 = l2 + int32(36) + v36<<(uint(int32(3))%32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+8)) = uint8(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v51
	v56 = v36 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v56 < v57 {
		v36 = v56
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v64 = v56
	goto L5
L8:
	;
	goto L7
L9:
	;
	v75 = l2 + int32(36)
	v81 = int32(0)
	v82 = v64
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l1
	F_qsort_arg(m, v22, v19, int32(12), int32(22), v15+int32(8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v91 = v22 + v82*int32(12)
	v93 = v81 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v95 = int32(3)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v75+v94<<(uint(v95)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v75+v101<<(uint(v95)%32)+v93)))
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)) = uint8(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v106
	v113 = v81 + v107
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v113 < v114 {
		v81 = v113
		v82 = v82 + v107
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
	v136 = int32(1)
	if int32(2) <= v19 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v144 = v136
	v145 = int32(1)
	goto L19
L17:
	;
	v181 = v136
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v181
	m.G0 = v15 + int32(16)
	return v22
L19:
	;
	v152 = int32(12)
	v154 = v22 + v145*v152
	v159 = F_compare_expanded_ranges(m, v154-v152, v154, v15+int32(8))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v181 = v173
	goto L18
L21:
	;
	if v159 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v144 != v145 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v173 = v144
	goto L24
L24:
	;
	v175 = v145 + int32(1)
	if v175 != v19 {
		v144 = v173
		v145 = v175
		goto L19
	} else {
		goto L28
	}
L25:
	;
	v164 = v22 + v144*int32(12)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v165
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v167
	goto L27
L26:
	;
	goto L27
L27:
	;
	v173 = v144 + int32(1)
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
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
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
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
	v93 = int32(0)
	if base.B2i32(l1 == v93)|base.B2i32(l3 == v93) != 0 {
		v130 = l1
		v131 = l3
		goto L22
	} else {
		goto L23
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	F_MemoryContextReset(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	F_domain_check(m, int32(0), int32(1), v38, l0+int32(104), v41)
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
	v80 = v16
	v82 = v33
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, v46+int32(18), int32(0), v72, l0+int32(104), v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v80 = l0 + int32(96)
	v82 = v65
	goto L4
L21:
	;
	goto L3
L22:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v135 = v133 & int32(224)
	v136 = int32(0)
	if base.B2i32(l2 == v136)|base.B2i32(v130 == v136) != 0 {
		v157 = v130
		v159 = v135
		goto L32
	} else {
		goto L33
	}
L23:
	;
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+20)))
	if v100&int32(4) == v98 {
		v130 = l1
		v131 = v98
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v105 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v120 = int32(_a_F_expanded_record_set_tuple_0)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v119
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v125 = F_toast_flatten_tuple(m, l1, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L31
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = F_AllocSetContextCreateInternal(m, v108, int32(_a_F_expanded_record_set_tuple_1), int32(0), int32(1024), int32(_a_F_expanded_record_set_tuple_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_MemoryContextReset(m, v105)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v113
	v119 = v113
	goto L25
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v119 = v118
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v121
	v130 = v125
	v131 = int32(1)
	goto L22
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v157 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v141 = int32(_a_F_expanded_record_set_tuple_0)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v144
	v146 = F_heap_copytuple(m, v130)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_tuple[0])) = v142
	v151 = v135 | int32(2)
	if v131 == int32(0) {
		v157 = v146
		v159 = v151
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v157 = v146
	v159 = v151
	goto L32
L37:
	;
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v182
	if v133&int32(8) == v183 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v157
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v164 + v166
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+20)))
	if v172&int32(4) != 0 {
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
	v182 = v159
	goto L37
L41:
	;
	v175 = int32(17)
	goto L43
L42:
	;
	v175 = int32(1)
	goto L43
L43:
	;
	v182 = v175 | v159
	goto L37
L44:
	;
	if v133&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v190 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v196 = int32(0)
	v199 = v190
	goto L47
L47:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v196))))
	if v206 != 0 {
		v227 = v199
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L44
L49:
	;
	v229 = v196 + int32(1)
	if v229 < v227 {
		v196 = v229
		v199 = v227
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v207<<(uint(int32(4))%32)+v196*int32(100))+102)))
	if v214 != 0 {
		v227 = v199
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215+v196<<(uint(int32(2))%32))))
	if base.B2i32(base.Ui32(v161) <= base.Ui32(v219))&base.B2i32(base.Ui32(v219) < base.Ui32(v160)) != 0 {
		v227 = v199
		goto L49
	} else {
		goto L52
	}
L52:
	;
	F_pfree(m, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v227 = v225
	goto L49
L54:
	;
	goto L48
L55:
	;
	F_pfree(m, v162)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
