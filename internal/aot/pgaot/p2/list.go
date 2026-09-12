package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_list_concat_unique_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = int32(0)
	if l1 == v3 {
		v69 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v69 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = l0
	v15 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v15<<(uint(int32(2))%32))))
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v69 = v58
	goto L1
L6:
	;
	v66 = v15 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v66 < v67 {
		v13 = v58
		v15 = v66
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v54 = F_lappend(m, v13, v24)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v35 = int32(0)
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30+v35<<(uint(int32(2))%32))))
	if v42 == v24 {
		v58 = v13
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v45 = v35 + int32(1)
	if v27 != v45 {
		v35 = v45
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	v58 = v54
	goto L6
L16:
	;
	goto L5
}
func F_list_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	v3 = int32(0)
	if l0 == v3 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return l0
L5:
	;
	v16 = v3
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = v20 + v16<<(uint(int32(2))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = F_equal(m, v24, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v29 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	return int32(0)
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v31 = v16 + int32(1)
	if v31 < v29 {
		v16 = v31
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L4
L13:
	;
	if l0+int32(16) != v33 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v46 = v23 + int32(4)
	v50 = int32(2)
	v54 = (v29 + int32(base.Ui32(v23-v33^int32(-1))>>(uint(v50)%32))) << (uint(v50) % 32)
	if v23 == v46 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	F_pfree(m, v33)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199 - int32(1)
	goto L4
L22:
	;
	goto L21
L23:
	;
	v58 = v23 + v54
	if base.Ui32(v46-v58) <= base.Ui32(int32(0)-v54<<(uint(int32(1))%32)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v65 = F___memcpy(m, v23, v46, v54)
	mBase = m.M
	goto L21
L25:
	;
	goto L26
L26:
	;
	v68 = (v23 ^ v46) & int32(3)
	if base.Ui32(v23) < base.Ui32(v46) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v170 == int32(0) {
		goto L22
	} else {
		goto L63
	}
L28:
	;
	if base.Ui32(v148) <= base.Ui32(int32(3)) {
		v169 = v147
		v170 = v148
		v171 = v149
		goto L27
	} else {
		goto L59
	}
L29:
	;
	if v68 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v68 != 0 {
		v130 = v54
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v169 = v46
	v170 = v54
	v171 = v23
	goto L27
L33:
	;
	goto L34
L34:
	;
	if v23&int32(3) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v147 = v46
	v148 = v54
	v149 = v23
	goto L28
L36:
	;
	goto L37
L37:
	;
	v75 = v46
	v76 = v54
	v77 = v23
	goto L38
L38:
	;
	if v76 == int32(0) {
		goto L22
	} else {
		goto L40
	}
L39:
	;
	v147 = v84
	v148 = v86
	v149 = v88
	goto L28
L40:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v81)
	v83 = int32(1)
	v84 = v75 + v83
	v86 = v76 - v83
	v88 = v77 + v83
	if v88&int32(3) != 0 {
		v75 = v84
		v76 = v86
		v77 = v88
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v130 == int32(0) {
		goto L22
	} else {
		goto L55
	}
L43:
	;
	if v58&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v95 = v54
	goto L47
L45:
	;
	v110 = v54
	goto L46
L46:
	;
	if base.Ui32(v110) <= base.Ui32(int32(3)) {
		v130 = v110
		goto L42
	} else {
		goto L51
	}
L47:
	;
	if v95 == int32(0) {
		goto L22
	} else {
		goto L49
	}
L48:
	;
	v110 = v101
	goto L46
L49:
	;
	v101 = v95 - int32(1)
	v102 = v23 + v101
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v104)
	if v102&int32(3) != 0 {
		v95 = v101
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v117 = v110
	goto L52
L52:
	;
	v121 = v117 - int32(4)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v46+v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v121))) = v124
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v117 = v121
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v130 = v121
	goto L42
L54:
	;
	goto L53
L55:
	;
	v137 = v130
	goto L56
L56:
	;
	v141 = v137 - int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v141))) = uint8(v144)
	if v141 != 0 {
		v137 = v141
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L22
L58:
	;
	goto L57
L59:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L60
L60:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v160 = int32(4)
	v161 = v154 + v160
	v163 = v156 + v160
	v165 = v155 - v160
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v154 = v161
		v155 = v165
		v156 = v163
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v169 = v161
	v170 = v165
	v171 = v163
	goto L27
L62:
	;
	goto L61
L63:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L64
L64:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v180)
	v182 = int32(1)
	v187 = v177 - v182
	if v187 != 0 {
		v176 = v176 + v182
		v177 = v187
		v178 = v178 + v182
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L22
L66:
	;
	goto L65
}
func F_list_make2_impl(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(17179869186)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v5 + int32(16)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v18
		return v5
	}
}
func F_list_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v2 = int32(0)
	if l0 == v2 {
		v68 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v68
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		v68 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v2
	v14 = v2
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v18 + v14<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v13 == int32(0) {
		v52 = v22
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v68 = v58
	goto L1
L6:
	;
	v64 = v14 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v64 < v65 {
		v13 = v58
		v14 = v64
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v55 = F_lappend(m, v13, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v26 <= v25 {
		v52 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = v25
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v33<<(uint(int32(2))%32))))
	v40 = F_equal(m, v39, v22)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v52 = v48
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	if v40 != 0 {
		v58 = v13
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v33 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v45 < v46 {
		v33 = v45
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v58 = v55
	goto L6
L17:
	;
	goto L5
}
func F_writeListPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(8208)
	m.G0 = v17
	if l1 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = int32(1)
	v38 = int32(4548788)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v40 + v37
	v44 = int32(16)
	if l1 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+(l1^int32(-1))<<(uint(int32(2))%32))))
	v36 = v28
	goto L1
