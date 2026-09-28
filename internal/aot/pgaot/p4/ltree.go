package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v243 int64
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	v2 = int32(0)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = base.I32_wrap_i64(v12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v2 {
		v32 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v32&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v19 == int32(0) {
		v32 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(7) {
		v32 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v25 != int32(17) {
		v32 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)))
	v32 = v28 ^ int32(1)
	goto L2
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_get_fn_opclass_options(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v41 = int32(28)
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+18)))
	if v43 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	return int64(0)
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v41 = v40
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L10
	} else {
		goto L56
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L52
	}
L14:
	;
	return v243
L15:
	;
	v243 = base.I64_extend_i32_u(v221)
	goto L14
L16:
	;
	v209 = F_palloc(m, int32(24))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L51
	}
L17:
	;
	v46 = F_pg_detoast_datum(m, v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v153&int32(2) != 0 {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v51 = F_ArrayGetNItemsSafe(m, v48, v46+int32(16))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if int32(2) <= v53 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v57 = F_array_contains_nulls(m, v46)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v57 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v59 = int32(0)
	v63 = F_ltree_gist_alloc(m, v59, v59, v41, v59, v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	if v51 <= int32(0) {
		v203 = v63
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v56 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v73 = v56
	goto L29
L28:
	;
	v73 = (v53<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L29
L29:
	;
	v81 = v46 + v73
	v83 = v51
	goto L30
L30:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+4)))
	if v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v203 = v63
	goto L16
L32:
	;
	v93 = v81 + int32(8)
	v94 = v90
	goto L35
L33:
	;
	goto L34
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v149 = int32(1)
	if v149 < v83 {
		v81 = v81 + (int32(base.Ui32(v141)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v83 = v83 - v149
		goto L30
	} else {
		goto L39
	}
L35:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
	v107 = F_ltree_crc32_sz(m, v93+int32(2), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	v109 = base.I32_rem_u_s(v107, v41<<(uint(int32(3))%32))
	v112 = v63 + int32(8) + int32(base.Ui32(v109)>>(uint(int32(3))%32))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v114 = int32(1)
	v118 = v113 | v114<<(uint(v109&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v118)
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
	if base.Ui32(v114) < base.Ui32(v94) {
		v93 = v93 + (v120+int32(9))&int32(_a_F__ltree_compress_0)
		v94 = v94 - v114
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L31
L40:
	;
	v221 = v13
	goto L15
L41:
	;
	goto L42
L42:
	;
	v157 = v42 + int32(8)
	v158 = int32(0)
	if v158 < v41 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v163 = v158
	goto L46
L44:
	;
	goto L45
L45:
	;
	v193 = int32(0)
	v195 = F_ltree_gist_alloc(m, int32(1), v157, v41, v193, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L50
	}
L46:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v157))))
	if v175 != int32(255) {
		v243 = v12 & int64(4294967295)
		goto L14
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v179 = v163 + int32(1)
	if v179 != v41 {
		v163 = v179
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v203 = v195
	goto L16
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_extend_i32_u(v203)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = v215
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v209)+18)) = uint8(v218)
	*(*uint16)(unsafe.Add(mBase, uint32(v209)+16)) = uint16(v217)
	v221 = v209
	goto L15
L52:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F__ltree_compress_1), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F__ltree_compress_2), int32(65), int32(_a_F__ltree_compress_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F__ltree_compress_4), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F__ltree_compress_2), int32(69), int32(_a_F__ltree_compress_3))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltree_isparent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14226(m, l0, int32(_a_F__ltree_isparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F__ltree_same(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14295(m, l0, int32(1), int32(2), int32(28))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
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
	v20 = F_pg_database_locale(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l1 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ltree_crc32_sz[0])) = v20
	goto L3
L6:
	;
	v138 = int32(0)
	goto L8
L7:
	;
	v32 = l0
	v33 = l1
	v34 = int32(-1)
	goto L9
L8:
	;
	m.G0 = v13 + int32(16)
	return v138
L9:
	;
	v43 = v13 + int32(3)
	v45 = F_pg_mblen_range(m, v32, l0+l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	v138 = v114 ^ int32(-1)
	goto L8
L11:
	;
	v123 = v33 - v45
	if int32(0) < v123 {
		v32 = v32 + v45
		v33 = v123
		v34 = v114
		goto L9
	} else {
		goto L22
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ltree_crc32_sz[0]))
	v49 = F_pg_strfold(m, v43, int32(13), v32, v45, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v49 == int32(0) {
		v114 = v34
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v49&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32((v34^v55)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v69 = v63 ^ int32(base.Ui32(v34)>>(uint(int32(8))%32))
	v70 = v13 + int32(4)
	v71 = v49 - int32(1)
	goto L17
L16:
	;
	v69 = v34
	v70 = v43
	v71 = v49
	goto L17
L17:
	;
	if v49 == int32(1) {
		v114 = v69
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v76 = v69
	v78 = v70
	v81 = v71
	goto L19
L19:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v86 = int32(255)
	v88 = int32(2)
	v92 = *(*int32)(unsafe.Add(mBase, uint32((v84^v76)&v86<<(uint(v88)%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v93 = int32(8)
	v95 = v92 ^ int32(base.Ui32(v76)>>(uint(v93)%32))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32((v95^v96)&v86<<(uint(v88)%32))+uint32(_c_F_ltree_crc32_sz[1])))
	v107 = v104 ^ int32(base.Ui32(v95)>>(uint(v93)%32))
	v111 = v81 - v88
	if v111 != 0 {
		v76 = v107
		v78 = v78 + v88
		v81 = v111
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v114 = v107
	goto L11
L21:
	;
	goto L20
L22:
	;
	goto L10
}
func F_ltree_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v157 != v14 {
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v156 = base.I64_extend_i32_u(base.B2i32(v21 != v22))
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v39 = v21
	v40 = v22
	goto L7
L7:
	;
	v45 = int32(2)
	v46 = v35 + v45
	v48 = v36 + v45
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v49) < base.Ui32(v50) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v52 = v49
	goto L11
L10:
	;
	v52 = v50
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v52) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v114|base.B2i32(v49 != v50) != 0 {
		v156 = int64(1)
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v114 = int32(0)
	goto L12
L14:
	;
	v88 = v83
	v89 = v84
	v90 = v85
	goto L24
L15:
	;
	if (v46|v48)&int32(3) != 0 {
		v83 = v46
		v84 = v48
		v85 = v52
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v76 = v46
	v77 = v48
	v78 = v52
	goto L17
L17:
	;
	if v78 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v60 = v46
	v61 = v48
	v62 = v52
	goto L19
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v65 != v66 {
		v83 = v60
		v84 = v61
		v85 = v62
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v76 = v71
	v77 = v69
	v78 = v73
	goto L17
L21:
	;
	v68 = int32(4)
	v69 = v61 + v68
	v71 = v60 + v68
	v73 = v62 - v68
	if base.Ui32(int32(3)) < base.Ui32(v73) {
		v60 = v71
		v61 = v69
		v62 = v73
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v83 = v76
	v84 = v77
	v85 = v78
	goto L14
L24:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v93 == v94 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v114 = v93 - v94
	goto L12
L26:
	;
	v96 = int32(1)
	v101 = v90 - v96
	if v101 != 0 {
		v88 = v88 + v96
		v89 = v89 + v96
		v90 = v101
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
	if v40 < int32(2) {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v119 = int32(1)
	v124 = (v49 + int32(9)) & int32(_a_F_ltree_ne_0)
	if v119 < v39 {
		v35 = v124 + v35
		v36 = v36 + v124
		v39 = v39 - v119
		v40 = v40 - v119
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
L33:
	;
	F_pfree(m, v14)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v161 != v19 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	F_pfree(m, v19)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	return v156
L40:
	;
	goto L39
}
func F_ltree_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int64
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v110 float64
	_ = v110
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v167 float32
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
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 float32
	_ = v199
	var v202 float32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v256 float64
	_ = v256
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v313 float32
	_ = v313
	var v314 float32
	_ = v314
	var v317 float32
	_ = v317
	v2 = int32(0)
	v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)))
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
	return int64(0)
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
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
	if base.B2i32(v65 == v57)|base.B2i32(v68 == v57) != 0 {
		v140 = v65
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v168&int32(1) != 0 {
		v182 = v50
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v167 = base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v65-v68)), float64(10)), base.F64_convert_i32_u(v140+int32(1))))
	goto L18
L20:
	;
	v72 = int32(8)
	v76 = v48 + v72
	v77 = v56 + v72
	v80 = v65
	v83 = v68
	goto L21
L21:
	;
	v86 = int32(2)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76))))
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	if base.Ui32(v90) < base.Ui32(v91) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v140 = v120
	goto L19
L23:
	;
	v120 = v80 - int32(1)
	if v80 < int32(2) {
		v140 = v120
		goto L19
	} else {
		goto L34
	}
L24:
	;
	v93 = v90
	goto L26
L25:
	;
	v93 = v91
	goto L26
L26:
	;
	v94 = F_memcmp(m, v76+v86, v77+v86, v93)
	mBase = m.M
	if v94 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v90 == v91 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v110 = base.F64_convert_i32_u(v80 + int32(1))
	if v94 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v167 = base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v90-v91)), float64(10)), base.F64_convert_i32_u(v80+int32(1))))
	goto L18
