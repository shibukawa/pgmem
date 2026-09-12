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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
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
	if v26 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if int32(0) < v71 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v64 = v5
	goto L3
L5:
	;
	goto L6
L6:
	;
	v36 = v5
	goto L7
L7:
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
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v64 = v56
	goto L3
L9:
	;
	goto L8
L10:
	;
	v75 = l2 + int32(36)
	v81 = int32(0)
	v82 = v64
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l1
	F_qsort_arg(m, v22, v19, int32(12), int32(22), v15+int32(8))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v91 = v22 + v82*int32(12)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v93 = int32(1)
	v96 = int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v92<<(uint(v93)%32)+v81)<<(uint(v96)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v101<<(uint(v93)%32)+v81)<<(uint(v96)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)) = uint8(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v108
	v115 = v81 + v93
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v115 < v116 {
		v81 = v115
		v82 = v82 + v93
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	goto L14
L16:
	;
	if v19 < int32(2) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184
	m.G0 = v15 + int32(16)
	return v22
L18:
	;
	v184 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v147 = int32(1)
	v148 = int32(1)
	goto L21
L21:
	;
	v155 = int32(12)
	v157 = v22 + v148*v155
	v162 = F_compare_expanded_ranges(m, v157-v155, v157, v15+int32(8))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v184 = v175
	goto L17
L23:
	;
	if v162 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v147 != v148 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v175 = v147
	goto L26
L26:
	;
	v178 = v148 + int32(1)
	if v178 != v19 {
		v147 = v175
		v148 = v178
		goto L21
	} else {
		goto L30
	}
L27:
	;
	v167 = v22 + v147*int32(12)
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v170
	goto L29
L28:
	;
	goto L29
L29:
	;
	v175 = v147 + int32(1)
	goto L26
L30:
	;
	goto L22
}
func F_expanded_record_set_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
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
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v11&int32(64) != 0 {
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
	if l1 == int32(0) {
		v132 = l1
		v133 = l3
		goto L22
	} else {
		goto L23
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	F_MemoryContextReset(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L12
	} else {
		goto L21
	}
L5:
	;
	v17 = l0 + int32(96)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v18 == int32(0) {
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
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v33 = int32(4442992)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v32
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, int32(0), int32(1), v39, l0+int32(104), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = F_AllocSetContextCreateInternal(m, v21, int32(58198), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_MemoryContextReset(m, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v26
	v32 = v26
	goto L8
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v32 = v31
	goto L8
L15:
	;
	v81 = v17
	v83 = v34
	goto L4
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+84)) = l1
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+88)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v52 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+92)) = v49 + v51
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
	if v59&int32(4) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v52 | int32(17)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v65 = int32(4442992)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v68
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, v47+int32(18), int32(0), v73, l0+int32(104), v76)
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
	v83 = v66
	goto L4
L21:
	;
	goto L3
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v137 = v135 & int32(224)
	if l2 == int32(0) {
		v158 = v132
		v160 = v137
		goto L33
	} else {
		goto L34
	}
L23:
	;
	if l3 == int32(0) {
		v132 = l1
		v133 = l3
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+20)))
	if v102&int32(4) == v100 {
		v132 = l1
		v133 = v100
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v107 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v122 = int32(4442992)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v121
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v127 = F_toast_flatten_tuple(m, l1, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L32
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = F_AllocSetContextCreateInternal(m, v110, int32(58198), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L12
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_MemoryContextReset(m, v107)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v115
	v121 = v115
	goto L26
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v121 = v120
	goto L26
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v123
	v132 = v127
	v133 = int32(1)
	goto L22
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v158 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	if v132 == int32(0) {
		v158 = v132
		v160 = v137
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v142 = int32(4442992)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v145
	v147 = F_heap_copytuple(m, v132)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v143
	v152 = v137 | int32(2)
	if v133 == int32(0) {
		v158 = v147
		v160 = v152
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	v158 = v147
	v160 = v152
	goto L33
L39:
	;
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v183
	if v135&int32(8) == v184 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v158
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v165 + v167
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+20)))
	if v173&int32(4) != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+84)) = int64(0)
	v183 = v160
	goto L39
L43:
	;
	v176 = int32(17)
	goto L45
L44:
	;
	v176 = int32(1)
	goto L45
L45:
	;
	v183 = v176 | v160
	goto L39
L46:
	;
	if v135&int32(2) != 0 {
		goto L57
	} else {
		goto L58
	}
L47:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v191 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v199 = int32(0)
	v202 = v191
	goto L49
L49:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v199))))
	if v210 != 0 {
		v231 = v202
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L46
L51:
	;
	v233 = v199 + int32(1)
	if v233 < v231 {
		v199 = v233
		v202 = v231
		goto L49
	} else {
		goto L56
	}
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+int32(102)+v211<<(uint(int32(4))%32)+v199*int32(100)))))
	if v218 != 0 {
		v231 = v202
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v199<<(uint(int32(2))%32))))
	if base.B2i32(base.Ui32(v162) <= base.Ui32(v223))&base.B2i32(base.Ui32(v223) < base.Ui32(v161)) != 0 {
		v231 = v202
		goto L51
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v223)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v231 = v229
	goto L51
L56:
	;
	goto L50
L57:
	;
	F_pfree(m, v163)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L12
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return
L60:
	;
	goto L59
}