L3:
	;
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v36 = v30 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	if int32(0) < l3 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	F_PageInit(m, v62, int32(8192), int32(8))
	mBase = m.M
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)))
	v67 = v62 + v66
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v44)
	goto L5
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(l1^int32(-1))<<(uint(int32(2))%32))))
	v62 = v54
	goto L6
L8:
	;
	goto L9
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v62 = v56 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L20
	} else {
		goto L45
	}
L11:
	;
	v83 = v6
	v84 = v17 + int32(16)
	v85 = v37
	v86 = v6
	goto L14
L12:
	;
	v122 = v6
	goto L13
L13:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v36+v128))) = l4
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	if l4 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+v86<<(uint(int32(2))%32))))
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)))
	v95 = v93 & int32(8191)
	if v95 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v122 = v110
	goto L13
L16:
	;
	v101 = F_PageAddItemExtended(m, v36, v92, v95, v85&int32(65535), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v96 = F__emscripten_memcpy_bulkmem(m, v84, v92, v95)
	mBase = m.M
	v97 = v96
	goto L19
L18:
	;
	v97 = v84
	goto L19
L19:
	;
	goto L16
L20:
	;
	return int32(0)
L21:
	;
	if v101 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v107 = int32(1)
	v110 = v95 + v83
	v112 = v86 + v107
	if v112 != l3 {
		v83 = v110
		v84 = v95 + v97
		v85 = v85 + v107
		v86 = v112
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v134 = v131 + v36
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+6)))
	v137 = v135 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+6)) = uint16(v137)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
	v142 = int32(1)
	v143 = v140
	goto L26
L25:
	;
	v142 = v6
	v143 = v131
	goto L26
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v143+v36)+4)) = uint16(v142)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+118)))
	if v149 != int32(112) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+14)))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
	v185 = v183 - v184
	v186 = int32(0)
	if v186 < v185 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v153 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v156 != 0 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L20
	} else {
		goto L35
	}
L33:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v157 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v162 = int32(8)
	F_XLogRegisterData(m, v17+v162, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(6))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	F_XLogRegisterBufData(m, int32(0), v17+int32(16), v122)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v178 = F_XLogInsert(m, int32(13), int32(112))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = base.I64_rotr(v178, int64(32))
	goto L28
L40:
	;
	F_UnlockReleaseBuffer(m, l1)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L20
	} else {
		goto L44
	}
L41:
	;
	v189 = v185
	goto L43
L42:
	;
	v189 = v186
	goto L43
L43:
	;
	goto L40
L44:
	;
	v192 = int32(4548788)
	v194 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v194 - int32(1)
	m.G0 = v17 + int32(8208)
	return v189
L45:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v206 + int32(4)
	F_errmsg_internal(m, int32(744597), v17)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(514923), int32(90), int32(427823))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