L31:
	;
	v167 = base.F32_demote_f64(base.F64_mul(v110, float64(-10)))
	goto L18
L32:
	;
	goto L33
L33:
	;
	v167 = base.F32_demote_f64(base.F64_mul(v110, float64(10)))
	goto L18
L34:
	;
	v123 = int32(9)
	v125 = int32(_a_F_ltree_penalty_0)
	v133 = int32(1)
	if v133 < v83 {
		v76 = v76 + (v90+v123)&v125
		v77 = v77 + (v91+v123)&v125
		v80 = v120
		v83 = v83 - v133
		goto L21
	} else {
		goto L35
	}
L35:
	;
	goto L22
L36:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v183&int32(1) != 0 {
		v197 = v42
		goto L42
	} else {
		goto L43
	}
L37:
	;
	if v168&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v174 = int32(0)
	goto L40
L39:
	;
	v174 = v40
	goto L40
L40:
	;
	v175 = v50 + v174
	if v168&int32(4) != 0 {
		v182 = v175
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v182 = v175 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
	goto L36
L42:
	;
	v199 = float32(0)
	if base.F32_gt(v167, v199) != 0 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v183&int32(2) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v189 = int32(0)
	goto L46
L45:
	;
	v189 = v40
	goto L46
L46:
	;
	v190 = v42 + v189
	if v183&int32(4) != 0 {
		v197 = v190
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v197 = v190 + int32(base.Ui32(v193)>>(uint(int32(2))%32))
	goto L42
L48:
	;
	v202 = v167
	goto L50
L49:
	;
	v202 = v199
	goto L50
L50:
	;
	v203 = int32(0)
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+4)))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	if base.B2i32(v211 == v203)|base.B2i32(v214 == v203) != 0 {
		v286 = v211
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v314 = float32(0)
	if base.F32_gt(v313, v314) != 0 {
		goto L69
	} else {
		goto L70
	}
L52:
	;
	v313 = base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v211-v214)), float64(10)), base.F64_convert_i32_u(v286+int32(1))))
	goto L51
