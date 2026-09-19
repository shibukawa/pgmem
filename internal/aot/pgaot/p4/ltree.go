package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v2 {
		v30 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v30&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v17 == int32(0) {
		v30 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v20 != int32(7) {
		v30 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v23 != int32(17) {
		v30 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
	v30 = v26 ^ int32(1)
	goto L2
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_get_fn_opclass_options(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v39 = int32(28)
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	if v41 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(0)
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v39 = v38
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L55
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L51
	}
L14:
	;
	v201 = F_palloc(m, int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L50
	}
L15:
	;
	v44 = F_pg_detoast_datum(m, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
	if v148&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v49 = F_ArrayGetNItemsSafe(m, v46, v44+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(2) <= v51 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v55 = F_array_contains_nulls(m, v44)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	if v55 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v57 = int32(0)
	v61 = F_ltree_gist_alloc(m, v57, v57, v39, v57, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v49 <= int32(0) {
		v196 = v61
		goto L14
	} else {
		goto L24
	}
L24:
	;
	if v54 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v71 = v54
	goto L27
L26:
	;
	v71 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L27
L27:
	;
	v79 = v44 + v71
	v81 = v49
	goto L28
L28:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	if v87 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v196 = v61
	goto L14
L30:
	;
	v90 = v79 + int32(8)
	v91 = v87
	goto L33
L31:
	;
	goto L32
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v144 = int32(1)
	if v144 < v81 {
		v79 = v79 + (int32(base.Ui32(v136)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v81 = v81 - v144
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	v103 = F_ltree_crc32_sz(m, v90+int32(2), v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v105 = base.I32_rem_u_s(v103, v39<<(uint(int32(3))%32))
	v108 = v61 + int32(8) + int32(base.Ui32(v105)>>(uint(int32(3))%32))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v110 = int32(1)
	v114 = v109 | v110<<(uint(v105&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v114)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	if base.Ui32(v110) < base.Ui32(v91) {
		v90 = v90 + (v116+int32(9))&int32(_a_F__ltree_compress_0)
		v91 = v91 - v110
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L29
L38:
	;
	return v11
L39:
	;
	goto L40
L40:
	;
	v153 = v40 + int32(8)
	v154 = int32(0)
	if v39 <= v154 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v186 = int32(0)
	v188 = F_ltree_gist_alloc(m, int32(1), v153, v39, v186, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L49
	}
L42:
	;
	v157 = v154
	goto L43
L43:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v153))))
	if v168 == int32(255) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	return v11
L45:
	;
	v172 = v157 + int32(1)
	if v39 != v172 {
		v157 = v172
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L41
L49:
	;
	v196 = v188
	goto L14
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v196
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v206
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+14)) = uint8(v209)
	*(*uint16)(unsafe.Add(mBase, uint32(v201)+12)) = uint16(v208)
	return v201
L51:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F__ltree_compress_1), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F__ltree_compress_2), int32(67), int32(_a_F__ltree_compress_3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F__ltree_compress_4), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F__ltree_compress_2), int32(71), int32(_a_F__ltree_compress_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltree_isparent(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13845(m, l0, int32(_a_F__ltree_isparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F__ltree_same(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13906(m, l0, int32(1), int32(2), int32(28))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_ltree_crc32_sz(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_crc32_sz[0]))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_pg_newlocale_from_collation(m, int32(100))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = v16
	goto L3
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+3)))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ltree_crc32_sz[0])) = v21
	v26 = v21
	goto L3
L6:
	;
	m.G0 = v13 + int32(16)
	return v173 ^ int32(-1)
L7:
	;
	v77 = l0
	v78 = l1
	v79 = int32(-1)
	goto L20
L8:
	;
	if int32(0) < l1 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = int32(-1)
	if l1 <= int32(0) {
		v173 = v33
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v173 = int32(-1)
	goto L6
L12:
	;
	v36 = l0
	v37 = l1
	v38 = v33
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32((v46-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v173 = v66
	goto L6
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32((v55^v38)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v66 = v63 ^ int32(base.Ui32(v38)>>(uint(int32(8))%32))
	v67 = int32(1)
	if base.Ui32(v67) < base.Ui32(v37) {
		v36 = v36 + v67
		v37 = v37 - v67
		v38 = v66
		goto L13
	} else {
		goto L19
	}
L16:
	;
	v55 = v46 | int32(32)
	goto L18
L17:
	;
	v55 = v46
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L14
L20:
	;
	v88 = v13 + int32(4)
	v90 = F_pg_mblen_range(m, v77, l0+l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v173 = v159
	goto L6
L22:
	;
	v168 = v78 - v90
	if int32(0) < v168 {
		v77 = v77 + v90
		v78 = v168
		v79 = v159
		goto L20
	} else {
		goto L33
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_crc32_sz[0]))
	v94 = F_pg_strfold(m, v88, int32(12), v77, v90, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v94 == int32(0) {
		v159 = v79
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v94&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32((v79^v100)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v114 = v108 ^ int32(base.Ui32(v79)>>(uint(int32(8))%32))
	v115 = v13 + int32(5)
	v116 = v94 - int32(1)
	goto L28
L27:
	;
	v114 = v79
	v115 = v88
	v116 = v94
	goto L28
L28:
	;
	if v94 == int32(1) {
		v159 = v114
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v121 = v114
	v122 = v115
	v126 = v116
	goto L30
L30:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v131 = int32(255)
	v133 = int32(2)
	v137 = *(*int32)(unsafe.Add(mBase, uint32((v129^v121)&v131<<(uint(v133)%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v138 = int32(8)
	v140 = v137 ^ int32(base.Ui32(v121)>>(uint(v138)%32))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32((v140^v141)&v131<<(uint(v133)%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v152 = v149 ^ int32(base.Ui32(v140)>>(uint(v138)%32))
	v156 = v126 - v133
	if v156 != 0 {
		v121 = v152
		v122 = v122 + v133
		v126 = v156
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v159 = v152
	goto L22
L32:
	;
	goto L31
L33:
	;
	goto L21
}
func F_ltree_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
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
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v24 = int32(0)
	if base.B2i32(v23 == v24)|base.B2i32(v22 == v24) != 0 {
		v152 = v23
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v174 != v15 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v172 = (v152 + int32(1)) * (v23 - v22) * int32(10)
	goto L4
L6:
	;
	v29 = int32(8)
	v37 = v15 + v29
	v38 = v20 + v29
	v39 = v23
	v46 = v22
	goto L7
L7:
	;
	v47 = int32(2)
	v48 = v37 + v47
	v50 = v38 + v47
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v51) < base.Ui32(v52) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v125
	goto L5
L9:
	;
	v54 = v51
	goto L11
L10:
	;
	v54 = v52
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v116 != 0 {
		v172 = int32(1)
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v116 = int32(0)
	goto L12
L14:
	;
	v90 = v85
	v91 = v86
	v92 = v87
	goto L24
L15:
	;
	if (v48|v50)&int32(3) != 0 {
		v85 = v48
		v86 = v50
		v87 = v54
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v78 = v48
	v79 = v50
	v80 = v54
	goto L17
L17:
	;
	if v80 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v62 = v48
	v63 = v50
	v64 = v54
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v67 != v68 {
		v85 = v62
		v86 = v63
		v87 = v64
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v78 = v73
	v79 = v71
	v80 = v75
	goto L17
L21:
	;
	v70 = int32(4)
	v71 = v63 + v70
	v73 = v62 + v70
	v75 = v64 - v70
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		v62 = v73
		v63 = v71
		v64 = v75
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v85 = v78
	v86 = v79
	v87 = v80
	goto L14
L24:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 == v96 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v116 = v95 - v96
	goto L12
L26:
	;
	v98 = int32(1)
	v103 = v92 - v98
	if v103 != 0 {
		v90 = v90 + v98
		v91 = v91 + v98
		v92 = v103
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	if v51 != v52 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = int32(10)
	v172 = (v39*v118 + v118) * (v51 - v52)
	goto L4
L32:
	;
	goto L33
L33:
	;
	v125 = v39 - int32(1)
	if v39 < int32(2) {
		v152 = v125
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v128 = int32(9)
	v130 = int32(_a_F_ltree_ne_0)
	v138 = int32(1)
	if v138 < v46 {
		v37 = v37 + (v51+v128)&v130
		v38 = v38 + (v52+v128)&v130
		v39 = v125
		v46 = v46 - v138
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v15)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v178 != v20 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v20)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.B2i32(v172 != int32(0))
L43:
	;
	goto L42
}
func F_ltree_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v2 {
		v31 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v31&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v31 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(7) {
		v31 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 != int32(17) {
		v31 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	v31 = v27 ^ int32(1)
	goto L2
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_get_fn_opclass_options(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v40 = int32(8)
	goto L9
L9:
	;
	v42 = v12 + int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v44&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v39
	goto L9
L12:
	;
	v47 = int32(0)
	goto L14
L13:
	;
	v47 = v40
	goto L14
L14:
	;
	v48 = v42 + v47
	v50 = v10 + int32(8)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v52&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = int32(0)
	goto L17
L16:
	;
	v55 = v40
	goto L17
L17:
	;
	v56 = v50 + v55
	v57 = int32(0)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
	if base.B2i32(v64 == v57)|base.B2i32(v67 == v57) != 0 {
		v131 = v64
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v153&int32(1) != 0 {
		v167 = v50
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v152 = (v131 + int32(1)) * (v64 - v67) * int32(10)
	goto L18
L20:
	;
	v71 = int32(8)
	v75 = v48 + v71
	v76 = v56 + v71
	v79 = v64
	v82 = v67
	goto L21
L21:
	;
	v84 = int32(2)
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75))))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76))))
	if base.Ui32(v88) < base.Ui32(v89) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v131 = v111
	goto L19
L23:
	;
	v111 = v79 - int32(1)
	if v79 < int32(2) {
		v131 = v111
		goto L19
	} else {
		goto L34
	}
L24:
	;
	v91 = v88
	goto L26
L25:
	;
	v91 = v89
	goto L26
L26:
	;
	v92 = F_memcmp(m, v75+v84, v76+v84, v91)
	mBase = m.M
	if v92 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v88 == v89 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v92 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v96 = int32(10)
	v152 = (v79*v96 + v96) * (v88 - v89)
	goto L18
L31:
	;
	v108 = int32(-10)
	goto L33
L32:
	;
	v108 = int32(10)
	goto L33
L33:
	;
	v152 = (v79 + int32(1)) * v108
	goto L18
L34:
	;
	v114 = int32(9)
	v116 = int32(_a_F_ltree_penalty_0)
	v124 = int32(1)
	if v124 < v82 {
		v75 = v75 + (v88+v114)&v116
		v76 = v76 + (v89+v114)&v116
		v79 = v111
		v82 = v82 - v124
		goto L21
	} else {
		goto L35
	}
L35:
	;
	goto L22
L36:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v168&int32(1) != 0 {
		v182 = v42
		goto L42
	} else {
		goto L43
	}
L37:
	;
	if v153&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v159 = int32(0)
	goto L40
L39:
	;
	v159 = v40
	goto L40
L40:
	;
	v160 = v50 + v159
	if v153&int32(4) != 0 {
		v167 = v160
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v167 = v160 + int32(base.Ui32(v163)>>(uint(int32(2))%32))
	goto L36
L42:
	;
	v183 = int32(0)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+4)))
	if base.B2i32(v190 == v183)|base.B2i32(v193 == v183) != 0 {
		v257 = v190
		goto L49
	} else {
		goto L50
	}
L43:
	;
	if v168&int32(2) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v174 = int32(0)
	goto L46
L45:
	;
	v174 = v40
	goto L46
L46:
	;
	v175 = v42 + v174
	if v168&int32(4) != 0 {
		v182 = v175
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v182 = v175 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
	goto L42
L48:
	;
	v279 = int32(0)
	if v279 < v278 {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	v278 = (v257 + int32(1)) * (v190 - v193) * int32(10)
	goto L48
L50:
	;
	v197 = int32(8)
	v201 = v167 + v197
	v202 = v182 + v197
	v205 = v190
	v208 = v193
	goto L51
L51:
	;
	v210 = int32(2)
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202))))
	if base.Ui32(v214) < base.Ui32(v215) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v257 = v237
	goto L49
L53:
	;
	v237 = v205 - int32(1)
	if v205 < int32(2) {
		v257 = v237
		goto L49
	} else {
		goto L64
	}
L54:
	;
	v217 = v214
	goto L56
L55:
	;
	v217 = v215
	goto L56
L56:
	;
	v218 = F_memcmp(m, v201+v210, v202+v210, v217)
	mBase = m.M
	if v218 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v214 == v215 {
		goto L53
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v218 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v222 = int32(10)
	v278 = (v205*v222 + v222) * (v214 - v215)
	goto L48
L61:
	;
	v234 = int32(-10)
	goto L63
L62:
	;
	v234 = int32(10)
	goto L63
L63:
	;
	v278 = (v205 + int32(1)) * v234
	goto L48
L64:
	;
	v240 = int32(9)
	v242 = int32(_a_F_ltree_penalty_0)
	v250 = int32(1)
	if v250 < v208 {
		v201 = v201 + (v214+v240)&v242
		v202 = v202 + (v215+v240)&v242
		v205 = v237
		v208 = v208 - v250
		goto L51
	} else {
		goto L65
	}
L65:
	;
	goto L52
L66:
	;
	v282 = v278
	goto L68
L67:
	;
	v282 = v279
	goto L68
L68:
	;
	v283 = int32(0)
	if v283 < v152 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = v152
	goto L71
L70:
	;
	v286 = v283
	goto L71
L71:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v8))) = base.F32_convert_i32_u(v282 + v286)
	return v8
}
func F_ltree_prefix_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	if base.Ui32(l1) <= base.Ui32(l3) {
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
	v52 = int32(1)
	goto L3
L3:
	;
	return base.B2i32(v52 == int32(0))
L4:
	;
	v52 = v50
	goto L3
L5:
	;
	v50 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v12 = l0
	v13 = l2
	v14 = l1
	v15 = v11
	goto L12
L9:
	;
	v38 = l2
	v42 = int32(0)
	goto L10
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v50 = v42 - v43
	goto L4
L11:
	;
	v38 = v33
	v42 = v35
	goto L10
L12:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v15 != v17)|base.B2i32(v17 == int32(0)) != 0 {
		v33 = v13
		v35 = v15
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v33 = v27
	v35 = int32(0)
	goto L11
L14:
	;
	v23 = v14 - int32(1)
	if v23 == int32(0) {
		v33 = v13
		v35 = v15
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v26 = int32(1)
	v27 = v13 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v28 != 0 {
		v12 = v12 + v26
		v13 = v27
		v14 = v23
		v15 = v28
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
}
func F_parse_ltree(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == v3 {
		v66 = v15
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v290
L2:
	;
	v76 = F_palloc(m, v66<<(uint(int32(4))%32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L15
	}
L3:
	;
	v22 = l0
	v23 = v3
	goto L4
L4:
	;
	v30 = F_pg_mblen_cstr(m, v22)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v41 = v37 + int32(1)
	if base.Ui32(v37) < base.Ui32(int32(_a_F_parse_ltree_0)) {
		v66 = v41
		goto L2
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v37 = v23 + base.B2i32(v34 == int32(46))
	v38 = v22 + v30
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 != 0 {
		v22 = v38
		v23 = v37
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v44 = F_errsave_start(m, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v44 == int32(0) {
		v290 = v3
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = int32(_a_F_parse_ltree_0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
	F_errmsg(m, int32(_a_F_parse_ltree_1), v13+int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(67), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v290 = v3
	goto L1
L15:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v78 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v230 = v227 + int32(8)
	v231 = F_palloc0(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L64
	}
L17:
	;
	v219 = v76
	v227 = v3
	goto L16
L18:
	;
	goto L19
L19:
	;
	v82 = l0
	v84 = v76
	v85 = int32(0)
	v89 = v15
	v90 = v3
	goto L20
L20:
	;
	v92 = F_pg_mblen_cstr(m, v82)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	if v172 != 0 {
		goto L52
	} else {
		goto L53
	}
L22:
	;
	if v85 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	v176 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+12)) = v175 + v176
	v180 = v89 + v176
	v181 = v82 + v92
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v182 != 0 {
		v82 = v181
		v84 = v171
		v85 = v172
		v89 = v180
		v90 = v174
		goto L20
	} else {
		goto L51
	}
L24:
	;
	v96 = F_t_isalnum_cstr(m, v82)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v125 == int32(46) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v108 = int32(0)
	v109 = F_errsave_start(m, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v82
	v171 = v84
	v172 = int32(1)
	v174 = v90
	goto L23
L29:
	;
	if v96 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v98 == int32(95) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if v98 != int32(45) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	if v109 == int32(0) {
		v290 = v108
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v89
	F_errmsg(m, int32(_a_F_parse_ltree_4), v13)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(84), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v290 = v108
	goto L1
L38:
	;
	v128 = int32(0)
	v130 = F_finish_nodeitem(m, v84, v82, v128, v89, l1)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v143 = int32(1)
	v144 = F_t_isalnum_cstr(m, v82)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	if v130 == int32(0) {
		v290 = v128
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v171 = v84 + int32(16)
	v172 = int32(0)
	v174 = (v134+int32(9))&int32(-8) + v90
	goto L23
L43:
	;
	if v144 != 0 {
		v171 = v84
		v172 = v143
		v174 = v90
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.B2i32(v146 == int32(45))|base.B2i32(v146 == int32(95)) != 0 {
		v171 = v84
		v172 = v143
		v174 = v90
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v152 = int32(0)
	v153 = F_errsave_start(m, l1)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v153 == int32(0) {
		v290 = v152
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v89
	F_errmsg(m, int32(_a_F_parse_ltree_4), v13+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(96), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v290 = v152
	goto L1
L51:
	;
	goto L21
L52:
	;
	v183 = int32(0)
	v185 = F_finish_nodeitem(m, v171, v181, v183, v180, l1)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v76 == v171 {
		v219 = v76
		v227 = v174
		goto L16
	} else {
		goto L57
	}
L55:
	;
	if v185 == int32(0) {
		v290 = v183
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v219 = v171 + int32(16)
	v227 = (v191+int32(9))&int32(-8) + v174
	goto L16
L57:
	;
	v198 = int32(0)
	v199 = F_errsave_start(m, l1)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v199 == int32(0) {
		v290 = v198
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_parse_ltree_5), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errdetail(m, int32(_a_F_parse_ltree_6), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(118), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v290 = v198
	goto L1
L64:
	;
	v233 = v219 - v76
	v235 = int32(base.Ui32(v233) >> (uint(int32(4)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+4)) = uint16(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v230 << (uint(int32(2)) % 32)
	if v233&int32(_a_F_parse_ltree_7) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v244 = v76
	v246 = v231 + int32(8)
	goto L68
L66:
	;
	goto L67
L67:
	;
	F_pfree(m, v76)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L74
	}
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v254)
	if v254 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	base.MemoryCopy(m, v246+int32(2), v258, v254)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+4)))
	v269 = v244 + int32(16)
	if (v269-v76)>>(uint(int32(4))%32) < v267 {
		v244 = v269
		v246 = v246 + (v254&int32(_a_F_parse_ltree_0)+int32(9))&int32(_a_F_parse_ltree_8)
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v290 = v231
	goto L1
}