L53:
	;
	v218 = int32(8)
	v222 = v182 + v218
	v223 = v197 + v218
	v226 = v211
	v229 = v214
	goto L54
L54:
	;
	v232 = int32(2)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223))))
	if base.Ui32(v236) < base.Ui32(v237) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v286 = v266
	goto L52
L56:
	;
	v266 = v226 - int32(1)
	if v226 < int32(2) {
		v286 = v266
		goto L52
	} else {
		goto L67
	}
L57:
	;
	v239 = v236
	goto L59
L58:
	;
	v239 = v237
	goto L59
L59:
	;
	v240 = F_memcmp(m, v222+v232, v223+v232, v239)
	mBase = m.M
	if v240 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v236 == v237 {
		goto L56
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v256 = base.F64_convert_i32_u(v226 + int32(1))
	if v240 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v313 = base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v236-v237)), float64(10)), base.F64_convert_i32_u(v226+int32(1))))
	goto L51
L64:
	;
	v313 = base.F32_demote_f64(base.F64_mul(v256, float64(-10)))
	goto L51
L65:
	;
	goto L66
L66:
	;
	v313 = base.F32_demote_f64(base.F64_mul(v256, float64(10)))
	goto L51
L67:
	;
	v269 = int32(9)
	v271 = int32(_a_F_ltree_penalty_0)
	v279 = int32(1)
	if v279 < v229 {
		v222 = v222 + (v236+v269)&v271
		v223 = v223 + (v237+v269)&v271
		v226 = v266
		v229 = v229 - v279
		goto L54
	} else {
		goto L68
	}
L68:
	;
	goto L55
L69:
	;
	v317 = v313
	goto L71
L70:
	;
	v317 = v314
	goto L71
L71:
	;
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8)))) = base.F32_add(v202, v317)
	return v8
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
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
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
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
	return v289
L2:
	;
	v75 = F_palloc_mul(m, int32(16), v66)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
		v289 = v3
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
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(68), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v289 = v3
	goto L1
L15:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v77 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v229 = v226 + int32(8)
	v230 = F_palloc0(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L64
	}
L17:
	;
	v218 = v75
	v226 = v3
	goto L16
L18:
	;
	goto L19
L19:
	;
	v81 = l0
	v83 = v75
	v84 = int32(0)
	v88 = v15
	v89 = v3
	goto L20
L20:
	;
	v91 = F_pg_mblen_cstr(m, v81)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	if v171 != 0 {
		goto L52
	} else {
		goto L53
	}
L22:
	;
	if v84 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+12)) = v174 + v175
	v179 = v88 + v175
	v180 = v81 + v91
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != 0 {
		v81 = v180
		v83 = v170
		v84 = v171
		v88 = v179
		v89 = v173
		goto L20
	} else {
		goto L51
	}
L24:
	;
	v95 = F_t_isalnum_cstr(m, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v124 == int32(46) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v107 = int32(0)
	v108 = F_errsave_start(m, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v81
	v170 = v83
	v171 = int32(1)
	v173 = v89
	goto L23
L29:
	;
	if v95 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v97 == int32(95) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if v97 != int32(45) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	if v108 == int32(0) {
		v289 = v107
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v88
	F_errmsg(m, int32(_a_F_parse_ltree_4), v13)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(85), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v289 = v107
	goto L1
L38:
	;
	v127 = int32(0)
	v129 = F_finish_nodeitem(m, v83, v81, v127, v88, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v142 = int32(1)
	v143 = F_t_isalnum_cstr(m, v81)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	if v129 == int32(0) {
		v289 = v127
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v170 = v83 + int32(16)
	v171 = int32(0)
	v173 = (v133+int32(9))&int32(-8) + v89
	goto L23
L43:
	;
	if v143 != 0 {
		v170 = v83
		v171 = v142
		v173 = v89
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.B2i32(v145 == int32(45))|base.B2i32(v145 == int32(95)) != 0 {
		v170 = v83
		v171 = v142
		v173 = v89
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v151 = int32(0)
	v152 = F_errsave_start(m, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v152 == int32(0) {
		v289 = v151
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v88
	F_errmsg(m, int32(_a_F_parse_ltree_4), v13+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(97), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v289 = v151
	goto L1
L51:
	;
	goto L21
L52:
	;
	v182 = int32(0)
	v184 = F_finish_nodeitem(m, v170, v180, v182, v179, l1)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v75 == v170 {
		v218 = v75
		v226 = v173
		goto L16
	} else {
		goto L57
	}
L55:
	;
	if v184 == int32(0) {
		v289 = v182
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v218 = v170 + int32(16)
	v226 = (v190+int32(9))&int32(-8) + v173
	goto L16
L57:
	;
	v197 = int32(0)
	v198 = F_errsave_start(m, l1)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v198 == int32(0) {
		v289 = v197
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_parse_ltree_5), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v211 = F_errdetail(m, int32(_a_F_parse_ltree_6), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errsave_finish(m, l1, int32(_a_F_parse_ltree_2), int32(119), int32(_a_F_parse_ltree_3))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v289 = v197
	goto L1
L64:
	;
	v232 = v218 - v75
	v234 = int32(base.Ui32(v232) >> (uint(int32(4)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)) = uint16(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v229 << (uint(int32(2)) % 32)
	if v232&int32(_a_F_parse_ltree_7) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v243 = v75
	v245 = v230 + int32(8)
	goto L68
L66:
	;
	goto L67
L67:
	;
	F_pfree(m, v75)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L74
	}
L68:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v253)
	if v253 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	base.MemoryCopy(m, v245+int32(2), v257, v253)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)))
	v268 = v243 + int32(16)
	if (v268-v75)>>(uint(int32(4))%32) < v266 {
		v243 = v268
		v245 = v245 + (v253&int32(_a_F_parse_ltree_0)+int32(9))&int32(_a_F_parse_ltree_8)
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v289 = v230
	goto L1
}
